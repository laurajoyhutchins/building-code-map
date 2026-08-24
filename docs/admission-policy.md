# Admission Policy

This repository is the public Building Code Map product boundary. The archived alpha repository is historical source material, not an upstream branch and not an automatically trusted source.

## Governing rule

Nothing enters this repository merely because it existed in Building Code Map alpha.

Every candidate from alpha or another external source must first be classified by material type and intended destination. The default disposition is `review_required`. If provenance, license terms, or redistribution rights are unclear, the material stays out of the public repository until the uncertainty is resolved.

The machine-readable invariants are in `policy/public-boundary.json`. Human review may make a decision more restrictive than that file, but not less restrictive without changing and reviewing the policy itself.

## Material classes

| Material | Public repository | Hosted response | Production/internal | Admission requirement |
| --- | --- | --- | --- | --- |
| First-party software and documentation | Allowed | Allowed | Allowed | Confirm authorship and intended Apache-2.0 release |
| Synthetic fixtures | Allowed | Allowed | Allowed | Use fictional identifiers and invented evidence |
| Public laws and government records | Reference-only by default | Citations and derived facts | Allowed | Check source-specific terms before redistributing bytes |
| Open third-party material | Conditional | Conditional | Conditional | Compatible license and preserved attribution |
| Model codes, standards, and other restricted works | Prohibited by default | Derived facts and citations only | Conditional | Lawful access; no unlicensed redistribution |
| Maintained production regulatory corpus | Prohibited | Project-scoped resolved records only | Allowed | No bulk export through the public product |
| Material from BCM alpha | Review required | Review required | Review required | Reclassify under a non-alpha class, review provenance/license, then clean-copy or reimplement |

`alpha_material` is a quarantine classification, not a permanent content class. Admission is complete only after the candidate has been classified by what it actually is.

## Alpha migration procedure

For each useful alpha artifact:

1. Identify what the artifact is: first-party code, third-party code, source material, derived regulatory fact, generated output, fixture, documentation, or something else.
2. Identify authorship, upstream source, notices, and license terms that affect the artifact.
3. Choose the destination independently of where alpha stored it.
4. For public first-party code or documentation, copy only the reviewed content into the new repository or reimplement it. Do not merge alpha Git history into this repository.
5. Preserve required third-party attribution. Do not copy restricted source text merely because alpha contained it.
6. Keep production research, evidence captures, jurisdiction-by-jurisdiction records, and maintained derived facts outside this public repository.
7. Demonstrate public behavior with synthetic fixtures whenever production data is unnecessary.

Generated output does not become public merely because it was generated. Its disposition follows the rights and visibility constraints of its inputs and the information it reveals.

## Fail-closed cases

Do not admit material when any of these remain unresolved:

- who authored it;
- whether third-party notices or license obligations apply;
- whether it contains model-code, standard, or other restricted text;
- whether it exposes maintained production research or bulk derived facts;
- whether a supposedly synthetic fixture represents a real jurisdiction or source;
- whether publication would reveal data that the hosted-service boundary intentionally withholds.

Uncertainty is a reason to keep material out, not a reason to assign a permissive default.

## Mechanical enforcement

`npm run validate:boundary` enforces a deliberately small set of invariants that software can know without legal judgment:

- alpha material can never be configured as automatically public;
- production corpus and restricted standards remain prohibited from the public repository;
- JSON fixtures under `demo/` must use the reserved synthetic jurisdiction prefix;
- top-level `corpus/`, `production-data/`, and `research/` directories are rejected.

The validator is a guardrail, not a license analyzer. Source-specific rights still require review.
