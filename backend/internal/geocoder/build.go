package geocoder

import (
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	_ "modernc.org/sqlite"
)

func BuildSnapshot(options BuildOptions) error {
	if strings.TrimSpace(options.OutputPath) == "" {
		return errors.New("output path is required")
	}
	if strings.TrimSpace(options.SourceName) == "" || strings.TrimSpace(options.SourceVintage) == "" {
		return errors.New("source name and source vintage are required")
	}
	if strings.TrimSpace(options.AddressPointsCSV) == "" && strings.TrimSpace(options.StreetRangesCSV) == "" {
		return errors.New("at least one address-point or street-range CSV is required")
	}

	outputPath, err := filepath.Abs(options.OutputPath)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(outputPath), 0o755); err != nil {
		return err
	}
	temporary, err := os.CreateTemp(filepath.Dir(outputPath), ".geocoder-*.sqlite")
	if err != nil {
		return err
	}
	temporaryPath := temporary.Name()
	if err := temporary.Close(); err != nil {
		return err
	}
	if err := os.Remove(temporaryPath); err != nil {
		return err
	}
	defer os.Remove(temporaryPath)

	db, err := sql.Open("sqlite", temporaryPath)
	if err != nil {
		return err
	}
	closed := false
	defer func() {
		if !closed {
			_ = db.Close()
		}
	}()
	if _, err := db.Exec(schemaSQL); err != nil {
		return fmt.Errorf("create geocoder schema: %w", err)
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	if _, err := tx.Exec(`INSERT INTO geocoder_metadata(key, value) VALUES
		('schema_version', ?),
		('source_name', ?),
		('source_vintage', ?);`, schemaVersion, options.SourceName, options.SourceVintage); err != nil {
		_ = tx.Rollback()
		return err
	}
	if options.AddressPointsCSV != "" {
		if err := importAddressPoints(tx, options); err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	if options.StreetRangesCSV != "" {
		if err := importStreetRanges(tx, options); err != nil {
			_ = tx.Rollback()
			return err
		}
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	if err := db.Close(); err != nil {
		return err
	}
	closed = true

	if err := os.Rename(temporaryPath, outputPath); err == nil {
		return nil
	}
	if err := os.Remove(outputPath); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	if err := os.Rename(temporaryPath, outputPath); err != nil {
		return fmt.Errorf("replace geocoder snapshot: %w", err)
	}
	return nil
}

func importAddressPoints(tx *sql.Tx, options BuildOptions) error {
	required := []string{
		"source_record_id", "address_number", "street", "city", "state", "postal_code",
		"longitude", "latitude", "matched_address",
	}
	return readCSVRows(options.AddressPointsCSV, required, func(row map[string]string, line int) error {
		state, ok := normalizeState(row["state"])
		if !ok {
			return fmt.Errorf("address points line %d: invalid state %q", line, row["state"])
		}
		longitude, latitude, err := parseCoordinates(row["longitude"], row["latitude"])
		if err != nil {
			return fmt.Errorf("address points line %d: %w", line, err)
		}
		addressNumber := normalizeWords(row["address_number"])
		street := normalizeStreet(row["street"])
		city := normalizeWords(row["city"])
		postalCode := normalizePostalCode(row["postal_code"])
		recordID := strings.TrimSpace(row["source_record_id"])
		matchedAddress := strings.TrimSpace(row["matched_address"])
		if recordID == "" || addressNumber == "" || street == "" || city == "" || matchedAddress == "" {
			return fmt.Errorf("address points line %d: required field is empty", line)
		}
		_, err = tx.Exec(`
INSERT INTO address_points (
  address_number, street_name, city, state, postal_code, matched_address,
  longitude, latitude, source_name, source_record_id, source_vintage
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`,
			addressNumber, street, city, state, postalCode, matchedAddress,
			longitude, latitude, options.SourceName, recordID, options.SourceVintage,
		)
		if err != nil {
			return fmt.Errorf("address points line %d: %w", line, err)
		}
		return nil
	})
}

func importStreetRanges(tx *sql.Tx, options BuildOptions) error {
	required := []string{
		"source_record_id", "from_number", "to_number", "parity", "street", "city", "state",
		"postal_code", "from_longitude", "from_latitude", "to_longitude", "to_latitude",
	}
	return readCSVRows(options.StreetRangesCSV, required, func(row map[string]string, line int) error {
		fromNumber, err := strconv.Atoi(strings.TrimSpace(row["from_number"]))
		if err != nil {
			return fmt.Errorf("street ranges line %d: invalid from_number", line)
		}
		toNumber, err := strconv.Atoi(strings.TrimSpace(row["to_number"]))
		if err != nil {
			return fmt.Errorf("street ranges line %d: invalid to_number", line)
		}
		parity := strings.ToUpper(strings.TrimSpace(row["parity"]))
		if parity != "B" && parity != "E" && parity != "O" {
			return fmt.Errorf("street ranges line %d: parity must be B, E, or O", line)
		}
		state, ok := normalizeState(row["state"])
		if !ok {
			return fmt.Errorf("street ranges line %d: invalid state %q", line, row["state"])
		}
		fromLongitude, fromLatitude, err := parseCoordinates(row["from_longitude"], row["from_latitude"])
		if err != nil {
			return fmt.Errorf("street ranges line %d: %w", line, err)
		}
		toLongitude, toLatitude, err := parseCoordinates(row["to_longitude"], row["to_latitude"])
		if err != nil {
			return fmt.Errorf("street ranges line %d: %w", line, err)
		}
		recordID := strings.TrimSpace(row["source_record_id"])
		street := normalizeStreet(row["street"])
		city := normalizeWords(row["city"])
		if recordID == "" || street == "" || city == "" || fromNumber == toNumber {
			return fmt.Errorf("street ranges line %d: required field is empty or range has zero length", line)
		}
		_, err = tx.Exec(`
INSERT INTO street_ranges (
  from_number, to_number, parity, street_name, city, state, postal_code,
  from_longitude, from_latitude, to_longitude, to_latitude,
  source_name, source_record_id, source_vintage
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`,
			fromNumber, toNumber, parity, street, city, state, normalizePostalCode(row["postal_code"]),
			fromLongitude, fromLatitude, toLongitude, toLatitude,
			options.SourceName, recordID, options.SourceVintage,
		)
		if err != nil {
			return fmt.Errorf("street ranges line %d: %w", line, err)
		}
		return nil
	})
}

func readCSVRows(path string, required []string, visit func(map[string]string, int) error) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true
	header, err := reader.Read()
	if err != nil {
		return err
	}
	indexes := make(map[string]int, len(header))
	for index, name := range header {
		indexes[strings.TrimSpace(name)] = index
	}
	for _, name := range required {
		if _, ok := indexes[name]; !ok {
			return fmt.Errorf("%s: missing required column %q", path, name)
		}
	}

	for line := 2; ; line++ {
		record, err := reader.Read()
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return fmt.Errorf("%s line %d: %w", path, line, err)
		}
		row := make(map[string]string, len(indexes))
		for name, index := range indexes {
			if index < len(record) {
				row[name] = record[index]
			}
		}
		if err := visit(row, line); err != nil {
			return err
		}
	}
}

func parseCoordinates(longitudeValue, latitudeValue string) (float64, float64, error) {
	longitude, err := strconv.ParseFloat(strings.TrimSpace(longitudeValue), 64)
	if err != nil || longitude < -180 || longitude > 180 {
		return 0, 0, errors.New("longitude is outside the valid range")
	}
	latitude, err := strconv.ParseFloat(strings.TrimSpace(latitudeValue), 64)
	if err != nil || latitude < -90 || latitude > 90 {
		return 0, 0, errors.New("latitude is outside the valid range")
	}
	return longitude, latitude, nil
}

func normalizePostalCode(value string) string {
	trimmed := strings.TrimSpace(value)
	if match := zipPattern.FindStringSubmatch(trimmed); len(match) == 2 {
		return match[1]
	}
	return ""
}
