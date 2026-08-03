# Authority and Snapshot Hardening Implementation Plan

**Goal:** Make Building Code Map preserve the geocoder, boundary-observation, and regulatory-policy authority boundaries at runtime, reject semantically invalid snapshots before serving them, and report partial capability honestly.

**Architecture:** Keep the existing three-stage pipeline. Public callers provide a point or address, the backend derives geographic observations from the loaded boundary snapshot, and only the internal regulatory resolver consumes normalized geographic context. Snapshot loaders validate one shared semantic contract before constructing the HTTP handler. Readiness reports each independently usable capability without turning optional geocoding or regulatory data into process liveness.

**Constraints:**

- Modify only `laurajoyhutchins/building-code-map`.
- Preserve local-first lookup and repository-contained runtime paths.
- Keep hydrated SQLite and DuckDB files outside Git.
- Do not infer legal authority from polygon containment.
- Preserve overlapping geographic observations and fail closed when a single authority-bearing context cannot be selected safely.
- Keep legacy DuckDB readable only through explicit configuration.
- Do not broaden pilot-state regulatory coverage or merge unreviewed research from other branches.
- Add behavioral tests before or with each implementation slice and use existing Go, Vitest, Python, and LORE verification lanes.

**Verification:**

- `pnpm install --frozen-lockfile`
- `pnpm check`
- `pnpm lint`
- `pnpm format:check`
- `pnpm test`
- `pnpm build`
- `(cd backend && go test ./... && go vet ./...)`
- `python -m unittest discover -s reports/tools -p test_regulatory.py -v`
- `python reports/tools/regulatory.py validate reports/pilots/*.md`
- `python reports/tools/regulatory.py compile reports/pilots/*.md --output-dir backend/data/regulatory --check`
- pinned LORE extraction and validation through GitHub Actions

## Task 1: Validate and resolve boundary snapshots conservatively

**Outcome:** Every SQLite or explicitly selected DuckDB snapshot satisfies one semantic contract before it reaches the API, and implicit snapshot selection never escapes the repository or guesses an unknown format.

**Files:**

- Create `backend/internal/snapshot/validate.go`.
- Modify `backend/internal/snapshot/snapshot.go`.
- Modify `backend/internal/snapshot/snapshot_test.go`.
- Modify `docs/reference/configuration.md`.
- Modify `README.md`.

**Interfaces:**

- Produce `func (Snapshot) Validate() error`.
- Produce exported sentinel `ErrUnsupportedSnapshotFormat` for `errors.Is` checks.
- Keep `LoadFile(path string) (Snapshot, error)` but accept only `.sqlite`, `.db`, and `.duckdb`.
- Make `DefaultCachePath(repoRoot string)` return only `backend/data/tigerweb.sqlite`.
- Keep repository-contained environment overrides as the only implicit override path.

**Dependencies:** None.

**Test cycle:**

1. Add failing tests for unknown extensions, workspace-local defaults, duplicate layer keys, duplicate composite feature identities, missing layer references, invalid geometry types or coordinate ranges, missing feature identity fields, and invalid refresh states.
2. Run `cd backend && go test ./internal/snapshot` and confirm failures describe the missing validator or current permissive behavior.
3. Implement validation and loader integration.
4. Re-run the focused package, then `go test ./...` and `go vet ./...`.

**Implementation notes:**

- Validate after the compatibility NERIS layer is installed so every returned feature references the final registry.
- Do not require a fixed allowlist of layer-family names; validate stable identifiers and references so locally authored authority-observation layers remain possible.
- Polygon rings must be non-empty, finite WGS 84 coordinates with at least four positions and a closed first/last position. MultiPolygons must contain valid polygons.
- Open SQLite snapshots read-only with query-only mode so runtime access cannot mutate local evidence.
- Missing defaults should produce the missing SQLite path, not a DuckDB executable error.

**Completion evidence:** Focused snapshot tests pass; all loaders call the same validator; documentation matches the implemented precedence.

## Task 2: Enforce the public point-to-boundary authority boundary

**Outcome:** `POST /resolve` cannot accept caller-authored geographic context, and boundary overlaps are returned as deterministic ambiguity evidence rather than silently selecting whichever polygon appeared first.

**Files:**

- Modify `backend/internal/httpapi/resolve.go`.
- Modify `backend/internal/httpapi/geocode.go`.
- Modify `backend/internal/httpapi/resolve_test.go`.
- Modify `docs/reference/runtime-and-api.md`.
- Modify `README.md` where needed.

**Interfaces:**

- Add a private HTTP DTO containing `point`, `code_family`, `project_type`, and `applicability_date` only.
- Keep `regulatory.ResolutionRequest` as the internal policy-resolver request with normalized context.
- Add a typed boundary-ambiguity error carrying sorted `BoundaryMatch` observations.
- Return HTTP `409` with stable code `boundary_ambiguous`, a safe message, the ambiguous layer family, and all tied observations.

**Dependencies:** Task 1 ensures feature identities and geometry are valid.

**Test cycle:**

1. Add a failing request-boundary test proving a supplied `context` field is rejected.
2. Add failing overlap tests for multiple state, county, and municipality matches and deterministic observation ordering.
3. Add a success test proving special-area, tribal-area, and fire-jurisdiction overlaps remain preserved when the authority-bearing state/county/municipality context is unambiguous.
4. Implement the private DTO, deterministic collection, and classified HTTP errors.
5. Run `cd backend && go test ./internal/httpapi`, then the full backend checks.

**Implementation notes:**

- A point on a polygon edge remains an observation, but multiple authority-bearing matches must not become an arbitrary policy choice.
- Preserve all non-exclusive special, tribal, and fire observations.
- Do not infer that NERIS or another containing layer grants building-code authority.

**Completion evidence:** Public context injection fails with `400`; ambiguous containment fails with `409` and complete sorted observations; ordinary point and address lookup still resolve through the same internal policy path.

## Task 3: Report capability readiness and support browser preflight

**Outcome:** Liveness, required boundary readiness, optional geocoding, and regulatory capability are independently visible, and configured cross-origin JSON POST requests can complete their CORS preflight.

**Files:**

- Create `backend/internal/httpapi/readiness.go`.
- Modify `backend/internal/httpapi/handler.go`.
- Modify `backend/internal/httpapi/handler_test.go`.
- Modify `docs/reference/runtime-and-api.md`.

**Interfaces:**

- `GET /health` remains process liveness.
- `GET /ready` returns overall `ready`, `degraded`, or `not_ready` plus capability records for boundary resolution, coordinate resolution, local geocoding, regulatory resolution, and composed address lookup.
- Return `503` only when required boundary data is unusable; optional capability loss returns `200` with `degraded`.
- Allowed `OPTIONS` requests return `204` with `Access-Control-Allow-Methods`, `Access-Control-Allow-Headers`, and origin-specific `Vary` behavior.

**Dependencies:** Task 1 defines when a loaded boundary snapshot is semantically usable.

**Test cycle:**

1. Replace the old unconditional-ready assertion with failing full, degraded, and not-ready cases.
2. Add failing allowed and denied preflight cases.
3. Implement the smallest readiness and CORS behavior.
4. Run the focused HTTP tests, then all backend checks.

**Implementation notes:** Do not expose private absolute paths or raw request bodies in readiness responses.

**Completion evidence:** Readiness distinguishes all capability states; JSON POST preflight works only for configured origins.

## Task 4: Disclose an assumed applicability date

**Outcome:** A resolver result never silently presents the server's current date as though the user supplied it.

**Files:**

- Modify `backend/internal/regulatory/resolver.go`.
- Modify `backend/internal/regulatory/resolver_test.go`.
- Modify `docs/reference/regulatory-resolution.md` or `docs/reference/runtime-and-api.md`.

**Interfaces:** Keep `applicability_date` optional for compatibility, but append a deterministic warning whenever the resolver defaults it to the current UTC date.

**Dependencies:** None.

**Test cycle:** Add a failing warning assertion for an omitted date, implement the warning for both known- and unknown-profile paths, and run the regulatory package plus full backend checks.

**Completion evidence:** Omitted dates are visible as assumptions; explicit dates do not receive the warning.

## Task 5: Reconcile documentation, semantic records, and review findings

**Outcome:** Reader documentation, LORE accepted records, and the repository review describe the implemented authority boundaries without making generated projections authoritative or duplicating the open archaeology backfill.

**Files:**

- Create `docs/reviews/2026-08-03-comprehensive-repository-review.md`.
- Add a new append-only revision under `.lore/records/decision/decision.point-based-resolution-boundary/` after the implementation commit exists.
- Update `docs/README.md` if needed for the review index.
- Do not edit or duplicate the unmerged Deciduous archaeology graph from PR #40.

**Interfaces:** The new LORE record must cite an exact prior implementation revision and supersede the bootstrap record's public normalized-context wording. Generated LORE projections remain non-authoritative.

**Dependencies:** Tasks 1 through 4.

**Test cycle:** Use formatting, link/path inspection, pinned LORE extraction and validation, and a requirement-by-requirement self-review.

**Implementation notes:** The review must distinguish implemented fixes, open same-repository work, cross-repository follow-up, and limitations caused by unavailable hydrated datasets or live jurisdiction evidence.

**Completion evidence:** Documentation and accepted semantic records agree with code; PR #40 remains the sole archaeology backfill path; no generated prompt-bearing graph export is committed.

## Task 6: Publish and verify the draft pull request

**Outcome:** The branch is reviewable as an exact source state with CI evidence and a clear residual-risk ledger.

**Files:** No additional product files unless review findings require correction.

**Interfaces:** Draft pull request against `main` with exact base/head, summary, findings resolved, tests, limitations, and out-of-scope items.

**Dependencies:** Tasks 1 through 5.

**Test cycle:** Inspect the complete branch diff, run GitHub Actions at the exact head, inspect failed-job logs if needed, perform requirement and implementation-quality self-review, and rerun invalidated checks after any correction.

**Completion evidence:** Draft PR exists; exact-head workflow results are reported accurately; mergeability is not presented as correctness; any unavailable or failing evidence is explicit.
