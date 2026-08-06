# Evidence index

This index points reviewers to the current evidence surfaces that support the maintained archaeology projections. The frozen graph contains its own full-SHA evidence references.

## Runtime and trust boundaries

- `README.md`: delivered architecture, public API boundary, readiness semantics, pilot inventory, and roadmap.
- `backend/internal/httpapi/`: point-only public request contracts, composed lookup, ambiguity responses, and date-default warnings.
- `backend/internal/snapshot/`: snapshot discovery, semantic validation, SQLite default, and explicit compatibility paths.
- `backend/internal/regulatory/`: generic profile-driven resolver and rule application.
- `schemas/regulatory/resolution-result.schema.json`: current public result schema, version `1.0`.

The runtime hardening was merged through PR #41. Its exact verified branch head was `738b1ee89f45f2817e641ef71ab32c5bec030ad3` before squash merge.

## Regulatory profiles

Executable compiled profiles and source-health rules are under `backend/data/regulatory/`.

Current profile inventory:

- `colorado.json`
- `florida.json`
- `new-jersey.json`
- `virginia.json`
- `oregon.json`
- `north-carolina.json`

The second pilot wave was merged through PR #39 after reconciliation with the runtime and readiness changes. Its verified integration head was `c1be688a25dfd3fc72a14df7708665e41aec107d`.

Source reports are under `reports/pilots/`. Nationwide authority-structure, finding, and source research artifacts are under `reports/yaml/` and remain broader but less uniform than executable coverage.

## Production-readiness evidence

- `reports/production/`: declared pilot scopes for Colorado, Florida, and New Jersey.
- `reports/tools/production_readiness.py`: readiness audit.
- `reports/tools/test_production_readiness.py`: gate regressions.
- `docs/reference/pilot-production-readiness.md`: maintained interpretation of the gate and its limitations.
- `.github/workflows/official-source-health.yml`: scheduled official-source checking.

The readiness contract was merged through PR #38. Its verified branch head was `1cd4f25754381793a2ae3541b92d216dc82075a3`.

## Geocoding and boundary evidence

- `backend/internal/geocoder/schema.go`: geocoder schema version `1`.
- `backend/internal/geocoder/`: normalization, candidate ranking, ambiguity, interpolation, provenance, and snapshot construction.
- `backend/internal/snapshot/`: boundary feature loading and semantic validation.
- `docs/how-to/build-local-geocoder-snapshot.md`: reproducible local snapshot construction.
- `DATA_SOURCES.md`: source roles, licensing, attribution, and redistribution boundaries.

## Repository knowledge and archaeology

- `.deciduous/deciduous.sql.zlib.b85.part-*`: frozen Deciduous causal baseline.
- `.deciduous/narratives.md`: baseline narratives plus current reconciliation.
- `tools/validate_deciduous_archaeology.py`: frozen graph validator.
- `tools/validate_deciduous_current.py`: current profile and readiness inventory wrapper.
- `.lore/records/`: accepted semantic repository knowledge.
- `docs/generated/`: non-authoritative LORE projections.

## Verification evidence

The current CI runs four independent jobs:

- frontend type-check, lint, formatting, tests, and build;
- backend tests and vet;
- regulatory compilation, coverage, production gate, and report artifacts;
- Deciduous frozen-graph and current-inventory validation;
- LORE extraction and semantic validation.

Exact-head success proves that the tested tree is internally coherent. It does not prove that external regulatory evidence is complete or that every U.S. jurisdiction is implemented.
