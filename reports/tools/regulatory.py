from __future__ import annotations

import argparse
import json
import re
import sys
from datetime import date
from pathlib import Path
from typing import Any, Iterable

PROFILE_FENCE = re.compile(r"```json\s+jurisdiction-profile\s*\n(?P<body>.*?)\n```", re.DOTALL)
RULE_FENCE = re.compile(r"```json\s+jurisdiction-rules\s*\n(?P<body>.*?)\n```", re.DOTALL)
FRONTMATTER = re.compile(r"\A---\s*\n(?P<body>.*?)\n---\s*\n", re.DOTALL)
DATE = re.compile(r"^\d{4}-\d{2}-\d{2}$")
RESOLUTION_STATUSES = {
    "resolved",
    "partially_resolved",
    "local_record_required",
    "ambiguous",
    "conflicting",
    "insufficient_evidence",
}
PROFILE_STATUSES = {"draft", "partially_verified", "verified", "deprecated"}
VERIFICATION_STATUSES = {"verified", "partially_verified", "needs_review", "unresolved"}
ADOPTION_STATUSES = {"current", "prior", "future", "pending", "superseded", "unknown"}
CLAIM_STATUSES = {"supported", "conflicting", "unknown", "not_applicable", "superseded"}
SOURCE_AVAILABILITY = {"available", "unavailable", "moved", "unknown"}
RULE_GROUPS = ("applicability_rules", "date_rules", "amendment_rules", "enforcement_rules")


class ProfileError(ValueError):
    pass


def _extract(
    text: str,
    pattern: re.Pattern[str],
    label: str,
    source: str,
    required: bool,
) -> dict[str, Any] | None:
    matches = list(pattern.finditer(text))
    if not matches:
        if required:
            raise ProfileError(f"{source}: missing ```json {label} fenced block")
        return None
    if len(matches) != 1:
        raise ProfileError(f"{source}: expected one {label} block, found {len(matches)}")
    try:
        value = json.loads(matches[0].group("body"))
    except json.JSONDecodeError as exc:
        raise ProfileError(
            f"{source}: invalid {label} JSON at line {exc.lineno}: {exc.msg}"
        ) from exc
    if not isinstance(value, dict):
        raise ProfileError(f"{source}: {label} root must be an object")
    return value


def extract_profile(text: str, source: str = "<memory>") -> dict[str, Any]:
    profile = _extract(text, PROFILE_FENCE, "jurisdiction-profile", source, True)
    assert profile is not None
    validate_profile(profile, source)
    return profile


def extract_rule_pack(
    text: str,
    profile: dict[str, Any],
    source: str = "<memory>",
    *,
    required: bool = False,
) -> dict[str, Any] | None:
    pack = _extract(text, RULE_FENCE, "jurisdiction-rules", source, required)
    if pack is not None:
        validate_rule_pack(pack, profile, source)
    return pack


def _list(
    owner: dict[str, Any],
    key: str,
    problems: list[str],
    required: bool = False,
) -> list[Any]:
    value = owner.get(key, [])
    if not isinstance(value, list) or (required and not value):
        problems.append(f"{key} must be {'a non-empty ' if required else 'an '}array")
        return []
    return value


def _ids(items: list[Any], kind: str, problems: list[str]) -> set[str]:
    result: set[str] = set()
    for item in items:
        if not isinstance(item, dict) or not isinstance(item.get("id"), str) or not item["id"]:
            problems.append(f"{kind} id is required")
            continue
        if item["id"] in result:
            problems.append(f"duplicate {kind} id {item['id']}")
        result.add(item["id"])
    return result


def _source_refs(
    item: dict[str, Any],
    owner: str,
    known: set[str],
    problems: list[str],
) -> None:
    refs = item.get("source_ids")
    if not isinstance(refs, list) or not refs:
        problems.append(f"{owner} must cite at least one source")
        return
    for ref in refs:
        if ref not in known:
            problems.append(f"{owner} references unknown source {ref}")


def _date_field(
    value: Any,
    owner: str,
    problems: list[str],
    required: bool = False,
) -> None:
    if value in (None, "") and not required:
        return
    if not isinstance(value, str) or not DATE.fullmatch(value):
        problems.append(f"{owner} must use YYYY-MM-DD")
        return
    try:
        date.fromisoformat(value)
    except ValueError:
        problems.append(f"{owner} must be a valid date")


def _verification(value: Any, owner: str, problems: list[str]) -> None:
    if not isinstance(value, dict) or value.get("status") not in VERIFICATION_STATUSES:
        problems.append(f"{owner} verification status is invalid")
        return
    confidence = value.get("confidence")
    if confidence is not None and (
        not isinstance(confidence, (int, float)) or confidence < 0 or confidence > 1
    ):
        problems.append(f"{owner} verification confidence must be between 0 and 1")


def validate_profile(profile: dict[str, Any], source: str = "<memory>") -> None:
    problems: list[str] = []
    required_strings = (
        "schema_version",
        "profile_id",
        "state_id",
        "state_name",
        "state_abbreviation",
        "state_fips",
        "status",
        "last_verified",
    )
    for key in required_strings:
        if not isinstance(profile.get(key), str) or not profile[key].strip():
            problems.append(f"{key} must be a non-empty string")
    if profile.get("schema_version") != "1.0":
        problems.append("schema_version must be '1.0'")
    if profile.get("status") not in PROFILE_STATUSES:
        problems.append(f"profile status is invalid: {profile.get('status')}")
    if not re.fullmatch(r"\d{2}", str(profile.get("state_fips", ""))):
        problems.append("state_fips must be exactly two digits")
    _date_field(profile.get("last_verified"), "last_verified", problems, required=True)

    sources = _list(profile, "sources", problems, required=True)
    authorities = _list(profile, "authorities", problems, required=True)
    relationships = _list(profile, "relationships", problems)
    adoptions = _list(profile, "adoptions", problems)
    source_ids = _ids(sources, "source", problems)
    authority_ids = _ids(authorities, "authority", problems)
    adoption_ids = _ids(adoptions, "adoption", problems)

    for item in sources:
        if not isinstance(item, dict):
            problems.append("sources entries must be objects")
            continue
        _date_field(
            item.get("accessed_at"),
            f"source {item.get('id')} accessed_at",
            problems,
            required=True,
        )
        _date_field(
            item.get("last_checked_at"),
            f"source {item.get('id')} last_checked_at",
            problems,
        )
        if (
            item.get("availability") is not None
            and item.get("availability") not in SOURCE_AVAILABILITY
        ):
            problems.append(
                f"source {item.get('id')} availability is invalid: {item.get('availability')}"
            )
    for item in authorities:
        if isinstance(item, dict):
            _source_refs(item, f"authority {item.get('id')}", source_ids, problems)
            _verification(item.get("verification"), f"authority {item.get('id')}", problems)
    for item in relationships:
        if isinstance(item, dict):
            _source_refs(item, f"relationship {item.get('id')}", source_ids, problems)
            _verification(
                item.get("verification"),
                f"relationship {item.get('id')}",
                problems,
            )
            if item.get("from_id") not in authority_ids:
                problems.append(
                    f"relationship {item.get('id')} references unknown from_id {item.get('from_id')}"
                )
    for item in adoptions:
        if not isinstance(item, dict):
            continue
        _source_refs(item, f"adoption {item.get('id')}", source_ids, problems)
        _verification(item.get("verification"), f"adoption {item.get('id')}", problems)
        if item.get("status") not in ADOPTION_STATUSES:
            problems.append(
                f"adoption {item.get('id')} status is invalid: {item.get('status')}"
            )
        if item.get("adopting_authority_id") not in authority_ids:
            problems.append(
                f"adoption {item.get('id')} references unknown authority {item.get('adopting_authority_id')}"
            )
        dates = item.get("dates", {})
        if not isinstance(dates, dict):
            problems.append(f"adoption {item.get('id')} dates must be an object")
        else:
            for key, value in dates.items():
                _date_field(value, f"adoption {item.get('id')} dates.{key}", problems)

    defaults = profile.get("defaults")
    if not isinstance(defaults, dict):
        problems.append("defaults must be an object")
        defaults = {}
    policies: list[tuple[str, Any]] = [
        ("defaults.incorporated", defaults.get("incorporated")),
        ("defaults.unincorporated", defaults.get("unincorporated")),
    ]
    for group in ("code_family_overrides", "project_type_overrides"):
        values = profile.get(group, {})
        if not isinstance(values, dict):
            problems.append(f"{group} must be an object")
        else:
            policies.extend((f"{group}.{key}", value) for key, value in values.items())
    for name, policy in policies:
        if not isinstance(policy, dict):
            problems.append(f"{name} must be an object")
            continue
        if policy.get("status") not in RESOLUTION_STATUSES:
            problems.append(f"{name}.status is invalid: {policy.get('status')}")
        _source_refs(policy, name, source_ids, problems)
        for adoption_id in policy.get("adoption_ids", []):
            if adoption_id not in adoption_ids:
                problems.append(f"{name} references unknown adoption {adoption_id}")
        for candidate in policy.get("authority_candidates", []):
            if not isinstance(candidate, dict):
                problems.append(f"{name}.authority_candidates entries must be objects")
                continue
            _source_refs(candidate, f"{name} candidate", source_ids, problems)
            if (
                candidate.get("kind") == "state_authority"
                and candidate.get("authority_id") not in authority_ids
            ):
                problems.append(
                    f"{name} references unknown authority {candidate.get('authority_id')}"
                )
    _verification(profile.get("verification"), "profile", problems)
    if problems:
        raise ProfileError(f"{source}: " + "; ".join(sorted(set(problems))))


def validate_rule_pack(
    pack: dict[str, Any],
    profile: dict[str, Any],
    source: str = "<memory>",
) -> None:
    problems: list[str] = []
    if pack.get("schema_version") != "1.0":
        problems.append("rule pack schema_version must be '1.0'")
    if pack.get("state_id") != profile.get("state_id"):
        problems.append(
            f"rule pack state_id {pack.get('state_id')} does not match profile {profile.get('state_id')}"
        )

    source_ids = {
        item.get("id") for item in profile.get("sources", []) if isinstance(item, dict)
    }
    authority_ids = {
        item.get("id")
        for item in profile.get("authorities", [])
        if isinstance(item, dict)
    }
    subjects = {profile.get("profile_id")}
    subjects.update(authority_ids)
    subjects.update(
        item.get("id")
        for item in profile.get("relationships", [])
        if isinstance(item, dict)
    )
    subjects.update(
        item.get("id")
        for item in profile.get("adoptions", [])
        if isinstance(item, dict)
    )

    health = pack.get("source_health", {})
    if not isinstance(health, dict):
        problems.append("source_health must be an object")
    else:
        for source_id, record in health.items():
            if source_id not in source_ids:
                problems.append(f"source_health references unknown source {source_id}")
            if not isinstance(record, dict):
                problems.append(f"source_health.{source_id} must be an object")
                continue
            _date_field(
                record.get("last_checked_at"),
                f"source_health.{source_id}.last_checked_at",
                problems,
            )
            if (
                record.get("availability") is not None
                and record.get("availability") not in SOURCE_AVAILABILITY
            ):
                problems.append(
                    f"source_health.{source_id}.availability is invalid: {record.get('availability')}"
                )

    rule_ids: set[str] = set()
    for group in RULE_GROUPS:
        for rule in _list(pack, group, problems):
            if not isinstance(rule, dict):
                problems.append(f"{group} entries must be objects")
                continue
            rule_id = rule.get("id")
            if not isinstance(rule_id, str) or not rule_id:
                problems.append(f"{group} id is required")
                continue
            if rule_id in rule_ids:
                problems.append(f"duplicate rule id {rule_id}")
            rule_ids.add(rule_id)
            subjects.add(rule_id)
            _source_refs(rule, f"{group} {rule_id}", source_ids, problems)
            _verification(rule.get("verification"), f"{group} {rule_id}", problems)
            if not isinstance(rule.get("summary"), str) or not rule["summary"].strip():
                problems.append(f"{group} {rule_id} summary is required")
            if group == "applicability_rules" and not rule.get("trigger"):
                problems.append(f"applicability rule {rule_id} trigger is required")
            elif group == "date_rules":
                if not rule.get("rule_type") or not rule.get("trigger"):
                    problems.append(
                        f"date rule {rule_id} rule_type and trigger are required"
                    )
                _date_field(
                    rule.get("start_date"),
                    f"date rule {rule_id} start_date",
                    problems,
                )
                _date_field(
                    rule.get("end_date"),
                    f"date rule {rule_id} end_date",
                    problems,
                )
                if (
                    rule.get("start_date")
                    and rule.get("end_date")
                    and rule["start_date"] > rule["end_date"]
                ):
                    problems.append(
                        f"date rule {rule_id} start_date must not follow end_date"
                    )
            elif group == "amendment_rules":
                if not rule.get("level") or not rule.get("posture"):
                    problems.append(
                        f"amendment rule {rule_id} level and posture are required"
                    )
                if (
                    rule.get("authority_id")
                    and rule.get("authority_id") not in authority_ids
                ):
                    problems.append(
                        f"amendment rule {rule_id} references unknown authority {rule.get('authority_id')}"
                    )
            elif group == "enforcement_rules":
                if (
                    not rule.get("model")
                    or not isinstance(rule.get("entity_kinds"), list)
                    or not rule["entity_kinds"]
                ):
                    problems.append(
                        f"enforcement rule {rule_id} model and entity_kinds are required"
                    )
                for authority_id in rule.get("authority_ids", []):
                    if authority_id not in authority_ids:
                        problems.append(
                            f"enforcement rule {rule_id} references unknown authority {authority_id}"
                        )

    claim_ids: set[str] = set()
    for claim in _list(pack, "claims", problems):
        if not isinstance(claim, dict):
            problems.append("claims entries must be objects")
            continue
        claim_id = claim.get("id")
        if not isinstance(claim_id, str) or not claim_id:
            problems.append("claim id is required")
            continue
        if claim_id in claim_ids:
            problems.append(f"duplicate claim id {claim_id}")
        claim_ids.add(claim_id)
        if claim.get("subject_id") not in subjects:
            problems.append(
                f"claim {claim_id} references unknown subject {claim.get('subject_id')}"
            )
        if not claim.get("field") or claim.get("status") not in CLAIM_STATUSES:
            problems.append(f"claim {claim_id} requires a field and valid status")
        if claim.get("status") == "conflicting" and not claim.get("conflict_group"):
            problems.append(f"conflicting claim {claim_id} requires conflict_group")
        _source_refs(claim, f"claim {claim_id}", source_ids, problems)
        _verification(claim.get("verification"), f"claim {claim_id}", problems)

    fixtures = pack.get("resolver_fixture_ids", [])
    if not isinstance(fixtures, list) or any(
        not isinstance(item, str) or not item for item in fixtures
    ):
        problems.append("resolver_fixture_ids must be an array of non-empty strings")
    if len(fixtures) != len(set(fixtures)):
        problems.append("resolver_fixture_ids must be unique")
    if problems:
        raise ProfileError(f"{source}: " + "; ".join(sorted(set(problems))))


def canonical_json(value: dict[str, Any]) -> str:
    return json.dumps(value, indent=2, ensure_ascii=False) + "\n"


def _load_external_pack(
    profile: dict[str, Any],
    report: Path,
    output_dir: Path | None = None,
) -> dict[str, Any] | None:
    name = _slug(profile["state_name"]) + ".json"
    candidates: list[Path] = []
    if output_dir is not None:
        candidates.append(output_dir / "rules" / name)
    for parent in report.parents:
        candidates.append(
            parent / "backend" / "data" / "regulatory" / "rules" / name
        )
    for candidate in candidates:
        if not candidate.exists():
            continue
        try:
            value = json.loads(candidate.read_text(encoding="utf-8"))
        except json.JSONDecodeError as exc:
            raise ProfileError(
                f"{candidate}: invalid rule pack JSON: {exc.msg}"
            ) from exc
        if not isinstance(value, dict):
            raise ProfileError(f"{candidate}: rule pack root must be an object")
        validate_rule_pack(value, profile, str(candidate))
        return value
    return None


def compile_reports(
    reports: Iterable[Path],
    output_dir: Path,
    check: bool = False,
) -> list[Path]:
    output_dir.mkdir(parents=True, exist_ok=True)
    written: list[Path] = []
    seen: set[str] = set()
    for report in sorted(reports):
        text = report.read_text(encoding="utf-8")
        profile = extract_profile(text, str(report))
        name = _slug(profile["state_name"]) + ".json"
        if name in seen:
            raise ProfileError(f"duplicate output profile {name}")
        seen.add(name)
        profile_path = output_dir / name
        write_or_check(profile_path, canonical_json(profile), check)
        written.append(profile_path)
        pack = extract_rule_pack(text, profile, str(report), required=False)
        if pack is None:
            pack = _load_external_pack(profile, report, output_dir)
        if pack is not None:
            rule_path = output_dir / "rules" / name
            write_or_check(rule_path, canonical_json(pack), check)
            written.append(rule_path)
    return written


def discover_reports(root: Path) -> list[Path]:
    return [
        path
        for path in sorted(root.rglob("*.md"))
        if path.name not in {"_template.md", "README.md"}
        and "generated" not in path.parts
    ]


def parse_frontmatter(text: str) -> dict[str, str]:
    match = FRONTMATTER.search(text)
    if not match:
        return {}
    values: dict[str, str] = {}
    section: str | None = None
    mapping = {
        ("state", "state_id"): "state_id",
        ("state", "name"): "state_name",
        ("state", "abbreviation"): "abbreviation",
        ("report", "status"): "report_status",
        ("report", "last_updated"): "last_updated",
        ("report", "last_verified"): "last_verified",
        ("risk", "open_questions_count"): "open_questions",
    }
    for raw in match.group("body").splitlines():
        line = raw.rstrip()
        if not line or line.lstrip().startswith("#"):
            continue
        indent = len(line) - len(line.lstrip())
        stripped = line.strip()
        if indent == 0 and stripped.endswith(":"):
            section = stripped[:-1]
            continue
        if ":" not in stripped:
            continue
        key, value = stripped.split(":", 1)
        canonical = mapping.get((section, key))
        if canonical:
            values[canonical] = value.strip().strip("\"'")
    return values


def _policies(profile: dict[str, Any]) -> Iterable[dict[str, Any]]:
    for key in ("incorporated", "unincorporated"):
        value = profile.get("defaults", {}).get(key)
        if isinstance(value, dict):
            yield value
    for group in ("code_family_overrides", "project_type_overrides"):
        for value in profile.get(group, {}).values():
            if isinstance(value, dict):
                yield value


def _int(value: str | None) -> int | None:
    try:
        return int(value) if value not in (None, "") else None
    except ValueError:
        return None


def _parse_date(value: Any) -> date | None:
    try:
        return date.fromisoformat(value) if isinstance(value, str) else None
    except ValueError:
        return None


def build_coverage(
    reports: Iterable[Path],
    root: Path,
    *,
    as_of: date | None = None,
    stale_after_days: int = 365,
) -> list[dict[str, Any]]:
    as_of = as_of or date.today()
    rows: list[dict[str, Any]] = []
    for path in sorted(reports):
        text = path.read_text(encoding="utf-8")
        metadata = parse_frontmatter(text)
        relative = path.relative_to(root).as_posix() if path.is_relative_to(root) else str(path)
        row: dict[str, Any] = {
            "report": relative,
            "state_id": metadata.get("state_id"),
            "state_name": metadata.get("state_name"),
            "abbreviation": metadata.get("abbreviation"),
            "report_status": metadata.get("report_status"),
            "last_updated": metadata.get("last_updated"),
            "last_verified": metadata.get("last_verified"),
            "open_questions": _int(metadata.get("open_questions")),
            "profile_present": False,
            "profile_valid": False,
            "profile_status": None,
            "rule_pack_present": False,
            "rule_pack_valid": False,
            "source_count": 0,
            "authority_count": 0,
            "adoption_count": 0,
            "code_families": [],
            "stale_source_count": 0,
            "unavailable_source_count": 0,
            "applicability_rule_count": 0,
            "date_rule_count": 0,
            "amendment_rule_count": 0,
            "enforcement_rule_count": 0,
            "claim_count": 0,
            "conflicting_claim_count": 0,
            "resolver_fixture_count": 0,
            "readiness": "blocked",
            "unresolved_signals": [],
            "error": None,
        }
        try:
            profile = extract_profile(text, str(path))
        except ProfileError as exc:
            row["profile_present"] = "missing" not in str(exc)
            row["unresolved_signals"] = ["profile_missing_or_invalid"]
            row["error"] = str(exc)
            rows.append(row)
            continue
        try:
            pack = extract_rule_pack(text, profile, str(path), required=False)
            if pack is None:
                pack = _load_external_pack(profile, path)
        except ProfileError as exc:
            row["profile_present"] = True
            row["profile_valid"] = True
            row["rule_pack_present"] = True
            row["unresolved_signals"] = ["rule_pack_invalid"]
            row["error"] = str(exc)
            rows.append(row)
            continue

        signals: list[str] = []
        policies = list(_policies(profile))
        if profile.get("status") != "verified":
            signals.append("profile_not_verified")
        if profile.get("verification", {}).get("status") != "verified":
            signals.append("verification_incomplete")
        if any(
            policy.get("status")
            in RESOLUTION_STATUSES - {"resolved", "partially_resolved"}
            for policy in policies
        ):
            signals.append("resolver_follow_up_required")
        if not profile.get("adoptions"):
            signals.append("no_state_adoptions_recorded")
        if row["open_questions"]:
            signals.append("open_questions")
        if pack is None:
            signals.append("rule_pack_missing")
            counts = {group: 0 for group in RULE_GROUPS}
            claims: list[Any] = []
            fixtures = 0
            stale = unavailable = conflicting = 0
        else:
            counts = {group: len(pack.get(group, [])) for group in RULE_GROUPS}
            if any(counts[group] == 0 for group in RULE_GROUPS):
                signals.append("rule_family_missing")
            claims = pack.get("claims", [])
            fixtures = len(pack.get("resolver_fixture_ids", []))
            if fixtures == 0:
                signals.append("resolver_fixtures_missing")
            conflicting = sum(
                isinstance(claim, dict) and claim.get("status") == "conflicting"
                for claim in claims
            )
            if conflicting:
                signals.append("conflicting_claims")
            stale = unavailable = 0
            for health in pack.get("source_health", {}).values():
                if not isinstance(health, dict):
                    continue
                checked = _parse_date(health.get("last_checked_at"))
                if checked is None or (as_of - checked).days > stale_after_days:
                    stale += 1
                if health.get("availability") in {"unavailable", "moved"}:
                    unavailable += 1
            if stale:
                signals.append("stale_sources")
            if unavailable:
                signals.append("unavailable_sources")

        pilot_ready = (
            bool(pack)
            and all(counts[group] > 0 for group in RULE_GROUPS)
            and fixtures > 0
        )
        production_ready = (
            pilot_ready
            and profile.get("status") == "verified"
            and profile.get("verification", {}).get("status") == "verified"
            and stale == 0
            and unavailable == 0
            and conflicting == 0
            and not row["open_questions"]
            and not any(
                policy.get("status") not in {"resolved", "partially_resolved"}
                for policy in policies
            )
        )
        readiness = (
            "production_ready"
            if production_ready
            else "pilot_ready"
            if pilot_ready
            else "needs_review"
        )
        row.update(
            {
                "state_id": profile["state_id"],
                "state_name": profile["state_name"],
                "abbreviation": profile["state_abbreviation"],
                "profile_present": True,
                "profile_valid": True,
                "profile_status": profile["status"],
                "rule_pack_present": pack is not None,
                "rule_pack_valid": pack is not None,
                "source_count": len(profile["sources"]),
                "authority_count": len(profile["authorities"]),
                "adoption_count": len(profile.get("adoptions", [])),
                "code_families": sorted(
                    {adoption["code_family"] for adoption in profile.get("adoptions", [])}
                ),
                "stale_source_count": stale,
                "unavailable_source_count": unavailable,
                "applicability_rule_count": counts["applicability_rules"],
                "date_rule_count": counts["date_rules"],
                "amendment_rule_count": counts["amendment_rules"],
                "enforcement_rule_count": counts["enforcement_rules"],
                "claim_count": len(claims),
                "conflicting_claim_count": conflicting,
                "resolver_fixture_count": fixtures,
                "readiness": readiness,
                "unresolved_signals": sorted(set(signals)),
            }
        )
        rows.append(row)
    return rows


def coverage_json(rows: list[dict[str, Any]]) -> str:
    return (
        json.dumps(
            {
                "schema_version": "1.0",
                "summary": {
                    "reports": len(rows),
                    "profiles_present": sum(
                        bool(row["profile_present"]) for row in rows
                    ),
                    "profiles_valid": sum(bool(row["profile_valid"]) for row in rows),
                    "rule_packs_valid": sum(
                        bool(row["rule_pack_valid"]) for row in rows
                    ),
                    "pilot_ready": sum(
                        row["readiness"] == "pilot_ready" for row in rows
                    ),
                    "production_ready": sum(
                        row["readiness"] == "production_ready" for row in rows
                    ),
                    "reports_requiring_follow_up": sum(
                        bool(row["unresolved_signals"]) for row in rows
                    ),
                    "stale_sources": sum(row["stale_source_count"] for row in rows),
                    "unavailable_sources": sum(
                        row["unavailable_source_count"] for row in rows
                    ),
                    "conflicting_claims": sum(
                        row["conflicting_claim_count"] for row in rows
                    ),
                },
                "states": rows,
            },
            indent=2,
            ensure_ascii=False,
        )
        + "\n"
    )


def coverage_markdown(rows: list[dict[str, Any]]) -> str:
    lines = [
        "# Regulatory Coverage and Quality",
        "",
        "This report is generated from evidence, rules, claims, source health, and resolver fixtures. File presence alone never implies production readiness.",
        "",
        "| State | Profile | Rules A/D/M/E | Claims/conflicts | Fixtures | Stale/unavailable | Readiness | Follow-up signals |",
        "| --- | --- | --- | --- | ---: | --- | --- | --- |",
    ]
    for row in rows:
        lines.append(
            f"| {row['state_name'] or row['state_id'] or 'Unknown'} | "
            f"{row['profile_status'] if row['profile_valid'] else 'missing/invalid'} | "
            f"{row['applicability_rule_count']}/{row['date_rule_count']}/{row['amendment_rule_count']}/{row['enforcement_rule_count']} | "
            f"{row['claim_count']}/{row['conflicting_claim_count']} | {row['resolver_fixture_count']} | "
            f"{row['stale_source_count']}/{row['unavailable_source_count']} | {row['readiness']} | "
            f"{', '.join(row['unresolved_signals']) or 'none'} |"
        )
    lines.extend(
        [
            "",
            "## Readiness criteria",
            "",
            "- `pilot_ready` requires a valid profile, a valid rule pack, all four rule families, and resolver fixtures.",
            "- `production_ready` additionally requires fully verified profile data, current and available sources, no conflicting claims, no unresolved policies, and no open questions.",
            "- `resolver_follow_up_required` may be legally necessary for local records and is not the same as a failed lookup.",
            "",
        ]
    )
    return "\n".join(lines)


def write_or_check(path: Path, content: str, check: bool) -> None:
    if check:
        if not path.exists() or path.read_text(encoding="utf-8") != content:
            raise ProfileError(f"generated artifact is stale: {path}")
        return
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(content, encoding="utf-8")


def _slug(value: str) -> str:
    return re.sub(r"[^a-z0-9]+", "-", value.lower()).strip("-")


def parser() -> argparse.ArgumentParser:
    root = argparse.ArgumentParser(
        description="Compile and audit Building Code Map regulatory profiles."
    )
    commands = root.add_subparsers(dest="command", required=True)
    validate = commands.add_parser("validate")
    validate.add_argument("reports", nargs="+", type=Path)
    compile_cmd = commands.add_parser("compile")
    compile_cmd.add_argument("reports", nargs="+", type=Path)
    compile_cmd.add_argument("--output-dir", type=Path, required=True)
    compile_cmd.add_argument("--check", action="store_true")
    coverage = commands.add_parser("coverage")
    coverage.add_argument("--reports-root", type=Path, default=Path("reports"))
    coverage.add_argument("--json-output", type=Path, required=True)
    coverage.add_argument("--markdown-output", type=Path, required=True)
    coverage.add_argument("--check", action="store_true")
    coverage.add_argument("--as-of", type=date.fromisoformat)
    coverage.add_argument("--stale-after-days", type=int, default=365)
    return root


def main(argv: list[str] | None = None) -> int:
    args = parser().parse_args(argv)
    try:
        if args.command == "validate":
            for report in args.reports:
                text = report.read_text(encoding="utf-8")
                profile = extract_profile(text, str(report))
                pack = extract_rule_pack(text, profile, str(report), required=False)
                if pack is None:
                    _load_external_pack(profile, report)
        elif args.command == "compile":
            compile_reports(args.reports, args.output_dir, check=args.check)
        elif args.command == "coverage":
            rows = build_coverage(
                discover_reports(args.reports_root),
                args.reports_root,
                as_of=args.as_of,
                stale_after_days=args.stale_after_days,
            )
            write_or_check(args.json_output, coverage_json(rows), args.check)
            write_or_check(args.markdown_output, coverage_markdown(rows), args.check)
        return 0
    except (OSError, ProfileError) as exc:
        print(exc, file=sys.stderr)
        return 1


if __name__ == "__main__":
    raise SystemExit(main())
