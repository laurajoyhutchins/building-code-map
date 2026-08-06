# Building Code Map archaeology

This directory projects the repository's causal Deciduous graph into reviewable views. The canonical compressed graph is described in [`.deciduous/README.md`](../../.deciduous/README.md).

## Two-layer contract

The archaeology has two deliberately different layers:

1. The 98-node, 112-edge graph is a frozen reconstruction of the repository's causal baseline when the archaeology effort began.
2. These Markdown projections are maintained reconciliations against the current repository.

Historical nodes are not silently rewritten when implementation advances. CI validates the frozen graph and separately checks current profile and readiness inventories.

## Current state

Building Code Map currently provides:

- a React, TypeScript, Vite, and MapLibre public lookup and explorer;
- a Go HTTP API;
- deterministic local SQLite geocoding with provenance-bearing address-point and street-range results;
- semantically validated local boundary snapshots;
- point-only public regulatory resolution;
- deterministic ambiguity responses for overlapping peer state, county, or municipality observations;
- candidate state, county, municipality, special-area, tribal-area, and NERIS observations without treating containment as legal authority;
- six executable source-backed pilot profiles: Colorado, Florida, New Jersey, Virginia, Oregon, and North Carolina;
- scoped production-readiness manifests for Colorado, Florida, and New Jersey;
- explicit uncertainty, local-record requirements, source evidence, and applicability-date warnings;
- a pinned LORE documentation trust root.

The six profiles are not complete nationwide coverage. The three production-ready findings apply only to declared scopes and fixtures, not every locality, amendment, occupancy, or historical interval in those states.

## Reading order

1. Start with [decision narratives](../../.deciduous/narratives.md).
2. Read the [current architecture](current-architecture.md).
3. Follow the focused projections:
   - [address-to-result pipeline](address-to-result-pipeline.md)
   - [geocoder evolution](geocoder-evolution.md)
   - [jurisdiction and authority model](jurisdiction-authority-model.md)
   - [adoption and amendment lifecycle](adoption-amendment-lifecycle.md)
   - [dataset sources and provenance](dataset-sources-provenance.md)
   - [state coverage maturity](state-coverage-maturity.md)
   - [Building Code Map versus Building Code AST](building-code-map-vs-ast.md)
   - [website and product evolution](website-product-evolution.md)
4. Finish with the [risk register](risk-register.md), [adversarial review](adversarial-review.md), and [maintenance procedure](maintenance.md).

## Evidence rules

- A merged pull request is evidence of intent only after reconciliation with the resulting tree, schemas, tests, and later corrections.
- Geographic containment is not legal authority.
- Model-code publication is not legal adoption.
- Adoption, effective, mandatory, and repeal dates are distinct.
- A state profile does not imply complete local coverage.
- A production-ready manifest is bounded by its declared scope.
- Deterministic output does not guarantee legal correctness.
- Green tests do not establish nationwide completeness.
- Generated archaeology and LORE projections are not canonical regulatory records.
