package snapshot

// BoundaryMapRecord carries only the fields required to render and select a
// boundary. Raw source attributes and detail-only metadata stay behind the
// feature-detail endpoint.
type BoundaryMapRecord struct {
	LayerFamily    string   `json:"layer_family"`
	FeatureID      string   `json:"feature_id"`
	Title          string   `json:"title"`
	Subtitle       string   `json:"subtitle"`
	SourceID       string   `json:"source_id"`
	GeometrySource string   `json:"geometry_source"`
	Geometry       Geometry `json:"geometry"`
}

func (feature BoundaryFeature) MapRecord() BoundaryMapRecord {
	return BoundaryMapRecord{
		LayerFamily:    feature.LayerFamily,
		FeatureID:      feature.FeatureID,
		Title:          feature.Title,
		Subtitle:       feature.Subtitle,
		SourceID:       feature.SourceID,
		GeometrySource: feature.GeometrySource,
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
