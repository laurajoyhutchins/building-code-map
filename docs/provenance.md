# W3C PROV-O source lineage

BCM publishes provenance as **additional, source-safe evidence metadata**, not as a replacement for its one canonical legal-evidence verifier. W3C [PROV-O](https://www.w3.org/TR/prov-o/) describes who or what contributed to data and how it was derived; it does **not** certify that a statute is operative or that a project is code-compliant.

## Ownership and contracts

- `provenance/graph.go` and `provenance/parse.go`: no-dependency deterministic JSON-LD 1.1 serialization, strict external import and closed-graph validation with genuine PROV-O IRIs and type-correct `prov:wasDerivedFrom`, `prov:used`, `prov:wasGeneratedBy`, `prov:wasAssociatedWith` and `prov:wasAttributedTo` relations. Input order does not change bytes. No remote context retrieval occurs.
- `engine/prov.go`: read-only projection of one existing `EvidenceClaim`. Document, SHA-pinned artifact, anchored excerpt, source link and claim remain separate entities. The link's source-evidence role is preserved. `bcm:hasEvidenceLink` is **not** equivalent to `prov:wasDerivedFrom` or an approved legal claim.
- `schemas/prov-graph.schema.json` + `schemas/prov-graph.odcs.json`: the paired public cross-repository data contract. Both are required where another repository consumes the export; the JSON Schema is structural, and the Go validator also checks unique IDs, closed references and predicate domain/range.
- `engine.VerifyClaimEvidence`, runtime-owned `SourcePolicy`, `engine.ProjectVerifier` and the existing project-code-basis contract remain sole owners of evidence admission and project verdicts.

## Export a claim

In Go:

```go
g, err := engine.ProjectClaimPROV(claim)
if err != nil { return err }
payload, err := g.JSONLD()
if err != nil { return err }
```

The call refuses missing or malformed source identities and documents, conflicting reused IDs, invalid SHA-256 and unknown evidence roles. The claim entity's URI binds its original ID, kind, jurisdiction, applicability date and legal-materiality flag; these properties are also represented explicitly in the graph. Source-evidence `supersedes` edges retain their specific referenced link via `bcm:targetsEvidenceLink` and refuse missing or self-targets. It does not use `Verified` status, add review agents, or generate an activity unless that activity exists independently in trusted source records. All source bytes remain in their appropriate custody domain.

## Command-line export

The public, read-only `bcm-prov` CLI accepts exactly one JSON `EvidenceClaim` object on standard input and emits source-safe PROV-O JSON-LD on standard output:

```sh
go run ./cmd/bcm-prov < test/fixtures/prov/claim.json > /tmp/bcm-provenance.jsonld
```

To **admit externally authored** PROV-O JSON-LD, run the same repository-owned whole-graph validator before any other processing:

```sh
go run ./cmd/bcm-prov validate < test/fixtures/prov/actor.jsonld > /tmp/validated.jsonld
```

The validator rejects unknown fields, namespace substitution, malformed literals, duplicated node IDs, orphan target references, invalid predicate domain/range, incomplete or altered claim scopes, and invalid supersession references. Successful output is deterministically serialized JSON-LD. **JSON Schema and ODCS checks alone are insufficient:** they cannot close cross-node references or reject duplicate @id values. Cross-repository consumers must pass the graph through `bcm-prov validate` (or call `provenance.ParseJSONLD` in Go) after schema checks and before trusting or storing an imported graph. There is no second source registry or project-verdict path.

The example uses invented `DEMO-XX` records. The CLI rejects unknown input fields, more than 1 MiB, trailing JSON, malformed source links, orphaned document/anchor identities and digest mismatches. The output is portable metadata, not an approved claim. Cross-repository readers must validate the paired JSON Schema and ODCS contract and retain the canonical source-policy decision outside the graph.

## Add independently observed activities and agents

Where a source-operations adapter can actually prove an execution event, it may add `provenance.Node{Kind: provenance.Activity}` and `provenance.Node{Kind: provenance.Agent}` to the graph, plus typed `provenance.Edge` relationships. For example, a recorded extraction activity can `prov:used` the exact source artifact; a retained result may `prov:wasGeneratedBy` that activity. Associate the actor through `prov:wasAssociatedWith` **only if the actor's identity is established**. Re-run `Graph.Validate()` or `Graph.JSONLD()` after extending the graph.

Do not invent activities, download times, reviewers, or source digest evidence to fill gaps. If the original publisher, activity, or actor is unknown, leave it unasserted. An event log saying a process ran does not establish independent verification of its output.

## PROV-O semantics and data preservation

The `@context` fixes `prov:` to `http://www.w3.org/ns/prov#`, `xsd:` to XML Schema types, and `bcm:` to the project namespace. Reference targets are explicit JSON-LD `{"@id": "..."}` values; SHA-256 hashes and existing source IDs are literals, never inferred source URLs. Stable node URNs are deterministically computed from the **typed** original source IDs using SHA-256. The original IDs remain on `bcm:sourceID`. The graph can be joined to underlying private receipts without including protected normative text.

Do not use `prov:wasAttributedTo` or `prov:wasAssociatedWith` as vague 'the source must have had an author' placeholders. Do not treat a source URL as an immutable artifact identity. Do not use `prov:wasDerivedFrom` to imply that a contradictory or corroborating source established a claim.

## Verification boundary

The Go tests exercise source-pointer closure, deterministic serialization, scoped claim identities, supersession targets, reused-ID conflicts, digests, invalid timestamps and invalid PROV predicate domain/range. External import is negative-tested against graphs that satisfy the JSON Schema but violate whole-graph invariants. The Node contract tests validate the same JSON-LD representation against the published JSON Schema and the paired ODCS contract using the repository's pinned validators. The test data is fictional.

An independent pinned RDFLib parser in CI additionally expands the actual Go output and an independently authored synthetic Activity/Agent graph to W3C PROV RDF triples. These remain **structural PROV-O/JSON-LD checks**, not a full RDF reasoner, remote-source authenticity check, or legally reviewed source-to-verdict proof. A production adapter still needs immutable source acquisition/review receipts and independently validated provenance of its asserted activities. Promote a real claim only through the existing runtime-owned source policy, never because its JSON-LD passes a schema.
