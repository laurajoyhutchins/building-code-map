# State Jurisdiction Resolution Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Deliver deterministic, evidence-bearing location-to-authority resolution from accumulated state research.

**Architecture:** Compile evidence-rich pilot reports into validated declarative profiles, interpret them with a generic Go resolver, join the resolver to local snapshot geography through a versioned HTTP endpoint, and expose uncertainty and evidence in the React explorer. A separate quality command measures nationwide readiness without treating file presence as completion.

**Tech Stack:** Go 1.25, Python 3.13 standard library, React 18, TypeScript, Vitest, JSON Schema, GitHub Actions.

## Global Constraints

- Preserve evidence reports as the human-readable source and maintenance layer.
- Do not include or republish model-code or standards text.
- Do not infer local adopted-code editions from state-level data.
- State-specific behavior must remain declarative.
- Unknown, unresolved, ambiguous, conflicting, and local-record-required states must remain explicit.
- All generated artifacts must be deterministic and fail closed on malformed references.

---

### Task 1: Canonical regulatory model

**Files:**

- Create: `backend/internal/regulatory/model.go`
- Create: `backend/internal/regulatory/validate.go`
- Create: `backend/internal/regulatory/catalog.go`
- Create: `schemas/regulatory/state-profile.schema.json`
- Test: `backend/internal/regulatory/resolver_test.go`

**Interfaces:**

- Produces: `StateProfile`, `Catalog`, `LoadCatalog(string)`, and `ValidateProfile(StateProfile)`.

- [x] Write failing tests for invalid references and duplicate profile keys.
- [x] Verify the tests fail before the package exists.
- [x] Implement the versioned types, validation, deterministic loading, and schema.
- [x] Run `cd backend && go test ./internal/regulatory` and verify PASS.

### Task 2: Source-report compiler and pilot profiles

**Files:**

- Create: `reports/tools/regulatory.py`
- Create: `reports/tools/test_regulatory.py`
- Create: `reports/pilots/colorado.md`
- Create: `reports/pilots/florida.md`
- Create: `reports/pilots/new-jersey.md`
- Create: `backend/data/regulatory/colorado.json`
- Create: `backend/data/regulatory/florida.json`
- Create: `backend/data/regulatory/new-jersey.json`

**Interfaces:**

- Produces: `validate`, `compile`, and `coverage` CLI commands and canonical `StateProfile` JSON.

- [x] Write failing unit tests for malformed source/adoption references and deterministic output.
- [x] Verify failures before implementing the compiler.
- [x] Implement strict fenced-profile extraction, validation, compilation, and coverage reporting.
- [x] Encode three structurally distinct pilots from official sources.
- [x] Verify compiled output exactly matches checked-in JSON.

### Task 3: Generic resolver

**Files:**

- Create: `backend/internal/regulatory/resolver.go`
- Test: `backend/internal/regulatory/resolver_test.go`

**Interfaces:**

- Consumes: `Catalog`, `ResolutionRequest`, and normalized `GeographicContext`.
- Produces: `Resolve(Catalog, ResolutionRequest) (ResolutionResult, error)`.

- [x] Write failing tests for Colorado local records, Colorado electrical override, Florida statewide baseline, New Jersey fire authority, and unknown states.
- [x] Implement policy overlay and replacement semantics.
- [x] Verify deterministic results while ignoring the generated timestamp.
- [x] Run the package test suite and verify PASS.

### Task 4: Geography and HTTP API

**Files:**

- Create: `backend/internal/httpapi/resolve.go`
- Create: `backend/internal/httpapi/resolve_test.go`
- Modify: `backend/internal/httpapi/handler.go`
- Modify: `backend/cmd/server/main.go`
- Create: `schemas/regulatory/resolution-result.schema.json`

**Interfaces:**

- Produces: strict `POST /resolve` and runtime `--regulatory-data` configuration.

- [x] Write failing endpoint tests for point matching, unknown JSON fields, and polygon holes.
- [x] Implement Polygon/MultiPolygon matching and normalized geographic context.
- [x] Load the profile catalog at startup and fail the endpoint closed when unavailable.
- [x] Run `cd backend && go test ./...` and `go vet ./...`.

### Task 5: Explorer integration

**Files:**

- Create: `src/components/ResolutionPanel.tsx`
- Create: `src/resolution.css`
- Modify: `src/types.ts`
- Modify: `src/lib/api.ts`
- Modify: `src/lib/api.test.ts`
- Modify: `src/App.tsx`
- Modify: `src/main.tsx`

**Interfaces:**

- Produces: `fetchResolution`, `decodeResolutionResult`, and an independent resolution panel.

- [x] Write a failing response-decoder test.
- [x] Implement the typed API mapping and form.
- [x] Render certainty status, candidates, adoptions, required local records, warnings, and evidence.
- [ ] Run `pnpm check`, `pnpm lint`, `pnpm test`, and `pnpm build` in repository CI.

### Task 6: Verification, documentation, and publication

**Files:**

- Modify: `.github/workflows/ci.yml`
- Modify: `package.json`
- Create: `docs/reference/regulatory-resolution.md`
- Create: `docs/superpowers/specs/2026-07-20-state-jurisdiction-resolution-design.md`
- Create: `docs/superpowers/plans/2026-07-20-state-jurisdiction-resolution.md`

- [x] Add Python profile verification and deterministic compile checks to CI.
- [x] Document architecture, commands, API, runtime behavior, and limitations.
- [ ] Publish the implementation branch and open a pull request closing GitHub issues #24–#29.
- [ ] Verify all repository checks and merge only after the exact PR head is green.
- [ ] Update GitHub and Linear trackers with the delivered outcome.
