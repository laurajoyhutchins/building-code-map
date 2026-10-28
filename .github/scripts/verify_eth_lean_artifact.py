#!/usr/bin/env python3
"""Fail-closed replay of the immutable exact-source ETH Lean proof evidence."""
import argparse
import hashlib
import json
import pathlib
import re

SOURCE = "7ef5c0cc5ca4b0e072b7a050fca09ab24bc6fa3a"
MANIFEST = "cf6105a25d9dca2f166b241d9191bd12c7e13305890dc9c4d0952351cccc0794"
LEAN_COMMIT = "5045d0056413266e57c625dcd7c365b10e377c52"
PREFIX = "OAI.Algorithms.SubexponentialSAT.SqrtLogSpace."
EXPECTED = {
    "compiled_exact_count_and_sat", "exact_count_and_sat",
    "exact_count_eth", "sat_polynomial_space", "original_sat_main",
}
PERMITTED = {"propext", "Quot.sound", "Classical.choice"}
REPORT = re.compile(
    r"^'(?P<name>OAI\.Algorithms\.SubexponentialSAT\.SqrtLogSpace\.[A-Za-z0-9_]+)' "
    r"(?:(?:depends on axioms: \[(?P<axioms>[^\]]*)\])|does not depend on any axioms)",
    re.MULTILINE,
)


def parse_axioms(content: str) -> dict[str, list[str]]:
    """Lean prints multi-line axiom lists, not one report per text line."""
    observed: dict[str, list[str]] = {}
    matches = list(REPORT.finditer(content))
    report_starts = sum(line.startswith("'" + PREFIX) for line in content.splitlines())
    if len(matches) != report_starts:
        raise ValueError("Malformed or unparsed Lean axiom report")
    for match in matches:
        full = match.group("name")
        short = full[len(PREFIX):]
        if short not in EXPECTED or short in observed:
            raise ValueError("Unexpected or duplicate axiom report: " + full)
        axioms = [a.strip() for a in (match.group("axioms") or "").split(",") if a.strip()]
        if len(axioms) != len(set(axioms)):
            raise ValueError("Duplicate axiom in report: " + full)
        observed[short] = sorted(axioms)
    missing = EXPECTED - observed.keys()
    if missing:
        raise ValueError("Missing axiom reports: " + ", ".join(sorted(missing)))
    disallowed = {name: sorted(set(axioms) - PERMITTED)
                  for name, axioms in observed.items() if set(axioms) - PERMITTED}
    if disallowed:
        raise ValueError("Disallowed axioms: " + json.dumps(disallowed, sort_keys=True))
    if "sorryAx" in content:
        raise ValueError("sorryAx appears in final Lean output")
    if re.search(r"(?m)^error:", content):
        raise ValueError("Lean error appears in final output")
    return observed


def key_values(text: str) -> dict[str, str]:
    lines = [line for line in text.splitlines() if line.strip()]
    if any(line.count("=") != 1 for line in lines):
        raise ValueError("Malformed identity/status line")
    values = dict(line.split("=", 1) for line in lines)
    if len(values) != len(lines):
        raise ValueError("Duplicate identity/status key")
    return values


def verify(directory: pathlib.Path) -> dict:
    identity = key_values((directory / "identity.txt").read_text())
    status = key_values((directory / "status.txt").read_text())
    expected_identity = {
        "source": SOURCE,
        "compiler": (
            "Lean (version 4.34.1, x86_64-unknown-linux-gnu, "
            "commit " + LEAN_COMMIT + ", Release)"
        ),
        "manifest_sha256": MANIFEST,
        "stage": "03",
    }
    if identity != expected_identity:
        raise ValueError("Source/compiler/manifest/stage identity mismatch")
    if status != {"lake-build": "0", "lean-final": "0"}:
        raise ValueError("Compilation or final elaboration did not succeed")
    build = (directory / "lake-build.log").read_text(errors="replace")
    if "Build completed successfully (11492 jobs)." not in build:
        raise ValueError("Missing successful exact-source build completion")
    final = (directory / "lean-final.log").read_text(errors="replace")
    observed = parse_axioms(final)
    return {
        "accepted": True,
        "source": SOURCE,
        "compiler": identity["compiler"],
        "manifest_sha256": MANIFEST,
        "stage": "03",
        "lake_build": 0,
        "lean_final": 0,
        "observed": observed,
        "missing": [],
        "disallowed": {},
        "sorry_ax": False,
        "lean_final_sha256": hashlib.sha256(final.encode()).hexdigest(),
        "lake_build_sha256": hashlib.sha256(build.encode()).hexdigest(),
    }


def main() -> None:
    parser = argparse.ArgumentParser()
    parser.add_argument("evidence_dir", type=pathlib.Path)
    parser.add_argument("--output", type=pathlib.Path)
    args = parser.parse_args()
    result = verify(args.evidence_dir)
    rendered = json.dumps(result, indent=2, sort_keys=True) + "\n"
    if args.output:
        args.output.write_text(rendered)
    print(rendered, end="")


if __name__ == "__main__":
    main()
