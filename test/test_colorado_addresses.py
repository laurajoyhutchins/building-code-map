import csv
import json
from pathlib import Path
import tempfile
import unittest
from unittest.mock import patch

from scripts import ingest_colorado_addresses as importer
from scripts.ingest_colorado_addresses import export, parse_row, verify_export


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

    @staticmethod
    def one_record_fetch(object_id):
        def fake(params):
            if params.get("returnCountOnly") == "true":
                return {"count": 1}
            return {"features": [row(object_id)]}
        return fake

    def test_verified_export_requires_matching_csv_bytes_and_receipt(self):
        with tempfile.TemporaryDirectory() as temp:
            path = Path(temp) / "points.csv"
            receipt = export(self.one_record_fetch(1), path, "2026-10-20")
            self.assertEqual(verify_export(path), receipt)
            path.write_text("unverified replacement", encoding="utf-8")
            with self.assertRaisesRegex(ValueError, "differs from"):
                verify_export(path)

    def test_successful_replacement_keeps_consistent_pair(self):
        with tempfile.TemporaryDirectory() as temp:
            path = Path(temp) / "points.csv"
            first = export(self.one_record_fetch(1), path, "2026-10-20")
            second = export(self.one_record_fetch(2), path, "2026-10-21")
            self.assertNotEqual(first["csv_sha256"], second["csv_sha256"])
            self.assertEqual(verify_export(path), second)
            self.assertIn("co-oit-2", path.read_text(encoding="utf-8"))
            self.assertEqual(list(path.parent.glob("*.partial")), [])
            self.assertEqual(list(path.parent.glob(".points.csv.recovery.*")), [])

    def test_receipt_serialization_failure_preserves_prior_verified_pair(self):
        with tempfile.TemporaryDirectory() as temp:
            path = Path(temp) / "points.csv"
            previous = export(self.one_record_fetch(1), path, "2026-10-20")
            with patch.object(importer.json, "dumps", side_effect=OSError("receipt serialization")):
                with self.assertRaisesRegex(OSError, "receipt serialization"):
                    export(self.one_record_fetch(2), path, "2026-10-21")
            self.assertEqual(verify_export(path), previous)

    def test_existing_pair_restored_if_receipt_rename_fails(self):
        with tempfile.TemporaryDirectory() as temp:
            path = Path(temp) / "points.csv"
            previous = export(self.one_record_fetch(1), path, "2026-10-20")
            receipt_path = importer.receipt_path(path)
            original_replace = importer.os.replace

            def reject_receipt(src, dst):
                if Path(dst) == receipt_path and str(src).endswith(".partial"):
                    raise OSError("sidecar rename rejected")
                return original_replace(src, dst)

            with patch.object(importer.os, "replace", side_effect=reject_receipt):
                with self.assertRaisesRegex(OSError, "sidecar rename rejected"):
                    export(self.one_record_fetch(2), path, "2026-10-21")
            self.assertEqual(verify_export(path), previous)
            self.assertNotIn("co-oit-2", path.read_text(encoding="utf-8"))
            self.assertEqual(list(path.parent.glob(".points.csv.recovery.*")), [])

    def test_ambiguous_receipt_rename_rolls_back_from_actual_files(self):
        with tempfile.TemporaryDirectory() as temp:
            path = Path(temp) / "points.csv"
            previous = export(self.one_record_fetch(1), path, "2026-10-20")
            sidecar = importer.receipt_path(path)
            original_replace = importer.os.replace

            def error_after_rename(src, dst):
                result = original_replace(src, dst)
                if Path(dst) == sidecar and str(src).endswith(".partial"):
                    raise OSError("lost response after sidecar rename")
                return result

            with patch.object(importer.os, "replace", side_effect=error_after_rename):
                with self.assertRaisesRegex(OSError, "lost response after sidecar rename"):
                    export(self.one_record_fetch(2), path, "2026-10-21")
            self.assertEqual(verify_export(path), previous)

    def test_no_previous_pair_is_removed_after_receipt_publish_failure(self):
        with tempfile.TemporaryDirectory() as temp:
            path = Path(temp) / "points.csv"
            sidecar = importer.receipt_path(path)
            original_replace = importer.os.replace

            def reject(src, dst):
                if Path(dst) == sidecar and str(src).endswith(".partial"):
                    raise OSError("sidecar publication failed")
                return original_replace(src, dst)

            with patch.object(importer.os, "replace", side_effect=reject):
                with self.assertRaisesRegex(OSError, "sidecar publication failed"):
                    export(self.one_record_fetch(1), path, "2026-10-20")
            self.assertFalse(path.exists())
            self.assertFalse(sidecar.exists())

    def test_missing_old_receipt_blocks_replacement_without_source_loss(self):
        with tempfile.TemporaryDirectory() as temp:
            path = Path(temp) / "points.csv"
            path.write_text("unreceipted historical address data", encoding="utf-8")
            with self.assertRaisesRegex(ValueError, "incomplete prior export pair"):
                export(self.one_record_fetch(1), path, "2026-10-20")
            self.assertEqual(path.read_text(encoding="utf-8"),
                             "unreceipted historical address data")


    def test_rejects_incomplete_page(self):
        def fake_fetch(params):
            return {"count": 2} if "returnCountOnly" in params else {"features": [row(1)]}
        with tempfile.TemporaryDirectory() as temp:
            with self.assertRaisesRegex(ValueError, "pagination"):
                export(fake_fetch, Path(temp)/"points.csv", "2026-10-20")


if __name__ == "__main__":
    unittest.main()
