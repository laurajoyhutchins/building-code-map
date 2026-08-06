package httpapi

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"building-code-map/backend/engine"
	"building-code-map/backend/regulatory"
	"building-code-map/backend/snapshot"
)

func TestV1ResolveReturnsEngineResultAndRequiresApplicabilityDate(t *testing.T) {
	catalog, err := regulatory.LoadCatalog(filepath.Join("..", "..", "data", "regulatory"))
	if err != nil {
		t.Fatal(err)
	}
	authority, err := engine.New(engine.Config{
		Snapshot: snapshot.Snapshot{BoundaryFeatures: []snapshot.BoundaryFeature{{
			LayerFamily: "states", FeatureID: "08", Title: "Colorado", Attributes: map[string]any{"STATEFP": "08"}, Geometry: testPolygon(-109, 37, -102, 41),
		}}},
		RegulatoryCatalog: catalog,
	})
	if err != nil {
		t.Fatal(err)
	}
	handler := NewHandler(authority, Options{Snapshot: snapshot.Snapshot{BoundaryFeatures: []snapshot.BoundaryFeature{{LayerFamily: "states", FeatureID: "08"}}}})

	request := httptest.NewRequest(http.MethodPost, "/v1/resolve", bytes.NewBufferString(`{"point":{"longitude":-104.99,"latitude":39.74},"applicability_date":"2026-08-06","code_family":"building"}`))
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusOK {
		t.Fatalf("status=%d body=%s", recorder.Code, recorder.Body.String())
	}
	var result engine.Result
	if err := json.Unmarshal(recorder.Body.Bytes(), &result); err != nil {
		t.Fatal(err)
	}
	if result.SchemaVersion != engine.SchemaVersion || result.Provenance.BoundarySnapshotDigest != "" {
		t.Fatalf("result=%#v", result)
	}

	request = httptest.NewRequest(http.MethodPost, "/v1/resolve", bytes.NewBufferString(`{"point":{"longitude":-104.99,"latitude":39.74}}`))
	recorder = httptest.NewRecorder()
	handler.ServeHTTP(recorder, request)
	if recorder.Code != http.StatusBadRequest || !bytes.Contains(recorder.Body.Bytes(), []byte(`"code":"invalid_query"`)) {
		t.Fatalf("status=%d body=%s", recorder.Code, recorder.Body.String())
	}
}

func TestV1BundleAndReadinessUseEngineIdentity(t *testing.T) {
	identity := engine.BundleIdentity{SourceCommit: "0123456789abcdef0123456789abcdef01234567", BundleManifestDigest: "sha256:bundle"}
	authority, err := engine.New(engine.Config{BundleIdentity: identity})
	if err != nil {
		t.Fatal(err)
	}
	handler := NewHandler(authority, Options{})
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/v1/bundle", nil))
	if recorder.Code != http.StatusOK || !bytes.Contains(recorder.Body.Bytes(), []byte(`"bundle_manifest_digest":"sha256:bundle"`)) {
		t.Fatalf("bundle status=%d body=%s", recorder.Code, recorder.Body.String())
	}
}
