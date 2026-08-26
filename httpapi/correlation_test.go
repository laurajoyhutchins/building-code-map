package httpapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/laurajoyhutchins/building-code-map/engine"
)

type identityProvider struct{}

func (identityProvider) Resolve(_ context.Context, query engine.NormalizedQuery) (engine.Resolution, error) {
	return engine.Resolution{
		Query: query,
		Jurisdiction: engine.JurisdictionContext{
			CountryCode:             "US",
			PrimaryJurisdictionID:   "DEMO-XX",
			PrimaryJurisdictionName: "Demonstration Jurisdiction",
		},
		Codes: []engine.CodeResolution{{
			Family:  "building",
			Edition: "DEMO-2024",
			Status:  engine.StatusResolved,
			Basis:   "Synthetic correlation test",
		}},
		Provenance: engine.Provenance{
			EngineVersion: "demo-engine",
			SourceCommit:  "DEMO-COMMIT",
			BundleID:      "DEMO-BUNDLE",
		},
	}, nil
}

func correlatedTestHandler(sink CompletionSink) http.Handler {
	runtime := engine.NewRuntime(identityProvider{})
	return NewHandler(engine.NewProjectVerifier(runtime), sink)
}

func TestRequestCorrelationPreservesSafeIDAndRecordsCanonicalIdentity(t *testing.T) {
	var records []CompletionRecord
	handler := correlatedTestHandler(CompletionSinkFunc(func(record CompletionRecord) error {
		records = append(records, record)
		return nil
	}))
	request := httptest.NewRequest(http.MethodPost, ProjectVerificationPath, strings.NewReader(projectRequestBody()))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set(RequestIDHeader, "client.123-abc:7")
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body=%s", response.Code, http.StatusOK, response.Body.String())
	}
	if got := response.Header().Get(RequestIDHeader); got != "client.123-abc:7" {
		t.Fatalf("request id = %q", got)
	}
	if len(records) != 1 {
		t.Fatalf("records = %d, want 1", len(records))
	}
	record := records[0]
	if record.RequestID != "client.123-abc:7" || record.Method != http.MethodPost || record.RouteClass != ProjectVerificationRouteClass || record.Status != http.StatusOK {
		t.Fatalf("record = %#v", record)
	}
	if record.DurationMS < 0 {
		t.Fatalf("duration_ms = %d", record.DurationMS)
	}
	if record.RuntimeIdentity == nil || record.RuntimeIdentity.EngineVersion != "demo-engine" || record.RuntimeIdentity.SourceCommit != "DEMO-COMMIT" || record.RuntimeIdentity.BundleID != "DEMO-BUNDLE" {
		t.Fatalf("runtime identity = %#v", record.RuntimeIdentity)
	}
}

func TestRequestCorrelationReplacesUnsafeOrMissingIDs(t *testing.T) {
	tests := map[string]string{
		"missing":  "",
		"spaces":   "unsafe request id",
		"too long": strings.Repeat("a", maxRequestIDLength+1),
	}
	for name, supplied := range tests {
		t.Run(name, func(t *testing.T) {
			var record CompletionRecord
			handler := correlatedTestHandler(CompletionSinkFunc(func(value CompletionRecord) error {
				record = value
				return nil
			}))
			request := httptest.NewRequest(http.MethodGet, ProjectVerificationPath, nil)
			if supplied != "" {
				request.Header.Set(RequestIDHeader, supplied)
			}
			response := httptest.NewRecorder()

			handler.ServeHTTP(response, request)

			got := response.Header().Get(RequestIDHeader)
			if got == "" || got == supplied || !isSafeRequestID(got) {
				t.Fatalf("effective request id = %q for supplied %q", got, supplied)
			}
			if record.RequestID != got || record.Status != http.StatusMethodNotAllowed {
				t.Fatalf("record = %#v", record)
			}
		})
	}
}

func TestEveryHTTPResponseCarriesRequestID(t *testing.T) {
	handler := correlatedTestHandler(nil)
	for _, test := range []struct {
		method string
		path   string
		body   string
	}{
		{method: http.MethodGet, path: ProjectVerificationPath},
		{method: http.MethodPost, path: ProjectVerificationPath, body: `{"project_id":`},
		{method: http.MethodGet, path: "/not-a-route"},
	} {
		request := httptest.NewRequest(test.method, test.path, strings.NewReader(test.body))
		response := httptest.NewRecorder()
		handler.ServeHTTP(response, request)
		if got := response.Header().Get(RequestIDHeader); !isSafeRequestID(got) {
			t.Fatalf("%s %s request id = %q", test.method, test.path, got)
		}
	}
}

func TestCompletionRecordExcludesAddressAndRawBody(t *testing.T) {
	const privateAddress = "987 Distinctive Private Path Apt 123"
	requestBody := `{"project_id":"DEMO-PRIVATE","address":"` + privateAddress + `","applicability_date":"2026-08-25","facts":{"occupancy":"classified-demo-fact"}}`
	var record CompletionRecord
	handler := correlatedTestHandler(CompletionSinkFunc(func(value CompletionRecord) error {
		record = value
		return nil
	}))
	request := httptest.NewRequest(http.MethodPost, ProjectVerificationPath, strings.NewReader(requestBody))
	request.Header.Set("Content-Type", "application/json")
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	payload, err := json.Marshal(record)
	if err != nil {
		t.Fatalf("marshal record: %v", err)
	}
	serialized := string(payload)
	for _, forbidden := range []string{privateAddress, requestBody, "classified-demo-fact", "address", "facts"} {
		if strings.Contains(serialized, forbidden) {
			t.Fatalf("completion record leaked %q: %s", forbidden, serialized)
		}
	}
}

func TestCompletionSinkFailureAndPanicAreNonFatal(t *testing.T) {
	tests := map[string]CompletionSink{
		"error": CompletionSinkFunc(func(CompletionRecord) error { return errors.New("sink unavailable") }),
		"panic": CompletionSinkFunc(func(CompletionRecord) error { panic("sink exploded") }),
	}
	for name, sink := range tests {
		t.Run(name, func(t *testing.T) {
			handler := correlatedTestHandler(sink)
			request := httptest.NewRequest(http.MethodPost, ProjectVerificationPath, strings.NewReader(projectRequestBody()))
			request.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()

			handler.ServeHTTP(response, request)

			if response.Code != http.StatusOK {
				t.Fatalf("status = %d, want %d; body=%s", response.Code, http.StatusOK, response.Body.String())
			}
			if !isSafeRequestID(response.Header().Get(RequestIDHeader)) {
				t.Fatalf("request id = %q", response.Header().Get(RequestIDHeader))
			}
		})
	}
}

func TestCompletionRecordBoundsUnknownMethodAndRoute(t *testing.T) {
	var record CompletionRecord
	handler := correlatedTestHandler(CompletionSinkFunc(func(value CompletionRecord) error {
		record = value
		return nil
	}))
	request := httptest.NewRequest("HIGH-CARDINALITY-METHOD-123456", "/private/absolute-looking/path/123456", nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	if record.Method != OtherMethodClass || record.RouteClass != UnmatchedRouteClass || record.Status != http.StatusNotFound {
		t.Fatalf("record = %#v", record)
	}
}

func TestJSONLineSinkWritesOnlyCompletionRecord(t *testing.T) {
	var buffer bytes.Buffer
	sink := &JSONLineSink{Writer: &buffer}
	record := CompletionRecord{RequestID: "demo-1", Method: http.MethodPost, RouteClass: ProjectVerificationRouteClass, Status: http.StatusOK, DurationMS: 4}
	if err := sink.RecordCompletion(record); err != nil {
		t.Fatalf("RecordCompletion() error = %v", err)
	}
	if !strings.HasSuffix(buffer.String(), "\n") {
		t.Fatalf("JSON line = %q", buffer.String())
	}
	var decoded CompletionRecord
	if err := json.Unmarshal(bytes.TrimSpace(buffer.Bytes()), &decoded); err != nil {
		t.Fatalf("unmarshal JSON line: %v", err)
	}
	if decoded != record {
		t.Fatalf("decoded = %#v, want %#v", decoded, record)
	}
}
