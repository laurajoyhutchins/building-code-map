package httpapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"

	"building-code-map/backend/internal/geocoder"
	"building-code-map/backend/internal/regulatory"
	"building-code-map/backend/internal/snapshot"
)

type readinessGeocoder struct{}

func (readinessGeocoder) Geocode(context.Context, geocoder.Query) (geocoder.Result, error) {
	return geocoder.Result{}, nil
}

func TestReadinessReportsFullDegradedAndUnavailableCapabilities(t *testing.T) {
	catalog, err := regulatory.LoadCatalog(filepath.Join("..", "..", "data", "regulatory"))
	if err != nil {
		t.Fatal(err)
	}
	boundarySnapshot := snapshot.Snapshot{
		LayerFamilies: []snapshot.LayerFamily{{Key: "states"}},
		BoundaryFeatures: []snapshot.BoundaryFeature{{
			LayerFamily: "states",
			FeatureID:   "08",
		}},
	}

	tests := []struct {
		name            string
		handler         http.Handler
		wantStatus      int
		wantReadiness   string
		wantAddressState string
	}{
		{
			name: "ready",
			handler: NewHandler(boundarySnapshot, Options{
				RegulatoryCatalog: catalog,
				Geocoder:           readinessGeocoder{},
			}),
			wantStatus:       http.StatusOK,
			wantReadiness:    "ready",
			wantAddressState: "available",
		},
		{
			name:             "degraded",
			handler:          NewHandler(boundarySnapshot),
			wantStatus:       http.StatusOK,
			wantReadiness:    "degraded",
			wantAddressState: "unavailable",
		},
		{
			name: "not ready",
			handler: NewHandler(snapshot.Snapshot{}, Options{
				RegulatoryCatalog: catalog,
				Geocoder:           readinessGeocoder{},
			}),
			wantStatus:       http.StatusServiceUnavailable,
			wantReadiness:    "not_ready",
			wantAddressState: "unavailable",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/ready", nil)
			rec := httptest.NewRecorder()
			test.handler.ServeHTTP(rec, req)
			if rec.Code != test.wantStatus {
				t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
			}
			var response struct {
				Readiness    string `json:"readiness"`
				Capabilities map[string]struct {
					Status string `json:"status"`
				} `json:"capabilities"`
			}
			if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
				t.Fatal(err)
			}
			if response.Readiness != test.wantReadiness {
				t.Fatalf("readiness=%q want=%q", response.Readiness, test.wantReadiness)
			}
			if response.Capabilities["address_lookup"].Status != test.wantAddressState {
				t.Fatalf("address_lookup=%#v", response.Capabilities["address_lookup"])
			}
		})
	}
}

func TestCORSPreflightAllowsConfiguredOrigin(t *testing.T) {
	handler := NewHandler(snapshot.Snapshot{}, Options{AllowedOrigins: []string{"https://example.com"}})
	req := httptest.NewRequest(http.MethodOptions, "/lookup", nil)
	req.Header.Set("Origin", "https://example.com")
	req.Header.Set("Access-Control-Request-Method", http.MethodPost)
	req.Header.Set("Access-Control-Request-Headers", "content-type")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
	}
	if got := rec.Header().Get("Access-Control-Allow-Origin"); got != "https://example.com" {
		t.Fatalf("allow origin=%q", got)
	}
	if got := rec.Header().Get("Access-Control-Allow-Methods"); !strings.Contains(got, http.MethodPost) {
		t.Fatalf("allow methods=%q", got)
	}
	if got := rec.Header().Get("Access-Control-Allow-Headers"); !strings.Contains(strings.ToLower(got), "content-type") {
		t.Fatalf("allow headers=%q", got)
	}
}

func TestCORSPreflightRejectsUnconfiguredOrigin(t *testing.T) {
	handler := NewHandler(snapshot.Snapshot{}, Options{AllowedOrigins: []string{"https://example.com"}})
	req := httptest.NewRequest(http.MethodOptions, "/lookup", nil)
	req.Header.Set("Origin", "https://untrusted.example")
	req.Header.Set("Access-Control-Request-Method", http.MethodPost)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusForbidden {
		t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
	}
	if got := rec.Header().Get("Access-Control-Allow-Origin"); got != "" {
		t.Fatalf("unexpected allow origin=%q", got)
	}
}
