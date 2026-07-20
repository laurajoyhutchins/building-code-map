import type { FeatureRecord, FeatureSummary } from "../types";

interface FeatureInspectorProps {
  feature: FeatureRecord | null;
  selectedFeature: FeatureSummary | null;
  isLoading: boolean;
  error: string | null;
}

function formatValue(value: unknown) {
  if (value === null || value === undefined || value === "") {
    return <span className="is-muted">Not populated</span>;
  }

  if (typeof value === "boolean") {
    return value ? "true" : "false";
  }

  return String(value);
}

function formatGeometrySource(value: string | null | undefined) {
  if (!value) {
    return <span className="is-muted">Not populated</span>;
  }

  if (value === "neris_department_jurisdiction") {
    return "NERIS DepartmentJurisdiction";
  }

  if (value === "county_approximation") {
    return "County approximation";
  }

  if (value === "tigerweb_live") {
    return "TIGERweb live";
  }

  return value;
}

export function FeatureInspector({
  feature,
  selectedFeature,
  isLoading,
  error,
}: FeatureInspectorProps): JSX.Element {
  const activeSummary = selectedFeature ?? feature;

  if (isLoading) {
    return (
      <aside className="panel inspector">
        <div className="panel__header">
          <h2>Feature record</h2>
          <p>Complete raw boundary fields keyed by the selected feature ID.</p>
        </div>
        <div className="empty-state">
          <h3>{activeSummary?.title ?? "Loading feature record"}</h3>
          <p>Waiting for the cached boundary record to load.</p>
        </div>
      </aside>
    );
  }

  if (error) {
    return (
      <aside className="panel inspector">
        <div className="panel__header">
          <h2>Feature record</h2>
          <p>Complete raw boundary fields keyed by the selected feature ID.</p>
        </div>
        <div className="empty-state">
          <h3>{activeSummary?.title ?? "Feature unavailable"}</h3>
          <p>{error}</p>
        </div>
      </aside>
    );
  }

  return (
    <aside className="panel inspector">
      <div className="panel__header">
        <h2>Feature record</h2>
        <p>Complete raw boundary fields keyed by the selected feature ID.</p>
      </div>

      {feature ? (
        <>
          <div className="inspector__summary">
            <div>
              <span className="eyebrow">{feature.layerFamily}</span>
              <h3>{feature.title}</h3>
              <p>{feature.subtitle}</p>
            </div>
            <dl className="meta-grid">
              <div>
                <dt>Feature ID</dt>
                <dd>{feature.featureId}</dd>
              </div>
              <div>
                <dt>Source ID</dt>
                <dd>{feature.sourceId}</dd>
              </div>
              <div>
                <dt>Geometry</dt>
                <dd>{feature.geometryLabel}</dd>
              </div>
              <div>
                <dt>Geometry source</dt>
                <dd>{formatGeometrySource(feature.geometrySource)}</dd>
              </div>
              <div>
                <dt>Last synced</dt>
                <dd>{feature.lastSyncedAt}</dd>
              </div>
            </dl>
          </div>

          <div className="attribute-table">
            {Object.entries(feature.attributes).map(([key, value]) => (
              <div className="attribute-row" key={key}>
                <span className="attribute-row__key">{key}</span>
                <span className="attribute-row__value">{formatValue(value)}</span>
              </div>
            ))}
          </div>
        </>
      ) : (
        <div className="empty-state">
          <h3>{activeSummary?.title ?? "No feature selected"}</h3>
          <p>
            {activeSummary
              ? "Waiting for the cached record to load."
              : "Click a boundary in the map stage or select a cached feature to inspect its record."}
          </p>
        </div>
      )}
    </aside>
  );
}
