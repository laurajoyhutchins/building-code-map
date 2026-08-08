package httpapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"building-code-map/backend/engine"
	"building-code-map/backend/internal/regulatory"
)

type pointResolutionRequest struct {
	Point             *regulatory.Point `json:"point"`
	CodeFamily        string            `json:"code_family,omitempty"`
	ProjectType       string            `json:"project_type,omitempty"`
	ApplicabilityDate string            `json:"applicability_date,omitempty"`
}

type boundaryAmbiguityResponse struct {
	Error        string                     `json:"error"`
	Code         string                     `json:"code"`
	LayerFamily  string                     `json:"layer_family"`
	Observations []regulatory.BoundaryMatch `json:"observations"`
}

func (h *Handler) handleResolve(w http.ResponseWriter, r *http.Request) {
	if h.regulatoryCatalog.Len() == 0 {
		h.writeJSON(w, http.StatusServiceUnavailable, map[string]string{"error": "regulatory policy catalog is unavailable"})
		return
	}

	request, err := decodeStrictRequest[pointResolutionRequest](w, r)
	if err != nil {
		h.writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	if request.Point == nil {
		h.writeJSON(w, http.StatusBadRequest, map[string]string{"error": "point is required"})
		return
	}

	dateWasAssumed := request.ApplicabilityDate == ""
	if dateWasAssumed {
		request.ApplicabilityDate = h.engineClock.Now().UTC().Format("2006-01-02")
	}
	result, err := h.engine.Resolve(r.Context(), engine.Query{
		Point:             &engine.Point{Longitude: request.Point.Longitude, Latitude: request.Point.Latitude},
		CodeFamily:        request.CodeFamily,
		ProjectType:       request.ProjectType,
		ApplicabilityDate: request.ApplicabilityDate,
	})
	if err != nil {
		h.writeResolutionError(w, err)
		return
	}
	if dateWasAssumed {
		result.Resolution.Warnings = append(result.Resolution.Warnings, fmt.Sprintf(
			"Applicability date was omitted; the server used %s in UTC. Confirm the date that governs the project.",
			result.Resolution.ApplicabilityDate,
		))
	}
	h.writeJSON(w, http.StatusOK, result.Resolution)
}

func (h *Handler) writeResolutionError(w http.ResponseWriter, err error) {
	var ambiguity *engine.BoundaryAmbiguityError
	if errors.As(err, &ambiguity) {
		h.writeJSON(w, http.StatusConflict, boundaryAmbiguityResponse{
			Error:        ambiguity.Error(),
			Code:         "boundary_ambiguous",
			LayerFamily:  ambiguity.LayerFamily,
			Observations: ambiguity.Matches,
		})
		return
	}
	var engineErr engine.EngineError
	if errors.As(err, &engineErr) && engineErr.Code == engine.ErrorBoundaryAmbiguous {
		response := boundaryAmbiguityResponse{Error: engineErr.Message, Code: string(engineErr.Code)}
		if engineErr.Details != nil {
			if layerFamily, ok := engineErr.Details["layer_family"].(string); ok {
				response.LayerFamily = layerFamily
			}
			if observations, ok := engineErr.Details["observations"].([]regulatory.BoundaryMatch); ok {
				response.Observations = observations
			}
		}
		h.writeJSON(w, http.StatusConflict, response)
		return
	}
	h.writeJSON(w, http.StatusUnprocessableEntity, map[string]string{"error": err.Error()})
}

func ensureSingleJSONValue(decoder *json.Decoder) error {
	var extra any
	if err := decoder.Decode(&extra); err == io.EOF {
		return nil
	} else if err != nil {
		return fmt.Errorf("invalid trailing JSON: %w", err)
	}
	return errors.New("request body must contain one JSON value")
}
