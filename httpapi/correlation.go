package httpapi

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/laurajoyhutchins/building-code-map/engine"
)

const (
	RequestIDHeader         = "X-Request-ID"
	UnmatchedRouteClass     = "unmatched"
	OtherMethodClass        = "OTHER"
	maxRequestIDLength      = 128
	generatedRequestIDBytes = 16
)

type RuntimeIdentity struct {
	EngineVersion string `json:"engine_version,omitempty"`
	SourceCommit  string `json:"source_commit,omitempty"`
	BundleID      string `json:"bundle_id,omitempty"`
}

type CompletionRecord struct {
	RequestID       string           `json:"request_id"`
	Method          string           `json:"method"`
	RouteClass      string           `json:"route_class"`
	Status          int              `json:"status"`
	DurationMS      int64            `json:"duration_ms"`
	RuntimeIdentity *RuntimeIdentity `json:"runtime_identity,omitempty"`
}

type CompletionSink interface {
	RecordCompletion(record CompletionRecord) error
}

type CompletionSinkFunc func(record CompletionRecord) error

func (sink CompletionSinkFunc) RecordCompletion(record CompletionRecord) error {
	return sink(record)
}

type JSONLineSink struct {
	Writer io.Writer
	mu     sync.Mutex
}

func (sink *JSONLineSink) RecordCompletion(record CompletionRecord) error {
	if sink == nil || sink.Writer == nil {
		return errors.New("completion sink writer is not configured")
	}
	payload, err := json.Marshal(record)
	if err != nil {
		return err
	}
	payload = append(payload, '\n')
	sink.mu.Lock()
	defer sink.mu.Unlock()
	_, err = sink.Writer.Write(payload)
	return err
}

type correlationHandler struct {
	next http.Handler
	sink CompletionSink
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

type runtimeIdentityCollector struct {
	identity *RuntimeIdentity
}

type runtimeIdentityContextKey struct{}

var fallbackRequestIDSequence atomic.Uint64

func withRequestCorrelation(next http.Handler, sink CompletionSink) http.Handler {
	return correlationHandler{next: next, sink: sink}
}

func (handler correlationHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	started := time.Now()
	requestID := effectiveRequestID(request.Header.Get(RequestIDHeader))
	response.Header().Set(RequestIDHeader, requestID)

	collector := &runtimeIdentityCollector{}
	request = request.WithContext(context.WithValue(request.Context(), runtimeIdentityContextKey{}, collector))
	recorder := &statusRecorder{ResponseWriter: response}
	handler.next.ServeHTTP(recorder, request)

	status := recorder.status
	if status == 0 {
		status = http.StatusOK
	}
	record := CompletionRecord{
		RequestID:       requestID,
		Method:          boundedMethod(request.Method),
		RouteClass:      routeClass(request.URL.Path),
		Status:          status,
		DurationMS:      time.Since(started).Milliseconds(),
		RuntimeIdentity: collector.identity,
	}
	recordCompletionSafely(handler.sink, record)
}

func (recorder *statusRecorder) WriteHeader(status int) {
	if recorder.status != 0 {
		return
	}
	recorder.status = status
	recorder.ResponseWriter.WriteHeader(status)
}

func (recorder *statusRecorder) Write(payload []byte) (int, error) {
	if recorder.status == 0 {
		recorder.WriteHeader(http.StatusOK)
	}
	return recorder.ResponseWriter.Write(payload)
}

func effectiveRequestID(candidate string) string {
	if isSafeRequestID(candidate) {
		return candidate
	}
	var random [generatedRequestIDBytes]byte
	if _, err := rand.Read(random[:]); err == nil {
		return "req_" + hex.EncodeToString(random[:])
	}
	sequence := fallbackRequestIDSequence.Add(1)
	seed := time.Now().UTC().Format(time.RFC3339Nano) + ":" + strconv.FormatUint(sequence, 10)
	digest := sha256.Sum256([]byte(seed))
	return "req_" + hex.EncodeToString(digest[:generatedRequestIDBytes])
}

func isSafeRequestID(value string) bool {
	if len(value) < 1 || len(value) > maxRequestIDLength || !isAlphaNumeric(value[0]) {
		return false
	}
	for index := 1; index < len(value); index++ {
		character := value[index]
		if isAlphaNumeric(character) || character == '.' || character == '_' || character == ':' || character == '-' {
			continue
		}
		return false
	}
	return true
}

func isAlphaNumeric(value byte) bool {
	return value >= 'a' && value <= 'z' || value >= 'A' && value <= 'Z' || value >= '0' && value <= '9'
}

func boundedMethod(method string) string {
	switch method {
	case http.MethodConnect, http.MethodDelete, http.MethodGet, http.MethodHead, http.MethodOptions,
		http.MethodPatch, http.MethodPost, http.MethodPut, http.MethodTrace:
		return method
	default:
		return OtherMethodClass
	}
}

func routeClass(path string) string {
	if path == ProjectVerificationPath {
		return ProjectVerificationRouteClass
	}
	return UnmatchedRouteClass
}

func observeRuntimeIdentity(ctx context.Context, provenance engine.Provenance) {
	collector, _ := ctx.Value(runtimeIdentityContextKey{}).(*runtimeIdentityCollector)
	if collector == nil {
		return
	}
	identity := &RuntimeIdentity{
		EngineVersion: strings.TrimSpace(provenance.EngineVersion),
		SourceCommit:  strings.TrimSpace(provenance.SourceCommit),
		BundleID:      strings.TrimSpace(provenance.BundleID),
	}
	if identity.EngineVersion == "" && identity.SourceCommit == "" && identity.BundleID == "" {
		return
	}
	collector.identity = identity
}

func recordCompletionSafely(sink CompletionSink, record CompletionRecord) {
	if sink == nil {
		return
	}
	defer func() {
		_ = recover()
	}()
	_ = sink.RecordCompletion(record)
}
