package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/laurajoyhutchins/building-code-map/engine"
)

// This is an operator-pinned, protected review artifact, NOT a provider claim.
// A matching SHA-256 proves byte identity only; source custody, independent
// polygon/annexation review, and reviewer authority are separate obligations.
const maxSiteAuthorityPolicyBytes = 1 << 20

// rejectSiteDuplicateKeys recursively rejects duplicate field names before
// Go's JSON decoder can silently replace one geographic/legal witness.
func rejectSiteDuplicateKeys(dec *json.Decoder) error {
	token, err := dec.Token()
	if err != nil {
		return errors.New("site policy JSON cannot be decoded")
	}
	delimiter, isDelimiter := token.(json.Delim)
	if !isDelimiter {
		return nil
	}
	switch delimiter {
	case '{':
		seen := map[string]struct{}{}
		for dec.More() {
			keyToken, err := dec.Token()
			if err != nil {
				return errors.New("site policy object key is invalid")
			}
			key, ok := keyToken.(string)
			if !ok {
				return errors.New("site policy object key is not a string")
			}
			if _, exists := seen[key]; exists {
				return errors.New("site policy contains duplicate JSON keys")
			}
			seen[key] = struct{}{}
			if err := rejectSiteDuplicateKeys(dec); err != nil {
				return err
			}
		}
		_, err = dec.Token()
		if err != nil {
			return errors.New("site policy JSON object is incomplete")
		}
	case '[':
		for dec.More() {
			if err := rejectSiteDuplicateKeys(dec); err != nil {
				return err
			}
		}
		_, err = dec.Token()
		if err != nil {
			return errors.New("site policy JSON array is incomplete")
		}
	default:
		return errors.New("unexpected site policy JSON delimiter")
	}
	return nil
}

func loadPinnedSiteAuthorityPolicy(path, expectedSHA string) (engine.SiteAuthorityPolicy, error) {
	var empty engine.SiteAuthorityPolicy
	if strings.TrimSpace(path) == "" || len(expectedSHA) != 64 ||
		strings.ToLower(expectedSHA) != expectedSHA {
		return empty, errors.New("site authority policy needs a path and exact lowercase SHA-256")
	}
	expected, err := hex.DecodeString(expectedSHA)
	if err != nil || len(expected) != sha256.Size {
		return empty, errors.New("site authority policy digest is malformed")
	}
	file, err := os.Open(path)
	if err != nil {
		return empty, fmt.Errorf("cannot read protected site authority policy: %w", err)
	}
	defer file.Close()
	raw, err := io.ReadAll(io.LimitReader(file, maxSiteAuthorityPolicyBytes+1))
	if err != nil || len(raw) == 0 || len(raw) > maxSiteAuthorityPolicyBytes {
		return empty, errors.New("site authority policy unavailable or oversized")
	}
	observed := sha256.Sum256(raw)
	if !bytes.Equal(observed[:], expected) {
		return empty, errors.New("site authority policy SHA-256 mismatch")
	}
	keys := json.NewDecoder(bytes.NewReader(raw))
	if err := rejectSiteDuplicateKeys(keys); err != nil {
		return empty, err
	}
	if _, err := keys.Token(); err != io.EOF {
		return empty, errors.New("site authority policy contains trailing JSON")
	}
	decoder := json.NewDecoder(bytes.NewReader(raw))
	decoder.DisallowUnknownFields()
	var policy engine.SiteAuthorityPolicy
	if err := decoder.Decode(&policy); err != nil {
		return empty, errors.New("site authority policy has unknown or invalid fields")
	}
	if _, err := decoder.Token(); err != io.EOF {
		return empty, errors.New("site authority policy has trailing content")
	}
	if policy.SchemaVersion != "1.0.0" || len(policy.Rules) == 0 {
		return empty, errors.New("site authority policy version or rules missing")
	}
	ids := map[string]bool{}
	for _, rule := range policy.Rules {
		if ids[rule.ID] || !engine.ValidateSiteAuthorityRule(rule) {
			return empty, errors.New("site authority policy contains duplicate or incomplete reviewed rule")
		}
		ids[rule.ID] = true
	}
	return policy, nil
}
