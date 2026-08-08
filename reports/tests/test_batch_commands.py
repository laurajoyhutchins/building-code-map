from __future__ import annotations

import sys
from pathlib import Path

import batch_convert
import batch_table_to_yaml

from test_conversion import HEADING, complete_rows, report_with


def test_batch_table_conversion_returns_nonzero_after_partial_failure(tmp_path: Path, monkeypatch) -> None:
    valid_report = tmp_path / "valid.md"
    valid_report.write_text(report_with(complete_rows()), encoding="utf-8")
    invalid_report = tmp_path / "invalid.md"
    invalid_report.write_text("# Missing requested section\n", encoding="utf-8")
    output_dir = tmp_path / "yaml"

    monkeypatch.setattr(
        sys,
        "argv",
        [
            "batch_table_to_yaml.py",
            "--output-dir",
            str(output_dir),
            str(valid_report),
            str(invalid_report),
        ],
    )

    assert batch_table_to_yaml.main() == 1
    assert (output_dir / "valid_authority_structure.yaml").exists()


def test_batch_table_best_effort_mode_is_explicit(tmp_path: Path, monkeypatch) -> None:
    invalid_report = tmp_path / "invalid.md"
    invalid_report.write_text("# Missing requested section\n", encoding="utf-8")

    monkeypatch.setattr(
        sys,
        "argv",
        [
            "batch_table_to_yaml.py",
            "--best-effort",
            "--output-dir",
            str(tmp_path / "yaml"),
            str(invalid_report),
        ],
    )

    assert batch_table_to_yaml.main() == 0


def test_batch_section_replacement_treats_missing_yaml_as_failure(tmp_path: Path, monkeypatch) -> None:
    report = tmp_path / "missing-yaml.md"
    report.write_text("### Key Findings\n\n| A | B |\n| --- | --- |\n| x | y |\n", encoding="utf-8")
    yaml_dir = tmp_path / "yaml"
    yaml_dir.mkdir()

    monkeypatch.setattr(
        sys,
        "argv",
        [
            "batch_convert.py",
            "--yaml-dir",
            str(yaml_dir),
            str(report),
        ],
    )

    assert batch_convert.main() == 1
