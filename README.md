# Building Code Map

**Evidence-first regulatory intelligence for building projects.**

Building Code Map is being built to answer a deceptively difficult question: given a project location, an applicability date, and the facts of the project, which regulatory authorities and code instruments can actually be verified, what evidence supports that answer, and what still requires confirmation?

This repository is the public foundation of that product. It contains the contracts, methodology, examples, and open-source components needed to make Building Code Map inspectable and interoperable.

## Core semantics

Building Code Map distinguishes three top-level outcomes:

- `verified`: available evidence is sufficient for the stated project scope and date;
- `conditional`: the answer depends on unresolved project facts or local evidence;
- `not_verified`: the available evidence is insufficient to make the requested claim.

Missing evidence is never treated as evidence that a requirement does not apply.

## Trust model

The project is built around a few durable rules:

- **Evidence before claims.** Regulatory conclusions should be traceable to evidence.
- **Time matters.** Adoption, transition, supersession, and effective dates are part of the answer.
- **Project facts matter.** Occupancy, scope, permit timing, location, and other facts can change applicability.
- **Uncertainty stays visible.** Ambiguous or unsupported conditions fail closed instead of being silently guessed away.
- **Public interfaces are deliberate.** A useful project-scoped answer does not require publishing every underlying source artifact or the complete maintained regulatory corpus.

See [Methodology](docs/methodology.md) for more detail.

## What is here

- public schemas and interface contracts;
- methodology for evidence-backed regulatory resolution;
- synthetic fixtures for examples and tests;
- open-source engine and client components as they are released;
- publication and data-distribution guardrails.

The current public contract starts with [`project-code-basis.schema.json`](schemas/project-code-basis.schema.json) and the fictional [`DEMO-XX`](demo/DEMO-XX/project-code-basis.json) example.

## Public runtime

The [`engine`](engine/) package is the first public-safe Authority Engine slice. It preserves the project-authored query and error contracts while putting regulatory resolution behind an injected `Provider` interface. That boundary is deliberate: the open-source runtime can execute and validate synthetic resolutions without requiring the maintained regulatory corpus, production geocoder or boundary databases, or restricted source material.

Per-code results keep `unsupported`, `not_applicable`, and `insufficient_evidence` distinct. `not_applicable` requires an affirmative basis; missing evidence cannot silently become non-applicability. Geographic results use country and primary-jurisdiction identity, with state-equivalent metadata optional rather than structural.

## Project Code Verification

[`ProjectVerifier.VerifyProject`](engine/project.go) is the primary project-level public workflow. Callers provide a project ID, location, applicability date, and any known project facts; they do not have to choose one code family before asking for an answer.

The verifier uses the canonical runtime result and deterministically reduces it to a project verdict. It preserves user-supplied and BCM-derived facts, per-code states, unresolved requirements, evidence identifiers, and runtime provenance. Only structured `project_fact` requirements become follow-up questions. Local-record requirements and evidence defects remain visibly different failure modes.

See [Project Code Verification](docs/project-verification.md) for the public reduction rules and contract semantics.

## Public HTTP API

The [`httpapi`](httpapi/) package exposes the canonical project verifier at `POST /v1/project-code-basis`. It is a transport adapter only: successful responses are the existing project-code-basis representation, and typed engine errors retain their stable codes across the HTTP boundary.

The adapter uses strict bounded JSON decoding, rejects unknown fields and trailing values, and is dependency-injected so public tests remain synthetic. It does not add production regulatory data, geocoder or boundary stores, source bytes, secrets, or a second decision path.

See [Project Verification HTTP API](docs/http-api.md) for the wire contract and error mapping.

## Data and licensing

This repository is not a bulk download of maintained jurisdiction-by-jurisdiction regulatory research. A hosted Building Code Map service may return project-scoped conclusions, citations, provenance, and unresolved requirements without exposing the complete maintained corpus or restricted source material.

Third-party laws, government records, model codes, standards, datasets, and other source materials retain their own terms. See the [Publication Policy](docs/publication-policy.md), [Data Visibility Contract](docs/data-visibility.md), and [`NOTICE`](NOTICE).

## Validate the public foundation

Requires Node.js 22 or later and Go 1.23 or later.

```sh
go test ./...
npm test
npm run validate
```

The validation gate checks the public distribution policy and keeps public fixtures synthetic.

## Project status

Building Code Map is in active early development. Public runtime, engine, and client capabilities will be added incrementally as their contracts and distribution boundaries are ready to support them.

## Contributing

Contributions are welcome. Start with [`CONTRIBUTING.md`](CONTRIBUTING.md). Security issues should be reported according to [`SECURITY.md`](SECURITY.md).

## License

Project-authored software and documentation in this repository are licensed under Apache-2.0 unless a file states otherwise. Third-party materials retain their own terms.
