# Use the Decision Graph for Repo Work

## Goal

Capture the repo's meaningful project choices in Deciduous so future sessions can recover the current model, not just the current files.

## When To Use This

Use this workflow when a task changes one of the repo's actual workstreams:

- state building-code reports under `reports/`
- canonical jurisdiction geometry and source provenance
- FEMA BCAT validation and discrepancy review
- production vector-tile serving from mirrored boundary data
- frontend/backend behavior that changes the delivered map experience

Do not use it for your own reading, planning, or shell activity.

## Prerequisites

- Run commands from the repo root.
- Use `.\tools\deciduous.cmd` instead of assuming `deciduous` is on `PATH`.
- Use `.\tools\deciduous-audit.cmd` when you want a quick graph-hygiene check.
- Use `.\tools\deciduous-recover.cmd` at session start to rehydrate the current graph state.
- Use `.\tools\deciduous-work.cmd` when you are starting a new repo-visible initiative and want to create the goal node first.
- Have the verbatim user request available when you create a new goal.

## Steps

1. Check the current graph state before starting substantial work:

```powershell
.\tools\deciduous-recover.cmd
```

This bundles the current pulse summary with the repo audit helper. If you want the full pulse output, run `.\tools\deciduous.cmd pulse` directly.

2. If the user request starts a new repo-visible initiative, create a root goal with the verbatim prompt:

```powershell
@'
<paste the full user request here>
'@ | .\tools\deciduous-work.cmd "Meaningful project goal"
```

This creates the goal node and then runs the recovery summary so you can see the updated graph state immediately.

3. Add option nodes for the real approaches the repo could take:
   - For state reports, options should distinguish official-source-first work from shortcut or secondary-source approaches.
   - For jurisdiction geometry, options should distinguish real NERIS service areas from proxy or fallback geometry.
   - For BCAT, options should distinguish validation-only use from canonical-source use.
   - For production tile work, add options even when the repo has not committed to one yet.

4. Link each option to its parent goal immediately:

```powershell
.\tools\deciduous.cmd link <goal_id> <option_id> -r "possible_approach"
```

5. When the repo has clearly chosen an approach, add a decision node and link it from the chosen option:

```powershell
.\tools\deciduous.cmd add decision "Chosen approach" -c 90
.\tools\deciduous.cmd link <option_id> <decision_id> -r "chosen"
```

6. Before you make the repo-visible change, add an action node for the implementation step:

```powershell
.\tools\deciduous.cmd add action "Implement the chosen change" -c 85
.\tools\deciduous.cmd link <decision_id> <action_id> -r "implementation"
```

7. After verification, add an outcome node that says what actually happened:

```powershell
.\tools\deciduous.cmd add outcome "Verified result of the change" -c 90
.\tools\deciduous.cmd link <action_id> <outcome_id> -r "result"
```

8. Re-run the pulse and look for orphan goals, missing options, or stale open questions:

```powershell
.\tools\deciduous.cmd pulse --summary
```

9. Run the repo audit helper before you stop if you changed the graph substantially:

```powershell
.\tools\deciduous-audit.cmd
```

Use `.\tools\deciduous-audit.cmd -Strict` when you want the command to fail on missing options, missing actions, missing outcomes, or orphan non-goal nodes.

## What To Capture In This Repo

- Decisions about canonical legal/code authority versus validation data.
- Decisions about what counts as canonical geometry and when fallbacks are allowed.
- Decisions that change the runtime boundary between the delivered cache-backed API and deferred production tile serving.
- Outcomes that a future session would need in order to trust the current direction.

## What Not To Capture

- "Read the codebase"
- "Planned the next step"
- "Ran tests"
- "Opened the docs"

Those are session process details, not project decisions.

## Result

The graph stays useful for handoff and recovery because it records the repo's actual model choices, not just a trail of generic activity.
