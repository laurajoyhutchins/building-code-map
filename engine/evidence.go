package engine

import (
	"reflect"
	"regexp"
	"sort"
	"strings"
	"time"
)

type EvidenceRole string

const (
	EvidenceEstablishes  EvidenceRole = "establishes"
	EvidenceCorroborates EvidenceRole = "corroborates"
	EvidenceContradicts  EvidenceRole = "contradicts"
	EvidenceDefines      EvidenceRole = "defines"
	EvidenceIncorporates EvidenceRole = "incorporates"
	EvidenceSupersedes   EvidenceRole = "supersedes"
)

var allowedEvidenceRoles = map[EvidenceRole]struct{}{
	EvidenceEstablishes: {}, EvidenceCorroborates: {}, EvidenceContradicts: {},
	EvidenceDefines: {}, EvidenceIncorporates: {}, EvidenceSupersedes: {},
}

var sha256Pattern = regexp.MustCompile(`^[0-9a-f]{64}$`)

type SourceDocument struct {
	ID             string `json:"id"`
	Kind           string `json:"kind"`
	JurisdictionID string `json:"jurisdiction_id"`
	Title          string `json:"title,omitempty"`
	CanonicalURL   string `json:"canonical_url,omitempty"`
}

type SourceArtifact struct {
	ID         string `json:"id"`
	DocumentID string `json:"document_id"`
	SHA256     string `json:"sha256"`
	SizeBytes  int64  `json:"size_bytes"`
	SourceURL  string `json:"source_url,omitempty"`
	ObservedAt string `json:"observed_at,omitempty"`
}

type TextAnchor struct {
	ID           string `json:"id"`
	ArtifactID   string `json:"artifact_id"`
	LegalLocator string `json:"legal_locator,omitempty"`
	StartOffset  *int64 `json:"start_offset,omitempty"`
	EndOffset    *int64 `json:"end_offset,omitempty"`
	TextSHA256   string `json:"text_sha256"`
}

type EvidenceLink struct {
	ID           string         `json:"id"`
	ClaimID      string         `json:"claim_id"`
	Role         EvidenceRole   `json:"role"`
	Document     SourceDocument `json:"document"`
	Artifact     SourceArtifact `json:"artifact"`
	Anchor       TextAnchor     `json:"anchor"`
	TargetLinkID string         `json:"target_link_id,omitempty"`
}

type SourcePolicyRule struct {
	ID             string `json:"id"`
	ClaimKind      string `json:"claim_kind"`
	JurisdictionID string `json:"jurisdiction_id"`
	DocumentID     string `json:"document_id"`
	ArtifactSHA256 string `json:"artifact_sha256"`
	TextSHA256     string `json:"text_sha256"`
	LegalLocator   string `json:"legal_locator"`
	EffectiveFrom  string `json:"effective_from,omitempty"`
	EffectiveTo    string `json:"effective_to,omitempty"`
}

type SourcePolicy struct {
	Rules []SourcePolicyRule `json:"rules"`
}

type EvidenceClaim struct {
	ID              string         `json:"id"`
	Kind            string         `json:"kind"`
	JurisdictionID  string         `json:"jurisdiction_id"`
	AsOf            string         `json:"as_of"`
	LegallyMaterial bool           `json:"legally_material"`
	Links           []EvidenceLink `json:"links"`
}

type EvidenceDecision string

const (
	EvidenceVerified     EvidenceDecision = "verified"
	EvidenceInsufficient EvidenceDecision = "insufficient_evidence"
	EvidenceConflicting  EvidenceDecision = "conflicting"
	EvidenceSuperseded   EvidenceDecision = "superseded"
)

type EvidenceVerification struct {
	Decision             EvidenceDecision `json:"decision"`
	QualifyingLinkIDs    []string         `json:"qualifying_link_ids,omitempty"`
	DisqualifyingLinkIDs []string         `json:"disqualifying_link_ids,omitempty"`
}

func VerifyClaimEvidence(claim EvidenceClaim, policy SourcePolicy) (EvidenceVerification, error) {
	if strings.TrimSpace(claim.ID) == "" || strings.TrimSpace(claim.Kind) == "" || strings.TrimSpace(claim.JurisdictionID) == "" {
		return EvidenceVerification{}, EngineError{Code: ErrorDataBundleInvalid, Message: "evidence claim requires id, kind, and jurisdiction_id"}
	}
	asOf, err := time.Parse(time.DateOnly, strings.TrimSpace(claim.AsOf))
	if err != nil {
		return EvidenceVerification{}, EngineError{Code: ErrorDataBundleInvalid, Message: "evidence claim as_of must use YYYY-MM-DD"}
	}
	if !claim.LegallyMaterial {
		return EvidenceVerification{Decision: EvidenceVerified}, nil
	}

	qualifying := make([]string, 0)
	contradicting := make([]string, 0)
	supersededTargets := map[string]struct{}{}
	linksByID := map[string]EvidenceLink{}

	for _, link := range claim.Links {
		if err := validateEvidenceLink(link, claim.ID); err != nil {
			return EvidenceVerification{}, err
		}
		if err := acceptEvidenceLinkIdentity(linksByID, link); err != nil {
			return EvidenceVerification{}, err
		}
		if link.Role == EvidenceContradicts {
			contradicting = append(contradicting, link.ID)
		}
		if link.Role == EvidenceSupersedes && strings.TrimSpace(link.TargetLinkID) != "" {
			supersededTargets[link.TargetLinkID] = struct{}{}
		}
	}

	if len(contradicting) > 0 {
		sort.Strings(contradicting)
		return EvidenceVerification{Decision: EvidenceConflicting, DisqualifyingLinkIDs: contradicting}, nil
	}

	for _, link := range claim.Links {
		if link.Role != EvidenceEstablishes {
			continue
		}
		if _, superseded := supersededTargets[link.ID]; superseded {
			continue
		}
		if policyAllows(policy, claim.Kind, claim.JurisdictionID, asOf, link) {
			qualifying = append(qualifying, link.ID)
		}
	}

	if len(qualifying) > 0 {
		sort.Strings(qualifying)
		return EvidenceVerification{Decision: EvidenceVerified, QualifyingLinkIDs: qualifying}, nil
	}

	if len(supersededTargets) > 0 {
		targets := make([]string, 0, len(supersededTargets))
		for target := range supersededTargets {
			if _, exists := linksByID[target]; exists {
				targets = append(targets, target)
			}
		}
		sort.Strings(targets)
		if len(targets) > 0 {
			return EvidenceVerification{Decision: EvidenceSuperseded, DisqualifyingLinkIDs: targets}, nil
		}
	}

	return EvidenceVerification{Decision: EvidenceInsufficient}, nil
}

// acceptEvidenceLinkIdentity prevents one identifier from silently referring
// to different claims, roles or source records. Repeated identical records
// are harmless; conflicting records are an invalid evidence bundle.
func acceptEvidenceLinkIdentity(seen map[string]EvidenceLink, link EvidenceLink) error {
	if previous, exists := seen[link.ID]; exists && !reflect.DeepEqual(previous, link) {
		return EngineError{Code: ErrorDataBundleInvalid, Message: "evidence link id refers to conflicting records"}
	}
	seen[link.ID] = link
	return nil
}

func validateEvidenceLink(link EvidenceLink, claimID string) error {
	if strings.TrimSpace(link.ID) == "" || strings.TrimSpace(link.ClaimID) != claimID {
		return EngineError{Code: ErrorDataBundleInvalid, Message: "evidence link requires a stable id and matching claim_id"}
	}
	if _, ok := allowedEvidenceRoles[link.Role]; !ok {
		return EngineError{Code: ErrorDataBundleInvalid, Message: "evidence link role is invalid"}
	}
	if strings.TrimSpace(link.Document.ID) == "" || strings.TrimSpace(link.Document.Kind) == "" || strings.TrimSpace(link.Document.JurisdictionID) == "" {
		return EngineError{Code: ErrorDataBundleInvalid, Message: "source document requires id, kind, and jurisdiction_id"}
	}
	if strings.TrimSpace(link.Artifact.ID) == "" || strings.TrimSpace(link.Artifact.DocumentID) != link.Document.ID {
		return EngineError{Code: ErrorDataBundleInvalid, Message: "source artifact must identify its source document"}
	}
	if !sha256Pattern.MatchString(strings.ToLower(strings.TrimSpace(link.Artifact.SHA256))) || link.Artifact.SizeBytes <= 0 {
		return EngineError{Code: ErrorDataBundleInvalid, Message: "source artifact requires immutable sha256 and positive size_bytes"}
	}
	if strings.TrimSpace(link.Anchor.ID) == "" || strings.TrimSpace(link.Anchor.ArtifactID) != link.Artifact.ID || !sha256Pattern.MatchString(strings.ToLower(strings.TrimSpace(link.Anchor.TextSHA256))) {
		return EngineError{Code: ErrorDataBundleInvalid, Message: "text anchor requires artifact identity and exact text sha256"}
	}
	legalLocator := strings.TrimSpace(link.Anchor.LegalLocator)
	hasOffsets := link.Anchor.StartOffset != nil || link.Anchor.EndOffset != nil
	if legalLocator == "" && !hasOffsets {
		return EngineError{Code: ErrorDataBundleInvalid, Message: "text anchor requires a legal locator or machine offsets"}
	}
	if hasOffsets {
		if link.Anchor.StartOffset == nil || link.Anchor.EndOffset == nil || *link.Anchor.StartOffset < 0 || *link.Anchor.EndOffset <= *link.Anchor.StartOffset {
			return EngineError{Code: ErrorDataBundleInvalid, Message: "text anchor offsets must be a valid half-open range"}
		}
	}
	return nil
}

func policyAllows(policy SourcePolicy, claimKind, jurisdictionID string, asOf time.Time, link EvidenceLink) bool {
	for _, rule := range policy.Rules {
		if strings.TrimSpace(rule.ClaimKind) != claimKind ||
			strings.TrimSpace(rule.JurisdictionID) != jurisdictionID ||
			strings.TrimSpace(link.Document.JurisdictionID) != jurisdictionID ||
			strings.TrimSpace(rule.DocumentID) != link.Document.ID ||
			!sha256Pattern.MatchString(rule.ArtifactSHA256) ||
			!sha256Pattern.MatchString(rule.TextSHA256) ||
			strings.TrimSpace(rule.LegalLocator) == "" ||
			rule.ArtifactSHA256 != link.Artifact.SHA256 ||
			rule.TextSHA256 != link.Anchor.TextSHA256 ||
			rule.LegalLocator != link.Anchor.LegalLocator {
			continue
		}
		if rule.EffectiveFrom != "" {
			from, err := time.Parse(time.DateOnly, rule.EffectiveFrom)
			if err != nil || asOf.Before(from) {
				continue
			}
		}
		if rule.EffectiveTo != "" {
			to, err := time.Parse(time.DateOnly, rule.EffectiveTo)
			if err != nil || asOf.After(to) {
				continue
			}
		}
		return true
	}
	return false
}
