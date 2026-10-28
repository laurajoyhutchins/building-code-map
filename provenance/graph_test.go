package provenance

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func testGraph() Graph {
	return Graph{
		Nodes: []Node{
			{ID: "urn:bcm:prov:doc:a", Kind: Entity, SourceID: "doc:a", SourceURL: "https://example.org/law"},
			{ID: "urn:bcm:prov:copy:b", Kind: Entity, SourceID: "artifact:b", SHA256: strings.Repeat("a", 64), ObservedAt: "2026-10-28T12:00:00Z"},
			{ID: "urn:bcm:prov:action:c", Kind: Activity, StartedAt: "2026-10-28T12:01:00Z", EndedAt: "2026-10-28T12:02:00Z"},
			{ID: "urn:bcm:prov:agent:d", Kind: Agent, SourceID: "reviewer:d"},
		},
		Edges: []Edge{
			{From: "urn:bcm:prov:copy:b", Predicate: "prov:wasDerivedFrom", To: "urn:bcm:prov:doc:a"},
			{From: "urn:bcm:prov:action:c", Predicate: "prov:used", To: "urn:bcm:prov:copy:b"},
			{From: "urn:bcm:prov:action:c", Predicate: "prov:wasAssociatedWith", To: "urn:bcm:prov:agent:d"},
		},
	}
}

func TestJSONLDDeterministicWithExplicitPROVOIRIs(t *testing.T) {
	g := testGraph()
	data, err := g.JSONLD()
	if err != nil {
		t.Fatal(err)
	}
	g.Nodes[0], g.Nodes[3] = g.Nodes[3], g.Nodes[0]
	g.Edges[0], g.Edges[2] = g.Edges[2], g.Edges[0]
	again, err := g.JSONLD()
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(data, again) {
		t.Fatal("JSON-LD depends on input ordering")
	}
	var doc map[string]any
	if err := json.Unmarshal(data, &doc); err != nil {
		t.Fatal(err)
	}
	ctx := doc["@context"].(map[string]any)
	if ctx["prov"] != provNS || ctx["bcm"] != bcmNS {
		t.Fatalf("unexpected vocabulary bindings: %v", ctx)
	}
	nodes := doc["@graph"].([]any)
	if len(nodes) != 4 {
		t.Fatalf("nodes = %d", len(nodes))
	}
	var sawUse, sawTime, sawSourceURL bool
	for _, raw := range nodes {
		n := raw.(map[string]any)
		if n["@id"] == "urn:bcm:prov:action:c" {
			refs := n["prov:used"].([]any)
			sawUse = len(refs) == 1 && refs[0].(map[string]any)["@id"] == "urn:bcm:prov:copy:b"
			started := n["prov:startedAtTime"].(map[string]any)
			sawTime = started["@type"] == "xsd:dateTime"
		}
		if n["@id"] == "urn:bcm:prov:doc:a" {
			sawSourceURL = n["bcm:sourceURL"].(map[string]any)["@id"] == "https://example.org/law"
		}
	}
	if !sawUse || !sawTime || !sawSourceURL {
		t.Fatalf("PROV predicates, typed time, or source IRI missing: %s", data)
	}
}

func TestStableIDUsesOpaqueSourceIdentity(t *testing.T) {
	a, err := StableID("document", "untrusted id?<>")
	if err != nil || !strings.HasPrefix(a, "urn:bcm:prov:document:") {
		t.Fatal(a, err)
	}
	b, err := StableID("artifact", "untrusted id?<>")
	if err != nil || a == b {
		t.Fatalf("typed stable IDs collapsed")
	}
	if _, err := StableID("../document", "x"); err == nil {
		t.Fatal("unsafe tag accepted")
	}
	if _, err := StableID("document", "  "); err == nil {
		t.Fatal("empty ID accepted")
	}
}

func TestPROVGraphRejectsInvalidLineage(t *testing.T) {
	base := testGraph()
	tests := []struct {
		name   string
		change func(*Graph)
	}{
		{"missing entity", func(g *Graph) { g.Nodes = g.Nodes[1:] }},
		{"duplicate node", func(g *Graph) { g.Nodes = append(g.Nodes, g.Nodes[0]) }},
		{"duplicate edge", func(g *Graph) { g.Edges = append(g.Edges, g.Edges[0]) }},
		{"wrong usage domain", func(g *Graph) { g.Edges[1].From = "urn:bcm:prov:doc:a" }},
		{"wrong association range", func(g *Graph) { g.Edges[2].To = "urn:bcm:prov:copy:b" }},
		{"wrong derivation range", func(g *Graph) { g.Edges[0].To = "urn:bcm:prov:action:c" }},
		{"self derivation", func(g *Graph) { g.Edges[0].To = g.Edges[0].From }},
		{"invented predicate", func(g *Graph) { g.Edges[0].Predicate = "prov:wasApprovedBy" }},
		{"unsafe IRI", func(g *Graph) { g.Nodes[0].ID = "https://evil.test/ bad" }},
		{"invalid digest", func(g *Graph) { g.Nodes[1].SHA256 = "not-pinned" }},
		{"invalid date", func(g *Graph) { g.Nodes[1].ObservedAt = "tomorrow" }},
		{"inverted activity", func(g *Graph) { g.Nodes[2].EndedAt = "2026-10-28T11:00:00Z" }},
		{"entity metadata on agent", func(g *Graph) { g.Nodes[3].SHA256 = strings.Repeat("a", 64) }},
		{"unrecognized evidence role", func(g *Graph) { g.Nodes[0].EvidenceRole = "probably_verified" }},
		{"source URL with credentials", func(g *Graph) { g.Nodes[0].SourceURL = "https://token@example.org/" }},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			g := Graph{Nodes: append([]Node(nil), base.Nodes...), Edges: append([]Edge(nil), base.Edges...)}
			tc.change(&g)
			if _, err := g.JSONLD(); err == nil {
				t.Fatal("invalid graph exported")
			}
		})
	}
}
