# Address-to-result pipeline map

## 1. Preserve input

The public client retains the entered address or coordinate string. Coordinate parsing and address
classification are separate. A direct coordinate never needs to be converted into a synthetic
address.

## 2. Normalize a civic address

Current normalization:

- trims and collapses whitespace;
- uppercases comparison fields;
- removes presentation punctuation;
- requires exactly three comma-delimited parts;
- requires a numeric house number with an optional trailing letter;
- normalizes a bounded list of street suffixes and cardinal directions;
- normalizes full U.S. state names or two-letter abbreviations;
- accepts an optional five-digit ZIP or ZIP+4 and stores the five-digit ZIP;
- rejects PO boxes, blank, incomplete, and malformed input.

The original query and deterministic normalized form both remain available. Unit designators and
ambiguous locality resolution are not implemented as independent concepts.

## 3. Geocode locally

The service queries address points first, constrained by house number, normalized street, city,
state, and optional ZIP. It uses street ranges only when no address point qualifies. Street ranges
must contain the house number and match parity. Interpolation is linear between range endpoints.

Candidate ordering is stable: confidence descending, then source record ID. Materially tied
qualifying candidates return `ambiguous`; the system does not choose arbitrarily.

## 4. Establish a point and its provenance

A selected candidate includes coordinates, matched address, precision, ranking confidence, source
name, source record ID, and source vintage. Precision is `address_point` or `interpolated`. It does
not independently prove rooftop, parcel, entrance, or unit accuracy.

## 5. Resolve supported geometries

The point is tested against local polygon and multipolygon features. The normalized context may
include state, county, municipality, incorporated status, special areas, tribal areas, and fire
jurisdictions. The geometry result is a containment fact, not yet a legal-authority conclusion.

## 6. Interpret authority policy

The resolver selects a version-compatible state profile, applies incorporated or unincorporated
policy, overlays code-family and project-type rules, expands authority candidates, resolves adoption
identifiers, and keeps source and verification references.

## 7. Apply date-aware adoption records

An applicability date filters records and rules. The resolver must not answer a historical query
with a merely current adoption. Unsupported history should remain insufficient or require local
records.

## 8. Emit inspectable outputs

`POST /lookup` returns both geocoding and regulatory results. `POST /resolve` remains the direct
point or normalized-context surface. The public result shows the matched location, authorities,
adoptions, special conditions, local confirmation items, jurisdiction structure, warnings, and
sources.

## Failure boundaries

- invalid address: request rejected;
- geocoder unavailable: address endpoints unavailable, coordinate resolution remains available;
- no qualifying candidate: `not_found`;
- materially tied candidates: `ambiguous`;
- point outside supported state geometry: resolution fails;
- missing profile: `insufficient_evidence`;
- missing local adoption or enforcement record: `local_record_required`;
- conflicting evidence: `conflicting`.

Each boundary prevents a weaker stage from silently manufacturing confidence for the next stage.
