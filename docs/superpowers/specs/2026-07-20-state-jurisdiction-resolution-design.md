# State jurisdiction resolution design

## Goal

Turn accumulated state research into a deterministic, source-backed policy layer that converts local geographic matches into authority candidates, supported statewide adoption records, and explicit local follow-up without guessing.

## Architecture

The implementation preserves three boundaries:

1. Markdown reports remain evidence dossiers and the research-maintenance interface.
2. Versioned canonical JSON profiles contain executable authorities, adoption records, policies, sources, and verification state.
3. A generic Go resolver combines a profile with normalized geography and optional code-family/project context to produce a separately versioned result.

State behavior is represented in data. The resolver does not branch on state names.

## Canonical profile

A profile contains stable state metadata, official sources, authorities, authority relationships, adoption records, incorporated and unincorporated defaults, code-family overrides, project-type overrides, and verification metadata. Policies may replace or extend default authority/adoption selections. Unknown, conflicting, ambiguous, local-record-required, and insufficient-evidence outcomes remain first-class values.

Source and verification references are attached to material records. Validation rejects malformed dates, duplicate identifiers, missing sources, and references to unknown authorities or adoptions.

## Report compilation

Pilot reports embed exactly one fenced `json jurisdiction-profile` block. A standard-library Python tool validates and compiles profiles deterministically. It fails closed for malformed or internally inconsistent profiles. A separate coverage command scans all state reports and records missing profiles, invalid profiles, incomplete verification, open questions, code-family coverage, and resolver follow-up as distinct quality signals.

## Pilot states

- Colorado exercises predominantly local adoption plus a state electrical override.
- Florida exercises a statewide building-code baseline with local enforcement and amendment follow-up.
- New Jersey exercises a statewide construction code and a separate operational-fire authority path.

The profiles are intentionally marked partially verified until broader field-level validation supports production readiness.

## Resolver

Inputs are either coordinates or normalized geographic context, plus optional code family, project type, and applicability date. Coordinate resolution uses supported local snapshot polygons for state, county, municipality, special areas, tribal areas, and NERIS jurisdictions.

The resolver selects incorporated/unincorporated policy, overlays code-family and project-type rules, expands declarative candidate kinds, resolves adoption IDs, deduplicates results, and emits only cited evidence. Missing profile data returns `insufficient_evidence`; missing local adoption data returns `local_record_required` rather than an inferred edition.

## API and explorer

`POST /resolve` accepts a strict, size-limited JSON body. Unknown fields, multiple JSON values, invalid coordinates, and unmatched state boundaries fail explicitly. The explorer provides a separate resolution panel showing status, candidates, adoptions, required local records, warnings, and source links. Raw boundary inspection remains independent.

## Testing and operations

Go tests cover profile validation, deterministic resolution, all three pilot models, incorporated and unincorporated behavior, code-family overrides, unknown states, strict API parsing, coordinate matching, and polygon holes. Python tests cover deterministic compilation, reference validation, and quality semantics. Frontend tests cover response decoding. CI verifies all three stacks and checks that compiled profiles match their source reports.
