"""Boundary generator hostile cases use synthetic polygons, never actual legal maps."""
from __future__ import annotations

from hashlib import sha256
import json
from pathlib import Path
from tempfile import TemporaryDirectory
import unittest

from pyproj import Transformer
from shapely.geometry import Point, box, mapping

from scripts import bulk_query
from scripts import generate_boundary_pairs as edges


class TestBoundaryPairs(unittest.TestCase):
    def source(self, folder, named):
        path = Path(folder) / "source.geojson"
        path.write_text(json.dumps({"type": "FeatureCollection", "features": [
            {"type": "Feature", "properties": {"id": label}, "geometry": mapping(polygon)}
            for label, polygon in named]}, sort_keys=True))
        return path, sha256(path.read_bytes()).hexdigest()

    def test_exact_shared_edges_and_exclusive_membership_three_features(self):
        with TemporaryDirectory() as folder:
            examples = [("B", box(-105, 39, -104, 40)),
                        ("C", box(-104, 39, -103, 40)),
                        ("A", box(-106, 39, -105, 40))]
            source, digest = self.source(folder, examples)
            features, observed = edges.load_features(source, digest)
            self.assertEqual(observed, digest)
            rows, summary = edges.generate(features, digest, 100)
            self.assertEqual(summary["sampled_pairs"], 2)
            self.assertEqual(summary["candidate_touch_pairs"], 2)
            self.assertEqual(summary["overlapping_source_pairs"], 0)
            self.assertEqual(summary["unresolved_touch_pairs"], 0)
            self.assertEqual(len(rows), 4)
            self.assertEqual(len({r["id"] for r in rows}), 4)
            forward = Transformer.from_crs("EPSG:4326", "EPSG:6933", always_xy=True).transform
            for row in rows:
                self.assertEqual(row["source_sha256"], digest)
                self.assertEqual(bulk_query.validate_sample(row), row)
                self.assertEqual(row["boundary"]["classification"], "source_polygon_candidate_only")
                self.assertNotIn("jurisdiction_id", row)
                self.assertNotIn("code_set", row)
                point = Point(*forward(row["point"]["longitude"], row["point"]["latitude"]))
                matches = [key for key, polygon in features.items() if polygon.contains(point)]
                self.assertEqual(matches, [row["boundary"]["feature_id"]])
            # Original feature order changes the source digest but not the
            # paired coordinates, IDs, or source feature identity.
            reversed_source, reverse_digest = self.source(folder, list(reversed(examples)))
            reverse_features, _ = edges.load_features(reversed_source, reverse_digest)
            reversed_rows, _ = edges.generate(reverse_features, reverse_digest, 100)
            self.assertEqual([r["id"] for r in rows], [r["id"] for r in reversed_rows])
            self.assertEqual([r["point"] for r in rows], [r["point"] for r in reversed_rows])

    def test_overlapping_polygons_are_not_a_boundary(self):
        with TemporaryDirectory() as folder:
            source, digest = self.source(folder, [("A", box(-106, 39, -104, 40)),
                                                   ("B", box(-105, 39, -103, 40))])
            rows, summary = edges.generate(edges.load_features(source, digest)[0], digest, 100)
            self.assertEqual(rows, [])
            self.assertEqual(summary["overlapping_source_pairs"], 1)
            self.assertEqual(summary["candidate_touch_pairs"], 0)

    def test_point_only_contact_is_not_a_shared_edge(self):
        with TemporaryDirectory() as folder:
            source, digest = self.source(folder, [("A", box(-106, 39, -105, 40)),
                                                   ("B", box(-105, 40, -104, 41))])
            rows, summary = edges.generate(edges.load_features(source, digest)[0], digest, 100)
            self.assertEqual(rows, [])
            self.assertEqual(summary["candidate_touch_pairs"], 0)

    def test_checksum_duplicates_and_bounded_pair_limit(self):
        with TemporaryDirectory() as folder:
            source, digest = self.source(folder, [("A", box(-106, 39, -105, 40)),
                                                   ("B", box(-105, 39, -104, 40)),
                                                   ("C", box(-104, 39, -103, 40))])
            with self.assertRaisesRegex(ValueError, "SHA-256"):
                edges.load_features(source, "0" * 64)
            features, _ = edges.load_features(source, digest)
            with self.assertRaisesRegex(ValueError, "max-pairs"):
                edges.generate(features, digest, 100, max_pairs=1)
            for spacing in [0, -1, 5001, float("nan"), float("inf")]:
                with self.subTest(spacing=spacing), self.assertRaises(ValueError):
                    edges.generate(features, digest, spacing)
            source, digest = self.source(folder, [("A", box(-106, 39, -105, 40)),
                                                   ("A", box(-105, 39, -104, 40))])
            with self.assertRaisesRegex(ValueError, "unique"):
                edges.load_features(source, digest)

    def test_cli_output_sha_and_no_partial_overwrite_on_failure(self):
        with TemporaryDirectory() as folder:
            source, digest = self.source(folder, [("A", box(-106, 39, -105, 40)),
                                                   ("B", box(-105, 39, -104, 40))])
            output = Path(folder) / "pairs.jsonl"
            summary = Path(folder) / "summary.json"
            args = ["--geometry", str(source), "--sha256", digest,
                    "--output", str(output), "--summary", str(summary),
                    "--max-pairs", "10"]
            edges.main(args)
            first = output.read_bytes()
            receipt = json.loads(summary.read_text())
            self.assertEqual(receipt["schema_version"], "1")
            self.assertEqual(receipt["source_geometry_sha256"], digest)
            self.assertEqual(receipt["workload_sha256"], sha256(first).hexdigest())
            self.assertEqual(receipt["samples"], 2)
            edges.main(args)
            self.assertEqual(output.read_bytes(), first)
            args[-1] = "0"
            with self.assertRaisesRegex(ValueError, "max pairs"):
                edges.main(args)
            self.assertEqual(output.read_bytes(), first)


if __name__ == "__main__":
    unittest.main()
