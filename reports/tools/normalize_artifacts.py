from __future__ import annotations

import argparse
from pathlib import Path

import yaml

from schemas import canonicalize_document
from yaml_io import dump_yaml, load_yaml_file


class ArtifactNormalizationError(ValueError):
    """Raised when a legacy artifact cannot be normalized without losing data."""


def normalize_artifact_file(path: Path) -> str:
    try:
        document = load_yaml_file(path)
        canonical = canonicalize_document(document, str(path))
        return f"---\n{dump_yaml(canonical)}"
    except (OSError, yaml.YAMLError, ValueError) as exc:
        raise ArtifactNormalizationError(f"{path}: {exc}") from exc


def main() -> int:
    parser = argparse.ArgumentParser(
        description="Normalize generated report YAML into the declared canonical schemas."
    )
    parser.add_argument(
        "artifact_dir",
        nargs="?",
        default=str(Path(__file__).resolve().parent.parent / "yaml"),
    )
    mode = parser.add_mutually_exclusive_group(required=True)
    mode.add_argument(
        "--check",
        action="store_true",
        help="Return non-zero if any artifact is not already canonical.",
    )
    mode.add_argument(
        "--write",
        action="store_true",
        help="Rewrite artifacts in place using canonical schemas.",
    )
    mode.add_argument(
        "--output-dir",
        help="Write canonical artifacts under a separate directory.",
    )
    args = parser.parse_args()

    artifact_dir = Path(args.artifact_dir)
    output_dir = Path(args.output_dir) if args.output_dir else None
    paths = sorted(artifact_dir.rglob("*.yaml"))
    failures = 0
    changed = 0
    unchanged = 0

    for path in paths:
        try:
            canonical_text = normalize_artifact_file(path)
            current_text = path.read_text(encoding="utf-8")
            is_changed = canonical_text != current_text
            if args.check:
                if is_changed:
                    print(f"ERR  {path}: artifact is not canonical")
                    failures += 1
                else:
                    print(f"OK   {path}")
                    unchanged += 1
                continue

            target = path if args.write else output_dir / path.relative_to(artifact_dir)
            target.parent.mkdir(parents=True, exist_ok=True)
            target.write_text(canonical_text, encoding="utf-8")
            if is_changed:
                print(f"WRITE {target}")
                changed += 1
            else:
                print(f"COPY  {target}")
                unchanged += 1
        except ArtifactNormalizationError as exc:
            print(f"ERR  {exc}")
            failures += 1

    if not paths:
        print(f"ERR  No YAML artifacts found under {artifact_dir}.")
        return 1
    print(f"Done: {changed} changed, {unchanged} unchanged, {failures} errors")
    return 1 if failures else 0


if __name__ == "__main__":
    raise SystemExit(main())
