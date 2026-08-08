package snapshotbuild

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type sourceMutationExecutor struct {
	fakeExecutor
	originalPath  string
	originalBytes []byte
	mutatedBytes  []byte
}

func (f *sourceMutationExecutor) Run(ctx context.Context, invocation Invocation) ([]byte, error) {
	if err := os.WriteFile(f.originalPath, f.mutatedBytes, 0o644); err != nil {
		return nil, err
	}
	if strings.Contains(invocation.SQL, filepath.ToSlash(f.originalPath)) {
		return nil, os.ErrInvalid
	}
	stagedPath := filepath.Join(invocation.WorkingDirectory, "sources", "000"+filepath.Ext(f.originalPath))
	stagedBytes, err := os.ReadFile(stagedPath)
	if err != nil {
		return nil, err
	}
	if string(stagedBytes) != string(f.originalBytes) {
		return nil, os.ErrInvalid
	}
	return f.fakeExecutor.Run(ctx, invocation)
}

func TestBuildManifestUsesStagedSourceIdentity(t *testing.T) {
	request := fixtureRequest(t, KindGeocoder)
	originalBytes, err := os.ReadFile(request.Sources[0].Path)
	if err != nil {
		t.Fatal(err)
	}
	preparedDigest, _, err := digestFile(request.Sources[0].Path)
	if err != nil {
		t.Fatal(err)
	}
	executor := &sourceMutationExecutor{
		fakeExecutor: fakeExecutor{
			version: "DuckDB v1.5.0",
			output:  []byte(`[{"accepted":1,"rejected":0,"duplicate":0,"quarantined":0}]`),
		},
		originalPath:  request.Sources[0].Path,
		originalBytes: originalBytes,
		mutatedBytes:  []byte("mutated after staging\n"),
	}

	result, err := (Pipeline{Executor: executor}).Build(context.Background(), request, func(_ context.Context, _ map[string]string, output string) error {
		return os.WriteFile(output, []byte("canonical sqlite bytes"), 0o644)
	})
	if err != nil {
		t.Fatalf("build did not consume the staged source: %v", err)
	}
	if len(result.Manifest.Sources) != 1 {
		t.Fatalf("unexpected manifest sources: %+v", result.Manifest.Sources)
	}
	if result.Manifest.Sources[0].SHA256 != preparedDigest {
		t.Fatalf("manifest digest = %q, want staged digest %q", result.Manifest.Sources[0].SHA256, preparedDigest)
	}
	if result.Receipt.Sources[0].SHA256 != preparedDigest {
		t.Fatalf("receipt digest = %q, want staged digest %q", result.Receipt.Sources[0].SHA256, preparedDigest)
	}
}

func TestBuildManifestDoesNotRereadOriginalSource(t *testing.T) {
	request := fixtureRequest(t, KindGeocoder)
	preparedDigest, _, err := digestFile(request.Sources[0].Path)
	if err != nil {
		t.Fatal(err)
	}
	executor := &fakeExecutor{
		version: "DuckDB v1.5.0",
		output:  []byte(`[{"accepted":1,"rejected":0,"duplicate":0,"quarantined":0}]`),
	}

	result, err := (Pipeline{Executor: executor}).Build(context.Background(), request, func(_ context.Context, _ map[string]string, output string) error {
		if err := os.Remove(request.Sources[0].Path); err != nil {
			return err
		}
		return os.WriteFile(output, []byte("canonical sqlite bytes"), 0o644)
	})
	if err != nil {
		t.Fatalf("build reread an original source after staging: %v", err)
	}
	if result.Manifest.Sources[0].SHA256 != preparedDigest {
		t.Fatalf("manifest digest = %q, want staged digest %q", result.Manifest.Sources[0].SHA256, preparedDigest)
	}
}
