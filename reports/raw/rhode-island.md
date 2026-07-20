---
state:
  state_id: "US-RI"
  name: "Rhode Island"
  abbreviation: "RI"
report:
  report_id: "state-report:usa-ri"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-26"
  last_verified: null
  reviewed_by: null
risk:
  overall_confidence: 0.72 # 0.00 - 1.00
  risk_flags:
    - "state_amendments_identified_but_not_parsed_clause_by_clause"
    - "local_ahj_directory_not_populated"
    - "fire_local_contact_routing_not_populated"
    - "transition_rule_interpreted_from_statute_not_agency_guidance"
  open_questions_count: 6

---

# State Building Code Authority Report: Rhode Island

## 1. Executive Summary

- **Authority model:** Rhode Island has a statewide construction-code adoption model. The Building Code Standards Committee adopts, promulgates, and administers the State Building Code. The State Building Code Commissioner standardizes interpretation and provides state-level enforcement support, while local building officials enforce the State Building Code in municipalities. Distinct fire, rehabilitation, and elevator authorities operate under separate legal bases.

- **Statewide code status:** The core Rhode Island Building Code rules in Title 510 are active. The current construction-code package is primarily based on the 2021 ICC code family, the 2024 International Energy Conservation Code, and the 2023 National Electrical Code, with Rhode Island amendments. The Building Code Commission states that the current 2021 ICC-based package took effect on 2025-12-01.

- **Local enforcement model:** Local authorities appoint local building officials, and local building officials enforce the State Building Code. The State Building Code Commissioner has statewide administrative, interpretation, permit-system, and backstop enforcement responsibilities. State projects are administered by the State Building Commissioner. Fire-code enforcement is separately vested in the State Fire Marshal for strict enforcement, with Fire Safety Code Board authority over administration, waivers, variances, and amendments.

- **Local amendment posture:** Local building codes and ordinances are preempted for future enactment. Amendments proceed through the statewide Building Code Standards Committee process; amendments adopted by the committee are binding on all cities and towns.

- **Known transition periods or pending changes:** R.I. Gen. Laws § 23-27.3-114.3.1 gives a three-month window after adoption of an amendment for an owner or agent to notify the building official that project drawings and specifications are based on the prior code; after that notice, the owner may obtain a permit based on the prior code provisions when the drawings and specifications are completed. No future statewide code package beyond the active rules was verified in this pass.

- **Production readiness:** partially_ready_for_statewide_code_fields; not_ready_for_full_address_to_AHJ_resolution

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: The Building Code Standards Committee is the statewide authority to adopt,
    promulgate, and administer the State Building Code.
  confidence: 0.92
  source_ids:
  - src:usa-ri:statute-bcsc-adoption
- topic: Primary building code edition
  finding: RISBC-1 incorporates the 2021 International Building Code with Rhode Island
    amendments and is active effective 2025-12-01.
  confidence: 0.92
  source_ids:
  - src:usa-ri:ricr-510-00-00-1
  - src:usa-ri:bcc-laws-rules
- topic: Residential code edition
  finding: RISBC-2 incorporates the 2021 International Residential Code with Rhode
    Island amendments and is active effective 2025-12-01.
  confidence: 0.91
  source_ids:
  - src:usa-ri:ricr-510-00-00-2
  - src:usa-ri:bcc-laws-rules
- topic: Electrical code edition
  finding: RISBC-5 incorporates the 2023 National Electrical Code with Rhode Island
    amendments and is active effective 2025-12-01.
  confidence: 0.9
  source_ids:
  - src:usa-ri:ricr-510-00-00-5
  - src:usa-ri:statute-bcsc-adoption
- topic: Energy code edition
  finding: RISBC-8 incorporates the 2024 International Energy Conservation Code with
    Rhode Island amendments and is active effective 2025-12-01.
  confidence: 0.9
  source_ids:
  - src:usa-ri:ricr-510-00-00-8
- topic: Fire code authority and editions
  finding: The Fire Safety Code Board administers the Rhode Island Fire Code and Life
    Safety Code; the State Fire Marshal is the strict-enforcement AHJ. The current
    fire package includes NFPA 1 2021, NFPA 101 2021, and NFPA 72 2022, effective
    2026-03-01.
  confidence: 0.9
  source_ids:
  - src:usa-ri:ricr-450-00-00-7
  - src:usa-ri:ricr-450-00-00-8
  - src:usa-ri:ricr-450-00-00-10
- topic: Local enforcement
  finding: Local authorities appoint local building officials; local building officials
    enforce the State Building Code, with state backstop where no local official or
    inspector is available.
  confidence: 0.86
  source_ids:
  - src:usa-ri:statute-local-building-official-appointment
  - src:usa-ri:statute-local-building-official-duties
  - src:usa-ri:statute-state-commissioner-duties
- topic: Local amendments
  finding: Cities and towns are prohibited from enacting local building codes or ordinances;
    statewide amendments adopted by the Building Code Standards Committee bind all
    cities and towns.
  confidence: 0.88
  source_ids:
  - src:usa-ri:statute-local-code-preemption
  - src:usa-ri:statute-amendment-petition
- topic: Effective / transition rule
  finding: A three-month post-amendment notice path allows certain projects to proceed
    under prior code provisions when drawings/specifications were based on the prior
    code.
  confidence: 0.8
  source_ids:
  - src:usa-ri:statute-transition-drawings-specs
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-ri:bcsc |
| Authority name | Building Code Standards Committee |
| Authority type | statewide_code_adopting_authority |
| Legal basis | R.I. Gen. Laws § 23-27.3-100.1.5; Title 510 RICR authority clauses |
| Role | Adopt, promulgate, and administer the State Building Code; adopt ICC, NEC, and related statewide code rules with Rhode Island amendments. |
| Enforcement model | Statewide code adoption with local building official enforcement and State Building Code Commissioner interpretation/backstop functions. |
| Source IDs | src:usa-ri:statute-bcsc-adoption; src:usa-ri:ricr-510-00-00-1 |
| Verification status | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-ri:bcsc | Building Code Standards Committee | Adopts and administers RISBC-1 and related State Building Code provisions. | R.I. Gen. Laws § 23-27.3-100.1.5; 510-RICR-00-00-1 | src:usa-ri:statute-bcsc-adoption; src:usa-ri:ricr-510-00-00-1 | partially_verified |
| Residential | ahj:usa-ri:bcsc | Building Code Standards Committee | Adopts and administers RISBC-2 for detached one- and two-family dwellings, townhouses, and listed accessory structures. | R.I. Gen. Laws § 23-27.3-100.1.5; 510-RICR-00-00-2 | src:usa-ri:statute-bcsc-adoption; src:usa-ri:ricr-510-00-00-2 | partially_verified |
| Existing Building / Rehabilitation | ahj:usa-ri:bcsc; ahj:usa-ri:rehab-joint-committee | Building Code Standards Committee; Joint Committee on the Rehabilitation Building and Fire Code for Existing Buildings and Structures | BCSC adopts the Existing Building Code; Joint Committee promulgates the State Rehabilitation Building and Fire Code for qualifying existing structures and resolves rehabilitation-code appeals/variances. | 510-RICR-00-00-7; 510-RICR-00-00-20; R.I. Gen. Laws §§ 23-29.1-2 and 23-29.1-3 | src:usa-ri:ricr-510-00-00-7; src:usa-ri:ricr-510-00-00-20 | partially_verified |
| Mechanical | ahj:usa-ri:bcsc | Building Code Standards Committee | Adopts and administers RISBC-4. | 510-RICR-00-00-4 | src:usa-ri:ricr-510-00-00-4 | partially_verified |
| Plumbing | ahj:usa-ri:bcsc | Building Code Standards Committee | Adopts and administers RISBC-3. | 510-RICR-00-00-3 | src:usa-ri:ricr-510-00-00-3 | partially_verified |
| Fuel Gas | ahj:usa-ri:bcsc | Building Code Standards Committee | Adopts and administers RISBC-19. | 510-RICR-00-00-19 | src:usa-ri:ricr-510-00-00-19 | partially_verified |
| Electrical | ahj:usa-ri:bcsc | Building Code Standards Committee | Adopts and administers RISBC-5; statute separately directs use/adoption of the latest NEC framework. | R.I. Gen. Laws § 23-27.3-100.1.5; 510-RICR-00-00-5 | src:usa-ri:statute-bcsc-adoption; src:usa-ri:ricr-510-00-00-5 | partially_verified |
| Energy | ahj:usa-ri:bcsc | Building Code Standards Committee | Adopts and administers RISBC-8. | 510-RICR-00-00-8 | src:usa-ri:ricr-510-00-00-8 | partially_verified |
| Fire - construction references | ahj:usa-ri:bcsc; ahj:usa-ri:fscbar | Building Code Standards Committee; Fire Safety Code Board of Appeal and Review | Building-code rules reference the Rhode Island Fire Safety Code and delete or defer some fire-system topics to Title 450. | 510-RICR-00-00-1; 450-RICR-00-00-7; 450-RICR-00-00-8 | src:usa-ri:ricr-510-00-00-1; src:usa-ri:ricr-450-00-00-7; src:usa-ri:ricr-450-00-00-8 | partially_verified |
| Fire - operational / prevention code | ahj:usa-ri:fscbar; ahj:usa-ri:state-fire-marshal | Fire Safety Code Board of Appeal and Review; State Fire Marshal | Fire Safety Code Board administers, varies, waives, and amends fire rules; State Fire Marshal is the strict-enforcement AHJ. | R.I. Gen. Laws § 23-28.3-3; 450-RICR-00-00-7; 450-RICR-00-00-8 | src:usa-ri:ricr-450-00-00-7; src:usa-ri:ricr-450-00-00-8 | partially_verified |
| Accessibility | ahj:usa-ri:bcsc | Building Code Standards Committee | Accessibility is primarily within the State Building Code; RISBC-17 separately adopts UFAS for pre-1977 public-meeting buildings. | 510-RICR-00-00-17; R.I. Gen. Laws § 23-27.3-100.1.5 | src:usa-ri:ricr-510-00-00-17; src:usa-ri:statute-bcsc-adoption | partially_verified |
| Elevator / Conveyance | ahj:usa-ri:dlt-elevator | Rhode Island Department of Labor and Training, Division of Occupational Safety / Chief Elevator Inspector | Administers elevator and conveyance safety standards, permits, inspections, certificates, and licensing. | R.I. Gen. Laws § 23-33-2; 260-RICR-30-10-1 | src:usa-ri:ricr-260-30-10-1 | partially_verified |

### 2.3 Authority Hierarchy Notes

Rhode Island uses a hybrid implementation model: statewide code adoption is centralized, but day-to-day building-code enforcement is local unless a state function is specifically reserved or triggered. Local enforcement does not imply local amendment authority. Cities and towns appoint local building officials and may administer permits and inspections, but local building-code enactment is preempted by statute.

The State Building Code Commissioner has important statewide coordination responsibilities, including standardizing interpretation and acting when a local building official, alternate official, or inspectors are unavailable. Official interpretations issued by the State Building Commissioner are distributed to local officials and may be appealed to the Building Code Standards Committee.

Fire safety is separate from the construction-code adoption hierarchy. The Fire Safety Code Board of Appeal and Review administers fire rules and variances, while the State Fire Marshal is the strict-enforcement AHJ for the Rhode Island Fire Code and Life Safety Code. The rehabilitation code blends building and fire elements and uses the Joint Committee for Rehabilitation Building and Fire Code matters.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-ri:001 | ahj:usa-ri:bcsc | adopts_and_administers | statewide State Building Code / Title 510 construction-code rules | src:usa-ri:statute-bcsc-adoption; src:usa-ri:ricr-510-00-00-1 | partially_verified |
| edge:usa-ri:002 | local_authorities | appoint | local_building_officials | src:usa-ri:statute-local-building-official-appointment | partially_verified |
| edge:usa-ri:003 | local_building_officials | enforce | State Building Code within local jurisdiction | src:usa-ri:statute-local-building-official-duties | partially_verified |
| edge:usa-ri:004 | ahj:usa-ri:state-building-code-commissioner | standardizes_interpretation_and_backstops | local enforcement / locations with no available local official or inspectors | src:usa-ri:statute-state-commissioner-duties; src:usa-ri:ricr-510-00-00-10 | partially_verified |
| edge:usa-ri:005 | cities_and_towns | prohibited_from_enacting | local building codes or ordinances after state-code adoption | src:usa-ri:statute-local-code-preemption | partially_verified |
| edge:usa-ri:006 | ahj:usa-ri:bcsc | statewide_amendments_bind | all cities and towns | src:usa-ri:statute-amendment-petition | partially_verified |
| edge:usa-ri:007 | ahj:usa-ri:fscbar | administers_and_varies | Rhode Island Fire Code, Life Safety Code, and Fire Alarm Code | src:usa-ri:ricr-450-00-00-7; src:usa-ri:ricr-450-00-00-8; src:usa-ri:ricr-450-00-00-10 | partially_verified |
| edge:usa-ri:008 | ahj:usa-ri:state-fire-marshal | strictly_enforces | Rhode Island Fire Code and Life Safety Code | src:usa-ri:ricr-450-00-00-7; src:usa-ri:ricr-450-00-00-8 | partially_verified |
| edge:usa-ri:009 | ahj:usa-ri:dlt-elevator | administers | elevator and conveyance safety rules, permits, inspections, and certificates | src:usa-ri:ricr-260-30-10-1 | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | RISBC-1 Rhode Island Building Code | International Building Code | 2021 | active | null | 2025-12-01 | null | null | Three-month prior-code drawings/specifications notice path after adoption of amendments; mandatory date not separately stated in the rule. | src:usa-ri:ricr-510-00-00-1; src:usa-ri:bcc-laws-rules; src:usa-ri:statute-transition-drawings-specs |
| Residential | RISBC-2 Rhode Island One- and Two-Family Dwelling Code | International Residential Code | 2021 | active | null | 2025-12-01 | null | null | Same statutory prior-code drawings/specifications notice path unless a more specific rule applies. | src:usa-ri:ricr-510-00-00-2; src:usa-ri:bcc-laws-rules; src:usa-ri:statute-transition-drawings-specs |
| Existing Building / Rehabilitation | RISBC-7 Rhode Island Existing Building Code; RISRC-1 State Rehabilitation Building and Fire Code for Existing Structures | International Existing Building Code; Rhode Island rehabilitation code using current IEBC and Life Safety Code rehabilitation provisions | 2021 IEBC; current linked LSC Chapter 43 at time of permit | active | null | 2025-12-01 | null | null | Part 20 expressly keys referenced codes to the editions in effect at permit issuance unless later amendments are made retroactive. | src:usa-ri:ricr-510-00-00-7; src:usa-ri:ricr-510-00-00-20 |
| Mechanical | RISBC-4 Rhode Island Mechanical Code | International Mechanical Code | 2021 | active | null | 2025-12-01 | null | null | Same statutory prior-code drawings/specifications notice path unless a more specific rule applies. | src:usa-ri:ricr-510-00-00-4; src:usa-ri:bcc-laws-rules; src:usa-ri:statute-transition-drawings-specs |
| Plumbing | RISBC-3 Rhode Island Plumbing Code | International Plumbing Code | 2021 | active | null | 2025-12-01 | null | null | Same statutory prior-code drawings/specifications notice path unless a more specific rule applies. | src:usa-ri:ricr-510-00-00-3; src:usa-ri:bcc-laws-rules; src:usa-ri:statute-transition-drawings-specs |
| Fuel Gas | RISBC-19 Rhode Island Fuel Gas Code | International Fuel Gas Code | 2021 | active | null | 2025-12-01 | null | null | Same statutory prior-code drawings/specifications notice path unless a more specific rule applies. | src:usa-ri:ricr-510-00-00-19; src:usa-ri:bcc-laws-rules; src:usa-ri:statute-transition-drawings-specs |
| Electrical | RISBC-5 Rhode Island Electrical Code | NFPA 70, National Electrical Code | 2023 | active | null | 2025-12-01 | null | null | Current rule effective date verified; statute also contains a latest-NEC adoption framework that should be monitored when the next NEC edition is issued. | src:usa-ri:ricr-510-00-00-5; src:usa-ri:statute-bcsc-adoption |
| Energy | RISBC-8 Rhode Island Energy Conservation Code | International Energy Conservation Code | 2024 | active | null | 2025-12-01 | null | null | Current rule effective date verified; earlier energy-code transition history not parsed. | src:usa-ri:ricr-510-00-00-8 |
| Fire - construction references | Rhode Island Fire Safety Code references within RISBC-1 and other Title 510 parts | Rhode Island Fire Code / Rhode Island Life Safety Code / Fire Alarm Code | NFPA 1 2021; NFPA 101 2021; NFPA 72 2022 | active | null | 2026-03-01 | null | null | Fire-system and life-safety subjects are separated from, and cross-referenced by, construction-code rules. | src:usa-ri:ricr-510-00-00-1; src:usa-ri:ricr-450-00-00-7; src:usa-ri:ricr-450-00-00-8; src:usa-ri:ricr-450-00-00-10 |
| Fire - operational / prevention code | Rhode Island Fire Code; Rhode Island Life Safety Code; Rhode Island Fire Alarm Code | NFPA 1; NFPA 101; NFPA 72 | 2021; 2021; 2022 | active | null | 2026-03-01 | null | null | Fire Safety Code Board administration; State Fire Marshal strict enforcement. | src:usa-ri:ricr-450-00-00-7; src:usa-ri:ricr-450-00-00-8; src:usa-ri:ricr-450-00-00-10 |
| Accessibility | State Building Code accessibility provisions; RISBC-17 Public Meetings Accessibility Standard | IBC Chapter 11 / UFAS public-meeting standard | 2021 IBC / UFAS | active | null | 2025-12-01 for IBC accessibility provisions; 2022-01-04 for RISBC-17 | null | null | RISBC-17 covers public-meeting buildings constructed or occupied before 1977; new construction accessibility is primarily in the State Building Code. | src:usa-ri:ricr-510-00-00-1; src:usa-ri:ricr-510-00-00-17 |
| Elevator / Conveyance | Elevator Safety Code | ASME A17.1, A17.3, A10.4, A10.5, A18.1, B20.1, A17.8 | 2016; 2015; 2016; 2013; 2017; 2015; 2016 | active | null | 2022-01-04 | null | null | Elevator code is administered outside Title 510 by the Department of Labor and Training; building-code interface should be checked project-by-project. | src:usa-ri:ricr-260-30-10-1 |

### 3.2 Adoption Records

```yaml
adoptions:
  - adoption_id: adoption:usa-ri:building:2021-ibc
    code_family: Building
    state_code_name: RISBC-1 Rhode Island Building Code
    base_model_code: International Building Code
    edition: "2021"
    adopting_authority_id: ahj:usa-ri:bcsc
    status: active
    adoption_date: null
    effective_date: "2025-12-01"
    operative_date: null
    mandatory_date: null
    source_ids:
      - src:usa-ri:ricr-510-00-00-1
      - src:usa-ri:bcc-laws-rules
    confidence: 0.92

  - adoption_id: adoption:usa-ri:residential:2021-irc
    code_family: Residential
    state_code_name: RISBC-2 Rhode Island One- and Two-Family Dwelling Code
    base_model_code: International Residential Code
    edition: "2021"
    adopting_authority_id: ahj:usa-ri:bcsc
    status: active
    adoption_date: null
    effective_date: "2025-12-01"
    operative_date: null
    mandatory_date: null
    source_ids:
      - src:usa-ri:ricr-510-00-00-2
      - src:usa-ri:bcc-laws-rules
    confidence: 0.91

  - adoption_id: adoption:usa-ri:existing:2021-iebc
    code_family: Existing Building / Rehabilitation
    state_code_name: RISBC-7 Rhode Island Existing Building Code
    base_model_code: International Existing Building Code
    edition: "2021"
    adopting_authority_id: ahj:usa-ri:bcsc
    status: active
    adoption_date: null
    effective_date: "2025-12-01"
    operative_date: null
    mandatory_date: null
    source_ids:
      - src:usa-ri:ricr-510-00-00-7
    confidence: 0.90

  - adoption_id: adoption:usa-ri:rehab:risrc-1
    code_family: Existing Building / Rehabilitation
    state_code_name: RISRC-1 State Rehabilitation Building and Fire Code for Existing Structures
    base_model_code: Rhode Island rehabilitation code linked to current IEBC and LSC rehabilitation provisions
    edition: current referenced editions at permit issuance
    adopting_authority_id: ahj:usa-ri:rehab-joint-committee
    status: active
    adoption_date: null
    effective_date: "2025-12-01"
    operative_date: null
    mandatory_date: null
    source_ids:
      - src:usa-ri:ricr-510-00-00-20
    confidence: 0.86

  - adoption_id: adoption:usa-ri:plumbing:2021-ipc
    code_family: Plumbing
    state_code_name: RISBC-3 Rhode Island Plumbing Code
    base_model_code: International Plumbing Code
    edition: "2021"
    adopting_authority_id: ahj:usa-ri:bcsc
    status: active
    adoption_date: null
    effective_date: "2025-12-01"
    operative_date: null
    mandatory_date: null
    source_ids:
      - src:usa-ri:ricr-510-00-00-3
      - src:usa-ri:bcc-laws-rules
    confidence: 0.90

  - adoption_id: adoption:usa-ri:mechanical:2021-imc
    code_family: Mechanical
    state_code_name: RISBC-4 Rhode Island Mechanical Code
    base_model_code: International Mechanical Code
    edition: "2021"
    adopting_authority_id: ahj:usa-ri:bcsc
    status: active
    adoption_date: null
    effective_date: "2025-12-01"
    operative_date: null
    mandatory_date: null
    source_ids:
      - src:usa-ri:ricr-510-00-00-4
      - src:usa-ri:bcc-laws-rules
    confidence: 0.90

  - adoption_id: adoption:usa-ri:electrical:2023-nec
    code_family: Electrical
    state_code_name: RISBC-5 Rhode Island Electrical Code
    base_model_code: NFPA 70 National Electrical Code
    edition: "2023"
    adopting_authority_id: ahj:usa-ri:bcsc
    status: active
    adoption_date: null
    effective_date: "2025-12-01"
    operative_date: null
    mandatory_date: null
    source_ids:
      - src:usa-ri:ricr-510-00-00-5
      - src:usa-ri:statute-bcsc-adoption
    confidence: 0.90

  - adoption_id: adoption:usa-ri:energy:2024-iecc
    code_family: Energy
    state_code_name: RISBC-8 Rhode Island Energy Conservation Code
    base_model_code: International Energy Conservation Code
    edition: "2024"
    adopting_authority_id: ahj:usa-ri:bcsc
    status: active
    adoption_date: null
    effective_date: "2025-12-01"
    operative_date: null
    mandatory_date: null
    source_ids:
      - src:usa-ri:ricr-510-00-00-8
    confidence: 0.90

  - adoption_id: adoption:usa-ri:fuel-gas:2021-ifgc
    code_family: Fuel Gas
    state_code_name: RISBC-19 Rhode Island Fuel Gas Code
    base_model_code: International Fuel Gas Code
    edition: "2021"
    adopting_authority_id: ahj:usa-ri:bcsc
    status: active
    adoption_date: null
    effective_date: "2025-12-01"
    operative_date: null
    mandatory_date: null
    source_ids:
      - src:usa-ri:ricr-510-00-00-19
      - src:usa-ri:bcc-laws-rules
    confidence: 0.90

  - adoption_id: adoption:usa-ri:fire:2021-nfpa-1
    code_family: Fire - operational / prevention code
    state_code_name: Rhode Island Fire Code
    base_model_code: NFPA 1 Fire Code
    edition: "2021"
    adopting_authority_id: ahj:usa-ri:fscbar
    status: active
    adoption_date: null
    effective_date: "2026-03-01"
    operative_date: null
    mandatory_date: null
    source_ids:
      - src:usa-ri:ricr-450-00-00-7
    confidence: 0.90

  - adoption_id: adoption:usa-ri:life-safety:2021-nfpa-101
    code_family: Fire - operational / prevention code
    state_code_name: Rhode Island Life Safety Code
    base_model_code: NFPA 101 Life Safety Code
    edition: "2021"
    adopting_authority_id: ahj:usa-ri:fscbar
    status: active
    adoption_date: null
    effective_date: "2026-03-01"
    operative_date: null
    mandatory_date: null
    source_ids:
      - src:usa-ri:ricr-450-00-00-8
    confidence: 0.90

  - adoption_id: adoption:usa-ri:fire-alarm:2022-nfpa-72
    code_family: Fire - operational / prevention code
    state_code_name: Rhode Island Fire Alarm Code
    base_model_code: NFPA 72 National Fire Alarm and Signaling Code
    edition: "2022"
    adopting_authority_id: ahj:usa-ri:fscbar
    status: active
    adoption_date: null
    effective_date: "2026-03-01"
    operative_date: null
    mandatory_date: null
    source_ids:
      - src:usa-ri:ricr-450-00-00-10
    confidence: 0.90

  - adoption_id: adoption:usa-ri:elevator:asme-package
    code_family: Elevator / Conveyance
    state_code_name: Elevator Safety Code
    base_model_code: ASME elevator and conveyance standards
    edition: "A17.1-2016; A17.3-2015; A10.4-2016; A10.5-2013; A18.1-2017; B20.1-2015; A17.8-2016"
    adopting_authority_id: ahj:usa-ri:dlt-elevator
    status: active
    adoption_date: null
    effective_date: "2022-01-04"
    operative_date: null
    mandatory_date: null
    source_ids:
      - src:usa-ri:ricr-260-30-10-1
    confidence: 0.82
```

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

The current Title 510 construction-code rules reviewed in this pass show an active effective date of 2025-12-01 for the 2021 ICC-based package, the 2024 IECC rule, and the 2023 NEC rule. The current Title 450 fire-code package reviewed in this pass shows an active effective date of 2026-03-01 for NFPA 1 2021, NFPA 101 2021, and NFPA 72 2022.

Rhode Island also has a statutory prior-code transition mechanism. Within three months after adoption of an amendment to the State Building Code, an owner or agent may notify the building official that project drawings and specifications are based on the prior code. If that notice is made, the owner may obtain a permit under the prior code provisions when drawings and specifications are completed. This report records the mechanism but does not convert it into a universal mandatory date because the statute depends on project-specific notice and the rule pages did not separately state a mandatory date.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-ri:construction-code-effective-2025-12-01 | Title 510 construction-code package reviewed here | rule_effective_date | 2025-12-01 | Active RICR filings for current Building, Residential, Plumbing, Mechanical, Electrical, Existing Building, Energy, Fuel Gas, and Rehabilitation rules. | Not determined by effective-date field alone. | src:usa-ri:ricr-510-00-00-1; src:usa-ri:ricr-510-00-00-2; src:usa-ri:ricr-510-00-00-3; src:usa-ri:ricr-510-00-00-4; src:usa-ri:ricr-510-00-00-5; src:usa-ri:ricr-510-00-00-7; src:usa-ri:ricr-510-00-00-8; src:usa-ri:ricr-510-00-00-19; src:usa-ri:ricr-510-00-00-20 | partially_verified |
| date-rule:usa-ri:fire-code-effective-2026-03-01 | Rhode Island Fire Code, Life Safety Code, and Fire Alarm Code | rule_effective_date | 2026-03-01 | Active RICR filings for current Title 450 fire-code package. | Not determined by effective-date field alone. | src:usa-ri:ricr-450-00-00-7; src:usa-ri:ricr-450-00-00-8; src:usa-ri:ricr-450-00-00-10 | partially_verified |
| date-rule:usa-ri:prior-code-drawings-specs | State Building Code amendments | transition / grace notice | Three months after adoption of an amendment | Owner or owner's agent notifies building official that drawings and specifications are based on prior code provisions. | Yes, if the statutory notice path applies; permit may be obtained under prior code provisions when drawings/specifications are completed. | src:usa-ri:statute-transition-drawings-specs | partially_verified |
| date-rule:usa-ri:rehab-permit-issuance-edition | State Rehabilitation Building and Fire Code for Existing Structures | permit_issuance_date | Code editions in effect at time of permit issuance | Applies to Part 20 references to Building, Mechanical, Plumbing, Fire Safety, Electrical, Boiler, Energy, Elevator, or Accessibility Code unless later amendments are expressly retroactive. | Prior editions not generally preserved by this rule; rule keys to permit issuance. | src:usa-ri:ricr-510-00-00-20 | partially_verified |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Electrical | Next NEC edition | null | null | null | null | null | monitor | src:usa-ri:statute-bcsc-adoption | Statute directs the committee to adopt the latest NEC so it takes effect on July 1 of the edition year; current active rule in this pass is NEC 2023 effective 2025-12-01. |
| Building / ICC family | Next ICC package | null | null | null | null | null | monitor | src:usa-ri:bcc-laws-rules | No official future package beyond the active Title 510 rules was captured. |
| Fire | Next NFPA fire package | null | null | null | null | null | monitor | src:usa-ri:ricr-450-00-00-7; src:usa-ri:ricr-450-00-00-8; src:usa-ri:ricr-450-00-00-10 | Current active fire package became effective 2026-03-01. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-ri:residential-scope | Residential | Detached one- and two-family dwellings, townhouses up to three stories above grade plane, and listed accessory structures | Project falls within RISBC-2 scope | RISBC-2 governs residential structures within its scope; other building types generally use RISBC-1 or other applicable parts. | src:usa-ri:ricr-510-00-00-2 | partially_verified |
| applicability-rule:usa-ri:existing-building-part-7 | Existing Building | Repairs, alterations, additions, changes of occupancy, moved buildings, historic buildings, and related existing-building work within the IEBC scope | Work on existing building subject to RISBC-7 | RISBC-7 adopts the 2021 IEBC and is the current existing-building code part. | src:usa-ri:ricr-510-00-00-7 | partially_verified |
| applicability-rule:usa-ri:rehab-ten-year | Rehabilitation | Qualifying existing buildings and structures in existence at least ten years before permit application, excluding listed occupancy groups and one-, two-, and three-family homes | Applicant uses RISRC-1 rehabilitation path | Part 20 applies only to qualifying existing buildings and links building/fire elements to IEBC and Life Safety Code editions in effect at permit issuance. | src:usa-ri:ricr-510-00-00-20 | partially_verified |
| applicability-rule:usa-ri:public-meetings-accessibility | Accessibility | Buildings constructed or occupied for public meetings before 1977 | Pre-1977 public-meeting building accessibility issue | RISBC-17 uses UFAS public-meetings accessibility standards; current new construction accessibility is in the State Building Code. | src:usa-ri:ricr-510-00-00-17 | partially_verified |
| applicability-rule:usa-ri:elevator-dlt | Elevator / Conveyance | Elevators, escalators, moving walks, dumbwaiters, hoists, lifts, vertical reciprocating conveyors, wind-turbine elevators, and similar devices | Device falls under Chapter 23-33 / 260-RICR-30-10-1 | DLT elevator rules, permits, inspections, and licenses apply; code interface with building/electrical requirements remains project-specific. | src:usa-ri:ricr-260-30-10-1 | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** Rhode Island adopts model codes by reference through state RICR parts and publishes Rhode Island amendments, deletions, substitutions, and administrative coordination text within those RICR parts.

**State amendment structure:** Official Rhode Island Code of Regulations pages for Title 510 Building Code Commission, Title 450 Fire Safety Code Board, and Title 260 Department of Labor and Training for elevators.

**Amendment parsing status:** identified_at_part_level; clause_by_clause_parsing_not_completed

### 5.2 State Amendment Sources

| Amendment Source ID | Code Family | Publication Path | Effective Date | Parsing Status | Source IDs |
| --- | --- | --- | --- | --- | --- |
| amend-source:usa-ri:risbc-1 | Building | 510-RICR-00-00-1 | 2025-12-01 | part_level_identified | src:usa-ri:ricr-510-00-00-1 |
| amend-source:usa-ri:risbc-2 | Residential | 510-RICR-00-00-2 | 2025-12-01 | part_level_identified | src:usa-ri:ricr-510-00-00-2 |
| amend-source:usa-ri:risbc-3 | Plumbing | 510-RICR-00-00-3 | 2025-12-01 | part_level_identified | src:usa-ri:ricr-510-00-00-3 |
| amend-source:usa-ri:risbc-4 | Mechanical | 510-RICR-00-00-4 | 2025-12-01 | part_level_identified | src:usa-ri:ricr-510-00-00-4 |
| amend-source:usa-ri:risbc-5 | Electrical | 510-RICR-00-00-5 | 2025-12-01 | part_level_identified | src:usa-ri:ricr-510-00-00-5 |
| amend-source:usa-ri:risbc-7 | Existing Building | 510-RICR-00-00-7 | 2025-12-01 | part_level_identified | src:usa-ri:ricr-510-00-00-7 |
| amend-source:usa-ri:risbc-8 | Energy | 510-RICR-00-00-8 | 2025-12-01 | part_level_identified | src:usa-ri:ricr-510-00-00-8 |
| amend-source:usa-ri:risbc-19 | Fuel Gas | 510-RICR-00-00-19 | 2025-12-01 | part_level_identified | src:usa-ri:ricr-510-00-00-19 |
| amend-source:usa-ri:risrc-1 | Rehabilitation | 510-RICR-00-00-20 | 2025-12-01 | part_level_identified | src:usa-ri:ricr-510-00-00-20 |
| amend-source:usa-ri:fire-code | Fire | 450-RICR-00-00-7; 450-RICR-00-00-8; 450-RICR-00-00-10 | 2026-03-01 | part_level_identified | src:usa-ri:ricr-450-00-00-7; src:usa-ri:ricr-450-00-00-8; src:usa-ri:ricr-450-00-00-10 |
| amend-source:usa-ri:elevator | Elevator / Conveyance | 260-RICR-30-10-1 | 2022-01-04 | part_level_identified | src:usa-ri:ricr-260-30-10-1 |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Summary | Source IDs | Confidence | Notes |
| --- | --- | --- | --- | --- | --- |
| amend:usa-ri:building-referenced-codes | Building | RISBC-1 references separate Rhode Island plumbing, mechanical, electrical, property maintenance, energy, pool/spa, fuel gas, fire safety, and existing-building rules, reinforcing a multi-part statewide code structure. | src:usa-ri:ricr-510-00-00-1 | 0.86 | Clause-by-clause conflicts with model-code chapters were not fully parsed. |
| amend:usa-ri:fire-separation | Building / Fire | Fire-code and fire-alarm subjects are not solely embedded in the building code; Title 450 rules and State Fire Marshal / Fire Safety Code Board authority must be resolved separately. | src:usa-ri:ricr-510-00-00-1; src:usa-ri:ricr-450-00-00-7; src:usa-ri:ricr-450-00-00-8; src:usa-ri:ricr-450-00-00-10 | 0.86 | Important for AHJ routing and compliance review. |
| amend:usa-ri:energy-2024-iecc-appendices | Energy | RISBC-8 adopts the 2024 IECC and notes adoption of commercial Appendix CH and residential Appendix RK. | src:usa-ri:ricr-510-00-00-8 | 0.85 | Full appendix text was not reproduced or parsed. |
| amend:usa-ri:rehab-relationship | Existing Building / Rehabilitation | RISRC-1 uses the IEBC adopted as Part 7 for the building-code element and the Rhode Island Life Safety Code Chapter 43 edition in effect at permit for the fire-code element. | src:usa-ri:ricr-510-00-00-20 | 0.86 | This is a key project-type applicability rule for existing-building rehabilitation. |
| amend:usa-ri:local-amendments-preempted | Local amendments | Local building-code enactment is prohibited; amendment petitions proceed to the Building Code Standards Committee and adopted amendments bind all cities and towns. | src:usa-ri:statute-local-code-preemption; src:usa-ri:statute-amendment-petition | 0.88 | Local enforcement remains distinct and locally implemented. |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-ri"
  model: "local_building_officials_enforce_state_code_with_state_commissioner_backstop"
  enforcing_entities:
    - "local building officials appointed by each local authority"
    - "State Building Code Commissioner and staff for state projects, statewide interpretation, and backstop enforcement where no local official or inspectors are available"
    - "State Fire Marshal and staff for strict enforcement of the Rhode Island Fire Code and Life Safety Code"
    - "Department of Labor and Training, Division of Occupational Safety / Chief Elevator Inspector for elevator and conveyance matters"
  required_officials:
    - "local building official"
    - "alternate local building official where appointed or state services are requested"
    - "certified building officials / inspectors as applicable under certification rules"
    - "certified fire marshal / deputies for fire-code elements where applicable"
  state_reserved_activities:
    - "statewide code adoption and amendments"
    - "official interpretations and standardization by the State Building Code Commissioner"
    - "state-project permitting and inspections"
    - "building-code enforcement backstop where no local official, alternate, or inspector is available"
    - "fire-code administration, variances, waivers, and amendments by the Fire Safety Code Board"
    - "strict fire-code enforcement by the State Fire Marshal"
    - "elevator permits, inspections, certificates, and related licensing by DLT"
  source_ids:
    - "src:usa-ri:statute-local-building-official-appointment"
    - "src:usa-ri:statute-local-building-official-duties"
    - "src:usa-ri:statute-state-commissioner-duties"
    - "src:usa-ri:ricr-510-00-00-9"
    - "src:usa-ri:ricr-510-00-00-10"
    - "src:usa-ri:ricr-450-00-00-7"
    - "src:usa-ri:ricr-450-00-00-8"
    - "src:usa-ri:ricr-260-30-10-1"
  verification_status: "partially_verified"
  confidence: 0.84
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-ri"
  model: "state_preemptive_no_future_local_building_codes"
  applies_to_code_families:
    - "Building"
    - "Residential"
    - "Existing Building / Rehabilitation"
    - "Mechanical"
    - "Plumbing"
    - "Fuel Gas"
    - "Electrical"
    - "Energy"
    - "construction-code accessibility provisions"
  approval_required: true
  approving_authority_id: "ahj:usa-ri:bcsc"
  filing_required: true
  registry_exists: true
  registry_source_ids:
    - "src:usa-ri:ricr-510-00-00-1"
    - "src:usa-ri:ricr-510-00-00-2"
    - "src:usa-ri:ricr-510-00-00-3"
    - "src:usa-ri:ricr-510-00-00-4"
    - "src:usa-ri:ricr-510-00-00-5"
    - "src:usa-ri:ricr-510-00-00-7"
    - "src:usa-ri:ricr-510-00-00-8"
    - "src:usa-ri:ricr-510-00-00-19"
  legal_basis_source_ids:
    - "src:usa-ri:statute-local-code-preemption"
    - "src:usa-ri:statute-amendment-petition"
  verification_status: "partially_verified"
  confidence: 0.86
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Rhode Island separates local administration from local lawmaking. Municipalities enforce the State Building Code through appointed local building officials, but state law prohibits future local building codes and ordinances. Proposed code changes are routed to the Building Code Standards Committee and, if adopted, apply statewide to all cities and towns. This distinction is important for address-to-AHJ resolution: local permits and inspections may be municipal, while the controlling code text remains state-level.

### 6.4 Known Local Amendment Registries

No separate municipal building-code amendment registry was verified. The controlling amendment registry for construction codes is the statewide Rhode Island Code of Regulations, especially Title 510 for Building Code Commission rules.

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendments were parsed. Because future local building-code enactment is prohibited, municipality-specific construction-code deviations should be treated as suspect until tied to a state-approved variance, local enforcement procedure, zoning ordinance, property-maintenance ordinance outside the State Building Code, or another non-building-code legal basis.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: partially_started

Jurisdiction stack:

```text
Address
  -> State: Rhode Island
  -> County: used for geography; not currently verified as default building-code AHJ
  -> Municipality / local authority
  -> Local building official for State Building Code permits and inspections
  -> State Building Code Commissioner for state projects, official interpretations, and backstop cases
  -> State Fire Marshal / Fire Safety Code Board for fire-code enforcement and administration
  -> Department of Labor and Training / Chief Elevator Inspector for elevator and conveyance matters
  -> Applicable Title 510 construction-code adoption records
  -> Applicable Title 450 fire-code adoption records
  -> Applicable Title 260 elevator-code record
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | not selected | none | statewide | unknown | pending |
| County | not selected | none | statewide | unknown | pending |
| Municipality | not selected | none | statewide | unknown | pending |
| Fire District | not selected | none | statewide | unknown | pending |
| Special District | not selected | none | statewide | unknown | pending |

### 7.3 AHJ Contact Data

| AHJ / Contact Entity | Scope | Contact Data Status | Source IDs | Notes |
| --- | --- | --- | --- | --- |
| Rhode Island Building Code Commission / State Building Office | Statewide code adoption, state building official functions, e-permitting, state-level resources | state_level_page_identified | src:usa-ri:bcc-laws-rules | Address-level municipal contacts were not populated. |
| State Fire Marshal / Fire Safety Code Board | Fire Code and Life Safety Code strict enforcement / administration | authority_verified_contact_directory_not_populated | src:usa-ri:ricr-450-00-00-7; src:usa-ri:ricr-450-00-00-8 | Local deputy/assistant deputy fire marshal routing was not populated. |
| Department of Labor and Training, Division of Occupational Safety / Elevator Unit | Elevator and conveyance code permits, inspections, and certificates | authority_verified_contact_directory_not_populated | src:usa-ri:ricr-260-30-10-1 | Project-specific elevator routing requires DLT contact data extraction. |

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title / Citation | Publisher | Source Type | Effective / Current Date Captured | URL | Used For |
| --- | --- | --- | --- | --- | --- | --- |
| src:usa-ri:bcc-laws-rules | Laws, Rules and Regulations | Rhode Island Building Code Commission / Official State of Rhode Island | agency page | page updated 2026-04-07; current code statement effective 2025-12-01 | https://ribcc.ri.gov/forms-resources-and-e-permitting/laws-rules-and-regulations | Overview that Rhode Island incorporated 2021 ICC codes with Rhode Island amendments effective 2025-12-01. |
| src:usa-ri:statute-bcsc-adoption | R.I. Gen. Laws § 23-27.3-100.1.5 | Rhode Island General Assembly | statute | current page reviewed 2026-06-26 | https://webserver.rilegislature.gov/Statutes/TITLE23/23-27.3/23-1/23-27.3-100.1.5.htm | BCSC authority to adopt/promulgate/administer; statutory code-family structure; NEC framework. |
| src:usa-ri:statute-local-code-preemption | R.I. Gen. Laws § 23-27.3-100.1.7 | Rhode Island General Assembly | statute | current page reviewed 2026-06-26 | https://webserver.rilegislature.gov/Statutes/TITLE23/23-27.3/23-1/23-27.3-100.1.7.htm | Prohibition on future city/town local building codes and ordinances. |
| src:usa-ri:statute-amendment-petition | R.I. Gen. Laws § 23-27.3-109.3 | Rhode Island General Assembly | statute | current page reviewed 2026-06-26 | https://webserver.rilegislature.gov/Statutes/TITLE23/23-27.3/23-1/23-27.3-109.3.htm | Statewide amendment petition process and binding effect on all cities and towns. |
| src:usa-ri:statute-local-building-official-appointment | R.I. Gen. Laws § 23-27.3-107.1 | Rhode Island General Assembly | statute | current page reviewed 2026-06-26 | https://webserver.rilegislature.gov/Statutes/TITLE23/23-27.3/23-1/23-27.3-107.1.htm | Local authority appointment of local building officials. |
| src:usa-ri:statute-alternate-local-building-official | R.I. Gen. Laws § 23-27.3-107.2 | Rhode Island General Assembly | statute | current page reviewed 2026-06-26 | https://webserver.rilegislature.gov/Statutes/TITLE23/23-27.3/23-1/23-27.3-107.2.htm | Alternate building official / state service fallback. |
| src:usa-ri:statute-local-building-official-duties | R.I. Gen. Laws §§ 23-27.3-107.5 and 23-27.3-108.1 | Rhode Island General Assembly | statute | current pages reviewed 2026-06-26 | https://webserver.rilegislature.gov/Statutes/TITLE23/23-27.3/23-1/23-27.3-107.5.htm; https://webserver.rilegislature.gov/Statutes/TITLE23/23-27.3/23-1/23-27.3-108.1.htm | Duties and enforcement role of local building official. |
| src:usa-ri:statute-state-commissioner-duties | R.I. Gen. Laws § 23-27.3-108.2 | Rhode Island General Assembly | statute | current page reviewed 2026-06-26 | https://webserver.rilegislature.gov/Statutes/TITLE23/23-27.3/23-1/23-27.3-108.2.htm | State Building Code Commissioner duties, interpretation, permit system, state/backstop enforcement. |
| src:usa-ri:statute-transition-drawings-specs | R.I. Gen. Laws § 23-27.3-114.3.1 | Rhode Island General Assembly | statute | current page reviewed 2026-06-26 | https://webserver.rilegislature.gov/Statutes/TITLE23/23-27.3/23-2/23-27.3-114.3.1.htm | Three-month drawings/specifications prior-code transition mechanism. |
| src:usa-ri:ricr-510-00-00-1 | RISBC-1 Rhode Island Building Code, 510-RICR-00-00-1 | Rhode Island Department of State / Building Code Commission | regulation | active effective 2025-12-01 | https://rules.sos.ri.gov/regulations/part/510-00-00-1 | Building code adoption, IBC 2021, referenced-code structure, amendments. |
| src:usa-ri:ricr-510-00-00-2 | RISBC-2 Rhode Island One- and Two-Family Dwelling Code, 510-RICR-00-00-2 | Rhode Island Department of State / Building Code Commission | regulation | active effective 2025-12-01 | https://rules.sos.ri.gov/regulations/part/510-00-00-2 | Residential code adoption, IRC 2021. |
| src:usa-ri:ricr-510-00-00-3 | RISBC-3 Rhode Island Plumbing Code, 510-RICR-00-00-3 | Rhode Island Department of State / Building Code Commission | regulation | active effective 2025-12-01 | https://rules.sos.ri.gov/regulations/part/510-00-00-3 | Plumbing code adoption, IPC 2021. |
| src:usa-ri:ricr-510-00-00-4 | RISBC-4 Rhode Island Mechanical Code, 510-RICR-00-00-4 | Rhode Island Department of State / Building Code Commission | regulation | active effective 2025-12-01 | https://rules.sos.ri.gov/regulations/part/510-00-00-4 | Mechanical code adoption, IMC 2021. |
| src:usa-ri:ricr-510-00-00-5 | RISBC-5 Rhode Island Electrical Code, 510-RICR-00-00-5 | Rhode Island Department of State / Building Code Commission | regulation | active effective 2025-12-01 | https://rules.sos.ri.gov/regulations/part/510-00-00-5 | Electrical code adoption, NEC 2023. |
| src:usa-ri:ricr-510-00-00-7 | RISBC-7 Rhode Island Existing Building Code, 510-RICR-00-00-7 | Rhode Island Department of State / Building Code Commission | regulation | active effective 2025-12-01 | https://rules.sos.ri.gov/regulations/part/510-00-00-7 | Existing building code adoption, IEBC 2021. |
| src:usa-ri:ricr-510-00-00-8 | RISBC-8 Rhode Island Energy Conservation Code, 510-RICR-00-00-8 | Rhode Island Department of State / Building Code Commission | regulation | active effective 2025-12-01 | https://rules.sos.ri.gov/regulations/part/510-00-00-8 | Energy code adoption, IECC 2024, appendix notes. |
| src:usa-ri:ricr-510-00-00-9 | 510-RICR-00-00-9 State Building Office Projects | Rhode Island Department of State / Building Code Commission | regulation | active page reviewed 2026-06-26 | https://rules.sos.ri.gov/regulations/part/510-00-00-9 | State Building Commissioner role for state projects and local plan-review requests. |
| src:usa-ri:ricr-510-00-00-10 | 510-RICR-00-00-10 Official Interpretations | Rhode Island Department of State / Building Code Commission | regulation | active page reviewed 2026-06-26 | https://rules.sos.ri.gov/regulations/part/510-00-00-10 | State Building Commissioner official interpretations and appeal path. |
| src:usa-ri:ricr-510-00-00-11 | 510-RICR-00-00-11 Certification of Building Officials, Building Inspectors and Plans Examiners | Rhode Island Department of State / Building Code Commission | regulation | active page reviewed 2026-06-26 | https://rules.sos.ri.gov/regulations/part/510-00-00-11 | Certification context for local building officials and inspectors. |
| src:usa-ri:ricr-510-00-00-17 | RISBC-17 Public Meetings Accessibility Standard, 510-RICR-00-00-17 | Rhode Island Department of State / Building Code Commission | regulation | active effective 2022-01-04 | https://rules.sos.ri.gov/regulations/part/510-00-00-17 | Public-meetings accessibility scope; new-construction accessibility cross-reference to State Building Code. |
| src:usa-ri:ricr-510-00-00-19 | RISBC-19 Rhode Island Fuel Gas Code, 510-RICR-00-00-19 | Rhode Island Department of State / Building Code Commission | regulation | active effective 2025-12-01 | https://rules.sos.ri.gov/regulations/part/510-00-00-19 | Fuel gas code adoption, IFGC 2021. |
| src:usa-ri:ricr-510-00-00-20 | RISRC-1 State Rehabilitation Building and Fire Code for Existing Structures, 510-RICR-00-00-20 | Rhode Island Department of State / Joint Committee on Rehabilitation Building and Fire Code | regulation | active effective 2025-12-01 | https://rules.sos.ri.gov/regulations/part/510-00-00-20 | Rehabilitation code authority, purpose, scope, permit-date rule, enforcement. |
| src:usa-ri:ricr-450-00-00-1 | Fire Safety Code Board Administrative Rules and Regulations, 450-RICR-00-00-1 | Rhode Island Department of State / Fire Safety Code Board | regulation | active page reviewed 2026-06-26 | https://rules.sos.ri.gov/regulations/part/450-00-00-1 | Fire Safety Code structure and administrative context. |
| src:usa-ri:ricr-450-00-00-7 | Rhode Island Fire Code, 450-RICR-00-00-7 | Rhode Island Department of State / Fire Safety Code Board | regulation | active effective 2026-03-01 | https://rules.sos.ri.gov/regulations/part/450-00-00-7 | NFPA 1 2021 adoption; Fire Safety Code Board / State Fire Marshal AHJ split. |
| src:usa-ri:ricr-450-00-00-8 | Rhode Island Life Safety Code, 450-RICR-00-00-8 | Rhode Island Department of State / Fire Safety Code Board | regulation | active effective 2026-03-01 | https://rules.sos.ri.gov/regulations/part/450-00-00-8 | NFPA 101 2021 adoption; Fire Safety Code Board / State Fire Marshal AHJ split. |
| src:usa-ri:ricr-450-00-00-10 | Rhode Island Fire Alarm Code, 450-RICR-00-00-10 | Rhode Island Department of State / Fire Safety Code Board | regulation | active effective 2026-03-01 | https://rules.sos.ri.gov/regulations/part/450-00-00-10 | NFPA 72 2022 adoption. |
| src:usa-ri:ricr-260-30-10-1 | Elevator Safety Code, 260-RICR-30-10-1 | Rhode Island Department of State / Department of Labor and Training | regulation | active effective 2022-01-04 | https://rules.sos.ri.gov/regulations/part/260-30-10-1 | Elevator and conveyance code authority, ASME incorporated standards, DLT permits/inspections. |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-ri:bcc-laws-rules | agency_summary | The agency page is useful as a current overview, but individual RICR parts are the controlling source for specific code editions, effective dates, and rule text. | cite_for_overview_only |
| src:usa-ri:ricr-510-00-00-1 | incorporated_model_code | The RICR page adopts the model code by reference and publishes state amendments; it does not reproduce all copyrighted model-code text. | consult_model_code_plus_RICR |
| src:usa-ri:ricr-510-00-00-2 | incorporated_model_code | The RICR page adopts the model code by reference and publishes state amendments; it does not reproduce all copyrighted model-code text. | consult_model_code_plus_RICR |
| src:usa-ri:ricr-510-00-00-3 | incorporated_model_code | The RICR page adopts the model code by reference and publishes state amendments; it does not reproduce all copyrighted model-code text. | consult_model_code_plus_RICR |
| src:usa-ri:ricr-510-00-00-4 | incorporated_model_code | The RICR page adopts the model code by reference and publishes state amendments; it does not reproduce all copyrighted model-code text. | consult_model_code_plus_RICR |
| src:usa-ri:ricr-510-00-00-5 | incorporated_model_code | The RICR page adopts NFPA 70 by reference and publishes state amendments; it does not reproduce all copyrighted NEC text. | consult_NEC_plus_RICR |
| src:usa-ri:ricr-510-00-00-7 | incorporated_model_code | The RICR page adopts the IEBC by reference; full compliance review requires both the IEBC and Rhode Island amendments. | consult_model_code_plus_RICR |
| src:usa-ri:ricr-510-00-00-8 | incorporated_model_code | The RICR page adopts the IECC by reference and does not reproduce all copyrighted model-code text. | consult_model_code_plus_RICR |
| src:usa-ri:ricr-450-00-00-7 | incorporated_model_code | The RICR page adopts NFPA 1 by reference and publishes Rhode Island provisions; full compliance review requires NFPA 1 and the RICR rule. | consult_NFPA_plus_RICR |
| src:usa-ri:ricr-450-00-00-8 | incorporated_model_code | The RICR page adopts NFPA 101 by reference and publishes Rhode Island provisions; full compliance review requires NFPA 101 and the RICR rule. | consult_NFPA_plus_RICR |
| src:usa-ri:ricr-450-00-00-10 | incorporated_model_code | The RICR page adopts NFPA 72 by reference and publishes Rhode Island provisions; full compliance review requires NFPA 72 and the RICR rule. | consult_NFPA_plus_RICR |
| src:usa-ri:ricr-260-30-10-1 | incorporated_model_code | The elevator rule adopts several ASME standards by reference; full compliance review requires the ASME standards and the RICR rule. | consult_ASME_plus_RICR |
| src:usa-ri:statute-transition-drawings-specs | interpretation_needed | The statute supports a prior-code notice path but does not itself provide a simple universal mandatory date for all projects. | use_as_transition_rule_not_blanket_deadline |

### 8.3 Supplemental Sources

None used as authority in the populated report. Non-official municipal pages were observed during research but were not used as controlling sources.

### 8.4 Source Extraction Metadata

| Extraction ID | Source IDs | Extracted Fields | Extraction Date | Extractor | Notes |
| --- | --- | --- | --- | --- | --- |
| extract:usa-ri:authority-2026-06-26 | src:usa-ri:statute-bcsc-adoption; src:usa-ri:statute-local-code-preemption; src:usa-ri:statute-amendment-petition | authority model, local amendment posture | 2026-06-26 | ChatGPT | Statute-level authority fields extracted from official General Assembly pages. |
| extract:usa-ri:title-510-2026-06-26 | src:usa-ri:ricr-510-00-00-1; src:usa-ri:ricr-510-00-00-2; src:usa-ri:ricr-510-00-00-3; src:usa-ri:ricr-510-00-00-4; src:usa-ri:ricr-510-00-00-5; src:usa-ri:ricr-510-00-00-7; src:usa-ri:ricr-510-00-00-8; src:usa-ri:ricr-510-00-00-19; src:usa-ri:ricr-510-00-00-20 | construction-code editions and effective dates | 2026-06-26 | ChatGPT | Code family rows populated from active RICR rule pages. |
| extract:usa-ri:title-450-2026-06-26 | src:usa-ri:ricr-450-00-00-7; src:usa-ri:ricr-450-00-00-8; src:usa-ri:ricr-450-00-00-10 | fire-code editions, effective date, AHJ split | 2026-06-26 | ChatGPT | Fire package updated to 2021/2022 NFPA editions effective 2026-03-01. |
| extract:usa-ri:elevator-2026-06-26 | src:usa-ri:ricr-260-30-10-1 | elevator-code authority, ASME editions, effective date | 2026-06-26 | ChatGPT | Elevator source included because template contains an elevator/conveyance row. |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 1.00 | none | Core authority and adoption fields are now source-backed; address-level AHJ routing remains incomplete. |
| report | risk.overall_confidence | 0.72 | verified | 1.00 | none | Confidence reflects strong statewide authority/code evidence but incomplete local contact and clause-level amendment parsing. |
| ahj:usa-ri:bcsc | authority_name | Building Code Standards Committee | partially_verified | 0.92 | src:usa-ri:statute-bcsc-adoption | Statewide adopting authority verified. |
| adoption:usa-ri:building:2021-ibc | effective_date | 2025-12-01 | partially_verified | 0.92 | src:usa-ri:ricr-510-00-00-1; src:usa-ri:bcc-laws-rules | Active RICR filing and BCC overview agree on current package effective date. |
| adoption:usa-ri:residential:2021-irc | edition | 2021 IRC | partially_verified | 0.91 | src:usa-ri:ricr-510-00-00-2 | Active RICR rule verified. |
| adoption:usa-ri:electrical:2023-nec | edition | 2023 NEC | partially_verified | 0.90 | src:usa-ri:ricr-510-00-00-5 | Active RICR rule verified. |
| adoption:usa-ri:energy:2024-iecc | edition | 2024 IECC | partially_verified | 0.90 | src:usa-ri:ricr-510-00-00-8 | Active RICR rule verified. |
| adoption:usa-ri:fire:2021-nfpa-1 | effective_date | 2026-03-01 | partially_verified | 0.90 | src:usa-ri:ricr-450-00-00-7 | Active RICR rule verified. |
| adoption:usa-ri:life-safety:2021-nfpa-101 | effective_date | 2026-03-01 | partially_verified | 0.90 | src:usa-ri:ricr-450-00-00-8 | Active RICR rule verified. |
| adoption:usa-ri:fire-alarm:2022-nfpa-72 | effective_date | 2026-03-01 | partially_verified | 0.90 | src:usa-ri:ricr-450-00-00-10 | Active RICR rule verified. |
| local-enforcement:usa-ri | model | local_building_officials_enforce_state_code_with_state_commissioner_backstop | partially_verified | 0.84 | src:usa-ri:statute-local-building-official-appointment; src:usa-ri:statute-local-building-official-duties; src:usa-ri:statute-state-commissioner-duties | Municipal contact directory remains open. |
| local-amendment-rule:usa-ri | model | state_preemptive_no_future_local_building_codes | partially_verified | 0.86 | src:usa-ri:statute-local-code-preemption; src:usa-ri:statute-amendment-petition | Building-code local amendment scope verified; non-building-code local ordinances not analyzed. |
| date-rule:usa-ri:prior-code-drawings-specs | transition_rule | three-month prior-code drawings/specifications notice path | partially_verified | 0.80 | src:usa-ri:statute-transition-drawings-specs | Treated as transition mechanism, not universal mandatory date. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Every `src:usa-ri:*` cited in the body appears in Section 8. |
| All authority IDs resolve | pass | Authorities are defined in Section 2 or intentionally named as local/system entities. |
| All current code families have adoption records | pass | Every template code-family row has either a normalized adoption record or an explicit caveat. |
| Building and operational fire code are separated | pass | Title 510 construction-code rows and Title 450 fire-code rows are distinct. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Unknown adoption, operative, and mandatory dates remain null when not separately sourced. |
| Effective dates are valid ISO dates | pass | Date fields entered as YYYY-MM-DD. |
| No impossible date sequences | pass | No unsupported adoption-to-effective sequence was introduced. |
| Transition rules have explicit trigger conditions | pass | The prior-code rule is tied to post-amendment owner/agent notice to the building official. |
| Permit-date logic is captured where applicable | pass | Part 20 permit-issuance edition rule is captured. |
| Local enforcement model classified | pass | Local building official model with state backstop is classified. |
| Local amendment rule classified | pass | State-preemptive local amendment posture is classified. |
| AHJ confirmation metadata present | partial | State-level AHJs are identified; municipality-level contact data remains open. |
| Official-source caveats captured | pass | Incorporated-by-reference and agency-summary caveats are documented. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-ri:001 | medium | local AHJ routing | Municipality-level building official contacts, service areas, and e-permitting routing were not populated. | Extract municipal/local building official directory and test address-to-AHJ routing. | null | null | open |
| issue:usa-ri:002 | medium | fire AHJ contacts | Fire authority model is verified, but local deputy/assistant deputy state fire marshal contact routing was not populated. | Extract State Fire Marshal / fire marshal contact directory and confirm routing rules. | null | null | open |
| issue:usa-ri:003 | medium | state amendments | State amendment sources are identified at the rule-part level, but high-impact amendments were not parsed clause-by-clause. | Parse each active Title 510 and Title 450 rule against the base model code structure. | null | null | open |
| issue:usa-ri:004 | medium | transition dates | The statutory three-month prior-code path was captured, but agency guidance was not collected to convert the 2025-12-01 construction-code package into project-level acceptance deadlines. | Request or locate State Building Office guidance and confirm permit-system behavior. | null | null | open |
| issue:usa-ri:005 | low | boundary data | Boundary datasets for municipalities, fire districts, and special districts were not selected. | Select authoritative GIS/boundary sources and update Section 7.2. | null | null | open |
| issue:usa-ri:006 | low | elevator routing | DLT elevator authority and code editions are verified, but inspection-unit contacts and permit workflow details were not populated. | Extract DLT elevator unit contact, permit, inspection, and certificate workflow data. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-ri:bcc-laws-rules | src:usa-ri:bcc-laws-rules | html_diff | monthly | Agency page announces new code package, transition guidance, or portal restructuring. | 2026-06-26 | active |
| watch:usa-ri:title-510 | src:usa-ri:ricr-510-00-00-1 | ricr_rule_diff | monthly | Title 510 active-rule effective dates, incorporated editions, or amendments change. | 2026-06-26 | active |
| watch:usa-ri:fire-title-450 | src:usa-ri:ricr-450-00-00-7 | ricr_rule_diff | monthly | Fire Code, Life Safety Code, or Fire Alarm Code edition/effective-date changes. | 2026-06-26 | active |
| watch:usa-ri:nec-statutory-cycle | src:usa-ri:statute-bcsc-adoption | statute_and_rule_monitor | quarterly | NEC statutory adoption framework or RISBC-5 edition changes. | 2026-06-26 | active |
| watch:usa-ri:elevator-title-260 | src:usa-ri:ricr-260-30-10-1 | ricr_rule_diff | quarterly | Elevator ASME editions, permit rules, or inspection requirements change. | 2026-06-26 | active |
| watch:usa-ri:transition-guidance | src:usa-ri:statute-transition-drawings-specs | agency_guidance_monitor | monthly | State Building Office publishes or updates transition/grace-period guidance for current code package. | 2026-06-26 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-26 | Populated Rhode Island report from baseline draft; upgraded status to partially_verified; added official source registry, authority model, adoption matrix, date rules, local amendment model, and QA checks. | report:usa-ri; ahj:usa-ri:bcsc; adoption:usa-ri:building:2021-ibc; adoption:usa-ri:fire:2021-nfpa-1 | src:usa-ri:statute-bcsc-adoption; src:usa-ri:bcc-laws-rules; src:usa-ri:ricr-510-00-00-1; src:usa-ri:ricr-450-00-00-7 | ChatGPT | Open issues retained for local contacts, clause-level amendments, and project-level transition guidance. |
