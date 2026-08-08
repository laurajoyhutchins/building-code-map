package snapshot

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	_ "modernc.org/sqlite"
)

type Snapshot struct {
	LayerFamilies    []LayerFamily     `json:"layer_families"`
	BoundaryFeatures []BoundaryFeature `json:"boundary_features"`
	RefreshStatus    RefreshStatus     `json:"refresh_status"`
}

type LayerFamily struct {
	Key            string `json:"key"`
	Label          string `json:"label"`
	MartinLayerID  string `json:"martin_layer_id"`
	Description    string `json:"description"`
	DefaultEnabled bool   `json:"default_enabled"`
}

type Geometry struct {
	Type        string `json:"type"`
	Coordinates any    `json:"coordinates"`
}

type FeatureRecord struct {
	LayerFamily    string         `json:"layer_family"`
	FeatureID      string         `json:"feature_id"`
	Title          string         `json:"title"`
	Subtitle       string         `json:"subtitle"`
	SourceID       string         `json:"source_id"`
	GeometryLabel  string         `json:"geometry_label"`
	GeometrySource string         `json:"geometry_source,omitempty"`
	LastSyncedAt   time.Time      `json:"last_synced_at"`
	Attributes     map[string]any `json:"attributes"`
}

type BoundaryFeature struct {
	LayerFamily    string         `json:"layer_family"`
	FeatureID      string         `json:"feature_id"`
	Title          string         `json:"title"`
	Subtitle       string         `json:"subtitle"`
	SourceID       string         `json:"source_id"`
	GeometryLabel  string         `json:"geometry_label"`
	GeometrySource string         `json:"geometry_source,omitempty"`
	LastSyncedAt   time.Time      `json:"last_synced_at"`
	Geometry       Geometry       `json:"geometry"`
	Attributes     map[string]any `json:"attributes"`
}

type RefreshStatus struct {
	Status                  string    `json:"status"`
	LatestSuccessfulRefresh time.Time `json:"latest_successful_refresh"`
	LatestAttempt           time.Time `json:"latest_attempt"`
	NextScheduledRefresh    time.Time `json:"next_scheduled_refresh"`
	Message                 string    `json:"message"`
}

type storedLayerFamily struct {
	Key            string `json:"key"`
	Label          string `json:"label"`
	MartinLayerID  string `json:"martin_layer_id"`
	Description    string `json:"description"`
	DefaultEnabled bool   `json:"default_enabled"`
}

type storedBoundaryFeature struct {
	LayerFamily    string `json:"layer_family"`
	FeatureID      string `json:"feature_id"`
	Title          string `json:"title"`
	Subtitle       string `json:"subtitle"`
	SourceID       string `json:"source_id"`
	GeometryLabel  string `json:"geometry_label"`
	GeometrySource string `json:"geometry_source"`
	LastSyncedAt   string `json:"last_synced_at"`
	GeometryJSON   string `json:"geometry_json"`
	AttributesJSON string `json:"attributes_json"`
}

type storedRefreshStatus struct {
	Status                  string `json:"status"`
	LatestSuccessfulRefresh string `json:"latest_successful_refresh"`
	LatestAttempt           string `json:"latest_attempt"`
	NextScheduledRefresh    string `json:"next_scheduled_refresh"`
	Message                 string `json:"message"`
}

var (
	ErrUnsupportedSnapshotFormat = errors.New("unsupported snapshot format")
	errInvalidSnapshot            = errors.New("invalid snapshot")
)

var nerisLayerFamily = LayerFamily{
	Key:            "neris_jurisdictions",
	Label:          "NERIS jurisdictions",
	MartinLayerID:  "neris.department_jurisdictions",
	Description:    "NERIS department jurisdiction polygons retained as geographic observations, not automatic building-code authority.",
	DefaultEnabled: false,
}

func (feature BoundaryFeature) Record() FeatureRecord {
	return FeatureRecord{
		LayerFamily:    feature.LayerFamily,
		FeatureID:      feature.FeatureID,
		Title:          feature.Title,
		Subtitle:       feature.Subtitle,
		SourceID:       feature.SourceID,
		GeometryLabel:  feature.GeometryLabel,
		GeometrySource: feature.GeometrySource,
		LastSyncedAt:   feature.LastSyncedAt,
		Attributes:     feature.Attributes,
	}
}

// LoadFile loads a supported boundary snapshot and rejects unknown formats
// before attempting any compatibility-tool discovery.
func LoadFile(path string) (Snapshot, error) {
	extension := strings.ToLower(filepath.Ext(path))
	switch extension {
	case ".sqlite", ".db":
		return LoadSQLite(path)
	case ".duckdb":
		return LoadDuckDB(path)
	default:
		return Snapshot{}, fmt.Errorf("%w: %q", ErrUnsupportedSnapshotFormat, extension)
	}
}

func LoadSQLite(path string) (Snapshot, error) {
	if _, err := os.Stat(path); err != nil {
		return Snapshot{}, fmt.Errorf("open SQLite snapshot %s: %w", path, err)
	}

	dsn, err := sqliteReadOnlyDSN(path)
	if err != nil {
		return Snapshot{}, err
	}
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return Snapshot{}, err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return Snapshot{}, err
	}

	layerRows, err := querySQLiteLayerFamilies(db)
	if err != nil {
		return Snapshot{}, err
	}
	boundaryRows, err := querySQLiteBoundaryFeatures(db)
	if err != nil {
		return Snapshot{}, err
	}
	refreshRows, err := querySQLiteRefreshStatus(db)
	if err != nil {
		return Snapshot{}, err
	}
	return buildValidatedSnapshot(layerRows, boundaryRows, refreshRows)
}

func LoadDuckDB(path string) (Snapshot, error) {
	cachePath, err := filepath.Abs(path)
	if err != nil {
		return Snapshot{}, err
	}
	if _, err := os.Stat(cachePath); err != nil {
		return Snapshot{}, fmt.Errorf("open DuckDB snapshot %s: %w", cachePath, err)
	}

	duckdbPath, err := findDuckDBExecutable()
	if err != nil {
		return Snapshot{}, err
	}

	layerRows, err := queryDuckDBRows[storedLayerFamily](duckdbPath, cachePath, `
SELECT
  key,
  label,
  martin_layer_id,
  description,
  default_enabled
FROM layer_families
ORDER BY key;
`)
	if err != nil {
		return Snapshot{}, err
	}

	boundaryRows, err := queryDuckDBRows[storedBoundaryFeature](duckdbPath, cachePath, `
SELECT
  layer_family,
  feature_id,
  title,
  subtitle,
  source_id,
  geometry_label,
  geometry_source,
  strftime(last_synced_at, '%Y-%m-%dT%H:%M:%SZ') AS last_synced_at,
  geometry_json,
  attributes_json
FROM boundary_features
ORDER BY layer_family, feature_id;
`)
	if err != nil {
		return Snapshot{}, err
	}

	refreshRows, err := queryDuckDBRows[storedRefreshStatus](duckdbPath, cachePath, `
SELECT
  status,
  strftime(latest_successful_refresh, '%Y-%m-%dT%H:%M:%SZ') AS latest_successful_refresh,
  strftime(latest_attempt, '%Y-%m-%dT%H:%M:%SZ') AS latest_attempt,
  strftime(next_scheduled_refresh, '%Y-%m-%dT%H:%M:%SZ') AS next_scheduled_refresh,
  message
FROM refresh_status;
`)
	if err != nil {
		return Snapshot{}, err
	}

	return buildValidatedSnapshot(layerRows, boundaryRows, refreshRows)
}

func buildValidatedSnapshot(layerRows []storedLayerFamily, boundaryRows []storedBoundaryFeature, refreshRows []storedRefreshStatus) (Snapshot, error) {
	if len(refreshRows) != 1 {
		return Snapshot{}, fmt.Errorf("%w: expected one refresh_status row, got %d", errInvalidSnapshot, len(refreshRows))
	}

	boundaryFeatures := make([]BoundaryFeature, 0, len(boundaryRows))
	for _, row := range boundaryRows {
		feature, err := row.toBoundaryFeature()
		if err != nil {
			return Snapshot{}, err
		}
		boundaryFeatures = append(boundaryFeatures, feature)
	}

	refreshStatus, err := refreshRows[0].toRefreshStatus()
	if err != nil {
		return Snapshot{}, err
	}

	snap := Snapshot{
		LayerFamilies:    ensureNerisLayerFamily(convertLayerFamilies(layerRows)),
		BoundaryFeatures: boundaryFeatures,
		RefreshStatus:    refreshStatus,
	}
	if err := snap.Validate(); err != nil {
		return Snapshot{}, err
	}
	return snap, nil
}

// DefaultCachePath is intentionally workspace-local and SQLite-first. Legacy
// DuckDB snapshots remain supported only when selected explicitly.
func DefaultCachePath(repoRoot string) string {
	return filepath.Join(repoRoot, "backend", "data", "tigerweb.sqlite")
}

func ResolveCachePath(repoRoot string) string {
	repoRootAbs, err := filepath.Abs(repoRoot)
	if err != nil {
		return DefaultCachePath(repoRoot)
	}

	for _, envName := range []string{"TIGERWEB_CACHE_PATH", "TIGERWEB_HYDRATED_CACHE_PATH"} {
		if resolved, ok := resolveWorkspacePath(repoRootAbs, os.Getenv(envName)); ok {
			return resolved
		}
	}
	return DefaultCachePath(repoRoot)
}

func sqliteReadOnlyDSN(path string) (string, error) {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	slashPath := filepath.ToSlash(absolutePath)
	if filepath.VolumeName(absolutePath) != "" && !strings.HasPrefix(slashPath, "/") {
		slashPath = "/" + slashPath
	}
	location := url.URL{Scheme: "file", Path: slashPath}
	query := location.Query()
	query.Set("mode", "ro")
	query.Add("_pragma", "query_only(1)")
	location.RawQuery = query.Encode()
	return location.String(), nil
}

func queryDuckDBRows[T any](duckdbPath string, dbPath string, statement string) ([]T, error) {
	cmd := exec.Command(duckdbPath, dbPath, "-json", "-c", statement)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("duckdb query failed: %w: %s", err, strings.TrimSpace(string(output)))
	}

	var rows []T
	if err := json.Unmarshal(output, &rows); err != nil {
		return nil, fmt.Errorf("decode duckdb output: %w", err)
	}
	return rows, nil
}

func querySQLiteLayerFamilies(db *sql.DB) ([]storedLayerFamily, error) {
	rows, err := db.Query(`
SELECT
  key,
  label,
  martin_layer_id,
  description,
  default_enabled
FROM layer_families
ORDER BY key;
`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []storedLayerFamily
	for rows.Next() {
		var row storedLayerFamily
		var defaultEnabled int64
		if err := rows.Scan(&row.Key, &row.Label, &row.MartinLayerID, &row.Description, &defaultEnabled); err != nil {
			return nil, err
		}
		row.DefaultEnabled = defaultEnabled != 0
		result = append(result, row)
	}
	return result, rows.Err()
}

func querySQLiteBoundaryFeatures(db *sql.DB) ([]storedBoundaryFeature, error) {
	rows, err := db.Query(`
SELECT
  layer_family,
  feature_id,
  title,
  subtitle,
  source_id,
  geometry_label,
  geometry_source,
  last_synced_at,
  geometry_json,
  attributes_json
FROM boundary_features
ORDER BY layer_family, feature_id;
`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []storedBoundaryFeature
	for rows.Next() {
		var row storedBoundaryFeature
		var geometrySource sql.NullString
		if err := rows.Scan(
			&row.LayerFamily,
			&row.FeatureID,
			&row.Title,
			&row.Subtitle,
			&row.SourceID,
			&row.GeometryLabel,
			&geometrySource,
			&row.LastSyncedAt,
			&row.GeometryJSON,
			&row.AttributesJSON,
		); err != nil {
			return nil, err
		}
		if geometrySource.Valid {
			row.GeometrySource = geometrySource.String
		}
		result = append(result, row)
	}
	return result, rows.Err()
}

func querySQLiteRefreshStatus(db *sql.DB) ([]storedRefreshStatus, error) {
	rows, err := db.Query(`
SELECT
  status,
  latest_successful_refresh,
  latest_attempt,
  next_scheduled_refresh,
  message
FROM refresh_status;
`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []storedRefreshStatus
	for rows.Next() {
		var row storedRefreshStatus
		if err := rows.Scan(
			&row.Status,
			&row.LatestSuccessfulRefresh,
			&row.LatestAttempt,
			&row.NextScheduledRefresh,
			&row.Message,
		); err != nil {
			return nil, err
		}
		result = append(result, row)
	}
	return result, rows.Err()
}

func convertLayerFamilies(rows []storedLayerFamily) []LayerFamily {
	layers := make([]LayerFamily, 0, len(rows))
	for _, row := range rows {
		layers = append(layers, LayerFamily{
			Key:            row.Key,
			Label:          row.Label,
			MartinLayerID:  row.MartinLayerID,
			Description:    row.Description,
			DefaultEnabled: row.DefaultEnabled,
		})
	}
	return layers
}

func (row storedBoundaryFeature) toBoundaryFeature() (BoundaryFeature, error) {
	geometry, err := parseJSONGeometry(row.LayerFamily, row.FeatureID, row.GeometryJSON)
	if err != nil {
		return BoundaryFeature{}, err
	}
	attributes, err := parseJSONAttributes(row.LayerFamily, row.FeatureID, row.AttributesJSON)
	if err != nil {
		return BoundaryFeature{}, err
	}
	lastSyncedAt, err := time.Parse(time.RFC3339, row.LastSyncedAt)
	if err != nil {
		return BoundaryFeature{}, fmt.Errorf("parse last_synced_at for %s/%s: %w", row.LayerFamily, row.FeatureID, err)
	}
	return BoundaryFeature{
		LayerFamily:    row.LayerFamily,
		FeatureID:      row.FeatureID,
		Title:          row.Title,
		Subtitle:       row.Subtitle,
		SourceID:       row.SourceID,
		GeometryLabel:  row.GeometryLabel,
		GeometrySource: row.GeometrySource,
		LastSyncedAt:   lastSyncedAt,
		Geometry:       geometry,
		Attributes:     attributes,
	}, nil
}

func (row storedRefreshStatus) toRefreshStatus() (RefreshStatus, error) {
	latestSuccessfulRefresh, err := time.Parse(time.RFC3339, row.LatestSuccessfulRefresh)
	if err != nil {
		return RefreshStatus{}, fmt.Errorf("parse latest_successful_refresh: %w", err)
	}
	latestAttempt, err := time.Parse(time.RFC3339, row.LatestAttempt)
	if err != nil {
		return RefreshStatus{}, fmt.Errorf("parse latest_attempt: %w", err)
	}
	nextScheduledRefresh, err := time.Parse(time.RFC3339, row.NextScheduledRefresh)
	if err != nil {
		return RefreshStatus{}, fmt.Errorf("parse next_scheduled_refresh: %w", err)
	}
	return RefreshStatus{
		Status:                  row.Status,
		LatestSuccessfulRefresh: latestSuccessfulRefresh,
		LatestAttempt:           latestAttempt,
		NextScheduledRefresh:    nextScheduledRefresh,
		Message:                 row.Message,
	}, nil
}

func parseJSONGeometry(layerFamily string, featureID string, raw string) (Geometry, error) {
	var geometry Geometry
	if err := json.Unmarshal([]byte(raw), &geometry); err != nil {
		return Geometry{}, fmt.Errorf("parse geometry for %s/%s: %w", layerFamily, featureID, err)
	}
	return geometry, nil
}

func parseJSONAttributes(layerFamily string, featureID string, raw string) (map[string]any, error) {
	var attributes map[string]any
	if err := json.Unmarshal([]byte(raw), &attributes); err != nil {
		return nil, fmt.Errorf("parse attributes for %s/%s: %w", layerFamily, featureID, err)
	}
	return attributes, nil
}

func findDuckDBExecutable() (string, error) {
	for _, envName := range []string{"DUCKDB_EXE", "DUCKDB_CLI_PATH"} {
		if candidate := strings.TrimSpace(os.Getenv(envName)); candidate != "" {
			if _, err := os.Stat(candidate); err == nil {
				return candidate, nil
			}
		}
	}
	if path, err := exec.LookPath("duckdb"); err == nil {
		return path, nil
	}
	return "", errors.New("duckdb executable not found on PATH")
}

func ensureNerisLayerFamily(layers []LayerFamily) []LayerFamily {
	for _, layer := range layers {
		if layer.Key == nerisLayerFamily.Key {
			return layers
		}
	}
	return append(layers, nerisLayerFamily)
}

func resolveWorkspacePath(repoRootAbs string, rawValue string) (string, bool) {
	candidate := strings.TrimSpace(rawValue)
	if candidate == "" {
		return "", false
	}
	path := candidate
	if !filepath.IsAbs(path) {
		path = filepath.Join(repoRootAbs, path)
	}
	resolvedPath, err := filepath.Abs(path)
	if err != nil {
		return "", false
	}
	relPath, err := filepath.Rel(repoRootAbs, resolvedPath)
	if err != nil {
		return "", false
	}
	if relPath == ".." || strings.HasPrefix(relPath, ".."+string(filepath.Separator)) {
		return "", false
	}
	if _, err := os.Stat(resolvedPath); err != nil {
		return "", false
	}
	return resolvedPath, true
}
