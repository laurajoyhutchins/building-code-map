"""Independent RDF/JSON-LD expansion acceptance for BCM's actual Go exporter.

Only synthetic inputs are permitted. This verifies real PROV-O triples, not
whether an ordinance, source snapshot, activity, or legal decision is authentic.
"""
import json
import subprocess
import unittest
from pathlib import Path

from rdflib import Graph, Literal, Namespace, URIRef
from rdflib.namespace import RDF, XSD

ROOT = Path(__file__).resolve().parents[1]
PROV = Namespace("http://www.w3.org/ns/prov#")
BCM = Namespace("https://github.com/laurajoyhutchins/building-code-map/ns/provenance#")


def parsed(value):
    return Graph().parse(data=json.dumps(value), format="json-ld")


class PROVRDFTests(unittest.TestCase):
    def test_real_go_export_expands_to_w3c_prov_triples(self):
        fixture = (ROOT / "test/fixtures/prov/claim.json").read_bytes()
        result = subprocess.run(
            ["go", "run", "./cmd/bcm-prov"],
            input=fixture, capture_output=True, cwd=ROOT, check=True, timeout=60,
        )
        exported = json.loads(result.stdout)
        graph = parsed(exported)
        entities = {s for s, _, _ in graph.triples((None, RDF.type, PROV.Entity))}
        self.assertEqual(len(entities), 5)
        self.assertEqual(list(graph.triples((None, RDF.type, PROV.Activity))), [])
        self.assertEqual(list(graph.triples((None, RDF.type, PROV.Agent))), [])
        self.assertTrue(any(graph.triples((None, PROV.wasDerivedFrom, None))))
        self.assertTrue(any(graph.triples((None, BCM.hasEvidenceLink, None))))
        self.assertTrue(any(graph.triples((None, BCM.legallyMaterial, Literal(True)))))
        self.assertTrue(any(graph.triples((None, BCM.asOf, Literal("2026-10-28", datatype=XSD.date)))))
        self.assertTrue(any(graph.triples((None, BCM.sourceID, Literal("DEMO-ARTIFACT-1")))))

    def test_process_graph_expands_agent_activity_and_datetimes(self):
        fixture = json.loads((ROOT / "test/fixtures/prov/actor.jsonld").read_text())
        validated = subprocess.run(
            ["go", "run", "./cmd/bcm-prov", "validate"],
            input=json.dumps(fixture).encode(), capture_output=True,
            cwd=ROOT, check=True, timeout=60,
        )
        graph = parsed(json.loads(validated.stdout))
        doc = URIRef("urn:bcm:prov:demo:document")
        activity = URIRef("urn:bcm:prov:demo:activity")
        agent = URIRef("urn:bcm:prov:demo:agent")
        output = URIRef("urn:bcm:prov:demo:output")
        self.assertIn((activity, RDF.type, PROV.Activity), graph)
        self.assertIn((agent, RDF.type, PROV.Agent), graph)
        self.assertIn((activity, PROV.used, doc), graph)
        self.assertIn((activity, PROV.wasAssociatedWith, agent), graph)
        self.assertIn((output, PROV.wasGeneratedBy, activity), graph)
        self.assertIn((output, PROV.wasDerivedFrom, doc), graph)
        self.assertIn(
            (activity, PROV.startedAtTime, Literal("2026-10-28T12:00:00Z", datatype=XSD.dateTime)),
            graph,
        )


if __name__ == "__main__":
    unittest.main()
