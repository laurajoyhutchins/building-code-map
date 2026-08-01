import type { Map as MapLibreMap } from "maplibre-gl";
import { describe, expect, it, vi } from "vitest";
import { featureRefFromMapProperties } from "../lib/featureIdentity";
import type { LayerFamilyDefinition, LayerSelectionMap } from "../types";
import { reconcileBoundaryLayers } from "./MapStage";

const statesLayer: LayerFamilyDefinition = {
  key: "states",
  label: "States",
  martinLayerId: "states",
  description: "State boundaries",
  defaultEnabled: true,
};

const enabledLayers: LayerSelectionMap = {
  states: true,
  counties: false,
  municipalities: false,
  special_areas: false,
  tribal_areas: false,
  neris_jurisdictions: false,
};

const collection = {
  type: "FeatureCollection" as const,
  features: [],
};

function createFakeMap() {
  const sources = new Map<string, { setData: ReturnType<typeof vi.fn> }>();
  const layers = new Set<string>();
  const visibility = new Map<string, unknown>();

  const map = {
    getSource: vi.fn((id: string) => sources.get(id)),
    addSource: vi.fn((id: string) => {
      sources.set(id, { setData: vi.fn() });
    }),
    getLayer: vi.fn((id: string) => (layers.has(id) ? { id } : undefined)),
    addLayer: vi.fn((layer: { id: string }) => {
      layers.add(layer.id);
    }),
    setLayoutProperty: vi.fn((id: string, property: string, value: unknown) => {
      visibility.set(`${id}:${property}`, value);
    }),
  };

  return {
    map: map as unknown as MapLibreMap,
    sources,
    layers,
    visibility,
    addSource: map.addSource,
    addLayer: map.addLayer,
  };
}

describe("featureRefFromMapProperties", () => {
  it("returns both identity fields from a rendered map feature", () => {
    expect(
      featureRefFromMapProperties({
        layerFamily: "counties",
        featureId: "08",
      }),
    ).toEqual({
      layerFamily: "counties",
      featureId: "08",
    });
  });

  it("rejects incomplete or unknown map feature identities", () => {
    expect(featureRefFromMapProperties({ featureId: "08" })).toBeNull();
    expect(
      featureRefFromMapProperties({
        layerFamily: "unsupported",
        featureId: "08",
      }),
    ).toBeNull();
  });
});

describe("reconcileBoundaryLayers", () => {
  it("adds registry layers when the map becomes ready before registry hydration", () => {
    const fake = createFakeMap();

    reconcileBoundaryLayers(fake.map, [], enabledLayers, collection);
    expect(fake.layers).toHaveLength(0);

    reconcileBoundaryLayers(fake.map, [statesLayer], enabledLayers, collection);

    expect(fake.addSource).toHaveBeenCalledTimes(1);
    expect([...fake.layers]).toEqual([
      "tigerweb-boundaries-states-fill",
      "tigerweb-boundaries-states-line",
      "tigerweb-boundaries-states-selected",
    ]);
    expect(fake.visibility.get("tigerweb-boundaries-states-fill:visibility")).toBe("visible");
  });

  it("is idempotent when the registry is ready before map reconciliation", () => {
    const fake = createFakeMap();

    reconcileBoundaryLayers(fake.map, [statesLayer], enabledLayers, collection);
    const firstLayerState = [...fake.layers];
    reconcileBoundaryLayers(fake.map, [statesLayer], enabledLayers, collection);

    expect(fake.addSource).toHaveBeenCalledTimes(1);
    expect(fake.addLayer).toHaveBeenCalledTimes(3);
    expect([...fake.layers]).toEqual(firstLayerState);
    expect(fake.sources.get("tigerweb-boundaries")?.setData).toHaveBeenCalledTimes(2);
  });
});
