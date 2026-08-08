package engine

import (
	"context"
	"errors"
	"testing"

	"building-code-map/backend/regulatory"
	"building-code-map/backend/snapshot"
)

func TestSnapshotGeographyResolverHonorsPolygonHolesAndStableOverlaps(t *testing.T) {
	resolver := NewSnapshotGeographyResolver(snapshot.Snapshot{BoundaryFeatures: []snapshot.BoundaryFeature{
		{LayerFamily: "states", FeatureID: "08", Title: "Colorado", Attributes: map[string]any{"STATEFP": "08"}, Geometry: rectangle(0, 0, 10, 10)},
		{LayerFamily: "special_areas", FeatureID: "z", Title: "Second", Geometry: rectangle(0, 0, 10, 10)},
		{LayerFamily: "special_areas", FeatureID: "a", Title: "First", Geometry: rectangle(0, 0, 10, 10)},
		{LayerFamily: "counties", FeatureID: "08031", Title: "Denver", Geometry: snapshot.Geometry{Type: "Polygon", Coordinates: [][][]float64{
			{{0, 0}, {10, 0}, {10, 10}, {0, 10}, {0, 0}},
			{{4, 4}, {6, 4}, {6, 6}, {4, 6}, {4, 4}},
		}}},
	}}, regulatory.EmptyCatalog())

	geography, err := resolver.ResolveGeography(context.Background(), Point{Longitude: 2, Latitude: 2})
	if err != nil {
		t.Fatal(err)
	}
	if geography.StateFIPS != "08" || len(geography.SpecialAreas) != 2 || geography.SpecialAreas[0].FeatureID != "a" {
		t.Fatalf("context=%#v", geography)
	}
	if geography.County == nil {
		t.Fatal("point inside polygon should match county")
	}
	holeContext, err := resolver.ResolveGeography(context.Background(), Point{Longitude: 5, Latitude: 5})
	if err != nil {
		t.Fatal(err)
	}
	if holeContext.County != nil {
		t.Fatal("point inside polygon hole must not match county")
	}
}

func TestSnapshotGeographyResolverReturnsTypedBoundaryAmbiguity(t *testing.T) {
	resolver := NewSnapshotGeographyResolver(snapshot.Snapshot{BoundaryFeatures: []snapshot.BoundaryFeature{
		{LayerFamily: "states", FeatureID: "08", Title: "Colorado", Geometry: rectangle(0, 0, 10, 10)},
		{LayerFamily: "states", FeatureID: "12", Title: "Florida", Geometry: rectangle(0, 0, 10, 10)},
	}}, regulatory.EmptyCatalog())

	_, err := resolver.ResolveGeography(context.Background(), Point{Longitude: 5, Latitude: 5})
	var ambiguity *BoundaryAmbiguityError
	if !errors.As(err, &ambiguity) || ambiguity.LayerFamily != "states" || len(ambiguity.Matches) != 2 {
		t.Fatalf("error=%T %v", err, err)
	}
}

func rectangle(minX, minY, maxX, maxY float64) snapshot.Geometry {
	return snapshot.Geometry{Type: "Polygon", Coordinates: [][][]float64{{
		{minX, minY}, {maxX, minY}, {maxX, maxY}, {minX, maxY}, {minX, minY},
	}}}
}
