package bcastconsumer

import (
	"context"
	"os"
	"strings"
	"testing"
)

// Retained byte-for-byte from public BCAST at commit
// 8273a20f82fa147d8f5718ca3a7a7aacb8ac4160, Git blob
// 0bfc2f2be3fc32de6a7f196d0c4171da5379808e.
// This tests a real independently maintained 0.2 contract vector offline.
func unicodeFixture(t *testing.T) map[string]any {
	t.Helper()
	data, err := os.ReadFile("testdata/package-0.2.0-unicode.json")
	if err != nil {
		t.Fatal(err)
	}
	decoded, err := decode(data)
	if err != nil {
		t.Fatal(err)
	}
	return decoded.(map[string]any)
}

// Alter a canonical coordinate and recompute ALL dependent identities.
// A failing test must prove normalization admission, not an unrelated stale hash.
func rehashUnicodeCoordinates(t *testing.T, pkg map[string]any) string {
	t.Helper()
	pub := pkg["publication"].(map[string]any)
	fields := map[string]any{
		"namespace": pub["namespace"],
		"family":    pub["family"],
		"edition":   pub["edition"],
	}
	if revision, ok := pub["revision"]; ok {
		fields["revision"] = revision
	}
	publicationID, err := id("bcastpub", fields)
	if err != nil {
		t.Fatal(err)
	}
	pub["publication_id"] = publicationID
	renamed := map[string]string{}
	for _, element := range pkg["objects"].([]any) {
		obj := element.(map[string]any)
		oldID := asString(obj, "object_id")
		newID, err := id("bcastobj", map[string]any{
			"publication_id": publicationID,
			"kind":           obj["kind"],
			"locator":        obj["locator"],
		})
		if err != nil {
			t.Fatal(err)
		}
		obj["object_id"] = newID
		renamed[oldID] = newID
	}
	for _, element := range pkg["objects"].([]any) {
		obj := element.(map[string]any)
		if parent, present := obj["parent_id"]; present {
			obj["parent_id"] = renamed[parent.(string)]
		}
	}
	return rehashPublicBCASTPackage(t, pkg)
}

func TestNFCUnicodeCanonicalFixtureRoundTrip(t *testing.T) {
	pkg := unicodeFixture(t)
	requested := asString(pkg, "package_id")
	accepted, err := validate(pkg, requested)
	if err != nil {
		t.Fatalf("valid public BCAST NFC Unicode fixture rejected: %v", err)
	}
	pub := pkg["publication"].(map[string]any)
	if pub["edition"] != "Édition 2026" || pub["revision"] != "Révision 1" {
		t.Fatal("wrong exact upstream Unicode fixture")
	}
	if accepted.PublicationID != asString(pub, "publication_id") {
		t.Fatal("publication identity differs from public BCAST fixture")
	}

	s := server(t, pkg)
	defer s.Close()
	client, err := New(s.URL)
	if err != nil {
		t.Fatal(err)
	}
	remote, err := client.Fetch(context.Background(), requested)
	if err != nil {
		t.Fatalf("Unicode BCAST HTTP package roundtrip: %v", err)
	}
	if err := client.VerifyNavigation(context.Background(), remote); err != nil {
		t.Fatalf("Unicode BCAST HTTP navigation: %v", err)
	}
}

func TestNonNFCUnicodeIdentityFailsAfterFullRehash(t *testing.T) {
	for _, tc := range []struct {
		name  string
		field string
		value string
	}{
		{"decomposed edition", "edition", "E\u0301dition 2026"},
		{"decomposed revision", "revision", "Re\u0301vision 1"},
		{"decomposed Hangul", "edition", "가 2026"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			pkg := unicodeFixture(t)
			pkg["publication"].(map[string]any)[tc.field] = tc.value
			requested := rehashUnicodeCoordinates(t, pkg)
			if _, err := validate(pkg, requested); err == nil || !strings.Contains(err.Error(), "NFC") {
				t.Fatalf("expected NFC admission failure for fully rehashed %s, got %v", tc.field, err)
			}
		})
	}
}

func TestNFCCoordinatesObservePublicUnicodeLength(t *testing.T) {
	for _, n := range []int{128, 129} {
		pkg := unicodeFixture(t)
		pkg["publication"].(map[string]any)["edition"] = strings.Repeat("é", n)
		requested := rehashUnicodeCoordinates(t, pkg)
		_, err := validate(pkg, requested)
		if n == 128 && err != nil {
			t.Fatalf("valid 128-rune NFC edition rejected: %v", err)
		}
		if n == 129 && err == nil {
			t.Fatal("oversize 129-rune edition accepted")
		}
	}
}
