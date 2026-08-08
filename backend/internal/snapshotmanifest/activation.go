package snapshotmanifest

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type ActivationResult struct {
	SnapshotPath string
	Receipt      ActivationReceipt
}

type fileReplacement struct {
	staged      string
	destination string
	backup      string
	backedUp    bool
	installed   bool
}

func Activate(candidatePath, activePath string, expectedKind Kind, activatedAt time.Time) (ActivationResult, error) {
	verified, err := LoadAndVerify(candidatePath, expectedKind)
	if err != nil {
		return ActivationResult{}, err
	}

	activeDir := filepath.Dir(activePath)
	if err := os.MkdirAll(activeDir, 0o755); err != nil {
		return ActivationResult{}, fmt.Errorf("create active snapshot directory: %w", err)
	}

	priorID := ""
	if prior, err := LoadAndVerify(activePath, expectedKind); err == nil {
		priorID = prior.Manifest.SnapshotID
	}

	stagedSnapshot := activePath + ".staging"
	stagedManifest := ManifestPath(stagedSnapshot)
	stagedReceipt := ActivationPath(stagedSnapshot)
	cleanup := func() {
		_ = os.Remove(stagedSnapshot)
		_ = os.Remove(stagedManifest)
		_ = os.Remove(stagedReceipt)
	}
	defer cleanup()

	if err := copyFile(candidatePath, stagedSnapshot); err != nil {
		return ActivationResult{}, err
	}
	manifestBytes, err := os.ReadFile(ManifestPath(candidatePath))
	if err != nil {
		return ActivationResult{}, fmt.Errorf("read candidate manifest: %w", err)
	}
	if err := os.WriteFile(stagedManifest, manifestBytes, 0o600); err != nil {
		return ActivationResult{}, fmt.Errorf("stage manifest: %w", err)
	}
	manifestDigest := sha256.Sum256(manifestBytes)
	receipt := ActivationReceipt{
		SchemaVersion:           SchemaVersion,
		SnapshotID:              verified.Manifest.SnapshotID,
		ActivatedAt:             activatedAt.UTC().Format(time.RFC3339),
		PriorActiveSnapshotID:   priorID,
		LastKnownGoodSnapshotID: verified.Manifest.SnapshotID,
		ManifestSHA256:          hex.EncodeToString(manifestDigest[:]),
	}
	receiptBytes, err := json.MarshalIndent(receipt, "", "  ")
	if err != nil {
		return ActivationResult{}, fmt.Errorf("encode activation receipt: %w", err)
	}
	receiptBytes = append(receiptBytes, '\n')
	if err := os.WriteFile(stagedReceipt, receiptBytes, 0o600); err != nil {
		return ActivationResult{}, fmt.Errorf("stage activation receipt: %w", err)
	}

	if _, err := LoadAndVerify(stagedSnapshot, expectedKind); err != nil {
		return ActivationResult{}, fmt.Errorf("verify staged snapshot: %w", err)
	}

	replacements := []fileReplacement{
		{staged: stagedSnapshot, destination: activePath, backup: activePath + ".rollback"},
		{staged: stagedManifest, destination: ManifestPath(activePath), backup: ManifestPath(activePath) + ".rollback"},
		{staged: stagedReceipt, destination: ActivationPath(activePath), backup: ActivationPath(activePath) + ".rollback"},
	}
	if err := replaceGeneration(replacements); err != nil {
		return ActivationResult{}, err
	}
	return ActivationResult{SnapshotPath: activePath, Receipt: receipt}, nil
}

func replaceGeneration(replacements []fileReplacement) error {
	for i := range replacements {
		if _, err := os.Stat(replacements[i].backup); err == nil {
			if err := os.Remove(replacements[i].backup); err != nil {
				return fmt.Errorf("clear prior rollback file %q: %w", replacements[i].backup, err)
			}
		} else if !os.IsNotExist(err) {
			return fmt.Errorf("inspect rollback file %q: %w", replacements[i].backup, err)
		}
	}

	for i := range replacements {
		if _, err := os.Stat(replacements[i].destination); err == nil {
			if err := os.Rename(replacements[i].destination, replacements[i].backup); err != nil {
				rollbackErr := rollbackGeneration(replacements)
				if rollbackErr != nil {
					return errors.Join(fmt.Errorf("preserve prior active file %q: %w", replacements[i].destination, err), rollbackErr)
				}
				return fmt.Errorf("preserve prior active file %q: %w", replacements[i].destination, err)
			}
			replacements[i].backedUp = true
		} else if !os.IsNotExist(err) {
			rollbackErr := rollbackGeneration(replacements)
			if rollbackErr != nil {
				return errors.Join(fmt.Errorf("inspect active file %q: %w", replacements[i].destination, err), rollbackErr)
			}
			return fmt.Errorf("inspect active file %q: %w", replacements[i].destination, err)
		}
	}

	for i := range replacements {
		if err := os.Rename(replacements[i].staged, replacements[i].destination); err != nil {
			rollbackErr := rollbackGeneration(replacements)
			if rollbackErr != nil {
				return errors.Join(fmt.Errorf("activate staged file %q: %w", replacements[i].destination, err), rollbackErr)
			}
			return fmt.Errorf("activate staged file %q: %w", replacements[i].destination, err)
		}
		replacements[i].installed = true
	}
	return nil
}

func rollbackGeneration(replacements []fileReplacement) error {
	var rollbackErrors []error
	for i := len(replacements) - 1; i >= 0; i-- {
		if replacements[i].installed {
			if err := os.Remove(replacements[i].destination); err != nil && !os.IsNotExist(err) {
				rollbackErrors = append(rollbackErrors, fmt.Errorf("remove partially activated file %q: %w", replacements[i].destination, err))
			}
		}
		if replacements[i].backedUp {
			if err := os.Rename(replacements[i].backup, replacements[i].destination); err != nil {
				rollbackErrors = append(rollbackErrors, fmt.Errorf("restore prior active file %q: %w", replacements[i].destination, err))
			}
		}
	}
	return errors.Join(rollbackErrors...)
}

func copyFile(source, destination string) error {
	data, err := os.ReadFile(source)
	if err != nil {
		return fmt.Errorf("read candidate snapshot: %w", err)
	}
	if err := os.WriteFile(destination, data, 0o600); err != nil {
		return fmt.Errorf("write staged snapshot: %w", err)
	}
	return nil
}
