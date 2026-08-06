# Building Code Authority Engine Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Deliver a deterministic, provenance-bearing Building Code Authority Engine as seven stacked draft pull requests descended temporarily from exact predecessor PR #48.

**Architecture:** Preserve the existing snapshot, geocoder, and regulatory implementations while moving orchestration into `backend/engine`. Promote reusable packages and make HTTP, CLI, MCP, and the website adapters of that engine. Compose the aggregate bundle manifest from the existing PR #45 component identities.

**Tech Stack:** Go 1.25 backend, standard-library HTTP/JSON/CLI/MCP framing, modernc SQLite, React/TypeScript/Vite frontend, JSON Schema, GitHub stacked branches, offline file verification.

## Global Constraints

- Preserve offline-first operation, deterministic execution, exact source and bundle identity, structured provenance, stable ordering, typed errors, and no network dependency during query execution.
- Treat PR #45, PR #47, and PR #48 as predecessors and do not recreate their functionality.
- Do not merge PRs during this implementation session.
- Keep PR #14, PR #35, the modernc SQLite upgrade, unrelated Dependabot PRs, and Issue #7’s branch separate unless a concrete conflict requires a documented compatibility change.
- Require explicit `--as-of` or `applicability_date` for new resolution contracts; retain legacy omitted-date behavior only in compatibility routes.
- Preserve geocoder ranking and interpolation provenance from PR #48 through engine, adapters, and frontend evidence.
- Do not modify or repurpose LORE trust-root material, generated LORE projections, or Offline Execution.
- Use TDD for new behavior: write a failing test, run it, implement the smallest change, rerun focused and full verification.
- Record exact verified head SHAs and actual command output before committing, pushing, or claiming completion.

---

### Task 1: Contracts branch and compatibility baseline

**Branch:** `refactor/authority-engine-contracts`, based on exact PR #48 head `d7908d9c4fc55f34d27f204310acc5cdd03919e7`.

**Files:**
- Create: `backend/engine/contracts.go`, `backend/engine/errors.go`, `backend/engine/clock.go`, `backend/engine/contracts_test.go`
- Create: `backend/internal/httpapi/compatibility_test.go`
- Modify: `backend/internal/regulatory/resolver.go` only after a failing clock test demonstrates the seam needed for testability; preserve legacy behavior.
- Modify: `docs/reference/authority-engine.md` with the contract summary.

**Interfaces:**
- Produce `engine.Query`, `engine.Result`, `engine.EngineError`, `engine.ErrorCode`, `engine.Clock`, `engine.RealClock`, and `engine.FixedClock`.
- Preserve current `internal/httpapi` constructor and response shapes.

- [ ] Write compatibility tests for health, readiness, layers, boundaries, feature lookup, geocode success/failure, lookup success/failure, resolve success, malformed JSON, invalid coordinates, missing catalog, unsupported state, boundary ambiguity, refresh status, and disabled refresh trigger.
- [ ] Run `cd backend && go test ./engine ./internal/httpapi -count=1` and verify the new tests fail only where the contract types or fixtures are absent.
- [ ] Define stable JSON contracts, error codes, clock implementations, and deterministic sorting helpers with unit tests for normalization, date validation, and fixed time.
- [ ] Run the focused tests, then `cd backend && go test ./... -count=1`.
- [ ] Commit only the contracts, compatibility tests, documentation, and required test seam as `backend: define authority engine contracts`.
- [ ] Push the branch and open draft PR 1 targeting `feature/geocoder-provenance-ranking`; explain the temporary external base and non-mergeability into `main` until #45–#48 merge.

### Task 2: Core authority engine

**Branch:** `refactor/authority-engine-core`, based on the exact PR 1 head.

**Files:**
- Create: `backend/engine/engine.go`, `backend/engine/geography.go`, `backend/engine/provenance.go`, `backend/engine/engine_test.go`, `backend/engine/geography_test.go`
- Modify: `backend/internal/httpapi/resolve.go`, `backend/internal/httpapi/geocode.go`, `backend/internal/regulatory/resolver.go`, and `backend/cmd/server/main.go`

**Interfaces:**
- `engine.GeographyResolver.ResolveGeography(context.Context, engine.Point) (regulatory.GeographicContext, error)`.
- `engine.Config` contains snapshot, catalog, geocoder, clock, and bundle identity.
- `engine.Engine` exposes `Resolve`, `Geocode`, `Lookup`, and `Readiness`.

- [ ] Add failing geography tests for polygon holes, multipolygons, boundary edges, deterministic ordering, state/county/municipality/special/tribal/fire matches, unsupported coverage, and ambiguity observations.
- [ ] Add failing engine tests for point/address exclusivity, explicit applicability date, address geocoding, typed errors, complete provenance, deterministic diagnostics, and fixed-clock byte-equivalence.
- [ ] Move the geometry and context construction into the engine geography resolver with the smallest reusable extraction from HTTP.
- [ ] Thread an explicit date/clock through regulatory resolution and retain generated-at timestamps under engine control.
- [ ] Construct one engine instance in `cmd/server` and make legacy handlers delegate without changing legacy response JSON.
- [ ] Run focused engine/regulatory/http/server tests and the full Go suite; record any unavailable-tool blocker rather than claiming success.
- [ ] Commit and open draft PR 2 targeting the exact PR 1 head.

### Task 3: Public packages and HTTP transport

**Branch:** `refactor/authority-engine-transports`, based on exact PR 2 head.

**Files:**
- Move: `backend/internal/geocoder/` to `backend/geocoder/` incrementally.
- Move: `backend/internal/regulatory/` to `backend/regulatory/` incrementally.
- Move: `backend/internal/snapshot/` to `backend/snapshot/` incrementally.
- Create: `backend/transport/httpapi/` and its tests.
- Modify: import sites and `backend/cmd/server/main.go`.

- [ ] Add import-boundary tests or compile checks before each package move.
- [ ] Move one package, update imports, run `cd backend && go test ./... -count=1`, and commit only after the suite is green.
- [ ] Add failing v1 route tests for `/v1/geocode`, `/v1/lookup`, `/v1/resolve`, `/v1/readiness`, and `/v1/bundle` plus strict JSON and typed error mapping.
- [ ] Implement `NewHandler(engine.Engine, Options) http.Handler`; keep CORS, routing, decoding, status mapping, and legacy shaping in transport only.
- [ ] Run `rg "geometryContainsPoint|regulatory\.Resolve|BoundaryFeatures" backend/transport` and manually confirm no production transport domain decisions remain.
- [ ] Run `cd backend && go test ./transport/httpapi ./engine ./... -count=1 && go vet ./...`.
- [ ] Commit and open draft PR 3 targeting exact PR 2.

### Task 4: CLI, schemas, and aggregate bundle

**Branch:** `feat/bcm-cli-bundle`, based on exact PR 3 head.

**Files:**
- Create: `backend/cmd/bcm/` command files and tests.
- Create: `schemas/engine-query.schema.json`, `schemas/engine-result.schema.json`, `schemas/engine-error.schema.json`, `schemas/engine-bundle.schema.json`.
- Create: `backend/bundle/manifest.go`, `backend/bundle/verify.go`, and tests.
- Modify: `backend/cmd/server` into a compatibility wrapper around the shared serving composition.

- [ ] Add failing CLI tests for explicit `--as-of`, compact/pretty JSON, stderr-only logs, all command forms, and exit codes 0–4.
- [ ] Add failing bundle tests for missing components, digest mismatch, traversal, duplicate roles, unsupported schema, and metadata inconsistency.
- [ ] Implement shared composition and CLI commands without duplicating engine decisions.
- [ ] Implement aggregate manifest validation and deterministic canonical serialization while composing PR #45 component manifests.
- [ ] Validate actual Go JSON output against all JSON Schemas.
- [ ] Run CLI help, bundle inspection, schema validation, and the full Go suite.
- [ ] Commit and open draft PR 4 targeting exact PR 3.

### Task 5: MCP stdio adapter

**Branch:** `feat/bcm-mcp`, based on exact PR 4 head.

**Files:**
- Create: `backend/transport/mcp/stdio.go`, `backend/transport/mcp/stdio_test.go`.
- Modify: `backend/cmd/bcm` to add `serve --mcp-stdio`.

- [ ] Add failing protocol tests for initialize, tool discovery, success, structured engine failure, malformed request, cancellation, and stdout/stderr separation.
- [ ] Implement exactly five bounded tools mapped to engine methods/read operations, with schemas from engine contracts and no prompts, memory, planning, or fallback logic.
- [ ] Run a real fixture-backed stdio initialization and tool call, then `cd backend && go test ./transport/mcp ./cmd/bcm ./... -count=1`.
- [ ] Commit and open draft PR 5 targeting exact PR 4.

### Task 6: Website v1 migration

**Branch:** `refactor/website-engine-v1`, based on exact PR 5 head.

**Files:**
- Modify: `src/lib/api.ts`, `src/lib/apiPayloads.ts`, `src/lib/boundaryPayloads.ts`, `src/lib/runtimeDecode.ts`, `src/types.ts`, and lookup components.
- Create or modify: typed v1 client tests and presentation-model tests.
- Modify: `README.md`, `AGENTS.md`, and hand-authored architecture/reference documentation.

- [ ] Add failing decoder tests for engine result/error schemas, geocoder provenance preservation, readiness identities, unknown optional fields, and cancellation.
- [ ] Implement v1 client calls by extending existing runtime decoders and preserving structured `ApiResponseError` semantics.
- [ ] Map engine results to user-facing presentation models without rendering raw diagnostics as unexplained copy.
- [ ] Preserve legacy routes and Issue #7 map/detail cancellation/cache behavior.
- [ ] Run `pnpm check`, `pnpm lint`, `pnpm test`, `pnpm build`, `pnpm backend:lint`, and `pnpm backend:test` when toolchains are available.
- [ ] Commit and open draft PR 6 targeting exact PR 5.

### Task 7: Offline release and cleanup

**Branch:** `build/authority-engine-offline-release`, based on exact PR 6 head.

**Files:**
- Create: `building-code-engine/bin/bcm`, `building-code-engine/data/`, `building-code-engine/manifests/`, `building-code-engine/schemas/`, `building-code-engine/licenses/`, `building-code-engine/README.md` according to repository conventions.
- Create: cold-room verification scripts and exact-head evidence documentation.
- Modify: temporary wrappers and documentation only after all consumers use the engine.

- [ ] Add failing cold-room tests for missing bundle, wrong digest, offline startup, point resolution, address resolution, HTTP readiness/v1 resolution, MCP initialization, and one MCP tool call.
- [ ] Build content-addressed release artifacts and record source, binary, bundle, and component digests.
- [ ] Run the complete root/backend/regulatory suite and the cold-room workflow with network access disabled.
- [ ] Keep legacy routes for the documented compatibility window and remove `cmd/server` only if every script and document uses `bcm serve --http`.
- [ ] Commit and open draft PR 7 targeting exact PR 6; do not modify Offline Execution.
- [ ] After predecessor merges, transplant/rebase all seven branches onto updated `main`, retarget each PR, and rerun exact-head verification in order.

