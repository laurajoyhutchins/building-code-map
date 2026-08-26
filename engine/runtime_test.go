package engine

import (
	"context"
	"errors"
	"testing"
)

type syntheticProvider struct{ codes []CodeResolution }

func (provider syntheticProvider) Resolve(_ context.Context, _ NormalizedQuery) (Resolution, error) {
	return Resolution{
		Jurisdiction: JurisdictionContext{
			CountryCode: "US", PrimaryJurisdictionID: "DEMO-XX", PrimaryJurisdictionName: "Demonstration Jurisdiction",
		},
		Codes:      provider.codes,
		Provenance: Provenance{EngineVersion: "demo", SourceCommit: "DEMO-COMMIT", BundleID: "DEMO-BUNDLE"},
	}, nil
}

func TestRuntimeExecutesSyntheticQuery(t *testing.T) {
	runtime := NewRuntime(syntheticProvider{codes: []CodeResolution{
		{Family: "building", Edition: "DEMO-IBC-2024", Status: StatusResolved, Basis: "DEMO ordinance"},
		{Family: "fire", Status: StatusUnsupported},
		{Family: "energy", Status: StatusInsufficientEvidence, Unresolved: []string{"DEMO local record"}},
	}})

	result, err := runtime.Resolve(context.Background(), Query{Address: "  100 Demo Plaza  ", ApplicabilityDate: "2026-08-25"})
	if err != nil {
		t.Fatalf("Resolve() error = %v", err)
	}
	if result.Query.Address != "100 Demo Plaza" {
		t.Fatalf("normalized address = %q", result.Query.Address)
	}
	if result.Jurisdiction.PrimaryJurisdictionID != "DEMO-XX" {
		t.Fatalf("primary jurisdiction = %q", result.Jurisdiction.PrimaryJurisdictionID)
	}
	if got := result.Codes[1].Status; got != StatusUnsupported {
		t.Fatalf("unsupported status = %q", got)
	}
	if got := result.Codes[2].Status; got != StatusInsufficientEvidence {
		t.Fatalf("insufficient evidence status = %q", got)
	}
}

func TestNotApplicableRequiresAffirmativeBasis(t *testing.T) {
	runtime := NewRuntime(syntheticProvider{codes: []CodeResolution{{Family: "fire", Status: StatusNotApplicable}}})
	_, err := runtime.Resolve(context.Background(), Query{Address: "100 Demo Plaza", ApplicabilityDate: "2026-08-25"})
	var engineErr EngineError
	if !errors.As(err, &engineErr) || engineErr.Code != ErrorDataBundleInvalid {
		t.Fatalf("Resolve() error = %#v, want data_bundle_invalid", err)
	}
}

func TestRuntimeRejectsMalformedExactEvidence(t *testing.T) {
	link := demoEvidenceLink("DEMO-LINK", EvidenceEstablishes)
	link.Artifact.SHA256 = ""
	runtime := NewRuntime(syntheticProvider{codes: []CodeResolution{{
		Family: "building", Edition: "DEMO-IBC-2024", Status: StatusResolved, Basis: "DEMO ordinance", ExactEvidence: []EvidenceLink{link},
	}}})

	_, err := runtime.Resolve(context.Background(), Query{Address: "100 Demo Plaza", ApplicabilityDate: "2026-08-25"})
	var engineErr EngineError
	if !errors.As(err, &engineErr) || engineErr.Code != ErrorDataBundleInvalid {
		t.Fatalf("Resolve() error = %#v, want data_bundle_invalid", err)
	}
}

func TestRuntimePreservesValidExactEvidence(t *testing.T) {
	link := demoEvidenceLink("DEMO-LINK", EvidenceEstablishes)
	runtime := NewRuntime(syntheticProvider{codes: []CodeResolution{{
		Family: "building", Edition: "DEMO-IBC-2024", Status: StatusResolved, Basis: "DEMO ordinance", ExactEvidence: []EvidenceLink{link},
	}}})

	result, err := runtime.Resolve(context.Background(), Query{Address: "100 Demo Plaza", ApplicabilityDate: "2026-08-25"})
	if err != nil {
		t.Fatalf("Resolve() error = %v", err)
	}
	if len(result.Codes[0].ExactEvidence) != 1 || result.Codes[0].ExactEvidence[0].ID != link.ID {
		t.Fatalf("exact evidence = %#v", result.Codes[0].ExactEvidence)
	}
}

func TestStatusVocabularyKeepsUnsupportedNotApplicableAndInsufficientEvidenceDistinct(t *testing.T) {
	statuses := []ResolutionStatus{StatusUnsupported, StatusNotApplicable, StatusInsufficientEvidence}
	if statuses[0] == statuses[1] || statuses[0] == statuses[2] || statuses[1] == statuses[2] {
		t.Fatalf("status vocabulary collapsed: %#v", statuses)
	}
}
