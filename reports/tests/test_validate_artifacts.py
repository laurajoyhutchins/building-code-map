from __future__ import annotations

from pathlib import Path

import pytest

from validate_artifacts import ArtifactValidationError, main, validate_artifact_file


CANONICAL_AUTHORITY = """---
authority_structure:
- authority_id: ahj:usa-nj:dca
  authority_name: New Jersey DCA
  authority_type: state_department
  legal_basis:
  - N.J.S.A. 52:27D-119 et seq.
  role: adopts
  enforcement_model: hybrid
  source_ids:
  - src:001
  verification_status:
  - partially_verified
  parent_agency: null
"""


def test_validator_accepts_schema_conforming_authority_artifact(tmp_path: Path) -> None:
    artifact = tmp_path / "authority-structure" / "new-jersey.yaml"
    artifact.parent.mkdir()
    artifact.write_text(CANONICAL_AUTHORITY, encoding="utf-8")

    validate_artifact_file(artifact)


def test_validator_accepts_schema_conforming_official_source(tmp_path: Path) -> None:
    artifact = tmp_path / "sources" / "new-jersey.yaml"
    artifact.parent.mkdir()
    artifact.write_text(
        """---
official_sources:
- source_id: src:usa-nj:dca
  title: Current Construction Codes
  source_type: agency_page
  publisher: New Jersey DCA
  url: https://example.test/codes
  accessed_date: '2026-06-26'
  snapshot_id: null
  checksum: null
  status: verified
  supports:
  - Current code matrix
  caveats: []
  notes: []
  details: []
""",
        encoding="utf-8",
    )

    validate_artifact_file(artifact)


def test_validator_rejects_legacy_shape_instead_of_normalizing_it(tmp_path: Path) -> None:
    artifact = tmp_path / "authority-structure" / "legacy.yaml"
    artifact.parent.mkdir()
    artifact.write_text(
        """---
authority_structure:
- field: Authority ID
  value: ahj:usa-nj:dca
""",
        encoding="utf-8",
    )

    with pytest.raises(ArtifactValidationError, match="missing required fields"):
        validate_artifact_file(artifact)


def test_validator_rejects_scalar_collection_field(tmp_path: Path) -> None:
    artifact = tmp_path / "sources" / "scalar-supports.yaml"
    artifact.parent.mkdir()
    artifact.write_text(
        """---
official_sources:
- source_id: src:001
  title: Source
  source_type: null
  publisher: null
  url: null
  accessed_date: null
  snapshot_id: null
  checksum: null
  status: null
  supports: one field
  caveats: []
  notes: []
  details: []
""",
        encoding="utf-8",
    )

    with pytest.raises(ArtifactValidationError, match="string_list"):
        validate_artifact_file(artifact)


def test_validator_rejects_duplicate_yaml_keys(tmp_path: Path) -> None:
    artifact = tmp_path / "key-findings" / "duplicate.yaml"
    artifact.parent.mkdir()
    artifact.write_text(
        """---
key_findings:
- topic: A
  topic: B
  finding: C
  confidence: 0.8
  source_ids: []
""",
        encoding="utf-8",
    )

    with pytest.raises(ArtifactValidationError, match="duplicate key"):
        validate_artifact_file(artifact)


def test_validator_cli_returns_nonzero_when_any_artifact_fails(tmp_path: Path, monkeypatch) -> None:
    bad = tmp_path / "authority-structure" / "bad.yaml"
    bad.parent.mkdir()
    bad.write_text("authority_structure: []\n", encoding="utf-8")
    monkeypatch.setattr("sys.argv", ["validate_artifacts.py", str(tmp_path)])

    assert main() == 1
