package httpapi

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"building-code-map/backend/internal/regulatory"
	"building-code-map/backend/internal/snapshot"
)

var defaultAllowedOrigins = []string{
	"http://localhost:5173",
	"http://127.0.0.1:5173",
	"http://[::1]:5173",
}

func ParseAllowedOrigins(rawValue string) []string {
	rawValue = strings.TrimSpace(rawValue)
	if rawValue == "" {
		return append([]string(nil), defaultAllowedOrigins...)
	}

	origins := strings.Split(rawValue, ",")
	validOrigins := make([]string, 0, len(origins))
	for _, origin := range origins {
		candidate := strings.TrimRight(strings.TrimSpace(origin), "/")
		if candidate == "" {
			continue
		}

		if isValidOrigin(candidate) {
			validOrigins = append(validOrigins, candidate)
		}
	}

	if len(validOrigins) == 0 {
		return append([]string(nil), defaultAllowedOrigins...)
	}

	return validOrigins
}

type Options struct {
	AllowedOrigins    []string
	RegulatoryCatalog regulatory.Catalog
}

type Handler struct {
	snapshot          snapshot.Snapshot
	layerIndex        map[string]snapshot.LayerFamily
	featureIndex      map[string]snapshot.BoundaryFeature
	allowedOrigins    map[string]struct{}
	regulatoryCatalog regulatory.Catalog
}

func NewHandler(snap snapshot.Snapshot, options ...Options) http.Handler {
	opt := Options{RegulatoryCatalog: regulatory.EmptyCatalog()}
	if len(options) > 0 {
		opt = options[0]
	}

	allowedOrigins := opt.AllowedOrigins
	if len(allowedOrigins) == 0 {
		allowedOrigins = defaultAllowedOrigins
	}

	handler := &Handler{
		snapshot:          snap,
		layerIndex:        make(map[string]snapshot.LayerFamily, len(snap.LayerFamilies)),
		featureIndex:      make(map[string]snapshot.BoundaryFeature, len(snap.BoundaryFeatures)),
		allowedOrigins:    make(map[string]struct{}, len(allowedOrigins)),
		regulatoryCatalog: opt.RegulatoryCatalog,
	}
	for _, origin := range allowedOrigins {
		handler.allowedOrigins[origin] = struct{}{}
	}
	for _, layer := range snap.LayerFamilies {
		handler.layerIndex[layer.Key] = layer
	}
	for _, feature := range snap.BoundaryFeatures {
		handler.featureIndex[featureKey(feature.LayerFamily, feature.FeatureID)] = feature
	}

	return handler
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.applyCORS(w, r)

	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/health":
		h.writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	case r.Method == http.MethodGet && r.URL.Path == "/ready":
		h.writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	case r.Method == http.MethodGet && r.URL.Path == "/layers":
		h.writeJSON(w, http.StatusOK, h.snapshot.LayerFamilies)
	case r.Method == http.MethodGet && r.URL.Path == "/boundaries":
		h.writeJSON(w, http.StatusOK, h.snapshot.BoundaryFeatures)
	case r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/features/"):
		h.handleFeature(w, r)
	case r.Method == http.MethodPost && r.URL.Path == "/resolve":
		h.handleResolve(w, r)
	case r.Method == http.MethodGet && r.URL.Path == "/refresh/status":
		h.writeJSON(w, http.StatusOK, h.snapshot.RefreshStatus)
	case r.Method == http.MethodPost && r.URL.Path == "/refresh/trigger":
		h.writeJSON(w, http.StatusOK, map[string]string{
			"status":  "disabled",
			"message": "Live refresh is disabled for the cached snapshot.",
		})
	default:
		http.NotFound(w, r)
	}
}

func (h *Handler) handleFeature(w http.ResponseWriter, r *http.Request) {
	layerFamily, featureID, ok := strings.Cut(strings.TrimPrefix(r.URL.Path, "/features/"), "/")
	if !ok || layerFamily == "" || featureID == "" {
		http.NotFound(w, r)
		return
	}

	feature, found := h.featureIndex[featureKey(layerFamily, featureID)]
	if !found {
		http.NotFound(w, r)
		return
	}

	h.writeJSON(w, http.StatusOK, feature.Record())
}

func (h *Handler) applyCORS(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	if origin == "" {
		return
	}
	if _, ok := h.allowedOrigins[origin]; ok {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Vary", "Origin")
	}
}

func (h *Handler) writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func isValidOrigin(origin string) bool {
	parsed, err := url.Parse(origin)
	if err != nil {
		return false
	}

	if parsed.Scheme != "http" && parsed.Scheme != "https" {
		return false
	}
	if parsed.Host == "" {
		return false
	}
	if parsed.Path != "" || parsed.RawQuery != "" || parsed.Fragment != "" {
		return false
	}

	return true
}

func featureKey(layerFamily, featureID string) string {
	return layerFamily + "\x00" + featureID
}
