package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"testing"
)

const syntheticIdentities = `[{"id":"1","city":"90001","first_city":"FICTIONAL CITY"}]`
const syntheticGeometry = `[{"id":"1","city":"90001","first_city":"FICTIONAL CITY","the_geom":{"type":"MultiPolygon","coordinates":[[[[-105,40],[-104,40],[-104,41],[-105,40]]]]}}]`

func hashBytes(s string) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}

func TestPinnedGeometryMatchesSourceIdentities(t *testing.T) {
	got, err := identityAudit([]byte(syntheticGeometry), []byte(syntheticIdentities), hashBytes(syntheticIdentities), 1)
	if err != nil || got != hashBytes(syntheticGeometry) {
		t.Fatalf("failed to verify matching fictional records: hash %s err %v", got, err)
	}
}

func TestPinnedGeometryRejectsDrift(t *testing.T) {
	tests := []struct {
		name       string
		geometry   string
		identities string
		count      int
	}{
		{"place code drift", strings.Replace(syntheticGeometry, "90001", "90002", 1), syntheticIdentities, 1},
		{"name drift", strings.Replace(syntheticGeometry, "FICTIONAL CITY", "OTHER", 1), syntheticIdentities, 1},
		{"geometry missing", strings.Replace(syntheticGeometry, "MultiPolygon", "Point", 1), syntheticIdentities, 1},
		{"incomplete", syntheticGeometry, syntheticIdentities, 274},
		{"duplicate identity", syntheticGeometry, `[{"id":"1","city":"90001","first_city":"FICTIONAL CITY"},{"id":"1","city":"90002","first_city":"OTHER"}]`, 2},
		{"missing identity geometry separation", syntheticGeometry, syntheticGeometry, 1},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := identityAudit([]byte(tc.geometry), []byte(tc.identities), hashBytes(tc.identities), tc.count); err == nil {
				t.Fatal("accepted mismatched or malformed source")
			}
		})
	}
}

func TestPinnedGeometryRejectsInvalidHash(t *testing.T) {
	for _, hash := range []string{"", strings.Repeat("0", 64), strings.ToUpper(hashBytes(syntheticIdentities))} {
		if _, err := identityAudit([]byte(syntheticGeometry), []byte(syntheticIdentities), hash, 1); err == nil {
			t.Fatalf("accepted invalid identity SHA256: %q", hash)
		}
	}
}
