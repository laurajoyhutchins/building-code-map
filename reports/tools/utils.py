from __future__ import annotations

import re
from pathlib import Path

import yaml


def parse_heading(heading: str) -> tuple[int | None, str]:
    stripped = heading.strip()
    if not stripped:
        raise ValueError("Heading cannot be empty.")

    level = 0
    while level < len(stripped) and stripped[level] == "#":
        level += 1
    if level > 0:
        title = stripped[level:].strip()
        if not title:
            raise ValueError("Heading must include a title.")
        return level, title

    return None, stripped


def is_heading_line(line: str, level: int | None, title: str) -> bool:
    stripped = line.strip()
    if not stripped.startswith("#"):
        return False

    current_level = 0
    while current_level < len(stripped) and stripped[current_level] == "#":
        current_level += 1
    if current_level == 0:
        return False

    content = stripped[current_level:].strip()
    return content == title and (level is None or current_level == level)


def parse_value(val: str, header: str = "") -> str | int | float | list[str] | None:
    stripped = val.strip().strip("`")

    if stripped.lower() in ("null", "none", "n/a"):
        return None

    is_plural_header = header.endswith("s")

    if "`" in val and (";" in val or "," in val):
        items = re.findall(r"`([^`]+)`", val)
        if items:
            return items if len(items) > 1 else items[0]
        return [i.strip().strip("`") for i in re.split(r"[;,]", val) if i.strip()]

    if is_plural_header:
        if ";" in val or "," in val:
            items = [i.strip().strip("`") for i in re.split(r"[;,]", val) if i.strip()]
            return items if len(items) > 1 else items[0] if items else stripped
        return stripped

    try:
        if "." in stripped:
            return float(stripped)
        return int(stripped)
    except ValueError:
        return stripped


def find_table_under_heading(lines: list[str], heading: str) -> list[str]:
    level, title = parse_heading(heading)

    start_index = None
    for index, line in enumerate(lines):
        if is_heading_line(line, level, title):
            start_index = index
            break

    if start_index is None:
        raise ValueError(f"No '{heading}' heading found.")

    i = start_index + 1
    while i < len(lines) and lines[i].strip() == "":
        i += 1

    table_lines = []
    while i < len(lines):
        line = lines[i].strip()
        if line == "" or line.startswith("#") or line == "---":
            break
        if "|" in line:
            table_lines.append(line)
            i += 1
            continue
        break

    return table_lines


class _CleanDumper(yaml.Dumper):
    """YAML dumper that produces clean, consistently indented output."""

    def _prepare_tag(self, tag):
        return tag


def _represent_none(dumper, _):
    return dumper.represent_scalar("tag:yaml.org,2002:null", "null")


def _represent_str(dumper, data):
    if data == "":
        return dumper.represent_scalar("tag:yaml.org,2002:str", data, style="'")
    return dumper.represent_scalar("tag:yaml.org,2002:str", data)


_CleanDumper.add_representer(type(None), _represent_none)
_CleanDumper.add_representer(str, _represent_str)


def table_to_yaml(table_lines: list[str], root_key: str) -> str:
    if len(table_lines) < 3:
        raise ValueError(f"Invalid Markdown table under '{root_key}'.")

    headers = [col.strip() for col in table_lines[0].split("|")[1:-1]]
    data_rows = table_lines[2:]
    yaml_data = []

    for row in data_rows:
        columns = [col.strip() for col in row.split("|")[1:-1]]
        if len(columns) != len(headers):
            continue

        row_dict = {}
        for header, val in zip(headers, columns):
            clean_key = header.lower().replace(" ", "_").replace("/", "_")
            row_dict[clean_key] = parse_value(val, clean_key)
        yaml_data.append(row_dict)

    return yaml.dump(
        {root_key: yaml_data},
        Dumper=_CleanDumper,
        sort_keys=False,
        default_flow_style=False,
    )


def resolve_report_paths(reports_arg: list[str] | None, raw_dir: Path) -> list[Path]:
    if reports_arg:
        return [Path(report) for report in reports_arg]
    return sorted(raw_dir.glob("*.md"))
