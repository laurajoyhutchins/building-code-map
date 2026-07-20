package regulatory

import (
	"encoding/json"
	"fmt"
	"io"
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
		if err := ValidateCompleteProfile(profile); err != nil {
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
	profiles, err := loadProfiles(dir)
	if err != nil {
		return Catalog{}, err
	}
	packs, err := loadRulePacks(filepath.Join(dir, "rules"))
	if err != nil {
		return Catalog{}, err
	}

	profileIndex := make(map[string]int, len(profiles))
	for index, profile := range profiles {
		profileIndex[strings.ToUpper(strings.TrimSpace(profile.StateID))] = index
	}
	for _, pack := range packs {
		key := strings.ToUpper(strings.TrimSpace(pack.StateID))
		index, ok := profileIndex[key]
		if !ok {
			return Catalog{}, fmt.Errorf("rule pack references unknown state profile %s", pack.StateID)
		}
		merged, err := MergeRulePack(profiles[index], pack)
		if err != nil {
			return Catalog{}, err
		}
		profiles[index] = merged
	}
	return NewCatalog(profiles)
}

func loadProfiles(dir string) ([]StateProfile, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	profiles := make([]StateProfile, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() || strings.ToLower(filepath.Ext(entry.Name())) != ".json" {
			continue
		}
		path := filepath.Join(dir, entry.Name())
		var profile StateProfile
		if err := decodeJSONFile(path, &profile); err != nil {
			return nil, err
		}
		profiles = append(profiles, profile)
	}
	if len(profiles) == 0 {
		return nil, fmt.Errorf("no regulatory profiles found in %s", dir)
	}
	return profiles, nil
}

func loadRulePacks(dir string) ([]StateRulePack, error) {
	entries, err := os.ReadDir(dir)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	packs := make([]StateRulePack, 0, len(entries))
	seen := map[string]struct{}{}
	for _, entry := range entries {
		if entry.IsDir() || strings.ToLower(filepath.Ext(entry.Name())) != ".json" {
			continue
		}
		path := filepath.Join(dir, entry.Name())
		var pack StateRulePack
		if err := decodeJSONFile(path, &pack); err != nil {
			return nil, err
		}
		key := strings.ToUpper(strings.TrimSpace(pack.StateID))
		if key == "" {
			return nil, fmt.Errorf("decode %s: state_id is required", path)
		}
		if _, exists := seen[key]; exists {
			return nil, fmt.Errorf("duplicate rule pack for %s", key)
		}
		seen[key] = struct{}{}
		packs = append(packs, pack)
	}
	sort.Slice(packs, func(i, j int) bool { return packs[i].StateID < packs[j].StateID })
	return packs, nil
}

func decodeJSONFile(path string, target any) error {
	raw, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(strings.NewReader(string(raw)))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(target); err != nil {
		return fmt.Errorf("decode %s: %w", path, err)
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		if err == nil {
			return fmt.Errorf("decode %s: multiple JSON values", path)
		}
		return fmt.Errorf("decode %s: %w", path, err)
	}
	return nil
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
	if err := ValidateCompleteProfile(profile); err != nil {
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
