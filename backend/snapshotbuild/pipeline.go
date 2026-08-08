package snapshotbuild

import (
	"context"
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

	"building-code-map/backend/internal/snapshotmanifest"
)

var (
	ErrInvalidRequest = errors.New("invalid snapshot build request")
	ErrUnsafeContract = errors.New("unsafe DuckDB contract")
)

type Pipeline struct {
	Executor Executor
}

func (p Pipeline) Audit(ctx context.Context, request Request) (AuditReport, Receipt, error) {
	report, receipt, cleanup, err := p.execute(ctx, request, false)
	if err != nil {
		return AuditReport{}, Receipt{}, err
	}
	defer cleanup()
	report.Normalized = nil
	return report, receipt, nil
}

func (p Pipeline) Build(ctx context.Context, request Request, finalize FinalizeFunc) (Result, error) {
	if request.Kind == KindBoundary && strings.TrimSpace(request.SpatialExtensionPath) == "" {
		return Result{}, fmt.Errorf("%w: boundary build requires an explicitly provisioned spatial extension", ErrInvalidRequest)
	}
	if finalize == nil {
		return Result{}, fmt.Errorf("%w: finalizer is required", ErrInvalidRequest)
	}
	report, receipt, cleanup, err := p.execute(ctx, request, true)
	if err != nil {
		return Result{}, err
	}
	defer cleanup()
	if err := finalize(ctx, report.Normalized, request.OutputPath); err != nil {
		return Result{}, fmt.Errorf("finalize canonical SQLite snapshot: %w", err)
	}
	outputDigest, outputSize, err := digestFile(request.OutputPath)
	if err != nil {
		return Result{}, fmt.Errorf("hash canonical SQLite snapshot: %w", err)
	}
	receipt.OutputSHA256 = outputDigest
	receipt.OutputSize = outputSize

	manifest, err := buildManifest(request, receipt)
	if err != nil {
		return Result{}, err
	}
	manifest, err = snapshotmanifest.FinalizeAndWrite(request.OutputPath, manifest)
	if err != nil {
		return Result{}, fmt.Errorf("finalize snapshot manifest: %w", err)
	}
	report.Normalized = nil
	return Result{Report: report, Receipt: receipt, Manifest: manifest}, nil
}

func (p Pipeline) execute(ctx context.Context, request Request, requireOutput bool) (AuditReport, Receipt, func(), error) {
	prepared, err := prepare(request, requireOutput)
	if err != nil {
		return AuditReport{}, Receipt{}, nil, err
	}
	executor := p.Executor
	if executor == nil {
		executor = CommandExecutor{}
	}

	version, err := executor.Version(ctx, prepared.request.DuckDBPath)
	if err != nil {
		return AuditReport{}, Receipt{}, nil, fmt.Errorf("read DuckDB version: %w", err)
	}
	if !versionMatches(version, prepared.request.ExpectedDuckDBVersion) {
		return AuditReport{}, Receipt{}, nil, fmt.Errorf("%w: expected DuckDB version %q, got %q", ErrInvalidRequest, prepared.request.ExpectedDuckDBVersion, strings.TrimSpace(version))
	}
	duckDigest, _, err := digestFile(prepared.request.DuckDBPath)
	if err != nil {
		return AuditReport{}, Receipt{}, nil, fmt.Errorf("hash DuckDB executable: %w", err)
	}
	if !strings.EqualFold(duckDigest, prepared.request.ExpectedDuckDBSHA256) {
		return AuditReport{}, Receipt{}, nil, fmt.Errorf("%w: DuckDB executable digest mismatch", ErrInvalidRequest)
	}
	var extensionDigest string
	if prepared.request.SpatialExtensionPath != "" {
		extensionDigest, _, err = digestFile(prepared.request.SpatialExtensionPath)
		if err != nil {
			return AuditReport{}, Receipt{}, nil, fmt.Errorf("hash DuckDB spatial extension: %w", err)
		}
		if !strings.EqualFold(extensionDigest, prepared.request.ExpectedSpatialExtensionSHA256) {
			return AuditReport{}, Receipt{}, nil, fmt.Errorf("%w: DuckDB spatial extension digest mismatch", ErrInvalidRequest)
		}
	}

	runRoot, err := os.MkdirTemp(prepared.request.WorkingDirectory, ".bcm-duckdb-build-*")
	if err != nil {
		return AuditReport{}, Receipt{}, nil, fmt.Errorf("create DuckDB work directory: %w", err)
	}
	cleanup := func() { _ = os.RemoveAll(runRoot) }
	fail := func(err error) (AuditReport, Receipt, func(), error) {
		cleanup()
		return AuditReport{}, Receipt{}, nil, err
	}

	prepared, err = stageSources(prepared, runRoot)
	if err != nil {
		return fail(err)
	}
	sqlText, normalized, err := renderContract(prepared.request, runRoot)
	if err != nil {
		return fail(err)
	}
	if err := validateContract(sqlText); err != nil {
		return fail(err)
	}
	databasePath := filepath.Join(runRoot, "workspace.duckdb")
	output, err := executor.Run(ctx, Invocation{
		Executable:         prepared.request.DuckDBPath,
		DatabasePath:       databasePath,
		SQL:                sqlText,
		WorkingDirectory:   runRoot,
		ExtensionDirectory: extensionDirectory(prepared.request.SpatialExtensionPath),
	})
	if err != nil {
		return fail(fmt.Errorf("run DuckDB snapshot contract: %w", err))
	}
	counts, err := decodeCounts(output)
	if err != nil {
		return fail(err)
	}
	for name, path := range normalized {
		if _, err := os.Stat(path); err != nil {
			return fail(fmt.Errorf("normalized output %q: %w", name, err))
		}
	}

	checks := []Check{
		{Name: "duckdb_contract_executed", Status: "passed"},
		{Name: "duplicates_classified", Status: "passed", Detail: fmt.Sprintf("%d duplicate records excluded", counts.Duplicate)},
		{Name: "rejections_classified", Status: "passed", Detail: fmt.Sprintf("%d rejected records excluded", counts.Rejected)},
		{Name: "quarantine_classified", Status: "passed", Detail: fmt.Sprintf("%d quarantined records excluded", counts.Quarantined)},
	}
	if prepared.request.Kind == KindBoundary {
		status := "not_run"
		detail := "spatial extension was not supplied"
		if prepared.request.SpatialExtensionPath != "" {
			status = "passed"
			detail = "geometry validity was evaluated by the explicitly provisioned spatial extension"
		}
		checks = append(checks, Check{Name: "geometry_validity", Status: status, Detail: detail})
	}
	sort.Slice(checks, func(i, j int) bool { return checks[i].Name < checks[j].Name })

	report := AuditReport{
		SchemaVersion: ReceiptSchemaVersion,
		Kind:          prepared.request.Kind,
		Counts:        counts,
		Checks:        checks,
		Normalized:    normalized,
	}
	receipt := Receipt{
		SchemaVersion: ReceiptSchemaVersion,
		Kind:          prepared.request.Kind,
		BuiltAt:       prepared.request.BuiltAt.UTC().Format("2006-01-02T15:04:05Z"),
		Tool: ToolIdentity{
			Pipeline:               PipelineIdentity,
			DuckDBVersion:          strings.TrimSpace(version),
			DuckDBSHA256:           duckDigest,
			SpatialExtensionSHA256: extensionDigest,
			SQLContractSHA256:      contractDigest(prepared.request.Kind),
		},
		Sources: prepared.sourceReceipts,
		Counts:  counts,
		Checks:  checks,
	}
	return report, receipt, cleanup, nil
}

func extensionDirectory(path string) string {
	if strings.TrimSpace(path) == "" {
		return ""
	}
	return filepath.Dir(path)
}

type preparedRequest struct {
	request        Request
	sourceReceipts []SourceReceipt
}

func prepare(request Request, requireOutput bool) (preparedRequest, error) {
	if request.Kind != KindBoundary && request.Kind != KindGeocoder {
		return preparedRequest{}, fmt.Errorf("%w: kind must be boundary or geocoder", ErrInvalidRequest)
	}
	required := map[string]string{
		"DuckDB path":       request.DuckDBPath,
		"DuckDB version":    request.ExpectedDuckDBVersion,
		"DuckDB SHA-256":    request.ExpectedDuckDBSHA256,
		"working directory": request.WorkingDirectory,
		"builder version":   request.BuilderVersion,
		"builder revision":  request.BuilderRevision,
	}
	if requireOutput {
		required["output path"] = request.OutputPath
	}
	for name, value := range required {
		if strings.TrimSpace(value) == "" {
			return preparedRequest{}, fmt.Errorf("%w: %s is required", ErrInvalidRequest, name)
		}
	}
	if request.BuiltAt.IsZero() {
		return preparedRequest{}, fmt.Errorf("%w: built_at is required", ErrInvalidRequest)
	}
	if !validSHA256(request.ExpectedDuckDBSHA256) {
		return preparedRequest{}, fmt.Errorf("%w: expected DuckDB SHA-256 must be a 64-character hex digest", ErrInvalidRequest)
	}
	if !filepath.IsAbs(request.DuckDBPath) {
		return preparedRequest{}, fmt.Errorf("%w: DuckDB path must be absolute", ErrInvalidRequest)
	}
	if err := validateLocalFile(request.DuckDBPath, "DuckDB executable"); err != nil {
		return preparedRequest{}, err
	}
	if request.SpatialExtensionPath != "" {
		if !validSHA256(request.ExpectedSpatialExtensionSHA256) {
			return preparedRequest{}, fmt.Errorf("%w: expected spatial extension SHA-256 must be provided", ErrInvalidRequest)
		}
		if !filepath.IsAbs(request.SpatialExtensionPath) {
			return preparedRequest{}, fmt.Errorf("%w: spatial extension path must be absolute", ErrInvalidRequest)
		}
		if err := validateLocalFile(request.SpatialExtensionPath, "spatial extension"); err != nil {
			return preparedRequest{}, err
		}
	}
	if len(request.Sources) == 0 {
		return preparedRequest{}, fmt.Errorf("%w: at least one source is required", ErrInvalidRequest)
	}
	seen := map[string]bool{}
	for index, source := range request.Sources {
		if strings.TrimSpace(source.Name) == "" || strings.TrimSpace(source.Path) == "" {
			return preparedRequest{}, fmt.Errorf("%w: source %d name and path are required", ErrInvalidRequest, index)
		}
		if strings.TrimSpace(source.Publisher) == "" || strings.TrimSpace(source.Product) == "" ||
			strings.TrimSpace(source.Vintage) == "" || strings.TrimSpace(source.Locator) == "" ||
			strings.TrimSpace(source.RetrievedAt) == "" || strings.TrimSpace(source.LicenseReviewStatus) == "" ||
			strings.TrimSpace(source.RedistributionStatus) == "" {
			return preparedRequest{}, fmt.Errorf("%w: source %d provenance metadata is incomplete", ErrInvalidRequest, index)
		}
		if seen[source.Name] {
			return preparedRequest{}, fmt.Errorf("%w: duplicate source name %q", ErrInvalidRequest, source.Name)
		}
		seen[source.Name] = true
		if err := validateLocalFile(source.Path, "source "+source.Name); err != nil {
			return preparedRequest{}, err
		}
		if strings.Contains(source.Path, "://") {
			return preparedRequest{}, fmt.Errorf("%w: remote source locators are not executable inputs", ErrInvalidRequest)
		}
	}
	if err := requireSourceNames(request.Kind, seen); err != nil {
		return preparedRequest{}, err
	}
	return preparedRequest{request: request}, nil
}

func stageSources(prepared preparedRequest, runRoot string) (preparedRequest, error) {
	sourcesRoot := filepath.Join(runRoot, "sources")
	if err := os.MkdirAll(sourcesRoot, 0o700); err != nil {
		return preparedRequest{}, fmt.Errorf("create staged source directory: %w", err)
	}
	stagedRequest := prepared.request
	stagedRequest.Sources = append([]Source(nil), prepared.request.Sources...)
	receipts := make([]SourceReceipt, 0, len(stagedRequest.Sources))
	for index := range stagedRequest.Sources {
		source := &stagedRequest.Sources[index]
		stagedPath := filepath.Join(sourcesRoot, fmt.Sprintf("%03d%s", index, filepath.Ext(source.Path)))
		digest, err := copyAndDigest(source.Path, stagedPath)
		if err != nil {
			return preparedRequest{}, fmt.Errorf("stage source %q: %w", source.Name, err)
		}
		source.Path = stagedPath
		receipts = append(receipts, SourceReceipt{Name: source.Name, Locator: source.Locator, SHA256: digest})
	}
	sort.Slice(receipts, func(i, j int) bool { return receipts[i].Name < receipts[j].Name })
	return preparedRequest{request: stagedRequest, sourceReceipts: receipts}, nil
}

func copyAndDigest(sourcePath, stagedPath string) (string, error) {
	source, err := os.Open(filepath.Clean(sourcePath))
	if err != nil {
		return "", err
	}
	defer source.Close()
	staged, err := os.OpenFile(stagedPath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o600)
	if err != nil {
		return "", err
	}
	removeStaged := true
	defer func() {
		_ = staged.Close()
		if removeStaged {
			_ = os.Remove(stagedPath)
		}
	}()
	hash := sha256.New()
	if _, err := io.Copy(io.MultiWriter(staged, hash), source); err != nil {
		return "", err
	}
	if err := staged.Sync(); err != nil {
		return "", err
	}
	if err := staged.Close(); err != nil {
		return "", err
	}
	removeStaged = false
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func requireSourceNames(kind Kind, seen map[string]bool) error {
	switch kind {
	case KindGeocoder:
		if !seen["address_points"] && !seen["street_ranges"] {
			return fmt.Errorf("%w: geocoder requires address_points or street_ranges", ErrInvalidRequest)
		}
	case KindBoundary:
		for _, name := range []string{"layer_families", "boundary_features", "refresh_status"} {
			if !seen[name] {
				return fmt.Errorf("%w: boundary requires source %q", ErrInvalidRequest, name)
			}
		}
	}
	return nil
}

func validateLocalFile(path, label string) error {
	info, err := os.Stat(filepath.Clean(path))
	if err != nil {
		return fmt.Errorf("%w: %s: %v", ErrInvalidRequest, label, err)
	}
	if !info.Mode().IsRegular() {
		return fmt.Errorf("%w: %s must be a regular file", ErrInvalidRequest, label)
	}
	return nil
}

func decodeCounts(raw []byte) (Counts, error) {
	lines := strings.Split(strings.TrimSpace(string(raw)), "\n")
	for index := len(lines) - 1; index >= 0; index-- {
		candidate := strings.TrimSpace(lines[index])
		if candidate == "" {
			continue
		}
		var rows []Counts
		if err := json.Unmarshal([]byte(candidate), &rows); err != nil || len(rows) != 1 {
			continue
		}
		counts := rows[0]
		if counts.Accepted < 0 || counts.Rejected < 0 || counts.Duplicate < 0 || counts.Quarantined < 0 {
			return Counts{}, errors.New("decode DuckDB audit output: record counts must be non-negative")
		}
		return counts, nil
	}
	return Counts{}, errors.New("decode DuckDB audit output: no single-row audit result found")
}

func buildManifest(request Request, receipt Receipt) (snapshotmanifest.Manifest, error) {
	digests := make(map[string]string, len(receipt.Sources))
	for _, source := range receipt.Sources {
		digests[source.Name] = source.SHA256
	}
	sources := make([]snapshotmanifest.SourceArtifact, 0, len(request.Sources))
	for _, source := range request.Sources {
		digest, ok := digests[source.Name]
		if !ok || !validSHA256(digest) {
			return snapshotmanifest.Manifest{}, fmt.Errorf("prepared source digest missing or invalid for %q", source.Name)
		}
		sources = append(sources, snapshotmanifest.SourceArtifact{
			Publisher:            source.Publisher,
			Product:              source.Product,
			Layer:                source.Layer,
			Vintage:              source.Vintage,
			Locator:              source.Locator,
			SHA256:               digest,
			RetrievedAt:          source.RetrievedAt,
			LicenseReviewStatus:  source.LicenseReviewStatus,
			RedistributionStatus: source.RedistributionStatus,
			InputCRS:             source.InputCRS,
			Transformation:       source.Transformation,
		})
	}
	kind := snapshotmanifest.KindGeocoder
	if request.Kind == KindBoundary {
		kind = snapshotmanifest.KindBoundary
	}
	checks := make([]snapshotmanifest.IntegrityCheck, 0, len(receipt.Checks)+1)
	for _, check := range receipt.Checks {
		if check.Status == "not_run" {
			continue
		}
		checks = append(checks, snapshotmanifest.IntegrityCheck{Name: check.Name, Status: check.Status, Detail: check.Detail})
	}
	checks = append(checks, snapshotmanifest.IntegrityCheck{Name: "canonical_sqlite_finalized", Status: "passed"})
	manifest := snapshotmanifest.Manifest{
		SchemaVersion: snapshotmanifest.SchemaVersion,
		SnapshotID:    string(request.Kind) + "-" + receipt.OutputSHA256[:16],
		Kind:          kind,
		Sources:       sources,
		Builder: snapshotmanifest.Builder{
			Tool:      "bcm/" + PipelineIdentity,
			Version:   request.BuilderVersion,
			Revision:  builderRevision(request.BuilderRevision, receipt.Tool),
			BuiltAt:   request.BuiltAt.UTC().Format("2006-01-02T15:04:05Z"),
			OutputCRS: request.OutputCRS,
		},
		RecordCounts: snapshotmanifest.RecordCounts{
			Accepted: receipt.Counts.Accepted, Rejected: receipt.Counts.Rejected,
			Duplicate: receipt.Counts.Duplicate, Quarantined: receipt.Counts.Quarantined,
		},
		OutputSHA256:    receipt.OutputSHA256,
		OutputSizeBytes: receipt.OutputSize,
		IntegrityChecks: checks,
	}
	if err := manifest.Validate(kind); err != nil {
		return snapshotmanifest.Manifest{}, fmt.Errorf("validate generated snapshot manifest: %w", err)
	}
	return manifest, nil
}

func digestFile(path string) (string, int64, error) {
	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return "", 0, err
	}
	defer file.Close()
	hash := sha256.New()
	size, err := io.Copy(hash, file)
	if err != nil {
		return "", 0, err
	}
	return hex.EncodeToString(hash.Sum(nil)), size, nil
}

func atomicWrite(path string, content []byte) error {
	absolute, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(absolute), 0o755); err != nil {
		return err
	}
	temp, err := os.CreateTemp(filepath.Dir(absolute), "."+filepath.Base(absolute)+".tmp-*")
	if err != nil {
		return err
	}
	tempPath := temp.Name()
	defer os.Remove(tempPath)
	if _, err := temp.Write(content); err != nil {
		temp.Close()
		return err
	}
	if err := temp.Sync(); err != nil {
		temp.Close()
		return err
	}
	if err := temp.Close(); err != nil {
		return err
	}
	return os.Rename(tempPath, absolute)
}

func WriteReceipt(path string, receipt Receipt) error {
	raw, err := json.MarshalIndent(receipt, "", "  ")
	if err != nil {
		return err
	}
	return atomicWrite(path, append(raw, '\n'))
}

func builderRevision(revision string, tool ToolIdentity) string {
	parts := []string{
		strings.TrimSpace(revision),
		"duckdb=" + tool.DuckDBSHA256[:16],
		"sql=" + tool.SQLContractSHA256[:16],
	}
	if tool.SpatialExtensionSHA256 != "" {
		parts = append(parts, "spatial="+tool.SpatialExtensionSHA256[:16])
	}
	return strings.Join(parts, ";")
}

func validSHA256(value string) bool {
	decoded, err := hex.DecodeString(strings.TrimSpace(value))
	return err == nil && len(decoded) == sha256.Size
}

func versionMatches(actual, expected string) bool {
	actual = strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(actual), "DuckDB"))
	actual = strings.TrimSpace(strings.TrimPrefix(actual, "v"))
	expected = strings.TrimSpace(strings.TrimPrefix(expected, "v"))
	fields := strings.Fields(actual)
	if len(fields) > 0 {
		actual = fields[0]
	}
	return actual == expected
}
