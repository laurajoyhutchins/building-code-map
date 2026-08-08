package httpapi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"building-code-map/backend/engine"
	"building-code-map/backend/internal/regulatory"
	"building-code-map/backend/internal/snapshot"
)

func TestResolveEndpointFindsGeographyAndReturnsLocalRecordRequirement(t *testing.T) {
	catalog, err := regulatory.LoadCatalog(filepath.Join("..", "..", "data", "regulatory"))
	if err != nil {
		t.Fatal(err)
	}
	handler := NewHandler(snapshot.Snapshot{BoundaryFeatures: []snapshot.BoundaryFeature{
		{LayerFamily: "states", FeatureID: "08", Title: "Colorado", SourceID: "GEOID=08", Attributes: map[string]any{"STATEFP": "08"}, Geometry: testPolygon(-109, 37, -102, 41)},
		{LayerFamily: "counties", FeatureID: "08031", Title: "Denver County", Geometry: testPolygon(-105.2, 39.5, -104.5, 40)},
		{LayerFamily: "municipalities", FeatureID: "0820000", Title: "Denver", Geometry: testPolygon(-105.1, 39.6, -104.7, 39.9)},
	}}, Options{RegulatoryCatalog: catalog})
	body, _ := json.Marshal(regulatory.ResolutionRequest{Point: &regulatory.Point{Longitude: -104.99, Latitude: 39.74}, CodeFamily: "building"})
	req := httptest.NewRequest(http.MethodPost, "/resolve", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
	}
	var result regulatory.ResolutionResult
	if err := json.Unmarshal(rec.Body.Bytes(), &result); err != nil {
		t.Fatal(err)
	}
	if result.Status != "local_record_required" {
		t.Fatalf("status=%q", result.Status)
	}
	if result.Geography.Municipality == nil || result.Geography.Municipality.Name != "Denver" {
		t.Fatalf("geography=%#v", result.Geography)
	}
}

func TestResolveEndpointRejectsUnknownFields(t *testing.T) {
	catalog, err := regulatory.LoadCatalog(filepath.Join("..", "..", "data", "regulatory"))
	if err != nil {
		t.Fatal(err)
	}
	handler := NewHandler(snapshot.Snapshot{}, Options{RegulatoryCatalog: catalog})
	req := httptest.NewRequest(http.MethodPost, "/resolve", bytes.NewBufferString(`{"wat":true}`))
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
	}
}

func TestResolveEndpointRejectsCallerAuthoredGeographicContext(t *testing.T) {
	catalog, err := regulatory.LoadCatalog(filepath.Join("..", "..", "data", "regulatory"))
	if err != nil {
		t.Fatal(err)
	}
	handler := NewHandler(snapshot.Snapshot{}, Options{RegulatoryCatalog: catalog})
	req := httptest.NewRequest(http.MethodPost, "/resolve", bytes.NewBufferString(`{
		"context": {
			"state_id": "US-FL",
			"state_fips": "12",
			"incorporated": true
		},
		"code_family": "building"
	}`))
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
	}
}

func TestResolveEndpointReturnsAllAuthorityBearingBoundaryMatchesWhenAmbiguous(t *testing.T) {
	catalog, err := regulatory.LoadCatalog(filepath.Join("..", "..", "data", "regulatory"))
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name            string
		ambiguousFamily string
		features        []snapshot.BoundaryFeature
	}{
		{
			name:            "states",
			ambiguousFamily: "states",
			features: []snapshot.BoundaryFeature{
				{LayerFamily: "states", FeatureID: "12", Title: "Florida", SourceID: "GEOID=12", Attributes: map[string]any{"STATEFP": "12"}, Geometry: testPolygon(0, 0, 10, 10)},
				{LayerFamily: "states", FeatureID: "08", Title: "Colorado", SourceID: "GEOID=08", Attributes: map[string]any{"STATEFP": "08"}, Geometry: testPolygon(0, 0, 10, 10)},
			},
		},
		{
			name:            "counties",
			ambiguousFamily: "counties",
			features: []snapshot.BoundaryFeature{
				{LayerFamily: "states", FeatureID: "08", Title: "Colorado", SourceID: "GEOID=08", Attributes: map[string]any{"STATEFP": "08"}, Geometry: testPolygon(0, 0, 10, 10)},
				{LayerFamily: "counties", FeatureID: "08059", Title: "Jefferson County", SourceID: "GEOID=08059", Geometry: testPolygon(0, 0, 10, 10)},
				{LayerFamily: "counties", FeatureID: "08031", Title: "Denver County", SourceID: "GEOID=08031", Geometry: testPolygon(0, 0, 10, 10)},
			},
		},
		{
			name:            "municipalities",
			ambiguousFamily: "municipalities",
			features: []snapshot.BoundaryFeature{
				{LayerFamily: "states", FeatureID: "08", Title: "Colorado", SourceID: "GEOID=08", Attributes: map[string]any{"STATEFP": "08"}, Geometry: testPolygon(0, 0, 10, 10)},
				{LayerFamily: "municipalities", FeatureID: "0820001", Title: "Denver Annex", SourceID: "GEOID=0820001", Geometry: testPolygon(0, 0, 10, 10)},
				{LayerFamily: "municipalities", FeatureID: "0820000", Title: "Denver", SourceID: "GEOID=0820000", Geometry: testPolygon(0, 0, 10, 10)},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			handler := NewHandler(snapshot.Snapshot{BoundaryFeatures: test.features}, Options{RegulatoryCatalog: catalog})
			body := bytes.NewBufferString(`{"point":{"longitude":5,"latitude":5},"code_family":"building"}`)
			req := httptest.NewRequest(http.MethodPost, "/resolve", body)
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, req)
			if rec.Code != http.StatusConflict {
				t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
			}
			var response struct {
				Code         string                     `json:"code"`
				LayerFamily  string                     `json:"layer_family"`
				Observations []regulatory.BoundaryMatch `json:"observations"`
			}
			if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
				t.Fatal(err)
			}
			if response.Code != "boundary_ambiguous" || response.LayerFamily != test.ambiguousFamily {
				t.Fatalf("response=%#v", response)
			}
			if len(response.Observations) != 2 {
				t.Fatalf("observations=%#v", response.Observations)
			}
			if response.Observations[0].FeatureID > response.Observations[1].FeatureID {
				t.Fatalf("observations are not deterministic: %#v", response.Observations)
			}
		})
	}
}

func TestResolveGeographicContextPreservesNonExclusiveOverlaps(t *testing.T) {
	catalog, err := regulatory.LoadCatalog(filepath.Join("..", "..", "data", "regulatory"))
	if err != nil {
		t.Fatal(err)
	}
	snap := snapshot.Snapshot{BoundaryFeatures: []snapshot.BoundaryFeature{
		{LayerFamily: "states", FeatureID: "08", Title: "Colorado", SourceID: "GEOID=08", Attributes: map[string]any{"STATEFP": "08"}, Geometry: testPolygon(0, 0, 10, 10)},
		{LayerFamily: "counties", FeatureID: "08031", Title: "Denver County", SourceID: "GEOID=08031", Geometry: testPolygon(0, 0, 10, 10)},
		{LayerFamily: "municipalities", FeatureID: "0820000", Title: "Denver", SourceID: "GEOID=0820000", Geometry: testPolygon(0, 0, 10, 10)},
		{LayerFamily: "special_areas", FeatureID: "z", Title: "Second special area", SourceID: "special=z", Geometry: testPolygon(0, 0, 10, 10)},
		{LayerFamily: "special_areas", FeatureID: "a", Title: "First special area", SourceID: "special=a", Geometry: testPolygon(0, 0, 10, 10)},
		{LayerFamily: "tribal_areas", FeatureID: "b", Title: "Second tribal observation", SourceID: "tribal=b", Geometry: testPolygon(0, 0, 10, 10)},
		{LayerFamily: "tribal_areas", FeatureID: "a", Title: "First tribal observation", SourceID: "tribal=a", Geometry: testPolygon(0, 0, 10, 10)},
		{LayerFamily: "neris_jurisdictions", FeatureID: "FD2", Title: "Second fire observation", SourceID: "neris=FD2", Geometry: testPolygon(0, 0, 10, 10)},
		{LayerFamily: "neris_jurisdictions", FeatureID: "FD1", Title: "First fire observation", SourceID: "neris=FD1", Geometry: testPolygon(0, 0, 10, 10)},
	}}

	resolver := engine.NewSnapshotGeographyResolver(snap, catalog)
	context, err := resolver.ResolveGeography(t.Context(), engine.Point{Longitude: 5, Latitude: 5})
	if err != nil {
		t.Fatal(err)
	}
	if len(context.SpecialAreas) != 2 || len(context.TribalAreas) != 2 || len(context.FireJurisdictions) != 2 {
		t.Fatalf("context did not preserve overlaps: %#v", context)
	}
	if context.SpecialAreas[0].FeatureID != "a" || context.TribalAreas[0].FeatureID != "a" || context.FireJurisdictions[0].FeatureID != "FD1" {
		t.Fatalf("overlap ordering is not deterministic: %#v", context)
	}
}

func TestGeometryContainsPointHonorsPolygonHoles(t *testing.T) {
	geometry := snapshot.Geometry{Type: "Polygon", Coordinates: [][][]float64{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}}, {{4, 4}, {6, 4}, {6, 6}, {4, 6}, {4, 4}}}}
	resolver := engine.NewSnapshotGeographyResolver(snapshot.Snapshot{BoundaryFeatures: []snapshot.BoundaryFeature{{
		LayerFamily: "states", FeatureID: "08", Geometry: geometry,
	}}}, regulatory.EmptyCatalog())
	if _, err := resolver.ResolveGeography(t.Context(), engine.Point{Longitude: 2, Latitude: 2}); err != nil {
		t.Fatalf("outer point should match: %v", err)
	}
	if _, err := resolver.ResolveGeography(t.Context(), engine.Point{Longitude: 5, Latitude: 5}); err == nil {
		t.Fatal("hole point should not match")
	}
}

func testPolygon(minX, minY, maxX, maxY float64) snapshot.Geometry {
	return snapshot.Geometry{Type: "Polygon", Coordinates: [][][]float64{{
		{minX, minY},
		{maxX, minY},
		{maxX, maxY},
		{minX, maxY},
		{minX, minY},
	}}}
}
