package regulatory

import "time"

const (
	ProfileSchemaVersion = "1.0"
	ResultSchemaVersion  = "1.0"
)

type Verification struct {
	Status     string  `json:"status"`
	Confidence float64 `json:"confidence,omitempty"`
	Notes      string  `json:"notes,omitempty"`
}

type Source struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	URL           string `json:"url"`
	Kind          string `json:"kind"`
	AccessedAt    string `json:"accessed_at"`
	LastCheckedAt string `json:"last_checked_at,omitempty"`
	Availability  string `json:"availability,omitempty"`
	Caveat        string `json:"caveat,omitempty"`
}

type Authority struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Type         string       `json:"type"`
	Roles        []string     `json:"roles"`
	SourceIDs    []string     `json:"source_ids"`
	Verification Verification `json:"verification"`
}

type AuthorityRelationship struct {
	ID           string       `json:"id"`
	FromID       string       `json:"from_id"`
	Relationship string       `json:"relationship"`
	To           string       `json:"to"`
	Scope        []string     `json:"scope,omitempty"`
	Summary      string       `json:"summary,omitempty"`
	SourceIDs    []string     `json:"source_ids"`
	Verification Verification `json:"verification"`
}

type ModelCode struct {
	Publisher   string `json:"publisher,omitempty"`
	Name        string `json:"name,omitempty"`
	EditionYear int    `json:"edition_year,omitempty"`
}

type AdoptionDates struct {
	AdoptionDate    string `json:"adoption_date,omitempty"`
	EffectiveDate   string `json:"effective_date,omitempty"`
	OperativeDate   string `json:"operative_date,omitempty"`
	MandatoryDate   string `json:"mandatory_date,omitempty"`
	ReplacementDate string `json:"replacement_date,omitempty"`
}

type Adoption struct {
	ID                  string        `json:"id"`
	CodeFamily          string        `json:"code_family"`
	Status              string        `json:"status"`
	StateCodeName       string        `json:"state_code_name"`
	BaseModelCode       *ModelCode    `json:"base_model_code,omitempty"`
	AdoptingAuthorityID string        `json:"adopting_authority_id"`
	EnforcementModel    string        `json:"enforcement_model"`
	AppliesTo           []string      `json:"applies_to,omitempty"`
	Exclusions          []string      `json:"exclusions,omitempty"`
	Dates               AdoptionDates `json:"dates"`
	SourceIDs           []string      `json:"source_ids"`
	Verification        Verification  `json:"verification"`
}

type ApplicabilityRule struct {
	ID           string       `json:"id"`
	CodeFamily   string       `json:"code_family,omitempty"`
	ProjectTypes []string     `json:"project_types,omitempty"`
	Trigger      string       `json:"trigger"`
	AppliesTo    []string     `json:"applies_to,omitempty"`
	Exclusions   []string     `json:"exclusions,omitempty"`
	Summary      string       `json:"summary"`
	SourceIDs    []string     `json:"source_ids"`
	Verification Verification `json:"verification"`
}

type DateRule struct {
	ID               string       `json:"id"`
	CodeFamily       string       `json:"code_family,omitempty"`
	RuleType         string       `json:"rule_type"`
	Trigger          string       `json:"trigger"`
	StartDate        string       `json:"start_date,omitempty"`
	EndDate          string       `json:"end_date,omitempty"`
	PriorCodeAllowed *bool        `json:"prior_code_allowed,omitempty"`
	Summary          string       `json:"summary"`
	SourceIDs        []string     `json:"source_ids"`
	Verification     Verification `json:"verification"`
}

type AmendmentRule struct {
	ID           string       `json:"id"`
	CodeFamily   string       `json:"code_family,omitempty"`
	Level        string       `json:"level"`
	AuthorityID  string       `json:"authority_id,omitempty"`
	Posture      string       `json:"posture"`
	Summary      string       `json:"summary"`
	SourceIDs    []string     `json:"source_ids"`
	Verification Verification `json:"verification"`
}

type EnforcementRule struct {
	ID           string       `json:"id"`
	CodeFamily   string       `json:"code_family,omitempty"`
	Model        string       `json:"model"`
	EntityKinds  []string     `json:"entity_kinds"`
	AuthorityIDs []string     `json:"authority_ids,omitempty"`
	Summary      string       `json:"summary"`
	SourceIDs    []string     `json:"source_ids"`
	Verification Verification `json:"verification"`
}

type Claim struct {
	ID            string       `json:"id"`
	SubjectID     string       `json:"subject_id"`
	Field         string       `json:"field"`
	Status        string       `json:"status"`
	Value         any          `json:"value,omitempty"`
	ConflictGroup string       `json:"conflict_group,omitempty"`
	SourceIDs     []string     `json:"source_ids"`
	Verification  Verification `json:"verification"`
}

type CandidateRule struct {
	Kind        string   `json:"kind"`
	AuthorityID string   `json:"authority_id,omitempty"`
	Label       string   `json:"label,omitempty"`
	Roles       []string `json:"roles,omitempty"`
	SourceIDs   []string `json:"source_ids"`
}

type ResolutionPolicy struct {
	ReplaceAuthorityCandidates bool            `json:"replace_authority_candidates,omitempty"`
	ReplaceAdoptions           bool            `json:"replace_adoptions,omitempty"`
	Status                     string          `json:"status"`
	AuthorityCandidates        []CandidateRule `json:"authority_candidates,omitempty"`
	AdoptionIDs                []string        `json:"adoption_ids,omitempty"`
	RequiredLocalRecords       []string        `json:"required_local_records,omitempty"`
	Warnings                   []string        `json:"warnings,omitempty"`
	SourceIDs                  []string        `json:"source_ids"`
}

type DefaultPolicies struct {
	Incorporated   ResolutionPolicy `json:"incorporated"`
	Unincorporated ResolutionPolicy `json:"unincorporated"`
}

type StateProfile struct {
	SchemaVersion        string                      `json:"schema_version"`
	ProfileID            string                      `json:"profile_id"`
	StateID              string                      `json:"state_id"`
	StateName            string                      `json:"state_name"`
	StateAbbreviation    string                      `json:"state_abbreviation"`
	StateFIPS            string                      `json:"state_fips"`
	Status               string                      `json:"status"`
	LastVerified         string                      `json:"last_verified"`
	Sources              []Source                    `json:"sources"`
	Authorities          []Authority                 `json:"authorities"`
	Relationships        []AuthorityRelationship     `json:"relationships,omitempty"`
	Adoptions            []Adoption                  `json:"adoptions,omitempty"`
	ApplicabilityRules   []ApplicabilityRule         `json:"applicability_rules,omitempty"`
	DateRules            []DateRule                  `json:"date_rules,omitempty"`
	AmendmentRules       []AmendmentRule             `json:"amendment_rules,omitempty"`
	EnforcementRules     []EnforcementRule           `json:"enforcement_rules,omitempty"`
	Claims               []Claim                     `json:"claims,omitempty"`
	ResolverFixtureIDs   []string                    `json:"resolver_fixture_ids,omitempty"`
	Defaults             DefaultPolicies             `json:"defaults"`
	CodeFamilyOverrides  map[string]ResolutionPolicy `json:"code_family_overrides,omitempty"`
	ProjectTypeOverrides map[string]ResolutionPolicy `json:"project_type_overrides,omitempty"`
	Verification         Verification                `json:"verification"`
}

type BoundaryMatch struct {
	LayerFamily string `json:"layer_family"`
	FeatureID   string `json:"feature_id"`
	Name        string `json:"name"`
	SourceID    string `json:"source_id,omitempty"`
}

type GeographicContext struct {
	StateID           string          `json:"state_id,omitempty"`
	StateFIPS         string          `json:"state_fips,omitempty"`
	StateName         string          `json:"state_name,omitempty"`
	County            *BoundaryMatch  `json:"county,omitempty"`
	Municipality      *BoundaryMatch  `json:"municipality,omitempty"`
	Incorporated      bool            `json:"incorporated"`
	SpecialAreas      []BoundaryMatch `json:"special_areas,omitempty"`
	TribalAreas       []BoundaryMatch `json:"tribal_areas,omitempty"`
	FireJurisdictions []BoundaryMatch `json:"fire_jurisdictions,omitempty"`
}

type Point struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type ResolutionRequest struct {
	Point             *Point             `json:"point,omitempty"`
	Context           *GeographicContext `json:"context,omitempty"`
	CodeFamily        string             `json:"code_family,omitempty"`
	ProjectType       string             `json:"project_type,omitempty"`
	ApplicabilityDate string             `json:"applicability_date,omitempty"`
}

type AuthorityCandidate struct {
	Kind         string       `json:"kind"`
	AuthorityID  string       `json:"authority_id,omitempty"`
	Name         string       `json:"name"`
	Roles        []string     `json:"roles,omitempty"`
	SourceIDs    []string     `json:"source_ids"`
	Verification Verification `json:"verification"`
}

type RuleReference struct {
	ID           string       `json:"id"`
	Kind         string       `json:"kind"`
	CodeFamily   string       `json:"code_family,omitempty"`
	Summary      string       `json:"summary"`
	SourceIDs    []string     `json:"source_ids"`
	Verification Verification `json:"verification"`
}

type PolicyBasis struct {
	Status               string       `json:"status"`
	RequiredLocalRecords []string     `json:"required_local_records,omitempty"`
	Warnings             []string     `json:"warnings,omitempty"`
	SourceIDs            []string     `json:"source_ids"`
	Verification         Verification `json:"verification"`
}

type ResolutionResult struct {
	SchemaVersion        string                  `json:"schema_version"`
	GeneratedAt          time.Time               `json:"generated_at"`
	ProfileID            string                  `json:"profile_id"`
	ProfileLastVerified  string                  `json:"profile_last_verified"`
	Geography            GeographicContext       `json:"geography"`
	CodeFamily           string                  `json:"code_family,omitempty"`
	ProjectType          string                  `json:"project_type,omitempty"`
	ApplicabilityDate    string                  `json:"applicability_date,omitempty"`
	Status               string                  `json:"status"`
	PolicyBasis          *PolicyBasis            `json:"policy_basis,omitempty"`
	AuthorityCandidates  []AuthorityCandidate    `json:"authority_candidates"`
	AuthorityPath        []AuthorityRelationship `json:"authority_path,omitempty"`
	Adoptions            []Adoption              `json:"adoptions,omitempty"`
	ApplicableRules      []RuleReference         `json:"applicable_rules,omitempty"`
	SupportingClaims     []Claim                 `json:"supporting_claims,omitempty"`
	RequiredLocalRecords []string                `json:"required_local_records,omitempty"`
	Warnings             []string                `json:"warnings,omitempty"`
	Evidence             []Source                `json:"evidence"`
}
