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

type Candidate struct {
	MatchedAddress string    `json:"matched_address"`
	Longitude      float64   `json:"longitude"`
	Latitude       float64   `json:"latitude"`
	Precision      Precision `json:"precision"`
	Confidence     float64   `json:"confidence"`
	Source         string    `json:"source"`
	SourceRecordID string    `json:"source_record_id"`
	SourceVintage  string    `json:"source_vintage"`
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
