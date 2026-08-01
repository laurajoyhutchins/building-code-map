# Local SQLite Geocoder Design

**Date:** 2026-08-01

## Goal

Accept a United States site address, resolve it locally to a provenance-bearing point, and pass that point into the existing deterministic jurisdiction resolver without making regulatory policy depend on address parsing or a network geocoding service.

## Product boundary

The geocoder answers **where is this address?** The regulatory resolver continues to answer **which geographic boundaries, authorities, and supported adoption records apply at this point?**

The existing point-based `POST /resolve` contract remains authoritative and unchanged. New address-oriented endpoints compose with it instead of replacing it.

## First vertical slice

The first slice supports a reproducible SQLite snapshot containing authoritative address points and optional TIGER-style street ranges.

- Exact or strongly constrained address-point matches are preferred.
- Street-range interpolation is a lower-confidence fallback and is labeled `interpolated`.
- Near-equal candidates produce an `ambiguous` result rather than an arbitrary selection.
- PO boxes and inputs without a usable civic address fail closed.
- No browser or backend request depends on an external geocoding API.
- The repository does not commit hydrated nationwide databases.

## Architecture

### `backend/internal/geocoder`

This package owns normalization, candidate ranking, SQLite access, provenance, and ambiguity policy.

Its public boundary is intentionally small:

```go
type Service interface {
    Geocode(ctx context.Context, query Query) (Result, error)
}
```

`Result` contains the normalized query, status, selected candidate when one exists, candidate list, and warnings. Every candidate includes longitude, latitude, precision, confidence, source, source record identifier, and source vintage.

### SQLite snapshot

The geocoder snapshot is separate from the boundary snapshot. This keeps refresh cadence, redistribution review, and provenance independent.

Required tables:

- `geocoder_metadata(key, value)`
- `address_points(...)`
- `street_ranges(...)`

The schema stores pre-normalized comparison fields. Request-time code performs bounded indexed queries rather than scanning or calling a fuzzy-search service.

### HTTP composition

- `POST /geocode` returns the geocoding result only.
- `POST /lookup` accepts an address plus the existing regulatory options, geocodes it, and resolves the selected point.
- `POST /resolve` remains point or normalized-context based.

`/lookup` returns both the geocoding result and the regulatory result so callers can inspect how the point was obtained.

### Public frontend

The location field accepts either coordinates or a civic address.

- Coordinate input continues to call `/resolve` directly.
- Address input calls `/lookup`.
- When an address was geocoded, the result shows the matched address, precision, source, and source vintage.
- Ambiguous or unmatched results are shown as errors and do not silently call the regulatory resolver.

## Normalization and matching

Normalization is deterministic and deliberately conservative:

1. Trim and collapse whitespace.
2. Parse house number, street, locality, state, and optional ZIP from a comma-delimited United States civic address.
3. Uppercase comparison fields.
4. Remove punctuation used only for presentation.
5. Normalize a bounded list of street suffixes and cardinal directions.
6. Query address points using house number, normalized street, and state; use locality and ZIP as ranking constraints.
7. Accept one top candidate only when it clears the minimum score and the score gap from the runner-up.
8. Otherwise return `ambiguous` or `not_found`.
9. If no address point matches, query compatible street ranges and interpolate only when parity and range constraints permit it.

The selected candidate is never inferred from regulatory boundaries.

## Data ingestion

A small Go command imports normalized CSV files into a new SQLite snapshot transactionally.

```text
go run ./cmd/geocoder-build --output data/geocoder.sqlite --address-points input.csv --street-ranges ranges.csv
```

The builder:

- creates a temporary database;
- validates coordinates and required provenance fields;
- rejects duplicate source identities;
- creates indexes after loading;
- writes source metadata;
- atomically replaces the requested output only after validation succeeds.

Hydrated SQLite files remain ignored. Small CSV fixtures used by tests may be committed.

## Error handling

- Missing geocoder snapshot: `/geocode` and address-based `/lookup` return `503`; coordinate `/resolve` remains available.
- Invalid civic address: `400`.
- Valid but unmatched address: `422` with `status: not_found`.
- Ambiguous address: `409` with candidate summaries.
- Geocoded point outside supported boundaries: retain the existing resolver `422` behavior.
- Corrupt or incompatible snapshot: server logs the load failure and leaves geocoding unavailable rather than partially loading it.

## Security and privacy

- Request bodies remain limited to 1 MiB and reject unknown fields and trailing JSON.
- SQLite paths use the existing repository-contained configuration pattern.
- Queries are parameterized.
- Logs must not emit full user-entered addresses at normal log level.
- No address history or analytics storage is introduced.

## Verification

Backend tests cover normalization, exact matching, ambiguity, interpolation, provenance, snapshot validation, HTTP status behavior, and end-to-end address-to-resolution composition.

Frontend tests cover coordinate/address classification, lookup response decoding, and provenance presentation helpers.

Repository verification remains:

```text
pnpm check
pnpm lint
pnpm format:check
pnpm test
pnpm build
cd backend && go test ./... && go vet ./...
```

## LORE compatibility

LORE is initialized as a repository-local trust root without making generated prose authoritative.

- `lore.yaml` declares paths, supported extractors, projections under `docs/generated/`, and the maintenance skill.
- LORE schemas and the exact neutral maintenance skill are copied from the pinned LORE revision used by CI.
- Initial bootstrap records describe the repository, the geocoder component, the point-based separation decision, the local-only constraint, and the snapshot-building procedure.
- Each record cites an exact full Git commit and repository path.
- Accepted records are append-only. Future semantic changes must be proposed with `lore-proposal/v1`; generated documents and extracted facts are never edited directly.
- CI checks LORE validation against a pinned public LORE commit.

The hand-authored design and implementation plan are evidence. They are not substitutes for accepted semantic records.

## Deferred work

- Nationwide data acquisition and redistribution decisions.
- Unit, building, parcel, landmark, and intersection geocoding.
- Probabilistic or machine-learned ranking.
- Network-provider fallback.
- Reverse geocoding.
- Historical address snapshots.
- Automated source refresh scheduling.

## Acceptance criteria

- A supported civic address can resolve locally to a provenance-bearing point and then to the existing regulatory result.
- Coordinates still work when geocoding is unavailable.
- Interpolated, ambiguous, unmatched, and exact address-point outcomes remain distinct.
- No external network request is required at lookup time.
- Hydrated databases are not committed.
- LORE validates the manifest, schemas, skill, and bootstrap records against exact Git evidence.
