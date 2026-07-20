---
state:
  state_id: "US-MT"
  name: "Montana"
  abbreviation: "MT"
report:
  report_id: "state-report:usa-mt"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-26"
  last_verified: null
  reviewed_by: null
risk:
  overall_confidence: 0.68 # 0.00 - 1.00
  risk_flags:
    - "official_rules_site_dynamic_html_requires_manual_capture_for_some_arm_pages"
    - "adoption_dates_separate_from_effective_dates_unresolved_for_current_cycle"
    - "fire_code_effective_date_unresolved"
    - "local_amendment_registry_not_verified"
    - "pending_2026_code_update_is_draft_only"
  open_questions_count: 6

---

# State Building Code Authority Report: Montana

## 1. Executive Summary

- **Authority model:** Montana uses a statewide state-building-code model administered by the Montana Department of Labor & Industry, Business Standards Division, Building Codes Program / Building and Commercial Measurements Bureau. The department is the state-level building-code rulemaking authority; the Building Codes Council is advisory, not the primary adopting authority.

- **Statewide code status:** The current DLI code list identifies the 2021 IBC, 2021 IRC, 2021 IEBC, 2021 IMC, 2021 IFGC, 2021 IECC, 2021 UPC, 2020 NEC, 2017 ICC A117.1, 2021 ISPSC, 2021 IWUIC, and listed elevator/boiler standards as adopted and amended through ARM Title 24, Chapter 301, with most listed effective dates of 2022-06-11.

- **Local enforcement model:** Cities, counties, and towns may choose to enforce local building, electrical, plumbing, and mechanical codes, in whole or in part, only through a certified local code-enforcement program. Areas without a certified local program remain under the State Building Codes Bureau / Building Codes Program. State-owned agency buildings remain under DLI jurisdiction for plan review, permitting, and inspections.

- **Local amendment posture:** Local building codes may include only codes adopted by the Building Codes Program. Certified local governments may make, amend, and repeal rules for administration/enforcement and fee collection, but this pass did not verify a statewide registry of local technical amendments.

- **Known transition periods or pending changes:** ARM 24.301.146 contains a permit-issued transition rule for new code requirements; projects with a legal building permit already issued are not required to meet later new requirements, while later alterations/remodels use requirements in effect at the time of permit issuance for that new work. DLI has also posted a draft MAR notice titled “2026 Updates to the State Building Code” proposing to update several model-code editions, including 2024 I-Codes/UPC and 2023 NEC; no final adoption or effective date was verified in this pass.

- **Production readiness:** limited_production_for_statewide_core_fields

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: DLI administers Title 50, chapter 60, parts 1 through 7 and must adopt
    rules that constitute the state building code.
  confidence: 0.9
  source_ids:
  - src:usa-mt:mca-50-60-103
  - src:usa-mt:mca-50-60-203
- topic: Sole state building-regulation authority
  finding: DLI is the only state agency that may promulgate building regulations,
    with an express Department of Justice exception for regulations relating to use
    of buildings and installation of equipment.
  confidence: 0.9
  source_ids:
  - src:usa-mt:mca-50-60-202
- topic: Primary building code edition
  finding: DLI lists the 2021 IBC as adopted and effective 2022-06-11, amended through
    ARM Title 24, Chapter 301.
  confidence: 0.85
  source_ids:
  - src:usa-mt:dli-current-codes
  - src:usa-mt:arm-24-301-current-rules
- topic: Electrical code edition
  finding: DLI lists the 2020 NEC as adopted and effective 2022-06-11, with a technical
    advisory issued 2022-08-12.
  confidence: 0.8
  source_ids:
  - src:usa-mt:dli-current-codes
- topic: Fire code authority
  finding: The Department of Justice Fire Prevention and Investigation Section / State
    Fire Marshal has statewide fire-prevention authority and ARM 23.12.601 adopts
    the 2021 IFC; the effective date of the IFC rule was not captured.
  confidence: 0.7
  source_ids:
  - src:usa-mt:mca-50-3-102
  - src:usa-mt:doj-fpi
  - src:usa-mt:arm-23-12-601
- topic: Local enforcement
  finding: Local governments may enforce selected code programs only if certified;
    otherwise the State Building Codes Bureau has jurisdiction.
  confidence: 0.9
  source_ids:
  - src:usa-mt:dli-building-codes-program
  - src:usa-mt:dli-certified-local-programs
  - src:usa-mt:mca-50-60-302
- topic: Local amendment scope
  finding: Local codes may include only codes adopted by the Building Codes Program;
    local authority verified in this pass is primarily adoption/enforcement/administration,
    not independent technical-code substitution.
  confidence: 0.75
  source_ids:
  - src:usa-mt:dli-certified-local-programs
  - src:usa-mt:mca-50-60-106
- topic: Effective / transition rule
  finding: DLI lists effective dates; ARM 24.301.146 contains a permit-issued transition
    rule, but adoption dates and mandatory dates were not separately verified.
  confidence: 0.7
  source_ids:
  - src:usa-mt:dli-current-codes
  - src:usa-mt:arm-24-301-146-permit-transition
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-mt:dli-building-codes-program |
| Authority name | Montana Department of Labor & Industry, Business Standards Division, Building Codes Program / Building and Commercial Measurements Bureau |
| Authority type | state_agency_program |
| Legal basis | MCA 50-60-103; MCA 50-60-202; MCA 50-60-203 |
| Role | Administers Montana building-construction standards; adopts the state building code by rule; issues state permits and inspections where the state retains jurisdiction; certifies local government code-enforcement programs. |
| Enforcement model | mixed_state_and_certified_local_enforcement |
| Source IDs | src:usa-mt:mca-50-60-103, src:usa-mt:mca-50-60-202, src:usa-mt:mca-50-60-203, src:usa-mt:dli-building-codes-program, src:usa-mt:dli-certified-local-programs |
| Verification status | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-mt:dli-building-codes-program | DLI Building Codes Program / Building and Commercial Measurements Bureau | Adopts and enforces state building code; certifies local enforcement programs. | MCA 50-60-103; MCA 50-60-203; ARM Title 24, Chapter 301 | src:usa-mt:mca-50-60-103, src:usa-mt:mca-50-60-203, src:usa-mt:dli-current-codes | partially_verified |
| Residential | ahj:usa-mt:dli-building-codes-program | DLI Building Codes Program / Building and Commercial Measurements Bureau | Adopts and enforces residential provisions where applicable; local programs may enforce if certified. | MCA 50-60-203; ARM Title 24, Chapter 301 | src:usa-mt:mca-50-60-203, src:usa-mt:dli-current-codes, src:usa-mt:dli-certified-local-programs | partially_verified |
| Existing Building / Rehabilitation | ahj:usa-mt:dli-building-codes-program | DLI Building Codes Program / Building and Commercial Measurements Bureau | Adopts IEBC as alternate existing-building method. | MCA 50-60-203; ARM Title 24, Chapter 301 | src:usa-mt:mca-50-60-203, src:usa-mt:dli-current-codes | partially_verified |
| Mechanical | ahj:usa-mt:dli-building-codes-program | DLI Building Codes Program / Building and Commercial Measurements Bureau | Adopts and enforces mechanical code; certified local programs may enforce. | MCA 50-60-203; ARM Title 24, Chapter 301 | src:usa-mt:mca-50-60-203, src:usa-mt:dli-current-codes, src:usa-mt:dli-certified-local-programs | partially_verified |
| Plumbing | ahj:usa-mt:dli-building-codes-program | DLI Building Codes Program / Building and Commercial Measurements Bureau | Adopts UPC and related plumbing requirements; certified local programs may enforce. | MCA 50-60-203; ARM Title 24, Chapter 301 | src:usa-mt:mca-50-60-203, src:usa-mt:dli-current-codes | partially_verified |
| Fuel Gas | ahj:usa-mt:dli-building-codes-program | DLI Building Codes Program / Building and Commercial Measurements Bureau | Adopts IFGC; certified local programs may enforce. | MCA 50-60-203; ARM Title 24, Chapter 301 | src:usa-mt:mca-50-60-203, src:usa-mt:dli-current-codes | partially_verified |
| Electrical | ahj:usa-mt:dli-building-codes-program | DLI Building Codes Program / Building and Commercial Measurements Bureau | Adopts NEC and administers state electrical permits/inspections where state retains jurisdiction; professional licensing is handled through DLI professional licensing structures. | MCA 50-60-203; ARM Title 24, Chapter 301 | src:usa-mt:mca-50-60-203, src:usa-mt:dli-current-codes, src:usa-mt:dli-electrical-permits | partially_verified |
| Energy | ahj:usa-mt:dli-building-codes-program | DLI Building Codes Program / Building and Commercial Measurements Bureau | Adopts IECC energy-conservation requirements as part of state building code. | MCA 50-60-203; MCA 50-60-803; ARM Title 24, Chapter 301 | src:usa-mt:mca-50-60-203, src:usa-mt:dli-current-codes | partially_verified |
| Fire - construction references | ahj:usa-mt:dli-building-codes-program | DLI Building Codes Program / Building and Commercial Measurements Bureau | State building-code authority; building-code amendments coordinate with fire authority and replace certain IFC references with the fire code adopted by the fire AHJ. | MCA 50-60-202; MCA 50-60-203; ARM 24.301.146 | src:usa-mt:mca-50-60-202, src:usa-mt:mca-50-60-203, src:usa-mt:arm-24-301-146-permit-transition, src:usa-mt:dli-2026-draft-code-update | partially_verified |
| Fire - operational / prevention code | ahj:usa-mt:doj-fpis | Montana Department of Justice, Fire Prevention and Investigation Section / State Fire Marshal | Promotes fire safety, assists with fire-safety regulations, conducts/assists inspections and investigations, and adopts operational fire code rules. | MCA 50-3-102; MCA 50-60-202; ARM 23.12.601 | src:usa-mt:mca-50-3-102, src:usa-mt:mca-50-60-202, src:usa-mt:doj-fpi, src:usa-mt:arm-23-12-601 | partially_verified |
| Accessibility | ahj:usa-mt:dli-building-codes-program | DLI Building Codes Program / Building and Commercial Measurements Bureau | Adopts accessibility standards within state building code. | MCA 50-60-203; ARM Title 24, Chapter 301 | src:usa-mt:mca-50-60-203, src:usa-mt:dli-current-codes | partially_verified |
| Elevator / Conveyance | ahj:usa-mt:dli-building-codes-program | DLI Building Codes Program / Building and Commercial Measurements Bureau | Performs elevator-safety inspections and adopts listed elevator/conveyance safety standards. | MCA 50-60-103; ARM Title 24, Chapter 301 | src:usa-mt:mca-50-60-103, src:usa-mt:dli-building-codes-program, src:usa-mt:dli-current-codes | partially_verified |
| Boiler | ahj:usa-mt:dli-building-codes-program | DLI Building Codes Program / Building and Commercial Measurements Bureau | Performs boiler-safety inspections and adopts listed boiler standards. | MCA 50-60-103; ARM Title 24, Chapter 301 | src:usa-mt:mca-50-60-103, src:usa-mt:dli-building-codes-program, src:usa-mt:dli-current-codes | partially_verified |

### 2.3 Authority Hierarchy Notes

Montana distinguishes state building-code rulemaking from local enforcement. MCA 50-60-202 makes DLI the sole state agency for building regulations, except that the Department of Justice may promulgate regulations relating to use of buildings and installation of equipment. MCA 50-60-203 provides that DLI-adopted rules constitute the state building code. MCA 50-60-302 and DLI's certified-local-program page show that local governments cannot enforce a building code unless their enforcement program is certified and the adopted code, fee list, and enforcement plan are filed with and approved by DLI.

Operational fire prevention is a separate authority path. MCA 50-3-102 gives the Department of Justice fire-prevention rulemaking and inspection powers, while MCA 50-60-202 preserves DLI's building-regulation primacy and gives the State Fire Prevention and Investigation Section plan/regulation review responsibilities for conformity with DLI rules.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-mt:001 | ahj:usa-mt:legislature | authorizes | ahj:usa-mt:dli-building-codes-program to administer and adopt state building code by rule | src:usa-mt:mca-50-60-103, src:usa-mt:mca-50-60-203 | verified |
| edge:usa-mt:002 | ahj:usa-mt:legislature | reserves_state_rulemaking_to | ahj:usa-mt:dli-building-codes-program for building regulations, subject to DOJ exception | src:usa-mt:mca-50-60-202 | verified |
| edge:usa-mt:003 | ahj:usa-mt:dli-building-codes-program | certifies_enforcement_programs_for | certified cities, counties, and towns | src:usa-mt:dli-certified-local-programs, src:usa-mt:mca-50-60-302 | verified |
| edge:usa-mt:004 | certified cities/counties/towns | enforce_within | city/town limits or county areas not covered by certified city/town code | src:usa-mt:mca-50-60-304, src:usa-mt:mca-50-60-106 | verified |
| edge:usa-mt:005 | ahj:usa-mt:dli-building-codes-program | retains_jurisdiction_for | state-owned agency buildings and areas without certified local enforcement programs | src:usa-mt:dli-building-codes-program, src:usa-mt:mca-50-60-106 | verified |
| edge:usa-mt:006 | ahj:usa-mt:doj-fpis | adopts_and_enforces | operational fire-prevention rules / 2021 IFC adoption path | src:usa-mt:mca-50-3-102, src:usa-mt:doj-fpi, src:usa-mt:arm-23-12-601 | partially_verified |
| edge:usa-mt:007 | ahj:usa-mt:doj-fpis | reviews_for_conformity_with | DLI building-code rules and building plans/regulations | src:usa-mt:mca-50-60-202 | verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | Montana State Building Code / International Building Code as adopted and amended | International Building Code | 2021 | active | null | 2022-06-11 | null | null | Permit-issued transition rule in ARM 24.301.146; effective date verified from DLI current-code list. | src:usa-mt:dli-current-codes, src:usa-mt:arm-24-301-current-rules, src:usa-mt:arm-24-301-146-permit-transition |
| Residential | Montana Residential Code / International Residential Code as adopted and amended | International Residential Code | 2021 | active | null | 2022-06-11 | null | null | Permit-issued transition rule applies to new requirements generally; IRC-specific adoption date not separately captured. | src:usa-mt:dli-current-codes, src:usa-mt:arm-24-301-current-rules |
| Existing Building / Rehabilitation | International Existing Building Code as adopted and amended | International Existing Building Code | 2021 | active | null | 2022-06-11 | null | null | Existing building work may use IEBC/IBC pathway per ARM adoption framework; adoption date not separately captured. | src:usa-mt:dli-current-codes, src:usa-mt:arm-24-301-current-rules |
| Mechanical | International Mechanical Code as adopted and amended | International Mechanical Code | 2021 | active | null | 2022-06-11 | null | null | Permit-issued transition rule generally applies to new requirements. | src:usa-mt:dli-current-codes, src:usa-mt:arm-24-301-current-rules |
| Plumbing | Uniform Plumbing Code as adopted and amended | Uniform Plumbing Code | 2021 | active | null | 2022-06-11 | null | null | Permit-issued transition rule generally applies to new requirements. | src:usa-mt:dli-current-codes, src:usa-mt:arm-24-301-current-rules |
| Fuel Gas | International Fuel Gas Code as adopted and amended | International Fuel Gas Code | 2021 | active | null | 2022-06-11 | null | null | Permit-issued transition rule generally applies to new requirements. | src:usa-mt:dli-current-codes, src:usa-mt:arm-24-301-current-rules |
| Electrical | National Electrical Code as adopted and amended | NFPA 70, National Electrical Code | 2020 | active | null | 2022-06-11 | null | null | DLI current-code page notes technical advisory dated 2022-08-12; adoption date not separately captured. | src:usa-mt:dli-current-codes, src:usa-mt:dli-electrical-permits, src:usa-mt:arm-24-301-current-rules |
| Energy | International Energy Conservation Code as adopted and amended | International Energy Conservation Code | 2021 | active | null | 2022-06-11 | null | null | Permit-issued transition rule generally applies to new requirements. | src:usa-mt:dli-current-codes, src:usa-mt:arm-24-301-current-rules |
| Fire - construction references | Fire-related construction provisions in state building code; references coordinated with fire authority | IBC fire/life-safety provisions as amended; fire code adopted by fire AHJ where referenced | 2021 IBC cycle | active | null | 2022-06-11 | null | null | Current ARM transition rule applies to new state building-code requirements; fire-AHJ code reference language needs final official ARM capture. | src:usa-mt:dli-current-codes, src:usa-mt:mca-50-60-202, src:usa-mt:arm-24-301-146-permit-transition |
| Fire - operational / prevention code | Uniform Fire Code / International Fire Code as adopted by FPIS | International Fire Code | 2021 | active | null | null | null | null | Operational fire-code transition/effective-date details were not captured from official ARM in this pass. | src:usa-mt:arm-23-12-601, src:usa-mt:mca-50-3-102, src:usa-mt:doj-fpi |
| Accessibility | Accessibility standard adopted through ARM Title 24, Chapter 301 | ICC A117.1, Accessible and Usable Buildings and Facilities | 2017 | active | null | 2022-06-11 | null | null | No separate accessibility transition rule captured. | src:usa-mt:dli-current-codes, src:usa-mt:arm-24-301-current-rules |
| Elevator / Conveyance | Elevator and conveyance safety standards | ASME A17.1; ASME A17.3; ASME A18.1 | A17.1-2019; A17.3-2017; A18.1-2017 | active | null | 2022-06-11 | null | null | DLI notes ARM text refers to ASME A17.3 2018 but says correct edition is A17.3-2017. | src:usa-mt:dli-current-codes |
| Boiler | Boiler and controls standards | ASME Boiler and Pressure Vessel Code; ASME CSD-1 | BPVC-2021; CSD-1-2018 | active | null | 2021-11-06 | null | null | No separate boiler transition rule captured. | src:usa-mt:dli-current-codes, src:usa-mt:dli-building-codes-program |
| Swimming Pool and Spa | International Swimming Pool and Spa Code as adopted and amended | International Swimming Pool and Spa Code | 2021 | active | null | 2022-06-11 | null | null | No separate ISPSC transition rule captured. | src:usa-mt:dli-current-codes, src:usa-mt:arm-24-301-current-rules |
| Wildland-Urban Interface | International Wildland-Urban Interface Code as adopted and amended | International Wildland-Urban Interface Code | 2021 | active | null | 2022-06-11 | null | null | Applicability depends on declared/designated WUI areas; detailed locality applicability not resolved. | src:usa-mt:dli-current-codes, src:usa-mt:arm-24-301-current-rules |

### 3.2 Adoption Records

| Record ID | Code Family | State Code Name | Model Code | Edition | Effective Date | Adoption Date | Operative Date | Mandatory Date | Adoption Rule / Publication | Source IDs | Verification Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| adoption:usa-mt:ibc-2021 | Building | Montana State Building Code / IBC as adopted and amended | International Building Code | 2021 | 2022-06-11 | null | null | null | ARM Title 24, Chapter 301; DLI current-code list | src:usa-mt:dli-current-codes, src:usa-mt:arm-24-301-current-rules | partially_verified |
| adoption:usa-mt:irc-2021 | Residential | IRC as adopted and amended | International Residential Code | 2021 | 2022-06-11 | null | null | null | ARM Title 24, Chapter 301; DLI current-code list | src:usa-mt:dli-current-codes, src:usa-mt:arm-24-301-current-rules | partially_verified |
| adoption:usa-mt:iebc-2021 | Existing Building / Rehabilitation | IEBC as adopted and amended | International Existing Building Code | 2021 | 2022-06-11 | null | null | null | ARM Title 24, Chapter 301; DLI current-code list | src:usa-mt:dli-current-codes, src:usa-mt:arm-24-301-current-rules | partially_verified |
| adoption:usa-mt:imc-2021 | Mechanical | IMC as adopted and amended | International Mechanical Code | 2021 | 2022-06-11 | null | null | null | ARM Title 24, Chapter 301; DLI current-code list | src:usa-mt:dli-current-codes, src:usa-mt:arm-24-301-current-rules | partially_verified |
| adoption:usa-mt:ifgc-2021 | Fuel Gas | IFGC as adopted and amended | International Fuel Gas Code | 2021 | 2022-06-11 | null | null | null | ARM Title 24, Chapter 301; DLI current-code list | src:usa-mt:dli-current-codes, src:usa-mt:arm-24-301-current-rules | partially_verified |
| adoption:usa-mt:iecc-2021 | Energy | IECC as adopted and amended | International Energy Conservation Code | 2021 | 2022-06-11 | null | null | null | ARM Title 24, Chapter 301; DLI current-code list | src:usa-mt:dli-current-codes, src:usa-mt:arm-24-301-current-rules | partially_verified |
| adoption:usa-mt:upc-2021 | Plumbing | UPC as adopted and amended | Uniform Plumbing Code | 2021 | 2022-06-11 | null | null | null | ARM Title 24, Chapter 301; DLI current-code list | src:usa-mt:dli-current-codes, src:usa-mt:arm-24-301-current-rules | partially_verified |
| adoption:usa-mt:nec-2020 | Electrical | NEC as adopted and amended | NFPA 70, National Electrical Code | 2020 | 2022-06-11 | null | null | null | ARM Title 24, Chapter 301; DLI current-code list | src:usa-mt:dli-current-codes, src:usa-mt:dli-electrical-permits | partially_verified |
| adoption:usa-mt:a117-2017 | Accessibility | ICC A117.1 accessibility standard | ICC A117.1 | 2017 | 2022-06-11 | null | null | null | ARM Title 24, Chapter 301; DLI current-code list | src:usa-mt:dli-current-codes | partially_verified |
| adoption:usa-mt:ifc-2021 | Fire - operational / prevention code | IFC adopted by FPIS | International Fire Code | 2021 | null | null | null | null | ARM 23.12.601 | src:usa-mt:arm-23-12-601, src:usa-mt:mca-50-3-102 | partial_rule_verified_effective_date_unresolved |
| adoption:usa-mt:elevator-a17-1-2019 | Elevator / Conveyance | Elevator safety standard | ASME A17.1 | 2019 | 2022-06-11 | null | null | null | ARM Title 24, Chapter 301; DLI current-code list | src:usa-mt:dli-current-codes | partially_verified |
| adoption:usa-mt:elevator-a17-3-2017 | Elevator / Conveyance | Existing elevator safety standard | ASME A17.3 | 2017 | 2022-06-11 | null | null | null | ARM Title 24, Chapter 301; DLI current-code list | src:usa-mt:dli-current-codes | partially_verified |
| adoption:usa-mt:conveyance-a18-1-2017 | Elevator / Conveyance | Platform lifts and stairway chairlifts | ASME A18.1 | 2017 | 2022-06-11 | null | null | null | ARM Title 24, Chapter 301; DLI current-code list | src:usa-mt:dli-current-codes | partially_verified |
| adoption:usa-mt:boiler-bpvc-2021 | Boiler | Boiler and Pressure Vessel Code | ASME BPVC | 2021 | 2021-11-06 | null | null | null | ARM Title 24, Chapter 301; DLI current-code list | src:usa-mt:dli-current-codes | partially_verified |
| adoption:usa-mt:boiler-csd1-2018 | Boiler | Controls and Safety Devices | ASME CSD-1 | 2018 | 2021-11-06 | null | null | null | ARM Title 24, Chapter 301; DLI current-code list | src:usa-mt:dli-current-codes | partially_verified |
| adoption:usa-mt:ispsc-2021 | Swimming Pool and Spa | ISPSC as adopted and amended | International Swimming Pool and Spa Code | 2021 | 2022-06-11 | null | null | null | ARM Title 24, Chapter 301; DLI current-code list | src:usa-mt:dli-current-codes, src:usa-mt:arm-24-301-current-rules | partially_verified |
| adoption:usa-mt:iwui-2021 | Wildland-Urban Interface | IWUIC as adopted and amended | International Wildland-Urban Interface Code | 2021 | 2022-06-11 | null | null | null | ARM Title 24, Chapter 301; DLI current-code list | src:usa-mt:dli-current-codes, src:usa-mt:arm-24-301-current-rules | partially_verified |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

The DLI current-code page provides effective dates for the current model-code editions but does not, by itself, resolve adoption dates, operative dates, or mandatory dates. The principal transition rule captured in this pass is a permit-issued rule in ARM 24.301.146: projects with a legal building permit issued before a new requirement/effective date are not required to comply with the new requirement; subsequent alterations/remodels use requirements in effect at the time of permit issuance for the new work; and the building official may determine case by case whether the permit-issuance process was substantially complete enough to warrant exemption.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-mt:001 | Current statewide DLI code adoptions | effective_date_listing | 2022-06-11 for most building-code families; 2021-11-06 for listed boiler standards | DLI current-code page lists the model-code edition and effective date. | Not stated by this source alone. | src:usa-mt:dli-current-codes | partially_verified |
| date-rule:usa-mt:002 | New requirements, administrative rules, and new model-code editions | permit_issued_transition | no fixed period captured | A legal building permit has been issued before the effective date of the new requirement/rule/code edition. | Yes, for the permitted building/project; later alteration/remodel uses requirements in effect at permit issuance for the new work. | src:usa-mt:arm-24-301-146-permit-transition | partially_verified_with_source_caveat |
| date-rule:usa-mt:003 | Projects in process when permit is not yet issued | substantial_completion_discretion | no fixed period captured | Building official decides case by case whether the process for issuance of a legal permit was substantially complete enough to warrant exemption. | Possible at building official discretion. | src:usa-mt:arm-24-301-146-permit-transition | partially_verified_with_source_caveat |
| date-rule:usa-mt:004 | Operational fire code / 2021 IFC | effective_date_unresolved | null | ARM 23.12.601 adoption identified, but effective date was not captured from official ARM text. | unresolved | src:usa-mt:arm-23-12-601 | unresolved |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | 2024 IBC | null | null | null | null | null | draft_rulemaking_watch | src:usa-mt:dli-2026-draft-code-update | Draft MAR Notice No. 2026-95.2 proposes replacing 2021 with 2024 in ARM 24.301.131. |
| Residential | 2024 IRC | null | null | null | null | null | draft_rulemaking_watch | src:usa-mt:dli-2026-draft-code-update | Draft notice proposes replacing 2021 with 2024 in ARM 24.301.154. |
| Existing Building / Rehabilitation | 2024 IEBC | null | null | null | null | null | draft_rulemaking_watch | src:usa-mt:dli-2026-draft-code-update | Draft notice proposes replacing 2021 with 2024 in ARM 24.301.171. |
| Mechanical | 2024 IMC | null | null | null | null | null | draft_rulemaking_watch | src:usa-mt:dli-2026-draft-code-update | Draft notice proposes replacing 2021 with 2024 in ARM 24.301.172. |
| Fuel Gas | 2024 IFGC | null | null | null | null | null | draft_rulemaking_watch | src:usa-mt:dli-2026-draft-code-update | Draft notice proposes replacing 2021 with 2024 in ARM 24.301.173. |
| Energy | 2024 IECC | null | null | null | null | null | draft_rulemaking_watch | src:usa-mt:dli-2026-draft-code-update | Draft notice proposes 2024 IECC amendments. |
| Plumbing | 2024 UPC | null | null | null | null | null | draft_rulemaking_watch | src:usa-mt:dli-2026-draft-code-update | Draft notice includes UPC in the list of proposed updated codes. |
| Electrical | 2023 NEC | null | null | null | null | null | draft_rulemaking_watch | src:usa-mt:dli-2026-draft-code-update | Draft notice proposes replacing 2020 with 2023 in ARM 24.301.401. |
| Swimming Pool and Spa | 2024 ISPSC | null | null | null | null | null | draft_rulemaking_watch | src:usa-mt:dli-2026-draft-code-update | Draft notice proposes replacing 2021 with 2024 in ARM 24.301.175. |
| Wildland-Urban Interface | 2024 IWUIC | null | null | null | null | null | draft_rulemaking_watch | src:usa-mt:dli-2026-draft-code-update | Draft notice proposes replacing 2021 with 2024 in ARM 24.301.181. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-mt:001 | State building code generally | State-owned agency buildings | Building owned by any state agency | State-owned agency buildings are subject to DLI jurisdiction for plan review, permitting, and inspections; DLI must provide plans and opportunity for local comment to affected cities, towns, and counties. | src:usa-mt:mca-50-60-106 | verified |
| applicability-rule:usa-mt:002 | Local enforcement | Certified local jurisdictions | City/county/town has adopted and certified a code-enforcement program | Certified local governments carry out plan examination, permits, inspections, orders, and certificates of occupancy within their jurisdictional areas. | src:usa-mt:mca-50-60-106, src:usa-mt:mca-50-60-302, src:usa-mt:mca-50-60-304 | verified |
| applicability-rule:usa-mt:003 | State enforcement fallback | Noncertified local areas or revoked certification | Local jurisdiction has not chosen/certified enforcement or certification is revoked | State Building Codes Bureau / Building Codes Program jurisdiction applies in remaining areas; if local certification is revoked, state resumes original jurisdiction, subject to completion responsibilities for existing local permits. | src:usa-mt:dli-building-codes-program, src:usa-mt:dli-certified-local-programs, src:usa-mt:mca-50-60-302 | verified |
| applicability-rule:usa-mt:004 | Residential energy | Residential buildings | Energy-conservation provisions adopted under MCA 50-60-203 | Energy-conservation applicability and enforcement details require follow-up extraction of MCA 50-60-102 and ARM energy provisions; only the general state energy-code adoption was captured. | src:usa-mt:mca-50-60-203, src:usa-mt:dli-current-codes | unresolved_scope |
| applicability-rule:usa-mt:005 | Operational fire code | Public, business, industrial, state buildings and local fire jurisdictions | Fire-prevention inspection/rule authority | DOJ may inspect state buildings and public/business/industrial buildings as provided by statute, and may adopt rules necessary to safeguard life/property from fire hazards if not conflicting with DLI building regulations. Local fire-AHJ adoption/enforcement boundaries were not fully mapped. | src:usa-mt:mca-50-3-102, src:usa-mt:doj-fpi, src:usa-mt:mca-50-60-202 | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** model codes are adopted and amended primarily through Administrative Rules of Montana Title 24, Chapter 301 for the state building-code family. Operational fire-code amendments are in ARM Title 23, Chapter 12, including ARM 23.12.601 for the IFC.

**Where amendments are published:** ARM Title 24, Chapter 301; ARM Title 23, Chapter 12; DLI current-code and rule-notice pages; Montana Administrative Register notices.

**Amendment parsing status:** partial_core_sources_registered

### 5.2 State Amendment Sources

| Amendment Source ID | Scope | Publication Path | Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- |
| amendment-source:usa-mt:arm-24-301 | Building-code family amendments | Administrative Rules of Montana Title 24, Chapter 301 | partially_verified | src:usa-mt:arm-24-301-current-rules, src:usa-mt:dli-current-codes | Official rules site is dynamic; this pass registered the official source and relied on DLI's current-code page for current editions. |
| amendment-source:usa-mt:arm-24-301-146 | IBC modifications and transition language | ARM 24.301.146 | partially_verified_with_source_caveat | src:usa-mt:arm-24-301-146-permit-transition, src:usa-mt:dli-2026-draft-code-update | Transition language was captured through searchable rule extract and DLI draft notice. Manual official-rule capture recommended. |
| amendment-source:usa-mt:arm-23-12-601 | Operational fire code / IFC adoption | ARM 23.12.601 | partially_verified_with_source_caveat | src:usa-mt:arm-23-12-601 | Rule adoption identified; effective date and full amendments require follow-up official-rule extraction. |
| amendment-source:usa-mt:dli-2026-draft | Proposed 2026 updates to state building code | Draft MAR Notice No. 2026-95.2 PDF posted by DLI | draft_watch | src:usa-mt:dli-2026-draft-code-update | Draft only; do not treat as adopted law. |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Topic | Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| amendment:usa-mt:001 | Building / all current code updates | Permit-issued transition | A legal building permit issued before a new requirement/effective date shields that project from the new requirement; later alterations/remodels use the requirements in effect at permit issuance for the new work. | src:usa-mt:arm-24-301-146-permit-transition | partially_verified_with_source_caveat |
| amendment:usa-mt:002 | Building / state-vs-local administration | Local government fees | Local governments certified to enforce the state building code may establish their own permit fees and valuation method. | src:usa-mt:dli-2026-draft-code-update | draft_reconfirmed_current_rule_followup_needed |
| amendment:usa-mt:003 | Fire-construction interface | Fire-code reference path | The DLI draft notice reflects amendments replacing references to the International Fire Code with the fire code adopted by the fire authority having jurisdiction; current final ARM text should be rechecked before production use. | src:usa-mt:dli-2026-draft-code-update, src:usa-mt:mca-50-60-202 | draft_watch |
| amendment:usa-mt:004 | Elevator / conveyance | ASME A17.3 edition caveat | DLI current-code page states ARM refers to ASME A17.3 2018 but says the correct edition is A17.3-2017. | src:usa-mt:dli-current-codes | verified_agency_note |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-mt"
  model: "optional_certified_local_enforcement_with_state_fallback"
  enforcing_entities:
    - "Montana Department of Labor & Industry Building Codes Program / Building and Commercial Measurements Bureau"
    - "Certified city building-code enforcement programs"
    - "Certified county building-code enforcement programs"
    - "Certified town building-code enforcement programs"
  required_officials:
    - "local building official / inspectors meeting DLI certification or licensure requirements"
    - "state inspectors for state-jurisdiction work"
  state_reserved_activities:
    - "plan review, permitting, and inspections for all buildings owned by state agencies"
    - "enforcement in cities, counties, and towns without certified code-enforcement programs"
    - "jurisdiction resumes if local certification is revoked"
  source_ids:
    - "src:usa-mt:dli-building-codes-program"
    - "src:usa-mt:dli-certified-local-programs"
    - "src:usa-mt:mca-50-60-106"
    - "src:usa-mt:mca-50-60-302"
    - "src:usa-mt:mca-50-60-304"
  verification_status: "partially_verified"
  confidence: 0.86
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-mt"
  model: "state_code_limited_local_administration"
  applies_to_code_families:
    - "building"
    - "electrical"
    - "plumbing"
    - "mechanical"
    - "fuel_gas"
    - "residential"
    - "energy"
    - "existing_building"
  approval_required: true
  approving_authority_id: "ahj:usa-mt:dli-building-codes-program"
  filing_required: true
  registry_exists: null
  registry_source_ids: []
  legal_basis_source_ids:
    - "src:usa-mt:dli-certified-local-programs"
    - "src:usa-mt:mca-50-60-106"
    - "src:usa-mt:mca-50-60-302"
  verification_status: "partially_verified"
  confidence: 0.72
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Montana local governments may be enforcement AHJs without being independent technical-code authors. The verified model is: a city, county, or town may adopt a building code ordinance and enforce a code program only if certified by DLI; the adopted code, fee list, and enforcement plan must be filed with and approved by DLI; and local codes may include only codes adopted by the Building Codes Program. Separately, certified local governments may make, amend, and repeal rules for administration/enforcement and fee collection. This report therefore treats local enforcement authority as verified and local technical-amendment authority as state-limited and not fully resolved beyond the sources cited.

### 6.4 Known Local Amendment Registries

No statewide public registry of local technical amendments was verified in this pass. The DLI certified-local-program page provides a list of certified local government building-code jurisdictions and inspectors, but this pass did not parse it into a registry table.

### 6.5 Municipality-Specific Known Amendments

No municipality-specific local amendments were parsed into this report. A production AHJ resolver should separately ingest the DLI certified-jurisdiction list and any locally adopted ordinances for a target jurisdiction.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: partially_designed_not_implemented

Jurisdiction stack:

```text
Address
  -> State of Montana
  -> County
  -> Municipality / unincorporated county
  -> Certified city/county/town building-code program, if any
  -> State Building Codes Bureau / DLI Building Codes Program fallback, if no certified local program
  -> State-owned agency building check; if yes, DLI jurisdiction for plan review, permitting, and inspections
  -> Fire protection agency / State Fire Marshal / certified fire-inspection jurisdiction, as applicable
  -> Applicable state code adoption records
  -> Applicable local administrative/enforcement rules and local ordinances
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | not selected | none | statewide | unknown | pending |
| County | not selected | none | statewide | unknown | pending |
| Municipality | not selected | none | statewide | unknown | pending |
| Certified building-code jurisdictions | DLI certified-local-program list | src:usa-mt:dli-certified-local-programs | certified jurisdictions statewide | unknown | source_identified_not_extracted |
| Fire inspection jurisdictions | DOJ Fire Prevention and Investigation page links to certified fire-inspection jurisdictions | src:usa-mt:doj-fpi | fire-inspection jurisdictions statewide | unknown | source_identified_not_extracted |
| Special District | not selected | none | statewide | unknown | pending |

### 7.3 AHJ Contact Data

No AHJ contact data was normalized into this report. DLI and DOJ contact pages were identified, but the current pass focused on authority and code-adoption fields.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Source Type | Title / Citation | Publisher | URL | Accessed | Supports |
| --- | --- | --- | --- | --- | --- | --- |
| src:usa-mt:mca-50-60-103 | statute | MCA 50-60-103, Administration by department | Montana Legislature | https://mca.legmt.gov/bills/mca/title_0500/chapter_0600/part_0010/section_0030/0500-0600-0010-0030.html | 2026-06-26 | DLI administration duties and consultation with Building Codes Council |
| src:usa-mt:mca-50-60-202 | statute | MCA 50-60-202, Department to be sole state agency to promulgate building regulations -- exception | Montana Legislature | https://mca.legmt.gov/bills/mca/title_0500/chapter_0600/part_0020/section_0020/0500-0600-0020-0020.html | 2026-06-26 | DLI sole state building-regulation authority; DOJ exception; fire plan/regulation review |
| src:usa-mt:mca-50-60-203 | statute | MCA 50-60-203, Department to adopt state building code by rule | Montana Legislature | https://mca.legmt.gov/bills/mca/title_0500/chapter_0600/part_0020/section_0030/0500-0600-0020-0030.html | 2026-06-26 | Rulemaking authority and state-building-code formation |
| src:usa-mt:mca-50-60-106 | statute | MCA 50-60-106, Powers and duties of counties, cities, and towns -- public buildings owned or operated by state agencies | Montana Legislature | https://mca.legmt.gov/bills/mca/title_0500/chapter_0600/part_0010/section_0060/0500-0600-0010-0060.html | 2026-06-26 | Certified local duties; local administration/enforcement powers; state-owned building jurisdiction |
| src:usa-mt:mca-50-60-302 | statute | MCA 50-60-302, Certification of county, city, or town building codes | Montana Legislature | https://mca.legmt.gov/bills/mca/title_0500/chapter_0600/part_0030/section_0020/0500-0600-0030-0020.html | 2026-06-26 | Local enforcement certification, filing, approval, inspector requirements, state resumption on revocation |
| src:usa-mt:mca-50-60-304 | statute | MCA 50-60-304, Area of applicability of county, city, or town building code -- enforcement | Montana Legislature | https://mca.legmt.gov/bills/mca/title_0500/chapter_0600/part_0030/section_0040/0500-0600-0030-0040.html | 2026-06-26 | City/town limits, county coverage, contract enforcement |
| src:usa-mt:mca-50-3-102 | statute | MCA 50-3-102, Powers and duties of department regarding state fire prevention and investigation -- rules | Montana Legislature | https://mca.legmt.gov/bills/mca/title_0500/chapter_0030/part_0010/section_0020/0500-0030-0010-0020.html | 2026-06-26 | Department of Justice / fire-prevention powers and rulemaking |
| src:usa-mt:dli-building-codes-program | agency_page | Montana Building Codes Program | Montana Department of Labor & Industry | https://bsd.dli.mt.gov/building-codes-permits/ | 2026-06-26 | Program role, statewide code adoption/enforcement, state/local split, boiler/elevator inspections |
| src:usa-mt:dli-current-codes | agency_page | Current Codes | Montana Department of Labor & Industry | https://bsd.dli.mt.gov/building-codes-permits/current-codes | 2026-06-26 | Current model-code editions and effective dates |
| src:usa-mt:dli-certified-local-programs | agency_page | Certified City, County and Town Programs | Montana Department of Labor & Industry | https://bsd.dli.mt.gov/building-codes-permits/certified-government | 2026-06-26 | Local certification requirements, local adoption limits, state fallback jurisdiction, certified-jurisdiction list |
| src:usa-mt:dli-building-codes-council | agency_page | Building Codes Council | Montana Department of Labor & Industry | https://bsd.dli.mt.gov/building-codes-permits/building-codes-council | 2026-06-26 | Advisory council role and membership structure |
| src:usa-mt:dli-electrical-permits | agency_page | Electrical Permits | Montana Department of Labor & Industry | https://bsd.dli.mt.gov/building-codes-permits/permit-applications/electrical-permits/ | 2026-06-26 | Electrical permit and inspection administration; NEC compliance statement |
| src:usa-mt:doj-fpi | agency_page | Fire Prevention and Investigations | Montana Department of Justice | https://dojmt.gov/dci-home/investigations-bureau/fire-prevention-and-investigations/ | 2026-06-26 | State Fire Marshal / FPIS role, statewide fire prevention/investigation services, local-agency support, regulation assistance |
| src:usa-mt:arm-24-301-current-rules | regulation | Administrative Rules of Montana, Title 24, Chapter 301, Building Codes | Montana Secretary of State | https://rules.mt.gov/browse/collections/aec52c46-128e-4279-9068-8af5d5432d74/sections/e89dd2b8-de8d-48d1-9fa2-4cf9bbacb24b | 2026-06-26 | Official rules chapter containing building-code adoptions and amendments |
| src:usa-mt:arm-24-301-146-permit-transition | regulation_extract | ARM 24.301.146, Modifications to the International Building Code applicable to both DLI and local government code enforcement programs | Montana Secretary of State / official ARM page; searchable extract used due dynamic page limitations | https://rules.mt.gov/search?query=24.301.146 | 2026-06-26 | Permit-issued transition rule; IBC amendments; fire-reference interface |
| src:usa-mt:arm-23-12-601 | regulation | ARM 23.12.601, Adoption of the International Fire Code (2021 Edition) | Montana Secretary of State | https://rules.mt.gov/browse/collections/aec52c46-128e-4279-9068-8af5d5432d74/policies/3bc741e2-2367-4657-be3f-ccbb026e8746 | 2026-06-26 | 2021 IFC adoption by FPIS |
| src:usa-mt:dli-2026-draft-code-update | register_notice_draft_pdf | Draft Notice of Proposed Rulemaking, MAR Notice No. 2026-95.2, 2026 Updates to the State Building Code | Montana Department of Labor & Industry / Montana Administrative Register draft | https://bsd.dli.mt.gov/_docs/building-codes-permits/2024-BUILDING-CODE-ADOPTION-AMENDMENTS.pdf | 2026-06-26 | Pending/draft updates to 2024 model codes and 2023 NEC; source caveat applies |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-mt:arm-24-301-current-rules | dynamic_html | The official ARM rules site is dynamic and did not expose full line-level text through the parser in this pass. | Retain as official canonical source; manually download/print-to-PDF or use official API/export before full production verification. |
| src:usa-mt:arm-24-301-146-permit-transition | search_extract | The transition rule text was identified through a searchable extract and mirrored current rule text, while the official rule page requires manual capture for line-level proof. | Use for partial verification only; production should attach official ARM PDF/export. |
| src:usa-mt:arm-23-12-601 | dynamic_html | Official ARM page identifies the rule and 2021 IFC adoption, but the parser did not capture full rule text/effective date. | Use for core fire-code authority only; resolve effective date and amendments before verified status. |
| src:usa-mt:dli-2026-draft-code-update | draft_notice | The PDF is explicitly marked as a draft notice of proposed rulemaking. It is not a final adopted rule. | Watch item only; do not update current adopted-code fields until final adoption/effective date is published. |
| src:usa-mt:mca-50-60-103 | official_internet_disclaimer | MCA internet version states printed version prevails in case of inconsistencies or errors. | Accept for partial verification; retain citation to official Montana Legislature site. |
| src:usa-mt:mca-50-60-202 | official_internet_disclaimer | MCA internet version states printed version prevails in case of inconsistencies or errors. | Accept for partial verification; retain citation to official Montana Legislature site. |
| src:usa-mt:mca-50-60-203 | official_internet_disclaimer | MCA internet version states printed version prevails in case of inconsistencies or errors. | Accept for partial verification; retain citation to official Montana Legislature site. |

### 8.3 Supplemental Sources

No non-official sources were used as controlling authority in this report. Search-result snippets from non-official rule mirrors were used only to triage official ARM rule numbers and are not listed as production sources.

### 8.4 Source Extraction Metadata

| Source ID | Extraction Method | Extracted Fields | Extraction Date | Quality |
| --- | --- | --- | --- | --- |
| src:usa-mt:mca-50-60-103 | web_html_parse | statute title, text summary, administration duties | 2026-06-26 | high |
| src:usa-mt:mca-50-60-202 | web_html_parse | statute title, DLI sole authority, DOJ exception | 2026-06-26 | high |
| src:usa-mt:mca-50-60-203 | web_html_parse | rulemaking authority, state building code definition, prohibitions | 2026-06-26 | high |
| src:usa-mt:mca-50-60-106 | web_html_parse | certified local duties, local administrative powers, state-owned buildings | 2026-06-26 | high |
| src:usa-mt:mca-50-60-302 | web_html_parse | local certification requirements, revocation/fallback | 2026-06-26 | high |
| src:usa-mt:mca-50-60-304 | web_html_parse | local area applicability | 2026-06-26 | high |
| src:usa-mt:mca-50-3-102 | web_html_parse | fire-prevention duties and rulemaking | 2026-06-26 | high |
| src:usa-mt:dli-building-codes-program | web_html_parse | program role, state/local enforcement model | 2026-06-26 | high |
| src:usa-mt:dli-current-codes | web_html_parse | current code families, editions, effective dates | 2026-06-26 | high |
| src:usa-mt:dli-certified-local-programs | web_html_parse | local certification, adoption limits, jurisdiction list pointer | 2026-06-26 | high |
| src:usa-mt:dli-building-codes-council | web_html_parse | advisory council purpose and structure | 2026-06-26 | high |
| src:usa-mt:dli-electrical-permits | web_html_parse | electrical permit/inspection notes | 2026-06-26 | medium |
| src:usa-mt:doj-fpi | web_html_parse | State Fire Marshal / FPIS role and services | 2026-06-26 | high |
| src:usa-mt:arm-24-301-current-rules | search_and_dynamic_html | official rules location and chapter identity | 2026-06-26 | medium_with_caveat |
| src:usa-mt:arm-24-301-146-permit-transition | search_extract_and_draft_notice | transition rule and amendments | 2026-06-26 | medium_with_caveat |
| src:usa-mt:arm-23-12-601 | search_and_dynamic_html | 2021 IFC adoption identification | 2026-06-26 | medium_with_caveat |
| src:usa-mt:dli-2026-draft-code-update | pdf_text_parse_and_spot_screenshot | draft proposed update scope, future model-code editions, NEC update | 2026-06-26 | medium_draft_only |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 1.00 | none | Status reflects source-backed core fields plus unresolved items. |
| report | risk.overall_confidence | 0.68 | verified | 1.00 | none | Conservative confidence due ARM dynamic-source caveats and unresolved effective/adoption date distinctions. |
| ahj:usa-mt:dli-building-codes-program | authority_name | Montana Department of Labor & Industry, Business Standards Division, Building Codes Program / Building and Commercial Measurements Bureau | partially_verified | 0.90 | src:usa-mt:dli-building-codes-program, src:usa-mt:dli-current-codes | Agency program naming appears across DLI pages; exact bureau style should be normalized in schema. |
| ahj:usa-mt:dli-building-codes-program | legal_basis | MCA 50-60-103; MCA 50-60-202; MCA 50-60-203 | verified | 0.92 | src:usa-mt:mca-50-60-103, src:usa-mt:mca-50-60-202, src:usa-mt:mca-50-60-203 | Core statewide authority verified. |
| adoption:usa-mt:ibc-2021 | edition/effective_date | 2021 IBC; 2022-06-11 | verified | 0.85 | src:usa-mt:dli-current-codes | Adoption date remains unresolved. |
| adoption:usa-mt:irc-2021 | edition/effective_date | 2021 IRC; 2022-06-11 | verified | 0.85 | src:usa-mt:dli-current-codes | Adoption date remains unresolved. |
| adoption:usa-mt:nec-2020 | edition/effective_date | 2020 NEC; 2022-06-11 | verified | 0.82 | src:usa-mt:dli-current-codes | Official DLI current-code list controls over non-official summaries. |
| adoption:usa-mt:ifc-2021 | edition | 2021 IFC | partially_verified | 0.70 | src:usa-mt:arm-23-12-601 | Effective date and full amendments unresolved. |
| local-enforcement:usa-mt | model | optional_certified_local_enforcement_with_state_fallback | verified | 0.86 | src:usa-mt:dli-building-codes-program, src:usa-mt:dli-certified-local-programs, src:usa-mt:mca-50-60-302 | Local certified program model verified. |
| local-amendment-rule:usa-mt | model | state_code_limited_local_administration | partially_verified | 0.72 | src:usa-mt:dli-certified-local-programs, src:usa-mt:mca-50-60-106 | Technical local amendment scope needs deeper statute/rule review and local ordinance examples. |
| date-rule:usa-mt:002 | transition_trigger | legal building permit issued before effective date of new requirement | partially_verified_with_source_caveat | 0.70 | src:usa-mt:arm-24-301-146-permit-transition | Official ARM page should be manually captured for production. |
| pending:usa-mt:2026-code-update | future_code_status | draft rulemaking watch only | verified_as_draft | 0.85 | src:usa-mt:dli-2026-draft-code-update | Do not treat as adopted. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Every source ID used in the body is listed in Section 8. |
| All authority IDs resolve | pass | Primary and specialized authority IDs are defined in Section 2. |
| All current code families have adoption records | pass | Current core families have adoption records; unresolved dates are explicit. |
| Building and operational fire code are separated | pass | DLI construction/building-code authority is separated from DOJ/FPIS operational fire authority. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Dates are separate; unknown dates remain null. |
| Effective dates are valid ISO dates | pass | Verified date values use YYYY-MM-DD. |
| No impossible date sequences | pass | No conflicting adoption/effective/operative/mandatory sequence was introduced. |
| Transition rules have explicit trigger conditions | pass | The permit-issued transition trigger is explicit; fire-code transition remains unresolved. |
| Permit-date logic is captured where applicable | pass | Captured as legal-permit-issued trigger with caveat. |
| Local enforcement model classified | pass | Classified as optional certified local enforcement with state fallback. |
| Local amendment rule classified | pass | Classified conservatively as state-code-limited local administration. |
| AHJ confirmation metadata present | fail | AHJ contacts and boundary datasets were not normalized. |
| Official-source caveats captured | pass | Dynamic ARM and draft notice caveats are captured in Section 8.2. |
| Pending/future change status conservative | pass | Draft 2026 rulemaking is only a watch item. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-mt:001 | high | ARM rule extraction | Official ARM dynamic pages for Title 24, Chapter 301 and ARM 23.12.601 need line-level or PDF/export capture for production-grade citations. | Download official ARM chapter/rules as PDF or official export; attach to source registry; revalidate section 3 and 5. | null | null | open |
| issue:usa-mt:002 | high | fire code effective date | ARM 23.12.601 adoption of the 2021 IFC was identified, but effective date, amendments, and transition provisions were not fully parsed. | Extract full ARM 23.12.601 and related ARM Title 23, Chapter 12 rules. | null | null | open |
| issue:usa-mt:003 | medium | adoption dates vs effective dates | DLI current-code page verifies effective dates but not adoption dates. | Locate final MAR adoption notices for the 2021 code cycle and record adoption dates separately. | null | null | open |
| issue:usa-mt:004 | medium | local amendment scope | Local technical-amendment authority and registry status remain only partially verified. | Review MCA 50-60-301, ARM 24.301.202/203, and certified local government filing rules; determine whether any local amendment registry exists. | null | null | open |
| issue:usa-mt:005 | medium | certified jurisdiction normalization | DLI certified local jurisdiction list was identified but not parsed into structured AHJ data. | Ingest the DLI list, normalize jurisdiction, code-family scope, official names, and contacts. | null | null | open |
| issue:usa-mt:006 | medium | pending 2026 update | DLI draft MAR notice proposes 2024 code updates and 2023 NEC, but no final adoption/effective dates were verified. | Monitor DLI rule notices and Montana Administrative Register for final notice/adoption. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-mt:dli-current-codes | src:usa-mt:dli-current-codes | html_diff | monthly | current-code edition or effective-date list changes | 2026-06-26 | active |
| watch:usa-mt:arm-24-301 | src:usa-mt:arm-24-301-current-rules | rules_diff | monthly | Title 24, Chapter 301 rule amendments or new effective dates | 2026-06-26 | active |
| watch:usa-mt:arm-23-12-601 | src:usa-mt:arm-23-12-601 | rules_diff | monthly | IFC edition/effective-date/amendment changes | 2026-06-26 | active |
| watch:usa-mt:dli-2026-draft | src:usa-mt:dli-2026-draft-code-update | rulemaking_status | weekly_until_resolved | final proposed/adopted MAR notice published; dates added | 2026-06-26 | active |
| watch:usa-mt:certified-locals | src:usa-mt:dli-certified-local-programs | html_or_pdf_diff | quarterly | certified jurisdiction list changes | 2026-06-26 | active |
| watch:usa-mt:doj-fpi | src:usa-mt:doj-fpi | html_diff | quarterly | fire inspection jurisdiction list or State Fire Marshal guidance changes | 2026-06-26 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-26 | Populated Montana draft from official state statutes, DLI agency pages, DOJ fire authority page, ARM source registry, and DLI draft rulemaking notice. | report:usa-mt; ahj:usa-mt:dli-building-codes-program; adoption:usa-mt:*; local-enforcement:usa-mt | src:usa-mt:mca-50-60-103, src:usa-mt:mca-50-60-202, src:usa-mt:mca-50-60-203, src:usa-mt:dli-current-codes, src:usa-mt:dli-certified-local-programs, src:usa-mt:mca-50-3-102, src:usa-mt:doj-fpi | GPT-5.5 Thinking | Upgraded status to partially_verified after validation pass; unresolved items remain explicit. |
