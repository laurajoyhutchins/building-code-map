# Maintain the Building Code Map Deciduous archaeology

## Authority order

1. Resulting repository tree and active schemas.
2. Merge ancestry and exact commit content.
3. Passing verification at the exact head.
4. Accepted LORE records and transaction receipts.
5. Merged pull-request and issue evidence.
6. Open or closed-unmerged work as proposal or historical evidence.
7. Retrospective summaries and chat context only when stronger evidence is unavailable.

Never promote a lower layer over a conflicting higher layer without an explicit unresolved node.

## Start a maintenance session

```bash
python tools/validate_deciduous_archaeology.py
.\tools\deciduous-recover.cmd
```

The Python validator checks the reproducible archaeology seed and projections. The existing Windows
wrapper checks the generated local Deciduous database and graph hygiene.

## Add a causal change

1. Identify the actual problem, observation, or incident.
2. Add or reuse a `goal` or `observation`.
3. Record the meaningful alternatives as `option` nodes.
4. Mark chosen, rejected, superseded, or unresolved paths with typed edges.
5. Add the governing `decision`.
6. Add the repository-visible `action`.
7. Add the verified `outcome`.
8. Put exact evidence in `metadata_json.evidence`.
9. Use native Deciduous status values and the extended lifecycle in
   `metadata_json.lifecycle_status`.
10. Regenerate and validate.

Do not add nodes for reading files, running routine tests, or every state file and CSS tweak. Group
them under the decision, data limitation, incident, or product concern that caused the work.

## Generate artifacts

```bash
python tools/validate_deciduous_archaeology.py --write-sql --write-db --write-export
```

The command expands the canonical compressed SQL seed, recreates the local Deciduous database, and
writes the upstream-compatible JSON export in stable order. Commit changed seed parts, narratives,
and curated projections. Do not commit the materialized SQL, database, or export.

## Evidence metadata

Preferred evidence objects:

```json
{"repository_path":"backend/internal/geocoder/sqlite.go","revision":"<full SHA>"}
{"pull_request":37,"merged":true,"merge_commit":"<full SHA>","head":"<full SHA>"}
{"issue":31}
{"external_repository":"laurajoyhutchins/building-code-ast","path":"README.md"}
```

Paths must exist in the current tree unless the metadata explicitly identifies another revision or
external repository. Use full SHAs. Keep source limitations in the node description rather than
inflating confidence.

## Lifecycle rules

- `active`: implemented and governing current behavior.
- `experimental`: implemented but not established as durable.
- `proposed`: open or draft work not on main.
- `superseded`: replaced by a later decision or implementation.
- `rejected`: considered and not selected.
- `abandoned`: stopped without a successor.
- `incomplete-data`: machinery exists but evidence coverage is materially incomplete.
- `compatibility-only`: retained to support older artifacts or callers.
- `historical-only`: retained solely as evidence of prior direction.
- `unresolved`: evidence conflicts or implementation state is not settled.

## State-coverage rule

Do not add one node per state unless a state materially changes the authority model. Update the
state-coverage projection when a state crosses a maturity threshold. A research report, canonical
profile, executable pilot, and production-ready scope are separate outcomes.

## External repositories

Reference only the decision owned elsewhere. Do not copy the full histories of Building Code AST,
Electrical Equipment Lineage, LORE, Deciduous, or predecessor interface repositories. Record which
repository owns the schema or decision and how Building Code Map consumes or depends on it.

## Adversarial review before acceptance

Ask:

- Did current architecture get projected backward?
- Did one successful fixture become a coverage claim?
- Did geometry become legal authority?
- Did a statewide rule become complete local applicability?
- Did a secondary source become controlling law?
- Did an edition lose its amendments or date?
- Did an address point become a rooftop guarantee?
- Did deterministic output become legal correctness?
- Did public deployment become data maturity?
- Did a generated view become canonical data?
- Did LORE compatibility become local reimplementation?
- Did an external repository's history get copied?
- Did compatibility debt become permanent architecture?
- Did generic trust language reappear as a current design principle?

Resolve material findings in the graph, not only in prose.
