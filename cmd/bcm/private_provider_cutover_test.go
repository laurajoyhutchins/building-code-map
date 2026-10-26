package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"

	"github.com/laurajoyhutchins/building-code-map/engine"
	"github.com/laurajoyhutchins/building-code-map/httpapi"
)

// Only the explicitly configured private provider origin may receive project
// inputs or its service credential. Go otherwise replays a POST for HTTP 307.
func TestPrivateProviderNeverFollowsUpstreamRedirect(t *testing.T) {
	var forwarded atomic.Int32
	other := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		forwarded.Add(1)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer other.Close()

	origin := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != "Bearer provider-only-secret" {
			t.Errorf("initial request missing private authentication: %q", got)
		}
		w.Header().Set("Location", other.URL+"/unexpected-receiver")
		w.WriteHeader(http.StatusTemporaryRedirect)
	}))
	defer origin.Close()

	provider, err := newPrivateProvider(origin.URL, "provider-only-secret")
	if err != nil {
		t.Fatal(err)
	}
	status, err := provider.request(context.Background(), http.MethodPost, "/v1/resolve",
		map[string]string{"address": "100 Synthetic Plaza", "applicability_date": "2026-10-20"}, &struct{}{})
	if status != http.StatusTemporaryRedirect || err == nil {
		t.Fatalf("redirect was not rejected: status=%d err=%v", status, err)
	}
	if forwarded.Load() != 0 {
		t.Fatal("private request was forwarded to an unreviewed origin")
	}
}

// A legacy resolver observation is translated through the public Runtime and
// ProjectVerifier, not returned as an independent private project verdict.
// Unrecognized upstream fields must never enter a public project response.
func TestPrivateProviderCanonicalHTTPExcludesPrivateSourceBody(t *testing.T) {
	const privateMarker = "SYNTHETIC_PRIVATE_SOURCE_BODY_MUST_NOT_ESCAPE"
	sample := strings.Replace(providerSample,
		`"source_ids":["CO-SRC-1"]`,
		`"source_ids":["CO-SRC-1"],"private_source_body":"`+privateMarker+`"`, 1)
	if sample == providerSample {
		t.Fatal("fixture private-body injection failed")
	}
	server := mockLegacyProvider(t, true, sample)
	defer server.Close()

	provider, err := newPrivateProvider(server.URL, "")
	if err != nil {
		t.Fatal(err)
	}
	provider.expectedBundle = "sha256:reviewed"
	handler := handlerWithProvider(provider, provider, "https://laurajoyhutchins.github.io")
	req := httptest.NewRequest(http.MethodPost, httpapi.ProjectVerificationPath,
		strings.NewReader(`{"project_id":"SYNTHETIC-1","address":"100 Main St, Denver, CO","applicability_date":"2026-10-20"}`))
	req.Header.Set("Content-Type", "application/json")
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("canonical project endpoint rejected response: status=%d body=%s", recorder.Code, recorder.Body.String())
	}
	if strings.Contains(recorder.Body.String(), privateMarker) {
		t.Fatal("private source body escaped into public project response")
	}
	var basis engine.ProjectCodeBasis
	if err := json.Unmarshal(recorder.Body.Bytes(), &basis); err != nil {
		t.Fatal(err)
	}
	if basis.ProjectID != "SYNTHETIC-1" || basis.AsOf != "2026-10-20" ||
		basis.Provenance.BundleID != "sha256:reviewed" {
		t.Fatalf("project or runtime identity was lost: %#v", basis)
	}
	if basis.Verdict != engine.ProjectNotVerified || len(basis.CodeSet) != 5 {
		t.Fatalf("legacy source IDs cannot establish a verified project: %#v", basis)
	}
	if len(basis.ExactEvidence) != 0 || len(basis.Evidence) != 1 ||
		basis.Evidence[0].ID != "CO-SRC-1" {
		t.Fatalf("source reference and exact-evidence boundary changed: %#v", basis)
	}
}
