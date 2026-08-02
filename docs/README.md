# Documentation

This repository uses Diataxis to separate reader intent.

## Current Coverage

- [Tutorials](tutorials/getting-started.md): learn the local workflow by following along.
- [How-to Guides](how-to/add-boundary-layer.md): add a supported geographic boundary family.
- [How-to Guides](how-to/build-local-geocoder-snapshot.md): build a provenance-bearing SQLite snapshot from reviewed address CSV files.
- [How-to Guides](how-to/populate-state-report-template.md): turn the state report template into a source-backed report.
- [How-to Guides](how-to/use-decision-graph-for-repo-work.md): keep Deciduous aligned to the repo's real workstreams.
- [Reference](reference/runtime-and-api.md): look up scripts, endpoints, data shapes, and status behavior.
- [Reference](reference/configuration.md): see which files own runtime, formatting, linting, and generated state.
- [Reference](reference/repository-governance.md): review merge, ruleset, security, ownership, and exact-head verification policy.
- [Reference](reference/fema_bcat.md): review the FEMA BCAT source contract and update flow.
- [Explanation](explanation/tigerweb-visualization-design.md): understand the original boundary-explorer architecture choices.
- [Explanation](explanation/launch-scripts-implementation-plan.md): understand the launch-script architecture and rollout plan.
- [Design](superpowers/specs/2026-08-01-local-sqlite-geocoder-design.md): understand the address-to-point boundary and local SQLite geocoder design.

## Generated LORE Documentation

`docs/generated/` is reserved for LORE projections. Generated prose is non-authoritative and may be regenerated from repository facts and accepted semantic records. Do not edit generated files directly.

## Scope

Keep pages inside `tutorials/`, `how-to/`, `reference/`, or `explanation/` aligned to one reader intent. Design and implementation planning evidence lives under `superpowers/`. Machine-reviewed semantic knowledge lives under `.lore/records/` and changes through the LORE proposal and transaction workflow.
