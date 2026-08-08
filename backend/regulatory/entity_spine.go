package regulatory

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

const (
	GovernmentalEntitySchemaVersion = "1.0"
	EntityInventorySchemaVersion    = "1.0"
)

type GovernmentalEntityType string

const (
	EntityTypeState                  GovernmentalEntityType = "state"
	EntityTypeCountyEquivalent       GovernmentalEntityType = "county_equivalent"
	EntityTypeMunicipality           GovernmentalEntityType = "municipality"
	EntityTypeMinorCivilDivision     GovernmentalEntityType = "minor_civil_division"
	EntityTypeConsolidatedGovernment GovernmentalEntityType = "consolidated_government"
	EntityTypeIndependentCity        GovernmentalEntityType = "independent_city"
	EntityTypeSpecialDistrict        GovernmentalEntityType = "special_district"
	EntityTypeTribalGovernment       GovernmentalEntityType = "tribal_government"
	EntityTypeOtherGovernmentalEntity GovernmentalEntityType = "other_governmental_entity"
)

type JurisdictionClassification string

const (
	ClassificationCoveredByState    JurisdictionClassification = "covered_by_state"
	ClassificationLocalAdopter      JurisdictionClassification = "local_adopter"
	ClassificationLocalAmender      JurisdictionClassification = "local_amender"
	ClassificationEnforcementOnly   JurisdictionClassification = "enforcement_only"
	ClassificationDelegated         JurisdictionClassification = "delegated"
	ClassificationNoGeneralAuthority JurisdictionClassification = "no_general_authority"
	ClassificationInactive          JurisdictionClassification = "inactive"
	ClassificationNotApplicable     JurisdictionClassification = "not_applicable"
	ClassificationUnresolved        JurisdictionClassification = "unresolved"
)

type LegalStatus string

const (
	LegalStatusActive   LegalStatus = "active"
	LegalStatusInactive LegalStatus = "inactive"
	LegalStatusUnknown  LegalStatus = "unknown"
)

type HistoricalGeographyStatus string

const (
	HistoricalGeographySupported   HistoricalGeographyStatus = "supported"
	HistoricalGeographyCurrentOnly HistoricalGeographyStatus = "current_only"
	HistoricalGeographyUnavailable HistoricalGeographyStatus = "unavailable"
)

type CanonicalEntityIdentity struct {
	Namespace string `json:"namespace"`
	Value     string `json:"value"`
}

type ExternalEntityIdentifier struct {
	Scheme string `json:"scheme"`
	Value  string `json:"value"`
}

type EntityEffectiveInterval struct {
	StartDate string `json:"start_date,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
}

type EntityGeographyReference struct {
	Kind       string `json:"kind"`
	Identifier string `json:"identifier"`
	Vintage    string `json:"vintage"`
}

type GovernmentalEntity struct {
	SchemaVersion             string                     `json:"schema_version"`
	EntityID                  string                     `json:"entity_id"`
	OfficialName              string                     `json:"official_name"`
	EntityType                GovernmentalEntityType     `json:"entity_type"`
	StateID                   string                     `json:"state_id"`
	StateFIPS                 string                     `json:"state_fips"`
	Identity                  CanonicalEntityIdentity    `json:"identity"`
	ExternalIdentifiers       []ExternalEntityIdentifier `json:"external_identifiers,omitempty"`
	LegalStatus               LegalStatus                `json:"legal_status"`
	EffectiveInterval         EntityEffectiveInterval    `json:"effective_interval,omitempty"`
	Geographies               []EntityGeographyReference `json:"geographies,omitempty"`
	ParentEntityID            string                     `json:"parent_entity_id,omitempty"`
	SuccessorEntityIDs        []string                   `json:"successor_entity_ids,omitempty"`
	Classification           JurisdictionClassification `json:"classification"`
	HistoricalGeographyStatus HistoricalGeographyStatus  `json:"historical_geography_status"`
	SourceIDs                 []string                   `json:"source_ids"`
}

type GovernmentalEntityCandidateInput struct {
	OfficialName              string
	EntityType                GovernmentalEntityType
	StateID                   string
	StateFIPS                 string
	Identity                  CanonicalEntityIdentity
	ExternalIdentifiers       []ExternalEntityIdentifier
	LegalStatus               LegalStatus
	EffectiveInterval         EntityEffectiveInterval
	Geographies               []EntityGeographyReference
	ParentEntityID            string
	SuccessorEntityIDs        []string
	HistoricalGeographyStatus HistoricalGeographyStatus
	SourceIDs                 []string
}

type EntityInventory struct {
	SchemaVersion string                `json:"schema_version"`
	InventoryID   string                `json:"inventory_id"`
	GeneratedAt   string                `json:"generated_at"`
	Sources       []Source              `json:"sources"`
	Entities      []GovernmentalEntity  `json:"entities"`
}

var (
	stateIDPattern       = regexp.MustCompile(`^US-[A-Z]{2}$`)
	stateFIPSPattern     = regexp.MustCompile(`^[0-9]{2}$`)
	identityPartPattern  = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9._-]*$`)
)

func NewGovernmentalEntityCandidate(input GovernmentalEntityCandidateInput) (GovernmentalEntity, error) {
	identity := CanonicalEntityIdentity{
		Namespace: strings.ToLower(strings.TrimSpace(input.Identity.Namespace)),
		Value:     strings.TrimSpace(input.Identity.Value),
	}
	entityID, err := CanonicalGovernmentalEntityID(input.StateID, input.EntityType, identity)
	if err != nil {
		return GovernmentalEntity{}, err
	}
	historical := input.HistoricalGeographyStatus
	if historical == "" {
		historical = HistoricalGeographyCurrentOnly
	}
	entity := GovernmentalEntity{
		SchemaVersion:             GovernmentalEntitySchemaVersion,
		EntityID:                  entityID,
		OfficialName:              strings.TrimSpace(input.OfficialName),
		EntityType:                input.EntityType,
		StateID:                   strings.TrimSpace(input.StateID),
		StateFIPS:                 strings.TrimSpace(input.StateFIPS),
		Identity:                  identity,
		ExternalIdentifiers:       append([]ExternalEntityIdentifier(nil), input.ExternalIdentifiers...),
		LegalStatus:               input.LegalStatus,
		EffectiveInterval:         input.EffectiveInterval,
		Geographies:               append([]EntityGeographyReference(nil), input.Geographies...),
		ParentEntityID:            strings.TrimSpace(input.ParentEntityID),
		SuccessorEntityIDs:        append([]string(nil), input.SuccessorEntityIDs...),
		Classification:           ClassificationUnresolved,
		HistoricalGeographyStatus: historical,
		SourceIDs:                 append([]string(nil), input.SourceIDs...),
	}
	if err := validateGovernmentalEntity(entity); err != nil {
		return GovernmentalEntity{}, err
	}
	return entity, nil
}

func CanonicalGovernmentalEntityID(stateID string, entityType GovernmentalEntityType, identity CanonicalEntityIdentity) (string, error) {
	stateID = strings.TrimSpace(stateID)
	if !stateIDPattern.MatchString(stateID) {
		return "", fmt.Errorf("state_id must match US-XX")
	}
	if !validGovernmentalEntityType(entityType) {
		return "", fmt.Errorf("unsupported governmental entity type %q", entityType)
	}
	namespace := strings.ToLower(strings.TrimSpace(identity.Namespace))
	if namespace == "" {
		return "", fmt.Errorf("identity namespace is required")
	}
	if !identityPartPattern.MatchString(namespace) {
		return "", fmt.Errorf("identity namespace %q contains unsupported characters", namespace)
	}
	value := strings.TrimSpace(identity.Value)
	if value == "" {
		return "", fmt.Errorf("identity value is required")
	}
	if !identityPartPattern.MatchString(value) {
		return "", fmt.Errorf("identity value %q contains unsupported characters", value)
	}
	return fmt.Sprintf("gov:%s:%s:%s:%s", strings.ToLower(stateID), entityType, namespace, value), nil
}

func ValidateEntityInventory(inventory EntityInventory) error {
	if inventory.SchemaVersion != EntityInventorySchemaVersion {
		return fmt.Errorf("unsupported entity inventory schema_version %q", inventory.SchemaVersion)
	}
	if strings.TrimSpace(inventory.InventoryID) == "" {
		return fmt.Errorf("inventory_id is required")
	}
	if _, err := time.Parse(time.RFC3339, inventory.GeneratedAt); err != nil {
		return fmt.Errorf("generated_at must be RFC3339: %w", err)
	}
	if len(inventory.Sources) == 0 {
		return fmt.Errorf("entity inventory requires at least one source")
	}
	knownSources := make(map[string]struct{}, len(inventory.Sources))
	for index, source := range inventory.Sources {
		if err := validateInventorySource(source); err != nil {
			return fmt.Errorf("source %d: %w", index, err)
		}
		if _, duplicate := knownSources[source.ID]; duplicate {
			return fmt.Errorf("duplicate source id %q", source.ID)
		}
		knownSources[source.ID] = struct{}{}
	}
	if len(inventory.Entities) == 0 {
		return fmt.Errorf("entity inventory requires at least one entity")
	}
	seenEntities := make(map[string]struct{}, len(inventory.Entities))
	for index, entity := range inventory.Entities {
		if err := validateGovernmentalEntity(entity); err != nil {
			return fmt.Errorf("entity %d: %w", index, err)
		}
		if _, duplicate := seenEntities[entity.EntityID]; duplicate {
			return fmt.Errorf("duplicate entity_id %q", entity.EntityID)
		}
		seenEntities[entity.EntityID] = struct{}{}
		for _, sourceID := range entity.SourceIDs {
			if _, ok := knownSources[sourceID]; !ok {
				return fmt.Errorf("entity %q references unknown source %q", entity.EntityID, sourceID)
			}
		}
	}
	return nil
}

func validateGovernmentalEntity(entity GovernmentalEntity) error {
	if entity.SchemaVersion != GovernmentalEntitySchemaVersion {
		return fmt.Errorf("unsupported governmental entity schema_version %q", entity.SchemaVersion)
	}
	if strings.TrimSpace(entity.OfficialName) == "" {
		return fmt.Errorf("official_name is required")
	}
	if !stateIDPattern.MatchString(entity.StateID) {
		return fmt.Errorf("state_id must match US-XX")
	}
	if !stateFIPSPattern.MatchString(entity.StateFIPS) {
		return fmt.Errorf("state_fips must contain two digits")
	}
	canonicalID, err := CanonicalGovernmentalEntityID(entity.StateID, entity.EntityType, entity.Identity)
	if err != nil {
		return err
	}
	if entity.EntityID != canonicalID {
		return fmt.Errorf("entity_id %q does not match canonical identity %q", entity.EntityID, canonicalID)
	}
	if !validLegalStatus(entity.LegalStatus) {
		return fmt.Errorf("unsupported legal_status %q", entity.LegalStatus)
	}
	if !validClassification(entity.Classification) {
		return fmt.Errorf("classification must be explicit and supported, got %q", entity.Classification)
	}
	if !validHistoricalGeographyStatus(entity.HistoricalGeographyStatus) {
		return fmt.Errorf("unsupported historical_geography_status %q", entity.HistoricalGeographyStatus)
	}
	if len(entity.SourceIDs) == 0 {
		return fmt.Errorf("source_ids must contain evidence")
	}
	for _, sourceID := range entity.SourceIDs {
		if strings.TrimSpace(sourceID) == "" {
			return fmt.Errorf("source_ids must not contain empty values")
		}
	}
	if err := validateEffectiveInterval(entity.EffectiveInterval); err != nil {
		return err
	}
	for index, identifier := range entity.ExternalIdentifiers {
		if strings.TrimSpace(identifier.Scheme) == "" || strings.TrimSpace(identifier.Value) == "" {
			return fmt.Errorf("external identifier %d requires scheme and value", index)
		}
	}
	for index, geography := range entity.Geographies {
		if strings.TrimSpace(geography.Kind) == "" || strings.TrimSpace(geography.Identifier) == "" || strings.TrimSpace(geography.Vintage) == "" {
			return fmt.Errorf("geography %d requires kind, identifier, and vintage", index)
		}
	}
	return nil
}

func validateInventorySource(source Source) error {
	if strings.TrimSpace(source.ID) == "" || strings.TrimSpace(source.Title) == "" || strings.TrimSpace(source.URL) == "" || strings.TrimSpace(source.Kind) == "" {
		return fmt.Errorf("id, title, url, and kind are required")
	}
	if _, err := time.Parse(time.DateOnly, source.AccessedAt); err != nil {
		return fmt.Errorf("accessed_at must be YYYY-MM-DD: %w", err)
	}
	if source.LastCheckedAt != "" {
		if _, err := time.Parse(time.DateOnly, source.LastCheckedAt); err != nil {
			return fmt.Errorf("last_checked_at must be YYYY-MM-DD: %w", err)
		}
	}
	return nil
}

func validateEffectiveInterval(interval EntityEffectiveInterval) error {
	var start, end time.Time
	var err error
	if interval.StartDate != "" {
		start, err = time.Parse(time.DateOnly, interval.StartDate)
		if err != nil {
			return fmt.Errorf("effective_interval.start_date must be YYYY-MM-DD: %w", err)
		}
	}
	if interval.EndDate != "" {
		end, err = time.Parse(time.DateOnly, interval.EndDate)
		if err != nil {
			return fmt.Errorf("effective_interval.end_date must be YYYY-MM-DD: %w", err)
		}
	}
	if !start.IsZero() && !end.IsZero() && end.Before(start) {
		return fmt.Errorf("effective_interval end_date precedes start_date")
	}
	return nil
}

func validGovernmentalEntityType(value GovernmentalEntityType) bool {
	switch value {
	case EntityTypeState,
		EntityTypeCountyEquivalent,
		EntityTypeMunicipality,
		EntityTypeMinorCivilDivision,
		EntityTypeConsolidatedGovernment,
		EntityTypeIndependentCity,
		EntityTypeSpecialDistrict,
		EntityTypeTribalGovernment,
		EntityTypeOtherGovernmentalEntity:
		return true
	default:
		return false
	}
}

func validClassification(value JurisdictionClassification) bool {
	switch value {
	case ClassificationCoveredByState,
		ClassificationLocalAdopter,
		ClassificationLocalAmender,
		ClassificationEnforcementOnly,
		ClassificationDelegated,
		ClassificationNoGeneralAuthority,
		ClassificationInactive,
		ClassificationNotApplicable,
		ClassificationUnresolved:
		return true
	default:
		return false
	}
}

func validLegalStatus(value LegalStatus) bool {
	switch value {
	case LegalStatusActive, LegalStatusInactive, LegalStatusUnknown:
		return true
	default:
		return false
	}
}

func validHistoricalGeographyStatus(value HistoricalGeographyStatus) bool {
	switch value {
	case HistoricalGeographySupported, HistoricalGeographyCurrentOnly, HistoricalGeographyUnavailable:
		return true
	default:
		return false
	}
}
