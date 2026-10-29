"""Source-date candidates, with synthetic fixtures only; never accepted enacted law."""
from __future__ import annotations

from hashlib import sha256
import json
from pathlib import Path
from tempfile import TemporaryDirectory
import unittest

from scripts import bulk_query
from scripts import generate_temporal_probes as temporal

EVENT = {"id": "synthetic-transition", "transition_date": "2025-12-31",
         "source_sha256": "a" * 64}


class TestTemporalWorkload(unittest.TestCase):
    def fixture(self, root, events=None):
        base = root / "points.jsonl"
        base.write_text(json.dumps({
            "id": "fixture-1", "point": {"latitude": 39.75, "longitude": -104.99},
            "applicability_date": "2020-01-01", "project_type": "commercial",
            "facts": {"occupancy": "business"},
        }) + "\n")
        event_path = root / "events.json"
        event_path.write_text(json.dumps({"schema_version": "1",
                                          "classification": "source_date_candidate_only",
                                          "events": events if events is not None else [EVENT]}))
        return base, event_path

    def test_year_boundary_event_generates_three_distinct_bounded_requests(self):
        with TemporaryDirectory() as directory:
            root = Path(directory)
            base, events = self.fixture(root)
            parsed, _ = temporal.parse_events(events)
            rows, digest, source_samples = temporal.expand(base, parsed, 3)
            self.assertEqual(source_samples, 1)
            self.assertEqual(digest, sha256(base.read_bytes()).hexdigest())
            self.assertEqual(len(rows), 3)
            self.assertEqual(len({row["id"] for row in rows}), 3)
            self.assertEqual({row["applicability_date"] for row in rows},
                             {"2025-12-30", "2025-12-31", "2026-01-01"})
            for row in rows:
                self.assertEqual(bulk_query.validate_sample(row), row)
                self.assertEqual(row["temporal"]["classification"], "source_date_candidate_only")
                self.assertEqual(row["temporal"]["source_sha256"], "a" * 64)
                self.assertEqual(row["project_type"], "commercial")
                self.assertEqual(row["facts"], {"occupancy": "business"})
                self.assertNotIn("code_set", row)
            with self.assertRaisesRegex(ValueError, "exceed max requests"):
                temporal.expand(base, parsed, 2)

    def test_manifest_rejects_claims_duplicates_bad_dates_and_missing_hash(self):
        with TemporaryDirectory() as directory:
            root = Path(directory)
            _, events = self.fixture(root)
            examples = [
                {"schema_version": "1", "classification": "verified_legal_transition", "events": [EVENT]},
                {"schema_version": "1", "classification": "source_date_candidate_only",
                 "events": [EVENT, EVENT]},
                {"schema_version": "1", "classification": "source_date_candidate_only",
                 "events": [{**EVENT, "transition_date": "2025-02-30"}]},
                {"schema_version": "1", "classification": "source_date_candidate_only",
                 "events": [{**EVENT, "transition_date": "0001-01-01"}]},
                {"schema_version": "1", "classification": "source_date_candidate_only",
                 "events": [{**EVENT, "transition_date": "9999-12-31"}]},
                {"schema_version": "1", "classification": "source_date_candidate_only",
                 "events": [{**EVENT, "source_sha256": "unbound"}]},
                {"schema_version": "1", "classification": "source_date_candidate_only",
                 "events": [{**EVENT, "authority": "fictional"}]},
            ]
            for doc in examples:
                with self.subTest(doc=doc):
                    events.write_text(json.dumps(doc))
                    with self.assertRaises(ValueError):
                        temporal.parse_events(events)

    def test_recursive_temporal_and_duplicate_ids_refused(self):
        with TemporaryDirectory() as directory:
            root = Path(directory)
            base, events = self.fixture(root)
            parsed, _ = temporal.parse_events(events)
            row = json.loads(base.read_text())
            base.write_text(json.dumps(row) + "\n" + json.dumps(row) + "\n")
            with self.assertRaisesRegex(ValueError, "duplicate"):
                temporal.expand(base, parsed, 100)
            row["temporal"] = {"classification": "already-temporal"}
            base.write_text(json.dumps(row) + "\n")
            with self.assertRaisesRegex(ValueError, "recursive temporal"):
                temporal.expand(base, parsed, 100)

    def test_cli_retains_byte_digests_and_does_not_overwrite_on_failure(self):
        with TemporaryDirectory() as directory:
            root = Path(directory)
            base, events = self.fixture(root)
            output, report = root / "output.jsonl", root / "receipt.json"
            args = ["--input", str(base), "--events", str(events), "--output", str(output),
                    "--summary", str(report), "--max-requests", "3"]
            temporal.main(args)
            original = output.read_bytes()
            summary = json.loads(report.read_text())
            self.assertEqual(summary["samples"], 3)
            self.assertEqual(summary["source_samples"], 1)
            self.assertEqual(summary["candidate_events"], 1)
            self.assertEqual(summary["workload_sha256"], sha256(original).hexdigest())
            self.assertEqual(summary["transition_manifest_sha256"],
                             sha256(events.read_bytes()).hexdigest())
            temporal.main(args)
            self.assertEqual(output.read_bytes(), original)
            args[-1] = "2"
            with self.assertRaisesRegex(ValueError, "exceed max requests"):
                temporal.main(args)
            self.assertEqual(output.read_bytes(), original)


if __name__ == "__main__":
    unittest.main()
