# Runtime and API Reference

## Repository Scripts

Run from the repository root:

- `pnpm dev`: start the Vite dev server with a loopback-only `/api` proxy to the backend.
- `pnpm build`: produce a production bundle.
- `pnpm check`: run the TypeScript compiler without emitting files.
- `pnpm decision:recover`: run the repo's Deciduous session-start recovery summary.
- `pnpm decision:audit`: run the repo's Deciduous graph-hygiene audit.
- `pnpm decision:audit:strict`: run the Deciduous audit in failing mode for hard gaps.
- `pnpm lint`: run ESLint on `src/`.
- `pnpm lint:fix`: apply safe ESLint fixes.
- `pnpm backend:lint`: run `go vet` against the backend module.
- `pnpm backend:format`: run `go fmt` across the backend module.
- `pnpm backend:test`: run the backend test suite.
- `pnpm format`: rewrite checked-in files with Prettier.
- `pnpm format:check`: verify Prettier formatting.
- `pnpm test`: run the Vitest suite once.
- `pnpm test:watch`: keep Vitest running in watch mode.

## Direct Backend Commands

Run from `backend/`:

- `go run ./cmd/server --addr 127.0.0.1:8000`: start the backend service.
- `go run ./cmd/geocoder-build --help`: inspect the local geocoder snapshot builder.
- `go test ./...`: run the backend test suite.
- `go vet ./...`: run the backend static checks.

## Environment

- `VITE_API_BASE_URL`: optional frontend API base URL override. Default: `/api`, which uses the Vite dev proxy in local development.
- `BACKEND_CORS_ALLOWED_ORIGINS`: comma-separated list of allowed browser origins for the backend CORS middleware. Default: `http://localhost:5173,http://127.0.0.1:5173,http://[::1]:5173`.
- `TIGERWEB_CACHE_PATH`: optional boundary snapshot path inside the repo checkout.
- `TIGERWEB_HYDRATED_CACHE_PATH`: optional hydrated boundary-snapshot override inside the repo checkout.
- `GEOCODER_DATA_PATH`: optional local geocoder SQLite path inside the backend checkout. It overrides `--geocoder-data`.
- `DUCKDB_EXE`: optional path to the DuckDB CLI binary used only for legacy DuckDB snapshot loading. `DUCKDB_CLI_PATH` is accepted as a legacy fallback.

The frontend requests the API at `/api/*` by default. In local development, Vite rewrites that path to the loopback backend.

## API Endpoints

- `GET /health`: liveness check.
- `GET /ready`: readiness check for the cached boundary snapshot.
- `GET /layers`: list the supported layer families.
- `GET /boundaries`: list the cached boundary feature records with geometry.
- `GET /features/{layer_family}/{feature_id}`: fetch a full feature record.
- `POST /geocode`: resolve a civic address using the optional local SQLite geocoder.
- `POST /lookup`: geocode a civic address and pass the selected point into regulatory resolution.
- `POST /resolve`: resolve an existing longitude/latitude point or normalized geographic context.
- `GET /refresh/status`: read refresh status metadata.
- `POST /refresh/trigger`: report that live refresh is disabled for the cached snapshot.

## `POST /geocode`

Request:

```json
{
  "address": "1600 N Broadway, Denver, CO 80202"
}
```

A successful result contains the normalized query, selected candidate, candidate list, and warnings. Each candidate contains:

- `matched_address`
- `longitude`
- `latitude`
- `precision`: `address_point` or `interpolated`
- `confidence`
- `source`
- `source_record_id`
- `source_vintage`

Status behavior:

- `200`: one candidate was selected.
- `400`: the input is not a supported civic address or the JSON contract is invalid.
- `409`: multiple candidates remain materially tied; no candidate is selected.
- `422`: the address is valid but absent from the local snapshot.
- `503`: no usable geocoder snapshot is loaded.

## `POST /lookup`

Request:

```json
{
  "address": "1600 N Broadway, Denver, CO 80202",
  "code_family": "building",
  "applicability_date": "2026-08-01"
}
```

A successful response contains two top-level fields:

```json
{
  "geocode": {},
  "resolution": {}
}
```

The geocoding result remains visible so callers can inspect whether the point came from an address point or interpolation and which source supplied it. Ambiguous and unmatched geocoding results do not proceed to regulatory resolution.

## `POST /resolve`

The existing point-based contract remains independent from geocoding:

```json
{
  "point": {
    "longitude": -104.9903,
    "latitude": 39.7392
  },
  "code_family": "building",
  "applicability_date": "2026-08-01"
}
```

Coordinate resolution remains available when the geocoder is absent or invalid.

## Snapshot behavior

The backend boundary snapshot and geocoder snapshot are separate artifacts.

- The boundary snapshot supplies state, county, municipality, special-area, tribal-area, and NERIS geometries.
- The geocoder snapshot supplies address points and optional street ranges.
- Regulatory profiles supply authority, adoption, rule, and evidence records.

This separation prevents address matching from becoming regulatory policy and allows each source family to carry its own provenance, refresh cadence, and redistribution review.

The geocoder prefers address points. It uses parity-aware street-range interpolation only when no qualifying address point exists. Interpolation is labeled and returned with a warning.

## Layer Family Keys

- `states`
- `counties`
- `municipalities`
- `special_areas`
- `tribal_areas`
- `neris_jurisdictions`

## Shared Boundary Record Fields

- `layerFamily`
- `featureId`
- `title`
- `subtitle`
- `sourceId`
- `geometryLabel`
- `geometrySource`
- `lastSyncedAt`
- `attributes`

## Refresh Status Fields

- `status`
- `latestSuccessfulRefresh`
- `latestAttempt`
- `nextScheduledRefresh`
- `message`
