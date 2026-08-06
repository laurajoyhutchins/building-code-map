import type {
  AttributeValue,
  BoundaryMapRecord,
  FeatureRecord,
  GeoJSONGeometry,
  LayerFamilyKey,
} from "../types";
import {
  arrayValue,
  enumValue,
  isoDateTime,
  latitude,
  longitude,
  nonEmptyString,
  optionalString,
  record,
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

export function decodeBoundaryMapRecords(value: unknown): BoundaryMapRecord[] {
  return arrayValue(value, "boundaries", decodeBoundaryMapRecord);
}

function decodeBoundaryMapRecord(value: unknown, path: string): BoundaryMapRecord {
  const raw = record(value, path);
  if ("attributes" in raw) {
    throw new Error(`${path}.attributes: detail-only fields are forbidden in the map payload`);
  }
  return {
    ...decodeSummary(raw, path),
    geometry: decodeGeometry(raw.geometry, `${path}.geometry`),
  };
}

export function decodeFeatureDetail(value: unknown): FeatureRecord {
  const raw = record(value, "feature");
  if ("geometry" in raw) {
    throw new Error("feature.geometry: map geometry is not part of the detail payload");
  }
  return {
    ...decodeSummary(raw, "feature"),
    attributes: decodeAttributes(raw.attributes, "feature.attributes"),
  };
}

function decodeSummary(raw: Record<string, unknown>, path: string) {
  return {
    layerFamily: enumValue(raw.layer_family, layerKeys, `${path}.layer_family`) as LayerFamilyKey,
    featureId: nonEmptyString(raw.feature_id, `${path}.feature_id`),
    title: nonEmptyString(raw.title, `${path}.title`),
    subtitle: stringValue(raw.subtitle, `${path}.subtitle`),
    sourceId: nonEmptyString(raw.source_id, `${path}.source_id`),
    geometryLabel: nonEmptyString(raw.geometry_label, `${path}.geometry_label`),
    geometrySource: optionalString(raw.geometry_source, `${path}.geometry_source`),
    lastSyncedAt: isoDateTime(raw.last_synced_at, `${path}.last_synced_at`),
  };
}

function decodeAttributes(value: unknown, path: string): Record<string, AttributeValue> {
  const raw = record(value, path);
  const result: Record<string, AttributeValue> = {};
  for (const [key, item] of Object.entries(raw)) {
    if (item !== null && typeof item !== "string" && typeof item !== "number" && typeof item !== "boolean") {
      throw new Error(`${path}.${key}: expected a scalar attribute`);
    }
    result[key] = item as AttributeValue;
  }
  return result;
}

function decodeGeometry(value: unknown, path: string): GeoJSONGeometry {
  const raw = record(value, path);
  const type = enumValue(raw.type, ["Polygon", "MultiPolygon"] as const, `${path}.type`);
  if (!Array.isArray(raw.coordinates)) throw new Error(`${path}.coordinates: expected an array`);
  validateCoordinateTree(raw.coordinates, `${path}.coordinates`);
  return { type, coordinates: raw.coordinates } as GeoJSONGeometry;
}

function validateCoordinateTree(value: unknown, path: string): void {
  if (!Array.isArray(value) || value.length === 0) throw new Error(`${path}: expected a non-empty array`);
  if (value.length >= 2 && typeof value[0] === "number" && typeof value[1] === "number") {
    longitude(value[0], `${path}[0]`);
    latitude(value[1], `${path}[1]`);
    return;
  }
  value.forEach((item, index) => validateCoordinateTree(item, `${path}[${index}]`));
}
