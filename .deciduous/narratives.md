# Building Code Map Decision Narratives

These narratives are root-level reading paths into the canonical Deciduous graph in
`.deciduous/deciduous.sql`. They summarize causal arcs; they do not replace node evidence.

## 1. From “which code?” to inspectable applicability

Building Code Map began with a deceptively compact engineering question: given a site, which
authorities and adopted codes apply? Model-code publication, statewide summaries, municipal
adoption, enforcement responsibility, amendments, special districts, and effective dates are
different facts. A map or city-to-edition table could display one dimension, but it could not show
why an edition applies, who enforces it, or what remains unknown. The project therefore became an
applicability resolver whose useful answer may be “local record required” rather than a convenient
edition.

Root: `bcm-goal-inspectable-location-applicability`

## 2. The address-to-result pipeline was separated to preserve evidence

The user-facing contract eventually became address or coordinates in, applicable authority and
adoption records out. That experience is implemented as explicit stages: normalize a civic address,
geocode it to a provenance-bearing point, match supported geometries, interpret authority policy,
filter adoption records by context and date, and render evidence and warnings. `POST /resolve`
remains point-based. `POST /lookup` composes geocoding with the same resolver. This boundary makes
a weak address match visible instead of allowing it to masquerade as a regulatory conclusion.

Root: `bcm-goal-address-to-adopted-code-contract`

## 3. Local-first means local execution over refreshable public data

The architecture is local-first because address privacy, deterministic replay, vendor outages,
per-query cost, source inspection, and engineering reproducibility matter. It is not a pledge never
to use networked public sources. TIGER/Line, reviewed address data, NERIS resources, statutes,
regulations, ordinances, and agency publications enter ingestion or research workflows. Once
versioned artifacts exist, lookup does not depend on a geocoding or AI service. Dataset size,
licensing, setup, refresh cadence, and stale-data detection are the price of that independence.

Root: `bcm-goal-deterministic-local-first`

## 4. RAG was rejected as final authority, not as a research aid

A retrieval-augmented answer can miss the controlling instrument, retrieve commentary instead of
law, confuse publication with adoption, overlook supersession, and return a plausible result for
the wrong jurisdiction. Those failure modes are unacceptable as the final authority-resolution
engine. The repository instead compiles declarative policy and resolves it deterministically.
AI-assisted research, extraction, comparison, review, and explanation can still help produce or
check proposals, but accepted data and validated source relationships govern the result.

Root: `bcm-goal-deterministic-local-first`

## 5. Geocoding moved from a deferred assumption to a bounded SQLite component

The first usable milestone intentionally accepted coordinates because address normalization and
data acquisition were separate problems. The later local geocoder added a conservative U.S.
civic-address parser, address-point preference, locality-constrained indexed queries, deterministic
ranking, ambiguity handling, and parity-aware street-range interpolation. Snapshot schema version 1
uses `geocoder_metadata`, `address_points`, and `street_ranges`. “Address point” describes the input
record type, not independently verified rooftop accuracy; interpolation is explicitly approximate.
Units, parcels, landmarks, intersections, rural-route specialization, reverse geocoding, and
historical address snapshots remain deferred.

Root: `bcm-goal-local-address-front-door`

## 6. Geographic containment is evidence, not legal authority

A coordinate can be contained by several nested and overlapping polygons. State, county,
municipality, tribal area, fire jurisdiction, flood or special area, and utility territory do not
automatically have the same adopting or enforcement role. The boundary resolver therefore produces
geographic context, while state policy profiles identify authority candidates, roles, relationships,
code-family overrides, and local-record requirements. The current result contract is
`resolution-result` schema 1.0. No separately versioned current `authority_graph` schema was found
on `main`; authority is expressed through candidates and `authority_path`.

Root: `bcm-goal-resolve-overlapping-jurisdictions`

## 7. Adoption records preserve dates, amendments, conflicts, and missing evidence

A model-code edition can have a publication date, an adopting instrument can have an adoption date,
and the law can have a different effective, operative, mandatory, replacement, or transition date.
Statewide minimums can coexist with local amendments or local adoption. The resolver therefore
supports date-aware records, amendment and enforcement rule kinds, verification state, claims,
conflicts, required local records, and warnings. The representation is more capable than the current
dataset: complete municipal amendments, optional appendices, design criteria, climatic values, and
historical supersession remain major gaps.

Root: `bcm-goal-date-aware-adoption-records`

## 8. Three pilots validated the model without establishing national coverage

Colorado, Florida, and New Jersey were selected because they exercise contrasting authority
patterns: predominantly local adoption with a state electrical override, a statewide building-code
baseline with local enforcement and amendment follow-up, and statewide construction subcodes with a
separate operational-fire path. They are executable profiles on `main`, but intentionally remain
partially verified. Draft PRs for stricter production-readiness manifests and a six-state research
wave are not current support. A state report, a draft profile, and complete local coverage are three
different maturity levels.

Root: `bcm-goal-expand-state-coverage-with-thresholds`

## 9. Source roles and publication rights shaped runtime boundaries

Census geometry identifies places but does not establish code adoption. Official statutes,
regulations, ordinances, and agency records can support legal claims. Secondary summaries help
research but are not substituted for controlling records. Model-code publishers identify
publications without proving local adoption. Address and fire-service data have separate ownership,
cadence, and redistribution questions. Hydrated boundary and geocoder databases stay out of Git;
the repository publishes reproducible logic, schemas, fixtures, provenance requirements, and
carefully bounded records instead.

Root: `bcm-goal-govern-source-classes`

## 10. Building Code Map and Building Code AST form a pipeline, not one product

Building Code Map owns location, jurisdictions, authority relationships, adoptions, amendments,
effective dates, and applicability. Building Code AST owns faithful structure of selected
publications: hierarchy, definitions, exceptions, cross-references, tables, figures, source spans,
and diagnostics. The intended composition is location to authority and adopted source, then source
artifact to document AST and selected provision. Electrical Equipment Lineage and future compliance
tools may consume applicability context, but each retains its own evidence and professional-review
boundary.

Root: `bcm-goal-repository-boundaries`

## 11. LORE was adopted as a trust protocol rather than recreated locally

An earlier draft prepared repository-local LORE-compatible scaffolding while the upstream workflow
was stabilizing. The merged implementation instead pins LORE, uses its shipped maintenance skill,
stores accepted semantic records, reserves proposals and transaction receipts, extracts
deterministic facts, and treats generated documents as projections. PR #37, not the earlier staged
draft, is the active trust-root implementation.

Root: `bcm-goal-durable-repository-knowledge`

## 12. The website learned to demonstrate trust instead of announcing it

The technical GIS explorer was not an adequate public front door. The public root became a
search-first lookup while `/explorer` retained the workbench. Generic trust badges, promotional
claims, decorative metrics, and slogan-like assurance copy were removed. Current results demonstrate
credibility through matched-address provenance, authorities, adopted-code records, effective dates,
local confirmation items, jurisdiction relationships, and dated source links. Necessary caveats
remain, but they are not used as decorative credibility tokens.

Root: `bcm-goal-public-product-interface`

## 13. Compatibility debt remains part of the current architecture

SQLite is the documented primary local snapshot format, but the boundary loader still retains
DuckDB support, treats unknown extensions as DuckDB, and searches legacy machine-global `C:\tmp`
candidates. Map hydration ordering and composite boundary identity defects were fixed, but snapshot
semantic validation, detail loading, path provenance, temporal geometry, and some overlap behavior
remain unresolved. These mechanisms are compatibility-bound, not evidence that the target
architecture was always DuckDB-shaped.

Root: `bcm-goal-preserve-compatible-runtime-while-maturing`

## 14. Public source availability is not production-grade data coverage

The public README, read-only CI, contribution rules, and source-publication boundaries made the
repository understandable and safer to contribute to. They did not make the datasets nationwide or
legally complete. Building Code Map remains pre-1.0 research and engineering. The public site and
green tests prove that the implemented paths are coherent, not that every jurisdiction or amendment
has been verified.

Root: `bcm-goal-public-repository-credibility`
