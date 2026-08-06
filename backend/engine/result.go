package engine

import (
	"sort"
	"time"
)

type Result struct {
	SchemaVersion string          `json:"schema_version"`
	Query         NormalizedQuery `json:"query"`
	Location      LocationResult  `json:"location"`
	Resolution    Resolution      `json:"resolution"`
	Provenance    Provenance      `json:"provenance"`
	Diagnostics   []Diagnostic    `json:"diagnostics"`
}

type LocationResult struct {
	Point   *Point         `json:"point,omitempty"`
	Address string         `json:"address,omitempty"`
	Geocode *GeocodeResult `json:"geocode,omitempty"`
}

type GeocodeResult struct {
	Query      string             `json:"query"`
	Normalized string             `json:"normalized,omitempty"`
	Status     string             `json:"status"`
	Selected   *GeocodeCandidate  `json:"selected,omitempty"`
	Candidates []GeocodeCandidate `json:"candidates"`
	Warnings   []string           `json:"warnings"`
}

type GeocodeCandidate struct {
	MatchedAddress       string                 `json:"matched_address"`
	Longitude            float64                `json:"longitude"`
	Latitude             float64                `json:"latitude"`
	Precision            string                 `json:"precision"`
	Confidence           float64                `json:"confidence"`
	ScoreKind            string                 `json:"score_kind"`
	ScoreFactors         map[string]float64     `json:"score_factors"`
	RankingPolicyVersion string                 `json:"ranking_policy_version"`
	Source               string                 `json:"source"`
	SourceRecordID       string                 `json:"source_record_id"`
	SourceVintage        string                 `json:"source_vintage"`
	Interpolation        *InterpolationEvidence `json:"interpolation,omitempty"`
}

type InterpolationEvidence struct {
	SourceRangeID             string     `json:"source_range_id"`
	RequestedHouseNumber      int        `json:"requested_house_number"`
	FromNumber                int        `json:"from_number"`
	ToNumber                  int        `json:"to_number"`
	RangeDirection            string     `json:"range_direction"`
	Parity                    string     `json:"parity"`
	Side                      string     `json:"side,omitempty"`
	FromCoordinate            Coordinate `json:"from_coordinate"`
	ToCoordinate              Coordinate `json:"to_coordinate"`
	Fraction                  float64    `json:"fraction"`
	DerivedCoordinate         Coordinate `json:"derived_coordinate"`
	CoordinateReferenceSystem string     `json:"coordinate_reference_system"`
	TransformationIdentity    string     `json:"transformation_identity"`
	MethodVersion             string     `json:"method_version"`
	PositionalQuality         string     `json:"positional_quality"`
}

type Coordinate struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type Resolution struct {
	SchemaVersion        string                  `json:"schema_version"`
	GeneratedAt          time.Time               `json:"generated_at"`
	ProfileID            string                  `json:"profile_id,omitempty"`
	ProfileLastVerified  string                  `json:"profile_last_verified,omitempty"`
	Geography            GeographicContext       `json:"geography"`
	CodeFamily           string                  `json:"code_family,omitempty"`
	ProjectType          string                  `json:"project_type,omitempty"`
	ApplicabilityDate    string                  `json:"applicability_date"`
	Status               string                  `json:"status"`
	PolicyBasis          *PolicyBasis            `json:"policy_basis,omitempty"`
	AuthorityCandidates  []AuthorityCandidate    `json:"authority_candidates"`
	AuthorityPath        []AuthorityRelationship `json:"authority_path"`
	Adoptions            []Adoption              `json:"adoptions"`
	ApplicableRules      []RuleReference         `json:"applicable_rules"`
	SupportingClaims     []Claim                 `json:"supporting_claims"`
	RequiredLocalRecords []string                `json:"required_local_records"`
	Warnings             []string                `json:"warnings"`
	Evidence             []Source                `json:"evidence"`
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
	SpecialAreas      []BoundaryMatch `json:"special_areas"`
	TribalAreas       []BoundaryMatch `json:"tribal_areas"`
	FireJurisdictions []BoundaryMatch `json:"fire_jurisdictions"`
}

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

type PolicyBasis struct {
	Status               string       `json:"status"`
	RequiredLocalRecords []string     `json:"required_local_records"`
	Warnings             []string     `json:"warnings"`
	SourceIDs            []string     `json:"source_ids"`
	Verification         Verification `json:"verification"`
}

type AuthorityCandidate struct {
	Kind         string       `json:"kind"`
	AuthorityID  string       `json:"authority_id,omitempty"`
	Name         string       `json:"name"`
	Roles        []string     `json:"roles"`
	SourceIDs    []string     `json:"source_ids"`
	Verification Verification `json:"verification"`
}

type AuthorityRelationship struct {
	ID           string       `json:"id"`
	FromID       string       `json:"from_id"`
	Relationship string       `json:"relationship"`
	To           string       `json:"to"`
	Scope        []string     `json:"scope"`
	Summary      string       `json:"summary,omitempty"`
	SourceIDs    []string     `json:"source_ids"`
	Verification Verification `json:"verification"`
}

type Adoption struct {
	ID               string            `json:"id"`
	CodeFamily       string            `json:"code_family"`
	Status           string            `json:"status"`
	StateCodeName    string            `json:"state_code_name"`
	EnforcementModel string            `json:"enforcement_model"`
	Dates            map[string]string `json:"dates"`
	SourceIDs        []string          `json:"source_ids"`
	Verification     Verification      `json:"verification"`
}

type RuleReference struct {
	ID           string       `json:"id"`
	Kind         string       `json:"kind"`
	CodeFamily   string       `json:"code_family,omitempty"`
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

type Provenance struct {
	SourceCommit            string `json:"source_commit"`
	EngineVersion           string `json:"engine_version"`
	BundleManifestDigest    string `json:"bundle_manifest_digest"`
	BoundarySnapshotDigest  string `json:"boundary_snapshot_digest"`
	RegulatoryCatalogDigest string `json:"regulatory_catalog_digest"`
	GeocoderSnapshotDigest  string `json:"geocoder_snapshot_digest,omitempty"`
}

type Diagnostic struct {
	Severity string `json:"severity"`
	Code     string `json:"code"`
	Message  string `json:"message"`
	Path     string `json:"path,omitempty"`
}

func SortDiagnostics(diagnostics []Diagnostic) {
	sort.SliceStable(diagnostics, func(i, j int) bool {
		left := diagnostics[i]
		right := diagnostics[j]
		if left.Severity != right.Severity {
			return left.Severity < right.Severity
		}
		if left.Code != right.Code {
			return left.Code < right.Code
		}
		if left.Path != right.Path {
			return left.Path < right.Path
		}
		return left.Message < right.Message
	})
}
