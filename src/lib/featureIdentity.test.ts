import { describe, expect, it } from "vitest";
import type { FeatureSummary } from "../types";
import { encodeFeatureRef, featureRefsEqual, findFeatureByRef } from "./featureIdentity";

const collidingFeatures: FeatureSummary[] = [
  {
    layerFamily: "states",
    featureId: "08",
    title: "Colorado",
    subtitle: "State boundary",
    sourceId: "STATEFP=08",
  },
  {
    layerFamily: "counties",
    featureId: "08",
    title: "County 08",
    subtitle: "County boundary",
    sourceId: "COUNTYFP=08",
  },
];

describe("feature identity helpers", () => {
  it("does not treat two missing references as a selected identity", () => {
    expect(featureRefsEqual(null, null)).toBe(false);
  });

  it("resolves the exact feature when layer families reuse a feature id", () => {
    expect(
      findFeatureByRef(collidingFeatures, {
        layerFamily: "counties",
        featureId: "08",
      }),
    ).toEqual(collidingFeatures[1]);
  });

  it("encodes different layer families as different collection keys", () => {
    expect(
      encodeFeatureRef({
        layerFamily: "states",
        featureId: "08",
      }),
    ).not.toBe(
      encodeFeatureRef({
        layerFamily: "counties",
        featureId: "08",
      }),
    );
  });
});
