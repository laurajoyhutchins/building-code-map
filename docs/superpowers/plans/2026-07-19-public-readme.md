# Public README Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Reorganize the public README so a first-time visitor understands the project immediately and a developer can run the current frontend/backend workflow within five minutes.

**Architecture:** Keep `README.md` as a concise landing page using progressive disclosure. Link to the existing Diátaxis documentation for detailed runtime, configuration, data, and workflow material rather than duplicating it.

**Tech Stack:** Markdown, pnpm 10.12.1, Node.js 22, Go 1.25, React, Vite, MapLibre GL JS, SQLite.

## Global Constraints

- Describe the delivered runtime as React/MapLibre frontend → Go API → local SQLite snapshot.
- State clearly that the repository is pre-1.0 and not an authoritative compliance determination.
- Do not present PostGIS, Martin, geocoding, historical versioning, or nationwide coverage as implemented.
- Do not commit runtime SQLite/DuckDB files or redistribution-sensitive datasets.
- Preserve links to `DATA_SOURCES.md`, `SECURITY.md`, `LICENSE`, and `NOTICE`.

---

### Task 1: Rewrite the README landing experience

**Files:**
- Modify: `README.md`

**Interfaces:**
- Consumes: scripts declared in `package.json`, Go version and module in `backend/go.mod`, runtime behavior documented in `docs/reference/runtime-and-api.md` and `docs/reference/configuration.md`.
- Produces: a public landing page with product framing, quick start, architecture summary, roadmap, and documentation links.

- [ ] **Step 1: Replace the opening with the product promise and honest maturity statement**

Use language that distinguishes the target deterministic authority/code lookup from the currently delivered boundary explorer and research foundation.

- [ ] **Step 2: Add the five-minute quick start**

Include exact commands:

```bash
pnpm install --frozen-lockfile
cd backend
go run ./cmd/server --addr 127.0.0.1:8000
```

and, in a second terminal:

```bash
pnpm dev
```

Document `http://127.0.0.1:8000/health`, `http://127.0.0.1:8000/ready`, and the frontend URL printed by Vite.

- [ ] **Step 3: Explain the missing-snapshot state**

State that database snapshots are intentionally excluded and that `/health` can succeed while `/ready` reports the runtime is not ready until a supported snapshot is supplied.

- [ ] **Step 4: Add the delivered architecture diagram and current/next/later roadmap**

Use the text diagram from the design spec and ensure deferred systems are listed only under later work.

- [ ] **Step 5: Keep detailed material linked rather than duplicated**

Link the tutorial, runtime/API reference, configuration reference, data-source policy, security policy, Apache license, and third-party notice.

- [ ] **Step 6: Run documentation formatting**

Run:

```bash
pnpm format:check
```

Expected: Prettier reports all matched files use consistent formatting.

- [ ] **Step 7: Run repository verification**

Run:

```bash
pnpm check
pnpm lint
pnpm test
pnpm build
pnpm backend:test
pnpm backend:lint
```

Expected: every command exits successfully with no test, type-check, lint, build, Go test, or Go vet failures.

- [ ] **Step 8: Commit**

```bash
git add README.md docs/superpowers/specs/2026-07-19-public-readme-design.md docs/superpowers/plans/2026-07-19-public-readme.md
git commit -m "docs: improve public README quick start"
```
