#!/usr/bin/env python3
"""Normalize an exact Census 2024 state-boundary shapefile ZIP to 50 states + DC GeoJSON.

Uses the existing pyshp library for the shapefile format; does not implement a
custom geometry parser. Source archive and derived geometry SHA-256 are reported.
"""
from __future__ import annotations

import argparse
from hashlib import sha256
from io import BytesIO
import json
from pathlib import Path
from zipfile import ZipFile

STATE_FIPS = frozenset((
    "01", "02", "04", "05", "06", "08", "09", "10", "11", "12", "13",
    "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25",
    "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36",
    "37", "38", "39", "40", "41", "42", "44", "45", "46", "47", "48",
    "49", "50", "51", "53", "54", "55", "56",
))
CENSUS_SOURCE_URL = "https://www2.census.gov/geo/tiger/GENZ2024/shp/cb_2024_us_state_500k.zip"


def prepare(archive: Path, expected_sha256: str | None = None):
    import shapefile
    from shapely.geometry import mapping, shape

    original = archive.read_bytes()
    original_sha = sha256(original).hexdigest()
    if expected_sha256 is not None and original_sha != expected_sha256.lower():
        raise ValueError("Census source archive SHA-256 mismatch")
    with ZipFile(BytesIO(original)) as source:
        paths = source.namelist()
        entries = {}
        for extension in ("shp", "shx", "dbf"):
            selected = [path for path in paths if path.endswith("." + extension) and
                        Path(path).name == "cb_2024_us_state_500k." + extension]
            if len(selected) != 1:
                raise ValueError("expected one Census state ." + extension + " entry")
            entries[extension] = BytesIO(source.read(selected[0]))
    reader = shapefile.Reader(shp=entries["shp"], shx=entries["shx"], dbf=entries["dbf"])
    features = []
    seen = set()
    for record in reader.iterShapeRecords():
        attributes = record.record.as_dict()
        fips = str(attributes.get("STATEFP", ""))
        if fips not in STATE_FIPS:
            continue
        if fips in seen:
            raise ValueError("duplicate state FIPS in source: " + fips)
        geometry = shape(record.shape.__geo_interface__)
        if geometry.geom_type not in ("Polygon", "MultiPolygon") or geometry.is_empty or not geometry.is_valid:
            raise ValueError("invalid Census state geometry: " + fips)
        seen.add(fips)
        features.append({"type": "Feature", "properties": {"STATEFP": fips},
                         "geometry": mapping(geometry)})
    if seen != STATE_FIPS:
        raise ValueError("Census source missing 50-state-and-DC coverage: " +
                         ", ".join(sorted(STATE_FIPS - seen)))
    features.sort(key=lambda value: value["properties"]["STATEFP"])
    return ({"type": "FeatureCollection", "features": features}, original_sha)


def main():
    p = argparse.ArgumentParser(description=__doc__)
    p.add_argument("--archive", type=Path, required=True, help="downloaded Census 2024 500k state ZIP")
    p.add_argument("--archive-sha256", help="optional already-pinned ZIP SHA-256; exact match required")
    p.add_argument("--output", type=Path, required=True, help="derived WGS84 GeoJSON path")
    args = p.parse_args()
    geometry, archive_sha = prepare(args.archive, args.archive_sha256)
    content = (json.dumps(geometry, sort_keys=True, separators=(",", ":")) + "\n").encode()
    target = args.output.with_name(args.output.name + ".tmp")
    try:
        target.write_bytes(content)
        target.replace(args.output)
    finally:
        target.unlink(missing_ok=True)
    print(json.dumps({"census_source_url": CENSUS_SOURCE_URL,
                      "archive_sha256": archive_sha, "geometry_sha256": sha256(content).hexdigest(),
                      "state_equivalents": len(geometry["features"]),
                      "output": str(args.output)}, sort_keys=True))


if __name__ == "__main__":
    main()
