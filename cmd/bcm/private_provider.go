package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/laurajoyhutchins/building-code-map/engine"
)

// privateProvider only translates observations from the existing private BCM Engine.
// The public ProjectVerifier remains the sole owner of project-level verdicts.
// No corpus, source text, boundary geometry or private Git history is transferred.
type privateProvider struct {
	base           string
	token          string
	client         *http.Client
	expectedBundle string
	identityToken  func(context.Context, string) (string, error)
}

type privateReadiness struct {
	Status       string `json:"status"`
	Capabilities map[string]struct {
		Status string `json:"status"`
	} `json:"capabilities"`
}

type privateResolution struct {
	SchemaVersion string `json:"schema_version"`
	Resolution    struct {
		Status            string `json:"status"`
		ApplicabilityDate string `json:"applicability_date"`
		Geography         struct {
			JurisdictionID            string `json:"jurisdiction_id"`
			JurisdictionName          string `json:"jurisdiction_name"`
			CensusStateEquivalentFIPS string `json:"census_state_equivalent_fips"`
		} `json:"geography"`
		Adoptions []struct {
			CodeFamily    string `json:"code_family"`
			CodeName      string `json:"code_name"`
			StateCodeName string `json:"state_code_name"`
			Status        string `json:"status"`
			BaseModelCode *struct {
				EditionYear int `json:"edition_year"`
			} `json:"base_model_code"`
			SourceIDs     []string              `json:"source_ids"`
			ExactEvidence []engine.EvidenceLink `json:"exact_evidence"`
			Verification  struct {
				Status string `json:"status"`
			} `json:"verification"`
		} `json:"adoptions"`
		RequiredLocalRecords []string `json:"required_local_records"`
		Warnings             []string `json:"warnings"`
	} `json:"resolution"`
	Provenance struct {
		EngineVersion        string `json:"engine_version"`
		SourceCommit         string `json:"source_commit"`
		BundleManifestDigest string `json:"bundle_manifest_digest"`
	} `json:"provenance"`
}

func newPrivateProvider(raw, token string) (*privateProvider, error) {
	uri, err := url.Parse(raw)
	if err != nil || uri.Host == "" || uri.User != nil || uri.RawQuery != "" || uri.Fragment != "" ||
		uri.Path != "" && uri.Path != "/" {
		return nil, errors.New("BCM_PROVIDER_URL must be an HTTPS origin")
	}
	host := uri.Hostname()
	loopback := net.ParseIP(host) != nil && net.ParseIP(host).IsLoopback() || host == "localhost"
	if uri.Scheme != "https" && !(loopback && uri.Scheme == "http") {
		return nil, errors.New("BCM_PROVIDER_URL requires HTTPS outside loopback")
	}
	if uri.Port() != "" {
		if _, err := strconv.Atoi(uri.Port()); err != nil {
			return nil, errors.New("invalid provider port")
		}
	}
	// The configured origin is the sole authorized destination. A private
	// resolver must not redirect bearer credentials or project inputs to an
	// unreviewed host (including a same-domain subdomain).
	client := &http.Client{
		Timeout: 10 * time.Second,
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	return &privateProvider{base: strings.TrimSuffix(uri.String(), "/"), token: token, client: client}, nil
}

// Successful private-provider responses may contain at most 1 MiB of JSON and whitespace.
const maxPrivateProviderResponseBytes = 1 << 20

func (p *privateProvider) request(ctx context.Context, method, path string, payload any, dst any) (int, error) {
	var body io.Reader
	if payload != nil {
		data, err := json.Marshal(payload)
		if err != nil {
			return 0, err
		}
		body = bytes.NewReader(data)
	}
	req, err := http.NewRequestWithContext(ctx, method, p.base+path, body)
	if err != nil {
		return 0, err
	}
	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if p.identityToken != nil {
		token, err := p.identityToken(ctx, p.base)
		if err != nil || token == "" {
			return 0, errors.New("provider identity token unavailable")
		}
		req.Header.Set("Authorization", "Bearer "+token)
	} else if p.token != "" {
		req.Header.Set("Authorization", "Bearer "+p.token)
	}
	resp, err := p.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		if resp.StatusCode >= 500 {
			return resp.StatusCode, errors.New("upstream unavailable")
		}
		var e struct {
			Code string `json:"code"`
		}
		_ = json.NewDecoder(io.LimitReader(resp.Body, 1<<16)).Decode(&e)
		return resp.StatusCode, fmt.Errorf("upstream rejected request: %s", e.Code)
	}
	// Read one byte past the hard ceiling so even trailing whitespace cannot
	// disguise an oversized response. This bounds memory independently of
	// the JSON decoder's consumption of the first value.
	responseBytes, err := io.ReadAll(io.LimitReader(resp.Body, maxPrivateProviderResponseBytes+1))
	if err != nil {
		return 0, err
	}
	if len(responseBytes) > maxPrivateProviderResponseBytes {
		return 0, errors.New("provider response exceeds 1 MiB limit")
	}
	decoder := json.NewDecoder(bytes.NewReader(responseBytes))
	if err := decoder.Decode(dst); err != nil {
		return 0, err
	}
	var extra any
	if err := decoder.Decode(&extra); err != io.EOF {
		return 0, errors.New("provider response contains trailing content")
	}
	return resp.StatusCode, nil
}

func (p *privateProvider) readiness(ctx context.Context) (privateReadiness, error) {
	var out privateReadiness
	_, err := p.request(ctx, http.MethodGet, "/v1/readiness", nil, &out)
	if err != nil {
		return out, err
	}
	if p.expectedBundle != "" {
		var identity struct {
			BundleManifestDigest string `json:"bundle_manifest_digest"`
		}
		if _, err := p.request(ctx, http.MethodGet, "/v1/bundle", nil, &identity); err != nil {
			return out, err
		}
		if identity.BundleManifestDigest != p.expectedBundle {
			return out, errors.New("provider runtime bundle does not match reviewed identity")
		}
	}
	return out, nil
}

// Resolve asks the private authority for every covered code family. The legacy
// endpoint resolves one family at a time; an unscoped query returns no adoption
// evidence and must never masquerade as a complete project code basis.
func (p *privateProvider) Resolve(ctx context.Context, query engine.NormalizedQuery) (engine.Resolution, error) {
	readiness, err := p.readiness(ctx)
	capability := "address_lookup"
	if query.Point != nil {
		capability = "coordinate_resolution"
	}
	if err != nil || readiness.Status == "not_ready" ||
		readiness.Capabilities[capability].Status != "available" {
		return engine.Resolution{}, engine.EngineError{
			Code:      engine.ErrorRegulatoryCatalogUnavailable,
			Message:   "BCM " + capability + " is not ready",
			Retryable: true,
		}
	}

	families := []string{"building", "electrical", "energy", "fire_operational", "residential"}
	if query.CodeFamily != "" {
		families = []string{query.CodeFamily}
	}
	var result engine.Resolution
	for index, family := range families {
		part, err := p.resolveFamily(ctx, query, family)
		if err != nil {
			return engine.Resolution{}, err
		}
		if index == 0 {
			result = part
			continue
		}
		if result.Jurisdiction != part.Jurisdiction || result.Provenance != part.Provenance {
			return engine.Resolution{}, engine.EngineError{
				Code:    engine.ErrorDataBundleInvalid,
				Message: "BCM provider returned inconsistent jurisdiction or bundle identity",
			}
		}
		result.Codes = append(result.Codes, part.Codes...)
		result.Requirements = append(result.Requirements, part.Requirements...)
	}
	if len(query.ProjectFacts) > 0 {
		result.Requirements = append(result.Requirements, engine.ResolutionRequirement{
			ID: "project-facts-review", Kind: engine.RequirementLocalRecord,
			Prompt: "Project facts still require review against local amendments.",
		})
	}
	return result, nil
}

func (p *privateProvider) resolveFamily(ctx context.Context, query engine.NormalizedQuery, family string) (engine.Resolution, error) {
	request := struct {
		Point             *engine.Point `json:"point,omitempty"`
		Address           string        `json:"address,omitempty"`
		ApplicabilityDate string        `json:"applicability_date"`
		ProjectType       string        `json:"project_type,omitempty"`
		CodeFamily        string        `json:"code_family"`
	}{query.Point, query.Address, query.ApplicabilityDate, query.ProjectType, family}
	var upstream privateResolution
	status, err := p.request(ctx, http.MethodPost, "/v1/resolve", request, &upstream)
	if err != nil {
		code := engine.ErrorRegulatoryCatalogUnavailable
		switch status {
		case http.StatusBadRequest:
			code = engine.ErrorInvalidQuery
		case http.StatusConflict:
			code = engine.ErrorAddressAmbiguous
		case http.StatusUnprocessableEntity:
			code = engine.ErrorOutsideSupportedCoverage
		}
		return engine.Resolution{}, engine.EngineError{
			Code: code, Message: "BCM could not resolve this location and date",
		}
	}
	r := upstream.Resolution
	if upstream.SchemaVersion != "1" || r.Geography.JurisdictionID == "" ||
		r.Geography.JurisdictionName == "" || r.ApplicabilityDate != query.ApplicabilityDate ||
		upstream.Provenance.BundleManifestDigest == "" || upstream.Provenance.SourceCommit == "" ||
		(p.expectedBundle != "" && upstream.Provenance.BundleManifestDigest != p.expectedBundle) {
		return engine.Resolution{}, engine.EngineError{
			Code:    engine.ErrorDataBundleInvalid,
			Message: "BCM provider response has incomplete identity",
		}
	}
	country := "US"
	if strings.HasPrefix(r.Geography.JurisdictionID, "CA-") {
		country = "CA"
	} else if r.Geography.CensusStateEquivalentFIPS == "" {
		return engine.Resolution{}, engine.EngineError{
			Code: engine.ErrorDataBundleInvalid, Message: "BCM provider country is unknown",
		}
	}
	result := engine.Resolution{
		Jurisdiction: engine.JurisdictionContext{
			CountryCode:               country,
			PrimaryJurisdictionID:     r.Geography.JurisdictionID,
			PrimaryJurisdictionName:   r.Geography.JurisdictionName,
			CensusStateEquivalentFIPS: r.Geography.CensusStateEquivalentFIPS,
		},
		Provenance: engine.Provenance{
			EngineVersion: upstream.Provenance.EngineVersion,
			SourceCommit:  upstream.Provenance.SourceCommit,
			BundleID:      upstream.Provenance.BundleManifestDigest,
		},
	}
	codeStatus := engine.StatusInsufficientEvidence
	switch r.Status {
	case "local_record_required":
		codeStatus = engine.StatusLocalRecordRequired
	case "partially_resolved":
		codeStatus = engine.StatusPartiallyResolved
	case "conflicting":
		codeStatus = engine.StatusConflicting
	case "unsupported":
		codeStatus = engine.StatusUnsupported
	case "resolved", "verified":
		codeStatus = engine.StatusPartiallyResolved // A verified *project* needs independent evidence.
	}
	for _, ad := range r.Adoptions {
		if ad.CodeFamily != family {
			return engine.Resolution{}, engine.EngineError{
				Code:    engine.ErrorDataBundleInvalid,
				Message: "BCM provider returned an adoption for a different code family",
			}
		}
		status := codeStatus
		// Legacy source IDs or a private "verified" flag cannot establish a
		// public claim. Only a scoped, exact-evidence-bearing observation may
		// be proposed as resolved; Runtime independently checks the hashes,
		// legal locator, contradiction and operative date against trusted policy.
		if (r.Status == "resolved" || r.Status == "verified") &&
			ad.Verification.Status == "verified" && len(ad.ExactEvidence) > 0 &&
			len(r.RequiredLocalRecords) == 0 && len(r.Warnings) == 0 &&
			len(query.ProjectFacts) == 0 {
			status = engine.StatusResolved
		} else if ad.Verification.Status != "verified" && status == engine.StatusPartiallyResolved {
			status = engine.StatusInsufficientEvidence
		}
		edition := ad.StateCodeName
		if edition == "" {
			edition = ad.CodeName
		}
		if edition == "" && ad.BaseModelCode != nil && ad.BaseModelCode.EditionYear > 0 {
			edition = strconv.Itoa(ad.BaseModelCode.EditionYear)
		}
		code := engine.CodeResolution{
			Family: family, Edition: edition, Status: status,
			Basis:         "Source-backed adoption; local applicability may require confirmation.",
			ExactEvidence: ad.ExactEvidence,
		}
		for _, id := range ad.SourceIDs {
			if id != "" {
				code.Evidence = append(code.Evidence, engine.EvidenceRef{
					ID: id, Kind: "regulatory_source",
				})
			}
		}
		result.Codes = append(result.Codes, code)
	}
	if len(result.Codes) == 0 {
		result.Codes = []engine.CodeResolution{{
			Family: family, Status: engine.StatusInsufficientEvidence,
			Basis:      "No supported code adoption returned for this family, location, and date.",
			Unresolved: []string{"Confirm the locally applicable " + family + " code adoption."},
		}}
	}
	for i, message := range r.RequiredLocalRecords {
		if message != "" {
			result.Requirements = append(result.Requirements, engine.ResolutionRequirement{
				ID: fmt.Sprintf("%s-local-%d", family, i), Kind: engine.RequirementLocalRecord,
				Prompt: family + ": " + message,
			})
		}
	}
	for i, message := range r.Warnings {
		if message != "" {
			result.Requirements = append(result.Requirements, engine.ResolutionRequirement{
				ID: fmt.Sprintf("%s-warning-%d", family, i), Kind: engine.RequirementEvidenceDefect,
				Prompt: family + ": " + message,
			})
		}
	}
	return result, nil
}

func cloudRunIdentityToken(ctx context.Context, audience string) (string, error) {
	metadataURL := "http://metadata.google.internal/computeMetadata/v1/instance/service-accounts/default/identity?audience=" + url.QueryEscape(audience)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, metadataURL, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Metadata-Flavor", "Google")
	metadataClient := &http.Client{Timeout: 3 * time.Second}
	resp, err := metadataClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("cloud identity endpoint unavailable")
	}
	encoded, err := io.ReadAll(io.LimitReader(resp.Body, 8192))
	if err != nil {
		return "", err
	}
	token := strings.TrimSpace(string(encoded))
	if len(token) < 30 {
		return "", errors.New("cloud identity token is empty or invalid")
	}
	return token, nil
}
