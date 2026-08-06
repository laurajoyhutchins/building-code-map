package httpapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"building-code-map/backend/geocoder"
	"building-code-map/backend/regulatory"
	"building-code-map/backend/snapshot"
)

type fakeGeocoder struct {
	result geocoder.Result
	err    error
}

func (fake fakeGeocoder) Geocode(context.Context, geocoder.Query) (geocoder.Result, error) {
	return fake.result, fake.err
}

func TestGeocodeEndpointReturnsUnavailableWithoutSnapshot(t *testing.T) {
	handler := NewLegacyHandler(snapshot.Snapshot{})
	req := httptest.NewRequest(http.MethodPost, "/geocode", bytes.NewBufferString(`{"address":"1600 N Broadway, Denver, CO"}`))
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusServiceUnavailable {
		t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
	}
}

func TestGeocodeEndpointReturnsExactResult(t *testing.T) {
	candidate := geocoder.Candidate{
		MatchedAddress: "1600 N Broadway St Denver CO 80202",
		Longitude:      -104.99, Latitude: 39.74,
		Precision:  geocoder.PrecisionAddressPoint,
		Confidence: 1,
		Source:     "fixture", SourceRecordID: "one", SourceVintage: "2026-08-01",
	}
	handler := NewLegacyHandler(snapshot.Snapshot{}, Options{Geocoder: fakeGeocoder{result: geocoder.Result{
		Query: "1600 N Broadway, Denver, CO", Normalized: "1600 N BROADWAY, DENVER, CO",
		Status: geocoder.StatusMatched, Selected: &candidate, Candidates: []geocoder.Candidate{candidate}, Warnings: []string{},
	}}})
	req := httptest.NewRequest(http.MethodPost, "/geocode", bytes.NewBufferString(`{"address":"1600 N Broadway, Denver, CO"}`))
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
	}
}

func TestGeocodeEndpointReturnsConflictForAmbiguity(t *testing.T) {
	handler := NewLegacyHandler(snapshot.Snapshot{}, Options{Geocoder: fakeGeocoder{result: geocoder.Result{
		Query: "100 Main St, Springfield, NJ", Status: geocoder.StatusAmbiguous,
		Candidates: []geocoder.Candidate{{SourceRecordID: "one"}, {SourceRecordID: "two"}}, Warnings: []string{},
	}}})
	req := httptest.NewRequest(http.MethodPost, "/geocode", bytes.NewBufferString(`{"address":"100 Main St, Springfield, NJ"}`))
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusConflict {
		t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
	}
}

func TestGeocodeEndpointRejectsUnknownFields(t *testing.T) {
	handler := NewLegacyHandler(snapshot.Snapshot{}, Options{Geocoder: fakeGeocoder{err: errors.New("not reached")}})
	req := httptest.NewRequest(http.MethodPost, "/geocode", bytes.NewBufferString(`{"address":"1600 N Broadway, Denver, CO","wat":true}`))
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
	}
}

func TestLookupEndpointComposesGeocoderAndResolver(t *testing.T) {
	catalog, err := regulatory.LoadCatalog(filepath.Join("..", "..", "data", "regulatory"))
	if err != nil {
		t.Fatal(err)
	}
	candidate := geocoder.Candidate{
		MatchedAddress: "1600 N Broadway St Denver CO 80202",
		Longitude:      -104.99, Latitude: 39.74,
		Precision:  geocoder.PrecisionAddressPoint,
		Confidence: 1,
		Source:     "fixture", SourceRecordID: "one", SourceVintage: "2026-08-01",
	}
	handler := NewLegacyHandler(snapshot.Snapshot{BoundaryFeatures: []snapshot.BoundaryFeature{
		{LayerFamily: "states", FeatureID: "08", Title: "Colorado", SourceID: "GEOID=08", Attributes: map[string]any{"STATEFP": "08"}, Geometry: snapshot.Geometry{Type: "Polygon", Coordinates: [][][]float64{{{-109, 37}, {-102, 37}, {-102, 41}, {-109, 41}, {-109, 37}}}}},
		{LayerFamily: "counties", FeatureID: "08031", Title: "Denver County", Geometry: snapshot.Geometry{Type: "Polygon", Coordinates: [][][]float64{{{-105.2, 39.5}, {-104.5, 39.5}, {-104.5, 40}, {-105.2, 40}, {-105.2, 39.5}}}}},
		{LayerFamily: "municipalities", FeatureID: "0820000", Title: "Denver", Geometry: snapshot.Geometry{Type: "Polygon", Coordinates: [][][]float64{{{-105.1, 39.6}, {-104.7, 39.6}, {-104.7, 39.9}, {-105.1, 39.9}, {-105.1, 39.6}}}}},
	}}, Options{
		RegulatoryCatalog: catalog,
		Geocoder: fakeGeocoder{result: geocoder.Result{
			Query: "1600 N Broadway, Denver, CO", Status: geocoder.StatusMatched,
			Selected: &candidate, Candidates: []geocoder.Candidate{candidate}, Warnings: []string{},
		}},
	})
	body, _ := json.Marshal(map[string]any{
		"address":     "1600 N Broadway, Denver, CO",
		"code_family": "building",
	})
	req := httptest.NewRequest(http.MethodPost, "/lookup", bytes.NewReader(body))
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("status=%d body=%s", rec.Code, rec.Body.String())
	}
	var response struct {
		Geocode    geocoder.Result             `json:"geocode"`
		Resolution regulatory.ResolutionResult `json:"resolution"`
	}
	if err := json.Unmarshal(rec.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}
	if response.Resolution.Geography.Municipality == nil || response.Resolution.Geography.Municipality.Name != "Denver" {
		t.Fatalf("resolution=%#v", response.Resolution)
	}
}
