package utilitycontext

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
)

func fixture(t *testing.T, mutate func(*view)) (string, string) {
	t.Helper()
	document := view{Format: "bcm/eia-county-candidates/v1", ReleaseID: "eia861-2025-aabbcc-v0.1.2",
		ReportingYear: 2025, SourceSHA256: strings.Repeat("a", 64), SQLiteSHA256: strings.Repeat("b", 64),
		SourceURL: "https://www.eia.gov/electricity/data/eia861/", Scope: "EIA-reported county candidates only; no verified serving utility",
		Records: []row{{State: "CO", County: "Denver", EIAID: 15466, ReportedName: "Public Service Co of Colorado", SourceSheet: "Counties_States", SourceRow: 7633,
			Activities: [][]string{{"Distribution", "Y"}, {"Retail Marketing", "N"}}}},
	}
	if mutate != nil {
		mutate(&document)
	}
	data, err := json.Marshal(document)
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(t.TempDir(), "view.json")
	if err := os.WriteFile(path, data, 0600); err != nil {
		t.Fatal(err)
	}
	return path, fmt.Sprintf("%x", sha256.Sum256(data))
}

func TestPinnedCountyCandidatesDoNotClaimSiteService(t *testing.T) {
	path, digest := fixture(t, nil)
	catalog, err := Open(path, digest)
	if err != nil {
		t.Fatal(err)
	}
	got, err := catalog.Lookup("co", "DENVER")
	if err != nil {
		t.Fatal(err)
	}
	if got.Status != "reported_county_candidates_not_site_verified" || got.ReleaseID == "" ||
		got.SourceSHA256 != strings.Repeat("a", 64) || len(got.Candidates) != 1 ||
		got.Candidates[0].EIAID != 15466 || got.Candidates[0].Activities[0].Value != "Y" {
		t.Fatalf("incorrect county candidate: %+v", got)
	}
	none, err := catalog.Lookup("CO", "Larimer")
	if err != nil || none.Candidates == nil || len(none.Candidates) != 0 || none.Disclaimer == "" {
		t.Fatalf("no-match result must be explicit and non-authoritative: %+v %v", none, err)
	}
}

func TestRejectTamperingAndUnsupportedCountyEvidence(t *testing.T) {
	path, digest := fixture(t, nil)
	if _, err := Open(path, strings.Repeat("c", 64)); err == nil {
		t.Fatal("accepted wrong digest")
	}
	if err := os.WriteFile(path, []byte("tampered"), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := Open(path, digest); err == nil {
		t.Fatal("accepted tampered source")
	}
	for _, edit := range []func(*view){
		func(v *view) { v.Scope = "confirmed serving utilities" },
		func(v *view) { v.Records = append(v.Records, v.Records[0]) },
		func(v *view) { v.Records[0].Activities[0][1] = "maybe" },
		func(v *view) { v.Records[0].SourceSheet = "user_entered" },
		func(v *view) { v.SourceSHA256 = "unknown" },
	} {
		candidatePath, candidateHash := fixture(t, edit)
		if _, err := Open(candidatePath, candidateHash); err == nil {
			t.Fatal("accepted invalid view")
		}
	}
}

func TestHTTPBoundedLookupAndMethodIsolation(t *testing.T) {
	path, digest := fixture(t, nil)
	catalog, err := Open(path, digest)
	if err != nil {
		t.Fatal(err)
	}
	for _, test := range []struct {
		method, uri string
		want        int
	}{
		{http.MethodGet, Path + "?state=CO&county=Denver", 200},
		{http.MethodGet, Path + "?state=CO&county=Denver&all=true", 400},
		{http.MethodGet, Path + "?state=CO&county=Denver&county=Adams", 400},
		{http.MethodGet, Path + "?state=CO", 400},
		{http.MethodGet, Path + "?state=CO&county=Denver%0AInjected", 400},
		{http.MethodPost, Path + "?state=CO&county=Denver", 405},
	} {
		recorder := httptest.NewRecorder()
		catalog.Handler().ServeHTTP(recorder, httptest.NewRequest(test.method, test.uri, nil))
		if recorder.Code != test.want {
			t.Fatalf("%s %s: got %d want %d", test.method, test.uri, recorder.Code, test.want)
		}
	}
}
