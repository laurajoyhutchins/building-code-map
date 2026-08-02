package geocoder

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func buildFixtureService(t *testing.T) *SQLiteService {
	t.Helper()
	output := filepath.Join(t.TempDir(), "geocoder.sqlite")
	if err := BuildSnapshot(BuildOptions{
		OutputPath:       output,
		AddressPointsCSV: filepath.Join("testdata", "address_points.csv"),
		StreetRangesCSV:  filepath.Join("testdata", "street_ranges.csv"),
		SourceName:       "fixture-addresses",
		SourceVintage:    "2026-08-01",
	}); err != nil {
		t.Fatal(err)
	}
	service, err := Open(output)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = service.Close() })
	return service
}

func buildAddressPointService(t *testing.T, contents string) *SQLiteService {
	t.Helper()
	directory := t.TempDir()
	input := filepath.Join(directory, "address-points.csv")
	if err := os.WriteFile(input, []byte(contents), 0o600); err != nil {
		t.Fatal(err)
	}
	output := filepath.Join(directory, "geocoder.sqlite")
	if err := BuildSnapshot(BuildOptions{
		OutputPath:       output,
		AddressPointsCSV: input,
		SourceName:       "generated-fixture",
		SourceVintage:    "2026-08-01",
	}); err != nil {
		t.Fatal(err)
	}
	service, err := Open(output)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = service.Close() })
	return service
}

func TestSQLiteServiceReturnsExactAddressPointWithProvenance(t *testing.T) {
	service := buildFixtureService(t)
	result, err := service.Geocode(context.Background(), Query{Address: "1600 N Broadway, Denver, CO 80202"})
	if err != nil {
		t.Fatal(err)
	}
	if result.Status != StatusMatched || result.Selected == nil {
		t.Fatalf("result=%#v", result)
	}
	if result.Selected.Precision != PrecisionAddressPoint {
		t.Fatalf("precision=%q", result.Selected.Precision)
	}
	if result.Selected.Source != "fixture-addresses" || result.Selected.SourceVintage != "2026-08-01" {
		t.Fatalf("provenance=%#v", result.Selected)
	}
	if result.Selected.SourceRecordID != "co-denver-1600" {
		t.Fatalf("source record=%q", result.Selected.SourceRecordID)
	}
}

func TestSQLiteServiceRejectsContradictoryLocalityEvidence(t *testing.T) {
	service := buildFixtureService(t)
	for _, address := range []string{
		"1600 N Broadway, Denver, CO 99999",
		"1600 N Broadway, Aurora, CO 80202",
	} {
		t.Run(address, func(t *testing.T) {
			result, err := service.Geocode(context.Background(), Query{Address: address})
			if err != nil {
				t.Fatal(err)
			}
			if result.Status != StatusNotFound || result.Selected != nil {
				t.Fatalf("result=%#v", result)
			}
		})
	}
}

func TestSQLiteServiceAppliesCandidateLimitAfterLocalityConstraints(t *testing.T) {
	var input strings.Builder
	input.WriteString("source_record_id,address_number,street,city,state,postal_code,longitude,latitude,matched_address\n")
	for index := 0; index < 20; index++ {
		fmt.Fprintf(
			&input,
			"%02d-decoy,100,MAIN ST,AURORA,CO,80010,-104.8000,39.7000,100 Main St Aurora CO 80010\n",
			index,
		)
	}
	input.WriteString("zz-denver,100,MAIN ST,DENVER,CO,80202,-104.9900,39.7400,100 Main St Denver CO 80202\n")

	service := buildAddressPointService(t, input.String())
	result, err := service.Geocode(context.Background(), Query{Address: "100 Main St, Denver, CO 80202"})
	if err != nil {
		t.Fatal(err)
	}
	if result.Status != StatusMatched || result.Selected == nil {
		t.Fatalf("result=%#v", result)
	}
	if result.Selected.SourceRecordID != "zz-denver" {
		t.Fatalf("source record=%q", result.Selected.SourceRecordID)
	}
}

func TestSQLiteServiceReturnsAmbiguousCandidatesWithoutSelectingOne(t *testing.T) {
	service := buildFixtureService(t)
	result, err := service.Geocode(context.Background(), Query{Address: "100 Main St, Springfield, NJ 07081"})
	if err != nil {
		t.Fatal(err)
	}
	if result.Status != StatusAmbiguous || result.Selected != nil {
		t.Fatalf("result=%#v", result)
	}
	if len(result.Candidates) != 2 {
		t.Fatalf("candidates=%d", len(result.Candidates))
	}
}

func TestSQLiteServiceReturnsNotFound(t *testing.T) {
	service := buildFixtureService(t)
	result, err := service.Geocode(context.Background(), Query{Address: "999 Missing Rd, Denver, CO 80202"})
	if err != nil {
		t.Fatal(err)
	}
	if result.Status != StatusNotFound || result.Selected != nil || len(result.Candidates) != 0 {
		t.Fatalf("result=%#v", result)
	}
}

func TestSQLiteServiceInterpolatesCompatibleStreetRange(t *testing.T) {
	service := buildFixtureService(t)
	result, err := service.Geocode(context.Background(), Query{Address: "1510 Market St, Denver, CO 80202"})
	if err != nil {
		t.Fatal(err)
	}
	if result.Status != StatusMatched || result.Selected == nil {
		t.Fatalf("result=%#v", result)
	}
	if result.Selected.Precision != PrecisionInterpolated {
		t.Fatalf("precision=%q", result.Selected.Precision)
	}
	if len(result.Warnings) == 0 {
		t.Fatal("expected interpolation warning")
	}
}

func TestSQLiteServiceRejectsWrongRangeParity(t *testing.T) {
	service := buildFixtureService(t)
	result, err := service.Geocode(context.Background(), Query{Address: "1511 Market St, Denver, CO 80202"})
	if err != nil {
		t.Fatal(err)
	}
	if result.Status != StatusNotFound {
		t.Fatalf("result=%#v", result)
	}
}
