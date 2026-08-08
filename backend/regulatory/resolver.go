package regulatory

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func Resolve(catalog Catalog, request ResolutionRequest) (ResolutionResult, error) {
	return ResolveAt(catalog, request, time.Now().UTC())
}

// ResolveAt resolves a request using the supplied generation time. The
// explicit time keeps engine results reproducible while Resolve preserves the
// legacy wall-clock behavior for existing callers.
func ResolveAt(catalog Catalog, request ResolutionRequest, generatedAt time.Time) (ResolutionResult, error) {
	if request.Context == nil {
		return ResolutionResult{}, fmt.Errorf("geographic context is required")
	}
	generatedAt = generatedAt.UTC()
	applicabilityDate, err := resolveApplicabilityDate(request.ApplicabilityDate, generatedAt)
	if err != nil {
		return ResolutionResult{}, err
	}
	context := *request.Context
	codeFamily := normalizeKey(request.CodeFamily)
	projectType := normalizeKey(request.ProjectType)
	profile, ok := catalog.Profile(context.StateID, context.StateFIPS)
	if !ok {
		return ResolutionResult{
			SchemaVersion:     ResultSchemaVersion,
			GeneratedAt:       generatedAt,
			Geography:         context,
			CodeFamily:        codeFamily,
			ProjectType:       projectType,
			ApplicabilityDate: applicabilityDate,
			Status:            "insufficient_evidence",
			Warnings:          []string{"No validated state jurisdiction profile is available for the matched state."},
			Evidence:          []Source{},
		}, nil
	}

	policy := profile.Defaults.Unincorporated
	if context.Incorporated || context.Municipality != nil {
		context.Incorporated = true
		policy = profile.Defaults.Incorporated
	}
	if override, exists := profile.CodeFamilyOverrides[codeFamily]; codeFamily != "" && exists {
		policy = mergePolicies(policy, override)
	}
	if override, exists := profile.ProjectTypeOverrides[projectType]; projectType != "" && exists {
		policy = mergePolicies(policy, override)
	}
	policyBasis := &PolicyBasis{
		Status:               policy.Status,
		RequiredLocalRecords: deduplicateStrings(policy.RequiredLocalRecords),
		Warnings:             deduplicateStrings(policy.Warnings),
		SourceIDs:            deduplicateStrings(policy.SourceIDs),
		Verification:         profile.Verification,
	}

	authorities := make(map[string]Authority, len(profile.Authorities))
	for _, authority := range profile.Authorities {
		authorities[authority.ID] = authority
	}
	adoptions := make(map[string]Adoption, len(profile.Adoptions))
	for _, adoption := range profile.Adoptions {
		adoptions[adoption.ID] = adoption
	}

	candidates := make([]AuthorityCandidate, 0, len(policy.AuthorityCandidates))
	evidenceIDs := append([]string(nil), policy.SourceIDs...)
	for _, rule := range policy.AuthorityCandidates {
		candidate, include := expandCandidate(rule, context, authorities)
		if !include {
			continue
		}
		candidates = append(candidates, candidate)
		evidenceIDs = append(evidenceIDs, candidate.SourceIDs...)
	}
	candidates = deduplicateCandidates(candidates)

	status := policy.Status
	warnings := append([]string(nil), policy.Warnings...)
	requiredLocalRecords := append([]string(nil), policy.RequiredLocalRecords...)
	resolvedAdoptions := make([]Adoption, 0, len(policy.AdoptionIDs))
	for _, id := range policy.AdoptionIDs {
		adoption, exists := adoptions[id]
		if !exists {
			return ResolutionResult{}, fmt.Errorf("validated profile references missing adoption %s", id)
		}
		if !adoptionAppliesOn(adoption, applicabilityDate, profile.LastVerified) {
			continue
		}
		resolvedAdoptions = append(resolvedAdoptions, adoption)
		evidenceIDs = append(evidenceIDs, adoption.SourceIDs...)
	}
	if len(policy.AdoptionIDs) > 0 && len(resolvedAdoptions) == 0 {
		status = "insufficient_evidence"
		warnings = append(warnings, fmt.Sprintf("No supported adoption record in the profile applies on %s.", applicabilityDate))
		requiredLocalRecords = append(requiredLocalRecords, fmt.Sprintf("Adoption and amendment record effective on %s", applicabilityDate))
	}

	applicableRules, ruleSubjects, ruleEvidence := selectApplicableRules(profile, codeFamily, projectType, applicabilityDate)
	evidenceIDs = append(evidenceIDs, ruleEvidence...)
	authorityPath, relationshipSubjects, relationshipEvidence := selectAuthorityPath(profile.Relationships, codeFamily)
	evidenceIDs = append(evidenceIDs, relationshipEvidence...)

	subjects := map[string]struct{}{profile.ProfileID: {}}
	for _, candidate := range candidates {
		if candidate.AuthorityID != "" {
			subjects[candidate.AuthorityID] = struct{}{}
		}
	}
	for _, adoption := range resolvedAdoptions {
		subjects[adoption.ID] = struct{}{}
	}
	for id := range ruleSubjects {
		subjects[id] = struct{}{}
	}
	for id := range relationshipSubjects {
		subjects[id] = struct{}{}
	}
	supportingClaims, claimEvidence, hasConflictingClaims := selectSupportingClaims(profile.Claims, subjects)
	evidenceIDs = append(evidenceIDs, claimEvidence...)
	if hasConflictingClaims {
		status = "conflicting"
		warnings = append(warnings, "Relevant source-backed claims conflict and require review before relying on the result.")
	}

	evidence, evidenceWarnings := collectEvidence(profile.Sources, evidenceIDs)
	warnings = append(warnings, evidenceWarnings...)

	return ResolutionResult{
		SchemaVersion:        ResultSchemaVersion,
		GeneratedAt:          generatedAt,
		ProfileID:            profile.ProfileID,
		ProfileLastVerified:  profile.LastVerified,
		Geography:            context,
		CodeFamily:           codeFamily,
		ProjectType:          projectType,
		ApplicabilityDate:    applicabilityDate,
		Status:               status,
		PolicyBasis:          policyBasis,
		AuthorityCandidates:  candidates,
		AuthorityPath:        authorityPath,
		Adoptions:            resolvedAdoptions,
		ApplicableRules:      applicableRules,
		SupportingClaims:     supportingClaims,
		RequiredLocalRecords: deduplicateStrings(requiredLocalRecords),
		Warnings:             deduplicateStrings(warnings),
		Evidence:             evidence,
	}, nil
}

func resolveApplicabilityDate(value string, now time.Time) (string, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return now.UTC().Format(time.DateOnly), nil
	}
	if _, err := time.Parse(time.DateOnly, value); err != nil {
		return "", fmt.Errorf("applicability_date must use YYYY-MM-DD")
	}
	return value, nil
}

func adoptionAppliesOn(adoption Adoption, applicabilityDate, profileLastVerified string) bool {
	requested, err := time.Parse(time.DateOnly, applicabilityDate)
	if err != nil {
		return false
	}
	start := firstNonEmpty(
		adoption.Dates.OperativeDate,
		adoption.Dates.EffectiveDate,
		adoption.Dates.MandatoryDate,
		adoption.Dates.AdoptionDate,
	)
	if start != "" {
		parsed, err := time.Parse(time.DateOnly, start)
		if err != nil || requested.Before(parsed) {
			return false
		}
	} else {
		if adoption.Status == "future" || adoption.Status == "pending" {
			return false
		}
		verifiedAt, err := time.Parse(time.DateOnly, profileLastVerified)
		if err != nil || requested.Before(verifiedAt) {
			return false
		}
	}
	if replacement := strings.TrimSpace(adoption.Dates.ReplacementDate); replacement != "" {
		parsed, err := time.Parse(time.DateOnly, replacement)
		if err != nil || !requested.Before(parsed) {
			return false
		}
	}
	return true
}

func selectApplicableRules(profile StateProfile, codeFamily, projectType, applicabilityDate string) ([]RuleReference, map[string]struct{}, []string) {
	refs := make([]RuleReference, 0)
	subjects := map[string]struct{}{}
	evidenceIDs := make([]string, 0)
	add := func(id, kind, family, summary string, sources []string, verification Verification) {
		refs = append(refs, RuleReference{
			ID:           id,
			Kind:         kind,
			CodeFamily:   family,
			Summary:      summary,
			SourceIDs:    append([]string(nil), sources...),
			Verification: verification,
		})
		subjects[id] = struct{}{}
		evidenceIDs = append(evidenceIDs, sources...)
	}
	for _, rule := range profile.ApplicabilityRules {
		if !familyMatches(rule.CodeFamily, codeFamily) || !projectTypesMatch(rule.ProjectTypes, projectType) {
			continue
		}
		add(rule.ID, "applicability", rule.CodeFamily, rule.Summary, rule.SourceIDs, rule.Verification)
	}
	for _, rule := range profile.DateRules {
		if !familyMatches(rule.CodeFamily, codeFamily) || !dateRuleApplies(rule, applicabilityDate) {
			continue
		}
		add(rule.ID, "date", rule.CodeFamily, rule.Summary, rule.SourceIDs, rule.Verification)
	}
	for _, rule := range profile.AmendmentRules {
		if !familyMatches(rule.CodeFamily, codeFamily) {
			continue
		}
		add(rule.ID, "amendment", rule.CodeFamily, rule.Summary, rule.SourceIDs, rule.Verification)
	}
	for _, rule := range profile.EnforcementRules {
		if !familyMatches(rule.CodeFamily, codeFamily) {
			continue
		}
		add(rule.ID, "enforcement", rule.CodeFamily, rule.Summary, rule.SourceIDs, rule.Verification)
	}
	sort.Slice(refs, func(i, j int) bool {
		if refs[i].Kind == refs[j].Kind {
			return refs[i].ID < refs[j].ID
		}
		return refs[i].Kind < refs[j].Kind
	})
	return refs, subjects, evidenceIDs
}

func dateRuleApplies(rule DateRule, applicabilityDate string) bool {
	requested, err := time.Parse(time.DateOnly, applicabilityDate)
	if err != nil {
		return false
	}
	if rule.StartDate != "" {
		start, err := time.Parse(time.DateOnly, rule.StartDate)
		if err != nil || requested.Before(start) {
			return false
		}
	}
	if rule.EndDate != "" {
		end, err := time.Parse(time.DateOnly, rule.EndDate)
		if err != nil || requested.After(end) {
			return false
		}
	}
	return true
}

func selectAuthorityPath(relationships []AuthorityRelationship, codeFamily string) ([]AuthorityRelationship, map[string]struct{}, []string) {
	selected := make([]AuthorityRelationship, 0)
	subjects := map[string]struct{}{}
	evidenceIDs := make([]string, 0)
	for _, relationship := range relationships {
		if !relationshipScopeMatches(relationship.Scope, codeFamily) {
			continue
		}
		selected = append(selected, relationship)
		subjects[relationship.ID] = struct{}{}
		evidenceIDs = append(evidenceIDs, relationship.SourceIDs...)
	}
	sort.Slice(selected, func(i, j int) bool { return selected[i].ID < selected[j].ID })
	return selected, subjects, evidenceIDs
}

func relationshipScopeMatches(scope []string, codeFamily string) bool {
	if len(scope) == 0 || codeFamily == "" {
		return true
	}
	for _, value := range scope {
		normalized := normalizeKey(value)
		if normalized == "all" || normalized == codeFamily {
			return true
		}
		if normalized == "construction_code" && codeFamily != "fire_operational" {
			return true
		}
	}
	return false
}

func selectSupportingClaims(claims []Claim, subjects map[string]struct{}) ([]Claim, []string, bool) {
	selected := make([]Claim, 0)
	evidenceIDs := make([]string, 0)
	hasConflict := false
	for _, claim := range claims {
		if _, ok := subjects[claim.SubjectID]; !ok {
			continue
		}
		selected = append(selected, claim)
		evidenceIDs = append(evidenceIDs, claim.SourceIDs...)
		if claim.Status == "conflicting" {
			hasConflict = true
		}
	}
	sort.Slice(selected, func(i, j int) bool { return selected[i].ID < selected[j].ID })
	return selected, evidenceIDs, hasConflict
}

func collectEvidence(sources []Source, evidenceIDs []string) ([]Source, []string) {
	sourceByID := make(map[string]Source, len(sources))
	for _, source := range sources {
		sourceByID[source.ID] = source
	}
	evidence := make([]Source, 0)
	warnings := make([]string, 0)
	seenSource := map[string]struct{}{}
	for _, id := range evidenceIDs {
		if _, seen := seenSource[id]; seen {
			continue
		}
		source, exists := sourceByID[id]
		if !exists {
			continue
		}
		evidence = append(evidence, source)
		seenSource[id] = struct{}{}
		switch source.Availability {
		case "unavailable":
			warnings = append(warnings, fmt.Sprintf("Evidence source %q is currently unavailable.", source.Title))
		case "moved":
			warnings = append(warnings, fmt.Sprintf("Evidence source %q has moved and should be revalidated.", source.Title))
		}
	}
	sort.Slice(evidence, func(i, j int) bool { return evidence[i].ID < evidence[j].ID })
	return evidence, warnings
}

func familyMatches(ruleFamily, requested string) bool {
	if requested == "" || strings.TrimSpace(ruleFamily) == "" {
		return true
	}
	normalizedRule := normalizeKey(ruleFamily)
	if normalizedRule == "construction_code" {
		return requested != "fire_operational"
	}
	return normalizedRule == requested
}

func projectTypesMatch(projectTypes []string, requested string) bool {
	if len(projectTypes) == 0 {
		return true
	}
	if requested == "" {
		return false
	}
	for _, projectType := range projectTypes {
		if normalizeKey(projectType) == requested {
			return true
		}
	}
	return false
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value = strings.TrimSpace(value); value != "" {
			return value
		}
	}
	return ""
}

func mergePolicies(base, override ResolutionPolicy) ResolutionPolicy {
	merged := ResolutionPolicy{
		Status:               base.Status,
		AuthorityCandidates:  append([]CandidateRule(nil), base.AuthorityCandidates...),
		AdoptionIDs:          append([]string(nil), base.AdoptionIDs...),
		RequiredLocalRecords: append([]string(nil), base.RequiredLocalRecords...),
		Warnings:             append([]string(nil), base.Warnings...),
		SourceIDs:            append([]string(nil), base.SourceIDs...),
	}
	if override.Status != "" {
		merged.Status = override.Status
	}
	if override.ReplaceAuthorityCandidates {
		merged.AuthorityCandidates = append([]CandidateRule(nil), override.AuthorityCandidates...)
	} else {
		merged.AuthorityCandidates = append(merged.AuthorityCandidates, override.AuthorityCandidates...)
	}
	if override.ReplaceAdoptions {
		merged.AdoptionIDs = append([]string(nil), override.AdoptionIDs...)
	} else {
		merged.AdoptionIDs = deduplicateStrings(append(merged.AdoptionIDs, override.AdoptionIDs...))
	}
	merged.RequiredLocalRecords = deduplicateStrings(append(merged.RequiredLocalRecords, override.RequiredLocalRecords...))
	merged.Warnings = deduplicateStrings(append(merged.Warnings, override.Warnings...))
	merged.SourceIDs = deduplicateStrings(append(merged.SourceIDs, override.SourceIDs...))
	return merged
}

func expandCandidate(rule CandidateRule, context GeographicContext, authorities map[string]Authority) (AuthorityCandidate, bool) {
	candidate := AuthorityCandidate{
		Kind:         rule.Kind,
		Roles:        append([]string(nil), rule.Roles...),
		SourceIDs:    append([]string(nil), rule.SourceIDs...),
		Verification: Verification{Status: "partially_verified"},
	}
	switch rule.Kind {
	case "state_authority":
		authority, ok := authorities[rule.AuthorityID]
		if !ok {
			return AuthorityCandidate{}, false
		}
		candidate.AuthorityID = authority.ID
		candidate.Name = authority.Name
		candidate.Roles = deduplicateStrings(append(candidate.Roles, authority.Roles...))
		candidate.SourceIDs = deduplicateStrings(append(candidate.SourceIDs, authority.SourceIDs...))
		candidate.Verification = authority.Verification
	case "municipality":
		if context.Municipality == nil {
			return AuthorityCandidate{}, false
		}
		candidate.Name = context.Municipality.Name
		if rule.Label != "" {
			candidate.Name += " " + rule.Label
		}
	case "county":
		if context.County == nil {
			return AuthorityCandidate{}, false
		}
		candidate.Name = context.County.Name
		if rule.Label != "" {
			candidate.Name += " " + rule.Label
		}
	case "fire_authority":
		if len(context.FireJurisdictions) == 0 {
			candidate.Name = "Local fire code enforcing agency"
			if rule.Label != "" {
				candidate.Name = rule.Label
			}
			return candidate, true
		}
		candidate.Name = context.FireJurisdictions[0].Name
		if rule.Label != "" {
			candidate.Name += " " + rule.Label
		}
	case "special_area":
		if len(context.SpecialAreas) == 0 {
			return AuthorityCandidate{}, false
		}
		candidate.Name = context.SpecialAreas[0].Name
	case "tribal_authority":
		if len(context.TribalAreas) == 0 {
			return AuthorityCandidate{}, false
		}
		candidate.Name = context.TribalAreas[0].Name
	default:
		return AuthorityCandidate{}, false
	}
	return candidate, true
}

func normalizeKey(value string) string {
	value = strings.ToLower(strings.TrimSpace(value))
	value = strings.NewReplacer("-", "_", " ", "_", "/", "_").Replace(value)
	for strings.Contains(value, "__") {
		value = strings.ReplaceAll(value, "__", "_")
	}
	switch value {
	case "fire", "fire_code", "operational_fire":
		return "fire_operational"
	case "existing", "existing_building_code":
		return "existing_building"
	}
	return value
}

func deduplicateStrings(values []string) []string {
	seen := map[string]struct{}{}
	result := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}
		if _, exists := seen[value]; exists {
			continue
		}
		seen[value] = struct{}{}
		result = append(result, value)
	}
	return result
}

func deduplicateCandidates(values []AuthorityCandidate) []AuthorityCandidate {
	seen := map[string]struct{}{}
	result := make([]AuthorityCandidate, 0, len(values))
	for _, candidate := range values {
		key := candidate.Kind + "\x00" + candidate.AuthorityID + "\x00" + candidate.Name
		if _, exists := seen[key]; exists {
			continue
		}
		seen[key] = struct{}{}
		result = append(result, candidate)
	}
	return result
}
