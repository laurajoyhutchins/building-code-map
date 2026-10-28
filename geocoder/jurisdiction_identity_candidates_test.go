package geocoder

import (
    "bytes"
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "strings"
    "testing"
)

func identityFixture(t *testing.T) (map[string]any, PlaceCandidateSource) {
    t.Helper()
    observation := map[string]any{
        "observation_id": "src:synthetic:1", "kind": "government_identity",
        "source_receipt_id": "receipt:synthetic", "source_record_key": "1",
        "observed_at": "2026-10-21T00:48:20.738Z", "provenance_entity": "urn:example:entity:1",
    }
    r := map[string]any{
        "schema_version": "1.0.0", "record_id": "source-view:src:synthetic:1", "scope_id": "US-CO",
        "government": map[string]any{
            "publisher_name": "Hudson", "kind": "municipality",
            "publisher_identifiers": []any{map[string]any{"system": "co-dola-candidate-place", "value": "37820"}},
            "match": map[string]any{"status": "candidate", "candidate_government_entity_id": "gov:synthetic:2"},
        },
        "effective_period": map[string]any{"start": nil, "end_exclusive": nil, "basis": "unknown"},
        "boundaries": []any{}, "observations": []any{observation}, "relations": []any{},
        "uncertainty": map[string]any{
            "state": "unresolved",
            "reasons": []any{"legal_authority_unreviewed", "boundary_missing", "effective_date_unknown", "candidate_crosswalk"},
            "notes": "The Census place identifier is a source candidate, never legal authority.",
        },
        "provenance": map[string]any{
            "graph_uri": "urn:bcm:co:dola:source-observation-graph:2026-05-04",
            "entity_id": "urn:example:entity:1",
        },
        "legal_authority": map[string]any{"status": "not_established"},
    }
    return r, PlaceCandidateSource{ScopeID: "US-CO", SourceReceiptID: "receipt:synthetic", ExpectedRows: 1}
}

func ingestCanonical(t *testing.T, fixture map[string]any, source PlaceCandidateSource) (*PlaceCandidateIndex, error) {
    t.Helper()
    raw, err := json.Marshal(fixture)
    if err != nil { t.Fatal(err) }
    raw = append(raw, '\n')
    sum := sha256.Sum256(raw)
    source.SHA256 = hex.EncodeToString(sum[:])
    return ReadJurisdictionIdentityCandidates(bytes.NewReader(raw), source)
}

func TestCanonicalIdentitySourceOnlyLookup(t *testing.T) {
    rec, config := identityFixture(t)
    idx, err := ingestCanonical(t, rec, config)
    if err != nil { t.Fatal(err) }
    got := idx.ByPlaceCode("37820")
    if got.Status != "candidate_only" || got.LegalVerdictAuthorized ||
        len(got.Observations) != 1 || idx.Summary().CodeCandidateCount != 1 ||
        idx.Summary().AuthorityStatus != "not_established" {
        t.Fatalf("source identity was promoted or lost: %+v %+v", got, idx.Summary())
    }
    address, err := ParseAddress("100 Main St, Hudson, CO")
    if err != nil { t.Fatal(err) }
    if idx.ForAddress(address).Status != "candidate_only" { t.Fatal("source name lookup missing") }
    // A source-only unmatched governmental observation must remain address-unverified.
    gov := rec["government"].(map[string]any)
    gov["match"] = map[string]any{"status": "unmatched"}
    rec["uncertainty"].(map[string]any)["reasons"] = []any{
        "legal_authority_unreviewed", "boundary_missing", "effective_date_unknown",
    }
    idx, err = ingestCanonical(t, rec, config)
    if err != nil || idx.Summary().UnmatchedCount != 1 ||
        idx.ByPlaceCode("37820").Status != "unmatched" {
        t.Fatalf("unmatched identity source incorrectly promoted: %v %+v", err, idx)
    }
}

func TestCanonicalIdentityRejectsHostileSourceChanges(t *testing.T) {
    tests := []struct {
        name string
        modify func(map[string]any)
    }{
        {"false authority", func(r map[string]any) { r["legal_authority"] = map[string]any{"status":"verified"} }},
        {"forged project verdict", func(r map[string]any) { r["project_verdict"] = "verified" }},
        {"promoted reviewed correspondence", func(r map[string]any) {
            r["government"].(map[string]any)["match"] = map[string]any{"status":"reviewed_correspondence", "reviewed_government_entity_id":"gov:123"}
        }},
        {"invented effective date", func(r map[string]any) { r["effective_period"].(map[string]any)["start"] = "2026-01-01" }},
        {"fictitious geometry", func(r map[string]any) { r["boundaries"] = []any{map[string]any{"boundary_id":"fake"}} }},
        {"forged PROV root", func(r map[string]any) { r["provenance"].(map[string]any)["entity_id"] = "urn:fake" }},
        {"missing source receipt", func(r map[string]any) {
            r["observations"].([]any)[0].(map[string]any)["source_receipt_id"] = "other"
        }},
        {"missing uncertainty", func(r map[string]any) {
            r["uncertainty"].(map[string]any)["reasons"] = []any{"candidate_crosswalk"}
        }},
        {"missing crosswalk", func(r map[string]any) {
            r["government"].(map[string]any)["match"] = map[string]any{"status":"candidate"}
        }},
        {"unmatched hidden crosswalk", func(r map[string]any) {
            r["government"].(map[string]any)["match"] = map[string]any{
                "status":"unmatched","candidate_government_entity_id":"gov:synthetic:2"}
        }},
        {"unreviewed relation", func(r map[string]any) { r["relations"] = []any{map[string]any{"relation":"part_of"}} }},
        {"unknown nested authority", func(r map[string]any) {
            r["government"].(map[string]any)["legal_primary_jurisdiction"] = "gov:synthetic:2"
        }},
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            rec, cfg := identityFixture(t)
            tc.modify(rec)
            if _, err := ingestCanonical(t, rec, cfg); err == nil {
                t.Fatal("canonical intake accepted unsupported source assertion")
            }
        })
    }
}

func TestCanonicalIdentityRejectsNestedDuplicateKeysAndTampering(t *testing.T) {
    rec, cfg := identityFixture(t)
    raw, err := json.Marshal(rec)
    if err != nil { t.Fatal(err) }
    raw = bytes.Replace(raw, []byte(`"status":"candidate"`),
        []byte(`"status":"candidate","status":"reviewed_correspondence"`), 1)
    if !bytes.Contains(raw, []byte(`"status":"candidate","status":`)) {
        t.Fatal("negative duplicate field not constructed")
    }
    raw = append(raw, '\n')
    digest := sha256.Sum256(raw)
    cfg.SHA256 = hex.EncodeToString(digest[:])
    if _, err := ReadJurisdictionIdentityCandidates(bytes.NewReader(raw), cfg); err == nil {
        t.Fatal("duplicate nested identity field was accepted")
    }
    rec, cfg = identityFixture(t)
    _, err = ingestCanonical(t, rec, cfg)
    if err != nil { t.Fatal(err) }
    payload, _ := json.Marshal(rec)
    payload = append(payload, '\n')
    sha := sha256.Sum256(payload)
    cfg.SHA256 = hex.EncodeToString(sha[:])
    cfg.ExpectedRows = 2
    if _, err := ReadJurisdictionIdentityCandidates(bytes.NewReader(payload), cfg); err == nil {
        t.Fatal("partial source export accepted")
    }
    cfg.ExpectedRows = 1
    cfg.SHA256 = strings.Repeat("0", 64)
    if _, err := ReadJurisdictionIdentityCandidates(bytes.NewReader(payload), cfg); err == nil {
        t.Fatal("untrusted source digest accepted")
    }
}
