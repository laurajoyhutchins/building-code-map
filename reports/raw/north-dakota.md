---
state:
  state_id: "US-ND"
  name: "North Dakota"
  abbreviation: "ND"
report:
  report_id: "state-report:usa-nd"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-26"
  last_verified: "2026-06-26"
  reviewed_by: null
risk:
  overall_confidence: 0.68 # 0.00 - 1.00
  risk_flags:
    - "fire_2024_update_requires_final_codification_check"
    - "local_amendment_registry_not_found"
    - "elevator_code_edition_not_verified"
    - "permit_transition_rules_not_fully_verified"
  open_questions_count: 5

---

# State Building Code Authority Report: North Dakota

## 1. Executive Summary

- **Authority model:** North Dakota has a state building code framework administered by the North Dakota Department of Commerce, Division of Community Services, with state-code update and amendment action tied to the Building Code Advisory Committee and eligible local jurisdictions. The State Building Code applies as the required code for cities, townships, and counties that elect to administer and enforce a building code, and separate trade or safety programs administer electrical, plumbing, fire, and conveyance matters.

- **Statewide code status:** The current North Dakota State Building Code is based on the 2024 IBC, 2024 IRC, 2024 IMC, 2024 IFGC, 2024 IECC, and 2024 IEBC, with state amendments, effective 2026-01-01. The State Building Code does not include the International Property Maintenance Code or the International Fire Code; those are treated separately.

- **Local enforcement model:** Local enforcement is a hybrid opt-in model. Cities, townships, and counties that elect to administer and enforce a building code must adopt and enforce the State Building Code, with counties covering areas not administered by a city or township. State agencies, local government owners, and schools retain compliance responsibilities for their own buildings in the circumstances described in the cited state sources.

- **Local amendment posture:** For the State Building Code, cities, townships, and counties that elect to enforce may amend the State Building Code to conform to local needs. A centralized local-amendment registry or state approval requirement for those local building amendments was not confirmed in the sources reviewed. Electrical local requirements may be made more stringent by ordinance, subject to statutory limits. Plumbing and fire local amendment scope remains narrower and requires follow-up.

- **Known transition periods or pending changes:** The current State Building Code became effective 2026-01-01. Electrical rules reviewed from the State Electrical Board identify 2023 NEC-based standards effective 2024-07-01 and an official administrative-code supplement indicating transition to 2026 NEC-based standards effective 2026-07-01. A 2026 State Fire Marshal rulemaking copy indicates a proposed or pending move from 2021 IFC to 2024 IFC, but final codification was not treated as verified in this report.

- **Production readiness:** partial

### Key Findings

```yaml
---
key_findings:
- topic: Primary statewide building-code authority
  finding: The North Dakota Department of Commerce, Division of Community Services,
    administers the State Building Code framework, with update and amendment voting
    by the Building Code Advisory Committee and eligible local jurisdictions.
  confidence: 0.9
  source_ids:
  - src:usa-nd:commerce-building-codes
  - src:usa-nd:state-building-code-2026-book
  - src:usa-nd:ndcc-54-21-3
  - src:usa-nd:ndac-108-01-01
- topic: Current State Building Code editions
  finding: The 2026 State Building Code uses the 2024 IBC, IRC, IMC, IFGC, IECC, and
    IEBC, with state amendments, effective 2026-01-01.
  confidence: 0.92
  source_ids:
  - src:usa-nd:commerce-building-codes
  - src:usa-nd:state-building-code-2026-book
- topic: Local building-code enforcement
  finding: Local jurisdictions that elect to enforce a building code must adopt and
    enforce the State Building Code; counties may administer areas outside city or
    township administration.
  confidence: 0.85
  source_ids:
  - src:usa-nd:state-building-code-2026-book
  - src:usa-nd:ndcc-54-21-3
  - src:usa-nd:ndac-108-01-01
- topic: Local building amendments
  finding: Electing cities, townships, and counties may amend the State Building Code
    to conform to local needs; a statewide amendment registry was not found in the
    reviewed sources.
  confidence: 0.75
  source_ids:
  - src:usa-nd:state-building-code-2026-book
  - src:usa-nd:ndcc-54-21-3
  - src:usa-nd:ndac-108-01-01
- topic: Electrical authority and code basis
  finding: The State Electrical Board has statewide jurisdiction over electrical installations
    and identifies 2023 NEC-based standards effective 2024-07-01, with a 2026 NEC
    transition indicated for 2026-07-01.
  confidence: 0.88
  source_ids:
  - src:usa-nd:electrical-board-laws-rules
  - src:usa-nd:ndac-24-1-06-01
  - src:usa-nd:ndac-24-1-2026-supplement
  - src:usa-nd:ndcc-43-09
- topic: Plumbing authority and code basis
  finding: The State Plumbing Board administers the North Dakota Plumbing Code, based
    on the 2018 Uniform Plumbing Code, effective 2020-04-01.
  confidence: 0.84
  source_ids:
  - src:usa-nd:plumbing-board-laws-rules
  - src:usa-nd:ndac-62-01-01
  - src:usa-nd:ndac-62-03-1-01
- topic: Fire authority and operational fire code
  finding: The State Fire Marshal administers the State Fire Code, currently verified
    from codified NDAC as the 2021 IFC with North Dakota modifications, effective
    2024-01-01. A 2026 rulemaking copy points toward 2024 IFC but needs final codification
    review.
  confidence: 0.75
  source_ids:
  - src:usa-nd:ndac-45-18-01
  - src:usa-nd:insurance-fire-rulemaking-2026
  - src:usa-nd:ndcc-18-01
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-nd:commerce-dcs |
| Authority name | North Dakota Department of Commerce, Division of Community Services |
| Authority type | state agency division |
| Legal basis | North Dakota Century Code chapter 54-21.3 and North Dakota Administrative Code article 108-01-01 |
| Role | Administers State Building Code rules, publishes the State Building Code book, coordinates Building Code Advisory Committee and eligible-jurisdiction amendment/update process |
| Enforcement model | hybrid: state code baseline plus local election to administer and enforce, with specific state/local government and school owner responsibilities |
| Source IDs | src:usa-nd:commerce-building-codes; src:usa-nd:state-building-code-2026-book; src:usa-nd:ndcc-54-21-3; src:usa-nd:ndac-108-01-01 |
| Verification status | verified_core |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-nd:commerce-dcs; ahj:usa-nd:local-building-jurisdiction | ND Department of Commerce, DCS; electing cities, townships, and counties | State code administration and publication; local adoption and enforcement where elected | NDCC ch. 54-21.3; NDAC art. 108-01-01 | src:usa-nd:state-building-code-2026-book; src:usa-nd:ndcc-54-21-3; src:usa-nd:ndac-108-01-01 | verified_core |
| Residential | ahj:usa-nd:commerce-dcs; ahj:usa-nd:local-building-jurisdiction | ND Department of Commerce, DCS; electing local jurisdictions | Same model as building code for IRC-covered work | NDCC ch. 54-21.3; NDAC art. 108-01-01 | src:usa-nd:state-building-code-2026-book; src:usa-nd:ndcc-54-21-3; src:usa-nd:ndac-108-01-01 | verified_core |
| Existing Building / Rehabilitation | ahj:usa-nd:commerce-dcs; ahj:usa-nd:local-building-jurisdiction | ND Department of Commerce, DCS; electing local jurisdictions | IEBC included in current State Building Code book and local enforcement framework | State Building Code publication and administrative rules | src:usa-nd:commerce-building-codes; src:usa-nd:state-building-code-2026-book; src:usa-nd:ndac-108-01-01 | partially_verified |
| Mechanical | ahj:usa-nd:commerce-dcs; ahj:usa-nd:local-building-jurisdiction | ND Department of Commerce, DCS; electing local jurisdictions | IMC state-code administration and local enforcement where elected | NDCC ch. 54-21.3; NDAC art. 108-01-01 | src:usa-nd:state-building-code-2026-book; src:usa-nd:ndcc-54-21-3; src:usa-nd:ndac-108-01-01 | verified_core |
| Plumbing | ahj:usa-nd:state-plumbing-board | North Dakota State Plumbing Board | Adopts and administers the North Dakota Plumbing Code; inspectors and executive director administer plumbing laws, rules, installation standards, and code | NDCC ch. 43-18; NDAC title 62 | src:usa-nd:plumbing-board-laws-rules; src:usa-nd:ndac-62-01-01; src:usa-nd:ndac-62-03-1-01 | verified_core |
| Fuel Gas | ahj:usa-nd:commerce-dcs; ahj:usa-nd:local-building-jurisdiction | ND Department of Commerce, DCS; electing local jurisdictions | IFGC state-code administration and local enforcement where elected | NDCC ch. 54-21.3; NDAC art. 108-01-01 | src:usa-nd:state-building-code-2026-book; src:usa-nd:ndcc-54-21-3; src:usa-nd:ndac-108-01-01 | verified_core |
| Electrical | ahj:usa-nd:state-electrical-board | North Dakota State Electrical Board | Licensing, standards, and inspection jurisdiction for electrical installations; conveyance regulation added under NDCC ch. 43-09 | NDCC ch. 43-09; NDAC title 24.1 | src:usa-nd:electrical-board-laws-rules; src:usa-nd:ndac-24-1-06-01; src:usa-nd:ndcc-43-09 | verified_core |
| Energy | ahj:usa-nd:commerce-dcs; ahj:usa-nd:local-building-jurisdiction | ND Department of Commerce, DCS; electing local jurisdictions | IECC and IRC energy provisions in State Building Code; local enforcement where elected | State Building Code publication and NDAC art. 108-01-01 | src:usa-nd:commerce-building-codes; src:usa-nd:state-building-code-2026-book; src:usa-nd:ndac-108-01-01 | verified_core |
| Fire - construction references | ahj:usa-nd:commerce-dcs; ahj:usa-nd:state-fire-marshal | ND Department of Commerce, DCS; State Fire Marshal | Building-code fire-safe construction provisions in State Building Code; Fire Code separately references State Building Code for fire-safe construction | NDCC ch. 54-21.3; NDAC ch. 45-18-01 | src:usa-nd:state-building-code-2026-book; src:usa-nd:ndac-45-18-01 | partially_verified |
| Fire - operational / prevention code | ahj:usa-nd:state-fire-marshal | North Dakota State Fire Marshal | State Fire Code administration, including operational and prevention provisions based on IFC with North Dakota modifications | NDCC ch. 18-01; NDAC ch. 45-18-01 | src:usa-nd:ndac-45-18-01; src:usa-nd:ndcc-18-01 | verified_core |
| Accessibility | ahj:usa-nd:commerce-dcs; ahj:usa-nd:state-and-local-public-owners | ND Department of Commerce, DCS; state agencies and political subdivisions | State-law reference to 2010 ADA Standards and required statement process; building-code amendment limitations for federal ADA standards | NDCC 54-21.3-04.1; NDAC article 108-01-01 | src:usa-nd:state-building-code-2026-book; src:usa-nd:ndac-108-01-01 | partially_verified |
| Elevator / Conveyance | ahj:usa-nd:state-electrical-board | North Dakota State Electrical Board | Conveyance registration, permits, certificates of operation, inspection, testing, and ASME-standard rulemaking authority | NDCC ch. 43-09 | src:usa-nd:ndcc-43-09 | partially_verified |

### 2.3 Authority Hierarchy Notes

The building-code framework is not a single statewide inspection program for all private construction. The state establishes and updates the State Building Code. Cities, townships, and counties that choose to administer and enforce a building code must use the State Building Code as the baseline, may amend it to local needs, and may designate an enforcement agency or use joint/private enforcement arrangements. Counties are the default local governmental enforcement option for areas not administered by a city or township.

Trade and safety codes are separated from the State Building Code chapter. Electrical, plumbing, and fire codes are expressly outside NDCC ch. 54-21.3. Electrical and plumbing are administered by their own boards. The operational fire code is administered under the State Fire Marshal rules. Conveyance regulation is tied to the State Electrical Board under the 2025 statutory provisions reviewed.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-nd:001 | ahj:usa-nd:commerce-dcs | administers_and_publishes | North Dakota State Building Code | src:usa-nd:commerce-building-codes; src:usa-nd:state-building-code-2026-book; src:usa-nd:ndac-108-01-01 | verified_core |
| edge:usa-nd:002 | ahj:usa-nd:commerce-dcs | coordinates_update_votes_with | ahj:usa-nd:building-code-advisory-committee and eligible local jurisdictions | src:usa-nd:state-building-code-2026-book; src:usa-nd:ndcc-54-21-3; src:usa-nd:ndac-108-01-01 | verified_core |
| edge:usa-nd:003 | ahj:usa-nd:local-building-jurisdiction | must_adopt_and_enforce_when_electing | North Dakota State Building Code | src:usa-nd:state-building-code-2026-book; src:usa-nd:ndcc-54-21-3; src:usa-nd:ndac-108-01-01 | verified_core |
| edge:usa-nd:004 | ahj:usa-nd:local-building-jurisdiction | may_amend_to_conform_to | local needs | src:usa-nd:state-building-code-2026-book; src:usa-nd:ndcc-54-21-3; src:usa-nd:ndac-108-01-01 | verified_core |
| edge:usa-nd:005 | ahj:usa-nd:local-building-jurisdiction | county_may_administer | areas not administered by city or township | src:usa-nd:state-building-code-2026-book; src:usa-nd:ndac-108-01-01 | verified_core |
| edge:usa-nd:006 | ahj:usa-nd:state-electrical-board | has_jurisdiction_over | electrical installations and conveyances | src:usa-nd:electrical-board-laws-rules; src:usa-nd:ndcc-43-09 | verified_core |
| edge:usa-nd:007 | ahj:usa-nd:state-electrical-board | city_may_inspect_within | corporate limits, subject to statutory conditions | src:usa-nd:ndcc-43-09 | partially_verified |
| edge:usa-nd:008 | ahj:usa-nd:state-plumbing-board | administers | North Dakota Plumbing Code | src:usa-nd:plumbing-board-laws-rules; src:usa-nd:ndac-62-01-01; src:usa-nd:ndac-62-03-1-01 | verified_core |
| edge:usa-nd:009 | ahj:usa-nd:state-fire-marshal | administers | State Fire Code | src:usa-nd:ndac-45-18-01; src:usa-nd:ndcc-18-01 | verified_core |
| edge:usa-nd:010 | ahj:usa-nd:state-and-local-public-owners | responsible_for_compliance_of | state buildings, local government buildings, and schools in the circumstances described by NDAC article 108-01-01 | src:usa-nd:state-building-code-2026-book; src:usa-nd:ndac-108-01-01 | verified_core |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | North Dakota State Building Code | International Building Code | 2024 | current | 2025-09-11 | 2026-01-01 | 2026-01-01 | 2026-01-01 for electing or previously electing local jurisdictions and covered public/school buildings | Electing city, township, or county must adopt and enforce the State Building Code; counties administer areas not administered by city or township. | src:usa-nd:commerce-building-codes; src:usa-nd:state-building-code-2026-book; src:usa-nd:ndcc-54-21-3; src:usa-nd:ndac-108-01-01 |
| Residential | North Dakota State Building Code | International Residential Code | 2024 | current | 2025-09-11 | 2026-01-01 | 2026-01-01 | 2026-01-01 for electing or previously electing local jurisdictions and covered public/school buildings | Same State Building Code local-election framework. | src:usa-nd:commerce-building-codes; src:usa-nd:state-building-code-2026-book; src:usa-nd:ndcc-54-21-3; src:usa-nd:ndac-108-01-01 |
| Existing Building / Rehabilitation | North Dakota State Building Code | International Existing Building Code | 2024 | current | 2025-09-11 | 2026-01-01 | 2026-01-01 | 2026-01-01 for electing or previously electing local jurisdictions and covered public/school buildings | Included in the 2026 State Building Code book; local enforcement follows State Building Code model. | src:usa-nd:commerce-building-codes; src:usa-nd:state-building-code-2026-book |
| Mechanical | North Dakota State Building Code | International Mechanical Code | 2024 | current | 2025-09-11 | 2026-01-01 | 2026-01-01 | 2026-01-01 for electing or previously electing local jurisdictions and covered public/school buildings | Same State Building Code local-election framework. | src:usa-nd:commerce-building-codes; src:usa-nd:state-building-code-2026-book; src:usa-nd:ndcc-54-21-3; src:usa-nd:ndac-108-01-01 |
| Plumbing | North Dakota Plumbing Code | Uniform Plumbing Code | 2018 | current | null | 2020-04-01 | 2020-04-01 | 2020-04-01 for plumbing governed by ND plumbing law and board rules | Plumbing work must meet or exceed NDAC article 62-03.1 and the North Dakota Plumbing Code; local ordinances or other legal requirements may also apply. | src:usa-nd:plumbing-board-laws-rules; src:usa-nd:ndac-62-01-01; src:usa-nd:ndac-62-03-1-01; src:usa-nd:ndac-62-03-1-02 |
| Fuel Gas | North Dakota State Building Code | International Fuel Gas Code | 2024 | current | 2025-09-11 | 2026-01-01 | 2026-01-01 | 2026-01-01 for electing or previously electing local jurisdictions and covered public/school buildings | Same State Building Code local-election framework. | src:usa-nd:commerce-building-codes; src:usa-nd:state-building-code-2026-book; src:usa-nd:ndcc-54-21-3; src:usa-nd:ndac-108-01-01 |
| Electrical | North Dakota Laws, Rules and Wiring Standards | National Electrical Code / NFPA 70 | 2023 | current_until_2026-07-01 | null | 2024-07-01 | 2024-07-01 | 2024-07-01 for electrical installations governed by State Electrical Board standards | 2023 NEC-based standards are current in the reviewed rules; official 2026 supplement indicates a 2026 NEC transition effective 2026-07-01. | src:usa-nd:electrical-board-laws-rules; src:usa-nd:ndac-24-1-06-01; src:usa-nd:ndac-24-1-2026-supplement; src:usa-nd:ndcc-43-09 |
| Energy | North Dakota State Building Code | International Energy Conservation Code and IRC energy provisions | 2024 | current | 2025-09-11 | 2026-01-01 | 2026-01-01 | 2026-01-01 for electing or previously electing local jurisdictions and covered public/school buildings | The 2026 State Building Code includes the 2024 IECC and retains IRC Chapter 11 energy provisions, subject to state amendments. | src:usa-nd:commerce-building-codes; src:usa-nd:state-building-code-2026-book |
| Fire - construction references | North Dakota State Building Code fire-safe construction provisions | IBC/IRC/related construction codes; Fire Code references State Building Code for fire-safe construction | 2024 State Building Code; Fire Code codified reference still needs 2026 codification check | current_with_caveat | 2025-09-11 for State Building Code | 2026-01-01 for State Building Code | 2026-01-01 | 2026-01-01 for building-code enforcement contexts; fire-code cross-reference should be checked after codification | State Fire Code is separate from State Building Code; current State Building Code excludes IFC. | src:usa-nd:state-building-code-2026-book; src:usa-nd:ndac-45-18-01; src:usa-nd:insurance-fire-rulemaking-2026 |
| Fire - operational / prevention code | North Dakota State Fire Code | International Fire Code | 2021 | current_codified | null | 2024-01-01 | 2024-01-01 | 2024-01-01 | Existing lawful conditions may continue unless the State Fire Marshal determines a distinct hazard; 2026 rulemaking copy for 2024 IFC requires codification check. | src:usa-nd:ndac-45-18-01; src:usa-nd:insurance-fire-rulemaking-2026 |
| Accessibility | North Dakota accessibility reference for covered buildings and facilities | 2010 ADA Standards for Accessible Design, 28 CFR parts 35 and 36 | 2010 | current_reference | null | null | null | null | Covered buildings and facilities subject to ADA must conform to the referenced federal standards; state-law statement process applies. | src:usa-nd:state-building-code-2026-book; src:usa-nd:ndac-108-01-01 |
| Elevator / Conveyance | North Dakota conveyance safety program | ASME standards to be implemented by State Electrical Board rule | edition_unresolved | partially_verified | null | 2025-07-01 statutory program baseline | 2025-07-01 | null | Conveyance registration, permit, operation certificate, inspection, and testing requirements appear in NDCC ch. 43-09; specific adopted ASME edition was not confirmed. | src:usa-nd:ndcc-43-09 |

### 3.2 Adoption Records

| Adoption Record ID | Code Family | Authority ID | State Code Name | Base Model Code | Edition | Adoption Date | Effective Date | Operative Date | Mandatory Date | Applicability Summary | Source IDs | Verification Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| adoption:usa-nd:building:2024-ibc | Building | ahj:usa-nd:commerce-dcs | North Dakota State Building Code | International Building Code | 2024 | 2025-09-11 | 2026-01-01 | 2026-01-01 | 2026-01-01 | Applies through the State Building Code framework for electing local jurisdictions and covered public/school buildings. | src:usa-nd:commerce-building-codes; src:usa-nd:state-building-code-2026-book | verified_core |
| adoption:usa-nd:residential:2024-irc | Residential | ahj:usa-nd:commerce-dcs | North Dakota State Building Code | International Residential Code | 2024 | 2025-09-11 | 2026-01-01 | 2026-01-01 | 2026-01-01 | Same State Building Code applicability model as building. | src:usa-nd:commerce-building-codes; src:usa-nd:state-building-code-2026-book | verified_core |
| adoption:usa-nd:existing-building:2024-iebc | Existing Building / Rehabilitation | ahj:usa-nd:commerce-dcs | North Dakota State Building Code | International Existing Building Code | 2024 | 2025-09-11 | 2026-01-01 | 2026-01-01 | 2026-01-01 | Included in 2026 State Building Code publication. | src:usa-nd:commerce-building-codes; src:usa-nd:state-building-code-2026-book | partially_verified |
| adoption:usa-nd:mechanical:2024-imc | Mechanical | ahj:usa-nd:commerce-dcs | North Dakota State Building Code | International Mechanical Code | 2024 | 2025-09-11 | 2026-01-01 | 2026-01-01 | 2026-01-01 | Same State Building Code applicability model as building. | src:usa-nd:commerce-building-codes; src:usa-nd:state-building-code-2026-book | verified_core |
| adoption:usa-nd:fuel-gas:2024-ifgc | Fuel Gas | ahj:usa-nd:commerce-dcs | North Dakota State Building Code | International Fuel Gas Code | 2024 | 2025-09-11 | 2026-01-01 | 2026-01-01 | 2026-01-01 | Same State Building Code applicability model as building. | src:usa-nd:commerce-building-codes; src:usa-nd:state-building-code-2026-book | verified_core |
| adoption:usa-nd:energy:2024-iecc | Energy | ahj:usa-nd:commerce-dcs | North Dakota State Building Code | International Energy Conservation Code and IRC energy provisions | 2024 | 2025-09-11 | 2026-01-01 | 2026-01-01 | 2026-01-01 | 2024 IECC and IRC energy provisions included with state amendments. | src:usa-nd:commerce-building-codes; src:usa-nd:state-building-code-2026-book | verified_core |
| adoption:usa-nd:plumbing:2018-upc | Plumbing | ahj:usa-nd:state-plumbing-board | North Dakota Plumbing Code | Uniform Plumbing Code | 2018 | null | 2020-04-01 | 2020-04-01 | 2020-04-01 | Applies to plumbing governed by NDCC ch. 43-18 and State Plumbing Board rules. | src:usa-nd:plumbing-board-laws-rules; src:usa-nd:ndac-62-03-1-01 | verified_core |
| adoption:usa-nd:electrical:2023-nec | Electrical | ahj:usa-nd:state-electrical-board | North Dakota Laws, Rules and Wiring Standards | National Electrical Code / NFPA 70 | 2023 | null | 2024-07-01 | 2024-07-01 | 2024-07-01 | Applies to electrical installations governed by State Electrical Board standards. | src:usa-nd:electrical-board-laws-rules; src:usa-nd:ndac-24-1-06-01; src:usa-nd:ndcc-43-09 | verified_core |
| adoption:usa-nd:fire:2021-ifc | Fire - operational / prevention code | ahj:usa-nd:state-fire-marshal | North Dakota State Fire Code | International Fire Code | 2021 | null | 2024-01-01 | 2024-01-01 | 2024-01-01 | Codified State Fire Code rule with exceptions and modifications. | src:usa-nd:ndac-45-18-01 | verified_core |
| adoption:usa-nd:accessibility:2010-ada | Accessibility | ahj:usa-nd:commerce-dcs | North Dakota accessibility reference for covered buildings and facilities | 2010 ADA Standards for Accessible Design | 2010 | null | null | null | null | State law references ADA standards for covered buildings and facilities and requires statements for covered work. | src:usa-nd:state-building-code-2026-book; src:usa-nd:ndac-108-01-01 | partially_verified |
| adoption:usa-nd:elevator:asme-unresolved | Elevator / Conveyance | ahj:usa-nd:state-electrical-board | North Dakota conveyance safety program | ASME standards | edition_unresolved | null | 2025-07-01 | 2025-07-01 | null | Statutory authority and program requirements verified; rule edition not confirmed. | src:usa-nd:ndcc-43-09 | unresolved_edition |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

North Dakota keeps several date concepts separate. The State Building Code's 2024 model-code package was adopted by eligible cities/counties and eligible Building Code Advisory Committee members on 2025-09-11 and became effective 2026-01-01. The local-government trigger is not simply the existence of a building anywhere in the state; it is whether a city, township, or county elects or has elected to administer and enforce a building code, plus special coverage for state/local government buildings and schools. The reviewed sources did not resolve permit-application, permit-issuance, or projects-in-progress transition rules.

The electrical program has a near-term transition item: current 2023 NEC-based standards were effective 2024-07-01, and an official 2026 administrative-code supplement indicates 2026 NEC-based standards effective 2026-07-01. The current codified State Fire Code source reviewed adopts the 2021 IFC effective 2024-01-01; the 2026 fire rulemaking copy indicates a 2024 IFC update path that should be checked against final codified NDAC.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-nd:building-2026-effective | Building, Residential, Existing Building, Mechanical, Fuel Gas, Energy | effective_date | 2026-01-01 | Current State Building Code edition package becomes effective. | Not resolved for projects already in process. | src:usa-nd:commerce-building-codes; src:usa-nd:state-building-code-2026-book | verified_core |
| date-rule:usa-nd:building-local-election | Building, Residential, Existing Building, Mechanical, Fuel Gas, Energy | local_adoption_trigger | ongoing; statutory framework baseline 1994-08-01 | City, township, or county elects or has elected to administer and enforce a building code. | Not resolved for private projects outside enforcing jurisdictions. | src:usa-nd:state-building-code-2026-book; src:usa-nd:ndcc-54-21-3; src:usa-nd:ndac-108-01-01 | verified_core |
| date-rule:usa-nd:public-buildings-schools | Building, Residential, Existing Building, Mechanical, Fuel Gas, Energy, Accessibility | owner_responsibility | ongoing | State agencies, local government owners, and public/private schools have compliance responsibilities in the circumstances described by state rules. | Not applicable as a prior-code grace rule. | src:usa-nd:state-building-code-2026-book; src:usa-nd:ndac-108-01-01 | verified_core |
| date-rule:usa-nd:plumbing-2018-upc | Plumbing | effective_date | 2020-04-01 | North Dakota Plumbing Code adoption of 2018 UPC under NDAC article 62-03.1. | Not resolved. | src:usa-nd:ndac-62-03-1-01 | verified_core |
| date-rule:usa-nd:electrical-2023-nec | Electrical | effective_date | 2024-07-01 | Current State Electrical Board standards based on 2023 NEC became effective. | Not resolved. | src:usa-nd:electrical-board-laws-rules; src:usa-nd:ndac-24-1-06-01 | verified_core |
| date-rule:usa-nd:electrical-2026-nec | Electrical | future_effective_date | 2026-07-01 | Official 2026 administrative-code supplement indicates updated 2026 NEC-based standards. | Check final application and any grace period after the effective date. | src:usa-nd:ndac-24-1-2026-supplement | watch |
| date-rule:usa-nd:fire-2021-ifc | Fire - operational / prevention code | effective_date | 2024-01-01 | Codified State Fire Code adopts 2021 IFC with North Dakota amendments. | Existing conditions that were lawful may continue unless the State Fire Marshal finds a distinct hazard. | src:usa-nd:ndac-45-18-01 | verified_core |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Electrical | 2026 NEC-based State Electrical Board standards | null | null | 2026-07-01 | 2026-07-01 | 2026-07-01 | scheduled_check_needed | src:usa-nd:ndac-24-1-2026-supplement | Official supplement reviewed; confirm final codified NDAC after effective date. |
| Fire - operational / prevention code | 2024 International Fire Code with North Dakota modifications | null | null | null | null | null | codification_check_needed | src:usa-nd:insurance-fire-rulemaking-2026 | Rulemaking copy is treated as a pending or proposed source, not as current codified law in this report. |
| Building, Residential, Existing Building, Mechanical, Fuel Gas, Energy | none identified beyond current 2024 model-code package | null | null | null | null | null | monitor | src:usa-nd:commerce-building-codes; src:usa-nd:state-building-code-2026-book | Current State Building Code package already effective 2026-01-01. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-nd:ifc-ipmc-excluded | Building; Fire; Property Maintenance | International Fire Code and International Property Maintenance Code | Local jurisdiction wants IFC or IPMC coverage | The State Building Code book states that the State Building Code does not include IPMC or IFC; those codes must be adopted separately by each city, county, or township. | src:usa-nd:state-building-code-2026-book | verified_core |
| applicability-rule:usa-nd:appendices-excluded | Building, Residential, Existing Building, Mechanical, Fuel Gas, Energy | Model-code appendices | Appendix chapter is not specifically adopted | Appendix chapters are not part of the State Building Code unless specifically adopted as a state amendment. | src:usa-nd:state-building-code-2026-book; src:usa-nd:ndac-108-01-01 | verified_core |
| applicability-rule:usa-nd:sprinkler-limitation | Residential | One-family and two-family dwellings | State or local building code would require a fire sprinkler system in a one-family or two-family dwelling | NDCC ch. 54-21.3 limits state and local building codes from requiring fire sprinkler systems in one-family or two-family dwellings. | src:usa-nd:state-building-code-2026-book; src:usa-nd:ndcc-54-21-3 | verified_core |
| applicability-rule:usa-nd:fire-existing-conditions | Fire - operational / prevention code | Existing conditions | Existing condition complied with law when established | Existing legal noncompliant conditions may continue unless the State Fire Marshal finds a distinct hazard to life or property. | src:usa-nd:ndac-45-18-01 | verified_core |
| applicability-rule:usa-nd:local-directory-caveat | Local enforcement | Jurisdiction lookup | User relies on Commerce local-code directory | Commerce's directory should be verified with the local jurisdiction and may not list every jurisdiction that has adopted a building code. | src:usa-nd:state-building-code-2026-book | verified_core |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** consolidated State Building Code book for the 2024 model-code package, with state amendments published by the Division of Community Services. The amendment process uses the Building Code Advisory Committee and eligible city/county voting structure described in statute and rule.

**Where amendments are published:** North Dakota Department of Commerce, Division of Community Services State Building Code book, with supporting administrative rules in NDAC article 108-01-01 and statutory authority in NDCC ch. 54-21.3.

**Amendment parsing status:** core_high_impact_parsed

### 5.2 State Amendment Sources

| Amendment Source ID | Code Families | Publication Path | Coverage | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| amendment-source:usa-nd:2026-state-building-code-book | Building, Residential, Existing Building, Mechanical, Fuel Gas, Energy, Accessibility-related state notes | Commerce/DCS State Building Code book | 2024 IBC, IRC, IMC, IFGC, IECC, IEBC and North Dakota amendments effective 2026-01-01 | src:usa-nd:state-building-code-2026-book; src:usa-nd:commerce-building-codes | verified_core |
| amendment-source:usa-nd:ndac-108-process | State Building Code amendment procedure | NDAC article 108-01-01 | Voting, publication, limitation, and local enforcement framework | src:usa-nd:ndac-108-01-01; src:usa-nd:ndcc-54-21-3 | verified_core |
| amendment-source:usa-nd:state-fire-code-2024-current | Fire - operational / prevention code | NDAC chapter 45-18-01 | 2021 IFC with North Dakota modifications, effective 2024-01-01 | src:usa-nd:ndac-45-18-01 | verified_core |
| amendment-source:usa-nd:state-fire-code-2026-rulemaking | Fire - operational / prevention code | Insurance Department rulemaking PDF | Proposed or pending 2024 IFC update; not treated as final current law | src:usa-nd:insurance-fire-rulemaking-2026 | watch |
| amendment-source:usa-nd:electrical-board-2024 | Electrical | State Electrical Board laws, rules, and wiring standards; NDAC 24.1-06-01 | 2023 NEC-based electrical standards effective 2024-07-01 | src:usa-nd:electrical-board-laws-rules; src:usa-nd:ndac-24-1-06-01 | verified_core |
| amendment-source:usa-nd:plumbing-board-2018-upc | Plumbing | NDAC article 62-03.1 | 2018 UPC as North Dakota Plumbing Code, with board administration | src:usa-nd:ndac-62-03-1-01 | verified_core |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Topic | Summary | Source IDs | Verification Status |
| --- | --- | --- | --- | --- | --- |
| amendment:usa-nd:state-building-code-excludes-ifc-ipmc | Building; Fire; Property Maintenance | IFC and IPMC exclusion from State Building Code | The State Building Code book states that the State Building Code does not include the International Fire Code or International Property Maintenance Code; each city, county, or township must separately adopt those codes if desired. | src:usa-nd:state-building-code-2026-book | verified_core |
| amendment:usa-nd:appendix-exclusion | Building, Residential, Existing Building, Mechanical, Fuel Gas, Energy | Model-code appendices | Appendix chapters are not part of the State Building Code unless specifically adopted as state amendments. | src:usa-nd:state-building-code-2026-book; src:usa-nd:ndac-108-01-01 | verified_core |
| amendment:usa-nd:residential-sprinkler-limitation | Residential | One-family and two-family dwelling fire sprinklers | State and local building codes may not require fire sprinkler systems in one-family or two-family dwellings under the cited statute. | src:usa-nd:state-building-code-2026-book; src:usa-nd:ndcc-54-21-3 | verified_core |
| amendment:usa-nd:ada-nonamendability | Accessibility | ADA federal standards | Federal ADA standards referenced by state law cannot be amended through the State Building Code process. | src:usa-nd:state-building-code-2026-book; src:usa-nd:ndac-108-01-01 | verified_core |
| amendment:usa-nd:fire-group-r-sprinkler-exception | Fire - operational / prevention code | Group R sprinklers | Current State Fire Code amendments include a Group R sprinkler exception tied to the fire-code text. | src:usa-nd:ndac-45-18-01 | partially_verified |
| amendment:usa-nd:electrical-local-more-stringent | Electrical | More stringent local ordinances | Cities may make more stringent electrical requirements by ordinance, subject to the limitation that such ordinances apply only to individuals licensed by the board. | src:usa-nd:ndcc-43-09 | verified_core |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-nd"
  model: "hybrid_opt_in_local_enforcement"
  enforcing_entities:
    - "city or township within its jurisdictional area, when it elects to administer and enforce a building code"
    - "county in areas not administered by a city or township, when the county elects to administer and enforce a building code"
    - "designated enforcement agency, joint enforcement arrangement, or private enforcement arrangement authorized by governing body agreement"
    - "state agencies, local government owners, and public/private schools for their own buildings in the circumstances described in NDAC article 108-01-01"
  required_officials:
    - "local building official or designated enforcement agency where local jurisdiction administers the State Building Code"
    - "State Electrical Board inspectors or city electrical inspectors where authorized"
    - "State Plumbing Board executive director and inspectors for plumbing-code administration"
    - "State Fire Marshal or authorized representative for State Fire Code enforcement"
  state_reserved_activities:
    - "State Building Code rule administration, publication, and state amendment process through Commerce/DCS"
    - "electrical licensing, standards, and inspection jurisdiction through the State Electrical Board"
    - "plumbing-code administration through the State Plumbing Board"
    - "State Fire Code administration through the State Fire Marshal"
    - "conveyance registration, permits, certificates, inspection, testing, and ASME-rule implementation through the State Electrical Board"
  source_ids:
    - "src:usa-nd:state-building-code-2026-book"
    - "src:usa-nd:ndcc-54-21-3"
    - "src:usa-nd:ndac-108-01-01"
    - "src:usa-nd:ndcc-43-09"
    - "src:usa-nd:ndac-62-03-1-01"
    - "src:usa-nd:ndac-45-18-01"
  verification_status: "verified_core"
  confidence: 0.82
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-nd"
  model: "local_building_amendments_allowed_to_conform_to_local_needs"
  applies_to_code_families:
    - "Building"
    - "Residential"
    - "Existing Building / Rehabilitation"
    - "Mechanical"
    - "Fuel Gas"
    - "Energy"
  approval_required: "state approval requirement not found in reviewed sources"
  approving_authority_id: null
  filing_required: "not confirmed"
  registry_exists: "not found"
  registry_source_ids: []
  legal_basis_source_ids:
    - "src:usa-nd:state-building-code-2026-book"
    - "src:usa-nd:ndcc-54-21-3"
    - "src:usa-nd:ndac-108-01-01"
  related_trade_notes:
    electrical: "Cities may make more stringent electrical requirements by ordinance, subject to statutory limits."
    plumbing: "NDAC 62-03.1-02 preserves the possibility of local ordinances or other legal requirements, but the scope was not fully classified."
    fire: "Local fire-code adoption or amendment scope was not fully classified; the State Building Code excludes IFC and IPMC, and State Fire Code authority is separate."
  verification_status: "partially_verified"
  confidence: 0.70
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement and local amendment authority are related but not identical in North Dakota. A city, township, or county may elect to administer and enforce a building code; once it does, it must adopt and enforce the State Building Code. That same local-government category may amend the State Building Code to conform to local needs. The sources reviewed do not confirm a central state registry for those local building amendments, so jurisdiction-specific verification remains necessary for production use.

Electrical, plumbing, and fire should not be folded into the building-code amendment model. Electrical has its own statewide board and a statutory path for more-stringent city ordinances. Plumbing is administered by the State Plumbing Board, with local ordinances or other legal requirements preserved by rule. The operational fire code is administered through the State Fire Marshal and must be evaluated separately from building-code enforcement.

### 6.4 Known Local Amendment Registries

| Registry ID | Registry Name | Scope | Source IDs | Status | Notes |
| --- | --- | --- | --- | --- | --- |
| registry:usa-nd:commerce-local-enforcement-directory | Commerce local code enforcement directory | Local enforcement contacts, not a verified local-amendment registry | src:usa-nd:state-building-code-2026-book; src:usa-nd:commerce-building-codes | limited | Directory is useful for AHJ discovery, but Commerce cautions that jurisdictions not listed may have adopted a building code and that users should verify locally. |
| registry:usa-nd:local-building-amendments | Statewide local building amendment registry | Local amendments to State Building Code | none | not_found | No statewide registry was confirmed from the reviewed sources. |

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendment set was parsed. Production use for a specific address should verify the city, township, county, and any local fire-code adoption directly with the AHJ.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: core_model_started

Jurisdiction stack:

```text
Address
  -> State of North Dakota
  -> County
  -> Municipality or township, if present
  -> Determine whether city/township administers and enforces the State Building Code
  -> If not administered by city/township, determine whether county administers and enforces in that area
  -> Local building AHJ or designated/joint/private enforcement agency, if local enforcement exists
  -> State Building Code adoption records for current model-code package
  -> State Fire Marshal / State Fire Code for operational fire-prevention requirements
  -> State Electrical Board and any authorized city electrical inspection program
  -> State Plumbing Board and any local ordinance or district-health/private-sewage requirement
  -> State Electrical Board conveyance program, if conveyance equipment is present
  -> Applicable local amendments or separate IFC/IPMC local adoptions, if confirmed
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | not selected | none | statewide | unknown | pending |
| County | not selected | none | statewide | unknown | pending |
| Municipality | not selected | none | statewide | unknown | pending |
| Township | not selected | none | statewide | unknown | pending |
| Fire District | not selected | none | statewide | unknown | pending |
| Special District | not selected | none | statewide | unknown | pending |
| Local code enforcement contacts | Commerce local code enforcement directory | src:usa-nd:commerce-building-codes; src:usa-nd:state-building-code-2026-book | listed jurisdictions only | unknown | useful_with_caveat |

### 7.3 AHJ Contact Data

No address-level AHJ contact data was populated. The Commerce local-code enforcement directory should be used as an initial contact source, but local confirmation is required because the state source cautions that unlisted jurisdictions may still have adopted a building code.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Source Name | Publisher / Custodian | Source Type | URL | Key Coverage | Status |
| --- | --- | --- | --- | --- | --- | --- |
| src:usa-nd:commerce-building-codes | Building Codes | North Dakota Department of Commerce | agency page | https://www.commerce.nd.gov/community-services/building-codes | Current State Building Code summary, 2024 model-code adoption statement, 2026 effective date, links to code book and local enforcement directory | official_current |
| src:usa-nd:state-building-code-2026-book | 2026 North Dakota State Building Code Book | North Dakota Department of Commerce, Division of Community Services | official PDF | https://www.medialibrary.nd.gov/assetbank-nd/assetfile/156311.pdf | 2026 State Building Code, code editions, history, statutory excerpts, administrative rules, local enforcement directory, amendment text | official_current |
| src:usa-nd:ndcc-54-21-3 | North Dakota Century Code Chapter 54-21.3, State Building Code | North Dakota Legislative Branch | statute PDF | https://ndlegis.gov/cencode/t54c21-3.pdf | State Building Code authority, local adoption/enforcement, local amendments, trade-code exclusions | official_current |
| src:usa-nd:ndac-108-01-01 | North Dakota Administrative Code Chapter 108-01-01, State Building Code | North Dakota Legislative Branch | administrative rule PDF | https://ndlegis.gov/prod/acdata/pdf/108-01-01.pdf | Scope, administration, local enforcement, amendment process, limitations, publication requirements | official_current |
| src:usa-nd:ndac-45-18-01 | North Dakota Administrative Code Chapter 45-18-01, State Fire Code | North Dakota Legislative Branch | administrative rule PDF | https://ndlegis.gov/information/acdata/pdf/45-18-01.pdf | State Fire Marshal authority, State Fire Code scope, 2021 IFC adoption with North Dakota modifications, 2024-01-01 effective history | official_current |
| src:usa-nd:insurance-fire-rulemaking-2026 | Article 45-18 Fire Marshal Rulemaking PDF | North Dakota Insurance Department / State Fire Marshal rulemaking | rulemaking PDF | https://www.insurance.nd.gov/sites/www/files/documents/Communications/Rulemaking/2026/Article%2045-18%20-%20Fire%20Marshal.pdf | 2026 fire-code rulemaking/redline indicating 2024 IFC update path and State Building Code cross-reference update | official_rulemaking_caveated |
| src:usa-nd:ndcc-18-01 | North Dakota Century Code Chapter 18-01, Fire Prevention | North Dakota Legislative Branch | statute PDF | https://ndlegis.gov/cencode/t18c01.pdf | State Fire Marshal statutory authority and delegation/enforcement framework | official_current |
| src:usa-nd:electrical-board-laws-rules | Laws, Rules and Wiring Standards | North Dakota State Electrical Board | agency page | https://www.ndseb.com/laws-rules-wiring-standards/ | Electrical Board standards page, 2023 laws/rules effective 2024-07-01, NEC/NFPA references | official_current |
| src:usa-nd:ndac-24-1-06-01 | North Dakota Administrative Code Chapter 24.1-06-01, Wiring Standards | North Dakota Legislative Branch | administrative rule PDF | https://ndlegis.gov/information/acdata/pdf/24.1-06-01.pdf | Electrical wiring standards, 2023 NEC references, effective-history information | official_current |
| src:usa-nd:ndac-24-1-2026-supplement | North Dakota Administrative Code 2026 Supplement for Electrical Board Rules | North Dakota Legislative Branch | administrative-code supplement PDF | https://ndlegis.gov/admcode/arc202607400.pdf | 2026 NEC transition language and July 1, 2026 effective-history markers | official_future_caveated |
| src:usa-nd:ndcc-43-09 | North Dakota Century Code Chapter 43-09, Electricians | North Dakota Legislative Branch | statute PDF | https://ndlegis.gov/cencode/t43c09.pdf | State Electrical Board jurisdiction, local electrical ordinances, inspection authority, conveyance statutory program | official_current |
| src:usa-nd:plumbing-board-laws-rules | Plumbing Laws and Rules | North Dakota State Plumbing Board | agency page | https://www.ndplumbingboard.com/laws-rules/ | Plumbing Board statutory background and links to NDAC plumbing rules | official_current |
| src:usa-nd:ndac-62-01-01 | North Dakota Administrative Code Chapter 62-01-01, Organization of Board | North Dakota Legislative Branch | administrative rule PDF | https://ndlegis.gov/information/acdata/pdf/62-01-01.pdf | State Plumbing Board organization and rulemaking/code authority | official_current |
| src:usa-nd:ndac-62-03-1-01 | North Dakota Administrative Code Chapter 62-03.1-01, General Provisions | North Dakota Legislative Branch | administrative rule PDF | https://ndlegis.gov/information/acdata/pdf/62-03.1-01.pdf | 2018 UPC adoption as North Dakota Plumbing Code and board administration | official_current |
| src:usa-nd:ndac-62-03-1-02 | North Dakota Administrative Code Chapter 62-03.1-02, Local Ordinances | North Dakota Legislative Branch | administrative rule PDF | https://ndlegis.gov/information/acdata/pdf/62-03.1-02.pdf | Plumbing interaction with local ordinances and other legal requirements | official_current |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-nd:state-building-code-2026-book | consolidated_publication | The code book is an official Commerce/DCS consolidation with statutory and rule excerpts, but individual statute and NDAC sources remain the controlling legal sources for legal interpretation. | use_with_primary_crosscheck |
| src:usa-nd:insurance-fire-rulemaking-2026 | rulemaking_or_redline | The fire rulemaking PDF appears to show edits toward 2024 IFC and 2026 State Building Code references. It was not treated as final codified NDAC in this report. | monitor_until_codified |
| src:usa-nd:ndac-24-1-2026-supplement | future_supplement | The electrical supplement contains redline-like text and future effective markers. Use it to flag the 2026-07-01 transition, then confirm final codified rule text. | monitor_until_effective_and_codified |
| src:usa-nd:commerce-building-codes | agency_summary | The Commerce page is official and current-looking, but the detailed legal basis should be cross-checked against the code book, NDCC, and NDAC sources. | use_with_primary_crosscheck |
| src:usa-nd:state-building-code-2026-book | directory_caveat | The local enforcement directory is not guaranteed to list every jurisdiction that has adopted a code; local verification remains required. | verify_locally |

### 8.3 Supplemental Sources

None used for report facts. Non-official sources were intentionally avoided for the core authority and adoption fields.

### 8.4 Source Extraction Metadata

| Extraction ID | Source IDs | Extracted Fields | Extraction Date | Method | Notes |
| --- | --- | --- | --- | --- | --- |
| extract:usa-nd:001 | src:usa-nd:commerce-building-codes; src:usa-nd:state-building-code-2026-book | Current State Building Code editions, adoption date, effective date, local enforcement summary, amendment publication path | 2026-06-26 | web review of official agency page and official PDF text/screenshot | Core State Building Code facts verified. |
| extract:usa-nd:002 | src:usa-nd:ndcc-54-21-3; src:usa-nd:ndac-108-01-01 | Statutory and administrative authority, local enforcement, local amendments, exclusions | 2026-06-26 | official ND legislative PDFs | Used to separate state authority, local enforcement, and trade-code exclusions. |
| extract:usa-nd:003 | src:usa-nd:electrical-board-laws-rules; src:usa-nd:ndac-24-1-06-01; src:usa-nd:ndac-24-1-2026-supplement; src:usa-nd:ndcc-43-09 | Electrical authority, NEC edition status, future NEC transition, local electrical ordinance note, conveyance program | 2026-06-26 | official agency page and ND legislative PDFs | 2026 transition flagged for follow-up after codification/effective date. |
| extract:usa-nd:004 | src:usa-nd:plumbing-board-laws-rules; src:usa-nd:ndac-62-01-01; src:usa-nd:ndac-62-03-1-01; src:usa-nd:ndac-62-03-1-02 | Plumbing authority, 2018 UPC adoption, local ordinance interaction | 2026-06-26 | official agency page and ND legislative PDFs | Plumbing code date derived from NDAC effective-history text. |
| extract:usa-nd:005 | src:usa-nd:ndac-45-18-01; src:usa-nd:insurance-fire-rulemaking-2026; src:usa-nd:ndcc-18-01 | Fire authority, current codified State Fire Code, 2026 fire rulemaking watch item | 2026-06-26 | official NDAC and agency rulemaking PDF | Current fire row remains 2021 IFC until codified 2024 IFC update is confirmed. |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 1.00 | none | Status reflects source-backed core authority and adoption fields with unresolved fire/elevator/local registry gaps. |
| report | risk.overall_confidence | 0.68 | verified | 1.00 | none | Confidence is moderate because core codes are source-backed but not all local and trade details are complete. |
| ahj:usa-nd:commerce-dcs | authority.name | North Dakota Department of Commerce, Division of Community Services | verified_core | 0.90 | src:usa-nd:commerce-building-codes; src:usa-nd:state-building-code-2026-book; src:usa-nd:ndac-108-01-01 | Primary State Building Code administration. |
| adoption:usa-nd:building:2024-ibc | edition/effective_date | 2024 IBC / 2026-01-01 | verified_core | 0.92 | src:usa-nd:commerce-building-codes; src:usa-nd:state-building-code-2026-book | Current State Building Code package. |
| adoption:usa-nd:residential:2024-irc | edition/effective_date | 2024 IRC / 2026-01-01 | verified_core | 0.92 | src:usa-nd:commerce-building-codes; src:usa-nd:state-building-code-2026-book | Current State Building Code package. |
| adoption:usa-nd:mechanical:2024-imc | edition/effective_date | 2024 IMC / 2026-01-01 | verified_core | 0.92 | src:usa-nd:commerce-building-codes; src:usa-nd:state-building-code-2026-book | Current State Building Code package. |
| adoption:usa-nd:fuel-gas:2024-ifgc | edition/effective_date | 2024 IFGC / 2026-01-01 | verified_core | 0.92 | src:usa-nd:commerce-building-codes; src:usa-nd:state-building-code-2026-book | Current State Building Code package. |
| adoption:usa-nd:energy:2024-iecc | edition/effective_date | 2024 IECC / 2026-01-01 | verified_core | 0.88 | src:usa-nd:commerce-building-codes; src:usa-nd:state-building-code-2026-book | Includes IRC energy-provision note. |
| adoption:usa-nd:plumbing:2018-upc | edition/effective_date | 2018 UPC / 2020-04-01 | verified_core | 0.84 | src:usa-nd:ndac-62-03-1-01 | Plumbing adoption verified from NDAC. |
| adoption:usa-nd:electrical:2023-nec | edition/effective_date | 2023 NEC / 2024-07-01 | verified_core | 0.88 | src:usa-nd:electrical-board-laws-rules; src:usa-nd:ndac-24-1-06-01 | Future 2026 NEC transition separately flagged. |
| adoption:usa-nd:fire:2021-ifc | edition/effective_date | 2021 IFC / 2024-01-01 | verified_core | 0.75 | src:usa-nd:ndac-45-18-01 | Current codified fire rule reviewed; 2026 rulemaking not treated as final. |
| local-enforcement:usa-nd | model | hybrid_opt_in_local_enforcement | verified_core | 0.82 | src:usa-nd:state-building-code-2026-book; src:usa-nd:ndcc-54-21-3; src:usa-nd:ndac-108-01-01 | Local enforcement separated from local amendment authority. |
| local-amendment-rule:usa-nd | model | local_building_amendments_allowed_to_conform_to_local_needs | partially_verified | 0.70 | src:usa-nd:state-building-code-2026-book; src:usa-nd:ndcc-54-21-3; src:usa-nd:ndac-108-01-01 | Registry, filing, and approval details still unresolved. |
| adoption:usa-nd:elevator:asme-unresolved | base_model_code | ASME standards, edition unresolved | unresolved_edition | 0.55 | src:usa-nd:ndcc-43-09 | Statutory program verified; adopted edition not confirmed. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Every body source ID is listed in section 8. |
| All authority IDs resolve | pass | Authority IDs used in the report are defined in sections 2 and 6. |
| All current code families have adoption records | pass | Adoption rows are present for all matrix families; unresolved edition/date details are explicit. |
| Building and operational fire code are separated | pass | State Building Code excludes IFC; State Fire Code is tracked separately. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Dates are separate columns, with null used where a date was not confirmed. |
| Effective dates are valid ISO dates | pass | Populated dates use YYYY-MM-DD format; unresolved dates remain null. |
| No impossible date sequences | pass | No adoption/effective/operative/mandatory sequence conflict was introduced. |
| Transition rules have explicit trigger conditions | pass | Local-election, public-owner, fire-existing-condition, and electrical-transition triggers are identified. |
| Permit-date logic is captured where applicable | fail | Permit application, permit issuance, and projects-in-progress rules were not found in the reviewed core sources. |
| Local enforcement model classified | pass | Classified as hybrid opt-in local enforcement. |
| Local amendment rule classified | pass | Building-code local amendment rule classified; registry and filing details remain unresolved. |
| AHJ confirmation metadata present | fail | Address-level AHJ contacts and boundary data were not populated. |
| Official-source caveats captured | pass | Caveats recorded for consolidated code book, rulemaking/redline sources, future electrical supplement, and local directory. |
| Fire-code 2024 update final status resolved | fail | 2026 fire rulemaking source requires final codified NDAC check. |
| Elevator / conveyance code edition resolved | fail | ASME edition and detailed administrative rules were not verified. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-nd:001 | high | fire-code update | Determine whether the 2026 Fire Marshal rulemaking has been finally adopted and codified, and whether North Dakota has moved from 2021 IFC to 2024 IFC. | Recheck NDAC 45-18-01 and Insurance Department rulemaking docket after effective/codification milestones. | null | null | open |
| issue:usa-nd:002 | medium | local amendment registry | Confirm whether any state filing, approval, or registry process exists for local building-code amendments. | Review Commerce guidance, local ordinance filing statutes, and sample local ordinances; contact DCS if needed. | null | null | open |
| issue:usa-nd:003 | medium | permit transition rules | Identify whether North Dakota or local AHJs use permit-application date, permit-issuance date, design contract date, grace periods, or concurrent-use rules for the 2026 State Building Code transition. | Review 2026 State Building Code book sections not parsed here and local AHJ transition notices. | null | null | open |
| issue:usa-nd:004 | medium | elevator/conveyance standards | Confirm the specific ASME standards and editions adopted for conveyance equipment. | Review State Electrical Board conveyance rules and any implementing guidance after the 2025 statutory changes. | null | null | open |
| issue:usa-nd:005 | medium | AHJ boundaries and contacts | Build an address-level resolver for city/township/county building AHJ, local fire AHJ, electrical inspection, and plumbing enforcement contacts. | Combine official boundary data with Commerce local enforcement directory and local confirmation. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-nd:commerce-building-codes | src:usa-nd:commerce-building-codes | html_diff | monthly | Commerce posts new State Building Code book, directory, schedule, or amendment instructions | 2026-06-26 | active |
| watch:usa-nd:state-building-code-book | src:usa-nd:state-building-code-2026-book | pdf_diff | monthly | Code book revised or replacement PDF issued | 2026-06-26 | active |
| watch:usa-nd:ndac-108 | src:usa-nd:ndac-108-01-01 | pdf_diff | quarterly | State Building Code administrative rule changes | 2026-06-26 | active |
| watch:usa-nd:fire-code | src:usa-nd:ndac-45-18-01 | pdf_diff | monthly | State Fire Code codifies 2024 IFC or changes amendments | 2026-06-26 | active |
| watch:usa-nd:fire-rulemaking-2026 | src:usa-nd:insurance-fire-rulemaking-2026 | docket_check | monthly | Fire Marshal 2026 rulemaking final adoption, effective date, or codification update | 2026-06-26 | active |
| watch:usa-nd:electrical-2026 | src:usa-nd:ndac-24-1-2026-supplement | pdf_diff | monthly | 2026 NEC transition becomes codified/current or grace-period guidance is issued | 2026-06-26 | active |
| watch:usa-nd:plumbing-code | src:usa-nd:ndac-62-03-1-01 | pdf_diff | quarterly | Plumbing Code edition or amendments change | 2026-06-26 | active |
| watch:usa-nd:electrical-statute | src:usa-nd:ndcc-43-09 | statute_check | quarterly | Electrical or conveyance authority changes | 2026-06-26 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-26 | Populated North Dakota report from baseline draft and upgraded status to partially_verified | state-report:usa-nd | src:usa-nd:commerce-building-codes; src:usa-nd:state-building-code-2026-book; src:usa-nd:ndcc-54-21-3; src:usa-nd:ndac-108-01-01; src:usa-nd:ndac-45-18-01; src:usa-nd:electrical-board-laws-rules; src:usa-nd:ndac-24-1-06-01; src:usa-nd:ndcc-43-09; src:usa-nd:plumbing-board-laws-rules; src:usa-nd:ndac-62-03-1-01 | ChatGPT | Core authority, adoption matrix, source registry, QA, and open issues populated from official sources. |
