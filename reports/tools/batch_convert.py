from __future__ import annotations

import argparse
import os
import sys
from pathlib import Path

sys.path.insert(0, os.path.dirname(__file__))

from replace_section_with_yaml import process_report, resolve_yaml_path
from utils import resolve_report_paths


def main() -> int:
    parser = argparse.ArgumentParser(
        description="Batch replace a markdown section with matching YAML blocks."
    )
    parser.add_argument(
        "--heading",
        default="### Key Findings",
        help="Section heading to replace. Defaults to '### Key Findings'.",
    )
    parser.add_argument(
        "--yaml-dir",
        default=str(Path(__file__).resolve().parent.parent / "yaml" / "key-findings"),
        help="Directory containing YAML files.",
    )
    parser.add_argument(
        "--yaml-template",
        default="{stem}_key_findings.yaml",
        help="Filename template for YAML files. Supports {stem} and {name}.",
    )
    parser.add_argument(
        "reports",
        nargs="*",
        help="Optional report files to process. Defaults to all reports/raw/*.md files.",
    )
    args = parser.parse_args()

    reports_dir = Path(__file__).resolve().parent.parent
    raw_dir = reports_dir / "raw"

    report_paths = resolve_report_paths(args.reports, raw_dir)

    yaml_dir = Path(args.yaml_dir)

    changed = 0
    skipped = 0
    for report_path in report_paths:
        try:
            yaml_path = resolve_yaml_path(
                report_path,
                yaml_file=None,
                yaml_dir=yaml_dir,
                yaml_template=args.yaml_template,
            )
            if process_report(report_path, args.heading, yaml_path):
                changed += 1
            else:
                skipped += 1
        except Exception as exc:
            print(f"ERR  {report_path.name}: {exc}")
            skipped += 1

    print(f"Done: {changed} updated, {skipped} unchanged or skipped")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
