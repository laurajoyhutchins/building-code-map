package regulatory

import "testing"

func TestPilotClassificationColoradoSeparatesBuildingAdoptionFromElectricalEnforcement(t *testing.T) {
	inventories, regimes := loadPilotClassificationData(t)
	entity := firstActiveEntityOfType(t, inventories.Inventories["US-CO"], EntityTypeMunicipality)
	regime := regimes["US-CO"]

	building, err := ClassifyGovernmentalEntity(EntityClassificationRequest{
		Entity: entity,
		Regime: regime,
		CodeFamily: "building",
		ApplicabilityDate: "2026-08-08",
	})
	if err != nil {
		t.Fatal(err)
	}
	if building.Classification != ClassificationLocalAdopter {
		t.Fatalf("Colorado building classification = %q, want %q", building.Classification, ClassificationLocalAdopter)
	}

	electrical, err := ClassifyGovernmentalEntity(EntityClassificationRequest{
		Entity: entity,
		Regime: regime,
		CodeFamily: "electrical",
		ApplicabilityDate: "2026-08-08",
	})
	if err != nil {
		t.Fatal(err)
	}
	if electrical.Classification != ClassificationEnforcementOnly {
		t.Fatalf("Colorado electrical classification = %q, want %q", electrical.Classification, ClassificationEnforcementOnly)
	}
	if !hasClassificationRecordKind(electrical.Records, ClassificationRecordStateInheritance) {
		t.Fatalf("Colorado electrical records = %#v, want state inheritance", electrical.Records)
	}
}

func TestPilotClassificationStatewideCodeStatesRetainLocalEnforcementAndInheritance(t *testing.T) {
	inventories, regimes := loadPilotClassificationData(t)
	cases := []struct {
		stateID string
		family string
		date string
	}{
		{stateID: "US-FL", family: "building", date: "2026-08-08"},
		{stateID: "US-NC", family: "building", date: "2026-08-08"},
		{stateID: "US-NJ", family: "building", date: "2026-08-08"},
		{stateID: "US-OR", family: "building", date: "2026-08-08"},
		{stateID: "US-VA", family: "building", date: "2026-08-08"},
	}
	for _, tc := range cases {
		t.Run(tc.stateID, func(t *testing.T) {
			entity := firstActiveEntityOfType(t, inventories.Inventories[tc.stateID], EntityTypeMunicipality)
			decision, err := ClassifyGovernmentalEntity(EntityClassificationRequest{
				Entity: entity,
				Regime: regimes[tc.stateID],
				CodeFamily: tc.family,
				ApplicabilityDate: tc.date,
			})
			if err != nil {
				t.Fatal(err)
			}
			if decision.Classification != ClassificationEnforcementOnly {
				t.Fatalf("classification = %q, want %q; decision=%#v", decision.Classification, ClassificationEnforcementOnly, decision)
			}
			if !hasClassificationRecordKind(decision.Records, ClassificationRecordStateInheritance) {
				t.Fatalf("records = %#v, want state inheritance", decision.Records)
			}
		})
	}
}

func TestPilotClassificationIndependentCityUsesVirginiaMunicipalEnforcement(t *testing.T) {
	inventories, regimes := loadPilotClassificationData(t)
	entity := firstActiveEntityOfType(t, inventories.Inventories["US-VA"], EntityTypeIndependentCity)
	decision, err := ClassifyGovernmentalEntity(EntityClassificationRequest{
		Entity: entity,
		Regime: regimes["US-VA"],
		CodeFamily: "building",
		ApplicabilityDate: "2026-08-08",
	})
	if err != nil {
		t.Fatal(err)
	}
	if decision.Classification != ClassificationEnforcementOnly {
		t.Fatalf("classification = %q, want %q", decision.Classification, ClassificationEnforcementOnly)
	}
	if !hasClassificationRecordKind(decision.Records, ClassificationRecordStateInheritance) {
		t.Fatalf("records = %#v, want state inheritance", decision.Records)
	}
}

func loadPilotClassificationData(t *testing.T) (CensusEntityInventoryBuild, map[string]StateRegimeSpec) {
	t.Helper()
	inventories, err := LoadCensusEntityInventoryDirectory("../data/governmental-entities/2025")
	if err != nil {
		t.Fatalf("LoadCensusEntityInventoryDirectory() error = %v", err)
	}
	specs, err := LoadStateRegimeDirectory("../data/state-regimes")
	if err != nil {
		t.Fatalf("LoadStateRegimeDirectory() error = %v", err)
	}
	regimes := make(map[string]StateRegimeSpec, len(specs))
	for _, spec := range specs {
		regimes[spec.StateID] = spec
	}
	return inventories, regimes
}

func firstActiveEntityOfType(t *testing.T, inventory EntityInventory, entityType GovernmentalEntityType) GovernmentalEntity {
	t.Helper()
	for _, entity := range inventory.Entities {
		if entity.EntityType == entityType && entity.LegalStatus == LegalStatusActive {
			return entity
		}
	}
	t.Fatalf("no active %s entity found in %s", entityType, inventory.InventoryID)
	return GovernmentalEntity{}
}
