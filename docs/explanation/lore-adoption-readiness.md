# Prepare Building Code Map for LORE

## Status

Building Code Map is bootstrap-ready but not yet activated. The trust root and maintainer protocol are compatible with LORE main `a8792bd594db34fcbf3065239b70ef4383f7e186`. Package installation, generated facts, accepted records, projections, and CI enforcement remain gated on the published `@laurajoyhutchins/lore@0.1.0` artifact.

## Why the repository is staged now

LORE separates deterministic facts from reviewed semantic records. Building Code Map already has strong hand-authored Diátaxis documentation and a Deciduous work graph, but it does not yet have append-only, evidence-backed semantic history or deterministic architecture projections.

The staged integration gives LORE a safe landing zone without making an unpublished package part of the public build.

## Documentation boundaries

- `README.md` and `docs/tutorials/`, `docs/how-to/`, `docs/reference/`, and `docs/explanation/` remain hand-authored.
- `docs/generated/lore/` is reserved for LORE projections.
- `.lore/extracted/` is replaceable deterministic output.
- `.lore/records/` and `.lore/transactions/` are accepted append-only history.
- `.lore/proposals/` contains untrusted proposed changes.
- Deciduous remains the active work and decision-process graph. LORE records durable accepted meaning.

## Deliberate README protection

The current LORE `readme` projection renders product-specific LORE copy. Building Code Map therefore omits that projection. Running `lore project` may only write the five paths declared under `docs/generated/lore/`.

This should remain true until LORE provides a repository-generic README projection or an explicit template contract.

## Activate after `0.1.0`

Install the exact published development dependency:

```bash
pnpm add -D @laurajoyhutchins/lore@0.1.0
pnpm exec lore version --json
```

Require the reported package name and version to match the lockfile and release evidence.

Generate deterministic facts:

```bash
pnpm exec lore extract
pnpm exec lore validate
```

Validate and apply the staged proposal:

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

Then add those three read-only checks to CI and run the complete existing frontend, backend, and regulatory verification suite.

## Maintaining knowledge

For a change with possible documentation impact:

1. create a LORE task describing paths, components, tags, audiences, and history needs;
2. run `lore context` or `lore hydrate`;
3. give the resulting context packet and committed maintainer skill to the maintainer;
4. require exactly one `lore-proposal/v1` artifact;
5. run `lore validate-proposal`;
6. review the evidence, uncertainty, identity, and supersession;
7. run `lore apply`;
8. regenerate and check projections.

Never edit accepted records, transactions, extracted facts, or generated documents directly.

## Release feedback discovered here

External adoption exposed one release-safety concern: `lore init` currently enables the `readme` projection by default while the renderer contains LORE-specific copy. A general-purpose release should either omit that projection during initialization, require explicit opt-in, or provide repository-generic templating.
