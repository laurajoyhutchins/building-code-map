import { useEffect, useRef, useState } from "react";
import type { GeoJSONSource, Map as MapLibreMap, StyleSpecification } from "maplibre-gl";
import {
  encodeFeatureRef,
  featureRefFromMapProperties,
  featureRefsEqual,
} from "../lib/featureIdentity";
import { buildBoundaryFeatureCollection, calculateBoundaryBounds } from "../lib/mapData";
import type {
  BoundaryFeatureRecord,
  FeatureRef,
  FeatureSummary,
  LayerFamilyDefinition,
  LayerSelectionMap,
  RefreshStatus,
} from "../types";

interface MapStageProps {
  layers: LayerFamilyDefinition[];
  selectedFeature: BoundaryFeatureRecord | null;
  enabledLayers: LayerSelectionMap;
  onSelectFeature: (feature: FeatureRef) => void;
  refreshStatus: RefreshStatus;
  featureSummaries: FeatureSummary[];
  boundaryFeatures: readonly BoundaryFeatureRecord[];
}

const sourceId = "tigerweb-boundaries";

const familyStyling: Record<
  LayerFamilyDefinition["key"],
  { fillColor: string; lineColor: string; selectedColor: string }
> = {
  states: {
    fillColor: "rgba(31, 78, 121, 0.26)",
    lineColor: "#1f4e79",
    selectedColor: "#10263f",
  },
  counties: {
    fillColor: "rgba(47, 109, 85, 0.22)",
    lineColor: "#2f6d55",
    selectedColor: "#183927",
  },
  municipalities: {
    fillColor: "rgba(161, 92, 31, 0.22)",
    lineColor: "#a15c1f",
    selectedColor: "#67390f",
  },
  special_areas: {
    fillColor: "rgba(123, 78, 29, 0.18)",
    lineColor: "#7b4e1d",
    selectedColor: "#4d2f10",
  },
  tribal_areas: {
    fillColor: "rgba(14, 107, 109, 0.22)",
    lineColor: "#0e6b6d",
    selectedColor: "#083f40",
  },
  neris_jurisdictions: {
    fillColor: "rgba(72, 93, 126, 0.22)",
    lineColor: "#485d7e",
    selectedColor: "#2f3d53",
  },
};

function getLayerPaints(layerKey: LayerFamilyDefinition["key"]) {
  const styling = familyStyling[layerKey];

  return {
    fillColor: styling.fillColor,
    lineColor: styling.lineColor,
    selectedColor: styling.selectedColor,
  };
}

function createMapStyle() {
  return {
    version: 8,
    sources: {
      osm: {
        type: "raster",
        tiles: ["https://tile.openstreetmap.org/{z}/{x}/{y}.png"],
        tileSize: 256,
        attribution: "© OpenStreetMap contributors",
      },
    },
    layers: [
      {
        id: "osm-basemap",
        type: "raster",
        source: "osm",
      },
    ],
  } satisfies StyleSpecification;
}

function layerIds(layerKey: LayerFamilyDefinition["key"]) {
  return {
    fill: `${sourceId}-${layerKey}-fill`,
    line: `${sourceId}-${layerKey}-line`,
    selected: `${sourceId}-${layerKey}-selected`,
  };
}

export function reconcileBoundaryLayers(
  map: MapLibreMap,
  layers: readonly LayerFamilyDefinition[],
  enabledLayers: LayerSelectionMap,
  boundaryCollection: ReturnType<typeof buildBoundaryFeatureCollection>,
): void {
  if (!map.getSource(sourceId)) {
    map.addSource(sourceId, {
      type: "geojson",
      data: boundaryCollection,
    });
  }

  const source = map.getSource(sourceId) as GeoJSONSource | undefined;
  source?.setData(boundaryCollection);

  for (const layer of layers) {
    const ids = layerIds(layer.key);
    const { fillColor, lineColor, selectedColor } = getLayerPaints(layer.key);
    const visibility = enabledLayers[layer.key] ? "visible" : "none";

    if (!map.getLayer(ids.fill)) {
      map.addLayer({
        id: ids.fill,
        type: "fill",
        source: sourceId,
        filter: ["==", ["get", "layerFamily"], layer.key],
        paint: {
          "fill-color": fillColor,
          "fill-opacity": 0.95,
          "fill-outline-color": lineColor,
        },
        layout: { visibility },
      });
    }

    if (!map.getLayer(ids.line)) {
      map.addLayer({
        id: ids.line,
        type: "line",
        source: sourceId,
        filter: ["==", ["get", "layerFamily"], layer.key],
        paint: {
          "line-color": lineColor,
          "line-width": 1.5,
        },
        layout: { visibility },
      });
    }

    if (!map.getLayer(ids.selected)) {
      map.addLayer({
        id: ids.selected,
        type: "line",
        source: sourceId,
        filter: [
          "all",
          ["==", ["get", "layerFamily"], layer.key],
          ["==", ["get", "selected"], true],
        ],
        paint: {
          "line-color": selectedColor,
          "line-width": 3.5,
        },
        layout: { visibility },
      });
    }

    for (const id of [ids.fill, ids.line, ids.selected]) {
      if (map.getLayer(id)) {
        map.setLayoutProperty(id, "visibility", visibility);
      }
    }
  }
}

export function MapStage({
  layers,
  selectedFeature,
  enabledLayers,
  onSelectFeature,
  refreshStatus,
  featureSummaries,
  boundaryFeatures,
}: MapStageProps): JSX.Element {
  const mapContainerRef = useRef<HTMLDivElement | null>(null);
  const mapRef = useRef<MapLibreMap | null>(null);
  const previousSelectedFeatureKeyRef = useRef<string | null>(null);
  const [mapStatus, setMapStatus] = useState<"loading" | "ready" | "error">("loading");
  const [mapError, setMapError] = useState<string | null>(null);

  const boundaryCollection = buildBoundaryFeatureCollection(
    boundaryFeatures,
    enabledLayers,
    selectedFeature,
  );
  const layersRef = useRef(layers);
  const enabledLayersRef = useRef(enabledLayers);
  const boundaryCollectionRef = useRef(boundaryCollection);
  const onSelectFeatureRef = useRef(onSelectFeature);

  useEffect(() => {
    layersRef.current = layers;
  }, [layers]);

  useEffect(() => {
    enabledLayersRef.current = enabledLayers;
  }, [enabledLayers]);

  useEffect(() => {
    boundaryCollectionRef.current = boundaryCollection;
  }, [boundaryCollection]);

  useEffect(() => {
    onSelectFeatureRef.current = onSelectFeature;
  }, [onSelectFeature]);

  useEffect(() => {
    if (!mapContainerRef.current || mapRef.current) {
      return;
    }

    let cancelled = false;

    void import("maplibre-gl")
      .then(({ default: maplibregl }) => {
        if (cancelled || !mapContainerRef.current) {
          return;
        }

        const map = new maplibregl.Map({
          container: mapContainerRef.current,
          style: createMapStyle(),
          center: [-105.4, 39.2],
          zoom: 5.35,
          pitch: 0,
          attributionControl: { compact: true },
        });

        mapRef.current = map;
        map.addControl(new maplibregl.NavigationControl({ showCompass: false }), "top-right");

        map.on("load", () => {
          if (cancelled) {
            return;
          }

          reconcileBoundaryLayers(
            map,
            layersRef.current,
            enabledLayersRef.current,
            boundaryCollectionRef.current,
          );

          map.on("click", (event) => {
            const existingFillLayers = layersRef.current
              .map((layer) => layerIds(layer.key).fill)
              .filter((id) => Boolean(map.getLayer(id)));
            if (existingFillLayers.length === 0) {
              return;
            }

            const renderedFeatures = map.queryRenderedFeatures(event.point, {
              layers: existingFillLayers,
            });
            const clickedFeature = featureRefFromMapProperties(renderedFeatures[0]?.properties);

            if (clickedFeature) {
              onSelectFeatureRef.current(clickedFeature);
            }
          });

          setMapError(null);
          setMapStatus("ready");
        });

        map.on("error", (event) => {
          if (cancelled) {
            return;
          }

          const message =
            event.error instanceof Error
              ? event.error.message
              : "MapLibre failed to load the OpenStreetMap basemap.";

          setMapError(message);
          if (!map.getSource(sourceId)) {
            setMapStatus("error");
          }
        });
      })
      .catch((error: unknown) => {
        if (cancelled) {
          return;
        }

        setMapError(error instanceof Error ? error.message : "MapLibre could not be initialized.");
        setMapStatus("error");
      });

    return () => {
      cancelled = true;
      mapRef.current?.remove();
      mapRef.current = null;
    };
  }, []);

  useEffect(() => {
    const map = mapRef.current;
    if (!map || mapStatus !== "ready") {
      return;
    }

    reconcileBoundaryLayers(map, layers, enabledLayers, boundaryCollection);

    if (!selectedFeature) {
      previousSelectedFeatureKeyRef.current = null;
      return;
    }

    const selectedFeatureKey = encodeFeatureRef(selectedFeature);
    if (selectedFeatureKey !== previousSelectedFeatureKeyRef.current) {
      previousSelectedFeatureKeyRef.current = selectedFeatureKey;
      map.fitBounds(calculateBoundaryBounds(selectedFeature.geometry), {
        padding: 48,
        duration: 650,
        maxZoom: 9,
      });
    }
  }, [boundaryCollection, enabledLayers, layers, mapStatus, selectedFeature]);

  return (
    <section className="map-stage">
      <div className="panel panel--map">
        <div className="map-stage__topline">
          <div>
            <p className="eyebrow">OpenStreetMap basemap</p>
            <h2>TIGERweb and NERIS boundary explorer</h2>
            <p className="lede">
              MapLibre renders cached boundary layers and keeps the selected feature highlighted
              while the inspector loads the full backend record.
            </p>
          </div>

          <div className={`status-chip status-chip--${refreshStatus.status}`}>
            {refreshStatus.status === "ok" ? "Mirror healthy" : "Mirror attention needed"}
          </div>
        </div>

        <div className="map-canvas">
          <div className="map-canvas__surface">
            <div ref={mapContainerRef} className="map-canvas__map" />

            <div className="map-canvas__overlay">
              <div className="map-canvas__label">
                {mapStatus === "ready"
                  ? "MapLibre cached map"
                  : mapStatus === "error"
                    ? "MapLibre unavailable"
                    : "Loading map"}
              </div>

              <div className="map-canvas__layers">
                {layers.map((layer) => (
                  <span
                    className={`layer-pill ${enabledLayers[layer.key] ? "is-enabled" : "is-disabled"}`}
                    key={layer.key}
                  >
                    {layer.label}
                  </span>
                ))}
              </div>

              <p className="map-canvas__hint">
                {mapStatus === "error"
                  ? (mapError ?? "The map surface could not be initialized.")
                  : selectedFeature
                    ? `Selected: ${selectedFeature.title}`
                    : "Click a boundary to inspect the cached attribute record."}
              </p>
            </div>
          </div>
        </div>
      </div>

      <div className="panel panel--tight">
        <div className="panel__header">
          <h2>Cached click targets</h2>
          <p>Local selection list derived from the runtime snapshot.</p>
        </div>
        <div className="feature-list">
          {featureSummaries.map((feature) => {
            const active = featureRefsEqual(selectedFeature, feature);

            return (
              <button
                className={`feature-card ${active ? "is-active" : ""}`}
                key={encodeFeatureRef(feature)}
                onClick={() =>
                  onSelectFeature({
                    layerFamily: feature.layerFamily,
                    featureId: feature.featureId,
                  })
                }
                type="button"
              >
                <strong>{feature.title}</strong>
                <span>{feature.subtitle}</span>
                <small>
                  {feature.layerFamily} {feature.sourceId}
                </small>
              </button>
            );
          })}
        </div>
      </div>

      <div className="panel panel--tight">
        <div className="panel__header">
          <h2>Backend contract</h2>
          <p>Endpoints exposed by the Go snapshot service.</p>
        </div>
        <dl className="meta-grid meta-grid--stack">
          <div>
            <dt>Health</dt>
            <dd>/health</dd>
          </div>
          <div>
            <dt>Readiness</dt>
            <dd>/ready</dd>
          </div>
          <div>
            <dt>Layer registry</dt>
            <dd>/layers</dd>
          </div>
          <div>
            <dt>Feature lookup</dt>
            <dd>{"/features/{layer_family}/{feature_id}"}</dd>
          </div>
          <div>
            <dt>Boundary list</dt>
            <dd>/boundaries</dd>
          </div>
          <div>
            <dt>Refresh status</dt>
            <dd>/refresh/status</dd>
          </div>
        </dl>
      </div>
    </section>
  );
}
