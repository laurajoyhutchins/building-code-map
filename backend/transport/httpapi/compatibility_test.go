package httpapi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
	"time"

	"building-code-map/backend/geocoder"
	"building-code-map/backend/regulatory"
	"building-code-map/backend/snapshot"
)

func TestLegacyHTTPCompatibilityStatusMatrix(t *testing.T) {
	catalog, err := regulatory.LoadCatalog(filepath.Join("..", "..", "data", "regulatory"))
	if err != nil {
		t.Fatal(err)
	}

	candidate := geocoder.Candidate{
		MatchedAddress:       "1600 N Broadway St Denver CO 80202",
		Longitude:            -104.99,
		Latitude:             39.74,
		Precision:            geocoder.PrecisionAddressPoint,
		Confidence:           1,
		ScoreKind:            "deterministic_quality",
		ScoreFactors:         map[string]float64{"address_point": 1},
		RankingPolicyVersion: "geocoder-ranking-1.0",
		Source:               "fixture",
		SourceRecordID:       "one",
		SourceVintage:        "2026-08-01",
	}

	tests := []struct {
		name    string
		handler http.Handler
		method  string
		path    string
		body    string
		status  int
	}{
		{
			name:    "health",
			handler: NewLegacyHandler(snapshot.Snapshot{}),
			method:  http.MethodGet,
			path:    "/health",
			status:  http.StatusOK,
		},
		{
			name:    "readiness",
			handler: NewLegacyHandler(compatibilitySnapshot()),
			method:  http.MethodGet,
			path:    "/ready",
			status:  http.StatusOK,
		},
		{
			name:    "layers",
			handler: NewLegacyHandler(compatibilitySnapshot()),
			method:  http.MethodGet,
			path:    "/layers",
			status:  http.StatusOK,
		},
		{
			name:    "boundaries",
			handler: NewLegacyHandler(compatibilitySnapshot()),
			method:  http.MethodGet,
			path:    "/boundaries",
			status:  http.StatusOK,
		},
		{
			name:    "feature lookup",
			handler: NewLegacyHandler(compatibilitySnapshot()),
			method:  http.MethodGet,
			path:    "/features/states/08",
			status:  http.StatusOK,
		},
		{
			name: "geocode success",
			handler: NewLegacyHandler(snapshot.Snapshot{}, Options{Geocoder: fakeGeocoder{result: geocoder.Result{
				Query: "1600 N Broadway, Denver, CO", Status: geocoder.StatusMatched,
				Selected: &candidate, Candidates: []geocoder.Candidate{candidate}, Warnings: []string{},
			}}}),
			method: http.MethodPost, path: "/geocode", body: `{"address":"1600 N Broadway, Denver, CO"}`, status: http.StatusOK,
		},
		{
			name: "geocode failure",
			handler: NewLegacyHandler(snapshot.Snapshot{}, Options{Geocoder: fakeGeocoder{result: geocoder.Result{
				Query: "unknown", Status: geocoder.StatusNotFound, Candidates: []geocoder.Candidate{}, Warnings: []string{},
			}}}),
			method: http.MethodPost, path: "/geocode", body: `{"address":"unknown"}`, status: http.StatusUnprocessableEntity,
		},
		{
			name: "lookup success",
			handler: NewLegacyHandler(compatibilitySnapshot(), Options{
				RegulatoryCatalog: catalog,
				Geocoder: fakeGeocoder{result: geocoder.Result{
					Query: "1600 N Broadway, Denver, CO", Status: geocoder.StatusMatched,
					Selected: &candidate, Candidates: []geocoder.Candidate{candidate}, Warnings: []string{},
				}},
			}),
			method: http.MethodPost, path: "/lookup", body: `{"address":"1600 N Broadway, Denver, CO","applicability_date":"2026-08-06"}`, status: http.StatusOK,
		},
		{
			name: "lookup failure",
			handler: NewLegacyHandler(compatibilitySnapshot(), Options{
				RegulatoryCatalog: catalog,
				Geocoder: fakeGeocoder{result: geocoder.Result{
					Query: "unknown", Status: geocoder.StatusNotFound, Candidates: []geocoder.Candidate{}, Warnings: []string{},
				}},
			}),
			method: http.MethodPost, path: "/lookup", body: `{"address":"unknown","applicability_date":"2026-08-06"}`, status: http.StatusUnprocessableEntity,
		},
		{
			name:    "resolve success",
			handler: NewLegacyHandler(compatibilitySnapshot(), Options{RegulatoryCatalog: catalog}),
			method:  http.MethodPost,
			path:    "/resolve",
			body:    `{"point":{"longitude":-104.99,"latitude":39.74},"applicability_date":"2026-08-06"}`,
			status:  http.StatusOK,
		},
		{
			name:    "malformed JSON",
			handler: NewLegacyHandler(compatibilitySnapshot(), Options{RegulatoryCatalog: catalog}),
			method:  http.MethodPost,
			path:    "/resolve",
			body:    `{"point":`,
			status:  http.StatusBadRequest,
		},
		{
			name:    "invalid coordinates",
			handler: NewLegacyHandler(compatibilitySnapshot(), Options{RegulatoryCatalog: catalog}),
			method:  http.MethodPost,
			path:    "/resolve",
			body:    `{"point":{"longitude":181,"latitude":39.74},"applicability_date":"2026-08-06"}`,
			status:  http.StatusUnprocessableEntity,
		},
		{
			name:    "missing regulatory catalog",
			handler: NewLegacyHandler(compatibilitySnapshot()),
			method:  http.MethodPost,
			path:    "/resolve",
			body:    `{"point":{"longitude":-104.99,"latitude":39.74},"applicability_date":"2026-08-06"}`,
			status:  http.StatusServiceUnavailable,
		},
		{
			name:    "unsupported state coverage",
			handler: NewLegacyHandler(snapshot.Snapshot{}, Options{RegulatoryCatalog: catalog}),
			method:  http.MethodPost,
			path:    "/resolve",
			body:    `{"point":{"longitude":-104.99,"latitude":39.74},"applicability_date":"2026-08-06"}`,
			status:  http.StatusUnprocessableEntity,
		},
		{
			name: "boundary ambiguity",
			handler: NewLegacyHandler(snapshot.Snapshot{BoundaryFeatures: []snapshot.BoundaryFeature{
				{LayerFamily: "states", FeatureID: "08", Title: "Colorado", Attributes: map[string]any{"STATEFP": "08"}, Geometry: testPolygon(0, 0, 10, 10)},
				{LayerFamily: "states", FeatureID: "12", Title: "Florida", Attributes: map[string]any{"STATEFP": "12"}, Geometry: testPolygon(0, 0, 10, 10)},
			}}, Options{RegulatoryCatalog: catalog}),
			method: http.MethodPost, path: "/resolve", body: `{"point":{"longitude":5,"latitude":5},"applicability_date":"2026-08-06"}`, status: http.StatusConflict,
		},
		{
			name:    "refresh status",
			handler: NewLegacyHandler(snapshot.Snapshot{}),
			method:  http.MethodGet,
			path:    "/refresh/status",
			status:  http.StatusOK,
		},
		{
			name:    "disabled refresh trigger",
			handler: NewLegacyHandler(snapshot.Snapshot{}),
			method:  http.MethodPost,
			path:    "/refresh/trigger",
			status:  http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(test.method, test.path, bytes.NewBufferString(test.body))
			rec := httptest.NewRecorder()
			test.handler.ServeHTTP(rec, req)
			if rec.Code != test.status {
				t.Fatalf("status=%d body=%s, want %d", rec.Code, rec.Body.String(), test.status)
			}
			if test.status != http.StatusNoContent {
				var payload any
				if err := json.Unmarshal(rec.Body.Bytes(), &payload); err != nil {
					t.Fatalf("response is not JSON: %v", err)
				}
			}
		})
	}
}

func compatibilitySnapshot() snapshot.Snapshot {
	return snapshot.Snapshot{
		LayerFamilies: []snapshot.LayerFamily{{Key: "states", Label: "States"}},
		BoundaryFeatures: []snapshot.BoundaryFeature{{
			LayerFamily: "states", FeatureID: "08", Title: "Colorado", SourceID: "GEOID=08",
			GeometryLabel: "Polygon", LastSyncedAt: time.Date(2026, 8, 6, 0, 0, 0, 0, time.UTC),
			Attributes: map[string]any{"STATEFP": "08"}, Geometry: testPolygon(-109, 37, -102, 41),
		}},
		RefreshStatus: snapshot.RefreshStatus{Status: "ok", LatestAttempt: time.Date(2026, 8, 6, 0, 0, 0, 0, time.UTC)},
	}
}

var _ geocoder.Service = fakeGeocoder{}
