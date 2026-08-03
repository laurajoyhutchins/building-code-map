# Building Code Map archaeology

This directory projects the repository's causal Deciduous graph into reviewable views. The canonical
graph is the schema-native compressed seed described in [`.deciduous/README.md`](../../.deciduous/README.md).
The upstream-compatible `docs/graph-data.json` export is generated on demand and remains local.

## Scope

The backfill contains **98 nodes** and **112 typed causal edges**. It covers the original
applicability problem, the address-to-result pipeline, local-first execution, rejection of RAG as
final authority, normalization and geocoding, polygon routing, the authority model, adoption and
amendments, data-source roles, state expansion, deterministic outputs, repository boundaries, LORE
adoption, website evolution, compatibility debt, public-repository framing, and current risks.

## Current state in one paragraph

`main` at archaeology start was `761ddc49bf6dffa50dc0454a5632f04fa3959594`. It provided a React,
TypeScript, Vite, and MapLibre public lookup and explorer; a Go API; local SQLite geocoding schema 1;
local SQLite boundary snapshots with legacy DuckDB compatibility; point-in-polygon context for state,
county, municipality, special-area, tribal-area, and NERIS fire-jurisdiction layers; executable
source-backed policy profiles for Colorado, Florida, and New Jersey; resolution-result schema 1.0;
and a pinned LORE trust root. The three state profiles were pilots with partial verification, not
production-ready national coverage. Hydrated boundary and address databases were not committed.

## How to read the graph

1. Start with [decision narratives](../../.deciduous/narratives.md).
2. Read the [current architecture](current-architecture.md).
3. Follow a focused projection:
   - [address-to-result pipeline](address-to-result-pipeline.md)
   - [geocoder evolution](geocoder-evolution.md)
   - [jurisdiction and authority model](jurisdiction-authority-model.md)
   - [adoption and amendment lifecycle](adoption-amendment-lifecycle.md)
   - [dataset sources and provenance](dataset-sources-provenance.md)
   - [state coverage maturity](state-coverage-maturity.md)
   - [Building Code Map versus Building Code AST](building-code-map-vs-ast.md)
   - [website and product design evolution](website-product-evolution.md)
4. Finish with the [risk register](risk-register.md) and [maintenance procedure](maintenance.md).

## Evidence rule

A merged pull request or README claim is evidence of intent only until reconciled with the resulting
tree, schemas, tests, and later corrections. Open draft PRs are represented as proposed or historical,
not current implementation. Related repositories are referenced only to establish ownership
boundaries.

## Key qualifications

- Geographic containment is not legal authority.
- A model-code publication is not a legal adoption.
- Adoption date and effective date are not interchangeable.
- A state record does not imply complete local coverage.
- Deterministic output does not guarantee legal correctness.
- Green tests do not establish nationwide completeness.
- Public website availability does not establish production-grade data.
- Generated reports and projections are not canonical regulatory records.
