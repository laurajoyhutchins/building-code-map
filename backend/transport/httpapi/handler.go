package httpapi

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"building-code-map/backend/engine"
	"building-code-map/backend/geocoder"
	"building-code-map/backend/regulatory"
	"building-code-map/backend/snapshot"
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
	Snapshot           snapshot.Snapshot
	AllowedOrigins     []string
	RegulatoryCatalog  regulatory.Catalog
	Geocoder           geocoder.Service
	Clock              engine.Clock
	BundleIdentity     engine.BundleIdentity
	BoundarySnapshotID string
	GeocoderSnapshotID string
}

type Handler struct {
	snapshot           snapshot.Snapshot
	boundaryMapRecords []snapshot.BoundaryMapRecord
	layerIndex         map[string]snapshot.LayerFamily
	featureIndex       map[string]snapshot.BoundaryFeature
	allowedOrigins     map[string]struct{}
	regulatoryCatalog  regulatory.Catalog
	geocoder           geocoder.Service
	engine             engine.Engine
	engineClock        engine.Clock
	boundarySnapshotID string
	geocoderSnapshotID string
}

// NewHandler adapts one authority engine to the HTTP transport. The engine
// owns all location and regulatory decisions; options only provide transport
// configuration and cached records used by legacy browsing endpoints.
func NewHandler(authority engine.Engine, opt Options) http.Handler {
	if authority == nil {
		panic("httpapi: nil authority engine")
	}
	if opt.Clock == nil {
		opt.Clock = engine.RealClock{}
	}
	snap := opt.Snapshot
	allowedOrigins := opt.AllowedOrigins
	if len(allowedOrigins) == 0 {
		allowedOrigins = defaultAllowedOrigins
	}
	handler := &Handler{
		snapshot:           snap,
		boundaryMapRecords: snapshot.MapRecords(snap.BoundaryFeatures),
		layerIndex:         make(map[string]snapshot.LayerFamily, len(snap.LayerFamilies)),
		featureIndex:       make(map[string]snapshot.BoundaryFeature, len(snap.BoundaryFeatures)),
		allowedOrigins:     make(map[string]struct{}, len(allowedOrigins)),
		regulatoryCatalog:  opt.RegulatoryCatalog,
		geocoder:           opt.Geocoder,
		engine:             authority,
		engineClock:        opt.Clock,
		boundarySnapshotID: strings.TrimSpace(opt.BoundarySnapshotID),
		geocoderSnapshotID: strings.TrimSpace(opt.GeocoderSnapshotID),
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

// NewLegacyHandler keeps the current website-server construction usable while
// callers migrate to NewHandler and an explicitly constructed engine.
func NewLegacyHandler(snap snapshot.Snapshot, options ...Options) http.Handler {
	opt := Options{RegulatoryCatalog: regulatory.EmptyCatalog()}
	if len(options) > 0 {
		opt = options[0]
	}
	clock := opt.Clock
	if clock == nil {
		clock = engine.RealClock{}
	}
	identity := opt.BundleIdentity
	if identity.BoundarySnapshotDigest == "" {
		identity.BoundarySnapshotDigest = strings.TrimSpace(opt.BoundarySnapshotID)
	}
	if identity.GeocoderSnapshotDigest == "" {
		identity.GeocoderSnapshotDigest = strings.TrimSpace(opt.GeocoderSnapshotID)
	}
	authorityEngine, err := engine.New(engine.Config{
		Snapshot:          snap,
		RegulatoryCatalog: opt.RegulatoryCatalog,
		Geocoder:          opt.Geocoder,
		Clock:             clock,
		BundleIdentity:    identity,
	})
	if err != nil {
		panic(err)
	}

	allowedOrigins := opt.AllowedOrigins
	if len(allowedOrigins) == 0 {
		allowedOrigins = defaultAllowedOrigins
	}

	handler := &Handler{
		snapshot:           snap,
		boundaryMapRecords: snapshot.MapRecords(snap.BoundaryFeatures),
		layerIndex:         make(map[string]snapshot.LayerFamily, len(snap.LayerFamilies)),
		featureIndex:       make(map[string]snapshot.BoundaryFeature, len(snap.BoundaryFeatures)),
		allowedOrigins:     make(map[string]struct{}, len(allowedOrigins)),
		regulatoryCatalog:  opt.RegulatoryCatalog,
		geocoder:           opt.Geocoder,
		engine:             authorityEngine,
		engineClock:        clock,
		boundarySnapshotID: strings.TrimSpace(opt.BoundarySnapshotID),
		geocoderSnapshotID: strings.TrimSpace(opt.GeocoderSnapshotID),
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
	if h.handleCORS(w, r) {
		return
	}

	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/health":
		h.writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	case r.Method == http.MethodGet && r.URL.Path == "/ready":
		h.handleReadiness(w)
	case r.Method == http.MethodGet && r.URL.Path == "/layers":
		h.writeJSON(w, http.StatusOK, h.snapshot.LayerFamilies)
	case r.Method == http.MethodGet && r.URL.Path == "/boundaries":
		h.writeJSON(w, http.StatusOK, h.boundaryMapRecords)
	case r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/features/"):
		h.handleFeature(w, r)
	case r.Method == http.MethodPost && r.URL.Path == "/geocode":
		h.handleGeocode(w, r)
	case r.Method == http.MethodPost && r.URL.Path == "/lookup":
		h.handleLookup(w, r)
	case r.Method == http.MethodPost && r.URL.Path == "/resolve":
		h.handleResolve(w, r)
	case r.Method == http.MethodPost && r.URL.Path == "/v1/geocode":
		h.handleV1Geocode(w, r)
	case r.Method == http.MethodPost && r.URL.Path == "/v1/lookup":
		h.handleV1Lookup(w, r)
	case r.Method == http.MethodPost && r.URL.Path == "/v1/resolve":
		h.handleV1Resolve(w, r)
	case r.Method == http.MethodGet && r.URL.Path == "/v1/readiness":
		h.handleV1Readiness(w, r)
	case r.Method == http.MethodGet && r.URL.Path == "/v1/bundle":
		h.handleV1Bundle(w, r)
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

func (h *Handler) handleCORS(w http.ResponseWriter, r *http.Request) bool {
	origin := r.Header.Get("Origin")
	_, originAllowed := h.allowedOrigins[origin]
	if origin != "" && originAllowed {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Vary", "Origin, Access-Control-Request-Method, Access-Control-Request-Headers")
	}

	if r.Method != http.MethodOptions {
		return false
	}
	if origin == "" || !originAllowed {
		h.writeJSON(w, http.StatusForbidden, map[string]string{"error": "CORS origin is not allowed"})
		return true
	}
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Max-Age", "600")
	w.WriteHeader(http.StatusNoContent)
	return true
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
