from __future__ import annotations

from pathlib import Path

import yaml
from yaml.constructor import ConstructorError


class _UniqueKeyLoader(yaml.SafeLoader):
    pass


def _construct_unique_mapping(loader: _UniqueKeyLoader, node: yaml.MappingNode, deep: bool = False):
    mapping: dict[object, object] = {}
    for key_node, value_node in node.value:
        key = loader.construct_object(key_node, deep=deep)
        if key in mapping:
            raise ConstructorError(
                "while constructing a mapping",
                node.start_mark,
                f"found duplicate key {key!r}",
                key_node.start_mark,
            )
        mapping[key] = loader.construct_object(value_node, deep=deep)
    return mapping


_UniqueKeyLoader.add_constructor(
    yaml.resolver.BaseResolver.DEFAULT_MAPPING_TAG,
    _construct_unique_mapping,
)


class CleanDumper(yaml.SafeDumper):
    pass


def _represent_none(dumper, _):
    return dumper.represent_scalar("tag:yaml.org,2002:null", "null")


def _represent_str(dumper, data):
    if data == "":
        return dumper.represent_scalar("tag:yaml.org,2002:str", data, style="'")
    return dumper.represent_scalar("tag:yaml.org,2002:str", data)


CleanDumper.add_representer(type(None), _represent_none)
CleanDumper.add_representer(str, _represent_str)


def load_yaml(text: str) -> object:
    return yaml.load(text, Loader=_UniqueKeyLoader)


def load_yaml_file(path: Path) -> object:
    return load_yaml(path.read_text(encoding="utf-8"))


def dump_yaml(document: object) -> str:
    return yaml.dump(
        document,
        Dumper=CleanDumper,
        sort_keys=False,
        default_flow_style=False,
        allow_unicode=True,
    )
