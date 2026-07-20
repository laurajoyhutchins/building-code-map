# Building Code Map

Building Code Map is an early-stage, local-first project for answering a difficult regulatory question deterministically: **given a location, which authorities and adopted building codes apply?**

The long-term product is a source-backed location-to-authority and code-evidence workflow. What exists today is the geographic and research foundation for that product: a boundary explorer, a Go API over local snapshots, and structured state building-code research artifacts.

> **Project status:** pre-1.0 research and engineering. Results are not a substitute for confirmation by the controlling authority having jurisdiction, a licensed professional, or current official publications.

## What Works Today

- React, TypeScript, and Vite frontend
- MapLibre GL JS with an attributed OpenStreetMap basemap
- Go HTTP API
- local SQLite runtime snapshots, with legacy DuckDB migration compatibility
- state, county, municipality, special land-use, American Indian area, and NERIS jurisdiction layers
- boundary selection and full-record inspection
- source-backed state building-code reports and normalization tooling
- read-only GitHub Actions verification for the frontend and backend

The repository intentionally excludes hydrated SQLite and DuckDB databases. Runtime snapshots can be large, stale, or redistribution-sensitive, so they must be generated or supplied locally with provenance intact.

## Five-Minute Quick Start

### Prerequisites

- Node.js 22
- pnpm 10.12.1-compatible tooling
- the Go version declared in [`backend/go.mod`](backend/go.mod)
- a supported local snapshot, normally `backend/data/tigerweb.sqlite`, for a ready data-backed runtime

Install frontend dependencies from the repository root:

```bash
pnpm install --frozen-lockfile
```

Start the backend in one terminal:

```bash
cd backend
go run ./cmd/server --addr 127.0.0.1:8000
```

Start the frontend in a second terminal from the repository root:

```bash
pnpm dev
```

Then check:

- backend liveness: `http://127.0.0.1:8000/health`
- backend readiness: `http://127.0.0.1:8000/ready`
- frontend: the local URL printed by Vite, normally `http://127.0.0.1:5173`

`/health` can succeed without a runtime snapshot. `/ready` reports whether the API has a usable supported snapshot. See the [configuration reference](docs/reference/configuration.md) for supported paths and environment variables.

On Windows, `./tools/start.cmd` starts both services and waits for backend readiness. `./tools/health.cmd` probes the running services, and `./tools/stop.cmd` stops the processes launched by the repository scripts.

## Delivered Architecture

```text
React + MapLibre frontend
        |
        | HTTP through the Vite development proxy
        v
Go API
        |
        v
Local SQLite snapshot
```

TIGERweb, NERIS, and official state sources feed ingestion and research workflows. They are not direct browser dependencies.

The backend currently exposes:

- `GET /health`
- `GET /ready`
- `GET /layers`
- `GET /boundaries`
- `GET /features/{layer_family}/{feature_id}`
- `GET /refresh/status`
- `POST /refresh/trigger`

The API prefers SQLite. DuckDB support remains only for migration compatibility.

## Repository Layout

- `src/`: React frontend and map interaction logic
- `backend/`: Go API, snapshot readers, and backend tests
- `reports/`: state building-code research artifacts and transformation tooling
- `docs/`: tutorials, how-to guides, references, and explanations
- `tools/`: Windows launch, health, stop, reporting, and Deciduous helpers
- `.github/workflows/ci.yml`: read-only frontend and backend verification

## Roadmap

### Current

- cached boundary exploration and attribute inspection
- Go API over local snapshots
- source-backed state research artifacts
- explicit provenance and redistribution boundaries

### Next

- a deterministic location-to-authority result
- normalized authority relationships with source evidence
- adopted-code output that preserves unresolved or conflicting states instead of guessing
- a reproducible snapshot ingestion path from official sources

### Later

- local geocoding and address normalization
- PostGIS-backed storage and Martin vector-tile serving where scale requires them
- historical versioning for boundaries, authorities, adoptions, and amendments
- broader jurisdiction coverage with source-level verification

Deferred systems are architectural direction, not current delivered functionality.

## Verification

Run the frontend checks from the repository root:

```bash
pnpm check
pnpm lint
pnpm format:check
pnpm test
pnpm build
```

Run backend verification either directly from `backend/`:

```bash
cd backend
go test ./...
go vet ./...
```

or through the root scripts:

```bash
pnpm backend:test
pnpm backend:lint
```

## Documentation

The documentation is organized by intent under [`docs/`](docs/README.md):

- [Getting started tutorial](docs/tutorials/getting-started.md)
- [Runtime and API reference](docs/reference/runtime-and-api.md)
- [Configuration reference](docs/reference/configuration.md)
- [Add a boundary layer](docs/how-to/add-boundary-layer.md)
- [Populate a state report](docs/how-to/populate-state-report-template.md)
- [TIGERweb visualization design](docs/explanation/tigerweb-visualization-design.md)
- [Data sources and redistribution boundaries](DATA_SOURCES.md)
- [Security policy](SECURITY.md)

## Data and Publication Boundaries

- OpenStreetMap attribution remains visible in the map.
- Census and other official-source provenance must remain attached to transformed records.
- NERIS operational data is not assumed to be redistributable; individual fire departments retain ownership of their data.
- Model-code and standards text is not included merely because a jurisdiction adopts it.
- Environment files, logs, local caches, decision databases, prompt-bearing graph exports, and database binaries are ignored.

See [`DATA_SOURCES.md`](DATA_SOURCES.md) for the maintained attribution and redistribution boundary.

## License

Project-authored software and documentation are licensed under the [Apache License 2.0](LICENSE). Third-party components, data, and services retain their own terms; see [`NOTICE`](NOTICE) and [`DATA_SOURCES.md`](DATA_SOURCES.md).
