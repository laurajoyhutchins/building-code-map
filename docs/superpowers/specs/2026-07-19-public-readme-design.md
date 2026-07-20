# Public README Experience Design

## Goal

Make the public repository understandable within one screenful and runnable by a developer within five minutes, without duplicating the detailed documentation under `docs/`.

## Audience

The README serves two primary audiences:

1. A first-time visitor deciding what Building Code Map is and whether the project is relevant.
2. A developer trying to run the checked-in boundary explorer locally.

## Information architecture

The README uses progressive disclosure:

1. Product promise and explicit pre-1.0 status.
2. What works today.
3. Five-minute quick start.
4. Compact architecture and repository layout.
5. Current, next, and later roadmap.
6. Links to detailed documentation, data policy, security policy, and licensing.

## Product framing

The target product is a deterministic, source-backed workflow that maps a location to the applicable authority graph and adopted building-code evidence. The current release is the geographic and research foundation for that product, not a finished code-compliance service.

The README must not present deferred PostGIS, Martin vector tiles, geocoding, nationwide code coverage, or historical versioning as implemented.

## Quick-start contract

The quick start must:

- require Node.js 22, pnpm 10.12.1-compatible tooling, and the Go version declared in `backend/go.mod`;
- explain that runtime database snapshots are intentionally not committed;
- provide separate backend and frontend terminal commands;
- identify the expected local URLs;
- describe the expected behavior when no snapshot is available;
- link to detailed configuration and runtime documentation.

## Architecture summary

Use a text-only diagram that reflects the delivered runtime:

```text
React + MapLibre frontend
        |
        | HTTP / Vite proxy
        v
Go API
        |
        v
Local SQLite snapshot
```

TIGERweb, NERIS, and official state sources are inputs to ingestion and research workflows, not direct browser dependencies.

## Roadmap framing

- **Current:** boundary explorer, Go API, local snapshot readers, source-backed state research artifacts.
- **Next:** deterministic location-to-authority result with explicit source evidence and unresolved states.
- **Later:** reproducible ingestion, geocoding, PostGIS/Martin serving, historical versioning, and broader jurisdiction coverage.

## Validation

Before completion:

- cross-check commands against `package.json`, `backend/go.mod`, and `tools/`;
- run Prettier formatting checks;
- run frontend type-check, lint, tests, and build;
- run Go tests and `go vet ./...`;
- confirm the README contains no stale FastAPI/DuckDB-primary language or personal paths.
