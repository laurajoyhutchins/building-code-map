package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"building-code-map/backend/bundle"
	"building-code-map/backend/engine"
	"building-code-map/backend/geocoder"
	"building-code-map/backend/regulatory"
	"building-code-map/backend/snapshot"
	"building-code-map/backend/transport/httpapi"
	"building-code-map/backend/transport/mcp"
)

const (
	exitSuccess = iota
	exitInternal
	exitArguments
	exitOutcome
	exitData
)

type runtime struct {
	manifest bundle.Verified
	snapshot snapshot.Snapshot
	catalog  regulatory.Catalog
	geocoder geocoder.Service
	database *geocoder.SQLiteService
	engine   engine.Engine
}

func main() { os.Exit(run(os.Args[1:])) }

func run(args []string) int {
	if len(args) == 0 {
		usage()
		return exitArguments
	}
	switch args[0] {
	case "resolve":
		return runResolve(args[1:])
	case "geocode":
		return runGeocode(args[1:])
	case "lookup":
		return runLookup(args[1:])
	case "inspect":
		return runInspect(args[1:])
	case "snapshot":
		return runSnapshot(args[1:])
	case "serve":
		return runServe(args[1:])
	case "--help", "-h", "help":
		usage()
		return exitSuccess
	default:
		logError("unknown command: " + args[0])
		return exitArguments
	}
}

func runResolve(args []string) int {
	flags := flag.NewFlagSet("resolve", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	address := flags.String("address", "", "civic address")
	pointValue := flags.String("point", "", "longitude,latitude")
	asOf := flags.String("as-of", "", "applicability date (YYYY-MM-DD)")
	bundlePath := flags.String("bundle", "bundle.json", "bundle manifest")
	pretty := flags.Bool("pretty", false, "indent JSON output")
	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return exitSuccess
		}
		return exitArguments
	}
	point, err := parsePoint(*pointValue)
	if err != nil {
		logError(err.Error())
		return exitArguments
	}
	if strings.TrimSpace(*asOf) == "" {
		logError("--as-of is required for resolution")
		return exitArguments
	}
	if (strings.TrimSpace(*address) == "") == (point == nil) {
		logError("provide exactly one of --address or --point")
		return exitArguments
	}
	runtime, err := loadRuntime(*bundlePath)
	if err != nil {
		logError(err.Error())
		return exitData
	}
	defer runtime.close()
	query := engine.Query{Address: *address, ApplicabilityDate: *asOf}
	if point != nil {
		query.Address = ""
		query.Point = point
	}
	result, err := runtime.engine.Resolve(context.Background(), query)
	if err != nil {
		writeEngineError(err)
		return outcomeExit(err)
	}
	writeJSON(result, *pretty)
	if result.Resolution.Status == "insufficient_evidence" || result.Resolution.Status == "conflicting" {
		return exitOutcome
	}
	return exitSuccess
}

func runGeocode(args []string) int {
	flags := flag.NewFlagSet("geocode", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	address := flags.String("address", "", "civic address")
	bundlePath := flags.String("bundle", "bundle.json", "bundle manifest")
	pretty := flags.Bool("pretty", false, "indent JSON output")
	if err := flags.Parse(args); err != nil || strings.TrimSpace(*address) == "" {
		logError("--address is required")
		return exitArguments
	}
	runtime, err := loadRuntime(*bundlePath)
	if err != nil {
		logError(err.Error())
		return exitData
	}
	defer runtime.close()
	result, err := runtime.engine.Geocode(context.Background(), *address)
	if err != nil {
		if result.Query != "" {
			writeJSON(result, *pretty)
		} else {
			writeEngineError(err)
		}
		return outcomeExit(err)
	}
	writeJSON(result, *pretty)
	return exitSuccess
}

func runLookup(args []string) int {
	flags := flag.NewFlagSet("lookup", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	pointValue := flags.String("point", "", "longitude,latitude")
	bundlePath := flags.String("bundle", "bundle.json", "bundle manifest")
	pretty := flags.Bool("pretty", false, "indent JSON output")
	if err := flags.Parse(args); err != nil {
		return exitArguments
	}
	point, err := parsePoint(*pointValue)
	if err != nil {
		logError(err.Error())
		return exitArguments
	}
	runtime, err := loadRuntime(*bundlePath)
	if err != nil {
		logError(err.Error())
		return exitData
	}
	defer runtime.close()
	result, err := runtime.engine.Lookup(context.Background(), *point)
	if err != nil {
		writeEngineError(err)
		return outcomeExit(err)
	}
	writeJSON(result, *pretty)
	return exitSuccess
}

func runInspect(args []string) int {
	if len(args) == 0 {
		logError("inspect requires bundle or jurisdiction")
		return exitArguments
	}
	switch args[0] {
	case "bundle":
		flags := flag.NewFlagSet("inspect bundle", flag.ContinueOnError)
		flags.SetOutput(os.Stderr)
		path := flags.String("bundle", "bundle.json", "bundle manifest")
		pretty := flags.Bool("pretty", false, "indent JSON output")
		if err := flags.Parse(args[1:]); err != nil {
			return exitArguments
		}
		verified, err := bundle.LoadAndVerify(*path)
		if err != nil {
			logError(err.Error())
			return exitData
		}
		writeJSON(map[string]any{"manifest": verified.Manifest, "bundle_manifest_digest": verified.Digest}, *pretty)
		return exitSuccess
	case "jurisdiction":
		flags := flag.NewFlagSet("inspect jurisdiction", flag.ContinueOnError)
		flags.SetOutput(os.Stderr)
		id := flags.String("id", "", "state ID or FIPS")
		path := flags.String("bundle", "bundle.json", "bundle manifest")
		pretty := flags.Bool("pretty", false, "indent JSON output")
		if err := flags.Parse(args[1:]); err != nil || strings.TrimSpace(*id) == "" {
			logError("--id is required")
			return exitArguments
		}
		runtime, err := loadRuntime(*path)
		if err != nil {
			logError(err.Error())
			return exitData
		}
		defer runtime.close()
		profile, ok := runtime.catalog.Profile(*id, *id)
		if !ok {
			logError("jurisdiction profile not found")
			return exitOutcome
		}
		writeJSON(profile, *pretty)
		return exitSuccess
	default:
		logError("unknown inspect target: " + args[0])
		return exitArguments
	}
}

func runServe(args []string) int {
	flags := flag.NewFlagSet("serve", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	address := flags.String("http", "127.0.0.1:8000", "HTTP listen address")
	bundlePath := flags.String("bundle", "bundle.json", "bundle manifest")
	mcpStdio := flags.Bool("mcp-stdio", false, "serve MCP over stdio")
	if err := flags.Parse(args); err != nil {
		return exitArguments
	}
	runtime, err := loadRuntime(*bundlePath)
	if err != nil {
		logError(err.Error())
		return exitData
	}
	if *mcpStdio {
		defer runtime.close()
		if err := mcp.Serve(context.Background(), os.Stdin, os.Stdout, os.Stderr, runtime.engine); err != nil {
			if errors.Is(err, context.Canceled) {
				return exitSuccess
			}
			logError(err.Error())
			return exitInternal
		}
		return exitSuccess
	}
	options := httpapi.Options{
		Snapshot: snapshotFromRuntime(runtime), RegulatoryCatalog: runtime.catalog, Geocoder: runtime.geocoder,
		BundleIdentity: runtime.engine.BundleIdentity(context.Background()),
	}
	if err := http.ListenAndServe(*address, httpapi.NewHandler(runtime.engine, options)); err != nil {
		runtime.close()
		logError(err.Error())
		return exitInternal
	}
	return exitSuccess
}

func loadRuntime(path string) (*runtime, error) {
	verified, err := bundle.LoadAndVerify(path)
	if err != nil {
		return nil, err
	}
	root := filepath.Dir(filepath.Clean(path))
	boundary := verified.Manifest.Components["boundary_snapshot"]
	boundaryPath := filepath.Join(root, filepath.FromSlash(boundary.Path))
	snap, err := loadSnapshot(boundaryPath)
	if err != nil {
		return nil, fmt.Errorf("load boundary snapshot: %w", err)
	}
	regulatoryComponent := verified.Manifest.Components["regulatory_catalog"]
	regulatoryPath := filepath.Join(root, filepath.FromSlash(regulatoryComponent.Path))
	catalog, err := regulatory.LoadCatalog(regulatoryCatalogRoot(regulatoryComponent, regulatoryPath))
	if err != nil {
		return nil, fmt.Errorf("load regulatory catalog: %w", err)
	}
	var service geocoder.Service
	var database *geocoder.SQLiteService
	if component, ok := verified.Manifest.Components["geocoder"]; ok {
		geocoderPath := filepath.Join(root, filepath.FromSlash(component.Path))
		database, err = geocoder.Open(geocoderPath)
		if err != nil {
			return nil, fmt.Errorf("load geocoder snapshot: %w", err)
		}
		service = database
	}
	authority, err := engine.New(engine.Config{
		Snapshot: snap, RegulatoryCatalog: catalog, Geocoder: service,
		Clock: fixedBundleClock(verified.Manifest.Coverage.AsOf),
		BundleIdentity: engine.BundleIdentity{
			SourceCommit: verified.Manifest.SourceCommit, EngineVersion: verified.Manifest.EngineVersion,
			BundleManifestDigest:   verified.Digest,
			BoundarySnapshotDigest: boundary.SHA256, RegulatoryCatalogDigest: regulatoryComponent.SHA256,
		},
	})
	if err != nil {
		if database != nil {
			database.Close()
		}
		return nil, err
	}
	return &runtime{manifest: verified, snapshot: snap, catalog: catalog, geocoder: service, database: database, engine: authority}, nil
}

func regulatoryCatalogRoot(component bundle.Component, path string) string {
	var shape struct {
		Recursive bool `json:"recursive"`
	}
	if raw, err := json.Marshal(component); err == nil {
		_ = json.Unmarshal(raw, &shape)
	}
	if shape.Recursive {
		return path
	}
	return filepath.Dir(path)
}

func fixedBundleClock(asOf string) engine.Clock {
	value, err := time.Parse(time.DateOnly, asOf)
	if err != nil {
		return engine.RealClock{}
	}
	return engine.NewFixedClock(value.UTC())
}

func loadSnapshot(path string) (snapshot.Snapshot, error) {
	if strings.EqualFold(filepath.Ext(path), ".json") {
		raw, err := os.ReadFile(path)
		if err != nil {
			return snapshot.Snapshot{}, err
		}
		var snap snapshot.Snapshot
		if err := json.Unmarshal(raw, &snap); err != nil {
			return snapshot.Snapshot{}, err
		}
		return snap, nil
	}
	return snapshot.LoadFile(path)
}

func snapshotFromRuntime(value *runtime) snapshot.Snapshot { return value.snapshot }

func (value *runtime) close() {
	if value.database != nil {
		_ = value.database.Close()
	}
}

func parsePoint(raw string) (*engine.Point, error) {
	if strings.TrimSpace(raw) == "" {
		return nil, nil
	}
	parts := strings.Split(raw, ",")
	if len(parts) != 2 {
		return nil, errors.New("--point must use longitude,latitude")
	}
	longitude, err := strconv.ParseFloat(strings.TrimSpace(parts[0]), 64)
	if err != nil {
		return nil, errors.New("--point longitude is invalid")
	}
	latitude, err := strconv.ParseFloat(strings.TrimSpace(parts[1]), 64)
	if err != nil {
		return nil, errors.New("--point latitude is invalid")
	}
	return &engine.Point{Longitude: longitude, Latitude: latitude}, nil
}

func outcomeExit(err error) int {
	var engineErr engine.EngineError
	if errors.As(err, &engineErr) {
		switch engineErr.Code {
		case engine.ErrorDataBundleInvalid, engine.ErrorRegulatoryCatalogUnavailable:
			return exitData
		case engine.ErrorInternal:
			return exitInternal
		default:
			return exitOutcome
		}
	}
	return exitInternal
}

func writeEngineError(err error) {
	var engineErr engine.EngineError
	if !errors.As(err, &engineErr) {
		engineErr = engine.EngineError{Code: engine.ErrorInternal, Message: err.Error()}
	}
	writeJSON(engineErr, false)
}

func writeJSON(value any, pretty bool) {
	var raw []byte
	var err error
	if pretty {
		raw, err = json.MarshalIndent(value, "", "  ")
	} else {
		raw, err = json.Marshal(value)
	}
	if err != nil {
		logError(err.Error())
		return
	}
	_, _ = os.Stdout.Write(append(raw, '\n'))
}

func logError(message string) { _, _ = fmt.Fprintln(os.Stderr, message) }

func usage() {
	logError("bcm resolve|geocode|lookup|inspect|snapshot|serve")
}
