package engine

import (
	"context"
	"strings"
	"testing"
)

func syntheticSiteAuthorityPolicy(address, country, jurisdiction string) SiteAuthorityPolicy {
	normalized, err := NormalizeQuery(Query{Address: address, ApplicabilityDate: "2026-08-25"})
	if err != nil {
		panic(err)
	}
	return SiteAuthorityPolicy{
		SchemaVersion: "1.0.0",
		Rules: []SiteAuthorityRule{{
			ID: "synthetic-site-1", LocationSHA256: SiteLocationSHA256(normalized),
			LocationPrecision: "address_point", CountryCode: country,
			JurisdictionID: jurisdiction, AuthorityID: "synthetic-local-enforcer",
			EffectiveFrom: "2026-01-01", EffectiveTo: "2026-12-31",
			ReviewState: "reviewed",
			Boundary: SiteAuthorityBoundary{
				ObjectURI: "gs://synthetic-review-only/fictional-boundary",
				ObjectGeneration: "1", SHA256: strings.Repeat("a", 64),
				CRS: "EPSG:4326", SiteRelation: "interior_reviewed",
				TopologyReviewSHA: strings.Repeat("b", 64),
				CrosswalkReviewSHA: strings.Repeat("c", 64),
			},
			Legal: SiteAuthorityLegal{
				InstrumentID: "synthetic-annexation", SourceArtifactSHA256: strings.Repeat("d", 64),
				ExactTextSHA256: strings.Repeat("e", 64),
				LegalLocator: "Synthetic locality clause 1", EffectiveFrom: "2026-01-01",
				EffectiveTo: "2026-12-31", EnforcementAuthority: "synthetic-local-enforcer",
				EnforcementMode: "direct",
			},
			Provenance: SiteAuthorityProvenance{
				EntityURI: "urn:synthetic:source", ActivityURI: "urn:synthetic:review",
				AgentURI: "urn:synthetic:reviewer", ReviewReceiptSHA256: strings.Repeat("f", 64),
			},
		}},
	}
}

// Fixed fictional proof is used only in tests. Real City/County legal authority,
// source bytes, GCS generations and independent reviewer signatures are absent.
func TestIndependentSiteAuthorityRequiresSourceAndGeographicEvidence(t *testing.T) {
	const address = "100 Demo Plaza"
	base := syntheticSiteAuthorityPolicy(address, "US", "DEMO-XX")
	normalized, _ := NormalizeQuery(Query{Address: address, ApplicabilityDate: "2026-08-25"})
	jurisdiction := JurisdictionContext{CountryCode: "US", PrimaryJurisdictionID: "DEMO-XX"}
	if got := evaluateSiteAuthority(normalized, jurisdiction, base); got.Status != "reviewed" ||
		got.LegalVerdictAuthorized || got.ProvenanceEntity != "urn:synthetic:source" {
		t.Fatalf("reviewed synthetic site witness rejected or promoted: %+v", got)
	}

	cases := []struct {
		name string
		mutate func(*SiteAuthorityPolicy)
		status string
		reason string
	}{
		{"no independent rules", func(p *SiteAuthorityPolicy) { p.Rules = nil }, "unresolved", "missing"},
		{"candidate place code is not legal identity", func(p *SiteAuthorityPolicy) {
			p.Rules[0].ReviewState = "candidate_only"
		}, "unresolved", "incomplete"},
		{"interpolated geocode", func(p *SiteAuthorityPolicy) {
			p.Rules[0].LocationPrecision = "interpolated"
		}, "unresolved", "incomplete"},
		{"unresolved municipal boundary edge", func(p *SiteAuthorityPolicy) {
			p.Rules[0].Boundary.SiteRelation = "boundary_edge"
		}, "unresolved", "incomplete"},
		{"missing topology review", func(p *SiteAuthorityPolicy) {
			p.Rules[0].Boundary.TopologyReviewSHA = ""
		}, "unresolved", "incomplete"},
		{"missing Census disagreement review", func(p *SiteAuthorityPolicy) {
			p.Rules[0].Boundary.CrosswalkReviewSHA = ""
		}, "unresolved", "incomplete"},
		{"unobserved GCS generation", func(p *SiteAuthorityPolicy) {
			p.Rules[0].Boundary.ObjectGeneration = ""
		}, "unresolved", "incomplete"},
		{"fake GCS generation zero", func(p *SiteAuthorityPolicy) {
			p.Rules[0].Boundary.ObjectGeneration = "0"
		}, "unresolved", "incomplete"},
		{"annexation after project", func(p *SiteAuthorityPolicy) {
			p.Rules[0].Legal.EffectiveFrom = "2027-01-01"
			p.Rules[0].Legal.EffectiveTo = "2027-12-31"
		}, "unresolved", "date"},
		{"expired site review", func(p *SiteAuthorityPolicy) {
			p.Rules[0].EffectiveTo = "2026-01-02"
		}, "unresolved", "date"},
		{"missing legal instrument source", func(p *SiteAuthorityPolicy) {
			p.Rules[0].Legal.SourceArtifactSHA256 = ""
		}, "unresolved", "incomplete"},
		{"invalid legal effective interval", func(p *SiteAuthorityPolicy) {
			p.Rules[0].Legal.EffectiveTo = "2025-12-31"
		}, "unresolved", "incomplete"},
		{"wrong enforcing body", func(p *SiteAuthorityPolicy) {
			p.Rules[0].Legal.EnforcementAuthority = "another-office"
		}, "unresolved", "incomplete"},
		{"unsigned delegated authority", func(p *SiteAuthorityPolicy) {
			p.Rules[0].Legal.EnforcementMode = "delegated"
		}, "unresolved", "incomplete"},
		{"direct authority cannot smuggle delegation", func(p *SiteAuthorityPolicy) {
			p.Rules[0].Legal.DelegationArtifactSHA = strings.Repeat("a", 64)
		}, "unresolved", "incomplete"},
		{"no independent reviewing agent", func(p *SiteAuthorityPolicy) {
			p.Rules[0].Provenance.AgentURI = ""
		}, "unresolved", "incomplete"},
		{"no independent review receipt", func(p *SiteAuthorityPolicy) {
			p.Rules[0].Provenance.ReviewReceiptSHA256 = ""
		}, "unresolved", "incomplete"},
		{"wrong government from provider", func(p *SiteAuthorityPolicy) {
			p.Rules[0].JurisdictionID = "DEMO-NEIGHBOR"
		}, "unresolved", "mismatch"},
		{"two enforcing bodies at same site", func(p *SiteAuthorityPolicy) {
			second := p.Rules[0]
			second.ID = "synthetic-site-2"
			second.AuthorityID = "second-enforcer"
			second.Legal.EnforcementAuthority = "second-enforcer"
			p.Rules = append(p.Rules, second)
		}, "conflicting", "multiple"},
		{"Hudson overlapping source code geometries", func(p *SiteAuthorityPolicy) {
			second := p.Rules[0]
			second.ID = "hudson-source-37820"
			second.Boundary.ObjectURI = "gs://synthetic-review-only/hudson-other"
			p.Rules = append(p.Rules, second)
		}, "conflicting", "multiple"},
		{"Central City/Central duplicate jurisdiction", func(p *SiteAuthorityPolicy) {
			second := p.Rules[0]
			second.ID = "central-source-alias"
			second.JurisdictionID = "different-city"
			p.Rules = append(p.Rules, second)
		}, "conflicting", "multiple"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			copyPolicy := SiteAuthorityPolicy{SchemaVersion: base.SchemaVersion, Rules: append([]SiteAuthorityRule(nil), base.Rules...)}
			tc.mutate(&copyPolicy)
			got := evaluateSiteAuthority(normalized, jurisdiction, copyPolicy)
			if got.Status != tc.status || !strings.Contains(got.Reason, tc.reason) || got.LegalVerdictAuthorized {
				t.Fatalf("unreviewed/ambiguous site promoted: %+v", got)
			}
		})
	}
}

// A source-policy-approved adoption cannot establish its own government
// identity. Only operator-reviewed site evidence makes a verified project
// possible, through the existing Runtime -> ProjectVerifier reducer.
func TestCodeAdoptionDoesNotSelfAuthorizeSiteJurisdiction(t *testing.T) {
	query := ProjectRequest{ProjectID: "SYNTHETIC-CO-1", Address: "100 Demo Plaza", ApplicabilityDate: "2026-08-25"}
	provider := syntheticProvider{codes: []CodeResolution{{
		Family: "building", Status: StatusResolved, Edition: "DEMO-2024",
		ExactEvidence: []EvidenceLink{demoEvidenceLink("DEMO-ESTABLISHES", EvidenceEstablishes)},
	}}}
	missing, err := NewProjectVerifier(NewRuntimeWithSourcePolicy(provider, demoPolicy())).VerifyProject(context.Background(), query)
	if err != nil || missing.Verdict != ProjectNotVerified ||
		len(missing.Unresolved) != 1 || missing.Unresolved[0].Kind != RequirementEvidenceDefect ||
		missing.SiteAuthority == nil || missing.SiteAuthority.Status != "unresolved" ||
		len(missing.FollowUpQuestions) != 0 {
		t.Fatalf("code evidence incorrectly granted site jurisdiction: result=%+v err=%v", missing, err)
	}
	approved := syntheticSiteAuthorityPolicy(query.Address, "US", "DEMO-XX")
	pass, err := NewProjectVerifier(NewRuntimeWithPolicies(provider, demoPolicy(), approved)).
		VerifyProject(context.Background(), query)
	if err != nil || pass.Verdict != ProjectVerified ||
		pass.SiteAuthority == nil || pass.SiteAuthority.Status != "reviewed" ||
		pass.SiteAuthority.ProvenanceEntity != "urn:synthetic:source" ||
		pass.SiteAuthority.LegalVerdictAuthorized {
		t.Fatalf("independent synthetic site and code evidence cannot be composed: %+v err=%v", pass, err)
	}
	// A previously approved fictional site must never be applied to a
	// different address, country, coordinate, or project date.
	for _, altered := range []ProjectRequest{
		{ProjectID: "wrong address", Address: "101 Demo Plaza", ApplicabilityDate: "2026-08-25"},
		{ProjectID: "wrong date", Address: "100 Demo Plaza", ApplicabilityDate: "2027-01-01"},
		{ProjectID: "wrong place", Point: &Point{Longitude: -104.99, Latitude: 39.74}, ApplicabilityDate: "2026-08-25"},
	} {
		other, err := NewProjectVerifier(NewRuntimeWithPolicies(provider, demoPolicy(), approved)).
			VerifyProject(context.Background(), altered)
		if altered.ApplicabilityDate == "2027-01-01" {
			// Exact code-adoption policy also expires, and rejects first.
			if err == nil {
				t.Fatalf("expired code evidence unexpectedly valid: %+v", other)
			}
			continue
		}
		if err != nil || other.Verdict != ProjectNotVerified ||
			other.SiteAuthority == nil || other.SiteAuthority.Status == "reviewed" {
			t.Fatalf("old site witness reused for another site: %+v err=%v", other, err)
		}
	}
}
