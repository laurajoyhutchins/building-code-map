package httpapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"sort"
	"strings"

	"building-code-map/backend/internal/regulatory"
	"building-code-map/backend/internal/snapshot"
)

type pointResolutionRequest struct {
	Point             *regulatory.Point `json:"point"`
	CodeFamily        string            `json:"code_family,omitempty"`
	ProjectType       string            `json:"project_type,omitempty"`
	ApplicabilityDate string            `json:"applicability_date,omitempty"`
}

type boundaryAmbiguityError struct {
	LayerFamily string
	Matches     []regulatory.BoundaryMatch
}

func (err *boundaryAmbiguityError) Error() string {
	return fmt.Sprintf("point matched multiple %s boundary observations; confirm the controlling boundary locally", err.LayerFamily)
}

type boundaryAmbiguityResponse struct {
	Error        string                     `json:"error"`
	Code         string                     `json:"code"`
	LayerFamily  string                     `json:"layer_family"`
	Observations []regulatory.BoundaryMatch `json:"observations"`
}

func (h *Handler) handleResolve(w http.ResponseWriter, r *http.Request) {
	if h.regulatoryCatalog.Len() == 0 {
		h.writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "regulatory policy catalog is unavailable"})
		return
	}

	request, err := decodeStrictRequest[pointResolutionRequest](w, r)
	if err != nil {
		h.writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if request.Point == nil {
		h.writeJSON(w, http.StatusBadRequest, map[string]string{"error": "point is required"})
		return
	}

	result, err := h.resolveRequest(regulatory.ResolutionRequest{
		Point:             request.Point,
		CodeFamily:        request.CodeFamily,
		ProjectType:       request.ProjectType,
		ApplicabilityDate: request.ApplicabilityDate,
	})
	if err != nil {
		h.writeResolutionError(w, err)
		return
	}
	h.writeJSON(w, http.StatusOK, result)
}

func (h *Handler) writeResolutionError(w http.ResponseWriter, err error) {
	var ambiguity *boundaryAmbiguityError
	if errors.As(err, &ambiguity) {
		h.writeJSON(w, http.StatusConflict, boundaryAmbiguityResponse{
			Error:        ambiguity.Error(),
			Code:         "boundary_ambiguous",
			LayerFamily:  ambiguity.LayerFamily,
			Observations: ambiguity.Matches,
		})
		return
	}
	h.writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
}

func ensureSingleJSONValue(decoder *json.Decoder) error {
	var extra any
	if err := decoder.Decode(&extra); err == io.EOF {
		return nil
	} else if err != nil {
		return fmt.Errorf("invalid trailing JSON: %w", err)
	}
	return errors.New("request body must contain one JSON value")
}

func resolveGeographicContext(snap snapshot.Snapshot, catalog regulatory.Catalog, point regulatory.Point) (regulatory.GeographicContext, error) {
	if point.Longitude < -180 || point.Longitude > 180 || point.Latitude < -90 || point.Latitude > 90 {
		return regulatory.GeographicContext{}, errors.New("point coordinates are outside valid longitude/latitude ranges")
	}

	matchedFeatures := map[string][]snapshot.BoundaryFeature{}
	for _, feature := range snap.BoundaryFeatures {
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
			return regulatory.GeographicContext{}, &boundaryAmbiguityError{
				LayerFamily: layerFamily,
				Matches:     boundaryMatches(matchedFeatures[layerFamily]),
			}
		}
	}

	context := regulatory.GeographicContext{}
	stateFeatures := matchedFeatures["states"]
	if len(stateFeatures) == 0 {
		return regulatory.GeographicContext{}, errors.New("point did not match a supported state boundary")
	}
	stateFeature := stateFeatures[0]
	fips := firstString(stateFeature.Attributes, "STATEFP", "GEOID")
	if fips == "" && len(stateFeature.FeatureID) == 2 {
		fips = stateFeature.FeatureID
	}
	context.StateFIPS = fips
	context.StateName = stateFeature.Title
	if profile, ok := catalog.Profile("", fips); ok {
		context.StateID = profile.StateID
	}

	if counties := matchedFeatures["counties"]; len(counties) == 1 {
		match := boundaryMatch(counties[0])
		context.County = &match
	}
	if municipalities := matchedFeatures["municipalities"]; len(municipalities) == 1 {
		match := boundaryMatch(municipalities[0])
		context.Municipality = &match
		context.Incorporated = true
	}
	context.SpecialAreas = boundaryMatches(matchedFeatures["special_areas"])
	context.TribalAreas = boundaryMatches(matchedFeatures["tribal_areas"])
	context.FireJurisdictions = boundaryMatches(matchedFeatures["neris_jurisdictions"])
	return context, nil
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
