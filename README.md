# Building Code Map

A local-first boundary and authority explorer for answering the eventual question: **given a location, which authorities and adopted building codes apply?**

The current application focuses on the geographic foundation of that workflow:

- React, TypeScript, and Vite frontend
- MapLibre GL JS with an OpenStreetMap basemap
- Go HTTP backend
- SQLite runtime snapshots cached from TIGERweb and related sources
- legacy DuckDB migration fallback
- boundary families for states, counties, municipalities, special land-use areas, American Indian areas, and NERIS jurisdictions
- source-backed state building-code research artifacts

## Project Status

This is a pre-1.0 research and engineering project. It is not a substitute for confirmation by the controlling authority having jurisdiction, a licensed professional, or current official publications.

The repository intentionally does **not** include hydrated SQLite or DuckDB runtime databases. Those files can contain large, stale, or redistribution-sensitive data and must be generated or supplied locally with provenance intact.

## What Is Implemented

### Frontend

- application header and refresh-status banner
- boundary-family toggles
- live MapLibre map with OpenStreetMap attribution
- cached boundary selection and click-to-select overlays
- full-record inspector backed by API detail lookup
- loading, empty, and error states

### Backend

- `GET /health`
- `GET /ready`
- `GET /layers`
- `GET /boundaries`
- `GET /features/{layer_family}/{feature_id}`
- `GET /refresh/status`
- `POST /refresh/trigger`

The backend prefers a local SQLite snapshot and retains DuckDB support only for migration compatibility.

## Repository Layout

- `src/`: React frontend
- `backend/`: Go API and snapshot readers
- `reports/`: state building-code research artifacts and report tooling
- `docs/`: Diátaxis documentation
- `tools/`: Windows launch, health, stop, reporting, and Deciduous helpers
- `.github/workflows/ci.yml`: read-only frontend and backend verification

## Documentation

The documentation is organized by intent under [`docs/`](docs/README.md):

- Tutorial: [`docs/tutorials/getting-started.md`](docs/tutorials/getting-started.md)
- How-to: [`docs/how-to/add-boundary-layer.md`](docs/how-to/add-boundary-layer.md)
- How-to: [`docs/how-to/populate-state-report-template.md`](docs/how-to/populate-state-report-template.md)
- How-to: [`docs/how-to/use-decision-graph-for-repo-work.md`](docs/how-to/use-decision-graph-for-repo-work.md)
- Reference: [`docs/reference/runtime-and-api.md`](docs/reference/runtime-and-api.md)
- Reference: [`docs/reference/configuration.md`](docs/reference/configuration.md)
- Reference: [`docs/reference/fema_bcat.md`](docs/reference/fema_bcat.md)
- Explanation: [`docs/explanation/tigerweb-visualization-design.md`](docs/explanation/tigerweb-visualization-design.md)
- Data and attribution: [`DATA_SOURCES.md`](DATA_SOURCES.md)
- Security reporting: [`SECURITY.md`](SECURITY.md)

## Prerequisites

- Node.js 22
- pnpm 10
- Go version declared by [`backend/go.mod`](backend/go.mod)
- a locally generated or supplied TIGERweb snapshot, normally `backend/data/tigerweb.sqlite`

## Frontend

Install and run from the repository root:

```bash
pnpm install --frozen-lockfile
pnpm dev
```

Useful checks:

```bash
pnpm check
pnpm lint
pnpm format:check
pnpm test
pnpm build
```

`VITE_API_BASE_URL` points the frontend at the API. It defaults to `/api`, which the Vite development server proxies to `127.0.0.1:8000`.

## Backend

Run from `backend/`:

```bash
cd backend
go run ./cmd/server --addr 127.0.0.1:8000
```

Verify the backend:

```bash
go test ./...
go vet ./...
```

Configuration:

- `TIGERWEB_CACHE_PATH`: optional snapshot path inside the repository checkout
- `TIGERWEB_HYDRATED_CACHE_PATH`: compatibility override inside the checkout
- `BACKEND_CORS_ALLOWED_ORIGINS`: comma-separated allowed browser origins; defaults to loopback-only origins
- `DUCKDB_EXE` or `DUCKDB_CLI_PATH`: optional legacy DuckDB CLI path

## Windows Development Scripts

- `.\tools\start.cmd`: start frontend and backend
- `.\tools\start.cmd -FrontendOnly`: start only the frontend
- `.\tools\start.cmd -BackendOnly`: start only the backend
- `.\tools\health.cmd`: probe both services
- `.\tools\stop.cmd`: stop the launched services
- `.\tools\deciduous-work.cmd`: start a repo-visible decision initiative
- `.\tools\deciduous-recover.cmd`: summarize decision state
- `.\tools\deciduous-audit.cmd`: inspect graph hygiene

The launcher resolves `go` from `PATH`, starts the API with `go run ./cmd/server`, and waits on `GET /ready` before reporting success.

## Data and Publication Boundaries

- OpenStreetMap attribution remains visible in the map.
- Census and other official-source provenance must remain attached to transformed records.
- NERIS operational data is not assumed to be redistributable; individual fire departments retain ownership of their data.
- Model-code and standards text is not included merely because a jurisdiction adopts it.
- Environment files, logs, local caches, decision databases, generated prompt-bearing graph exports, and database binaries are ignored.

See [`DATA_SOURCES.md`](DATA_SOURCES.md) for the maintained attribution and redistribution boundary.

## Next Implementation Steps

- replace demo or cached GeoJSON overlays with Martin vector tiles backed by the PostGIS mirror
- make snapshot ingestion reproducible and refreshable from official live sources
- connect location lookup to the normalized authority graph and state-report evidence
- add historical versioning for boundary, authority, adoption, and amendment changes
- expand jurisdiction coverage while preserving source-level verification and unresolved-state labels

## License

Project-authored software and documentation are licensed under the Apache License 2.0. Third-party data and services retain their own terms; see [`DATA_SOURCES.md`](DATA_SOURCES.md).