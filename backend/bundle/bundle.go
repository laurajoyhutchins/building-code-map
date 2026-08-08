package bundle

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const SchemaVersion = "1"

var (
	ErrInvalid = errors.New("data bundle invalid")
	ErrDigest  = errors.New("data bundle component digest mismatch")
)

type Component struct {
	Path         string `json:"path"`
	SHA256       string `json:"sha256"`
	ManifestPath string `json:"manifest_path,omitempty"`
	Optional     bool   `json:"optional,omitempty"`
	Recursive    bool   `json:"recursive,omitempty"`
}

type Coverage struct {
	States       []string `json:"states"`
	CodeFamilies []string `json:"code_families"`
	AsOf         string   `json:"as_of"`
}

type Manifest struct {
	SchemaVersion string               `json:"schema_version"`
	EngineVersion string               `json:"engine_version"`
	SourceCommit  string               `json:"source_commit"`
	CreatedAt     string               `json:"created_at"`
	Components    map[string]Component `json:"components"`
	Coverage      Coverage             `json:"coverage"`
}

type Verified struct {
	Manifest Manifest
	Digest   string
	Path     string
}

func (manifest Manifest) Validate() error {
	if manifest.SchemaVersion != SchemaVersion {
		return fmt.Errorf("%w: unsupported schema_version %q", ErrInvalid, manifest.SchemaVersion)
	}
	if strings.TrimSpace(manifest.EngineVersion) == "" || !isSHA256Commit(manifest.SourceCommit) {
		return fmt.Errorf("%w: engine_version and 40-hex source_commit are required", ErrInvalid)
	}
	if _, err := time.Parse(time.RFC3339, manifest.CreatedAt); err != nil {
		return fmt.Errorf("%w: created_at: %v", ErrInvalid, err)
	}
	if _, err := time.Parse(time.DateOnly, manifest.Coverage.AsOf); err != nil {
		return fmt.Errorf("%w: coverage.as_of must use YYYY-MM-DD", ErrInvalid)
	}
	for _, role := range []string{"boundary_snapshot", "regulatory_catalog"} {
		component, ok := manifest.Components[role]
		if !ok {
			return fmt.Errorf("%w: required component %q is missing", ErrInvalid, role)
		}
		if err := component.Validate(role); err != nil {
			return err
		}
	}
	for role, component := range manifest.Components {
		if role != "boundary_snapshot" && role != "regulatory_catalog" && role != "geocoder" {
			return fmt.Errorf("%w: unsupported component role %q", ErrInvalid, role)
		}
		if err := component.Validate(role); err != nil {
			return err
		}
	}
	return nil
}

func (component Component) Validate(role string) error {
	rawPath := strings.TrimSpace(component.Path)
	path := filepath.Clean(rawPath)
	pathParts := strings.Split(strings.ReplaceAll(rawPath, `\`, "/"), "/")
	for _, part := range pathParts {
		if part == ".." {
			return fmt.Errorf("%w: component %q has unsafe path", ErrInvalid, role)
		}
	}
	if path == "." || filepath.IsAbs(path) || path != filepath.ToSlash(path) || path == ".." || strings.HasPrefix(path, ".."+string(filepath.Separator)) {
		return fmt.Errorf("%w: component %q has unsafe path", ErrInvalid, role)
	}
	if !validDigest(component.SHA256) {
		return fmt.Errorf("%w: component %q has invalid sha256", ErrInvalid, role)
	}
	if component.Optional && role != "geocoder" {
		return fmt.Errorf("%w: only geocoder may be optional", ErrInvalid)
	}
	if component.Recursive && role != "regulatory_catalog" {
		return fmt.Errorf("%w: only regulatory_catalog may be recursive", ErrInvalid)
	}
	return nil
}

func LoadAndVerify(path string) (Verified, error) {
	cleanPath := filepath.Clean(path)
	raw, err := os.ReadFile(cleanPath)
	if err != nil {
		return Verified{}, fmt.Errorf("%w: read manifest: %v", ErrInvalid, err)
	}
	var manifest Manifest
	decoder := json.NewDecoder(strings.NewReader(string(raw)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&manifest); err != nil {
		return Verified{}, fmt.Errorf("%w: decode manifest: %v", ErrInvalid, err)
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return Verified{}, fmt.Errorf("%w: manifest must contain one JSON value", ErrInvalid)
	}
	if err := manifest.Validate(); err != nil {
		return Verified{}, err
	}
	root := filepath.Dir(cleanPath)
	seenPaths := map[string]string{}
	for role, component := range manifest.Components {
		componentPath := filepath.Join(root, filepath.FromSlash(component.Path))
		if previousRole, exists := seenPaths[component.Path]; exists {
			return Verified{}, fmt.Errorf("%w: components %q and %q reuse path %q", ErrInvalid, previousRole, role, component.Path)
		}
		seenPaths[component.Path] = role
		digest, _, digestErr := ComponentDigest(componentPath, component.Recursive)
		if digestErr != nil {
			if component.Optional && os.IsNotExist(digestErr) {
				continue
			}
			return Verified{}, fmt.Errorf("%w: component %q: %v", ErrInvalid, role, digestErr)
		}
		if !strings.EqualFold(normalizeDigest(digest), normalizeDigest(component.SHA256)) {
			return Verified{}, fmt.Errorf("%w: component %q", ErrDigest, role)
		}
		if component.ManifestPath != "" {
			manifestPath := filepath.Join(root, filepath.FromSlash(component.ManifestPath))
			if _, err := os.Stat(manifestPath); err != nil {
				return Verified{}, fmt.Errorf("%w: component %q manifest: %v", ErrInvalid, role, err)
			}
		}
	}
	digest := sha256.Sum256(raw)
	return Verified{Manifest: manifest, Digest: hex.EncodeToString(digest[:]), Path: cleanPath}, nil
}

func WriteManifest(path string, manifest Manifest) error {
	if err := manifest.Validate(); err != nil {
		return err
	}
	raw, err := CanonicalJSON(manifest)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, append(raw, '\n'), 0o644)
}

func CanonicalJSON(manifest Manifest) ([]byte, error) {
	roles := make([]string, 0, len(manifest.Components))
	for role := range manifest.Components {
		roles = append(roles, role)
	}
	sort.Strings(roles)
	ordered := make(map[string]Component, len(roles))
	for _, role := range roles {
		ordered[role] = manifest.Components[role]
	}
	manifest.Components = ordered
	return json.MarshalIndent(manifest, "", "  ")
}

func ComponentDigest(path string, recursive bool) (string, int64, error) {
	if !recursive {
		return FileDigest(path)
	}
	return DirectoryDigest(path)
}

func DirectoryDigest(path string) (string, int64, error) {
	root := filepath.Clean(path)
	info, err := os.Stat(root)
	if err != nil {
		return "", 0, err
	}
	if !info.IsDir() {
		return "", 0, fmt.Errorf("recursive component is not a directory")
	}

	var files []string
	if err := filepath.WalkDir(root, func(current string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.Type().IsRegular() {
			files = append(files, current)
		}
		return nil
	}); err != nil {
		return "", 0, err
	}
	if len(files) == 0 {
		return "", 0, fmt.Errorf("recursive component directory is empty")
	}
	sort.Strings(files)

	hash := sha256.New()
	var total int64
	for _, file := range files {
		relative, err := filepath.Rel(root, file)
		if err != nil {
			return "", 0, err
		}
		digest, size, err := FileDigest(file)
		if err != nil {
			return "", 0, err
		}
		if _, err := fmt.Fprintf(hash, "%s\x00%s\n", filepath.ToSlash(relative), digest); err != nil {
			return "", 0, err
		}
		total += size
	}
	return hex.EncodeToString(hash.Sum(nil)), total, nil
}

func FileDigest(path string) (string, int64, error) {
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

func validDigest(value string) bool {
	value = normalizeDigest(value)
	if len(value) != sha256.Size*2 {
		return false
	}
	_, err := hex.DecodeString(value)
	return err == nil
}

func normalizeDigest(value string) string {
	return strings.TrimPrefix(strings.ToLower(strings.TrimSpace(value)), "sha256:")
}

func isSHA256Commit(value string) bool {
	if len(value) != 40 {
		return false
	}
	_, err := hex.DecodeString(value)
	return err == nil
}
