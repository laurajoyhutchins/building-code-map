from __future__ import annotations

import re
from dataclasses import dataclass
from typing import Literal

ValueKind = Literal[
    "string",
    "string_list",
    "number",
    "nullable_string",
    "detail_list",
]


@dataclass(frozen=True)
class FieldSpec:
    kind: ValueKind


@dataclass(frozen=True)
class ArtifactSchema:
    root_key: str
    fields: dict[str, FieldSpec]


AUTHORITY_FIELDS = {
    "authority_id": FieldSpec("string"),
    "authority_name": FieldSpec("string"),
    "authority_type": FieldSpec("string"),
    "legal_basis": FieldSpec("string_list"),
    "role": FieldSpec("string"),
    "enforcement_model": FieldSpec("string"),
    "source_ids": FieldSpec("string_list"),
    "verification_status": FieldSpec("string_list"),
    "parent_agency": FieldSpec("nullable_string"),
}

KEY_FINDING_FIELDS = {
    "topic": FieldSpec("string"),
    "finding": FieldSpec("string"),
    "confidence": FieldSpec("number"),
    "source_ids": FieldSpec("string_list"),
}

OFFICIAL_SOURCE_FIELDS = {
    "source_id": FieldSpec("string"),
    "title": FieldSpec("string"),
    "source_type": FieldSpec("nullable_string"),
    "publisher": FieldSpec("nullable_string"),
    "url": FieldSpec("nullable_string"),
    "accessed_date": FieldSpec("nullable_string"),
    "snapshot_id": FieldSpec("nullable_string"),
    "checksum": FieldSpec("nullable_string"),
    "status": FieldSpec("nullable_string"),
    "supports": FieldSpec("string_list"),
    "caveats": FieldSpec("string_list"),
    "notes": FieldSpec("string_list"),
    "details": FieldSpec("detail_list"),
}

SCHEMAS = {
    schema.root_key: schema
    for schema in (
        ArtifactSchema("authority_structure", AUTHORITY_FIELDS),
        ArtifactSchema("key_findings", KEY_FINDING_FIELDS),
        ArtifactSchema("official_sources", OFFICIAL_SOURCE_FIELDS),
    )
}

AUTHORITY_ALIASES = {
    "authority_id": "authority_id",
    "authority_name": "authority_name",
    "authority_type": "authority_type",
    "legal_basis": "legal_basis",
    "legal_regulatory_basis": "legal_basis",
    "role": "role",
    "verified_role": "role",
    "enforcement_model": "enforcement_model",
    "source_ids": "source_ids",
    "status": "verification_status",
    "verification_status": "verification_status",
    "parent_agency": "parent_agency",
}

OFFICIAL_SOURCE_ALIASES = {
    "source_id": "source_id",
    "title": "title",
    "source_name": "title",
    "title_citation": "title",
    "title_description": "title",
    "title_locator": "title",
    "source_type": "source_type",
    "type": "source_type",
    "publisher": "publisher",
    "publisher_agency": "publisher",
    "publisher_authority": "publisher",
    "publisher_custodian": "publisher",
    "publisher_maintainer": "publisher",
    "publisher_owner": "publisher",
    "issuer": "publisher",
    "issuer_publisher": "publisher",
    "issuing_authority": "publisher",
    "url": "url",
    "url_citation": "url",
    "url_location": "url",
    "url_locator": "url",
    "locator": "url",
    "access_path": "url",
    "accessed_date": "accessed_date",
    "accessed": "accessed_date",
    "access_date": "accessed_date",
    "date_accessed": "accessed_date",
    "last_checked": "accessed_date",
    "snapshot_id": "snapshot_id",
    "checksum": "checksum",
    "status": "status",
    "official_status": "status",
    "supports": "supports",
    "key_fields_supported": "supports",
    "key_facts_supported": "supports",
    "key_facts_used": "supports",
    "key_coverage": "supports",
    "key_use": "supports",
    "records_supported": "supports",
    "supported_fields": "supports",
    "used_for": "supports",
    "primary_use": "supports",
    "primary_extracted_fields": "supports",
    "coverage": "supports",
    "caveat": "caveats",
    "caveats": "caveats",
    "notes": "notes",
    "access_notes": "notes",
}


def normalize_field_name(value: str) -> str:
    stripped = value.strip()
    if stripped.startswith("**") and stripped.endswith("**") and len(stripped) >= 4:
        stripped = stripped[2:-2].strip()
    return re.sub(r"[^a-z0-9]+", "_", stripped.lower()).strip("_")


def get_schema(root_key: str) -> ArtifactSchema:
    try:
        return SCHEMAS[root_key]
    except KeyError as exc:
        supported = ", ".join(sorted(SCHEMAS))
        raise ValueError(
            f"No explicit report schema is registered for '{root_key}'. "
            f"Supported schemas: {supported}."
        ) from exc


def _value_matches_kind(value: object, kind: ValueKind) -> bool:
    if kind == "string":
        return isinstance(value, str)
    if kind == "string_list":
        return isinstance(value, list) and all(isinstance(item, str) for item in value)
    if kind == "nullable_string":
        return value is None or isinstance(value, str)
    if kind == "number":
        return isinstance(value, (int, float)) and not isinstance(value, bool)
    if kind == "detail_list":
        return (
            isinstance(value, list)
            and all(
                isinstance(item, dict)
                and set(item) == {"name", "values"}
                and isinstance(item["name"], str)
                and isinstance(item["values"], list)
                and all(isinstance(entry, str) for entry in item["values"])
                for item in value
            )
        )
    raise AssertionError(f"Unhandled value kind: {kind}")


def validate_record(record: object, schema: ArtifactSchema, location: str) -> None:
    if not isinstance(record, dict):
        raise ValueError(f"{location} must be a mapping.")

    expected = set(schema.fields)
    actual = set(record)
    missing = sorted(expected - actual)
    unknown = sorted(actual - expected)
    if missing:
        raise ValueError(f"{location} is missing required fields: {', '.join(missing)}.")
    if unknown:
        raise ValueError(f"{location} contains unknown fields: {', '.join(unknown)}.")

    for field, spec in schema.fields.items():
        if not _value_matches_kind(record[field], spec.kind):
            raise ValueError(
                f"{location}.{field} must match schema type '{spec.kind}', "
                f"got {type(record[field]).__name__}."
            )
        if spec.kind == "detail_list":
            names = [item["name"] for item in record[field]]
            if len(names) != len(set(names)):
                raise ValueError(f"{location}.{field} contains duplicate detail names.")


def validate_root_value(value: object, schema: ArtifactSchema) -> None:
    if not isinstance(value, list):
        raise ValueError(f"{schema.root_key} must be a list.")
    if not value:
        raise ValueError(f"{schema.root_key} must contain at least one record.")
    for index, record in enumerate(value):
        validate_record(record, schema, f"{schema.root_key}[{index}]")


def validate_document(document: object) -> None:
    if not isinstance(document, dict):
        raise ValueError("document root must be a mapping.")
    if len(document) != 1:
        raise ValueError("expected exactly one generated-artifact root key.")
    root_key, value = next(iter(document.items()))
    if not isinstance(root_key, str):
        raise ValueError("artifact root key must be a string.")
    validate_root_value(value, get_schema(root_key))


def _string(value: object, location: str, *, allow_singleton_list: bool = False) -> str:
    if isinstance(value, str):
        return value
    if allow_singleton_list and isinstance(value, list) and len(value) == 1 and isinstance(value[0], str):
        return value[0]
    raise ValueError(f"{location} must be a string.")


def _nullable_string(value: object, location: str) -> str | None:
    if value is None or isinstance(value, str):
        return value
    if isinstance(value, (int, float)) and not isinstance(value, bool):
        return str(value)
    if isinstance(value, list) and len(value) == 1 and isinstance(value[0], str):
        return value[0]
    raise ValueError(f"{location} must be a string or null.")


def _split_legacy_items(value: str, delimiters: tuple[str, ...]) -> list[str]:
    if value == "":
        return []
    if not delimiters:
        return [value]
    pattern = "[" + re.escape("".join(delimiters)) + "]"
    return [item.strip() for item in re.split(pattern, value) if item.strip()]


def _string_list(
    value: object,
    location: str,
    *,
    delimiters: tuple[str, ...] = (),
) -> list[str]:
    if value is None:
        return []
    if isinstance(value, str):
        return _split_legacy_items(value, delimiters)
    if isinstance(value, list):
        converted: list[str] = []
        for index, item in enumerate(value):
            if not isinstance(item, str):
                raise ValueError(f"{location}[{index}] must be a string.")
            converted.extend(_split_legacy_items(item, delimiters))
        return converted
    if isinstance(value, (int, float)) and not isinstance(value, bool):
        return [str(value)]
    raise ValueError(f"{location} must be a string or a list of strings.")


def _number(value: object, location: str) -> int | float:
    if isinstance(value, (int, float)) and not isinstance(value, bool):
        return value
    if isinstance(value, str):
        try:
            return float(value)
        except ValueError as exc:
            raise ValueError(f"{location} must be numeric.") from exc
    raise ValueError(f"{location} must be numeric.")


def _details_values(value: object, location: str) -> list[str]:
    if value is None:
        return []
    if isinstance(value, str):
        return [value]
    if isinstance(value, (int, float)) and not isinstance(value, bool):
        return [str(value)]
    if isinstance(value, list):
        converted: list[str] = []
        for index, item in enumerate(value):
            if isinstance(item, str):
                converted.append(item)
            elif isinstance(item, (int, float)) and not isinstance(item, bool):
                converted.append(str(item))
            else:
                raise ValueError(f"{location}[{index}] must be scalar.")
        return converted
    raise ValueError(f"{location} must be scalar or a list of scalars.")


def _assign_alias(
    output: dict[str, object],
    canonical_key: str,
    value: object,
    location: str,
) -> None:
    if canonical_key in output:
        raise ValueError(f"{location} maps more than one input field to '{canonical_key}'.")
    output[canonical_key] = value


def _canonicalize_authority_record(record: object, location: str) -> dict[str, object]:
    if not isinstance(record, dict):
        raise ValueError(f"{location} must be a mapping.")
    mapped: dict[str, object] = {}
    for raw_key, value in record.items():
        if not isinstance(raw_key, str):
            raise ValueError(f"{location} has a non-string field name.")
        normalized = normalize_field_name(raw_key)
        canonical = AUTHORITY_ALIASES.get(normalized)
        if canonical is None:
            raise ValueError(f"{location} contains unknown authority field '{raw_key}'.")
        _assign_alias(mapped, canonical, value, location)

    required = {
        "authority_id",
        "authority_name",
        "authority_type",
        "legal_basis",
        "role",
        "enforcement_model",
        "verification_status",
    }
    missing = sorted(required - set(mapped))
    if missing:
        raise ValueError(f"{location} is missing required fields: {', '.join(missing)}.")

    canonical_record = {
        "authority_id": _string(mapped["authority_id"], f"{location}.authority_id"),
        "authority_name": _string(mapped["authority_name"], f"{location}.authority_name"),
        "authority_type": _string(mapped["authority_type"], f"{location}.authority_type"),
        "legal_basis": _string_list(
            mapped["legal_basis"], f"{location}.legal_basis", delimiters=(";",)
        ),
        "role": _string(mapped["role"], f"{location}.role"),
        "enforcement_model": _string(mapped["enforcement_model"], f"{location}.enforcement_model"),
        "source_ids": _string_list(
            mapped.get("source_ids"), f"{location}.source_ids", delimiters=(",", ";")
        ),
        "verification_status": _string_list(
            mapped["verification_status"],
            f"{location}.verification_status",
            delimiters=(";",),
        ),
        "parent_agency": _nullable_string(mapped.get("parent_agency"), f"{location}.parent_agency"),
    }
    validate_record(canonical_record, SCHEMAS["authority_structure"], location)
    return canonical_record


def _canonicalize_authority(value: object, location: str) -> list[dict[str, object]]:
    if isinstance(value, dict):
        return [_canonicalize_authority_record(value, f"{location}[0]")]
    if not isinstance(value, list):
        raise ValueError(f"{location} must be a list or mapping.")
    if value and all(isinstance(item, dict) and set(item) == {"field", "value"} for item in value):
        record: dict[str, object] = {}
        for index, item in enumerate(value):
            raw_field = item["field"]
            if not isinstance(raw_field, str):
                raise ValueError(f"{location}[{index}].field must be a string.")
            field = normalize_field_name(raw_field)
            if field in record:
                raise ValueError(f"{location}[{index}] duplicates field '{field}'.")
            record[field] = item["value"]
        return [_canonicalize_authority_record(record, f"{location}[0]")]
    return [
        _canonicalize_authority_record(record, f"{location}[{index}]")
        for index, record in enumerate(value)
    ]


def _canonicalize_key_findings(value: object, location: str) -> list[dict[str, object]]:
    if not isinstance(value, list):
        raise ValueError(f"{location} must be a list.")
    result: list[dict[str, object]] = []
    for index, record in enumerate(value):
        item_location = f"{location}[{index}]"
        if not isinstance(record, dict):
            raise ValueError(f"{item_location} must be a mapping.")
        unknown = sorted(set(record) - set(KEY_FINDING_FIELDS))
        missing = sorted(set(KEY_FINDING_FIELDS) - set(record))
        if unknown:
            raise ValueError(f"{item_location} contains unknown fields: {', '.join(unknown)}.")
        if missing:
            raise ValueError(f"{item_location} is missing required fields: {', '.join(missing)}.")
        canonical_record = {
            "topic": _string(record["topic"], f"{item_location}.topic"),
            "finding": _string(
                record["finding"], f"{item_location}.finding", allow_singleton_list=True
            ),
            "confidence": _number(record["confidence"], f"{item_location}.confidence"),
            "source_ids": _string_list(
                record["source_ids"],
                f"{item_location}.source_ids",
                delimiters=(",", ";"),
            ),
        }
        validate_record(canonical_record, SCHEMAS["key_findings"], item_location)
        result.append(canonical_record)
    return result


def _normalize_details(value: object, location: str) -> list[dict[str, object]]:
    if value is None:
        return []
    if not isinstance(value, list):
        raise ValueError(f"{location} must be a list.")
    details: list[dict[str, object]] = []
    seen: set[str] = set()
    for index, item in enumerate(value):
        item_location = f"{location}[{index}]"
        if not isinstance(item, dict) or set(item) != {"name", "values"}:
            raise ValueError(f"{item_location} must contain name and values.")
        name = _string(item["name"], f"{item_location}.name")
        if name in seen:
            raise ValueError(f"{location} duplicates detail name '{name}'.")
        seen.add(name)
        details.append(
            {
                "name": name,
                "values": _details_values(item["values"], f"{item_location}.values"),
            }
        )
    return details


def _canonicalize_official_source_record(record: object, location: str) -> dict[str, object]:
    if not isinstance(record, dict):
        raise ValueError(f"{location} must be a mapping.")
    mapped: dict[str, object] = {}
    detail_items = _normalize_details(record.get("details"), f"{location}.details")
    detail_names = {item["name"] for item in detail_items}
    for raw_key, value in record.items():
        if not isinstance(raw_key, str):
            raise ValueError(f"{location} has a non-string field name.")
        normalized = normalize_field_name(raw_key)
        if normalized == "details":
            continue
        canonical = OFFICIAL_SOURCE_ALIASES.get(normalized)
        if canonical is None:
            if normalized in detail_names:
                raise ValueError(f"{location} duplicates detail name '{normalized}'.")
            detail_names.add(normalized)
            detail_items.append(
                {
                    "name": normalized,
                    "values": _details_values(value, f"{location}.{raw_key}"),
                }
            )
            continue
        _assign_alias(mapped, canonical, value, location)

    missing = sorted({"source_id", "title"} - set(mapped))
    if missing:
        raise ValueError(f"{location} is missing required fields: {', '.join(missing)}.")

    canonical_record = {
        "source_id": _string(mapped["source_id"], f"{location}.source_id"),
        "title": _string(mapped["title"], f"{location}.title"),
        "source_type": _nullable_string(mapped.get("source_type"), f"{location}.source_type"),
        "publisher": _nullable_string(mapped.get("publisher"), f"{location}.publisher"),
        "url": _nullable_string(mapped.get("url"), f"{location}.url"),
        "accessed_date": _nullable_string(mapped.get("accessed_date"), f"{location}.accessed_date"),
        "snapshot_id": _nullable_string(mapped.get("snapshot_id"), f"{location}.snapshot_id"),
        "checksum": _nullable_string(mapped.get("checksum"), f"{location}.checksum"),
        "status": _nullable_string(mapped.get("status"), f"{location}.status"),
        "supports": _string_list(mapped.get("supports"), f"{location}.supports"),
        "caveats": _string_list(mapped.get("caveats"), f"{location}.caveats"),
        "notes": _string_list(mapped.get("notes"), f"{location}.notes"),
        "details": sorted(detail_items, key=lambda item: item["name"]),
    }
    validate_record(canonical_record, SCHEMAS["official_sources"], location)
    return canonical_record


def _canonicalize_official_sources(value: object, location: str) -> list[dict[str, object]]:
    if not isinstance(value, list):
        raise ValueError(f"{location} must be a list.")
    return [
        _canonicalize_official_source_record(record, f"{location}[{index}]")
        for index, record in enumerate(value)
    ]


def canonicalize_document(document: object, source_name: str = "<memory>") -> dict[str, object]:
    if not isinstance(document, dict):
        raise ValueError(f"{source_name}: document root must be a mapping.")
    if len(document) != 1:
        raise ValueError(f"{source_name}: expected exactly one generated-artifact root key.")
    root_key, value = next(iter(document.items()))
    if not isinstance(root_key, str):
        raise ValueError(f"{source_name}: artifact root key must be a string.")
    get_schema(root_key)
    location = f"{source_name}:{root_key}"
    if root_key == "authority_structure":
        canonical_value = _canonicalize_authority(value, location)
    elif root_key == "key_findings":
        canonical_value = _canonicalize_key_findings(value, location)
    elif root_key == "official_sources":
        canonical_value = _canonicalize_official_sources(value, location)
    else:
        raise AssertionError(f"Unhandled root key: {root_key}")
    canonical = {root_key: canonical_value}
    validate_document(canonical)
    return canonical
