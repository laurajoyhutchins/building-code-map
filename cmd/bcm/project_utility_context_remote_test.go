package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/laurajoyhutchins/building-code-map/utilitycontext"
)

type remoteCatalogStub struct {
	value utilitycontext.Response
	err   error
}

func (r remoteCatalogStub) Lookup(_, _ string) (utilitycontext.Response, error) {
	return r.value, r.err
}

func (r remoteCatalogStub) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	})
}

func TestProjectUtilityContextRejectsUnavailableCatalogAfterGeographyReview(t *testing.T) {
	g := &geoFixture{}
	geo := mockGeography(t, g)
	defer geo.Close()
	private, err := newPrivateProvider(geo.URL, "provider-test-token")
	if err != nil {
		t.Fatal(err)
	}
	private.expectedBundle = "sha256:reviewed"
	handler := projectUtilityContextHandler(private, remoteCatalogStub{
		err: utilitycontext.ErrSourceUnavailable,
	})
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, httptest.NewRequest(http.MethodPost,
		ProjectUtilityContextPath, strings.NewReader(validGeoBody)))
	if rec.Code != http.StatusServiceUnavailable ||
		!strings.Contains(rec.Body.String(), "source_unavailable") {
		t.Fatalf("unavailable EIA service leaked or masqueraded as geography: %d %s",
			rec.Code, rec.Body.String())
	}
	if len(g.calls) != 4 {
		t.Fatalf("catalog consulted before geographic evidence: %v", g.calls)
	}
}

func TestProjectUtilityContextPreservesIndependentCatalogFailure(t *testing.T) {
	g := &geoFixture{stateFIPS: "06"}
	geo := mockGeography(t, g)
	defer geo.Close()
	private, err := newPrivateProvider(geo.URL, "provider-test-token")
	if err != nil {
		t.Fatal(err)
	}
	private.expectedBundle = "sha256:reviewed"
	handler := projectUtilityContextHandler(private, remoteCatalogStub{
		err: errors.New("should not read EIA candidates after incorrect county"),
	})
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, httptest.NewRequest(http.MethodPost,
		ProjectUtilityContextPath, strings.NewReader(validGeoBody)))
	if rec.Code != http.StatusUnprocessableEntity ||
		!strings.Contains(rec.Body.String(), "county_not_resolved") {
		t.Fatalf("upstream EIA status obscured independent county denial: %d %s",
			rec.Code, rec.Body.String())
	}
}
