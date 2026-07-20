# Configuration Reference

This page records which files own which kinds of configuration so the repository stays easy to maintain.

## Frontend Tooling

- [`package.json`](/package.json): canonical package scripts and frontend dependency declarations.
- [`pnpm-lock.yaml`](/pnpm-lock.yaml): reproducible JavaScript dependency resolution.
- [`eslint.config.js`](/eslint.config.js): ESLint rules for `src/`.
- [`.prettierrc.json`](/.prettierrc.json): Prettier formatting defaults.
- [`.prettierignore`](/.prettierignore): files Prettier skips.
- [`.editorconfig`](/.editorconfig): editor-wide newline and indentation defaults.
- [`tsconfig.json`](/tsconfig.json): application TypeScript settings.

## Backend Tooling

- [`backend/go.mod`](/backend/go.mod): Go module metadata and direct dependencies.
- [`backend/go.sum`](/backend/go.sum): reproducible Go dependency checksums.
- [`backend/cmd/server/main.go`](/backend/cmd/server/main.go): Go backend entrypoint and runtime flag handling.
- [`backend/internal/httpapi/handler.go`](/backend/internal/httpapi/handler.go): HTTP routing, CORS handling, and JSON response helpers.
- [`backend/internal/snapshot/snapshot.go`](/backend/internal/snapshot/snapshot.go): SQLite snapshot loading and cache-path resolution, with DuckDB retained only as a legacy migration fallback.

## Backend Snapshot

- `TIGERWEB_CACHE_PATH`: optional snapshot path inside the repository checkout. If unset, the backend prefers `backend/data/tigerweb.sqlite` and falls back to `backend/data/tigerweb.duckdb` during migration.
- `TIGERWEB_HYDRATED_CACHE_PATH`: optional hydrated-snapshot override inside the repository checkout. Paths resolving outside the checkout are rejected.
- `BACKEND_CORS_ALLOWED_ORIGINS`: optional comma-separated browser origins accepted by the backend. Invalid or empty values fall back to loopback-only defaults.
- `DUCKDB_EXE`: optional path to the DuckDB CLI used by the legacy loader. `DUCKDB_CLI_PATH` remains a compatibility alias.
- `backend/data/`: conventional local runtime-snapshot location. Database files in this directory are ignored and must be generated or supplied locally; they are not publication artifacts.

## Operational Scripts

- [`tools/start.cmd`](/tools/start.cmd): Windows launcher for the frontend and backend services.
- [`tools/health.cmd`](/tools/health.cmd): HTTP health checks for local services.
- [`tools/stop.cmd`](/tools/stop.cmd): stops services recorded by the launcher.
- [`tools/deciduous.cmd`](/tools/deciduous.cmd): Windows wrapper for the Deciduous CLI.
- [`tools/deciduous-work.cmd`](/tools/deciduous-work.cmd): starts a goal from the supplied user request, then shows recovery state.
- [`tools/deciduous-recover.cmd`](/tools/deciduous-recover.cmd): session-start recovery for the current graph.
- [`tools/deciduous-audit.cmd`](/tools/deciduous-audit.cmd): graph-hygiene audit for goals, decisions, actions, and outcomes.
- [`tools/start.ps1`](/tools/start.ps1), [`tools/health.ps1`](/tools/health.ps1), and [`tools/stop.ps1`](/tools/stop.ps1): PowerShell implementations behind the service wrappers.
- [`tools/deciduous.ps1`](/tools/deciduous.ps1), [`tools/deciduous-work.ps1`](/tools/deciduous-work.ps1), [`tools/deciduous-recover.ps1`](/tools/deciduous-recover.ps1), and [`tools/deciduous-audit.ps1`](/tools/deciduous-audit.ps1): PowerShell implementations behind the decision-workflow wrappers.
- [`tools/dev.mjs`](/tools/dev.mjs): loopback-only Vite development server with the `/api` proxy.
- [`tools/preview.mjs`](/tools/preview.mjs): loopback-only Vite preview server with the same proxy behavior.
- [`reports/tools/generate_report_tracker.py`](/reports/tools/generate_report_tracker.py): report-set dashboard generator for completion, issues, source coverage, and consistency checks.
- [`reports/tools/reports-status.cmd`](/reports/tools/reports-status.cmd): Windows wrapper for report status.

The backend launcher resolves `go` from `PATH`, runs `go run ./cmd/server`, and waits on `GET /ready` before reporting success. The frontend launcher and package scripts use the same `/api` proxy path.

## Generated or Ignored State

- `dist/`: generated frontend build output.
- `node_modules/`: installed JavaScript dependencies.
- `.venv/`: local Python environment used by report tooling when needed.
- `tools/.state/` and `tools/.logs/`: launcher state and logs.
- `cache/`: local downloads, hydrated snapshots, and tool caches.
- `.deciduous/`: local decision-graph database.
- `docs/graph-data.json` and `docs/git-history.json`: optional generated decision-graph exports. These are ignored because they can contain verbatim prompts, commit metadata, and stale implementation state.
- `*.sqlite`, `*.duckdb`, and their journal or temporary files: local runtime data, never source-controlled.

## Canonical Commands

- `pnpm install --frozen-lockfile`: install locked frontend dependencies.
- `pnpm check`: run TypeScript checks.
- `pnpm lint`: run ESLint.
- `pnpm format:check`: verify formatting.
- `pnpm test`: run frontend tests.
- `pnpm build`: build the frontend.
- `pnpm decision:recover`: recover current Deciduous graph state.
- `pnpm decision:audit`: summarize graph hygiene.
- `pnpm decision:audit:strict`: fail on hard graph-hygiene gaps.
- `cd backend; go test ./...`: run backend tests.
- `cd backend; go vet ./...`: run backend static checks.
- `.\tools\start.cmd`, `.\tools\health.cmd`, and `.\tools\stop.cmd`: operate the local service pair on Windows.

## Cleanup Rules

- Keep `.editorconfig` and Prettier aligned on indentation and line endings.
- Keep frontend rules in ESLint and backend checks in Go tooling.
- Prefer adding a package-script alias over duplicating a command in multiple documents.
- Keep secrets, private prompts, absolute home-directory paths, generated decision exports, and hydrated data out of Git history.
- Treat state reports as evidence-bearing research artifacts: retain official-source provenance and label unresolved claims explicitly.
