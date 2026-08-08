package snapshot

import (
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	_ "modernc.org/sqlite"
)

type SQLiteBuildOptions struct {
	OutputPath          string
	LayerFamiliesCSV    string
	BoundaryFeaturesCSV string
	RefreshStatusCSV    string
}

const sqliteSchema = `
PRAGMA page_size = 4096;
PRAGMA auto_vacuum = NONE;
PRAGMA journal_mode = OFF;
PRAGMA synchronous = OFF;
CREATE TABLE layer_families (
  key TEXT PRIMARY KEY NOT NULL,
  label TEXT NOT NULL,
  martin_layer_id TEXT NOT NULL,
  description TEXT NOT NULL,
  default_enabled INTEGER NOT NULL CHECK (default_enabled IN (0, 1))
);
CREATE TABLE boundary_features (
  layer_family TEXT NOT NULL,
  feature_id TEXT NOT NULL,
  title TEXT NOT NULL,
  subtitle TEXT NOT NULL,
  source_id TEXT NOT NULL,
  geometry_label TEXT NOT NULL,
  geometry_source TEXT,
  last_synced_at TEXT NOT NULL,
  geometry_json TEXT NOT NULL,
  attributes_json TEXT NOT NULL,
  PRIMARY KEY (layer_family, feature_id),
  FOREIGN KEY (layer_family) REFERENCES layer_families(key)
);
CREATE INDEX boundary_features_layer_family ON boundary_features(layer_family, feature_id);
CREATE TABLE refresh_status (
  status TEXT NOT NULL,
  latest_successful_refresh TEXT NOT NULL,
  latest_attempt TEXT NOT NULL,
  next_scheduled_refresh TEXT NOT NULL,
  message TEXT NOT NULL
);
`

func BuildSQLite(options SQLiteBuildOptions) error {
	for label, value := range map[string]string{
		"output path":           options.OutputPath,
		"layer families CSV":    options.LayerFamiliesCSV,
		"boundary features CSV": options.BoundaryFeaturesCSV,
		"refresh status CSV":    options.RefreshStatusCSV,
	} {
		if strings.TrimSpace(value) == "" {
			return fmt.Errorf("%s is required", label)
		}
	}
	outputPath, err := filepath.Abs(options.OutputPath)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		return err
	}
	temp, err := os.CreateTemp(filepath.Dir(outputPath), ".boundary-*.sqlite")
	if err != nil {
		return err
	}
	tempPath := temp.Name()
	if err := temp.Close(); err != nil {
		return err
	}
	if err := os.Remove(tempPath); err != nil {
		return err
	}
	defer os.Remove(tempPath)

	db, err := sql.Open("sqlite", tempPath)
	if err != nil {
		return err
	}
	closed := false
	defer func() {
		if !closed {
			_ = db.Close()
		}
	}()
	if _, err := db.Exec(sqliteSchema); err != nil {
		return fmt.Errorf("create boundary snapshot schema: %w", err)
	}
	if _, err := db.Exec("PRAGMA foreign_keys = ON;"); err != nil {
		return err
	}
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	if err := importLayerFamilies(tx, options.LayerFamiliesCSV); err != nil {
		_ = tx.Rollback()
		return err
	}
	if err := importBoundaryFeatures(tx, options.BoundaryFeaturesCSV); err != nil {
		_ = tx.Rollback()
		return err
	}
	if err := importRefreshStatus(tx, options.RefreshStatusCSV); err != nil {
		_ = tx.Rollback()
		return err
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	if _, err := db.Exec("VACUUM;"); err != nil {
		return fmt.Errorf("vacuum boundary snapshot: %w", err)
	}
	if err := db.Close(); err != nil {
		return err
	}
	closed = true
	if _, err := LoadSQLite(tempPath); err != nil {
		return fmt.Errorf("validate built boundary snapshot: %w", err)
	}
	return replaceSQLiteSnapshot(tempPath, outputPath)
}

func importLayerFamilies(tx *sql.Tx, path string) error {
	return readBuildCSV(path, []string{"key", "label", "martin_layer_id", "description", "default_enabled"}, func(row map[string]string, line int) error {
		enabled, err := strconv.ParseBool(strings.TrimSpace(row["default_enabled"]))
		if err != nil {
			value, intErr := strconv.Atoi(strings.TrimSpace(row["default_enabled"]))
			if intErr != nil || (value != 0 && value != 1) {
				return fmt.Errorf("%s line %d: invalid default_enabled", path, line)
			}
			enabled = value == 1
		}
		if strings.TrimSpace(row["key"]) == "" || strings.TrimSpace(row["label"]) == "" {
			return fmt.Errorf("%s line %d: key and label are required", path, line)
		}
		_, err = tx.Exec(`
INSERT INTO layer_families(key, label, martin_layer_id, description, default_enabled)
VALUES (?, ?, ?, ?, ?);`,
			strings.TrimSpace(row["key"]), strings.TrimSpace(row["label"]),
			strings.TrimSpace(row["martin_layer_id"]), strings.TrimSpace(row["description"]), boolInt(enabled),
		)
		return err
	})
}

func importBoundaryFeatures(tx *sql.Tx, path string) error {
	required := []string{
		"layer_family", "feature_id", "title", "subtitle", "source_id", "geometry_label",
		"geometry_source", "last_synced_at", "geometry_json", "attributes_json",
	}
	return readBuildCSV(path, required, func(row map[string]string, line int) error {
		for _, field := range []string{"layer_family", "feature_id", "title", "source_id", "geometry_label", "last_synced_at", "geometry_json", "attributes_json"} {
			if strings.TrimSpace(row[field]) == "" {
				return fmt.Errorf("%s line %d: %s is required", path, line, field)
			}
		}
		_, err := tx.Exec(`
INSERT INTO boundary_features(
  layer_family, feature_id, title, subtitle, source_id, geometry_label,
  geometry_source, last_synced_at, geometry_json, attributes_json
) VALUES (?, ?, ?, ?, ?, ?, NULLIF(?, ''), ?, ?, ?);`,
			strings.TrimSpace(row["layer_family"]), strings.TrimSpace(row["feature_id"]),
			strings.TrimSpace(row["title"]), strings.TrimSpace(row["subtitle"]),
			strings.TrimSpace(row["source_id"]), strings.TrimSpace(row["geometry_label"]),
			strings.TrimSpace(row["geometry_source"]), strings.TrimSpace(row["last_synced_at"]),
			strings.TrimSpace(row["geometry_json"]), strings.TrimSpace(row["attributes_json"]),
		)
		if err != nil {
			return fmt.Errorf("%s line %d: %w", path, line, err)
		}
		return nil
	})
}

func importRefreshStatus(tx *sql.Tx, path string) error {
	count := 0
	err := readBuildCSV(path, []string{
		"status", "latest_successful_refresh", "latest_attempt", "next_scheduled_refresh", "message",
	}, func(row map[string]string, line int) error {
		count++
		if count > 1 {
			return fmt.Errorf("%s: expected exactly one refresh status row", path)
		}
		_, err := tx.Exec(`
INSERT INTO refresh_status(
  status, latest_successful_refresh, latest_attempt, next_scheduled_refresh, message
) VALUES (?, ?, ?, ?, ?);`,
			strings.TrimSpace(row["status"]), strings.TrimSpace(row["latest_successful_refresh"]),
			strings.TrimSpace(row["latest_attempt"]), strings.TrimSpace(row["next_scheduled_refresh"]),
			strings.TrimSpace(row["message"]),
		)
		return err
	})
	if err != nil {
		return err
	}
	if count != 1 {
		return fmt.Errorf("%s: expected exactly one refresh status row", path)
	}
	return nil
}

func readBuildCSV(path string, required []string, visit func(map[string]string, int) error) error {
	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		return err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true
	header, err := reader.Read()
	if err != nil {
		return err
	}
	indexes := map[string]int{}
	for index, name := range header {
		indexes[strings.TrimSpace(name)] = index
	}
	for _, name := range required {
		if _, ok := indexes[name]; !ok {
			return fmt.Errorf("%s: missing required column %q", path, name)
		}
	}
	for line := 2; ; line++ {
		record, err := reader.Read()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return fmt.Errorf("%s line %d: %w", path, line, err)
		}
		row := map[string]string{}
		for name, index := range indexes {
			if index < len(record) {
				row[name] = record[index]
			}
		}
		if err := visit(row, line); err != nil {
			return err
		}
	}
}

func replaceSQLiteSnapshot(tempPath, outputPath string) error {
	if err := os.Rename(tempPath, outputPath); err == nil {
		return nil
	}
	if info, err := os.Stat(outputPath); err != nil {
		return fmt.Errorf("replace boundary snapshot: %w", err)
	} else if info.IsDir() {
		return errors.New("replace boundary snapshot: output path is a directory")
	}
	backup := outputPath + ".previous"
	_ = os.Remove(backup)
	if err := os.Rename(outputPath, backup); err != nil {
		return fmt.Errorf("backup boundary snapshot: %w", err)
	}
	if err := os.Rename(tempPath, outputPath); err != nil {
		restoreErr := os.Rename(backup, outputPath)
		if restoreErr != nil {
			return errors.Join(fmt.Errorf("replace boundary snapshot: %w", err), fmt.Errorf("restore previous boundary snapshot: %w", restoreErr))
		}
		return fmt.Errorf("replace boundary snapshot: %w", err)
	}
	if err := os.Remove(backup); err != nil && !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("remove previous boundary snapshot: %w", err)
	}
	return nil
}

func boolInt(value bool) int {
	if value {
		return 1
	}
	return 0
}
