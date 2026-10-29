# Independent regulatory accuracy benchmark

BCM cannot infer the accuracy of its legal conclusions from code coverage, a green test suite, or a count of represented jurisdictions. The `accuracy` package compares actual project-code-basis results with **separately reviewed primary-source expectations**.

The benchmark is an evaluator, not a second engine or an alternative project-verdict reducer. It does not call the provider, interpret legal text, or fill missing code families.

## Select cases for independent review without seeing provider answers

A transport-grid pass, source-boundary pair, temporal probe, and source-gap
probe all produce **candidate** locations, not legal truth. A blind queue
can select an intentionally diverse, reproducible set of workload IDs
before the reviewer sees any BCM provider result:

```sh
python3 -m scripts.select_accuracy_review \
  --input /secure/composed-bcm-workload.jsonl \
  --max-cases 100 \
  --output /secure/pending-review-queue.json
```

Selection balances grid, boundary, source-gap, transition-date and
unclassified samples via stable ID hashes; it never reads API result
files. The queue contains only candidate IDs/tags and an exact SHA-256 of
the source workload, **not** coordinates, addresses, project facts, legal
text, observed provider verdicts or expected code editions. It is
always `pending_independent_review` with
`legal_accuracy_reviewed:false` and `expected_results_present:false`.

The reviewer must retrieve selected cases from the **original pinned**
protected workload, separately obtain and inspect current or historical
primary legal instruments and effective timelines, preserve source
artifact identities and exact supporting text, resolve governmental
authority/geometry ambiguity and document explicit uncertainties.
Only an independently adjudicated, custody-appropriate reference
becomes a `bcm-accuracy` truth case. A valid JSON file or a nonempty
`source_ids` array is **not** proof of independent legal review.

The public contract is `schemas/accuracy-review-queue-v1.schema.json`,
paired with `schemas/bulk-query.odcs.json`. The tool creates no
automatic conversion into a legal expected result and never commits a
private case to public Git.
## Inputs and custody

Prepare `truth.json` independently of BCM's derived regulatory bundles. Each case identifies a project, jurisdiction, code family, applicability date, review date, explicit expected code states/editions, and primary-source identifiers. An independent reviewer must verify the source text and applicable legal timeline before admitting a case. A non-empty source ID is a mechanical prerequisite, not proof that review was performed.

```json
{
  "schema_version": "1",
  "cases": [
    {
      "id": "DEMO-BUILDING",
      "project_id": "DEMO-PROJECT",
      "jurisdiction_id": "DEMO-XX",
      "code_family": "building",
      "as_of": "2026-08-25",
      "reviewed_at": "2026-08-26",
      "source_ids": ["DEMO-PRIMARY-SOURCE"],
      "expected": [{"status": "resolved", "edition": "DEMO-2024"}]
    }
  ]
}
```

`results.json` wraps actual public `project-code-basis` responses without altering their content:

```json
{"results": [{"schema_version": "0.3", "project_id": "DEMO-PROJECT",
  "as_of": "2026-08-25", "jurisdiction_id": "DEMO-XX",
  "verdict": "verified",
  "code_set": [{"family": "building", "edition": "DEMO-2024", "status": "resolved"}]}]}
```

The short example is synthetic input for the evaluator, not a valid standalone production project-basis schema fixture. Actual provider results must be generated through the canonical runtime and API. Keep real project inputs, private primary-source mappings, and benchmark results requiring restricted custody in the private regulatory-data repository. Do not copy private alpha artifacts into public history.

## Run

```sh
go run ./cmd/bcm-accuracy -truth /secure/truth.json -results /secure/results.json
go run ./cmd/bcm-accuracy -strict -truth /secure/truth.json -results /secure/results.json
```

The command writes a deterministic machine-readable report with case-level outcomes and counts. Its default exit code is nonzero for an incorrect affirmative assertion or falsely verified result. `-strict` additionally rejects abstentions and missing case results. Invalid input produces a separate error code.

## Metrics that must stay separate

- **Correct:** matching jurisdiction, applicability date, family, status, and edition or editions.
- **Incorrect assertion:** an affirmative `resolved` or `not_applicable` claim contradicts the independent reference, or its jurisdiction/date is wrong.
- **Abstention:** BCM cannot make the known assertion, or the benchmarked code family is absent. Safe refusal is not a correct legal determination.
- **Missing:** no project result was supplied for a reviewed case.
- **False verified:** the full-project `verified` verdict is asserted where the reviewed case is missing, mismatched, or still legally conditional.

Cases can share one project but must use different code families. Include expected families that BCM may omit; otherwise a provider can appear accurate simply by returning only the families it understands. Represent statutory transition periods with multiple expected code states rather than forcing one edition.

This benchmark proves neither that the reviewed corpus is representative nor that its primary-source readings are true. Separate sampling, evidence review, freshness assessment, and domain-by-domain denominators remain necessary. Never label these results a nationwide accuracy rate without a documented representative sampling frame.

## Recommended benchmark strata

Review cases across incorporated/unincorporated locations, counties and townships, delegated or special authorities, fire/building/electrical/energy disciplines, historical and transition dates, local amendments, unresolved permit facts, explicit non-applicability, and contradictory sources. Colorado electrical transition and North Dakota elective township enforcement are useful initial independently reviewed cases once their actual project outputs are available.
