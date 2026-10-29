#!/usr/bin/env python3
"""Fail-closed one-project canary before any live BCM bulk request workload.

This checks runtime identity and a concrete coordinate result, not legal truth.
"""
from __future__ import annotations

import argparse
from hashlib import sha256
import json
import os
from pathlib import Path
import re
from urllib.error import HTTPError, URLError
from urllib.parse import urlsplit
from urllib.request import Request, build_opener

from scripts.bulk_query import (MAX_RESPONSE, NoRedirect, PATH, RateLimiter,
                                classify, query, validate_sample, validate_target)


def json_get(url: str, token: str, timeout: float) -> dict:
    headers = {"Accept": "application/json"}
    if token:
        headers["Authorization"] = "Bearer " + token
    opener = build_opener(NoRedirect())
    request = Request(url, headers=headers, method="GET")
    try:
        with opener.open(request, timeout=timeout) as response:
            if response.status != 200:
                raise ValueError("readiness returned non-200")
            raw = response.read(MAX_RESPONSE + 1)
    except HTTPError as error:
        error.close()
        raise ValueError("readiness endpoint refused request") from None
    except (OSError, URLError) as error:
        raise ValueError("readiness endpoint unavailable") from None
    if len(raw) > MAX_RESPONSE:
        raise ValueError("oversized readiness response")
    try:
        value = json.loads(raw)
    except (ValueError, UnicodeError):
        raise ValueError("malformed readiness JSON") from None
    if not isinstance(value, dict):
        raise ValueError("readiness is not an object")
    return value


def run(args: argparse.Namespace) -> dict:
    validate_target(args.endpoint, args.allow_remote)
    if not re.fullmatch(r"\d{4}-\d{2}-\d{2}", args.date):
        raise ValueError("date must be YYYY-MM-DD")
    from datetime import date
    date.fromisoformat(args.date)
    if (not re.fullmatch(r"[a-zA-Z0-9:._-]{8,200}", args.expected_bundle)
            or "demo" in args.expected_bundle.lower()
            or not re.fullmatch(r"[a-f0-9]{40}", args.expected_source)):
        raise ValueError("real runs require non-demo bundle and exact source revision pins")
    if not 1 <= args.timeout <= 20:
        raise ValueError("timeout must be 1..20 seconds")
    token = os.environ.get(args.token_env, "") if args.token_env else ""
    if args.token_env and not token:
        raise ValueError("token environment variable missing")
    if args.token_env and args.token_env.startswith("BCM_PUBLIC"):
        raise ValueError("public page token variables are forbidden")
    # Input must be precisely one JSONL sample. Never silently choose an
    # arbitrary location from a full national grid as proof of readiness.
    source = Path(args.canary)
    data = source.read_bytes()
    if len(data) > 65536:
        raise ValueError("canary input oversized")
    rows = [line for line in data.splitlines() if line.strip()]
    if len(rows) != 1:
        raise ValueError("canary input must contain exactly one sample")
    sample = validate_sample(json.loads(rows[0]))
    if "point" not in sample:
        raise ValueError("live canary requires an explicit reviewed coordinate")
    if sample.get("applicability_date", args.date) != args.date:
        raise ValueError("canary effective date mismatch")
    origin = args.endpoint.removesuffix(PATH)
    # Unauthenticated synthetic service cannot pass this explicit source-backed
    # adapter check. No bulk requests are issued if readiness is unsupported.
    health = json_get(origin + "/health", token, args.timeout)
    if health.get("status") != "ok" or health.get("data_scope") != "private_provider_adapter":
        raise ValueError("canonical live provider adapter not established")
    readiness = json_get(origin + "/v1/readiness", token, args.timeout)
    if readiness.get("status") != "ready":
        raise ValueError("private provider readiness not established")
    result = query(sample, args.endpoint, args.date, "", args.timeout, token,
                   RateLimiter(1), args.expected_bundle, args.expected_source)
    if result["category"] not in {"verified", "conditional", "not_verified"}:
        raise ValueError("live point canary failed identity or availability verification")
    return {
        "schema_version": "1",
        "category": "runtime_canary_only",
        "canary_input_sha256": sha256(data).hexdigest(),
        "expected_runtime_identity": {
            "bundle_id": args.expected_bundle,
            "source_commit": args.expected_source,
        },
        "sample_id": sample["id"],
        "http_status": result["http_status"],
        "project_verdict": result["verdict"],
        "jurisdiction_id": result["jurisdiction_id"],
        "code_families": result["code_families"],
        "latency_ms": result["latency_ms"],
        "legal_accuracy_reviewed": False,
        "bulk_run_authorized": False,
    }


def parser():
    p = argparse.ArgumentParser(description=__doc__)
    p.add_argument("--endpoint", required=True, help="HTTPS origin plus " + PATH)
    p.add_argument("--canary", required=True, help="single reviewed geographic JSONL query")
    p.add_argument("--date", required=True)
    p.add_argument("--expected-bundle", required=True)
    p.add_argument("--expected-source", required=True)
    p.add_argument("--token-env", default="")
    p.add_argument("--timeout", type=float, default=8)
    p.add_argument("--allow-remote", action="store_true")
    p.add_argument("--output", required=True, help="sanitized JSON canary receipt")
    return p


if __name__ == "__main__":
    args = parser().parse_args()
    try:
        result = run(args)
        from scripts.bulk_query import atomic_json
        atomic_json(Path(args.output), result)
        print("live canary transport/identity accepted; legal review not established")
    except (OSError, ValueError, TypeError) as error:
        raise SystemExit("live canary blocked: " + str(error)) from None
