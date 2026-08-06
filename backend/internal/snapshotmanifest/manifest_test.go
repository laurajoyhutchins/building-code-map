package snapshotmanifest

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestCanonicalJSONIsDeterministic(t *testing.T) {
	manifest := validManifest("abc", 3)
	manifest.Sources = append(manifest.Sources, SourceArtifact{
		Publisher: "A", Product: "First", Vintage: "2025", Locator: "a", SHA256: digestOf("a"), RetrievedAt: "2026-08-05T12:00:00Z", LicenseReviewStatus: "reviewed", RedistributionStatus: "restricted",
	})
	first, err := CanonicalJSON(manifest)
	if err != nil {
		t.Fatal(err)
	}
	second, err := CanonicalJSON(manifest)
	if err != nil {
		t.Fatal(err)
	}
	if string(first) != string(second) {
		t.Fatalf("canonical JSON changed between calls")
	}
}

func TestLoadAndVerifyAcceptsMatchingSnapshot(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "boundary.sqlite")
	if err := os.WriteFile(path, []byte("abc"), 0o600); err != nil {
		t.Fatal(err)
	}
	manifest := validManifest(digestOf("abc"), 3)
	writeManifestAndReceipt(t, path, manifest)

	verified, err := LoadAndVerify(path, KindBoundary)
	if err != nil {
		t.Fatal(err)
	}
	if verified.Manifest.SnapshotID != "boundary-test" {
		t.Fatalf("unexpected snapshot id %q", verified.Manifest.SnapshotID)
	}
}

func TestLoadAndVerifyRejectsChecksumMismatch(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "boundary.sqlite")
	if err := os.WriteFile(path, []byte("changed"), 0o600); err != nil {
		t.Fatal(err)
	}
	manifest := validManifest(digestOf("abc"), 3)
	writeManifestAndReceipt(t, path, manifest)

	_, err := LoadAndVerify(path, KindBoundary)
	if !errors.Is(err, ErrChecksumMismatch) {
		t.Fatalf("expected checksum mismatch, got %v", err)
	}
}

func TestLoadAndVerifyRejectsWrongKind(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "geocoder.sqlite")
	if err := os.WriteFile(path, []byte("abc"), 0o600); err != nil {
		t.Fatal(err)
	}
	manifest := validManifest(digestOf("abc"), 3)
	writeManifestAndReceipt(t, path, manifest)

	_, err := LoadAndVerify(path, KindGeocoder)
	if !errors.Is(err, ErrManifestInvalid) {
		t.Fatalf("expected invalid manifest, got %v", err)
	}
}

func TestLoadAndVerifyRejectsFailedIntegrityCheck(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "boundary.sqlite")
	if err := os.WriteFile(path, []byte("abc"), 0o600); err != nil {
		t.Fatal(err)
	}
	manifest := validManifest(digestOf("abc"), 3)
	manifest.IntegrityChecks[0].Status = "failed"
	writeManifestAndReceipt(t, path, manifest)

	_, err := LoadAndVerify(path, KindBoundary)
	if !errors.Is(err, ErrManifestInvalid) {
		t.Fatalf("expected invalid manifest, got %v", err)
	}
}

func validManifest(outputDigest string, outputSize int64) Manifest {
	return Manifest{
		SchemaVersion: SchemaVersion,
		SnapshotID:    "boundary-test",
		Kind:          KindBoundary,
		Sources: []SourceArtifact{{
			Publisher: "Census", Product: "TIGER/Line", Vintage: "2025", Locator: "source.zip", SHA256: digestOf("source"), RetrievedAt: "2026-08-05T12:00:00Z", LicenseReviewStatus: "reviewed", RedistributionStatus: "allowed",
		}},
		Builder: Builder{Tool: "boundary-build", Version: "1.0.0", Revision: "0123456789abcdef0123456789abcdef01234567", BuiltAt: "2026-08-05T12:30:00Z", OutputCRS: "EPSG:4326"},
		RecordCounts: RecordCounts{Accepted: 1},
		OutputSHA256: outputDigest,
		OutputSizeBytes: outputSize,
		IntegrityChecks: []IntegrityCheck{{Name: "sqlite-integrity", Status: "passed"}},
	}
}

func writeManifestAndReceipt(t *testing.T, path string, manifest Manifest) {
	t.Helper()
	manifestBytes, err := CanonicalJSON(manifest)
	if err != nil {
		t.Fatal(err)
	}
	manifestBytes = append(manifestBytes, '\n')
	if err := os.WriteFile(ManifestPath(path), manifestBytes, 0o600); err != nil {
		t.Fatal(err)
	}
	manifestHash := sha256.Sum256(manifestBytes)
	receipt := ActivationReceipt{
		SchemaVersion: SchemaVersion,
		SnapshotID: manifest.SnapshotID,
		ActivatedAt: "2026-08-05T12:45:00Z",
		LastKnownGoodSnapshotID: manifest.SnapshotID,
		ManifestSHA256: hex.EncodeToString(manifestHash[:]),
	}
	receiptBytes, err := json.Marshal(receipt)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(ActivationPath(path), receiptBytes, 0o600); err != nil {
		t.Fatal(err)
	}
}

func digestOf(value string) string {
	digest := sha256.Sum256([]byte(value))
	return hex.EncodeToString(digest[:])
}
