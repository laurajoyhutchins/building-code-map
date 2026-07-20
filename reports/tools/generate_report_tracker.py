#!/usr/bin/env python3
"""Generate a status dashboard for the state report set.

The script scans `reports/*.md`, excluding `_template.md`, and summarizes:
- completion by report status
- issue counts and open issues
- source registry coverage
- consistency warnings such as leftover template markers or unresolved sources

It prints a Markdown dashboard by default. Use `--json` for JSON output, and
`--write-markdown` / `--write-json` to save artifacts to disk.
"""

from __future__ import annotations

import argparse
import json
import re
import sys
from collections import Counter
from dataclasses import asdict, dataclass
from datetime import datetime, timezone
from pathlib import Path




STATE_ORDER: list[tuple[str, str, str]] = [
    ("alabama", "AL", "Alabama"),
    ("alaska", "AK", "Alaska"),
    ("arizona", "AZ", "Arizona"),
    ("arkansas", "AR", "Arkansas"),
    ("california", "CA", "California"),
    ("colorado", "CO", "Colorado"),
    ("connecticut", "CT", "Connecticut"),
    ("delaware", "DE", "Delaware"),
    ("florida", "FL", "Florida"),
    ("georgia", "GA", "Georgia"),
    ("hawaii", "HI", "Hawaii"),
    ("idaho", "ID", "Idaho"),
    ("illinois", "IL", "Illinois"),
    ("indiana", "IN", "Indiana"),
    ("iowa", "IA", "Iowa"),
    ("kansas", "KS", "Kansas"),
    ("kentucky", "KY", "Kentucky"),
    ("louisiana", "LA", "Louisiana"),
    ("maine", "ME", "Maine"),
    ("maryland", "MD", "Maryland"),
    ("massachusetts", "MA", "Massachusetts"),
    ("michigan", "MI", "Michigan"),
    ("minnesota", "MN", "Minnesota"),
    ("mississippi", "MS", "Mississippi"),
    ("missouri", "MO", "Missouri"),
    ("montana", "MT", "Montana"),
    ("nebraska", "NE", "Nebraska"),
    ("nevada", "NV", "Nevada"),
    ("new-hampshire", "NH", "New Hampshire"),
    ("new-jersey", "NJ", "New Jersey"),
    ("new-mexico", "NM", "New Mexico"),
    ("new-york", "NY", "New York"),
    ("north-carolina", "NC", "North Carolina"),
    ("north-dakota", "ND", "North Dakota"),
    ("ohio", "OH", "Ohio"),
    ("oklahoma", "OK", "Oklahoma"),
    ("oregon", "OR", "Oregon"),
    ("pennsylvania", "PA", "Pennsylvania"),
    ("rhode-island", "RI", "Rhode Island"),
    ("south-carolina", "SC", "South Carolina"),
    ("south-dakota", "SD", "South Dakota"),
    ("tennessee", "TN", "Tennessee"),
    ("texas", "TX", "Texas"),
    ("utah", "UT", "Utah"),
    ("vermont", "VT", "Vermont"),
    ("virginia", "VA", "Virginia"),
    ("washington", "WA", "Washington"),
    ("west-virginia", "WV", "West Virginia"),
    ("wisconsin", "WI", "Wisconsin"),
    ("wyoming", "WY", "Wyoming"),
]

EXPECTED_STATE_BY_SLUG = {slug: (abbr, name) for slug, abbr, name in STATE_ORDER}
EXPECTED_SLUG_BY_ABBR = {abbr: slug for slug, abbr, _ in STATE_ORDER}

PLACEHOLDER_PATTERN = re.compile(
    r"\{\{|\bTBD\b|\bTODO\b|<null>|unknown pending|not verified in this pass",
    re.IGNORECASE,
)
SOURCE_ID_PATTERN = re.compile(r"\bsrc:[a-z0-9][a-z0-9:-]*", re.IGNORECASE)


@dataclass(slots=True)
class ReportRow:
    slug: str
    state_name: str
    abbreviation: str
    file_name: str
    present: bool
    report_status: str
    last_updated: str | None
    last_verified: str | None
    open_questions_count: int | None
    issue_count: int
    open_issue_count: int
    source_count: int
    validation_fail_count: int
    placeholder_count: int
    missing_sections: list[str]
    unresolved_sources: list[str]
    warnings: list[str]


def parse_args(argv: list[str]) -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Generate a tracker for building-code-map state reports."
    )
    parser.add_argument(
        "--root",
        type=Path,
        default=Path(__file__).resolve().parents[2],
        help="Repository root. Defaults to two levels above this script.",
    )
    parser.add_argument(
        "--json",
        action="store_true",
        help="Print JSON instead of Markdown to stdout.",
    )
    parser.add_argument(
        "--write-markdown",
        type=Path,
        help="Write the Markdown dashboard to this path.",
    )
    parser.add_argument(
        "--write-json",
        type=Path,
        help="Write the JSON payload to this path.",
    )
    parser.add_argument(
        "--strict",
        action="store_true",
        help="Exit non-zero when the tracker finds warnings or missing reports.",
    )
    return parser.parse_args(argv)


def parse_scalar(raw_value: str) -> str | int | float | bool | None:
    value = raw_value.strip()
    if not value or value in {"null", "~"}:
        return None
    if "#" in value:
        value = value.split("#", 1)[0].rstrip()
    if (
        len(value) >= 2
        and ((value[0] == '"' and value[-1] == '"') or (value[0] == "'" and value[-1] == "'"))
    ):
        return value[1:-1]
    lowered = value.lower()
    if lowered == "true":
        return True
    if lowered == "false":
        return False
    if re.fullmatch(r"-?\d+", value):
        return int(value)
    if re.fullmatch(r"-?\d+\.\d+", value):
        return float(value)
    return value


def parse_front_matter(text: str) -> tuple[dict[str, dict[str, object]], list[str], list[str]]:
    lines = text.splitlines()
    if not lines or lines[0].strip() != "---":
        return {}, [], lines

    end_index = None
    for index in range(1, len(lines)):
        if lines[index].strip() == "---":
            end_index = index
            break
    if end_index is None:
        return {}, [], lines

    front_lines = lines[1:end_index]
    body_lines = lines[end_index + 1 :]
    parsed: dict[str, dict[str, object]] = {}
    current_block: str | None = None

    for line in front_lines:
        if not line.strip():
            continue
        block_match = re.match(r"^([A-Za-z_]+):\s*$", line)
        if block_match and not line.startswith(" "):
            current_block = block_match.group(1)
            parsed.setdefault(current_block, {})
            continue

        if current_block is None:
            continue

        field_match = re.match(r"^\s{2}([A-Za-z_]+):\s*(.*)$", line)
        if field_match:
            key = field_match.group(1)
            raw_value = field_match.group(2)
            parsed[current_block][key] = parse_scalar(raw_value)

    return parsed, front_lines, body_lines


def extract_section(
    lines: list[str], start_pattern: str, stop_patterns: list[str]
) -> list[str]:
    start_index = None
    for index, line in enumerate(lines):
        if re.match(start_pattern, line):
            start_index = index
            break
    if start_index is None:
        return []

    stop_index = len(lines)
    for index in range(start_index + 1, len(lines)):
        if any(re.match(pattern, lines[index]) for pattern in stop_patterns):
            stop_index = index
            break
    return lines[start_index:stop_index]


def table_rows(section_lines: list[str]) -> list[list[str]]:
    rows: list[list[str]] = []
    for line in section_lines:
        stripped = line.strip()
        if not stripped.startswith("|"):
            continue
        if set(stripped) <= {"|", "-", " "}:
            continue
        cells = [cell.strip() for cell in stripped.strip("|").split("|")]
        rows.append(cells)
    return rows


def normalize_cell(value: str) -> str:
    return value.replace("\n", " ").strip()


def count_placeholders(text: str) -> int:
    return len(PLACEHOLDER_PATTERN.findall(text))


def collect_source_ids(text: str) -> set[str]:
    return set(SOURCE_ID_PATTERN.findall(text))


def is_open_issue_status(status: str) -> bool:
    return status.lower() in {"open", "in_progress", "pending", "needs_review"}


def missing_sections(body_text: str) -> list[str]:
    expected = [str(index) for index in range(1, 12)]
    missing: list[str] = []
    for section in expected:
        if not re.search(rf"^##\s+{re.escape(section)}\.", body_text, re.MULTILINE):
            missing.append(section)
    return missing


def summarize_report(path: Path, expected_name: str, expected_abbr: str) -> ReportRow:
    text = path.read_text(encoding="utf-8")
    front_matter, _, body_lines = parse_front_matter(text)
    body_text = "\n".join(body_lines)

    state_block = front_matter.get("state", {})
    report_block = front_matter.get("report", {})
    risk_block = front_matter.get("risk", {})

    state_name = str(state_block.get("name") or expected_name)
    abbreviation = str(state_block.get("abbreviation") or expected_abbr)
    report_status = str(report_block.get("status") or "unknown")
    last_updated = report_block.get("last_updated")
    last_verified = report_block.get("last_verified")
    open_questions_count = risk_block.get("open_questions_count")
    open_questions_count = int(open_questions_count) if isinstance(open_questions_count, int) else None

    report_id = str(report_block.get("report_id") or "")
    state_id = str(state_block.get("state_id") or "")

    warnings: list[str] = []
    if abbreviation != expected_abbr:
        warnings.append(f"abbreviation mismatch: expected {expected_abbr}, found {abbreviation}")
    expected_report_id = f"state-report:usa-{expected_abbr.lower()}"
    if report_id and report_id != expected_report_id:
        warnings.append(f"report_id mismatch: expected {expected_report_id}, found {report_id}")
    expected_state_id = f"US-{expected_abbr}"
    if state_id and state_id != expected_state_id:
        warnings.append(f"state_id mismatch: expected {expected_state_id}, found {state_id}")
    if report_status not in {"draft", "partially_verified", "verified", "deprecated"}:
        warnings.append(f"unexpected report status: {report_status}")

    missing = missing_sections(body_text)
    if missing:
        warnings.append(f"missing sections: {', '.join(missing)}")

    placeholder_count = count_placeholders(text)
    if placeholder_count:
        warnings.append(f"placeholder markers present: {placeholder_count}")

    section_8 = extract_section(
        body_lines,
        r"^##\s+8\.\s+Source Registry\s*$",
        [r"^##\s+9\."],
    )
    section_8_1 = extract_section(
        body_lines,
        r"^###\s+8\.1\s+Official Sources\s*$",
        [r"^###\s+8\.2\s+", r"^##\s+9\."],
    )
    section_9_2 = extract_section(
        body_lines,
        r"^###\s+9\.2\s+Validation Checks\s*$",
        [r"^##\s+10\."],
    )
    section_10 = extract_section(
        body_lines,
        r"^##\s+10\.\s+Open Issues and Review Queue\s*$",
        [r"^##\s+11\."],
    )

    defined_source_ids = collect_source_ids("\n".join(section_8))
    referenced_source_ids = collect_source_ids(body_text)
    unresolved_sources = sorted(referenced_source_ids - defined_source_ids)
    if unresolved_sources:
        warnings.append(
            f"unresolved source ids: {', '.join(unresolved_sources[:6])}"
            + (" ..." if len(unresolved_sources) > 6 else "")
        )

    issue_rows = []
    for row in table_rows(section_10):
        if not row or not row[0].startswith("issue:"):
            continue
        issue_rows.append(row)

    issue_count = len(issue_rows)
    open_issue_count = 0
    for row in issue_rows:
        status = row[-1] if row else ""
        if is_open_issue_status(status):
            open_issue_count += 1

    validation_fail_count = 0
    for row in table_rows(section_9_2):
        if not row or row[0] == "Check":
            continue
        status = row[1] if len(row) > 1 else ""
        if status.lower() == "fail":
            validation_fail_count += 1

    source_count = 0
    for row in table_rows(section_8_1):
        if row and row[0].startswith("src:"):
            source_count += 1

    return ReportRow(
        slug=path.stem,
        state_name=state_name,
        abbreviation=abbreviation,
        file_name=path.name,
        present=True,
        report_status=report_status,
        last_updated=str(last_updated) if last_updated is not None else None,
        last_verified=str(last_verified) if last_verified is not None else None,
        open_questions_count=open_questions_count,
        issue_count=issue_count,
        open_issue_count=open_issue_count,
        source_count=source_count,
        validation_fail_count=validation_fail_count,
        placeholder_count=placeholder_count,
        missing_sections=missing,
        unresolved_sources=unresolved_sources,
        warnings=warnings,
    )


def missing_report(slug: str, abbr: str, name: str) -> ReportRow:
    return ReportRow(
        slug=slug,
        state_name=name,
        abbreviation=abbr,
        file_name=f"{slug}.md",
        present=False,
        report_status="missing",
        last_updated=None,
        last_verified=None,
        open_questions_count=None,
        issue_count=0,
        open_issue_count=0,
        source_count=0,
        validation_fail_count=0,
        placeholder_count=0,
        missing_sections=[str(index) for index in range(1, 12)],
        unresolved_sources=[],
        warnings=["report file is missing"],
    )


def markdown_table(headers: list[str], rows: list[list[str]]) -> str:
    def esc(value: object) -> str:
        text = "" if value is None else str(value)
        return text.replace("|", r"\|").replace("\n", " ")

    lines = [
        "| " + " | ".join(esc(header) for header in headers) + " |",
        "| " + " | ".join("---" for _ in headers) + " |",
    ]
    for row in rows:
        lines.append("| " + " | ".join(esc(cell) for cell in row) + " |")
    return "\n".join(lines)


def render_markdown(rows: list[ReportRow], generated_at: str) -> str:
    present_rows = [row for row in rows if row.present]
    missing_rows = [row for row in rows if not row.present]

    status_counts = Counter(row.report_status for row in present_rows)
    open_issues_total = sum(row.open_issue_count for row in present_rows)
    issue_total = sum(row.issue_count for row in present_rows)
    source_total = sum(row.source_count for row in present_rows)
    placeholder_total = sum(row.placeholder_count for row in present_rows)
    validation_fail_total = sum(row.validation_fail_count for row in present_rows)
    open_questions_total = sum(row.open_questions_count or 0 for row in present_rows)
    total_warnings = sum(len(row.warnings) for row in rows)

    lines: list[str] = [
        "# State Report Tracker",
        "",
        f"_Generated: {generated_at}_",
        "",
        "## Overview",
        f"- Reports present: {len(present_rows)} / {len(rows)}",
        f"- Missing reports: {len(missing_rows)}",
        f"- Total source rows: {source_total}",
        f"- Total issue rows: {issue_total}",
        f"- Open issues: {open_issues_total}",
        f"- Open questions: {open_questions_total}",
        f"- Placeholder markers: {placeholder_total}",
        f"- Validation fails: {validation_fail_total}",
        f"- Warnings: {total_warnings}",
        "",
        "## Status Counts",
        markdown_table(
            ["Status", "Count"],
            [[status, count] for status, count in sorted(status_counts.items())],
        ),
        "",
        "## State Summary",
        markdown_table(
            [
                "State",
                "File",
                "Status",
                "Updated",
                "Questions",
                "Issues",
                "Open",
                "Sources",
                "Val Fail",
                "Placeholders",
                "Warnings",
            ],
            [
                [
                    row.state_name,
                    row.file_name,
                    row.report_status,
                    row.last_updated or "-",
                    row.open_questions_count if row.open_questions_count is not None else "-",
                    row.issue_count,
                    row.open_issue_count,
                    row.source_count,
                    row.validation_fail_count,
                    row.placeholder_count,
                    "; ".join(row.warnings) if row.warnings else "-",
                ]
                for row in rows
            ],
        ),
    ]

    if missing_rows:
        lines.extend(
            [
                "",
                "## Missing Reports",
                markdown_table(
                    ["Slug", "State", "Abbreviation", "File"],
                    [[row.slug, row.state_name, row.abbreviation, row.file_name] for row in missing_rows],
                ),
            ]
        )

    flagged_rows = [row for row in rows if row.warnings and row.present]
    if flagged_rows:
        lines.extend(["", "## Warnings"])
        for row in flagged_rows:
            lines.append(f"- **{row.state_name}**: {'; '.join(row.warnings)}")

    return "\n".join(lines).rstrip() + "\n"


def to_json_payload(rows: list[ReportRow], generated_at: str) -> dict[str, object]:
    present_rows = [row for row in rows if row.present]
    missing_rows = [row for row in rows if not row.present]
    status_counts = Counter(row.report_status for row in present_rows)

    return {
        "generated_at": generated_at,
        "summary": {
            "present": len(present_rows),
            "expected": len(rows),
            "missing": len(missing_rows),
            "status_counts": dict(sorted(status_counts.items())),
            "source_rows": sum(row.source_count for row in present_rows),
            "issue_rows": sum(row.issue_count for row in present_rows),
            "open_issues": sum(row.open_issue_count for row in present_rows),
            "open_questions": sum(row.open_questions_count or 0 for row in present_rows),
            "validation_fails": sum(row.validation_fail_count for row in present_rows),
            "placeholder_markers": sum(row.placeholder_count for row in present_rows),
        },
        "rows": [asdict(row) for row in rows],
    }


def write_text(path: Path, text: str) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(text, encoding="utf-8")


def main(argv: list[str]) -> int:
    args = parse_args(argv)
    root: Path = args.root.resolve()
    reports_dir = root / "reports"

    expected_rows: list[ReportRow] = []
    discovered = {path.stem: path for path in reports_dir.glob("*.md") if path.name != "_template.md"}

    for slug, abbr, name in STATE_ORDER:
        path = discovered.pop(slug, None)
        if path is None:
            expected_rows.append(missing_report(slug, abbr, name))
        else:
            expected_rows.append(summarize_report(path, name, abbr))

    extra_rows: list[ReportRow] = []
    for slug, path in sorted(discovered.items()):
        if slug == "_template":
            continue
        extra_rows.append(
            ReportRow(
                slug=slug,
                state_name=slug.replace("-", " ").title(),
                abbreviation="??",
                file_name=path.name,
                present=True,
                report_status="unexpected",
                last_updated=None,
                last_verified=None,
                open_questions_count=None,
                issue_count=0,
                open_issue_count=0,
                source_count=0,
                validation_fail_count=0,
                placeholder_count=0,
                missing_sections=[],
                unresolved_sources=[],
                warnings=["unexpected report file"],
            )
        )

    rows = expected_rows + extra_rows
    generated_at = datetime.now(timezone.utc).isoformat(timespec="seconds")
    json_payload = to_json_payload(rows, generated_at)
    markdown_output = render_markdown(rows, generated_at)

    if args.write_markdown:
        write_text(args.write_markdown, markdown_output)
    if args.write_json:
        write_text(args.write_json, json.dumps(json_payload, indent=2, sort_keys=True) + "\n")

    if args.json:
        sys.stdout.write(json.dumps(json_payload, indent=2, sort_keys=True) + "\n")
    else:
        sys.stdout.write(markdown_output)

    if args.strict:
        warnings_present = any(row.warnings for row in rows)
        missing_present = any(not row.present for row in rows)
        unresolved_sources_present = any(row.unresolved_sources for row in rows if row.present)
        if warnings_present or missing_present or unresolved_sources_present:
            return 1

    return 0


if __name__ == "__main__":
    raise SystemExit(main(sys.argv[1:]))
