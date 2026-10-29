package geocoder

import "regexp"

// The geographic source-candidate consumer is deliberately NOT a municipal
// jurisdiction resolver. In particular, candidate_government_entity_id must
// never be substituted for the legal primary_jurisdiction_id used by BCM's
// canonical Runtime and ProjectVerifier.
const maxPlaceCandidateBytes = 16 << 20
const maxPlaceCandidateRecords = 100000

var usStateScope = regexp.MustCompile(`^US-[A-Z]{2}$`)
var exactSHA256 = regexp.MustCompile(`^[0-9a-f]{64}$`)
var fipsPlaceCode = regexp.MustCompile(`^[0-9]{5}$`)

// PlaceCandidateSource is selected independently of the untrusted input.
// A digest obtained by hashing the same unreviewed file is not a trusted pin.
type PlaceCandidateSource struct {
	SHA256          string
	ScopeID         string
	SourceReceiptID string
	ExpectedRows    int
}

// PlaceCandidateObservation is the candidate-only projection of the
// canonical jurisdiction-identity JSON Schema + ODCS contract.
// This is not a legal government object, address-bound polygon, or adoption.
type PlaceCandidateObservation struct {
	SchemaVersion               string `json:"schema_version"`
	ObservationID               string `json:"observation_id"`
	Kind                        string `json:"kind"`
	ScopeID                     string `json:"scope_id"`
	SourceReceiptID             string `json:"source_receipt_id"`
	SourceRecordKey             string `json:"source_record_key"`
	ObservedAt                  string `json:"observed_at"`
	ProvenanceEntity            string `json:"provenance_entity"`
	ReviewState                 string `json:"review_state"`
	PublisherCode               string `json:"publisher_code"`
	PublisherName               string `json:"publisher_name"`
	IdentityMatchBasis          string `json:"identity_match_basis"`
	CandidateGovernmentEntityID string `json:"candidate_government_entity_id,omitempty"`
}

type PlaceCandidateSummary struct {
	SchemaVersion          string `json:"schema_version"`
	ScopeID                string `json:"scope_id"`
	SourceReceiptID        string `json:"source_receipt_id"`
	SourceSHA256           string `json:"source_sha256"`
	ObservationCount       int    `json:"observation_count"`
	CodeCandidateCount     int    `json:"code_candidate_count"`
	UnmatchedCount         int    `json:"unmatched_count"`
	ConflictedCount        int    `json:"conflicted_count"`
	AuthorityStatus        string `json:"authority_status"`
	LegalVerdictAuthorized bool   `json:"legal_verdict_authorized"`
}

// PlaceCandidateMatch is a source-only lookup. Even one source-code match
// cannot establish an address/parcel jurisdiction, enforcement body, or law.
type PlaceCandidateMatch struct {
	Status                 string                      `json:"status"`
	Observations           []PlaceCandidateObservation `json:"observations"`
	LegalVerdictAuthorized bool                        `json:"legal_verdict_authorized"`
}

type PlaceCandidateIndex struct {
	byCode  map[string][]PlaceCandidateObservation
	byName  map[string][]PlaceCandidateObservation
	summary PlaceCandidateSummary
}

func (idx *PlaceCandidateIndex) Summary() PlaceCandidateSummary {
	return idx.summary
}

func placeCandidatesFor(rows []PlaceCandidateObservation) PlaceCandidateMatch {
	out := PlaceCandidateMatch{
		Status:       "unmatched",
		Observations: append([]PlaceCandidateObservation{}, rows...),
	}
	switch {
	case len(rows) > 1:
		out.Status = "ambiguous"
	case len(rows) == 1 && rows[0].ReviewState == "conflicted":
		out.Status = "conflicted"
	case len(rows) == 1 && rows[0].CandidateGovernmentEntityID != "":
		out.Status = "candidate_only"
	}
	return out
}

// ByPlaceCode never performs a name-based alias or auto-admits government IDs.
func (idx *PlaceCandidateIndex) ByPlaceCode(placeCode string) PlaceCandidateMatch {
	return placeCandidatesFor(idx.byCode[placeCode])
}

// ForAddress accepts the already-normalized civic address from the existing
// geocoder, but matches ONLY the municipality display name. It cannot infer
// site-boundary membership or legal authority from that text.
func (idx *PlaceCandidateIndex) ForAddress(address ParsedAddress) PlaceCandidateMatch {
	if "US-"+address.State != idx.summary.ScopeID || address.City == "" {
		return placeCandidatesFor(nil)
	}
	return placeCandidatesFor(idx.byName[normalizeWords(address.City)])
}
