package engine

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"regexp"
	"strings"
	"time"
)

// SiteAuthorityPolicy is supplied by the operator independently of the
// Provider and its claims. Configuring a digest-shaped rule is NOT itself
// evidence of real source custody, a reviewed law, or GCS readback.
// Admission of these inputs remains an independent human/source operation.
type SiteAuthorityPolicy struct {
	SchemaVersion string              `json:"schema_version"`
	Rules         []SiteAuthorityRule `json:"rules"`
}

type SiteAuthorityRule struct {
	ID                string                  `json:"id"`
	LocationSHA256    string                  `json:"location_sha256"`
	LocationPrecision string                  `json:"location_precision"`
	CountryCode       string                  `json:"country_code"`
	JurisdictionID    string                  `json:"jurisdiction_id"`
	AuthorityID       string                  `json:"authority_id"`
	EffectiveFrom     string                  `json:"effective_from"`
	EffectiveTo       string                  `json:"effective_to"`
	ReviewState       string                  `json:"review_state"`
	Boundary          SiteAuthorityBoundary   `json:"boundary"`
	Legal             SiteAuthorityLegal      `json:"legal"`
	Provenance        SiteAuthorityProvenance `json:"provenance"`
}

type SiteAuthorityBoundary struct {
	ObjectURI          string `json:"object_uri"`
	ObjectGeneration   string `json:"object_generation"`
	SHA256             string `json:"sha256"`
	CRS                string `json:"crs"`
	SiteRelation       string `json:"site_relation"`
	TopologyReviewSHA  string `json:"topology_review_sha256"`
	CrosswalkReviewSHA string `json:"crosswalk_review_sha256"`
}

type SiteAuthorityLegal struct {
	InstrumentID          string `json:"instrument_id"`
	SourceArtifactSHA256  string `json:"source_artifact_sha256"`
	ExactTextSHA256       string `json:"exact_text_sha256"`
	LegalLocator          string `json:"legal_locator"`
	EffectiveFrom         string `json:"effective_from"`
	EffectiveTo           string `json:"effective_to"`
	EnforcementAuthority  string `json:"enforcement_authority_id"`
	EnforcementMode       string `json:"enforcement_mode"`
	DelegationArtifactSHA string `json:"delegation_artifact_sha256,omitempty"`
}

type SiteAuthorityProvenance struct {
	EntityURI           string `json:"entity_uri"`
	ActivityURI         string `json:"activity_uri"`
	AgentURI            string `json:"agent_uri"`
	ReviewReceiptSHA256 string `json:"review_receipt_sha256"`
}

// SiteAuthorityAssessment is safe, bounded project-level provenance.
// It never includes a private legal text, source URL, or bulk geometry.
type SiteAuthorityAssessment struct {
	Status                 string `json:"status"`
	Reason                 string `json:"reason"`
	RuleID                 string `json:"rule_id,omitempty"`
	ProvenanceEntity       string `json:"provenance_entity,omitempty"`
	ReviewActivity         string `json:"review_activity,omitempty"`
	ReviewReceiptSHA256    string `json:"review_receipt_sha256,omitempty"`
	BoundaryGeneration     string `json:"boundary_generation,omitempty"`
	LegalVerdictAuthorized bool   `json:"legal_verdict_authorized"`
}

var positiveGCSGeneration = regexp.MustCompile(`^[1-9][0-9]*$`)
var siteAuthorityGCSObject = regexp.MustCompile(`^gs://[^/]+/.+$`)

func siteAuthorityURI(value string) bool {
	return strings.HasPrefix(value, "urn:") && len(value) > 4 ||
		strings.HasPrefix(value, "https://") && len(value) > 8
}

func siteAuthorityDate(value string) (time.Time, bool) {
	parsed, err := time.Parse(time.DateOnly, value)
	return parsed, err == nil && parsed.Format(time.DateOnly) == value
}

func validSiteAuthorityInterval(from, to string) bool {
	start, ok := siteAuthorityDate(from)
	if !ok {
		return false
	}
	end, ok := siteAuthorityDate(to)
	return ok && !end.Before(start)
}

// ValidateSiteAuthorityRule only establishes structural and identity fitness.
// It cannot independently inspect external objects, issue rights, or the
// reviewing agent's authority. Only independently admitted, digest-pinned
// operator rules may reach this predicate.
func ValidateSiteAuthorityRule(rule SiteAuthorityRule) bool {
	if strings.TrimSpace(rule.ID) == "" ||
		!sha256Pattern.MatchString(rule.LocationSHA256) ||
		strings.TrimSpace(rule.CountryCode) == "" ||
		strings.TrimSpace(rule.JurisdictionID) == "" ||
		strings.TrimSpace(rule.AuthorityID) == "" ||
		rule.ReviewState != "reviewed" ||
		(rule.LocationPrecision != "address_point" && rule.LocationPrecision != "parcel") ||
		!validSiteAuthorityInterval(rule.EffectiveFrom, rule.EffectiveTo) {
		return false
	}
	b := rule.Boundary
	if !siteAuthorityGCSObject.MatchString(b.ObjectURI) ||
		!positiveGCSGeneration.MatchString(b.ObjectGeneration) ||
		!sha256Pattern.MatchString(b.SHA256) ||
		!sha256Pattern.MatchString(b.TopologyReviewSHA) ||
		!sha256Pattern.MatchString(b.CrosswalkReviewSHA) ||
		b.CRS != "EPSG:4326" || b.SiteRelation != "interior_reviewed" {
		return false
	}
	l := rule.Legal
	if strings.TrimSpace(l.InstrumentID) == "" ||
		!sha256Pattern.MatchString(l.SourceArtifactSHA256) ||
		!sha256Pattern.MatchString(l.ExactTextSHA256) ||
		strings.TrimSpace(l.LegalLocator) == "" ||
		!validSiteAuthorityInterval(l.EffectiveFrom, l.EffectiveTo) ||
		l.EnforcementAuthority != rule.AuthorityID {
		return false
	}
	if l.EnforcementMode != "direct" && l.EnforcementMode != "delegated" {
		return false
	}
	if l.EnforcementMode == "delegated" && !sha256Pattern.MatchString(l.DelegationArtifactSHA) {
		return false
	}
	if l.EnforcementMode == "direct" && l.DelegationArtifactSHA != "" {
		return false
	}
	p := rule.Provenance
	return siteAuthorityURI(p.EntityURI) && siteAuthorityURI(p.ActivityURI) &&
		siteAuthorityURI(p.AgentURI) && sha256Pattern.MatchString(p.ReviewReceiptSHA256)
}

// SiteLocationSHA256 binds a trusted geographic witness to the exact
// normalized request site, not a free-form provider-reported location.
// Project facts, code family, and date are assessed independently.
func SiteLocationSHA256(query NormalizedQuery) string {
	binding := struct {
		Version string `json:"version"`
		Address string `json:"address,omitempty"`
		Point   *Point `json:"point,omitempty"`
	}{
		Version: "bcm-site-location-v1", Address: query.Address, Point: query.Point,
	}
	raw, _ := json.Marshal(binding)
	digest := sha256.Sum256(raw)
	return hex.EncodeToString(digest[:])
}

func unresolvedSiteAuthority(reason string) SiteAuthorityAssessment {
	return SiteAuthorityAssessment{Status: "unresolved", Reason: reason}
}

// evaluateSiteAuthority considers *all* contemporaneous witnesses for the
// same request site. A provider cannot choose just the favorable source.
// Contradictions and two enforcement bodies fail closed.
func evaluateSiteAuthority(query NormalizedQuery, jurisdiction JurisdictionContext, policy SiteAuthorityPolicy) SiteAuthorityAssessment {
	if policy.SchemaVersion != "1.0.0" {
		return unresolvedSiteAuthority("independent_site_authority_policy_missing")
	}
	digest := SiteLocationSHA256(query)
	asOf, ok := siteAuthorityDate(query.ApplicabilityDate)
	if !ok {
		return unresolvedSiteAuthority("site_authority_date_invalid")
	}
	matchingSite := 0
	active := make([]SiteAuthorityRule, 0)
	for _, rule := range policy.Rules {
		if rule.LocationSHA256 != digest {
			continue
		}
		matchingSite++
		if !ValidateSiteAuthorityRule(rule) {
			return unresolvedSiteAuthority("site_authority_review_or_source_incomplete")
		}
		from, _ := siteAuthorityDate(rule.EffectiveFrom)
		to, _ := siteAuthorityDate(rule.EffectiveTo)
		if asOf.Before(from) || asOf.After(to) {
			continue
		}
		legalFrom, _ := siteAuthorityDate(rule.Legal.EffectiveFrom)
		legalTo, _ := siteAuthorityDate(rule.Legal.EffectiveTo)
		if asOf.Before(legalFrom) || asOf.After(legalTo) {
			continue
		}
		active = append(active, rule)
	}
	if len(active) > 1 {
		return SiteAuthorityAssessment{Status: "conflicting", Reason: "multiple_site_authorities_or_overlapping_boundaries"}
	}
	if len(active) == 0 {
		if matchingSite > 0 {
			return unresolvedSiteAuthority("site_authority_not_effective_for_project_date")
		}
		return unresolvedSiteAuthority("independent_site_to_jurisdiction_evidence_missing")
	}
	rule := active[0]
	if rule.JurisdictionID != jurisdiction.PrimaryJurisdictionID ||
		rule.CountryCode != jurisdiction.CountryCode {
		return unresolvedSiteAuthority("site_authority_government_identity_mismatch")
	}
	return SiteAuthorityAssessment{
		Status: "reviewed", Reason: "site_authority_policy_witness_matched",
		RuleID:              rule.ID,
		ProvenanceEntity:    rule.Provenance.EntityURI,
		ReviewActivity:      rule.Provenance.ActivityURI,
		ReviewReceiptSHA256: rule.Provenance.ReviewReceiptSHA256,
		BoundaryGeneration:  rule.Boundary.ObjectGeneration,
		// This is never a grant of legal authority from the provider's own
		// geometry/code. Project verification additionally requires all
		// exact adopted-code SourcePolicy evidence.
		LegalVerdictAuthorized: false,
	}
}
