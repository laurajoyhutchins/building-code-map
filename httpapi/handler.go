package httpapi

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"mime"
	"net/http"
	"strings"

	"github.com/laurajoyhutchins/building-code-map/engine"
)

const (
	ProjectVerificationPath       = "/v1/project-code-basis"
	ProjectVerificationRouteClass = "project_code_verification"
	MaxRequestBodyBytes     int64 = 64 * 1024
)

type ProjectVerifier interface {
	VerifyProject(ctx context.Context, request engine.ProjectRequest) (engine.ProjectCodeBasis, error)
}

type Handler struct {
	verifier ProjectVerifier
}

type ErrorEnvelope struct {
	Error ErrorBody `json:"error"`
}

type ErrorBody struct {
	Code      string         `json:"code"`
	Message   string         `json:"message"`
	Details   map[string]any `json:"details,omitempty"`
	Retryable bool           `json:"retryable"`
}

type requestDecodeError struct {
	status  int
	code    string
	message string
}

func NewHandler(verifier ProjectVerifier) http.Handler {
	return Handler{verifier: verifier}
}

func (handler Handler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != ProjectVerificationPath {
		writeError(response, http.StatusNotFound, ErrorBody{Code: "not_found", Message: "route not found"})
		return
	}
	if request.Method != http.MethodPost {
		response.Header().Set("Allow", http.MethodPost)
		writeError(response, http.StatusMethodNotAllowed, ErrorBody{Code: "method_not_allowed", Message: "method must be POST"})
		return
	}
	if handler.verifier == nil {
		writeError(response, http.StatusInternalServerError, ErrorBody{Code: string(engine.ErrorInternal), Message: "project verifier is not configured"})
		return
	}
	if err := requireJSONContentType(request); err != nil {
		writeError(response, err.status, ErrorBody{Code: err.code, Message: err.message})
		return
	}

	projectRequest, decodeErr := decodeProjectRequest(response, request)
	if decodeErr != nil {
		writeError(response, decodeErr.status, ErrorBody{Code: decodeErr.code, Message: decodeErr.message})
		return
	}

	basis, err := handler.verifier.VerifyProject(request.Context(), projectRequest)
	if err != nil {
		engineErr := engine.NormalizeError(err)
		writeError(response, statusForEngineError(engineErr), ErrorBody{
			Code:      string(engineErr.Code),
			Message:   engineErr.Message,
			Details:   engineErr.Details,
			Retryable: engineErr.Retryable,
		})
		return
	}
	writeJSON(response, http.StatusOK, basis)
}

func requireJSONContentType(request *http.Request) *requestDecodeError {
	value := strings.TrimSpace(request.Header.Get("Content-Type"))
	if value == "" {
		return nil
	}
	mediaType, _, err := mime.ParseMediaType(value)
	if err != nil || mediaType != "application/json" {
		return &requestDecodeError{
			status:  http.StatusUnsupportedMediaType,
			code:    "unsupported_media_type",
			message: "Content-Type must be application/json",
		}
	}
	return nil
}

func decodeProjectRequest(response http.ResponseWriter, request *http.Request) (engine.ProjectRequest, *requestDecodeError) {
	request.Body = http.MaxBytesReader(response, request.Body, MaxRequestBodyBytes)
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()

	var projectRequest engine.ProjectRequest
	if err := decoder.Decode(&projectRequest); err != nil {
		return engine.ProjectRequest{}, classifyDecodeError(err)
	}

	var trailing any
	if err := decoder.Decode(&trailing); !errors.Is(err, io.EOF) {
		if err != nil {
			return engine.ProjectRequest{}, classifyDecodeError(err)
		}
		return engine.ProjectRequest{}, &requestDecodeError{
			status:  http.StatusBadRequest,
			code:    "invalid_json",
			message: "request body must contain exactly one JSON object",
		}
	}
	return projectRequest, nil
}

func classifyDecodeError(err error) *requestDecodeError {
	var maxBytesErr *http.MaxBytesError
	if errors.As(err, &maxBytesErr) {
		return &requestDecodeError{
			status:  http.StatusRequestEntityTooLarge,
			code:    "request_too_large",
			message: "request body exceeds the maximum size",
		}
	}
	return &requestDecodeError{
		status:  http.StatusBadRequest,
		code:    "invalid_json",
		message: "request body must be one valid project request JSON object",
	}
}

func statusForEngineError(err engine.EngineError) int {
	switch err.Code {
	case engine.ErrorInvalidQuery, engine.ErrorInvalidCoordinates:
		return http.StatusBadRequest
	case engine.ErrorAddressNotFound, engine.ErrorJurisdictionProfileMissing:
		return http.StatusNotFound
	case engine.ErrorAddressAmbiguous, engine.ErrorBoundaryAmbiguous, engine.ErrorOutsideSupportedCoverage:
		return http.StatusUnprocessableEntity
	case engine.ErrorGeocoderUnavailable, engine.ErrorRegulatoryCatalogUnavailable:
		return http.StatusServiceUnavailable
	case engine.ErrorDataBundleInvalid, engine.ErrorInternal:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

func writeError(response http.ResponseWriter, status int, body ErrorBody) {
	writeJSON(response, status, ErrorEnvelope{Error: body})
}

func writeJSON(response http.ResponseWriter, status int, value any) {
	response.Header().Set("Content-Type", "application/json; charset=utf-8")
	response.WriteHeader(status)
	_ = json.NewEncoder(response).Encode(value)
}
