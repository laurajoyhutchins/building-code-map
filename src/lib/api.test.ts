import { describe, expect, it } from "vitest";
import {
  buildApiUrl,
  decodeGeocodeResult,
  decodeLookupResult,
  decodeResolutionResult,
  getApiBaseUrl,
} from "./api";

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
      source: "Denver address points",
      source_record_id: "co-denver-1600",
      source_vintage: "2026-08-01",
    },
    candidates: [],
    warnings: [],
  };

  it("preserves selected-point provenance", () => {
    const result = decodeGeocodeResult(rawGeocode);
    expect(result.selected?.sourceRecordId).toBe("co-denver-1600");
    expect(result.selected?.precision).toBe("address_point");
    expect(result.selected?.sourceVintage).toBe("2026-08-01");
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
