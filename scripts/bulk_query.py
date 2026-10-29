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
    if "applicability_date" in row:
        value = row["applicability_date"]
        if not isinstance(value, str) or not re.fullmatch(r"\d{4}-\d{2}-\d{2}", value):
            raise ValueError("applicability_date must be YYYY-MM-DD")
        date.fromisoformat(value)
    return row


def classify(status: int, body: bytes, sample_id: str, as_of: str,
             expected_bundle: str = "", expected_source: str = "") -> dict:
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
                any(not isinstance(c, dict) or not isinstance(c.get("family"), str)
                    or not c["family"] or not isinstance(c.get("status"), str) or not c["status"]
                    for c in codes)):
            return dict(result, category="invalid_response")
        verdict = payload.get("verdict")
        if verdict not in {"verified", "conditional", "not_verified"}:
            return dict(result, category="invalid_response")
        evidence = payload.get("evidence", [])
        exact = payload.get("exact_evidence", [])
        if not isinstance(evidence, list) or not isinstance(exact, list):
            return dict(result, category="invalid_response")
        if expected_bundle or expected_source:
            provenance = payload.get("provenance")
            if (not isinstance(provenance, dict) or
                    (expected_bundle and provenance.get("bundle_id") != expected_bundle) or
                    (expected_source and provenance.get("source_commit") != expected_source) or
                    payload["jurisdiction_id"].upper().startswith("DEMO") or
                    (isinstance(payload.get("jurisdiction"), dict) and
                     payload["jurisdiction"].get("country_code") == "DEMO")):
                return dict(result, category="runtime_identity_mismatch")
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
          timeout: float, token: str, limiter: RateLimiter,
          expected_bundle: str = "", expected_source: str = "") -> dict:
    effective_date = row.get("applicability_date", as_of)
    payload = {"project_id": row["id"], "applicability_date": effective_date}
    for key in ("point", "address", "facts", "project_type"):
        if key in row:
            payload[key] = row[key]
    if default_type and not payload.get("project_type"):
        payload["project_type"] = default_type
    data = json.dumps(payload, allow_nan=False, separators=(",", ":")).encode()
    headers = {"Content-Type": "application/json", "X-Request-ID": row["id"]}
    if token:
        headers["Authorization"] = "Bearer " + token
    limiter.acquire()
    started = time.perf_counter()
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
            result = classify(status, body, row["id"], effective_date,
                              expected_bundle, expected_source)
    except (OSError, ValueError, URLError):
        result = {"id": row["id"], "category": "transport_error"}
    result["applicability_date"] = effective_date
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
    if not re.fullmatch(r"\d{4}-\d{2}-\d{2}", args.date):
        raise ValueError("date must be YYYY-MM-DD")
    date.fromisoformat(args.date)
    if not (0 <= args.shard_index < args.shard_count <= 1024):
        raise ValueError("shard-index must be 0..shard-count-1; shard-count 1..1024")
    if not (1 <= args.concurrency <= 64 and 1 <= args.qps <= 100
            and 1 <= args.max_requests <= 5_000_000 and 1 <= args.timeout <= 120):
        raise ValueError("concurrency 1..64, qps 1..100, max-requests 1..5000000, timeout 1..120")
    abort_after = getattr(args, "abort_after_errors", 0)
    if not isinstance(abort_after, int) or isinstance(abort_after, bool) or not 0 <= abort_after <= 100000:
        raise ValueError("abort-after-errors must be 0..100000")
    paths = [Path(args.input).resolve(), Path(args.results).resolve(), Path(args.summary).resolve()]
    if len(set(paths)) != 3:
        raise ValueError("input, results and summary paths must differ")
    expected_bundle = getattr(args, "expected_bundle", "")
    expected_source = getattr(args, "expected_source", "")
    if bool(expected_bundle) != bool(expected_source):
        raise ValueError("live runs must pin both --expected-bundle and --expected-source")
    if expected_bundle and (
            not re.fullmatch(r"[a-zA-Z0-9:._-]{8,200}", expected_bundle)
            or not re.fullmatch(r"[a-f0-9]{40}", expected_source)
            or "demo" in expected_bundle.lower()):
        raise ValueError("invalid or demo runtime identity pin")
    token = os.environ.get(args.token_env, "") if args.token_env else ""
    if args.token_env and not token:
        raise ValueError("token environment variable is empty")
    counts, status_counts, families = Counter(), Counter(), Counter()
    latencies: list[float] = []
    digest = sha256()
    limiter = RateLimiter(args.qps)
    seen: set[str] = set()
    started = time.perf_counter()
    skipped = 0
    errors_seen = 0
    circuit_breaker_tripped = False
    input_validation_complete = True
    fatal_categories = {"runtime_identity_mismatch", "invalid_response",
                        "unsupported_verification", "service_error",
                        "http_error", "transport_error"}
    results_path = paths[1]
    with tempfile.NamedTemporaryFile("w", encoding="utf-8", dir=results_path.parent,
                                     prefix=results_path.name + ".", delete=False) as output:
        tmp = Path(output.name)
        try:
            with open(paths[0], "rb") as source, ThreadPoolExecutor(max_workers=args.concurrency) as pool:
                pending = set()
                sample_count = 0

                def collect(done):
                    nonlocal errors_seen
                    for future in done:
                        outcome = future.result()
                        output.write(json.dumps(outcome, sort_keys=True) + "\n")
                        counts[outcome["category"]] += 1
                        if outcome["category"] in fatal_categories:
                            errors_seen += 1
                        status_counts[str(outcome.get("http_status", 0))] += 1
                        families.update(outcome.get("code_families", []))
                        latencies.append(outcome["latency_ms"])

                for line in source:
                    digest.update(line)
                    if not line.strip():
                        continue
                    try:
                        row = validate_sample(json.loads(line))
                    except (ValueError, TypeError, KeyError) as error:
                        raise ValueError("invalid sample at input line " + str(sample_count + 1)) from error
                    if row["id"] in seen:
                        raise ValueError("duplicate sample id: " + row["id"])
                    seen.add(row["id"])
                    if int.from_bytes(sha256(row["id"].encode()).digest()[:8], "big") % args.shard_count != args.shard_index:
                        skipped += 1
                        continue
                    if sample_count >= args.max_requests:
                        raise ValueError("workload shard exceeds --max-requests")
                    pending.add(pool.submit(query, row, args.endpoint, args.date,
                                            args.project_type, args.timeout, token, limiter,
                                            expected_bundle, expected_source))
                    sample_count += 1
                    # With a breaker armed, never queue more requests than
                    # the available concurrent slots. Already-in-flight calls
                    # may complete, but no more work is submitted once tripped.
                    limit = args.concurrency if abort_after else args.concurrency * 2
                    if len(pending) >= limit:
                        done, pending = wait(pending, return_when=FIRST_COMPLETED)
                        collect(done)
                        if abort_after and errors_seen >= abort_after:
                            circuit_breaker_tripped = True
                            input_validation_complete = False
                            # Hash the complete immutable byte input even though
                            # unexamined rows must not be counted as validated.
                            digest.update(source.read())
                            break
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
        "samples": sample_count, "skipped_samples": skipped,
        "shard_index": args.shard_index, "shard_count": args.shard_count,
        "categories": dict(sorted(counts.items())),
        **({"circuit_breaker_tripped": circuit_breaker_tripped,
            "input_validation_complete": input_validation_complete,
            "fatal_error_count": errors_seen,
            "abort_after_errors": abort_after} if abort_after else {}),
        **({"expected_runtime_identity": {"bundle_id": expected_bundle,
                                          "source_commit": expected_source}}
           if expected_bundle else {}),
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
    p.add_argument("--max-requests", type=int, default=5000,
                   help="maximum executed samples per shard")
    p.add_argument("--shard-index", type=int, default=0)
    p.add_argument("--shard-count", type=int, default=1)
    p.add_argument("--timeout", type=float, default=8)
    p.add_argument("--expected-bundle", default="", help="independently pinned non-demo provider bundle")
    p.add_argument("--expected-source", default="", help="independently pinned 40-character provider source commit")
    p.add_argument("--abort-after-errors", type=int, default=0,
                   help="stop scheduling new calls after N fatal results (0 disables breaker)")
    p.add_argument("--fail-on-errors", action="store_true",
                   help="exit nonzero for transport, server, invalid response or identity drift")
    p.add_argument("--allow-remote", action="store_true")
    return p


if __name__ == "__main__":
    try:
        args = parser().parse_args()
        report = run(args)
        if report.get("circuit_breaker_tripped"):
            raise ValueError("circuit breaker tripped; inspect sanitized partial results")
        if args.fail_on_errors:
            defects = {"runtime_identity_mismatch", "invalid_response",
                       "unsupported_verification", "service_error",
                       "http_error", "transport_error"}
            if any(report["categories"].get(category, 0) for category in defects):
                raise ValueError("run has non-accepted outcomes; inspect sanitized summary")
    except (OSError, ValueError) as error:
        raise SystemExit("bulk run failed: " + str(error)) from error
