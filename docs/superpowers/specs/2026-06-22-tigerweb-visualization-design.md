# TIGERweb Visualization Design

**Date:** 2026-06-22

**Goal:** Let a user explore AHJ-relevant TIGERweb boundary data on an OpenStreetMap basemap, click a feature, and inspect the full TIGERweb attribute record.

**Scope:** Frontend in TypeScript with MapLibre GL JS. Boundary data mirrored into PostGIS, served as vector tiles by Martin, and accessed through a FastAPI backend for detail lookup and refresh orchestration.

---

## Product Summary

This product is a map-first explorer for TIGERweb boundary data. The initial experience should be simple: show boundaries, let the user toggle layers, and let them click a feature to inspect the full record. The application is intended as a springboard for future GIS and AHJ workflow work, so the architecture should be easy to extend without being overbuilt now.

## In Scope

- OpenStreetMap basemap.
- Vector-tile rendering of TIGERweb boundary layers.
- Boundary layers relevant to AHJs:
  - States
  - Counties
  - Municipalities / incorporated places
  - Military and other special land use areas
  - American Indian areas
- Layer toggles for turning boundary families on and off.
- Click-to-identify behavior for boundary features.
- Attribute panel showing the full TIGERweb record for the clicked feature.
- Scheduled refresh of mirrored TIGERweb data into PostGIS.
- Basic loading, empty, and error states.

## Deferred

These items are explicitly out of scope for the MVP:

- Search.
- Geocoding.
- Historical versioning.
- Editing or data correction workflows.
- Authentication and user accounts.
- Advanced analytics, reporting, or exports.
- Curated summaries of attributes.

## User Experience

1. The user opens the app and sees an OpenStreetMap basemap with TIGERweb boundary overlays.
2. A layer control lets the user show or hide each boundary family.
3. When the user clicks a boundary feature, the feature highlights.
4. A side panel or drawer opens with the complete TIGERweb record for that feature.
5. The user can click a different feature and the panel updates.

The UI should be functional and readable, not fancy. The main design goal is clarity and a stable base for future work.

## Frontend

### Responsibilities

- Render the map with MapLibre GL JS.
- Load the OSM basemap and Martin vector tile sources.
- Present layer toggles for the supported TIGERweb boundary families.
- Handle feature clicks, highlight state, and attribute display.
- Request full feature details from the backend when a feature is selected.

### Behavior

- The map starts with a sensible default set of boundary layers enabled.
- Each layer can be toggled independently.
- Clicking a feature should be forgiving: if multiple features overlap, the topmost or nearest visible feature may be selected, but the selection must be deterministic.
- The attribute panel should render the complete record in a straightforward field/value layout.
- Raw TIGERweb field names should be preserved.
- Null or empty values should still be represented clearly so the user can tell whether a field is absent or just not populated.

### Frontend Data Contract

The frontend should not depend on tile internals beyond:

- A stable feature identifier.
- A layer family identifier.
- A geometry suitable for highlight/select behavior.

The detail panel should be fed by a backend lookup keyed by the selected feature identifier rather than by tile payload alone.

## Backend

### Responsibilities

- Serve application metadata and layer configuration.
- Resolve clicked feature identifiers to full TIGERweb records.
- Expose refresh status and operational health.
- Orchestrate or report on scheduled ingestion into PostGIS.

### Suggested Endpoints

- `GET /health` for basic liveness.
- `GET /layers` for the supported layer families and display metadata.
- `GET /features/{layer_family}/{feature_id}` for the full TIGERweb record.
- `GET /refresh/status` for the most recent refresh outcome and timestamps.

The exact route names may change, but the backend must provide these capabilities.

## Data Model

### Canonical Storage

PostGIS is the system of record for the mirrored TIGERweb data. The mirror should preserve the original TIGERweb fields, not just a simplified subset.

Recommended storage shape:

- A canonical feature table or set of tables keyed by layer family.
- Stable internal feature identifiers.
- Original TIGERweb source identifiers.
- Geometry in PostGIS geometry columns.
- Original attributes preserved in columns and/or JSONB so the full record can be reconstructed.
- Timestamps for ingestion and refresh tracking.

### Indexing

The stored data should be indexed for the two primary access patterns:

- Tile generation by geometry and layer family.
- Point lookup by feature identifier for the detail endpoint.

## Tile Serving

Martin should read from PostGIS and serve vector tiles for the supported boundary layers.

Requirements:

- Tiles must be generated from the mirrored PostGIS data, not directly from TIGERweb at request time.
- Layer names in Martin should map cleanly to the frontend layer registry.
- Tile payloads should be small enough for responsive map interaction.
- Tile feature identifiers must be stable enough to support click selection and backend lookup.

The tile layer is responsible for map rendering, not for showing the full attribute record.

## Refresh Pipeline

TIGERweb data should be refreshed on a schedule.

Requirements:

- Refreshes should update the mirrored PostGIS data automatically.
- Refreshes should not require manual map or tile editing.
- Refreshes should avoid leaving the system in a partially updated state.
- The latest successful refresh time should be visible through backend status metadata.

Implementation can use a scheduled job, a worker, or a cron-driven process, as long as the result is a repeatable scheduled refresh.

## Error Handling

The app should handle the following cases gracefully:

- Basemap unavailable.
- Tile source unavailable.
- Feature detail lookup fails.
- Refresh data is stale or currently failing.
- No feature is selected.

Behavior expectations:

- Display a clear, non-blocking error message when a layer fails.
- Keep the map usable even if one boundary family is unavailable.
- Show a loading state in the attribute panel while details are being fetched.
- If the clicked feature cannot be resolved, keep the selection visible and explain that details are unavailable.

## Non-Goals and Constraints

- No search or geocoding in the MVP.
- No historical snapshots or time travel.
- No statistical-area layers.
- No custom cartographic styling beyond basic legibility.
- No data editing or admin UI.
- No attempt to normalize TIGERweb into a heavily curated domain model.

## Acceptance Criteria

- The app loads an OpenStreetMap basemap and can display the chosen TIGERweb boundary families.
- A user can toggle states, counties, municipalities, special land use areas, and American Indian areas on and off.
- A user can click a feature and see the full TIGERweb attribute record.
- The detail panel shows complete raw fields, not only curated summaries.
- TIGERweb data is mirrored into PostGIS and refreshed on a schedule.
- Martin serves the vector tiles from the mirrored PostGIS data.
- The app remains usable when one data source or tile layer fails.

## Open Questions

- Refresh cadence is intentionally left configurable; the system only requires scheduled refresh, not a specific interval.
- The exact shape of the PostGIS tables can be chosen during implementation as long as the full record and stable feature lookup are preserved.
- The exact UI placement of the attribute panel can be chosen during implementation, provided it is easy to find and easy to dismiss.

