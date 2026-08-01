# LORE Adoption Readiness Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Install a release-compatible LORE trust root and bootstrap task without introducing an unpublished package dependency or committing an inherently stale proposal.

**Architecture:** Copy the current LORE initialization assets into Building Code Map, configure only non-README projections under `docs/generated/lore/`, and commit the trust root and task. After `0.1.0` is published, install and commit the package, generate one proposal as an ignored working-tree artifact bound to current `HEAD`, apply it transactionally before any head movement, and commit only the accepted outputs.

**Tech Stack:** Git, YAML, JSON Schema draft 2020-12, Node.js 22, pnpm, future `@laurajoyhutchins/lore@0.1.0`.

## Global Constraints

- Do not add an unpublished or Git-based LORE package dependency.
- Do not generate or edit accepted records, transactions, extracted facts, or LORE projections manually.
- Do not configure LORE to overwrite `README.md`.
- Preserve the existing Diataxis and Deciduous workflows.
- Keep `.lore/proposals/` ignored and uncommitted.
- Bind each proposal to the current committed `HEAD`; any head movement requires regeneration.
- Anchor every semantic claim to a full 40-character Git commit and an existing path.
- Treat schema, skill, evidence, revision, and identity failures as blocking.

---

### Task 1: Install the bootstrap trust root

**Files:**
- Create: `BOOTSTRAP.md`
- Create: `lore.yaml`
- Create: `schemas/manifest.schema.json`
- Create: `schemas/record.schema.json`
- Create: `schemas/proposal.schema.json`
- Create: `schemas/task.schema.json`
- Create: `schemas/hydration.schema.json`
- Create: `schemas/extracted-facts.schema.json`
- Create: `schemas/transaction.schema.json`
- Create: `skills/maintain-repository-documentation/SKILL.md`
- Create: `skills/maintain-repository-documentation/INPUTS.md`
- Create: `skills/maintain-repository-documentation/OUTPUTS.md`
- Create: `skills/maintain-repository-documentation/schemas/proposal.schema.json`
- Modify: `.prettierignore`

**Interfaces:**
- Consumes: LORE initialization assets at the documented compatibility baseline.
- Produces: A target-repository trust root that the released CLI can preserve and validate.

- [ ] **Step 1: Copy the initialization assets without semantic modification**

Copy the files listed above from the compatibility baseline.

- [ ] **Step 2: Configure the repository manifest**

Set repository ID `building-code-map`, generated-doc root `docs/generated/lore`, the four supported extractors, and all non-README projections.

- [ ] **Step 3: Preserve trust-root schema bytes**

Exclude the copied root LORE schemas from Building Code Map's separate Prettier policy rather than rewriting the upstream schema representation.

- [ ] **Step 4: Verify static invariants**

Parse every JSON schema, confirm the two proposal schemas are byte-identical, and confirm the skill SHA-256 is `a4638b4544a102a51a09878406c3a2820bc6d4038b5c35b2898d2966ca037144`.

- [ ] **Step 5: Commit the trust root**

```bash
git add BOOTSTRAP.md lore.yaml schemas skills .prettierignore
git commit -m "docs: install LORE bootstrap trust root"
```

### Task 2: Commit the bootstrap task and release runbook

**Files:**
- Create: `.lore/tasks/bootstrap-repository-knowledge.yaml`
- Modify: `.gitignore`
- Create: `docs/explanation/lore-adoption-readiness.md`
- Create: `docs/superpowers/specs/2026-08-01-lore-adoption-readiness-design.md`
- Create: `docs/superpowers/plans/2026-08-01-lore-adoption-readiness.md`
- Modify: `docs/README.md`

**Interfaces:**
- Consumes: Task schema version 1 and the existing repository documentation structure.
- Produces: A repeatable maintenance request and a human activation runbook.

- [ ] **Step 1: Add the task artifact**

Use task ID `bootstrap-repository-knowledge`, list the public lookup, explorer, backend, regulatory, and governance evidence paths, request architecture and maintainer audiences, and set `history: false`.

- [ ] **Step 2: Ignore proposal inputs**

Add `.lore/proposals/` to `.gitignore` with a note that proposals are ephemeral inputs bound to current `HEAD`.

- [ ] **Step 3: Document the release gate and exact-head proposal lifecycle**

State the exact `0.1.0` installation, extraction, context, validation, apply, project, and check commands. Explain why README projection and CI activation are deferred and why proposals must not be committed.

- [ ] **Step 4: Link the readiness guide**

Add the guide to `docs/README.md` under Explanation.

- [ ] **Step 5: Commit the readiness packet**

```bash
git add .lore/tasks .gitignore docs
git commit -m "docs: define LORE adoption readiness"
```

### Task 3: Activate after the stable package is published

**Files:**
- Modify after release: `package.json`
- Modify after release: `pnpm-lock.yaml`
- Modify after release: `.github/workflows/ci.yml`
- Generate after release: `.lore/extracted/*.yaml`
- Generate ephemerally after release: `.lore/proposals/bootstrap-context.yaml`
- Generate ephemerally after release: `.lore/proposals/bootstrap-repository-knowledge.yaml`
- Generate after release: `.lore/records/**`
- Generate after release: `.lore/transactions/*.yaml`
- Generate after release: `docs/generated/lore/*.md`

**Interfaces:**
- Consumes: Published `@laurajoyhutchins/lore@0.1.0`, the committed task, and current repository evidence.
- Produces: Validated accepted history, deterministic projections, and CI drift enforcement.

- [ ] **Step 1: Install and commit the exact stable version**

```bash
pnpm add -D @laurajoyhutchins/lore@0.1.0
pnpm exec lore version --json
git add package.json pnpm-lock.yaml
git commit -m "chore: install LORE 0.1.0"
```

Require package name `@laurajoyhutchins/lore`, version `0.1.0`, Node.js 22 or newer, and schema version 1 for manifest, record, proposal, task, hydration, and transaction.

- [ ] **Step 2: Generate deterministic facts and validate**

```bash
pnpm exec lore extract
pnpm exec lore validate
```

- [ ] **Step 3: Generate bounded maintainer context at current HEAD**

```bash
pnpm exec lore context .lore/tasks/bootstrap-repository-knowledge.yaml > .lore/proposals/bootstrap-context.yaml
git rev-parse HEAD
```

Use the committed maintainer skill and context packet to produce exactly one proposal. Set its `base_revision` to the printed current `HEAD`, include the skill digest from that revision, and keep the file uncommitted.

- [ ] **Step 4: Validate and apply without moving HEAD**

```bash
pnpm exec lore validate-proposal .lore/proposals/bootstrap-repository-knowledge.yaml
pnpm exec lore apply .lore/proposals/bootstrap-repository-knowledge.yaml
```

If any commit or merge changes `HEAD`, regenerate the proposal before applying it.

- [ ] **Step 5: Generate and check projections**

```bash
pnpm exec lore project
pnpm exec lore validate
pnpm exec lore extract --check
pnpm exec lore project --check
```

- [ ] **Step 6: Commit accepted outputs, not proposals**

```bash
git add package.json pnpm-lock.yaml .lore/extracted .lore/records .lore/transactions docs/generated/lore
git commit -m "docs: activate LORE repository knowledge"
```

Confirm `.lore/proposals/` remains absent from the commit.

- [ ] **Step 7: Add CI enforcement**

Add release-pinned install plus `lore extract --check`, `lore validate`, and `lore project --check` to the read-only verification workflow.

- [ ] **Step 8: Run the full repository verification**

```bash
pnpm check
pnpm lint
pnpm format:check
pnpm test
pnpm build
pnpm backend:test
pnpm backend:lint
pnpm regulatory:test
pnpm regulatory:compile:check
pnpm exec lore extract --check
pnpm exec lore validate
pnpm exec lore project --check
```
