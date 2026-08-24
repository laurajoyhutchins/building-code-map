# BCM Alpha Disposition

This document records the product boundary created when `building-code-map-alpha` was archived on August 24, 2026.

## Current authority

`building-code-map-alpha` is frozen historical product state. It is not an upstream branch, an active production source repository, a fallback implementation, or a place for new Building Code Map development.

This repository, `building-code-map`, is the clean-history public product repository. The maintained production regulatory knowledge base belongs behind a separate private data/service boundary and must not be recreated here as a bulk corpus.

Alpha remains useful as evidence of prior product work. Reuse is selective and governed by the admission policy, not by repository ancestry.

## Disposition by material class

### Public software candidates

First-party software and contracts may be reviewed for clean reimplementation or deliberate copying under the admission policy. The strongest candidates from alpha are:

- the transport-independent Authority Engine boundary;
- typed fail-closed, ambiguity, and unsupported outcomes;
- project-level code-set verification contracts and structured unresolved requirements;
- deterministic bundle/provenance identity and capability readiness semantics;
- deterministic geocoder ranking, interpolation, and provenance mechanics when local geocoding remains in product scope;
- source-safe robustness and executable capability-test concepts.

These are candidates, not inherited implementation. Alpha Git history is not merged into this repository.

### Private production knowledge candidates

Maintained jurisdiction research, governmental-entity inventories, regulatory evidence metadata, exact-text legal evidence, source-capability policy, readiness records, and other real-jurisdiction knowledge do not belong in this public repository merely because alpha stored them in Git.

Useful material in this class must be reviewed for provenance, source terms, currentness, and visibility before being admitted to the private production knowledge base. Alpha remains the historical source until that review occurs; it is not the ongoing authority afterward.

### Regenerate, do not migrate

Deterministic projections should be rebuilt from reviewed authoritative facts instead of copied from alpha. This includes normalized regulatory profiles and rules, jurisdiction regimes, coverage projections, status summaries, generated documentation, assembled executables, and hydrated runtime snapshots.

Generated output does not become authoritative because it was committed in alpha.

### Historical operating machinery

Alpha-specific publication, deployment, repository-maintenance, and recovery choreography is historical by default. In particular, the old public-release exporter completed its purpose once the clean public repository was created. Alpha-specific Pages deployment, LORE/Deciduous state, Actions-retention tuning, object-hydration choreography, and materialization workflow plumbing should not be ported automatically.

If the successor later needs an equivalent capability, implement the smallest mechanism required by the successor architecture rather than preserving alpha compatibility.

### Third-party and restricted material

Model-code and standards text, third-party data, address datasets, department-owned data, retained source captures, and other externally sourced bytes retain their own terms. Accessibility in alpha is not evidence of redistribution permission. Public release remains fail-closed unless the applicable rights are established.

## Open alpha work at archive time

An open alpha pull request is historical candidate work, not accepted successor state.

Several open pull requests contain real-jurisdiction research. Those branches may contain useful research for the future private corpus, but duplicate and superseded candidates must be reconciled before any fact is accepted.

PR #253, the project-code-verification branch, is a particularly useful public-software design source. Its multi-family verification model, structured requirements, deterministic verification identity, and thin transport boundary should be reviewed as a clean successor implementation rather than merged wholesale.

PR #263, the automatic object-transport branch, represents alpha-specific storage/recovery machinery and is not a default successor dependency. PR #265 is an obsolete Actions-retention adjustment. Documentation-only alpha branches may supply historical context, but the successor documentation is authoritative for current product boundaries.

## Re-energization boundary

Substantive successor development should proceed only through post-alpha paths:

1. public software and schemas live in this repository;
2. maintained real-jurisdiction knowledge has a separate private authoritative home;
3. generated projections are derived deterministically from authoritative facts;
4. public interfaces consume synthetic fixtures or a hosted service boundary rather than importing the production corpus;
5. automation and orchestration target the successor architecture and treat alpha as retired.

The goal is not to reproduce alpha in pieces. The goal is to preserve the useful product semantics while removing accidental coupling between public software, proprietary or redistribution-sensitive knowledge, generated state, and historical operating machinery.
