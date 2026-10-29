"""State-stratified geographic candidate coverage tests; all fixtures synthetic."""
from __future__ import annotations

from hashlib import sha256
import json
from pathlib import Path
from tempfile import TemporaryDirectory
import unittest

from shapely.geometry import mapping, box

from scripts import bulk_query
from scripts import state_coverage


class StateCoverageTests(unittest.TestCase):
    def fixture(self, directory):
        target = Path(directory) / "states.geojson"
        features = [
            {"type": "Feature", "properties": {"STATEFP": fips},
             "geometry": mapping(geometry)}
            for fips, geometry in [
                ("01", box(-103, 35, -102, 36)),
                ("02", box(-102, 35, -101, 36)),
            ]
        ]
        target.write_text(json.dumps({"type": "FeatureCollection", "features": features}))
        return target, sha256(target.read_bytes()).hexdigest()

    def test_boundary_overlap_unknown_and_unsampled_state_are_separate(self):
        with TemporaryDirectory() as d:
            source, digest = self.fixture(d)
            states, observed = state_coverage.load_states(source, digest)
            self.assertEqual(observed, digest)
            grid = Path(d) / "grid.jsonl"
            rows = [
                {"id": "within-A", "point": {"longitude": -102.5, "latitude": 35.5}},
                {"id": "boundary-A-B", "point": {"longitude": -102, "latitude": 35.5}},
                {"id": "outside", "point": {"longitude": -110, "latitude": 35.5}},
            ]
            grid.write_text("".join(json.dumps(row) + "\n" for row in rows))
            result = state_coverage.assess(states, digest, grid)
            self.assertEqual(result["samples"], 3)
            self.assertEqual(result["state_candidate_counts"], {"01": 1, "02": 0})
            self.assertEqual(result["unsampled_states"], ["02"])
            self.assertEqual(result["unmatched_points"], 1)
            self.assertEqual(result["overlapping_candidate_points"], 1)
            self.assertEqual(result["overlapping_ids"], ["boundary-A-B"])
            self.assertEqual(result["unmatched_ids"], ["outside"])

    def test_representatives_have_source_hash_and_are_canonical_workload_points(self):
        with TemporaryDirectory() as d:
            source, digest = self.fixture(d)
            states, _ = state_coverage.load_states(source, digest)
            rows = list(state_coverage.state_representatives(states, digest))
            self.assertEqual(len(rows), 2)
            self.assertEqual([r["id"] for r in rows], ["state-probe:01", "state-probe:02"])
            for row in rows:
                self.assertEqual(row["source_sha256"], digest)
                self.assertEqual(bulk_query.validate_sample(row), row)
                self.assertNotIn("jurisdiction_id", row)
                self.assertNotIn("code_set", row)
            with self.assertRaises(ValueError):
                state_coverage.load_states(source, "0" * 64)
            with self.assertRaises(ValueError):
                state_coverage.load_states(source, digest, expect_51=True)

    def test_duplicate_state_fips_and_grid_ids_rejected(self):
        with TemporaryDirectory() as d:
            source, digest = self.fixture(d)
            content = json.loads(source.read_text())
            content["features"].append(content["features"][0])
            source.write_text(json.dumps(content))
            with self.assertRaises(ValueError):
                state_coverage.load_states(source, sha256(source.read_bytes()).hexdigest())
            source, digest = self.fixture(d)
            states, _ = state_coverage.load_states(source, digest)
            grid = Path(d) / "grid.jsonl"
            row = {"id": "repeat", "point": {"longitude": -102.5, "latitude": 35.5}}
            grid.write_text(json.dumps(row) + "\n" + json.dumps(row) + "\n")
            with self.assertRaises(ValueError):
                state_coverage.assess(states, digest, grid)


if __name__ == "__main__":
    unittest.main()
