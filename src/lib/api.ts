import type {
  BoundaryAmbiguityDetails,
  BoundaryMapRecord,
  FeatureRecord,
  GeocodeResult,
  LayerFamilyDefinition,
  LayerFamilyKey,
  LookupResult,
  ReadinessResult,
  RefreshStatus,
  ResolutionBoundaryMatch,
  ResolutionResult,
} from "../types";
import {
  decodeGeocodeResult,
  decodeHealth,
  decodeLayers,
  decodeLookupResult,
  decodeReadiness,
  decodeRefreshStatus,
  decodeResolutionResult,
} from "./apiPayloads";
import { decodeBoundaryMapRecords, decodeFeatureDetail } from "./boundaryPayloads";
import {
  decodeEngineProvenance,
  decodeEngineResult,
  type EngineProvenance,
  type EngineResult,
} from "./engineV1";
import { arrayValue, nonEmptyString, record } from "./runtimeDecode";

const API_BASE = import.meta.env.VITE_API_BASE_URL?.trim() || "/api";

type Decoder<T> = (value: unknown) => T;

export class ApiResponseError extends Error {
  readonly status: number;
  readonly code?: string;
  readonly details?: BoundaryAmbiguityDetails;

  constructor(message: string, status: number, code?: string, details?: BoundaryAmbiguityDetails) {
    super(message);
    this.name = "ApiResponseError";
    this.status = status;
    this.code = code;
    this.details = details;
  }
}

export function buildApiUrl(path: string): string {
  const normalizedPath = path.startsWith("/") ? path : `/${path}`;
  if (/^https?:\/\//i.test(API_BASE)) {
    return new URL(normalizedPath.slice(1), `${API_BASE.replace(/\/+$/, "")}/`).toString();
  }
  return `${API_BASE.replace(/\/+$/, "")}${normalizedPath}`;
}

async function readJson<T>(path: string, decode: Decoder<T>, init?: RequestInit): Promise<T> {
  const { headers: requestHeaders, ...rest } = init ?? {};
  const headers = new Headers(requestHeaders);
  headers.set("Accept", "application/json");
  const response = await fetch(buildApiUrl(path), { headers, ...rest });
  let payload: unknown;
  try {
    payload = await response.json();
  } catch {
    throw new ApiResponseError(`Request to ${path} returned invalid JSON`, response.status);
  }
  if (!response.ok) {
    throw decodeApiError(payload, response.status, path);
  }
  try {
    return decode(payload);
  } catch (error) {
    const detail = error instanceof Error ? error.message : "unknown payload error";
    throw new ApiResponseError(`Invalid response from ${path}: ${detail}`, response.status);
  }
}

function decodeApiError(payload: unknown, status: number, path: string): ApiResponseError {
  let raw: Record<string, unknown>;
  try {
    raw = record(payload, `error response from ${path}`);
  } catch {
    return new ApiResponseError(`Request to ${path} failed with ${status}`, status);
  }

  const message =
    typeof raw.message === "string" && raw.message.trim() !== ""
      ? raw.message
      : typeof raw.error === "string" && raw.error.trim() !== ""
        ? raw.error
        : `Request to ${path} failed with ${status}`;
  const code = typeof raw.code === "string" ? raw.code : undefined;

  if (code === "boundary_ambiguous") {
    try {
      const details = raw.details === undefined ? raw : record(raw.details, "error.details");
      const layerFamily = nonEmptyString(details.layer_family, "error.details.layer_family");
      const observations = arrayValue(
        details.observations,
        "error.details.observations",
        decodeBoundaryMatch,
      );
      return new ApiResponseError(message, status, code, { layerFamily, observations });
    } catch {
      return new ApiResponseError(message, status, code);
    }
  }
  if (code === "address_ambiguous" || raw.status === "ambiguous") {
    const details =
      typeof raw.details === "object" && raw.details !== null && !Array.isArray(raw.details)
        ? (raw.details as Record<string, unknown>)
        : raw;
    const candidateCount =
      typeof details.candidate_count === "number"
        ? details.candidate_count
        : Array.isArray(raw.candidates)
          ? raw.candidates.length
          : undefined;
    return new ApiResponseError(
      `The address matched ${candidateCount || "multiple"} locations. Add a ZIP code or more specific locality.`,
      status,
      "geocoder_ambiguous",
    );
  }
  if (code === "address_not_found" || raw.status === "not_found") {
    return new ApiResponseError(
      "The local geocoder could not match that address.",
      status,
      "geocoder_not_found",
    );
  }
  return new ApiResponseError(message, status, code);
}

function decodeBoundaryMatch(value: unknown, path: string): ResolutionBoundaryMatch {
  const raw = record(value, path);
  return {
    layerFamily: nonEmptyString(raw.layer_family, `${path}.layer_family`),
    featureId: nonEmptyString(raw.feature_id, `${path}.feature_id`),
    name: nonEmptyString(raw.name, `${path}.name`),
    sourceId: typeof raw.source_id === "string" ? raw.source_id : undefined,
  };
}

export function getApiBaseUrl(): string {
  return API_BASE;
}

export function fetchHealth(): Promise<{ status: string }> {
  return readJson("/health", decodeHealth);
}

export function fetchReadiness(signal?: AbortSignal): Promise<ReadinessResult> {
  return readJson("/v1/readiness", decodeReadiness, { signal });
}

export function fetchLayers(): Promise<LayerFamilyDefinition[]> {
  return readJson("/layers", decodeLayers);
}

export function fetchBoundaryFeatures(): Promise<BoundaryMapRecord[]> {
  return readJson("/boundaries", decodeBoundaryMapRecords);
}

export function fetchRefreshStatus(): Promise<RefreshStatus> {
  return readJson("/refresh/status", decodeRefreshStatus);
}

export function fetchFeature(
  layerFamily: LayerFamilyKey,
  featureId: FeatureRecord["featureId"],
  signal?: AbortSignal,
): Promise<FeatureRecord> {
  return readJson(
    `/features/${layerFamily}/${encodeURIComponent(featureId)}`,
    decodeFeatureDetail,
    {
      signal,
    },
  );
}

export interface ResolutionRequestInput {
  longitude: number;
  latitude: number;
  codeFamily?: string;
  projectType?: string;
  applicabilityDate?: string;
  signal?: AbortSignal;
}

export interface LookupRequestInput {
  address: string;
  codeFamily?: string;
  projectType?: string;
  applicabilityDate?: string;
  signal?: AbortSignal;
}

export function fetchResolution(input: ResolutionRequestInput): Promise<ResolutionResult> {
  return readJson("/resolve", decodeResolutionResult, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      point: { longitude: input.longitude, latitude: input.latitude },
      code_family: input.codeFamily || undefined,
      project_type: input.projectType || undefined,
      applicability_date: input.applicabilityDate || undefined,
    }),
    signal: input.signal,
  });
}

export function fetchEngineResolution(input: ResolutionRequestInput): Promise<EngineResult> {
  return readJson("/v1/resolve", decodeEngineResult, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      point: { longitude: input.longitude, latitude: input.latitude },
      code_family: input.codeFamily || undefined,
      project_type: input.projectType || undefined,
      applicability_date: input.applicabilityDate || undefined,
    }),
    signal: input.signal,
  });
}

export function fetchGeocode(address: string, signal?: AbortSignal): Promise<GeocodeResult> {
  return readJson("/v1/geocode", decodeGeocodeResult, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ address }),
    signal,
  });
}

export function fetchLookup(input: LookupRequestInput): Promise<LookupResult> {
  return readJson("/lookup", decodeLookupResult, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      address: input.address,
      code_family: input.codeFamily || undefined,
      project_type: input.projectType || undefined,
      applicability_date: input.applicabilityDate || undefined,
    }),
    signal: input.signal,
  });
}

export function fetchEngineAddressResolution(input: LookupRequestInput): Promise<EngineResult> {
  return readJson("/v1/resolve", decodeEngineResult, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      address: input.address,
      code_family: input.codeFamily || undefined,
      project_type: input.projectType || undefined,
      applicability_date: input.applicabilityDate || undefined,
    }),
    signal: input.signal,
  });
}

export function fetchBundle(signal?: AbortSignal): Promise<EngineProvenance> {
  return readJson("/v1/bundle", decodeEngineProvenance, { signal });
}
