package snapshotbuild

import (
	"context"
	"time"

	"building-code-map/backend/internal/snapshotmanifest"
)

const (
	ReceiptSchemaVersion = "1.0"
	PipelineIdentity     = "duckdb-snapshot-build-1"
)

type Kind string

const (
	KindBoundary Kind = "boundary"
	KindGeocoder Kind = "geocoder"
)

type Source struct {
	Name                 string
	Path                 string
	Locator              string
	Publisher            string
	Product              string
	Layer                string
	Vintage              string
	RetrievedAt          string
	LicenseReviewStatus  string
	RedistributionStatus string
	InputCRS             string
	Transformation       string
}

type Request struct {
	Kind                            Kind
	DuckDBPath                      string
	ExpectedDuckDBVersion           string
	ExpectedDuckDBSHA256            string
	SpatialExtensionPath            string
	ExpectedSpatialExtensionSHA256 string
	WorkingDirectory                string
	OutputPath                      string
	BuilderVersion                  string
	BuilderRevision                 string
	BuiltAt                         time.Time
	OutputCRS                       string
	Sources                         []Source
}

type Invocation struct {
	Executable         string
	DatabasePath       string
	SQL                string
	WorkingDirectory   string
	ExtensionDirectory string
}

type Executor interface {
	Version(context.Context, string) (string, error)
	Run(context.Context, Invocation) ([]byte, error)
}

type Counts struct {
	Accepted    int64 `json:"accepted"`
	Rejected    int64 `json:"rejected"`
	Duplicate   int64 `json:"duplicate"`
	Quarantined int64 `json:"quarantined"`
}

type Check struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Detail string `json:"detail,omitempty"`
}

type AuditReport struct {
	SchemaVersion string            `json:"schema_version"`
	Kind          Kind              `json:"kind"`
	Counts        Counts            `json:"record_counts"`
	Checks        []Check           `json:"checks"`
	Normalized    map[string]string `json:"normalized,omitempty"`
}

type ToolIdentity struct {
	Pipeline               string `json:"pipeline"`
	DuckDBVersion          string `json:"duckdb_version"`
	DuckDBSHA256           string `json:"duckdb_sha256"`
	SpatialExtensionSHA256 string `json:"spatial_extension_sha256,omitempty"`
	SQLContractSHA256      string `json:"sql_contract_sha256"`
}

type Receipt struct {
	SchemaVersion string          `json:"schema_version"`
	Kind          Kind            `json:"kind"`
	BuiltAt       string          `json:"built_at"`
	Tool          ToolIdentity    `json:"tool"`
	Sources       []SourceReceipt `json:"sources"`
	Counts        Counts          `json:"record_counts"`
	Checks        []Check         `json:"checks"`
	OutputSHA256  string          `json:"output_sha256,omitempty"`
	OutputSize    int64           `json:"output_size_bytes,omitempty"`
	ManifestPath  string          `json:"manifest_path,omitempty"`
}

type SourceReceipt struct {
	Name    string `json:"name"`
	Locator string `json:"locator"`
	SHA256  string `json:"sha256"`
}

type Result struct {
	Report   AuditReport
	Receipt  Receipt
	Manifest snapshotmanifest.Manifest
}

type FinalizeFunc func(context.Context, map[string]string, string) error
