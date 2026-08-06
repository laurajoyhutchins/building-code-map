# Dataset sources and provenance map

## Runtime, build, research, and reference roles

| Source class                                                    | Contribution                                                       | Authority and caveat                                                                         | Runtime role                                       |
| --------------------------------------------------------------- | ------------------------------------------------------------------ | -------------------------------------------------------------------------------------------- | -------------------------------------------------- |
| Local geocoder snapshot                                         | Address points and street ranges                                   | Accuracy, currency, licensing, and redistribution depend on the imported source              | Runtime local artifact                             |
| Local boundary snapshot                                         | State, county, municipality, special, tribal, and fire geometry    | Geometry establishes containment, not code authority                                         | Runtime local artifact                             |
| U.S. Census TIGER/Line and TIGERweb                             | Boundaries and geographic identifiers                              | Public government geography; vintage and transformation must be preserved                    | Ingestion/build input                              |
| NERIS resources                                                 | Fire-jurisdiction geometry and department context where permitted  | Individual departments retain ownership of operational data; redistribution must be reviewed | Ingestion/reference, not hidden request dependency |
| FEMA and NFIP resources                                         | Floodplain or NFIP-community context where integrated              | Authoritative for their published program data, not automatically building-code authority    | Research or future ingestion                       |
| Utility service territories                                     | Utility context where reliable data exists                         | Service territory does not automatically establish code enforcement                          | Future ingestion/research                          |
| State statutes, regulations, registers, and agency publications | State authority, adoption, dates, enforcement, and amendment rules | Usually primary legal or official administrative evidence                                    | Research/build input compiled to policy data       |
| Local ordinances and municipal-code hosts                       | Local adoption, amendments, effective dates, and enforcement       | Controlling instrument and current codification must be reconciled                           | Research/build input                               |
| State code-adoption summaries                                   | Discovery and cross-checking                                       | Secondary unless the state office is itself publishing an official record                    | Research aid                                       |
| ICC, NFPA, ASHRAE, and other publishers                         | Publication identity and edition metadata                          | Publisher material does not prove local adoption; protected text is not bundled              | Reference only                                     |
| OpenStreetMap tiles                                             | Public basemap                                                     | Presentation layer with attribution and service-policy obligations                           | Browser presentation dependency                    |
| Manually curated jurisdiction records                           | Reconciled structured claims                                       | Must retain source IDs, review state, and unresolved conflicts                               | Build input and canonical policy data              |

## Provenance requirements

Material records should retain publisher, source title, stable identifier, URL, access or verification
date, source vintage, availability, caveat, transformation history, and verification status. Address
sources additionally require coordinate reference system, source-record identity, license or
public-record review, and limitations on private roads, rural routes, units, and update frequency.

## Authority hierarchy

An authoritative government source can support a legal claim within its scope. A model-code
publisher can support publication identity. A secondary site can help locate evidence. None should be
silently promoted beyond its role.

## Publication boundaries

The repository does not relicense third-party data. Hydrated SQLite and DuckDB files are local
artifacts. Model-code text, standards, protected commentary, NERIS operational records, and
unreviewed address datasets are not committed merely because they are obtainable.

## Refresh expectations

Boundary, address, source-health, adoption, and amendment data have independent cadences. Local-first
execution therefore requires periodically regenerated artifacts rather than permanently frozen data.
Automated refresh and stale-data enforcement are incomplete.
