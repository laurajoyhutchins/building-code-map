#!/usr/bin/env python3
"""Regenerate the self-contained Building Code Authority Engine package."""

from __future__ import annotations

import argparse
import hashlib
import json
import shutil
import subprocess
from datetime import datetime, timezone
from pathlib import Path
from typing import Any

ROOT = Path(__file__).resolve().parents[1]
PACKAGE = ROOT / "building-code-engine"
MANIFESTS = PACKAGE / "manifests"
MANIFEST_DATA = MANIFESTS / "data"
BUNDLE_PATH = MANIFESTS / "bundle.json"
VERIFICATION_PATH = MANIFESTS / "verification.json"
CHECKSUMS_PATH = MANIFESTS / "checksums.json"


def source_created_at(source_commit: str) -> str:
    raw = subprocess.check_output(
        ["git", "show", "-s", "--format=%cI", source_commit],
        cwd=ROOT,
        text=True,
    ).strip()
    parsed = datetime.fromisoformat(raw.replace("Z", "+00:00"))
    return parsed.astimezone(timezone.utc).replace(microsecond=0).isoformat().replace("+00:00", "Z")


def sha256_file(path: Path) -> str:
    digest = hashlib.sha256()
    with path.open("rb") as handle:
        for chunk in iter(lambda: handle.read(1024 * 1024), b""):
            digest.update(chunk)
    return digest.hexdigest()


def sha256_directory(path: Path) -> str:
    digest = hashlib.sha256()
    files = sorted(item for item in path.rglob("*") if item.is_file())
    if not files:
        raise ValueError(f"recursive component directory is empty: {path}")
    for item in files:
        relative = item.relative_to(path).as_posix()
        digest.update(f"{relative}\0{sha256_file(item)}\n".encode())
    return digest.hexdigest()


def write_json(path: Path, value: Any) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(json.dumps(value, indent=2, sort_keys=False) + "\n", encoding="utf-8")


def sync_tree(source: Path, destination: Path) -> None:
    if destination.exists():
        shutil.rmtree(destination)
    shutil.copytree(source, destination)


def sync_engine_schemas() -> None:
    destination = PACKAGE / "schemas"
    if destination.exists():
        shutil.rmtree(destination)
    destination.mkdir(parents=True)
    sources = sorted((ROOT / "schemas").glob("engine-*.schema.json"))
    if not sources:
        raise ValueError("no engine JSON Schemas found")
    for source in sources:
        shutil.copy2(source, destination / source.name)


def sync_manifest_data() -> None:
    if MANIFEST_DATA.exists():
        shutil.rmtree(MANIFEST_DATA)
    MANIFEST_DATA.mkdir(parents=True)
    shutil.copy2(PACKAGE / "data/authority.snapshot.json", MANIFEST_DATA / "authority.snapshot.json")
    shutil.copy2(PACKAGE / "data/geocoder.sqlite", MANIFEST_DATA / "geocoder.sqlite")
    shutil.copy2(
        PACKAGE / "data/geocoder.sqlite.manifest.json",
        MANIFEST_DATA / "geocoder.sqlite.manifest.json",
    )
    sync_tree(PACKAGE / "data/regulatory", MANIFEST_DATA / "regulatory")


def coverage_state(state_id: str) -> str:
    prefix = "US-"
    if not state_id.startswith(prefix) or len(state_id) != len(prefix) + 2:
        raise ValueError(f"unsupported regulatory state_id: {state_id}")
    return state_id.removeprefix(prefix)


def coverage() -> dict[str, Any]:
    existing = json.loads((MANIFESTS / "coverage.json").read_text(encoding="utf-8"))
    profiles = sorted((ROOT / "backend/data/regulatory").glob("*.json"))
    states = sorted(
        coverage_state(json.loads(path.read_text(encoding="utf-8"))["state_id"])
        for path in profiles
    )
    if states != sorted(existing["states"]):
        raise ValueError(f"coverage states {existing['states']} do not match profiles {states}")
    return {
        "states": states,
        "code_families": sorted(existing["code_families"]),
        "as_of": existing["as_of"],
    }


def verification_receipt(
    *, source_commit: str, created_at: str, status: str, go_version: str, cold_room_mode: str | None, workflow_run: str | None
) -> dict[str, Any]:
    checks = [
        "go-test",
        "go-vet",
        "deterministic-binary-build",
        "geocoder-rebuild",
        "geocoder-byte-reproducibility",
        "geocoder-manifest-reproducibility",
        "bundle-inspect",
        "cli-resolve",
        "cli-geocode",
        "http-readiness",
        "http-resolve",
        "mcp-initialize",
        "offline-verification",
        "cold-room-verification",
    ]
    receipt: dict[str, Any] = {
        "schema_version": "1",
        "generated_at": created_at,
        "source_commit": source_commit,
        "build": {
            "go_version": go_version,
            "goos": "linux",
            "goarch": "amd64",
            "cgo_enabled": "0",
            "trimpath": True,
            "build_id_cleared": True,
            "geocoder_source_name": "building-code-map-release-fixtures",
            "geocoder_source_vintage": source_commit,
        },
        "binary": {
            "path": "bin/bcm",
            "sha256": sha256_file(PACKAGE / "bin/bcm"),
        },
        "geocoder": {
            "path": "data/geocoder.sqlite",
            "sha256": sha256_file(PACKAGE / "data/geocoder.sqlite"),
            "snapshot_manifest": {
                "path": "data/geocoder.sqlite.manifest.json",
                "sha256": sha256_file(PACKAGE / "data/geocoder.sqlite.manifest.json"),
            },
        },
        "checks": [{"name": name, "status": status} for name in checks],
    }
    if cold_room_mode is not None:
        receipt["cold_room_mode"] = cold_room_mode
    if workflow_run is not None:
        receipt["workflow_run"] = workflow_run
    return receipt


def checksum_paths() -> list[Path]:
    return sorted(
        (
            path
            for path in PACKAGE.rglob("*")
            if path.is_file() and path != CHECKSUMS_PATH
        ),
        key=lambda item: item.relative_to(PACKAGE).as_posix(),
    )


def write_checksums() -> None:
    manifest = json.loads(BUNDLE_PATH.read_text(encoding="utf-8"))
    write_json(
        CHECKSUMS_PATH,
        {
            "schema_version": "1",
            "source_commit": manifest["source_commit"],
            "files": {
                path.relative_to(PACKAGE).as_posix(): sha256_file(path)
                for path in checksum_paths()
            },
        },
    )


def prepare(args: argparse.Namespace) -> None:
    source_commit = args.source_commit.lower()
    if len(source_commit) != 40 or any(character not in "0123456789abcdef" for character in source_commit):
        raise ValueError("source commit must be a 40-character hexadecimal Git commit")
    created_at = source_created_at(source_commit)

    shutil.copy2(ROOT / "backend/testdata/sample_snapshot.json", PACKAGE / "data/authority.snapshot.json")
    sync_tree(ROOT / "backend/data/regulatory", PACKAGE / "data/regulatory")
    sync_engine_schemas()

    geocoder = PACKAGE / "data/geocoder.sqlite"
    geocoder_manifest = PACKAGE / "data/geocoder.sqlite.manifest.json"
    if not geocoder.is_file():
        raise ValueError("source-built geocoder snapshot is missing")
    if not geocoder_manifest.is_file():
        raise ValueError("source-built geocoder snapshot manifest is missing")
    sync_manifest_data()

    package_coverage = coverage()
    write_json(MANIFESTS / "coverage.json", package_coverage)

    components: dict[str, Any] = {
        "boundary_snapshot": {
            "path": "data/authority.snapshot.json",
            "sha256": sha256_file(MANIFEST_DATA / "authority.snapshot.json"),
        },
        "regulatory_catalog": {
            "path": "data/regulatory",
            "sha256": sha256_directory(MANIFEST_DATA / "regulatory"),
            "recursive": True,
        },
        "geocoder": {
            "path": "data/geocoder.sqlite",
            "sha256": sha256_file(MANIFEST_DATA / "geocoder.sqlite"),
            "optional": True,
        },
    }

    write_json(
        BUNDLE_PATH,
        {
            "schema_version": "1",
            "engine_version": "0.1.0",
            "source_commit": source_commit,
            "created_at": created_at,
            "components": components,
            "coverage": package_coverage,
        },
    )
    write_json(
        VERIFICATION_PATH,
        verification_receipt(
            source_commit=source_commit,
            created_at=created_at,
            status="pending",
            go_version=args.go_version,
            cold_room_mode=None,
            workflow_run=None,
        ),
    )
    write_checksums()


def finalize(args: argparse.Namespace) -> None:
    manifest = json.loads(BUNDLE_PATH.read_text(encoding="utf-8"))
    if manifest["source_commit"] != args.source_commit.lower():
        raise ValueError("verification source commit does not match bundle source commit")
    write_json(
        VERIFICATION_PATH,
        verification_receipt(
            source_commit=args.source_commit.lower(),
            created_at=manifest["created_at"],
            status="passed",
            go_version=args.go_version,
            cold_room_mode=args.cold_room_mode,
            workflow_run=args.workflow_run,
        ),
    )
    write_checksums()


def current_commit() -> str:
    return subprocess.check_output(["git", "rev-parse", "HEAD"], cwd=ROOT, text=True).strip()


def parser() -> argparse.ArgumentParser:
    result = argparse.ArgumentParser()
    subparsers = result.add_subparsers(dest="command", required=True)
    for name in ("prepare", "finalize"):
        command = subparsers.add_parser(name)
        command.add_argument("--source-commit", default=current_commit())
        command.add_argument("--go-version", required=True)
        if name == "finalize":
            command.add_argument("--cold-room-mode", choices=("isolated", "fallback"), required=True)
            command.add_argument("--workflow-run", required=True)
    return result


def main() -> None:
    args = parser().parse_args()
    if args.command == "prepare":
        prepare(args)
    else:
        finalize(args)


if __name__ == "__main__":
    main()
