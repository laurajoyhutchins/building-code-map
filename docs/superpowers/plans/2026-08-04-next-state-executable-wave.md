# Next-State Executable Wave Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Promote Virginia, Oregon, and North Carolina from research-only draft reports into conservative, source-backed executable regulatory profiles.

**Architecture:** Preserve the existing Markdown-to-compiled-JSON pipeline. Each state report remains the human-reviewable source artifact, each rule pack carries executable applicability/date/amendment/enforcement claims, and the Go resolver consumes only checked-in compiled artifacts. The wave adds no state-name branches and does not extend the resolver for municipality-conditioned policies, geographic exceptions, or occupancy routing.

**Tech Stack:** Markdown reports with embedded JSON, Python regulatory compiler and validator, Go regulatory catalog and resolver, GitHub Actions.

## Global Constraints

- Do not claim complete statewide, municipal, amendment, code-family, project-type, or historical coverage.
- Preserve unknown dates and condition-triggered future transitions without inventing calendar dates.
- Keep Massachusetts, New York, and California outside executable pilot discovery until their resolver prerequisites exist.
- Use official state sources and record source-health dates.
- Keep the generic resolver free of state-name branches.
- A pending adoption must not be returned as current merely because it exists in a profile.

---

### Task 1: Define resolver acceptance behavior

**Files:**
- Create: `backend/internal/regulatory/next_states_test.go`

**Interfaces:**
- Consumes: `LoadCatalog`, `Resolve`, `ResolutionRequest`, and the existing resolver test helpers.
- Produces: regression expectations for Virginia, Oregon, and North Carolina.

- [ ] Add a Virginia test that expects the 2021 Virginia Construction Code, statewide adopting authority, local enforcement candidate, and one-year transition rule.
- [ ] Add an Oregon test that expects the 2025 OSSC after its mandatory date and preserves the 2022/2025 phase-in rule for a phase-in date.
- [ ] Add a North Carolina test that returns the current 2018 code while excluding the conditionally pending 2024 code and surfacing the certification record requirement.
- [ ] Push the test-only commit and confirm CI fails because the three compiled profiles are absent.

### Task 2: Promote the three state reports

**Files:**
- Create: `reports/pilots/virginia.md`
- Create: `reports/pilots/oregon.md`
- Create: `reports/pilots/north-carolina.md`
- Delete: `reports/virginia.md`
- Delete: `reports/oregon.md`
- Delete: `reports/north-carolina.md`

**Interfaces:**
- Consumes: jurisdiction-profile schema 1.0 and existing pilot-report front matter.
- Produces: partially verified, compiler-discoverable state reports.

- [ ] Recheck all official source pages as of 2026-08-04.
- [ ] Change report and profile status from `draft` to `partially_verified` while retaining scope and uncertainty flags.
- [ ] Preserve only source-backed adoption and authority paths.
- [ ] Keep the North Carolina 2024 code as `pending` with no operative date.
- [ ] Keep the Oregon 2026 residential code as `pending` with no operative date.

### Task 3: Add executable rule packs and compiled profiles

**Files:**
- Create: `backend/data/regulatory/virginia.json`
- Create: `backend/data/regulatory/oregon.json`
- Create: `backend/data/regulatory/north-carolina.json`
- Create: `backend/data/regulatory/rules/virginia.json`
- Create: `backend/data/regulatory/rules/oregon.json`
- Create: `backend/data/regulatory/rules/north-carolina.json`

**Interfaces:**
- Consumes: embedded profile JSON from the three reports.
- Produces: deterministic catalog inputs for the Go resolver.

- [ ] Compile each report into canonical indented JSON.
- [ ] Add source-health records for every cited source.
- [ ] Encode Virginia's one-year prior-edition option as a date rule without supplying an unsupported prior-edition adoption.
- [ ] Encode Oregon's code-family-specific optional and mandatory dates.
- [ ] Encode North Carolina's current 2018 baseline and condition-based 2024 transition as a supported claim, not a calendar date.
- [ ] Add amendment and enforcement rules that preserve required local records.

### Task 4: Reconcile research and public status documentation

**Files:**
- Modify: `docs/research/2026-08-02-next-state-wave.md`
- Modify: `README.md`

**Interfaces:**
- Produces: documentation that distinguishes six executable pilots from production or nationwide completeness.

- [ ] Mark Virginia, Oregon, and North Carolina as the implemented executable wave.
- [ ] Keep Massachusetts, New York, and California as blocked research tracks with their exact resolver prerequisites.
- [ ] Name all six executable pilot states in the README while retaining the pre-1.0 and local-confirmation warnings.

### Task 5: Verify exact head

**Files:**
- No new files.

- [ ] Run Python compiler unit tests.
- [ ] Validate every pilot report.
- [ ] Check compiled regulatory artifacts for drift.
- [ ] Generate the nationwide quality report.
- [ ] Run all Go tests and `go vet ./...`.
- [ ] Run frontend type checking, linting, formatting, tests, and build.
- [ ] Run LORE validation through the repository CI workflow.
- [ ] Inspect the exact-head diff and update PR #39 with scope, limitations, and verification evidence.
