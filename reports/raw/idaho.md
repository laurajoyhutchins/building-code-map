---
state:
  state_id: "US-ID"
  name: "Idaho"
  abbreviation: "ID"
report:
  report_id: "state-report:usa-id"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-24"
  last_verified: "2026-06-24"
  reviewed_by: null
risk:
  overall_confidence: 0.81 # 0.00 - 1.00
  risk_flags:
    - "fire_code_authority_not_verified"
    - "local_amendment_scope_not_verified"
    - "elevator_and_lpg_are_separate_programs"
  open_questions_count: 4

---

# State Building Code Authority Report: Idaho

## 1. Executive Summary

- **Authority model:** Idaho uses a state-administered, board-driven code system under the Division of Occupational and Professional Licenses (DOPL) / Division of Building Safety. The Idaho Building Code Board handles the core building code, while separate electrical, plumbing, HVAC, elevator, and liquefied petroleum gas programs handle trade-specific code families.

- **Statewide code status:** Strong statewide code program for building, residential, existing building, energy, accessibility, mechanical, fuel gas, plumbing, electrical, HVAC, elevator, and LPG safety. No Idaho statewide fire-code source was confirmed in this pass.

- **Local enforcement model:** State-administered. DOPL provides plan review, permits, and inspections for building work, public buildings, school facilities, modular structures, and manufactured homes. A local enforcement agency workflow appears in the modular-building process, but the core code program is state-run.

- **Local amendment posture:** No local amendment registry or municipal code overlay was verified in the sources reviewed. Treat local amendment authority as unresolved until the underlying statute is checked directly.

- **Known transition periods or pending changes:** The current code chapters show recent effective dates on their face, but no separate Idaho-wide transition program was extracted beyond the current rule text and program updates.

- **Production readiness:** usable_with_caution

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: Idaho Building Code Board and DOPL administer the Idaho Building Code Act
    and related rule chapters.
  confidence: 0.98
  source_ids:
  - src:usa-id:bld-home
  - src:usa-id:bld-statutes-rules
  - src:usa-id:bld-rules-2024
- topic: Primary building code edition
  finding: Idaho Building Code - 2018 IBC, with Idaho amendments and accessibility
    appendices.
  confidence: 0.99
  source_ids:
  - src:usa-id:bld-rules-2024
- topic: Electrical code authority
  finding: Idaho Electrical Board adopts the 2023 NEC.
  confidence: 0.99
  source_ids:
  - src:usa-id:ele-home
  - src:usa-id:ele-rules-2025
- topic: Fire code authority
  finding: Not verified in this pass; no statewide Idaho fire-code source was confirmed.
  confidence: 0.1
  source_ids:
  - none
- topic: Local enforcement
  finding: DOPL issues building permits, plan review, and inspections; the building
    rules apply to construction in Idaho and to public entities and modular building
    manufacturers.
  confidence: 0.94
  source_ids:
  - src:usa-id:bld-rules-2024
  - src:usa-id:plan-review
- topic: Local amendments
  finding: No local amendment pathway or registry was verified in the reviewed source
    set.
  confidence: 0.68
  source_ids:
  - src:usa-id:bld-statutes-rules
  - src:usa-id:plan-review
- topic: Effective / operative date rule
  finding: Current rule chapters show effective dates of 2024-07-01 for building,
    2023-03-28 for plumbing/HVAC, and 2025-04-04 for electrical.
  confidence: 0.91
  source_ids:
  - src:usa-id:bld-rules-2024
  - src:usa-id:plb-rules-2023
  - src:usa-id:hvac-rules-2023
  - src:usa-id:ele-rules-2025
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-id:idaho-building-code-board |
| Authority name | Idaho Building Code Board |
| Authority type | state_board |
| Legal basis | Idaho Code Title 39, Chapter 41; Idaho Code 39-4107; 39-4109; 39-4112; 39-4113; 39-9701 |
| Role | Adopts the building code bundle and oversees plan review, permits, and inspections through DOPL |
| Enforcement model | state_administered |
| Source IDs | src:usa-id:bld-home; src:usa-id:bld-statutes-rules; src:usa-id:bld-rules-2024; src:usa-id:plan-review |
| Verification status | verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-id:idaho-building-code-board | Idaho Building Code Board | Adopts the Idaho Building Code - 2018 IBC | Title 39, Chapter 41 / IDAPA 24.39.30 | src:usa-id:bld-rules-2024 | verified |
| Residential | ahj:usa-id:idaho-building-code-board | Idaho Building Code Board | Adopts the Idaho Residential Code - 2018 IRC | Title 39, Chapter 41 / IDAPA 24.39.30 | src:usa-id:bld-rules-2024 | verified |
| Existing Building / Rehabilitation | ahj:usa-id:idaho-building-code-board | Idaho Building Code Board | Adopts the Idaho Existing Building Code - 2018 IEBC | Title 39, Chapter 41 / IDAPA 24.39.30 | src:usa-id:bld-rules-2024 | verified |
| Mechanical | ahj:usa-id:idaho-hvac-board | Idaho HVAC Board | Adopts the Idaho Mechanical Code - 2018 IMC | Title 54, Chapter 50 / IDAPA 24.39.70 | src:usa-id:hvac-home; src:usa-id:hvac-rules-2023 | verified |
| Plumbing | ahj:usa-id:idaho-plumbing-board | State Plumbing Board | Adopts the Idaho State Plumbing Code - 2015 UPC | Title 54, Chapter 26 / IDAPA 24.39.20 | src:usa-id:plb-home; src:usa-id:plb-rules-2023 | verified |
| Fuel Gas | ahj:usa-id:idaho-hvac-board | Idaho HVAC Board | Adopts the 2018 IFGC and IRC Part VI fuel-gas provisions | Title 54, Chapter 50 / IDAPA 24.39.70 | src:usa-id:hvac-home; src:usa-id:hvac-rules-2023 | verified |
| Electrical | ahj:usa-id:idaho-electrical-board | Idaho Electrical Board | Adopts the Idaho Electrical Code - 2023 NEC | Title 54, Chapter 10 / IDAPA 24.39.10 | src:usa-id:ele-home; src:usa-id:ele-rules-2025 | verified |
| Energy | ahj:usa-id:idaho-building-code-board | Idaho Building Code Board | Adopts the Idaho Energy Code - 2018 IECC | Title 39, Chapter 41 / IDAPA 24.39.30 | src:usa-id:bld-rules-2024 | verified |
| Fire - construction references | ahj:usa-id:unknown | Unknown | Not verified | Not verified | none | unresolved |
| Fire - operational / prevention code | ahj:usa-id:unknown | Unknown | Not verified | Not verified | none | unresolved |
| Accessibility | ahj:usa-id:idaho-building-code-board | Idaho Building Code Board | Uses IBC accessibility appendices and related accessibility references; not a separate standalone code in the sources reviewed | Title 39, Chapter 41 / IDAPA 24.39.30 | src:usa-id:bld-rules-2024; src:usa-id:elev-codes | verified |
| Elevator / Conveyance | ahj:usa-id:idaho-elevator-program | Idaho Elevator Program | Adopts the Idaho elevator standards bundle (ASME A17.1, A17.3, A17.4, A17.5, A17.6, A17.7, A17.8, A18.1, QEI-1, and ICC/ANSI A117.1) | DOPL Elevator Program | src:usa-id:elev-home; src:usa-id:elev-codes | verified |

### 2.3 Authority Hierarchy Notes

Idaho is primarily state-administered rather than county-code-overlay driven. The building code board and DOPL own the building-code workflow, while separate trade boards own electrical, plumbing, HVAC/fuel gas, elevator, and LPG safety programs. Factory-built structures are handled through a separate board track inside the same DOPL umbrella.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-id:001 | ahj:usa-id:idaho-building-code-board | delegates_enforcement_to | division_of_occupational_and_professional_licenses | src:usa-id:bld-rules-2024; src:usa-id:plan-review | verified |
| edge:usa-id:002 | ahj:usa-id:idaho-building-code-board | coordinates_with | modular_building_local_enforcement_agency_notice_workflow | src:usa-id:plan-review | verified |
| edge:usa-id:003 | ahj:usa-id:idaho-elevator-program | reserves_review_for | elevator_inspectors_and_program_staff | src:usa-id:elev-home; src:usa-id:elev-codes | verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | Idaho Building Code | IBC | 2018 | current | null | 2024-07-01 | null | null | Current rule chapter adopts the board code bundle. | src:usa-id:bld-rules-2024 |
| Residential | Idaho Residential Code | IRC | 2018 | current | null | 2024-07-01 | null | null | Current rule chapter adopts the residential code bundle. | src:usa-id:bld-rules-2024 |
| Existing Building / Rehabilitation | Idaho Existing Building Code | IEBC | 2018 | current | null | 2024-07-01 | null | null | Current rule chapter adopts the existing-building code bundle. | src:usa-id:bld-rules-2024 |
| Mechanical | Idaho Mechanical Code | IMC | 2018 | current | null | 2023-03-28 | null | null | HVAC rule chapter adopts the mechanical code bundle. | src:usa-id:hvac-rules-2023 |
| Plumbing | Idaho State Plumbing Code | UPC | 2015 | current | null | 2023-03-28 | null | null | Plumbing rule chapter adopts the state plumbing code bundle. | src:usa-id:plb-rules-2023 |
| Fuel Gas | Idaho Mechanical Code / Fuel Gas provisions | IFGC / IRC Part VI | 2018 | current | null | 2023-03-28 | null | null | Fuel-gas provisions are adopted in the HVAC rule chapter. | src:usa-id:hvac-rules-2023 |
| Electrical | Idaho Electrical Code | NEC | 2023 | current | null | 2025-04-04 | null | null | Electrical rule chapter adopts the 2023 NEC. | src:usa-id:ele-rules-2025 |
| Energy | Idaho Energy Code | IECC | 2018 | current | null | 2024-07-01 | null | null | Energy provisions are embedded in the building-code rule chapter. | src:usa-id:bld-rules-2024 |
| Fire - construction references | unknown | unknown | unknown | unresolved | null | null | null | null | No statewide fire-code source was verified in this pass. | none |
| Fire - operational / prevention code | unknown | unknown | unknown | unresolved | null | null | null | null | No statewide fire-code source was verified in this pass. | none |
| Accessibility | Idaho Building Code accessibility appendices | IBC appendices | 2018 | current | null | 2024-07-01 | null | null | Accessibility is embedded in the building-code rule chapter. | src:usa-id:bld-rules-2024 |
| Elevator / Conveyance | Idaho Elevator Program adopted codes | ASME / ICC bundle | mixed | current | null | null | null | null | Current program page lists the adopted standards bundle. | src:usa-id:elev-codes |

### 3.2 Adoption Records

```yaml
adoption_id: "adoption:usa-id:building:idaho-building-code-2018"
state_id: "US-ID"
code_family: "building"
status: "current"
state_code:
  name: "Idaho Building Code"
  edition_label: "2018 IBC"
  codification: "Idaho Code Title 39, Chapter 41 / IDAPA 24.39.30"
base_model_code:
  publisher: "ICC"
  code_name: "IBC"
  edition_year: 2018
  incorporated_by_reference: true
authority:
  adopting_authority_id: "ahj:usa-id:idaho-building-code-board"
  enforcing_authority_model: "state"
  interpretation_authority_id: "ahj:usa-id:idaho-building-code-board"
dates:
  adoption_date: null
  effective_date: "2024-07-01"
  operative_date: null
  mandatory_date: null
  replacement_date: null
applicability:
  date_trigger: "permit_application_date"
  applies_to:
    - "new_construction"
    - "alteration"
    - "addition"
    - "public_buildings"
    - "school_facilities"
  exclusions: []
  special_conditions:
    - "Includes plan review and permit workflows administered by DOPL."
transition:
  exists: false
  rule_id: null
  start_date: null
  end_date: null
  prior_code_allowed: false
  prior_code_condition: "Not separately extracted in this pass."
amendments:
  state_amended: true
  amendment_set_ids:
    - "amendment-set:usa-id:building:2024"
  amendment_source_ids:
    - "src:usa-id:bld-rules-2024"
provenance:
  source_ids:
    - "src:usa-id:bld-rules-2024"
  field_sources:
    state_code.name: ["src:usa-id:bld-rules-2024"]
    state_code.codification: ["src:usa-id:bld-rules-2024"]
    authority.adopting_authority_id: ["src:usa-id:bld-home", "src:usa-id:bld-statutes-rules"]
verification:
  status: "verified"
  confidence: 0.99
  notes: "Rule chapter explicitly lists the 2018 IBC and the building-code board authority."

adoption_id: "adoption:usa-id:residential:idaho-residential-code-2018"
state_id: "US-ID"
code_family: "residential"
status: "current"
state_code:
  name: "Idaho Residential Code"
  edition_label: "2018 IRC"
  codification: "Idaho Code Title 39, Chapter 41 / IDAPA 24.39.30"
base_model_code:
  publisher: "ICC"
  code_name: "IRC"
  edition_year: 2018
  incorporated_by_reference: true
authority:
  adopting_authority_id: "ahj:usa-id:idaho-building-code-board"
  enforcing_authority_model: "state"
  interpretation_authority_id: "ahj:usa-id:idaho-building-code-board"
dates:
  adoption_date: null
  effective_date: "2024-07-01"
  operative_date: null
  mandatory_date: null
  replacement_date: null
applicability:
  date_trigger: "permit_application_date"
  applies_to:
    - "one_family"
    - "two_family"
    - "townhouse"
  exclusions: []
  special_conditions: []
transition:
  exists: false
  rule_id: null
  start_date: null
  end_date: null
  prior_code_allowed: false
  prior_code_condition: "Not separately extracted in this pass."
amendments:
  state_amended: true
  amendment_set_ids:
    - "amendment-set:usa-id:residential:2024"
  amendment_source_ids:
    - "src:usa-id:bld-rules-2024"
provenance:
  source_ids:
    - "src:usa-id:bld-rules-2024"
  field_sources:
    state_code.name: ["src:usa-id:bld-rules-2024"]
    state_code.codification: ["src:usa-id:bld-rules-2024"]
verification:
  status: "verified"
  confidence: 0.98
  notes: "Residential code is embedded in the building-code rule chapter."

adoption_id: "adoption:usa-id:existing-building:idaho-existing-building-code-2018"
state_id: "US-ID"
code_family: "existing_building"
status: "current"
state_code:
  name: "Idaho Existing Building Code"
  edition_label: "2018 IEBC"
  codification: "Idaho Code Title 39, Chapter 41 / IDAPA 24.39.30"
base_model_code:
  publisher: "ICC"
  code_name: "IEBC"
  edition_year: 2018
  incorporated_by_reference: true
authority:
  adopting_authority_id: "ahj:usa-id:idaho-building-code-board"
  enforcing_authority_model: "state"
  interpretation_authority_id: "ahj:usa-id:idaho-building-code-board"
dates:
  adoption_date: null
  effective_date: "2024-07-01"
  operative_date: null
  mandatory_date: null
  replacement_date: null
applicability:
  date_trigger: "permit_application_date"
  applies_to:
    - "existing_building"
    - "rehabilitation"
    - "alteration"
  exclusions: []
  special_conditions: []
transition:
  exists: false
  rule_id: null
  start_date: null
  end_date: null
  prior_code_allowed: false
  prior_code_condition: "Not separately extracted in this pass."
amendments:
  state_amended: true
  amendment_set_ids:
    - "amendment-set:usa-id:existing-building:2024"
  amendment_source_ids:
    - "src:usa-id:bld-rules-2024"
provenance:
  source_ids:
    - "src:usa-id:bld-rules-2024"
  field_sources:
    state_code.name: ["src:usa-id:bld-rules-2024"]
verification:
  status: "verified"
  confidence: 0.98
  notes: "Rule chapter explicitly lists the 2018 IEBC."

adoption_id: "adoption:usa-id:energy:idaho-energy-code-2018"
state_id: "US-ID"
code_family: "energy"
status: "current"
state_code:
  name: "Idaho Energy Code"
  edition_label: "2018 IECC"
  codification: "Idaho Code Title 39, Chapter 41 / IDAPA 24.39.30"
base_model_code:
  publisher: "ICC"
  code_name: "IECC"
  edition_year: 2018
  incorporated_by_reference: true
authority:
  adopting_authority_id: "ahj:usa-id:idaho-building-code-board"
  enforcing_authority_model: "state"
  interpretation_authority_id: "ahj:usa-id:idaho-building-code-board"
dates:
  adoption_date: null
  effective_date: "2024-07-01"
  operative_date: null
  mandatory_date: null
  replacement_date: null
applicability:
  date_trigger: "permit_application_date"
  applies_to:
    - "new_construction"
    - "commercial"
    - "residential"
  exclusions: []
  special_conditions:
    - "Commercial and residential energy provisions are embedded in the building-code rule chapter."
transition:
  exists: false
  rule_id: null
  start_date: null
  end_date: null
  prior_code_allowed: false
  prior_code_condition: "Not separately extracted in this pass."
amendments:
  state_amended: true
  amendment_set_ids:
    - "amendment-set:usa-id:energy:2024"
  amendment_source_ids:
    - "src:usa-id:bld-rules-2024"
provenance:
  source_ids:
    - "src:usa-id:bld-rules-2024"
  field_sources:
    state_code.name: ["src:usa-id:bld-rules-2024"]
verification:
  status: "verified"
  confidence: 0.97
  notes: "Energy code appears in the building-code rule chapter."

adoption_id: "adoption:usa-id:accessibility:idaho-accessibility-framework-2018"
state_id: "US-ID"
code_family: "accessibility"
status: "current"
state_code:
  name: "Idaho Building Code accessibility appendices"
  edition_label: "2018 IBC appendices / accessibility references"
  codification: "Idaho Code Title 39, Chapter 41 / IDAPA 24.39.30"
base_model_code:
  publisher: "ICC"
  code_name: "IBC appendices and accessibility references"
  edition_year: 2018
  incorporated_by_reference: true
authority:
  adopting_authority_id: "ahj:usa-id:idaho-building-code-board"
  enforcing_authority_model: "state"
  interpretation_authority_id: "ahj:usa-id:idaho-building-code-board"
dates:
  adoption_date: null
  effective_date: "2024-07-01"
  operative_date: null
  mandatory_date: null
  replacement_date: null
applicability:
  date_trigger: "permit_application_date"
  applies_to:
    - "new_construction"
    - "alteration"
    - "public_buildings"
  exclusions: []
  special_conditions:
    - "Accessibility is not a standalone code in the sources reviewed."
transition:
  exists: false
  rule_id: null
  start_date: null
  end_date: null
  prior_code_allowed: false
  prior_code_condition: "Not separately extracted in this pass."
amendments:
  state_amended: true
  amendment_set_ids:
    - "amendment-set:usa-id:accessibility:2024"
  amendment_source_ids:
    - "src:usa-id:bld-rules-2024"
provenance:
  source_ids:
    - "src:usa-id:bld-rules-2024"
  field_sources:
    state_code.name: ["src:usa-id:bld-rules-2024"]
verification:
  status: "verified"
  confidence: 0.94
  notes: "Building rules reference accessibility appendices rather than a separate accessibility code."

adoption_id: "adoption:usa-id:mechanical:idaho-mechanical-code-2018"
state_id: "US-ID"
code_family: "mechanical"
status: "current"
state_code:
  name: "Idaho Mechanical Code"
  edition_label: "2018 IMC"
  codification: "Idaho Code Title 54, Chapter 50 / IDAPA 24.39.70"
base_model_code:
  publisher: "ICC"
  code_name: "IMC"
  edition_year: 2018
  incorporated_by_reference: true
authority:
  adopting_authority_id: "ahj:usa-id:idaho-hvac-board"
  enforcing_authority_model: "state"
  interpretation_authority_id: "ahj:usa-id:idaho-hvac-board"
dates:
  adoption_date: null
  effective_date: "2023-03-28"
  operative_date: null
  mandatory_date: null
  replacement_date: null
applicability:
  date_trigger: "permit_application_date"
  applies_to:
    - "hvac_installations"
    - "commercial"
    - "residential"
  exclusions: []
  special_conditions: []
transition:
  exists: false
  rule_id: null
  start_date: null
  end_date: null
  prior_code_allowed: false
  prior_code_condition: "Not separately extracted in this pass."
amendments:
  state_amended: true
  amendment_set_ids:
    - "amendment-set:usa-id:mechanical:2023"
  amendment_source_ids:
    - "src:usa-id:hvac-rules-2023"
provenance:
  source_ids:
    - "src:usa-id:hvac-rules-2023"
  field_sources:
    state_code.name: ["src:usa-id:hvac-rules-2023"]
verification:
  status: "verified"
  confidence: 0.97
  notes: "HVAC rule chapter explicitly states the 2018 IMC adoption."

adoption_id: "adoption:usa-id:fuel-gas:idaho-fuel-gas-provisions-2018"
state_id: "US-ID"
code_family: "fuel_gas"
status: "current"
state_code:
  name: "Idaho Mechanical Code / Fuel Gas provisions"
  edition_label: "2018 IFGC and IRC Part VI"
  codification: "Idaho Code Title 54, Chapter 50 / IDAPA 24.39.70"
base_model_code:
  publisher: "ICC"
  code_name: "IFGC / IRC Part VI"
  edition_year: 2018
  incorporated_by_reference: true
authority:
  adopting_authority_id: "ahj:usa-id:idaho-hvac-board"
  enforcing_authority_model: "state"
  interpretation_authority_id: "ahj:usa-id:idaho-hvac-board"
dates:
  adoption_date: null
  effective_date: "2023-03-28"
  operative_date: null
  mandatory_date: null
  replacement_date: null
applicability:
  date_trigger: "permit_application_date"
  applies_to:
    - "fuel_gas_piping"
    - "hearth_appliances"
    - "nonducted_oil_furnaces"
  exclusions:
    - "plumbing"
    - "electrical"
    - "ductwork"
  special_conditions:
    - "Specialty HVAC scope is limited to the fuel-gas subset described in the rule chapter."
transition:
  exists: false
  rule_id: null
  start_date: null
  end_date: null
  prior_code_allowed: false
  prior_code_condition: "Not separately extracted in this pass."
amendments:
  state_amended: true
  amendment_set_ids:
    - "amendment-set:usa-id:fuel-gas:2023"
  amendment_source_ids:
    - "src:usa-id:hvac-rules-2023"
provenance:
  source_ids:
    - "src:usa-id:hvac-rules-2023"
  field_sources:
    state_code.name: ["src:usa-id:hvac-rules-2023"]
verification:
  status: "verified"
  confidence: 0.96
  notes: "HVAC rule chapter explicitly lists fuel-gas adoption and the specialty scope."

adoption_id: "adoption:usa-id:plumbing:idaho-state-plumbing-code-2015"
state_id: "US-ID"
code_family: "plumbing"
status: "current"
state_code:
  name: "Idaho State Plumbing Code"
  edition_label: "2015 UPC"
  codification: "Idaho Code Title 54, Chapter 26 / IDAPA 24.39.20"
base_model_code:
  publisher: "IAPMO"
  code_name: "UPC"
  edition_year: 2015
  incorporated_by_reference: true
authority:
  adopting_authority_id: "ahj:usa-id:idaho-plumbing-board"
  enforcing_authority_model: "state"
  interpretation_authority_id: "ahj:usa-id:idaho-plumbing-board"
dates:
  adoption_date: null
  effective_date: "2023-03-28"
  operative_date: null
  mandatory_date: null
  replacement_date: null
applicability:
  date_trigger: "permit_application_date"
  applies_to:
    - "plumbing_installations"
    - "commercial"
    - "residential"
  exclusions: []
  special_conditions: []
transition:
  exists: false
  rule_id: null
  start_date: null
  end_date: null
  prior_code_allowed: false
  prior_code_condition: "Not separately extracted in this pass."
amendments:
  state_amended: true
  amendment_set_ids:
    - "amendment-set:usa-id:plumbing:2023"
  amendment_source_ids:
    - "src:usa-id:plb-rules-2023"
provenance:
  source_ids:
    - "src:usa-id:plb-rules-2023"
  field_sources:
    state_code.name: ["src:usa-id:plb-home", "src:usa-id:plb-rules-2023"]
verification:
  status: "verified"
  confidence: 0.98
  notes: "Plumbing board page and rule chapter together identify the state plumbing code."

adoption_id: "adoption:usa-id:electrical:idaho-electrical-code-2023"
state_id: "US-ID"
code_family: "electrical"
status: "current"
state_code:
  name: "Idaho Electrical Code"
  edition_label: "2023 NEC"
  codification: "Idaho Code Title 54, Chapter 10 / IDAPA 24.39.10"
base_model_code:
  publisher: "NFPA"
  code_name: "NEC"
  edition_year: 2023
  incorporated_by_reference: true
authority:
  adopting_authority_id: "ahj:usa-id:idaho-electrical-board"
  enforcing_authority_model: "state"
  interpretation_authority_id: "ahj:usa-id:idaho-electrical-board"
dates:
  adoption_date: null
  effective_date: "2025-04-04"
  operative_date: null
  mandatory_date: null
  replacement_date: null
applicability:
  date_trigger: "permit_application_date"
  applies_to:
    - "electrical_installations"
    - "commercial"
    - "residential"
  exclusions: []
  special_conditions: []
transition:
  exists: false
  rule_id: null
  start_date: null
  end_date: null
  prior_code_allowed: false
  prior_code_condition: "Not separately extracted in this pass."
amendments:
  state_amended: true
  amendment_set_ids:
    - "amendment-set:usa-id:electrical:2025"
  amendment_source_ids:
    - "src:usa-id:ele-rules-2025"
provenance:
  source_ids:
    - "src:usa-id:ele-rules-2025"
  field_sources:
    state_code.name: ["src:usa-id:ele-rules-2025"]
verification:
  status: "verified"
  confidence: 0.99
  notes: "Electrical board rule chapter explicitly adopts the 2023 NEC."

adoption_id: "adoption:usa-id:elevator:idaho-elevator-adopted-codes"
state_id: "US-ID"
code_family: "elevator"
status: "current"
state_code:
  name: "Idaho Elevator Program adopted codes"
  edition_label: "ASME A17.1 2022 and related bundle"
  codification: "DOPL Elevator Program"
base_model_code:
  publisher: "ASME / ICC"
  code_name: "ASME A17.1, A17.3, A17.4, A17.5, A17.6, A17.7, A17.8, A18.1, QEI-1, ICC/ANSI A117.1"
  edition_year: 2022
  incorporated_by_reference: true
authority:
  adopting_authority_id: "ahj:usa-id:idaho-elevator-program"
  enforcing_authority_model: "state"
  interpretation_authority_id: "ahj:usa-id:idaho-elevator-program"
dates:
  adoption_date: null
  effective_date: null
  operative_date: null
  mandatory_date: null
  replacement_date: null
applicability:
  date_trigger: "certificate_to_operate_or_inspection_date"
  applies_to:
    - "elevators"
    - "escalators"
    - "moving_walks"
    - "platform_lifts"
    - "chairlifts"
    - "dumbwaiters"
  exclusions: []
  special_conditions:
    - "The program page lists a bundled set of adopted standards rather than a single model code."
transition:
  exists: false
  rule_id: null
  start_date: null
  end_date: null
  prior_code_allowed: false
  prior_code_condition: "Not separately extracted in this pass."
amendments:
  state_amended: true
  amendment_set_ids:
    - "amendment-set:usa-id:elevator:bundle"
  amendment_source_ids:
    - "src:usa-id:elev-codes"
provenance:
  source_ids:
    - "src:usa-id:elev-codes"
    - "src:usa-id:elev-home"
  field_sources:
    state_code.name: ["src:usa-id:elev-codes"]
    state_code.edition_label: ["src:usa-id:elev-codes"]
verification:
  status: "verified"
  confidence: 0.95
  notes: "Program page lists the adopted standards bundle directly."
```

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

Idaho's current code structure is set through board and program rule chapters under DOPL. The building-code chapter is current as of 2024-07-01, the plumbing and HVAC chapters are current as of 2023-03-28, and the electrical chapter is current as of 2025-04-04. The elevator program publishes a bundle of adopted standards, but a separate effective-date transition was not extracted here.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-id:001 | building and related code families | effective_date | 2024-07-01 | Building-code rule chapter became current | false | src:usa-id:bld-rules-2024 | verified |
| date-rule:usa-id:002 | mechanical, fuel gas, plumbing | effective_date | 2023-03-28 | HVAC and plumbing rule chapters became current | false | src:usa-id:hvac-rules-2023; src:usa-id:plb-rules-2023 | verified |
| date-rule:usa-id:003 | electrical | effective_date | 2025-04-04 | Electrical rule chapter became current | false | src:usa-id:ele-rules-2025 | verified |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| building | rule chapter update | null | null | null | null | null | monitoring | src:usa-id:bld-home; src:usa-id:bld-rules-2024 | Watch for board rulemaking and updated code editions. |
| electrical | NEC cycle update | null | null | null | null | null | monitoring | src:usa-id:ele-home; src:usa-id:ele-rules-2025 | Watch for the next NEC cycle. |
| plumbing | UPC cycle update | null | null | null | null | null | monitoring | src:usa-id:plb-home; src:usa-id:plb-rules-2023 | Watch for any move off the 2015 UPC. |
| hvac / fuel_gas | IMC / IFGC cycle update | null | null | null | null | null | monitoring | src:usa-id:hvac-home; src:usa-id:hvac-rules-2023 | Watch for HVAC board amendments. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-id:001 | building | public buildings and school facilities | DOPL review / inspection workflow | The building rule chapter explicitly covers public buildings and school facilities. | src:usa-id:bld-rules-2024 | verified |
| applicability-rule:usa-id:002 | building | modular structures and manufactured homes | modular permit workflow | Modular structures and manufactured homes are handled in the building and factory-built programs. | src:usa-id:bld-rules-2024; src:usa-id:plan-review | verified |
| applicability-rule:usa-id:003 | hvac | specialty fuel-gas scope | specialty certificate scope | Specialty HVAC work can be limited to fuel-gas piping and designated appliances. | src:usa-id:hvac-rules-2023 | verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** board and program rule chapters adopted under DOPL, with separate chapters for building, electrical, plumbing, HVAC, elevator, and LPG safety

**Where amendments are published:** Idaho Administrative Code, DOPL board pages, and board-referenced documents

**Amendment parsing status:** partial

### 5.2 State Amendment Sources

| Amendment Set ID | Code Family | Publication Format | Source IDs | Parsed? | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| amendment-set:usa-id:building:2024 | building / energy / accessibility | DOPL building rule chapter | src:usa-id:bld-rules-2024 | yes | Building, residential, existing-building, energy, and accessibility references are in one rule chapter. |
| amendment-set:usa-id:mechanical:2023 | mechanical | HVAC rule chapter | src:usa-id:hvac-rules-2023 | yes | IMC 2018 adopted with amendments. |
| amendment-set:usa-id:fuel-gas:2023 | fuel_gas | HVAC rule chapter | src:usa-id:hvac-rules-2023 | yes | IFGC and IRC Part VI fuel-gas provisions adopted with amendments. |
| amendment-set:usa-id:plumbing:2023 | plumbing | plumbing rule chapter | src:usa-id:plb-rules-2023 | yes | 2015 UPC adopted as the Idaho State Plumbing Code. |
| amendment-set:usa-id:electrical:2025 | electrical | electrical rule chapter | src:usa-id:ele-rules-2025 | yes | 2023 NEC adopted with amendments. |
| amendment-set:usa-id:elevator:bundle | elevator | elevator adopted-codes page | src:usa-id:elev-codes | partial | Program page lists a bundled set of adopted standards. |
| amendment-set:usa-id:lpg:2017 | lpg_safety | board-referenced documents page | src:usa-id:lpg-documents | partial | NFPA 58 2017 edition is identified on the board documents page. |

### 5.3 High-Impact State Amendments

| Code Family | Section | Amendment Type | Summary | Source IDs | Confidence |
| --- | --- | --- | --- | --- | --- |
| building | 600 / accessibility references | adopt | The building rule chapter adopts the 2018 IBC and includes accessibility appendices/references. | src:usa-id:bld-rules-2024 | 0.98 |
| building | 600 / energy references | adopt | The building rule chapter adopts the 2018 IECC as Idaho's energy code. | src:usa-id:bld-rules-2024 | 0.97 |
| mechanical | 600 | adopt | The HVAC rule chapter adopts the 2018 IMC. | src:usa-id:hvac-rules-2023 | 0.97 |
| fuel_gas | 600 | adopt | The HVAC rule chapter adopts the 2018 IFGC and IRC Part VI fuel-gas provisions. | src:usa-id:hvac-rules-2023 | 0.96 |
| plumbing | 600 | adopt | The plumbing rule chapter adopts the 2015 UPC as the Idaho State Plumbing Code. | src:usa-id:plb-rules-2023 | 0.98 |
| electrical | 600 | adopt | The electrical rule chapter adopts the 2023 NEC. | src:usa-id:ele-rules-2025 | 0.99 |
| elevator | adopted codes list | adopt | The elevator program lists ASME and ICC/ANSI standards for conveyance safety and accessibility. | src:usa-id:elev-codes | 0.95 |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-id"
  model: "state_administered"
  enforcing_entities:
    - "division_of_occupational_and_professional_licenses"
    - "division_of_building_safety"
    - "trade_board_inspector_programs"
  required_officials:
    - "building_inspector"
    - "plan_reviewer"
    - "trade_inspector"
  state_reserved_activities:
    - "code_adoption"
    - "plan_review"
    - "permit_issuance"
    - "inspection_services"
  source_ids:
    - "src:usa-id:bld-rules-2024"
    - "src:usa-id:plan-review"
    - "src:usa-id:ele-home"
    - "src:usa-id:plb-home"
    - "src:usa-id:hvac-home"
  verification_status: "verified"
  confidence: 0.94
```

### 6.2 Local Amendment Rule

local_amendment_rule:
  rule_id: "local-amendment-rule:usa-id"
  model: "state_uniform"
  applies_to_code_families:
    - "building"
    - "residential"
    - "existing_building"
    - "mechanical"
    - "plumbing"
    - "fuel_gas"
    - "electrical"
    - "energy"
    - "elevator"
  approval_required: false
  approving_authority_id: null
  filing_required: false
  registry_exists: false
  registry_source_ids:
    - "src:usa-id:bld-statutes-rules"
  legal_basis_source_ids:
    - "src:usa-id:bld-rules-2024"
  verification_status: "partially_verified"
  confidence: 0.68

### 6.3 Local Enforcement vs. Local Amendment Summary

Idaho's code program is centralized at the state level. DOPL and its boards run permits, plan review, and inspections, while the modular-building workflow introduces a local enforcement agency notice step. I did not verify a local municipal amendment regime in the source set reviewed here.

### 6.4 Known Local Amendment Registries

| Registry ID | Scope | Maintained By | Source ID | Machine Readable? | Parsed? | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| registry:usa-id:local-amendments | statewide | not verified | none | no | no | No statewide local-amendment registry was identified in the sources reviewed. |

### 6.5 Municipality-Specific Known Amendments

| Jurisdiction | Code Family | Amendment Set ID | Approval Status | Effective Date | Source IDs | Parsed? |
| --- | --- | --- | --- | --- | --- | --- |
| none verified | none | none | none | null | none | no |

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: state_administered

Jurisdiction stack:

```text
Address
  -> State
  -> County, only where the modular workflow references a local enforcement agency
  -> Municipality / unincorporated county
  -> Special districts, if applicable
  -> Building AHJ
  -> Fire AHJ
  -> Trade-specific AHJs
  -> Applicable state code adoption records
  -> Applicable local amendment records
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- | --- |
| State | DOPL / Division of Building Safety | src:usa-id:bld-home | statewide | periodic | active |
| County | modular-building local enforcement agency workflow | src:usa-id:plan-review | limited | periodic | partial |
| Municipality | not separately identified | none | unknown | unknown | pending |
| Fire District | not verified | none | unknown | unknown | pending |
| Special District | not verified | none | unknown | unknown | pending |

### 7.3 AHJ Contact Data

| AHJ ID | Jurisdiction | Department | Role | Website | Phone | Email | Last Verified | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| ahj:usa-id:building-code | Statewide | Idaho Building Code Board / DOPL | building | https://dopl.idaho.gov/bld/ | 208-334-3233 | BCRE-PERMITS@DOPL.IDAHO.GOV | 2026-06-24 | src:usa-id:bld-home; src:usa-id:plan-review |
| ahj:usa-id:electrical-board | Statewide | Idaho Electrical Board / DOPL | electrical | https://dopl.idaho.gov/ele/ | 208-334-3233 | BCRE-Licensing@DOPL.IDAHO.GOV | 2026-06-24 | src:usa-id:ele-home |
| ahj:usa-id:plumbing-board | Statewide | State Plumbing Board / DOPL | plumbing | https://dopl.idaho.gov/plb/ | 208-334-3233 | BCRE-Licensing@DOPL.IDAHO.GOV | 2026-06-24 | src:usa-id:plb-home |
| ahj:usa-id:hvac-board | Statewide | HVAC Board / DOPL | mechanical / fuel gas | https://dopl.idaho.gov/hvac/ | 208-334-3233 | BCRE-Licensing@DOPL.IDAHO.GOV | 2026-06-24 | src:usa-id:hvac-home |
| ahj:usa-id:elevator-program | Statewide | Idaho Elevator Program / DOPL | elevators | https://dopl.idaho.gov/elev/ | 208-334-3233 | safety@dopl.idaho.gov | 2026-06-24 | src:usa-id:elev-home; src:usa-id:elev-codes |
| ahj:usa-id:lpg-safety | Statewide | Board of Liquefied Petroleum Gas Safety / DOPL | LPG safety | https://dopl.idaho.gov/lpg/ | 208-334-3233 | BCRE-Licensing@DOPL.IDAHO.GOV | 2026-06-24 | src:usa-id:lpg-home; src:usa-id:lpg-documents |

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Source Type | Publisher | URL | Accessed Date | Snapshot ID | Checksum | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-id:bld-home | Idaho Building Code Board & Factory Built Structures Board | agency_page | DOPL | https://dopl.idaho.gov/bld/ | 2026-06-24 | snapshot:usa-id-bld-home | null | active |
| src:usa-id:bld-statutes-rules | Statutes and Rules | agency_page | DOPL | https://dopl.idaho.gov/bld/bld-statutes-and-rules/ | 2026-06-24 | snapshot:usa-id-bld-statutes-rules | null | active |
| src:usa-id:bld-rules-2024 | Rules of Building Safety (Building Code Rules) | regulation_pdf | Idaho Administrative Code | https://adminrules.idaho.gov/rules/current/24/243930.pdf | 2026-06-24 | snapshot:usa-id-bld-rules-2024 | null | active |
| src:usa-id:plan-review | Plan Review and Permits | agency_page | DOPL | https://dopl.idaho.gov/bld/bld-plan-review-and-permits/ | 2026-06-24 | snapshot:usa-id-plan-review | null | active |
| src:usa-id:ele-home | Idaho Electrical Board | agency_page | DOPL | https://dopl.idaho.gov/ele/ | 2026-06-24 | snapshot:usa-id-ele-home | null | active |
| src:usa-id:ele-rules-2025 | Rules of the Idaho Electrical Board | regulation_pdf | Idaho Administrative Code | https://adminrules.idaho.gov/rules/current/24/243910.pdf | 2026-06-24 | snapshot:usa-id-ele-rules-2025 | null | active |
| src:usa-id:plb-home | State Plumbing Board | agency_page | DOPL | https://dopl.idaho.gov/plb/ | 2026-06-24 | snapshot:usa-id-plb-home | null | active |
| src:usa-id:plb-rules-2023 | Rules Governing Plumbing | regulation_pdf | Idaho Administrative Code | https://adminrules.idaho.gov/rules/current/24/243920.pdf | 2026-06-24 | snapshot:usa-id-plb-rules-2023 | null | active |
| src:usa-id:hvac-home | Heating, Ventilation and Air Conditioning (HVAC) Board | agency_page | DOPL | https://dopl.idaho.gov/hvac/ | 2026-06-24 | snapshot:usa-id-hvac-home | null | active |
| src:usa-id:hvac-rules-2023 | Rules Governing Installation of Heating, Ventilation, and Air Conditioning Systems | regulation_pdf | Idaho Administrative Code | https://adminrules.idaho.gov/rules/current/24/243970.pdf | 2026-06-24 | snapshot:usa-id-hvac-rules-2023 | null | active |
| src:usa-id:elev-home | Idaho Elevator Program | agency_page | DOPL | https://dopl.idaho.gov/elev/ | 2026-06-24 | snapshot:usa-id-elev-home | null | active |
| src:usa-id:elev-codes | Adopted Codes | agency_page | DOPL | https://dopl.idaho.gov/elev/elev-adopted-codes/ | 2026-06-24 | snapshot:usa-id-elev-codes | null | active |
| src:usa-id:lpg-home | Board of Liquefied Petroleum Gas Safety | agency_page | DOPL | https://dopl.idaho.gov/lpg/ | 2026-06-24 | snapshot:usa-id-lpg-home | null | active |
| src:usa-id:lpg-documents | Board Referenced Documents | agency_page | DOPL | https://dopl.idaho.gov/lpg/lpg-documents/ | 2026-06-24 | snapshot:usa-id-lpg-documents | null | active |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-id:bld-rules-2024 | rule_chapter | The building rule chapter is the canonical source for the 2018 IBC/IRC/IEBC/IECC bundle, but it does not separately resolve fire-code authority. | use for building-family adoption and permit/inspection rules; do not infer fire-code adoption |
| src:usa-id:elev-codes | standards_bundle | The elevator program publishes a standards bundle rather than a single model code. | use as the authoritative adopted bundle and keep the standards list explicit |
| src:usa-id:lpg-documents | standards_bundle | The LPG page identifies NFPA 58, 2017 edition, for public inspection. | use for LPG safety program modeling; keep it separate from HVAC fuel-gas adoption |

### 8.3 Supplemental Sources

| Source ID | Title | Source Type | Publisher | URL | Use | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| none | none | none | none | none | none | none |

### 8.4 Source Extraction Metadata

| Source ID | Parser | Parser Version | Extracted At | Extraction Confidence | Requires OCR | Requires Browser Automation | Manual Review Required |
| --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-id:bld-home | web | 2026-06-24 | 2026-06-24T00:00:00Z | 0.98 | no | no | no |
| src:usa-id:bld-rules-2024 | web | 2026-06-24 | 2026-06-24T00:00:00Z | 0.99 | no | no | no |
| src:usa-id:plan-review | web | 2026-06-24 | 2026-06-24T00:00:00Z | 0.97 | no | no | no |
| src:usa-id:ele-home | web | 2026-06-24 | 2026-06-24T00:00:00Z | 0.98 | no | no | no |
| src:usa-id:ele-rules-2025 | web | 2026-06-24 | 2026-06-24T00:00:00Z | 0.99 | no | no | no |
| src:usa-id:plb-home | web | 2026-06-24 | 2026-06-24T00:00:00Z | 0.97 | no | no | no |
| src:usa-id:plb-rules-2023 | web | 2026-06-24 | 2026-06-24T00:00:00Z | 0.98 | no | no | no |
| src:usa-id:hvac-home | web | 2026-06-24 | 2026-06-24T00:00:00Z | 0.98 | no | no | no |
| src:usa-id:hvac-rules-2023 | web | 2026-06-24 | 2026-06-24T00:00:00Z | 0.99 | no | no | no |
| src:usa-id:elev-home | web | 2026-06-24 | 2026-06-24T00:00:00Z | 0.96 | no | no | no |
| src:usa-id:elev-codes | web | 2026-06-24 | 2026-06-24T00:00:00Z | 0.95 | no | no | no |
| src:usa-id:lpg-home | web | 2026-06-24 | 2026-06-24T00:00:00Z | 0.96 | no | no | no |
| src:usa-id:lpg-documents | web | 2026-06-24 | 2026-06-24T00:00:00Z | 0.95 | no | no | no |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| adoption:usa-id:building:idaho-building-code-2018 | state_code.name | Idaho Building Code | verified | 0.99 | src:usa-id:bld-rules-2024 | Current building code. |
| adoption:usa-id:electrical:idaho-electrical-code-2023 | state_code.name | Idaho Electrical Code | verified | 0.99 | src:usa-id:ele-rules-2025 | Current electrical code. |
| adoption:usa-id:plumbing:idaho-state-plumbing-code-2015 | state_code.name | Idaho State Plumbing Code | verified | 0.98 | src:usa-id:plb-rules-2023 | Current plumbing code. |
| adoption:usa-id:mechanical:idaho-mechanical-code-2018 | state_code.name | Idaho Mechanical Code | verified | 0.97 | src:usa-id:hvac-rules-2023 | Current HVAC mechanical code. |
| adoption:usa-id:fuel-gas:idaho-fuel-gas-provisions-2018 | state_code.name | Idaho Mechanical Code / Fuel Gas provisions | verified | 0.96 | src:usa-id:hvac-rules-2023 | Fuel-gas scope is carved out in HVAC rules. |
| local-enforcement:usa-id | model | state_administered | verified | 0.94 | src:usa-id:bld-rules-2024; src:usa-id:plan-review | DOPL runs the permit and inspection workflow. |
| local-amendment-rule:usa-id | model | state_uniform | partially_verified | 0.68 | src:usa-id:bld-statutes-rules; src:usa-id:plan-review | No local amendment registry was identified. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Source IDs map to Idaho DOPL pages, rule PDFs, and program pages. |
| All authority IDs resolve | pass | Board/program authority IDs are internally consistent. |
| All current code families have adoption records | pass | Building, residential, existing, mechanical, plumbing, fuel gas, electrical, energy, accessibility, and elevator are covered. |
| Building and operational fire code are separated | fail | Fire-code authority was not verified in this pass. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Only verified effective dates were entered. |
| Effective dates are valid ISO dates | pass | Entered dates are valid ISO dates. |
| No impossible date sequences | pass | No contradictory date chain was asserted. |
| Transition rules have explicit trigger conditions | pass | Permit-date triggers are stated where applicable. |
| Permit-date logic is captured where applicable | pass | Permit/plan-review triggers were recorded for the code families covered here. |
| Local enforcement model classified | pass | Classified as state_administered. |
| Local amendment rule classified | pass | Classified as state_uniform with partial verification. |
| AHJ confirmation metadata present | pass | AHJ contact data is present for the active programs. |
| Official-source caveats captured | pass | Caveats call out fire-code and standards-bundle gaps. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-id:001 | high | statewide fire code | Confirm whether Idaho has a statewide fire-code adoption path or a separate fire-marshals program. | Extract the state fire authority and current fire code. | null | null | open |
| issue:usa-id:002 | medium | local amendment authority | Confirm whether any local amendment authority exists outside the DOPL rule chapters. | Read the underlying statute directly. | null | null | open |
| issue:usa-id:003 | medium | elevator scope | Confirm whether elevator conveyance should remain a separate program or be cross-referenced into the building code family. | Cross-check the elevator statutes and DOPL rules. | null | null | open |
| issue:usa-id:004 | low | LPG crosswalk | Confirm whether LPG safety should be modeled separately from HVAC fuel-gas scope in downstream reports. | Reconcile the LPG board program with the HVAC fuel-gas chapter. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-id:bld-board | src:usa-id:bld-home | html_diff | monthly | building-board page or rule chapter changes | 2026-06-24 | active |
| watch:usa-id:ele-board | src:usa-id:ele-home | html_diff | monthly | electrical-board page or rule chapter changes | 2026-06-24 | active |
| watch:usa-id:plb-board | src:usa-id:plb-home | html_diff | monthly | plumbing-board page or rule chapter changes | 2026-06-24 | active |
| watch:usa-id:hvac-board | src:usa-id:hvac-home | html_diff | monthly | HVAC board page or rule chapter changes | 2026-06-24 | active |
| watch:usa-id:elev-program | src:usa-id:elev-home | html_diff | quarterly | elevator program code bundle changes | 2026-06-24 | active |
| watch:usa-id:lpg-program | src:usa-id:lpg-home | html_diff | quarterly | LPG board document or page changes | 2026-06-24 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-24 | Replaced Idaho stub with a source-backed report using DOPL boards, Idaho Administrative Code rule chapters, and program pages. | report:usa-id | src:usa-id:bld-home; src:usa-id:bld-statutes-rules; src:usa-id:bld-rules-2024; src:usa-id:plan-review; src:usa-id:ele-home; src:usa-id:ele-rules-2025; src:usa-id:plb-home; src:usa-id:plb-rules-2023; src:usa-id:hvac-home; src:usa-id:hvac-rules-2023; src:usa-id:elev-home; src:usa-id:elev-codes; src:usa-id:lpg-home; src:usa-id:lpg-documents | Codex | Idaho is a state-administered, board-driven code program with a remaining fire-code gap in this pass. |