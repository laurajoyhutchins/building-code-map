// Package provenance emits a bounded, deterministic PROV-O JSON-LD graph.
// It models recorded lineage only: a graph is never an approval or legal verdict.
package provenance

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
)

const (
	provNS = "http://www.w3.org/ns/prov#"
	bcmNS  = "https://github.com/laurajoyhutchins/building-code-map/ns/provenance#"
)

type Kind string

const (
	Entity   Kind = "prov:Entity"
	Activity Kind = "prov:Activity"
	Agent    Kind = "prov:Agent"
)

type Node struct {
	ID           string
	Kind         Kind
	SourceID     string
	SHA256       string
	SourceURL    string
	LegalLocator string
	EvidenceRole string
	ObservedAt   string
	StartedAt    string
	EndedAt      string
	ClaimKind    string
	JurisdictionID string
	AsOf         string
	LegallyMaterial *bool
}

type Edge struct {
	From      string
	Predicate string
	To        string
}

type Graph struct {
	Nodes []Node
	Edges []Edge
}

var digestRE = regexp.MustCompile("^[0-9a-f]{64}$")
var tagRE = regexp.MustCompile("^[a-z][a-z0-9-]*$")

// StableID creates an opaque, collision-resistant URI from an existing
// source identity. Never construct URIs by concatenating untrusted IDs.
func StableID(tag, sourceID string) (string, error) {
	if !tagRE.MatchString(tag) || strings.TrimSpace(sourceID) == "" || len(sourceID) > 4096 {
		return "", fmt.Errorf("invalid provenance identity")
	}
	sum := sha256.Sum256([]byte(sourceID))
	return "urn:bcm:prov:" + tag + ":" + hex.EncodeToString(sum[:]), nil
}

// ClaimID binds legal scope and applicability to immutable provenance identity.
// A JSON tuple avoids ambiguous string concatenation and delimiter collisions.
func ClaimID(sourceID, kind, jurisdiction, asOf string, legallyMaterial bool) (string, error) {
	if strings.TrimSpace(sourceID) == "" || strings.TrimSpace(kind) == "" || strings.TrimSpace(jurisdiction) == "" {
		return "", fmt.Errorf("claim scope is incomplete")
	}
	if _, err := time.Parse(time.DateOnly, asOf); err != nil {
		return "", fmt.Errorf("claim applicability date is invalid")
	}
	scope, err := json.Marshal([]any{sourceID, kind, jurisdiction, asOf, legallyMaterial})
	if err != nil { return "", err }
	return StableID("claim", string(scope))
}

func validIRI(value string) bool {
	if value == "" || strings.ContainsAny(value, " \t\r\n<>\"{}|\\") {
		return false
	}
	parsed, err := url.Parse(value)
	if err != nil || parsed.Scheme == "" || parsed.Fragment != "" {
		return false
	}
	if parsed.Scheme == "urn" {
		return parsed.Opaque != ""
	}
	return parsed.Scheme == "https" && parsed.Host != "" && parsed.User == nil
}

func validSourceURL(value string) bool {
	if value == "" {
		return true
	}
	parsed, err := url.Parse(value)
	return err == nil && (parsed.Scheme == "https" || parsed.Scheme == "http") &&
		parsed.Host != "" && parsed.User == nil
}

func validDateTime(value string) bool {
	if value == "" {
		return true
	}
	_, err := time.Parse(time.RFC3339Nano, value)
	return err == nil
}

func permittedRole(role string) bool {
	switch role {
	case "", "establishes", "corroborates", "contradicts", "defines", "incorporates", "supersedes":
		return true
	}
	return false
}

// Validate checks closed references and PROV-O property domain/range types.
// This is structural and semantic-admission validation, not an RDF reasoner.
func (g Graph) Validate() error {
	if len(g.Nodes) == 0 {
		return fmt.Errorf("PROV graph has no nodes")
	}
	nodes := make(map[string]Node, len(g.Nodes))
	for _, n := range g.Nodes {
		if !validIRI(n.ID) {
			return fmt.Errorf("invalid PROV node IRI %q", n.ID)
		}
		if n.Kind != Entity && n.Kind != Activity && n.Kind != Agent {
			return fmt.Errorf("invalid PROV node kind %q", n.Kind)
		}
		if _, exists := nodes[n.ID]; exists {
			return fmt.Errorf("duplicate PROV node %q", n.ID)
		}
		nodes[n.ID] = n
		if n.SHA256 != "" && !digestRE.MatchString(n.SHA256) {
			return fmt.Errorf("invalid PROV SHA-256")
		}
		if !validSourceURL(n.SourceURL) || !validDateTime(n.ObservedAt) ||
			!validDateTime(n.StartedAt) || !validDateTime(n.EndedAt) {
			return fmt.Errorf("invalid PROV metadata at %q", n.ID)
		}
		if !permittedRole(n.EvidenceRole) {
			return fmt.Errorf("unknown evidence role at %q", n.ID)
		}
		if n.Kind != Entity &&
			(n.SHA256 != "" || n.SourceURL != "" || n.LegalLocator != "" || n.EvidenceRole != "" || n.ObservedAt != "") {
			return fmt.Errorf("entity-only properties on %q", n.ID)
		}
		if n.Kind != Entity && (n.ClaimKind != "" || n.JurisdictionID != "" || n.AsOf != "" || n.LegallyMaterial != nil) {
			return fmt.Errorf("claim properties on non-entity %q", n.ID)
		}
		hasClaim := n.ClaimKind != "" || n.JurisdictionID != "" || n.AsOf != "" || n.LegallyMaterial != nil || strings.HasPrefix(n.ID, "urn:bcm:prov:claim:")
		if hasClaim {
			if n.Kind != Entity || n.LegallyMaterial == nil {
				return fmt.Errorf("incomplete claim scope at %q", n.ID)
			}
			expected, err := ClaimID(n.SourceID, n.ClaimKind, n.JurisdictionID, n.AsOf, *n.LegallyMaterial)
			if err != nil || n.ID != expected {
				return fmt.Errorf("claim identity does not bind legal scope at %q", n.ID)
			}
		}
		if n.Kind != Activity && (n.StartedAt != "" || n.EndedAt != "") {
			return fmt.Errorf("activity times on non-activity %q", n.ID)
		}
		if n.StartedAt != "" && n.EndedAt != "" {
			start, _ := time.Parse(time.RFC3339Nano, n.StartedAt)
			end, _ := time.Parse(time.RFC3339Nano, n.EndedAt)
			if end.Before(start) {
				return fmt.Errorf("activity ends before it starts")
			}
		}
	}
	seen := map[Edge]bool{}
	for _, e := range g.Edges {
		from, okFrom := nodes[e.From]
		to, okTo := nodes[e.To]
		if !okFrom || !okTo {
			return fmt.Errorf("orphan PROV relation %q", e.Predicate)
		}
		if seen[e] {
			return fmt.Errorf("duplicate PROV relation %q", e.Predicate)
		}
		seen[e] = true
		switch e.Predicate {
		case "prov:wasDerivedFrom", "prov:specializationOf", "bcm:hasEvidenceLink":
			if from.Kind != Entity || to.Kind != Entity || e.From == e.To {
				return fmt.Errorf("invalid entity relation %q", e.Predicate)
			}
		case "bcm:targetsEvidenceLink":
			if from.Kind != Entity || to.Kind != Entity || e.From == e.To || from.EvidenceRole != "supersedes" || to.EvidenceRole == "" {
				return fmt.Errorf("invalid superseded evidence target")
			}
		case "prov:wasGeneratedBy":
			if from.Kind != Entity || to.Kind != Activity {
				return fmt.Errorf("invalid generation relation")
			}
		case "prov:used":
			if from.Kind != Activity || to.Kind != Entity {
				return fmt.Errorf("invalid usage relation")
			}
		case "prov:wasAssociatedWith":
			if from.Kind != Activity || to.Kind != Agent {
				return fmt.Errorf("invalid association relation")
			}
		case "prov:wasAttributedTo":
			if from.Kind != Entity || to.Kind != Agent {
				return fmt.Errorf("invalid attribution relation")
			}
		default:
			return fmt.Errorf("unsupported PROV predicate %q", e.Predicate)
		}
	}
	return nil
}

// JSONLD emits self-contained JSON-LD 1.1 with PROV-O IRIs and stable ordering.
// Reference objects use @id; literals never become accidental IRI references.
func (g Graph) JSONLD() ([]byte, error) {
	if err := g.Validate(); err != nil {
		return nil, err
	}
	nodes := append([]Node(nil), g.Nodes...)
	edges := append([]Edge(nil), g.Edges...)
	sort.Slice(nodes, func(i, j int) bool { return nodes[i].ID < nodes[j].ID })
	sort.Slice(edges, func(i, j int) bool {
		if edges[i].From != edges[j].From {
			return edges[i].From < edges[j].From
		}
		if edges[i].Predicate != edges[j].Predicate {
			return edges[i].Predicate < edges[j].Predicate
		}
		return edges[i].To < edges[j].To
	})
	graph := make([]map[string]any, 0, len(nodes))
	for _, n := range nodes {
		item := map[string]any{"@id": n.ID, "@type": string(n.Kind)}
		if n.SourceID != "" {
			item["bcm:sourceID"] = n.SourceID
		}
		if n.SHA256 != "" {
			item["bcm:sha256"] = n.SHA256
		}
		if n.SourceURL != "" {
			item["bcm:sourceURL"] = map[string]string{"@id": n.SourceURL}
		}
		if n.LegalLocator != "" {
			item["bcm:legalLocator"] = n.LegalLocator
		}
		if n.EvidenceRole != "" {
			item["bcm:evidenceRole"] = n.EvidenceRole
		}
		if n.ClaimKind != "" {
			item["bcm:claimKind"] = n.ClaimKind
			item["bcm:jurisdictionID"] = n.JurisdictionID
			item["bcm:asOf"] = map[string]string{"@value": n.AsOf, "@type": "xsd:date"}
			item["bcm:legallyMaterial"] = *n.LegallyMaterial
		}
		if n.ObservedAt != "" {
			item["bcm:observedAt"] = dateTimeLiteral(n.ObservedAt)
		}
		if n.StartedAt != "" {
			item["prov:startedAtTime"] = dateTimeLiteral(n.StartedAt)
		}
		if n.EndedAt != "" {
			item["prov:endedAtTime"] = dateTimeLiteral(n.EndedAt)
		}
		for _, e := range edges {
			if e.From != n.ID {
				continue
			}
			refs, _ := item[e.Predicate].([]map[string]string)
			item[e.Predicate] = append(refs, map[string]string{"@id": e.To})
		}
		graph = append(graph, item)
	}
	doc := map[string]any{
		"@context": map[string]string{"prov": provNS, "bcm": bcmNS, "xsd": "http://www.w3.org/2001/XMLSchema#"},
		"@graph":   graph,
	}
	result, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		return nil, err
	}
	return append(result, '\n'), nil
}

func dateTimeLiteral(value string) map[string]string {
	return map[string]string{"@value": value, "@type": "xsd:dateTime"}
}
