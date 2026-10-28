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

func syntheticPrivateSitePolicyFor(address, country, jurisdiction string) engine.SiteAuthorityPolicy {
	normalized, err := engine.NormalizeQuery(engine.Query{Address: address, ApplicabilityDate: "2026-10-25"})
	if err != nil {
		panic(err)
	}
	return engine.SiteAuthorityPolicy{
		SchemaVersion: "1.0.0", Rules: []engine.SiteAuthorityRule{{
			ID: "fictional-site-witness",
			LocationSHA256: engine.SiteLocationSHA256(normalized),
			LocationPrecision: "address_point", CountryCode: country,
			JurisdictionID: jurisdiction, AuthorityID: "fictional-authority",
			EffectiveFrom: "2026-01-01", EffectiveTo: "2026-12-31",
			ReviewState: "reviewed",
			Boundary: engine.SiteAuthorityBoundary{
				ObjectURI: "gs://fictional-review-only/demo-geometry",
				ObjectGeneration: "11",
				SHA256: strings.Repeat("a", 64), CRS: "EPSG:4326",
				SiteRelation: "interior_reviewed",
				TopologyReviewSHA: strings.Repeat("b", 64),
				CrosswalkReviewSHA: strings.Repeat("c", 64),
			},
			Legal: engine.SiteAuthorityLegal{
				InstrumentID: "fictional-ordinance",
				SourceArtifactSHA256: strings.Repeat("d", 64),
				ExactTextSHA256: strings.Repeat("e", 64), LegalLocator: "Fictional section",
				EffectiveFrom: "2026-01-01", EffectiveTo: "2026-12-31",
				EnforcementAuthority: "fictional-authority", EnforcementMode: "direct",
			},
			Provenance: engine.SiteAuthorityProvenance{
				EntityURI: "urn:fictional:source",
				ActivityURI: "urn:fictional:review",
				AgentURI: "urn:fictional:reviewer",
				ReviewReceiptSHA256: strings.Repeat("f", 64),
			},
		}},
	}
}

func TestOperatorPinnedSiteAuthorityPolicyRejectsDriftAndFalseReview(t *testing.T) {
	policy := syntheticPrivateSitePolicyFor(demoAddress, "US", "DEMO-XX")
	path := filepath.Join(t.TempDir(), "protected-site-review.json")
	encode := func(value engine.SiteAuthorityPolicy) []byte {
		t.Helper()
		data, err := json.Marshal(value)
		if err != nil { t.Fatal(err) }
		return data
	}
	load := func(data []byte, shouldPass bool) {
		t.Helper()
		if err := os.WriteFile(path, data, 0600); err != nil { t.Fatal(err) }
		hash := sha256.Sum256(data)
		out, err := loadPinnedSiteAuthorityPolicy(path, hex.EncodeToString(hash[:]))
		if shouldPass {
			if err != nil || len(out.Rules) != 1 {
				t.Fatalf("synthetic pinned site review wrongly rejected: policy=%+v err=%v", out, err)
			}
		} else if err == nil {
			t.Fatalf("unreviewed policy accepted: %+v", out)
		}
	}
	original := encode(policy)
	load(original, true)
	if err := os.WriteFile(path, append(append([]byte{}, original...), byte(' ')), 0600); err != nil { t.Fatal(err) }
	h := sha256.Sum256(original)
	if _, err := loadPinnedSiteAuthorityPolicy(path, hex.EncodeToString(h[:])); err == nil {
		t.Fatal("unreviewed mutated policy passed independently pinned SHA")
	}
	changes := []struct {
		name string
		edit func(*engine.SiteAuthorityPolicy)
	}{
		{"missing review receipt", func(p *engine.SiteAuthorityPolicy) { p.Rules[0].Provenance.ReviewReceiptSHA256 = "" }},
		{"wrong review state", func(p *engine.SiteAuthorityPolicy) { p.Rules[0].ReviewState = "candidate_only" }},
		{"interpolated geocode is not legal site", func(p *engine.SiteAuthorityPolicy) { p.Rules[0].LocationPrecision = "interpolated" }},
		{"unreviewed boundary", func(p *engine.SiteAuthorityPolicy) { p.Rules[0].Boundary.SiteRelation = "overlaps" }},
		{"missing GCS generation", func(p *engine.SiteAuthorityPolicy) { p.Rules[0].Boundary.ObjectGeneration = "" }},
		{"expired interval", func(p *engine.SiteAuthorityPolicy) {
			p.Rules[0].EffectiveFrom, p.Rules[0].EffectiveTo = "2027-01-01", "2026-01-01"
		}},
		{"wrong enforcement authority", func(p *engine.SiteAuthorityPolicy) { p.Rules[0].Legal.EnforcementAuthority = "unknown-office" }},
		{"missing delegation instrument", func(p *engine.SiteAuthorityPolicy) { p.Rules[0].Legal.EnforcementMode = "delegated" }},
		{"duplicate rule", func(p *engine.SiteAuthorityPolicy) { p.Rules = append(p.Rules, p.Rules[0]) }},
	}
	for _, test := range changes {
		t.Run(test.name, func(t *testing.T) {
			copyPolicy := engine.SiteAuthorityPolicy{SchemaVersion: policy.SchemaVersion, Rules: append([]engine.SiteAuthorityRule(nil), policy.Rules...)}
			test.edit(&copyPolicy)
			load(encode(copyPolicy), false)
		})
	}
	load(append(append([]byte{}, original...), []byte("{}")...), false)
	load([]byte(strings.Replace(string(original), `"review_state":"reviewed"`, `"review_state":"reviewed","review_state":"reviewed"`, 1)), false)
	load([]byte(strings.Replace(string(original), `"legal_locator":"Fictional section"`, `"legal_locator":"Fictional section","unexpected":"fake"`, 1)), false)
	if _, err := loadPinnedSiteAuthorityPolicy(path, "not-a-hash"); err == nil {
		t.Fatal("malformed caller-supplied digest accepted")
	}
}

func TestConfiguredSiteAuthorityDoesNotTrustProviderClaims(t *testing.T) {
	provider := syntheticEvidenceProvider{}
	source := engine.SourcePolicy{Rules: []engine.SourcePolicyRule{{
		ID: "synthetic-adoption", ClaimKind: "code_adoption", JurisdictionID: "DEMO-XX",
		DocumentID: "demo-instrument", ArtifactSHA256: strings.Repeat("a", 64),
		TextSHA256: strings.Repeat("b", 64), LegalLocator: "DEMO-XX 1.1",
		EffectiveFrom: "2026-01-01", EffectiveTo: "2026-12-31",
	}}}
	request := engine.ProjectRequest{ProjectID: "fictional-site", Address: "1 Synthetic Lane, DEMO-XX", ApplicabilityDate: "2026-10-25"}
	without, err := engine.NewProjectVerifier(engine.NewRuntimeWithSourcePolicy(provider, source)).VerifyProject(context.Background(), request)
	if err != nil || without.Verdict != engine.ProjectNotVerified ||
		without.SiteAuthority == nil || without.SiteAuthority.Status != "unresolved" {
		t.Fatalf("provider self-authorized a project: %+v %v", without, err)
	}
	site := syntheticPrivateSitePolicyFor(request.Address, "DEMO", "DEMO-XX")
	with, err := engine.NewProjectVerifier(engine.NewRuntimeWithPolicies(provider, source, site)).
		VerifyProject(context.Background(), request)
	if err != nil || with.Verdict != engine.ProjectVerified || with.SiteAuthority == nil ||
		with.SiteAuthority.Status != "reviewed" {
		t.Fatalf("fictional independent site witness and adopted-code proof were not composed: %+v %v", with, err)
	}
}
