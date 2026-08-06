from __future__ import annotations

import argparse
import json
import socket
import sys
from datetime import date, timedelta
from pathlib import Path
from typing import Any, Callable
from urllib.error import HTTPError, URLError
from urllib.request import Request, urlopen

REQUIRED_FIXTURE_KINDS = {
    "incorporated",
    "unincorporated",
    "special_project",
    "missing_evidence",
    "conflict",
    "historical",
}
REQUIRED_PILOT_SLUGS = {"colorado", "florida", "new-jersey"}
RESOLVED_STATUSES = {"resolved", "partially_resolved"}
VERIFIED = "verified"
SOURCE_OK = "available"
NON_PRIMARY_SOURCE_KINDS = {
    "session_law_summary",
    "secondary",
    "news",
    "blog",
    "vendor_summary",
}


class ReadinessError(ValueError):
    pass


def _parse_date(value: Any, owner: str) -> date:
    if not isinstance(value, str):
        raise ReadinessError(f"{owner} must be an ISO date")
    try:
        return date.fromisoformat(value)
    except ValueError as exc:
        raise ReadinessError(f"{owner} must be an ISO date") from exc


def _string_list(
    value: Any, owner: str, *, required: bool = True
) -> list[str]:
    if not isinstance(value, list) or any(
        not isinstance(item, str) or not item for item in value
    ):
        raise ReadinessError(f"{owner} must be an array of non-empty strings")
    if required and not value:
        raise ReadinessError(f"{owner} must not be empty")
    if len(value) != len(set(value)):
        raise ReadinessError(f"{owner} must contain unique values")
    return value


def _verification_status(value: Any) -> str | None:
    return value.get("status") if isinstance(value, dict) else None


def validate_manifest(
    manifest: dict[str, Any],
    profile: dict[str, Any],
    pack: dict[str, Any],
) -> None:
    if manifest.get("schema_version") != "1.0":
        raise ReadinessError("manifest schema_version must be '1.0'")
    state_id = manifest.get("state_id")
    if state_id != profile.get("state_id") or state_id != pack.get("state_id"):
        raise ReadinessError(
            "manifest, profile, and rule pack state_id values must match"
        )

    scope = manifest.get("scope")
    if not isinstance(scope, dict):
        raise ReadinessError("scope must be an object")
    families = set(
        _string_list(scope.get("code_families"), "scope.code_families")
    )
    start = _parse_date(scope.get("start_date"), "scope.start_date")
    end = _parse_date(scope.get("end_date"), "scope.end_date")
    if start > end:
        raise ReadinessError("scope.start_date must not follow scope.end_date")

    sources = {
        item.get("id")
        for item in profile.get("sources", [])
        if isinstance(item, dict) and isinstance(item.get("id"), str)
    }
    primary = set(
        _string_list(manifest.get("primary_source_ids"), "primary_source_ids")
    )
    unknown_primary = sorted(primary - sources)
    if unknown_primary:
        raise ReadinessError(
            "primary_source_ids reference unknown sources: "
            + ", ".join(unknown_primary)
        )

    required_kinds = set(
        _string_list(
            manifest.get("required_fixture_kinds"),
            "required_fixture_kinds",
        )
    )
    if required_kinds != REQUIRED_FIXTURE_KINDS:
        missing = sorted(REQUIRED_FIXTURE_KINDS - required_kinds)
        extra = sorted(required_kinds - REQUIRED_FIXTURE_KINDS)
        details: list[str] = []
        if missing:
            details.append("missing " + ", ".join(missing))
        if extra:
            details.append("unknown " + ", ".join(extra))
        suffix = ": " + "; ".join(details) if details else ""
        raise ReadinessError(
            "required_fixture_kinds must exactly match the issue #31 "
            f"acceptance classes{suffix}"
        )

    adoptions = {
        item.get("id")
        for item in profile.get("adoptions", [])
        if isinstance(item, dict) and isinstance(item.get("id"), str)
    }
    timelines = manifest.get("timelines")
    if not isinstance(timelines, list):
        raise ReadinessError("timelines must be an array")
    for index, timeline in enumerate(timelines):
        owner = f"timelines[{index}]"
        if not isinstance(timeline, dict):
            raise ReadinessError(f"{owner} must be an object")
        if timeline.get("code_family") not in families:
            raise ReadinessError(f"{owner}.code_family must be within scope")
        segment_start = _parse_date(
            timeline.get("start_date"), f"{owner}.start_date"
        )
        raw_end = timeline.get("end_date")
        if raw_end is not None:
            segment_end = _parse_date(raw_end, f"{owner}.end_date")
            if segment_start > segment_end:
                raise ReadinessError(
                    f"{owner}.start_date must not follow end_date"
                )
        adoption_refs = set(
            _string_list(
                timeline.get("adoption_ids"), f"{owner}.adoption_ids"
            )
        )
        unknown_adoptions = sorted(adoption_refs - adoptions)
        if unknown_adoptions:
            raise ReadinessError(
                f"{owner} references unknown adoptions: "
                + ", ".join(unknown_adoptions)
            )
        source_refs = set(
            _string_list(timeline.get("source_ids"), f"{owner}.source_ids")
        )
        unknown_sources = sorted(source_refs - sources)
        if unknown_sources:
            raise ReadinessError(
                f"{owner} references unknown sources: "
                + ", ".join(unknown_sources)
            )

    registered = set(
        _string_list(
            pack.get("resolver_fixture_ids", []),
            "resolver_fixture_ids",
            required=False,
        )
    )
    fixtures = manifest.get("fixtures")
    if not isinstance(fixtures, list):
        raise ReadinessError("fixtures must be an array")
    fixture_ids: set[str] = set()
    for index, fixture in enumerate(fixtures):
        owner = f"fixtures[{index}]"
        if not isinstance(fixture, dict):
            raise ReadinessError(f"{owner} must be an object")
        fixture_id = fixture.get("id")
        if not isinstance(fixture_id, str) or not fixture_id:
            raise ReadinessError(f"{owner}.id is required")
        if fixture_id in fixture_ids:
            raise ReadinessError(f"duplicate fixture id {fixture_id}")
        fixture_ids.add(fixture_id)
        if fixture.get("kind") not in REQUIRED_FIXTURE_KINDS:
            raise ReadinessError(f"{owner}.kind is invalid")
        fixture_families = set(
            _string_list(
                fixture.get("code_families"), f"{owner}.code_families"
            )
        )
        if not fixture_families.issubset(families):
            raise ReadinessError(f"{owner}.code_families must be within scope")
        _parse_date(
            fixture.get("applicability_date"),
            f"{owner}.applicability_date",
        )
        if not isinstance(fixture.get("production_supported"), bool):
            raise ReadinessError(
                f"{owner}.production_supported must be boolean"
            )
        source_refs = set(
            _string_list(fixture.get("source_ids"), f"{owner}.source_ids")
        )
        local_refs = set(
            _string_list(
                fixture.get("local_source_ids", []),
                f"{owner}.local_source_ids",
                required=False,
            )
        )
        unknown_sources = sorted((source_refs | local_refs) - sources)
        if unknown_sources:
            raise ReadinessError(
                f"{owner} references unknown sources: "
                + ", ".join(unknown_sources)
            )
        if not local_refs.issubset(source_refs):
            raise ReadinessError(
                f"{owner}.local_source_ids must also appear in source_ids"
            )
        if fixture_id not in registered:
            raise ReadinessError(
                f"fixture {fixture_id} is not registered in resolver_fixture_ids"
            )


def _timeline_signals(
    profile: dict[str, Any],
    manifest: dict[str, Any],
    primary: set[str],
) -> list[str]:
    signals: list[str] = []
    scope = manifest["scope"]
    scope_start = _parse_date(scope["start_date"], "scope.start_date")
    scope_end = _parse_date(scope["end_date"], "scope.end_date")
    adoptions = {
        item["id"]: item
        for item in profile.get("adoptions", [])
        if isinstance(item, dict) and isinstance(item.get("id"), str)
    }

    for family in scope["code_families"]:
        if not any(
            item.get("code_family") == family for item in adoptions.values()
        ):
            signals.append(f"scope_family_missing_adoption:{family}")
        segments = sorted(
            [
                item
                for item in manifest["timelines"]
                if item.get("code_family") == family
            ],
            key=lambda item: item["start_date"],
        )
        if not segments:
            signals.append(f"timeline_missing:{family}")
            continue
        cursor = scope_start
        for segment in segments:
            start = _parse_date(
                segment["start_date"], "timeline.start_date"
            )
            raw_end = segment.get("end_date")
            end = (
                _parse_date(raw_end, "timeline.end_date")
                if raw_end
                else scope_end
            )
            if end < scope_start or start > scope_end:
                continue
            start = max(start, scope_start)
            end = min(end, scope_end)
            if start > cursor:
                signals.append(
                    f"timeline_gap:{family}:{cursor.isoformat()}:"
                    f"{(start - timedelta(days=1)).isoformat()}"
                )
            if _verification_status(segment.get("verification")) != VERIFIED:
                signals.append(
                    f"timeline_not_verified:{family}:{segment['start_date']}"
                )
            if not set(segment.get("source_ids", [])).issubset(primary):
                signals.append(
                    f"timeline_source_not_primary:{family}:"
                    f"{segment['start_date']}"
                )
            for adoption_id in segment.get("adoption_ids", []):
                adoption = adoptions.get(adoption_id)
                if adoption and _verification_status(
                    adoption.get("verification")
                ) != VERIFIED:
                    signals.append(f"adoption_not_verified:{adoption_id}")
                if adoption and not set(
                    adoption.get("source_ids", [])
                ).issubset(primary):
                    signals.append(
                        f"adoption_source_not_primary:{adoption_id}"
                    )
            cursor = max(cursor, end + timedelta(days=1))
            if cursor > scope_end:
                break
        if cursor <= scope_end:
            signals.append(
                f"timeline_gap:{family}:{cursor.isoformat()}:"
                f"{scope_end.isoformat()}"
            )
    return signals


def audit_state(
    profile: dict[str, Any],
    pack: dict[str, Any],
    manifest: dict[str, Any],
    *,
    as_of: date | None = None,
    stale_after_days: int = 365,
) -> dict[str, Any]:
    validate_manifest(manifest, profile, pack)
    as_of = as_of or date.today()
    signals: list[str] = []
    if profile.get("status") != VERIFIED:
        signals.append("profile_not_verified")
    if _verification_status(profile.get("verification")) != VERIFIED:
        signals.append("profile_verification_incomplete")

    primary = set(manifest["primary_source_ids"])
    sources = {
        item.get("id"): item
        for item in profile.get("sources", [])
        if isinstance(item, dict) and isinstance(item.get("id"), str)
    }
    for source_id in sorted(primary):
        source = sources.get(source_id, {})
        if source.get("kind") in NON_PRIMARY_SOURCE_KINDS:
            signals.append(
                f"source_not_primary:{source_id}:{source.get('kind')}"
            )

    source_health = pack.get("source_health", {})
    for source_id in sorted(primary):
        health = source_health.get(source_id)
        if not isinstance(health, dict):
            signals.append(f"primary_source_health_missing:{source_id}")
            continue
        if health.get("availability") != SOURCE_OK:
            availability = health.get("availability", "unknown")
            signals.append(
                f"primary_source_unhealthy:{source_id}:{availability}"
            )
        try:
            checked = _parse_date(
                health.get("last_checked_at"),
                f"source_health.{source_id}.last_checked_at",
            )
        except ReadinessError:
            signals.append(f"primary_source_check_date_missing:{source_id}")
        else:
            if (as_of - checked).days > stale_after_days:
                signals.append(f"primary_source_stale:{source_id}")

    signals.extend(_timeline_signals(profile, manifest, primary))

    registered = set(pack.get("resolver_fixture_ids", []))
    fixtures = manifest["fixtures"]
    fixture_ids = {fixture["id"] for fixture in fixtures}
    for fixture_id in sorted(registered - fixture_ids):
        signals.append(f"registered_fixture_undefined:{fixture_id}")
    for fixture_id in sorted(fixture_ids - registered):
        signals.append(f"fixture_not_registered:{fixture_id}")

    present_kinds = {fixture["kind"] for fixture in fixtures}
    for kind in sorted(REQUIRED_FIXTURE_KINDS - present_kinds):
        signals.append(f"fixture_kind_missing:{kind}")

    for fixture in fixtures:
        fixture_id = fixture["id"]
        if _verification_status(fixture.get("verification")) != VERIFIED:
            signals.append(f"fixture_not_verified:{fixture_id}")
        fixture_sources = set(fixture.get("source_ids", []))
        if not fixture_sources.issubset(primary):
            signals.append(f"fixture_source_not_primary:{fixture_id}")
        if fixture.get("production_supported"):
            if fixture.get("expected_status") not in RESOLVED_STATUSES:
                signals.append(f"production_fixture_unresolved:{fixture_id}")
            if fixture.get("kind") in {
                "incorporated",
                "unincorporated",
                "historical",
            }:
                local_sources = set(fixture.get("local_source_ids", []))
                if not local_sources or not local_sources.issubset(primary):
                    signals.append(
                        f"fixture_local_evidence_missing:{fixture_id}"
                    )
        elif (
            fixture.get("kind") == "missing_evidence"
            and fixture.get("expected_status") != "insufficient_evidence"
        ):
            signals.append(f"negative_fixture_status_mismatch:{fixture_id}")
        elif (
            fixture.get("kind") == "conflict"
            and fixture.get("expected_status") != "conflicting"
        ):
            signals.append(f"negative_fixture_status_mismatch:{fixture_id}")

    for claim in pack.get("claims", []):
        if not isinstance(claim, dict):
            continue
        claim_id = claim.get("id", "unknown")
        if claim.get("status") == "conflicting":
            signals.append(f"unresolved_conflicting_claim:{claim_id}")
        if claim.get("status") == "supported":
            if _verification_status(claim.get("verification")) != VERIFIED:
                signals.append(f"supported_claim_not_verified:{claim_id}")
            if not set(claim.get("source_ids", [])).issubset(primary):
                signals.append(
                    f"supported_claim_source_not_primary:{claim_id}"
                )

    unique_signals = sorted(set(signals))
    return {
        "state_id": profile.get("state_id"),
        "state_name": profile.get("state_name"),
        "scope": manifest.get("scope"),
        "fixture_count": len(fixtures),
        "production_fixture_count": sum(
            bool(item.get("production_supported")) for item in fixtures
        ),
        "timeline_segment_count": len(manifest.get("timelines", [])),
        "primary_source_count": len(primary),
        "readiness": (
            "production_ready" if not unique_signals else "needs_review"
        ),
        "premature_verified_profile": (
            profile.get("status") == VERIFIED and bool(unique_signals)
        ),
        "signals": unique_signals,
    }


def audit_repository(
    compiled_root: Path,
    manifests_root: Path,
    *,
    as_of: date | None = None,
    stale_after_days: int = 365,
    expected_state_slugs: set[str] | None = None,
) -> dict[str, Any]:
    manifests = {
        path.stem: path for path in manifests_root.glob("*.json")
    }
    slugs = set(expected_state_slugs or REQUIRED_PILOT_SLUGS)
    slugs.update(manifests)
    rows: list[dict[str, Any]] = []
    for slug in sorted(slugs):
        manifest_path = manifests.get(slug, manifests_root / f"{slug}.json")
        profile_path = compiled_root / f"{slug}.json"
        pack_path = compiled_root / "rules" / f"{slug}.json"
        profile: dict[str, Any] | None = None
        try:
            profile = json.loads(profile_path.read_text(encoding="utf-8"))
            pack = json.loads(pack_path.read_text(encoding="utf-8"))
            if not manifest_path.exists():
                row = {
                    "state_id": profile.get("state_id"),
                    "state_name": profile.get("state_name", slug),
                    "scope": None,
                    "fixture_count": 0,
                    "production_fixture_count": 0,
                    "timeline_segment_count": 0,
                    "primary_source_count": 0,
                    "readiness": "blocked",
                    "premature_verified_profile": (
                        profile.get("status") == VERIFIED
                    ),
                    "signals": ["manifest_missing"],
                    "manifest": manifest_path.as_posix(),
                }
            else:
                manifest = json.loads(
                    manifest_path.read_text(encoding="utf-8")
                )
                row = audit_state(
                    profile,
                    pack,
                    manifest,
                    as_of=as_of,
                    stale_after_days=stale_after_days,
                )
                row["manifest"] = manifest_path.as_posix()
        except (OSError, json.JSONDecodeError, ReadinessError) as exc:
            row = {
                "state_id": profile.get("state_id") if profile else None,
                "state_name": (
                    profile.get("state_name", slug) if profile else slug
                ),
                "scope": None,
                "fixture_count": 0,
                "production_fixture_count": 0,
                "timeline_segment_count": 0,
                "primary_source_count": 0,
                "readiness": "blocked",
                "premature_verified_profile": bool(
                    profile and profile.get("status") == VERIFIED
                ),
                "signals": ["manifest_or_artifact_invalid"],
                "error": str(exc),
                "manifest": manifest_path.as_posix(),
            }
        rows.append(row)

    return {
        "schema_version": "1.0",
        "as_of": (as_of or date.today()).isoformat(),
        "summary": {
            "states": len(rows),
            "production_ready": sum(
                row["readiness"] == "production_ready" for row in rows
            ),
            "needs_review": sum(
                row["readiness"] == "needs_review" for row in rows
            ),
            "blocked": sum(
                row["readiness"] == "blocked" for row in rows
            ),
            "premature_verified_profiles": sum(
                bool(row.get("premature_verified_profile")) for row in rows
            ),
        },
        "states": rows,
    }


def readiness_markdown(report: dict[str, Any]) -> str:
    lines = [
        "# Pilot Production Readiness",
        "",
        "This audit evaluates the explicit production scope, primary-source "
        "set, complete adoption timelines, and evidence-backed resolver "
        "fixtures required by issue #31.",
        "",
        "| State | Scope | Primary sources | Timelines | Fixtures | "
        "Production fixtures | Readiness | Gaps |",
        "| --- | --- | ---: | ---: | ---: | ---: | --- | --- |",
    ]
    for row in report["states"]:
        scope = row.get("scope") or {}
        families = ", ".join(scope.get("code_families", [])) or "n/a"
        period = (
            f"{scope.get('start_date')} to {scope.get('end_date')}"
            if scope
            else "n/a"
        )
        gaps = ", ".join(row.get("signals", [])) or "none"
        lines.append(
            f"| {row.get('state_name') or row.get('state_id') or 'Unknown'} | "
            f"{families} ({period}) | "
            f"{row.get('primary_source_count', 0)} | "
            f"{row.get('timeline_segment_count', 0)} | "
            f"{row.get('fixture_count', 0)} | "
            f"{row.get('production_fixture_count', 0)} | "
            f"{row.get('readiness')} | {gaps} |"
        )
    lines.extend(
        [
            "",
            "## Promotion rule",
            "",
            "A pilot is `production_ready` only when every scoped code "
            "family has a continuous verified timeline, every required "
            "fixture class is represented by a verified fixture, supported "
            "local fixtures cite local primary sources, all primary sources "
            "are healthy and current, and no supported claim remains "
            "conflicting or partially verified.",
            "",
        ]
    )
    return "\n".join(lines)


def check_source(
    source_id: str,
    url: str,
    *,
    timeout: float = 20.0,
    opener: Callable[..., Any] = urlopen,
    checked_at: date | None = None,
) -> dict[str, Any]:
    request = Request(
        url,
        headers={
            "User-Agent": "building-code-map-source-health/1.0",
            "Range": "bytes=0-0",
            "Accept": (
                "text/html,application/json,application/pdf,*/*;q=0.1"
            ),
        },
    )
    result: dict[str, Any] = {
        "source_id": source_id,
        "url": url,
        "checked_at": (checked_at or date.today()).isoformat(),
    }
    try:
        with opener(request, timeout=timeout) as response:
            response.read(1)
            status = int(getattr(response, "status", 200))
            final_url = response.geturl()
            result.update(
                {
                    "status_code": status,
                    "final_url": final_url,
                    "availability": (
                        "moved" if final_url != url else "available"
                    ),
                }
            )
    except HTTPError as exc:
        result.update(
            {
                "status_code": exc.code,
                "final_url": exc.geturl(),
                "availability": "unavailable",
                "error": str(exc),
            }
        )
    except (URLError, TimeoutError, socket.timeout, OSError) as exc:
        result.update(
            {
                "status_code": None,
                "final_url": None,
                "availability": "unavailable",
                "error": str(exc),
            }
        )
    return result


def check_repository_sources(
    compiled_root: Path,
    *,
    timeout: float = 20.0,
    opener: Callable[..., Any] = urlopen,
) -> dict[str, Any]:
    records: list[dict[str, Any]] = []
    for profile_path in sorted(compiled_root.glob("*.json")):
        profile = json.loads(profile_path.read_text(encoding="utf-8"))
        for source in profile.get("sources", []):
            if (
                not isinstance(source, dict)
                or not source.get("id")
                or not source.get("url")
            ):
                continue
            record = check_source(
                source["id"],
                source["url"],
                timeout=timeout,
                opener=opener,
            )
            record["state_id"] = profile.get("state_id")
            record["title"] = source.get("title")
            records.append(record)
    return {
        "schema_version": "1.0",
        "checked_at": date.today().isoformat(),
        "summary": {
            "sources": len(records),
            "available": sum(
                item["availability"] == "available" for item in records
            ),
            "moved": sum(
                item["availability"] == "moved" for item in records
            ),
            "unavailable": sum(
                item["availability"] == "unavailable" for item in records
            ),
        },
        "sources": records,
    }


def source_health_markdown(report: dict[str, Any]) -> str:
    lines = [
        "# Regulatory Source Health",
        "",
        "| State | Source | Availability | HTTP | Final URL |",
        "| --- | --- | --- | ---: | --- |",
    ]
    for item in report["sources"]:
        lines.append(
            f"| {item.get('state_id', '')} | "
            f"{item.get('source_id', '')} | "
            f"{item.get('availability', 'unknown')} | "
            f"{item.get('status_code') or ''} | "
            f"{item.get('final_url') or ''} |"
        )
    lines.append("")
    return "\n".join(lines)


def _write(path: Path, content: str) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(content, encoding="utf-8")


def parser() -> argparse.ArgumentParser:
    root = argparse.ArgumentParser(
        description=(
            "Audit scoped pilot production readiness and official source "
            "health."
        )
    )
    commands = root.add_subparsers(dest="command", required=True)

    audit = commands.add_parser("audit")
    audit.add_argument(
        "--compiled-root",
        type=Path,
        default=Path("backend/data/regulatory"),
    )
    audit.add_argument(
        "--manifests-root",
        type=Path,
        default=Path("reports/production"),
    )
    audit.add_argument("--json-output", type=Path, required=True)
    audit.add_argument("--markdown-output", type=Path, required=True)
    audit.add_argument("--as-of", type=date.fromisoformat)
    audit.add_argument("--stale-after-days", type=int, default=365)
    audit.add_argument(
        "--fail-on-premature-verification", action="store_true"
    )

    health = commands.add_parser("check-sources")
    health.add_argument(
        "--compiled-root",
        type=Path,
        default=Path("backend/data/regulatory"),
    )
    health.add_argument("--json-output", type=Path, required=True)
    health.add_argument("--markdown-output", type=Path, required=True)
    health.add_argument("--timeout", type=float, default=20.0)
    health.add_argument("--fail-on-unhealthy", action="store_true")
    return root


def main(argv: list[str] | None = None) -> int:
    args = parser().parse_args(argv)
    try:
        if args.command == "audit":
            report = audit_repository(
                args.compiled_root,
                args.manifests_root,
                as_of=args.as_of,
                stale_after_days=args.stale_after_days,
            )
            _write(
                args.json_output,
                json.dumps(report, indent=2) + "\n",
            )
            _write(args.markdown_output, readiness_markdown(report))
            if (
                args.fail_on_premature_verification
                and report["summary"]["premature_verified_profiles"]
            ):
                return 1
        else:
            report = check_repository_sources(
                args.compiled_root,
                timeout=args.timeout,
            )
            _write(
                args.json_output,
                json.dumps(report, indent=2) + "\n",
            )
            _write(
                args.markdown_output,
                source_health_markdown(report),
            )
            unhealthy = (
                report["summary"]["moved"]
                + report["summary"]["unavailable"]
            )
            if args.fail_on_unhealthy and unhealthy:
                return 1
        return 0
    except (OSError, json.JSONDecodeError, ReadinessError) as exc:
        print(exc, file=sys.stderr)
        return 1


if __name__ == "__main__":
    raise SystemExit(main())
