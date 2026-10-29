#!/usr/bin/env python3
"""Audit coarse grid observations against source-state geometry, without legal authority.

The official Census state polygons are statistical/cartographic observations,
not jurisdictional authority, building code adoption or permit offices.
"""
from __future__ import annotations

import argparse
from collections import Counter
from hashlib import sha256
import json
from pathlib import Path

from shapely.geometry import Point, shape

from scripts.bulk_query import validate_sample
from scripts.prepare_census_us_land import STATE_FIPS


def load_states(path: Path, expected_sha256: str, expect_51: bool = False):
    content = path.read_bytes()
    digest = sha256(content).hexdigest()
    if len(expected_sha256) != 64 or digest != expected_sha256.lower():
        raise ValueError("source geometry digest mismatch")
    root = json.loads(content)
    if root.get("type") != "FeatureCollection" or "crs" in root:
        raise ValueError("expected RFC 7946 WGS84 FeatureCollection")
    states = {}
    for feature in root.get("features", []):
        fips = feature.get("properties", {}).get("STATEFP")
        geometry = shape(feature.get("geometry"))
        if (not isinstance(fips, str) or len(fips) != 2 or not fips.isdigit()
                or geometry.geom_type not in ("Polygon", "MultiPolygon")
                or geometry.is_empty or not geometry.is_valid or fips in states):
            raise ValueError("invalid or duplicate source-state feature")
        states[fips] = geometry
    if not states or (expect_51 and set(states) != STATE_FIPS):
        raise ValueError("incomplete source-state geographic inventory")
    return states, digest


def state_representatives(states, digest):
    for fips in sorted(states):
        p = states[fips].representative_point()
        if not states[fips].covers(p):
            raise ValueError("representative point is not on source-state land")
        yield {"id": "state-probe:" + fips,
               "point": {"longitude": round(p.x, 7), "latitude": round(p.y, 7)},
               "source_sha256": digest}


def assess(states, digest, workload_path: Path):
    counts = Counter({fips: 0 for fips in states})
    seen = set()
    missing, overlapping, rows = [], [], 0
    hash_value = sha256()
    with workload_path.open("rb") as handle:
        for line in handle:
            hash_value.update(line)
            if not line.strip():
                continue
            row = validate_sample(json.loads(line))
            if row["id"] in seen:
                raise ValueError("duplicate grid sample ID")
            seen.add(row["id"])
            if "point" not in row:
                raise ValueError("state audit requires a geographic point workload")
            p = Point(row["point"]["longitude"], row["point"]["latitude"])
            matches = [fips for fips in sorted(states) if states[fips].covers(p)]
            if not matches:
                missing.append(row["id"])
            elif len(matches) > 1:
                overlapping.append(row["id"])
            else:
                counts[matches[0]] += 1
            rows += 1
    return {
        "schema_version": "1",
        "classification": "source_geography_candidate_only",
        "source_geometry_sha256": digest,
        "workload_sha256": hash_value.hexdigest(),
        "samples": rows,
        "states": len(states),
        "state_candidate_counts": dict(sorted(counts.items())),
        "unsampled_states": sorted(fips for fips, n in counts.items() if n == 0),
        "unmatched_points": len(missing),
        "overlapping_candidate_points": len(overlapping),
        "unmatched_ids": missing,
        "overlapping_ids": overlapping,
    }


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--geometry", type=Path, required=True)
    parser.add_argument("--sha256", required=True)
    parser.add_argument("--grid", type=Path, required=True)
    parser.add_argument("--audit", type=Path, required=True)
    parser.add_argument("--representatives", type=Path, required=True)
    parser.add_argument("--expect-51", action="store_true")
    args = parser.parse_args()
    if args.audit.resolve() == args.representatives.resolve():
        parser.error("audit and representatives must have separate paths")
    states, digest = load_states(args.geometry, args.sha256, args.expect_51)
    audit = assess(states, digest, args.grid)
    representative_lines = [json.dumps(row, sort_keys=True, separators=(",", ":"))
                            for row in state_representatives(states, digest)]
    args.audit.write_text(json.dumps(audit, sort_keys=True, indent=2) + "\n")
    args.representatives.write_text("\n".join(representative_lines) + "\n")
    print(json.dumps({"source_state_count": audit["states"],
                      "grid_samples": audit["samples"],
                      "unsampled_state_count": len(audit["unsampled_states"]),
                      "unmatched_points": audit["unmatched_points"],
                      "overlapping_candidate_points": audit["overlapping_candidate_points"],
                      "extra_state_probes": len(representative_lines)}, sort_keys=True))


if __name__ == "__main__":
    main()
