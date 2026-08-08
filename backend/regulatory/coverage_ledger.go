package regulatory

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

const EntityCoverageLedgerSchemaVersion = "1.0"

type CoverageEvidenceStatus string

const (
	CoverageEvidenceVerifiedInheritance                  CoverageEvidenceStatus = "verified_inheritance"
	CoverageEvidenceVerifiedInheritanceWithLocalEvidence CoverageEvidenceStatus = "verified_inheritance_with_local_evidence"
	CoverageEvidenceClassifiedRequiresLocalEvidence      CoverageEvidenceStatus = "classified_requires_local_evidence"
	CoverageEvidenceClassifiedNotVerified                CoverageEvidenceStatus = "classified_not_verified"
	CoverageEvidenceUnresolved                           CoverageEvidenceStatus = "unresolved"
	CoverageEvidenceConflicting                          CoverageEvidenceStatus = "conflicting"
	CoverageEvidenceInactive                             CoverageEvidenceStatus = "inactive"
)

type CoverageAvailability struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}

type EntityCoveragePopulationSummary struct {
	KnownEntities int   `json:"known_entities"`
	Total         int64 `json:"total"`
	Classified    int64 `json:"classified"`
	Verified      int64 `json:"verified"`
	Unresolved    int64 `json:"unresolved"`
}

type EntityCoverageSummary struct {
	ExpectedEntities              int                                    `json:"expected_entities"`
	ActiveEntities                int                                    `json:"active_entities"`
	ClassifiedEntities            int                                    `json:"classified_entities"`
	VerifiedEntities              int                                    `json:"verified_entities"`
	InheritedEntities             int                                    `json:"inherited_entities"`
	RequiresLocalEvidenceEntities int                                    `json:"requires_local_evidence_entities"`
	UnresolvedEntities            int                                    `json:"unresolved_entities"`
	InactiveEntities              int                                    `json:"inactive_entities"`
	ConflictingEntities           int                                    `json:"conflicting_entities"`
	EntityTypeCounts              map[GovernmentalEntityType]int         `json:"entity_type_counts"`
	ClassificationCounts          map[JurisdictionClassification]int     `json:"classification_counts"`
	EvidenceStatusCounts          map[CoverageEvidenceStatus]int          `json:"evidence_status_counts"`
	Population                    EntityCoveragePopulationSummary         `json:"population"`
}

type EntityCoverageEntry struct {
	EntityID               string                       `json:"entity_id"`
	OfficialName           string                       `json:"official_name"`
	EntityType             GovernmentalEntityType       `json:"entity_type"`
	LegalStatus            LegalStatus                  `json:"legal_status"`
	Classification         JurisdictionClassification   `json:"classification"`
	EvidenceStatus         CoverageEvidenceStatus       `json:"evidence_status"`
	Inherited              bool                         `json:"inherited"`
	Verified               bool                         `json:"verified"`
	RequiresLocalEvidence  bool                         `json:"requires_local_evidence"`
	Population             *EntityPopulation            `json:"population,omitempty"`
	RequiredLocalEvidence  []string                     `json:"required_local_evidence,omitempty"`
	UnresolvedReasons      []string                     `json:"unresolved_reasons,omitempty"`
	UnresolvedDimensions   []string                     `json:"unresolved_dimensions,omitempty"`
	SourceIDs              []string                     `json:"source_ids"`
	ClassificationRecords  []EntityClassificationRecord `json:"classification_records,omitempty"`
}

type StateEntityCoverage struct {
	StateID            string                `json:"state_id"`
	RegimeID           string                `json:"regime_id"`
	RegimeLastVerified string                `json:"regime_last_verified"`
	Summary            EntityCoverageSummary `json:"summary"`
	Entities           []EntityCoverageEntry  `json:"entities"`
}

type EntityCoverageLedger struct {
	SchemaVersion      string                `json:"schema_version"`
	InventorySourceID  string                `json:"inventory_source_id"`
	InventoryGeneratedAt string              `json:"inventory_generated_at"`
	CodeFamily         string                `json:"code_family"`
	ProjectType        string                `json:"project_type,omitempty"`
	ApplicabilityDate  string                `json:"applicability_date"`
	StateIDs           []string              `json:"state_ids"`
	Summary            EntityCoverageSummary `json:"summary"`
	States             []StateEntityCoverage  `json:"states"`
	Freshness          CoverageAvailability   `json:"freshness"`
	LandArea           CoverageAvailability   `json:"land_area"`
}

type EntityCoverageLedgerRequest struct {
	Inventories       CensusEntityInventoryBuild
	Regimes           []StateRegimeSpec
	StateIDs          []string
	CodeFamily        string
	ProjectType       string
	ApplicabilityDate string
}

func BuildEntityCoverageLedger(request EntityCoverageLedgerRequest) (EntityCoverageLedger, error) {
	codeFamily := normalizeKey(request.CodeFamily)
	if codeFamily == "" {
		return EntityCoverageLedger{}, fmt.Errorf("code family is required")
	}
	projectType := normalizeKey(request.ProjectType)
	date := strings.TrimSpace(request.ApplicabilityDate)
	if _, err := time.Parse(time.DateOnly, date); err != nil {
		return EntityCoverageLedger{}, fmt.Errorf("applicability_date must use YYYY-MM-DD")
	}
	if request.Inventories.Index.SourceID == "" || request.Inventories.Index.GeneratedAt == "" {
		return EntityCoverageLedger{}, fmt.Errorf("inventory index identity is required")
	}
	stateIDs, err := normalizedCoverageStateIDs(request.StateIDs)
	if err != nil {
		return EntityCoverageLedger{}, err
	}
	regimes := make(map[string]StateRegimeSpec, len(request.Regimes))
	for _, regime := range request.Regimes {
		if err := ValidateStateRegimeSpec(regime); err != nil {
			return EntityCoverageLedger{}, fmt.Errorf("regime %s: %w", regime.StateID, err)
		}
		if _, duplicate := regimes[regime.StateID]; duplicate {
			return EntityCoverageLedger{}, fmt.Errorf("duplicate regime for %s", regime.StateID)
		}
		regimes[regime.StateID] = regime
	}

	ledger := EntityCoverageLedger{
		SchemaVersion:       EntityCoverageLedgerSchemaVersion,
		InventorySourceID:   request.Inventories.Index.SourceID,
		InventoryGeneratedAt: request.Inventories.Index.GeneratedAt,
		CodeFamily:          codeFamily,
		ProjectType:         projectType,
		ApplicabilityDate:   date,
		StateIDs:            stateIDs,
		Summary:             newEntityCoverageSummary(),
		States:              make([]StateEntityCoverage, 0, len(stateIDs)),
		Freshness: CoverageAvailability{
			Status: "not_assessed",
			Reason: "State regime projections expose last_verified but do not yet define a per-source staleness threshold or freshness contract.",
		},
		LandArea: CoverageAvailability{
			Status: "unavailable",
			Reason: "The canonical 2025 governmental-entity inventory does not yet carry source-backed land-area measures.",
		},
	}

	for _, stateID := range stateIDs {
		inventory, ok := request.Inventories.Inventories[stateID]
		if !ok {
			return EntityCoverageLedger{}, fmt.Errorf("no governmental entity inventory for %s", stateID)
		}
		if err := ValidateEntityInventory(inventory); err != nil {
			return EntityCoverageLedger{}, fmt.Errorf("inventory %s: %w", stateID, err)
		}
		regime, ok := regimes[stateID]
		if !ok {
			return EntityCoverageLedger{}, fmt.Errorf("no state regime for %s", stateID)
		}
		state, err := buildStateEntityCoverage(inventory, regime, codeFamily, projectType, date)
		if err != nil {
			return EntityCoverageLedger{}, fmt.Errorf("coverage %s: %w", stateID, err)
		}
		ledger.States = append(ledger.States, state)
		mergeCoverageSummary(&ledger.Summary, state.Summary)
	}
	return ledger, nil
}

func normalizedCoverageStateIDs(values []string) ([]string, error) {
	if len(values) == 0 {
		return nil, fmt.Errorf("at least one state id is required")
	}
	seen := map[string]bool{}
	result := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.ToUpper(strings.TrimSpace(value))
		if len(value) != 5 || !strings.HasPrefix(value, "US-") {
			return nil, fmt.Errorf("invalid state id %q", value)
		}
		if seen[value] {
			return nil, fmt.Errorf("duplicate state id %s", value)
		}
		seen[value] = true
		result = append(result, value)
	}
	sort.Strings(result)
	return result, nil
}

func buildStateEntityCoverage(inventory EntityInventory, regime StateRegimeSpec, codeFamily, projectType, date string) (StateEntityCoverage, error) {
	if regime.StateID == "" || len(inventory.Entities) == 0 {
		return StateEntityCoverage{}, fmt.Errorf("state regime and inventory entities are required")
	}
	entries := make([]EntityCoverageEntry, 0, len(inventory.Entities))
	summary := newEntityCoverageSummary()
	for _, entity := range inventory.Entities {
		entry, err := buildEntityCoverageEntry(entity, regime, codeFamily, projectType, date)
		if err != nil {
			return StateEntityCoverage{}, err
		}
		entries = append(entries, entry)
		addCoverageEntry(&summary, entry)
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].EntityID < entries[j].EntityID })
	return StateEntityCoverage{
		StateID:            regime.StateID,
		RegimeID:           regime.RegimeID,
		RegimeLastVerified: regime.LastVerified,
		Summary:            summary,
		Entities:           entries,
	}, nil
}

func buildEntityCoverageEntry(entity GovernmentalEntity, regime StateRegimeSpec, codeFamily, projectType, date string) (EntityCoverageEntry, error) {
	if conflictSources := coverageConflictSources(entity, regime, codeFamily, projectType); len(conflictSources) > 0 && entity.LegalStatus != LegalStatusInactive {
		return EntityCoverageEntry{
			EntityID:              entity.EntityID,
			OfficialName:          entity.OfficialName,
			EntityType:            entity.EntityType,
			LegalStatus:           entity.LegalStatus,
			Classification:        ClassificationUnresolved,
			EvidenceStatus:        CoverageEvidenceConflicting,
			Population:            entity.Population,
			UnresolvedReasons:     []string{"applicable regime policy is conflicting"},
			UnresolvedDimensions:  append([]string(nil), regime.UnresolvedDimensions...),
			SourceIDs:             sortedUnique(append(append([]string(nil), entity.SourceIDs...), conflictSources...)),
		}, nil
	}
	decision, err := ClassifyGovernmentalEntity(EntityClassificationRequest{
		Entity:            entity,
		Regime:            regime,
		CodeFamily:        codeFamily,
		ProjectType:       projectType,
		ApplicabilityDate: date,
	})
	if err != nil {
		return EntityCoverageEntry{}, err
	}
	inherited := hasClassificationRecordKindValue(decision.Records, ClassificationRecordStateInheritance)
	requiredLocal := coverageRequiresLocalEvidence(decision)
	verified := entity.LegalStatus == LegalStatusActive && inherited && regime.Verification.Status == "verified"
	status := coverageEvidenceStatus(entity, decision, inherited, verified, requiredLocal)
	required := coverageRequiredLocalEvidence(decision)
	if requiredLocal {
		required = append(required, regime.RequiredLocalEvidence...)
	}
	return EntityCoverageEntry{
		EntityID:              entity.EntityID,
		OfficialName:          entity.OfficialName,
		EntityType:            entity.EntityType,
		LegalStatus:           entity.LegalStatus,
		Classification:        decision.Classification,
		EvidenceStatus:        status,
		Inherited:             inherited,
		Verified:              verified,
		RequiresLocalEvidence: requiredLocal,
		Population:            entity.Population,
		RequiredLocalEvidence: sortedUnique(required),
		UnresolvedReasons:     append([]string(nil), decision.UnresolvedReasons...),
		UnresolvedDimensions:  append([]string(nil), decision.UnresolvedDimensions...),
		SourceIDs:             sortedUnique(append(append([]string(nil), entity.SourceIDs...), decision.SourceIDs...)),
		ClassificationRecords: append([]EntityClassificationRecord(nil), decision.Records...),
	}, nil
}

func coverageEvidenceStatus(entity GovernmentalEntity, decision EntityClassificationDecision, inherited, verified, requiredLocal bool) CoverageEvidenceStatus {
	if entity.LegalStatus == LegalStatusInactive || decision.Classification == ClassificationInactive {
		return CoverageEvidenceInactive
	}
	if decision.Classification == ClassificationUnresolved {
		return CoverageEvidenceUnresolved
	}
	if verified && inherited && requiredLocal {
		return CoverageEvidenceVerifiedInheritanceWithLocalEvidence
	}
	if verified && inherited {
		return CoverageEvidenceVerifiedInheritance
	}
	if requiredLocal {
		return CoverageEvidenceClassifiedRequiresLocalEvidence
	}
	return CoverageEvidenceClassifiedNotVerified
}

func coverageRequiresLocalEvidence(decision EntityClassificationDecision) bool {
	switch decision.Classification {
	case ClassificationLocalAdopter, ClassificationLocalAmender, ClassificationEnforcementOnly, ClassificationDelegated:
		return true
	}
	for _, record := range decision.Records {
		if len(record.RequiredLocalEvidence) > 0 {
			return true
		}
	}
	return false
}

func coverageRequiredLocalEvidence(decision EntityClassificationDecision) []string {
	var required []string
	for _, record := range decision.Records {
		required = append(required, record.RequiredLocalEvidence...)
	}
	return sortedUnique(required)
}

func coverageConflictSources(entity GovernmentalEntity, regime StateRegimeSpec, codeFamily, projectType string) []string {
	var sources []string
	if policy, ok := regime.CodeFamilyPolicies[codeFamily]; ok && policy.Status == "conflicting" {
		sources = append(sources, policy.SourceIDs...)
	}
	if projectType != "" {
		if policy, ok := regime.ProjectTypeExceptions[projectType]; ok && policy.Status == "conflicting" {
			sources = append(sources, policy.SourceIDs...)
		}
	}
	for _, scope := range classificationTerritoryScopes(entity.EntityType) {
		var policy ResolutionPolicy
		switch scope {
		case "incorporated":
			policy = regime.Territory.Incorporated
		case "unincorporated":
			policy = regime.Territory.Unincorporated
		}
		if policy.Status == "conflicting" {
			sources = append(sources, policy.SourceIDs...)
		}
	}
	return sortedUnique(sources)
}

func hasClassificationRecordKindValue(records []EntityClassificationRecord, kind ClassificationRecordKind) bool {
	for _, record := range records {
		if record.Kind == kind {
			return true
		}
	}
	return false
}

func newEntityCoverageSummary() EntityCoverageSummary {
	return EntityCoverageSummary{
		EntityTypeCounts:     map[GovernmentalEntityType]int{},
		ClassificationCounts: map[JurisdictionClassification]int{},
		EvidenceStatusCounts: map[CoverageEvidenceStatus]int{},
	}
}

func addCoverageEntry(summary *EntityCoverageSummary, entry EntityCoverageEntry) {
	summary.ExpectedEntities++
	summary.EntityTypeCounts[entry.EntityType]++
	summary.ClassificationCounts[entry.Classification]++
	summary.EvidenceStatusCounts[entry.EvidenceStatus]++
	if entry.LegalStatus == LegalStatusActive {
		summary.ActiveEntities++
	}
	if entry.Classification != ClassificationUnresolved {
		summary.ClassifiedEntities++
	}
	if entry.Verified {
		summary.VerifiedEntities++
	}
	if entry.Inherited {
		summary.InheritedEntities++
	}
	if entry.RequiresLocalEvidence {
		summary.RequiresLocalEvidenceEntities++
	}
	if entry.Classification == ClassificationUnresolved {
		summary.UnresolvedEntities++
	}
	if entry.Classification == ClassificationInactive {
		summary.InactiveEntities++
	}
	if entry.EvidenceStatus == CoverageEvidenceConflicting {
		summary.ConflictingEntities++
	}
	if entry.Population != nil && entry.LegalStatus == LegalStatusActive {
		summary.Population.KnownEntities++
		summary.Population.Total += entry.Population.Count
		if entry.Classification != ClassificationUnresolved {
			summary.Population.Classified += entry.Population.Count
		}
		if entry.Verified {
			summary.Population.Verified += entry.Population.Count
		}
		if entry.Classification == ClassificationUnresolved {
			summary.Population.Unresolved += entry.Population.Count
		}
	}
}

func mergeCoverageSummary(target *EntityCoverageSummary, source EntityCoverageSummary) {
	target.ExpectedEntities += source.ExpectedEntities
	target.ActiveEntities += source.ActiveEntities
	target.ClassifiedEntities += source.ClassifiedEntities
	target.VerifiedEntities += source.VerifiedEntities
	target.InheritedEntities += source.InheritedEntities
	target.RequiresLocalEvidenceEntities += source.RequiresLocalEvidenceEntities
	target.UnresolvedEntities += source.UnresolvedEntities
	target.InactiveEntities += source.InactiveEntities
	target.ConflictingEntities += source.ConflictingEntities
	for key, value := range source.EntityTypeCounts {
		target.EntityTypeCounts[key] += value
	}
	for key, value := range source.ClassificationCounts {
		target.ClassificationCounts[key] += value
	}
	for key, value := range source.EvidenceStatusCounts {
		target.EvidenceStatusCounts[key] += value
	}
	target.Population.KnownEntities += source.Population.KnownEntities
	target.Population.Total += source.Population.Total
	target.Population.Classified += source.Population.Classified
	target.Population.Verified += source.Population.Verified
	target.Population.Unresolved += source.Population.Unresolved
}
