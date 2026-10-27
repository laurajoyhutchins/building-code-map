package httpapi

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/laurajoyhutchins/building-code-map/engine"
)

type malformedProjectProvider struct {
	derivedFacts []engine.DerivedFact
	evidence     []engine.EvidenceRef
}

func (provider malformedProjectProvider) Resolve(_ context.Context, _ engine.NormalizedQuery) (engine.Resolution, error) {
	return engine.Resolution{
		Jurisdiction: engine.JurisdictionContext{
			CountryCode:             "US",
			PrimaryJurisdictionID:   "DEMO-XX",
			PrimaryJurisdictionName: "Synthetic jurisdiction",
		},
		Codes: []engine.CodeResolution{{
			Family:   "building",
			Status:   engine.StatusPartiallyResolved,
			Evidence: provider.evidence,
		}},
		DerivedFacts: provider.derivedFacts,
	}, nil
}

func TestProjectHTTPRejectsMalformedProviderFieldsBeforeSuccess(t *testing.T) {
	tests := []struct {
		name     string
		provider malformedProjectProvider
	}{
		{name: "missing derived key", provider: malformedProjectProvider{
			derivedFacts: []engine.DerivedFact{{Value: "DEMO-XX"}},
		}},
		{name: "missing derived value", provider: malformedProjectProvider{
			derivedFacts: []engine.DerivedFact{{Key: "jurisdiction"}},
		}},
		{name: "missing ordinary evidence id", provider: malformedProjectProvider{
			evidence: []engine.EvidenceRef{{Kind: "synthetic"}},
		}},
		{name: "missing ordinary evidence kind", provider: malformedProjectProvider{
			evidence: []engine.EvidenceRef{{ID: "DEMO-SOURCE"}},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handler := handlerForProvider(tt.provider)
			request := httptest.NewRequest(http.MethodPost, ProjectVerificationPath, strings.NewReader(projectRequestBody()))
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			handler.ServeHTTP(response, request)
			if response.Code == http.StatusOK {
				t.Fatalf("malformed provider data incorrectly returned HTTP 200: %s", response.Body.String())
			}
			envelope := decodeErrorEnvelope(t, response)
			if envelope.Error.Code != string(engine.ErrorDataBundleInvalid) {
				t.Fatalf("error = %s, want data_bundle_invalid", envelope.Error.Code)
			}
		})
	}
}
