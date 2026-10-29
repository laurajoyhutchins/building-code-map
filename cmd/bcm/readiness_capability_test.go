package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

// A real source provider can expose coordinate_resolution before it has
// address lookup. The single-point canary must not be blocked by geocoding.
func TestReadinessRespectsRequestedCoordinateCapability(t *testing.T) {
    upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if r.URL.Path != "/v1/readiness" {
            http.NotFound(w,r)
            return
        }
        w.Header().Set("Content-Type","application/json")
        _, _ = w.Write([]byte(`{"status":"ok","capabilities":{"address_lookup":{"status":"unavailable"},"coordinate_resolution":{"status":"available"}}}`))
    }))
    defer upstream.Close()
    provider, err := newPrivateProvider(upstream.URL,"")
    if err != nil {t.Fatal(err)}
    handler := handlerWithProvider(provider,provider,"")
    cases := []struct{ query string; status int }{
        {"?capability=coordinate_resolution",http.StatusOK},
        {"",http.StatusServiceUnavailable},
        {"?capability=address_lookup",http.StatusServiceUnavailable},
        {"?capability=unknown",http.StatusBadRequest},
    }
    for _, tc := range cases {
        response := httptest.NewRecorder()
        handler.ServeHTTP(response, httptest.NewRequest(http.MethodGet,"/v1/readiness"+tc.query,nil))
        if response.Code != tc.status {
            t.Errorf("query %q returned %d, expected %d: %s",tc.query,response.Code,tc.status,response.Body.String())
        }
    }
}
