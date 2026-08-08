from __future__ import annotations

import pytest
import yaml

from section_table_to_yaml import table_section_to_yaml
from utils import ConversionError, parse_value


HEADING = "### 2.1 Primary Building Code Authorities"


def report_with(rows: list[str]) -> str:
    return "\n".join(
        [
            "# State report",
            "",
            HEADING,
            "",
            "| Field | Value |",
            "| --- | --- |",
            *rows,
            "",
        ]
    )


def complete_rows(*, authority_id: str = "ahj:usa-08:001", source_ids: str = "src:001") -> list[str]:
    return [
        f"| **Authority ID** | {authority_id} |",
        "| **Authority name** | Example Authority |",
        "| **Authority type** | state_agency |",
        r"| **Legal basis** | Code 08.001 and section A\|B |",
        "| **Role** | adopts |",
        "| **Enforcement model** | state_enforced |",
        f"| **Source IDs** | {source_ids} |",
        "| **Verification status** | partially_verified |",
    ]


def test_explicit_string_fields_preserve_leading_zero_and_version_like_values() -> None:
    assert parse_value("08", "string") == "08"
    assert parse_value("001", "string") == "001"
    assert parse_value("5.23.1", "string") == "5.23.1"


def test_field_value_authority_emits_canonical_record_and_consistent_lists() -> None:
    yaml_text = table_section_to_yaml(
        report_with(complete_rows(source_ids="src:001")),
        HEADING,
        "authority_structure",
        source_name="new-jersey.md",
    )

    parsed = yaml.safe_load(yaml_text)
    authority = parsed["authority_structure"][0]
    assert authority["authority_id"] == "ahj:usa-08:001"
    assert authority["legal_basis"] == ["Code 08.001 and section A|B"]
    assert authority["source_ids"] == ["src:001"]
    assert authority["verification_status"] == ["partially_verified"]
    assert authority["parent_agency"] is None


def test_authority_schema_emits_many_and_zero_source_ids_as_lists() -> None:
    many = table_section_to_yaml(
        report_with(complete_rows(source_ids="`src:001`; `src:002`")),
        HEADING,
        "authority_structure",
        source_name="many.md",
    )
    zero = table_section_to_yaml(
        report_with(complete_rows(source_ids="N/A")),
        HEADING,
        "authority_structure",
        source_name="zero.md",
    )

    assert yaml.safe_load(many)["authority_structure"][0]["source_ids"] == [
        "src:001",
        "src:002",
    ]
    assert yaml.safe_load(zero)["authority_structure"][0]["source_ids"] == []


def test_multi_record_authority_table_maps_legacy_headers_to_canonical_fields() -> None:
    markdown = "\n".join(
        [
            HEADING,
            "",
            "| Authority ID | Authority Name | Authority Type | Verified Role | Legal / Regulatory Basis | Enforcement Model | Status |",
            "| --- | --- | --- | --- | --- | --- | --- |",
            "| ahj:usa-al:001 | Agency A | state agency | adopts | Code 08; Rule 001 | state | verified |",
            "| ahj:usa-al:002 | Agency B | local government | enforces | Local act | local | partially_verified |",
            "",
        ]
    )

    parsed = yaml.safe_load(
        table_section_to_yaml(
            markdown,
            HEADING,
            "authority_structure",
            source_name="alabama.md",
        )
    )

    assert parsed["authority_structure"] == [
        {
            "authority_id": "ahj:usa-al:001",
            "authority_name": "Agency A",
            "authority_type": "state agency",
            "legal_basis": ["Code 08", "Rule 001"],
            "role": "adopts",
            "enforcement_model": "state",
            "source_ids": [],
            "verification_status": ["verified"],
            "parent_agency": None,
        },
        {
            "authority_id": "ahj:usa-al:002",
            "authority_name": "Agency B",
            "authority_type": "local government",
            "legal_basis": ["Local act"],
            "role": "enforces",
            "enforcement_model": "local",
            "source_ids": [],
            "verification_status": ["partially_verified"],
            "parent_agency": None,
        },
    ]


def test_malformed_row_reports_file_heading_and_row_number() -> None:
    rows = complete_rows()
    rows[3] = "| **Legal basis** | missing trailing boundary"

    with pytest.raises(ConversionError) as exc_info:
        table_section_to_yaml(
            report_with(rows),
            HEADING,
            "authority_structure",
            source_name="broken.md",
        )

    message = str(exc_info.value)
    assert "broken.md" in message
    assert HEADING in message
    assert "row 4" in message


def test_missing_required_schema_field_fails_loudly() -> None:
    rows = complete_rows()[:-1]

    with pytest.raises(ConversionError, match="verification_status"):
        table_section_to_yaml(
            report_with(rows),
            HEADING,
            "authority_structure",
            source_name="partial.md",
        )


def test_key_findings_schema_declares_numeric_confidence_and_list_sources() -> None:
    markdown = "\n".join(
        [
            "### Key Findings",
            "",
            "| Topic | Finding | Confidence | Source IDs |",
            "| --- | --- | --- | --- |",
            r"| State authority | Section 08\|001 controls | 0.82 | `src:001`; `src:002` |",
            "",
        ]
    )

    yaml_text = table_section_to_yaml(
        markdown,
        "### Key Findings",
        "key_findings",
        source_name="findings.md",
    )

    parsed = yaml.safe_load(yaml_text)
    assert parsed == {
        "key_findings": [
            {
                "topic": "State authority",
                "finding": "Section 08|001 controls",
                "confidence": 0.82,
                "source_ids": ["src:001", "src:002"],
            }
        ]
    }


def test_missing_heading_identifies_source_file() -> None:
    with pytest.raises(ConversionError) as exc_info:
        table_section_to_yaml(
            "# No requested section\n",
            HEADING,
            "authority_structure",
            source_name="missing.md",
        )

    assert "missing.md" in str(exc_info.value)
    assert HEADING in str(exc_info.value)
