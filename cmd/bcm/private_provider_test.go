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
			_, _ = w.Write([]byte(result))
		default:
			http.NotFound(w, r)
		}
	}))
}

func TestPrivateProviderUsesOneEngineAndPreservesUncertainty(t *testing.T) {
	server := mockLegacyProvider(t, true, providerSample)
	defer server.Close()
	provider, err := newPrivateProvider(server.URL, "")
	if err != nil { t.Fatal(err) }
	result, err := engine.NewProjectVerifier(engine.NewRuntime(provider)).VerifyProject(context.Background(), engine.ProjectRequest{
		ProjectID:"SITE-123",Address:"100 Main St, Denver, CO",ApplicabilityDate:"2026-10-20",
	})
	if err != nil { t.Fatal(err) }
	if result.Verdict != engine.ProjectConditional || result.JurisdictionID != "08" || result.AsOf != "2026-10-20" {
		t.Fatalf("unexpected project basis: %#v", result)
	}
	if len(result.CodeSet)!=1 || result.CodeSet[0].Family!="electrical" || result.CodeSet[0].Status!=engine.StatusLocalRecordRequired {
		t.Fatalf("adoption incorrectly mapped: %#v", result.CodeSet)
	}
	if len(result.Evidence)!=1 || result.Evidence[0].ID!="CO-SRC-1" || result.Provenance.BundleID!="sha256:reviewed" {
		t.Fatalf("missing real source identity: %#v", result)
	}
	if len(result.Unresolved)!=1 || result.Unresolved[0].Message!="Check municipal amendments" {
		t.Fatalf("local record not preserved: %#v", result.Unresolved)
	}
}

func TestUnavailablePrivateDataFailsClosed(t *testing.T) {
	server:=mockLegacyProvider(t,false,providerSample)
	defer server.Close()
	provider,err:=newPrivateProvider(server.URL,"")
	if err!=nil {t.Fatal(err)}
	_,err=engine.NewRuntime(provider).Resolve(context.Background(),engine.Query{
		Address:"100 Main St, Denver, CO",ApplicabilityDate:"2026-10-20",
	})
	if err==nil || !strings.Contains(err.Error(),"not ready") {t.Fatalf("expected fail-closed error, got %v",err)}
}

func TestMissingBundleIdentityFailsClosed(t *testing.T) {
	invalid:=strings.Replace(providerSample, "sha256:reviewed", "", 1)
	server:=mockLegacyProvider(t,true,invalid)
	defer server.Close()
	provider,err:=newPrivateProvider(server.URL,"")
	if err!=nil {t.Fatal(err)}
	_,err=engine.NewRuntime(provider).Resolve(context.Background(),engine.Query{
		Address:"100 Main St, Denver, CO",ApplicabilityDate:"2026-10-20",
	})
	if err==nil || !strings.Contains(err.Error(),"incomplete identity") {t.Fatalf("expected data identity rejection, got %v",err)}
}

func TestProviderURLRejectsUnencryptedNonLocalhost(t *testing.T) {
	if _,err:=newPrivateProvider("http://example.com",""); err==nil {t.Fatal("non-HTTPS upstream accepted")}
	if _,err:=newPrivateProvider("https://example.com/untrusted/path",""); err==nil {t.Fatal("path-based upstream accepted")}
}

func TestPublicBrowserOriginIsStrictlyScoped(t *testing.T) {
	server:=mockLegacyProvider(t,true,providerSample)
	defer server.Close()
	provider,err:=newPrivateProvider(server.URL,"")
	if err!=nil {t.Fatal(err)}
	handler:=handlerWithProvider(provider,provider,"https://laurajoyhutchins.github.io")
	for _,origin:=range []string{"https://laurajoyhutchins.github.io","https://untrusted.example"} {
		req:=httptest.NewRequest(http.MethodOptions,httpapi.ProjectVerificationPath,nil)
		req.Header.Set("Origin",origin)
		response:=httptest.NewRecorder()
		handler.ServeHTTP(response,req)
		if origin=="https://laurajoyhutchins.github.io" {
			if response.Code!=http.StatusNoContent || response.Header().Get("Access-Control-Allow-Origin")!=origin {
				t.Fatalf("allowed preflight rejected: %d %#v",response.Code,response.Header())
			}
		} else if response.Header().Get("Access-Control-Allow-Origin")!="" {
			t.Fatal("untrusted origin was allowed")
		}
	}
}
