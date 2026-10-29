# Bulk project-query stress testing

BCM stress workloads exercise the canonical `POST /v1/project-code-basis`
API, not a second code-adoption engine. The first workload is an approximately
100-mile-square **equal-area** grid over a pinned U.S. land polygon. The same
runner accepts arbitrary point or address JSON Lines workloads for boundary
pairs, historical date transitions, ambiguity cases, concurrency and soak runs.

## Geometry and sampling

Use a rights-cleared Census land-geometry snapshot converted into WGS84
GeoJSON and retain its official origin and SHA-256 in private run custody.
The input must contain only the selected coverage area, for example all
50 states and DC, not Puerto Rico or the territories if the benchmark scope
is 50 states and DC. Do not commit national geometry or real bulk results to
this public source repository.

Use the Census 2024 1:500,000 **state** cartographic boundary ZIP
(`https://www2.census.gov/geo/tiger/GENZ2024/shp/cb_2024_us_state_500k.zip`).
The preparer includes only the 50 states and DC (51 state equivalents),
checks original Census field identities and reports both ZIP and derived
geometry digests. Keep both and the source URL with the private run receipt.
It uses off-the-shelf shapefile reading, polygon operations and projection,
not a homegrown geospatial parser.

```sh
python3 -m pip install 'pyshp==2.3.1' 'shapely==2.1.2' 'pyproj==3.7.2'
curl --fail --location --output /secure/census-states.zip \\
  https://www2.census.gov/geo/tiger/GENZ2024/shp/cb_2024_us_state_500k.zip
python3 scripts/prepare_census_us_land.py \\
  --archive /secure/census-states.zip --output /secure/us-land.geojson
# Save the printed source archive and derived geometry SHA-256 values.
python3 scripts/generate_grid.py \
  --geometry /secure/us-land.geojson \
  --sha256 ACTUAL_64_CHARACTER_SHA256 \
  --miles 100 \
  --output /secure/us-100mi.jsonl
```

The generator uses the world equal-area EPSG:6933 CRS with a fixed global
origin and a cell side of 100 statute miles (160,934.4 metres). It emits
every cell with **positive area of land overlap**, using the projected
center when on land, otherwise a representative point from the land
intersection. This is an area-equivalent square lattice, not a promise
that both geodesic sides are exactly 100 miles at every latitude.
This coarse cartographic sampling cannot prove address-level boundary
precision or detect every enclave, narrow municipality or disputed claim.
Antimeridian-crossing source rings must be normalized rather than treated
as world-width polygons. Geometry snapshots are not legal authority.

IDs are deterministic for identical geometry/projection/grid spacing, and
samples include source geometry SHA-256, EPSG, cell row/column and estimated
land fraction. A coastal cell is sampled **on land**, not blindly at sea.
No code or jurisdiction is inferred by grid generation.

## State-stratified complement

A 100-mile grid intersects states that its selected representative points may
not actually sample, especially small northeastern states and DC. Audit
Census **candidate** state membership separately, preserving zero-count
states, overlap/edge ambiguities and unmatched points. Then generate one
supplemental positive-land representative point for each Census state
equivalent. This is a geographic sampling probe, **never** evidence of a
verified governmental authority or adopted code.

```sh
python3 -m scripts.state_coverage \
  --geometry /secure/us-land.geojson \
  --sha256 ACTUAL_64_CHARACTER_SHA256 \
  --grid /secure/us-100mi.jsonl --expect-51 \
  --audit /secure/us-100mi-state-audit.json \
  --representatives /secure/us-state-probes.jsonl
python3 scripts/bulk_query.py \
  --input /secure/us-state-probes.jsonl \
  --results /secure/state-probe-results.jsonl \
  --summary /secure/state-probe-summary.json \
  --date 2026-10-28 --concurrency 4 --qps 2 --max-requests 51
```

The versioned `schemas/bulk-state-audit.schema.json` allows only source
coverage counts and ambiguity diagnostics. The additional state probes use
the existing `bulk-workload-v1` contract and must not carry regulatory
verdicts. Neither the grid nor this representative sample can establish
address-level applicability or jurisdiction boundary accuracy.

## Safe first execution

Start the public fictional demo service in a separate terminal:

```sh
go run ./cmd/bcm
```

Then run a bounded workload; every real geographic sample will be refused
by this *synthetic-only* service, which is the expected first end-to-end
transport test, not an assessment of national code coverage:

```sh
python3 scripts/bulk_query.py \
  --input /secure/us-100mi.jsonl \
  --results /secure/100mi-results.jsonl \
  --summary /secure/100mi-summary.json \
  --date 2026-10-28 \
  --concurrency 4 --qps 2 --max-requests 5000
```

To target a real, separately approved deployment, use `--endpoint
https://<approved-host>/v1/project-code-basis --allow-remote` and optionally
`--token-env BCM_BULK_TOKEN`. A missing token fails before traffic starts.
Remote HTTP, embedded credentials, redirects, unexpected paths and implicit
production targets are rejected. Keep `--qps`, `--concurrency`, per-request
`--timeout`, and `--max-requests` explicit. No tool call creates cloud
infrastructure or automatically runs a paid load. The overall limit counts
**executed samples per shard**, not raw lines in the input file.

For deterministic parallel partitions, add `--shard-count 4 --shard-index 0`
(and use indices 1, 2, 3 separately). All shards read the same input bytes.
They select IDs by SHA-256 modulo the shard count. Aggregate results only
after checking that the source SHA-256, date/scope, shard count and all
distinct indices match; do not sum p95 latency across shards.

## Source-bound paired-edge workload (offline)

The national 100-mile grid cannot resolve narrow municipalities or
jurisdiction boundaries. Generate a separate adversarial workload from a
**digest-pinned, rights-cleared WGS84 GeoJSON FeatureCollection** of source
polygons with stable unique string IDs (default \`properties.id\`). The
input may be private. Never substitute unreviewed Census polygons for a
legally effective authority boundary, and do not commit private source
geometry or individual results to public Git.

\`\`\`sh
python3 -m scripts.generate_boundary_pairs \
  --geometry /secure/approved-geographic-source.geojson \
  --sha256 ACTUAL_SOURCE_SHA256_64_LOWERCASE_HEX \
  --id-field id \
  --offset-metres 100 \
  --max-pairs 10000 \
  --output /secure/boundary-pairs.jsonl \
  --summary /secure/boundary-pairs-receipt.json

# Replay via the EXISTING canonical API after the same live preflight
# and bounded, approved source/bundle identity checks described below.
python3 scripts/bulk_query.py \
  --input /secure/boundary-pairs.jsonl \
  --results /secure/boundary-results.jsonl \
  --summary /secure/boundary-run-summary.json \
  --date 2026-10-28 --concurrency 1 --qps 1 --max-requests 100
\`\`\`

This generator uses Shapely STRtree, projected EPSG:6933 geometry, and
PyProj rather than custom polygon parsers. For each *exact*, positive-length
shared edge between two source polygons, it tries points on either side
at a requested **projected** offset, shrinking only when needed. It
accepts a pair only when each rounded WGS84 point lies strictly inside
exactly its own source polygon and no other input polygon. Feature
ordering does not change pair IDs or coordinates. The same-source SHA-256
binds rows to the input snapshot. A final receipt records eligible
touching pairs, produced pairs, unresolved edges, overlapping interiors,
and the workload bytes' SHA-256.

Overlaps and missing exact shared edges must not be repaired silently:
a point-only contact is not an edge; an overlap is not a valid pair;
gaps and imprecise noncoincident boundaries produce no pair. Boundaries
at very narrow features or multi-feature junctions may have zero safe
samples. The CLI refuses unbounded counts and refuses overwriting
the output when zero pairs or a limit violation is detected.

\`boundary.classification\` is always
\`source_polygon_candidate_only\`. Its \`feature_id\` is *not* the
verified enforcing governmental authority. The offset is in EPSG:6933
projected metres, not guaranteed geodesic distance. The generator
does not evaluate annexations, legal effectivity, addresses or code
adoption. Pair outcomes require independent source and legal review
in the actual provider, not merely a changed API jurisdiction field.

Schemas: \`bulk-workload-v1.schema.json\` provides the optional typed
boundary metadata, \`bulk-boundary-summary-v1.schema.json\` specifies
the generation receipt, and \`bulk-query.odcs.json\` binds both.

## Controlled live-provider preflight

The original 516-cell source-backed 2024 Census workload and the 51
state-stratified companion probes have run against the **synthetic-only** public
API, not against a real reviewed regulatory provider. A real Cloud Run or
equivalent service must be demonstrably present; the tester must not start a
new deployed service, publish a private corpus, infer operator SourcePolicy
from a source ID, or silently send a nationwide job to an unverified URL.

Before a single real-provider request, obtain independent custody/review
for the exact private provider bundle and source commit and check the live
service's origin, authentication and scope. The local Go server must be
configured as a private provider adapter, not a synthetic demo. Keep credentials
in the server's private environment and use `--token-env` only when the
**public** canonical API itself requires an explicit bearer token. Cloud Run
workload identity is the preferred server-to-private-provider boundary.
Never place tokens in the command line, git or page configuration.

Prepare exactly **one** reviewed geographic sample (JSONL, single point;
one record only) in private storage. It is a transport canary, not a truth
case. Do not use a synthetic DEMO-XX address or substitute a random grid cell
for a documented site:

```sh
python3 -m scripts.live_bulk_preflight \
  --endpoint https://APPROVED_BCM_API/v1/project-code-basis \
  --allow-remote \
  --canary /secure/reviewed-canary.jsonl \
  --date 2026-10-28 \
  --expected-bundle EXACT_NON_DEMO_BUNDLE_ID \
  --expected-source EXACT_40_CHAR_GIT_SOURCE_COMMIT \
  --token-env BCM_BULK_TOKEN \
  --output /secure/live-canary-receipt.json
```

If the public API does not require its own bearer token, omit
`--token-env`. The canary checks `GET /health` reports
`private_provider_adapter`, `GET /v1/readiness` reports ready, and one
point's canonical 200 response carries the independently pinned bundle and
source commit. It explicitly refuses the fictional demo, drift, unavailable
coordinate support or provider errors. The receipt declares
`legal_accuracy_reviewed:false` and `bulk_run_authorized:false`; it
cannot replace manual operative law and site-authority review.

Only after independently approving the **same target, scope and QPS** run
a small separate sampled input (for example one source-state point), then
the full verified Census input. When using the existing runner against a real
endpoint always pass the independent `--expected-bundle` and
`--expected-source` pins, plus `--fail-on-errors`:

```sh
python3 scripts/bulk_query.py \
  --input /secure/approved-100mi-subset.jsonl \
  --results /secure/live-outcomes.jsonl \
  --summary /secure/live-summary.json \
  --date 2026-10-28 \
  --endpoint https://APPROVED_BCM_API/v1/project-code-basis \
  --allow-remote \
  --expected-bundle EXACT_NON_DEMO_BUNDLE_ID \
  --expected-source EXACT_40_CHAR_GIT_SOURCE_COMMIT \
  --concurrency 1 --qps 1 --max-requests 10 --timeout 8 \
  --fail-on-errors
```

The batch runner keeps `runtime_identity_mismatch` distinct from legal
unknowns, rejects returned DEMO jurisdiction identity under an explicit
live pin, and exits nonzero for HTTP/transport errors, invalid provider
response or falsely asserted verification when `--fail-on-errors` is set.
A 404/422 `unresolved` sample has **no provider identity witness** and
therefore cannot prove the remote backend revision per point. A successful
canary pin does not prove that a rolling deployment never changed. For
rigorous full-batch replay, retain authenticated service revision identity
and provider bundle snapshots separately, and prevent deployment changes
during the approved run.

Neither a valid provider identity nor a `verified` API field constitutes
independently checked legal truth. Use the separately reviewed
`bcm-accuracy` suite, including amendments, effective dates and AHJ,
to distinguish real correct assertions from plausible but unsupported ones.

## Workload and output contracts

- `schemas/bulk-workload-v1.schema.json`: one JSONL sample (point **or**
  address, optional project type/facts/per-sample applicability date, optional
  grid provenance metadata).
- `schemas/bulk-result-v1.schema.json`: one sanitized result per submitted
  sample. Only status/category, relevant code family labels, evidence count,
  jurisdiction ID and latency are retained; never copy full source responses.
- `schemas/bulk-summary-v1.schema.json`: aggregate counts, deterministic
  input SHA-256, shard identity, throughput, and p50/p95/p99 latency.
- `schemas/live-bulk-canary.schema.json`: one non-authoritative live provider
  identity/transport receipt.
- `schemas/bulk-query.odcs.json`: paired ODCS cross-repository contract.

A 200 response is checked for matching project ID/date, v0.3 schema version,
jurisdiction, and nonempty code results. The categories `verified`,
`conditional`, `not_verified`, `unresolved`, `unsupported_verification`,
`invalid_response`, `service_error`, `http_error` and
`transport_error` remain separate. Even a seemingly `verified` answer
is **not independently proven correct** by this transport harness. Compare
selected cases against independent reviewed evidence with `bcm-accuracy`.
A `not_verified` or unresolved reply is not a successful legal conclusion,
and a response with an unsupported verified assertion is flagged.

## Additional workloads and acceptance

Keep input creation separate from execution; the runner is reusable without
changing BCM's query semantics. Next workloads: both sides of municipal and
county boundaries; tiny enclaves; overlapping source polygons; transition
effective dates; adversarial ambiguity; high-contention and soak; and an
independently reviewed stratified truth set. Runs should pin API source SHA,
provider bundle identity, geometry/source snapshot hashes, applicability date,
runtime environment and execution receipt. The public summary alone is
**not** sufficient evidence of nationwide coverage or legal correctness.
