# National coverage evidence frontier

This is an **internal source-safe research worklist**, not a second
regulatory verifier, ordinance extractor, authority resolver, or assertion
of real legal coverage. It reuses the existing family ledger's schema
and integrity admission gate in scripts/code-family-coverage-ledger.mjs.

## Run

Produce an operator-supplied expected JSON input with exactly these keys:
as_of (ISO date), project_class (nonempty scope), and
projection_snapshot_sha256 (64 lowercase hex characters from a separately
approved government identity projection). The following invocation accepts
already reviewed, same-snapshot family ledger files:

    node scripts/national-coverage-frontier.mjs \
      /private/expected.json /private/building-ledger.json \
      /private/energy-ledger.json

Pass no family ledger files to report all six required families as missing,
rather than as having zero coverage. Malformed, inconsistent or rejected
inputs never emit a partial JSON result.

## Meaning

The resulting classification is evidence_gap_queue_only. Each work item
carries code family, candidate governmental entity ID (or null for a
family-wide gap), lane, reason, and existing source identity if present.
There are three independent work lanes:

- inventory: absent family ledger, unaudited denominator, or government
  ID in an audited denominator without a record. An unaudited denominator
  cannot generate an invented missing-government list.
- adoption: no source candidate, a candidate lacking independent review,
  contradictory source evidence, or unresolved/conflicting legal authority.
- project: missing site, provider or execution evidence. This is explicitly
  **not** jurisdiction-wide adoption work and does not increase a national
  legal-source gap count.

No URL, digest, source-review label or receipt-shaped string self-certifies
a legal or project verdict. This report records statuses and work items,
not a nationally verified numerator or percentage. It carries no bulk
private regulatory corpus or publisher geometry.

## Ownership and release requirements

Arcata Jurisdictions owns governmental identity candidates, source-native
geography and protected source custody. Arcata RSO owns original legal
sources, rights, enactment/amendment review, contradictions and regulatory
proof obligations. The one BCM Runtime, SourcePolicy, SiteAuthorityPolicy,
ProjectVerifier and POST /v1/project-code-basis remain authoritative for
project applicability. Do not implement a second verdict reducer here.

This CLI does not write protected storage, independently review original
documents, approve a provider, run a project, or deploy cloud resources.
Keep output in private reviewed custody when it references a maintained
jurisdiction corpus; use create-only generations and separately
authenticated readback for accepted receipts. Output counts are numbers
of *work items*, not unique legal regimes or verified projects. A family
cannot inherit another family's coverage or a state's legal source
without separately establishing the law's scope and priority.

The output is an internal planning report, not yet an accepted
cross-repository data contract. Before exporting these rows into another
repository, add paired JSON Schema 2020-12 + ODCS and W3C PROV-O
source/review lineage under BCM's existing contract ownership.

Work items align with BCM #87 national family ledger, BCM #74 Denver
real-source acceptance, BCM #90/#91 site and source review admission,
and RSO #77/#78 source proof closure. The synthetic public demo and
this worklist are not national verification.
