#!/usr/bin/env python3
"""Bounded workload runner for BCM's canonical project verification HTTP API.

This is a transport/load probe, not a second legal verdict implementation.
"""
from __future__ import annotations

import argparse
from collections import Counter
from concurrent.futures import FIRST_COMPLETED, ThreadPoolExecutor, wait
from datetime import date
from hashlib import sha256
import ipaddress
import json
import math
import os
from pathlib import Path
import re
import tempfile
import threading
import time
from urllib.error import HTTPError, URLError
from urllib.parse import urlsplit
from urllib.request import HTTPRedirectHandler, Request, build_opener

PATH = "/v1/project-code-basis"
MAX_RESPONSE = 1024 * 1024
EXPECTED_UNRESOLVED = {
    "address_not_found", "jurisdiction_profile_missing",
    "address_ambiguous", "boundary_ambiguous", "outside_supported_coverage",
}


class NoRedirect(HTTPRedirectHandler):
    def redirect_request(self, request, fp, code, msg, headers, newurl):
        raise ValueError("redirects are forbidden")


def validate_target(endpoint: str, allow_remote: bool) -> None:
    target = urlsplit(endpoint)
    if (target.scheme not in ("http", "https") or not target.hostname
            or target.path != PATH or target.query or target.fragment
            or target.username is not None or target.password is not None):
        raise ValueError("endpoint must be an HTTP(S) origin plus " + PATH)
    host = target.hostname
    try:
        local = ipaddress.ip_address(host).is_loopback
    except ValueError:
        local = host.lower() == "localhost"
    if not local and (not allow_remote or target.scheme != "https"):
        raise ValueError("remote targets require HTTPS and --allow-remote")


def validate_sample(row: dict) -> dict:
    if not isinstance(row, dict) or not re.fullmatch(r"[a-zA-Z0-9][a-zA-Z0-9._:-]{0,127}", row.get("id", "")):
        raise ValueError("sample requires a safe unique id (1..128 characters)")
    point, address = row.get("point"), row.get("address")
    if (point is None) == (not bool(address)):
        raise ValueError("exactly one of point or address is required")
    if point is not None:
        if not isinstance(point, dict) or set(point) != {"latitude", "longitude"}:
            raise ValueError("point must contain latitude and longitude")
        for name, low, high in (("latitude", -90, 90), ("longitude", -180, 180)):
            value = point[name]
            if isinstance(value, bool) or not isinstance(value, (float, int)) or not math.isfinite(value) or not low <= value <= high:
                raise ValueError("invalid " + name)
    elif not isinstance(address, str) or not address.strip():
        raise ValueError("address must be nonempty")
    if "facts" in row and (not isinstance(row["facts"], dict) or
       any(not isinstance(k, str) or not k.strip() or not isinstance(v, str) or not v.strip() for k, v in row["facts"].items())):
        raise ValueError("facts must be a string-to-string map")
    if "project_type" in row and not isinstance(row["project_type"], str):
        raise ValueError("project_type must be a string")
    return row


def classify(status: int, body: bytes, sample_id: str, as_of: str) -> dict:
    result = {"id": sample_id, "http_status": status}
    try:
        payload = json.loads(body)
    except (ValueError, UnicodeError):
        return dict(result, category="invalid_response")
    if not isinstance(payload, dict):
        return dict(result, category="invalid_response")
    if status == 200:
        codes = payload.get("code_set")
        if (payload.get("schema_version") != "0.3" or
                payload.get("project_id") != sample_id or payload.get("as_of") != as_of or
                not isinstance(payload.get("jurisdiction_id"), str) or
                not payload["jurisdiction_id"] or
                not isinstance(codes, list) or not codes or
                any(not isinstance(c, dict) or not c.get("family") or not c.get("status") for c in codes)):
            return dict(result, category="invalid_response")
        verdict = payload.get("verdict")
        if verdict not in {"verified", "conditional", "not_verified"}:
            return dict(result, category="invalid_response")
        evidence = payload.get("evidence", [])
        exact = payload.get("exact_evidence", [])
        if not isinstance(evidence, list) or not isinstance(exact, list):
            return dict(result, category="invalid_response")
        if verdict == "verified" and not evidence and not exact:
            return dict(result, category="unsupported_verification")
        return dict(result, category=verdict, verdict=verdict,
                    jurisdiction_id=payload["jurisdiction_id"],
                    code_families=sorted({c["family"] for c in codes}),
                    evidence_count=len(evidence) + len(exact))
    code = payload.get("error", {}).get("code") if isinstance(payload.get("error"), dict) else None
    if isinstance(code, str):
        result["error_code"] = code
    if status in (404, 422) and code in EXPECTED_UNRESOLVED:
        return dict(result, category="unresolved")
    if status == 429 or status >= 500:
        return dict(result, category="service_error")
    return dict(result, category="http_error")


class RateLimiter:
    def __init__(self, rate: int):
        self.period = 1.0 / rate
        self.next_slot = time.monotonic()
        self.lock = threading.Lock()

    def acquire(self):
        with self.lock:
            now = time.monotonic()
            slot = max(now, self.next_slot)
            self.next_slot = slot + self.period
        time.sleep(max(0, slot - now))


def query(row: dict, endpoint: str, as_of: str, default_type: str,
          timeout: float, token: str, limiter: RateLimiter) -> dict:
    payload = {"project_id": row["id"], "applicability_date": as_of}
    for key in ("point", "address", "facts", "project_type"):
        if key in row:
            payload[key] = row[key]
    if default_type and not payload.get("project_type"):
        payload["project_type"] = default_type
    data = json.dumps(payload, allow_nan=False, separators=(",", ":")).encode()
    headers = {"Content-Type": "application/json", "X-Request-ID": row["id"]}
    if token:
        headers["Authorization"] = "Bearer " + token
    started = time.perf_counter()
    limiter.acquire()
    try:
        request = Request(endpoint, data=data, headers=headers, method="POST")
        opener = build_opener(NoRedirect())
        try:
            response = opener.open(request, timeout=timeout)
        except HTTPError as error:
            response = error
        with response:
            status = response.status
            body = response.read(MAX_RESPONSE + 1)
        if len(body) > MAX_RESPONSE:
            result = {"id": row["id"], "http_status": status, "category": "invalid_response"}
        else:
            result = classify(status, body, row["id"], as_of)
    except (OSError, ValueError, URLError):
        result = {"id": row["id"], "category": "transport_error"}
    result["latency_ms"] = round((time.perf_counter() - started) * 1000, 3)
    return result


def percentile(sorted_values: list[float], fraction: float) -> float:
    return sorted_values[max(0, math.ceil(len(sorted_values) * fraction) - 1)] if sorted_values else 0


def atomic_json(path: Path, value: dict):
    with tempfile.NamedTemporaryFile("w", encoding="utf-8", dir=path.parent,
                                     prefix=path.name + ".", delete=False) as handle:
        temporary = Path(handle.name)
        json.dump(value, handle, sort_keys=True, indent=2)
        handle.write("\n")
    os.replace(temporary, path)


def run(args: argparse.Namespace) -> dict:
    validate_target(args.endpoint, args.allow_remote)
    date.fromisoformat(args.date)
    if not (1 <= args.concurrency <= 64 and 1 <= args.qps <= 100
            and 1 <= args.max_requests <= 5_000_000 and 1 <= args.timeout <= 120):
        raise ValueError("concurrency 1..64, qps 1..100, max-requests 1..5000000, timeout 1..120")
    paths = [Path(args.input).resolve(), Path(args.results).resolve(), Path(args.summary).resolve()]
    if len(set(paths)) != 3:
        raise ValueError("input, results and summary paths must differ")
    token = os.environ.get(args.token_env, "") if args.token_env else ""
    if args.token_env and not token:
        raise ValueError("token environment variable is empty")
    counts, status_counts, families = Counter(), Counter(), Counter()
    latencies: list[float] = []
    digest = sha256()
    limiter = RateLimiter(args.qps)
    seen: set[str] = set()
    started = time.perf_counter()
    results_path = paths[1]
    with tempfile.NamedTemporaryFile("w", encoding="utf-8", dir=results_path.parent,
                                     prefix=results_path.name + ".", delete=False) as output:
        tmp = Path(output.name)
        try:
            with open(paths[0], "rb") as source, ThreadPoolExecutor(max_workers=args.concurrency) as pool:
                pending = set()
                sample_count = 0

                def collect(done):
                    for future in done:
                        outcome = future.result()
                        output.write(json.dumps(outcome, sort_keys=True) + "\n")
                        counts[outcome["category"]] += 1
                        status_counts[str(outcome.get("http_status", 0))] += 1
                        families.update(outcome.get("code_families", []))
                        latencies.append(outcome["latency_ms"])

                for line in source:
                    digest.update(line)
                    if not line.strip():
                        continue
                    if sample_count >= args.max_requests:
                        raise ValueError("workload exceeds --max-requests")
                    try:
                        row = validate_sample(json.loads(line))
                    except (ValueError, TypeError, KeyError) as error:
                        raise ValueError("invalid sample at input line " + str(sample_count + 1)) from error
                    if row["id"] in seen:
                        raise ValueError("duplicate sample id: " + row["id"])
                    seen.add(row["id"])
                    pending.add(pool.submit(query, row, args.endpoint, args.date,
                                            args.project_type, args.timeout, token, limiter))
                    sample_count += 1
                    if len(pending) >= args.concurrency * 2:
                        done, pending = wait(pending, return_when=FIRST_COMPLETED)
                        collect(done)
                while pending:
                    done, pending = wait(pending, return_when=FIRST_COMPLETED)
                    collect(done)
            output.flush()
            os.fsync(output.fileno())
        except BaseException:
            tmp.unlink(missing_ok=True)
            raise
    os.replace(tmp, results_path)
    latencies.sort()
    duration = max(time.perf_counter() - started, 0.000001)
    summary = {
        "schema_version": "1", "input_sha256": digest.hexdigest(),
        "endpoint_path": PATH, "applicability_date": args.date,
        "samples": sample_count, "categories": dict(sorted(counts.items())),
        "http_statuses": dict(sorted(status_counts.items())),
        "code_families": dict(sorted(families.items())),
        "duration_ms": round(duration * 1000),
        "requests_per_second": round(sample_count / duration, 3),
        "latency_p50_ms": percentile(latencies, .5),
        "latency_p95_ms": percentile(latencies, .95),
        "latency_p99_ms": percentile(latencies, .99),
    }
    atomic_json(paths[2], summary)
    return summary


def parser() -> argparse.ArgumentParser:
    p = argparse.ArgumentParser(description=__doc__)
    p.add_argument("--input", required=True, help="JSONL sample workload")
    p.add_argument("--results", required=True, help="per-sample JSONL outcomes")
    p.add_argument("--summary", required=True, help="aggregate JSON")
    p.add_argument("--date", required=True, help="YYYY-MM-DD applicability date")
    p.add_argument("--endpoint", default="http://127.0.0.1:8787" + PATH)
    p.add_argument("--project-type", default="")
    p.add_argument("--token-env", default="", help="name of bearer-token environment variable")
    p.add_argument("--concurrency", type=int, default=4)
    p.add_argument("--qps", type=int, default=2)
    p.add_argument("--max-requests", type=int, default=5000)
    p.add_argument("--timeout", type=float, default=8)
    p.add_argument("--allow-remote", action="store_true")
    return p


if __name__ == "__main__":
    try:
        run(parser().parse_args())
    except (OSError, ValueError) as error:
        raise SystemExit("bulk run failed: " + str(error)) from error
