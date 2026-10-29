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
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/laurajoyhutchins/building-code-map/engine"
	"github.com/laurajoyhutchins/building-code-map/httpapi"
	"github.com/laurajoyhutchins/building-code-map/utilitycontext"
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
	return handlerWithUtilityContext(provider, remote, allowedOrigin, nil)
}

func handlerWithUtilityContext(provider engine.Provider, remote *privateProvider, allowedOrigin string, catalog utilitycontext.Provider) http.Handler {
	return handlerWithSourcePolicy(provider, remote, allowedOrigin, catalog, engine.SourcePolicy{})
}

// The operator's reviewed policy is independent of the private Provider.
// Legacy endpoints without exact evidence remain conditional or unverified.
func handlerWithSourcePolicy(provider engine.Provider, remote *privateProvider, allowedOrigin string, catalog utilitycontext.Provider, policy engine.SourcePolicy) http.Handler {
	return handlerWithPolicies(provider, remote, allowedOrigin, catalog, policy, engine.SiteAuthorityPolicy{})
}

// Site evidence is admitted by the operator, independently of the remote provider.
// A missing policy cannot authorize an otherwise source-backed real-world project.
func handlerWithPolicies(provider engine.Provider, remote *privateProvider, allowedOrigin string, catalog utilitycontext.Provider, policy engine.SourcePolicy, sitePolicy engine.SiteAuthorityPolicy) http.Handler {
	mux := http.NewServeMux()
	if catalog != nil {
		mux.Handle(utilitycontext.Path, catalog.Handler())
		if remote != nil {
			mux.Handle(ProjectUtilityContextPath, projectUtilityContextHandler(remote, catalog))
		}
	}
	runtime := engine.NewRuntimeWithPolicies(provider, policy, sitePolicy)
	mux.Handle(httpapi.ProjectVerificationPath, httpapi.NewHandler(engine.NewProjectVerifier(runtime), nil))
	if remote != nil {
		mux.HandleFunc("/v1/readiness", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				w.Header().Set("Allow", "GET")
				http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
				return
			}
			// Preserve address readiness as the default for existing clients.
			// Coordinate-based site canaries must instead check the actual
			// coordinate capability, not a possibly absent address geocoder.
			capability := r.URL.Query().Get("capability")
			if capability == "" {
				capability = "address_lookup"
			}
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Cache-Control", "no-store")
			if capability != "address_lookup" && capability != "coordinate_resolution" {
				w.WriteHeader(http.StatusBadRequest)
				_ = json.NewEncoder(w).Encode(map[string]string{"status": "invalid_capability"})
				return
			}
			readiness, err := remote.readiness(r.Context())
			if err != nil || readiness.Capabilities[capability].Status != "available" ||
				readiness.Status == "not_ready" {
				w.WriteHeader(http.StatusServiceUnavailable)
				_ = json.NewEncoder(w).Encode(map[string]string{"status": "unavailable"})
				return
			}
			_ = json.NewEncoder(w).Encode(map[string]string{"status": "ready"})
		})
	}
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Cache-Control", "no-store")
		scope := "synthetic_demo_only"
		if remote != nil {
			scope = "private_provider_adapter"
		}
		_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok", "data_scope": scope})
	})
	static, err := fs.Sub(webFiles, "web")
	if err != nil {
		panic(err)
	}
	mux.Handle("/", http.FileServer(http.FS(static)))
	if allowedOrigin == "" {
		return mux
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Origin") == allowedOrigin {
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
			w.Header().Set("Vary", "Origin")
			if r.Method == http.MethodOptions && (r.URL.Path == httpapi.ProjectVerificationPath || r.URL.Path == ProjectUtilityContextPath) {
				w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
				w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
				w.WriteHeader(http.StatusNoContent)
				return
			}
		}
		mux.ServeHTTP(w, r)
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
		parsedOrigin, err := url.Parse(origin)
		if err != nil || parsedOrigin.Scheme != "https" || parsedOrigin.Host == "" ||
			parsedOrigin.Path != "" || parsedOrigin.RawQuery != "" || parsedOrigin.Fragment != "" {
			log.Fatal("BCM_PUBLIC_ORIGIN must be an HTTPS origin")
		}
		switch os.Getenv("BCM_PROVIDER_AUTH_MODE") {
		case "gcp_id_token":
			if os.Getenv("BCM_PROVIDER_TOKEN") != "" {
				log.Fatal("Do not set BCM_PROVIDER_TOKEN with Cloud Run workload identity")
			}
			remote.identityToken = cloudRunIdentityToken
		case "bearer":
			if os.Getenv("BCM_PROVIDER_TOKEN") == "" {
				log.Fatal("BCM_PROVIDER_TOKEN is required in bearer mode")
			}
		default:
			log.Fatal("BCM_PROVIDER_AUTH_MODE must be gcp_id_token or bearer for live service")
		}
		remote.expectedBundle = os.Getenv("BCM_PROVIDER_BUNDLE_DIGEST")
		if remote.expectedBundle == "" {
			log.Fatal("BCM_PROVIDER_BUNDLE_DIGEST must pin a reviewed, non-demo runtime bundle")
		}
		var catalog utilitycontext.Provider
		viewPath := os.Getenv("BCM_UTILITY_CONTEXT_VIEW_PATH")
		viewSHA := os.Getenv("BCM_UTILITY_CONTEXT_VIEW_SHA256")
		catalogOrigin := os.Getenv("BCM_UTILITY_CATALOG_API_URL")
		catalogViewSHA := os.Getenv("BCM_UTILITY_CATALOG_VIEW_SHA256")
		if catalogOrigin != "" || catalogViewSHA != "" {
			if catalogOrigin == "" || catalogViewSHA == "" || viewPath != "" || viewSHA != "" {
				log.Fatal("catalog API needs a pinned HTTPS origin and view SHA-256, not a parallel local EIA source")
			}
			catalog, err = utilitycontext.NewRemote(catalogOrigin, catalogViewSHA)
			if err != nil {
				log.Fatalf("catalog API source rejected: %v", err)
			}
		} else if viewPath != "" || viewSHA != "" {
			if viewPath == "" || viewSHA == "" {
				log.Fatal("BCM utility context requires both view path and exact SHA-256")
			}
			catalog, err = utilitycontext.Open(viewPath, viewSHA)
			if err != nil {
				log.Fatalf("BCM utility context source rejected: %v", err)
			}
		}
		policyPath := os.Getenv("BCM_SOURCE_POLICY_PATH")
		policySHA := os.Getenv("BCM_SOURCE_POLICY_SHA256")
		var policy engine.SourcePolicy
		if policyPath != "" || policySHA != "" {
			if policyPath == "" || policySHA == "" {
				log.Fatal("BCM source policy needs both private file path and exact SHA-256")
			}
			policy, err = loadPinnedSourcePolicy(policyPath, policySHA)
			if err != nil {
				log.Fatalf("BCM reviewed source policy rejected: %v", err)
			}
		}
		sitePath := os.Getenv("BCM_SITE_AUTHORITY_POLICY_PATH")
		siteSHA := os.Getenv("BCM_SITE_AUTHORITY_POLICY_SHA256")
		var sitePolicy engine.SiteAuthorityPolicy
		if sitePath != "" || siteSHA != "" {
			if sitePath == "" || siteSHA == "" {
				log.Fatal("BCM site authority policy needs both private file path and exact SHA-256")
			}
			sitePolicy, err = loadPinnedSiteAuthorityPolicy(sitePath, siteSHA)
			if err != nil {
				log.Fatal("BCM independently reviewed site authority policy rejected: ", err)
			}
		}
		handler = handlerWithPolicies(remote, remote, origin, catalog, policy, sitePolicy)
		log.Printf("BCM project verifier listening at %s; remote provider configured", *addr)
	} else {
		if os.Getenv("BCM_SITE_AUTHORITY_POLICY_PATH") != "" || os.Getenv("BCM_SITE_AUTHORITY_POLICY_SHA256") != "" {
			log.Fatal("BCM site authority policy requires the configured live backend, not the fictional demo")
		}
		if os.Getenv("BCM_UTILITY_CONTEXT_VIEW_PATH") != "" || os.Getenv("BCM_UTILITY_CONTEXT_VIEW_SHA256") != "" ||
			os.Getenv("BCM_UTILITY_CATALOG_API_URL") != "" || os.Getenv("BCM_UTILITY_CATALOG_VIEW_SHA256") != "" {
			log.Fatal("BCM utility context requires the configured live backend, not the fictional demo")
		}
		log.Printf("BCM fictional reference app at http://%s (no real-world coverage)", *addr)
	}
	server := &http.Server{Addr: *addr, Handler: handler, ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout: 15 * time.Second, WriteTimeout: 20 * time.Second, IdleTimeout: 60 * time.Second}
	log.Fatal(server.ListenAndServe())
}
