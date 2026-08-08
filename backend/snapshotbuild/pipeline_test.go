package snapshotbuild

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
	"time"
)

type fakeExecutor struct {
	version    string
	output     []byte
	invocation Invocation
}

func (f *fakeExecutor) Version(context.Context, string) (string, error) {
	return f.version, nil
}

func (f *fakeExecutor) Run(_ context.Context, invocation Invocation) ([]byte, error) {
	f.invocation = invocation
	matches := regexp.MustCompile(`(?i)\bTO\s+'([^']+)'`).FindAllStringSubmatch(invocation.SQL, -1)
	for _, match := range matches {
		path := filepath.FromSlash(strings.ReplaceAll(match[1], "''", "'"))
		if err := os.WriteFile(path, []byte("header\nvalue\n"), 0o644); err != nil {
			return nil, err
		}
	}
	return f.output, nil
}

func TestAuditBindsExplicitDuckDBAndContractIdentity(t *testing.T) {
	request := fixtureRequest(t, KindGeocoder)
	executor := &fakeExecutor{
		version: "DuckDB v1.5.0",
		output:  []byte(`[{"accepted":2,"rejected":1,"duplicate":1,"quarantined":0}]`),
	}
	report, receipt, err := (Pipeline{Executor: executor}).Audit(context.Background(), request)
	if err != nil {
		t.Fatal(err)
	}
	if report.Counts.Accepted != 2 || report.Counts.Rejected != 1 || report.Counts.Duplicate != 1 {
		t.Fatalf("unexpected counts: %+v", report.Counts)
	}
	if receipt.Tool.Pipeline != PipelineIdentity || receipt.Tool.DuckDBVersion != "DuckDB v1.5.0" {
		t.Fatalf("unexpected tool identity: %+v", receipt.Tool)
	}
	if len(receipt.Tool.DuckDBSHA256) != 64 || len(receipt.Tool.SQLContractSHA256) != 64 {
		t.Fatalf("missing content identities: %+v", receipt.Tool)
	}
	if strings.Contains(strings.ToUpper(executor.invocation.SQL), "INSTALL ") {
		t.Fatal("contract attempted implicit extension installation")
	}
	if strings.Contains(executor.invocation.SQL, "https://") {
		t.Fatal("contract contained a remote source")
	}
	if receipt.Sources[0].Locator == request.Sources[0].Path {
		t.Fatalf("receipt leaked machine-local source path: %+v", receipt.Sources[0])
	}
}

func TestBoundaryAuditLoadsOnlyExplicitSpatialExtension(t *testing.T) {
	request := fixtureRequest(t, KindBoundary)
	extension := filepath.Join(t.TempDir(), "spatial.duckdb_extension")
	if err := os.WriteFile(extension, []byte("extension"), 0o644); err != nil {
		t.Fatal(err)
	}
	request.SpatialExtensionPath = extension
	extensionDigest, _, err := digestFile(extension)
	if err != nil {
		t.Fatal(err)
	}
	request.ExpectedSpatialExtensionSHA256 = extensionDigest
	executor := &fakeExecutor{
		version: "DuckDB v1.5.0",
		output:  []byte(`[{"accepted":1,"rejected":0,"duplicate":0,"quarantined":0}]`),
	}
	report, receipt, err := (Pipeline{Executor: executor}).Audit(context.Background(), request)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(executor.invocation.SQL, "LOAD '"+filepath.ToSlash(extension)+"';") {
		t.Fatalf("explicit extension was not loaded:\n%s", executor.invocation.SQL)
	}
	if strings.Contains(strings.ToUpper(executor.invocation.SQL), "INSTALL ") {
		t.Fatal("contract attempted to install an extension")
	}
	if receipt.Tool.SpatialExtensionSHA256 == "" {
		t.Fatal("extension digest was not recorded")
	}
	found := false
	for _, check := range report.Checks {
		if check.Name == "geometry_validity" && check.Status == "passed" {
			found = true
		}
	}
	if !found {
		t.Fatalf("geometry validity evidence missing: %+v", report.Checks)
	}
}

func TestBuildFinalizesOutputAndWritesPR45Manifest(t *testing.T) {
	request := fixtureRequest(t, KindGeocoder)
	request.OutputPath = filepath.Join(t.TempDir(), "geocoder.sqlite")
	executor := &fakeExecutor{
		version: "DuckDB v1.5.0",
		output:  []byte(`[{"accepted":1,"rejected":0,"duplicate":0,"quarantined":0}]`),
	}
	result, err := (Pipeline{Executor: executor}).Build(context.Background(), request, func(_ context.Context, normalized map[string]string, output string) error {
		path := normalized["address_points"]
		if path == "" {
			t.Fatal("normalized address points were not provided")
		}
		if _, err := os.ReadFile(path); err != nil {
			t.Fatalf("normalized address points were removed before finalization: %v", err)
		}
		return os.WriteFile(output, []byte("canonical sqlite bytes"), 0o644)
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Manifest.Kind != "geocoder" {
		t.Fatalf("unexpected manifest kind: %s", result.Manifest.Kind)
	}
	if result.Manifest.Builder.Tool != "bcm/"+PipelineIdentity {
		t.Fatalf("unexpected builder tool: %+v", result.Manifest.Builder)
	}
	if !strings.Contains(result.Manifest.Builder.Revision, "duckdb=") || !strings.Contains(result.Manifest.Builder.Revision, "sql=") {
		t.Fatalf("builder revision does not bind tool inputs: %q", result.Manifest.Builder.Revision)
	}
	manifestPath := request.OutputPath + ".manifest.json"
	raw, err := os.ReadFile(manifestPath)
	if err != nil {
		t.Fatal(err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded["snapshot_id"] == "" {
		t.Fatal("manifest lacks snapshot identity")
	}
}

func TestAuditRejectsRemoteAndImplicitInstallContracts(t *testing.T) {
	if err := validateContract("INSTALL spatial;"); err == nil {
		t.Fatal("INSTALL should be rejected")
	}
	request := fixtureRequest(t, KindGeocoder)
	request.Sources[0].Path = "https://example.com/addresses.csv"
	_, _, err := (Pipeline{Executor: &fakeExecutor{}}).Audit(context.Background(), request)
	if err == nil {
		t.Fatal("remote source should be rejected")
	}
}

func TestAuditRejectsUnexpectedDuckDBVersionOrDigest(t *testing.T) {
	request := fixtureRequest(t, KindGeocoder)
	executor := &fakeExecutor{
		version: "DuckDB v9.9.9",
		output:  []byte(`[{"accepted":1,"rejected":0,"duplicate":0,"quarantined":0}]`),
	}
	if _, _, err := (Pipeline{Executor: executor}).Audit(context.Background(), request); err == nil {
		t.Fatal("unexpected DuckDB version should fail")
	}

	request = fixtureRequest(t, KindGeocoder)
	request.ExpectedDuckDBSHA256 = strings.Repeat("0", 64)
	executor.version = "DuckDB v1.5.0"
	if _, _, err := (Pipeline{Executor: executor}).Audit(context.Background(), request); err == nil {
		t.Fatal("unexpected DuckDB digest should fail")
	}
}

func TestBoundaryBuildRequiresVerifiedSpatialExtension(t *testing.T) {
	request := fixtureRequest(t, KindBoundary)
	executor := &fakeExecutor{
		version: "DuckDB v1.5.0",
		output:  []byte(`[{"accepted":1,"rejected":0,"duplicate":0,"quarantined":0}]`),
	}
	_, err := (Pipeline{Executor: executor}).Build(context.Background(), request, func(context.Context, map[string]string, string) error {
		return nil
	})
	if err == nil {
		t.Fatal("boundary build without a verified spatial extension should fail")
	}
}

func TestDecodeCountsUsesFinalDuckDBJSONResult(t *testing.T) {
	raw := []byte("[{\"Count\":1}]\n[{\"accepted\":3,\"rejected\":2,\"duplicate\":1,\"quarantined\":0}]\n")
	counts, err := decodeCounts(raw)
	if err != nil {
		t.Fatal(err)
	}
	if counts.Accepted != 3 || counts.Rejected != 2 || counts.Duplicate != 1 {
		t.Fatalf("unexpected counts: %+v", counts)
	}
}

func fixtureRequest(t *testing.T, kind Kind) Request {
	t.Helper()
	root := t.TempDir()
	duckdb := filepath.Join(root, "duckdb")
	if err := os.WriteFile(duckdb, []byte("duckdb binary"), 0o755); err != nil {
		t.Fatal(err)
	}
	work := filepath.Join(root, "work")
	if err := os.MkdirAll(work, 0o755); err != nil {
		t.Fatal(err)
	}
	source := func(name string) Source {
		path := filepath.Join(root, name+".csv")
		if err := os.WriteFile(path, []byte("header\nvalue\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		return Source{
			Name: name, Path: path, Locator: "fixture://" + name, Publisher: "Test publisher", Product: "Test product",
			Layer: name, Vintage: "2026", RetrievedAt: "2026-08-06T12:00:00Z",
			LicenseReviewStatus: "reviewed", RedistributionStatus: "fixture-only",
		}
	}
	sources := []Source{source("address_points")}
	if kind == KindBoundary {
		sources = []Source{source("layer_families"), source("boundary_features"), source("refresh_status")}
	}
	duckDigest, _, err := digestFile(duckdb)
	if err != nil {
		t.Fatal(err)
	}
	return Request{
		Kind: kind, DuckDBPath: duckdb, ExpectedDuckDBVersion: "1.5.0",
		ExpectedDuckDBSHA256: duckDigest, WorkingDirectory: work,
		OutputPath:     filepath.Join(root, "output.sqlite"),
		BuilderVersion: "0.1.0", BuilderRevision: "abc123",
		BuiltAt:   time.Date(2026, 8, 6, 12, 0, 0, 0, time.UTC),
		OutputCRS: "EPSG:4326", Sources: sources,
	}
}
