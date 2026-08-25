package engine

import "testing"

const demoSHA = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
const demoTextSHA = "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"

func demoEvidenceLink(id string, role EvidenceRole) EvidenceLink {
	return EvidenceLink{
		ID: id, ClaimID: "DEMO-CLAIM", Role: role,
		Document: SourceDocument{ID: "DEMO-DOC", Kind: "ordinance", JurisdictionID: "DEMO-XX", CanonicalURL: "https://example.invalid/demo"},
		Artifact: SourceArtifact{ID: "DEMO-ARTIFACT", DocumentID: "DEMO-DOC", SHA256: demoSHA, SizeBytes: 1234, SourceURL: "https://example.invalid/demo.pdf"},
		Anchor: TextAnchor{ID: "DEMO-ANCHOR", ArtifactID: "DEMO-ARTIFACT", LegalLocator: "Section 1", TextSHA256: demoTextSHA},
	}
}

func demoPolicy() SourcePolicy {
	return SourcePolicy{Rules: []SourcePolicyRule{{ID: "DEMO-POLICY", ClaimKind: "code_adoption", JurisdictionID: "DEMO-XX", DocumentID: "DEMO-DOC", EffectiveFrom: "2026-01-01"}}}
}

func demoClaim(links ...EvidenceLink) EvidenceClaim {
	return EvidenceClaim{ID: "DEMO-CLAIM", Kind: "code_adoption", JurisdictionID: "DEMO-XX", AsOf: "2026-08-25", LegallyMaterial: true, Links: links}
}

func TestExactTextEvidenceQualifiesLegallyMaterialClaim(t *testing.T) {
	result, err := VerifyClaimEvidence(demoClaim(demoEvidenceLink("DEMO-LINK", EvidenceEstablishes)), demoPolicy())
	if err != nil {
		t.Fatalf("VerifyClaimEvidence() error = %v", err)
	}
	if result.Decision != EvidenceVerified || len(result.QualifyingLinkIDs) != 1 {
		t.Fatalf("result = %#v", result)
	}
}

func TestURLOnlyIdentityCannotEstablishClaim(t *testing.T) {
	link := demoEvidenceLink("DEMO-LINK", EvidenceEstablishes)
	link.Artifact.SHA256 = ""
	_, err := VerifyClaimEvidence(demoClaim(link), demoPolicy())
	if err == nil {
		t.Fatal("VerifyClaimEvidence() error = nil, want immutable artifact identity failure")
	}
}

func TestCorroborationWithoutEstablishingEvidenceIsInsufficient(t *testing.T) {
	result, err := VerifyClaimEvidence(demoClaim(demoEvidenceLink("DEMO-CORROBORATION", EvidenceCorroborates)), demoPolicy())
	if err != nil {
		t.Fatalf("VerifyClaimEvidence() error = %v", err)
	}
	if result.Decision != EvidenceInsufficient {
		t.Fatalf("decision = %q", result.Decision)
	}
}

func TestContradictingEvidenceFailsClosed(t *testing.T) {
	result, err := VerifyClaimEvidence(demoClaim(
		demoEvidenceLink("DEMO-ESTABLISHES", EvidenceEstablishes),
		demoEvidenceLink("DEMO-CONTRADICTS", EvidenceContradicts),
	), demoPolicy())
	if err != nil {
		t.Fatalf("VerifyClaimEvidence() error = %v", err)
	}
	if result.Decision != EvidenceConflicting {
		t.Fatalf("decision = %q", result.Decision)
	}
}

func TestSupersededEstablishingEvidenceDoesNotVerify(t *testing.T) {
	establishes := demoEvidenceLink("DEMO-OLD", EvidenceEstablishes)
	supersedes := demoEvidenceLink("DEMO-SUPERSEDES", EvidenceSupersedes)
	supersedes.TargetLinkID = establishes.ID
	result, err := VerifyClaimEvidence(demoClaim(establishes, supersedes), demoPolicy())
	if err != nil {
		t.Fatalf("VerifyClaimEvidence() error = %v", err)
	}
	if result.Decision != EvidenceSuperseded {
		t.Fatalf("decision = %q", result.Decision)
	}
}

func TestSourcePolicyUsesDocumentIdentityNotURLOrTier(t *testing.T) {
	link := demoEvidenceLink("DEMO-LINK", EvidenceEstablishes)
	link.Document.ID = "OTHER-DOC"
	link.Artifact.DocumentID = "OTHER-DOC"
	result, err := VerifyClaimEvidence(demoClaim(link), demoPolicy())
	if err != nil {
		t.Fatalf("VerifyClaimEvidence() error = %v", err)
	}
	if result.Decision != EvidenceInsufficient {
		t.Fatalf("decision = %q", result.Decision)
	}
}
