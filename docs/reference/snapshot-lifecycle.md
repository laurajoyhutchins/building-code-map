# Snapshot lifecycle

Building Code Map treats boundary and geocoder databases as immutable runtime capsules. A database file is not eligible for runtime use merely because SQLite can open it.

## Required sidecars

For a snapshot at `path/to/snapshot.sqlite`, the runtime requires:

- `path/to/snapshot.sqlite.manifest.json`
- `path/to/snapshot.sqlite.activation.json`

The manifest uses `schemas/snapshots/snapshot-manifest.schema.json`. The activation receipt uses `schemas/snapshots/activation-receipt.schema.json`.

## Lifecycle

```text
identified source artifacts
    -> validated ingestion
    -> immutable candidate database
    -> deterministic manifest
    -> checksum and integrity verification
    -> staged activation
    -> activation receipt
    -> read-only runtime
```

The manifest records source publishers, products, layers, vintages, locators, source checksums, retrieval timestamps, licensing review state, coordinate systems and transformations, builder identity, repository revision, record counts, output checksum, output size, and integrity checks.

The activation receipt records the active snapshot identity, activation time, prior active identity, last-known-good identity, and exact manifest checksum.

## Runtime behavior

The boundary snapshot is required. Startup fails closed when its manifest or activation receipt is missing, malformed, has the wrong snapshot kind, reports a failed integrity check, or does not match the database checksum and size.

The geocoder is optional. When its database exists but verification fails, the server leaves address geocoding unavailable and preserves coordinate-based resolution. It does not load an unidentified geocoder.

`GET /ready` reports only stable snapshot identities, never private absolute paths. A loaded database without an identity is reported as `unidentified`; an absent optional database is `unavailable`; a fully verified snapshot is `verified`.

## Activation and recovery

`backend/internal/snapshotmanifest.Activate` verifies the candidate before mutating the active location, writes staged database and sidecar files, verifies the staged set, then replaces the active files. The receipt retains the prior active identity and last-known-good identity.

A failed candidate verification leaves the active snapshot untouched. Replacement failures restore the individual prior file where possible and return an actionable error. Hydrated databases and sidecars remain outside Git unless their source and redistribution status are separately approved.

## Builder responsibilities

Snapshot builders must:

1. require source metadata rather than inventing it;
2. compute source and output SHA-256 checksums;
3. record accepted, rejected, duplicate, and quarantined counts;
4. run format-specific integrity checks;
5. write canonical manifest JSON;
6. produce an activation receipt only after candidate verification;
7. never replace the active snapshot before all validation passes.
