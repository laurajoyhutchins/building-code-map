from __future__ import annotations

import pytest
import yaml

from utils import ConversionError, table_to_yaml


HEADING = "### Official Sources"


def test_additional_official_source_metadata_is_preserved_in_details() -> None:
    table = [
        "| Source ID | Title | Effective Date |",
        "| --- | --- | --- |",
        "| src:001 | State code portal | 2026-01-15 |",
    ]

    parsed = yaml.safe_load(
        table_to_yaml(
            table,
            "official_sources",
            source_name="colorado.md",
            heading=HEADING,
        )
    )

    source = parsed["official_sources"][0]
    assert source["source_id"] == "src:001"
    assert source["title"] == "State code portal"
    assert source["details"] == [
        {"name": "effective_date", "values": ["2026-01-15"]}
    ]


def test_additional_official_source_details_are_serialized_deterministically() -> None:
    table = [
        "| Source ID | Title | Revision Label | Effective Date |",
        "| --- | --- | --- | --- |",
        "| src:001 | State code portal | 2026 edition | 2026-01-15 |",
    ]

    first = table_to_yaml(
        table,
        "official_sources",
        source_name="colorado.md",
        heading=HEADING,
    )
    second = table_to_yaml(
        table,
        "official_sources",
        source_name="colorado.md",
        heading=HEADING,
    )

    assert first == second
    assert yaml.safe_load(first)["official_sources"][0]["details"] == [
        {"name": "effective_date", "values": ["2026-01-15"]},
        {"name": "revision_label", "values": ["2026 edition"]},
    ]


def test_duplicate_normalized_additional_headers_fail_with_context() -> None:
    table = [
        "| Source ID | Title | Effective Date | effective-date |",
        "| --- | --- | --- | --- |",
        "| src:001 | State code portal | 2026-01-15 | 2026-01-16 |",
    ]

    with pytest.raises(
        ConversionError,
        match="Duplicate normalized table header 'effective_date' in colorado.md, heading '### Official Sources'",
    ):
        table_to_yaml(
            table,
            "official_sources",
            source_name="colorado.md",
            heading=HEADING,
        )
