package accuracy

import (
	"testing"

	"github.com/laurajoyhutchins/building-code-map/engine"
)

func demoTruth(status engine.ResolutionStatus, edition string) Case {
	return Case{
		ID: "DEMO-CASE", ProjectID: "DEMO-PROJECT", JurisdictionID: "DEMO-XX",
		Family: "building", AsOf: "2026-08-25", ReviewedAt: "2026-08-26",
		SourceIDs: []string{"DEMO-PRIMARY-SOURCE"},
		Expected: []ExpectedCode{{Status: status, Edition: edition}},
	}
}

func demoResult(status engine.ResolutionStatus, edition string, verdict engine.ProjectVerdict) engine.ProjectCodeBasis {
	return engine.ProjectCodeBasis{
		ProjectID: "DEMO-PROJECT", JurisdictionID: "DEMO-XX", AsOf: "2026-08-25",
		Verdict: verdict,
		CodeSet: []engine.ProjectCodeEntry{{Family: "building", Edition: edition, Status: status}},
	}
}

func TestAccuracyBenchmarkDistinguishesCorrectClaimsFromAbstention(t *testing.T) {
	for _, tt := range []struct {
		name    string
		truth   Case
		actual  engine.ProjectCodeBasis
		outcome Outcome
		falseVerified bool
	}{
		{
			name: "correct verified edition",
			truth: demoTruth(engine.StatusResolved, "DEMO-2024"),
			actual: demoResult(engine.StatusResolved, "DEMO-2024", engine.ProjectVerified),
			outcome: OutcomeCorrect,
		},
		{
			name: "wrong edition is an incorrect assertion",
			truth: demoTruth(engine.StatusResolved, "DEMO-2024"),
			actual: demoResult(engine.StatusResolved, "DEMO-2021", engine.ProjectVerified),
			outcome: OutcomeFalseClaim, falseVerified: true,
		},
		{
			name: "safe refusal is not an incorrect legal claim",
			truth: demoTruth(engine.StatusResolved, "DEMO-2024"),
			actual: demoResult(engine.StatusInsufficientEvidence, "", engine.ProjectNotVerified),
			outcome: OutcomeAbstained,
		},
		{
			name: "affirmative non-applicability needs matching legal truth",
			truth: demoTruth(engine.StatusResolved, "DEMO-2024"),
			actual: demoResult(engine.StatusNotApplicable, "", engine.ProjectVerified),
			outcome: OutcomeFalseClaim, falseVerified: true,
		},
		{
			name: "correct conditional assessment",
			truth: demoTruth(engine.StatusLocalRecordRequired, ""),
			actual: demoResult(engine.StatusLocalRecordRequired, "", engine.ProjectConditional),
			outcome: OutcomeCorrect,
		},
		{
			name: "verified verdict for unresolved reference is misleading",
			truth: demoTruth(engine.StatusLocalRecordRequired, ""),
			actual: demoResult(engine.StatusLocalRecordRequired, "", engine.ProjectVerified),
			outcome: OutcomeCorrect, falseVerified: true,
		},
		{
			name: "verified result in wrong jurisdiction",
			truth: demoTruth(engine.StatusResolved, "DEMO-2024"),
			actual: func() engine.ProjectCodeBasis {
				basis := demoResult(engine.StatusResolved, "DEMO-2024", engine.ProjectVerified)
				basis.JurisdictionID = "DEMO-OTHER"
				return basis
			}(),
			outcome: OutcomeFalseClaim, falseVerified: true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			report, err := Assess([]Case{tt.truth}, []engine.ProjectCodeBasis{tt.actual})
			if err != nil {
				t.Fatal(err)
			}
			if len(report.Cases) != 1 || report.Cases[0].Outcome != tt.outcome ||
				report.Cases[0].FalseVerified != tt.falseVerified {
				t.Fatalf("report = %#v, want %q falseVerified=%t", report, tt.outcome, tt.falseVerified)
			}
		})
	}
}

func TestAccuracyBenchmarkFlagsOmittedCodeFamilyFromVerifiedProject(t *testing.T) {
	c := demoTruth(engine.StatusResolved, "DEMO-2024")
	basis := demoResult(engine.StatusResolved, "DEMO-2024", engine.ProjectVerified)
	basis.CodeSet = []engine.ProjectCodeEntry{{Family: "electrical", Edition: "DEMO-2023", Status: engine.StatusResolved}}
	report, err := Assess([]Case{c}, []engine.ProjectCodeBasis{basis})
	if err != nil {
		t.Fatal(err)
	}
	if report.Summary.Abstentions != 1 || report.Summary.FalseVerified != 1 {
		t.Fatalf("omission must not be reported as accurate: %#v", report)
	}
}

func TestAccuracyBenchmarkRequiresIndependentSourceDenominator(t *testing.T) {
	c := demoTruth(engine.StatusResolved, "DEMO-2024")
	for _, tt := range []struct {
		name string
		change func(*Case)
	}{
		{name: "no primary sources", change: func(c *Case) { c.SourceIDs = nil }},
		{name: "no review date", change: func(c *Case) { c.ReviewedAt = "" }},
		{name: "no project", change: func(c *Case) { c.ProjectID = "" }},
		{name: "bad applicability date", change: func(c *Case) { c.AsOf = "today" }},
		{name: "resolved missing edition", change: func(c *Case) { c.Expected[0].Edition = "" }},
	} {
		t.Run(tt.name, func(t *testing.T) {
			copyCase := c
			copyCase.Expected = append([]ExpectedCode(nil), c.Expected...)
			tt.change(&copyCase)
			if err := ValidateCases([]Case{copyCase}); err == nil {
				t.Fatal("invalid independent case was accepted")
			}
		})
	}
	if _, err := Assess([]Case{c, c}, nil); err == nil {
		t.Fatal("duplicate case was accepted")
	}
}

func TestAccuracyBenchmarkCountsMissingResultsSeparately(t *testing.T) {
	c := demoTruth(engine.StatusResolved, "DEMO-2024")
	report, err := Assess([]Case{c}, nil)
	if err != nil {
		t.Fatal(err)
	}
	if report.Summary.MissingResults != 1 || report.Summary.IncorrectAssertions != 0 {
		t.Fatalf("missing results must not become incorrect assertions: %#v", report)
	}
}
