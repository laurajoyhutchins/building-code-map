// Test-only executable evidence: canonical provider → Runtime → ProjectVerifier → HTTP.
// It does not define a second verdict path or read actual regulatory source material.
package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"

	"github.com/laurajoyhutchins/building-code-map/engine"
	"github.com/laurajoyhutchins/building-code-map/httpapi"
)

type syntheticProvider struct{}

func (syntheticProvider) Resolve(_ context.Context, _ engine.NormalizedQuery) (engine.Resolution, error) {
	return engine.Resolution{
		Jurisdiction: engine.JurisdictionContext{
			CountryCode:             "US",
			PrimaryJurisdictionID:   "DEMO-XX",
			PrimaryJurisdictionName: "Fictional demonstration jurisdiction",
		},
		Codes: []engine.CodeResolution{{
			Family:   "building",
			Edition:  "DEMO-2024",
			Status:   engine.StatusPartiallyResolved,
			Basis:    "Synthetic unresolved adoption, never production authority",
			Evidence: []engine.EvidenceRef{{ID: "DEMO-SOURCE", Kind: "synthetic"}},
		}},
		Requirements: []engine.ResolutionRequirement{{
			ID:      "DEMO-OCCUPANCY",
			Kind:    engine.RequirementProjectFact,
			FactKey: "occupancy",
			Prompt:  "What is the fictional occupancy?",
		}},
		DerivedFacts: []engine.DerivedFact{{
			Key:   "jurisdiction",
			Value: "DEMO-XX",
			Basis: "Synthetic location fixture",
		}},
		Provenance: engine.Provenance{
			EngineVersion: "synthetic-schema-fixture",
			SourceCommit:  "DEMO-SOURCE-ONLY",
			BundleID:      "DEMO-BUNDLE",
		},
	}, nil
}

func main() {
	verifier := engine.NewProjectVerifier(engine.NewRuntime(syntheticProvider{}))
	handler := httpapi.NewHandler(verifier, nil)
	request := httptest.NewRequest(http.MethodPost, httpapi.ProjectVerificationPath,
		strings.NewReader(`{"project_id":"DEMO-SCHEMA","address":"100 Demo Plaza","applicability_date":"2026-08-25"}`))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		fmt.Fprintf(os.Stderr, "canonical HTTP response failed with %d: %s\n", response.Code, response.Body.String())
		os.Exit(1)
	}
	if _, err := os.Stdout.Write(response.Body.Bytes()); err != nil {
		fmt.Fprintf(os.Stderr, "write canonical schema fixture: %v\n", err)
		os.Exit(1)
	}
}
