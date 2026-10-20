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
	base   string
	token  string
	client *http.Client
	expectedBundle string
}

type privateReadiness struct {
	Status       string `json:"status"`
	Capabilities map[string]struct{ Status string `json:"status"` } `json:"capabilities"`
}

type privateResolution struct {
	SchemaVersion string `json:"schema_version"`
	Resolution struct {
		Status               string `json:"status"`
		ApplicabilityDate    string `json:"applicability_date"`
		Geography struct {
			JurisdictionID            string `json:"jurisdiction_id"`
			JurisdictionName          string `json:"jurisdiction_name"`
			CensusStateEquivalentFIPS string `json:"census_state_equivalent_fips"`
		} `json:"geography"`
		Adoptions []struct {
			CodeFamily  string `json:"code_family"`
			CodeName    string `json:"code_name"`
			Status      string `json:"status"`
			BaseModelCode *struct {
				EditionYear int `json:"edition_year"`
			} `json:"base_model_code"`
			SourceIDs []string `json:"source_ids"`
			Verification struct { Status string `json:"status"` } `json:"verification"`
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
		if _, err := strconv.Atoi(uri.Port()); err != nil { return nil, errors.New("invalid provider port") }
	}
	return &privateProvider{base: strings.TrimSuffix(uri.String(), "/"), token: token, client: &http.Client{Timeout: 10 * time.Second}}, nil
}

func (p *privateProvider) request(ctx context.Context, method, path string, payload any, dst any) (int,error) {
	var body io.Reader
	if payload != nil {
		data, err := json.Marshal(payload)
		if err != nil { return 0,err }
		body = bytes.NewReader(data)
	}
	req, err := http.NewRequestWithContext(ctx,method,p.base+path,body)
	if err != nil { return 0,err }
	if payload != nil { req.Header.Set("Content-Type", "application/json") }
	if p.token != "" { req.Header.Set("Authorization","Bearer "+p.token) }
	resp,err:=p.client.Do(req)
	if err != nil { return 0,err }
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		if resp.StatusCode >= 500 { return resp.StatusCode, errors.New("upstream unavailable") }
		var e struct { Code string `json:"code"` }
		_ = json.NewDecoder(io.LimitReader(resp.Body, 1<<16)).Decode(&e)
		return resp.StatusCode,fmt.Errorf("upstream rejected request: %s",e.Code)
	}
	if err := json.NewDecoder(io.LimitReader(resp.Body,1<<20)).Decode(dst); err != nil {return 0,err}
	return resp.StatusCode,nil
}

func (p *privateProvider) readiness(ctx context.Context) (privateReadiness,error) {
	var out privateReadiness
	_,err:=p.request(ctx,http.MethodGet,"/v1/readiness",nil,&out)
	if err != nil {return out,err}
	if p.expectedBundle != "" {
		var identity struct { BundleManifestDigest string `json:"bundle_manifest_digest"` }
		if _,err:=p.request(ctx,http.MethodGet,"/v1/bundle",nil,&identity); err!=nil {return out,err}
		if identity.BundleManifestDigest != p.expectedBundle {
			return out,errors.New("provider runtime bundle does not match reviewed identity")
		}
	}
	return out,nil
}

func (p *privateProvider) Resolve(ctx context.Context, query engine.NormalizedQuery) (engine.Resolution,error) {
	if readiness,err:=p.readiness(ctx); err != nil ||
		readiness.Capabilities["address_lookup"].Status != "available" {
		return engine.Resolution{},engine.EngineError{
			Code:engine.ErrorRegulatoryCatalogUnavailable,Message:"BCM address lookup is not ready",Retryable:true,
		}
	}
	request := struct{
		Point *engine.Point `json:"point,omitempty"`
		Address string `json:"address,omitempty"`
		ApplicabilityDate string `json:"applicability_date"`
		ProjectType string `json:"project_type,omitempty"`
	}{query.Point,query.Address,query.ApplicabilityDate,query.ProjectType}
	var upstream privateResolution
	status,err:=p.request(ctx,http.MethodPost,"/v1/resolve",request,&upstream)
	if err!=nil {
		code:=engine.ErrorRegulatoryCatalogUnavailable
		switch status {
		case 400: code=engine.ErrorInvalidQuery
		case 409: code=engine.ErrorAddressAmbiguous
		case 422: code=engine.ErrorOutsideSupportedCoverage
		}
		return engine.Resolution{},engine.EngineError{Code:code,Message:"BCM could not resolve this location and date"}
	}
	r:=upstream.Resolution
	if upstream.SchemaVersion!="1" || r.Geography.JurisdictionID=="" ||
		r.Geography.JurisdictionName=="" || r.ApplicabilityDate!=query.ApplicabilityDate ||
		upstream.Provenance.BundleManifestDigest=="" ||
		upstream.Provenance.SourceCommit=="" ||
		(p.expectedBundle!="" && upstream.Provenance.BundleManifestDigest!=p.expectedBundle) {
		return engine.Resolution{},engine.EngineError{Code:engine.ErrorDataBundleInvalid,Message:"BCM provider response has incomplete identity"}
	}
	country:="US"
	if strings.HasPrefix(r.Geography.JurisdictionID,"CA-") { country="CA" } else if r.Geography.CensusStateEquivalentFIPS=="" {
		return engine.Resolution{},engine.EngineError{Code:engine.ErrorDataBundleInvalid,Message:"BCM provider country is unknown"}
	}
	result:=engine.Resolution{
		Jurisdiction:engine.JurisdictionContext{
			CountryCode:country,PrimaryJurisdictionID:r.Geography.JurisdictionID,
			PrimaryJurisdictionName:r.Geography.JurisdictionName,
			CensusStateEquivalentFIPS:r.Geography.CensusStateEquivalentFIPS,
		},
		Provenance:engine.Provenance{
			EngineVersion:upstream.Provenance.EngineVersion,
			SourceCommit:upstream.Provenance.SourceCommit,
			BundleID:upstream.Provenance.BundleManifestDigest,
		},
	}
	statusCode:=engine.StatusInsufficientEvidence
	switch r.Status {
	case "local_record_required": statusCode=engine.StatusLocalRecordRequired
	case "conflicting":statusCode=engine.StatusConflicting
	case "unsupported":statusCode=engine.StatusUnsupported
	case "resolved","verified":statusCode=engine.StatusPartiallyResolved // No verified project claim until contract-level evidence is reviewed.
	}
	for _,ad:=range r.Adoptions {
		if strings.TrimSpace(ad.CodeFamily)=="" {continue}
		codeStatus:=statusCode
		if ad.Verification.Status!="verified" && codeStatus==engine.StatusPartiallyResolved {
			codeStatus=engine.StatusInsufficientEvidence
		}
		edition:=ad.CodeName
		if edition=="" && ad.BaseModelCode!=nil && ad.BaseModelCode.EditionYear>0 {
			edition=strconv.Itoa(ad.BaseModelCode.EditionYear)
		}
		code:=engine.CodeResolution{
			Family:ad.CodeFamily,Edition:edition,Status:codeStatus,
			Basis:"Adoption reported by the BCM source catalog; local applicability may require confirmation.",
		}
		for _,id:=range ad.SourceIDs {
			if id!="" {code.Evidence=append(code.Evidence,engine.EvidenceRef{ID:id,Kind:"regulatory_source"})}
		}
		result.Codes=append(result.Codes,code)
	}
	if len(result.Codes)==0 {
		// Keep unresolved results visible without inventing an adopted edition.
		family:="unspecified"
		if query.CodeFamily!="" {family=query.CodeFamily}
		result.Codes=[]engine.CodeResolution{{Family:family,Status:engine.StatusInsufficientEvidence,
			Basis:"No supported code adoption was returned for this location and date."}}
	}
	for i,message:=range r.RequiredLocalRecords {
		if message!="" {result.Requirements=append(result.Requirements,engine.ResolutionRequirement{
			ID:fmt.Sprintf("local-%d",i),Kind:engine.RequirementLocalRecord,Prompt:message,
		})}
	}
	for i,message:=range r.Warnings {
		if message!="" {result.Requirements=append(result.Requirements,engine.ResolutionRequirement{
			ID:fmt.Sprintf("warning-%d",i),Kind:engine.RequirementEvidenceDefect,Prompt:message,
		})}
	}
	if len(query.ProjectFacts)>0 {
		result.Requirements=append(result.Requirements,engine.ResolutionRequirement{
			ID:"project-facts-review",Kind:engine.RequirementLocalRecord,
			Prompt:"Project-specific facts have not yet been checked against local amendments.",
		})
	}
	return result,nil
}
