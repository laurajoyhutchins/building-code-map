import type { FeatureRef, LayerFamilyKey } from "../types";

const layerFamilyKeys = new Set<LayerFamilyKey>([
  "states",
  "counties",
  "municipalities",
  "special_areas",
  "tribal_areas",
  "neris_jurisdictions",
]);

export function featureRefsEqual(
  left: FeatureRef | null | undefined,
  right: FeatureRef | null | undefined,
): boolean {
  return (
    left?.layerFamily === right?.layerFamily &&
    left?.featureId === right?.featureId &&
    Boolean(left) === Boolean(right)
  );
}

export function encodeFeatureRef(feature: FeatureRef): string {
  return `${feature.layerFamily}:${encodeURIComponent(feature.featureId)}`;
}

export function featureRefFromMapProperties(properties: unknown): FeatureRef | null {
  if (!properties || typeof properties !== "object") {
    return null;
  }

  const { layerFamily, featureId } = properties as Record<string, unknown>;
  if (
    typeof layerFamily !== "string" ||
    !layerFamilyKeys.has(layerFamily as LayerFamilyKey) ||
    typeof featureId !== "string" ||
    featureId.length === 0
  ) {
    return null;
  }

  return {
    layerFamily: layerFamily as LayerFamilyKey,
    featureId,
  };
}
