---
state:
  state_id: "US-WA"
  name: "Washington"
  abbreviation: "WA"
report:
  report_id: "state-report:usa-wa"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-26"
  last_verified: "2026-06-26"
  reviewed_by: null
risk:
  overall_confidence: 0.87 # 0.00 - 1.00
  risk_flags:
    - "2024_code_cycle_future_effective_date_monitored"
    - "building_permit_date_trigger_not_fully_normalized"
    - "local_amendment_registry_not_parsed"
    - "ahj_contact_data_not_populated"
  open_questions_count: 5
---

# State Building Code Authority Report: Washington

## 1. Executive Summary

- **Authority model:** Washington uses a statewide State Building Code administered through the Washington State Building Code Council (SBCC). Electrical and elevator/conveyance codes are administered through the Washington State Department of Labor & Industries (L&I) under separate statutory and WAC chapters.
- **Statewide code status:** The statewide building-code set is on the 2021 code cycle. Major SBCC-administered chapters became effective statewide on 2024-03-15. The current electrical code adopts the 2023 NEC until the 2026 NEC takes effect on 2026-12-31. Elevator/conveyance rules currently include ASME A17.1-2019/CSA B44-19 for new elevators and escalators.
- **Local enforcement model:** Counties and cities enforce the State Building Code. Electrical inspection may be performed by L&I or by a city with electrical inspection jurisdiction. Elevator/conveyance enforcement is a state L&I program, with local building departments still relevant for associated building permits.
- **Local amendment posture:** Counties and cities may amend State Building Code chapters only within RCW 19.27.060 limits. Residential amendments affecting single-family or multifamily buildings generally require SBCC approval. Local enforcement authority is separate from local amendment authority.
- **Known transition periods or pending changes:** The 2024 Washington State Building Codes are in the adoption cycle; SBCC states final adoption was delayed to 2026-08-21 and the 2024-code effective date will be 2027-05-03. L&I states the 2026 NEC will replace the current 2023 NEC on 2026-12-31.
- **Production readiness:** usable_with_caution

### Key Findings

```yaml
---
key_findings:
- topic: '**State adopting authority**'
  finding: The Washington State Building Code Council maintains the State Building
    Code under chapter 19.27 RCW and WAC Title 51.
  confidence: 0.95
  source_ids:
  - src:usa-wa:statute:rcw-19-27
  - src:usa-wa:regulation:wac-title-51
- topic: '**Primary building code edition**'
  finding: WAC 51-50 adopts and amends the 2021 International Building Code, including
    Appendix E; the chapter became effective in all counties and cities on 2024-03-15.
  confidence: 0.97
  source_ids:
  - src:usa-wa:regulation:wac-51-50
  - src:usa-wa:agency:sbcc-building-amendments
- topic: '**Electrical code authority**'
  finding: Electrical rules are administered under chapter 19.28 RCW and chapter 296-46B
    WAC; WAC 296-46B-010 currently adopts the 2023 NEC and sets a future replacement
    by the 2026 NEC on 2026-12-31.
  confidence: 0.94
  source_ids:
  - src:usa-wa:statute:rcw-19-28
  - src:usa-wa:regulation:wac-296-46b-010
  - src:usa-wa:agency:lni-electrical-laws
- topic: '**Fire code authority**'
  finding: WAC 51-54A adopts and amends the 2021 International Fire Code. Counties
    and cities enforce the State Building Code under RCW 19.27.050, while fire-code-official
    functions remain local/AHJ-specific.
  confidence: 0.92
  source_ids:
  - src:usa-wa:regulation:wac-51-54a
  - src:usa-wa:statute:rcw-19-27-050
- topic: '**Local enforcement**'
  finding: The State Building Code is enforced by counties and cities; a county or
    city without a building department must contract with an approved enforcement
    entity.
  confidence: 0.96
  source_ids:
  - src:usa-wa:statute:rcw-19-27-050
- topic: '**Local amendments**'
  finding: Counties and cities may amend the State Building Code within statutory
    limits, but local amendments may not reduce minimum state standards and residential
    amendments generally require SBCC approval.
  confidence: 0.92
  source_ids:
  - src:usa-wa:statute:rcw-19-27-060
- topic: '**Effective / operative date rule**'
  finding: Major 2021 SBCC-administered code chapters state 2024-03-15 as the statewide
    effective date. Building permit-date trigger details are still AHJ-sensitive;
    electrical standards apply by electrical permit issue date except for stated exceptions.
  confidence: 0.86
  source_ids:
  - src:usa-wa:regulation:wac-51-50
  - src:usa-wa:regulation:wac-51-51
  - src:usa-wa:regulation:wac-51-52
  - src:usa-wa:regulation:wac-51-54a
  - src:usa-wa:regulation:wac-51-56
  - src:usa-wa:regulation:wac-51-11c
  - src:usa-wa:regulation:wac-51-11r
  - src:usa-wa:regulation:wac-296-46b-010
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| **Authority ID** | ahj:usa-wa:state-building-code-council |
| **Authority name** | Washington State Building Code Council |
| **Authority type** | state_council |
| **Legal basis** | Chapter 19.27 RCW; WAC Title 51 |
| **Role** | Maintains statewide building-code adoptions and Washington amendments for building, residential, mechanical/fuel-gas, plumbing, energy, fire, existing-building, accessibility-linked building-code provisions, and related model-code insert pages |
| **Enforcement model** | locally_enforced for State Building Code; specialized state programs for electrical and elevator/conveyance |
| **Source IDs** | src:usa-wa:statute:rcw-19-27; src:usa-wa:regulation:wac-title-51; src:usa-wa:agency:sbcc-current-codes |
| **Verification status** | verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| **Building** | ahj:usa-wa:state-building-code-council | Washington State Building Code Council | Adopts and amends the state building code | WAC 51-50; RCW 19.27 | src:usa-wa:regulation:wac-51-50; src:usa-wa:statute:rcw-19-27 | verified |
| **Residential** | ahj:usa-wa:state-building-code-council | Washington State Building Code Council | Adopts and amends the state residential code | WAC 51-51; RCW 19.27 | src:usa-wa:regulation:wac-51-51; src:usa-wa:statute:rcw-19-27 | verified |
| **Existing Building / Rehabilitation** | ahj:usa-wa:state-building-code-council | Washington State Building Code Council | Includes the 2021 IEBC within the IBC adoption and SBCC building-code insert pages | WAC 51-50; SBCC building-code amendments | src:usa-wa:regulation:wac-51-50; src:usa-wa:agency:sbcc-building-amendments | verified |
| **Mechanical** | ahj:usa-wa:state-building-code-council | Washington State Building Code Council | Adopts and amends the 2021 IMC | WAC 51-52; RCW 19.27 | src:usa-wa:regulation:wac-51-52; src:usa-wa:agency:sbcc-mechanical-amendments | verified |
| **Plumbing** | ahj:usa-wa:state-building-code-council | Washington State Building Code Council | Adopts and amends the 2021 UPC with stated exclusions | WAC 51-56; RCW 19.27 | src:usa-wa:regulation:wac-51-56; src:usa-wa:regulation:wac-title-51 | verified |
| **Fuel Gas** | ahj:usa-wa:state-building-code-council | Washington State Building Code Council | Adopts fuel-gas provisions through the mechanical code chapter, including the 2021 IFGC and NFPA fuel-gas standards | WAC 51-52 | src:usa-wa:regulation:wac-51-52; src:usa-wa:agency:sbcc-mechanical-amendments | verified |
| **Electrical** | ahj:usa-wa:lni-electrical | Washington State Department of Labor & Industries | Administers electrical laws and rules; WAC 296-46B adopts and amends NEC-based electrical standards | RCW 19.28; WAC 296-46B | src:usa-wa:statute:rcw-19-28; src:usa-wa:regulation:wac-296-46b; src:usa-wa:regulation:wac-296-46b-010; src:usa-wa:agency:lni-electrical-laws | verified |
| **Energy** | ahj:usa-wa:state-building-code-council | Washington State Building Code Council | Adopts the Washington State Energy Code, commercial and residential | WAC 51-11C; WAC 51-11R; chapter 19.27A RCW | src:usa-wa:regulation:wac-51-11c; src:usa-wa:regulation:wac-51-11r; src:usa-wa:agency:sbcc-energy-code | verified |
| **Fire - construction references** | ahj:usa-wa:state-building-code-council | Washington State Building Code Council | Adopts and amends the IFC within the State Building Code | WAC 51-54A; RCW 19.27 | src:usa-wa:regulation:wac-51-54a; src:usa-wa:agency:sbcc-fire-amendments | verified |
| **Fire - operational / prevention code** | ahj:usa-wa:local-fire-code-official | Local fire code officials / local governments | Administer operational fire-code functions locally under the state IFC framework | WAC 51-54A; RCW 19.27.050 | src:usa-wa:regulation:wac-51-54a; src:usa-wa:statute:rcw-19-27-050 | partially_verified |
| **Accessibility** | ahj:usa-wa:state-building-code-council | Washington State Building Code Council | Adopts accessibility-linked building-code provisions through the IBC adoption, including Appendix E and ICC A117.1-2017 in SBCC insert-page publication path | WAC 51-50; chapter 70.92 RCW references in IBC rule authority | src:usa-wa:regulation:wac-51-50; src:usa-wa:agency:sbcc-building-amendments | partially_verified |
| **Elevator / Conveyance** | ahj:usa-wa:lni-elevator | Washington State Department of Labor & Industries | Administers elevator/conveyance laws and rules | RCW 70.87; WAC 296-96 | src:usa-wa:agency:lni-elevator-laws; src:usa-wa:regulation:wac-296-96; src:usa-wa:regulation:wac-296-96-00650 | verified |

### 2.3 Authority Hierarchy Notes

Washington's general construction-code model separates statewide code adoption from local building-code enforcement. SBCC administers the State Building Code adoption and amendment framework. Counties and cities enforce the State Building Code locally, and local jurisdictions may make limited amendments under RCW 19.27.060.

Electrical and elevator/conveyance programs are not simply additional WAC Title 51 chapters. Electrical rules are administered by L&I under chapter 19.28 RCW and chapter 296-46B WAC, with inspection jurisdiction split between L&I and authorized cities. Elevator/conveyance rules are administered by L&I under chapter 70.87 RCW and chapter 296-96 WAC. For projects, local building permits, local fire review, electrical permits, and elevator permits may therefore involve different AHJs.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-wa:001 | ahj:usa-wa:state-building-code-council | delegates_enforcement_to | counties_and_cities_for_state_building_code | src:usa-wa:statute:rcw-19-27-050 | verified |
| edge:usa-wa:002 | ahj:usa-wa:state-building-code-council | limits_local_amendments_by | minimum_state_performance_standards_and_residential_approval_rule | src:usa-wa:statute:rcw-19-27-060 | verified |
| edge:usa-wa:003 | ahj:usa-wa:state-building-code-council | adopts_and_amends | WAC_Title_51_state_building_code_chapters | src:usa-wa:regulation:wac-title-51 | verified |
| edge:usa-wa:004 | ahj:usa-wa:lni-electrical | administers | NEC_based_electrical_safety_standards_and_inspections | src:usa-wa:statute:rcw-19-28; src:usa-wa:regulation:wac-296-46b-010 | verified |
| edge:usa-wa:005 | ahj:usa-wa:lni-elevator | administers | elevator_and_conveyance_safety_rules | src:usa-wa:agency:lni-elevator-laws; src:usa-wa:regulation:wac-296-96 | verified |
| edge:usa-wa:006 | ahj:usa-wa:local-fire-code-official | enforces_or_administers | IFC_operational_permits_and_local_fire_official_functions | src:usa-wa:regulation:wac-51-54a; src:usa-wa:statute:rcw-19-27-050 | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| **Building** | Washington State Building Code | IBC | 2021 | current | null | 2024-03-15 | 2024-03-15 | 2024-03-15 | WAC 51-50 states the 2021 IBC adoption is effective in all counties and cities on 2024-03-15. Project-level trigger remains AHJ-sensitive. | src:usa-wa:regulation:wac-51-50; src:usa-wa:agency:sbcc-building-amendments |
| **Residential** | Washington State Residential Code | IRC | 2021 | current | null | 2024-03-15 | 2024-03-15 | 2024-03-15 | WAC 51-51 states the 2021 IRC adoption is effective in all counties and cities on 2024-03-15; chapters 11 and 25-43 are not adopted. | src:usa-wa:regulation:wac-51-51; src:usa-wa:agency:sbcc-residential-amendments |
| **Existing Building / Rehabilitation** | Washington State Building Code | IEBC included in IBC adoption | 2021 | current | null | 2024-03-15 | 2024-03-15 | 2024-03-15 | The 2021 IEBC is included in the IBC adoption and amended in WAC 51-50-480000. | src:usa-wa:regulation:wac-51-50; src:usa-wa:agency:sbcc-building-amendments |
| **Mechanical** | Washington State Mechanical Code | IMC | 2021 | current | null | 2024-03-15 | 2024-03-15 | 2024-03-15 | WAC 51-52 states the 2021 IMC adoption is effective in all counties and cities on 2024-03-15. | src:usa-wa:regulation:wac-51-52; src:usa-wa:agency:sbcc-mechanical-amendments |
| **Plumbing** | Washington State Plumbing Code | UPC | 2021 | current | null | 2024-03-15 | 2024-03-15 | 2024-03-15 | WAC 51-56 states the 2021 UPC adoption is effective in all counties and cities on 2024-03-15, with stated chapter and subject exclusions. | src:usa-wa:regulation:wac-51-56 |
| **Fuel Gas** | Washington State Mechanical Code / Fuel Gas provisions | IFGC; NFPA 58; NFPA 54 | 2021 IFGC; 2020 NFPA 58; 2021 NFPA 54 | current | null | 2024-03-15 | 2024-03-15 | 2024-03-15 | WAC 51-52 regulates fuel-gas distribution piping and equipment through IFGC, with LP-gas standards identified as NFPA 58 and ANSI Z223.1/NFPA 54. | src:usa-wa:regulation:wac-51-52; src:usa-wa:agency:sbcc-mechanical-amendments |
| **Electrical** | Electrical Safety Standards, Administration, and Installation | NEC / NFPA 70 | 2023 | current | null | 2024-04-01 | 2024-04-01 | 2024-04-01 | WAC 296-46B-010 adopts the 2023 NEC; adopted standards apply to installations when electrical permit issue dates are on and after adoption dates, subject to stated exceptions. | src:usa-wa:regulation:wac-296-46b-010; src:usa-wa:agency:lni-electrical-laws |
| **Energy** | Washington State Energy Code - Commercial | IECC / Washington State Energy Code | 2021 | current | null | 2024-03-15 | 2024-03-15 | 2024-03-15 | WAC 51-11C adopts the 2021 Washington State Energy Code for commercial buildings effective statewide on 2024-03-15. | src:usa-wa:regulation:wac-51-11c; src:usa-wa:agency:sbcc-energy-code |
| **Energy** | Washington State Energy Code - Residential | IECC / Washington State Energy Code | 2021 | current | null | 2024-03-15 | 2024-03-15 | 2024-03-15 | WAC 51-11R adopts the 2021 Washington State Energy Code for residential buildings effective statewide on 2024-03-15. | src:usa-wa:regulation:wac-51-11r; src:usa-wa:agency:sbcc-energy-code |
| **Fire - construction references** | Washington State Fire Code | IFC | 2021 | current | null | 2024-03-15 | 2024-03-15 | 2024-03-15 | WAC 51-54A adopts the 2021 IFC effective in all counties and cities on 2024-03-15. | src:usa-wa:regulation:wac-51-54a; src:usa-wa:agency:sbcc-fire-amendments |
| **Fire - operational / prevention code** | Washington State Fire Code | IFC | 2021 | current | null | 2024-03-15 | 2024-03-15 | 2024-03-15 | Same IFC adoption; operational administration is local/AHJ-specific under the State Building Code enforcement model. | src:usa-wa:regulation:wac-51-54a; src:usa-wa:statute:rcw-19-27-050 |
| **Accessibility** | Washington State Building Code accessibility provisions | IBC Appendix E; ICC A117.1 | 2021 IBC / ICC A117.1-2017 | current | null | 2024-03-15 | 2024-03-15 | 2024-03-15 | SBCC building-code amendment page states the 2021 building code includes Appendix E and ICC/ANSI A117.1-2017; federal ADA and project-specific accessibility review are outside this matrix. | src:usa-wa:regulation:wac-51-50; src:usa-wa:agency:sbcc-building-amendments |
| **Elevator / Conveyance** | Safety Regulations and Fees for All Elevators, Dumbwaiters, Escalators and Other Conveyances | ASME A17.1/CSA B44 and related standards | ASME A17.1-2019/CSA B44-19 | current | null | 2023-10-02 | 2023-10-02 | 2023-10-02 | WAC 296-96-00650 lists ASME A17.1-2019/CSA B44-19 as current beginning 2023-10-02. | src:usa-wa:regulation:wac-296-96-00650; src:usa-wa:agency:lni-elevator-laws |

### 3.2 Adoption Records

```yaml
adoption_records:
  - adoption_id: "adoption:usa-wa:building:2021-ibc"
    state_id: "US-WA"
    code_family: "building"
    status: "current"
    state_code:
      name: "Washington State Building Code"
      edition_label: "2021 IBC with Washington amendments"
      codification: "WAC 51-50"
    base_model_code:
      publisher: "ICC"
      code_name: "International Building Code"
      edition_year: 2021
      incorporated_by_reference: true
    authority:
      adopting_authority_id: "ahj:usa-wa:state-building-code-council"
      enforcing_authority_model: "local"
      interpretation_authority_id: "ahj:usa-wa:state-building-code-council"
    dates:
      adoption_date: null
      effective_date: "2024-03-15"
      operative_date: "2024-03-15"
      mandatory_date: "2024-03-15"
      replacement_date: null
    applicability:
      date_trigger: "unknown"
      applies_to:
        - "new_construction"
        - "alteration"
        - "repair"
        - "addition"
        - "change_of_occupancy"
        - "commercial"
      exclusions: []
      special_conditions:
        - "WAC implementation rule gives a statewide effective date but does not, by itself, normalize project-level permit application or permit issuance triggers."
    transition:
      exists: true
      rule_id: "date-rule:usa-wa:sbcc-2021-statewide-effective"
      start_date: "2024-03-15"
      end_date: null
      prior_code_allowed: null
      prior_code_condition: "Resolve with local AHJ for filings around the transition date."
    amendments:
      state_amended: true
      amendment_set_ids:
        - "amendment-set:usa-wa:building:2021-ibc"
      amendment_source_ids:
        - "src:usa-wa:regulation:wac-51-50"
        - "src:usa-wa:agency:sbcc-building-amendments"
    provenance:
      source_ids:
        - "src:usa-wa:regulation:wac-51-50"
        - "src:usa-wa:agency:sbcc-building-amendments"
      field_sources:
        state_code.name: ["src:usa-wa:regulation:wac-51-50"]
        state_code.codification: ["src:usa-wa:regulation:wac-51-50"]
        base_model_code.code_name: ["src:usa-wa:regulation:wac-51-50"]
        base_model_code.edition_year: ["src:usa-wa:regulation:wac-51-50"]
        authority.adopting_authority_id: ["src:usa-wa:statute:rcw-19-27", "src:usa-wa:regulation:wac-title-51"]
        dates.effective_date: ["src:usa-wa:regulation:wac-51-50"]
        dates.operative_date: ["src:usa-wa:regulation:wac-51-50"]
        dates.mandatory_date: ["src:usa-wa:regulation:wac-51-50"]
    verification:
      status: "partially_verified"
      confidence: 0.91
      notes: "Code edition and statewide effective date verified; permit-trigger normalization requires AHJ or rulemaking detail."

  - adoption_id: "adoption:usa-wa:residential:2021-irc"
    state_id: "US-WA"
    code_family: "residential"
    status: "current"
    state_code:
      name: "Washington State Residential Code"
      edition_label: "2021 IRC with Washington amendments"
      codification: "WAC 51-51"
    base_model_code:
      publisher: "ICC"
      code_name: "International Residential Code"
      edition_year: 2021
      incorporated_by_reference: true
    authority:
      adopting_authority_id: "ahj:usa-wa:state-building-code-council"
      enforcing_authority_model: "local"
      interpretation_authority_id: "ahj:usa-wa:state-building-code-council"
    dates:
      adoption_date: null
      effective_date: "2024-03-15"
      operative_date: "2024-03-15"
      mandatory_date: "2024-03-15"
      replacement_date: null
    applicability:
      date_trigger: "unknown"
      applies_to:
        - "detached_one_two_family"
        - "townhouses"
        - "adult_family_homes"
        - "accessory_structures"
        - "alteration"
        - "repair"
        - "addition"
      exclusions:
        - "IRC Chapter 11 and Chapters 25 through 43 are not adopted."
      special_conditions:
        - "Energy, plumbing, and electrical are regulated in separate Washington chapters identified in WAC 51-51."
    transition:
      exists: true
      rule_id: "date-rule:usa-wa:sbcc-2021-statewide-effective"
      start_date: "2024-03-15"
      end_date: null
      prior_code_allowed: null
      prior_code_condition: "Resolve with local AHJ for filings around the transition date."
    amendments:
      state_amended: true
      amendment_set_ids:
        - "amendment-set:usa-wa:residential:2021-irc"
      amendment_source_ids:
        - "src:usa-wa:regulation:wac-51-51"
        - "src:usa-wa:agency:sbcc-residential-amendments"
    provenance:
      source_ids:
        - "src:usa-wa:regulation:wac-51-51"
        - "src:usa-wa:agency:sbcc-residential-amendments"
      field_sources:
        state_code.name: ["src:usa-wa:agency:sbcc-residential-amendments"]
        state_code.codification: ["src:usa-wa:regulation:wac-51-51"]
        base_model_code.code_name: ["src:usa-wa:regulation:wac-51-51"]
        base_model_code.edition_year: ["src:usa-wa:regulation:wac-51-51"]
        authority.adopting_authority_id: ["src:usa-wa:statute:rcw-19-27", "src:usa-wa:regulation:wac-title-51"]
        dates.effective_date: ["src:usa-wa:regulation:wac-51-51"]
    verification:
      status: "partially_verified"
      confidence: 0.91
      notes: "Core edition, exclusions, and statewide effective date verified."

  - adoption_id: "adoption:usa-wa:existing-building:2021-iebc"
    state_id: "US-WA"
    code_family: "existing_building"
    status: "current"
    state_code:
      name: "Washington State Building Code"
      edition_label: "2021 IEBC included in 2021 IBC adoption"
      codification: "WAC 51-50"
    base_model_code:
      publisher: "ICC"
      code_name: "International Existing Building Code"
      edition_year: 2021
      incorporated_by_reference: true
    authority:
      adopting_authority_id: "ahj:usa-wa:state-building-code-council"
      enforcing_authority_model: "local"
      interpretation_authority_id: "ahj:usa-wa:state-building-code-council"
    dates:
      adoption_date: null
      effective_date: "2024-03-15"
      operative_date: "2024-03-15"
      mandatory_date: "2024-03-15"
      replacement_date: null
    applicability:
      date_trigger: "unknown"
      applies_to:
        - "existing_buildings"
        - "alteration"
        - "repair"
        - "change_of_occupancy"
      exclusions: []
      special_conditions:
        - "The 2021 IEBC is included through the 2021 IBC adoption and is amended in WAC 51-50-480000."
    transition:
      exists: true
      rule_id: "date-rule:usa-wa:sbcc-2021-statewide-effective"
      start_date: "2024-03-15"
      end_date: null
      prior_code_allowed: null
      prior_code_condition: "Resolve with local AHJ for filings around the transition date."
    amendments:
      state_amended: true
      amendment_set_ids:
        - "amendment-set:usa-wa:existing-building:2021-iebc"
      amendment_source_ids:
        - "src:usa-wa:regulation:wac-51-50"
        - "src:usa-wa:agency:sbcc-building-amendments"
    provenance:
      source_ids:
        - "src:usa-wa:regulation:wac-51-50"
        - "src:usa-wa:agency:sbcc-building-amendments"
      field_sources:
        state_code.codification: ["src:usa-wa:regulation:wac-51-50"]
        base_model_code.code_name: ["src:usa-wa:regulation:wac-51-50"]
        base_model_code.edition_year: ["src:usa-wa:agency:sbcc-building-amendments"]
        dates.effective_date: ["src:usa-wa:regulation:wac-51-50"]
    verification:
      status: "partially_verified"
      confidence: 0.88
      notes: "Included status verified; detailed IEBC amendments not parsed."

  - adoption_id: "adoption:usa-wa:mechanical:2021-imc"
    state_id: "US-WA"
    code_family: "mechanical"
    status: "current"
    state_code:
      name: "Washington State Mechanical Code"
      edition_label: "2021 IMC with Washington amendments"
      codification: "WAC 51-52"
    base_model_code:
      publisher: "ICC"
      code_name: "International Mechanical Code"
      edition_year: 2021
      incorporated_by_reference: true
    authority:
      adopting_authority_id: "ahj:usa-wa:state-building-code-council"
      enforcing_authority_model: "local"
      interpretation_authority_id: "ahj:usa-wa:state-building-code-council"
    dates:
      adoption_date: null
      effective_date: "2024-03-15"
      operative_date: "2024-03-15"
      mandatory_date: "2024-03-15"
      replacement_date: null
    applicability:
      date_trigger: "unknown"
      applies_to:
        - "mechanical_systems"
        - "commercial"
        - "alteration"
        - "repair"
      exclusions: []
      special_conditions:
        - "Fuel gas distribution piping and equipment are regulated by the International Fuel Gas Code within WAC 51-52."
    transition:
      exists: true
      rule_id: "date-rule:usa-wa:sbcc-2021-statewide-effective"
      start_date: "2024-03-15"
      end_date: null
      prior_code_allowed: null
      prior_code_condition: "Resolve with local AHJ for filings around the transition date."
    amendments:
      state_amended: true
      amendment_set_ids:
        - "amendment-set:usa-wa:mechanical:2021-imc"
      amendment_source_ids:
        - "src:usa-wa:regulation:wac-51-52"
        - "src:usa-wa:agency:sbcc-mechanical-amendments"
    provenance:
      source_ids:
        - "src:usa-wa:regulation:wac-51-52"
        - "src:usa-wa:agency:sbcc-mechanical-amendments"
      field_sources:
        state_code.codification: ["src:usa-wa:regulation:wac-51-52"]
        base_model_code.code_name: ["src:usa-wa:regulation:wac-51-52"]
        base_model_code.edition_year: ["src:usa-wa:regulation:wac-51-52"]
        dates.effective_date: ["src:usa-wa:regulation:wac-51-52"]
    verification:
      status: "partially_verified"
      confidence: 0.91
      notes: "Core edition and effective date verified."

  - adoption_id: "adoption:usa-wa:fuel-gas:2021-ifgc"
    state_id: "US-WA"
    code_family: "fuel_gas"
    status: "current"
    state_code:
      name: "Washington State Mechanical Code fuel-gas provisions"
      edition_label: "2021 IFGC; 2020 NFPA 58; 2021 ANSI Z223.1/NFPA 54"
      codification: "WAC 51-52"
    base_model_code:
      publisher: "ICC/NFPA"
      code_name: "International Fuel Gas Code; Liquefied Petroleum Gas Code; National Fuel Gas Code"
      edition_year: 2021
      incorporated_by_reference: true
    authority:
      adopting_authority_id: "ahj:usa-wa:state-building-code-council"
      enforcing_authority_model: "local"
      interpretation_authority_id: "ahj:usa-wa:state-building-code-council"
    dates:
      adoption_date: null
      effective_date: "2024-03-15"
      operative_date: "2024-03-15"
      mandatory_date: "2024-03-15"
      replacement_date: null
    applicability:
      date_trigger: "unknown"
      applies_to:
        - "fuel_gas_piping"
        - "fuel_gas_utilization_equipment"
        - "gaseous_hydrogen_systems"
        - "regulated_accessories"
      exclusions:
        - "Detached one- and two-family dwellings and qualifying townhouses follow the IRC."
      special_conditions:
        - "LP-gas installations use the NFPA 58 and ANSI Z223.1/NFPA 54 editions identified in WAC 51-52."
    transition:
      exists: true
      rule_id: "date-rule:usa-wa:sbcc-2021-statewide-effective"
      start_date: "2024-03-15"
      end_date: null
      prior_code_allowed: null
      prior_code_condition: "Resolve with local AHJ for filings around the transition date."
    amendments:
      state_amended: true
      amendment_set_ids:
        - "amendment-set:usa-wa:fuel-gas:2021-ifgc"
      amendment_source_ids:
        - "src:usa-wa:regulation:wac-51-52"
        - "src:usa-wa:agency:sbcc-mechanical-amendments"
    provenance:
      source_ids:
        - "src:usa-wa:regulation:wac-51-52"
        - "src:usa-wa:agency:sbcc-mechanical-amendments"
      field_sources:
        state_code.codification: ["src:usa-wa:regulation:wac-51-52"]
        base_model_code.code_name: ["src:usa-wa:regulation:wac-51-52"]
        base_model_code.edition_year: ["src:usa-wa:agency:sbcc-mechanical-amendments"]
        dates.effective_date: ["src:usa-wa:regulation:wac-51-52"]
    verification:
      status: "partially_verified"
      confidence: 0.89
      notes: "Fuel-gas scope verified from WAC 51-52; detailed IFGC amendments not parsed."

  - adoption_id: "adoption:usa-wa:plumbing:2021-upc"
    state_id: "US-WA"
    code_family: "plumbing"
    status: "current"
    state_code:
      name: "Washington State Plumbing Code"
      edition_label: "2021 UPC with Washington amendments"
      codification: "WAC 51-56"
    base_model_code:
      publisher: "IAPMO"
      code_name: "Uniform Plumbing Code"
      edition_year: 2021
      incorporated_by_reference: true
    authority:
      adopting_authority_id: "ahj:usa-wa:state-building-code-council"
      enforcing_authority_model: "local"
      interpretation_authority_id: "ahj:usa-wa:state-building-code-council"
    dates:
      adoption_date: null
      effective_date: "2024-03-15"
      operative_date: "2024-03-15"
      mandatory_date: "2024-03-15"
      replacement_date: null
    applicability:
      date_trigger: "unknown"
      applies_to:
        - "plumbing_systems"
        - "new_construction"
        - "alteration"
        - "repair"
      exclusions:
        - "UPC chapters 12 and 14 are not adopted."
        - "UPC requirements relating to venting and combustion air of fuel-fired appliances in chapter 5, and portions addressing building sewers, are not adopted."
      special_conditions: []
    transition:
      exists: true
      rule_id: "date-rule:usa-wa:sbcc-2021-statewide-effective"
      start_date: "2024-03-15"
      end_date: null
      prior_code_allowed: null
      prior_code_condition: "Resolve with local AHJ for filings around the transition date."
    amendments:
      state_amended: true
      amendment_set_ids:
        - "amendment-set:usa-wa:plumbing:2021-upc"
      amendment_source_ids:
        - "src:usa-wa:regulation:wac-51-56"
    provenance:
      source_ids:
        - "src:usa-wa:regulation:wac-51-56"
      field_sources:
        state_code.codification: ["src:usa-wa:regulation:wac-51-56"]
        base_model_code.code_name: ["src:usa-wa:regulation:wac-51-56"]
        base_model_code.edition_year: ["src:usa-wa:regulation:wac-51-56"]
        dates.effective_date: ["src:usa-wa:regulation:wac-51-56"]
    verification:
      status: "partially_verified"
      confidence: 0.91
      notes: "Core edition, effective date, and major exclusions verified."

  - adoption_id: "adoption:usa-wa:electrical:2023-nec"
    state_id: "US-WA"
    code_family: "electrical"
    status: "current"
    state_code:
      name: "Electrical Safety Standards, Administration, and Installation"
      edition_label: "2023 NEC with Washington amendments"
      codification: "WAC 296-46B"
    base_model_code:
      publisher: "NFPA"
      code_name: "National Electrical Code"
      edition_year: 2023
      incorporated_by_reference: true
    authority:
      adopting_authority_id: "ahj:usa-wa:lni-electrical"
      enforcing_authority_model: "hybrid"
      interpretation_authority_id: "ahj:usa-wa:lni-electrical"
    dates:
      adoption_date: null
      effective_date: "2024-04-01"
      operative_date: "2024-04-01"
      mandatory_date: "2024-04-01"
      replacement_date: "2026-12-31"
    applicability:
      date_trigger: "permit_issuance_date"
      applies_to:
        - "electrical_installations"
        - "new_construction"
        - "alteration"
        - "repair"
      exclusions:
        - "WAC 296-46B-010 provides exceptions for certain dwelling projects based on building permit issue date and plan-review timing."
      special_conditions:
        - "Effective 2026-12-31, the 2026 NEC replaces the 2023 NEC under the current WAC text."
    transition:
      exists: true
      rule_id: "date-rule:usa-wa:electrical-permit-issue-date"
      start_date: "2024-04-01"
      end_date: "2026-12-30"
      prior_code_allowed: true
      prior_code_condition: "Only as allowed by WAC 296-46B-010 exceptions tied to building permit issue date or accepted plan review."
    amendments:
      state_amended: true
      amendment_set_ids:
        - "amendment-set:usa-wa:electrical:2023-nec"
      amendment_source_ids:
        - "src:usa-wa:regulation:wac-296-46b"
        - "src:usa-wa:regulation:wac-296-46b-010"
    provenance:
      source_ids:
        - "src:usa-wa:statute:rcw-19-28"
        - "src:usa-wa:regulation:wac-296-46b"
        - "src:usa-wa:regulation:wac-296-46b-010"
        - "src:usa-wa:agency:lni-electrical-laws"
      field_sources:
        state_code.codification: ["src:usa-wa:regulation:wac-296-46b"]
        base_model_code.code_name: ["src:usa-wa:regulation:wac-296-46b-010"]
        base_model_code.edition_year: ["src:usa-wa:regulation:wac-296-46b-010"]
        authority.adopting_authority_id: ["src:usa-wa:statute:rcw-19-28", "src:usa-wa:agency:lni-electrical-laws"]
        dates.effective_date: ["src:usa-wa:regulation:wac-296-46b-010"]
        dates.replacement_date: ["src:usa-wa:regulation:wac-296-46b-010", "src:usa-wa:agency:lni-electrical-laws"]
        applicability.date_trigger: ["src:usa-wa:regulation:wac-296-46b-010"]
    verification:
      status: "verified"
      confidence: 0.94
      notes: "Current edition, future NEC replacement date, and electrical permit-issue trigger verified from current WAC text."

  - adoption_id: "adoption:usa-wa:energy-commercial:2021-wsec"
    state_id: "US-WA"
    code_family: "energy"
    status: "current"
    state_code:
      name: "Washington State Energy Code - Commercial"
      edition_label: "2021 WSEC Commercial"
      codification: "WAC 51-11C"
    base_model_code:
      publisher: "ICC/state"
      code_name: "International Energy Conservation Code / Washington State Energy Code"
      edition_year: 2021
      incorporated_by_reference: true
    authority:
      adopting_authority_id: "ahj:usa-wa:state-building-code-council"
      enforcing_authority_model: "local"
      interpretation_authority_id: "ahj:usa-wa:state-building-code-council"
    dates:
      adoption_date: null
      effective_date: "2024-03-15"
      operative_date: "2024-03-15"
      mandatory_date: "2024-03-15"
      replacement_date: null
    applicability:
      date_trigger: "unknown"
      applies_to:
        - "commercial"
        - "group_r_sleeping_units_as_commercial_when_required"
        - "new_construction"
        - "alteration"
      exclusions:
        - "Temporary growing structures used solely for commercial horticultural production are excluded from WAC 51-11C scope as stated in WAC 51-11C-10100."
      special_conditions:
        - "Commercial includes buildings not covered under the residential energy code."
    transition:
      exists: true
      rule_id: "date-rule:usa-wa:sbcc-2021-statewide-effective"
      start_date: "2024-03-15"
      end_date: null
      prior_code_allowed: null
      prior_code_condition: "Resolve with local AHJ for filings around the transition date."
    amendments:
      state_amended: true
      amendment_set_ids:
        - "amendment-set:usa-wa:energy-commercial:2021-wsec"
      amendment_source_ids:
        - "src:usa-wa:regulation:wac-51-11c"
        - "src:usa-wa:agency:sbcc-energy-code"
    provenance:
      source_ids:
        - "src:usa-wa:regulation:wac-51-11c"
        - "src:usa-wa:agency:sbcc-energy-code"
      field_sources:
        state_code.name: ["src:usa-wa:regulation:wac-51-11c", "src:usa-wa:agency:sbcc-energy-code"]
        state_code.codification: ["src:usa-wa:regulation:wac-51-11c"]
        base_model_code.code_name: ["src:usa-wa:agency:sbcc-energy-code"]
        base_model_code.edition_year: ["src:usa-wa:regulation:wac-51-11c"]
        dates.effective_date: ["src:usa-wa:regulation:wac-51-11c", "src:usa-wa:agency:sbcc-energy-code"]
    verification:
      status: "partially_verified"
      confidence: 0.92
      notes: "Edition, scope, and statewide effective date verified; detailed energy amendments not parsed."

  - adoption_id: "adoption:usa-wa:energy-residential:2021-wsec"
    state_id: "US-WA"
    code_family: "energy"
    status: "current"
    state_code:
      name: "Washington State Energy Code - Residential"
      edition_label: "2021 WSEC Residential"
      codification: "WAC 51-11R"
    base_model_code:
      publisher: "ICC/state"
      code_name: "International Energy Conservation Code / Washington State Energy Code"
      edition_year: 2021
      incorporated_by_reference: true
    authority:
      adopting_authority_id: "ahj:usa-wa:state-building-code-council"
      enforcing_authority_model: "local"
      interpretation_authority_id: "ahj:usa-wa:state-building-code-council"
    dates:
      adoption_date: null
      effective_date: "2024-03-15"
      operative_date: "2024-03-15"
      mandatory_date: "2024-03-15"
      replacement_date: null
    applicability:
      date_trigger: "unknown"
      applies_to:
        - "detached_one_two_family"
        - "townhouses"
        - "r_3_buildings_three_stories_or_less"
        - "r_2_buildings_three_stories_or_less_with_exterior_unit_access"
      exclusions: []
      special_conditions:
        - "SBCC states residential includes one- and two-family dwellings, townhouses, certain R-3 buildings, and certain R-2 buildings with direct exterior dwelling-unit access."
    transition:
      exists: true
      rule_id: "date-rule:usa-wa:sbcc-2021-statewide-effective"
      start_date: "2024-03-15"
      end_date: null
      prior_code_allowed: null
      prior_code_condition: "Resolve with local AHJ for filings around the transition date."
    amendments:
      state_amended: true
      amendment_set_ids:
        - "amendment-set:usa-wa:energy-residential:2021-wsec"
      amendment_source_ids:
        - "src:usa-wa:regulation:wac-51-11r"
        - "src:usa-wa:agency:sbcc-energy-code"
    provenance:
      source_ids:
        - "src:usa-wa:regulation:wac-51-11r"
        - "src:usa-wa:agency:sbcc-energy-code"
      field_sources:
        state_code.name: ["src:usa-wa:regulation:wac-51-11r", "src:usa-wa:agency:sbcc-energy-code"]
        state_code.codification: ["src:usa-wa:regulation:wac-51-11r"]
        base_model_code.code_name: ["src:usa-wa:agency:sbcc-energy-code"]
        base_model_code.edition_year: ["src:usa-wa:regulation:wac-51-11r"]
        dates.effective_date: ["src:usa-wa:regulation:wac-51-11r", "src:usa-wa:agency:sbcc-energy-code"]
    verification:
      status: "partially_verified"
      confidence: 0.92
      notes: "Edition, scope, and statewide effective date verified; detailed energy amendments not parsed."

  - adoption_id: "adoption:usa-wa:fire:2021-ifc"
    state_id: "US-WA"
    code_family: "fire_operational"
    status: "current"
    state_code:
      name: "Washington State Fire Code"
      edition_label: "2021 IFC with Washington amendments"
      codification: "WAC 51-54A"
    base_model_code:
      publisher: "ICC"
      code_name: "International Fire Code"
      edition_year: 2021
      incorporated_by_reference: true
    authority:
      adopting_authority_id: "ahj:usa-wa:state-building-code-council"
      enforcing_authority_model: "local"
      interpretation_authority_id: "ahj:usa-wa:local-fire-code-official"
    dates:
      adoption_date: null
      effective_date: "2024-03-15"
      operative_date: "2024-03-15"
      mandatory_date: "2024-03-15"
      replacement_date: null
    applicability:
      date_trigger: "unknown"
      applies_to:
        - "new_construction"
        - "existing_buildings"
        - "operational_fire_code"
        - "commercial"
      exclusions: []
      special_conditions:
        - "WAC 51-54A contains an emergency rule path for IFC Section 308.1.4 noted by SBCC; operational permit administration should be resolved with local fire AHJ."
    transition:
      exists: true
      rule_id: "date-rule:usa-wa:sbcc-2021-statewide-effective"
      start_date: "2024-03-15"
      end_date: null
      prior_code_allowed: null
      prior_code_condition: "Resolve with local fire AHJ for filings around the transition date."
    amendments:
      state_amended: true
      amendment_set_ids:
        - "amendment-set:usa-wa:fire:2021-ifc"
      amendment_source_ids:
        - "src:usa-wa:regulation:wac-51-54a"
        - "src:usa-wa:agency:sbcc-fire-amendments"
    provenance:
      source_ids:
        - "src:usa-wa:regulation:wac-51-54a"
        - "src:usa-wa:agency:sbcc-fire-amendments"
      field_sources:
        state_code.codification: ["src:usa-wa:regulation:wac-51-54a"]
        base_model_code.code_name: ["src:usa-wa:regulation:wac-51-54a"]
        base_model_code.edition_year: ["src:usa-wa:regulation:wac-51-54a"]
        dates.effective_date: ["src:usa-wa:regulation:wac-51-54a", "src:usa-wa:agency:sbcc-fire-amendments"]
    verification:
      status: "partially_verified"
      confidence: 0.91
      notes: "Edition and statewide effective date verified; local operational details remain AHJ-specific."

  - adoption_id: "adoption:usa-wa:accessibility:2021-ibc-a117"
    state_id: "US-WA"
    code_family: "accessibility"
    status: "current"
    state_code:
      name: "Washington State Building Code accessibility provisions"
      edition_label: "2021 IBC Appendix E and ICC/ANSI A117.1-2017"
      codification: "WAC 51-50"
    base_model_code:
      publisher: "ICC"
      code_name: "International Building Code Appendix E; ICC A117.1"
      edition_year: 2021
      incorporated_by_reference: true
    authority:
      adopting_authority_id: "ahj:usa-wa:state-building-code-council"
      enforcing_authority_model: "local"
      interpretation_authority_id: "ahj:usa-wa:state-building-code-council"
    dates:
      adoption_date: null
      effective_date: "2024-03-15"
      operative_date: "2024-03-15"
      mandatory_date: "2024-03-15"
      replacement_date: null
    applicability:
      date_trigger: "unknown"
      applies_to:
        - "accessible_routes"
        - "public_buildings"
        - "commercial"
        - "residential_as_required_by_code"
      exclusions:
        - "Federal ADA compliance and project-specific civil-rights review are outside this state-code adoption record."
      special_conditions:
        - "SBCC building-code amendment page states the 2021 building code includes Appendix E and ICC/ANSI A117.1-2017."
    transition:
      exists: true
      rule_id: "date-rule:usa-wa:sbcc-2021-statewide-effective"
      start_date: "2024-03-15"
      end_date: null
      prior_code_allowed: null
      prior_code_condition: "Resolve with local AHJ for filings around the transition date."
    amendments:
      state_amended: true
      amendment_set_ids:
        - "amendment-set:usa-wa:accessibility:2021-ibc-a117"
      amendment_source_ids:
        - "src:usa-wa:regulation:wac-51-50"
        - "src:usa-wa:agency:sbcc-building-amendments"
    provenance:
      source_ids:
        - "src:usa-wa:regulation:wac-51-50"
        - "src:usa-wa:agency:sbcc-building-amendments"
      field_sources:
        state_code.codification: ["src:usa-wa:regulation:wac-51-50"]
        base_model_code.code_name: ["src:usa-wa:agency:sbcc-building-amendments"]
        base_model_code.edition_year: ["src:usa-wa:agency:sbcc-building-amendments"]
        dates.effective_date: ["src:usa-wa:agency:sbcc-building-amendments"]
    verification:
      status: "partially_verified"
      confidence: 0.83
      notes: "State building-code accessibility publication path verified; detailed accessibility scoping not parsed."

  - adoption_id: "adoption:usa-wa:elevator:2019-asme-a17-1"
    state_id: "US-WA"
    code_family: "elevator"
    status: "current"
    state_code:
      name: "Safety regulations and fees for all elevators, dumbwaiters, escalators and other conveyances"
      edition_label: "ASME A17.1-2019/CSA B44-19 and related standards"
      codification: "WAC 296-96"
    base_model_code:
      publisher: "ASME/CSA"
      code_name: "Safety Code for Elevators and Escalators"
      edition_year: 2019
      incorporated_by_reference: true
    authority:
      adopting_authority_id: "ahj:usa-wa:lni-elevator"
      enforcing_authority_model: "state"
      interpretation_authority_id: "ahj:usa-wa:lni-elevator"
    dates:
      adoption_date: null
      effective_date: "2023-10-02"
      operative_date: "2023-10-02"
      mandatory_date: "2023-10-02"
      replacement_date: null
    applicability:
      date_trigger: "unknown"
      applies_to:
        - "elevators"
        - "escalators"
        - "dumbwaiters"
        - "residential_elevators"
        - "special_purpose_elevators"
        - "alteration"
      exclusions: []
      special_conditions:
        - "WAC 296-96 includes related current standards beyond ASME A17.1; this record captures the primary new-elevator/escalator model-code anchor only."
    transition:
      exists: true
      rule_id: "date-rule:usa-wa:elevator-current-standards"
      start_date: "2023-10-02"
      end_date: null
      prior_code_allowed: null
      prior_code_condition: "Resolve by equipment type and permit date under WAC 296-96."
    amendments:
      state_amended: true
      amendment_set_ids:
        - "amendment-set:usa-wa:elevator:296-96"
      amendment_source_ids:
        - "src:usa-wa:regulation:wac-296-96"
        - "src:usa-wa:regulation:wac-296-96-00650"
    provenance:
      source_ids:
        - "src:usa-wa:agency:lni-elevator-laws"
        - "src:usa-wa:regulation:wac-296-96"
        - "src:usa-wa:regulation:wac-296-96-00650"
      field_sources:
        state_code.codification: ["src:usa-wa:regulation:wac-296-96"]
        base_model_code.code_name: ["src:usa-wa:regulation:wac-296-96-00650"]
        base_model_code.edition_year: ["src:usa-wa:regulation:wac-296-96-00650"]
        dates.effective_date: ["src:usa-wa:regulation:wac-296-96-00650"]
    verification:
      status: "partially_verified"
      confidence: 0.88
      notes: "Primary current ASME A17.1/CSA B44 standard and effective date verified; equipment-specific rules not fully parsed."
```

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

Washington's SBCC-administered 2021 code cycle is anchored by WAC implementation sections and SBCC pages stating a statewide effective date of 2024-03-15. The WAC implementation language says the adopted codes become effective in all counties and cities on that date. This report does not convert that statewide effective-date rule into a universal building-permit application or issuance trigger because local permit intake practices may still affect transition handling.

Electrical code transition logic is more explicit. WAC 296-46B-010 states that adopted standards apply to installations when electrical permit issue dates are on and after adoption dates, with stated exceptions for certain dwelling projects and accepted plan review. Elevator/conveyance transition logic remains equipment-specific under WAC 296-96.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-wa:sbcc-2021-statewide-effective | WAC 51-50, 51-51, 51-52, 51-54A, 51-56, 51-11C, 51-11R | effective_date / operative_date | 2024-03-15 | Code chapter is effective in all counties and cities; project filing trigger remains AHJ-sensitive | unknown | src:usa-wa:regulation:wac-51-50; src:usa-wa:regulation:wac-51-51; src:usa-wa:regulation:wac-51-52; src:usa-wa:regulation:wac-51-54a; src:usa-wa:regulation:wac-51-56; src:usa-wa:regulation:wac-51-11c; src:usa-wa:regulation:wac-51-11r; src:usa-wa:agency:sbcc-current-codes | partially_verified |
| date-rule:usa-wa:electrical-permit-issue-date | WAC 296-46B electrical standards | effective_date / permit_issuance_date | 2024-04-01 through 2026-12-30 for 2023 NEC; 2026-12-31 for 2026 NEC replacement | Electrical permit issue date on or after standard adoption date, subject to WAC 296-46B-010 exceptions | yes, limited | src:usa-wa:regulation:wac-296-46b-010; src:usa-wa:agency:lni-electrical-laws | verified |
| date-rule:usa-wa:elevator-current-standards | WAC 296-96 elevator/conveyance standards | effective_date | 2023-10-02 for ASME A17.1-2019/CSA B44-19 | Equipment-specific conveyance rules and permit/inspection requirements | unknown | src:usa-wa:regulation:wac-296-96-00650; src:usa-wa:agency:lni-elevator-laws | partially_verified |
| date-rule:usa-wa:sbcc-2024-future-cycle | 2024 Washington State Building Codes | future_effective_date | Final adoption delayed to 2026-08-21; effective date 2027-05-03 | SBCC 2024 code adoption schedule and final rulemaking | not applicable | src:usa-wa:agency:sbcc-home-2024-cycle | partially_verified |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| SBCC-administered State Building Code families | 2024 Washington State Building Codes | 2026-01-23 | 2026-08-21 | 2027-05-03 | 2027-05-03 | 2027-05-03 | monitoring | src:usa-wa:agency:sbcc-home-2024-cycle | SBCC states final adoption was delayed until 2026-08-21 and the 2024-code effective date will be 2027-05-03. |
| Electrical | 2026 NEC / NFPA 70-2026 | null | null | 2026-12-31 | 2026-12-31 | 2026-12-31 | monitoring | src:usa-wa:regulation:wac-296-46b-010; src:usa-wa:agency:lni-electrical-laws | WAC 296-46B-010 states the 2026 NEC replaces the 2023 NEC effective 2026-12-31. |
| Wildland-Urban Interface | 2021 WUI code status / future wildfire mapping dependent amendments | null | null | null | null | null | monitoring | src:usa-wa:agency:sbcc-home-2024-cycle; src:usa-wa:regulation:wac-title-51 | WAC Title 51 lists WAC 51-55 for the 2021 WUI code, while SBCC homepage notes WUI amendments are in a reserved status pending statewide wildfire risk and hazard mapping. This report does not normalize WUI as a standard code-family row. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-wa:local-building-code-enforcement | all SBCC State Building Code chapters | local building permits and inspections | local permit or inspection | Counties and cities enforce the State Building Code; if they lack a building department, they must contract for enforcement. | src:usa-wa:statute:rcw-19-27-050 | verified |
| applicability-rule:usa-wa:local-amendments | all SBCC State Building Code chapters | local code amendments | local ordinance or resolution | Counties and cities may amend the codes listed in RCW 19.27.031 within RCW 19.27.060 limits; amendments may not reduce minimum performance standards, and residential amendments generally require SBCC approval. | src:usa-wa:statute:rcw-19-27-060 | verified |
| applicability-rule:usa-wa:electrical-transition | electrical | electrical installations | electrical permit issue date | Adopted electrical standards apply when electrical permit issue dates are on and after standard adoption dates, with WAC 296-46B-010 exceptions. | src:usa-wa:regulation:wac-296-46b-010 | verified |
| applicability-rule:usa-wa:energy-commercial-residential-scope | energy | commercial versus residential energy-code scope | building type and occupancy | SBCC distinguishes WSEC Residential from WSEC Commercial; residential includes listed low-rise dwelling categories, and commercial includes buildings not covered under residential. | src:usa-wa:agency:sbcc-energy-code; src:usa-wa:regulation:wac-51-11c; src:usa-wa:regulation:wac-51-11r | verified |

---

## 5. State Amendments

### 5.1 Amendment Model

- **State amendment structure:** Washington amendments to national model codes plus state-written/state-specific Washington State Energy Code chapters.
- **Where amendments are published:** WAC Title 51 chapters, SBCC code-amendment pages, and SBCC insert pages for model codes. Electrical amendments are in WAC 296-46B. Elevator/conveyance amendments are in WAC 296-96.
- **Amendment parsing status:** partial

### 5.2 State Amendment Sources

| Amendment Set ID | Code Family | Publication Format | Source IDs | Parsed? | Notes |
| --- | --- | --- | --- | --- | --- |
| amendment-set:usa-wa:building:2021-ibc | building | WAC chapter and SBCC insert pages | src:usa-wa:regulation:wac-51-50; src:usa-wa:agency:sbcc-building-amendments | partial | Core edition, Appendix E, IEBC inclusion, A117.1 reference, and statewide effective date verified. |
| amendment-set:usa-wa:residential:2021-irc | residential | WAC chapter and SBCC insert pages | src:usa-wa:regulation:wac-51-51; src:usa-wa:agency:sbcc-residential-amendments | partial | Core edition, chapter exclusions, and statewide effective date verified. |
| amendment-set:usa-wa:existing-building:2021-iebc | existing_building | WAC chapter and SBCC insert pages | src:usa-wa:regulation:wac-51-50; src:usa-wa:agency:sbcc-building-amendments | partial | IEBC inclusion verified; detailed section amendments not parsed. |
| amendment-set:usa-wa:mechanical:2021-imc | mechanical | WAC chapter and SBCC insert pages | src:usa-wa:regulation:wac-51-52; src:usa-wa:agency:sbcc-mechanical-amendments | partial | Core edition and effective date verified. |
| amendment-set:usa-wa:fuel-gas:2021-ifgc | fuel_gas | WAC chapter and SBCC insert pages | src:usa-wa:regulation:wac-51-52; src:usa-wa:agency:sbcc-mechanical-amendments | partial | IFGC/NFPA fuel-gas standards verified. |
| amendment-set:usa-wa:plumbing:2021-upc | plumbing | WAC chapter | src:usa-wa:regulation:wac-51-56 | partial | Core edition, appendices, and exclusions verified. |
| amendment-set:usa-wa:energy-commercial:2021-wsec | energy | state-written WAC chapter and SBCC page | src:usa-wa:regulation:wac-51-11c; src:usa-wa:agency:sbcc-energy-code | partial | Scope and effective date verified. |
| amendment-set:usa-wa:energy-residential:2021-wsec | energy | state-written WAC chapter and SBCC page | src:usa-wa:regulation:wac-51-11r; src:usa-wa:agency:sbcc-energy-code | partial | Scope and effective date verified. |
| amendment-set:usa-wa:fire:2021-ifc | fire | WAC chapter and SBCC insert pages | src:usa-wa:regulation:wac-51-54a; src:usa-wa:agency:sbcc-fire-amendments | partial | Core edition, effective date, and emergency rule path noted. |
| amendment-set:usa-wa:accessibility:2021-ibc-a117 | accessibility | WAC chapter and SBCC insert pages | src:usa-wa:regulation:wac-51-50; src:usa-wa:agency:sbcc-building-amendments | partial | Appendix E and ICC/ANSI A117.1-2017 publication path verified. |
| amendment-set:usa-wa:electrical:2023-nec | electrical | WAC chapter | src:usa-wa:regulation:wac-296-46b; src:usa-wa:regulation:wac-296-46b-010 | partial | NEC adoption, future 2026 replacement, and permit-issue trigger verified. |
| amendment-set:usa-wa:elevator:296-96 | elevator | WAC chapter | src:usa-wa:regulation:wac-296-96; src:usa-wa:regulation:wac-296-96-00650 | partial | Current ASME A17.1/CSA B44 anchor verified. |

### 5.3 High-Impact State Amendments

| Code Family | Section | Amendment Type | Summary | Source IDs | Confidence |
| --- | --- | --- | --- | --- | --- |
| building | WAC 51-50-007 / WAC 51-50-480000 | include / amend | The 2021 IEBC is included in the IBC adoption and amended in WAC 51-50-480000. | src:usa-wa:regulation:wac-51-50 | 0.88 |
| residential | WAC 51-51-007 | delete / not adopt | IRC Chapter 11 and Chapters 25 through 43 are not adopted; energy, plumbing, and electrical are separately regulated. | src:usa-wa:regulation:wac-51-51; src:usa-wa:agency:sbcc-residential-amendments | 0.90 |
| plumbing | WAC 51-56-003 | delete / not adopt | UPC Chapters 12 and 14 are not adopted; certain venting/combustion-air and building-sewer provisions are excluded. | src:usa-wa:regulation:wac-51-56 | 0.88 |
| electrical | WAC 296-46B-010 | replace future edition | The 2026 NEC replaces the 2023 NEC on 2026-12-31 under current WAC text. | src:usa-wa:regulation:wac-296-46b-010 | 0.94 |
| fire | WAC 51-54A / SBCC fire amendment page | emergency rule path | SBCC notes an emergency rule on IFC Section 308.1.4 effective 2024-03-15. | src:usa-wa:agency:sbcc-fire-amendments | 0.78 |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-wa"
  model: "hybrid"
  enforcing_entities:
    - "county"
    - "city"
    - "approved_inspection_agency_by_contract"
    - "local_fire_code_official"
    - "labor_and_industries_for_state_electrical_jurisdiction"
    - "city_with_electrical_inspection_jurisdiction"
    - "labor_and_industries_elevator_program"
  required_officials:
    - "building_official"
    - "fire_code_official"
    - "electrical_inspector"
    - "elevator_inspector"
  state_reserved_activities:
    - "statewide_code_adoption"
    - "statewide_amendment_process"
    - "electrical_rule_administration_under_wac_296_46b"
    - "elevator_conveyance_rule_administration_under_wac_296_96"
  source_ids:
    - "src:usa-wa:statute:rcw-19-27-050"
    - "src:usa-wa:statute:rcw-19-28"
    - "src:usa-wa:regulation:wac-296-46b-010"
    - "src:usa-wa:agency:lni-elevator-laws"
  verification_status: "partially_verified"
  confidence: 0.88
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-wa"
  model: "allowed_with_limits_and_residential_sbc_approval"
  applies_to_code_families:
    - "building"
    - "residential"
    - "existing_building"
    - "mechanical"
    - "plumbing"
    - "fuel_gas"
    - "energy"
    - "fire"
    - "accessibility"
  approval_required: true
  approving_authority_id: "ahj:usa-wa:state-building-code-council"
  filing_required: null
  registry_exists: null
  registry_source_ids:
    - "src:usa-wa:regulation:wac-51-04"
  legal_basis_source_ids:
    - "src:usa-wa:statute:rcw-19-27-060"
  verification_status: "partially_verified"
  confidence: 0.84
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement and local amendment authority are separate questions in Washington. Counties and cities enforce the statewide State Building Code locally. That enforcement role does not create unlimited authority to change the statewide code. RCW 19.27.060 allows counties and cities to amend the codes, but the amendments may not reduce minimum state performance standards; amendments affecting single-family or multifamily residential buildings generally must be approved by SBCC.

Electrical inspection jurisdiction is also separate from general building-code enforcement. WAC 296-46B refers to inspection by the department or a city authorized to do electrical inspections. Elevator/conveyance regulation is an L&I state program under a separate WAC chapter, even though a local building permit may still be required for associated construction.

### 6.4 Known Local Amendment Registries

| Registry ID | Scope | Maintained By | Source ID | Machine Readable? | Parsed? | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| registry:usa-wa:local-amendments-sbcc | statewide local code-amendment process | Washington State Building Code Council | src:usa-wa:regulation:wac-51-04 | partial | no | WAC 51-04 is identified as the chapter for statewide and local amendment procedures, but this pass did not parse a local amendment registry or approval database. |

### 6.5 Municipality-Specific Known Amendments

| Jurisdiction | Code Family | Amendment Set ID | Approval Status | Effective Date | Source IDs | Parsed? |
| --- | --- | --- | --- | --- | --- | --- |
| none populated | none | none | none | null | none | no |

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

- **Resolver status:** state_only

**Jurisdiction stack:**

```text
Address
  -> State of Washington
  -> County
  -> City or unincorporated county
  -> Local building department or contracted inspection agency
  -> Local fire code official / fire marshal
  -> L&I or city electrical inspection jurisdiction
  -> L&I elevator/conveyance program if conveyances are present
  -> Applicable WAC Title 51 adoption records
  -> Applicable local amendments, if any
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| **State** | not selected | none | statewide | unknown | pending |
| **County** | not selected | none | statewide | unknown | pending |
| **Municipality** | not selected | none | statewide | unknown | pending |
| **Fire District** | not selected | none | unknown | unknown | pending |
| **Electrical inspection jurisdiction** | not selected | none | statewide / city-specific | unknown | pending |
| **Elevator jurisdiction** | L&I program assumed statewide except local project coordination | src:usa-wa:agency:lni-elevator-laws | statewide | unknown | partial |

### 7.3 AHJ Contact Data

| AHJ ID | Jurisdiction | Department | Role | Website | Phone | Email | Last Verified | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| none populated | none | none | none | none | none | none | null | none |

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Source Type | Publisher | URL | Accessed Date | Snapshot ID | Checksum | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-wa:agency:sbcc-current-codes | State Codes, Regulations & Guidelines | agency_page | Washington State Building Code Council | https://sbcc.wa.gov/state-codes-regulations-guidelines | 2026-06-26 | null | null | official |
| src:usa-wa:agency:sbcc-home-2024-cycle | Washington State Building Code Council homepage / 2024 Codes Delay notice | agency_page | Washington State Building Code Council | https://sbcc.wa.gov/ | 2026-06-26 | null | null | official |
| src:usa-wa:agency:sbcc-building-amendments | Building Code Amendments | agency_page | Washington State Building Code Council | https://sbcc.wa.gov/state-codes-regulations-guidelines/state-building-code/building-code-amendments | 2026-06-26 | null | null | official |
| src:usa-wa:agency:sbcc-residential-amendments | Residential Code Amendments | agency_page | Washington State Building Code Council | https://sbcc.wa.gov/state-codes-regulations-guidelines/state-building-code/residential-code-amendments | 2026-06-26 | null | null | official |
| src:usa-wa:agency:sbcc-mechanical-amendments | Mechanical Code Amendments | agency_page | Washington State Building Code Council | https://sbcc.wa.gov/state-codes-regulations-guidelines/state-building-code/mechanical-code-amendments | 2026-06-26 | null | null | official |
| src:usa-wa:agency:sbcc-fire-amendments | Fire Code Amendments | agency_page | Washington State Building Code Council | https://sbcc.wa.gov/state-codes-regulations-guidelines/state-building-code/fire-code-amendments | 2026-06-26 | null | null | official |
| src:usa-wa:agency:sbcc-energy-code | Energy Code | agency_page | Washington State Building Code Council | https://sbcc.wa.gov/state-codes-regulations-guidelines/state-building-code/energy-code | 2026-06-26 | null | null | official |
| src:usa-wa:agency:lni-electrical-laws | Electrical Laws & Rules | agency_page | Washington State Department of Labor & Industries | https://www.lni.wa.gov/licensing-permits/electrical/laws-rules-policies | 2026-06-26 | null | null | official |
| src:usa-wa:agency:lni-elevator-laws | Elevator Laws, Rules & Policies | agency_page | Washington State Department of Labor & Industries | https://lni.wa.gov/licensing-permits/elevators/laws-rules-policies | 2026-06-26 | null | null | official |
| src:usa-wa:statute:rcw-19-27 | Chapter 19.27 RCW: State Building Code | statute | Washington State Legislature | https://app.leg.wa.gov/RCW/default.aspx?cite=19.27 | 2026-06-26 | null | null | official |
| src:usa-wa:statute:rcw-19-27-050 | RCW 19.27.050: Enforcement | statute | Washington State Legislature | https://app.leg.wa.gov/RCW/default.aspx?cite=19.27.050 | 2026-06-26 | null | null | official |
| src:usa-wa:statute:rcw-19-27-060 | RCW 19.27.060: Local amendments and permit exemptions | statute | Washington State Legislature | https://app.leg.wa.gov/RCW/default.aspx?cite=19.27.060 | 2026-06-26 | null | null | official |
| src:usa-wa:statute:rcw-19-28 | Chapter 19.28 RCW: Electricians and Electrical Installations | statute | Washington State Legislature | https://app.leg.wa.gov/RCW/default.aspx?cite=19.28 | 2026-06-26 | null | null | official |
| src:usa-wa:statute:rcw-19-28-031 | RCW 19.28.031: Rules, Regulations, and Standards | statute | Washington State Legislature | https://app.leg.wa.gov/RCW/default.aspx?cite=19.28.031 | 2026-06-26 | null | null | official |
| src:usa-wa:regulation:wac-title-51 | Title 51 WAC: Enterprise Services, Department of (Building Code Council) | regulation | Washington State Legislature | https://app.leg.wa.gov/WAC/default.aspx?cite=51 | 2026-06-26 | null | null | official |
| src:usa-wa:regulation:wac-51-04 | Chapter 51-04 WAC: Policies and Procedures for Consideration of Statewide and Local Amendments to the State Building Code | regulation | Washington State Legislature | https://app.leg.wa.gov/WAC/default.aspx?cite=51-04 | 2026-06-26 | null | null | official_not_parsed |
| src:usa-wa:regulation:wac-51-50 | Chapter 51-50 WAC: 2021 International Building Code adoption and amendment | regulation | Washington State Legislature | https://app.leg.wa.gov/WAC/default.aspx?cite=51-50 | 2026-06-26 | null | null | official |
| src:usa-wa:regulation:wac-51-51 | Chapter 51-51 WAC: 2021 International Residential Code adoption and amendment | regulation | Washington State Legislature | https://app.leg.wa.gov/WAC/default.aspx?cite=51-51 | 2026-06-26 | null | null | official |
| src:usa-wa:regulation:wac-51-52 | Chapter 51-52 WAC: 2021 International Mechanical Code adoption and amendment | regulation | Washington State Legislature | https://app.leg.wa.gov/WAC/default.aspx?cite=51-52 | 2026-06-26 | null | null | official |
| src:usa-wa:regulation:wac-51-54a | Chapter 51-54A WAC: 2021 International Fire Code adoption and amendment | regulation | Washington State Legislature | https://app.leg.wa.gov/WAC/default.aspx?cite=51-54A | 2026-06-26 | null | null | official |
| src:usa-wa:regulation:wac-51-56 | Chapter 51-56 WAC: 2021 Uniform Plumbing Code adoption and amendment | regulation | Washington State Legislature | https://app.leg.wa.gov/WAC/default.aspx?cite=51-56 | 2026-06-26 | null | null | official |
| src:usa-wa:regulation:wac-51-11c | Chapter 51-11C WAC: 2021 International Energy Conservation Code, Commercial | regulation | Washington State Legislature | https://app.leg.wa.gov/WAC/default.aspx?cite=51-11C | 2026-06-26 | null | null | official |
| src:usa-wa:regulation:wac-51-11r | Chapter 51-11R WAC: 2021 International Energy Conservation Code, Residential | regulation | Washington State Legislature | https://app.leg.wa.gov/WAC/default.aspx?cite=51-11R | 2026-06-26 | null | null | official |
| src:usa-wa:regulation:wac-296-46b | Chapter 296-46B WAC: Electrical Safety Standards, Administration, and Installation | regulation | Washington State Legislature | https://app.leg.wa.gov/WAC/default.aspx?cite=296-46B | 2026-06-26 | null | null | official |
| src:usa-wa:regulation:wac-296-46b-010 | WAC 296-46B-010: General / Adopted Standards | regulation | Washington State Legislature | https://app.leg.wa.gov/WAC/default.aspx?cite=296-46B-010 | 2026-06-26 | null | null | official |
| src:usa-wa:regulation:wac-296-96 | Chapter 296-96 WAC: Safety Regulations and Fees for All Elevators, Dumbwaiters, Escalators and Other Conveyances | regulation | Washington State Legislature | https://app.leg.wa.gov/WAC/default.aspx?cite=296-96 | 2026-06-26 | null | null | official |
| src:usa-wa:regulation:wac-296-96-00650 | WAC 296-96-00650: Adopted Standards | regulation | Washington State Legislature | https://app.leg.wa.gov/WAC/default.aspx?cite=296-96-00650 | 2026-06-26 | null | null | official |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-wa:agency:sbcc-current-codes | agency_summary | SBCC agency pages summarize code editions and amendment publication paths but are not the codified legal text. | Use with WAC chapters for legal anchoring. |
| src:usa-wa:agency:sbcc-home-2024-cycle | dynamic_agency_page | Homepage text may change as the 2024 cycle proceeds. | Monitor monthly or before production use. |
| src:usa-wa:regulation:wac-title-51 | index_page | Title 51 index verifies chapter names and editions, but not all section-level details. | Use individual WAC chapters for field-level dates and exceptions. |
| src:usa-wa:regulation:wac-51-04 | not_parsed | Local amendment procedures chapter identified but not parsed. | Parse before using as a registry or procedural source. |
| src:usa-wa:regulation:wac-296-96-00650 | table_source | Adopted standards are listed in a table-like WAC section. | Review equipment-specific standards before project-level elevator determinations. |
| src:usa-wa:agency:lni-electrical-laws | dynamic_agency_page | L&I rulemaking page includes current rules and future rulemaking timeline; page can change. | Pair with current WAC 296-46B text. |

### 8.3 Supplemental Sources

No supplemental sources were used.

### 8.4 Source Extraction Metadata

| Source ID | Parser | Parser Version | Extracted At | Extraction Confidence | Requires OCR | Requires Browser Automation | Manual Review Required |
| --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-wa:agency:sbcc-current-codes | manual_web_review | 2026-06-26 | 2026-06-26 | 0.90 | no | no | yes |
| src:usa-wa:agency:sbcc-home-2024-cycle | manual_web_review | 2026-06-26 | 2026-06-26 | 0.85 | no | no | yes |
| src:usa-wa:agency:sbcc-building-amendments | manual_web_review | 2026-06-26 | 2026-06-26 | 0.88 | no | no | yes |
| src:usa-wa:agency:sbcc-residential-amendments | manual_web_review | 2026-06-26 | 2026-06-26 | 0.88 | no | no | yes |
| src:usa-wa:agency:sbcc-mechanical-amendments | manual_web_review | 2026-06-26 | 2026-06-26 | 0.88 | no | no | yes |
| src:usa-wa:agency:sbcc-fire-amendments | manual_web_review | 2026-06-26 | 2026-06-26 | 0.84 | no | no | yes |
| src:usa-wa:agency:sbcc-energy-code | manual_web_review | 2026-06-26 | 2026-06-26 | 0.90 | no | no | yes |
| src:usa-wa:agency:lni-electrical-laws | manual_web_review | 2026-06-26 | 2026-06-26 | 0.88 | no | no | yes |
| src:usa-wa:agency:lni-elevator-laws | manual_web_review | 2026-06-26 | 2026-06-26 | 0.86 | no | no | yes |
| src:usa-wa:statute:rcw-19-27 | manual_web_review | 2026-06-26 | 2026-06-26 | 0.88 | no | no | yes |
| src:usa-wa:statute:rcw-19-27-050 | manual_web_review | 2026-06-26 | 2026-06-26 | 0.96 | no | no | no |
| src:usa-wa:statute:rcw-19-27-060 | manual_web_review | 2026-06-26 | 2026-06-26 | 0.92 | no | no | yes |
| src:usa-wa:statute:rcw-19-28 | manual_web_review | 2026-06-26 | 2026-06-26 | 0.88 | no | no | yes |
| src:usa-wa:statute:rcw-19-28-031 | manual_web_review | 2026-06-26 | 2026-06-26 | 0.90 | no | no | yes |
| src:usa-wa:regulation:wac-title-51 | manual_web_review | 2026-06-26 | 2026-06-26 | 0.92 | no | no | yes |
| src:usa-wa:regulation:wac-51-04 | not_extracted | 2026-06-26 | 2026-06-26 | 0.30 | no | no | yes |
| src:usa-wa:regulation:wac-51-50 | manual_web_review | 2026-06-26 | 2026-06-26 | 0.92 | no | no | yes |
| src:usa-wa:regulation:wac-51-51 | manual_web_review | 2026-06-26 | 2026-06-26 | 0.92 | no | no | yes |
| src:usa-wa:regulation:wac-51-52 | manual_web_review | 2026-06-26 | 2026-06-26 | 0.92 | no | no | yes |
| src:usa-wa:regulation:wac-51-54a | manual_web_review | 2026-06-26 | 2026-06-26 | 0.92 | no | no | yes |
| src:usa-wa:regulation:wac-51-56 | manual_web_review | 2026-06-26 | 2026-06-26 | 0.92 | no | no | yes |
| src:usa-wa:regulation:wac-51-11c | manual_web_review | 2026-06-26 | 2026-06-26 | 0.90 | no | no | yes |
| src:usa-wa:regulation:wac-51-11r | manual_web_review | 2026-06-26 | 2026-06-26 | 0.90 | no | no | yes |
| src:usa-wa:regulation:wac-296-46b | manual_web_review | 2026-06-26 | 2026-06-26 | 0.91 | no | no | yes |
| src:usa-wa:regulation:wac-296-46b-010 | manual_web_review | 2026-06-26 | 2026-06-26 | 0.94 | no | no | no |
| src:usa-wa:regulation:wac-296-96 | manual_web_review | 2026-06-26 | 2026-06-26 | 0.86 | no | no | yes |
| src:usa-wa:regulation:wac-296-96-00650 | manual_web_review | 2026-06-26 | 2026-06-26 | 0.88 | no | no | yes |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| adoption:usa-wa:building:2021-ibc | base_model_code.edition_year | 2021 | verified | 0.97 | src:usa-wa:regulation:wac-51-50 | WAC chapter title and adoption section verify edition. |
| adoption:usa-wa:building:2021-ibc | dates.effective_date | 2024-03-15 | verified | 0.96 | src:usa-wa:regulation:wac-51-50 | WAC implementation section gives statewide effective date. |
| adoption:usa-wa:residential:2021-irc | base_model_code.edition_year | 2021 | verified | 0.97 | src:usa-wa:regulation:wac-51-51 | WAC adoption section verifies edition. |
| adoption:usa-wa:residential:2021-irc | exclusions | IRC Chapter 11 and Chapters 25 through 43 are not adopted | verified | 0.90 | src:usa-wa:regulation:wac-51-51; src:usa-wa:agency:sbcc-residential-amendments | Exclusion captured because it affects scope. |
| adoption:usa-wa:mechanical:2021-imc | base_model_code.edition_year | 2021 | verified | 0.96 | src:usa-wa:regulation:wac-51-52 | WAC adoption section verifies edition. |
| adoption:usa-wa:fuel-gas:2021-ifgc | base_model_code.code_name | International Fuel Gas Code; NFPA 58; NFPA 54 | verified | 0.89 | src:usa-wa:regulation:wac-51-52; src:usa-wa:agency:sbcc-mechanical-amendments | Fuel-gas scope and standards captured. |
| adoption:usa-wa:plumbing:2021-upc | base_model_code.edition_year | 2021 | verified | 0.96 | src:usa-wa:regulation:wac-51-56 | WAC adoption section verifies UPC edition and exclusions. |
| adoption:usa-wa:electrical:2023-nec | base_model_code.edition_year | 2023 | verified | 0.96 | src:usa-wa:regulation:wac-296-46b-010 | Current WAC text adopts the 2023 NEC and future 2026 NEC. |
| adoption:usa-wa:electrical:2023-nec | applicability.date_trigger | permit_issuance_date | verified | 0.93 | src:usa-wa:regulation:wac-296-46b-010 | WAC text uses electrical permit issue date. |
| adoption:usa-wa:energy-commercial:2021-wsec | dates.effective_date | 2024-03-15 | verified | 0.94 | src:usa-wa:regulation:wac-51-11c; src:usa-wa:agency:sbcc-energy-code | WAC and SBCC energy page align. |
| adoption:usa-wa:energy-residential:2021-wsec | dates.effective_date | 2024-03-15 | verified | 0.94 | src:usa-wa:regulation:wac-51-11r; src:usa-wa:agency:sbcc-energy-code | WAC and SBCC energy page align. |
| adoption:usa-wa:fire:2021-ifc | base_model_code.edition_year | 2021 | verified | 0.96 | src:usa-wa:regulation:wac-51-54a | WAC adoption section verifies edition. |
| adoption:usa-wa:elevator:2019-asme-a17-1 | base_model_code.edition_year | 2019 | verified | 0.88 | src:usa-wa:regulation:wac-296-96-00650 | Current standard table lists ASME A17.1-2019/CSA B44-19 as current from 2023-10-02. |
| date-rule:usa-wa:sbcc-2021-statewide-effective | trigger_condition | statewide effective in counties and cities; project trigger unresolved | partially_verified | 0.80 | src:usa-wa:regulation:wac-51-50; src:usa-wa:regulation:wac-51-51; src:usa-wa:regulation:wac-51-52; src:usa-wa:regulation:wac-51-54a; src:usa-wa:regulation:wac-51-56; src:usa-wa:regulation:wac-51-11c; src:usa-wa:regulation:wac-51-11r | Effective date verified; permit application/issuance trigger not normalized for building-code chapters. |
| local-enforcement:usa-wa | model | hybrid | partially_verified | 0.88 | src:usa-wa:statute:rcw-19-27-050; src:usa-wa:regulation:wac-296-46b-010; src:usa-wa:agency:lni-elevator-laws | General building local enforcement verified; electrical/elevator specialized authorities verified at a statewide level. |
| local-amendment-rule:usa-wa | model | allowed_with_limits_and_residential_sbc_approval | partially_verified | 0.84 | src:usa-wa:statute:rcw-19-27-060; src:usa-wa:regulation:wac-51-04 | Legal rule verified from statute; procedural registry requires WAC 51-04 parsing. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| **All source IDs resolve** | pass | Source-ID scan found body source IDs in Section 8. |
| **All authority IDs resolve** | pass | Authority IDs used in adoption records are defined in Section 2. |
| **All current code families have adoption records** | pass | Adoption records included for matrix rows. |
| **Building and operational fire code are separated** | pass | Matrix and authority section separate construction references from operational/prevention administration. |
| **Adoption/effective/operative/mandatory dates are not conflated** | pass | Adoption dates remain null unless directly verified; effective/operative/mandatory are separate fields. |
| **Effective dates are valid ISO dates** | pass | Populated date fields use YYYY-MM-DD. |
| **No impossible date sequences** | pass | No replacement dates precede effective dates. |
| **Transition rules have explicit trigger conditions** | partial | Electrical trigger is explicit; building-code trigger still AHJ-sensitive beyond statewide effective date. |
| **Permit-date logic is captured where applicable** | partial | Electrical permit-issue logic captured; building permit application/issuance transition logic remains open. |
| **Local enforcement model classified** | pass | Classified as hybrid with local building enforcement and specialized state programs. |
| **Local amendment rule classified** | partial | Statutory model classified; detailed WAC 51-04 procedure and registry not parsed. |
| **AHJ confirmation metadata present** | fail | No AHJ contacts or local boundary resolution were populated. |
| **Official-source caveats captured** | pass | Caveats table covers agency summaries, index page, dynamic pages, and unparsed procedure chapter. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-wa:001 | high | Building permit transition trigger | WAC implementation sections verify statewide effective dates but do not fully normalize permit application versus permit issuance treatment for building-code chapters. | Parse SBCC rulemaking material and representative AHJ transition notices; determine whether a statewide permit-trigger rule exists. | null | null | open |
| issue:usa-wa:002 | high | Local amendment procedure and registry | RCW 19.27.060 authority is captured, but WAC 51-04 local amendment filing/approval process and any registry were not parsed. | Parse WAC 51-04 and SBCC local amendment materials. | null | null | open |
| issue:usa-wa:003 | medium | Fire operational AHJ scope | IFC adoption is verified, but local fire-code-official administration, operational permits, and fire district/city/county split are not normalized. | Add local fire AHJ resolver and contact data. | null | null | open |
| issue:usa-wa:004 | medium | Accessibility scope | Appendix E and ICC/ANSI A117.1-2017 are captured, but state accessibility statutes, federal ADA overlays, and local enforcement pathways are not parsed. | Parse relevant WAC/RCW accessibility provisions and project-type scoping. | null | null | open |
| issue:usa-wa:005 | medium | Elevator/conveyance detail | Primary L&I elevator authority and ASME A17.1-2019 anchor are captured, but equipment-specific applicability and transition details are not fully normalized. | Parse WAC 296-96 by equipment type and permit/inspection trigger. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-wa:sbcc-2024-cycle | src:usa-wa:agency:sbcc-home-2024-cycle | html_diff | monthly | final adoption date, effective date, or public comment schedule changes | 2026-06-26 | active |
| watch:usa-wa:wac-title-51 | src:usa-wa:regulation:wac-title-51 | chapter_index_diff | monthly | WAC Title 51 chapter title or code-edition changes | 2026-06-26 | active |
| watch:usa-wa:wac-51-50 | src:usa-wa:regulation:wac-51-50 | regulation_text_diff | quarterly | IBC edition, effective date, or amendment text changes | 2026-06-26 | active |
| watch:usa-wa:wac-51-11c | src:usa-wa:regulation:wac-51-11c | regulation_text_diff | quarterly | WSEC commercial edition, effective date, or amendment text changes | 2026-06-26 | active |
| watch:usa-wa:wac-51-11r | src:usa-wa:regulation:wac-51-11r | regulation_text_diff | quarterly | WSEC residential edition, effective date, or amendment text changes | 2026-06-26 | active |
| watch:usa-wa:wac-296-46b | src:usa-wa:regulation:wac-296-46b-010 | regulation_text_diff | monthly_until_2026-12-31 | NEC future replacement date or permit-trigger language changes | 2026-06-26 | active |
| watch:usa-wa:lni-electrical | src:usa-wa:agency:lni-electrical-laws | html_diff | monthly | electrical rulemaking timeline or adopted-rule status changes | 2026-06-26 | active |
| watch:usa-wa:wac-296-96 | src:usa-wa:regulation:wac-296-96-00650 | regulation_text_diff | quarterly | elevator adopted standards table changes | 2026-06-26 | active |
| watch:usa-wa:local-amendments | src:usa-wa:regulation:wac-51-04 | regulation_text_diff | quarterly | local-amendment process or filing/approval requirements change | 2026-06-26 | pending |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-26 | Rebuilt Washington report from the template structure, updated last_updated, populated authority model, source registry, code matrix, date rules, amendment sources, QA, and open issues. | state-report:usa-wa | src:usa-wa:statute:rcw-19-27; src:usa-wa:regulation:wac-title-51; src:usa-wa:regulation:wac-296-46b-010; src:usa-wa:regulation:wac-296-96-00650 | ChatGPT | Status remains partially_verified because AHJ data, local amendment procedures, and building permit-trigger logic still require follow-up. |
