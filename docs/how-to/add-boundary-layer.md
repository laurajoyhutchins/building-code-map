# Add a Boundary Layer

## Goal

Add a new TIGERweb boundary family to the frontend and backend scaffold.

## Prerequisites

- You know the boundary family key and Martin layer id.
- The backend schema and mock data can be updated together.
- The frontend layer registry can be updated in the same change.

## Steps

1. Add the new layer family to the backend schema and mock data.
2. Add the same layer family to the frontend layer registry and default selection map.
3. Update tests that assert the supported layer count or the default-enabled set.
4. Confirm the backend returns the new family from `GET /layers`.
5. Verify the map shows the toggle and the layer can be enabled or disabled.

## Result

The new boundary family appears in the layer control and can participate in feature lookup and map rendering.
