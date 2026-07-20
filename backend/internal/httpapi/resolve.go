package httpapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strings"

	"building-code-map/backend/internal/regulatory"
	"building-code-map/backend/internal/snapshot"
)

func (h *Handler) handleResolve(w http.ResponseWriter, r *http.Request) {
	if h.regulatoryCatalog.Len() == 0 {
		h.writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "regulatory policy catalog is unavailable"})
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	var request regulatory.ResolutionRequest
	if err := decoder.Decode(&request); err != nil {
		h.writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid resolution request: " + err.Error()})
		return
	}
	if err := ensureSingleJSONValue(decoder); err != nil {
		h.writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if request.Context == nil {
		if request.Point == nil {
			h.writeJSON(w, http.StatusBadRequest, map[string]string{"error": "point or context is required"})
			return
		}
		context, err := resolveGeographicContext(h.snapshot, h.regulatoryCatalog, *request.Point)
		if err != nil {
			h.writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
			return
		}
		request.Context = &context
	}
	result, err := regulatory.Resolve(h.regulatoryCatalog, request)
	if err != nil {
		h.writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
		return
	}
	h.writeJSON(w, http.StatusOK, result)
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
	context := regulatory.GeographicContext{}
	for _, feature := range snap.BoundaryFeatures {
		if !geometryContainsPoint(feature.Geometry, point.Longitude, point.Latitude) {
			continue
		}
		match := regulatory.BoundaryMatch{LayerFamily: feature.LayerFamily, FeatureID: feature.FeatureID, Name: feature.Title, SourceID: feature.SourceID}
		switch feature.LayerFamily {
		case "states":
			fips := firstString(feature.Attributes, "STATEFP", "GEOID")
			if fips == "" && len(feature.FeatureID) == 2 {
				fips = feature.FeatureID
			}
			context.StateFIPS = fips
			context.StateName = feature.Title
			if profile, ok := catalog.Profile("", fips); ok {
				context.StateID = profile.StateID
			}
		case "counties":
			if context.County == nil {
				context.County = &match
			}
		case "municipalities":
			if context.Municipality == nil {
				context.Municipality = &match
				context.Incorporated = true
			}
		case "special_areas":
			context.SpecialAreas = append(context.SpecialAreas, match)
		case "tribal_areas":
			context.TribalAreas = append(context.TribalAreas, match)
		case "neris_jurisdictions":
			context.FireJurisdictions = append(context.FireJurisdictions, match)
		}
	}
	if context.StateFIPS == "" && context.StateID == "" {
		return regulatory.GeographicContext{}, errors.New("point did not match a supported state boundary")
	}
	return context, nil
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
	rv := reflect.ValueOf(value)
	if rv.Kind() != reflect.Slice && rv.Kind() != reflect.Array {
		return nil, false
	}
	result := make([]any, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		result[i] = rv.Index(i).Interface()
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
