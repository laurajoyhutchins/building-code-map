# EIA utility context (non-regulatory)

BCM's optional county context reader consumes a **content-pinned projection** of an independently verified [Electric Utility Catalog](https://github.com/laurajoyhutchins/electric-utility-catalog) SQLite release. It does not consume the regulatory Provider, geocode addresses, change BCM jurisdiction selection, or contribute to `ProjectCodeBasis.verdict`.

The EIA Form 861 respondent identifier is a statistical reporting ID. A respondent may be a marketer or administrator, rather than a physical distribution utility. A reported county association is only a **candidate**: it does not establish the retail utility at a parcel, a facility's point of connection, a franchise boundary, or any electrical service characteristics. Absence from the candidate list is not proof that nobody serves the county. Confirmation requires separately identified utility documents or project-site evidence.

## Ownership migration to Electric Utility Catalog

The source-independent EIA projection and standalone county API now have their
target implementation in [Electric Utility Catalog](https://github.com/laurajoyhutchins/electric-utility-catalog/blob/main/docs/county-candidate-api.md).
BCM is a **consumer** only: it establishes a county using source-bound
geography and then reads the independently published candidate list.

For a separately reviewed and deployed catalog service, configure the BCM
live backend (both values are mandatory):

```sh
BCM_UTILITY_CATALOG_API_URL=https://catalog-service.example.com
BCM_UTILITY_CATALOG_VIEW_SHA256=<exact-operator-reviewed-64-character-sha256>
```

The origin must be HTTPS and has no path, query, user info or fragment.
The remote adapter rejects redirects, malformed/oversize payloads, mismatched
state or county, changed release/view provenance, duplicate EIA IDs, unexpected
fields and promotional serving-utility claims. It never sends the project
address to the catalog service. Missing or inconsistent catalog results
return `503 source_unavailable` independently of the code-basis verdict.

The original `BCM_UTILITY_CONTEXT_VIEW_PATH` and
`BCM_UTILITY_CONTEXT_VIEW_SHA256` are **transitional rollback settings**.
BCM refuses to start with both local and remote inputs configured, rather than
silently combining two EIA releases. The standalone BCM image, local exporter,
and local catalog reader must not be removed until live catalog deployment,
pinned source readback, UI parity and rollback have been independently shown.
A merged PR alone does not activate the remote service.

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

## Geographic county evidence from the live BCM backend

When the **existing** authenticated private provider is configured and the
checksum-pinned EIA projection is loaded, the public Go API additionally
exposes `POST /v1/project-utility-context` with bounded JSON input:

```json
{"project_id":"SITE-1","address":"100 Main St, Denver, CO"}
```

The bridge invokes the private engine's existing `/v1/geocode` and
`/v1/lookup` operations, only after `/v1/readiness` and the exact provider
bundle identity check. It requires one selected **address point** rather than a
street-range interpolation, an unambiguous county boundary with source IDs,
matched Census state FIPS, identical returned coordinates, and matching pinned
private bundle provenance. It preserves boundary/geocoder source digests in
`geography` and nests the original non-authoritative EIA county result under
`utility_context`. The input address is not echoed back. Missing, ambiguous,
unsupported, or contradictory geography returns an error, not a utility guess.

This is **not** a second regulatory decision path, a public geocoder database,
or an address-level service-territory claim. It is disabled in DEMO-only mode,
and it is not deployed by merging code. Production activation still requires
private source and deployment verification and a separate user-interface review.

## Project Utility Context in the BCM user interface

The separately labeled Utility Context section is available **only in live
API mode**. The browser sends the same project ID and address as the user's
code-basis request to the existing `POST /v1/project-utility-context` operation.
BCM's server-side bridge remains solely responsible for authenticating its
private provider, geocoding to a unique address point, checking a source-bound
county boundary, and reading the pinned EIA county-candidate view. The browser
never derives state/county from entered address text and it never makes a
jurisdiction or serving-utility determination.

The code-basis result and utility context are independent. If the utility
lookup is unavailable, unresolved, or ambiguous, its section shows that
limitation without replacing the code verdict. Both views invalidate when any
project input changes. Late HTTP responses from a previous project are ignored.
The live section shows source IDs, the private provider bundle identity, EIA
reporting year/release/view hash, and a capped candidate list with a clear
non-confirmation notice. A candidate is **not** a verified service connection;
no UI control asserts it is. The static Pages demo never issues this request.

Enabling the section is not a production deployment. It remains inactive until
the configured Go backend has its independently verified provider and EIA view
installed, and public Pages is explicitly switched to that live backend.
