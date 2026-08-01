# Prepare Building Code Map for LORE

## Status

Building Code Map is bootstrap-ready but not yet activated. The repository contains a LORE-compatible manifest, trust-root schemas, maintainer skill, and bootstrap task. Package installation, deterministic extraction, semantic proposal generation, accepted records, projections, and CI enforcement remain gated on the published `@laurajoyhutchins/lore@0.1.0` artifact.

## Why the repository is staged now

LORE separates deterministic facts from reviewed semantic records. Building Code Map already has hand-authored Diataxis documentation and a Deciduous work graph, but it does not yet have append-only, evidence-backed semantic history or deterministic architecture projections.

The staged integration gives the released CLI a safe landing zone without making an unpublished package part of the public build.

## Documentation boundaries

- `README.md` and `docs/tutorials/`, `docs/how-to/`, `docs/reference/`, and `docs/explanation/` remain hand-authored.
- `docs/generated/lore/` is reserved for LORE projections.
- `.lore/extracted/` is replaceable deterministic output.
- `.lore/records/` and `.lore/transactions/` are accepted append-only history.
- `.lore/proposals/` contains ephemeral, untrusted maintainer output and is ignored by Git.
- Deciduous remains the active work and decision-process graph. LORE records durable accepted meaning.

## Deliberate README protection

The current LORE `readme` projection renders product-specific LORE copy. Building Code Map therefore omits that projection. Running `lore project` may only write the five paths declared under `docs/generated/lore/`.

This should remain true until LORE provides a repository-generic README projection or an explicit template contract.

## Proposal lifecycle

A LORE proposal is bound to an exact committed `HEAD`. `lore apply` rejects it after any intervening commit. The proposal must therefore remain uncommitted:

1. Start from a clean, committed branch head.
2. Generate the maintainer context from `.lore/tasks/bootstrap-repository-knowledge.yaml`.
3. Write the maintainer's single `lore-proposal/v1` artifact under the ignored `.lore/proposals/` directory.
4. Set `base_revision` to the current `git rev-parse HEAD` value and use the maintainer skill digest from that revision.
5. Run `validate-proposal` and `apply` before committing or moving `HEAD`.
6. Commit the accepted records, transaction receipt, extracted facts, generated projections, and package changes. Do not commit the proposal input.

If `HEAD` moves, discard or regenerate the proposal rather than weakening the stale-base check.

## Activate after `0.1.0`

Create an activation branch from current `main`, install the exact published dependency, and commit the package and lockfile before generating a proposal:

```bash
pnpm add -D @laurajoyhutchins/lore@0.1.0
pnpm exec lore version --json
git add package.json pnpm-lock.yaml
git commit -m "chore: install LORE 0.1.0"
```

Require the reported package name and version to match the lockfile and release evidence.

Generate deterministic facts and validate the repository:

```bash
pnpm exec lore extract
pnpm exec lore validate
```

Generate bounded maintainer context without changing `HEAD`:

```bash
pnpm exec lore context .lore/tasks/bootstrap-repository-knowledge.yaml > .lore/proposals/bootstrap-context.yaml
git rev-parse HEAD
```

Use the committed maintainer skill and that context to produce exactly one uncommitted proposal at `.lore/proposals/bootstrap-repository-knowledge.yaml`. Then validate and apply it immediately:

```bash
pnpm exec lore validate-proposal .lore/proposals/bootstrap-repository-knowledge.yaml
pnpm exec lore apply .lore/proposals/bootstrap-repository-knowledge.yaml
```

Generate derived documentation and prove convergence:

```bash
pnpm exec lore project
pnpm exec lore validate
pnpm exec lore extract --check
pnpm exec lore project --check
```

Commit the resulting authoritative and generated artifacts, then add the three read-only checks to CI and run the complete existing frontend, backend, and regulatory verification suite.

## Expected initial knowledge coverage

The bootstrap maintainer should consider records for:

- repository purpose and support boundaries;
- the public building-code lookup and technical GIS explorer;
- the Go API, local geographic snapshot, and regulatory-resolution pipeline;
- frontend-to-API, API-to-snapshot, and report-to-catalog relationships;
- local-first deterministic resolution, uncertainty preservation, and SQLite primacy;
- source provenance, redistribution limits, database publication boundaries, and exact-head verification;
- the separation between Diataxis, Deciduous, and LORE projections.

The maintainer must derive the final identities and assertions from the exact activation revision rather than copying a prerelease proposal.

## Maintaining knowledge

For a change with possible documentation impact:

1. create or update a LORE task describing paths, components, tags, audiences, and history needs;
2. run `lore context` or `lore hydrate` at a clean committed `HEAD`;
3. give the resulting context packet and committed maintainer skill to the maintainer;
4. require exactly one uncommitted `lore-proposal/v1` artifact bound to that `HEAD`;
5. run `lore validate-proposal`;
6. review the evidence, uncertainty, identity, and supersession;
7. run `lore apply` before any commit or head movement;
8. regenerate, check, and commit the accepted outputs.

Never edit accepted records, transactions, extracted facts, or generated documents directly.

## Release feedback discovered here

External adoption exposed two release-safety concerns:

- `lore init` currently enables the `readme` projection by default while the renderer contains LORE-specific copy.
- `lore apply` requires the proposal base to equal current `HEAD`, so public documentation must describe proposals as ephemeral inputs rather than checked-in artifacts.
