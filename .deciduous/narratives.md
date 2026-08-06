# Building Code Map decision narratives

These narratives are reading paths through the frozen Deciduous graph. The graph records the repository's causal baseline when archaeology began. The reconciliation below records later merged work without rewriting historical nodes.

## Current reconciliation

Since the baseline was reconstructed, Building Code Map has:

- hardened public resolution so callers provide a point rather than trusted jurisdiction context;
- made overlapping state, county, and municipality observations deterministic ambiguity instead of encounter-order selection;
- admitted boundary snapshots only after semantic validation;
- replaced broad legacy snapshot guessing with a workspace-local SQLite default and explicit DuckDB compatibility;
- added capability-specific readiness and visible applicability-date defaulting;
- established scoped production-readiness manifests for Colorado, Florida, and New Jersey;
- expanded executable regulatory profiles to Virginia, Oregon, and North Carolina.

The repository therefore has six executable pilot profiles and three production-ready **declared scopes**. Neither statement means complete statewide, municipal, historical, or nationwide coverage.

## 1. From “which code?” to inspectable applicability

Building Code Map began with a compact question: given a site, which authorities and adopted codes may apply? Publication, adoption, amendment, enforcement, project scope, and effective date are distinct facts. The product therefore became an applicability resolver whose correct result may require a local record rather than fabricate a convenient edition.

Root: `bcm-goal-inspectable-location-applicability`

## 2. The address-to-result pipeline preserves evidence

Address normalization, geocoding, point provenance, boundary observation, regulatory interpretation, and result rendering are explicit stages. `POST /lookup` composes geocoding with the point-only resolver. A weak or ambiguous location match remains visible and cannot silently become a legal conclusion.

Root: `bcm-goal-address-to-adopted-code-contract`

## 3. Local-first means deterministic execution over reviewable artifacts

Public datasets and official legal sources enter acquisition and research workflows. Lookup uses local, versioned artifacts rather than depending on a live geocoder, vendor, or language model. That choice improves privacy and replayability while creating obligations for source licensing, snapshot identity, refresh, and stale-data detection.

Root: `bcm-goal-deterministic-local-first`

## 4. Retrieval is research assistance, not final authority

Retrieval can miss controlling instruments, confuse commentary with law, overlook supersession, or answer for the wrong jurisdiction. Accepted declarative policy and verified source relationships govern deterministic results. AI may assist research, extraction, comparison, and review, but it does not become the authority surface.

Root: `bcm-goal-deterministic-local-first`

## 5. Geocoding became a bounded SQLite component

The initial coordinate-first milestone avoided pretending address resolution was solved. The later geocoder added conservative civic-address parsing, address-point preference, deterministic ranking, explicit ambiguity, and parity-aware street-range interpolation. Address points and interpolated points carry different precision and provenance.

Root: `bcm-goal-local-address-front-door`

## 6. Geographic containment is not legal authority

A point may intersect nested or overlapping governmental and special-purpose polygons. The boundary layer reports observations. Regulatory profiles interpret candidate authority, code-family scope, inheritance, enforcement, and local-record requirements. Multiple peer state, county, or municipality observations now return explicit ambiguity instead of selecting the first row.

Root: `bcm-goal-resolve-overlapping-jurisdictions`

## 7. Adoption records preserve time and uncertainty

Publication date, adoption date, effective date, mandatory date, transition period, repeal, and source observation are not interchangeable. Profiles and rules preserve pending instruments, verification state, conflicts, special conditions, and required local records. Historical and municipal depth remains incomplete even where current pilot behavior is executable.

Root: `bcm-goal-date-aware-adoption-records`

## 8. Pilot waves test regimes without claiming national completeness

Colorado, Florida, and New Jersey established the first executable profiles and now have production-ready declared scopes. Virginia, Oregon, and North Carolina add statewide-code, specialty-code, local-enforcement, and conditional-transition patterns. All six remain pilot implementations, and the three readiness determinations apply only to their declared manifestations and fixtures.

Root: `bcm-goal-expand-state-coverage-with-thresholds`

## 9. Source roles and publication rights shape runtime boundaries

Census geometry can identify a place without proving legal authority. Statutes, regulations, ordinances, and official agency records support legal claims. Secondary summaries aid discovery. Model-code publications do not prove adoption. Hydrated address and boundary databases remain outside Git unless their provenance and redistribution terms permit publication.

Root: `bcm-goal-govern-source-classes`

## 10. Building Code Map and Building Code AST are adjacent products

Building Code Map owns location, candidate authorities, adopted instruments, amendments, dates, and applicability uncertainty. Building Code AST owns faithful structure of selected publications, including hierarchy, cross-references, source spans, and diagnostics. Their composition is location to governing source, then source to provision, without merging their evidence boundaries.

Root: `bcm-goal-repository-boundaries`

## 11. LORE is the durable knowledge trust protocol

The repository pins LORE, stores accepted semantic records, extracts deterministic facts, and treats generated documentation as projections. Deciduous preserves causal development history; LORE preserves accepted semantic repository knowledge. Neither replaces source-backed regulatory records.

Root: `bcm-goal-durable-repository-knowledge`

## 12. The website demonstrates evidence rather than announcing trust

The public root became a search-first lookup while the explorer remained available as a technical workbench. Generic trust slogans were removed. Results instead expose matched-location provenance, authority candidates, adoption records, effective dates, warnings, local confirmation items, and dated sources.

Root: `bcm-goal-public-product-interface`

## 13. Compatibility is now bounded rather than guessed

SQLite is the workspace-local default for boundary snapshots. Legacy DuckDB remains available only through explicit configuration. Snapshot loaders reject unsupported formats and invalid semantic content before runtime admission. Full manifest, activation, rollback, and historical-boundary contracts remain follow-up work.

Root: `bcm-goal-preserve-compatible-runtime-while-maturing`

## 14. Public availability is not nationwide production coverage

The public repository, green CI, six executable pilots, and three production-ready scopes establish that implemented paths are coherent and reviewable. They do not establish complete legal coverage for every municipality, special authority, occupancy, amendment, or historical date. Building Code Map remains pre-1.0 research and engineering.

Root: `bcm-goal-public-repository-credibility`
