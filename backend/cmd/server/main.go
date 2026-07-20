package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"building-code-map/backend/internal/httpapi"
	"building-code-map/backend/internal/regulatory"
	"building-code-map/backend/internal/snapshot"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:8000", "listen address")
	cachePath := flag.String("cache", snapshot.DefaultCachePath(".."), "cache snapshot file")
	regulatoryPath := flag.String("regulatory-data", filepath.Join("data", "regulatory"), "directory containing compiled state jurisdiction profiles")
	corsOrigins := flag.String("cors-origins", "", "comma-separated CORS origins")
	flag.Parse()

	resolvedCachePath := *cachePath
	if strings.TrimSpace(resolvedCachePath) == snapshot.DefaultCachePath("..") {
		resolvedCachePath = snapshot.ResolveCachePath("..")
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

	options := httpapi.Options{
		AllowedOrigins:    httpapi.ParseAllowedOrigins(os.Getenv("BACKEND_CORS_ALLOWED_ORIGINS")),
		RegulatoryCatalog: catalog,
	}
	if trimmed := strings.TrimSpace(*corsOrigins); trimmed != "" {
		options.AllowedOrigins = httpapi.ParseAllowedOrigins(trimmed)
	}

	handler := httpapi.NewHandler(snap, options)
	slog.Info("starting Go backend", "addr", *addr, "cache", resolvedCachePath, "regulatory_profiles", catalog.Len())
	if err := http.ListenAndServe(*addr, handler); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
