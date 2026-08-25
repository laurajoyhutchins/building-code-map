# Project Code Verification

Project Code Verification is the primary public Building Code Map workflow. It asks a project-level question instead of requiring callers to inspect one code family at a time:

> Given this location, applicability date, and the project facts currently known, what code set can BCM verify, what remains conditional, and what cannot yet be verified?

## One decision path

`ProjectVerifier.VerifyProject` calls the same `Runtime.Resolve` path used by lower-level engine clients. It does not maintain a parallel regulatory model. The project verifier only reduces the canonical runtime result into a project-level answer.

The input does not contain a required code-family selector. A provider may therefore return all relevant code or regulatory regime entries for the project in one resolution.

## Verdict reduction

The project verdict is deterministic:

- `verified` when every returned entry is `resolved` or affirmatively `not_applicable` and there are no unresolved requirements;
- `conditional` when a result is partially resolved, requires a local record, or contains a structured unresolved requirement;
- `not_verified` when any required entry is unsupported, ambiguous, conflicting, or lacks sufficient evidence.

Missing evidence never becomes `not_applicable`. The runtime separately requires an affirmative basis for every `not_applicable` result.

## Unresolved work

Providers return unresolved requirements using three public categories:

- `project_fact`: information the project team can supply and that can change the answer;
- `local_record`: a project or authority record that must be obtained;
- `evidence_defect`: evidence is absent, ambiguous, conflicting, or otherwise insufficient.

Only `project_fact` requirements become follow-up questions. This keeps user questioning deterministic and prevents evidence defects from being disguised as questions the project team can answer.

## Reproducibility

A `project-code-basis` record preserves the normalized inputs, user-supplied facts, BCM-derived facts, per-code statuses, unresolved requirements, evidence identifiers, and runtime provenance. It is a project-scoped result, not a bulk regulatory corpus export.
