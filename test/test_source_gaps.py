"""Hostile topology and base-grid coverage tests; synthetic geometry only."""
from __future__ import annotations

from hashlib import sha256
import json
from pathlib import Path
from tempfile import TemporaryDirectory
import unittest

from shapely.geometry import Point, Polygon, box, mapping
from pyproj import Transformer

from scripts import bulk_query
from scripts import generate_source_gaps as gaps
from scripts.generate_boundary_pairs import load_features


class TestSourceGaps(unittest.TestCase):
    def fixture(self, root, features, baseline):
        geometry = root / "geometry.geojson"
        geometry.write_text(json.dumps({
            "type": "FeatureCollection",
            "features": [{"type": "Feature", "properties": {"id": name},
                          "geometry": mapping(g)} for name, g in features],
        }, sort_keys=True))
        workload = root / "baseline.jsonl"
        workload.write_text("".join(json.dumps({
            "id": f"grid:{i}", "point": {"longitude": xy[0], "latitude": xy[1]},
        }) + "\n" for i, xy in enumerate(baseline)))
        return geometry, sha256(geometry.read_bytes()).hexdigest(), workload, sha256(workload.read_bytes()).hexdigest()

    def test_missing_small_polygon_is_sampled_even_when_outer_is_represented(self):
        with TemporaryDirectory() as folder:
            root = Path(folder)
            geo, digest, base, baseline_digest = self.fixture(root, [
                ("county", box(-106, 39, -103, 42)),
                ("town", box(-105.10, 40.10, -105.05, 40.15)),
            ], [(-104, 39.5)])
            features, _ = load_features(geo, digest)
            rows, receipt = gaps.generate(features, digest, base, baseline_digest)
            self.assertEqual(receipt["baseline_represented_features"], 1)
            self.assertEqual(receipt["missing_source_features"], 1)
            self.assertEqual(receipt["generated_samples"], 1)
            row = rows[0]
            self.assertEqual(row["gap"]["kind"], "nested_feature_candidate")
            self.assertEqual(row["gap"]["source_feature_id"], "town")
            self.assertEqual(row["gap"]["containing_feature_ids"], ["county", "town"])
            self.assertEqual(row["source_sha256"], digest)
            bulk_query.validate_sample(row)
            self.assertNotIn("jurisdiction_id", row)

    def test_empty_source_hole_probe_does_not_imply_legal_gap(self):
        with TemporaryDirectory() as folder:
            root = Path(folder)
            shell = Polygon([(-106, 39), (-103, 39), (-103, 42), (-106, 42)],
                            holes=[[(-105, 40), (-104, 40), (-104, 41), (-105, 41)]])
            geo, digest, base, baseline_digest = self.fixture(root,
                [("shell", shell), ("other", box(-102, 39, -101, 40))], [(-105.5, 39.5)])
            features, _ = load_features(geo, digest)
            rows, receipt = gaps.generate(features, digest, base, baseline_digest)
            holes = [row for row in rows if row["gap"]["kind"] == "interior_void_candidate"]
            self.assertEqual(len(holes), 1)
            self.assertEqual(holes[0]["gap"]["containing_feature_ids"], [])
            self.assertEqual(holes[0]["gap"]["classification"], "source_geometry_candidate_only")
            pt = holes[0]["point"]
            projected = Point(*Transformer.from_crs("EPSG:4326", "EPSG:6933", always_xy=True)
                              .transform(pt["longitude"], pt["latitude"]))
            self.assertEqual([name for name, shape in features.items() if shape.contains(projected)], [])
            self.assertGreaterEqual(receipt["generated_samples"], 1)

    def test_hole_filled_by_source_feature_is_not_a_void_probe(self):
        with TemporaryDirectory() as folder:
            root = Path(folder)
            shell = Polygon([(-106, 39), (-103, 39), (-103, 42), (-106, 42)],
                            holes=[[(-105, 40), (-104, 40), (-104, 41), (-105, 41)]])
            geo, digest, base, baseline_digest = self.fixture(root, [
                ("shell", shell), ("filler", box(-105, 40, -104, 41)),
            ], [(-105.5, 39.5)])
            rows, receipt = gaps.generate(load_features(geo, digest)[0], digest, base, baseline_digest)
            self.assertEqual([x for x in rows if x["gap"]["kind"] == "interior_void_candidate"], [])
            self.assertGreaterEqual(receipt["categories"].get("occupied_or_unsampleable_holes", 0), 1)

    def test_partial_overlap_cannot_infer_exclusive_feature(self):
        with TemporaryDirectory() as folder:
            root = Path(folder)
            geo, digest, base, baseline_digest = self.fixture(root, [
                ("one", box(-106, 39, -104, 41)), ("two", box(-105.5, 39.5, -103.5, 41.5)),
            ], [(-105.8, 39.2)])
            rows, receipt = gaps.generate(load_features(geo, digest)[0], digest, base, baseline_digest)
            self.assertEqual(rows, [])
            self.assertEqual(receipt["categories"]["ambiguous_overlap_skipped"], 1)

    def test_digests_bounds_duplicate_ids_fail_without_clobber(self):
        with TemporaryDirectory() as folder:
            root = Path(folder)
            geo, digest, base, baseline_digest = self.fixture(root, [
                ("big", box(-106, 39, -103, 42)),
                ("small", box(-105.1, 40.1, -105.05, 40.15)),
            ], [(-104, 39.5)])
            features, _ = load_features(geo, digest)
            with self.assertRaisesRegex(ValueError, "SHA-256"):
                gaps.generate(features, digest, base, "0" * 64)
            with self.assertRaisesRegex(ValueError, "max samples"):
                gaps.generate(features, digest, base, baseline_digest, max_samples=0)
            with self.assertRaisesRegex(ValueError, "max samples"):
                gaps.generate(features, digest, base, baseline_digest, max_samples=100001)
            base.write_text(base.read_text() * 2)
            with self.assertRaisesRegex(ValueError, "duplicate baseline"):
                gaps.generate(features, digest, base, sha256(base.read_bytes()).hexdigest())
            # Source-mode CLI never overwrites a successful result after bad input.
            base.write_text(json.dumps({"id": "fixture", "point": {"latitude": 39.5, "longitude": -104}}) + "\n")
            baseline_digest = sha256(base.read_bytes()).hexdigest()
            output, summary = root / "workload.jsonl", root / "summary.json"
            args = ["--geometry", str(geo), "--sha256", digest, "--baseline", str(base),
                    "--baseline-sha256", baseline_digest, "--output", str(output),
                    "--summary", str(summary), "--max-samples", "10"]
            gaps.main(args)
            original = output.read_bytes()
            summary_obj = json.loads(summary.read_text())
            self.assertEqual(summary_obj["workload_sha256"], sha256(original).hexdigest())
            gaps.main(args)
            self.assertEqual(output.read_bytes(), original)
            args[-1] = "0"
            with self.assertRaisesRegex(ValueError, "max samples"):
                gaps.main(args)
            self.assertEqual(output.read_bytes(), original)


if __name__ == "__main__":
    unittest.main()
