#!/usr/bin/env python3
"""Generate reproducible equal-area land-intersecting grid samples from pinned GeoJSON.

Requires shapely==2.1.2 and pyproj==3.7.2. Input must be an authorized
WGS84 land polygon, not unreviewed jurisdiction or code-adoption evidence.
"""
from __future__ import annotations

import argparse
from hashlib import sha256
import json
import math
from pathlib import Path

METRES_PER_MILE = 1609.344
CRS = "EPSG:6933"


def load_polygon(path: Path, expected_sha256: str):
    from shapely.geometry import shape
    from shapely.ops import unary_union

    content = path.read_bytes()
    actual = sha256(content).hexdigest()
    if len(expected_sha256) != 64 or expected_sha256.lower() != actual:
        raise ValueError("geometry SHA-256 mismatch")
    doc = json.loads(content)
    if "crs" in doc:
        raise ValueError("GeoJSON must use RFC 7946 WGS84, without a custom crs")
    if doc.get("type") == "FeatureCollection":
        features = doc.get("features", [])
        geometries = [feature["geometry"] for feature in features]
    elif doc.get("type") == "Feature":
        geometries = [doc["geometry"]]
    else:
        geometries = [doc]
    if not geometries:
        raise ValueError("empty land geometry")
    polygons = []
    for item in geometries:
        geom = shape(item)
        if geom.geom_type not in ("Polygon", "MultiPolygon") or geom.is_empty or not geom.is_valid:
            raise ValueError("geometry must contain valid land Polygon/MultiPolygon features")
        # A geographic span over 180 degrees can indicate an unsplit dateline
        # crossing. Requiring preprocessing avoids an accidental world-width grid.
        if geom.bounds[2] - geom.bounds[0] > 180:
            raise ValueError("unsplit antimeridian polygon; normalize the source first")
        polygons.append(geom)
    return unary_union(polygons), actual


def generate(geometry, source_sha256: str, miles: float):
    from pyproj import Transformer
    from shapely.geometry import box
    from shapely.ops import transform

    if not math.isfinite(miles) or miles <= 0 or miles > 1000:
        raise ValueError("spacing must be greater than zero and at most 1000 miles")
    step = miles * METRES_PER_MILE
    forward = Transformer.from_crs("EPSG:4326", CRS, always_xy=True).transform
    inverse = Transformer.from_crs(CRS, "EPSG:4326", always_xy=True).transform
    land = transform(forward, geometry)
    minx, miny, maxx, maxy = land.bounds
    first_column, last_column = math.floor(minx / step), math.ceil(maxx / step)
    first_row, last_row = math.floor(miny / step), math.ceil(maxy / step)
    # Column/row and anchored projected origin, not feature ordering, determine IDs.
    for row in range(first_row, last_row):
        for column in range(first_column, last_column):
            cell = box(column * step, row * step, (column + 1) * step, (row + 1) * step)
            if not land.intersects(cell):
                continue
            overlap = land.intersection(cell)
            if overlap.is_empty or overlap.area <= 0:
                continue
            midpoint = cell.centroid
            representative = midpoint if land.covers(midpoint) else overlap.representative_point()
            lon, lat = inverse(representative.x, representative.y)
            if not (-180 <= lon <= 180 and -90 <= lat <= 90):
                raise ValueError("inverse coordinate transform produced an invalid point")
            yield {
                "id": f"grid:{miles:g}mi:r{row}:c{column}",
                "point": {"longitude": round(lon, 7), "latitude": round(lat, 7)},
                "source_sha256": source_sha256,
                "cell": {
                    "crs": CRS,
                    "column": column,
                    "row": row,
                    "step_metres": round(step, 6),
                    "land_fraction": round(min(1.0, overlap.area / cell.area), 8),
                    "sample_kind": "cell_center" if land.covers(midpoint) else "land_representative_point",
                },
            }


def main(argv=None):
    p = argparse.ArgumentParser(description=__doc__)
    p.add_argument("--geometry", type=Path, required=True, help="source WGS84 land GeoJSON")
    p.add_argument("--sha256", required=True, help="exact SHA-256 of original geometry bytes")
    p.add_argument("--miles", type=float, default=100.0, help="equal-area cell side, default 100")
    p.add_argument("--output", type=Path, required=True, help="JSONL output")
    args = p.parse_args(argv)
    geometry, checksum = load_polygon(args.geometry, args.sha256)
    # Avoid partially valid sample files after an invalid input or projection.
    target = args.output.with_name(args.output.name + ".tmp")
    try:
        count = 0
        with target.open("w", encoding="utf-8") as handle:
            for row in generate(geometry, checksum, args.miles):
                handle.write(json.dumps(row, sort_keys=True, separators=(",", ":")) + "\n")
                count += 1
        if not count:
            raise ValueError("no positive-area land cells in geometry")
        target.replace(args.output)
    finally:
        target.unlink(missing_ok=True)
    print(json.dumps({"samples": count, "geometry_sha256": checksum, "grid_crs": CRS,
                      "miles": args.miles, "output": str(args.output)}, sort_keys=True))


if __name__ == "__main__":
    main()
