package snapshotmanifest

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const SchemaVersion = "1.0"

type Kind string

const (
	KindBoundary Kind = "boundary"
	KindGeocoder Kind = "geocoder"
)

var (
	ErrManifestMissing  = errors.New("snapshot manifest missing")
	ErrManifestInvalid  = errors.New("snapshot manifest invalid")
	ErrChecksumMismatch = errors.New("snapshot checksum mismatch")
)

type SourceArtifact struct {
	Publisher            string `json:"publisher"`
	Product              string `json:"product"`
	Layer                string `json:"layer,omitempty"`
	Vintage              string `json:"vintage"`
	Locator              string `json:"locator"`
	SHA256               string `json:"sha256"`
	RetrievedAt          string `json:"retrieved_at"`
	LicenseReviewStatus  string `json:"license_review_status"`
	RedistributionStatus string `json:"redistribution_status"`
	InputCRS             string `json:"input_crs,omitempty"`
	Transformation       string `json:"transformation,omitempty"`
}

type Builder struct {
	Tool      string `json:"tool"`
	Version   string `json:"version"`
	Revision  string `json:"revision"`
	BuiltAt   string `json:"built_at"`
	OutputCRS string `json:"output_crs,omitempty"`
}

type RecordCounts struct {
	Accepted    int64 `json:"accepted"`
	Rejected    int64 `json:"rejected"`
	Duplicate   int64 `json:"duplicate"`
	Quarantined int64 `json:"quarantined"`
}

type IntegrityCheck struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Detail string `json:"detail,omitempty"`
}

type Manifest struct {
	SchemaVersion   string           `json:"schema_version"`
	SnapshotID      string           `json:"snapshot_id"`
	Kind            Kind             `json:"kind"`
	Sources         []SourceArtifact `json:"sources"`
	Builder         Builder          `json:"builder"`
	RecordCounts    RecordCounts     `json:"record_counts"`
	OutputSHA256    string           `json:"output_sha256"`
	OutputSizeBytes int64            `json:"output_size_bytes"`
	IntegrityChecks []IntegrityCheck `json:"integrity_checks"`
}

type ActivationReceipt struct {
	SchemaVersion           string `json:"schema_version"`
	SnapshotID              string `json:"snapshot_id"`
	ActivatedAt             string `json:"activated_at"`
	PriorActiveSnapshotID   string `json:"prior_active_snapshot_id,omitempty"`
	LastKnownGoodSnapshotID string `json:"last_known_good_snapshot_id"`
	ManifestSHA256          string `json:"manifest_sha256"`
}

type Verified struct {
	Manifest Manifest
	Receipt  ActivationReceipt
}

func ManifestPath(snapshotPath string) string {
	return snapshotPath + ".manifest.json"
}

func ActivationPath(snapshotPath string) string {
	return snapshotPath + ".activation.json"
}

func FinalizeAndWrite(snapshotPath string, manifest Manifest) (Manifest, error) {
	if manifest.Kind != KindBoundary && manifest.Kind != KindGeocoder {
		return Manifest{}, fmt.Errorf("%w: unsupported kind %q", ErrManifestInvalid, manifest.Kind)
	}
	digest, size, err := fileDigest(snapshotPath)
	if err != nil {
		return Manifest{}, err
	}
	manifest.OutputSHA256 = digest
	manifest.OutputSizeBytes = size
	if err := manifest.Validate(manifest.Kind); err != nil {
		return Manifest{}, err
	}
	manifestBytes, err := CanonicalJSON(manifest)
	if err != nil {
		return Manifest{}, fmt.Errorf("encode snapshot manifest: %w", err)
	}
	manifestBytes = append(manifestBytes, '\n')
	manifestPath := ManifestPath(snapshotPath)
	temporaryPath := manifestPath + ".tmp"
	if err := os.WriteFile(temporaryPath, manifestBytes, 0o644); err != nil {
		return Manifest{}, fmt.Errorf("write snapshot manifest: %w", err)
	}
	if err := os.Rename(temporaryPath, manifestPath); err != nil {
		_ = os.Remove(temporaryPath)
		return Manifest{}, fmt.Errorf("publish snapshot manifest: %w", err)
	}
	return manifest, nil
}

func LoadAndVerify(snapshotPath string, expectedKind Kind) (Verified, error) {
	manifest, manifestBytes, err := loadAndVerifySnapshot(snapshotPath, expectedKind)
	if err != nil {
		return Verified{}, err
	}

	receiptBytes, err := os.ReadFile(ActivationPath(snapshotPath))
	if err != nil {
		if os.IsNotExist(err) {
			return Verified{}, fmt.Errorf("%w: activation receipt missing", ErrManifestInvalid)
		}
		return Verified{}, fmt.Errorf("read activation receipt: %w", err)
	}
	var receipt ActivationReceipt
	if err := json.Unmarshal(receiptBytes, &receipt); err != nil {
		return Verified{}, fmt.Errorf("%w: decode activation receipt: %v", ErrManifestInvalid, err)
	}
	if err := receipt.Validate(manifest, manifestBytes); err != nil {
		return Verified{}, err
	}
	return Verified{Manifest: manifest, Receipt: receipt}, nil
}

func loadAndVerifySnapshot(snapshotPath string, expectedKind Kind) (Manifest, []byte, error) {
	manifestPath := ManifestPath(snapshotPath)
	manifestBytes, err := os.ReadFile(manifestPath)
	if err != nil {
		if os.IsNotExist(err) {
			return Manifest{}, nil, fmt.Errorf("%w: %s", ErrManifestMissing, manifestPath)
		}
		return Manifest{}, nil, fmt.Errorf("read snapshot manifest: %w", err)
	}

	var manifest Manifest
	if err := json.Unmarshal(manifestBytes, &manifest); err != nil {
		return Manifest{}, nil, fmt.Errorf("%w: decode manifest: %v", ErrManifestInvalid, err)
	}
	if err := manifest.Validate(expectedKind); err != nil {
		return Manifest{}, nil, err
	}

	digest, size, err := fileDigest(snapshotPath)
	if err != nil {
		return Manifest{}, nil, err
	}
	if !strings.EqualFold(digest, manifest.OutputSHA256) || size != manifest.OutputSizeBytes {
		return Manifest{}, nil, fmt.Errorf("%w: expected sha256=%s size=%d, got sha256=%s size=%d", ErrChecksumMismatch, manifest.OutputSHA256, manifest.OutputSizeBytes, digest, size)
	}
	return manifest, manifestBytes, nil
}

func (m Manifest) Validate(expectedKind Kind) error {
	if m.SchemaVersion != SchemaVersion {
		return fmt.Errorf("%w: unsupported schema_version %q", ErrManifestInvalid, m.SchemaVersion)
	}
	if m.Kind != expectedKind {
		return fmt.Errorf("%w: expected kind %q, got %q", ErrManifestInvalid, expectedKind, m.Kind)
	}
	if strings.TrimSpace(m.SnapshotID) == "" || len(m.Sources) == 0 {
		return fmt.Errorf("%w: snapshot_id and sources are required", ErrManifestInvalid)
	}
	if !validDigest(m.OutputSHA256) || m.OutputSizeBytes < 1 {
		return fmt.Errorf("%w: output checksum and size are required", ErrManifestInvalid)
	}
	if strings.TrimSpace(m.Builder.Tool) == "" || strings.TrimSpace(m.Builder.Version) == "" || strings.TrimSpace(m.Builder.Revision) == "" {
		return fmt.Errorf("%w: builder tool, version, and revision are required", ErrManifestInvalid)
	}
	if _, err := time.Parse(time.RFC3339, m.Builder.BuiltAt); err != nil {
		return fmt.Errorf("%w: builder built_at: %v", ErrManifestInvalid, err)
	}
	if m.RecordCounts.Accepted < 0 || m.RecordCounts.Rejected < 0 || m.RecordCounts.Duplicate < 0 || m.RecordCounts.Quarantined < 0 {
		return fmt.Errorf("%w: record counts must be non-negative", ErrManifestInvalid)
	}
	for index, source := range m.Sources {
		if strings.TrimSpace(source.Publisher) == "" || strings.TrimSpace(source.Product) == "" || strings.TrimSpace(source.Vintage) == "" || strings.TrimSpace(source.Locator) == "" || !validDigest(source.SHA256) {
			return fmt.Errorf("%w: source %d is incomplete", ErrManifestInvalid, index)
		}
		if strings.TrimSpace(source.LicenseReviewStatus) == "" || strings.TrimSpace(source.RedistributionStatus) == "" {
			return fmt.Errorf("%w: source %d licensing status is incomplete", ErrManifestInvalid, index)
		}
		if _, err := time.Parse(time.RFC3339, source.RetrievedAt); err != nil {
			return fmt.Errorf("%w: source %d retrieved_at: %v", ErrManifestInvalid, index, err)
		}
	}
	for _, check := range m.IntegrityChecks {
		if strings.TrimSpace(check.Name) == "" || (check.Status != "passed" && check.Status != "failed") {
			return fmt.Errorf("%w: invalid integrity check", ErrManifestInvalid)
		}
		if check.Status != "passed" {
			return fmt.Errorf("%w: integrity check %q did not pass", ErrManifestInvalid, check.Name)
		}
	}
	return nil
}

func (r ActivationReceipt) Validate(manifest Manifest, manifestBytes []byte) error {
	if r.SchemaVersion != SchemaVersion || r.SnapshotID != manifest.SnapshotID || r.LastKnownGoodSnapshotID == "" {
		return fmt.Errorf("%w: activation receipt identity mismatch", ErrManifestInvalid)
	}
	if _, err := time.Parse(time.RFC3339, r.ActivatedAt); err != nil {
		return fmt.Errorf("%w: activation timestamp: %v", ErrManifestInvalid, err)
	}
	digest := sha256.Sum256(manifestBytes)
	if !strings.EqualFold(r.ManifestSHA256, hex.EncodeToString(digest[:])) {
		return fmt.Errorf("%w: activation receipt manifest checksum mismatch", ErrManifestInvalid)
	}
	return nil
}

func CanonicalJSON(manifest Manifest) ([]byte, error) {
	sort.Slice(manifest.Sources, func(i, j int) bool {
		left := manifest.Sources[i].Publisher + "\x00" + manifest.Sources[i].Product + "\x00" + manifest.Sources[i].Layer + "\x00" + manifest.Sources[i].Locator
		right := manifest.Sources[j].Publisher + "\x00" + manifest.Sources[j].Product + "\x00" + manifest.Sources[j].Layer + "\x00" + manifest.Sources[j].Locator
		return left < right
	})
	sort.Slice(manifest.IntegrityChecks, func(i, j int) bool {
		return manifest.IntegrityChecks[i].Name < manifest.IntegrityChecks[j].Name
	})
	return json.MarshalIndent(manifest, "", "  ")
}

func fileDigest(path string) (string, int64, error) {
	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return "", 0, fmt.Errorf("open snapshot for checksum: %w", err)
	}
	defer file.Close()
	hash := sha256.New()
	size, err := io.Copy(hash, file)
	if err != nil {
		return "", 0, fmt.Errorf("hash snapshot: %w", err)
	}
	return hex.EncodeToString(hash.Sum(nil)), size, nil
}

func validDigest(value string) bool {
	decoded, err := hex.DecodeString(value)
	return err == nil && len(decoded) == sha256.Size
}
