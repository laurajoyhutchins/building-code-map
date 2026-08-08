from __future__ import annotations

import argparse
from pathlib import Path

import yaml

from schemas import validate_document
from yaml_io import load_yaml_file


class ArtifactValidationError(ValueError):
    """Raised when a generated YAML artifact violates its declared schema."""


def validate_artifact_file(path: Path) -> None:
    try:
        document = load_yaml_file(path)
        validate_document(document)
    except (OSError, yaml.YAMLError, ValueError) as exc:
        raise ArtifactValidationError(f"{path}: {exc}") from exc


def main() -> int:
    parser = argparse.ArgumentParser(
        description="Validate generated report YAML without modifying it."
    )
    parser.add_argument(
        "artifact_dir",
        nargs="?",
        default=str(Path(__file__).resolve().parent.parent / "yaml"),
        help="Generated artifact directory. Defaults to reports/yaml.",
    )
    args = parser.parse_args()

    artifact_dir = Path(args.artifact_dir)
    paths = sorted(artifact_dir.rglob("*.yaml"))
    failures = 0
    for path in paths:
        try:
            validate_artifact_file(path)
            print(f"OK   {path}")
        except ArtifactValidationError as exc:
            print(f"ERR  {exc}")
            failures += 1

    if not paths:
        print(f"ERR  No YAML artifacts found under {artifact_dir}.")
        return 1
    print(f"Done: {len(paths) - failures} valid, {failures} errors")
    return 1 if failures else 0


if __name__ == "__main__":
    raise SystemExit(main())
