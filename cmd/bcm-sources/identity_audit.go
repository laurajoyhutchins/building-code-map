package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// identityAudit proves that candidate geometry matches a pinned DOLA identity
// projection. Neither source establishes a legally authoritative boundary.
func identityAudit(geometry, identities []byte, pinnedHash string, expected int) (string, error) {
	if expected < 1 || expected > 10000 {
		return "", errors.New("expected record count required")
	}
	if len(pinnedHash) != 64 || strings.Trim(pinnedHash, "0123456789abcdef") != "" {
		return "", errors.New("identity SHA256 must be pinned lowercase hex")
	}
	if len(geometry) > 256<<20 || len(identities) > 2<<20 {
		return "", errors.New("source is too large")
	}
	hash := sha256.Sum256(identities)
	if hex.EncodeToString(hash[:]) != pinnedHash {
		return "", errors.New("identity source hash mismatch")
	}
	type row struct {
		ID       json.Number     `json:"id"`
		Code     string          `json:"city"`
		Name     string          `json:"first_city"`
		Geometry json.RawMessage `json:"the_geom"`
	}
	var names []row
	var shapes []row
	if err := json.Unmarshal(identities, &names); err != nil {
		return "", err
	}
	if err := json.Unmarshal(geometry, &shapes); err != nil {
		return "", err
	}
	if len(names) != expected || len(shapes) != expected {
		return "", errors.New("incomplete or changed municipal dataset")
	}
	known := make(map[string]row, expected)
	codes := make(map[string]bool, expected)
	for _, x := range names {
		id := string(x.ID)
		n, err := strconv.Atoi(id)
		if err != nil || n < 1 || strconv.Itoa(n) != id || known[id].Name != "" || len(x.Geometry) != 0 {
			return "", fmt.Errorf("invalid or duplicate identity record %s", id)
		}
		if len(x.Code) != 5 || strings.Trim(x.Code, "0123456789") != "" || codes[x.Code] || strings.TrimSpace(x.Name) == "" {
			return "", fmt.Errorf("invalid identity code/name at %s", id)
		}
		known[id] = x
		codes[x.Code] = true
	}
	seen := make(map[string]bool, expected)
	for _, x := range shapes {
		id := string(x.ID)
		y, ok := known[id]
		if !ok || seen[id] || x.Code != y.Code || x.Name != y.Name {
			return "", fmt.Errorf("geometry identity mismatch at %s", id)
		}
		var geometryType struct {
			Type string `json:"type"`
		}
		if json.Unmarshal(x.Geometry, &geometryType) != nil || geometryType.Type != "MultiPolygon" {
			return "", fmt.Errorf("missing multipolygon at %s", id)
		}
		seen[id] = true
	}
	sum := sha256.Sum256(geometry)
	return hex.EncodeToString(sum[:]), nil
}

func auditColoradoFiles(geometryPath, identityPath, pinnedHash string, count int) (string, error) {
	geo, err := os.ReadFile(geometryPath)
	if err != nil {
		return "", err
	}
	ids, err := os.ReadFile(identityPath)
	if err != nil {
		return "", err
	}
	return identityAudit(geo, ids, pinnedHash, count)
}
