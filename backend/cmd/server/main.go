package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strings"

	"building-code-map/backend/internal/httpapi"
	"building-code-map/backend/internal/snapshot"
)

func main() {
	addr := flag.String("addr", "127.0.0.1:8000", "listen address")
	cachePath := flag.String("cache", snapshot.DefaultCachePath(".."), "cache snapshot file")
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

	options := httpapi.Options{AllowedOrigins: httpapi.ParseAllowedOrigins(os.Getenv("BACKEND_CORS_ALLOWED_ORIGINS"))}
	if trimmed := strings.TrimSpace(*corsOrigins); trimmed != "" {
		options.AllowedOrigins = httpapi.ParseAllowedOrigins(trimmed)
	}

	handler := httpapi.NewHandler(snap, options)
	slog.Info("starting Go backend", "addr", *addr, "cache", resolvedCachePath)
	if err := http.ListenAndServe(*addr, handler); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
