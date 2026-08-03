#!/usr/bin/env python3
"""Validate and regenerate the Building Code Map Deciduous archaeology."""

from __future__ import annotations

import argparse
import json
import re
import sqlite3
import sys
from collections import defaultdict, deque
from pathlib import Path

NODE_TYPES = {"goal", "option", "decision", "action", "outcome", "observation", "revisit"}
NATIVE_STATUSES = {"pending", "active", "completed", "rejected", "superseded"}
LIFECYCLES = {
    "active",
    "experimental",
    "proposed",
    "superseded",
    "rejected",
    "abandoned",
    "incomplete-data",
    "compatibility-only",
    "historical-only",
    "unresolved",
}
EDGE_TYPES = {"leads_to", "chosen", "rejected", "requires", "blocks", "enables", "supersedes"}
SEMANTIC_ID = re.compile(r"^bcm-[a-z0-9]+(?:-[a-z0-9]+)*$")
FULL_SHA = re.compile(r"^[0-9a-f]{40}$")

EXPECTED_NODE_COUNT = 98
EXPECTED_EDGE_COUNT = 112
EXPECTED_PROFILES = {"colorado.json", "florida.json", "new-jersey.json"}
REQUIRED_PROJECTIONS = {
    "README.md",
    "current-architecture.md",
    "address-to-result-pipeline.md",
    "geocoder-evolution.md",
    "jurisdiction-authority-model.md",
    "adoption-amendment-lifecycle.md",
    "dataset-sources-provenance.md",
    "state-coverage-maturity.md",
    "building-code-map-vs-ast.md",
    "website-product-evolution.md",
    "risk-register.md",
    "maintenance.md",
    "adversarial-review.md",
    "evidence-index.md",
}


def repository_root() -> Path:
    return Path(__file__).resolve().parents[1]


def load_seed(root: Path) -> sqlite3.Connection:
    seed_paths = sorted((root / ".deciduous").glob("deciduous.sql.zlib.b85.part-*"))
    if not seed_paths:
        raise AssertionError("missing canonical Deciduous SQL seed parts")
    import base64
    import zlib

    encoded = "".join(path.read_text(encoding="ascii").strip() for path in seed_paths)
    compressed = base64.b85decode(encoded)
    sql = zlib.decompress(compressed).decode("utf-8")
    connection = sqlite3.connect(":memory:")
    connection.row_factory = sqlite3.Row
    connection.executescript(sql)
    return connection


def rows(connection: sqlite3.Connection, table: str) -> list[dict[str, object]]:
    return [dict(row) for row in connection.execute(f"SELECT * FROM {table} ORDER BY id")]


def export_payload(connection: sqlite3.Connection) -> dict[str, object]:
    version = connection.execute(
        "SELECT version FROM schema_versions ORDER BY id DESC LIMIT 1"
    ).fetchone()[0]
    generated_at = connection.execute(
        "SELECT created_at FROM decision_nodes ORDER BY id LIMIT 1"
    ).fetchone()[0]
    return {
        "schema_version": version,
        "generated_at": generated_at,
        "repository": "laurajoyhutchins/building-code-map",
        "source": ".deciduous/deciduous.sql.zlib.b85.part-*",
        "nodes": rows(connection, "decision_nodes"),
        "edges": rows(connection, "decision_edges"),
    }


def validate_schema(connection: sqlite3.Connection) -> None:
    version = connection.execute(
        "SELECT version, name, features FROM schema_versions ORDER BY id DESC LIMIT 1"
    ).fetchone()
    assert version is not None, "missing schema version"
    assert version[0] == "1.1.0", f"unexpected Deciduous schema version: {version[0]}"
    assert version[1] == "decision-graph", f"unexpected schema name: {version[1]}"
    assert "decision_nodes" in version[2] and "decision_edges" in version[2]


def validate_nodes(nodes: list[dict[str, object]], root: Path, check_paths: bool) -> None:
    assert len(nodes) == EXPECTED_NODE_COUNT, (
        f"expected {EXPECTED_NODE_COUNT} nodes, found {len(nodes)}"
    )
    semantic_ids: set[str] = set()
    current_decisions = 0

    for node in nodes:
        node_type = str(node["node_type"])
        status = str(node["status"])
        change_id = str(node["change_id"])
        assert node_type in NODE_TYPES, f"unsupported node type: {node_type}"
        assert status in NATIVE_STATUSES, f"unsupported native status: {status}"
        assert SEMANTIC_ID.fullmatch(change_id), f"unstable semantic ID: {change_id}"
        assert change_id not in semantic_ids, f"duplicate semantic ID: {change_id}"
        semantic_ids.add(change_id)

        metadata = json.loads(str(node["metadata_json"]))
        assert metadata["semantic_id"] == change_id
        lifecycle = metadata.get("lifecycle_status")
        assert lifecycle in LIFECYCLES, f"unsupported lifecycle {lifecycle!r} for {change_id}"
        evidence = metadata.get("evidence")
        assert isinstance(evidence, list) and evidence, f"missing evidence for {change_id}"

        if metadata.get("current_architecture") and node_type == "decision":
            current_decisions += 1

        for item in evidence:
            assert isinstance(item, dict), f"invalid evidence object for {change_id}"
            revision = item.get("revision")
            if revision is not None:
                assert FULL_SHA.fullmatch(str(revision)), (
                    f"evidence revision must be a full SHA for {change_id}: {revision}"
                )
            for key in ("merge_commit", "head"):
                value = item.get(key)
                if value is not None:
                    assert FULL_SHA.fullmatch(str(value)), (
                        f"{key} must be a full SHA for {change_id}: {value}"
                    )
            repo_path = item.get("repository_path")
            if check_paths and repo_path:
                path = root / str(repo_path)
                assert path.exists(), f"missing evidence path for {change_id}: {repo_path}"

    assert current_decisions >= 10, "current architecture is not marked on enough decisions"


def validate_edges(
    nodes: list[dict[str, object]], edges: list[dict[str, object]]
) -> None:
    assert len(edges) == EXPECTED_EDGE_COUNT, (
        f"expected {EXPECTED_EDGE_COUNT} edges, found {len(edges)}"
    )
    node_by_id = {int(node["id"]): node for node in nodes}
    outgoing: dict[int, list[dict[str, object]]] = defaultdict(list)
    incoming: dict[int, list[dict[str, object]]] = defaultdict(list)
    seen: set[tuple[int, int, str]] = set()

    for edge in edges:
        source = int(edge["from_node_id"])
        target = int(edge["to_node_id"])
        edge_type = str(edge["edge_type"])
        assert source in node_by_id, f"edge source does not exist: {source}"
        assert target in node_by_id, f"edge target does not exist: {target}"
        assert source != target, f"self edge at node {source}"
        assert edge_type in EDGE_TYPES, f"unsupported edge type: {edge_type}"
        key = (source, target, edge_type)
        assert key not in seen, f"duplicate edge: {key}"
        seen.add(key)
        assert edge["from_change_id"] == node_by_id[source]["change_id"]
        assert edge["to_change_id"] == node_by_id[target]["change_id"]
        assert str(edge["rationale"]).strip(), f"missing edge rationale: {key}"
        outgoing[source].append(edge)
        incoming[target].append(edge)

    indegree = {node_id: len(incoming[node_id]) for node_id in node_by_id}
    queue = deque(sorted(node_id for node_id, degree in indegree.items() if degree == 0))
    visited = 0
    while queue:
        source = queue.popleft()
        visited += 1
        for edge in outgoing[source]:
            target = int(edge["to_node_id"])
            indegree[target] -= 1
            if indegree[target] == 0:
                queue.append(target)
    assert visited == len(nodes), "decision graph contains a cycle"

    for node_id, node in node_by_id.items():
        node_type = str(node["node_type"])
        children = [node_by_id[int(edge["to_node_id"])] for edge in outgoing[node_id]]
        if node_type != "goal":
            assert incoming[node_id], f"orphan non-goal: {node['change_id']}"
        if node_type == "goal":
            assert any(child["node_type"] == "option" for child in children), (
                f"goal has no option: {node['change_id']}"
            )
            options = [child for child in children if child["node_type"] == "option"]
            decisions = []
            for option in options:
                decisions.extend(
                    node_by_id[int(edge["to_node_id"])]
                    for edge in outgoing[int(option["id"])]
                    if node_by_id[int(edge["to_node_id"])] ["node_type"] == "decision"
                )
            assert decisions, f"goal awaits a decision: {node['change_id']}"
        if node_type == "decision":
            assert any(child["node_type"] == "action" for child in children), (
                f"decision has no action: {node['change_id']}"
            )
        if node_type == "action":
            assert any(child["node_type"] == "outcome" for child in children), (
                f"action has no outcome: {node['change_id']}"
            )


def validate_export(connection: sqlite3.Connection, root: Path) -> None:
    export_path = root / "docs" / "graph-data.json"
    if not export_path.is_file():
        return
    expected = export_payload(connection)
    actual = json.loads(export_path.read_text(encoding="utf-8"))
    assert actual == expected, (
        "docs/graph-data.json is stale; run "
        "python tools/validate_deciduous_archaeology.py --write-export"
    )


def validate_projections(root: Path) -> None:
    projection_root = root / "docs" / "archaeology"
    actual = {path.name for path in projection_root.glob("*.md")}
    missing = REQUIRED_PROJECTIONS - actual
    assert not missing, f"missing archaeology projections: {sorted(missing)}"

    narratives = root / ".deciduous" / "narratives.md"
    assert narratives.is_file(), "missing Deciduous narratives"
    narrative_text = narratives.read_text(encoding="utf-8")
    assert narrative_text.count("Root: `bcm-goal-") >= 14, (
        "root narratives do not cover the major arcs"
    )


def validate_current_repository_claims(root: Path) -> None:
    profiles_root = root / "backend" / "data" / "regulatory"
    if profiles_root.is_dir():
        profiles = {path.name for path in profiles_root.glob("*.json")}
        assert profiles == EXPECTED_PROFILES, (
            f"state coverage claim mismatch: expected {sorted(EXPECTED_PROFILES)}, "
            f"found {sorted(profiles)}"
        )

    result_schema = root / "schemas" / "regulatory" / "resolution-result.schema.json"
    if result_schema.is_file():
        schema = json.loads(result_schema.read_text(encoding="utf-8"))
        assert schema["properties"]["schema_version"]["const"] == "1.0"

    geocoder_schema = root / "backend" / "internal" / "geocoder" / "schema.go"
    if geocoder_schema.is_file():
        text = geocoder_schema.read_text(encoding="utf-8")
        assert 'const schemaVersion = "1"' in text

    public_source = root / "src"
    if public_source.is_dir():
        source_text = "\n".join(
            path.read_text(encoding="utf-8")
            for path in public_source.rglob("*")
            if path.suffix in {".ts", ".tsx", ".css"}
        )
        removed_badge_row = "Source-backed · Uncertainty preserved · Verify with the AHJ"
        assert removed_badge_row not in source_text


def write_sql(root: Path) -> None:
    import base64
    import zlib

    seed_paths = sorted((root / ".deciduous").glob("deciduous.sql.zlib.b85.part-*"))
    encoded = "".join(path.read_text(encoding="ascii").strip() for path in seed_paths)
    sql = zlib.decompress(base64.b85decode(encoded))
    destination = root / ".deciduous" / "deciduous.sql"
    destination.write_bytes(sql)


def write_database(connection: sqlite3.Connection, root: Path) -> None:
    destination = root / ".deciduous" / "deciduous.db"
    destination.parent.mkdir(parents=True, exist_ok=True)
    if destination.exists():
        destination.unlink()
    disk = sqlite3.connect(destination)
    connection.backup(disk)
    disk.close()


def write_export(connection: sqlite3.Connection, root: Path) -> None:
    destination = root / "docs" / "graph-data.json"
    destination.write_text(
        json.dumps(export_payload(connection), indent=2, ensure_ascii=False) + "\n",
        encoding="utf-8",
    )


def main() -> int:
    parser = argparse.ArgumentParser()
    parser.add_argument("--write-db", action="store_true")
    parser.add_argument("--write-sql", action="store_true")
    parser.add_argument("--write-export", action="store_true")
    parser.add_argument(
        "--skip-repository-paths",
        action="store_true",
        help="validate graph-only material outside a full checkout",
    )
    args = parser.parse_args()

    root = repository_root()
    connection = load_seed(root)
    try:
        validate_schema(connection)
        node_rows = rows(connection, "decision_nodes")
        edge_rows = rows(connection, "decision_edges")
        validate_nodes(node_rows, root, not args.skip_repository_paths)
        validate_edges(node_rows, edge_rows)
        if args.write_export:
            write_export(connection, root)
        validate_export(connection, root)
        validate_projections(root)
        if not args.skip_repository_paths:
            validate_current_repository_claims(root)
        if args.write_sql:
            write_sql(root)
        if args.write_db:
            write_database(connection, root)
    finally:
        connection.close()

    print(
        "Deciduous archaeology valid: "
        f"{EXPECTED_NODE_COUNT} nodes, {EXPECTED_EDGE_COUNT} edges, "
        f"{len(REQUIRED_PROJECTIONS)} projections"
    )
    return 0


if __name__ == "__main__":
    try:
        raise SystemExit(main())
    except AssertionError as error:
        print(f"error: {error}", file=sys.stderr)
        raise SystemExit(1)
