package httpapi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"

	"building-code-map/backend/regulatory"
	"building-code-map/backend/snapshot"
)

func TestResolveEndpointDisclosesAssumedApplicabilityDate(t *testing.T) {
	catalog, err := regulatory.LoadCatalog(filepath.Join("..", "..", "data", "regulatory"))
	if err != nil {
		t.Fatal(err)
	}
	handler := NewLegacyHandler(snapshot.Snapshot{BoundaryFeatures: []snapshot.BoundaryFeature{{
		LayerFamily: "states",
		FeatureID:   "08",
		Title:       "Colorado",
		SourceID:    "GEOID=08",
		Attributes:  map[string]any{"STATEFP": "08"},
		Geometry:    testPolygon(-109, 37, -102, 41),
	}}}, Options{RegulatoryCatalog: catalog})

	for _, test := range []struct {
		name        string
		body        string
		wantWarning bool
	}{
		{
			name:        "omitted",
			body:        `{"point":{"longitude":-104.99,"latitude":39.74},"code_family":"building"}`,
			wantWarning: true,
		},
		{
			name:        "explicit",
			body:        `{"point":{"longitude":-104.99,"latitude":39.74},"code_family":"building","applicability_date":"2026-08-03"}`,
			wantWarning: false,
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/resolve", bytes.NewBufferString(test.body))
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, req)
			if rec.Code != http.StatusOK {
				t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
			}
			var result regulatory.ResolutionResult
			if err := json.Unmarshal(rec.Body.Bytes(), &result); err != nil {
				t.Fatal(err)
			}
			hasWarning := strings.Contains(strings.Join(result.Warnings, " "), "Applicability date was omitted")
			if hasWarning != test.wantWarning {
				t.Fatalf("warnings=%#v want assumed-date warning=%t", result.Warnings, test.wantWarning)
			}
		})
	}
}
