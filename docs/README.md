# Documentation

This repository uses Diataxis to separate reader intent.

## Current Coverage

- [Tutorials](tutorials/getting-started.md): learn the local workflow by following along.
- [How-to Guides](how-to/add-boundary-layer.md): complete one task without extra background.
- [How-to Guides](how-to/populate-state-report-template.md): turn the state report template into a source-backed report.
- [How-to Guides](how-to/use-decision-graph-for-repo-work.md): keep Deciduous aligned to the repo's real workstreams.
- [Reference](reference/runtime-and-api.md): look up scripts, endpoints, and data shapes.
- [Reference](reference/configuration.md): see which files own formatting, linting, and generated state.
- [Reference](reference/repository-governance.md): review merge, ruleset, security, ownership, and exact-head verification policy.
- [Reference](reference/fema_bcat.md): review the FEMA BCAT source contract and update flow.
- [Reference](reference/regulatory-resolution.md): inspect the evidence-to-resolution data model and API contract.
- [Explanation](explanation/tigerweb-visualization-design.md): understand the product and architecture choices.
- [Explanation](explanation/launch-scripts-implementation-plan.md): understand the launch-script architecture and rollout plan.
- [Explanation](explanation/lore-adoption-readiness.md): understand the staged LORE trust root, documentation boundaries, and stable-release activation gate.

## Scope

The docs in this scaffold stay aligned to one intent per page. Keep pages inside `tutorials/`, `how-to/`, `reference/`, or `explanation/`, and split mixed-intent topics into separate pages instead of forcing one document to do everything.

LORE-generated architecture and decision views belong only under `docs/generated/lore/`. They are projections, not authoritative source material, and must not replace the hand-authored Diataxis pages.
