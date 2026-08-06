package geocoder

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	_ "modernc.org/sqlite"
)

type SQLiteService struct {
	db     *sql.DB
	policy RankingPolicy
}

func Open(path string) (*SQLiteService, error) {
	return OpenWithPolicy(path, DefaultRankingPolicy())
}

func OpenWithPolicy(path string, policy RankingPolicy) (*SQLiteService, error) {
	if strings.TrimSpace(path) == "" {
		return nil, errors.New("geocoder snapshot path is required")
	}
	if strings.TrimSpace(policy.Version) == "" {
		return nil, errors.New("geocoder ranking policy version is required")
	}
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1)
	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, err
	}
	var version string
	if err := db.QueryRow(`SELECT value FROM geocoder_metadata WHERE key = 'schema_version'`).Scan(&version); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("read geocoder schema version: %w", err)
	}
	if version != schemaVersion {
		_ = db.Close()
		return nil, fmt.Errorf("unsupported geocoder schema version %q", version)
	}
	if _, err := db.Exec(`PRAGMA query_only = ON;`); err != nil {
		_ = db.Close()
		return nil, err
	}
	return &SQLiteService{db: db, policy: policy}, nil
}

func (service *SQLiteService) Close() error {
	if service == nil || service.db == nil {
		return nil
	}
	return service.db.Close()
}

func ResolveDataPath(root, candidate string) (string, error) {
	rootPath, err := filepath.Abs(root)
	if err != nil {
		return "", err
	}
	candidatePath := candidate
	if !filepath.IsAbs(candidatePath) {
		candidatePath = filepath.Join(rootPath, candidatePath)
	}
	candidatePath, err = filepath.Abs(candidatePath)
	if err != nil {
		return "", err
	}
	relative, err := filepath.Rel(rootPath, candidatePath)
	if err != nil {
		return "", err
	}
	if relative == ".." || strings.HasPrefix(relative, ".."+string(filepath.Separator)) {
		return "", errors.New("geocoder snapshot path must stay inside the backend checkout")
	}
	return candidatePath, nil
}

func (service *SQLiteService) Geocode(ctx context.Context, query Query) (Result, error) {
	parsed, err := ParseAddress(query.Address)
	if err != nil {
		return Result{}, err
	}
	result := Result{
		Query:      query.Address,
		Normalized: parsed.Normalized,
		Status:     StatusNotFound,
		Candidates: []Candidate{},
		Warnings:   []string{},
	}

	addressCandidates, err := service.addressPointCandidates(ctx, parsed)
	if err != nil {
		return Result{}, err
	}
	if selected := chooseCandidatesWithGap(result, addressCandidates, service.policy.MinimumAddressPointQuality, service.policy.AmbiguityGap); selected.Status != StatusNotFound {
		return selected, nil
	}

	rangeCandidates, err := service.streetRangeCandidates(ctx, parsed)
	if err != nil {
		return Result{}, err
	}
	selected := chooseCandidatesWithGap(result, rangeCandidates, service.policy.MinimumStreetRangeQuality, service.policy.AmbiguityGap)
	if selected.Status == StatusMatched {
		selected.Warnings = append(selected.Warnings, "The point was interpolated from a street range and is not a parcel, entrance, or address-point location.")
	}
	return selected, nil
}

func (service *SQLiteService) addressPointCandidates(ctx context.Context, parsed ParsedAddress) ([]Candidate, error) {
	rows, err := service.db.QueryContext(ctx, `
SELECT matched_address, longitude, latitude, source_name, source_record_id, source_vintage,
       street_name, city, postal_code
FROM address_points
WHERE address_number = ?
  AND (street_name = ? OR street_name LIKE ?)
  AND city = ?
  AND state = ?
  AND (? = '' OR postal_code = ?)
ORDER BY source_name, source_record_id
LIMIT 20;`,
		parsed.HouseNumber,
		parsed.Street,
		parsed.Street+" %",
		parsed.City,
		parsed.State,
		parsed.PostalCode,
		parsed.PostalCode,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	candidates := []Candidate{}
	for rows.Next() {
		var candidate Candidate
		var street, city, postalCode string
		if err := rows.Scan(
			&candidate.MatchedAddress,
			&candidate.Longitude,
			&candidate.Latitude,
			&candidate.Source,
			&candidate.SourceRecordID,
			&candidate.SourceVintage,
			&street,
			&city,
			&postalCode,
		); err != nil {
			return nil, err
		}
		candidate.Precision = PrecisionAddressPoint
		candidate.ScoreKind = "deterministic_quality"
		candidate.RankingPolicyVersion = service.policy.Version
		candidate.ScoreFactors = map[string]float64{"address_point_base": service.policy.AddressPointBase}
		candidate.Confidence = service.policy.AddressPointBase
		if street == parsed.Street {
			candidate.Confidence += addFactor(candidate.ScoreFactors, "exact_street", service.policy.ExactStreet)
		}
		if city == parsed.City {
			candidate.Confidence += addFactor(candidate.ScoreFactors, "exact_city", service.policy.ExactCity)
		}
		if parsed.PostalCode == "" {
			candidate.Confidence += addFactor(candidate.ScoreFactors, "postal_not_supplied", service.policy.PostalCodeNotSupplied)
		} else if postalCode == parsed.PostalCode {
			candidate.Confidence += addFactor(candidate.ScoreFactors, "exact_postal_code", service.policy.ExactPostalCode)
		}
		candidate.Confidence += addFactor(candidate.ScoreFactors, "source_priority", service.policy.sourcePriority(candidate.Source))
		candidates = append(candidates, candidate)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return candidates, nil
}

func (service *SQLiteService) streetRangeCandidates(ctx context.Context, parsed ParsedAddress) ([]Candidate, error) {
	houseNumber, err := strconv.Atoi(parsed.HouseNumber)
	if err != nil {
		return nil, nil
	}
	rows, err := service.db.QueryContext(ctx, `
SELECT from_number, to_number, parity, street_name, city, postal_code,
       from_longitude, from_latitude, to_longitude, to_latitude,
       source_name, source_record_id, source_vintage
FROM street_ranges
WHERE (street_name = ? OR street_name LIKE ?)
  AND city = ?
  AND state = ?
  AND (? = '' OR postal_code = ?)
  AND ? BETWEEN MIN(from_number, to_number) AND MAX(from_number, to_number)
ORDER BY source_name, source_record_id
LIMIT 20;`,
		parsed.Street,
		parsed.Street+" %",
		parsed.City,
		parsed.State,
		parsed.PostalCode,
		parsed.PostalCode,
		houseNumber,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	candidates := []Candidate{}
	for rows.Next() {
		var fromNumber, toNumber int
		var parity, street, city, postalCode string
		var fromLongitude, fromLatitude, toLongitude, toLatitude float64
		var source, sourceRecordID, sourceVintage string
		if err := rows.Scan(
			&fromNumber,
			&toNumber,
			&parity,
			&street,
			&city,
			&postalCode,
			&fromLongitude,
			&fromLatitude,
			&toLongitude,
			&toLatitude,
			&source,
			&sourceRecordID,
			&sourceVintage,
		); err != nil {
			return nil, err
		}
		if fromNumber == toNumber || !parityMatches(parity, houseNumber) {
			continue
		}
		ratio := float64(houseNumber-fromNumber) / float64(toNumber-fromNumber)
		longitude := fromLongitude + ratio*(toLongitude-fromLongitude)
		latitude := fromLatitude + ratio*(toLatitude-fromLatitude)
		direction := "ascending"
		if toNumber < fromNumber {
			direction = "descending"
		}
		candidate := Candidate{
			MatchedAddress:       parsed.Normalized,
			Longitude:            longitude,
			Latitude:             latitude,
			Precision:            PrecisionInterpolated,
			Confidence:           service.policy.StreetRangeBase,
			ScoreKind:            "deterministic_quality",
			ScoreFactors:         map[string]float64{"street_range_base": service.policy.StreetRangeBase},
			RankingPolicyVersion: service.policy.Version,
			Source:               source,
			SourceRecordID:       sourceRecordID,
			SourceVintage:        sourceVintage,
			Interpolation: &InterpolationProvenance{
				SourceRangeID:             sourceRecordID,
				RequestedHouseNumber:      houseNumber,
				FromNumber:                fromNumber,
				ToNumber:                  toNumber,
				RangeDirection:            direction,
				Parity:                    parity,
				FromCoordinate:            Coordinate{Longitude: fromLongitude, Latitude: fromLatitude},
				ToCoordinate:              Coordinate{Longitude: toLongitude, Latitude: toLatitude},
				Fraction:                  ratio,
				DerivedCoordinate:         Coordinate{Longitude: longitude, Latitude: latitude},
				CoordinateReferenceSystem: "EPSG:4326",
				TransformationIdentity:    "none",
				MethodVersion:             InterpolationMethodVersion,
				PositionalQuality:         "street_range_interpolation",
			},
		}
		if street == parsed.Street {
			candidate.Confidence += addFactor(candidate.ScoreFactors, "exact_street", service.policy.ExactStreet)
		}
		if city == parsed.City {
			candidate.Confidence += addFactor(candidate.ScoreFactors, "exact_city", service.policy.ExactCity)
		}
		if parsed.PostalCode == "" {
			candidate.Confidence += addFactor(candidate.ScoreFactors, "postal_not_supplied", service.policy.PostalCodeNotSupplied)
		} else if postalCode == parsed.PostalCode {
			candidate.Confidence += addFactor(candidate.ScoreFactors, "exact_postal_code", service.policy.ExactPostalCode)
		}
		candidate.Confidence += addFactor(candidate.ScoreFactors, "parity_matched", service.policy.ParityMatched)
		candidate.Confidence += addFactor(candidate.ScoreFactors, "source_priority", service.policy.sourcePriority(candidate.Source))
		candidates = append(candidates, candidate)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return candidates, nil
}

func chooseCandidates(base Result, candidates []Candidate, minimumConfidence float64) Result {
	return chooseCandidatesWithGap(base, candidates, minimumConfidence, DefaultRankingPolicy().AmbiguityGap)
}

func chooseCandidatesWithGap(base Result, candidates []Candidate, minimumConfidence, ambiguityGap float64) Result {
	sort.SliceStable(candidates, func(left, right int) bool {
		if candidates[left].Confidence == candidates[right].Confidence {
			if candidates[left].Source != candidates[right].Source {
				return candidates[left].Source < candidates[right].Source
			}
			return candidates[left].SourceRecordID < candidates[right].SourceRecordID
		}
		return candidates[left].Confidence > candidates[right].Confidence
	})
	if len(candidates) == 0 || candidates[0].Confidence < minimumConfidence {
		return base
	}

	eligible := make([]Candidate, 0, len(candidates))
	for _, candidate := range candidates {
		if candidate.Confidence >= minimumConfidence {
			eligible = append(eligible, candidate)
		}
	}
	base.Candidates = eligible
	if len(eligible) > 1 && eligible[0].Confidence-eligible[1].Confidence < ambiguityGap {
		base.Status = StatusAmbiguous
		return base
	}
	selected := eligible[0]
	base.Status = StatusMatched
	base.Selected = &selected
	return base
}

func parityMatches(parity string, number int) bool {
	switch parity {
	case "B":
		return true
	case "E":
		return number%2 == 0
	case "O":
		return number%2 != 0
	default:
		return false
	}
}
