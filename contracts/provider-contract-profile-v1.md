# BCM Provider Contract Profile v1

Status: Normative draft  
Profile version: `1.0.0`  
ODCS base: `v3.1.0`

## Purpose

This profile defines the public logical contract between Building Code Map (BCM) and a regulatory-data provider. It adopts compatible concepts from the Open Data Contract Standard (ODCS) v3.1.0 while preserving BCM's existing regulatory, evidence, uncertainty, and publication semantics.

This profile is intentionally narrower than a general data-platform contract. It describes what a BCM provider must mean and identify. It does not prescribe database tables, object-store paths, compiler internals, private provider mappings, deployment topology, or the contents of a maintained regulatory corpus.

The key words **MUST**, **MUST NOT**, **SHOULD**, **SHOULD NOT**, and **MAY** are normative requirements of this BCM profile.

## Authority and compatibility

A BCM provider contract conforming to profile `1.x`:

1. MUST use ODCS `apiVersion: v3.1.0` and `kind: DataContract` for its contract envelope;
2. MUST satisfy the stricter BCM requirements in this document, including fields that ODCS itself makes optional;
3. MUST preserve the existing BCM project-verdict, unresolved-requirement, exact-evidence, and source-policy semantics;
4. MUST use BCM-defined extension semantics for regulatory concepts that ODCS does not define; and
5. MUST NOT interpret ODCS metadata as authority to expose material that BCM's publication and data-visibility policies keep non-public.

Conformance to this profile means **ODCS-aligned BCM provider contract**, not that BCM's regulatory graph is itself an ODCS data-schema model. ODCS supplies the contract envelope, stable-ID discipline, authoritative-definition mechanism, and extension mechanism. BCM remains authoritative for regulatory meaning.

Normative BCM definitions referenced by this profile are:

- [Methodology](../docs/methodology.md)
- [Data Visibility Contract](../docs/data-visibility.md)
- [Publication Policy](../docs/publication-policy.md)
- [Exact-Text Evidence schema](../schemas/evidence.schema.json)
- [Project Code Basis schema](../schemas/project-code-basis.schema.json)
- [Source Policy schema](../schemas/source-policy.schema.json)

If this profile and an existing BCM domain schema appear to conflict about evidence identity, verdicts, unresolved requirements, or source publication policy, the BCM domain schema or policy is authoritative for that domain concept until the conflict is resolved by a versioned profile change.

## Contract envelope

BCM tightens the ODCS v3.1.0 fundamentals for provider use.

| Field or section | ODCS v3.1.0 | BCM profile v1 | BCM rule |
| --- | --- | --- | --- |
| `apiVersion` | Required | Required | MUST equal `v3.1.0`. |
| `kind` | Required | Required | MUST equal `DataContract`. |
| `id` | Required | Required | Stable contract identity. MUST NOT depend on a filename or storage path. |
| `version` | Required | Required | Provider-contract version. |
| `status` | Required | Required | Lifecycle status. |
| `name` | Optional | Required | Human-readable only. MUST NOT be reference identity. |
| `description.purpose` | Optional | Required | States what the provider contract is intended to support. |
| `description.usage` | Optional | Required | States intended use, including supported runtime use where applicable. |
| `description.limitations` | Optional | Required | States coverage, uncertainty, visibility, or other material limitations. |
| `authoritativeDefinitions` | Optional | Required | MUST link to the applicable BCM methodology, schemas, and publication/source policy definitions. |
| `description.authoritativeDefinitions` | Optional | Optional | MAY add context-specific definitions. |
| `customProperties` | Optional | Conditional | MUST be used for BCM-specific envelope extensions defined by this profile or a successor. |
| `domain`, `dataProduct`, `tenant`, `tags` | Optional | Optional / pass-through | MAY be supplied. BCM MUST NOT use them as regulatory authority or stable regulatory identity. |
| ODCS data schema | Optional | Not used by BCM v1 | MAY be present for other consumers. It does not define BCM regulatory-object semantics. |
| team, role, support metadata | Optional | Not used by BCM v1 | MAY be present. It does not affect BCM resolution. |
| server / infrastructure metadata | Optional | Not used by BCM v1 | MUST NOT be required for BCM conformance. Private connection details MUST NOT be exposed. |
| pricing metadata | Optional | Not used by BCM v1 | MUST NOT be required for BCM conformance. |
| SLA metadata | Optional | Not used by BCM v1 | MAY be present but does not alter regulatory semantics. |
| data-quality metadata | Optional | Not used by BCM v1 | MAY be present but does not replace BCM evidence validation. |

A provider MUST NOT claim BCM profile conformance merely because the document validates as an ODCS v3.1.0 contract. The BCM requirements above and the domain rules below also apply.

## Stable regulatory identity

Every object that can be the target of a cross-object reference MUST have a stable, non-empty string `id`.

At minimum this requirement applies to:

- jurisdiction;
- authority;
- regulatory instrument;
- adoption;
- amendment;
- code family;
- code edition;
- source document;
- source artifact;
- claim;
- evidence link;
- source-policy rule; and
- provider bundle.

### ID rules

1. An ID MUST identify a logical object, not its current presentation.
2. Human-readable names, labels, titles, array positions, row numbers, filenames, and display ordering MUST NOT be identity.
3. Moving an object between files, tables, stores, or bundle layouts MUST NOT by itself change the object's ID.
4. IDs MUST NOT contain secrets or encode private storage locations.
5. An ID MUST be unique within the scope in which its provider contract promises uniqueness.
6. Cross-bundle references MUST include enough stable context to identify both the bundle and target object, such as a stable URI or the pair `(provider_bundle_id, object_id)`.
7. A provider MAY use UUIDs, URIs, or another deterministic opaque identifier scheme if it obeys these rules.

ODCS v3.1.0's stable-ID and fully qualified reference concepts inform this rule. BCM does not adopt ODCS's foreign-key reference notation as the complete regulatory reference language. BCM regulatory references have domain meaning beyond relational foreign keys.

## Exact evidence remains exact

Stable contract references do not weaken BCM's exact-evidence model.

The public `evidence.schema.json` already distinguishes:

- a `source_document` identity;
- a content-addressed `source_artifact` with SHA-256 and byte size;
- a `text_anchor` bound to an artifact and text hash; and
- an `evidence_link` connecting a claim to that exact evidence.

A provider contract MUST preserve those identities when supplying exact evidence. It MUST NOT substitute a document title, mutable URL, regulatory-object name, or ODCS relationship for the artifact and anchor identity required by the BCM evidence schema.

A logical regulatory object MAY refer to an evidence claim or evidence link by stable ID. The referenced evidence still has to satisfy the BCM evidence contract independently.

## Logical regulatory model

The provider contract describes a logical model. Physical implementation is explicitly non-normative.

The canonical payloads for BCM regulatory authority objects are defined separately from this profile. This profile establishes the identity and reference rules those payloads must obey so that the provider boundary can evolve without coupling the public contract to an alpha database or private production layout.

The logical model MUST preserve these distinctions:

- a **jurisdiction** is a governed place or scope, not automatically the regulatory authority for every subject;
- an **authority** is an entity or authority scope that can issue, adopt, amend, administer, or enforce regulatory instruments;
- a **regulatory instrument** is distinct from the evidence used to prove its content or applicability;
- an **adoption** and an **amendment** are distinct regulatory acts and MUST remain separately identifiable when both exist;
- a **code family** and a **code edition** are separately identifiable concepts;
- a **claim** is a proposition evaluated by BCM, not the source material itself;
- an **evidence link** states how exact evidence relates to a claim;
- a **source-policy rule** governs source treatment and publication independently from whether a source is technically retrievable; and
- a **provider bundle** identifies a coherent supplied set and is suitable for provenance in a project result.

A database primary key MAY implement a stable ID. A database primary key is not required. No SQL table, object key, local path, cache key, compiler symbol, or private provider adapter becomes part of the public logical contract merely because an implementation uses it.

## BCM extensions and namespacing

BCM-specific semantics that do not map directly to an ODCS v3.1.0 field MUST be treated as BCM extensions.

For ODCS `customProperties` encoding:

- BCM-owned property names MUST use the camel-case prefix `bcm`, for example `bcmProfileVersion` or `bcmProviderBundleId`;
- the prefix `bcm` is reserved for properties defined by a versioned BCM contract or schema;
- providers MUST NOT invent new `bcm*` meanings locally;
- provider-specific extensions MUST use a distinct, stable provider prefix and MUST NOT collide with ODCS or BCM names; and
- an extension that becomes required for BCM interoperability MUST be promoted into a versioned BCM profile or schema before BCM relies on it.

Profile v1 requires a root custom property named `bcmProfileVersion` with value `1.0.0`. A provider that supplies a concrete provider bundle MUST also expose the stable bundle identity through the BCM domain model; it MAY mirror that identity in `bcmProviderBundleId` for contract-level discovery.

The machine-enforceable encoding of BCM regulatory objects is a separate schema concern. This profile deliberately does not smuggle a database-shaped object model into `customProperties`.

## Authoritative definitions

Root `authoritativeDefinitions` MUST provide links sufficient for a consumer to discover the BCM definitions governing the provider contract.

At minimum, a production BCM provider contract MUST identify authoritative definitions for:

1. BCM methodology;
2. the applicable BCM provider/profile contract;
3. project-result and exact-evidence schemas used by the runtime;
4. source/publication policy; and
5. any additional versioned BCM domain schema required to interpret the provider bundle.

Definitions SHOULD use stable public URLs. A definition link is a pointer to normative meaning, not evidence that a real-world regulatory proposition is true.

Authoritative definitions MUST NOT point to a private deployment console, private object-store location, private research notebook, secret-bearing URL, or restricted source byte location as a requirement for understanding the public contract.

## Limitations are first-class

`description.limitations` is required by BCM even though ODCS v3.1.0 makes it optional.

A provider contract MUST state material limits that affect how a consumer should interpret the provider. Examples include:

- geographic or subject-matter coverage;
- supported as-of dates or known temporal bounds;
- categories that require local records or human verification;
- incomplete or unavailable evidence classes;
- non-public operational dependencies at a product-safe level; and
- restrictions on redistribution or bulk export.

A limitation MUST NOT be hidden in implementation notes when it can affect whether a BCM answer is verified, conditional, or not verified.

## Public and private boundary

Provider-contract conformance does not grant publication rights.

A public BCM provider contract MUST NOT require or expose:

- the maintained jurisdiction-by-jurisdiction production corpus;
- private project data;
- private research notes or verification annotations;
- restricted model-code, standard, or publisher source bytes;
- bulk evidence mappings that BCM intentionally keeps non-public;
- private provider mappings or adapters;
- compiler internals or normalization heuristics that are outside the public product contract;
- database credentials, private server details, object-store paths, or deployment topology; or
- source material whose authorship, licensing, attribution, or publication status remains unresolved.

Public contracts SHOULD prefer stable identifiers, citations, derived regulatory facts, provenance, and synthetic examples over redistributing source bytes.

A provider MAY use non-public operational data behind the provider boundary. The fact that BCM can use material operationally does not make that material part of the public provider contract.

## Relationship to project results

The provider contract supplies inputs to BCM's existing resolution and verification path. It does not create a second verdict system.

The existing `project-code-basis.schema.json` remains authoritative for public project-result concepts including:

- `verified`, `conditional`, and `not_verified` verdicts;
- per-code resolution states such as `resolved`, `partially_resolved`, `local_record_required`, `insufficient_evidence`, `ambiguous`, and `conflicting`;
- unresolved project facts, local records, and evidence defects; and
- provenance including a provider `bundle_id`.

Providers MUST preserve uncertainty rather than converting missing or conflicting evidence into a confident adoption conclusion.

## Versioning and backward compatibility

The BCM provider profile uses semantic versioning.

- **Major**: changes that can invalidate a previously conforming provider, change identity/reference semantics, reinterpret an existing field, or change the meaning of a required domain concept.
- **Minor**: backward-compatible additions, including new optional metadata or extension semantics that old consumers can safely ignore.
- **Patch**: clarifications and corrections that do not change conformance semantics.

A provider contract's own ODCS `version` is the version of that provider contract. `bcmProfileVersion` identifies the BCM profile against which it claims conformance. They are distinct version coordinates.

Profile `1.x` is pinned to ODCS `v3.1.0`. Adopting a different ODCS major or minor version requires an explicit BCM profile revision and compatibility review rather than silently changing `apiVersion`.

A BCM runtime MUST fail closed when it cannot safely interpret the claimed profile version. It MUST NOT silently downgrade, reinterpret an unsupported major version, or treat ordinary ODCS validation as proof of BCM provider compatibility.

## Synthetic example

This fragment is illustrative only. It contains no real jurisdiction record or restricted source material.

```yaml
apiVersion: v3.1.0
kind: DataContract
id: bcm-provider-demo-xx
name: BCM synthetic provider
version: 0.1.0
status: active

description:
  purpose: Exercise the public BCM provider contract with synthetic regulatory objects.
  usage: Development and conformance testing only.
  limitations: DEMO-XX is fictional and provides no real regulatory coverage.

authoritativeDefinitions:
  - id: bcm_methodology
    type: businessDefinition
    url: https://github.com/laurajoyhutchins/building-code-map/blob/main/docs/methodology.md
    description: BCM evidence and resolution methodology.
  - id: bcm_profile
    type: canonicalUrl
    url: https://github.com/laurajoyhutchins/building-code-map/blob/main/contracts/provider-contract-profile-v1.md
    description: BCM provider contract profile.
  - id: bcm_publication_policy
    type: businessDefinition
    url: https://github.com/laurajoyhutchins/building-code-map/blob/main/docs/publication-policy.md
    description: BCM publication boundary.

customProperties:
  - id: bcm_profile_version
    property: bcmProfileVersion
    value: 1.0.0
  - id: bcm_provider_bundle_id
    property: bcmProviderBundleId
    value: bundle:DEMO-XX:synthetic-v1
```

The synthetic example does not define the full BCM regulatory-object payload. Canonical object schemas and mechanical conformance rules are separate versioned deliverables built on this profile.

## Conformance summary

A provider conforms to BCM Provider Contract Profile v1 only when all of the following are true:

- its ODCS envelope uses `v3.1.0` and satisfies the BCM-required fundamentals;
- cross-object targets have stable IDs and names or array positions are never used as identity;
- BCM-specific semantics follow the extension namespace rules;
- logical semantics are independent of physical storage and private implementation;
- authoritative definitions identify the governing BCM public contracts;
- limitations are explicit;
- BCM evidence, project-verdict, unresolved-requirement, and source-policy semantics remain authoritative; and
- the contract does not expose or require private production implementation, maintained corpus contents, private project data, or restricted source material.
