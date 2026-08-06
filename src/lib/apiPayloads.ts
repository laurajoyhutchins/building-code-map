import type {
  AttributeValue,
  BoundaryFeatureRecord,
  CapabilityReadiness,
  FeatureRecord,
  GeoJSONGeometry,
  GeocodeCandidate,
  GeocodeResult,
  LayerFamilyDefinition,
  LayerFamilyKey,
  LookupResult,
  ReadinessResult,
  RefreshStatus,
  ResolutionBoundaryMatch,
  ResolutionResult,
  SnapshotReadiness,
  VerificationRecord,
} from "../types";
import {
  arrayValue,
  booleanValue,
  enumValue,
  finiteNumber,
  httpUrl,
  isoDate,
  isoDateTime,
  latitude,
  longitude,
  nonEmptyString,
  optionalArray,
  optionalString,
  record,
  stringArray,
  stringValue,
} from "./runtimeDecode";

const layerKeys = [
  "states",
  "counties",
  "municipalities",
  "special_areas",
  "tribal_areas",
  "neris_jurisdictions",
] as const;
const resolutionStatuses = [
  "resolved",
  "partially_resolved",
  "local_record_required",
  "ambiguous",
  "conflicting",
  "insufficient_evidence",
] as const;

export function decodeHealth(value: unknown): { status: string } {
  const raw = record(value, "health");
  return { status: nonEmptyString(raw.status, "health.status") };
}

export function decodeReadiness(value: unknown): ReadinessResult {
  const raw = record(value, "ready");
  return {
    status: enumValue(raw.status, ["ok", "not_ready"] as const, "ready.status"),
    readiness: enumValue(
      raw.readiness,
      ["ready", "degraded", "not_ready"] as const,
      "ready.readiness",
    ),
    capabilities: decodeNamedRecord(raw.capabilities, "ready.capabilities", decodeCapability),
    snapshots: decodeNamedRecord(raw.snapshots, "ready.snapshots", decodeSnapshot),
  };
}

function decodeCapability(value: unknown, path: string): CapabilityReadiness {
  const raw = record(value, path);
  return {
    status: enumValue(raw.status, ["available", "unavailable"] as const, `${path}.status`),
    required: booleanValue(raw.required, `${path}.required`),
    message: nonEmptyString(raw.message, `${path}.message`),
  };
}

function decodeSnapshot(value: unknown, path: string): SnapshotReadiness {
  const raw = record(value, path);
  const status = enumValue(
    raw.status,
    ["verified", "unidentified", "unavailable"] as const,
    `${path}.status`,
  );
  const snapshotId = optionalString(raw.snapshot_id, `${path}.snapshot_id`);
  if (status === "verified" && !snapshotId) {
    throw new Error(`${path}.snapshot_id: required for a verified snapshot`);
  }
  return { status, snapshotId };
}

export function decodeLayers(value: unknown): LayerFamilyDefinition[] {
  return arrayValue(value, "layers", (item, path) => {
    const raw = record(item, path);
    return {
      key: enumValue(raw.key, layerKeys, `${path}.key`) as LayerFamilyKey,
      label: nonEmptyString(raw.label, `${path}.label`),
      martinLayerId: nonEmptyString(raw.martin_layer_id, `${path}.martin_layer_id`),
      description: stringValue(raw.description, `${path}.description`),
      defaultEnabled: booleanValue(raw.default_enabled, `${path}.default_enabled`),
    };
  });
}

export function decodeBoundaryFeatures(value: unknown): BoundaryFeatureRecord[] {
  return arrayValue(
    value,
    "boundaries",
    (item, path) => decodeFeature(item, path, true) as BoundaryFeatureRecord,
  );
}

export function decodeFeatureRecord(value: unknown): FeatureRecord {
  return decodeFeature(value, "feature", false);
}

function decodeFeature(value: unknown, path: string, geometryRequired: boolean): FeatureRecord {
  const raw = record(value, path);
  const result: FeatureRecord = {
    layerFamily: enumValue(raw.layer_family, layerKeys, `${path}.layer_family`) as LayerFamilyKey,
    featureId: nonEmptyString(raw.feature_id, `${path}.feature_id`),
    title: nonEmptyString(raw.title, `${path}.title`),
    subtitle: stringValue(raw.subtitle, `${path}.subtitle`),
    sourceId: nonEmptyString(raw.source_id, `${path}.source_id`),
    geometryLabel: nonEmptyString(raw.geometry_label, `${path}.geometry_label`),
    geometrySource: optionalString(raw.geometry_source, `${path}.geometry_source`),
    lastSyncedAt: isoDateTime(raw.last_synced_at, `${path}.last_synced_at`),
    attributes: decodeAttributes(raw.attributes, `${path}.attributes`),
  };
  if (raw.geometry !== undefined) {
    result.geometry = decodeGeometry(raw.geometry, `${path}.geometry`);
  } else if (geometryRequired) {
    throw new Error(`${path}.geometry: required`);
  }
  return result;
}

function decodeAttributes(value: unknown, path: string): Record<string, AttributeValue> {
  const raw = record(value, path);
  const result: Record<string, AttributeValue> = {};
  for (const [key, item] of Object.entries(raw)) {
    if (
      item !== null &&
      typeof item !== "string" &&
      typeof item !== "number" &&
      typeof item !== "boolean"
    ) {
      throw new Error(`${path}.${key}: expected a scalar attribute`);
    }
    result[key] = item as AttributeValue;
  }
  return result;
}

function decodeGeometry(value: unknown, path: string): GeoJSONGeometry {
  const raw = record(value, path);
  const type = enumValue(raw.type, ["Polygon", "MultiPolygon"] as const, `${path}.type`);
  if (!Array.isArray(raw.coordinates)) {
    throw new Error(`${path}.coordinates: expected an array`);
  }
  validateCoordinateTree(raw.coordinates, `${path}.coordinates`);
  return { type, coordinates: raw.coordinates } as GeoJSONGeometry;
}

function validateCoordinateTree(value: unknown, path: string): void {
  if (!Array.isArray(value) || value.length === 0) {
    throw new Error(`${path}: expected a non-empty coordinate array`);
  }
  if (value.length >= 2 && typeof value[0] === "number" && typeof value[1] === "number") {
    longitude(value[0], `${path}[0]`);
    latitude(value[1], `${path}[1]`);
    return;
  }
  value.forEach((item, index) => validateCoordinateTree(item, `${path}[${index}]`));
}

export function decodeRefreshStatus(value: unknown): RefreshStatus {
  const raw = record(value, "refresh");
  return {
    status: enumValue(raw.status, ["ok", "warning", "error"] as const, "refresh.status"),
    latestSuccessfulRefresh:
      raw.latest_successful_refresh === null
        ? null
        : isoDateTime(raw.latest_successful_refresh, "refresh.latest_successful_refresh"),
    latestAttempt: isoDateTime(raw.latest_attempt, "refresh.latest_attempt"),
    nextScheduledRefresh: isoDateTime(raw.next_scheduled_refresh, "refresh.next_scheduled_refresh"),
    message: stringValue(raw.message, "refresh.message"),
  };
}

export function decodeGeocodeResult(value: unknown, path = "geocode"): GeocodeResult {
  const raw = record(value, path);
  const selected =
    raw.selected === undefined
      ? undefined
      : decodeGeocodeCandidate(raw.selected, `${path}.selected`);
  return {
    query: nonEmptyString(raw.query, `${path}.query`),
    normalized: optionalString(raw.normalized, `${path}.normalized`),
    status: enumValue(raw.status, ["matched", "ambiguous", "not_found"] as const, `${path}.status`),
    selected,
    candidates: optionalArray(raw.candidates, `${path}.candidates`, decodeGeocodeCandidate),
    warnings: raw.warnings === undefined ? [] : stringArray(raw.warnings, `${path}.warnings`),
  };
}

function decodeGeocodeCandidate(value: unknown, path: string): GeocodeCandidate {
  const raw = record(value, path);
  return {
    matchedAddress: nonEmptyString(raw.matched_address, `${path}.matched_address`),
    longitude: longitude(raw.longitude, `${path}.longitude`),
    latitude: latitude(raw.latitude, `${path}.latitude`),
    precision: enumValue(
      raw.precision,
      ["address_point", "interpolated"] as const,
      `${path}.precision`,
    ),
    confidence: finiteNumber(raw.confidence, `${path}.confidence`),
    scoreKind: enumValue(raw.score_kind, ["deterministic_quality"] as const, `${path}.score_kind`),
    scoreFactors: decodeScoreFactors(raw.score_factors, `${path}.score_factors`),
    rankingPolicyVersion: nonEmptyString(
      raw.ranking_policy_version,
      `${path}.ranking_policy_version`,
    ),
    source: nonEmptyString(raw.source, `${path}.source`),
    sourceRecordId: nonEmptyString(raw.source_record_id, `${path}.source_record_id`),
    sourceVintage: nonEmptyString(raw.source_vintage, `${path}.source_vintage`),
    interpolation:
      raw.interpolation === undefined
        ? undefined
        : decodeInterpolationProvenance(raw.interpolation, `${path}.interpolation`),
  };
}

function decodeScoreFactors(value: unknown, path: string): Record<string, number> {
  const raw = record(value, path);
  const result: Record<string, number> = {};
  for (const [key, item] of Object.entries(raw)) {
    result[key] = finiteNumber(item, `${path}.${key}`);
  }
  return result;
}

function decodeInterpolationProvenance(
  value: unknown,
  path: string,
): NonNullable<GeocodeCandidate["interpolation"]> {
  const raw = record(value, path);
  return {
    sourceRangeId: nonEmptyString(raw.source_range_id, `${path}.source_range_id`),
    requestedHouseNumber: finiteNumber(
      raw.requested_house_number,
      `${path}.requested_house_number`,
    ),
    fromNumber: finiteNumber(raw.from_number, `${path}.from_number`),
    toNumber: finiteNumber(raw.to_number, `${path}.to_number`),
    rangeDirection: enumValue(
      raw.range_direction,
      ["ascending", "descending"] as const,
      `${path}.range_direction`,
    ),
    parity: nonEmptyString(raw.parity, `${path}.parity`),
    side: optionalString(raw.side, `${path}.side`),
    fromCoordinate: decodeCoordinateEvidence(raw.from_coordinate, `${path}.from_coordinate`),
    toCoordinate: decodeCoordinateEvidence(raw.to_coordinate, `${path}.to_coordinate`),
    fraction: finiteNumber(raw.fraction, `${path}.fraction`),
    derivedCoordinate: decodeCoordinateEvidence(
      raw.derived_coordinate,
      `${path}.derived_coordinate`,
    ),
    coordinateReferenceSystem: nonEmptyString(
      raw.coordinate_reference_system,
      `${path}.coordinate_reference_system`,
    ),
    transformationIdentity: nonEmptyString(
      raw.transformation_identity,
      `${path}.transformation_identity`,
    ),
    methodVersion: nonEmptyString(raw.method_version, `${path}.method_version`),
    positionalQuality: nonEmptyString(raw.positional_quality, `${path}.positional_quality`),
  };
}

function decodeCoordinateEvidence(
  value: unknown,
  path: string,
): { longitude: number; latitude: number } {
  const raw = record(value, path);
  return {
    longitude: longitude(raw.longitude, `${path}.longitude`),
    latitude: latitude(raw.latitude, `${path}.latitude`),
  };
}

export function decodeLookupResult(value: unknown): LookupResult {
  const raw = record(value, "lookup");
  return {
    geocode: decodeGeocodeResult(raw.geocode, "lookup.geocode"),
    resolution: decodeResolutionResult(raw.resolution, "lookup.resolution"),
  };
}

export function decodeResolutionResult(value: unknown, path = "resolution"): ResolutionResult {
  const raw = record(value, path);
  const geography = record(raw.geography, `${path}.geography`);
  const boundary = (item: unknown, itemPath: string): ResolutionBoundaryMatch => {
    const match = record(item, itemPath);
    return {
      layerFamily: nonEmptyString(match.layer_family, `${itemPath}.layer_family`),
      featureId: nonEmptyString(match.feature_id, `${itemPath}.feature_id`),
      name: nonEmptyString(match.name, `${itemPath}.name`),
      sourceId: optionalString(match.source_id, `${itemPath}.source_id`),
    };
  };
  return {
    schemaVersion: nonEmptyString(raw.schema_version, `${path}.schema_version`),
    generatedAt: isoDateTime(raw.generated_at, `${path}.generated_at`),
    profileId: optionalString(raw.profile_id, `${path}.profile_id`),
    profileLastVerified:
      raw.profile_last_verified === undefined
        ? undefined
        : isoDate(raw.profile_last_verified, `${path}.profile_last_verified`),
    geography: {
      stateId: optionalString(geography.state_id, `${path}.geography.state_id`),
      stateFips: optionalString(geography.state_fips, `${path}.geography.state_fips`),
      stateName: optionalString(geography.state_name, `${path}.geography.state_name`),
      county:
        geography.county === undefined
          ? undefined
          : boundary(geography.county, `${path}.geography.county`),
      municipality:
        geography.municipality === undefined
          ? undefined
          : boundary(geography.municipality, `${path}.geography.municipality`),
      incorporated: booleanValue(geography.incorporated, `${path}.geography.incorporated`),
      specialAreas: optionalArray(
        geography.special_areas,
        `${path}.geography.special_areas`,
        boundary,
      ),
      tribalAreas: optionalArray(
        geography.tribal_areas,
        `${path}.geography.tribal_areas`,
        boundary,
      ),
      fireJurisdictions: optionalArray(
        geography.fire_jurisdictions,
        `${path}.geography.fire_jurisdictions`,
        boundary,
      ),
    },
    codeFamily: optionalString(raw.code_family, `${path}.code_family`),
    projectType: optionalString(raw.project_type, `${path}.project_type`),
    applicabilityDate:
      raw.applicability_date === undefined
        ? undefined
        : isoDate(raw.applicability_date, `${path}.applicability_date`),
    status: enumValue(raw.status, resolutionStatuses, `${path}.status`),
    policyBasis:
      raw.policy_basis === undefined
        ? undefined
        : decodePolicyBasis(raw.policy_basis, `${path}.policy_basis`),
    authorityCandidates: optionalArray(
      raw.authority_candidates,
      `${path}.authority_candidates`,
      decodeAuthority,
    ),
    authorityPath: optionalArray(raw.authority_path, `${path}.authority_path`, decodeRelationship),
    adoptions: optionalArray(raw.adoptions, `${path}.adoptions`, decodeAdoption),
    applicableRules: optionalArray(raw.applicable_rules, `${path}.applicable_rules`, decodeRule),
    supportingClaims: optionalArray(
      raw.supporting_claims,
      `${path}.supporting_claims`,
      decodeClaim,
    ),
    requiredLocalRecords:
      raw.required_local_records === undefined
        ? []
        : stringArray(raw.required_local_records, `${path}.required_local_records`),
    warnings: raw.warnings === undefined ? [] : stringArray(raw.warnings, `${path}.warnings`),
    evidence: optionalArray(raw.evidence, `${path}.evidence`, decodeSource),
  };
}

function decodeVerification(value: unknown, path: string): VerificationRecord {
  const raw = record(value, path);
  return {
    status: nonEmptyString(raw.status, `${path}.status`),
    confidence:
      raw.confidence === undefined ? undefined : finiteNumber(raw.confidence, `${path}.confidence`),
    notes: optionalString(raw.notes, `${path}.notes`),
  };
}

function decodePolicyBasis(value: unknown, path: string): ResolutionResult["policyBasis"] {
  const raw = record(value, path);
  return {
    status: enumValue(raw.status, resolutionStatuses, `${path}.status`),
    requiredLocalRecords:
      raw.required_local_records === undefined
        ? []
        : stringArray(raw.required_local_records, `${path}.required_local_records`),
    warnings: raw.warnings === undefined ? [] : stringArray(raw.warnings, `${path}.warnings`),
    sourceIds:
      raw.source_ids === undefined ? [] : stringArray(raw.source_ids, `${path}.source_ids`),
    verification: decodeVerification(raw.verification, `${path}.verification`),
  };
}

function decodeAuthority(
  value: unknown,
  path: string,
): ResolutionResult["authorityCandidates"][number] {
  const raw = record(value, path);
  return {
    kind: nonEmptyString(raw.kind, `${path}.kind`),
    authorityId: optionalString(raw.authority_id, `${path}.authority_id`),
    name: nonEmptyString(raw.name, `${path}.name`),
    roles: raw.roles === undefined ? [] : stringArray(raw.roles, `${path}.roles`),
    sourceIds:
      raw.source_ids === undefined ? [] : stringArray(raw.source_ids, `${path}.source_ids`),
    verification: decodeVerification(raw.verification, `${path}.verification`),
  };
}

function decodeRelationship(
  value: unknown,
  path: string,
): ResolutionResult["authorityPath"][number] {
  const raw = record(value, path);
  return {
    id: nonEmptyString(raw.id, `${path}.id`),
    fromId: nonEmptyString(raw.from_id, `${path}.from_id`),
    relationship: nonEmptyString(raw.relationship, `${path}.relationship`),
    to: nonEmptyString(raw.to, `${path}.to`),
    scope: raw.scope === undefined ? [] : stringArray(raw.scope, `${path}.scope`),
    summary: optionalString(raw.summary, `${path}.summary`),
    sourceIds:
      raw.source_ids === undefined ? [] : stringArray(raw.source_ids, `${path}.source_ids`),
    verification: decodeVerification(raw.verification, `${path}.verification`),
  };
}

function decodeAdoption(value: unknown, path: string): ResolutionResult["adoptions"][number] {
  const raw = record(value, path);
  const datesRaw = raw.dates === undefined ? {} : record(raw.dates, `${path}.dates`);
  const dates: Record<string, string> = {};
  for (const [key, date] of Object.entries(datesRaw))
    dates[key] = isoDate(date, `${path}.dates.${key}`);
  return {
    id: nonEmptyString(raw.id, `${path}.id`),
    codeFamily: nonEmptyString(raw.code_family, `${path}.code_family`),
    status: nonEmptyString(raw.status, `${path}.status`),
    stateCodeName: nonEmptyString(raw.state_code_name, `${path}.state_code_name`),
    enforcementModel: nonEmptyString(raw.enforcement_model, `${path}.enforcement_model`),
    dates,
    sourceIds:
      raw.source_ids === undefined ? [] : stringArray(raw.source_ids, `${path}.source_ids`),
    verification: decodeVerification(raw.verification, `${path}.verification`),
  };
}

function decodeRule(value: unknown, path: string): ResolutionResult["applicableRules"][number] {
  const raw = record(value, path);
  return {
    id: nonEmptyString(raw.id, `${path}.id`),
    kind: enumValue(
      raw.kind,
      ["applicability", "date", "amendment", "enforcement"] as const,
      `${path}.kind`,
    ),
    codeFamily: optionalString(raw.code_family, `${path}.code_family`),
    summary: nonEmptyString(raw.summary, `${path}.summary`),
    sourceIds:
      raw.source_ids === undefined ? [] : stringArray(raw.source_ids, `${path}.source_ids`),
    verification: decodeVerification(raw.verification, `${path}.verification`),
  };
}

function decodeClaim(value: unknown, path: string): ResolutionResult["supportingClaims"][number] {
  const raw = record(value, path);
  return {
    id: nonEmptyString(raw.id, `${path}.id`),
    subjectId: nonEmptyString(raw.subject_id, `${path}.subject_id`),
    field: nonEmptyString(raw.field, `${path}.field`),
    status: nonEmptyString(raw.status, `${path}.status`),
    value: raw.value,
    conflictGroup: optionalString(raw.conflict_group, `${path}.conflict_group`),
    sourceIds:
      raw.source_ids === undefined ? [] : stringArray(raw.source_ids, `${path}.source_ids`),
    verification: decodeVerification(raw.verification, `${path}.verification`),
  };
}

function decodeSource(value: unknown, path: string): ResolutionResult["evidence"][number] {
  const raw = record(value, path);
  return {
    id: nonEmptyString(raw.id, `${path}.id`),
    title: nonEmptyString(raw.title, `${path}.title`),
    url: httpUrl(raw.url, `${path}.url`),
    kind: nonEmptyString(raw.kind, `${path}.kind`),
    accessedAt: isoDate(raw.accessed_at, `${path}.accessed_at`),
    lastCheckedAt:
      raw.last_checked_at === undefined
        ? undefined
        : isoDate(raw.last_checked_at, `${path}.last_checked_at`),
    availability:
      raw.availability === undefined
        ? undefined
        : enumValue(
            raw.availability,
            ["available", "unavailable", "moved", "unknown"] as const,
            `${path}.availability`,
          ),
    caveat: optionalString(raw.caveat, `${path}.caveat`),
  };
}

function decodeNamedRecord<T>(
  value: unknown,
  path: string,
  decode: (item: unknown, itemPath: string) => T,
): Record<string, T> {
  const raw = record(value, path);
  const result: Record<string, T> = {};
  for (const [key, item] of Object.entries(raw)) result[key] = decode(item, `${path}.${key}`);
  return result;
}
