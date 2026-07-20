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
  it("maps the complete versioned provenance contract without inventing missing records", () => {
    const result = decodeResolutionResult({
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
      status: "partially_resolved",
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
          kind: "date",
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
          availability: "available",
        },
      ],
    });

    expect(result.geography.municipality?.name).toBe("Orlando");
    expect(result.applicabilityDate).toBe("2026-07-20");
    expect(result.authorityPath[0]?.relationship).toBe("delegates_enforcement_to");
    expect(result.applicableRules[0]?.kind).toBe("date");
    expect(result.supportingClaims[0]?.value).toBe("2023-12-31");
    expect(result.evidence[0]?.availability).toBe("available");
    expect(result.adoptions).toEqual([]);
    expect(result.requiredLocalRecords).toEqual([]);
  });
});
