import { useState, type FormEvent, type ReactNode } from "react";
import { fetchResolution } from "../lib/api";
import {
  formatCodeFamily,
  getResolutionNotice,
  getResolutionPlace,
  parseCoordinateQuery,
} from "../lib/publicLookup";
import type { ResolutionResult } from "../types";

const codeFamilies = [
  ["building", "Building"],
  ["residential", "Residential"],
  ["electrical", "Electrical"],
  ["energy", "Energy"],
  ["fire_operational", "Operational fire"],
] as const;

const today = new Date().toISOString().slice(0, 10);

export function PublicLookup(): JSX.Element {
  const [coordinates, setCoordinates] = useState("");
  const [codeFamily, setCodeFamily] = useState("building");
  const [applicabilityDate, setApplicabilityDate] = useState(today);
  const [result, setResult] = useState<ResolutionResult | null>(null);
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);

  const explorerHref = new URL("explorer", window.location.href).pathname;

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();

    let point;
    try {
      point = parseCoordinateQuery(coordinates);
    } catch (coordinateError) {
      setError(
        coordinateError instanceof Error ? coordinateError.message : "Enter valid coordinates.",
      );
      return;
    }

    setIsLoading(true);
    setError(null);
    try {
      const nextResult = await fetchResolution({
        longitude: point.longitude,
        latitude: point.latitude,
        codeFamily,
        applicabilityDate,
      });
      setResult(nextResult);
    } catch (requestError) {
      setResult(null);
      setError(requestError instanceof Error ? requestError.message : "The lookup failed.");
    } finally {
      setIsLoading(false);
    }
  }

  return (
    <div className="public-shell">
      <header className="public-header">
        <a className="public-brand" href="./" aria-label="Building Code Map home">
          Building Code Map
        </a>
        <a className="public-header__link" href={explorerHref}>
          Explorer
        </a>
      </header>

      <main className="public-main">
        <section className="public-intro" aria-labelledby="public-title">
          <h1 id="public-title">Building Code Map</h1>
          <p>Find the building codes and authorities associated with a location.</p>
        </section>

        <form className="public-search" onSubmit={handleSubmit}>
          <label className="public-search__location">
            <span>Coordinates</span>
            <input
              autoComplete="off"
              inputMode="decimal"
              placeholder="39.7392, -104.9903"
              value={coordinates}
              onChange={(event) => setCoordinates(event.target.value)}
              required
            />
            <small>Enter latitude and longitude. The reverse order is also accepted.</small>
          </label>

          <div className="public-search__options">
            <label>
              <span>Code</span>
              <select value={codeFamily} onChange={(event) => setCodeFamily(event.target.value)}>
                {codeFamilies.map(([value, label]) => (
                  <option key={value} value={value}>
                    {label}
                  </option>
                ))}
              </select>
            </label>

            <label>
              <span>Date</span>
              <input
                type="date"
                value={applicabilityDate}
                onChange={(event) => setApplicabilityDate(event.target.value)}
                required
              />
            </label>

            <button type="submit" disabled={isLoading}>
              {isLoading ? "Searching…" : "Search"}
            </button>
          </div>
        </form>

        {error ? (
          <p className="public-message public-message--error" role="alert">
            {error}
          </p>
        ) : null}

        {result ? <PublicResult result={result} /> : null}
      </main>

      <footer className="public-footer">
        <p>
          Building Code Map summarizes public records and does not replace confirmation by the
          applicable authority.
        </p>
      </footer>
    </div>
  );
}

function PublicResult({ result }: { result: ResolutionResult }): JSX.Element {
  const notice = getResolutionNotice(result.status);
  const authorities = result.authorityCandidates;
  const adoptions = result.adoptions;

  return (
    <article className="public-result" aria-live="polite">
      <header className="public-result__header">
        <p>{formatCodeFamily(result.codeFamily)}</p>
        <h2>{getResolutionPlace(result)}</h2>
        {result.applicabilityDate ? <p>Applicable on {result.applicabilityDate}</p> : null}
        {notice ? <div className="public-message">{notice}</div> : null}
      </header>

      <ResultSection title="Authorities">
        {authorities.length > 0 ? (
          <ul className="public-record-list">
            {authorities.map((authority) => (
              <li key={`${authority.kind}:${authority.authorityId ?? authority.name}`}>
                <strong>{authority.name}</strong>
                {authority.roles.length > 0 ? <span>{authority.roles.join(", ")}</span> : null}
              </li>
            ))}
          </ul>
        ) : (
          <p>No supported authority record is available.</p>
        )}
      </ResultSection>

      <ResultSection title="Adopted codes">
        {adoptions.length > 0 ? (
          <ul className="public-record-list">
            {adoptions.map((adoption) => (
              <li key={adoption.id}>
                <strong>{adoption.stateCodeName}</strong>
                <span>{formatAdoptionDetails(adoption)}</span>
              </li>
            ))}
          </ul>
        ) : (
          <p>No supported adoption record is available.</p>
        )}
      </ResultSection>

      {result.applicableRules.length > 0 ? (
        <ResultSection title="Special conditions">
          <ul className="public-record-list">
            {result.applicableRules.map((rule) => (
              <li key={rule.id}>
                <strong>{formatRuleKind(rule.kind)}</strong>
                <span>{rule.summary}</span>
              </li>
            ))}
          </ul>
        </ResultSection>
      ) : null}

      {result.requiredLocalRecords.length > 0 || result.warnings.length > 0 ? (
        <ResultSection title="Confirm locally">
          <ul className="public-record-list public-record-list--plain">
            {[...result.requiredLocalRecords, ...result.warnings].map((item) => (
              <li key={item}>{item}</li>
            ))}
          </ul>
        </ResultSection>
      ) : null}

      {result.authorityPath.length > 0 ? (
        <details className="public-details">
          <summary>Jurisdiction structure</summary>
          <ul className="public-record-list">
            {result.authorityPath.map((relationship) => (
              <li key={relationship.id}>
                <strong>{relationship.summary ?? relationship.relationship.replace(/_/g, " ")}</strong>
                <span>
                  {relationship.fromId} to {relationship.to}
                </span>
              </li>
            ))}
          </ul>
        </details>
      ) : null}

      <ResultSection title="Sources">
        {result.evidence.length > 0 ? (
          <ul className="public-source-list">
            {result.evidence.map((source) => (
              <li key={source.id}>
                <a href={source.url} target="_blank" rel="noreferrer">
                  {source.title}
                </a>
                <span>
                  Accessed {source.accessedAt}
                  {source.lastCheckedAt ? ` · checked ${source.lastCheckedAt}` : ""}
                </span>
              </li>
            ))}
          </ul>
        ) : (
          <p>No source links are available for this result.</p>
        )}
      </ResultSection>
    </article>
  );
}

function formatAdoptionDetails(adoption: ResolutionResult["adoptions"][number]): string {
  const effectiveDate = adoption.dates.effective_date;
  const parts = [formatCodeFamily(adoption.codeFamily)];
  if (effectiveDate) {
    parts.push(`effective ${effectiveDate}`);
  }
  return parts.join(" · ");
}

function formatRuleKind(kind: ResolutionResult["applicableRules"][number]["kind"]): string {
  return kind.charAt(0).toUpperCase() + kind.slice(1);
}

function ResultSection({ title, children }: { title: string; children: ReactNode }): JSX.Element {
  return (
    <section className="public-result__section">
      <h3>{title}</h3>
      <div>{children}</div>
    </section>
  );
}
