package snapshot

import (
	"database/sql"
	"os"
	"path/filepath"
	"testing"

	_ "modernc.org/sqlite"
)

func TestLoadFileLoadsSQLiteSnapshot(t *testing.T) {
	t.Setenv("DUCKDB_EXE", "")
	t.Setenv("DUCKDB_CLI_PATH", "")
	t.Setenv("PATH", "")

	cachePath := filepath.Join(t.TempDir(), "tigerweb.sqlite")
	buildSQLiteSnapshot(t, cachePath)

	snap, err := LoadFile(cachePath)
	if err != nil {
		t.Fatalf("LoadFile: %v", err)
	}

	if got, want := len(snap.LayerFamilies), 3; got != want {
		t.Fatalf("LayerFamilies length = %d, want %d", got, want)
	}
	if snap.LayerFamilies[0].Key != "counties" || snap.LayerFamilies[1].Key != "states" || snap.LayerFamilies[2].Key != "neris_jurisdictions" {
		t.Fatalf("unexpected layer order: %#v", snap.LayerFamilies)
	}
	if got, want := len(snap.BoundaryFeatures), 1; got != want {
		t.Fatalf("BoundaryFeatures length = %d, want %d", got, want)
	}
	if snap.RefreshStatus.Status != "ok" {
		t.Fatalf("RefreshStatus.Status = %q, want %q", snap.RefreshStatus.Status, "ok")
	}
	if snap.BoundaryFeatures[0].Geometry.Type != "Polygon" {
		t.Fatalf("Geometry.Type = %q, want %q", snap.BoundaryFeatures[0].Geometry.Type, "Polygon")
	}
	if got := snap.BoundaryFeatures[0].Attributes["STATEFP"]; got != "08" {
		t.Fatalf("Attributes[STATEFP] = %#v, want %q", got, "08")
	}
}

func TestResolveCachePathPrefersSQLiteCache(t *testing.T) {
	repoRoot := t.TempDir()
	sqlitePath := filepath.Join(repoRoot, "backend", "data", "tigerweb.sqlite")
	if err := os.MkdirAll(filepath.Dir(sqlitePath), 0o755); err != nil {
		t.Fatalf("mkdir sqlite path: %v", err)
	}
	if err := os.WriteFile(sqlitePath, []byte(""), 0o600); err != nil {
		t.Fatalf("write sqlite path: %v", err)
	}

	t.Setenv("TIGERWEB_CACHE_PATH", "")
	t.Setenv("TIGERWEB_HYDRATED_CACHE_PATH", "")

	got := ResolveCachePath(repoRoot)
	if got != sqlitePath {
		t.Fatalf("ResolveCachePath() = %q, want %q", got, sqlitePath)
	}
}

func TestResolveCachePathHonorsExplicitOverride(t *testing.T) {
	repoRoot := t.TempDir()
	cachePath := filepath.Join(repoRoot, "cache", "hydrated.duckdb")
	if err := os.MkdirAll(filepath.Dir(cachePath), 0o755); err != nil {
		t.Fatalf("mkdir cache path: %v", err)
	}
	if err := os.WriteFile(cachePath, []byte(""), 0o600); err != nil {
		t.Fatalf("write cache path: %v", err)
	}

	t.Setenv("TIGERWEB_CACHE_PATH", "")
	t.Setenv("TIGERWEB_HYDRATED_CACHE_PATH", filepath.Join("cache", "hydrated.duckdb"))

	got := ResolveCachePath(repoRoot)
	if got != cachePath {
		t.Fatalf("ResolveCachePath() = %q, want %q", got, cachePath)
	}
}

func TestResolveCachePathIgnoresOutsideWorkspaceOverride(t *testing.T) {
	repoRoot := t.TempDir()
	t.Setenv("TIGERWEB_CACHE_PATH", "")
	t.Setenv("TIGERWEB_HYDRATED_CACHE_PATH", "..\\global-cache\\tigerweb.duckdb")

	got := ResolveCachePath(repoRoot)
	want := DefaultCachePath(repoRoot)
	if got != want {
		t.Fatalf("ResolveCachePath() = %q, want %q", got, want)
	}
}

func buildSQLiteSnapshot(t *testing.T, cachePath string) {
	t.Helper()

	db, err := sql.Open("sqlite", cachePath)
	if err != nil {
		t.Fatalf("open sqlite snapshot: %v", err)
	}
	defer db.Close()

	statements := []string{
		`CREATE TABLE layer_families (
			key TEXT NOT NULL,
			label TEXT NOT NULL,
			martin_layer_id TEXT NOT NULL,
			description TEXT NOT NULL,
			default_enabled INTEGER NOT NULL
		);`,
		`CREATE TABLE boundary_features (
			layer_family TEXT NOT NULL,
			feature_id TEXT NOT NULL,
			title TEXT NOT NULL,
			subtitle TEXT NOT NULL,
			source_id TEXT NOT NULL,
			geometry_label TEXT NOT NULL,
			geometry_source TEXT NOT NULL,
			last_synced_at TEXT NOT NULL,
			geometry_json TEXT NOT NULL,
			attributes_json TEXT NOT NULL
		);`,
		`CREATE TABLE refresh_status (
			status TEXT NOT NULL,
			latest_successful_refresh TEXT NOT NULL,
			latest_attempt TEXT NOT NULL,
			next_scheduled_refresh TEXT NOT NULL,
			message TEXT NOT NULL
		);`,
		`INSERT INTO layer_families (key, label, martin_layer_id, description, default_enabled) VALUES
			('counties', 'Counties', 'tigerweb.counties', 'County boundaries used for jurisdiction lookups.', 0),
			('states', 'States', 'tigerweb.states', 'State boundaries mirrored from TIGERweb.', 1);`,
		`INSERT INTO boundary_features (
			layer_family, feature_id, title, subtitle, source_id, geometry_label, geometry_source, last_synced_at, geometry_json, attributes_json
		) VALUES (
			'states',
			'08',
			'Colorado',
			'State boundary',
			'GEOID=08',
			'Polygon',
			'tigerweb_live',
			'2026-06-22T12:00:00Z',
			'{"type":"Polygon","coordinates":[[[-109.06,37.0],[-102.04,37.0],[-102.04,41.0],[-109.06,41.0],[-109.06,37.0]]]}',
			'{"STATEFP":"08","NAME":"Colorado"}'
		);`,
		`INSERT INTO refresh_status (
			status, latest_successful_refresh, latest_attempt, next_scheduled_refresh, message
		) VALUES (
			'ok',
			'2026-06-22T09:40:00Z',
			'2026-06-22T09:40:00Z',
			'2026-06-23T09:40:00Z',
			'Cached TIGERweb snapshot is current.'
		);`,
	}

	for _, statement := range statements {
		if _, err := db.Exec(statement); err != nil {
			t.Fatalf("exec sqlite statement: %v", err)
		}
	}
}
