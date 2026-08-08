package geocoder

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

var (
	poBoxPattern = regexp.MustCompile(`(?i)^\s*P(?:OST)?\.?\s*O(?:FFICE)?\.?\s+BOX\b`)
	housePattern = regexp.MustCompile(`^([0-9]+[A-Z]?)\s+(.+)$`)
	zipPattern   = regexp.MustCompile(`^([0-9]{5})(?:-[0-9]{4})?$`)
)

var directionAliases = map[string]string{
	"NORTH":     "N",
	"SOUTH":     "S",
	"EAST":      "E",
	"WEST":      "W",
	"NORTHEAST": "NE",
	"NORTHWEST": "NW",
	"SOUTHEAST": "SE",
	"SOUTHWEST": "SW",
}

var suffixAliases = map[string]string{
	"ALLEY":      "ALY",
	"AVENUE":     "AVE",
	"BOULEVARD":  "BLVD",
	"CIRCLE":     "CIR",
	"COURT":      "CT",
	"DRIVE":      "DR",
	"EXPRESSWAY": "EXPY",
	"HIGHWAY":    "HWY",
	"LANE":       "LN",
	"PARKWAY":    "PKWY",
	"PLACE":      "PL",
	"ROAD":       "RD",
	"SQUARE":     "SQ",
	"STREET":     "ST",
	"TERRACE":    "TER",
	"TRAIL":      "TRL",
	"WAY":        "WAY",
}

var stateAliases = map[string]string{
	"ALABAMA": "AL", "ALASKA": "AK", "ARIZONA": "AZ", "ARKANSAS": "AR",
	"CALIFORNIA": "CA", "COLORADO": "CO", "CONNECTICUT": "CT", "DELAWARE": "DE",
	"DISTRICT OF COLUMBIA": "DC", "FLORIDA": "FL", "GEORGIA": "GA", "HAWAII": "HI",
	"IDAHO": "ID", "ILLINOIS": "IL", "INDIANA": "IN", "IOWA": "IA",
	"KANSAS": "KS", "KENTUCKY": "KY", "LOUISIANA": "LA", "MAINE": "ME",
	"MARYLAND": "MD", "MASSACHUSETTS": "MA", "MICHIGAN": "MI", "MINNESOTA": "MN",
	"MISSISSIPPI": "MS", "MISSOURI": "MO", "MONTANA": "MT", "NEBRASKA": "NE",
	"NEVADA": "NV", "NEW HAMPSHIRE": "NH", "NEW JERSEY": "NJ", "NEW MEXICO": "NM",
	"NEW YORK": "NY", "NORTH CAROLINA": "NC", "NORTH DAKOTA": "ND", "OHIO": "OH",
	"OKLAHOMA": "OK", "OREGON": "OR", "PENNSYLVANIA": "PA", "RHODE ISLAND": "RI",
	"SOUTH CAROLINA": "SC", "SOUTH DAKOTA": "SD", "TENNESSEE": "TN", "TEXAS": "TX",
	"UTAH": "UT", "VERMONT": "VT", "VIRGINIA": "VA", "WASHINGTON": "WA",
	"WEST VIRGINIA": "WV", "WISCONSIN": "WI", "WYOMING": "WY",
}

func ParseAddress(value string) (ParsedAddress, error) {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" || poBoxPattern.MatchString(trimmed) {
		return ParsedAddress{}, ErrInvalidAddress
	}

	parts := strings.Split(trimmed, ",")
	if len(parts) != 3 {
		return ParsedAddress{}, ErrInvalidAddress
	}

	streetPart := normalizeWords(parts[0])
	match := housePattern.FindStringSubmatch(streetPart)
	if len(match) != 3 {
		return ParsedAddress{}, ErrInvalidAddress
	}
	houseNumber := match[1]
	street := normalizeStreet(match[2])
	city := normalizeWords(parts[1])
	if street == "" || city == "" {
		return ParsedAddress{}, ErrInvalidAddress
	}

	stateAndZip := strings.Fields(strings.TrimSpace(parts[2]))
	if len(stateAndZip) == 0 {
		return ParsedAddress{}, ErrInvalidAddress
	}
	postalCode := ""
	if zipMatch := zipPattern.FindStringSubmatch(stateAndZip[len(stateAndZip)-1]); len(zipMatch) == 2 {
		postalCode = zipMatch[1]
		stateAndZip = stateAndZip[:len(stateAndZip)-1]
	}
	state, ok := normalizeState(strings.Join(stateAndZip, " "))
	if !ok {
		return ParsedAddress{}, fmt.Errorf("%w: use a United States state name or abbreviation", ErrInvalidAddress)
	}

	normalized := fmt.Sprintf("%s %s, %s, %s", houseNumber, street, city, state)
	if postalCode != "" {
		normalized += " " + postalCode
	}
	return ParsedAddress{
		HouseNumber: houseNumber,
		Street:      street,
		City:        city,
		State:       state,
		PostalCode:  postalCode,
		Normalized:  normalized,
	}, nil
}

func normalizeStreet(value string) string {
	words := strings.Fields(normalizeWords(value))
	for index, word := range words {
		if replacement, ok := directionAliases[word]; ok {
			words[index] = replacement
		}
	}
	if len(words) > 0 {
		last := len(words) - 1
		if replacement, ok := suffixAliases[words[last]]; ok {
			words[last] = replacement
		}
	}
	return strings.Join(words, " ")
}

func normalizeWords(value string) string {
	var builder strings.Builder
	lastWasSpace := true
	for _, char := range strings.ToUpper(strings.TrimSpace(value)) {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			builder.WriteRune(char)
			lastWasSpace = false
			continue
		}
		if !lastWasSpace {
			builder.WriteByte(' ')
			lastWasSpace = true
		}
	}
	return strings.TrimSpace(builder.String())
}

func normalizeState(value string) (string, bool) {
	normalized := normalizeWords(value)
	if len(normalized) == 2 {
		for _, abbreviation := range stateAliases {
			if abbreviation == normalized {
				return normalized, true
			}
		}
		return "", false
	}
	abbreviation, ok := stateAliases[normalized]
	return abbreviation, ok
}
