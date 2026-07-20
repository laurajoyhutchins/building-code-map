package regulatory

import (
	"fmt"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
)

var twoDigitFIPS = regexp.MustCompile(`^[0-9]{2}$`)

func ValidateProfile(profile StateProfile) error {
	var problems []string
	if profile.SchemaVersion != ProfileSchemaVersion {
		problems = append(problems, fmt.Sprintf("schema_version must be %q", ProfileSchemaVersion))
	}
	if strings.TrimSpace(profile.ProfileID) == "" {
		problems = append(problems, "profile_id is required")
	}
	if strings.TrimSpace(profile.StateID) == "" {
		problems = append(problems, "state_id is required")
	}
	if !twoDigitFIPS.MatchString(profile.StateFIPS) {
		problems = append(problems, "state_fips must be exactly two digits")
	}
	if _, err := time.Parse("2006-01-02", profile.LastVerified); err != nil {
		problems = append(problems, "last_verified must use YYYY-MM-DD")
	}

	sourceIDs := map[string]struct{}{}
	for _, source := range profile.Sources {
		if source.ID == "" {
			problems = append(problems, "source id is required")
			continue
		}
		if _, exists := sourceIDs[source.ID]; exists {
			problems = append(problems, "duplicate source id "+source.ID)
		}
		sourceIDs[source.ID] = struct{}{}
		parsed, err := url.Parse(source.URL)
		if err != nil || (parsed.Scheme != "http" && parsed.Scheme != "https") || parsed.Host == "" {
			problems = append(problems, "source "+source.ID+" must use an absolute http(s) URL")
		}
		if _, err := time.Parse("2006-01-02", source.AccessedAt); err != nil {
			problems = append(problems, "source "+source.ID+" accessed_at must use YYYY-MM-DD")
		}
	}

	authorityIDs := map[string]struct{}{}
	for _, authority := range profile.Authorities {
		if authority.ID == "" {
			problems = append(problems, "authority id is required")
			continue
		}
		if _, exists := authorityIDs[authority.ID]; exists {
			problems = append(problems, "duplicate authority id "+authority.ID)
		}
		authorityIDs[authority.ID] = struct{}{}
		problems = append(problems, missingSources("authority "+authority.ID, authority.SourceIDs, sourceIDs)...)
	}

	adoptionIDs := map[string]struct{}{}
	for _, adoption := range profile.Adoptions {
		if adoption.ID == "" {
			problems = append(problems, "adoption id is required")
			continue
		}
		if _, exists := adoptionIDs[adoption.ID]; exists {
			problems = append(problems, "duplicate adoption id "+adoption.ID)
		}
		adoptionIDs[adoption.ID] = struct{}{}
		if _, ok := authorityIDs[adoption.AdoptingAuthorityID]; !ok {
			problems = append(problems, "adoption "+adoption.ID+" references unknown authority "+adoption.AdoptingAuthorityID)
		}
		problems = append(problems, missingSources("adoption "+adoption.ID, adoption.SourceIDs, sourceIDs)...)
	}

	for _, relationship := range profile.Relationships {
		if _, ok := authorityIDs[relationship.FromID]; !ok {
			problems = append(problems, "relationship "+relationship.ID+" references unknown from_id "+relationship.FromID)
		}
		problems = append(problems, missingSources("relationship "+relationship.ID, relationship.SourceIDs, sourceIDs)...)
	}

	validatePolicy := func(name string, policy ResolutionPolicy) {
		if !validResolutionStatus(policy.Status) {
			problems = append(problems, name+" has invalid status "+policy.Status)
		}
		problems = append(problems, missingSources(name, policy.SourceIDs, sourceIDs)...)
		for _, candidate := range policy.AuthorityCandidates {
			if candidate.Kind == "state_authority" {
				if _, ok := authorityIDs[candidate.AuthorityID]; !ok {
					problems = append(problems, name+" references unknown authority "+candidate.AuthorityID)
				}
			}
			problems = append(problems, missingSources(name+" candidate "+candidate.Kind, candidate.SourceIDs, sourceIDs)...)
		}
		for _, adoptionID := range policy.AdoptionIDs {
			if _, ok := adoptionIDs[adoptionID]; !ok {
				problems = append(problems, name+" references unknown adoption "+adoptionID)
			}
		}
	}
	validatePolicy("defaults.incorporated", profile.Defaults.Incorporated)
	validatePolicy("defaults.unincorporated", profile.Defaults.Unincorporated)
	for key, policy := range profile.CodeFamilyOverrides {
		validatePolicy("code_family_overrides."+key, policy)
	}
	for key, policy := range profile.ProjectTypeOverrides {
		validatePolicy("project_type_overrides."+key, policy)
	}

	if len(problems) > 0 {
		sort.Strings(problems)
		return fmt.Errorf("invalid profile %s: %s", profile.ProfileID, strings.Join(problems, "; "))
	}
	return nil
}

func missingSources(owner string, ids []string, known map[string]struct{}) []string {
	if len(ids) == 0 {
		return []string{owner + " must cite at least one source"}
	}
	var problems []string
	for _, id := range ids {
		if _, ok := known[id]; !ok {
			problems = append(problems, owner+" references unknown source "+id)
		}
	}
	return problems
}

func validResolutionStatus(status string) bool {
	switch status {
	case "resolved", "partially_resolved", "local_record_required", "ambiguous", "conflicting", "insufficient_evidence":
		return true
	default:
		return false
	}
}
