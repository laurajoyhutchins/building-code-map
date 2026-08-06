# Geocoder Ranking and Interpolation Reference

## Contract identity

Building Code Map ranks local geocoder candidates with the versioned policy `geocoder-ranking-1.0`.

Every candidate exposes:

- `score_kind: "deterministic_quality"`;
- `score_factors`, an additive factor breakdown;
- `ranking_policy_version`;
- source name, source-record identity, and source vintage;
- `precision`, which remains either `address_point` or `interpolated`.

The JSON field is still named `confidence` for compatibility. It contains a deterministic additive quality score. It is not a probability, statistical confidence interval, calibrated likelihood, or statement that a location is legally authoritative. Positive configured source-priority factors can make the score greater than `1.0`.

## Policy factors

The default policy uses these contributions:

- address-point base: `0.70`;
- street-range base: `0.55`;
- exact normalized street: `0.05`;
- exact normalized city: `0.15`;
- exact supplied postal code: `0.10`;
- postal code not supplied: `0.05`;
- street-range parity matched: `0.05`;
- source priority: an optional configured contribution for a named source.

The minimum default quality is `0.85` for address points and `0.75` for street-range interpolation. Candidates within `0.05` of the highest eligible score remain ambiguous.

Source and record identifiers are used to order returned observations deterministically. Their lexical order does not break a quality tie and does not select a winner.

## Source priority

`RankingPolicy.SourcePriority` is an explicit map from source name to additive score contribution. `OpenWithPolicy` accepts a caller-supplied versioned policy, and tests exercise configured ordering.

The production server currently opens the geocoder with `Open`, which uses the default policy and an empty source-priority map. This pull request does not add an environment variable, configuration file, or command-line surface for production priority overrides. Until such a surface is reviewed, production does not silently prefer one source over another.

## Source vintage

Source vintage is provenance-only in `geocoder-ranking-1.0`. It is returned on every candidate but contributes no score and does not break a tie. A newer vintage is therefore not automatically treated as more authoritative, more accurate, or more applicable.

This is deliberate. A future policy may rank vintage only after defining source-specific date semantics and changing the ranking-policy version.

## Duplicate and conflicting candidates

Duplicate address points from distinct source records remain separate observations. Equal or materially close candidates return `ambiguous`; no candidate is selected.

City and supplied postal-code constraints are applied in the database query. Contradictory locality or postal evidence returns no match rather than awarding a partial score to an incompatible record.

## Street-range interpolation

An interpolated candidate exposes:

- source range identity;
- requested house number;
- normalized numeric range endpoints;
- ascending or descending range direction;
- parity;
- segment endpoint coordinates;
- interpolation fraction;
- derived coordinate;
- coordinate reference system;
- transformation identity;
- interpolation method version;
- positional-quality class.

The current method identity is `linear-street-range-1.0`. It performs linear interpolation between endpoint coordinates in `EPSG:4326`. Reversed numeric ranges are handled explicitly. Zero-length ranges are rejected during snapshot construction.

Interpolation establishes only a derived point along a source street segment. It does not establish parcel, structure, entrance, rooftop, driveway, delivery, emergency-access, or authoritative address-point precision. API responses preserve `precision: "interpolated"` and include a warning stating this limitation.

## Known remaining gaps

Issue #44 remains open for work not completed by this contract:

- ingesting and preserving source-provided side-of-street values;
- preserving original source text for range endpoints where it differs from normalized integers;
- wiring reviewed source-priority configuration into the production server;
- proving precision and ranking identity in every log and downstream persisted record.

These gaps must not be inferred as complete from the presence of interpolation or ranking fields in the API response.
