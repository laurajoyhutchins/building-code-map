# BCM Provider Conformance Rules v1

Status: Normative draft  
Provider profile: `1.0.0`

## Purpose

This contract defines deterministic, public-safe conformance checks for a BCM provider boundary. It turns the provider profile's regulatory-integrity requirements into stable machine diagnostics without introducing another regulatory decision path.

Provider conformance is not a project verdict. A provider can pass every rule here while a particular project result is still `conditional` or `not_verified` because of project facts, local records, ambiguity, conflict, or evidence conditions.

## Stable rule catalog

| Rule | Quality dimensions | Requirement |
| --- | --- | --- |
| `BCM-CONF-001` | uniqueness | Required provider object identifiers are unique within their logical namespace. |
| `BCM-CONF-002` | consistency, conformity | Stable references resolve to the expected logical object type. |
| `BCM-CONF-003` | completeness, conformity | Every adoption identifies an adopting authority. |
| `BCM-CONF-004` | consistency, timeliness | Effective dates use `YYYY-MM-DD`; a bounded interval is not inverted. |
| `BCM-CONF-005` | accuracy, completeness | A provider claim declared `resolved` satisfies BCM's existing claim-evidence requirements for that claim kind and as-of date. |
| `BCM-CONF-006` | accuracy, conformity | Exact evidence satisfies BCM's existing exact `EvidenceLink` validation. |
| `BCM-CONF-007` | accuracy, completeness | `not_applicable` is affirmative and therefore retains a non-empty basis. |
| `BCM-CONF-008` | consistency | An evidence defect remains an evidence defect: it produces `not_verified` and is not emitted as a project-fact follow-up question. |
| `BCM-CONF-009` | coverage, completeness | Machine-checkable contract bundle identity, limitations, and declared jurisdiction coverage agree with represented provider content. |

Rule IDs are stable public identifiers. Future checks, including freshness or change-monitoring checks, must receive new IDs rather than silently changing the meaning of an existing ID.

## Reuse of BCM domain semantics

The rule layer deliberately delegates to existing BCM engine semantics:

- `BCM-CONF-005` calls `VerifyClaimEvidence` and uses the existing `SourcePolicy` model.
- `BCM-CONF-006` uses the existing exact-evidence link validator, including artifact hashes, byte sizes, anchor identity, and locators.
- `BCM-CONF-007` shares the same affirmative-`not_applicable` predicate used by runtime resolution validation.
- `BCM-CONF-008` checks the output invariant established by `ProjectVerifier`; it does not compute or replace the project verdict.

These checks therefore expose regulatory-integrity failures with stable rule IDs without recreating evidence or verdict logic.

## Logical validation projection

The public `demo/DEMO-XX/provider-bundle.json` is an illustrative logical validation projection. Its object shape is not a database schema, storage contract, compiler API, or promise about private provider implementation.

Fields used by these checks make relationships statically inspectable: stable IDs, references, effective intervals, claim status/as-of metadata, exact-evidence references, source-policy coordinates, and declared jurisdiction coverage. A private provider may map its own implementation into the same logical validation semantics without publishing its storage model or production corpus.

## Diagnostics

Diagnostics contain a stable `rule_id`, logical `path`, and human-readable `message`. Diagnostics are sorted by rule ID, path, and message so identical provider state produces deterministic output.

`BCM-CONF-009` does not attempt natural-language inference over limitation prose. It checks only facts that are safely mechanical: the contract declares limitations, the bundle declares limitations, the contract and bundle identities agree, and structured jurisdiction coverage equals represented jurisdictions.

## Public fixture boundary

CI exercises these rules with the fully synthetic `DEMO-XX` golden provider and synthetic failing fixtures. Public conformance fixtures must not contain real jurisdiction research, maintained production corpus data, restricted standards or model-code text, private provider mappings, private project data, credentials, storage locations, compiler heuristics, or deployment topology.

The normal `go test ./...` CI step is the public conformance execution path. Passing this rule set does not imply that a real project claim is verified.