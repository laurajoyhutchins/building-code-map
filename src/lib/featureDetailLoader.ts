import type { FeatureRecord, FeatureRef, LayerFamilyKey } from "../types";
import { encodeFeatureRef } from "./featureIdentity";

export interface FeatureDetailState {
  feature: FeatureRecord | null;
  isLoading: boolean;
  error: string | null;
}

export type FeatureDetailFetcher = (
  layerFamily: LayerFamilyKey,
  featureId: string,
  signal?: AbortSignal,
) => Promise<FeatureRecord>;

type PublishFeatureDetail = (state: FeatureDetailState) => void;

export function createFeatureDetailLoader(fetchDetail: FeatureDetailFetcher) {
  const cache = new Map<string, FeatureRecord>();
  let generation = 0;
  let activeController: AbortController | null = null;

  return {
    async load(ref: FeatureRef, publish: PublishFeatureDetail): Promise<void> {
      const key = encodeFeatureRef(ref);
      const cached = cache.get(key);

      activeController?.abort();
      activeController = null;
      const requestGeneration = ++generation;

      if (cached) {
        publish({ feature: cached, isLoading: false, error: null });
        return;
      }

      const controller = new AbortController();
      activeController = controller;
      publish({ feature: null, isLoading: true, error: null });

      try {
        const feature = await fetchDetail(ref.layerFamily, ref.featureId, controller.signal);
        if (controller.signal.aborted || requestGeneration !== generation) {
          return;
        }
        cache.set(key, feature);
        activeController = null;
        publish({ feature, isLoading: false, error: null });
      } catch {
        if (controller.signal.aborted || requestGeneration !== generation) {
          return;
        }
        activeController = null;
        publish({
          feature: null,
          isLoading: false,
          error: "Feature details could not be loaded.",
        });
      }
    },

    cancel(): void {
      generation += 1;
      activeController?.abort();
      activeController = null;
    },
  };
}
