# LORE Adoption Readiness Design

**Status:** Approved  
**Date:** 2026-08-01  
**Target repository:** `laurajoyhutchins/building-code-map`  
**LORE compatibility baseline:** `laurajoyhutchins/LORE@a8792bd594db34fcbf3065239b70ef4383f7e186`  
**Release candidate observed:** `@laurajoyhutchins/lore@0.1.0` work on `ac57c8ae03a7aacb8f13eab8abfbb281e91a0f51`

## 1. Goal

Prepare Building Code Map to adopt the versioned LORE CLI without making the public repository depend on an unpublished package, allowing a prerelease projection to overwrite established documentation, or committing a proposal that cannot satisfy LORE's exact-HEAD transaction gate.

## 2. Architecture

The repository receives LORE's current bootstrap trust root, strict schemas, and provider-neutral maintainer skill. `lore.yaml` enables deterministic extraction for repository metadata, package scripts, TypeScript modules, and Vitest tests. Human-facing projections are isolated under `docs/generated/lore/`.

The existing `README.md` and Diataxis documentation remain hand-authored. The LORE `readme` projection is deliberately disabled because the current renderer contains LORE-product-specific prose and is unsafe for an external repository.

A committed task describes the bootstrap maintenance job. After the stable package is installed and committed, the maintainer skill produces one uncommitted `lore-proposal/v1` artifact anchored to that exact current `HEAD`. LORE validates and applies it before any intervening commit. Accepted records, transactions, extracted facts, and generated projections are then committed; the proposal input is ignored and discarded.

## 3. Knowledge boundaries

LORE records describe durable repository meaning:

- repository purpose and boundaries;
- the public building-code lookup and technical GIS explorer;
- frontend, backend, snapshot, regulatory-data, and documentation components;
- data-flow relationships;
- architectural decisions;
- provenance and redistribution constraints;
- exact-head verification procedures.

Deciduous continues to describe active work: goals, alternatives, decisions in progress, implementation actions, and outcomes. A verified outcome may justify a LORE proposal, but Deciduous state is not automatically authoritative semantic history.

Diataxis remains the documentation system for tutorials, how-to guides, reference pages, and explanations. LORE projections provide derived architecture and decision views rather than replacing those pages.

## 4. Proposal lifecycle

`planTransaction` requires `proposal.base_revision` to equal current repository `HEAD`. A committed proposal cannot meet that contract because the commit containing it changes `HEAD`, and a Git object cannot name itself.

Therefore:

- `.lore/proposals/` is ignored by Git;
- proposals are generated only after all prerequisite package and trust-root changes are committed;
- the proposal names the current `HEAD` and skill digest at that revision;
- `validate-proposal` and `apply` run before any commit or head movement;
- accepted outputs are committed, but the proposal input is not;
- any head movement invalidates the proposal and requires regeneration.

## 5. Release gate

This readiness change does not add `@laurajoyhutchins/lore` to `package.json`, modify CI, run extraction, create accepted records or transactions, or generate projections. Activation waits for a published, installation-tested `0.1.0` package.

After release, activation is:

```bash
pnpm add -D @laurajoyhutchins/lore@0.1.0
pnpm exec lore version --json
git add package.json pnpm-lock.yaml
git commit -m "chore: install LORE 0.1.0"
pnpm exec lore extract
pnpm exec lore validate
pnpm exec lore context .lore/tasks/bootstrap-repository-knowledge.yaml > .lore/proposals/bootstrap-context.yaml
git rev-parse HEAD
# Produce one uncommitted proposal bound to the current HEAD.
pnpm exec lore validate-proposal .lore/proposals/bootstrap-repository-knowledge.yaml
pnpm exec lore apply .lore/proposals/bootstrap-repository-knowledge.yaml
pnpm exec lore project
pnpm exec lore validate
pnpm exec lore extract --check
pnpm exec lore project --check
```

The package version, schema versions, skill digest, evidence revisions, generated output, and exact-head CI result must all be reviewed before merge.

## 6. Safety properties

- Existing README content is never a generated target.
- Generated output is confined to `docs/generated/lore/`.
- The maintainer skill is committed before any proposal claims its digest.
- Proposals are ephemeral and bound to the exact current committed `HEAD`.
- Every semantic assertion cites a full Git commit and repository path.
- Initial records enter through one untrusted proposal and transactional application.
- Missing evidence, schema drift, stale revisions, or ambiguous record identity fail closed.
- Hydrated databases, protected source text, and restricted operational data remain outside LORE records and projections.

## 7. Verification

Before activation, verify:

1. all copied schemas parse as JSON;
2. root and skill proposal schemas are byte-identical;
3. the committed skill SHA-256 is `a4638b4544a102a51a09878406c3a2820bc6d4038b5c35b2898d2966ca037144`;
4. `lore.yaml` excludes the `readme` projection;
5. `.lore/proposals/` is ignored;
6. the bootstrap task conforms to task schema version 1 and includes the current public lookup and explorer paths;
7. a generated proposal's base revision equals current `HEAD` immediately before validation and application;
8. every proposal evidence path exists at its cited revision;
9. no accepted record, transaction, extracted fact, or generated projection is created manually.
