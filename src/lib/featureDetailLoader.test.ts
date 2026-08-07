import { describe, expect, it, vi } from "vitest";
import type { FeatureRecord, FeatureRef } from "../types";
import { createFeatureDetailLoader, type FeatureDetailState } from "./featureDetailLoader";

const coloradoRef: FeatureRef = { layerFamily: "states", featureId: "08" };
const denverRef: FeatureRef = { layerFamily: "counties", featureId: "08031" };

function feature(ref: FeatureRef, title: string): FeatureRecord {
  return {
    ...ref,
    title,
    subtitle: "Boundary",
    sourceId: `source:${ref.featureId}`,
    geometryLabel: "Polygon",
    geometrySource: "tigerweb_live",
    lastSyncedAt: "2026-06-22T12:00:00Z",
    attributes: { NAME: title },
  };
}

function deferred<T>() {
  let resolve!: (value: T) => void;
  let reject!: (reason?: unknown) => void;
  const promise = new Promise<T>((nextResolve, nextReject) => {
    resolve = nextResolve;
    reject = nextReject;
  });
  return { promise, resolve, reject };
}

function latest(states: FeatureDetailState[]): FeatureDetailState {
  const state = states[states.length - 1];
  if (!state) throw new Error("expected at least one published state");
  return state;
}

describe("feature detail loader", () => {
  it("publishes loading before a successful detail response", async () => {
    const pending = deferred<FeatureRecord>();
    const fetchDetail = vi.fn(() => pending.promise);
    const loader = createFeatureDetailLoader(fetchDetail);
    const states: FeatureDetailState[] = [];

    const loading = loader.load(coloradoRef, (state) => states.push(state));

    expect(latest(states)).toEqual({ feature: null, isLoading: true, error: null });

    const colorado = feature(coloradoRef, "Colorado");
    pending.resolve(colorado);
    await loading;

    expect(latest(states)).toEqual({ feature: colorado, isLoading: false, error: null });
  });

  it("publishes a bounded error when detail loading fails", async () => {
    const fetchDetail = vi.fn(async () => {
      throw new Error("backend unavailable");
    });
    const loader = createFeatureDetailLoader(fetchDetail);
    const states: FeatureDetailState[] = [];

    await loader.load(coloradoRef, (state) => states.push(state));

    expect(latest(states)).toEqual({
      feature: null,
      isLoading: false,
      error: "Feature details could not be loaded.",
    });
  });

  it("never publishes a stale response after selection changes", async () => {
    const coloradoPending = deferred<FeatureRecord>();
    const denverPending = deferred<FeatureRecord>();
    const fetchDetail = vi.fn((_: FeatureRef["layerFamily"], featureId: string) =>
      featureId === coloradoRef.featureId ? coloradoPending.promise : denverPending.promise,
    );
    const loader = createFeatureDetailLoader(fetchDetail);
    const states: FeatureDetailState[] = [];

    const firstLoad = loader.load(coloradoRef, (state) => states.push(state));
    const secondLoad = loader.load(denverRef, (state) => states.push(state));

    coloradoPending.resolve(feature(coloradoRef, "Colorado"));
    await firstLoad;
    expect(states.some((state) => state.feature?.featureId === coloradoRef.featureId)).toBe(false);

    const denver = feature(denverRef, "Denver County");
    denverPending.resolve(denver);
    await secondLoad;

    expect(latest(states)).toEqual({ feature: denver, isLoading: false, error: null });
  });

  it("serves successful repeat selections from cache", async () => {
    const colorado = feature(coloradoRef, "Colorado");
    const fetchDetail = vi.fn(async () => colorado);
    const loader = createFeatureDetailLoader(fetchDetail);
    const states: FeatureDetailState[] = [];

    await loader.load(coloradoRef, (state) => states.push(state));
    await loader.load(coloradoRef, (state) => states.push(state));

    expect(fetchDetail).toHaveBeenCalledTimes(1);
    expect(latest(states)).toEqual({ feature: colorado, isLoading: false, error: null });
  });
});
