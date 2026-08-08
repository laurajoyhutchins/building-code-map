package regulatory

import "testing"

func TestClassifyGovernmentalEntityInactiveWinsWithoutInference(t *testing.T) {
	entity := testClassificationEntity(t, EntityTypeMunicipality)
	entity.LegalStatus = LegalStatusInactive
	entity.Classification = ClassificationInactive

	decision, err := ClassifyGovernmentalEntity(EntityClassificationRequest{
		Entity: entity,
		Regime: testClassificationRegime(),
		CodeFamily: "building",
	})
	if err != nil {
		t.Fatalf("ClassifyGovernmentalEntity() error = %v", err)
	}
	if decision.Classification != ClassificationInactive {
		t.Fatalf("classification = %q, want %q", decision.Classification, ClassificationInactive)
	}
	if len(decision.Records) != 1 || decision.Records[0].Kind != ClassificationRecordLegalStatus {
		t.Fatalf("records = %#v, want one legal-status record", decision.Records)
	}
}

func TestClassifyGovernmentalEntityUsesExplicitLocalAdoptionPower(t *testing.T) {
	regime := testClassificationRegime()
	regime.EntityPowerEvidence = []EntityPowerEvidence{{
		EntityKind: "municipality",
		Scope: "incorporated",
		Powers: []RegimePower{RegimePowerAdopt, RegimePowerEnforce},
		SourceIDs: []string{"src:local-adoption"},
	}}

	decision, err := ClassifyGovernmentalEntity(EntityClassificationRequest{
		Entity: testClassificationEntity(t, EntityTypeMunicipality),
		Regime: regime,
		CodeFamily: "building",
	})
	if err != nil {
		t.Fatalf("ClassifyGovernmentalEntity() error = %v", err)
	}
	if decision.Classification != ClassificationLocalAdopter {
		t.Fatalf("classification = %q, want %q", decision.Classification, ClassificationLocalAdopter)
	}
	if !containsString(decision.SourceIDs, "src:local-adoption") {
		t.Fatalf("source_ids = %#v, want local adoption evidence", decision.SourceIDs)
	}
}

func TestClassifyGovernmentalEntityUsesLocalAmendmentBeforeEnforcementOnly(t *testing.T) {
	regime := testClassificationRegime()
	regime.EntityPowerEvidence = []EntityPowerEvidence{{
		EntityKind: "municipality",
		Scope: "incorporated",
		Powers: []RegimePower{RegimePowerAdminister, RegimePowerEnforce},
		SourceIDs: []string{"src:local-enforcement"},
	}}
	regime.AmendmentRules = []AmendmentRule{{
		ID: "rule:local-amendment",
		CodeFamily: "building",
		Level: "local",
		Posture: "local_amendment_permitted",
		Summary: "Local technical amendments are evidence-backed.",
		SourceIDs: []string{"src:local-amendment"},
		Verification: Verification{Status: "verified"},
	}}

	decision, err := ClassifyGovernmentalEntity(EntityClassificationRequest{
		Entity: testClassificationEntity(t, EntityTypeMunicipality),
		Regime: regime,
		CodeFamily: "building",
	})
	if err != nil {
		t.Fatalf("ClassifyGovernmentalEntity() error = %v", err)
	}
	if decision.Classification != ClassificationLocalAmender {
		t.Fatalf("classification = %q, want %q", decision.Classification, ClassificationLocalAmender)
	}
}

func TestClassifyGovernmentalEntityBuildsStateInheritanceRecordWithInterval(t *testing.T) {
	regime := testClassificationRegime()
	regime.CodeFamilyPolicies = map[string]ResolutionPolicy{
		"building": {
			Status: "partially_resolved",
			AdoptionIDs: []string{"adoption:state:building:2024"},
			SourceIDs: []string{"src:state-adoption"},
		},
	}
	regime.DateRules = []DateRule{{
		ID: "rule:building-window",
		CodeFamily: "building",
		RuleType: "effective_window",
		Trigger: "applicability_date",
		StartDate: "2024-01-01",
		EndDate: "2027-12-31",
		Summary: "State building code applicability window.",
		SourceIDs: []string{"src:state-date"},
		Verification: Verification{Status: "verified"},
	}}

	decision, err := ClassifyGovernmentalEntity(EntityClassificationRequest{
		Entity: testClassificationEntity(t, EntityTypeMunicipality),
		Regime: regime,
		CodeFamily: "building",
		ApplicabilityDate: "2026-08-08",
	})
	if err != nil {
		t.Fatalf("ClassifyGovernmentalEntity() error = %v", err)
	}
	if decision.Classification != ClassificationCoveredByState {
		t.Fatalf("classification = %q, want %q", decision.Classification, ClassificationCoveredByState)
	}
	if len(decision.Records) != 1 {
		t.Fatalf("records = %#v, want one inheritance record", decision.Records)
	}
	record := decision.Records[0]
	if record.Kind != ClassificationRecordStateInheritance {
		t.Fatalf("record kind = %q, want %q", record.Kind, ClassificationRecordStateInheritance)
	}
	if record.EffectiveInterval.StartDate != "2024-01-01" || record.EffectiveInterval.EndDate != "2027-12-31" {
		t.Fatalf("effective interval = %#v", record.EffectiveInterval)
	}
	if !containsString(record.LegalBasisSourceIDs, "src:state-adoption") || !containsString(record.LegalBasisSourceIDs, "src:state-date") {
		t.Fatalf("legal basis = %#v, want adoption and date sources", record.LegalBasisSourceIDs)
	}
}

func TestClassifyGovernmentalEntityDoesNotInferStateInheritanceWithoutTemporalBasis(t *testing.T) {
	regime := testClassificationRegime()
	regime.CodeFamilyPolicies = map[string]ResolutionPolicy{
		"building": {
			Status: "partially_resolved",
			AdoptionIDs: []string{"adoption:state:building:2024"},
			SourceIDs: []string{"src:state-adoption"},
		},
	}

	decision, err := ClassifyGovernmentalEntity(EntityClassificationRequest{
		Entity: testClassificationEntity(t, EntityTypeMunicipality),
		Regime: regime,
		CodeFamily: "building",
	})
	if err != nil {
		t.Fatalf("ClassifyGovernmentalEntity() error = %v", err)
	}
	if decision.Classification != ClassificationUnresolved {
		t.Fatalf("classification = %q, want %q", decision.Classification, ClassificationUnresolved)
	}
	if !containsString(decision.UnresolvedReasons, "state inheritance lacks an applicable temporal rule") {
		t.Fatalf("unresolved_reasons = %#v", decision.UnresolvedReasons)
	}
}

func TestClassifyGovernmentalEntityCodeFamilyOverrideDoesNotLeakGenericLocalAdoption(t *testing.T) {
	regime := testClassificationRegime()
	regime.EntityPowerEvidence = []EntityPowerEvidence{{
		EntityKind: "municipality",
		Scope: "incorporated",
		Powers: []RegimePower{RegimePowerAdopt, RegimePowerEnforce},
		SourceIDs: []string{"src:generic-local"},
	}}
	regime.CodeFamilyPolicies = map[string]ResolutionPolicy{
		"electrical": {
			Status: "partially_resolved",
			AdoptionIDs: []string{"adoption:state:electrical:2026"},
			SourceIDs: []string{"src:state-electrical"},
		},
	}
	regime.DateRules = []DateRule{{
		ID: "rule:electrical-window",
		CodeFamily: "electrical",
		RuleType: "effective_window",
		Trigger: "applicability_date",
		StartDate: "2026-08-01",
		Summary: "State electrical code window.",
		SourceIDs: []string{"src:state-electrical-date"},
		Verification: Verification{Status: "verified"},
	}}
	regime.EnforcementRules = []EnforcementRule{{
		ID: "rule:electrical-local-enforcement",
		CodeFamily: "electrical",
		Model: "hybrid",
		EntityKinds: []string{"municipal_electrical_authority"},
		Summary: "Approved municipal electrical authorities may administer and enforce the state code.",
		SourceIDs: []string{"src:local-electrical-enforcement"},
		Verification: Verification{Status: "verified"},
	}}

	decision, err := ClassifyGovernmentalEntity(EntityClassificationRequest{
		Entity: testClassificationEntity(t, EntityTypeMunicipality),
		Regime: regime,
		CodeFamily: "electrical",
		ApplicabilityDate: "2026-08-08",
	})
	if err != nil {
		t.Fatalf("ClassifyGovernmentalEntity() error = %v", err)
	}
	if decision.Classification != ClassificationEnforcementOnly {
		t.Fatalf("classification = %q, want %q", decision.Classification, ClassificationEnforcementOnly)
	}
}

func TestClassifyGovernmentalEntityIndependentCityUsesMunicipalEvidence(t *testing.T) {
	regime := testClassificationRegime()
	regime.EntityPowerEvidence = []EntityPowerEvidence{{
		EntityKind: "municipality",
		Scope: "incorporated",
		Powers: []RegimePower{RegimePowerAdopt},
		SourceIDs: []string{"src:city-power"},
	}}

	decision, err := ClassifyGovernmentalEntity(EntityClassificationRequest{
		Entity: testClassificationEntity(t, EntityTypeIndependentCity),
		Regime: regime,
		CodeFamily: "building",
	})
	if err != nil {
		t.Fatalf("ClassifyGovernmentalEntity() error = %v", err)
	}
	if decision.Classification != ClassificationLocalAdopter {
		t.Fatalf("classification = %q, want %q", decision.Classification, ClassificationLocalAdopter)
	}
}

func testClassificationEntity(t *testing.T, entityType GovernmentalEntityType) GovernmentalEntity {
	t.Helper()
	entity, err := NewGovernmentalEntityCandidate(GovernmentalEntityCandidateInput{
		OfficialName: "Test Government",
		EntityType: entityType,
		StateID: "US-TS",
		StateFIPS: "99",
		Identity: CanonicalEntityIdentity{Namespace: "test", Value: "000001"},
		LegalStatus: LegalStatusActive,
		HistoricalGeographyStatus: HistoricalGeographyUnavailable,
		SourceIDs: []string{"src:entity"},
	})
	if err != nil {
		t.Fatalf("NewGovernmentalEntityCandidate() error = %v", err)
	}
	return entity
}

func testClassificationRegime() StateRegimeSpec {
	return StateRegimeSpec{
		SchemaVersion: StateRegimeSchemaVersion,
		RegimeID: "state-regime:us-ts",
		StateID: "US-TS",
		StateName: "Test State",
		StateAbbreviation: "TS",
		StateFIPS: "99",
		SourceProfileID: "state-profile:us-ts",
		SourceProfileStatus: "verified",
		LastVerified: "2026-08-08",
		Dimensions: StateRegimeDimensions{
			LocalAdoption: RegimeDimensionEvidence{Status: RegimeEvidenceUnresolved},
			LocalAmendment: RegimeDimensionEvidence{Status: RegimeEvidenceUnresolved},
			Enforcement: RegimeDimensionEvidence{Status: RegimeEvidenceUnresolved},
			DelegationContracting: RegimeDimensionEvidence{Status: RegimeEvidenceUnresolved},
			HomeRule: RegimeDimensionEvidence{Status: RegimeEvidenceUnresolved},
			StatewideInheritance: RegimeDimensionEvidence{Status: RegimeEvidenceUnresolved},
			Unincorporated: RegimeDimensionEvidence{Status: RegimeEvidenceUnresolved},
			Temporal: RegimeDimensionEvidence{Status: RegimeEvidenceUnresolved},
		},
		Territory: StateRegimeTerritory{
			Incorporated: ResolutionPolicy{Status: "insufficient_evidence", SourceIDs: []string{"src:regime"}},
			Unincorporated: ResolutionPolicy{Status: "insufficient_evidence", SourceIDs: []string{"src:regime"}},
		},
		SourceIDs: []string{"src:regime"},
		Verification: Verification{Status: "verified"},
	}
}

func containsString(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}
