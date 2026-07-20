---
state:
  state_id: "US-WY"
  name: "Wyoming"
  abbreviation: "WY"
report:
  report_id: "state-report:usa-wy"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-26"
  last_verified: null
  reviewed_by: null
risk:
  overall_confidence: 0.58 # 0.00 - 1.00
  risk_flags:
    - "standalone_residential_plumbing_energy_scope_unresolved"
    - "state_amendment_text_not_fully_parsed"
    - "local_delegation_registry_not_captured"
    - "ahj_contact_data_not_populated"
    - "elevator_conveyance_authority_unresolved"
  open_questions_count: 6

---

# State Building Code Authority Report: Wyoming

## 1. Executive Summary

- **Authority model:** Wyoming uses a state-administered minimum-code model for fire, building, existing-building, mechanical, fuel-gas, and electrical safety standards, with optional delegation of local enforcement to municipalities and counties that apply and meet statutory conditions. The Council on Fire Prevention and Electrical Safety in Buildings adopts the minimum fire/building-related standards, while the Department of Fire Prevention and Electrical Safety and State Fire Marshal administer and enforce the program unless local enforcement authority has been delegated.

- **Statewide code status:** Core statewide/currently identified adoptions are the 2024 International Fire Code, 2024 International Building Code, 2024 International Fuel Gas Code, 2024 International Mechanical Code, 2024 International Existing Building Code, and the 2023 National Electrical Code. Standalone statewide adoption of the International Residential Code, International Plumbing Code, and International Energy Conservation Code was not confirmed; the Council rule text mirrored by Cornell indicates only referenced provisions of those codes are incorporated where referenced by the adopted codes.

- **Local enforcement model:** Municipalities and counties may receive complete delegated authority to enforce and interpret local or state fire, building, existing-building, or electrical safety standards if their standards and inspectors meet statutory requirements. If delegated jurisdictions do not keep standards at least as strong as state standards after state updates, enforcement authority reverts to the Department.

- **Local amendment posture:** Delegated local jurisdictions must adopt standards by ordinance or resolution that are equivalent to or more stringent than applicable Department standards. For certain state-owned or leased buildings, including schools, local provisions that are more stringent than state codes prevail where the local governmental entity has assumed the authorized plan-review/inspection role. A statewide registry of local amendments or delegations was not captured in this pass.

- **Known transition periods or pending changes:** Official sources reviewed show the 2024 International Codes rule effective on 2024-06-26 and the State Fire Marshal plan-review page states that new plan review submissions from 2024-06-28 must comply. The 2023 NEC is stated effective 2023-07-01. Local delegated jurisdictions have a six-month statutory update window after Department adoption of new standards.

- **Production readiness:** partially_ready_for_authority_and_core_code_fields; not_ready_for_full_jurisdiction_resolution

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: The Council on Fire Prevention and Electrical Safety in Buildings adopts
    minimum standards for fire/building-related codes; DFPES and the State Fire Marshal
    administer/enforce.
  confidence: 0.78
  source_ids:
  - src:usa-wy:statutes:title-35-ch9
  - src:usa-wy:rules:council-ch1-2024
  - src:usa-wy:dfpes:plan-review
- topic: Primary building code edition
  finding: 2024 IBC is the current verified statewide building-code base for DFPES-administered
    plan review and minimum standards.
  confidence: 0.82
  source_ids:
  - src:usa-wy:rules:council-ch1-2024
  - src:usa-wy:dfpes:plan-review
  - src:usa-wy:rules:council-ch1-lii
- topic: Fire code authority
  finding: 2024 IFC is adopted under the Council/DFPES framework and listed by the
    State Fire Marshal as an adopted state code.
  confidence: 0.82
  source_ids:
  - src:usa-wy:rules:council-ch1-2024
  - src:usa-wy:dfpes:plan-review
  - src:usa-wy:statutes:title-35-ch9
- topic: Electrical code authority
  finding: Electrical installations are subject to the NEC; the State Fire Marshal
    page identifies the 2023 NEC as effective 2023-07-01, and the chief electrical
    inspector enforces minimum electrical requirements unless local authority is delegated.
  confidence: 0.76
  source_ids:
  - src:usa-wy:statutes:title-35-ch9
  - src:usa-wy:rules:electrical-board-2023
  - src:usa-wy:dfpes:plan-review
- topic: Local enforcement
  finding: 'Local enforcement is optional/delegated, not purely home-rule: municipalities
    and counties must apply, adopt equivalent-or-stronger standards, maintain certified
    inspectors, and receive transfer of authority.'
  confidence: 0.8
  source_ids:
  - src:usa-wy:statutes:title-35-ch9
- topic: Health care facilities
  finding: The Department of Health has jurisdiction over construction/remodeling
    of state-licensed health care facilities, except electrical installation; NFPA
    101 or federal life-safety certification requirements can control.
  confidence: 0.75
  source_ids:
  - src:usa-wy:statutes:title-35-ch9
  - src:usa-wy:health:chapter-20-jurisdiction
  - src:usa-wy:health:hls-construction
- topic: Effective / operative date rule
  finding: 2024 International Codes rule effective date is 2024-06-26; State Fire
    Marshal page uses 2024-06-28 as the new-plan-review compliance trigger.
  confidence: 0.72
  source_ids:
  - src:usa-wy:rules:council-ch1-2024
  - src:usa-wy:dfpes:plan-review
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-wy:dfpes-state-fire-marshal |
| Authority name | Wyoming Department of Fire Prevention and Electrical Safety / Wyoming State Fire Marshal |
| Authority type | state agency / state fire marshal |
| Legal basis | W.S. 35-9-101 through 35-9-130; State Fire Marshal appointed as director of the Department; Council and Electrical Board created within the Department |
| Role | Administers the fire prevention and electrical safety program, performs plan review and inspections within state jurisdiction, and enforces minimum fire and electrical safety standards unless authority is delegated locally |
| Enforcement model | state_default_with_optional_local_delegation |
| Source IDs | src:usa-wy:statutes:title-35-ch9; src:usa-wy:dfpes:plan-review |
| Verification status | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-wy:council-fire-prevention-electrical-safety | Council on Fire Prevention and Electrical Safety in Buildings | Adopts minimum standards not exceeding IBC/related model codes; hears certain appeals/variances | W.S. 35-9-103; W.S. 35-9-106 | src:usa-wy:statutes:title-35-ch9; src:usa-wy:rules:council-ch1-2024 | partially_verified |
| Residential | ahj:usa-wy:dfpes-state-fire-marshal | DFPES / State Fire Marshal | Standalone statewide IRC authority unresolved; referenced IRC provisions appear only where referenced in adopted codes | Council rule mirror; official text not fully extracted | src:usa-wy:rules:council-ch1-lii | unresolved |
| Existing Building / Rehabilitation | ahj:usa-wy:council-fire-prevention-electrical-safety | Council on Fire Prevention and Electrical Safety in Buildings | Adopts IEBC as a minimum standard and local delegated jurisdictions may enforce/interpret existing-building standards | W.S. 35-9-106; W.S. 35-9-121 | src:usa-wy:statutes:title-35-ch9; src:usa-wy:rules:council-ch1-2024; src:usa-wy:dfpes:plan-review | partially_verified |
| Mechanical | ahj:usa-wy:council-fire-prevention-electrical-safety | Council on Fire Prevention and Electrical Safety in Buildings | Adopts IMC as a minimum standard | W.S. 35-9-106 | src:usa-wy:statutes:title-35-ch9; src:usa-wy:rules:council-ch1-2024; src:usa-wy:dfpes:plan-review | partially_verified |
| Plumbing | ahj:usa-wy:unresolved-plumbing | Unresolved | Standalone statewide IPC adoption/agency not confirmed; referenced IPC provisions appear only where referenced in adopted codes | Council rule mirror; plumbing-specific authority not parsed | src:usa-wy:rules:council-ch1-lii | unresolved |
| Fuel Gas | ahj:usa-wy:council-fire-prevention-electrical-safety | Council on Fire Prevention and Electrical Safety in Buildings | Adopts IFGC as a minimum standard; LP gas conflict rule reserves NFPA 58 and ANSI Z223.1/NFPA 54 priority | W.S. 35-9-106 | src:usa-wy:statutes:title-35-ch9; src:usa-wy:rules:council-ch1-2024; src:usa-wy:dfpes:plan-review | partially_verified |
| Electrical | ahj:usa-wy:chief-electrical-inspector | Chief Electrical Inspector / Electrical Safety Division | Enforces minimum electrical installation requirements and supervises deputy electrical inspectors unless local electrical enforcement authority has been delegated | W.S. 35-9-119; W.S. 35-9-120 | src:usa-wy:statutes:title-35-ch9; src:usa-wy:dfpes:plan-review | partially_verified |
| Energy | ahj:usa-wy:unresolved-energy | Unresolved | Standalone statewide IECC adoption/agency not confirmed; referenced IECC provisions appear only where referenced in adopted codes | Council rule mirror; energy-specific authority not parsed | src:usa-wy:rules:council-ch1-lii | unresolved |
| Fire - construction references | ahj:usa-wy:council-fire-prevention-electrical-safety | Council on Fire Prevention and Electrical Safety in Buildings | Adopts IFC and building-related fire standards | W.S. 35-9-106 | src:usa-wy:statutes:title-35-ch9; src:usa-wy:rules:council-ch1-2024; src:usa-wy:dfpes:plan-review | partially_verified |
| Fire - operational / prevention code | ahj:usa-wy:dfpes-state-fire-marshal | DFPES / State Fire Marshal | Implements fire safety programs, preventive inspections, corrective activities, and statewide fire-safety enforcement within state jurisdiction | W.S. 35-9-107 | src:usa-wy:statutes:title-35-ch9; src:usa-wy:dfpes:plan-review | partially_verified |
| Accessibility | ahj:usa-wy:unresolved-accessibility | Unresolved | State accessibility code authority not parsed; federal ADA/FHA overlays outside this state-code pass | unresolved | none | unresolved |
| Elevator / Conveyance | ahj:usa-wy:unresolved-elevator | Unresolved | Elevator/conveyance inspection authority not resolved beyond a general Title 27 safety-device statute hit | unresolved | src:usa-wy:statutes:title-27-elevator-safety-device | unresolved |
| Health care facilities | ahj:usa-wy:department-of-health-hls | Wyoming Department of Health / Healthcare Licensing and Surveys | Jurisdiction over construction/remodeling of state-licensed health care facilities except electrical installation; may delegate plan review/inspection to counties/municipalities with certified personnel | W.S. 35-9-121.1; Department of Health Chapter 20 | src:usa-wy:statutes:title-35-ch9; src:usa-wy:health:chapter-20-jurisdiction; src:usa-wy:health:hls-construction | partially_verified |

### 2.3 Authority Hierarchy Notes

The verified Wyoming model separates **adoption**, **state administration/enforcement**, and **delegated local enforcement**. The Council adopts minimum fire/building-related standards. The State Fire Marshal is the director of DFPES and implements statewide fire safety programs, plan review, and inspections where state jurisdiction applies. The chief electrical inspector administers NEC-based electrical enforcement unless a municipality or county has received electrical enforcement authority. Municipalities and counties may enforce and interpret local or state standards only after statutory delegation conditions are met.

Health care facilities are a specialized track: the Department of Health has jurisdiction over construction/remodeling of state-licensed health care facilities except electrical installation. For those facilities, NFPA 101/federal certification requirements can supersede conflicting state or local code requirements.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-wy:001 | ahj:usa-wy:council-fire-prevention-electrical-safety | adopts_minimum_standards_for | IFC, IBC, IMC, IEBC, IFGC and referenced provisions | src:usa-wy:statutes:title-35-ch9; src:usa-wy:rules:council-ch1-2024 | partially_verified |
| edge:usa-wy:002 | ahj:usa-wy:dfpes-state-fire-marshal | administers_and_enforces | state plan review, fire prevention, state buildings, nondelegated jurisdictions | src:usa-wy:statutes:title-35-ch9; src:usa-wy:dfpes:plan-review | partially_verified |
| edge:usa-wy:003 | ahj:usa-wy:dfpes-state-fire-marshal | delegates_enforcement_to | municipalities and counties that apply and meet W.S. 35-9-121 conditions | src:usa-wy:statutes:title-35-ch9 | partially_verified |
| edge:usa-wy:004 | ahj:usa-wy:chief-electrical-inspector | enforces_except_when_delegated | NEC-based electrical installation requirements | src:usa-wy:statutes:title-35-ch9 | partially_verified |
| edge:usa-wy:005 | ahj:usa-wy:department-of-health-hls | has_jurisdiction_over | state-licensed health care facility construction/remodeling except electrical installation | src:usa-wy:statutes:title-35-ch9; src:usa-wy:health:chapter-20-jurisdiction | partially_verified |
| edge:usa-wy:006 | ahj:usa-wy:department-of-health-hls | delegates_plan_review_and_inspection_to | certified counties or municipalities upon written request | src:usa-wy:statutes:title-35-ch9; src:usa-wy:health:hls-construction | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | Wyoming minimum fire/building standards | International Building Code | 2024 | current_verified | 2024-06-26 | 2024-06-26 | 2024-06-28 | 2024-06-28 | Rule effective 2024-06-26; all new plan review submissions from 2024-06-28 must comply | src:usa-wy:rules:council-ch1-2024; src:usa-wy:dfpes:plan-review; src:usa-wy:rules:council-ch1-lii |
| Residential | Referenced residential provisions only | International Residential Code | not confirmed | referenced_only_unresolved | null | null | null | null | Standalone statewide IRC not confirmed; mirrored Council rule lists only referenced provisions where referenced by adopted codes | src:usa-wy:rules:council-ch1-lii |
| Existing Building / Rehabilitation | Wyoming minimum existing-building standards | International Existing Building Code | 2024 | current_verified | 2024-06-26 | 2024-06-26 | 2024-06-28 | 2024-06-28 | Rule effective 2024-06-26; all new plan review submissions from 2024-06-28 must comply | src:usa-wy:rules:council-ch1-2024; src:usa-wy:dfpes:plan-review; src:usa-wy:rules:council-ch1-lii |
| Mechanical | Wyoming minimum mechanical standards | International Mechanical Code | 2024 | current_verified | 2024-06-26 | 2024-06-26 | 2024-06-28 | 2024-06-28 | Rule effective 2024-06-26; all new plan review submissions from 2024-06-28 must comply | src:usa-wy:rules:council-ch1-2024; src:usa-wy:dfpes:plan-review; src:usa-wy:rules:council-ch1-lii |
| Plumbing | Referenced plumbing provisions only | International Plumbing Code | not confirmed | referenced_only_unresolved | null | null | null | null | Standalone statewide IPC not confirmed; mirrored Council rule lists only referenced provisions where referenced by adopted codes | src:usa-wy:rules:council-ch1-lii |
| Fuel Gas | Wyoming minimum fuel-gas standards | International Fuel Gas Code | 2024 | current_verified | 2024-06-26 | 2024-06-26 | 2024-06-28 | 2024-06-28 | Rule effective 2024-06-26; all new plan review submissions from 2024-06-28 must comply; LP gas conflict rule preserved separately | src:usa-wy:rules:council-ch1-2024; src:usa-wy:dfpes:plan-review; src:usa-wy:statutes:title-35-ch9; src:usa-wy:rules:council-ch1-lii |
| Electrical | Wyoming electrical safety standards | National Electrical Code | 2023 | current_verified | null | 2023-07-01 | 2023-07-01 | 2023-07-01 | State Fire Marshal page states 2023 NEC effective 2023-07-01; separate formal adoption date not parsed | src:usa-wy:dfpes:plan-review; src:usa-wy:statutes:title-35-ch9; src:usa-wy:rules:electrical-board-2023 |
| Energy | Referenced energy provisions only | International Energy Conservation Code | not confirmed | referenced_only_unresolved | null | null | null | null | Standalone statewide IECC not confirmed; mirrored Council rule lists only referenced provisions where referenced by adopted codes | src:usa-wy:rules:council-ch1-lii |
| Fire - construction references | Wyoming minimum fire standards | International Fire Code | 2024 | current_verified | 2024-06-26 | 2024-06-26 | 2024-06-28 | 2024-06-28 | Rule effective 2024-06-26; all new plan review submissions from 2024-06-28 must comply | src:usa-wy:rules:council-ch1-2024; src:usa-wy:dfpes:plan-review; src:usa-wy:rules:council-ch1-lii |
| Fire - operational / prevention code | Wyoming minimum fire standards | International Fire Code | 2024 | current_verified | 2024-06-26 | 2024-06-26 | 2024-06-28 | 2024-06-28 | Operational/prevention scope is tied to state fire-safety enforcement and IFC-based standards; local delegation can shift enforcement | src:usa-wy:rules:council-ch1-2024; src:usa-wy:dfpes:plan-review; src:usa-wy:statutes:title-35-ch9 |
| Accessibility | unresolved | unresolved | unresolved | unresolved | null | null | null | null | State accessibility adoption not parsed | none |
| Elevator / Conveyance | unresolved | unresolved | unresolved | unresolved | null | null | null | null | Elevator/conveyance code adoption not parsed | src:usa-wy:statutes:title-27-elevator-safety-device |

### 3.2 Adoption Records

#### code-adoption:usa-wy:ifc-2024

- **Authority ID:** ahj:usa-wy:council-fire-prevention-electrical-safety
- **Code family:** Fire - construction references; Fire - operational / prevention code
- **State code name:** Wyoming minimum fire standards
- **Base model code:** International Fire Code
- **Edition:** 2024
- **Adoption/effective basis:** Council rule effective 2024-06-26; State Fire Marshal page states Wyoming adopted the 2024 International Codes on 2024-06-28 and requires new plan review submissions from 2024-06-28 to comply.
- **Source IDs:** src:usa-wy:rules:council-ch1-2024; src:usa-wy:dfpes:plan-review; src:usa-wy:rules:council-ch1-lii
- **Record status:** partially_verified

#### code-adoption:usa-wy:ibc-2024

- **Authority ID:** ahj:usa-wy:council-fire-prevention-electrical-safety
- **Code family:** Building
- **State code name:** Wyoming minimum building/fire-safety standards
- **Base model code:** International Building Code
- **Edition:** 2024
- **Adoption/effective basis:** Council rule effective 2024-06-26; State Fire Marshal page applies 2024 IBC to new plan review submissions from 2024-06-28.
- **Source IDs:** src:usa-wy:rules:council-ch1-2024; src:usa-wy:dfpes:plan-review; src:usa-wy:rules:council-ch1-lii
- **Record status:** partially_verified

#### code-adoption:usa-wy:ifgc-2024

- **Authority ID:** ahj:usa-wy:council-fire-prevention-electrical-safety
- **Code family:** Fuel Gas
- **State code name:** Wyoming minimum fuel-gas standards
- **Base model code:** International Fuel Gas Code
- **Edition:** 2024
- **Adoption/effective basis:** Council rule effective 2024-06-26; State Fire Marshal page applies 2024 IFGC to new plan review submissions from 2024-06-28.
- **Important conflict rule:** For liquefied petroleum gas installations, current NFPA 58 and ANSI Z223.1/NFPA 54 control over conflicting IFGC provisions.
- **Source IDs:** src:usa-wy:rules:council-ch1-2024; src:usa-wy:dfpes:plan-review; src:usa-wy:statutes:title-35-ch9; src:usa-wy:rules:council-ch1-lii
- **Record status:** partially_verified

#### code-adoption:usa-wy:imc-2024

- **Authority ID:** ahj:usa-wy:council-fire-prevention-electrical-safety
- **Code family:** Mechanical
- **State code name:** Wyoming minimum mechanical standards
- **Base model code:** International Mechanical Code
- **Edition:** 2024
- **Adoption/effective basis:** Council rule effective 2024-06-26; State Fire Marshal page applies 2024 IMC to new plan review submissions from 2024-06-28.
- **Source IDs:** src:usa-wy:rules:council-ch1-2024; src:usa-wy:dfpes:plan-review; src:usa-wy:rules:council-ch1-lii
- **Record status:** partially_verified

#### code-adoption:usa-wy:iebc-2024

- **Authority ID:** ahj:usa-wy:council-fire-prevention-electrical-safety
- **Code family:** Existing Building / Rehabilitation
- **State code name:** Wyoming minimum existing-building standards
- **Base model code:** International Existing Building Code
- **Edition:** 2024
- **Adoption/effective basis:** Council rule effective 2024-06-26; State Fire Marshal page applies 2024 IEBC to new plan review submissions from 2024-06-28.
- **Source IDs:** src:usa-wy:rules:council-ch1-2024; src:usa-wy:dfpes:plan-review; src:usa-wy:rules:council-ch1-lii
- **Record status:** partially_verified

#### code-adoption:usa-wy:nec-2023

- **Authority ID:** ahj:usa-wy:chief-electrical-inspector
- **Code family:** Electrical
- **State code name:** Wyoming electrical safety standards
- **Base model code:** National Electrical Code
- **Edition:** 2023
- **Adoption/effective basis:** State Fire Marshal page states the 2023 NEC is effective 2023-07-01; W.S. 35-9-120 subjects electrical installations to applicable NEC minimum requirements.
- **Source IDs:** src:usa-wy:dfpes:plan-review; src:usa-wy:statutes:title-35-ch9; src:usa-wy:rules:electrical-board-2023
- **Record status:** partially_verified

#### code-adoption:usa-wy:referenced-irc-ipc-iecc-provisions

- **Authority ID:** ahj:usa-wy:council-fire-prevention-electrical-safety
- **Code family:** Residential; Plumbing; Energy; property-maintenance cross-references
- **State code name:** Referenced provisions in adopted model codes
- **Base model code:** IRC, IPMC, IPC, IECC provisions referenced in chapters 1 through 10 of adopted I-Codes
- **Edition:** Not independently confirmed in this pass
- **Adoption/effective basis:** Mirrored Council rule text states referenced provisions are incorporated only as referenced by the adopted IBC, IFC, IMC, and IFGC. This is not treated as proof of standalone statewide IRC, IPC, or IECC adoption.
- **Source IDs:** src:usa-wy:rules:council-ch1-lii
- **Record status:** unresolved_official_text_needed

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

For the 2024 International Codes, the official rules search result identifies the Council rule as effective 2024-06-26. The State Fire Marshal plan-review page states that all new plan review submissions from 2024-06-28 must comply with the adopted state codes and referenced standards. For electrical work, the State Fire Marshal page states that the 2023 NEC is effective 2023-07-01. For delegated local jurisdictions, W.S. 35-9-121 creates a six-month update rule after Department adoption of new standards.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-wy:001 | 2024 IFC, IBC, IFGC, IMC, IEBC | rule_effective_date | 2024-06-26 | Council rule amended/effective date in official rules search result | unresolved | src:usa-wy:rules:council-ch1-2024 | partially_verified |
| date-rule:usa-wy:002 | New plan review submissions | operative_and_mandatory_date | 2024-06-28 | New plan review submission to DFPES from 2024-06-28 | no for new submissions within DFPES plan-review scope | src:usa-wy:dfpes:plan-review | verified |
| date-rule:usa-wy:003 | 2023 NEC | effective_date | 2023-07-01 | Electrical code effective date stated by State Fire Marshal | no statewide prior-code allowance verified | src:usa-wy:dfpes:plan-review; src:usa-wy:rules:electrical-board-2023 | partially_verified |
| date-rule:usa-wy:004 | Delegated local enforcement jurisdictions | local_update_window | 6 months after Department adoption of new standards | Municipality or county with delegated authority must adopt/maintain standards at least meeting state standards; otherwise authority reverts | local standards must be at least statewide minimum | src:usa-wy:statutes:title-35-ch9 | verified |
| date-rule:usa-wy:005 | State Fire Marshal plan review | review_deadline | 21 working days for initial complete plan review; 10 working days for corrected plans | State Fire Marshal receipt of plans/corrections for covered projects | not applicable | src:usa-wy:statutes:title-35-ch9; src:usa-wy:dfpes:plan-review | partially_verified |
| date-rule:usa-wy:006 | DFPES permit to construct | work_start_rule | before work begins | Plans must be submitted before work, fees paid before permit issuance, and no work before permit to construct is issued | no | src:usa-wy:dfpes:plan-review; src:usa-wy:statutes:title-35-ch9 | partially_verified |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Fire / Building / Mechanical / Existing / Fuel Gas | none captured | null | null | null | null | null | monitor | src:usa-wy:dfpes:plan-review; src:usa-wy:rules:council-ch1-2024 | Current sources show 2024 I-Codes already in effect; no later pending code action captured. |
| Electrical | none captured | null | null | null | null | null | monitor | src:usa-wy:dfpes:plan-review; src:usa-wy:rules:electrical-board-2023 | Current source shows 2023 NEC effective; no later NEC update captured. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-wy:001 | Building / Fire / Mechanical / Fuel Gas / Existing / Electrical | State Fire Marshal plan review | Covered project before construction/remodeling begins | Plans must be submitted for certain state/local-government buildings, public buildings over 5,000 square feet, multistory public buildings, child-care centers over 10 children, certain public assembly/restaurant/bar occupancies, and aboveground fuel dispensing facilities; local jurisdictions with fire-safety enforcement authority are excepted for applicable scope. | src:usa-wy:statutes:title-35-ch9; src:usa-wy:dfpes:plan-review | partially_verified |
| applicability-rule:usa-wy:002 | Electrical | Electrical installations | Installation in/on buildings, mobile homes, and premises | Electrical installations are subject to applicable NEC minimum requirements; NEC controls over conflicting provisions in IFC, IBC, IMC, IEBC, and IFGC. | src:usa-wy:statutes:title-35-ch9 | verified |
| applicability-rule:usa-wy:003 | Fuel Gas / LP gas | Liquefied petroleum gas installations | Conflict between IFGC and LP gas standards | Current NFPA 58 and ANSI Z223.1/NFPA 54 control over conflicting IFGC provisions for LP gas installations. | src:usa-wy:statutes:title-35-ch9 | verified |
| applicability-rule:usa-wy:004 | Health care facilities | State-licensed health care facility construction/remodeling | Licensed health care facility project | Department of Health has jurisdiction except electrical installation; NFPA 101 or federal fire/life-safety certification requirements can prevail over conflicting state/local code. | src:usa-wy:statutes:title-35-ch9; src:usa-wy:health:chapter-20-jurisdiction | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** incorporation_by_reference_with_state_rules_and_statutory_conflict_rules

**Where amendments are published:** Wyoming Secretary of State administrative rules for the Department of Fire Prevention and Electrical Safety, Council on Fire Prevention and Electrical Safety in Buildings, and Electrical Board; DFPES website links to the rules and adopted-code list.

**Amendment parsing status:** partial_authority_and_high_impact_rules_only

### 5.2 State Amendment Sources

| Amendment Source ID | Description | Source IDs | Parsing Status | Notes |
| --- | --- | --- | --- | --- |
| amendment-source:usa-wy:council-ch1 | Council on Fire Prevention and Electrical Safety in Buildings, Chapter 1, General Provisions / Codes and Standards | src:usa-wy:rules:council-ch1-2024; src:usa-wy:rules:council-ch1-lii | partial | Official rules search identified the rule and effective date; Cornell mirror text was used as supplemental text for incorporated-code details. |
| amendment-source:usa-wy:dfpes-rules | Rules of the Department of Fire Prevention and Electrical Safety | src:usa-wy:dfpes:plan-review | not_parsed | DFPES links to rules; full rule text and fee/plan-review chapters need separate extraction. |
| amendment-source:usa-wy:electrical-board | Electrical Board Rules and Regulations | src:usa-wy:rules:electrical-board-2023; src:usa-wy:dfpes:plan-review | not_parsed | Effective 2023 NEC verified from DFPES page; full Electrical Board amendments not parsed. |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Topic | Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| amendment:usa-wy:001 | Electrical | NEC conflict priority | If IFC, IBC, IMC, IEBC, or IFGC conflicts with the NEC for electrical installation standards, the NEC controls. | src:usa-wy:statutes:title-35-ch9 | verified |
| amendment:usa-wy:002 | Fuel Gas | LP gas standards conflict priority | Current NFPA 58 and ANSI Z223.1/NFPA 54 control over conflicting IFGC requirements for liquefied petroleum gas installations. | src:usa-wy:statutes:title-35-ch9 | verified |
| amendment:usa-wy:003 | Fire | IFC appendices | Mirrored Council rule text lists IFC Appendix D, E, F, and G as included. Official rule PDF should be directly extracted before production use. | src:usa-wy:rules:council-ch1-lii; src:usa-wy:rules:council-ch1-2024 | partially_verified |
| amendment:usa-wy:004 | Residential / Plumbing / Energy | Referenced provisions only | Mirrored Council rule text includes provisions of IRC, IPMC, IPC, and IECC only as referenced in chapters 1 through 10 of adopted codes; this is not treated as standalone adoption. | src:usa-wy:rules:council-ch1-lii | unresolved_official_text_needed |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-wy"
  model: "state_default_optional_local_delegation"
  enforcing_entities:
    - "Department of Fire Prevention and Electrical Safety / State Fire Marshal, where state jurisdiction applies or local authority has not been delegated"
    - "Chief Electrical Inspector / Electrical Safety Division, for electrical installations unless local electrical enforcement authority has been delegated"
    - "Municipalities and counties that apply for and receive delegated enforcement authority"
    - "Department of Health / Healthcare Licensing and Surveys for state-licensed health care facility construction/remodeling except electrical installation"
  required_officials:
    - "Certified fire inspector or building inspector for delegated fire/building enforcement"
    - "Certified and state-licensed electrical inspector for delegated electrical enforcement"
    - "Certified plan reviewer/inspectors for certain state-owned, school, or health care facility delegations"
  state_reserved_activities:
    - "State-owned or leased buildings remain under State Fire Marshal/chief electrical inspector authority except where local governmental entity assumes authority under W.S. 35-9-121(b)"
    - "Health care facility fire and life-safety inspections required for federal certification remain with the Department of Health after construction/remodeling"
    - "Electrical installation remains outside Department of Health health care facility construction jurisdiction"
  source_ids:
    - "src:usa-wy:statutes:title-35-ch9"
    - "src:usa-wy:health:chapter-20-jurisdiction"
    - "src:usa-wy:health:hls-construction"
  verification_status: "partially_verified"
  confidence: 0.78
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-wy"
  model: "equivalent_or_more_stringent_minimums_for_delegated_jurisdictions"
  applies_to_code_families:
    - "fire"
    - "building"
    - "existing building"
    - "electrical"
    - "state-owned/school buildings when local entity assumes authority under W.S. 35-9-121(b)"
  approval_required: true
  approving_authority_id: "ahj:usa-wy:dfpes-state-fire-marshal"
  filing_required: "unresolved"
  registry_exists: "unresolved"
  registry_source_ids: []
  legal_basis_source_ids:
    - "src:usa-wy:statutes:title-35-ch9"
  verification_status: "partially_verified"
  confidence: 0.70
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement authority and local amendment/local standard authority are linked but not identical. A municipality or county must receive delegated authority to enforce and interpret fire, building, existing-building, or electrical safety standards. Before initial delegation, the State Fire Marshal must determine that the local governing body has adopted minimum standards by ordinance or resolution that are equivalent to or more stringent than applicable Department standards. After Department code updates, a delegated municipality or county has six months to adopt or maintain standards at least meeting statewide standards; otherwise enforcement authority reverts to the Department.

### 6.4 Known Local Amendment Registries

No statewide local amendment or delegation registry was captured. The report should not assume that every municipality or county has local enforcement authority.

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendments were parsed in this pass. Local jurisdictions with delegated authority should be researched individually before jurisdiction-specific compliance advice is generated.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: authority_model_started_boundary_and_registry_data_missing

Jurisdiction stack:

```text
Address
  -> State of Wyoming
  -> County
  -> Municipality / unincorporated county
  -> Determine whether municipality or county has DFPES-delegated enforcement authority for fire, building, existing-building, and/or electrical standards
  -> If no delegated authority or for reserved state scope, DFPES / State Fire Marshal or Chief Electrical Inspector
  -> If state-licensed health care facility, Department of Health HLS for construction/remodeling except electrical installation
  -> Fire AHJ / building AHJ / electrical AHJ based on delegated authority and project type
  -> Applicable state code adoption records
  -> Applicable local ordinance/amendment records, if any
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | not selected | none | statewide | unknown | pending |
| County | not selected | none | statewide | unknown | pending |
| Municipality | not selected | none | statewide | unknown | pending |
| Fire District | not selected | none | statewide | unknown | pending |
| Special District | not selected | none | statewide | unknown | pending |
| Delegated-code-enforcement jurisdictions | not captured | none | statewide if registry exists | unknown | unresolved |
| Health care facility delegated jurisdictions | Department of Health list referenced but not extracted | src:usa-wy:health:hls-construction | partial | unknown | unresolved |

### 7.3 AHJ Contact Data

No AHJ contact data was populated. DFPES and Department of Health pages provide program-level contact points, but a production AHJ resolver needs jurisdiction-specific delegated-authority records, inspector/service areas, and local ordinance links.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Publisher / Authority | Source Type | URL / Locator | Currentness / Effective Date | Supports | Caveats |
| --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-wy:statutes:title-35-ch9 | Wyoming Statutes, Title 35, Chapter 9 - Fire Protection | Wyoming Legislature | statute PDF | https://wyoleg.gov/statutes/compress/title35.pdf | Legislature PDF search result published last month; accessed 2026-06-26 | DFPES creation, Council/Board creation, State Fire Marshal duties, plan review, electrical standards, local enforcement, health care facility jurisdiction | Official compressed-title PDF; line extraction from PDF was readable, but production should pin exact statutory sections in a statute database if available. |
| src:usa-wy:statutes:title-15-municipal-powers | Wyoming Statutes, Title 15 - Cities and Towns | Wyoming Legislature | statute PDF | https://wyoleg.gov/statutes/compress/title15.pdf | Legislature PDF search result published last month; accessed 2026-06-26 | Municipal fire/building-related police powers background | Official compressed-title PDF; used as background only, not primary local-delegation authority. |
| src:usa-wy:statutes:title-27-elevator-safety-device | Wyoming Statutes, Title 27 - Labor and Employment, W.S. 27-1-103 safety devices on elevators and machinery | Wyoming Legislature | statute PDF | https://wyoleg.gov/statutes/compress/title27.pdf | Legislature PDF search result published last month; accessed 2026-06-26 | Elevator/conveyance open issue marker | Search result indicates safety-device provisions only; not sufficient to classify elevator code authority. |
| src:usa-wy:dfpes:plan-review | Plan Review | Wyoming State Fire Marshal / Department of Fire Prevention and Electrical Safety | agency web page | https://wsfm.wyo.gov/fire-prevention/plan-review | Accessed 2026-06-26; page states 2024 International Codes adopted on 2024-06-28 and new submissions from 2024-06-28 must comply | Adopted code list, plan review process, operative compliance trigger, 2023 NEC effective date | Agency website page; page footer did not expose a reliable last-updated date in parsed text. |
| src:usa-wy:dfpes:rules-and-statutes | Rules and Statutes | Wyoming State Fire Marshal / Department of Fire Prevention and Electrical Safety | agency web page | https://wsfm.wyo.gov/electrical-safety/rules-statutes-and-jurisdictions | Accessed 2026-06-26 | Adopted code list and links to statute/rule documents | Agency page duplicates adopted-code text from plan-review page; linked rule PDFs hosted through Google Drive were not fetched directly. |
| src:usa-wy:rules:council-ch1-2024 | Fire Prevention & Electrical Safety, Dept. of - Council on Fire Prevention, Electrical Safety, Chapter 1 | Wyoming Secretary of State Administrative Rules | official rule PDF/search result | rules.wyo.gov source_id=24233, filename=24233.pdf | Effective 2024-06-26 to Current | 2024 I-Code adoption rule, effective date, Council code standards | Official PDF search result was found, but the web fetch returned an internal fetch error; code-list details were cross-checked using the agency page and Cornell mirror. |
| src:usa-wy:rules:electrical-board-2023 | Fire Prevention & Electrical Safety, Dept. of - Electrical Board rules | Wyoming Secretary of State Administrative Rules | official rule PDF/search result | rules.wyo.gov source_id=21789, filename=21789.pdf | Search result date 2023-07-19; accessed 2026-06-26 | Electrical Board/NEC context | Official PDF search result located; full rule text not parsed. The 2023 NEC effective date is taken from the State Fire Marshal page. |
| src:usa-wy:health:hls-construction | Facility Construction | Wyoming Department of Health, Healthcare Licensing and Surveys | agency web page | https://health.wyo.gov/aging/hls/healthcare-facility-construction/ | Accessed 2026-06-26 | HLS construction plan review, final licensure surveys, electrical review coordination, delegated health-care construction jurisdiction list reference | Agency page; linked jurisdiction delegation list was not extracted. |
| src:usa-wy:health:chapter-20-jurisdiction | Chapter 20: Health Care Facilities Jurisdiction and Delegation | Wyoming Department of Health / Healthcare Licensing and Surveys | administrative rule PDF | https://health.wyo.gov/wp-content/uploads/2016/11/HLS-Rule-Ch-20-Jurisdiction-and-Delegation.pdf | Current rule per PDF metadata generated 2016-11-07; effective 2004-07-15 to Current | Department of Health AHJ definition, health-care construction jurisdiction/delegation, NFPA 101/federal life-safety certification priority | Official agency PDF; production should confirm against current Secretary of State rules database if the PDF copy is stale. |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-wy:rules:council-ch1-2024 | fetch_gap | Official Secretary of State PDF result was located, including effective date and snippets, but the PDF could not be directly opened by the browsing tool. | Use as official locator/effective-date source; extract PDF separately before verified status. |
| src:usa-wy:rules:electrical-board-2023 | fetch_gap | Official Secretary of State Electrical Board rule result was located, but full text was not extracted. | Use as official locator only; rely on DFPES page for 2023 NEC effective date until rule text is extracted. |
| src:usa-wy:dfpes:plan-review | agency_page | Agency page states code adoptions and transition trigger; it is not a codified rule text. | Use for current agency-facing adoption and submission requirements; confirm with rules for verified status. |
| src:usa-wy:health:chapter-20-jurisdiction | agency_pdf | PDF appears to be an agency-hosted generated rule copy and may lag the SOS database. | Use for partial verification; confirm against SOS before verified status. |

### 8.3 Supplemental Sources

| Source ID | Title | Publisher | Source Type | URL | Supports | Caveats |
| --- | --- | --- | --- | --- | --- | --- |
| src:usa-wy:rules:council-ch1-lii | 041-1 Wyo. Code R. §§ 1-2 - Codes and Standards | Cornell Legal Information Institute | unofficial regulation mirror | https://www.law.cornell.edu/regulations/wyoming/041-1-Wyo-Code-R-SSSS-1-2 | Text of incorporated 2024 I-Codes, referenced IRC/IPMC/IPC/IECC provisions, IFC appendices, amended effective date | Supplemental only; not an official Wyoming source. Use to guide extraction, not as final legal authority. |

### 8.4 Source Extraction Metadata

| Source ID | Extracted Fields | Extraction Method | Extracted By | Extraction Date | Confidence |
| --- | --- | --- | --- | --- | --- |
| src:usa-wy:statutes:title-35-ch9 | authority creation, Council powers, State Fire Marshal duties, plan review, electrical, local enforcement, health care facilities | browser PDF text extraction plus targeted screenshot QA | ChatGPT | 2026-06-26 | 0.78 |
| src:usa-wy:dfpes:plan-review | adopted codes, plan review process, compliance trigger, NEC effective date | browser HTML extraction | ChatGPT | 2026-06-26 | 0.82 |
| src:usa-wy:rules:council-ch1-2024 | official rule locator and effective date | search result extraction | ChatGPT | 2026-06-26 | 0.60 |
| src:usa-wy:rules:council-ch1-lii | incorporated-code list, IFC appendices, referenced-code provision treatment | browser HTML extraction from supplemental mirror | ChatGPT | 2026-06-26 | 0.55 |
| src:usa-wy:health:hls-construction | HLS construction review, delegated jurisdictions reference, electrical review coordination | browser HTML extraction | ChatGPT | 2026-06-26 | 0.70 |
| src:usa-wy:health:chapter-20-jurisdiction | Department of Health AHJ and health-care construction jurisdiction | search result/PDF text extraction | ChatGPT | 2026-06-26 | 0.66 |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 1.00 | none | Core authority and current code fields are populated; unresolved code families and local registries remain explicit. |
| report | risk.overall_confidence | 0.58 | verified | 1.00 | none | Confidence reflects a source-backed but incomplete report. |
| ahj:usa-wy:council-fire-prevention-electrical-safety | authority.role | adopts minimum fire/building-related standards | partially_verified | 0.78 | src:usa-wy:statutes:title-35-ch9; src:usa-wy:rules:council-ch1-2024 | Statute and official rules search support authority; official rule PDF needs direct extraction. |
| ahj:usa-wy:dfpes-state-fire-marshal | authority.role | administers and enforces state program | partially_verified | 0.80 | src:usa-wy:statutes:title-35-ch9; src:usa-wy:dfpes:plan-review | Strong statutory support. |
| code-adoption:usa-wy:ibc-2024 | edition | 2024 IBC | partially_verified | 0.82 | src:usa-wy:rules:council-ch1-2024; src:usa-wy:dfpes:plan-review; src:usa-wy:rules:council-ch1-lii | Official agency page and rules locator align. |
| code-adoption:usa-wy:ifc-2024 | edition | 2024 IFC | partially_verified | 0.82 | src:usa-wy:rules:council-ch1-2024; src:usa-wy:dfpes:plan-review; src:usa-wy:rules:council-ch1-lii | Official agency page and rules locator align. |
| code-adoption:usa-wy:nec-2023 | effective_date | 2023-07-01 | partially_verified | 0.76 | src:usa-wy:dfpes:plan-review; src:usa-wy:rules:electrical-board-2023 | Agency page supports effective date; formal rule text not parsed. |
| local-enforcement:usa-wy | model | state_default_optional_local_delegation | verified | 0.80 | src:usa-wy:statutes:title-35-ch9 | Statute directly describes delegation process and reversion. |
| local-amendment-rule:usa-wy | model | equivalent_or_more_stringent_minimums_for_delegated_jurisdictions | verified | 0.70 | src:usa-wy:statutes:title-35-ch9 | Registry/filing mechanics remain unresolved. |
| ahj:usa-wy:department-of-health-hls | jurisdiction | health care facility construction/remodeling except electrical installation | partially_verified | 0.75 | src:usa-wy:statutes:title-35-ch9; src:usa-wy:health:chapter-20-jurisdiction | Statute and HLS rule support scope. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Every `src:usa-wy:` identifier used in the body appears in section 8. |
| All authority IDs resolve | pass | All non-unresolved authority IDs are defined in section 2. |
| All current code families have adoption records | partial | Core verified statewide code families have records; unresolved standalone IRC/IPC/IECC/accessibility/elevator rows remain explicit. |
| Building and operational fire code are separated | pass | Fire construction references and operational/prevention fire code are separate rows. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Dates are split; NEC adoption date is left null because only effective date was verified. |
| Effective dates are valid ISO dates | pass | Entered dates use ISO format. |
| No impossible date sequences | pass | 2024 rule effective date precedes 2024-06-28 plan-review trigger; NEC dates are consistent. |
| Transition rules have explicit trigger conditions | pass | Plan-submission and local-update triggers are stated. |
| Permit-date logic is captured where applicable | partial | State plan-review submission/permit-to-construct logic is captured; local permit concurrency rules are not captured. |
| Local enforcement model classified | pass | Classified as state default with optional local delegation. |
| Local amendment rule classified | partial | Equivalent-or-more-stringent rule captured; filing/registry details unresolved. |
| AHJ confirmation metadata present | fail | No jurisdiction-specific AHJ contacts, delegation list, or local ordinance records were populated. |
| Official-source caveats captured | pass | Official-source limitations are recorded in section 8.2. |
| Leftover template markers absent | pass | No template placeholder markers intentionally remain. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-wy:001 | high | official rule PDF extraction | Direct extraction of the Secretary of State Council rule PDF was blocked by browsing-tool fetch error. | Download or manually extract rules.wyo.gov source_id=24233 and compare against the Cornell mirror. | null | null | open |
| issue:usa-wy:002 | high | residential / plumbing / energy | Standalone statewide IRC, IPC, and IECC adoption was not confirmed; only referenced provisions were identified in supplemental mirrored text. | Confirm official rule text and any separate agency authority for residential, plumbing, and energy codes. | null | null | open |
| issue:usa-wy:003 | high | local delegation registry | No statewide list of municipalities/counties with delegated authority for fire, building, existing-building, or electrical enforcement was captured. | Obtain DFPES delegation records and local ordinance links. | null | null | open |
| issue:usa-wy:004 | medium | state amendments | Full Department, Council, and Electrical Board amendment/administrative chapters were not parsed. | Extract and normalize all current DFPES, Council, and Electrical Board rules. | null | null | open |
| issue:usa-wy:005 | medium | health care facility delegation | HLS page references a list of jurisdictions delegated plan-review/inspection responsibilities; the linked list was not extracted. | Fetch and parse Form 116 / jurisdiction delegation list. | null | null | open |
| issue:usa-wy:006 | medium | elevator / conveyance | Elevator/conveyance code authority remains unresolved beyond general Title 27 safety-device statute result. | Research current Wyoming elevator/conveyance inspection statutes/rules and responsible agency. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-wy:dfpes-plan-review | src:usa-wy:dfpes:plan-review | html_diff | monthly | adopted code list, plan-review trigger date, process, or fee references change | 2026-06-26 | active |
| watch:usa-wy:dfpes-rules-and-statutes | src:usa-wy:dfpes:rules-and-statutes | html_diff | monthly | DFPES rule/statute links or adopted code list change | 2026-06-26 | active |
| watch:usa-wy:rules-council | src:usa-wy:rules:council-ch1-2024 | rules_diff | monthly | Council rules amended, superseded, or proposed | 2026-06-26 | active |
| watch:usa-wy:rules-electrical-board | src:usa-wy:rules:electrical-board-2023 | rules_diff | monthly | Electrical Board rules amended, superseded, or proposed | 2026-06-26 | active |
| watch:usa-wy:statutes-title35 | src:usa-wy:statutes:title-35-ch9 | statute_diff | quarterly | W.S. 35-9-101 through 35-9-130 amended | 2026-06-26 | active |
| watch:usa-wy:health-hls | src:usa-wy:health:hls-construction | html_diff | quarterly | HLS construction rules, delegation list, or plan review process changes | 2026-06-26 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-24 | Generated baseline draft report | report:usa-wy | none | Codex | Original uploaded draft contained placeholder/source-gap content. |
| 2026-06-26 | Populated source-backed Wyoming authority, core adoption, local enforcement, and QA sections | report:usa-wy; ahj:usa-wy:dfpes-state-fire-marshal; ahj:usa-wy:council-fire-prevention-electrical-safety; code-adoption:usa-wy:ibc-2024; code-adoption:usa-wy:ifc-2024; code-adoption:usa-wy:ifgc-2024; code-adoption:usa-wy:imc-2024; code-adoption:usa-wy:iebc-2024; code-adoption:usa-wy:nec-2023; local-enforcement:usa-wy; local-amendment-rule:usa-wy | src:usa-wy:statutes:title-35-ch9; src:usa-wy:dfpes:plan-review; src:usa-wy:rules:council-ch1-2024; src:usa-wy:rules:council-ch1-lii; src:usa-wy:health:hls-construction; src:usa-wy:health:chapter-20-jurisdiction | ChatGPT | Status upgraded to partially_verified for core fields; unresolved fields remain explicit. |
