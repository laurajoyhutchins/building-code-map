import { afterEach, describe, expect, it, vi } from "vitest";
import { buildApiUrl, fetchReadiness, getApiBaseUrl } from "./api";
import {
  decodeGeocodeResult,
  decodeLookupResult,
  decodeReadiness,
  decodeResolutionResult,
} from "./apiPayloads";

afterEach(() => {
  vi.unstubAllGlobals();
});

describe("api url helpers", () => {
  it("defaults to the dev proxy path", () => {
    expect(getApiBaseUrl()).toBe("/api");
    expect(buildApiUrl("/health")).toBe("/api/health");
  });

  it("normalizes missing leading slashes", () => {
    expect(buildApiUrl("ready")).toBe("/api/ready");
  });
});

const rawResolution = {
  schema_version: "1.0",
  generated_at: "2026-07-20T00:00:00Z",
  profile_id: "state-profile:us-fl",
  profile_last_verified: "2026-07-20",
  geography: {
    state_id: "US-FL",
    state_fips: "12",
    state_name: "Florida",
    county: {
      layer_family: "counties",
      feature_id: "12095",
      name: "Orange County",
    },
    municipality: {
      layer_family: "municipalities",
      feature_id: "1253000",
      name: "Orlando",
    },
    incorporated: true,
  },
  code_family: "building",
  applicability_date: "2026-07-20",
  status: "partially_resolved" as const,
  policy_basis: {
    status: "partially_resolved" as const,
    required_local_records: ["Local amendment record"],
    source_ids: ["src:us-fl:statute-553-73"],
    verification: { status: "partially_verified", confidence: 0.9 },
  },
  authority_path: [
    {
      id: "edge:us-fl:state-code-local-enforcement",
      from_id: "auth:us-fl:building-commission",
      relationship: "delegates_enforcement_to",
      to: "local enforcing agencies",
      source_ids: ["src:us-fl:statute-553-73"],
      verification: { status: "partially_verified", confidence: 0.8 },
    },
  ],
  applicable_rules: [
    {
      id: "rule:us-fl:2023-code-effective",
      kind: "date" as const,
      code_family: "building",
      summary: "The current edition applies on the requested date.",
      source_ids: ["src:us-fl:fbc-current"],
      verification: { status: "verified", confidence: 0.95 },
    },
  ],
  supporting_claims: [
    {
      id: "claim:us-fl:building-effective-date",
      subject_id: "adoption:us-fl:building:2023",
      field: "dates.effective_date",
      status: "supported",
      value: "2023-12-31",
      source_ids: ["src:us-fl:fbc-current"],
      verification: { status: "verified", confidence: 0.95 },
    },
  ],
  evidence: [
    {
      id: "src:us-fl:fbc-current",
      title: "Current Florida Building Code",
      url: "https://example.com/code",
      kind: "official_web",
      accessed_at: "2026-07-20",
      last_checked_at: "2026-07-20",
      availability: "available" as const,
    },
  ],
  future_optional_field: { retained_by_server: true },
};

describe("resolution response decoding", () => {
  it("maps the complete versioned provenance contract without inventing missing records", () => {
    const result = decodeResolutionResult(rawResolution);

    expect(result.geography.municipality?.name).toBe("Orlando");
    expect(result.applicabilityDate).toBe("2026-07-20");
    expect(result.policyBasis?.sourceIds).toEqual(["src:us-fl:statute-553-73"]);
    expect(result.authorityPath[0]?.relationship).toBe("delegates_enforcement_to");
    expect(result.applicableRules[0]?.kind).toBe("date");
    expect(result.supportingClaims[0]?.value).toBe("2023-12-31");
    expect(result.evidence[0]?.availability).toBe("available");
    expect(result.adoptions).toEqual([]);
    expect(result.requiredLocalRecords).toEqual([]);
  });

  it("rejects malformed dates and nested verification records", () => {
    expect(() => decodeResolutionResult({ ...rawResolution, generated_at: "yesterday" })).toThrow(
      /generated_at/,
    );
    expect(() =>
      decodeResolutionResult({
        ...rawResolution,
        policy_basis: { ...rawResolution.policy_basis, verification: null },
      }),
    ).toThrow(/policy_basis\.verification/);
  });

  it("rejects invalid source URLs", () => {
    expect(() =>
      decodeResolutionResult({
        ...rawResolution,
        evidence: [{ ...rawResolution.evidence[0], url: "file:///private/code" }],
      }),
    ).toThrow(/HTTP or HTTPS/);
  });
});

describe("readiness response decoding", () => {
  const rawReadiness = {
    status: "ok",
    readiness: "degraded",
    capabilities: {
      boundary_resolution: {
        status: "available",
        required: true,
        message: "Boundary snapshot loaded.",
      },
    },
    snapshots: {
      boundary: { status: "verified", snapshot_id: "boundary-2026-08-05" },
      geocoder: { status: "unavailable" },
    },
  };

  it("preserves capability and snapshot identity", () => {
    const result = decodeReadiness(rawReadiness);
    expect(result.snapshots.boundary?.snapshotId).toBe("boundary-2026-08-05");
    expect(result.capabilities.boundary_resolution?.required).toBe(true);
  });

  it("rejects verified snapshots without identities", () => {
    expect(() =>
      decodeReadiness({
        status: "ok",
        readiness: "degraded",
        capabilities: {},
        snapshots: { boundary: { status: "verified" } },
      }),
    ).toThrow(/snapshot_id/);
  });

  it("forwards cancellation to the readiness request", async () => {
    const fetchMock = vi.fn(async () => ({
      ok: true,
      status: 200,
      json: async () => rawReadiness,
    }));
    vi.stubGlobal("fetch", fetchMock);
    const controller = new AbortController();

    await fetchReadiness(controller.signal);

    expect(fetchMock).toHaveBeenCalledWith(
      "/api/ready",
      expect.objectContaining({ signal: controller.signal }),
    );
  });

  it("keeps malformed error payloads inside the typed API boundary", async () => {
    vi.stubGlobal(
      "fetch",
      vi.fn(async () => ({
        ok: false,
        status: 503,
        json: async () => ["unexpected", "shape"],
      })),
    );

    await expect(fetchReadiness()).rejects.toMatchObject({
      name: "ApiResponseError",
      status: 503,
      message: "Request to /ready failed with 503",
    });
  });
});

describe("geocoder response decoding", () => {
  const rawGeocode = {
    query: "1600 N Broadway, Denver, CO 80202",
    normalized: "1600 N BROADWAY, DENVER, CO 80202",
    status: "matched" as const,
    selected: {
      matched_address: "1600 N Broadway St Denver CO 80202",
      longitude: -104.9876,
      latitude: 39.7411,
      precision: "address_point" as const,
      confidence: 1,
      score_kind: "deterministic_quality",
      score_factors: {
        address_point_base: 0.7,
        exact_street: 0.05,
        exact_city: 0.15,
        exact_postal_code: 0.1,
      },
      ranking_policy_version: "geocoder-ranking-1.0",
      source: "Denver address points",
      source_record_id: "co-denver-1600",
      source_vintage: "2026-08-01",
    },
    candidates: [],
    warnings: [],
  };

  it("preserves selected-point provenance and deterministic scoring identity", () => {
    const result = decodeGeocodeResult(rawGeocode);
    expect(result.selected?.sourceRecordId).toBe("co-denver-1600");
    expect(result.selected?.precision).toBe("address_point");
    expect(result.selected?.sourceVintage).toBe("2026-08-01");
    expect(result.selected?.rankingPolicyVersion).toBe("geocoder-ranking-1.0");
    expect(result.selected?.scoreFactors).toEqual(rawGeocode.selected.score_factors);
  });

  it("preserves interpolation provenance and deterministic quality scoring", () => {
    const result = decodeGeocodeResult({
      ...rawGeocode,
      selected: {
        ...rawGeocode.selected,
        precision: "interpolated" as const,
        confidence: 0.9,
        score_factors: {
          street_range_base: 0.55,
          exact_street: 0.05,
          exact_city: 0.15,
          exact_postal_code: 0.1,
          parity_matched: 0.05,
        },
        interpolation: {
          source_range_id: "range-17",
          requested_house_number: 1510,
          from_number: 1500,
          to_number: 1520,
          range_direction: "ascending",
          parity: "E",
          from_coordinate: { longitude: -104.99, latitude: 39.74 },
          to_coordinate: { longitude: -104.98, latitude: 39.75 },
          fraction: 0.5,
          derived_coordinate: { longitude: -104.985, latitude: 39.745 },
          coordinate_reference_system: "EPSG:4326",
          transformation_identity: "none",
          method_version: "linear-street-range-1.0",
          positional_quality: "street_range_interpolation",
        },
      },
    });
    expect(result.selected?.interpolation?.sourceRangeId).toBe("range-17");
    expect(result.selected?.interpolation?.fraction).toBe(0.5);
    expect(result.selected?.scoreFactors.street_range_base).toBe(0.55);
  });

  it("rejects invalid coordinates and precision", () => {
    expect(() =>
      decodeGeocodeResult({
        ...rawGeocode,
        selected: { ...rawGeocode.selected, longitude: 190 },
      }),
    ).toThrow(/longitude/);
    expect(() =>
      decodeGeocodeResult({
        ...rawGeocode,
        selected: { ...rawGeocode.selected, precision: "parcel" },
      }),
    ).toThrow(/address_point, interpolated/);
  });

  it("decodes a composed address lookup", () => {
    const result = decodeLookupResult({
      geocode: rawGeocode,
      resolution: rawResolution,
    });
    expect(result.geocode.selected?.matchedAddress).toContain("Broadway");
    expect(result.resolution.geography.stateName).toBe("Florida");
  });
});
