package engine

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/laurajoyhutchins/building-code-map/provenance"
)

// ProjectClaimPROV is a read-only projection of one existing EvidenceClaim.
// It retains both favorable and adverse evidence relationships. It deliberately
// does not call VerifyClaimEvidence, infer review agents, manufacture extraction
// activities, or promote any claim to verified.
func ProjectClaimPROV(claim EvidenceClaim) (provenance.Graph, error) {
	if strings.TrimSpace(claim.ID) == "" || strings.TrimSpace(claim.Kind) == "" ||
		strings.TrimSpace(claim.JurisdictionID) == "" {
		return provenance.Graph{}, fmt.Errorf("PROV claim identity is incomplete")
	}
	if _, err := time.Parse(time.DateOnly, claim.AsOf); err != nil {
		return provenance.Graph{}, fmt.Errorf("PROV claim has invalid applicability date")
	}
	if len(claim.Links) == 0 {
		return provenance.Graph{}, fmt.Errorf("PROV claim has no cited evidence")
	}
	claimID, err := provenance.StableID("claim", claim.ID)
	if err != nil {
		return provenance.Graph{}, err
	}

	g := provenance.Graph{}
	nodes := map[string]provenance.Node{}
	add := func(n provenance.Node) error {
		if prev, ok := nodes[n.ID]; ok {
			if !reflect.DeepEqual(prev, n) {
				return fmt.Errorf("conflicting PROV node %s", n.ID)
			}
			return nil
		}
		nodes[n.ID] = n
		g.Nodes = append(g.Nodes, n)
		return nil
	}
	if err := add(provenance.Node{ID: claimID, Kind: provenance.Entity, SourceID: claim.ID}); err != nil {
		return provenance.Graph{}, err
	}
	seen := map[string]EvidenceLink{}
	artifactDocument := map[string]string{}
	anchorArtifact := map[string]string{}
	for _, link := range claim.Links {
		if err := validateEvidenceLink(link, claim.ID); err != nil {
			return provenance.Graph{}, err
		}
		if err := acceptEvidenceLinkIdentity(seen, link); err != nil {
			return provenance.Graph{}, err
		}
		if link.Document.JurisdictionID != claim.JurisdictionID {
			return provenance.Graph{}, fmt.Errorf("source jurisdiction differs from claim")
		}
		docID, err := provenance.StableID("document", link.Document.ID)
		if err != nil {
			return provenance.Graph{}, err
		}
		artID, err := provenance.StableID("artifact", link.Artifact.ID)
		if err != nil {
			return provenance.Graph{}, err
		}
		anchorID, err := provenance.StableID("anchor", link.Anchor.ID)
		if err != nil {
			return provenance.Graph{}, err
		}
		evidenceID, err := provenance.StableID("evidence-link", link.ID)
		if err != nil {
			return provenance.Graph{}, err
		}
		if owner, ok := artifactDocument[artID]; ok && owner != docID {
			return provenance.Graph{}, fmt.Errorf("artifact identity refers to multiple documents")
		}
		artifactDocument[artID] = docID
		if owner, ok := anchorArtifact[anchorID]; ok && owner != artID {
			return provenance.Graph{}, fmt.Errorf("text anchor identity refers to multiple artifacts")
		}
		anchorArtifact[anchorID] = artID
		items := []provenance.Node{
			{ID: docID, Kind: provenance.Entity, SourceID: link.Document.ID, SourceURL: link.Document.CanonicalURL},
			{ID: artID, Kind: provenance.Entity, SourceID: link.Artifact.ID,
				SHA256: strings.ToLower(link.Artifact.SHA256), SourceURL: link.Artifact.SourceURL,
				ObservedAt: link.Artifact.ObservedAt},
			{ID: anchorID, Kind: provenance.Entity, SourceID: link.Anchor.ID,
				SHA256: strings.ToLower(link.Anchor.TextSHA256), LegalLocator: link.Anchor.LegalLocator},
			{ID: evidenceID, Kind: provenance.Entity, SourceID: link.ID, EvidenceRole: string(link.Role)},
		}
		for _, n := range items {
			if err := add(n); err != nil {
				return provenance.Graph{}, err
			}
		}
		// A legal document, its acquired manifestation, and an exact-text anchor
		// are distinct entities. The evidence link is derived from the anchor.
		// The claim HAS an evidence link; it is not automatically legally supported
		// or derived from every link (a link may contradict the claim).
		g.Edges = append(g.Edges,
			provenance.Edge{From: artID, Predicate: "prov:wasDerivedFrom", To: docID},
			provenance.Edge{From: anchorID, Predicate: "prov:wasDerivedFrom", To: artID},
			provenance.Edge{From: evidenceID, Predicate: "prov:wasDerivedFrom", To: anchorID},
			provenance.Edge{From: claimID, Predicate: "bcm:hasEvidenceLink", To: evidenceID},
		)
	}
	// Duplicate identical evidence links are harmless in the source evidence
	// verifier. Deduplicate identical edges here while rejecting disagreements.
	unique := make(map[provenance.Edge]bool, len(g.Edges))
	edges := g.Edges[:0]
	for _, edge := range g.Edges {
		if !unique[edge] {
			unique[edge] = true
			edges = append(edges, edge)
		}
	}
	g.Edges = edges
	if err := g.Validate(); err != nil {
		return provenance.Graph{}, err
	}
	return g, nil
}
