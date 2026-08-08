package regulatory

import (
	"reflect"
	"testing"
)

func TestBuildStateRegimeSpecPreservesTerritoryEvidenceAndUnresolvedDimensions(t *testing.T) {
	profile := StateProfile{
		SchemaVersion:     ProfileSchemaVersion,
		ProfileID:         "state-profile:us-co",
		StateID:           "US-CO",
		StateName:         "Colorado",
		StateAbbreviation: "CO",
		StateFIPS:         "08",
		Status:            "verified",
		LastVerified:      "2026-08-02",
		Sources: []Source{
			{ID: "src:local", Title: "Local authority source", URL: "https://example.invalid/local", Kind: "statute", AccessedAt: "2026-08-02"},
			{ID: "src:electrical", Title: "Electrical source", URL: "https://example.invalid/electrical", Kind: "rule", AccessedAt: "2026-08-02"},
		},
		Defaults: DefaultPolicies{
			Incorporated: ResolutionPolicy{
				Status: "local_record_required",
				AuthorityCandidates: []CandidateRule{
					{Kind: "municipality", Roles: []string{"adopts", "enforces"}, SourceIDs: []string{"src:local"}},
				},
				RequiredLocalRecords: []string{"Current municipal adoption instrument"},
				SourceIDs:            []string{"src:local"},
			},
			Unincorporated: ResolutionPolicy{
				Status: "local_record_required",
				AuthorityCandidates: []CandidateRule{
					{Kind: "county", Roles: []string{"adopts", "enforces"}, SourceIDs: []string{"src:local"}},
				},
				RequiredLocalRecords: []string{"Current county adoption instrument"},
				SourceIDs:            []string{"src:local"},
			},
		},
		CodeFamilyOverrides: map[string]ResolutionPolicy{
			"electrical": {
				Status:               "partially_resolved",
				RequiredLocalRecords: []string{"Local inspection authority"},
				SourceIDs:            []string{"src:electrical"},
			},
		},
		Verification: Verification{Status: "verified", Confidence: 0.95},
	}

	spec, err := BuildStateRegimeSpec(profile)
	if err != nil {
		t.Fatal(err)
	}
	if spec.RegimeID != "state-regime:us-co" || spec.SourceProfileID != profile.ProfileID {
		t.Fatalf("identity=%#v", spec)
	}
	if spec.Territory.Incorporated.Status != "local_record_required" || spec.Territory.Unincorporated.Status != "local_record_required" {
		t.Fatalf("territory=%#v", spec.Territory)
	}
	if spec.Dimensions.LocalAdoption.Status != RegimeEvidenceEvidenced {
		t.Fatalf("local adoption=%#v", spec.Dimensions.LocalAdoption)
	}
	if spec.Dimensions.Enforcement.Status != RegimeEvidenceEvidenced {
		t.Fatalf("enforcement=%#v", spec.Dimensions.Enforcement)
	}
	if spec.Dimensions.HomeRule.Status != RegimeEvidenceUnresolved {
		t.Fatalf("home rule=%#v", spec.Dimensions.HomeRule)
	}
	if spec.Dimensions.DelegationContracting.Status != RegimeEvidenceUnresolved {
		t.Fatalf("delegation=%#v", spec.Dimensions.DelegationContracting)
	}
	if spec.Dimensions.StatewideInheritance.Status != RegimeEvidenceUnresolved {
		t.Fatalf("inheritance=%#v", spec.Dimensions.StatewideInheritance)
	}
	if !reflect.DeepEqual(spec.RequiredLocalEvidence, []string{
		"Current county adoption instrument",
		"Current municipal adoption instrument",
		"Local inspection authority",
	}) {
		t.Fatalf("required local evidence=%#v", spec.RequiredLocalEvidence)
	}
	if len(spec.EntityPowerEvidence) != 2 {
		t.Fatalf("entity power evidence=%#v", spec.EntityPowerEvidence)
	}
	if spec.EntityPowerEvidence[0].EntityKind != "county" || !reflect.DeepEqual(spec.EntityPowerEvidence[0].Powers, []RegimePower{RegimePowerAdopt, RegimePowerEnforce}) {
		t.Fatalf("first power evidence=%#v", spec.EntityPowerEvidence[0])
	}
	if spec.CodeFamilyPolicies["electrical"].Status != "partially_resolved" {
		t.Fatalf("electrical=%#v", spec.CodeFamilyPolicies["electrical"])
	}
}

func TestBuildStateRegimeSpecRecognizesStatewideInheritanceFromDefaultAdoptions(t *testing.T) {
	profile := StateProfile{
		SchemaVersion:     ProfileSchemaVersion,
		ProfileID:         "state-profile:us-fl",
		StateID:           "US-FL",
		StateName:         "Florida",
		StateAbbreviation: "FL",
		StateFIPS:         "12",
		Status:            "verified",
		LastVerified:      "2026-08-02",
		Sources: []Source{
			{ID: "src:state", Title: "State code", URL: "https://example.invalid/state", Kind: "statute", AccessedAt: "2026-08-02"},
			{ID: "src:enforcement", Title: "Enforcement", URL: "https://example.invalid/enforcement", Kind: "statute", AccessedAt: "2026-08-02"},
		},
		Defaults: DefaultPolicies{
			Incorporated: ResolutionPolicy{
				Status:      "resolved",
				AdoptionIDs: []string{"adoption:building"},
				SourceIDs:   []string{"src:state"},
			},
			Unincorporated: ResolutionPolicy{
				Status:      "resolved",
				AdoptionIDs: []string{"adoption:building"},
				SourceIDs:   []string{"src:state"},
			},
		},
		Adoptions: []Adoption{
			{ID: "adoption:building", CodeFamily: "building", Status: "current", StateCodeName: "State Building Code", AdoptingAuthorityID: "auth:state", EnforcementModel: "local", Dates: AdoptionDates{EffectiveDate: "2023-01-01"}, SourceIDs: []string{"src:state"}, Verification: Verification{Status: "verified", Confidence: 0.95}},
		},
		EnforcementRules: []EnforcementRule{
			{ID: "rule:local-enforcement", CodeFamily: "building", Model: "local", EntityKinds: []string{"municipality", "county"}, Summary: "Local enforcing agencies administer the statewide code.", SourceIDs: []string{"src:enforcement"}, Verification: Verification{Status: "verified", Confidence: 0.95}},
		},
		Verification: Verification{Status: "verified", Confidence: 0.95},
	}

	spec, err := BuildStateRegimeSpec(profile)
	if err != nil {
		t.Fatal(err)
	}
	if spec.Dimensions.StatewideInheritance.Status != RegimeEvidenceEvidenced {
		t.Fatalf("inheritance=%#v", spec.Dimensions.StatewideInheritance)
	}
	if !reflect.DeepEqual(spec.Dimensions.StatewideInheritance.SourceIDs, []string{"src:state"}) {
		t.Fatalf("inheritance sources=%#v", spec.Dimensions.StatewideInheritance.SourceIDs)
	}
	if spec.Dimensions.Enforcement.Status != RegimeEvidenceEvidenced {
		t.Fatalf("enforcement=%#v", spec.Dimensions.Enforcement)
	}
	if spec.Dimensions.Temporal.Status != RegimeEvidenceEvidenced {
		t.Fatalf("temporal=%#v", spec.Dimensions.Temporal)
	}
}

func TestBuildStateRegimeSpecRecognizesAmendmentAndDelegationEvidenceWithoutInventingHomeRule(t *testing.T) {
	profile := StateProfile{
		SchemaVersion:     ProfileSchemaVersion,
		ProfileID:         "state-profile:us-test",
		StateID:           "US-TX",
		StateName:         "Test",
		StateAbbreviation: "TX",
		StateFIPS:         "48",
		Status:            "verified",
		LastVerified:      "2026-08-02",
		Sources: []Source{
			{ID: "src:amend", Title: "Amendment statute", URL: "https://example.invalid/amend", Kind: "statute", AccessedAt: "2026-08-02"},
			{ID: "src:delegate", Title: "Delegation statute", URL: "https://example.invalid/delegate", Kind: "statute", AccessedAt: "2026-08-02"},
		},
		Defaults: DefaultPolicies{
			Incorporated:   ResolutionPolicy{Status: "unresolved", SourceIDs: []string{"src:amend"}},
			Unincorporated: ResolutionPolicy{Status: "unresolved", SourceIDs: []string{"src:amend"}},
		},
		AmendmentRules: []AmendmentRule{
			{ID: "amend:local", Level: "local", Posture: "permitted_with_limits", Summary: "Local amendments are limited.", SourceIDs: []string{"src:amend"}, Verification: Verification{Status: "verified", Confidence: 0.9}},
		},
		Relationships: []AuthorityRelationship{
			{ID: "edge:delegate", FromID: "auth:state", Relationship: "delegates_enforcement_to", To: "local agencies", SourceIDs: []string{"src:delegate"}, Verification: Verification{Status: "verified", Confidence: 0.9}},
		},
		Verification: Verification{Status: "verified", Confidence: 0.9},
	}

	spec, err := BuildStateRegimeSpec(profile)
	if err != nil {
		t.Fatal(err)
	}
	if spec.Dimensions.LocalAmendment.Status != RegimeEvidenceEvidenced {
		t.Fatalf("amendment=%#v", spec.Dimensions.LocalAmendment)
	}
	if spec.Dimensions.DelegationContracting.Status != RegimeEvidenceEvidenced {
		t.Fatalf("delegation=%#v", spec.Dimensions.DelegationContracting)
	}
	if spec.Dimensions.HomeRule.Status != RegimeEvidenceUnresolved {
		t.Fatalf("home rule was inferred from unrelated evidence: %#v", spec.Dimensions.HomeRule)
	}
}
