from __future__ import annotations

import argparse
from pathlib import Path

from utils import is_heading_line, parse_heading


def normalize_heading(heading: str) -> str:
    stripped = heading.strip()
    if stripped.startswith("#"):
        return stripped
    _, title = parse_heading(heading)
    return f"### {title}"


def find_section_bounds(lines: list[str], heading: str) -> tuple[int, int, bool]:
    level, title = parse_heading(heading)
    heading_index = None
    for index, line in enumerate(lines):
        if is_heading_line(line, level, title):
            heading_index = index
            break
    if heading_index is None:
        raise ValueError(f"No '{heading}' heading found.")
    start = heading_index + 1
    while start < len(lines) and lines[start].strip() == "":
        start += 1
    end = start
    if end < len(lines) and lines[end].strip().startswith("```yaml"):
        end += 1
        while end < len(lines) and lines[end].strip() != "```":
            end += 1
        if end >= len(lines):
            raise ValueError(f"Unclosed YAML code fence under '{heading}'.")
        end += 1
    else:
        while end < len(lines):
            stripped = lines[end].strip()
            if stripped == "" or stripped == "---" or stripped.startswith("#"):
                break
            if "|" in stripped:
                end += 1
                continue
            break
    while end < len(lines) and lines[end].strip() == "":
        end += 1
    has_separator = end < len(lines) and lines[end].strip() == "---"
    if has_separator:
        end += 1
        while end < len(lines) and lines[end].strip() == "":
            end += 1
    return heading_index, end, has_separator


def replace_section(report_text: str, heading: str, yaml_text: str) -> str:
    eol = "\r\n" if "\r\n" in report_text else "\n"
    normalized_report = report_text.replace("\r\n", "\n")
    lines = normalized_report.split("\n")
    heading_index, end, has_separator = find_section_bounds(lines, heading)
    heading_line = normalize_heading(heading)
    yaml_block = yaml_text.replace("\r\n", "\n").strip()
    replacement = [heading_line, "", "```yaml", yaml_block, "```", ""]
    if has_separator:
        replacement.extend(["---", ""])
    after = end + 1
    while after < len(lines) and lines[after].strip() == "":
        after += 1
    updated_lines = lines[:heading_index] + replacement + lines[after:]
    return "\n".join(updated_lines).replace("\n", eol)


def resolve_yaml_path(report_path: Path, yaml_file: Path | None, yaml_dir: Path | None, yaml_template: str | None) -> Path:
    if yaml_file is not None:
        return yaml_file
    if yaml_dir is None or yaml_template is None:
        raise ValueError("Specify either --yaml-file or both --yaml-dir and --yaml-template.")
    return yaml_dir / yaml_template.format(stem=report_path.stem, name=report_path.name)


def process_report(report_path: Path, heading: str, yaml_path: Path) -> bool:
    if not yaml_path.exists():
        raise FileNotFoundError(
            f"Missing YAML input for {report_path}: {yaml_path}"
        )
    report_text = report_path.read_text(encoding="utf-8")
    yaml_text = yaml_path.read_text(encoding="utf-8")
    updated = replace_section(report_text, heading, yaml_text)
    if updated != report_text:
        report_path.write_text(updated, encoding="utf-8")
        print(f"OK   {report_path.name}")
        return True
    print(f"NOOP {report_path.name}")
    return False


def main() -> int:
    parser = argparse.ArgumentParser(
        description="Replace a markdown section table with a matching YAML block."
    )
    parser.add_argument(
        "--heading",
        required=True,
        help="Section heading to replace, for example '### Key Findings'.",
    )
    parser.add_argument(
        "--yaml-file",
        help="YAML file to insert. Use with one or more report files.",
    )
    parser.add_argument(
        "--yaml-dir",
        help="Directory containing YAML files when using --yaml-template.",
    )
    parser.add_argument(
        "--yaml-template",
        help="Filename template for YAML files. Supports {stem} and {name}.",
    )
    parser.add_argument(
        "reports",
        nargs="*",
        help="Optional report files to process. Defaults to all reports/raw/*.md files.",
    )
    args = parser.parse_args()

    tools_dir = Path(__file__).resolve().parent
    reports_dir = tools_dir.parent
    raw_dir = reports_dir / "raw"

    if args.reports:
        report_paths = [Path(report) for report in args.reports]
    else:
        report_paths = sorted(raw_dir.glob("*.md"))

    yaml_file = Path(args.yaml_file) if args.yaml_file else None
    yaml_dir = Path(args.yaml_dir) if args.yaml_dir else None

    changed = 0
    unchanged = 0
    failures = 0
    for report_path in report_paths:
        try:
            yaml_path = resolve_yaml_path(
                report_path, yaml_file, yaml_dir, args.yaml_template
            )
            if process_report(report_path, args.heading, yaml_path):
                changed += 1
            else:
                unchanged += 1
        except Exception as exc:
            print(f"ERR  {report_path.name}: {exc}")
            failures += 1

    print(f"Done: {changed} updated, {unchanged} unchanged, {failures} errors")
    return 1 if failures else 0


if __name__ == "__main__":
    raise SystemExit(main())
