// Command bcm serves the public, synthetic Building Code Map reference application.
// It does not contain or infer production regulatory data.
package main

import (
	"context"
	"embed"
	"encoding/json"
	"flag"
	"io/fs"
	"log"
	"os"
	"net/http"
	"strings"

	"github.com/laurajoyhutchins/building-code-map/engine"
	"github.com/laurajoyhutchins/building-code-map/httpapi"
)

//go:embed web/*
var webFiles embed.FS

const demoAddress = "100 Demo Plaza, DEMO-XX"

type demoProvider struct{}

func (demoProvider) Resolve(_ context.Context, query engine.NormalizedQuery) (engine.Resolution, error) {
	if !strings.EqualFold(strings.TrimSpace(query.Address), demoAddress) || query.Point != nil {
		return engine.Resolution{}, engine.EngineError{
			Code:    engine.ErrorOutsideSupportedCoverage,
			Message: "This public reference server supports only the fictional DEMO-XX address; it cannot verify real locations.",
		}
	}

	result := engine.Resolution{
		Jurisdiction: engine.JurisdictionContext{
			CountryCode:             "DEMO",
			PrimaryJurisdictionID:   "DEMO-XX",
			PrimaryJurisdictionName: "Demonstration Jurisdiction (fictional)",
		},
		Codes: []engine.CodeResolution{{
			Family:     "building",
			Edition:    "DEMO-2024",
			Status:     engine.StatusLocalRecordRequired,
			Basis:      "Fictional code instrument for interface demonstration only.",
			Unresolved: []string{"Obtain the fictional DEMO-XX local adoption record before treating this as verified."},
		}},
		Provenance: engine.Provenance{
			EngineVersion: "public-demo",
			SourceCommit:  "DEMO-SYNTHETIC",
			BundleID:      "DEMO-XX-ONLY",
		},
	}
	if query.ProjectFacts["occupancy"] == "" {
		result.Requirements = []engine.ResolutionRequirement{{
			ID:      "DEMO-OCCUPANCY",
			Kind:    engine.RequirementProjectFact,
			FactKey: "occupancy",
			Prompt:  "What is the project's occupancy?",
		}}
	}
	return result, nil
}

func newHandler() http.Handler { return handlerWithProvider(demoProvider{}, nil, "") }

func handlerWithProvider(provider engine.Provider, remote *privateProvider, allowedOrigin string) http.Handler {
	mux := http.NewServeMux()
	runtime := engine.NewRuntime(provider)
	mux.Handle(httpapi.ProjectVerificationPath, httpapi.NewHandler(engine.NewProjectVerifier(runtime), nil))
	if remote != nil {
		mux.HandleFunc("/v1/readiness", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				w.Header().Set("Allow", "GET")
				http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
				return
			}
			readiness, err := remote.readiness(r.Context())
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Cache-Control", "no-store")
			if err != nil || readiness.Capabilities["address_lookup"].Status != "available" ||
				readiness.Status == "not_ready" {
				w.WriteHeader(http.StatusServiceUnavailable)
				_ = json.NewEncoder(w).Encode(map[string]string{"status":"unavailable"})
				return
			}
			_ = json.NewEncoder(w).Encode(map[string]string{"status":"ready"})
		})
	}
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Cache-Control", "no-store")
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok", "data_scope": "synthetic_demo_only"})
	})
	static, err := fs.Sub(webFiles, "web")
	if err != nil {
		panic(err)
	}
	mux.Handle("/", http.FileServer(http.FS(static)))
	if allowedOrigin == "" { return mux }
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Origin") == allowedOrigin {
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
			w.Header().Set("Vary", "Origin")
			if r.Method == http.MethodOptions && r.URL.Path == httpapi.ProjectVerificationPath {
				w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
		mux.ServeHTTP(w,r)
	})
}

func main() {
	endpoint := os.Getenv("BCM_PROVIDER_URL")
	defaultAddress := "127.0.0.1:8787"
	if endpoint != "" {
		defaultAddress = ":8080"
		if port := os.Getenv("PORT"); port != "" {
			defaultAddress = ":" + port
		}
	}
	addr := flag.String("addr", defaultAddress, "HTTP listen address")
	flag.Parse()
	handler := newHandler()
	if endpoint != "" {
		remote, err := newPrivateProvider(endpoint, os.Getenv("BCM_PROVIDER_TOKEN"))
		if err != nil {
			log.Fatal(err)
		}
		origin := os.Getenv("BCM_PUBLIC_ORIGIN")
		if origin == "" { log.Fatal("BCM_PUBLIC_ORIGIN is required for the live provider") }
		handler = handlerWithProvider(remote, remote, origin)
		log.Printf("BCM project verifier listening at %s; remote provider configured", *addr)
	} else {
		log.Printf("BCM fictional reference app at http://%s (no real-world coverage)", *addr)
	}
	log.Fatal(http.ListenAndServe(*addr, handler))
}
