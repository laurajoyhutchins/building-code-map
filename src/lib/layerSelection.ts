import type { LayerFamilyDefinition, LayerSelectionMap } from "../types";

export function createLayerSelectionMap(
  layers: readonly LayerFamilyDefinition[],
): LayerSelectionMap {
  const selection: Record<LayerFamilyDefinition["key"], boolean> = {
    states: false,
    counties: false,
    municipalities: false,
    special_areas: false,
    tribal_areas: false,
    neris_jurisdictions: false,
  };

  for (const layer of layers) {
    selection[layer.key] = layer.defaultEnabled;
  }

  return selection;
}
