import { describe, expect, it } from "vitest";
import {
  formatCodeFamily,
  getResolutionNotice,
  getResolutionPlace,
  parseCoordinateQuery,
} from "./publicLookup";
import type { ResolutionResult } from "../types";

function makeResult(overrides: Partial<ResolutionResult> = {}): ResolutionResult {
  return {
    schemaVersion: "1.0.0",
    generatedAt: "2026-08-01T00:00:00Z",
    geography: {
      stateName: "Colorado",
      incorporated: true,
      municipality: {
        layerFamily: "municipalities",
        featureId: "denver",
        name: "Denver",
      },
      county: {
        layerFamily: "counties",
        featureId: "denver-county",
        name: "Denver County",
      },
      specialAreas: [],
      tribalAreas: [],
      fireJurisdictions: [],
    },
    status: "resolved",
    authorityCandidates: [],
    authorityPath: [],
    adoptions: [],
    applicableRules: [],
    supportingClaims: [],
    requiredLocalRecords: [],
    warnings: [],
    evidence: [],
    ...overrides,
  };
}

describe("parseCoordinateQuery", () => {
  it("parses latitude followed by longitude", () => {
    expect(parseCoordinateQuery("39.7392, -104.9903")).toEqual({
      latitude: 39.7392,
      longitude: -104.9903,
    });
  });

  it("accepts longitude followed by latitude when the first value cannot be a latitude", () => {
    expect(parseCoordinateQuery("-104.9903 39.7392")).toEqual({
      latitude: 39.7392,
      longitude: -104.9903,
    });
  });

  it("rejects incomplete and out-of-range coordinates", () => {
    expect(() => parseCoordinateQuery("39.7392")).toThrow("two coordinates");
    expect(() => parseCoordinateQuery("95, 200")).toThrow("valid latitude and longitude");
  });
});

describe("public result presentation", () => {
  it("uses the most specific matched place", () => {
    expect(getResolutionPlace(makeResult())).toBe("Denver, Colorado");
  });

  it("does not decorate a resolved result with a status message", () => {
    expect(getResolutionNotice("resolved")).toBeNull();
  });

  it("explains incomplete states in ordinary language", () => {
    expect(getResolutionNotice("local_record_required")).toBe(
      "Contact the applicable local authority to confirm the records identified below.",
    );
    expect(getResolutionNotice("conflicting")).toBe("The available official records conflict.");
  });

  it("formats machine code-family names for display", () => {
    expect(formatCodeFamily("fire_operational")).toBe("Operational fire");
    expect(formatCodeFamily(undefined)).toBe("Building codes");
  });
});
