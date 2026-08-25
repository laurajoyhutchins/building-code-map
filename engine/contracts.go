package engine

import (
	"math"
	"sort"
	"strings"
	"time"
)

const SchemaVersion = "1"

type IncludeKey string

type Point struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type Query struct {
	Point             *Point       `json:"point,omitempty"`
	Address           string       `json:"address,omitempty"`
	CodeFamily        string       `json:"code_family,omitempty"`
	ProjectType       string       `json:"project_type,omitempty"`
	ApplicabilityDate string       `json:"applicability_date"`
	Include           []IncludeKey `json:"include,omitempty"`
}

type NormalizedQuery struct {
	Point             *Point       `json:"point,omitempty"`
	Address           string       `json:"address,omitempty"`
	CodeFamily        string       `json:"code_family,omitempty"`
	ProjectType       string       `json:"project_type,omitempty"`
	ApplicabilityDate string       `json:"applicability_date"`
	Include           []IncludeKey `json:"include,omitempty"`
}

func NormalizeQuery(query Query) (NormalizedQuery, error) {
	address := strings.TrimSpace(query.Address)
	pointProvided := query.Point != nil
	addressProvided := address != ""
	if pointProvided == addressProvided {
		return NormalizedQuery{}, EngineError{Code: ErrorInvalidQuery, Message: "exactly one of point or address is required"}
	}

	applicabilityDate := strings.TrimSpace(query.ApplicabilityDate)
	if applicabilityDate == "" {
		return NormalizedQuery{}, EngineError{Code: ErrorInvalidQuery, Message: "applicability_date is required for engine queries"}
	}
	parsedDate, err := time.Parse(time.DateOnly, applicabilityDate)
	if err != nil {
		return NormalizedQuery{ApplicabilityDate: applicabilityDate}, EngineError{
			Code: ErrorInvalidQuery, Message: "applicability_date must use YYYY-MM-DD", Details: map[string]any{"value": applicabilityDate},
		}
	}

	normalized := NormalizedQuery{
		Address:           address,
		CodeFamily:        normalizeText(query.CodeFamily),
		ProjectType:       normalizeText(query.ProjectType),
		ApplicabilityDate: parsedDate.Format(time.DateOnly),
		Include:           normalizeIncludes(query.Include),
	}
	if pointProvided {
		if err := validatePoint(*query.Point); err != nil {
			return NormalizedQuery{}, err
		}
		point := *query.Point
		normalized.Point = &point
	}
	return normalized, nil
}

func normalizeText(value string) string { return strings.ToLower(strings.TrimSpace(value)) }

func normalizeIncludes(values []IncludeKey) []IncludeKey {
	seen := make(map[IncludeKey]struct{}, len(values))
	result := make([]IncludeKey, 0, len(values))
	for _, value := range values {
		normalized := IncludeKey(normalizeText(string(value)))
		if normalized == "" {
			continue
		}
		if _, exists := seen[normalized]; exists {
			continue
		}
		seen[normalized] = struct{}{}
		result = append(result, normalized)
	}
	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
	return result
}

func validatePoint(point Point) error {
	if math.IsNaN(point.Longitude) || math.IsInf(point.Longitude, 0) || math.IsNaN(point.Latitude) || math.IsInf(point.Latitude, 0) {
		return EngineError{Code: ErrorInvalidCoordinates, Message: "point coordinates must be finite"}
	}
	if point.Longitude < -180 || point.Longitude > 180 || point.Latitude < -90 || point.Latitude > 90 {
		return EngineError{Code: ErrorInvalidCoordinates, Message: "point coordinates are outside valid longitude/latitude ranges"}
	}
	return nil
}
