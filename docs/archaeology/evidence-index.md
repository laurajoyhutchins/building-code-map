# Evidence index

## Building Code Map

| Evidence | Role in archaeology |
| --- | --- |
| `88cea0da46bc6f72a9b628d7b7cec55d98b8d7d0` / PR #5 | Public README, quick start, delivered versus deferred architecture |
| PR #13, head `495fb62e873fddec7db98d700118d9cf53e86f6c` | Public governance, contribution, security, and data boundaries |
| `25d8ff5a345099b9d252899e129660763d10443b` / PR #30 | Canonical profiles, rule packs, resolver, three pilots, API, quality reporting |
| `0d7efe36ef26ff42b961f4b3f4c77d22c6e02345` / PR #32 | Map layer reconciliation correction |
| `286bd3f06454387008de7c1de4b0d8c4b4fffc7a` / PR #34 | Composite boundary identity correction |
| `a93a7aab56cb7d319c8e4d27fd683cf8b6de6e48` / PR #36 | Restrained public lookup and trust-signaling removal |
| `761ddc49bf6dffa50dc0454a5632f04fa3959594` / PR #37 | Local SQLite geocoder and active LORE trust root |
| PR #14, head `3006ebc65592acb9fa39d6e8db8a7746ab837a92` | Open report-conversion work, not current main |
| PR #35, head `0759d7618ba09f8c3fd4862d000a580fd3d5150c` | Superseded staged LORE readiness |
| PR #38, head `1cd4f25754381793a2ae3541b92d216dc82075a3` | Unmerged production-readiness manifests |
| PR #39, head `886affd903e68e61868b5b57b690f181ba5fe567` | Unmerged six-state research wave |
| Issues #4, #24–#31 | Product milestone, domain model, compilation, pilots, resolver, quality, readiness |
| Issues #7, #9, #11 | Boundary detail, snapshot validation, and path/provenance debt |
| `README.md` | Current declared architecture and roadmap |
| `DATA_SOURCES.md` | Source roles, rights, provenance, and runtime publication boundary |
| `backend/internal/geocoder/*` | Implemented normalization, schema, matching, interpolation, and ranking |
| `backend/internal/httpapi/resolve.go` | Implemented containment and point-based resolver boundary |
| `schemas/regulatory/resolution-result.schema.json` | Current executable result contract version 1.0 |
| `src/components/PublicLookup.tsx`, `src/public.css` | Current public product behavior and responsive visual implementation |
| `.lore/records/*`, `lore.yaml`, CI | Active LORE trust root and verification |
| CI run `30720401310` at PR #37 exact head | 28 frontend tests, 7 regulatory-tool tests, Go package tests/vet, build, LORE validation |

## External repositories

| Repository | Referenced decision ownership |
| --- | --- |
| `notactuallytreyanastasio/deciduous` | Canonical graph node, edge, schema, narrative, and export conventions |
| `laurajoyhutchins/LORE` | Accepted knowledge, proposals, transactions, extraction, projection, and repository contract |
| `laurajoyhutchins/building-code-ast` | Publication structure, source spans, tables, figures, definitions, exceptions, and provision ASTs |
| `laurajoyhutchins/electrical-equipment-lineage` | Manufacturer, product-family, UL, transaction, replacement, and compatibility claims |

## Evidence not promoted

Pull-request bodies, issue plans, README claims, and chat summaries were not treated as proof of
implementation without merge or tree evidence. Draft PRs are represented as proposed or historical.
No current deployment, production database, or nationwide coverage claim was inferred.
