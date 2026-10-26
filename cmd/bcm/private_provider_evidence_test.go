package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/laurajoyhutchins/building-code-map/engine"
	"github.com/laurajoyhutchins/building-code-map/httpapi"
)

const syntheticArtifactHash = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
const syntheticPassageHash = "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"

var syntheticFamilies = []string{"building", "electrical", "energy", "fire_operational", "residential"}

func exactLegacyLink(family string) engine.EvidenceLink {
	id := "DEMO-" + family
	return engine.EvidenceLink{
		ID: id + "-LINK", ClaimID: id + "-CLAIM", Role: engine.EvidenceEstablishes,
		Document: engine.SourceDocument{ID: id + "-DOCUMENT", Kind: "synthetic_ordinance",
			JurisdictionID: "DEMO-XX"},
		Artifact: engine.SourceArtifact{ID: id + "-ARTIFACT", DocumentID: id + "-DOCUMENT",
			SHA256: syntheticArtifactHash, SizeBytes: 123},
		Anchor: engine.TextAnchor{ID: id + "-ANCHOR", ArtifactID: id + "-ARTIFACT",
			TextSHA256: syntheticPassageHash, LegalLocator: "Fictional section 1"},
	}
}

func exactLegacyPolicy() engine.SourcePolicy {
	policy := engine.SourcePolicy{}
	for _, family := range syntheticFamilies {
		link := exactLegacyLink(family)
		policy.Rules = append(policy.Rules, engine.SourcePolicyRule{
			ID: "DEMO-POLICY-" + family, ClaimKind: "code_adoption",
			JurisdictionID: "DEMO-XX", DocumentID: link.Document.ID,
			ArtifactSHA256: syntheticArtifactHash, TextSHA256: syntheticPassageHash,
			LegalLocator:  "Fictional section 1",
			EffectiveFrom: "2026-01-01", EffectiveTo: "2026-12-31",
		})
	}
	return policy
}

type exactLegacyObservation struct {
	omitEvidence  bool
	mutatePassage bool
	localRecord   bool
}

// Synthetic transport fixture; the private legacy API still resolves one
// family per request and returns observation records, not a project verdict.
func mockExactLegacyProvider(t *testing.T, scenario exactLegacyObservation) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/v1/readiness":
			_, _ = w.Write([]byte(`{"status":"ok","capabilities":{"address_lookup":{"status":"available"}}}`))
		case "/v1/bundle":
			_, _ = w.Write([]byte(`{"bundle_manifest_digest":"sha256:synthetic"}`))
		case "/v1/resolve":
			var query struct {
				Address           string `json:"address"`
				CodeFamily        string `json:"code_family"`
				ApplicabilityDate string `json:"applicability_date"`
			}
			if err := json.NewDecoder(r.Body).Decode(&query); err != nil {
				t.Error(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			if query.Address != demoAddress {
				t.Errorf("wrong project address: %q", query.Address)
			}
			link := exactLegacyLink(query.CodeFamily)
			if scenario.mutatePassage {
				link.Anchor.TextSHA256 = strings.Repeat("c", 64)
			}
			links := []engine.EvidenceLink{link}
			if scenario.omitEvidence {
				links = nil
			}
			required := []string{}
			if scenario.localRecord {
				required = []string{"Check fictional amendment record"}
			}
			out := map[string]any{
				"schema_version": "1",
				"resolution": map[string]any{
					"status":             "verified",
					"applicability_date": query.ApplicabilityDate,
					"geography": map[string]any{
						"jurisdiction_id":              "DEMO-XX",
						"jurisdiction_name":            "Fictional jurisdiction",
						"census_state_equivalent_fips": "08",
					},
					"adoptions": []any{map[string]any{
						"code_family":    query.CodeFamily,
						"code_name":      "DEMO-CODE-" + query.CodeFamily,
						"verification":   map[string]any{"status": "verified"},
						"source_ids":     []string{"DEMO-SOURCE-" + query.CodeFamily},
						"exact_evidence": links,
					}},
					"required_local_records": required,
				},
				"provenance": map[string]any{
					"engine_version":         "DEMO-ENGINE",
					"source_commit":          strings.Repeat("d", 40),
					"bundle_manifest_digest": "sha256:synthetic",
				},
			}
			if err := json.NewEncoder(w).Encode(out); err != nil {
				t.Error(err)
			}
		default:
			http.NotFound(w, r)
		}
	}))
}

func exactLegacyProject(t *testing.T, scenario exactLegacyObservation, policy engine.SourcePolicy,
	asOf string, facts map[string]string) (engine.ProjectCodeBasis, error) {
	t.Helper()
	server := mockExactLegacyProvider(t, scenario)
	defer server.Close()
	provider, err := newPrivateProvider(server.URL, "")
	if err != nil {
		t.Fatal(err)
	}
	provider.expectedBundle = "sha256:synthetic"
	return engine.NewProjectVerifier(engine.NewRuntimeWithSourcePolicy(provider, policy)).
		VerifyProject(context.Background(), engine.ProjectRequest{
			ProjectID: "DEMO-PROJECT", Address: demoAddress,
			ApplicabilityDate: asOf, Facts: facts,
		})
}

func TestPrivateExactEvidenceUsesCanonicalProjectVerifier(t *testing.T) {
	// The private provider can furnish an exact witness, but only an
	// independently configured policy and Runtime may accept that witness.
	result, err := exactLegacyProject(t, exactLegacyObservation{}, exactLegacyPolicy(), "2026-10-25", nil)
	if err != nil {
		t.Fatal(err)
	}
	if result.Verdict != engine.ProjectVerified || len(result.CodeSet) != 5 ||
		len(result.ExactEvidence) != 5 || result.Provenance.BundleID != "sha256:synthetic" {
		t.Fatalf("source-exact project did not round-trip: %#v", result)
	}
	for _, code := range result.CodeSet {
		if code.Status != engine.StatusResolved {
			t.Fatalf("source-exact adoption not accepted for %q: %s", code.Family, code.Status)
		}
	}
}

func TestPrivateExactEvidenceFailsClosedForMissingOrUntrustedWitness(t *testing.T) {
	for _, tt := range []struct {
		name        string
		scenario    exactLegacyObservation
		policy      engine.SourcePolicy
		date        string
		facts       map[string]string
		wantVerdict engine.ProjectVerdict
		wantError   bool
	}{
		{name: "no exact evidence remains conditional", scenario: exactLegacyObservation{omitEvidence: true},
			policy: exactLegacyPolicy(), date: "2026-10-25", wantVerdict: engine.ProjectConditional},
		{name: "unapproved source", policy: engine.SourcePolicy{}, date: "2026-10-25", wantError: true},
		{name: "tampered passage", scenario: exactLegacyObservation{mutatePassage: true},
			policy: exactLegacyPolicy(), date: "2026-10-25", wantError: true},
		{name: "outside the approved effective interval", policy: exactLegacyPolicy(),
			date: "2027-01-01", wantError: true},
		{name: "pending local record remains conditional", scenario: exactLegacyObservation{localRecord: true},
			policy: exactLegacyPolicy(), date: "2026-10-25", wantVerdict: engine.ProjectConditional},
		{name: "unreviewed project facts remain conditional", policy: exactLegacyPolicy(),
			date: "2026-10-25", facts: map[string]string{"occupancy": "business"},
			wantVerdict: engine.ProjectConditional},
	} {
		t.Run(tt.name, func(t *testing.T) {
			result, err := exactLegacyProject(t, tt.scenario, tt.policy, tt.date, tt.facts)
			if tt.wantError {
				if got := engine.NormalizeError(err).Code; got != engine.ErrorDataBundleInvalid {
					t.Fatalf("wanted data_bundle_invalid, got %v", err)
				}
				return
			}
			if err != nil || result.Verdict != tt.wantVerdict {
				t.Fatalf("result = %#v; err = %v, want %s", result, err, tt.wantVerdict)
			}
		})
	}
}

func TestPrivateExactEvidenceSurvivesOnlyCanonicalHTTPPath(t *testing.T) {
	server := mockExactLegacyProvider(t, exactLegacyObservation{})
	defer server.Close()
	provider, err := newPrivateProvider(server.URL, "")
	if err != nil {
		t.Fatal(err)
	}
	provider.expectedBundle = "sha256:synthetic"
	handler := handlerWithSourcePolicy(provider, provider, "", nil, exactLegacyPolicy())
	req := httptest.NewRequest(http.MethodPost, httpapi.ProjectVerificationPath,
		strings.NewReader(`{"project_id":"DEMO-PROJECT","address":"100 Demo Plaza, DEMO-XX","applicability_date":"2026-10-25"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("canonical HTTP response: status=%d body=%s", rec.Code, rec.Body.String())
	}
	var result engine.ProjectCodeBasis
	if err := json.Unmarshal(rec.Body.Bytes(), &result); err != nil {
		t.Fatal(err)
	}
	if result.Verdict != engine.ProjectVerified || len(result.ExactEvidence) != 5 {
		t.Fatalf("lost verified evidence at HTTP boundary: %#v", result)
	}
}
