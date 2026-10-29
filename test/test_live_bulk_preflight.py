"""Fail-closed real-provider preflight, synthetic local HTTP fixtures only."""
from __future__ import annotations

from http.server import BaseHTTPRequestHandler, ThreadingHTTPServer
import json
from pathlib import Path
from tempfile import TemporaryDirectory
from threading import Thread
from types import SimpleNamespace
import unittest

from scripts import bulk_query
from scripts import live_bulk_preflight as live

BUNDLE = "sha256:independently-reviewed-12345678"
SOURCE = "f" * 40
DATE = "2026-10-28"


class TestLiveBulkPreflight(unittest.TestCase):
    def test_live_canary_requires_pinned_source_and_private_adapter(self):
        state = {"demo": False, "ready": True, "bundle": BUNDLE, "hits": 0}

        class Handler(BaseHTTPRequestHandler):
            def log_message(self, *_):
                pass

            def reply(self, status, value):
                raw = json.dumps(value).encode()
                self.send_response(status)
                self.send_header("Content-Type", "application/json")
                self.send_header("Content-Length", str(len(raw)))
                self.end_headers()
                self.wfile.write(raw)

            def do_GET(self):
                if self.path == "/health":
                    self.reply(200, {"status": "ok", "data_scope":
                                    "synthetic_demo_only" if state["demo"] else "private_provider_adapter"})
                elif self.path == "/v1/readiness?capability=coordinate_resolution":
                    self.reply(200, {"status": "ready" if state["ready"] else "unavailable"})
                else:
                    self.reply(404, {})

            def do_POST(self):
                state["hits"] += 1
                req = json.loads(self.rfile.read(int(self.headers["Content-Length"])))
                self.reply(200, {"schema_version": "0.3", "project_id": req["project_id"],
                                 "as_of": req["applicability_date"], "jurisdiction_id": "CO-DENVER",
                                 "verdict": "not_verified",
                                 "code_set": [{"family": "building", "status": "insufficient_evidence"}],
                                 "evidence": [], "provenance": {
                                     "bundle_id": state["bundle"],
                                     "source_commit": SOURCE,
                                 }})

        server = ThreadingHTTPServer(("127.0.0.1", 0), Handler)
        thread = Thread(target=server.serve_forever, daemon=True)
        thread.start()
        try:
            with TemporaryDirectory() as directory:
                canary = Path(directory) / "canary.jsonl"
                canary.write_text(json.dumps({
                    "id": "reviewed-canary",
                    "point": {"longitude": -104.97, "latitude": 39.77},
                }) + "\n")
                args = SimpleNamespace(
                    endpoint="http://127.0.0.1:" + str(server.server_port) + bulk_query.PATH,
                    canary=str(canary), date=DATE, expected_bundle=BUNDLE,
                    expected_source=SOURCE, timeout=2.0,
                    token_env="", allow_remote=False,
                )
                accepted = live.run(args)
                self.assertEqual(accepted["category"], "runtime_canary_only")
                self.assertEqual(accepted["project_verdict"], "not_verified")
                self.assertFalse(accepted["legal_accuracy_reviewed"])
                self.assertFalse(accepted["bulk_run_authorized"])
                self.assertEqual(accepted["expected_runtime_identity"]["bundle_id"], BUNDLE)
                self.assertEqual(state["hits"], 1)

                state["bundle"] = "sha256:different-bundle"
                with self.assertRaisesRegex(ValueError, "identity"):
                    live.run(args)
                state["bundle"] = BUNDLE

                state["demo"] = True
                with self.assertRaisesRegex(ValueError, "adapter"):
                    live.run(args)
                state["demo"] = False
                self.assertEqual(state["hits"], 2)

                state["ready"] = False
                with self.assertRaisesRegex(ValueError, "readiness"):
                    live.run(args)
                state["ready"] = True
                self.assertEqual(state["hits"], 2)

                args.expected_bundle = "DEMO-XX"
                with self.assertRaises(ValueError):
                    live.run(args)
                args.expected_bundle = BUNDLE
                args.expected_source = "z" * 40
                with self.assertRaises(ValueError):
                    live.run(args)
                args.expected_source = SOURCE

                canary.write_text(canary.read_text() * 2)
                with self.assertRaisesRegex(ValueError, "exactly one sample"):
                    live.run(args)
        finally:
            server.shutdown()
            server.server_close()
            thread.join(timeout=2)

    def test_bulk_result_keeps_identity_drift_distinct_from_legal_uncertainty(self):
        def response(bundle):
            return json.dumps({
                "schema_version": "0.3", "project_id": "sample", "as_of": DATE,
                "jurisdiction_id": "CO-DENVER", "verdict": "not_verified",
                "code_set": [{"family": "building", "status": "insufficient_evidence"}],
                "provenance": {"bundle_id": bundle, "source_commit": SOURCE},
            }).encode()
        self.assertEqual(
            bulk_query.classify(200, response(BUNDLE), "sample", DATE, BUNDLE, SOURCE)["category"],
            "not_verified")
        self.assertEqual(
            bulk_query.classify(200, response("sha256:other"), "sample", DATE, BUNDLE, SOURCE)["category"],
            "runtime_identity_mismatch")
        self.assertEqual(
            bulk_query.classify(200, response(BUNDLE), "sample", DATE, BUNDLE, "a" * 40)["category"],
            "runtime_identity_mismatch")
        self.assertEqual(
            bulk_query.classify(200, response(BUNDLE).replace(b"CO-DENVER", b"DEMO-XX"),
                                "sample", DATE, BUNDLE, SOURCE)["category"],
            "runtime_identity_mismatch")


if __name__ == "__main__":
    unittest.main()
