"""Cross-language, source-read-only verification for the public BCM product."""

from __future__ import annotations

import argparse
import subprocess
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]


def run(label: str, argv: list[str]) -> bool:
    print(f"==> {label}", flush=True)
    result = subprocess.run(argv, cwd=ROOT, check=False)
    return result.returncode == 0


def main() -> int:
    parser = argparse.ArgumentParser(description="Check Go, Python and Node using repository tooling.")
    parser.add_argument("--preflight", action="store_true", help="Run static checks only")
    parser.add_argument("--fix-style", action="store_true", help="Apply gofmt before preflight")
    args = parser.parse_args()
    if args.fix_style and not args.preflight:
        parser.error("--fix-style requires --preflight")

    tracked = subprocess.run(
        ["git", "ls-files", "-z", "--", "*.go"], cwd=ROOT, capture_output=True, check=True
    ).stdout.decode().split(chr(0))
    go_files = [path for path in tracked if path]
    if not go_files:
        raise SystemExit("No tracked Go files; cannot claim formatting passed")
    if args.fix_style:
        if not run("Apply Go formatting", ["gofmt", "-w", *go_files]):
            return 1

    failures: list[str] = []
    result = subprocess.run(["gofmt", "-l", *go_files], cwd=ROOT, capture_output=True, text=True)
    if result.returncode or result.stdout.strip():
        print("Unformatted Go source:\n" + result.stdout + result.stderr, file=sys.stderr)
        subprocess.run(["gofmt", "-d", *go_files], cwd=ROOT, check=False)
        failures.append("Go formatting")
    checks = [
        ("Go vet", ["go", "vet", "./..."]),
        ("Python syntax", [sys.executable, "-m", "compileall", "-q", "scripts", "test"]),
    ]
    checks.extend(
        (f"Node syntax: {path.name}", ["node", "--check", str(path.relative_to(ROOT))])
        for path in sorted((ROOT / "scripts").glob("*.mjs"))
    )
    checks.extend(
        (f"Node TypeScript module: {path.name}", [
            "node", "--experimental-strip-types", "--input-type=module", "--eval",
            "await import('./" + path.relative_to(ROOT).as_posix() + "')",
        ])
        for path in sorted((ROOT / "scripts").glob("*.ts"))
    )
    if not args.preflight:
        checks += [
            ("Go tests", ["go", "test", "./..."]),
            ("Python tests", [sys.executable, "-m", "unittest", "discover", "-s", "test", "-p", "test_*.py"]),
            ("Node tests", ["npm", "test"]),
            ("Public boundary and contracts", ["npm", "run", "validate"]),
            ("Static pages build", ["npm", "run", "build:pages"]),
        ]
    for label, argv in checks:
        if not run(label, argv):
            failures.append(label)
    if failures:
        print("Failed checks: " + ", ".join(failures), file=sys.stderr)
        return 1
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
