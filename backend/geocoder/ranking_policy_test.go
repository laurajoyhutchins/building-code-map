package geocoder

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"
	"strings"
	"testing"

	_ "modernc.org/sqlite"
)

const rankingAddressCSV = `source_record_id,address_number,street,city,state,postal_code,longitude,latitude,matched_address
base,100,MAIN ST,DENVER,CO,80202,-104.9900,39.7400,100 Main St Denver CO 80202
`

func buildRankingSnapshot(t *testing.T, sourceName, sourceVintage, contents string) string {
	t.Helper()
	directory := t.TempDir()
	input := filepath.Join(directory, "input.csv")
	if err := os.WriteFile(input, []byte(contents), 0o600); err != nil {
		t.Fatal(err)
	}
	output := filepath.Join(directory, "geocoder.sqlite")
	options := BuildOptions{
		OutputPath:    output,
		SourceName:    sourceName,
		SourceVintage: sourceVintage,
	}
	if strings.Contains(contents, "address_number") {
		options.AddressPointsCSV = input
	} else {
		options.StreetRangesCSV = input
	}
	if err := BuildSnapshot(options); err != nil {
		t.Fatal(err)
	}
	return output
}

func insertAddressPoint(
	t *testing.T,
	path string,
	sourceName string,
	recordID string,
	vintage string,
	longitude float64,
	latitude float64,
) {
	t.Helper()
	db, err := sql.Open("sqlite", path)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	_, err = db.Exec(`
INSERT INTO address_points (
  address_number, street_name, city, state, postal_code, matched_address,
  longitude, latitude, source_name, source_record_id, source_vintage
) VALUES ('100', 'MAIN ST', 'DENVER', 'CO', '80202',
  '100 Main St Denver CO 80202', ?, ?, ?, ?, ?);`,
		longitude,
		latitude,
		sourceName,
		recordID,
		vintage,
	)
	if err != nil {
		t.Fatal(err)
	}
}

func TestConfiguredSourcePriorityChangesSelectionAndIsVersioned(t *testing.T) {
	path := buildRankingSnapshot(t, "base-source", "2026-08-01", rankingAddressCSV)
	insertAddressPoint(t, path, "preferred-source", "preferred", "2025-01-01", -104.991, 39.741)

	policy := DefaultRankingPolicy()
	policy.Version = "geocoder-ranking-priority-fixture"
	policy.SourcePriority = map[string]float64{"preferred-source": 0.10}
	service, err := OpenWithPolicy(path, policy)
	if err != nil {
		t.Fatal(err)
	}
	defer service.Close()

	result, err := service.Geocode(
		context.Background(),
		Query{Address: "100 Main St, Denver, CO 80202"},
	)
	if err != nil {
		t.Fatal(err)
	}
	if result.Status != StatusMatched || result.Selected == nil {
		t.Fatalf("result=%#v", result)
	}
	if result.Selected.Source != "preferred-source" {
		t.Fatalf("selected source=%q", result.Selected.Source)
	}
	if result.Selected.ScoreFactors["source_priority"] != 0.10 {
		t.Fatalf("score factors=%#v", result.Selected.ScoreFactors)
	}
	if result.Selected.RankingPolicyVersion != policy.Version {
		t.Fatalf("ranking policy=%q", result.Selected.RankingPolicyVersion)
	}
}

func TestSourceVintageRemainsProvenanceOnly(t *testing.T) {
	path := buildRankingSnapshot(t, "current-source", "2026-08-01", rankingAddressCSV)
	insertAddressPoint(t, path, "older-source", "older", "2019-01-01", -104.991, 39.741)
	service, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer service.Close()

	result, err := service.Geocode(
		context.Background(),
		Query{Address: "100 Main St, Denver, CO 80202"},
	)
	if err != nil {
		t.Fatal(err)
	}
	if result.Status != StatusAmbiguous || result.Selected != nil || len(result.Candidates) != 2 {
		t.Fatalf("result=%#v", result)
	}
	vintages := map[string]bool{}
	for _, candidate := range result.Candidates {
		vintages[candidate.SourceVintage] = true
		if _, ranked := candidate.ScoreFactors["source_vintage"]; ranked {
			t.Fatalf("source vintage unexpectedly affected score: %#v", candidate.ScoreFactors)
		}
	}
	if !vintages["2026-08-01"] || !vintages["2019-01-01"] {
		t.Fatalf("vintages=%#v", vintages)
	}
}

func TestDuplicateAddressPointsRemainAmbiguous(t *testing.T) {
	path := buildRankingSnapshot(t, "fixture-source", "2026-08-01", rankingAddressCSV)
	insertAddressPoint(t, path, "fixture-source", "duplicate", "2026-08-01", -104.9900, 39.7400)
	service, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer service.Close()

	result, err := service.Geocode(
		context.Background(),
		Query{Address: "100 Main St, Denver, CO 80202"},
	)
	if err != nil {
		t.Fatal(err)
	}
	if result.Status != StatusAmbiguous || result.Selected != nil || len(result.Candidates) != 2 {
		t.Fatalf("result=%#v", result)
	}
}

func TestDescendingStreetRangePreservesDirectionAndFraction(t *testing.T) {
	contents := `source_record_id,from_number,to_number,parity,street,city,state,postal_code,from_longitude,from_latitude,to_longitude,to_latitude
range-descending,1598,1500,E,MARKET ST,DENVER,CO,80202,-104.9850,39.7480,-105.0000,39.7480
`
	path := buildRankingSnapshot(t, "range-source", "2026-08-01", contents)
	service, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer service.Close()

	result, err := service.Geocode(
		context.Background(),
		Query{Address: "1510 Market St, Denver, CO 80202"},
	)
	if err != nil {
		t.Fatal(err)
	}
	if result.Status != StatusMatched || result.Selected == nil || result.Selected.Interpolation == nil {
		t.Fatalf("result=%#v", result)
	}
	provenance := result.Selected.Interpolation
	if provenance.RangeDirection != "descending" {
		t.Fatalf("range direction=%q", provenance.RangeDirection)
	}
	if provenance.Fraction <= 0 || provenance.Fraction >= 1 {
		t.Fatalf("fraction=%f", provenance.Fraction)
	}
	if provenance.DerivedCoordinate.Longitude != result.Selected.Longitude ||
		provenance.DerivedCoordinate.Latitude != result.Selected.Latitude {
		t.Fatalf("derived coordinate=%#v selected=%#v", provenance.DerivedCoordinate, result.Selected)
	}
}

func TestSnapshotBuildRejectsZeroLengthStreetRange(t *testing.T) {
	directory := t.TempDir()
	input := filepath.Join(directory, "ranges.csv")
	output := filepath.Join(directory, "geocoder.sqlite")
	contents := `source_record_id,from_number,to_number,parity,street,city,state,postal_code,from_longitude,from_latitude,to_longitude,to_latitude
zero,1500,1500,E,MARKET ST,DENVER,CO,80202,-105.0000,39.7480,-104.9850,39.7480
`
	if err := os.WriteFile(input, []byte(contents), 0o600); err != nil {
		t.Fatal(err)
	}
	err := BuildSnapshot(BuildOptions{
		OutputPath:       output,
		StreetRangesCSV:  input,
		SourceName:       "range-source",
		SourceVintage:    "2026-08-01",
	})
	if err == nil || !strings.Contains(err.Error(), "zero length") {
		t.Fatalf("error=%v", err)
	}
	if _, statErr := os.Stat(output); !os.IsNotExist(statErr) {
		t.Fatalf("snapshot should not be published, stat error=%v", statErr)
	}
}
