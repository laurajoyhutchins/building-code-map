package engine

import (
	"context"
	"testing"
)

type projectProvider struct{ result Resolution }

func (provider projectProvider) Resolve(_ context.Context, query NormalizedQuery) (Resolution, error) {
	result := provider.result
	result.Query = query
	return result, nil
}

func TestProjectVerificationDoesNotRequireCodeFamily(t *testing.T) {
	runtime := NewRuntime(projectProvider{result: Resolution{
		Jurisdiction: JurisdictionContext{CountryCode: "US", PrimaryJurisdictionID: "DEMO-XX", PrimaryJurisdictionName: "Demonstration Jurisdiction"},
		Codes:        []CodeResolution{{Family: "building", Edition: "DEMO-2024", Status: StatusResolved, Basis: "DEMO ordinance"}},
	}})
	basis, err := NewProjectVerifier(runtime).VerifyProject(context.Background(), ProjectRequest{ProjectID: "DEMO-001", Address: "100 Demo Plaza", ApplicabilityDate: "2026-08-25"})
	if err != nil {
		t.Fatalf("VerifyProject() error = %v", err)
	}
	if basis.Verdict != ProjectVerified {
		t.Fatalf("verdict = %q, want verified", basis.Verdict)
	}
	if basis.Inputs.Address != "100 Demo Plaza" {
		t.Fatalf("address = %q", basis.Inputs.Address)
	}
}

func TestProjectVerificationDerivesConditionalFollowUpFromStructuredRequirement(t *testing.T) {
	runtime := NewRuntime(projectProvider{result: Resolution{
		Jurisdiction: JurisdictionContext{CountryCode: "US", PrimaryJurisdictionID: "DEMO-XX", PrimaryJurisdictionName: "Demonstration Jurisdiction"},
		Codes:        []CodeResolution{{Family: "building", Edition: "DEMO-2024", Status: StatusPartiallyResolved, Basis: "DEMO ordinance"}},
		Requirements: []ResolutionRequirement{{ID: "occupancy", Kind: RequirementProjectFact, FactKey: "occupancy", Prompt: "What is the project occupancy?"}},
	}})
	basis, err := NewProjectVerifier(runtime).VerifyProject(context.Background(), ProjectRequest{ProjectID: "DEMO-002", Address: "100 Demo Plaza", ApplicabilityDate: "2026-08-25"})
	if err != nil {
		t.Fatalf("VerifyProject() error = %v", err)
	}
	if basis.Verdict != ProjectConditional {
		t.Fatalf("verdict = %q, want conditional", basis.Verdict)
	}
	if len(basis.FollowUpQuestions) != 1 || basis.FollowUpQuestions[0] != "What is the project occupancy?" {
		t.Fatalf("follow_up_questions = %#v", basis.FollowUpQuestions)
	}
	if len(basis.Unresolved) != 1 || basis.Unresolved[0].Kind != RequirementProjectFact {
		t.Fatalf("unresolved = %#v", basis.Unresolved)
	}
}

func TestProjectVerificationMissingEvidenceIsNotNotApplicable(t *testing.T) {
	runtime := NewRuntime(projectProvider{result: Resolution{
		Jurisdiction: JurisdictionContext{CountryCode: "US", PrimaryJurisdictionID: "DEMO-XX", PrimaryJurisdictionName: "Demonstration Jurisdiction"},
		Codes:        []CodeResolution{{Family: "energy", Status: StatusInsufficientEvidence, Unresolved: []string{"DEMO adoption record missing"}}},
	}})
	basis, err := NewProjectVerifier(runtime).VerifyProject(context.Background(), ProjectRequest{ProjectID: "DEMO-003", Address: "100 Demo Plaza", ApplicabilityDate: "2026-08-25"})
	if err != nil {
		t.Fatalf("VerifyProject() error = %v", err)
	}
	if basis.Verdict != ProjectNotVerified {
		t.Fatalf("verdict = %q, want not_verified", basis.Verdict)
	}
	if basis.CodeSet[0].Status != StatusInsufficientEvidence {
		t.Fatalf("status = %q, want insufficient_evidence", basis.CodeSet[0].Status)
	}
}

func TestProjectVerificationPreservesFactsEvidenceAndProvenance(t *testing.T) {
	exact := demoEvidenceLink("DEMO-EXACT", EvidenceEstablishes)
	runtime := NewRuntime(projectProvider{result: Resolution{
		Jurisdiction: JurisdictionContext{CountryCode: "US", PrimaryJurisdictionID: "DEMO-XX", PrimaryJurisdictionName: "Demonstration Jurisdiction"},
		Codes: []CodeResolution{{
			Family: "building", Edition: "DEMO-2024", Status: StatusResolved, Basis: "DEMO ordinance",
			Evidence: []EvidenceRef{{ID: "DEMO-EVIDENCE", Kind: "synthetic"}}, ExactEvidence: []EvidenceLink{exact},
		}},
		DerivedFacts: []DerivedFact{{Key: "primary_jurisdiction", Value: "DEMO-XX", Basis: "synthetic resolver"}},
		Provenance:   Provenance{EngineVersion: "demo", SourceCommit: "DEMO-COMMIT", BundleID: "DEMO-BUNDLE"},
	}})
	basis, err := NewProjectVerifier(runtime).VerifyProject(context.Background(), ProjectRequest{
		ProjectID: "DEMO-004", Address: "100 Demo Plaza", ApplicabilityDate: "2026-08-25", Facts: map[string]string{" Occupancy ": "Business"},
	})
	if err != nil {
		t.Fatalf("VerifyProject() error = %v", err)
	}
	if basis.UserFacts["occupancy"] != "Business" {
		t.Fatalf("user facts = %#v", basis.UserFacts)
	}
	if len(basis.DerivedFacts) != 1 || basis.DerivedFacts[0].Key != "primary_jurisdiction" {
		t.Fatalf("derived facts = %#v", basis.DerivedFacts)
	}
	if len(basis.Evidence) != 1 || basis.Evidence[0].ID != "DEMO-EVIDENCE" {
		t.Fatalf("evidence = %#v", basis.Evidence)
	}
	if len(basis.ExactEvidence) != 1 || basis.ExactEvidence[0].Anchor.ID != "DEMO-ANCHOR" {
		t.Fatalf("exact evidence = %#v", basis.ExactEvidence)
	}
	if basis.Provenance.BundleID != "DEMO-BUNDLE" {
		t.Fatalf("provenance = %#v", basis.Provenance)
	}
}
