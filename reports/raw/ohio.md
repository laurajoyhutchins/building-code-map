---
state:
  state_id: "US-OH"
  name: "Ohio"
  abbreviation: "OH"
report:
  report_id: "state-report:usa-oh"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-26"
  last_verified: null
  reviewed_by: null
risk:
  overall_confidence: 0.67 # 0.00 - 1.00
  risk_flags:
    - "transition_rule_needs_primary_bbs_memo_confirmation"
    - "residential_base_code_needs_full_rule_parse"
    - "elevator_and_boiler_authority_not_fully_parsed"
    - "local_amendment_registry_not_identified"
    - "ahj_contact_data_not_populated"
  open_questions_count: 7

---

# State Building Code Authority Report: Ohio

## 1. Executive Summary

- **Authority model:** Ohio uses a statewide minimum-code model led by the Ohio Board of Building Standards (BBS). BBS is the primary authority for Ohio construction-code rules under ORC 3781.10 and administers certified local building departments and certified personnel under OAC Chapter 4101:7.

- **Statewide code status:** Core statewide construction-code adoptions are partially verified. The 2024 Ohio Building Code, Mechanical Code, Plumbing Code, and Existing Building Code package became effective on 2024-03-01 and is based on 2021 ICC model codes. Current OAC 4101:1-1-01 remains organized around the 2021 IBC incorporation, with later amendments effective 2025-10-15. The 2025 Ohio Fire Code became effective on 2025-11-20 and incorporates 2021 IFC chapters with 2024 IFC Chapter 12.

- **Local enforcement model:** Municipalities, townships, and counties may seek BBS certification to exercise building-code enforcement authority. Certified local departments enforce within their certified scope; where authority is not certified or is suspended/revoked, state or other certified enforcement paths may apply. Plumbing enforcement has special limits and may involve certified municipal departments, county building departments with health-board agreements, local boards of health, or the Department of Commerce Division of Industrial Compliance.

- **Local amendment posture:** Ohio provides uniform minimum statewide standards. Municipal corporations may adopt additional building regulations only when they do not conflict with ORC Chapters 3781/3791 or BBS rules. Local residential regulations by counties, townships, or municipalities require a BBS notice/request path and BBS conflict review.

- **Known transition periods or pending changes:** A 2024 transition rule is identified through secondary references to BBS Memo 414: applications/submittals before 2024-03-01 remain under the prior code in effect at first application; applications/submittals on or after 2024-03-01 use the 2024 codes. The primary memo text was not captured and remains an open issue. The 2025 Ohio Fire Code is current as of this report.

- **Production readiness:** partially_ready. The core statewide authority and adoption model are source-backed. Production use should wait for full RCO parsing, official BBS transition-memo capture, local amendment registry confirmation, and AHJ/contact enrichment.

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: The Ohio Board of Building Standards formulates and adopts rules governing
    construction, repair, alteration, and maintenance of buildings and related equipment.
  confidence: 0.95
  source_ids:
  - src:usa-oh:orc-3781-10
- topic: Primary building code edition
  finding: The 2024 Ohio Building Code package adopted 2021 IBC chapters 2-35 and
    Appendix H by reference with Ohio amendments, effective 2024-03-01; current administrative
    rule text shows OAC 4101:1-1-01 effective 2025-10-15.
  confidence: 0.9
  source_ids:
  - src:usa-oh:commerce-2024-obc-rules
  - src:usa-oh:oac-4101-1-1-01-2025
- topic: Mechanical and plumbing editions
  finding: The 2024 Ohio Mechanical Code and 2024 Ohio Plumbing Code were adopted
    as part of Amendments Group 100, effective 2024-03-01, incorporating the 2021
    IMC and 2021 IPC by reference with Ohio amendments.
  confidence: 0.9
  source_ids:
  - src:usa-oh:commerce-2024-omc-rules
  - src:usa-oh:commerce-2024-opc-rules
- topic: Fire code authority
  finding: The state fire marshal adopts and enforces the state fire code; the current
    Ohio Fire Code rule is effective 2025-11-20.
  confidence: 0.9
  source_ids:
  - src:usa-oh:orc-3737-22
  - src:usa-oh:orc-3737-82
  - src:usa-oh:oac-1301-7-7-01-2025
- topic: Local enforcement
  finding: BBS-certified municipalities, townships, and counties may exercise building-code
    enforcement authority; certification and certified personnel are prerequisites.
  confidence: 0.85
  source_ids:
  - src:usa-oh:oac-4101-7-1-01
  - src:usa-oh:oac-4101-7-2-01
- topic: Local amendments
  finding: Local regulations may supplement state standards only within conflict limits;
    local residential regulations have a BBS conflict-review path.
  confidence: 0.8
  source_ids:
  - src:usa-oh:orc-3781-01
  - src:usa-oh:orc-3781-11
  - src:usa-oh:oac-4101-1-1-01-2025
- topic: Effective / operative date rule
  finding: 2024 code package effective date is verified; precise application-date
    transition rule still needs the primary BBS memo.
  confidence: 0.55
  source_ids:
  - src:usa-oh:commerce-2024-obc-rules
  - src:usa-oh:commerce-2024-omc-rules
  - src:usa-oh:commerce-2024-opc-rules
  - src:usa-oh:bbs-memo-414-secondary
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-oh:bbs |
| Authority name | Ohio Board of Building Standards |
| Authority type | statewide building-code board |
| Legal basis | ORC 3781.10; related certification rules in OAC Chapter 4101:7 |
| Role | Formulates and adopts statewide construction-code rules; certifies local building departments, personnel, and related program participants. |
| Enforcement model | Statewide minimum code with certified local enforcement and state backstop/specialized enforcement paths. |
| Source IDs | src:usa-oh:orc-3781-10; src:usa-oh:oac-4101-7-1-01; src:usa-oh:oac-4101-7-2-01 |
| Verification status | verified_core |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-oh:bbs | Ohio Board of Building Standards | Adopts statewide building-code rules and certifies enforcement departments/personnel. | ORC 3781.10; OAC 4101:1-1-01; OAC 4101:7 | src:usa-oh:orc-3781-10; src:usa-oh:oac-4101-1-1-01-2025; src:usa-oh:oac-4101-7-2-01 | verified_core |
| Residential | ahj:usa-oh:bbs | Ohio Board of Building Standards | Adopts residential-code rules and residential-code amendments; certifies residential building departments/personnel. | ORC 3781.10; OAC 4101:8; OAC 4101:7 | src:usa-oh:commerce-rco-2019-rules; src:usa-oh:commerce-rco-ag101; src:usa-oh:oac-4101-7-2-01 | partially_verified |
| Existing Building / Rehabilitation | ahj:usa-oh:bbs | Ohio Board of Building Standards | Administers Ohio Existing Building Code through OAC 4101:1-34-01 / OBC Chapter 34 and IEBC incorporation. | OAC 4101:1-1-01; OBC Chapter 34 | src:usa-oh:oac-4101-1-1-01-2025; src:usa-oh:commerce-2024-overview | verified_core |
| Mechanical | ahj:usa-oh:bbs | Ohio Board of Building Standards | Adopts Ohio Mechanical Code rules. | OAC 4101:2; ORC 3781.10 | src:usa-oh:commerce-2024-omc-rules; src:usa-oh:orc-3781-10 | verified_core |
| Plumbing | ahj:usa-oh:bbs | Ohio Board of Building Standards | Adopts Ohio Plumbing Code rules; enforcement may involve certified municipal/county departments, health boards, or Division of Industrial Compliance depending on certification and agreements. | OAC 4101:3; ORC 3781.03; OAC 4101:7-2-01 | src:usa-oh:commerce-2024-opc-rules; src:usa-oh:orc-3781-03; src:usa-oh:oac-4101-7-2-01 | verified_core |
| Fuel Gas | ahj:usa-oh:bbs | Ohio Board of Building Standards | Fuel-gas provisions are administered as part of the 2024 code package and OBC administrative scope. | OAC 4101:1-1-01; 2024 Building Codes overview | src:usa-oh:oac-4101-1-1-01-2025; src:usa-oh:commerce-2024-overview | partially_verified |
| Electrical | ahj:usa-oh:bbs | Ohio Board of Building Standards | Electrical requirements are administered through OBC Chapter 27 and RCO Chapter 34; 2024 standards overview identifies NFPA 70-2023. | OAC 4101:1-27-01; OAC 4101:8-34-01 | src:usa-oh:oac-4101-1-27-01; src:usa-oh:commerce-2024-overview; src:usa-oh:commerce-rco-ag101 | partially_verified |
| Energy | ahj:usa-oh:bbs | Ohio Board of Building Standards | Energy provisions are administered through OBC Chapter 13 / energy standards within the 2024 code package. | OAC 4101:1-13; OBC Chapter 13 | src:usa-oh:commerce-2024-overview; src:usa-oh:oac-4101-1-1-01-2025 | partially_verified |
| Fire - construction references | ahj:usa-oh:bbs | Ohio Board of Building Standards | OBC construction rules include fire-protection-system and safeguards-during-construction provisions and coordinate with Ohio Fire Code references. | OAC 4101:1-1-01; OBC Chapters 9, 33, 35 | src:usa-oh:oac-4101-1-1-01-2025; src:usa-oh:commerce-2024-overview | partially_verified |
| Fire - operational / prevention code | ahj:usa-oh:fire-marshal | Ohio State Fire Marshal | Adopts and enforces Ohio Fire Code / fire-prevention and life-safety rules. | ORC 3737.22; ORC 3737.82; OAC 1301:7-7-01 | src:usa-oh:orc-3737-22; src:usa-oh:orc-3737-82; src:usa-oh:oac-1301-7-7-01-2025 | verified_core |
| Accessibility | ahj:usa-oh:bbs | Ohio Board of Building Standards | Accessibility requirements are administered within OBC Chapter 11 and referenced standards. | OAC 4101:1-11; OBC Chapter 11 | src:usa-oh:commerce-2024-overview; src:usa-oh:oac-4101-1-1-01-2025 | partially_verified |
| Elevator / Conveyance | ahj:usa-oh:bbs-unresolved-elevator | Ohio Board of Building Standards / Ohio Department of Commerce, elevator program | OBC administrative rule references the Ohio Elevator Code, but the dedicated elevator authority and current edition were not fully parsed. | OAC 4101:5; OAC 4101:1-1-01 | src:usa-oh:oac-4101-1-1-01-2025; src:usa-oh:commerce-2024-overview | unresolved_partial |
| Boiler / Pressure Vessel | ahj:usa-oh:bbs-unresolved-boiler | Ohio Department of Commerce, boiler program | OBC administrative rule references Ohio Boiler and Pressure Vessel Rules; detailed authority not fully parsed. | OAC 4101:4; OAC 4101:1-1-01 | src:usa-oh:oac-4101-1-1-01-2025; src:usa-oh:commerce-2024-overview | unresolved_partial |

### 2.3 Authority Hierarchy Notes

Ohio separates statewide rule adoption from local enforcement certification. BBS is the primary statewide construction-code rulemaking body. Local political subdivisions do not automatically become construction-code AHJs; they must be certified by BBS as building departments or sub-building departments and must use appropriately certified personnel. Plumbing enforcement is more specialized than general building enforcement and depends on certification, health-board relationships, and Division of Industrial Compliance fallback roles. The state fire marshal is separate from BBS for the operational/prevention fire code.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-oh:001 | ahj:usa-oh:bbs | adopts_rules_for | statewide construction codes | src:usa-oh:orc-3781-10 | verified_core |
| edge:usa-oh:002 | ahj:usa-oh:bbs | certifies | municipal, township, and county building departments | src:usa-oh:oac-4101-7-2-01 | verified_core |
| edge:usa-oh:003 | certified local building departments | enforce_within_certified_scope | construction-document review, approvals, and inspections | src:usa-oh:oac-4101-7-2-01 | verified_core |
| edge:usa-oh:004 | ahj:usa-oh:bbs | supersedes_conflicting_rules | conflicting local/state construction standards except preserved statutory exceptions | src:usa-oh:orc-3781-11; src:usa-oh:oac-4101-1-1-01-2025 | partially_verified |
| edge:usa-oh:005 | municipalities/counties/townships | may_adopt_additional_residential_regulations_subject_to | BBS conflict review and temporary-variance process | src:usa-oh:orc-3781-01 | verified_core |
| edge:usa-oh:006 | ahj:usa-oh:fire-marshal | adopts_and_enforces | state fire code / operational fire prevention | src:usa-oh:orc-3737-22; src:usa-oh:orc-3737-82 | verified_core |
| edge:usa-oh:007 | local fire-code officials | enforce | Ohio Fire Code within delegated/local jurisdiction | src:usa-oh:oac-1301-7-7-01-2025 | partially_verified |
| edge:usa-oh:008 | political subdivision with local fire-code authority | may_adopt | IFC appendices as local fire code | src:usa-oh:oac-1301-7-7-01-2025 | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | 2024 Ohio Building Code, with later OAC 4101:1-1-01 amendments | 2021 International Building Code, Chapters 2-35 and Appendix H | 2021 | current_core_verified | 2023-08-11 | 2024-03-01 | 2024-03-01 | 2024-03-01 | 2024-03-01 effective date verified; application-date rule requires primary BBS Memo 414 capture. | src:usa-oh:commerce-2024-obc-rules; src:usa-oh:oac-4101-1-1-01-2025; src:usa-oh:bbs-memo-414-secondary |
| Residential | 2019 Residential Code of Ohio, with 2024 targeted amendments | 2018 International Residential Code; RCO Chapter 34 updated to NFPA 70-2023 by 2024 amendment | 2018 / 2023 electrical | current_partial | 2018-12-14; 2024-01-26 for AG101 amendments | 2019-07-01; 2024-03-01 for AG101 amendments | 2019-07-01; 2024-03-01 for AG101 amendments | unknown | Base RCO transition not parsed; AG101 effective date verified. | src:usa-oh:commerce-rco-2019-rules; src:usa-oh:commerce-rco-ag101 |
| Existing Building / Rehabilitation | 2024 Ohio Existing Building Code / OBC Chapter 34 | 2021 International Existing Building Code | 2021 | current_core_verified | 2023-08-11 | 2024-03-01 | 2024-03-01 | 2024-03-01 | Same 2024 code-package transition issue as OBC; primary memo unresolved. | src:usa-oh:commerce-2024-overview; src:usa-oh:oac-4101-1-1-01-2025; src:usa-oh:bbs-memo-414-secondary |
| Mechanical | 2024 Ohio Mechanical Code | 2021 International Mechanical Code | 2021 | current_core_verified | 2023-08-11 | 2024-03-01 | 2024-03-01 | 2024-03-01 | 2024-03-01 effective date verified; application-date rule requires primary BBS Memo 414 capture. | src:usa-oh:commerce-2024-omc-rules; src:usa-oh:bbs-memo-414-secondary |
| Plumbing | 2024 Ohio Plumbing Code | 2021 International Plumbing Code | 2021 | current_core_verified | 2023-08-11 | 2024-03-01 | 2024-03-01 | 2024-03-01 | 2024-03-01 effective date verified; application-date rule requires primary BBS Memo 414 capture. | src:usa-oh:commerce-2024-opc-rules; src:usa-oh:bbs-memo-414-secondary |
| Fuel Gas | Fuel-gas provisions in 2024 Ohio code package | 2021 International Fuel Gas Code | 2021 | current_partial | 2023-08-11 | 2024-03-01 | 2024-03-01 | 2024-03-01 | Same as OBC where administered through building-code package. | src:usa-oh:commerce-2024-overview; src:usa-oh:oac-4101-1-1-01-2025 |
| Electrical | OBC Chapter 27 and RCO Chapter 34 electrical provisions | NFPA 70-2023 referenced in 2024 standards / RCO AG101; OBC Chapter 27 part of 2024 OBC | 2023 NFPA 70 / 2021 IBC Chapter 27 | current_partial | 2023-08-11; 2024-01-26 for RCO AG101 | 2024-03-01 | 2024-03-01 | 2024-03-01 for verified amended provisions | Need full OAC 4101:1-27-01 and RCO 4101:8-34-01 parse for all scopes. | src:usa-oh:oac-4101-1-27-01; src:usa-oh:commerce-2024-overview; src:usa-oh:commerce-rco-ag101 |
| Energy | OBC Chapter 13 Energy Efficiency | 2021 IECC and ASHRAE 90.1-2019 | 2021 / 2019 | current_partial | 2023-08-11 | 2024-03-01 | 2024-03-01 | 2024-03-01 | Same as OBC where administered through building-code package. | src:usa-oh:commerce-2024-overview; src:usa-oh:oac-4101-1-1-01-2025 |
| Fire - construction references | OBC fire-protection and construction-safeguard provisions | 2021 IBC with IFC references | 2021 | current_partial | 2023-08-11 | 2024-03-01 | 2024-03-01 | 2024-03-01 | OBC construction references coordinate with separate Ohio Fire Code operational requirements. | src:usa-oh:commerce-2024-overview; src:usa-oh:oac-4101-1-1-01-2025 |
| Fire - operational / prevention code | 2025 Ohio Fire Code | 2021 International Fire Code Chapters 1-11 and 20-80; 2024 IFC Chapter 12 | 2021 / 2024 | current_core_verified | 2025-11-10 | 2025-11-20 | 2025-11-20 | 2025-11-20 | Prior conflicting fire codes/regulations superseded as of effective date; local appendices may be adopted only by political subdivisions with local fire-code authority. | src:usa-oh:oac-1301-7-7-01-2025 |
| Accessibility | OBC Chapter 11 Accessibility | 2021 IBC Chapter 11; ICC A117.1 standards by scope | 2021 / 2017 / 2009 | current_partial | 2023-08-11 | 2024-03-01 | 2024-03-01 | 2024-03-01 | Same as OBC; detailed standards need chapter-level parse. | src:usa-oh:commerce-2024-overview; src:usa-oh:oac-4101-1-1-01-2025 |
| Elevator / Conveyance | Ohio Elevator Code | OAC 4101:5 rules; model-code edition unresolved | unknown | unresolved_partial | null | null | null | null | OBC references Ohio Elevator Code, but dedicated adoption dates/editions were not parsed. | src:usa-oh:oac-4101-1-1-01-2025; src:usa-oh:commerce-2024-overview |
| Boiler / Pressure Vessel | Ohio Boiler and Pressure Vessel Rules | OAC 4101:4 rules; edition unresolved | unknown | unresolved_partial | null | null | null | null | OBC references boiler rules, but dedicated adoption dates/editions were not parsed. | src:usa-oh:oac-4101-1-1-01-2025; src:usa-oh:commerce-2024-overview |

### 3.2 Adoption Records

#### adoption:usa-oh:obc-2024

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-oh:bbs |
| Code family | Building |
| State code name | 2024 Ohio Building Code |
| Base model code | 2021 International Building Code |
| Scope parsed | OAC 4101:1-1-01 incorporates 2021 IBC chapters 2-35 and Appendix H by reference, subject to Ohio amendments. |
| Adoption date | 2023-08-11 |
| Effective date | 2024-03-01 |
| Current rule update | OAC 4101:1-1-01 effective 2025-10-15 |
| Source IDs | src:usa-oh:commerce-2024-obc-rules; src:usa-oh:oac-4101-1-1-01-2025 |
| Verification status | verified_core |

#### adoption:usa-oh:omc-2024

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-oh:bbs |
| Code family | Mechanical |
| State code name | 2024 Ohio Mechanical Code |
| Base model code | 2021 International Mechanical Code |
| Adoption date | 2023-08-11 |
| Effective date | 2024-03-01 |
| Source IDs | src:usa-oh:commerce-2024-omc-rules |
| Verification status | verified_core |

#### adoption:usa-oh:opc-2024

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-oh:bbs |
| Code family | Plumbing |
| State code name | 2024 Ohio Plumbing Code |
| Base model code | 2021 International Plumbing Code |
| Adoption date | 2023-08-11 |
| Effective date | 2024-03-01 |
| Source IDs | src:usa-oh:commerce-2024-opc-rules |
| Verification status | verified_core |

#### adoption:usa-oh:oebc-2024

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-oh:bbs |
| Code family | Existing Building / Rehabilitation |
| State code name | 2024 Ohio Existing Building Code / OAC 4101:1-34-01 |
| Base model code | 2021 International Existing Building Code |
| Adoption date | 2023-08-11 |
| Effective date | 2024-03-01 |
| Source IDs | src:usa-oh:commerce-2024-overview; src:usa-oh:oac-4101-1-1-01-2025 |
| Verification status | verified_core |

#### adoption:usa-oh:rco-2019-base

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-oh:bbs |
| Code family | Residential |
| State code name | 2019 Residential Code of Ohio |
| Base model code | 2018 International Residential Code |
| Adoption date | 2018-12-14 |
| Effective date | 2019-07-01 |
| Source IDs | src:usa-oh:commerce-rco-2019-rules |
| Verification status | partially_verified |
| Notes | Base RCO PDF/source was identified, but the full rule text was not parsed in this pass. |

#### adoption:usa-oh:rco-ag101-2024

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-oh:bbs |
| Code family | Residential / electrical / referenced standards |
| State code name | Amendments Group 101 to the Residential Code of Ohio |
| Base model / standards | NFPA 70-2023 for RCO Chapter 34; updated referenced standards for Chapter 44 |
| Adoption date | 2024-01-26 |
| Effective date | 2024-03-01 |
| Source IDs | src:usa-oh:commerce-rco-ag101 |
| Verification status | verified_core_for_amendment_notice |

#### adoption:usa-oh:ofc-2025

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-oh:fire-marshal |
| Code family | Fire - operational / prevention code |
| State code name | 2025 Ohio Fire Code |
| Base model code | 2021 International Fire Code Chapters 1-11 and 20-80; 2024 International Fire Code Chapter 12 |
| Rule final date | 2025-11-10 |
| Effective date | 2025-11-20 |
| Source IDs | src:usa-oh:oac-1301-7-7-01-2025; src:usa-oh:orc-3737-22; src:usa-oh:orc-3737-82 |
| Verification status | verified_core |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

The official Ohio Commerce 2024 OBC, OMC, and OPC rule packages show a common 2024-03-01 effective date after BBS adoption on 2023-08-11. The 2024 Building Codes overview also identifies the 2024 building-code package as effective 2024-03-01. Secondary sources quoting or summarizing BBS Memo 414 identify an application/submittal trigger: applications before 2024-03-01 remain under the prior code at first application, and applications on or after 2024-03-01 use the 2024 codes. Because the primary BBS Memo 414 text was not captured, the application-date rule is recorded as partially verified rather than fully verified.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-oh:001 | 2024 OBC / OMC / OPC / OEBC package | effective_date | 2024-03-01 | Rule package effective date. | unknown | src:usa-oh:commerce-2024-obc-rules; src:usa-oh:commerce-2024-omc-rules; src:usa-oh:commerce-2024-opc-rules; src:usa-oh:commerce-2024-overview | verified_core |
| date-rule:usa-oh:002 | 2024 OBC / OMC / OPC / OEBC package | application_or_submittal_date_transition | 2024-03-01 | Secondary references to BBS Memo 414 indicate that first application/submittal date controls code edition. | yes, for applications/submittals before 2024-03-01, subject to primary memo confirmation | src:usa-oh:bbs-memo-414-secondary | partially_verified |
| date-rule:usa-oh:003 | RCO Amendments Group 101 | amendment_effective_date | 2024-03-01 | RCO Chapter 4, 34, and 44 amendments adopted 2024-01-26 become effective. | unknown | src:usa-oh:commerce-rco-ag101 | verified_core_for_amendment_notice |
| date-rule:usa-oh:004 | 2025 Ohio Fire Code | effective_date | 2025-11-20 | OAC 1301:7-7-01 effective date. | conflicting prior fire codes/regulations superseded as of effective date | src:usa-oh:oac-1301-7-7-01-2025 | verified_core |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Fire - operational / prevention code | 2025 Ohio Fire Code | 2025-11-20 effective rule identified | 2025-11-10 | 2025-11-20 | 2025-11-20 | 2025-11-20 | current | src:usa-oh:oac-1301-7-7-01-2025 | Current as of this report; monitor for later State Fire Marshal/OAC changes. |
| Building / mechanical / plumbing / existing building | Later post-2024 code amendments | null | null | 2025-10-15 for current OAC 4101:1-1-01 admin rule | null | null | active_monitoring | src:usa-oh:oac-4101-1-1-01-2025 | Current administrative rule text effective 2025-10-15 should be compared against 2024 package for amendment deltas. |
| Residential | Later RCO amendments beyond AG101 | null | null | null | null | null | unresolved | src:usa-oh:commerce-rco-ag101 | Official 2024 overview references RCO amendments to three rules; full RCO amendment history remains open. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-oh:001 | Building / Residential | Detached one-, two-, and three-family dwellings and accessory structures | Building type / scope | OBC administrative scope excludes work within RCO scope. | src:usa-oh:oac-4101-1-1-01-2025 | verified_core |
| applicability-rule:usa-oh:002 | Fire - operational / prevention | Existing and new conditions affecting fire/life safety | Fire-code scope | Ohio Fire Code applies statewide to fire-safety matters and operational/prevention conditions, while structural building requirements covered by OBC are exempted from OFC where applicable. | src:usa-oh:oac-1301-7-7-01-2025 | verified_core |
| applicability-rule:usa-oh:003 | Plumbing | Plumbing enforcement jurisdiction | Local certification and health-board agreement status | Counties and townships have limited plumbing certification paths; local boards of health and Division of Industrial Compliance may enforce where local departments are not certified or agreements apply. | src:usa-oh:oac-4101-7-2-01; src:usa-oh:orc-3781-03 | partially_verified |
| applicability-rule:usa-oh:004 | Fire code appendices | Political subdivisions with local fire-code authority | Local adoption of appendices | IFC appendices are not adopted statewide; political subdivisions with authority to enact local fire code may adopt appendices locally. | src:usa-oh:oac-1301-7-7-01-2025 | verified_core |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** Ohio incorporates model codes by reference and publishes Ohio amendments through OAC chapters and official Commerce/BBS rule packages. For the 2024 construction-code package, the Board adopted by reference the 2021 ICC model codes and maintained Ohio-specific amendments in OAC chapter/rule structure.

**Where amendments are published:** Ohio Administrative Code, Ohio Laws and Administrative Rules portal, Ohio Department of Commerce / Board of Building Standards PDF rule packages, and ICC-published Ohio code volumes where available.

**Amendment parsing status:** partial. Administrative scope, core adoption rules, and major adoption notices were reviewed. Chapter-level amendments were not fully parsed across all chapters.

### 5.2 State Amendment Sources

| Amendment Source ID | Code Family | Publication | Scope | Status | Notes |
| --- | --- | --- | --- | --- | --- |
| amendset:usa-oh:obc-2024 | Building / Existing Building / energy / electrical references / accessibility | 2024 Ohio Building Code Rules Effective 2024-03-01 and current OAC 4101:1-1-01 | OBC administrative provisions, IBC incorporation, Chapter 34 IEBC, referenced standards and Ohio amendments | partially_parsed | Current OAC effective 2025-10-15 should be diffed against 2024 rule package. |
| amendset:usa-oh:omc-2024 | Mechanical | 2024 Ohio Mechanical Code Rules Effective 2024-03-01 | OMC chapters 1-15 incorporating 2021 IMC | partially_parsed | Covers adoption notice and scope; chapter-level deltas unresolved. |
| amendset:usa-oh:opc-2024 | Plumbing | 2024 Ohio Plumbing Code Rules Effective 2024-03-01 | OPC chapters 1-15 incorporating 2021 IPC | partially_parsed | Covers adoption notice and scope; chapter-level deltas unresolved. |
| amendset:usa-oh:rco-ag101 | Residential / electrical / referenced standards | RCO Amendments Group 101 notice package | RCO chapters 4, 34, 44 | partially_parsed | Notice confirms NFPA 70-2023 for RCO Chapter 34; full RCO base remains partially parsed. |
| amendset:usa-oh:ofc-2025 | Fire | OAC 1301:7-7-01 effective 2025-11-20 | Ohio Fire Code, IFC incorporation, scope, appendices, supersession | partially_parsed | Rule 1301:7-7-01 parsed for high-value administration and scope. |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Topic | Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| amend:usa-oh:001 | Building | Model-code incorporation | OBC incorporates 2021 IBC Chapters 2-35 and Appendix H by reference with Ohio amendments. | src:usa-oh:commerce-2024-obc-rules; src:usa-oh:oac-4101-1-1-01-2025 | verified_core |
| amend:usa-oh:002 | Existing Building | IEBC / Chapter 34 | OBC Chapter 34 uses the 2021 IEBC as modified for existing-building work. | src:usa-oh:oac-4101-1-1-01-2025; src:usa-oh:commerce-2024-overview | verified_core |
| amend:usa-oh:003 | Energy | Energy standards | 2024 overview identifies OBC Chapter 13 energy efficiency as 2021 IECC and ASHRAE 90.1-2019. | src:usa-oh:commerce-2024-overview | partially_verified |
| amend:usa-oh:004 | Electrical | NFPA 70-2023 | 2024 overview and RCO AG101 identify NFPA 70-2023; RCO Chapter 34 was amended to adopt NFPA 70-2023 with Ohio modifications. | src:usa-oh:commerce-2024-overview; src:usa-oh:commerce-rco-ag101 | partially_verified |
| amend:usa-oh:005 | Fire | Fire-code appendices | IFC appendices are not adopted statewide; local political subdivisions with local fire-code authority may adopt appendices locally. | src:usa-oh:oac-1301-7-7-01-2025 | verified_core |
| amend:usa-oh:006 | Local conflict | State minimum uniformity | BBS construction-code rules supersede conflicting construction standards, subject to statutory exceptions and permitted non-conflicting local regulations. | src:usa-oh:orc-3781-11; src:usa-oh:oac-4101-1-1-01-2025 | verified_core |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-oh"
  model: "statewide_minimum_code_with_bbs_certified_local_enforcement"
  enforcing_entities:
    - "Ohio Board of Building Standards-certified municipal building departments"
    - "Ohio Board of Building Standards-certified township building departments"
    - "Ohio Board of Building Standards-certified county building departments"
    - "Certified sub-building departments acting within certified scope"
    - "Division of Industrial Compliance or approved certified department where local certification is absent, suspended, or revoked"
    - "State fire marshal and fire-code officials for Ohio Fire Code operational/prevention scope"
  required_officials:
    - "Certified building official / master plans examiner / building inspector as applicable"
    - "Certified residential building official / residential plans examiner / residential building inspector as applicable"
    - "Certified plumbing inspector for plumbing enforcement where certification and jurisdiction allow"
  state_reserved_activities:
    - "Certification of building departments and personnel"
    - "Rule adoption and statewide code amendments"
    - "Fire-code rule adoption by state fire marshal"
    - "Fallback or specialized enforcement when local certification is absent or revoked"
  source_ids:
    - "src:usa-oh:oac-4101-7-1-01"
    - "src:usa-oh:oac-4101-7-2-01"
    - "src:usa-oh:orc-3781-03"
    - "src:usa-oh:orc-3737-22"
  verification_status: "partially_verified"
  confidence: 0.82
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-oh"
  model: "statewide_uniform_minimum_with_limited_nonconflicting_local_additions"
  applies_to_code_families:
    - "building"
    - "residential"
    - "fire appendices and local fire code where local authority exists"
  approval_required: "conditional"
  approving_authority_id: "ahj:usa-oh:bbs for local residential conflict review; local fire-code appendices depend on political subdivision authority"
  filing_required: "unresolved"
  registry_exists: "unresolved"
  registry_source_ids: []
  legal_basis_source_ids:
    - "src:usa-oh:orc-3781-01"
    - "src:usa-oh:orc-3781-11"
    - "src:usa-oh:oac-4101-1-1-01-2025"
    - "src:usa-oh:oac-1301-7-7-01-2025"
  verification_status: "partially_verified"
  confidence: 0.76
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Ohio local enforcement authority and local amendment authority are distinct. A county, township, or municipality may enforce the construction codes only after BBS certification and only within the certified scope. That certification does not itself prove that the local government can amend the state code. Local amendment authority is constrained by ORC 3781.01 and ORC 3781.11 conflict rules, and local residential regulations have a specific BBS notice/review process. Fire-code appendices are a separate issue: Ohio Fire Code rule text states that IFC appendices are not adopted statewide, but political subdivisions with authority to enact a local fire code may adopt appendices locally.

### 6.4 Known Local Amendment Registries

| Registry ID | Registry Name | Code Family | Source IDs | Status | Notes |
| --- | --- | --- | --- | --- | --- |
| registry:usa-oh:local-building-amendments | statewide local building amendment registry | building / residential | none | unresolved | No statewide registry was confirmed. Need BBS search and local-code filing review. |
| registry:usa-oh:local-fire-appendices | local fire-code appendix adoption registry | fire | none | unresolved | Ohio Fire Code permits local appendix adoption by authorized political subdivisions, but a statewide registry was not identified. |

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendments were parsed for this report. Do not infer local code modifications from certification status alone.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: started_conceptual; data_not_loaded

Jurisdiction stack:

```text
Address
  -> State of Ohio
  -> County
  -> Municipality / township / unincorporated county
  -> BBS-certified building department or sub-building department, if any
  -> Plumbing enforcement entity based on certification and health-board agreement status
  -> Fire-code official / state fire marshal / local fire department with applicable authority
  -> Specialized state programs such as elevator, boiler, or industrialized units where applicable
  -> Applicable statewide code adoption records
  -> Applicable local amendment or local fire appendix records, if verified
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | not selected | none | statewide | unknown | pending |
| County | not selected | none | statewide | unknown | pending |
| Municipality | not selected | none | statewide | unknown | pending |
| Township | not selected | none | statewide | unknown | pending |
| Certified building department coverage | BBS certified department list | src:usa-oh:oac-4101-7-2-01 | statewide expected | unknown | source_basis_verified_list_not_extracted |
| Fire District / fire-code official coverage | not selected | none | statewide expected | unknown | pending |
| Special District | not selected | none | statewide expected | unknown | pending |

### 7.3 AHJ Contact Data

No AHJ contact data was populated. The rule basis for a BBS-maintained list of certified building departments was verified, but the actual list was not extracted.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Publisher / Custodian | Type | Date / Effective Date | Locator | Reliability | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-oh:orc-3781-10 | Ohio Revised Code § 3781.10 - Board of building standards; powers and duties | Ohio Laws and Administrative Rules / Ohio Revised Code | statute | current page reviewed 2026-06-26 | https://codes.ohio.gov/ohio-revised-code/section-3781.10 | official | Primary authority for BBS construction-code rulemaking. |
| src:usa-oh:orc-3781-03 | Ohio Revised Code § 3781.03 - Enforcement provisions | Ohio Laws and Administrative Rules / Ohio Revised Code | statute | current page reviewed 2026-06-26 | https://codes.ohio.gov/ohio-revised-code/section-3781.03 | official | Supports enforcement split among fire, construction, and plumbing authorities. |
| src:usa-oh:orc-3781-01 | Ohio Revised Code § 3781.01 - Municipal and residential local regulations | Ohio Laws and Administrative Rules / Ohio Revised Code | statute | current page reviewed 2026-06-26 | https://codes.ohio.gov/ohio-revised-code/section-3781.01 | official | Supports local non-conflict and residential conflict-review rules. |
| src:usa-oh:orc-3781-11 | Ohio Revised Code § 3781.11 - Rules of board; uniform standards; conflict/supersession | Ohio Laws and Administrative Rules / Ohio Revised Code | statute | current page reviewed 2026-06-26 | https://codes.ohio.gov/ohio-revised-code/section-3781.11 | official | Supports statewide uniform minimum code and supersession rules. |
| src:usa-oh:oac-4101-1-1-01-2025 | OAC 4101:1-1-01 - Administration, effective 2025-10-15 | Ohio Laws and Administrative Rules / Ohio Administrative Code | administrative rule / authenticated PDF | 2025-10-15 | https://codes.ohio.gov/assets/laws/administrative-code/rules/4101/1/4101%241-1-01_eff_10_15_25.pdf | official | Current OBC administrative rule; includes 2021 IBC incorporation, scope, local regulation text, IEBC, and referenced code coordination. |
| src:usa-oh:oac-4101-1-27-01 | OAC 4101:1-27-01 - Electrical | Ohio Laws and Administrative Rules / Ohio Administrative Code | administrative rule | 2024-03-01 | https://codes.ohio.gov/ohio-administrative-code/rule-4101:1-27-01 | official | Rule identified; full electrical scope not fully parsed. |
| src:usa-oh:oac-4101-7-1-01 | OAC 4101:7-1-01 - Certification general provisions | Ohio Laws and Administrative Rules / Ohio Administrative Code | administrative rule | current page reviewed 2026-06-26 | https://codes.ohio.gov/ohio-administrative-code/rule-4101:7-1-01 | official | Supports BBS certification program and prerequisite nature of certification. |
| src:usa-oh:oac-4101-7-2-01 | OAC 4101:7-2-01 - Building department certification | Ohio Laws and Administrative Rules / Ohio Administrative Code | administrative rule | current page reviewed 2026-06-26 | https://codes.ohio.gov/ohio-administrative-code/rule-4101:7-2-01 | official | Supports local enforcement certification model, plumbing distinctions, and certified department list. |
| src:usa-oh:commerce-2024-obc-rules | 2024 Ohio Building Code Rules Effective March 1, 2024 | Ohio Department of Commerce / Ohio Board of Building Standards | official PDF rule package | adopted 2023-08-11; effective 2024-03-01 | https://dam.assets.ohio.gov/image/upload/com.ohio.gov/documents/2024%20Ohio%20Building%20Code%20Rules%20Effective%20March%201%2C%202024.pdf | official | Official rule package for 2024 OBC; model code text incorporated by reference, not reproduced in full. |
| src:usa-oh:commerce-2024-omc-rules | 2024 Ohio Mechanical Code Rules Effective March 1, 2024 | Ohio Department of Commerce / Ohio Board of Building Standards | official PDF rule package | adopted 2023-08-11; effective 2024-03-01 | https://dam.assets.ohio.gov/image/upload/com.ohio.gov/documents/2024%20Ohio%20Mechanical%20Code%20Rules%20Effective%20March%201%2C%202024.pdf | official | Official rule package for 2024 OMC; incorporates 2021 IMC by reference with Ohio amendments. |
| src:usa-oh:commerce-2024-opc-rules | 2024 Ohio Plumbing Code Rules Effective March 1, 2024 | Ohio Department of Commerce / Ohio Board of Building Standards | official PDF rule package | adopted 2023-08-11; effective 2024-03-01 | https://dam.assets.ohio.gov/image/upload/com.ohio.gov/documents/2024%20Ohio%20Plumbing%20Code%20Rules%20Effective%20March%201%2C%202024.pdf | official | Official rule package for 2024 OPC; incorporates 2021 IPC by reference with Ohio amendments. |
| src:usa-oh:commerce-2024-overview | Overview 2024 Building Codes | Ohio Department of Commerce / Ohio Board of Building Standards | official presentation PDF | 2024-01-30 | https://dam.assets.ohio.gov/image/upload/com.ohio.gov/documents/2024%20OBC%20Executive%20Summary%201.pdf | official_supplemental | Helpful summary for 2021 ICC model-code package, IECC/ASHRAE/electrical/accessibility standards, and other rules; not a substitute for rule text. |
| src:usa-oh:commerce-rco-2019-rules | Residential Code of Ohio Rules Effective July 1, 2019 | Ohio Department of Commerce / Ohio Board of Building Standards | official PDF rule package | adopted 2018-12-14; effective 2019-07-01 | https://dam.assets.ohio.gov/image/upload/com.ohio.gov/documents/bbst_ResidentialCodeofOhioEffectiveJuly1%2C2019.pdf | official_partial | Source identified for RCO base; full text not parsed in this pass. |
| src:usa-oh:commerce-rco-ag101 | Amendments Group 101 Notice Package - Residential Code of Ohio | Ohio Department of Commerce / Ohio Board of Building Standards | official PDF notice package | adopted 2024-01-26; effective 2024-03-01 | https://dam.assets.ohio.gov/image/upload/com.ohio.gov/documents/AG101%20Notice%20package.pdf | official | Confirms RCO Chapter 4, 34, and 44 amendments and NFPA 70-2023 reference. |
| src:usa-oh:orc-3737-22 | Ohio Revised Code § 3737.22 - Fire marshal duties and code enforcement | Ohio Laws and Administrative Rules / Ohio Revised Code | statute | current page reviewed 2026-06-26 | https://codes.ohio.gov/ohio-revised-code/section-3737.22 | official | Primary source for fire marshal fire-code adoption/enforcement duties. |
| src:usa-oh:orc-3737-82 | Ohio Revised Code § 3737.82 - State fire code | Ohio Laws and Administrative Rules / Ohio Revised Code | statute | current page reviewed 2026-06-26 | https://codes.ohio.gov/ohio-revised-code/section-3737.82 | official | Requires the fire marshal to adopt a state fire code and permits incorporation by reference. |
| src:usa-oh:oac-1301-7-7-01-2025 | OAC 1301:7-7-01 - Ohio Fire Code, effective 2025-11-20 | Ohio Laws and Administrative Rules / Ohio Administrative Code | administrative rule / authenticated PDF | final 2025-11-10; effective 2025-11-20 | https://codes.ohio.gov/assets/laws/administrative-code/rules/1301/7/1301%247-7-01_eff_11_20_25.pdf | official | Current Ohio Fire Code administrative rule; incorporates 2021 IFC chapters and 2024 IFC Chapter 12. |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-oh:commerce-2024-obc-rules | incorporated_model_code | Official Ohio rule package incorporates copyrighted model-code chapters by reference and does not reproduce all model-code text. | use_for_adoption_and_ohio_amendments; do_not_treat_as_full_model_code_text |
| src:usa-oh:oac-4101-1-1-01-2025 | current_rule_after_2024_package | Current rule effective 2025-10-15 may include amendments after the 2024 package. | compare_current_rule_against_2024_package_before_final_versioning |
| src:usa-oh:commerce-2024-overview | summary_not_rule_text | Presentation is official but summarizes code changes; rule text controls. | supplemental_only_for_standards_summary |
| src:usa-oh:commerce-rco-2019-rules | partial_parse | Source identified but full base RCO text was not parsed. | complete_full_parse_before_verified_status |
| src:usa-oh:bbs-memo-414-secondary | secondary_transition_source | Transition rule was found through secondary references to BBS Memo 414, not the primary memo. | cannot_mark_transition_rule_verified_until_primary_memo_captured |
| src:usa-oh:oac-1301-7-7-01-2025 | authenticated_pdf | Fire-code rule PDF is authoritative for the administrative rule, but incorporated IFC text may require separate lawful access. | use_for_scope_and_incorporation; do_not_extract_full_ifc_text |

### 8.3 Supplemental Sources

| Source ID | Title | Publisher | Type | Date | Locator | Use | Caveat |
| --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-oh:bbs-memo-414-secondary | Secondary references to BBS Memo 414 transition guidance | ICC NTA / local Ohio building department notices | secondary web references | 2024 | secondary references reviewed 2026-06-26 | Supports unresolved transition-rule lead only. | Not an official primary source; replace with official BBS memo when found. |

### 8.4 Source Extraction Metadata

| Source ID | Extracted Fields | Extraction Method | Extracted On | Extractor | Notes |
| --- | --- | --- | --- | --- | --- |
| src:usa-oh:orc-3781-10 | BBS rulemaking authority | Ohio Laws HTML | 2026-06-26 | ChatGPT | Core authority parsed. |
| src:usa-oh:orc-3781-03 | enforcement split | Ohio Laws HTML | 2026-06-26 | ChatGPT | Fire/construction/plumbing split parsed at high level. |
| src:usa-oh:orc-3781-01 | local regulation rules | Ohio Laws HTML | 2026-06-26 | ChatGPT | Local residential BBS conflict process captured. |
| src:usa-oh:orc-3781-11 | uniformity and supersession | Ohio Laws HTML | 2026-06-26 | ChatGPT | Uniform minimum and conflict rules captured. |
| src:usa-oh:oac-4101-1-1-01-2025 | OBC administration, IBC incorporation, local regulation, IEBC, references | Ohio Laws HTML/PDF and screenshot spot-check | 2026-06-26 | ChatGPT | Current rule effective 2025-10-15. |
| src:usa-oh:oac-4101-7-1-01 | certification program scope | Ohio Laws HTML | 2026-06-26 | ChatGPT | Certification framework captured. |
| src:usa-oh:oac-4101-7-2-01 | local building department certification and plumbing distinctions | Ohio Laws HTML | 2026-06-26 | ChatGPT | High-value enforcement details captured. |
| src:usa-oh:commerce-2024-obc-rules | OBC adoption/effective dates and IBC incorporation | official PDF and screenshot spot-check | 2026-06-26 | ChatGPT | Adoption package parsed for high-level records. |
| src:usa-oh:commerce-2024-omc-rules | OMC adoption/effective dates and IMC incorporation | official PDF and screenshot spot-check | 2026-06-26 | ChatGPT | Adoption package parsed for high-level records. |
| src:usa-oh:commerce-2024-opc-rules | OPC adoption/effective dates and IPC incorporation | official PDF and screenshot spot-check | 2026-06-26 | ChatGPT | Adoption package parsed for high-level records. |
| src:usa-oh:commerce-2024-overview | model-code package and standards summary | official PDF and screenshot spot-check | 2026-06-26 | ChatGPT | Used as official supplemental summary. |
| src:usa-oh:commerce-rco-ag101 | RCO AG101 adoption/effective dates and NFPA 70-2023 | official PDF and screenshot spot-check | 2026-06-26 | ChatGPT | Targeted RCO amendment notice parsed. |
| src:usa-oh:orc-3737-22 | fire marshal authority | Ohio Laws HTML | 2026-06-26 | ChatGPT | Core fire authority parsed. |
| src:usa-oh:orc-3737-82 | state fire code adoption authority | Ohio Laws HTML | 2026-06-26 | ChatGPT | Core fire code authority parsed. |
| src:usa-oh:oac-1301-7-7-01-2025 | Ohio Fire Code incorporation, scope, appendices, supersession | Ohio Laws HTML/PDF and screenshot spot-check | 2026-06-26 | ChatGPT | Current fire-code adoption parsed. |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 1.00 | none | Status reflects source-backed core fields plus unresolved gaps. |
| report | risk.overall_confidence | 0.67 | verified | 1.00 | none | Confidence is intentionally below verified because several scope-specific records remain open. |
| ahj:usa-oh:bbs | authority.name | Ohio Board of Building Standards | verified_core | 0.95 | src:usa-oh:orc-3781-10 | Primary construction-code authority. |
| ahj:usa-oh:fire-marshal | authority.name | Ohio State Fire Marshal | verified_core | 0.90 | src:usa-oh:orc-3737-22; src:usa-oh:orc-3737-82 | Primary operational/prevention fire-code authority. |
| adoption:usa-oh:obc-2024 | base_model_code | 2021 International Building Code | verified_core | 0.90 | src:usa-oh:commerce-2024-obc-rules; src:usa-oh:oac-4101-1-1-01-2025 | Rule package and current OAC both support IBC incorporation. |
| adoption:usa-oh:omc-2024 | base_model_code | 2021 International Mechanical Code | verified_core | 0.90 | src:usa-oh:commerce-2024-omc-rules | Official OMC rule package. |
| adoption:usa-oh:opc-2024 | base_model_code | 2021 International Plumbing Code | verified_core | 0.90 | src:usa-oh:commerce-2024-opc-rules | Official OPC rule package. |
| adoption:usa-oh:ofc-2025 | effective_date | 2025-11-20 | verified_core | 0.90 | src:usa-oh:oac-1301-7-7-01-2025 | Current Ohio Fire Code rule. |
| local-enforcement:usa-oh | model | statewide_minimum_code_with_bbs_certified_local_enforcement | partially_verified | 0.82 | src:usa-oh:oac-4101-7-1-01; src:usa-oh:oac-4101-7-2-01 | Actual AHJ list not extracted. |
| local-amendment-rule:usa-oh | model | statewide_uniform_minimum_with_limited_nonconflicting_local_additions | partially_verified | 0.76 | src:usa-oh:orc-3781-01; src:usa-oh:orc-3781-11 | Registry and local examples unresolved. |
| date-rule:usa-oh:002 | application_or_submittal_date_transition | first application/submittal date controls 2024 code applicability | partially_verified | 0.55 | src:usa-oh:bbs-memo-414-secondary | Primary BBS Memo 414 still required. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Every `src:usa-oh:` identifier used in body is represented in section 8, except literal `none` values. |
| All authority IDs resolve | pass | Authority IDs used in section 2 and 6 are defined or explicitly marked unresolved_partial. |
| All current code families have adoption rows | pass | Matrix rows are present for core and unresolved families; unresolved records stay explicit. |
| Building and operational fire code are separated | pass | BBS construction-code scope and fire marshal/OFC operational-prevention scope are separate records. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Date columns are separate; unknown mandatory dates are left as `unknown` or `null`. |
| Effective dates are valid ISO dates | pass | Entered dates use YYYY-MM-DD format. |
| No impossible date sequences | pass | No adoption/effective sequence conflicts found in populated records. |
| Transition rules have explicit trigger conditions | partial | Effective dates are clear; application/submittal trigger needs primary BBS Memo 414. |
| Permit-date logic is captured where applicable | partial | Application/submittal rule captured as secondary-source lead only. |
| Local enforcement model classified | pass | Certified-local-enforcement model captured; AHJ list not extracted. |
| Local amendment rule classified | partial | Non-conflict and residential-review rules captured; registry unresolved. |
| AHJ confirmation metadata present | fail | No AHJ contacts or certified department coverage data extracted. |
| Official-source caveats captured | pass | Caveats recorded for PDFs, current-rule deltas, summary sources, and secondary transition source. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-oh:001 | high | 2024 transition rule | Application/submittal transition rule depends on secondary references to BBS Memo 414. | Locate and extract official BBS Memo 414 or equivalent primary BBS guidance. | null | null | open |
| issue:usa-oh:002 | high | Residential code base | 2019 RCO base was identified but not fully parsed; current RCO amendment state may include later amendments beyond AG101. | Parse full RCO OAC 4101:8 rules and amendment history, including 2024 April amendments if applicable. | null | null | open |
| issue:usa-oh:003 | medium | Electrical scope | Electrical provisions are partially captured through OBC Chapter 27 / NFPA 70-2023 and RCO Chapter 34, but exact statewide scope needs chapter parse. | Extract OAC 4101:1-27-01 and 4101:8-34-01 in full. | null | null | open |
| issue:usa-oh:004 | medium | Elevator / boiler authority | OBC references OAC 4101:5 elevator and OAC 4101:4 boiler rules, but current editions and authority records are unresolved. | Parse elevator and boiler chapters and Department of Commerce program pages. | null | null | open |
| issue:usa-oh:005 | medium | Local amendment registry | No statewide local amendment or local fire appendix registry was confirmed. | Search BBS, State Fire Marshal, and Secretary/local filing sources for registries or filing duties. | null | null | open |
| issue:usa-oh:006 | medium | Certified department list | Rule basis for BBS certified building department list is verified; actual list was not captured. | Extract current BBS certified department list and normalize AHJ coverage. | null | null | open |
| issue:usa-oh:007 | low | Boundary and contact data | No boundary files or AHJ contacts were loaded. | Select authoritative state/county/municipality/township/fire district boundary data and contact sources. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-oh:oac-4101-1-1-01 | src:usa-oh:oac-4101-1-1-01-2025 | html_or_pdf_diff | monthly | OBC administration rule amended or new effective date posted | 2026-06-26 | active |
| watch:usa-oh:oac-1301-7-7-01 | src:usa-oh:oac-1301-7-7-01-2025 | html_or_pdf_diff | monthly | Ohio Fire Code rule amended or new effective date posted | 2026-06-26 | active |
| watch:usa-oh:bbs-code-packages | src:usa-oh:commerce-2024-obc-rules | portal_diff | monthly | New BBS code package, amendment group, or memo posted | 2026-06-26 | active |
| watch:usa-oh:bbs-certified-departments | src:usa-oh:oac-4101-7-2-01 | data_pull | monthly | Certified department list changes | 2026-06-26 | pending |
| watch:usa-oh:local-amendments | none | targeted_research | quarterly | Evidence of statewide local amendment filing or registry appears | 2026-06-26 | pending |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-24 | Generated baseline Ohio draft report | report:usa-oh | none | Codex | Initial placeholder draft. |
| 2026-06-26 | Populated core Ohio authority, adoption, enforcement, amendment, source-registry, and QA sections | report:usa-oh; ahj:usa-oh:bbs; ahj:usa-oh:fire-marshal; adoption:usa-oh:obc-2024; adoption:usa-oh:omc-2024; adoption:usa-oh:opc-2024; adoption:usa-oh:ofc-2025 | src:usa-oh:orc-3781-10; src:usa-oh:oac-4101-1-1-01-2025; src:usa-oh:commerce-2024-obc-rules; src:usa-oh:commerce-2024-omc-rules; src:usa-oh:commerce-2024-opc-rules; src:usa-oh:oac-1301-7-7-01-2025 | ChatGPT | Status upgraded to partially_verified; unresolved items preserved explicitly. |
