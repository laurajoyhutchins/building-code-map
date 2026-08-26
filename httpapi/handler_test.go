package httpapi

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/laurajoyhutchins/building-code-map/engine"
)

type recordingProvider struct {
	query engine.NormalizedQuery
	err   error
}

func (provider *recordingProvider) Resolve(_ context.Context, query engine.NormalizedQuery) (engine.Resolution, error) {
	provider.query = query
	if provider.err != nil {
		return engine.Resolution{}, provider.err
	}
	return engine.Resolution{
		Jurisdiction: engine.JurisdictionContext{
			CountryCode:             "US",
			PrimaryJurisdictionID:   "DEMO-XX",
			PrimaryJurisdictionName: "Demonstration Jurisdiction",
		},
		Codes: []engine.CodeResolution{{
			Family:  "building",
			Edition: "DEMO-2024",
			Status:  engine.StatusResolved,
			Basis:   "Synthetic HTTP test",
		}},
		Provenance: engine.Provenance{EngineVersion: "demo-http"},
	}, nil
}

func handlerForProvider(provider engine.Provider) http.Handler {
	runtime := engine.NewRuntime(provider)
	return NewHandler(engine.NewProjectVerifier(runtime), nil)
}

func projectRequestBody() string {
	return `{"project_id":"DEMO-HTTP","address":"100 Demo Plaza","applicability_date":"2026-08-25","facts":{"occupancy":"business"}}`
}

func decodeErrorEnvelope(t *testing.T, recorder *httptest.ResponseRecorder) ErrorEnvelope {
	t.Helper()
	var envelope ErrorEnvelope
	if err := json.NewDecoder(recorder.Body).Decode(&envelope); err != nil {
		t.Fatalf("decode error response: %v; body=%q", err, recorder.Body.String())
	}
	return envelope
}

func TestProjectVerificationHTTPUsesCanonicalProjectVerifier(t *testing.T) {
	provider := &recordingProvider{}
	handler := handlerForProvider(provider)
	request := httptest.NewRequest(http.MethodPost, ProjectVerificationPath, strings.NewReader(projectRequestBody()))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body=%s", response.Code, http.StatusOK, response.Body.String())
	}
	var basis engine.ProjectCodeBasis
	if err := json.NewDecoder(response.Body).Decode(&basis); err != nil {
		t.Fatalf("decode basis: %v", err)
	}
	if basis.ProjectID != "DEMO-HTTP" || basis.Verdict != engine.ProjectVerified {
		t.Fatalf("basis = %#v", basis)
	}
	if provider.query.CodeFamily != "" {
		t.Fatalf("HTTP adapter injected code family %q", provider.query.CodeFamily)
	}
	if provider.query.ProjectFacts["occupancy"] != "business" {
		t.Fatalf("project facts = %#v", provider.query.ProjectFacts)
	}
}

func TestProjectVerificationHTTPRejectsWrongMethod(t *testing.T) {
	handler := handlerForProvider(&recordingProvider{})
	request := httptest.NewRequest(http.MethodGet, ProjectVerificationPath, nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusMethodNotAllowed {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusMethodNotAllowed)
	}
	if got := response.Header().Get("Allow"); got != http.MethodPost {
		t.Fatalf("Allow = %q, want POST", got)
	}
	if got := decodeErrorEnvelope(t, response).Error.Code; got != "method_not_allowed" {
		t.Fatalf("error code = %q", got)
	}
}

func TestProjectVerificationHTTPRejectsInvalidJSONShapes(t *testing.T) {
	tests := map[string]string{
		"unknown field": `{"project_id":"DEMO","address":"100 Demo Plaza","applicability_date":"2026-08-25","surprise":true}`,
		"malformed":     `{"project_id":`,
		"trailing value": projectRequestBody() + ` {}`,
	}
	for name, body := range tests {
		t.Run(name, func(t *testing.T) {
			handler := handlerForProvider(&recordingProvider{})
			request := httptest.NewRequest(http.MethodPost, ProjectVerificationPath, strings.NewReader(body))
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()

			handler.ServeHTTP(response, request)

			if response.Code != http.StatusBadRequest {
				t.Fatalf("status = %d, want %d; body=%s", response.Code, http.StatusBadRequest, response.Body.String())
			}
			if got := decodeErrorEnvelope(t, response).Error.Code; got != "invalid_json" {
				t.Fatalf("error code = %q", got)
			}
		})
	}
}

func TestProjectVerificationHTTPRejectsOversizedBody(t *testing.T) {
	handler := handlerForProvider(&recordingProvider{})
	body := `{"project_id":"DEMO","address":"` + strings.Repeat("x", int(MaxRequestBodyBytes)+1024) + `","applicability_date":"2026-08-25"}`
	request := httptest.NewRequest(http.MethodPost, ProjectVerificationPath, strings.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusRequestEntityTooLarge {
		t.Fatalf("status = %d, want %d; body=%s", response.Code, http.StatusRequestEntityTooLarge, response.Body.String())
	}
	if got := decodeErrorEnvelope(t, response).Error.Code; got != "request_too_large" {
		t.Fatalf("error code = %q", got)
	}
}

func TestProjectVerificationHTTPPreservesTypedEngineError(t *testing.T) {
	provider := &recordingProvider{err: engine.EngineError{
		Code:      engine.ErrorRegulatoryCatalogUnavailable,
		Message:   "catalog unavailable",
		Retryable: true,
	}}
	handler := handlerForProvider(provider)
	request := httptest.NewRequest(http.MethodPost, ProjectVerificationPath, strings.NewReader(projectRequestBody()))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusServiceUnavailable {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusServiceUnavailable)
	}
	envelope := decodeErrorEnvelope(t, response)
	if envelope.Error.Code != string(engine.ErrorRegulatoryCatalogUnavailable) || !envelope.Error.Retryable {
		t.Fatalf("error = %#v", envelope.Error)
	}
}

func TestProjectVerificationHTTPRejectsNonJSONContentType(t *testing.T) {
	handler := handlerForProvider(&recordingProvider{})
	request := httptest.NewRequest(http.MethodPost, ProjectVerificationPath, strings.NewReader(projectRequestBody()))
	request.Header.Set("Content-Type", "text/plain")
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusUnsupportedMediaType {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusUnsupportedMediaType)
	}
	if got := decodeErrorEnvelope(t, response).Error.Code; got != "unsupported_media_type" {
		t.Fatalf("error code = %q", got)
	}
}
