package snapshot

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestBuildSQLiteProducesDeterministicValidatedBoundarySnapshot(t *testing.T) {
	root := t.TempDir()
	layers := writeFixtureCSV(t, root, "layers.csv", `key,label,martin_layer_id,description,default_enabled
counties,Counties,counties.boundaries,County boundary observations,true
`)
	features := writeFixtureCSV(t, root, "features.csv", `layer_family,feature_id,title,subtitle,source_id,geometry_label,geometry_source,last_synced_at,geometry_json,attributes_json
counties,001,Example County,Test,source-001,County polygon,fixture,2026-08-06T12:00:00Z,"{""type"":""Polygon"",""coordinates"":[[[-105,39],[-104,39],[-104,40],[-105,39]]]} ","{""state"":""CO""}"
`)
	refresh := writeFixtureCSV(t, root, "refresh.csv", `status,latest_successful_refresh,latest_attempt,next_scheduled_refresh,message
ok,2026-08-06T12:00:00Z,2026-08-06T12:00:00Z,2026-08-07T12:00:00Z,Fixture is current
`)
	first := filepath.Join(root, "first.sqlite")
	second := filepath.Join(root, "second.sqlite")
	options := func(output string) SQLiteBuildOptions {
		return SQLiteBuildOptions{
			OutputPath: output, LayerFamiliesCSV: layers,
			BoundaryFeaturesCSV: features, RefreshStatusCSV: refresh,
		}
	}
	if err := BuildSQLite(options(first)); err != nil {
		t.Fatal(err)
	}
	if err := BuildSQLite(options(second)); err != nil {
		t.Fatal(err)
	}
	firstBytes, err := os.ReadFile(first)
	if err != nil {
		t.Fatal(err)
	}
	secondBytes, err := os.ReadFile(second)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(firstBytes, secondBytes) {
		t.Fatal("identical canonical inputs did not produce byte-identical SQLite snapshots")
	}
	snap, err := LoadSQLite(first)
	if err != nil {
		t.Fatal(err)
	}
	if len(snap.BoundaryFeatures) != 1 || snap.BoundaryFeatures[0].FeatureID != "001" {
		t.Fatalf("unexpected snapshot: %+v", snap)
	}
}

func writeFixtureCSV(t *testing.T, root, name, content string) string {
	t.Helper()
	path := filepath.Join(root, name)
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}
