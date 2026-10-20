package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/laurajoyhutchins/building-code-map/engine"
	"github.com/laurajoyhutchins/building-code-map/httpapi"
)

func verify(t *testing.T, address string, facts string) *httptest.ResponseRecorder {
	t.Helper()
	body := `{"project_id":"DEMO-001","address":"` + address + `","applicability_date":"2026-08-25","facts":` + facts + `}`
	request := httptest.NewRequest(http.MethodPost, httpapi.ProjectVerificationPath, strings.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()
	newHandler().ServeHTTP(response, request)
	return response
}

func TestSyntheticDemoHasCanonicalContractAndConditionalEvidence(t *testing.T) {
	response := verify(t, demoAddress, "{}")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	var basis engine.ProjectCodeBasis
	if err := json.NewDecoder(response.Body).Decode(&basis); err != nil {
		t.Fatal(err)
	}
	if basis.SchemaVersion != engine.ProjectCodeBasisSchemaVersion || basis.JurisdictionID != "DEMO-XX" {
		t.Fatalf("noncanonical demo basis: %#v", basis)
	}
	if basis.Verdict != engine.ProjectConditional || len(basis.FollowUpQuestions) != 1 {
		t.Fatalf("demo must remain conditional with a project fact question: %#v", basis)
	}
	if len(basis.Unresolved) != 2 || basis.CodeSet[0].Status != engine.StatusLocalRecordRequired {
		t.Fatalf("demo must visibly require local evidence: %#v", basis)
	}
}

func TestSyntheticDemoRefusesRealWorldQueries(t *testing.T) {
	response := verify(t, "1600 N Broadway Denver CO", "{}")
	if response.Code != http.StatusUnprocessableEntity {
		t.Fatalf("real-world address status = %d, want 422", response.Code)
	}
	if !strings.Contains(response.Body.String(), string(engine.ErrorOutsideSupportedCoverage)) {
		t.Fatalf("unexpected real-world address response: %s", response.Body.String())
	}
}

func TestSyntheticDemoAnswersFactWithoutInventingLocalEvidence(t *testing.T) {
	response := verify(t, demoAddress, `{"occupancy":"business"}`)
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, body=%s", response.Code, response.Body.String())
	}
	var basis engine.ProjectCodeBasis
	if err := json.NewDecoder(response.Body).Decode(&basis); err != nil {
		t.Fatal(err)
	}
	if len(basis.FollowUpQuestions) != 0 || basis.Verdict != engine.ProjectConditional {
		t.Fatalf("answer must clear question but not substitute for local record: %#v", basis)
	}
}

func TestPublicInterfaceIsServedAlongsideCanonicalHTTP(t *testing.T) {
	for _, path := range []string{"/", "/app.js", "/styles.css"} {
		response := httptest.NewRecorder()
		newHandler().ServeHTTP(response, httptest.NewRequest(http.MethodGet, path, nil))
		if response.Code != http.StatusOK {
			t.Fatalf("%s = %d", path, response.Code)
		}
	}
}
