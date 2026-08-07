import { useEffect, useRef, useState } from "react";
import { FeatureInspector } from "./components/FeatureInspector";
import { LayerToggleList } from "./components/LayerToggleList";
import { MapStage } from "./components/MapStage";
import { PublicLookup } from "./components/PublicLookup";
import { ResolutionPanel } from "./components/ResolutionPanel";
import { StatusBanner } from "./components/StatusBanner";
import {
  fetchBoundaryFeatures,
  fetchFeature,
  fetchLayers,
  fetchRefreshStatus,
  getApiBaseUrl,
} from "./lib/api";
import { createFeatureDetailLoader } from "./lib/featureDetailLoader";
import { findFeatureByRef } from "./lib/featureIdentity";
import { createLayerSelectionMap } from "./lib/layerSelection";
import type {
  BoundaryFeatureRecord,
  FeatureRecord,
  FeatureRef,
  LayerFamilyDefinition,
  LayerFamilyKey,
  LayerSelectionMap,
  RefreshStatus,
} from "./types";

const loadingRefreshStatus: RefreshStatus = {
  status: "warning",
  latestSuccessfulRefresh: null,
  latestAttempt: new Date(0).toISOString(),
  nextScheduledRefresh: new Date(0).toISOString(),
  message: "Loading cached boundary snapshot.",
};

function App(): JSX.Element {
  const routeName = window.location.pathname.split("/").filter(Boolean).at(-1);
  return routeName === "explorer" ? <ExplorerApp /> : <PublicLookup />;
}

function ExplorerApp(): JSX.Element {
  const [layerRegistry, setLayerRegistry] = useState<LayerFamilyDefinition[]>([]);
  const [enabledLayers, setEnabledLayers] = useState<LayerSelectionMap>(() =>
    createLayerSelectionMap([]),
  );
  const [refreshStatus, setRefreshStatus] = useState<RefreshStatus | null>(null);
  const [boundaryFeatures, setBoundaryFeatures] = useState<BoundaryFeatureRecord[]>([]);
  const [selectedFeatureRef, setSelectedFeatureRef] = useState<FeatureRef | null>(null);
  const [selectedFeature, setSelectedFeature] = useState<FeatureRecord | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [loadError, setLoadError] = useState<string | null>(null);
  const [isDetailLoading, setIsDetailLoading] = useState(false);
  const [detailError, setDetailError] = useState<string | null>(null);
  const detailLoaderRef = useRef<ReturnType<typeof createFeatureDetailLoader> | null>(null);

  if (!detailLoaderRef.current) {
    detailLoaderRef.current = createFeatureDetailLoader(fetchFeature);
  }

  const selectedMapFeature = findFeatureByRef(boundaryFeatures, selectedFeatureRef);
  const selectedSummary = selectedMapFeature;

  useEffect(() => {
    let cancelled = false;

    async function loadMetadata() {
      try {
        const [nextLayers, nextRefreshStatus, nextBoundaryFeatures] = await Promise.all([
          fetchLayers(),
          fetchRefreshStatus(),
          fetchBoundaryFeatures(),
        ]);

        if (cancelled) {
          return;
        }

        const firstFeature = nextBoundaryFeatures[0];
        const firstFeatureRef: FeatureRef | null = firstFeature
          ? {
              layerFamily: firstFeature.layerFamily,
              featureId: firstFeature.featureId,
            }
          : null;

        setLayerRegistry(nextLayers);
        setEnabledLayers(createLayerSelectionMap(nextLayers));
        setRefreshStatus(nextRefreshStatus);
        setBoundaryFeatures(nextBoundaryFeatures);
        setLoadError(null);
        setSelectedFeatureRef((current) => current ?? firstFeatureRef);
      } catch {
        if (!cancelled) {
          setLayerRegistry([]);
          setEnabledLayers(createLayerSelectionMap([]));
          setRefreshStatus(null);
          setBoundaryFeatures([]);
          setSelectedFeatureRef(null);
          setLoadError("Cached boundary data could not be loaded.");
        }
      } finally {
        if (!cancelled) {
          setIsLoading(false);
        }
      }
    }

    void loadMetadata();

    return () => {
      cancelled = true;
    };
  }, []);

  useEffect(() => {
    const loader = detailLoaderRef.current;
    if (!loader) {
      return;
    }

    if (!selectedFeatureRef) {
      loader.cancel();
      setSelectedFeature(null);
      setIsDetailLoading(false);
      setDetailError(null);
      return;
    }

    void loader.load(selectedFeatureRef, (state) => {
      setSelectedFeature(state.feature);
      setIsDetailLoading(state.isLoading);
      setDetailError(state.error);
    });

    return () => loader.cancel();
  }, [selectedFeatureRef]);

  const handleToggleLayer = (key: LayerFamilyKey) => {
    setEnabledLayers((current) => ({
      ...current,
      [key]: !current[key],
    }));
  };

  const apiBaseUrl = getApiBaseUrl();

  return (
    <div className="app-shell">
      <header className="app-header">
        <div>
          <p className="eyebrow">Local GIS + source-backed jurisdiction policy</p>
          <h1>Building Code Map Explorer</h1>
        </div>
        <div className="app-header__stats">
          <div>
            <span>Layer families</span>
            <strong>{layerRegistry.length}</strong>
          </div>
          <div>
            <span>Cached records</span>
            <strong>{boundaryFeatures.length}</strong>
          </div>
          <div>
            <span>Refresh state</span>
            <strong>{refreshStatus?.status ?? "loading"}</strong>
          </div>
        </div>
      </header>

      <main className="app-grid">
        <aside className="sidebar">
          <StatusBanner refreshStatus={refreshStatus} apiBaseUrl={apiBaseUrl} error={loadError} />
          <ResolutionPanel />
          <LayerToggleList
            layers={layerRegistry}
            enabledLayers={enabledLayers}
            onToggle={handleToggleLayer}
          />
        </aside>

        <MapStage
          layers={layerRegistry}
          selectedFeature={selectedMapFeature}
          enabledLayers={enabledLayers}
          onSelectFeature={(featureRef) => setSelectedFeatureRef(featureRef)}
          refreshStatus={refreshStatus ?? loadingRefreshStatus}
          featureSummaries={boundaryFeatures.map(
            ({ layerFamily, featureId, title, subtitle, sourceId }) => ({
              layerFamily,
              featureId,
              title,
              subtitle,
              sourceId,
            }),
          )}
          boundaryFeatures={boundaryFeatures}
        />

        <FeatureInspector
          feature={selectedFeature}
          selectedFeature={selectedSummary}
          isLoading={isLoading || isDetailLoading}
          error={detailError ?? loadError}
        />
      </main>
    </div>
  );
}

export default App;
