# Authority Engine Contract

The Building Code Authority Engine is the transport-independent authority for location, jurisdiction, applicability, regulatory evidence, and bundle identity. It is owned by Building Code Map. Building Code AST remains responsible for parsing and navigating code-document content.

## Current stack base

The authority-engine stack currently descends from predecessor PR #48 at `d7908d9c4fc55f34d27f204310acc5cdd03919e7`. PR #45 supplies component snapshot manifests and activation identity, PR #47 supplies frontend runtime trust-boundary decoders, and PR #48 supplies geocoder ranking and interpolation provenance. The first engine pull request is temporarily dependent on that open predecessor stack and is not independently mergeable into `main` until those predecessors merge.

## Query contract

New engine callers provide exactly one location source and an explicit applicability date:

```json
{
  "point": { "longitude": -104.99, "latitude": 39.74 },
  "code_family": "building",
  "project_type": "new construction",
  "applicability_date": "2026-08-06",
  "include": ["evidence", "location"]
}
```

An address is geocoded through the local snapshot before geographic resolution. A point query does not invent an address. Coordinates are validated as finite longitude/latitude values within `[-180, 180]` and `[-90, 90]`. Query normalization trims text, uses stable lower-case keys, deduplicates and sorts include keys, and preserves the explicit date.

## Error codes

Engine errors are structured values with `code`, `message`, optional `details`, and `retryable`. Stable codes are:

`invalid_query`, `invalid_coordinates`, `address_not_found`, `address_ambiguous`, `boundary_ambiguous`, `outside_supported_coverage`, `regulatory_catalog_unavailable`, `regulatory_profile_missing`, `data_bundle_invalid`, and `internal_error`.

Legacy HTTP routes retain their current response shapes while the engine contract is introduced. Versioned transports will map these errors without changing their meaning.

## Provenance

Every result carries the source Git commit, engine version, aggregate bundle-manifest digest, boundary-snapshot digest, and regulatory-catalog digest. The geocoder-snapshot digest is present when address data was used. Geocoder candidates preserve the ranking-policy version, deterministic score factors, source identity and vintage, and interpolation derivation evidence established by PR #48.

## Determinism

Production code uses an injected `Clock`; tests use `FixedClock`. A fixed query, source data, clock, and bundle identity must produce byte-equivalent normalized results. Diagnostics are explicitly sortable by severity, code, path, and message. No query-time network access is permitted.

## CLI and bundle

The machine-facing CLI is `bcm`. Resolution requires an explicit applicability
date and writes only JSON to stdout:

```text
bcm resolve --point <longitude,latitude> --as-of YYYY-MM-DD
bcm resolve --address <address> --as-of YYYY-MM-DD
bcm geocode --address <address>
bcm lookup --point <longitude,latitude>
bcm inspect bundle
bcm inspect jurisdiction --id <id>
bcm serve --http 127.0.0.1:8000
```

`--pretty` is available for indented output. Exit code 0 means success, 1 an
unexpected internal failure, 2 invalid arguments, 3 an ambiguous/unresolved
outcome, and 4 invalid or unavailable bundle data.

The aggregate bundle manifest is content-addressed and identifies the exact
engine source, boundary snapshot, regulatory catalog, and optional geocoder
files. PR #45 component manifests remain the authority for component
activation and last-known-good identity; the aggregate manifest composes those
identities for a single engine result provenance record.
