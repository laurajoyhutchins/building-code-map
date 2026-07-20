---
state:
  state_id: "US-KS"
  name: "Kansas"
  abbreviation: "KS"
report:
  report_id: "state-report:usa-ks"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-25"
  last_verified: null
  reviewed_by: null
risk:
  overall_confidence: 0.62 # 0.00 - 1.00
  risk_flags:
    - "statewide_model_code_adoption_limited_to_fire_life_safety_scope"
    - "mechanical_plumbing_energy_adoptions_unresolved"
    - "local_building_code_adoptions_not_parsed"
    - "ifc_2024_update_rulemaking_pending"
    - "some_revisor_statute_pages_may_lag_2026_kansas_register_text"
  open_questions_count: 6

---

# State Building Code Authority Report: Kansas

## 1. Executive Summary

- **Authority model:** Kansas has a statewide fire/life-safety code model centered on the Office of the State Fire Marshal rather than a confirmed comprehensive statewide building department for every construction-code family. The State Fire Marshal adopts and administers the Kansas fire prevention code, which has uniform statewide force and includes nationally recognized code references, including the 2006 IBC, 2006 IFC, and NFPA standards. Cities and counties retain home-rule/local-code roles, but local rules cannot be inconsistent with the Kansas fire prevention code.

- **Statewide code status:** The verified statewide adoption record in this pass is the Kansas fire prevention code. Current official Kansas State Fire Marshal materials identify the 2006 IBC and 2006 IFC as currently enforced under K.A.R. 22-1-3. Separate verified statewide records also exist for school-building construction references, limited publicly assisted dwelling accessibility standards, and elevator/conveyance safety.

- **Local enforcement model:** Hybrid. The State Fire Marshal has statewide authority for the Kansas fire prevention code, certain plan-review/code-footprint functions, school fire/life-safety oversight, and elevator safety. Local building and fire officials may enforce locally adopted building/fire codes and may collaborate with KSFM through memoranda of understanding, but local AHJ coverage and local model-code editions were not parsed jurisdiction by jurisdiction.

- **Local amendment posture:** Local fire-code amendments are constrained by state law. A municipality that adopts a nationally recognized fire code or the fire-protection segment of a building code and modifies a section must submit a summary of modifications to the State Fire Marshal for approval or rejection. Municipal enactments inconsistent with the Kansas fire prevention code are subject to State Fire Marshal determination. General local building-code amendment authority outside that fire/life-safety scope remains unresolved.

- **Known transition periods or pending changes:** HB 2739, published in the Kansas Register on 2026-04-16, created a special rulemaking path for updating the Kansas fire prevention code with the 2024 International Fire Code, with appropriate amendments, but the source reviewed still requires Attorney General review and legislative ratification of the rules. Elevator ASME A17.1 (2022) applies on and after 2024-07-01, with ASME A17.1 (2019) allowed for elevators designed, approved, or contracted before that date unless the Fire Marshal grants a variance.

- **Production readiness:** narrow_use_only

### Key Findings

```yaml
---
key_findings:
- topic: State fire/life-safety adopting authority
  finding: The State Fire Marshal adopts rules and regulations for safeguarding life
    and property from fire, explosion, and hazardous materials; the Kansas fire prevention
    code has uniform statewide force.
  confidence: 0.85
  source_ids:
  - src:usa-ks:ksreg-2026-hb2739
  - src:usa-ks:kar-22-1-fire-prevention-code
- topic: Primary statewide construction-code reference
  finding: The verified statewide construction-code reference is the 2006 IBC as adopted
    in K.A.R. 22-1-3 and treated as part of the Kansas fire prevention code, not as
    a fully verified all-purpose statewide building code.
  confidence: 0.78
  source_ids:
  - src:usa-ks:kar-22-1-fire-prevention-code
  - src:usa-ks:ksfm-code-listing
  - src:usa-ks:ksfm-plan-review
- topic: Operational fire code
  finding: The verified operational fire code reference is the 2006 IFC with stated
    exclusions in K.A.R. 22-1-3.
  confidence: 0.82
  source_ids:
  - src:usa-ks:kar-22-1-fire-prevention-code
  - src:usa-ks:ksfm-code-listing
- topic: One- and two-family dwellings
  finding: Fire Marshal rules under K.S.A. 31-133 do not apply to buildings used wholly
    as dwelling houses containing no more than two families. No statewide IRC adoption
    was verified.
  confidence: 0.74
  source_ids:
  - src:usa-ks:ksreg-2026-hb2739
- topic: Schools
  finding: School-building construction must meet the IBC adopted by reference in
    K.A.R. 22-1-3, and school electric wiring must conform to the NEC adopted by reference
    in K.A.R. 22-1-3.
  confidence: 0.82
  source_ids:
  - src:usa-ks:ksreg-2024-kar-22-18-3
  - src:usa-ks:ksa-31-150
- topic: Local fire amendments
  finding: Municipal fire-code modifications must be submitted to the State Fire Marshal
    and reviewed for approval or rejection within 30 days.
  confidence: 0.86
  source_ids:
  - src:usa-ks:kar-22-1-fire-prevention-code
- topic: Elevator code
  finding: KSFM implemented ASME A17.1 (2022) effective 2024-07-01; pre-2024-07-01
    designed, approved, or contracted elevators may comply with ASME A17.1 (2019)
    unless a variance applies.
  confidence: 0.86
  source_ids:
  - src:usa-ks:ksfm-asme-a17-1-2022
  - src:usa-ks:ksfm-elevator-program
- topic: Pending fire-code update
  finding: HB 2739 authorizes special handling for rules updating the Kansas fire
    prevention code with the 2024 IFC, but still requires Attorney General review
    and legislative ratification.
  confidence: 0.78
  source_ids:
  - src:usa-ks:ksreg-2026-hb2739
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-ks:state-fire-marshal |
| Authority name | Kansas Office of the State Fire Marshal |
| Authority type | state_fire_marshal / state fire and life-safety code authority |
| Legal basis | K.S.A. 31-133 and 31-134 as amended and published in 2026 HB 2739; K.A.R. Article 22-1 |
| Role | Adopts and administers the Kansas fire prevention code; may incorporate nationally recognized fire-prevention code editions; reviews specified plans/code footprints and local fire-code modifications; administers elevator safety. |
| Enforcement model | statewide fire/life-safety authority with local collaboration and local code administration; not confirmed as a comprehensive statewide building-code agency for all model-code families |
| Source IDs | src:usa-ks:ksreg-2026-hb2739; src:usa-ks:kar-22-1-fire-prevention-code; src:usa-ks:ksfm-plan-review; src:usa-ks:ksfm-elevator-program |
| Verification status | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-ks:state-fire-marshal | Kansas Office of the State Fire Marshal | State fire/life-safety adoption and plan-review reference using IBC; general building-code administration outside this scope remains local/unresolved. | K.S.A. 31-133, 31-134, 31-134a; K.A.R. 22-1-3 | src:usa-ks:ksreg-2026-hb2739; src:usa-ks:ksa-31-134a; src:usa-ks:kar-22-1-fire-prevention-code; src:usa-ks:ksfm-plan-review | partially_verified |
| Residential | ahj:usa-ks:local-jurisdictions-unresolved | Local jurisdictions, where codes are adopted | No statewide IRC adoption verified; one- and two-family dwellings are outside the Fire Marshal rule scope stated in K.S.A. 31-133. | K.S.A. 31-133 as amended by 2026 HB 2739 | src:usa-ks:ksreg-2026-hb2739 | unresolved |
| Existing Building / Rehabilitation | ahj:usa-ks:state-fire-marshal | Kansas Office of the State Fire Marshal | Existing facilities may continue if predating rules and not determined to be a distinct hazard; no IEBC adoption verified. | K.S.A. 31-133(c) | src:usa-ks:ksreg-2026-hb2739 | partially_verified_limited |
| Mechanical | ahj:usa-ks:unknown | Unknown | No statewide mechanical-code adoption verified in this pass. | unresolved | none | unresolved |
| Plumbing | ahj:usa-ks:unknown | Unknown | No statewide plumbing-code adoption verified in this pass. | unresolved | none | unresolved |
| Fuel Gas | ahj:usa-ks:unknown | Unknown | LP-gas rules exist under KSFM, but fuel-gas construction-code adoption was not parsed. | unresolved | none | unresolved |
| Electrical | ahj:usa-ks:state-fire-marshal | Kansas Office of the State Fire Marshal | NFPA 70/NEC is adopted as a Kansas fire prevention code standard and school electric wiring reference; no broader statewide electrical-code regime was parsed. | K.A.R. 22-1-3; K.A.R. 22-18-3 | src:usa-ks:kar-22-1-fire-prevention-code; src:usa-ks:ksreg-2024-kar-22-18-3 | partially_verified_limited |
| Energy | ahj:usa-ks:unknown | Unknown | No statewide energy-code adoption verified in this pass. | unresolved | none | unresolved |
| Fire - construction references | ahj:usa-ks:state-fire-marshal | Kansas Office of the State Fire Marshal | Adopts IBC and IFC references for Kansas fire prevention code compliance and plan-review/code-footprint purposes. | K.A.R. 22-1-2; K.A.R. 22-1-3; K.S.A. 31-134a | src:usa-ks:kar-22-1-fire-prevention-code; src:usa-ks:ksa-31-134a; src:usa-ks:ksfm-plan-review | partially_verified |
| Fire - operational / prevention code | ahj:usa-ks:state-fire-marshal | Kansas Office of the State Fire Marshal | Adopts and administers the Kansas fire prevention code and IFC/NFPA fire-safety standards. | K.S.A. 31-133, 31-134; K.A.R. 22-1-3 | src:usa-ks:ksreg-2026-hb2739; src:usa-ks:kar-22-1-fire-prevention-code; src:usa-ks:ksfm-code-listing | partially_verified |
| Accessibility | ahj:usa-ks:public-assistance-program-admins | Public financial assistance program administrators / Division of Housing reference | Limited state accessibility standards apply to new dwellings constructed with public financial assistance; full enforcement path not parsed. | K.S.A. 58-1401 through 58-1404, as affected by 2026 HB 2739 | src:usa-ks:ksreg-2026-hb2739; src:usa-ks:ksa-58-1402; src:usa-ks:ksa-58-1403; src:usa-ks:ksa-58-1404 | partially_verified_limited |
| Elevator / Conveyance | ahj:usa-ks:state-fire-marshal | Kansas Office of the State Fire Marshal | Administers the Elevator Safety Act, elevator registration, licensing, installation/alteration permits, certificates of operation, and inspections, with city/county program exceptions. | K.S.A. 44-1801 through 44-1815; KSFM elevator program guidance | src:usa-ks:ksa-44-1801; src:usa-ks:ksa-44-1804; src:usa-ks:ksa-44-1814; src:usa-ks:ksa-44-1815; src:usa-ks:ksfm-elevator-program; src:usa-ks:ksfm-asme-a17-1-2022 | partially_verified |

### 2.3 Authority Hierarchy Notes

Kansas should be modeled as a hybrid authority state for this dataset:

1. The Kansas Office of the State Fire Marshal is the statewide fire/life-safety code authority for the Kansas fire prevention code.
2. The Kansas fire prevention code has uniform force and effect statewide, and municipalities cannot enforce inconsistent ordinances or rules.
3. Cities have constitutional home-rule authority over local affairs, and counties have statutory home-rule authority, both subject to uniform statewide enactments and other limitations.
4. Local building-code adoption and enforcement may be the controlling path for model building, residential, mechanical, plumbing, energy, and local administrative details, but those local adoptions were not parsed here.
5. Elevators are a distinct statewide KSFM program, except for city or county programs that meet or exceed state standards; KSFM materials identify Wichita as conducting its own elevator inspections.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-ks:001 | ahj:usa-ks:state-fire-marshal | adopts | Kansas fire prevention code | src:usa-ks:ksreg-2026-hb2739; src:usa-ks:kar-22-1-fire-prevention-code | partially_verified |
| edge:usa-ks:002 | Kansas fire prevention code | has_uniform_force_and_effect | statewide | src:usa-ks:ksreg-2026-hb2739 | partially_verified |
| edge:usa-ks:003 | ahj:usa-ks:state-fire-marshal | determines_inconsistency | municipal enactments inconsistent with Kansas fire prevention code | src:usa-ks:ksreg-2026-hb2739; src:usa-ks:kar-22-1-fire-prevention-code | partially_verified |
| edge:usa-ks:004 | municipalities | submit_modifications_to | ahj:usa-ks:state-fire-marshal for modified fire codes or fire-protection building-code segments | src:usa-ks:kar-22-1-fire-prevention-code | partially_verified |
| edge:usa-ks:005 | ahj:usa-ks:state-fire-marshal | collaborates_with | local fire and building officials through MOUs for plan review and inspections | src:usa-ks:ksfm-plan-review | partially_verified |
| edge:usa-ks:006 | cities | exercise_home_rule | local affairs subject to statewide/uniform enactments and charter-ordinance limits | src:usa-ks:kan-const-art12-sec5 | partially_verified |
| edge:usa-ks:007 | counties | exercise_home_rule | county local legislation and administration subject to statutory limitations | src:usa-ks:ksa-19-101a | partially_verified |
| edge:usa-ks:008 | city_or_county_elevator_programs | may_supersede_state_elevator_permit_or_inspection_sections_if | local requirements meet or exceed Elevator Safety Act and rules | src:usa-ks:ksa-44-1804; src:usa-ks:ksa-44-1814; src:usa-ks:ksa-44-1815 | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | Kansas fire prevention code building-code reference | International Building Code | 2006 | adopted_fire_life_safety_reference | null | 2011-02-04 | null | null | Existing facilities predating rules may continue unless distinct hazard; local building-code editions still require AHJ lookup. | src:usa-ks:kar-22-1-fire-prevention-code; src:usa-ks:ksfm-code-listing; src:usa-ks:ksfm-plan-review |
| Residential | No statewide IRC adoption verified | International Residential Code | unknown | unresolved_local_or_none_statewide | null | null | null | null | K.S.A. 31-133 excludes buildings wholly used as dwelling houses containing no more than two families from the Fire Marshal rule scope. | src:usa-ks:ksreg-2026-hb2739 |
| Existing Building / Rehabilitation | Distinct-hazard continuation rule | no IEBC verified | no IEBC verified | limited_state_rule_only | null | 2026-04-16 | null | null | Facilities in service before effective rules may continue if not a distinct hazard to life or property. | src:usa-ks:ksreg-2026-hb2739 |
| Mechanical | unknown | unknown | unknown | unresolved | null | null | null | null | No statewide mechanical-code adoption verified. | none |
| Plumbing | unknown | unknown | unknown | unresolved | null | null | null | null | No statewide plumbing-code adoption verified. | none |
| Fuel Gas | unknown | unknown | unknown | unresolved | null | null | null | null | Fuel-gas construction-code adoption not parsed; LP-gas regulations require separate review. | none |
| Electrical | Kansas fire prevention code electrical standard / school wiring reference | NFPA 70 National Electrical Code | 2008 | adopted_fire_life_safety_reference | null | 2011-02-04 | null | null | School electric wiring must conform to the NEC adopted in K.A.R. 22-1-3; broader electrical-code authority remains unresolved. | src:usa-ks:kar-22-1-fire-prevention-code; src:usa-ks:ksreg-2024-kar-22-18-3 |
| Energy | unknown | unknown | unknown | unresolved | null | null | null | null | No statewide energy-code adoption verified. | none |
| Fire - construction references | Kansas fire prevention code | IBC and IFC construction-related provisions | 2006 IBC; 2006 IFC | adopted_statewide_fire_code_reference | null | 2011-02-04 | null | null | Existing facilities continuation rule applies unless a distinct hazard is determined. | src:usa-ks:kar-22-1-fire-prevention-code; src:usa-ks:ksfm-code-listing; src:usa-ks:ksfm-plan-review |
| Fire - operational / prevention code | Kansas fire prevention code | International Fire Code and NFPA standards | 2006 IFC; selected NFPA editions | adopted_statewide | null | 2011-02-04 | null | null | Municipal fire-code modifications require State Fire Marshal review; inconsistent local enactments are subject to State Fire Marshal determination. | src:usa-ks:ksreg-2026-hb2739; src:usa-ks:kar-22-1-fire-prevention-code; src:usa-ks:ksfm-code-listing |
| School buildings | Kansas school-building construction requirements | IBC; NEC; Life Safety Code for mobile/modular/portable/relocatable school buildings | IBC/NEC/LSC editions adopted in K.A.R. 22-1-3 | adopted_limited_scope | null | 2024-03-29 | null | null | Applies to school buildings; K.S.A. 31-150 includes plan-certification and postsecondary code-footprint provisions. | src:usa-ks:ksa-31-150; src:usa-ks:ksreg-2024-kar-22-18-3 |
| Accessibility | Kansas accessibility standards for publicly assisted dwellings | statute-specific state accessibility standards | no model edition | adopted_limited_scope | null | 2026-04-16 | null | null | Applies to new dwellings constructed with public financial assistance, with 2026 HB 2739 exceptions for specified housing programs. | src:usa-ks:ksreg-2026-hb2739; src:usa-ks:ksa-58-1402; src:usa-ks:ksa-58-1403; src:usa-ks:ksa-58-1404 |
| Elevator / Conveyance | Kansas Elevator Safety Act / KSFM elevator program | ASME A17.1 Safety Code for Elevators and Escalators | 2022 for post-2024-07-01 applicability; 2019 for qualifying prior projects | implemented_statewide_with_local_program_exceptions | 2023-12-21 | 2024-07-01 | 2024-07-01 | 2024-07-01 | ASME A17.1 (2019) remains allowed for elevators designed, approved, or contracted before 2024-07-01 unless variance applies. | src:usa-ks:ksfm-asme-a17-1-2022; src:usa-ks:ksfm-elevator-program; src:usa-ks:ksa-44-1814; src:usa-ks:ksa-44-1815 |

### 3.2 Adoption Records

```yaml
adoption_record:
  record_id: "adoption:usa-ks:kfpc-ibc-2006"
  code_family: "Building / fire-life-safety construction reference"
  state_code_name: "Kansas fire prevention code"
  base_model_code: "International Building Code"
  edition: "2006"
  adoption_date: null
  effective_date: "2011-02-04"
  operative_date: null
  mandatory_date: null
  adoption_instrument: "K.A.R. 22-1-3(a)"
  scope_note: "Adopted within the Kansas fire prevention code; not treated here as a verified comprehensive statewide building code for all local building-permit purposes."
  source_ids:
    - "src:usa-ks:kar-22-1-fire-prevention-code"
    - "src:usa-ks:ksfm-code-listing"
    - "src:usa-ks:ksfm-plan-review"
  confidence: 0.78

adoption_record:
  record_id: "adoption:usa-ks:kfpc-ifc-2006"
  code_family: "Fire - operational / prevention code"
  state_code_name: "Kansas fire prevention code"
  base_model_code: "International Fire Code"
  edition: "2006"
  adoption_date: null
  effective_date: "2011-02-04"
  operative_date: null
  mandatory_date: null
  adoption_instrument: "K.A.R. 22-1-3(b)"
  scope_note: "Adopted with stated chapter/reference exclusions."
  source_ids:
    - "src:usa-ks:kar-22-1-fire-prevention-code"
    - "src:usa-ks:ksfm-code-listing"
  confidence: 0.82

adoption_record:
  record_id: "adoption:usa-ks:kfpc-nec-2008"
  code_family: "Electrical - fire/life-safety reference"
  state_code_name: "Kansas fire prevention code"
  base_model_code: "NFPA 70 National Electrical Code"
  edition: "2008"
  adoption_date: null
  effective_date: "2011-02-04"
  operative_date: null
  mandatory_date: null
  adoption_instrument: "K.A.R. 22-1-3(n)"
  scope_note: "Adopted as one of the NFPA standards in the Kansas fire prevention code; K.A.R. 22-18-3 separately points school electric wiring to the NEC adopted by reference in K.A.R. 22-1-3."
  source_ids:
    - "src:usa-ks:kar-22-1-fire-prevention-code"
    - "src:usa-ks:ksreg-2024-kar-22-18-3"
  confidence: 0.74

adoption_record:
  record_id: "adoption:usa-ks:schools-kar-22-18-3"
  code_family: "School buildings"
  state_code_name: "Construction requirements for school buildings"
  base_model_code: "IBC, NEC, and Life Safety Code as adopted by K.A.R. 22-1-3"
  edition: "editions adopted in K.A.R. 22-1-3"
  adoption_date: null
  effective_date: "2024-03-29"
  operative_date: "2024-03-29"
  mandatory_date: null
  adoption_instrument: "K.A.R. 22-18-3, permanent administrative regulation"
  scope_note: "Applies to school-building construction; mobile, modular, portable, or relocatable school buildings must meet the Life Safety Code adopted by K.A.R. 22-1-3."
  source_ids:
    - "src:usa-ks:ksa-31-150"
    - "src:usa-ks:ksreg-2024-kar-22-18-3"
  confidence: 0.82

adoption_record:
  record_id: "adoption:usa-ks:elevators-asme-a17-1-2022"
  code_family: "Elevator / Conveyance"
  state_code_name: "Kansas Elevator Safety Act / KSFM elevator program"
  base_model_code: "ASME A17.1 Safety Code for Elevators and Escalators"
  edition: "2022"
  adoption_date: "2023-12-21"
  effective_date: "2024-07-01"
  operative_date: "2024-07-01"
  mandatory_date: "2024-07-01"
  adoption_instrument: "KSFM ASME A17.1 (2022) implementation notice"
  scope_note: "Elevators designed, approved, or contracted before 2024-07-01 may comply with ASME A17.1 (2019) unless the Fire Marshal grants an applicable variance."
  source_ids:
    - "src:usa-ks:ksfm-asme-a17-1-2022"
    - "src:usa-ks:ksfm-elevator-program"
  confidence: 0.86
```

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

Kansas distinguishes several different date concepts in the sources reviewed. The Kansas fire prevention code records reviewed do not provide a separate statewide permit-date transition rule for model building-code changes. They do, however, include an existing-facility continuation rule: facilities in service before the effective date of Fire Marshal rules may continue if not in strict conformity, so long as the State Fire Marshal does not determine them to be a distinct hazard to life or property. School-building regulation amendments became effective on 2024-03-29. The elevator program has a clear 2024-07-01 ASME A17.1 transition, and HB 2739 became effective upon Kansas Register publication on 2026-04-16 while leaving 2024 IFC adoption to rulemaking and ratification.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-ks:001 | Kansas fire prevention code facilities | existing_facility_continuation | facilities in service before effective date of rules | Facility is not in strict conformity with current rules but is not determined by State Fire Marshal to be a distinct hazard to life or property. | yes, limited by distinct-hazard determination | src:usa-ks:ksreg-2026-hb2739 | partially_verified |
| date-rule:usa-ks:002 | Municipal fire-code modifications | agency_review_period | 30 days after receipt of modification summary | Municipality adopts a nationally recognized fire code or fire-protection segment of a building code and modifies a section. | not applicable | src:usa-ks:kar-22-1-fire-prevention-code | partially_verified |
| date-rule:usa-ks:003 | Elevators | edition_transition | 2024-07-01 | Elevator designed, approved, or contracted on or after 2024-07-01. | yes, ASME A17.1 (2019) for elevators designed, approved, or contracted before 2024-07-01 unless variance applies | src:usa-ks:ksfm-asme-a17-1-2022 | partially_verified |
| date-rule:usa-ks:004 | Elevator inspections | inspection_interval | annually, every two years, or every three years by county population | Elevator owner/agent inspection responsibility under K.S.A. 44-1815 and KSFM program guidance. | not applicable | src:usa-ks:ksa-44-1815; src:usa-ks:ksfm-elevator-program | partially_verified |
| date-rule:usa-ks:005 | 2024 IFC update | rulemaking_review | Attorney General review within 120 days after proposed rules are submitted; legislative ratification still required | State Fire Marshal rulemaking to update Kansas fire prevention code with 2024 IFC and appropriate amendments. | current KFPC remains until rule adoption/ratification is confirmed | src:usa-ks:ksreg-2026-hb2739 | partially_verified_pending |
| date-rule:usa-ks:006 | School-building construction regulation | effective_date | 2024-03-29 | Permanent amendment to K.A.R. 22-18-3 became effective. | not specified in source | src:usa-ks:ksreg-2024-kar-22-18-3 | partially_verified |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Fire - operational / prevention code | International Fire Code | 2026-04-16 | null | null | null | null | active_watch | src:usa-ks:ksreg-2026-hb2739 | HB 2739 references rules to update the Kansas fire prevention code with the 2024 IFC, with appropriate amendments. The reviewed text does not itself complete adoption of 2024 IFC rules. |
| Building / fire-life-safety construction references | possible related IBC/KFPC amendments | null | null | null | null | null | monitor | src:usa-ks:ksreg-2026-hb2739 | If KSFM converts KFPC to a 2024 IFC-based structure, related construction-reference changes may be proposed. No completed adoption found in this pass. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-ks:001 | Fire/life safety | Buildings and places where people work, live, or congregate, except buildings used wholly as dwelling houses containing no more than two families | Fire Marshal rules adopted under K.S.A. 31-133 | Fire Marshal rule scope includes exits/fire escapes, fire control/detection/extinguishment equipment, and other safeguards, but excludes wholly residential dwelling houses containing no more than two families. | src:usa-ks:ksreg-2026-hb2739 | partially_verified |
| applicability-rule:usa-ks:002 | School buildings | Construction, reconstruction, or renovation of school buildings | School-building work subject to K.S.A. 31-150 and K.A.R. 22-18-3 | School buildings must meet IBC requirements, school wiring must conform to NEC, mobile/modular/portable/relocatable school buildings must meet Life Safety Code, and postsecondary projects must submit code footprints to KSFM. | src:usa-ks:ksa-31-150; src:usa-ks:ksreg-2024-kar-22-18-3 | partially_verified |
| applicability-rule:usa-ks:003 | Accessibility | New dwellings constructed with public financial assistance | Public financial assistance for covered dwellings | Covered dwellings must satisfy state accessibility standards; program administrators must require an affidavit of intent to comply before releasing funds. | src:usa-ks:ksreg-2026-hb2739; src:usa-ks:ksa-58-1402; src:usa-ks:ksa-58-1403; src:usa-ks:ksa-58-1404 | partially_verified_limited |
| applicability-rule:usa-ks:004 | Elevator / Conveyance | Elevators, escalators, and conveyances within scope of Elevator Safety Act | Elevator subject to state program unless exempt or covered by qualifying city/county program | Elevators must be registered and inspected, and installation/alteration generally requires a State Fire Marshal permit unless a city/county program meets or exceeds state standards. | src:usa-ks:ksfm-elevator-program; src:usa-ks:ksa-44-1804; src:usa-ks:ksa-44-1814; src:usa-ks:ksa-44-1815 | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** incorporation by reference with Kansas-specific exclusions, substitutions, and special statutory/regulatory provisions.

**Where amendments are published:** Kansas Administrative Regulations, Kansas Register notices, State Fire Marshal code-listing and program pages, and agency PDF guidance for some implementation notices.

**Amendment parsing status:** partial_core_fire_code_only

### 5.2 State Amendment Sources

| Amendment Source ID | Description | Source IDs | Parsing Status |
| --- | --- | --- | --- |
| amend-set:usa-ks:kar-22-1 | K.A.R. Article 22-1 Kansas fire prevention code; includes municipal modification review, compliance by certain building codes, and adoption of 2006 IBC, 2006 IFC, and NFPA standards with exclusions. | src:usa-ks:kar-22-1-fire-prevention-code | partially_parsed |
| amend-set:usa-ks:hb2739-2026 | 2026 statutory changes affecting Fire Marshal rules, municipal inconsistency, and 2024 IFC update rulemaking path. | src:usa-ks:ksreg-2026-hb2739 | partially_parsed |
| amend-set:usa-ks:kar-22-18-3-2024 | School-building construction requirements updated to refer to codes adopted in K.A.R. 22-1-3. | src:usa-ks:ksreg-2024-kar-22-18-3 | parsed_core |
| amend-set:usa-ks:elevator-asme-2022 | KSFM ASME A17.1 (2022) implementation notice with transition from 2019 edition for qualifying prior projects. | src:usa-ks:ksfm-asme-a17-1-2022 | parsed_core |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Base Code Provision / Area | Kansas Amendment / Rule | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| amend:usa-ks:kfpc-ibc-nfpa-exclusions | Building / fire construction references | IBC Chapter 35 references | K.A.R. 22-1-3(a) adopts the 2006 IBC including appendices but excludes Chapter 35 references to NFPA 13, 13D, 13R, 14, 30, 72, 101, and 110. | src:usa-ks:kar-22-1-fire-prevention-code | partially_verified |
| amend:usa-ks:kfpc-ifc-exclusions | Fire - operational / prevention code | IFC chapters and Chapter 45 references | K.A.R. 22-1-3(b) adopts the 2006 IFC including appendices but excludes Chapters 22, 30, 33, 34, 35, 36, and 38, plus specified NFPA references in Chapter 45. | src:usa-ks:kar-22-1-fire-prevention-code | partially_verified |
| amend:usa-ks:kfpc-local-modification-review | Fire - local amendments | Local modifications to nationally recognized fire codes or fire-protection building-code segments | Municipal summaries of modifications must be submitted to KSFM for approval or rejection within 30 days. | src:usa-ks:kar-22-1-fire-prevention-code | partially_verified |
| amend:usa-ks:kfpc-two-family-exclusion | Fire/life safety | One- and two-family dwelling-house scope | Fire Marshal rules under K.S.A. 31-133 do not apply to buildings used wholly as dwelling houses containing no more than two families. | src:usa-ks:ksreg-2026-hb2739 | partially_verified |
| amend:usa-ks:kfpc-existing-facilities | Fire/life safety | Existing facilities | Facilities in service before the effective date of Fire Marshal rules may continue in service if not a distinct hazard to life or property. | src:usa-ks:ksreg-2026-hb2739 | partially_verified |
| amend:usa-ks:elevator-pre-2024-projects | Elevator / Conveyance | ASME A17.1 2022 transition | Elevators designed, approved, or contracted before 2024-07-01 comply with ASME A17.1 (2019) unless KSFM grants an applicable variance. | src:usa-ks:ksfm-asme-a17-1-2022 | partially_verified |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-ks"
  model: "hybrid_state_fire_life_safety_plus_local_home_rule"
  enforcing_entities:
    - "Kansas Office of the State Fire Marshal for Kansas fire prevention code, selected plan-review/code-footprint functions, school fire/life-safety oversight, and elevator safety"
    - "Local fire chiefs or fire inspectors for annual school inspections in cities of the first and second class with a full-time fire chief or fire inspector"
    - "Local fire and building officials where local jurisdictions have adopted codes and where MOU collaboration or local review applies"
    - "City/county elevator programs that meet or exceed state requirements, where approved or applicable"
  required_officials:
    - "State Fire Marshal"
    - "Licensed elevator inspector, mechanic, contractor, or other authorized inspector for elevator inspections and work"
    - "Local fire chief or fire inspector for certain school inspections"
  state_reserved_activities:
    - "Kansas fire prevention code adoption and inconsistency determinations"
    - "Review of municipal fire-code modifications"
    - "Code-footprint review for covered projects and certain postsecondary school projects"
    - "Elevator registration, state permits, certificates, and inspection-program administration unless a qualifying local program applies"
  source_ids:
    - "src:usa-ks:ksreg-2026-hb2739"
    - "src:usa-ks:kar-22-1-fire-prevention-code"
    - "src:usa-ks:ksfm-plan-review"
    - "src:usa-ks:ksa-31-144"
    - "src:usa-ks:ksfm-elevator-program"
    - "src:usa-ks:ksa-44-1814"
    - "src:usa-ks:ksa-44-1815"
  verification_status: "partially_verified"
  confidence: 0.68
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-ks"
  model: "state_fire_code_consistency_review_plus_unresolved_general_building_amendment_scope"
  applies_to_code_families:
    - "Fire - operational / prevention code"
    - "Fire - construction references"
    - "Fire-protection segment of a nationally recognized building code"
    - "Elevator programs that meet or exceed state standards"
    - "General local building, residential, mechanical, plumbing, and energy amendments remain unresolved"
  approval_required: true
  approving_authority_id: "ahj:usa-ks:state-fire-marshal"
  filing_required: true
  registry_exists: null
  registry_source_ids: []
  legal_basis_source_ids:
    - "src:usa-ks:kar-22-1-fire-prevention-code"
    - "src:usa-ks:ksreg-2026-hb2739"
    - "src:usa-ks:ksa-44-1804"
  verification_status: "partially_verified"
  confidence: 0.70
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement and local amendment authority are separate in Kansas. Local jurisdictions may enforce locally adopted building and fire codes and may participate in KSFM plan-review or inspection coordination, but local fire-code modifications are subject to KSFM review and the Kansas fire prevention code has statewide uniform force. Local model-code adoptions for construction permitting still require jurisdiction-level lookup.

### 6.4 Known Local Amendment Registries

No statewide local-amendment registry was verified. K.A.R. 22-1-1 requires municipalities to submit summaries of modified fire-code provisions to the State Fire Marshal, but no public registry of those summaries was found in the official sources reviewed.

### 6.5 Municipality-Specific Known Amendments

No municipality-specific building-code or fire-code amendments were parsed. Wichita is identified by KSFM as conducting its own elevator inspections/program, but detailed Wichita elevator standards were not parsed.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: partial_state_layer_only

Expected jurisdiction stack:

Address
  -> State of Kansas
  -> County
  -> Municipality / unincorporated county
  -> Special districts, if applicable
  -> Local building AHJ, if local code adopted
  -> Local fire AHJ, if local fire code or inspection role applies
  -> Kansas Office of the State Fire Marshal for statewide fire/life-safety overlay
  -> Kansas Office of the State Fire Marshal elevator program, unless qualifying local program applies
  -> Applicable state code adoption records
  -> Applicable local adoption and amendment records

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | not selected | none | statewide | unknown | pending |
| County | not selected | none | statewide | unknown | pending |
| Municipality | not selected | none | statewide | unknown | pending |
| Fire District | not selected | none | statewide | unknown | pending |
| Special District | not selected | none | statewide | unknown | pending |

### 7.3 AHJ Contact Data

No AHJ contact data was populated. The only statewide contact path captured in this pass is the Kansas Office of the State Fire Marshal source registry entry.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Source Type | Title / Citation | URL | Date Accessed | Used For |
| --- | --- | --- | --- | --- | --- |
| src:usa-ks:ksreg-2026-hb2739 | session law / Kansas Register publication | Kansas Register, Volume 45, Issue 16, House Bill 2739, Doc. No. 054070 | https://sos.ks.gov/publications/register/Volume-45/Issues/Issue-16/04-16-26-54070.html | 2026-06-25 | Current amendments to K.S.A. 31-133, 31-134, and 58-1401; 2024 IFC rulemaking path; statewide force/inconsistency provisions; effective date from publication. |
| src:usa-ks:ksa-31-132 | statute | K.S.A. 31-132, Fire safety and prevention definitions | https://ksrevisor.gov/statutes/chapters/ch31/031_001_0032.html | 2026-06-25 | Definitions of municipality, nationally recognized fire prevention code, and nationally recognized building code. |
| src:usa-ks:ksa-31-134a | statute | K.S.A. 31-134a, compliance with certain building codes deemed compliance with fire prevention code | https://ksrevisor.gov/statutes/chapters/ch31/031_001_0034a.html | 2026-06-25 | Building-code compliance relationship and State Fire Marshal authority to specify subsequent editions or equivalent nationally recognized building codes. |
| src:usa-ks:kar-22-1-fire-prevention-code | administrative regulation / agency PDF | State Fire Marshal Permanent Administrative Regulations, Article 1, Kansas Fire Prevention Code, K.A.R. 22-1-1 through 22-1-4 | https://firemarshal.ks.gov/DocumentCenter/View/126/KAR22-1---Kansas-Fire-Prevention-Code-PDF | 2026-06-25 | K.A.R. 22-1-1 municipal modification review; 22-1-2 building-code compliance; 22-1-3 adopted IBC/IFC/NFPA standards and exclusions; variances. |
| src:usa-ks:ksfm-code-listing | agency page | Kansas Office of the State Fire Marshal, Code Listing | https://firemarshal.ks.gov/181/Code-Listing | 2026-06-25 | Agency-facing current code listing, including 1997 UBC, 2006 IBC, 2006 IFC, and adopted NFPA editions. |
| src:usa-ks:ksfm-plan-review | agency page | Kansas Office of the State Fire Marshal, Plans Review & Code Footprint Requirements | https://firemarshal.ks.gov/189/Plans-Review-Code-Footprint | 2026-06-25 | Current KSFM plan-review code-footprint summary; current codes for new construction/additions/changes of occupancy; MOU/local jurisdiction collaboration. |
| src:usa-ks:ksa-31-144 | statute | K.S.A. 31-144, school buildings; inspection; correction of violations | https://ksrevisor.gov/statutes/chapters/ch31/031_001_0044.html | 2026-06-25 | Annual school-building inspection allocation between local fire officials and State Fire Marshal. |
| src:usa-ks:ksa-31-150 | statute | K.S.A. 31-150, school buildings construction, reconstruction, renovation requirements | https://ksrevisor.gov/statutes/chapters/ch31/031_001_0050.html | 2026-06-25 | School construction code basis, accessibility, plan certification, postsecondary code-footprint submission, and State Fire Marshal later-edition authority. |
| src:usa-ks:ksreg-2024-kar-22-18-3 | administrative regulation / Kansas Register publication | Kansas Register, Volume 43, Issue 11, State Fire Marshal Permanent Administrative Regulations, K.A.R. 22-18-3 | https://www.kssos.org/publications/Register/Volume-43/Issues/Issue-11/03-14-24-51960.html | 2026-06-25 | Current school-building construction regulation tying school construction, NEC wiring, and mobile/modular/portable/relocatable school buildings to K.A.R. 22-1-3. |
| src:usa-ks:kan-const-art12-sec5 | state constitution | Kansas Constitution, Article 12, Section 5, cities' powers of home rule | https://ksrevisor.gov/kanconst/093_012_0005.html | 2026-06-25 | City home-rule baseline and limitations. |
| src:usa-ks:ksa-19-101a | statute | K.S.A. 19-101a, county home rule powers | https://ksrevisor.gov/statutes/chapters/ch19/019_001_0001a.html | 2026-06-25 | County home-rule baseline and limitations. |
| src:usa-ks:ksa-58-1402 | statute | K.S.A. 58-1402, accessibility standards for covered dwellings | https://ksrevisor.gov/statutes/chapters/ch58/058_014_0002.html | 2026-06-25 | State accessibility standards for dwellings receiving public financial assistance. |
| src:usa-ks:ksa-58-1403 | statute | K.S.A. 58-1403, application to new dwellings | https://ksrevisor.gov/statutes/chapters/ch58/058_014_0003.html | 2026-06-25 | Scope of accessibility standards for new dwellings. |
| src:usa-ks:ksa-58-1404 | statute | K.S.A. 58-1404, condition of release of public financial assistance | https://ksrevisor.gov/statutes/chapters/ch58/058_014_0004.html | 2026-06-25 | Affidavit and future assistance consequence for covered accessibility requirements. |
| src:usa-ks:ksfm-elevator-program | agency page | Kansas Office of the State Fire Marshal, Elevator Safety Program | https://firemarshal.ks.gov/342/Elevator-Safety-Program | 2026-06-25 | Elevator program authority, registration, inspections, licensing, certificate of operation, Wichita local program note, and program scope. |
| src:usa-ks:ksa-44-1801 | statute | K.S.A. 44-1801, Elevator Safety Act citation | https://ksrevisor.gov/statutes/chapters/ch44/044_018_0001.html | 2026-06-25 | Statutory citation for Elevator Safety Act. |
| src:usa-ks:ksa-44-1804 | statute | K.S.A. 44-1804, city and county standards not restricted if meeting/exceeding state act | https://ksrevisor.gov/statutes/chapters/ch44/044_018_0004.html | 2026-06-25 | Local elevator-program meet-or-exceed rule and notice to State Fire Marshal. |
| src:usa-ks:ksa-44-1814 | statute | K.S.A. 44-1814, elevator construction or alteration permits | https://ksrevisor.gov/statutes/chapters/ch44/044_018_0014.html | 2026-06-25 | State permit requirement and city/county exception for local standards meeting/exceeding state act. |
| src:usa-ks:ksa-44-1815 | statute | K.S.A. 44-1815, elevator inspections and testing | https://ksrevisor.gov/statutes/chapters/ch44/044_018_0015.html | 2026-06-25 | Elevator inspection intervals, report, correction period, extensions, load tests, and exceptions. |
| src:usa-ks:ksfm-asme-a17-1-2022 | agency notice / PDF | Kansas Office of the State Fire Marshal, ASME A17.1 (2022) Effective July 1, 2024 notice | https://firemarshal.ks.gov/DocumentCenter/View/1606/07012024-Implementation-of-ASME-A171-2022 | 2026-06-25 | ASME A17.1 (2022) implementation date and transition for pre-2024-07-01 projects. |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-ks:ksreg-2026-hb2739 | recently_enacted_law | The Kansas Register publication is the current source for 2026 HB 2739 amendments. Some Revisor statute pages viewed in this pass may lag the 2026 amendments. | prefer_for_current_31_133_31_134_58_1401_text |
| src:usa-ks:kar-22-1-fire-prevention-code | agency_pdf | The KSFM PDF is an official agency-hosted regulation excerpt, not a direct rules.ks.gov HTML page. It was cross-checked against KSFM current code-listing and plan-review pages. | acceptable_with_caveat |
| src:usa-ks:ksfm-code-listing | agency_summary | The page states it lists “some of the codes and editions that are currently being enforced,” so it is a useful official summary but not a substitute for adopted regulation text. | use_as_supplement_to_regulation |
| src:usa-ks:ksfm-plan-review | agency_summary | The page is an official practice guide for plan review/code footprints and should be reconciled with K.A.R. text for legal publication. | use_as_operational_source |
| src:usa-ks:ksfm-asme-a17-1-2022 | agency_pdf_notice | The source is an implementation notice rather than codified regulation text. It is official KSFM guidance and should be checked against final elevator regulations as they mature. | monitor_for_codification |
| src:usa-ks:ksa-58-1402 | statute_may_need_2026_reconciliation | K.S.A. 58-1402 was used for standards text; K.S.A. 58-1401 definitions were affected by 2026 HB 2739 and should be reconciled once the Revisor updates the chapter. | pair_with_hb2739 |

### 8.3 Supplemental Sources

No non-official source was relied on for report conclusions. Cornell/Justia-style regulation mirrors appeared in search results but were not used as controlling sources in the report body.

### 8.4 Source Extraction Metadata

| Extraction ID | Source IDs | Extraction Method | Extracted Fields | Extracted By | Date |
| --- | --- | --- | --- | --- | --- |
| extract:usa-ks:001 | src:usa-ks:ksreg-2026-hb2739 | official HTML review | K.S.A. 31-133/31-134 amendments, 2024 IFC rulemaking path, effective date | ChatGPT | 2026-06-25 |
| extract:usa-ks:002 | src:usa-ks:kar-22-1-fire-prevention-code | agency PDF text plus page screenshot review | K.A.R. 22-1-1, 22-1-2, 22-1-3 adopted codes/exclusions | ChatGPT | 2026-06-25 |
| extract:usa-ks:003 | src:usa-ks:ksfm-code-listing; src:usa-ks:ksfm-plan-review | official agency HTML review | current code listing; plan-review and MOU notes | ChatGPT | 2026-06-25 |
| extract:usa-ks:004 | src:usa-ks:ksa-31-144; src:usa-ks:ksa-31-150; src:usa-ks:ksreg-2024-kar-22-18-3 | official statute and Kansas Register review | school inspection and construction requirements | ChatGPT | 2026-06-25 |
| extract:usa-ks:005 | src:usa-ks:ksfm-elevator-program; src:usa-ks:ksa-44-1801; src:usa-ks:ksa-44-1804; src:usa-ks:ksa-44-1814; src:usa-ks:ksa-44-1815; src:usa-ks:ksfm-asme-a17-1-2022 | official statute, agency page, and agency PDF review | elevator authority, permits, inspections, ASME A17.1 transition | ChatGPT | 2026-06-25 |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified_for_current_file | 1.00 | none | Status reflects source-backed core fields and unresolved non-core fields. |
| report | risk.overall_confidence | 0.62 | verified_for_current_file | 1.00 | none | Confidence is limited by unresolved local adoptions and non-fire code families. |
| ahj:usa-ks:state-fire-marshal | authority.name | Kansas Office of the State Fire Marshal | verified | 0.90 | src:usa-ks:ksreg-2026-hb2739; src:usa-ks:ksfm-plan-review | Primary statewide fire/life-safety authority. |
| adoption:usa-ks:kfpc-ibc-2006 | edition | 2006 IBC | verified | 0.82 | src:usa-ks:kar-22-1-fire-prevention-code; src:usa-ks:ksfm-code-listing | Verified as part of Kansas fire prevention code, not a standalone statewide building code. |
| adoption:usa-ks:kfpc-ifc-2006 | edition | 2006 IFC | verified | 0.86 | src:usa-ks:kar-22-1-fire-prevention-code; src:usa-ks:ksfm-code-listing | K.A.R. 22-1-3 includes exclusions. |
| adoption:usa-ks:elevators-asme-a17-1-2022 | effective_date | 2024-07-01 | verified | 0.90 | src:usa-ks:ksfm-asme-a17-1-2022 | Agency notice gives clear transition date. |
| local-amendment-rule:usa-ks | approval_required | true for municipal fire-code modifications | verified_limited_scope | 0.86 | src:usa-ks:kar-22-1-fire-prevention-code | General local building-code amendment scope still unresolved. |
| date-rule:usa-ks:005 | 2024_IFC_update_status | pending rulemaking/ratification | verified | 0.78 | src:usa-ks:ksreg-2026-hb2739 | HB 2739 does not by itself complete rule adoption. |
| school-buildings | effective_date | 2024-03-29 | verified | 0.82 | src:usa-ks:ksreg-2024-kar-22-18-3 | Effective date from Kansas Register permanent regulation entry. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Source IDs used in body are present in section 8. |
| All authority IDs resolve | pass | Non-unknown authority IDs are introduced in section 2; unresolved IDs are explicit. |
| All current code families have adoption matrix rows | pass | Rows are present; unresolved families are retained explicitly. |
| Building and operational fire code are separated | pass | IBC construction reference and IFC operational/prevention code are separate rows. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Unknown dates are null; effective and mandatory dates are separated. |
| Effective dates are valid ISO dates | pass | Dates use YYYY-MM-DD. |
| No impossible date sequences | pass | No adoption/effective/operative/mandatory sequence conflicts detected in populated rows. |
| Transition rules have explicit trigger conditions | pass | Captured transition rows include trigger conditions. |
| Permit-date logic is captured where applicable | fail | No statewide building-permit date rule was found for general model-code transitions. Elevator installation/alteration permit rules were captured. |
| Local enforcement model classified | pass | Hybrid state fire/life-safety plus local home-rule model recorded. |
| Local amendment rule classified | pass | Fire-code modification review is classified; general local building amendments remain unresolved. |
| AHJ confirmation metadata present | fail | No jurisdiction-specific AHJ contacts or local adoption records were entered. |
| Official-source caveats captured | pass | Caveats for recently enacted law, agency PDFs, and agency summaries are listed. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-ks:001 | high | general building-code authority | Determine whether any statewide authority beyond the Kansas fire prevention code administers general building, residential, mechanical, plumbing, or energy code adoption. | Review Kansas Department of Administration, local-government enabling statutes, and any trade licensing boards. | null | null | open |
| issue:usa-ks:002 | high | 2024 IFC update | Confirm whether KSFM has proposed, adopted, and obtained ratification for rules updating the Kansas fire prevention code to the 2024 IFC. | Monitor Kansas Register proposed/permanent regulations, Attorney General review, and legislative ratification records. | null | null | open |
| issue:usa-ks:003 | high | local model-code adoptions | Build city/county adoption layer for IBC, IRC, IPC, IMC, IFGC, IECC, NEC, and IFC editions. | Prioritize Wichita, Overland Park, Kansas City, Olathe, Topeka, Lawrence, Johnson County, Sedgwick County, and unincorporated counties. | null | null | open |
| issue:usa-ks:004 | medium | local fire amendment registry | Determine whether KSFM maintains a public or internal registry of municipal fire-code modifications submitted under K.A.R. 22-1-1. | Search KSFM records or request registry/index if public. | null | null | open |
| issue:usa-ks:005 | medium | accessibility enforcement | Reconcile 2026 HB 2739 amendments to K.S.A. 58-1401 with updated Revisor text and identify administering/enforcing officials for accessibility affidavits. | Re-check Revisor chapter 58 after codification and review housing program forms. | null | null | open |
| issue:usa-ks:006 | medium | elevator regulations | Verify whether final K.A.R. elevator rules now codify ASME A17.1 (2022) and implementation guidance. | Review Kansas Administrative Regulations, KSFM elevator FAQ, and Kansas Register future/permanent regulations. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-ks:kansas-register-fire-code | src:usa-ks:ksreg-2026-hb2739 | register_search | monthly | proposed or permanent KSFM regulations updating Kansas fire prevention code or adopting 2024 IFC | 2026-06-25 | active |
| watch:usa-ks:ksfm-code-listing | src:usa-ks:ksfm-code-listing | html_diff | monthly | code-listing edition changes or new code families | 2026-06-25 | active |
| watch:usa-ks:ksfm-plan-review | src:usa-ks:ksfm-plan-review | html_diff | monthly | plan-review code list, code-footprint rules, or MOU guidance changes | 2026-06-25 | active |
| watch:usa-ks:kar-22-1 | src:usa-ks:kar-22-1-fire-prevention-code | regulation_diff | monthly | K.A.R. 22-1-1 through 22-1-3 amendments | 2026-06-25 | active |
| watch:usa-ks:elevator-program | src:usa-ks:ksfm-elevator-program | html_diff | quarterly | ASME edition, inspection interval, certificate, or local program changes | 2026-06-25 | active |
| watch:usa-ks:revisor-post-hb2739 | src:usa-ks:ksreg-2026-hb2739 | statute_codification_check | monthly | Revisor updates K.S.A. 31-133, 31-134, and 58-1401 after HB 2739 | 2026-06-25 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-25 | Replaced Kansas baseline draft with partially verified source-backed report. | report:usa-ks; ahj:usa-ks:state-fire-marshal; adoption:usa-ks:kfpc-ibc-2006; adoption:usa-ks:kfpc-ifc-2006; adoption:usa-ks:elevators-asme-a17-1-2022 | src:usa-ks:ksreg-2026-hb2739; src:usa-ks:kar-22-1-fire-prevention-code; src:usa-ks:ksfm-code-listing; src:usa-ks:ksfm-plan-review; src:usa-ks:ksfm-elevator-program | ChatGPT | Core authority, fire-code adoption, school-building, elevator, local amendment, and pending 2024 IFC fields populated. |
