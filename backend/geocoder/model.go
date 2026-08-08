package geocoder

import (
	"context"
	"errors"
)

var ErrInvalidAddress = errors.New("enter a civic address with house number, street, city, and state")

type Status string

const (
	StatusMatched   Status = "matched"
	StatusAmbiguous Status = "ambiguous"
	StatusNotFound  Status = "not_found"
)

type Precision string

const (
	PrecisionAddressPoint Precision = "address_point"
	PrecisionInterpolated Precision = "interpolated"
)

type Query struct {
	Address string `json:"address"`
}

type ParsedAddress struct {
	HouseNumber string `json:"house_number"`
	Street      string `json:"street"`
	City        string `json:"city"`
	State       string `json:"state"`
	PostalCode  string `json:"postal_code,omitempty"`
	Normalized  string `json:"normalized"`
}

type Coordinate struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type InterpolationProvenance struct {
	SourceRangeID             string     `json:"source_range_id"`
	RequestedHouseNumber      int        `json:"requested_house_number"`
	FromNumber                int        `json:"from_number"`
	ToNumber                  int        `json:"to_number"`
	RangeDirection            string     `json:"range_direction"`
	Parity                    string     `json:"parity"`
	Side                      string     `json:"side,omitempty"`
	FromCoordinate            Coordinate `json:"from_coordinate"`
	ToCoordinate              Coordinate `json:"to_coordinate"`
	Fraction                  float64    `json:"fraction"`
	DerivedCoordinate         Coordinate `json:"derived_coordinate"`
	CoordinateReferenceSystem string     `json:"coordinate_reference_system"`
	TransformationIdentity    string     `json:"transformation_identity"`
	MethodVersion             string     `json:"method_version"`
	PositionalQuality         string     `json:"positional_quality"`
}

type Candidate struct {
	MatchedAddress      string                   `json:"matched_address"`
	Longitude           float64                  `json:"longitude"`
	Latitude            float64                  `json:"latitude"`
	Precision           Precision                `json:"precision"`
	Confidence          float64                  `json:"confidence"`
	ScoreKind           string                   `json:"score_kind"`
	ScoreFactors        map[string]float64       `json:"score_factors"`
	RankingPolicyVersion string                  `json:"ranking_policy_version"`
	Source              string                  `json:"source"`
	SourceRecordID      string                  `json:"source_record_id"`
	SourceVintage       string                  `json:"source_vintage"`
	Interpolation       *InterpolationProvenance `json:"interpolation,omitempty"`
}

type Result struct {
	Query      string      `json:"query"`
	Normalized string      `json:"normalized,omitempty"`
	Status     Status      `json:"status"`
	Selected   *Candidate  `json:"selected,omitempty"`
	Candidates []Candidate `json:"candidates"`
	Warnings   []string    `json:"warnings"`
}

type Service interface {
	Geocode(context.Context, Query) (Result, error)
}

type BuildOptions struct {
	OutputPath       string
	AddressPointsCSV string
	StreetRangesCSV  string
	SourceName       string
	SourceVintage    string
}
