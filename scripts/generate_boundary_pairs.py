#!/usr/bin/env python3
"""Generate test points across *source* polygon edges, never legal boundaries.

Requires shapely 2.1 and pyproj 3.7. Does not admit legal authority or codes.
"""
from __future__ import annotations

import argparse
from hashlib import sha256
import json
import math
from pathlib import Path
import re

from pyproj import Transformer
from shapely.geometry import Point, shape
from shapely.ops import transform
from shapely.strtree import STRtree

CRS = "EPSG:6933"
IDENT = re.compile(r"^[A-Za-z0-9][A-Za-z0-9._:-]{0,100}$")


def load_features(path: Path, expected_sha256: str, id_field: str = "id", *, allow_integer_ids: bool = False):
    raw = path.read_bytes()
    digest = sha256(raw).hexdigest()
    if not re.fullmatch(r"[a-f0-9]{64}", expected_sha256) or digest != expected_sha256:
        raise ValueError("source boundary geometry SHA-256 mismatch")
    document = json.loads(raw)
    if not isinstance(document, dict) or document.get("type") != "FeatureCollection" or "crs" in document:
        raise ValueError("source must be a WGS84 RFC 7946 FeatureCollection")
    if not IDENT.fullmatch(id_field):
        raise ValueError("unsafe source feature identifier field")
    entries = document.get("features", [])
    if not isinstance(entries, list) or len(entries) > 100000:
        raise ValueError("invalid or too many source features")
    forward = Transformer.from_crs("EPSG:4326", CRS, always_xy=True).transform
    features = {}
    for feature in entries:
        if not isinstance(feature, dict) or feature.get("type") != "Feature":
            raise ValueError("expected GeoJSON Feature")
        properties = feature.get("properties")
        if not isinstance(properties, dict):
            raise ValueError("source feature lacks properties")
        ident = properties.get(id_field)
        if allow_integer_ids and type(ident) is int and ident > 0:
            ident = str(ident)
        if not isinstance(ident, str) or not IDENT.fullmatch(ident) or ident in features:
            raise ValueError("feature IDs must be safe, unique strings")
        geom = shape(feature.get("geometry"))
        if geom.geom_type not in ("Polygon", "MultiPolygon") or geom.is_empty or not geom.is_valid:
            raise ValueError("invalid polygon geometry")
        parts = list(geom.geoms) if geom.geom_type == "MultiPolygon" else [geom]
        if any(part.bounds[2] - part.bounds[0] > 180 for part in parts):
            raise ValueError("unsplit antimeridian geometry")
        projected = transform(forward, geom)
        if projected.is_empty or not projected.is_valid or projected.area <= 0:
            raise ValueError("invalid projected polygon")
        features[ident] = projected
    if len(features) < 2:
        raise ValueError("at least two source polygons required")
    return dict(sorted(features.items())), digest


def shared_lines(geometry):
    if geometry.geom_type in ("LineString", "LinearRing"):
        return [geometry] if geometry.length > 0 else []
    if geometry.geom_type in ("MultiLineString", "GeometryCollection"):
        return [line for part in geometry.geoms for line in shared_lines(part)]
    return []


def candidate_pair(a, b, i, j, all_geometries, tree, line, offset, inverse, forward):
    # Evaluate both sides by exact geometric membership, never by ring winding.
    centre = line.interpolate(0.5, normalized=True)
    tangent = min(10, line.length / 4)
    before = line.interpolate(max(0, line.length / 2 - tangent))
    after = line.interpolate(min(line.length, line.length / 2 + tangent))
    dx, dy = after.x - before.x, after.y - before.y
    norm = math.hypot(dx, dy)
    if not norm:
        return None
    nx, ny = -dy / norm, dx / norm
    for distance in (offset, offset / 2, offset / 4, offset / 8):
        points = [Point(centre.x + sign * distance * nx, centre.y + sign * distance * ny)
                  for sign in (1, -1)]
        if not ((a.contains(points[0]) and b.contains(points[1])) or
                (a.contains(points[1]) and b.contains(points[0]))):
            continue
        checked = []
        for point in points:
            lon, lat = inverse(point.x, point.y)
            if not (-180 <= lon <= 180 and -90 <= lat <= 90):
                break
            coordinate = {"longitude": round(lon, 7), "latitude": round(lat, 7)}
            check = Point(*forward(coordinate["longitude"], coordinate["latitude"]))
            members = {int(k) for k in tree.query(check)
                       if all_geometries[int(k)].contains(check)}
            checked.append((coordinate, members))
        if len(checked) != 2:
            continue
        for first, second, local_first in ((i, j, 0), (j, i, 1)):
            if checked[0][1] == {first} and checked[1][1] == {second}:
                return checked[0][0], checked[1][0], round(distance, 6), local_first
    return None


def generate(features: dict, digest: str, offset_metres: float, max_pairs: int = 10000):
    if not math.isfinite(offset_metres) or not 1 <= offset_metres <= 5000:
        raise ValueError("offset must be 1..5000 projected metres")
    if not 1 <= max_pairs <= 100000:
        raise ValueError("max pairs must be 1..100000")
    names = list(features)
    polygons = [features[name] for name in names]
    tree = STRtree(polygons)
    inverse = Transformer.from_crs(CRS, "EPSG:4326", always_xy=True).transform
    forward = Transformer.from_crs("EPSG:4326", CRS, always_xy=True).transform
    report = {"schema_version": "1", "source_geometry_sha256": digest,
              "source_features": len(features), "candidate_touch_pairs": 0,
              "sampled_pairs": 0, "unresolved_touch_pairs": 0,
              "overlapping_source_pairs": 0, "max_pairs": max_pairs,
              "requested_offset_metres": offset_metres,
              "classification": "source_polygon_candidate_only"}
    rows = []
    for i, first in enumerate(polygons):
        for other in tree.query(first):
            j = int(other)
            if j <= i:
                continue
            second = polygons[j]
            if first.intersection(second).area > 0:
                report["overlapping_source_pairs"] += 1
                continue
            lines = shared_lines(first.boundary.intersection(second.boundary))
            if not lines:
                continue
            report["candidate_touch_pairs"] += 1
            if report["candidate_touch_pairs"] > max_pairs:
                raise ValueError("source exceeds --max-pairs; refuse silent truncation")
            found = None
            for line in sorted(lines, key=lambda part: -part.length):
                found = candidate_pair(first, second, i, j, polygons, tree, line,
                                       offset_metres, inverse, forward)
                if found:
                    break
            if found is None:
                report["unresolved_touch_pairs"] += 1
                continue
            p0, p1, distance, first_side = found
            pair_hash = sha256((names[i] + "\0" + names[j]).encode()).hexdigest()[:24]
            for point, local_side in ((p0, first_side), (p1, 1 - first_side)):
                side = "a" if local_side == 0 else "b"
                rows.append({"id": "edge:" + pair_hash + ":" + side,
                             "point": point, "source_sha256": digest,
                             "boundary": {"pair_key": pair_hash,
                                          "feature_id": names[(i, j)[local_side]],
                                          "neighbor_feature_id": names[(i, j)[1-local_side]],
                                          "side": side,
                                          "offset_metres": distance,
                                          "crs": CRS,
                                          "classification": "source_polygon_candidate_only"}})
            report["sampled_pairs"] += 1
    report["samples"] = len(rows)
    return sorted(rows, key=lambda row: row["id"]), report


def main(argv=None):
    p = argparse.ArgumentParser(description=__doc__)
    p.add_argument("--geometry", type=Path, required=True)
    p.add_argument("--sha256", required=True)
    p.add_argument("--id-field", default="id")
    p.add_argument("--offset-metres", type=float, default=100)
    p.add_argument("--max-pairs", type=int, default=10000)
    p.add_argument("--output", type=Path, required=True)
    p.add_argument("--summary", type=Path, required=True)
    args = p.parse_args(argv)
    if len({args.geometry.resolve(), args.output.resolve(), args.summary.resolve()}) != 3:
        p.error("source, output, summary must be different files")
    features, digest = load_features(args.geometry, args.sha256, args.id_field)
    rows, report = generate(features, digest, args.offset_metres, args.max_pairs)
    if not rows:
        raise ValueError("no unambiguous positive-length shared-edge pairs were sampled")
    import os
    from scripts.bulk_query import atomic_json
    target = args.output.with_name(args.output.name + ".tmp")
    try:
        with target.open("w", encoding="utf-8") as handle:
            for row in rows:
                handle.write(json.dumps(row, sort_keys=True, separators=(",", ":")) + "\n")
            handle.flush()
            os.fsync(handle.fileno())
        report["workload_sha256"] = sha256(target.read_bytes()).hexdigest()
        target.replace(args.output)
        atomic_json(args.summary, report)
    finally:
        target.unlink(missing_ok=True)
    print(json.dumps(report, sort_keys=True))


if __name__ == "__main__":
    main()
