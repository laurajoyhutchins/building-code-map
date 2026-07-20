import { describe, expect, it } from "vitest";
import { createLayerSelectionMap } from "./layerSelection";

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

describe("createLayerSelectionMap", () => {
  it("preserves the configured default visibility for each boundary family", () => {
    const selection = createLayerSelectionMap(layerFamilies);

    expect(selection).toEqual({
      states: true,
      counties: true,
      municipalities: true,
      special_areas: false,
      tribal_areas: false,
      neris_jurisdictions: false,
    });
  });
});
