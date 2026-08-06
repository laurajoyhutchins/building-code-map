package bundle

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
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

func TestLoadAndVerifyChecksRecursiveRegulatoryCatalog(t *testing.T) {
	root := t.TempDir()
	boundary := []byte(`{"kind":"boundary"}`)
	if err := os.WriteFile(filepath.Join(root, "boundary.json"), boundary, 0o644); err != nil {
		t.Fatal(err)
	}
	catalog := filepath.Join(root, "regulatory")
	if err := os.MkdirAll(filepath.Join(catalog, "rules"), 0o755); err != nil {
		t.Fatal(err)
	}
	for path, content := range map[string]string{
		filepath.Join(catalog, "colorado.json"):          `{"state":"CO"}`,
		filepath.Join(catalog, "rules", "colorado.json"): `{"rules":[]}`,
	} {
		if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	boundaryDigest := sha256.Sum256(boundary)
	catalogDigest, _, err := DirectoryDigest(catalog)
	if err != nil {
		t.Fatal(err)
	}
	manifest := Manifest{
		SchemaVersion: "1",
		EngineVersion: "0.1.0",
		SourceCommit:  "0123456789abcdef0123456789abcdef01234567",
		CreatedAt:     "2026-08-06T00:00:00Z",
		Components: map[string]Component{
			"boundary_snapshot": {
				Path:   "boundary.json",
				SHA256: hex.EncodeToString(boundaryDigest[:]),
			},
			"regulatory_catalog": {
				Path:      "regulatory",
				SHA256:    catalogDigest,
				Recursive: true,
			},
		},
		Coverage: Coverage{States: []string{"CO"}, AsOf: "2026-08-06"},
	}
	manifestPath := filepath.Join(root, "bundle.json")
	if err := WriteManifest(manifestPath, manifest); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadAndVerify(manifestPath); err != nil {
		t.Fatal(err)
	}

	if err := os.WriteFile(filepath.Join(catalog, "colorado.json"), []byte(`{"state":"CO","changed":true}`), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := LoadAndVerify(manifestPath); !errors.Is(err, ErrDigest) {
		t.Fatalf("LoadAndVerify() error = %v, want ErrDigest", err)
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

func TestManifestRejectsRecursiveNonCatalogComponent(t *testing.T) {
	manifest := Manifest{
		SchemaVersion: "1",
		EngineVersion: "0.1.0",
		SourceCommit:  "0123456789abcdef0123456789abcdef01234567",
		CreatedAt:     "2026-08-06T00:00:00Z",
		Components: map[string]Component{
			"boundary_snapshot": {
				Path:      "boundary",
				SHA256:    "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
				Recursive: true,
			},
			"regulatory_catalog": {
				Path:   "catalog.json",
				SHA256: "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
			},
		},
		Coverage: Coverage{AsOf: "2026-08-06"},
	}
	if err := manifest.Validate(); err == nil {
		t.Fatal("Validate() error = nil")
	}
}
