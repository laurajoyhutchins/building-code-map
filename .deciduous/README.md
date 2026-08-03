# Deciduous archaeology

The canonical Building Code Map archaeology is stored as
`deciduous.sql.zlib.b85.part-*`, deterministic ordered parts of a base85-encoded zlib stream of a
schema-native SQL seed targeting Deciduous decision-graph schema `1.1.0`. Compression keeps the
large evidence-bearing graph practical to review and transport through repository tools without
inventing a second graph schema.

```bash
python tools/validate_deciduous_archaeology.py
python tools/validate_deciduous_archaeology.py --write-sql --write-db --write-export
```

The second command materializes the readable `.deciduous/deciduous.sql`, local
`.deciduous/deciduous.db`, and upstream-compatible `docs/graph-data.json`. Those generated files
remain ignored because a working database may accumulate local sessions, prompts, command logs, and
other machine-local state. The validator always expands and validates the canonical seed in memory.

The graph uses upstream node types (`goal`, `option`, `decision`, `action`, `outcome`,
`observation`, and `revisit`) and typed causal edges. Extended lifecycle labels such as
`incomplete-data`, `compatibility-only`, `historical-only`, and `unresolved` are stored in each
node's `metadata_json.lifecycle_status`; native Deciduous `status` values remain compatible.

Root narratives are in [`narratives.md`](narratives.md). Curated projections are under
[`docs/archaeology/`](../docs/archaeology/).
