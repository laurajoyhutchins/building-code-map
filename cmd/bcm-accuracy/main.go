// Command bcm-accuracy evaluates project-code-basis outputs against an
// independently curated evidence-backed regulatory truth set.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/laurajoyhutchins/building-code-map/accuracy"
	"github.com/laurajoyhutchins/building-code-map/engine"
)

type truthFile struct {
	SchemaVersion string          `json:"schema_version"`
	Cases         []accuracy.Case `json:"cases"`
}

type resultsFile struct {
	Results []engine.ProjectCodeBasis `json:"results"`
}

func readJSON(path string, target any) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(target); err != nil {
		return err
	}
	var trailing any
	if err := decoder.Decode(&trailing); err != io.EOF {
		return fmt.Errorf("unexpected trailing JSON in %s", path)
	}
	return nil
}

func run(args []string, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("bcm-accuracy", flag.ContinueOnError)
	flags.SetOutput(stderr)
	truthPath := flags.String("truth", "", "path to independent reviewed cases (JSON)")
	resultsPath := flags.String("results", "", "path to project-code-basis results (JSON)")
	strict := flags.Bool("strict", false, "also fail when any reviewed case is missing or abstained")
	if err := flags.Parse(args); err != nil {
		return 2
	}
	if *truthPath == "" || *resultsPath == "" {
		fmt.Fprintln(stderr, "both -truth and -results are required")
		return 2
	}
	var truth truthFile
	if err := readJSON(*truthPath, &truth); err != nil {
		fmt.Fprintf(stderr, "read reviewed cases: %v\n", err)
		return 2
	}
	if truth.SchemaVersion != "1" {
		fmt.Fprintln(stderr, "unsupported accuracy truth schema_version")
		return 2
	}
	var results resultsFile
	if err := readJSON(*resultsPath, &results); err != nil {
		fmt.Fprintf(stderr, "read observed project results: %v\n", err)
		return 2
	}
	report, err := accuracy.Assess(truth.Cases, results.Results)
	if err != nil {
		fmt.Fprintf(stderr, "invalid accuracy assessment: %v\n", err)
		return 2
	}
	encoder := json.NewEncoder(stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(report); err != nil {
		fmt.Fprintf(stderr, "write accuracy report: %v\n", err)
		return 2
	}
	if report.Summary.IncorrectAssertions != 0 || report.Summary.FalseVerified != 0 {
		return 1
	}
	if *strict && (report.Summary.Abstentions != 0 || report.Summary.MissingResults != 0) {
		return 1
	}
	return 0
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}
