---
state:
  state_id: "US-VA"
  name: "Virginia"
  abbreviation: "VA"
report:
  report_id: "state-report:usa-va"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-26"
  last_verified: "2026-06-26"
  reviewed_by: null
risk:
  overall_confidence: 0.72 # 0.00 - 1.00
  risk_flags:
    - "full_state_amendment_text_not_parsed"
    - "local_fire_ordinance_registry_unverified"
    - "ahj_boundary_and_contact_data_not_populated"
    - "state_owned_building_process_not_fully_traced"
    - "2024_code_development_cycle_pending"
  open_questions_count: 5

---

# State Building Code Authority Report: Virginia

## 1. Executive Summary

- **Authority model:** Virginia uses a statewide model. The Virginia Board of Housing and Community Development is the primary statewide adopting authority for the Virginia Uniform Statewide Building Code (USBC), and DHCD administers the state code program. The USBC supersedes local building codes and regulations, while local building departments enforce construction and rehabilitation requirements.

- **Statewide code status:** The current Virginia code cycle is the 2021 cycle, effective `2024-01-18`. DHCD states that Virginia adopted the 2021 I-Codes as referenced in Virginia Construction Code Part I and the 2021 Statewide Fire Prevention Code, plus the 2020 National Electrical Code.

- **Local enforcement model:** Construction and rehabilitation enforcement is local-building-department based. Counties and municipalities must have a local building department and local appeals board or enter an approved enforcement/appeals agreement. For towns under 3,500 population that do not elect to enforce the Building Code, the county enforces it for the town.

- **Local amendment posture:** Building-code amendments are primarily statewide: the statutory USBC supersedes local building codes and regulations, subject to statutory exceptions. Fire prevention is different: local governments may adopt fire prevention regulations that are more restrictive or more extensive than the Statewide Fire Prevention Code, but they may not affect the manner of construction or materials used in erection, alteration, repair, or use of a building or structure.

- **Known transition periods or pending changes:** For construction permit applications submitted during the one-year period beginning `2024-01-18`, applicants may choose either the 2021 code or the immediately prior edition. This yields a calculated mandatory date of `2025-01-18` for permit applications governed by that transition rule. Virginia also has an active 2024 Code Development Cycle page with 2025-2026 workgroup and board materials; no final 2024-code effective date was captured in this report.

- **Production readiness:** partially_ready_for_authority_and_current_adoption_fields. Not ready for full production use until state amendments are parsed, AHJ contact/boundary data is added, and 2024-cycle monitoring is completed.

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: The Board of Housing and Community Development adopts and promulgates the
    USBC; DHCD administers the program.
  confidence: 0.9
  source_ids:
  - src:usa-va:code-36-98
  - src:usa-va:dhcd-codes
  - src:usa-va:dhcd-usbc
- topic: Primary building code edition
  finding: Virginia Construction Code / USBC Part I is the 2021 edition and incorporates
    2021 IBC Chapters 2-35.
  confidence: 0.9
  source_ids:
  - src:usa-va:vac-13-5-63
  - src:usa-va:register-usbc-2021-final
- topic: Electrical code edition
  finding: The current statewide code set references the 2020 National Electrical
    Code / NFPA 70.
  confidence: 0.85
  source_ids:
  - src:usa-va:dhcd-codes
  - src:usa-va:vac-13-5-63
- topic: Fire code authority
  finding: The Board of Housing and Community Development adopts the SFPC, cooperatively
    developed with the Fire Services Board.
  confidence: 0.88
  source_ids:
  - src:usa-va:code-27-97
  - src:usa-va:register-sfpc-2021-final
- topic: Operational fire code edition
  finding: The SFPC incorporates the 2021 International Fire Code, with Virginia administrative
    replacement of IFC Chapter 1 and state amendments.
  confidence: 0.88
  source_ids:
  - src:usa-va:vac-13-5-52-30
  - src:usa-va:register-sfpc-2021-final
- topic: Local building-code enforcement
  finding: Local building departments enforce construction and rehabilitation requirements;
    local appeals precede state review-board appeals.
  confidence: 0.88
  source_ids:
  - src:usa-va:code-36-105
  - src:usa-va:dhcd-usbc
- topic: Local fire-code enforcement
  finding: Local governments may enforce the SFPC; the State Fire Marshal enforces
    in jurisdictions where local governments do not enforce it.
  confidence: 0.84
  source_ids:
  - src:usa-va:code-27-98
- topic: Local amendments
  finding: Building codes are superseded statewide; local fire prevention regulations
    may be more restrictive or broader if they do not regulate construction manner
    or materials.
  confidence: 0.8
  source_ids:
  - src:usa-va:code-36-98
  - src:usa-va:code-27-97
- topic: Effective / operative date rule
  finding:
  - '2024-01-18'
  confidence: 0.86
  source_ids:
  - src:usa-va:dhcd-codes
  - src:usa-va:vac-13-5-63
  - src:usa-va:register-usbc-2021-final
  - src:usa-va:register-sfpc-2021-final
- topic: Pending changes
  finding: The 2024 Code Development Cycle is active, with base documents and 2026
    BHCD meeting materials posted.
  confidence: 0.76
  source_ids:
  - src:usa-va:dhcd-2024-cdc
  - src:usa-va:dhcd-bhcd
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-va:bhcd |
| Authority name | Virginia Board of Housing and Community Development |
| Authority type | statewide_adopting_board |
| Legal basis | Code of Virginia § 36-98 directs and empowers the Board to adopt and promulgate the Uniform Statewide Building Code. |
| Role | Adopts and amends the USBC; initiates regulatory actions to incorporate new model-code editions; final regulations are published through the Virginia Register / Virginia Administrative Code. |
| Enforcement model | statewide_code_local_enforcement |
| Source IDs | src:usa-va:code-36-98; src:usa-va:dhcd-codes; src:usa-va:dhcd-usbc; src:usa-va:register-usbc-2021-final |
| Verification status | verified_core_authority |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-va:bhcd | Virginia Board of Housing and Community Development | Adopts/amends USBC / VCC. | Code of Virginia §§ 36-98, 36-99; 13VAC5-63. | src:usa-va:code-36-98; src:usa-va:code-36-99; src:usa-va:vac-13-5-63 | verified |
| Residential | ahj:usa-va:bhcd | Virginia Board of Housing and Community Development | Adopts/amends residential provisions within the VCC / referenced IRC. | Code of Virginia §§ 36-98, 36-99; 13VAC5-63. | src:usa-va:code-36-98; src:usa-va:vac-13-5-63 | verified |
| Existing Building / Rehabilitation | ahj:usa-va:bhcd | Virginia Board of Housing and Community Development | Adopts/amends USBC Part II, Virginia Existing Building Code. | Code of Virginia §§ 36-98, 36-99; 13VAC5-63 Part II. | src:usa-va:code-36-98; src:usa-va:vac-13-5-63 | verified |
| Mechanical | ahj:usa-va:bhcd | Virginia Board of Housing and Community Development | Administers mechanical provisions through VCC referenced codes. | 13VAC5-63 incorporation of 2021 IBC and major referenced codes. | src:usa-va:vac-13-5-63 | verified_core_reference |
| Plumbing | ahj:usa-va:bhcd | Virginia Board of Housing and Community Development | Administers plumbing provisions through VCC referenced codes. | 13VAC5-63 incorporation of 2021 IBC and major referenced codes. | src:usa-va:vac-13-5-63 | verified_core_reference |
| Fuel Gas | ahj:usa-va:bhcd | Virginia Board of Housing and Community Development | Administers fuel-gas provisions through VCC referenced codes. | 13VAC5-63 incorporation of 2021 IBC and major referenced codes. | src:usa-va:vac-13-5-63 | verified_core_reference |
| Electrical | ahj:usa-va:bhcd | Virginia Board of Housing and Community Development | Administers electrical provisions through VCC referenced NFPA 70 / NEC. | DHCD current-code page and 13VAC5-63 references. | src:usa-va:dhcd-codes; src:usa-va:vac-13-5-63 | verified_core_reference |
| Energy | ahj:usa-va:bhcd | Virginia Board of Housing and Community Development | Administers energy provisions through VCC referenced IECC and related Virginia provisions. | 13VAC5-63 incorporation of 2021 IBC and major referenced codes. | src:usa-va:vac-13-5-63 | verified_core_reference |
| Fire - construction references | ahj:usa-va:bhcd | Virginia Board of Housing and Community Development | Construction fire/life-safety requirements are administered through the VCC and referenced IBC/IFC-related provisions. | 13VAC5-63 and Code of Virginia §§ 36-98, 36-99. | src:usa-va:vac-13-5-63; src:usa-va:code-36-99 | verified_core_reference |
| Fire - operational / prevention code | ahj:usa-va:bhcd-fire | Virginia Board of Housing and Community Development, with Fire Services Board cooperation | Adopts/promulgates the Statewide Fire Prevention Code; SFPC is cooperatively developed with the Fire Services Board. | Code of Virginia § 27-97; 13VAC5-52. | src:usa-va:code-27-97; src:usa-va:vac-13-5-52-30; src:usa-va:register-sfpc-2021-final | verified |
| Accessibility | ahj:usa-va:bhcd | Virginia Board of Housing and Community Development | Adopts accessibility provisions within the USBC/VCC and referenced standards. | Code of Virginia § 36-99; 13VAC5-63 documents incorporated by reference. | src:usa-va:code-36-99; src:usa-va:vac-13-5-63-dir | verified_core_reference |
| Elevator / Conveyance | ahj:usa-va:bhcd | Virginia Board of Housing and Community Development | Adopts elevator/conveyance provisions within the Building Code; local governing bodies inspect/enforce elevators, escalators, and related conveyances for existing buildings. | Code of Virginia §§ 36-105, 36-105.01; 13VAC5-63 documents incorporated by reference. | src:usa-va:code-36-105; src:usa-va:code-36-105-01; src:usa-va:vac-13-5-63-dir | partially_verified |
| Appeals / interpretations | ahj:usa-va:sbctrb | State Building Code Technical Review Board | Hears appeals from enforcement actions under the USBC, SFPC, and related regulations; provides interpretations. | Code of Virginia Article 2 of Chapter 6; DHCD SBCTRB page. | src:usa-va:code-title36-ch6; src:usa-va:dhcd-sbctrb | verified_core_role |
| State Fire Marshal enforcement fallback | ahj:usa-va:sfmo | Virginia State Fire Marshal | Enforces SFPC in jurisdictions where local governments do not enforce it; also may enforce in cooperation with local governments. | Code of Virginia § 27-98. | src:usa-va:code-27-98 | verified_core_role |

### 2.3 Authority Hierarchy Notes

The statewide adopting authority and local enforcement model are separate. The Board of Housing and Community Development adopts/amends the USBC and SFPC. DHCD administers the code program and supports the code development process. Local building departments enforce construction and rehabilitation provisions of the Building Code. Local fire-code enforcement may be local or, where a locality does not enforce the SFPC, through the State Fire Marshal. Appeals generally begin at a local board when local enforcement is involved and proceed to the State Building Code Technical Review Board if needed.

This report does not fully trace state-owned building procedures, Department of General Services roles, industrialized/modular building programs, manufactured housing rules, or amusement-device regulations.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-va:001 | ahj:usa-va:bhcd | adopts_and_amends | Virginia Uniform Statewide Building Code / Virginia Construction Code | src:usa-va:code-36-98; src:usa-va:dhcd-usbc | verified |
| edge:usa-va:002 | ahj:usa-va:bhcd | supersedes | local building codes and regulations, subject to statutory exceptions | src:usa-va:code-36-98 | verified |
| edge:usa-va:003 | ahj:usa-va:bhcd | bases_technical_requirements_on | nationally accepted model codes and standards | src:usa-va:dhcd-usbc; src:usa-va:code-36-99 | verified |
| edge:usa-va:004 | ahj:usa-va:local-building-departments | enforces | construction and rehabilitation provisions of Building Code | src:usa-va:code-36-105; src:usa-va:dhcd-usbc | verified |
| edge:usa-va:005 | ahj:usa-va:local-building-departments | first_level_appeals_to | local board of Building Code appeals | src:usa-va:code-36-105 | verified |
| edge:usa-va:006 | ahj:usa-va:local-building-appeals-board | further_appeals_to | ahj:usa-va:sbctrb | src:usa-va:code-36-105; src:usa-va:dhcd-sbctrb | verified |
| edge:usa-va:007 | ahj:usa-va:bhcd-fire | cooperatively_develops_with | Virginia Fire Services Board / SFPC | src:usa-va:code-27-97; src:usa-va:register-sfpc-2021-final | verified |
| edge:usa-va:008 | ahj:usa-va:local-fire-enforcing-agency | may_enforce | SFPC in entirety or selected statutory categories | src:usa-va:code-27-98 | verified |
| edge:usa-va:009 | ahj:usa-va:sfmo | fallback_enforces | SFPC in jurisdictions without local enforcement | src:usa-va:code-27-98 | verified |
| edge:usa-va:010 | ahj:usa-va:local-governments | may_adopt_more_restrictive_fire_prevention_regulations | local fire-prevention scope excluding construction manner/materials | src:usa-va:code-27-97 | verified |
| edge:usa-va:011 | ahj:usa-va:sbctrb | hears_appeals_from | USBC/SFPC enforcement actions after required local appeal path | src:usa-va:dhcd-sbctrb; src:usa-va:code-27-98; src:usa-va:code-36-105 | verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | Virginia Construction Code / USBC Part I | International Building Code | 2021 | current | null | 2024-01-18 | 2024-01-18 | 2025-01-18 | One-year permit-application choice period; prior edition allowed for applications submitted from 2024-01-18 through the one-year transition window. | src:usa-va:dhcd-codes; src:usa-va:vac-13-5-63; src:usa-va:register-usbc-2021-final |
| Residential | Virginia Residential Code provisions within VCC | International Residential Code | 2021 | current | null | 2024-01-18 | 2024-01-18 | 2025-01-18 | Same VCC permit-application transition where applicable. | src:usa-va:dhcd-codes; src:usa-va:vac-13-5-63 |
| Existing Building / Rehabilitation | Virginia Existing Building Code / USBC Part II | International Existing Building Code | 2021 | current | null | 2024-01-18 | 2024-01-18 | 2025-01-18 | Transition treated as tied to USBC permit-application logic; applicability should be checked project-by-project. | src:usa-va:vac-13-5-63; src:usa-va:register-usbc-2021-final |
| Mechanical | Virginia Construction Code referenced mechanical provisions | International Mechanical Code | 2021 | current_reference | null | 2024-01-18 | 2024-01-18 | 2025-01-18 | Same VCC permit-application transition where applicable. | src:usa-va:vac-13-5-63 |
| Plumbing | Virginia Construction Code referenced plumbing provisions | International Plumbing Code | 2021 | current_reference | null | 2024-01-18 | 2024-01-18 | 2025-01-18 | Same VCC permit-application transition where applicable. | src:usa-va:vac-13-5-63 |
| Fuel Gas | Virginia Construction Code referenced fuel-gas provisions | International Fuel Gas Code | 2021 | current_reference | null | 2024-01-18 | 2024-01-18 | 2025-01-18 | Same VCC permit-application transition where applicable. | src:usa-va:vac-13-5-63 |
| Electrical | Virginia Construction Code referenced electrical provisions | NFPA 70 / National Electrical Code | 2020 | current_reference | null | 2024-01-18 | 2024-01-18 | 2025-01-18 | Same VCC permit-application transition where applicable. | src:usa-va:dhcd-codes; src:usa-va:vac-13-5-63 |
| Energy | Virginia Energy Conservation Code provisions referenced through VCC | International Energy Conservation Code | 2021 | current_reference | null | 2024-01-18 | 2024-01-18 | 2025-01-18 | Same VCC permit-application transition where applicable. | src:usa-va:vac-13-5-63 |
| Fire - construction references | VCC fire/life-safety construction provisions | International Building Code and related referenced standards | 2021 | current | null | 2024-01-18 | 2024-01-18 | 2025-01-18 | Construction fire/life-safety provisions follow VCC permit-application transition. | src:usa-va:vac-13-5-63 |
| Fire - operational / prevention code | Virginia Statewide Fire Prevention Code | International Fire Code | 2021 | current | null | 2024-01-18 | 2024-01-18 | 2024-01-18 | No one-year SFPC transition was captured; report treats effective date as operative/mandatory unless later source indicates otherwise. | src:usa-va:code-27-97; src:usa-va:vac-13-5-52-30; src:usa-va:register-sfpc-2021-final |
| Accessibility | VCC accessibility provisions and incorporated accessibility standards | 2021 IBC accessibility provisions; ICC/ANSI A117.1-17 | 2021 / 2017 | current_reference | null | 2024-01-18 | 2024-01-18 | 2025-01-18 | Same VCC permit-application transition where applicable. | src:usa-va:code-36-99; src:usa-va:vac-13-5-63; src:usa-va:vac-13-5-63-dir |
| Elevator / Conveyance | VCC elevator and conveying-system provisions | 2021 IBC Chapter 30; ASME A17.1/CSA B44-19; ASME A17.3-2008; ASME A18.1-2008 | mixed | partially_verified | null | 2024-01-18 | 2024-01-18 | null | New-construction provisions appear tied to VCC timing; existing-elevator inspection timing requires deeper parsing. | src:usa-va:code-36-105; src:usa-va:code-36-105-01; src:usa-va:vac-13-5-63; src:usa-va:vac-13-5-63-dir |

### 3.2 Adoption Records

```yaml
- record_id: adoption:usa-va:vcc-2021
  code_family: Building
  state_code_name: Virginia Construction Code / USBC Part I
  base_model_code: International Building Code
  edition: "2021"
  authority_id: ahj:usa-va:bhcd
  adoption_date: null
  final_publication_date: "2023-12-18"
  effective_date: "2024-01-18"
  operative_date: "2024-01-18"
  mandatory_date: "2025-01-18"
  transition_rule_id: date-rule:usa-va:001
  source_ids:
    - src:usa-va:dhcd-codes
    - src:usa-va:vac-13-5-63
    - src:usa-va:register-usbc-2021-final

- record_id: adoption:usa-va:vrc-2021
  code_family: Residential
  state_code_name: Virginia Residential Code provisions within VCC
  base_model_code: International Residential Code
  edition: "2021"
  authority_id: ahj:usa-va:bhcd
  adoption_date: null
  final_publication_date: "2023-12-18"
  effective_date: "2024-01-18"
  operative_date: "2024-01-18"
  mandatory_date: "2025-01-18"
  transition_rule_id: date-rule:usa-va:001
  source_ids:
    - src:usa-va:dhcd-codes
    - src:usa-va:vac-13-5-63

- record_id: adoption:usa-va:vebc-2021
  code_family: Existing Building / Rehabilitation
  state_code_name: Virginia Existing Building Code / USBC Part II
  base_model_code: International Existing Building Code
  edition: "2021"
  authority_id: ahj:usa-va:bhcd
  adoption_date: null
  final_publication_date: "2023-12-18"
  effective_date: "2024-01-18"
  operative_date: "2024-01-18"
  mandatory_date: "2025-01-18"
  transition_rule_id: date-rule:usa-va:001
  source_ids:
    - src:usa-va:vac-13-5-63
    - src:usa-va:register-usbc-2021-final

- record_id: adoption:usa-va:referenced-trades-2021
  code_family:
    - Mechanical
    - Plumbing
    - Fuel Gas
    - Electrical
    - Energy
  state_code_name: VCC referenced technical codes
  base_model_code:
    - 2021 International Mechanical Code
    - 2021 International Plumbing Code
    - 2021 International Fuel Gas Code
    - 2020 NFPA 70 / National Electrical Code
    - 2021 International Energy Conservation Code
  edition: mixed
  authority_id: ahj:usa-va:bhcd
  adoption_date: null
  final_publication_date: "2023-12-18"
  effective_date: "2024-01-18"
  operative_date: "2024-01-18"
  mandatory_date: "2025-01-18"
  transition_rule_id: date-rule:usa-va:001
  source_ids:
    - src:usa-va:dhcd-codes
    - src:usa-va:vac-13-5-63

- record_id: adoption:usa-va:sfpc-2021
  code_family: Fire - operational / prevention code
  state_code_name: Virginia Statewide Fire Prevention Code
  base_model_code: International Fire Code
  edition: "2021"
  authority_id: ahj:usa-va:bhcd-fire
  adoption_date: null
  final_publication_date: "2023-12-18"
  effective_date: "2024-01-18"
  operative_date: "2024-01-18"
  mandatory_date: "2024-01-18"
  transition_rule_id: date-rule:usa-va:003
  source_ids:
    - src:usa-va:code-27-97
    - src:usa-va:vac-13-5-52-30
    - src:usa-va:register-sfpc-2021-final
```

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

Virginia's 2021 USBC/SFPC effective date is `2024-01-18`. For construction permit applications subject to the VCC transition clause, a one-year period beginning on the 2021 code effective date allows the applicant to choose either the 2021 code or the immediately prior edition. The report calculates the end of that one-year window as `2025-01-18`; legal confirmation of inclusive/exclusive day counting remains an open implementation detail. Properly issued permits under a prior edition are protected from required changes to approved construction documents, design, or construction if the permit is not suspended or revoked.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-va:001 | VCC / USBC construction permit applications and related referenced construction codes | permit_application_grace_period | 2024-01-18/2025-01-18 | Permit application submitted to local building department during one-year period beginning on effective date of 2021 edition. | yes | src:usa-va:vac-13-5-63; src:usa-va:register-usbc-2021-final | verified_core_rule |
| date-rule:usa-va:002 | Previously issued permits under prior USBC editions | vested_prior_permit | null | Permit properly issued under a previous edition and not suspended or revoked. | yes | src:usa-va:vac-13-5-63 | verified_core_rule |
| date-rule:usa-va:003 | Statewide Fire Prevention Code | effective_date | 2024-01-18 | Final regulation effective date for 13VAC5-52 SFPC. | no_grace_period_captured | src:usa-va:register-sfpc-2021-final; src:usa-va:vac-13-5-52-30 | partially_verified |
| date-rule:usa-va:004 | Future code cycle | monitor | null | 2024 Code Development Cycle advances to final regulation / effective date. | unresolved | src:usa-va:dhcd-2024-cdc; src:usa-va:dhcd-bhcd | monitoring |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| USBC / SFPC and related regulations | 2024 Virginia codes under development | null | null | null | null | null | active_monitoring | src:usa-va:dhcd-2024-cdc; src:usa-va:dhcd-bhcd | DHCD has an active 2024 Code Development Cycle page with base documents and stakeholder meetings; BHCD page lists May 11, 2026 code-cycle materials. No final effective date captured. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-va:001 | Residential | Detached one-family and two-family dwellings and townhouses | Residential scope under VCC | The IRC is applicable to detached one-family and two-family dwellings and townhouses as set out in VCC Section 310. | src:usa-va:vac-13-5-63 | verified_core_rule |
| applicability-rule:usa-va:002 | Existing Building / Rehabilitation | Construction and rehabilitation in existing buildings | Existing-building activity except where specifically addressed in VCC | USBC Part II / VEBC applies to construction and rehabilitation activities in existing buildings and structures, except where specifically addressed in the VCC. | src:usa-va:vac-13-5-63 | verified_core_rule |
| applicability-rule:usa-va:003 | Operational fire prevention | Maintenance and operational fire-safety functions | Existing structures and operational fire-safety activities | SFPC governs maintenance of fire-protection aspects of existing structures and operational functions relating to fire safety. | src:usa-va:register-sfpc-2021-final; src:usa-va:code-27-97 | verified_core_rule |
| applicability-rule:usa-va:004 | Farm buildings and structures | Farm building exemption | Farm building or structure not otherwise excepted | Farm buildings and structures are generally exempt from the Building Code, with statutory and regulatory exceptions such as farm restaurants and flood/mudslide conditions. | src:usa-va:code-36-99; src:usa-va:vac-13-5-63 | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** Virginia incorporates model codes by reference and then publishes Virginia-specific administrative provisions and amendments in the Virginia Administrative Code. The VCC is comprised of Chapter 1 administration, incorporated 2021 IBC Chapters 2-35, and specifically identified changes to the incorporated IBC chapters. The VEBC and VPMC follow the same general structure for the IEBC and IPMC. The SFPC incorporates the 2021 IFC, deletes IFC Chapter 1, replaces administrative/procedural matters with SFPC Chapter 1, and applies Virginia BHCD amendments where they conflict with unchanged IFC or referenced-standard text.

**Where amendments are published:** Virginia Administrative Code Title 13, Agency 5, Chapter 63 for the USBC/VCC/VEBC/VPMC and Chapter 52 for the SFPC. DHCD also links Virginia code collections and state amendment booklets.

**Amendment parsing status:** registry_identified_core_not_fully_parsed

### 5.2 State Amendment Sources

| Amendment Source ID | Scope | Publication Path | Status | Notes |
| --- | --- | --- | --- | --- |
| amendment-source:usa-va:13vac5-63 | USBC / VCC / VEBC / VPMC | Virginia Administrative Code 13VAC5-63 | identified | Full amendment text is long and has not been converted into granular amendment records. |
| amendment-source:usa-va:13vac5-63-dir | Documents incorporated by reference for USBC | Virginia Administrative Code 13VAC5-63 documents incorporated by reference | identified | Confirms major model-code editions and standards, including ICC, NFPA 70, ICC/ANSI A117.1, and ASME elevator standards. |
| amendment-source:usa-va:13vac5-52 | SFPC | Virginia Administrative Code 13VAC5-52 | identified | Core incorporation and administrative replacement verified; granular amendments not fully parsed. |
| amendment-source:usa-va:13vac5-52-dir | Documents incorporated by reference for SFPC | Virginia Administrative Code 13VAC5-52 documents incorporated by reference | identified | Confirms IFC 2021 and fire-related referenced standards. |
| amendment-source:usa-va:dhcd-regulations-archive | DHCD Virginia Building and Fire Regulations | DHCD code/archive and amendment booklets | identified | Useful for publication monitoring; not parsed into amendment records in this report. |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Topic | Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| amendment:usa-va:vcc-ch1-admin | Building / construction | Administrative chapter replacement | VCC Chapter 1 is established in Virginia administrative text; IBC Chapter 1 is not incorporated as part of the USBC except where Virginia text gives meaning to a referenced administrative provision. | src:usa-va:vac-13-5-63 | verified_core_amendment |
| amendment:usa-va:vebc-ch1-admin | Existing Building / Rehabilitation | Administrative chapter replacement | VEBC Chapter 1 is established by Virginia; IEBC Chapters 2-16 are incorporated, with state amendments and order-of-precedence rules. | src:usa-va:vac-13-5-63 | verified_core_amendment |
| amendment:usa-va:vpmc-ch1-admin | Maintenance | Administrative chapter replacement | VPMC Chapter 1 is established by Virginia; IPMC Chapters 2-8 are incorporated, with state amendments and order-of-precedence rules. | src:usa-va:vac-13-5-63 | verified_core_amendment |
| amendment:usa-va:sfpc-ifc-ch1-delete | Fire - operational / prevention | Administrative chapter replacement | SFPC incorporates the 2021 IFC, deletes IFC Chapter 1, and replaces administrative/procedural provisions with SFPC Chapter 1. | src:usa-va:vac-13-5-52-30 | verified_core_amendment |
| amendment:usa-va:vcc-transition | Building / construction | One-year transition | Permit applicants during the one-year period beginning `2024-01-18` may choose 2021 code or the immediately prior edition. | src:usa-va:vac-13-5-63 | verified_core_amendment |
| amendment:usa-va:elevator-ch30 | Elevator / Conveyance | Elevator and conveying systems | VCC Chapter 30 includes state amendments, and documents incorporated by reference include ASME A17.1/CSA B44-19, ASME A17.3-2008, and ASME A18.1-2008. | src:usa-va:vac-13-5-63; src:usa-va:vac-13-5-63-dir | partially_verified |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-va"
  model: "statewide_code_local_building_department_enforcement"
  enforcing_entities:
    - "local building department for construction and rehabilitation"
    - "local governing body-designated agency or department for existing-building inspection/enforcement where authorized"
    - "local governing body for elevator/escalator/related conveyance inspection/enforcement, carried out by designated agency or department"
    - "local fire enforcing agency for SFPC where locality elects to enforce"
    - "State Fire Marshal for SFPC where local governments do not enforce"
  required_officials:
    - "local building department"
    - "local board of Building Code appeals or approved appeals agreement"
    - "local fire official / local board of fire prevention code appeals where locality enforces SFPC"
  state_reserved_activities:
    - "adoption and amendment of USBC by Board of Housing and Community Development"
    - "adoption and amendment of SFPC by Board of Housing and Community Development, cooperatively developed with Fire Services Board"
    - "State Building Code Technical Review Board appeals and interpretations"
    - "State Fire Marshal enforcement fallback for SFPC in non-enforcing localities"
  source_ids:
    - "src:usa-va:code-36-105"
    - "src:usa-va:dhcd-usbc"
    - "src:usa-va:code-27-98"
    - "src:usa-va:dhcd-sbctrb"
  verification_status: "verified_core_model"
  confidence: 0.84
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-va"
  model: "building_code_statewide_preemption_with_limited_fire_prevention_local_authority"
  applies_to_code_families:
    - "USBC / VCC / VEBC / VPMC: statewide preemption of local building codes and regulations, subject to statutory exceptions"
    - "SFPC: local governments may adopt more restrictive or more extensive fire prevention regulations, constrained by construction/material limits"
  approval_required: null
  approving_authority_id: null
  filing_required: null
  registry_exists: null
  registry_source_ids: []
  legal_basis_source_ids:
    - "src:usa-va:code-36-98"
    - "src:usa-va:code-27-97"
    - "src:usa-va:vac-13-5-52-30"
  verification_status: "partially_verified"
  confidence: 0.78
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Virginia's local enforcement model should not be confused with local code-amendment authority. Local building departments enforce the statewide Building Code, but the USBC supersedes local building codes and regulations. Fire prevention has a broader local-ordinance pathway: local governments may adopt fire prevention regulations that are more restrictive or more extensive than the SFPC, but the statute preserves the USBC's construction/material domain by prohibiting local fire-prevention regulations from affecting the manner of construction or materials used in construction, alteration, repair, or use of buildings or structures.

### 6.4 Known Local Amendment Registries

No statewide registry of local building-code amendments or local fire-prevention ordinances was identified in this report. Because the USBC is statewide and preemptive, any local building-code amendment registry should not be assumed to exist. Local fire-prevention ordinances may exist locally and should be resolved at the city/county/town level.

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendments or local fire-prevention ordinances were parsed in this report.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: not_started

Jurisdiction stack:

```text
Address
  -> Virginia
  -> County / independent city
  -> Municipality / town status / unincorporated county area
  -> Local building department or approved enforcement agreement
  -> Local board of Building Code appeals or approved appeals agreement
  -> Local fire enforcing agency if locality enforces SFPC
  -> State Fire Marshal fallback if locality does not enforce SFPC
  -> State Building Code Technical Review Board for final administrative appeals after required local path
  -> Applicable statewide code adoption records
  -> Applicable local fire-prevention ordinance records, if any
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | not selected | none | statewide | unknown | pending |
| County / independent city | not selected | none | statewide | unknown | pending |
| Municipality / town | not selected | none | statewide | unknown | pending |
| Fire enforcement locality status | not selected | none | statewide | unknown | pending |
| Special district | not selected | none | statewide | unknown | pending |

### 7.3 AHJ Contact Data

No AHJ contact dataset was populated. For production lookup, collect local building department contacts, local fire-code enforcement status, local boards of appeals, and any local fire-prevention ordinance source for each locality.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Source Type | Title | Publisher / Agency | URL | Accessed | Key Fields Supported |
| --- | --- | --- | --- | --- | --- | --- |
| src:usa-va:dhcd-codes | agency_page | Codes | Virginia Department of Housing and Community Development | https://www.dhcd.virginia.gov/codes | 2026-06-26 | Current effective date; 2021 I-Codes; 2020 NEC; DHCD/BHCD code-adoption program. |
| src:usa-va:dhcd-usbc | agency_page | Virginia Uniform Statewide Building Code (USBC) | Virginia Department of Housing and Community Development | https://www.dhcd.virginia.gov/virginia-uniform-statewide-building-code-usbc | 2026-06-26 | USBC parts; Board adopts/amends; local building inspections enforcement; appeals overview. |
| src:usa-va:code-36-98 | statute | Code of Virginia § 36-98, Board to promulgate Statewide Code; other codes and regulations superseded; exceptions | Virginia Law / Division of Legislative Automated Systems | https://law.lis.virginia.gov/vacode/title36/chapter6/section36-98/ | 2026-06-26 | Primary building-code authority; statewide preemption. |
| src:usa-va:code-36-99 | statute | Code of Virginia § 36-99, Provisions of Code; modifications | Virginia Law / Division of Legislative Automated Systems | https://law.lis.virginia.gov/vacode/title36/chapter6/section36-99/ | 2026-06-26 | Code content, administration/enforcement procedures, model-code standards, accessibility/health/safety policy. |
| src:usa-va:code-36-105 | statute | Code of Virginia § 36-105, Enforcement of Code; appeals; inspections; permits | Virginia Law / Division of Legislative Automated Systems | https://law.lis.virginia.gov/vacode/title36/chapter6/section36-105/ | 2026-06-26 | Local building department enforcement; local appeals board; small-town rule; inspection/fees; existing-building and elevator enforcement. |
| src:usa-va:code-36-105-01 | statute | Code of Virginia § 36-105.01, Elevator inspections by contract | Virginia Law / Division of Legislative Automated Systems | https://law.lis.virginia.gov/vacode/title36/chapter6/section36-105.01/ | 2026-06-26 | Existing elevator inspection and approved-agency certification authority. |
| src:usa-va:code-36-105-1 | statute | Code of Virginia § 36-105.1, Inspection and review of plans of buildings under construction | Virginia Law / Division of Legislative Automated Systems | https://law.lis.virginia.gov/vacode/title36/chapter6/section36-105.1/ | 2026-06-26 | Local inspectors' plan review and inspection role; state fire safety transition after completion in non-SFPC-enforcing localities. |
| src:usa-va:code-title36-ch6 | statute_index | Code of Virginia Title 36, Chapter 6, Uniform Statewide Building Code | Virginia Law / Division of Legislative Automated Systems | https://law.lis.virginia.gov/vacode/title36/chapter6/ | 2026-06-26 | Chapter map; SBCTRB article and appeal provisions for follow-up. |
| src:usa-va:vac-13-5-63 | regulation | 13VAC5-63, Virginia Uniform Statewide Building Code | Virginia Administrative Code / Virginia Law | https://law.lis.virginia.gov/admincodefull/title13/agency5/chapter63/ | 2026-06-26 | VCC/VEBC/VPMC structure; IBC/IEBC/IPMC incorporation; referenced major codes; transition rule; state amendments. |
| src:usa-va:vac-13-5-63-dir | regulation | Documents Incorporated by Reference (13VAC5-63) | Virginia Administrative Code / Virginia Law | https://law.lis.virginia.gov/admincode/title13/agency5/chapter63/section9999/ | 2026-06-26 | Model code and standard editions, including IBC, IECC, IEBC, IFC, IFGC, IMC, IPMC, IPC, IRC, NFPA 70, ICC/ANSI A117.1, ASME elevator standards. |
| src:usa-va:register-usbc-2021-final | register_notice | Vol. 40 Issue 9 Final Regulation, 13VAC5-63 Virginia Uniform Statewide Building Code | Virginia Register of Regulations | https://register.dls.virginia.gov/details.aspx?id=10921 | 2026-06-26 | Final regulation publication; statutory authority; effective date `2024-01-18`; code-development process summary. |
| src:usa-va:code-27-97 | statute | Code of Virginia § 27-97, Adoption of Fire Prevention Code | Virginia Law / Division of Legislative Automated Systems | https://law.lis.virginia.gov/vacode/title27/chapter9/section27-97/ | 2026-06-26 | SFPC adopting authority; cooperative development with Fire Services Board; local fire-prevention ordinance limits. |
| src:usa-va:code-27-98 | statute | Code of Virginia § 27-98, Enforcement of Fire Prevention Code; appeals | Virginia Law / Division of Legislative Automated Systems | https://law.lis.virginia.gov/vacode/title27/chapter9/section27-98/ | 2026-06-26 | Local SFPC enforcement election; State Fire Marshal fallback; appeals path; fees/inspection procedures. |
| src:usa-va:vac-13-5-52-30 | regulation | 13VAC5-52-30, SFPC Section 103 Incorporation by reference | Virginia Administrative Code / Virginia Law | https://law.lis.virginia.gov/admincode/title13/agency5/chapter52/section30/ | 2026-06-26 | 2021 IFC incorporation; IFC Chapter 1 deletion; state amendments and order of precedence; no retroactive IFC fire-protection-system enforcement unless specified by USBC. |
| src:usa-va:vac-13-5-52-dir | regulation | Documents Incorporated by Reference (13VAC5-52) | Virginia Administrative Code / Virginia Law | https://law.lis.virginia.gov/admincode/title13/agency5/chapter52/section9999/ | 2026-06-26 | SFPC incorporated standards including 2021 IFC and fire-protection NFPA standards. |
| src:usa-va:register-sfpc-2021-final | register_notice | Vol. 40 Issue 9 Final Regulation, 13VAC5-51/52 Virginia Statewide Fire Prevention Code | Virginia Register of Regulations | https://register.dls.virginia.gov/details.aspx?id=10917 | 2026-06-26 | SFPC final action; repeal/replace of 13VAC5-51 with 13VAC5-52; statutory authority; effective date `2024-01-18`; SFPC scope. |
| src:usa-va:dhcd-sbctrb | agency_page | State Building Code Technical Review Board (SBCTRB) | Virginia Department of Housing and Community Development | https://www.dhcd.virginia.gov/state-building-code-technical-review-board-sbctrb | 2026-06-26 | Appeals and interpretations for USBC/SFPC and related regulations. |
| src:usa-va:dhcd-2024-cdc | agency_page | 2024 Code Development Cycle | Virginia Department of Housing and Community Development | https://www.dhcd.virginia.gov/2024-code-development-cycle | 2026-06-26 | Pending 2024 code cycle; workgroups; base documents; stakeholder meeting dates. |
| src:usa-va:dhcd-bhcd | agency_page | Board of Housing and Community Development (BHCD) | Virginia Department of Housing and Community Development | https://www.dhcd.virginia.gov/board-housing-and-community-development-bhcd | 2026-06-26 | Board composition; 2024 code-development meeting materials; monitoring target. |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-va:dhcd-codes | agency_page | Agency page is concise and current-facing; use with VAC/Register sources for legal text. | acceptable_with_primary_crosscheck |
| src:usa-va:dhcd-usbc | agency_page | Agency summary; not a substitute for statutory/VAC text. | acceptable_for_program_summary |
| src:usa-va:vac-13-5-63 | very_long_html | Full VAC chapter is long; this report extracted core sections but did not parse every amendment. | requires_granular_amendment_pass |
| src:usa-va:vac-13-5-63-dir | incorporation_list | The VAC warns that website addresses for incorporated documents are convenience links and the described source document should be used for accuracy. | use_for_edition_identification_not_full_model_code_text |
| src:usa-va:vac-13-5-52-30 | regulation_section | Core SFPC incorporation section verified; full Chapter 52 amendments not fully parsed. | requires_granular_amendment_pass |
| src:usa-va:register-usbc-2021-final | register_notice | Final notice supports effective date and regulatory action history; current VAC text controls for current codified language. | use_for_effective_date_and_history |
| src:usa-va:register-sfpc-2021-final | register_notice | Final notice supports effective date and repeal/replace history; current VAC text controls for current codified language. | use_for_effective_date_and_history |
| src:usa-va:dhcd-2024-cdc | pending_cycle_page | Indicates ongoing code-development activity but not a final adopted 2024-code effective date. | monitor_for_final_action |
| src:usa-va:dhcd-bhcd | pending_cycle_page | Board materials may include PDFs and drafts; this report did not parse individual packet PDFs. | monitor_and_parse_when_needed |

### 8.3 Supplemental Sources

None used in this report.

### 8.4 Source Extraction Metadata

| Source ID | Extraction Method | Extracted Fields | Extracted On | Notes |
| --- | --- | --- | --- | --- |
| src:usa-va:dhcd-codes | web_html | effective date, current code cycle, DHCD/BHCD program summary | 2026-06-26 | Official agency page. |
| src:usa-va:dhcd-usbc | web_html | USBC parts, Board adoption role, local enforcement summary | 2026-06-26 | Official agency page. |
| src:usa-va:code-36-98 | web_html | primary authority and preemption | 2026-06-26 | Primary statute. |
| src:usa-va:code-36-99 | web_html | code purpose, procedures, model-code standards, accessibility policy | 2026-06-26 | Primary statute. |
| src:usa-va:code-36-105 | web_html | local enforcement and appeals | 2026-06-26 | Primary statute. |
| src:usa-va:code-27-97 | web_html | SFPC adoption and local fire ordinance scope | 2026-06-26 | Primary statute. |
| src:usa-va:code-27-98 | web_html | SFPC enforcement and State Fire Marshal fallback | 2026-06-26 | Primary statute. |
| src:usa-va:vac-13-5-63 | web_html | VCC/VEBC/VPMC incorporation and transition | 2026-06-26 | Current VAC full chapter. |
| src:usa-va:vac-13-5-52-30 | web_html | SFPC IFC incorporation and administrative replacement | 2026-06-26 | Current VAC section. |
| src:usa-va:register-usbc-2021-final | web_html | final regulation effective date and authority | 2026-06-26 | Virginia Register notice. |
| src:usa-va:register-sfpc-2021-final | web_html | final regulation effective date and authority | 2026-06-26 | Virginia Register notice. |
| src:usa-va:dhcd-2024-cdc | web_html | pending 2024 cycle monitoring | 2026-06-26 | No final effective date captured. |
| src:usa-va:dhcd-bhcd | web_html | pending 2024 cycle board materials | 2026-06-26 | No PDF packet parsing performed. |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 1.00 | none | Core authority and adoption fields have official source support; full amendment/AHJ work remains open. |
| report | risk.overall_confidence | 0.72 | verified | 0.80 | none | Confidence reflects strong core-state support but incomplete amendment and AHJ coverage. |
| ahj:usa-va:bhcd | authority_name | Virginia Board of Housing and Community Development | verified | 0.95 | src:usa-va:code-36-98; src:usa-va:dhcd-usbc | Primary adopting authority. |
| ahj:usa-va:bhcd | legal_basis | Code of Virginia § 36-98 | verified | 0.95 | src:usa-va:code-36-98 | Statute directs and empowers Board to adopt/promulgate USBC. |
| adoption:usa-va:vcc-2021 | effective_date | 2024-01-18 | verified | 0.95 | src:usa-va:dhcd-codes; src:usa-va:register-usbc-2021-final | DHCD and Virginia Register align. |
| adoption:usa-va:vcc-2021 | base_model_code | 2021 International Building Code | verified | 0.95 | src:usa-va:vac-13-5-63 | VCC incorporates IBC Chapters 2-35. |
| adoption:usa-va:referenced-trades-2021 | electrical_base | 2020 NFPA 70 / NEC | verified | 0.90 | src:usa-va:dhcd-codes; src:usa-va:vac-13-5-63 | DHCD states 2020 National Electrical Code; VCC references NFPA 70. |
| adoption:usa-va:sfpc-2021 | base_model_code | 2021 International Fire Code | verified | 0.95 | src:usa-va:vac-13-5-52-30 | SFPC incorporates 2021 IFC. |
| date-rule:usa-va:001 | transition_period | 2024-01-18/2025-01-18 | partially_verified | 0.82 | src:usa-va:vac-13-5-63 | Source says one-year period beginning effective date; end date calculated. |
| local-enforcement:usa-va | model | statewide_code_local_building_department_enforcement | verified | 0.88 | src:usa-va:code-36-105; src:usa-va:dhcd-usbc | Local building department enforcement verified. |
| local-amendment-rule:usa-va | building_preemption | USBC supersedes local building codes/regulations | verified | 0.90 | src:usa-va:code-36-98 | Statutory preemption verified. |
| local-amendment-rule:usa-va | fire_local_regulation_scope | More restrictive/extensive local fire prevention regulations allowed within statutory construction/material limits | verified | 0.86 | src:usa-va:code-27-97 | Scope verified; local registry not identified. |
| watch:usa-va:2024-code-cycle | status | active_monitoring | verified | 0.76 | src:usa-va:dhcd-2024-cdc; src:usa-va:dhcd-bhcd | Current-cycle activity found; final adoption/effective date not captured. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | All `src:usa-va:*` identifiers used in the report are registered in section 8. Rows using `none` are unresolved data placeholders, not source identifiers. |
| All authority IDs resolve | pass | Authority IDs are defined in sections 2 and 6. |
| All current code families have adoption records | pass | Matrix rows and normalized records cover verified current statewide code families; some are bundled as referenced technical codes. |
| Building and operational fire code are separated | pass | VCC construction fire/life-safety and SFPC operational/prevention code are tracked separately. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Adoption date remains `null` where Board adoption date was not extracted; final publication date is kept separate in normalized records. |
| Effective dates are valid ISO dates | pass | Entered effective dates use `YYYY-MM-DD`; missing dates use `null`. |
| No impossible date sequences | pass | No date sequence conflict was introduced; one mandatory date is calculated from the one-year transition and flagged. |
| Transition rules have explicit trigger conditions | pass | Permit-application trigger and prior-permit protection are captured. |
| Permit-date logic is captured where applicable | pass | VCC permit-application transition is captured. |
| Local enforcement model classified | pass | Local building department enforcement and SFPC local/state fallback are classified. |
| Local amendment rule classified | pass | Building-code preemption and limited local fire-prevention authority are classified. |
| AHJ confirmation metadata present | fail | AHJ contacts and boundary data remain unpopulated. |
| Official-source caveats captured | pass | Agency-page, VAC, Register, and pending-cycle caveats are recorded. |
| Full state amendment parsing complete | fail | Only core amendment structure and selected high-impact amendments were captured. |
| Pending 2024 code cycle checked | pass | Monitoring source identified; final adoption/effective fields remain null until final action. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-va:001 | high | state amendments | Full Virginia amendments in 13VAC5-63 and 13VAC5-52 were not parsed into granular amendment records. | Extract amendment deltas by code family and section; identify high-impact deviations from model codes. | null | null | open |
| issue:usa-va:002 | medium | local fire ordinances | Local fire-prevention regulations may be more restrictive or broader than SFPC within statutory limits, but no statewide local registry was identified. | Determine whether local fire-prevention ordinances can be collected from municipal/county code publishers or a state clearinghouse. | null | null | open |
| issue:usa-va:003 | medium | AHJ resolution | No AHJ boundary or contact data was loaded for local building departments, fire officials, or boards of appeals. | Build locality-level AHJ dataset and distinguish counties, independent cities, and towns under 3,500 population. | null | null | open |
| issue:usa-va:004 | medium | state-owned buildings | State-owned building review/enforcement and Department of General Services interactions were not fully traced. | Extract Code of Virginia §§ 36-98.1, 36-98.2, and related administrative provisions. | null | null | open |
| issue:usa-va:005 | medium | pending 2024 cycle | The 2024 Code Development Cycle is active and may change current-code status when final regulations are adopted. | Monitor DHCD, BHCD meeting materials, Virginia Town Hall, and Virginia Register for final 2024 regulations and effective dates. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-va:dhcd-codes | src:usa-va:dhcd-codes | html_diff | monthly | DHCD current-code effective date or edition changes | 2026-06-26 | active |
| watch:usa-va:vac-13-5-63 | src:usa-va:vac-13-5-63 | html_diff | monthly | USBC / VCC / VEBC / VPMC text, historical notes, or incorporated codes change | 2026-06-26 | active |
| watch:usa-va:vac-13-5-52 | src:usa-va:vac-13-5-52-30 | html_diff | monthly | SFPC incorporation, amendments, or effective history changes | 2026-06-26 | active |
| watch:usa-va:register-usbc | src:usa-va:register-usbc-2021-final | register_monitor | monthly | New final or proposed USBC regulation appears | 2026-06-26 | active |
| watch:usa-va:register-sfpc | src:usa-va:register-sfpc-2021-final | register_monitor | monthly | New final or proposed SFPC regulation appears | 2026-06-26 | active |
| watch:usa-va:2024-code-cycle | src:usa-va:dhcd-2024-cdc | html_diff | biweekly | 2024 cycle base documents, summaries, proposals, or final action updates | 2026-06-26 | active |
| watch:usa-va:bhcd | src:usa-va:dhcd-bhcd | html_diff | biweekly | Board meeting materials or minutes indicate adoption / final approval of 2024 codes | 2026-06-26 | active |
| watch:usa-va:local-fire | src:usa-va:code-27-97 | legal_research_queue | quarterly | Local fire-prevention regulation collection method identified | 2026-06-26 | pending |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-24 | Generated baseline draft report | report:usa-va | none | Codex | Baseline contained unresolved placeholders and no official sources. |
| 2026-06-26 | Populated core Virginia authority, current-code adoption, transition, enforcement, local-amendment posture, source registry, and QA sections | report:usa-va; ahj:usa-va:bhcd; adoption:usa-va:vcc-2021; adoption:usa-va:sfpc-2021; local-enforcement:usa-va; local-amendment-rule:usa-va | src:usa-va:dhcd-codes; src:usa-va:dhcd-usbc; src:usa-va:code-36-98; src:usa-va:vac-13-5-63; src:usa-va:code-36-105; src:usa-va:code-27-97; src:usa-va:code-27-98; src:usa-va:vac-13-5-52-30 | OpenAI GPT-5.5 Thinking | Upgraded status to `partially_verified` because core authority/adoption fields are source-backed; kept amendment/AHJ/pending-cycle issues open. |
