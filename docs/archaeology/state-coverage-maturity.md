# State coverage maturity map

## Maturity vocabulary

- **Research record:** a human-readable state report or source inventory exists.
- **Canonical profile:** a versioned machine-readable profile validates.
- **Executable pilot:** the generic resolver has fixtures and returns source-backed outcomes.
- **Production-ready scope:** every returned material conclusion for a defined code-family and
  project scope has current primary-source evidence, complete timelines, local amendment and
  enforcement records, conflict resolution, and source-health checks.
- **Nationwide coverage:** production-ready scopes exist broadly enough to answer the intended
  location population, including local jurisdictions. No repository evidence supports this label
  today.

A state file is not complete jurisdictional coverage.

## Current main

| State      | Current status on `main`             | Why it was selected                                                             | Important gaps                                                                               |
| ---------- | ------------------------------------ | ------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| Colorado   | Executable pilot, partially verified | Predominantly local adoption plus state electrical override                     | Municipal adoptions, amendment detail, historical timelines, named local fixtures            |
| Florida    | Executable pilot, partially verified | Statewide building-code baseline with local enforcement and amendment follow-up | Local enforcement and amendment evidence, historical and special-project fixtures            |
| New Jersey | Executable pilot, partially verified | Statewide construction subcodes plus separate operational-fire path             | Local records, subcode timelines, fire-path details, historical and special-project fixtures |

These are the only confirmed executable pilot profiles on main at the archaeology start SHA.

## Broader report inventory

The repository has broader state research and quality-report machinery. The graph deliberately does
not equate the presence of a report with executable or local coverage. Draft PR #14 attempted a much
larger canonical report conversion but remained open and incomplete.

## Production-readiness work

Draft PR #38 introduces stricter manifests and source-health checks. Its reported state is zero
production-ready pilots and all three needing review. Because it is unmerged, those manifests are
proposal evidence rather than current architecture.

## Additional state wave

Draft PR #39 covers official-source inventory for Virginia, Oregon, North Carolina, Massachusetts,
New York, and California and drafts profiles for Virginia, Oregon, and North Carolina. All remain
needs-review research. None should be named as supported by main.

## Criteria for adding a state

A state should advance through:

1. primary-source inventory;
2. authority-model analysis;
3. adoption and date records;
4. local-adoption and amendment posture;
5. declarative profile and rule pack;
6. schema validation;
7. contrasting incorporated, unincorporated, code-family, date, conflict, and missing-evidence fixtures;
8. field-aware review;
9. source-health and staleness controls;
10. explicit production-ready scope.

Home-rule complexity, local-code hosts, inaccessible amendments, historical transitions, and special
facility authority can make a seemingly simple state substantially harder than another.
