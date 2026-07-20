import { useState, type FormEvent, type ReactNode } from "react";
import { fetchResolution } from "../lib/api";
import type { ResolutionResult, ResolutionStatus } from "../types";

const codeFamilies = [
  ["building", "Building"],
  ["residential", "Residential"],
  ["electrical", "Electrical"],
  ["energy", "Energy"],
  ["fire_operational", "Operational fire"],
] as const;

const projectTypes = [
  ["", "Ordinary project"],
  ["state_owned", "State-owned building"],
  ["public_school", "Public school"],
] as const;

const statusLabels: Record<ResolutionStatus, string> = {
  resolved: "Resolved",
  partially_resolved: "Partially resolved",
  local_record_required: "Local record required",
  ambiguous: "Ambiguous",
  conflicting: "Conflicting evidence",
  insufficient_evidence: "Insufficient evidence",
};

export function ResolutionPanel(): JSX.Element {
  const [longitude, setLongitude] = useState("-104.9903");
  const [latitude, setLatitude] = useState("39.7392");
  const [codeFamily, setCodeFamily] = useState("building");
  const [projectType, setProjectType] = useState("");
  const [result, setResult] = useState<ResolutionResult | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();
    const nextLongitude = Number(longitude);
    const nextLatitude = Number(latitude);
    if (!Number.isFinite(nextLongitude) || !Number.isFinite(nextLatitude)) {
      setError("Enter numeric longitude and latitude values.");
      return;
    }

    setIsLoading(true);
    setError(null);
    try {
      const nextResult = await fetchResolution({
        longitude: nextLongitude,
        latitude: nextLatitude,
        codeFamily,
        projectType: projectType || undefined,
      });
      setResult(nextResult);
    } catch (requestError) {
      setResult(null);
      setError(
        requestError instanceof Error
          ? requestError.message
          : "Resolution failed.",
      );
    } finally {
      setIsLoading(false);
    }
  }

  return (
    <section
      className="panel panel--tight resolution-panel"
      aria-labelledby="resolution-heading"
    >
      <div className="panel__header">
        <h2 id="resolution-heading">Resolve authority</h2>
        <p>
          Enter a coordinate to identify the state policy, likely authorities,
          and records that still require local verification.
        </p>
      </div>

      <form className="resolution-form" onSubmit={handleSubmit}>
        <div className="resolution-form__coordinates">
          <label>
            <span>Longitude</span>
            <input
              inputMode="decimal"
              value={longitude}
              onChange={(event) => setLongitude(event.target.value)}
              required
            />
          </label>
          <label>
            <span>Latitude</span>
            <input
              inputMode="decimal"
              value={latitude}
              onChange={(event) => setLatitude(event.target.value)}
              required
            />
          </label>
        </div>
        <label>
          <span>Code family</span>
          <select
            value={codeFamily}
            onChange={(event) => setCodeFamily(event.target.value)}
          >
            {codeFamilies.map(([value, label]) => (
              <option key={value} value={value}>
                {label}
              </option>
            ))}
          </select>
        </label>
        <label>
          <span>Project context</span>
          <select
            value={projectType}
            onChange={(event) => setProjectType(event.target.value)}
          >
            {projectTypes.map(([value, label]) => (
              <option key={value || "ordinary"} value={value}>
                {label}
              </option>
            ))}
          </select>
        </label>
        <button
          className="resolution-form__submit"
          type="submit"
          disabled={isLoading}
        >
          {isLoading ? "Resolving…" : "Resolve location"}
        </button>
      </form>

      {error ? (
        <p className="resolution-message resolution-message--error">{error}</p>
      ) : null}
      {result ? <ResolutionSummary result={result} /> : null}
    </section>
  );
}

function ResolutionSummary({
  result,
}: {
  result: ResolutionResult;
}): JSX.Element {
  const place =
    result.geography.municipality?.name ??
    result.geography.county?.name ??
    result.geography.stateName ??
    "Matched location";

  return (
    <div className="resolution-result" aria-live="polite">
      <div className="resolution-result__heading">
        <div>
          <span className="resolution-result__place">{place}</span>
          <strong>
            {result.codeFamily
              ? result.codeFamily.replace(/_/g, " ")
              : "All code families"}
          </strong>
        </div>
        <span
          className={`resolution-status resolution-status--${result.status}`}
        >
          {statusLabels[result.status]}
        </span>
      </div>

      <ResultSection title="Authority candidates">
        {result.authorityCandidates.length > 0 ? (
          <ul>
            {result.authorityCandidates.map((candidate) => (
              <li
                key={`${candidate.kind}:${candidate.authorityId ?? candidate.name}`}
              >
                <strong>{candidate.name}</strong>
                <span>
                  {candidate.roles.join(", ") ||
                    candidate.kind.replace(/_/g, " ")}
                </span>
              </li>
            ))}
          </ul>
        ) : (
          <p>No supported authority candidate is available.</p>
        )}
      </ResultSection>

      {result.adoptions.length > 0 ? (
        <ResultSection title="Supported statewide records">
          <ul>
            {result.adoptions.map((adoption) => (
              <li key={adoption.id}>
                <strong>{adoption.stateCodeName}</strong>
                <span>{adoption.codeFamily.replace(/_/g, " ")}</span>
              </li>
            ))}
          </ul>
        </ResultSection>
      ) : null}

      {result.requiredLocalRecords.length > 0 ? (
        <ResultSection title="Still verify locally">
          <ul>
            {result.requiredLocalRecords.map((record) => (
              <li key={record}>{record}</li>
            ))}
          </ul>
        </ResultSection>
      ) : null}

      {result.warnings.length > 0 ? (
        <div className="resolution-message resolution-message--warning">
          {result.warnings.map((warning) => (
            <p key={warning}>{warning}</p>
          ))}
        </div>
      ) : null}

      <details className="resolution-evidence">
        <summary>Evidence ({result.evidence.length})</summary>
        <ul>
          {result.evidence.map((source) => (
            <li key={source.id}>
              <a href={source.url} target="_blank" rel="noreferrer">
                {source.title}
              </a>
              <span>Accessed {source.accessedAt}</span>
            </li>
          ))}
        </ul>
      </details>
    </div>
  );
}

function ResultSection({
  title,
  children,
}: {
  title: string;
  children: ReactNode;
}): JSX.Element {
  return (
    <section className="resolution-result__section">
      <h3>{title}</h3>
      {children}
    </section>
  );
}
