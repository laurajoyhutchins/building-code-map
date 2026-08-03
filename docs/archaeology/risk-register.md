# Unresolved evidence, data-quality, and coverage risks

## Critical product risks

| Risk | Current evidence | Lifecycle |
| --- | --- | --- |
| Local applicability overclaim | Three pilot profiles remain partially verified; issue #31 is open | `incomplete-data` |
| Amendment incompleteness | Rule type exists, but municipal and historical amendment coverage is sparse | `incomplete-data` |
| Historical-date gaps | Resolver is date-aware, but complete adoption, operative, and supersession timelines are not present | `incomplete-data` |
| Geocoder source availability | Engine exists; production address sources and redistribution decisions do not ship | `incomplete-data` |
| Positional accuracy overstatement | Address points inherit source accuracy; interpolation is approximate | `active limitation` |
| Geometry versus legal authority | Polygon containment cannot establish adopting or enforcement roles | `active boundary` |
| Boundary snapshot provenance | Legacy machine-global path search and unknown-extension DuckDB fallback remain | `compatibility-only`, `unresolved` |
| Snapshot semantic validation | Issue #9 describes missing startup invariants; merge evidence was not found | `unresolved` |
| Boundary detail payload | Issue #7 remains evidence of an unresolved summary/detail split | `unresolved` |
| Point-on-boundary and overlap policy | Segment points count as inside; first county/municipality wins | `unresolved` |
| Temporal boundary change | Geometry is not historically versioned | `proposed` |
| Source staleness | Coverage tooling reports signals, but recurring refresh and enforcement are incomplete | `incomplete-data` |
| National support claim | Draft next-state work is unmerged; report presence is not executable support | `rejected claim` |
| AI-generated authority | RAG is not the final resolver; AI output must remain a proposal or research aid | `rejected architecture` |
| Code-text rights | Adopted editions are recorded without bundling protected publications | `active boundary` |
| Public site versus data maturity | The interface is usable while data remains pre-1.0 | `active qualification` |

## Claims downgraded during archaeology

- “Authority graph version 0.4” was not found as a current main contract. The confirmed executable
  result schema is `1.0`, and authority relationships are represented through candidates and
  `authority_path`.
- “Nationwide coverage” was downgraded to three executable, partially verified pilots plus broader
  research reports.
- “Rooftop geocoder” was downgraded to source-provided address points or labeled interpolation.
- “SQLite-only boundary runtime” was downgraded to SQLite-primary with legacy DuckDB compatibility.
- “LORE adoption began with PR #35” was corrected: PR #35 was staged and unmerged; PR #37 is the
  active merged trust root.
- “Trust badges communicate rigor” was reversed: current product demonstrates provenance in results
  and deliberately omits generic badge rows.
- “Green tests prove correctness” was qualified: they prove deterministic implementation behavior
  for fixtures, not complete legal or geographic data.

## Evidence still unresolved

- Exact earliest product naming and any private predecessor design not preserved in the public tree.
- Complete inventory of state research files versus their field-level freshness.
- Current production deployment URL and runtime snapshot vintage, because repository evidence does
  not establish deployment state.
- Whether issue #7, #9, or #11 has a later unmerged implementation outside the reviewed branches.
- Complete Go test-function count; standard CI reports package pass status rather than a test total.
- Independent accessibility audit and mobile-device verification beyond source and automated tests.
