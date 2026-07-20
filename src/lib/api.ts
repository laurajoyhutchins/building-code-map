import type {
  BoundaryFeatureRecord,
  FeatureRecord,
  LayerFamilyDefinition,
  LayerFamilyKey,
  RefreshStatus,
  ResolutionResult,
} from "../types";

const API_BASE = import.meta.env.VITE_API_BASE_URL?.trim() || "/api";

export function buildApiUrl(path: string): string {
  const normalizedPath = path.startsWith("/") ? path : `/${path}`;

  if (/^https?:\/\//i.test(API_BASE)) {
    return new URL(
      normalizedPath.slice(1),
      `${API_BASE.replace(/\/+$/, "")}/`,
    ).toString();
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

interface RawVerificationRecord {
  status: string;
  confidence?: number;
  notes?: string;
}

interface RawBoundaryMatch {
  layer_family: string;
  feature_id: string;
  name: string;
  source_id?: string;
}

interface RawResolutionResult {
  schema_version: string;
  generated_at: string;
  profile_id?: string;
  profile_last_verified?: string;
  geography: {
    state_id?: string;
    state_fips?: string;
    state_name?: string;
    county?: RawBoundaryMatch;
    municipality?: RawBoundaryMatch;
    incorporated: boolean;
    special_areas?: RawBoundaryMatch[];
    tribal_areas?: RawBoundaryMatch[];
    fire_jurisdictions?: RawBoundaryMatch[];
  };
  code_family?: string;
  project_type?: string;
  applicability_date?: string;
  status: ResolutionResult["status"];
  authority_candidates?: Array<{
    kind: string;
    authority_id?: string;
    name: string;
    roles?: string[];
    source_ids?: string[];
    verification: RawVerificationRecord;
  }>;
  adoptions?: Array<{
    id: string;
    code_family: string;
    status: string;
    state_code_name: string;
    enforcement_model: string;
    dates?: Record<string, string>;
    source_ids?: string[];
    verification: RawVerificationRecord;
  }>;
  required_local_records?: string[];
  warnings?: string[];
  evidence?: Array<{
    id: string;
    title: string;
    url: string;
    kind: string;
    accessed_at: string;
    caveat?: string;
  }>;
}

export interface ResolutionRequestInput {
  longitude: number;
  latitude: number;
  codeFamily?: string;
  projectType?: string;
  applicabilityDate?: string;
}

export function decodeResolutionResult(
  raw: RawResolutionResult,
): ResolutionResult {
  const mapBoundary = (match: RawBoundaryMatch) => ({
    layerFamily: match.layer_family,
    featureId: match.feature_id,
    name: match.name,
    sourceId: match.source_id,
  });

  return {
    schemaVersion: raw.schema_version,
    generatedAt: raw.generated_at,
    profileId: raw.profile_id,
    profileLastVerified: raw.profile_last_verified,
    geography: {
      stateId: raw.geography.state_id,
      stateFips: raw.geography.state_fips,
      stateName: raw.geography.state_name,
      county: raw.geography.county
        ? mapBoundary(raw.geography.county)
        : undefined,
      municipality: raw.geography.municipality
        ? mapBoundary(raw.geography.municipality)
        : undefined,
      incorporated: raw.geography.incorporated,
      specialAreas: (raw.geography.special_areas ?? []).map(mapBoundary),
      tribalAreas: (raw.geography.tribal_areas ?? []).map(mapBoundary),
      fireJurisdictions: (raw.geography.fire_jurisdictions ?? []).map(
        mapBoundary,
      ),
    },
    codeFamily: raw.code_family,
    projectType: raw.project_type,
    applicabilityDate: raw.applicability_date,
    status: raw.status,
    authorityCandidates: (raw.authority_candidates ?? []).map((candidate) => ({
      kind: candidate.kind,
      authorityId: candidate.authority_id,
      name: candidate.name,
      roles: candidate.roles ?? [],
      sourceIds: candidate.source_ids ?? [],
      verification: candidate.verification,
    })),
    adoptions: (raw.adoptions ?? []).map((adoption) => ({
      id: adoption.id,
      codeFamily: adoption.code_family,
      status: adoption.status,
      stateCodeName: adoption.state_code_name,
      enforcementModel: adoption.enforcement_model,
      dates: adoption.dates ?? {},
      sourceIds: adoption.source_ids ?? [],
      verification: adoption.verification,
    })),
    requiredLocalRecords: raw.required_local_records ?? [],
    warnings: raw.warnings ?? [],
    evidence: (raw.evidence ?? []).map((source) => ({
      id: source.id,
      title: source.title,
      url: source.url,
      kind: source.kind,
      accessedAt: source.accessed_at,
      caveat: source.caveat,
    })),
  };
}

export async function fetchResolution(
  input: ResolutionRequestInput,
): Promise<ResolutionResult> {
  const raw = await readJson<RawResolutionResult>("/resolve", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({
      point: { longitude: input.longitude, latitude: input.latitude },
      code_family: input.codeFamily || undefined,
      project_type: input.projectType || undefined,
      applicability_date: input.applicabilityDate || undefined,
    }),
  });

  return decodeResolutionResult(raw);
}
