import type {
  FeatureRef,
  GeoJSONGeometry,
  LayerFamilyKey,
  LayerSelectionMap,
  PolygonGeometry,
} from "../types";
import { featureRefsEqual } from "./featureIdentity";

export interface BoundaryRenderableFeature extends FeatureRef {
  title: string;
  subtitle: string;
  sourceId: string;
  geometry: GeoJSONGeometry;
  geometrySource?: string | null;
}

export interface BoundaryGeoJsonFeature {
  type: "Feature";
  geometry: GeoJSONGeometry;
  properties: {
    layerFamily: LayerFamilyKey;
    featureId: string;
    title: string;
    subtitle: string;
    sourceId: string;
    geometrySource?: string | null;
    selected: boolean;
  };
}

export interface BoundaryFeatureCollection {
  type: "FeatureCollection";
  features: BoundaryGeoJsonFeature[];
}

function getGeometryBounds(geometry: GeoJSONGeometry): [[number, number], [number, number]] {
  let minLng = Number.POSITIVE_INFINITY;
  let minLat = Number.POSITIVE_INFINITY;
  let maxLng = Number.NEGATIVE_INFINITY;
  let maxLat = Number.NEGATIVE_INFINITY;

  for (const polygon of iterPolygons(geometry)) {
    for (const ring of polygon.coordinates) {
      for (const [lng, lat] of ring) {
        minLng = Math.min(minLng, lng);
        minLat = Math.min(minLat, lat);
        maxLng = Math.max(maxLng, lng);
        maxLat = Math.max(maxLat, lat);
      }
    }
  }

  if (
    minLng === Number.POSITIVE_INFINITY ||
    minLat === Number.POSITIVE_INFINITY ||
    maxLng === Number.NEGATIVE_INFINITY ||
    maxLat === Number.NEGATIVE_INFINITY
  ) {
    throw new Error("Polygon geometry does not contain any coordinates");
  }

  return [
    [minLng, minLat],
    [maxLng, maxLat],
  ];
}

export function calculateBoundaryBounds(
  geometry: GeoJSONGeometry,
): [[number, number], [number, number]] {
  return getGeometryBounds(geometry);
}

export function buildBoundaryFeatureCollection(
  features: readonly BoundaryRenderableFeature[],
  enabledLayers: LayerSelectionMap,
  selectedFeature: FeatureRef | null,
): BoundaryFeatureCollection {
  return {
    type: "FeatureCollection",
    features: features
      .filter((feature) => Boolean(enabledLayers[feature.layerFamily]))
      .map((feature) => ({
        type: "Feature",
        geometry: feature.geometry,
        properties: {
          layerFamily: feature.layerFamily,
          featureId: feature.featureId,
          title: feature.title,
          subtitle: feature.subtitle,
          sourceId: feature.sourceId,
          geometrySource: feature.geometrySource ?? null,
          selected: featureRefsEqual(feature, selectedFeature),
        },
      })),
  };
}

function iterPolygons(geometry: GeoJSONGeometry): PolygonGeometry[] {
  if (geometry.type === "Polygon") {
    return [geometry];
  }

  return geometry.coordinates.map((coordinates) => ({
    type: "Polygon",
    coordinates,
  })) as PolygonGeometry[];
}
