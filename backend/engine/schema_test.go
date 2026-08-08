package engine

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestGoProducedResultContainsSchemaContractFields(t *testing.T) {
	raw, err := json.Marshal(Result{
		SchemaVersion: SchemaVersion,
		Query:         NormalizedQuery{ApplicabilityDate: "2026-08-06"},
		Provenance:    Provenance{SourceCommit: "0123456789abcdef0123456789abcdef01234567", EngineVersion: "0.1.0", BundleManifestDigest: "digest", BoundarySnapshotDigest: "boundary", RegulatoryCatalogDigest: "regulatory"},
		Diagnostics:   []Diagnostic{},
	})
	if err != nil {
		t.Fatal(err)
	}
	var object map[string]json.RawMessage
	if err := json.Unmarshal(raw, &object); err != nil {
		t.Fatal(err)
	}
	for _, key := range []string{"schema_version", "query", "location", "resolution", "provenance", "diagnostics"} {
		if _, ok := object[key]; !ok {
			t.Fatalf("produced result missing %q", key)
		}
	}
	for _, schemaName := range []string{"engine-query.schema.json", "engine-result.schema.json", "engine-error.schema.json", "engine-bundle.schema.json"} {
		schema, err := os.ReadFile(filepath.Join("..", "..", "schemas", schemaName))
		if err != nil {
			t.Fatal(err)
		}
		var schemaObject map[string]any
		if err := json.Unmarshal(schema, &schemaObject); err != nil {
			t.Fatalf("schema %s: %v", schemaName, err)
		}
		if schemaObject["$schema"] == nil {
			t.Fatalf("schema %s has no $schema", schemaName)
		}
	}
}
