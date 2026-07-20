---
state:
  state_id: "US-FL"
  name: "Florida"
  abbreviation: "FL"
report:
  report_id: "state-report:usa-fl"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-25"
  last_verified: "2026-06-25"
  reviewed_by: null
risk:
  overall_confidence: 0.82 # 0.00 - 1.00
  risk_flags:
    - "local_amendments_allowed_only_as_more_stringent"
    - "fire_code_is_separate_from_building_code"
    - "base_model_code_verified_at_cycle_level_not_volume_text_level"
    - "ahj_contacts_not_populated"
  open_questions_count: 3

---

# State Building Code Authority Report: Florida

## 1. Executive Summary

- **Authority model:** Statewide building code adopted and maintained by the Florida Building Commission, with enforcement primarily vested in local governments and legally constituted enforcement districts, subject to statutory state-agency exceptions.

- **Statewide code status:** Florida Building Code, 8th Edition (2023), effective December 31, 2023. The 8th Edition development process selected the 2021 I-Codes and 2020 NEC as model codes for review; this report records those as cycle-level base references and leaves volume-level parsing open.

- **Fire code status:** Florida Fire Prevention Code, 8th Edition (2023), effective December 31, 2023, adopted by the State Fire Marshal / Department of Financial Services and enforced by local fire officials in counties, municipalities, and special fire districts.

- **Local enforcement model:** Hybrid local/state model. Local governments and enforcement districts generally administer permits, plan review, and inspections, while specified facility types and subjects are reserved to state agencies or special state review paths.

- **Local amendment posture:** Local building-code and fire-code amendments may be more stringent than the statewide minimums if statutory process requirements are met. No state pre-approval requirement was verified for local building-code technical amendments in this pass; transmission, publication, availability, and challenge/review procedures are separately documented.

- **Known transition periods or pending changes:** Florida Building Commission and State Fire Marshal materials show 9th Edition (2026) development or proposed rulemaking activity, but this report has not verified an adopted 2026 effective date.

- **Production readiness:** usable_with_caution

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: The Florida Building Commission adopts the Florida Building Code under
    Fla. Stat. 553.73.
  confidence: 0.99
  source_ids:
  - src:usa-fl:statute:553.73
- topic: Primary building code edition
  finding: Florida Building Code, 8th Edition (2023), is the current statewide building-code
    cycle and is effective December 31, 2023.
  confidence: 0.99
  source_ids:
  - src:usa-fl:agency:bcis-homepage
  - src:usa-fl:agency:fbc-annual-report-2021
- topic: Cycle-level model code basis
  finding: The 8th Edition (2023) update process selected the 2021 I-Codes and 2020
    NEC for review, not the 2023 I-Codes.
  confidence: 0.96
  source_ids:
  - src:usa-fl:agency:fbc-annual-report-2021
- topic: Electrical model reference
  finding: The 8th Edition cycle-level electrical model reference is the 2020 NEC;
    the Florida Building Commission uses NFPA 70 in code updates.
  confidence: 0.95
  source_ids:
  - src:usa-fl:agency:fbc-annual-report-2021
  - src:usa-fl:statute:553.73
- topic: Fire code authority
  finding: The Florida Fire Prevention Code and Life Safety Code are adopted, modified,
    interpreted, and maintained by the Department of Financial Services / State Fire
    Marshal, not the Florida Building Commission.
  confidence: 0.99
  source_ids:
  - src:usa-fl:statute:553.73
  - src:usa-fl:statute:633.208
- topic: Fire code edition
  finding: The current Florida Fire Prevention Code is the 8th Edition (2023), effective
    December 31, 2023, incorporating NFPA 1 Florida 2021 Edition and NFPA 101 Florida
    2021 Edition through Rule Chapter 69A-60.
  confidence: 0.98
  source_ids:
  - src:usa-fl:agency:sfm-ffpc
  - src:usa-fl:rule:69a-60
  - src:usa-fl:rule:69a-60.003
  - src:usa-fl:rule:69a-60.004
- topic: Local enforcement
  finding: Local governments and legally constituted enforcement districts generally
    enforce the Florida Building Code, with state-agency exceptions listed in Fla.
    Stat. 553.80.
  confidence: 0.98
  source_ids:
  - src:usa-fl:statute:553.80
- topic: Local amendments
  finding: Local amendments must be more stringent, transmitted to the commission
    within 30 days, publicly available, and, for local technical amendments, are not
    effective until 30 days after receipt and publication by the commission.
  confidence: 0.98
  source_ids:
  - src:usa-fl:statute:553.73
- topic: Permit-date rule
  finding: The edition of the Florida Building Code in effect on the permit-application
    date governs permitted work for the life of the permit and extensions.
  confidence: 0.97
  source_ids:
  - src:usa-fl:statute:553.73
- topic: Expired-permit closure rule
  finding: An expired permit with substantially completed work may be closed under
    the building code in effect when the local enforcement agency received the permit
    application.
  confidence: 0.94
  source_ids:
  - src:usa-fl:statute:553.79
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-fl:florida-building-commission |
| Authority name | Florida Building Commission |
| Authority type | state_commission |
| Legal basis | Fla. Stat. 553.73 |
| Role | Adopts and maintains the Florida Building Code |
| Enforcement model | locally_enforced_with_state_exceptions |
| Source IDs | src:usa-fl:statute:553.73, src:usa-fl:statute:553.80 |
| Verification status | verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-fl:florida-building-commission | Florida Building Commission | Adopts statewide building code | Fla. Stat. 553.73 | src:usa-fl:statute:553.73 | verified |
| Residential | ahj:usa-fl:florida-building-commission | Florida Building Commission | Adopts residential volume | Fla. Stat. 553.73 | src:usa-fl:statute:553.73 | verified |
| Existing Building / Rehabilitation | ahj:usa-fl:florida-building-commission | Florida Building Commission | Adopts existing building volume | Fla. Stat. 553.73 | src:usa-fl:statute:553.73 | verified |
| Mechanical | ahj:usa-fl:florida-building-commission | Florida Building Commission | Adopts mechanical volume | Fla. Stat. 553.73 | src:usa-fl:statute:553.73 | verified |
| Plumbing | ahj:usa-fl:florida-building-commission | Florida Building Commission | Adopts plumbing volume | Fla. Stat. 553.73 | src:usa-fl:statute:553.73 | verified |
| Fuel Gas | ahj:usa-fl:florida-building-commission | Florida Building Commission | Adopts fuel-gas volume | Fla. Stat. 553.73 | src:usa-fl:statute:553.73 | verified |
| Electrical | ahj:usa-fl:florida-building-commission | Florida Building Commission | Adopts electrical provisions using NEC / NFPA 70 as the cycle-level model reference | Fla. Stat. 553.73(3) | src:usa-fl:statute:553.73, src:usa-fl:agency:fbc-annual-report-2021 | verified |
| Energy | ahj:usa-fl:florida-building-commission | Florida Building Commission | Adopts energy conservation volume | Fla. Stat. 553.73 | src:usa-fl:statute:553.73 | verified |
| Fire - construction references | ahj:usa-fl:state-fire-marshal | State Fire Marshal / Department of Financial Services | Maintains fire prevention and life-safety code references used alongside the building code | Fla. Stat. 553.73(1)(c) | src:usa-fl:statute:553.73 | verified |
| Fire - operational / prevention code | ahj:usa-fl:state-fire-marshal | State Fire Marshal / Department of Financial Services | Adopts and maintains the Florida Fire Prevention Code | Fla. Stat. 633.208; Rule Chapter 69A-60, F.A.C. | src:usa-fl:statute:633.208, src:usa-fl:rule:69a-60 | verified |
| Accessibility | ahj:usa-fl:florida-building-commission | Florida Building Commission | Incorporates technical portions of the Florida Accessibility Code for Building Construction into the Florida Building Code | Fla. Stat. 553.73(1)(b) | src:usa-fl:statute:553.73 | verified |
| Elevator / Conveyance | ahj:usa-fl:bureau-of-elevators | Bureau of Elevators, DBPR | Enforces elevator equipment under exclusive state authority | Fla. Stat. 553.80(1)(b) | src:usa-fl:statute:553.80 | verified |

### 2.3 Authority Hierarchy Notes

Florida is a local-enforcement state for most building work. The Florida Building Commission sets the statewide Florida Building Code. Counties, municipalities, and legally constituted enforcement districts generally enforce the code through permits, plan review, inspections, certificates, and local procedures. State agencies retain authority over defined subjects and facility classes, including correctional facilities, elevator equipment, certain health-care facilities, public schools and higher-education facilities, turnpike toll facilities, and secure mental-health treatment facilities.

Fire prevention and life safety are separated from the building-code adoption authority. The Florida Fire Prevention Code is adopted by the State Fire Marshal / Department of Financial Services, deemed adopted by municipalities, counties, and special districts with firesafety responsibilities, and enforced locally as the minimum firesafety code. Building-code and fire-code conflicts are resolved in favor of the requirement that provides the greater degree of life safety or equivalent life-safety protection.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-fl:001 | ahj:usa-fl:florida-building-commission | establishes_statewide_code_for | local governments, enforcement districts, and state agencies with applicable enabling authority | src:usa-fl:statute:553.73, src:usa-fl:statute:553.80 | verified |
| edge:usa-fl:002 | ahj:usa-fl:florida-building-commission | preempts_general_building_codes_except | authorized local amendments and statutory exceptions | src:usa-fl:statute:553.79, src:usa-fl:statute:553.73 | verified |
| edge:usa-fl:003 | ahj:usa-fl:state-fire-marshal | adopts_and_maintains | Florida Fire Prevention Code and Life Safety Code | src:usa-fl:statute:553.73, src:usa-fl:rule:69a-60 | verified |
| edge:usa-fl:004 | ahj:usa-fl:state-fire-marshal | local_fire_enforcement_by | municipalities, counties, and special districts with firesafety responsibilities | src:usa-fl:statute:633.208, src:usa-fl:agency:sfm-ffpc | verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | Florida Building Code, Building, 8th Edition (2023) | 2021 IBC, cycle-level model reference | 2023 | current | null | 2023-12-31 | 2023-12-31 | 2023-12-31 | The code effective on the permit-application date governs permitted work for the life of the permit and extensions. | src:usa-fl:agency:bcis-homepage, src:usa-fl:agency:fbc-annual-report-2021, src:usa-fl:statute:553.73 |
| Residential | Florida Building Code, Residential, 8th Edition (2023) | 2021 IRC, cycle-level model reference | 2023 | current | null | 2023-12-31 | 2023-12-31 | 2023-12-31 | Same statewide permit-application-date rule. | src:usa-fl:agency:bcis-homepage, src:usa-fl:agency:fbc-annual-report-2021, src:usa-fl:statute:553.73 |
| Existing Building / Rehabilitation | Florida Building Code, Existing Building, 8th Edition (2023) | 2021 IEBC, cycle-level model reference | 2023 | current | null | 2023-12-31 | 2023-12-31 | 2023-12-31 | Same statewide permit-application-date rule. | src:usa-fl:agency:bcis-homepage, src:usa-fl:agency:fbc-annual-report-2021, src:usa-fl:statute:553.73 |
| Mechanical | Florida Building Code, Mechanical, 8th Edition (2023) | 2021 IMC, cycle-level model reference | 2023 | current | null | 2023-12-31 | 2023-12-31 | 2023-12-31 | Same statewide permit-application-date rule. | src:usa-fl:agency:bcis-homepage, src:usa-fl:agency:fbc-annual-report-2021, src:usa-fl:statute:553.73 |
| Plumbing | Florida Building Code, Plumbing, 8th Edition (2023) | 2021 IPC, cycle-level model reference | 2023 | current | null | 2023-12-31 | 2023-12-31 | 2023-12-31 | Same statewide permit-application-date rule. | src:usa-fl:agency:bcis-homepage, src:usa-fl:agency:fbc-annual-report-2021, src:usa-fl:statute:553.73 |
| Fuel Gas | Florida Building Code, Fuel Gas, 8th Edition (2023) | 2021 IFGC, cycle-level model reference | 2023 | current | null | 2023-12-31 | 2023-12-31 | 2023-12-31 | Same statewide permit-application-date rule. | src:usa-fl:agency:bcis-homepage, src:usa-fl:agency:fbc-annual-report-2021, src:usa-fl:statute:553.73 |
| Electrical | Florida Building Code, Electrical, 8th Edition (2023) | 2020 NFPA 70 / NEC, cycle-level model reference | 2023 | current | null | 2023-12-31 | 2023-12-31 | 2023-12-31 | Same statewide permit-application-date rule. | src:usa-fl:agency:bcis-homepage, src:usa-fl:agency:fbc-annual-report-2021, src:usa-fl:statute:553.73 |
| Energy | Florida Building Code, Energy Conservation, 8th Edition (2023) | 2021 IECC, cycle-level model reference; volume-level energy alternatives not parsed | 2023 | current | null | 2023-12-31 | 2023-12-31 | 2023-12-31 | Same statewide permit-application-date rule; energy effective date may be delayed under statute if compliance software is not approved at least 3 months before the updated code effective date. | src:usa-fl:agency:bcis-homepage, src:usa-fl:agency:fbc-annual-report-2021, src:usa-fl:statute:553.73 |
| Fire - construction references | Florida Building Code fire and life-safety references | Florida Fire Prevention Code / Life Safety Code references | 2023 | current | null | 2023-12-31 | 2023-12-31 | 2023-12-31 | Fire references operate in conjunction with the building code; conflicts are resolved in favor of greater life safety or equivalent life-safety protection. | src:usa-fl:statute:553.73, src:usa-fl:statute:633.208 |
| Fire - operational / prevention code | Florida Fire Prevention Code, 8th Edition (2023) | NFPA 1, Fire Code, Florida 2021 Edition; NFPA 101, Life Safety Code, Florida 2021 Edition | 2023 | current | null | 2023-12-31 | 2023-12-31 | 2023-12-31 | Enforced by municipalities, counties, and special districts with firesafety responsibilities as the minimum firesafety code. | src:usa-fl:agency:sfm-ffpc, src:usa-fl:rule:69a-60, src:usa-fl:rule:69a-60.003, src:usa-fl:rule:69a-60.004, src:usa-fl:statute:633.208 |
| Accessibility | Florida Building Code, Accessibility, 8th Edition (2023) | Florida Accessibility Code for Building Construction; exact external base not parsed | 2023 | current | null | 2023-12-31 | 2023-12-31 | 2023-12-31 | Technical portions of the Florida Accessibility Code are contained in the Florida Building Code. | src:usa-fl:agency:bcis-homepage, src:usa-fl:statute:553.73 |
| Elevator / Conveyance | Florida Building Code elevator provisions / state elevator equipment regulation | ASME A17.1 reference unresolved | unknown | current | null | null | null | null | Elevator equipment enforcement is reserved to DBPR Bureau of Elevators; exact code-volume reference and date remain unresolved. | src:usa-fl:statute:553.80 |

### 3.2 Adoption Records

```yaml
adoptions:
  - adoption_id: "adoption:usa-fl:fbc:2023-cycle"
    code_families:
      - building
      - residential
      - existing_building
      - mechanical
      - plumbing
      - fuel_gas
      - electrical
      - energy
      - accessibility
    state_code: "Florida Building Code, 8th Edition (2023)"
    status: "current"
    adoption_date: null
    effective_date: "2023-12-31"
    operative_date: "2023-12-31"
    mandatory_date: "2023-12-31"
    model_code_basis:
      i_codes: "2021 I-Codes selected for review"
      electrical: "2020 NEC selected for review"
    authority_id: "ahj:usa-fl:florida-building-commission"
    source_ids:
      - "src:usa-fl:agency:bcis-homepage"
      - "src:usa-fl:agency:fbc-annual-report-2021"
      - "src:usa-fl:statute:553.73"
    verification_status: "partially_verified"
    notes: "Cycle-level base model verified; volume-specific code text parsing remains open."

  - adoption_id: "adoption:usa-fl:ffpc:2023"
    code_families:
      - fire_operational
    state_code: "Florida Fire Prevention Code, 8th Edition (2023)"
    status: "current"
    adoption_date: null
    effective_date: "2023-12-31"
    operative_date: "2023-12-31"
    mandatory_date: "2023-12-31"
    model_code_basis:
      nfpa_1: "NFPA 1, Fire Code, Florida 2021 Edition"
      nfpa_101: "NFPA 101, Life Safety Code, Florida 2021 Edition"
    authority_id: "ahj:usa-fl:state-fire-marshal"
    source_ids:
      - "src:usa-fl:agency:sfm-ffpc"
      - "src:usa-fl:rule:69a-60"
      - "src:usa-fl:rule:69a-60.003"
      - "src:usa-fl:rule:69a-60.004"
      - "src:usa-fl:statute:633.208"
    verification_status: "verified"
```

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

Florida uses a statewide effective-date and permit-application-date model. The Florida Building Code cycle currently in effect is the 8th Edition (2023), effective December 31, 2023. The edition in effect on the permit-application date governs permitted work for the life of the permit and extensions. Local technical amendments have a separate effectiveness rule: they are not effective until 30 days after receipt and publication by the Florida Building Commission and are generally effective only until the next statewide code edition unless incorporated or readopted. The current Florida Fire Prevention Code 8th Edition is also effective December 31, 2023.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-fl:fbc-2023-effective-date | statewide current building-code families | effective_date | 2023-12-31 | current 8th Edition Florida Building Code cycle | false | src:usa-fl:agency:bcis-homepage, src:usa-fl:agency:fbc-annual-report-2021 | verified |
| date-rule:usa-fl:permit-application-governs | permitted work under Florida Building Code | permit_application_date | code in effect on permit-application date | permit application received for work authorized by the code | true | src:usa-fl:statute:553.73 | verified |
| date-rule:usa-fl:update-six-month-minimum | updated Florida Building Code | minimum_delay | no sooner than 6 months after publication | rule updating the Florida Building Code | true | src:usa-fl:statute:553.73 | verified |
| date-rule:usa-fl:energy-software-delay | energy provisions | possible_delay | up to 3 additional months | energy compliance software not approved at least 3 months before updated-code effective date | true | src:usa-fl:statute:553.73 | verified |
| date-rule:usa-fl:permit-closure-old-code | expired permits with substantial completion | permit_closure_rule | code in effect when local enforcement agency received the permit application | local enforcement closes an expired permit with substantially completed requirements | true | src:usa-fl:statute:553.79 | verified |
| date-rule:usa-fl:local-amendment-publication | local technical building-code amendments | publication_delay | 30 days after receipt and publication | local amendment transmitted to commission | false | src:usa-fl:statute:553.73 | verified |
| date-rule:usa-fl:ffpc-2023-effective-date | Florida Fire Prevention Code | effective_date | 2023-12-31 | 8th Edition (2023) Florida Fire Prevention Code | false | src:usa-fl:agency:sfm-ffpc, src:usa-fl:rule:69a-60 | verified |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building and related FBC volumes | Draft 9th Edition (2026) Florida Building Code | unknown | unknown | unknown | unknown | unknown | monitoring | src:usa-fl:agency:bcis-homepage | The commission site links to draft 9th Edition materials; no adopted effective date was verified in this pass. |
| Fire - operational / prevention code | 9th Edition Florida Fire Prevention Code | unknown | unknown | unknown | unknown | unknown | monitoring | src:usa-fl:rule:69a-60.003, src:usa-fl:rule:69a-60.004 | Rule-history entries show 2025 and 2026 activity to develop or propose 9th Edition FFPC updates; current final adopted rules remain effective 2023-12-31. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-fl:state-exceptions | building, elevator, health-care, education, correctional, toll, secure mental health | statutory state-agency exceptions | work falls under an exception listed in Fla. Stat. 553.80 | Certain facility types or subjects are enforced by state agencies or through special review paths instead of the default local AHJ path. | src:usa-fl:statute:553.80 | verified |
| applicability-rule:usa-fl:fire-plan-review | fire, life_safety | building permit applications subject to firesafety review | permit plans require firesafety review | The appropriate certified firesafety inspector reviews plans for compliance with the Florida Fire Prevention Code and Life Safety Code unless a building or structure is not subject to a firesafety code. | src:usa-fl:statute:553.79 | verified |
| applicability-rule:usa-fl:one-two-family-fire-review | fire, residential | one-family and two-family detached dwellings | detached one- or two-family residential dwelling unit | Such units are not subject to local fire official plan review or inspection unless expressly made subject by local ordinance. | src:usa-fl:statute:553.79 | verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** Florida-specific provisions, technical amendments, statewide updates, errata, and referenced standards are layered into the statewide Florida Building Code through the Florida Building Commission process. Fire-code amendments are handled through Department of Financial Services / State Fire Marshal rulemaking and Rule Chapter 69A-60.

**Where amendments are published:** Florida Building Commission materials, Florida Building Code volumes, commission-published local amendment records, and Florida Administrative Code / Florida Administrative Register materials for fire-code rules.

**Amendment parsing status:** partial

### 5.2 State Amendment Sources

| Amendment Set ID | Code Family | Publication Format | Source IDs | Parsed? | Notes |
| --- | --- | --- | --- | --- | --- |
| amendment-set:usa-fl:fbc-2023-cycle | building, residential, existing_building, mechanical, plumbing, fuel_gas, electrical, energy, accessibility | code volumes / commission publications | src:usa-fl:agency:bcis-homepage, src:usa-fl:agency:fbc-annual-report-2021 | no | Current cycle baseline and model-code selection verified; detailed volume-level text parsing pending. |
| amendment-set:usa-fl:ffpc-2023 | fire_operational | Florida Administrative Code / State Fire Marshal FFPC publication | src:usa-fl:agency:sfm-ffpc, src:usa-fl:rule:69a-60, src:usa-fl:rule:69a-60.003, src:usa-fl:rule:69a-60.004 | partial | NFPA 1 and NFPA 101 edition/effective dates verified; detailed Florida-specific amendment text not parsed. |
| amendment-set:usa-fl:local-building-amendments | local building-code amendments | commission local amendment records | src:usa-fl:statute:553.73 | no | Statutory transmission, publication, and availability requirements verified; actual local amendment registry not parsed. |
| amendment-set:usa-fl:local-fire-amendments | local fire-code amendments | local jurisdiction records / State Fire Marshal local amendment link | src:usa-fl:agency:sfm-ffpc, src:usa-fl:statute:633.208 | no | Local fire amendments may exist and are enforceable only by the local jurisdiction; no statewide inventory was parsed. |

### 5.3 High-Impact State Amendments

| Code Family | Section | Amendment Type | Summary | Source IDs | Confidence |
| --- | --- | --- | --- | --- | --- |
| building | local technical amendments | process_limit | Local governments may adopt local technical amendments only when more stringent, locally justified, non-discriminatory, within the subject matter of the Florida Building Code, transmitted to the commission, publicly available, and published before becoming effective. | src:usa-fl:statute:553.73 | 0.98 |
| building | model-code updates | transition_limit | Updated Florida Building Code rules take effect no sooner than 6 months after publication, except immediate-threat amendments. | src:usa-fl:statute:553.73 | 0.97 |
| electrical | NEC update | cycle_reference | The 8th Edition (2023) process selected the 2020 NEC as the model electrical code for review. | src:usa-fl:agency:fbc-annual-report-2021 | 0.95 |
| fire_operational | local fire-safety amendments | process_limit | Local governments and special districts with firesafety responsibilities may adopt more stringent firesafety standards if statutory local-conditions and process requirements are met. | src:usa-fl:statute:633.208 | 0.96 |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-fl"
  model: "hybrid"
  enforcing_entities:
    - "municipality"
    - "county"
    - "legally constituted enforcement district with statutory authority"
    - "state agency where authorized by enabling legislation"
  required_or_referenced_officials:
    - "building_official"
    - "building_code_administrator"
    - "plans_examiner"
    - "inspector"
    - "certified_firesafety_inspector"
    - "local_fire_official"
  state_reserved_activities:
    - "correctional_facilities_under_Department_of_Corrections_or_Department_of_Juvenile_Justice"
    - "elevator_equipment_under_DBPR_Bureau_of_Elevators"
    - "healthcare_facility_plan_review_and_construction_survey_paths"
    - "public_schools_state_universities_and_Florida_College_System_institutions"
    - "turnpike_enterprise_toll_collection_facilities"
    - "secure_mental_health_treatment_facilities"
  source_ids:
    - "src:usa-fl:statute:553.79"
    - "src:usa-fl:statute:553.80"
    - "src:usa-fl:statute:633.208"
  verification_status: "verified"
  confidence: 0.98
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-fl"
  model: "allowed_if_stricter"
  applies_to_code_families:
    - "building"
    - "residential"
    - "existing_building"
    - "mechanical"
    - "plumbing"
    - "fuel_gas"
    - "electrical"
    - "energy"
    - "accessibility_limited_by_statute"
    - "fire_operational"
  building_code_local_amendments:
    approval_required_before_local_adoption: false
    approving_authority_id: null
    filing_required: true
    filing_recipient_authority_id: "ahj:usa-fl:florida-building-commission"
    transmission_deadline: "within_30_days_after_local_adoption"
    publication_required: true
    effective_delay: "not_effective_until_30_days_after_receipt_and_publication_by_commission"
    public_availability_required: true
    challenge_or_review_path: "countywide_compliance_review_board_for_challenged_local_technical_amendments"
    expiration_or_cycle_rule: "effective_only_until_next_new_FBC_edition_unless_commission_incorporates_or_local_government_readopts_after_rescission"
  fire_code_local_amendments:
    approval_required_before_local_adoption: false
    approving_authority_id: null
    more_stringent_allowed: true
    local_conditions_required: true
    local_jurisdiction_enforcement_only: true
    appeal_to_state_fire_marshal_verified: false
  registry_exists: true
  registry_source_ids:
    - "src:usa-fl:statute:553.73"
    - "src:usa-fl:agency:sfm-ffpc"
  legal_basis_source_ids:
    - "src:usa-fl:statute:553.73"
    - "src:usa-fl:statute:633.208"
  verification_status: "partially_verified"
  confidence: 0.96
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement authority and local amendment authority are distinct. Florida local governments usually administer building permits, plan review, and inspections under the statewide Florida Building Code. That enforcement role does not create unrestricted authority to rewrite the code. Local building-code amendments must satisfy statutory stringency, local-need, public-hearing, transmittal, publication, and availability requirements. Local fire-code amendments are also constrained and remain local in enforcement scope.

### 6.4 Known Local Amendment Registries

| Registry ID | Scope | Maintained By | Source ID | Machine Readable? | Parsed? | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| registry:usa-fl:local-building-amendments | statewide / local | Florida Building Commission | src:usa-fl:statute:553.73 | partial | no | The statute requires local amendments to be transmitted to the commission, maintained in a usable public format, and published before local technical amendments become effective. |
| registry:usa-fl:local-fire-amendments | local fire-code amendments | State Fire Marshal page links to access local amendments | src:usa-fl:agency:sfm-ffpc | partial | no | The State Fire Marshal page indicates local FFPC amendments may exist and are enforceable only by the local jurisdiction. |

### 6.5 Municipality-Specific Known Amendments

No municipality-specific building-code or fire-code amendment set was parsed for this pass.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: state_only

Expected jurisdiction stack:

Address
  -> State
  -> County
  -> Municipality / unincorporated county
  -> Special districts, if applicable
  -> Building AHJ
  -> Fire AHJ
  -> Trade-specific AHJs
  -> Applicable state code adoption records
  -> Applicable local amendment records

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- | --- |
| State | Florida legislative and commission sources | src:usa-fl:statute:553.73 | statewide | when statutes / rules change | verified |
| County | county GIS / local government sources | unresolved | statewide | varies | unresolved |
| Municipality | municipal GIS / local government sources | unresolved | statewide | varies | unresolved |
| Fire District | local fire authority sources | unresolved | varies | varies | unresolved |
| Special District | state agency or district source | unresolved | varies | varies | unresolved |

### 7.3 AHJ Contact Data

AHJ contact data has not been populated for Florida in this pass. State-level authority and enforcement model are verified, but jurisdiction-specific building and fire AHJ contacts remain unresolved.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Source Type | Publisher | URL | Accessed Date | Snapshot ID | Checksum | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-fl:agency:bcis-homepage | Florida Building Commission homepage | agency_page | Florida Building Commission | https://www.floridabuilding.org/ | 2026-06-25 | snapshot-pending | snapshot-pending | verified |
| src:usa-fl:agency:fbc-annual-report-2021 | Annual Report FY 2020-2021 / FBC Report to the 2022 Florida Legislature | agency_report_pdf | Florida Building Commission / DBPR | https://www.floridabuilding.org/fbc/commission/FBC_1021/2020-2021_Annual_Report.pdf | 2026-06-25 | snapshot-pending | snapshot-pending | verified |
| src:usa-fl:statute:553.73 | Florida Statutes 553.73, Florida Building Code | statute | Florida Legislature | https://www.leg.state.fl.us/Statutes/index.cfm?App_mode=Display_Statute&URL=0500-0599/0553/Sections/0553.73.html | 2026-06-25 | snapshot-pending | snapshot-pending | verified |
| src:usa-fl:statute:553.79 | Florida Statutes 553.79, permits; applications; issuance; inspections | statute | Florida Legislature | https://www.leg.state.fl.us/statutes/index.cfm?App_mode=Display_Statute&URL=0500-0599/0553/Sections/0553.79.html | 2026-06-25 | snapshot-pending | snapshot-pending | verified |
| src:usa-fl:statute:553.80 | Florida Statutes 553.80, enforcement | statute | Florida Legislature | https://www.leg.state.fl.us/statutes/index.cfm?App_mode=Display_Statute&URL=0500-0599/0553/Sections/0553.80.html | 2026-06-25 | snapshot-pending | snapshot-pending | verified |
| src:usa-fl:statute:633.208 | Florida Statutes 633.208, minimum firesafety standards | statute | Florida Legislature | https://www.leg.state.fl.us/statutes/index.cfm?App_mode=Display_Statute&URL=0600-0699/0633/Sections/0633.208.html | 2026-06-25 | snapshot-pending | snapshot-pending | verified |
| src:usa-fl:agency:sfm-ffpc | Florida Fire Prevention Code page | agency_page | Florida State Fire Marshal / Department of Financial Services | https://www.myfloridacfo.com/division/sfm/bfp/florida-fire-prevention-code | 2026-06-25 | snapshot-pending | snapshot-pending | verified |
| src:usa-fl:rule:69a-60 | Rule Chapter 69A-60, Florida Administrative Code, The Florida Fire Prevention Code | administrative_rule_chapter | Florida Department of State / Florida Administrative Code | https://flrules.org/gateway/ChapterHome.asp?Chapter=69A-60 | 2026-06-25 | snapshot-pending | snapshot-pending | verified |
| src:usa-fl:rule:69a-60.003 | Rule 69A-60.003, NFPA 1, the Fire Code, Florida 2021 Edition, Adopted | administrative_rule | Florida Department of State / Florida Administrative Code | https://flrules.org/gateway/RuleNo.asp?ID=69A-60.003&title=THE+FLORIDA+FIRE+PREVENTION+CODE | 2026-06-25 | snapshot-pending | snapshot-pending | verified |
| src:usa-fl:rule:69a-60.004 | Rule 69A-60.004, NFPA 101, Life Safety Code, Florida 2021 Edition, Adopted | administrative_rule | Florida Department of State / Florida Administrative Code | https://flrules.org/gateway/RuleNo.asp?ID=69A-60.004&title=THE+FLORIDA+FIRE+PREVENTION+CODE | 2026-06-25 | snapshot-pending | snapshot-pending | verified |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-fl:agency:bcis-homepage | homepage_summary | The homepage is a current-summary source; use code volumes, rules, and archived commission documents for cycle-by-cycle history and volume-specific adoption details. | informational_for_current_cycle_status |
| src:usa-fl:agency:fbc-annual-report-2021 | agency_report_pdf | Official commission PDF supports the 2021 I-Codes / 2020 NEC cycle-level model-code selection and scheduled effective date; it is not a substitute for parsing each final code volume. | authoritative_for_cycle_level_selection |
| src:usa-fl:statute:553.73 | codified_law | Statute defines authority, update process, local amendment rules, and permit-application-date logic, but does not itself list every current volume-specific code section. | authoritative |
| src:usa-fl:statute:553.79 | codified_law | Statute defines permitting and closure rules; local procedures may add administrative detail. | authoritative |
| src:usa-fl:statute:553.80 | codified_law | Statute defines enforcement allocation and state exceptions; facility-specific enabling laws may add detail not parsed here. | authoritative |
| src:usa-fl:agency:sfm-ffpc | agency_page | State Fire Marshal page is a current-summary page for FFPC access, edition, and local amendment access; rule text remains the controlling source for adopted rules. | authoritative_for_current_ffpc_summary |
| src:usa-fl:rule:69a-60 | administrative_rule_chapter | Rule chapter listing verifies current rule titles and effective dates; individual rule text should be parsed for detailed amendment text. | authoritative |
| src:usa-fl:rule:69a-60.003 | administrative_rule | Rule page verifies NFPA 1 Florida 2021 Edition and effective date; detailed incorporated text may require NFPA access. | authoritative_for_rule_metadata |
| src:usa-fl:rule:69a-60.004 | administrative_rule | Rule page verifies NFPA 101 Florida 2021 Edition and effective date; detailed incorporated text may require NFPA access. | authoritative_for_rule_metadata |

### 8.3 Supplemental Sources

None used in this pass.

### 8.4 Source Extraction Metadata

| Source ID | Parser | Parser Version | Extracted At | Extraction Confidence | Requires OCR | Requires Browser Automation | Manual Review Required |
| --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-fl:agency:bcis-homepage | browser_manual | 0.1 | 2026-06-25T00:00:00Z | 0.95 | no | yes | yes |
| src:usa-fl:agency:fbc-annual-report-2021 | browser_pdf_text_and_screenshot | 0.1 | 2026-06-25T00:00:00Z | 0.98 | no | no | yes |
| src:usa-fl:statute:553.73 | browser_manual | 0.1 | 2026-06-25T00:00:00Z | 0.99 | no | yes | yes |
| src:usa-fl:statute:553.79 | browser_manual | 0.1 | 2026-06-25T00:00:00Z | 0.99 | no | yes | yes |
| src:usa-fl:statute:553.80 | browser_manual | 0.1 | 2026-06-25T00:00:00Z | 0.99 | no | yes | yes |
| src:usa-fl:statute:633.208 | browser_manual | 0.1 | 2026-06-25T00:00:00Z | 0.99 | no | yes | yes |
| src:usa-fl:agency:sfm-ffpc | browser_manual | 0.1 | 2026-06-25T00:00:00Z | 0.98 | no | yes | yes |
| src:usa-fl:rule:69a-60 | browser_manual | 0.1 | 2026-06-25T00:00:00Z | 0.98 | no | yes | yes |
| src:usa-fl:rule:69a-60.003 | browser_manual | 0.1 | 2026-06-25T00:00:00Z | 0.98 | no | yes | yes |
| src:usa-fl:rule:69a-60.004 | browser_manual | 0.1 | 2026-06-25T00:00:00Z | 0.98 | no | yes | yes |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report:usa-fl | report.status | partially_verified | verified | 0.90 | src:usa-fl:agency:bcis-homepage, src:usa-fl:statute:553.73, src:usa-fl:rule:69a-60 | Core authority and current editions verified; AHJ contacts and volume-level parsing remain open. |
| adoption:usa-fl:fbc:2023-cycle | state_code.name | Florida Building Code, 8th Edition (2023) | verified | 0.99 | src:usa-fl:agency:bcis-homepage, src:usa-fl:agency:fbc-annual-report-2021 | Current statewide baseline. |
| adoption:usa-fl:fbc:2023-cycle | dates.effective_date | 2023-12-31 | verified | 0.99 | src:usa-fl:agency:bcis-homepage, src:usa-fl:agency:fbc-annual-report-2021 | Homepage and commission report support the date. |
| adoption:usa-fl:fbc:2023-cycle | model_code_basis | 2021 I-Codes and 2020 NEC | verified | 0.96 | src:usa-fl:agency:fbc-annual-report-2021 | Corrects prior 2023 I-Code base-model entries. |
| adoption:usa-fl:ffpc:2023 | state_code.name | Florida Fire Prevention Code, 8th Edition (2023) | verified | 0.98 | src:usa-fl:agency:sfm-ffpc, src:usa-fl:rule:69a-60 | Current fire-code baseline. |
| adoption:usa-fl:ffpc:2023 | model_code_basis | NFPA 1 Florida 2021 Edition and NFPA 101 Florida 2021 Edition | verified | 0.98 | src:usa-fl:rule:69a-60.003, src:usa-fl:rule:69a-60.004 | Rule metadata verified. |
| ahj:usa-fl:state-fire-marshal | authority.role | adopts and maintains FFPC / Life Safety Code | verified | 0.99 | src:usa-fl:statute:553.73, src:usa-fl:statute:633.208 | Separate from Florida Building Commission authority. |
| local-enforcement:usa-fl | model | hybrid | verified | 0.98 | src:usa-fl:statute:553.80 | Local enforcement plus state exceptions. |
| local-amendment-rule:usa-fl | model | allowed_if_stricter | verified | 0.96 | src:usa-fl:statute:553.73, src:usa-fl:statute:633.208 | State pre-approval before local adoption was not verified; transmission/publication/challenge rules are verified. |
| date-rule:usa-fl:permit-application-governs | rule | permit-application-date code governs permitted work | verified | 0.97 | src:usa-fl:statute:553.73 | Applies for life of permit and extensions. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| Template markers absent | pass | The repository template-marker scan returned no remaining matches. |
| All source IDs resolve | pass | All `src:usa-fl:*` identifiers cited in the body are listed in section 8. |
| All authority IDs resolve | pass | Authority IDs used in the report are defined in section 2. |
| All current code families have adoption rows | partial | Matrix rows are present, but detailed volume-level adoption records remain open. |
| Building and operational fire code are separated | pass | Separate authority and rule tracks are captured. |
| Fire code edition resolved | pass | FFPC 8th Edition (2023), NFPA 1 Florida 2021, and NFPA 101 Florida 2021 are verified. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Known dates are separated; unknown adoption dates remain null. |
| Effective dates are valid ISO dates | pass | Known date values use ISO format. |
| No impossible date sequences | pass | No contradictory sequences were introduced. |
| Transition rules have explicit trigger conditions | pass | Permit-application, 6-month publication, energy-software delay, local-amendment publication, and expired-permit closure rules are stated. |
| Permit-date logic is captured where applicable | pass | Section 4.2 includes both general permit-application-date and expired-permit closure rules. |
| Local enforcement model classified | pass | Section 6.1 uses hybrid. |
| Local amendment rule classified | pass | Section 6.2 uses allowed_if_stricter and separates building-code and fire-code paths. |
| AHJ confirmation metadata present | fail | AHJ contact rows and local boundary sources remain unresolved. |
| Official-source caveats captured | pass | Section 8.2 includes caveats for homepage, PDF, statutes, and rules. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-fl:001 | medium | Volume-level adoption records | The report verifies the current statewide cycle and cycle-level model-code references, but does not parse every final 8th Edition volume, supplement, or section-level amendment. | Parse final code volumes and supplements and expand normalized adoption records by code family. | null | null | open |
| issue:usa-fl:002 | low | AHJ registry and boundary data | County, municipal, fire-district, special-district boundaries and AHJ contact data are not populated. | Add jurisdiction-specific boundary sources and AHJ contact records. | null | null | open |
| issue:usa-fl:003 | medium | 2026 update monitoring | Draft/proposed 9th Edition (2026) activity exists, but no adopted effective date was verified in this pass. | Monitor Florida Building Commission and State Fire Marshal rulemaking for final 2026 adoption and effective dates. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-fl:agency-code-page | src:usa-fl:agency:bcis-homepage | html_diff | monthly | code edition/date text changes or 9th Edition materials become final | 2026-06-25 | active |
| watch:usa-fl:fbc-statute | src:usa-fl:statute:553.73 | statute_text_diff | quarterly | building-code authority, permit-date, amendment-process, or update-timing changes | 2026-06-25 | active |
| watch:usa-fl:permitting-statute | src:usa-fl:statute:553.79 | statute_text_diff | quarterly | permit issuance, plan review, expired permit, or special applicability changes | 2026-06-25 | active |
| watch:usa-fl:enforcement-statute | src:usa-fl:statute:553.80 | statute_text_diff | quarterly | enforcement allocation or state-exception changes | 2026-06-25 | active |
| watch:usa-fl:ffpc-page | src:usa-fl:agency:sfm-ffpc | html_diff | monthly | FFPC edition/effective-date changes or local amendment access changes | 2026-06-25 | active |
| watch:usa-fl:ffpc-rule-chapter | src:usa-fl:rule:69a-60 | rule_chapter_diff | monthly | proposed or adopted FFPC rule activity | 2026-06-25 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-25 | Updated Florida report; corrected 8th Edition base model references from 2023 I-Codes to 2021 I-Codes / 2020 NEC; verified FFPC 8th Edition and NFPA 1 / NFPA 101 Florida 2021 rule metadata; revised local amendment pre-approval posture; replaced unresolved source placeholders. | report:usa-fl, adoption:usa-fl:fbc:2023-cycle, adoption:usa-fl:ffpc:2023, local-amendment-rule:usa-fl, validation:usa-fl | src:usa-fl:agency:bcis-homepage, src:usa-fl:agency:fbc-annual-report-2021, src:usa-fl:statute:553.73, src:usa-fl:statute:553.79, src:usa-fl:statute:553.80, src:usa-fl:statute:633.208, src:usa-fl:agency:sfm-ffpc, src:usa-fl:rule:69a-60, src:usa-fl:rule:69a-60.003, src:usa-fl:rule:69a-60.004 | ChatGPT | Partially verified draft; volume-level code parsing and AHJ registry remain open. |
