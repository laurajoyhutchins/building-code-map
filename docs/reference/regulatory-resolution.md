# Regulatory resolution

Building Code Map separates regulatory research from executable resolution so that evidence and uncertainty survive transformation.

## Data layers

1. **Evidence reports** in `reports/` remain the human-readable research and maintenance interface.
2. **Canonical jurisdiction profiles** in `backend/data/regulatory/` contain validated authorities, adoption records, declarative policies, source references, and verification state.
3. **State rule packs** in `backend/data/regulatory/rules/` add source-health records, applicability, date, amendment, and enforcement rules, field-level claims, conflict groups, and resolver-fixture identifiers without duplicating the base profile.
4. **Resolution results** combine a geographic context, applicability date, profile, and rule pack without pretending that missing local records are known.

The canonical state-profile contract is versioned in `schemas/regulatory/state-profile.schema.json`. Rule packs use `schemas/regulatory/state-rule-pack.schema.json`. The public response contract is versioned separately in `schemas/regulatory/resolution-result.schema.json`.

## Pilot coverage

The first resolver fixtures deliberately exercise different models:

- **Colorado:** ordinary construction generally requires a municipal or county adoption record; electrical regulation has a state-level override.
- **Florida:** a statewide building-code baseline is returned while local enforcement and amendment records remain visible follow-up.
- **New Jersey:** the statewide Uniform Construction Code and the operational fire-code authority are resolved through distinct code-family paths.

These profiles are `partially_verified`, not blanket production-ready claims. Their source registries, rule packs, claims, and unresolved local records are returned with results.

## Compile and validate profiles

Pilot reports contain exactly one fenced `json jurisdiction-profile` block. A report may also contain a fenced `json jurisdiction-rules` block. The pilot rule packs are maintained as separately versioned canonical overlays in `backend/data/regulatory/rules/`; the compiler validates either representation against the same profile references.

Compile the reports with:

```bash
python reports/tools/regulatory.py compile reports/pilots/*.md --output-dir backend/data/regulatory
```

Verify checked-in profile and rule-pack artifacts are current:

```bash
python reports/tools/regulatory.py compile reports/pilots/*.md --output-dir backend/data/regulatory --check
```

Malformed JSON, unknown source IDs, unknown authority, adoption, rule, or claim references, invalid status values, duplicate identifiers, invalid dates, and broken conflict records fail non-zero with the source path and context.

## Coverage and quality reporting

Generate a repository-wide, field-aware report:

```bash
python reports/tools/regulatory.py coverage \
  --reports-root reports \
  --json-output .tmp/regulatory-coverage.json \
  --markdown-output .tmp/regulatory-coverage.md
```

The report tracks profile validity, all four rule families, claims and conflicts, resolver fixtures, stale or unavailable sources, local follow-up, and open questions separately. File presence never counts as readiness.

Readiness labels mean:

- `blocked`: no valid profile can be evaluated;
- `needs_review`: a profile exists but lacks a complete executable pilot rule set;
- `pilot_ready`: a valid profile and rule pack cover applicability, dates, amendments, enforcement, and resolver fixtures;
- `production_ready`: the pilot conditions are met and the profile is fully verified, sources are current and available, claims do not conflict, policies do not require unresolved lookup, and open questions are zero.

A state may be useful and `pilot_ready` while still intentionally returning `local_record_required` for a location. That is a legal-data requirement, not a failed resolver.

## Resolution API

`POST /resolve` accepts either coordinates:

```json
{
  "point": { "longitude": -104.9903, "latitude": 39.7392 },
  "code_family": "building",
  "project_type": "state_owned",
  "applicability_date": "2026-07-20"
}
```

or an already normalized `context`. Coordinate requests are matched against state, county, municipality, special-area, tribal-area, and NERIS jurisdiction polygons available in the local snapshot.

The applicability date defaults to the current UTC date. Current or future adoption records that do not apply on the requested date are excluded. When a profile points to adoption records but none apply on that date, the resolver fails closed with `insufficient_evidence` and asks for the effective adoption and amendment record for that date.

The response distinguishes:

- geographic matches;
- resolver-selected authority candidates;
- authority relationships that explain the enforcement path;
- adoption records supported for the supplied date and scope;
- applicable applicability, date, amendment, and enforcement rules;
- field-level supporting or conflicting claims;
- required local records;
- source health, warnings, and evidence;
- `resolved`, `partially_resolved`, `local_record_required`, `ambiguous`, `conflicting`, and `insufficient_evidence` states.

The resolver interprets data. It does not contain a state-name switch statement, silently choose between conflicting claims, reproduce model-code text, substitute a statewide rule for a required local adoption, or return a currently recorded edition for a date before it became operative.

## Runtime configuration

The server loads base profiles and optional `rules/` overlays from `backend/data/regulatory` by default when started from `backend/`. Override the path with:

```bash
go run ./cmd/server --regulatory-data /path/to/regulatory
```

Rule packs are matched to profiles by `state_id`. Duplicate packs, packs for unknown states, invalid references, or unknown fields fail catalog loading. If profiles are unavailable or invalid, boundary exploration can still start, but `/resolve` fails closed with `503 Service Unavailable`.
