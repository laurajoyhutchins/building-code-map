from __future__ import annotations

from pathlib import Path

import yaml

from normalize_artifacts import main, normalize_artifact_file


def test_normalizer_migrates_field_value_authority_without_losing_identifiers(tmp_path: Path) -> None:
    artifact = tmp_path / "authority.yaml"
    artifact.write_text(
        """---
authority_structure:
- field: Authority ID
  value: ahj:usa-08:001
- field: Authority name
  value: Example Authority
- field: Authority type
  value: state_agency
- field: Legal basis
  value: Code 08.001
- field: Role
  value: adopts
- field: Enforcement model
  value: statewide
- field: Source IDs
  value: src:001
- field: Verification status
  value: verified
""",
        encoding="utf-8",
    )

    normalized = yaml.safe_load(normalize_artifact_file(artifact))
    record = normalized["authority_structure"][0]
    assert record["authority_id"] == "ahj:usa-08:001"
    assert record["legal_basis"] == ["Code 08.001"]
    assert record["source_ids"] == ["src:001"]
    assert record["verification_status"] == ["verified"]


def test_normalizer_canonicalizes_official_source_aliases_and_preserves_details(tmp_path: Path) -> None:
    artifact = tmp_path / "sources.yaml"
    artifact.write_text(
        """---
official_sources:
- source_id: src:001
  title___citation: Statute 08.001
  publisher___custodian: State Agency
  type: statute
  url___locator: https://example.test/08.001
  date___version: 2026
  supports: Current code
  caveat: Courtesy copy
  status: verified
""",
        encoding="utf-8",
    )

    normalized = yaml.safe_load(normalize_artifact_file(artifact))
    record = normalized["official_sources"][0]
    assert record["title"] == "Statute 08.001"
    assert record["publisher"] == "State Agency"
    assert record["source_type"] == "statute"
    assert record["url"] == "https://example.test/08.001"
    assert record["supports"] == ["Current code"]
    assert record["caveats"] == ["Courtesy copy"]
    assert record["details"] == [{"name": "date_version", "values": ["2026"]}]


def test_normalizer_check_mode_reports_legacy_artifact(tmp_path: Path, monkeypatch) -> None:
    artifact = tmp_path / "key-findings.yaml"
    artifact.write_text(
        """---
key_findings:
- topic: Effective date
  finding:
  - 2026-01-01
  confidence: 0.9
  source_ids: []
""",
        encoding="utf-8",
    )
    monkeypatch.setattr(
        "sys.argv", ["normalize_artifacts.py", str(tmp_path), "--check"]
    )

    assert main() == 1
