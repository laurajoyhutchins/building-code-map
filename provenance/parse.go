package provenance

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// ParseJSONLD is the mandatory admission boundary for externally authored BCM
// provenance. JSON Schema alone cannot check global identity and reference
// closure; parsed graphs pass the SAME typed graph validation as local exports.
// Unknown fields, unexpected contexts, JSON null, malformed references and
// conflicting identities fail closed without invoking a network context loader.
func ParseJSONLD(data []byte) (Graph, error) {
	if len(data) == 0 || len(data) > 16<<20 {
		return Graph{}, fmt.Errorf("PROV input must be 1 byte to 16 MiB")
	}
	root, err := objectFields(data, "@context", "@graph")
	if err != nil {
		return Graph{}, fmt.Errorf("invalid PROV root: %w", err)
	}
	if len(root) != 2 {
		return Graph{}, fmt.Errorf("PROV root requires only @context and @graph")
	}
	ctx, err := objectFields(root["@context"], "prov", "bcm", "xsd")
	if err != nil || len(ctx) != 3 {
		return Graph{}, fmt.Errorf("PROV context must fix prov, bcm and xsd")
	}
	for name, expected := range map[string]string{
		"prov": provNS,
		"bcm": bcmNS,
		"xsd": "http://www.w3.org/2001/XMLSchema#",
	} {
		value, err := stringField(ctx[name])
		if err != nil || value != expected {
			return Graph{}, fmt.Errorf("unapproved PROV context %s", name)
		}
	}
	var rawNodes []json.RawMessage
	if err := json.Unmarshal(root["@graph"], &rawNodes); err != nil || len(rawNodes) == 0 || len(rawNodes) > 250000 {
		return Graph{}, fmt.Errorf("invalid PROV node collection")
	}
	graph := Graph{Nodes: make([]Node, 0, len(rawNodes))}
	for _, raw := range rawNodes {
		node, edges, err := readNode(raw)
		if err != nil {
			return Graph{}, err
		}
		graph.Nodes = append(graph.Nodes, node)
		graph.Edges = append(graph.Edges, edges...)
	}
	if err := graph.Validate(); err != nil {
		return Graph{}, err
	}
	return graph, nil
}

var allowedNodeFields = map[string]bool{
	"@id": true, "@type": true,
	"bcm:sourceID": true, "bcm:sha256": true, "bcm:sourceURL": true,
	"bcm:legalLocator": true, "bcm:evidenceRole": true, "bcm:observedAt": true,
	"bcm:claimKind": true, "bcm:jurisdictionID": true,
	"bcm:asOf": true, "bcm:legallyMaterial": true,
	"prov:startedAtTime": true, "prov:endedAtTime": true,
	"prov:wasDerivedFrom": true, "prov:specializationOf": true,
	"prov:wasGeneratedBy": true, "prov:used": true,
	"prov:wasAssociatedWith": true, "prov:wasAttributedTo": true,
	"bcm:hasEvidenceLink": true, "bcm:targetsEvidenceLink": true,
}

var relationFields = map[string]bool{
	"prov:wasDerivedFrom": true, "prov:specializationOf": true,
	"prov:wasGeneratedBy": true, "prov:used": true,
	"prov:wasAssociatedWith": true, "prov:wasAttributedTo": true,
	"bcm:hasEvidenceLink": true, "bcm:targetsEvidenceLink": true,
}

func readNode(raw json.RawMessage) (Node, []Edge, error) {
	fields, err := objectFields(raw)
	if err != nil {
		return Node{}, nil, err
	}
	for key := range fields {
		if !allowedNodeFields[key] {
			return Node{}, nil, fmt.Errorf("unapproved PROV property %s", key)
		}
	}
	if _, ok := fields["@id"]; !ok {
		return Node{}, nil, fmt.Errorf("PROV node missing @id")
	}
	if _, ok := fields["@type"]; !ok {
		return Node{}, nil, fmt.Errorf("PROV node missing @type")
	}
	id, err := stringField(fields["@id"])
	if err != nil { return Node{}, nil, err }
	kind, err := stringField(fields["@type"])
	if err != nil { return Node{}, nil, err }
	n := Node{ID: id, Kind: Kind(kind)}
	literals := map[string]*string{
		"bcm:sourceID": &n.SourceID, "bcm:sha256": &n.SHA256,
		"bcm:legalLocator": &n.LegalLocator, "bcm:evidenceRole": &n.EvidenceRole,
		"bcm:claimKind": &n.ClaimKind, "bcm:jurisdictionID": &n.JurisdictionID,
	}
	for key, destination := range literals {
		if raw, exists := fields[key]; exists {
			value, err := stringField(raw)
			if err != nil { return Node{}, nil, fmt.Errorf("invalid %s: %w", key, err) }
			*destination = value
		}
	}
	if raw, exists := fields["bcm:legallyMaterial"]; exists {
		if invalidNull(raw) { return Node{}, nil, fmt.Errorf("invalid materiality null") }
		var value bool
		if err := json.Unmarshal(raw, &value); err != nil {
			return Node{}, nil, fmt.Errorf("invalid claim materiality: %w", err)
		}
		n.LegallyMaterial = &value
	}
	if raw, exists := fields["bcm:sourceURL"]; exists {
		value, err := referenceIRI(raw)
		if err != nil { return Node{}, nil, err }
		n.SourceURL = value
	}
	dates := map[string]*string{
		"bcm:observedAt": &n.ObservedAt,
		"prov:startedAtTime": &n.StartedAt,
		"prov:endedAtTime": &n.EndedAt,
		"bcm:asOf": &n.AsOf,
	}
	for key, destination := range dates {
		if raw, exists := fields[key]; exists {
			kind := "xsd:dateTime"
			if key == "bcm:asOf" { kind = "xsd:date" }
			value, err := typedLiteral(raw, kind)
			if err != nil { return Node{}, nil, fmt.Errorf("invalid %s: %w", key, err) }
			*destination = value
		}
	}
	var edges []Edge
	for key := range relationFields {
		raw, exists := fields[key]
		if !exists { continue }
		if invalidNull(raw) { return Node{}, nil, fmt.Errorf("null PROV edge %s", key) }
		var links []json.RawMessage
		if err := json.Unmarshal(raw, &links); err != nil || len(links) == 0 {
			return Node{}, nil, fmt.Errorf("invalid PROV references %s", key)
		}
		for _, link := range links {
			iri, err := referenceIRI(link)
			if err != nil { return Node{}, nil, fmt.Errorf("invalid relation %s: %w", key, err) }
			edges = append(edges, Edge{From: id, Predicate: key, To: iri})
		}
	}
	return n, edges, nil
}

func invalidNull(raw json.RawMessage) bool {
	return bytes.Equal(bytes.TrimSpace(raw), []byte("null"))
}

func objectFields(raw json.RawMessage, allowed ...string) (map[string]json.RawMessage, error) {
	if len(raw) == 0 || invalidNull(raw) {
		return nil, fmt.Errorf("missing JSON object")
	}
	if !strings.HasPrefix(strings.TrimSpace(string(raw)), "{") {
		return nil, fmt.Errorf("JSON object expected")
	}
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(raw, &obj); err != nil { return nil, err }
	if allowed != nil {
		for key := range obj {
			found := false
			for _, item := range allowed { if item == key { found = true; break } }
			if !found { return nil, fmt.Errorf("unknown JSON property %s", key) }
		}
	}
	return obj, nil
}

func stringField(raw json.RawMessage) (string, error) {
	if len(raw) == 0 || invalidNull(raw) {
		return "", fmt.Errorf("string property missing or null")
	}
	var result string
	if err := json.Unmarshal(raw, &result); err != nil {
		return "", err
	}
	return result, nil
}

func referenceIRI(raw json.RawMessage) (string, error) {
	fields, err := objectFields(raw, "@id")
	if err != nil || len(fields) != 1 {
		return "", fmt.Errorf("reference must contain exactly one @id")
	}
	return stringField(fields["@id"])
}

func typedLiteral(raw json.RawMessage, kind string) (string, error) {
	fields, err := objectFields(raw, "@value", "@type")
	if err != nil || len(fields) != 2 {
		return "", fmt.Errorf("typed literal must contain @value and @type")
	}
	actualKind, err := stringField(fields["@type"])
	if err != nil || actualKind != kind {
		return "", fmt.Errorf("typed literal requires %s", kind)
	}
	return stringField(fields["@value"])
}
