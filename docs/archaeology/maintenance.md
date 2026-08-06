# Archaeology maintenance

## Authority of each artifact

- `.deciduous/deciduous.sql.zlib.b85.part-*` is the frozen causal baseline.
- `.deciduous/narratives.md` explains the baseline and records a current reconciliation note.
- `docs/archaeology/*.md` are maintained projections of current repository state and historical context.
- `backend/data/regulatory/*.json` are generated executable profiles.
- `reports/production/*.json` are declared production-readiness scopes.
- LORE records and projections preserve accepted semantic repository knowledge.
- None of these replaces official regulatory evidence.

## Validation

Run:

```bash
python tools/validate_deciduous_current.py
python tools/validate_deciduous_current.py --write-sql --write-db --write-export
```

The wrapper preserves the frozen graph validator and supplies the current repository inventory. CI checks:

- Deciduous schema `1.1.0`;
- 98 frozen nodes and 112 frozen edges;
- semantic identifiers, evidence references, graph topology, and root narratives;
- required archaeology projections;
- six executable profile files;
- three production-scope manifests;
- result schema `1.0`;
- geocoder schema `1`;
- removal of obsolete trust-badge copy.

Generated SQL, SQLite, and JSON graph material remains ignored. The compressed seed is the reviewable canonical graph artifact.

## Updating current projections

Update projections when a merged change alters a maintained claim, including:

- public API trust boundaries;
- snapshot admission or compatibility behavior;
- executable profile inventory;
- production-readiness inventory;
- regulatory result schema;
- geocoder schema;
- repository ownership boundaries;
- known risks or deferred systems.

Do not change frozen graph nodes merely to make their old observations look current. Instead:

1. preserve the historical node and evidence;
2. update the reconciliation in narratives and projections;
3. add a new graph revision only when intentionally extending the archaeology baseline;
4. update the validator's live inventory checks;
5. run all repository verification at the exact branch head.

## Adding a state profile

A new executable profile requires more than adding JSON:

1. add or promote a compiler-discoverable pilot report;
2. generate the profile and source-health rule pack;
3. add resolver regressions for current, transition, pending, and missing-evidence behavior;
4. update nationwide quality reporting;
5. update `tools/validate_deciduous_current.py` and maintained state-coverage projections;
6. keep verification status and unresolved local records explicit.

A production-ready scope additionally requires a manifest, continuous verified timelines inside the declared scope, primary-source evidence, source health, and evidence-backed fixtures. Do not infer statewide completeness from a passing scope.

## Review discipline

- Reconcile PR descriptions with their final exact heads.
- Regenerate generated artifacts after branch integration, not before.
- Treat open PRs as proposed, not current.
- Preserve containment versus authority, publication versus adoption, and test coverage versus geographic completeness.
- Prefer deleting stale claims over layering contradictory documentation.
