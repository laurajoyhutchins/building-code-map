---
state:
  state_id: "US-MA"
  name: "Massachusetts"
  abbreviation: "MA"
report:
  report_id: "state-report:usa-ma"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-25"
  last_verified: null
  reviewed_by: null
risk:
  overall_confidence: 0.66 # 0.00 - 1.00
  risk_flags:
    - "official_massgov_regulation_pages_are_unofficial_copies"
    - "electrical_code_effective_and_permit_trigger_dates_require_legal_review"
    - "local_amendment_scope_partially_resolved_only"
    - "municipal_energy_code_adoption_registry_not_parsed"
  open_questions_count: 5

---

# State Building Code Authority Report: Massachusetts

## 1. Executive Summary

- **Authority model:** Massachusetts uses a statewide code model. The State Board of Building Regulations and Standards (BBRS) is the primary statewide building-code adopting and administering authority for the State Building Code, codified at 780 CMR. The Board of Fire Prevention Regulations (BFPR), Board of State Examiners of Plumbers and Gas Fitters, Architectural Access Board, and Board of Elevator Regulations administer important specialized code areas.

- **Statewide code status:** The 10th Edition of the Massachusetts State Building Code is the current statewide building code. It became effective on 2024-10-11, and the 9th/10th Edition concurrency period ended on 2025-06-30, making the 10th Edition mandatory for building permit applications filed on or after 2025-07-01.

- **Local enforcement model:** Building-code enforcement is local for ordinary city and town work: each city or town employs a building commissioner or inspector of buildings and local inspectors to administer and enforce the State Building Code. State inspectors have authority for Commonwealth-owned buildings and supervisory review functions.

- **Local amendment posture:** For the base building code, Massachusetts is treated here as a statewide/preemptive code model because the statute provides that the State Building Code is the code for all buildings and structures within the city or town. Local variation is separately recognized for energy-code pathways where municipalities may adopt the Stretch Code or Specialized Opt-in Code, and for historic/architectural-district exterior-feature conflicts. This pass did not identify a general municipal building-code amendment mechanism outside those supported exceptions.

- **Known transition periods or pending changes:** The 10th Edition building-code concurrency period ended 2025-06-30. The current Massachusetts Electrical Code is the 2026 NFPA 70-based 527 CMR 12.00, labeled effective 2026-04-24; Rule 11 states it applies to installations for which a permit was granted after 2026-02-28. BFPR has announced that it is no longer accepting proposals for the 2024 Edition of 527 CMR 1.00 and expects to accept proposals for the 2027 Edition later in 2026.

- **Production readiness:** partially_ready_for_narrow_validation

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: BBRS is established in the Division of Occupational Licensure and must
    adopt and administer the State Building Code.
  confidence: 0.95
  source_ids:
  - src:usa-ma:mgl-c143-s93
- topic: Primary building code edition
  finding: The current State Building Code is 780 CMR 10th Edition, based on modified
    2021 I-Codes, effective 2024-10-11 and mandatory after the 2025-06-30 concurrency
    period.
  confidence: 0.9
  source_ids:
  - src:usa-ma:780-cmr-10th-handbook
  - src:usa-ma:780-cmr-10th-toc
- topic: Electrical code authority
  finding: BFPR promulgates electrical wiring rules; 527 CMR 12.00 is the Massachusetts
    Electrical Code and is currently based on NFPA 70, 2026 edition.
  confidence: 0.85
  source_ids:
  - src:usa-ma:mgl-c143-s3l
  - src:usa-ma:527-cmr-12-2026
- topic: Fire code authority
  finding: BFPR is required to promulgate a comprehensive fire safety code; current
    operational code is 527 CMR 1.00 based on NFPA 1, 2021 edition with Massachusetts
    amendments.
  confidence: 0.85
  source_ids:
  - src:usa-ma:mgl-c22d-s4
  - src:usa-ma:527-cmr-1-2021
  - src:usa-ma:fire-code-page
- topic: Local enforcement
  finding: City and town building commissioners/inspectors and local inspectors administer
    and enforce the State Building Code.
  confidence: 0.9
  source_ids:
  - src:usa-ma:mgl-c143-s3
  - src:usa-ma:mgl-c143-s3a
- topic: Local amendments
  finding: General local building-code amendment authority was not identified; supported
    exceptions include municipal energy-code adoption paths and historic/architectural
    district exterior features.
  confidence: 0.62
  source_ids:
  - src:usa-ma:mgl-c143-s3a
  - src:usa-ma:energy-codes
  - src:usa-ma:municipal-energy-registry
- topic: Effective / operative date rule
  finding: '10th Edition building-code transition is permit-application based: 9th
    Edition projects had to file permit applications on or before 2025-06-30.'
  confidence: 0.88
  source_ids:
  - src:usa-ma:780-cmr-10th-handbook
  - src:usa-ma:10th-edition-concurrency-extension
  - src:usa-ma:780-cmr-residential-ch1
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-ma:bbrs |
| Authority name | State Board of Building Regulations and Standards |
| Authority type | statewide_board |
| Legal basis | M.G.L. c. 143, §§ 93-94 |
| Role | Adopts, administers, revises, and amends the Massachusetts State Building Code, including the non-specialized building-code provisions and energy provisions integrated into the State Building Code. |
| Enforcement model | Local enforcement by city/town building commissioner or inspector of buildings, with state inspector authority and supervisory review for specified state-owned or state-supervised work. |
| Source IDs | src:usa-ma:mgl-c143-s93; src:usa-ma:mgl-c143-s94; src:usa-ma:mgl-c143-s3; src:usa-ma:mgl-c143-s3a |
| Verification status | verified_core |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-ma:bbrs | State Board of Building Regulations and Standards | Adopt/administer State Building Code, including commercial building provisions. | M.G.L. c. 143, §§ 93-94 | src:usa-ma:mgl-c143-s93; src:usa-ma:mgl-c143-s94; src:usa-ma:780-cmr-10th-handbook | verified_core |
| Residential | ahj:usa-ma:bbrs | State Board of Building Regulations and Standards | Adopt/administer 780 CMR residential provisions. | M.G.L. c. 143, §§ 93-94 | src:usa-ma:mgl-c143-s93; src:usa-ma:780-cmr-residential-ch1 | verified_core |
| Existing Building / Rehabilitation | ahj:usa-ma:bbrs | State Board of Building Regulations and Standards | Adopt/administer 780 CMR Chapter 34 / IEBC-based existing-building provisions. | M.G.L. c. 143, § 94(a)(ii) | src:usa-ma:mgl-c143-s94; src:usa-ma:780-cmr-ch34 | verified_core |
| Mechanical | ahj:usa-ma:bbrs | State Board of Building Regulations and Standards | Adopt/administer mechanical-system provisions generally governed by 2021 IMC through 780 CMR. | M.G.L. c. 143, §§ 93-94 | src:usa-ma:780-cmr-10th-ch1; src:usa-ma:780-cmr-ch28 | partially_verified |
| Plumbing | ahj:usa-ma:plumbers-gasfitters-board | Board of State Examiners of Plumbers and Gas Fitters | Regulates plumbing and administers 248 CMR plumbing rules; local plumbing inspectors inspect permitted work. | M.G.L. c. 142; 248 CMR | src:usa-ma:mgl-c142-s11; src:usa-ma:248-cmr-10; src:usa-ma:plumbers-board-page | partially_verified |
| Fuel Gas | ahj:usa-ma:plumbers-gasfitters-board | Board of State Examiners of Plumbers and Gas Fitters | Regulates gas fitting and Massachusetts Fuel Gas Code provisions in 248 CMR 4.00-8.00. | M.G.L. c. 142; 248 CMR | src:usa-ma:248-cmr-4; src:usa-ma:248-cmr-5; src:usa-ma:248-cmr-8; src:usa-ma:plumbers-board-page | partially_verified |
| Electrical | ahj:usa-ma:bfpr | Board of Fire Prevention Regulations | Promulgates Massachusetts Electrical Code, 527 CMR 12.00; local inspectors of wires and state examiners enforce. | M.G.L. c. 143, § 3L | src:usa-ma:mgl-c143-s3l; src:usa-ma:527-cmr-12-2026 | verified_core |
| Energy | ahj:usa-ma:bbrs-doer | BBRS in consultation with Department of Energy Resources | Integrates IECC-based energy provisions into 780 CMR and supports Base, Stretch, and Specialized Opt-in energy-code pathways. | M.G.L. c. 143, § 94(o)-(q); 225 CMR 22-23; 780 CMR | src:usa-ma:mgl-c143-s94; src:usa-ma:energy-codes; src:usa-ma:225-cmr-22; src:usa-ma:225-cmr-23 | partially_verified |
| Fire - construction references | ahj:usa-ma:bbrs | State Board of Building Regulations and Standards | Construction fire/life-safety requirements embedded in the State Building Code, with specialized-code coordination. | M.G.L. c. 143, §§ 93-94 | src:usa-ma:mgl-c143-s93; src:usa-ma:mgl-c143-s94; src:usa-ma:780-cmr-10th-handbook | partially_verified |
| Fire - operational / prevention code | ahj:usa-ma:bfpr | Board of Fire Prevention Regulations | Promulgates Massachusetts Comprehensive Fire Safety Code, 527 CMR 1.00. | M.G.L. c. 22D, § 4 | src:usa-ma:mgl-c22d-s4; src:usa-ma:fire-code-page; src:usa-ma:527-cmr-1-2021 | verified_core |
| Accessibility | ahj:usa-ma:aab | Architectural Access Board | Promulgates accessibility regulations, 521 CMR, treated as a specialized code. | M.G.L. c. 22, § 13A | src:usa-ma:mgl-c22-s13a; src:usa-ma:521-cmr-physical-accessibility | partially_verified |
| Elevator / Conveyance | ahj:usa-ma:elevator-board | Board of Elevator Regulations / Commissioner of Division of Occupational Licensure | Regulates construction, installation, alteration, inspection, and operation of elevators and associated conveyances. | M.G.L. c. 143, § 62; 524 CMR | src:usa-ma:mgl-c143-s62; src:usa-ma:524-cmr-elevator-list; src:usa-ma:524-cmr-35 | partially_verified |

### 2.3 Authority Hierarchy Notes

Massachusetts uses a statewide building-code model with local administration. BBRS is the central State Building Code authority, but Massachusetts also uses specialized boards for electrical, operational fire prevention, plumbing/gas fitting, accessibility, and elevators. Chapter 143 distinguishes the State Building Code from "specialized codes," so the building-code report should not treat all code families as exclusively controlled by BBRS.

Local enforcement and local amendment authority are separate questions. City and town building departments enforce the statewide State Building Code, but that enforcement role does not by itself establish authority to amend the State Building Code. Energy-code adoption is an exception-like local pathway because municipalities can be Base, Stretch, or Specialized Opt-in code communities under the state energy-code framework.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-ma:001 | ahj:usa-ma:bbrs | adopts_and_administers | state_building_code:780-cmr | src:usa-ma:mgl-c143-s93; src:usa-ma:mgl-c143-s94 | verified |
| edge:usa-ma:002 | ahj:usa-ma:bbrs | delegates_or_relies_on_enforcement_by | local_building_commissioners_and_inspectors | src:usa-ma:mgl-c143-s3; src:usa-ma:mgl-c143-s3a | verified |
| edge:usa-ma:003 | ahj:usa-ma:bfpr | promulgates | state_fire_code:527-cmr-1 | src:usa-ma:mgl-c22d-s4; src:usa-ma:527-cmr-1-2021 | verified |
| edge:usa-ma:004 | ahj:usa-ma:bfpr | promulgates | state_electrical_code:527-cmr-12 | src:usa-ma:mgl-c143-s3l; src:usa-ma:527-cmr-12-2026 | verified |
| edge:usa-ma:005 | ahj:usa-ma:plumbers-gasfitters-board | regulates | plumbing_and_gas_fitting:248-cmr | src:usa-ma:plumbers-board-page; src:usa-ma:248-cmr-10; src:usa-ma:248-cmr-4 | partially_verified |
| edge:usa-ma:006 | ahj:usa-ma:aab | promulgates | accessibility_code:521-cmr | src:usa-ma:mgl-c22-s13a; src:usa-ma:521-cmr-physical-accessibility | partially_verified |
| edge:usa-ma:007 | ahj:usa-ma:elevator-board | regulates | elevators_and_conveyances:524-cmr | src:usa-ma:mgl-c143-s62; src:usa-ma:524-cmr-elevator-list | partially_verified |
| edge:usa-ma:008 | local_municipalities | may_select_energy_pathway | base_or_stretch_or_specialized_energy_code | src:usa-ma:energy-codes; src:usa-ma:municipal-energy-registry; src:usa-ma:225-cmr-22; src:usa-ma:225-cmr-23 | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | 780 CMR, 10th Edition Massachusetts State Building Code, commercial/base provisions | International Building Code | 2021 | current_statewide | null | 2024-10-11 | 2024-10-11 | 2025-07-01 | 9th and 10th Editions concurrent through 2025-06-30; permit applications using 9th Edition had to be filed on or before 2025-06-30. | src:usa-ma:780-cmr-10th-handbook; src:usa-ma:10th-edition-concurrency-extension; src:usa-ma:780-cmr-10th-toc |
| Residential | 780 CMR 51.00, Massachusetts Residential Code | International Residential Code, with Massachusetts amendments and related I-Code references | 2021 | current_statewide | null | 2024-10-11 | 2024-10-11 | 2025-07-01 | Residential permit applications and related documents filed through 2025-06-30 could use either the prior or current 780 CMR edition. | src:usa-ma:780-cmr-residential-ch1; src:usa-ma:780-cmr-10th-handbook |
| Existing Building / Rehabilitation | 780 CMR Chapter 34, Existing Building Code | International Existing Building Code | 2021 | current_statewide | null | 2024-10-11 | 2024-10-11 | 2025-07-01 | Same 10th Edition concurrency rule; Chapter 34 applies to repair, alteration, change of occupancy, addition, and relocation of existing buildings. | src:usa-ma:780-cmr-ch34; src:usa-ma:780-cmr-10th-handbook |
| Mechanical | 780 CMR mechanical provisions | International Mechanical Code | 2021 | current_statewide | null | 2024-10-11 | 2024-10-11 | 2025-07-01 | Same 10th Edition concurrency rule. | src:usa-ma:780-cmr-10th-ch1; src:usa-ma:780-cmr-ch28; src:usa-ma:780-cmr-10th-handbook |
| Plumbing | 248 CMR 10.00, Uniform State Plumbing Code | state-specific plumbing code | current 248 CMR 10.00; model edition not identified | current_statewide | null | 2023-12-08 | 2023-12-08 | 2023-12-08 | New regulation changes to 248 CMR 10.00 went into effect 2023-12-08; project-specific transition rule not parsed. | src:usa-ma:248-cmr-10; src:usa-ma:plumbing-2023-faq |
| Fuel Gas | 248 CMR 4.00-8.00, Massachusetts Fuel Gas Code | NFPA 54 and NFPA 58 amendments plus Massachusetts provisions | editions not fully parsed | current_statewide | null | 2021-04-30 | 2021-04-30 | 2021-04-30 | 248 CMR 4.00 through 8.00 govern fuel-gas piping systems and utilization equipment; detailed transition rule not parsed. | src:usa-ma:248-cmr-4; src:usa-ma:248-cmr-5; src:usa-ma:248-cmr-8 |
| Electrical | 527 CMR 12.00, Massachusetts Electrical Code | NFPA 70, National Electrical Code | 2026 | current_statewide | null | 2026-04-24 | 2026-03-01 | 2026-03-01 | Source labels code "Effective 4/24/26"; Rule 11 states applicability to installations with permits granted after 2026-02-28. Date sequence needs legal review. | src:usa-ma:527-cmr-12-2026; src:usa-ma:electrical-code-page |
| Energy | Massachusetts Base Energy Code within 780 CMR; Stretch and Specialized Opt-in energy codes in 225 CMR 22-23 | International Energy Conservation Code | 2021 | current_statewide_with_municipal_options | null | 2024-10-11 | 2024-10-11 | 2025-07-01 | Base code follows 780 CMR 10th Edition; municipalities may have Stretch or Specialized Opt-in status and effective dates in state registry. | src:usa-ma:energy-codes; src:usa-ma:225-cmr-22; src:usa-ma:225-cmr-23; src:usa-ma:municipal-energy-registry |
| Fire - construction references | Fire/life-safety construction provisions in 780 CMR | International Building Code and related I-Codes | 2021 | current_statewide | null | 2024-10-11 | 2024-10-11 | 2025-07-01 | Same 10th Edition concurrency rule; operational fire-code permits remain separately governed by 527 CMR. | src:usa-ma:780-cmr-10th-handbook; src:usa-ma:780-cmr-10th-toc |
| Fire - operational / prevention code | 527 CMR 1.00, Massachusetts Comprehensive Fire Safety Code | NFPA 1 Fire Code | 2021 | current_statewide | null | 2023-05-12 | 2023-05-12 | 2023-05-12 | Current Mass.gov fire-code page identifies the 2021 Edition PDF as effective 2023-05-12; earlier advisory references 2022-12-09 revisions and should be reconciled before verification. | src:usa-ma:fire-code-page; src:usa-ma:527-cmr-1-2021; src:usa-ma:fire-code-advisory-2022 |
| Accessibility | 521 CMR, Architectural Access Board regulations | Massachusetts accessibility regulations | current 521 CMR; edition/date not fully parsed | current_statewide_specialized_code | null | null | null | null | Date fields unresolved; statute and current Mass.gov accessibility page support authority and code family. | src:usa-ma:mgl-c22-s13a; src:usa-ma:521-cmr-physical-accessibility; src:usa-ma:521-cmr-1996 |
| Elevator / Conveyance | 524 CMR, Board of Elevator Regulations | ASME A17.1-2013 / CSA B44-13 and ASME A18.1-2014 with Massachusetts modifications, among other chapters | mixed model editions | current_statewide_specialized_code | null | 2018-06-01 | 2018-06-01 | 2018-06-01 | Mass.gov elevator regulations list is labeled effective 2018-06-01; detailed chapter-by-chapter transition not parsed. | src:usa-ma:524-cmr-elevator-list; src:usa-ma:524-cmr-35; src:usa-ma:mgl-c143-s62 |

### 3.2 Adoption Records

```yaml
adoptions:
  - adoption_id: adoption:usa-ma:780-cmr-10th-building
    code_family: Building
    authority_id: ahj:usa-ma:bbrs
    state_code_name: "780 CMR, 10th Edition Massachusetts State Building Code"
    base_model_code: "International Building Code"
    base_model_edition: "2021"
    adoption_date: null
    effective_date: "2024-10-11"
    operative_date: "2024-10-11"
    mandatory_date: "2025-07-01"
    transition_rule_id: date-rule:usa-ma:780-cmr-10th-concurrency
    source_ids:
      - src:usa-ma:780-cmr-10th-handbook
      - src:usa-ma:10th-edition-concurrency-extension
      - src:usa-ma:780-cmr-10th-toc
    confidence: 0.90

  - adoption_id: adoption:usa-ma:780-cmr-10th-residential
    code_family: Residential
    authority_id: ahj:usa-ma:bbrs
    state_code_name: "780 CMR 51.00, Massachusetts Residential Code"
    base_model_code: "International Residential Code and related I-Code references"
    base_model_edition: "2021"
    adoption_date: null
    effective_date: "2024-10-11"
    operative_date: "2024-10-11"
    mandatory_date: "2025-07-01"
    transition_rule_id: date-rule:usa-ma:780-cmr-10th-concurrency
    source_ids:
      - src:usa-ma:780-cmr-residential-ch1
      - src:usa-ma:780-cmr-10th-handbook
    confidence: 0.84

  - adoption_id: adoption:usa-ma:780-cmr-10th-existing-building
    code_family: Existing Building / Rehabilitation
    authority_id: ahj:usa-ma:bbrs
    state_code_name: "780 CMR Chapter 34, Existing Building Code"
    base_model_code: "International Existing Building Code"
    base_model_edition: "2021"
    adoption_date: null
    effective_date: "2024-10-11"
    operative_date: "2024-10-11"
    mandatory_date: "2025-07-01"
    transition_rule_id: date-rule:usa-ma:780-cmr-10th-concurrency
    source_ids:
      - src:usa-ma:780-cmr-ch34
      - src:usa-ma:780-cmr-10th-handbook
    confidence: 0.84

  - adoption_id: adoption:usa-ma:527-cmr-12-2026
    code_family: Electrical
    authority_id: ahj:usa-ma:bfpr
    state_code_name: "527 CMR 12.00, Massachusetts Electrical Code"
    base_model_code: "NFPA 70, National Electrical Code"
    base_model_edition: "2026"
    adoption_date: null
    effective_date: "2026-04-24"
    operative_date: "2026-03-01"
    mandatory_date: "2026-03-01"
    transition_rule_id: date-rule:usa-ma:527-cmr-12-2026-permit
    source_ids:
      - src:usa-ma:527-cmr-12-2026
      - src:usa-ma:electrical-code-page
    confidence: 0.78
    caveat: "Effective-date label and permit-trigger date should be reconciled before verified status."

  - adoption_id: adoption:usa-ma:527-cmr-1-2021
    code_family: Fire - operational / prevention code
    authority_id: ahj:usa-ma:bfpr
    state_code_name: "527 CMR 1.00, Massachusetts Comprehensive Fire Safety Code"
    base_model_code: "NFPA 1 Fire Code"
    base_model_edition: "2021"
    adoption_date: null
    effective_date: "2023-05-12"
    operative_date: "2023-05-12"
    mandatory_date: "2023-05-12"
    transition_rule_id: null
    source_ids:
      - src:usa-ma:fire-code-page
      - src:usa-ma:527-cmr-1-2021
    confidence: 0.80
    caveat: "Earlier 2022 advisory references a 2022-12-09 revision; keep date caveat until official register history is reviewed."

  - adoption_id: adoption:usa-ma:248-cmr-10
    code_family: Plumbing
    authority_id: ahj:usa-ma:plumbers-gasfitters-board
    state_code_name: "248 CMR 10.00, Uniform State Plumbing Code"
    base_model_code: "state-specific"
    base_model_edition: null
    adoption_date: null
    effective_date: "2023-12-08"
    operative_date: "2023-12-08"
    mandatory_date: "2023-12-08"
    transition_rule_id: null
    source_ids:
      - src:usa-ma:248-cmr-10
      - src:usa-ma:plumbing-2023-faq
    confidence: 0.72

  - adoption_id: adoption:usa-ma:248-cmr-4-through-8
    code_family: Fuel Gas
    authority_id: ahj:usa-ma:plumbers-gasfitters-board
    state_code_name: "248 CMR 4.00 through 8.00, Massachusetts Fuel Gas Code"
    base_model_code: "NFPA 54 and NFPA 58, as amended, plus Massachusetts provisions"
    base_model_edition: null
    adoption_date: null
    effective_date: "2021-04-30"
    operative_date: "2021-04-30"
    mandatory_date: "2021-04-30"
    transition_rule_id: null
    source_ids:
      - src:usa-ma:248-cmr-4
      - src:usa-ma:248-cmr-5
      - src:usa-ma:248-cmr-8
    confidence: 0.72
```

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

The highest-value verified date rule is the 10th Edition State Building Code transition. The 10th Edition became effective on 2024-10-11. Massachusetts then allowed a concurrency period in which projects could use the 9th or 10th Edition. That concurrency period was extended through 2025-06-30. Projects using the 9th Edition needed building permit applications filed on or before 2025-06-30; on and after 2025-07-01, the 10th Edition is the mandatory building-code edition.

The electrical-code transition needs a specific caveat. The Mass.gov electrical-code page identifies an unofficial 2026 Massachusetts Electrical Code labeled effective 2026-04-24, while the code text excerpt states the code applies to installations for permits granted after 2026-02-28. Both dates are recorded separately and flagged for legal review rather than merged.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-ma:780-cmr-10th-concurrency | 780 CMR 10th Edition State Building Code | concurrency_then_mandatory | 2024-10-11 through 2025-06-30; mandatory 2025-07-01 | Building permit application filing date; projects using 9th Edition needed applications filed by 2025-06-30. | yes, through 2025-06-30 | src:usa-ma:780-cmr-10th-handbook; src:usa-ma:10th-edition-concurrency-extension; src:usa-ma:780-cmr-residential-ch1 | verified_core |
| date-rule:usa-ma:527-cmr-12-2026-permit | 527 CMR 12.00 Massachusetts Electrical Code | permit_trigger | Applies to permits granted after 2026-02-28; code source also labeled effective 2026-04-24 | Electrical installation permit grant date | unresolved | src:usa-ma:527-cmr-12-2026; src:usa-ma:electrical-code-page | partially_verified |
| date-rule:usa-ma:527-cmr-1-2021-effective | 527 CMR 1.00 Massachusetts Comprehensive Fire Safety Code | effective_date | 2023-05-12 | State fire-code effective date per current Mass.gov code page/PDF label | unresolved | src:usa-ma:fire-code-page; src:usa-ma:527-cmr-1-2021; src:usa-ma:fire-code-advisory-2022 | partially_verified |
| date-rule:usa-ma:225-cmr-municipal-energy | Stretch and Specialized Opt-in energy codes | municipal_effective_date | municipality-specific | Local adoption of Stretch or Specialized Opt-in Code; state registry lists municipal status and effective dates. | varies_by_municipality | src:usa-ma:municipal-energy-registry; src:usa-ma:225-cmr-22; src:usa-ma:225-cmr-23 | partially_verified |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Fire - operational / prevention code | 2024 Edition of 527 CMR 1.00 / future 2027 cycle | null | null | null | null | null | watch | src:usa-ma:submit-fire-code-changes | BFPR states it is no longer accepting Code Change Proposals for the 2024 Edition and expects to accept proposals for the 2027 Edition later in 2026. |
| Building | future 11th Edition / next I-Code cycle | null | null | null | null | null | watch | src:usa-ma:bbrs-page; src:usa-ma:780-cmr-10th-handbook | No official future adoption date identified in this pass. |
| Energy | municipal Stretch/Specialized adoption updates | null | null | null | null | null | active_registry_watch | src:usa-ma:municipal-energy-registry | Registry should be monitored because individual municipal effective dates change over time. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-ma:commonwealth-owned-buildings | Building | Buildings owned in whole or in part by the Commonwealth or its agencies/authorities | ownership by Commonwealth | State inspector enforces the State Building Code for such buildings and has local-inspector powers. | src:usa-ma:mgl-c143-s3a | verified |
| applicability-rule:usa-ma:historic-exterior-features | Building | Historic district, regional historic district, or architecturally controlled district exterior architectural features | conflict between State Building Code and district statute/ordinance/bylaw | The historic/architectural district exterior-feature regulation prevails in the event of conflict. | src:usa-ma:mgl-c143-s3a | verified |
| applicability-rule:usa-ma:energy-municipal-option | Energy | Projects in municipalities that have adopted Stretch or Specialized Opt-in Code | municipal adoption and effective date | Energy requirements may be Base, Stretch, or Specialized depending on municipality; consult state municipal registry. | src:usa-ma:energy-codes; src:usa-ma:municipal-energy-registry | partially_verified |
| applicability-rule:usa-ma:aab-accessibility | Accessibility | Public buildings and covered private/multiple-dwelling accessibility scopes | scope under M.G.L. c. 22, § 13A and 521 CMR | Accessibility rules are a specialized code and are enforced alongside State Building Code administration. | src:usa-ma:mgl-c22-s13a; src:usa-ma:521-cmr-physical-accessibility | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** Massachusetts publishes state-specific amendments and adopted/integrated code chapters in the Code of Massachusetts Regulations. The core building-code amendments are codified in 780 CMR, with Mass.gov providing unofficial convenience copies and the Massachusetts Register/official CMR as the official publication path. Specialized code amendments appear in other CMR titles, including 527 CMR, 248 CMR, 225 CMR, 521 CMR, and 524 CMR.

**Where amendments are published:** Official CMR / Massachusetts Register, with Mass.gov convenience pages and PDFs used for this draft.

**Amendment parsing status:** partial_source_registry_only

### 5.2 State Amendment Sources

| Source ID | Code Family | Publication Path | Amendment Model | Status | Notes |
| --- | --- | --- | --- | --- | --- |
| src:usa-ma:780-cmr-10th-toc | Building / Residential / Existing / Mechanical / Energy construction | 780 CMR 10th Edition table of contents and chapter PDFs | Massachusetts amendments to 2021 I-Codes | partially_parsed | Mass.gov page indicates unofficial copy; official version is Massachusetts Register. |
| src:usa-ma:527-cmr-12-2026 | Electrical | 527 CMR 12.00 PDF | Massachusetts amendments to NFPA 70, 2026 | partially_parsed | Date caveat retained. |
| src:usa-ma:527-cmr-1-2021 | Fire | 527 CMR 1.00 PDF | Massachusetts amendments to NFPA 1, 2021 | partially_parsed | Date caveat retained. |
| src:usa-ma:248-cmr-10 | Plumbing | 248 CMR 10.00 | Uniform State Plumbing Code | not_parsed | Regulation source identified; technical amendments not extracted. |
| src:usa-ma:248-cmr-4 | Fuel Gas | 248 CMR 4.00 through 8.00 | Massachusetts Fuel Gas Code | not_parsed | Regulation source identified; technical amendments not extracted. |
| src:usa-ma:225-cmr-22 | Energy | 225 CMR 22.00 | Residential Stretch/Specialized energy-code amendments | not_parsed | Needs municipality-level linkage. |
| src:usa-ma:225-cmr-23 | Energy | 225 CMR 23.00 | Commercial/multifamily Stretch/Specialized energy-code amendments | not_parsed | Needs municipality-level linkage. |
| src:usa-ma:521-cmr-physical-accessibility | Accessibility | 521 CMR | Architectural Access Board accessibility regulations | not_parsed | Current scope and effective dates need deeper extraction. |
| src:usa-ma:524-cmr-elevator-list | Elevator / Conveyance | 524 CMR | Elevator regulations | not_parsed | Chapter-by-chapter model references need extraction. |

### 5.3 High-Impact State Amendments

No technical state amendments were extracted into normalized amendment records in this pass. High-impact amendment review should prioritize:

1. 780 CMR Chapter 1 administration, including appeals, permit, and concurrency language.
2. 780 CMR Chapter 34 existing-building scope and IEBC modifications.
3. 780 CMR/225 CMR energy provisions, especially Base/Stretch/Specialized municipality routing.
4. 527 CMR 12.00 Rule 11 transition language and Massachusetts NEC amendments.
5. 527 CMR 1.00 Massachusetts amendments to NFPA 1.

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-ma"
  model: "municipal_building_official_enforcement_of_statewide_code"
  enforcing_entities:
    - "city_or_town_building_commissioner_or_inspector_of_buildings"
    - "local_inspectors"
    - "state_inspector_for_commonwealth_owned_buildings_and_supervisory_review"
    - "local_inspector_of_wires_for_electrical_work"
    - "local_plumbing_or_gas_fitting_inspectors_for_permitted_plumbing_or_gas_work"
  required_officials:
    - "inspector_of_buildings_or_building_commissioner"
    - "local_inspectors_as_needed"
  state_reserved_activities:
    - "state inspector enforcement for Commonwealth-owned buildings"
    - "state inspector review of local orders or decisions"
    - "BBRS certification of building officials"
    - "state elevator inspection system through commissioner and Board of Elevator Regulations"
  source_ids:
    - "src:usa-ma:mgl-c143-s3"
    - "src:usa-ma:mgl-c143-s3a"
    - "src:usa-ma:mgl-c143-s3l"
    - "src:usa-ma:mgl-c142-s11"
    - "src:usa-ma:mgl-c143-s62"
  verification_status: "partially_verified"
  confidence: 0.82
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-ma"
  model: "statewide_base_code_with_limited_supported_local_variation"
  applies_to_code_families:
    - "Building"
    - "Residential"
    - "Existing Building / Rehabilitation"
    - "Energy"
  approval_required: "unresolved_for_general_local_amendments"
  approving_authority_id: null
  filing_required: "unresolved_for_general_local_amendments"
  registry_exists: "yes_for_energy_code_status_by_municipality"
  registry_source_ids:
    - "src:usa-ma:municipal-energy-registry"
  legal_basis_source_ids:
    - "src:usa-ma:mgl-c143-s3a"
    - "src:usa-ma:energy-codes"
    - "src:usa-ma:225-cmr-22"
    - "src:usa-ma:225-cmr-23"
  verification_status: "partially_verified"
  confidence: 0.62
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Massachusetts local officials administer and enforce a statewide State Building Code. That local AHJ role should not be interpreted as local building-code adoption or amendment authority. The most clearly supported local variation in this pass is energy-code status: municipalities may be Base Code, Stretch Code, or Specialized Opt-in Code communities, and the state maintains a municipal energy-code adoption registry. M.G.L. c. 143, § 3A also preserves certain historic/architectural district exterior-feature regulations in case of conflict.

### 6.4 Known Local Amendment Registries

| Registry ID | Scope | Registry Source ID | Status | Notes |
| --- | --- | --- | --- | --- |
| registry:usa-ma:municipal-energy-code-status | Municipal Base/Stretch/Specialized energy-code status and effective dates | src:usa-ma:municipal-energy-registry | identified_not_parsed | Needs extraction before address-level energy-code routing. |

### 6.5 Municipality-Specific Known Amendments

No municipality-specific local amendments were parsed. Energy-code routing should not be inferred without consulting the municipal registry for the project location.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: partially_started

Jurisdiction stack:

```text
Address
  -> State: Massachusetts
  -> County: informational / limited code-administration role unless shared services are used
  -> Municipality: city or town
  -> Municipal building department: building commissioner / inspector of buildings
  -> Local inspectors: building, plumbing/gas, wires as applicable
  -> Fire department / fire prevention authority for operational fire-code matters
  -> Specialized state boards and state inspectors where applicable
  -> 780 CMR State Building Code adoption record
  -> 527 CMR / 248 CMR / 521 CMR / 524 CMR specialized-code adoption records
  -> Municipal energy-code status: Base, Stretch, or Specialized Opt-in
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | U.S. Census TIGER/Line or MassGIS state boundary | none | statewide | annual or state-managed | pending |
| County | U.S. Census TIGER/Line or MassGIS county boundaries | none | statewide | annual or state-managed | pending |
| Municipality | MassGIS municipal boundaries | none | statewide | state-managed | pending |
| Fire District | not selected | none | unknown | unknown | pending |
| Energy Code Municipality Status | Massachusetts municipal energy-code adoption registry | src:usa-ma:municipal-energy-registry | all 351 municipalities | state-managed / updated as adoptions change | identified_not_parsed |
| Special District | not selected | none | unknown | unknown | pending |

### 7.3 AHJ Contact Data

No AHJ contact data was populated. Minimum next pass should add municipal building department contacts and municipal energy-code status for a representative sample before treating the resolver as usable.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Publisher / Authority | URL / Citation | Source Type | Key Fields Supported | Status |
| --- | --- | --- | --- | --- | --- | --- |
| src:usa-ma:mgl-c143-s93 | M.G.L. c. 143, § 93, State board of building regulations and standards; establishment | Massachusetts Legislature | https://malegislature.gov/Laws/GeneralLaws/PartI/TitleXX/Chapter143/Section93 | statute | BBRS establishment; adoption/administering role | parsed_core |
| src:usa-ma:mgl-c143-s94 | M.G.L. c. 143, § 94, Powers and duties | Massachusetts Legislature | https://malegislature.gov/laws/generallaws/parti/titlexx/chapter143/section94 | statute | BBRS powers; State Building Code scope; energy update duties | parsed_core |
| src:usa-ma:mgl-c143-s3 | M.G.L. c. 143, § 3, Local building officials | Massachusetts Legislature | https://malegislature.gov/Laws/GeneralLaws/PartI/TitleXX/Chapter143/Section3 | statute | local enforcement officials; certification | parsed_core |
| src:usa-ma:mgl-c143-s3a | M.G.L. c. 143, § 3A, Enforcement of State Building Code | Massachusetts Legislature | https://malegislature.gov/Laws/GeneralLaws/PartI/TitleXX/Chapter143/Section3A | statute | local enforcement; State Building Code as code for all city/town buildings; Commonwealth-owned buildings; historic district exception | parsed_core |
| src:usa-ma:mgl-c143-s3l | M.G.L. c. 143, § 3L, Electrical wiring and fixtures | Massachusetts Legislature | https://malegislature.gov/Laws/GeneralLaws/PartI/TitleXX/Chapter143/Section3L | statute | BFPR electrical rulemaking and enforcement | parsed_core |
| src:usa-ma:mgl-c22d-s4 | M.G.L. c. 22D, § 4, Board of Fire Prevention Regulations | Massachusetts Legislature | https://malegislature.gov/Laws/GeneralLaws/PartI/TitleII/Chapter22d/Section4 | statute | BFPR establishment; comprehensive fire safety code duty | parsed_core |
| src:usa-ma:mgl-c142-s11 | M.G.L. c. 142, § 11, Plumbing/gas inspectors | Massachusetts Legislature | https://malegislature.gov/Laws/GeneralLaws/PartI/TitleXX/Chapter142/Section11 | statute | local plumbing/gas inspection model | parsed_core |
| src:usa-ma:mgl-c22-s13a | M.G.L. c. 22, § 13A, Architectural Access Board | Massachusetts Legislature | https://malegislature.gov/Laws/GeneralLaws/PartI/TitleII/Chapter22/Section13A | statute | AAB authority; specialized-code status | parsed_core |
| src:usa-ma:mgl-c143-s62 | M.G.L. c. 143, § 62, Elevator inspection system | Massachusetts Legislature | https://malegislature.gov/Laws/GeneralLaws/PartI/TitleXX/Chapter143/Section62 | statute | elevator inspection and regulation authority | parsed_core |
| src:usa-ma:780-cmr-10th-handbook | Tenth edition of the MA State Building Code 780 | Mass.gov / Office of Public Safety and Inspections | https://www.mass.gov/handbook/tenth-edition-of-the-ma-state-building-code-780 | official agency page / handbook | 10th Edition effective date; concurrency; 2021 I-Code basis | parsed_core |
| src:usa-ma:10th-edition-concurrency-extension | Concurrency Period for 9th and 10th Edition of the State Building Code Extended Until June 30, 2025 | Mass.gov / BBRS | https://www.mass.gov/news/concurrency-period-for-9th-and-10th-edition-of-the-state-building-code-extended-until-june-30-2025 | official agency notice | permit filing date through 2025-06-30; mandatory 10th Edition after concurrency | parsed_core |
| src:usa-ma:780-cmr-10th-toc | 780 CMR Tenth Edition: Table of Contents | Mass.gov / Office of Public Safety and Inspections | https://www.mass.gov/regulations/780-CMR-tenth-edition-table-of-contents | regulation convenience page | 780 CMR 10th Edition structure | parsed_core |
| src:usa-ma:780-cmr-10th-ch1 | 10th Edition Chapter 1 Scope and Administration of Amendments | Mass.gov / BBRS | https://www.mass.gov/doc/10th-edition-chapter-1-scope-and-administration-of-amendments/download | regulation PDF convenience copy | IBC adoption; IMC reference; administrative scope | partially_parsed |
| src:usa-ma:780-cmr-ch28 | 10th Edition Chapter 28 Mechanical Systems | Mass.gov / BBRS | https://www.mass.gov/doc/10th-edition-chapter-28-mechanical-systems/download | regulation PDF convenience copy | mechanical provisions | source_identified |
| src:usa-ma:780-cmr-ch34 | 780 CMR Tenth Edition Chapter 34 Existing Building Code Amendments | Mass.gov / BBRS | https://www.mass.gov/regulations/780-CMR-tenth-edition-chapter-34-existing-building-code-amendments | regulation convenience page | IEBC 2021 existing building scope | parsed_core |
| src:usa-ma:780-cmr-residential-ch1 | 10th Edition Residential Chapter 1 Scope and Application | Mass.gov / BBRS | https://www.mass.gov/doc/10th-edition-residential-chapter-1-scope-and-application/download | regulation PDF convenience copy | residential transition; residential code scope | parsed_core |
| src:usa-ma:energy-codes | Massachusetts Building Energy Codes | Mass.gov / Department of Energy Resources | https://www.mass.gov/info-details/massachusetts-building-energy-codes | official agency page | Base/Stretch/Specialized energy-code pathways; IECC 2021 basis | parsed_core |
| src:usa-ma:225-cmr-22 | 225 CMR 22.00, Massachusetts Stretch and Specialized Code for Low-Rise Residential | Mass.gov / Department of Energy Resources | https://www.mass.gov/doc/225-cmr-22-massachusetts-stretch-code-and-specialized-code-for-low-rise-residential-2025-residential-low-rise-amendments-to-iecc2021-and-irc-2021-chapter-11-energy-efficiency/download | regulation PDF convenience copy | residential Stretch/Specialized code | source_identified |
| src:usa-ma:225-cmr-23 | 225 CMR 23.00, Massachusetts Stretch and Specialized Code for Commercial, Multi-Family, and Other Construction | Mass.gov / Department of Energy Resources | https://www.mass.gov/regulations/225-CMR-2300-massachusetts-stretch-code-and-specialized-code-for-commercial-multi-family-and-all-other-construction-2025-amendments-to-iecc2021-and-ashrae-standards-901-2019 | regulation convenience page | commercial/multifamily Stretch/Specialized code | source_identified |
| src:usa-ma:municipal-energy-registry | Massachusetts Building Energy Code Adoption by Municipality | Mass.gov / Department of Energy Resources | https://www.mass.gov/info-details/massachusetts-building-energy-code-adoption-by-municipality | official agency registry | municipal Base/Stretch/Specialized status and effective dates | identified_not_parsed |
| src:usa-ma:fire-code-page | Massachusetts Fire Code | Mass.gov / Department of Fire Services | https://www.mass.gov/info-details/massachusetts-fire-code | official agency page | current fire code; NFPA 1 2021 basis | parsed_core |
| src:usa-ma:527-cmr-1-2021 | Massachusetts 527 CMR 1.00, 2021 Edition, effective May 12, 2023 | Mass.gov / Department of Fire Services | https://www.mass.gov/doc/massachusetts-527-cmr-100-2021-edition-effective-may-12-2023/download | regulation PDF convenience copy | current 527 CMR 1.00, NFPA 1 2021 with amendments | parsed_core |
| src:usa-ma:fire-code-advisory-2022 | Advisory: Revisions to 527 CMR 1.00, effective December 9, 2022 | Mass.gov / Department of Fire Services | https://www.mass.gov/doc/advisory-2021-amendment-package-memo/download | official agency advisory | historical revision date caveat | source_identified |
| src:usa-ma:submit-fire-code-changes | Submit Proposed Changes to the State Fire Code | Mass.gov / Department of Fire Services | https://www.mass.gov/how-to/submit-proposed-changes-to-the-state-fire-code | official agency page | future fire-code cycle watch | parsed_core |
| src:usa-ma:electrical-code-page | Massachusetts Electrical Code | Mass.gov / Department of Fire Services | https://www.mass.gov/info-details/massachusetts-electrical-code | official agency page | current 527 CMR 12.00; NFPA 70 2026 basis | parsed_core |
| src:usa-ma:527-cmr-12-2026 | 527 CMR 12.00 Massachusetts Electrical Code, effective 4/24/26 | Mass.gov / Board of Fire Prevention Regulations | https://www.mass.gov/doc/527-cmr-1200-massachusetts-electrical-code-effective-42426/download | regulation PDF convenience copy | 2026 NEC adoption; permit trigger | parsed_core |
| src:usa-ma:248-cmr-10 | 248 CMR 10.00, Uniform State Plumbing Code | Mass.gov / Board of State Examiners of Plumbers and Gas Fitters | https://www.mass.gov/regulations/248-CMR-1000-uniform-state-plumbing-code | regulation convenience page | current plumbing code | parsed_core |
| src:usa-ma:plumbing-2023-faq | Board of State Examiners of Plumbers and Gas Fitters FAQ for new regulations | Mass.gov / Board of State Examiners of Plumbers and Gas Fitters | https://www.mass.gov/lists/board-of-state-examiners-of-plumbers-and-gas-fitters-frequently-asked-questions-for-new-regulations | official agency FAQ | 2023-12-08 changes to 248 CMR 10.00 | parsed_core |
| src:usa-ma:248-cmr-4 | 248 CMR 4.00, Massachusetts Fuel Gas Code | Mass.gov / Board of State Examiners of Plumbers and Gas Fitters | https://www.mass.gov/regulations/248-CMR-400-massachusetts-fuel-gas-code | regulation convenience page | fuel gas code scope and effective date | parsed_core |
| src:usa-ma:248-cmr-5 | 248 CMR 5.00, Amendments to NFPA 54 | Mass.gov / Board of State Examiners of Plumbers and Gas Fitters | https://www.mass.gov/regulations/248-CMR-500-amendments-to-nfpa-54 | regulation convenience page | NFPA 54 amendments | source_identified |
| src:usa-ma:248-cmr-8 | 248 CMR 8.00, Amendments to NFPA 58 | Mass.gov / Board of State Examiners of Plumbers and Gas Fitters | https://www.mass.gov/regulations/248-CMR-800-amendments-to-nfpa-58 | regulation convenience page | NFPA 58 amendments | source_identified |
| src:usa-ma:plumbers-board-page | Board of State Examiners of Plumbers and Gas Fitters | Mass.gov / Division of Occupational Licensure | https://www.mass.gov/orgs/board-of-state-examiners-of-plumbers-and-gas-fitters | official agency page | board role in plumbing and gas fitting | parsed_core |
| src:usa-ma:521-cmr-physical-accessibility | Physical accessibility requirements | Mass.gov / Architectural Access Board | https://www.mass.gov/info-details/physical-accessibility-requirements | official agency page | 521 CMR current accessibility regulations | source_identified |
| src:usa-ma:521-cmr-1996 | 521 CMR 1996 Edition | Mass.gov / Architectural Access Board | https://www.mass.gov/doc/521-cmr-1996-edition/download | regulation PDF convenience copy | 521 CMR edition history / effective-date context | source_identified |
| src:usa-ma:524-cmr-elevator-list | Elevator laws and regulations, effective June 1, 2018 | Mass.gov / Board of Elevator Regulations | https://www.mass.gov/lists/elevator-laws-and-regulations-effective-june-1-2018 | official agency list | 524 CMR current regulation set | source_identified |
| src:usa-ma:524-cmr-35 | 524 CMR 35.00, Safety Code for Elevators and Escalators A17.1-2013 | Mass.gov / Board of Elevator Regulations | https://www.mass.gov/doc/524-cmr-chapter-35-safety-code-for-elevators-and-escalators-a171-2013-and-the-massachusetts-modifications-of-that-code/download | regulation PDF convenience copy | ASME A17.1-2013 incorporation | source_identified |
| src:usa-ma:bbrs-page | Board of Building Regulations and Standards | Mass.gov / Office of Public Safety and Inspections | https://www.mass.gov/orgs/board-of-building-regulations-and-standards | official agency page | agency identity and monitoring target | source_identified |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-ma:780-cmr-10th-toc | unofficial_convenience_copy | Mass.gov regulation pages state that posted versions are unofficial convenience copies and that the official version is published by the Massachusetts Register. | use_for_draft; verify_against_register_for_verified_status |
| src:usa-ma:780-cmr-10th-ch1 | pdf_convenience_copy | PDF was not fully parsed chapter-by-chapter in this pass. | use_for_core_scope_only |
| src:usa-ma:780-cmr-residential-ch1 | pdf_convenience_copy | Residential PDF source supports transition language but was not fully parsed. | use_for_core_transition_only |
| src:usa-ma:527-cmr-12-2026 | date_sequence_caveat | Source label gives effective 2026-04-24, while Rule 11 uses permits granted after 2026-02-28. | legal_review_required_before_verified_status |
| src:usa-ma:527-cmr-1-2021 | date_history_caveat | Current PDF label uses 2023-05-12, while advisory materials reference 2022-12-09 revisions. | reconcile_register_history_before_verified_status |
| src:usa-ma:municipal-energy-registry | dynamic_registry | Municipal energy status can change by local adoption. | monitor_and_refresh_before_address_level_use |
| src:usa-ma:521-cmr-physical-accessibility | effective_date_gap | Current 521 CMR applicability and edition dates were not normalized. | keep_date_fields_unresolved |
| src:usa-ma:524-cmr-elevator-list | chapter_level_gap | 524 CMR chapter-level model references and transition dates were not fully parsed. | source_identified_only |

### 8.3 Supplemental Sources

None used in this pass. Third-party summaries were intentionally excluded from the source registry.

### 8.4 Source Extraction Metadata

| Extraction ID | Source IDs | Extracted Fields | Extracted On | Method | Notes |
| --- | --- | --- | --- | --- | --- |
| extraction:usa-ma:2026-06-25-authority | src:usa-ma:mgl-c143-s93; src:usa-ma:mgl-c143-s94; src:usa-ma:mgl-c143-s3; src:usa-ma:mgl-c143-s3a | BBRS authority; local enforcement | 2026-06-25 | official_html_review | Core authority fields verified. |
| extraction:usa-ma:2026-06-25-fire-electrical | src:usa-ma:mgl-c22d-s4; src:usa-ma:mgl-c143-s3l; src:usa-ma:fire-code-page; src:usa-ma:527-cmr-1-2021; src:usa-ma:electrical-code-page; src:usa-ma:527-cmr-12-2026 | fire and electrical authority/adoption | 2026-06-25 | official_html_and_pdf_snippet_review | Date caveats retained. |
| extraction:usa-ma:2026-06-25-building-code | src:usa-ma:780-cmr-10th-handbook; src:usa-ma:10th-edition-concurrency-extension; src:usa-ma:780-cmr-10th-toc; src:usa-ma:780-cmr-residential-ch1 | 10th Edition status and transition | 2026-06-25 | official_agency_page_review | Core transition verified. |
| extraction:usa-ma:2026-06-25-specialized | src:usa-ma:248-cmr-10; src:usa-ma:248-cmr-4; src:usa-ma:mgl-c22-s13a; src:usa-ma:mgl-c143-s62 | specialized authorities and code families | 2026-06-25 | official_source_identification | Technical provisions not parsed. |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 1.00 | none | Core authority and adoption fields are source-backed, but specialized code details remain partial. |
| report | risk.overall_confidence | 0.66 | verified | 1.00 | none | Confidence reflects verified core plus unresolved specialized details. |
| ahj:usa-ma:bbrs | authority_name | State Board of Building Regulations and Standards | verified | 0.95 | src:usa-ma:mgl-c143-s93 | Statute establishes board. |
| ahj:usa-ma:bbrs | role | adopt and administer State Building Code | verified | 0.95 | src:usa-ma:mgl-c143-s93; src:usa-ma:mgl-c143-s94 | Statutes support role and scope. |
| adoption:usa-ma:780-cmr-10th-building | effective_date | 2024-10-11 | verified_core | 0.90 | src:usa-ma:780-cmr-10th-handbook | Agency page supports date. |
| adoption:usa-ma:780-cmr-10th-building | mandatory_date | 2025-07-01 | verified_core | 0.88 | src:usa-ma:10th-edition-concurrency-extension; src:usa-ma:780-cmr-10th-handbook | Derived from concurrency ending 2025-06-30. |
| local-enforcement:usa-ma | model | municipal_building_official_enforcement_of_statewide_code | verified_core | 0.90 | src:usa-ma:mgl-c143-s3; src:usa-ma:mgl-c143-s3a | City/town officials enforce State Building Code. |
| local-amendment-rule:usa-ma | model | statewide_base_code_with_limited_supported_local_variation | partially_verified | 0.62 | src:usa-ma:mgl-c143-s3a; src:usa-ma:energy-codes | General local amendment authority not identified; energy local options verified at registry level only. |
| adoption:usa-ma:527-cmr-12-2026 | base_model_edition | NFPA 70, 2026 | verified_core | 0.85 | src:usa-ma:electrical-code-page; src:usa-ma:527-cmr-12-2026 | Date caveat retained. |
| adoption:usa-ma:527-cmr-1-2021 | base_model_edition | NFPA 1, 2021 | verified_core | 0.85 | src:usa-ma:fire-code-page; src:usa-ma:527-cmr-1-2021 | Effective date history caveat retained. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Every `src:usa-ma:*` reference in body appears in section 8. |
| All authority IDs resolve | pass | Authority IDs used in sections 2 and 3 are defined or directly described. |
| All current code families have adoption records | fail | Matrix rows are present for all code families, but normalized adoption records are only included for supported high-value families. |
| Building and operational fire code are separated | pass | Construction fire references and operational/prevention fire code are distinct rows. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Date fields are separate; unknown dates are null. |
| Effective dates are valid ISO dates | pass | Known dates use YYYY-MM-DD. |
| No impossible date sequences | pass_with_caveat | Electrical record intentionally flags source date sequence conflict. |
| Transition rules have explicit trigger conditions | pass | Highest-value transition rules include permit-application or permit-grant triggers. |
| Permit-date logic is captured where applicable | pass | 780 CMR and 527 CMR 12.00 permit triggers are captured. |
| Local enforcement model classified | pass | Municipal enforcement of statewide code is classified. |
| Local amendment rule classified | pass_with_caveat | Classified conservatively; general local amendment authority still needs legal review. |
| AHJ confirmation metadata present | fail | No AHJ contact data populated. |
| Official-source caveats captured | pass | Mass.gov unofficial-copy and date caveats included. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-ma:001 | high | official regulation text | Mass.gov regulation pages are convenience copies; official Massachusetts Register/CMR text was not independently checked. | Verify 780 CMR, 527 CMR, 248 CMR, 225 CMR, 521 CMR, and 524 CMR against the official CMR/Massachusetts Register publication path. | null | null | open |
| issue:usa-ma:002 | high | electrical transition dates | 527 CMR 12.00 is labeled effective 2026-04-24, but Rule 11 applies to permits granted after 2026-02-28. | Review filing/register history and BFPR adoption materials to reconcile effective and permit-trigger dates. | null | null | open |
| issue:usa-ma:003 | medium | fire-code date history | Current fire-code source labels 2021 Edition effective 2023-05-12, but advisory materials reference revisions effective 2022-12-09. | Review 527 CMR 1.00 register filings and amendment history. | null | null | open |
| issue:usa-ma:004 | medium | municipal energy routing | State registry of municipal Base/Stretch/Specialized energy-code status was identified but not parsed. | Extract municipal status and effective dates for all 351 municipalities. | null | null | open |
| issue:usa-ma:005 | medium | specialized code technical parsing | 248 CMR, 521 CMR, and 524 CMR technical provisions and model-code references were only partially identified. | Parse specialized code chapters and normalize adoption records where appropriate. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-ma:780-cmr | src:usa-ma:780-cmr-10th-handbook | html_diff | monthly | new edition, revised concurrency note, or new chapter links | 2026-06-25 | active |
| watch:usa-ma:bbrs | src:usa-ma:bbrs-page | html_diff | monthly | BBRS meeting notices, code update notices, or new 780 CMR material | 2026-06-25 | active |
| watch:usa-ma:fire-code | src:usa-ma:fire-code-page | html_diff | monthly | new 527 CMR 1.00 edition or errata | 2026-06-25 | active |
| watch:usa-ma:electrical-code | src:usa-ma:electrical-code-page | html_diff | monthly | new 527 CMR 12.00 edition, errata, or interpretation | 2026-06-25 | active |
| watch:usa-ma:energy-municipal-registry | src:usa-ma:municipal-energy-registry | data_diff | monthly | municipal Base/Stretch/Specialized status or effective-date changes | 2026-06-25 | active |
| watch:usa-ma:plumbing-gas | src:usa-ma:plumbers-board-page | html_diff | quarterly | new 248 CMR amendments or board policy changes | 2026-06-25 | active |
| watch:usa-ma:aab | src:usa-ma:521-cmr-physical-accessibility | html_diff | quarterly | 521 CMR updates or AAB guidance changes | 2026-06-25 | active |
| watch:usa-ma:elevator | src:usa-ma:524-cmr-elevator-list | html_diff | quarterly | 524 CMR updates or model-code reference changes | 2026-06-25 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-25 | Replaced Massachusetts baseline draft with source-backed partially verified report. | report:usa-ma; ahj:usa-ma:bbrs; adoption:usa-ma:780-cmr-10th-building; local-enforcement:usa-ma | src:usa-ma:mgl-c143-s93; src:usa-ma:mgl-c143-s94; src:usa-ma:780-cmr-10th-handbook; src:usa-ma:10th-edition-concurrency-extension | ChatGPT | Core authority and adoption matrix populated; caveats retained. |
| 2026-06-25 | Added specialized authority rows and unresolved technical parsing issues. | ahj:usa-ma:bfpr; ahj:usa-ma:plumbers-gasfitters-board; ahj:usa-ma:aab; ahj:usa-ma:elevator-board | src:usa-ma:mgl-c22d-s4; src:usa-ma:mgl-c143-s3l; src:usa-ma:248-cmr-10; src:usa-ma:mgl-c22-s13a; src:usa-ma:mgl-c143-s62 | ChatGPT | Specialized code details remain partial. |
