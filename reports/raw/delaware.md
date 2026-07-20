---
state:
  state_id: "US-DE"
  name: "Delaware"
  abbreviation: "DE"
report:
  report_id: "state-report:usa-de"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-25"
  last_verified: "2026-06-25"
  reviewed_by: null
risk:
  overall_confidence: 0.80 # 0.00 - 1.00
  risk_flags:
    - "no_single_statewide_building_code_verified"
    - "local_building_code_editions_jurisdiction_specific"
    - "hybrid_state_local_enforcement"
    - "energy_implementation_cutover_requires_local_followup"
    - "accessibility_authority_unresolved"
  open_questions_count: 5

---

# State Building Code Authority Report: Delaware

## 1. Executive Summary

- **Authority model:** Delaware is a split-authority state. County and municipal governments have statutory authority to adopt and enforce local building, plumbing, electrical, and similar codes, but plumbing, fire prevention, and energy conservation each have distinct statewide rules or statewide minimums. The result is not a single statewide IBC/IRC adoption.

- **Statewide code status:** Statewide or state-minimum rules were verified for plumbing, fire prevention, and energy conservation. The Delaware Energy Office adopted the 2024 IECC and ASHRAE 90.1-2022 package effective 2026-04-11. The State Fire Prevention Commission/State Fire Marshal announced the 2026 Delaware State Fire Prevention Regulations effective 2026-01-01, including 2024 NFPA 1, 2024 NFPA 101, and 2023 NFPA 70 references. The Delaware Plumbing Program states that Delaware uses the 2021 IPC with Delaware amendments.

- **Local enforcement model:** Local jurisdictions administer local building codes and enforce state energy requirements locally. Plumbing enforcement is statewide by statute with local enforcement responsibilities and local modifications requiring state review. Fire prevention is statewide through the State Fire Prevention Commission and State Fire Marshal.

- **Local amendment posture:** Mixed. Local building-code adoption is jurisdiction-specific. Local plumbing additions or modifications must be submitted to the Division of Public Health for review and approval. Local energy stretch codes are allowed, but a local jurisdiction may not fully replace the state energy code without the process directed by the Delaware Energy Office. More stringent local fire regulations may govern if not inconsistent with the state fire code and good engineering practice.

- **Known transition periods or pending changes:** Fire projects submitted on or after 2025-12-31 must comply with the 2026 Delaware State Fire Prevention Regulations. The energy package became effective 2026-04-11 and local jurisdictions have 12 months from promulgation to implement and enforce the revised code. The statutory zero-net-energy-capable milestones remain flagged because the 2026 DNREC order stated that implementation was not appropriate at that time and deferred implementation to future legislative or executive direction.

- **Production readiness:** usable_with_caution

### Key Findings

```yaml
---
key_findings:
- topic: Statewide building-code authority
  finding: No single statewide IBC/IRC adopting authority was verified; county and
    municipal governments have local code authority.
  confidence: 0.94
  source_ids:
  - src:usa-de:code:title16:ch76
- topic: Primary local building authority
  finding: County Councils, Levy Court, and municipalities may adopt and enforce building,
    plumbing, electrical, and similar codes.
  confidence: 0.96
  source_ids:
  - src:usa-de:code:title16:ch76
- topic: State energy code
  finding: The Delaware Energy Office adopted a 2026 package based on the 2024 IECC
    and ASHRAE 90.1-2022, effective 2026-04-11, with a 12-month local implementation
    window.
  confidence: 0.97
  source_ids:
  - src:usa-de:dnrec:building-energy-codes
  - src:usa-de:dnrec:energy-order-2026-cce-0002
- topic: State plumbing code
  finding: Delaware's plumbing program uses the 2021 IPC as the Delaware Plumbing
    Code with Delaware-specific amendments, and DPH has statutory authority to adopt
    and enforce the IPC.
  confidence: 0.94
  source_ids:
  - src:usa-de:code:title16:ch79
  - src:usa-de:dhss:plumbing-program
- topic: Fire prevention authority
  finding: The State Fire Prevention Commission promulgates regulations with statewide
    force, and the State Fire Marshal enforces fire laws, fire-prevention regulations,
    plan review, and listed fire-safety systems.
  confidence: 0.98
  source_ids:
  - src:usa-de:code:title16:ch66-sc01
  - src:usa-de:code:title16:ch66-sc02
- topic: Current fire regulation package
  finding: The 2026 Delaware State Fire Prevention Regulations became effective 2026-01-01
    and require projects submitted on or after 2025-12-31 to comply.
  confidence: 0.93
  source_ids:
  - src:usa-de:firemarshal:2026-dsfpr-update
  - src:usa-de:regulations:title1:701-2025-final
- topic: Electrical code reference
  finding: Electrical work is tied to the NEC as adopted through the State Fire Prevention
    Commission; the 2026 fire-regulation update lists NFPA 70, National Electrical
    Code, 2023 edition.
  confidence: 0.86
  source_ids:
  - src:usa-de:firemarshal:2026-dsfpr-update
  - src:usa-de:regulations:title24:1400
- topic: Local amendments
  finding: Local building adoption is broad; plumbing local modifications require
    DPH approval; energy stretch codes are allowed with limits; more stringent local
    fire regulations may apply if not inconsistent.
  confidence: 0.9
  source_ids:
  - src:usa-de:code:title16:ch76
  - src:usa-de:code:title16:ch79
  - src:usa-de:code:title16:ch66-sc01
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-de:split-building-authority |
| Authority name | Split state/local code authority; county and municipal building-code authorities for general building codes |
| Authority type | hybrid_state_local |
| Legal basis | 16 Del. C. § 7601; 16 Del. C. ch. 79; 16 Del. C. ch. 66; 7 DE Admin. Code 2101 |
| Role | Delaware does not present one verified statewide general building-code adoption in the sources reviewed. General building-code adoption is local; statewide or state-minimum rules apply for energy, plumbing, and fire prevention. |
| Enforcement model | hybrid |
| Source IDs | src:usa-de:code:title16:ch76; src:usa-de:code:title16:ch79; src:usa-de:code:title16:ch66-sc01; src:usa-de:dnrec:building-energy-codes |
| Verification status | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-de:local-building-authorities | County Councils, Levy Court, and municipal governments | Adopt and enforce local building codes | 16 Del. C. § 7601 | src:usa-de:code:title16:ch76 | verified |
| Residential | ahj:usa-de:local-building-authorities | County Councils, Levy Court, and municipal governments | Adopt and enforce local residential building codes where enacted locally | 16 Del. C. § 7601 | src:usa-de:code:title16:ch76 | partially_verified |
| Existing Building / Rehabilitation | ahj:usa-de:local-building-authorities | County Councils, Levy Court, and municipal governments | Existing-building code adoption is jurisdiction-specific; statewide edition not identified | 16 Del. C. § 7601; local ordinances | src:usa-de:code:title16:ch76; src:usa-de:new-castle:county-laws-code | partially_verified |
| Mechanical | ahj:usa-de:local-building-authorities | County Councils, Levy Court, and municipal governments | Mechanical-code adoption is jurisdiction-specific | 16 Del. C. § 7601; local ordinances | src:usa-de:code:title16:ch76; src:usa-de:new-castle:county-laws-code | partially_verified |
| Plumbing | ahj:usa-de:dph-plumbing-program | Delaware Division of Public Health, Plumbing Program | Adopts and enforces the Delaware Plumbing Code; local jurisdictions enforce the DPH-adopted code and submit additions or modifications for approval | 16 Del. C. §§ 7903, 7904, 7906 | src:usa-de:code:title16:ch79; src:usa-de:dhss:plumbing-program | verified |
| Fuel Gas | ahj:usa-de:local-building-authorities | County Councils, Levy Court, and municipal governments, with fire-code references where applicable | Local fuel-gas/mechanical code adoption; fire regulations also reference NFPA 54, National Fuel Gas Code | 16 Del. C. § 7601; 1 DE Admin. Code 701 | src:usa-de:code:title16:ch76; src:usa-de:firemarshal:2026-dsfpr-update; src:usa-de:new-castle:county-laws-code | partially_verified |
| Electrical | ahj:usa-de:state-fire-prevention-commission | State Fire Prevention Commission; Board of Electrical Examiners; local code officials | NEC reference is adopted through the State Fire Prevention Commission; electrical work also remains subject to applicable local building code | 1 DE Admin. Code 701; 24 DE Admin. Code 1400 | src:usa-de:firemarshal:2026-dsfpr-update; src:usa-de:regulations:title24:1400 | partially_verified |
| Energy | ahj:usa-de:delaware-energy-office | Delaware Energy Office / DNREC Division of Climate, Coastal and Energy | Adopts statewide energy-code minimums and reviews updates triennially | 16 Del. C. § 7602; 7 DE Admin. Code 2101 | src:usa-de:code:title16:ch76; src:usa-de:dnrec:building-energy-codes; src:usa-de:dnrec:energy-order-2026-cce-0002 | verified |
| Fire - construction references | ahj:usa-de:state-fire-marshal | State Fire Marshal | Reviews plans and specifications for covered buildings, occupancy changes, fire-protection systems, alarms, and similar fire-safety work | 16 Del. C. § 6607 | src:usa-de:code:title16:ch66-sc02 | verified |
| Fire - operational / prevention code | ahj:usa-de:state-fire-prevention-commission | State Fire Prevention Commission / State Fire Marshal | Promulgates and enforces statewide fire-prevention regulations | 16 Del. C. ch. 66; 1 DE Admin. Code 701 | src:usa-de:code:title16:ch66-sc01; src:usa-de:code:title16:ch66-sc02; src:usa-de:firemarshal:2026-dsfpr-update | verified |
| Accessibility | ahj:usa-de:unresolved-accessibility-authority | unresolved | Statewide accessibility-specific construction-code authority was not fully traced in this pass | unresolved | none | unresolved |
| Elevator / Conveyance | ahj:usa-de:local-elevator-communication-rule-authorities | County and municipal governments | May adopt rules to effectuate passenger-elevator emergency communication requirements | 16 Del. C. § 8701(c) | src:usa-de:code:title16:ch87 | verified |

### 2.3 Authority Hierarchy Notes

Delaware should be modeled by code family, not as a single statewide building-code adoption. County and municipal governments may adopt and enforce local general building, plumbing, electrical, and similar codes under Chapter 76, but Chapter 79 creates a statewide plumbing framework under the Division of Public Health. Chapter 66 creates statewide fire-prevention authority through the State Fire Prevention Commission and State Fire Marshal. Section 7602 and 7 DE Admin. Code 2101 establish the energy-code floor and update process through the Delaware Energy Office.

Local examples should not be promoted to statewide facts. New Castle County, for example, lists 2024 IBC/IRC/IEBC effective 2026-01-01, 2021 IMC/IFGC effective 2024-01-01, 2021 IPC effective 2024-01-01, and 2018 IECC effective 2021-01-01, but those are local county records rather than a statewide code matrix.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-de:001 | ahj:usa-de:local-building-authorities | adopts_and_enforces | local building, residential, mechanical, electrical, plumbing, and similar codes | src:usa-de:code:title16:ch76 | verified |
| edge:usa-de:002 | ahj:usa-de:dph-plumbing-program | sets_floor_for | Delaware Plumbing Code and DPH-approved local plumbing modifications | src:usa-de:code:title16:ch79; src:usa-de:dhss:plumbing-program | verified |
| edge:usa-de:003 | ahj:usa-de:state-fire-prevention-commission | promulgates | statewide fire-prevention regulations with force of law | src:usa-de:code:title16:ch66-sc01 | verified |
| edge:usa-de:004 | ahj:usa-de:state-fire-marshal | enforces | fire laws, Commission regulations, plan review, exits, fire protection systems, and related matters | src:usa-de:code:title16:ch66-sc02 | verified |
| edge:usa-de:005 | ahj:usa-de:delaware-energy-office | adopts_floor_for | statewide energy conservation code | src:usa-de:code:title16:ch76; src:usa-de:dnrec:building-energy-codes | verified |
| edge:usa-de:006 | ahj:usa-de:delaware-energy-office | reviews_limits_on | local energy stretch-code or replacement processes | src:usa-de:code:title16:ch76 | verified |
| edge:usa-de:007 | ahj:usa-de:state-fire-prevention-commission | permits_more_stringent_local_rules_if | local fire rules are not inconsistent with state code and good engineering practice | src:usa-de:code:title16:ch66-sc01 | verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | Local county/municipal building codes | jurisdiction-specific | jurisdiction-specific | local_current | null | null | null | null | Local adoption controls. No statewide IBC edition was verified. | src:usa-de:code:title16:ch76 |
| Residential | Local county/municipal residential building codes | jurisdiction-specific | jurisdiction-specific | local_current | null | null | null | null | Local adoption controls. No statewide IRC edition was verified. | src:usa-de:code:title16:ch76 |
| Existing Building / Rehabilitation | Local county/municipal existing-building codes | jurisdiction-specific | jurisdiction-specific | local_current | null | null | null | null | Local adoption controls; New Castle County is only a local example. | src:usa-de:code:title16:ch76; src:usa-de:new-castle:county-laws-code |
| Mechanical | Local county/municipal mechanical codes | jurisdiction-specific | jurisdiction-specific | local_current | null | null | null | null | Local adoption controls; New Castle County is only a local example. | src:usa-de:code:title16:ch76; src:usa-de:new-castle:county-laws-code |
| Plumbing | Delaware Plumbing Code | International Plumbing Code | 2021 IPC with Delaware amendments | current | null | null | null | null | DPH adopts and enforces the IPC; local jurisdictions enforce the DPH-adopted code and submit additions or modifications for DPH review and approval. | src:usa-de:code:title16:ch79; src:usa-de:dhss:plumbing-program |
| Fuel Gas | Local fuel-gas/mechanical codes plus fire-code references | jurisdiction-specific; NFPA 54 where incorporated by fire regulations | jurisdiction-specific; NFPA 54 2024 listed in 2026 DSFPR update | partially_verified | 2025-12-16 | 2026-01-01 | 2025-12-31 | null | Local adoption controls for fuel-gas/mechanical codes; fire-regulation projects submitted on or after 2025-12-31 follow 2026 DSFPR references. | src:usa-de:code:title16:ch76; src:usa-de:firemarshal:2026-dsfpr-update |
| Electrical | NEC as adopted through the State Fire Prevention Commission; local electrical codes where applicable | NFPA 70, National Electrical Code | 2023 | partially_verified | 2025-12-16 | 2026-01-01 | 2025-12-31 | null | The 2026 DSFPR update lists NEC 2023; projects submitted on or after 2025-12-31 follow the 2026 DSFPR. Local code requirements still apply. | src:usa-de:firemarshal:2026-dsfpr-update; src:usa-de:regulations:title24:1400 |
| Energy - residential | Delaware State Energy Conservation Code | 2024 IECC | 2024 | current | 2026-03-10 | 2026-04-11 | null | null | Revised 7 DE Admin. Code 2101 effective 2026-04-11; local jurisdictions have 12 months from promulgation to implement and enforce. | src:usa-de:dnrec:building-energy-codes; src:usa-de:dnrec:energy-order-2026-cce-0002 |
| Energy - commercial and high-rise residential | Delaware State Energy Conservation Code | 2024 IECC; ASHRAE 90.1 | 2024 IECC; ASHRAE 90.1-2022 | current | 2026-03-10 | 2026-04-11 | null | null | Commercial/high-rise residential compliance options include 2024 IECC prescriptive, 2024 IECC simulated performance, or ASHRAE 90.1-2022. | src:usa-de:dnrec:building-energy-codes; src:usa-de:dnrec:energy-order-2026-cce-0002 |
| Energy - EV and solar-ready appendices | Delaware State Energy Conservation Code amendments | 2024 IECC Appendix RE; 2024 IECC Appendix CB | 2024 | current | 2026-03-10 | 2026-04-11 | null | null | 2026 order adopts EV-capable Appendix RE and solar-ready Appendix CB with Delaware amendments; Chapter 80 and § 7605 remain separate statutory references. | src:usa-de:dnrec:energy-order-2026-cce-0002; src:usa-de:code:title16:ch80; src:usa-de:code:title16:ch76 |
| Fire - construction references | Delaware State Fire Prevention Regulations | NFPA 1, NFPA 101, and listed NFPA standards | 2026 DSFPR; NFPA 1 2024; NFPA 101 2024 | current | 2025-12-16 | 2026-01-01 | 2025-12-31 | null | Projects submitted on or after 2025-12-31 must comply with the 2026 DSFPR. | src:usa-de:firemarshal:2026-dsfpr-update; src:usa-de:code:title16:ch66-sc02 |
| Fire - operational / prevention code | Delaware State Fire Prevention Regulations | NFPA 1 and listed NFPA standards | 2026 DSFPR; NFPA 1 2024 | current | 2025-12-16 | 2026-01-01 | 2025-12-31 | null | Commission regulations have statewide force; more stringent local rules may apply if consistent with state code and good engineering practice. | src:usa-de:firemarshal:2026-dsfpr-update; src:usa-de:code:title16:ch66-sc01 |
| Accessibility | unresolved statewide accessibility construction-code record | unresolved | null | unresolved | null | null | null | null | Statewide accessibility construction-code authority was not resolved. | none |
| Elevator / Conveyance | Passenger-elevator emergency communication requirement | state statute | statutory requirement | current | null | 1997-01-26 | 1997-01-26 | 1997-01-26 | Passenger elevators must have direct emergency communication; private residential elevator exception applies; counties and municipalities may adopt implementing rules. | src:usa-de:code:title16:ch87 |

### 3.2 Adoption Records

```yaml
adoption_id: "adoption:usa-de:building:local-codes"
state_id: "US-DE"
code_family: "building"
status: "local_current"
state_code:
  name: "County/municipal building codes"
  edition_label: "jurisdiction-specific"
  codification: "16 Del. C. § 7601"
base_model_code:
  publisher: "jurisdiction-specific"
  code_name: null
  edition_year: null
  incorporated_by_reference: null
authority:
  adopting_authority_id: "ahj:usa-de:local-building-authorities"
  enforcing_authority_model: "local"
  interpretation_authority_id: "ahj:usa-de:local-building-authorities"
dates:
  adoption_date: null
  effective_date: null
  operative_date: null
  mandatory_date: null
  replacement_date: null
applicability:
  date_trigger: "local_permit_or_local_ordinance_rule"
  applies_to:
    - "new_construction"
    - "alteration"
    - "repair"
    - "addition"
    - "commercial"
    - "residential"
  exclusions: []
  special_conditions:
    - "No statewide IBC or IRC edition was verified from the statewide sources reviewed."
    - "Local code content must be resolved by county, municipality, and project location."
transition:
  exists: true
  rule_id: "date-rule:usa-de:building:local-adoption"
  start_date: null
  end_date: null
  prior_code_allowed: null
  prior_code_condition: "Local ordinance and local permit administration control."
amendments:
  state_amended: false
  amendment_set_ids:
    - "amendment-set:usa-de:building:local-codes"
  amendment_source_ids:
    - "src:usa-de:code:title16:ch76"
provenance:
  source_ids:
    - "src:usa-de:code:title16:ch76"
  field_sources:
    state_code.name: ["src:usa-de:code:title16:ch76"]
    state_code.codification: ["src:usa-de:code:title16:ch76"]
    authority.adopting_authority_id: ["src:usa-de:code:title16:ch76"]
    applicability.date_trigger: ["src:usa-de:code:title16:ch76"]
verification:
  status: "partially_verified"
  confidence: 0.90
  notes: "Local code authority verified; statewide IBC/IRC adoption not found; local editions remain open."

adoption_id: "adoption:usa-de:plumbing:delaware-plumbing-code"
state_id: "US-DE"
code_family: "plumbing"
status: "current"
state_code:
  name: "Delaware Plumbing Code"
  edition_label: "2021 IPC with Delaware-specific amendments"
  codification: "16 Del. C. ch. 79; Delaware Plumbing Code"
base_model_code:
  publisher: "International Code Council"
  code_name: "International Plumbing Code"
  edition_year: 2021
  incorporated_by_reference: true
authority:
  adopting_authority_id: "ahj:usa-de:dph-plumbing-program"
  enforcing_authority_model: "hybrid_state_local"
  interpretation_authority_id: "ahj:usa-de:dph-plumbing-program"
dates:
  adoption_date: null
  effective_date: null
  operative_date: null
  mandatory_date: null
  replacement_date: null
applicability:
  date_trigger: "permit_or_state_plumbing_program_rule"
  applies_to:
    - "construction"
    - "alteration"
    - "repair"
    - "plumbing_systems"
  exclusions:
    - "owner_work_on_owner_occupied_single_family_home_subject_to_permit"
  special_conditions:
    - "DPH adopts and enforces the most recent IPC within one calendar year of issuance."
    - "Every political subdivision must enforce the IPC as adopted and modified by DPH."
    - "Local additions or modifications must be submitted to DPH for review and approval."
transition:
  exists: true
  rule_id: "date-rule:usa-de:plumbing:dph-ipc-adoption"
  start_date: null
  end_date: null
  prior_code_allowed: null
  prior_code_condition: "DPH and local jurisdiction implementation details require project-specific confirmation."
amendments:
  state_amended: true
  amendment_set_ids:
    - "amendment-set:usa-de:plumbing:delaware-amendments"
  amendment_source_ids:
    - "src:usa-de:dhss:plumbing-program"
    - "src:usa-de:code:title16:ch79"
provenance:
  source_ids:
    - "src:usa-de:code:title16:ch79"
    - "src:usa-de:dhss:plumbing-program"
  field_sources:
    state_code.name: ["src:usa-de:dhss:plumbing-program"]
    state_code.edition_label: ["src:usa-de:dhss:plumbing-program"]
    authority.adopting_authority_id: ["src:usa-de:code:title16:ch79"]
    amendments.state_amended: ["src:usa-de:dhss:plumbing-program"]
verification:
  status: "partially_verified"
  confidence: 0.94
  notes: "Current edition verified from official program page; adoption/effective date for the 2021 IPC amendment set was not extracted."

adoption_id: "adoption:usa-de:energy:2024-iecc-ashrae-2022"
state_id: "US-DE"
code_family: "energy"
status: "current"
state_code:
  name: "Delaware State Energy Conservation Code"
  edition_label: "2024 IECC and ASHRAE 90.1-2022 package"
  codification: "7 DE Admin. Code 2101; 16 Del. C. § 7602"
base_model_code:
  publisher: "International Code Council; ASHRAE"
  code_name: "2024 IECC; ASHRAE 90.1-2022"
  edition_year: 2024
  incorporated_by_reference: true
authority:
  adopting_authority_id: "ahj:usa-de:delaware-energy-office"
  enforcing_authority_model: "local_with_state_floor"
  interpretation_authority_id: "ahj:usa-de:delaware-energy-office"
dates:
  adoption_date: "2026-03-10"
  effective_date: "2026-04-11"
  operative_date: null
  mandatory_date: null
  replacement_date: null
applicability:
  date_trigger: "local_implementation_or_permit_rule"
  applies_to:
    - "new_residential"
    - "new_commercial"
    - "high_rise_residential"
    - "energy_alterations_where_code_applies"
  exclusions:
    - "agricultural_structures_under_section_7602"
  special_conditions:
    - "Residential buildings use 2024 IECC."
    - "Commercial and high-rise residential compliance options include 2024 IECC prescriptive, 2024 IECC simulated performance, and ASHRAE 90.1-2022."
    - "Local jurisdictions have 12 months from promulgation to implement and enforce."
    - "Local stretch codes are allowed; full local replacement of the state code requires a formal process under Delaware Energy Office guidance."
transition:
  exists: true
  rule_id: "date-rule:usa-de:energy:2026-update"
  start_date: "2026-04-11"
  end_date: null
  prior_code_allowed: null
  prior_code_condition: "Local implementation date and local permit cutover must be confirmed by jurisdiction."
amendments:
  state_amended: true
  amendment_set_ids:
    - "amendment-set:usa-de:energy:2026-2101"
  amendment_source_ids:
    - "src:usa-de:dnrec:energy-order-2026-cce-0002"
provenance:
  source_ids:
    - "src:usa-de:code:title16:ch76"
    - "src:usa-de:dnrec:building-energy-codes"
    - "src:usa-de:dnrec:energy-order-2026-cce-0002"
  field_sources:
    state_code.name: ["src:usa-de:dnrec:building-energy-codes"]
    state_code.edition_label: ["src:usa-de:dnrec:building-energy-codes", "src:usa-de:dnrec:energy-order-2026-cce-0002"]
    dates.adoption_date: ["src:usa-de:dnrec:energy-order-2026-cce-0002"]
    dates.effective_date: ["src:usa-de:dnrec:building-energy-codes", "src:usa-de:dnrec:energy-order-2026-cce-0002"]
    transition.rule_id: ["src:usa-de:code:title16:ch76", "src:usa-de:dnrec:building-energy-codes"]
verification:
  status: "partially_verified"
  confidence: 0.97
  notes: "State adoption and effective date verified; local implementation cutover remains jurisdiction-specific."

adoption_id: "adoption:usa-de:fire:2026-dsfpr"
state_id: "US-DE"
code_family: "fire_operational"
status: "current"
state_code:
  name: "Delaware State Fire Prevention Regulations"
  edition_label: "2026 edition"
  codification: "1 DE Admin. Code 701; 16 Del. C. ch. 66"
base_model_code:
  publisher: "NFPA and Delaware State Fire Prevention Commission"
  code_name: "NFPA 1; NFPA 101; listed NFPA standards"
  edition_year: 2024
  incorporated_by_reference: true
authority:
  adopting_authority_id: "ahj:usa-de:state-fire-prevention-commission"
  enforcing_authority_model: "state"
  interpretation_authority_id: "ahj:usa-de:state-fire-marshal"
dates:
  adoption_date: "2025-12-16"
  effective_date: "2026-01-01"
  operative_date: "2025-12-31"
  mandatory_date: null
  replacement_date: null
applicability:
  date_trigger: "project_submission_date"
  applies_to:
    - "new_construction"
    - "alteration"
    - "repair"
    - "addition"
    - "change_of_occupancy"
    - "fire_protection_systems"
    - "fire_alarm_systems"
    - "exits"
    - "existing_buildings_where_regulations_apply"
  exclusions:
    - "one_and_two_family_dwellings_for_some_plan_review_functions"
    - "farming_buildings_for_some_plan_review_functions"
  special_conditions:
    - "All projects submitted on or after 2025-12-31 must comply with the 2026 DSFPR."
    - "The 2026 update list includes NFPA 1 2024 and NFPA 101 2024."
    - "The 2026 update list includes NFPA 70, National Electrical Code, 2023."
transition:
  exists: true
  rule_id: "date-rule:usa-de:fire:2026-dsfpr"
  start_date: "2025-12-31"
  end_date: null
  prior_code_allowed: false
  prior_code_condition: "Projects submitted on or after the cutover must comply with the 2026 DSFPR."
amendments:
  state_amended: true
  amendment_set_ids:
    - "amendment-set:usa-de:fire:2026-dsfpr"
  amendment_source_ids:
    - "src:usa-de:firemarshal:2026-dsfpr-update"
    - "src:usa-de:regulations:title1:701-2025-final"
provenance:
  source_ids:
    - "src:usa-de:code:title16:ch66-sc01"
    - "src:usa-de:code:title16:ch66-sc02"
    - "src:usa-de:firemarshal:2026-dsfpr-update"
    - "src:usa-de:regulations:title1:701-2025-final"
  field_sources:
    state_code.name: ["src:usa-de:firemarshal:2026-dsfpr-update"]
    state_code.edition_label: ["src:usa-de:firemarshal:2026-dsfpr-update"]
    authority.adopting_authority_id: ["src:usa-de:code:title16:ch66-sc01"]
    authority.interpretation_authority_id: ["src:usa-de:code:title16:ch66-sc02"]
    dates.effective_date: ["src:usa-de:firemarshal:2026-dsfpr-update"]
    transition.rule_id: ["src:usa-de:firemarshal:2026-dsfpr-update"]
verification:
  status: "partially_verified"
  confidence: 0.93
  notes: "Fire authority and current transition notice verified; full 2026 regulation text was not exhaustively parsed."

adoption_id: "adoption:usa-de:electrical:nec-2023-dsfpr"
state_id: "US-DE"
code_family: "electrical"
status: "partially_verified"
state_code:
  name: "Electrical code under State Fire Prevention Commission NEC adoption and applicable local code"
  edition_label: "NFPA 70, National Electrical Code, 2023"
  codification: "1 DE Admin. Code 701; 24 DE Admin. Code 1400"
base_model_code:
  publisher: "NFPA"
  code_name: "NFPA 70, National Electrical Code"
  edition_year: 2023
  incorporated_by_reference: true
authority:
  adopting_authority_id: "ahj:usa-de:state-fire-prevention-commission"
  enforcing_authority_model: "hybrid_state_local"
  interpretation_authority_id: "ahj:usa-de:state-fire-marshal"
dates:
  adoption_date: "2025-12-16"
  effective_date: "2026-01-01"
  operative_date: "2025-12-31"
  mandatory_date: null
  replacement_date: null
applicability:
  date_trigger: "project_submission_date_or_local_permit_rule"
  applies_to:
    - "electrical_work"
    - "fire_alarm_systems"
    - "construction_projects_where_nec_applies"
  exclusions: []
  special_conditions:
    - "Electrical work is also subject to applicable local building codes."
    - "The current professional-licensing/regulatory pathway should be rechecked against the official Delaware Administrative Code before final verification."
transition:
  exists: true
  rule_id: "date-rule:usa-de:electrical:nec-2023-dsfpr"
  start_date: "2025-12-31"
  end_date: null
  prior_code_allowed: null
  prior_code_condition: "Fire-regulation project-submission cutover verified; local electrical permitting cutover unresolved."
amendments:
  state_amended: true
  amendment_set_ids:
    - "amendment-set:usa-de:electrical:nec-2023"
  amendment_source_ids:
    - "src:usa-de:firemarshal:2026-dsfpr-update"
    - "src:usa-de:regulations:title24:1400"
provenance:
  source_ids:
    - "src:usa-de:firemarshal:2026-dsfpr-update"
    - "src:usa-de:regulations:title24:1400"
  field_sources:
    state_code.edition_label: ["src:usa-de:firemarshal:2026-dsfpr-update"]
    dates.effective_date: ["src:usa-de:firemarshal:2026-dsfpr-update"]
    applicability.special_conditions: ["src:usa-de:regulations:title24:1400"]
verification:
  status: "partially_verified"
  confidence: 0.86
  notes: "NEC edition verified from fire-regulation update; official electrical licensing/admin-code text needs a focused extraction."
```

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

Delaware uses different cutover rules by code family. General building-code cutovers are local. Plumbing has a DPH-led statewide adoption structure and local enforcement. Energy code updates use the Delaware Energy Office process and a 12-month local implementation window. Fire regulations have a statewide effective date and a project-submission cutover announced by the State Fire Marshal. Elevator emergency communication has a statutory compliance date.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-de:building:local-adoption | county/municipal building codes | local_operative_date | jurisdiction-specific | Local ordinance adoption and local permit administration | null | src:usa-de:code:title16:ch76 | verified |
| date-rule:usa-de:plumbing:dph-ipc-adoption | Delaware Plumbing Code | statutory_update_rule | within 1 calendar year of IPC issuance | Division of Public Health adopts and enforces the most recent IPC | null | src:usa-de:code:title16:ch79 | verified |
| date-rule:usa-de:energy:2026-update | Delaware State Energy Conservation Code | effective_date_with_local_implementation | 2026-04-11; 12-month local implementation window | DNREC/Delaware Energy Office promulgation of revised 7 DE Admin. Code 2101 | null | src:usa-de:dnrec:building-energy-codes; src:usa-de:dnrec:energy-order-2026-cce-0002 | verified |
| date-rule:usa-de:energy:statutory-zne-residential-2025 | Delaware energy statute | statutory_milestone | 2025-12-31 | New residential building construction zero-net-energy-capable milestone in § 7602(c) | unresolved | src:usa-de:code:title16:ch76; src:usa-de:dnrec:energy-order-2026-cce-0002 | partially_verified |
| date-rule:usa-de:energy:statutory-zne-commercial-2030 | Delaware energy statute | statutory_milestone | 2030-12-31 | New commercial building construction zero-net-energy-capable milestone in § 7602(c) | unresolved | src:usa-de:code:title16:ch76; src:usa-de:dnrec:energy-order-2026-cce-0002 | partially_verified |
| date-rule:usa-de:fire:2026-dsfpr | Delaware State Fire Prevention Regulations | project_submission_cutover | 2025-12-31 project-submission cutover; 2026-01-01 regulation effective date | Projects submitted to the Office of the State Fire Marshal on or after 2025-12-31 | false | src:usa-de:firemarshal:2026-dsfpr-update | verified |
| date-rule:usa-de:electrical:nec-2023-dsfpr | NEC reference in 2026 DSFPR | project_submission_cutover | 2025-12-31 project-submission cutover; 2026-01-01 regulation effective date | Projects submitted under the 2026 DSFPR | null | src:usa-de:firemarshal:2026-dsfpr-update | partially_verified |
| date-rule:usa-de:elevator:emergency-communication | passenger elevators | statutory_compliance_date | 1997-01-26 | Passenger elevator use or permit/CO/license issuance or renewal | false | src:usa-de:code:title16:ch87 | verified |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| energy | local implementation of 2026 7 DE Admin. Code 2101 update | 2026-03-10 | 2026-03-10 | 2026-04-11 | null | null | active | src:usa-de:dnrec:building-energy-codes; src:usa-de:dnrec:energy-order-2026-cce-0002 | Confirm whether each local jurisdiction uses the 12-month window from publication, effective date, or another local cutover event. |
| energy | zero-net-energy-capable implementation | null | null | null | null | 2030-12-31 | monitoring | src:usa-de:code:title16:ch76; src:usa-de:dnrec:energy-order-2026-cce-0002 | The statute contains 2025 and 2030 milestones, but the 2026 DNREC order deferred implementation. |
| fire_operational | future DSFPR/NFPA updates | null | null | null | null | null | monitoring | src:usa-de:regulations:title1:701-2025-final; src:usa-de:firemarshal:2026-dsfpr-update | Monitor State Fire Prevention Commission notices and State Fire Marshal update letters. |
| plumbing | next IPC adoption | null | null | null | null | null | monitoring | src:usa-de:code:title16:ch79; src:usa-de:dhss:plumbing-program | Statute requires DPH adoption of the most recent IPC within one calendar year of issuance; current program page identifies 2021 IPC. |
| local_building | county and municipal code updates | null | null | null | null | null | monitoring | src:usa-de:code:title16:ch76; src:usa-de:new-castle:county-laws-code | Local IBC/IRC/IMC/IFGC/IEBC editions require jurisdiction-by-jurisdiction tracking. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-de:001 | energy | local stretch codes | local adoption | Local jurisdictions may adopt stretch codes but cannot fully replace or supersede the state energy code without a formal Delaware Energy Office process at least 6 months before adoption. | src:usa-de:code:title16:ch76 | verified |
| applicability-rule:usa-de:002 | energy | residential and commercial construction | 2026 state energy update | Residential uses 2024 IECC; commercial/high-rise uses 2024 IECC or ASHRAE 90.1-2022 paths; local implementation is required within 12 months. | src:usa-de:dnrec:building-energy-codes; src:usa-de:dnrec:energy-order-2026-cce-0002 | verified |
| applicability-rule:usa-de:003 | fire_operational | plan review | new buildings, additions, occupancy changes, fire systems, alarms, flammable/combustible liquids and similar scopes | State Fire Marshal review applies to specified construction and fire-safety work, with listed exceptions. | src:usa-de:code:title16:ch66-sc02 | verified |
| applicability-rule:usa-de:004 | plumbing | local plumbing modifications | local additions or modifications | Local plumbing additions or modifications must be submitted to DPH for review and approval before implementation. | src:usa-de:code:title16:ch79 | verified |
| applicability-rule:usa-de:005 | elevator | passenger elevators | occupied elevator use and permit/CO/license issuance | Passenger elevators need direct emergency communication; private residential elevators are excepted; county and municipal rules may effectuate the section. | src:usa-de:code:title16:ch87 | verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** mixed local code adoption plus statewide plumbing, fire-prevention, and energy-code amendments.

**Where amendments are published:** Delaware Code Online, Delaware Regulations, DNREC/Delaware Energy Office rulemaking materials, DHSS Plumbing Program materials, Office of the State Fire Marshal notices, and local code portals.

**Amendment parsing status:** partial.

### 5.2 State Amendment Sources

| Amendment Set ID | Code Family | Publication Format | Source IDs | Parsed? | Notes |
| --- | --- | --- | --- | --- | --- |
| amendment-set:usa-de:building:local-codes | building | county or municipal code text | src:usa-de:code:title16:ch76; src:usa-de:new-castle:county-laws-code | partial | Local adoption is jurisdiction-specific; New Castle County parsed only as an example. |
| amendment-set:usa-de:plumbing:delaware-amendments | plumbing | DPH plumbing regulations and program materials | src:usa-de:code:title16:ch79; src:usa-de:dhss:plumbing-program | partial | DPH program page confirms 2021 IPC with Delaware-specific amendments; full amendment text not fully parsed. |
| amendment-set:usa-de:energy:2026-2101 | energy | 7 DE Admin. Code 2101 final regulation and DNREC order | src:usa-de:dnrec:building-energy-codes; src:usa-de:dnrec:energy-order-2026-cce-0002; src:usa-de:dnrec:energy-hearing-2025-r-cce-0008 | partial | 2026 order confirms model codes, effective date, EV appendix, and solar-ready appendix adoption. |
| amendment-set:usa-de:fire:2026-dsfpr | fire_operational | 1 DE Admin. Code 701 final regulation and State Fire Marshal update | src:usa-de:regulations:title1:701-2025-final; src:usa-de:firemarshal:2026-dsfpr-update | partial | State Fire Marshal update lists common NFPA codes and the cutover rule; final regulatory text requires full extraction. |
| amendment-set:usa-de:electrical:nec-2023 | electrical | 2026 DSFPR references and electrical licensing regulations | src:usa-de:firemarshal:2026-dsfpr-update; src:usa-de:regulations:title24:1400 | partial | NEC 2023 edition verified through 2026 DSFPR update; official 24 DE Admin. Code extraction remains open. |
| amendment-set:usa-de:elevator:emergency-communication | elevator | state statute | src:usa-de:code:title16:ch87 | yes | Narrow statutory communication requirement only. |

### 5.3 High-Impact State Amendments

| Code Family | Section | Amendment Type | Summary | Source IDs | Confidence |
| --- | --- | --- | --- | --- | --- |
| energy | 7 DE Admin. Code 2101 | replace | Delaware adopted a 2026 energy package based on 2024 IECC for residential and 2024 IECC or ASHRAE 90.1-2022 options for commercial/high-rise residential, effective 2026-04-11. | src:usa-de:dnrec:building-energy-codes; src:usa-de:dnrec:energy-order-2026-cce-0002 | 0.97 |
| energy | 2024 IECC Appendix RE | add | The 2026 energy order adopts the 2024 IECC EV-capable Appendix RE with Delaware-specific amendments. | src:usa-de:dnrec:energy-order-2026-cce-0002; src:usa-de:code:title16:ch80 | 0.94 |
| energy | 2024 IECC Appendix CB | add | The 2026 energy order adopts the 2024 IECC solar-ready Appendix CB with Delaware-specific amendments. | src:usa-de:dnrec:energy-order-2026-cce-0002; src:usa-de:code:title16:ch76 | 0.94 |
| energy | 16 Del. C. § 7602(c) | unresolved_implementation | Delaware Code contains 2025 and 2030 zero-net-energy-capable milestones, but the 2026 DNREC order stated that implementation was not appropriate at that time and deferred implementation to future legislative or executive action. | src:usa-de:code:title16:ch76; src:usa-de:dnrec:energy-order-2026-cce-0002 | 0.88 |
| fire_operational | 1 DE Admin. Code 701 | replace | 2026 DSFPR update became effective 2026-01-01; commonly used standards include NFPA 1 2024, NFPA 101 2024, and NFPA 70 2023. | src:usa-de:firemarshal:2026-dsfpr-update; src:usa-de:regulations:title1:701-2025-final | 0.93 |
| plumbing | Delaware Plumbing Code | amend | The official program page identifies the 2021 IPC as the Delaware Plumbing Code with Delaware-specific amendments. | src:usa-de:dhss:plumbing-program; src:usa-de:code:title16:ch79 | 0.94 |
| elevator | 16 Del. C. § 8701 | add | Delaware requires direct emergency communication in passenger elevators and allows counties or municipalities to adopt implementing rules. | src:usa-de:code:title16:ch87 | 0.95 |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-de"
  model: "hybrid"
  enforcing_entities:
    - "county_council"
    - "levy_court"
    - "municipal_government"
    - "division_of_public_health"
    - "local_plumbing_authorities"
    - "state_fire_marshal"
    - "state_fire_marshal_deputy"
    - "local_energy_code_officials"
  required_officials:
    - "county_or_municipal_building_official"
    - "local_code_official"
    - "plumbing_program_or_local_plumbing_inspector"
    - "state_fire_marshal"
    - "state_fire_marshal_deputy"
  state_reserved_activities:
    - "fire_prevention_regulation"
    - "state_fire_plan_review"
    - "state_energy_code_floor_setting"
    - "state_plumbing_code_adoption_and_local_modification_review"
  local_reserved_or_local_primary_activities:
    - "general_building_code_adoption"
    - "local_permit_administration"
    - "local_energy_code_implementation"
    - "local_building_code_enforcement"
  source_ids:
    - "src:usa-de:code:title16:ch76"
    - "src:usa-de:code:title16:ch79"
    - "src:usa-de:code:title16:ch66-sc02"
    - "src:usa-de:dnrec:building-energy-codes"
  verification_status: "partially_verified"
  confidence: 0.90
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-de"
  model: "mixed_by_code_family"
  applies_to_code_families:
    - "building"
    - "residential"
    - "existing_building"
    - "mechanical"
    - "plumbing"
    - "fuel_gas"
    - "electrical"
    - "energy"
    - "fire_operational"
  code_family_rules:
    building:
      model: "jurisdiction_specific_local_adoption"
      approval_required: null
      approving_authority_id: null
      filing_required: null
      source_ids:
        - "src:usa-de:code:title16:ch76"
    plumbing:
      model: "state_floor_with_local_modifications_requiring_review"
      approval_required: true
      approving_authority_id: "ahj:usa-de:dph-plumbing-program"
      filing_required: true
      source_ids:
        - "src:usa-de:code:title16:ch79"
    energy:
      model: "state_floor_with_local_stretch_codes"
      approval_required: "required_for_full_replacement_or_supersession"
      approving_authority_id: "ahj:usa-de:delaware-energy-office"
      filing_required: true
      source_ids:
        - "src:usa-de:code:title16:ch76"
        - "src:usa-de:dnrec:building-energy-codes"
    fire_operational:
      model: "state_minimum_with_more_stringent_local_rules_if_consistent"
      approval_required: null
      approving_authority_id: null
      filing_required: null
      source_ids:
        - "src:usa-de:code:title16:ch66-sc01"
    electrical:
      model: "state_nec_reference_plus_applicable_local_code"
      approval_required: null
      approving_authority_id: null
      filing_required: null
      source_ids:
        - "src:usa-de:firemarshal:2026-dsfpr-update"
        - "src:usa-de:regulations:title24:1400"
  registry_exists: false
  registry_source_ids: []
  legal_basis_source_ids:
    - "src:usa-de:code:title16:ch76"
    - "src:usa-de:code:title16:ch79"
    - "src:usa-de:code:title16:ch66-sc01"
  verification_status: "partially_verified"
  confidence: 0.88
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement and local amendment authority are separate in Delaware. Local governments may adopt and enforce many general construction codes, but the state still reserves or sets floors for specific code families. Plumbing local modifications are subject to DPH review and approval. Energy local stretch codes are allowed, but full local replacement of the state energy code requires a formal Delaware Energy Office process. Fire prevention regulations are statewide minimum rules, while more stringent local rules may govern if they are not inconsistent with the state code and good engineering practice.

### 6.4 Known Local Amendment Registries

| Registry ID | Scope | Maintained By | Source ID | Machine Readable? | Parsed? | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| registry:usa-de:local-building-amendments | county/municipal building codes | county and municipal governments | src:usa-de:code:title16:ch76 | no | no | No statewide machine-readable local building amendment registry was found in the sources reviewed. |
| registry:usa-de:plumbing-local-modifications | local plumbing additions/modifications | Division of Public Health review process | src:usa-de:code:title16:ch79 | no | no | Statute requires submission to DPH for review and approval, but a public consolidated registry was not located. |
| registry:usa-de:energy-local-stretch-codes | local energy stretch codes | Delaware Energy Office process | src:usa-de:code:title16:ch76 | no | no | Statute allows stretch codes and limits full replacement; a consolidated live registry was not found. |

### 6.5 Municipality-Specific Known Amendments

| Jurisdiction | Code Family | Amendment Set ID | Approval Status | Effective Date | Source IDs | Parsed? |
| --- | --- | --- | --- | --- | --- | --- |
| New Castle County | building/residential/existing/mechanical/fuel_gas/plumbing/energy | local-amendment:usa-de:new-castle:chapter-6 | local_example_only | multiple | src:usa-de:new-castle:county-laws-code | partial |
| unresolved statewide set | local building amendments | local-amendment:usa-de:unresolved | unresolved | null | none | no |

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: state_plus_local_required

Jurisdiction stack:

```text
Address
  -> State of Delaware
  -> County
  -> Municipality or unincorporated county area
  -> Local building-code jurisdiction
  -> Local plumbing jurisdiction and/or DPH plumbing program
  -> State Fire Marshal district / State Fire Marshal plan-review scope
  -> Energy-code enforcing jurisdiction
  -> Applicable local amendments, stretch codes, and permit cutover rules
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | Delaware Code Online, Title 16 | src:usa-de:code:title16:home | statewide | periodic | active |
| County | County governments and Chapter 76 local code authority | src:usa-de:code:title16:ch76 | statewide | periodic | active |
| Municipality | municipal governments and local code portals | src:usa-de:code:title16:ch76 | statewide | periodic | active |
| Fire AHJ | State Fire Prevention Commission / State Fire Marshal | src:usa-de:code:title16:ch66-sc01; src:usa-de:code:title16:ch66-sc02 | statewide | periodic | active |
| Plumbing AHJ | Division of Public Health Plumbing Program and local jurisdictions | src:usa-de:code:title16:ch79; src:usa-de:dhss:plumbing-program | statewide | periodic | active |
| Energy AHJ | Delaware Energy Office and local code-enforcement jurisdictions | src:usa-de:dnrec:building-energy-codes | statewide | triennial review plus local implementation | active |
| Special District | unresolved | none | unknown | unknown | pending |

### 7.3 AHJ Contact Data

No AHJ contact data was populated for this pass.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Source Type | Publisher | URL | Accessed Date | Snapshot ID | Checksum | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-de:code:title16:home | Title 16 - Health and Safety | code_index | Delaware Code Online | https://delcode.delaware.gov/title16/ | 2026-06-25 | snapshot:de-title16-2026-06-25 | null | active |
| src:usa-de:code:title16:ch66-sc01 | Chapter 66, Subchapter I - State Fire Prevention Commission | statute | Delaware Code Online | https://delcode.delaware.gov/title16/c066/sc01/index.html | 2026-06-25 | snapshot:de-ch66-sc01-2026-06-25 | null | active |
| src:usa-de:code:title16:ch66-sc02 | Chapter 66, Subchapter II - State Fire Marshal | statute | Delaware Code Online | https://delcode.delaware.gov/title16/c066/sc02/index.html | 2026-06-25 | snapshot:de-ch66-sc02-2026-06-25 | null | active |
| src:usa-de:code:title16:ch76 | Chapter 76 - County or Municipal Building, Plumbing, Electrical and Other Codes | statute | Delaware Code Online | https://delcode.delaware.gov/title16/c076/index.html | 2026-06-25 | snapshot:de-ch76-2026-06-25 | null | active |
| src:usa-de:code:title16:ch79 | Chapter 79 - Plumbing Code | statute | Delaware Code Online | https://delcode.delaware.gov/title16/c079/index.html | 2026-06-25 | snapshot:de-ch79-2026-06-25 | null | active |
| src:usa-de:code:title16:ch80 | Chapter 80 - Electric Vehicle Charging Infrastructure for Residential Dwellings | statute | Delaware Code Online | https://delcode.delaware.gov/title16/c080/index.html | 2026-06-25 | snapshot:de-ch80-2026-06-25 | null | active |
| src:usa-de:code:title16:ch87 | Chapter 87 - Elevators | statute | Delaware Code Online | https://delcode.delaware.gov/title16/c087/index.html | 2026-06-25 | snapshot:de-ch87-2026-06-25 | null | active |
| src:usa-de:dhss:plumbing-program | Plumbing Program | agency_page | Delaware Department of Health and Social Services, Division of Public Health | https://dhss.delaware.gov/dph/homepage/about/sections/hsp/licenses-and-permits/plumbing/ | 2026-06-25 | snapshot:de-dhss-plumbing-program-2026-06-25 | null | active |
| src:usa-de:dnrec:building-energy-codes | Building Energy Codes | agency_page | DNREC Division of Climate, Coastal and Energy | https://dnrec.delaware.gov/climate-coastal-energy/efficiency/building-energy-codes/ | 2026-06-25 | snapshot:de-dnrec-building-energy-codes-2026-06-25 | null | active |
| src:usa-de:dnrec:energy-order-2026-cce-0002 | Secretary's Order No. 2026-CCE-0002 | secretary_order_pdf | DNREC | https://documents.dnrec.delaware.gov/Admin/Orders/Secretarys-Order-No-2026-CCE-0002.pdf | 2026-06-25 | snapshot:de-dnrec-order-2026-cce-0002 | null | active |
| src:usa-de:dnrec:energy-hearing-2025-r-cce-0008 | Public Hearing: Regulations for State Energy Conservation Code | rulemaking_page | DNREC | https://dnrec.delaware.gov/events/public-hearing-regulations-for-state-energy-conservation-code-2/ | 2026-06-25 | snapshot:de-dnrec-energy-hearing-2025-r-cce-0008 | null | active |
| src:usa-de:regulations:title1:701-2025-final | 1 DE Admin. Code 701, Administration and Enforcement - Final Regulation, 29 DE Reg. 509 | final_regulation | Delaware Register of Regulations | https://regulations.delaware.gov/register/december2025/final/29%20DE%20Reg%20509%2012-01-25.htm | 2026-06-25 | snapshot:de-reg-701-final-2025-12 | null | active |
| src:usa-de:firemarshal:2026-dsfpr-update | 2026 Delaware State Fire Prevention Regulations update notice | agency_pdf | Delaware Office of the State Fire Marshal | https://statefiremarshal.delaware.gov/wp-content/uploads/sites/110/2025/12/2026-DSFPR-update.pdf | 2026-06-25 | snapshot:de-firemarshal-2026-dsfpr-update | null | active |
| src:usa-de:regulations:title24:1400 | Board of Electrical Examiners, 24 DE Admin. Code 1400 | regulation | Delaware Regulations | https://regulations.delaware.gov/AdminCode/title24/1400/ | 2026-06-25 | snapshot:de-title24-1400-2026-06-25 | null | active |
| src:usa-de:new-castle:county-laws-code | County Laws and Code | local_code_page | New Castle County | https://www.newcastlede.gov/229/County-Laws-Code | 2026-06-25 | snapshot:new-castle-code-page-2026-06-25 | null | active |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-de:code:title16:ch76 | official_html | State statute verifies local code authority and energy-code framework, but it does not list current local building-code editions. | use for authority and state floor; do not infer local editions |
| src:usa-de:code:title16:ch79 | official_html | Statute verifies DPH plumbing authority and local review process, but current edition details were taken from the official DPH program page. | pair with DPH program source for current edition |
| src:usa-de:dhss:plumbing-program | official_agency_page | Program page confirms 2021 IPC with Delaware-specific amendments, but it is not itself the full regulatory amendment text. | use for current edition; extract regulations separately before verified status |
| src:usa-de:dnrec:energy-order-2026-cce-0002 | official_pdf | PDF text was machine-readable and spot-checked visually; it is an order, not a consolidated regulatory code text. | use for adoption/effective dates and adopted model-code package |
| src:usa-de:regulations:title1:701-2025-final | official_html | Official final-regulation page is the canonical rulemaking publication, but full text extraction was incomplete in this pass. | pair with State Fire Marshal update letter and re-extract before verified status |
| src:usa-de:firemarshal:2026-dsfpr-update | official_agency_pdf | Agency update letter summarizes the 2026 DSFPR and common NFPA references, but is not the consolidated final regulation text. | use for cutover and high-level code list; verify detailed section text in 1 DE Admin. Code 701 |
| src:usa-de:regulations:title24:1400 | official_html | Official electrical administrative-code source was identified, but full section extraction was not completed. | use cautiously; re-extract before final electrical verification |
| src:usa-de:new-castle:county-laws-code | local_official_page | Local county page is useful as a local example only and does not establish statewide code editions. | do not globalize to statewide matrix |

### 8.3 Supplemental Sources

| Source ID | Title | Source Type | Publisher | URL | Use | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| src:usa-de:supp:doe-energy-code-status | Delaware Current State Code | federal_status_page | U.S. Department of Energy, Building Energy Codes Program | https://www.energycodes.gov/status/states/delaware | cross-check | DOE status page aligns with the 2026 Delaware energy adoption but is supplemental to DNREC and Delaware Code sources. |

### 8.4 Source Extraction Metadata

| Source ID | Parser | Parser Version | Extracted At | Extraction Confidence | Requires OCR | Requires Browser Automation | Manual Review Required |
| --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-de:code:title16:ch66-sc01 | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.97 | no | no | no |
| src:usa-de:code:title16:ch66-sc02 | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.97 | no | no | no |
| src:usa-de:code:title16:ch76 | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.98 | no | no | no |
| src:usa-de:code:title16:ch79 | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.97 | no | no | no |
| src:usa-de:code:title16:ch80 | web_search | 2026-06-25 | 2026-06-25T00:00:00Z | 0.82 | no | no | yes |
| src:usa-de:code:title16:ch87 | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.96 | no | no | no |
| src:usa-de:dhss:plumbing-program | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.93 | no | no | yes |
| src:usa-de:dnrec:building-energy-codes | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.97 | no | no | no |
| src:usa-de:dnrec:energy-order-2026-cce-0002 | web_pdf_screenshot | 2026-06-25 | 2026-06-25T00:00:00Z | 0.96 | no | yes | no |
| src:usa-de:dnrec:energy-hearing-2025-r-cce-0008 | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.90 | no | no | yes |
| src:usa-de:regulations:title1:701-2025-final | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.84 | no | no | yes |
| src:usa-de:firemarshal:2026-dsfpr-update | web_pdf_screenshot | 2026-06-25 | 2026-06-25T00:00:00Z | 0.95 | no | yes | yes |
| src:usa-de:regulations:title24:1400 | web_search | 2026-06-25 | 2026-06-25T00:00:00Z | 0.78 | no | no | yes |
| src:usa-de:new-castle:county-laws-code | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.93 | no | no | yes |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| adoption:usa-de:building:local-codes | authority.adopting_authority_id | ahj:usa-de:local-building-authorities | verified | 0.96 | src:usa-de:code:title16:ch76 | Chapter 76 supports local code adoption and enforcement. |
| adoption:usa-de:building:local-codes | state_code.edition_label | jurisdiction-specific | partially_verified | 0.90 | src:usa-de:code:title16:ch76; src:usa-de:new-castle:county-laws-code | Statewide IBC/IRC edition not identified; local example confirms jurisdiction-specific adoption. |
| adoption:usa-de:plumbing:delaware-plumbing-code | state_code.edition_label | 2021 IPC with Delaware-specific amendments | verified | 0.94 | src:usa-de:dhss:plumbing-program | DPH program page identifies the current IPC edition and amendment status. |
| adoption:usa-de:plumbing:delaware-plumbing-code | authority.adopting_authority_id | ahj:usa-de:dph-plumbing-program | verified | 0.96 | src:usa-de:code:title16:ch79 | DPH authority and local review process are statutory. |
| adoption:usa-de:energy:2024-iecc-ashrae-2022 | dates.adoption_date | 2026-03-10 | verified | 0.95 | src:usa-de:dnrec:energy-order-2026-cce-0002 | Secretary's order issuance date. |
| adoption:usa-de:energy:2024-iecc-ashrae-2022 | dates.effective_date | 2026-04-11 | verified | 0.97 | src:usa-de:dnrec:building-energy-codes; src:usa-de:dnrec:energy-order-2026-cce-0002 | Agency page and order align. |
| adoption:usa-de:energy:2024-iecc-ashrae-2022 | transition.rule_id | date-rule:usa-de:energy:2026-update | partially_verified | 0.91 | src:usa-de:code:title16:ch76; src:usa-de:dnrec:building-energy-codes | 12-month local implementation window verified; exact local cutover remains open. |
| adoption:usa-de:fire:2026-dsfpr | state_code.edition_label | 2026 edition | verified | 0.93 | src:usa-de:firemarshal:2026-dsfpr-update; src:usa-de:regulations:title1:701-2025-final | Current edition and transition notice verified from official sources. |
| adoption:usa-de:fire:2026-dsfpr | dates.effective_date | 2026-01-01 | verified | 0.93 | src:usa-de:firemarshal:2026-dsfpr-update | State Fire Marshal update letter. |
| adoption:usa-de:electrical:nec-2023-dsfpr | state_code.edition_label | NFPA 70, National Electrical Code, 2023 | verified | 0.90 | src:usa-de:firemarshal:2026-dsfpr-update | Listed in 2026 DSFPR update. |
| local-amendment-rule:usa-de | code_family_rules.plumbing.approval_required | true | verified | 0.94 | src:usa-de:code:title16:ch79 | Local plumbing additions/modifications require DPH review and approval. |
| local-amendment-rule:usa-de | code_family_rules.energy.model | state_floor_with_local_stretch_codes | verified | 0.95 | src:usa-de:code:title16:ch76 | Stretch-code and full-replacement rules are statutory. |
| local-amendment-rule:usa-de | code_family_rules.fire_operational.model | state_minimum_with_more_stringent_local_rules_if_consistent | verified | 0.94 | src:usa-de:code:title16:ch66-sc01 | State fire regulations are minimum requirements and more stringent consistent local rules may govern. |
| date-rule:usa-de:elevator:emergency-communication | Date / Period | 1997-01-26 | verified | 0.95 | src:usa-de:code:title16:ch87 | Statutory compliance date. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Every `src:usa-de:*` ID used in the body appears in Section 8. |
| All authority IDs are internally consistent | pass | Authority IDs used in records are defined in Section 2 or explicitly unresolved. |
| Core authority fields replaced | pass | Frontmatter and primary authority placeholders are replaced. |
| Current statewide code families have explicit rows | pass | Building/local, plumbing, energy, fire, electrical, elevator, and unresolved accessibility rows are explicit. |
| Building and operational fire code are separated | pass | Local building authority is separate from Commission/State Fire Marshal fire authority. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Energy, fire, electrical, and elevator dates are separated. |
| Effective dates are ISO formatted | pass | Verified dates are in YYYY-MM-DD format. |
| No impossible date sequences introduced | pass | Fire, energy, and elevator sequences were checked for obvious contradictions. |
| Transition rules have explicit trigger conditions | pass | Local, DPH, energy, fire, electrical, and elevator triggers are listed. |
| Permit or project-submission logic captured where applicable | pass | Fire and electrical project-submission cutovers are captured; local building and energy cutovers are flagged as local. |
| Local enforcement model classified | pass | Classified as hybrid. |
| Local amendment model classified | pass | Classified as mixed by code family. |
| AHJ confirmation metadata present | pass | AHJ resolution section requires local lookup and states that no contact data was populated. |
| Official-source caveats captured | pass | Imperfect extraction and local-example caveats are listed. |
| Unresolved items remain explicit | pass | Accessibility, local edition coverage, electrical extraction, and local registry gaps remain in the open-issue queue. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-de:001 | high | local building code editions | Current IBC, IRC, IEBC, IMC, IFGC, and local amendment editions must be collected for each county and relevant municipality. | Extract county and municipal code pages and ordinances; normalize per jurisdiction. | null | null | open |
| issue:usa-de:002 | medium | electrical administrative-code extraction | NEC edition was verified through the fire-regulation update, but the official 24 DE Admin. Code 1400 text should be fully extracted for licensing and local-code interplay. | Re-extract 24 DE Admin. Code 1400 and Board of Electrical Examiners materials. | null | null | open |
| issue:usa-de:003 | medium | energy local implementation cutover | The 2026 energy update and 12-month local implementation window are verified, but local permit-trigger dates remain unresolved. | Confirm DNREC guidance and each enforcing jurisdiction's implementation date. | null | null | open |
| issue:usa-de:004 | medium | accessibility authority | Statewide accessibility construction-code authority was not resolved beyond local code structures and federal background law. | Review Title 22, Title 29, state ADA/accessibility rules, and local building-code amendments. | null | null | open |
| issue:usa-de:005 | low | local amendment registry | No consolidated registry for county/municipal building, plumbing, or energy amendments was found. | Search county, municipal, DPH, and Delaware Energy Office registries or guidance pages. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-de:title16-ch76 | src:usa-de:code:title16:ch76 | html_diff | monthly | local code authority or energy-code framework changes | 2026-06-25 | active |
| watch:usa-de:title16-ch79 | src:usa-de:code:title16:ch79 | html_diff | monthly | DPH plumbing-code adoption or local modification process changes | 2026-06-25 | active |
| watch:usa-de:dnrec-energy-code | src:usa-de:dnrec:building-energy-codes | html_diff | monthly | 7 DE Admin. Code 2101 guidance, local implementation, or training changes | 2026-06-25 | active |
| watch:usa-de:dnrec-energy-order | src:usa-de:dnrec:energy-order-2026-cce-0002 | rulemaking_followup | quarterly | new secretary order or local implementation guidance | 2026-06-25 | active |
| watch:usa-de:fire-regs | src:usa-de:regulations:title1:701-2025-final | register_notice_scan | monthly | State Fire Prevention Commission regulation updates | 2026-06-25 | active |
| watch:usa-de:firemarshal-update | src:usa-de:firemarshal:2026-dsfpr-update | pdf_or_page_diff | monthly | 2026 DSFPR update revisions or new transition notices | 2026-06-25 | active |
| watch:usa-de:dhss-plumbing | src:usa-de:dhss:plumbing-program | html_diff | monthly | plumbing code edition or amendment updates | 2026-06-25 | active |
| watch:usa-de:local-codes | src:usa-de:new-castle:county-laws-code | local_page_sample | quarterly | local code edition or amendment changes in tracked jurisdictions | 2026-06-25 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-25 | Replaced Delaware draft with a partially verified source-backed report. Added DPH plumbing authority, current Delaware Plumbing Code edition, 2026 Delaware energy-code update, 2026 DSFPR/fire-code transition, NEC 2023 fire-regulation reference, and explicit unresolved queues for local building editions and accessibility. | report:usa-de | src:usa-de:code:title16:ch66-sc01; src:usa-de:code:title16:ch66-sc02; src:usa-de:code:title16:ch76; src:usa-de:code:title16:ch79; src:usa-de:dnrec:building-energy-codes; src:usa-de:dnrec:energy-order-2026-cce-0002; src:usa-de:firemarshal:2026-dsfpr-update; src:usa-de:dhss:plumbing-program | system | State status upgraded from draft to partially_verified after validation checks; unresolved items remain explicit. |
