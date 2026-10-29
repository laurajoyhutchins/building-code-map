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

Install the two existing geometry libraries rather than implementing a
coordinate projection or polygon clipping algorithm:

```sh
python3 -m pip install 'shapely==2.1.2' 'pyproj==3.7.2'
sha256sum /secure/us-land.geojson
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

## Workload and output contracts

- `schemas/bulk-workload-v1.schema.json`: one JSONL sample (point **or**
  address, optional project type/facts/per-sample applicability date, optional
  grid provenance metadata).
- `schemas/bulk-result-v1.schema.json`: one sanitized result per submitted
  sample. Only status/category, relevant code family labels, evidence count,
  jurisdiction ID and latency are retained; never copy full source responses.
- `schemas/bulk-summary-v1.schema.json`: aggregate counts, deterministic
  input SHA-256, shard identity, throughput, and p50/p95/p99 latency.
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
