package snapshotbuild

import (
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

//go:embed sql/*.sql
var contracts embed.FS

func renderContract(request Request, runRoot string) (string, map[string]string, error) {
	sources := make(map[string]string, len(request.Sources))
	for _, source := range request.Sources {
		sources[source.Name] = filepath.Clean(source.Path)
	}
	normalized := map[string]string{}
	var scripts []string
	if request.SpatialExtensionPath != "" {
		scripts = append(scripts, "LOAD "+sqlLiteral(filepath.Clean(request.SpatialExtensionPath))+";")
	}

	switch request.Kind {
	case KindGeocoder:
		if source, ok := sources["address_points"]; ok {
			output := filepath.Join(runRoot, "address_points.csv")
			script, err := readContract("sql/geocoder_address_points.sql")
			if err != nil {
				return "", nil, err
			}
			script = strings.ReplaceAll(script, "{{SOURCE}}", sqlLiteral(source))
			script = strings.ReplaceAll(script, "{{OUTPUT}}", sqlLiteral(output))
			scripts = append(scripts, script)
			normalized["address_points"] = output
		}
		if source, ok := sources["street_ranges"]; ok {
			output := filepath.Join(runRoot, "street_ranges.csv")
			script, err := readContract("sql/geocoder_street_ranges.sql")
			if err != nil {
				return "", nil, err
			}
			script = strings.ReplaceAll(script, "{{SOURCE}}", sqlLiteral(source))
			script = strings.ReplaceAll(script, "{{OUTPUT}}", sqlLiteral(output))
			scripts = append(scripts, script)
			normalized["street_ranges"] = output
		}
	case KindBoundary:
		script, err := readContract("sql/boundary.sql")
		if err != nil {
			return "", nil, err
		}
		layerOutput := filepath.Join(runRoot, "layer_families.csv")
		featureOutput := filepath.Join(runRoot, "boundary_features.csv")
		refreshOutput := filepath.Join(runRoot, "refresh_status.csv")
		replacements := map[string]string{
			"{{LAYER_SOURCE}}":   sqlLiteral(sources["layer_families"]),
			"{{FEATURE_SOURCE}}": sqlLiteral(sources["boundary_features"]),
			"{{REFRESH_SOURCE}}": sqlLiteral(sources["refresh_status"]),
			"{{LAYER_OUTPUT}}":   sqlLiteral(layerOutput),
			"{{FEATURE_OUTPUT}}": sqlLiteral(featureOutput),
			"{{REFRESH_OUTPUT}}": sqlLiteral(refreshOutput),
		}
		geometryValid := "try_cast(trim(geometry_json) AS JSON) IS NOT NULL"
		if request.SpatialExtensionPath != "" {
			geometryValid = "(try_cast(trim(geometry_json) AS JSON) IS NOT NULL AND ST_IsValid(ST_GeomFromGeoJSON(trim(geometry_json))))"
		}
		replacements["{{GEOMETRY_VALID}}"] = geometryValid
		for token, value := range replacements {
			script = strings.ReplaceAll(script, token, value)
		}
		scripts = append(scripts, script)
		normalized["layer_families"] = layerOutput
		normalized["boundary_features"] = featureOutput
		normalized["refresh_status"] = refreshOutput
	default:
		return "", nil, fmt.Errorf("%w: unsupported kind %q", ErrInvalidRequest, request.Kind)
	}
	scripts = append(scripts, `
SELECT
  CAST(coalesce(sum(CASE WHEN metric = 'accepted' THEN value ELSE 0 END), 0) AS BIGINT) AS accepted,
  CAST(coalesce(sum(CASE WHEN metric = 'rejected' THEN value ELSE 0 END), 0) AS BIGINT) AS rejected,
  CAST(coalesce(sum(CASE WHEN metric = 'duplicate' THEN value ELSE 0 END), 0) AS BIGINT) AS duplicate,
  CAST(coalesce(sum(CASE WHEN metric = 'quarantined' THEN value ELSE 0 END), 0) AS BIGINT) AS quarantined
FROM bcm_audit;
`)
	return strings.Join(scripts, "\n"), normalized, nil
}

func readContract(name string) (string, error) {
	raw, err := contracts.ReadFile(name)
	if err != nil {
		return "", fmt.Errorf("read embedded DuckDB contract %s: %w", name, err)
	}
	return string(raw), nil
}

func contractDigest(kind Kind) string {
	names := []string{}
	switch kind {
	case KindGeocoder:
		names = []string{"sql/geocoder_address_points.sql", "sql/geocoder_street_ranges.sql"}
	case KindBoundary:
		names = []string{"sql/boundary.sql"}
	}
	sort.Strings(names)
	hash := sha256.New()
	hash.Write([]byte(PipelineIdentity))
	for _, name := range names {
		raw, err := contracts.ReadFile(name)
		if err != nil {
			panic(err)
		}
		hash.Write([]byte{0})
		hash.Write([]byte(name))
		hash.Write([]byte{0})
		hash.Write(raw)
	}
	return hex.EncodeToString(hash.Sum(nil))
}

func sqlLiteral(value string) string {
	return "'" + strings.ReplaceAll(filepath.ToSlash(value), "'", "''") + "'"
}

func validateContract(sqlText string) error {
	upper := strings.ToUpper(sqlText)
	for _, forbidden := range []string{
		"INSTALL ",
		"HTTP://",
		"HTTPS://",
		"S3://",
		"LOAD HTTPFS",
		"LOAD 'HTTPFS",
		"LOAD \"HTTPFS",
	} {
		if strings.Contains(upper, forbidden) {
			return fmt.Errorf("%w: contract contains forbidden token %q", ErrUnsafeContract, forbidden)
		}
	}
	return nil
}
