# Publication Policy

Building Code Map is a public software project with a deliberately bounded data surface. Publication decisions are based on what a material is, who owns it, what its terms allow, and what information it reveals.

## Governing rule

Material is not public merely because it is accessible, useful, or already present in another system.

Anything entering this repository must be classified by material type and reviewed for authorship, licensing, attribution, provenance, and intended visibility. Externally sourced material is `review_required` by default. If those questions are unresolved, the material stays out of the public repository until they are resolved.

Machine-checkable invariants live in [`policy/public-boundary.json`](../policy/public-boundary.json). Human review may make a publication decision more restrictive than that policy, but not less restrictive without changing and reviewing the policy itself.

## Material classes

| Material | Public repository | Hosted response | Non-public operations | Requirement |
| --- | --- | --- | --- | --- |
| Project-authored software and documentation | Allowed | Allowed | Allowed | Confirm authorship and intended Apache-2.0 release |
| Synthetic fixtures | Allowed | Allowed | Allowed | Use fictional identifiers and invented evidence |
| Public laws and government records | Reference-only by default | Citations and derived facts | Allowed | Check source-specific terms before redistributing bytes |
| Open third-party material | Conditional | Conditional | Conditional | Compatible license and preserved attribution |
| Model codes, standards, and other restricted works | Prohibited by default | Derived facts and citations only | Conditional | Lawful access and no unlicensed redistribution |
| Maintained production regulatory corpus | Prohibited | Project-scoped resolved records only | Allowed | No bulk export through the public product |
| Other externally sourced material | Review required | Review required | Review required | Classify, review provenance and rights, then choose a destination |

Generated output follows the rights and visibility constraints of its inputs and the information it reveals. Generation does not create publication rights.

## Fail-closed cases

Do not publish material when any of these remain unresolved:

- who authored or owns it;
- whether third-party notices or license obligations apply;
- whether it contains model-code, standard, or other restricted text;
- whether it exposes maintained production research or bulk derived facts;
- whether a supposedly synthetic fixture represents a real jurisdiction or source;
- whether publication would reveal information intentionally outside the public product contract.

Uncertainty is a reason to withhold material pending review, not a reason to assume permissive terms.

## Mechanical enforcement

`npm run validate:boundary` enforces a small set of invariants that software can know without making legal judgments:

- externally sourced material cannot be configured as automatically public by default;
- the maintained production corpus and restricted standards remain prohibited from the public repository;
- JSON fixtures under `demo/` must use the reserved synthetic jurisdiction prefix;
- top-level `corpus/`, `production-data/`, and `research/` directories are rejected.

The validator is a guardrail, not a license analyzer. Source-specific rights still require review.
