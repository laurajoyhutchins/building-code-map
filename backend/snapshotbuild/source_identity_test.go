package snapshotbuild

import (
	"context"
	"os"
	"testing"
)

func TestBuildManifestUsesPreparedSourceIdentity(t *testing.T) {
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
		t.Fatalf("build reread a source after preparation: %v", err)
	}
	if len(result.Manifest.Sources) != 1 {
		t.Fatalf("unexpected manifest sources: %+v", result.Manifest.Sources)
	}
	if result.Manifest.Sources[0].SHA256 != preparedDigest {
		t.Fatalf("manifest digest = %q, want prepared digest %q", result.Manifest.Sources[0].SHA256, preparedDigest)
	}
	if result.Receipt.Sources[0].SHA256 != preparedDigest {
		t.Fatalf("receipt digest = %q, want prepared digest %q", result.Receipt.Sources[0].SHA256, preparedDigest)
	}
}
