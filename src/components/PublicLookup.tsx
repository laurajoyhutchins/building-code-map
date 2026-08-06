import { useEffect, useRef, useState, type FormEvent, type ReactNode } from "react";
import { fetchLookup, fetchReadiness, fetchResolution } from "../lib/api";
import {
  classifyLocationQuery,
  formatCodeFamily,
  formatGeocodeSummary,
  getResolutionNotice,
  getResolutionPlace,
} from "../lib/publicLookup";
import type { GeocodeResult, ReadinessResult, ResolutionResult } from "../types";
import { ApiEvidenceNotice, ReadinessNotice } from "./ApiEvidenceNotice";

const codeFamilies = [
  ["building", "Building"],
  ["residential", "Residential"],
  ["electrical", "Electrical"],
  ["energy", "Energy"],
  ["fire_operational", "Operational fire"],
] as const;

const today = new Date().toISOString().slice(0, 10);

export function PublicLookup(): JSX.Element {
  const [location, setLocation] = useState("");
  const [codeFamily, setCodeFamily] = useState("building");
  const [applicabilityDate, setApplicabilityDate] = useState(today);
  const [result, setResult] = useState<ResolutionResult | null>(null);
  const [geocode, setGeocode] = useState<GeocodeResult | null>(null);
  const [error, setError] = useState<unknown>(null);
  const [readiness, setReadiness] = useState<ReadinessResult | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const requestController = useRef<AbortController | null>(null);

  const explorerHref = new URL("explorer", window.location.href).pathname;

  useEffect(() => {
    const controller = new AbortController();
    void fetchReadiness()
      .then(setReadiness)
      .catch((readinessError: unknown) => {
        if (!controller.signal.aborted) setError(readinessError);
      });
    return () => controller.abort();
  }, []);

  useEffect(() => () => requestController.current?.abort(), []);

  async function handleSubmit(event: FormEvent<HTMLFormElement>) {
    event.preventDefault();

    let query;
    try {
      query = classifyLocationQuery(location);
    } catch (locationError) {
      setError(locationError instanceof Error ? locationError : new Error("Enter a valid location."));
      return;
    }

    requestController.current?.abort();
    const controller = new AbortController();
    requestController.current = controller;
    setIsLoading(true);
    setError(null);
    try {
      if (query.kind === "coordinates") {
        const nextResult = await fetchResolution({
          longitude: query.point.longitude,
          latitude: query.point.latitude,
          codeFamily,
          applicabilityDate,
          signal: controller.signal,
        });
        setGeocode(null);
        setResult(nextResult);
      } else {
        const lookup = await fetchLookup({
          address: query.address,
          codeFamily,
          applicabilityDate,
          signal: controller.signal,
        });
        setGeocode(lookup.geocode);
        setResult(lookup.resolution);
      }
    } catch (requestError) {
      if (controller.signal.aborted) return;
      setResult(null);
      setGeocode(null);
      setError(requestError instanceof Error ? requestError : new Error("The lookup failed."));
    } finally {
      if (requestController.current === controller) {
        requestController.current = null;
        setIsLoading(false);
      }
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

        <ReadinessNotice readiness={readiness} />

        <form className="public-search" onSubmit={handleSubmit}>
          <label className="public-search__location">
            <span>Address or coordinates</span>
            <input
              autoComplete="street-address"
              placeholder="1600 N Broadway, Denver, CO 80202"
              value={location}
              onChange={(event) => setLocation(event.target.value)}
              required
            />
            <small>
              Enter a United States civic address or latitude and longitude. Address lookup uses a
              local snapshot when available.
            </small>
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

            <button type="submit" disabled={isLoading || readiness?.readiness === "not_ready"}>
              {isLoading ? "Searching…" : "Search"}
            </button>
          </div>
        </form>

        <ApiEvidenceNotice error={error} />

        {result ? <PublicResult result={result} geocode={geocode} /> : null}
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

function PublicResult({
  result,
  geocode,
}: {
  result: ResolutionResult;
  geocode: GeocodeResult | null;
}): JSX.Element {
  const notice = getResolutionNotice(result.status);
  const authorities = result.authorityCandidates;
  const adoptions = result.adoptions;

  return (
    <article className="public-result" aria-live="polite">
      <header className="public-result__header">
        <p>{formatCodeFamily(result.codeFamily)}</p>
        <h2>{getResolutionPlace(result)}</h2>
        {geocode?.selected ? (
          <div className="public-location-match">
            <strong>{geocode.selected.matchedAddress}</strong>
            <span>{formatGeocodeSummary(geocode.selected)}</span>
            {geocode.warnings.map((warning) => (
              <span key={warning}>{warning}</span>
            ))}
          </div>
        ) : null}
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
                <span>Verification: {authority.verification.status.replace(/_/g, " ")}</span>
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
                <span>Verification: {adoption.verification.status.replace(/_/g, " ")}</span>
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
                <strong>
                  {relationship.summary ?? relationship.relationship.replace(/_/g, " ")}
                </strong>
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
                  {source.availability ? ` · ${source.availability}` : ""}
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
