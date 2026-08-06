import { ApiResponseError } from "../lib/api";
import type { ReadinessResult } from "../types";

export function ApiEvidenceNotice({ error }: { error: unknown }): JSX.Element | null {
  if (!(error instanceof ApiResponseError)) {
    return error instanceof Error ? (
      <p className="public-message public-message--error" role="alert">
        {error.message}
      </p>
    ) : null;
  }

  if (error.code === "boundary_ambiguous" && error.details) {
    return (
      <section className="public-message public-message--error" role="alert">
        <strong>{error.message}</strong>
        <p>
          The point intersects multiple {error.details.layerFamily.replace(/_/g, " ")} records. No
          authority result was selected.
        </p>
        <ul>
          {error.details.observations.map((observation) => (
            <li key={`${observation.layerFamily}:${observation.featureId}`}>
              {observation.name} ({observation.featureId})
              {observation.sourceId ? ` · source ${observation.sourceId}` : ""}
            </li>
          ))}
        </ul>
      </section>
    );
  }

  return (
    <p className="public-message public-message--error" role="alert">
      {error.message}
    </p>
  );
}

export function ReadinessNotice({
  readiness,
}: {
  readiness: ReadinessResult | null;
}): JSX.Element | null {
  if (!readiness || readiness.readiness === "ready") {
    return null;
  }
  const unavailable = Object.entries(readiness.capabilities).filter(
    ([, capability]) => capability.status === "unavailable",
  );
  return (
    <section className="public-message" aria-live="polite">
      <strong>
        {readiness.readiness === "not_ready"
          ? "Lookup service not ready"
          : "Some workflows unavailable"}
      </strong>
      {unavailable.length > 0 ? (
        <ul>
          {unavailable.map(([name, capability]) => (
            <li key={name}>
              {name.replace(/_/g, " ")}: {capability.message}
            </li>
          ))}
        </ul>
      ) : null}
      {Object.entries(readiness.snapshots).map(([kind, snapshot]) => (
        <p key={kind}>
          {kind} snapshot: {snapshot.status}
          {snapshot.snapshotId ? ` (${snapshot.snapshotId})` : ""}
        </p>
      ))}
    </section>
  );
}
