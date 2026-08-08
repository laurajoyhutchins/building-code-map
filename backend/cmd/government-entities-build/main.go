package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"time"

	"building-code-map/backend/regulatory"
)

const (
	exitSuccess   = 0
	exitFailure   = 1
	exitArguments = 2
)

type sourceArtifact struct {
	Path   string `json:"path"`
	SHA256 string `json:"sha256"`
}

type generalPurposeContract struct {
	Sheet       string   `json:"sheet"`
	RecordCount int      `json:"record_count"`
	UnitTypes   []string `json:"unit_types"`
	Columns     []string `json:"columns"`
}

type scopeContract struct {
	StatesAndDC int    `json:"states_and_dc"`
	Note        string `json:"note"`
}

type censusSourceContract struct {
	SchemaVersion     string                 `json:"schema_version"`
	SourceID          string                 `json:"source_id"`
	Publisher         string                 `json:"publisher"`
	Product           string                 `json:"product"`
	ReleaseDate       string                 `json:"release_date"`
	AccessedAt        string                 `json:"accessed_at"`
	LandingPage       string                 `json:"landing_page"`
	ArchiveURL        string                 `json:"archive_url"`
	ArchiveSHA256     string                 `json:"archive_sha256"`
	Workbook          sourceArtifact         `json:"workbook"`
	Documentation     sourceArtifact         `json:"documentation"`
	GeneralPurpose    generalPurposeContract `json:"general_purpose"`
	Scope             scopeContract          `json:"scope"`
	AuthorityBoundary string                 `json:"authority_boundary"`
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintln(stderr, "government-entities-build requires generate or validate")
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
	input := flags.String("input", "", "normalized General Purpose CSV")
	outputDir := flags.String("output-dir", "", "output directory for per-state inventories")
	sourceContractPath := flags.String("source-contract", "", "pinned Census source contract JSON")
	generatedAtValue := flags.String("generated-at", "", "RFC3339 generation timestamp")
	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return exitSuccess
		}
		return exitArguments
	}
	if strings.TrimSpace(*input) == "" || strings.TrimSpace(*outputDir) == "" || strings.TrimSpace(*sourceContractPath) == "" {
		fmt.Fprintln(stderr, "--input, --output-dir, and --source-contract are required")
		return exitArguments
	}
	generatedAt, err := time.Parse(time.RFC3339, strings.TrimSpace(*generatedAtValue))
	if err != nil {
		fmt.Fprintf(stderr, "--generated-at must be RFC3339: %v\n", err)
		return exitArguments
	}
	contract, err := loadSourceContract(*sourceContractPath)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return exitFailure
	}
	file, err := os.Open(*input)
	if err != nil {
		fmt.Fprintf(stderr, "open normalized Census input: %v\n", err)
		return exitFailure
	}
	defer file.Close()
	build, err := regulatory.BuildCensusEntityInventories(file, contract.source(), generatedAt)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return exitFailure
	}
	if err := regulatory.ValidateNationalCensusInventory(build, contract.GeneralPurpose.RecordCount); err != nil {
		fmt.Fprintln(stderr, err)
		return exitFailure
	}
	if err := regulatory.WriteCensusEntityInventoryDirectory(*outputDir, build); err != nil {
		fmt.Fprintln(stderr, err)
		return exitFailure
	}
	fmt.Fprintf(stdout, "generated %d governmental entities across %d state/DC inventories\n", build.Index.TotalEntities, len(build.Inventories))
	return exitSuccess
}

func runValidate(args []string, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("validate", flag.ContinueOnError)
	flags.SetOutput(stderr)
	root := flags.String("root", "", "governmental entity inventory directory")
	sourceContractPath := flags.String("source-contract", "", "pinned Census source contract JSON")
	if err := flags.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return exitSuccess
		}
		return exitArguments
	}
	if strings.TrimSpace(*root) == "" || strings.TrimSpace(*sourceContractPath) == "" {
		fmt.Fprintln(stderr, "--root and --source-contract are required")
		return exitArguments
	}
	contract, err := loadSourceContract(*sourceContractPath)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return exitFailure
	}
	build, err := regulatory.LoadCensusEntityInventoryDirectory(*root)
	if err != nil {
		fmt.Fprintln(stderr, err)
		return exitFailure
	}
	if build.Index.SourceID != contract.SourceID {
		fmt.Fprintf(stderr, "inventory source_id %q does not match contract %q\n", build.Index.SourceID, contract.SourceID)
		return exitFailure
	}
	if err := regulatory.ValidateNationalCensusInventory(build, contract.GeneralPurpose.RecordCount); err != nil {
		fmt.Fprintln(stderr, err)
		return exitFailure
	}
	fmt.Fprintf(stdout, "validated %d governmental entities across %d state/DC inventories\n", build.Index.TotalEntities, len(build.Inventories))
	return exitSuccess
}

func loadSourceContract(path string) (censusSourceContract, error) {
	file, err := os.Open(path)
	if err != nil {
		return censusSourceContract{}, fmt.Errorf("open Census source contract: %w", err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	var contract censusSourceContract
	if err := decoder.Decode(&contract); err != nil {
		return censusSourceContract{}, fmt.Errorf("decode Census source contract: %w", err)
	}
	if err := contract.validate(); err != nil {
		return censusSourceContract{}, err
	}
	return contract, nil
}

func (contract censusSourceContract) validate() error {
	if contract.SchemaVersion != "1.0" {
		return fmt.Errorf("unsupported Census source contract schema_version %q", contract.SchemaVersion)
	}
	if contract.SourceID != "src:census-government-units-2025" {
		return fmt.Errorf("unexpected Census source_id %q", contract.SourceID)
	}
	if strings.TrimSpace(contract.Publisher) == "" || strings.TrimSpace(contract.Product) == "" || strings.TrimSpace(contract.LandingPage) == "" || strings.TrimSpace(contract.AuthorityBoundary) == "" {
		return fmt.Errorf("Census source contract publisher, product, landing_page, and authority_boundary are required")
	}
	if _, err := time.Parse(time.DateOnly, contract.ReleaseDate); err != nil {
		return fmt.Errorf("Census source release_date must be YYYY-MM-DD: %w", err)
	}
	if _, err := time.Parse(time.DateOnly, contract.AccessedAt); err != nil {
		return fmt.Errorf("Census source accessed_at must be YYYY-MM-DD: %w", err)
	}
	if !validSHA256(contract.ArchiveSHA256) || !validSHA256(contract.Workbook.SHA256) || !validSHA256(contract.Documentation.SHA256) {
		return fmt.Errorf("Census source artifact SHA-256 values are invalid")
	}
	if contract.GeneralPurpose.Sheet != "General Purpose" || contract.GeneralPurpose.RecordCount < 1 {
		return fmt.Errorf("Census General Purpose sheet contract is invalid")
	}
	if !reflect.DeepEqual(contract.GeneralPurpose.Columns, regulatory.CensusGeneralPurposeHeader) {
		return fmt.Errorf("Census General Purpose column contract does not match the importer")
	}
	if contract.Scope.StatesAndDC != 51 {
		return fmt.Errorf("Census source scope must contain 51 states/DC entries")
	}
	return nil
}

func (contract censusSourceContract) source() regulatory.Source {
	return regulatory.Source{
		ID:            contract.SourceID,
		Title:         contract.Product,
		URL:           contract.LandingPage,
		Kind:          "governmental_inventory",
		AccessedAt:    contract.AccessedAt,
		LastCheckedAt: contract.AccessedAt,
		Availability:  "available",
		Caveat:        contract.AuthorityBoundary,
	}
}

func validSHA256(value string) bool {
	if len(value) != 64 {
		return false
	}
	for _, character := range strings.ToLower(value) {
		if (character < '0' || character > '9') && (character < 'a' || character > 'f') {
			return false
		}
	}
	return true
}
