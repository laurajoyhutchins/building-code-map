---
state:
  state_id: "US-LA"
  name: "Louisiana"
  abbreviation: "LA"
report:
  report_id: "state-report:usa-la"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-25"
  last_verified: "2026-06-25"
  reviewed_by: null
risk:
  overall_confidence: 0.68 # 0.00 - 1.00
  risk_flags:
    - "scheduled_2026_statutory_changes"
    - "fire_code_transition_requires_follow_up"
    - "elevator_conveyance_authority_unresolved"
    - "local_amendment_registry_unresolved"
    - "consolidated_amendment_source_has_caveat"
  open_questions_count: 5

---

# State Building Code Authority Report: Louisiana

## 1. Executive Summary

- **Authority model:** Louisiana uses a statewide uniform construction-code model. The Louisiana State Uniform Construction Code Council (LSUCCC) is the primary adopting and amendment body for the State Uniform Construction Code, while municipalities and parishes enforce the state code locally using council-certified personnel or certified third-party providers. The Office of State Fire Marshal (OSFM) retains distinct fire, life-safety, accessibility, and fire-prevention authority.

- **Statewide code status:** The core Louisiana State Uniform Construction Code currently includes the 2021 IBC, 2021 IEBC, 2021 IRC, 2021 IMC, 2021 IPC, 2021 IFGC, 2020 NEC, and 2021 IECC with Louisiana amendments. The 2021 I-code / 2020 NEC package is reported by LSUCCC as effective 2023-01-01, with the 2021 energy-code provisions effective 2023-07-01 and subsequent amendment packages, including the 2024-03-20 consolidated amendment document.

- **Local enforcement model:** Local enforcement is mandatory for municipalities and parishes. Local governments must enforce only the construction codes provided in Part IV-A and must use building code enforcement officers or certified third-party providers. Each parish and municipality must appoint a council-certified building official or contract for that function so the unincorporated area of the parish is under a council-certified building official.

- **Local amendment posture:** Local governments may adopt ordinances for local administration and enforcement procedures, but R.S. 40:1730.40 states that local governments may not avoid enforcement or otherwise amend the mandatory construction-code provisions required by R.S. 40:1730.28. Fire-prevention local variances are treated separately: effective 2026-07-01, R.S. 40:1578.7 provides a pathway for more stringent local fire-prevention requirements based on local factors after state fire marshal review and approval.

- **Known transition periods or pending changes:** Several statutory notes are scheduled for 2026-07-01, including changes to R.S. 40:1730.23 enforcement terminology, R.S. 40:1578.7 fire-prevention code language, and R.S. 40:1733 accessibility. These changes should be rechecked after 2026-07-01 before raising status above partially verified.

- **Production readiness:** partial_ready_for_core_building_code_fields

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: LSUCCC is the primary body to review and adopt the State Uniform Construction
    Code and to determine justified amendments.
  confidence: 0.9
  source_ids:
  - src:usa-la:rs-40-1730-22
  - src:usa-la:rs-40-1730-26
- topic: Primary building code edition
  finding: The current building-code adoption is the 2021 IBC, excluding Chapter 1,
    Chapter 11, and Chapter 27, with Louisiana amendments.
  confidence: 0.82
  source_ids:
  - src:usa-la:lac-17-i-ucc-2024-03-20
  - src:usa-la:lsuccc-current-codes-2023-10-19
  - src:usa-la:osfm-codes-rules-laws
- topic: Electrical code authority
  finding: The LSUCCC-adopted construction-code matrix includes the 2020 National
    Electric Code for regulation of construction in Louisiana.
  confidence: 0.78
  source_ids:
  - src:usa-la:lac-17-i-ucc-2024-03-20
  - src:usa-la:rs-40-1730-28
- topic: Fire code authority
  finding: OSFM enforces life-safety, accessibility, and fire-code requirements; current
    OSFM materials list NFPA 101 Life Safety Code 2015 and fire-protection/egress
    provisions of LSUCC codes. R.S. 40:1578.7 separately adopts the State Uniform
    Fire Prevention Code.
  confidence: 0.72
  source_ids:
  - src:usa-la:osfm-codes-rules-laws
  - src:usa-la:rs-40-1578-7
- topic: Local enforcement
  finding: Municipalities and parishes must enforce only the construction codes in
    Part IV-A and must use building code enforcement officers or certified third-party
    providers.
  confidence: 0.88
  source_ids:
  - src:usa-la:rs-40-1730-23
  - src:usa-la:rs-40-1730-24
  - src:usa-la:rs-40-1730-25
- topic: Local amendments
  finding: Local administrative/enforcement ordinances are allowed, but local governments
    may not avoid enforcement or amend mandatory construction-code provisions required
    by R.S. 40:1730.28.
  confidence: 0.8
  source_ids:
  - src:usa-la:rs-40-1730-40
- topic: Effective / operative date rule
  finding: The 2021 I-code / 2020 NEC package is reported as effective 2023-01-01;
    energy provisions are effective 2023-07-01; local grace-period or permit-date
    transition rules were not resolved.
  confidence: 0.7
  source_ids:
  - src:usa-la:lsuccc-current-codes-2023-10-19
  - src:usa-la:rs-40-1730-28-5
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-la:lsuccc |
| Authority name | Louisiana State Uniform Construction Code Council |
| Authority type | statewide construction-code council |
| Legal basis | R.S. 40:1730.21 through 40:1730.40.1; especially R.S. 40:1730.22, 40:1730.26, and 40:1730.28 |
| Role | Reviews, adopts, modifies, and promulgates the State Uniform Construction Code; determines justified amendments; establishes certification and continuing-education requirements for code enforcement personnel |
| Enforcement model | Statewide code with local municipal/parish enforcement and certified third-party provider option |
| Source IDs | src:usa-la:rs-40-1730-21, src:usa-la:rs-40-1730-22, src:usa-la:rs-40-1730-26, src:usa-la:rs-40-1730-28 |
| Verification status | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-la:lsuccc | Louisiana State Uniform Construction Code Council | Adopt and amend statewide IBC provisions | R.S. 40:1730.22; R.S. 40:1730.26; R.S. 40:1730.28 | src:usa-la:rs-40-1730-22, src:usa-la:rs-40-1730-26, src:usa-la:rs-40-1730-28, src:usa-la:lac-17-i-ucc-2024-03-20 | partially_verified |
| Residential | ahj:usa-la:lsuccc | Louisiana State Uniform Construction Code Council | Adopt and amend statewide IRC provisions, with statutory limits on scope and optional local Appendix J | R.S. 40:1730.28 | src:usa-la:rs-40-1730-28, src:usa-la:lac-17-i-ucc-2024-03-20 | partially_verified |
| Existing Building / Rehabilitation | ahj:usa-la:lsuccc | Louisiana State Uniform Construction Code Council | Adopt and amend IEBC provisions; OSFM transition for some existing-building fire/life-safety enforcement requires follow-up | R.S. 40:1730.28; R.S. 40:1574-related source not fully parsed | src:usa-la:rs-40-1730-28, src:usa-la:lac-17-i-ucc-2024-03-20 | partially_verified |
| Mechanical | ahj:usa-la:lsuccc | Louisiana State Uniform Construction Code Council | Adopt and amend IMC provisions | R.S. 40:1730.28 | src:usa-la:rs-40-1730-28, src:usa-la:lac-17-i-ucc-2024-03-20 | partially_verified |
| Plumbing | ahj:usa-la:lsuccc | Louisiana State Uniform Construction Code Council | Adopt and amend IPC and related plumbing provisions; Department of Health sanitary-code carveouts require separate issue tracking | R.S. 40:1730.28; R.S. 40:1730.28.1 not fully parsed in this file | src:usa-la:rs-40-1730-28, src:usa-la:lac-17-i-ucc-2024-03-20 | partially_verified |
| Fuel Gas | ahj:usa-la:lsuccc | Louisiana State Uniform Construction Code Council | Adopt and amend IFGC provisions | R.S. 40:1730.28 | src:usa-la:rs-40-1730-28, src:usa-la:lac-17-i-ucc-2024-03-20 | partially_verified |
| Electrical | ahj:usa-la:lsuccc | Louisiana State Uniform Construction Code Council | Adopt and amend NEC provisions for construction | R.S. 40:1730.28 | src:usa-la:rs-40-1730-28, src:usa-la:lac-17-i-ucc-2024-03-20 | partially_verified |
| Energy | ahj:usa-la:lsuccc | Louisiana State Uniform Construction Code Council, with statutory energy-amendment process | Adopt IECC and IRC Part IV Energy Efficiency provisions | R.S. 40:1730.28.5 | src:usa-la:rs-40-1730-28-5, src:usa-la:lac-17-i-ucc-2024-03-20 | partially_verified |
| Fire - construction references | ahj:usa-la:osfm | Louisiana Office of State Fire Marshal | Enforce fire protection, egress, life-safety, and accessibility provisions in OSFM jurisdiction | R.S. 40:1574 through 40:1593; R.S. 40:1578.7; OSFM codes page | src:usa-la:osfm-codes-rules-laws, src:usa-la:rs-40-1578-7 | partially_verified |
| Fire - operational / prevention code | ahj:usa-la:osfm | Louisiana Office of State Fire Marshal | Enforce State Uniform Fire Prevention Code and NFPA 101 Life Safety Code | R.S. 40:1578.7; OSFM codes page | src:usa-la:rs-40-1578-7, src:usa-la:osfm-codes-rules-laws | partially_verified |
| Accessibility | ahj:usa-la:osfm | Louisiana Office of State Fire Marshal | Administer accessibility standards; current and 2026-07-01 statutory versions differ | R.S. 40:1733; OSFM codes page | src:usa-la:rs-40-1733, src:usa-la:osfm-codes-rules-laws | partially_verified |
| Elevator / Conveyance | ahj:usa-la:unresolved-elevator | Unresolved | Separate elevator/conveyance authority was not confirmed from official sources reviewed | unresolved | none | unresolved |

### 2.3 Authority Hierarchy Notes

The construction-code hierarchy is hybrid but not purely local. Louisiana law creates a statewide uniform construction code and assigns the LSUCCC the adoption and amendment function. Local governments enforce the statewide code, but R.S. 40:1730.23 limits them to the construction codes in the Part IV-A framework. R.S. 40:1730.24 allows municipalities and parishes, contractors, homeowners, and specified others to use certified third-party providers. R.S. 40:1730.25 requires each parish and municipality to appoint or contract for a council-certified building official so unincorporated parish areas are under a certified official.

The OSFM authority should not be collapsed into the LSUCCC construction-code authority. OSFM materials list both LSUCC building-code references and distinct fire, life-safety, accessibility, and fire-prevention codes. The fire-prevention statute has 2026-07-01 amendments that could affect local fire-code variance and adoption notes.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-la:001 | ahj:usa-la:legislature | creates | ahj:usa-la:lsuccc / State Uniform Construction Code Council | src:usa-la:rs-40-1730-22 | verified |
| edge:usa-la:002 | ahj:usa-la:lsuccc | adopts_and_amends | Louisiana State Uniform Construction Code | src:usa-la:rs-40-1730-22, src:usa-la:rs-40-1730-26, src:usa-la:rs-40-1730-28 | verified |
| edge:usa-la:003 | ahj:usa-la:lsuccc | requires_local_enforcement_by | municipalities_and_parishes | src:usa-la:rs-40-1730-23 | verified |
| edge:usa-la:004 | municipalities_and_parishes | may_contract_with | other_governmental_entities_or_certified_third_party_providers | src:usa-la:rs-40-1730-24 | verified |
| edge:usa-la:005 | ahj:usa-la:osfm | enforces | fire_life_safety_accessibility_fire_prevention | src:usa-la:osfm-codes-rules-laws, src:usa-la:rs-40-1578-7, src:usa-la:rs-40-1733 | partially_verified |
| edge:usa-la:006 | municipalities_and_parishes | prohibited_from_amending | mandatory_construction_code_provisions_required_by_R.S._40:1730.28 | src:usa-la:rs-40-1730-40 | verified |
| edge:usa-la:007 | local_fire_governing_authority | may_seek_approval_for_more_stringent_requirements_after_2026-07-01 | ahj:usa-la:osfm / State Uniform Fire Prevention Code variances | src:usa-la:rs-40-1578-7 | pending |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | Louisiana State Uniform Construction Code | International Building Code | 2021 | current_with_state_amendments | null | 2023-01-01 | 2023-01-01 | 2023-01-01 | No permit-date grace period resolved; current amendment set includes later amendments through 2024-03-20. IBC excludes Chapter 1, Chapter 11, and Chapter 27 in current rule text. | src:usa-la:lsuccc-current-codes-2023-10-19, src:usa-la:lac-17-i-ucc-2024-03-20, src:usa-la:osfm-codes-rules-laws |
| Residential | Louisiana State Uniform Construction Code | International Residential Code | 2021 | current_with_state_amendments | null | 2023-01-01 | 2023-01-01 | 2023-01-01 | Mandatory for new construction, reconstruction, additions to homes previously built to IRC, and extensive alterations; Appendix J may be adopted/enforced only at option of parish, municipality, or regional planning commission. | src:usa-la:rs-40-1730-28, src:usa-la:lac-17-i-ucc-2024-03-20 |
| Existing Building / Rehabilitation | Louisiana State Uniform Construction Code | International Existing Building Code | 2021 | current_with_state_amendments | null | 2023-01-01 | 2023-01-01 | 2023-01-01 | No general permit-date transition resolved; separate OSFM existing-building transition for structures permitted on or after 2026-01-01 requires follow-up. | src:usa-la:lac-17-i-ucc-2024-03-20 |
| Mechanical | Louisiana State Uniform Construction Code | International Mechanical Code | 2021 | current_with_state_amendments | null | 2023-01-01 | 2023-01-01 | 2023-01-01 | No permit-date grace period resolved. | src:usa-la:lac-17-i-ucc-2024-03-20 |
| Plumbing | Louisiana State Uniform Construction Code | International Plumbing Code | 2021 | current_with_state_amendments | null | 2023-01-01 | 2023-01-01 | 2023-01-01 | Appendices may be adopted as needed if specifically referenced; Department of Health sanitary-code carveouts not fully parsed. | src:usa-la:lac-17-i-ucc-2024-03-20 |
| Fuel Gas | Louisiana State Uniform Construction Code | International Fuel Gas Code | 2021 | current_with_state_amendments | null | 2023-01-01 | 2023-01-01 | 2023-01-01 | No permit-date grace period resolved. | src:usa-la:lac-17-i-ucc-2024-03-20 |
| Electrical | Louisiana State Uniform Construction Code | National Electric Code | 2020 | current_with_state_amendments | null | 2023-01-01 | 2023-01-01 | 2023-01-01 | No permit-date grace period resolved. | src:usa-la:lac-17-i-ucc-2024-03-20 |
| Energy | Louisiana State Uniform Construction Code | International Energy Conservation Code; IRC Part IV Energy Efficiency | 2021 | current_with_state_amendments | null | 2023-07-01 | 2023-07-01 | 2023-07-01 | R.S. 40:1730.28.5 states adopted energy-code provisions become effective 2023-07-01; later emergency/additional amendments require register-level reconciliation. | src:usa-la:rs-40-1730-28-5, src:usa-la:lsuccc-current-codes-2023-10-19, src:usa-la:lac-17-i-ucc-2024-03-20 |
| Fire - construction references | OSFM fire/life-safety provisions within LSUCC context | Fire protection and egress provisions of adopted LSUCC codes | mixed | current | null | 2023-01-01 | 2023-01-01 | 2023-01-01 | OSFM page lists fire-protection and egress provisions of LSUCC adopted codes; exact permit-date trigger for OSFM plan review unresolved. | src:usa-la:osfm-codes-rules-laws |
| Fire - operational / prevention code | State Uniform Fire Prevention Code and Life Safety Code | NFPA 1; NFPA 101 | NFPA 1 1997; NFPA 101 2015 | current_with_2026_transition | null | 2017-07-01 | 2017-07-01 | 2017-07-01 | OSFM lists NFPA 101 2015 as current for new and existing buildings effective 2017-07-01; R.S. 40:1578.7 has 2026-07-01 amendments allowing NFPA 1 or IFC framework. | src:usa-la:osfm-codes-rules-laws, src:usa-la:rs-40-1578-7 |
| Accessibility | OSFM accessibility standards | ADA Standards / ADA-ABA Accessibility Guidelines; future IBC Chapter 11 reference | 2010 ADA-ABA / ADA Standards in effect 2011-03-15; future IBC Chapter 11 | current_with_2026_transition | null | 2011-10-01 | 2011-10-01 | 2011-10-01 | R.S. 40:1733 current text adopts ADA Standards as of 2011-10-01; 2026-07-01 version references IBC Chapter 11 and ICC A117.1. | src:usa-la:rs-40-1733, src:usa-la:osfm-codes-rules-laws |
| Elevator / Conveyance | unresolved | unresolved | unresolved | unresolved | null | null | null | null | Distinct statewide elevator/conveyance code authority was not resolved. | none |

### 3.2 Adoption Records

| Adoption Record ID | Code Family | Adopted Code / Edition | Adoption Mechanism | Effective Date | Operative Date | Mandatory Date | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| adopt:usa-la:ibc-2021 | Building | 2021 IBC with Louisiana amendments | LAC Title 17, Part I; LSUCCC current-code summary points to Louisiana Register October 2022 pages 2577-2607 | 2023-01-01 | 2023-01-01 | 2023-01-01 | src:usa-la:lsuccc-current-codes-2023-10-19, src:usa-la:lac-17-i-ucc-2024-03-20 | Current consolidated rule text is March 2024 and includes later amendment history. |
| adopt:usa-la:iebc-2021 | Existing Building / Rehabilitation | 2021 IEBC with Louisiana amendments | LAC Title 17, Part I | 2023-01-01 | 2023-01-01 | 2023-01-01 | src:usa-la:lsuccc-current-codes-2023-10-19, src:usa-la:lac-17-i-ucc-2024-03-20 | Chapter 1 excluded. |
| adopt:usa-la:irc-2021 | Residential | 2021 IRC with Louisiana amendments | LAC Title 17, Part I | 2023-01-01 | 2023-01-01 | 2023-01-01 | src:usa-la:rs-40-1730-28, src:usa-la:lac-17-i-ucc-2024-03-20 | Parts I-Administrative and VIII-Electrical excluded; Appendix J local option noted. |
| adopt:usa-la:imc-2021 | Mechanical | 2021 IMC with Louisiana amendments | LAC Title 17, Part I | 2023-01-01 | 2023-01-01 | 2023-01-01 | src:usa-la:lac-17-i-ucc-2024-03-20 | Current consolidated rule text includes 2024 amendments. |
| adopt:usa-la:ipc-2021 | Plumbing | 2021 IPC with Louisiana amendments | LAC Title 17, Part I | 2023-01-01 | 2023-01-01 | 2023-01-01 | src:usa-la:lac-17-i-ucc-2024-03-20 | Plumbing-specific statutory transition commission provisions not fully parsed. |
| adopt:usa-la:ifgc-2021 | Fuel Gas | 2021 IFGC with Louisiana amendments | LAC Title 17, Part I | 2023-01-01 | 2023-01-01 | 2023-01-01 | src:usa-la:lac-17-i-ucc-2024-03-20 | Current consolidated rule text includes 2024 amendments. |
| adopt:usa-la:nec-2020 | Electrical | 2020 NEC with Louisiana amendments | LAC Title 17, Part I | 2023-01-01 | 2023-01-01 | 2023-01-01 | src:usa-la:lac-17-i-ucc-2024-03-20 | Rule heading uses "National Electric Code"; base model is NEC. |
| adopt:usa-la:iecc-2021 | Energy | 2021 IECC and IRC Part IV Energy Efficiency | R.S. 40:1730.28.5; LAC Title 17, Part I | 2023-07-01 | 2023-07-01 | 2023-07-01 | src:usa-la:rs-40-1730-28-5, src:usa-la:lac-17-i-ucc-2024-03-20 | Current-code summary also flags 2023-06-30 and 2023-07-30 emergency/additional amendment references for reconciliation. |
| adopt:usa-la:nfpa-101-2015 | Fire - operational / prevention code | NFPA 101 Life Safety Code, 2015 edition | OSFM code page; LAC 55:V:103 reference | 2017-07-01 | 2017-07-01 | 2017-07-01 | src:usa-la:osfm-codes-rules-laws | OSFM page says current edition for new and existing buildings. |
| adopt:usa-la:state-fire-prevention | Fire - operational / prevention code | State Uniform Fire Prevention Code | R.S. 40:1578.7 | null | null | null | src:usa-la:rs-40-1578-7 | Current statutory text adopts NFPA 1 1997 until 2026-07-01; future statutory text broadens to NFPA 1 or IFC. |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

The principal effective-date facts captured in this pass are: the 2021 I-code / 2020 NEC package is described by LSUCCC as effective 2023-01-01; 2021 energy-code provisions are effective 2023-07-01; freeboard amendments are identified by LSUCCC as effective 2023-08-01; and LSUCCC publishes a consolidated 2024-03-20 amendment document. This pass did not resolve a general grace-period, permit-application-date, permit-issuance-date, or concurrency rule for the 2021 code transition. Where a project-specific determination matters, local AHJ and OSFM plan-review rules should be checked.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-la:001 | 2021 IBC, IEBC, IRC, IMC, IPC, IFGC; 2020 NEC | effective_date | 2023-01-01 | LSUCCC current-code summary states 2021 codes and amendments effective 2023-01-01 | unresolved | src:usa-la:lsuccc-current-codes-2023-10-19 | partially_verified |
| date-rule:usa-la:002 | 2021 IECC; IRC Part IV Energy Efficiency | statutory_effective_date | 2023-07-01 | R.S. 40:1730.28.5 states adopted energy-code provisions become effective 2023-07-01 | unresolved | src:usa-la:rs-40-1730-28-5 | verified |
| date-rule:usa-la:003 | 2021 IRC / IBC freeboard amendments | amendment_effective_date | 2023-08-01 | LSUCCC current-code summary identifies freeboard amendments effective 2023-08-01 | unresolved | src:usa-la:lsuccc-current-codes-2023-10-19 | partially_verified |
| date-rule:usa-la:004 | Multiple construction-code amendments | amendment_effective_date | 2023-10-19 | LSUCCC current-code summary identifies new amendments effective 2023-10-19 | unresolved | src:usa-la:lsuccc-current-codes-2023-10-19 | partially_verified |
| date-rule:usa-la:005 | Consolidated LSUCCC amendment document | publication_or_consolidation_date | 2024-03-20 | LSUCCC Codes & Standards page labels the document "Uniform Construction Codes and Amendments effective 03/20/2024" | unresolved | src:usa-la:lsuccc-codes-standards, src:usa-la:lac-17-i-ucc-2024-03-20 | partially_verified |
| date-rule:usa-la:006 | Fire-prevention and accessibility statutes | statutory_change_date | 2026-07-01 | Legislative notes in R.S. 40:1578.7 and R.S. 40:1733 identify amendments effective 2026-07-01 | unresolved | src:usa-la:rs-40-1578-7, src:usa-la:rs-40-1733 | pending |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building / Existing Building / Electrical / I-codes | 2023 NEC and 2024 I-Codes proposed language | 2026-06-25 | null | null | null | null | watch | src:usa-la:lsuccc-home-announcements | LSUCCC homepage announcement states proposed language for the 2023 NEC and 2024 I-Codes was drafted for review/vote at the 2026-07-07 council meeting. |
| Fire - operational / prevention code | State Uniform Fire Prevention Code statutory update to NFPA 1 or IFC framework | null | null | 2026-07-01 | 2026-07-01 | 2026-07-01 | watch | src:usa-la:rs-40-1578-7 | Legislative text shows current and future versions effective 2026-07-01. |
| Accessibility | Future statutory shift to IBC Chapter 11 and ICC A117.1 | null | null | 2026-07-01 | 2026-07-01 | 2026-07-01 | watch | src:usa-la:rs-40-1733 | R.S. 40:1733 future version references IBC Chapter 11 as adopted by LSUCCC. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-la:001 | Residential | Existing homes and alterations | IRC scope rule | IRC enforcement is mandatory for new construction, reconstruction, additions to homes previously built to IRC, and extensive alterations. Appendix J may be adopted/enforced only at local option. | src:usa-la:rs-40-1730-28, src:usa-la:lac-17-i-ucc-2024-03-20 | verified |
| applicability-rule:usa-la:002 | Residential | One- and two-family dwelling sprinklers | IRC R313 amendments | Louisiana amendments state the council and municipalities/parishes shall not require fire-protection sprinklers in one- or two-family dwellings; voluntary installations remain allowed. | src:usa-la:lac-17-i-ucc-2024-03-20 | verified |
| applicability-rule:usa-la:003 | Fire - operational / prevention code | Local fire-prevention requirements after 2026-07-01 | State fire marshal review and approval | A local governing authority may provide more stringent State Uniform Fire Prevention Code requirements when based on local climatic, geologic, topographic, or public safety factors after prior OSFM review and approval. | src:usa-la:rs-40-1578-7 | pending |
| applicability-rule:usa-la:004 | Building / flood | Freeboard | Freeboard amendments effective 2023-08-01 | LSUCCC current-code summary identifies 2021 IRC/IBC freeboard amendments effective 2023-08-01; detailed provisions should be parsed from Register pages before production use. | src:usa-la:lsuccc-current-codes-2023-10-19, src:usa-la:la-admin-code-june-2023-excerpt | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** Louisiana publishes amendments through LAC Title 17, Part I for the State Uniform Construction Code and through Louisiana Register rulemaking. LSUCCC has authority to determine justified amendments and adopt rules under the Administrative Procedure Act. R.S. 40:1730.26 provides review/adoption/modification procedures and legislative oversight hooks for certain code updates.

**Where amendments are published:** Louisiana Administrative Code / Louisiana Register, LSUCCC Codes & Standards page, and ICC custom online editions for Louisiana amendments.

**Amendment parsing status:** partial_summary_only

### 5.2 State Amendment Sources

| Amendment Source ID | Source Type | Coverage | Effective / Publication Date | Source IDs | Caveat |
| --- | --- | --- | --- | --- | --- |
| amend-source:usa-la:lac-17-2024-03-20 | consolidated_pdf | Current LSUCCC consolidated amendment document for Title 17, Part I, Uniform Construction Code | 2024-03-20 | src:usa-la:lac-17-i-ucc-2024-03-20 | LSUCCC site labels the PDF effective 2024-03-20; full Register-level provenance should be reconciled for all line-item amendments. |
| amend-source:usa-la:codes-standards | agency_webpage | LSUCCC page listing current, prior, and custom online code resources | 2025 site footer; page accessed 2026-06-25 | src:usa-la:lsuccc-codes-standards | Webpage is a source registry and pointer page, not itself codified rule text. |
| amend-source:usa-la:current-codes-summary | agency_pdf_summary | Summary of current code components as of 2023-10-19, including Register page references | 2023-10-19 | src:usa-la:lsuccc-current-codes-2023-10-19 | Summary says references added together make up current adopted codes and amendments; it also says a correction document was being prepared. |
| amend-source:usa-la:icc-custom-online | custom_online_code | ICC-hosted Louisiana custom editions with red/blue amendments | current page link | src:usa-la:lsuccc-codes-standards | Supplemental access path; not treated as primary legal authority in this pass. |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Section / Topic | Change Summary | Source IDs | Confidence |
| --- | --- | --- | --- | --- | --- |
| amend:usa-la:ibc-exclusions | Building | IBC scope exclusions | 2021 IBC adopted without Chapter 1 Administration, Chapter 11 Accessibility, and Chapter 27 Electrical. | src:usa-la:lac-17-i-ucc-2024-03-20, src:usa-la:rs-40-1730-28 | 0.85 |
| amend:usa-la:irc-scope | Residential | IRC applicability and appendices | 2021 IRC adopted without Parts I-Administrative and VIII-Electrical; mandatory only for specified residential scopes; Appendix J is local option. | src:usa-la:rs-40-1730-28, src:usa-la:lac-17-i-ucc-2024-03-20 | 0.88 |
| amend:usa-la:irc-sprinklers | Residential | IRC R313 sprinklers | Louisiana amendments bar adoption or enforcement of a requirement for fire-protection sprinkler systems in one- or two-family dwellings, while allowing voluntary systems. | src:usa-la:lac-17-i-ucc-2024-03-20 | 0.84 |
| amend:usa-la:energy-climate-zone | Energy | IECC C301.2 warm humid counties | LAC text states all parishes in Louisiana are Climate Zone 2A warm humid climates. | src:usa-la:lac-17-i-ucc-2024-03-20 | 0.80 |
| amend:usa-la:freeboard | Building / Residential / flood | IBC / IRC freeboard | LSUCCC current-code summary identifies 2021 IRC/IBC freeboard amendments effective 2023-08-01; detailed provisions require register reconciliation. | src:usa-la:lsuccc-current-codes-2023-10-19, src:usa-la:la-admin-code-june-2023-excerpt | 0.62 |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-la"
  model: "mandatory_local_enforcement_of_statewide_uniform_code"
  enforcing_entities:
    - "municipalities"
    - "parishes"
    - "regional planning commissions where applicable"
    - "certified third-party providers where authorized or contracted"
    - "Office of State Fire Marshal for fire/life-safety/accessibility scopes"
  required_officials:
    - "council-certified building official or contracted equivalent for each parish and municipality"
    - "building code enforcement officers or certified third-party providers"
  state_reserved_activities:
    - "LSUCCC adoption and amendment of State Uniform Construction Code"
    - "OSFM fire prevention, life-safety, accessibility, and related plan-review/enforcement functions"
  source_ids:
    - "src:usa-la:rs-40-1730-23"
    - "src:usa-la:rs-40-1730-24"
    - "src:usa-la:rs-40-1730-25"
    - "src:usa-la:osfm-codes-rules-laws"
  verification_status: "partially_verified"
  confidence: 0.86
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-la"
  model: "local_administration_allowed_mandatory_code_amendments_preempted"
  applies_to_code_families:
    - "building"
    - "residential"
    - "existing building"
    - "mechanical"
    - "plumbing"
    - "fuel gas"
    - "electrical"
    - "energy"
  approval_required: false
  approving_authority_id: null
  filing_required: null
  registry_exists: null
  registry_source_ids: []
  legal_basis_source_ids:
    - "src:usa-la:rs-40-1730-40"
  verification_status: "partially_verified"
  confidence: 0.78

fire_local_variance_rule:
  rule_id: "local-fire-variance-rule:usa-la"
  model: "more_stringent_local_fire_prevention_requirements_after_osfm_review"
  applies_to_code_families:
    - "fire - operational / prevention code"
  approval_required: true
  approving_authority_id: "ahj:usa-la:osfm"
  filing_required: unresolved
  registry_exists: unresolved
  legal_basis_source_ids:
    - "src:usa-la:rs-40-1578-7"
  verification_status: "pending_2026-07-01"
  confidence: 0.70
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Louisiana separates local administration/enforcement from local code amendment authority. Local governments are central to enforcement, but that does not give them authority to replace or weaken mandatory statewide code provisions. The best-supported building-code rule is that local governments may adopt ordinances for local administration and enforcement procedures, while R.S. 40:1730.40 bars avoiding enforcement or amending mandatory R.S. 40:1730.28 construction-code provisions. Fire prevention is a separate track: the future 2026-07-01 version of R.S. 40:1578.7 includes a more-stringent-local-requirement pathway after OSFM review and approval.

### 6.4 Known Local Amendment Registries

No statewide local building-code amendment registry was verified. LSUCCC registrant searches exist for code enforcement officers, third-party providers, and public entities, but a local amendment registry was not confirmed in this pass.

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendments were parsed. Local administrative ordinances and local floodplain/freeboard requirements may exist, but they should be treated as local administration or special-hazard ordinances unless a legal basis for amendment of statewide mandatory provisions is established.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: partial_model_only

Jurisdiction stack:

```yaml
Address
  -> State of Louisiana
  -> Parish
  -> Municipality or unincorporated parish
  -> Floodplain manager / local floodplain ordinance, if flood hazard area
  -> Local building official or contracted certified third-party provider
  -> Office of State Fire Marshal for OSFM plan-review / fire / life-safety / accessibility scopes
  -> Trade-specific reviewers as locally assigned
  -> Applicable statewide LSUCC adoption records
  -> Applicable local administrative ordinances and confirmed special-hazard requirements
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | U.S. Census Bureau TIGER/Line or equivalent authoritative state boundary source | none | statewide | annual | pending |
| Parish | U.S. Census Bureau TIGER/Line county-equivalent boundaries | none | statewide | annual | pending |
| Municipality | U.S. Census Bureau Places or Louisiana geospatial portal | none | statewide | annual / variable | pending |
| Flood Hazard | FEMA NFHL and local floodplain ordinances | none | statewide where mapped | variable | pending |
| Fire District | OSFM / local fire district source | none | statewide where available | variable | pending |
| Special District | local / state special district datasets | none | variable | variable | pending |

### 7.3 AHJ Contact Data

No AHJ contact data was populated. The statewide source registry includes LSUCCC and OSFM source pages, but local AHJ records should be added separately during jurisdiction-specific implementation.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Issuer / Publisher | Source Type | URL / Citation | Date / Version | Used For | Accessed |
| --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-la:rs-40-1730-21 | R.S. 40:1730.21, Public policy for state uniform construction code | Louisiana State Legislature | statute | https://www.legis.la.gov/legis/Law.aspx?d=97795 | current as displayed 2026-06-25 | State uniform construction code purpose and certification purpose | 2026-06-25 |
| src:usa-la:rs-40-1730-22 | R.S. 40:1730.22, Louisiana State Uniform Construction Code Council | Louisiana State Legislature | statute | https://www.legis.la.gov/Legis/Law.aspx?d=97796 | current as displayed 2026-06-25 | LSUCCC creation, membership, function, amendment authority | 2026-06-25 |
| src:usa-la:rs-40-1730-23 | R.S. 40:1730.23, Enforcement of building codes by municipalities and parishes | Louisiana State Legislature | statute | https://www.legis.la.gov/Legis/Law.aspx?d=97797 | current text with 2026-07-01 notes | Local enforcement mandate and third-party provider limitations | 2026-06-25 |
| src:usa-la:rs-40-1730-24 | R.S. 40:1730.24, Agreements with other governmental entities; private agreements | Louisiana State Legislature | statute | https://www.legis.la.gov/legis/Law.aspx?d=97798 | current as displayed 2026-06-25 | Third-party provider and intergovernmental enforcement agreements | 2026-06-25 |
| src:usa-la:rs-40-1730-25 | R.S. 40:1730.25, Appointment of building official or contractual arrangement | Louisiana State Legislature | statute | https://www.legis.la.gov/legis/Law.aspx?d=97799 | current as displayed 2026-06-25 | Required local building official / contracted equivalent | 2026-06-25 |
| src:usa-la:rs-40-1730-26 | R.S. 40:1730.26, Adoption and promulgation procedures | Louisiana State Legislature | statute | https://www.legis.la.gov/legis/Law.aspx?d=97800 | current as displayed 2026-06-25 | LSUCCC adoption, review, update, and amendment procedure | 2026-06-25 |
| src:usa-la:rs-40-1730-28 | R.S. 40:1730.28, Mandatory adoption of certain nationally recognized codes | Louisiana State Legislature | statute | https://www.legis.la.gov/Legis/Law.aspx?d=97802 | current text with 2026-07-01 notes | Statutory code families and IRC applicability | 2026-06-25 |
| src:usa-la:rs-40-1730-28-5 | R.S. 40:1730.28.5, Energy code provisions | Louisiana State Legislature | statute | https://www.legis.la.gov/legis/Law.aspx?d=1296031 | current as displayed 2026-06-25 | IECC and IRC energy provisions; 2023-07-01 effective date | 2026-06-25 |
| src:usa-la:rs-40-1730-40 | R.S. 40:1730.40, Municipalities and parishes; home rule charter | Louisiana State Legislature | statute | https://www.legis.la.gov/legis/Law.aspx?d=587987 | current as displayed 2026-06-25 | Local administration authority and limitation on local amendment of mandatory code provisions | 2026-06-25 |
| src:usa-la:rs-40-1578-7 | R.S. 40:1578.7, State Uniform Fire Prevention Code | Louisiana State Legislature | statute | https://www.legis.la.gov/legis/Law.aspx?d=97686 | current and 2026-07-01 future text | Fire prevention code, local fire-code path, OSFM rule authority | 2026-06-25 |
| src:usa-la:rs-40-1733 | R.S. 40:1733, ADA Standards / standards to prevent architectural barriers | Louisiana State Legislature | statute | https://www.legis.la.gov/legis/Law.aspx?d=97809 | current and 2026-07-01 future text | Accessibility standard and pending transition | 2026-06-25 |
| src:usa-la:lsuccc-codes-standards | LSUCCC Codes & Standards page | Louisiana State Uniform Construction Code Council | agency webpage | https://lsuccc.la/codes-standards/ | page footer 2025; accessed 2026-06-25 | Source registry of current/prior code documents and ICC Louisiana custom editions | 2026-06-25 |
| src:usa-la:lac-17-i-ucc-2024-03-20 | Uniform Construction Codes and Amendments effective 03/20/2024, LAC Title 17 Part I | Louisiana State Uniform Construction Code Council | administrative-code PDF / agency-published consolidated rule text | https://lsuccc.la/wp-content/uploads/2025/10/ucc_amendments_03-20-24.pdf | Louisiana Administrative Code March 2024; effective label 2024-03-20 | Current code editions and state amendments | 2026-06-25 |
| src:usa-la:lsuccc-current-codes-2023-10-19 | Current codes effective 10-19-23 | Louisiana State Uniform Construction Code Council | agency PDF summary | https://lsuccc.la/wp-content/uploads/2025/10/links_to_current_codes_effective_10-19-23.pdf | 2023-10-19 | Effective-date summary and Register page map | 2026-06-25 |
| src:usa-la:osfm-codes-rules-laws | Codes, Rules & Laws Enforced by the Louisiana State Fire Marshal | Louisiana Office of State Fire Marshal | agency webpage | https://www.lasfm.org/plan-review/plan-review-codes-rules-laws/ | current page accessed 2026-06-25 | OSFM fire, life-safety, accessibility, and adopted LSUCC code references | 2026-06-25 |
| src:usa-la:la-admin-code-june-2023-excerpt | LAC Title 17 Part I, June 2023 excerpt including energy/freeboard updates | Louisiana DOTD-hosted document / Louisiana Administrative Code excerpt | administrative-code PDF excerpt | https://dotd.la.gov/media/ta0exrvw/updated-ibc-irc-freeboard.pdf | Louisiana Administrative Code June 2023 | Freeboard and energy update context | 2026-06-25 |
| src:usa-la:lsuccc-home-announcements | LSUCCC homepage announcements | Louisiana State Uniform Construction Code Council | agency webpage | https://lsuccc.la/ | current page accessed 2026-06-25 | Pending 2023 NEC / 2024 I-Code proposed-language watch item | 2026-06-25 |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-la:lac-17-i-ucc-2024-03-20 | consolidated_pdf | LSUCCC site labels the PDF as effective 2024-03-20 and the text contains historical Register citations, but it is a consolidated PDF rather than a direct Register issue page for each amendment. | acceptable_for_partial_verification; reconcile against Register before verified |
| src:usa-la:lsuccc-current-codes-2023-10-19 | agency_summary | The PDF is an agency summary of multiple Register publications and states that all references added together make the current adopted codes; it also mentions a correction document in progress. | use_for_date_map; verify each Register entry for production |
| src:usa-la:la-admin-code-june-2023-excerpt | agency-hosted_excerpt | Hosted on DOTD and appears to be LAC text; not the primary OSR Register source. | supplemental_only |
| src:usa-la:osfm-codes-rules-laws | agency_webpage | OSFM page is a practical code list and enforcement reference, not a codified statute or rule page for every listed standard. | use_with_statutes; verify LAC 55:V for production |
| src:usa-la:rs-40-1730-23 | future_notes | Page includes current and future versions around 2026-07-01. | recheck after 2026-07-01 |
| src:usa-la:rs-40-1578-7 | future_notes | Page includes current and future versions around 2026-07-01. | recheck after 2026-07-01 |
| src:usa-la:rs-40-1733 | future_notes | Page includes current and future versions around 2026-07-01. | recheck after 2026-07-01 |

### 8.3 Supplemental Sources

None relied on for core conclusions. ICC pages and private summaries were viewed only as orientation and were not used as authority-source IDs in the report body.

### 8.4 Source Extraction Metadata

| Extraction ID | Source ID | Extracted Fields | Method | Extracted By | Extracted Date | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| extract:usa-la:001 | src:usa-la:rs-40-1730-22 | LSUCCC creation, function, amendment authority | browser text extraction | GPT-5.5 Thinking | 2026-06-25 | Official legislature HTML |
| extract:usa-la:002 | src:usa-la:lac-17-i-ucc-2024-03-20 | Current code editions and amendment examples | browser PDF text extraction plus screenshot spot-check | GPT-5.5 Thinking | 2026-06-25 | PDF pages spot-checked for table of contents and current code sections |
| extract:usa-la:003 | src:usa-la:lsuccc-current-codes-2023-10-19 | effective-date map | browser PDF text extraction plus screenshot spot-check | GPT-5.5 Thinking | 2026-06-25 | PDF is one-page agency summary |
| extract:usa-la:004 | src:usa-la:osfm-codes-rules-laws | OSFM fire/life-safety/accessibility code list | browser text extraction | GPT-5.5 Thinking | 2026-06-25 | Agency webpage |
| extract:usa-la:005 | src:usa-la:rs-40-1730-40 | local amendment limitation | browser text extraction | GPT-5.5 Thinking | 2026-06-25 | Official legislature HTML |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 0.90 | none | Core authority and code-adoption fields now cite official sources; unresolved fields remain explicit. |
| authority:lsuccc | Authority name | Louisiana State Uniform Construction Code Council | verified | 0.95 | src:usa-la:rs-40-1730-22 | Official statute creates the council. |
| authority:lsuccc | Role | review/adopt State Uniform Construction Code and determine justified amendments | verified | 0.90 | src:usa-la:rs-40-1730-22, src:usa-la:rs-40-1730-26 | Role supported by statutes. |
| adoption:ibc | Base model code and edition | 2021 IBC | partially_verified | 0.82 | src:usa-la:lac-17-i-ucc-2024-03-20, src:usa-la:osfm-codes-rules-laws | Consolidated administrative-code PDF and OSFM list agree. |
| adoption:irc | Base model code and edition | 2021 IRC | verified | 0.85 | src:usa-la:lac-17-i-ucc-2024-03-20 | Scope caveats captured. |
| adoption:iecc | Effective date | 2023-07-01 | verified | 0.92 | src:usa-la:rs-40-1730-28-5 | Statute gives effective date. |
| local-enforcement | model | mandatory_local_enforcement_of_statewide_uniform_code | verified | 0.88 | src:usa-la:rs-40-1730-23, src:usa-la:rs-40-1730-24, src:usa-la:rs-40-1730-25 | Local enforcement and third-party provider model supported. |
| local-amendment | model | local_administration_allowed_mandatory_code_amendments_preempted | partially_verified | 0.78 | src:usa-la:rs-40-1730-40 | Local amendment registry unresolved. |
| fire-code | operational code | NFPA 101 2015 and State Uniform Fire Prevention Code | partially_verified | 0.72 | src:usa-la:osfm-codes-rules-laws, src:usa-la:rs-40-1578-7 | 2026 statutory transition flagged. |
| elevator | authority | unresolved | unresolved | 0.20 | none | Needs separate source pass. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Every src:usa-la:* cited in body is present in section 8. |
| All authority IDs resolve | pass | Active authority IDs appear in section 2. |
| Core current code families have adoption records | pass | IBC, IEBC, IRC, IMC, IPC, IFGC, NEC, IECC, fire/life-safety/accessibility are recorded; elevator/conveyance remains unresolved. |
| Building and operational fire code are separated | pass | LSUCCC construction-code authority is separated from OSFM fire/life-safety/fire-prevention authority. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Fields are separate; unresolved adoption dates remain null. |
| Effective dates are valid ISO dates | pass | Dates entered use YYYY-MM-DD. |
| No impossible date sequences | pass | No date sequences conflict based on current evidence. |
| Transition rules have explicit trigger conditions | fail | General permit-date / issuance-date grace-period logic remains unresolved. |
| Permit-date logic is captured where applicable | fail | No general permit-date rule was verified. |
| Local enforcement model classified | pass | Mandatory local enforcement model supported by statutes. |
| Local amendment rule classified | pass | Building-code local administration vs amendment posture is classified from R.S. 40:1730.40. |
| AHJ confirmation metadata present | fail | No local AHJ contacts were populated. |
| Official-source caveats captured | pass | Caveats for consolidated PDFs, summaries, and future statutory notes are recorded. |
| Pending 2026 statutory changes captured | pass | Fire, accessibility, enforcement, and proposed code adoption watch items are listed. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-la:001 | high | 2026 statutory changes | Multiple statute pages include future 2026-07-01 versions affecting enforcement terminology, fire-prevention code, and accessibility. | Reopen official statutes after 2026-07-01 and update authority/adoption rows. | null | 2026-07-15 | open |
| issue:usa-la:002 | high | 2024 / 2026 code adoption activity | LSUCCC homepage references proposed 2023 NEC and 2024 I-Code language for a 2026-07-07 council meeting. | Monitor meeting packet/minutes and any Notice of Intent / Louisiana Register publication. | null | 2026-08-15 | open |
| issue:usa-la:003 | medium | Register-level reconciliation | LSUCCC summary and consolidated PDF should be reconciled against Louisiana Register October 2022, June 2023, emergency-rule, October 2023, and March 2024 entries. | Extract OSR Register publications and map each effective date to official entries. | null | null | open |
| issue:usa-la:004 | medium | elevator / conveyance authority | Distinct elevator/conveyance code and inspection authority was not confirmed. | Search Louisiana statutes/rules and OSFM or workforce/elevator program sources. | null | null | open |
| issue:usa-la:005 | medium | local amendment registry | No statewide local building-code amendment registry was verified; local fire-prevention more-stringent requirements may require OSFM approval after 2026-07-01. | Confirm whether LSUCCC or OSFM keeps local administrative or variance records. | null | null | open |
| issue:usa-la:006 | low | AHJ data | No parish/municipality AHJ contacts or public-entity registrations were added. | Use LSUCCC public-entity search and local websites to populate jurisdiction records. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-la:lsuccc-codes-standards | src:usa-la:lsuccc-codes-standards | html_diff | monthly | New effective code/amendment PDF or custom online edition link | 2026-06-25 | active |
| watch:usa-la:lac-17-ucc-pdf | src:usa-la:lac-17-i-ucc-2024-03-20 | pdf_diff | monthly | New consolidated Title 17, Part I PDF or changes to code editions/amendments | 2026-06-25 | active |
| watch:usa-la:lsuccc-home | src:usa-la:lsuccc-home-announcements | html_diff | weekly_until_2026-08-15 | Action on 2023 NEC / 2024 I-Code proposed language | 2026-06-25 | active |
| watch:usa-la:legislature-1730 | src:usa-la:rs-40-1730-22 | statute_check | after_2026-07-01_then_quarterly | Statutory changes to LSUCCC composition/function/effective notes | 2026-06-25 | active |
| watch:usa-la:fire-prevention | src:usa-la:rs-40-1578-7 | statute_check | after_2026-07-01_then_quarterly | State Uniform Fire Prevention Code future version becomes current | 2026-06-25 | active |
| watch:usa-la:osfm-codes | src:usa-la:osfm-codes-rules-laws | html_diff | monthly | Updated fire, life-safety, accessibility code editions or plan-review notes | 2026-06-25 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-25 | Populated Louisiana report from baseline draft | report:usa-la, ahj:usa-la:lsuccc, local-enforcement:usa-la, adoption records | src:usa-la:rs-40-1730-21, src:usa-la:rs-40-1730-22, src:usa-la:rs-40-1730-23, src:usa-la:rs-40-1730-24, src:usa-la:rs-40-1730-25, src:usa-la:rs-40-1730-26, src:usa-la:rs-40-1730-28, src:usa-la:rs-40-1730-28-5, src:usa-la:rs-40-1730-40, src:usa-la:rs-40-1578-7, src:usa-la:rs-40-1733, src:usa-la:lsuccc-codes-standards, src:usa-la:lac-17-i-ucc-2024-03-20, src:usa-la:lsuccc-current-codes-2023-10-19, src:usa-la:osfm-codes-rules-laws | GPT-5.5 Thinking | Upgraded status to partially_verified for core fields; unresolved elevator, AHJ, permit-date transition, and source reconciliation items left explicit. |
