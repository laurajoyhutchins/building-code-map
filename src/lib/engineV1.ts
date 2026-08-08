import type { GeocodeResult, ResolutionResult } from "../types";
import { decodeGeocodeResult, decodeResolutionResult } from "./apiPayloads";
import {
  isoDate,
  latitude,
  longitude,
  nonEmptyString,
  optionalArray,
  optionalString,
  record,
  stringArray,
} from "./runtimeDecode";

export interface EnginePoint {
  longitude: number;
  latitude: number;
}

export interface EngineProvenance {
  sourceCommit: string;
  engineVersion: string;
  bundleManifestDigest: string;
  boundarySnapshotDigest: string;
  regulatoryCatalogDigest: string;
  geocoderSnapshotDigest?: string;
}

export interface EngineDiagnostic {
  severity: string;
  code: string;
  message: string;
  path?: string;
}

export interface EngineResult {
  schemaVersion: string;
  query: {
    point?: EnginePoint;
    address?: string;
    codeFamily?: string;
    projectType?: string;
    applicabilityDate: string;
    include: string[];
  };
  location: {
    point?: EnginePoint;
    address?: string;
    geocode?: GeocodeResult;
  };
  resolution: ResolutionResult;
  provenance: EngineProvenance;
  diagnostics: EngineDiagnostic[];
  unknownFields?: Record<string, unknown>;
}

function decodePoint(value: unknown, path: string): EnginePoint {
  const raw = record(value, path);
  return {
    longitude: longitude(raw.longitude, `${path}.longitude`),
    latitude: latitude(raw.latitude, `${path}.latitude`),
  };
}

export function decodeEngineProvenance(
  value: unknown,
  path = "engine provenance",
): EngineProvenance {
  const raw = record(value, path);
  return {
    sourceCommit: nonEmptyString(raw.source_commit, `${path}.source_commit`),
    engineVersion: nonEmptyString(raw.engine_version, `${path}.engine_version`),
    bundleManifestDigest: nonEmptyString(
      raw.bundle_manifest_digest,
      `${path}.bundle_manifest_digest`,
    ),
    boundarySnapshotDigest: nonEmptyString(
      raw.boundary_snapshot_digest,
      `${path}.boundary_snapshot_digest`,
    ),
    regulatoryCatalogDigest: nonEmptyString(
      raw.regulatory_catalog_digest,
      `${path}.regulatory_catalog_digest`,
    ),
    geocoderSnapshotDigest: optionalString(
      raw.geocoder_snapshot_digest,
      `${path}.geocoder_snapshot_digest`,
    ),
  };
}

export function decodeEngineResult(value: unknown, path = "engine result"): EngineResult {
  const raw = record(value, path);
  const query = record(raw.query, `${path}.query`);
  const location = record(raw.location, `${path}.location`);
  const diagnostics = optionalArray(
    raw.diagnostics,
    `${path}.diagnostics`,
    (item, diagnosticPath) => {
      const diagnostic = record(item, diagnosticPath);
      return {
        severity: nonEmptyString(diagnostic.severity, `${diagnosticPath}.severity`),
        code: nonEmptyString(diagnostic.code, `${diagnosticPath}.code`),
        message: nonEmptyString(diagnostic.message, `${diagnosticPath}.message`),
        path: optionalString(diagnostic.path, `${diagnosticPath}.path`),
      };
    },
  );
  const knownKeys = new Set([
    "schema_version",
    "query",
    "location",
    "resolution",
    "provenance",
    "diagnostics",
  ]);
  const unknownFields = Object.fromEntries(
    Object.entries(raw).filter(([key]) => !knownKeys.has(key)),
  );

  return {
    schemaVersion: nonEmptyString(raw.schema_version, `${path}.schema_version`),
    query: {
      point:
        query.point === undefined ? undefined : decodePoint(query.point, `${path}.query.point`),
      address: optionalString(query.address, `${path}.query.address`),
      codeFamily: optionalString(query.code_family, `${path}.query.code_family`),
      projectType: optionalString(query.project_type, `${path}.query.project_type`),
      applicabilityDate: isoDate(query.applicability_date, `${path}.query.applicability_date`),
      include:
        query.include === undefined ? [] : stringArray(query.include, `${path}.query.include`),
    },
    location: {
      point:
        location.point === undefined
          ? undefined
          : decodePoint(location.point, `${path}.location.point`),
      address: optionalString(location.address, `${path}.location.address`),
      geocode:
        location.geocode === undefined
          ? undefined
          : decodeGeocodeResult(location.geocode, `${path}.location.geocode`),
    },
    resolution: decodeResolutionResult(raw.resolution, `${path}.resolution`),
    provenance: decodeEngineProvenance(raw.provenance, `${path}.provenance`),
    diagnostics,
    ...(Object.keys(unknownFields).length > 0 ? { unknownFields } : {}),
  };
}
