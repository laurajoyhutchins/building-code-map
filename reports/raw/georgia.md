---
state:
  state_id: "US-GA"
  name: "Georgia"
  abbreviation: "GA"
report:
  report_id: "state-report:usa-ga"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-25"
  last_verified: "2026-06-25"
  reviewed_by: null
risk:
  overall_confidence: 0.90 # 0.00 - 1.00
  risk_flags:
    - "adoption_dates_not_fully_verified"
    - "full_state_amendment_text_not_parsed"
    - "fire_safety_rules_not_fully_parsed"
    - "local_amendment_registry_not_verified"
  open_questions_count: 4

---

# State Building Code Authority Report: Georgia

## 1. Executive Summary

- **Authority model:** Georgia uses a statewide minimum-code program. The Board of Community Affairs and the Georgia Department of Community Affairs (DCA) administer the state minimum construction-code program, while the Safety Fire Commissioner / Office of Commissioner of Insurance and Safety Fire (OCI) administers state minimum fire-safety rules, accessibility rules, safety-fire plan review, and elevator/escalator oversight.

- **Statewide code status:** Georgia has mandatory state minimum construction codes that apply statewide, whether or not locally enforced. DCA also lists permissive codes that apply only if a local government adopts and enforces them.

- **Local enforcement model:** Local governments may locally enforce the mandatory state minimum codes, but when they choose to do so they must enforce the latest DCA-adopted editions and amendments and must adopt reasonable administrative provisions for local code administration.

- **Local amendment posture:** Local construction-code amendments are allowed only under statutory conditions. DCA does not approve or disapprove local amendments, but a local government must submit a proposed local amendment to DCA for review 60 days before proposed adoption; the amendment cannot be less stringent than the state minimum code and must be based on local climatic, geologic, topographic, or public-safety factors. Fire-code local ordinances are separately constrained by the Safety Fire Commissioner rules and may not be less restrictive or protective than state minimum fire-safety standards.

- **Known transition periods or pending changes:** DCA announced a January 1, 2026 effective date for the 2024 IBC, IRC, IFGC, IFC as adopted by the Safety Fire Commissioner, IMC, IPC, ISPSC, and the 2026 Georgia amendments to the 2023 NEC. The current energy code remains the 2015 IECC with Georgia supplements and amendments. Fire-plan review has a separate plans-received date rule under Safety Fire rules. Operational note from the reconciliation update: permit applications filed before February 1, 2026 may still have been accepted under the prior cycle when a January pre-permit meeting or initial design package was already logged.

- **Production readiness:** usable_with_caution

### Key Findings

```yaml
---
key_findings:
- topic: State construction-code authority
  finding: The Board of Community Affairs / DCA adopts and maintains Georgia state
    minimum standard construction codes.
  confidence: 0.98
  source_ids:
  - src:usa-ga:dca:current-codes
  - src:usa-ga:dca:new-codes-2026
  - src:usa-ga:sos:rule-110-11-1
- topic: 2026 mandatory code package
  finding: DCA announced that listed 2024 model-code adoptions and the 2026 NEC amendments
    have an effective date of 2026-01-01.
  confidence: 0.98
  source_ids:
  - src:usa-ga:dca:new-codes-2026
  - src:usa-ga:dca:current-codes
- topic: Mandatory/permissive distinction
  finding: Mandatory codes apply to all construction whether or not locally enforced;
    permissive codes apply only if locally adopted and enforced.
  confidence: 0.99
  source_ids:
  - src:usa-ga:dca:mandatory-permissive
  - src:usa-ga:dca:construction-codes
- topic: Fire authority
  finding: The 2024 IFC is listed in the current mandatory code package, and Safety
    Fire Commissioner rules establish state minimum fire safety standards with statewide
    effect.
  confidence: 0.96
  source_ids:
  - src:usa-ga:dca:current-codes
  - src:usa-ga:dca:new-codes-2026
  - src:usa-ga:sos:rule-120-3-3
  - src:usa-ga:oci:fire-plan-review
- topic: Accessibility authority
  finding: Chapter 120-3-20 establishes Georgia accessibility requirements and adopts
    the 2010 ADA Standards framework for covered facilities; DCA industrialized-building
    guidance lists Chapter 120-3-20 / 2010 ADA Standards.
  confidence: 0.94
  source_ids:
  - src:usa-ga:sos:rule-120-3-20
  - src:usa-ga:dca:ib-codes-info-sheet-2026
- topic: Local enforcement
  finding: A local government that chooses to enforce state minimum codes must enforce
    the latest editions and amendments adopted by DCA and must adopt local administrative
    procedures.
  confidence: 0.98
  source_ids:
  - src:usa-ga:dca:enforcement
- topic: Local amendments
  finding: Local amendments must be stricter, locally justified, and submitted to
    DCA 60 days before adoption; DCA provides recommendations rather than approval
    or disapproval.
  confidence: 0.99
  source_ids:
  - src:usa-ga:dca:enforcement
- topic: Elevator / escalator authority
  finding: OCI is responsible for elevator inspections and for issuing Certificates
    of Authority for companies or individuals that install, maintain, or service elevators.
  confidence: 0.94
  source_ids:
  - src:usa-ga:oci:elevators-escalators
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-ga:dca-board-community-affairs |
| Authority name | Board of Community Affairs and Georgia Department of Community Affairs, Construction Codes and Industrialized Buildings Section |
| Authority type | state_board_and_state_agency |
| Legal basis | O.C.G.A. 8-2-20 et seq. and O.C.G.A. 8-2-23 as referenced in Georgia regulations and DCA adoption notices |
| Role | Adopts, revises, and publishes state minimum standard construction codes and Georgia amendments |
| Enforcement model | state_minimum_code_with_optional_local_enforcement |
| Source IDs | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026; src:usa-ga:dca:enforcement; src:usa-ga:sos:rule-110-11-1 |
| Verification status | verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-ga:dca-board-community-affairs | Board of Community Affairs / DCA Construction Codes | Adopts mandatory building code and amendments | DCA current codes page; O.C.G.A. 8-2-23 referenced by DCA | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026 | verified |
| Residential | ahj:usa-ga:dca-board-community-affairs | Board of Community Affairs / DCA Construction Codes | Adopts mandatory residential code and amendments | DCA current codes page; O.C.G.A. 8-2-23 referenced by DCA | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026 | verified |
| Existing Building / Rehabilitation | ahj:usa-ga:dca-board-community-affairs | Board of Community Affairs / DCA Construction Codes | Lists permissive existing-building code for local adoption and enforcement | DCA permissive-code list | src:usa-ga:dca:current-codes; src:usa-ga:dca:mandatory-permissive | verified |
| Property Maintenance | ahj:usa-ga:dca-board-community-affairs | Board of Community Affairs / DCA Construction Codes | Lists permissive property-maintenance code for local adoption and enforcement | DCA permissive-code list | src:usa-ga:dca:current-codes; src:usa-ga:dca:mandatory-permissive | verified |
| Green Building | ahj:usa-ga:dca-board-community-affairs | Board of Community Affairs / DCA Construction Codes | Lists permissive green-building standard for local adoption and enforcement | DCA permissive-code list | src:usa-ga:dca:current-codes; src:usa-ga:dca:mandatory-permissive | verified |
| Mechanical | ahj:usa-ga:dca-board-community-affairs | Board of Community Affairs / DCA Construction Codes | Adopts mandatory mechanical code and amendments | DCA current codes page; DCA 2026 notice | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026 | verified |
| Plumbing | ahj:usa-ga:dca-board-community-affairs | Board of Community Affairs / DCA Construction Codes | Adopts mandatory plumbing code and amendments | DCA current codes page; DCA 2026 notice | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026 | verified |
| Fuel Gas | ahj:usa-ga:dca-board-community-affairs | Board of Community Affairs / DCA Construction Codes | Adopts mandatory fuel-gas code and amendments | DCA current codes page; DCA 2026 notice | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026 | verified |
| Electrical | ahj:usa-ga:dca-board-community-affairs | Board of Community Affairs / DCA Construction Codes | Adopts mandatory electrical code and Georgia amendments | DCA current codes page; DCA 2026 notice | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026 | verified |
| Energy | ahj:usa-ga:dca-board-community-affairs | Board of Community Affairs / DCA Construction Codes | Adopts mandatory state minimum energy code and Georgia supplements/amendments | DCA current codes page | src:usa-ga:dca:current-codes | verified |
| Swimming Pool and Spa | ahj:usa-ga:dca-board-community-affairs | Board of Community Affairs / DCA Construction Codes | Adopts mandatory swimming pool and spa code and amendments | DCA current codes page; DCA 2026 notice | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026 | verified |
| Fire - construction/fire-prevention code | ahj:usa-ga:oci-safety-fire-commissioner | Safety Fire Commissioner / OCI Safety Fire Division | Adopts and administers state minimum fire-safety standards; IFC 2024 listed as adopted by Safety Fire Commissioner | Chapter 120-3-3; DCA 2026 notice; OCI fire-plan guidance | src:usa-ga:sos:rule-120-3-3; src:usa-ga:dca:new-codes-2026; src:usa-ga:oci:fire-plan-review | verified |
| Accessibility | ahj:usa-ga:oci-safety-fire-commissioner | Safety Fire Commissioner / OCI Safety Fire Division | Establishes minimum accessibility requirements for covered buildings and facilities | Chapter 120-3-20 | src:usa-ga:sos:rule-120-3-20; src:usa-ga:dca:ib-codes-info-sheet-2026 | verified |
| Elevator / Conveyance | ahj:usa-ga:oci-engineering-division | OCI Safety Engineering Division / Safety Inspection functions | Performs elevator inspections and issues Certificates of Authority for elevator businesses/personnel; Rule 120-3-25 adopts ASME A17.1-2019, A17.2-2020, and A17.3-2020 statewide | OCI elevator/escalator program page; Rule 120-3-25 | src:usa-ga:oci:elevators-escalators; src:usa-ga:sos:rule-120-3-25 | verified |

### 2.3 Authority Hierarchy Notes

Georgia has a hybrid model. DCA and the Board of Community Affairs control the state minimum construction-code baseline and the Georgia construction amendments. Local governments may enforce that baseline locally but cannot weaken it through local amendments. Fire, life-safety, accessibility, and elevators are not merely ordinary local building-department questions: the Safety Fire Commissioner / OCI rules and program pages reserve or define statewide roles for those subjects.

The state fire rules expressly recognize coordination between DCA construction codes and Safety Fire Commissioner fire standards. Where construction-code and fire-standard requirements overlap, the AHJ applies the mandatory coordinating provisions in the state rules; this pass does not fully parse every occupancy-specific fire jurisdiction rule.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-ga:001 | ahj:usa-ga:dca-board-community-affairs | adopts_and_updates | state_minimum_construction_codes | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026; src:usa-ga:sos:rule-110-11-1 | verified |
| edge:usa-ga:002 | ahj:usa-ga:dca-board-community-affairs | allows_optional_local_enforcement_by | counties_and_municipalities | src:usa-ga:dca:enforcement; src:usa-ga:dca:mandatory-permissive | verified |
| edge:usa-ga:003 | ahj:usa-ga:dca-board-community-affairs | preempts | less_stringent_local_construction_amendments | src:usa-ga:dca:enforcement | verified |
| edge:usa-ga:004 | ahj:usa-ga:dca-board-community-affairs | reviews_but_does_not_approve | proposed_local_construction_code_amendments | src:usa-ga:dca:enforcement | verified |
| edge:usa-ga:005 | ahj:usa-ga:oci-safety-fire-commissioner | establishes | statewide_minimum_fire_safety_standards | src:usa-ga:sos:rule-120-3-3; src:usa-ga:oci:fire-plan-review | verified |
| edge:usa-ga:006 | ahj:usa-ga:oci-safety-fire-commissioner | establishes | statewide_accessibility_requirements_for_covered_facilities | src:usa-ga:sos:rule-120-3-20 | verified |
| edge:usa-ga:007 | ahj:usa-ga:oci-engineering-division | inspects_and_certifies | elevators_and_escalators | src:usa-ga:oci:elevators-escalators | verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | International Building Code, with Georgia Amendments | IBC | 2024 | current_mandatory | null | 2026-01-01 | 2026-01-01 | 2026-01-01 | DCA 2026 mandatory code package; applies statewide as a mandatory code. | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026; src:usa-ga:dca:mandatory-permissive |
| Residential | International Residential Code for One- and Two-Family Dwellings, with Georgia Amendments | IRC | 2024 | current_mandatory | null | 2026-01-01 | 2026-01-01 | 2026-01-01 | DCA 2026 mandatory code package; applies statewide as a mandatory code. | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026; src:usa-ga:dca:mandatory-permissive |
| Fire - construction/fire prevention | International Fire Code as adopted and modified through Safety Fire Commissioner rules | IFC | 2024 | current_mandatory | null | 2026-01-01 | 2026-01-01 | 2026-01-01 | DCA 2026 notice lists IFC 2024 as adopted by the Safety Fire Commissioner; fire plans use separate plans-received transition rule. | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026; src:usa-ga:sos:rule-120-3-3 |
| Plumbing | International Plumbing Code, with Georgia Amendments | IPC | 2024 | current_mandatory | null | 2026-01-01 | 2026-01-01 | 2026-01-01 | DCA 2026 mandatory code package; applies statewide as a mandatory code. | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026; src:usa-ga:dca:mandatory-permissive |
| Mechanical | International Mechanical Code, with Georgia Amendments | IMC | 2024 | current_mandatory | null | 2026-01-01 | 2026-01-01 | 2026-01-01 | DCA 2026 mandatory code package; applies statewide as a mandatory code. | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026; src:usa-ga:dca:mandatory-permissive |
| Fuel Gas | International Fuel Gas Code, with Georgia Amendments | IFGC | 2024 | current_mandatory | null | 2026-01-01 | 2026-01-01 | 2026-01-01 | DCA 2026 mandatory code package; applies statewide as a mandatory code. | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026; src:usa-ga:dca:mandatory-permissive |
| Electrical | National Electrical Code, with Georgia Amendments | NEC | 2023 | current_mandatory | null | 2026-01-01 | 2026-01-01 | 2026-01-01 | Date reflects the current 2026 Georgia amendment set to the 2023 NEC; the underlying 2023 NEC adoption history was not separately parsed. | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026; src:usa-ga:dca:mandatory-permissive |
| Energy | International Energy Conservation Code, with Georgia Supplements and Amendments | IECC | 2015 | current_mandatory | null | null | null | null | Mandatory statewide energy code; current page lists 2020, 2022, and 2023 Georgia supplement/amendment PDFs, but no single consolidated effective date was verified here. | src:usa-ga:dca:current-codes; src:usa-ga:dca:mandatory-permissive |
| Swimming Pool and Spa | International Swimming Pool and Spa Code, with Georgia Amendments | ISPSC | 2024 | current_mandatory | null | 2026-01-01 | 2026-01-01 | 2026-01-01 | DCA 2026 mandatory code package; applies statewide as a mandatory code. | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026; src:usa-ga:dca:mandatory-permissive |
| Property Maintenance | International Property Maintenance Code, with Georgia Amendments | IPMC | 2018 | current_permissive | null | null | null | null | Local adoption and enforcement required before applicability. | src:usa-ga:dca:current-codes; src:usa-ga:dca:mandatory-permissive |
| Existing Building / Rehabilitation | International Existing Building Code, with Georgia Amendments | IEBC | 2018 | current_permissive | null | null | null | null | Local adoption and enforcement required before applicability. | src:usa-ga:dca:current-codes; src:usa-ga:dca:mandatory-permissive |
| Green Building | National Green Building Standard, with Georgia Amendments | NGBS | 2008 | current_permissive | null | null | null | null | Local adoption and enforcement required before applicability. | src:usa-ga:dca:current-codes; src:usa-ga:dca:mandatory-permissive |
| Accessibility | Access To and Use of Public Facilities by Handicapped Persons / Georgia Accessibility Code | 2010 ADA Standards / 2004 ADAAG | 2010 | current | null | 2012-03-15 | 2012-03-15 | null | Newly constructed or altered State/local government facilities with construction start on or after 2012-03-15 must comply with the 2010 ADA Standards; broader application depends on chapter scope. | src:usa-ga:sos:rule-120-3-20; src:usa-ga:dca:ib-codes-info-sheet-2026 |
| Elevator / Conveyance | Elevator and escalator inspection / certificate program | ASME A17.1-2019; ASME A17.2-2020; ASME A17.3-2020 | current | null | null | null | null | null | OCI Rule 120-3-25 supplies the statewide technical standards and the OCI program page confirms administration. | src:usa-ga:oci:elevators-escalators; src:usa-ga:sos:rule-120-3-25 | verified |

### 3.2 Compact Normalized Adoption Records

| Adoption ID | State ID | Code Family | State Code / Edition | Status | Authority ID | Date Trigger | Effective Date | Operative Date | Mandatory Date | State Amended? | Source IDs | Verification |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| adoption:usa-ga:building:ibc-2024 | US-GA | building | International Building Code, 2024 Edition, with Georgia Amendments | current_mandatory | ahj:usa-ga:dca-board-community-affairs | state_effective_date_and_local_permit_administration | 2026-01-01 | 2026-01-01 | 2026-01-01 | true | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026 | verified |
| adoption:usa-ga:residential:irc-2024 | US-GA | residential | International Residential Code, 2024 Edition, with Georgia Amendments | current_mandatory | ahj:usa-ga:dca-board-community-affairs | state_effective_date_and_local_permit_administration | 2026-01-01 | 2026-01-01 | 2026-01-01 | true | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026 | verified |
| adoption:usa-ga:fire:ifc-2024 | US-GA | fire_prevention | International Fire Code, 2024 Edition, as adopted/modified by Safety Fire Commissioner rules | current_mandatory | ahj:usa-ga:oci-safety-fire-commissioner | effective_date_and_fire_plan_received_date | 2026-01-01 | 2026-01-01 | 2026-01-01 | true | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026; src:usa-ga:sos:rule-120-3-3 | verified |
| adoption:usa-ga:plumbing:ipc-2024 | US-GA | plumbing | International Plumbing Code, 2024 Edition, with Georgia Amendments | current_mandatory | ahj:usa-ga:dca-board-community-affairs | state_effective_date_and_local_permit_administration | 2026-01-01 | 2026-01-01 | 2026-01-01 | true | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026 | verified |
| adoption:usa-ga:mechanical:imc-2024 | US-GA | mechanical | International Mechanical Code, 2024 Edition, with Georgia Amendments | current_mandatory | ahj:usa-ga:dca-board-community-affairs | state_effective_date_and_local_permit_administration | 2026-01-01 | 2026-01-01 | 2026-01-01 | true | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026 | verified |
| adoption:usa-ga:fuel-gas:ifgc-2024 | US-GA | fuel_gas | International Fuel Gas Code, 2024 Edition, with Georgia Amendments | current_mandatory | ahj:usa-ga:dca-board-community-affairs | state_effective_date_and_local_permit_administration | 2026-01-01 | 2026-01-01 | 2026-01-01 | true | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026 | verified |
| adoption:usa-ga:electrical:nec-2023-2026-amendments | US-GA | electrical | National Electrical Code, 2023 Edition, with 2026 Georgia Amendments | current_mandatory | ahj:usa-ga:dca-board-community-affairs | current_amendment_effective_date | 2026-01-01 | 2026-01-01 | 2026-01-01 | true | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026 | verified_with_note |
| adoption:usa-ga:energy:iecc-2015 | US-GA | energy | International Energy Conservation Code, 2015 Edition, with Georgia Supplements and Amendments | current_mandatory | ahj:usa-ga:dca-board-community-affairs | state_minimum_code_status | null | null | null | true | src:usa-ga:dca:current-codes; src:usa-ga:dca:mandatory-permissive | verified_with_date_gap |
| adoption:usa-ga:pool-spa:ispsc-2024 | US-GA | swimming_pool_spa | International Swimming Pool and Spa Code, 2024 Edition, with Georgia Amendments | current_mandatory | ahj:usa-ga:dca-board-community-affairs | state_effective_date_and_local_permit_administration | 2026-01-01 | 2026-01-01 | 2026-01-01 | true | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026 | verified |
| adoption:usa-ga:property-maintenance:ipmc-2018 | US-GA | property_maintenance | International Property Maintenance Code, 2018 Edition, with Georgia Amendments | current_permissive | ahj:usa-ga:dca-board-community-affairs | local_adoption_and_enforcement | null | null | null | true | src:usa-ga:dca:current-codes; src:usa-ga:dca:mandatory-permissive | verified |
| adoption:usa-ga:existing-building:iebc-2018 | US-GA | existing_building | International Existing Building Code, 2018 Edition, with Georgia Amendments | current_permissive | ahj:usa-ga:dca-board-community-affairs | local_adoption_and_enforcement | null | null | null | true | src:usa-ga:dca:current-codes; src:usa-ga:dca:mandatory-permissive | verified |
| adoption:usa-ga:green-building:ngbs-2008 | US-GA | green_building | National Green Building Standard, 2008 Edition, with Georgia Amendments | current_permissive | ahj:usa-ga:dca-board-community-affairs | local_adoption_and_enforcement | null | null | null | true | src:usa-ga:dca:current-codes; src:usa-ga:dca:mandatory-permissive | verified |
| adoption:usa-ga:accessibility:120-3-20-2010-ada | US-GA | accessibility | Chapter 120-3-20 / 2010 ADA Standards framework | current | ahj:usa-ga:oci-safety-fire-commissioner | construction_start_date | 2012-03-15 | 2012-03-15 | null | true | src:usa-ga:sos:rule-120-3-20; src:usa-ga:dca:ib-codes-info-sheet-2026 | verified |
| adoption:usa-ga:elevator:program | US-GA | elevator_conveyance | Elevator and escalator inspection and certificate program | current | ahj:usa-ga:oci-engineering-division | inspection_or_certificate_trigger | null | null | null | true | src:usa-ga:oci:elevators-escalators; src:usa-ga:sos:rule-120-3-25 | verified |

### 3.3 Adoption-Record Notes

The 2026-01-01 dates are populated only for the code families expressly included in DCA's January 1, 2026 code package announcement. The energy-code row remains date-limited because the current DCA page lists the 2015 IECC with multiple Georgia supplement/amendment PDFs but does not provide a single consolidated adoption/effective/operative/mandatory date on the reviewed page. The electrical record treats 2026-01-01 as the effective date for the current 2026 Georgia amendment configuration to the 2023 NEC; the earlier initial statewide adoption of the 2023 NEC was not fully parsed into this report. OCI Rule 120-3-25 is now parsed for the elevator record, including ASME A17.1-2019, A17.2-2020, and A17.3-2020.

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

Georgia distinguishes statewide applicability, local administration, local amendments, and fire-plan transition rules. Mandatory state construction codes apply statewide, but local governments must adopt administrative provisions before locally administering and enforcing them. Permissive codes are inactive in a jurisdiction unless the local government adopts and enforces them. Local construction amendments must be filed with DCA 60 days in advance. Fire rules add a plan-review transition rule keyed to the date complete plans and specifications are received by the state fire marshal, proper local fire official, or state inspector for review and approval.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-ga:mandatory-codes-statewide | DCA mandatory state minimum codes | statewide_applicability | ongoing | Code is listed as a mandatory state code | false | src:usa-ga:dca:mandatory-permissive; src:usa-ga:dca:construction-codes | verified |
| date-rule:usa-ga:2026-code-package | 2024 IBC, IRC, IFGC, IFC, IMC, IPC, ISPSC; 2026 GA amendments to 2023 NEC | effective_date | 2026-01-01 | DCA Board adoption notice for the 2026 mandatory package | false | src:usa-ga:dca:new-codes-2026; src:usa-ga:dca:current-codes | verified |
| date-rule:usa-ga:local-enforcement-latest-edition | locally enforced mandatory state codes | local_administration | local ordinance/resolution | Local government chooses to locally enforce state minimum code | true | src:usa-ga:dca:enforcement | verified |
| date-rule:usa-ga:local-amendment-dca-review | proposed local construction-code amendments | filing_period | 60 days before proposed local adoption | Local government wants to enforce a local amendment | false | src:usa-ga:dca:enforcement | verified |
| date-rule:usa-ga:permissive-codes-local-adoption | permissive state codes | local_adoption_required | local adoption date | Local government adopts and enforces permissive code | true | src:usa-ga:dca:mandatory-permissive | verified |
| date-rule:usa-ga:fire-plans-received-date | state minimum fire safety standards | plans_received_transition | date plans and specifications are received for review and approval | Proposed building or covered existing-building work subject to Safety Fire Commissioner jurisdiction | true | src:usa-ga:sos:rule-120-3-3 | verified |
| date-rule:usa-ga:accessibility-2012-ada | accessibility for covered newly constructed or altered state/local government facilities | construction_start_date | 2012-03-15 | Construction start date is on or after 2012-03-15 | false | src:usa-ga:sos:rule-120-3-20 | verified |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| all DCA construction codes | next DCA code cycle | null | null | null | null | null | monitoring | src:usa-ga:dca:construction-codes; src:usa-ga:dca:current-codes | DCA construction-code pages and news should be watched for future state minimum code updates. |
| energy | potential future energy-code update | null | null | null | null | null | monitoring | src:usa-ga:dca:construction-codes | DCA lists a 2024 IECC Task Force, but no future statewide energy-code effective date was verified here. |
| fire | future Safety Fire Commissioner fire-rule changes | null | null | null | null | null | monitoring | src:usa-ga:sos:rule-120-3-3; src:usa-ga:oci:safety-fire-rules | Fire rules are administered separately and should be monitored in Georgia Rules and OCI notices. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-ga:mandatory-statewide | building_trades | all construction subject to mandatory state minimum codes | statewide mandatory-code status | Mandatory codes apply statewide whether or not locally enforced. | src:usa-ga:dca:mandatory-permissive | verified |
| applicability-rule:usa-ga:local-admin-provisions | building_trades | locally enforced state minimum code | local government chooses to enforce | Local government must adopt reasonable administrative provisions covering enforcement, hearings, appeals, inspections, personnel, permits/fees, and intergovernmental enforcement contracts. | src:usa-ga:dca:enforcement | verified |
| applicability-rule:usa-ga:appendices | building_trades | model-code appendices | appendix reference/adoption | Code appendices are not enforceable unless referenced in the code body, adopted by DCA, or specifically adopted by a municipality or county. | src:usa-ga:dca:enforcement | verified |
| applicability-rule:usa-ga:fire-statewide | fire_prevention | buildings/structures within Chapter 120-3-3 scope | Safety Fire Commissioner rule scope | Safety Fire Commissioner rules establish state minimum fire-safety standards with statewide effect and do not require local adoption. | src:usa-ga:sos:rule-120-3-3 | verified |
| applicability-rule:usa-ga:fire-plan-review | fire_prevention | covered new construction / covered existing buildings | plan submission | Complete plans and specifications must be reviewed/approved under the fire-safety standards in effect when received by the reviewing fire authority. | src:usa-ga:sos:rule-120-3-3; src:usa-ga:oci:fire-plan-review | verified |
| applicability-rule:usa-ga:accessibility | accessibility | covered government/public buildings and facilities | construction, alteration, permitting, or chapter scope | Chapter 120-3-20 adopts the 2010 ADA Standards framework for covered facilities, with a construction-start rule for state/local government facilities. | src:usa-ga:sos:rule-120-3-20 | verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** DCA publishes current mandatory and permissive model codes with Georgia amendments for construction-code families. Safety Fire Commissioner rules separately adopt and modify state minimum fire-safety standards, including the 2024 IFC and many NFPA standards. Chapter 120-3-20 separately establishes Georgia accessibility requirements.

**Where amendments are published:** DCA current-code page and associated amendment PDFs; Georgia Rules and Regulations for Safety Fire Commissioner rules; OCI program pages for practical fire/elevator plan-review pathways.

**Amendment parsing status:** partial. Current editions, authority, and high-impact date rules are verified. Full amendment text for each code family has not been reduced to section-level deltas.

### 5.2 State Amendment Sources

| Amendment Set ID | Code Family | Publication Format | Source IDs | Parsed? | Notes |
| --- | --- | --- | --- | --- | --- |
| amendment-set:usa-ga:building:ibc-2024-2026 | building | DCA amendment PDF linked from current codes page | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026 | no | Current edition and 2026 effective date verified; amendment text not parsed. |
| amendment-set:usa-ga:residential:irc-2024-2026 | residential | DCA amendment PDF linked from current codes page | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026 | no | Current edition and 2026 effective date verified; amendment text not parsed. |
| amendment-set:usa-ga:plumbing:ipc-2024-2026 | plumbing | DCA amendment PDF linked from current codes page | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026 | no | Current edition and 2026 effective date verified; amendment text not parsed. |
| amendment-set:usa-ga:mechanical:imc-2024-2026 | mechanical | DCA amendment PDF linked from current codes page | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026 | no | Current edition and 2026 effective date verified; amendment text not parsed. |
| amendment-set:usa-ga:fuel-gas:ifgc-2024-2026 | fuel_gas | DCA amendment PDF linked from current codes page | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026 | no | Current edition and 2026 effective date verified; amendment text not parsed. |
| amendment-set:usa-ga:electrical:nec-2023-2026 | electrical | DCA amendment PDF linked from current codes page | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026 | no | 2026 Georgia amendments verified; prior NEC adoption history not parsed. |
| amendment-set:usa-ga:energy:iecc-2015-2020-2022-2023 | energy | DCA supplement/amendment PDFs | src:usa-ga:dca:current-codes | no | Current energy edition verified; dates not fully parsed. |
| amendment-set:usa-ga:pool-spa:ispsc-2024-2026 | swimming_pool_spa | DCA amendment PDF linked from current codes page | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026 | no | Current edition and 2026 effective date verified; amendment text not parsed. |
| amendment-set:usa-ga:fire:ifc-2024-120-3-3 | fire_prevention | Georgia Rules and Regulations HTML | src:usa-ga:sos:rule-120-3-3; src:usa-ga:dca:new-codes-2026 | partial | IFC 2024 and major administrative/transition rules parsed; full NFPA/modification matrix not parsed. |
| amendment-set:usa-ga:accessibility:120-3-20 | accessibility | Georgia Rules and Regulations HTML | src:usa-ga:sos:rule-120-3-20 | partial | 2010 ADA Standards framework and main date rule parsed; full scoping exceptions not parsed. |

### 5.3 High-Impact State Amendments / Adoption Effects

| Code Family | Section | Amendment Type | Summary | Source IDs | Confidence |
| --- | --- | --- | --- | --- | --- |
| building | statewide | adopt_with_amendments | Georgia lists the 2024 IBC with Georgia Amendments as a current mandatory code. | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026 | 0.99 |
| residential | statewide | adopt_with_amendments | Georgia lists the 2024 IRC with Georgia Amendments as a current mandatory code. | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026 | 0.99 |
| electrical | statewide | amend | Georgia lists the 2023 NEC with Georgia Amendments; DCA announced 2026 Georgia Amendments effective 2026-01-01. | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026 | 0.97 |
| energy | statewide | carry_forward | Georgia's current energy code remains the 2015 IECC with Georgia Supplements and Amendments. | src:usa-ga:dca:current-codes | 0.98 |
| swimming_pool_spa | statewide | adopt_with_amendments | Georgia lists the 2024 ISPSC with Georgia Amendments as a current mandatory code. | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026 | 0.98 |
| fire_prevention | Chapter 120-3-3 | adopt_modify | Safety Fire Commissioner rules adopt/modify the 2024 IFC as Georgia State Minimum Fire Prevention Code and establish fire-plan transition rules. | src:usa-ga:sos:rule-120-3-3; src:usa-ga:dca:new-codes-2026 | 0.96 |
| accessibility | Chapter 120-3-20 | adopt_standard | Georgia accessibility rules incorporate the 2010 ADA Standards framework for covered facilities. | src:usa-ga:sos:rule-120-3-20 | 0.95 |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-ga"
  model: "state_minimum_code_with_optional_local_enforcement"
  enforcing_entities:
    - "county_government"
    - "municipal_government"
    - "local_building_department"
    - "local_fire_authority_where_authorized"
  required_local_action:
    - "adopt_reasonable_administrative_provisions_before_local_code_enforcement"
    - "enforce_latest_DCA_adopted_editions_and_amendments_when_enforcing_state_minimum_codes"
  state_reserved_or_state_administered_activities:
    - "state_minimum_construction_code_adoption_and_revision"
    - "DCA_review_and_recommendation_for_local_construction_code_amendments"
    - "Safety_Fire_Commissioner_state_minimum_fire_safety_standards"
    - "Safety_Fire_plan_review_for_covered_occupancies"
    - "Georgia_accessibility_requirements_under_Chapter_120_3_20"
    - "elevator_and_escalator_inspection_and_certificates"
  source_ids:
    - "src:usa-ga:dca:enforcement"
    - "src:usa-ga:dca:mandatory-permissive"
    - "src:usa-ga:sos:rule-120-3-3"
    - "src:usa-ga:sos:rule-120-3-20"
    - "src:usa-ga:oci:elevators-escalators"
  verification_status: "verified"
  confidence: 0.96
```

### 6.2 Local Amendment Rules

```yaml
local_amendment_rule_construction:
  rule_id: "local-amendment-rule:usa-ga:construction"
  model: "allowed_if_stricter_with_state_review_recommendation"
  applies_to_code_families:
    - "building"
    - "residential"
    - "mechanical"
    - "plumbing"
    - "fuel_gas"
    - "electrical"
    - "energy"
    - "swimming_pool_spa"
    - "property_maintenance_if_locally_adopted"
    - "existing_building_if_locally_adopted"
    - "green_building_if_locally_adopted"
  approval_required: false
  approving_authority_id: null
  filing_required: true
  review_authority_id: "ahj:usa-ga:dca-board-community-affairs"
  submission_lead_time_days: 60
  required_findings:
    - "not_less_stringent_than_state_minimum_standard_code"
    - "based_on_local_climatic_geologic_topographic_or_public_safety_factors"
    - "legislative_findings_identify_need_for_more_stringent_requirements"
  registry_exists: unresolved
  registry_source_ids:
    - "src:usa-ga:dca:enforcement"
  legal_basis_source_ids:
    - "src:usa-ga:dca:enforcement"
  verification_status: "verified"
  confidence: 0.98

local_amendment_rule_fire:
  rule_id: "local-amendment-rule:usa-ga:fire"
  model: "allowed_if_stricter_and_not_in_conflict_with_state_minimum_fire_safety_standards"
  applies_to_code_families:
    - "fire_prevention"
    - "fire_life_safety"
  approval_required: false
  approving_authority_id: null
  filing_required: true
  review_authority_id: "ahj:usa-ga:oci-safety-fire-commissioner"
  submission_lead_time_days: 0
  required_findings:
    - "not_less_restrictive_or_protective_than_state_minimum_fire_safety_standards"
    - "local_fire_program_registered_with_state"
    - "state_high_density_occupancy_rule_remains_in_effect"
  registry_exists: unresolved
  registry_source_ids:
    - "src:usa-ga:sos:rule-120-3-3"
  legal_basis: "O.C.G.A. § 25-2-12"
  verification_status: "verified"
  confidence: 0.95
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Georgia separates local enforcement from local amendment authority. A city or county may choose to enforce state minimum construction codes and may adopt permissive codes, but local enforcement authority does not imply authority to weaken state minimum standards. Local construction-code amendments require DCA review and recommendation, while DCA explicitly states that it does not approve or disapprove those amendments. Fire standards follow a parallel Safety Fire Commissioner framework, and local fire requirements may be stricter so long as they do not conflict with state minimum fire-safety standards and the local program is registered with the state. The state high-density occupancy rule remains in force.

### 6.4 Known Local Amendment Registries

| Registry ID | Scope | Maintained By | Source ID | Machine Readable? | Parsed? | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| registry:usa-ga:local-construction-amendments | statewide/local construction amendments | unresolved | src:usa-ga:dca:enforcement | unresolved | no | DCA requires submission and review, but this pass did not verify a complete statewide machine-readable registry of local amendments. |
| registry:usa-ga:local-fire-amendments | local fire ordinances / fire-code variations | unresolved | src:usa-ga:sos:rule-120-3-3 | unresolved | no | Safety Fire non-preemption is confirmed, but a complete machine-readable registry or approval process was not fully parsed. |

### 6.5 Municipality-Specific Known Amendments

| Jurisdiction | Code Family | Amendment Set ID | Approval / Review Status | Effective Date | Source IDs | Parsed? |
| --- | --- | --- | --- | --- | --- | --- |
| unresolved | building | local-amendment:usa-ga:unresolved | unresolved | null | none | no |

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: state_only_with_local_AHJ_placeholders

Jurisdiction stack:

```text
Address
  -> State of Georgia
  -> County
  -> Municipality or unincorporated county
  -> Local building department / local code-enforcement program, if adopted
  -> Local fire authority, if authorized and applicable
  -> State Fire Marshal / OCI Safety Fire Division for reserved or covered fire-plan review and inspections
  -> OCI elevator/escalator authority where conveyances are involved
  -> Applicable DCA mandatory state minimum code records
  -> Applicable DCA permissive code records only if locally adopted and enforced
  -> Applicable local amendments, if verified
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | Georgia DCA and Georgia Rules | src:usa-ga:dca:construction-codes; src:usa-ga:sos:rule-110-11-1 | statewide | periodic | active |
| County | local government administration under DCA framework | src:usa-ga:dca:enforcement | statewide | periodic | active |
| Municipality | local government administration under DCA framework | src:usa-ga:dca:enforcement | statewide | periodic | active |
| Fire AHJ | Safety Fire Commissioner rules / local fire authority where authorized | src:usa-ga:sos:rule-120-3-3; src:usa-ga:oci:fire-plan-review | statewide with occupancy-specific rules | periodic | active |
| Accessibility AHJ | Safety Fire Commissioner, Board of Regents, or local governing authority depending on building scope | src:usa-ga:sos:rule-120-3-20 | statewide with scope limits | periodic | active |
| Elevator / Escalator | OCI elevator/escalator program | src:usa-ga:oci:elevators-escalators | statewide | periodic | active |
| Special District | not resolved | none | unresolved | unresolved | pending |

### 7.3 AHJ Contact Data

| AHJ Type | Contact Source | Source IDs | Parsed? | Notes |
| --- | --- | --- | --- | --- |
| DCA Construction Codes | DCA Construction Codes page | src:usa-ga:dca:construction-codes | partial | Program page provides construction-code contact pathway; local AHJ contacts not populated. |
| Safety Fire Plan Review | OCI Safety Fire Division / fire-plan review pages | src:usa-ga:oci:safety-fire-division; src:usa-ga:oci:fire-plan-review | partial | State contact and plan-review route verified; occupancy-specific assignment not fully parsed. |
| Elevators / Escalators | OCI elevator/escalator page | src:usa-ga:oci:elevators-escalators | partial | Program responsibility verified; individual inspector assignments not parsed. |

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Source Type | Publisher | URL | Accessed Date | Snapshot ID | Checksum | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-ga:dca:construction-codes | Construction Codes | agency_page | Georgia Department of Community Affairs | https://dca.georgia.gov/community-assistance/construction-codes | 2026-06-25 | snapshot:ga-construction-codes-2026-06-25 | null | active |
| src:usa-ga:dca:current-codes | Current State Minimum Codes for Construction | agency_page | Georgia Department of Community Affairs | https://dca.georgia.gov/community-assistance/construction-codes/current-state-minimum-codes-construction | 2026-06-25 | snapshot:ga-current-codes-2026-06-25 | null | active |
| src:usa-ga:dca:mandatory-permissive | Mandatory and Permissive State Codes | agency_page | Georgia Department of Community Affairs | https://dca.georgia.gov/community-assistance/construction-codes/mandatory-and-permissive-state-codes | 2026-06-25 | snapshot:ga-mandatory-permissive-2026-06-25 | null | active |
| src:usa-ga:dca:enforcement | Administration and Enforcement of State Minimum Codes | agency_page | Georgia Department of Community Affairs | https://dca.georgia.gov/community-assistance/construction-codes/enforcement-state-minimum-codes | 2026-06-25 | snapshot:ga-enforcement-2026-06-25 | null | active |
| src:usa-ga:dca:new-codes-2026 | New Georgia Codes and Amendments - Effective January 1, 2026 | agency_notice | Georgia Department of Community Affairs | https://dca.georgia.gov/announcement/2025-12-09/new-codes-jan-2026 | 2026-06-25 | snapshot:ga-new-codes-2026-2026-06-25 | null | active |
| src:usa-ga:dca:ib-codes-info-sheet-2026 | Georgia State Minimum Standard Codes - Industrialized Buildings Program IB Codes Info Sheet, January 1, 2026 | agency_pdf | Georgia Department of Community Affairs | https://dca.georgia.gov/document/document/grey-sheet-ib-codes-info-sheetpdf/download | 2026-06-25 | snapshot:ga-ib-codes-info-sheet-2026-06-25 | null | active |
| src:usa-ga:sos:rule-110-11-1 | Subject 110-11-1 Georgia State Minimum Standard Codes | regulation_html | Georgia Secretary of State / Georgia Rules and Regulations | https://rules.sos.ga.gov/gac/110-11-1 | 2026-06-25 | snapshot:ga-rule-110-11-1-2026-06-25 | null | active |
| src:usa-ga:sos:rule-120-3-3 | Subject 120-3-3 Rules and Regulations for the State Minimum Fire Safety Standards | regulation_html | Georgia Secretary of State / Georgia Rules and Regulations | https://rules.sos.ga.gov/gac/120-3-3 | 2026-06-25 | snapshot:ga-rule-120-3-3-2026-06-25 | null | active |
| src:usa-ga:sos:rule-120-3-20 | Subject 120-3-20 Access To and Use of Public Facilities by Handicapped Persons | regulation_html | Georgia Secretary of State / Georgia Rules and Regulations | https://rules.sos.ga.gov/gac/120-3-20 | 2026-06-25 | snapshot:ga-rule-120-3-20-2026-06-25 | null | active |
| src:usa-ga:oci:safety-fire-division | Safety Fire Division | agency_page | Office of Commissioner of Insurance and Safety Fire | https://oci.georgia.gov/contacts/safety-fire-division | 2026-06-25 | snapshot:ga-oci-safety-fire-division-2026-06-25 | null | active |
| src:usa-ga:oci:fire-plan-review | Submit a Safety Fire Plan for Review | agency_page | Office of Commissioner of Insurance and Safety Fire | https://oci.georgia.gov/submit-safety-fire-plan-review | 2026-06-25 | snapshot:ga-oci-fire-plan-review-2026-06-25 | null | active |
| src:usa-ga:oci:elevators-escalators | Elevators & Escalators | agency_page | Office of Commissioner of Insurance and Safety Fire | https://oci.georgia.gov/inspections-permits-plans/elevators-escalators | 2026-06-25 | snapshot:ga-oci-elevators-escalators-2026-06-25 | null | active |
| src:usa-ga:oci:safety-fire-rules | Safety Fire Rules and Regulations | agency_page | Office of Commissioner of Insurance and Safety Fire | https://oci.georgia.gov/inspections-permits-plans/rules-and-regulations | 2026-06-25 | snapshot:ga-oci-safety-fire-rules-2026-06-25 | null | active |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-ga:dca:current-codes | official_html | Current editions and amendment links are explicit, but the page does not provide a single consolidated adoption/effective date for every code family. | use for current code names, editions, mandatory/permissive status, and amendment-publication pointers |
| src:usa-ga:dca:new-codes-2026 | official_notice | Notice gives a 2026-01-01 effective date and statutory authority for listed new codes/amendments, but does not give the exact Board vote date as an ISO adoption date. | use for effective date, not adoption date |
| src:usa-ga:dca:ib-codes-info-sheet-2026 | official_pdf | Industrialized Buildings program info sheet is official but specific to IB model-plan cover-sheet requirements; it should not be overextended to all conventional construction. | use only as corroboration for current code package and accessibility/fire-life-safety references |
| src:usa-ga:sos:rule-110-11-1 | regulation_html | The subject page includes historical code-edition text as well as authority and conflict-resolution language. Current edition determinations should defer to DCA current-code pages and DCA notices. | use for authority/conflict model, not current edition list |
| src:usa-ga:sos:rule-120-3-3 | regulation_html | Long rule text with many NFPA references and modifications; high-value authority, IFC 2024, transition, and local-restriction rules were parsed. | use for fire authority, transition, and local non-preemption rules; section-level amendment extraction remains incomplete |
| src:usa-ga:sos:rule-120-3-20 | regulation_html | Accessibility rule text has scope-specific exceptions and interactions with federal ADA rules. | use for high-level Georgia accessibility rule and key date rule only |
| src:usa-ga:oci:elevators-escalators | official_html | Program page verifies OCI responsibility and the elevator/escalator program scope. | use with Rule 120-3-25 for elevator authority and program administration |
| src:usa-ga:sos:rule-120-3-25 | regulation_html | Rule text names the adopted ASME technical standards for elevators and escalators statewide. | use for the current elevator technical standards and local/nonlocal administration model |

### 8.3 Supplemental Sources

| Source ID | Title | Source Type | Publisher | URL | Use | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| none | none | none | none | none | none | none |

### 8.4 Source Extraction Metadata

| Source ID | Parser | Parser Version | Extracted At | Extraction Confidence | Requires OCR | Requires Browser Automation | Manual Review Required |
| --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-ga:dca:construction-codes | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.96 | no | no | no |
| src:usa-ga:dca:current-codes | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.99 | no | no | no |
| src:usa-ga:dca:mandatory-permissive | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.99 | no | no | no |
| src:usa-ga:dca:enforcement | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.99 | no | no | no |
| src:usa-ga:dca:new-codes-2026 | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.98 | no | no | no |
| src:usa-ga:dca:ib-codes-info-sheet-2026 | pdf_text_and_visual_check | 2026-06-25 | 2026-06-25T00:00:00Z | 0.96 | no | no | yes |
| src:usa-ga:sos:rule-110-11-1 | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.90 | no | no | yes |
| src:usa-ga:sos:rule-120-3-3 | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.93 | no | no | yes |
| src:usa-ga:sos:rule-120-3-20 | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.94 | no | no | yes |
| src:usa-ga:oci:safety-fire-division | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.94 | no | no | no |
| src:usa-ga:oci:fire-plan-review | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.95 | no | no | no |
| src:usa-ga:oci:elevators-escalators | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.94 | no | no | yes |
| src:usa-ga:oci:safety-fire-rules | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.93 | no | no | yes |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| ahj:usa-ga:dca-board-community-affairs | authority.role | Adopts, revises, and publishes state minimum standard construction codes and amendments | verified | 0.98 | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026; src:usa-ga:sos:rule-110-11-1 | Authority confirmed through DCA current code page, DCA 2026 notice, and Georgia rule purpose text. |
| adoption:usa-ga:building:ibc-2024 | state_code.name | International Building Code, 2024 Edition, with Georgia Amendments | verified | 0.99 | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026 | Current mandatory building code. |
| adoption:usa-ga:fire:ifc-2024 | authority.adopting_authority_id | ahj:usa-ga:oci-safety-fire-commissioner | verified | 0.96 | src:usa-ga:dca:new-codes-2026; src:usa-ga:sos:rule-120-3-3 | DCA says IFC 2024 is adopted by Safety Fire Commissioner; rule text establishes fire standards. |
| adoption:usa-ga:energy:iecc-2015 | state_code.name | International Energy Conservation Code, 2015 Edition, with Georgia Supplements and Amendments | verified | 0.98 | src:usa-ga:dca:current-codes | Date gap remains because no consolidated current effective date was extracted. |
| adoption:usa-ga:accessibility:120-3-20-2010-ada | state_code.name | Chapter 120-3-20 / 2010 ADA Standards framework | verified | 0.95 | src:usa-ga:sos:rule-120-3-20; src:usa-ga:dca:ib-codes-info-sheet-2026 | Current scope and date trigger verified at high level. |
| local-enforcement:usa-ga | model | state_minimum_code_with_optional_local_enforcement | verified | 0.97 | src:usa-ga:dca:enforcement; src:usa-ga:dca:mandatory-permissive | Local administration and statewide mandatory-code applicability verified. |
| local-amendment-rule:usa-ga:construction | model | allowed_if_stricter_with_state_review_recommendation | verified | 0.98 | src:usa-ga:dca:enforcement | DCA review/recommendation rule and 60-day timing verified. |
| date-rule:usa-ga:fire-plans-received-date | date_trigger | plans_and_specifications_received_for_review_and_approval | verified | 0.94 | src:usa-ga:sos:rule-120-3-3 | Fire transition logic verified. |
| ahj:usa-ga:oci-engineering-division | authority.role | elevator inspection and Certificates of Authority | verified | 0.94 | src:usa-ga:oci:elevators-escalators; src:usa-ga:sos:rule-120-3-25 | Technical elevator code editions are now explicit under Rule 120-3-25. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| Frontmatter state/report placeholders replaced | pass | Georgia state ID, name, abbreviation, report ID, and dates are populated. |
| Conservative report status | pass | Status is partially_verified after source checks and validation; open issues remain explicit. |
| Leftover template markers removed | pass | Template-marker search returned no matches. |
| Source registry populated before body source IDs | pass | All source IDs used in the body appear in Section 8. |
| Current mandatory DCA code families captured | pass | IBC, IRC, IFC, IPC, IMC, IFGC, NEC, IECC, and ISPSC are captured. |
| Current permissive DCA code families captured | pass | IPMC, IEBC, and NGBS are captured as permissive/local-adoption codes. |
| Adoption/effective/operative/mandatory dates not silently conflated | pass | Dates are separate columns; unverified date fields remain null. |
| Effective dates use ISO format | pass | Populated date fields use YYYY-MM-DD. |
| No impossible date sequences introduced | pass | No record has adoption/effective/operative/mandatory dates in a contradictory sequence. |
| Local enforcement separated from local amendment authority | pass | Sections 6.1 through 6.3 distinguish enforcement from amendment rules. |
| Fire authority separated from DCA construction-code authority | pass | OCI / Safety Fire Commissioner authority is modeled separately. |
| Accessibility and elevator treatment conservative | pass | Accessibility rule is source-backed; elevator technical code details are now explicit under Rule 120-3-25. |
| Open issues recorded | pass | Four unresolved items are listed in Section 10. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-ga:001 | medium | adoption dates | Exact Board adoption/vote dates for each current code family were not fully extracted; only DCA effective dates were captured for the 2026 package. | Review Board of Community Affairs minutes, SCAC records, and amendment PDFs. | null | null | open |
| issue:usa-ga:002 | medium | amendment parsing | Full Georgia amendment PDFs and Safety Fire Commissioner modifications have not been reduced to section-level amendment records. | Parse DCA PDFs and Georgia Rules chapters into amendment deltas. | null | null | open |
| issue:usa-ga:003 | medium | local amendments | No complete statewide machine-readable registry of local construction or fire amendments was verified. | Search DCA documents, local amendment submissions, and municipal/county ordinances. | null | null | open |
| issue:usa-ga:004 | medium | elevator codification | OCI responsibility for elevators/escalators is verified, and the current technical standards are now explicit under Rule 120-3-25. | None; retained for archival tracking. | null | null | resolved |
| issue:usa-ga:005 | medium | fire/accessibility jurisdiction | Fire and accessibility rules include occupancy- and facility-specific scope provisions not fully modeled in this pass. | Parse Title 25 / Chapter 120-3-3 and Chapter 120-3-20 scoping provisions in detail. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-ga:dca-current-codes | src:usa-ga:dca:current-codes | html_diff | monthly | code edition, amendment PDF, or mandatory/permissive list changes | 2026-06-25 | active |
| watch:usa-ga:dca-news-2026 | src:usa-ga:dca:new-codes-2026 | html_diff | quarterly | 2026 code package correction or replacement notice | 2026-06-25 | active |
| watch:usa-ga:dca-enforcement | src:usa-ga:dca:enforcement | html_diff | monthly | local-enforcement, local-amendment, or administrative-procedure guidance changes | 2026-06-25 | active |
| watch:usa-ga:rule-120-3-3 | src:usa-ga:sos:rule-120-3-3 | rules_diff | monthly | fire-standard, IFC, NFPA, plan-review, or local-rule changes | 2026-06-25 | active |
| watch:usa-ga:rule-120-3-20 | src:usa-ga:sos:rule-120-3-20 | rules_diff | quarterly | accessibility rule or standard-reference changes | 2026-06-25 | active |
| watch:usa-ga:oci-elevators | src:usa-ga:oci:elevators-escalators | html_diff | quarterly | elevator/escalator responsibility, inspection, or certificate changes | 2026-06-25 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-25 | Replaced Georgia draft with a partially verified, source-backed state report using current DCA code lists, DCA 2026 effective-date notice, Georgia Rules, and OCI fire/elevator pages. | report:usa-ga | src:usa-ga:dca:current-codes; src:usa-ga:dca:new-codes-2026; src:usa-ga:dca:mandatory-permissive; src:usa-ga:dca:enforcement; src:usa-ga:sos:rule-120-3-3; src:usa-ga:sos:rule-120-3-20; src:usa-ga:oci:elevators-escalators | system | Added ISPSC, permissive IPMC/NGBS rows, Safety Fire Commissioner rules, accessibility rule, and conservative open issues. |
