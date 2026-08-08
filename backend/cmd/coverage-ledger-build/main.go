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
		fmt.Fprintln(stderr, "coverage-ledger-build requires generate or validate")
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

type options struct {
	inventoryRoot     string
	regimesRoot       string
	output            string
	states            string
	codeFamily        string
	projectType       string
	applicabilityDate string
}

func parseOptions(command string, args []string, stderr io.Writer) (options, int) {
	flags := flag.NewFlagSet(command, flag.ContinueOnError)
	flags.SetOutput(stderr)
	var result options
	flags.StringVar(&result.inventoryRoot, "inventory-root", "", "governmental entity inventory directory")
	flags.StringVar(&result.regimesRoot, "regimes-root", "", "state regime specification directory")
	flags.StringVar(&result.output, "output", "", "coverage ledger JSON path")
	flags.StringVar(&result.states, "states", "", "comma-separated state IDs")
	flags.StringVar(&result.codeFamily, "code-family", "", "code family to evaluate")
	flags.StringVar(&result.projectType, "project-type", "", "optional project type")
	flags.StringVar(&result.applicabilityDate, "applicability-date", "", "applicability date YYYY-MM-DD")
	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return options{}, exitSuccess
		}
		return options{}, exitArguments
	}
	if strings.TrimSpace(result.inventoryRoot) == "" || strings.TrimSpace(result.regimesRoot) == "" || strings.TrimSpace(result.output) == "" || strings.TrimSpace(result.states) == "" || strings.TrimSpace(result.codeFamily) == "" || strings.TrimSpace(result.applicabilityDate) == "" {
		fmt.Fprintln(stderr, "--inventory-root, --regimes-root, --output, --states, --code-family, and --applicability-date are required")
		return options{}, exitArguments
	}
	return result, -1
}

func buildRequest(options options) (regulatory.EntityCoverageLedgerRequest, error) {
	inventories, err := regulatory.LoadCensusEntityInventoryDirectory(options.inventoryRoot)
	if err != nil {
		return regulatory.EntityCoverageLedgerRequest{}, err
	}
	regimes, err := regulatory.LoadStateRegimeDirectory(options.regimesRoot)
	if err != nil {
		return regulatory.EntityCoverageLedgerRequest{}, err
	}
	stateIDs := make([]string, 0)
	for _, value := range strings.Split(options.states, ",") {
		if value = strings.TrimSpace(value); value != "" {
			stateIDs = append(stateIDs, value)
		}
	}
	return regulatory.EntityCoverageLedgerRequest{
		Inventories:       inventories,
		Regimes:           regimes,
		StateIDs:          stateIDs,
		CodeFamily:        options.codeFamily,
		ProjectType:       options.projectType,
		ApplicabilityDate: options.applicabilityDate,
	}, nil
}

func runGenerate(args []string, stdout, stderr io.Writer) int {
	options, status := parseOptions("generate", args, stderr)
	if status >= 0 {
		return status
	}
	request, err := buildRequest(options)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return exitFailure
	}
	ledger, err := regulatory.BuildEntityCoverageLedger(request)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return exitFailure
	}
	if err := regulatory.WriteEntityCoverageLedger(options.output, ledger); err != nil {
		fmt.Fprintln(stderr, err)
		return exitFailure
	}
	fmt.Fprintf(stdout, "generated coverage ledger for %d states and %d entities\n", len(ledger.States), ledger.Summary.ExpectedEntities)
	return exitSuccess
}

func runValidate(args []string, stdout, stderr io.Writer) int {
	options, status := parseOptions("validate", args, stderr)
	if status >= 0 {
		return status
	}
	request, err := buildRequest(options)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return exitFailure
	}
	if err := regulatory.ValidateEntityCoverageLedgerFile(request, options.output); err != nil {
		fmt.Fprintln(stderr, err)
		return exitFailure
	}
	ledger, err := regulatory.LoadEntityCoverageLedger(options.output)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return exitFailure
	}
	fmt.Fprintf(stdout, "validated coverage ledger for %d states and %d entities\n", len(ledger.States), ledger.Summary.ExpectedEntities)
	return exitSuccess
}
