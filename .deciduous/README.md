# Deciduous archaeology

The canonical Building Code Map archaeology is stored as the schema-native SQL seed
[`deciduous.sql`](deciduous.sql). It targets Deciduous decision-graph schema `1.1.0` and uses only
the upstream node types (`goal`, `option`, `decision`, `action`, `outcome`, `observation`, and
`revisit`) plus typed causal edges.

The repository does not commit `.deciduous/deciduous.db`. The binary database is generated from the
reviewable seed because a working Deciduous database may accumulate local sessions, prompts, command
logs, and other machine-local state.

```bash
python tools/validate_deciduous_archaeology.py
python tools/validate_deciduous_archaeology.py --write-db
```

The first command validates the SQL seed, graph topology, evidence metadata, strict goal-to-outcome
coverage, generated public export, projections, and forbidden present-day claims. The second also
writes `.deciduous/deciduous.db`, which existing `tools/deciduous*` wrappers can open.

Extended lifecycle labels such as `incomplete-data`, `compatibility-only`, `historical-only`, and
`unresolved` are stored in each node's `metadata_json.lifecycle_status`. Native Deciduous `status`
values remain compatible with the upstream graph.

Root narratives are in [`narratives.md`](narratives.md). Generated and curated projections are
under [`docs/archaeology/`](../docs/archaeology/).
