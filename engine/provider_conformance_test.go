package engine

import (
	"encoding/json"
	"os"
	"sort"
	"testing"
)

func readProviderFixture(t *testing.T, path string, target any) {
	t.Helper()
	contents, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	if err := json.Unmarshal(contents, target); err != nil {
		t.Fatalf("decode %s: %v", path, err)
	}
}

func diagnosticRuleIDs(diagnostics []ProviderConformanceDiagnostic) []string {
	seen := map[string]struct{}{}
	for _, diagnostic := range diagnostics {
		seen[diagnostic.RuleID] = struct{}{}
	}
	ids := make([]string, 0, len(seen))
	for id := range seen {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	return ids
}

func TestProviderConformanceRuleCatalogIsStableAndQualityScoped(t *testing.T) {
	expected := []string{
		"BCM-CONF-001",
		"BCM-CONF-002",
		"BCM-CONF-003",
		"BCM-CONF-004",
		"BCM-CONF-005",
		"BCM-CONF-006",
		"BCM-CONF-007",
		"BCM-CONF-008",
		"BCM-CONF-009",
	}
	allowedDimensions := map[string]bool{
		"accuracy": true,
		"completeness": true,
		"conformity": true,
		"consistency": true,
		"coverage": true,
		"timeliness": true,
		"uniqueness": true,
	}

	rules := ProviderConformanceRuleCatalog()
	if len(rules) != len(expected) {
		t.Fatalf("rule count = %d, want %d", len(rules), len(expected))
	}
	for index, rule := range rules {
		if rule.ID != expected[index] {
			t.Fatalf("rule[%d].ID = %q, want %q", index, rule.ID, expected[index])
		}
		if rule.Summary == "" {
			t.Fatalf("rule %s has empty summary", rule.ID)
		}
		if len(rule.QualityDimensions) == 0 {
			t.Fatalf("rule %s has no quality dimension", rule.ID)
		}
		for _, dimension := range rule.QualityDimensions {
			if !allowedDimensions[dimension] {
				t.Fatalf("rule %s uses unsupported quality dimension %q", rule.ID, dimension)
			}
		}
	}
}

func TestDEMOGoldenProviderConforms(t *testing.T) {
	contractBytes, err := os.ReadFile("../demo/DEMO-XX/provider-contract.json")
	if err != nil {
		t.Fatal(err)
	}
	contract, err := ParseProviderContractDeclaration(contractBytes)
	if err != nil {
		t.Fatalf("parse provider contract: %v", err)
	}
	var bundle ProviderBundle
	readProviderFixture(t, "../demo/DEMO-XX/provider-bundle.json", &bundle)

	if diagnostics := ValidateProviderConformance(contract, bundle); len(diagnostics) != 0 {
		t.Fatalf("DEMO-XX provider diagnostics = %#v", diagnostics)
	}
}

func TestSyntheticFailingFixtureExercisesEveryInitialRule(t *testing.T) {
	var fixture struct {
		Contract   ProviderContractDeclaration `json:"contract"`
		Bundle     ProviderBundle              `json:"bundle"`
		Resolution Resolution                  `json:"resolution"`
		Project    ProjectCodeBasis            `json:"project_code_basis"`
	}
	readProviderFixture(t, "../test/fixtures/provider-conformance/failing.json", &fixture)

	diagnostics := ValidateProviderConformance(fixture.Contract, fixture.Bundle)
	diagnostics = append(diagnostics, ValidateResolutionConformance(fixture.Resolution)...)
	diagnostics = append(diagnostics, ValidateProjectConformance(fixture.Project)...)

	expected := []string{
		"BCM-CONF-001",
		"BCM-CONF-002",
		"BCM-CONF-003",
		"BCM-CONF-004",
		"BCM-CONF-005",
		"BCM-CONF-006",
		"BCM-CONF-007",
		"BCM-CONF-008",
		"BCM-CONF-009",
	}
	actual := diagnosticRuleIDs(diagnostics)
	if len(actual) != len(expected) {
		t.Fatalf("diagnostic rule IDs = %v, want all %v; diagnostics=%#v", actual, expected, diagnostics)
	}
	for index := range expected {
		if actual[index] != expected[index] {
			t.Fatalf("diagnostic rule IDs = %v, want %v", actual, expected)
		}
	}
}

func TestProviderConformanceDiagnosticsAreDeterministic(t *testing.T) {
	var fixture struct {
		Contract ProviderContractDeclaration `json:"contract"`
		Bundle   ProviderBundle              `json:"bundle"`
	}
	readProviderFixture(t, "../test/fixtures/provider-conformance/failing.json", &fixture)

	first := ValidateProviderConformance(fixture.Contract, fixture.Bundle)
	second := ValidateProviderConformance(fixture.Contract, fixture.Bundle)
	firstJSON, _ := json.Marshal(first)
	secondJSON, _ := json.Marshal(second)
	if string(firstJSON) != string(secondJSON) {
		t.Fatalf("diagnostics are not deterministic:\n%s\n%s", firstJSON, secondJSON)
	}
	for index := 1; index < len(first); index++ {
		previous := first[index-1].RuleID + "\x00" + first[index-1].Path + "\x00" + first[index-1].Message
		current := first[index].RuleID + "\x00" + first[index].Path + "\x00" + first[index].Message
		if previous > current {
			t.Fatalf("diagnostics are not sorted at %d: %#v then %#v", index, first[index-1], first[index])
		}
	}
}

func TestNotApplicableRuleMatchesRuntimeValidation(t *testing.T) {
	resolution := Resolution{
		Jurisdiction: JurisdictionContext{CountryCode: "US", PrimaryJurisdictionID: "DEMO-XX", PrimaryJurisdictionName: "Demo"},
		Codes: []CodeResolution{{Family: "building", Status: StatusNotApplicable}},
	}

	diagnostics := ValidateResolutionConformance(resolution)
	if got := diagnosticRuleIDs(diagnostics); len(got) != 1 || got[0] != "BCM-CONF-007" {
		t.Fatalf("rule IDs = %v, want BCM-CONF-007", got)
	}
	if err := validateResolution(resolution); err == nil || err.Error() == "" {
		t.Fatal("runtime validation accepted not_applicable without an affirmative basis")
	}
}

func TestEvidenceDefectRuleChecksProjectOutputWithoutCreatingAnotherVerdictPath(t *testing.T) {
	project := ProjectCodeBasis{
		Verdict: ProjectConditional,
		Unresolved: []ProjectUnresolved{{Kind: RequirementEvidenceDefect, Message: "Synthetic evidence defect"}},
		FollowUpQuestions: []string{"Synthetic evidence defect"},
	}

	diagnostics := ValidateProjectConformance(project)
	if got := diagnosticRuleIDs(diagnostics); len(got) != 1 || got[0] != "BCM-CONF-008" {
		t.Fatalf("rule IDs = %v, want BCM-CONF-008", got)
	}
}