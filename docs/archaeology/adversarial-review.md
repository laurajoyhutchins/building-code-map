# Adversarial review

This review asks how Building Code Map could produce a plausible but unsupported answer even while the software behaves deterministically.

## Attacks the current architecture resists

### Caller supplies the desired jurisdiction

The public resolver no longer accepts normalized jurisdiction context as trusted input. It accepts a point and derives geographic observations from the admitted snapshot.

### Overlapping polygons produce an arbitrary winner

Peer state, county, and municipality overlaps return deterministic ambiguity with all tied observations. Encounter order is not treated as evidence.

### A malformed or semantically incomplete snapshot is admitted

Snapshot loaders validate the schema and invariants required by indexes and spatial resolution before making the capability ready.

### An unknown file is guessed to be DuckDB

SQLite is the workspace-local default. Legacy DuckDB is explicitly selected rather than inferred from an extension or machine-global fallback path.

### A future code edition appears current

Profiles preserve pending, optional, transition, effective, and mandatory states. The Oregon and North Carolina future records remain excluded until their conditions are supported.

### A pilot is called production-ready because tests are green

The production gate requires declared scopes, continuous timelines inside those scopes, primary evidence, source health, and evidence-backed fixtures. It also rejects premature verification.

## Remaining attack paths

### Correct state profile, wrong local authority

The repository lacks a canonical national legal-entity and local-authority registry. A state profile may correctly require local records while the system cannot yet prove that every candidate locality has been classified.

### Current geometry answers a historical query

Regulatory dates can be historical while address and boundary snapshots are current. Historical containment must remain unsupported unless dated geometry exists.

### Official source is stale, moved, or superseded

Source health checks cover declared pilot scopes, not all state and local records nationwide. A reachable source may also be legally obsolete.

### Scope boundary is misunderstood

A production-ready manifest can be misread as a statewide certification. User interfaces and reports must continue to expose the declared scope and unresolved classes.

### Special-area containment becomes authority by implication

Tribal, fire-service, flood, utility, campus, airport, and other special-purpose observations may be relevant without being adopting authorities. Only explicit policy and evidence may establish the relationship.

### Geocoder confidence becomes parcel certainty

An address point or high-confidence interpolation may still represent the wrong parcel, entrance, unit, or side of a jurisdiction boundary.

### Deterministic compiler preserves a bad interpretation

Determinism makes a rule reproducible, not legally correct. Accepted profiles still require skeptical source review and bounded claims.

## Required reviewer questions

1. What exact input and as-of date were used?
2. Was the point sourced or interpolated?
3. Were any boundary observations ambiguous?
4. Which legal rule turns containment into candidate authority?
5. Which instrument performs the adoption?
6. Are adoption, effective, mandatory, and repeal dates distinguished?
7. Is the result inside a declared production-ready scope?
8. What local records or special programs remain unresolved?
9. Are source identity, retrieval date, and freshness visible?
10. Would the same evidence support the answer if the map were removed?
