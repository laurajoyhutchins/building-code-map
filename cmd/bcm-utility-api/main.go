// Command bcm-utility-api serves source-pinned EIA county candidates only.
// It does not provide address geocoding, regulatory verdicts, or site utility verification.
package main

import (
 "encoding/json"
 "errors"
 "log"
 "net/http"
 "os"
 "regexp"
 "time"

 "github.com/laurajoyhutchins/building-code-map/utilitycontext"
)

const pagesOrigin = "https://laurajoyhutchins.github.io"
var portPattern = regexp.MustCompile(`^[0-9]{2,5}$`)

func handler(catalog *utilitycontext.Catalog) http.Handler {
 mux := http.NewServeMux()
 mux.Handle(utilitycontext.Path, catalog.Handler())
 mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodGet { w.Header().Set("Allow", "GET"); http.Error(w,"method not allowed",http.StatusMethodNotAllowed); return }
  w.Header().Set("Content-Type","application/json")
  w.Header().Set("Cache-Control","no-store")
  _ = json.NewEncoder(w).Encode(map[string]string{"status":"ok","capability":"county_utility_candidates_only"})
 })
 return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
  if r.Header.Get("Origin") == pagesOrigin {
   w.Header().Set("Access-Control-Allow-Origin",pagesOrigin)
   w.Header().Set("Vary","Origin")
  }
  if r.Method == http.MethodOptions && r.URL.Path == utilitycontext.Path {
   if r.Header.Get("Origin") != pagesOrigin { http.Error(w,"origin not allowed",http.StatusForbidden); return }
   w.Header().Set("Access-Control-Allow-Methods","GET, OPTIONS")
   w.Header().Set("Access-Control-Allow-Headers","Content-Type")
   w.WriteHeader(http.StatusNoContent)
   return
  }
  mux.ServeHTTP(w,r)
 })
}

func run() error {
 path := os.Getenv("BCM_UTILITY_CONTEXT_VIEW_PATH")
 digest := os.Getenv("BCM_UTILITY_CONTEXT_VIEW_SHA256")
 if path == "" || digest == "" {return errors.New("BCM utility view path and exact SHA-256 are required")}
 catalog, err := utilitycontext.Open(path,digest)
 if err != nil {return err}
 port := os.Getenv("PORT")
 if port == "" { port = "8080" }
 if !portPattern.MatchString(port) {return errors.New("invalid PORT")}
 srv:= &http.Server{Addr:":"+port, Handler:handler(catalog),
  ReadHeaderTimeout:5*time.Second,ReadTimeout:15*time.Second,
  WriteTimeout:20*time.Second,IdleTimeout:60*time.Second}
 log.Printf("BCM county candidate API on %s (no geocoding or regulatory authority)",srv.Addr)
 return srv.ListenAndServe()
}

func main() {if err:=run();err!=nil{log.Fatal(err)}}
