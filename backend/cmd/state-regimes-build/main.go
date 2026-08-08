package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"building-code-map/backend/regulatory"
)

const (
	exitSuccess   = 0
	exitFailure   = 1
	exitArguments = 2
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(stderr, "state-regimes-build requires generate or validate")
		return exitArguments
	}
	switch args[0] {
	case "generate":
		return runGenerate(args[1:], stdout, stderr)
	case "validate":
		return runValidate(args[1:], stdout, stderr)
	default:
		fmt.Fprintf(stderr, "unknown command %q\n", args[0])
		return exitArguments
	}
}

func runGenerate(args []string, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("generate", flag.ContinueOnError)
	flags.SetOutput(stderr)
	profilesRoot := flags.String("profiles-root", "", "compiled regulatory profile directory")
	outputDir := flags.String("output-dir", "", "output directory for state regime specs")
	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return exitSuccess
		}
		return exitArguments
	}
	if strings.TrimSpace(*profilesRoot) == "" || strings.TrimSpace(*outputDir) == "" {
		fmt.Fprintln(stderr, "--profiles-root and --output-dir are required")
		return exitArguments
	}
	catalog, err := regulatory.LoadCatalog(*profilesRoot)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return exitFailure
	}
	specs, err := regulatory.BuildStateRegimeSpecs(catalog)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return exitFailure
	}
	if err := regulatory.WriteStateRegimeDirectory(*outputDir, specs); err != nil {
		fmt.Fprintln(stderr, err)
		return exitFailure
	}
	fmt.Fprintf(stdout, "generated %d state regime specs\n", len(specs))
	return exitSuccess
}

func runValidate(args []string, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("validate", flag.ContinueOnError)
	flags.SetOutput(stderr)
	profilesRoot := flags.String("profiles-root", "", "compiled regulatory profile directory")
	root := flags.String("root", "", "state regime spec directory")
	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return exitSuccess
		}
		return exitArguments
	}
	if strings.TrimSpace(*profilesRoot) == "" || strings.TrimSpace(*root) == "" {
		fmt.Fprintln(stderr, "--profiles-root and --root are required")
		return exitArguments
	}
	catalog, err := regulatory.LoadCatalog(*profilesRoot)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return exitFailure
	}
	if err := regulatory.ValidateStateRegimeDirectory(catalog, *root); err != nil {
		fmt.Fprintln(stderr, err)
		return exitFailure
	}
	fmt.Fprintf(stdout, "validated %d state regime specs\n", catalog.Len())
	return exitSuccess
}
