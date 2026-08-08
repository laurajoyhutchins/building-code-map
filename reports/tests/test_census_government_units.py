from __future__ import annotations

import csv
import io
import zipfile
from pathlib import Path

import pytest

from census_government_units import EXPECTED_GENERAL_PURPOSE_HEADER, extract_general_purpose


def write_fixture_workbook(path: Path, rows: list[list[str]], *, header: tuple[str, ...] = EXPECTED_GENERAL_PURPOSE_HEADER) -> None:
    strings: list[str] = []
    lookup: dict[str, int] = {}

    def shared(value: str) -> int:
        if value not in lookup:
            lookup[value] = len(strings)
            strings.append(value)
        return lookup[value]

    all_rows = [list(header), *rows]
    sheet_rows: list[str] = []
    for row_number, row in enumerate(all_rows, start=1):
        cells: list[str] = []
        for column_number, value in enumerate(row, start=1):
            reference = f"{column_name(column_number)}{row_number}"
            index = shared(str(value))
            cells.append(f'<c r="{reference}" t="s"><v>{index}</v></c>')
        sheet_rows.append(f'<row r="{row_number}">{"".join(cells)}</row>')

    shared_strings = "".join(f"<si><t>{escape_xml(value)}</t></si>" for value in strings)
    with zipfile.ZipFile(path, "w") as archive:
        archive.writestr(
            "xl/workbook.xml",
            '<?xml version="1.0" encoding="UTF-8"?>'
            '<workbook xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main" '
            'xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships">'
            '<sheets><sheet name="General Purpose" sheetId="1" r:id="rId1"/></sheets>'
            '</workbook>',
        )
        archive.writestr(
            "xl/_rels/workbook.xml.rels",
            '<?xml version="1.0" encoding="UTF-8"?>'
            '<Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships">'
            '<Relationship Id="rId1" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/worksheet" Target="worksheets/sheet1.xml"/>'
            '</Relationships>',
        )
        archive.writestr(
            "xl/sharedStrings.xml",
            '<?xml version="1.0" encoding="UTF-8"?>'
            f'<sst xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main" count="{len(strings)}" uniqueCount="{len(strings)}">'
            f"{shared_strings}</sst>",
        )
        archive.writestr(
            "xl/worksheets/sheet1.xml",
            '<?xml version="1.0" encoding="UTF-8"?>'
            '<worksheet xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main">'
            f'<sheetData>{"".join(sheet_rows)}</sheetData>'
            '</worksheet>',
        )


def test_extract_general_purpose_preserves_source_identifiers_and_leading_zeroes(tmp_path: Path) -> None:
    workbook = tmp_path / "government-units.xlsx"
    output = tmp_path / "general-purpose.csv"
    row = [
        "100001",
        "COUNTY OF AUTAUGA",
        "1 - COUNTY",
        "ADMINISTRATOR",
        "135 N COURT ST",
        "STE B",
        "PRATTVILLE",
        "AL",
        "36067",
        "3049",
        "http://www.autaugaco.org",
        "COUNTY",
        "60342",
        "2023",
        "01",
        "001",
        "99001",
        "AUTAUGA",
        "Y",
    ]
    write_fixture_workbook(workbook, [row])

    count = extract_general_purpose(workbook, output)

    assert count == 1
    with output.open(newline="", encoding="utf-8") as handle:
        records = list(csv.DictReader(handle))
    assert records == [dict(zip(EXPECTED_GENERAL_PURPOSE_HEADER, row, strict=True))]
    assert records[0]["FIPS_STATE"] == "01"
    assert records[0]["FIPS_COUNTY"] == "001"
    assert records[0]["FIPS_PLACE"] == "99001"


def test_extract_general_purpose_rejects_source_contract_drift(tmp_path: Path) -> None:
    workbook = tmp_path / "government-units.xlsx"
    output = tmp_path / "general-purpose.csv"
    changed_header = tuple(
        "RENAMED_UNIT_TYPE" if value == "UNIT_TYPE" else value
        for value in EXPECTED_GENERAL_PURPOSE_HEADER
    )
    write_fixture_workbook(workbook, [], header=changed_header)

    with pytest.raises(ValueError, match="General Purpose header"):
        extract_general_purpose(workbook, output)

    assert not output.exists()


def column_name(number: int) -> str:
    result = ""
    while number:
        number, remainder = divmod(number - 1, 26)
        result = chr(ord("A") + remainder) + result
    return result


def escape_xml(value: str) -> str:
    return (
        value.replace("&", "&amp;")
        .replace("<", "&lt;")
        .replace(">", "&gt;")
        .replace('"', "&quot;")
    )
