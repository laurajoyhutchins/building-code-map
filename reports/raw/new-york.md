---
state:
  state_id: "US-NY"
  name: "New York"
  abbreviation: "NY"
report:
  report_id: "state-report:usa-ny"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-26"
  last_verified: null
  reviewed_by: null
risk:
  overall_confidence: 0.72 # 0.00 - 1.00
  risk_flags:
    - "nyc_uniform_code_exception"
    - "energy_code_applies_statewide_including_nyc"
    - "fossil_fuel_equipment_provisions_suspended_by_court_order"
    - "electrical_code_scope_limited_to_referenced_standards"
    - "local_amendment_registry_partial_or_stale"
    - "ahj_boundary_and_contact_data_not_populated"
  open_questions_count: 7

---

# State Building Code Authority Report: New York

## 1. Executive Summary

- **Authority model:** New York uses a statewide code-council model. The State Fire Prevention and Building Code Council, located in the Department of State, develops and maintains the New York State Uniform Fire Prevention and Building Code (Uniform Code) and the State Energy Conservation Construction Code (Energy Code). The Department of State Division of Building Standards and Codes provides code-development, administration, and enforcement support.

- **Statewide code status:** The 2025 Uniform Code and 2025 Energy Code are the current statewide editions for the fields verified in this pass. The Code Council adopted update rules on 2025-07-25; those rules became effective on 2025-12-31. The Uniform Code incorporates the 2025 New York State Residential, Building, Plumbing, Mechanical, Fuel Gas, Fire, Property Maintenance, Existing Building, and Rail Station publications. The Energy Code incorporates the 2025 Energy Conservation Construction Code of New York State and the New York State ASHRAE 90.1-2025 publication.

- **Local enforcement model:** Local administration and enforcement is the default model. Executive Law §381 charges local governments with administration and enforcement of the Uniform Code and Energy Code, subject to county and Secretary of State fallback mechanisms. State agencies are responsible for administration and enforcement for buildings, premises, and equipment in their custody.

- **Local amendment posture:** Local enforcement authority and local amendment authority are separate. For the Uniform Code, cities, towns, villages, and Nassau County may adopt higher or more restrictive construction standards only through the Executive Law §379 notice/petition and Code Council approval process. For the Energy Code, local governments and certain other local entities may adopt more stringent local energy codes under Energy Law §11-109, with filing and Code Council review requirements.

- **Known transition periods or pending changes:** The Uniform Code transition period allowed regulated parties to comply with either the 2020 or 2025 Uniform Code between publication of the Notice of Adoption and 2025-12-31. Beginning 2025-12-31, building permit applications must comply with the 2025 Uniform Code. The Energy Code update has no transition period; building permit applications submitted on or after 2025-12-31 must comply with the 2025 Energy Code. The effective date of 19 NYCRR §1240.6 and 19 NYCRR Subpart 1229-2 is suspended pending final appellate/certiorari disposition under the court stipulation and order captured in this report.

- **Production readiness:** partially_ready_for_authority_and_current_code_adoption_fields_only

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: The State Fire Prevention and Building Code Council has statutory responsibility
    for developing and maintaining the Uniform Code and Energy Code.
  confidence: 0.95
  source_ids:
  - src:usa-ny:dos-code-development
  - src:usa-ny:exec-law-374
  - src:usa-ny:exec-law-375
- topic: Primary building code edition
  finding: The 2025 Building Code of New York State is incorporated as one of the
    2025 NYS Uniform Code Books.
  confidence: 0.9
  source_ids:
  - src:usa-ny:dos-2025-notice-adoption
  - src:usa-ny:19nycrr-1219
  - src:usa-ny:19nycrr-1221
- topic: Residential code edition
  finding: The 2025 Residential Code of New York State is incorporated in the Uniform
    Code and applies to the residential buildings identified in 19 NYCRR Part 1220.
  confidence: 0.9
  source_ids:
  - src:usa-ny:dos-2025-notice-adoption
  - src:usa-ny:19nycrr-1219
  - src:usa-ny:19nycrr-1220
- topic: Energy code edition
  finding: The Energy Code update incorporates the 2025 Energy Conservation Construction
    Code of New York State and New York State ASHRAE 90.1-2025, effective 2025-12-31.
  confidence: 0.9
  source_ids:
  - src:usa-ny:dos-2025-notice-adoption
  - src:usa-ny:19nycrr-1240
- topic: Fire code authority
  finding: The operational/prevention fire code is within the statewide Uniform Code
    structure; no distinct statewide fire-code adopting authority separate from the
    Code Council was verified in this pass.
  confidence: 0.78
  source_ids:
  - src:usa-ny:dos-uniform-code
  - src:usa-ny:dos-2025-notice-adoption
  - src:usa-ny:19nycrr-1219
- topic: Local enforcement
  finding: Local governments administer and enforce the Uniform Code and Energy Code,
    with county and Secretary of State fallback mechanisms.
  confidence: 0.9
  source_ids:
  - src:usa-ny:exec-law-381
  - src:usa-ny:dos-enforcement-programs
- topic: Local amendments
  finding: "Uniform Code local standards must be higher or more restrictive and require\
    \ Code Council approval; local energy codes must be more stringent and are subject\
    \ to filing/review under Energy Law \xA711-109."
  confidence: 0.86
  source_ids:
  - src:usa-ny:exec-law-379
  - src:usa-ny:energy-law-11-109
  - src:usa-ny:dos-code-development
  - src:usa-ny:dos-lg07-local-authority
- topic: NYC treatment
  finding: The Uniform Code applies statewide except New York City; the Energy Code
    applies statewide, including New York City.
  confidence: 0.86
  source_ids:
  - src:usa-ny:dos-code-council-page
  - src:usa-ny:dos-uniform-code
  - src:usa-ny:dos-faq
- topic: Suspended provisions
  finding: "Fossil-fuel-equipment/building-system provisions in 19 NYCRR \xA71240.6\
    \ and Subpart 1229-2 are suspended under a court stipulation and order."
  confidence: 0.88
  source_ids:
  - src:usa-ny:dos-2025-notice-adoption
  - src:usa-ny:mulhern-stipulation-order
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-ny:state-fire-prevention-building-code-council |
| Authority name | State Fire Prevention and Building Code Council |
| Authority type | state_code_council |
| Legal basis | Executive Law Article 18, including §§374-375; Energy Law Article 11 for the Energy Code framework; Title 19 NYCRR implementing rules |
| Role | Develops, maintains, amends, and updates the Uniform Code and Energy Code; approves qualifying local higher/more restrictive Uniform Code standards and qualifying more stringent local energy codes |
| Enforcement model | statewide_code_with_local_administration_and_enforcement; NYC excluded from Uniform Code but included for Energy Code |
| Source IDs | src:usa-ny:dos-code-development; src:usa-ny:exec-law-374; src:usa-ny:exec-law-375; src:usa-ny:dos-code-council-page |
| Verification status | partially_verified |

### 2.2 Related State Authorities and Administrative Units

| Authority ID | Authority Name | Authority Type | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| ahj:usa-ny:secretary-of-state | Secretary of State | state_officer | Chairs the Code Council; promulgates minimum standards for administration and enforcement; may assume or direct enforcement when local/county enforcement fails or is declined | Executive Law §§374 and 381 | src:usa-ny:exec-law-374; src:usa-ny:exec-law-381 | partially_verified |
| ahj:usa-ny:dos-dbsc | Department of State Division of Building Standards and Codes | state_agency_division | Provides services for development, administration, and enforcement of the Uniform Code and Energy Code; supports local AHJs and state agency enforcement reporting | Department of State program materials; Title 19 NYCRR framework | src:usa-ny:dos-bsc-overview; src:usa-ny:dos-enforcement-programs | partially_verified |
| ahj:usa-ny:state-agencies | New York State agencies with custody/control of state property | state_agency_enforcement_entities | Administer and enforce the Uniform Code and Energy Code for buildings, premises, and equipment in state agency custody | State agency enforcement program materials | src:usa-ny:dos-enforcement-programs | partially_verified |
| ahj:usa-ny:local-governments | Cities, towns, villages, and applicable counties | local_enforcement_entities | Default local administration and enforcement of the Uniform Code and Energy Code; may petition/file for qualifying local code standards | Executive Law §§379 and 381; Energy Law §11-109 | src:usa-ny:exec-law-379; src:usa-ny:exec-law-381; src:usa-ny:energy-law-11-109 | partially_verified |

### 2.3 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-ny:state-fire-prevention-building-code-council | State Fire Prevention and Building Code Council | Adopts and maintains 2025 BCNYS within the Uniform Code | Executive Law Article 18; 19 NYCRR Parts 1219 and 1221 | src:usa-ny:dos-2025-notice-adoption; src:usa-ny:19nycrr-1219; src:usa-ny:19nycrr-1221 | partially_verified |
| Residential | ahj:usa-ny:state-fire-prevention-building-code-council | State Fire Prevention and Building Code Council | Adopts and maintains 2025 RCNYS within the Uniform Code | Executive Law Article 18; 19 NYCRR Parts 1219 and 1220 | src:usa-ny:dos-2025-notice-adoption; src:usa-ny:19nycrr-1219; src:usa-ny:19nycrr-1220 | partially_verified |
| Existing Building / Rehabilitation | ahj:usa-ny:state-fire-prevention-building-code-council | State Fire Prevention and Building Code Council | Adopts and maintains 2025 EBCNYS within the Uniform Code | Executive Law Article 18; 19 NYCRR Part 1219 | src:usa-ny:dos-2025-notice-adoption; src:usa-ny:19nycrr-1219 | partially_verified |
| Mechanical | ahj:usa-ny:state-fire-prevention-building-code-council | State Fire Prevention and Building Code Council | Adopts and maintains 2025 MCNYS within the Uniform Code | Executive Law Article 18; 19 NYCRR Part 1219 | src:usa-ny:dos-2025-notice-adoption; src:usa-ny:19nycrr-1219 | partially_verified |
| Plumbing | ahj:usa-ny:state-fire-prevention-building-code-council | State Fire Prevention and Building Code Council | Adopts and maintains 2025 PCNYS within the Uniform Code | Executive Law Article 18; 19 NYCRR Part 1219 | src:usa-ny:dos-2025-notice-adoption; src:usa-ny:19nycrr-1219 | partially_verified |
| Fuel Gas | ahj:usa-ny:state-fire-prevention-building-code-council | State Fire Prevention and Building Code Council | Adopts and maintains 2025 FGCNYS within the Uniform Code | Executive Law Article 18; 19 NYCRR Part 1219 | src:usa-ny:dos-2025-notice-adoption; src:usa-ny:19nycrr-1219 | partially_verified |
| Electrical | ahj:usa-ny:state-fire-prevention-building-code-council | State Fire Prevention and Building Code Council | NFPA 70-2023 is incorporated as a referenced standard in Energy Code materials; a standalone statewide electrical code authority was not fully parsed | Energy Code referenced standards; local electrical regulation not fully parsed | src:usa-ny:19nycrr-1240; src:usa-ny:dos-faq | limited_verified |
| Energy | ahj:usa-ny:state-fire-prevention-building-code-council | State Fire Prevention and Building Code Council | Adopts and maintains the Energy Code, including 2025 ECCCNYS and NYS ASHRAE 90.1-2025 | Energy Law Article 11; 19 NYCRR Part 1240 | src:usa-ny:dos-2025-notice-adoption; src:usa-ny:19nycrr-1240; src:usa-ny:energy-law-11-109 | partially_verified |
| Fire - construction references | ahj:usa-ny:state-fire-prevention-building-code-council | State Fire Prevention and Building Code Council | Administers fire-safety construction requirements through the Uniform Code, including Building and Fire Code provisions | Executive Law Article 18; 19 NYCRR Part 1219 | src:usa-ny:dos-uniform-code; src:usa-ny:dos-2025-notice-adoption; src:usa-ny:19nycrr-1219 | partially_verified |
| Fire - operational / prevention code | ahj:usa-ny:state-fire-prevention-building-code-council | State Fire Prevention and Building Code Council | Adopts and maintains 2025 FCNYS within the Uniform Code | Executive Law Article 18; 19 NYCRR Part 1219 | src:usa-ny:dos-uniform-code; src:usa-ny:dos-2025-notice-adoption; src:usa-ny:19nycrr-1219 | partially_verified |
| Property Maintenance | ahj:usa-ny:state-fire-prevention-building-code-council | State Fire Prevention and Building Code Council | Adopts and maintains 2025 PMCNYS within the Uniform Code | Executive Law Article 18; 19 NYCRR Part 1219 | src:usa-ny:dos-2025-notice-adoption; src:usa-ny:19nycrr-1219 | partially_verified |
| Rail Stations | ahj:usa-ny:state-fire-prevention-building-code-council | State Fire Prevention and Building Code Council | Adopts 2025 Uniform Code Provisions for Rail Stations within the Uniform Code | Executive Law Article 18; 19 NYCRR Part 1219 | src:usa-ny:dos-2025-notice-adoption; src:usa-ny:19nycrr-1219 | partially_verified |
| Accessibility | ahj:usa-ny:state-fire-prevention-building-code-council | State Fire Prevention and Building Code Council | Accessibility requirements are expected to be embedded in the Uniform Code; the standalone accessibility authority and amendment path were not fully parsed | Uniform Code framework; detailed accessibility provisions not extracted | src:usa-ny:dos-2025-notice-adoption; src:usa-ny:19nycrr-1219 | unresolved_detail |
| Elevator / Conveyance | ahj:usa-ny:unresolved-elevator-authority | Unresolved | A statewide elevator/conveyance authority and code-adoption record remain unresolved after this update | unresolved | none | unresolved |

### 2.4 Authority Hierarchy Notes

New York's verified statewide code hierarchy is:

State Fire Prevention and Building Code Council
  -> Department of State / Division of Building Standards and Codes support
  -> Title 19 NYCRR Uniform Code and Energy Code regulations
  -> 2025 NYS code books and incorporated referenced standards
  -> local governments and state agencies administering/enforcing in assigned jurisdictions
  -> county or Secretary of State fallback when local enforcement is declined or fails

The Uniform Code is statewide except New York City. The Energy Code is statewide and includes New York City. New York City therefore requires separate handling in any AHJ resolver: Uniform Code and local construction-code rules should branch to New York City law, while Energy Code routing must still account for statewide Energy Code limits and any approved NYC local energy code.

### 2.5 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-ny:001 | ahj:usa-ny:state-fire-prevention-building-code-council | maintains | Uniform Code and Energy Code | src:usa-ny:dos-code-development; src:usa-ny:exec-law-375 | partially_verified |
| edge:usa-ny:002 | ahj:usa-ny:secretary-of-state | chairs_and_administers | State Fire Prevention and Building Code Council / enforcement standards | src:usa-ny:exec-law-374; src:usa-ny:exec-law-381 | partially_verified |
| edge:usa-ny:003 | ahj:usa-ny:dos-dbsc | supports | code development, local enforcement support, and state agency reporting | src:usa-ny:dos-bsc-overview; src:usa-ny:dos-enforcement-programs | partially_verified |
| edge:usa-ny:004 | ahj:usa-ny:local-governments | administer_and_enforce | Uniform Code and Energy Code within local jurisdiction | src:usa-ny:exec-law-381; src:usa-ny:dos-enforcement-programs | partially_verified |
| edge:usa-ny:005 | ahj:usa-ny:counties | fallback_enforces_for | local governments that opt out or otherwise trigger county enforcement | src:usa-ny:exec-law-381 | partially_verified |
| edge:usa-ny:006 | ahj:usa-ny:secretary-of-state | fallback_enforces_for | local/county enforcement gaps and failures | src:usa-ny:exec-law-381 | partially_verified |
| edge:usa-ny:007 | ahj:usa-ny:local-governments | petitions_for_higher_more_restrictive_standards | Code Council approval under Executive Law §379 | src:usa-ny:exec-law-379; src:usa-ny:dos-lg07-local-authority | partially_verified |
| edge:usa-ny:008 | ahj:usa-ny:local-governments | files_more_stringent_energy_code | Code Council review under Energy Law §11-109 | src:usa-ny:energy-law-11-109; src:usa-ny:dos-code-development | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | Building Code of New York State | 2024 ICC family, New York-specific 2025 publication | 2025 | active | 2025-07-25; amended 2025-12-05 | 2025-12-31 | 2025-12-31 | 2025-12-31 | Uniform Code transition allowed either 2020 or 2025 Uniform Code from Notice publication to 2025-12-31; permit applications on/after 2025-12-31 must comply with 2025 Uniform Code | src:usa-ny:dos-2025-notice-adoption; src:usa-ny:19nycrr-1219; src:usa-ny:19nycrr-1221 |
| Residential | Residential Code of New York State | 2024 ICC family, New York-specific 2025 publication | 2025 | active | 2025-07-25; amended 2025-12-05 | 2025-12-31 | 2025-12-31 | 2025-12-31 | Same Uniform Code transition rule | src:usa-ny:dos-2025-notice-adoption; src:usa-ny:19nycrr-1219; src:usa-ny:19nycrr-1220 |
| Existing Building / Rehabilitation | Existing Building Code of New York State | 2024 ICC family, New York-specific 2025 publication | 2025 | active | 2025-07-25; amended 2025-12-05 | 2025-12-31 | 2025-12-31 | 2025-12-31 | Same Uniform Code transition rule | src:usa-ny:dos-2025-notice-adoption; src:usa-ny:19nycrr-1219 |
| Mechanical | Mechanical Code of New York State | 2024 ICC family, New York-specific 2025 publication | 2025 | active | 2025-07-25; amended 2025-12-05 | 2025-12-31 | 2025-12-31 | 2025-12-31 | Same Uniform Code transition rule | src:usa-ny:dos-2025-notice-adoption; src:usa-ny:19nycrr-1219 |
| Plumbing | Plumbing Code of New York State | 2024 ICC family, New York-specific 2025 publication | 2025 | active | 2025-07-25; amended 2025-12-05 | 2025-12-31 | 2025-12-31 | 2025-12-31 | Same Uniform Code transition rule | src:usa-ny:dos-2025-notice-adoption; src:usa-ny:19nycrr-1219 |
| Fuel Gas | Fuel Gas Code of New York State | 2024 ICC family, New York-specific 2025 publication | 2025 | active | 2025-07-25; amended 2025-12-05 | 2025-12-31 | 2025-12-31 | 2025-12-31 | Same Uniform Code transition rule | src:usa-ny:dos-2025-notice-adoption; src:usa-ny:19nycrr-1219 |
| Electrical | NFPA 70 National Electrical Code as referenced standard | NFPA 70 | 2023 | referenced_standard_limited_scope | null | 2025-12-31 for Energy Code references | 2025-12-31 for Energy Code references | 2025-12-31 for applicable Energy Code references | No standalone electrical transition rule verified; treat as referenced-standard use only until the electrical authority scope is parsed | src:usa-ny:19nycrr-1240; src:usa-ny:dos-faq |
| Energy | Energy Conservation Construction Code of New York State; New York State ASHRAE 90.1-2025 | 2024 IECC / ASHRAE 90.1-2022 modified by New York publication | 2025 | active_except_suspended_sections | 2025-07-25; amended 2025-12-05 | 2025-12-31 except suspended provisions | 2025-12-31 except suspended provisions | 2025-12-31 except suspended provisions | No transition period; permit applications on/after 2025-12-31 must comply with 2025 Energy Code. 19 NYCRR §1240.6 is suspended pending court-order termination conditions | src:usa-ny:dos-2025-notice-adoption; src:usa-ny:19nycrr-1240; src:usa-ny:mulhern-stipulation-order |
| Fire - construction references | Fire-safety construction requirements within the 2025 Uniform Code | 2024 ICC family, New York-specific 2025 publications | 2025 | active_except_suspended_subpart_1229_2 | 2025-07-25; amended 2025-12-05 | 2025-12-31 except suspended Subpart 1229-2 | 2025-12-31 except suspended Subpart 1229-2 | 2025-12-31 except suspended Subpart 1229-2 | Same Uniform Code transition rule; 19 NYCRR Subpart 1229-2 is suspended pending court-order termination conditions | src:usa-ny:dos-2025-notice-adoption; src:usa-ny:19nycrr-1219; src:usa-ny:mulhern-stipulation-order |
| Fire - operational / prevention code | Fire Code of New York State | 2024 ICC family, New York-specific 2025 publication | 2025 | active | 2025-07-25; amended 2025-12-05 | 2025-12-31 | 2025-12-31 | 2025-12-31 | Same Uniform Code transition rule | src:usa-ny:dos-uniform-code; src:usa-ny:dos-2025-notice-adoption; src:usa-ny:19nycrr-1219 |
| Property Maintenance | Property Maintenance Code of New York State | 2024 ICC family, New York-specific 2025 publication | 2025 | active | 2025-07-25; amended 2025-12-05 | 2025-12-31 | 2025-12-31 | 2025-12-31 | Same Uniform Code transition rule | src:usa-ny:dos-2025-notice-adoption; src:usa-ny:19nycrr-1219 |
| Rail Stations | Uniform Code Provisions for Rail Stations | New York-specific Uniform Code publication | 2025 | active | 2025-07-25; amended 2025-12-05 | 2025-12-31 | 2025-12-31 | 2025-12-31 | Same Uniform Code transition rule | src:usa-ny:dos-2025-notice-adoption; src:usa-ny:19nycrr-1219 |
| Accessibility | Accessibility provisions within the Uniform Code | New York-specific 2025 Uniform Code publications; detailed base provisions not parsed | 2025 | included_in_uniform_code_not_independently_parsed | 2025-07-25; amended 2025-12-05 | 2025-12-31 | 2025-12-31 | 2025-12-31 | Same Uniform Code transition rule, but detailed accessibility scope remains open | src:usa-ny:dos-2025-notice-adoption; src:usa-ny:19nycrr-1219 |
| Elevator / Conveyance | unresolved | unresolved | unresolved | unresolved | null | null | null | null | Statewide elevator/conveyance authority and transition rule remain unresolved after this update | none |

### 3.2 Adoption Records

| Record ID | Code Families | Action | Adoption Date | Effective Date | Operative Date | Mandatory Date | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| adoption:usa-ny:2025-uniform-code | Building; Residential; Existing Building; Mechanical; Plumbing; Fuel Gas; Fire; Property Maintenance; Rail Stations | Repealed the 2020 Uniform Code edition and adopted the 2025 NYS Uniform Code Books | 2025-07-25 | 2025-12-31 | 2025-12-31 | 2025-12-31 | src:usa-ny:dos-2025-notice-adoption | Uniform Code had a transition option before the effective date. |
| adoption:usa-ny:2025-energy-code | Energy | Repealed the 2020 Energy Code edition and adopted the 2025 ECCCNYS and NYS ASHRAE 90.1-2025 | 2025-07-25 | 2025-12-31 | 2025-12-31 | 2025-12-31 | src:usa-ny:dos-2025-notice-adoption; src:usa-ny:19nycrr-1240 | No transition period. |
| adoption:usa-ny:2025-amended-rule | Uniform Code; Energy | Amended the 2025 code update rules before the effective date | 2025-12-05 | 2025-12-31 | 2025-12-31 | 2025-12-31 | src:usa-ny:dos-2025-notice-adoption | Includes targeted corrections and clarifications; no separate transition period identified for amended rule. |
| adoption:usa-ny:fossil-fuel-provision-suspension | Energy; Uniform Code Subpart 1229-2 | Suspended effective date for 19 NYCRR §1240.6 and 19 NYCRR Subpart 1229-2 pending appellate/certiorari finality conditions | 2025-11-18 | null | null | null | src:usa-ny:mulhern-stipulation-order; src:usa-ny:dos-2025-notice-adoption | Treat suspended provisions as not effective/enforceable until suspension terminates under the order. |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

The verified current transition model separates the Uniform Code and Energy Code. The Uniform Code update has an adoption date of 2025-07-25, an effective date of 2025-12-31, and a transition period from Notice publication to the effective date. During that transition period, regulated parties could choose either the 2020 or 2025 Uniform Code. Beginning 2025-12-31, regulated parties submitting building permit applications must comply with the 2025 Uniform Code.

The Energy Code update also has an adoption date of 2025-07-25 and effective date of 2025-12-31, but the Notice states that there is no transition period. Beginning 2025-12-31, regulated parties submitting building permit applications must comply with the 2025 Energy Code. The exception is the court-suspended fossil-fuel-equipment/building-system provisions in 19 NYCRR §1240.6 and Subpart 1229-2.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-ny:uniform-2025-transition | 2025 Uniform Code | transition_period | 2025-10-01 through 2025-12-30, based on Notice publication to effective date | Regulated party submits or proceeds during transition period before the 2025 Uniform Code effective date | yes; 2020 or 2025 Uniform Code allowed during transition | src:usa-ny:dos-2025-notice-adoption | partially_verified |
| date-rule:usa-ny:uniform-2025-permit-mandatory | 2025 Uniform Code | permit_application_date | 2025-12-31 | Building permit application submitted on or after 2025-12-31 | no | src:usa-ny:dos-2025-notice-adoption | partially_verified |
| date-rule:usa-ny:energy-2025-no-transition | 2025 Energy Code | no_transition_period | 2025-12-31 | Building permit application submitted on or after 2025-12-31 | no | src:usa-ny:dos-2025-notice-adoption | partially_verified |
| date-rule:usa-ny:fossil-fuel-suspension | 19 NYCRR §1240.6 and 19 NYCRR Subpart 1229-2 | court_suspension | from 2025-11-18 until termination conditions in order | Appellate/certiorari finality conditions not yet satisfied | not applicable | src:usa-ny:mulhern-stipulation-order; src:usa-ny:dos-2025-notice-adoption | partially_verified |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Energy; Fire / construction references | Suspended fossil-fuel equipment and building-system provisions | 2025-11-18 court order; DOS notice confirms suspended status | null | null | null | null | active_watch | src:usa-ny:mulhern-stipulation-order; src:usa-ny:dos-2025-notice-adoption | The current watch target is termination or modification of the suspension, not a new code edition. |
| Uniform Code; Energy Code | Next code cycle | null | null | null | null | null | monitor | src:usa-ny:dos-code-development | No next statewide edition was verified in this pass. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-ny:uniform-nyc-exception | Uniform Code | Buildings and structures in New York City | Project is located in New York City | Uniform Code applies statewide except New York City, which has its own construction code | src:usa-ny:dos-code-council-page; src:usa-ny:dos-uniform-code; src:usa-ny:dos-faq | partially_verified |
| applicability-rule:usa-ny:energy-statewide-including-nyc | Energy Code | Buildings and structures statewide | Project is subject to State Energy Code | Energy Code applies statewide, including New York City; approved local energy codes may also apply if more stringent | src:usa-ny:dos-code-council-page; src:usa-ny:dos-faq; src:usa-ny:energy-law-11-109 | partially_verified |
| applicability-rule:usa-ny:state-agency-property | Uniform Code; Energy Code | Buildings, premises, and equipment in custody of state agencies | State agency has custody/control | State of New York, through the relevant state agency framework, is responsible for administration and enforcement for state-agency property | src:usa-ny:dos-enforcement-programs | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** integrated_state_code_books_and_title_19_nycrr

**Where amendments are published:** Title 19 NYCRR Parts 1219-1229 for the Uniform Code, Title 19 NYCRR Part 1240 for the Energy Code, the incorporated 2025 NYS code books, and State Register adoption notices.

**Amendment parsing status:** partially_parsed_for_adoption_dates_and_high_impact_items

### 5.2 State Amendment Sources

| Amendment Source ID | Description | Source IDs | Parsing Status | Notes |
| --- | --- | --- | --- | --- |
| amendment-source:usa-ny:uniform-code-parts-1219-1229 | Title 19 NYCRR Uniform Code rule parts and incorporated 2025 NYS Uniform Code Books | src:usa-ny:dos-laws-regs; src:usa-ny:19nycrr-1219; src:usa-ny:19nycrr-1220; src:usa-ny:19nycrr-1221 | partially_parsed | Parts 1222-1229 were not individually extracted in this pass; Part 1219 and the Notice identify incorporated code books. |
| amendment-source:usa-ny:energy-part-1240 | Title 19 NYCRR Energy Code rule part and incorporated 2025 ECCCNYS / NYS ASHRAE 90.1-2025 | src:usa-ny:19nycrr-1240; src:usa-ny:dos-2025-notice-adoption | partially_parsed | Fossil-fuel provisions are subject to court suspension. |
| amendment-source:usa-ny:notice-of-adoption | DOS Notice of Adoption and Amended Notice of Adoption | src:usa-ny:dos-2025-notice-adoption | parsed_for_dates_and_publications | Used for adoption, transition, and effective-date fields. |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Topic | Summary | Effective / Applicability | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| amendment:usa-ny:energy-u-factor-correction | Energy | Table R402.1.2 correction | Amended rule corrected crawl-space wall insulation maximum U-factor for Climate Zone 4 from 0.55 to 0.055 | 2025-12-31 | src:usa-ny:dos-2025-notice-adoption | partially_verified |
| amendment:usa-ny:building-snow-load-importance-factor | Building | Snow load importance factor | Amended rule added snow-load importance-factor formulas and equations to BCNYS Chapter 16 | 2025-12-31 | src:usa-ny:dos-2025-notice-adoption; src:usa-ny:19nycrr-1221 | partially_verified |
| amendment:usa-ny:energy-appendix-g | Energy | ASHRAE 90.1 Appendix G replacement | Part 1240 indicates Appendix G in ASHRAE 90.1-2022 is replaced by DOS Normative Appendix G for the NYS ASHRAE 90.1-2025 publication | 2025-12-31 | src:usa-ny:19nycrr-1240 | partially_verified |
| amendment:usa-ny:fossil-fuel-suspension | Energy; Uniform Code Subpart 1229-2 | Fossil-fuel equipment and building systems | Effective date of 19 NYCRR §1240.6 and Subpart 1229-2 is suspended pending appellate/certiorari finality conditions | suspended as of 2025-11-18 order | src:usa-ny:mulhern-stipulation-order; src:usa-ny:dos-2025-notice-adoption | partially_verified |
| amendment:usa-ny:truss-symbols | Fire / building administration | Truss-type construction symbol references | Parts 1264 and 1265 were amended to update references from the 2020 FCNYS to the 2025 FCNYS, with no substantive changes identified in the Notice | 2025-12-31 | src:usa-ny:dos-2025-notice-adoption | partially_verified |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-ny"
  model: "local_government_default_with_county_and_secretary_fallback"
  enforcing_entities:
    - "cities"
    - "towns"
    - "villages"
    - "counties when local government opts out or when otherwise responsible"
    - "Secretary of State directly or by contract when local and county enforcement is declined or fails"
    - "state agencies for buildings, premises, and equipment in their custody"
  required_officials:
    - "code enforcement officials and local AHJs; specific credential and staffing requirements not fully parsed"
  state_reserved_activities:
    - "minimum administration and enforcement standards"
    - "state oversight, hearings, and corrective action for deficient local enforcement"
    - "state agency enforcement reporting"
    - "variance and appeal procedures under state rules"
  source_ids:
    - "src:usa-ny:exec-law-381"
    - "src:usa-ny:dos-enforcement-programs"
    - "src:usa-ny:dos-faq"
  verification_status: "partially_verified"
  confidence: 0.86
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-ny"
  model: "uniform_code_higher_more_restrictive_with_code_council_approval; energy_code_more_stringent_with_filing_and_code_council_review"
  applies_to_code_families:
    - "Uniform Code building and fire prevention standards"
    - "State Energy Conservation Construction Code"
  approval_required: "true for Uniform Code higher/more restrictive standards; Energy Code requires filing and is subject to Code Council determination if not more stringent"
  approving_authority_id: "ahj:usa-ny:state-fire-prevention-building-code-council"
  filing_required: true
  registry_exists: "partial"
  registry_source_ids:
    - "src:usa-ny:dos-code-council-page"
    - "src:usa-ny:dos-local-standard-list"
  legal_basis_source_ids:
    - "src:usa-ny:exec-law-379"
    - "src:usa-ny:energy-law-11-109"
    - "src:usa-ny:dos-code-development"
    - "src:usa-ny:dos-lg07-local-authority"
  verification_status: "partially_verified"
  confidence: 0.84
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement authority does not imply local amendment authority. New York local governments generally administer and enforce the statewide codes in their jurisdictions, but local construction standards that are higher or more restrictive than the Uniform Code require the Executive Law §379 process and Code Council approval before they become effective. Local energy codes follow a separate Energy Law §11-109 more-stringent-code filing and review model.

### 6.4 Known Local Amendment Registries

| Registry ID | Registry Name | Coverage | Source IDs | Status | Caveat |
| --- | --- | --- | --- | --- | --- |
| registry:usa-ny:more-restrictive-local-standards | Code Council list of approved more restrictive local standards and more stringent local energy codes | Statewide list maintained on the Code Council page | src:usa-ny:dos-code-council-page; src:usa-ny:dos-local-standard-list | partial | The visible page includes a last-updated date of 2022-09-27 for at least part of the list. Treat as useful but not complete for current production use without confirmation. |

### 6.5 Municipality-Specific Known Amendments

| Locality | Code Family | Amendment / Local Code | Source IDs | Status | Notes |
| --- | --- | --- | --- | --- | --- |
| New York City | Energy | Approved local energy code, details not parsed in this pass | src:usa-ny:dos-faq; src:usa-ny:dos-code-council-page | identified_not_parsed | The FAQ confirms NYC has an approved local energy code. The actual local code text and version were not parsed here. |
| Village of Saltaire | Uniform Code / fire sprinklers | Local law related to fire sprinkler systems in one- and two-family dwellings identified in DOS local-standard table | src:usa-ny:dos-local-standard-list | identified_not_parsed | Table source is stale/partial; details require current confirmation. |

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: not_started

Jurisdiction stack:

```text
Address
  -> State: New York
  -> New York City branch check
  -> County
  -> City / town / village or other local government
  -> State agency custody check, if applicable
  -> Local code administration/enforcement program
  -> County fallback or Secretary of State fallback, if local enforcement declined or deficient
  -> Building AHJ
  -> Fire prevention / operational fire AHJ
  -> Trade-specific AHJs, if applicable
  -> 2025 Uniform Code adoption records, except NYC Uniform Code branch
  -> 2025 Energy Code adoption record, statewide including NYC
  -> Approved higher/more restrictive local standards and more stringent local energy codes
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | not selected | none | statewide | unknown | pending |
| County | not selected | none | statewide | unknown | pending |
| City / town / village | not selected | none | statewide | unknown | pending |
| New York City boundary | not selected | none | NYC | unknown | pending |
| Fire District | not selected | none | statewide | unknown | pending |
| State agency custody / facility jurisdiction | not selected | none | statewide | unknown | pending |

### 7.3 AHJ Contact Data

No AHJ contact data was populated for this pass. AHJ contact and boundary records remain open before address-level production use.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Publisher / Agency | Source Type | URL | Date / Version | Supports | Caveats |
| --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-ny:dos-bsc-overview | Building Standards and Codes | New York Department of State | agency_page | https://dos.ny.gov/building-standards-and-codes | accessed 2026-06-26 | DBSC role; Title 19 NYCRR framework | Agency web page; use statutes/regulations for legal text. |
| src:usa-ny:dos-code-development | Code Development | New York Department of State | agency_page | https://dos.ny.gov/code/code-development | accessed 2026-06-26 | Code Council authority; local more restrictive / more stringent code review | Agency summary; cite statutes for controlling legal text. |
| src:usa-ny:dos-code-council-page | State Fire Prevention and Building Code Council | New York Department of State | agency_page | https://dos.ny.gov/state-fire-prevention-and-building-code-council | accessed 2026-06-26 | NYC exception; Code Council page; local standards table | Local standards table may be stale/partial. |
| src:usa-ny:dos-local-standard-list | Code Council approved local standards / energy filings table | New York Department of State | agency_page_table | https://dos.ny.gov/state-fire-prevention-and-building-code-council | visible last-updated marker 2022-09-27 | Local amendment registry existence and examples | Treat as partial until DOS confirms current registry status. |
| src:usa-ny:dos-laws-regs | Laws and Regulations, Division of Building Standards and Codes | New York Department of State | agency_page | https://dos.ny.gov/laws-and-regulations-division-building-standards-and-codes | courtesy copies current as of 2026-02-20 per page | Title 19 NYCRR parts and statutory links | Courtesy copy; verify against official NYCRR/State Register for production. |
| src:usa-ny:dos-uniform-code | Uniform Fire Prevention and Building Code | New York Department of State | agency_page | https://dos.ny.gov/uniform-fire-prevention-and-building-code | accessed 2026-06-26 | Uniform Code scope, NYC exception, local higher/more restrictive standards | Agency summary; use statutes/regulations for controlling language. |
| src:usa-ny:dos-faq | Division of Building Standards and Codes Frequently Asked Questions | New York Department of State | agency_faq | https://dos.ny.gov/division-building-standards-and-codes-frequently-asked-questions | accessed 2026-06-26 | NYC exception; statewide Energy Code; local enforcement summary; 2025 code current status | FAQ includes some legacy 2020 wording elsewhere; use only verified current statements and cross-check with Notice/NYCRR. |
| src:usa-ny:dos-2025-notice-adoption | Notice of Adoption: 2025 Uniform Code and Energy Code update | New York Department of State | adoption_notice | https://dos.ny.gov/notice-adoption | original adoption 2025-07-25; amended adoption 2025-12-05; effective 2025-12-31 | 2025 code books; adoption/effective/mandatory dates; transition rules; suspended provisions note | Agency page summarizing State Register notices; State Register text should be archived for long-term legal provenance. |
| src:usa-ny:mulhern-stipulation-order | Mulhern et al. v. Mosley So Ordered Stipulation | U.S. District Court / hosted by NY DOS | court_order_pdf | https://dos.ny.gov/mulhern-so-ordered-stipulation-11-18-25 | ordered 2025-11-18 | Suspension of 19 NYCRR §1240.6 and Subpart 1229-2 | PDF hosted by DOS; docket should be checked for later modifications. |
| src:usa-ny:19nycrr-1219 | 19 NYCRR Part 1219: Uniform Code definitions and incorporated publications | New York Department of State | regulation_pdf | https://dos.ny.gov/19-nycrr-part-1219 | 2025 rule text | Uniform Code composition and code-book definitions | Courtesy/agency PDF; verify official NYCRR text for production. |
| src:usa-ny:19nycrr-1220 | 19 NYCRR Part 1220: Residential Construction | New York Department of State | regulation_pdf | https://dos.ny.gov/19-nycrr-part-1220 | 2025 rule text | 2025 RCNYS applicability and incorporation | Courtesy/agency PDF; verify official NYCRR text for production. |
| src:usa-ny:19nycrr-1221 | 19 NYCRR Part 1221: Building Construction | New York Department of State | regulation_pdf | https://dos.ny.gov/19-nycrr-part-1221 | 2025 rule text | 2025 BCNYS applicability, exceptions, amendments | Courtesy/agency PDF; verify official NYCRR text for production. |
| src:usa-ny:19nycrr-1240 | 19 NYCRR Part 1240: State Energy Conservation Construction Code | New York Department of State | regulation_pdf | https://dos.ny.gov/state-energy-conservation-construction-code-0 | 2025 rule text | 2025 ECCCNYS, NYS ASHRAE 90.1-2025, NFPA 70 references, Energy Code applicability | Courtesy/agency PDF; fossil-fuel provisions partly suspended by order. |
| src:usa-ny:exec-law-374 | Executive Law §374: State Fire Prevention and Building Code Council | NYSenate.gov / Open Legislation | statute_html | https://www.nysenate.gov/legislation/laws/EXC/374 | accessed 2026-06-26 | Code Council creation and composition | Online statute; verify session-law amendments if needed. |
| src:usa-ny:exec-law-375 | Executive Law §375: Council powers and duties | NYSenate.gov / Open Legislation | statute_html | https://www.nysenate.gov/legislation/laws/EXC/375 | accessed 2026-06-26 | Code Council powers/duties | Online statute; verify session-law amendments if needed. |
| src:usa-ny:exec-law-379 | Executive Law §379: Higher local standards | NYSenate.gov / Open Legislation | statute_html | https://www.nysenate.gov/legislation/laws/EXC/379 | accessed 2026-06-26 | Local higher/more restrictive construction standards | Online statute; verify session-law amendments if needed. |
| src:usa-ny:exec-law-381 | Executive Law §381: Administration and enforcement | NYSenate.gov / Open Legislation | statute_html | https://www.nysenate.gov/legislation/laws/EXC/381 | accessed 2026-06-26 | Local enforcement model; county/SOS fallback | Online statute; verify session-law amendments if needed. |
| src:usa-ny:energy-law-11-109 | Energy Law §11-109: Local energy conservation construction codes | NYSenate.gov / Open Legislation | statute_html | https://www.nysenate.gov/legislation/laws/ENG/11-109 | accessed 2026-06-26 | Local more stringent energy code filing/review | Online statute; verify session-law amendments if needed. |
| src:usa-ny:dos-enforcement-programs | Local Government and State Agency Enforcement Programs | New York Department of State | agency_page | https://dos.ny.gov/code/local-government-state-agency-enforcement-programs | accessed 2026-06-26 | Local and state agency enforcement reporting and support | Agency summary; detailed Part 1203 rules not fully parsed. |
| src:usa-ny:dos-lg07-local-authority | Legal Memorandum LG07: The Uniform Code and Local Authority | New York Department of State | agency_legal_memorandum | https://dos.ny.gov/legal-memorandum-lg07-uniform-code-and-local-authority | accessed 2026-06-26 | Local authority limits and approval process | Informal/non-binding explanatory memorandum; not a substitute for statute/regulation. |
| src:usa-ny:dos-code-outreach-2026-01 | 2026 Code Outreach Program: 19 NYCRR changes for 2025 codes | New York Department of State | agency_guidance_pdf | https://dos.ny.gov/2026-01-nycrr-changes-2025-codes | January 2026 | Outreach explanation of 2025 code structure and notable changes | Guidance/outreach; do not rely on it as controlling legal text. |
| src:usa-ny:dos-manufactured-housing-2025-uniform | 2025 Uniform Code Adoption PDF / Manufactured Homes Program notice | New York Department of State | agency_program_pdf | https://dos.ny.gov/2025-uniform-code-adoptionpdf | 2025 | Additional confirmation of Uniform Code adoption/effective date for a program context | Program notice; use Notice/NYCRR for core adoption fields. |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-ny:dos-laws-regs | courtesy_copy | DOS states that listed documents are courtesy of the Secretary of State and current as of 2026-02-20. | Cross-check final production records against official NYCRR and State Register text. |
| src:usa-ny:dos-faq | mixed_currency | FAQ includes current 2025 code statements but also contains legacy 2020 Energy Code wording in another section. | Use only cross-checked statements; prefer Notice/NYCRR for code adoption. |
| src:usa-ny:dos-local-standard-list | partial_registry | Local standards / local energy filings table appears to include a 2022 last-updated marker and may not be complete after the 2025 code update. | Treat as registry existence/example evidence only until confirmed. |
| src:usa-ny:mulhern-stipulation-order | litigation_status | Court order may be superseded or terminated by later appellate events. | Monitor docket/DOS notices before relying on suspension status in production. |
| src:usa-ny:dos-lg07-local-authority | informal_guidance | Legal memorandum is explanatory and not binding legal advice. | Use as supplemental interpretation only; cite statutes for controlling requirements. |

### 8.3 Supplemental Sources

None used in this pass. All cited source IDs are official state, statutory, regulatory, or court-order sources, except that agency guidance and legal memoranda are treated as official but non-controlling support.

### 8.4 Source Extraction Metadata

| Extraction ID | Source IDs | Extraction Method | Extracted Fields | Extracted On | Extracted By | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| extraction:usa-ny:authority-001 | src:usa-ny:dos-code-development; src:usa-ny:exec-law-374; src:usa-ny:exec-law-375 | official web text review | primary authority, Code Council role | 2026-06-26 | GPT-5.5 Thinking | Verified authority model at source level. |
| extraction:usa-ny:adoption-001 | src:usa-ny:dos-2025-notice-adoption; src:usa-ny:19nycrr-1219; src:usa-ny:19nycrr-1220; src:usa-ny:19nycrr-1221; src:usa-ny:19nycrr-1240 | official web/PDF text review | code books, adoption/effective/mandatory dates, transition rules | 2026-06-26 | GPT-5.5 Thinking | Individual Parts 1222-1229 not fully parsed. |
| extraction:usa-ny:local-001 | src:usa-ny:exec-law-379; src:usa-ny:exec-law-381; src:usa-ny:energy-law-11-109; src:usa-ny:dos-lg07-local-authority | statute and agency guidance review | local enforcement and local amendment model | 2026-06-26 | GPT-5.5 Thinking | Local registry remains incomplete. |
| extraction:usa-ny:litigation-001 | src:usa-ny:mulhern-stipulation-order; src:usa-ny:dos-2025-notice-adoption | court-order and agency-notice review | suspended provisions | 2026-06-26 | GPT-5.5 Thinking | Docket monitoring required. |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 1.00 | none | Status upgraded only after source IDs were registered and validation checks were run. |
| report | report.last_updated | 2026-06-26 | verified | 1.00 | none | Current update date for this pass. |
| authority:primary | primary_authority | State Fire Prevention and Building Code Council | partially_verified | 0.95 | src:usa-ny:dos-code-development; src:usa-ny:exec-law-374; src:usa-ny:exec-law-375 | Statutory/agency support captured. |
| adoption:uniform | current_uniform_code | 2025 NYS Uniform Code Books | partially_verified | 0.90 | src:usa-ny:dos-2025-notice-adoption; src:usa-ny:19nycrr-1219 | Individual Uniform Code parts not all parsed. |
| adoption:energy | current_energy_code | 2025 ECCCNYS and NYS ASHRAE 90.1-2025 | partially_verified | 0.90 | src:usa-ny:dos-2025-notice-adoption; src:usa-ny:19nycrr-1240 | Suspended provisions excluded from effective/enforceable status. |
| date-rule:uniform | mandatory_date | 2025-12-31 | partially_verified | 0.90 | src:usa-ny:dos-2025-notice-adoption | Permit-application trigger captured. |
| date-rule:energy | transition_rule | no transition period | partially_verified | 0.90 | src:usa-ny:dos-2025-notice-adoption | Permit-application trigger captured. |
| local-enforcement | model | local_government_default_with_county_and_secretary_fallback | partially_verified | 0.86 | src:usa-ny:exec-law-381; src:usa-ny:dos-enforcement-programs | Detailed credential/staffing rules remain open. |
| local-amendment | model | Uniform higher/more restrictive approval; Energy more stringent filing/review | partially_verified | 0.84 | src:usa-ny:exec-law-379; src:usa-ny:energy-law-11-109; src:usa-ny:dos-lg07-local-authority | Registry completeness remains open. |
| adoption:electrical | electrical_code_scope | NFPA 70-2023 referenced in Energy Code materials; standalone authority unresolved | limited_verified | 0.55 | src:usa-ny:19nycrr-1240; src:usa-ny:dos-faq | Electrical licensing and standalone adoption not fully parsed. |
| adoption:elevator | elevator_conveyance_authority | unresolved | unresolved | 0.10 | none | Dedicated research needed. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Every src:usa-ny ID used in the body is defined in Section 8. |
| All authority IDs resolve | pass | Authority IDs used in authority tables are defined or deliberately marked unresolved. |
| All current code families have adoption records | partial_pass | Core statewide Uniform Code and Energy Code families have adoption records; elevator/conveyance remains unresolved. |
| Building and operational fire code are separated | pass | Fire construction references and operational/prevention fire code are separate rows. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Date columns are separate, and transition language distinguishes triggers. |
| Effective dates are valid ISO dates | pass | Verified date fields use YYYY-MM-DD or null. |
| No impossible date sequences | pass | Adoption dates precede effective/mandatory dates for verified records. |
| Transition rules have explicit trigger conditions | pass | Uniform and Energy transition/permit-date triggers are captured. |
| Permit-date logic is captured where applicable | pass | 2025 Uniform and Energy permit-application triggers are captured. |
| Local enforcement model classified | pass | Local default / county fallback / Secretary fallback model is captured. |
| Local amendment rule classified | pass | Uniform Code and Energy Code local amendment models are separated. |
| AHJ confirmation metadata present | fail | AHJ contact and boundary data were not populated. |
| Official-source caveats captured | pass | Courtesy-copy, stale-registry, FAQ, and litigation caveats are captured. |
| Leftover template markers removed | pass | Automated scan found no template marker patterns listed in the validation instructions. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-ny:001 | medium | official NYCRR verification | DOS regulation PDFs are treated as courtesy/agency copies. | Cross-check Parts 1219-1229 and 1240 against official NYCRR or State Register records. | null | null | open |
| issue:usa-ny:002 | high | suspended fossil-fuel provisions | The court-order suspension may terminate or be modified by later appellate/certiorari events. | Monitor the docket and DOS notices for changes to the suspension of 19 NYCRR §1240.6 and Subpart 1229-2. | null | null | open |
| issue:usa-ny:003 | medium | electrical authority | NFPA 70-2023 was verified as a referenced standard, but standalone statewide electrical code authority and local electrical licensing scope were not parsed. | Research state and local electrical licensing/adoption statutes and any separate electrical board authority. | null | null | open |
| issue:usa-ny:004 | medium | accessibility authority | Accessibility provisions are included in the Uniform Code framework, but standalone state accessibility authority and amendment sources were not parsed. | Parse 2025 BCNYS accessibility provisions and relevant New York accessibility statutes/rules. | null | null | open |
| issue:usa-ny:005 | medium | elevator / conveyance | A statewide elevator/conveyance authority and adoption record were not verified. | Research New York elevator/conveyance statutes, agency jurisdiction, and referenced code editions. | null | null | open |
| issue:usa-ny:006 | high | local amendment registry | DOS local-standard / local energy table may be stale or incomplete after the 2025 code update. | Obtain or confirm current Code Council registry of approved local standards and filed local energy codes. | null | null | open |
| issue:usa-ny:007 | high | AHJ resolver | Boundary data, AHJ contacts, and NYC/non-NYC resolver logic were not populated. | Select boundary data sources and build local enforcement/AHJ contact dataset. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-ny:notice-adoption | src:usa-ny:dos-2025-notice-adoption | html_diff | monthly | New adoption notice, amended notice, effective-date change, or suspension update | 2026-06-26 | active |
| watch:usa-ny:laws-regs | src:usa-ny:dos-laws-regs | html_diff | monthly | Title 19 NYCRR parts or courtesy-copy currency date changes | 2026-06-26 | active |
| watch:usa-ny:code-council | src:usa-ny:dos-code-council-page | html_diff | monthly | Code Council meeting/action, local standard approval, or registry update | 2026-06-26 | active |
| watch:usa-ny:local-standard-list | src:usa-ny:dos-local-standard-list | manual_registry_check | monthly | Local standards or local energy-code table changes | 2026-06-26 | active |
| watch:usa-ny:exec-law-article-18 | src:usa-ny:exec-law-374 | statute_diff | quarterly | Executive Law Article 18 amendment | 2026-06-26 | active |
| watch:usa-ny:energy-law-11 | src:usa-ny:energy-law-11-109 | statute_diff | quarterly | Energy Law Article 11 amendment affecting local energy codes | 2026-06-26 | active |
| watch:usa-ny:mulhern-order | src:usa-ny:mulhern-stipulation-order | docket_or_agency_notice_check | monthly | Suspension terminated, modified, appealed, or superseded | 2026-06-26 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-24 | Baseline draft report existed with unresolved placeholder content | report:usa-ny | none | prior draft generator | Starting point for this pass. |
| 2026-06-26 | Replaced baseline stub with partially verified New York report | report:usa-ny; authority:primary; adoption:usa-ny:2025-uniform-code; adoption:usa-ny:2025-energy-code; local-enforcement:usa-ny; local-amendment-rule:usa-ny | src:usa-ny:dos-code-development; src:usa-ny:dos-2025-notice-adoption; src:usa-ny:exec-law-381; src:usa-ny:exec-law-379; src:usa-ny:energy-law-11-109 | GPT-5.5 Thinking | Source registry populated; unresolved items retained explicitly; validation pass completed with AHJ/elevator/accessibility/electrical gaps open. |
