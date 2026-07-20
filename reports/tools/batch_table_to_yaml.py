from __future__ import annotations

import argparse
from pathlib import Path

from section_table_to_yaml import table_section_to_yaml
from utils import resolve_report_paths


def main() -> int:
    parser = argparse.ArgumentParser(
        description="Batch convert a markdown table section to YAML files."
    )
    parser.add_argument(
        "--heading",
        default="### 2.1 Primary Building Code Authorities",
        help="Section heading to extract.",
    )
    parser.add_argument(
        "--root-key",
        default="authority_structure",
        help="Top-level YAML key to write.",
    )
    parser.add_argument(
        "--output-dir",
        default=str(Path(__file__).resolve().parent.parent / "yaml" / "authority-structure"),
        help="Directory for generated YAML files.",
    )
    parser.add_argument(
        "--yaml-template",
        default="{stem}_authority_structure.yaml",
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
    output_dir = Path(args.output_dir)
    output_dir.mkdir(parents=True, exist_ok=True)

    report_paths = resolve_report_paths(args.reports, raw_dir)

    changed = 0
    skipped = 0
    for report_path in report_paths:
        output_path = output_dir / args.yaml_template.format(stem=report_path.stem, name=report_path.name)
        try:
            markdown_input = report_path.read_text(encoding="utf-8")
            yaml_output = table_section_to_yaml(markdown_input, args.heading, args.root_key)
            output_path.write_text(f"---\n{yaml_output}", encoding="utf-8")
            print(f"OK   {report_path.name} -> {output_path.name}")
            changed += 1
        except Exception as exc:
            print(f"ERR  {report_path.name}: {exc}")
            skipped += 1

    print(f"Done: {changed} converted, {skipped} errors")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
