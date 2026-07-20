package regulatory

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Catalog struct {
	byStateID map[string]StateProfile
	byFIPS    map[string]StateProfile
	profiles  []StateProfile
}

func EmptyCatalog() Catalog {
	return Catalog{byStateID: map[string]StateProfile{}, byFIPS: map[string]StateProfile{}}
}

func NewCatalog(profiles []StateProfile) (Catalog, error) {
	catalog := EmptyCatalog()
	catalog.profiles = append([]StateProfile(nil), profiles...)
	for _, profile := range profiles {
		if err := ValidateProfile(profile); err != nil {
			return Catalog{}, err
		}
		stateKey := strings.ToUpper(strings.TrimSpace(profile.StateID))
		if _, exists := catalog.byStateID[stateKey]; exists {
			return Catalog{}, fmt.Errorf("duplicate state profile for %s", stateKey)
		}
		if _, exists := catalog.byFIPS[profile.StateFIPS]; exists {
			return Catalog{}, fmt.Errorf("duplicate state profile for FIPS %s", profile.StateFIPS)
		}
		catalog.byStateID[stateKey] = profile
		catalog.byFIPS[profile.StateFIPS] = profile
	}
	sort.Slice(catalog.profiles, func(i, j int) bool { return catalog.profiles[i].StateID < catalog.profiles[j].StateID })
	return catalog, nil
}

func LoadCatalog(dir string) (Catalog, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return Catalog{}, err
	}
	profiles := make([]StateProfile, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() || strings.ToLower(filepath.Ext(entry.Name())) != ".json" {
			continue
		}
		path := filepath.Join(dir, entry.Name())
		raw, err := os.ReadFile(path)
		if err != nil {
			return Catalog{}, err
		}
		var profile StateProfile
		decoder := json.NewDecoder(strings.NewReader(string(raw)))
		decoder.DisallowUnknownFields()
		if err := decoder.Decode(&profile); err != nil {
			return Catalog{}, fmt.Errorf("decode %s: %w", path, err)
		}
		profiles = append(profiles, profile)
	}
	if len(profiles) == 0 {
		return Catalog{}, fmt.Errorf("no regulatory profiles found in %s", dir)
	}
	return NewCatalog(profiles)
}

func (catalog Catalog) Profile(stateID, fips string) (StateProfile, bool) {
	if key := strings.ToUpper(strings.TrimSpace(stateID)); key != "" {
		if profile, ok := catalog.byStateID[key]; ok {
			return profile, true
		}
	}
	if profile, ok := catalog.byFIPS[strings.TrimSpace(fips)]; ok {
		return profile, true
	}
	return StateProfile{}, false
}

func (catalog Catalog) Profiles() []StateProfile {
	return append([]StateProfile(nil), catalog.profiles...)
}

func (catalog Catalog) Len() int { return len(catalog.profiles) }

func WriteProfile(path string, profile StateProfile) error {
	if err := ValidateProfile(profile); err != nil {
		return err
	}
	raw, err := json.MarshalIndent(profile, "", "  ")
	if err != nil {
		return err
	}
	raw = append(raw, '\n')
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.WriteFile(path, raw, fs.FileMode(0o644))
}
