# Data Sources and Attribution

The Apache 2.0 license in this repository applies to project-authored software and documentation. It does not relicense third-party map tiles, government records, department-owned data, statutes, regulations, standards, or other source material.

## OpenStreetMap

The interactive basemap requests standard raster tiles from OpenStreetMap and displays `© OpenStreetMap contributors` through MapLibre's attribution control.

- Copyright and license: https://www.openstreetmap.org/copyright
- Standard tile usage policy: https://operations.osmfoundation.org/policies/tiles/

Production deployments must continue to show attribution and must comply with the tile service's usage policy. A high-volume or offline deployment should use an appropriate hosted or self-operated tile service rather than treating the community tile servers as an unlimited production CDN.

## U.S. Census Bureau TIGER/Line and TIGERweb

Boundary geometry and geographic identifiers may be derived from U.S. Census Bureau TIGER/Line and TIGERweb services.

- TIGER/Line product page: https://www.census.gov/geographies/mapping-files/time-series/geo/tiger-line-file.html
- Census geography mapping files: https://www.census.gov/geographies/mapping-files.html

Preserve the source vintage, layer or service identifier, retrieval date, and transformation history when material is mirrored or transformed. Census geography products describe boundaries and identifiers; they should not be represented as containing demographic data unless a separate demographic source is explicitly joined.

## NERIS

The project may model National Emergency Response Information System jurisdiction geometry and department attributes. The U.S. Fire Administration states that individual fire departments retain ownership of their NERIS data.

- NERIS data FAQ: https://www.usfa.fema.gov/nfirs/neris/faq/data/
- NERIS overview: https://www.usfa.fema.gov/nfirs/neris/about-neris/

This repository must not publish department-owned operational, incident, personnel, exposure, or otherwise restricted NERIS records without confirmed authorization and redistribution terms. Public schemas, documentation, identifiers, and explicitly public jurisdiction resources should remain distinguishable from protected operational data.

## Building-Code Research

State reports and normalized records should be grounded in official sources such as statutes, administrative rules, agency publications, registers, and adoption documents. Each evidence-bearing claim should retain:

- a stable source identifier;
- the official publisher and source URL;
- an access or verification date;
- the relevant adoption, effective, operative, or mandatory date;
- a conservative verification status and confidence;
- explicit unresolved fields instead of inferred certainty.

Commercial model-code text, standards, and copyrighted commentary are not included merely because a jurisdiction adopts them. Record editions, adoption instruments, amendments, applicability, and authority without copying protected code text beyond what is legally permitted.

## Runtime Snapshots

Hydrated SQLite and DuckDB snapshots are local runtime artifacts and are ignored by Git. A public release should distribute reproducible ingestion logic and provenance metadata rather than local databases whose contents and redistribution rights have not been independently reviewed.