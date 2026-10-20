import csv
import json
from pathlib import Path
import tempfile
import unittest

from scripts.ingest_colorado_addresses import export, parse_row


def row(index, **overrides):
    attrs = {
        "OBJECTID": index, "AddrNum": "1600", "StreetName": "Broadway",
        "PreDir": "N", "PlaceName": "Denver", "Zipcode": "80202",
        "Latitude": 39.7411, "Longitude": -104.9876,
    }
    attrs.update(overrides)
    return {"attributes": attrs}


class ColoradoAddressExportTest(unittest.TestCase):
    def test_reviewed_row_shape(self):
        result = parse_row(row(1))
        self.assertEqual(result["source_record_id"], "co-oit-1")
        self.assertEqual(result["street"], "N Broadway")
        self.assertEqual(result["city"], "Denver")
        self.assertEqual(result["state"], "CO")

    def test_paged_export_and_attribution(self):
        requests = []
        def fake_fetch(params):
            requests.append(params)
            if params.get("returnCountOnly") == "true":
                return {"count": 2}
            return {"features": [row(1), row(2, AddrNum="1601")]}
        with tempfile.TemporaryDirectory() as temp:
            path = Path(temp) / "points.csv"
            receipt = export(fake_fetch, path, "2026-10-20")
            with path.open(newline="", encoding="utf-8") as handle:
                data = list(csv.DictReader(handle))
            self.assertEqual(len(data), 2)
            self.assertEqual(receipt["accepted"], 2)
            self.assertEqual(receipt["invalid"], 0)
            self.assertFalse(receipt["production_approved"])
            self.assertIn("State of Colorado Public addresses", receipt["attribution"])
            self.assertEqual(requests[1]["orderByFields"], "OBJECTID ASC")
            saved = json.loads(path.with_suffix(".csv.receipt.json").read_text())
            self.assertEqual(receipt, saved)

    def test_rejects_missing_city_without_partial_publication(self):
        def fake_fetch(params):
            if params.get("returnCountOnly") == "true":
                return {"count": 1}
            return {"features": [row(1, PlaceName=None)]}
        with tempfile.TemporaryDirectory() as temp:
            path = Path(temp) / "points.csv"
            with self.assertRaisesRegex(ValueError, "invalid records"):
                export(fake_fetch, path, "2026-10-20")
            self.assertFalse(path.exists())

    def test_rejects_duplicate_and_wrong_order(self):
        def fake_fetch(params):
            if params.get("returnCountOnly") == "true":
                return {"count": 2}
            return {"features": [row(2), row(1)]}
        with tempfile.TemporaryDirectory() as temp:
            with self.assertRaisesRegex(ValueError, "nonmonotonic"):
                export(fake_fetch, Path(temp)/"points.csv", "2026-10-20")

    def test_rejects_non_colorado_coordinates(self):
        with self.assertRaisesRegex(ValueError, "outside Colorado"):
            parse_row(row(1, Longitude=-74.0))

    def test_rejects_incomplete_page(self):
        def fake_fetch(params):
            return {"count": 2} if "returnCountOnly" in params else {"features": [row(1)]}
        with tempfile.TemporaryDirectory() as temp:
            with self.assertRaisesRegex(ValueError, "pagination"):
                export(fake_fetch, Path(temp)/"points.csv", "2026-10-20")


if __name__ == "__main__":
    unittest.main()
