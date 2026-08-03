# Comprehensive Repository Review: 2026-08-03

## Review conclusion

Building Code Map has the correct product center: a local-first, deterministic chain from location evidence to geographic observations to source-backed regulatory conclusions. Its strongest existing design decisions are the separation of geocoding from policy, the explicit regulatory evidence model, the refusal to treat pilot coverage as nationwide production coverage, the ignored runtime databases, and the compiler's fail-closed `--check` workflow.

Before this review, the implementation did not fully enforce that architecture. Public callers could submit normalized geographic context directly to `POST /resolve`, containing polygons were selected by encounter order for state, county, and municipality, syntactically valid but semantically invalid boundary snapshots could reach the HTTP layer, machine-global Windows caches could be selected implicitly, and readiness always returned a green response without describing unavailable capabilities.

This pass fixes those high-confidence authority and runtime-trust defects. The repository can now answer the core question conservatively when all of the following are true:

1. a semantically valid local boundary snapshot covers the input point;
2. address input, when used, resolves to one unambiguous provenance-bearing candidate;
3. state, county, and municipality containment is not ambiguous;
4. a validated regulatory profile supports the matched state, code family, project type, and date;
5. the result's unresolved records, warnings, evidence, and pilot limitations are preserved through the caller or interface.

It still cannot support a nationwide regulatory-coverage claim, complete historical applicability, or production-grade snapshot refresh and rollback. Those limits are explicit below.

## Skill routing

The review followed the repository owner's skill-routing contract rather than blending all guidance into one checklist.

| Phase or boundary | Owning skill | Result |
| --- | --- | --- |
| End-to-end source and snapshot trust | `data-engineering-design` | Primary owner; traced source observations through snapshots, API outputs, and publication boundaries |
| System authority and dependency direction | `architecture-review` | Enforced geocoder → boundary observation → regulatory policy separation |
| Backend changes | `go-idiomatic` | Added typed internal errors, deterministic ordering, read-only file access, focused packages, and tests |
| SQLite contract | `sql-idiomatic` | Validated schema-facing semantics and opened boundary SQLite snapshots read-only/query-only |
| HTTP behavior | `api-philosophy` | Made `/resolve` point-only, classified overlap ambiguity, added capability readiness and CORS preflight |
| Frontend contract assessment | `typescript-idiomatic` | Confirmed that compile-time interfaces currently substitute for complete runtime response validation; follow-up required |
| Operational evidence | `logging-philosophy` | Separated readiness evidence from regulatory conclusions; request correlation remains follow-up work |
| Repository state | `repo-config-governance` | Preserved pins, ignored snapshots, CI lanes, LORE authority, and Deciduous publication boundaries |
| Documentation | `diataxis-docs` | Updated overview, configuration, and API reference; added this review without turning it into a tutorial |
| Delivery | `writing-plans`, `test-driven-development`, `testing-philosophy`, `verification-before-completion`, `requesting-code-review`, `yeet` | Committed an implementation plan, added regression tests, used exact-head GitHub Actions, self-reviewed, and published draft PR #41 |

## Evidence inspected

The review inspected the current `main` branch at `761ddc49bf6dffa50dc0454a5632f04fa3959594`, recent commits and pull requests, open architecture and data-integrity issues, and the implementation areas listed below:

- repository identity, governance, data-source, security, contribution, license, notice, and configuration documents;
- frontend package, TypeScript, Vite, API-decoding, public lookup, map, and test contracts;
- Go server, HTTP routing, geocoder, boundary loader, regulatory catalog and resolver, and tests;
- SQLite geocoder schema, snapshot builder, ranking, interpolation, and read-only runtime behavior;
- legacy DuckDB compatibility;
- regulatory reports, compiler, validation, coverage reporting, generated profiles, and CI;
- LORE manifest, schemas, accepted records, generated projections, and pinned CI validation;
- Deciduous wrappers and the separately scoped archaeology work in draft PR #40;
- Windows launch and repository-governance tooling;
- open issues relevant to current authority and runtime behavior.

No hydrated boundary or address database was available through the connected GitHub environment. No result in this review claims live geographic coverage, source freshness, or production jurisdiction correctness from synthetic fixtures alone.

## Current-state authority model

### User input

A user may supply a civic address or WGS 84 coordinate pair. Original address input is preserved. Address normalization is a matching form, not a fabricated official mailing address.

### Geocoder observation

The local SQLite geocoder returns address-point or street-range candidates with source name, source record identity, source vintage, match precision, coordinates, and warnings. It does not own jurisdiction or regulatory policy.

### Boundary observation

The backend evaluates the point against loaded Polygon and MultiPolygon features. State, county, and municipality layers are treated as exclusive context inputs only when exactly one containing feature exists. Special-area, tribal-area, and NERIS layers remain non-exclusive geographic observations.

Containment is not legal authority. NERIS containment does not establish building-code, fire-code, or electrical-code authority. Census geography does not establish enforcement responsibility.

### Regulatory assertion

The internal regulatory resolver consumes normalized geographic context produced by the boundary stage. State profiles, rule packs, claims, adoptions, authorities, relationships, evidence, dates, conflicts, and required local records determine what can be stated.

### Presentation

The frontend displays a selected geocoding candidate and the regulatory result. It must retain warnings, evidence, unresolved records, and effective-date context. It must not present technical detail or a nationwide-looking map as proof of nationwide regulatory coverage.

## Findings resolved in this pass

### Critical: public context injection bypassed boundary evidence

**Previous behavior:** `POST /resolve` decoded `regulatory.ResolutionRequest`, including a caller-supplied `context`. A caller could assert state, municipality, incorporated status, special areas, and fire jurisdictions without a boundary observation.

**Resolution:** The HTTP endpoint now decodes a private point-only DTO. A `context` field is rejected by strict JSON decoding. Address lookup resolves only from the geocoder's selected point. The internal regulatory resolver may still accept normalized context from trusted repository code.

**Evidence:**

- `backend/internal/httpapi/resolve.go`
- `backend/internal/httpapi/geocode.go`
- `backend/internal/httpapi/resolve_test.go`

### Critical: overlapping authority-bearing polygons were selected arbitrarily

**Previous behavior:** state was overwritten by the last containing feature; county and municipality retained the first containing feature. Snapshot order could change the regulatory context.

**Resolution:** Matches are collected and sorted by stable feature identity. More than one containing state, county, or municipality returns HTTP `409` with code `boundary_ambiguous` and every tied observation. Special, tribal, and fire observations remain complete sorted arrays.

This is intentionally conservative. A future explicit authority rule may resolve a known overlap, but feature order may not.

### High: boundary snapshots lacked a semantic admission contract

**Previous behavior:** loaders proved that tables, JSON, and timestamps could be read. They did not reject duplicate identities, unknown layer references, unsupported geometry, non-finite or out-of-range coordinates, open rings, missing source identity, or unsupported refresh states. The HTTP feature index could silently overwrite duplicate records.

**Resolution:** `Snapshot.Validate()` now enforces:

- non-empty, unique, stable layer-family keys;
- non-empty, unique `(layer_family, feature_id)` identities;
- registered layer references;
- required title, source identity, geometry label, sync time, and object-shaped attributes;
- Polygon or MultiPolygon geometry;
- non-empty finite WGS 84 coordinates;
- at least four positions per ring and closed rings;
- supported refresh states and coherent timestamps.

Every supported loader calls the same validator before returning a snapshot.

### High: implicit snapshot discovery escaped the repository trust boundary

**Previous behavior:** default discovery could select `C:\tmp\tigerweb_hydrated.sqlite` or `C:\tmp\tigerweb_hydrated.duckdb`. Unknown extensions were treated as DuckDB, producing misleading compatibility-tool errors.

**Resolution:** The only default is `backend/data/tigerweb.sqlite`. Environment overrides must exist inside the repository checkout. Legacy DuckDB remains supported only through an explicitly selected `.duckdb` path. Unknown extensions return `ErrUnsupportedSnapshotFormat` before DuckDB discovery.

SQLite boundary snapshots are opened read-only and query-only.

### High: readiness overstated available behavior

**Previous behavior:** `/ready` always returned `{"status":"ok"}`. It did not distinguish boundary data, geocoding, or regulatory profiles.

**Resolution:** `/ready` now reports `ready`, `degraded`, or `not_ready` plus capability records for boundary resolution, coordinate resolution, address geocoding, regulatory resolution, and composed address lookup. Required boundary unavailability returns `503`; optional capability loss returns `200` with `degraded`.

### High: omitted applicability dates were silent assumptions

**Previous behavior:** the regulatory resolver defaulted an omitted date to the current UTC date without identifying that assumption in the public result.

**Resolution:** Public point and address resolution append a warning that identifies the server-selected UTC date and directs the caller to confirm the governing project date. The optional field remains compatible, but the assumption is no longer invisible.

### Medium: configured cross-origin JSON POST requests lacked preflight behavior

**Previous behavior:** normal allowed-origin responses included `Access-Control-Allow-Origin`, but browser `OPTIONS` requests fell through to `404`.

**Resolution:** allowed preflight requests return `204` with explicit methods, `Content-Type`, cache age, and `Vary`; unconfigured origins receive `403`. CORS is documented as a browser boundary, not authentication.

## Strong existing controls retained

### Regulatory data model and compiler

The repository already preserves authorities, relationships, adoptions, rules, claims, conflicts, verification records, evidence sources, access dates, and source health. The compiler validates pilot reports, produces deterministic output, and supports `--check` without rewriting tracked data. CI executes unit tests, report validation, stale-output detection, and nationwide quality reporting.

This pass did not weaken `partially_verified`, `local_record_required`, `conflicting`, or `insufficient_evidence` states to make demonstrations look complete.

### Geocoder composition

Address points are preferred over ranges. Ranges are parity-aware, support reversed ranges, use deterministic source-record ordering, and emit an interpolation warning. Snapshot replacement uses a temporary database and restores the prior output when replacement fails.

### Publication boundaries

Hydrated SQLite and DuckDB files remain ignored. Address data, Census products, NERIS records, OpenStreetMap tiles, statutes, regulations, and copyrighted code text retain independent licensing and attribution boundaries. The Apache 2.0 repository license does not relicense third-party data.

### LORE

Accepted records remain under `.lore/records/`; deterministic extracts and generated documentation remain non-authoritative projections. Pinned LORE CI validates the trust root. This pass adds an append-only decision revision rather than rewriting bootstrap history.

### Deciduous

The repository contains wrappers for work, recovery, and audit. The comprehensive archaeology graph is still isolated in draft PR #40. This pass did not create a second graph or commit prompt-bearing generated exports.

## Material work that remains

### Snapshot provenance, activation, and rollback identity

The current boundary snapshot stores layers, features, and one refresh-status row. It does not yet carry a complete source manifest, source artifact checksums, schema version, builder version, repository revision, transformation history, record counts, snapshot checksum, activation receipt, last-known-good identity, or rollback history.

The geocoder metadata table records schema version, one source name, and one source vintage. It does not identify input file checksums, build tool revision, build timestamp, row counts, coordinate reference system declaration, quarantine counts, output checksum, or activation receipt.

**Required next result:** define one versioned manifest contract per snapshot family, validate it before activation, build candidate snapshots in staging, run integrity and data-quality checks, atomically activate by identity, and retain a bounded last-known-good rollback reference.

### Geocoder positional provenance and ranking contract

An interpolated candidate identifies the source range record but does not return side of street, parity, original endpoints, normalized range direction, interpolation fraction, or the explicit scoring factors that produced confidence. Source priority is implicit rather than a versioned source-specific rule.

**Required next result:** extend candidate provenance and define confidence as a documented quality classification or score decomposition, not an unexplained decimal that downstream interfaces may mistake for probability.

### Frontend runtime validation

`src/lib/api.ts` uses TypeScript interfaces and generic JSON casts. It transforms snake-case payloads but does not validate all required fields, discriminants, coordinate ranges, URLs, dates, array members, or error envelopes at runtime.

**Required next result:** validate every independently versioned backend response at the network boundary, reject malformed payloads with actionable errors, and test missing, extra, invalid, and forward-compatible fields. Compile-time types must remain a developer aid rather than evidence that network data is valid.

### Frontend ambiguity and degraded-mode presentation

The backend now returns structured boundary ambiguity and capability readiness, but the frontend's generic error path does not yet render tied boundary observations or the readiness matrix. Technical explorer detail can also feel more authoritative merely because it is more detailed.

**Required next result:** present boundary ambiguity as evidence requiring selection or confirmation, distinguish geocoder and regulatory outages, define every verification label, and ensure uncertainty is conveyed through text and structure rather than color or an icon alone.

### Historical validity

Profiles support applicability dates and fail closed where current records do not prove historical applicability. Coverage is not complete for enactment, administrative adoption, effective, operative, mandatory, transition, repeal, supersession, project-grandfathering, and source-observation timelines.

**Required next result:** retain append-only adoption timelines and explicit unsupported historical intervals. Do not infer a historical answer from the latest source.

### API surface and operational evidence

Remaining API work includes:

- list/detail and pagination redesign for `/boundaries`, tracked by issue #7;
- versioned error envelopes and stable machine codes beyond boundary ambiguity;
- request timeouts and cancellation evidence at expensive boundaries;
- request correlation and structured logs carrying snapshot and profile identities;
- safe limits for bounding boxes and future geometry queries;
- explicit API versioning policy;
- refresh authorization before mutation is enabled.

`POST /refresh/trigger` remains disabled and therefore does not currently expose arbitrary ingestion.

### Windows process ownership

Issue #12 documents that launcher cleanup can terminate an unverified port owner. This pass did not modify process lifecycle scripts because that work is independent from regulatory and snapshot authority and needs a focused PowerShell ownership test harness.

### Repository protection and external contribution posture

Issue #3 remains the owner for branch rules, required checks, dependency alerts, merge methods, and public contribution boundaries. CI is healthy, but repository-platform policy is separate from code correctness.

## Testing assessment

| Area | Current evidence | Important missing evidence |
| --- | --- | --- |
| Address normalization | Unit fixtures cover ordinary civic addresses and selected aliases | Rural routes, highway forms, fractions, leading zero policy, locality alias governance, Unicode property cases |
| Address points | SQLite integration tests and deterministic ranking | Multiple source-priority policies, duplicate coordinates across publishers, stale-source arbitration |
| Street interpolation | Parity, reversed ranges, zero-length rejection, rollback-safe build | Side of street, endpoint and fraction response provenance, positional uncertainty bands |
| Boundary geometry | Polygon holes, semantic validator, overlap ambiguity | Enclaves, coastal topology, invalid-geometry repair policy, large multipolygons, antimeridian rejection or support |
| Regulatory resolution | Three contrasting pilot states, date failure, authority paths, conflicts, source evidence | Complete local/home-rule fixtures, historical timelines, state-owned and special occupancy breadth |
| Snapshots | SQLite loader integration, semantic negative tests, explicit formats | Manifest/checksum/activation/rollback integration and corruption recovery |
| HTTP | Strict JSON, body limit, coordinate ranges, point-only contract, CORS, readiness | Stable general error schema, timeout/cancellation, rate and geometry-size limits |
| Frontend | Type-check, lint, Vitest, build, geocoder provenance display | Runtime response validation, boundary ambiguity UI, readiness/degraded UI, accessibility audit for uncertainty |
| Regulatory compiler | Python unit tests, report validation, deterministic `--check`, coverage report | Primary-source health automation and complete production-ready state fixtures |
| LORE | Pinned extract and validate CI | Transactional proposal workflow exercised for this repository, projection freshness display in UI/docs |
| Deciduous | Wrappers and separate archaeology PR | Merge and audit the canonical graph after PR #40 review; avoid duplicate backfill |

## Cross-repository boundaries and follow-up

No other repository was modified.

### Building Code AST

Building Code Map should eventually emit stable adopted-instrument identifiers, family, edition, amendment-set identity, scope, and date context that Building Code AST can resolve to structured corpus records. Building Code AST remains responsible for code text, document structure, tables, figures, references, and requirement assertions.

### Building Code Dashboard

Before expanding presentation surfaces here, inspect and document the dashboard's current product contract. A portfolio or nationwide analysis view should not be duplicated inside the local lookup product.

### Revit Agent

A future consumer contract should pass source-backed authority and adoption results with uncertainty and as-of date. This repository must not mutate Revit models or decide project compliance.

### Electrical Equipment Lineage

Adopted electrical-code evidence may provide context to equipment research, but it cannot establish listing, compatibility, replacement acceptability, or suitability for a specific installation.

### LORE

LORE owns accepted knowledge transaction semantics. Building Code Map may consume a pinned LORE toolchain and maintain repository-local records, but it should not redefine the LORE kernel or treat LORE validation as software, GIS, or regulatory verification.

## Verification evidence

The first CI run exposed a Go syntax defect in `backend/internal/httpapi/resolve.go`; frontend, regulatory-data, and LORE jobs passed. The defect was corrected and invalidated checks were rerun.

At implementation head `742909b252a1c548a4d243fecdee463aa40bc8ae`, GitHub Actions run `30855496982` completed successfully:

- Backend: `go test ./...` passed; `go vet ./...` passed.
- Frontend: locked install, TypeScript check, ESLint, Prettier check, Vitest, and Vite build passed.
- Regulatory data: Python tests, pilot validation, compiled-profile `--check`, and coverage generation passed.
- LORE documentation: pinned LORE build, extraction, and validation passed.

Subsequent documentation and semantic-record commits require their own exact-head CI result before the pull request can be described as fully verified.

## Final assessment

Building Code Map is not yet a nationwide adopted-code authority. It is a credible pre-1.0 evidence and resolution system for loaded geography and supported pilot policy, provided callers respect its warnings and local-confirmation requirements.

The repository's next quality frontier is not a larger map or more prose. It is stronger snapshot identity, richer positional provenance, runtime contract validation, historical timelines, and explicit authority mappings. Those additions should deepen auditability without collapsing geography, law, enforcement, code content, or professional interpretation into one convenient but unreliable answer.
