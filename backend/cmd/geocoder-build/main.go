package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"building-code-map/backend/geocoder"
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
	if err := flags.Parse(args); err != nil {
		return 2
	}
	if strings.TrimSpace(*output) == "" || strings.TrimSpace(*sourceName) == "" || strings.TrimSpace(*sourceVintage) == "" {
		fmt.Fprintln(stderr, "--output, --source-name, and --source-vintage are required")
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
	return 0
}
