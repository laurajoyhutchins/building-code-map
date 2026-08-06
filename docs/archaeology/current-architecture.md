# Current architecture

## System boundary

Building Code Map is a local-first applicability resolver. It accepts an address or coordinate, derives a provenance-bearing point, observes containing geographic features, applies source-backed regulatory policy for an as-of date, and returns candidate authorities, adopted instruments, uncertainty, and required confirmation.

It is not a general geocoder, national GIS warehouse, code-text corpus, legal-opinion engine, or substitute for the authority having jurisdiction.

## Runtime flow

```text
address input
    -> deterministic normalization
    -> local SQLite geocoder
    -> address point or labeled interpolation
    -> provenance-bearing point

coordinate input
    -> provenance-bearing point

point
    -> validated local boundary snapshot
    -> geographic observations
    -> ambiguity check
    -> state regulatory profile
    -> date and project-scope rules
    -> authority candidates and adoptions
    -> evidence, warnings, and local-record requirements
```

## Public API boundary

`POST /resolve` accepts a point. It does not accept caller-authored normalized jurisdiction context as trusted input. The server derives geographic observations from its admitted boundary snapshot.

`POST /lookup` geocodes an address locally and invokes the same point path. The response preserves both geocoding and regulatory evidence.

If an applicability date is omitted, the compatibility default is the current UTC date and the result carries an explicit warning. The default is not evidence of project applicability.

## Boundary behavior

The runtime admits a snapshot only after checking the schema and semantic invariants consumed by spatial indexes and resolution. SQLite is the workspace-local default. Legacy DuckDB compatibility is explicit rather than inferred from an unknown extension or machine-global path.

State, county, municipality, special-area, American Indian area, and NERIS features are geographic observations. Multiple peer state, county, or municipality matches return deterministic `409 boundary_ambiguous` evidence instead of encounter-order selection. Special-area, tribal-area, and NERIS containment is never automatically promoted to legal authority.

## Regulatory behavior

The generic resolver loads declarative profiles and rule packs. It has no state-name branches. Profiles distinguish current, historical, transition, and pending records; authority candidates; enforcement follow-up; code-family overrides; source references; verification status; and required local records.

Current executable pilot profiles are:

- Colorado
- Florida
- New Jersey
- Virginia
- Oregon
- North Carolina

Colorado, Florida, and New Jersey also have production-ready declared-scope manifests. Those findings are scoped and do not assert complete statewide municipal, amendment, project-type, or historical coverage.

## Readiness

`/health` proves process response. `/ready` reports capability-specific readiness for boundary, geocoder, regulatory, and composed lookup paths. A missing geocoder may leave coordinate resolution usable. Missing regulatory profiles leave geographic inspection available without permitting code conclusions.

## Remaining architectural gaps

- a canonical national governmental-entity and legal-authority registry;
- explicit state-regime and local-jurisdiction inheritance contracts;
- production snapshot manifests, checksums, activation receipts, and rollback identity;
- nationwide local-adoption and amendment coverage;
- historical address, boundary, authority, and adoption timelines;
- frontend runtime decoding for every API payload;
- source-specific geocoder ranking and complete interpolation derivation.
