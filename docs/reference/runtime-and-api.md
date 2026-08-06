# Runtime and API Reference

## Repository Scripts

Run from the repository root:

- `pnpm dev`: start the Vite dev server with a loopback-only `/api` proxy to the backend.
- `pnpm build`: produce a production bundle.
- `pnpm check`: run the TypeScript compiler without emitting files.
- `pnpm decision:recover`: run the repository's Deciduous session-start recovery summary.
- `pnpm decision:audit`: run the repository's Deciduous graph-hygiene audit.
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

- `VITE_API_BASE_URL`: optional frontend API base URL override. Default: `/api`, which uses the Vite development proxy.
- `BACKEND_CORS_ALLOWED_ORIGINS`: comma-separated browser origins accepted by the backend. Default: `http://localhost:5173,http://127.0.0.1:5173,http://[::1]:5173`.
- `TIGERWEB_CACHE_PATH`: optional repository-contained boundary snapshot path.
- `TIGERWEB_HYDRATED_CACHE_PATH`: legacy repository-contained alias for the boundary snapshot path.
- `GEOCODER_DATA_PATH`: optional local geocoder SQLite path inside the backend checkout. It overrides `--geocoder-data`.
- `DUCKDB_EXE`: optional path to the DuckDB CLI used only when an explicitly selected boundary snapshot has a `.duckdb` extension. `DUCKDB_CLI_PATH` is a legacy alias.

The frontend requests the API at `/api/*` by default. In local development, Vite rewrites that path to the loopback backend.

## API Endpoints

- `GET /health`: process liveness only.
- `GET /ready`: capability-specific readiness for boundary, coordinate, geocoder, regulatory, and composed lookup paths.
- `GET /layers`: list supported layer families.
- `GET /boundaries`: list cached boundary records with geometry. This remains a bulk endpoint and is tracked for list/detail redesign in issue #7.
- `GET /features/{layer_family}/{feature_id}`: fetch a full feature record.
- `POST /geocode`: resolve a civic address using the optional local SQLite geocoder.
- `POST /lookup`: geocode a civic address and pass the selected point into boundary and regulatory resolution.
- `POST /resolve`: resolve a longitude/latitude point. Caller-authored geographic context is not accepted.
- `POST /v1/geocode`, `POST /v1/lookup`, and `POST /v1/resolve`: versioned engine contracts with typed errors and exact provenance.
- `GET /v1/readiness` and `GET /v1/bundle`: versioned readiness and identity records.
- `GET /refresh/status`: read snapshot refresh metadata.
- `POST /refresh/trigger`: report that live refresh is disabled for the cached snapshot.

## Liveness and readiness

`GET /health` returns `200` with `{"status":"ok"}` when the process can answer HTTP. It does not validate datasets.

`GET /ready` returns:

- `200`, `readiness: "ready"` when boundary, geocoder, and regulatory capabilities are available;
- `200`, `readiness: "degraded"` when the required boundary snapshot is usable but one or more optional capabilities are unavailable;
- `503`, `readiness: "not_ready"` when boundary-backed operation is unavailable.

The response includes capability records for:

- `boundary_resolution`
- `coordinate_resolution`
- `address_geocoding`
- `regulatory_resolution`
- `address_lookup`

A capability record states whether it is available, whether it is required for process readiness, and why it is available or unavailable.

## `POST /geocode`

Request:

```json
{
  "address": "1600 N Broadway, Denver, CO 80202"
}
```

A successful result contains the original query, normalized matching form, selected candidate, candidate list, and warnings. Each candidate contains:

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

An interpolated result remains labeled and includes a warning. It is not an authoritative address point or parcel-level location.

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

The public contract is point-only:

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

The backend derives normalized geographic context from its boundary snapshot. A request containing `context` is rejected as an unknown field. Internal regulatory code may consume normalized context after it has been produced by boundary resolution, but public callers cannot assert jurisdiction, incorporation, or authority facts directly.

If a point matches more than one state, county, or municipality polygon, the endpoint returns `409`:

```json
{
  "error": "point matched multiple counties boundary observations; confirm the controlling boundary locally",
  "code": "boundary_ambiguous",
  "layer_family": "counties",
  "observations": []
}
```

The observations are sorted deterministically and retain layer, feature, name, and source identities. Special-area, tribal-area, and NERIS overlaps are preserved as non-exclusive geographic observations. Their containment does not establish building-code authority.

When `applicability_date` is omitted, the server uses the current UTC date and adds a warning identifying the assumed date. Callers should supply the relevant project date whenever applicability matters.

Coordinate resolution remains available when the geocoder is absent. Regulatory conclusions remain unavailable when no validated profile is loaded for the matched state.

## CORS

Configured browser origins receive `Access-Control-Allow-Origin` on normal responses. An allowed `OPTIONS` preflight receives `204` with:

- `Access-Control-Allow-Methods: GET, POST, OPTIONS`
- `Access-Control-Allow-Headers: Content-Type`
- `Access-Control-Max-Age: 600`

An unconfigured preflight origin receives `403`. CORS configuration is a browser boundary, not an authentication mechanism.

## Snapshot behavior

The boundary snapshot and geocoder snapshot are separate local artifacts.

- The boundary snapshot supplies geographic observations.
- The geocoder snapshot supplies address points and optional street ranges.
- Regulatory profiles supply authority, adoption, rule, claim, and evidence records.

The boundary loader accepts `.sqlite`, `.db`, and explicitly selected `.duckdb` files. Unknown extensions fail before DuckDB discovery. The default is only `backend/data/tigerweb.sqlite`; no machine-global cache path is searched implicitly.

SQLite boundary snapshots are opened read-only. Every supported boundary loader validates the same semantic contract before returning data, including:

- unique layer identities;
- unique `(layer_family, feature_id)` identities;
- registered layer references;
- required feature and source identities;
- Polygon or MultiPolygon geometry;
- finite, in-range, closed rings;
- object-shaped attributes;
- supported refresh status and coherent timestamps.

The current snapshot contract does not yet contain complete source manifests, checksums, build-tool identities, activation receipts, or rollback identities. Those remain follow-up work and must not be inferred from the refresh-status row.

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
