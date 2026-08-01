import type { FeatureRef, LayerFamilyKey } from "../types";

const layerFamilyKeyLookup = {
  states: true,
  counties: true,
  municipalities: true,
  special_areas: true,
  tribal_areas: true,
  neris_jurisdictions: true,
} satisfies Record<LayerFamilyKey, true>;

function isLayerFamilyKey(value: string): value is LayerFamilyKey {
  return Object.prototype.hasOwnProperty.call(layerFamilyKeyLookup, value);
}

export function featureRefsEqual(
  left: FeatureRef | null | undefined,
  right: FeatureRef | null | undefined,
): boolean {
  if (!left || !right) {
    return false;
  }

  return left.layerFamily === right.layerFamily && left.featureId === right.featureId;
}

export function findFeatureByRef<T extends FeatureRef>(
  features: readonly T[],
  featureRef: FeatureRef | null | undefined,
): T | null {
  if (!featureRef) {
    return null;
  }

  return features.find((feature) => featureRefsEqual(feature, featureRef)) ?? null;
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
    !isLayerFamilyKey(layerFamily) ||
    typeof featureId !== "string" ||
    featureId.length === 0
  ) {
    return null;
  }

  return {
    layerFamily,
    featureId,
  };
}
