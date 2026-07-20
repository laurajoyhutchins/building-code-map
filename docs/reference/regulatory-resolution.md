# Regulatory resolution

Building Code Map separates regulatory research from executable resolution so that evidence and uncertainty survive transformation.

## Data layers

1. **Evidence reports** in `reports/` remain the human-readable research and maintenance interface.
2. **Canonical jurisdiction profiles** in `backend/data/regulatory/` contain validated authorities, adoption records, declarative policies, source references, and verification state.
3. **Resolution results** combine a geographic context with a profile without pretending that missing local records are known.

The canonical state-profile contract is versioned in `schemas/regulatory/state-profile.schema.json`. The public response contract is versioned separately in `schemas/regulatory/resolution-result.schema.json`.

## Pilot coverage

The first resolver fixtures deliberately exercise different models:

- **Colorado:** ordinary construction generally requires a municipal or county adoption record; electrical regulation has a state-level override.
- **Florida:** a statewide building-code baseline is returned while local enforcement and amendment records remain visible follow-up.
- **New Jersey:** the statewide Uniform Construction Code and the operational fire-code authority are resolved through distinct code-family paths.

These profiles are `partially_verified`, not blanket production-ready claims. Their source registries and unresolved local records are returned with results.

## Compile and validate profiles

Pilot reports contain exactly one fenced `json jurisdiction-profile` block. Compile them with:

```bash
python reports/tools/regulatory.py compile reports/pilots/*.md --output-dir backend/data/regulatory
```

Verify checked-in artifacts are current:

```bash
python reports/tools/regulatory.py compile reports/pilots/*.md --output-dir backend/data/regulatory --check
```

Malformed JSON, unknown source IDs, unknown authority/adoption references, invalid status values, or duplicate identifiers fail non-zero with the report path and context.

## Coverage and quality reporting

Generate a repository-wide, field-aware report:

```bash
python reports/tools/regulatory.py coverage \
  --reports-root reports \
  --json-output .tmp/regulatory-coverage.json \
  --markdown-output .tmp/regulatory-coverage.md
```

A missing profile is a quality signal, not an invented failure percentage. The report keeps profile validity, verification, sources, authorities, adoptions, code-family coverage, local follow-up, and open questions distinct.

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

The response distinguishes:

- geographic matches;
- supported state regulatory facts;
- resolver-selected authority candidates;
- statewide adoption records supported for the supplied scope;
- required local records;
- warnings and evidence;
- `resolved`, `partially_resolved`, `local_record_required`, `ambiguous`, `conflicting`, and `insufficient_evidence` states.

The resolver interprets data. It does not contain a state-name switch statement, silently choose between conflicting claims, reproduce model-code text, or substitute a statewide rule for a required local adoption.

## Runtime configuration

The server loads profiles from `backend/data/regulatory` by default when started from `backend/`. Override the path with:

```bash
go run ./cmd/server --regulatory-data /path/to/regulatory
```

If profiles are unavailable or invalid, boundary exploration can still start, but `/resolve` fails closed with `503 Service Unavailable`.
