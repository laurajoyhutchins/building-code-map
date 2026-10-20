package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAccuracyCLISeparatesFalseClaimsFromSafeAbstentions(t *testing.T) {
	dir := t.TempDir()
	truth := filepath.Join(dir, "truth.json")
	results := filepath.Join(dir, "results.json")
	content := `{"schema_version":"1","cases":[{"id":"DEMO-CASE","project_id":"DEMO-PROJECT","jurisdiction_id":"DEMO-XX","code_family":"building","as_of":"2026-08-25","reviewed_at":"2026-08-26","source_ids":["DEMO-PRIMARY"],"expected":[{"status":"resolved","edition":"DEMO-2024"}]}]}`
	if err := os.WriteFile(truth, []byte(content), 0600); err != nil {
		t.Fatal(err)
	}
	for _, tt := range []struct {
		name string
		observed string
		strict bool
		code int
		summary string
	}{
		{
			name: "wrong edition fails", observed: `{"results":[{"project_id":"DEMO-PROJECT","jurisdiction_id":"DEMO-XX","as_of":"2026-08-25","verdict":"verified","code_set":[{"family":"building","status":"resolved","edition":"DEMO-2021"}]}]}`,
			code: 1, summary: "\"incorrect_assertions\": 1",
		},
		{
			name: "safe abstention", observed: `{"results":[{"project_id":"DEMO-PROJECT","jurisdiction_id":"DEMO-XX","as_of":"2026-08-25","verdict":"not_verified","code_set":[{"family":"building","status":"insufficient_evidence"}]}]}`,
			code: 0, summary: "\"abstentions\": 1",
		},
		{
			name: "strict abstention gate", observed: `{"results":[{"project_id":"DEMO-PROJECT","jurisdiction_id":"DEMO-XX","as_of":"2026-08-25","verdict":"not_verified","code_set":[{"family":"building","status":"insufficient_evidence"}]}]}`,
			strict: true, code: 1, summary: "\"abstentions\": 1",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			if err := os.WriteFile(results, []byte(tt.observed), 0600); err != nil {
				t.Fatal(err)
			}
			args := []string{"-truth", truth, "-results", results}
			if tt.strict {
				args = append(args, "-strict")
			}
			var stdout, stderr bytes.Buffer
			got := run(args, &stdout, &stderr)
			if got != tt.code || !strings.Contains(stdout.String(), tt.summary) {
				t.Fatalf("exit=%d stdout=%s stderr=%s", got, stdout.String(), stderr.String())
			}
		})
	}
}
