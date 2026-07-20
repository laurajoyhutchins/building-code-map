---
state:
  state_id: "US-UT"
  name: "Utah"
  abbreviation: "UT"
report:
  report_id: "state-report:usa-ut"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-26"
  last_verified: "2026-06-26"
  reviewed_by: null
risk:
  overall_confidence: 0.62 # 0.00 - 1.00
  risk_flags:
    - "pending_code_transition_2026_07_01"
    - "state_amendments_not_fully_parsed"
    - "local_amendment_scope_not_exhaustively_parsed"
    - "project_level_transition_trigger_unresolved"
    - "ahj_boundary_contact_data_not_populated"
  open_questions_count: 5

---

# State Building Code Authority Report: Utah

## 1. Executive Summary

- **Authority model:** Utah uses a statewide legislative adoption model for construction codes. The Utah Legislature adopts the State Construction Code by statute, primarily through Title 15A, Chapters 2 through 4 and Chapter 6. The Uniform Building Code Commission (UBCC) advises the Division of Professional Licensing (DOPL), supports code-analysis work, and reports through the legislative process, but the Legislature is the primary statewide adopting authority for the State Construction Code. Source IDs: `src:usa-ut:code-15a-1-204`, `src:usa-ut:code-15a-2-103`, `src:usa-ut:code-15a-1-203`, `src:usa-ut:dopl-uniform-building-codes`.

- **Statewide code status:** As of 2026-06-26, Utah's current State Construction Code provisions in Section 15A-2-103 are effective 2025-07-01 and superseded 2026-07-01. The current construction-code package includes the 2021 IBC, 2021 IRC, 2021 IPC, 2021 IMC, 2021 IFGC, 2023 NEC, 2021 IECC, 2021 IEBC, residential provisions of the 2021 ISPSC, and additional manufactured housing, modular, historic-property, and wildland-urban-interface provisions. Source ID: `src:usa-ut:code-15a-2-103`.

- **Local enforcement model:** The State Construction Code is a statewide standard that the state and each political subdivision must follow. Enforcement is carried out through compliance agencies, local regulators, state regulators, and, where authorized, third-party inspection firms. A local regulator is a political subdivision that employs or contracts with a qualified building official and is empowered to regulate construction, alteration, remodeling, building, repair, installation, inspection, or related code-covered activities. Source IDs: `src:usa-ut:code-15a-1-202`, `src:usa-ut:code-15a-1-204`.

- **Local amendment posture:** Utah is generally state-preemptive for construction code content. Statewide amendments are codified in Title 15A, Chapter 3. Local amendments incorporated into the State Construction Code are codified in Title 15A, Chapter 4. Section 15A-1-204 restricts state executive-branch entities and political subdivisions from adopting or enforcing post-2016 ordinances, rules, or requirements that are specifically addressed by and more restrictive than the State Construction Code, except where an exception or state law supports it. Source IDs: `src:usa-ut:code-15a-1-204`, `src:usa-ut:code-chapter-3`, `src:usa-ut:code-chapter-4`.

- **Known transition periods or pending changes:** On 2026-07-01, Utah's codified future construction-code package moves several code families to 2024 editions, while retaining the 2021 IRC and 2023 NEC. On the same date, the State Fire Code moves from the 2021 IFC package to the 2024 IFC package. A further codified change on 2027-01-01 replaces the prior Utah Wildland Urban Interface Code reference with the 2024 International Wildland-Urban Interface Code. Project-level trigger rules such as permit application date, plan-submittal date, permit issuance date, or grace-period/concurrency rules were not resolved in this pass. Source IDs: `src:usa-ut:code-15a-2-103-future`, `src:usa-ut:code-15a-5-103`.

- **Production readiness:** partial_ready_for_narrow_validation

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: The Utah Legislature adopts the State Construction Code by legislation;
    UBCC advises DOPL and supports legislative code-analysis workflows.
  confidence: 0.86
  source_ids:
  - src:usa-ut:code-15a-1-204
  - src:usa-ut:code-15a-1-203
  - src:usa-ut:dopl-uniform-building-codes
- topic: Current construction code editions
  finding: Current 15A-2-103 package is effective 2025-07-01 and superseded 2026-07-01;
    it includes 2021 IBC, 2021 IRC, 2021 IPC, 2021 IMC, 2021 IFGC, 2023 NEC, 2021
    IECC, 2021 IEBC, and residential provisions of the 2021 ISPSC.
  confidence: 0.9
  source_ids:
  - src:usa-ut:code-15a-2-103
- topic: Pending construction code update
  finding: Effective 2026-07-01, 15A-2-103 moves building, plumbing, mechanical, fuel
    gas, energy, existing building, and residential swimming-pool/spa references to
    2024 editions while retaining 2021 IRC and 2023 NEC.
  confidence: 0.82
  source_ids:
  - src:usa-ut:code-15a-2-103-future
  - src:usa-ut:hb65-2026
- topic: Fire code authority and editions
  finding: The State Fire Code is separately codified under Chapter 5; current text
    incorporates the 2021 IFC package and future 2026-07-01 text incorporates the
    2024 IFC package.
  confidence: 0.88
  source_ids:
  - src:usa-ut:code-15a-5-103
  - src:usa-ut:hb45-2026
- topic: Local enforcement
  finding: Local regulators and compliance agencies are defined in Title 15A; local
    regulators must employ or contract with a qualified building official.
  confidence: 0.78
  source_ids:
  - src:usa-ut:code-15a-1-202
- topic: Local amendments
  finding: Utah local amendments are not a free-form local-adoption model; codified
    local amendments appear in Title 15A, Chapter 4, and more-restrictive local requirements
    are constrained by Section 15A-1-204.
  confidence: 0.7
  source_ids:
  - src:usa-ut:code-15a-1-204
  - src:usa-ut:code-chapter-4
- topic: Effective / operative date rule
  finding: The general legislative adoption rule points to July 1 after enactment
    unless legislation states otherwise; statewide project-level permit-date or issuance-date
    logic remains unresolved.
  confidence: 0.58
  source_ids:
  - src:usa-ut:code-15a-1-204
  - src:usa-ut:hb65-2026
  - src:usa-ut:hb45-2026
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | `ahj:usa-ut:legislature` |
| Authority name | Utah Legislature |
| Authority type | statewide legislative code-adopting authority |
| Legal basis | Utah Code § 15A-1-204; Utah Code § 15A-2-103; Utah Code § 15A-2-101 |
| Role | Adopts the State Construction Code by legislation, including nationally recognized model codes and Utah modifications. |
| Enforcement model | Statewide code followed by the state and political subdivisions; administered/enforced by compliance agencies, local regulators, and state regulators as applicable. |
| Source IDs | `src:usa-ut:code-15a-1-204`; `src:usa-ut:code-15a-2-103`; `src:usa-ut:code-15a-1-202` |
| Verification status | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | `ahj:usa-ut:legislature` | Utah Legislature | Adopts statewide State Construction Code; current base code is 2021 IBC until 2026-07-01. | Utah Code §§ 15A-1-204 and 15A-2-103 | `src:usa-ut:code-15a-1-204`; `src:usa-ut:code-15a-2-103` | partially_verified |
| Residential | `ahj:usa-ut:legislature` | Utah Legislature | Adopts statewide residential code; current and 2026-07-01 future text retain 2021 IRC. | Utah Code § 15A-2-103 | `src:usa-ut:code-15a-2-103`; `src:usa-ut:code-15a-2-103-future` | partially_verified |
| Existing Building / Rehabilitation | `ahj:usa-ut:legislature` | Utah Legislature | Adopts statewide existing-building code; current 2021 IEBC, future 2024 IEBC effective 2026-07-01. | Utah Code § 15A-2-103 | `src:usa-ut:code-15a-2-103`; `src:usa-ut:code-15a-2-103-future` | partially_verified |
| Mechanical | `ahj:usa-ut:legislature` | Utah Legislature | Adopts statewide mechanical code; current 2021 IMC, future 2024 IMC effective 2026-07-01. | Utah Code § 15A-2-103 | `src:usa-ut:code-15a-2-103`; `src:usa-ut:code-15a-2-103-future` | partially_verified |
| Plumbing | `ahj:usa-ut:legislature` | Utah Legislature | Adopts statewide plumbing code; current 2021 IPC, future 2024 IPC effective 2026-07-01. | Utah Code § 15A-2-103 | `src:usa-ut:code-15a-2-103`; `src:usa-ut:code-15a-2-103-future` | partially_verified |
| Fuel Gas | `ahj:usa-ut:legislature` | Utah Legislature | Adopts statewide fuel-gas code; current 2021 IFGC, future 2024 IFGC effective 2026-07-01. | Utah Code § 15A-2-103 | `src:usa-ut:code-15a-2-103`; `src:usa-ut:code-15a-2-103-future` | partially_verified |
| Electrical | `ahj:usa-ut:legislature` | Utah Legislature | Adopts statewide electrical code; current and 2026-07-01 future text use 2023 NEC. | Utah Code § 15A-2-103 | `src:usa-ut:code-15a-2-103`; `src:usa-ut:code-15a-2-103-future` | partially_verified |
| Energy | `ahj:usa-ut:legislature` | Utah Legislature | Adopts statewide energy code; current 2021 IECC, future 2024 IECC effective 2026-07-01. | Utah Code § 15A-2-103 | `src:usa-ut:code-15a-2-103`; `src:usa-ut:code-15a-2-103-future` | partially_verified |
| Fire - construction references | `ahj:usa-ut:legislature`; `ahj:usa-ut:state-fire-marshal` | Utah Legislature; Utah State Fire Marshal | Legislature adopts State Fire Code; State Fire Marshal / deputies / local fire enforcement authority function as AHJ under Chapter 5 definitions. | Utah Code §§ 15A-5-102 and 15A-5-103 | `src:usa-ut:code-15a-5-103`; `src:usa-ut:fire-prevention-board` | partially_verified |
| Fire - operational / prevention code | `ahj:usa-ut:legislature`; `ahj:usa-ut:fire-prevention-board`; `ahj:usa-ut:state-fire-marshal` | Utah Legislature; Utah Fire Prevention Board; Utah State Fire Marshal | State Fire Code is separately codified; Fire Prevention Board and State Fire Marshal sources support fire-code administration and code-update process. | Utah Code Chapter 5; Utah Code Title 53 fire-prevention framework not fully parsed here | `src:usa-ut:code-15a-5-103`; `src:usa-ut:fire-prevention-board` | partially_verified |
| Accessibility | `ahj:usa-ut:legislature` | Utah Legislature | Accessibility is handled through the adopted IBC package and Title 15A amendments; separate accessibility-agency authority was not populated. | Utah Code § 15A-2-103; Title 15A Chapter 3 | `src:usa-ut:code-15a-2-103`; `src:usa-ut:code-chapter-3` | partially_verified |
| Swimming Pool and Spa | `ahj:usa-ut:legislature` | Utah Legislature | Current residential provisions of the 2021 ISPSC; future residential provisions of the 2024 ISPSC effective 2026-07-01. | Utah Code § 15A-2-103 | `src:usa-ut:code-15a-2-103`; `src:usa-ut:code-15a-2-103-future` | partially_verified |
| Elevator / Conveyance | `ahj:usa-ut:unresolved-conveyance` | Unresolved | Elevator/conveyance code authority was not populated from the Title 15A sources reviewed. | unresolved | none | unresolved |

### 2.3 Authority Hierarchy Notes

- `ahj:usa-ut:legislature` is the primary statewide code-adopting authority for the State Construction Code and State Fire Code.
- `ahj:usa-ut:ubcc` advises DOPL on code administration, supports advisory committees, reports to the Business and Labor Interim Committee, and participates in the Unified Code Analysis Council process.
- `ahj:usa-ut:dopl` supports UBCC functions, training-fund activity, and statewide code-administration support; DOPL's Building Code Analyst role is described as collaborative, evaluative, and not disciplinary or enforcement-focused.
- `ahj:usa-ut:compliance-agencies` and `ahj:usa-ut:local-regulators` are enforcement/administration entities, not independent statewide code-adopting authorities.
- Fire-code authority is related but distinct: Chapter 5 adopts the State Fire Code, and fire-code AHJ language includes the State Fire Marshal, authorized deputies, and local fire enforcement authorities.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-ut:001 | `ahj:usa-ut:legislature` | adopts | State Construction Code, including model-code editions and amendments | `src:usa-ut:code-15a-1-204`; `src:usa-ut:code-15a-2-103` | partially_verified |
| edge:usa-ut:002 | `ahj:usa-ut:legislature` | adopts | State Fire Code | `src:usa-ut:code-15a-5-103` | partially_verified |
| edge:usa-ut:003 | `ahj:usa-ut:ubcc` | advises | `ahj:usa-ut:dopl` on code administration and recommendations | `src:usa-ut:code-15a-1-203`; `src:usa-ut:dopl-uniform-building-codes` | partially_verified |
| edge:usa-ut:004 | `ahj:usa-ut:dopl` | supports | statewide consistency, training fund, Building Code Analyst support | `src:usa-ut:dopl-uniform-building-codes` | partially_verified |
| edge:usa-ut:005 | `ahj:usa-ut:compliance-agencies` | enforces/administers | State Construction Code within authorized scope | `src:usa-ut:code-15a-1-202`; `src:usa-ut:code-15a-1-204` | partially_verified |
| edge:usa-ut:006 | `ahj:usa-ut:local-regulators` | enforces/administers | local construction regulation through qualified building official/designee | `src:usa-ut:code-15a-1-202` | partially_verified |
| edge:usa-ut:007 | `ahj:usa-ut:state-fire-marshal` | acts_as_AHJ_or_administers | State Fire Code with authorized deputies and local fire enforcement authorities | `src:usa-ut:code-15a-5-103`; `src:usa-ut:fire-prevention-board` | partially_verified |
| edge:usa-ut:008 | `ahj:usa-ut:legislature` | constrains_more_restrictive_local_rules | state executive branch entities and political subdivisions, subject to exceptions/state law | `src:usa-ut:code-15a-1-204`; `src:usa-ut:code-chapter-4` | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

Current-state rows are stated as of 2026-06-26. Utah has codified future changes effective 2026-07-01 and 2027-01-01; those are recorded in Sections 3.2 and 4.3.

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | State Construction Code | International Building Code, including Appendices C and J | 2021 | adopted-current; superseded 2026-07-01 | null | 2025-07-01 | 2025-07-01 | null | Current 15A-2-103 text is superseded 2026-07-01; no project-level grace rule resolved. | `src:usa-ut:code-15a-2-103` |
| Residential | State Construction Code | International Residential Code | 2021 | adopted-current; retained in 2026-07-01 future text | null | 2025-07-01 | 2025-07-01 | null | Current 15A-2-103 text is superseded 2026-07-01; future text still lists 2021 IRC. | `src:usa-ut:code-15a-2-103`; `src:usa-ut:code-15a-2-103-future` |
| Existing Building / Rehabilitation | State Construction Code | International Existing Building Code | 2021 | adopted-current; superseded by 2024 IEBC on 2026-07-01 | null | 2025-07-01 | 2025-07-01 | null | Current 15A-2-103 text is superseded 2026-07-01; no project-level grace rule resolved. | `src:usa-ut:code-15a-2-103`; `src:usa-ut:code-15a-2-103-future` |
| Mechanical | State Construction Code | International Mechanical Code | 2021 | adopted-current; superseded by 2024 IMC on 2026-07-01 | null | 2025-07-01 | 2025-07-01 | null | Current 15A-2-103 text is superseded 2026-07-01; no project-level grace rule resolved. | `src:usa-ut:code-15a-2-103`; `src:usa-ut:code-15a-2-103-future` |
| Plumbing | State Construction Code | International Plumbing Code | 2021 | adopted-current; superseded by 2024 IPC on 2026-07-01 | null | 2025-07-01 | 2025-07-01 | null | Current 15A-2-103 text is superseded 2026-07-01; no project-level grace rule resolved. | `src:usa-ut:code-15a-2-103`; `src:usa-ut:code-15a-2-103-future` |
| Fuel Gas | State Construction Code | International Fuel Gas Code | 2021 | adopted-current; superseded by 2024 IFGC on 2026-07-01 | null | 2025-07-01 | 2025-07-01 | null | Current 15A-2-103 text is superseded 2026-07-01; no project-level grace rule resolved. | `src:usa-ut:code-15a-2-103`; `src:usa-ut:code-15a-2-103-future` |
| Electrical | State Construction Code | National Electrical Code | 2023 | adopted-current; retained in 2026-07-01 future text | null | 2025-07-01 | 2025-07-01 | null | Current 15A-2-103 text is superseded 2026-07-01; future text still lists 2023 NEC. | `src:usa-ut:code-15a-2-103`; `src:usa-ut:code-15a-2-103-future` |
| Energy | State Construction Code | International Energy Conservation Code | 2021 | adopted-current; superseded by 2024 IECC on 2026-07-01 | null | 2025-07-01 | 2025-07-01 | null | Current 15A-2-103 text is superseded 2026-07-01; no project-level grace rule resolved. | `src:usa-ut:code-15a-2-103`; `src:usa-ut:code-15a-2-103-future` |
| Fire - construction references | State Fire Code | International Fire Code, excluding appendices, with Part 2 amendments | 2021 | adopted-current; superseded by 2024 IFC on 2026-07-01 | null | null | null | null | Current Chapter 5 text is superseded 2026-07-01; no project-level grace rule resolved. | `src:usa-ut:code-15a-5-103` |
| Fire - operational / prevention code | State Fire Code | 2021 IFC; NFPA 1 Ch. 38 (2018); NFPA 54 (2024); NFPA 58 (2024) | mixed | adopted-current; superseded by 2024 IFC/NFPA package on 2026-07-01 | null | null | null | null | Current Chapter 5 text is superseded 2026-07-01. | `src:usa-ut:code-15a-5-103` |
| Accessibility | State Construction Code | IBC accessibility provisions and Utah amendments | 2021 IBC-based | partially_verified | null | 2025-07-01 | 2025-07-01 | null | Separate state accessibility authority not populated. | `src:usa-ut:code-15a-2-103`; `src:usa-ut:code-chapter-3` |
| Swimming Pool and Spa | State Construction Code | Residential provisions of the International Swimming Pool and Spa Code | 2021 | adopted-current; superseded by 2024 residential ISPSC provisions on 2026-07-01 | null | 2025-07-01 | 2025-07-01 | null | Current 15A-2-103 text is superseded 2026-07-01. | `src:usa-ut:code-15a-2-103`; `src:usa-ut:code-15a-2-103-future` |
| Wildland-Urban Interface | State Construction Code / local option rule | 2006 Utah Wildland Urban Interface Code may be adopted locally as a local amendment | 2006 | local-option-current; future IWUIC shift effective 2027-01-01 | null | 2025-07-01 | 2025-07-01 | null | 2026-07-01 text keeps local adoption of 2006 UWUI; 2027-01-01 future text lists 2024 IWUIC. | `src:usa-ut:code-15a-2-103`; `src:usa-ut:code-15a-2-103-future` |
| Manufactured Housing / Modular | State Construction Code | HUD Code; 2005 NFPA 225; MBI Standards 1200 and 1205 | mixed | adopted-current | null | 2025-07-01 | 2025-07-01 | null | Applicability is subject to Title 15A manufactured-housing and modular-unit provisions not fully parsed here. | `src:usa-ut:code-15a-2-103` |
| Elevator / Conveyance | unresolved | unresolved | unresolved | unresolved | null | null | null | null | Elevator/conveyance authority and adopted standard were not populated from the sources reviewed. | none |

### 3.2 Adoption Records

| Adoption Record ID | Code Families | State Code Name | Base Code / Edition | Adoption Date | Effective Date | Operative Date | Mandatory Date | Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `adoption:usa-ut:construction:2025-07-01-current` | Building; Residential; Existing Building; Mechanical; Plumbing; Fuel Gas; Electrical; Energy; Swimming Pool and Spa; manufactured/modular/historic-property references | State Construction Code | 2021 IBC, 2021 IRC, 2021 IPC, 2021 IMC, 2021 IFGC, 2023 NEC, 2021 IECC, 2021 IEBC, 2021 residential ISPSC, plus listed specialty standards | null | 2025-07-01 | 2025-07-01 | null | current_until_2026-07-01 | `src:usa-ut:code-15a-2-103` | Statutory text states effective 2025-07-01 and superseded 2026-07-01. |
| `adoption:usa-ut:construction:2026-07-01-future` | Building; Existing Building; Mechanical; Plumbing; Fuel Gas; Energy; Swimming Pool and Spa; specialty standards | State Construction Code | 2024 IBC, 2021 IRC, 2024 IPC, 2024 IMC, 2024 IFGC, 2023 NEC, 2024 IECC, 2024 IEBC, residential provisions of 2024 ISPSC, plus listed specialty standards | null | 2026-07-01 | 2026-07-01 | null | future_codified | `src:usa-ut:code-15a-2-103-future`; `src:usa-ut:hb65-2026` | Future code text is effective 2026-07-01 and superseded 2027-01-01 because of the WUI transition. |
| `adoption:usa-ut:wui:2027-01-01-future` | Wildland-Urban Interface | State Construction Code | 2024 International Wildland-Urban Interface Code | null | 2027-01-01 | 2027-01-01 | null | future_codified | `src:usa-ut:code-15a-2-103-future` | Future definitions shift from UWUI to IWUIC effective 2027-01-01; detailed applicability and mapping process were not fully parsed. |
| `adoption:usa-ut:fire:current-through-2026-06-30` | Fire - construction references; Fire - operational / prevention code | State Fire Code | 2021 IFC excluding appendices; NFPA 1 Ch. 38 (2018); NFPA 54 (2024); NFPA 58 (2024) | null | null | null | null | current_until_2026-07-01 | `src:usa-ut:code-15a-5-103` | Current text is shown as superseded 2026-07-01. |
| `adoption:usa-ut:fire:2026-07-01-future` | Fire - construction references; Fire - operational / prevention code | State Fire Code | 2024 IFC excluding appendices; NFPA 1 Ch. 38 (2024); NFPA 54 (2024); NFPA 58 (2024) | null | 2026-07-01 | 2026-07-01 | null | future_codified | `src:usa-ut:code-15a-5-103`; `src:usa-ut:hb45-2026` | Future text is effective 2026-07-01. |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

Utah's general code-adoption statute states that legislation adopting a new State Construction Code should provide that the code takes effect on July 1 after the legislation is enacted unless the legislation expressly states another date. Current construction-code text is effective 2025-07-01 and superseded 2026-07-01. Future construction-code text is effective 2026-07-01 and, for WUI-related references, superseded 2027-01-01 by a further codified version. Current State Fire Code text is superseded 2026-07-01, and future fire-code text is effective 2026-07-01.

This pass did not resolve whether Utah has a statewide project-level transition trigger based on permit application date, plan-submittal acceptance date, permit issuance date, construction start date, or a formal grace/concurrency period.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-ut:001 | State Construction Code | legislative_default_effective_date | July 1 after enactment unless legislation states otherwise | Legislature adopts a new State Construction Code or amendments by legislation | unresolved | `src:usa-ut:code-15a-1-204` | partially_verified |
| date-rule:usa-ut:002 | Current construction-code package | statutory_effective_and_supersession | effective 2025-07-01; superseded 2026-07-01 | Current Section 15A-2-103 version | unresolved | `src:usa-ut:code-15a-2-103` | verified |
| date-rule:usa-ut:003 | Future construction-code package | statutory_effective_and_supersession | effective 2026-07-01; superseded 2027-01-01 for WUI-related future version | Future Section 15A-2-103 version | unresolved | `src:usa-ut:code-15a-2-103-future`; `src:usa-ut:hb65-2026` | partially_verified |
| date-rule:usa-ut:004 | State Fire Code future package | statutory_effective_date | effective 2026-07-01 | Future Section 15A-5-103 version | unresolved | `src:usa-ut:code-15a-5-103`; `src:usa-ut:hb45-2026` | partially_verified |
| date-rule:usa-ut:005 | Individual projects | project_level_transition_trigger | unresolved | Permit application, plan submission, permit issuance, inspection, or construction start trigger not resolved | unresolved | none | unresolved |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | 2024 IBC, including Appendices C and J | null | null | 2026-07-01 | 2026-07-01 | null | active_until_effective | `src:usa-ut:code-15a-2-103-future`; `src:usa-ut:hb65-2026` | Future text codified in 15A-2-103. |
| Residential | 2021 IRC | null | null | 2026-07-01 | 2026-07-01 | null | active_until_effective | `src:usa-ut:code-15a-2-103-future` | Future text continues the 2021 IRC. |
| Existing Building / Rehabilitation | 2024 IEBC | null | null | 2026-07-01 | 2026-07-01 | null | active_until_effective | `src:usa-ut:code-15a-2-103-future`; `src:usa-ut:hb65-2026` | Future text codified in 15A-2-103. |
| Mechanical | 2024 IMC | null | null | 2026-07-01 | 2026-07-01 | null | active_until_effective | `src:usa-ut:code-15a-2-103-future`; `src:usa-ut:hb65-2026` | Future text codified in 15A-2-103. |
| Plumbing | 2024 IPC | null | null | 2026-07-01 | 2026-07-01 | null | active_until_effective | `src:usa-ut:code-15a-2-103-future`; `src:usa-ut:hb65-2026` | Future text codified in 15A-2-103. |
| Fuel Gas | 2024 IFGC | null | null | 2026-07-01 | 2026-07-01 | null | active_until_effective | `src:usa-ut:code-15a-2-103-future`; `src:usa-ut:hb65-2026` | Future text codified in 15A-2-103. |
| Electrical | 2023 NEC | null | null | 2026-07-01 | 2026-07-01 | null | monitor_for_next_cycle | `src:usa-ut:code-15a-2-103-future` | Future text continues the 2023 NEC. |
| Energy | 2024 IECC | null | null | 2026-07-01 | 2026-07-01 | null | active_until_effective | `src:usa-ut:code-15a-2-103-future`; `src:usa-ut:hb65-2026` | Future text codified in 15A-2-103. |
| Fire - construction references | 2024 IFC excluding appendices | null | null | 2026-07-01 | 2026-07-01 | null | active_until_effective | `src:usa-ut:code-15a-5-103`; `src:usa-ut:hb45-2026` | Future text codified in 15A-5-103. |
| Fire - operational / prevention code | 2024 IFC; NFPA 1 Ch. 38 (2024); NFPA 54 (2024); NFPA 58 (2024) | null | null | 2026-07-01 | 2026-07-01 | null | active_until_effective | `src:usa-ut:code-15a-5-103`; `src:usa-ut:hb45-2026` | Future text codified in 15A-5-103. |
| Swimming Pool and Spa | Residential provisions of 2024 ISPSC | null | null | 2026-07-01 | 2026-07-01 | null | active_until_effective | `src:usa-ut:code-15a-2-103-future`; `src:usa-ut:hb65-2026` | Future text codified in 15A-2-103. |
| Wildland-Urban Interface | 2024 IWUIC | null | null | 2027-01-01 | 2027-01-01 | null | active_until_effective | `src:usa-ut:code-15a-2-103-future` | Separate WUI transition appears in codified future versions. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-ut:001 | State Construction Code | New construction, and certain additions, alterations, repairs, installations, remodeling, renovation, or change of use | Work subject to State Construction Code | A person must comply with applicable State Construction Code provisions when new construction is involved and in listed work categories such as addition, alteration, repair, installation, remodeling, renovation, or change of use. | `src:usa-ut:code-15a-1-204` | partially_verified |
| applicability-rule:usa-ut:002 | Wildland-Urban Interface | Local WUI adoption / WUI construction | Local adoption or future IWUIC applicability | Current and 2026-07-01 text allows the 2006 Utah Wildland Urban Interface Code to be adopted locally as a local amendment; 2027-01-01 future text references 2024 IWUIC. Detailed mapping and trigger rules were not parsed. | `src:usa-ut:code-15a-2-103`; `src:usa-ut:code-15a-2-103-future` | partially_verified |
| applicability-rule:usa-ut:003 | Manufactured Housing / Modular | Manufactured housing, factory-built housing, modular units | Specialty statutory scope | Section 15A-2-103 incorporates the HUD Code, NFPA 225, and MBI 1200/1205; detailed program jurisdiction was not parsed. | `src:usa-ut:code-15a-2-103` | partially_verified |
| applicability-rule:usa-ut:004 | Elevator / Conveyance | Elevator and conveyance work | unresolved | Elevator/conveyance authority and code adoption were not populated from the Title 15A sources reviewed. | none | unresolved |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** Utah codifies state amendments in Title 15A. Chapter 3 contains "Statewide Amendments Incorporated as Part of State Construction Code." Chapter 5 Part 2 contains State Fire Code amendments and additions. Chapter 4 contains local amendments that are incorporated as part of the State Construction Code for specified jurisdictions or situations.

**Where amendments are published:** Utah Legislature Title 15A, primarily Chapters 3, 4, and 5.

**Amendment parsing status:** partial_index_level. The amendment publication structure and legal basis were verified, but individual amendments were not fully extracted or normalized.

### 5.2 State Amendment Sources

| Amendment Source ID | Applies To | Publication Path | Coverage | Parsing Status | Source IDs | Caveats |
| --- | --- | --- | --- | --- | --- | --- |
| amend-source:usa-ut:chapter-3 | State Construction Code statewide amendments | Utah Code Title 15A, Chapter 3 | Statewide amendments to IBC and other adopted construction-code families | indexed_not_fully_parsed | `src:usa-ut:code-chapter-3` | Requires section-level amendment extraction before high-impact amendments can be relied on. |
| amend-source:usa-ut:chapter-4 | Local amendments incorporated into State Construction Code | Utah Code Title 15A, Chapter 4 | Local amendments for specified jurisdictions and code families | indexed_not_fully_parsed | `src:usa-ut:code-chapter-4` | This is not a comprehensive local ordinance registry outside Title 15A; local ordinances still require jurisdictional confirmation. |
| amend-source:usa-ut:chapter-5-part-2 | State Fire Code amendments | Utah Code Title 15A, Chapter 5, Part 2 | Amendments and additions to the International Fire Code incorporated as part of State Fire Code | indexed_not_fully_parsed | `src:usa-ut:code-15a-5-103` | Fire-code amendments should be parsed separately from construction-code amendments. |
| amend-source:usa-ut:commission-emergency | UBCC urgent amendments by rule | Utah Code § 15A-1-204 | Commission may amend by rule in urgent circumstances, with filing, publication, and legislative follow-up requirements | framework_verified_only | `src:usa-ut:code-15a-1-204` | No active emergency amendment was identified or normalized in this pass. |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Topic | Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| amendment:usa-ut:structure:publication-model | All State Construction Code families | Publication model | Statewide amendments are codified in Title 15A, Chapter 3, and are part of the State Construction Code. | `src:usa-ut:code-chapter-3`; `src:usa-ut:code-15a-2-103` | partially_verified |
| amendment:usa-ut:local:publication-model | Local amendments | Publication model | Codified local amendments are in Title 15A, Chapter 4, and are incorporated as part of the State Construction Code for specified local scopes. | `src:usa-ut:code-chapter-4`; `src:usa-ut:code-15a-2-103` | partially_verified |
| amendment:usa-ut:fire:publication-model | Fire Code | Publication model | International Fire Code amendments and additions are handled in Chapter 5 Part 2, separate from Chapter 3 construction-code amendments. | `src:usa-ut:code-15a-5-103` | partially_verified |
| amendment:usa-ut:content:unparsed | All | Individual amendment content | High-impact substantive amendments, such as energy, residential sprinkler, snow-load, WUI, fire, and local amendments, require detailed extraction before production use. | `src:usa-ut:code-chapter-3`; `src:usa-ut:code-chapter-4`; `src:usa-ut:code-15a-5-103` | unresolved |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-ut"
  model: "statewide_code_with_state_and_local_compliance_agencies"
  enforcing_entities:
    - "compliance agencies issuing construction permits or empowered to enforce the codes"
    - "local regulators that employ or contract with a qualified building official"
    - "state regulators empowered to regulate code-covered construction activities"
    - "third-party inspection firms where authorized"
    - "State Fire Marshal, authorized deputies, and local fire enforcement authorities for fire-code AHJ functions"
  required_officials:
    - "qualified building official for a local regulator"
    - "certified/licensed inspection personnel requirements exist but were not exhaustively parsed"
  state_reserved_activities:
    - "statewide legislative adoption of State Construction Code"
    - "statewide legislative adoption of State Fire Code"
    - "limits on more-restrictive local ordinances, rules, or requirements specifically addressed by the State Construction Code"
  source_ids:
    - "src:usa-ut:code-15a-1-202"
    - "src:usa-ut:code-15a-1-204"
    - "src:usa-ut:code-15a-5-103"
  verification_status: "partially_verified"
  confidence: 0.74
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-ut"
  model: "state_preemptive_with_codified_local_amendments_and_statutory_exceptions"
  applies_to_code_families:
    - "State Construction Code families"
    - "Wildland-urban-interface local adoption where authorized"
    - "specified local amendments codified in Title 15A Chapter 4"
  approval_required: "conditional_or_statutory; exact local-submission and approval workflow unresolved"
  approving_authority_id: "ahj:usa-ut:legislature"
  filing_required: "unresolved outside codified Chapter 4 amendments"
  registry_exists: "partial_statutory_registry"
  registry_source_ids:
    - "src:usa-ut:code-chapter-4"
  legal_basis_source_ids:
    - "src:usa-ut:code-15a-1-204"
    - "src:usa-ut:code-chapter-4"
    - "src:usa-ut:code-15a-2-103"
  verification_status: "partially_verified"
  confidence: 0.68
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Utah local enforcement is not the same as local amendment authority. Local regulators and compliance agencies may enforce/administer the statewide code within their authorized scope, but local governments do not appear to have broad independent authority to adopt more-restrictive code requirements that are already specifically addressed by the State Construction Code. Codified local amendments are part of Title 15A, Chapter 4, while the detailed workflow for new local amendments remains unresolved.

### 6.4 Known Local Amendment Registries

| Registry ID | Registry Name | Scope | Source ID | Status | Notes |
| --- | --- | --- | --- | --- | --- |
| registry:usa-ut:title-15a-chapter-4 | Utah Code Title 15A, Chapter 4, Local Amendments Incorporated as Part of State Construction Code | Statutory local amendments incorporated into the State Construction Code | `src:usa-ut:code-chapter-4` | partially_verified | Treat as the primary statutory publication path for codified local amendments. |
| registry:usa-ut:local-ordinances | Local jurisdiction ordinances and policies | Municipal/county ordinances outside the state code | none | unresolved | Not searched or normalized in this pass; needed for address-level AHJ/local-amendment resolution. |

### 6.5 Municipality-Specific Known Amendments

| Jurisdiction | Code Family | Amendment Topic | Source IDs | Status | Notes |
| --- | --- | --- | --- | --- | --- |
| Park City Corporation / Park City Fire District | unresolved | Codified local amendment references observed in Chapter 4 index-level source | `src:usa-ut:code-chapter-4` | index_only | Individual amendment content was not extracted. |
| Salt Lake City | unresolved | Codified local amendment references observed in Chapter 4 index-level source | `src:usa-ut:code-chapter-4` | index_only | Individual amendment content was not extracted. |
| South Jordan | unresolved | Codified local amendment references observed in Chapter 4 index-level source | `src:usa-ut:code-chapter-4` | index_only | Individual amendment content was not extracted. |

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: partial_framework_only

Jurisdiction stack:

```text
Address
  -> State of Utah
  -> County
  -> Municipality / unincorporated county
  -> Compliance agency or state regulator with permit/enforcement authority
  -> Local regulator, if a political subdivision employs or contracts with a qualified building official
  -> Fire AHJ: State Fire Marshal, authorized deputy, or local fire enforcement authority
  -> Applicable State Construction Code adoption records
  -> Applicable State Fire Code adoption records
  -> Applicable Title 15A Chapter 3 statewide amendments
  -> Applicable Title 15A Chapter 4 local amendments
  -> Local ordinance/policy checks for enforcement administration, forms, fees, and local processes
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | not selected | none | statewide | unresolved | pending |
| County | not selected | none | statewide | unresolved | pending |
| Municipality | not selected | none | statewide | unresolved | pending |
| Fire District | not selected | none | statewide | unresolved | pending |
| Special District | not selected | none | statewide | unresolved | pending |
| Compliance agency / local regulator contacts | not selected | none | statewide | unresolved | pending |

### 7.3 AHJ Contact Data

No address-level AHJ contact data was populated. A production resolver still needs authoritative state, county, municipal, fire-district, and compliance-agency datasets.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Source Type | Official Status | Title / Locator | URL | Key Facts Supported | Caveats |
| --- | --- | --- | --- | --- | --- | --- |
| `src:usa-ut:code-title-15a` | statute | official | Utah Code Title 15A, State Construction and Fire Codes Act | https://le.utah.gov/xcode/Title15a/15a.html | Overall Title 15A structure; Chapters 1 through 6. | HTML navigation can omit details shown in PDFs; verify section-level current/future versions before relying on dates. |
| `src:usa-ut:code-15a-1-202` | statute | official | Utah Code § 15A-1-202, Definitions | https://le.utah.gov/xcode/Title15A/Chapter1/15A-1-S202.html | Definitions of commission, compliance agency, construction code, local regulator, qualified building official, and state regulator. | Date-sensitive; section has effective-version history. |
| `src:usa-ut:code-15a-1-203` | statute | official | Utah Code § 15A-1-203, Uniform Building Code Commission -- Unified Code Analysis Council | https://le.utah.gov/xcode/Title15A/Chapter1/15A-1-S203.html | UBCC creation, advisory role, member structure, reporting and code-analysis role. | Detailed member roster should be cross-checked against DOPL for current appointments. |
| `src:usa-ut:code-15a-1-204` | statute | official | Utah Code § 15A-1-204, Adoption of State Construction Code | https://le.utah.gov/xcode/Title15A/Chapter1/15A-1-S204.html | State Construction Code definition; statewide compliance; legislative adoption process; July 1 default; commission emergency amendment framework; local more-restrictive-rule constraint. | Some future/effective text may appear in PDF more clearly than in dynamic HTML. |
| `src:usa-ut:code-15a-2-103` | statute | official | Utah Code § 15A-2-103, current effective 2025-07-01 and superseded 2026-07-01 | https://le.utah.gov/xcode/Title15A/Chapter2/15A-2-S103.html | Current construction-code editions and specialty standards. | Must distinguish current from future/superseded versions. |
| `src:usa-ut:code-15a-2-103-future` | statute | official | Utah Code § 15A-2-103, future effective 2026-07-01 and 2027-01-01 versions | https://le.utah.gov/xcode/Title15A/Chapter2/C15A-2_1800010118000101.pdf | Future construction-code editions and WUI transition. | Compiled PDF includes multiple effective versions; do not collapse 2026-07-01 and 2027-01-01 changes. |
| `src:usa-ut:code-chapter-3` | statute | official | Utah Code Title 15A, Chapter 3, Statewide Amendments Incorporated as Part of State Construction Code | https://le.utah.gov/xcode/Title15a/Chapter3/15a-3.html | Statewide amendment publication structure and amendment index. | Individual amendments were not fully extracted. |
| `src:usa-ut:code-chapter-4` | statute | official | Utah Code Title 15A, Chapter 4, Local Amendments Incorporated as Part of State Construction Code | https://le.utah.gov/xcode/Title15A/Chapter4/15A-4-P1.html | Statutory publication path for codified local amendments. | Index-level use only; individual local amendments require section-level extraction and local confirmation. |
| `src:usa-ut:code-15a-5-103` | statute | official | Utah Code § 15A-5-103, Nationally recognized codes incorporated by reference into State Fire Code | https://le.utah.gov/xcode/Title15A/Chapter5/15A-5-S103.html | Current and future State Fire Code editions, including 2021 IFC current and 2024 IFC future effective 2026-07-01. | Fire-code amendments in Chapter 5 Part 2 were not fully parsed. |
| `src:usa-ut:dopl-uniform-building-codes` | agency page | official | Utah Department of Commerce / DOPL, Uniform Building Codes | https://commerce.utah.gov/dopl/uniform-building-codes/ | DOPL Building Code Analyst support role, training fund, building-official management training announcement, links to boards/commissions. | Agency page is supplemental; statutes control legal adoption. |
| `src:usa-ut:dopl-ubcc` | agency page | official | Utah Department of Commerce / DOPL, Uniform Building Code Commission | https://commerce.utah.gov/dopl/uniform-building-codes/uniform-building-code-commission/ | Current UBCC board page and roster context. | Roster dates can change; use only for administrative context unless refreshed. |
| `src:usa-ut:fire-prevention-board` | agency page | official | Utah State Fire Marshal Office, Fire Prevention Board | https://firemarshal.utah.gov/boards/fire-prevention-board/ | Fire Prevention Board context and fire-code update process/meeting records. | Meeting-page content is not codified law; use for process monitoring, not as sole adoption authority. |
| `src:usa-ut:hb65-2026` | legislation | official | 2026 H.B. 65, State Construction Code Modifications | https://le.utah.gov/Session/2026/bills/introduced/HB0065S04.pdf | 2026 construction-code transition and effective-date support. | Use codified 15A-2-103 future text as controlling source once published. |
| `src:usa-ut:hb45-2026` | legislation | official | 2026 H.B. 45, State Fire Code Modifications | https://le.utah.gov/Session/2026/bills/enrolled/HB0045.pdf | 2026 fire-code transition and effective-date support. | Use codified 15A-5-103 future text as controlling source once published. |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| `src:usa-ut:code-15a-2-103-future` | multiple_effective_versions | The compiled Chapter 2 PDF includes current, 2026-07-01, and 2027-01-01 text. | Must version adoption records by effective date. |
| `src:usa-ut:code-15a-5-103` | multiple_effective_versions | Chapter 5 text includes current State Fire Code text and future 2026-07-01 text. | Must version fire adoption records by effective date. |
| `src:usa-ut:dopl-uniform-building-codes` | supplemental_agency_page | DOPL page explains administration and training support but does not itself adopt codes. | Use for administrative context only. |
| `src:usa-ut:fire-prevention-board` | meeting_monitoring_source | Fire Prevention Board page/minutes help monitor pending changes but statutes control adoption. | Use as watch target and context only. |
| `src:usa-ut:code-chapter-4` | local_scope_incomplete | Chapter 4 is a statutory local-amendment publication path, not a full local ordinance/AHJ dataset. | Pair with municipal/county ordinance research for address-level output. |

### 8.3 Supplemental Sources

None used as authority for legal conclusions. All findings above rely on official Utah legislative or agency sources.

### 8.4 Source Extraction Metadata

| Source ID | Extracted Fields | Extraction Method | Extracted On | Confidence | Notes |
| --- | --- | --- | --- | --- | --- |
| `src:usa-ut:code-15a-1-202` | local regulator, compliance agency, construction code, qualified building official definitions | official statute HTML/PDF review | 2026-06-26 | 0.82 | Definitions are high value for enforcement model. |
| `src:usa-ut:code-15a-1-204` | adoption authority, statewide compliance, amendment process, local-preemption constraint, date rule | official statute HTML/PDF review | 2026-06-26 | 0.84 | Project-level transition trigger remains unresolved. |
| `src:usa-ut:code-15a-2-103` | current construction code editions and dates | official statute HTML/PDF review | 2026-06-26 | 0.90 | Current version effective 2025-07-01 and superseded 2026-07-01. |
| `src:usa-ut:code-15a-2-103-future` | future construction code editions and WUI transition | official compiled statute PDF review | 2026-06-26 | 0.82 | Multiple future versions require careful date slicing. |
| `src:usa-ut:code-15a-5-103` | current and future fire-code editions | official statute HTML/PDF review | 2026-06-26 | 0.88 | Fire amendments not individually parsed. |
| `src:usa-ut:code-chapter-3` | state amendment publication path | official statute index review | 2026-06-26 | 0.70 | Section-level amendments require extraction. |
| `src:usa-ut:code-chapter-4` | local amendment publication path | official statute index review | 2026-06-26 | 0.68 | Section-level amendments require extraction. |
| `src:usa-ut:dopl-uniform-building-codes` | DOPL / UBCC administration context | official agency-page review | 2026-06-26 | 0.70 | Administrative rather than adoption authority. |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 1.00 | none | Set after source-backed authority, adoption, fire-code, and local-enforcement fields were populated and QA checks were run. |
| report | risk.overall_confidence | 0.62 | verified | 1.00 | none | Moderate confidence due to verified core statewide authority and code-adoption fields, with unresolved local/project-level details. |
| `ahj:usa-ut:legislature` | primary authority | Utah Legislature | partially_verified | 0.86 | `src:usa-ut:code-15a-1-204`; `src:usa-ut:code-15a-2-103` | Statewide legislative adoption model verified. |
| `ahj:usa-ut:ubcc` | advisory role | Advises DOPL and participates in code-analysis processes | partially_verified | 0.80 | `src:usa-ut:code-15a-1-203`; `src:usa-ut:dopl-uniform-building-codes` | Commission role verified; full rule/procedure flow not fully parsed. |
| `adoption:usa-ut:construction:2025-07-01-current` | current code editions | 2021 IBC/IRC/IPC/IMC/IFGC, 2023 NEC, 2021 IECC/IEBC, 2021 residential ISPSC | verified | 0.90 | `src:usa-ut:code-15a-2-103` | Current package is superseded 2026-07-01. |
| `adoption:usa-ut:construction:2026-07-01-future` | future code editions | 2024 IBC/IPC/IMC/IFGC/IECC/IEBC/ISPSC, 2021 IRC, 2023 NEC | partially_verified | 0.82 | `src:usa-ut:code-15a-2-103-future`; `src:usa-ut:hb65-2026` | Future text is codified, but project-level transition trigger unresolved. |
| `adoption:usa-ut:fire:current-through-2026-06-30` | current fire code | 2021 IFC package | verified | 0.88 | `src:usa-ut:code-15a-5-103` | Current text superseded 2026-07-01. |
| `adoption:usa-ut:fire:2026-07-01-future` | future fire code | 2024 IFC package | partially_verified | 0.86 | `src:usa-ut:code-15a-5-103`; `src:usa-ut:hb45-2026` | Future text effective 2026-07-01. |
| `local-enforcement:usa-ut` | enforcement model | statewide code with local/state compliance agencies and local regulators | partially_verified | 0.74 | `src:usa-ut:code-15a-1-202`; `src:usa-ut:code-15a-1-204` | Contact datasets and AHJ resolver not populated. |
| `local-amendment-rule:usa-ut` | local amendment posture | state-preemptive with codified local amendments and exceptions | partially_verified | 0.68 | `src:usa-ut:code-15a-1-204`; `src:usa-ut:code-chapter-4` | New-local-amendment workflow unresolved. |
| `date-rule:usa-ut:005` | project-level transition trigger | unresolved | unresolved | 0.20 | none | Permit/application/issuance/concurrency rule needs further research. |
| `ahj:usa-ut:unresolved-conveyance` | elevator/conveyance authority | unresolved | unresolved | 0.10 | none | Not covered by this Title 15A pass. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| Leftover template markers removed | pass | Checked for leftover placeholder/template tokens; no matches were found in this file. |
| All source IDs resolve | pass | Every `src:usa-ut:*` cited in the body appears in Section 8. |
| All authority IDs resolve | pass | Authority IDs used in tables are defined or explicitly labeled unresolved. |
| All current code families have adoption rows | pass | Rows exist for the template code families plus Utah-specific swimming-pool/spa, WUI, and manufactured/modular categories; elevator/conveyance remains unresolved. |
| Building and operational fire code are separated | pass | Construction-code and fire-code adoption records are separate. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Date fields are separated; unresolved dates remain null. |
| Effective dates are valid ISO dates | pass | Populated dates use YYYY-MM-DD format. |
| No impossible date sequences | pass | Future effective/supersession sequence is date-sliced and no backward sequence was introduced. |
| Transition rules have explicit trigger conditions | partial | Statutory effective/supersession triggers are captured; project-level triggers remain unresolved. |
| Permit-date logic is captured where applicable | fail | Statewide permit-application/permit-issuance/concurrency rule was not resolved. |
| Local enforcement model classified | pass | Local regulator and compliance-agency framework is classified. |
| Local amendment rule classified | partial | State-preemptive posture and Chapter 4 publication path are captured; workflow still unresolved. |
| AHJ confirmation metadata present | fail | No address-level AHJ contact or boundary datasets are populated. |
| Official-source caveats captured | pass | Caveats for multiple effective versions, agency pages, and Chapter 4 local scope are captured. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-ut:001 | high | project-level transition logic | Determine whether statewide transition uses permit application, plan-submittal acceptance, permit issuance, construction start, grace period, or local policy. | Search Utah statutes/rules and DOPL/UBCC guidance; verify whether local policies control. | null | null | open |
| issue:usa-ut:002 | high | state amendments | Parse Title 15A Chapter 3 and Chapter 5 Part 2 into normalized amendment records, including high-impact residential, energy, structural, fire, and WUI changes. | Section-level extraction and amendment impact tagging. | null | null | open |
| issue:usa-ut:003 | high | local amendment workflow | Confirm the exact statutory workflow for new local amendments, including submission, legislative approval, filing, publication, and exceptions to the more-restrictive-local-rule constraint. | Parse Chapter 4, Section 15A-1-204 exceptions, and any DOPL/UBCC procedures. | null | null | open |
| issue:usa-ut:004 | medium | elevator/conveyance authority | Identify Utah elevator/conveyance code authority, adopted standards, inspection entity, and enforcement model. | Search Utah statutes/rules outside Title 15A if needed. | null | null | open |
| issue:usa-ut:005 | medium | AHJ boundary/contact resolver | Populate state, county, municipal, fire-district, compliance-agency, and local-regulator boundary/contact datasets. | Select authoritative GIS/contact sources and create resolver records. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-ut:title-15a | `src:usa-ut:code-title-15a` | html_diff | monthly | Title 15A chapter or section restructuring | 2026-06-26 | active |
| watch:usa-ut:construction-adoptions | `src:usa-ut:code-15a-2-103` | html_diff | monthly_until_2027-01-01 | new effective version, supersession date, or code-edition change | 2026-06-26 | active |
| watch:usa-ut:construction-future | `src:usa-ut:code-15a-2-103-future` | pdf_diff | monthly_until_2027-01-01 | changes to 2026-07-01 or 2027-01-01 future construction-code versions | 2026-06-26 | active |
| watch:usa-ut:fire-adoptions | `src:usa-ut:code-15a-5-103` | html_diff | monthly_until_2026-07-01 | changes to State Fire Code adoption text or future effective versions | 2026-06-26 | active |
| watch:usa-ut:chapter-3-amendments | `src:usa-ut:code-chapter-3` | html_diff | quarterly | new statewide amendment section or amendment change | 2026-06-26 | active |
| watch:usa-ut:chapter-4-local-amendments | `src:usa-ut:code-chapter-4` | html_diff | quarterly | new or amended codified local amendment | 2026-06-26 | active |
| watch:usa-ut:dopl-ubc | `src:usa-ut:dopl-uniform-building-codes` | html_diff | monthly | new DOPL guidance, building-official training rule, or code analyst update | 2026-06-26 | active |
| watch:usa-ut:fire-prevention-board | `src:usa-ut:fire-prevention-board` | html_diff | monthly | Fire Prevention Board agenda/minutes regarding code adoption or amendments | 2026-06-26 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-24 | Generated baseline draft report | report:usa-ut | none | Codex | Initial stub contained unresolved placeholders and no official Utah sources. |
| 2026-06-26 | Populated Utah authority, adoption, fire-code, local-enforcement, amendment-publication, source-registry, QA, and monitoring sections | `ahj:usa-ut:legislature`; `adoption:usa-ut:construction:2025-07-01-current`; `adoption:usa-ut:construction:2026-07-01-future`; `adoption:usa-ut:fire:current-through-2026-06-30`; `adoption:usa-ut:fire:2026-07-01-future`; `local-enforcement:usa-ut`; `local-amendment-rule:usa-ut` | `src:usa-ut:code-15a-1-202`; `src:usa-ut:code-15a-1-203`; `src:usa-ut:code-15a-1-204`; `src:usa-ut:code-15a-2-103`; `src:usa-ut:code-15a-2-103-future`; `src:usa-ut:code-15a-5-103`; `src:usa-ut:code-chapter-3`; `src:usa-ut:code-chapter-4`; `src:usa-ut:dopl-uniform-building-codes`; `src:usa-ut:dopl-ubcc`; `src:usa-ut:fire-prevention-board`; `src:usa-ut:hb65-2026`; `src:usa-ut:hb45-2026` | OpenAI | Upgraded status to partially_verified after validation checks; kept unresolved issues explicit for project-level transition triggers, full amendment parsing, local amendment workflow, elevator/conveyance authority, and AHJ datasets. |
