# Building Code Authority Engine Design

## Status and base

This design implements the approved Building Code Authority Engine refactor as seven stacked draft pull requests. The current implementation branch is temporarily based on the exact head of predecessor PR #48, `d7908d9c4fc55f34d27f204310acc5cdd03919e7`, because PRs #45, #47, and #48 remain open drafts and the owner instruction is to leave them unmerged during this session.

The intended ancestry is:

```text
main
  #45 snapshot manifests and activation identity
    #47 frontend runtime decoders
      #48 geocoder ranking and provenance
        authority-engine-contracts
          authority-engine-core
            authority-engine-transports
              bcm-cli-bundle
                bcm-mcp
                  website-engine-v1
                    authority-engine-offline-release
```

After the predecessor stack merges, this stack must be transplanted or rebased onto the resulting `main`, retargeted, and reverified at each exact head.

## Ownership boundaries

Building Code Map remains the authority for compiled boundary snapshots, geocoding snapshots, regulatory catalogs, jurisdiction resolution, evidence, provenance, packaging, and the public website. Building Code AST remains responsible for parsing and navigating code documents; this engine resolves which authorities and code records apply and does not answer code-content questions.

PR #45 owns component snapshot manifests, verification, activation, prior-active identity, last-known-good identity, fail-closed boundary startup, optional geocoder degradation, and readiness snapshot identities. The aggregate engine bundle composes those component identities and does not replace their manifest contract.

PR #47 owns frontend runtime decoding, path-aware trust-boundary failures, structured API errors, ambiguity distinctions, readiness identity decoding, unknown optional-field preservation, and cancellation behavior. The v1 client extends those decoders rather than bypassing them.

PR #48 owns deterministic geocoder ranking and its full provenance chain. The engine preserves ranking-policy identity, factor contributions, source and vintage identity, interpolation derivation, coordinate evidence, and equal-quality ambiguity through every adapter.

PR #14 remains the owner of report conversion safety and canonical regulatory artifacts. PR #35 remains a staged LORE adoption lane. Neither is absorbed into the engine stack, and engine documentation remains hand-authored under the existing documentation tree rather than `docs/generated/lore/`.

Issue #7 has an active but unmerged branch, `feature/on-demand-boundary-detail`, at `e34c2a9`. Its map/detail DTO and cancellation work is preserved as a coordination dependency; the engine stack does not recreate it.

## Architecture

The engine is a transport-independent Go package composed from verified snapshots and the existing regulatory catalog:

```text
verified component snapshots
       ↓
Building Code Authority Engine
       ↓
CLI | HTTP v1 | MCP | website
```

The engine owns query validation, address-to-point orchestration, geography resolution, regulatory invocation, typed outcomes, stable ordering, explicit applicability dates, and complete provenance. Geography matching moves out of HTTP into a reusable resolver. Regulatory resolution receives an explicit date or injected clock and never reads wall-clock time directly.

Legacy HTTP routes remain compatibility projections. Versioned HTTP, CLI, MCP, and the website all consume the same engine methods and result/error contracts. Adapters may change framing, status, and presentation, but cannot choose authorities, applicability, or evidence.

## Contracts

The public `backend/engine` package defines `Query`, normalized query, point and location results, resolution results, provenance, diagnostics, readiness, bundle identity, and typed `EngineError` values. Resolution requires an explicit `applicability_date`; legacy routes retain their current omitted-date compatibility behavior until the v1 migration is complete.

Stable error codes distinguish invalid input, invalid coordinates, geocoder outcomes, unsupported coverage, missing regulatory data, bundle invalidity, ambiguity, insufficiency, and internal failure. Results are JSON-stable through explicit schema versions, deterministic collection sorting, and exact component identity fields.

The aggregate bundle manifest is content-addressed and validates schema version, safe relative paths, component role uniqueness, component digests, source commit, engine version, and coverage metadata. It composes the existing boundary and geocoder snapshot manifests and adds the regulatory catalog digest.

## Testing and verification

Each PR has its own focused contract and full-suite gate. Tests use small explicit fixtures and cover legacy HTTP behavior, query normalization, coordinate and boundary edge cases, deterministic clocks, provenance completeness, bundle digest validation, CLI stream separation and exit codes, v1 schemas, MCP framing and cancellation, runtime frontend decoding, and cold-room offline startup.

Exact-head evidence records the branch, base, commit SHA, commands, and results in every PR body. Missing Go tooling in the current environment is a verification blocker to report honestly; it is not a reason to weaken the contracts or silently claim a passing check.

