// Package accuracy compares independently reviewed regulatory expectations with
// published BCM project results. It is an evaluator, never a source of authority.
package accuracy

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/laurajoyhutchins/building-code-map/engine"
)

type ExpectedCode struct {
	Edition string                  `json:"edition,omitempty"`
	Status  engine.ResolutionStatus `json:"status"`
}

// Case is curated from source records independently of the BCM result under test.
type Case struct {
	ID             string         `json:"id"`
	ProjectID      string         `json:"project_id"`
	JurisdictionID string         `json:"jurisdiction_id"`
	Family         string         `json:"code_family"`
	AsOf           string         `json:"as_of"`
	ReviewedAt     string         `json:"reviewed_at"`
	SourceIDs      []string       `json:"source_ids"`
	Expected       []ExpectedCode `json:"expected"`
}

type Outcome string

const (
	OutcomeCorrect      Outcome = "correct"
	OutcomeFalseClaim   Outcome = "incorrect_assertion"
	OutcomeAbstained    Outcome = "abstained"
	OutcomeMissing      Outcome = "missing_result"
)

type CaseResult struct {
	ID             string         `json:"id"`
	ProjectID      string         `json:"project_id"`
	JurisdictionID string         `json:"jurisdiction_id"`
	Family         string         `json:"code_family"`
	Outcome        Outcome        `json:"outcome"`
	Observed       []ExpectedCode `json:"observed,omitempty"`
	Verdict        engine.ProjectVerdict `json:"project_verdict,omitempty"`
	FalseVerified  bool           `json:"false_verified,omitempty"`
}

type Summary struct {
	ReviewedCases     int `json:"reviewed_cases"`
	Correct           int `json:"correct"`
	IncorrectAssertions int `json:"incorrect_assertions"`
	Abstentions       int `json:"abstentions"`
	MissingResults    int `json:"missing_results"`
	FalseVerified     int `json:"false_verified"`
}

type Report struct {
	SchemaVersion string       `json:"schema_version"`
	Summary       Summary      `json:"summary"`
	Cases         []CaseResult `json:"cases"`
}

func validStatus(status engine.ResolutionStatus) bool {
	switch status {
	case engine.StatusResolved, engine.StatusNotApplicable, engine.StatusPartiallyResolved,
		engine.StatusLocalRecordRequired, engine.StatusUnsupported, engine.StatusInsufficientEvidence,
		engine.StatusAmbiguous, engine.StatusConflicting:
		return true
	default:
		return false
	}
}

func asserted(status engine.ResolutionStatus) bool {
	return status == engine.StatusResolved || status == engine.StatusNotApplicable
}

func sortedCodes(codes []ExpectedCode) []ExpectedCode {
	result := append([]ExpectedCode(nil), codes...)
	sort.Slice(result, func(i, j int) bool {
		if result[i].Status == result[j].Status {
			return result[i].Edition < result[j].Edition
		}
		return result[i].Status < result[j].Status
	})
	return result
}

func matchingCodes(expected, observed []ExpectedCode) bool {
	if len(expected) != len(observed) {
		return false
	}
	for i := range expected {
		if expected[i] != observed[i] {
			return false
		}
	}
	return true
}

func ValidateCases(cases []Case) error {
	if len(cases) == 0 {
		return fmt.Errorf("accuracy benchmark requires at least one independently reviewed case")
	}
	seenIDs := make(map[string]bool, len(cases))
	seenProjectFamilies := make(map[string]bool, len(cases))
	for _, c := range cases {
		if strings.TrimSpace(c.ID) == "" || strings.TrimSpace(c.ProjectID) == "" ||
			strings.TrimSpace(c.JurisdictionID) == "" || strings.TrimSpace(c.Family) == "" {
			return fmt.Errorf("case identity, project, jurisdiction, and code family are required")
		}
		if seenIDs[c.ID] {
			return fmt.Errorf("duplicate accuracy case id %q", c.ID)
		}
		seenIDs[c.ID] = true
		key := c.ProjectID + "\x00" + c.Family
		if seenProjectFamilies[key] {
			return fmt.Errorf("duplicate benchmark project/family %q/%q", c.ProjectID, c.Family)
		}
		seenProjectFamilies[key] = true
		if _, err := time.Parse(time.DateOnly, c.AsOf); err != nil {
			return fmt.Errorf("case %s requires a valid applicability date: %w", c.ID, err)
		}
		if _, err := time.Parse(time.DateOnly, c.ReviewedAt); err != nil {
			return fmt.Errorf("case %s requires a valid independent-review date: %w", c.ID, err)
		}
		if len(c.SourceIDs) == 0 {
			return fmt.Errorf("case %s requires primary source ids", c.ID)
		}
		for _, id := range c.SourceIDs {
			if strings.TrimSpace(id) == "" {
				return fmt.Errorf("case %s contains an empty source id", c.ID)
			}
		}
		if len(c.Expected) == 0 {
			return fmt.Errorf("case %s requires at least one expected code state", c.ID)
		}
		seenExpected := make(map[ExpectedCode]bool, len(c.Expected))
		for _, expected := range c.Expected {
			if !validStatus(expected.Status) {
				return fmt.Errorf("case %s has unknown expected status %q", c.ID, expected.Status)
			}
			if expected.Status == engine.StatusResolved && strings.TrimSpace(expected.Edition) == "" {
				return fmt.Errorf("case %s has a resolved claim with no edition", c.ID)
			}
			if seenExpected[expected] {
				return fmt.Errorf("case %s contains duplicate expected code state", c.ID)
			}
			seenExpected[expected] = true
		}
	}
	return nil
}

// Assess compares observed output to an independent reference. A correct
// assertion and a safe refusal to assert are intentionally different outcomes.
func Assess(cases []Case, results []engine.ProjectCodeBasis) (Report, error) {
	if err := ValidateCases(cases); err != nil {
		return Report{}, err
	}
	byProject := make(map[string]engine.ProjectCodeBasis, len(results))
	for _, basis := range results {
		if basis.ProjectID == "" {
			return Report{}, fmt.Errorf("observed project result has no project id")
		}
		if _, duplicate := byProject[basis.ProjectID]; duplicate {
			return Report{}, fmt.Errorf("duplicate observed project %q", basis.ProjectID)
		}
		byProject[basis.ProjectID] = basis
	}
	report := Report{SchemaVersion: "1", Summary: Summary{ReviewedCases: len(cases)}}
	for _, c := range cases {
		entry := CaseResult{ID: c.ID, ProjectID: c.ProjectID, JurisdictionID: c.JurisdictionID, Family: c.Family}
		basis, exists := byProject[c.ProjectID]
		if !exists {
			entry.Outcome = OutcomeMissing
		} else {
			entry.Verdict = basis.Verdict
			observed := make([]ExpectedCode, 0)
			for _, code := range basis.CodeSet {
				if code.Family == c.Family {
					observed = append(observed, ExpectedCode{Edition: code.Edition, Status: code.Status})
				}
			}
			entry.Observed = sortedCodes(observed)
			reference := sortedCodes(c.Expected)
			if len(entry.Observed) > 0 && basis.JurisdictionID == c.JurisdictionID &&
				basis.AsOf == c.AsOf && matchingCodes(reference, entry.Observed) {
				entry.Outcome = OutcomeCorrect
			} else if basis.JurisdictionID != c.JurisdictionID || basis.AsOf != c.AsOf {
				entry.Outcome = OutcomeFalseClaim
			} else {
				entry.Outcome = OutcomeAbstained
				for _, code := range entry.Observed {
					if asserted(code.Status) {
						entry.Outcome = OutcomeFalseClaim
						break
					}
				}
			}
			entry.FalseVerified = basis.Verdict == engine.ProjectVerified &&
				(entry.Outcome != OutcomeCorrect || !allAssertive(reference))
		}
		switch entry.Outcome {
		case OutcomeCorrect:
			report.Summary.Correct++
		case OutcomeFalseClaim:
			report.Summary.IncorrectAssertions++
		case OutcomeAbstained:
			report.Summary.Abstentions++
		case OutcomeMissing:
			report.Summary.MissingResults++
		}
		if entry.FalseVerified {
			report.Summary.FalseVerified++
		}
		report.Cases = append(report.Cases, entry)
	}
	sort.Slice(report.Cases, func(i, j int) bool { return report.Cases[i].ID < report.Cases[j].ID })
	return report, nil
}

func allAssertive(codes []ExpectedCode) bool {
	for _, code := range codes {
		if !asserted(code.Status) {
			return false
		}
	}
	return true
}
