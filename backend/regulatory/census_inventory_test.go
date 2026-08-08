package regulatory

import (
	"strings"
	"testing"
	"time"
)

func TestBuildCensusEntityInventoriesGroupsSortsAndCountsStates(t *testing.T) {
	csv := strings.Join([]string{
		strings.Join(CensusGeneralPurposeHeader, ","),
		"200002,CITY OF BETA,2 - MUNICIPAL,,,,,CO,,,,CITY,200,2023,08,001,20000,,Y",
		"100001,COUNTY OF ALPHA,1 - COUNTY,,,,,AL,,,,COUNTY,100,2023,01,001,99001,,Y",
		"200001,TOWN OF ALPHA,2 - MUNICIPAL,,,,,CO,,,,TOWN,50,2023,08,001,10000,,N",
	}, "\n") + "\n"
	source := Source{
		ID:         "src:census-government-units-2025",
		Title:      "2025 Government Units Listing",
		URL:        "https://example.invalid/census",
		Kind:       "governmental_inventory",
		AccessedAt: "2026-08-08",
	}

	result, err := BuildCensusEntityInventories(
		strings.NewReader(csv),
		source,
		time.Date(2026, 8, 8, 16, 30, 0, 0, time.UTC),
	)
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Inventories) != 2 {
		t.Fatalf("inventories=%d", len(result.Inventories))
	}
	colorado := result.Inventories["US-CO"]
	if len(colorado.Entities) != 2 {
		t.Fatalf("Colorado entities=%d", len(colorado.Entities))
	}
	if colorado.Entities[0].EntityID >= colorado.Entities[1].EntityID {
		t.Fatalf("entities are not sorted: %q >= %q", colorado.Entities[0].EntityID, colorado.Entities[1].EntityID)
	}
	if result.Index.TotalEntities != 3 {
		t.Fatalf("total=%d", result.Index.TotalEntities)
	}
	if result.Index.ClassificationCounts[ClassificationUnresolved] != 2 {
		t.Fatalf("unresolved=%d", result.Index.ClassificationCounts[ClassificationUnresolved])
	}
	if result.Index.ClassificationCounts[ClassificationInactive] != 1 {
		t.Fatalf("inactive=%d", result.Index.ClassificationCounts[ClassificationInactive])
	}
	if result.Index.States[0].StateID != "US-AL" || result.Index.States[1].StateID != "US-CO" {
		t.Fatalf("state summaries=%#v", result.Index.States)
	}
}

func TestBuildCensusEntityInventoriesRejectsHeaderDriftAndDuplicateGovernmentUnits(t *testing.T) {
	source := Source{
		ID:         "src:census-government-units-2025",
		Title:      "2025 Government Units Listing",
		URL:        "https://example.invalid/census",
		Kind:       "governmental_inventory",
		AccessedAt: "2026-08-08",
	}
	generatedAt := time.Date(2026, 8, 8, 16, 30, 0, 0, time.UTC)

	badHeader := strings.Replace(strings.Join(CensusGeneralPurposeHeader, ","), "UNIT_TYPE", "RENAMED_UNIT_TYPE", 1) + "\n"
	if _, err := BuildCensusEntityInventories(strings.NewReader(badHeader), source, generatedAt); err == nil || !strings.Contains(err.Error(), "header") {
		t.Fatalf("header drift error=%v", err)
	}

	row := "100001,COUNTY OF AUTAUGA,1 - COUNTY,,,,,AL,,,,COUNTY,60342,2023,01,001,99001,,Y"
	duplicate := strings.Join([]string{
		strings.Join(CensusGeneralPurposeHeader, ","),
		row,
		row,
	}, "\n") + "\n"
	if _, err := BuildCensusEntityInventories(strings.NewReader(duplicate), source, generatedAt); err == nil || !strings.Contains(err.Error(), "duplicate entity_id") {
		t.Fatalf("duplicate error=%v", err)
	}
}

func TestValidateNationalCensusInventoryRequiresAllStatesAndExpectedDenominator(t *testing.T) {
	result := CensusEntityInventoryBuild{
		Inventories: map[string]EntityInventory{
			"US-DC": {
				SchemaVersion: EntityInventorySchemaVersion,
				InventoryID:   "census-government-units-2025:us-dc",
				GeneratedAt:   "2026-08-08T16:30:00Z",
				Sources: []Source{{
					ID: "src:census-government-units-2025", Title: "Fixture", URL: "https://example.invalid", Kind: "governmental_inventory", AccessedAt: "2026-08-08",
				}},
				Entities: []GovernmentalEntity{},
			},
		},
		Index: EntityInventoryIndex{TotalEntities: 1},
	}
	if err := ValidateNationalCensusInventory(result, 38704); err == nil || !strings.Contains(err.Error(), "51") {
		t.Fatalf("missing states error=%v", err)
	}
}
