# DEMO-XX provider example

DEMO-XX is a fully synthetic public example. It has **no real regulatory coverage** and must not be used to answer real building-code questions.

This directory shows three deliberately separate layers:

1. `provider-contract.json` is the ODCS v3.1.0 contract envelope. It identifies the provider agreement, BCM profile version, bundle identity, purpose, usage, limitations, and public authoritative definitions.
2. `provider-bundle.json` is an **illustrative** logical BCM provider bundle. Its object shape is **not normative**. It demonstrates stable regulatory identities and references without making database tables, storage paths, compiler internals, or private provider mappings part of the public contract. Machine-readable fields for coverage, effective dates, claim evidence, and source policy exist so the public conformance rules can inspect the logical relationships; they do not turn this example into a physical storage schema.
3. `project-code-basis.json` is a runtime/project result. It records BCM's verdict, unresolved questions, exact evidence, and runtime provenance. It is not a provider contract and it is not the provider bundle.

Executable regulatory-integrity checks are documented in [`contracts/provider-conformance-v1.md`](../../contracts/provider-conformance-v1.md).

## Identity continuity

The example intentionally reuses synthetic identities already present in `project-code-basis.json`:

- provider bundle: `DEMO-BUNDLE`;
- jurisdiction: `DEMO-XX`;
- claim: `DEMO-CLAIM-001`;
- evidence link: `DEMO-LINK-001`;
- source document: `DEMO-DOC-001`;
- source artifact: `DEMO-ARTIFACT-001`;
- text anchor: `DEMO-ANCHOR-001`.

The illustrative bundle adds stable synthetic identities for an authority, regulatory instrument, adoption, amendment, code families, code editions, and source-policy rule. Human-readable names and array positions are never used as identity.

## Quality vocabulary

`provider-bundle.json` includes small illustrative expectations for conformity, coverage, and completeness. These are examples of provider-quality vocabulary, not project verdicts. Passing provider conformance does not mean a project claim is `verified`. Evidence and project verdicts continue through the existing BCM engine path.

## Public/private boundary

Everything here is fictional. URLs use either this public repository or `example.invalid`. The example contains no real jurisdiction research, maintained production corpus, restricted model-code or standards text, private project data, private provider mappings, credentials, storage locations, compiler heuristics, or deployment topology.