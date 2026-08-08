package engine

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"reflect"
	"sort"
	"strings"

	"building-code-map/backend/regulatory"
	"building-code-map/backend/snapshot"
)

// GeographyResolver turns a point into the boundary observations used by
// regulatory resolution. It is deliberately independent of HTTP and address
// presentation concerns.
type GeographyResolver interface {
	ResolveGeography(context.Context, Point) (regulatory.GeographicContext, error)
}

type SnapshotGeographyResolver struct {
	snapshot snapshot.Snapshot
	catalog  regulatory.Catalog
}

func NewSnapshotGeographyResolver(snap snapshot.Snapshot, catalog regulatory.Catalog) GeographyResolver {
	return SnapshotGeographyResolver{snapshot: snap, catalog: catalog}
}

type BoundaryAmbiguityError struct {
	LayerFamily string
	Matches     []regulatory.BoundaryMatch
}

func (err *BoundaryAmbiguityError) Error() string {
	return fmt.Sprintf("point matched multiple %s boundary observations; confirm the controlling boundary locally", err.LayerFamily)
}

func (resolver SnapshotGeographyResolver) ResolveGeography(_ context.Context, point Point) (regulatory.GeographicContext, error) {
	if math.IsNaN(point.Longitude) || math.IsInf(point.Longitude, 0) || math.IsNaN(point.Latitude) || math.IsInf(point.Latitude, 0) ||
		point.Longitude < -180 || point.Longitude > 180 || point.Latitude < -90 || point.Latitude > 90 {
		return regulatory.GeographicContext{}, errors.New("point coordinates are outside valid longitude/latitude ranges")
	}

	matchedFeatures := map[string][]snapshot.BoundaryFeature{}
	for _, feature := range resolver.snapshot.BoundaryFeatures {
		if geometryContainsPoint(feature.Geometry, point.Longitude, point.Latitude) {
			matchedFeatures[feature.LayerFamily] = append(matchedFeatures[feature.LayerFamily], feature)
		}
	}
	for layerFamily := range matchedFeatures {
		sort.Slice(matchedFeatures[layerFamily], func(i, j int) bool {
			left := matchedFeatures[layerFamily][i]
			right := matchedFeatures[layerFamily][j]
			if left.FeatureID != right.FeatureID {
				return left.FeatureID < right.FeatureID
			}
			if left.Title != right.Title {
				return left.Title < right.Title
			}
			return left.SourceID < right.SourceID
		})
	}

	for _, layerFamily := range []string{"states", "counties", "municipalities"} {
		if len(matchedFeatures[layerFamily]) > 1 {
			return regulatory.GeographicContext{}, &BoundaryAmbiguityError{
				LayerFamily: layerFamily,
				Matches:     boundaryMatches(matchedFeatures[layerFamily]),
			}
		}
	}

	stateFeatures := matchedFeatures["states"]
	if len(stateFeatures) == 0 {
		return regulatory.GeographicContext{}, errors.New("point did not match a supported state boundary")
	}
	stateFeature := stateFeatures[0]
	stateFIPS := firstString(stateFeature.Attributes, "STATEFP", "GEOID")
	if stateFIPS == "" && len(stateFeature.FeatureID) == 2 {
		stateFIPS = stateFeature.FeatureID
	}
	result := regulatory.GeographicContext{
		StateFIPS: stateFIPS,
		StateName: stateFeature.Title,
	}
	if profile, ok := resolver.catalog.Profile("", stateFIPS); ok {
		result.StateID = profile.StateID
	}
	if counties := matchedFeatures["counties"]; len(counties) == 1 {
		match := boundaryMatch(counties[0])
		result.County = &match
	}
	if municipalities := matchedFeatures["municipalities"]; len(municipalities) == 1 {
		match := boundaryMatch(municipalities[0])
		result.Municipality = &match
		result.Incorporated = true
	}
	result.SpecialAreas = boundaryMatches(matchedFeatures["special_areas"])
	result.TribalAreas = boundaryMatches(matchedFeatures["tribal_areas"])
	result.FireJurisdictions = boundaryMatches(matchedFeatures["neris_jurisdictions"])
	return result, nil
}

func boundaryMatches(features []snapshot.BoundaryFeature) []regulatory.BoundaryMatch {
	matches := make([]regulatory.BoundaryMatch, 0, len(features))
	for _, feature := range features {
		matches = append(matches, boundaryMatch(feature))
	}
	return matches
}

func boundaryMatch(feature snapshot.BoundaryFeature) regulatory.BoundaryMatch {
	return regulatory.BoundaryMatch{
		LayerFamily: feature.LayerFamily,
		FeatureID:   feature.FeatureID,
		Name:        feature.Title,
		SourceID:    feature.SourceID,
	}
}

func firstString(values map[string]any, keys ...string) string {
	for _, key := range keys {
		if value, ok := values[key]; ok {
			if text, ok := value.(string); ok && strings.TrimSpace(text) != "" {
				return strings.TrimSpace(text)
			}
		}
	}
	return ""
}

func geometryContainsPoint(geometry snapshot.Geometry, longitude, latitude float64) bool {
	switch geometry.Type {
	case "Polygon":
		polygon, ok := coordinateSlice(geometry.Coordinates)
		return ok && polygonContainsPoint(polygon, longitude, latitude)
	case "MultiPolygon":
		polygons, ok := coordinateSlice(geometry.Coordinates)
		if !ok {
			return false
		}
		for _, polygonValue := range polygons {
			polygon, ok := coordinateSlice(polygonValue)
			if ok && polygonContainsPoint(polygon, longitude, latitude) {
				return true
			}
		}
	}
	return false
}

func polygonContainsPoint(polygon []any, longitude, latitude float64) bool {
	if len(polygon) == 0 {
		return false
	}
	outer, ok := coordinateSlice(polygon[0])
	if !ok || !ringContainsPoint(outer, longitude, latitude) {
		return false
	}
	for _, holeValue := range polygon[1:] {
		hole, ok := coordinateSlice(holeValue)
		if ok && ringContainsPoint(hole, longitude, latitude) {
			return false
		}
	}
	return true
}

func ringContainsPoint(ring []any, longitude, latitude float64) bool {
	if len(ring) < 3 {
		return false
	}
	inside := false
	j := len(ring) - 1
	for i := 0; i < len(ring); i++ {
		xi, yi, okI := position(ring[i])
		xj, yj, okJ := position(ring[j])
		if !okI || !okJ {
			return false
		}
		if pointOnSegment(longitude, latitude, xi, yi, xj, yj) {
			return true
		}
		intersects := (yi > latitude) != (yj > latitude) && longitude < (xj-xi)*(latitude-yi)/(yj-yi)+xi
		if intersects {
			inside = !inside
		}
		j = i
	}
	return inside
}

func pointOnSegment(px, py, x1, y1, x2, y2 float64) bool {
	const epsilon = 1e-9
	cross := (px-x1)*(y2-y1) - (py-y1)*(x2-x1)
	if cross > epsilon || cross < -epsilon {
		return false
	}
	dot := (px-x1)*(px-x2) + (py-y1)*(py-y2)
	return dot <= epsilon
}

func coordinateSlice(value any) ([]any, bool) {
	if value == nil {
		return nil, false
	}
	reflected := reflect.ValueOf(value)
	if reflected.Kind() != reflect.Slice && reflected.Kind() != reflect.Array {
		return nil, false
	}
	result := make([]any, reflected.Len())
	for i := 0; i < reflected.Len(); i++ {
		result[i] = reflected.Index(i).Interface()
	}
	return result, true
}

func position(value any) (float64, float64, bool) {
	parts, ok := coordinateSlice(value)
	if !ok || len(parts) < 2 {
		return 0, 0, false
	}
	x, okX := number(parts[0])
	y, okY := number(parts[1])
	return x, y, okX && okY
}

func number(value any) (float64, bool) {
	switch n := value.(type) {
	case float64:
		return n, true
	case float32:
		return float64(n), true
	case int:
		return float64(n), true
	case int64:
		return float64(n), true
	case json.Number:
		v, err := n.Float64()
		return v, err == nil
	default:
		return 0, false
	}
}
