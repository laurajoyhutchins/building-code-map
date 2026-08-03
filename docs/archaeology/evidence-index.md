# Evidence index

## Building Code Map

- `88cea0da46bc6f72a9b628d7b7cec55d98b8d7d0` / PR #5: public README,
  quick start, and delivered-versus-deferred architecture.
- PR #13, head `495fb62e873fddec7db98d700118d9cf53e86f6c`: public governance,
  contribution, security, and data boundaries.
- `25d8ff5a345099b9d252899e129660763d10443b` / PR #30: canonical profiles,
  rule packs, resolver, three pilots, API, and quality reporting.
- `0d7efe36ef26ff42b961f4b3f4c77d22c6e02345` / PR #32: map-layer
  reconciliation correction.
- `286bd3f06454387008de7c1de4b0d8c4b4fffc7a` / PR #34: composite boundary
  identity correction.
- `a93a7aab56cb7d319c8e4d27fd683cf8b6de6e48` / PR #36: restrained public
  lookup and trust-signaling removal.
- `761ddc49bf6dffa50dc0454a5632f04fa3959594` / PR #37: local SQLite geocoder
  and active LORE trust root.
- PR #14, head `3006ebc65592acb9fa39d6e8db8a7746ab837a92`: open report-conversion
  work, not current main.
- PR #35, head `0759d7618ba09f8c3fd4862d000a580fd3d5150c`: superseded staged LORE
  readiness.
- PR #38, head `1cd4f25754381793a2ae3541b92d216dc82075a3`: unmerged
  production-readiness manifests.
- PR #39, head `886affd903e68e61868b5b57b690f181ba5fe567`: unmerged six-state
  research wave.
- Issues #4 and #24 through #31: product milestone, domain model, compilation,
  pilots, resolver, quality, and readiness.
- Issues #7, #9, and #11: boundary detail, snapshot validation, and
  path/provenance debt.
- `README.md`: current declared architecture and roadmap.
- `DATA_SOURCES.md`: source roles, rights, provenance, and runtime publication
  boundary.
- `backend/internal/geocoder/*`: implemented normalization, schema, matching,
  interpolation, and ranking.
- `backend/internal/httpapi/resolve.go`: implemented containment and point-based
  resolver boundary.
- `schemas/regulatory/resolution-result.schema.json`: current executable result
  contract version 1.0.
- `src/components/PublicLookup.tsx` and `src/public.css`: current public product
  behavior and responsive visual implementation.
- `.lore/records/*`, `lore.yaml`, and CI: active LORE trust root and
  verification.
- CI run `30720401310` at the PR #37 exact head: 28 frontend tests, 7
  regulatory-tool tests, Go package tests and vet, build, and LORE validation.

## External repositories

- `notactuallytreyanastasio/deciduous`: canonical graph node, edge, schema,
  narrative, and export conventions.
- `laurajoyhutchins/LORE`: accepted knowledge, proposals, transactions,
  extraction, projection, and repository contract.
- `laurajoyhutchins/building-code-ast`: publication structure, source spans,
  tables, figures, definitions, exceptions, and provision ASTs.
- `laurajoyhutchins/electrical-equipment-lineage`: manufacturer,
  product-family, UL, transaction, replacement, and compatibility claims.

## Evidence not promoted

Pull-request bodies, issue plans, README claims, and chat summaries were not
treated as proof of implementation without merge or tree evidence. Draft PRs
are represented as proposed or historical. No current deployment, production
database, or nationwide coverage claim was inferred.
