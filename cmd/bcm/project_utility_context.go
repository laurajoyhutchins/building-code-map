package main

import (
	"encoding/json"
	"errors"
	"io"
	"math"
	"net/http"
	"regexp"
	"strings"

	"github.com/laurajoyhutchins/building-code-map/geocoder"
	"github.com/laurajoyhutchins/building-code-map/utilitycontext"
)

// ProjectUtilityContextPath is a non-regulatory, source-bound lookup. The
// address is never exposed in the response or recorded in application logs.
const ProjectUtilityContextPath = "/v1/project-utility-context"

var projectIdentity = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9_.:-]{0,119}$`)

// Census state FIPS is checked against the caller's parsed state only AFTER a
// real boundary lookup; an address string never supplies county evidence.
var stateFIPS = map[string]string{
	"AL": "01", "AK": "02", "AZ": "04", "AR": "05", "CA": "06", "CO": "08",
	"CT": "09", "DE": "10", "DC": "11", "FL": "12", "GA": "13", "HI": "15",
	"ID": "16", "IL": "17", "IN": "18", "IA": "19", "KS": "20", "KY": "21",
	"LA": "22", "ME": "23", "MD": "24", "MA": "25", "MI": "26", "MN": "27",
	"MS": "28", "MO": "29", "MT": "30", "NE": "31", "NV": "32", "NH": "33",
	"NJ": "34", "NM": "35", "NY": "36", "NC": "37", "ND": "38", "OH": "39",
	"OK": "40", "OR": "41", "PA": "42", "RI": "44", "SC": "45", "SD": "46",
	"TN": "47", "TX": "48", "UT": "49", "VT": "50", "VA": "51", "WA": "53",
	"WV": "54", "WI": "55", "WY": "56", "AS": "60", "GU": "66", "MP": "69",
	"PR": "72", "VI": "78",
}

type projectUtilityRequest struct {
	ProjectID string `json:"project_id"`
	Address   string `json:"address"`
}

type geographicObservation struct {
	State                  string `json:"state"`
	StateFIPS              string `json:"state_fips"`
	County                 string `json:"county"`
	CountyFeatureID        string `json:"county_feature_id"`
	CountySourceID         string `json:"county_source_id"`
	AddressPrecision       string `json:"address_precision"`
	AddressSourceID        string `json:"address_source_id"`
	GeocoderSnapshotSHA256 string `json:"geocoder_snapshot_sha256"`
	BoundarySnapshotSHA256 string `json:"boundary_snapshot_sha256"`
	ProviderBundleDigest   string `json:"provider_bundle_digest"`
}

type projectUtilityResponse struct {
	ProjectID      string                  `json:"project_id"`
	Status         string                  `json:"status"`
	Geography      geographicObservation   `json:"geography"`
	UtilityContext utilitycontext.Response `json:"utility_context"`
}

// Attach to the existing configured private provider, rather than introducing
// a second geocoder, provider, live database, or regulatory verdict path.
func projectUtilityContextHandler(remote *privateProvider, catalog *utilitycontext.Catalog) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Cache-Control", "no-store")
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			projectUtilityError(w, http.StatusMethodNotAllowed, "method_not_allowed")
			return
		}
		if r.URL.Path != ProjectUtilityContextPath || r.URL.RawQuery != "" {
			projectUtilityError(w, http.StatusBadRequest, "invalid_request")
			return
		}
		r.Body = http.MaxBytesReader(w, r.Body, 4096)
		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()
		var input projectUtilityRequest
		if err := decoder.Decode(&input); err != nil {
			projectUtilityError(w, http.StatusBadRequest, "invalid_request")
			return
		}
		var extra any
		if err := decoder.Decode(&extra); !errors.Is(err, io.EOF) {
			projectUtilityError(w, http.StatusBadRequest, "invalid_request")
			return
		}
		if !projectIdentity.MatchString(input.ProjectID) || len(input.Address) > 512 {
			projectUtilityError(w, http.StatusBadRequest, "invalid_request")
			return
		}
		parsed, err := geocoder.ParseAddress(input.Address)
		if err != nil {
			projectUtilityError(w, http.StatusBadRequest, "invalid_address")
			return
		}
		if remote == nil || catalog == nil || remote.expectedBundle == "" {
			projectUtilityError(w, http.StatusServiceUnavailable, "source_unavailable")
			return
		}
		ready, err := remote.readiness(r.Context())
		if err != nil || ready.Status == "not_ready" ||
			ready.Capabilities["address_lookup"].Status != "available" ||
			ready.Capabilities["coordinate_resolution"].Status != "available" {
			projectUtilityError(w, http.StatusServiceUnavailable, "geography_unavailable")
			return
		}

		var geocode struct {
			Status   string `json:"status"`
			Selected *struct {
				Longitude      float64 `json:"longitude"`
				Latitude       float64 `json:"latitude"`
				Precision      string  `json:"precision"`
				SourceRecordID string  `json:"source_record_id"`
			} `json:"selected"`
		}
		if _, err := remote.request(r.Context(), http.MethodPost, "/v1/geocode", map[string]string{"address": input.Address}, &geocode); err != nil {
			projectUtilityError(w, http.StatusUnprocessableEntity, "address_not_resolved")
			return
		}
		if geocode.Status != "matched" || geocode.Selected == nil ||
			geocode.Selected.Precision != "address_point" ||
			geocode.Selected.SourceRecordID == "" ||
			!validPoint(geocode.Selected.Longitude, geocode.Selected.Latitude) {
			projectUtilityError(w, http.StatusUnprocessableEntity, "address_not_resolved")
			return
		}
		point := struct {
			Longitude float64 `json:"longitude"`
			Latitude  float64 `json:"latitude"`
		}{geocode.Selected.Longitude, geocode.Selected.Latitude}
		var lookup struct {
			Point struct {
				Longitude float64 `json:"longitude"`
				Latitude  float64 `json:"latitude"`
			} `json:"point"`
			Geography struct {
				CensusStateEquivalentFIPS string `json:"census_state_equivalent_fips"`
				County                    *struct {
					LayerFamily string `json:"layer_family"`
					FeatureID   string `json:"feature_id"`
					Name        string `json:"name"`
					SourceID    string `json:"source_id"`
				} `json:"county"`
			} `json:"geography"`
			Provenance struct {
				BundleManifestDigest   string `json:"bundle_manifest_digest"`
				BoundarySnapshotDigest string `json:"boundary_snapshot_digest"`
				GeocoderSnapshotDigest string `json:"geocoder_snapshot_digest"`
			} `json:"provenance"`
		}
		if _, err := remote.request(r.Context(), http.MethodPost, "/v1/lookup", map[string]any{"point": point}, &lookup); err != nil {
			projectUtilityError(w, http.StatusUnprocessableEntity, "county_not_resolved")
			return
		}
		county := lookup.Geography.County
		if lookup.Point != point || county == nil ||
			county.LayerFamily != "counties" || county.FeatureID == "" ||
			county.SourceID == "" || strings.TrimSpace(county.Name) == "" ||
			lookup.Geography.CensusStateEquivalentFIPS != stateFIPS[parsed.State] ||
			lookup.Provenance.BundleManifestDigest != remote.expectedBundle ||
			lookup.Provenance.BoundarySnapshotDigest == "" ||
			lookup.Provenance.GeocoderSnapshotDigest == "" {
			projectUtilityError(w, http.StatusUnprocessableEntity, "county_not_resolved")
			return
		}
		// Only a known suffix is removed. Original county source identity is retained.
		countyKey := strings.TrimSpace(county.Name)
		if strings.HasSuffix(countyKey, " County") {
			countyKey = strings.TrimSuffix(countyKey, " County")
		}
		candidates, err := catalog.Lookup(parsed.State, countyKey)
		if err != nil {
			projectUtilityError(w, http.StatusUnprocessableEntity, "county_not_resolved")
			return
		}
		response := projectUtilityResponse{
			ProjectID: input.ProjectID,
			Status:    "geography_observed_utility_candidates_unconfirmed",
			Geography: geographicObservation{
				State:                  parsed.State,
				StateFIPS:              lookup.Geography.CensusStateEquivalentFIPS,
				County:                 county.Name,
				CountyFeatureID:        county.FeatureID,
				CountySourceID:         county.SourceID,
				AddressPrecision:       geocode.Selected.Precision,
				AddressSourceID:        geocode.Selected.SourceRecordID,
				GeocoderSnapshotSHA256: lookup.Provenance.GeocoderSnapshotDigest,
				BoundarySnapshotSHA256: lookup.Provenance.BoundarySnapshotDigest,
				ProviderBundleDigest:   lookup.Provenance.BundleManifestDigest,
			},
			UtilityContext: candidates,
		}
		_ = json.NewEncoder(w).Encode(response)
	})
}

func projectUtilityError(w http.ResponseWriter, status int, code string) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": code})
}

func validPoint(longitude, latitude float64) bool {
	return !math.IsNaN(longitude) && !math.IsInf(longitude, 0) &&
		!math.IsNaN(latitude) && !math.IsInf(latitude, 0) &&
		longitude >= -180 && longitude <= 180 && latitude >= -90 && latitude <= 90
}
