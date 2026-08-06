package httpapi

import (
	"errors"
	"net/http"

	"building-code-map/backend/engine"
)

type v1GeocodeRequest struct {
	Address string `json:"address"`
}

func (h *Handler) handleV1Geocode(w http.ResponseWriter, r *http.Request) {
	request, err := decodeStrictRequest[v1GeocodeRequest](w, r)
	if err != nil {
		h.writeV1Error(w, engine.EngineError{Code: engine.ErrorInvalidQuery, Message: err.Error()})
		return
	}
	result, err := h.engine.Geocode(r.Context(), request.Address)
	if err != nil {
		h.writeV1Error(w, err)
		return
	}
	h.writeJSON(w, http.StatusOK, result)
}

func (h *Handler) handleV1Lookup(w http.ResponseWriter, r *http.Request) {
	request, err := decodeStrictRequest[engine.Query](w, r)
	if err != nil {
		h.writeV1Error(w, engine.EngineError{Code: engine.ErrorInvalidQuery, Message: err.Error()})
		return
	}
	if request.Address != "" {
		result, err := h.engine.Resolve(r.Context(), request)
		if err != nil {
			h.writeV1Error(w, err)
			return
		}
		h.writeJSON(w, http.StatusOK, result)
		return
	}
	if request.Point == nil {
		h.writeV1Error(w, engine.EngineError{Code: engine.ErrorInvalidQuery, Message: "point or address is required"})
		return
	}
	result, err := h.engine.Lookup(r.Context(), *request.Point)
	if err != nil {
		h.writeV1Error(w, err)
		return
	}
	h.writeJSON(w, http.StatusOK, result)
}

func (h *Handler) handleV1Resolve(w http.ResponseWriter, r *http.Request) {
	request, err := decodeStrictRequest[engine.Query](w, r)
	if err != nil {
		h.writeV1Error(w, engine.EngineError{Code: engine.ErrorInvalidQuery, Message: err.Error()})
		return
	}
	result, err := h.engine.Resolve(r.Context(), request)
	if err != nil {
		h.writeV1Error(w, err)
		return
	}
	h.writeJSON(w, http.StatusOK, result)
}

func (h *Handler) handleV1Readiness(w http.ResponseWriter, r *http.Request) {
	readiness := h.engine.Readiness(r.Context())
	status := http.StatusOK
	if readiness.Status == "not_ready" {
		status = http.StatusServiceUnavailable
	}
	h.writeJSON(w, status, readiness)
}

func (h *Handler) handleV1Bundle(w http.ResponseWriter, r *http.Request) {
	h.writeJSON(w, http.StatusOK, h.engine.BundleIdentity(r.Context()))
}

func (h *Handler) writeV1Error(w http.ResponseWriter, err error) {
	var engineErr engine.EngineError
	if !errors.As(err, &engineErr) {
		engineErr = engine.EngineError{Code: engine.ErrorInternal, Message: err.Error(), Retryable: false}
	}
	status := http.StatusUnprocessableEntity
	switch engineErr.Code {
	case engine.ErrorInvalidQuery, engine.ErrorInvalidCoordinates:
		status = http.StatusBadRequest
	case engine.ErrorAddressAmbiguous, engine.ErrorBoundaryAmbiguous:
		status = http.StatusConflict
	case engine.ErrorAddressNotFound, engine.ErrorOutsideSupportedCoverage, engine.ErrorRegulatoryProfileMissing:
		status = http.StatusUnprocessableEntity
	case engine.ErrorRegulatoryCatalogUnavailable, engine.ErrorDataBundleInvalid:
		status = http.StatusServiceUnavailable
	case engine.ErrorInternal:
		status = http.StatusInternalServerError
	}
	h.writeJSON(w, status, engineErr)
}
