package geocoder

// Canonical source-identity intake for the public candidate-only geocoder.
// The authoritative shape is arcata-jurisdictions/contracts/jurisdiction-identity.schema.json
// (paired ODCS v3.1.0). This subset is deliberately narrower than that schema:
// reviewed identity, geometry, dated authority and other source kinds are rejected,
// not silently downgraded into a government or project verdict.

import (
    "bufio"
    "bytes"
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "strings"
    "time"
)

type canonicalPeriod struct {
    Start *string `json:"start"`
    EndExclusive *string `json:"end_exclusive"`
    Basis string `json:"basis"`
}
type canonicalIdentifier struct {
    System string `json:"system"`
    Value string `json:"value"`
}
type canonicalMatch struct {
    Status string `json:"status"`
    CandidateGovernmentEntityID string `json:"candidate_government_entity_id,omitempty"`
}
type canonicalGovernment struct {
    PublisherName string `json:"publisher_name"`
    Kind string `json:"kind"`
    PublisherIdentifiers []canonicalIdentifier `json:"publisher_identifiers"`
    Match canonicalMatch `json:"match"`
}
type canonicalObservation struct {
    ObservationID string `json:"observation_id"`
    Kind string `json:"kind"`
    SourceReceiptID string `json:"source_receipt_id"`
    SourceRecordKey string `json:"source_record_key"`
    ObservedAt string `json:"observed_at"`
    ProvenanceEntity string `json:"provenance_entity"`
}
type canonicalUncertainty struct {
    State string `json:"state"`
    Reasons []string `json:"reasons"`
    Notes string `json:"notes"`
}
type canonicalProvenance struct {
    GraphURI string `json:"graph_uri"`
    EntityID string `json:"entity_id"`
}
type canonicalLegalAuthority struct {
    Status string `json:"status"`
}
type canonicalIdentity struct {
    SchemaVersion string `json:"schema_version"`
    RecordID string `json:"record_id"`
    ScopeID string `json:"scope_id"`
    Government canonicalGovernment `json:"government"`
    EffectivePeriod canonicalPeriod `json:"effective_period"`
    Boundaries []json.RawMessage `json:"boundaries"`
    Observations []canonicalObservation `json:"observations"`
    Relations []json.RawMessage `json:"relations"`
    Uncertainty canonicalUncertainty `json:"uncertainty"`
    Provenance canonicalProvenance `json:"provenance"`
    LegalAuthority canonicalLegalAuthority `json:"legal_authority"`
}

// rejectDuplicateJSONKeys traverses nested objects, not just top-level keys.
func rejectDuplicateJSONKeys(raw []byte) error {
    d := json.NewDecoder(bytes.NewReader(raw))
    var visit func() error
    visit = func() error {
        tok, err := d.Token()
        if err != nil { return err }
        opener, isDelimiter := tok.(json.Delim)
        if !isDelimiter { return nil }
        switch opener {
        case '{':
            seen := map[string]bool{}
            for d.More() {
                keyTok, err := d.Token()
                if err != nil { return err }
                key, ok := keyTok.(string)
                if !ok || seen[key] { return errors.New("duplicate canonical JSON field") }
                seen[key] = true
                if err := visit(); err != nil { return err }
            }
            closer, err := d.Token()
            if err != nil || closer != json.Delim('}') { return errors.New("invalid JSON object") }
        case '[':
            for d.More() { if err := visit(); err != nil { return err } }
            closer, err := d.Token()
            if err != nil || closer != json.Delim(']') { return errors.New("invalid JSON array") }
        default:
            return errors.New("unexpected JSON terminator")
        }
        return nil
    }
    if err := visit(); err != nil { return err }
    if _, err := d.Token(); err != io.EOF { return errors.New("trailing canonical JSON") }
    return nil
}

func decodeCanonicalCandidate(raw []byte, cfg PlaceCandidateSource) (PlaceCandidateObservation, error) {
    var out PlaceCandidateObservation
    if err := rejectDuplicateJSONKeys(raw); err != nil { return out, err }
    var r canonicalIdentity
    d := json.NewDecoder(bytes.NewReader(raw))
    d.DisallowUnknownFields()
    if err := d.Decode(&r); err != nil { return out, fmt.Errorf("canonical JSON Schema subset rejected: %w", err) }
    if r.SchemaVersion != "1.0.0" || !strings.HasPrefix(r.RecordID, "source-view:") ||
        r.ScopeID != cfg.ScopeID || r.Government.Kind != "municipality" ||
        strings.TrimSpace(r.Government.PublisherName) == "" ||
        r.LegalAuthority.Status != "not_established" ||
        r.EffectivePeriod.Basis != "unknown" || r.EffectivePeriod.Start != nil ||
        r.EffectivePeriod.EndExclusive != nil ||
        len(r.Boundaries) != 0 || len(r.Relations) != 0 ||
        len(r.Observations) != 1 ||
        r.Uncertainty.State != "unresolved" ||
        strings.TrimSpace(r.Uncertainty.Notes) == "" {
        return out, errors.New("canonical government cannot authorize reviewed identity, dates, boundaries or legal authority")
    }
    required := map[string]bool{"legal_authority_unreviewed": false, "boundary_missing": false, "effective_date_unknown": false}
    seenReason := map[string]bool{}
    for _, reason := range r.Uncertainty.Reasons {
        if seenReason[reason] { return out, errors.New("duplicate uncertainty reason") }
        seenReason[reason] = true
        if _, ok := required[reason]; ok { required[reason] = true }
    }
    for _, ok := range required { if !ok { return out, errors.New("missing source uncertainty") } }
    if len(r.Government.PublisherIdentifiers) != 1 ||
        r.Government.PublisherIdentifiers[0].System != "co-dola-candidate-place" ||
        !fipsPlaceCode.MatchString(r.Government.PublisherIdentifiers[0].Value) {
        return out, errors.New("unrecognized publisher place identifier")
    }
    match := r.Government.Match
    switch match.Status {
    case "candidate":
        if strings.TrimSpace(match.CandidateGovernmentEntityID) == "" || !seenReason["candidate_crosswalk"] {
            return out, errors.New("candidate crosswalk requires explicit unresolved ID")
        }
    case "unmatched":
        if match.CandidateGovernmentEntityID != "" || seenReason["candidate_crosswalk"] {
            return out, errors.New("unmatched publisher cannot assert a crosswalk")
        }
    default:
        return out, errors.New("reviewed or conflicted identity cannot enter candidate-only Colorado intake")
    }
    observation := r.Observations[0]
    if observation.SourceReceiptID != cfg.SourceReceiptID ||
        observation.Kind != "government_identity" ||
        observation.ObservationID == "" || observation.SourceRecordKey == "" ||
        r.RecordID != "source-view:"+observation.ObservationID ||
        observation.ProvenanceEntity == "" ||
        r.Provenance.EntityID != observation.ProvenanceEntity ||
        r.Provenance.GraphURI != "urn:bcm:co:dola:source-observation-graph:2026-05-04" {
        return out, errors.New("canonical observation and W3C PROV source witness disagree")
    }
    if _, err := time.Parse(time.RFC3339Nano, observation.ObservedAt); err != nil {
        return out, errors.New("canonical observation lacks valid source timestamp")
    }
    out = PlaceCandidateObservation{
        SchemaVersion: "1.0.0", ObservationID: observation.ObservationID,
        Kind: "government_identity", ScopeID: r.ScopeID,
        SourceReceiptID: observation.SourceReceiptID, SourceRecordKey: observation.SourceRecordKey,
        ObservedAt: observation.ObservedAt, ProvenanceEntity: observation.ProvenanceEntity,
        ReviewState: "candidate_only",
        PublisherCode: r.Government.PublisherIdentifiers[0].Value,
        PublisherName: r.Government.PublisherName,
        CandidateGovernmentEntityID: match.CandidateGovernmentEntityID,
    }
    if match.Status == "candidate" { out.IdentityMatchBasis = "publisher_code_candidate" } else {
        out.IdentityMatchBasis = "unmatched"
    }
    return out, nil
}

// ReadJurisdictionIdentityCandidates requires a trusted, immutable SHA-256
// and complete count, then consumes canonical identity records only. The
// returned index is an existing source-candidate view, never legal identity.
func ReadJurisdictionIdentityCandidates(reader io.Reader, source PlaceCandidateSource) (*PlaceCandidateIndex, error) {
    if !exactSHA256.MatchString(source.SHA256) || !usStateScope.MatchString(source.ScopeID) ||
        source.ScopeID != "US-CO" || strings.TrimSpace(source.SourceReceiptID) == "" ||
        source.ExpectedRows < 1 || source.ExpectedRows > maxPlaceCandidateRecords {
        return nil, errors.New("trusted canonical identity source pin required")
    }
    raw, err := io.ReadAll(io.LimitReader(reader, maxPlaceCandidateBytes+1))
    if err != nil { return nil, err }
    if len(raw) == 0 || len(raw) > maxPlaceCandidateBytes || raw[len(raw)-1] != '\n' {
        return nil, errors.New("canonical identity file incomplete or oversized")
    }
    digest := sha256.Sum256(raw)
    if hex.EncodeToString(digest[:]) != source.SHA256 {
        return nil, errors.New("canonical identity export digest differs from trusted lock")
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
    seenIDs, seenKeys := map[string]bool{}, map[string]bool{}
    scan := bufio.NewScanner(bytes.NewReader(raw))
    scan.Buffer(make([]byte, 65536), 1<<20)
    for scan.Scan() {
        if idx.summary.ObservationCount >= source.ExpectedRows { return nil, errors.New("extra canonical identity record") }
        row, err := decodeCanonicalCandidate(scan.Bytes(), source)
        if err != nil { return nil, fmt.Errorf("canonical identity row %d: %w", idx.summary.ObservationCount+1, err) }
        if seenIDs[row.ObservationID] || seenKeys[row.SourceRecordKey] {
            return nil, errors.New("duplicate canonical source observation")
        }
        seenIDs[row.ObservationID], seenKeys[row.SourceRecordKey] = true, true
        idx.byCode[row.PublisherCode] = append(idx.byCode[row.PublisherCode], row)
        idx.byName[normalizeWords(row.PublisherName)] = append(idx.byName[normalizeWords(row.PublisherName)], row)
        idx.summary.ObservationCount++
        if row.CandidateGovernmentEntityID != "" { idx.summary.CodeCandidateCount++ } else { idx.summary.UnmatchedCount++ }
    }
    if err := scan.Err(); err != nil { return nil, err }
    if idx.summary.ObservationCount != source.ExpectedRows { return nil, errors.New("truncated canonical identity export") }
    return idx, nil
}
