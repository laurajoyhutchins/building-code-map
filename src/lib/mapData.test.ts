import { describe, expect, it } from "vitest";
import { createLayerSelectionMap } from "./layerSelection";
import { buildBoundaryFeatureCollection, calculateBoundaryBounds } from "./mapData";

const layerFamilies = [
  {
    key: "states",
    label: "States",
    martinLayerId: "tigerweb.states",
    description: "State boundaries mirrored from TIGERweb.",
    defaultEnabled: true,
  },
  {
    key: "counties",
    label: "Counties",
    martinLayerId: "tigerweb.counties",
    description: "County boundaries used for jurisdiction lookups.",
    defaultEnabled: true,
  },
  {
    key: "municipalities",
    label: "Municipalities",
    martinLayerId: "tigerweb.municipalities",
    description: "Incorporated places and municipal limits.",
    defaultEnabled: true,
  },
  {
    key: "special_areas",
    label: "Special areas",
    martinLayerId: "tigerweb.special_areas",
    description: "Military installations and other special land use areas.",
    defaultEnabled: false,
  },
  {
    key: "tribal_areas",
    label: "American Indian areas",
    martinLayerId: "tigerweb.tribal_areas",
    description: "Tribal boundaries and related areas.",
    defaultEnabled: false,
  },
  {
    key: "neris_jurisdictions",
    label: "NERIS jurisdictions",
    martinLayerId: "neris.department_jurisdictions",
    description: "Department jurisdiction polygons from NERIS.",
    defaultEnabled: false,
  },
] as const;

const demoGeometry = {
  type: "Polygon",
  coordinates: [
    [
      [-109.05, 37.0],
      [-102.05, 37.0],
      [-102.05, 41.0],
      [-109.05, 41.0],
      [-109.05, 37.0],
    ],
  ],
} as const;

const demoFeatures = [
  {
    layerFamily: "states",
    featureId: "08",
    title: "Colorado",
    subtitle: "State boundary",
    sourceId: "STATEFP=08",
    geometry: demoGeometry,
  },
  {
    layerFamily: "tribal_areas",
    featureId: "3925",
    title: "Southern Ute",
    subtitle: "American Indian area",
    sourceId: "AIANNH=3925",
    geometry: {
      type: "Polygon",
      coordinates: [
        [
          [-108.65, 36.85],
          [-107.8, 36.85],
          [-107.8, 37.4],
          [-108.65, 37.4],
          [-108.65, 36.85],
        ],
      ],
    } as const,
  },
  {
    layerFamily: "neris_jurisdictions",
    featureId: "FD01001122",
    title: "Independence Volunteer Fire Department",
    subtitle: "Department jurisdiction",
    sourceId: "neris_id=FD01001122",
    geometrySource: "neris_department_jurisdiction",
    geometry: {
      type: "Polygon",
      coordinates: [
        [
          [-105.25, 39.65],
          [-104.98, 39.65],
          [-104.98, 39.82],
          [-105.25, 39.82],
          [-105.25, 39.65],
        ],
      ],
    } as const,
  },
] as const;

describe("buildBoundaryFeatureCollection", () => {
  it("filters out disabled boundary families and marks the active selection", () => {
    const enabledLayers = {
      ...createLayerSelectionMap(layerFamilies),
      tribal_areas: false,
      neris_jurisdictions: true,
    };

    const collection = buildBoundaryFeatureCollection(demoFeatures, enabledLayers, "08");

    expect(collection.type).toBe("FeatureCollection");
    expect(collection.features).toHaveLength(2);
    expect(collection.features[0].properties).toMatchObject({
      layerFamily: "states",
      featureId: "08",
      selected: true,
    });
    expect(collection.features[1].properties).toMatchObject({
      layerFamily: "neris_jurisdictions",
      featureId: "FD01001122",
      geometrySource: "neris_department_jurisdiction",
      selected: false,
    });
  });
});

describe("calculateBoundaryBounds", () => {
  it("computes the bounding box for polygon coordinates", () => {
    expect(calculateBoundaryBounds(demoGeometry)).toEqual([
      [-109.05, 37.0],
      [-102.05, 41.0],
    ]);
  });

  it("computes the bounding box for multipolygon coordinates", () => {
    expect(
      calculateBoundaryBounds({
        type: "MultiPolygon",
        coordinates: [
          [
            [
              [-106.0, 39.0],
              [-105.5, 39.0],
              [-105.5, 39.5],
              [-106.0, 39.0],
            ],
          ],
          [
            [
              [-104.9, 39.1],
              [-104.7, 39.1],
              [-104.7, 39.3],
              [-104.9, 39.1],
            ],
          ],
        ],
      }),
    ).toEqual([
      [-106.0, 39.0],
      [-104.7, 39.5],
    ]);
  });
});
