package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/laurajoyhutchins/building-code-map/engine"
)

func TestPinnedSourcePolicyRejectsDriftAndInvalidRules(t *testing.T) {
	path := filepath.Join(t.TempDir(), "private-policy.json")
	valid := `{"rules":[{"id":"synthetic-demo-policy","claim_kind":"code_adoption","jurisdiction_id":"DEMO-XX","document_id":"demo-instrument","artifact_sha256":"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa","text_sha256":"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb","legal_locator":"DEMO-XX 1.1","effective_from":"2025-01-01","effective_to":"2026-12-31"}]}`
	verify := func(contents string, wantError bool) {
		t.Helper()
		if err := os.WriteFile(path, []byte(contents), 0600); err != nil {
			t.Fatal(err)
		}
		digest := sha256.Sum256([]byte(contents))
		policy, err := loadPinnedSourcePolicy(path, hex.EncodeToString(digest[:]))
		if wantError && err == nil {
			t.Fatalf("invalid rule accepted: %#v", policy)
		}
		if !wantError && (err != nil || len(policy.Rules) != 1) {
			t.Fatalf("valid policy rejected: %#v, %v", policy, err)
		}
	}
	verify(valid, false)
	if err := os.WriteFile(path, []byte(valid+" "), 0600); err != nil {
		t.Fatal(err)
	}
	digest := sha256.Sum256([]byte(valid))
	if _, err := loadPinnedSourcePolicy(path, hex.EncodeToString(digest[:])); err == nil {
		t.Fatal("changed policy passed original source digest")
	}
	verify(valid+`{"extra":true}`, true)
	verify(strings.Replace(valid, "2026-12-31", "2024-01-01", 1), true)
	verify(strings.Replace(valid, "code_adoption", "unknown_kind", 1), true)
	verify(strings.Replace(valid, "2025-01-01", "not-a-date", 1), true)
	verify(strings.Replace(valid, strings.Repeat("a", 64), strings.Repeat("z", 64), 1), true)
	verify(strings.Replace(valid, `"legal_locator":"DEMO-XX 1.1"`, `"legal_locator":""`, 1), true)
	verify(strings.Replace(valid, `"rules":`, `"unknown":1,"rules":`, 1), true)
	if _, err := loadPinnedSourcePolicy(path, "invalid"); err == nil {
		t.Fatal("invalid digest length accepted")
	}
}

type syntheticEvidenceProvider struct{}

func (syntheticEvidenceProvider) Resolve(_ context.Context, q engine.NormalizedQuery) (engine.Resolution, error) {
	link := engine.EvidenceLink{
		ID: "demo-establishing-link", ClaimID: "demo-claim",
		Role: engine.EvidenceEstablishes,
		Document: engine.SourceDocument{ID: "demo-instrument", Kind: "synthetic-law",
			JurisdictionID: "DEMO-XX"},
		Artifact: engine.SourceArtifact{ID: "demo-artifact", DocumentID: "demo-instrument",
			SHA256: strings.Repeat("a", 64), SizeBytes: 120},
		Anchor: engine.TextAnchor{ID: "demo-anchor", ArtifactID: "demo-artifact",
			LegalLocator: "DEMO-XX 1.1", TextSHA256: strings.Repeat("b", 64)},
	}
	return engine.Resolution{
		Jurisdiction: engine.JurisdictionContext{
			CountryCode: "DEMO", PrimaryJurisdictionID: "DEMO-XX",
			PrimaryJurisdictionName: "Fictional jurisdiction",
		},
		Codes: []engine.CodeResolution{{
			Family: "building", Edition: "DEMO-2025", Status: engine.StatusResolved,
			ExactEvidence: []engine.EvidenceLink{link},
		}},
		Provenance: engine.Provenance{EngineVersion: "synthetic-test", BundleID: "demo-fixture"},
	}, nil
}

func TestCanonicalVerifierNeedsIndependentEvidencePolicy(t *testing.T) {
	policy := engine.SourcePolicy{Rules: []engine.SourcePolicyRule{{
		ID: "demo-policy", ClaimKind: "code_adoption", JurisdictionID: "DEMO-XX",
		DocumentID: "demo-instrument", ArtifactSHA256: strings.Repeat("a", 64),
		TextSHA256: strings.Repeat("b", 64), LegalLocator: "DEMO-XX 1.1",
		EffectiveFrom: "2025-01-01",
	}}}
	req := engine.ProjectRequest{
		ProjectID: "DEMO-1", Address: "1 Synthetic Lane, DEMO-XX",
		ApplicabilityDate: "2026-10-25",
	}
	for _, tt := range []struct {
		name   string
		policy engine.SourcePolicy
		want   engine.ProjectVerdict
	}{
		{"independent-source-policy", policy, engine.ProjectVerified},
		{"provider-does-not-authorize-itself", engine.SourcePolicy{}, ""},
		{"wrong-document-policy", engine.SourcePolicy{Rules: []engine.SourcePolicyRule{{
			ID: "wrong", ClaimKind: "code_adoption", JurisdictionID: "DEMO-XX",
			DocumentID: "unrelated-document",
		}}}, ""},
	} {
		t.Run(tt.name, func(t *testing.T) {
			verifier := engine.NewProjectVerifier(engine.NewRuntimeWithPolicies(
				syntheticEvidenceProvider{}, tt.policy, syntheticPrivateSitePolicyFor("1 Synthetic Lane, DEMO-XX", "DEMO", "DEMO-XX")))
			result, err := verifier.VerifyProject(context.Background(), req)
			if tt.want == "" {
				if err == nil {
					t.Fatalf("unsupported claim promoted: %#v", result)
				}
				return
			}
			if err != nil || result.Verdict != tt.want {
				payload, _ := json.Marshal(result)
				t.Fatalf("canonical result rejected: %s %v", payload, err)
			}
		})
	}
}
