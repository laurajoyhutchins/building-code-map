---
state:
  state_id: "US-NV"
  name: "Nevada"
  abbreviation: "NV"
report:
  report_id: "state-report:usa-nv"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-26"
  last_verified: null
  reviewed_by: null
risk:
  overall_confidence: 0.58 # 0.00 - 1.00
  risk_flags:
    - "statewide_building_code_scope_is_limited"
    - "local_code_editions_not_statewide_normalized"
    - "local_amendment_registry_not_parsed"
    - "fire_code_transition_requires_final_nac_crosscheck"
    - "elevator_code_editions_not_fully_parsed"
  open_questions_count: 7

---

# State Building Code Authority Report: Nevada

## 1. Executive Summary

- **Authority model:** Nevada uses a split authority model. The State Public Works Division / State Public Works Board is the state building-code authority for state-owned buildings and buildings held in trust, while city and county governments adopt and enforce most local building codes for non-state property. Energy conservation standards are a statewide minimum administered by the Governor's Office of Energy, and fire/life-safety code authority is separately administered by the Nevada State Fire Marshal. Source IDs: src:usa-nv:spwd-permitting; src:usa-nv:nac-341-045; src:usa-nv:nrs-244-3675; src:usa-nv:nrs-268-413; src:usa-nv:nrs-278-580; src:usa-nv:nrs-701-220; src:usa-nv:nrs-477-030.

- **Statewide code status:** A single statewide building code for all private construction was not confirmed. State-owned buildings are governed by the code set adopted under NAC 341.045. Energy conservation is confirmed as a statewide minimum based on the 2024 International Energy Conservation Code effective 2024-08-18. The State Fire Marshal announced enforcement of 2024 codes effective 2026-01-01, and Approved Regulation R205-24 contains 2024-code updates to the fire-code regulations. Source IDs: src:usa-nv:nac-341-045; src:usa-nv:reg-r072-24; src:usa-nv:goe-state-adoption-status; src:usa-nv:sfm-hot-topics; src:usa-nv:reg-r205-24.

- **Local enforcement model:** Counties and incorporated cities are the default local construction-code adopters/enforcers outside the state-property track. NRS 278.585 confirms that persons and political subdivisions must comply with the applicable city or county building code, subject to statutory exceptions. Source IDs: src:usa-nv:nrs-244-3675; src:usa-nv:nrs-268-413; src:usa-nv:nrs-278-580; src:usa-nv:nrs-278-585.

- **Local amendment posture:** Local governments appear to have local code-adoption and modification authority, but it is bounded by state floors and limits. Confirmed limits include the statewide energy-code minimum, the rule that electrical-code modifications may not reduce National Electrical Code standards, and statutory limits on local residential sprinkler mandates. A statewide local-amendment registry was not found in this pass. Source IDs: src:usa-nv:nrs-701-220; src:usa-nv:nrs-278-583; src:usa-nv:nrs-278-586.

- **Known transition periods or pending changes:** Energy-code adoption is confirmed effective 2024-08-18. State Fire Marshal 2024-code enforcement is confirmed effective 2026-01-01. Local permit-date and grace-period rules vary by jurisdiction and were not normalized statewide. Source IDs: src:usa-nv:goe-state-adoption-status; src:usa-nv:sfm-hot-topics.

- **Production readiness:** partially_ready_for_core_state_authority. The report is source-backed for the main state, local, energy, and fire authority tracks, but it remains unsuitable for production use without local jurisdiction normalization and final register/NAC consolidation review.

### Key Findings

```yaml
---
key_findings:
- topic: State building-code authority
  finding: The State Public Works Division / Board administers code compliance for
    state-owned land and adopts minimum code standards for state buildings; it is
    not confirmed as a statewide private-building code authority.
  confidence: 0.78
  source_ids:
  - src:usa-nv:spwd-permitting
  - src:usa-nv:nrs-341-100
  - src:usa-nv:nrs-341-105
  - src:usa-nv:nac-341-045
- topic: Local building-code authority
  finding: Counties and cities have statutory authority to adopt building, electrical,
    plumbing, safety, and related codes, subject to state limits.
  confidence: 0.72
  source_ids:
  - src:usa-nv:nrs-244-3675
  - src:usa-nv:nrs-268-413
  - src:usa-nv:nrs-278-580
- topic: Energy code
  finding: The Governor's Office of Energy adopted the 2024 IECC effective 2024-08-18
    as the statewide energy-conservation code minimum.
  confidence: 0.82
  source_ids:
  - src:usa-nv:goe-state-adoption-status
  - src:usa-nv:nrs-701-220
  - src:usa-nv:nac-701-185
- topic: Fire code authority
  finding: The Nevada State Fire Marshal is the distinct fire/life-safety authority;
    2024-code enforcement is announced effective 2026-01-01, with R205-24 updating
    NAC 477 references.
  confidence: 0.7
  source_ids:
  - src:usa-nv:nrs-477-030
  - src:usa-nv:sfm-hot-topics
  - src:usa-nv:reg-r205-24
- topic: Electrical code authority
  finding: 'Electrical-code rules are split: local governments administer local construction
    codes; state buildings use the NEC code path under NAC 341.045; NRS 278.583 limits
    local NEC modifications.'
  confidence: 0.55
  source_ids:
  - src:usa-nv:nac-341-045
  - src:usa-nv:nrs-278-583
- topic: Local enforcement
  finding: Enforcement is primarily local for non-state property, with separate state
    enforcement on state-owned property and specialized state tracks for fire, energy,
    and conveyances.
  confidence: 0.62
  source_ids:
  - src:usa-nv:nrs-278-585
  - src:usa-nv:spwd-permitting
  - src:usa-nv:nrs-477-030
  - src:usa-nv:nrs-455c
- topic: Local amendments
  finding: Local amendments are allowed within state-law limits, but no complete statewide
    local-amendment filing or registry process was verified.
  confidence: 0.45
  source_ids:
  - src:usa-nv:nrs-701-220
  - src:usa-nv:nrs-278-583
  - src:usa-nv:nrs-278-586
- topic: Effective / operative date rule
  finding: Confirmed statewide dates are limited to the 2024 IECC effective date and
    the State Fire Marshal 2024-code enforcement date; local permit/application transition
    dates remain jurisdiction-specific.
  confidence: 0.5
  source_ids:
  - src:usa-nv:goe-state-adoption-status
  - src:usa-nv:sfm-hot-topics
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-nv:spwd |
| Authority name | Nevada State Public Works Division / State Public Works Board |
| Authority type | state_building_official_for_state_property |
| Legal basis | NRS 341.100; NRS 341.105; NRS 341.1455; NAC 341.045 |
| Role | Oversees code compliance and enforcement on state-owned land, issues permits for state-owned facilities, and adopts minimum standards for state buildings. |
| Enforcement model | Direct state enforcement for state-owned buildings and buildings held in trust; local enforcement model applies outside this track. |
| Source IDs | src:usa-nv:spwd-permitting; src:usa-nv:nrs-341-100; src:usa-nv:nrs-341-105; src:usa-nv:nrs-341-1455; src:usa-nv:nac-341-045 |
| Verification status | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-nv:spwd; ahj:usa-nv:local-building-governments | State Public Works Division / Board for state buildings; counties and cities for local buildings | SPWD adopts/enforces state-building standards; cities/counties adopt/enforce local building codes. | NRS 341.100; NAC 341.045; NRS 244.3675; NRS 268.413; NRS 278.580 | src:usa-nv:nac-341-045; src:usa-nv:spwd-permitting; src:usa-nv:nrs-244-3675; src:usa-nv:nrs-268-413; src:usa-nv:nrs-278-580 | partially_verified |
| Residential | ahj:usa-nv:spwd; ahj:usa-nv:local-building-governments | State Public Works Division / Board; local governments | SPWD state-building code set includes the IRC; local residential code edition varies by jurisdiction. | NAC 341.045; NRS 278.580 | src:usa-nv:nac-341-045; src:usa-nv:nrs-278-580 | partially_verified |
| Existing Building / Rehabilitation | ahj:usa-nv:spwd; ahj:usa-nv:sfm | State Public Works Division / Board; Nevada State Fire Marshal | SPWD state-building code set includes the IEBC; State Fire Marshal amendments also reference the IEBC for fire/life-safety scope. | NAC 341.045; NAC 477; R205-24 | src:usa-nv:nac-341-045; src:usa-nv:reg-r205-24 | partially_verified |
| Mechanical | ahj:usa-nv:spwd; ahj:usa-nv:local-building-governments; ahj:usa-nv:sfm | SPWD; local governments; State Fire Marshal | Nevada state-building and fire-rule paths substitute the Uniform Mechanical Code where model-code references would otherwise point to the International Mechanical Code. | NAC 341.045; NAC 477; local enabling statutes | src:usa-nv:nac-341-045; src:usa-nv:reg-r205-24; src:usa-nv:nrs-244-3675; src:usa-nv:nrs-268-413 | partially_verified |
| Plumbing | ahj:usa-nv:spwd; ahj:usa-nv:local-building-governments; ahj:usa-nv:sfm | SPWD; local governments; State Fire Marshal | Nevada state-building and fire-rule paths substitute the Uniform Plumbing Code where model-code references would otherwise point to the International Plumbing Code. | NAC 341.045; NAC 477; local enabling statutes | src:usa-nv:nac-341-045; src:usa-nv:reg-r205-24; src:usa-nv:nrs-244-3675; src:usa-nv:nrs-268-413 | partially_verified |
| Fuel Gas | ahj:usa-nv:sfm; ahj:usa-nv:local-building-governments | State Fire Marshal; local governments | Fire-code amendments replace International Fuel Gas Code references with NFPA 54/ANSI Z223.1 or the current LPG Board version; local fuel-gas code administration was not fully normalized. | NAC 477; NRS 477.030 | src:usa-nv:reg-r205-24; src:usa-nv:nrs-477-030 | partially_verified |
| Electrical | ahj:usa-nv:spwd; ahj:usa-nv:local-building-governments | SPWD; local governments | SPWD state-building path uses the National Electrical Code; local governments may adopt NEC modifications only if standards are not reduced. | NAC 341.045; NRS 278.583 | src:usa-nv:nac-341-045; src:usa-nv:nrs-278-583 | partially_verified |
| Energy | ahj:usa-nv:goe | Nevada Governor's Office of Energy | Adopts minimum energy-conservation standards; local governments must incorporate and enforce the standards and may adopt more stringent standards. | NRS 701.220; NAC 701.185 | src:usa-nv:nrs-701-220; src:usa-nv:nac-701-185; src:usa-nv:goe-state-adoption-status | partially_verified |
| Fire - construction references | ahj:usa-nv:sfm | Nevada State Fire Marshal | Adopts fire/life-safety regulations that include construction-code references and Nevada amendments. | NRS 477.030; NAC 477; R205-24 | src:usa-nv:nrs-477-030; src:usa-nv:reg-r205-24; src:usa-nv:sfm-hot-topics | partially_verified |
| Fire - operational / prevention code | ahj:usa-nv:sfm | Nevada State Fire Marshal | Enforces fire prevention and operational fire-code requirements within the State Fire Marshal scope, with local delegation possible under statute. | NRS 477.030; NAC 477 | src:usa-nv:nrs-477-030; src:usa-nv:sfm-hot-topics; src:usa-nv:reg-r205-24 | partially_verified |
| Accessibility | ahj:usa-nv:spwd; ahj:usa-nv:local-building-governments | SPWD; local governments | SPWD state-building code path substitutes ICC A117.1-2009 for ICC A117.1-2017 in the adopted IBC path; local accessibility administration and federal ADA overlay were not normalized. | NAC 341.045 | src:usa-nv:nac-341-045 | partially_verified |
| Elevator / Conveyance | ahj:usa-nv:dir-mcs | Nevada Department of Business and Industry, Division of Industrial Relations, Mechanical Compliance Section | Administers elevator and related conveyance compliance under NRS/NAC 455C; exact adopted edition list requires final parsing of NAC 455C.500. | NRS 455C; NAC 455C.500; NAC 455C.516 | src:usa-nv:nrs-455c; src:usa-nv:nac-455c-500; src:usa-nv:nac-455c-516; src:usa-nv:dir-mcs-information | partially_verified |

### 2.3 Authority Hierarchy Notes

Nevada should be modeled as a hybrid jurisdiction rather than as a single statewide building-code adopter. The first resolver question is whether the project is on state-owned land or in a building held in trust for any state agency. If yes, the SPWD state-building code path applies and SPWD permitting, plan review, and inspection are the controlling state-building track. If not, local city/county code adoption and enforcement generally control the building-code path, subject to statewide energy minimums, State Fire Marshal scope, elevator/conveyance regulation, and other special statutory limits.

The Governor's Office of Energy track is not merely advisory. NRS 701.220 and the GOE adoption page support a statewide minimum energy-code rule. Local governments may adopt stricter standards, but weaker local standards are not supported by the energy-code sources reviewed. Source IDs: src:usa-nv:nrs-701-220; src:usa-nv:goe-state-adoption-status.

The State Fire Marshal track is separate from local building-code adoption. NRS 477.030 gives the State Fire Marshal authority over fire-prevention and life-safety regulations in specified occupancies and contexts, while also allowing local-government enforcement arrangements in certain circumstances. Source IDs: src:usa-nv:nrs-477-030; src:usa-nv:sfm-fire-protection-engineering; src:usa-nv:sfm-hot-topics.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-nv:001 | ahj:usa-nv:spwd | enforces_codes_on | state-owned land and state-owned facilities | src:usa-nv:spwd-permitting; src:usa-nv:nrs-341-100; src:usa-nv:nrs-341-105 | partially_verified |
| edge:usa-nv:002 | ahj:usa-nv:spwd | adopts_minimum_standards_for | state buildings, except statutory exclusions | src:usa-nv:nac-341-045; src:usa-nv:reg-r072-24 | partially_verified |
| edge:usa-nv:003 | ahj:usa-nv:local-building-governments | adopt_and_enforce | local building, electrical, plumbing, safety, and related codes | src:usa-nv:nrs-244-3675; src:usa-nv:nrs-268-413; src:usa-nv:nrs-278-580; src:usa-nv:nrs-278-585 | partially_verified |
| edge:usa-nv:004 | ahj:usa-nv:goe | sets_minimum_energy_standards_for | local building-code governments and statewide building energy conservation | src:usa-nv:nrs-701-220; src:usa-nv:goe-state-adoption-status | partially_verified |
| edge:usa-nv:005 | ahj:usa-nv:local-building-governments | may_adopt_more_stringent_energy_standards | local energy-code provisions reported to GOE | src:usa-nv:nrs-701-220 | partially_verified |
| edge:usa-nv:006 | ahj:usa-nv:sfm | adopts_and_enforces | fire prevention, life safety, and fire-code construction references within SFM scope | src:usa-nv:nrs-477-030; src:usa-nv:reg-r205-24 | partially_verified |
| edge:usa-nv:007 | ahj:usa-nv:sfm | may_delegate_or_coordinate_enforcement_with | local governments where statute allows and qualifications are met | src:usa-nv:nrs-477-030 | partially_verified |
| edge:usa-nv:008 | ahj:usa-nv:dir-mcs | regulates | elevators and related conveyances | src:usa-nv:nrs-455c; src:usa-nv:nac-455c-500; src:usa-nv:nac-455c-516 | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | State-building minimum standards; local building codes outside state-property track | 2024 IBC for state buildings; local IBC edition varies | 2024 for state-building path | state_buildings_verified_local_varies | 2024-11-15 | null | null | null | SPWD permit path for state-owned facilities; local permit/application rules vary by jurisdiction. | src:usa-nv:nac-341-045; src:usa-nv:spwd-adopted-codes-2024; src:usa-nv:reg-r072-24; src:usa-nv:nrs-278-580 |
| Residential | State-building minimum standards; local residential codes outside state-property track | 2024 IRC for state-building path; local IRC edition varies | 2024 for state-building path | state_buildings_verified_local_varies | 2024-11-15 | null | null | null | SPWD permit path for state-owned facilities; local permit/application rules vary by jurisdiction. | src:usa-nv:nac-341-045; src:usa-nv:spwd-adopted-codes-2024; src:usa-nv:reg-r072-24; src:usa-nv:nrs-278-580 |
| Existing Building / Rehabilitation | State-building existing-building standards; SFM fire/life-safety references | 2024 IEBC | 2024 | state_buildings_and_sfm_scope_partially_verified | 2024-11-15 | null | null | null | SPWD date path unresolved; SFM enforcement of 2024 codes announced for 2026-01-01. | src:usa-nv:nac-341-045; src:usa-nv:reg-r072-24; src:usa-nv:reg-r205-24; src:usa-nv:sfm-hot-topics |
| Mechanical | State-building mechanical standards; local mechanical codes outside state-property track | 2024 Uniform Mechanical Code in state-building and SFM substitution paths | 2024 UMC | state_buildings_and_sfm_scope_partially_verified | 2024-11-15 | null | null | null | IMC references are replaced by UMC in the reviewed state-building/fire-rule sources; local dates vary. | src:usa-nv:nac-341-045; src:usa-nv:reg-r072-24; src:usa-nv:reg-r205-24 |
| Plumbing | State-building plumbing standards; local plumbing codes outside state-property track | 2024 Uniform Plumbing Code in state-building and SFM substitution paths | 2024 UPC | state_buildings_and_sfm_scope_partially_verified | 2024-11-15 | null | null | null | IPC references are replaced by UPC in the reviewed state-building/fire-rule sources; local dates vary. | src:usa-nv:nac-341-045; src:usa-nv:reg-r072-24; src:usa-nv:reg-r205-24 |
| Fuel Gas | Fire-rule fuel-gas substitution; local fuel-gas path unresolved | NFPA 54/ANSI Z223.1 National Fuel Gas Code or current LPG Board version under SFM amendment path | 2024 for NFPA 54 path | partially_verified | null | 2026-01-01 | null | null | SFM 2024-code enforcement date confirmed; local gas/LPG split requires additional review. | src:usa-nv:reg-r205-24; src:usa-nv:sfm-hot-topics |
| Electrical | State-building electrical standards; local electrical code authority | National Electrical Code for state-building path; local NEC modifications limited by NRS 278.583 | 2023 NEC for state-building path; local current-edition handling unresolved | partially_verified | 2024-11-15 | null | null | null | Local electrical-code modifications may not reduce NEC standards. | src:usa-nv:nac-341-045; src:usa-nv:reg-r072-24; src:usa-nv:nrs-278-583 |
| Energy | Nevada energy-conservation code | 2024 International Energy Conservation Code | 2024 IECC | statewide_minimum_verified | 2024-08-18 | 2024-08-18 | null | null | Local building-code governments must incorporate and enforce the minimum standards; stricter local standards are allowed. | src:usa-nv:goe-state-adoption-status; src:usa-nv:goe-building-energy-codes; src:usa-nv:nrs-701-220; src:usa-nv:nac-701-185 |
| Fire - construction references | State Fire Marshal fire/life-safety construction-code references | 2024 IBC, 2024 IEBC, 2024 IWUIC, 2024 UMC, 2024 UPC references in R205-24 | 2024 | sfm_scope_partially_verified | null | 2026-01-01 | null | null | SFM announced 2024-code enforcement effective 2026-01-01; final consolidated NAC cross-check remains open. | src:usa-nv:nrs-477-030; src:usa-nv:reg-r205-24; src:usa-nv:sfm-hot-topics |
| Fire - operational / prevention code | State Fire Marshal fire code | 2024 International Fire Code with Nevada amendments and NFPA references | 2024 IFC | sfm_scope_partially_verified | null | 2026-01-01 | null | null | SFM announced 2024-code enforcement effective 2026-01-01; local fire-AHJ transition rules are not normalized. | src:usa-nv:nrs-477-030; src:usa-nv:reg-r205-24; src:usa-nv:sfm-hot-topics |
| Accessibility | State-building accessibility reference path | 2024 IBC accessibility provisions with ICC A117.1-2009 substitution in state-building path | 2024 IBC / 2009 A117.1 substitution | state_buildings_partially_verified | 2024-11-15 | null | null | null | Federal ADA and local accessibility amendment paths were not normalized. | src:usa-nv:nac-341-045; src:usa-nv:reg-r072-24 |
| Elevator / Conveyance | Nevada elevator and conveyance safety program | ASME A17.1 / adopted standards under NAC 455C.500 | edition_list_unresolved | specialized_state_program_partially_verified | null | null | null | null | Owners and inspectors must follow adopted conveyance standards; exact adopted edition list requires final NAC 455C.500 parsing. | src:usa-nv:nrs-455c; src:usa-nv:nac-455c-500; src:usa-nv:nac-455c-516; src:usa-nv:dir-mcs-information |

### 3.2 Adoption Records

#### adoption:usa-nv:spwd-2024-building-codes

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-nv:spwd |
| Scope | State buildings and state-owned facilities, subject to statutory exclusions. |
| Code families | Building; Residential; Existing Building; Mechanical; Plumbing; Electrical; Energy; Accessibility |
| Base/model codes captured | 2024 IBC; 2024 IRC; 2024 IEBC; 2024 Uniform Plumbing Code; 2024 Uniform Mechanical Code; 2023 NEC; 2024 IECC; ASHRAE 90.1-2022; ICC A117.1-2009 substitution. |
| Adoption date | 2024-11-15 |
| Effective date | null |
| Operative date | null |
| Mandatory date | null |
| Date basis | Approved Regulation R072-24 and NAC 341.045 code text identify the adopted code set. The exact effective/operative sequence should be verified in the codified NAC and regulation history before elevating to verified. |
| Source IDs | src:usa-nv:nac-341-045; src:usa-nv:spwd-adopted-codes-2024; src:usa-nv:reg-r072-24 |
| Confidence | 0.70 |

#### adoption:usa-nv:goe-2024-iecc

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-nv:goe |
| Scope | Statewide minimum energy-conservation standard for buildings, with local incorporation and enforcement. |
| Code families | Energy |
| Base/model codes captured | 2024 International Energy Conservation Code |
| Adoption date | 2024-08-18 |
| Effective date | 2024-08-18 |
| Operative date | null |
| Mandatory date | null |
| Date basis | GOE adoption-status page states that the 2024 IECC was adopted effective 2024-08-18 by reference pursuant to NAC 701.185. |
| Source IDs | src:usa-nv:goe-state-adoption-status; src:usa-nv:nrs-701-220; src:usa-nv:nac-701-185 |
| Confidence | 0.82 |

#### adoption:usa-nv:sfm-2024-fire-codes

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-nv:sfm |
| Scope | State Fire Marshal fire-prevention, operational, and fire/life-safety code scope. |
| Code families | Fire - operational / prevention code; Fire - construction references; Fuel Gas; Mechanical; Plumbing; Existing Building / Rehabilitation |
| Base/model codes captured | 2024 IFC; 2024 IBC references; 2024 IEBC references; 2024 IWUIC; 2024 UMC; 2024 UPC; NFPA 54/ANSI Z223.1 2024 substitution for fuel gas. |
| Adoption date | null |
| Effective date | 2026-01-01 |
| Operative date | null |
| Mandatory date | null |
| Date basis | State Fire Marshal official notice states 2024-code enforcement begins 2026-01-01; Approved Regulation R205-24 supplies amendment text but final codified NAC timing still requires review. |
| Source IDs | src:usa-nv:sfm-hot-topics; src:usa-nv:reg-r205-24; src:usa-nv:nrs-477-030 |
| Confidence | 0.70 |

#### adoption:usa-nv:local-building-code-enabling-record

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-nv:local-building-governments |
| Scope | Local building-code adoption and enforcement by counties and cities outside the state-property path. |
| Code families | Building; Residential; Electrical; Plumbing; Mechanical; Safety; local fire/building interfaces |
| Base/model codes captured | Local editions vary and were not normalized statewide. |
| Adoption date | null |
| Effective date | null |
| Operative date | null |
| Mandatory date | null |
| Date basis | Local statutes provide authority; each jurisdiction's ordinance and local transition rules must be parsed separately. |
| Source IDs | src:usa-nv:nrs-244-3675; src:usa-nv:nrs-268-413; src:usa-nv:nrs-278-580; src:usa-nv:nrs-278-585 |
| Confidence | 0.62 |

#### adoption:usa-nv:elevator-conveyance-standards

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-nv:dir-mcs |
| Scope | Elevators and related conveyances. |
| Code families | Elevator / Conveyance |
| Base/model codes captured | ASME A17.1 and other standards adopted by reference under NAC 455C.500; exact edition list unresolved. |
| Adoption date | null |
| Effective date | null |
| Operative date | null |
| Mandatory date | null |
| Date basis | DIR and NAC 455C sources confirm an adopted-standards program but do not complete the edition matrix in this pass. |
| Source IDs | src:usa-nv:nrs-455c; src:usa-nv:nac-455c-500; src:usa-nv:nac-455c-516; src:usa-nv:dir-mcs-information |
| Confidence | 0.46 |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

Two statewide date rules were confirmed with enough confidence to record: the 2024 IECC effective date of 2024-08-18 and the State Fire Marshal's announced 2024-code enforcement date of 2026-01-01. For local building-code adoption, Nevada statutes support local adoption/enforcement authority, but local permit-date, permit-application-date, issuance-date, and grace-period logic must be parsed jurisdiction by jurisdiction. SPWD state-building code dates should be cross-checked against the codified NAC history before this report is marked verified.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-nv:energy-2024-iecc-effective | 2024 IECC statewide minimum | effective_date | 2024-08-18 | GOE adoption by reference pursuant to NAC 701.185 | unresolved for local transition/grace periods | src:usa-nv:goe-state-adoption-status; src:usa-nv:nac-701-185 | partially_verified |
| date-rule:usa-nv:sfm-2024-codes-enforcement | State Fire Marshal 2024-code package | enforcement_date | 2026-01-01 | State Fire Marshal announced enforcement of 2024 codes | unresolved for open permit/project transition | src:usa-nv:sfm-hot-topics; src:usa-nv:reg-r205-24 | partially_verified |
| date-rule:usa-nv:local-code-transition-varies | Local building-code adoptions | local_transition_rule | jurisdiction-specific | Local ordinance, permit application date, issuance date, or code-effective ordinance date | unresolved statewide | src:usa-nv:nrs-244-3675; src:usa-nv:nrs-268-413; src:usa-nv:nrs-278-580 | unresolved |
| date-rule:usa-nv:construction-begins-energy-trigger | Energy-conservation codes in counties/cities with statutory population trigger | construction_start_trigger | statutory effective date of local energy-code ordinance | Construction begins on or after the effective date of the applicable ordinance | unresolved for current local ordinances | src:usa-nv:nrs-278-581 | partially_verified |
| date-rule:usa-nv:spwd-2024-code-effective | SPWD 2024 state-building code set | effective_or_operative_date | null | Code set adoption under NAC 341.045 and R072-24 | unresolved | src:usa-nv:nac-341-045; src:usa-nv:reg-r072-24 | unresolved |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Energy | 2027 IECC cycle or later | null | null | null | null | null | watch | src:usa-nv:nrs-701-220; src:usa-nv:goe-state-adoption-status | NRS 701.220 requires recurring adoption of the most recent IECC; next-cycle timing should be monitored. |
| Fire | Next SFM code package after 2024 codes | null | null | null | null | null | watch | src:usa-nv:sfm-hot-topics; src:usa-nv:reg-r205-24 | No future post-2026 SFM code package was verified. |
| State buildings | Next SPWD code update after 2024 package | null | null | null | null | null | watch | src:usa-nv:nac-341-045; src:usa-nv:reg-r072-24 | No future state-building code package was verified. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-nv:state-owned-facilities | Building; Residential; Existing; Mechanical; Plumbing; Electrical; Energy; Accessibility | State-owned facilities and state-owned land | Project involves state-owned facilities or state-owned land | SPWD permitting, plan review, inspection, and NAC 341.045 code set apply. | src:usa-nv:spwd-permitting; src:usa-nv:nac-341-045 | partially_verified |
| applicability-rule:usa-nv:local-building-code-compliance | Building; Residential; trade codes | Local city/county building-code jurisdiction | Applicable local city or county has adopted a building code | Persons and political subdivisions must comply with the appropriate city/county building code, subject to statutory exceptions. | src:usa-nv:nrs-278-585; src:usa-nv:nrs-278-580 | partially_verified |
| applicability-rule:usa-nv:energy-statewide-minimum | Energy | Buildings subject to state energy-conservation standards | GOE code adoption and local incorporation/enforcement | Local standards must at least meet statewide energy-conservation standards; stricter local standards are allowed. | src:usa-nv:nrs-701-220; src:usa-nv:goe-state-adoption-status | partially_verified |
| applicability-rule:usa-nv:residential-sprinklers | Fire; Residential | Local residential sprinkler mandates | Local governing body attempts to require sprinklers in new residential units | NRS 278.586 places specific constraints on local sprinkler requirements for units below 5,000 square feet and allows requirements for larger units. | src:usa-nv:nrs-278-586 | partially_verified |
| applicability-rule:usa-nv:school-exception | Building | School facilities | Project falls under NRS 393.110 exception referenced in NRS 278.585 | The school-construction exception is flagged but not parsed in this pass. | src:usa-nv:nrs-278-585 | unresolved |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** split_by_authority_track

**Where amendments are published:** State-building amendments and substitutions are published in NAC 341.045 and related SPWD adopted-code documents. Fire-code amendments are published in NAC 477 and Nevada Register amendment packages such as R205-24. Energy-code adoption is published through GOE rules and adoption notices under NRS 701.220 / NAC 701.185. Local amendments are published in local ordinances and regional/local amendment packages, but no statewide registry was verified.

**Amendment parsing status:** partially_started

### 5.2 State Amendment Sources

| Amendment Source ID | Authority ID | Applies To | Publication Path | Parsing Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| amend-source:usa-nv:spwd-nac-341-045 | ahj:usa-nv:spwd | State-building code set | NAC 341.045 and SPWD adopted-code PDF | partial | src:usa-nv:nac-341-045; src:usa-nv:spwd-adopted-codes-2024; src:usa-nv:reg-r072-24 | Confirms code substitutions and state-building code editions; final effective-date path unresolved. |
| amend-source:usa-nv:sfm-r205-24 | ahj:usa-nv:sfm | Fire-code and fire/life-safety construction references | Nevada Register Approved Regulation R205-24 and NAC 477 | partial | src:usa-nv:reg-r205-24; src:usa-nv:nac-477-281; src:usa-nv:nac-477-283 | Confirms 2024-code replacement language and several high-impact substitutions; final codified NAC cross-check remains open. |
| amend-source:usa-nv:goe-nac-701 | ahj:usa-nv:goe | Energy conservation | NRS 701.220; NAC 701.185; GOE adoption page | partial | src:usa-nv:nrs-701-220; src:usa-nv:nac-701-185; src:usa-nv:goe-state-adoption-status | Confirms statewide 2024 IECC adoption/effective date and local ability to exceed the minimum. |
| amend-source:usa-nv:local-ordinances | ahj:usa-nv:local-building-governments | Local building, trade, and fire-code amendments | County/city ordinances and regional amendment packages | not_started | src:usa-nv:nrs-244-3675; src:usa-nv:nrs-268-413; src:usa-nv:nrs-278-580 | Local edition/amendment inventory not parsed statewide. |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Scope | Amendment Summary | Source IDs | Confidence | Status |
| --- | --- | --- | --- | --- | --- | --- |
| amend:usa-nv:spwd-upc-substitution | Plumbing | State-building path | NAC 341.045 state-building code path substitutes the Uniform Plumbing Code where the model-code path would otherwise use the International Plumbing Code. | src:usa-nv:nac-341-045; src:usa-nv:reg-r072-24 | 0.72 | partially_verified |
| amend:usa-nv:spwd-umc-substitution | Mechanical | State-building path | NAC 341.045 state-building code path substitutes the Uniform Mechanical Code where the model-code path would otherwise use the International Mechanical Code. | src:usa-nv:nac-341-045; src:usa-nv:reg-r072-24 | 0.72 | partially_verified |
| amend:usa-nv:spwd-nec-substitution | Electrical | State-building path | NAC 341.045 state-building code path substitutes the National Electrical Code for the International Code Council electrical-code reference. | src:usa-nv:nac-341-045; src:usa-nv:reg-r072-24 | 0.68 | partially_verified |
| amend:usa-nv:spwd-a117-2009 | Accessibility | State-building path | NAC 341.045 state-building code path uses ICC A117.1-2009 in place of the newer ICC A117.1 reference in the adopted IBC path. | src:usa-nv:nac-341-045 | 0.66 | partially_verified |
| amend:usa-nv:sfm-fuel-gas-nfpa54 | Fuel Gas; Fire | State Fire Marshal fire-code path | R205-24 replaces International Fuel Gas Code references with NFPA 54/ANSI Z223.1 National Fuel Gas Code, 2024 edition, or the current LPG Board version adopted under NAC 590.610. | src:usa-nv:reg-r205-24 | 0.72 | partially_verified |
| amend:usa-nv:sfm-umc-upc-substitutions | Mechanical; Plumbing; Fire | State Fire Marshal fire-code path | R205-24 replaces International Mechanical Code and International Plumbing Code references with the 2024 Uniform Mechanical Code and 2024 Uniform Plumbing Code. | src:usa-nv:reg-r205-24 | 0.72 | partially_verified |
| amend:usa-nv:sfm-high-rise-55ft | Fire; Building | State Fire Marshal fire/life-safety path | R205-24 includes a high-rise definition amendment tied to an occupied floor more than 55 feet above the lowest level of fire-department vehicle access. | src:usa-nv:reg-r205-24 | 0.62 | partially_verified |
| amend:usa-nv:local-nec-floor | Electrical | Local electrical-code modification | NRS 278.583 allows local NEC modifications only if they do not reduce the standards set by the NEC. | src:usa-nv:nrs-278-583 | 0.70 | partially_verified |
| amend:usa-nv:local-energy-floor | Energy | Local energy-code modification | NRS 701.220 allows local governments to adopt higher or more stringent energy standards and requires local incorporation/enforcement of the statewide minimum. | src:usa-nv:nrs-701-220; src:usa-nv:goe-state-adoption-status | 0.80 | partially_verified |
| amend:usa-nv:local-sprinkler-limits | Fire; Residential | Local residential sprinkler mandates | NRS 278.586 limits when local governments may require automatic fire sprinklers in new single-family residential units below 5,000 square feet. | src:usa-nv:nrs-278-586 | 0.68 | partially_verified |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-nv"
  model: "local_adoption_and_enforcement_with_state_specialized_overlays"
  enforcing_entities:
    - "counties for county building-code jurisdictions"
    - "incorporated cities for city building-code jurisdictions"
    - "State Public Works Division for state-owned land and state-owned facilities"
    - "Nevada State Fire Marshal and delegated or coordinated local fire authorities for fire/life-safety scope"
    - "Governor's Office of Energy sets energy-code minimums; local governments incorporate and enforce them"
    - "Division of Industrial Relations Mechanical Compliance Section for elevator/conveyance scope"
  required_officials:
    - "local building official or code official by local ordinance, not normalized statewide"
    - "SPWD building official for state-owned facilities"
    - "State Fire Marshal or approved local fire authority in SFM scope"
  state_reserved_activities:
    - "SPWD permitting and code enforcement for state-owned facilities"
    - "State Fire Marshal fire prevention and life-safety regulation within statutory scope"
    - "statewide minimum energy-conservation standard adoption"
    - "elevator/conveyance permitting and inspection program"
  source_ids:
    - "src:usa-nv:nrs-244-3675"
    - "src:usa-nv:nrs-268-413"
    - "src:usa-nv:nrs-278-580"
    - "src:usa-nv:nrs-278-585"
    - "src:usa-nv:spwd-permitting"
    - "src:usa-nv:nrs-477-030"
    - "src:usa-nv:nrs-701-220"
    - "src:usa-nv:nrs-455c"
  verification_status: "partially_verified"
  confidence: 0.62
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-nv"
  model: "local_amendments_allowed_with_state_floors_and_topic_specific_limits"
  applies_to_code_families:
    - "building"
    - "residential"
    - "electrical"
    - "plumbing"
    - "mechanical"
    - "energy"
    - "fire and life safety, subject to State Fire Marshal and local fire authority scope"
  approval_required: "varies by topic; no general statewide local-amendment approval workflow was confirmed"
  approving_authority_id: null
  filing_required: "energy standards more stringent than the statewide minimum must be reported to the GOE Director; other filing duties unresolved"
  registry_exists: null
  registry_source_ids: []
  legal_basis_source_ids:
    - "src:usa-nv:nrs-244-3675"
    - "src:usa-nv:nrs-268-413"
    - "src:usa-nv:nrs-278-580"
    - "src:usa-nv:nrs-701-220"
    - "src:usa-nv:nrs-278-583"
    - "src:usa-nv:nrs-278-586"
  verification_status: "partially_verified"
  confidence: 0.45
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement and local amendment authority should not be collapsed. Nevada local governments can function as local building-code AHJs and can adopt local codes or amendments within statutory limits, but statewide specialized authorities still apply. For example, a city or county building department may enforce a local building-code ordinance, while the GOE statewide energy-code minimum, SFM fire/life-safety rules, and SPWD state-building jurisdiction may independently affect the same project depending on ownership and scope.

### 6.4 Known Local Amendment Registries

| Registry ID | Jurisdiction / Region | Registry Name | Source ID | Coverage | Status | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| registry:usa-nv:statewide-local-amendments | statewide | none confirmed | none | statewide local amendments | unresolved | No statewide local amendment filing registry was verified. |
| registry:usa-nv:local-ordinances | counties and cities | local code ordinances | src:usa-nv:nrs-244-3675; src:usa-nv:nrs-268-413; src:usa-nv:nrs-278-580 | jurisdiction-specific | pending | Local ordinance parsing is required for production AHJ resolution. |

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendments were normalized into this report. Regional and local amendment packages may exist, but they require a separate jurisdiction-level extraction pass and should not be treated as statewide rules.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: partially_started

Jurisdiction stack:

```text
Address
  -> Determine whether the project is on state-owned land or in a state-owned facility
  -> If state-owned, route to SPWD state-building code path
  -> Determine county and incorporated city / unincorporated county status
  -> Determine local building department and local adopted code edition
  -> Determine fire AHJ and whether State Fire Marshal scope or local fire authority scope controls
  -> Apply statewide energy-conservation minimum and any stricter local energy standard
  -> Apply elevator/conveyance program if applicable
  -> Apply school, special-district, utility, floodplain, and other overlays if applicable
  -> Return applicable state code adoption records and applicable local amendment records
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | not selected | none | statewide | unresolved | pending |
| County | not selected | none | statewide | unresolved | pending |
| Municipality | not selected | none | statewide | unresolved | pending |
| Fire District | not selected | none | statewide | unresolved | pending |
| State-Owned Property | not selected | src:usa-nv:spwd-permitting | statewide state-property path | unresolved | pending |
| Special District | not selected | none | statewide | unresolved | pending |

### 7.3 AHJ Contact Data

No AHJ contact data was populated for this pass. Production use requires a county/city building department directory, local fire-AHJ mapping, and SPWD/SFM contact metadata for state or fire-scope projects.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Publisher | Type | URL / Locator | Key Fields Supported | Status |
| --- | --- | --- | --- | --- | --- | --- |
| src:usa-nv:nrs-244-3675 | NRS 244.3675 - County regulation of buildings, structures, property and adoption of codes | Nevada Legislature / Legislative Counsel Bureau | statute | https://www.leg.state.nv.us/NRS/NRS-244.html#NRS244Sec3675 | County local building-code authority | official_registry_entry |
| src:usa-nv:nrs-268-413 | NRS 268.413 - City regulation of buildings, structures, property and adoption of codes | Nevada Legislature / Legislative Counsel Bureau | statute | https://www.leg.state.nv.us/NRS/NRS-268.html#NRS268Sec413 | City local building-code authority | official_registry_entry |
| src:usa-nv:nrs-278-580 | NRS 278.580 - Adoption of building code and enforcement rules | Nevada Legislature / Legislative Counsel Bureau | statute | https://www.leg.state.nv.us/NRS/NRS-278.html#NRS278Sec580 | Local building-code adoption/enforcement authority | official_registry_entry |
| src:usa-nv:nrs-278-581 | NRS 278.581 - Construction and energy codes for counties/cities with population trigger | Nevada Legislature / Legislative Counsel Bureau | statute | https://www.leg.state.nv.us/NRS/NRS-278.html#NRS278Sec581 | Energy-code construction-start trigger and local enforcement context | official_registry_entry |
| src:usa-nv:nrs-278-583 | NRS 278.583 - National Electrical Code standards and local modifications | Nevada Legislature / Legislative Counsel Bureau | statute | https://www.leg.state.nv.us/NRS/NRS-278.html#NRS278Sec583 | NEC local modification floor | official_registry_entry |
| src:usa-nv:nrs-278-585 | NRS 278.585 - Compliance with city/county building codes | Nevada Legislature / Legislative Counsel Bureau | statute | https://www.leg.state.nv.us/NRS/NRS-278.html#NRS278Sec585 | Local code compliance and enforcement posture | official_registry_entry |
| src:usa-nv:nrs-278-586 | NRS 278.586 - Residential fire-sprinkler limits | Nevada Legislature / Legislative Counsel Bureau | statute | https://www.leg.state.nv.us/NRS/NRS-278.html#NRS278Sec586 | Local residential sprinkler amendment limits | official_registry_entry |
| src:usa-nv:nrs-341-100 | NRS 341.100 - Deputy Administrator for Public Works, Compliance and Code Enforcement | Nevada Legislature / Legislative Counsel Bureau | statute | https://www.leg.state.nv.us/NRS/NRS-341.html#NRS341Sec100 | SPWD building official / code enforcement authority | official_registry_entry |
| src:usa-nv:nrs-341-105 | NRS 341.105 - Enforcement authority of Deputy Administrator | Nevada Legislature / Legislative Counsel Bureau | statute | https://www.leg.state.nv.us/NRS/NRS-341.html#NRS341Sec105 | SPWD stop-work and enforcement authority | official_registry_entry |
| src:usa-nv:nrs-341-1455 | NRS 341.1455 - Permits for planning, maintenance or construction of state buildings | Nevada Legislature / Legislative Counsel Bureau | statute | https://www.leg.state.nv.us/NRS/NRS-341.html#NRS341Sec1455 | SPWD permit authority for state property | official_registry_entry |
| src:usa-nv:nrs-341-141 | NRS 341.141 - State-building exception referenced in NAC 341.045 | Nevada Legislature / Legislative Counsel Bureau | statute | https://www.leg.state.nv.us/NRS/NRS-341.html#NRS341Sec141 | State-building scope caveat | official_registry_entry |
| src:usa-nv:nac-341-045 | NAC 341.045 - Codes adopted by State Public Works Board | Nevada Legislature / Legislative Counsel Bureau | regulation | https://www.leg.state.nv.us/NAC/NAC-341.html#NAC341Sec045 | SPWD state-building code editions and substitutions | official_registry_entry |
| src:usa-nv:spwd-permitting | Permitting & Code Enforcement | Nevada State Public Works Division | agency_page | https://publicworks.nv.gov/Services/Permitting___Code_Enforcement/ | SPWD state-owned-land enforcement, permit, plan review, inspection scope | official_agency_page |
| src:usa-nv:spwd-statutes-regulations | Statutes and Regulations | Nevada State Public Works Division | agency_page | https://publicworks.nv.gov/About/Statutes_and_Regulations/ | SPWD statute/regulation locator | official_agency_page |
| src:usa-nv:spwd-adopted-codes-2024 | 2024 Adopted Codes | Nevada State Public Works Division | pdf | https://publicworks.nv.gov/uploadedFiles/publicworksnvgov/content/Services/2024%20Adopted%20Codes.pdf | SPWD 2024 adopted-code list | official_pdf_partial_parse |
| src:usa-nv:reg-r072-24 | Approved Regulation R072-24 | Nevada Register | pdf | https://www.leg.state.nv.us/Register/2024Register/R072-24AP.pdf | NAC 341.045 2024-code amendments | official_pdf_partial_parse |
| src:usa-nv:nrs-701-220 | NRS 701.220 - Energy conservation standards in buildings | Nevada Legislature / Legislative Counsel Bureau | statute | https://www.leg.state.nv.us/NRS/NRS-701.html#NRS701Sec220 | GOE authority, statewide minimum, local stricter standards | official_registry_entry |
| src:usa-nv:nac-701-185 | NAC 701.185 - Energy-code adoption by reference | Nevada Legislature / Legislative Counsel Bureau | regulation | https://www.leg.state.nv.us/NAC/NAC-701.html#NAC701Sec185 | IECC adoption by reference | official_registry_entry |
| src:usa-nv:goe-building-energy-codes | Building Energy Codes | Nevada Governor's Office of Energy | agency_page | https://energy.nv.gov/Programs/Building_Energy_Codes/ | GOE authority and local incorporation/enforcement summary | official_agency_page |
| src:usa-nv:goe-state-adoption-status | State Adoption Status | Nevada Governor's Office of Energy | agency_page | https://energy.nv.gov/Programs/Building_Energy_Codes/State_Adoption_Status/ | 2024 IECC adoption/effective date and statewide energy status | official_agency_page |
| src:usa-nv:nrs-477-030 | NRS 477.030 - Duties of State Fire Marshal; regulations; enforcement | Nevada Legislature / Legislative Counsel Bureau | statute | https://www.leg.state.nv.us/NRS/NRS-477.html#NRS477Sec030 | State Fire Marshal authority and local enforcement coordination | official_registry_entry |
| src:usa-nv:nac-477-281 | NAC 477.281 - Fire-code regulations adopted by reference | Nevada Legislature / Legislative Counsel Bureau | regulation | https://www.leg.state.nv.us/NAC/NAC-477.html#NAC477Sec281 | SFM adopted-code references | official_registry_entry |
| src:usa-nv:nac-477-283 | NAC 477.283 - Fire-code amendments | Nevada Legislature / Legislative Counsel Bureau | regulation | https://www.leg.state.nv.us/NAC/NAC-477.html#NAC477Sec283 | SFM amendments and code substitutions | official_registry_entry |
| src:usa-nv:reg-r205-24 | Approved Regulation R205-24 | Nevada Register | pdf | https://www.leg.state.nv.us/Register/2024Register/R205-24AP.pdf | 2024 SFM fire-code updates and substitutions | official_pdf_partial_parse |
| src:usa-nv:sfm-hot-topics | Hot Topics | Nevada State Fire Marshal Division | agency_page | https://fire.nv.gov/bureaus/FPL/HOT_TOPICS/ | SFM 2024-code enforcement date | official_agency_page |
| src:usa-nv:sfm-fire-protection-engineering | Fire Protection Engineering | Nevada State Fire Marshal Division | agency_page | https://fire.nv.gov/bureaus/FPL/Fire_Protection_Engineering/ | SFM code program context and enforcement notices | official_agency_page |
| src:usa-nv:nrs-455c | NRS Chapter 455C - Elevators and related conveyances | Nevada Legislature / Legislative Counsel Bureau | statute | https://www.leg.state.nv.us/NRS/NRS-455C.html | Elevator/conveyance program authority | official_registry_entry |
| src:usa-nv:nac-455c-500 | NAC 455C.500 - Codes adopted by reference | Nevada Legislature / Legislative Counsel Bureau | regulation | https://www.leg.state.nv.us/NAC/NAC-455C.html#NAC455CSec500 | Elevator/conveyance code adoption by reference | official_registry_entry |
| src:usa-nv:nac-455c-516 | NAC 455C.516 - Inspection and standards references | Nevada Legislature / Legislative Counsel Bureau | regulation | https://www.leg.state.nv.us/NAC/NAC-455C.html#NAC455CSec516 | Elevator/conveyance inspection standards | official_registry_entry |
| src:usa-nv:dir-mcs-information | Mechanical Compliance Section information | Nevada Division of Industrial Relations | agency_page | https://dir.nv.gov/Mechanical_Compliance_Section/Information/ | Elevator/conveyance program and adopted-publications context | official_agency_page |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-nv:nrs-244-3675; src:usa-nv:nrs-268-413; src:usa-nv:nrs-278-580; src:usa-nv:nrs-278-581; src:usa-nv:nrs-278-583; src:usa-nv:nrs-278-585; src:usa-nv:nrs-278-586; src:usa-nv:nrs-341-100; src:usa-nv:nrs-341-105; src:usa-nv:nrs-341-1455; src:usa-nv:nrs-701-220; src:usa-nv:nrs-477-030 | official_html_extraction_caveat | Official Nevada Legislature NRS/NAC locators are used as the authoritative source registry. Text extraction was cross-checked against agency pages and public legal mirrors where direct official HTML extraction was degraded. | verify_direct_official_text_before_verified_status |
| src:usa-nv:nac-341-045; src:usa-nv:nac-477-281; src:usa-nv:nac-477-283; src:usa-nv:nac-455c-500; src:usa-nv:nac-455c-516 | codification_timing | Codified NAC pages should be rechecked after regulation consolidation, especially for 2024 SFM and SPWD changes. | revalidate_codified_nac |
| src:usa-nv:reg-r072-24; src:usa-nv:reg-r205-24; src:usa-nv:spwd-adopted-codes-2024 | pdf_partial_parse | PDFs were used for code-edition/amendment extraction, but full page-level manual validation was not completed. | use_for_partially_verified_only |
| src:usa-nv:goe-state-adoption-status | agency_page_update_risk | GOE page supplies the 2024 IECC effective date and was treated as official; local implementation and local stricter provisions still require local ordinance review. | monitor_quarterly |
| src:usa-nv:sfm-hot-topics | agency_notice | SFM enforcement-date notice is official but should be paired with final NAC 477 codification review for verified production use. | monitor_monthly |

### 8.3 Supplemental Sources

| Supplemental Source ID | Title | Publisher | Purpose | Treatment |
| --- | --- | --- | --- | --- |
| supp:usa-nv:justia-nrs | Nevada Revised Statutes mirrored sections | Justia | Cross-check NRS text when official HTML extraction was degraded. | not primary; use only to support extraction QA |
| supp:usa-nv:public-law-nrs | Nevada Revised Statutes mirrored sections | Public.Law | Cross-check NRS text when official HTML extraction was degraded. | not primary; use only to support extraction QA |
| supp:usa-nv:cornell-nac | Nevada Administrative Code mirrored sections | Legal Information Institute / Cornell | Cross-check NAC text when official HTML extraction was degraded. | not primary; use only to support extraction QA |

### 8.4 Source Extraction Metadata

| Extraction ID | Source IDs | Extracted Fields | Extraction Method | Extracted On | Completeness | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| extract:usa-nv:001 | src:usa-nv:spwd-permitting; src:usa-nv:nac-341-045; src:usa-nv:reg-r072-24 | SPWD authority and state-building code set | official agency page plus regulation/PDF review | 2026-06-26 | partial | Effective/operative date for R072-24 remains unresolved. |
| extract:usa-nv:002 | src:usa-nv:goe-state-adoption-status; src:usa-nv:nrs-701-220; src:usa-nv:nac-701-185 | Energy authority, 2024 IECC date, local stricter-standard rule | official agency page plus statute/regulation locators | 2026-06-26 | moderate | Local implementation not normalized. |
| extract:usa-nv:003 | src:usa-nv:nrs-244-3675; src:usa-nv:nrs-268-413; src:usa-nv:nrs-278-580; src:usa-nv:nrs-278-585 | Local code adoption/enforcement authority | official statute locators plus legal mirror cross-check | 2026-06-26 | moderate | Local code editions and local transition rules still open. |
| extract:usa-nv:004 | src:usa-nv:nrs-477-030; src:usa-nv:sfm-hot-topics; src:usa-nv:reg-r205-24 | State Fire Marshal authority and 2024-code package | official agency notice and Nevada Register PDF | 2026-06-26 | partial | Final consolidated NAC 477 review remains open. |
| extract:usa-nv:005 | src:usa-nv:nrs-455c; src:usa-nv:nac-455c-500; src:usa-nv:nac-455c-516; src:usa-nv:dir-mcs-information | Elevator/conveyance program | official statute/regulation locators plus DIR page | 2026-06-26 | limited | Adopted edition list not fully parsed. |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 1.00 | none | Status reflects that core authority and code-adoption fields now have source IDs, while local normalization remains open. |
| report | risk.overall_confidence | 0.58 | verified | 1.00 | none | Confidence remains below production level due to local-code, SFM codification, and elevator edition gaps. |
| ahj:usa-nv:spwd | authority.name | Nevada State Public Works Division / State Public Works Board | partially_verified | 0.78 | src:usa-nv:spwd-permitting; src:usa-nv:nac-341-045 | Strong for state-property scope; not a statewide private-building authority. |
| ahj:usa-nv:local-building-governments | role | Local building-code adoption and enforcement | partially_verified | 0.72 | src:usa-nv:nrs-244-3675; src:usa-nv:nrs-268-413; src:usa-nv:nrs-278-580; src:usa-nv:nrs-278-585 | Local edition-by-jurisdiction extraction remains open. |
| ahj:usa-nv:goe | role | Statewide energy-conservation code adoption | partially_verified | 0.82 | src:usa-nv:nrs-701-220; src:usa-nv:goe-state-adoption-status | 2024 IECC effective date confirmed by GOE page. |
| ahj:usa-nv:sfm | role | State Fire Marshal fire/life-safety code authority | partially_verified | 0.70 | src:usa-nv:nrs-477-030; src:usa-nv:sfm-hot-topics; src:usa-nv:reg-r205-24 | SFM scope and 2024-code enforcement date confirmed; codified NAC cross-check open. |
| adoption:usa-nv:spwd-2024-building-codes | base_codes | 2024 state-building code package | partially_verified | 0.70 | src:usa-nv:nac-341-045; src:usa-nv:reg-r072-24; src:usa-nv:spwd-adopted-codes-2024 | State-building code editions captured, date sequence still open. |
| adoption:usa-nv:goe-2024-iecc | effective_date | 2024-08-18 | partially_verified | 0.82 | src:usa-nv:goe-state-adoption-status; src:usa-nv:nac-701-185 | Effective date and code edition confirmed. |
| adoption:usa-nv:sfm-2024-fire-codes | effective_date | 2026-01-01 | partially_verified | 0.70 | src:usa-nv:sfm-hot-topics; src:usa-nv:reg-r205-24 | Announced enforcement date confirmed; adoption-date sequence not entered. |
| local-amendment-rule:usa-nv | model | local_amendments_allowed_with_state_floors_and_topic_specific_limits | partially_verified | 0.45 | src:usa-nv:nrs-701-220; src:usa-nv:nrs-278-583; src:usa-nv:nrs-278-586 | No statewide local-amendment registry confirmed. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Every src:usa-nv identifier cited in sections 1-7 and 9-11 is registered in section 8. |
| All authority IDs resolve | pass | Authority IDs are introduced in section 2 or in local enforcement records. |
| All current code families have adoption rows | pass | Rows are present for all template code families; unresolved local scopes are explicit. |
| Building and operational fire code are separated | pass | Building/local/SPWD and State Fire Marshal fire-code tracks are separated. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Dates are maintained in separate columns; unresolved dates remain null. |
| Effective dates are valid ISO dates | pass | Entered effective dates use YYYY-MM-DD format. |
| No impossible date sequences | pass | Ambiguous adoption dates were left null where they could conflict with effective dates. |
| Transition rules have explicit trigger conditions | pass | Confirmed triggers are stated; local transition rules are explicitly unresolved. |
| Permit-date logic is captured where applicable | fail | Statewide local permit-date/application-date transition logic is not normalized. |
| Local enforcement model classified | pass | Model is classified as local adoption/enforcement with state specialized overlays. |
| Local amendment rule classified | pass | Model is classified, with unresolved registry and topic-specific limits documented. |
| AHJ confirmation metadata present | fail | AHJ contacts and local boundary data were not populated. |
| Official-source caveats captured | pass | Section 8.2 documents official HTML, PDF, and codification caveats. |
| No leftover template markers | pass | Template markers targeted by the validation regex were removed. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-nv:001 | high | local code editions | Local building, residential, fire, electrical, plumbing, mechanical, and existing-building editions are not normalized for each Nevada city/county. | Build jurisdiction-level local code matrix, starting with Clark County, Washoe County, Las Vegas, Henderson, North Las Vegas, Reno, and Sparks. | null | null | open |
| issue:usa-nv:002 | high | local amendment registry | No statewide registry or filing workflow for local building-code amendments was confirmed. | Search state agency records and local ordinance portals; determine whether any filing requirement exists beyond topic-specific rules. | null | null | open |
| issue:usa-nv:003 | high | SFM codification | R205-24 and SFM 2024-code enforcement were reviewed, but final consolidated NAC 477 adoption/effective-date sequence was not fully validated. | Recheck NAC 477.281 and NAC 477.283 after codification and compare with R205-24. | null | null | open |
| issue:usa-nv:004 | medium | SPWD effective date | R072-24 and NAC 341.045 support the SPWD 2024 state-building code package, but the effective/operative date sequence was not fully extracted. | Review R072-24 rulemaking history and codified NAC 341.045 effective notes. | null | null | open |
| issue:usa-nv:005 | medium | elevator editions | Elevator/conveyance standards are confirmed as an adopted-standards program, but the exact adopted edition list under NAC 455C.500 was not fully parsed. | Extract all standards and editions from NAC 455C.500 and DIR adopted-publications list. | null | null | open |
| issue:usa-nv:006 | medium | fuel gas and LPG split | Fire-code amendments reference NFPA 54 and the LPG Board version under NAC 590.610, but the gas/LPG authority split was not fully modeled. | Review NAC 590.610 and related LPG Board authority for integration with building/fire reports. | null | null | open |
| issue:usa-nv:007 | medium | school construction exception | NRS 278.585 references a school-construction exception under NRS 393.110, but school facility code authority was not parsed. | Review NRS 393.110 and education-facility construction authority. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-nv:spwd-permitting | src:usa-nv:spwd-permitting | html_diff | quarterly | changes to SPWD permitting scope or state-building code link | 2026-06-26 | active |
| watch:usa-nv:nac-341 | src:usa-nv:nac-341-045 | html_diff | monthly | codified change to state-building code editions or substitutions | 2026-06-26 | active |
| watch:usa-nv:goe-energy | src:usa-nv:goe-state-adoption-status | html_diff | quarterly | new IECC adoption date, local implementation notice, or code revision | 2026-06-26 | active |
| watch:usa-nv:nac-701 | src:usa-nv:nac-701-185 | html_diff | quarterly | codified IECC reference change | 2026-06-26 | active |
| watch:usa-nv:sfm-hot-topics | src:usa-nv:sfm-hot-topics | html_diff | monthly | change to SFM adopted/enforced code package or transition guidance | 2026-06-26 | active |
| watch:usa-nv:nac-477 | src:usa-nv:nac-477-281; src:usa-nv:nac-477-283 | html_diff | monthly | codified change to SFM code editions or amendments | 2026-06-26 | active |
| watch:usa-nv:nac-455c | src:usa-nv:nac-455c-500 | html_diff | quarterly | elevator/conveyance adopted-standard edition change | 2026-06-26 | pending |
| watch:usa-nv:local-major-jurisdictions | none | local_ordinance_review | quarterly | major local jurisdiction code edition/amendment changes | 2026-06-26 | pending |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-26 | Populated Nevada report from baseline stub | report:usa-nv | src:usa-nv:spwd-permitting; src:usa-nv:nac-341-045; src:usa-nv:goe-state-adoption-status; src:usa-nv:nrs-701-220; src:usa-nv:nrs-477-030; src:usa-nv:sfm-hot-topics; src:usa-nv:reg-r205-24 | ChatGPT | Replaced generic unresolved text with source-backed hybrid authority model and explicit open issues. |
