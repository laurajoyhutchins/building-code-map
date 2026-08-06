package snapshot

import "time"

// BoundaryMapRecord carries only the fields required to render and select a
// boundary. Raw source attributes remain behind the feature-detail endpoint.
type BoundaryMapRecord struct {
	LayerFamily    string   `json:"layer_family"`
	FeatureID      string   `json:"feature_id"`
	Title          string   `json:"title"`
	Subtitle       string   `json:"subtitle"`
	SourceID       string   `json:"source_id"`
	GeometryLabel  string   `json:"geometry_label"`
	GeometrySource string   `json:"geometry_source,omitempty"`
	LastSyncedAt   time.Time `json:"last_synced_at"`
	Geometry       Geometry `json:"geometry"`
}

func (feature BoundaryFeature) MapRecord() BoundaryMapRecord {
	return BoundaryMapRecord{
		LayerFamily:    feature.LayerFamily,
		FeatureID:      feature.FeatureID,
		Title:          feature.Title,
		Subtitle:       feature.Subtitle,
		SourceID:       feature.SourceID,
		GeometryLabel:  feature.GeometryLabel,
		GeometrySource: feature.GeometrySource,
		LastSyncedAt:   feature.LastSyncedAt,
		Geometry:       feature.Geometry,
	}
}

func MapRecords(features []BoundaryFeature) []BoundaryMapRecord {
	records := make([]BoundaryMapRecord, 0, len(features))
	for _, feature := range features {
		records = append(records, feature.MapRecord())
	}
	return records
}
