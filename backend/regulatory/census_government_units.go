package regulatory

import (
	"fmt"
	"strconv"
	"strings"
)

const CensusGovernmentUnitsClassificationSystem = "census-government-units-2025"

type CensusGovernmentUnitRecord struct {
	CensusID                 string
	UnitName                 string
	UnitType                 string
	PoliticalCodeDescription string
	State                    string
	FIPSState                string
	FIPSCounty               string
	FIPSPlace                string
	Population               string
	PopulationSourceYear     string
	Active                   string
}

func GovernmentalEntityFromCensusUnit(record CensusGovernmentUnitRecord, sourceID string) (GovernmentalEntity, error) {
	sourceID = strings.TrimSpace(sourceID)
	if sourceID == "" {
		return GovernmentalEntity{}, fmt.Errorf("source id is required")
	}
	state := strings.ToUpper(strings.TrimSpace(record.State))
	if len(state) != 2 {
		return GovernmentalEntity{}, fmt.Errorf("state abbreviation must contain two letters")
	}
	entityType, err := censusGovernmentalEntityType(record.UnitType, record.PoliticalCodeDescription, record.FIPSCounty)
	if err != nil {
		return GovernmentalEntity{}, err
	}
	legalStatus, classification, err := censusActivity(strings.TrimSpace(record.Active))
	if err != nil {
		return GovernmentalEntity{}, err
	}
	population, err := censusPopulation(record.Population, record.PopulationSourceYear)
	if err != nil {
		return GovernmentalEntity{}, err
	}
	identifiers := censusExternalIdentifiers(record)
	entity, err := NewGovernmentalEntityCandidate(GovernmentalEntityCandidateInput{
		OfficialName: strings.TrimSpace(record.UnitName),
		EntityType:   entityType,
		StateID:      "US-" + state,
		StateFIPS:    strings.TrimSpace(record.FIPSState),
		Identity: CanonicalEntityIdentity{
			Namespace: "census-government-unit",
			Value:     strings.TrimSpace(record.CensusID),
		},
		ExternalIdentifiers:       identifiers,
		LegalStatus:               legalStatus,
		HistoricalGeographyStatus: HistoricalGeographyUnavailable,
		Population:                population,
		SourceClassifications: []EntitySourceClassification{
			{
				System:                   CensusGovernmentUnitsClassificationSystem,
				UnitType:                 strings.TrimSpace(record.UnitType),
				PoliticalCodeDescription: strings.TrimSpace(record.PoliticalCodeDescription),
			},
		},
		SourceIDs: []string{sourceID},
	})
	if err != nil {
		return GovernmentalEntity{}, err
	}
	entity.Classification = classification
	if err := validateGovernmentalEntity(entity); err != nil {
		return GovernmentalEntity{}, err
	}
	return entity, nil
}

func censusGovernmentalEntityType(unitType, politicalCode, fipsCounty string) (GovernmentalEntityType, error) {
	unitType = strings.TrimSpace(unitType)
	politicalCode = strings.ToUpper(strings.TrimSpace(politicalCode))
	if isConsolidatedPoliticalCode(politicalCode) {
		return EntityTypeConsolidatedGovernment, nil
	}
	switch unitType {
	case "1 - COUNTY":
		return EntityTypeCountyEquivalent, nil
	case "2 - MUNICIPAL":
		if politicalCode == "CITY" && isCountyEquivalentFIPS(fipsCounty) {
			return EntityTypeIndependentCity, nil
		}
		return EntityTypeMunicipality, nil
	case "3 - TOWNSHIP":
		return EntityTypeMinorCivilDivision, nil
	default:
		return "", fmt.Errorf("unsupported Census government unit type %q", unitType)
	}
}

func isCountyEquivalentFIPS(value string) bool {
	value = strings.TrimSpace(value)
	if len(value) != 3 {
		return false
	}
	code, err := strconv.Atoi(value)
	return err == nil && code >= 500
}

func isConsolidatedPoliticalCode(value string) bool {
	switch value {
	case "CITY AND BOROUGH",
		"CITY AND COUNTY",
		"CITY-PARISH",
		"CONSOLIDATED GOVERNMENT",
		"METRO GOVERNMENT",
		"METROPOLITAN GOVERNMENT",
		"UNIFIED GOVERNMENT",
		"URBAN COUNTY GOVERNMENT":
		return true
	default:
		return false
	}
}

func censusActivity(value string) (LegalStatus, JurisdictionClassification, error) {
	switch strings.ToUpper(strings.TrimSpace(value)) {
	case "Y":
		return LegalStatusActive, ClassificationUnresolved, nil
	case "N":
		return LegalStatusInactive, ClassificationInactive, nil
	default:
		return "", "", fmt.Errorf("unsupported Census ACTIVE flag %q", value)
	}
}

func censusPopulation(value, sourceYear string) (*EntityPopulation, error) {
	value = strings.TrimSpace(value)
	sourceYear = strings.TrimSpace(sourceYear)
	if value == "" {
		if sourceYear != "" {
			return nil, fmt.Errorf("population source year is present without population")
		}
		return nil, nil
	}
	count, err := strconv.ParseInt(value, 10, 64)
	if err != nil || count < 0 {
		return nil, fmt.Errorf("invalid population %q", value)
	}
	if sourceYear == "" {
		return nil, fmt.Errorf("population source year is required with population")
	}
	return &EntityPopulation{Count: count, SourceYear: sourceYear}, nil
}

func censusExternalIdentifiers(record CensusGovernmentUnitRecord) []ExternalEntityIdentifier {
	pairs := []ExternalEntityIdentifier{
		{Scheme: "census-government-unit-pid6", Value: strings.TrimSpace(record.CensusID)},
		{Scheme: "government-units-fips-state", Value: strings.TrimSpace(record.FIPSState)},
		{Scheme: "government-units-fips-county", Value: strings.TrimSpace(record.FIPSCounty)},
		{Scheme: "government-units-fips-place", Value: strings.TrimSpace(record.FIPSPlace)},
	}
	result := make([]ExternalEntityIdentifier, 0, len(pairs))
	for _, pair := range pairs {
		if pair.Value != "" {
			result = append(result, pair)
		}
	}
	return result
}
