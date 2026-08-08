# Build and audit snapshots with DuckDB

DuckDB is a build-time normalization and analytical QA tool. The Building Code Authority Engine continues to read canonical SQLite snapshots at runtime.

The build command never downloads DuckDB, extensions, or source data. Supply exact local files and their expected identities.

## Geocoder build

```bash
bcm snapshot build \
  --kind geocoder \
  --duckdb /opt/bcm-tools/duckdb \
  --duckdb-version 1.5.0 \
  --duckdb-sha256 <sha256> \
  --work-dir .bcm-build \
  --output backend/data/geocoder.sqlite \
  --address-points sources/address-points.csv \
  --street-ranges sources/street-ranges.csv \
  --publisher "Source publisher" \
  --product "Address source" \
  --locator "source://address-source/2026" \
  --vintage 2026 \
  --retrieved-at 2026-08-06T12:00:00Z \
  --license-review-status reviewed \
  --redistribution-status internal-build-only \
  --builder-version 0.1.0 \
  --builder-revision <git-commit> \
  --built-at 2026-08-06T12:00:00Z
```

DuckDB classifies, orders, and exports canonical CSV rows. The existing Go geocoder builder then writes the established SQLite schema and indexes. Ranking, ambiguity, and interpolation semantics remain in Go.

## Boundary build

Boundary builds additionally require an explicitly provisioned spatial extension and its digest:

```bash
bcm snapshot build \
  --kind boundary \
  --duckdb /opt/bcm-tools/duckdb \
  --duckdb-version 1.5.0 \
  --duckdb-sha256 <sha256> \
  --spatial-extension /opt/bcm-tools/spatial.duckdb_extension \
  --spatial-extension-sha256 <sha256> \
  --work-dir .bcm-build \
  --output backend/data/tigerweb.sqlite \
  --layer-families sources/layer-families.csv \
  --boundary-features sources/boundary-features.csv \
  --refresh-status sources/refresh-status.csv \
  --publisher "Source publisher" \
  --product "Boundary source" \
  --locator "source://boundary-source/2026" \
  --vintage 2026 \
  --retrieved-at 2026-08-06T12:00:00Z \
  --license-review-status reviewed \
  --redistribution-status internal-build-only \
  --builder-version 0.1.0 \
  --builder-revision <git-commit> \
  --built-at 2026-08-06T12:00:00Z
```

The SQL contract uses `LOAD` with that exact extension path. It never executes `INSTALL`.

## Audit without mutation

Replace `snapshot build` with `snapshot audit` and omit `--output`. Audit runs the same normalization and classification contract but does not finalize or activate a SQLite snapshot.

## Outputs

A successful build creates:

- the canonical SQLite snapshot;
- the PR #45 component manifest, by default `<snapshot>.manifest.json`;
- a deterministic build receipt, by default `<snapshot>.build.json`.

The component manifest binds the `bcm/duckdb-snapshot-build-1` pipeline through its builder identity. The build receipt carries the complete DuckDB executable, spatial extension, SQL-contract, source, count, and output digests.

Build does not activate the snapshot. Activation and rollback remain owned by the existing snapshot lifecycle.
