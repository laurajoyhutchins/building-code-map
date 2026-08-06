package engine

import (
	"context"
	"errors"
	"path/filepath"
	"testing"
	"time"

	"building-code-map/backend/internal/geocoder"
	"building-code-map/backend/internal/regulatory"
	"building-code-map/backend/internal/snapshot"
)

func TestEngineResolveIsDeterministicWithFixedClock(t *testing.T) {
	catalog, err := regulatory.LoadCatalog(filepath.Join("..", "data", "regulatory"))
	if err != nil {
		t.Fatal(err)
	}
	clock := NewFixedClock(time.Date(2026, 8, 6, 12, 34, 56, 0, time.UTC))
	engine, err := New(Config{
		Snapshot: snapshot.Snapshot{BoundaryFeatures: []snapshot.BoundaryFeature{
			{LayerFamily: "states", FeatureID: "08", Title: "Colorado", Attributes: map[string]any{"STATEFP": "08"}, Geometry: rectangle(-109, 37, -102, 41)},
		}},
		RegulatoryCatalog: catalog,
		Clock:             clock,
		BundleIdentity: BundleIdentity{
			SourceCommit:            "0123456789abcdef0123456789abcdef01234567",
			EngineVersion:           "0.1.0",
			BundleManifestDigest:    "sha256:bundle",
			BoundarySnapshotDigest:  "sha256:boundary",
			RegulatoryCatalogDigest: "sha256:regulatory",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	query := Query{Point: &Point{Longitude: -104.99, Latitude: 39.74}, ApplicabilityDate: "2026-08-06", CodeFamily: "building"}
	first, err := engine.Resolve(context.Background(), query)
	if err != nil {
		t.Fatal(err)
	}
	second, err := engine.Resolve(context.Background(), query)
	if err != nil {
		t.Fatal(err)
	}
	if first.Resolution.GeneratedAt != second.Resolution.GeneratedAt || first.Provenance != second.Provenance {
		t.Fatalf("non-deterministic identity: first=%#v second=%#v", first, second)
	}
}

func TestEngineRejectsMissingRegulatoryCatalog(t *testing.T) {
	engine, err := New(Config{Snapshot: snapshot.Snapshot{BoundaryFeatures: []snapshot.BoundaryFeature{{LayerFamily: "states", FeatureID: "08", Geometry: rectangle(-109, 37, -102, 41)}}}})
	if err != nil {
		t.Fatal(err)
	}
	_, err = engine.Resolve(context.Background(), Query{Point: &Point{Longitude: -104, Latitude: 39}, ApplicabilityDate: "2026-08-06"})
	var engineErr EngineError
	if !errors.As(err, &engineErr) || engineErr.Code != ErrorRegulatoryCatalogUnavailable {
		t.Fatalf("error=%T %v", err, err)
	}
}

func TestEngineGeocodePreservesRankingAndInterpolationProvenance(t *testing.T) {
	service := fakeGeocoder{result: geocoder.Result{
		Query:  "10 Main St, Denver, CO",
		Status: geocoder.StatusMatched,
		Selected: &geocoder.Candidate{
			MatchedAddress:       "10 Main St, Denver, CO",
			Longitude:            -104.99,
			Latitude:             39.74,
			ScoreFactors:         map[string]float64{"custom_factor": 0.75},
			ScoreKind:            "quality_score",
			RankingPolicyVersion: "ranking-v3",
			Interpolation: &geocoder.InterpolationProvenance{
				SourceRangeID: "range-7", RequestedHouseNumber: 10, Fraction: 0.25,
				CoordinateReferenceSystem: "EPSG:4326", MethodVersion: "interp-v2",
			},
		},
	}}
	engine, err := New(Config{Geocoder: service})
	if err != nil {
		t.Fatal(err)
	}
	result, err := engine.Geocode(context.Background(), "10 Main St, Denver, CO")
	if err != nil {
		t.Fatal(err)
	}
	if result.Selected == nil || result.Selected.RankingPolicyVersion != "ranking-v3" || result.Selected.ScoreFactors["custom_factor"] != 0.75 {
		t.Fatalf("candidate provenance=%#v", result.Selected)
	}
	if result.Selected.Interpolation == nil || result.Selected.Interpolation.SourceRangeID != "range-7" || result.Selected.Interpolation.MethodVersion != "interp-v2" {
		t.Fatalf("interpolation provenance=%#v", result.Selected.Interpolation)
	}
}

type fakeGeocoder struct{ result geocoder.Result }

func (fake fakeGeocoder) Geocode(context.Context, geocoder.Query) (geocoder.Result, error) {
	return fake.result, nil
}
