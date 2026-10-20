package engine

import (
	"context"
	"errors"
	"testing"
)

func TestResolvedClaimRequiresRequestSpecificQualifyingEvidence(t *testing.T) {
	establishes := demoEvidenceLink("DEMO-ESTABLISHES", EvidenceEstablishes)
	conflicts := demoEvidenceLink("DEMO-CONFLICTS", EvidenceContradicts)
	supersedes := demoEvidenceLink("DEMO-SUPERSEDES", EvidenceSupersedes)
	supersedes.TargetLinkID = establishes.ID

	tests := []struct {
		name   string
		links  []EvidenceLink
		policy SourcePolicy
		wantOK bool
	}{
		{name: "bare resolved status", policy: demoPolicy()},
		{name: "source ID alone is not evidence", policy: demoPolicy()},
		{name: "no trusted policy", links: []EvidenceLink{establishes}},
		{name: "approved exact text", links: []EvidenceLink{establishes}, policy: demoPolicy(), wantOK: true},
		{name: "corroboration only", links: []EvidenceLink{demoEvidenceLink("DEMO-CORROBORATES", EvidenceCorroborates)}, policy: demoPolicy()},
		{name: "conflicting link", links: []EvidenceLink{establishes, conflicts}, policy: demoPolicy()},
		{name: "superseded link", links: []EvidenceLink{establishes, supersedes}, policy: demoPolicy()},
		{name: "unapproved source", links: []EvidenceLink{establishes}, policy: SourcePolicy{}},
		{name: "wrong applicability date", links: []EvidenceLink{establishes}, policy: SourcePolicy{Rules: []SourcePolicyRule{{
			ID: "DEMO-LATE", ClaimKind: "code_adoption", JurisdictionID: "DEMO-XX",
			DocumentID: "DEMO-DOC", EffectiveFrom: "2027-01-01",
		}}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code := CodeResolution{
				Family: "building", Edition: "DEMO-2024", Status: StatusResolved,
				Basis: "Synthetic adoption", ExactEvidence: tt.links,
				Evidence: []EvidenceRef{{ID: "DEMO-SOURCE", Kind: "synthetic"}},
			}
			runtime := NewRuntimeWithSourcePolicy(syntheticProvider{codes: []CodeResolution{code}}, tt.policy)
			basis, err := NewProjectVerifier(runtime).VerifyProject(context.Background(), ProjectRequest{
				ProjectID: "DEMO-001", Address: "100 Demo Plaza", ApplicabilityDate: "2026-08-25",
			})
			if tt.wantOK {
				if err != nil || basis.Verdict != ProjectVerified {
					t.Fatalf("qualified claim: basis=%#v error=%v", basis, err)
				}
				return
			}
			var engineErr EngineError
			if !errors.As(err, &engineErr) || engineErr.Code != ErrorDataBundleInvalid {
				t.Fatalf("rejected claim: error=%v, want data_bundle_invalid", err)
			}
		})
	}
}

func TestMalformedProviderFieldsNeverEnterProjectCodeBasis(t *testing.T) {
	tests := []struct {
		name      string
		facts     []DerivedFact
		references []EvidenceRef
	}{
		{name: "missing derived key", facts: []DerivedFact{{Value: "DEMO-XX"}}},
		{name: "missing derived value", facts: []DerivedFact{{Key: "jurisdiction"}}},
		{name: "blank derived key", facts: []DerivedFact{{Key: " ", Value: "DEMO-XX"}}},
		{name: "missing source ID", references: []EvidenceRef{{Kind: "synthetic"}}},
		{name: "missing source kind", references: []EvidenceRef{{ID: "DEMO-SOURCE"}}},
		{name: "blank source kind", references: []EvidenceRef{{ID: "DEMO-SOURCE", Kind: "  "}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			runtime := NewRuntime(projectProvider{result: Resolution{
				Jurisdiction: JurisdictionContext{
					CountryCode: "US", PrimaryJurisdictionID: "DEMO-XX", PrimaryJurisdictionName: "Demo",
				},
				Codes: []CodeResolution{{Family: "building", Status: StatusPartiallyResolved, Evidence: tt.references}},
				DerivedFacts: tt.facts,
			}})
			_, err := NewProjectVerifier(runtime).VerifyProject(context.Background(), ProjectRequest{
				ProjectID: "DEMO-001", Address: "100 Demo Plaza", ApplicabilityDate: "2026-08-25",
			})
			var engineErr EngineError
			if !errors.As(err, &engineErr) || engineErr.Code != ErrorDataBundleInvalid {
				t.Fatalf("error=%v, want data_bundle_invalid", err)
			}
		})
	}
}

func TestAffirmativeNonApplicabilityRequiresLegalProof(t *testing.T) {
	link := demoEvidenceLink("DEMO-NOT-APPLICABLE", EvidenceEstablishes)
	for _, tt := range []struct {
		name string
		policy SourcePolicy
		links []EvidenceLink
		wantVerified bool
	}{
		{name: "bare provider assertion"},
		{name: "approved claim", policy: SourcePolicy{Rules: []SourcePolicyRule{{
			ID: "DEMO-EXCLUSION", ClaimKind: "code_non_applicability",
			JurisdictionID: "DEMO-XX", DocumentID: "DEMO-DOC", EffectiveFrom: "2026-01-01",
		}}}, links: []EvidenceLink{link}, wantVerified: true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			runtime := NewRuntimeWithSourcePolicy(syntheticProvider{codes: []CodeResolution{{
				Family: "fire", Status: StatusNotApplicable, Basis: "Synthetic exclusion",
				ExactEvidence: tt.links,
			}}}, tt.policy)
			basis, err := NewProjectVerifier(runtime).VerifyProject(context.Background(), ProjectRequest{
				ProjectID: "DEMO-001", Address: "100 Demo Plaza", ApplicabilityDate: "2026-08-25",
			})
			if tt.wantVerified {
				if err != nil || basis.Verdict != ProjectVerified {
					t.Fatalf("basis=%#v error=%v", basis, err)
				}
				return
			}
			var engineErr EngineError
			if !errors.As(err, &engineErr) || engineErr.Code != ErrorDataBundleInvalid {
				t.Fatalf("error=%v, want data_bundle_invalid", err)
			}
		})
	}
}
