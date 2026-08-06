# Building Code Map

Building Code Map is a local-first project for answering a difficult regulatory question deterministically: **given a location and an as-of date, which authorities may govern it, which adopted-code instruments are supported by evidence, and what still requires confirmation?**

The product combines three deliberately separate systems: address-to-point geocoding, point-to-geographic-boundary observation, and source-backed regulatory policy. Keeping those boundaries separate preserves provenance and prevents an approximate address match or containing polygon from being mistaken for a regulatory conclusion.

> **Project status:** pre-1.0 research and engineering. Results are not a substitute for confirmation by the controlling authority having jurisdiction, a licensed professional, or current official publications.

## What Works Today

- public address-or-coordinate lookup plus a technical GIS explorer
- deterministic local SQLite geocoder engine and snapshot builder
- authoritative address-point preference with visibly labeled street-range interpolation
- React, TypeScript, Vite, and MapLibre GL JS frontend
- Go HTTP API
- validated, read-only local SQLite boundary snapshots
- explicitly configured legacy DuckDB migration compatibility
- state, county, municipality, special land-use, American Indian area, and NERIS jurisdiction observations
- source-backed jurisdiction and adopted-code policy for six executable pilot states: Colorado, Florida, New Jersey, Virginia, Oregon, and North Carolina
- point-only public regulatory resolution with deterministic boundary-ambiguity responses
- capability-specific readiness for boundary, geocoder, regulatory, and composed lookup paths
- LORE-compatible semantic documentation trust root
- read-only GitHub Actions verification for frontend, backend, regulatory data, and repository documentation

The repository intentionally excludes hydrated boundary and geocoder databases. Runtime snapshots can be large, stale, or redistribution-sensitive, so they must be generated or supplied locally with provenance intact.

## Five-Minute Quick Start

### Prerequisites

- Node.js 22
- pnpm 10.12.1-compatible tooling
- the Go version declared in [`backend/go.mod`](backend/go.mod)
- a supported boundary snapshot at `backend/data/tigerweb.sqlite`, unless a repository-contained override is configured explicitly
- optionally, a locally built `backend/data/geocoder.sqlite` for address lookup

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

`/health` proves only that the process responds. `/ready` reports `ready`, `degraded`, or `not_ready` plus capability records. A missing geocoder snapshot disables address geocoding without disabling boundary inspection. Missing regulatory profiles leave geographic capabilities available while code conclusions remain unavailable.

On Windows, `./tools/start.cmd` starts both services and waits for backend readiness. `./tools/health.cmd` probes the running services, and `./tools/stop.cmd` stops the processes launched by the repository scripts.

## Delivered Architecture

```text
Address input
    |
    v
Local SQLite geocoder -----> provenance-bearing point
                                  |
Coordinate input ----------------+
                                  |
                                  v
                    Geographic boundary observations
                                  |
                                  v
              Source-backed policy plus as-of date
                                  |
                                  v
                         Regulatory resolver
                                  |
                                  v
          Candidate authorities, adoptions, evidence,
                 uncertainty, and local confirmation
```

OpenStreetMap is a browser basemap. TIGERweb, NERIS, address sources, and official regulatory sources feed offline ingestion or research workflows; they are not lookup-time browser dependencies.

The backend exposes:

- `GET /health`
- `GET /ready`
- `GET /layers`
- `GET /boundaries`
- `GET /features/{layer_family}/{feature_id}`
- `POST /geocode`
- `POST /lookup`
- `POST /resolve`
- `POST /v1/geocode`
- `POST /v1/lookup`
- `POST /v1/resolve`
- `GET /v1/readiness`
- `GET /v1/bundle`
- `GET /refresh/status`
- `POST /refresh/trigger`

`POST /resolve` accepts a point, never caller-authored jurisdiction context. The backend derives boundary observations from its validated snapshot before invoking the internal regulatory resolver. Multiple containing states, counties, or municipalities return `409 boundary_ambiguous` with every tied observation instead of selecting the first polygon. `POST /lookup` composes local geocoding with the same point path and returns both results.

When `applicability_date` is omitted, the API uses the current UTC date for compatibility and emits an explicit warning. That default is not evidence of project applicability.

## Build a Local Geocoder Snapshot

The repository includes a deterministic CSV-to-SQLite builder:

```bash
cd backend
go run ./cmd/geocoder-build \
  --output data/geocoder.sqlite \
  --address-points ./path/to/address-points.csv \
  --street-ranges ./path/to/street-ranges.csv \
  --source-name "Reviewed local address data" \
  --source-vintage "2026-08-01"
```

Address sources require independent licensing and redistribution review. See [Build a local geocoder snapshot](docs/how-to/build-local-geocoder-snapshot.md) and [Data Sources and Attribution](DATA_SOURCES.md).

## Repository Layout

- `src/`: public lookup, explorer, map interaction, and frontend data contracts
- `backend/engine/`: deterministic authority resolution, geography orchestration, provenance, and typed contracts
- `backend/geocoder/`: address normalization, SQLite matching, interpolation, provenance, and snapshot construction
- `backend/transport/httpapi/`: thin versioned and legacy HTTP adapters
- `backend/regulatory/`: source-backed policy catalog and generic resolver
- `backend/snapshot/`: boundary snapshot loading and semantic validation
- `backend/bundle/`: content-addressed aggregate bundle verification
- `backend/cmd/bcm/`: machine-safe CLI and shared serving entrypoint
- `backend/cmd/geocoder-build/`: reproducible geocoder snapshot builder
- `reports/`: state building-code research artifacts and transformation tooling
- `docs/`: tutorials, how-to guides, references, explanations, reviews, and generated LORE projections
- `.lore/records/`: accepted semantic repository knowledge
- `skills/maintain-repository-documentation/`: neutral LORE maintenance protocol
- `tools/`: Windows launch, health, stop, reporting, and Deciduous helpers
- `.github/workflows/ci.yml`: read-only repository verification

## Roadmap

### Current

- local address or coordinate entry
- provenance-bearing address-point and interpolated geocoding
- validated cached boundary exploration and geographic resolution
- source-backed regulatory resolution for six executable pilot states
- explicit uncertainty, ambiguity, and required-local-record outputs

### Next

- acquire and review production address sources for selected jurisdictions
- add complete source manifests, checksums, build identities, activation receipts, and rollback identities to runtime snapshots
- deepen code-family, authority, amendment, and fixture coverage within the executable pilots
- broaden municipal adopted-code coverage
- add reproducible boundary refresh and atomic activation workflows
- improve geocoder ranking and interpolation provenance using deterministic source-specific rules
- validate frontend API payloads at runtime instead of relying on TypeScript assertions
- design municipality-conditioned policy, geographic exceptions, provision-level status, and multi-agency routing before promoting Massachusetts, New York, or California

### Later

- PostGIS-backed storage and vector-tile serving where scale requires them
- historical versioning for addresses, boundaries, authorities, adoptions, and amendments
- unit, parcel, landmark, intersection, and reverse geocoding where source quality supports them
- broader nationwide coverage with source-level verification

Deferred systems are architectural direction, not current delivered functionality.

## Verification

Run frontend checks from the repository root:

```bash
pnpm check
pnpm lint
pnpm format:check
pnpm test
pnpm build
```

Run backend verification from `backend/`:

```bash
go test ./...
go vet ./...
```

GitHub Actions also validates regulatory compilation and the repository-local LORE trust root against a pinned public LORE revision.

## Documentation

- [Getting started tutorial](docs/tutorials/getting-started.md)
- [Build a local geocoder snapshot](docs/how-to/build-local-geocoder-snapshot.md)
- [Runtime and API reference](docs/reference/runtime-and-api.md)
- [Configuration reference](docs/reference/configuration.md)
- [Repository governance reference](docs/reference/repository-governance.md)
- [Add a boundary layer](docs/how-to/add-boundary-layer.md)
- [Populate a state report](docs/how-to/populate-state-report-template.md)
- [Local SQLite geocoder design](docs/superpowers/specs/2026-08-01-local-sqlite-geocoder-design.md)
- [Comprehensive repository review, 2026-08-03](docs/reviews/2026-08-03-comprehensive-repository-review.md)
- [Data sources and redistribution boundaries](DATA_SOURCES.md)
- [Security policy](SECURITY.md)
- [Contribution guide](CONTRIBUTING.md)

`docs/generated/` contains non-authoritative LORE projections. Accepted semantic records, schemas, repository facts, and transaction history remain the source of truth.

## Data and Publication Boundaries

- OpenStreetMap attribution remains visible in the map.
- Boundary and address sources retain independent provenance and redistribution reviews.
- Census and NERIS geometry remains geographic evidence, not self-proving regulatory authority.
- NERIS operational data is not assumed to be redistributable; individual fire departments retain ownership of their data.
- Model-code and standards text is not included merely because a jurisdiction adopts it.
- Environment files, logs, local caches, decision databases, prompt-bearing graph exports, and database binaries are ignored.

See [`DATA_SOURCES.md`](DATA_SOURCES.md) for the maintained attribution and redistribution boundary.

## Contributing

Use the structured issue forms for software bugs, jurisdiction or code-data corrections, and feature requests. Pull requests must preserve provenance, publication boundaries, and the exact-head verification policy described in [`CONTRIBUTING.md`](CONTRIBUTING.md).

Security vulnerabilities must be reported privately according to [`SECURITY.md`](SECURITY.md), not through a public issue.

## License

Project-authored software and documentation are licensed under the [Apache License 2.0](LICENSE). Third-party components, data, and services retain their own terms; see [`NOTICE`](NOTICE) and [`DATA_SOURCES.md`](DATA_SOURCES.md).
