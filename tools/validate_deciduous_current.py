#!/usr/bin/env python3
"""Run the frozen archaeology validator against the current repository inventory."""

from __future__ import annotations

import sys
from pathlib import Path

import validate_deciduous_archaeology as archaeology

archaeology.EXPECTED_PROFILES = {
    "colorado.json",
    "florida.json",
    "new-jersey.json",
    "north-carolina.json",
    "oregon.json",
    "virginia.json",
}


def validate_current_readiness_inventory() -> None:
    root = Path(__file__).resolve().parents[1]
    actual = {path.stem for path in (root / "reports" / "production").glob("*.json")}
    expected = {"colorado", "florida", "new-jersey"}
    if actual != expected:
        raise AssertionError(
            f"production-scope inventory mismatch: expected {sorted(expected)}, found {sorted(actual)}"
        )


if __name__ == "__main__":
    validate_current_readiness_inventory()
    raise SystemExit(archaeology.main())
