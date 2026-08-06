package httpapi

import "net/http"

type capabilityReadiness struct {
	Status   string `json:"status"`
	Required bool   `json:"required"`
	Message  string `json:"message"`
}

type readinessResponse struct {
	Status       string                         `json:"status"`
	Readiness    string                         `json:"readiness"`
	Capabilities map[string]capabilityReadiness `json:"capabilities"`
}

func (h *Handler) handleReadiness(w http.ResponseWriter) {
	boundaryAvailable := len(h.snapshot.LayerFamilies) > 0 && len(h.snapshot.BoundaryFeatures) > 0
	geocoderAvailable := h.geocoder != nil
	regulatoryAvailable := h.regulatoryCatalog.Len() > 0

	capabilities := map[string]capabilityReadiness{
		"boundary_resolution": capabilityStatus(
			boundaryAvailable,
			true,
			"A validated local boundary snapshot is loaded.",
			"No usable local boundary snapshot is loaded.",
		),
		"coordinate_resolution": capabilityStatus(
			boundaryAvailable && regulatoryAvailable,
			false,
			"Coordinate input can proceed through boundary and regulatory resolution.",
			"Coordinate resolution requires both boundary data and regulatory profiles.",
		),
		"address_geocoding": capabilityStatus(
			geocoderAvailable,
			false,
			"A local geocoder snapshot is loaded.",
			"No usable local geocoder snapshot is loaded.",
		),
		"regulatory_resolution": capabilityStatus(
			regulatoryAvailable,
			false,
			"One or more validated regulatory profiles are loaded.",
			"No validated regulatory profiles are loaded.",
		),
		"address_lookup": capabilityStatus(
			boundaryAvailable && geocoderAvailable && regulatoryAvailable,
			false,
			"Address lookup can preserve geocoding evidence through regulatory resolution.",
			"Address lookup requires boundary data, a local geocoder, and regulatory profiles.",
		),
	}

	if !boundaryAvailable {
		h.writeJSON(w, http.StatusServiceUnavailable, readinessResponse{
			Status:       "not_ready",
			Readiness:    "not_ready",
			Capabilities: capabilities,
		})
		return
	}
	if !geocoderAvailable || !regulatoryAvailable {
		h.writeJSON(w, http.StatusOK, readinessResponse{
			Status:       "ok",
			Readiness:    "degraded",
			Capabilities: capabilities,
		})
		return
	}
	h.writeJSON(w, http.StatusOK, readinessResponse{
		Status:       "ok",
		Readiness:    "ready",
		Capabilities: capabilities,
	})
}

func capabilityStatus(available, required bool, availableMessage, unavailableMessage string) capabilityReadiness {
	if available {
		return capabilityReadiness{Status: "available", Required: required, Message: availableMessage}
	}
	return capabilityReadiness{Status: "unavailable", Required: required, Message: unavailableMessage}
}
