package snapshotmanifest

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestActivateReplacesSnapshotAndRecordsPriorIdentity(t *testing.T) {
	dir := t.TempDir()
	activePath := filepath.Join(dir, "active.sqlite")
	candidatePath := filepath.Join(dir, "candidate.sqlite")

	writeSnapshotFixture(t, activePath, "old", "old-snapshot")
	writeSnapshotFixture(t, candidatePath, "new", "new-snapshot")

	result, err := Activate(candidatePath, activePath, KindBoundary, time.Date(2026, 8, 5, 13, 0, 0, 0, time.UTC))
	if err != nil {
		t.Fatal(err)
	}
	if result.Receipt.PriorActiveSnapshotID != "old-snapshot" {
		t.Fatalf("prior id=%q", result.Receipt.PriorActiveSnapshotID)
	}
	verified, err := LoadAndVerify(activePath, KindBoundary)
	if err != nil {
		t.Fatal(err)
	}
	if verified.Manifest.SnapshotID != "new-snapshot" {
		t.Fatalf("active id=%q", verified.Manifest.SnapshotID)
	}
	contents, err := os.ReadFile(activePath)
	if err != nil {
		t.Fatal(err)
	}
	if string(contents) != "new" {
		t.Fatalf("active contents=%q", contents)
	}
}

func TestActivateAcceptsBuiltCandidateWithoutActivationReceipt(t *testing.T) {
	dir := t.TempDir()
	activePath := filepath.Join(dir, "active.sqlite")
	candidatePath := filepath.Join(dir, "candidate.sqlite")

	writeSnapshotFixture(t, activePath, "old", "old-snapshot")
	if err := os.WriteFile(candidatePath, []byte("new"), 0o600); err != nil {
		t.Fatal(err)
	}
	manifest := validManifest(digestOf("new"), int64(len("new")))
	manifest.SnapshotID = "new-snapshot"
	if _, err := FinalizeAndWrite(candidatePath, manifest); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(ActivationPath(candidatePath)); !os.IsNotExist(err) {
		t.Fatalf("candidate unexpectedly has activation receipt: %v", err)
	}

	result, err := Activate(candidatePath, activePath, KindBoundary, time.Date(2026, 8, 8, 15, 0, 0, 0, time.UTC))
	if err != nil {
		t.Fatal(err)
	}
	if result.Receipt.SnapshotID != "new-snapshot" {
		t.Fatalf("activated snapshot id=%q", result.Receipt.SnapshotID)
	}
	verified, err := LoadAndVerify(activePath, KindBoundary)
	if err != nil {
		t.Fatal(err)
	}
	if verified.Manifest.SnapshotID != "new-snapshot" {
		t.Fatalf("active id=%q", verified.Manifest.SnapshotID)
	}
}

func TestActivateLeavesActiveSnapshotUntouchedWhenCandidateFailsVerification(t *testing.T) {
	dir := t.TempDir()
	activePath := filepath.Join(dir, "active.sqlite")
	candidatePath := filepath.Join(dir, "candidate.sqlite")

	writeSnapshotFixture(t, activePath, "old", "old-snapshot")
	writeSnapshotFixture(t, candidatePath, "new", "new-snapshot")
	if err := os.WriteFile(candidatePath, []byte("corrupt"), 0o600); err != nil {
		t.Fatal(err)
	}

	if _, err := Activate(candidatePath, activePath, KindBoundary, time.Now()); err == nil {
		t.Fatal("expected activation failure")
	}
	verified, err := LoadAndVerify(activePath, KindBoundary)
	if err != nil {
		t.Fatal(err)
	}
	if verified.Manifest.SnapshotID != "old-snapshot" {
		t.Fatalf("active id=%q", verified.Manifest.SnapshotID)
	}
}

func TestActivateRollsBackWholeGenerationWhenSidecarReplacementFails(t *testing.T) {
	dir := t.TempDir()
	activePath := filepath.Join(dir, "active.sqlite")
	candidatePath := filepath.Join(dir, "candidate.sqlite")

	writeSnapshotFixture(t, activePath, "old", "old-snapshot")
	writeSnapshotFixture(t, candidatePath, "new", "new-snapshot")

	manifestRollbackBlocker := ManifestPath(activePath) + ".rollback"
	if err := os.MkdirAll(manifestRollbackBlocker, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(manifestRollbackBlocker, "keep"), []byte("block"), 0o600); err != nil {
		t.Fatal(err)
	}

	if _, err := Activate(candidatePath, activePath, KindBoundary, time.Now()); err == nil {
		t.Fatal("expected activation failure")
	}

	verified, err := LoadAndVerify(activePath, KindBoundary)
	if err != nil {
		t.Fatalf("active generation must remain valid after failed activation: %v", err)
	}
	if verified.Manifest.SnapshotID != "old-snapshot" {
		t.Fatalf("active id=%q", verified.Manifest.SnapshotID)
	}
	contents, err := os.ReadFile(activePath)
	if err != nil {
		t.Fatal(err)
	}
	if string(contents) != "old" {
		t.Fatalf("active contents=%q", contents)
	}
}

func writeSnapshotFixture(t *testing.T, path, contents, snapshotID string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(contents), 0o600); err != nil {
		t.Fatal(err)
	}
	manifest := validManifest(digestOf(contents), int64(len(contents)))
	manifest.SnapshotID = snapshotID
	writeManifestAndReceipt(t, path, manifest)
}
