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

func newHandler() http.Handler {
	mux := http.NewServeMux()
	runtime := engine.NewRuntime(demoProvider{})
	mux.Handle(httpapi.ProjectVerificationPath, httpapi.NewHandler(engine.NewProjectVerifier(runtime), nil))
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
	return mux
}

func main() {
	addr := flag.String("addr", "127.0.0.1:8787", "HTTP address for synthetic-only reference application")
	flag.Parse()
	log.Printf("BCM synthetic reference application at http://%s (no real regulatory coverage)", *addr)
	log.Fatal(http.ListenAndServe(*addr, newHandler()))
}
