# LORE Adoption Readiness Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Install a release-compatible LORE trust root and stage an evidence-backed bootstrap proposal without introducing an unpublished package dependency.

**Architecture:** Copy the current LORE initialization assets into Building Code Map, configure only non-README projections under `docs/generated/lore/`, and commit those assets before creating the proposal that authenticates the committed skill. Leave package installation, extraction, proposal application, projection, and CI enforcement release-gated.

**Tech Stack:** Git, YAML, JSON Schema draft 2020-12, Node.js 22, pnpm, future `@laurajoyhutchins/lore@0.1.0`.

## Global Constraints

- Do not add an unpublished or Git-based LORE package dependency.
- Do not generate or edit accepted records, transactions, extracted facts, or LORE projections manually.
- Do not configure LORE to overwrite `README.md`.
- Preserve the existing Diátaxis and Deciduous workflows.
- Anchor every proposal claim to a full 40-character Git commit and an existing path.
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

**Interfaces:**
- Consumes: LORE initialization assets at the documented compatibility baseline.
- Produces: A target-repository trust root that the released CLI can preserve and validate.

- [ ] **Step 1: Copy the initialization assets without modification**

Copy the files listed above from the compatibility baseline.

- [ ] **Step 2: Configure the repository manifest**

Set repository ID `building-code-map`, generated-doc root `docs/generated/lore`, the four supported extractors, and all non-README projections.

- [ ] **Step 3: Verify static invariants**

Parse every JSON schema, confirm the two proposal schemas are byte-identical, and confirm the skill SHA-256 is `a4638b4544a102a51a09878406c3a2820bc6d4038b5c35b2898d2966ca037144`.

- [ ] **Step 4: Commit the trust root**

```bash
git add BOOTSTRAP.md lore.yaml schemas skills
git commit -m "docs: install LORE bootstrap trust root"
```

Record this exact commit for the proposal base revision.

### Task 2: Stage the bootstrap maintenance task

**Files:**
- Create: `.lore/tasks/bootstrap-repository-knowledge.yaml`
- Create: `docs/explanation/lore-adoption-readiness.md`
- Create: `docs/superpowers/specs/2026-08-01-lore-adoption-readiness-design.md`
- Create: `docs/superpowers/plans/2026-08-01-lore-adoption-readiness.md`
- Modify: `docs/README.md`

**Interfaces:**
- Consumes: Task schema version 1 and the existing repository documentation structure.
- Produces: A repeatable maintenance request and a human activation runbook.

- [ ] **Step 1: Add the task artifact**

Use task ID `bootstrap-repository-knowledge`, list the evidence-bearing repository paths, request architecture and maintainer audiences, and set `history: false`.

- [ ] **Step 2: Document the release gate**

State the exact `0.1.0` installation, extraction, validation, proposal, apply, project, and check commands. Explain why README projection and CI activation are deferred.

- [ ] **Step 3: Link the readiness guide**

Add the guide to `docs/README.md` under Explanation.

- [ ] **Step 4: Commit the readiness packet**

```bash
git add .lore/tasks docs
git commit -m "docs: define LORE adoption readiness"
```

### Task 3: Produce the initial semantic proposal

**Files:**
- Create: `.lore/proposals/bootstrap-repository-knowledge.yaml`

**Interfaces:**
- Consumes: The exact trust-root commit, committed maintainer skill, task artifact, and repository evidence.
- Produces: One schema-conforming `lore-proposal/v1` artifact with append-only initial records.

- [ ] **Step 1: Use the committed skill**

Read `skills/maintain-repository-documentation/SKILL.md` from the exact proposal base revision and use its SHA-256 digest.

- [ ] **Step 2: Propose initial records**

Include repository, component, relationship, decision, constraint, and procedure records. Use revision 1, `supersedes: null`, repository scope `building-code-map`, and evidence paths that exist at the base revision.

- [ ] **Step 3: Validate the proposal structure**

Confirm one `changes_proposed` result, at least one operation, no duplicate record identities, valid component references, and no uncertainty hidden as certainty.

- [ ] **Step 4: Commit the proposal separately**

```bash
git add .lore/proposals/bootstrap-repository-knowledge.yaml
git commit -m "docs: stage initial LORE knowledge proposal"
```

The proposal must remain unapplied until the released CLI validates it.

### Task 4: Activate after the stable package is published

**Files:**
- Modify after release: `package.json`
- Modify after release: `pnpm-lock.yaml`
- Modify after release: `.github/workflows/ci.yml`
- Generate after release: `.lore/extracted/*.yaml`
- Generate after release: `.lore/records/**`
- Generate after release: `.lore/transactions/*.yaml`
- Generate after release: `docs/generated/lore/*.md`

**Interfaces:**
- Consumes: Published `@laurajoyhutchins/lore@0.1.0` and the staged proposal.
- Produces: Validated accepted history, deterministic projections, and CI drift enforcement.

- [ ] **Step 1: Install the exact stable version**

```bash
pnpm add -D @laurajoyhutchins/lore@0.1.0
pnpm exec lore version --json
```

Require package name `@laurajoyhutchins/lore`, version `0.1.0`, Node.js 22 or newer, and schema version 1 for manifest, record, proposal, task, hydration, and transaction.

- [ ] **Step 2: Generate deterministic facts and validate**

```bash
pnpm exec lore extract
pnpm exec lore validate
```

- [ ] **Step 3: Validate and apply the proposal**

```bash
pnpm exec lore validate-proposal .lore/proposals/bootstrap-repository-knowledge.yaml
pnpm exec lore apply .lore/proposals/bootstrap-repository-knowledge.yaml
```

- [ ] **Step 4: Generate and check projections**

```bash
pnpm exec lore project
pnpm exec lore validate
pnpm exec lore extract --check
pnpm exec lore project --check
```

- [ ] **Step 5: Add CI enforcement**

Add release-pinned install plus `lore extract --check`, `lore validate`, and `lore project --check` to the read-only verification workflow.

- [ ] **Step 6: Run the full repository verification**

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
