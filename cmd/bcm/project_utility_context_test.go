package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type geoFixture struct {
	geocodeStatus string
	precision     string
	countyName    string
	countySource  string
	stateFIPS     string
	bundle        string
	pointDrift    bool
	unavailable   bool
	calls         []string
}

func mockGeography(t *testing.T, g *geoFixture) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		g.calls = append(g.calls, r.URL.Path)
		if r.Header.Get("Authorization") != "Bearer provider-test-token" {
			t.Errorf("request to %s lacks provider identity", r.URL.Path)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		switch r.URL.Path {
		case "/v1/readiness":
			status := "available"
			if g.unavailable {
				status = "unavailable"
			}
			_, _ = fmt.Fprintf(w, `{"status":"ok","capabilities":{"address_lookup":{"status":%q},"coordinate_resolution":{"status":%q}}}`, status, status)
		case "/v1/bundle":
			_, _ = w.Write([]byte(`{"bundle_manifest_digest":"sha256:reviewed"}`))
		case "/v1/geocode":
			var request struct {
				Address string `json:"address"`
			}
			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
				t.Error(err)
			}
			if request.Address != "100 Main St, Denver, CO" {
				t.Errorf("unbound address: %q", request.Address)
			}
			status := g.geocodeStatus
			if status == "" {
				status = "matched"
			}
			precision := g.precision
			if precision == "" {
				precision = "address_point"
			}
			_, _ = fmt.Fprintf(w, `{"status":%q,"selected":{"longitude":-104.99,"latitude":39.74,"precision":%q,"source_record_id":"address-row-1"}}`, status, precision)
		case "/v1/lookup":
			var request struct {
				Point struct {
					Longitude float64 `json:"longitude"`
					Latitude  float64 `json:"latitude"`
				} `json:"point"`
			}
			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
				t.Error(err)
			}
			longitude := request.Point.Longitude
			if g.pointDrift {
				longitude += 0.1
			}
			county := g.countyName
			if county == "" {
				county = "Denver County"
			}
			source := g.countySource
			if source == "" {
				source = "co-county-boundaries-2025"
			}
			fips := g.stateFIPS
			if fips == "" {
				fips = "08"
			}
			bundle := g.bundle
			if bundle == "" {
				bundle = "sha256:reviewed"
			}
			_, _ = fmt.Fprintf(w, `{"point":{"longitude":%g,"latitude":%g},"geography":{"census_state_equivalent_fips":%q,"county":{"layer_family":"counties","feature_id":"08031","name":%q,"source_id":%q}},"provenance":{"bundle_manifest_digest":%q,"boundary_snapshot_digest":"boundary-sha","geocoder_snapshot_digest":"geocoder-sha"}}`, longitude, request.Point.Latitude, fips, county, source, bundle)
		default:
			t.Errorf("unexpected private endpoint %s", r.URL.Path)
			http.NotFound(w, r)
		}
	}))
}

func runProjectGeo(t *testing.T, g *geoFixture, body string) *httptest.ResponseRecorder {
	t.Helper()
	server := mockGeography(t, g)
	defer server.Close()
	remote, err := newPrivateProvider(server.URL, "provider-test-token")
	if err != nil {
		t.Fatal(err)
	}
	remote.expectedBundle = "sha256:reviewed"
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, ProjectUtilityContextPath, strings.NewReader(body))
	projectUtilityContextHandler(remote, testUtilityContext(t)).ServeHTTP(recorder, request)
	return recorder
}

const validGeoBody = `{"project_id":"SITE-1","address":"100 Main St, Denver, CO"}`

func TestProjectUtilityContextResolvesSourceBoundCountyWithoutRegulatoryClaim(t *testing.T) {
	g := &geoFixture{}
	rec := runProjectGeo(t, g, validGeoBody)
	if rec.Code != http.StatusOK {
		t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
	}
	var response projectUtilityResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}
	if response.ProjectID != "SITE-1" || response.Status != "geography_observed_utility_candidates_unconfirmed" || response.Geography.County != "Denver County" ||
		response.Geography.StateFIPS != "08" || response.Geography.ProviderBundleDigest != "sha256:reviewed" || response.Geography.CountySourceID == "" ||
		response.Geography.AddressSourceID != "address-row-1" || response.UtilityContext.Status != "reported_county_candidates_not_site_verified" || len(response.UtilityContext.Candidates) != 1 {
		t.Fatalf("geographic evidence not preserved: %+v", response)
	}
	if bytes.Contains(rec.Body.Bytes(), []byte("100 Main")) || bytes.Contains(rec.Body.Bytes(), []byte(`"verdict"`)) {
		t.Fatal("response leaked address or crossed regulatory verdict boundary")
	}
	if strings.Join(g.calls, ",") != "/v1/readiness,/v1/bundle,/v1/geocode,/v1/lookup" {
		t.Fatalf("wrong authority sequence: %+v", g.calls)
	}
}

func TestProjectUtilityContextFailsClosedOnBadGeographicEvidence(t *testing.T) {
	cases := []struct {
		name    string
		fixture geoFixture
		code    int
		calls   int
	}{
		{"ambiguous", geoFixture{geocodeStatus: "ambiguous"}, 422, 3},
		{"interpolated", geoFixture{precision: "interpolated"}, 422, 3},
		{"state mismatch", geoFixture{stateFIPS: "06"}, 422, 4},
		{"bundle mismatch", geoFixture{bundle: "sha256:stale"}, 422, 4},
		{"coordinate mismatch", geoFixture{pointDrift: true}, 422, 4},
		{"no county name", geoFixture{countyName: "  "}, 422, 4},
		{"unavailable provider", geoFixture{unavailable: true}, 503, 2},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			g := tt.fixture
			rec := runProjectGeo(t, &g, validGeoBody)
			if rec.Code != tt.code {
				t.Fatalf("got=%d want=%d body=%s", rec.Code, tt.code, rec.Body.String())
			}
			if len(g.calls) != tt.calls {
				t.Fatalf("called %v", g.calls)
			}
			if strings.Contains(rec.Body.String(), "100 Main") || strings.Contains(rec.Body.String(), "candidates") {
				t.Fatalf("error leaked request/evidence: %s", rec.Body.String())
			}
		})
	}
}

func TestProjectUtilityContextRejectsInvalidTransportBeforePrivateCalls(t *testing.T) {
	for _, body := range []string{
		`{"project_id":"SITE-1","address":"foo"}`,
		`{"project_id":"SITE-1","address":"100 Main St, Denver, CO","county":"Denver"}`,
		`{"project_id":"SITE-1","address":"100 Main St, Denver, CO"}{}`,
		`{"project_id":"bad space","address":"100 Main St, Denver, CO"}`,
		strings.Repeat("A", 5000),
	} {
		g := &geoFixture{}
		rec := runProjectGeo(t, g, body)
		if rec.Code != http.StatusBadRequest || len(g.calls) != 0 {
			t.Fatalf("invalid request admitted: %d calls=%v", rec.Code, g.calls)
		}
	}
}
