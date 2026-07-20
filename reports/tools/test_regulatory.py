from __future__ import annotations

import json
import tempfile
import unittest
from pathlib import Path

import regulatory


class RegulatoryToolTests(unittest.TestCase):
    def setUp(self) -> None:
        self.profile = json.loads(
            (Path(__file__).parents[2] / "backend/data/regulatory/colorado.json").read_text(
                encoding="utf-8"
            )
        )

    def report_text(self, profile: dict | None = None) -> str:
        profile = self.profile if profile is None else profile
        return (
            "---\n"
            "state:\n  state_id: US-CO\n  name: Colorado\n  abbreviation: CO\n"
            "report:\n  status: partially_verified\n  last_updated: 2026-07-20\n"
            "  last_verified: 2026-07-20\n"
            "risk:\n  open_questions_count: 2\n"
            "---\n\n# Colorado\n\n"
            "```json jurisdiction-profile\n"
            + json.dumps(profile, indent=2)
            + "\n```\n"
        )

    def test_extract_profile_validates_references(self) -> None:
        profile = regulatory.extract_profile(self.report_text(), "colorado.md")
        self.assertEqual(profile["state_id"], "US-CO")

    def test_unknown_adoption_reference_fails_closed(self) -> None:
        broken = json.loads(json.dumps(self.profile))
        broken["code_family_overrides"]["electrical"]["adoption_ids"] = ["missing"]
        with self.assertRaisesRegex(regulatory.ProfileError, "unknown adoption missing"):
            regulatory.extract_profile(self.report_text(broken), "broken.md")

    def test_compile_is_deterministic(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            report = root / "colorado.md"
            report.write_text(self.report_text(), encoding="utf-8")
            output = root / "out"
            first = regulatory.compile_reports([report], output)[0].read_text(encoding="utf-8")
            second = regulatory.compile_reports([report], output)[0].read_text(encoding="utf-8")
            self.assertEqual(first, second)
            self.assertTrue(first.endswith("\n"))

    def test_coverage_does_not_treat_file_presence_as_readiness(self) -> None:
        with tempfile.TemporaryDirectory() as directory:
            root = Path(directory)
            valid = root / "colorado.md"
            valid.write_text(self.report_text(), encoding="utf-8")
            missing = root / "other.md"
            missing.write_text("# No profile\n", encoding="utf-8")
            rows = regulatory.build_coverage([valid, missing], root)
            by_name = {row["report"]: row for row in rows}
            self.assertTrue(by_name["colorado.md"]["profile_valid"])
            self.assertIn(
                "resolver_follow_up_required", by_name["colorado.md"]["unresolved_signals"]
            )
            self.assertFalse(by_name["other.md"]["profile_valid"])
            self.assertIn(
                "profile_missing_or_invalid", by_name["other.md"]["unresolved_signals"]
            )


if __name__ == "__main__":
    unittest.main()
