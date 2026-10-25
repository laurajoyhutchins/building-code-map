package geocoder

import (
	"bytes"
	"crypto/sha256"
	_ "embed"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"strings"
)

const RankingPolicyPathEnvironment = "GEOCODER_RANKING_POLICY_PATH"

type rankingPolicyDocument struct {
	Version                    string             `json:"version"`
	AddressPointBase           float64            `json:"address_point_base"`
	StreetRangeBase            float64            `json:"street_range_base"`
	ExactStreet                float64            `json:"exact_street"`
	ExactCity                  float64            `json:"exact_city"`
	ExactPostalCode            float64            `json:"exact_postal_code"`
	PostalCodeNotSupplied      float64            `json:"postal_code_not_supplied"`
	ParityMatched              float64            `json:"parity_matched"`
	MinimumAddressPointQuality float64            `json:"minimum_address_point_quality"`
	MinimumStreetRangeQuality  float64            `json:"minimum_street_range_quality"`
	AmbiguityGap               float64            `json:"ambiguity_gap"`
	SourcePriority             map[string]float64 `json:"source_priority"`
}

var requiredRankingPolicyFields = []string{
	"version",
	"address_point_base",
	"street_range_base",
	"exact_street",
	"exact_city",
	"exact_postal_code",
	"postal_code_not_supplied",
	"parity_matched",
	"minimum_address_point_quality",
	"minimum_street_range_quality",
	"ambiguity_gap",
	"source_priority",
}

//go:embed config/geocoder-ranking-policy.json
var defaultRankingPolicyJSON []byte

func LoadDefaultRankingPolicy() (RankingPolicy, string, error) {
	if path := strings.TrimSpace(os.Getenv(RankingPolicyPathEnvironment)); path != "" {
		return LoadRankingPolicy(path)
	}
	return loadRankingPolicy(defaultRankingPolicyJSON)
}

func LoadRankingPolicy(path string) (RankingPolicy, string, error) {
	if strings.TrimSpace(path) == "" {
		return RankingPolicy{}, "", errors.New("geocoder ranking policy path is required")
	}
	contents, err := os.ReadFile(path)
	if err != nil {
		return RankingPolicy{}, "", fmt.Errorf("read geocoder ranking policy: %w", err)
	}
	return loadRankingPolicy(contents)
}

func loadRankingPolicy(contents []byte) (RankingPolicy, string, error) {
	var fields map[string]json.RawMessage
	if err := json.Unmarshal(contents, &fields); err != nil {
		return RankingPolicy{}, "", fmt.Errorf("decode geocoder ranking policy: %w", err)
	}
	for _, name := range requiredRankingPolicyFields {
		if _, ok := fields[name]; !ok {
			return RankingPolicy{}, "", fmt.Errorf("geocoder ranking policy field %q is required", name)
		}
	}

	decoder := json.NewDecoder(bytes.NewReader(contents))
	decoder.DisallowUnknownFields()
	var document rankingPolicyDocument
	if err := decoder.Decode(&document); err != nil {
		return RankingPolicy{}, "", fmt.Errorf("decode geocoder ranking policy: %w", err)
	}
	policy := RankingPolicy{
		Version:                    document.Version,
		AddressPointBase:           document.AddressPointBase,
		StreetRangeBase:            document.StreetRangeBase,
		ExactStreet:                document.ExactStreet,
		ExactCity:                  document.ExactCity,
		ExactPostalCode:            document.ExactPostalCode,
		PostalCodeNotSupplied:      document.PostalCodeNotSupplied,
		ParityMatched:              document.ParityMatched,
		MinimumAddressPointQuality: document.MinimumAddressPointQuality,
		MinimumStreetRangeQuality:  document.MinimumStreetRangeQuality,
		AmbiguityGap:               document.AmbiguityGap,
		SourcePriority:             document.SourcePriority,
	}
	if err := ValidateRankingPolicy(policy); err != nil {
		return RankingPolicy{}, "", err
	}
	identity, err := rankingPolicyIdentity(policy)
	if err != nil {
		return RankingPolicy{}, "", err
	}
	return policy, identity, nil
}

func rankingPolicyIdentity(policy RankingPolicy) (string, error) {
	if err := ValidateRankingPolicy(policy); err != nil {
		return "", err
	}
	document := rankingPolicyDocument{
		Version:                    policy.Version,
		AddressPointBase:           policy.AddressPointBase,
		StreetRangeBase:            policy.StreetRangeBase,
		ExactStreet:                policy.ExactStreet,
		ExactCity:                  policy.ExactCity,
		ExactPostalCode:            policy.ExactPostalCode,
		PostalCodeNotSupplied:      policy.PostalCodeNotSupplied,
		ParityMatched:              policy.ParityMatched,
		MinimumAddressPointQuality: policy.MinimumAddressPointQuality,
		MinimumStreetRangeQuality:  policy.MinimumStreetRangeQuality,
		AmbiguityGap:               policy.AmbiguityGap,
		SourcePriority:             policy.SourcePriority,
	}
	canonical, err := json.Marshal(document)
	if err != nil {
		return "", fmt.Errorf("canonicalize geocoder ranking policy: %w", err)
	}
	digest := sha256.Sum256(canonical)
	return "sha256:" + hex.EncodeToString(digest[:]), nil
}

func ValidateRankingPolicy(policy RankingPolicy) error {
	if strings.TrimSpace(policy.Version) == "" {
		return errors.New("geocoder ranking policy version is required")
	}
	weights := map[string]float64{
		"address_point_base":       policy.AddressPointBase,
		"street_range_base":        policy.StreetRangeBase,
		"exact_street":             policy.ExactStreet,
		"exact_city":               policy.ExactCity,
		"exact_postal_code":        policy.ExactPostalCode,
		"postal_code_not_supplied": policy.PostalCodeNotSupplied,
		"parity_matched":           policy.ParityMatched,
	}
	for name, value := range weights {
		if !finiteUnitInterval(value) {
			return fmt.Errorf("geocoder ranking policy %s must be finite and between 0 and 1", name)
		}
	}
	thresholds := map[string]float64{
		"minimum_address_point_quality": policy.MinimumAddressPointQuality,
		"minimum_street_range_quality":  policy.MinimumStreetRangeQuality,
		"ambiguity_gap":                 policy.AmbiguityGap,
	}
	for name, value := range thresholds {
		if !finiteUnitInterval(value) {
			return fmt.Errorf("geocoder ranking policy %s must be finite and between 0 and 1", name)
		}
	}
	if policy.SourcePriority == nil {
		return errors.New("geocoder ranking policy source_priority is required")
	}
	for source, weight := range policy.SourcePriority {
		if strings.TrimSpace(source) == "" || source != strings.TrimSpace(source) {
			return errors.New("geocoder ranking policy source_priority keys must be non-empty and trimmed")
		}
		if !finiteUnitInterval(weight) {
			return fmt.Errorf("geocoder ranking policy source priority %q must be finite and between 0 and 1", source)
		}
	}
	return nil
}

func finiteUnitInterval(value float64) bool {
	return !math.IsNaN(value) && !math.IsInf(value, 0) && value >= 0 && value <= 1
}
