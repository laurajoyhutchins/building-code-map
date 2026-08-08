package regulatory

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func ValidateCompleteProfile(profile StateProfile) error {
	if err := ValidateProfile(profile); err != nil {
		return err
	}
	var problems []string
	if !validProfileStatusExtended(profile.Status) {
		problems = append(problems, "profile status is invalid: "+profile.Status)
	}
	validateExtendedVerification("profile", profile.Verification, &problems)

	sourceIDs := map[string]struct{}{}
	for _, source := range profile.Sources {
		sourceIDs[source.ID] = struct{}{}
		validateExtendedDate("source "+source.ID+" last_checked_at", source.LastCheckedAt, false, &problems)
		if source.Availability != "" && !validAvailabilityExtended(source.Availability) {
			problems = append(problems, "source "+source.ID+" availability is invalid: "+source.Availability)
		}
	}
	authorityIDs := map[string]struct{}{}
	for _, authority := range profile.Authorities {
		authorityIDs[authority.ID] = struct{}{}
		validateExtendedVerification("authority "+authority.ID, authority.Verification, &problems)
	}
	adoptionIDs := map[string]struct{}{}
	for _, adoption := range profile.Adoptions {
		adoptionIDs[adoption.ID] = struct{}{}
		if !validAdoptionStatusExtended(adoption.Status) {
			problems = append(problems, "adoption "+adoption.ID+" status is invalid: "+adoption.Status)
		}
		validateExtendedVerification("adoption "+adoption.ID, adoption.Verification, &problems)
		validateExtendedDate("adoption "+adoption.ID+" adoption_date", adoption.Dates.AdoptionDate, false, &problems)
		validateExtendedDate("adoption "+adoption.ID+" effective_date", adoption.Dates.EffectiveDate, false, &problems)
		validateExtendedDate("adoption "+adoption.ID+" operative_date", adoption.Dates.OperativeDate, false, &problems)
		validateExtendedDate("adoption "+adoption.ID+" mandatory_date", adoption.Dates.MandatoryDate, false, &problems)
		validateExtendedDate("adoption "+adoption.ID+" replacement_date", adoption.Dates.ReplacementDate, false, &problems)
	}
	relationshipIDs := map[string]struct{}{}
	for _, relationship := range profile.Relationships {
		if _, exists := relationshipIDs[relationship.ID]; exists {
			problems = append(problems, "duplicate relationship id "+relationship.ID)
		}
		relationshipIDs[relationship.ID] = struct{}{}
		validateExtendedVerification("relationship "+relationship.ID, relationship.Verification, &problems)
	}

	ruleIDs := map[string]struct{}{}
	for _, rule := range profile.ApplicabilityRules {
		registerExtendedRule(rule.ID, "applicability rule", rule.SourceIDs, rule.Verification, sourceIDs, ruleIDs, &problems)
		if strings.TrimSpace(rule.Trigger) == "" || strings.TrimSpace(rule.Summary) == "" {
			problems = append(problems, "applicability rule "+rule.ID+" requires trigger and summary")
		}
	}
	for _, rule := range profile.DateRules {
		registerExtendedRule(rule.ID, "date rule", rule.SourceIDs, rule.Verification, sourceIDs, ruleIDs, &problems)
		if strings.TrimSpace(rule.RuleType) == "" || strings.TrimSpace(rule.Trigger) == "" || strings.TrimSpace(rule.Summary) == "" {
			problems = append(problems, "date rule "+rule.ID+" requires rule_type, trigger, and summary")
		}
		validateExtendedDate("date rule "+rule.ID+" start_date", rule.StartDate, false, &problems)
		validateExtendedDate("date rule "+rule.ID+" end_date", rule.EndDate, false, &problems)
		if rule.StartDate != "" && rule.EndDate != "" && rule.StartDate > rule.EndDate {
			problems = append(problems, "date rule "+rule.ID+" start_date must not follow end_date")
		}
	}
	for _, rule := range profile.AmendmentRules {
		registerExtendedRule(rule.ID, "amendment rule", rule.SourceIDs, rule.Verification, sourceIDs, ruleIDs, &problems)
		if strings.TrimSpace(rule.Level) == "" || strings.TrimSpace(rule.Posture) == "" || strings.TrimSpace(rule.Summary) == "" {
			problems = append(problems, "amendment rule "+rule.ID+" requires level, posture, and summary")
		}
		if rule.AuthorityID != "" {
			if _, ok := authorityIDs[rule.AuthorityID]; !ok {
				problems = append(problems, "amendment rule "+rule.ID+" references unknown authority "+rule.AuthorityID)
			}
		}
	}
	for _, rule := range profile.EnforcementRules {
		registerExtendedRule(rule.ID, "enforcement rule", rule.SourceIDs, rule.Verification, sourceIDs, ruleIDs, &problems)
		if strings.TrimSpace(rule.Model) == "" || strings.TrimSpace(rule.Summary) == "" || len(rule.EntityKinds) == 0 {
			problems = append(problems, "enforcement rule "+rule.ID+" requires model, entity_kinds, and summary")
		}
		for _, authorityID := range rule.AuthorityIDs {
			if _, ok := authorityIDs[authorityID]; !ok {
				problems = append(problems, "enforcement rule "+rule.ID+" references unknown authority "+authorityID)
			}
		}
	}

	knownSubjects := map[string]struct{}{profile.ProfileID: {}}
	for id := range authorityIDs {
		knownSubjects[id] = struct{}{}
	}
	for id := range adoptionIDs {
		knownSubjects[id] = struct{}{}
	}
	for id := range relationshipIDs {
		knownSubjects[id] = struct{}{}
	}
	for id := range ruleIDs {
		knownSubjects[id] = struct{}{}
	}
	claimIDs := map[string]struct{}{}
	for _, claim := range profile.Claims {
		if strings.TrimSpace(claim.ID) == "" {
			problems = append(problems, "claim id is required")
		} else if _, exists := claimIDs[claim.ID]; exists {
			problems = append(problems, "duplicate claim id "+claim.ID)
		}
		claimIDs[claim.ID] = struct{}{}
		if _, ok := knownSubjects[claim.SubjectID]; !ok {
			problems = append(problems, "claim "+claim.ID+" references unknown subject "+claim.SubjectID)
		}
		if strings.TrimSpace(claim.Field) == "" || !validClaimStatusExtended(claim.Status) {
			problems = append(problems, "claim "+claim.ID+" requires a field and valid status")
		}
		if claim.Status == "conflicting" && strings.TrimSpace(claim.ConflictGroup) == "" {
			problems = append(problems, "conflicting claim "+claim.ID+" requires conflict_group")
		}
		validateExtendedVerification("claim "+claim.ID, claim.Verification, &problems)
		problems = append(problems, missingSources("claim "+claim.ID, claim.SourceIDs, sourceIDs)...)
	}
	fixtureIDs := map[string]struct{}{}
	for _, fixtureID := range profile.ResolverFixtureIDs {
		if strings.TrimSpace(fixtureID) == "" {
			problems = append(problems, "resolver fixture id must not be empty")
		} else if _, exists := fixtureIDs[fixtureID]; exists {
			problems = append(problems, "duplicate resolver fixture id "+fixtureID)
		}
		fixtureIDs[fixtureID] = struct{}{}
	}

	if len(problems) > 0 {
		sort.Strings(problems)
		return fmt.Errorf("invalid profile %s: %s", profile.ProfileID, strings.Join(problems, "; "))
	}
	return nil
}

func registerExtendedRule(id, kind string, sourceRefs []string, verification Verification, sourceIDs, ruleIDs map[string]struct{}, problems *[]string) {
	if strings.TrimSpace(id) == "" {
		*problems = append(*problems, kind+" id is required")
	} else if _, exists := ruleIDs[id]; exists {
		*problems = append(*problems, "duplicate "+kind+" id "+id)
	}
	ruleIDs[id] = struct{}{}
	validateExtendedVerification(kind+" "+id, verification, problems)
	*problems = append(*problems, missingSources(kind+" "+id, sourceRefs, sourceIDs)...)
}

func validateExtendedDate(owner, value string, required bool, problems *[]string) {
	if value == "" && !required {
		return
	}
	if _, err := time.Parse("2006-01-02", value); err != nil {
		*problems = append(*problems, owner+" must use YYYY-MM-DD")
	}
}

func validateExtendedVerification(owner string, verification Verification, problems *[]string) {
	if !validVerificationStatusExtended(verification.Status) {
		*problems = append(*problems, owner+" verification status is invalid: "+verification.Status)
	}
	if verification.Confidence < 0 || verification.Confidence > 1 {
		*problems = append(*problems, owner+" verification confidence must be between 0 and 1")
	}
}

func validProfileStatusExtended(status string) bool {
	switch status {
	case "draft", "partially_verified", "verified", "deprecated":
		return true
	default:
		return false
	}
}

func validVerificationStatusExtended(status string) bool {
	switch status {
	case "verified", "partially_verified", "needs_review", "unresolved":
		return true
	default:
		return false
	}
}

func validAvailabilityExtended(status string) bool {
	switch status {
	case "available", "unavailable", "moved", "unknown":
		return true
	default:
		return false
	}
}

func validAdoptionStatusExtended(status string) bool {
	switch status {
	case "current", "prior", "future", "pending", "superseded", "unknown":
		return true
	default:
		return false
	}
}

func validClaimStatusExtended(status string) bool {
	switch status {
	case "supported", "conflicting", "unknown", "not_applicable", "superseded":
		return true
	default:
		return false
	}
}
