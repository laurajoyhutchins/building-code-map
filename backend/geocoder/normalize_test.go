package geocoder

import "testing"

func TestParseAddressNormalizesCivicAddress(t *testing.T) {
	parsed, err := ParseAddress(" 1600 North Broadway Street, Denver, Colorado 80202-1234 ")
	if err != nil {
		t.Fatal(err)
	}
	if parsed.HouseNumber != "1600" {
		t.Fatalf("house number=%q", parsed.HouseNumber)
	}
	if parsed.Street != "N BROADWAY ST" {
		t.Fatalf("street=%q", parsed.Street)
	}
	if parsed.City != "DENVER" || parsed.State != "CO" || parsed.PostalCode != "80202" {
		t.Fatalf("locality=%#v", parsed)
	}
	if parsed.Normalized != "1600 N BROADWAY ST, DENVER, CO 80202" {
		t.Fatalf("normalized=%q", parsed.Normalized)
	}
}

func TestParseAddressRejectsPOBox(t *testing.T) {
	if _, err := ParseAddress("PO Box 123, Denver, CO 80202"); err == nil {
		t.Fatal("expected PO box rejection")
	}
}

func TestParseAddressRejectsIncompleteAddress(t *testing.T) {
	if _, err := ParseAddress("1600 Broadway, Denver"); err == nil {
		t.Fatal("expected incomplete address rejection")
	}
}
