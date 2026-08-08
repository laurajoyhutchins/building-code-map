from __future__ import annotations

import re
from pathlib import Path

from schemas import (
    AUTHORITY_ALIASES,
    OFFICIAL_SOURCE_ALIASES,
    ValueKind,
    canonicalize_document,
    get_schema,
    normalize_field_name,
)
from yaml_io import dump_yaml


BACKSLASH = chr(92)


class ConversionError(ValueError):
    """Raised when report content cannot be converted without losing data."""


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


def _strip_markdown(value: str) -> str:
    stripped = value.strip()
    if stripped.startswith("**") and stripped.endswith("**") and len(stripped) >= 4:
        stripped = stripped[2:-2].strip()
    if stripped.startswith("`") and stripped.endswith("`") and stripped.count("`") == 2:
        stripped = stripped[1:-1].strip()
    return stripped


def _unescape(value: str) -> str:
    return (
        value.replace(BACKSLASH + "|", "|")
        .replace(BACKSLASH + ",", ",")
        .replace(BACKSLASH + ";", ";")
        .replace(BACKSLASH * 2, BACKSLASH)
    )


def _split_list(value: str, delimiters: tuple[str, ...] = (",", ";")) -> list[str]:
    stripped = value.strip()
    if not stripped or stripped.lower() in {"null", "none", "n/a"}:
        return []

    items: list[str] = []
    current: list[str] = []
    in_code = False
    escaped = False

    for char in stripped:
        if escaped:
            current.append(BACKSLASH)
            current.append(char)
            escaped = False
            continue
        if char == BACKSLASH:
            escaped = True
            continue
        if char == "`":
            in_code = not in_code
            current.append(char)
            continue
        if char in delimiters and not in_code:
            item = _unescape(_strip_markdown("".join(current)))
            if item:
                items.append(item)
            current = []
            continue
        current.append(char)

    if escaped:
        current.append(BACKSLASH)

    item = _unescape(_strip_markdown("".join(current)))
    if item:
        items.append(item)
    return items


def parse_value(
    value: str,
    kind: ValueKind = "string",
    *,
    delimiters: tuple[str, ...] = (",", ";"),
) -> str | int | float | list[str] | None:
    stripped = _unescape(_strip_markdown(value))

    if kind == "string":
        return stripped
    if kind == "string_list":
        return _split_list(value, delimiters)
    if kind == "nullable_string":
        return None if stripped.lower() in {"null", "none", "n/a"} else stripped
    if kind == "number":
        try:
            return float(stripped)
        except ValueError as exc:
            raise ConversionError(f"Expected numeric value, got {value!r}.") from exc
    raise ConversionError(f"Table input cannot directly populate schema type '{kind}'.")


def find_table_under_heading(lines: list[str], heading: str) -> list[str]:
    level, title = parse_heading(heading)

    start_index = None
    for index, line in enumerate(lines):
        if is_heading_line(line, level, title):
            start_index = index
            break

    if start_index is None:
        raise ConversionError(f"No '{heading}' heading found.")

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

    if not table_lines:
        raise ConversionError(f"No Markdown table found under '{heading}'.")
    return table_lines


def split_markdown_row(row: str) -> list[str]:
    stripped = row.strip()
    if not stripped.startswith("|") or not stripped.endswith("|"):
        raise ConversionError("Markdown table rows must begin and end with '|'.")

    body = stripped[1:-1]
    columns: list[str] = []
    current: list[str] = []
    in_code = False
    escaped = False

    for char in body:
        if escaped:
            if char in {"|", BACKSLASH}:
                current.append(char)
            else:
                current.extend((BACKSLASH, char))
            escaped = False
            continue
        if char == BACKSLASH:
            escaped = True
            continue
        if char == "`":
            in_code = not in_code
            current.append(char)
            continue
        if char == "|" and not in_code:
            columns.append("".join(current).strip())
            current = []
            continue
        current.append(char)

    if escaped:
        current.append(BACKSLASH)
    if in_code:
        raise ConversionError("Markdown table row contains an unclosed inline-code span.")
    columns.append("".join(current).strip())
    return columns


def _context(source_name: str, heading: str, row_number: int | None = None) -> str:
    context = f"{source_name}, heading '{heading}'"
    if row_number is not None:
        context += f", row {row_number}"
    return context


def _validate_separator(row: str, expected_columns: int, source_name: str, heading: str) -> None:
    columns = split_markdown_row(row)
    if len(columns) != expected_columns or any(
        re.fullmatch(r":?-{3,}:?", column) is None for column in columns
    ):
        raise ConversionError(
            f"Invalid Markdown separator row in {_context(source_name, heading)}."
        )


def _canonical_input_field(root_key: str, raw_field: str) -> tuple[str, str, ValueKind]:
    normalized = normalize_field_name(raw_field)
    schema = get_schema(root_key)
    if root_key == "authority_structure":
        canonical = AUTHORITY_ALIASES.get(normalized)
        if canonical is None:
            raise ConversionError(f"Unknown authority field {raw_field!r}.")
        return normalized, canonical, schema.fields[canonical].kind
    if root_key == "key_findings":
        if normalized not in schema.fields:
            raise ConversionError(f"Unknown key-finding field {raw_field!r}.")
        return normalized, normalized, schema.fields[normalized].kind
    if root_key == "official_sources":
        canonical = OFFICIAL_SOURCE_ALIASES.get(normalized)
        if canonical is None:
            return normalized, normalized, "string"
        return normalized, canonical, schema.fields[canonical].kind
    raise AssertionError(f"Unhandled root key: {root_key}")


def _parse_cell(
    root_key: str,
    field: str,
    value: str,
    source_name: str,
    heading: str,
    row_number: int,
) -> tuple[str, object]:
    try:
        normalized, canonical, kind = _canonical_input_field(root_key, field)
        delimiters: tuple[str, ...] = (",", ";")
        if kind == "string_list":
            if canonical in {"legal_basis", "verification_status"}:
                delimiters = (";",)
            elif canonical not in {"source_ids"}:
                delimiters = (";",)
        return normalized, parse_value(value, kind, delimiters=delimiters)
    except (ConversionError, ValueError) as exc:
        raise ConversionError(
            f"{_context(source_name, heading, row_number)}, field '{field}': {exc}"
        ) from exc


def _convert_field_value_table(
    table_lines: list[str], root_key: str, source_name: str, heading: str
) -> dict[str, object]:
    headers = tuple(_strip_markdown(header) for header in split_markdown_row(table_lines[0]))
    if headers != ("Field", "Value"):
        raise ConversionError(
            f"Unexpected field/value table headers in {_context(source_name, heading)}: "
            f"got {headers}."
        )
    _validate_separator(table_lines[1], len(headers), source_name, heading)

    record: dict[str, object] = {}
    for row_number, row in enumerate(table_lines[2:], start=1):
        try:
            columns = split_markdown_row(row)
        except ConversionError as exc:
            raise ConversionError(f"{_context(source_name, heading, row_number)}: {exc}") from exc
        if len(columns) != len(headers):
            raise ConversionError(
                f"{_context(source_name, heading, row_number)}: expected {len(headers)} columns, "
                f"got {len(columns)}."
            )
        field, parsed = _parse_cell(
            root_key, columns[0], columns[1], source_name, heading, row_number
        )
        if field in record:
            raise ConversionError(
                f"{_context(source_name, heading, row_number)}: duplicate field '{field}'."
            )
        record[field] = parsed
    return {root_key: record}


def _convert_records_table(
    table_lines: list[str], root_key: str, source_name: str, heading: str
) -> dict[str, object]:
    headers = [_strip_markdown(header) for header in split_markdown_row(table_lines[0])]
    _validate_separator(table_lines[1], len(headers), source_name, heading)

    normalized_headers: list[str] = []
    for header in headers:
        normalized, _, _ = _canonical_input_field(root_key, header)
        if normalized in normalized_headers:
            raise ConversionError(
                f"Duplicate normalized table header '{normalized}' in {_context(source_name, heading)}."
            )
        normalized_headers.append(normalized)

    records: list[dict[str, object]] = []
    for row_number, row in enumerate(table_lines[2:], start=1):
        try:
            columns = split_markdown_row(row)
        except ConversionError as exc:
            raise ConversionError(f"{_context(source_name, heading, row_number)}: {exc}") from exc
        if len(columns) != len(headers):
            raise ConversionError(
                f"{_context(source_name, heading, row_number)}: expected {len(headers)} columns, "
                f"got {len(columns)}."
            )

        record: dict[str, object] = {}
        for header, value in zip(headers, columns):
            field, parsed = _parse_cell(
                root_key, header, value, source_name, heading, row_number
            )
            record[field] = parsed
        records.append(record)
    return {root_key: records}


def table_to_yaml(
    table_lines: list[str],
    root_key: str,
    *,
    source_name: str = "<memory>",
    heading: str | None = None,
) -> str:
    display_heading = heading or root_key
    if len(table_lines) < 3:
        raise ConversionError(
            f"Invalid Markdown table in {_context(source_name, display_heading)}: "
            "expected a header, separator, and at least one data row."
        )

    get_schema(root_key)
    headers = tuple(_strip_markdown(header) for header in split_markdown_row(table_lines[0]))
    try:
        if headers == ("Field", "Value"):
            raw_document = _convert_field_value_table(
                table_lines, root_key, source_name, display_heading
            )
        else:
            raw_document = _convert_records_table(
                table_lines, root_key, source_name, display_heading
            )
        canonical = canonicalize_document(raw_document, source_name)
    except ValueError as exc:
        if isinstance(exc, ConversionError):
            raise
        raise ConversionError(f"{_context(source_name, display_heading)}: {exc}") from exc
    return dump_yaml(canonical)


def resolve_report_paths(reports_arg: list[str] | None, raw_dir: Path) -> list[Path]:
    if reports_arg:
        return [Path(report) for report in reports_arg]
    return sorted(raw_dir.glob("*.md"))
