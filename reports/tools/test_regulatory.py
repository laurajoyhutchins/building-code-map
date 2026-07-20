from __future__ import annotations

import json
import tempfile
import unittest
from datetime import date
from pathlib import Path

import regulatory


class RegulatoryToolTests(unittest.TestCase):
    def setUp(self) -> None:
        verification = {"status": "verified", "confidence": 0.95}
        self.profile = {
            "schema_version": "1.0",
            "profile_id": "state-profile:us-aa",
            "state_id": "US-AA",
            "state_name": "Alpha",
            "state_abbreviation": "AA",
            "state_fips": "01",
            "status": "verified",
            "last_verified": "2026-07-20",
            "sources": [
                {
                    "id": "src:aa",
                    "title": "Alpha official source",
                    "url": "https://example.com/alpha",
                    "kind": "official_web",
                    "accessed_at": "2026-07-20",
                }
            ],
            "authorities": [
                {
                    "id": "auth:aa",
                    "name": "Alpha Board",
                    "type": "state_board",
                    "roles": ["adopts"],
                    "source_ids": ["src:aa"],
                    "verification": verification,
                }
            ],
            "relationships": [
                {
                    "id": "edge:aa",
                    "from_id": "auth:aa",
                    "relationship": "delegates_enforcement_to",
                    "to": "local agencies",
                    "scope": ["building"],
                    "source_ids": ["src:aa"],
                    "verification": verification,
                }
            ],
            "adoptions": [
                {
                    "id": "adoption:aa",
                    "code_family": "building",
                    "status": "current",
                    "state_code_name": "Alpha Building Code",
                    "adopting_authority_id": "auth:aa",
                    "enforcement_model": "local",
                    "dates": {"effective_date": "2024-01-01"},
                    "source_ids": ["src:aa"],
                    "verification": verification,
                }
            ],
            "defaults": {
                "incorporated": {
                    "status": "partially_resolved",
                    "adoption_ids": ["adoption:aa"],
                    "source_ids": ["src:aa"],
                },
                "unincorporated": {
                    "status": "partially_resolved",
                    "adoption_ids": ["adoption:aa"],
                    "source_ids": ["src:aa"],
                },
            },
            "code_family_overrides": {},
            "project_type_overrides": {},
            "verification": verification,
        }
        self.pack = {
            "schema_version": "1.0",
            "state_id": "US-AA",
            "source_health": {
                "src:aa": {
                    "last_checked_at": "2026-07-20",
                    "availability": "available",
                }
            },
            "applicability_rules": [
                {
                    "id": "rule:aa:app",
                    "code_family": "building",
                    "trigger": "ordinary_project",
                    "summary": "The state code applies.",
                    "source_ids": ["src:aa"],
                    "verification": verification,
                }
            ],
            "date_rules": [
                {
                    "id": "rule:aa:date",
                    "code_family": "building",
                    "rule_type": "effective_date",
                    "trigger": "applicability_date",
                    "start_date": "2024-01-01",
                    "summary": "Effective on 2024-01-01.",
                    "source_ids": ["src:aa"],
                    "verification": verification,
                }
            ],
            "amendment_rules": [
                {
                    "id": "rule:aa:amend",
                    "code_family": "building",
                    "level": "local",
                    "posture": "check_local_record",
                    "summary": "Check local amendments.",
                    "source_ids": ["src:aa"],
                    "verification": verification,
                }
            ],
            "enforcement_rules": [
                {
                    "id": "rule:aa:enforce",
                    "code_family": "building",
                    "model": "local",
                    "entity_kinds": ["municipality"],
                    "authority_ids": ["auth:aa"],
                    "summary": "Local enforcement.",
                    "source_ids": ["src:aa"],
                    "verification": verification,
                }
            ],
            "claims": [
                {
                    "id": "claim:aa",
                    "subject_id": "adoption:aa",
                    "field": "dates.effective_date",
                    "status": "supported",
                    "value": "2024-01-01",
                    "source_ids": ["src:aa"],
                    "verification": verification,
                }
            ],
            "resolver_fixture_ids": ["fixture:aa:building"],
        }

    def report_text(self, profile: dict | None = None, pack: dict | None = None) -> str:
        profile = self.profile if profile is None else profile
        pack = self.pack if pack is None else pack
        return (
            "---\n"
            "state:\n  state_id: US-AA\n  name: Alpha\n  abbreviation: AA\n"
            "report:\n  status: verified\n  last_updated: 2026-07-20\n"
            "  last_verified: 2026-07-20\n"
            "risk:\n  open_questions_count: 0\n"
            "---\n\n# Alpha\n\n"
            "```json jurisdiction-profile\n"
            + json.dumps(profile, separators=(",", ":"))
            + "\n```\n\n"
            "```json jurisdiction-rules\n"
            + json.dumps(pack, separators=(",", ":"))
            + "\n```\n"
        )

    def test_extract_profile_and_rule_pack_validate_references(self) -> None:
        text = self.report_text()
        profile = regulatory.extract_profile(text, "alpha.md")
        pack = regulatory.extract_rule_pack(text, profile, "alpha.md", required=True)
        self.assertEqual(pack["state_id"], "US-AA")

    def test_unknown_adoption_reference_fails_closed(self) -> None:
        broken = json.loads(json.dumps(self.profile))
        broken["defaults"]["incorporated"]["adoption_ids"] = ["missing"]
        with self.assertRaisesRegex(regulatory.ProfileError, "unknown adoption missing"):
            regulatory.extract_profile(self.report_text(broken), "broken.md")

    def test_rule_pack_unknown_source_fails_closed(self) -> None:
        broken = json.loads(json.dumps(self.pack))
        broken["source_health"]["src:missing"] = {
            "last_checked_at": "2026-07-20",
            "availability": "available",
        }
        profile = regulatory.extract_profile(self.report_text(), "alpha.md")
        with self.assertRaisesRegex(regulatory.ProfileError, "unknown source src:missing"):
            regulatory.extract_rule_pack(
                self.report_text(pack=broken),
                profile,
                "broken.md",
                required=True,
            )

    def test_compile_is_deterministic_and_writes_rule_pack(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            report = root / "alpha.md"
            report.write_text(self.report_text(), encoding="utf-8")
            output = root / "out"
            first_paths = regulatory.compile_reports([report], output)
            first = [path.read_text(encoding="utf-8") for path in first_paths]
            second_paths = regulatory.compile_reports([report], output)
            second = [path.read_text(encoding="utf-8") for path in second_paths]
            self.assertEqual(first, second)
            self.assertTrue((output / "rules" / "alpha.json").exists())

    def test_coverage_reports_rule_families_and_readiness(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            report = root / "alpha.md"
            report.write_text(self.report_text(), encoding="utf-8")
            row = regulatory.build_coverage(
                [report],
                root,
                as_of=date(2026, 7, 20),
            )[0]
            self.assertEqual(row["readiness"], "production_ready")
            self.assertEqual(row["applicability_rule_count"], 1)
            self.assertEqual(row["resolver_fixture_count"], 1)

    def test_coverage_flags_stale_sources_and_conflicting_claims(self) -> None:
        pack = json.loads(json.dumps(self.pack))
        pack["source_health"]["src:aa"]["last_checked_at"] = "2024-01-01"
        pack["claims"][0]["status"] = "conflicting"
        pack["claims"][0]["conflict_group"] = "conflict:aa:date"
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            report = root / "alpha.md"
            report.write_text(self.report_text(pack=pack), encoding="utf-8")
            row = regulatory.build_coverage(
                [report],
                root,
                as_of=date(2026, 7, 20),
            )[0]
            self.assertIn("stale_sources", row["unresolved_signals"])
            self.assertIn("conflicting_claims", row["unresolved_signals"])
            self.assertEqual(row["readiness"], "pilot_ready")

    def test_coverage_does_not_treat_file_presence_as_readiness(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            missing = root / "other.md"
            missing.write_text("# No profile\n", encoding="utf-8")
            row = regulatory.build_coverage(
                [missing],
                root,
                as_of=date(2026, 7, 20),
            )[0]
            self.assertFalse(row["profile_valid"])
            self.assertEqual(row["readiness"], "blocked")
            self.assertIn("profile_missing_or_invalid", row["unresolved_signals"])


if __name__ == "__main__":
    unittest.main()
