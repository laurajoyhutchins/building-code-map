package httpapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"building-code-map/backend/internal/geocoder"
	"building-code-map/backend/internal/regulatory"
)

type geocodeRequest struct {
	Address string `json:"address"`
}

type lookupRequest struct {
	Address           string `json:"address"`
	CodeFamily        string `json:"code_family,omitempty"`
	ProjectType       string `json:"project_type,omitempty"`
	ApplicabilityDate string `json:"applicability_date,omitempty"`
}

type lookupResponse struct {
	Geocode   geocoder.Result             `json:"geocode"`
	Resolution regulatory.ResolutionResult `json:"resolution"`
}

func (h *Handler) handleGeocode(w http.ResponseWriter, r *http.Request) {
	if h.geocoder == nil {
		h.writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "local geocoder is unavailable"})
		return
	}
	request, err := decodeStrictRequest[geocodeRequest](w, r)
	if err != nil {
		h.writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	result, err := h.geocoder.Geocode(r.Context(), geocoder.Query{Address: request.Address})
	if err != nil {
		h.writeGeocoderError(w, err)
		return
	}
	h.writeJSON(w, geocodeStatusCode(result.Status), result)
}

func (h *Handler) handleLookup(w http.ResponseWriter, r *http.Request) {
	if h.geocoder == nil {
		h.writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "local geocoder is unavailable"})
		return
	}
	if h.regulatoryCatalog.Len() == 0 {
		h.writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "regulatory policy catalog is unavailable"})
		return
	}
	request, err := decodeStrictRequest[lookupRequest](w, r)
	if err != nil {
		h.writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	geocodeResult, err := h.geocoder.Geocode(r.Context(), geocoder.Query{Address: request.Address})
	if err != nil {
		h.writeGeocoderError(w, err)
		return
	}
	if geocodeResult.Status != geocoder.StatusMatched || geocodeResult.Selected == nil {
		h.writeJSON(w, geocodeStatusCode(geocodeResult.Status), geocodeResult)
		return
	}

	resolution, err := h.resolveRequest(regulatory.ResolutionRequest{
		Point: &regulatory.Point{
			Longitude: geocodeResult.Selected.Longitude,
			Latitude:  geocodeResult.Selected.Latitude,
		},
		CodeFamily:        request.CodeFamily,
		ProjectType:       request.ProjectType,
		ApplicabilityDate: request.ApplicabilityDate,
	})
	if err != nil {
		h.writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
		return
	}
	h.writeJSON(w, http.StatusOK, lookupResponse{Geocode: geocodeResult, Resolution: resolution})
}

func (h *Handler) resolveRequest(request regulatory.ResolutionRequest) (regulatory.ResolutionResult, error) {
	if request.Context == nil {
		if request.Point == nil {
			return regulatory.ResolutionResult{}, errors.New("point or context is required")
		}
		context, err := resolveGeographicContext(h.snapshot, h.regulatoryCatalog, *request.Point)
		if err != nil {
			return regulatory.ResolutionResult{}, err
		}
		request.Context = &context
	}
	return regulatory.Resolve(h.regulatoryCatalog, request)
}

func decodeStrictRequest[T any](w http.ResponseWriter, r *http.Request) (T, error) {
	var value T
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&value); err != nil {
		return value, fmt.Errorf("invalid request: %w", err)
	}
	if err := ensureSingleJSONValue(decoder); err != nil {
		return value, err
	}
	return value, nil
}

func (h *Handler) writeGeocoderError(w http.ResponseWriter, err error) {
	if errors.Is(err, geocoder.ErrInvalidAddress) {
		h.writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	h.writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "local geocoder query failed"})
}

func geocodeStatusCode(status geocoder.Status) int {
	switch status {
	case geocoder.StatusMatched:
		return http.StatusOK
	case geocoder.StatusAmbiguous:
		return http.StatusConflict
	case geocoder.StatusNotFound:
		return http.StatusUnprocessableEntity
	default:
		return http.StatusInternalServerError
	}
}
