# Source polygon candidates in BCM

The public BCM geocoder now has an **offline, source-only** point/polygon probe. It consumes digest-pinned private WGS84 publisher geometry without embedding it in public Git, and uses the same Shapely/PyProj geometry loader already used for boundary-gap and paired-edge stress workloads.

For the Larimer County City Limits layer, the trusted raw publisher candidate collection is a 23-feature GeoJSON `FeatureCollection` with integer `properties.OBJECTID` values. Following generation-pinned readback from the private GCS archive:

    python3 -m scripts.point_source_candidates \
      --geometry /secure/larimer-candidate.geojson \
      --sha256 TRUSTED_EXACT_GEOJSON_SHA256 \
      --expected-features 23 --id-field OBJECTID \
      --source-role municipal_boundary \
      --longitude -105.1 --latitude 40.4

This produces an explicitly **source_polygon_candidate_only** record with zero or more publisher ObjectIDs, separate boundary-touching IDs, and classifications `one_source_candidate`, `multiple_source_candidates`, `boundary_uncertain`, or `no_source_polygon_match`.

No result identifies a reviewed municipality, legal code enforcement authority, adopted edition, effective date, or BCM project verdict. A source exterior gap, polygon hole, overlap or touching boundary must remain an unresolved observation. Multiple disjoint or overlapping polygons sharing one municipal name retain separate ObjectIDs. Source bytes are checked against the **independently obtained SHA-256**, and the full expected feature count is required. Neither a guessed digest nor a manually supplied provenance field establishes private GCS custody.

The query output has a closed JSON Schema + ODCS v3.1.0 contract. It is a reusable geographic observation for BCM, not a competing authoritative resolver. The next consumer must join candidate feature selectors to the private `jurisdiction-identity` source contract and independently admitted geometry-custody receipts. BCM's existing `SiteAuthorityPolicy` and single `ProjectVerifier` remain the only route for accepted site/code conclusions.

All checked-in tests use fictional polygons, and no county geometry is published with the BCM code.
