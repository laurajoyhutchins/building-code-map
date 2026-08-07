package httpapi_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"

	"building-code-map/backend/internal/httpapi"
	"building-code-map/backend/internal/snapshot"
)

func TestBoundariesReturnOnlyMapContractFields(t *testing.T) {
	handler := boundaryContractHandler()
	response := requestBoundaryContract(t, handler, "/boundaries")

	var records []map[string]any
	if err := json.Unmarshal(response.Body.Bytes(), &records); err != nil {
		t.Fatalf("decode /boundaries response: %v", err)
	}
	if len(records) != 1 {
		t.Fatalf("/boundaries returned %d records, want 1", len(records))
	}

	wantKeys := []string{
		"feature_id",
		"geometry",
		"geometry_label",
		"geometry_source",
		"layer_family",
		"source_id",
		"subtitle",
		"title",
	}
	if got := sortedMapKeys(records[0]); !reflect.DeepEqual(got, wantKeys) {
		t.Fatalf("/boundaries keys = %v, want %v", got, wantKeys)
	}
	if _, ok := records[0]["geometry"]; !ok {
		t.Fatal("/boundaries omitted geometry required for map rendering")
	}
}

func TestFeatureDetailReturnsOnlyDetailContractFields(t *testing.T) {
	handler := boundaryContractHandler()
	response := requestBoundaryContract(t, handler, "/features/states/08")

	var record map[string]any
	if err := json.Unmarshal(response.Body.Bytes(), &record); err != nil {
		t.Fatalf("decode feature response: %v", err)
	}

	wantKeys := []string{
		"attributes",
		"feature_id",
		"geometry_label",
		"geometry_source",
		"last_synced_at",
		"layer_family",
		"source_id",
		"subtitle",
		"title",
	}
	if got := sortedMapKeys(record); !reflect.DeepEqual(got, wantKeys) {
		t.Fatalf("feature detail keys = %v, want %v", got, wantKeys)
	}
	if _, ok := record["geometry"]; ok {
		t.Fatal("feature detail unexpectedly duplicated map geometry")
	}
	attributes, ok := record["attributes"].(map[string]any)
	if !ok || attributes["detail_sentinel"] != "detail-only" {
		t.Fatalf("feature detail attributes = %#v, want detail sentinel", record["attributes"])
	}
}

func TestBoundaryPayloadSizeDoesNotScaleWithRawAttributes(t *testing.T) {
	handler := boundaryContractHandler()
	response := requestBoundaryContract(t, handler, "/boundaries")

	if strings.Contains(response.Body.String(), "detail-only") {
		t.Fatal("/boundaries leaked detail attributes")
	}
	if response.Body.Len() >= 4096 {
		t.Fatalf("/boundaries response is %d bytes, want less than 4096 for one compact map record", response.Body.Len())
	}
}

func boundaryContractHandler() http.Handler {
	return httpapi.NewHandler(snapshot.Snapshot{
		BoundaryFeatures: []snapshot.BoundaryFeature{
			{
				LayerFamily:    "states",
				FeatureID:      "08",
				Title:          "Colorado",
				Subtitle:       "State boundary",
				SourceID:       "GEOID=08",
				GeometryLabel:  "Polygon",
				GeometrySource: "tigerweb_live",
				LastSyncedAt:   time.Date(2026, 6, 22, 12, 0, 0, 0, time.UTC),
				Geometry: snapshot.Geometry{
					Type: "Polygon",
					Coordinates: []any{
						[]any{
							[]float64{-109.06, 37.0},
							[]float64{-102.04, 37.0},
							[]float64{-102.04, 41.0},
							[]float64{-109.06, 41.0},
							[]float64{-109.06, 37.0},
						},
					},
				},
				Attributes: map[string]any{
					"detail_sentinel": "detail-only",
					"large_raw_value": strings.Repeat("x", 64*1024),
				},
			},
		},
	})
}

func requestBoundaryContract(t *testing.T, handler http.Handler, path string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodGet, path, nil)
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, req)
	if response.Code != http.StatusOK {
		t.Fatalf("GET %s status = %d, want %d", path, response.Code, http.StatusOK)
	}
	return response
}

func sortedMapKeys(record map[string]any) []string {
	keys := make([]string, 0, len(record))
	for key := range record {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
