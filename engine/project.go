package engine

import (
	"context"
	"sort"
	"strings"
)

const ProjectCodeBasisSchemaVersion = "0.3"

type ProjectVerdict string

const (
	ProjectVerified    ProjectVerdict = "verified"
	ProjectConditional ProjectVerdict = "conditional"
	ProjectNotVerified ProjectVerdict = "not_verified"
)

type ProjectRequest struct {
	ProjectID         string            `json:"project_id"`
	Point             *Point            `json:"point,omitempty"`
	Address           string            `json:"address,omitempty"`
	ApplicabilityDate string            `json:"applicability_date"`
	ProjectType       string            `json:"project_type,omitempty"`
	Facts             map[string]string `json:"facts,omitempty"`
}

type ProjectInputs struct {
	Point       *Point            `json:"point,omitempty"`
	Address     string            `json:"address,omitempty"`
	ProjectType string            `json:"project_type,omitempty"`
	Facts       map[string]string `json:"facts,omitempty"`
}

type ProjectCodeEntry struct {
	Family  string           `json:"family"`
	Edition string           `json:"edition,omitempty"`
	Status  ResolutionStatus `json:"status"`
	Basis   string           `json:"basis,omitempty"`
}

type ProjectUnresolved struct {
	Kind    RequirementKind `json:"kind"`
	Message string          `json:"message"`
	FactKey string          `json:"fact_key,omitempty"`
	Code    string          `json:"code_family,omitempty"`
}

type ProjectCodeBasis struct {
	SchemaVersion     string              `json:"schema_version"`
	ProjectID         string              `json:"project_id"`
	AsOf              string              `json:"as_of"`
	Inputs            ProjectInputs       `json:"inputs"`
	JurisdictionID    string              `json:"jurisdiction_id"`
	Jurisdiction      JurisdictionContext `json:"jurisdiction"`
	Verdict           ProjectVerdict      `json:"verdict"`
	CodeSet           []ProjectCodeEntry  `json:"code_set"`
	UserFacts         map[string]string   `json:"user_facts,omitempty"`
	DerivedFacts      []DerivedFact       `json:"derived_facts,omitempty"`
	Unresolved        []ProjectUnresolved `json:"unresolved"`
	FollowUpQuestions []string            `json:"follow_up_questions"`
	Evidence          []EvidenceRef       `json:"evidence"`
	ExactEvidence     []EvidenceLink      `json:"exact_evidence"`
	Provenance        Provenance          `json:"provenance"`
}

type ProjectVerifier struct{ runtime Runtime }

func NewProjectVerifier(runtime Runtime) ProjectVerifier { return ProjectVerifier{runtime: runtime} }

func (verifier ProjectVerifier) VerifyProject(ctx context.Context, request ProjectRequest) (ProjectCodeBasis, error) {
	projectID := strings.TrimSpace(request.ProjectID)
	if projectID == "" {
		return ProjectCodeBasis{}, EngineError{Code: ErrorInvalidQuery, Message: "project_id is required for project verification"}
	}

	resolution, err := verifier.runtime.Resolve(ctx, Query{
		Point:             request.Point,
		Address:           request.Address,
		ProjectType:       request.ProjectType,
		ProjectFacts:      request.Facts,
		ApplicabilityDate: request.ApplicabilityDate,
	})
	if err != nil {
		return ProjectCodeBasis{}, err
	}

	codeSet := make([]ProjectCodeEntry, 0, len(resolution.Codes))
	unresolved := make([]ProjectUnresolved, 0, len(resolution.Requirements))
	evidence := make([]EvidenceRef, 0)
	evidenceSeen := map[string]struct{}{}
	exactEvidence := make([]EvidenceLink, 0)
	exactEvidenceSeen := map[string]struct{}{}
	followUps := make([]string, 0)
	verdict := ProjectVerified

	for _, requirement := range resolution.Requirements {
		unresolved = append(unresolved, ProjectUnresolved{Kind: requirement.Kind, Message: requirement.Prompt, FactKey: requirement.FactKey})
		if requirement.Kind == RequirementProjectFact {
			followUps = append(followUps, requirement.Prompt)
		}
		if verdict == ProjectVerified {
			verdict = ProjectConditional
		}
	}

	for _, code := range resolution.Codes {
		codeSet = append(codeSet, ProjectCodeEntry{Family: code.Family, Edition: code.Edition, Status: code.Status, Basis: code.Basis})
		switch code.Status {
		case StatusUnsupported, StatusInsufficientEvidence, StatusAmbiguous, StatusConflicting:
			verdict = ProjectNotVerified
		case StatusPartiallyResolved, StatusLocalRecordRequired:
			if verdict == ProjectVerified {
				verdict = ProjectConditional
			}
		}
		kind := unresolvedKindForStatus(code.Status)
		for _, message := range code.Unresolved {
			message = strings.TrimSpace(message)
			if message == "" {
				continue
			}
			unresolved = append(unresolved, ProjectUnresolved{Kind: kind, Message: message, Code: code.Family})
		}
		for _, ref := range code.Evidence {
			key := ref.Kind + "\x00" + ref.ID
			if _, ok := evidenceSeen[key]; ok {
				continue
			}
			evidenceSeen[key] = struct{}{}
			evidence = append(evidence, ref)
		}
		for _, link := range code.ExactEvidence {
			if _, ok := exactEvidenceSeen[link.ID]; ok {
				continue
			}
			exactEvidenceSeen[link.ID] = struct{}{}
			exactEvidence = append(exactEvidence, link)
		}
	}

	sort.Slice(codeSet, func(i, j int) bool {
		if codeSet[i].Family == codeSet[j].Family {
			return codeSet[i].Edition < codeSet[j].Edition
		}
		return codeSet[i].Family < codeSet[j].Family
	})
	sort.Slice(evidence, func(i, j int) bool {
		if evidence[i].Kind == evidence[j].Kind {
			return evidence[i].ID < evidence[j].ID
		}
		return evidence[i].Kind < evidence[j].Kind
	})
	sort.Slice(exactEvidence, func(i, j int) bool { return exactEvidence[i].ID < exactEvidence[j].ID })
	sort.Strings(followUps)

	return ProjectCodeBasis{
		SchemaVersion: ProjectCodeBasisSchemaVersion,
		ProjectID:     projectID,
		AsOf:          resolution.Query.ApplicabilityDate,
		Inputs: ProjectInputs{
			Point: resolution.Query.Point, Address: resolution.Query.Address, ProjectType: resolution.Query.ProjectType, Facts: cloneFacts(resolution.Query.ProjectFacts),
		},
		JurisdictionID:    resolution.Jurisdiction.PrimaryJurisdictionID,
		Jurisdiction:      resolution.Jurisdiction,
		Verdict:           verdict,
		CodeSet:           codeSet,
		UserFacts:         cloneFacts(resolution.Query.ProjectFacts),
		DerivedFacts:      append([]DerivedFact(nil), resolution.DerivedFacts...),
		Unresolved:        unresolved,
		FollowUpQuestions: followUps,
		Evidence:          evidence,
		ExactEvidence:     exactEvidence,
		Provenance:        resolution.Provenance,
	}, nil
}

func unresolvedKindForStatus(status ResolutionStatus) RequirementKind {
	switch status {
	case StatusLocalRecordRequired:
		return RequirementLocalRecord
	case StatusPartiallyResolved:
		return RequirementProjectFact
	default:
		return RequirementEvidenceDefect
	}
}

func cloneFacts(values map[string]string) map[string]string {
	if len(values) == 0 {
		return nil
	}
	cloned := make(map[string]string, len(values))
	for key, value := range values {
		cloned[key] = value
	}
	return cloned
}
