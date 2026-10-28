package engine

import (
	"strings"
	"testing"

	"github.com/laurajoyhutchins/building-code-map/provenance"
)

func TestProjectClaimPROVRetainsAdverseEvidenceWithoutVerdictPromotion(t *testing.T) {
	establishing := demoEvidenceLink("DEMO-LINK-1", EvidenceEstablishes)
	contradicting := demoEvidenceLink("DEMO-LINK-2", EvidenceContradicts)
	g, err := ProjectClaimPROV(demoClaim(establishing, contradicting))
	if err != nil {
		t.Fatal(err)
	}
	if len(g.Nodes) != 6 || len(g.Edges) != 6 {
		t.Fatalf("graph identities = %d nodes / %d links, want 6/6", len(g.Nodes), len(g.Edges))
	}
	var haveEstablishing, haveContradicting bool
	for _, n := range g.Nodes {
		if n.Kind == provenance.Activity || n.Kind == provenance.Agent {
			t.Fatal("invented review activity or agent")
		}
		if n.EvidenceRole == "establishes" {
			haveEstablishing = true
		}
		if n.EvidenceRole == "contradicts" {
			haveContradicting = true
		}
	}
	if !haveEstablishing || !haveContradicting {
		t.Fatal("lost exact evidence roles")
	}
	for _, edge := range g.Edges {
		if edge.Predicate == "prov:wasGeneratedBy" {
			t.Fatal("invented generation/review")
		}
		if edge.Predicate == "prov:wasDerivedFrom" {
			claimID, _ := provenance.StableID("claim", "DEMO-CLAIM")
			if edge.From == claimID {
				t.Fatal("claim silently derived from possibly contradictory evidence")
			}
		}
	}
	data, err := g.JSONLD()
	if err != nil || !strings.Contains(string(data), "bcm:hasEvidenceLink") {
		t.Fatalf("invalid JSON-LD evidence graph: %v / %s", err, data)
	}
}

func TestProjectClaimPROVDeduplicatesIdenticalEvidence(t *testing.T) {
	same := demoEvidenceLink("SAME", EvidenceEstablishes)
	g, err := ProjectClaimPROV(demoClaim(same, same))
	if err != nil || len(g.Nodes) != 5 || len(g.Edges) != 4 {
		t.Fatalf("deduplication = %d/%d, err %v", len(g.Nodes), len(g.Edges), err)
	}
}

func TestProjectClaimPROVRejectsInvalidEvidence(t *testing.T) {
	base := demoEvidenceLink("LINK", EvidenceEstablishes)
	cases := []struct {
		name  string
		claim EvidenceClaim
	}{
		{"no links", demoClaim()},
		{"invalid date", func() EvidenceClaim { c := demoClaim(base); c.AsOf = "tomorrow"; return c }()},
		{"wrong jurisdiction", func() EvidenceClaim { c := demoClaim(base); c.Links[0].Document.JurisdictionID = "US-CO"; return c }()},
		{"wrong artifact", func() EvidenceClaim { c := demoClaim(base); c.Links[0].Artifact.DocumentID = "OTHER"; return c }()},
		{"unknown role", func() EvidenceClaim { c := demoClaim(base); c.Links[0].Role = "approved"; return c }()},
		{"invalid sha", func() EvidenceClaim { c := demoClaim(base); c.Links[0].Artifact.SHA256 = "bad"; return c }()},
		{"same ID differing role", func() EvidenceClaim {
			c := demoClaim(base)
			different := demoEvidenceLink("LINK", EvidenceContradicts)
			c.Links = append(c.Links, different)
			return c
		}()},
		{"same artifact ID different hash", func() EvidenceClaim {
			c := demoClaim(base)
			other := demoEvidenceLink("OTHER", EvidenceEstablishes)
			other.Artifact.SHA256 = strings.Repeat("c", 64)
			c.Links = append(c.Links, other)
			return c
		}()},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := ProjectClaimPROV(tc.claim); err == nil {
				t.Fatal("invalid evidence projected")
			}
		})
	}
}

func TestProjectClaimPROVRejectsSourceRebindingAcrossLinks(t *testing.T) {
	base := demoEvidenceLink("DEMO-FIRST", EvidenceEstablishes)
	t.Run("same artifact identity has two document owners", func(t *testing.T) {
		second := demoEvidenceLink("DEMO-SECOND", EvidenceCorroborates)
		second.Document.ID = "DIFFERENT-DOC"
		second.Artifact.DocumentID = second.Document.ID
		second.Anchor.ID = "DIFFERENT-ANCHOR"
		if _, err := ProjectClaimPROV(demoClaim(base, second)); err == nil {
			t.Fatal("one artifact rebound to a second document")
		}
	})
	t.Run("same anchor identity has two artifact owners", func(t *testing.T) {
		second := demoEvidenceLink("DEMO-SECOND", EvidenceCorroborates)
		second.Artifact.ID = "DIFFERENT-ARTIFACT"
		second.Anchor.ArtifactID = second.Artifact.ID
		if _, err := ProjectClaimPROV(demoClaim(base, second)); err == nil {
			t.Fatal("one text anchor rebound to a second artifact")
		}
	})
}
