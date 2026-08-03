# Adoption and amendment lifecycle map

## Publication is not adoption

ICC, NFPA, ASHRAE, and other publishers issue model codes and standards. A jurisdiction applies
them through a statute, regulation, ordinance, resolution, administrative rule, or other controlling
instrument. Publication metadata can identify a source edition, but it cannot establish local law.

## Date distinctions

The data model must keep these dates separate when evidenced:

1. publication date of the model code;
2. adoption or execution date of the legal instrument;
3. effective date of the adopted requirement;
4. operative or mandatory date;
5. transition or permit-stage date;
6. replacement, repeal, or supersession date.

The current profiles support date records and applicability filtering, but completeness varies by
state and code family.

## Adoption record dimensions

A supported record may include:

- code family;
- model-code edition or state code name;
- adopting authority;
- adoption instrument;
- enforcement model;
- scope and project class;
- statewide baseline or local-adoption posture;
- date fields;
- source identifiers;
- verification status and confidence;
- conflict or unresolved state.

The resolver can return several records when code families or scopes differ.

## Amendments

The rule-pack model includes `amendment` as a first-class rule kind. It can describe statewide
amendment sets, local amendment authority, constraints, and source references. This representation
does not imply comprehensive amendment content.

Major remaining gaps include:

- complete municipal amendment inventories;
- deleted and inserted sections;
- local design criteria and climatic values;
- optional chapter and appendix adoption;
- fire-code appendices;
- electrical amendments;
- energy-code compliance pathways;
- conflicts between ordinances and codified hosts;
- amendment supersession and historical effective periods;
- inaccessible or redistribution-sensitive source documents.

## Conservative output

When a statewide rule does not establish local applicability, the result should identify candidate
authorities and required local records. Inferred amendments must not be presented as confirmed. A
missing date must not be silently replaced with the model-code publication date.

## Current maturity

Colorado, Florida, and New Jersey contain source-backed pilot records and exercise differing models,
but issue #31 remains the evidence-backed path to production-ready timelines, amendments, local
fixtures, and special-project coverage. Draft PR #38 is not merged and therefore does not change
current main classifications.
