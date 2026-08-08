from __future__ import annotations

import argparse
from pathlib import Path

from utils import find_table_under_heading, table_to_yaml


def table_section_to_yaml(
    markdown_text: str,
    heading: str,
    root_key: str,
    *,
    source_name: str = "<memory>",
) -> str:
    lines = markdown_text.split("\n")
    try:
        table_lines = find_table_under_heading(lines, heading)
    except ValueError as exc:
        raise type(exc)(f"{source_name}, heading '{heading}': {exc}") from exc
    return table_to_yaml(
        table_lines,
        root_key,
        source_name=source_name,
        heading=heading,
    )


def main() -> int:
    parser = argparse.ArgumentParser(
        description="Extract a markdown table from a section and convert it to YAML."
    )
    parser.add_argument("input_file", help="Path to the markdown file.")
    parser.add_argument("--heading", required=True, help="Section heading to extract.")
    parser.add_argument("--root-key", required=True, help="Registered output schema/root key.")
    parser.add_argument("-o", "--output", help="YAML output path. Defaults to stdout.")
    args = parser.parse_args()

    input_path = Path(args.input_file)
    markdown_input = input_path.read_text(encoding="utf-8")
    yaml_output = table_section_to_yaml(
        markdown_input,
        args.heading,
        args.root_key,
        source_name=str(input_path),
    )

    if args.output:
        Path(args.output).write_text(f"---\n{yaml_output}", encoding="utf-8")
    else:
        print("---")
        print(yaml_output, end="")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
