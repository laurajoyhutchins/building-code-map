package geocoder

import (
	"errors"
	"testing"
)

func TestParseAddressNormalizesSyntheticCivicAddress(t *testing.T) {
	parsed, err := ParseAddress(" 100 North Main Street, Demo City, Colorado 80202-1234 ")
	if err != nil {
		t.Fatal(err)
	}
	if parsed.HouseNumber != "100" || parsed.Street != "N MAIN ST" {
		t.Fatalf("unexpected street: %#v", parsed)
	}
	if parsed.City != "DEMO CITY" || parsed.State != "CO" || parsed.PostalCode != "80202" {
		t.Fatalf("unexpected locality: %#v", parsed)
	}
	if parsed.Normalized != "100 N MAIN ST, DEMO CITY, CO 80202" {
		t.Fatalf("normalized = %q", parsed.Normalized)
	}
}

func TestParseAddressRejectsUnusableInputs(t *testing.T) {
	for _, input := range []string{
		"",
		"PO Box 123, Demo City, CO 80202",
		"100 Demo St, Demo City",
		"100 Demo St, Demo City, XX",
		"Demo Main St, Demo City, CO",
	} {
		if _, err := ParseAddress(input); !errors.Is(err, ErrInvalidAddress) {
			t.Errorf("ParseAddress(%q) = %v, want ErrInvalidAddress", input, err)
		}
	}
}

func TestParseAddressAcceptsPostalCodeOmissionButDoesNotInventOne(t *testing.T) {
	parsed, err := ParseAddress("500 West Example Boulevard, Demo City, CO")
	if err != nil {
		t.Fatal(err)
	}
	if parsed.PostalCode != "" || parsed.Normalized != "500 W EXAMPLE BLVD, DEMO CITY, CO" {
		t.Fatalf("unexpected inferred postal code: %#v", parsed)
	}
}
