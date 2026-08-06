package geocoder

const (
	DefaultRankingPolicyVersion = "geocoder-ranking-1.0"
	InterpolationMethodVersion  = "linear-street-range-1.0"
)

type RankingPolicy struct {
	Version                       string
	AddressPointBase              float64
	StreetRangeBase               float64
	ExactStreet                   float64
	ExactCity                     float64
	ExactPostalCode               float64
	PostalCodeNotSupplied         float64
	ParityMatched                 float64
	MinimumAddressPointQuality    float64
	MinimumStreetRangeQuality     float64
	AmbiguityGap                  float64
	SourcePriority                map[string]float64
}

func DefaultRankingPolicy() RankingPolicy {
	return RankingPolicy{
		Version:                    DefaultRankingPolicyVersion,
		AddressPointBase:           0.70,
		StreetRangeBase:            0.55,
		ExactStreet:                0.05,
		ExactCity:                  0.15,
		ExactPostalCode:            0.10,
		PostalCodeNotSupplied:      0.05,
		ParityMatched:              0.05,
		MinimumAddressPointQuality: 0.85,
		MinimumStreetRangeQuality:  0.75,
		AmbiguityGap:               0.05,
		SourcePriority:             map[string]float64{},
	}
}

func (policy RankingPolicy) sourcePriority(source string) float64 {
	return policy.SourcePriority[source]
}

func addFactor(factors map[string]float64, name string, value float64) float64 {
	if value == 0 {
		return 0
	}
	factors[name] = value
	return value
}
