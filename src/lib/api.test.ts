import { describe, expect, it } from "vitest";
import { buildApiUrl, getApiBaseUrl } from "./api";

describe("api url helpers", () => {
  it("defaults to the dev proxy path", () => {
    expect(getApiBaseUrl()).toBe("/api");
    expect(buildApiUrl("/health")).toBe("/api/health");
  });

  it("normalizes missing leading slashes", () => {
    expect(buildApiUrl("ready")).toBe("/api/ready");
  });
});
