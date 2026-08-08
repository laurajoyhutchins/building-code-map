package regulatory

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

const EntityClassificationDecisionSchemaVersion = "1.0"

type ClassificationRecordKind string

const (
	ClassificationRecordLegalStatus      ClassificationRecordKind = "legal_status"
	ClassificationRecordLocalPower       ClassificationRecordKind = "local_power"
	ClassificationRecordEnforcement      ClassificationRecordKind = "enforcement"
	ClassificationRecordStateInheritance ClassificationRecordKind = "state_inheritance"
)

type EntityClassificationRequest struct {
	Entity            GovernmentalEntity
	Regime            StateRegimeSpec
	CodeFamily        string
	ProjectType       string
	ApplicabilityDate string
}

type EntityClassificationRecord struct {
	RecordID              string                     `json:"record_id"`
	Kind                  ClassificationRecordKind   `json:"kind"`
	Classification        JurisdictionClassification `json:"classification"`
	CodeFamily            string                     `json:"code_family,omitempty"`
	ProjectType           string                     `json:"project_type,omitempty"`
	Scope                 string                     `json:"scope,omitempty"`
	Powers                []RegimePower              `json:"powers,omitempty"`
	AdoptionIDs           []string                   `json:"adoption_ids,omitempty"`
	RuleIDs               []string                   `json:"rule_ids,omitempty"`
	PolicyStatus          string                     `json:"policy_status,omitempty"`
	LegalBasisSourceIDs   []string                   `json:"legal_basis_source_ids"`
	EffectiveInterval     EntityEffectiveInterval    `json:"effective_interval,omitempty"`
	Exclusions            []string                   `json:"exclusions,omitempty"`
	RequiredLocalEvidence []string                   `json:"required_local_evidence,omitempty"`
}

type EntityClassificationDecision struct {
	SchemaVersion      string                       `json:"schema_version"`
	EntityID           string                       `json:"entity_id"`
	StateID            string                       `json:"state_id"`
	Classification     JurisdictionClassification   `json:"classification"`
	CodeFamily         string                       `json:"code_family,omitempty"`
	ProjectType        string                       `json:"project_type,omitempty"`
	ApplicabilityDate  string                       `json:"applicability_date"`
	Records            []EntityClassificationRecord `json:"records"`
	UnresolvedReasons  []string                     `json:"unresolved_reasons,omitempty"`
	UnresolvedDimensions []string                   `json:"unresolved_dimensions,omitempty"`
	SourceIDs          []string                     `json:"source_ids"`
}

func ClassifyGovernmentalEntity(request EntityClassificationRequest) (EntityClassificationDecision, error) {
	if err := validateGovernmentalEntity(request.Entity); err != nil {
		return EntityClassificationDecision{}, fmt.Errorf("entity: %w", err)
	}
	if err := ValidateStateRegimeSpec(request.Regime); err != nil {
		return EntityClassificationDecision{}, fmt.Errorf("regime: %w", err)
	}
	if request.Entity.StateID != request.Regime.StateID {
		return EntityClassificationDecision{}, fmt.Errorf("entity state %s does not match regime state %s", request.Entity.StateID, request.Regime.StateID)
	}

	codeFamily := normalizeKey(request.CodeFamily)
	projectType := normalizeKey(request.ProjectType)
	applicabilityDate, err := classificationApplicabilityDate(request.ApplicabilityDate, request.Regime.LastVerified)
	if err != nil {
		return EntityClassificationDecision{}, err
	}

	decision := EntityClassificationDecision{
		SchemaVersion:        EntityClassificationDecisionSchemaVersion,
		EntityID:             request.Entity.EntityID,
		StateID:              request.Entity.StateID,
		Classification:       ClassificationUnresolved,
		CodeFamily:           codeFamily,
		ProjectType:          projectType,
		ApplicabilityDate:    applicabilityDate,
		Records:              []EntityClassificationRecord{},
		UnresolvedDimensions: append([]string(nil), request.Regime.UnresolvedDimensions...),
		SourceIDs:            []string{},
	}

	if request.Entity.LegalStatus == LegalStatusInactive || request.Entity.Classification == ClassificationInactive {
		record := EntityClassificationRecord{
			RecordID:            classificationRecordID(request.Entity.EntityID, ClassificationRecordLegalStatus, codeFamily, projectType),
			Kind:                ClassificationRecordLegalStatus,
			Classification:      ClassificationInactive,
			CodeFamily:          codeFamily,
			ProjectType:         projectType,
			LegalBasisSourceIDs: sortedUnique(request.Entity.SourceIDs),
		}
		decision.Classification = ClassificationInactive
		decision.Records = append(decision.Records, record)
		decision.SourceIDs = sortedUnique(request.Entity.SourceIDs)
		return decision, nil
	}

	if request.Entity.LegalStatus == LegalStatusUnknown {
		decision.UnresolvedReasons = append(decision.UnresolvedReasons, "entity legal status is unknown")
	}

	inheritance, inheritanceReason := stateInheritanceRecord(request.Entity, request.Regime, codeFamily, projectType, applicabilityDate)
	if inheritance != nil {
		decision.Records = append(decision.Records, *inheritance)
		decision.SourceIDs = append(decision.SourceIDs, inheritance.LegalBasisSourceIDs...)
	} else if inheritanceReason != "" {
		decision.UnresolvedReasons = append(decision.UnresolvedReasons, inheritanceReason)
	}

	localRecords := explicitLocalPowerRecords(request.Entity, request.Regime, codeFamily, projectType)
	decision.Records = append(decision.Records, localRecords...)
	for _, record := range localRecords {
		decision.SourceIDs = append(decision.SourceIDs, record.LegalBasisSourceIDs...)
	}

	enforcementRecords := explicitEnforcementRecords(request.Entity, request.Regime, codeFamily, projectType)
	decision.Records = append(decision.Records, enforcementRecords...)
	for _, record := range enforcementRecords {
		decision.SourceIDs = append(decision.SourceIDs, record.LegalBasisSourceIDs...)
	}

	decision.Classification = chooseEntityClassification(localRecords, enforcementRecords, inheritance)
	if decision.Classification == ClassificationUnresolved && len(decision.UnresolvedReasons) == 0 {
		decision.UnresolvedReasons = append(decision.UnresolvedReasons, "regime evidence does not classify this entity for the requested context")
	}
	decision.UnresolvedReasons = sortedUnique(decision.UnresolvedReasons)
	decision.UnresolvedDimensions = sortedUnique(decision.UnresolvedDimensions)
	decision.SourceIDs = sortedUnique(decision.SourceIDs)
	sortClassificationRecords(decision.Records)
	return decision, nil
}

func classificationApplicabilityDate(value, fallback string) (string, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		value = strings.TrimSpace(fallback)
	}
	if value == "" {
		return "", fmt.Errorf("applicability date is required when regime last_verified is empty")
	}
	if _, err := time.Parse(time.DateOnly, value); err != nil {
		return "", fmt.Errorf("applicability_date must use YYYY-MM-DD")
	}
	return value, nil
}

func stateInheritanceRecord(entity GovernmentalEntity, regime StateRegimeSpec, codeFamily, projectType, applicabilityDate string) (*EntityClassificationRecord, string) {
	if codeFamily == "" {
		return nil, "state inheritance requires an explicit code family"
	}
	policy, ok := regime.CodeFamilyPolicies[codeFamily]
	if !ok || len(policy.AdoptionIDs) == 0 {
		return nil, "state inheritance is not explicitly established for the requested code family"
	}
	if projectType != "" {
		if override, exists := regime.ProjectTypeExceptions[projectType]; exists {
			policy = mergePolicies(policy, override)
		}
	}
	if len(policy.AdoptionIDs) == 0 {
		return nil, "project exception removes the statewide adoption basis for the requested context"
	}

	dateRules := applicableClassificationDateRules(regime.DateRules, codeFamily, applicabilityDate)
	interval, ok := intersectClassificationIntervals(dateRules)
	if !ok {
		return nil, "state inheritance lacks an applicable temporal rule"
	}

	legalBasis := append([]string(nil), policy.SourceIDs...)
	ruleIDs := make([]string, 0, len(dateRules))
	for _, rule := range dateRules {
		legalBasis = append(legalBasis, rule.SourceIDs...)
		ruleIDs = append(ruleIDs, rule.ID)
	}
	exclusions := make([]string, 0)
	for _, rule := range regime.ApplicabilityRules {
		if !familyMatches(rule.CodeFamily, codeFamily) || !projectTypesMatch(rule.ProjectTypes, projectType) {
			continue
		}
		exclusions = append(exclusions, rule.Exclusions...)
		legalBasis = append(legalBasis, rule.SourceIDs...)
		ruleIDs = append(ruleIDs, rule.ID)
	}

	return &EntityClassificationRecord{
		RecordID:              classificationRecordID(entity.EntityID, ClassificationRecordStateInheritance, codeFamily, projectType),
		Kind:                  ClassificationRecordStateInheritance,
		Classification:        ClassificationCoveredByState,
		CodeFamily:            codeFamily,
		ProjectType:           projectType,
		Scope:                 "statewide",
		AdoptionIDs:           sortedUnique(policy.AdoptionIDs),
		RuleIDs:               sortedUnique(ruleIDs),
		PolicyStatus:          policy.Status,
		LegalBasisSourceIDs:   sortedUnique(legalBasis),
		EffectiveInterval:     interval,
		Exclusions:            sortedUnique(exclusions),
		RequiredLocalEvidence: sortedUnique(policy.RequiredLocalRecords),
	}, ""
}

func applicableClassificationDateRules(rules []DateRule, codeFamily, applicabilityDate string) []DateRule {
	result := make([]DateRule, 0)
	for _, rule := range rules {
		if !familyMatches(rule.CodeFamily, codeFamily) || !dateRuleApplies(rule, applicabilityDate) {
			continue
		}
		if strings.TrimSpace(rule.StartDate) == "" && strings.TrimSpace(rule.EndDate) == "" {
			continue
		}
		result = append(result, rule)
	}
	sort.Slice(result, func(i, j int) bool { return result[i].ID < result[j].ID })
	return result
}

func intersectClassificationIntervals(rules []DateRule) (EntityEffectiveInterval, bool) {
	if len(rules) == 0 {
		return EntityEffectiveInterval{}, false
	}
	var latestStart time.Time
	var earliestEnd time.Time
	for _, rule := range rules {
		if rule.StartDate != "" {
			start, err := time.Parse(time.DateOnly, rule.StartDate)
			if err != nil {
				return EntityEffectiveInterval{}, false
			}
			if latestStart.IsZero() || start.After(latestStart) {
				latestStart = start
			}
		}
		if rule.EndDate != "" {
			end, err := time.Parse(time.DateOnly, rule.EndDate)
			if err != nil {
				return EntityEffectiveInterval{}, false
			}
			if earliestEnd.IsZero() || end.Before(earliestEnd) {
				earliestEnd = end
			}
		}
	}
	if !latestStart.IsZero() && !earliestEnd.IsZero() && earliestEnd.Before(latestStart) {
		return EntityEffectiveInterval{}, false
	}
	interval := EntityEffectiveInterval{}
	if !latestStart.IsZero() {
		interval.StartDate = latestStart.Format(time.DateOnly)
	}
	if !earliestEnd.IsZero() {
		interval.EndDate = earliestEnd.Format(time.DateOnly)
	}
	if interval == (EntityEffectiveInterval{}) {
		return EntityEffectiveInterval{}, false
	}
	return interval, true
}

func explicitLocalPowerRecords(entity GovernmentalEntity, regime StateRegimeSpec, codeFamily, projectType string) []EntityClassificationRecord {
	stateBaseline := false
	if policy, ok := regime.CodeFamilyPolicies[codeFamily]; codeFamily != "" && ok && len(policy.AdoptionIDs) > 0 {
		stateBaseline = true
	}
	aliases := classificationEntityKinds(entity.EntityType)
	territoryScopes := classificationTerritoryScopes(entity.EntityType)
	records := make([]EntityClassificationRecord, 0)
	for _, evidence := range regime.EntityPowerEvidence {
		if !containsAnyString(aliases, strings.TrimSpace(evidence.EntityKind)) {
			continue
		}
		scope := strings.TrimSpace(evidence.Scope)
		isCodeFamilyScope := codeFamily != "" && scope == "code_family:"+codeFamily
		isProjectScope := projectType != "" && scope == "project_type:"+projectType
		isTerritoryScope := containsAnyString(territoryScopes, scope)
		if !isCodeFamilyScope && !isProjectScope && !isTerritoryScope {
			continue
		}

		powers := append([]RegimePower(nil), evidence.Powers...)
		if stateBaseline && isTerritoryScope {
			powers = removeClassificationPowers(powers, RegimePowerAdopt, RegimePowerAmend)
		}
		classification := classificationFromPowers(powers)
		if classification == ClassificationUnresolved {
			continue
		}
		records = append(records, EntityClassificationRecord{
			RecordID:            classificationRecordID(entity.EntityID, ClassificationRecordLocalPower, codeFamily, projectType)+":"+sanitizeClassificationID(scope),
			Kind:                ClassificationRecordLocalPower,
			Classification:      classification,
			CodeFamily:          codeFamily,
			ProjectType:         projectType,
			Scope:               scope,
			Powers:              powers,
			LegalBasisSourceIDs: sortedUnique(evidence.SourceIDs),
		})
	}
	return records
}

func explicitEnforcementRecords(entity GovernmentalEntity, regime StateRegimeSpec, codeFamily, projectType string) []EntityClassificationRecord {
	aliases := classificationEntityKinds(entity.EntityType)
	records := make([]EntityClassificationRecord, 0)
	for _, rule := range regime.EnforcementRules {
		if !familyMatches(rule.CodeFamily, codeFamily) {
			continue
		}
		if !classificationRuleKindsMatch(rule.EntityKinds, aliases) {
			continue
		}
		records = append(records, EntityClassificationRecord{
			RecordID:            classificationRecordID(entity.EntityID, ClassificationRecordEnforcement, codeFamily, projectType)+":"+sanitizeClassificationID(rule.ID),
			Kind:                ClassificationRecordEnforcement,
			Classification:      ClassificationEnforcementOnly,
			CodeFamily:          codeFamily,
			ProjectType:         projectType,
			Scope:               strings.TrimSpace(rule.Model),
			Powers:              []RegimePower{RegimePowerAdminister, RegimePowerEnforce},
			RuleIDs:             []string{rule.ID},
			LegalBasisSourceIDs: sortedUnique(rule.SourceIDs),
		})
	}
	return records
}

func chooseEntityClassification(localRecords, enforcementRecords []EntityClassificationRecord, inheritance *EntityClassificationRecord) JurisdictionClassification {
	best := ClassificationUnresolved
	for _, record := range localRecords {
		switch record.Classification {
		case ClassificationLocalAdopter:
			return ClassificationLocalAdopter
		case ClassificationLocalAmender:
			best = ClassificationLocalAmender
		case ClassificationEnforcementOnly:
			if best == ClassificationUnresolved {
				best = ClassificationEnforcementOnly
			}
		}
	}
	if best != ClassificationUnresolved {
		return best
	}
	if len(enforcementRecords) > 0 {
		return ClassificationEnforcementOnly
	}
	if inheritance != nil {
		return ClassificationCoveredByState
	}
	return ClassificationUnresolved
}

func classificationFromPowers(powers []RegimePower) JurisdictionClassification {
	if containsClassificationPower(powers, RegimePowerAdopt) {
		return ClassificationLocalAdopter
	}
	if containsClassificationPower(powers, RegimePowerAmend) {
		return ClassificationLocalAmender
	}
	if containsClassificationPower(powers, RegimePowerAdminister) || containsClassificationPower(powers, RegimePowerEnforce) {
		return ClassificationEnforcementOnly
	}
	return ClassificationUnresolved
}

func classificationEntityKinds(entityType GovernmentalEntityType) []string {
	switch entityType {
	case EntityTypeMunicipality, EntityTypeIndependentCity:
		return []string{"municipality"}
	case EntityTypeCountyEquivalent:
		return []string{"county"}
	case EntityTypeMinorCivilDivision:
		return []string{"minor_civil_division", "township"}
	case EntityTypeConsolidatedGovernment:
		return []string{"consolidated_government", "municipality", "county"}
	case EntityTypeSpecialDistrict:
		return []string{"special_district", "special_authority"}
	case EntityTypeTribalGovernment:
		return []string{"tribal_government", "tribal_authority"}
	default:
		return []string{string(entityType)}
	}
}

func classificationTerritoryScopes(entityType GovernmentalEntityType) []string {
	switch entityType {
	case EntityTypeMunicipality, EntityTypeIndependentCity:
		return []string{"incorporated"}
	case EntityTypeCountyEquivalent, EntityTypeMinorCivilDivision:
		return []string{"unincorporated"}
	case EntityTypeConsolidatedGovernment:
		return []string{"incorporated", "unincorporated"}
	default:
		return nil
	}
}

func classificationRuleKindsMatch(ruleKinds, aliases []string) bool {
	for _, ruleKind := range ruleKinds {
		normalized := normalizeKey(ruleKind)
		for _, alias := range aliases {
			alias = normalizeKey(alias)
			if normalized == alias || strings.HasPrefix(normalized, alias+"_") || strings.HasPrefix(normalized, strings.TrimSuffix(alias, "ality")+"al_") {
				return true
			}
			if alias == "municipality" && strings.HasPrefix(normalized, "municipal_") {
				return true
			}
		}
	}
	return false
}

func removeClassificationPowers(values []RegimePower, remove ...RegimePower) []RegimePower {
	blocked := map[RegimePower]bool{}
	for _, value := range remove {
		blocked[value] = true
	}
	result := make([]RegimePower, 0, len(values))
	for _, value := range values {
		if !blocked[value] {
			result = append(result, value)
		}
	}
	return result
}

func containsClassificationPower(values []RegimePower, want RegimePower) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}

func containsAnyString(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
}

func classificationRecordID(entityID string, kind ClassificationRecordKind, codeFamily, projectType string) string {
	parts := []string{"classification", entityID, string(kind)}
	if codeFamily != "" {
		parts = append(parts, codeFamily)
	}
	if projectType != "" {
		parts = append(parts, projectType)
	}
	return strings.Join(parts, ":")
}

func sanitizeClassificationID(value string) string {
	value = normalizeKey(value)
	if value == "" {
		return "general"
	}
	return value
}

func sortClassificationRecords(records []EntityClassificationRecord) {
	sort.Slice(records, func(i, j int) bool {
		if records[i].Kind != records[j].Kind {
			return records[i].Kind < records[j].Kind
		}
		return records[i].RecordID < records[j].RecordID
	})
}
