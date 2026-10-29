#!/usr/bin/env python3
"""Expand candidate transition dates to -1/0/+1 day canonical BCM test queries.

These are source-date candidates, not independently reviewed operative law.
"""
from __future__ import annotations

import argparse
from datetime import date, timedelta
from hashlib import sha256
import json
from pathlib import Path
import re

from scripts.bulk_query import atomic_json, validate_sample

SAFE_ID = re.compile(r"^[A-Za-z0-9][A-Za-z0-9._:-]{0,70}$")
HEX = re.compile(r"^[0-9a-f]{64}$")


def parse_events(path: Path):
    raw = path.read_bytes()
    if len(raw) > 256 * 1024:
        raise ValueError("transition manifest exceeds 256 KiB")
    doc = json.loads(raw)
    if (not isinstance(doc, dict) or
            set(doc) != {"schema_version", "classification", "events"} or
            doc["schema_version"] != "1" or
            doc["classification"] != "source_date_candidate_only"):
        raise ValueError("transition manifest must be candidate-only version 1")
    events = doc["events"]
    if not isinstance(events, list) or not 1 <= len(events) <= 100:
        raise ValueError("transition manifest must contain 1..100 events")
    seen = set()
    for event in events:
        if not isinstance(event, dict) or set(event) != {"id", "transition_date", "source_sha256"}:
            raise ValueError("candidate transition event has invalid fields")
        ident, day, digest = event["id"], event["transition_date"], event["source_sha256"]
        if not isinstance(ident, str) or not SAFE_ID.fullmatch(ident) or ident in seen:
            raise ValueError("transition event ID invalid or duplicated")
        if not isinstance(digest, str) or not HEX.fullmatch(digest):
            raise ValueError("transition event needs source observation SHA-256")
        if not isinstance(day, str) or not re.fullmatch(r"\d{4}-\d{2}-\d{2}", day):
            raise ValueError("candidate transition date must be YYYY-MM-DD")
        try:
            when = date.fromisoformat(day)
            _ = when - timedelta(days=1)
            _ = when + timedelta(days=1)
        except (ValueError, OverflowError):
            raise ValueError("transition date outside -1/+1 calendar-day window") from None
        seen.add(ident)
    return sorted(events, key=lambda event: event["id"]), sha256(raw).hexdigest()


def expand(input_path: Path, events: list, max_requests: int):
    if not 1 <= max_requests <= 200000:
        raise ValueError("max requests must be 1..200000")
    source_bytes = input_path.read_bytes()
    if len(source_bytes) > 32 * 1024 * 1024:
        raise ValueError("source workload exceeds 32 MiB")
    points, known = [], set()
    for line_number, line in enumerate(source_bytes.splitlines(), 1):
        if not line.strip():
            continue
        try:
            sample = validate_sample(json.loads(line))
        except (TypeError, ValueError, KeyError) as exc:
            raise ValueError(f"invalid base sample at line {line_number}") from exc
        if "temporal" in sample or sample["id"] in known:
            raise ValueError("recursive temporal or duplicate base sample")
        known.add(sample["id"])
        points.append(sample)
        if len(points) * len(events) * 3 > max_requests:
            raise ValueError("expanded temporal requests exceed max requests")
    if not points:
        raise ValueError("no base samples to expand")
    result = []
    for sample in points:
        for event in events:
            when = date.fromisoformat(event["transition_date"])
            for offset in (-1, 0, 1):
                actual_date = (when + timedelta(days=offset)).isoformat()
                identity = sha256((sample["id"] + "\0" + event["id"] + "\0" + str(offset)).encode()).hexdigest()[:24]
                row = {k: v for k, v in sample.items() if k not in {"id", "applicability_date"}}
                row.update({"id": "temporal:" + identity,
                            "applicability_date": actual_date,
                            "temporal": {"event_id": event["id"],
                                         "transition_date": event["transition_date"],
                                         "offset_days": offset,
                                         "source_sha256": event["source_sha256"],
                                         "classification": "source_date_candidate_only"}})
                result.append(row)
    result.sort(key=lambda row: row["id"])
    return result, sha256(source_bytes).hexdigest(), len(points)


def main(argv=None):
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--input", type=Path, required=True)
    parser.add_argument("--events", type=Path, required=True)
    parser.add_argument("--output", type=Path, required=True)
    parser.add_argument("--summary", type=Path, required=True)
    parser.add_argument("--max-requests", type=int, default=10000)
    args = parser.parse_args(argv)
    if len({x.resolve() for x in (args.input, args.events, args.output, args.summary)}) != 4:
        parser.error("input, events, output, summary must be separate paths")
    events, manifest_digest = parse_events(args.events)
    rows, input_digest, points = expand(args.input, events, args.max_requests)
    target = args.output.with_name(args.output.name + ".tmp")
    try:
        with target.open("w", encoding="utf-8") as handle:
            for row in rows:
                handle.write(json.dumps(row, sort_keys=True, separators=(",", ":")) + "\n")
            handle.flush()
            import os
            os.fsync(handle.fileno())
        report = {"schema_version": "1", "classification": "source_date_candidate_only",
                  "input_sha256": input_digest, "transition_manifest_sha256": manifest_digest,
                  "workload_sha256": sha256(target.read_bytes()).hexdigest(),
                  "source_samples": points, "candidate_events": len(events), "samples": len(rows)}
        target.replace(args.output)
        atomic_json(args.summary, report)
    finally:
        target.unlink(missing_ok=True)
    print(json.dumps(report, sort_keys=True))


if __name__ == "__main__":
    main()
