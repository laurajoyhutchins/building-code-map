# Runtime and API Reference

## Repository Scripts

Run from the repository root:

- `npm run dev`: start the Vite dev server with a loopback-only `/api` proxy to the backend.
- `npm run build`: type-check and produce a production bundle.
- `npm run check`: run the TypeScript compiler without emitting files.
- `npm run decision:recover`: run the repo's Deciduous session-start recovery summary.
- `npm run decision:audit`: run the repo's Deciduous graph-hygiene audit.
- `npm run decision:audit:strict`: run the Deciduous audit in failing mode for hard gaps.
- `npm run lint`: run ESLint on `src/`.
- `npm run lint:fix`: apply safe ESLint fixes.
- `npm run backend:lint`: run `go vet` against the backend module.
- `npm run backend:format`: run `go fmt` across the backend module.
- `npm run backend:test`: run the backend test suite.
- `npm run format`: rewrite checked-in files with Prettier.
- `npm run format:check`: verify Prettier formatting.
- `npm test`: run the Vitest suite once.
- `npm run test:watch`: keep Vitest running in watch mode.

## Direct Backend Commands

Run from `backend/`:

- `go run ./cmd/server --addr 127.0.0.1:8000`: start the backend service.
- `go test ./...`: run the backend test suite.
- `go vet ./...`: run the backend static checks.

The root-level npm scripts above are the preferred entry points for linting, formatting, and testing the backend during normal repo work.

## Environment

- `VITE_API_BASE_URL`: optional frontend API base URL override. Default: `/api`, which uses the Vite dev proxy in local development.
- `BACKEND_CORS_ALLOWED_ORIGINS`: comma-separated list of allowed browser origins for the backend CORS middleware. Default: `http://localhost:5173,http://127.0.0.1:5173,http://[::1]:5173`.
- `TIGERWEB_CACHE_PATH`: optional snapshot path inside the repo checkout. If unset, the Go backend prefers `backend/data/tigerweb.sqlite` and falls back to `backend/data/tigerweb.duckdb` during the migration period.
- `TIGERWEB_HYDRATED_CACHE_PATH`: optional hydrated-snapshot override inside the repo checkout. Paths that resolve outside the checkout are ignored.
- `DUCKDB_EXE`: optional path to the DuckDB CLI binary used only for legacy DuckDB snapshot loading. `DUCKDB_CLI_PATH` is accepted as a legacy fallback.

The frontend requests the API at `/api/*` by default. In local development, Vite rewrites that path to the loopback backend. If you override `VITE_API_BASE_URL` with an absolute URL, the frontend will call that URL directly instead.

## API Endpoints

- `GET /health`: liveness check.
- `GET /ready`: readiness check for the cached backend snapshot.
- `GET /layers`: list the supported layer families.
- `GET /boundaries`: list the cached boundary feature records with geometry.
- `GET /features/{layer_family}/{feature_id}`: fetch a full feature record.
- `GET /refresh/status`: read refresh status metadata.
- `POST /refresh/trigger`: report that live refresh is disabled for the cached snapshot.

The runtime backend reads all boundary layers, feature records, and refresh metadata from `backend/data/tigerweb.sqlite` when present, with `backend/data/tigerweb.duckdb` retained as a legacy fallback during migration. `TIGERWEB_CACHE_PATH` and `TIGERWEB_HYDRATED_CACHE_PATH` are only honored when they resolve inside the repo checkout, so sandboxed sessions do not accidentally depend on a machine-global cache.

The snapshot also includes a `neris_jurisdictions` layer family populated from the public NERIS DepartmentJurisdiction polygons so the app can render real NERIS service areas alongside the mirrored TIGERweb layers. County polygons are used only as an explicit fallback when no NERIS jurisdiction polygon can be resolved.

The refresh endpoints are intentionally offline-only in the runtime API. `GET /refresh/status` reports the cached snapshot status, and `POST /refresh/trigger` returns a disabled response rather than attempting a live refresh.

The launcher probes `GET /ready` rather than `GET /health` so startup only succeeds after the cache-backed endpoints are actually usable.

## Layer Family Keys

- `states`
- `counties`
- `municipalities`
- `special_areas`
- `tribal_areas`
- `neris_jurisdictions`

## Shared Record Fields

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
