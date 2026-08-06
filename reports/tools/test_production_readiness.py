from __future__ import annotations

import json
import tempfile
import unittest
from datetime import date
from pathlib import Path
from urllib.error import HTTPError

import production_readiness


VERIFIED = {"status": "verified", "confidence": 0.95}


def complete_inputs() -> tuple[dict, dict, dict]:
    profile = {
        "schema_version": "1.0",
        "profile_id": "state-profile:us-aa",
        "state_id": "US-AA",
        "state_name": "Alpha",
        "state_abbreviation": "AA",
        "state_fips": "01",
        "status": "verified",
        "last_verified": "2026-08-01",
        "sources": [
            {
                "id": "src:aa:state",
                "title": "State code",
                "url": "https://example.test/state",
                "kind": "statute",
                "accessed_at": "2026-08-01",
            },
            {
                "id": "src:aa:city",
                "title": "City ordinance",
                "url": "https://example.test/city",
                "kind": "ordinance",
                "accessed_at": "2026-08-01",
            },
            {
                "id": "src:aa:county",
                "title": "County resolution",
                "url": "https://example.test/county",
                "kind": "resolution",
                "accessed_at": "2026-08-01",
            },
        ],
        "authorities": [
            {
                "id": "auth:aa",
                "name": "Alpha Board",
                "type": "state_board",
                "roles": ["adopts"],
                "source_ids": ["src:aa:state"],
                "verification": VERIFIED,
            }
        ],
        "relationships": [],
        "adoptions": [
            {
                "id": "adoption:aa:2021",
                "code_family": "building",
                "status": "prior",
                "state_code_name": "Alpha 2021",
                "adopting_authority_id": "auth:aa",
                "enforcement_model": "local",
                "dates": {
                    "effective_date": "2021-01-01",
                    "replacement_date": "2024-01-01",
                },
                "source_ids": ["src:aa:state"],
                "verification": VERIFIED,
            },
            {
                "id": "adoption:aa:2024",
                "code_family": "building",
                "status": "current",
                "state_code_name": "Alpha 2024",
                "adopting_authority_id": "auth:aa",
                "enforcement_model": "local",
                "dates": {"effective_date": "2024-01-01"},
                "source_ids": ["src:aa:state"],
                "verification": VERIFIED,
            },
        ],
        "defaults": {
            "incorporated": {
                "status": "partially_resolved",
                "adoption_ids": ["adoption:aa:2024"],
                "source_ids": ["src:aa:state"],
            },
            "unincorporated": {
                "status": "partially_resolved",
                "adoption_ids": ["adoption:aa:2024"],
                "source_ids": ["src:aa:state"],
            },
        },
        "code_family_overrides": {},
        "project_type_overrides": {},
        "verification": VERIFIED,
    }
    fixture_ids = [
        "fixture:aa:city",
        "fixture:aa:county",
        "fixture:aa:special",
        "fixture:aa:missing",
        "fixture:aa:conflict",
        "fixture:aa:historical",
    ]
    pack = {
        "schema_version": "1.0",
        "state_id": "US-AA",
        "source_health": {
            source_id: {"last_checked_at": "2026-08-01", "availability": "available"}
            for source_id in ("src:aa:state", "src:aa:city", "src:aa:county")
        },
        "applicability_rules": [],
        "date_rules": [],
        "amendment_rules": [],
        "enforcement_rules": [],
        "claims": [
            {
                "id": "claim:aa:effective",
                "subject_id": "adoption:aa:2024",
                "field": "dates.effective_date",
                "status": "supported",
                "value": "2024-01-01",
                "source_ids": ["src:aa:state"],
                "verification": VERIFIED,
            }
        ],
        "resolver_fixture_ids": fixture_ids,
    }
    fixtures = [
        {
            "id": "fixture:aa:city",
            "kind": "incorporated",
            "jurisdiction": "Alpha City",
            "code_families": ["building"],
            "applicability_date": "2026-08-01",
            "expected_status": "partially_resolved",
            "production_supported": True,
            "source_ids": ["src:aa:state", "src:aa:city"],
            "local_source_ids": ["src:aa:city"],
            "verification": VERIFIED,
        },
        {
            "id": "fixture:aa:county",
            "kind": "unincorporated",
            "jurisdiction": "Alpha County",
            "code_families": ["building"],
            "applicability_date": "2026-08-01",
            "expected_status": "partially_resolved",
            "production_supported": True,
            "source_ids": ["src:aa:state", "src:aa:county"],
            "local_source_ids": ["src:aa:county"],
            "verification": VERIFIED,
        },
        {
            "id": "fixture:aa:special",
            "kind": "special_project",
            "jurisdiction": "State project",
            "code_families": ["building"],
            "applicability_date": "2026-08-01",
            "expected_status": "resolved",
            "production_supported": True,
            "source_ids": ["src:aa:state"],
            "local_source_ids": [],
            "verification": VERIFIED,
        },
        {
            "id": "fixture:aa:missing",
            "kind": "missing_evidence",
            "jurisdiction": "Unknown local record",
            "code_families": ["building"],
            "applicability_date": "2026-08-01",
            "expected_status": "insufficient_evidence",
            "production_supported": False,
            "source_ids": ["src:aa:state"],
            "local_source_ids": [],
            "verification": VERIFIED,
        },
        {
            "id": "fixture:aa:conflict",
            "kind": "conflict",
            "jurisdiction": "Conflicting local record",
            "code_families": ["building"],
            "applicability_date": "2026-08-01",
            "expected_status": "conflicting",
            "production_supported": False,
            "source_ids": ["src:aa:state", "src:aa:city"],
            "local_source_ids": ["src:aa:city"],
            "verification": VERIFIED,
        },
        {
            "id": "fixture:aa:historical",
            "kind": "historical",
            "jurisdiction": "Alpha City",
            "code_families": ["building"],
            "applicability_date": "2022-01-01",
            "expected_status": "partially_resolved",
            "production_supported": True,
            "source_ids": ["src:aa:state", "src:aa:city"],
            "local_source_ids": ["src:aa:city"],
            "verification": VERIFIED,
        },
    ]
    manifest = {
        "schema_version": "1.0",
        "state_id": "US-AA",
        "scope": {
            "code_families": ["building"],
            "start_date": "2021-01-01",
            "end_date": "2026-08-01",
        },
        "primary_source_ids": ["src:aa:state", "src:aa:city", "src:aa:county"],
        "required_fixture_kinds": sorted(production_readiness.REQUIRED_FIXTURE_KINDS),
        "timelines": [
            {
                "code_family": "building",
                "start_date": "2021-01-01",
                "end_date": "2023-12-31",
                "adoption_ids": ["adoption:aa:2021"],
                "source_ids": ["src:aa:state"],
                "verification": VERIFIED,
            },
            {
                "code_family": "building",
                "start_date": "2024-01-01",
                "end_date": None,
                "adoption_ids": ["adoption:aa:2024"],
                "source_ids": ["src:aa:state"],
                "verification": VERIFIED,
            },
        ],
        "fixtures": fixtures,
    }
    return profile, pack, manifest


class ProductionReadinessTests(unittest.TestCase):
    def test_complete_scoped_evidence_is_production_ready(self) -> None:
        profile, pack, manifest = complete_inputs()
        result = production_readiness.audit_state(
            profile, pack, manifest, as_of=date(2026, 8, 2)
        )
        self.assertEqual(result["readiness"], "production_ready")
        self.assertEqual(result["signals"], [])

    def test_missing_required_fixture_class_blocks_promotion(self) -> None:
        profile, pack, manifest = complete_inputs()
        manifest["fixtures"] = [
            fixture for fixture in manifest["fixtures"] if fixture["kind"] != "conflict"
        ]
        result = production_readiness.audit_state(
            profile, pack, manifest, as_of=date(2026, 8, 2)
        )
        self.assertEqual(result["readiness"], "needs_review")
        self.assertIn("fixture_kind_missing:conflict", result["signals"])

    def test_timeline_gap_blocks_promotion(self) -> None:
        profile, pack, manifest = complete_inputs()
        manifest["timelines"][1]["start_date"] = "2024-02-01"
        result = production_readiness.audit_state(
            profile, pack, manifest, as_of=date(2026, 8, 2)
        )
        self.assertIn("timeline_gap:building:2024-01-01:2024-01-31", result["signals"])
        self.assertEqual(result["readiness"], "needs_review")

    def test_production_local_fixture_requires_verified_local_primary_source(self) -> None:
        profile, pack, manifest = complete_inputs()
        city = next(item for item in manifest["fixtures"] if item["id"] == "fixture:aa:city")
        city["local_source_ids"] = []
        result = production_readiness.audit_state(
            profile, pack, manifest, as_of=date(2026, 8, 2)
        )
        self.assertIn("fixture_local_evidence_missing:fixture:aa:city", result["signals"])

    def test_registered_placeholder_fixture_without_definition_is_reported(self) -> None:
        profile, pack, manifest = complete_inputs()
        manifest["fixtures"] = manifest["fixtures"][:-1]
        result = production_readiness.audit_state(
            profile, pack, manifest, as_of=date(2026, 8, 2)
        )
        self.assertIn(
            "registered_fixture_undefined:fixture:aa:historical", result["signals"]
        )

    def test_session_law_summary_cannot_satisfy_primary_source_gate(self) -> None:
        profile, pack, manifest = complete_inputs()
        profile["sources"][0]["kind"] = "session_law_summary"
        result = production_readiness.audit_state(
            profile, pack, manifest, as_of=date(2026, 8, 2)
        )
        self.assertIn(
            "source_not_primary:src:aa:state:session_law_summary", result["signals"]
        )

    def test_stale_primary_source_blocks_promotion(self) -> None:
        profile, pack, manifest = complete_inputs()
        pack["source_health"]["src:aa:state"]["last_checked_at"] = "2024-01-01"
        result = production_readiness.audit_state(
            profile, pack, manifest, as_of=date(2026, 8, 2), stale_after_days=365
        )
        self.assertIn("primary_source_stale:src:aa:state", result["signals"])

    def test_repository_audit_flags_premature_verified_profile(self) -> None:
        profile, pack, manifest = complete_inputs()
        manifest["fixtures"] = []
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            compiled = root / "compiled"
            manifests = root / "manifests"
            (compiled / "rules").mkdir(parents=True)
            manifests.mkdir()
            (compiled / "alpha.json").write_text(json.dumps(profile), encoding="utf-8")
            (compiled / "rules" / "alpha.json").write_text(
                json.dumps(pack), encoding="utf-8"
            )
            (manifests / "alpha.json").write_text(
                json.dumps(manifest), encoding="utf-8"
            )
            report = production_readiness.audit_repository(
                compiled, manifests, as_of=date(2026, 8, 2)
            )
        self.assertTrue(report["summary"]["premature_verified_profiles"])


class _Response:
    def __init__(self, status: int, url: str):
        self.status = status
        self._url = url

    def __enter__(self):
        return self

    def __exit__(self, exc_type, exc, tb):
        return False

    def geturl(self) -> str:
        return self._url

    def read(self, size: int = -1) -> bytes:
        return b"x"


class SourceHealthTests(unittest.TestCase):
    def test_redirected_source_is_classified_as_moved(self) -> None:
        def opener(request, timeout):
            return _Response(200, "https://example.test/new")

        result = production_readiness.check_source(
            "src:test", "https://example.test/old", opener=opener
        )
        self.assertEqual(result["availability"], "moved")
        self.assertEqual(result["final_url"], "https://example.test/new")

    def test_http_error_is_classified_as_unavailable(self) -> None:
        def opener(request, timeout):
            raise HTTPError(request.full_url, 404, "not found", hdrs=None, fp=None)

        result = production_readiness.check_source(
            "src:test", "https://example.test/missing", opener=opener
        )
        self.assertEqual(result["availability"], "unavailable")
        self.assertEqual(result["status_code"], 404)


if __name__ == "__main__":
    unittest.main()
