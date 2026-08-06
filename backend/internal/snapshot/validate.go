package snapshot

import (
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"regexp"
	"strings"
)

var layerFamilyKeyPattern = regexp.MustCompile(`^[a-z][a-z0-9_]*$`)

// Validate checks the semantic contract required by boundary resolution and the
// HTTP API. Loaders must call it before returning a snapshot to the runtime.
func (snap Snapshot) Validate() error {
	if len(snap.LayerFamilies) == 0 {
		return fmt.Errorf("%w: no layer families", errInvalidSnapshot)
	}

	layerKeys := make(map[string]struct{}, len(snap.LayerFamilies))
	for index, layer := range snap.LayerFamilies {
		key := strings.TrimSpace(layer.Key)
		if !layerFamilyKeyPattern.MatchString(key) {
			return fmt.Errorf("%w: layer_families[%d].key %q is invalid", errInvalidSnapshot, index, layer.Key)
		}
		if _, exists := layerKeys[key]; exists {
			return fmt.Errorf("%w: duplicate layer family %q", errInvalidSnapshot, key)
		}
		if strings.TrimSpace(layer.Label) == "" {
			return fmt.Errorf("%w: layer family %q has an empty label", errInvalidSnapshot, key)
		}
		if strings.TrimSpace(layer.Description) == "" {
			return fmt.Errorf("%w: layer family %q has an empty description", errInvalidSnapshot, key)
		}
		layerKeys[key] = struct{}{}
	}

	if len(snap.BoundaryFeatures) == 0 {
		return fmt.Errorf("%w: no boundary features", errInvalidSnapshot)
	}

	featureKeys := make(map[string]struct{}, len(snap.BoundaryFeatures))
	for index, feature := range snap.BoundaryFeatures {
		layerFamily := strings.TrimSpace(feature.LayerFamily)
		featureID := strings.TrimSpace(feature.FeatureID)
		if _, exists := layerKeys[layerFamily]; !exists {
			return fmt.Errorf("%w: boundary_features[%d] references unknown layer family %q", errInvalidSnapshot, index, feature.LayerFamily)
		}
		if featureID == "" {
			return fmt.Errorf("%w: boundary_features[%d].feature_id is required", errInvalidSnapshot, index)
		}
		key := layerFamily + "\x00" + featureID
		if _, exists := featureKeys[key]; exists {
			return fmt.Errorf("%w: duplicate boundary feature %s/%s", errInvalidSnapshot, layerFamily, featureID)
		}
		featureKeys[key] = struct{}{}

		if strings.TrimSpace(feature.Title) == "" {
			return fmt.Errorf("%w: boundary feature %s/%s has an empty title", errInvalidSnapshot, layerFamily, featureID)
		}
		if strings.TrimSpace(feature.SourceID) == "" {
			return fmt.Errorf("%w: boundary feature %s/%s has an empty source_id", errInvalidSnapshot, layerFamily, featureID)
		}
		if strings.TrimSpace(feature.GeometryLabel) == "" {
			return fmt.Errorf("%w: boundary feature %s/%s has an empty geometry_label", errInvalidSnapshot, layerFamily, featureID)
		}
		if feature.LastSyncedAt.IsZero() {
			return fmt.Errorf("%w: boundary feature %s/%s has an empty last_synced_at", errInvalidSnapshot, layerFamily, featureID)
		}
		if feature.Attributes == nil {
			return fmt.Errorf("%w: boundary feature %s/%s attributes must be an object", errInvalidSnapshot, layerFamily, featureID)
		}
		if err := validateGeometry(feature.Geometry); err != nil {
			return fmt.Errorf("%w: boundary feature %s/%s: %v", errInvalidSnapshot, layerFamily, featureID, err)
		}
	}

	if err := validateRefreshStatus(snap.RefreshStatus); err != nil {
		return fmt.Errorf("%w: %v", errInvalidSnapshot, err)
	}
	return nil
}

func validateRefreshStatus(status RefreshStatus) error {
	switch status.Status {
	case "ok", "warning", "error":
	default:
		return fmt.Errorf("refresh status %q is unsupported", status.Status)
	}
	if status.LatestSuccessfulRefresh.IsZero() {
		return fmt.Errorf("latest_successful_refresh is required")
	}
	if status.LatestAttempt.IsZero() {
		return fmt.Errorf("latest_attempt is required")
	}
	if status.NextScheduledRefresh.IsZero() {
		return fmt.Errorf("next_scheduled_refresh is required")
	}
	if status.LatestSuccessfulRefresh.After(status.LatestAttempt) {
		return fmt.Errorf("latest_successful_refresh cannot be after latest_attempt")
	}
	if status.NextScheduledRefresh.Before(status.LatestAttempt) {
		return fmt.Errorf("next_scheduled_refresh cannot be before latest_attempt")
	}
	if strings.TrimSpace(status.Message) == "" {
		return fmt.Errorf("refresh status message is required")
	}
	return nil
}

func validateGeometry(geometry Geometry) error {
	switch geometry.Type {
	case "Polygon":
		return validatePolygonCoordinates(geometry.Coordinates)
	case "MultiPolygon":
		polygons, ok := sequence(geometry.Coordinates)
		if !ok || len(polygons) == 0 {
			return fmt.Errorf("MultiPolygon coordinates must contain at least one polygon")
		}
		for index, polygon := range polygons {
			if err := validatePolygonCoordinates(polygon); err != nil {
				return fmt.Errorf("polygon %d: %w", index, err)
			}
		}
		return nil
	default:
		return fmt.Errorf("unsupported geometry type %q", geometry.Type)
	}
}

func validatePolygonCoordinates(value any) error {
	rings, ok := sequence(value)
	if !ok || len(rings) == 0 {
		return fmt.Errorf("Polygon coordinates must contain at least one ring")
	}
	for index, ring := range rings {
		if err := validateRing(ring); err != nil {
			return fmt.Errorf("ring %d: %w", index, err)
		}
	}
	return nil
}

func validateRing(value any) error {
	positions, ok := sequence(value)
	if !ok || len(positions) < 4 {
		return fmt.Errorf("ring must contain at least four positions")
	}

	var firstLongitude, firstLatitude float64
	for index, rawPosition := range positions {
		position, ok := sequence(rawPosition)
		if !ok || len(position) < 2 {
			return fmt.Errorf("position %d must contain longitude and latitude", index)
		}
		longitude, okLongitude := coordinateNumber(position[0])
		latitude, okLatitude := coordinateNumber(position[1])
		if !okLongitude || !okLatitude || math.IsNaN(longitude) || math.IsNaN(latitude) || math.IsInf(longitude, 0) || math.IsInf(latitude, 0) {
			return fmt.Errorf("position %d contains a non-finite coordinate", index)
		}
		if longitude < -180 || longitude > 180 || latitude < -90 || latitude > 90 {
			return fmt.Errorf("position %d is outside longitude/latitude ranges", index)
		}
		if index == 0 {
			firstLongitude, firstLatitude = longitude, latitude
		}
		if index == len(positions)-1 && (longitude != firstLongitude || latitude != firstLatitude) {
			return fmt.Errorf("ring is not closed")
		}
	}
	return nil
}

func sequence(value any) ([]any, bool) {
	if value == nil {
		return nil, false
	}
	reflected := reflect.ValueOf(value)
	if reflected.Kind() != reflect.Slice && reflected.Kind() != reflect.Array {
		return nil, false
	}
	items := make([]any, reflected.Len())
	for index := 0; index < reflected.Len(); index++ {
		items[index] = reflected.Index(index).Interface()
	}
	return items, true
}

func coordinateNumber(value any) (float64, bool) {
	switch number := value.(type) {
	case float64:
		return number, true
	case float32:
		return float64(number), true
	case int:
		return float64(number), true
	case int8:
		return float64(number), true
	case int16:
		return float64(number), true
	case int32:
		return float64(number), true
	case int64:
		return float64(number), true
	case uint:
		return float64(number), true
	case uint8:
		return float64(number), true
	case uint16:
		return float64(number), true
	case uint32:
		return float64(number), true
	case uint64:
		return float64(number), true
	case json.Number:
		parsed, err := number.Float64()
		return parsed, err == nil
	default:
		return 0, false
	}
}
