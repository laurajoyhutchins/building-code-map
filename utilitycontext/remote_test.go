package utilitycontext

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func remoteFixture(t *testing.T, serve http.HandlerFunc) (*Remote, *httptest.Server) {
	t.Helper()
	server := httptest.NewTLSServer(serve)
	remote, err := NewRemote(server.URL, strings.Repeat("a", 64))
	if err != nil {
		server.Close()
		t.Fatal(err)
	}
	remote.client.Transport = server.Client().Transport
	return remote, server
}

func catalogResponse(t *testing.T) Response {
	t.Helper()
	path, digest := fixture(t, nil)
	catalog, err := Open(path, digest)
	if err != nil {
		t.Fatal(err)
	}
	value, err := catalog.Lookup("CO", "Denver")
	if err != nil {
		t.Fatal(err)
	}
	// The mock service's view identity is pinned independently.
	value.ViewSHA256 = strings.Repeat("a", 64)
	return value
}

func TestRemoteLookupAndVersionedHandler(t *testing.T) {
	payload := catalogResponse(t)
	remote, server := remoteFixture(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != Path ||
			r.URL.Query().Get("state") != "CO" || r.URL.Query().Get("county") != "Denver" ||
			len(r.URL.Query()) != 2 {
			t.Errorf("unexpected external request: %s %s", r.Method, r.URL)
			http.Error(w, "bad call", http.StatusBadRequest)
			return
		}
		_ = json.NewEncoder(w).Encode(payload)
	})
	defer server.Close()
	got, err := remote.Lookup("co", "Denver")
	if err != nil || got.Status != payload.Status ||
		got.ViewSHA256 != payload.ViewSHA256 || len(got.Candidates) != 1 {
		t.Fatalf("rejected valid catalog: result=%+v err=%v", got, err)
	}
	accepted := httptest.NewRecorder()
	remote.Handler().ServeHTTP(accepted, httptest.NewRequest(http.MethodGet,
		Path+"?state=co&county=Denver", nil))
	if accepted.Code != http.StatusOK {
		t.Fatalf("remote handler returned %d %s", accepted.Code, accepted.Body.String())
	}
	for _, path := range []string{
		Path+"?state=CO&county=Denver&all=true",
		Path+"?state=CO&county=Denver&county=Adams",
		Path+"?state=CO",
	} {
		rec := httptest.NewRecorder()
		remote.Handler().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, path, nil))
		if rec.Code != http.StatusBadRequest {
			t.Fatalf("accepted unbounded or ambiguous lookup: %s -> %d", path, rec.Code)
		}
	}
}

func TestRemoteRejectsCorruptOrOverclaimedCatalogResponses(t *testing.T) {
	original := catalogResponse(t)
	variant := []struct {
		name string
		edit func(*Response)
	}{
		{"wrong view digest", func(v *Response) { v.ViewSHA256 = strings.Repeat("b", 64) }},
		{"wrong state", func(v *Response) { v.State = "NY" }},
		{"wrong county", func(v *Response) { v.County = "Boulder" }},
		{"promoted status", func(v *Response) { v.Status = "verified_serving_utility" }},
		{"missing source identity", func(v *Response) { v.SourceSHA256 = "" }},
		{"missing disclaimer", func(v *Response) { v.Disclaimer = "" }},
		{"null candidate array", func(v *Response) { v.Candidates = nil }},
		{"invalid county source row", func(v *Response) { v.Candidates[0].SourceRow = 0 }},
		{"duplicate candidate", func(v *Response) { v.Candidates = append(v.Candidates, v.Candidates[0]) }},
	}
	for _, tc := range variant {
		t.Run(tc.name, func(t *testing.T) {
			value := original
			value.Candidates = append([]Candidate(nil), original.Candidates...)
			tc.edit(&value)
			remote, server := remoteFixture(t, func(w http.ResponseWriter, _ *http.Request) {
				_ = json.NewEncoder(w).Encode(value)
			})
			defer server.Close()
			if _, err := remote.Lookup("CO", "Denver"); err == nil {
				t.Fatal("accepted malformed or unsupported source claim")
			}
			handlerResult := httptest.NewRecorder()
			remote.Handler().ServeHTTP(handlerResult, httptest.NewRequest(http.MethodGet,
				Path+"?state=CO&county=Denver", nil))
			if handlerResult.Code != http.StatusServiceUnavailable ||
				strings.Contains(handlerResult.Body.String(), "candidates") {
				t.Fatalf("failed to deny uncertain EIA response: %d %s",
					handlerResult.Code, handlerResult.Body.String())
			}
		})
	}
}

func TestRemoteRejectsRedirectsUnknownFieldsAndUnavailableServer(t *testing.T) {
	remote, server := remoteFixture(t, func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://untrusted.example/steal", http.StatusFound)
	})
	if _, err := remote.Lookup("CO", "Denver"); err == nil {
		t.Fatal("followed or accepted redirect")
	}
	server.Close()
	if _, err := remote.Lookup("CO", "Denver"); err == nil {
		t.Fatal("accepted unavailable catalog")
	}
	remote, server = remoteFixture(t, func(w http.ResponseWriter, _ *http.Request) {
		_, _ = w.Write([]byte(`{"status":"reported_county_candidates_not_site_verified","unexpected":"yes"}`))
	})
	defer server.Close()
	if _, err := remote.Lookup("CO", "Denver"); err == nil {
		t.Fatal("accepted non-contract field")
	}
}

func TestRemoteRejectsUnpinnedAndUntrustedOrigins(t *testing.T) {
	for _, origin := range []string{
		"http://example.org", "https://example.org/path",
		"https://user@example.org", "https://example.org?x=1",
		"https://example.org#fragment",
	} {
		if _, err := NewRemote(origin, strings.Repeat("a", 64)); err == nil {
			t.Fatalf("accepted untrusted origin %q", origin)
		}
	}
	if _, err := NewRemote("https://example.org", "missing"); err == nil {
		t.Fatal("accepted unpinned source")
	}
}
