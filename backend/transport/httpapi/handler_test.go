package httpapi_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"building-code-map/backend/snapshot"
	"building-code-map/backend/transport/httpapi"
)

func TestHandlerServesGoBackendContractRoutes(t *testing.T) {
	handler := httpapi.NewLegacyHandler(snapshot.Snapshot{
		LayerFamilies: []snapshot.LayerFamily{
			{
				Key:            "states",
				Label:          "States",
				MartinLayerID:  "tigerweb.states",
				Description:    "State boundaries mirrored from TIGERweb.",
				DefaultEnabled: true,
			},
			{
				Key:            "counties",
				Label:          "Counties",
				MartinLayerID:  "tigerweb.counties",
				Description:    "County boundaries used for jurisdiction lookups.",
				DefaultEnabled: true,
			},
			{
				Key:            "neris_jurisdictions",
				Label:          "NERIS jurisdictions",
				MartinLayerID:  "neris.department_jurisdictions",
				Description:    "Real NERIS department jurisdiction polygons joined to department attributes.",
				DefaultEnabled: false,
			},
		},
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
					"STATEFP": "08",
					"NAME":    "Colorado",
					"LSAD":    nil,
				},
			},
			{
				LayerFamily:    "counties",
				FeatureID:      "08031",
				Title:          "Denver County",
				Subtitle:       "County boundary",
				SourceID:       "COUNTYFP=031",
				GeometryLabel:  "Polygon",
				GeometrySource: "tigerweb_live",
				LastSyncedAt:   time.Date(2026, 6, 22, 12, 0, 0, 0, time.UTC),
				Geometry: snapshot.Geometry{
					Type: "Polygon",
					Coordinates: []any{
						[]any{
							[]float64{-105.13, 39.63},
							[]float64{-104.55, 39.63},
							[]float64{-104.55, 39.92},
							[]float64{-105.13, 39.92},
							[]float64{-105.13, 39.63},
						},
					},
				},
				Attributes: map[string]any{
					"STATEFP":  "08",
					"COUNTYFP": "031",
					"NAME":     "Denver",
					"NAMELSAD": "Denver County",
				},
			},
			{
				LayerFamily:    "neris_jurisdictions",
				FeatureID:      "FD01001122",
				Title:          "Independence Volunteer Fire Department",
				Subtitle:       "Department jurisdiction",
				SourceID:       "neris_id=FD01001122",
				GeometryLabel:  "Polygon",
				GeometrySource: "neris_department_jurisdiction",
				LastSyncedAt:   time.Date(2026, 6, 22, 12, 0, 0, 0, time.UTC),
				Geometry: snapshot.Geometry{
					Type: "Polygon",
					Coordinates: []any{
						[]any{
							[]float64{-86.6949, 32.5746},
							[]float64{-86.6836, 32.5747},
							[]float64{-86.6832, 32.5746},
							[]float64{-86.6949, 32.5746},
						},
					},
				},
				Attributes: map[string]any{
					"neris_id":        "FD01001122",
					"name":            "Independence Volunteer Fire Department",
					"internal_id":     nil,
					"entity_type":     nil,
					"department_type": "VOLUNTEER",
					"region_type":     "JURISDICTION",
				},
			},
		},
		RefreshStatus: snapshot.RefreshStatus{
			Status:                  "ok",
			LatestSuccessfulRefresh: time.Date(2026, 6, 22, 9, 40, 0, 0, time.UTC),
			LatestAttempt:           time.Date(2026, 6, 22, 9, 40, 0, 0, time.UTC),
			NextScheduledRefresh:    time.Date(2026, 6, 23, 9, 40, 0, 0, time.UTC),
			Message:                 "Cached TIGERweb snapshot is current.",
		},
	})

	assertJSONResponse(t, handler, http.MethodGet, "/health", map[string]any{"status": "ok"})
	assertJSONResponse(t, handler, http.MethodGet, "/ready", map[string]any{"status": "ok"})

	assertJSONArrayResponse(t, handler, http.MethodGet, "/layers", []map[string]any{
		{
			"key":             "states",
			"label":           "States",
			"martin_layer_id": "tigerweb.states",
			"description":     "State boundaries mirrored from TIGERweb.",
			"default_enabled": true,
		},
		{
			"key":             "counties",
			"label":           "Counties",
			"martin_layer_id": "tigerweb.counties",
			"description":     "County boundaries used for jurisdiction lookups.",
			"default_enabled": true,
		},
		{
			"key":             "neris_jurisdictions",
			"label":           "NERIS jurisdictions",
			"martin_layer_id": "neris.department_jurisdictions",
			"description":     "Real NERIS department jurisdiction polygons joined to department attributes.",
			"default_enabled": false,
		},
	})

	assertJSONArrayResponse(t, handler, http.MethodGet, "/boundaries", []map[string]any{
		{
			"layer_family":    "states",
			"feature_id":      "08",
			"title":           "Colorado",
			"subtitle":        "State boundary",
			"source_id":       "GEOID=08",
			"geometry_label":  "Polygon",
			"geometry_source": "tigerweb_live",
		},
		{
			"layer_family":    "counties",
			"feature_id":      "08031",
			"title":           "Denver County",
			"subtitle":        "County boundary",
			"source_id":       "COUNTYFP=031",
			"geometry_label":  "Polygon",
			"geometry_source": "tigerweb_live",
		},
		{
			"layer_family":    "neris_jurisdictions",
			"feature_id":      "FD01001122",
			"title":           "Independence Volunteer Fire Department",
			"subtitle":        "Department jurisdiction",
			"source_id":       "neris_id=FD01001122",
			"geometry_label":  "Polygon",
			"geometry_source": "neris_department_jurisdiction",
		},
	})

	assertJSONResponse(t, handler, http.MethodGet, "/features/states/08", map[string]any{
		"layer_family":    "states",
		"feature_id":      "08",
		"title":           "Colorado",
		"subtitle":        "State boundary",
		"source_id":       "GEOID=08",
		"geometry_label":  "Polygon",
		"geometry_source": "tigerweb_live",
		"attributes": map[string]any{
			"STATEFP": "08",
			"NAME":    "Colorado",
			"LSAD":    nil,
		},
	})

	assertJSONResponse(t, handler, http.MethodGet, "/features/neris_jurisdictions/FD01001122", map[string]any{
		"layer_family":    "neris_jurisdictions",
		"feature_id":      "FD01001122",
		"title":           "Independence Volunteer Fire Department",
		"geometry_source": "neris_department_jurisdiction",
	})

	assertJSONResponse(t, handler, http.MethodGet, "/refresh/status", map[string]any{
		"status":                    "ok",
		"message":                   "Cached TIGERweb snapshot is current.",
		"latest_successful_refresh": "2026-06-22T09:40:00Z",
		"latest_attempt":            "2026-06-22T09:40:00Z",
		"next_scheduled_refresh":    "2026-06-23T09:40:00Z",
	})

	assertJSONResponse(t, handler, http.MethodPost, "/refresh/trigger", map[string]any{
		"status":  "disabled",
		"message": "Live refresh is disabled for the cached snapshot.",
	})
}

func TestHandlerAddsAllowedOriginCORSHeader(t *testing.T) {
	handler := httpapi.NewLegacyHandler(snapshot.Snapshot{}, httpapi.Options{
		AllowedOrigins: []string{"http://127.0.0.1:5173"},
	})

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	req.Header.Set("Origin", "http://127.0.0.1:5173")
	resp := httptest.NewRecorder()

	handler.ServeHTTP(resp, req)

	if got := resp.Header().Get("Access-Control-Allow-Origin"); got != "http://127.0.0.1:5173" {
		t.Fatalf("unexpected CORS header: got %q", got)
	}
}

func TestParseAllowedOriginsTrimsAndFiltersOrigins(t *testing.T) {
	got := httpapi.ParseAllowedOrigins(" http://localhost:3000/ , https://example.com , ftp://bad ")
	want := []string{"http://localhost:3000", "https://example.com"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ParseAllowedOrigins() = %#v, want %#v", got, want)
	}
}

func TestParseAllowedOriginsFallsBackToLoopbackDefaults(t *testing.T) {
	got := httpapi.ParseAllowedOrigins("ftp://bad, https://example.com/path")
	want := []string{
		"http://localhost:5173",
		"http://127.0.0.1:5173",
		"http://[::1]:5173",
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ParseAllowedOrigins() = %#v, want %#v", got, want)
	}
}

func assertJSONResponse(t *testing.T, handler http.Handler, method, path string, want map[string]any) {
	t.Helper()

	req := httptest.NewRequest(method, path, nil)
	resp := httptest.NewRecorder()

	handler.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("%s %s status = %d, want %d", method, path, resp.Code, http.StatusOK)
	}

	assertJSONContainsMap(t, resp.Body.Bytes(), want)
}

func assertJSONArrayResponse(t *testing.T, handler http.Handler, method, path string, want []map[string]any) {
	t.Helper()

	req := httptest.NewRequest(method, path, nil)
	resp := httptest.NewRecorder()

	handler.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("%s %s status = %d, want %d", method, path, resp.Code, http.StatusOK)
	}

	var got []map[string]any
	if err := json.Unmarshal(resp.Body.Bytes(), &got); err != nil {
		t.Fatalf("unmarshal response: %v", err)
	}

	if len(got) != len(want) {
		t.Fatalf("array length = %d, want %d", len(got), len(want))
	}

	for i := range want {
		assertContainsMap(t, got[i], want[i])
	}
}

func assertJSONContainsMap(t *testing.T, raw []byte, want map[string]any) {
	t.Helper()

	var got map[string]any
	if err := json.Unmarshal(raw, &got); err != nil {
		t.Fatalf("unmarshal response: %v", err)
	}

	assertContainsMap(t, got, want)
}

func assertContainsMap(t *testing.T, got map[string]any, want map[string]any) {
	t.Helper()

	for key, wantValue := range want {
		gotValue, ok := got[key]
		if !ok {
			t.Fatalf("missing key %q in %#v", key, got)
		}
		if !jsonValuesEqual(gotValue, wantValue) {
			t.Fatalf("key %q = %#v, want %#v", key, gotValue, wantValue)
		}
	}
}

func jsonValuesEqual(got any, want any) bool {
	gotJSON, err := json.Marshal(got)
	if err != nil {
		return false
	}

	wantJSON, err := json.Marshal(want)
	if err != nil {
		return false
	}

	return string(gotJSON) == string(wantJSON)
}
