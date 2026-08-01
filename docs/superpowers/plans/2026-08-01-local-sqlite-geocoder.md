# Local SQLite Geocoder Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Add a deterministic local SQLite address geocoder that composes with the existing point-based regulatory resolver and is documented through a LORE-compatible trust root.

**Architecture:** A focused `backend/internal/geocoder` package owns normalization, indexed SQLite matching, candidate policy, provenance, and interpolation. HTTP endpoints expose geocoding independently and compose it with the unchanged regulatory resolver. The public frontend classifies coordinate versus address input and displays geocoder provenance. LORE schemas, skill, manifest, bootstrap records, and pinned validation CI make semantic documentation machine-checkable.

**Tech Stack:** Go 1.25, `modernc.org/sqlite`, React 18, TypeScript 5.5, Vitest 2, pnpm 10.12.1, GitHub Actions, LORE pinned to public commit `a8792bd594db34fcbf3065239b70ef4383f7e186`.

## Global Constraints

- Lookup-time geocoding must not require a network request.
- `POST /resolve` remains point-based and backward compatible.
- Exact address points outrank street-range interpolation.
- Ambiguity must be returned rather than silently broken.
- Every selected candidate must include source identity and source vintage.
- Hydrated SQLite databases must remain untracked.
- LORE generated documents, extracted facts, accepted records, and transactions must not be edited outside their prescribed workflows.

---

### Task 1: Define geocoder behavior with failing tests

**Files:**
- Create: `backend/internal/geocoder/normalize_test.go`
- Create: `backend/internal/geocoder/sqlite_test.go`
- Create: `backend/internal/geocoder/testdata/address_points.csv`
- Create: `backend/internal/geocoder/testdata/street_ranges.csv`

**Interfaces:**
- Produces: `ParseAddress(string) (ParsedAddress, error)`
- Produces: `Open(string) (*SQLiteService, error)`
- Produces: `(*SQLiteService).Geocode(context.Context, Query) (Result, error)`

- [ ] **Step 1: Write normalization tests** for canonical suffixes, directions, ZIP handling, whitespace, PO-box rejection, and incomplete civic addresses.
- [ ] **Step 2: Run `cd backend && go test ./internal/geocoder -run TestParseAddress -v`** and confirm compilation fails because the package API does not exist.
- [ ] **Step 3: Write SQLite behavior tests** that build a fixture snapshot and assert exact selection, ambiguity, not-found behavior, address-point preference, interpolation precision, and provenance.
- [ ] **Step 4: Run `cd backend && go test ./internal/geocoder -v`** and confirm failure is caused by the missing implementation.

### Task 2: Implement deterministic normalization and result contracts

**Files:**
- Create: `backend/internal/geocoder/model.go`
- Create: `backend/internal/geocoder/normalize.go`

**Interfaces:**
- Produces: `Status`, `Precision`, `Query`, `ParsedAddress`, `Candidate`, `Result`, `Service`
- Produces: `ParseAddress(string) (ParsedAddress, error)`

- [ ] **Step 1: Implement the minimal data contracts** required by the tests, including JSON field names used by HTTP responses.
- [ ] **Step 2: Implement bounded deterministic normalization** with explicit suffix and direction maps.
- [ ] **Step 3: Run `cd backend && go test ./internal/geocoder -run TestParseAddress -v`** and confirm normalization tests pass.
- [ ] **Step 4: Run `cd backend && gofmt -w internal/geocoder/model.go internal/geocoder/normalize.go`**.

### Task 3: Implement the SQLite snapshot and matcher

**Files:**
- Create: `backend/internal/geocoder/schema.go`
- Create: `backend/internal/geocoder/sqlite.go`
- Create: `backend/internal/geocoder/build.go`

**Interfaces:**
- Consumes: `ParsedAddress`, `Candidate`, `Result`
- Produces: `Open(string) (*SQLiteService, error)`
- Produces: `BuildSnapshot(BuildOptions) error`

- [ ] **Step 1: Implement the SQLite schema** for metadata, address points, street ranges, and indexes.
- [ ] **Step 2: Implement parameterized address-point candidate queries** and deterministic ranking.
- [ ] **Step 3: Implement ambiguity thresholds** using minimum score and runner-up gap constants.
- [ ] **Step 4: Implement parity-aware street-range interpolation** as a fallback only.
- [ ] **Step 5: Implement transactional CSV snapshot building** with validation, duplicate-source rejection, and atomic replacement.
- [ ] **Step 6: Run `cd backend && go test ./internal/geocoder -v`** and confirm all geocoder tests pass.
- [ ] **Step 7: Run `cd backend && go vet ./internal/geocoder`**.

### Task 4: Add a reproducible snapshot builder command

**Files:**
- Create: `backend/cmd/geocoder-build/main.go`

**Interfaces:**
- Consumes: `geocoder.BuildSnapshot(BuildOptions)`
- Produces: CLI flags `--output`, `--address-points`, `--street-ranges`, `--source-name`, `--source-vintage`

- [ ] **Step 1: Write command-level argument tests or extract a testable `run` function**.
- [ ] **Step 2: Run the focused test and verify it fails before implementation**.
- [ ] **Step 3: Implement the command with explicit required flags and nonzero exit on validation failure**.
- [ ] **Step 4: Run `cd backend && go test ./cmd/geocoder-build ./internal/geocoder`**.

### Task 5: Add geocoding and address lookup HTTP contracts

**Files:**
- Create: `backend/internal/httpapi/geocode_test.go`
- Create: `backend/internal/httpapi/geocode.go`
- Modify: `backend/internal/httpapi/handler.go`
- Modify: `backend/internal/httpapi/resolve.go`

**Interfaces:**
- Adds: `Options.Geocoder geocoder.Service`
- Adds: `POST /geocode`
- Adds: `POST /lookup`
- Preserves: `POST /resolve`

- [ ] **Step 1: Write failing HTTP tests** for unavailable service, exact geocode, ambiguous geocode, unknown fields, address-to-resolution composition, and coordinate resolve remaining available.
- [ ] **Step 2: Run `cd backend && go test ./internal/httpapi -run 'TestGeocode|TestLookup' -v`** and confirm the routes are absent.
- [ ] **Step 3: Add the optional geocoder dependency to `Handler` and route dispatch**.
- [ ] **Step 4: Implement strict request decoding with the existing body-size and trailing-value protections**.
- [ ] **Step 5: Refactor the point-to-regulatory call into a shared helper used by `/resolve` and `/lookup`** without changing the `/resolve` response.
- [ ] **Step 6: Map geocoder statuses to `200`, `409`, and `422`, and missing service to `503`**.
- [ ] **Step 7: Run `cd backend && go test ./internal/httpapi -v`**.

### Task 6: Wire server configuration without breaking coordinate lookup

**Files:**
- Modify: `backend/cmd/server/main.go`
- Modify: `docs/reference/configuration.md`

**Interfaces:**
- Adds: `--geocoder-data`
- Adds: `GEOCODER_DATA_PATH`

- [ ] **Step 1: Write or extend server configuration tests where the repository pattern permits**.
- [ ] **Step 2: Load the geocoder snapshot only when a configured file exists**; on failure log a warning and leave address endpoints unavailable.
- [ ] **Step 3: Pass the service through `httpapi.Options` and include geocoder availability in startup logging**.
- [ ] **Step 4: Document path containment, defaults, and disabled behavior**.
- [ ] **Step 5: Run `cd backend && go test ./... && go vet ./...`**.

### Task 7: Add frontend address-or-coordinate lookup

**Files:**
- Modify: `src/types.ts`
- Modify: `src/lib/api.ts`
- Modify: `src/lib/api.test.ts`
- Modify: `src/lib/publicLookup.ts`
- Modify: `src/lib/publicLookup.test.ts`
- Modify: `src/components/PublicLookup.tsx`
- Modify: `src/styles.css`

**Interfaces:**
- Adds: `GeocodeCandidate`, `GeocodeResult`, `LookupResult`
- Adds: `fetchLookup(input)`
- Adds: `classifyLocationQuery(value)`

- [ ] **Step 1: Write failing tests** for coordinate classification, address classification, lookup decoding, and geocoder provenance formatting.
- [ ] **Step 2: Run `pnpm test -- src/lib/publicLookup.test.ts src/lib/api.test.ts`** and confirm failures are caused by missing APIs.
- [ ] **Step 3: Implement the frontend types and decoder** without weakening the existing resolution decoder.
- [ ] **Step 4: Update the form to call `/resolve` for coordinates and `/lookup` for addresses**.
- [ ] **Step 5: Render matched address, precision, source, and vintage in a compact result detail**.
- [ ] **Step 6: Keep coordinate-only behavior functional when geocoding is unavailable**.
- [ ] **Step 7: Run `pnpm check && pnpm lint && pnpm format:check && pnpm test && pnpm build`**.

### Task 8: Document ingestion and API usage

**Files:**
- Create: `docs/how-to/build-local-geocoder-snapshot.md`
- Modify: `docs/reference/runtime-and-api.md`
- Modify: `DATA_SOURCES.md`
- Modify: `README.md`

**Interfaces:**
- Documents: CSV input columns, provenance requirements, builder command, endpoint examples, status semantics, and redistribution boundary.

- [ ] **Step 1: Document the snapshot builder using committed fixture-shaped examples, not a fictional nationwide source**.
- [ ] **Step 2: Document `/geocode` and `/lookup` request and response contracts**.
- [ ] **Step 3: State that address-source licensing and redistribution must be reviewed independently of Census boundary data**.
- [ ] **Step 4: Update the README delivered-feature and roadmap sections without claiming nationwide data coverage**.
- [ ] **Step 5: Run formatting checks**.

### Task 9: Initialize and validate the LORE trust root

**Files:**
- Create: `lore.yaml`
- Create: `BOOTSTRAP.md`
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
- Create: `.lore/records/repository.building-code-map.1.yaml`
- Create: `.lore/records/component.local-sqlite-geocoder.1.yaml`
- Create: `.lore/records/decision.point-based-resolution-boundary.1.yaml`
- Create: `.lore/records/constraint.local-geocoding-no-network.1.yaml`
- Create: `.lore/records/procedure.build-geocoder-snapshot.1.yaml`
- Modify: `.github/workflows/ci.yml`

**Interfaces:**
- Uses: LORE public commit `a8792bd594db34fcbf3065239b70ef4383f7e186`
- Produces: schema-valid append-only bootstrap records with exact commit evidence

- [ ] **Step 1: Copy the exact LORE initialization templates from the pinned revision**.
- [ ] **Step 2: Configure supported extractors and projections under `docs/generated/` without making the hand-authored README generated**.
- [ ] **Step 3: Add bootstrap records only after implementation evidence exists at a full commit SHA**.
- [ ] **Step 4: Add a CI job that checks out and builds the pinned LORE revision in the runner temp directory, then runs `validate` against this repository**.
- [ ] **Step 5: Run or observe CI and correct every schema, evidence, path, and invariant failure**.

### Task 10: Complete exact-head verification

**Files:**
- Review all changed files

- [ ] **Step 1: Run or confirm frontend checks at the exact branch head**.
- [ ] **Step 2: Run or confirm backend tests and vet at the exact branch head**.
- [ ] **Step 3: Confirm the LORE validation job passes at the same head**.
- [ ] **Step 4: Review the final diff for committed databases, private paths, invented provenance, generated-doc edits, and unrelated refactoring**.
- [ ] **Step 5: Open a draft pull request with API, data, frontend, and documentation impact summarized**.
