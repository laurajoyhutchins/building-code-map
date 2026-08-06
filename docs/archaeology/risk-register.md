# Risk register

## Active high-priority risks

### No national governmental-entity spine

Boundary features and state profiles do not yet form a canonical nationwide registry of legal entities, adoption powers, enforcement assignments, inheritance, and validity intervals. This prevents complete jurisdiction enumeration and reliable national completion metrics.

### Local adoption and amendment gaps

The six executable profiles expose required local records but do not comprehensively cover every municipality, county, township, amendment, enforcement contract, special occupancy, or delegated program.

### Historical incompleteness

Selected regulatory transitions are modeled, but historical addresses, boundaries, authorities, adoptions, and amendments are not complete nationwide. Current geometry must not silently answer a historical question.

### Snapshot lifecycle remains incomplete

Snapshots are semantically validated before admission, but complete source manifests, checksums, build identities, atomic activation, rollback identity, and historical snapshot catalogs remain follow-up work.

### Frontend runtime decoding

TypeScript types do not prove that a runtime API payload conforms. The frontend still needs comprehensive decoding and explicit degraded or ambiguous evidence rendering.

### Geocoder provenance depth

The geocoder distinguishes address points from interpolation and preserves source metadata, but source-specific ranking, complete interpolation derivation, unit handling, parcel linkage, and historical address identity remain incomplete.

### Source freshness and legal change

The production gate checks declared sources, but nationwide monitoring does not yet exist. Moved links, superseding instruments, annexations, agency reorganizations, and amended local ordinances can make a once-correct result stale.

## Recently reduced risks

### Caller-authored authority context

Reduced by making public `/resolve` point-only and deriving geographic observations server-side.

### Encounter-order boundary selection

Reduced by returning deterministic ambiguity for overlapping peer state, county, and municipality observations.

### Invalid snapshot admission

Reduced by semantic validation before runtime use.

### Machine-global snapshot guessing

Reduced by using a workspace-local SQLite default and requiring explicit legacy DuckDB configuration.

### Undifferentiated readiness

Reduced by capability-specific readiness and by the scoped pilot production gate.

### Silent date defaulting

Reduced by returning an explicit warning when the server supplies the current UTC applicability date.

## Non-risks that must not be promoted into claims

- six executable profiles do not mean six complete states;
- three production-ready scopes do not mean three complete statewide datasets;
- a green pipeline does not establish legal correctness;
- a containing polygon does not establish adoption or enforcement authority;
- a state research artifact does not make the state executable;
- public availability does not make hydrated source data redistributable.
