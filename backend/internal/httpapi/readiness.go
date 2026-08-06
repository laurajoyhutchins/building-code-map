package httpapi

import "net/http"

type capabilityReadiness struct {
	Status   string `json:"status"`
	Required bool   `json:"required"`
	Message  string `json:"message"`
}

type snapshotReadiness struct {
	Status     string `json:"status"`
	SnapshotID string `json:"snapshot_id,omitempty"`
}

type readinessResponse struct {
	Status       string                         `json:"status"`
	Readiness    string                         `json:"readiness"`
	Capabilities map[string]capabilityReadiness `json:"capabilities"`
	Snapshots    map[string]snapshotReadiness   `json:"snapshots"`
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

	snapshots := map[string]snapshotReadiness{
		"boundary": snapshotStatus(boundaryAvailable, h.boundarySnapshotID),
		"geocoder": snapshotStatus(geocoderAvailable, h.geocoderSnapshotID),
	}

	if !boundaryAvailable {
		h.writeJSON(w, http.StatusServiceUnavailable, readinessResponse{
			Status:       "not_ready",
			Readiness:    "not_ready",
			Capabilities: capabilities,
			Snapshots:    snapshots,
		})
		return
	}
	if !geocoderAvailable || !regulatoryAvailable {
		h.writeJSON(w, http.StatusOK, readinessResponse{
			Status:       "ok",
			Readiness:    "degraded",
			Capabilities: capabilities,
			Snapshots:    snapshots,
		})
		return
	}
	h.writeJSON(w, http.StatusOK, readinessResponse{
		Status:       "ok",
		Readiness:    "ready",
		Capabilities: capabilities,
		Snapshots:    snapshots,
	})
}

func capabilityStatus(available, required bool, availableMessage, unavailableMessage string) capabilityReadiness {
	if available {
		return capabilityReadiness{Status: "available", Required: required, Message: availableMessage}
	}
	return capabilityReadiness{Status: "unavailable", Required: required, Message: unavailableMessage}
}

func snapshotStatus(available bool, snapshotID string) snapshotReadiness {
	if available && snapshotID != "" {
		return snapshotReadiness{Status: "verified", SnapshotID: snapshotID}
	}
	if available {
		return snapshotReadiness{Status: "unidentified"}
	}
	return snapshotReadiness{Status: "unavailable"}
}
