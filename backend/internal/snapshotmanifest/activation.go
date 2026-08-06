package snapshotmanifest

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type ActivationResult struct {
	SnapshotPath string
	Receipt      ActivationReceipt
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

	if err := replaceFile(stagedSnapshot, activePath); err != nil {
		return ActivationResult{}, err
	}
	if err := replaceFile(stagedManifest, ManifestPath(activePath)); err != nil {
		return ActivationResult{}, err
	}
	if err := replaceFile(stagedReceipt, ActivationPath(activePath)); err != nil {
		return ActivationResult{}, err
	}
	return ActivationResult{SnapshotPath: activePath, Receipt: receipt}, nil
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

func replaceFile(source, destination string) error {
	backup := destination + ".rollback"
	if _, err := os.Stat(destination); err == nil {
		_ = os.Remove(backup)
		if err := os.Rename(destination, backup); err != nil {
			return fmt.Errorf("preserve prior active file: %w", err)
		}
	}
	if err := os.Rename(source, destination); err != nil {
		if _, backupErr := os.Stat(backup); backupErr == nil {
			_ = os.Rename(backup, destination)
		}
		return fmt.Errorf("activate staged file: %w", err)
	}
	return nil
}
