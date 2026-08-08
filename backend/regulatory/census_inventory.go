package regulatory

import (
	"encoding/csv"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strings"
	"time"
)

var CensusGeneralPurposeHeader = []string{
	"CENSUS_ID_PID6",
	"UNIT_NAME",
	"UNIT_TYPE",
	"TITLE",
	"ADDRESS1",
	"ADDRESS2",
	"CITY",
	"STATE",
	"ZIP",
	"ZIP4",
	"WEB_ADDRESS",
	"POLITICAL_CODE_DESCRIPTION",
	"POPULATION",
	"POPULATION_SOURCE_YEAR",
	"FIPS_STATE",
	"FIPS_COUNTY",
	"FIPS_PLACE",
	"COUNTY_AREA_NAME",
	"ACTIVE",
}

type EntityInventoryStateSummary struct {
	StateID              string                                 `json:"state_id"`
	TotalEntities        int                                    `json:"total_entities"`
	EntityTypeCounts     map[GovernmentalEntityType]int         `json:"entity_type_counts"`
	ClassificationCounts map[JurisdictionClassification]int     `json:"classification_counts"`
}

type EntityInventoryIndex struct {
	SchemaVersion         string                                 `json:"schema_version"`
	SourceID              string                                 `json:"source_id"`
	GeneratedAt           string                                 `json:"generated_at"`
	TotalEntities         int                                    `json:"total_entities"`
	EntityTypeCounts      map[GovernmentalEntityType]int         `json:"entity_type_counts"`
	ClassificationCounts  map[JurisdictionClassification]int     `json:"classification_counts"`
	States                []EntityInventoryStateSummary           `json:"states"`
}

type CensusEntityInventoryBuild struct {
	Inventories map[string]EntityInventory
	Index       EntityInventoryIndex
}

func BuildCensusEntityInventories(reader io.Reader, source Source, generatedAt time.Time) (CensusEntityInventoryBuild, error) {
	if err := validateInventorySource(source); err != nil {
		return CensusEntityInventoryBuild{}, fmt.Errorf("source: %w", err)
	}
	if generatedAt.IsZero() {
		return CensusEntityInventoryBuild{}, fmt.Errorf("generated time is required")
	}
	csvReader := csv.NewReader(reader)
	csvReader.FieldsPerRecord = len(CensusGeneralPurposeHeader)
	header, err := csvReader.Read()
	if err != nil {
		return CensusEntityInventoryBuild{}, fmt.Errorf("read Census header: %w", err)
	}
	if !reflect.DeepEqual(header, CensusGeneralPurposeHeader) {
		return CensusEntityInventoryBuild{}, fmt.Errorf("Census General Purpose header does not match source contract")
	}

	entitiesByState := map[string][]GovernmentalEntity{}
	for rowNumber := 2; ; rowNumber++ {
		row, readErr := csvReader.Read()
		if readErr == io.EOF {
			break
		}
		if readErr != nil {
			return CensusEntityInventoryBuild{}, fmt.Errorf("read Census row %d: %w", rowNumber, readErr)
		}
		record := CensusGovernmentUnitRecord{
			CensusID:                 row[0],
			UnitName:                 row[1],
			UnitType:                 row[2],
			State:                    row[7],
			PoliticalCodeDescription: row[11],
			Population:               row[12],
			PopulationSourceYear:     row[13],
			FIPSState:                row[14],
			FIPSCounty:               row[15],
			FIPSPlace:                row[16],
			Active:                   row[18],
		}
		entity, mapErr := GovernmentalEntityFromCensusUnit(record, source.ID)
		if mapErr != nil {
			return CensusEntityInventoryBuild{}, fmt.Errorf("map Census row %d PID6 %q: %w", rowNumber, record.CensusID, mapErr)
		}
		entitiesByState[entity.StateID] = append(entitiesByState[entity.StateID], entity)
	}
	if len(entitiesByState) == 0 {
		return CensusEntityInventoryBuild{}, fmt.Errorf("Census input contains no government units")
	}

	generated := generatedAt.UTC().Format(time.RFC3339)
	inventories := make(map[string]EntityInventory, len(entitiesByState))
	stateIDs := make([]string, 0, len(entitiesByState))
	for stateID := range entitiesByState {
		stateIDs = append(stateIDs, stateID)
	}
	sort.Strings(stateIDs)

	index := EntityInventoryIndex{
		SchemaVersion:        EntityInventorySchemaVersion,
		SourceID:             source.ID,
		GeneratedAt:          generated,
		EntityTypeCounts:     map[GovernmentalEntityType]int{},
		ClassificationCounts: map[JurisdictionClassification]int{},
		States:               make([]EntityInventoryStateSummary, 0, len(stateIDs)),
	}
	for _, stateID := range stateIDs {
		entities := entitiesByState[stateID]
		sort.Slice(entities, func(i, j int) bool {
			return entities[i].EntityID < entities[j].EntityID
		})
		inventory := EntityInventory{
			SchemaVersion: EntityInventorySchemaVersion,
			InventoryID:   CensusGovernmentUnitsClassificationSystem + ":" + strings.ToLower(stateID),
			GeneratedAt:   generated,
			Sources:       []Source{source},
			Entities:      entities,
		}
		if err := ValidateEntityInventory(inventory); err != nil {
			return CensusEntityInventoryBuild{}, fmt.Errorf("validate %s inventory: %w", stateID, err)
		}
		inventories[stateID] = inventory

		stateSummary := EntityInventoryStateSummary{
			StateID:              stateID,
			TotalEntities:        len(entities),
			EntityTypeCounts:     map[GovernmentalEntityType]int{},
			ClassificationCounts: map[JurisdictionClassification]int{},
		}
		for _, entity := range entities {
			stateSummary.EntityTypeCounts[entity.EntityType]++
			stateSummary.ClassificationCounts[entity.Classification]++
			index.EntityTypeCounts[entity.EntityType]++
			index.ClassificationCounts[entity.Classification]++
			index.TotalEntities++
		}
		index.States = append(index.States, stateSummary)
	}
	return CensusEntityInventoryBuild{Inventories: inventories, Index: index}, nil
}

func ValidateNationalCensusInventory(result CensusEntityInventoryBuild, expectedTotal int) error {
	if len(result.Inventories) != len(nationalStateIDs) {
		return fmt.Errorf("national Census inventory must contain 51 state/DC inventories, got %d", len(result.Inventories))
	}
	for _, stateID := range nationalStateIDs {
		inventory, ok := result.Inventories[stateID]
		if !ok {
			return fmt.Errorf("national Census inventory is missing %s", stateID)
		}
		if err := ValidateEntityInventory(inventory); err != nil {
			return fmt.Errorf("validate national %s inventory: %w", stateID, err)
		}
	}
	if result.Index.TotalEntities != expectedTotal {
		return fmt.Errorf("national Census denominator is %d, expected %d", result.Index.TotalEntities, expectedTotal)
	}
	if len(result.Index.States) != len(nationalStateIDs) {
		return fmt.Errorf("national Census index must contain 51 state/DC summaries, got %d", len(result.Index.States))
	}
	return nil
}

var nationalStateIDs = []string{
	"US-AK", "US-AL", "US-AR", "US-AZ", "US-CA", "US-CO", "US-CT", "US-DC", "US-DE", "US-FL",
	"US-GA", "US-HI", "US-IA", "US-ID", "US-IL", "US-IN", "US-KS", "US-KY", "US-LA", "US-MA",
	"US-MD", "US-ME", "US-MI", "US-MN", "US-MO", "US-MS", "US-MT", "US-NC", "US-ND", "US-NE",
	"US-NH", "US-NJ", "US-NM", "US-NV", "US-NY", "US-OH", "US-OK", "US-OR", "US-PA", "US-RI",
	"US-SC", "US-SD", "US-TN", "US-TX", "US-UT", "US-VA", "US-VT", "US-WA", "US-WI", "US-WV",
	"US-WY",
}
