"""Deterministic reference projection and hostile source tests."""

from __future__ import annotations

import hashlib
import importlib.util
import json
import sqlite3
import tempfile
import unittest
from pathlib import Path

MODULE = Path(__file__).with_name("export_utility_context.py")
spec = importlib.util.spec_from_file_location("bcm_eia_export", MODULE)
assert spec is not None and spec.loader is not None
exporter = importlib.util.module_from_spec(spec)
spec.loader.exec_module(exporter)


class ExportTests(unittest.TestCase):
    def setUp(self) -> None:
        d = tempfile.TemporaryDirectory()
        self.addCleanup(d.cleanup)
        self.root = Path(d.name)
        self.db = self.root / "catalog.sqlite"
        self.manifest = self.root / "manifest.json"
        self.out = self.root / "candidates.json"
        self.source = "f" * 64
        with sqlite3.connect(self.db) as c:
            c.executescript("""
                CREATE TABLE release(reporting_year INTEGER, source_sha256 TEXT);
                CREATE TABLE service_county(eia_id INTEGER, reporting_year INTEGER, state TEXT,
                    county_name TEXT, county_fips TEXT, source_sheet TEXT, source_row INTEGER);
                CREATE TABLE utility_snapshot(eia_id INTEGER, reporting_year INTEGER, reported_name TEXT);
                CREATE TABLE operating_activity(eia_id INTEGER, reporting_year INTEGER, state TEXT,
                    activity_group TEXT, activity_code TEXT, reported_value TEXT);
            """)
            c.execute("INSERT INTO release VALUES(?,?)", (2025, self.source))
            c.execute("INSERT INTO utility_snapshot VALUES (15466,2025,'Public Service Co of Colorado')")
            c.execute("INSERT INTO service_county VALUES (15466,2025,'CO','Denver',NULL,'Counties_States',7633)")
            c.execute("INSERT INTO operating_activity VALUES (15466,2025,'CO','activity','distribution','Y')")
        self.write_manifest()

    def write_manifest(self) -> None:
        self.manifest.write_text(json.dumps({
            "format": "electric-utility-catalog/v1", "reporting_year": 2025,
            "release_id": "eia861-2025-" + self.source[:16] + "-v0.1.2",
            "source": {"sha256": self.source, "url": "https://www.eia.gov/electricity/data/eia861/f8612025.zip"},
            "artifact": {"sha256": exporter.sha256(self.db)},
            "counts": {"service_county": 1}, "geography": None,
        }))

    def test_source_bound_export_is_deterministic_and_labeled_candidate(self) -> None:
        digest = exporter.export(self.manifest, self.db, self.out)
        original = self.out.read_bytes()
        self.assertEqual(hashlib.sha256(original).hexdigest(), digest)
        exported = json.loads(original)
        self.assertEqual(exported["scope"], "EIA-reported county candidates only; no verified serving utility")
        self.assertEqual(exported["records"][0]["activities"], [["distribution", "Y"]])
        self.assertEqual(exported["records"][0]["eia_id"], 15466)
        self.assertEqual(exporter.export(self.manifest, self.db, self.out), digest)
        self.assertEqual(self.out.read_bytes(), original)

    def test_rejects_corrupted_file_and_manifest(self) -> None:
        with sqlite3.connect(self.db) as c:
            c.execute("UPDATE utility_snapshot SET reported_name='Tampered'")
        with self.assertRaisesRegex(ValueError, "digest mismatch"):
            exporter.export(self.manifest, self.db, self.out)

    def test_rejects_mismatched_count(self) -> None:
        m = json.loads(self.manifest.read_text())
        m["counts"]["service_county"] = 2
        self.manifest.write_text(json.dumps(m))
        with self.assertRaisesRegex(ValueError, "record count"):
            exporter.export(self.manifest, self.db, self.out)

    def test_rejects_unreviewed_enrichment(self) -> None:
        m = json.loads(self.manifest.read_text())
        m["geography"] = {"source": "unreviewed"}
        self.manifest.write_text(json.dumps(m))
        with self.assertRaisesRegex(ValueError, "geographic enrichment"):
            exporter.export(self.manifest, self.db, self.out)


if __name__ == "__main__":
    unittest.main()
