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
- [`backend/cmd/geocoder-build/main.go`](/backend/cmd/geocoder-build/main.go): deterministic CSV-to-SQLite geocoder snapshot builder.
- [`backend/internal/httpapi/handler.go`](/backend/internal/httpapi/handler.go): HTTP routing, CORS handling, and JSON response helpers.
- [`backend/internal/httpapi/readiness.go`](/backend/internal/httpapi/readiness.go): capability-specific readiness.
- [`backend/internal/geocoder/`](/backend/internal/geocoder/): address normalization, SQLite matching, interpolation, provenance, and snapshot construction.
- [`backend/internal/snapshot/snapshot.go`](/backend/internal/snapshot/snapshot.go): supported snapshot loaders and workspace-local cache selection.
- [`backend/internal/snapshot/validate.go`](/backend/internal/snapshot/validate.go): shared semantic validation for every boundary snapshot loader.

## Backend Snapshots

### Boundary snapshot

The only implicit default is:

```text
backend/data/tigerweb.sqlite
```

Configuration precedence is:

1. an explicit `--cache` server flag;
2. an existing repository-contained `TIGERWEB_CACHE_PATH` value;
3. an existing repository-contained `TIGERWEB_HYDRATED_CACHE_PATH` compatibility value;
4. `backend/data/tigerweb.sqlite`.

Environment overrides that resolve outside the checkout or do not exist are ignored. The server does not search `C:\tmp`, a home directory, or another machine-global cache location.

Supported formats are:

- `.sqlite`
- `.db`
- `.duckdb`, only when selected explicitly and a DuckDB CLI is available

Unknown extensions fail before compatibility-tool discovery. The default never falls back silently from missing SQLite data to a DuckDB file.

SQLite boundary snapshots are opened read-only and query-only. Every loader applies the same semantic validator before startup can proceed.

### Geocoder snapshot

- `GEOCODER_DATA_PATH`: optional local geocoder SQLite path inside the backend checkout. It overrides `--geocoder-data`.
- `--geocoder-data`: server flag for a repository-contained geocoder snapshot path.
- Default: `backend/data/geocoder.sqlite` when launched from `backend/`.

A missing or invalid geocoder snapshot leaves `/geocode` and address-based `/lookup` unavailable. Boundary inspection remains available. Coordinate-based regulatory resolution also requires validated regulatory profiles, as reported by `/ready`.

### Other runtime settings

- `BACKEND_CORS_ALLOWED_ORIGINS`: comma-separated browser origins accepted by the backend. Invalid or empty values fall back to loopback-only defaults.
- `DUCKDB_EXE`: optional DuckDB CLI path for an explicitly selected legacy boundary snapshot. `DUCKDB_CLI_PATH` remains a compatibility alias.
- `backend/data/`: conventional local runtime-snapshot location. Database files in this directory are ignored and must be generated or supplied locally; they are not publication artifacts.

## Operational Scripts

- [`tools/start.cmd`](/tools/start.cmd): Windows launcher for the frontend and backend services.
- [`tools/health.cmd`](/tools/health.cmd): HTTP health checks for local services.
- [`tools/stop.cmd`](/tools/stop.cmd): stops services recorded by the launcher.
- [`tools/configure-github-repository.cmd`](/tools/configure-github-repository.cmd): applies and verifies the documented GitHub merge, ruleset, and security settings through the authenticated GitHub CLI.
- [`tools/configure-github-repository.ps1`](/tools/configure-github-repository.ps1): idempotent PowerShell implementation behind the GitHub settings wrapper.
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

The backend launcher resolves `go` from `PATH`, runs `go run ./cmd/server`, and waits on `GET /ready`. A degraded but boundary-usable service returns HTTP `200`; required boundary unavailability returns `503`.

## Generated or Ignored State

- `dist/`: generated frontend build output.
- `node_modules/`: installed JavaScript dependencies.
- `.venv/`: local Python environment used by report tooling when needed.
- `tools/.state/` and `tools/.logs/`: launcher state and logs.
- `cache/`: local downloads, hydrated snapshots, and tool caches.
- `.deciduous/`: local decision-graph database.
- `.lore/extracted/`: deterministic LORE facts, regenerated by LORE rather than edited directly.
- `.lore/proposals/`: untrusted LORE proposal artifacts awaiting validation and application.
- `.lore/transactions/`: append-only LORE transaction receipts.
- `docs/generated/`: LORE projections. These are non-authoritative and regenerated from source, facts, and accepted semantic records.
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
- `cd backend; go run ./cmd/geocoder-build --help`: inspect the local geocoder snapshot builder.
- `.\tools\start.cmd`, `.\tools\health.cmd`, and `.\tools\stop.cmd`: operate the local service pair on Windows.
- `.\tools\configure-github-repository.cmd`: apply and verify public repository governance settings.

## Cleanup Rules

- Keep `.editorconfig` and Prettier aligned on indentation and line endings.
- Keep frontend rules in ESLint and backend checks in Go tooling.
- Prefer adding a package-script alias over duplicating a command in multiple documents.
- Keep secrets, private prompts, absolute home-directory paths, generated decision exports, and hydrated data out of Git history.
- Treat boundary, address, and regulatory sources as separate provenance domains with independently reviewed redistribution terms.
- Treat state reports as evidence-bearing research artifacts: retain official-source provenance and label unresolved claims explicitly.
- Do not describe refresh metadata as a complete source manifest, checksum, activation receipt, or rollback contract.
