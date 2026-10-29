"""Synthetic/hostile tests for BCM bulk grid generation and query driving."""
from __future__ import annotations

import hashlib
from http.server import BaseHTTPRequestHandler, ThreadingHTTPServer
import importlib.util
import json
from pathlib import Path
from tempfile import TemporaryDirectory
from threading import Thread
from types import SimpleNamespace
import unittest

from scripts import bulk_query as bulk
from scripts import generate_grid as grid


class TestBulkQuery(unittest.TestCase):
    def test_sample_rejects_bad_location_and_id(self):
        for row in [
            {}, {"id": "a", "point": {"longitude": float("nan"), "latitude": 2}},
            {"id": "a", "point": {"longitude": 200, "latitude": 2}},
            {"id": "a", "point": {"longitude": 1, "latitude": 2}, "address": "also"},
            {"id": "a", "point": {"longitude": 1}},
            {"id": "contains space", "address": "Demo"},
        ]:
            with self.subTest(row=row), self.assertRaises((ValueError, TypeError)):
                bulk.validate_sample(row)

    def test_target_requires_opt_in_for_remote_and_blocks_credentials(self):
        for target, opt in [
            ("https://example.org" + bulk.PATH, False),
            ("http://example.org" + bulk.PATH, True),
            ("http://user:secret@localhost:8000" + bulk.PATH, False),
            ("http://localhost:8000/other", False),
            ("http://localhost:8000" + bulk.PATH + "?token=x", False),
        ]:
            with self.subTest(target=target), self.assertRaises(ValueError):
                bulk.validate_target(target, opt)
        bulk.validate_target("http://127.0.0.1:8888" + bulk.PATH, False)

    def test_response_categories_do_not_promote_unknowns(self):
        def good(verdict, evidence=None):
            return json.dumps({
                "schema_version": "0.3", "project_id": "sample", "as_of": "2026-10-28",
                "jurisdiction_id": "DEMO-XX", "verdict": verdict,
                "code_set": [{"family": "building", "status": "resolved"}],
                "evidence": evidence or [],
            }).encode()
        self.assertEqual(bulk.classify(200, good("conditional"), "sample", "2026-10-28")["category"], "conditional")
        self.assertEqual(bulk.classify(200, good("verified"), "sample", "2026-10-28")["category"], "unsupported_verification")
        self.assertEqual(bulk.classify(200, good("verified", [{"id": "source"}]), "sample", "2026-10-28")["category"], "verified")
        self.assertEqual(bulk.classify(200, good("verified"), "different", "2026-10-28")["category"], "invalid_response")
        self.assertEqual(bulk.classify(200, b"not-json", "sample", "2026-10-28")["category"], "invalid_response")
        self.assertEqual(bulk.classify(422, b'{"error":{"code":"boundary_ambiguous"}}', "sample", "2026-10-28")["category"], "unresolved")
        self.assertEqual(bulk.classify(422, b'{"error":{"code":"other"}}', "sample", "2026-10-28")["category"], "http_error")
        self.assertEqual(bulk.classify(503, b'{"error":{"code":"geocoder_unavailable"}}', "sample", "2026-10-28")["category"], "service_error")

    def test_end_to_end_with_local_fake_api(self):
        class Handler(BaseHTTPRequestHandler):
            requests = []
            def do_POST(self):
                data = json.loads(self.rfile.read(int(self.headers["Content-Length"])))
                Handler.requests.append(data)
                self.send_header("Content-Type", "application/json")
                if data["project_id"] == "unknown":
                    status, body = 422, {"error": {"code": "outside_supported_coverage", "retryable": False}}
                elif data["project_id"] == "broken":
                    status, body = 503, {"error": {"code": "geocoder_unavailable", "retryable": True}}
                else:
                    status, body = 200, {
                        "schema_version": "0.3", "project_id": data["project_id"],
                        "as_of": data["applicability_date"], "jurisdiction_id": "DEMO-XX",
                        "verdict": "not_verified",
                        "code_set": [{"family": "building", "status": "insufficient_evidence"}],
                        "evidence": [],
                    }
                result = json.dumps(body).encode()
                self.send_response(status)
                self.send_header("Content-Length", str(len(result)))
                self.end_headers()
                self.wfile.write(result)
            def send_header(self, k, v):
                if not hasattr(self, "_headers_buffer"):
                    self._headers_buffer = []
                super().send_header(k, v)
            def log_message(self, *_):
                pass

        # Send headers only after status: patch the normal handler with a clean override.
        def response(self):
            data = json.loads(self.rfile.read(int(self.headers["Content-Length"])))
            Handler.requests.append(data)
            if data["project_id"] == "unknown":
                status, body = 422, {"error": {"code": "outside_supported_coverage"}}
            elif data["project_id"] == "broken":
                status, body = 503, {"error": {"code": "geocoder_unavailable"}}
            else:
                status, body = 200, {
                    "schema_version": "0.3", "project_id": data["project_id"],
                    "as_of": data["applicability_date"], "jurisdiction_id": "DEMO-XX",
                    "verdict": "not_verified",
                    "code_set": [{"family": "building", "status": "insufficient_evidence"}],
                    "evidence": [],
                }
            raw = json.dumps(body).encode()
            self.send_response(status)
            self.send_header("Content-Type", "application/json")
            self.send_header("Content-Length", str(len(raw)))
            self.end_headers()
            self.wfile.write(raw)
        Handler.do_POST = response
        server = ThreadingHTTPServer(("127.0.0.1", 0), Handler)
        thread = Thread(target=server.serve_forever, daemon=True)
        thread.start()
        try:
            with TemporaryDirectory() as directory:
                home = Path(directory)
                rows = [{"id": key, "point": {"latitude": 39.0, "longitude": -105.0}}
                        for key in ("unresolved", "unknown", "broken")]
                source = home / "samples.jsonl"
                source.write_text("".join(json.dumps(row) + "\n" for row in rows))
                args = SimpleNamespace(
                    input=str(source), results=str(home / "results.jsonl"),
                    summary=str(home / "summary.json"), endpoint=f"http://127.0.0.1:{server.server_port}" + bulk.PATH,
                    date="2026-10-28", concurrency=3, qps=100,
                    max_requests=10, timeout=2, project_type="", token_env="",
                    allow_remote=False, shard_index=0, shard_count=1,
                )
                report = bulk.run(args)
                self.assertEqual(report["samples"], 3)
                self.assertEqual(report["categories"], {"not_verified": 1, "unresolved": 1, "service_error": 1})
                self.assertEqual(report["input_sha256"], hashlib.sha256(source.read_bytes()).hexdigest())
                self.assertEqual(len((home / "results.jsonl").read_text().splitlines()), 3)
                self.assertTrue(all(r["applicability_date"] == "2026-10-28" for r in Handler.requests))
                args.max_requests = 2
                with self.assertRaises(ValueError):
                    bulk.run(args)
                args.max_requests = 10
                args.shard_count = 2
                args.shard_index = 0
                shard = bulk.run(args)
                expected = sum(
                    int.from_bytes(hashlib.sha256(row["id"].encode()).digest()[:8], "big") % 2 == 0
                    for row in rows
                )
                self.assertEqual(shard["samples"], expected)
                self.assertEqual(shard["skipped_samples"], len(rows) - expected)
        finally:
            server.shutdown()
            server.server_close()
            thread.join(timeout=2)


@unittest.skipUnless(
    importlib.util.find_spec("shapely") and importlib.util.find_spec("pyproj"),
    "grid geometry dependencies are not installed",
)
class TestGrid(unittest.TestCase):
    def test_pinned_land_geometry_and_deterministic_cells(self):
        from shapely.geometry import box, mapping
        with TemporaryDirectory() as directory:
            path = Path(directory) / "land.geojson"
            synthetic = {"type": "Feature", "properties": {},
                         "geometry": mapping(box(-106, 39, -103, 41))}
            path.write_text(json.dumps(synthetic, sort_keys=True))
            digest = hashlib.sha256(path.read_bytes()).hexdigest()
            geometry, verified = grid.load_polygon(path, digest)
            self.assertEqual(verified, digest)
            first = list(grid.generate(geometry, digest, 100))
            self.assertTrue(first)
            self.assertEqual(first, list(grid.generate(geometry, digest, 100)))
            self.assertEqual(len({r["id"] for r in first}), len(first))
            for row in first:
                self.assertEqual(row["source_sha256"], digest)
                self.assertEqual(row["cell"]["crs"], "EPSG:6933")
                self.assertTrue(0 < row["cell"]["land_fraction"] <= 1)
                bulk.validate_sample(row)
            with self.assertRaises(ValueError):
                grid.load_polygon(path, "0" * 64)

    def test_invalid_grid_spacing(self):
        from shapely.geometry import box
        for spacing in [0, -1, float("nan"), float("inf"), 1001]:
            with self.assertRaises(ValueError):
                list(grid.generate(box(-1, -1, 1, 1), "0" * 64, spacing))


if __name__ == "__main__":
    unittest.main()
