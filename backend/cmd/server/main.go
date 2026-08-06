package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"building-code-map/backend/internal/geocoder"
	"building-code-map/backend/internal/httpapi"
	"building-code-map/backend/internal/regulatory"
	"building-code-map/backend/internal/snapshot"
	"building-code-map/backend/internal/snapshotmanifest"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:8000", "listen address")
	cachePath := flag.String("cache", snapshot.DefaultCachePath(".."), "cache snapshot file")
	regulatoryPath := flag.String("regulatory-data", filepath.Join("data", "regulatory"), "directory containing compiled state jurisdiction profiles")
	geocoderPath := flag.String("geocoder-data", filepath.Join("data", "geocoder.sqlite"), "optional local geocoder SQLite snapshot")
	corsOrigins := flag.String("cors-origins", "", "comma-separated CORS origins")
	flag.Parse()

	resolvedCachePath := *cachePath
	if strings.TrimSpace(resolvedCachePath) == snapshot.DefaultCachePath("..") {
		resolvedCachePath = snapshot.ResolveCachePath("..")
	}

	boundaryVerification, err := snapshotmanifest.LoadAndVerify(resolvedCachePath, snapshotmanifest.KindBoundary)
	if err != nil {
		slog.Error("verify boundary snapshot manifest", "path", resolvedCachePath, "error", err)
		os.Exit(1)
	}
	snap, err := snapshot.LoadFile(resolvedCachePath)
	if err != nil {
		slog.Error("load snapshot", "path", resolvedCachePath, "error", err)
		os.Exit(1)
	}

	catalog, err := regulatory.LoadCatalog(*regulatoryPath)
	if err != nil {
		slog.Warn("load regulatory profiles", "path", *regulatoryPath, "error", err)
		catalog = regulatory.EmptyCatalog()
	}

	configuredGeocoderPath := *geocoderPath
	if environmentPath := strings.TrimSpace(os.Getenv("GEOCODER_DATA_PATH")); environmentPath != "" {
		configuredGeocoderPath = environmentPath
	}
	resolvedGeocoderPath, err := geocoder.ResolveDataPath(".", configuredGeocoderPath)
	if err != nil {
		slog.Warn("resolve geocoder snapshot", "error", err)
	}
	var geocoderService geocoder.Service
	var geocoderDatabase *geocoder.SQLiteService
	var geocoderSnapshotID string
	if err == nil {
		if _, statErr := os.Stat(resolvedGeocoderPath); statErr == nil {
			verification, verifyErr := snapshotmanifest.LoadAndVerify(resolvedGeocoderPath, snapshotmanifest.KindGeocoder)
			if verifyErr != nil {
				slog.Warn("verify geocoder snapshot manifest", "path", resolvedGeocoderPath, "error", verifyErr)
			} else {
				geocoderDatabase, err = geocoder.Open(resolvedGeocoderPath)
				if err != nil {
					slog.Warn("load geocoder snapshot", "path", resolvedGeocoderPath, "error", err)
				} else {
					geocoderService = geocoderDatabase
					geocoderSnapshotID = verification.Manifest.SnapshotID
					defer geocoderDatabase.Close()
				}
			}
		} else if !os.IsNotExist(statErr) {
			slog.Warn("inspect geocoder snapshot", "path", resolvedGeocoderPath, "error", statErr)
		}
	}

	options := httpapi.Options{
		AllowedOrigins:     httpapi.ParseAllowedOrigins(os.Getenv("BACKEND_CORS_ALLOWED_ORIGINS")),
		RegulatoryCatalog:  catalog,
		Geocoder:           geocoderService,
		BoundarySnapshotID: boundaryVerification.Manifest.SnapshotID,
		GeocoderSnapshotID: geocoderSnapshotID,
	}
	if trimmed := strings.TrimSpace(*corsOrigins); trimmed != "" {
		options.AllowedOrigins = httpapi.ParseAllowedOrigins(trimmed)
	}

	handler := httpapi.NewHandler(snap, options)
	slog.Info(
		"starting Go backend",
		"addr", *addr,
		"cache", resolvedCachePath,
		"boundary_snapshot_id", boundaryVerification.Manifest.SnapshotID,
		"geocoder_snapshot_id", geocoderSnapshotID,
		"regulatory_profiles", catalog.Len(),
		"geocoder_available", geocoderService != nil,
	)
	if err := http.ListenAndServe(*addr, handler); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
