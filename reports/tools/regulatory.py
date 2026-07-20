from __future__ import annotations

import argparse
import json
import re
import sys
from pathlib import Path
from typing import Any, Iterable

FENCE = re.compile(r"```json\s+jurisdiction-profile\s*\n(?P<body>.*?)\n```", re.DOTALL)
FRONTMATTER = re.compile(r"\A---\s*\n(?P<body>.*?)\n---\s*\n", re.DOTALL)
DATE = re.compile(r"^\d{4}-\d{2}-\d{2}$")
STATUSES = {
    "resolved",
    "partially_resolved",
    "local_record_required",
    "ambiguous",
    "conflicting",
    "insufficient_evidence",
}


class ProfileError(ValueError):
    pass


def extract_profile(text: str, source: str = "<memory>") -> dict[str, Any]:
    matches = list(FENCE.finditer(text))
    if not matches:
        raise ProfileError(f"{source}: missing ```json jurisdiction-profile fenced block")
    if len(matches) != 1:
        raise ProfileError(f"{source}: expected one jurisdiction-profile block, found {len(matches)}")
    try:
        profile = json.loads(matches[0].group("body"))
    except json.JSONDecodeError as exc:
        raise ProfileError(f"{source}: invalid profile JSON at line {exc.lineno}: {exc.msg}") from exc
    if not isinstance(profile, dict):
        raise ProfileError(f"{source}: profile root must be an object")
    validate_profile(profile, source)
    return profile


def validate_profile(profile: dict[str, Any], source: str = "<memory>") -> None:
    problems: list[str] = []
    for key in (
        "schema_version",
        "profile_id",
        "state_id",
        "state_name",
        "state_abbreviation",
        "state_fips",
        "status",
        "last_verified",
    ):
        if not isinstance(profile.get(key), str) or not profile[key].strip():
            problems.append(f"{key} must be a non-empty string")
    if profile.get("schema_version") != "1.0":
        problems.append("schema_version must be '1.0'")
    if not re.fullmatch(r"\d{2}", str(profile.get("state_fips", ""))):
        problems.append("state_fips must be exactly two digits")
    if not DATE.fullmatch(str(profile.get("last_verified", ""))):
        problems.append("last_verified must use YYYY-MM-DD")

    sources = _list(profile, "sources", problems, required=True)
    authorities = _list(profile, "authorities", problems, required=True)
    adoptions = _list(profile, "adoptions", problems)
    source_ids = _ids(sources, "source", problems)
    authority_ids = _ids(authorities, "authority", problems)
    adoption_ids = _ids(adoptions, "adoption", problems)

    for item in authorities:
        _source_refs(item, f"authority {item.get('id')}", source_ids, problems)
    for item in adoptions:
        _source_refs(item, f"adoption {item.get('id')}", source_ids, problems)
        if item.get("adopting_authority_id") not in authority_ids:
            problems.append(
                f"adoption {item.get('id')} references unknown authority {item.get('adopting_authority_id')}"
            )
    for item in profile.get("relationships", []):
        _source_refs(item, f"relationship {item.get('id')}", source_ids, problems)
        if item.get("from_id") not in authority_ids:
            problems.append(
                f"relationship {item.get('id')} references unknown from_id {item.get('from_id')}"
            )

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
            continue
        policies.extend((f"{group}.{key}", value) for key, value in values.items())

    for name, policy in policies:
        if not isinstance(policy, dict):
            problems.append(f"{name} must be an object")
            continue
        if policy.get("status") not in STATUSES:
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
                problems.append(f"{name} references unknown authority {candidate.get('authority_id')}")

    if problems:
        raise ProfileError(f"{source}: " + "; ".join(sorted(set(problems))))


def canonical_json(profile: dict[str, Any]) -> str:
    return json.dumps(profile, indent=2, ensure_ascii=False) + "\n"


def compile_reports(reports: Iterable[Path], output_dir: Path) -> list[Path]:
    output_dir.mkdir(parents=True, exist_ok=True)
    written: list[Path] = []
    seen: set[str] = set()
    for report in sorted(reports):
        profile = extract_profile(report.read_text(encoding="utf-8"), str(report))
        name = _slug(profile["state_name"]) + ".json"
        if name in seen:
            raise ProfileError(f"duplicate output profile {name}")
        seen.add(name)
        destination = output_dir / name
        destination.write_text(canonical_json(profile), encoding="utf-8")
        written.append(destination)
    return written


def discover_reports(root: Path) -> list[Path]:
    return [
        path
        for path in sorted(root.rglob("*.md"))
        if path.name not in {"_template.md", "README.md"} and "generated" not in path.parts
    ]


def build_coverage(reports: Iterable[Path], root: Path) -> list[dict[str, Any]]:
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
            "source_count": 0,
            "authority_count": 0,
            "adoption_count": 0,
            "code_families": [],
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
        policies = list(_policies(profile))
        signals: list[str] = []
        if profile.get("status") != "verified":
            signals.append("profile_not_verified")
        if profile.get("verification", {}).get("status") != "verified":
            signals.append("verification_incomplete")
        if any(p.get("status") in STATUSES - {"resolved", "partially_resolved"} for p in policies):
            signals.append("resolver_follow_up_required")
        if not profile.get("adoptions"):
            signals.append("no_state_adoptions_recorded")
        if row["open_questions"]:
            signals.append("open_questions")
        row.update(
            {
                "state_id": profile["state_id"],
                "state_name": profile["state_name"],
                "abbreviation": profile["state_abbreviation"],
                "profile_present": True,
                "profile_valid": True,
                "profile_status": profile["status"],
                "source_count": len(profile["sources"]),
                "authority_count": len(profile["authorities"]),
                "adoption_count": len(profile.get("adoptions", [])),
                "code_families": sorted(
                    {adoption["code_family"] for adoption in profile.get("adoptions", [])}
                ),
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
                    "profiles_present": sum(bool(row["profile_present"]) for row in rows),
                    "profiles_valid": sum(bool(row["profile_valid"]) for row in rows),
                    "profiles_verified": sum(row["profile_status"] == "verified" for row in rows),
                    "reports_requiring_follow_up": sum(
                        bool(row["unresolved_signals"]) for row in rows
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
        "This report is generated from field-level profile data. A report file existing does not imply production readiness.",
        "",
        "| State | Report | Profile | Sources | Authorities | Adoptions | Code families | Follow-up signals |",
        "| --- | --- | --- | ---: | ---: | ---: | --- | --- |",
    ]
    for row in rows:
        lines.append(
            f"| {row['state_name'] or row['state_id'] or 'Unknown'} | `{row['report']}` | "
            f"{row['profile_status'] if row['profile_valid'] else 'missing/invalid'} | "
            f"{row['source_count']} | {row['authority_count']} | {row['adoption_count']} | "
            f"{', '.join(row['code_families']) or 'none'} | "
            f"{', '.join(row['unresolved_signals']) or 'none'} |"
        )
    lines.extend(
        [
            "",
            "## Interpretation",
            "",
            "- `profile_missing_or_invalid` means the report cannot feed the resolver.",
            "- `resolver_follow_up_required` is expected where local records are legally necessary; it is not a failed lookup.",
            "- Production readiness requires validated sources, explicit uncertainty, and resolver fixtures.",
            "",
        ]
    )
    return "\n".join(lines)


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


def write_or_check(path: Path, content: str, check: bool) -> None:
    if check:
        if not path.exists() or path.read_text(encoding="utf-8") != content:
            raise ProfileError(f"generated artifact is stale: {path}")
        return
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(content, encoding="utf-8")


def _list(
    profile: dict[str, Any], key: str, problems: list[str], required: bool = False
) -> list[Any]:
    value = profile.get(key, [])
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
    item: dict[str, Any], owner: str, known: set[str], problems: list[str]
) -> None:
    refs = item.get("source_ids")
    if not isinstance(refs, list) or not refs:
        problems.append(f"{owner} must cite at least one source")
        return
    for ref in refs:
        if ref not in known:
            problems.append(f"{owner} references unknown source {ref}")


def _policies(profile: dict[str, Any]) -> Iterable[dict[str, Any]]:
    for key in ("incorporated", "unincorporated"):
        value = profile.get("defaults", {}).get(key)
        if isinstance(value, dict):
            yield value
    for group in ("code_family_overrides", "project_type_overrides"):
        for value in profile.get(group, {}).values():
            if isinstance(value, dict):
                yield value


def _slug(value: str) -> str:
    return re.sub(r"[^a-z0-9]+", "-", value.lower()).strip("-")


def _int(value: str | None) -> int | None:
    try:
        return int(value) if value not in (None, "") else None
    except ValueError:
        return None


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
    return root


def main(argv: list[str] | None = None) -> int:
    args = parser().parse_args(argv)
    try:
        if args.command == "validate":
            for report in args.reports:
                extract_profile(report.read_text(encoding="utf-8"), str(report))
        elif args.command == "compile":
            if args.check:
                for report in sorted(args.reports):
                    profile = extract_profile(report.read_text(encoding="utf-8"), str(report))
                    write_or_check(
                        args.output_dir / (_slug(profile["state_name"]) + ".json"),
                        canonical_json(profile),
                        True,
                    )
            else:
                compile_reports(args.reports, args.output_dir)
        elif args.command == "coverage":
            rows = build_coverage(discover_reports(args.reports_root), args.reports_root)
            write_or_check(args.json_output, coverage_json(rows), args.check)
            write_or_check(args.markdown_output, coverage_markdown(rows), args.check)
        return 0
    except (OSError, ProfileError) as exc:
        print(exc, file=sys.stderr)
        return 1


if __name__ == "__main__":
    raise SystemExit(main())
