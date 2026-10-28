package geocoder

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"
	"time"
)

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
// cross-repository geographic-observation JSON Schema + ODCS contract.
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

// Reject duplicate JSON keys. encoding/json by itself silently accepts a
// repeated legal or identity field and chooses the final value.
func requireUniqueCandidateFields(raw []byte) error {
	dec := json.NewDecoder(bytes.NewReader(raw))
	open, err := dec.Token()
	if err != nil || open != json.Delim('{') {
		return errors.New("candidate observation must be a JSON object")
	}
	seen := make(map[string]bool)
	for dec.More() {
		token, err := dec.Token()
		if err != nil {
			return errors.New("invalid candidate key")
		}
		key, ok := token.(string)
		if !ok || seen[key] {
			return errors.New("duplicate or invalid candidate key")
		}
		seen[key] = true
		var value json.RawMessage
		if err := dec.Decode(&value); err != nil {
			return errors.New("invalid candidate field")
		}
	}
	end, err := dec.Token()
	if err != nil || end != json.Delim('}') {
		return errors.New("candidate JSON object is incomplete")
	}
	if _, err := dec.Token(); err != io.EOF {
		return errors.New("trailing candidate JSON content")
	}
	return nil
}

func decodeCandidate(raw []byte, source PlaceCandidateSource) (PlaceCandidateObservation, error) {
	var row PlaceCandidateObservation
	if err := requireUniqueCandidateFields(raw); err != nil {
		return row, err
	}
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&row); err != nil {
		return row, errors.New("candidate does not satisfy supported geography field contract")
	}
	_, timeErr := time.Parse(time.RFC3339Nano, row.ObservedAt)
	if row.SchemaVersion != "1.0.0" || row.Kind != "government_identity" ||
		row.ScopeID != source.ScopeID || row.SourceReceiptID != source.SourceReceiptID ||
		strings.TrimSpace(row.ObservationID) == "" ||
		strings.TrimSpace(row.SourceRecordKey) == "" ||
		strings.TrimSpace(row.PublisherName) == "" ||
		!fipsPlaceCode.MatchString(row.PublisherCode) ||
		timeErr != nil ||
		!(strings.HasPrefix(row.ProvenanceEntity, "urn:") ||
			strings.HasPrefix(row.ProvenanceEntity, "https://")) {
		return row, errors.New("candidate observation has invalid source identity or contract fields")
	}
	if row.ReviewState != "candidate_only" && row.ReviewState != "conflicted" {
		return row, errors.New("source review cannot establish a legal verdict")
	}
	switch row.IdentityMatchBasis {
	case "publisher_code_candidate":
		if row.CandidateGovernmentEntityID == "" || row.ReviewState != "candidate_only" {
			return row, errors.New("candidate code crosswalk needs explicit candidate-only identity")
		}
	case "unmatched", "conflicted":
		if row.CandidateGovernmentEntityID != "" {
			return row, errors.New("unmatched or conflicted source must not contain an identity join")
		}
	default:
		return row, errors.New("reviewed legal identity cannot enter a source-only candidate index")
	}
	return row, nil
}

// ReadPlaceCandidates independently checks an explicitly trusted SHA-256,
// expected scope, receipt identity and row count before building a read-only
// lookup. It rejects authority fields, duplicate keys and malformed records.
// Source rights, PROV graph closure and full schema conformance remain separate
// producer/CI obligations; successful ingestion is never legal acceptance.
func ReadPlaceCandidates(reader io.Reader, source PlaceCandidateSource) (*PlaceCandidateIndex, error) {
	if !exactSHA256.MatchString(source.SHA256) ||
		!usStateScope.MatchString(source.ScopeID) ||
		strings.TrimSpace(source.SourceReceiptID) == "" ||
		source.ExpectedRows < 1 || source.ExpectedRows > maxPlaceCandidateRecords {
		return nil, errors.New("trusted geography source identity and row count are required")
	}
	raw, err := io.ReadAll(io.LimitReader(reader, maxPlaceCandidateBytes+1))
	if err != nil {
		return nil, fmt.Errorf("cannot read geographic source: %w", err)
	}
	if len(raw) > maxPlaceCandidateBytes || len(raw) == 0 || raw[len(raw)-1] != '\n' {
		return nil, errors.New("geography candidate file too large, empty or missing terminal newline")
	}
	digest := sha256.Sum256(raw)
	if hex.EncodeToString(digest[:]) != source.SHA256 {
		return nil, errors.New("candidate source SHA-256 mismatch")
	}
	idx := &PlaceCandidateIndex{
		byCode: make(map[string][]PlaceCandidateObservation),
		byName: make(map[string][]PlaceCandidateObservation),
		summary: PlaceCandidateSummary{
			SchemaVersion: "1.0.0", ScopeID: source.ScopeID,
			SourceReceiptID: source.SourceReceiptID, SourceSHA256: source.SHA256,
			AuthorityStatus: "not_established",
		},
	}
	seenID := make(map[string]bool)
	seenSourceKey := make(map[string]bool)
	scanner := bufio.NewScanner(bytes.NewReader(raw))
	scanner.Buffer(make([]byte, 65536), 1<<20)
	for scanner.Scan() {
		if idx.summary.ObservationCount >= source.ExpectedRows {
			return nil, errors.New("source contains more records than trusted count")
		}
		row, err := decodeCandidate(scanner.Bytes(), source)
		if err != nil {
			return nil, fmt.Errorf("candidate row %d rejected: %w", idx.summary.ObservationCount+1, err)
		}
		if seenID[row.ObservationID] || seenSourceKey[row.SourceRecordKey] {
			return nil, errors.New("duplicate source observation or publisher record key")
		}
		seenID[row.ObservationID], seenSourceKey[row.SourceRecordKey] = true, true
		idx.byCode[row.PublisherCode] = append(idx.byCode[row.PublisherCode], row)
		idx.byName[normalizeWords(row.PublisherName)] = append(idx.byName[normalizeWords(row.PublisherName)], row)
		idx.summary.ObservationCount++
		if row.CandidateGovernmentEntityID != "" {
			idx.summary.CodeCandidateCount++
		} else {
			idx.summary.UnmatchedCount++
		}
		if row.ReviewState == "conflicted" {
			idx.summary.ConflictedCount++
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, errors.New("oversized or invalid candidate observation record")
	}
	if idx.summary.ObservationCount != source.ExpectedRows {
		return nil, errors.New("incomplete geographic source observations")
	}
	return idx, nil
}
