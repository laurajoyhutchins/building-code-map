package regulatory

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
)

func WriteEntityCoverageLedger(path string, ledger EntityCoverageLedger) error {
	if err := ValidateEntityCoverageLedger(ledger); err != nil {
		return err
	}
	raw, err := marshalEntityCoverageLedger(ledger)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, raw, 0o644)
}

func LoadEntityCoverageLedger(path string) (EntityCoverageLedger, error) {
	var ledger EntityCoverageLedger
	if err := decodeJSONFile(path, &ledger); err != nil {
		return EntityCoverageLedger{}, err
	}
	if err := ValidateEntityCoverageLedger(ledger); err != nil {
		return EntityCoverageLedger{}, fmt.Errorf("validate %s: %w", path, err)
	}
	return ledger, nil
}

func ValidateEntityCoverageLedger(ledger EntityCoverageLedger) error {
	if ledger.SchemaVersion != EntityCoverageLedgerSchemaVersion {
		return fmt.Errorf("unsupported coverage ledger schema version %q", ledger.SchemaVersion)
	}
	if strings.TrimSpace(ledger.InventorySourceID) == "" || strings.TrimSpace(ledger.InventoryGeneratedAt) == "" {
		return fmt.Errorf("inventory source identity is required")
	}
	if strings.TrimSpace(ledger.CodeFamily) == "" {
		return fmt.Errorf("code_family is required")
	}
	if strings.TrimSpace(ledger.ApplicabilityDate) == "" {
		return fmt.Errorf("applicability_date is required")
	}
	if len(ledger.StateIDs) == 0 || len(ledger.States) != len(ledger.StateIDs) {
		return fmt.Errorf("state_ids and states must contain the same non-zero state set")
	}
	if !sort.StringsAreSorted(ledger.StateIDs) {
		return fmt.Errorf("state_ids must be sorted")
	}
	seen := map[string]bool{}
	recomputed := newEntityCoverageSummary()
	for index, state := range ledger.States {
		if state.StateID != ledger.StateIDs[index] {
			return fmt.Errorf("state %d is %s, want %s", index, state.StateID, ledger.StateIDs[index])
		}
		if seen[state.StateID] {
			return fmt.Errorf("duplicate state %s", state.StateID)
		}
		seen[state.StateID] = true
		if strings.TrimSpace(state.RegimeID) == "" || strings.TrimSpace(state.RegimeLastVerified) == "" {
			return fmt.Errorf("state %s regime identity is required", state.StateID)
		}
		stateSummary := newEntityCoverageSummary()
		lastEntityID := ""
		for _, entity := range state.Entities {
			if entity.EntityID == "" || entity.OfficialName == "" {
				return fmt.Errorf("state %s contains entity without identity", state.StateID)
			}
			if lastEntityID != "" && entity.EntityID <= lastEntityID {
				return fmt.Errorf("state %s entities must be sorted and unique by entity_id", state.StateID)
			}
			lastEntityID = entity.EntityID
			if entity.EvidenceStatus == "" {
				return fmt.Errorf("entity %s evidence_status is required", entity.EntityID)
			}
			if len(entity.SourceIDs) == 0 {
				return fmt.Errorf("entity %s source_ids are required", entity.EntityID)
			}
			addCoverageEntry(&stateSummary, entity)
		}
		if !reflect.DeepEqual(stateSummary, state.Summary) {
			return fmt.Errorf("state %s summary does not match entity rows", state.StateID)
		}
		mergeCoverageSummary(&recomputed, state.Summary)
	}
	if !reflect.DeepEqual(recomputed, ledger.Summary) {
		return fmt.Errorf("national summary does not match state summaries")
	}
	if ledger.Freshness.Status == "" || ledger.Freshness.Reason == "" {
		return fmt.Errorf("freshness availability must be explicit")
	}
	if ledger.LandArea.Status == "" || ledger.LandArea.Reason == "" {
		return fmt.Errorf("land-area availability must be explicit")
	}
	return nil
}

func ValidateEntityCoverageLedgerFile(request EntityCoverageLedgerRequest, path string) error {
	if _, err := LoadEntityCoverageLedger(path); err != nil {
		return err
	}
	rebuilt, err := BuildEntityCoverageLedger(request)
	if err != nil {
		return err
	}
	want, err := marshalEntityCoverageLedger(rebuilt)
	if err != nil {
		return err
	}
	got, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	if !bytes.Equal(got, want) {
		return fmt.Errorf("coverage ledger %s is stale; regenerate it", path)
	}
	return nil
}

func marshalEntityCoverageLedger(ledger EntityCoverageLedger) ([]byte, error) {
	raw, err := json.MarshalIndent(ledger, "", "  ")
	if err != nil {
		return nil, err
	}
	return append(raw, '\n'), nil
}
