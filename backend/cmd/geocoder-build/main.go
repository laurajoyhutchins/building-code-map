package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"building-code-map/backend/internal/geocoder"
	"building-code-map/backend/internal/snapshotmanifest"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stderr))
}

func run(args []string, stderr io.Writer) int {
	flags := flag.NewFlagSet("geocoder-build", flag.ContinueOnError)
	flags.SetOutput(stderr)
	output := flags.String("output", "", "output SQLite snapshot path")
	addressPoints := flags.String("address-points", "", "address-point CSV path")
	streetRanges := flags.String("street-ranges", "", "street-range CSV path")
	sourceName := flags.String("source-name", "", "source name recorded on every imported row")
	sourceVintage := flags.String("source-vintage", "", "source vintage recorded on every imported row")
	manifestTemplate := flags.String("manifest-template", "", "JSON manifest template with source, builder, record-count, and integrity metadata")
	if err := flags.Parse(args); err != nil {
		return 2
	}
	if strings.TrimSpace(*output) == "" || strings.TrimSpace(*sourceName) == "" || strings.TrimSpace(*sourceVintage) == "" || strings.TrimSpace(*manifestTemplate) == "" {
		fmt.Fprintln(stderr, "--output, --source-name, --source-vintage, and --manifest-template are required")
		return 2
	}
	if strings.TrimSpace(*addressPoints) == "" && strings.TrimSpace(*streetRanges) == "" {
		fmt.Fprintln(stderr, "at least one of --address-points or --street-ranges is required")
		return 2
	}
	if err := geocoder.BuildSnapshot(geocoder.BuildOptions{
		OutputPath:       *output,
		AddressPointsCSV: *addressPoints,
		StreetRangesCSV:  *streetRanges,
		SourceName:       *sourceName,
		SourceVintage:    *sourceVintage,
	}); err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	manifestBytes, err := os.ReadFile(*manifestTemplate)
	if err != nil {
		fmt.Fprintf(stderr, "read manifest template: %v\n", err)
		return 1
	}
	var manifest snapshotmanifest.Manifest
	if err := json.Unmarshal(manifestBytes, &manifest); err != nil {
		fmt.Fprintf(stderr, "decode manifest template: %v\n", err)
		return 1
	}
	if manifest.Kind != snapshotmanifest.KindGeocoder {
		fmt.Fprintf(stderr, "manifest template kind must be %q\n", snapshotmanifest.KindGeocoder)
		return 1
	}
	if _, err := snapshotmanifest.FinalizeAndWrite(*output, manifest); err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	return 0
}
