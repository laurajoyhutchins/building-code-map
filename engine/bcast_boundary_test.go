package engine

import (
	"context"
	"encoding/json"
	"testing"
)

// BCAST identifies one exact publication materialization. It does not
// establish adoption, effective dates, amendments, or AHJ applicability.
func TestBCASTSourcePointersRemainUnverifiedProjectTraceEvidence(t *testing.T) {
	packageID := "bcastpkg:sha256:aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	objectID := "bcastobj:sha256:bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"
	runtime := NewRuntime(projectProvider{result: Resolution{
		Jurisdiction: JurisdictionContext{CountryCode: "US", PrimaryJurisdictionID: "DEMO-XX", PrimaryJurisdictionName: "Demonstration Jurisdiction"},
		Codes: []CodeResolution{{
			Family: "building", Edition: "DEMO-2026", Status: StatusInsufficientEvidence,
			Unresolved: []string{"No independent DEMO-XX adoption evidence"},
			Evidence: []EvidenceRef{
				{ID: packageID, Kind: "bcast.package"},
				{ID: objectID, Kind: "bcast.object"},
			},
		}},
	}})
	basis, err := NewProjectVerifier(runtime).VerifyProject(context.Background(), ProjectRequest{
		ProjectID: "DEMO-BCAST-001", Address: "100 Demo Plaza", ApplicabilityDate: "2026-10-01",
	})
	if err != nil {
		t.Fatalf("VerifyProject: %v", err)
	}
	if basis.Verdict != ProjectNotVerified || basis.CodeSet[0].Status != StatusInsufficientEvidence {
		t.Fatalf("BCAST source identity must not promote adoption or applicability: %#v", basis)
	}
	if len(basis.ExactEvidence) != 0 {
		t.Fatalf("BCAST trace pointers must not become authoritative exact evidence: %#v", basis.ExactEvidence)
	}
	if len(basis.Evidence) != 2 {
		t.Fatalf("trace reference count = %d", len(basis.Evidence))
	}
	payload, err := json.Marshal(basis)
	if err != nil {
		t.Fatal(err)
	}
	var decoded ProjectCodeBasis
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded.Evidence[0].ID != objectID || decoded.Evidence[0].Kind != "bcast.object" ||
		decoded.Evidence[1].ID != packageID || decoded.Evidence[1].Kind != "bcast.package" {
		t.Fatalf("BCAST trace identifiers lost in existing project contract: %#v", decoded.Evidence)
	}
	if decoded.Verdict != ProjectNotVerified {
		t.Fatalf("trace reference changed verdict: %q", decoded.Verdict)
	}
}

func TestBCASTSourcePointerAloneCannotEstablishLegallyMaterialClaim(t *testing.T) {
	// A source pointer is not an establishing legal text anchor, even if a
	// provider knows its package and object identity. BCM requires explicit
	// authoritative SourceDocument/SourceArtifact/TextAnchor evidence.
	claim := EvidenceClaim{
		ID: "DEMO-BCAST-ADOPTION", Kind: "code_adoption",
		JurisdictionID: "DEMO-XX", AsOf: "2026-10-01",
		LegallyMaterial: true, Links: nil,
	}
	result, err := VerifyClaimEvidence(claim, demoPolicy())
	if err != nil {
		t.Fatal(err)
	}
	if result.Decision != EvidenceInsufficient {
		t.Fatalf("BCAST identity without a qualifying legal evidence link should be insufficient: %#v", result)
	}
}
