package engine

import (
	"encoding/json"
	"testing"
	"time"
)

func TestResultCarriesStableSchemaAndComponentProvenance(t *testing.T) {
	result := Result{
		SchemaVersion: SchemaVersion,
		Query: NormalizedQuery{
			Point:             &Point{Longitude: -104.99, Latitude: 39.74},
			ApplicabilityDate: "2026-08-06",
		},
		Provenance: Provenance{
			SourceCommit:            "0123456789abcdef0123456789abcdef01234567",
			EngineVersion:           "0.1.0",
			BundleManifestDigest:    "sha256:bundle",
			BoundarySnapshotDigest:  "sha256:boundary",
			RegulatoryCatalogDigest: "sha256:regulatory",
			GeocoderSnapshotDigest:  "sha256:geocoder",
		},
	}

	raw, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"schema_version", "query", "location", "resolution", "provenance", "diagnostics"} {
		if _, ok := decoded[key]; !ok {
			t.Fatalf("result JSON missing %q: %s", key, raw)
		}
	}
	provenance := decoded["provenance"].(map[string]any)
	for _, key := range []string{"source_commit", "engine_version", "bundle_manifest_digest", "boundary_snapshot_digest", "regulatory_catalog_digest", "geocoder_snapshot_digest"} {
		if _, ok := provenance[key]; !ok {
			t.Fatalf("provenance JSON missing %q: %s", key, raw)
		}
	}
}

func TestSortDiagnosticsProducesStableSeverityAndMessageOrder(t *testing.T) {
	diagnostics := []Diagnostic{
		{Severity: "warning", Code: "z", Message: "z"},
		{Severity: "error", Code: "b", Message: "b"},
		{Severity: "error", Code: "a", Message: "a"},
	}
	SortDiagnostics(diagnostics)
	if diagnostics[0].Code != "a" || diagnostics[1].Code != "b" || diagnostics[2].Code != "z" {
		t.Fatalf("diagnostics = %#v", diagnostics)
	}
}

func TestResultGeneratedAtUsesAnExplicitTimeValue(t *testing.T) {
	want := time.Date(2026, 8, 6, 12, 34, 56, 0, time.UTC)
	result := Result{Resolution: Resolution{GeneratedAt: want}}
	if !result.Resolution.GeneratedAt.Equal(want) {
		t.Fatalf("generated_at = %s, want %s", result.Resolution.GeneratedAt, want)
	}
}
