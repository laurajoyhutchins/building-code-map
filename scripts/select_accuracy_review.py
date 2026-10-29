#!/usr/bin/env python3
"""Create a blinded, *pending* legal-accuracy review queue from canonical samples.

No provider responses, observed verdicts, addresses or coordinate data in output.
This program cannot produce an accepted legal truth or bcm-accuracy case.
"""
from __future__ import annotations

import argparse
from collections import Counter, defaultdict
from hashlib import sha256
import json
from pathlib import Path

from scripts.bulk_query import atomic_json, validate_sample

COHORTS = ("grid", "boundary", "source_gap", "transition_date", "other")


def tags(row):
    result = []
    if "cell" in row:
        result.append("grid")
    if "boundary" in row:
        result.append("boundary")
    if "gap" in row:
        result.append("source_gap")
    if "temporal" in row:
        result.append("transition_date")
    return result or ["other"]


def select(input_path: Path, max_cases: int):
    if not 1 <= max_cases <= 1000:
        raise ValueError("max cases must be 1..1000")
    raw = input_path.read_bytes()
    if len(raw) > 32 * 1024 * 1024:
        raise ValueError("source workload exceeds 32 MiB")
    seen = set()
    groups = defaultdict(list)
    total = 0
    for line_number, line in enumerate(raw.splitlines(), 1):
        if not line.strip():
            continue
        try:
            sample = validate_sample(json.loads(line))
        except (ValueError, TypeError, KeyError) as error:
            raise ValueError(f"invalid workload sample at line {line_number}") from error
        if sample["id"] in seen:
            raise ValueError("duplicate workload sample ID")
        seen.add(sample["id"])
        total += 1
        rowtags = tags(sample)
        # For mixed cases prefer the most diagnostic explicit metadata.
        priority = ("transition_date", "source_gap", "boundary", "grid", "other")
        cohort = next(k for k in priority if k in rowtags)
        rank = sha256(sample["id"].encode()).hexdigest()
        groups[cohort].append((rank, sample["id"], {
            "case_id": "review:" + rank[:24],
            "workload_id": sample["id"],
            "candidate_tags": sorted(rowtags),
            "status": "pending_independent_review",
        }))
    if not total:
        raise ValueError("empty workload")
    for cohort in groups:
        groups[cohort].sort(key=lambda entry: (entry[0], entry[1]))
    selected = []
    counts = Counter()
    # Balanced rotating strata: one case per observed source category before
    # taking another, with no randomness and no invented expected verdict.
    while len(selected) < max_cases:
        progressed = False
        for cohort in COHORTS:
            index = counts[cohort]
            if cohort in groups and index < len(groups[cohort]) and len(selected) < max_cases:
                selected.append(groups[cohort][index][2])
                counts[cohort] += 1
                progressed = True
        if not progressed:
            break
    selected.sort(key=lambda row: row["case_id"])
    result = {
        "schema_version": "1", "status": "pending_independent_review",
        "input_sha256": sha256(raw).hexdigest(),
        "source_samples": total, "max_cases": max_cases,
        "case_count": len(selected),
        "selected_cohorts": dict(sorted((k, v) for k, v in counts.items() if v)),
        "cases": selected,
        "legal_accuracy_reviewed": False,
        "expected_results_present": False,
    }
    return result


def main(argv=None):
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--input", type=Path, required=True, help="canonical JSONL workload")
    parser.add_argument("--output", type=Path, required=True, help="private review queue JSON")
    parser.add_argument("--max-cases", type=int, default=100)
    args = parser.parse_args(argv)
    if args.input.resolve() == args.output.resolve():
        parser.error("input and output must differ")
    result = select(args.input, args.max_cases)
    atomic_json(args.output, result)
    print(json.dumps({"queue_sha256": sha256(args.output.read_bytes()).hexdigest(),
                      "input_sha256": result["input_sha256"],
                      "case_count": result["case_count"],
                      "status": result["status"]}, sort_keys=True))


if __name__ == "__main__":
    main()
