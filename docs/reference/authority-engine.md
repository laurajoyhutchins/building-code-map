# Authority Engine Contract

The Building Code Authority Engine is the transport-independent authority for location, jurisdiction, applicability, regulatory evidence, and bundle identity. It is owned by Building Code Map. Building Code AST remains responsible for parsing and navigating code-document content.

## Current stack base

The authority-engine stack currently descends from predecessor PR #48 at `0bc8c4b81996070a4568dd4896dea8ddc47a0c1f`. PR #45 supplies component snapshot manifests and activation identity, PR #47 supplies frontend runtime trust-boundary decoders, and PR #48 supplies geocoder ranking and interpolation provenance, including the completed frontend decoding and presentation path. The first engine pull request is temporarily dependent on that open predecessor stack and is not independently mergeable into `main` until those predecessors merge.

Frontend provenance behavior remains owned by PR #48. Authority-engine branches must not carry a second copy of those frontend changes. When a predecessor head changes, each descendant branch must be transplanted or rebased onto the exact new head, then re-verified before its prior evidence is reused.

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
