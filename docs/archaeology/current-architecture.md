# Current architecture projection

## Active request path

```text
address input
  -> conservative civic-address normalization
  -> local SQLite geocoder schema 1
  -> selected address point or labeled street-range interpolation
  -> WGS 84 point
  -> local boundary snapshot
  -> normalized geographic context
  -> versioned state profile and rule pack
  -> generic regulatory resolver
  -> resolution-result schema 1.0
  -> public lookup or technical explorer
```

Coordinate input enters at the WGS 84 point and remains usable when no geocoder snapshot is loaded.

## Active components

### Public application

The root route is a search-first public lookup. It accepts a U.S. civic address or coordinates, a
code family, and an applicability date. Results show geocoder provenance when applicable, authority
candidates, adopted-code records, special conditions, local-confirmation items, jurisdiction
relationships, and dated source links. `/explorer` retains the GIS workbench.

### Local geocoder

The Go geocoder uses a separate SQLite snapshot with schema version `1`. Address points and street
ranges are indexed by normalized fields and source identities. Ranking is deterministic. Address
points require a minimum confidence of `0.85`; street ranges require `0.75`; top candidates within
`0.05` are ambiguous. Those numbers are ranking thresholds, not calibrated probability estimates.

### Boundary resolver

The local snapshot exposes states, counties, municipalities, special areas, tribal areas, and NERIS
jurisdictions. Polygon and multipolygon containment includes holes. A point on a ring segment is
treated as contained. The current implementation keeps the first county and municipality match while
accumulating special, tribal, and fire matches.

### Regulatory resolver

The generic resolver consumes normalized geography plus optional code family, project type, and
applicability date. State behavior lives in data, not a state-name switch. The result can be
`resolved`, `partially_resolved`, `local_record_required`, `ambiguous`, `conflicting`, or
`insufficient_evidence`.

### Regulatory data

`main` contains executable pilot profiles and rule packs for Colorado, Florida, and New Jersey.
Their differing authority structures validate the data model, but the profiles remain partially
verified and do not establish full municipal amendment or historical coverage.

### Knowledge and verification

LORE is pinned in CI and validates accepted repository records. Deciduous wrappers support local
decision-graph recovery and audit. The archaeology adds a schema-native seed, generated graph export,
and deterministic validator.

## Compatibility-bound mechanisms

- Boundary snapshots can be SQLite or legacy DuckDB.
- Unknown boundary snapshot extensions currently fall back to DuckDB.
- Default boundary-path discovery still includes legacy machine-global temporary locations.
- The technical explorer remains alongside the public product.
- Existing Deciduous Windows wrappers locate a separately installed executable.

These are current behaviors, not necessarily target architecture.

## Proposed or unmerged

- Production-readiness manifests and source-health checks in draft PR #38.
- Additional state inventories and draft profiles in PR #39.
- Stricter report-schema conversion work in draft PR #14.
- Broader local adoption, amendment, and historical records.
- Automated boundary and geocoder refresh.
- PostGIS and vector-tile serving if scale requires them.
- Unit, parcel, landmark, intersection, reverse, and historical geocoding.

## Not implemented by this archaeology

No geocoder behavior, API schema, jurisdiction data, deployment, website behavior, production
database, external service, Linear record, release, or merge state is changed.
