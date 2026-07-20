import { describe, expect, it } from "vitest";
import { buildApiUrl, decodeResolutionResult, getApiBaseUrl } from "./api";

describe("api url helpers", () => {
  it("defaults to the dev proxy path", () => {
    expect(getApiBaseUrl()).toBe("/api");
    expect(buildApiUrl("/health")).toBe("/api/health");
  });

  it("normalizes missing leading slashes", () => {
    expect(buildApiUrl("ready")).toBe("/api/ready");
  });
});

describe("resolution response decoding", () => {
  it("maps the versioned snake-case API contract without inventing missing records", () => {
    const result = decodeResolutionResult({
      schema_version: "1.0",
      generated_at: "2026-07-20T00:00:00Z",
      profile_id: "state-profile:us-co",
      profile_last_verified: "2026-07-20",
      geography: {
        state_id: "US-CO",
        state_fips: "08",
        state_name: "Colorado",
        county: {
          layer_family: "counties",
          feature_id: "08031",
          name: "Denver County",
        },
        municipality: {
          layer_family: "municipalities",
          feature_id: "0820000",
          name: "Denver",
        },
        incorporated: true,
      },
      code_family: "building",
      status: "local_record_required",
      required_local_records: ["Current municipal building-code adoption ordinance"],
      warnings: ["State data does not establish a local edition."],
      evidence: [],
    });

    expect(result.geography.municipality?.name).toBe("Denver");
    expect(result.status).toBe("local_record_required");
    expect(result.adoptions).toEqual([]);
    expect(result.requiredLocalRecords).toEqual([
      "Current municipal building-code adoption ordinance",
    ]);
  });
});
