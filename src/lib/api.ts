import type {
  BoundaryFeatureRecord,
  FeatureRecord,
  LayerFamilyDefinition,
  LayerFamilyKey,
  RefreshStatus,
} from "../types";

const API_BASE = import.meta.env.VITE_API_BASE_URL?.trim() || "/api";

export function buildApiUrl(path: string): string {
  const normalizedPath = path.startsWith("/") ? path : `/${path}`;

  if (/^https?:\/\//i.test(API_BASE)) {
    return new URL(normalizedPath.slice(1), `${API_BASE.replace(/\/+$/, "")}/`).toString();
  }

  return `${API_BASE.replace(/\/+$/, "")}${normalizedPath}`;
}

async function readJson<T>(path: string, init?: RequestInit): Promise<T> {
  const { headers: requestHeaders, ...rest } = init ?? {};
  const headers = new Headers(requestHeaders);
  headers.set("Accept", "application/json");

  const response = await fetch(buildApiUrl(path), {
    headers,
    ...rest,
  });

  if (!response.ok) {
    throw new Error(`Request to ${path} failed with ${response.status}`);
  }

  return (await response.json()) as T;
}

export function getApiBaseUrl(): string {
  return API_BASE;
}

export function fetchHealth(): Promise<{ status: string }> {
  return readJson<{ status: string }>("/health");
}

export function fetchLayers(): Promise<LayerFamilyDefinition[]> {
  return readJson<LayerFamilyDefinition[]>("/layers");
}

export function fetchBoundaryFeatures(): Promise<BoundaryFeatureRecord[]> {
  return readJson<BoundaryFeatureRecord[]>("/boundaries");
}

export function fetchRefreshStatus(): Promise<RefreshStatus> {
  return readJson<RefreshStatus>("/refresh/status");
}

export function fetchFeature(
  layerFamily: LayerFamilyKey,
  featureId: FeatureRecord["featureId"],
): Promise<FeatureRecord> {
  return readJson<FeatureRecord>(`/features/${layerFamily}/${featureId}`);
}
