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

export interface FeatureSummary {
  layerFamily: LayerFamilyKey;
  featureId: string;
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
