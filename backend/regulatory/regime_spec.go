package regulatory

import (
	"fmt"
	"sort"
	"strings"
)

const StateRegimeSchemaVersion = "1.0"

type RegimeEvidenceStatus string

const (
	RegimeEvidenceEvidenced  RegimeEvidenceStatus = "evidenced"
	RegimeEvidencePartial    RegimeEvidenceStatus = "partial"
	RegimeEvidenceUnresolved RegimeEvidenceStatus = "unresolved"
)

type RegimePower string

const (
	RegimePowerAdopt      RegimePower = "adopt"
	RegimePowerAmend      RegimePower = "amend"
	RegimePowerAdminister RegimePower = "administer"
	RegimePowerEnforce    RegimePower = "enforce"
	RegimePowerDelegate   RegimePower = "delegate"
	RegimePowerContract   RegimePower = "contract"
)

type RegimeDimensionEvidence struct {
	Status    RegimeEvidenceStatus `json:"status"`
	SourceIDs []string             `json:"source_ids,omitempty"`
}

type StateRegimeDimensions struct {
	LocalAdoption          RegimeDimensionEvidence `json:"local_adoption"`
	LocalAmendment         RegimeDimensionEvidence `json:"local_amendment"`
	Enforcement            RegimeDimensionEvidence `json:"enforcement"`
	DelegationContracting  RegimeDimensionEvidence `json:"delegation_contracting"`
	HomeRule               RegimeDimensionEvidence `json:"home_rule"`
	StatewideInheritance   RegimeDimensionEvidence `json:"statewide_inheritance"`
	Unincorporated         RegimeDimensionEvidence `json:"unincorporated"`
	Temporal               RegimeDimensionEvidence `json:"temporal"`
}

type StateRegimeTerritory struct {
	Incorporated   ResolutionPolicy `json:"incorporated"`
	Unincorporated ResolutionPolicy `json:"unincorporated"`
}

type EntityPowerEvidence struct {
	EntityKind string        `json:"entity_kind"`
	Scope      string        `json:"scope"`
	Powers     []RegimePower `json:"powers"`
	SourceIDs  []string      `json:"source_ids"`
}

type StateRegimeSpec struct {
	SchemaVersion         string                      `json:"schema_version"`
	RegimeID              string                      `json:"regime_id"`
	StateID               string                      `json:"state_id"`
	StateName             string                      `json:"state_name"`
	StateAbbreviation     string                      `json:"state_abbreviation"`
	StateFIPS             string                      `json:"state_fips"`
	SourceProfileID       string                      `json:"source_profile_id"`
	SourceProfileStatus   string                      `json:"source_profile_status"`
	LastVerified          string                      `json:"last_verified"`
	Dimensions            StateRegimeDimensions       `json:"dimensions"`
	Territory             StateRegimeTerritory        `json:"territory"`
	CodeFamilyPolicies    map[string]ResolutionPolicy `json:"code_family_policies,omitempty"`
	ProjectTypeExceptions map[string]ResolutionPolicy `json:"project_type_exceptions,omitempty"`
	EntityPowerEvidence   []EntityPowerEvidence       `json:"entity_power_evidence,omitempty"`
	RequiredLocalEvidence []string                    `json:"required_local_evidence,omitempty"`
	ApplicabilityRules    []ApplicabilityRule         `json:"applicability_rules,omitempty"`
	DateRules             []DateRule                  `json:"date_rules,omitempty"`
	AmendmentRules        []AmendmentRule             `json:"amendment_rules,omitempty"`
	EnforcementRules      []EnforcementRule           `json:"enforcement_rules,omitempty"`
	Relationships         []AuthorityRelationship     `json:"relationships,omitempty"`
	UnresolvedDimensions  []string                    `json:"unresolved_dimensions,omitempty"`
	SourceIDs             []string                    `json:"source_ids"`
	Verification          Verification                `json:"verification"`
}

func BuildStateRegimeSpec(profile StateProfile) (StateRegimeSpec, error) {
	if profile.SchemaVersion != ProfileSchemaVersion {
		return StateRegimeSpec{}, fmt.Errorf("unsupported state profile schema_version %q", profile.SchemaVersion)
	}
	if strings.TrimSpace(profile.ProfileID) == "" || strings.TrimSpace(profile.StateID) == "" || strings.TrimSpace(profile.StateName) == "" {
		return StateRegimeSpec{}, fmt.Errorf("state profile identity is incomplete")
	}
	if len(profile.Sources) == 0 {
		return StateRegimeSpec{}, fmt.Errorf("state profile %q has no sources", profile.ProfileID)
	}

	powerEvidence := collectEntityPowerEvidence(profile)
	dimensions := StateRegimeDimensions{
		LocalAdoption:         dimensionFromPower(powerEvidence, RegimePowerAdopt),
		LocalAmendment:        localAmendmentDimension(profile.AmendmentRules),
		Enforcement:           enforcementDimension(profile.EnforcementRules, powerEvidence),
		DelegationContracting: delegationDimension(profile.Relationships),
		HomeRule:              homeRuleDimension(profile.Claims),
		StatewideInheritance:  statewideInheritanceDimension(profile.Defaults),
		Unincorporated:        territoryDimension(profile.Defaults.Unincorporated),
		Temporal:              temporalDimension(profile.DateRules, profile.Adoptions),
	}

	spec := StateRegimeSpec{
		SchemaVersion:         StateRegimeSchemaVersion,
		RegimeID:              "state-regime:" + strings.ToLower(profile.StateID),
		StateID:               profile.StateID,
		StateName:             profile.StateName,
		StateAbbreviation:     profile.StateAbbreviation,
		StateFIPS:             profile.StateFIPS,
		SourceProfileID:       profile.ProfileID,
		SourceProfileStatus:   profile.Status,
		LastVerified:          profile.LastVerified,
		Dimensions:            dimensions,
		Territory:             StateRegimeTerritory{Incorporated: profile.Defaults.Incorporated, Unincorporated: profile.Defaults.Unincorporated},
		CodeFamilyPolicies:    clonePolicyMap(profile.CodeFamilyOverrides),
		ProjectTypeExceptions: clonePolicyMap(profile.ProjectTypeOverrides),
		EntityPowerEvidence:   powerEvidence,
		RequiredLocalEvidence: collectRequiredLocalEvidence(profile),
		ApplicabilityRules:    append([]ApplicabilityRule(nil), profile.ApplicabilityRules...),
		DateRules:             append([]DateRule(nil), profile.DateRules...),
		AmendmentRules:        append([]AmendmentRule(nil), profile.AmendmentRules...),
		EnforcementRules:      append([]EnforcementRule(nil), profile.EnforcementRules...),
		Relationships:         append([]AuthorityRelationship(nil), profile.Relationships...),
		SourceIDs:             profileSourceIDs(profile.Sources),
		Verification:          profile.Verification,
	}
	spec.UnresolvedDimensions = unresolvedRegimeDimensions(dimensions)
	if err := ValidateStateRegimeSpec(spec); err != nil {
		return StateRegimeSpec{}, err
	}
	return spec, nil
}

func ValidateStateRegimeSpec(spec StateRegimeSpec) error {
	if spec.SchemaVersion != StateRegimeSchemaVersion {
		return fmt.Errorf("unsupported state regime schema_version %q", spec.SchemaVersion)
	}
	if strings.TrimSpace(spec.RegimeID) == "" || strings.TrimSpace(spec.StateID) == "" || strings.TrimSpace(spec.SourceProfileID) == "" {
		return fmt.Errorf("state regime identity is incomplete")
	}
	if spec.RegimeID != "state-regime:"+strings.ToLower(spec.StateID) {
		return fmt.Errorf("regime_id %q does not match state_id %q", spec.RegimeID, spec.StateID)
	}
	if len(spec.SourceIDs) == 0 {
		return fmt.Errorf("state regime requires source_ids")
	}
	for name, evidence := range map[string]RegimeDimensionEvidence{
		"local_adoption":          spec.Dimensions.LocalAdoption,
		"local_amendment":         spec.Dimensions.LocalAmendment,
		"enforcement":             spec.Dimensions.Enforcement,
		"delegation_contracting":  spec.Dimensions.DelegationContracting,
		"home_rule":               spec.Dimensions.HomeRule,
		"statewide_inheritance":   spec.Dimensions.StatewideInheritance,
		"unincorporated":          spec.Dimensions.Unincorporated,
		"temporal":                spec.Dimensions.Temporal,
	} {
		if !validRegimeEvidenceStatus(evidence.Status) {
			return fmt.Errorf("dimension %s has invalid status %q", name, evidence.Status)
		}
		if evidence.Status != RegimeEvidenceUnresolved && len(evidence.SourceIDs) == 0 {
			return fmt.Errorf("dimension %s is %s without source evidence", name, evidence.Status)
		}
	}
	return nil
}

func collectEntityPowerEvidence(profile StateProfile) []EntityPowerEvidence {
	var result []EntityPowerEvidence
	appendPolicyPowerEvidence := func(scope string, policy ResolutionPolicy) {
		for _, candidate := range policy.AuthorityCandidates {
			powers := powersFromRoles(candidate.Roles)
			if len(powers) == 0 {
				continue
			}
			result = append(result, EntityPowerEvidence{
				EntityKind: strings.TrimSpace(candidate.Kind),
				Scope:      scope,
				Powers:     powers,
				SourceIDs:  sortedUnique(candidate.SourceIDs),
			})
		}
	}
	appendPolicyPowerEvidence("incorporated", profile.Defaults.Incorporated)
	appendPolicyPowerEvidence("unincorporated", profile.Defaults.Unincorporated)

	codeFamilies := sortedMapKeys(profile.CodeFamilyOverrides)
	for _, codeFamily := range codeFamilies {
		appendPolicyPowerEvidence("code_family:"+codeFamily, profile.CodeFamilyOverrides[codeFamily])
	}
	projectTypes := sortedMapKeys(profile.ProjectTypeOverrides)
	for _, projectType := range projectTypes {
		appendPolicyPowerEvidence("project_type:"+projectType, profile.ProjectTypeOverrides[projectType])
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].EntityKind != result[j].EntityKind {
			return result[i].EntityKind < result[j].EntityKind
		}
		return result[i].Scope < result[j].Scope
	})
	return result
}

func powersFromRoles(roles []string) []RegimePower {
	seen := map[RegimePower]bool{}
	for _, role := range roles {
		switch strings.ToLower(strings.TrimSpace(role)) {
		case "adopts", "adopt":
			seen[RegimePowerAdopt] = true
		case "amends", "amend":
			seen[RegimePowerAmend] = true
		case "permits", "plan_review", "inspects", "administers", "administer":
			seen[RegimePowerAdminister] = true
		case "enforces", "enforce":
			seen[RegimePowerEnforce] = true
		case "delegates", "delegate":
			seen[RegimePowerDelegate] = true
		case "contracts", "contract":
			seen[RegimePowerContract] = true
		}
	}
	order := []RegimePower{RegimePowerAdopt, RegimePowerAmend, RegimePowerAdminister, RegimePowerEnforce, RegimePowerDelegate, RegimePowerContract}
	result := make([]RegimePower, 0, len(seen))
	for _, power := range order {
		if seen[power] {
			result = append(result, power)
		}
	}
	return result
}

func dimensionFromPower(evidence []EntityPowerEvidence, power RegimePower) RegimeDimensionEvidence {
	var sources []string
	for _, item := range evidence {
		if isStateAuthorityKind(item.EntityKind) {
			continue
		}
		if containsPower(item.Powers, power) {
			sources = append(sources, item.SourceIDs...)
		}
	}
	return evidencedIfSources(sources)
}

func localAmendmentDimension(rules []AmendmentRule) RegimeDimensionEvidence {
	var sources []string
	for _, rule := range rules {
		level := strings.ToLower(strings.TrimSpace(rule.Level))
		if strings.Contains(level, "local") || strings.Contains(level, "municip") || strings.Contains(level, "county") {
			sources = append(sources, rule.SourceIDs...)
		}
	}
	return evidencedIfSources(sources)
}

func enforcementDimension(rules []EnforcementRule, powers []EntityPowerEvidence) RegimeDimensionEvidence {
	var sources []string
	for _, rule := range rules {
		sources = append(sources, rule.SourceIDs...)
	}
	for _, item := range powers {
		if containsPower(item.Powers, RegimePowerEnforce) {
			sources = append(sources, item.SourceIDs...)
		}
	}
	return evidencedIfSources(sources)
}

func delegationDimension(relationships []AuthorityRelationship) RegimeDimensionEvidence {
	var sources []string
	for _, relationship := range relationships {
		name := strings.ToLower(relationship.Relationship)
		if strings.Contains(name, "delegat") || strings.Contains(name, "contract") {
			sources = append(sources, relationship.SourceIDs...)
		}
	}
	return evidencedIfSources(sources)
}

func homeRuleDimension(claims []Claim) RegimeDimensionEvidence {
	var sources []string
	for _, claim := range claims {
		field := strings.ToLower(strings.TrimSpace(claim.Field))
		if field == "home_rule" || field == "home_rule_status" || field == "home_rule_authority" {
			sources = append(sources, claim.SourceIDs...)
		}
	}
	return evidencedIfSources(sources)
}

func statewideInheritanceDimension(defaults DefaultPolicies) RegimeDimensionEvidence {
	shared := intersectStrings(defaults.Incorporated.AdoptionIDs, defaults.Unincorporated.AdoptionIDs)
	if len(shared) == 0 {
		return RegimeDimensionEvidence{Status: RegimeEvidenceUnresolved}
	}
	sources := append([]string(nil), defaults.Incorporated.SourceIDs...)
	sources = append(sources, defaults.Unincorporated.SourceIDs...)
	return RegimeDimensionEvidence{Status: RegimeEvidenceEvidenced, SourceIDs: sortedUnique(sources)}
}

func territoryDimension(policy ResolutionPolicy) RegimeDimensionEvidence {
	if strings.TrimSpace(policy.Status) == "" || len(policy.SourceIDs) == 0 {
		return RegimeDimensionEvidence{Status: RegimeEvidenceUnresolved}
	}
	return RegimeDimensionEvidence{Status: RegimeEvidenceEvidenced, SourceIDs: sortedUnique(policy.SourceIDs)}
}

func temporalDimension(dateRules []DateRule, adoptions []Adoption) RegimeDimensionEvidence {
	var sources []string
	for _, rule := range dateRules {
		sources = append(sources, rule.SourceIDs...)
	}
	for _, adoption := range adoptions {
		if adoption.Dates != (AdoptionDates{}) {
			sources = append(sources, adoption.SourceIDs...)
		}
	}
	return evidencedIfSources(sources)
}

func collectRequiredLocalEvidence(profile StateProfile) []string {
	var values []string
	values = append(values, profile.Defaults.Incorporated.RequiredLocalRecords...)
	values = append(values, profile.Defaults.Unincorporated.RequiredLocalRecords...)
	for _, key := range sortedMapKeys(profile.CodeFamilyOverrides) {
		values = append(values, profile.CodeFamilyOverrides[key].RequiredLocalRecords...)
	}
	for _, key := range sortedMapKeys(profile.ProjectTypeOverrides) {
		values = append(values, profile.ProjectTypeOverrides[key].RequiredLocalRecords...)
	}
	return sortedUnique(values)
}

func unresolvedRegimeDimensions(dimensions StateRegimeDimensions) []string {
	values := map[string]RegimeDimensionEvidence{
		"local_adoption":         dimensions.LocalAdoption,
		"local_amendment":        dimensions.LocalAmendment,
		"enforcement":            dimensions.Enforcement,
		"delegation_contracting": dimensions.DelegationContracting,
		"home_rule":              dimensions.HomeRule,
		"statewide_inheritance":  dimensions.StatewideInheritance,
		"unincorporated":         dimensions.Unincorporated,
		"temporal":               dimensions.Temporal,
	}
	var result []string
	for name, evidence := range values {
		if evidence.Status == RegimeEvidenceUnresolved {
			result = append(result, name)
		}
	}
	sort.Strings(result)
	return result
}

func profileSourceIDs(sources []Source) []string {
	result := make([]string, 0, len(sources))
	for _, source := range sources {
		result = append(result, source.ID)
	}
	return sortedUnique(result)
}

func evidencedIfSources(sources []string) RegimeDimensionEvidence {
	sources = sortedUnique(sources)
	if len(sources) == 0 {
		return RegimeDimensionEvidence{Status: RegimeEvidenceUnresolved}
	}
	return RegimeDimensionEvidence{Status: RegimeEvidenceEvidenced, SourceIDs: sources}
}

func containsPower(values []RegimePower, target RegimePower) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

func isStateAuthorityKind(kind string) bool {
	kind = strings.ToLower(strings.TrimSpace(kind))
	return strings.Contains(kind, "state") && !strings.Contains(kind, "estate")
}

func intersectStrings(left, right []string) []string {
	rightSet := map[string]bool{}
	for _, value := range right {
		rightSet[value] = true
	}
	var result []string
	for _, value := range left {
		if rightSet[value] {
			result = append(result, value)
		}
	}
	return sortedUnique(result)
}

func sortedUnique(values []string) []string {
	seen := map[string]bool{}
	result := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" || seen[value] {
			continue
		}
		seen[value] = true
		result = append(result, value)
	}
	sort.Strings(result)
	return result
}

func sortedMapKeys[T any](values map[string]T) []string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func clonePolicyMap(values map[string]ResolutionPolicy) map[string]ResolutionPolicy {
	if len(values) == 0 {
		return nil
	}
	result := make(map[string]ResolutionPolicy, len(values))
	for key, value := range values {
		result[key] = value
	}
	return result
}

func validRegimeEvidenceStatus(value RegimeEvidenceStatus) bool {
	switch value {
	case RegimeEvidenceEvidenced, RegimeEvidencePartial, RegimeEvidenceUnresolved:
		return true
	default:
		return false
	}
}
