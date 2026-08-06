# Deciduous archaeology

The canonical Building Code Map archaeology is stored as `deciduous.sql.zlib.b85.part-*`, deterministic ordered parts of a base85-encoded zlib stream containing a schema-native SQL seed for Deciduous decision-graph schema `1.1.0`.

The compressed graph is a **frozen causal baseline** reconstructed from repository evidence available when the archaeology effort began. It intentionally preserves decisions, constraints, rejected paths, compatibility debt, and then-current observations instead of rewriting history whenever `main` changes.

The maintained projections under [`docs/archaeology/`](../docs/archaeology/) reconcile that baseline with the current repository. CI validates both layers:

- the frozen seed remains structurally valid at 98 nodes and 112 typed edges;
- evidence references remain inspectable;
- maintained projections exist;
- current executable-profile and readiness inventories match the repository.

Use the current-inventory wrapper:

```bash
python tools/validate_deciduous_current.py
python tools/validate_deciduous_current.py --write-sql --write-db --write-export
```

The generation command materializes `.deciduous/deciduous.sql`, `.deciduous/deciduous.db`, and `docs/graph-data.json`. Those files remain ignored because a working database may accumulate local sessions, prompts, command logs, and machine-local state. The validator always expands and checks the canonical seed in memory.

The graph uses upstream node types (`goal`, `option`, `decision`, `action`, `outcome`, `observation`, and `revisit`) and typed causal edges. Extended lifecycle labels such as `incomplete-data`, `compatibility-only`, `historical-only`, and `unresolved` remain in node metadata while native Deciduous status values stay compatible.

Root narratives are in [`narratives.md`](narratives.md). They describe the causal baseline and include a current reconciliation note rather than silently changing historical nodes.
