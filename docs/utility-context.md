# EIA utility context (non-regulatory)

BCM's optional county context reader consumes a **content-pinned projection** of an independently verified [Electric Utility Catalog](https://github.com/laurajoyhutchins/electric-utility-catalog) SQLite release. It does not consume the regulatory Provider, geocode addresses, change BCM jurisdiction selection, or contribute to `ProjectCodeBasis.verdict`.

The EIA Form 861 respondent identifier is a statistical reporting ID. A respondent may be a marketer or administrator, rather than a physical distribution utility. A reported county association is only a **candidate**: it does not establish the retail utility at a parcel, a facility's point of connection, a franchise boundary, or any electrical service characteristics. Absence from the candidate list is not proof that nobody serves the county. Confirmation requires separately identified utility documents or project-site evidence.

## Prepare the read-only projection

Obtain the separately verified `catalog.sqlite` and `manifest.json` from one release in the private GCS bucket. Do not commit source data or the generated projection to this public repository. Run:

```sh
python3 scripts/export_utility_context.py \
  --manifest /secure/eia861-2025/manifest.json \
  --database /secure/eia861-2025/catalog.sqlite \
  --output /secure/bcm/eia861-county-candidates-v1.json
```

The program requires the exact SQLite digest recorded in the manifest, checks SQLite integrity and foreign keys, checks the single annual release against the manifest, enforces the source county count, and preserves EIA county source-sheet/row and state-qualified reported activity markers. It writes a deterministic JSON projection and prints its SHA-256. **The view is a derivative with its own digest**, not a replacement or modification of the immutable EIA manifest. Only deploy a projection after an independently checked catalog release has been admitted. The signed-in operator stages the view with private backend assets or an authorized, generation-pinned storage read; the Go service does not need any GCS API permission.

Configure the live **server process**, not GitHub Pages:

```sh
BCM_UTILITY_CONTEXT_VIEW_PATH=/secure/bcm/eia861-county-candidates-v1.json
BCM_UTILITY_CONTEXT_VIEW_SHA256=<sha256-printed-by-the-exporter>
```

Both settings are required together. The service fails to start if the view is missing, modified, oversized, or structurally incompatible. Utility context can only be served by the BCM live backend, not by the fictional DEMO-XX app. Configuration must be redeployed explicitly for a new catalog release. In particular, do not put GCS credentials or the bulk source data in Pages or embed it in the regulatory provider bundle.

## Separate HTTP operation

```text
GET /v1/utility-context?state=CO&county=Denver
```

The query requires an **already-known state and county name**. It cannot infer a county from a freeform address or coordinates and never treats an address guess as verified geography. Its versioned response contains:

- `status = reported_county_candidates_not_site_verified`
- `state`, `county`, `release_id`, `reporting_year`, `source_sha256`, `sqlite_sha256`, `view_sha256`, `source_url`
- `candidates[]`: EIA ID, reported name, source sheet/row and state-qualified `Y`/`N` activity observations
- an explicit serving-utility disclaimer, including on empty results

This is not `/v1/project-code-basis`. A caller can display a **Utility Context** section once it has verified the county independently, but must not merge these candidates into regulatory authority, code applicability, or the `verified` verdict. The browser must never guess a county from the input address. The selected EIA ID, release ID and source SHA can subsequently be carried into Arcata's separately reviewed project reference, with an independently recorded *serving-utility confirmation* if one exists.

The reference app's static GitHub Pages build stays synthetic. The endpoint is not automatically live just because the feature's tests pass: production staging, authenticating the backend, and an independent deployment readback are separate operational steps.
