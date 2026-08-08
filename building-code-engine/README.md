# Building Code Authority Engine offline package

This package is restored from Building Code Map source commit
`e45d3674c09333e1682509e1a1a544363e7ab127`. It contains the `bcm` binary,
content-addressed bundle data, regulatory profiles, schemas, and checksums.

From this directory:

```text
bin/bcm inspect bundle --bundle manifests/bundle.json
bin/bcm resolve --bundle manifests/bundle.json --point -104.99,39.74 --as-of 2026-08-06
bin/bcm serve --bundle manifests/bundle.json --http 127.0.0.1:8000
bin/bcm serve --bundle manifests/bundle.json --mcp-stdio
```

All normal queries are local and require no network access. The aggregate
manifest identifies the exact component bytes; PR #45 snapshot manifests and
activation receipts remain the component-level identity authority.
