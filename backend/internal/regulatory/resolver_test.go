package regulatory

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestResolveColoradoBuildingRequiresLocalAdoptionRecord(t *testing.T) {
	catalog := loadTestCatalog(t)
	result, err := Resolve(catalog, ResolutionRequest{
		Context: &GeographicContext{
			StateID:      "US-CO",
			StateFIPS:    "08",
			StateName:    "Colorado",
			Incorporated: true,
			Municipality: &BoundaryMatch{LayerFamily: "municipalities", FeatureID: "0820000", Name: "Denver"},
			County:       &BoundaryMatch{LayerFamily: "counties", FeatureID: "08031", Name: "Denver County"},
		},
		CodeFamily: "building",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Status != "local_record_required" {
		t.Fatalf("status = %q, want local_record_required", result.Status)
	}
	if len(result.Adoptions) != 0 {
		t.Fatalf("Colorado building resolution must not invent a statewide adoption: %#v", result.Adoptions)
	}
	assertCandidateKind(t, result, "municipality")
	assertContains(t, result.RequiredLocalRecords, "Current municipal building-code adoption ordinance and effective date")
}

func TestResolveColoradoElectricalAddsStateAdoption(t *testing.T) {
	catalog := loadTestCatalog(t)
	result, err := Resolve(catalog, ResolutionRequest{
		Context:    &GeographicContext{StateID: "US-CO", StateFIPS: "08", County: &BoundaryMatch{Name: "Larimer County"}},
		CodeFamily: "electrical",
	})
	if err != nil {
		t.Fatal(err)
	}
	if result.Status != "partially_resolved" {
		t.Fatalf("status = %q, want partially_resolved", result.Status)
	}
	if len(result.Adoptions) != 1 || result.Adoptions[0].CodeFamily != "electrical" {
		t.Fatalf("unexpected adoptions: %#v", result.Adoptions)
	}
	assertCandidateAuthority(t, result, "auth:us-co:state-electrical-board")
}

func TestResolveFloridaUsesStatewideBuildingCodeWithoutSuppressingLocalFollowUp(t *testing.T) {
	catalog := loadTestCatalog(t)
	result, err := Resolve(catalog, ResolutionRequest{
		Context: &GeographicContext{
			StateID:      "US-FL",
			StateFIPS:    "12",
			Incorporated: true,
			Municipality: &BoundaryMatch{Name: "Orlando"},
			County:       &BoundaryMatch{Name: "Orange County"},
		},
		CodeFamily: "building",
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(result.Adoptions) != 1 || result.Adoptions[0].StateCodeName != "Florida Building Code, Building, 8th Edition (2023)" {
		t.Fatalf("unexpected adoptions: %#v", result.Adoptions)
	}
	assertCandidateKind(t, result, "municipality")
	assertContains(t, result.RequiredLocalRecords, "Local technical amendments and administrative procedures")
}

func TestResolveNewJerseyOperationalFireUsesSeparateAuthority(t *testing.T) {
	catalog := loadTestCatalog(t)
	result, err := Resolve(catalog, ResolutionRequest{
		Context: &GeographicContext{
			StateID:           "US-NJ",
			StateFIPS:         "34",
			Incorporated:      true,
			Municipality:      &BoundaryMatch{Name: "Trenton"},
			FireJurisdictions: []BoundaryMatch{{Name: "Trenton Fire Prevention Bureau"}},
		},
		CodeFamily: "fire",
	})
	if err != nil {
		t.Fatal(err)
	}
	assertCandidateAuthority(t, result, "auth:us-nj:division-fire-safety")
	assertCandidateKind(t, result, "fire_authority")
	if len(result.Adoptions) != 1 || result.Adoptions[0].CodeFamily != "fire_operational" {
		t.Fatalf("unexpected fire adoptions: %#v", result.Adoptions)
	}
}

func TestResolveUnknownStateFailsClosed(t *testing.T) {
	catalog := loadTestCatalog(t)
	result, err := Resolve(catalog, ResolutionRequest{Context: &GeographicContext{StateID: "US-XX", StateFIPS: "99"}})
	if err != nil {
		t.Fatal(err)
	}
	if result.Status != "insufficient_evidence" {
		t.Fatalf("status = %q, want insufficient_evidence", result.Status)
	}
	if len(result.AuthorityCandidates) != 0 || len(result.Adoptions) != 0 {
		t.Fatalf("unknown state must not return unsupported conclusions: %#v", result)
	}
}

func TestProfilesRoundTripDeterministically(t *testing.T) {
	catalog := loadTestCatalog(t)
	for _, profile := range catalog.Profiles() {
		raw, err := json.Marshal(profile)
		if err != nil {
			t.Fatal(err)
		}
		var decoded StateProfile
		if err := json.Unmarshal(raw, &decoded); err != nil {
			t.Fatal(err)
		}
		if err := ValidateProfile(decoded); err != nil {
			t.Fatalf("round-trip profile %s: %v", profile.ProfileID, err)
		}
	}
}

func loadTestCatalog(t *testing.T) Catalog {
	t.Helper()
	catalog, err := LoadCatalog(filepath.Join("..", "..", "data", "regulatory"))
	if err == nil {
		return catalog
	}
	cwd, _ := os.Getwd()
	t.Fatalf("load catalog from %s: %v", cwd, err)
	return Catalog{}
}

func assertCandidateKind(t *testing.T, result ResolutionResult, kind string) {
	t.Helper()
	for _, candidate := range result.AuthorityCandidates {
		if candidate.Kind == kind {
			return
		}
	}
	t.Fatalf("missing candidate kind %q in %#v", kind, result.AuthorityCandidates)
}

func assertCandidateAuthority(t *testing.T, result ResolutionResult, id string) {
	t.Helper()
	for _, candidate := range result.AuthorityCandidates {
		if candidate.AuthorityID == id {
			return
		}
	}
	t.Fatalf("missing authority %q in %#v", id, result.AuthorityCandidates)
}

func assertContains(t *testing.T, values []string, wanted string) {
	t.Helper()
	for _, value := range values {
		if value == wanted {
			return
		}
	}
	t.Fatalf("missing %q in %#v", wanted, values)
}
