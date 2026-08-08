package regulatory

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
)

func WriteCensusEntityInventoryDirectory(root string, build CensusEntityInventoryBuild) error {
	root = filepath.Clean(root)
	parent := filepath.Dir(root)
	if err := os.MkdirAll(parent, 0o755); err != nil {
		return fmt.Errorf("create inventory parent: %w", err)
	}
	temporary, err := os.MkdirTemp(parent, ".governmental-entities-")
	if err != nil {
		return fmt.Errorf("create temporary inventory directory: %w", err)
	}
	keepTemporary := false
	defer func() {
		if !keepTemporary {
			_ = os.RemoveAll(temporary)
		}
	}()

	stateIDs := make([]string, 0, len(build.Inventories))
	for stateID := range build.Inventories {
		stateIDs = append(stateIDs, stateID)
	}
	sort.Strings(stateIDs)
	for _, stateID := range stateIDs {
		inventory := build.Inventories[stateID]
		if err := ValidateEntityInventory(inventory); err != nil {
			return fmt.Errorf("validate %s before write: %w", stateID, err)
		}
		if err := writeCanonicalJSON(filepath.Join(temporary, stateID+".json"), inventory); err != nil {
			return err
		}
	}
	if err := validateEntityInventoryIndex(build.Index); err != nil {
		return err
	}
	if err := writeCanonicalJSON(filepath.Join(temporary, "index.json"), build.Index); err != nil {
		return err
	}

	backup := root + ".previous"
	_ = os.RemoveAll(backup)
	if _, statErr := os.Stat(root); statErr == nil {
		if err := os.Rename(root, backup); err != nil {
			return fmt.Errorf("preserve prior inventory directory: %w", err)
		}
	} else if !os.IsNotExist(statErr) {
		return fmt.Errorf("inspect inventory directory: %w", statErr)
	}
	if err := os.Rename(temporary, root); err != nil {
		if _, backupErr := os.Stat(backup); backupErr == nil {
			_ = os.Rename(backup, root)
		}
		return fmt.Errorf("publish inventory directory: %w", err)
	}
	keepTemporary = true
	_ = os.RemoveAll(backup)
	return nil
}

func LoadCensusEntityInventoryDirectory(root string) (CensusEntityInventoryBuild, error) {
	root = filepath.Clean(root)
	var index EntityInventoryIndex
	if err := decodeStrictJSON(filepath.Join(root, "index.json"), &index); err != nil {
		return CensusEntityInventoryBuild{}, err
	}
	if err := validateEntityInventoryIndex(index); err != nil {
		return CensusEntityInventoryBuild{}, err
	}

	inventories := make(map[string]EntityInventory, len(index.States))
	for _, summary := range index.States {
		var inventory EntityInventory
		if err := decodeStrictJSON(filepath.Join(root, summary.StateID+".json"), &inventory); err != nil {
			return CensusEntityInventoryBuild{}, err
		}
		if err := ValidateEntityInventory(inventory); err != nil {
			return CensusEntityInventoryBuild{}, fmt.Errorf("validate %s inventory: %w", summary.StateID, err)
		}
		if _, duplicate := inventories[summary.StateID]; duplicate {
			return CensusEntityInventoryBuild{}, fmt.Errorf("duplicate state summary %q", summary.StateID)
		}
		computed := summarizeEntityInventory(summary.StateID, inventory.Entities)
		if !reflect.DeepEqual(computed, summary) {
			return CensusEntityInventoryBuild{}, fmt.Errorf("state summary drift for %s", summary.StateID)
		}
		inventories[summary.StateID] = inventory
	}
	computedIndex := indexFromInventories(inventories, index.SourceID, index.GeneratedAt)
	if !reflect.DeepEqual(computedIndex, index) {
		return CensusEntityInventoryBuild{}, fmt.Errorf("inventory index totals drift from state inventories")
	}
	return CensusEntityInventoryBuild{Inventories: inventories, Index: index}, nil
}

func validateEntityInventoryIndex(index EntityInventoryIndex) error {
	if index.SchemaVersion != EntityInventorySchemaVersion {
		return fmt.Errorf("unsupported entity inventory index schema_version %q", index.SchemaVersion)
	}
	if index.SourceID == "" || index.GeneratedAt == "" {
		return fmt.Errorf("entity inventory index source_id and generated_at are required")
	}
	if len(index.States) == 0 {
		return fmt.Errorf("entity inventory index requires state summaries")
	}
	if index.TotalEntities < 1 {
		return fmt.Errorf("entity inventory index total_entities must be positive")
	}
	return nil
}

func indexFromInventories(inventories map[string]EntityInventory, sourceID, generatedAt string) EntityInventoryIndex {
	stateIDs := make([]string, 0, len(inventories))
	for stateID := range inventories {
		stateIDs = append(stateIDs, stateID)
	}
	sort.Strings(stateIDs)
	index := EntityInventoryIndex{
		SchemaVersion:        EntityInventorySchemaVersion,
		SourceID:             sourceID,
		GeneratedAt:          generatedAt,
		EntityTypeCounts:     map[GovernmentalEntityType]int{},
		ClassificationCounts: map[JurisdictionClassification]int{},
		States:               make([]EntityInventoryStateSummary, 0, len(stateIDs)),
	}
	for _, stateID := range stateIDs {
		summary := summarizeEntityInventory(stateID, inventories[stateID].Entities)
		index.States = append(index.States, summary)
		index.TotalEntities += summary.TotalEntities
		for entityType, count := range summary.EntityTypeCounts {
			index.EntityTypeCounts[entityType] += count
		}
		for classification, count := range summary.ClassificationCounts {
			index.ClassificationCounts[classification] += count
		}
	}
	return index
}

func summarizeEntityInventory(stateID string, entities []GovernmentalEntity) EntityInventoryStateSummary {
	summary := EntityInventoryStateSummary{
		StateID:              stateID,
		TotalEntities:        len(entities),
		EntityTypeCounts:     map[GovernmentalEntityType]int{},
		ClassificationCounts: map[JurisdictionClassification]int{},
	}
	for _, entity := range entities {
		summary.EntityTypeCounts[entity.EntityType]++
		summary.ClassificationCounts[entity.Classification]++
	}
	return summary
}

func writeCanonicalJSON(path string, value any) error {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return fmt.Errorf("encode %s: %w", path, err)
	}
	data = append(data, '\n')
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", path, err)
	}
	return nil
}

func decodeStrictJSON(path string, target any) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open %s: %w", path, err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(target); err != nil {
		return fmt.Errorf("decode %s: %w", path, err)
	}
	return nil
}
