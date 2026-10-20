#!/usr/bin/env python3
"""Acquire State of Colorado address points into BCM's reviewed CSV contract.

Downloads source data only to a caller-owned local path; never checks data into Git.
Release admission still requires boundary, code and source-health verification.
"""
from __future__ import annotations

import argparse
import csv
import hashlib
import json
import os
from pathlib import Path
import tempfile
from urllib.parse import urlencode
from urllib.request import urlopen

SOURCE = (
    "https://gis.colorado.gov/public/rest/services/"
    "Address_and_Parcel/Colorado_Public_Addresses/FeatureServer/0/query"
)
FIELDS = (
    "OBJECTID,SAUID,AddrNum,NumSuf,St_PreMod,PreDir,PreType,St_PreSep,"
    "StreetName,PostType,PostDir,St_PosMod,PlaceName,Zipcode,"
    "Latitude,Longitude,AddrFull"
)
COLUMNS = (
    "source_record_id", "address_number", "street", "city", "state",
    "postal_code", "longitude", "latitude", "matched_address"
)
PAGE_SIZE = 2000
ATTRIBUTION = (
    "State of Colorado Public addresses were accessed on {date} from "
    "https://gis.colorado.gov/public/rest/services/Address_and_Parcel"
)


def fetch(params: dict[str, str | int]) -> dict:
    url = SOURCE + "?" + urlencode({**params, "f": "json"})
    with urlopen(url, timeout=45) as response:
        data = response.read(24 * 1024 * 1024)
    value = json.loads(data)
    if not isinstance(value, dict) or "error" in value:
        raise ValueError("Colorado GIS returned an error or invalid JSON payload")
    return value


def parse_row(feature: dict) -> dict[str, str]:
    attrs = feature.get("attributes")
    if not isinstance(attrs, dict):
        raise ValueError("missing attributes")
    object_id = attrs.get("OBJECTID")
    if not isinstance(object_id, int) or object_id <= 0:
        raise ValueError("missing OBJECTID")
    address_number = str(attrs.get("AddrNum") or "").strip()
    street = " ".join(
        str(attrs.get(field) or "").strip() for field in
        ("St_PreMod", "PreDir", "PreType", "St_PreSep", "StreetName",
         "PostType", "PostDir", "St_PosMod")
        if str(attrs.get(field) or "").strip()
    )
    suffix = str(attrs.get("NumSuf") or "").strip()
    if suffix:
        address_number += suffix
    city = str(attrs.get("PlaceName") or "").strip()
    zip_code = str(attrs.get("Zipcode") or "").strip()
    latitude, longitude = attrs.get("Latitude"), attrs.get("Longitude")
    if not address_number or not street or not city:
        raise ValueError("missing street number, street name or locality")
    if isinstance(latitude, bool) or isinstance(longitude, bool):
        raise ValueError("invalid coordinates")
    try:
        lat, lon = float(latitude), float(longitude)
    except (TypeError, ValueError) as error:
        raise ValueError("invalid coordinates") from error
    # Out-of-state location errors must not silently become valid BCM points.
    if not (-109.2 <= lon <= -102.0 and 36.9 <= lat <= 41.1):
        raise ValueError("outside Colorado bounding envelope")
    return {
        "source_record_id": f"co-oit-{object_id}",
        "address_number": address_number,
        "street": street,
        "city": city,
        "state": "CO",
        "postal_code": zip_code,
        "longitude": format(lon, ".10f"),
        "latitude": format(lat, ".10f"),
        "matched_address": f"{address_number} {street}, {city}, CO {zip_code}".strip(),
    }


def export(fetch_page, output: Path, date: str, max_invalid: int = 0,
           max_records: int = 5_000_000) -> dict:
    count_data = fetch_page({"where": "1=1", "returnCountOnly": "true"})
    total = count_data.get("count")
    if not isinstance(total, int) or total < 1 or total > max_records:
        raise ValueError("source count missing or outside configured limit")
    if max_invalid < 0:
        raise ValueError("max_invalid must be nonnegative")
    output.parent.mkdir(parents=True, exist_ok=True)
    with tempfile.NamedTemporaryFile(
        mode="w", encoding="utf-8", newline="", dir=output.parent,
        prefix=output.name + ".", suffix=".partial", delete=False
    ) as handle:
        tmp = Path(handle.name)
        writer = csv.DictWriter(handle, fieldnames=COLUMNS, lineterminator="\n")
        writer.writeheader()
        accepted, invalid, seen = 0, 0, set()
        last_object_id = 0
        try:
            for offset in range(0, total, PAGE_SIZE):
                page = fetch_page({
                    "where": "1=1", "outFields": FIELDS,
                    "orderByFields": "OBJECTID ASC",
                    "resultOffset": offset,
                    "resultRecordCount": min(PAGE_SIZE, total - offset),
                    "returnGeometry": "false", "outSR": 4326,
                })
                rows = page.get("features")
                if not isinstance(rows, list) or len(rows) != min(PAGE_SIZE, total - offset):
                    raise ValueError("source pagination was incomplete")
                for feature in rows:
                    record_id = feature.get("attributes", {}).get("OBJECTID")
                    if not isinstance(record_id, int) or record_id in seen or record_id <= last_object_id:
                        raise ValueError("duplicate or nonmonotonic source OBJECTID")
                    seen.add(record_id)
                    last_object_id = record_id
                    try:
                        parsed = parse_row(feature)
                    except ValueError:
                        invalid += 1
                        if invalid > max_invalid:
                            raise ValueError("source has invalid records; no snapshot published") from None
                    else:
                        writer.writerow(parsed)
                        accepted += 1
            if len(seen) != total or accepted == 0:
                raise ValueError("record count mismatch or empty acceptable dataset")
        except BaseException:
            handle.close()
            tmp.unlink(missing_ok=True)
            raise
    digest = hashlib.sha256(tmp.read_bytes()).hexdigest()
    os.replace(tmp, output)
    receipt = {
        "schema_version": "1",
        "source_url": SOURCE,
        "source_count": total,
        "accepted": accepted,
        "invalid": invalid,
        "csv_sha256": digest,
        "retrieved_at": date,
        "attribution": ATTRIBUTION.format(date=date),
        "output_path": str(output),
        "production_approved": False,
    }
    output.with_suffix(output.suffix + ".receipt.json").write_text(
        json.dumps(receipt, sort_keys=True, indent=2) + "\n", encoding="utf-8"
    )
    return receipt


def main() -> None:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--output", type=Path, required=True)
    parser.add_argument("--retrieved-at", required=True, help="Actual UTC acquisition date YYYY-MM-DD")
    parser.add_argument("--max-invalid", type=int, default=0)
    parser.add_argument("--max-records", type=int, default=5_000_000)
    args = parser.parse_args()
    from datetime import date
    if date.fromisoformat(args.retrieved_at).isoformat() != args.retrieved_at:
        parser.error("--retrieved-at must be an ISO date")
    receipt = export(fetch, args.output, args.retrieved_at,
                     args.max_invalid, args.max_records)
    print(json.dumps(receipt, sort_keys=True))


if __name__ == "__main__":
    main()
