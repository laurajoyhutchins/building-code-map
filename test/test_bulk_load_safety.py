"""Bounded localhost-only concurrency and circuit-breaker tests."""
from __future__ import annotations

from hashlib import sha256
from http.server import BaseHTTPRequestHandler, ThreadingHTTPServer
import json
from pathlib import Path
from tempfile import TemporaryDirectory
from threading import Lock, Thread
from time import monotonic, sleep
from types import SimpleNamespace
import unittest

from scripts import bulk_query


class TestBulkLoadSafety(unittest.TestCase):
    def serve(self, fail_first=False):
        state = {"started": 0, "active": 0, "max_active": 0, "timestamps": []}
        lock = Lock()

        class Handler(BaseHTTPRequestHandler):
            def do_POST(self):
                request = json.loads(self.rfile.read(int(self.headers["Content-Length"])))
                with lock:
                    state["started"] += 1
                    n = state["started"]
                    state["active"] += 1
                    state["max_active"] = max(state["max_active"], state["active"])
                    state["timestamps"].append(monotonic())
                # Fatal answer is made faster than the expected unresolved
                # responses so the breaker should trip before new scheduling.
                if not (fail_first and n == 1):
                    sleep(0.09 if fail_first else 0.04)
                status = 503 if fail_first and n == 1 else 422
                body = {"error": {"code":
                        "geocoder_unavailable" if status == 503 else "outside_supported_coverage"}}
                raw = json.dumps(body).encode()
                self.send_response(status)
                self.send_header("Content-Type", "application/json")
                self.send_header("Content-Length", str(len(raw)))
                self.end_headers()
                try:
                    self.wfile.write(raw)
                finally:
                    with lock:
                        state["active"] -= 1
            def log_message(self, *_):
                pass

        server = ThreadingHTTPServer(("127.0.0.1", 0), Handler)
        thread = Thread(target=server.serve_forever, daemon=True)
        thread.start()
        return server, thread, state

    def args(self, root, port, rows, max_requests=100, abort_after=0):
        source = root / "cases.jsonl"
        source.write_text("".join(json.dumps({
            "id": "case:" + str(i),
            "point": {"latitude": 39.75, "longitude": -104.99},
        }) + "\n" for i in range(rows)))
        return SimpleNamespace(
            input=str(source), results=str(root / "results.jsonl"),
            summary=str(root / "summary.json"),
            endpoint=f"http://127.0.0.1:{port}" + bulk_query.PATH,
            date="2026-10-29", concurrency=4, qps=100, max_requests=max_requests,
            timeout=2.0, project_type="", token_env="", allow_remote=False,
            shard_index=0, shard_count=1, expected_bundle="", expected_source="",
            abort_after_errors=abort_after,
        )

    def test_bounded_concurrent_local_requests_and_expected_422(self):
        server, thread, state = self.serve()
        try:
            with TemporaryDirectory() as directory:
                root = Path(directory)
                args = self.args(root, server.server_port, 18)
                report = bulk_query.run(args)
                self.assertEqual(report["samples"], 18)
                self.assertEqual(report["categories"], {"unresolved": 18})
                self.assertEqual(state["started"], 18)
                self.assertGreaterEqual(state["max_active"], 2)
                self.assertLessEqual(state["max_active"], args.concurrency)
                self.assertEqual(report["input_sha256"],
                                 sha256(Path(args.input).read_bytes()).hexdigest())
                self.assertEqual(len(Path(args.results).read_text().splitlines()), 18)
                self.assertNotIn("circuit_breaker_tripped", report)
        finally:
            server.shutdown()
            server.server_close()
            thread.join(timeout=2)

    def test_fatal_result_stops_new_scheduling_and_preserves_full_source_hash(self):
        server, thread, state = self.serve(fail_first=True)
        try:
            with TemporaryDirectory() as directory:
                root = Path(directory)
                args = self.args(root, server.server_port, 50, abort_after=1)
                report = bulk_query.run(args)
                self.assertTrue(report["circuit_breaker_tripped"])
                self.assertFalse(report["input_validation_complete"])
                self.assertEqual(report["abort_after_errors"], 1)
                self.assertGreaterEqual(report["fatal_error_count"], 1)
                self.assertGreaterEqual(report["categories"]["service_error"], 1)
                self.assertEqual(report["input_sha256"],
                                 sha256(Path(args.input).read_bytes()).hexdigest())
                self.assertLessEqual(state["started"], args.concurrency)
                self.assertEqual(report["samples"], state["started"])
                self.assertEqual(len(Path(args.results).read_text().splitlines()),
                                 state["started"])
        finally:
            server.shutdown()
            server.server_close()
            thread.join(timeout=2)

    def test_bounded_inputs_and_global_rate_limiter(self):
        for value in (-1, 100001, 1.5, True):
            with TemporaryDirectory() as folder:
                args = SimpleNamespace()
                with self.subTest(value=value):
                    server, thread, _ = self.serve()
                    try:
                        args = self.args(Path(folder), server.server_port, 1,
                                         abort_after=value)
                        with self.assertRaisesRegex(ValueError, "abort-after-errors"):
                            bulk_query.run(args)
                    finally:
                        server.shutdown()
                        server.server_close()
                        thread.join(timeout=2)
        limiter = bulk_query.RateLimiter(10)
        start = monotonic()
        for _ in range(3):
            limiter.acquire()
        self.assertGreaterEqual(monotonic() - start, 0.18)


if __name__ == "__main__":
    unittest.main()
