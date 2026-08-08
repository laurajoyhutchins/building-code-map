package regulatory

import (
	"strings"
	"testing"
)

func TestGovernmentalEntityCandidateHasStableSourceIdentityAndStartsUnresolved(t *testing.T) {
	input := GovernmentalEntityCandidateInput{
		OfficialName: "City and County of Denver",
		EntityType:   EntityTypeConsolidatedGovernment,
		StateID:      "US-CO",
		StateFIPS:    "08",
		Identity: CanonicalEntityIdentity{
			Namespace: "census-place",
			Value:     "0804000",
		},
		LegalStatus: LegalStatusActive,
		Geographies: []EntityGeographyReference{
			{Kind: "census_place", Identifier: "0804000", Vintage: "2025"},
		},
		SourceIDs: []string{"src:census-places-2025"},
	}

	candidate, err := NewGovernmentalEntityCandidate(input)
	if err != nil {
		t.Fatal(err)
	}
	if candidate.EntityID != "gov:us-co:consolidated_government:census-place:0804000" {
		t.Fatalf("entity id=%q", candidate.EntityID)
	}
	if candidate.Classification != ClassificationUnresolved {
		t.Fatalf("new geographic candidate classification=%q, want unresolved", candidate.Classification)
	}
	if candidate.HistoricalGeographyStatus != HistoricalGeographyCurrentOnly {
		t.Fatalf("historical geography status=%q", candidate.HistoricalGeographyStatus)
	}
	if len(candidate.Geographies) != 1 || candidate.Geographies[0].Identifier != "0804000" {
		t.Fatalf("geographies=%#v", candidate.Geographies)
	}

	renamed := input
	renamed.OfficialName = "Denver Consolidated Government"
	renamedCandidate, err := NewGovernmentalEntityCandidate(renamed)
	if err != nil {
		t.Fatal(err)
	}
	if renamedCandidate.EntityID != candidate.EntityID {
		t.Fatalf("display-name change altered identity: %q != %q", renamedCandidate.EntityID, candidate.EntityID)
	}
}

func TestCanonicalGovernmentalEntityIDRejectsGeographyWithoutLegalEntityIdentity(t *testing.T) {
	_, err := CanonicalGovernmentalEntityID("US-CO", EntityTypeMunicipality, CanonicalEntityIdentity{
		Namespace: "",
		Value:     "0804000",
	})
	if err == nil || !strings.Contains(err.Error(), "identity namespace") {
		t.Fatalf("error=%v", err)
	}
}

func TestValidateEntityInventoryRequiresExplicitClassificationAndSourceEvidence(t *testing.T) {
	entity, err := NewGovernmentalEntityCandidate(GovernmentalEntityCandidateInput{
		OfficialName: "Example City",
		EntityType:   EntityTypeMunicipality,
		StateID:      "US-CO",
		StateFIPS:    "08",
		Identity: CanonicalEntityIdentity{
			Namespace: "census-place",
			Value:     "0812345",
		},
		LegalStatus: LegalStatusActive,
		SourceIDs:   []string{"src:census-places-2025"},
	})
	if err != nil {
		t.Fatal(err)
	}
	inventory := EntityInventory{
		SchemaVersion: EntityInventorySchemaVersion,
		InventoryID:   "governmental-entities-2025",
		GeneratedAt:   "2026-08-08T15:00:00Z",
		Sources: []Source{
			{
				ID:         "src:census-places-2025",
				Title:      "Census place inventory",
				URL:        "https://example.invalid/census-places",
				Kind:       "governmental_inventory",
				AccessedAt: "2026-08-08",
			},
		},
		Entities: []GovernmentalEntity{entity},
	}
	if err := ValidateEntityInventory(inventory); err != nil {
		t.Fatalf("valid inventory: %v", err)
	}

	inventory.Entities[0].Classification = ""
	if err := ValidateEntityInventory(inventory); err == nil || !strings.Contains(err.Error(), "classification") {
		t.Fatalf("missing classification error=%v", err)
	}

	inventory.Entities[0] = entity
	inventory.Entities[0].SourceIDs = []string{"src:missing"}
	if err := ValidateEntityInventory(inventory); err == nil || !strings.Contains(err.Error(), "unknown source") {
		t.Fatalf("missing source evidence error=%v", err)
	}
}

func TestValidateEntityInventoryRejectsDuplicateCanonicalEntities(t *testing.T) {
	entity, err := NewGovernmentalEntityCandidate(GovernmentalEntityCandidateInput{
		OfficialName: "Example County",
		EntityType:   EntityTypeCountyEquivalent,
		StateID:      "US-VA",
		StateFIPS:    "51",
		Identity: CanonicalEntityIdentity{
			Namespace: "census-county",
			Value:     "51059",
		},
		LegalStatus: LegalStatusActive,
		SourceIDs:   []string{"src:fixture"},
	})
	if err != nil {
		t.Fatal(err)
	}
	inventory := EntityInventory{
		SchemaVersion: EntityInventorySchemaVersion,
		InventoryID:   "duplicate-fixture",
		GeneratedAt:   "2026-08-08T15:00:00Z",
		Sources: []Source{
			{ID: "src:fixture", Title: "Fixture", URL: "https://example.invalid", Kind: "fixture", AccessedAt: "2026-08-08"},
		},
		Entities: []GovernmentalEntity{entity, entity},
	}
	if err := ValidateEntityInventory(inventory); err == nil || !strings.Contains(err.Error(), "duplicate entity_id") {
		t.Fatalf("duplicate error=%v", err)
	}
}

func TestHistoricalGeographyCanBeExplicitlyUnavailable(t *testing.T) {
	entity, err := NewGovernmentalEntityCandidate(GovernmentalEntityCandidateInput{
		OfficialName: "Example Special District",
		EntityType:   EntityTypeSpecialDistrict,
		StateID:      "US-OR",
		StateFIPS:    "41",
		Identity: CanonicalEntityIdentity{
			Namespace: "state-registry",
			Value:     "district-42",
		},
		LegalStatus:               LegalStatusActive,
		HistoricalGeographyStatus: HistoricalGeographyUnavailable,
		SourceIDs:                 []string{"src:fixture"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if entity.HistoricalGeographyStatus != HistoricalGeographyUnavailable {
		t.Fatalf("historical geography status=%q", entity.HistoricalGeographyStatus)
	}
	if len(entity.Geographies) != 0 {
		t.Fatalf("unavailable historical geography invented records: %#v", entity.Geographies)
	}
}
