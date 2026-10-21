package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/laurajoyhutchins/building-code-map/engine"
	"github.com/laurajoyhutchins/building-code-map/httpapi"
)

const providerSample = `{
  "schema_version":"1",
  "resolution":{
    "status":"local_record_required",
    "applicability_date":"2026-10-20",
    "geography":{
      "jurisdiction_id":"08",
      "jurisdiction_name":"Colorado",
      "census_state_equivalent_fips":"08"
    },
    "adoptions":[{
      "code_family":"electrical",
      "code_name":"National Electrical Code",
      "status":"current",
      "verification":{"status":"verified"},
      "source_ids":["CO-SRC-1"]
    }],
    "required_local_records":["Check municipal amendments"]
  },
  "provenance":{
    "engine_version":"0.1.0",
    "source_commit":"0123456789abcdef0123456789abcdef01234567",
    "bundle_manifest_digest":"sha256:reviewed"
  }
}`

func sampleForFamily(t *testing.T, source, family string) []byte {
	t.Helper()
	var response map[string]any
	if err := json.Unmarshal([]byte(source), &response); err != nil {
		t.Fatal(err)
	}
	res := response["resolution"].(map[string]any)
	if family != "electrical" {
		res["adoptions"] = []any{}
		res["status"] = "insufficient_evidence"
		res["required_local_records"] = []any{}
	}
	b, err := json.Marshal(response)
	if err != nil {
		t.Fatal(err)
	}
	return b
}

func mockLegacyProvider(t *testing.T, ready bool, result string) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/v1/readiness":
			w.Header().Set("Content-Type", "application/json")
			if !ready {
				_, _ = w.Write([]byte(`{"status":"not_ready","capabilities":{"address_lookup":{"status":"unavailable"}}}`))
				return
			}
			_, _ = w.Write([]byte(`{"status":"ok","capabilities":{"address_lookup":{"status":"available"}}}`))
		case "/v1/bundle":
			_, _ = w.Write([]byte(`{"bundle_manifest_digest":"sha256:reviewed"}`))
		case "/v1/resolve":
			if r.Method != http.MethodPost {
				t.Errorf("method = %s", r.Method)
			}
			var query map[string]any
			if err := json.NewDecoder(r.Body).Decode(&query); err != nil {
				t.Errorf("decode query: %v", err)
			}
			if query["applicability_date"] != "2026-10-20" || query["address"] != "100 Main St, Denver, CO" {
				t.Errorf("unexpected query: %#v", query)
			}
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write(sampleForFamily(t, result, query["code_family"].(string)))
		default:
			http.NotFound(w, r)
		}
	}))
}

func TestPrivateProviderUsesOneEngineAndPreservesUncertainty(t *testing.T) {
	server := mockLegacyProvider(t, true, providerSample)
	defer server.Close()
	provider, err := newPrivateProvider(server.URL, "")
	if err != nil {
		t.Fatal(err)
	}
	result, err := engine.NewProjectVerifier(engine.NewRuntime(provider)).VerifyProject(context.Background(), engine.ProjectRequest{
		ProjectID: "SITE-123", Address: "100 Main St, Denver, CO", ApplicabilityDate: "2026-10-20",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Verdict != engine.ProjectNotVerified || result.JurisdictionID != "08" || result.AsOf != "2026-10-20" {
		t.Fatalf("unexpected project basis: %#v", result)
	}
	if len(result.CodeSet) != 5 {
		t.Fatalf("incomplete code family set: %#v", result.CodeSet)
	}
	var electricalFound bool
	for _, code := range result.CodeSet {
		if code.Family == "electrical" {
			electricalFound = true
			if code.Status != engine.StatusLocalRecordRequired || code.Edition != "National Electrical Code" {
				t.Fatalf("wrong adoption translation: %#v", code)
			}
		} else if code.Status != engine.StatusInsufficientEvidence {
			t.Fatalf("unsupported code family was promoted: %#v", code)
		}
	}
	if !electricalFound {
		t.Fatalf("missing electrical code from %#v", result.CodeSet)
	}
	if len(result.Evidence) != 1 || result.Evidence[0].ID != "CO-SRC-1" || result.Provenance.BundleID != "sha256:reviewed" {
		t.Fatalf("missing real source identity: %#v", result)
	}
	var foundLocalRecord bool
	for _, unresolved := range result.Unresolved {
		if unresolved.Message == "electrical: Check municipal amendments" {
			foundLocalRecord = true
		}
	}
	if !foundLocalRecord || len(result.Unresolved) < 5 {
		t.Fatalf("code family gaps or local record not preserved: %#v", result.Unresolved)
	}
}

func TestUnavailablePrivateDataFailsClosed(t *testing.T) {
	server := mockLegacyProvider(t, false, providerSample)
	defer server.Close()
	provider, err := newPrivateProvider(server.URL, "")
	if err != nil {
		t.Fatal(err)
	}
	_, err = engine.NewRuntime(provider).Resolve(context.Background(), engine.Query{
		Address: "100 Main St, Denver, CO", ApplicabilityDate: "2026-10-20",
	})
	if err == nil || !strings.Contains(err.Error(), "not ready") {
		t.Fatalf("expected fail-closed error, got %v", err)
	}
}

func TestMissingBundleIdentityFailsClosed(t *testing.T) {
	invalid := strings.Replace(providerSample, "sha256:reviewed", "", 1)
	server := mockLegacyProvider(t, true, invalid)
	defer server.Close()
	provider, err := newPrivateProvider(server.URL, "")
	if err != nil {
		t.Fatal(err)
	}
	_, err = engine.NewRuntime(provider).Resolve(context.Background(), engine.Query{
		Address: "100 Main St, Denver, CO", ApplicabilityDate: "2026-10-20",
	})
	if err == nil || !strings.Contains(err.Error(), "incomplete identity") {
		t.Fatalf("expected data identity rejection, got %v", err)
	}
}

func TestProviderURLRejectsUnencryptedNonLocalhost(t *testing.T) {
	if _, err := newPrivateProvider("http://example.com", ""); err == nil {
		t.Fatal("non-HTTPS upstream accepted")
	}
	if _, err := newPrivateProvider("https://example.com/untrusted/path", ""); err == nil {
		t.Fatal("path-based upstream accepted")
	}
}

func TestPublicBrowserOriginIsStrictlyScoped(t *testing.T) {
	server := mockLegacyProvider(t, true, providerSample)
	defer server.Close()
	provider, err := newPrivateProvider(server.URL, "")
	if err != nil {
		t.Fatal(err)
	}
	handler := handlerWithProvider(provider, provider, "https://laurajoyhutchins.github.io")
	for _, origin := range []string{"https://laurajoyhutchins.github.io", "https://untrusted.example"} {
		req := httptest.NewRequest(http.MethodOptions, httpapi.ProjectVerificationPath, nil)
		req.Header.Set("Origin", origin)
		response := httptest.NewRecorder()
		handler.ServeHTTP(response, req)
		if origin == "https://laurajoyhutchins.github.io" {
			if response.Code != http.StatusNoContent || response.Header().Get("Access-Control-Allow-Origin") != origin {
				t.Fatalf("allowed preflight rejected: %d %#v", response.Code, response.Header())
			}
		} else if response.Header().Get("Access-Control-Allow-Origin") != "" {
			t.Fatal("untrusted origin was allowed")
		}
	}
}

func TestPrivateProviderUsesWorkloadIdentityForEveryRequest(t *testing.T) {
	calls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer signed-test-identity" {
			t.Errorf("private request was not authenticated: %s", r.URL.Path)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		calls++
		switch r.URL.Path {
		case "/v1/readiness":
			_, _ = w.Write([]byte(`{"status":"ok","capabilities":{"address_lookup":{"status":"available"}}}`))
		case "/v1/bundle":
			_, _ = w.Write([]byte(`{"bundle_manifest_digest":"sha256:reviewed"}`))
		case "/v1/resolve":
			var request struct {
				CodeFamily string `json:"code_family"`
			}
			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
				t.Error(err)
			}
			_, _ = w.Write(sampleForFamily(t, providerSample, request.CodeFamily))
		default:
			http.NotFound(w, r)
		}
	}))
	defer server.Close()
	provider, err := newPrivateProvider(server.URL, "")
	if err != nil {
		t.Fatal(err)
	}
	provider.expectedBundle = "sha256:reviewed"
	provider.identityToken = func(ctx context.Context, audience string) (string, error) {
		if audience != server.URL {
			t.Errorf("wrong token audience: %s", audience)
		}
		return "signed-test-identity", nil
	}
	_, err = engine.NewProjectVerifier(engine.NewRuntime(provider)).VerifyProject(
		context.Background(),
		engine.ProjectRequest{ProjectID: "SITE-123", Address: "100 Main St, Denver, CO", ApplicabilityDate: "2026-10-20"},
	)
	if err != nil {
		t.Fatal(err)
	}
	if calls != 7 {
		t.Fatalf("wanted 7 authenticated calls, got %d", calls)
	}
}

func TestPrivateProviderResponseIsCompleteAndSizeBounded(t *testing.T) {
	const valid = `{"ok":true}`
	tests := []struct {
		name      string
		body      string
		wantError string
	}{
		{name: "single JSON value", body: valid},
		{name: "trailing JSON whitespace", body: valid + " \n\t"},
		{name: "exactly at size limit", body: valid + strings.Repeat(" ", maxPrivateProviderResponseBytes-len(valid))},
		{name: "second JSON object", body: valid + `{"other":true}`, wantError: "trailing content"},
		{name: "trailing non-JSON bytes", body: valid + " garbage", wantError: "trailing content"},
		{name: "oversized trailing whitespace", body: valid + strings.Repeat(" ", maxPrivateProviderResponseBytes-len(valid)+1), wantError: "exceeds 1 MiB"},
		{name: "oversized first JSON value", body: `{"blob":"` + strings.Repeat("x", maxPrivateProviderResponseBytes) + `"}`, wantError: "exceeds 1 MiB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(tt.body))
			}))
			defer server.Close()
			provider, err := newPrivateProvider(server.URL, "")
			if err != nil {
				t.Fatal(err)
			}
			var result struct {
				OK bool `json:"ok"`
			}
			status, err := provider.request(context.Background(), http.MethodGet, "/v1/test", nil, &result)
			if tt.wantError != "" {
				if err == nil || !strings.Contains(err.Error(), tt.wantError) {
					t.Fatalf("wanted %q error, got status %d, err %v", tt.wantError, status, err)
				}
				return
			}
			if err != nil || status != http.StatusOK || !result.OK {
				t.Fatalf("valid response rejected: status %d, result %#v, err %v", status, result, err)
			}
		})
	}
}
