package engine

import (
	"errors"
	"testing"
	"time"
)

func TestNormalizeQueryRequiresExactlyOneLocation(t *testing.T) {
	tests := []struct {
		name  string
		query Query
	}{
		{name: "neither", query: Query{ApplicabilityDate: "2026-08-06"}},
		{name: "both", query: Query{
			Point:             &Point{Longitude: -104.99, Latitude: 39.74},
			Address:           "1600 Pennsylvania Ave NW",
			ApplicabilityDate: "2026-08-06",
		}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := NormalizeQuery(test.query)
			if err == nil {
				t.Fatal("NormalizeQuery() error = nil, want invalid_query")
			}
			var engineErr EngineError
			if !errors.As(err, &engineErr) {
				t.Fatalf("NormalizeQuery() error = %T %v, want EngineError", err, err)
			}
			if engineErr.Code != ErrorInvalidQuery {
				t.Fatalf("error code = %q, want %q", engineErr.Code, ErrorInvalidQuery)
			}
		})
	}
}

func TestNormalizeQueryRequiresExplicitApplicabilityDate(t *testing.T) {
	_, err := NormalizeQuery(Query{Point: &Point{Longitude: -104.99, Latitude: 39.74}})
	if err == nil {
		t.Fatal("NormalizeQuery() error = nil, want invalid_query")
	}
	var engineErr EngineError
	if !errors.As(err, &engineErr) {
		t.Fatalf("NormalizeQuery() error = %T %v, want EngineError", err, err)
	}
	if engineErr.Code != ErrorInvalidQuery {
		t.Fatalf("error code = %q, want %q", engineErr.Code, ErrorInvalidQuery)
	}
}

func TestNormalizeQueryProducesStableNormalizedValues(t *testing.T) {
	got, err := NormalizeQuery(Query{
		Point:             &Point{Longitude: -104.99, Latitude: 39.74},
		CodeFamily:        " Building ",
		ProjectType:       " New Construction ",
		ApplicabilityDate: "2026-08-06",
		Include:           []IncludeKey{"evidence", "location", "evidence"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if got.CodeFamily != "building" || got.ProjectType != "new construction" {
		t.Fatalf("normalized query = %#v", got)
	}
	if got.ApplicabilityDate != "2026-08-06" {
		t.Fatalf("applicability date = %q", got.ApplicabilityDate)
	}
	if len(got.Include) != 2 || got.Include[0] != "evidence" || got.Include[1] != "location" {
		t.Fatalf("include = %#v", got.Include)
	}
}

func TestNormalizeQueryRejectsInvalidCoordinates(t *testing.T) {
	_, err := NormalizeQuery(Query{
		Point:             &Point{Longitude: -181, Latitude: 39.74},
		ApplicabilityDate: "2026-08-06",
	})
	if err == nil {
		t.Fatal("NormalizeQuery() error = nil, want invalid_coordinates")
	}
	var engineErr EngineError
	if !errors.As(err, &engineErr) {
		t.Fatalf("NormalizeQuery() error = %T %v, want EngineError", err, err)
	}
	if engineErr.Code != ErrorInvalidCoordinates {
		t.Fatalf("error code = %q, want %q", engineErr.Code, ErrorInvalidCoordinates)
	}
}

func TestFixedClockReturnsConfiguredUTCValue(t *testing.T) {
	want := time.Date(2026, 8, 6, 12, 34, 56, 0, time.FixedZone("MDT", -6*60*60))
	clock := NewFixedClock(want)
	if got := clock.Now(); !got.Equal(want.UTC()) {
		t.Fatalf("FixedClock.Now() = %s, want %s", got, want.UTC())
	}
}

func TestEngineErrorCarriesStableCodeAndDetails(t *testing.T) {
	err := EngineError{
		Code:      ErrorBoundaryAmbiguous,
		Message:   "multiple boundary observations matched",
		Details:   map[string]any{"layer_family": "states"},
		Retryable: false,
	}
	if err.Error() != "multiple boundary observations matched" {
		t.Fatalf("EngineError.Error() = %q", err.Error())
	}
	if err.Code != ErrorBoundaryAmbiguous || err.Details["layer_family"] != "states" {
		t.Fatalf("engine error = %#v", err)
	}
}
