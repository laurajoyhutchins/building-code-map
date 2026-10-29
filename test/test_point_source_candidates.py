"""Offline source point-in-polygon tests; geometry never admits legal authority."""
from __future__ import annotations

from hashlib import sha256
import json
from pathlib import Path
from tempfile import TemporaryDirectory
import unittest

from shapely.geometry import Polygon, box, mapping
from scripts.point_source_candidates import query
from scripts.generate_boundary_pairs import load_features


class SourcePolygonCandidates(unittest.TestCase):
    def fixture(self, root, geometries):
        file = root / "publisher.geojson"
        file.write_text(json.dumps({"type": "FeatureCollection",
            "features": [{"type": "Feature",
                           "properties": {"OBJECTID": ident, "NAME": name},
                           "geometry": mapping(polygon)}
                          for ident, name, polygon in geometries]},
                          allow_nan=False, sort_keys=True))
        return file, sha256(file.read_bytes()).hexdigest()

    def look(self, file, digest, lon, lat, n=2):
        return query(file, digest, longitude=lon, latitude=lat,
                     role="municipal_boundary", expect_features=n)

    def test_one_candidate_and_absent_source_never_authority(self):
        with TemporaryDirectory() as td:
            source, sha = self.fixture(Path(td), [
                (5, "FIRST", box(-106,39,-105,40)),
                (8, "SECOND", box(-104,39,-103,40))])
            answer = self.look(source, sha, -105.5, 39.5)
            self.assertEqual(answer["candidate_feature_ids"], ["5"])
            self.assertEqual(answer["status"], "one_source_candidate")
            self.assertFalse(answer["regulatory_verdict_authorized"])
            self.assertEqual(answer["legal_authority"], "not_established")
            self.assertEqual(answer["identity_review"], "not_established")
            outside = self.look(source, sha, -102.5, 39.5)
            self.assertEqual(outside["status"], "no_source_polygon_match")
            self.assertEqual(outside["candidate_feature_ids"], [])
            self.assertEqual(outside["legal_authority"], "not_established")

    def test_duplicate_names_remain_separate_object_ids(self):
        with TemporaryDirectory() as td:
            source, sha = self.fixture(Path(td), [
                (11, "FORT COLLINS", box(-106,39,-105,40)),
                (12, "FORT COLLINS", box(-105.9,39.1,-104.9,40.1))])
            both = self.look(source,sha,-105.5,39.5)
            self.assertEqual(both["candidate_feature_ids"], ["11","12"])
            self.assertEqual(both["status"], "multiple_source_candidates")
            self.assertFalse(both["regulatory_verdict_authorized"])

    def test_polygon_hole_and_boundary_are_unknown(self):
        with TemporaryDirectory() as td:
            shell = Polygon([(-106,39),(-103,39),(-103,42),(-106,42)],
                            holes=[[(-105,40),(-104,40),(-104,41),(-105,41)]])
            source, sha = self.fixture(Path(td), [
                (10,"SHELL",shell),(20,"REMOTE",box(-102,39,-101,40))])
            hole = self.look(source,sha,-104.5,40.5)
            self.assertEqual(hole["candidate_feature_ids"], [])
            self.assertEqual(hole["status"], "no_source_polygon_match")
            edge = self.look(source,sha,-106,40.5)
            self.assertEqual(edge["status"], "boundary_uncertain")
            self.assertEqual(edge["boundary_feature_ids"], ["10"])

    def test_digest_count_and_malformed_object_id_rejected(self):
        with TemporaryDirectory() as td:
            source, sha = self.fixture(Path(td), [
                (1,"CITY",box(-106,39,-105,40)),(2,"CITY",box(-104,39,-103,40))])
            with self.assertRaisesRegex(ValueError,"SHA-256"):
                self.look(source,"0"*64,-105.5,39.5)
            with self.assertRaisesRegex(ValueError,"feature count"):
                self.look(source,sha,-105.5,39.5,n=3)
            doc=json.loads(source.read_text())
            doc["features"][1]["properties"]["OBJECTID"]=1
            source.write_text(json.dumps(doc))
            with self.assertRaisesRegex(ValueError,"unique"):
                self.look(source,sha256(source.read_bytes()).hexdigest(),-105.5,39.5)
            doc["features"][1]["properties"]["OBJECTID"]=True
            source.write_text(json.dumps(doc))
            with self.assertRaisesRegex(ValueError,"unique"):
                self.look(source,sha256(source.read_bytes()).hexdigest(),-105.5,39.5)

    def test_invalid_layer_role_and_coordinate_refused(self):
        with TemporaryDirectory() as td:
            source, sha = self.fixture(Path(td), [
                (1,"ONE",box(-106,39,-105,40)),(2,"TWO",box(-104,39,-103,40))])
            for lon,lat in [(181,0),(0,91),(float('nan'),0),(True,0)]:
                with self.subTest(point=(lon,lat)),self.assertRaisesRegex(ValueError,"WGS84"):
                    self.look(source,sha,lon,lat)
            with self.assertRaisesRegex(ValueError,"role"):
                query(source,sha,longitude=-105.5,latitude=39.5,
                      role="verified_permit_AHJ",expect_features=2)

    def test_closed_schema_and_odcs_reject_promoted_verdict(self):
        from jsonschema import Draft202012Validator,FormatChecker
        root = Path(__file__).resolve().parents[1]
        schema=json.loads((root/"schemas/source-point-candidates-v1.schema.json").read_text())
        odcs=json.loads((root/"schemas/source-point-candidates-v1.odcs.json").read_text())
        Draft202012Validator.check_schema(schema)
        self.assertEqual(odcs["apiVersion"],"v3.1.0")
        self.assertEqual([x["name"] for x in odcs["schema"][0]["properties"]],list(schema["properties"]))
        with TemporaryDirectory() as td:
            source, sha = self.fixture(Path(td), [
                (1,"ONE",box(-106,39,-105,40)),(2,"TWO",box(-104,39,-103,40))])
            row=self.look(source,sha,-105.5,39.5)
            validator=Draft202012Validator(schema,format_checker=FormatChecker())
            validator.validate(row)
            row["legal_authority"]="reviewed"
            with self.assertRaises(Exception):
                validator.validate(row)
            row["legal_authority"]="not_established"
            row["adopted_code"]="IBC"
            with self.assertRaises(Exception):
                validator.validate(row)


if __name__=="__main__":
    unittest.main()
