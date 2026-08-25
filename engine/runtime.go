package engine

import (
	"context"
	"strings"
)

type ResolutionStatus string

const (
	StatusResolved             ResolutionStatus = "resolved"
	StatusPartiallyResolved    ResolutionStatus = "partially_resolved"
	StatusLocalRecordRequired  ResolutionStatus = "local_record_required"
	StatusUnsupported          ResolutionStatus = "unsupported"
	StatusNotApplicable        ResolutionStatus = "not_applicable"
	StatusInsufficientEvidence ResolutionStatus = "insufficient_evidence"
	StatusAmbiguous            ResolutionStatus = "ambiguous"
	StatusConflicting          ResolutionStatus = "conflicting"
)

var allowedResolutionStatuses = map[ResolutionStatus]struct{}{
	StatusResolved: {}, StatusPartiallyResolved: {}, StatusLocalRecordRequired: {}, StatusUnsupported: {},
	StatusNotApplicable: {}, StatusInsufficientEvidence: {}, StatusAmbiguous: {}, StatusConflicting: {},
}

type RequirementKind string

const (
	RequirementProjectFact    RequirementKind = "project_fact"
	RequirementLocalRecord    RequirementKind = "local_record"
	RequirementEvidenceDefect RequirementKind = "evidence_defect"
)

var allowedRequirementKinds = map[RequirementKind]struct{}{
	RequirementProjectFact: {}, RequirementLocalRecord: {}, RequirementEvidenceDefect: {},
}

type JurisdictionContext struct {
	CountryCode               string `json:"country_code"`
	PrimaryJurisdictionID     string `json:"primary_jurisdiction_id"`
	PrimaryJurisdictionName   string `json:"primary_jurisdiction_name"`
	CensusStateEquivalentFIPS string `json:"census_state_equivalent_fips,omitempty"`
}

type EvidenceRef struct {
	ID   string `json:"id"`
	Kind string `json:"kind"`
}

type DerivedFact struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Basis string `json:"basis,omitempty"`
}

type ResolutionRequirement struct {
	ID      string          `json:"id"`
	Kind    RequirementKind `json:"kind"`
	Prompt  string          `json:"prompt"`
	FactKey string          `json:"fact_key,omitempty"`
}

type CodeResolution struct {
	Family     string           `json:"family"`
	Edition    string           `json:"edition,omitempty"`
	Status     ResolutionStatus `json:"status"`
	Basis      string           `json:"basis,omitempty"`
	Unresolved []string         `json:"unresolved,omitempty"`
	Evidence   []EvidenceRef    `json:"evidence,omitempty"`
}

type Provenance struct {
	EngineVersion string `json:"engine_version,omitempty"`
	SourceCommit  string `json:"source_commit,omitempty"`
	BundleID      string `json:"bundle_id,omitempty"`
}

type Resolution struct {
	SchemaVersion string                  `json:"schema_version"`
	Query         NormalizedQuery         `json:"query"`
	Jurisdiction  JurisdictionContext     `json:"jurisdiction"`
	Codes         []CodeResolution        `json:"codes"`
	Requirements  []ResolutionRequirement `json:"requirements,omitempty"`
	DerivedFacts  []DerivedFact           `json:"derived_facts,omitempty"`
	Provenance    Provenance              `json:"provenance,omitempty"`
}

type Provider interface {
	Resolve(context.Context, NormalizedQuery) (Resolution, error)
}

type Runtime struct{ provider Provider }

func NewRuntime(provider Provider) Runtime { return Runtime{provider: provider} }

func (runtime Runtime) Resolve(ctx context.Context, query Query) (Resolution, error) {
	normalized, err := NormalizeQuery(query)
	if err != nil {
		return Resolution{}, err
	}
	if runtime.provider == nil {
		return Resolution{}, EngineError{Code: ErrorRegulatoryCatalogUnavailable, Message: "resolution provider is not configured", Retryable: false}
	}
	result, err := runtime.provider.Resolve(ctx, normalized)
	if err != nil {
		return Resolution{}, err
	}
	result.SchemaVersion = SchemaVersion
	result.Query = normalized
	if err := validateResolution(result); err != nil {
		return Resolution{}, err
	}
	return result, nil
}

func validateResolution(result Resolution) error {
	if strings.TrimSpace(result.Jurisdiction.CountryCode) == "" || strings.TrimSpace(result.Jurisdiction.PrimaryJurisdictionID) == "" || strings.TrimSpace(result.Jurisdiction.PrimaryJurisdictionName) == "" {
		return EngineError{Code: ErrorDataBundleInvalid, Message: "resolution requires country and primary-jurisdiction identity"}
	}
	if len(result.Codes) == 0 {
		return EngineError{Code: ErrorDataBundleInvalid, Message: "resolution requires at least one explicit code-family result"}
	}
	for index, code := range result.Codes {
		if strings.TrimSpace(code.Family) == "" {
			return EngineError{Code: ErrorDataBundleInvalid, Message: "code resolution family is required", Details: map[string]any{"index": index}}
		}
		if _, ok := allowedResolutionStatuses[code.Status]; !ok {
			return EngineError{Code: ErrorDataBundleInvalid, Message: "code resolution status is invalid", Details: map[string]any{"index": index, "status": code.Status}}
		}
		if code.Status == StatusNotApplicable && strings.TrimSpace(code.Basis) == "" {
			return EngineError{Code: ErrorDataBundleInvalid, Message: "not_applicable requires an affirmative basis", Details: map[string]any{"index": index, "family": code.Family}}
		}
	}
	for index, requirement := range result.Requirements {
		if strings.TrimSpace(requirement.ID) == "" || strings.TrimSpace(requirement.Prompt) == "" {
			return EngineError{Code: ErrorDataBundleInvalid, Message: "resolution requirement requires id and prompt", Details: map[string]any{"index": index}}
		}
		if _, ok := allowedRequirementKinds[requirement.Kind]; !ok {
			return EngineError{Code: ErrorDataBundleInvalid, Message: "resolution requirement kind is invalid", Details: map[string]any{"index": index, "kind": requirement.Kind}}
		}
		if requirement.Kind == RequirementProjectFact && strings.TrimSpace(requirement.FactKey) == "" {
			return EngineError{Code: ErrorDataBundleInvalid, Message: "project_fact requirement requires fact_key", Details: map[string]any{"index": index}}
		}
	}
	return nil
}
