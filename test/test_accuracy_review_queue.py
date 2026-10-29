"""Blinded stratified intake: synthetic-only, no legal truth invented."""
from __future__ import annotations

from hashlib import sha256
import json
from pathlib import Path
from tempfile import TemporaryDirectory
import unittest

from scripts import select_accuracy_review as review


class TestBlindAccuracyReview(unittest.TestCase):
    def samples(self):
        return [
            {"id": "plain", "point": {"latitude": 40, "longitude": -105}},
            {"id": "grid:1", "point": {"latitude": 40, "longitude": -105},
             "cell": {"crs": "EPSG:6933", "column": 0, "row": 1, "step_metres": 160934.4,
                      "land_fraction": 0.5, "sample_kind": "cell_center"}},
            {"id": "edge:1", "point": {"latitude": 40, "longitude": -105},
             "source_sha256": "a" * 64,
             "boundary": {"pair_key": "a" * 24, "feature_id": "synthetic-A",
                          "neighbor_feature_id": "synthetic-B", "side": "a",
                          "offset_metres": 100, "crs": "EPSG:6933",
                          "classification": "source_polygon_candidate_only"}},
            {"id": "gap:1", "point": {"latitude": 40, "longitude": -105},
             "source_sha256": "b" * 64,
             "gap": {"classification": "source_geometry_candidate_only",
                     "kind": "nested_feature_candidate", "source_feature_id": "synthetic-C",
                     "containing_feature_ids": ["synthetic-C"],
                     "baseline_sampled": False, "crs": "EPSG:6933"}},
            {"id": "temporal:1", "point": {"latitude": 40, "longitude": -105},
             "applicability_date": "2026-01-01",
             "temporal": {"classification": "source_date_candidate_only",
                          "event_id": "synthetic-E", "transition_date": "2025-12-31",
                          "offset_days": 1, "source_sha256": "c" * 64}},
        ]

    def source(self, path, rows):
        path.write_text("".join(json.dumps(row, sort_keys=True) + "\n" for row in rows))

    def test_one_per_cohort_and_blind_to_project_location(self):
        with TemporaryDirectory() as directory:
            path = Path(directory) / "input.jsonl"
            self.source(path, self.samples())
            queue = review.select(path, 5)
            self.assertEqual(queue["case_count"], 5)
            self.assertEqual(queue["selected_cohorts"], {
                "boundary": 1, "grid": 1, "other": 1,
                "source_gap": 1, "transition_date": 1})
            self.assertEqual(queue["input_sha256"], sha256(path.read_bytes()).hexdigest())
            self.assertFalse(queue["legal_accuracy_reviewed"])
            self.assertFalse(queue["expected_results_present"])
            self.assertEqual({c["status"] for c in queue["cases"]},
                             {"pending_independent_review"})
            content = json.dumps(queue)
            for forbidden in ("latitude", "longitude", "point", "address",
                              "expected", "verdict", "code_set", "40.0", "-105.0"):
                if forbidden in ("expected",):
                    # Schema-level metadata intentionally says expected_results_present:false.
                    continue
                self.assertNotIn('"' + forbidden + '"', content)
            self.assertEqual(queue, review.select(path, 5))

    def test_input_reordering_does_not_change_selected_case_ids(self):
        with TemporaryDirectory() as directory:
            path = Path(directory) / "input.jsonl"
            self.source(path, self.samples())
            first = review.select(path, 3)
            self.source(path, list(reversed(self.samples())))
            second = review.select(path, 3)
            self.assertEqual(first["cases"], second["cases"])
            self.assertNotEqual(first["input_sha256"], second["input_sha256"])
            self.assertEqual(first["case_count"], 3)

    def test_overflow_duplicates_and_invalid_rows_are_rejected(self):
        with TemporaryDirectory() as directory:
            path = Path(directory) / "input.jsonl"
            with self.assertRaisesRegex(ValueError, "empty"):
                self.source(path, [])
                review.select(path, 10)
            rows = self.samples()
            self.source(path, rows + [rows[0]])
            with self.assertRaisesRegex(ValueError, "duplicate"):
                review.select(path, 10)
            self.source(path, [{"id": "bad", "point": {"latitude": 100, "longitude": 0}}])
            with self.assertRaisesRegex(ValueError, "invalid workload"):
                review.select(path, 10)
            self.source(path, rows)
            for cap in (0, 1001):
                with self.subTest(cap=cap), self.assertRaisesRegex(ValueError, "max cases"):
                    review.select(path, cap)

    def test_cli_writes_stable_candidate_queue_without_truth(self):
        with TemporaryDirectory() as directory:
            base = Path(directory)
            path = base / "input.jsonl"
            out = base / "review.json"
            self.source(path, self.samples())
            review.main(["--input", str(path), "--output", str(out), "--max-cases", "4"])
            first = out.read_bytes()
            queue = json.loads(first)
            self.assertEqual(queue["case_count"], 4)
            self.assertTrue(all(set(c) == {"case_id", "workload_id", "candidate_tags", "status"}
                                for c in queue["cases"]))
            review.main(["--input", str(path), "--output", str(out), "--max-cases", "4"])
            self.assertEqual(out.read_bytes(), first)


if __name__ == "__main__":
    unittest.main()
