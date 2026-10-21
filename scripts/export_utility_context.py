"""Project only reported EIA county candidates into BCM's read-only reference view.

This is a consumer projection of a verified Electric Utility Catalog release,
not a second importer, identity authority, or serving-utility determination.
"""

from __future__ import annotations

import argparse
import hashlib
import json
import sqlite3
from pathlib import Path


def sha256(path: Path) -> str:
    digest = hashlib.sha256()
    with path.open("rb") as stream:
        for chunk in iter(lambda: stream.read(1 << 20), b""):
            digest.update(chunk)
    return digest.hexdigest()


def export(manifest_path: Path, database_path: Path, output_path: Path) -> str:
    manifest = json.loads(manifest_path.read_bytes())
    if manifest.get("format") != "electric-utility-catalog/v1":
        raise ValueError("unsupported source manifest format")
    if manifest.get("geography") is not None:
        raise ValueError("geographic enrichment requires separately reviewed projection")
    if not isinstance(manifest.get("release_id"), str) or not manifest["release_id"].startswith("eia861-"):
        raise ValueError("invalid EIA release identity")
    expected_sha = manifest.get("artifact", {}).get("sha256")
    if not isinstance(expected_sha, str) or sha256(database_path) != expected_sha:
        raise ValueError("source SQLite digest mismatch")
    year = manifest.get("reporting_year")
    if type(year) is not int or year < 2000 or year > 2100:
        raise ValueError("invalid reporting year")
    source = manifest.get("source", {})
    source_sha = source.get("sha256")
    if not isinstance(source_sha, str) or len(source_sha) != 64:
        raise ValueError("source identity missing")
    uri = "file:" + database_path.resolve().as_posix() + "?mode=ro&immutable=1"
    with sqlite3.connect(uri, uri=True) as conn:
        if conn.execute("PRAGMA integrity_check").fetchone() != ("ok",):
            raise ValueError("catalog SQLite integrity check failed")
        if conn.execute("PRAGMA foreign_key_check").fetchone() is not None:
            raise ValueError("catalog foreign key check failed")
        release = conn.execute("SELECT reporting_year, source_sha256 FROM release").fetchall()
        if release != [(year, source_sha)]:
            raise ValueError("catalog release identity differs from manifest")
        rows = conn.execute("""SELECT c.state, c.county_name, c.county_fips, c.eia_id,
                                   u.reported_name, c.source_sheet, c.source_row
                            FROM service_county c
                            JOIN utility_snapshot u
                              ON u.eia_id=c.eia_id AND u.reporting_year=c.reporting_year
                            WHERE c.reporting_year=?
                            ORDER BY c.state, lower(c.county_name), c.eia_id, c.county_name""", (year,)).fetchall()
        if len(rows) != manifest.get("counts", {}).get("service_county"):
            raise ValueError("county record count differs from manifest")
        activity_rows = conn.execute("""SELECT eia_id, state, activity_code, reported_value
                                        FROM operating_activity
                                        WHERE reporting_year=? AND activity_group='activity'
                                        ORDER BY eia_id, state, activity_code""", (year,)).fetchall()
    activities: dict[tuple[int, str], list[list[str]]] = {}
    for eia_id, state, code, value in activity_rows:
        if value not in ("Y", "N"):
            raise ValueError("unexpected reported activity marker")
        activities.setdefault((eia_id, state), []).append([code, value])
    entries = []
    for state, county, fips, eia_id, name, sheet, row in rows:
        if not all((state, county, name, sheet)) or type(eia_id) is not int or eia_id <= 0:
            raise ValueError("invalid county observation")
        if fips is not None:
            raise ValueError("unreviewed geographic code in EIA-only view")
        entries.append({
            "state": state, "county": county, "eia_id": eia_id,
            "reported_name": name, "source_sheet": sheet, "source_row": row,
            "activities": activities.get((eia_id, state), []),
        })
    payload = {
        "format": "bcm/eia-county-candidates/v1",
        "release_id": manifest["release_id"], "reporting_year": year,
        "source_sha256": source_sha, "sqlite_sha256": expected_sha,
        "source_url": source.get("url"),
        "scope": "EIA-reported county candidates only; no verified serving utility",
        "records": entries,
    }
    encoded = (json.dumps(payload, sort_keys=True, separators=(",", ":"), ensure_ascii=False) + "\n").encode()
    output_path.parent.mkdir(parents=True, exist_ok=True)
    # This projection has its own pinned digest; it is not inserted into or
    # mistaken for the already-published EIA-861 source release manifest.
    output_path.write_bytes(encoded)
    return hashlib.sha256(encoded).hexdigest()


def main() -> None:
    parser = argparse.ArgumentParser()
    parser.add_argument("--manifest", type=Path, required=True)
    parser.add_argument("--database", type=Path, required=True)
    parser.add_argument("--output", type=Path, required=True)
    a = parser.parse_args()
    print(export(a.manifest, a.database, a.output))


if __name__ == "__main__":
    main()
