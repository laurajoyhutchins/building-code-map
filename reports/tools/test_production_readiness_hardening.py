from __future__ import annotations

import json
import tempfile
import unittest
from datetime import date
from pathlib import Path

import production_readiness
from test_production_readiness import complete_inputs


class ProductionReadinessHardeningTests(unittest.TestCase):
    def test_manifest_must_require_every_acceptance_fixture_class(self) -> None:
        profile, pack, manifest = complete_inputs()
        manifest["required_fixture_kinds"] = ["incorporated"]
        with self.assertRaisesRegex(
            production_readiness.ReadinessError,
            "required_fixture_kinds must exactly match",
        ):
            production_readiness.audit_state(
                profile, pack, manifest, as_of=date(2026, 8, 2)
            )

    def test_fixture_code_families_must_be_inside_declared_scope(self) -> None:
        profile, pack, manifest = complete_inputs()
        manifest["fixtures"][0]["code_families"] = ["plumbing"]
        with self.assertRaisesRegex(
            production_readiness.ReadinessError,
            r"fixtures\[0\].code_families must be within scope",
        ):
            production_readiness.audit_state(
                profile, pack, manifest, as_of=date(2026, 8, 2)
            )

    def test_timeline_and_adoption_sources_must_be_primary(self) -> None:
        profile, pack, manifest = complete_inputs()
        profile["sources"].append(
            {
                "id": "src:aa:secondary",
                "title": "Secondary summary",
                "url": "https://example.test/secondary",
                "kind": "secondary",
                "accessed_at": "2026-08-01",
            }
        )
        pack["source_health"]["src:aa:secondary"] = {
            "last_checked_at": "2026-08-01",
            "availability": "available",
        }
        manifest["timelines"][0]["source_ids"] = ["src:aa:secondary"]
        profile["adoptions"][0]["source_ids"] = ["src:aa:secondary"]
        result = production_readiness.audit_state(
            profile, pack, manifest, as_of=date(2026, 8, 2)
        )
        self.assertIn(
            "timeline_source_not_primary:building:2021-01-01",
            result["signals"],
        )
        self.assertIn(
            "adoption_source_not_primary:adoption:aa:2021",
            result["signals"],
        )

    def test_missing_manifest_cannot_bypass_verified_profile_gate(self) -> None:
        profile, pack, _ = complete_inputs()
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            compiled = root / "compiled"
            manifests = root / "manifests"
            (compiled / "rules").mkdir(parents=True)
            manifests.mkdir()
            (compiled / "alpha.json").write_text(
                json.dumps(profile), encoding="utf-8"
            )
            (compiled / "rules" / "alpha.json").write_text(
                json.dumps(pack), encoding="utf-8"
            )
            report = production_readiness.audit_repository(
                compiled,
                manifests,
                as_of=date(2026, 8, 2),
                expected_state_slugs={"alpha"},
            )
        self.assertEqual(report["summary"]["states"], 1)
        self.assertEqual(report["states"][0]["readiness"], "blocked")
        self.assertIn("manifest_missing", report["states"][0]["signals"])
        self.assertTrue(report["summary"]["premature_verified_profiles"])


if __name__ == "__main__":
    unittest.main()
