package httpapi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"building-code-map/backend/internal/regulatory"
	"building-code-map/backend/internal/snapshot"
)

func TestResolveEndpointFindsGeographyAndReturnsLocalRecordRequirement(t *testing.T) {
	catalog, err := regulatory.LoadCatalog(filepath.Join("..", "..", "data", "regulatory"))
	if err != nil {
		t.Fatal(err)
	}
	handler := NewHandler(snapshot.Snapshot{BoundaryFeatures: []snapshot.BoundaryFeature{
		{LayerFamily: "states", FeatureID: "08", Title: "Colorado", SourceID: "GEOID=08", Attributes: map[string]any{"STATEFP": "08"}, Geometry: snapshot.Geometry{Type: "Polygon", Coordinates: [][][]float64{{{-109, 37}, {-102, 37}, {-102, 41}, {-109, 41}, {-109, 37}}}}},
		{LayerFamily: "counties", FeatureID: "08031", Title: "Denver County", Geometry: snapshot.Geometry{Type: "Polygon", Coordinates: [][][]float64{{{-105.2, 39.5}, {-104.5, 39.5}, {-104.5, 40}, {-105.2, 40}, {-105.2, 39.5}}}}},
		{LayerFamily: "municipalities", FeatureID: "0820000", Title: "Denver", Geometry: snapshot.Geometry{Type: "Polygon", Coordinates: [][][]float64{{{-105.1, 39.6}, {-104.7, 39.6}, {-104.7, 39.9}, {-105.1, 39.9}, {-105.1, 39.6}}}}},
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

func TestGeometryContainsPointHonorsPolygonHoles(t *testing.T) {
	geometry := snapshot.Geometry{Type: "Polygon", Coordinates: [][][]float64{{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}}, {{4, 4}, {6, 4}, {6, 6}, {4, 6}, {4, 4}}}}
	if !geometryContainsPoint(geometry, 2, 2) {
		t.Fatal("outer point should match")
	}
	if geometryContainsPoint(geometry, 5, 5) {
		t.Fatal("hole point should not match")
	}
}
