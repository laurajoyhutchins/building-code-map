# Building Code Map

Building Code Map is an evidence-first engine and interface for determining the regulatory basis of a building project.

The public project focuses on the open mechanics around that problem: contracts, schemas, methodology, interoperability, and software that can operate against synthetic data or a hosted Building Code Map service.

## What this repository contains

- public data contracts and schemas;
- methodology for evidence-backed, fail-closed regulatory resolution;
- synthetic examples for developing and testing integrations;
- open-source engine and client components as they are admitted to the public boundary.

## What this repository does not contain

The production regulatory corpus is not distributed here. In particular, this repository is not a bulk dataset of jurisdiction-by-jurisdiction adoption research, authority assignments, local amendments, effective-date histories, verification annotations, or production evidence mappings.

That separation is intentional. Public software can remain inspectable without turning the maintained production knowledge base into a downloadable artifact.

## Public boundary

The archived BCM alpha repository is historical product state, not an upstream development branch. Material from alpha is review-required by default and must pass the [admission policy](docs/admission-policy.md) before reuse.

The [data visibility contract](docs/data-visibility.md) defines what belongs in the public repository, what a hosted project-scoped response may reveal, and what remains in the production/internal layer. Machine-checkable invariants live in `policy/public-boundary.json` and are enforced by `npm run validate:boundary`.

## Core semantics

A Building Code Map result should distinguish three outcomes:

- `verified`: available evidence is sufficient for the stated project scope and date;
- `conditional`: the answer depends on unresolved project facts or local evidence;
- `not_verified`: the available evidence is insufficient to make the requested claim.

Missing evidence is not treated as evidence that a requirement does not apply.

## Synthetic example

`demo/DEMO-XX/project-code-basis.json` is fictional. `DEMO-XX` is reserved for public examples and tests and does not represent a real jurisdiction.

## License

The software and documentation in this public distribution are licensed under Apache-2.0 unless a file states otherwise. Public laws, government records, model codes, standards, and other third-party materials retain their own terms.
