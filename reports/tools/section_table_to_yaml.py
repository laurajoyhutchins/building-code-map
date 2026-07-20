from __future__ import annotations

import argparse
from pathlib import Path

from utils import find_table_under_heading, table_to_yaml


def table_section_to_yaml(markdown_text: str, heading: str, root_key: str) -> str:
    lines = markdown_text.split("\n")
    table_lines = find_table_under_heading(lines, heading)
    return table_to_yaml(table_lines, root_key)


def main() -> int:
    parser = argparse.ArgumentParser(
        description="Extract a markdown table from a section and convert it to YAML."
    )
    parser.add_argument("input_file", help="Path to the markdown file.")
    parser.add_argument(
        "--heading",
        required=True,
        help="Section heading to extract, for example '### 2.1 Primary Building Code Authorities'.",
    )
    parser.add_argument(
        "--root-key",
        required=True,
        help="Top-level YAML key to write.",
    )
    parser.add_argument(
        "-o",
        "--output",
        help="Path to the YAML output file. If omitted, prints to stdout.",
    )
    args = parser.parse_args()

    markdown_input = Path(args.input_file).read_text(encoding="utf-8")
    yaml_output = table_section_to_yaml(markdown_input, args.heading, args.root_key)

    if args.output:
        Path(args.output).write_text(f"---\n{yaml_output}", encoding="utf-8")
    else:
        print("---")
        print(yaml_output, end="")

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
