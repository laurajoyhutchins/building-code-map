#!/usr/bin/env python3
"""BCM source-only point/polygon candidates, never jurisdiction or code authority.

Reuses the existing Shapely/PyProj polygon loader and STRtree, including strict
geometry, source SHA, CRS and antimeridian checks. Not a second ProjectVerifier.
"""
from __future__ import annotations

import argparse
import json
import math
from pathlib import Path

from pyproj import Transformer
from shapely.geometry import Point
from shapely.strtree import STRtree

from scripts.generate_boundary_pairs import CRS, load_features

ROLES = {"municipal_boundary", "county_boundary", "fire_district", "zoning", "other"}
MAX_BYTES = 128 * 1024 * 1024


def query(geometry_path: Path, trusted_sha256: str, *,
          longitude: float, latitude: float, role: str,
          id_field: str = "OBJECTID", expect_features: int):
    if not (type(longitude) in (int, float) and type(latitude) in (int, float)
            and math.isfinite(longitude) and math.isfinite(latitude)
            and -180 <= longitude <= 180 and -90 <= latitude <= 90):
        raise ValueError("invalid WGS84 longitude or latitude")
    if role not in ROLES:
        raise ValueError("explicit published source layer role required")
    if not 1 <= expect_features <= 100000:
        raise ValueError("complete expected source feature count required")
    if not geometry_path.is_file() or geometry_path.stat().st_size > MAX_BYTES:
        raise ValueError("source geometry unavailable or exceeds 128 MiB")
    features, digest = load_features(
        geometry_path, trusted_sha256, id_field, allow_integer_ids=True)
    if len(features) != expect_features:
        raise ValueError("complete independently pinned feature count differs")
    ids = list(features)
    polygons = [features[key] for key in ids]
    tree = STRtree(polygons)
    forward = Transformer.from_crs("EPSG:4326", CRS, always_xy=True).transform
    point = Point(*forward(longitude, latitude))
    near = [int(i) for i in tree.query(point)]
    inside = sorted(ids[i] for i in near if polygons[i].contains(point))
    boundary = sorted(ids[i] for i in near if polygons[i].touches(point))
    status = ("boundary_uncertain" if boundary else
              "multiple_source_candidates" if len(inside) > 1 else
              "one_source_candidate" if len(inside) == 1 else
              "no_source_polygon_match")
    return {
        "schema_version": "1.0.0",
        "classification": "source_polygon_candidate_only",
        "source_geometry_sha256": digest,
        "source_layer_role": role,
        "source_features": len(features),
        "point": {"longitude": longitude, "latitude": latitude},
        "candidate_feature_ids": inside,
        "boundary_feature_ids": boundary,
        "status": status,
        "identity_review": "not_established",
        "boundary_effective_date": "unknown",
        "legal_authority": "not_established",
        "regulatory_verdict_authorized": False,
        "limitations": "Source polygon containment alone cannot identify a legally enforcing government, permit authority or adopted code.",
    }


def main(argv=None):
    p = argparse.ArgumentParser(description=__doc__)
    p.add_argument("--geometry", type=Path, required=True)
    p.add_argument("--sha256", required=True)
    p.add_argument("--expected-features", type=int, required=True)
    p.add_argument("--id-field", default="OBJECTID")
    p.add_argument("--source-role", required=True, choices=sorted(ROLES))
    p.add_argument("--longitude", type=float, required=True)
    p.add_argument("--latitude", type=float, required=True)
    options = p.parse_args(argv)
    result = query(options.geometry, options.sha256,
                   longitude=options.longitude, latitude=options.latitude,
                   role=options.source_role, id_field=options.id_field,
                   expect_features=options.expected_features)
    print(json.dumps(result, sort_keys=True, allow_nan=False))


if __name__ == "__main__":
    main()
