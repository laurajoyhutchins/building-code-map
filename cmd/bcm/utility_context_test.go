package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/laurajoyhutchins/building-code-map/engine"
	"github.com/laurajoyhutchins/building-code-map/utilitycontext"
)

func testUtilityContext(t *testing.T) *utilitycontext.Catalog {
	t.Helper()
	data := []byte(`{"format":"bcm/eia-county-candidates/v1","release_id":"eia861-2025-synthetic-v1","reporting_year":2025,"source_sha256":"` + strings.Repeat("a", 64) + `","sqlite_sha256":"` + strings.Repeat("b", 64) + `","source_url":"https://www.eia.gov/electricity/data/eia861/","scope":"EIA-reported county candidates only; no verified serving utility","records":[{"state":"CO","county":"Denver","eia_id":15466,"reported_name":"Test Utility","source_sheet":"Counties_States","source_row":7,"activities":[]}]}`)
	path := filepath.Join(t.TempDir(), "context.json")
	if err := os.WriteFile(path, data, 0600); err != nil {
		t.Fatal(err)
	}
	catalog, err := utilitycontext.Open(path, fmt.Sprintf("%x", sha256.Sum256(data)))
	if err != nil {
		t.Fatal(err)
	}
	return catalog
}

func TestOptionalUtilityContextNeverChangesCodeBasis(t *testing.T) {
	withContext := handlerWithUtilityContext(demoProvider{}, nil, "", testUtilityContext(t))
	without := handlerWithProvider(demoProvider{}, nil, "")
	body := `{"project_id":"DEMO-001","address":"100 Demo Plaza, DEMO-XX","applicability_date":"2026-08-25"}`
	check := func(handler http.Handler) engine.ProjectCodeBasis {
		recorder := httptest.NewRecorder()
		handler.ServeHTTP(recorder, httptest.NewRequest(http.MethodPost, "/v1/project-code-basis", strings.NewReader(body)))
		if recorder.Code != 200 {
			t.Fatalf("code basis status = %d", recorder.Code)
		}
		var basis engine.ProjectCodeBasis
		if err := json.Unmarshal(recorder.Body.Bytes(), &basis); err != nil {
			t.Fatal(err)
		}
		return basis
	}
	before, after := check(without), check(withContext)
	if before.Verdict != after.Verdict || len(before.CodeSet) != len(after.CodeSet) || before.JurisdictionID != after.JurisdictionID {
		t.Fatalf("optional utility context changed regulatory verdict: before=%+v after=%+v", before, after)
	}
	found := httptest.NewRecorder()
	withContext.ServeHTTP(found, httptest.NewRequest(http.MethodGet, "/v1/utility-context?state=CO&county=Denver", nil))
	if found.Code != 200 || !strings.Contains(found.Body.String(), "reported_county_candidates_not_site_verified") {
		t.Fatalf("utility context not mounted: %d %s", found.Code, found.Body.String())
	}
	notConfigured := httptest.NewRecorder()
	without.ServeHTTP(notConfigured, httptest.NewRequest(http.MethodGet, "/v1/utility-context?state=CO&county=Denver", nil))
	if notConfigured.Code == 200 {
		t.Fatal("unconfigured demo exposed utility data")
	}
}
