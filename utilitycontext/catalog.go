// Package utilitycontext reads a pinned, source-bound view of EIA-861 county
// candidates. It is intentionally independent of BCM's regulatory engine.
package utilitycontext

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const Path = "/v1/utility-context"
const maxViewBytes = 8 << 20

var lowercaseHash = regexp.MustCompile(`^[0-9a-f]{64}$`)
var stateCode = regexp.MustCompile(`^[A-Z]{2}$`)
var countyName = regexp.MustCompile(`^[\pL\pN .,'()/-]{1,120}$`)

type Activity struct {
	Code  string `json:"code"`
	Value string `json:"value"`
}

type Candidate struct {
	EIAID        int        `json:"eia_id"`
	ReportedName string     `json:"reported_name"`
	SourceSheet  string     `json:"source_sheet"`
	SourceRow    int        `json:"source_row"`
	Activities   []Activity `json:"activities"`
}

type row struct {
	State        string     `json:"state"`
	County       string     `json:"county"`
	EIAID        int        `json:"eia_id"`
	ReportedName string     `json:"reported_name"`
	SourceSheet  string     `json:"source_sheet"`
	SourceRow    int        `json:"source_row"`
	Activities   [][]string `json:"activities"`
}

type view struct {
	Format        string `json:"format"`
	ReleaseID     string `json:"release_id"`
	ReportingYear int    `json:"reporting_year"`
	SourceSHA256  string `json:"source_sha256"`
	SQLiteSHA256  string `json:"sqlite_sha256"`
	SourceURL     string `json:"source_url"`
	Scope         string `json:"scope"`
	Records       []row  `json:"records"`
}

type Catalog struct {
	releaseID  string
	year       int
	sourceHash string
	sqliteHash string
	viewHash   string
	sourceURL  string
	counties   map[string][]Candidate
}

// Open verifies the entire local projection against the operator-pinned digest.
// The projection is produced from a separately verified EIA SQLite/manifest pair.
// No network access, GCS credentials, or dynamic SQLite driver are needed.
func Open(path, pinnedSHA256 string) (*Catalog, error) {
	if !lowercaseHash.MatchString(pinnedSHA256) {
		return nil, errors.New("utility context requires a pinned lowercase SHA-256")
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := io.ReadAll(io.LimitReader(file, maxViewBytes+1))
	if err != nil {
		return nil, err
	}
	if len(data) == 0 || len(data) > maxViewBytes {
		return nil, errors.New("utility context view is empty or oversized")
	}
	computed := fmt.Sprintf("%x", sha256.Sum256(data))
	if computed != pinnedSHA256 {
		return nil, errors.New("utility context view SHA-256 mismatch")
	}
	var document view
	if err := json.Unmarshal(data, &document); err != nil {
		return nil, err
	}
	if document.Format != "bcm/eia-county-candidates/v1" ||
		!strings.HasPrefix(document.ReleaseID, "eia861-") ||
		document.ReportingYear < 2000 || document.ReportingYear > 2100 ||
		!lowercaseHash.MatchString(document.SourceSHA256) ||
		!lowercaseHash.MatchString(document.SQLiteSHA256) ||
		!strings.HasPrefix(document.SourceURL, "https://www.eia.gov/") ||
		document.Scope != "EIA-reported county candidates only; no verified serving utility" ||
		len(document.Records) == 0 || len(document.Records) > 20000 {
		return nil, errors.New("invalid utility context release identity or scope")
	}
	catalog := &Catalog{releaseID: document.ReleaseID, year: document.ReportingYear,
		sourceHash: document.SourceSHA256, sqliteHash: document.SQLiteSHA256,
		viewHash: computed, sourceURL: document.SourceURL, counties: make(map[string][]Candidate)}
	seen := make(map[string]struct{}, len(document.Records))
	for _, entry := range document.Records {
		if !stateCode.MatchString(entry.State) || !countyName.MatchString(entry.County) ||
			entry.EIAID <= 0 || entry.SourceRow <= 0 || entry.SourceSheet != "Counties_States" && entry.SourceSheet != "Counties_Territories" ||
			strings.TrimSpace(entry.ReportedName) == "" {
			return nil, errors.New("invalid utility context county record")
		}
		key := entry.State + "\x00" + strings.ToLower(entry.County)
		ident := fmt.Sprintf("%s\x00%d", key, entry.EIAID)
		if _, exists := seen[ident]; exists {
			return nil, errors.New("duplicate utility county candidate")
		}
		seen[ident] = struct{}{}
		candidate := Candidate{EIAID: entry.EIAID, ReportedName: entry.ReportedName,
			SourceSheet: entry.SourceSheet, SourceRow: entry.SourceRow, Activities: make([]Activity, 0, len(entry.Activities))}
		for _, reported := range entry.Activities {
			if len(reported) != 2 || reported[0] == "" || reported[1] != "Y" && reported[1] != "N" {
				return nil, errors.New("invalid EIA activity marker")
			}
			candidate.Activities = append(candidate.Activities, Activity{Code: reported[0], Value: reported[1]})
		}
		catalog.counties[key] = append(catalog.counties[key], candidate)
	}
	return catalog, nil
}

// Response is distinct from a ProjectCodeBasis and cannot establish a serving
// utility, a building-code jurisdiction, or any engineering design input.
type Response struct {
	Status        string      `json:"status"`
	State         string      `json:"state"`
	County        string      `json:"county"`
	ReleaseID     string      `json:"release_id"`
	ReportingYear int         `json:"reporting_year"`
	SourceSHA256  string      `json:"source_sha256"`
	SQLiteSHA256  string      `json:"sqlite_sha256"`
	ViewSHA256    string      `json:"view_sha256"`
	SourceURL     string      `json:"source_url"`
	Disclaimer    string      `json:"disclaimer"`
	Candidates    []Candidate `json:"candidates"`
}

func (catalog *Catalog) Lookup(state, county string) (Response, error) {
	state = strings.ToUpper(strings.TrimSpace(state))
	county = strings.TrimSpace(county)
	if !stateCode.MatchString(state) || !countyName.MatchString(county) {
		return Response{}, errors.New("state must be a two-letter code and county must be a reported name")
	}
	matches := catalog.counties[state+"\x00"+strings.ToLower(county)]
	candidates := make([]Candidate, len(matches))
	for i, item := range matches {
		candidates[i] = item
		candidates[i].Activities = append([]Activity{}, item.Activities...)
	}
	return Response{
		Status: "reported_county_candidates_not_site_verified", State: state, County: county,
		ReleaseID: catalog.releaseID, ReportingYear: catalog.year,
		SourceSHA256: catalog.sourceHash, SQLiteSHA256: catalog.sqliteHash,
		ViewSHA256: catalog.viewHash, SourceURL: catalog.sourceURL,
		Disclaimer: "EIA-861 reports county associations, not an address's actual serving utility. " +
			"An empty result does not mean no utility serves the location. Confirm provider identity independently.",
		Candidates: candidates,
	}, nil
}

// Handler exposes only a bounded single-county lookup, never a bulk export.
// Browser CORS and authentication remain responsibilities of the BCM host.
func (catalog *Catalog) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", "GET")
			w.WriteHeader(http.StatusMethodNotAllowed)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "method_not_allowed"})
			return
		}
		if r.URL.Path != Path || len(r.URL.Query()) != 2 || len(r.URL.Query()["state"]) != 1 || len(r.URL.Query()["county"]) != 1 {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "invalid_query"})
			return
		}
		answer, err := catalog.Lookup(r.URL.Query().Get("state"), r.URL.Query().Get("county"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "invalid_query"})
			return
		}
		_ = json.NewEncoder(w).Encode(answer)
	})
}
