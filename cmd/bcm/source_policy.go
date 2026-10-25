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
	"time"

	"github.com/laurajoyhutchins/building-code-map/engine"
)

const maxSourcePolicyBytes int64 = 1 << 20

// loadPinnedSourcePolicy loads trusted policy from operator-configured, private
// bytes. The remote Provider does not select its own evidentiary authority.
// A matching digest only proves byte identity; review/rights remain external.
func loadPinnedSourcePolicy(path, expectedSHA string) (engine.SourcePolicy, error) {
	if strings.TrimSpace(path) == "" || len(expectedSHA) != 64 {
		return engine.SourcePolicy{}, errors.New("source policy needs a path and exact SHA-256")
	}
	expected, err := hex.DecodeString(expectedSHA)
	if err != nil || len(expected) != sha256.Size {
		return engine.SourcePolicy{}, errors.New("source policy SHA-256 is invalid")
	}
	f, err := os.Open(path)
	if err != nil {
		return engine.SourcePolicy{}, fmt.Errorf("cannot read pinned source policy: %w", err)
	}
	defer f.Close()
	data, err := io.ReadAll(io.LimitReader(f, maxSourcePolicyBytes+1))
	if err != nil || int64(len(data)) > maxSourcePolicyBytes {
		return engine.SourcePolicy{}, errors.New("source policy unavailable or too large")
	}
	calculated := sha256.Sum256(data)
	if !bytes.Equal(calculated[:], expected) {
		return engine.SourcePolicy{}, errors.New("source policy content digest mismatch")
	}
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	var policy engine.SourcePolicy
	if err := decoder.Decode(&policy); err != nil {
		return engine.SourcePolicy{}, errors.New("invalid source policy JSON")
	}
	var trailing any
	if err := decoder.Decode(&trailing); err != io.EOF {
		return engine.SourcePolicy{}, errors.New("source policy contains trailing content")
	}
	if len(policy.Rules) == 0 {
		return engine.SourcePolicy{}, errors.New("source policy contains no rules")
	}
	ids := map[string]struct{}{}
	for _, rule := range policy.Rules {
		if strings.TrimSpace(rule.ID) == "" || strings.TrimSpace(rule.DocumentID) == "" ||
			strings.TrimSpace(rule.JurisdictionID) == "" ||
			(rule.ClaimKind != "code_adoption" && rule.ClaimKind != "code_non_applicability") {
			return engine.SourcePolicy{}, errors.New("source policy contains invalid rule identity")
		}
		// Only independently admitted immutable original bytes and a reviewed
		// exact passage may establish legal materiality. A digest-shaped
		// value asserted by the provider is never sufficient.
		if len(rule.ArtifactSHA256) != 64 || len(rule.TextSHA256) != 64 ||
			strings.TrimSpace(rule.LegalLocator) == "" {
			return engine.SourcePolicy{}, errors.New("source policy requires pinned artifact and passage identities")
		}
		for _, digest := range []string{rule.ArtifactSHA256, rule.TextSHA256} {
			decoded, err := hex.DecodeString(digest)
			if err != nil || len(decoded) != sha256.Size || digest != strings.ToLower(digest) {
				return engine.SourcePolicy{}, errors.New("source policy contains malformed source digest")
			}
		}
		if _, found := ids[rule.ID]; found {
			return engine.SourcePolicy{}, errors.New("source policy contains duplicate rule ID")
		}
		ids[rule.ID] = struct{}{}
		var start, end time.Time
		if rule.EffectiveFrom != "" {
			start, err = time.Parse(time.DateOnly, rule.EffectiveFrom)
			if err != nil {
				return engine.SourcePolicy{}, errors.New("source policy has malformed start date")
			}
		}
		if rule.EffectiveTo != "" {
			end, err = time.Parse(time.DateOnly, rule.EffectiveTo)
			if err != nil {
				return engine.SourcePolicy{}, errors.New("source policy has malformed end date")
			}
		}
		if !start.IsZero() && !end.IsZero() && end.Before(start) {
			return engine.SourcePolicy{}, errors.New("source policy has reversed date range")
		}
	}
	return policy, nil
}
