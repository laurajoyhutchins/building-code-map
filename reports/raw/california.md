---
state:
  state_id: "US-CA"
  name: "California"
  abbreviation: "CA"
report:
  report_id: "state-report:usa-ca"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-25"
  last_verified: null
  reviewed_by: null
risk:
  overall_confidence: 0.62 # 0.00 - 1.00
  risk_flags:
    - "model_code_base_editions_partly_supported_by_publisher_or_catalog_sources"
    - "state_amendments_not_section_by_section_parsed"
    - "local_ahj_boundary_sources_not_selected"
    - "elevator_conveyance_scope_requires_specialist_review"
  open_questions_count: 6

---

# State Building Code Authority Report: California

## 1. Executive Summary

- **Authority model:** California has a statewide building-standards system centered on the California Building Standards Commission (CBSC). State agencies propose or adopt building standards within their enabling authority, and those standards generally must be submitted to CBSC for approval or adoption before codification in the California Building Standards Code, Title 24, California Code of Regulations. `src:usa-ca:hsc-18930`, `src:usa-ca:hsc-18938`, `src:usa-ca:hsc-18942`, `src:usa-ca:cbsc-about`

- **Statewide code status:** The 2025 California Building Standards Code, Title 24, is the current statewide code edition identified in this pass. CBSC identifies it as effective January 1, 2026 through December 31, 2028, and its code page states the 2025 edition was published July 1, 2025 with an effective date of January 1, 2026. `src:usa-ca:cbsc-home`, `src:usa-ca:cbsc-codes`

- **Local enforcement model:** Enforcement is hybrid. City and county building departments enforce State Building Standards Code provisions for covered residential/housing occupancies within their jurisdictions; local housing, building, health, or environmental agencies also have maintenance/use/occupancy enforcement roles. Energy standards are primarily enforced by local building departments, with California Energy Commission (CEC) fallback or dispute roles. Fire and panic safety standards are enforced by the State Fire Marshal and local fire chiefs or districts in their respective areas. `src:usa-ca:hsc-17960`, `src:usa-ca:hsc-17961`, `src:usa-ca:prc-25402-1`, `src:usa-ca:hsc-13145`

- **Local amendment posture:** Cities and counties may adopt local modifications where authorized and supported by express findings of local climatic, geological, or topographical conditions, but residential-unit amendments are subject to a temporary moratorium and exceptions from October 1, 2025 through June 1, 2031. Findings and amendments must be filed before effectiveness for the covered local-modification paths. Fire protection districts have a separate more-stringent fire/panic-safety amendment path subject to city/county ratification. CBSC maintains a local ordinance filing registry. `src:usa-ca:hsc-17958`, `src:usa-ca:hsc-17958-5`, `src:usa-ca:hsc-17958-7`, `src:usa-ca:hsc-18941-5`, `src:usa-ca:hsc-13869-7`, `src:usa-ca:cbsc-local-ordinances`

- **Known transition periods or pending changes:** State building standards generally become effective 180 days after publication unless CBSC establishes a later date. California also applies a permit-application-date rule: building standards approved by CBSC and effective at the local level when a building permit application is submitted apply to the plans, specifications, and construction under that permit. CEC states that buildings with permit applications applied for on or after January 1, 2026 must comply with the 2025 Energy Code. CBSC has begun the 2025 Intervening Code Adoption Cycle for supplements to the 2025 edition. `src:usa-ca:hsc-18938`, `src:usa-ca:hsc-18938-5`, `src:usa-ca:cec-2025-energy`, `src:usa-ca:cbsc-2025-intervening-cycle`

- **Production readiness:** partially_ready_for_narrow_validation

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: CBSC is the statewide approval/adoption and publication hub for California
    building standards; state agency building standards must generally be submitted
    to and approved or adopted by CBSC before codification.
  confidence: 0.86
  source_ids:
  - src:usa-ca:hsc-18930
  - src:usa-ca:hsc-18938
  - src:usa-ca:hsc-18942
  - src:usa-ca:cbsc-about
- topic: Current Title 24 edition
  finding: 2025 California Building Standards Code, Title 24, effective January 1,
    2026 to December 31, 2028.
  confidence: 0.84
  source_ids:
  - src:usa-ca:cbsc-home
  - src:usa-ca:cbsc-codes
- topic: Primary building code edition
  finding: 2025 California Building Code, Title 24 Part 2; base model code identified
    as the 2024 International Building Code by publisher/catalog sources.
  confidence: 0.7
  source_ids:
  - src:usa-ca:cbsc-codes
  - src:usa-ca:icc-cbc-2025
- topic: Electrical code edition
  finding: 2025 California Electrical Code, Title 24 Part 3, based on the 2023 National
    Electrical Code according to catalog/publisher sources; current statewide effective
    date follows the 2025 Title 24 cycle.
  confidence: 0.68
  source_ids:
  - src:usa-ca:cbsc-codes
  - src:usa-ca:saclaw-cec-2025
  - src:usa-ca:icc-store-cec-2025
- topic: Energy code applicability
  finding: 2025 Energy Code applies to buildings with permit applications applied
    for on or after January 1, 2026.
  confidence: 0.9
  source_ids:
  - src:usa-ca:cec-2025-energy
- topic: Fire code authority
  finding: OSFM promulgates fire/life-safety regulations for inclusion in Title 24
    and the State Fire Marshal/local fire chiefs enforce fire and panic-safety standards
    in their respective areas.
  confidence: 0.82
  source_ids:
  - src:usa-ca:osfm-title24
  - src:usa-ca:hsc-13143
  - src:usa-ca:hsc-13145
  - src:usa-ca:osfm-fire-life-safety
- topic: Local amendments
  finding: City/county local modifications require statutory findings and filing;
    residential modifications are restricted from October 1, 2025 through June 1,
    2031 unless an exception applies.
  confidence: 0.83
  source_ids:
  - src:usa-ca:hsc-17958
  - src:usa-ca:hsc-17958-5
  - src:usa-ca:hsc-17958-7
  - src:usa-ca:hsc-18941-5
- topic: Permit-date rule
  finding: The applicable building standards are those approved by CBSC and effective
    locally when the building permit application is submitted.
  confidence: 0.88
  source_ids:
  - src:usa-ca:hsc-18938-5
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | `ahj:usa-ca:cbsc` |
| Authority name | California Building Standards Commission |
| Authority type | statewide building standards commission |
| Legal basis | California Building Standards Law, including Health and Safety Code §§ 18930, 18938, and 18942 |
| Role | Reviews, approves/adopts, codifies, and publishes statewide building standards in Title 24; publishes triennial editions and supplements; receives local ordinance filings where statutes require filing. |
| Enforcement model | Statewide code publication and approval body; local building, housing, fire, and specialized state agencies enforce within assigned scopes. |
| Source IDs | `src:usa-ca:hsc-18930`; `src:usa-ca:hsc-18938`; `src:usa-ca:hsc-18942`; `src:usa-ca:cbsc-about`; `src:usa-ca:cbsc-local-ordinances` |
| Verification status | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | `ahj:usa-ca:cbsc` | California Building Standards Commission | Statewide approval/adoption and publication of Title 24 building standards; state agencies participate by occupancy/scope. | HSC §§ 18930, 18938, 18942 | `src:usa-ca:hsc-18930`; `src:usa-ca:hsc-18938`; `src:usa-ca:hsc-18942`; `src:usa-ca:cbsc-codes` | partially_verified |
| Residential | `ahj:usa-ca:hcd` | California Department of Housing and Community Development, with CBSC approval/publication | Develops residential/housing building standards for Title 24; local building departments enforce housing construction standards within jurisdiction. | HSC §§ 17922, 17960; HSC Part 2.5 CBSC approval process | `src:usa-ca:hsc-17922`; `src:usa-ca:hsc-17960`; `src:usa-ca:hcd-title24`; `src:usa-ca:hcd-shl-laws` | partially_verified |
| Existing Building / Rehabilitation | `ahj:usa-ca:cbsc` | California Building Standards Commission plus proposing agencies by occupancy | Current edition identified as 2025 California Existing Building Code, Title 24 Part 10; section-by-section authority not parsed. | HSC §§ 18930, 18938, 18942 | `src:usa-ca:cbsc-codes`; `src:usa-ca:saclaw-cebc-2025`; `src:usa-ca:hsc-18930` | partially_verified |
| Mechanical | `ahj:usa-ca:cbsc` | California Building Standards Commission plus proposing agencies by occupancy | Current edition identified as 2025 California Mechanical Code, Title 24 Part 4; section-by-section authority not parsed. | HSC §§ 18930, 18938, 18942 | `src:usa-ca:cbsc-codes`; `src:usa-ca:saclaw-cmc-2025`; `src:usa-ca:hsc-18930` | partially_verified |
| Plumbing | `ahj:usa-ca:cbsc` | California Building Standards Commission plus proposing agencies by occupancy | Current edition identified as 2025 California Plumbing Code, Title 24 Part 5; section-by-section authority not parsed. | HSC §§ 18930, 18938, 18942 | `src:usa-ca:cbsc-codes`; `src:usa-ca:saclaw-cpc-2025`; `src:usa-ca:hsc-18930` | partially_verified |
| Fuel Gas | `ahj:usa-ca:unresolved-fuel-gas` | No separate statewide fuel-gas code authority confirmed | No separate California Fuel Gas Code row was confirmed; fuel-gas provisions appear to require follow-up within mechanical/plumbing code scopes. | unresolved | `src:usa-ca:cbsc-codes` | unresolved |
| Electrical | `ahj:usa-ca:cbsc` | California Building Standards Commission plus proposing agencies by occupancy | Current edition identified as 2025 California Electrical Code, Title 24 Part 3. | HSC §§ 18930, 18938, 18942 | `src:usa-ca:cbsc-codes`; `src:usa-ca:saclaw-cec-2025`; `src:usa-ca:icc-store-cec-2025` | partially_verified |
| Energy | `ahj:usa-ca:cec` | California Energy Commission, with CBSC approval/publication | Develops/adopts Building Energy Efficiency Standards; local building departments enforce, with CEC fallback/dispute authority. | PRC § 25402.1; Title 24 Part 6 publication by CBSC | `src:usa-ca:prc-25402-1`; `src:usa-ca:cec-2025-energy`; `src:usa-ca:cec-energy-adoption-news` | partially_verified |
| Fire - construction references | `ahj:usa-ca:osfm` | Office of the State Fire Marshal | Promulgates fire/life-safety regulations for inclusion in Title 24, including California Fire Code and related construction-code provisions. | HSC §§ 13143, 18930 | `src:usa-ca:osfm-title24`; `src:usa-ca:hsc-13143`; `src:usa-ca:hsc-18930` | partially_verified |
| Fire - operational / prevention code | `ahj:usa-ca:osfm` | Office of the State Fire Marshal plus local fire chiefs/districts | Adopts/enforces other fire-prevention regulations and enforces fire and panic-safety standards in respective state/local areas. | HSC §§ 13143, 13145 | `src:usa-ca:hsc-13143`; `src:usa-ca:hsc-13145`; `src:usa-ca:osfm-fire-life-safety` | partially_verified |
| Accessibility | `ahj:usa-ca:dsa-access` | Division of the State Architect, Access Compliance | DSA submits accessibility code changes to CBSC for review/approval; full Chapter 11A/11B split not parsed in this pass. | Title 24 accessibility code-development path; HSC § 18930 process | `src:usa-ca:dsa-access-code-development`; `src:usa-ca:hsc-18930`; `src:usa-ca:hcd-title24` | partially_verified |
| Elevator / Conveyance | `ahj:usa-ca:dir-dosh-elevator` | Department of Industrial Relations, Division of Occupational Safety and Health, Elevator Unit | Elevator Safety Orders under Title 8 apply to elevators in California subject to stated exceptions; Cal/OSHA Elevator Unit issues operating permits. | CCR Title 8 Elevator Safety Orders | `src:usa-ca:dir-elevator-safety-orders`; `src:usa-ca:dir-elevator-permits` | partially_verified |

### 2.3 Authority Hierarchy Notes

California's code-adoption structure is not a simple single-agency model. CBSC is the central statewide code approval/publication body. Specialized state agencies develop or adopt standards within statutory scope, including HCD for housing/residential standards, CEC for energy standards, OSFM for fire and panic safety, DSA for access compliance, and DIR/DOSH for elevator/conveyance safety orders. Local building and fire agencies enforce within jurisdictional and occupancy-specific scopes. Local amendment authority is separate from enforcement authority and is constrained by filing, findings, ratification, and temporary residential-moratorium rules.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| `edge:usa-ca:001` | `ahj:usa-ca:hcd` | submits_or_develops_for | `ahj:usa-ca:cbsc` / residential and housing building standards in Title 24 | `src:usa-ca:hcd-title24`; `src:usa-ca:hsc-17922`; `src:usa-ca:hsc-18930` | partially_verified |
| `edge:usa-ca:002` | `ahj:usa-ca:cec` | adopts_energy_standards_for_publication_by | `ahj:usa-ca:cbsc` / Title 24 Part 6 | `src:usa-ca:cec-energy-adoption-news`; `src:usa-ca:cec-2025-energy`; `src:usa-ca:prc-25402-1` | partially_verified |
| `edge:usa-ca:003` | `ahj:usa-ca:osfm` | promulgates_fire_life_safety_standards_for | Title 24 parts including Building, Fire, Electrical, Mechanical, Plumbing, and Historical Building codes | `src:usa-ca:osfm-title24`; `src:usa-ca:hsc-13143` | partially_verified |
| `edge:usa-ca:004` | `local-building-departments` | enforce | State Building Standards Code provisions for covered residential/housing occupancies within jurisdiction | `src:usa-ca:hsc-17960`; `src:usa-ca:hsc-17961` | partially_verified |
| `edge:usa-ca:005` | `local-building-departments` | enforce | Energy Code, with CEC fallback if no local department or failure to enforce | `src:usa-ca:prc-25402-1` | partially_verified |
| `edge:usa-ca:006` | `ahj:usa-ca:osfm` | shares_enforcement_with | local fire chiefs/districts for fire and panic-safety standards in respective areas | `src:usa-ca:hsc-13145`; `src:usa-ca:osfm-fire-life-safety` | partially_verified |
| `edge:usa-ca:007` | `local-governing-bodies` | may_modify_with_findings_and_filing | California Building Standards Code requirements where authorized | `src:usa-ca:hsc-17958-5`; `src:usa-ca:hsc-17958-7`; `src:usa-ca:hsc-18941-5` | partially_verified |
| `edge:usa-ca:008` | `fire-protection-districts` | may_adopt_more_stringent_fire_panic_standards_subject_to_ratification | applicable city/county/city-and-county and filing path | `src:usa-ca:hsc-13869-7` | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | 2025 California Building Code, Title 24 Part 2 | 2024 International Building Code | 2025 | current statewide Title 24 edition | null | 2026-01-01 | null | 2026-01-01 | Permit-application-date rule; standards effective locally when permit application is submitted apply. | `src:usa-ca:cbsc-codes`; `src:usa-ca:hsc-18938-5`; `src:usa-ca:icc-cbc-2025` |
| Residential | 2025 California Residential Code, Title 24 Part 2.5 | 2024 International Residential Code | 2025 | current statewide Title 24 edition | null | 2026-01-01 | null | 2026-01-01 | Permit-application-date rule; HCD residential scope requires local enforcement/follow-up by occupancy. | `src:usa-ca:cbsc-codes`; `src:usa-ca:hsc-18938-5`; `src:usa-ca:hcd-shl-laws`; `src:usa-ca:icc-crc-2025` |
| Existing Building / Rehabilitation | 2025 California Existing Building Code, Title 24 Part 10 | 2024 International Existing Building Code | 2025 | current statewide Title 24 edition | null | 2026-01-01 | null | 2026-01-01 | Permit-application-date rule; local amendment applicability may also depend on local ordinance effective date. | `src:usa-ca:cbsc-codes`; `src:usa-ca:hsc-18938-5`; `src:usa-ca:saclaw-cebc-2025` |
| Mechanical | 2025 California Mechanical Code, Title 24 Part 4 | 2024 Uniform Mechanical Code | 2025 | current statewide Title 24 edition | null | 2026-01-01 | null | 2026-01-01 | Permit-application-date rule. | `src:usa-ca:cbsc-codes`; `src:usa-ca:hsc-18938-5`; `src:usa-ca:saclaw-cmc-2025`; `src:usa-ca:icc-store-cmc-2025` |
| Plumbing | 2025 California Plumbing Code, Title 24 Part 5 | 2024 Uniform Plumbing Code | 2025 | current statewide Title 24 edition | null | 2026-01-01 | null | 2026-01-01 | Permit-application-date rule. | `src:usa-ca:cbsc-codes`; `src:usa-ca:hsc-18938-5`; `src:usa-ca:saclaw-cpc-2025` |
| Fuel Gas | no separate statewide fuel-gas code confirmed | unresolved; likely integrated through mechanical/plumbing provisions | unresolved | unresolved scope | null | null | null | null | Confirm whether fuel-gas requirements are administered through Part 4, Part 5, or another code path before normalizing as a separate adoption. | `src:usa-ca:cbsc-codes` |
| Electrical | 2025 California Electrical Code, Title 24 Part 3 | 2023 National Electrical Code | 2025 | current statewide Title 24 edition | null | 2026-01-01 | null | 2026-01-01 | Permit-application-date rule. | `src:usa-ca:cbsc-codes`; `src:usa-ca:hsc-18938-5`; `src:usa-ca:saclaw-cec-2025`; `src:usa-ca:icc-store-cec-2025` |
| Energy | 2025 Building Energy Efficiency Standards / 2025 California Energy Code, Title 24 Part 6 | California energy standards, not a simple model-code adoption | 2025 | current statewide Title 24 edition | 2024-09-11 | 2026-01-01 | null | 2026-01-01 | CEC states buildings with permit applications applied for on or after 2026-01-01 must comply. | `src:usa-ca:cec-energy-adoption-news`; `src:usa-ca:cec-2025-energy`; `src:usa-ca:prc-25402-1` |
| Fire - construction references | 2025 California Fire Code, Title 24 Part 9 | 2024 International Fire Code | 2025 | current statewide Title 24 edition | null | 2026-01-01 | null | 2026-01-01 | Permit-application-date rule; OSFM/local fire enforcement applies by occupancy and jurisdiction. | `src:usa-ca:cbsc-codes`; `src:usa-ca:hsc-18938-5`; `src:usa-ca:osfm-title24`; `src:usa-ca:icc-cfc-2025` |
| Fire - operational / prevention code | OSFM fire-prevention and fire/panic-safety regulations, including California Fire Code where applicable | California Fire Code / OSFM regulations | 2025 / Title 19 unresolved | partially_verified | null | 2026-01-01 | null | 2026-01-01 | Construction code follows Title 24 permit-date rule; non-building-standard Title 19 operational rules require separate extraction. | `src:usa-ca:hsc-13143`; `src:usa-ca:hsc-13145`; `src:usa-ca:osfm-title24`; `src:usa-ca:osfm-title19` |
| Accessibility | California Building Code accessibility provisions, including DSA Access Compliance amendments | California amendments aligned with federal minimum accessibility requirements where applicable | 2025 | current statewide Title 24 edition; chapter-level parsing incomplete | null | 2026-01-01 | null | 2026-01-01 | Permit-application-date rule; accessibility update exceptions to residential moratorium noted in HSC § 18930. | `src:usa-ca:hsc-18930`; `src:usa-ca:hsc-18938-5`; `src:usa-ca:dsa-access-code-development` |
| Elevator / Conveyance | Title 8 Elevator Safety Orders | California Title 8 safety orders; model-code base not established in this pass | current Title 8 | outside Title 24 core matrix; active rulemaking noted | null | null | null | null | Permit-to-operate and installation inspection rules require Title 8 specialist review; do not treat as a Title 24 adoption without follow-up. | `src:usa-ca:dir-elevator-safety-orders`; `src:usa-ca:dir-elevator-permits`; `src:usa-ca:dir-elevator-group-v-rulemaking` |

### 3.2 Adoption Records

#### `adoption:usa-ca:2025-title24-general`

| Field | Value |
| --- | --- |
| Applies to | Title 24, California Building Standards Code, except where specialized rules apply |
| Adopting / approving authority | `ahj:usa-ca:cbsc` |
| State code | 2025 California Building Standards Code, Title 24 |
| Publication date | 2025-07-01 |
| Effective date | 2026-01-01 |
| Operative date | null |
| Mandatory date | 2026-01-01 |
| Transition rule | Standards approved by CBSC and effective at the local level at permit-application submittal apply to the plans/specifications and construction under that permit. |
| Source IDs | `src:usa-ca:cbsc-codes`; `src:usa-ca:hsc-18938`; `src:usa-ca:hsc-18938-5` |
| Confidence | 0.83 |

#### `adoption:usa-ca:2025-energy-code`

| Field | Value |
| --- | --- |
| Applies to | Title 24 Part 6, Building Energy Efficiency Standards / Energy Code |
| Adopting / approving authority | `ahj:usa-ca:cec` with CBSC approval/publication path |
| State code | 2025 Building Energy Efficiency Standards / 2025 California Energy Code |
| CEC adoption date | 2024-09-11 |
| CBSC approval date | unresolved |
| Effective date | 2026-01-01 |
| Operative date | null |
| Mandatory date | 2026-01-01 |
| Transition rule | Buildings whose permit applications are applied for on or after January 1, 2026 must comply with the 2025 Energy Code. |
| Source IDs | `src:usa-ca:cec-energy-adoption-news`; `src:usa-ca:cec-2025-energy`; `src:usa-ca:prc-25402-1` |
| Confidence | 0.90 |

#### `adoption:usa-ca:2025-title24-fire`

| Field | Value |
| --- | --- |
| Applies to | California Fire Code and OSFM fire/life-safety building standards in Title 24 |
| Adopting / approving authority | `ahj:usa-ca:osfm` proposes/promulgates within scope; `ahj:usa-ca:cbsc` approves/adopts for Title 24 codification |
| State code | 2025 California Fire Code, Title 24 Part 9 |
| Publication date | 2025-07-01 |
| Effective date | 2026-01-01 |
| Operative date | null |
| Mandatory date | 2026-01-01 |
| Transition rule | Construction standards follow Title 24 permit-application-date rule; operational/prevention rules outside Title 24 require Title 19 extraction. |
| Source IDs | `src:usa-ca:cbsc-codes`; `src:usa-ca:osfm-title24`; `src:usa-ca:hsc-13143`; `src:usa-ca:hsc-13145`; `src:usa-ca:icc-cfc-2025` |
| Confidence | 0.74 |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

California separates publication, effective/local applicability, and permit-application triggers. CBSC publishes the California Building Standards Code once every three years and supplements as necessary. HSC § 18938 provides the default 180-day-after-publication effectiveness rule for building standards, unless CBSC establishes a later date. For project applicability, HSC § 18938.5 uses the building permit application date: the standards approved by CBSC and effective at the local level when the application is submitted govern the plans, specifications, and construction under that permit. Energy Code guidance is aligned for 2025: CEC states buildings with permit applications applied for on or after January 1, 2026 must comply with the 2025 Energy Code.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `date-rule:usa-ca:title24-publication-2025` | 2025 Title 24 | publication | 2025-07-01 | 2025 triennial edition publication | n/a | `src:usa-ca:cbsc-codes` | partially_verified |
| `date-rule:usa-ca:title24-effective-2025` | 2025 Title 24 | effective date | 2026-01-01 | CBSC-published effective date | prior code may govern earlier permit applications under permit-date rule | `src:usa-ca:cbsc-codes`; `src:usa-ca:cbsc-home`; `src:usa-ca:hsc-18938` | partially_verified |
| `date-rule:usa-ca:default-180-days` | Title 24 building standards | default effectiveness | 180 days after publication unless later date established | publication by CBSC | yes, for permit applications before local effective date | `src:usa-ca:hsc-18938`; `src:usa-ca:hsc-18941-5` | partially_verified |
| `date-rule:usa-ca:permit-application` | Title 24 building standards | permit application date | application date | building permit application submitted | yes, if earlier standards were effective locally at the time of permit application | `src:usa-ca:hsc-18938-5` | partially_verified |
| `date-rule:usa-ca:local-residential-ordinance` | local residential amendments | local ordinance application | after ordinance effective date | residential building permit application submitted after local ordinance effective date | generally yes for applications before ordinance effective date, subject to exceptions | `src:usa-ca:hsc-18938-5`; `src:usa-ca:hsc-17958-7` | partially_verified |
| `date-rule:usa-ca:energy-2025` | 2025 Energy Code | permit application date | 2026-01-01 onward | permit application applied for on or after 2026-01-01 | yes, for applications before 2026-01-01 unless another rule applies | `src:usa-ca:cec-2025-energy`; `src:usa-ca:prc-25402-1` | partially_verified |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Multiple Title 24 parts | 2025 Intervening Code Adoption Cycle supplements | 2025-10-15 | null | 2027-07-01 | null | null | active_watch | `src:usa-ca:cbsc-2025-intervening-cycle`; `src:usa-ca:cbsc-fall-2025`; `src:usa-ca:cbsc-spring-2026` | CBSC sources identify 2025 intervening-cycle activity for supplements to the 2025 edition; exact affected parts and final effective dates require cycle monitoring. |
| Residential units | temporary statutory moratorium on many state/local residential building standards | 2025-06-30 | 2025-06-30 | 2025-10-01 | null | 2031-06-01 end of moratorium period | active_watch | `src:usa-ca:hsc-18930`; `src:usa-ca:hsc-17958`; `src:usa-ca:hsc-17958-5`; `src:usa-ca:hsc-17958-7`; `src:usa-ca:hsc-18941-5` | Affects proposed building standards and local changes applicable to residential units unless statutory exceptions apply. |
| Energy | 2028 Building Energy Efficiency Standards | null | null | null | null | null | active_watch | `src:usa-ca:cec-2025-energy` | CEC page lists 2028 Energy Code pre-rulemaking activity; no adoption record created in this pass. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| `applicability-rule:usa-ca:model-home-design` | statewide and local building standards | residential dwellings based on an approved model home design in the same jurisdiction | model home design approved under standards in effect at permit application date | State/local standards in effect at the time of the model home permit application continue to apply to future dwellings based on that approved model design unless the design substantially changes or 10 years pass. | `src:usa-ca:hsc-18938-5` | partially_verified |
| `applicability-rule:usa-ca:additional-occupancy-model-code` | model code made applicable to additional occupancy | project submitted before effective date of model code | building permit submittal before model code effective date | The additional-occupancy model code does not apply to a project submitted before the effective date. | `src:usa-ca:hsc-18938-5` | partially_verified |
| `applicability-rule:usa-ca:fire-district-more-stringent` | fire/panic-safety local standards | fire protection district ordinance | district adoption plus city/county ratification | A fire protection district may adopt more stringent fire/panic-safety building standards; ordinance is not effective until ratified by the city/county/city-and-county where it applies. | `src:usa-ca:hsc-13869-7` | partially_verified |
| `applicability-rule:usa-ca:energy-local-fallback` | Energy Code | jurisdiction lacking local building department or failing to enforce | no local building department, or CEC notice after failure to enforce | CEC enforces where there is no local building department and may provide enforcement after notice if a local building department fails to enforce. | `src:usa-ca:prc-25402-1` | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** centralized Title 24 codification with decentralized subject-matter proposing/adopting agencies.

**Where amendments are published:** California Building Standards Code, Title 24, California Code of Regulations; CBSC-published triennial editions, errata, supplements, and rulemaking documents; agency rulemaking pages for pending or proposed changes.

**Amendment parsing status:** registry_level_only. This report identifies the statewide publication/adoption framework and current code edition, but it does not parse California amendments section by section.

California state agencies propose or adopt standards within statutory authority and submit them to CBSC. CBSC reviews those standards under statutory criteria and codifies/publishes approved standards in Title 24. OSFM standards promoting fire and panic safety require State Fire Marshal approval unless adopted by OSFM. HSC § 18930 and HSC § 18942 now also include temporary limits on proposed building standards affecting residential units from October 1, 2025 through June 1, 2031, subject to listed exceptions.

### 5.2 State Amendment Sources

| Amendment Source ID | Scope | Publication Path | Parsed? | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- |
| `amend-source:usa-ca:title24-2025` | 2025 California Building Standards Code, all Title 24 parts | CBSC code portal and linked publishers | partially | `src:usa-ca:cbsc-codes`; `src:usa-ca:cbsc-home` | Edition-level dates captured; section-level amendments not parsed. |
| `amend-source:usa-ca:2025-title24-code-changes` | 2025 substantive California code changes summary | CBSC resource page | not_started | `src:usa-ca:cbsc-2025-code-changes` | Official summary located, but full content fetch was blocked by site access; treat as a watch target and re-fetch manually if necessary. |
| `amend-source:usa-ca:osfm-title24` | fire and life-safety amendments across Title 24 parts | OSFM Title 24 Development page and CBSC rulemaking dockets | partially | `src:usa-ca:osfm-title24`; `src:usa-ca:hsc-13143` | Authority and active-cycle posture captured; individual code sections not parsed. |
| `amend-source:usa-ca:cec-energy-2025` | 2025 Energy Code | CEC 2025 Energy Code page, docket, rulemaking documents, Title 24 Part 6 | partially | `src:usa-ca:cec-2025-energy`; `src:usa-ca:cec-energy-adoption-news`; `src:usa-ca:prc-25402-1` | Adoption/effective/permit-application trigger captured; measure-by-measure amendments not parsed. |
| `amend-source:usa-ca:local-ordinance-filings` | local city/county/fire-district amendments | CBSC local ordinance filing registry | partially | `src:usa-ca:cbsc-local-ordinances`; `src:usa-ca:cbsc-2025-ordinances` | Registry existence captured; no municipality-specific amendments parsed. |

### 5.3 High-Impact State Amendments

| Record ID | Code Family | Summary | Source IDs | Status |
| --- | --- | --- | --- | --- |
| `state-amendment:usa-ca:ab130-residential-moratorium` | multiple residential-unit building standards | Statutory changes effective June 30, 2025 limit consideration/adoption of many state and local building standards affecting residential units from October 1, 2025 through June 1, 2031, unless an exception applies. | `src:usa-ca:hsc-18930`; `src:usa-ca:hsc-17958`; `src:usa-ca:hsc-17958-5`; `src:usa-ca:hsc-17958-7`; `src:usa-ca:hsc-18941-5` | partially_verified |
| `state-amendment:usa-ca:energy-2025-high-level` | Energy | CEC describes the 2025 Energy Code as expanding heat pump use in newly constructed residential buildings, encouraging electric readiness, strengthening ventilation standards, and applying to permit applications on or after January 1, 2026. | `src:usa-ca:cec-2025-energy`; `src:usa-ca:cec-energy-adoption-news` | partially_verified |
| `state-amendment:usa-ca:osfm-wui-exception` | fire / WUI | HSC § 18930 includes an exception for State Fire Marshal amendments to the California Wildland-Urban Interface Code during the residential standards moratorium period. | `src:usa-ca:hsc-18930`; `src:usa-ca:osfm-title24` | partially_verified |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-ca"
  model: "hybrid_statewide_code_local_enforcement_with_state_reserved_scopes"
  enforcing_entities:
    - "city building departments"
    - "county building departments"
    - "city/county housing or health departments for maintenance, sanitation, use, or occupancy where applicable"
    - "local fire chiefs, county fire chiefs, city-and-county fire chiefs, fire districts, and their authorized representatives for fire/panic safety within respective areas"
    - "Office of the State Fire Marshal for state-owned, specified state-occupied, and state institution fire/life-safety scopes"
    - "California Energy Commission where no local building department exists or after failure-to-enforce notice"
    - "DIR/DOSH Elevator Unit for elevator permits and Title 8 Elevator Safety Orders"
  required_officials:
    - "local building official or building department for building permits and code administration"
    - "local fire chief or authorized representative for fire/panic-safety enforcement where applicable"
  state_reserved_activities:
    - "CBSC approval/adoption/codification/publication of statewide building standards"
    - "CEC technical assistance, dispute resolution, fallback Energy Code enforcement, and state-building energy review path"
    - "OSFM plan review/inspection for state-owned, specified state-occupied, and state institution projects"
    - "DIR/DOSH elevator operating permits and Elevator Safety Orders"
  source_ids:
    - "src:usa-ca:hsc-17960"
    - "src:usa-ca:hsc-17961"
    - "src:usa-ca:prc-25402-1"
    - "src:usa-ca:hsc-13145"
    - "src:usa-ca:osfm-fire-life-safety"
    - "src:usa-ca:dir-elevator-safety-orders"
    - "src:usa-ca:dir-elevator-permits"
  verification_status: "partially_verified"
  confidence: 0.76
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-ca"
  model: "local_modifications_allowed_with_findings_filing_and_residential_moratorium_constraints"
  applies_to_code_families:
    - "building"
    - "residential"
    - "green building"
    - "fire and panic safety where adopted by fire protection districts under HSC 13869.7"
    - "other Title 24 building standards where authorized by statute"
  approval_required: "not generally CBSC approval, but filing and statutory compliance are required; fire district ordinances require city/county ratification"
  approving_authority_id: "city_or_county_governing_body_or_fire_district_with_ratification; CBSC may reject certain filings without required findings"
  filing_required: true
  registry_exists: true
  registry_source_ids:
    - "src:usa-ca:cbsc-local-ordinances"
    - "src:usa-ca:cbsc-2025-ordinances"
  legal_basis_source_ids:
    - "src:usa-ca:hsc-17958"
    - "src:usa-ca:hsc-17958-5"
    - "src:usa-ca:hsc-17958-7"
    - "src:usa-ca:hsc-18941-5"
    - "src:usa-ca:hsc-13869-7"
  verification_status: "partially_verified"
  confidence: 0.80
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement and local amendment authority are separate questions in California. A city or county building department may be responsible for enforcing statewide codes within its jurisdiction, but local amendments require a statutory path: express findings, filing, and for some fire district rules city/county ratification. The 2025 residential-unit moratorium further limits both state and local building-standard changes affecting residential units during the October 1, 2025 to June 1, 2031 period unless an exception applies.

### 6.4 Known Local Amendment Registries

| Registry ID | Registry Name | Maintainer | Coverage | Source IDs | Status | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| `registry:usa-ca:cbsc-local-ordinances` | Local Amendments to Building Standards - Ordinances | California Building Standards Commission | Local amendments to building standards filed with CBSC, selectable by code year | `src:usa-ca:cbsc-local-ordinances`; `src:usa-ca:cbsc-2025-ordinances` | partially_verified | Search snippets confirm registry existence; full content access returned 403 during this pass. |

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendments were parsed in this pass. Use the CBSC local ordinance registry before making project-level determinations.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: not_started

Jurisdiction stack:

```text
Address
  -> State of California
  -> County
  -> Municipality / unincorporated county
  -> Fire protection district or municipal/county fire department, if applicable
  -> Special districts, if applicable
  -> Local building department / building official
  -> Local fire chief / fire marshal or OSFM for reserved state scopes
  -> Trade-specific AHJs, including energy compliance and elevator/conveyance where applicable
  -> Applicable 2025 Title 24 statewide code adoption records
  -> Applicable local amendment records from CBSC registry and local ordinance sources
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | California state boundary source not selected | none | statewide | unknown | pending |
| County | California county boundary source not selected | none | statewide | unknown | pending |
| Municipality | California city/place boundary source not selected | none | statewide | unknown | pending |
| Fire District | Fire district boundary source not selected | none | statewide | unknown | pending |
| Special District | Special district boundary source not selected | none | statewide | unknown | pending |

### 7.3 AHJ Contact Data

No AHJ contact data was populated for this pass. Use local building department, local fire authority, OSFM, CEC, and DIR/DOSH contact datasets in a future jurisdiction-specific pass.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Publisher / Maintainer | Type | URL | Accessed | Key Fields Supported |
| --- | --- | --- | --- | --- | --- | --- |
| `src:usa-ca:hsc-18930` | Health and Safety Code § 18930 | California Legislative Information | statute | https://leginfo.legislature.ca.gov/faces/codes_displaySection.xhtml?lawCode=HSC&sectionNum=18930. | 2026-06-25 | CBSC approval/adoption before codification; State Fire Marshal approval for fire/panic standards; residential standards moratorium exceptions |
| `src:usa-ca:hsc-18938` | Health and Safety Code § 18938 | California Legislative Information | statute | https://leginfo.legislature.ca.gov/faces/codes_displaySection.xhtml?lawCode=HSC&sectionNum=18938. | 2026-06-25 | filing, codification, statewide applicability, 180-day default effective date |
| `src:usa-ca:hsc-18938-5` | Health and Safety Code § 18938.5 | California Legislative Information | statute | https://leginfo.legislature.ca.gov/faces/codes_displaySection.xhtml?lawCode=HSC&sectionNum=18938.5. | 2026-06-25 | permit-application-date rule; local residential ordinance applicability; model-home design rule |
| `src:usa-ca:hsc-18941-5` | Health and Safety Code § 18941.5 | California Legislative Information | statute | https://leginfo.legislature.ca.gov/faces/codes_displaySection.xhtml?lawCode=HSC&sectionNum=18941.5. | 2026-06-25 | local amendments, 180-day rule, more-restrictive local standards, residential moratorium |
| `src:usa-ca:hsc-18942` | Health and Safety Code § 18942 | California Legislative Information | statute | https://leginfo.legislature.ca.gov/faces/codes_displaySection.xhtml?lawCode=HSC&sectionNum=18942. | 2026-06-25 | CBSC triennial code publication; supplements; local code copies maintained by building officials |
| `src:usa-ca:hsc-17922` | Health and Safety Code § 17922 | California Legislative Information | statute | https://leginfo.legislature.ca.gov/faces/codes_displaySection.xhtml?lawCode=HSC&sectionNum=17922 | 2026-06-25 | HCD residential/housing standards and model-code families referenced for housing |
| `src:usa-ca:hsc-17958` | Health and Safety Code § 17958 | California Legislative Information | statute | https://leginfo.legislature.ca.gov/faces/codes_displaySection.xhtml?lawCode=HSC&sectionNum=17958 | 2026-06-25 | city/county changes to Title 24/HCD provisions; 2025-2031 residential constraints |
| `src:usa-ca:hsc-17958-5` | Health and Safety Code § 17958.5 | California Legislative Information | statute | https://leginfo.legislature.ca.gov/faces/codes_displaySection.xhtml?lawCode=HSC&sectionNum=17958.5 | 2026-06-25 | local climatic/geological/topographical findings basis; residential restrictions |
| `src:usa-ca:hsc-17958-7` | Health and Safety Code § 17958.7 | California Legislative Information | statute | https://leginfo.legislature.ca.gov/faces/codes_displaySection.xhtml?lawCode=HSC&sectionNum=17958.7 | 2026-06-25 | local findings and filing requirement; CBSC rejection authority for missing findings; residential restrictions |
| `src:usa-ca:hsc-17960` | Health and Safety Code § 17960 | California Legislative Information | statute | https://leginfo.legislature.ca.gov/faces/codes_displaySection.xhtml?lawCode=HSC&sectionNum=17960. | 2026-06-25 | city/county building department enforcement for covered residential/housing construction activities |
| `src:usa-ca:hsc-17961` | Health and Safety Code § 17961 | California Legislative Information | statute | https://leginfo.legislature.ca.gov/faces/codes_displaySection.xhtml?lawCode=HSC&sectionNum=17961. | 2026-06-25 | local enforcement for maintenance, sanitation, ventilation, use, or occupancy of covered housing occupancies |
| `src:usa-ca:hsc-13143` | Health and Safety Code § 13143 | California Legislative Information | statute | https://leginfo.legislature.ca.gov/faces/codes_displaySection.xhtml?lawCode=HSC&sectionNum=13143. | 2026-06-25 | State Fire Marshal fire/panic-safety building standards and regulations |
| `src:usa-ca:hsc-13145` | Health and Safety Code § 13145 | California Legislative Information | statute | https://leginfo.legislature.ca.gov/faces/codes_displaySection.xhtml?lawCode=HSC&sectionNum=13145. | 2026-06-25 | State Fire Marshal and local fire chiefs/districts enforce fire/panic-safety standards in respective areas |
| `src:usa-ca:hsc-13869-7` | Health and Safety Code § 13869.7 | California Legislative Information | statute | https://leginfo.legislature.ca.gov/faces/codes_displaySection.xhtml?lawCode=HSC&sectionNum=13869.7. | 2026-06-25 | fire protection district more-stringent fire/panic-safety building standards and ratification process |
| `src:usa-ca:prc-25402-1` | Public Resources Code § 25402.1 | California Legislative Information | statute | https://leginfo.legislature.ca.gov/faces/codes_displaySection.xhtml?lawCode=PRC&sectionNum=25402.1. | 2026-06-25 | CEC Energy Code implementation, manuals, local enforcement, CEC fallback enforcement and disputes |
| `src:usa-ca:cbsc-home` | Building Standards Commission home page | California Department of General Services / CBSC | agency page | https://www.dgs.ca.gov/bsc | 2026-06-25 | 2025 Title 24 current effective period, 2025 intervening-cycle notice |
| `src:usa-ca:cbsc-about` | CBSC About | California Department of General Services / CBSC | agency page | https://www.dgs.ca.gov/BSC/About | 2026-06-25 | CBSC charged by Building Standards Law to administer adoption, approval, publication, and implementation processes; local ordinance filing role |
| `src:usa-ca:cbsc-codes` | Codes | California Department of General Services / CBSC | agency page | https://www.dgs.ca.gov/bsc/codes | 2026-06-25 | 2025 Title 24 publication date and effective date; links to code parts |
| `src:usa-ca:cbsc-2025-code-changes` | 2025 Title 24 California Code Changes | California Department of General Services / CBSC | agency page | https://www.dgs.ca.gov/BSC/Resources/2025-Title-24-California-Code-Changes | 2026-06-25 | summary of substantive 2025 California changes; official rulemaking context |
| `src:usa-ca:cbsc-2025-intervening-cycle` | 2025 Intervening Code Adoption Cycle | California Department of General Services / CBSC | agency page | https://www.dgs.ca.gov/BSC/Rulemaking/2025-Intervening-Cycle | 2026-06-25 | pending supplement cycle and agency submission timing |
| `src:usa-ca:cbsc-fall-2025` | CALCode Quarterly Fall 2025 | California Department of General Services / CBSC | agency newsletter | https://www.dgs.ca.gov/BSC/News/Page-Content/CALCode-Quarterly-Newsletters-Archive/CALCode-Quarterly-Fall-2025 | 2026-06-25 | local filing guidance and intervening-cycle context |
| `src:usa-ca:cbsc-spring-2026` | CALCode Quarterly Spring 2026 | California Department of General Services / CBSC | agency newsletter | https://www.dgs.ca.gov/BSC/News/Page-Content/CALCode-Quarterly-Newsletters-Archive/CALCode-Quarterly-Spring-2026 | 2026-06-25 | supplements to 2025 edition watch target |
| `src:usa-ca:cbsc-local-ordinances` | Local Amendments to Building Standards - Ordinances | California Department of General Services / CBSC | registry page | https://www.dgs.ca.gov/bsc/codes/local-amendments-to-building-standards---ordinances | 2026-06-25 | local ordinance filing registry existence |
| `src:usa-ca:cbsc-2025-ordinances` | 2025 Ordinances | California Department of General Services / CBSC | registry page | https://www.dgs.ca.gov/BSC/Codes/2025-Ordinances | 2026-06-25 | 2025 local ordinance search by city/county |
| `src:usa-ca:hcd-title24` | Title 24 Rulemaking | California Department of Housing and Community Development | agency page | https://www.hcd.ca.gov/building-standards/title-24-rulemaking | 2026-06-25 | HCD Title 24 building standards development and housing accessibility scope |
| `src:usa-ca:hcd-shl-laws` | State Housing Law Program Laws & Regulations | California Department of Housing and Community Development | agency page | https://www.hcd.ca.gov/building-standards/shl/laws-and-regulations | 2026-06-25 | statewide residential building standards in Title 24; CBSC adoption/publication |
| `src:usa-ca:cec-2025-energy` | 2025 Building Energy Efficiency Standards | California Energy Commission | agency page | https://www.energy.ca.gov/programs-and-topics/programs/building-energy-efficiency-standards/2025-building-energy-efficiency | 2026-06-25 | 2025 Energy Code scope, permit-application trigger, compliance resources, local ordinances exceeding Energy Code |
| `src:usa-ca:cec-energy-adoption-news` | Energy Commission Adopts Updated Building Standards | California Energy Commission | agency news release | https://www.energy.ca.gov/news/2024-09/energy-commission-adopts-updated-building-standards-expanding-requirements-heat | 2026-06-25 | CEC adoption date and CBSC approval path for 2025 Energy Code |
| `src:usa-ca:osfm-title24` | Title 24 Development | Office of the State Fire Marshal | agency page | https://osfm.fire.ca.gov/what-we-do/code-development-and-analysis/title-24-development | 2026-06-25 | OSFM Title 24 fire/life-safety regulation role and affected codes |
| `src:usa-ca:osfm-title19` | Title 19 Development | Office of the State Fire Marshal | agency page | https://osfm.fire.ca.gov/what-we-do/code-development-and-analysis/title-19-development | 2026-06-25 | operational fire/prevention rulemaking watch target |
| `src:usa-ca:osfm-fire-life-safety` | Fire and Life Safety | Office of the State Fire Marshal | agency page | https://osfm.fire.ca.gov/what-we-do/fire-and-life-safety | 2026-06-25 | OSFM enforcement jurisdiction for state-owned, specified state-occupied, and state institutions; plan review/inspection |
| `src:usa-ca:dsa-access-code-development` | Access Compliance Code Development | Division of the State Architect | agency page | https://www.dgs.ca.gov/DSA/Resources/Access-Compliance-Code-Development | 2026-06-25 | DSA accessibility code-change proposals submitted to CBSC |
| `src:usa-ca:dir-elevator-safety-orders` | Title 8 Elevator Safety Orders | Department of Industrial Relations | regulation page | https://www.dir.ca.gov/title8/sub6.html | 2026-06-25 | Elevator Safety Orders scope and caveat |
| `src:usa-ca:dir-elevator-permits` | Elevator Permits | Department of Industrial Relations / Cal/OSHA | agency page | https://www.dir.ca.gov/dosh/ElevatorPermits.html | 2026-06-25 | Cal/OSHA Elevator Unit operating permits |
| `src:usa-ca:dir-elevator-group-v-rulemaking` | Elevator Safety Orders - Group V | Occupational Safety and Health Standards Board | rulemaking page | https://www.dir.ca.gov/OSHSB/ESO-Group-V.html | 2026-06-25 | active elevator rulemaking watch target |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| `src:usa-ca:cbsc-codes` | official_page_with_publisher_links | CBSC code page is official for edition publication/effective date and directs users to publisher sites for code text; it does not itself provide all model-code base details in parsed text. | Use for edition/status/date fields; use publisher/catalog sources for base-model details until official code text is extracted. |
| `src:usa-ca:cbsc-about` | full_fetch_blocked | Search result establishes the relevant CBSC role; full page fetch returned 403 during this pass. | Treat as official but re-fetch manually or via browser before verified status. |
| `src:usa-ca:cbsc-local-ordinances` | full_fetch_blocked | Search result establishes registry existence; full page fetch returned 403 during this pass. | Use only for registry existence, not for municipality-specific amendment content. |
| `src:usa-ca:cbsc-2025-ordinances` | full_fetch_blocked | Search result establishes 2025 ordinance search by city/county; full page fetch returned 403 during this pass. | Use as watch/navigation source; do not infer local amendment contents. |
| `src:usa-ca:cbsc-2025-code-changes` | full_fetch_blocked | Official summary page located, but full page fetch returned 403 during this pass. | Use as watch target; do not treat amendment list as parsed. |
| `src:usa-ca:dir-elevator-safety-orders` | agency_convenience_copy | DIR Title 8 pages state they are provided free of charge and include a current/accuracy disclaimer. | Use as a working official agency copy; confirm against official CCR if needed for final legal determinations. |
| `src:usa-ca:osfm-title19` | scope_not_parsed | Title 19 operational fire rulemaking exists, but current operational code text was not parsed. | Keep operational fire-code details partially unresolved. |

### 8.3 Supplemental Sources

| Source ID | Title | Publisher / Maintainer | Type | URL | Accessed | Key Fields Supported | Caveat |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `src:usa-ca:icc-cbc-2025` | 2025 California Building Code, Title 24 Part 2 | ICC Digital Codes | publisher page | https://codes.iccsafe.org/content/CABC2025P2 | 2026-06-25 | CBC title, Part 2, base 2024 IBC, errata note | Publisher page; not an agency statute/regulation page. |
| `src:usa-ca:icc-crc-2025` | 2025 California Residential Code, Title 24 Part 2.5 | ICC Digital Codes | publisher page | https://codes.iccsafe.org/content/CARC2025P2 | 2026-06-25 | CRC title, Part 2.5, base IRC context, errata note | Publisher page; use for base edition only until official code text is extracted. |
| `src:usa-ca:icc-cfc-2025` | 2025 California Fire Code, Title 24 Part 9 | ICC Digital Codes | publisher page | https://codes.iccsafe.org/content/CAFC2025P2 | 2026-06-25 | CFC title, Part 9, base 2024 IFC, errata note | Publisher page; not a replacement for OSFM/CBSC authority sources. |
| `src:usa-ca:saclaw-cebc-2025` | 2025 California Existing Building Code catalog record | Sacramento County Public Law Library | catalog record | https://catalog.saclaw.org/cgi-bin/koha/opac-ISBDdetail.pl?biblionumber=10551 | 2026-06-25 | CEBC title, Part 10, base 2024 IEBC, effective date | Library catalog; use only for bibliographic/base-code details. |
| `src:usa-ca:saclaw-cmc-2025` | 2025 California Mechanical Code catalog record | Sacramento County Public Law Library | catalog record | https://catalog.saclaw.org/cgi-bin/koha/opac-ISBDdetail.pl?biblionumber=10585 | 2026-06-25 | CMC title, Part 4, base 2024 UMC, effective date | Library catalog; use only for bibliographic/base-code details. |
| `src:usa-ca:saclaw-cpc-2025` | 2025 California Plumbing Code catalog record | Sacramento County Public Law Library | catalog record | https://catalog.saclaw.org/cgi-bin/koha/opac-ISBDdetail.pl?biblionumber=10586 | 2026-06-25 | CPC title, Part 5, base 2024 UPC, effective date | Library catalog; use only for bibliographic/base-code details. |
| `src:usa-ca:saclaw-cec-2025` | 2025 California Electrical Code catalog record | Sacramento County Public Law Library | catalog record | https://catalog.saclaw.org/bib/10571 | 2026-06-25 | CEC title, Part 3, base 2023 NEC, effective date | Library catalog; use only for bibliographic/base-code details. |
| `src:usa-ca:icc-store-cec-2025` | 2025 California Electrical Code, Title 24 Part 3 | ICC Store | publisher/store page | https://shop.iccsafe.org/2025-california-electrical-code-title-24-part-3.html | 2026-06-25 | Part 3, 2023 NEC base, effective date | Commercial publisher/store page. |
| `src:usa-ca:icc-store-cmc-2025` | 2025 California Mechanical Code, Title 24 Part 4 | ICC Store | publisher/store page | https://shop.iccsafe.org/2025-california-mechanical-code-title-24-part-4.html | 2026-06-25 | Part 4, 2024 UMC base, effective date | Commercial publisher/store page. |

### 8.4 Source Extraction Metadata

| Extraction ID | Date | Method | Sources Covered | Notes |
| --- | --- | --- | --- | --- |
| `extract:usa-ca:2026-06-25-web` | 2026-06-25 | targeted web research of official California statutes and agency pages, plus publisher/catalog sources for base-code bibliographic details | all sources in Sections 8.1 and 8.3 | Several DGS pages returned 403 on full fetch; search-result snippets and other official pages were used conservatively. |
| `extract:usa-ca:2026-06-25-template` | 2026-06-25 | populated from uploaded California draft stub | uploaded draft | The uploaded draft was a low-confidence baseline with unresolved placeholders; this file replaces those placeholders with sourced California-specific content. |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 1.00 | none | Body now includes official source registry and source-backed core fields; full verification still requires section-by-section code extraction. |
| report | risk.overall_confidence | 0.62 | verified | 1.00 | none | Confidence reflects strong authority/date support and weaker base-code/amendment parsing. |
| `ahj:usa-ca:cbsc` | authority name and role | California Building Standards Commission; Title 24 approval/adoption/publication hub | partially_verified | 0.86 | `src:usa-ca:hsc-18930`; `src:usa-ca:hsc-18938`; `src:usa-ca:hsc-18942`; `src:usa-ca:cbsc-about` | Legal role is supported by statutes; agency about page full fetch was blocked. |
| `adoption:usa-ca:2025-title24-general` | publication date | 2025-07-01 | partially_verified | 0.84 | `src:usa-ca:cbsc-codes` | CBSC code page supports publication date. |
| `adoption:usa-ca:2025-title24-general` | effective date | 2026-01-01 | partially_verified | 0.86 | `src:usa-ca:cbsc-codes`; `src:usa-ca:cbsc-home`; `src:usa-ca:hsc-18938` | Supported by CBSC and statute. |
| `date-rule:usa-ca:permit-application` | trigger | permit application submitted | partially_verified | 0.88 | `src:usa-ca:hsc-18938-5` | Strong statutory support. |
| `adoption:usa-ca:2025-energy-code` | CEC adoption date | 2024-09-11 | partially_verified | 0.90 | `src:usa-ca:cec-energy-adoption-news` | CEC adoption supported; CBSC approval date remains unresolved. |
| `adoption:usa-ca:2025-energy-code` | mandatory date | 2026-01-01 | partially_verified | 0.90 | `src:usa-ca:cec-2025-energy` | CEC states permit applications on or after this date must comply. |
| `local-enforcement:usa-ca` | model | hybrid_statewide_code_local_enforcement_with_state_reserved_scopes | partially_verified | 0.76 | `src:usa-ca:hsc-17960`; `src:usa-ca:hsc-17961`; `src:usa-ca:prc-25402-1`; `src:usa-ca:hsc-13145` | Coverage is verified for housing, energy, and fire scopes; full commercial building enforcement requires additional extraction. |
| `local-amendment-rule:usa-ca` | filing_required | true | partially_verified | 0.83 | `src:usa-ca:hsc-17958-7`; `src:usa-ca:cbsc-local-ordinances` | Filing requirement supported; registry content not parsed. |
| `ahj:usa-ca:dir-dosh-elevator` | elevator authority | DIR/DOSH Elevator Unit and Title 8 Elevator Safety Orders | partially_verified | 0.66 | `src:usa-ca:dir-elevator-safety-orders`; `src:usa-ca:dir-elevator-permits` | Needs specialist reconciliation with Title 24 elevator references and official CCR. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Body source IDs are listed in Section 8. |
| All authority IDs resolve | pass | Authority IDs used in Sections 2 and 6 are either defined or explicitly local/unresolved pseudo-authorities. |
| All current code families have adoption rows | pass | Rows are present; Fuel Gas and Elevator/Conveyance remain explicit unresolved/outside-core rows. |
| Building and operational fire code are separated | pass | Title 24 construction fire code and Title 19/operational fire-prevention scope are separated. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Fields remain separate; unknown adoption/operative dates are null. |
| Effective dates are valid ISO dates | pass | Entered date fields use ISO format. |
| No impossible date sequences | pass | 2025-07-01 publication followed by 2026-01-01 effectiveness; energy CEC adoption precedes effectiveness. |
| Transition rules have explicit trigger conditions | pass | Permit-application trigger and 180-day publication rule captured. |
| Permit-date logic is captured where applicable | pass | HSC § 18938.5 and CEC permit-application Energy Code rule captured. |
| Local enforcement model classified | pass | Hybrid local/state-reserved model classified. |
| Local amendment rule classified | pass | Findings/filing/moratorium model classified. |
| AHJ confirmation metadata present | fail | Jurisdiction-specific AHJ contacts and boundary sources not populated. |
| Official-source caveats captured | pass | DGS 403, publisher-link, DIR caveats captured. |
| Section-level state amendments parsed | fail | Amendment sources identified, but not parsed section-by-section. |
| Base model code editions official-only | fail | Some base model details rely on publisher/catalog sources rather than official state code text extraction. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `issue:usa-ca:001` | medium | CBSC approval dates | Title 24 effective/publication dates are captured, but CBSC approval/adoption dates by code part are not fully normalized. | Extract commission meeting actions and final rulemaking documents for each relevant 2024 Triennial Cycle package. | null | null | open |
| `issue:usa-ca:002` | medium | base model editions | Several base model code details are supported by publisher/catalog sources rather than official state code text. | Extract prefaces or official Title 24 part text from authorized publishers or official code files. | null | null | open |
| `issue:usa-ca:003` | high | state amendments | State amendments are not parsed section-by-section. | Parse 2025 Title 24 change summaries, final express terms, supplements, and errata by code family. | null | null | open |
| `issue:usa-ca:004` | medium | fuel gas scope | No separate California Fuel Gas Code adoption was confirmed. | Confirm whether fuel-gas provisions are administered through Part 4, Part 5, or another authority path. | null | null | open |
| `issue:usa-ca:005` | medium | operational fire code | Title 24 fire construction scope is captured, but Title 19 operational/prevention rules were not extracted. | Review OSFM Title 19 regulations and rulemaking documents. | null | null | open |
| `issue:usa-ca:006` | high | AHJ resolver | Boundary sources and jurisdiction-specific AHJ contacts are not selected. | Choose state/county/municipality/fire-district boundary datasets and contact sources. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| `watch:usa-ca:cbsc-codes` | `src:usa-ca:cbsc-codes` | html_diff | monthly | new errata, supplement, code edition, or link restructuring | 2026-06-25 | active |
| `watch:usa-ca:cbsc-2025-intervening` | `src:usa-ca:cbsc-2025-intervening-cycle` | html_diff | monthly | new 2025 supplement documents, commission actions, or effective dates | 2026-06-25 | active |
| `watch:usa-ca:cbsc-local-ordinances` | `src:usa-ca:cbsc-local-ordinances` | registry_diff | monthly | new or modified local ordinance filings | 2026-06-25 | active |
| `watch:usa-ca:cec-2025-energy` | `src:usa-ca:cec-2025-energy` | html_diff | monthly | new compliance manuals, local ordinance determinations, software approvals, or errata | 2026-06-25 | active |
| `watch:usa-ca:osfm-title24` | `src:usa-ca:osfm-title24` | html_diff | monthly | new fire/life-safety Title 24 rulemaking or WUI emergency action | 2026-06-25 | active |
| `watch:usa-ca:osfm-title19` | `src:usa-ca:osfm-title19` | html_diff | quarterly | new operational fire-prevention rulemaking | 2026-06-25 | active |
| `watch:usa-ca:dir-elevator` | `src:usa-ca:dir-elevator-group-v-rulemaking` | html_diff | monthly | filed/effective Elevator Safety Orders update | 2026-06-25 | active |
| `watch:usa-ca:hsc-18930` | `src:usa-ca:hsc-18930` | statutory_diff | quarterly | amendment to residential standards moratorium or CBSC approval authority | 2026-06-25 | active |
| `watch:usa-ca:hsc-17958-family` | `src:usa-ca:hsc-17958` | statutory_diff | quarterly | amendment to local amendment authority or moratorium exceptions | 2026-06-25 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-24 | Generated baseline draft report | `report:usa-ca` | none | Codex | Initial uploaded stub had unresolved California-specific content. |
| 2026-06-25 | Populated California authority, code-adoption, date-rule, enforcement, local-amendment, source-registry, QA, and monitoring sections | `ahj:usa-ca:cbsc`; `adoption:usa-ca:2025-title24-general`; `adoption:usa-ca:2025-energy-code`; `local-enforcement:usa-ca`; `local-amendment-rule:usa-ca` | `src:usa-ca:hsc-18930`; `src:usa-ca:cbsc-codes`; `src:usa-ca:cec-2025-energy`; `src:usa-ca:hsc-18938-5` | ChatGPT | Upgraded report status to partially_verified for narrow validation; kept unresolved issues explicit. |
