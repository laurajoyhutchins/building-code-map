package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"building-code-map/backend/internal/snapshotmanifest"
)

func TestSnapshotActivatePromotesVerifiedBuildCandidate(t *testing.T) {
	dir := t.TempDir()
	candidate := filepath.Join(dir, "candidate.sqlite")
	active := filepath.Join(dir, "active.sqlite")
	if err := os.WriteFile(candidate, []byte("candidate"), 0o600); err != nil {
		t.Fatal(err)
	}
	manifest := snapshotmanifest.Manifest{
		SchemaVersion: snapshotmanifest.SchemaVersion,
		SnapshotID:    "boundary-candidate-001",
		Kind:          snapshotmanifest.KindBoundary,
		Sources: []snapshotmanifest.SourceArtifact{
			{
				Publisher:            "fixture publisher",
				Product:              "fixture product",
				Vintage:              "2026-08-08",
				Locator:              "fixture://boundary",
				SHA256:               strings.Repeat("a", 64),
				RetrievedAt:          "2026-08-08T15:00:00Z",
				LicenseReviewStatus:  "reviewed",
				RedistributionStatus: "repository-owned",
			},
		},
		Builder: snapshotmanifest.Builder{
			Tool:     "fixture-builder",
			Version:  "1.0",
			Revision: "fixture-revision",
			BuiltAt:  "2026-08-08T15:00:00Z",
		},
		RecordCounts: snapshotmanifest.RecordCounts{Accepted: 1},
		IntegrityChecks: []snapshotmanifest.IntegrityCheck{
			{Name: "fixture", Status: "passed"},
		},
	}
	if _, err := snapshotmanifest.FinalizeAndWrite(candidate, manifest); err != nil {
		t.Fatal(err)
	}

	code := runSnapshot([]string{
		"activate",
		"--kind", "boundary",
		"--candidate", candidate,
		"--active", active,
		"--activated-at", "2026-08-08T15:30:00Z",
	})
	if code != exitSuccess {
		t.Fatalf("runSnapshot activate exit=%d", code)
	}

	verified, err := snapshotmanifest.LoadAndVerify(active, snapshotmanifest.KindBoundary)
	if err != nil {
		t.Fatal(err)
	}
	if verified.Manifest.SnapshotID != "boundary-candidate-001" {
		t.Fatalf("active snapshot=%q", verified.Manifest.SnapshotID)
	}
	if verified.Receipt.ActivatedAt != "2026-08-08T15:30:00Z" {
		t.Fatalf("activated_at=%q", verified.Receipt.ActivatedAt)
	}
}
