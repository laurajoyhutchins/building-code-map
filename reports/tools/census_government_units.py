from __future__ import annotations

import argparse
import csv
import re
import sys
import xml.etree.ElementTree as ET
import zipfile
from pathlib import Path, PurePosixPath
from typing import Iterator

EXPECTED_GENERAL_PURPOSE_HEADER = (
    "CENSUS_ID_PID6",
    "UNIT_NAME",
    "UNIT_TYPE",
    "TITLE",
    "ADDRESS1",
    "ADDRESS2",
    "CITY",
    "STATE",
    "ZIP",
    "ZIP4",
    "WEB_ADDRESS",
    "POLITICAL_CODE_DESCRIPTION",
    "POPULATION",
    "POPULATION_SOURCE_YEAR",
    "FIPS_STATE",
    "FIPS_COUNTY",
    "FIPS_PLACE",
    "COUNTY_AREA_NAME",
    "ACTIVE",
)

_MAIN_NS = "http://schemas.openxmlformats.org/spreadsheetml/2006/main"
_DOCUMENT_REL_NS = "http://schemas.openxmlformats.org/officeDocument/2006/relationships"
_PACKAGE_REL_NS = "http://schemas.openxmlformats.org/package/2006/relationships"
_CELL_REFERENCE = re.compile(r"^([A-Z]+)")


def extract_general_purpose(workbook_path: Path, output_path: Path) -> int:
    rows = iter_sheet_rows(workbook_path, "General Purpose")
    try:
        header = tuple(next(rows))
    except StopIteration as exc:
        raise ValueError("General Purpose sheet is empty") from exc

    if header != EXPECTED_GENERAL_PURPOSE_HEADER:
        raise ValueError(
            "General Purpose header does not match the pinned 2025 Census source contract: "
            f"got {header!r}"
        )

    output_path.parent.mkdir(parents=True, exist_ok=True)
    temporary = output_path.with_name(output_path.name + ".tmp")
    count = 0
    try:
        with temporary.open("w", newline="", encoding="utf-8") as handle:
            writer = csv.writer(handle, lineterminator="\n")
            writer.writerow(EXPECTED_GENERAL_PURPOSE_HEADER)
            for raw_row in rows:
                if not any(raw_row):
                    continue
                if len(raw_row) > len(EXPECTED_GENERAL_PURPOSE_HEADER) and any(
                    raw_row[len(EXPECTED_GENERAL_PURPOSE_HEADER) :]
                ):
                    raise ValueError("General Purpose row contains unexpected trailing columns")
                row = list(raw_row[: len(EXPECTED_GENERAL_PURPOSE_HEADER)])
                row.extend([""] * (len(EXPECTED_GENERAL_PURPOSE_HEADER) - len(row)))
                writer.writerow(row)
                count += 1
        temporary.replace(output_path)
    except Exception:
        temporary.unlink(missing_ok=True)
        raise
    return count


def iter_sheet_rows(workbook_path: Path, sheet_name: str) -> Iterator[list[str]]:
    with zipfile.ZipFile(workbook_path) as archive:
        shared_strings = _read_shared_strings(archive)
        sheet_path = _resolve_sheet_path(archive, sheet_name)
        with archive.open(sheet_path) as stream:
            for event, element in ET.iterparse(stream, events=("end",)):
                if element.tag != f"{{{_MAIN_NS}}}row":
                    continue
                values: dict[int, str] = {}
                maximum = -1
                for cell in element.findall(f"{{{_MAIN_NS}}}c"):
                    reference = cell.attrib.get("r", "")
                    column = _column_index(reference)
                    values[column] = _cell_value(cell, shared_strings)
                    maximum = max(maximum, column)
                if maximum < 0:
                    yield []
                else:
                    yield [values.get(index, "") for index in range(maximum + 1)]
                element.clear()


def _resolve_sheet_path(archive: zipfile.ZipFile, sheet_name: str) -> str:
    workbook = ET.fromstring(archive.read("xl/workbook.xml"))
    relationship_id = None
    for sheet in workbook.findall(f".//{{{_MAIN_NS}}}sheet"):
        if sheet.attrib.get("name") == sheet_name:
            relationship_id = sheet.attrib.get(f"{{{_DOCUMENT_REL_NS}}}id")
            break
    if not relationship_id:
        raise ValueError(f"workbook does not contain sheet {sheet_name!r}")

    relationships = ET.fromstring(archive.read("xl/_rels/workbook.xml.rels"))
    target = None
    for relationship in relationships.findall(f"{{{_PACKAGE_REL_NS}}}Relationship"):
        if relationship.attrib.get("Id") == relationship_id:
            target = relationship.attrib.get("Target")
            break
    if not target:
        raise ValueError(f"workbook relationship {relationship_id!r} is missing")

    normalized = PurePosixPath(target.lstrip("/"))
    if normalized.parts and normalized.parts[0] == "xl":
        return str(normalized)
    return str(PurePosixPath("xl") / normalized)


def _read_shared_strings(archive: zipfile.ZipFile) -> list[str]:
    try:
        data = archive.read("xl/sharedStrings.xml")
    except KeyError:
        return []
    root = ET.fromstring(data)
    values: list[str] = []
    for item in root.findall(f"{{{_MAIN_NS}}}si"):
        values.append("".join(text.text or "" for text in item.iter(f"{{{_MAIN_NS}}}t")))
    return values


def _cell_value(cell: ET.Element, shared_strings: list[str]) -> str:
    cell_type = cell.attrib.get("t", "")
    if cell_type == "inlineStr":
        return "".join(text.text or "" for text in cell.iter(f"{{{_MAIN_NS}}}t"))
    value = cell.find(f"{{{_MAIN_NS}}}v")
    raw = "" if value is None or value.text is None else value.text
    if cell_type == "s":
        try:
            return shared_strings[int(raw)]
        except (ValueError, IndexError) as exc:
            raise ValueError(f"invalid shared-string index {raw!r}") from exc
    return raw


def _column_index(reference: str) -> int:
    match = _CELL_REFERENCE.match(reference)
    if not match:
        raise ValueError(f"invalid cell reference {reference!r}")
    result = 0
    for character in match.group(1):
        result = result * 26 + (ord(character) - ord("A") + 1)
    return result - 1


def main(argv: list[str] | None = None) -> int:
    parser = argparse.ArgumentParser(description="Extract the pinned Census Government Units General Purpose sheet")
    parser.add_argument("workbook", type=Path)
    parser.add_argument("output", type=Path)
    args = parser.parse_args(argv)
    try:
        count = extract_general_purpose(args.workbook, args.output)
    except (OSError, ValueError, zipfile.BadZipFile) as exc:
        print(exc, file=sys.stderr)
        return 1
    print(f"extracted {count} general-purpose government units")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
