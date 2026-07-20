import type { LayerFamilyDefinition, LayerSelectionMap } from "../types";

interface LayerToggleListProps {
  layers: LayerFamilyDefinition[];
  enabledLayers: LayerSelectionMap;
  onToggle: (key: LayerFamilyDefinition["key"]) => void;
}

export function LayerToggleList({
  layers,
  enabledLayers,
  onToggle,
}: LayerToggleListProps): JSX.Element {
  return (
    <section className="panel panel--tight">
      <div className="panel__header">
        <h2>Boundary families</h2>
        <p>Independent layer toggles for the mirrored TIGERweb and NERIS feeds.</p>
      </div>
      <div className="toggle-list">
        {layers.map((layer) => {
          const active = Boolean(enabledLayers[layer.key]);

          return (
            <label className={`toggle-row ${active ? "is-active" : ""}`} key={layer.key}>
              <span className="toggle-row__copy">
                <strong>{layer.label}</strong>
                <span>{layer.description}</span>
              </span>
              <input
                type="checkbox"
                checked={active}
                onChange={() => onToggle(layer.key)}
                aria-label={`Toggle ${layer.label}`}
              />
            </label>
          );
        })}
      </div>
    </section>
  );
}
