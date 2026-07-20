# Repository Guidelines

## Project Structure & Module Organization

- `src/` contains the Vite + React frontend. Use `src/components/` for UI, `src/lib/` for client-side logic, and `src/types.ts` for shared TypeScript types.
- `backend/` contains the Go service. The executable entrypoint is `backend/cmd/server`, with implementation packages under `backend/internal/` and tests beside the Go code.
- `reports/` contains source-backed state building-code research artifacts and report tooling.
- `docs/` follows Diátaxis: tutorials, how-to guides, reference, and explanation.
- `tools/` contains Windows launch, health, stop, report, and Deciduous helpers. Generated runtime state belongs under ignored directories such as `tools/.state/`, `tools/.logs/`, and `cache/`.

## Build, Test, and Development Commands

Run frontend commands from the repository root:

- `pnpm install --frozen-lockfile` installs the locked JavaScript dependencies.
- `pnpm dev` starts the frontend and local development proxy.
- `pnpm check` runs TypeScript without emitting files.
- `pnpm lint` runs ESLint on `src/`.
- `pnpm format:check` verifies Prettier formatting.
- `pnpm test` runs the Vitest suite once.
- `pnpm build` creates the production frontend bundle.

Run backend commands from `backend/`:

- `go run ./cmd/server --addr 127.0.0.1:8000` starts the API.
- `go test ./...` runs backend tests.
- `go vet ./...` runs Go static checks.
- `go fmt ./...` formats Go source.

On Windows, `tools/start.cmd`, `tools/health.cmd`, and `tools/stop.cmd` operate both services together.

## Coding Style & Naming Conventions

- TypeScript uses 2-space indentation, semicolons, double quotes, PascalCase components, and camelCase functions, variables, and hooks.
- Keep frontend tests next to their implementation with the `.test.ts` or `.test.tsx` suffix.
- Go code must remain `gofmt`-formatted. Keep HTTP concerns in `backend/internal/httpapi`, snapshot loading in `backend/internal/snapshot`, and command wiring in `backend/cmd/`.
- `.editorconfig`, ESLint, Prettier, TypeScript, and Go tooling define the checked-in baseline. Prefer existing conventions over one-off formatting.

## Testing Guidelines

- Add or update Vitest coverage for frontend data transforms, API clients, and boundary-selection behavior.
- Add or update Go tests for HTTP responses, CORS behavior, snapshot parsing, and cache-path resolution.
- When a change crosses the frontend/backend boundary, verify both stacks and update the API reference when the contract changes.
- Do not commit runtime database snapshots, local environment files, logs, generated decision-graph exports, or test output.

## Commit & Pull Request Guidelines

- Use short conventional-style commit subjects such as `backend: validate snapshot timestamps` or `docs: clarify source provenance`.
- Keep each commit focused and use imperative wording.
- Pull requests should describe the change, link relevant issues or decision nodes, and state frontend, backend, data, and documentation impact.
- Include screenshots for visible UI changes and sample requests or responses for API changes.

## Configuration and Data Safety

- The frontend reads `VITE_API_BASE_URL` and defaults to `/api` through the local Vite proxy.
- `BACKEND_CORS_ALLOWED_ORIGINS` controls browser origins accepted by the Go API and defaults to loopback-only origins.
- `TIGERWEB_CACHE_PATH` and `TIGERWEB_HYDRATED_CACHE_PATH` may point only to files inside the repository checkout.
- Runtime SQLite and DuckDB files are intentionally ignored. Publish source and reproducible ingestion logic, not local hydrated databases.
- Treat NERIS-derived records conservatively: do not commit department-owned operational or incident data without confirmed redistribution rights.

## Decision Graph Workflow

Use the repository-local Deciduous wrappers for durable, repo-visible architectural decisions. Generated exports must not contain private conversation material or machine-specific paths.

The canonical flow is:

```text
goal -> options -> decision -> action -> outcome
```

- `tools/deciduous-work.cmd` starts a user-visible initiative.
- `tools/deciduous-recover.cmd` summarizes current graph state.
- `tools/deciduous-audit.cmd` reports graph hygiene gaps; `-Strict` fails on hard gaps.
- `pnpm decision:recover`, `pnpm decision:audit`, and `pnpm decision:audit:strict` provide package-script aliases.
- Root goals are the only valid orphan nodes. Link every option, decision, action, outcome, observation, and revisit to its parent immediately.
- Capture durable project choices and verified outcomes. Do not capture private chain-of-thought, credentials, personal paths, or unrelated conversation text.

See `docs/how-to/use-decision-graph-for-repo-work.md` and `docs/reference/configuration.md` for the maintained workflow.