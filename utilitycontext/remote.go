package utilitycontext

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Provider is the versioned county-candidate reader consumed by BCM.
// Its only authority is the independently pinned EIA release view.
type Provider interface {
	Lookup(state, county string) (Response, error)
	Handler() http.Handler
}

var ErrSourceUnavailable = errors.New("catalog county candidate source unavailable")
var ErrInvalidCounty = errors.New("invalid state or county")

// Remote consumes Electric Utility Catalog's v1 read-only HTTP service.
// It has no access to private EIA source bytes, the catalog database, or BCM's verdict.
type Remote struct {
	origin     string
	viewSHA256 string
	client     *http.Client
}

// NewRemote requires an operator-configured HTTPS origin and exact view digest.
// An arbitrary URL from an end user cannot select this service.
func NewRemote(origin, pinnedViewSHA256 string) (*Remote, error) {
	u, err := url.Parse(origin)
	if err != nil || u.Scheme != "https" || u.Host == "" ||
		u.User != nil || u.Path != "" || u.RawQuery != "" ||
		u.Fragment != "" || u.Opaque != "" || u.RawFragment != "" ||
		!lowercaseHash.MatchString(pinnedViewSHA256) {
		return nil, errors.New("utility catalog requires an HTTPS origin and pinned view SHA-256")
	}
	return &Remote{
		origin:     u.String(),
		viewSHA256: pinnedViewSHA256,
		client: &http.Client{
			Timeout: 8 * time.Second,
			CheckRedirect: func(*http.Request, []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}, nil
}

func (c *Remote) Lookup(state, county string) (Response, error) {
	state = strings.ToUpper(strings.TrimSpace(state))
	county = strings.TrimSpace(county)
	if !stateCode.MatchString(state) || !countyName.MatchString(county) {
		return Response{}, ErrInvalidCounty
	}
	query := url.Values{"state": {state}, "county": {county}}
	request, err := http.NewRequest(http.MethodGet, c.origin+Path+"?"+query.Encode(), nil)
	if err != nil {
		return Response{}, ErrSourceUnavailable
	}
	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()
	request = request.WithContext(ctx)
	response, err := c.client.Do(request)
	if err != nil {
		return Response{}, fmt.Errorf("%w: request failed", ErrSourceUnavailable)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return Response{}, fmt.Errorf("%w: HTTP %d", ErrSourceUnavailable, response.StatusCode)
	}
	limited := io.LimitReader(response.Body, (8<<20)+1)
	data, err := io.ReadAll(limited)
	if err != nil || len(data) == 0 || len(data) > 8<<20 {
		return Response{}, fmt.Errorf("%w: invalid response size", ErrSourceUnavailable)
	}
	decoder := json.NewDecoder(strings.NewReader(string(data)))
	decoder.DisallowUnknownFields()
	var value Response
	if err := decoder.Decode(&value); err != nil {
		return Response{}, fmt.Errorf("%w: malformed response", ErrSourceUnavailable)
	}
	var trailing any
	if err := decoder.Decode(&trailing); !errors.Is(err, io.EOF) ||
		!validRemoteResult(value, state, county, c.viewSHA256) {
		return Response{}, fmt.Errorf("%w: inconsistent source identity or schema", ErrSourceUnavailable)
	}
	return value, nil
}

func validRemoteResult(value Response, state, county, viewHash string) bool {
	if value.Status != "reported_county_candidates_not_site_verified" ||
		value.State != state || value.County != county ||
		value.ReleaseID == "" || !strings.HasPrefix(value.ReleaseID, "eia861-") ||
		value.ReportingYear < 2000 || value.ReportingYear > 2100 ||
		!lowercaseHash.MatchString(value.SourceSHA256) ||
		!lowercaseHash.MatchString(value.SQLiteSHA256) ||
		value.ViewSHA256 != viewHash ||
		!strings.HasPrefix(value.SourceURL, "https://www.eia.gov/") ||
		strings.TrimSpace(value.Disclaimer) == "" ||
		value.Candidates == nil || len(value.Candidates) > 20000 {
		return false
	}
	seen := make(map[int]struct{}, len(value.Candidates))
	for _, candidate := range value.Candidates {
		if candidate.EIAID <= 0 || candidate.SourceRow <= 0 ||
			strings.TrimSpace(candidate.ReportedName) == "" ||
			(candidate.SourceSheet != "Counties_States" && candidate.SourceSheet != "Counties_Territories") ||
			candidate.Activities == nil {
			return false
		}
		if _, exists := seen[candidate.EIAID]; exists {
			return false
		}
		seen[candidate.EIAID] = struct{}{}
		for _, activity := range candidate.Activities {
			if activity.Code == "" || (activity.Value != "Y" && activity.Value != "N") {
				return false
			}
		}
	}
	return true
}

// Handler preserves the old one-county HTTP contract for BCM's existing callers.
func (c *Remote) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", "GET")
			w.WriteHeader(http.StatusMethodNotAllowed)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "method_not_allowed"})
			return
		}
		if r.URL.Path != Path || len(r.URL.Query()) != 2 ||
			len(r.URL.Query()["state"]) != 1 || len(r.URL.Query()["county"]) != 1 {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": "invalid_query"})
			return
		}
		result, err := c.Lookup(r.URL.Query().Get("state"), r.URL.Query().Get("county"))
		if err != nil {
			code := http.StatusServiceUnavailable
			message := "source_unavailable"
			if errors.Is(err, ErrInvalidCounty) {
				code, message = http.StatusBadRequest, "invalid_query"
			}
			w.WriteHeader(code)
			_ = json.NewEncoder(w).Encode(map[string]string{"error": message})
			return
		}
		_ = json.NewEncoder(w).Encode(result)
	})
}
