package regulatory

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestWriteAndLoadCensusEntityInventoryDirectory(t *testing.T) {
	source := Source{
		ID:         "src:census-government-units-2025",
		Title:      "2025 Government Units Listing",
		URL:        "https://example.invalid/census",
		Kind:       "governmental_inventory",
		AccessedAt: "2026-08-08",
	}
	csv := strings.Join([]string{
		strings.Join(CensusGeneralPurposeHeader, ","),
		"100001,COUNTY OF AUTAUGA,1 - COUNTY,,,,,AL,,,,COUNTY,60342,2023,01,001,99001,,Y",
	}, "\n") + "\n"
	build, err := BuildCensusEntityInventories(strings.NewReader(csv), source, time.Date(2026, 8, 8, 16, 30, 0, 0, time.UTC))
	if err != nil {
		t.Fatal(err)
	}
	root := filepath.Join(t.TempDir(), "inventories")
	if err := WriteCensusEntityInventoryDirectory(root, build); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(root, "US-AL.json")); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(root, "index.json")); err != nil {
		t.Fatal(err)
	}
	loaded, err := LoadCensusEntityInventoryDirectory(root)
	if err != nil {
		t.Fatal(err)
	}
	if loaded.Index.TotalEntities != 1 || len(loaded.Inventories["US-AL"].Entities) != 1 {
		t.Fatalf("loaded=%#v", loaded)
	}
}

func TestLoadCensusEntityInventoryDirectoryRejectsIndexDrift(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "index.json"), []byte(`{"schema_version":"1.0","source_id":"src:census-government-units-2025","generated_at":"2026-08-08T16:30:00Z","total_entities":2,"entity_type_counts":{},"classification_counts":{},"states":[]}`+"\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadCensusEntityInventoryDirectory(root); err == nil || !strings.Contains(err.Error(), "state summaries") {
		t.Fatalf("index drift error=%v", err)
	}
}
