package engine

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"building-code-map/backend/internal/geocoder"
	"building-code-map/backend/internal/regulatory"
	"building-code-map/backend/internal/snapshot"
)

type BundleIdentity struct {
	SourceCommit            string `json:"source_commit"`
	EngineVersion           string `json:"engine_version"`
	BundleManifestDigest    string `json:"bundle_manifest_digest"`
	BoundarySnapshotDigest  string `json:"boundary_snapshot_digest"`
	RegulatoryCatalogDigest string `json:"regulatory_catalog_digest"`
	GeocoderSnapshotDigest  string `json:"geocoder_snapshot_digest,omitempty"`
}

type Config struct {
	Snapshot          snapshot.Snapshot
	RegulatoryCatalog regulatory.Catalog
	Geocoder          geocoder.Service
	Clock             Clock
	BundleIdentity    BundleIdentity
	Geography         GeographyResolver
}

type Engine interface {
	Resolve(context.Context, Query) (Result, error)
	Geocode(context.Context, string) (GeocodeResult, error)
	Lookup(context.Context, Point) (LookupResult, error)
	Readiness(context.Context) Readiness
}

type authorityEngine struct {
	snapshot  snapshot.Snapshot
	catalog   regulatory.Catalog
	geocoder  geocoder.Service
	clock     Clock
	identity  BundleIdentity
	geography GeographyResolver
}

type LookupResult struct {
	Point      Point                        `json:"point"`
	Geography  regulatory.GeographicContext `json:"geography"`
	Provenance Provenance                   `json:"provenance"`
}

type CapabilityReadiness struct {
	Status   string `json:"status"`
	Required bool   `json:"required"`
	Message  string `json:"message"`
}

type SnapshotReadiness struct {
	Status     string `json:"status"`
	SnapshotID string `json:"snapshot_id,omitempty"`
}

type Readiness struct {
	Status       string                         `json:"status"`
	Readiness    string                         `json:"readiness"`
	Capabilities map[string]CapabilityReadiness `json:"capabilities"`
	Snapshots    map[string]SnapshotReadiness   `json:"snapshots"`
}

func New(config Config) (Engine, error) {
	clock := config.Clock
	if clock == nil {
		clock = RealClock{}
	}
	geography := config.Geography
	if geography == nil {
		geography = NewSnapshotGeographyResolver(config.Snapshot, config.RegulatoryCatalog)
	}
	return &authorityEngine{
		snapshot:  config.Snapshot,
		catalog:   config.RegulatoryCatalog,
		geocoder:  config.Geocoder,
		clock:     clock,
		identity:  config.BundleIdentity,
		geography: geography,
	}, nil
}

func (engine *authorityEngine) Resolve(ctx context.Context, query Query) (Result, error) {
	normalized, err := NormalizeQuery(query)
	if err != nil {
		return Result{}, err
	}
	if engine.catalog.Len() == 0 {
		return Result{}, EngineError{Code: ErrorRegulatoryCatalogUnavailable, Message: "regulatory policy catalog is unavailable", Retryable: false}
	}

	result := Result{
		SchemaVersion: SchemaVersion,
		Query:         normalized,
		Location:      LocationResult{Address: normalized.Address},
		Provenance:    engine.provenance(),
		Diagnostics:   []Diagnostic{},
	}
	point := normalized.Point
	if strings.TrimSpace(normalized.Address) != "" {
		geocoded, geocodeErr := engine.Geocode(ctx, normalized.Address)
		if geocodeErr != nil {
			return Result{}, geocodeErr
		}
		result.Location.Geocode = &geocoded
		if geocoded.Selected == nil {
			return Result{}, EngineError{Code: ErrorAddressNotFound, Message: "address did not produce a selected coordinate", Retryable: false}
		}
		point = &Point{Longitude: geocoded.Selected.Longitude, Latitude: geocoded.Selected.Latitude}
	}
	if point == nil {
		return Result{}, EngineError{Code: ErrorInvalidQuery, Message: "point or address is required", Retryable: false}
	}
	result.Location.Point = &Point{Longitude: point.Longitude, Latitude: point.Latitude}

	geography, geographyErr := engine.geography.ResolveGeography(ctx, *point)
	if geographyErr != nil {
		return Result{}, mapGeographyError(geographyErr)
	}
	regulatoryResult, resolveErr := regulatory.ResolveAt(engine.catalog, regulatory.ResolutionRequest{
		Context:           &geography,
		CodeFamily:        normalized.CodeFamily,
		ProjectType:       normalized.ProjectType,
		ApplicabilityDate: normalized.ApplicabilityDate,
	}, engine.clock.Now())
	if resolveErr != nil {
		return Result{}, EngineError{Code: ErrorInternal, Message: resolveErr.Error(), Retryable: false}
	}
	result.Resolution, err = convertResolution(regulatoryResult)
	if err != nil {
		return Result{}, EngineError{Code: ErrorInternal, Message: err.Error(), Retryable: false}
	}
	SortDiagnostics(result.Diagnostics)
	return result, nil
}

func (engine *authorityEngine) Geocode(ctx context.Context, address string) (GeocodeResult, error) {
	if engine.geocoder == nil {
		return GeocodeResult{}, EngineError{Code: ErrorAddressNotFound, Message: "local geocoder is unavailable", Retryable: false}
	}
	result, err := engine.geocoder.Geocode(ctx, geocoder.Query{Address: address})
	if err != nil {
		if errors.Is(err, geocoder.ErrInvalidAddress) {
			return GeocodeResult{}, EngineError{Code: ErrorInvalidQuery, Message: err.Error(), Retryable: false}
		}
		return GeocodeResult{}, EngineError{Code: ErrorInternal, Message: "local geocoder query failed", Retryable: false}
	}
	converted := GeocodeResult{
		Query:      result.Query,
		Normalized: result.Normalized,
		Status:     string(result.Status),
		Candidates: make([]GeocodeCandidate, 0, len(result.Candidates)),
		Warnings:   append([]string(nil), result.Warnings...),
	}
	for _, candidate := range result.Candidates {
		convertedCandidate := GeocodeCandidate{
			MatchedAddress:       candidate.MatchedAddress,
			Longitude:            candidate.Longitude,
			Latitude:             candidate.Latitude,
			Precision:            string(candidate.Precision),
			Confidence:           candidate.Confidence,
			ScoreKind:            candidate.ScoreKind,
			ScoreFactors:         cloneScoreFactors(candidate.ScoreFactors),
			RankingPolicyVersion: candidate.RankingPolicyVersion,
			Source:               candidate.Source,
			SourceRecordID:       candidate.SourceRecordID,
			SourceVintage:        candidate.SourceVintage,
		}
		if candidate.Interpolation != nil {
			interpolation := candidate.Interpolation
			convertedCandidate.Interpolation = &InterpolationEvidence{
				SourceRangeID:             interpolation.SourceRangeID,
				RequestedHouseNumber:      interpolation.RequestedHouseNumber,
				FromNumber:                interpolation.FromNumber,
				ToNumber:                  interpolation.ToNumber,
				RangeDirection:            interpolation.RangeDirection,
				Parity:                    interpolation.Parity,
				Side:                      interpolation.Side,
				FromCoordinate:            Coordinate{Longitude: interpolation.FromCoordinate.Longitude, Latitude: interpolation.FromCoordinate.Latitude},
				ToCoordinate:              Coordinate{Longitude: interpolation.ToCoordinate.Longitude, Latitude: interpolation.ToCoordinate.Latitude},
				Fraction:                  interpolation.Fraction,
				DerivedCoordinate:         Coordinate{Longitude: interpolation.DerivedCoordinate.Longitude, Latitude: interpolation.DerivedCoordinate.Latitude},
				CoordinateReferenceSystem: interpolation.CoordinateReferenceSystem,
				TransformationIdentity:    interpolation.TransformationIdentity,
				MethodVersion:             interpolation.MethodVersion,
				PositionalQuality:         interpolation.PositionalQuality,
			}
		}
		converted.Candidates = append(converted.Candidates, convertedCandidate)
	}
	if result.Selected != nil {
		selected := cloneGeocodeCandidate(*result.Selected)
		converted.Selected = &selected
	}
	if result.Status == geocoder.StatusAmbiguous {
		return converted, EngineError{Code: ErrorAddressAmbiguous, Message: "address matched multiple equally suitable candidates", Details: map[string]any{"candidate_count": len(result.Candidates)}, Retryable: false}
	}
	if result.Status == geocoder.StatusNotFound {
		return converted, EngineError{Code: ErrorAddressNotFound, Message: "address was not found in the local geocoder", Retryable: false}
	}
	return converted, nil
}

func (engine *authorityEngine) Lookup(ctx context.Context, point Point) (LookupResult, error) {
	normalized, err := NormalizeQuery(Query{Point: &point, ApplicabilityDate: engine.clock.Now().UTC().Format("2006-01-02")})
	if err != nil {
		return LookupResult{}, err
	}
	geography, err := engine.geography.ResolveGeography(ctx, *normalized.Point)
	if err != nil {
		return LookupResult{}, mapGeographyError(err)
	}
	return LookupResult{Point: *normalized.Point, Geography: geography, Provenance: engine.provenance()}, nil
}

func (engine *authorityEngine) Readiness(context.Context) Readiness {
	boundaryAvailable := len(engine.snapshot.LayerFamilies) > 0 && len(engine.snapshot.BoundaryFeatures) > 0
	geocoderAvailable := engine.geocoder != nil
	regulatoryAvailable := engine.catalog.Len() > 0
	capabilities := map[string]CapabilityReadiness{
		"boundary_resolution":   capabilityReadiness(boundaryAvailable, true, "A validated local boundary snapshot is loaded.", "No usable local boundary snapshot is loaded."),
		"coordinate_resolution": capabilityReadiness(boundaryAvailable && regulatoryAvailable, false, "Coordinate input can proceed through boundary and regulatory resolution.", "Coordinate resolution requires both boundary data and regulatory profiles."),
		"address_geocoding":     capabilityReadiness(geocoderAvailable, false, "A local geocoder snapshot is loaded.", "No usable local geocoder snapshot is loaded."),
		"regulatory_resolution": capabilityReadiness(regulatoryAvailable, false, "One or more validated regulatory profiles are loaded.", "No validated regulatory profiles are loaded."),
		"address_lookup":        capabilityReadiness(boundaryAvailable && geocoderAvailable && regulatoryAvailable, false, "Address lookup can preserve geocoding evidence through regulatory resolution.", "Address lookup requires boundary data, a local geocoder, and regulatory profiles."),
	}
	snapshots := map[string]SnapshotReadiness{
		"boundary": snapshotReadiness(boundaryAvailable, engine.identity.BoundarySnapshotDigest),
		"geocoder": snapshotReadiness(geocoderAvailable, engine.identity.GeocoderSnapshotDigest),
	}
	readiness := "ready"
	status := "ok"
	if !boundaryAvailable {
		readiness, status = "not_ready", "not_ready"
	} else if !geocoderAvailable || !regulatoryAvailable {
		readiness = "degraded"
	}
	return Readiness{Status: status, Readiness: readiness, Capabilities: capabilities, Snapshots: snapshots}
}

func (engine *authorityEngine) provenance() Provenance {
	return Provenance{
		SourceCommit:            engine.identity.SourceCommit,
		EngineVersion:           engine.identity.EngineVersion,
		BundleManifestDigest:    engine.identity.BundleManifestDigest,
		BoundarySnapshotDigest:  engine.identity.BoundarySnapshotDigest,
		RegulatoryCatalogDigest: engine.identity.RegulatoryCatalogDigest,
		GeocoderSnapshotDigest:  engine.identity.GeocoderSnapshotDigest,
	}
}

func mapGeographyError(err error) error {
	var ambiguity *BoundaryAmbiguityError
	if errors.As(err, &ambiguity) {
		return EngineError{Code: ErrorBoundaryAmbiguous, Message: ambiguity.Error(), Details: map[string]any{"layer_family": ambiguity.LayerFamily, "observations": ambiguity.Matches}, Retryable: false}
	}
	if strings.Contains(err.Error(), "outside valid") {
		return EngineError{Code: ErrorInvalidCoordinates, Message: err.Error(), Retryable: false}
	}
	if strings.Contains(err.Error(), "supported state") {
		return EngineError{Code: ErrorOutsideSupportedCoverage, Message: err.Error(), Retryable: false}
	}
	return EngineError{Code: ErrorInternal, Message: err.Error(), Retryable: false}
}

func convertResolution(value regulatory.ResolutionResult) (Resolution, error) {
	raw, err := json.Marshal(value)
	if err != nil {
		return Resolution{}, err
	}
	var result Resolution
	if err := json.Unmarshal(raw, &result); err != nil {
		return Resolution{}, err
	}
	return result, nil
}

func cloneScoreFactors(factors map[string]float64) map[string]float64 {
	if factors == nil {
		return nil
	}
	clone := make(map[string]float64, len(factors))
	for key, value := range factors {
		clone[key] = value
	}
	return clone
}

func cloneGeocodeCandidate(candidate geocoder.Candidate) GeocodeCandidate {
	result := GeocodeCandidate{
		MatchedAddress:       candidate.MatchedAddress,
		Longitude:            candidate.Longitude,
		Latitude:             candidate.Latitude,
		Precision:            string(candidate.Precision),
		Confidence:           candidate.Confidence,
		ScoreKind:            candidate.ScoreKind,
		ScoreFactors:         cloneScoreFactors(candidate.ScoreFactors),
		RankingPolicyVersion: candidate.RankingPolicyVersion,
		Source:               candidate.Source,
		SourceRecordID:       candidate.SourceRecordID,
		SourceVintage:        candidate.SourceVintage,
	}
	if candidate.Interpolation != nil {
		interpolation := candidate.Interpolation
		result.Interpolation = &InterpolationEvidence{
			SourceRangeID:             interpolation.SourceRangeID,
			RequestedHouseNumber:      interpolation.RequestedHouseNumber,
			FromNumber:                interpolation.FromNumber,
			ToNumber:                  interpolation.ToNumber,
			RangeDirection:            interpolation.RangeDirection,
			Parity:                    interpolation.Parity,
			Side:                      interpolation.Side,
			FromCoordinate:            Coordinate{Longitude: interpolation.FromCoordinate.Longitude, Latitude: interpolation.FromCoordinate.Latitude},
			ToCoordinate:              Coordinate{Longitude: interpolation.ToCoordinate.Longitude, Latitude: interpolation.ToCoordinate.Latitude},
			Fraction:                  interpolation.Fraction,
			DerivedCoordinate:         Coordinate{Longitude: interpolation.DerivedCoordinate.Longitude, Latitude: interpolation.DerivedCoordinate.Latitude},
			CoordinateReferenceSystem: interpolation.CoordinateReferenceSystem,
			TransformationIdentity:    interpolation.TransformationIdentity,
			MethodVersion:             interpolation.MethodVersion,
			PositionalQuality:         interpolation.PositionalQuality,
		}
	}
	return result
}

func capabilityReadiness(available, required bool, availableMessage, unavailableMessage string) CapabilityReadiness {
	if available {
		return CapabilityReadiness{Status: "available", Required: required, Message: availableMessage}
	}
	return CapabilityReadiness{Status: "unavailable", Required: required, Message: unavailableMessage}
}

func snapshotReadiness(available bool, snapshotID string) SnapshotReadiness {
	if available && snapshotID != "" {
		return SnapshotReadiness{Status: "verified", SnapshotID: snapshotID}
	}
	if available {
		return SnapshotReadiness{Status: "unidentified"}
	}
	return SnapshotReadiness{Status: "unavailable"}
}

var _ Engine = (*authorityEngine)(nil)
