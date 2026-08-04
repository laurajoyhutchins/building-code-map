package regulatory

import "testing"

func TestResolveVirginiaBuildingUsesStatewideCodeAndLocalEnforcement(t *testing.T) {
	catalog := loadTestCatalog(t)
	result, err := Resolve(catalog, ResolutionRequest{
		Context: &GeographicContext{
			StateID:      "US-VA",
			StateFIPS:    "51",
			StateName:    "Virginia",
			Incorporated: true,
			Municipality: &BoundaryMatch{Name: "Richmond"},
		},
		CodeFamily:        "building",
		ApplicabilityDate: "2026-08-04",
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Adoptions) != 1 || result.Adoptions[0].StateCodeName != "2021 Virginia Construction Code" {
		t.Fatalf("unexpected Virginia adoptions: %#v", result.Adoptions)
	}
	assertCandidateAuthority(t, result, "auth:us-va:board-housing-community-development")
	assertCandidateKind(t, result, "municipality")
	assertRuleIDPresent(t, result, "rule:us-va:2021-vcc-mandatory")
}

func TestResolveVirginiaTransitionPreservesPriorEditionOption(t *testing.T) {
	catalog := loadTestCatalog(t)
	result, err := Resolve(catalog, ResolutionRequest{
		Context:           &GeographicContext{StateID: "US-VA", StateFIPS: "51", Municipality: &BoundaryMatch{Name: "Norfolk"}},
		CodeFamily:        "building",
		ApplicabilityDate: "2024-06-01",
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Adoptions) != 1 || result.Adoptions[0].StateCodeName != "2021 Virginia Construction Code" {
		t.Fatalf("unexpected Virginia transition adoptions: %#v", result.Adoptions)
	}
	assertRuleIDPresent(t, result, "rule:us-va:2021-vcc-transition")
	assertContains(t, result.RequiredLocalRecords, "Permit application date and documented code-edition election during the 2021-code transition period")
}

func TestResolveOregonBuildingUses2025OSSC(t *testing.T) {
	catalog := loadTestCatalog(t)
	result, err := Resolve(catalog, ResolutionRequest{
		Context: &GeographicContext{
			StateID:      "US-OR",
			StateFIPS:    "41",
			StateName:    "Oregon",
			Incorporated: true,
			Municipality: &BoundaryMatch{Name: "Portland"},
		},
		CodeFamily:        "building",
		ApplicabilityDate: "2026-08-04",
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Adoptions) != 1 || result.Adoptions[0].StateCodeName != "2025 Oregon Structural Specialty Code" {
		t.Fatalf("unexpected Oregon adoptions: %#v", result.Adoptions)
	}
	assertCandidateAuthority(t, result, "auth:us-or:building-codes-division")
	assertCandidateKind(t, result, "municipality")
	assertRuleIDPresent(t, result, "rule:us-or:2025-ossc-mandatory")
}

func TestResolveOregonPhaseInPreservesPriorCodeOption(t *testing.T) {
	catalog := loadTestCatalog(t)
	result, err := Resolve(catalog, ResolutionRequest{
		Context:           &GeographicContext{StateID: "US-OR", StateFIPS: "41", Municipality: &BoundaryMatch{Name: "Salem"}},
		CodeFamily:        "building",
		ApplicabilityDate: "2025-12-01",
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Adoptions) != 1 || result.Adoptions[0].StateCodeName != "2025 Oregon Structural Specialty Code" {
		t.Fatalf("unexpected Oregon phase-in adoptions: %#v", result.Adoptions)
	}
	assertRuleIDPresent(t, result, "rule:us-or:2025-ossc-phase-in")
	assertContains(t, result.RequiredLocalRecords, "Permit application date and documented code-edition election during any applicable phase-in period")
}

func TestResolveNorthCarolinaKeeps2024CodePending(t *testing.T) {
	catalog := loadTestCatalog(t)
	result, err := Resolve(catalog, ResolutionRequest{
		Context: &GeographicContext{
			StateID:      "US-NC",
			StateFIPS:    "37",
			StateName:    "North Carolina",
			Incorporated: true,
			Municipality: &BoundaryMatch{Name: "Raleigh"},
		},
		CodeFamily:        "building",
		ApplicabilityDate: "2026-08-04",
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Adoptions) != 1 || result.Adoptions[0].StateCodeName != "2018 North Carolina State Building Code: Building Code" {
		t.Fatalf("unexpected North Carolina adoptions: %#v", result.Adoptions)
	}
	for _, adoption := range result.Adoptions {
		if adoption.StateCodeName == "2024 North Carolina State Building Code" {
			t.Fatalf("conditionally pending 2024 code leaked into current result: %#v", result.Adoptions)
		}
	}
	assertCandidateAuthority(t, result, "auth:us-nc:state-fire-marshal")
	assertCandidateKind(t, result, "municipality")
	assertRuleIDPresent(t, result, "rule:us-nc:2018-building-current")
	assertContains(t, result.RequiredLocalRecords, "Current State Fire Marshal certification status for the 2024 code transition")
}

func assertRuleIDPresent(t *testing.T, result ResolutionResult, id string) {
	t.Helper()
	for _, rule := range result.ApplicableRules {
		if rule.ID == id {
			return
		}
	}
	t.Fatalf("missing rule %q in %#v", id, result.ApplicableRules)
}
