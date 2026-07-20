# FEMA BCAT

FEMA's National Building Code Adoption Tracking program publishes jurisdiction-level code adoption data and polygons through an ArcGIS FeatureServer. In this repo, BCAT is treated as a validation and enrichment source, not the legal source of truth.

## Endpoints

- Service metadata: `https://services3.arcgis.com/mQsIDpmRUmYltsYD/arcgis/rest/services/Building_Codes/FeatureServer`
- Layer metadata: `https://services3.arcgis.com/mQsIDpmRUmYltsYD/ArcGIS/rest/services/Building_Codes/FeatureServer/0`
- Query endpoint: `https://services3.arcgis.com/mQsIDpmRUmYltsYD/ArcGIS/rest/services/Building_Codes/FeatureServer/0/query`

## Fields Used

- `OBJECTID`
- `Building_Code`
- `Res_Code`
- `State_Building_Code_Link_1`
- `State_Building_Code_Link_2`
- `State_Building_Code_Link_3`
- `BRIC_Code`
- `BuildCode_sum`
- `BuildCode_Index`
- `BuildingCode_Int`
- `ResCode_Int`
- `Shape__Area`
- `Shape__Length`

## Fields Ignored

- Any field that is not needed for provenance, comparison, or the point lookup response.
- All raw values remain preserved in the normalized record payload, even when a field is not elevated into a dedicated column.

## Outputs

- Raw metadata snapshots:
  - `backend/data/raw/fema_bcat/service_metadata.json`
  - `backend/data/raw/fema_bcat/layer_0_metadata.json`
- Raw feature snapshot:
  - `backend/data/raw/fema_bcat/building_codes.geojson`
- Normalized data:
  - `backend/data/processed/fema_bcat/building_codes_normalized.parquet`
- Comparison output:
  - `backend/data/processed/fema_bcat/fema_bcat_tiger_comparison.parquet`
- Human-readable report:
  - `backend/data/processed/fema_bcat/fema_bcat_discrepancy_report.md`

## Known Limitations

- BCAT does not replace primary-source adoption ordinances, state code pages, amendments, effective dates, or AHJ records.
- `No Data` does not mean `No Building Code`.
- `No Building Code` is still not legal confirmation of absence without primary-source review.
- BCAT geometry can be simplified, dissolved, or otherwise different from TIGER-derived jurisdiction polygons.

## Relationship To TIGER

TIGER-derived polygons remain the canonical jurisdiction geometry in this project. BCAT polygons are compared against TIGER to surface:

- geometry that looks TIGER-derived
- geometry that looks TIGER-adjacent
- dissolved or simplified geometry
- geometry that does not plausibly match the TIGER layer

The comparison is intentionally conservative. Disagreement is surfaced for review instead of silently resolved.

## Update Procedure

1. Run the current backend and validation commands from `backend/` as needed for the workflow you are testing.
2. Inspect the raw metadata snapshots and the raw GeoJSON snapshot under `backend/data/raw/fema_bcat/`.
3. Review the normalized parquet output under `backend/data/processed/fema_bcat/`.
4. Compare the BCAT-TIGER report against the current jurisdiction cache.

## Discrepancy Review

- Treat BCAT mismatches as investigation items, not overwrites.
- Confirm adoption status against primary legal sources before changing any canonical code-stack record.
- Escalate cases where BCAT says `No Data`, `No Building Code`, or geometry confidence is low.
- Keep the BCAT output as provenance that can help explain, but not replace, the final legal decision.
