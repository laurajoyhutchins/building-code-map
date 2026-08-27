package engine

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
	"time"
)

type ProviderConformanceRule struct {
	ID                string   `json:"id"`
	QualityDimensions []string `json:"quality_dimensions"`
	Summary           string   `json:"summary"`
}

type ProviderConformanceDiagnostic struct {
	RuleID  string `json:"rule_id"`
	Path    string `json:"path"`
	Message string `json:"message"`
}

var providerConformanceRules = []ProviderConformanceRule{
	{ID: "BCM-CONF-001", QualityDimensions: []string{"uniqueness"}, Summary: "Required provider object identifiers are unique within their namespace."},
	{ID: "BCM-CONF-002", QualityDimensions: []string{"consistency", "conformity"}, Summary: "Stable provider references resolve to the expected logical object type."},
	{ID: "BCM-CONF-003", QualityDimensions: []string{"completeness", "conformity"}, Summary: "Every adoption identifies an adopting authority."},
	{ID: "BCM-CONF-004", QualityDimensions: []string{"consistency", "timeliness"}, Summary: "Declared effective intervals use valid dates and are not inverted."},
	{ID: "BCM-CONF-005", QualityDimensions: []string{"accuracy", "completeness"}, Summary: "Resolved legally material regulatory claims satisfy BCM claim-evidence requirements."},
	{ID: "BCM-CONF-006", QualityDimensions: []string{"accuracy", "conformity"}, Summary: "Exact evidence satisfies the existing BCM evidence-link validator."},
	{ID: "BCM-CONF-007", QualityDimensions: []string{"accuracy", "completeness"}, Summary: "A not_applicable resolution retains an affirmative basis."},
	{ID: "BCM-CONF-008", QualityDimensions: []string{"consistency"}, Summary: "Evidence defects remain evidence defects and force a not_verified project result."},
	{ID: "BCM-CONF-009", QualityDimensions: []string{"coverage", "completeness"}, Summary: "Machine-checkable contract coverage and limitations agree with represented provider content."},
}

func ProviderConformanceRuleCatalog() []ProviderConformanceRule {
	rules := make([]ProviderConformanceRule, len(providerConformanceRules))
	for index, rule := range providerConformanceRules {
		rules[index] = rule
		rules[index].QualityDimensions = append([]string(nil), rule.QualityDimensions...)
	}
	return rules
}

type ProviderContractDeclaration struct {
	BundleID    string `json:"bundle_id"`
	Limitations string `json:"limitations"`
}

func ParseProviderContractDeclaration(contents []byte) (ProviderContractDeclaration, error) {
	var envelope struct {
		Description struct {
			Limitations string `json:"limitations"`
		} `json:"description"`
		CustomProperties []struct {
			Property string `json:"property"`
			Value    any    `json:"value"`
		} `json:"customProperties"`
	}
	if err := json.Unmarshal(contents, &envelope); err != nil {
		return ProviderContractDeclaration{}, err
	}
	declaration := ProviderContractDeclaration{Limitations: strings.TrimSpace(envelope.Description.Limitations)}
	for _, property := range envelope.CustomProperties {
		if property.Property != "bcmProviderBundleId" {
			continue
		}
		if value, ok := property.Value.(string); ok {
			declaration.BundleID = strings.TrimSpace(value)
		}
	}
	return declaration, nil
}

type ProviderCoverage struct {
	JurisdictionIDs []string `json:"jurisdiction_ids"`
}

type ProviderJurisdiction struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

type ProviderAuthority struct {
	ID             string `json:"id"`
	JurisdictionID string `json:"jurisdiction_id"`
	Name           string `json:"name,omitempty"`
}

type ProviderRegulatoryInstrument struct {
	ID               string `json:"id"`
	JurisdictionID   string `json:"jurisdiction_id"`
	Title            string `json:"title,omitempty"`
	SourceDocumentID string `json:"source_document_id"`
}

type ProviderCodeFamily struct {
	ID   string `json:"id"`
	Name string `json:"name,omitempty"`
}

type ProviderCodeEdition struct {
	ID       string `json:"id"`
	FamilyID string `json:"family_id"`
	Edition  string `json:"edition,omitempty"`
}

type ProviderAdoption struct {
	ID            string `json:"id"`
	AuthorityID   string `json:"authority_id"`
	InstrumentID  string `json:"instrument_id"`
	CodeEditionID string `json:"code_edition_id"`
	EffectiveFrom string `json:"effective_from,omitempty"`
	EffectiveTo   string `json:"effective_to,omitempty"`
}

type ProviderAmendment struct {
	ID            string `json:"id"`
	AuthorityID   string `json:"authority_id"`
	InstrumentID  string `json:"instrument_id"`
	AdoptionID    string `json:"adoption_id"`
	EffectiveFrom string `json:"effective_from,omitempty"`
	EffectiveTo   string `json:"effective_to,omitempty"`
}

type ProviderClaim struct {
	ID              string           `json:"id"`
	JurisdictionID  string           `json:"jurisdiction_id"`
	Kind            string           `json:"kind"`
	SubjectRef      string           `json:"subject_ref"`
	Status          ResolutionStatus `json:"status,omitempty"`
	Basis           string           `json:"basis,omitempty"`
	AsOf            string           `json:"as_of,omitempty"`
	LegallyMaterial bool             `json:"legally_material,omitempty"`
}

type ProviderEvidenceLinkRef struct {
	ID         string       `json:"id"`
	ClaimID    string       `json:"claim_id"`
	DocumentID string       `json:"document_id"`
	ArtifactID string       `json:"artifact_id"`
	AnchorID   string       `json:"anchor_id"`
	Role       EvidenceRole `json:"role"`
}

type ProviderSourcePolicyRule struct {
	ID               string `json:"id"`
	ClaimKind        string `json:"claim_kind,omitempty"`
	JurisdictionID   string `json:"jurisdiction_id,omitempty"`
	DocumentID       string `json:"document_id,omitempty"`
	EffectiveFrom    string `json:"effective_from,omitempty"`
	EffectiveTo      string `json:"effective_to,omitempty"`
	MaterialClass    string `json:"material_class,omitempty"`
	PublicRepository string `json:"public_repository,omitempty"`
}

type ProviderQualityExpectation struct {
	Dimension string `json:"dimension"`
	Statement string `json:"statement"`
}

type ProviderBundle struct {
	ID                    string                       `json:"id"`
	Coverage              ProviderCoverage             `json:"coverage"`
	Jurisdictions         []ProviderJurisdiction       `json:"jurisdictions"`
	Authorities           []ProviderAuthority          `json:"authorities"`
	RegulatoryInstruments []ProviderRegulatoryInstrument `json:"regulatory_instruments"`
	Adoptions             []ProviderAdoption           `json:"adoptions"`
	Amendments            []ProviderAmendment          `json:"amendments"`
	CodeFamilies          []ProviderCodeFamily         `json:"code_families"`
	CodeEditions          []ProviderCodeEdition        `json:"code_editions"`
	SourceDocuments       []SourceDocument             `json:"source_documents"`
	SourceArtifacts       []SourceArtifact             `json:"source_artifacts"`
	TextAnchors           []TextAnchor                 `json:"text_anchors"`
	Claims                []ProviderClaim              `json:"claims"`
	EvidenceLinks         []ProviderEvidenceLinkRef    `json:"evidence_links"`
	SourcePolicyRules     []ProviderSourcePolicyRule   `json:"source_policy_rules"`
	QualityExpectations   []ProviderQualityExpectation `json:"quality_expectations"`
	Limitations           []string                     `json:"limitations"`
}

func conformanceDiagnostic(ruleID, path, message string) ProviderConformanceDiagnostic {
	return ProviderConformanceDiagnostic{RuleID: ruleID, Path: path, Message: message}
}

func sortConformanceDiagnostics(diagnostics []ProviderConformanceDiagnostic) []ProviderConformanceDiagnostic {
	sort.Slice(diagnostics, func(i, j int) bool {
		left := diagnostics[i].RuleID + "\x00" + diagnostics[i].Path + "\x00" + diagnostics[i].Message
		right := diagnostics[j].RuleID + "\x00" + diagnostics[j].Path + "\x00" + diagnostics[j].Message
		return left < right
	})
	return diagnostics
}

func stableIDs(values []string, namespace string) []ProviderConformanceDiagnostic {
	diagnostics := make([]ProviderConformanceDiagnostic, 0)
	seen := map[string]int{}
	for index, value := range values {
		id := strings.TrimSpace(value)
		path := fmt.Sprintf("%s[%d].id", namespace, index)
		if id == "" {
			diagnostics = append(diagnostics, conformanceDiagnostic("BCM-CONF-001", path, "stable id is required"))
			continue
		}
		if previous, ok := seen[id]; ok {
			diagnostics = append(diagnostics, conformanceDiagnostic("BCM-CONF-001", path, fmt.Sprintf("stable id duplicates %s[%d].id", namespace, previous)))
			continue
		}
		seen[id] = index
	}
	return diagnostics
}

func idsOfJurisdictions(values []ProviderJurisdiction) []string { result := make([]string, len(values)); for i := range values { result[i] = values[i].ID }; return result }
func idsOfAuthorities(values []ProviderAuthority) []string { result := make([]string, len(values)); for i := range values { result[i] = values[i].ID }; return result }
func idsOfInstruments(values []ProviderRegulatoryInstrument) []string { result := make([]string, len(values)); for i := range values { result[i] = values[i].ID }; return result }
func idsOfAdoptions(values []ProviderAdoption) []string { result := make([]string, len(values)); for i := range values { result[i] = values[i].ID }; return result }
func idsOfAmendments(values []ProviderAmendment) []string { result := make([]string, len(values)); for i := range values { result[i] = values[i].ID }; return result }
func idsOfFamilies(values []ProviderCodeFamily) []string { result := make([]string, len(values)); for i := range values { result[i] = values[i].ID }; return result }
func idsOfEditions(values []ProviderCodeEdition) []string { result := make([]string, len(values)); for i := range values { result[i] = values[i].ID }; return result }
func idsOfDocuments(values []SourceDocument) []string { result := make([]string, len(values)); for i := range values { result[i] = values[i].ID }; return result }
func idsOfArtifacts(values []SourceArtifact) []string { result := make([]string, len(values)); for i := range values { result[i] = values[i].ID }; return result }
func idsOfAnchors(values []TextAnchor) []string { result := make([]string, len(values)); for i := range values { result[i] = values[i].ID }; return result }
func idsOfClaims(values []ProviderClaim) []string { result := make([]string, len(values)); for i := range values { result[i] = values[i].ID }; return result }
func idsOfLinks(values []ProviderEvidenceLinkRef) []string { result := make([]string, len(values)); for i := range values { result[i] = values[i].ID }; return result }
func idsOfPolicies(values []ProviderSourcePolicyRule) []string { result := make([]string, len(values)); for i := range values { result[i] = values[i].ID }; return result }

func idSet[T any](values []T, id func(T) string) map[string]T {
	result := make(map[string]T, len(values))
	for _, value := range values {
		key := strings.TrimSpace(id(value))
		if key != "" {
			result[key] = value
		}
	}
	return result
}

func missingReference(path, value, expected string) ProviderConformanceDiagnostic {
	return conformanceDiagnostic("BCM-CONF-002", path, fmt.Sprintf("reference %q does not resolve to %s", value, expected))
}

func validateProviderReferences(bundle ProviderBundle) []ProviderConformanceDiagnostic {
	diagnostics := make([]ProviderConformanceDiagnostic, 0)
	jurisdictions := idSet(bundle.Jurisdictions, func(value ProviderJurisdiction) string { return value.ID })
	authorities := idSet(bundle.Authorities, func(value ProviderAuthority) string { return value.ID })
	instruments := idSet(bundle.RegulatoryInstruments, func(value ProviderRegulatoryInstrument) string { return value.ID })
	families := idSet(bundle.CodeFamilies, func(value ProviderCodeFamily) string { return value.ID })
	editions := idSet(bundle.CodeEditions, func(value ProviderCodeEdition) string { return value.ID })
	adoptions := idSet(bundle.Adoptions, func(value ProviderAdoption) string { return value.ID })
	amendments := idSet(bundle.Amendments, func(value ProviderAmendment) string { return value.ID })
	documents := idSet(bundle.SourceDocuments, func(value SourceDocument) string { return value.ID })
	artifacts := idSet(bundle.SourceArtifacts, func(value SourceArtifact) string { return value.ID })
	anchors := idSet(bundle.TextAnchors, func(value TextAnchor) string { return value.ID })
	claims := idSet(bundle.Claims, func(value ProviderClaim) string { return value.ID })

	for index, value := range bundle.Authorities {
		if _, ok := jurisdictions[value.JurisdictionID]; !ok { diagnostics = append(diagnostics, missingReference(fmt.Sprintf("authorities[%d].jurisdiction_id", index), value.JurisdictionID, "jurisdiction")) }
	}
	for index, value := range bundle.RegulatoryInstruments {
		if _, ok := jurisdictions[value.JurisdictionID]; !ok { diagnostics = append(diagnostics, missingReference(fmt.Sprintf("regulatory_instruments[%d].jurisdiction_id", index), value.JurisdictionID, "jurisdiction")) }
		if _, ok := documents[value.SourceDocumentID]; !ok { diagnostics = append(diagnostics, missingReference(fmt.Sprintf("regulatory_instruments[%d].source_document_id", index), value.SourceDocumentID, "source document")) }
	}
	for index, value := range bundle.CodeEditions {
		if _, ok := families[value.FamilyID]; !ok { diagnostics = append(diagnostics, missingReference(fmt.Sprintf("code_editions[%d].family_id", index), value.FamilyID, "code family")) }
	}
	for index, value := range bundle.Adoptions {
		if _, ok := instruments[value.InstrumentID]; !ok { diagnostics = append(diagnostics, missingReference(fmt.Sprintf("adoptions[%d].instrument_id", index), value.InstrumentID, "regulatory instrument")) }
		if _, ok := editions[value.CodeEditionID]; !ok { diagnostics = append(diagnostics, missingReference(fmt.Sprintf("adoptions[%d].code_edition_id", index), value.CodeEditionID, "code edition")) }
		if strings.TrimSpace(value.AuthorityID) != "" { if _, ok := authorities[value.AuthorityID]; !ok { diagnostics = append(diagnostics, missingReference(fmt.Sprintf("adoptions[%d].authority_id", index), value.AuthorityID, "authority")) } }
	}
	for index, value := range bundle.Amendments {
		if _, ok := authorities[value.AuthorityID]; !ok { diagnostics = append(diagnostics, missingReference(fmt.Sprintf("amendments[%d].authority_id", index), value.AuthorityID, "authority")) }
		if _, ok := instruments[value.InstrumentID]; !ok { diagnostics = append(diagnostics, missingReference(fmt.Sprintf("amendments[%d].instrument_id", index), value.InstrumentID, "regulatory instrument")) }
		if _, ok := adoptions[value.AdoptionID]; !ok { diagnostics = append(diagnostics, missingReference(fmt.Sprintf("amendments[%d].adoption_id", index), value.AdoptionID, "adoption")) }
	}
	for index, value := range bundle.SourceArtifacts { if _, ok := documents[value.DocumentID]; !ok { diagnostics = append(diagnostics, missingReference(fmt.Sprintf("source_artifacts[%d].document_id", index), value.DocumentID, "source document")) } }
	for index, value := range bundle.TextAnchors { if _, ok := artifacts[value.ArtifactID]; !ok { diagnostics = append(diagnostics, missingReference(fmt.Sprintf("text_anchors[%d].artifact_id", index), value.ArtifactID, "source artifact")) } }
	for index, value := range bundle.Claims {
		if _, ok := jurisdictions[value.JurisdictionID]; !ok { diagnostics = append(diagnostics, missingReference(fmt.Sprintf("claims[%d].jurisdiction_id", index), value.JurisdictionID, "jurisdiction")) }
		switch value.Kind {
		case "code_adoption": if _, ok := adoptions[value.SubjectRef]; !ok { diagnostics = append(diagnostics, missingReference(fmt.Sprintf("claims[%d].subject_ref", index), value.SubjectRef, "adoption")) }
		case "amendment": if _, ok := amendments[value.SubjectRef]; !ok { diagnostics = append(diagnostics, missingReference(fmt.Sprintf("claims[%d].subject_ref", index), value.SubjectRef, "amendment")) }
		}
	}
	for index, value := range bundle.EvidenceLinks {
		if _, ok := claims[value.ClaimID]; !ok { diagnostics = append(diagnostics, missingReference(fmt.Sprintf("evidence_links[%d].claim_id", index), value.ClaimID, "claim")) }
		if _, ok := documents[value.DocumentID]; !ok { diagnostics = append(diagnostics, missingReference(fmt.Sprintf("evidence_links[%d].document_id", index), value.DocumentID, "source document")) }
		if _, ok := artifacts[value.ArtifactID]; !ok { diagnostics = append(diagnostics, missingReference(fmt.Sprintf("evidence_links[%d].artifact_id", index), value.ArtifactID, "source artifact")) }
		if _, ok := anchors[value.AnchorID]; !ok { diagnostics = append(diagnostics, missingReference(fmt.Sprintf("evidence_links[%d].anchor_id", index), value.AnchorID, "text anchor")) }
	}
	for index, value := range bundle.SourcePolicyRules {
		if strings.TrimSpace(value.JurisdictionID) != "" { if _, ok := jurisdictions[value.JurisdictionID]; !ok { diagnostics = append(diagnostics, missingReference(fmt.Sprintf("source_policy_rules[%d].jurisdiction_id", index), value.JurisdictionID, "jurisdiction")) } }
		if strings.TrimSpace(value.DocumentID) != "" { if _, ok := documents[value.DocumentID]; !ok { diagnostics = append(diagnostics, missingReference(fmt.Sprintf("source_policy_rules[%d].document_id", index), value.DocumentID, "source document")) } }
	}
	return diagnostics
}

func validateInterval(ruleID, path, fromValue, toValue string) []ProviderConformanceDiagnostic {
	diagnostics := make([]ProviderConformanceDiagnostic, 0)
	var from, to time.Time
	var fromSet, toSet bool
	if strings.TrimSpace(fromValue) != "" {
		parsed, err := time.Parse(time.DateOnly, fromValue)
		if err != nil { diagnostics = append(diagnostics, conformanceDiagnostic(ruleID, path+".effective_from", "effective_from must use YYYY-MM-DD")) } else { from, fromSet = parsed, true }
	}
	if strings.TrimSpace(toValue) != "" {
		parsed, err := time.Parse(time.DateOnly, toValue)
		if err != nil { diagnostics = append(diagnostics, conformanceDiagnostic(ruleID, path+".effective_to", "effective_to must use YYYY-MM-DD")) } else { to, toSet = parsed, true }
	}
	if fromSet && toSet && from.After(to) { diagnostics = append(diagnostics, conformanceDiagnostic(ruleID, path, "effective interval is inverted")) }
	return diagnostics
}

func providerEvidence(bundle ProviderBundle) (map[string]EvidenceLink, map[string][]EvidenceLink) {
	documents := idSet(bundle.SourceDocuments, func(value SourceDocument) string { return value.ID })
	artifacts := idSet(bundle.SourceArtifacts, func(value SourceArtifact) string { return value.ID })
	anchors := idSet(bundle.TextAnchors, func(value TextAnchor) string { return value.ID })
	byID := map[string]EvidenceLink{}
	byClaim := map[string][]EvidenceLink{}
	for _, reference := range bundle.EvidenceLinks {
		document, documentOK := documents[reference.DocumentID]
		artifact, artifactOK := artifacts[reference.ArtifactID]
		anchor, anchorOK := anchors[reference.AnchorID]
		if !documentOK || !artifactOK || !anchorOK { continue }
		link := EvidenceLink{ID: reference.ID, ClaimID: reference.ClaimID, Role: reference.Role, Document: document, Artifact: artifact, Anchor: anchor}
		byID[link.ID] = link
		byClaim[link.ClaimID] = append(byClaim[link.ClaimID], link)
	}
	return byID, byClaim
}

func providerSourcePolicy(bundle ProviderBundle) SourcePolicy {
	rules := make([]SourcePolicyRule, 0, len(bundle.SourcePolicyRules))
	for _, rule := range bundle.SourcePolicyRules {
		if strings.TrimSpace(rule.ClaimKind) == "" || strings.TrimSpace(rule.JurisdictionID) == "" || strings.TrimSpace(rule.DocumentID) == "" { continue }
		rules = append(rules, SourcePolicyRule{ID: rule.ID, ClaimKind: rule.ClaimKind, JurisdictionID: rule.JurisdictionID, DocumentID: rule.DocumentID, EffectiveFrom: rule.EffectiveFrom, EffectiveTo: rule.EffectiveTo})
	}
	return SourcePolicy{Rules: rules}
}

func ValidateProviderConformance(contract ProviderContractDeclaration, bundle ProviderBundle) []ProviderConformanceDiagnostic {
	diagnostics := make([]ProviderConformanceDiagnostic, 0)

	if strings.TrimSpace(bundle.ID) == "" { diagnostics = append(diagnostics, conformanceDiagnostic("BCM-CONF-001", "id", "provider bundle stable id is required")) }
	diagnostics = append(diagnostics, stableIDs(idsOfJurisdictions(bundle.Jurisdictions), "jurisdictions")...)
	diagnostics = append(diagnostics, stableIDs(idsOfAuthorities(bundle.Authorities), "authorities")...)
	diagnostics = append(diagnostics, stableIDs(idsOfInstruments(bundle.RegulatoryInstruments), "regulatory_instruments")...)
	diagnostics = append(diagnostics, stableIDs(idsOfAdoptions(bundle.Adoptions), "adoptions")...)
	diagnostics = append(diagnostics, stableIDs(idsOfAmendments(bundle.Amendments), "amendments")...)
	diagnostics = append(diagnostics, stableIDs(idsOfFamilies(bundle.CodeFamilies), "code_families")...)
	diagnostics = append(diagnostics, stableIDs(idsOfEditions(bundle.CodeEditions), "code_editions")...)
	diagnostics = append(diagnostics, stableIDs(idsOfDocuments(bundle.SourceDocuments), "source_documents")...)
	diagnostics = append(diagnostics, stableIDs(idsOfArtifacts(bundle.SourceArtifacts), "source_artifacts")...)
	diagnostics = append(diagnostics, stableIDs(idsOfAnchors(bundle.TextAnchors), "text_anchors")...)
	diagnostics = append(diagnostics, stableIDs(idsOfClaims(bundle.Claims), "claims")...)
	diagnostics = append(diagnostics, stableIDs(idsOfLinks(bundle.EvidenceLinks), "evidence_links")...)
	diagnostics = append(diagnostics, stableIDs(idsOfPolicies(bundle.SourcePolicyRules), "source_policy_rules")...)

	diagnostics = append(diagnostics, validateProviderReferences(bundle)...)

	authorities := idSet(bundle.Authorities, func(value ProviderAuthority) string { return value.ID })
	for index, adoption := range bundle.Adoptions {
		if strings.TrimSpace(adoption.AuthorityID) == "" { diagnostics = append(diagnostics, conformanceDiagnostic("BCM-CONF-003", fmt.Sprintf("adoptions[%d].authority_id", index), "adoption must identify an adopting authority")); continue }
		if _, ok := authorities[adoption.AuthorityID]; !ok { diagnostics = append(diagnostics, conformanceDiagnostic("BCM-CONF-003", fmt.Sprintf("adoptions[%d].authority_id", index), "adopting authority does not resolve")) }
	}

	for index, value := range bundle.Adoptions { diagnostics = append(diagnostics, validateInterval("BCM-CONF-004", fmt.Sprintf("adoptions[%d]", index), value.EffectiveFrom, value.EffectiveTo)...)}
	for index, value := range bundle.Amendments { diagnostics = append(diagnostics, validateInterval("BCM-CONF-004", fmt.Sprintf("amendments[%d]", index), value.EffectiveFrom, value.EffectiveTo)...)}
	for index, value := range bundle.SourcePolicyRules { diagnostics = append(diagnostics, validateInterval("BCM-CONF-004", fmt.Sprintf("source_policy_rules[%d]", index), value.EffectiveFrom, value.EffectiveTo)...)}

	exactByID, exactByClaim := providerEvidence(bundle)
	for index, reference := range bundle.EvidenceLinks {
		link, ok := exactByID[reference.ID]
		if !ok { continue }
		if err := validateEvidenceLink(link, reference.ClaimID); err != nil { diagnostics = append(diagnostics, conformanceDiagnostic("BCM-CONF-006", fmt.Sprintf("evidence_links[%d]", index), err.Error())) }
	}

	policy := providerSourcePolicy(bundle)
	for index, claim := range bundle.Claims {
		if claim.Status != StatusResolved { continue }
		verification, err := VerifyClaimEvidence(EvidenceClaim{ID: claim.ID, Kind: claim.Kind, JurisdictionID: claim.JurisdictionID, AsOf: claim.AsOf, LegallyMaterial: claim.LegallyMaterial, Links: exactByClaim[claim.ID]}, policy)
		if err != nil { diagnostics = append(diagnostics, conformanceDiagnostic("BCM-CONF-005", fmt.Sprintf("claims[%d]", index), err.Error())); continue }
		if verification.Decision != EvidenceVerified { diagnostics = append(diagnostics, conformanceDiagnostic("BCM-CONF-005", fmt.Sprintf("claims[%d]", index), fmt.Sprintf("resolved claim evidence decision is %s", verification.Decision))) }
	}

	if strings.TrimSpace(contract.BundleID) == "" || contract.BundleID != bundle.ID { diagnostics = append(diagnostics, conformanceDiagnostic("BCM-CONF-009", "contract.bundle_id", "contract bundle identity must match represented provider bundle")) }
	if strings.TrimSpace(contract.Limitations) == "" { diagnostics = append(diagnostics, conformanceDiagnostic("BCM-CONF-009", "contract.limitations", "contract limitations must be explicit")) }
	if len(bundle.Limitations) == 0 { diagnostics = append(diagnostics, conformanceDiagnostic("BCM-CONF-009", "limitations", "provider bundle limitations must be explicit")) } else { for index, limitation := range bundle.Limitations { if strings.TrimSpace(limitation) == "" { diagnostics = append(diagnostics, conformanceDiagnostic("BCM-CONF-009", fmt.Sprintf("limitations[%d]", index), "provider bundle limitation must not be empty")) } } }
	represented := map[string]struct{}{}; for _, jurisdiction := range bundle.Jurisdictions { if id := strings.TrimSpace(jurisdiction.ID); id != "" { represented[id] = struct{}{} } }
	declared := map[string]struct{}{}; for _, id := range bundle.Coverage.JurisdictionIDs { if id = strings.TrimSpace(id); id != "" { declared[id] = struct{}{} } }
	if len(represented) != len(declared) { diagnostics = append(diagnostics, conformanceDiagnostic("BCM-CONF-009", "coverage.jurisdiction_ids", "declared jurisdiction coverage must equal represented jurisdictions")) } else { for id := range represented { if _, ok := declared[id]; !ok { diagnostics = append(diagnostics, conformanceDiagnostic("BCM-CONF-009", "coverage.jurisdiction_ids", "declared jurisdiction coverage must equal represented jurisdictions")); break } } }

	return sortConformanceDiagnostics(diagnostics)
}

func hasAffirmativeNotApplicableBasis(code CodeResolution) bool {
	return code.Status != StatusNotApplicable || strings.TrimSpace(code.Basis) != ""
}

func ValidateResolutionConformance(resolution Resolution) []ProviderConformanceDiagnostic {
	diagnostics := make([]ProviderConformanceDiagnostic, 0)
	for index, code := range resolution.Codes {
		if !hasAffirmativeNotApplicableBasis(code) { diagnostics = append(diagnostics, conformanceDiagnostic("BCM-CONF-007", fmt.Sprintf("codes[%d].basis", index), "not_applicable requires an affirmative basis")) }
	}
	return sortConformanceDiagnostics(diagnostics)
}

func ValidateProjectConformance(project ProjectCodeBasis) []ProviderConformanceDiagnostic {
	diagnostics := make([]ProviderConformanceDiagnostic, 0)
	followUps := map[string]struct{}{}
	for _, question := range project.FollowUpQuestions { if normalized := strings.TrimSpace(question); normalized != "" { followUps[normalized] = struct{}{} } }
	hasEvidenceDefect := false
	for index, unresolved := range project.Unresolved {
		if unresolved.Kind != RequirementEvidenceDefect { continue }
		hasEvidenceDefect = true
		message := strings.TrimSpace(unresolved.Message)
		if _, ok := followUps[message]; ok && message != "" { diagnostics = append(diagnostics, conformanceDiagnostic("BCM-CONF-008", fmt.Sprintf("unresolved[%d]", index), "evidence defect must not be emitted as a project-fact follow-up question")) }
	}
	if hasEvidenceDefect && project.Verdict != ProjectNotVerified { diagnostics = append(diagnostics, conformanceDiagnostic("BCM-CONF-008", "verdict", "evidence defect requires not_verified project verdict")) }
	return sortConformanceDiagnostics(diagnostics)
}