# Exact-Text Evidence

Building Code Map treats exact operative text as a verification invariant for legally material claims. The public implementation publishes the identities, relationships, and verification rules needed to prove a claim without requiring the source bytes themselves to be published in this repository.

## Identity model

The evidence model separates four concepts:

- `SourceDocument`: durable identity of a legal or publication object, such as an ordinance or administrative record;
- `SourceArtifact`: one immutable observed occurrence of that document, pinned by SHA-256 and byte size;
- `TextAnchor`: an exact portion of an artifact, identified by a legal locator and/or machine offsets plus a hash of the anchored text;
- `EvidenceLink`: a typed relationship between a claim and an exact anchor.

A URL is locator metadata. It is not artifact identity and cannot establish a legally material claim by itself.

## Evidence roles

Public evidence links support:

- `establishes`
- `corroborates`
- `contradicts`
- `defines`
- `incorporates`
- `supersedes`

`corroborates` is not a substitute for `establishes`. Contradicting evidence fails closed. An establishing link targeted by a `supersedes` relationship no longer qualifies as current establishing evidence.

## Verification invariant

`VerifyClaimEvidence` promotes a legally material claim to `verified` only when:

1. a non-superseded `establishes` link exists;
2. the link resolves to a valid exact `TextAnchor`;
3. the anchor belongs to an immutable `SourceArtifact` with a pinned SHA-256 and byte size;
4. the artifact belongs to the identified `SourceDocument`;
5. source policy authorizes that exact document identity for the claim kind, jurisdiction, and applicability date; and
6. no contradictory evidence is present.

Otherwise the verifier returns an explicit insufficient, conflicting, or superseded decision.

## Source policy

Source policy is keyed to document identity, claim kind, jurisdiction, and time. It does not authorize evidence merely because a URL, publisher, or informal source tier looks plausible.

The public schemas are [`evidence.schema.json`](../schemas/evidence.schema.json) and [`source-policy.schema.json`](../schemas/source-policy.schema.json).

## Publication boundary

This repository may publish project-authored evidence schemas, hashes, locators, policy rules, verification code, and synthetic fixtures. It does not need to publish restricted model-code or standards text, maintained production source artifacts, or externally controlled bytes to express or verify those identities.

`project-code-basis` 0.3 carries full `EvidenceLink` objects in `exact_evidence`, so exact document, artifact, and anchor identity can survive the project-level result without being collapsed into a URL or prose citation.
