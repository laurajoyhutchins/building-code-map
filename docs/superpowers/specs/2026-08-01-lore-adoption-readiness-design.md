# LORE Adoption Readiness Design

**Status:** Approved  
**Date:** 2026-08-01  
**Target repository:** `laurajoyhutchins/building-code-map`  
**LORE compatibility baseline:** `laurajoyhutchins/LORE@a8792bd594db34fcbf3065239b70ef4383f7e186`  
**Release candidate observed:** `@laurajoyhutchins/lore@0.1.0` work on `ac57c8ae03a7aacb8f13eab8abfbb281e91a0f51`

## 1. Goal

Prepare Building Code Map to adopt the versioned LORE CLI without making the public repository depend on an unpublished package or allowing a prerelease projection to overwrite established documentation.

## 2. Architecture

The repository receives LORE's current bootstrap trust root, strict schemas, and provider-neutral maintainer skill. `lore.yaml` enables deterministic extraction for repository metadata, package scripts, TypeScript modules, and Vitest tests. Human-facing projections are isolated under `docs/generated/lore/`.

The existing `README.md` and Diátaxis documentation remain hand-authored. The LORE `readme` projection is deliberately disabled because the current renderer contains LORE-product-specific prose and is unsafe for an external repository.

An initial task describes the bootstrap maintenance job. After the trust-root commit exists, the maintainer skill produces one `lore-proposal/v1` artifact anchored to that exact commit. Accepted records and generated projections are not hand-authored.

## 3. Knowledge boundaries

LORE records describe durable repository meaning:

- repository purpose and support boundary;
- frontend, backend, snapshot, regulatory-data, and documentation components;
- data-flow relationships;
- architectural decisions;
- provenance and redistribution constraints;
- exact-head verification procedures.

Deciduous continues to describe active work: goals, alternatives, decisions in progress, implementation actions, and outcomes. A verified outcome may justify a LORE proposal, but Deciduous state is not automatically authoritative semantic history.

Diátaxis remains the documentation system for tutorials, how-to guides, reference pages, and explanations. LORE projections provide derived architecture and decision views rather than replacing those pages.

## 4. Release gate

This change does not add `@laurajoyhutchins/lore` to `package.json`, modify CI, run generated projections, or apply the bootstrap proposal. Activation waits for a published, installation-tested `0.1.0` package.

After release, activation is:

```bash
pnpm add -D @laurajoyhutchins/lore@0.1.0
pnpm exec lore version --json
pnpm exec lore extract
pnpm exec lore validate
pnpm exec lore validate-proposal .lore/proposals/bootstrap-repository-knowledge.yaml
pnpm exec lore apply .lore/proposals/bootstrap-repository-knowledge.yaml
pnpm exec lore project
pnpm exec lore validate
pnpm exec lore project --check
```

The package version, schema versions, skill digest, generated output, and exact-head CI result must all be reviewed before merge.

## 5. Safety properties

- Existing README content is never a generated target.
- Generated output is confined to `docs/generated/lore/`.
- The maintainer skill is committed before any proposal claims its digest.
- Every semantic assertion cites a full Git commit and repository path.
- Initial records enter through one untrusted proposal and transactional application.
- Missing evidence, schema drift, stale revisions, or ambiguous record identity fail closed.
- Hydrated databases, protected source text, and restricted operational data remain outside LORE records and projections.

## 6. Verification

Before activation, verify:

1. all copied schemas parse as JSON;
2. root and skill proposal schemas are byte-identical;
3. the committed skill SHA-256 is `a4638b4544a102a51a09878406c3a2820bc6d4038b5c35b2898d2966ca037144`;
4. `lore.yaml` excludes the `readme` projection;
5. the bootstrap task conforms to task schema version 1;
6. the proposal base revision is the exact trust-root commit;
7. every proposal evidence path exists at that revision;
8. no accepted record, transaction, extracted fact, or generated projection is created manually.
