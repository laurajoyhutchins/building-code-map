package httpapi

import (
	"context"
	"net/http"
)

func (h *Handler) handleReadiness(w http.ResponseWriter) {
	readiness := h.engine.Readiness(context.Background())
	status := http.StatusOK
	if readiness.Status == "not_ready" {
		status = http.StatusServiceUnavailable
	}
	h.writeJSON(w, status, readiness)
}
