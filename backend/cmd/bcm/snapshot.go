package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"building-code-map/backend/geocoder"
	"building-code-map/backend/snapshot"
	"building-code-map/backend/snapshotbuild"
)

func runSnapshot(args []string) int {
	if len(args) == 0 {
		logError("snapshot requires build or audit")
		return exitArguments
	}
	switch args[0] {
	case "build":
		return runSnapshotBuild(args[1:])
	case "audit":
		return runSnapshotAudit(args[1:])
	default:
		logError("unknown snapshot command: " + args[0])
		return exitArguments
	}
}

type snapshotFlags struct {
	kind                   *string
	duckdb                 *string
	duckdbVersion          *string
	duckdbSHA256           *string
	spatialExtension       *string
	spatialExtensionSHA256 *string
	workDir                *string
	output                 *string
	receipt                *string
	builderVersion         *string
	builderRevision        *string
	builtAt                *string
	outputCRS              *string
	addressPoints          *string
	streetRanges           *string
	layerFamilies          *string
	boundaryFeatures       *string
	refreshStatus          *string
	publisher              *string
	locator                *string
	product                *string
	vintage                *string
	retrievedAt            *string
	licenseReviewStatus    *string
	redistributionStatus   *string
	inputCRS               *string
	transformation         *string
	pretty                 *bool
}

func addSnapshotFlags(flags *flag.FlagSet, build bool) snapshotFlags {
	values := snapshotFlags{
		kind:                   flags.String("kind", "", "snapshot kind: boundary or geocoder"),
		duckdb:                 flags.String("duckdb", "", "absolute path to the explicitly provisioned DuckDB executable"),
		duckdbVersion:          flags.String("duckdb-version", "", "expected DuckDB version"),
		duckdbSHA256:           flags.String("duckdb-sha256", "", "expected DuckDB executable SHA-256"),
		spatialExtension:       flags.String("spatial-extension", "", "absolute path to an explicitly provisioned DuckDB spatial extension"),
		spatialExtensionSHA256: flags.String("spatial-extension-sha256", "", "expected DuckDB spatial extension SHA-256"),
		workDir:                flags.String("work-dir", "", "directory for disposable DuckDB build state"),
		builderVersion:         flags.String("builder-version", "", "bcm builder version"),
		builderRevision:        flags.String("builder-revision", "", "source revision used for the build"),
		builtAt:                flags.String("built-at", "", "deterministic RFC3339 build timestamp"),
		outputCRS:              flags.String("output-crs", "EPSG:4326", "output coordinate reference system"),
		addressPoints:          flags.String("address-points", "", "geocoder address-point CSV source"),
		streetRanges:           flags.String("street-ranges", "", "geocoder street-range CSV source"),
		layerFamilies:          flags.String("layer-families", "", "boundary layer-family CSV source"),
		boundaryFeatures:       flags.String("boundary-features", "", "boundary feature CSV source"),
		refreshStatus:          flags.String("refresh-status", "", "boundary refresh-status CSV source"),
		publisher:              flags.String("publisher", "", "source publisher"),
		locator:                flags.String("locator", "", "source acquisition locator"),
		product:                flags.String("product", "", "source product"),
		vintage:                flags.String("vintage", "", "source vintage"),
		retrievedAt:            flags.String("retrieved-at", "", "source retrieval timestamp in RFC3339"),
		licenseReviewStatus:    flags.String("license-review-status", "", "source license review status"),
		redistributionStatus:   flags.String("redistribution-status", "", "source redistribution status"),
		inputCRS:               flags.String("input-crs", "", "source coordinate reference system"),
		transformation:         flags.String("transformation", "", "source transformation identity"),
		pretty:                 flags.Bool("pretty", false, "indent JSON output"),
	}
	if build {
		values.output = flags.String("output", "", "canonical output SQLite path")
		values.receipt = flags.String("receipt", "", "build receipt path; defaults beside the output")
	}
	return values
}

func runSnapshotAudit(args []string) int {
	flags := flag.NewFlagSet("snapshot audit", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	values := addSnapshotFlags(flags, false)
	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return exitSuccess
		}
		return exitArguments
	}
	request, err := values.request(false)
	if err != nil {
		logError(err.Error())
		return exitArguments
	}
	report, receipt, err := (snapshotbuild.Pipeline{}).Audit(context.Background(), request)
	if err != nil {
		logError(err.Error())
		return exitData
	}
	writeJSON(map[string]any{"report": report, "receipt": receipt}, *values.pretty)
	return exitSuccess
}

func runSnapshotBuild(args []string) int {
	flags := flag.NewFlagSet("snapshot build", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	values := addSnapshotFlags(flags, true)
	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return exitSuccess
		}
		return exitArguments
	}
	request, err := values.request(true)
	if err != nil {
		logError(err.Error())
		return exitArguments
	}
	result, err := (snapshotbuild.Pipeline{}).Build(context.Background(), request, finalizer(request, values))
	if err != nil {
		logError(err.Error())
		return exitData
	}
	receiptPath := strings.TrimSpace(*values.receipt)
	if receiptPath == "" {
		receiptPath = request.OutputPath + ".build.json"
	}
	if err := snapshotbuild.WriteReceipt(receiptPath, result.Receipt); err != nil {
		logError("write build receipt: " + err.Error())
		return exitData
	}
	writeJSON(map[string]any{
		"manifest": result.Manifest,
		"receipt":  result.Receipt,
		"report":   result.Report,
	}, *values.pretty)
	return exitSuccess
}

func (values snapshotFlags) request(build bool) (snapshotbuild.Request, error) {
	kind := snapshotbuild.Kind(strings.ToLower(strings.TrimSpace(*values.kind)))
	if kind != snapshotbuild.KindBoundary && kind != snapshotbuild.KindGeocoder {
		return snapshotbuild.Request{}, errors.New("--kind must be boundary or geocoder")
	}
	for label, value := range map[string]string{
		"--duckdb":                *values.duckdb,
		"--duckdb-version":        *values.duckdbVersion,
		"--duckdb-sha256":         *values.duckdbSHA256,
		"--work-dir":              *values.workDir,
		"--builder-version":       *values.builderVersion,
		"--builder-revision":      *values.builderRevision,
		"--built-at":              *values.builtAt,
		"--publisher":             *values.publisher,
		"--locator":               *values.locator,
		"--product":               *values.product,
		"--vintage":               *values.vintage,
		"--retrieved-at":          *values.retrievedAt,
		"--license-review-status": *values.licenseReviewStatus,
		"--redistribution-status": *values.redistributionStatus,
	} {
		if strings.TrimSpace(value) == "" {
			return snapshotbuild.Request{}, fmt.Errorf("%s is required", label)
		}
	}
	if build && strings.TrimSpace(*values.output) == "" {
		return snapshotbuild.Request{}, errors.New("--output is required")
	}
	if build && kind == snapshotbuild.KindBoundary && strings.TrimSpace(*values.spatialExtension) == "" {
		return snapshotbuild.Request{}, errors.New("boundary build requires --spatial-extension and --spatial-extension-sha256")
	}
	if strings.TrimSpace(*values.spatialExtension) != "" && strings.TrimSpace(*values.spatialExtensionSHA256) == "" {
		return snapshotbuild.Request{}, errors.New("--spatial-extension-sha256 is required with --spatial-extension")
	}
	builtAt, err := time.Parse(time.RFC3339, *values.builtAt)
	if err != nil {
		return snapshotbuild.Request{}, fmt.Errorf("--built-at must be RFC3339: %w", err)
	}
	if _, err := time.Parse(time.RFC3339, *values.retrievedAt); err != nil {
		return snapshotbuild.Request{}, fmt.Errorf("--retrieved-at must be RFC3339: %w", err)
	}
	duckdbPath, err := filepath.Abs(*values.duckdb)
	if err != nil {
		return snapshotbuild.Request{}, err
	}
	workDir, err := filepath.Abs(*values.workDir)
	if err != nil {
		return snapshotbuild.Request{}, err
	}
	var extensionPath string
	if strings.TrimSpace(*values.spatialExtension) != "" {
		extensionPath, err = filepath.Abs(*values.spatialExtension)
		if err != nil {
			return snapshotbuild.Request{}, err
		}
	}
	sources, err := values.sources(kind)
	if err != nil {
		return snapshotbuild.Request{}, err
	}
	return snapshotbuild.Request{
		Kind:                           kind,
		DuckDBPath:                     duckdbPath,
		ExpectedDuckDBVersion:          strings.TrimSpace(*values.duckdbVersion),
		ExpectedDuckDBSHA256:           strings.TrimSpace(*values.duckdbSHA256),
		SpatialExtensionPath:           extensionPath,
		ExpectedSpatialExtensionSHA256: strings.TrimSpace(*values.spatialExtensionSHA256),
		WorkingDirectory:               workDir,
		OutputPath:                     strings.TrimSpace(valueOrEmpty(values.output)),
		BuilderVersion:                 strings.TrimSpace(*values.builderVersion),
		BuilderRevision:                strings.TrimSpace(*values.builderRevision),
		BuiltAt:                        builtAt,
		OutputCRS:                      strings.TrimSpace(*values.outputCRS),
		Sources:                        sources,
	}, nil
}

func (values snapshotFlags) sources(kind snapshotbuild.Kind) ([]snapshotbuild.Source, error) {
	paths := map[string]string{}
	switch kind {
	case snapshotbuild.KindGeocoder:
		if strings.TrimSpace(*values.addressPoints) != "" {
			paths["address_points"] = *values.addressPoints
		}
		if strings.TrimSpace(*values.streetRanges) != "" {
			paths["street_ranges"] = *values.streetRanges
		}
		if len(paths) == 0 {
			return nil, errors.New("geocoder requires --address-points or --street-ranges")
		}
	case snapshotbuild.KindBoundary:
		paths["layer_families"] = *values.layerFamilies
		paths["boundary_features"] = *values.boundaryFeatures
		paths["refresh_status"] = *values.refreshStatus
		for name, path := range paths {
			if strings.TrimSpace(path) == "" {
				return nil, fmt.Errorf("boundary requires --%s", strings.ReplaceAll(name, "_", "-"))
			}
		}
	}
	names := make([]string, 0, len(paths))
	for name := range paths {
		names = append(names, name)
	}
	sortStrings(names)
	sources := make([]snapshotbuild.Source, 0, len(names))
	for _, name := range names {
		absolute, err := filepath.Abs(paths[name])
		if err != nil {
			return nil, err
		}
		sources = append(sources, snapshotbuild.Source{
			Name:                 name,
			Path:                 absolute,
			Locator:              strings.TrimSpace(*values.locator) + "#" + name,
			Publisher:            strings.TrimSpace(*values.publisher),
			Product:              strings.TrimSpace(*values.product),
			Layer:                name,
			Vintage:              strings.TrimSpace(*values.vintage),
			RetrievedAt:          strings.TrimSpace(*values.retrievedAt),
			LicenseReviewStatus:  strings.TrimSpace(*values.licenseReviewStatus),
			RedistributionStatus: strings.TrimSpace(*values.redistributionStatus),
			InputCRS:             strings.TrimSpace(*values.inputCRS),
			Transformation:       strings.TrimSpace(*values.transformation),
		})
	}
	return sources, nil
}

func finalizer(request snapshotbuild.Request, values snapshotFlags) snapshotbuild.FinalizeFunc {
	return func(_ context.Context, normalized map[string]string, output string) error {
		switch request.Kind {
		case snapshotbuild.KindGeocoder:
			return geocoder.BuildSnapshot(geocoder.BuildOptions{
				OutputPath:       output,
				AddressPointsCSV: normalized["address_points"],
				StreetRangesCSV:  normalized["street_ranges"],
				SourceName:       strings.TrimSpace(*values.product),
				SourceVintage:    strings.TrimSpace(*values.vintage),
			})
		case snapshotbuild.KindBoundary:
			return snapshot.BuildSQLite(snapshot.SQLiteBuildOptions{
				OutputPath:          output,
				LayerFamiliesCSV:    normalized["layer_families"],
				BoundaryFeaturesCSV: normalized["boundary_features"],
				RefreshStatusCSV:    normalized["refresh_status"],
			})
		default:
			return fmt.Errorf("unsupported snapshot kind %q", request.Kind)
		}
	}
}

func valueOrEmpty(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}

func sortStrings(values []string) {
	for i := 1; i < len(values); i++ {
		for j := i; j > 0 && values[j] < values[j-1]; j-- {
			values[j], values[j-1] = values[j-1], values[j]
		}
	}
}
