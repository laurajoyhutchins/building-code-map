#!/usr/bin/env python3
"""Audit source-polygon sample omissions and topology voids; not legal gaps."""
from __future__ import annotations

import argparse
from collections import Counter
from hashlib import sha256
import json
from pathlib import Path

from pyproj import Transformer
from shapely.geometry import Point, Polygon
from shapely.strtree import STRtree

from scripts.bulk_query import atomic_json, validate_sample
from scripts.generate_boundary_pairs import CRS, load_features


def load_baseline(path: Path, expected_sha: str, polygons, tree):
    raw = path.read_bytes()
    digest = sha256(raw).hexdigest()
    if digest != expected_sha or len(expected_sha) != 64:
        raise ValueError("baseline workload SHA-256 mismatch")
    if len(raw) > 32 * 1024 * 1024:
        raise ValueError("baseline exceeds 32 MiB")
    forward = Transformer.from_crs("EPSG:4326", CRS, always_xy=True).transform
    covered = set()
    seen = set()
    total = 0
    for i, line in enumerate(raw.splitlines(), 1):
        if not line.strip():
            continue
        try:
            row = validate_sample(json.loads(line))
        except (TypeError, ValueError, KeyError) as error:
            raise ValueError(f"invalid baseline line {i}") from error
        if row["id"] in seen:
            raise ValueError("duplicate baseline sample ID")
        if "point" not in row:
            raise ValueError("baseline for polygon audit must contain geographic points")
        seen.add(row["id"])
        total += 1
        point = Point(*forward(row["point"]["longitude"], row["point"]["latitude"]))
        covered.update(int(k) for k in tree.query(point) if polygons[int(k)].contains(point))
    if not total:
        raise ValueError("baseline is empty")
    return covered, total, digest


def generate(features, source_digest: str, baseline: Path, expected_baseline: str,
             max_samples: int = 10000):
    if not 1 <= max_samples <= 100000:
        raise ValueError("max samples must be 1..100000")
    names = list(features)
    polygons = [features[name] for name in names]
    tree = STRtree(polygons)
    covered, base_total, baseline_sha = load_baseline(baseline, expected_baseline, polygons, tree)
    inverse = Transformer.from_crs(CRS, "EPSG:4326", always_xy=True).transform
    forward = Transformer.from_crs("EPSG:4326", CRS, always_xy=True).transform
    counts = Counter()
    rows = []

    def memberships(candidate):
        lon, lat = inverse(candidate.x, candidate.y)
        coordinate = {"longitude": round(lon, 7), "latitude": round(lat, 7)}
        if not (-180 <= coordinate["longitude"] <= 180 and -90 <= coordinate["latitude"] <= 90):
            return None, set()
        point = Point(*forward(coordinate["longitude"], coordinate["latitude"]))
        return coordinate, {int(k) for k in tree.query(point) if polygons[int(k)].contains(point)}

    def add(kind, key, point, candidates, represented):
        if len(rows) >= max_samples:
            raise ValueError("source audit exceeds max samples; refusing truncation")
        rows.append({
            "id": "gap:" + sha256((kind + "\0" + key).encode()).hexdigest()[:24],
            "point": point, "source_sha256": source_digest,
            "gap": {"classification": "source_geometry_candidate_only",
                    "kind": kind, "source_feature_id": key,
                    "containing_feature_ids": sorted(candidates),
                    "baseline_sampled": represented,
                    "crs": CRS},
        })
        counts[kind] += 1

    for i, name in enumerate(names):
        geom = polygons[i]
        if i not in covered:
            point, member_indices = memberships(geom.representative_point())
            if point is None or i not in member_indices:
                counts["unsampleable_feature"] += 1
            else:
                partial_overlaps = [j for j in member_indices if j != i and not polygons[j].covers(geom)]
                if partial_overlaps:
                    counts["ambiguous_overlap_skipped"] += 1
                else:
                    nested = any(j != i for j in member_indices)
                    kind = "nested_feature_candidate" if nested else "unrepresented_feature_candidate"
                    add(kind, name, point, [names[j] for j in member_indices], False)
        # An interior ring is a geometry hole, not evidence of a legal void.
        parts = list(geom.geoms) if geom.geom_type == "MultiPolygon" else [geom]
        for part_index, part in enumerate(parts):
            for ring_index, ring in enumerate(part.interiors):
                hole = Polygon(ring)
                if hole.is_empty or not hole.is_valid or hole.area <= 0:
                    counts["invalid_holes_skipped"] += 1
                    continue
                point, memberships_at_hole = memberships(hole.representative_point())
                if point is None or memberships_at_hole:
                    counts["occupied_or_unsampleable_holes"] += 1
                    continue
                add("interior_void_candidate", f"{name}:{part_index}:{ring_index}",
                    point, [], False)

    report = {
        "schema_version": "1", "classification": "source_geometry_candidate_only",
        "source_geometry_sha256": source_digest, "baseline_sha256": baseline_sha,
        "source_features": len(polygons), "baseline_samples": base_total,
        "baseline_represented_features": len(covered),
        "missing_source_features": len(polygons) - len(covered),
        "generated_samples": len(rows), "categories": dict(sorted(counts.items())),
    }
    return sorted(rows, key=lambda row: row["id"]), report


def main(argv=None):
    p = argparse.ArgumentParser(description=__doc__)
    p.add_argument("--geometry", type=Path, required=True)
    p.add_argument("--sha256", required=True)
    p.add_argument("--id-field", default="id")
    p.add_argument("--baseline", type=Path, required=True)
    p.add_argument("--baseline-sha256", required=True)
    p.add_argument("--output", type=Path, required=True)
    p.add_argument("--summary", type=Path, required=True)
    p.add_argument("--max-samples", type=int, default=10000)
    args = p.parse_args(argv)
    paths = [x.resolve() for x in (args.geometry, args.baseline, args.output, args.summary)]
    if len(set(paths)) != 4:
        p.error("inputs and outputs must be four different paths")
    features, digest = load_features(args.geometry, args.sha256, args.id_field)
    rows, receipt = generate(features, digest, args.baseline, args.baseline_sha256,
                             args.max_samples)
    import os
    temp = args.output.with_name(args.output.name + ".tmp")
    try:
        with temp.open("w", encoding="utf-8") as fh:
            for row in rows:
                fh.write(json.dumps(row, sort_keys=True, separators=(",", ":")) + "\n")
            fh.flush()
            os.fsync(fh.fileno())
        receipt["workload_sha256"] = sha256(temp.read_bytes()).hexdigest()
        temp.replace(args.output)
        atomic_json(args.summary, receipt)
    finally:
        temp.unlink(missing_ok=True)
    print(json.dumps(receipt, sort_keys=True))


if __name__ == "__main__":
    main()
