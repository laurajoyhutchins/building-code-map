package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/laurajoyhutchins/building-code-map/utilitycontext"
)

func fixture(t *testing.T) *utilitycontext.Catalog {
	t.Helper()
	body := []byte(`{"format":"bcm/eia-county-candidates/v1","release_id":"eia861-2025-synthetic-v1","reporting_year":2025,"source_sha256":"` + strings.Repeat("a", 64) + `","sqlite_sha256":"` + strings.Repeat("b", 64) + `","source_url":"https://www.eia.gov/electricity/data/eia861/","scope":"EIA-reported county candidates only; no verified serving utility","records":[{"state":"CO","county":"Denver","eia_id":15466,"reported_name":"Synthetic utility","source_sheet":"Counties_States","source_row":8,"activities":[]}]}`)
	path := filepath.Join(t.TempDir(), "view.json")
	if err := os.WriteFile(path, body, 0600); err != nil {
		t.Fatal(err)
	}
	c, err := utilitycontext.Open(path, fmt.Sprintf("%x", sha256.Sum256(body)))
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func TestCountyOnlyNoRegulatoryRouteOrCrossOriginAccess(t *testing.T) {
	api := handler(fixture(t))
	cases := []struct {
		method, url, origin string
		status              int
	}{
		{"GET", "/v1/utility-context?state=CO&county=Denver", "", 200},
		{"GET", "/v1/utility-context?state=CO&county=Larimer", "", 200},
		{"GET", "/v1/utility-context?state=CO&county=Denver&all=true", "", 400},
		{"GET", "/v1/project-code-basis", "", 404},
		{"POST", "/v1/project-utility-context", "", 404},
		{"POST", "/v1/project-code-basis", "", 404},
		{"GET", "/healthz", "", 200},
		{"OPTIONS", "/v1/utility-context", "https://untrusted.example", 403},
		{"OPTIONS", "/v1/utility-context", "https://laurajoyhutchins.github.io", 204},
	}
	for _, tc := range cases {
		req := httptest.NewRequest(tc.method, tc.url, nil)
		if tc.origin != "" {
			req.Header.Set("Origin", tc.origin)
		}
		out := httptest.NewRecorder()
		api.ServeHTTP(out, req)
		if out.Code != tc.status {
			t.Errorf("%s %s: got %d want %d", tc.method, tc.url, out.Code, tc.status)
		}
		if tc.origin == "https://untrusted.example" && out.Header().Get("Access-Control-Allow-Origin") != "" {
			t.Fatal("untrusted origin allowed")
		}
		if tc.url == "/healthz" && out.Code == 200 {
			var v map[string]string
			if err := json.Unmarshal(out.Body.Bytes(), &v); err != nil {
				t.Fatal(err)
			}
			if v["capability"] != "county_utility_candidates_only" {
				t.Fatalf("wrong capability: %v", v)
			}
		}
	}
}

func TestServerRefusesUnpinnedCatalog(t *testing.T) {
	t.Setenv("BCM_UTILITY_CONTEXT_VIEW_PATH", "")
	t.Setenv("BCM_UTILITY_CONTEXT_VIEW_SHA256", "")
	if err := run(); err == nil {
		t.Fatal("started server without view pin")
	}
}
