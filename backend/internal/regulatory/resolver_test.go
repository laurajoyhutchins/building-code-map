package regulatory

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestResolveColoradoBuildingRequiresLocalAdoptionRecord(t *testing.T) {
	catalog := loadTestCatalog(t)
	result, err := Resolve(catalog, ResolutionRequest{
		Context: &GeographicContext{
			StateID:      "US-CO",
			StateFIPS:    "08",
			StateName:    "Colorado",
			Incorporated: true,
			Municipality: &BoundaryMatch{LayerFamily: "municipalities", FeatureID: "0820000", Name: "Denver"},
			County:       &BoundaryMatch{LayerFamily: "counties", FeatureID: "08031", Name: "Denver County"},
		},
		CodeFamily: "building",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Status != "local_record_required" {
		t.Fatalf("status = %q, want local_record_required", result.Status)
	}
	if len(result.Adoptions) != 0 {
		t.Fatalf("Colorado building resolution must not invent a statewide adoption: %#v", result.Adoptions)
	}
	assertCandidateKind(t, result, "municipality")
	assertContains(t, result.RequiredLocalRecords, "Current municipal building, fire, and energy-code adoption instruments")
	assertRuleKind(t, result, "applicability")
	assertRuleKind(t, result, "amendment")
	if result.PolicyBasis == nil || result.PolicyBasis.Status != "local_record_required" {
		t.Fatalf("missing source-bearing policy basis: %#v", result.PolicyBasis)
	}
	assertContains(t, result.PolicyBasis.SourceIDs, "src:us-co:energy-local")
}

func TestResolveColoradoElectricalAddsStateAdoption(t *testing.T) {
	catalog := loadTestCatalog(t)
	result, err := Resolve(catalog, ResolutionRequest{
		Context:    &GeographicContext{StateID: "US-CO", StateFIPS: "08", County: &BoundaryMatch{Name: "Larimer County"}},
		CodeFamily: "electrical",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Status != "partially_resolved" {
		t.Fatalf("status = %q, want partially_resolved", result.Status)
	}
	if len(result.Adoptions) != 1 || result.Adoptions[0].CodeFamily != "electrical" {
		t.Fatalf("unexpected adoptions: %#v", result.Adoptions)
	}
	assertCandidateAuthority(t, result, "auth:us-co:state-electrical-board")
	assertRuleKind(t, result, "enforcement")
}

func TestResolveFloridaUsesStatewideBuildingCodeWithoutSuppressingLocalFollowUp(t *testing.T) {
	catalog := loadTestCatalog(t)
	result, err := Resolve(catalog, ResolutionRequest{
		Context: &GeographicContext{
			StateID:      "US-FL",
			StateFIPS:    "12",
			Incorporated: true,
			Municipality: &BoundaryMatch{Name: "Orlando"},
			County:       &BoundaryMatch{Name: "Orange County"},
		},
		CodeFamily:        "building",
		ApplicabilityDate: "2026-07-20",
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Adoptions) != 1 || result.Adoptions[0].StateCodeName != "Florida Building Code, Building, 8th Edition (2023)" {
		t.Fatalf("unexpected adoptions: %#v", result.Adoptions)
	}
	assertCandidateKind(t, result, "municipality")
	assertContains(t, result.RequiredLocalRecords, "Current local technical amendments and administrative procedures")
	if len(result.AuthorityPath) != 1 {
		t.Fatalf("expected authority path, got %#v", result.AuthorityPath)
	}
	for _, kind := range []string{"applicability", "date", "amendment", "enforcement"} {
		assertRuleKind(t, result, kind)
	}
	if len(result.SupportingClaims) < 2 {
		t.Fatalf("expected adoption and relationship claims, got %#v", result.SupportingClaims)
	}
}

func TestResolveFloridaHistoricalDateUsesPriorEdition(t *testing.T) {
	catalog := loadTestCatalog(t)
	result, err := Resolve(catalog, ResolutionRequest{
		Context:           &GeographicContext{StateID: "US-FL", StateFIPS: "12", Municipality: &BoundaryMatch{Name: "Orlando"}},
		CodeFamily:        "building",
		ApplicabilityDate: "2023-01-01",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Status != "partially_resolved" || len(result.Adoptions) != 1 {
		t.Fatalf("historical result did not select one prior edition: %#v", result)
	}
	if result.Adoptions[0].ID != "adoption:us-fl:building:2020" {
		t.Fatalf("historical result selected wrong adoption: %#v", result.Adoptions)
	}
}

func TestResolveNewJerseyHistoricalDatePreservesTransitionEditions(t *testing.T) {
	catalog := loadTestCatalog(t)
	result, err := Resolve(catalog, ResolutionRequest{
		Context:           &GeographicContext{StateID: "US-NJ", StateFIPS: "34", Municipality: &BoundaryMatch{Name: "Trenton"}},
		CodeFamily:        "building",
		ApplicabilityDate: "2020-01-01",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Status != "partially_resolved" || len(result.Adoptions) != 2 {
		t.Fatalf("historical transition result = %#v, want two supported editions", result)
	}
	ids := map[string]bool{}
	for _, adoption := range result.Adoptions {
		ids[adoption.ID] = true
	}
	if !ids["adoption:us-nj:building:2015"] || !ids["adoption:us-nj:building:2018"] {
		t.Fatalf("historical transition editions = %#v", result.Adoptions)
	}
}

func TestResolveNewJerseyOperationalFireUsesSeparateAuthority(t *testing.T) {
	catalog := loadTestCatalog(t)
	result, err := Resolve(catalog, ResolutionRequest{
		Context: &GeographicContext{
			StateID:           "US-NJ",
			StateFIPS:         "34",
			Incorporated:      true,
			Municipality:      &BoundaryMatch{Name: "Trenton"},
			FireJurisdictions: []BoundaryMatch{{Name: "Trenton Fire Prevention Bureau"}},
		},
		CodeFamily: "fire",
	})
	if err != nil {
		t.Fatal(err)
	}
	assertCandidateAuthority(t, result, "auth:us-nj:division-fire-safety")
	assertCandidateKind(t, result, "fire_authority")
	if len(result.Adoptions) != 1 || result.Adoptions[0].CodeFamily != "fire_operational" {
		t.Fatalf("unexpected fire adoptions: %#v", result.Adoptions)
	}
	assertRuleKind(t, result, "enforcement")
}

func TestResolveFloridaOperationalFireExcludesConstructionRules(t *testing.T) {
	catalog := loadTestCatalog(t)
	result, err := Resolve(catalog, ResolutionRequest{
		Context:    &GeographicContext{StateID: "US-FL", StateFIPS: "12", Municipality: &BoundaryMatch{Name: "Orlando"}},
		CodeFamily: "fire",
	})
	if err != nil {
		t.Fatal(err)
	}
	assertRuleIDAbsent(t, result, "rule:us-fl:local-technical-amendments")
	assertRuleIDAbsent(t, result, "rule:us-fl:local-enforcement")
}

func TestResolveNewJerseyStateOwnedProjectAppliesProjectOverride(t *testing.T) {
	catalog := loadTestCatalog(t)
	result, err := Resolve(catalog, ResolutionRequest{
		Context:     &GeographicContext{StateID: "US-NJ", StateFIPS: "34", Municipality: &BoundaryMatch{Name: "Trenton"}},
		CodeFamily:  "building",
		ProjectType: "state_owned",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Status != "resolved" {
		t.Fatalf("status = %q, want resolved", result.Status)
	}
	assertCandidateAuthority(t, result, "auth:us-nj:bcpr")
}

func TestResolveSurfacesConflictingRelevantClaims(t *testing.T) {
	catalog := loadTestCatalog(t)
	profile, ok := catalog.Profile("US-FL", "")
	if !ok {
		t.Fatal("Florida profile missing")
	}
	profile.Claims = append(profile.Claims, Claim{
		ID:            "claim:us-fl:building-effective-date:conflict",
		SubjectID:     "adoption:us-fl:building:2023",
		Field:         "dates.effective_date",
		Status:        "conflicting",
		Value:         "2024-01-01",
		ConflictGroup: "conflict:us-fl:building-effective-date",
		SourceIDs:     []string{"src:us-fl:statute-553-73"},
		Verification:  Verification{Status: "needs_review", Confidence: 0.5},
	})
	conflictingCatalog, err := NewCatalog([]StateProfile{profile})
	if err != nil {
		t.Fatal(err)
	}
	result, err := Resolve(conflictingCatalog, ResolutionRequest{
		Context:           &GeographicContext{StateID: "US-FL", StateFIPS: "12", Municipality: &BoundaryMatch{Name: "Orlando"}},
		CodeFamily:        "building",
		ApplicabilityDate: "2026-07-20",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Status != "conflicting" {
		t.Fatalf("status = %q, want conflicting", result.Status)
	}
	if len(result.SupportingClaims) < 2 {
		t.Fatalf("conflicting claims were not preserved: %#v", result.SupportingClaims)
	}
}

func TestResolveRejectsInvalidApplicabilityDate(t *testing.T) {
	catalog := loadTestCatalog(t)
	_, err := Resolve(catalog, ResolutionRequest{
		Context:           &GeographicContext{StateID: "US-FL", StateFIPS: "12"},
		ApplicabilityDate: "yesterday",
	})
	if err == nil || !strings.Contains(err.Error(), "YYYY-MM-DD") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestResolveUnknownStateFailsClosed(t *testing.T) {
	catalog := loadTestCatalog(t)
	result, err := Resolve(catalog, ResolutionRequest{Context: &GeographicContext{StateID: "US-XX", StateFIPS: "99"}})
	if err != nil {
		t.Fatal(err)
	}
	if result.Status != "insufficient_evidence" {
		t.Fatalf("status = %q, want insufficient_evidence", result.Status)
	}
	if len(result.AuthorityCandidates) != 0 || len(result.Adoptions) != 0 {
		t.Fatalf("unknown state must not return unsupported conclusions: %#v", result)
	}
}

func TestResolveAtUsesExplicitGenerationTime(t *testing.T) {
	catalog := loadTestCatalog(t)
	want := time.Date(2026, 8, 6, 12, 34, 56, 0, time.UTC)
	result, err := ResolveAt(catalog, ResolutionRequest{
		Context:           &GeographicContext{StateID: "US-CO", StateFIPS: "08"},
		ApplicabilityDate: "2026-08-06",
	}, want)
	if err != nil {
		t.Fatal(err)
	}
	if !result.GeneratedAt.Equal(want) {
		t.Fatalf("generated_at=%s, want %s", result.GeneratedAt, want)
	}
}

func TestCatalogMergesRulePacksAndSourceHealth(t *testing.T) {
	catalog := loadTestCatalog(t)
	profile, ok := catalog.Profile("US-FL", "")
	if !ok {
		t.Fatal("Florida profile missing")
	}
	if len(profile.ApplicabilityRules) == 0 || len(profile.DateRules) == 0 || len(profile.AmendmentRules) == 0 || len(profile.EnforcementRules) == 0 {
		t.Fatalf("rule pack was not merged: %#v", profile)
	}
	if profile.Sources[0].LastCheckedAt == "" || profile.Sources[0].Availability != "available" {
		t.Fatalf("source health was not merged: %#v", profile.Sources)
	}
}

func TestProfilesRoundTripDeterministically(t *testing.T) {
	catalog := loadTestCatalog(t)
	for _, profile := range catalog.Profiles() {
		raw, err := json.Marshal(profile)
		if err != nil {
			t.Fatal(err)
		}
		var decoded StateProfile
		if err := json.Unmarshal(raw, &decoded); err != nil {
			t.Fatal(err)
		}
		if err := ValidateCompleteProfile(decoded); err != nil {
			t.Fatalf("round-trip profile %s: %v", profile.ProfileID, err)
		}
	}
}

func loadTestCatalog(t *testing.T) Catalog {
	t.Helper()
	catalog, err := LoadCatalog(filepath.Join("..", "..", "data", "regulatory"))
	if err == nil {
		return catalog
	}
	cwd, _ := os.Getwd()
	t.Fatalf("load catalog from %s: %v", cwd, err)
	return Catalog{}
}

func assertCandidateKind(t *testing.T, result ResolutionResult, kind string) {
	t.Helper()
	for _, candidate := range result.AuthorityCandidates {
		if candidate.Kind == kind {
			return
		}
	}
	t.Fatalf("missing candidate kind %q in %#v", kind, result.AuthorityCandidates)
}

func assertCandidateAuthority(t *testing.T, result ResolutionResult, id string) {
	t.Helper()
	for _, candidate := range result.AuthorityCandidates {
		if candidate.AuthorityID == id {
			return
		}
	}
	t.Fatalf("missing authority %q in %#v", id, result.AuthorityCandidates)
}

func assertRuleKind(t *testing.T, result ResolutionResult, kind string) {
	t.Helper()
	for _, rule := range result.ApplicableRules {
		if rule.Kind == kind {
			return
		}
	}
	t.Fatalf("missing rule kind %q in %#v", kind, result.ApplicableRules)
}

func assertRuleIDAbsent(t *testing.T, result ResolutionResult, id string) {
	t.Helper()
	for _, rule := range result.ApplicableRules {
		if rule.ID == id {
			t.Fatalf("unexpected rule %q in %#v", id, result.ApplicableRules)
		}
	}
}

func assertContains(t *testing.T, values []string, wanted string) {
	t.Helper()
	for _, value := range values {
		if value == wanted {
			return
		}
	}
	t.Fatalf("missing %q in %#v", wanted, values)
}
