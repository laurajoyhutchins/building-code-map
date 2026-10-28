package geocoder

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"testing"
)

// These cases test the existing source-only candidate index through its
// canonical identity input, not the retired parallel geographic JSONL decoder.
func twoCanonicalHudsons(t *testing.T, duplicateCode bool) ([]byte, PlaceCandidateSource) {
	t.Helper()
	first, cfg := identityFixture(t)
	payload, err := json.Marshal(first)
	if err != nil {
		t.Fatal(err)
	}
	second, _ := identityFixture(t)
	observation := second["observations"].([]any)[0].(map[string]any)
	observation["observation_id"] = "src:synthetic:2"
	observation["source_record_key"] = "2"
	observation["provenance_entity"] = "urn:example:entity:2"
	second["record_id"] = "source-view:src:synthetic:2"
	second["provenance"].(map[string]any)["entity_id"] = "urn:example:entity:2"
	second["government"].(map[string]any)["match"] = map[string]any{"status": "unmatched"}
	second["uncertainty"].(map[string]any)["reasons"] = []any{
		"legal_authority_unreviewed", "boundary_missing", "effective_date_unknown",
	}
	if !duplicateCode {
		second["government"].(map[string]any)["publisher_identifiers"].([]any)[0].(map[string]any)["value"] = "03782"
	}
	row2, err := json.Marshal(second)
	if err != nil {
		t.Fatal(err)
	}
	raw := append(append(append([]byte{}, payload...), '\n'), row2...)
	raw = append(raw, '\n')
	digest := sha256.Sum256(raw)
	cfg.SHA256 = hex.EncodeToString(digest[:])
	cfg.ExpectedRows = 2
	return raw, cfg
}

func TestCanonicalCandidatesRetainDistinctSameNameSources(t *testing.T) {
	raw, cfg := twoCanonicalHudsons(t, false)
	idx, err := ReadJurisdictionIdentityCandidates(bytes.NewReader(raw), cfg)
	if err != nil {
		t.Fatal(err)
	}
	summary := idx.Summary()
	if summary.ObservationCount != 2 || summary.CodeCandidateCount != 1 ||
		summary.UnmatchedCount != 1 || summary.LegalVerdictAuthorized ||
		summary.AuthorityStatus != "not_established" {
		t.Fatalf("canonical candidate count or authority changed: %+v", summary)
	}
	address, err := ParseAddress("120 Main St, Hudson, CO")
	if err != nil {
		t.Fatal(err)
	}
	ambiguous := idx.ForAddress(address)
	if ambiguous.Status != "ambiguous" || len(ambiguous.Observations) != 2 ||
		ambiguous.LegalVerdictAuthorized {
		t.Fatalf("name ambiguity lost: %+v", ambiguous)
	}
	if got := idx.ByPlaceCode("37820"); got.Status != "candidate_only" {
		t.Fatalf("candidate lost: %+v", got)
	}
	if got := idx.ByPlaceCode("03782"); got.Status != "unmatched" {
		t.Fatalf("unmatched promoted: %+v", got)
	}
	other, err := ParseAddress("120 Main St, Hudson, WY")
	if err != nil {
		t.Fatal(err)
	}
	if got := idx.ForAddress(other); got.Status != "unmatched" {
		t.Fatalf("cross-state match admitted: %+v", got)
	}
	if got := idx.ByPlaceCode("00000"); got.Status != "unmatched" || len(got.Observations) != 0 {
		t.Fatalf("unrepresented source code matched: %+v", got)
	}
	ambiguous.Observations[0].PublisherName = "Tampered"
	if got := idx.ForAddress(address); got.Observations[0].PublisherName != "Hudson" {
		t.Fatal("caller mutation changed canonical candidate index")
	}
}

func TestCanonicalCandidatesRetainDuplicatePlaceCodesAsAmbiguous(t *testing.T) {
	raw, cfg := twoCanonicalHudsons(t, true)
	idx, err := ReadJurisdictionIdentityCandidates(bytes.NewReader(raw), cfg)
	if err != nil {
		t.Fatal(err)
	}
	got := idx.ByPlaceCode("37820")
	if got.Status != "ambiguous" || len(got.Observations) != 2 || got.LegalVerdictAuthorized {
		t.Fatalf("duplicate source code was collapsed: %+v", got)
	}
}

func TestLegacyGeographicJSONLDoesNotEnterCanonicalConsumer(t *testing.T) {
	legacy := []byte("{\"schema_version\":\"1.0.0\",\"review_state\":\"candidate_only\"}\n")
	digest := sha256.Sum256(legacy)
	_, err := ReadJurisdictionIdentityCandidates(bytes.NewReader(legacy), PlaceCandidateSource{
		SHA256: hex.EncodeToString(digest[:]), ScopeID: "US-CO",
		SourceReceiptID: "receipt:synthetic", ExpectedRows: 1,
	})
	if err == nil {
		t.Fatal("legacy noncanonical geographic JSONL was accepted")
	}
}
