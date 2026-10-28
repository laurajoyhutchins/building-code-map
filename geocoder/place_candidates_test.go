package geocoder

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strings"
	"testing"
)

func candidateFixture() ([]byte, PlaceCandidateSource) {
	rows := []PlaceCandidateObservation{
		{
			SchemaVersion: "1.0.0", ObservationID: "obs:1",
			Kind: "government_identity", ScopeID: "US-CO",
			SourceReceiptID: "private:synthetic:receipt", SourceRecordKey: "1",
			ObservedAt: "2026-10-21T00:48:20.738Z", ProvenanceEntity: "urn:example:row:1",
			ReviewState: "candidate_only", PublisherCode: "03782", PublisherName: "Hudson",
			IdentityMatchBasis: "unmatched",
		},
		{
			SchemaVersion: "1.0.0", ObservationID: "obs:2",
			Kind: "government_identity", ScopeID: "US-CO",
			SourceReceiptID: "private:synthetic:receipt", SourceRecordKey: "2",
			ObservedAt: "2026-10-21T00:48:20.738Z", ProvenanceEntity: "urn:example:row:2",
			ReviewState: "candidate_only", PublisherCode: "37820", PublisherName: "Hudson",
			IdentityMatchBasis: "publisher_code_candidate",
			CandidateGovernmentEntityID: "gov:synthetic:2",
		},
		{
			SchemaVersion: "1.0.0", ObservationID: "obs:3",
			Kind: "government_identity", ScopeID: "US-CO",
			SourceReceiptID: "private:synthetic:receipt", SourceRecordKey: "3",
			ObservedAt: "2026-10-21T00:48:20.738Z", ProvenanceEntity: "urn:example:row:3",
			ReviewState: "candidate_only", PublisherCode: "20000", PublisherName: "Denver",
			IdentityMatchBasis: "publisher_code_candidate",
			CandidateGovernmentEntityID: "gov:synthetic:3",
		},
	}
	var raw []byte
	for _, row := range rows {
		payload, err := json.Marshal(row)
		if err != nil {
			panic(err)
		}
		raw = append(raw, append(payload, '\n')...)
	}
	digest := sha256.Sum256(raw)
	return raw, PlaceCandidateSource{
		SHA256: hex.EncodeToString(digest[:]), ScopeID: "US-CO",
		SourceReceiptID: "private:synthetic:receipt", ExpectedRows: 3,
	}
}

func TestSourceOnlyColoradoCandidateImportAndAddressLookup(t *testing.T) {
	raw, source := candidateFixture()
	idx, err := ReadPlaceCandidates(bytes.NewReader(raw), source)
	if err != nil {
		t.Fatal(err)
	}
	summary := idx.Summary()
	if summary.ObservationCount != 3 || summary.CodeCandidateCount != 2 ||
		summary.UnmatchedCount != 1 || summary.AuthorityStatus != "not_established" ||
		summary.LegalVerdictAuthorized {
		t.Fatalf("candidate-only summary corrupted: %+v", summary)
	}
	unknown := idx.ByPlaceCode("03782")
	if unknown.Status != "unmatched" || len(unknown.Observations) != 1 ||
		unknown.LegalVerdictAuthorized {
		t.Fatalf("unmatched source was promoted: %+v", unknown)
	}
	candidate := idx.ByPlaceCode("37820")
	if candidate.Status != "candidate_only" || len(candidate.Observations) != 1 ||
		candidate.LegalVerdictAuthorized {
		t.Fatalf("source-only code candidate changed: %+v", candidate)
	}
	hudson, err := ParseAddress("120 Main St, Hudson, CO")
	if err != nil {
		t.Fatal(err)
	}
	ambiguous := idx.ForAddress(hudson)
	if ambiguous.Status != "ambiguous" || len(ambiguous.Observations) != 2 ||
		ambiguous.Observations[0].PublisherCode == ambiguous.Observations[1].PublisherCode ||
		ambiguous.LegalVerdictAuthorized {
		t.Fatalf("distinct source Hudson identities collapsed: %+v", ambiguous)
	}
	denver, err := ParseAddress("100 Main St, Denver, CO")
	if err != nil {
		t.Fatal(err)
	}
	name := idx.ForAddress(denver)
	if name.Status != "candidate_only" || len(name.Observations) != 1 ||
		name.LegalVerdictAuthorized {
		t.Fatalf("single municipality name incorrectly granted authority: %+v", name)
	}
	wyoming, err := ParseAddress("100 Main St, Denver, WY")
	if err != nil {
		t.Fatal(err)
	}
	if result := idx.ForAddress(wyoming); result.Status != "unmatched" {
		t.Fatalf("cross-state civic city lookup was accepted: %+v", result)
	}
	if result := idx.ByPlaceCode("00000"); result.Status != "unmatched" || len(result.Observations) != 0 {
		t.Fatalf("unrepresented place code incorrectly matched: %+v", result)
	}
	candidate.Observations[0].PublisherName = "Changed"
	if got := idx.ByPlaceCode("37820"); got.Observations[0].PublisherName != "Hudson" {
		t.Fatal("caller can mutate the retained source candidate index")
	}
}

func TestSourceOnlyCandidateRejections(t *testing.T) {
	base, config := candidateFixture()
	cases := []struct {
		name   string
		change func([]byte, *PlaceCandidateSource) []byte
	}{
		{"missing trusted digest", func(b []byte, c *PlaceCandidateSource) []byte { c.SHA256 = ""; return b }},
		{"bad digest", func(b []byte, c *PlaceCandidateSource) []byte { c.SHA256 = strings.Repeat("0", 64); return b }},
		{"missing scope", func(b []byte, c *PlaceCandidateSource) []byte { c.ScopeID = ""; return b }},
		{"wrong publisher receipt", func(b []byte, c *PlaceCandidateSource) []byte { c.SourceReceiptID = "other"; return b }},
		{"incomplete count", func(b []byte, c *PlaceCandidateSource) []byte { c.ExpectedRows = 2; return b }},
		{"too many expected", func(b []byte, c *PlaceCandidateSource) []byte { c.ExpectedRows = 4; return b }},
		{"no terminal newline", func(b []byte, c *PlaceCandidateSource) []byte { return b[:len(b)-1] }},
		{"modified source", func(b []byte, c *PlaceCandidateSource) []byte { return append(b, '\n') }},
		{"legal authority field", func(b []byte, c *PlaceCandidateSource) []byte {
			return bytes.Replace(b, []byte(`"publisher_code":"03782"`), []byte(`"publisher_code":"03782","government_entity_id":"gov:synthetic:2"`), 1)
		}},
		{"duplicate source JSON field", func(b []byte, c *PlaceCandidateSource) []byte {
			return bytes.Replace(b, []byte(`"review_state":"candidate_only"`), []byte(`"review_state":"candidate_only","review_state":"candidate_only"`), 1)
		}},
		{"reviewed mapping substitution", func(b []byte, c *PlaceCandidateSource) []byte {
			return bytes.Replace(b, []byte(`"identity_match_basis":"publisher_code_candidate"`), []byte(`"identity_match_basis":"reviewed_mapping"`), 1)
		}},
		{"false verified review state", func(b []byte, c *PlaceCandidateSource) []byte {
			return bytes.Replace(b, []byte(`"review_state":"candidate_only"`), []byte(`"review_state":"verified"`), 1)
		}},
		{"cross-state observation", func(b []byte, c *PlaceCandidateSource) []byte {
			return bytes.Replace(b, []byte(`"scope_id":"US-CO"`), []byte(`"scope_id":"US-KS"`), 1)
		}},
		{"invalid date", func(b []byte, c *PlaceCandidateSource) []byte {
			return bytes.Replace(b, []byte(`"observed_at":"2026-10-21T00:48:20.738Z"`), []byte(`"observed_at":"yesterday"`), 1)
		}},
		{"mismatched candidate source basis", func(b []byte, c *PlaceCandidateSource) []byte {
			return bytes.Replace(b, []byte(`"identity_match_basis":"unmatched"`), []byte(`"identity_match_basis":"publisher_code_candidate"`), 1)
		}},
		{"duplicate source record key", func(b []byte, c *PlaceCandidateSource) []byte {
			return bytes.Replace(b, []byte(`"source_record_key":"2"`), []byte(`"source_record_key":"1"`), 1)
		}},
		{"duplicate observation id", func(b []byte, c *PlaceCandidateSource) []byte {
			return bytes.Replace(b, []byte(`"observation_id":"obs:2"`), []byte(`"observation_id":"obs:1"`), 1)
		}},
		{"duplicate place source cannot become legal government", func(b []byte, c *PlaceCandidateSource) []byte {
			return bytes.Replace(b, []byte(`"publisher_code":"37820"`), []byte(`"publisher_code":"03782"`), 1)
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			b := append([]byte{}, base...)
			cfg := config
			b = tc.change(b, &cfg)
			// For semantic negatives, recompute only the transport checksum to
			// prove the validation is not just a changed-byte hash rejection.
			if cfg.SHA256 == config.SHA256 && !bytes.Equal(b, base) {
				digest := sha256.Sum256(b)
				cfg.SHA256 = hex.EncodeToString(digest[:])
			}
			_, err := ReadPlaceCandidates(bytes.NewReader(b), cfg)
			if tc.name == "duplicate place source cannot become legal government" {
				// Duplicate publisher codes are retained as ambiguous, never
				// silently deduplicated or used as accepted jurisdiction.
				if err != nil {
					t.Fatal(err)
				}
				idx, _ := ReadPlaceCandidates(bytes.NewReader(b), cfg)
				if got := idx.ByPlaceCode("03782"); got.Status != "ambiguous" || len(got.Observations) != 2 {
					t.Fatalf("duplicate publisher code was silently accepted: %+v", got)
				}
				return
			}
			if err == nil {
				t.Fatal("untrusted geographic source was accepted")
			}
		})
	}
}

func TestSourceOnlyConflictedPublisherRecordStaysUnresolved(t *testing.T) {
	raw, cfg := candidateFixture()
	// Conflicted records can be staged, but may not carry candidate government ID.
	raw = bytes.Replace(raw, []byte(`"review_state":"candidate_only","publisher_code":"03782"`),
		[]byte(`"review_state":"conflicted","publisher_code":"03782"`), 1)
	raw = bytes.Replace(raw, []byte(`"identity_match_basis":"unmatched"`),
		[]byte(`"identity_match_basis":"conflicted"`), 1)
	digest := sha256.Sum256(raw)
	cfg.SHA256 = hex.EncodeToString(digest[:])
	idx, err := ReadPlaceCandidates(bytes.NewReader(raw), cfg)
	if err != nil {
		t.Fatal(err)
	}
	if found := idx.ByPlaceCode("03782"); found.Status != "conflicted" || found.LegalVerdictAuthorized {
		t.Fatalf("conflicted row illegally promoted: %+v", found)
	}
}
