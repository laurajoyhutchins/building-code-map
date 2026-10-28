"""Negative tests for the source-bound Lean axiom report verifier."""
import importlib.util
import pathlib
import unittest

SCRIPT = pathlib.Path(__file__).with_name("verify_eth_lean_artifact.py")
spec = importlib.util.spec_from_file_location("eth_verifier", SCRIPT)
verifier = importlib.util.module_from_spec(spec)
spec.loader.exec_module(verifier)


class AxiomReportTests(unittest.TestCase):
    def reports(self, axioms="propext,\n Classical.choice,\n Quot.sound"):
        return "\n".join(
            f"'{verifier.PREFIX}{name}' depends on axioms: [{axioms}]"
            for name in sorted(verifier.EXPECTED)
        ) + "\n"

    def test_actual_multiline_lean_format(self):
        observed = verifier.parse_axioms(self.reports())
        self.assertEqual(set(observed), verifier.EXPECTED)
        for axioms in observed.values():
            self.assertEqual(set(axioms), verifier.PERMITTED)

    def test_single_line_format(self):
        self.assertEqual(len(verifier.parse_axioms(
            self.reports("propext, Classical.choice, Quot.sound"))), 5)

    def test_disallowed_axiom_rejected(self):
        with self.assertRaisesRegex(ValueError, "Disallowed"):
            verifier.parse_axioms(self.reports("propext, Classical.choice, sorryAx"))

    def test_missing_report_rejected(self):
        with self.assertRaisesRegex(ValueError, "Missing"):
            verifier.parse_axioms(self.reports().splitlines()[0] + "\n")

    def test_duplicate_report_rejected(self):
        data = self.reports()
        with self.assertRaisesRegex(ValueError, "duplicate"):
            verifier.parse_axioms(data + data.split("]\n")[0] + "]\n")

    def test_malformed_report_rejected(self):
        data = self.reports().replace("depends on axioms:", "depends on axioms", 1)
        with self.assertRaisesRegex(ValueError, "Malformed"):
            verifier.parse_axioms(data)

    def test_lean_error_rejected(self):
        with self.assertRaisesRegex(ValueError, "Lean error"):
            verifier.parse_axioms(self.reports() + "error: failure\n")

    def test_sorry_ax_elsewhere_rejected(self):
        with self.assertRaisesRegex(ValueError, "sorryAx"):
            verifier.parse_axioms(self.reports() + "warning: sorryAx detected\n")

    def test_duplicate_status_key_rejected(self):
        with self.assertRaisesRegex(ValueError, "Duplicate"):
            verifier.key_values("lake-build=0\nlake-build=124\n")


if __name__ == "__main__":
    unittest.main()
