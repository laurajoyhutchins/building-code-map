package bundle

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadAndVerifyChecksComponentDigestsAndIdentity(t *testing.T) {
	root := t.TempDir()
	component := []byte(`{"kind":"boundary"}`)
	if err := os.WriteFile(filepath.Join(root, "boundary.json"), component, 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "catalog.json"), component, 0o644); err != nil {
		t.Fatal(err)
	}
	digest := sha256.Sum256(component)
	manifest := Manifest{
		SchemaVersion: "1",
		EngineVersion: "0.1.0",
		SourceCommit:  "0123456789abcdef0123456789abcdef01234567",
		CreatedAt:     "2026-08-06T00:00:00Z",
		Components: map[string]Component{
			"boundary_snapshot":  {Path: "boundary.json", SHA256: hex.EncodeToString(digest[:])},
			"regulatory_catalog": {Path: "catalog.json", SHA256: hex.EncodeToString(digest[:])},
		},
		Coverage: Coverage{States: []string{"CO"}, AsOf: "2026-08-06"},
	}
	manifestPath := filepath.Join(root, "bundle.json")
	if err := WriteManifest(manifestPath, manifest); err != nil {
		t.Fatal(err)
	}
	verified, err := LoadAndVerify(manifestPath)
	if err != nil {
		t.Fatal(err)
	}
	if verified.Manifest.SourceCommit != manifest.SourceCommit || verified.Digest == "" {
		t.Fatalf("verified=%#v", verified)
	}
}

func TestManifestRejectsPathTraversalAndUnsupportedSchema(t *testing.T) {
	manifest := Manifest{
		SchemaVersion: "2",
		EngineVersion: "0.1.0",
		SourceCommit:  "0123456789abcdef0123456789abcdef01234567",
		CreatedAt:     "2026-08-06T00:00:00Z",
		Components: map[string]Component{
			"boundary_snapshot":  {Path: "../boundary.json", SHA256: "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"},
			"regulatory_catalog": {Path: "catalog.json", SHA256: "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef"},
		},
	}
	if err := manifest.Validate(); err == nil {
		t.Fatal("Validate() error = nil")
	}
}
