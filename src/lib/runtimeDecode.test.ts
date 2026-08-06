import { describe, expect, it } from "vitest";
import {
  PayloadDecodeError,
  arrayValue,
  enumValue,
  httpUrl,
  isoDate,
  latitude,
  longitude,
  record,
} from "./runtimeDecode";

describe("runtime decoding primitives", () => {
  it("rejects arrays where an object is required", () => {
    expect(() => record([], "health")).toThrowError(
      new PayloadDecodeError("health", "expected an object"),
    );
  });

  it("validates coordinate ranges", () => {
    expect(longitude(-104.99, "point.longitude")).toBe(-104.99);
    expect(latitude(39.74, "point.latitude")).toBe(39.74);
    expect(() => longitude(181, "point.longitude")).toThrow(/-180 through 180/);
    expect(() => latitude(-91, "point.latitude")).toThrow(/-90 through 90/);
  });

  it("validates enum members", () => {
    expect(enumValue("ready", ["ready", "degraded"] as const, "ready.status")).toBe("ready");
    expect(() => enumValue("mystery", ["ready", "degraded"] as const, "ready.status")).toThrow(
      /ready, degraded/,
    );
  });

  it("identifies malformed array members by index", () => {
    expect(() => arrayValue([{}, null], "layers", (item, path) => record(item, path))).toThrow(
      /layers\[1\]/,
    );
  });

  it("validates dates and source URLs", () => {
    expect(isoDate("2026-08-05", "source.accessed_at")).toBe("2026-08-05");
    expect(httpUrl("https://example.com/code", "source.url")).toContain("example.com");
    expect(() => isoDate("August 5", "source.accessed_at")).toThrow(/ISO date/);
    expect(() => httpUrl("file:///tmp/code", "source.url")).toThrow(/HTTP or HTTPS/);
  });
});
