package regulatory

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func Resolve(catalog Catalog, request ResolutionRequest) (ResolutionResult, error) {
	if request.Context == nil {
		return ResolutionResult{}, fmt.Errorf("geographic context is required")
	}
	context := *request.Context
	profile, ok := catalog.Profile(context.StateID, context.StateFIPS)
	if !ok {
		return ResolutionResult{
			SchemaVersion: ResultSchemaVersion,
			GeneratedAt:   time.Now().UTC(),
			Geography:     context,
			CodeFamily:    normalizeKey(request.CodeFamily),
			ProjectType:   normalizeKey(request.ProjectType),
			Status:        "insufficient_evidence",
			Warnings:      []string{"No validated state jurisdiction profile is available for the matched state."},
			Evidence:      []Source{},
		}, nil
	}

	policy := profile.Defaults.Unincorporated
	if context.Incorporated || context.Municipality != nil {
		context.Incorporated = true
		policy = profile.Defaults.Incorporated
	}
	if codeFamily := normalizeKey(request.CodeFamily); codeFamily != "" {
		if override, exists := profile.CodeFamilyOverrides[codeFamily]; exists {
			policy = mergePolicies(policy, override)
		}
	}
	if projectType := normalizeKey(request.ProjectType); projectType != "" {
		if override, exists := profile.ProjectTypeOverrides[projectType]; exists {
			policy = mergePolicies(policy, override)
		}
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

	resolvedAdoptions := make([]Adoption, 0, len(policy.AdoptionIDs))
	for _, id := range policy.AdoptionIDs {
		adoption, exists := adoptions[id]
		if !exists {
			return ResolutionResult{}, fmt.Errorf("validated profile references missing adoption %s", id)
		}
		resolvedAdoptions = append(resolvedAdoptions, adoption)
		evidenceIDs = append(evidenceIDs, adoption.SourceIDs...)
	}

	sourceByID := make(map[string]Source, len(profile.Sources))
	for _, source := range profile.Sources {
		sourceByID[source.ID] = source
	}
	evidence := make([]Source, 0)
	seenSource := map[string]struct{}{}
	for _, id := range evidenceIDs {
		if _, seen := seenSource[id]; seen {
			continue
		}
		if source, exists := sourceByID[id]; exists {
			evidence = append(evidence, source)
			seenSource[id] = struct{}{}
		}
	}
	sort.Slice(evidence, func(i, j int) bool { return evidence[i].ID < evidence[j].ID })

	return ResolutionResult{
		SchemaVersion:        ResultSchemaVersion,
		GeneratedAt:          time.Now().UTC(),
		ProfileID:            profile.ProfileID,
		ProfileLastVerified:  profile.LastVerified,
		Geography:            context,
		CodeFamily:           normalizeKey(request.CodeFamily),
		ProjectType:          normalizeKey(request.ProjectType),
		ApplicabilityDate:    request.ApplicabilityDate,
		Status:               policy.Status,
		AuthorityCandidates:  candidates,
		Adoptions:            resolvedAdoptions,
		RequiredLocalRecords: deduplicateStrings(policy.RequiredLocalRecords),
		Warnings:             deduplicateStrings(policy.Warnings),
		Evidence:             evidence,
	}, nil
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
