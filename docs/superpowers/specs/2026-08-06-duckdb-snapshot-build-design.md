# DuckDB snapshot build lane design

## Decision

Use DuckDB as a deterministic build and analytical QA workspace. Continue to use SQLite as the runtime snapshot format and Go as the semantic authority.

```text
source files
    -> DuckDB normalization and audit
    -> canonical ordered interchange rows
    -> existing Go SQLite finalizers
    -> snapshot manifest and activation lifecycle
    -> Building Code Authority Engine
```

## Authority boundaries

DuckDB owns bulk source scanning, deterministic ordering, duplicate classification, rejection classification, quarantine classification, and optional spatial validity checks.

Go owns input contracts, tool identity verification, final SQLite schemas, runtime validation, snapshot manifests, activation, geocoder ranking, interpolation, ambiguity, and regulatory meaning.

The build lane must not add a DuckDB runtime dependency, CGO requirement, network acquisition behavior, alternative ranking policy, or shadow snapshot schema.

## Tool and extension policy

The caller supplies an absolute DuckDB executable path, expected version, and expected SHA-256. Boundary builds also supply an absolute spatial extension path and expected SHA-256.

SQL contracts reject installation and remote-source tokens. The process receives a bounded environment and runs inside disposable workspace state. Tool, extension, SQL-contract, source, and output identities are recorded.

## Determinism

The caller supplies `built_at` and source retrieval timestamps. SQL exports rows in canonical order. Go creates SQLite databases from clean temporary files in canonical insertion order and validates them before replacement.

The build receipt excludes machine-local input paths. Component manifests use the existing PR #45 schema and bind the detailed receipt through the builder revision identity rather than introducing a second manifest authority.

## Initial source contracts

Geocoder inputs are canonical address-point and street-range CSV families already accepted by the Go builder.

Boundary inputs are layer-family, boundary-feature, and refresh-status CSV families matching the runtime snapshot contract. The spatial extension validates GeoJSON geometry before accepted rows reach the Go SQLite finalizer.

This is a bounded build lane, not a general ETL framework.
