package regulatory

import "testing"

func TestBuildEntityCoverageLedgerSeparatesClassificationFromVerifiedApplicability(t *testing.T) {
	entities := []GovernmentalEntity{
		testCoverageEntity(t, "001", EntityTypeMunicipality, 100),
		testCoverageEntity(t, "002", EntityTypeCountyEquivalent, 200),
	}
	inventory := testCoverageInventory("US-TS", entities)
	regime := testCoverageRegime()
	regime.CodeFamilyPolicies = map[string]ResolutionPolicy{
		"building": {
			Status: "partially_resolved",
			AdoptionIDs: []string{"adoption:state:building"},
			SourceIDs: []string{"src:state-building"},
		},
	}
	regime.DateRules = []DateRule{{
		ID: "rule:state-building-window",
		CodeFamily: "building",
		RuleType: "effective_window",
		Trigger: "applicability_date",
		StartDate: "2025-01-01",
		Summary: "Current state building code window.",
		SourceIDs: []string{"src:state-building-date"},
		Verification: Verification{Status: "verified"},
	}}
	regime.EntityPowerEvidence = []EntityPowerEvidence{
		{
			EntityKind: "municipality",
			Scope: "incorporated",
			Powers: []RegimePower{RegimePowerAdminister, RegimePowerEnforce},
			SourceIDs: []string{"src:municipal-enforcement"},
		},
	}

	ledger, err := BuildEntityCoverageLedger(EntityCoverageLedgerRequest{
		Inventories: testCoverageBuild(inventory),
		Regimes: []StateRegimeSpec{regime},
		StateIDs: []string{"US-TS"},
		CodeFamily: "building",
		ApplicabilityDate: "2026-08-08",
	})
	if err != nil {
		t.Fatalf("BuildEntityCoverageLedger() error = %v", err)
	}

	summary := ledger.States[0].Summary
	if summary.ExpectedEntities != 2 || summary.ActiveEntities != 2 || summary.ClassifiedEntities != 2 {
		t.Fatalf("summary counts = %#v", summary)
	}
	if summary.InheritedEntities != 2 || summary.VerifiedEntities != 2 {
		t.Fatalf("inheritance/verified counts = %#v", summary)
	}
	if summary.RequiresLocalEvidenceEntities != 1 {
		t.Fatalf("requires-local-evidence = %d, want 1", summary.RequiresLocalEvidenceEntities)
	}
	if summary.UnresolvedEntities != 0 || summary.InactiveEntities != 0 || summary.ConflictingEntities != 0 {
		t.Fatalf("unexpected unresolved/inactive/conflicting counts = %#v", summary)
	}
	if summary.Population.Total != 300 || summary.Population.Verified != 300 || summary.Population.Unresolved != 0 {
		t.Fatalf("population summary = %#v", summary.Population)
	}

	municipality := ledger.States[0].Entities[1]
	if municipality.Classification != ClassificationEnforcementOnly || municipality.EvidenceStatus != CoverageEvidenceVerifiedInheritanceWithLocalEvidence {
		t.Fatalf("municipal coverage = %#v", municipality)
	}
	if !municipality.Inherited || !municipality.RequiresLocalEvidence {
		t.Fatalf("municipal evidence flags = %#v", municipality)
	}
}

func TestBuildEntityCoverageLedgerDoesNotCallLocalAdopterVerified(t *testing.T) {
	entity := testCoverageEntity(t, "001", EntityTypeMunicipality, 125)
	regime := testCoverageRegime()
	regime.EntityPowerEvidence = []EntityPowerEvidence{{
		EntityKind: "municipality",
		Scope: "incorporated",
		Powers: []RegimePower{RegimePowerAdopt, RegimePowerEnforce},
		SourceIDs: []string{"src:local-adoption-authority"},
	}}

	ledger, err := BuildEntityCoverageLedger(EntityCoverageLedgerRequest{
		Inventories: testCoverageBuild(testCoverageInventory("US-TS", []GovernmentalEntity{entity})),
		Regimes: []StateRegimeSpec{regime},
		StateIDs: []string{"US-TS"},
		CodeFamily: "building",
		ApplicabilityDate: "2026-08-08",
	})
	if err != nil {
		t.Fatal(err)
	}
	entry := ledger.States[0].Entities[0]
	if entry.Classification != ClassificationLocalAdopter {
		t.Fatalf("classification = %q", entry.Classification)
	}
	if entry.EvidenceStatus != CoverageEvidenceClassifiedRequiresLocalEvidence {
		t.Fatalf("evidence status = %q", entry.EvidenceStatus)
	}
	if ledger.States[0].Summary.VerifiedEntities != 0 || ledger.States[0].Summary.RequiresLocalEvidenceEntities != 1 {
		t.Fatalf("summary = %#v", ledger.States[0].Summary)
	}
}

func TestBuildEntityCoverageLedgerKeepsUnresolvedAndInactiveDistinct(t *testing.T) {
	active := testCoverageEntity(t, "001", EntityTypeMunicipality, 50)
	inactive := testCoverageEntity(t, "002", EntityTypeMunicipality, 75)
	inactive.LegalStatus = LegalStatusInactive
	inactive.Classification = ClassificationInactive

	ledger, err := BuildEntityCoverageLedger(EntityCoverageLedgerRequest{
		Inventories: testCoverageBuild(testCoverageInventory("US-TS", []GovernmentalEntity{active, inactive})),
		Regimes: []StateRegimeSpec{testCoverageRegime()},
		StateIDs: []string{"US-TS"},
		CodeFamily: "building",
		ApplicabilityDate: "2026-08-08",
	})
	if err != nil {
		t.Fatal(err)
	}
	summary := ledger.States[0].Summary
	if summary.UnresolvedEntities != 1 || summary.InactiveEntities != 1 || summary.ClassifiedEntities != 1 {
		t.Fatalf("summary = %#v", summary)
	}
	if summary.Population.Unresolved != 50 {
		t.Fatalf("unresolved population = %d, want 50", summary.Population.Unresolved)
	}
}

func TestBuildEntityCoverageLedgerSurfacesConflictingPolicyWithoutInventingCoverage(t *testing.T) {
	entity := testCoverageEntity(t, "001", EntityTypeMunicipality, 90)
	regime := testCoverageRegime()
	regime.Territory.Incorporated = ResolutionPolicy{
		Status: "conflicting",
		SourceIDs: []string{"src:conflict-a", "src:conflict-b"},
	}

	ledger, err := BuildEntityCoverageLedger(EntityCoverageLedgerRequest{
		Inventories: testCoverageBuild(testCoverageInventory("US-TS", []GovernmentalEntity{entity})),
		Regimes: []StateRegimeSpec{regime},
		StateIDs: []string{"US-TS"},
		CodeFamily: "building",
		ApplicabilityDate: "2026-08-08",
	})
	if err != nil {
		t.Fatal(err)
	}
	entry := ledger.States[0].Entities[0]
	if entry.Classification != ClassificationUnresolved || entry.EvidenceStatus != CoverageEvidenceConflicting {
		t.Fatalf("entry = %#v", entry)
	}
	if ledger.States[0].Summary.ConflictingEntities != 1 {
		t.Fatalf("summary = %#v", ledger.States[0].Summary)
	}
}

func TestBuildEntityCoverageLedgerRejectsStateOutsideRegimeCorpus(t *testing.T) {
	_, err := BuildEntityCoverageLedger(EntityCoverageLedgerRequest{
		Inventories: testCoverageBuild(testCoverageInventory("US-TS", []GovernmentalEntity{testCoverageEntity(t, "001", EntityTypeMunicipality, 1)})),
		Regimes: nil,
		StateIDs: []string{"US-TS"},
		CodeFamily: "building",
		ApplicabilityDate: "2026-08-08",
	})
	if err == nil {
		t.Fatal("expected missing regime error")
	}
}

func testCoverageEntity(t *testing.T, identity string, entityType GovernmentalEntityType, population int64) GovernmentalEntity {
	t.Helper()
	entity, err := NewGovernmentalEntityCandidate(GovernmentalEntityCandidateInput{
		OfficialName: "Test Government " + identity,
		EntityType: entityType,
		StateID: "US-TS",
		StateFIPS: "99",
		Identity: CanonicalEntityIdentity{Namespace: "test", Value: identity},
		LegalStatus: LegalStatusActive,
		HistoricalGeographyStatus: HistoricalGeographyUnavailable,
		Population: &EntityPopulation{Count: population, SourceYear: "2023"},
		SourceIDs: []string{"src:entity"},
	})
	if err != nil {
		t.Fatal(err)
	}
	return entity
}

func testCoverageInventory(stateID string, entities []GovernmentalEntity) EntityInventory {
	return EntityInventory{
		SchemaVersion: EntityInventorySchemaVersion,
		InventoryID: "test-inventory:" + stateID,
		GeneratedAt: "2026-08-08T00:00:00Z",
		Sources: []EntitySource{{ID: "src:test-inventory", Title: "Test inventory", Kind: "governmental_inventory", AccessedAt: "2026-08-08"}},
		Entities: entities,
	}
}

func testCoverageBuild(inventory EntityInventory) CensusEntityInventoryBuild {
	inventories := map[string]EntityInventory{inventory.Entities[0].StateID: inventory}
	return CensusEntityInventoryBuild{
		Inventories: inventories,
		Index: indexFromInventories(inventories, "src:test-inventory", "2026-08-08T00:00:00Z"),
	}
}

func testCoverageRegime() StateRegimeSpec {
	regime := testClassificationRegime()
	regime.Verification = Verification{Status: "verified", Confidence: 0.95}
	return regime
}
