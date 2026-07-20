package regulatory

import (
	"fmt"
	"sort"
	"strings"
)

type SourceHealth struct {
	LastCheckedAt string `json:"last_checked_at,omitempty"`
	Availability  string `json:"availability,omitempty"`
}

type StateRulePack struct {
	SchemaVersion      string                  `json:"schema_version"`
	StateID            string                  `json:"state_id"`
	SourceHealth       map[string]SourceHealth `json:"source_health,omitempty"`
	ApplicabilityRules []ApplicabilityRule     `json:"applicability_rules,omitempty"`
	DateRules          []DateRule              `json:"date_rules,omitempty"`
	AmendmentRules     []AmendmentRule         `json:"amendment_rules,omitempty"`
	EnforcementRules   []EnforcementRule       `json:"enforcement_rules,omitempty"`
	Claims             []Claim                 `json:"claims,omitempty"`
	ResolverFixtureIDs []string                `json:"resolver_fixture_ids,omitempty"`
}

func MergeRulePack(profile StateProfile, pack StateRulePack) (StateProfile, error) {
	if pack.SchemaVersion != ProfileSchemaVersion {
		return StateProfile{}, fmt.Errorf("rule pack schema_version must be %q", ProfileSchemaVersion)
	}
	if !strings.EqualFold(strings.TrimSpace(profile.StateID), strings.TrimSpace(pack.StateID)) {
		return StateProfile{}, fmt.Errorf("rule pack state_id %q does not match profile %q", pack.StateID, profile.StateID)
	}

	merged := profile
	merged.Sources = append([]Source(nil), profile.Sources...)
	knownSources := make(map[string]struct{}, len(merged.Sources))
	for index := range merged.Sources {
		knownSources[merged.Sources[index].ID] = struct{}{}
		health, ok := pack.SourceHealth[merged.Sources[index].ID]
		if !ok {
			continue
		}
		merged.Sources[index].LastCheckedAt = health.LastCheckedAt
		merged.Sources[index].Availability = health.Availability
	}
	unknown := make([]string, 0)
	for sourceID := range pack.SourceHealth {
		if _, ok := knownSources[sourceID]; !ok {
			unknown = append(unknown, sourceID)
		}
	}
	if len(unknown) > 0 {
		sort.Strings(unknown)
		return StateProfile{}, fmt.Errorf("rule pack references unknown source health IDs: %s", strings.Join(unknown, ", "))
	}

	merged.ApplicabilityRules = append(append([]ApplicabilityRule(nil), profile.ApplicabilityRules...), pack.ApplicabilityRules...)
	merged.DateRules = append(append([]DateRule(nil), profile.DateRules...), pack.DateRules...)
	merged.AmendmentRules = append(append([]AmendmentRule(nil), profile.AmendmentRules...), pack.AmendmentRules...)
	merged.EnforcementRules = append(append([]EnforcementRule(nil), profile.EnforcementRules...), pack.EnforcementRules...)
	merged.Claims = append(append([]Claim(nil), profile.Claims...), pack.Claims...)
	merged.ResolverFixtureIDs = deduplicateStrings(append(append([]string(nil), profile.ResolverFixtureIDs...), pack.ResolverFixtureIDs...))
	if err := ValidateCompleteProfile(merged); err != nil {
		return StateProfile{}, err
	}
	return merged, nil
}
