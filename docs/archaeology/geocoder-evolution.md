# Geocoder evolution map

## Deferred during the first product milestone

The first executable location-to-authority milestone deliberately accepted coordinates or normalized
geographic context. Issue #4 recorded address normalization and local geocoding as deferred. This
kept state-policy modeling from becoming entangled with an unsolved national address-data problem.

## External geocoding was not accepted as the durable request path

A per-request provider would make results depend on network availability, pricing, rate limits,
vendor ranking, address disclosure, and opaque updates. A RAG or generated answer would add further
nondeterminism. The selected approach instead builds reviewed address data into a separate local
snapshot.

## SQLite vertical slice

Merged PR #37 introduced:

- `backend/internal/geocoder/normalize.go`;
- `backend/internal/geocoder/schema.go`;
- `backend/internal/geocoder/sqlite.go`;
- a transactional CSV-to-SQLite builder;
- synthetic address-point and range fixtures;
- `/geocode` and `/lookup`;
- public address input and provenance presentation;
- a LORE record for the component and the no-network constraint.

Snapshot schema version `1` contains:

```text
geocoder_metadata(key, value)
address_points(...)
street_ranges(...)
```

Indexes cover normalized lookup fields. Source name plus source record ID is unique in each source
table. The service opens the snapshot read-only and rejects incompatible schema versions.

## Accuracy actually provided

The implementation distinguishes:

- `address_point`: coordinates supplied by the imported address record;
- `interpolated`: a calculated point along a compatible street range.

It does not compute or certify rooftop error, parcel placement, entrance placement, unit placement,
or rural-route quality. Ranking confidence is deterministic scoring based on field agreement, not a
calibrated positional probability.

## Address rules actually implemented

Implemented: house number, street, city, state, optional ZIP, bounded suffixes and directions,
punctuation cleanup, case normalization, exact/locality-constrained SQL, optional omitted suffix
matching, deterministic ambiguity, parity-aware ranges, and original-query preservation.

Not implemented: unit parsing, postal-only input, PO boxes, free-form locality inference, multiple
candidate localities, landmarks, intersections, parcels, buildings, reverse geocoding, historical
addresses, network fallback, machine-learned ranking, and nationwide source acquisition.

## Packaging and refresh

The builder creates a temporary database, validates input and provenance, builds indexes, and
replaces the target only after success. Hydrated databases remain ignored. A production snapshot
therefore requires an independently reviewed source, redistribution decision, source vintage, and
repeatable refresh process. Automated refresh scheduling remains deferred.
