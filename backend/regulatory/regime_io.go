package regulatory

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
)

func BuildStateRegimeSpecs(catalog Catalog) ([]StateRegimeSpec, error) {
	profiles := catalog.Profiles()
	if len(profiles) == 0 {
		return nil, fmt.Errorf("regulatory catalog contains no state profiles")
	}
	specs := make([]StateRegimeSpec, 0, len(profiles))
	for _, profile := range profiles {
		spec, err := BuildStateRegimeSpec(profile)
		if err != nil {
			return nil, fmt.Errorf("build regime for %s: %w", profile.StateID, err)
		}
		specs = append(specs, spec)
	}
	sort.Slice(specs, func(i, j int) bool { return specs[i].StateID < specs[j].StateID })
	return specs, nil
}

func WriteStateRegimeDirectory(root string, specs []StateRegimeSpec) error {
	if len(specs) == 0 {
		return fmt.Errorf("no state regime specs to write")
	}
	root = filepath.Clean(root)
	parent := filepath.Dir(root)
	if err := os.MkdirAll(parent, 0o755); err != nil {
		return fmt.Errorf("create state regime parent: %w", err)
	}
	temporary, err := os.MkdirTemp(parent, ".state-regimes-")
	if err != nil {
		return fmt.Errorf("create temporary state regime directory: %w", err)
	}
	published := false
	defer func() {
		if !published {
			_ = os.RemoveAll(temporary)
		}
	}()

	seen := map[string]bool{}
	for _, spec := range specs {
		if err := ValidateStateRegimeSpec(spec); err != nil {
			return fmt.Errorf("validate state regime %s: %w", spec.StateID, err)
		}
		if seen[spec.StateID] {
			return fmt.Errorf("duplicate state regime %s", spec.StateID)
		}
		seen[spec.StateID] = true
		data, err := json.MarshalIndent(spec, "", "  ")
		if err != nil {
			return fmt.Errorf("encode state regime %s: %w", spec.StateID, err)
		}
		data = append(data, '\n')
		if err := os.WriteFile(filepath.Join(temporary, spec.StateID+".json"), data, 0o644); err != nil {
			return fmt.Errorf("write state regime %s: %w", spec.StateID, err)
		}
	}

	backup := root + ".previous"
	_ = os.RemoveAll(backup)
	if _, statErr := os.Stat(root); statErr == nil {
		if err := os.Rename(root, backup); err != nil {
			return fmt.Errorf("preserve prior state regime directory: %w", err)
		}
	} else if !os.IsNotExist(statErr) {
		return fmt.Errorf("inspect state regime directory: %w", statErr)
	}
	if err := os.Rename(temporary, root); err != nil {
		if _, backupErr := os.Stat(backup); backupErr == nil {
			_ = os.Rename(backup, root)
		}
		return fmt.Errorf("publish state regime directory: %w", err)
	}
	published = true
	_ = os.RemoveAll(backup)
	return nil
}

func LoadStateRegimeDirectory(root string) ([]StateRegimeSpec, error) {
	entries, err := os.ReadDir(filepath.Clean(root))
	if err != nil {
		return nil, err
	}
	var specs []StateRegimeSpec
	for _, entry := range entries {
		if entry.IsDir() || strings.ToLower(filepath.Ext(entry.Name())) != ".json" {
			continue
		}
		path := filepath.Join(root, entry.Name())
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		decoder := json.NewDecoder(file)
		decoder.DisallowUnknownFields()
		var spec StateRegimeSpec
		decodeErr := decoder.Decode(&spec)
		closeErr := file.Close()
		if decodeErr != nil {
			return nil, fmt.Errorf("decode %s: %w", path, decodeErr)
		}
		if closeErr != nil {
			return nil, closeErr
		}
		if err := ValidateStateRegimeSpec(spec); err != nil {
			return nil, fmt.Errorf("validate %s: %w", path, err)
		}
		specs = append(specs, spec)
	}
	if len(specs) == 0 {
		return nil, fmt.Errorf("no state regime specs found in %s", root)
	}
	sort.Slice(specs, func(i, j int) bool { return specs[i].StateID < specs[j].StateID })
	return specs, nil
}

func ValidateStateRegimeDirectory(catalog Catalog, root string) error {
	expected, err := BuildStateRegimeSpecs(catalog)
	if err != nil {
		return err
	}
	actual, err := LoadStateRegimeDirectory(root)
	if err != nil {
		return err
	}
	if !reflect.DeepEqual(actual, expected) {
		return fmt.Errorf("state regime projections are stale; regenerate from current regulatory profiles")
	}
	return nil
}
