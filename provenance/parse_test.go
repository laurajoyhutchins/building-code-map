package provenance

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestExternalJSONLDReadbackHasSameGraph(t *testing.T) {
	g := testGraph()
	first, err := g.JSONLD()
	if err != nil {
		t.Fatal(err)
	}
	imported, err := ParseJSONLD(first)
	if err != nil {
		t.Fatal(err)
	}
	again, err := imported.JSONLD()
	if err != nil || !bytes.Equal(first, again) {
		t.Fatalf("JSON-LD readback is not deterministic: %v", err)
	}
}

func TestImportClosesGraphNotJustSchemaShape(t *testing.T) {
	first, err := testGraph().JSONLD()
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		name   string
		mutate func(map[string]any)
	}{
		{"orphan used target", func(v map[string]any) {
			v["@graph"].([]any)[2].(map[string]any)["prov:used"] = []any{map[string]any{"@id": "urn:bcm:prov:missing"}}
		}},
		{"wrong type derived target", func(v map[string]any) {
			v["@graph"].([]any)[1].(map[string]any)["prov:wasDerivedFrom"] = []any{map[string]any{"@id": "urn:bcm:prov:action:c"}}
		}},
		{"duplicate source node", func(v map[string]any) { nodes := v["@graph"].([]any); v["@graph"] = append(nodes, nodes[0]) }},
		{"duplicate relation", func(v map[string]any) {
			n := v["@graph"].([]any)[2].(map[string]any)
			n["prov:used"] = []any{map[string]any{"@id": "urn:bcm:prov:copy:b"}, map[string]any{"@id": "urn:bcm:prov:copy:b"}}
		}},
		{"unapproved context", func(v map[string]any) { v["@context"].(map[string]any)["prov"] = "https://example.invalid/fake#" }},
		{"forbidden root property", func(v map[string]any) { v["verified"] = true }},
		{"forbidden node property", func(v map[string]any) { v["@graph"].([]any)[0].(map[string]any)["bcm:legalVerdict"] = "verified" }},
		{"wrong source URL shape", func(v map[string]any) {
			v["@graph"].([]any)[1].(map[string]any)["bcm:sourceURL"] = "https://example.invalid/"
		}},
		{"invalid node null", func(v map[string]any) { v["@graph"].([]any)[0].(map[string]any)["bcm:sha256"] = nil }},
		{"wrong typed time", func(v map[string]any) {
			v["@graph"].([]any)[2].(map[string]any)["prov:startedAtTime"] = map[string]any{"@value": "2026-10-28T12:00:00Z", "@type": "xsd:string"}
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var value map[string]any
			if err := json.Unmarshal(first, &value); err != nil {
				t.Fatal(err)
			}
			tc.mutate(value)
			altered, err := json.Marshal(value)
			if err != nil {
				t.Fatal(err)
			}
			if _, err := ParseJSONLD(altered); err == nil {
				t.Fatal("invalid external lineage admitted")
			}
		})
	}
	for _, invalid := range []string{"", "[]", "null", "{}{}", strings.Repeat("x", 17<<20)} {
		if _, err := ParseJSONLD([]byte(invalid)); err == nil {
			t.Fatalf("accepted %d bytes invalid JSON-LD", len(invalid))
		}
	}
}

func TestClaimIdentityBindsLegalScope(t *testing.T) {
	first, err := ClaimID("DEMO-CLAIM", "code_adoption", "DEMO-XX", "2026-10-27", false)
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		kind, jurisdiction, date string
		material                 bool
	}{
		{"code_adoption", "DEMO-XX", "2026-10-28", false},
		{"code_adoption", "DEMO-XX", "2026-10-27", true},
		{"code_adoption", "DEMO-YY", "2026-10-27", false},
		{"other_claim", "DEMO-XX", "2026-10-27", false},
	}
	for _, c := range cases {
		id, err := ClaimID("DEMO-CLAIM", c.kind, c.jurisdiction, c.date, c.material)
		if err != nil || first == id {
			t.Fatalf("claim scope aliased %s: %v", id, err)
		}
	}
	if _, err := ClaimID("DEMO-CLAIM", "code_adoption", "DEMO-XX", "yesterday", false); err == nil {
		t.Fatal("invalid date admitted")
	}
	value := false
	g := Graph{Nodes: []Node{{ID: first, Kind: Entity, SourceID: "DEMO-CLAIM", ClaimKind: "code_adoption", JurisdictionID: "DEMO-XX", AsOf: "2026-10-27", LegallyMaterial: &value}}}
	data, err := g.JSONLD()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(data, []byte("\"bcm:legallyMaterial\": false")) {
		t.Fatal("false materiality omitted")
	}
	if _, err := ParseJSONLD(data); err != nil {
		t.Fatal(err)
	}
	g.Nodes[0].AsOf = "2026-10-28"
	if err := g.Validate(); err == nil {
		t.Fatal("claim @id not bound to applicability date")
	}
}
