# External public datasets: acquisition boundary

`cmd/bcm-sources` is a source-acquisition utility. It creates **local, content-identified, unreviewed source snapshots** and, for Colorado DOLA municipalities, a non-authoritative GeoJSON observation projection. It does **not** activate sources for `engine.Runtime`, decide jurisdiction, claim legal adoption, or publish a regulatory corpus.

The open-source software owns this acquisition code; `bcm-regulatory-data` owns privately maintained research, source review, and regulatory bundle admission. The existing `building-code-map-alpha` custody and cutover gates in [source ownership](source-ownership.md) remain in force.

## Listed sources (checked October 20, 2026)

| Source ID | Format | Role | Limit |
| --- | --- | --- | --- |
| `doe-residential` | XLSX | DOE modeled residential energy-code equivalence | No proof of enacted code |
| `doe-commercial` | XLSX | DOE modeled commercial energy-code equivalence | No proof of enacted code |
| `fema-bcat-2026-06` | ZIP of Esri geodatabase | Historical public third-party preservation of FEMA BCAT | Research leads, not controlling ordinances; large source |
| `usdot-nad-2026-09` | ZIP of text | National Address Database; candidate address points | Incomplete coverage; large source |
| `co-dola-municipal` | Socrata JSON | Best-effort Colorado municipal geographic observations | Not a legally authoritative boundary |
| `census-co-places-2026` | TIGER/Line ZIP | Colorado place polygons and Census IDs | Census geography is not code authority |

**Datasets intentionally not treated as bulk feeds:** FEMA National Flood Hazard Layer, USGS seismic design maps and PNNL IECC climate zones need scoped, edition-aware queries or specialized source adapters. Their data cannot be substituted for local code-adoption evidence.

## Acquire and inspect

Run from the public BCM repository, **writing all outputs to an approved private folder outside its Git history**:

```sh
go run ./cmd/bcm-sources list
go run ./cmd/bcm-sources fetch -source co-dola-municipal -out /private/bcm/co-municipal.json
go run ./cmd/bcm-sources normalize-colorado \
  -in /private/bcm/co-municipal.json \
  -receipt /private/bcm/co-municipal.json.receipt.json \
  -out /private/bcm/co-municipal.candidates.geojson
```

For subsequent acquisitions, pass `-sha256 <previously-approved-64-character-digest>` so changed bytes fail closed rather than silently replacing a reviewed source. Every acquisition records publisher, URL, retrieval time, SHA-256 digest, length, status and source purpose in a sibling receipt. Without a pinned hash the receipt remains `unreviewed`; with a matching hash it is `hash_matched_not_legal_evidence`. Both states are **insufficient** to certify regulatory claims.

BCAT and NAD require explicit `-allow-large`, have a bounded 4 GiB response limit, and should be acquired only when object-storage capacity and source terms have been reviewed. Downloads are not performed by builds or tests.

## Colorado geometry transformation

The `normalize-colorado` command requires the exact digest-bearing acquisition receipt, parses DOLA's documented `id`, `city`, and `the_geom` columns, validates IDs and MultiPolygon rings, and emits GeoJSON features with `status=geographic_observation_only` and `source_sha256`. It refuses a mismatched input digest, duplicate IDs, malformed geometry and output overwrites. It does not infer incorporation dates, historical annexations, official legal boundaries or code adoption.

An output should be compared against Census TIGER, dated DOLA annexation records, and local official boundaries before the existing boundary snapshot process can admit it. Adoption claims additionally require statutes, regulations or ordinances, exact evidence and effective dates. No connector bypasses those gates.

## Next promotion gates

1. Independently acquire each original file and confirm the endpoint, format, schema, version and redistribution conditions. This session verified public landing pages and source metadata, **not successful end-to-end source downloads**.
2. Implement source-specific normalization for DOE, FEMA and NAD based on inspected bytes, retaining research-candidate status.
3. Place retained, verified input objects and rights/provenance records under approved private custody.
4. Reuse the canonical provider and runtime bundle adapters once the ongoing [source-ownership migration](../docs/source-ownership.md) is complete. Test fail-closed behavior on contradictory and missing official evidence before any real-world `verified` outcome.

Useful source pages: [DOE state portal](https://www.energycodes.gov/state-portal), [FEMA BCAT extract](https://zenodo.org/records/21084129), [USDOT NAD](https://data.transportation.gov/dataset/National-Address-Database-NAD-Text-File/fc2s-wawr), [DOLA municipality dataset](https://data.colorado.gov/Local-Aggregation/Municipal-Boundaries-in-Colorado/w6jq-7em5), [Census TIGER/Line](https://www.census.gov/geographies/mapping-files/time-series/geo/tiger-line-file.html).
