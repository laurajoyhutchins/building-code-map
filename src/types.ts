export type LayerFamilyKey =
  | "states"
  | "counties"
  | "municipalities"
  | "special_areas"
  | "tribal_areas"
  | "neris_jurisdictions";

export type AttributeValue = string | number | boolean | null;
export type PolygonPosition = readonly [number, number];

export interface PolygonGeometry {
  type: "Polygon";
  coordinates: readonly (readonly PolygonPosition[])[];
}

export interface MultiPolygonGeometry {
  type: "MultiPolygon";
  coordinates: readonly (readonly (readonly PolygonPosition[])[])[];
}

export type GeoJSONGeometry = PolygonGeometry | MultiPolygonGeometry;

export interface LayerFamilyDefinition {
  key: LayerFamilyKey;
  label: string;
  martinLayerId: string;
  description: string;
  defaultEnabled: boolean;
}

export type LayerSelectionMap = Readonly<Record<LayerFamilyKey, boolean>>;

export interface FeatureRef {
  layerFamily: LayerFamilyKey;
  featureId: string;
}

export interface FeatureSummary extends FeatureRef {
  title: string;
  subtitle: string;
  sourceId: string;
}

export interface FeatureRecord extends FeatureSummary {
  geometryLabel: string;
  geometrySource?: string | null;
  lastSyncedAt: string;
  attributes: Record<string, AttributeValue>;
  geometry?: GeoJSONGeometry;
}

export interface BoundaryFeatureRecord extends FeatureRecord {
  geometry: GeoJSONGeometry;
}

export interface RefreshStatus {
  status: "ok" | "warning" | "error";
  latestSuccessfulRefresh: string | null;
  latestAttempt: string;
  nextScheduledRefresh: string;
  message: string;
}

export interface CapabilityReadiness {
  status: "available" | "unavailable";
  required: boolean;
  message: string;
}

export interface SnapshotReadiness {
  status: "verified" | "unidentified" | "unavailable";
  snapshotId?: string;
}

export interface ReadinessResult {
  status: "ok" | "not_ready";
  readiness: "ready" | "degraded" | "not_ready";
  capabilities: Record<string, CapabilityReadiness>;
  snapshots: Record<string, SnapshotReadiness>;
}

export type GeocodeStatus = "matched" | "ambiguous" | "not_found";
export type GeocodePrecision = "address_point" | "interpolated";

export interface CoordinateEvidence {
  longitude: number;
  latitude: number;
}

export interface InterpolationProvenance {
  sourceRangeId: string;
  requestedHouseNumber: number;
  fromNumber: number;
  toNumber: number;
  rangeDirection: "ascending" | "descending";
  parity: string;
  side?: string;
  fromCoordinate: CoordinateEvidence;
  toCoordinate: CoordinateEvidence;
  fraction: number;
  derivedCoordinate: CoordinateEvidence;
  coordinateReferenceSystem: string;
  transformationIdentity: string;
  methodVersion: string;
  positionalQuality: string;
}

export interface GeocodeCandidate {
  matchedAddress: string;
  longitude: number;
  latitude: number;
  precision: GeocodePrecision;
  confidence: number;
  scoreKind: "deterministic_quality";
  scoreFactors: Record<string, number>;
  rankingPolicyVersion: string;
  source: string;
  sourceRecordId: string;
  sourceVintage: string;
  interpolation?: InterpolationProvenance;
}

export interface GeocodeResult {
  query: string;
  normalized?: string;
  status: GeocodeStatus;
  selected?: GeocodeCandidate;
  candidates: GeocodeCandidate[];
  warnings: string[];
}

export type ResolutionStatus =
  | "resolved"
  | "partially_resolved"
  | "local_record_required"
  | "ambiguous"
  | "conflicting"
  | "insufficient_evidence";

export interface VerificationRecord {
  status: string;
  confidence?: number;
  notes?: string;
}

export interface RegulatorySource {
  id: string;
  title: string;
  url: string;
  kind: string;
  accessedAt: string;
  lastCheckedAt?: string;
  availability?: "available" | "unavailable" | "moved" | "unknown";
  caveat?: string;
}

export interface ResolutionBoundaryMatch {
  layerFamily: string;
  featureId: string;
  name: string;
  sourceId?: string;
}

export interface BoundaryAmbiguityDetails {
  layerFamily: string;
  observations: ResolutionBoundaryMatch[];
}

export interface ResolutionGeography {
  stateId?: string;
  stateFips?: string;
  stateName?: string;
  county?: ResolutionBoundaryMatch;
  municipality?: ResolutionBoundaryMatch;
  incorporated: boolean;
  specialAreas: ResolutionBoundaryMatch[];
  tribalAreas: ResolutionBoundaryMatch[];
  fireJurisdictions: ResolutionBoundaryMatch[];
}

export interface ResolutionPolicyBasis {
  status: ResolutionStatus;
  requiredLocalRecords: string[];
  warnings: string[];
  sourceIds: string[];
  verification: VerificationRecord;
}

export interface ResolutionAuthorityCandidate {
  kind: string;
  authorityId?: string;
  name: string;
  roles: string[];
  sourceIds: string[];
  verification: VerificationRecord;
}

export interface ResolutionAuthorityRelationship {
  id: string;
  fromId: string;
  relationship: string;
  to: string;
  scope: string[];
  summary?: string;
  sourceIds: string[];
  verification: VerificationRecord;
}

export interface ResolutionAdoption {
  id: string;
  codeFamily: string;
  status: string;
  stateCodeName: string;
  enforcementModel: string;
  dates: Record<string, string>;
  sourceIds: string[];
  verification: VerificationRecord;
}

export interface ResolutionRuleReference {
  id: string;
  kind: "applicability" | "date" | "amendment" | "enforcement";
  codeFamily?: string;
  summary: string;
  sourceIds: string[];
  verification: VerificationRecord;
}

export interface ResolutionClaim {
  id: string;
  subjectId: string;
  field: string;
  status: string;
  value?: unknown;
  conflictGroup?: string;
  sourceIds: string[];
  verification: VerificationRecord;
}

export interface ResolutionResult {
  schemaVersion: string;
  generatedAt: string;
  profileId?: string;
  profileLastVerified?: string;
  geography: ResolutionGeography;
  codeFamily?: string;
  projectType?: string;
  applicabilityDate?: string;
  status: ResolutionStatus;
  policyBasis?: ResolutionPolicyBasis;
  authorityCandidates: ResolutionAuthorityCandidate[];
  authorityPath: ResolutionAuthorityRelationship[];
  adoptions: ResolutionAdoption[];
  applicableRules: ResolutionRuleReference[];
  supportingClaims: ResolutionClaim[];
  requiredLocalRecords: string[];
  warnings: string[];
  evidence: RegulatorySource[];
}

export interface LookupResult {
  geocode: GeocodeResult;
  resolution: ResolutionResult;
}
