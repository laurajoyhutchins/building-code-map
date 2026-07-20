---
state:
  state_id: "US-CO"
  name: "Colorado"
  abbreviation: "CO"
report:
  report_id: "state-report:usa-co"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-25"
  last_verified: "2026-06-25"
  reviewed_by: null
risk:
  overall_confidence: 0.68 # 0.00 - 1.00
  risk_flags:
    - "no_single_statewide_ibc_irc_for_all_buildings_verified"
    - "local_building_code_adoption_model_with_state_energy_floor"
    - "fire_code_authority_verified_for_dfpc_administered_scopes_only"
    - "local_amendment_scope_partially_unresolved"
    - "ahj_boundary_and_contact_data_not_populated"
  open_questions_count: 6
---

# State Building Code Authority Report: Colorado

## 1. Executive Summary

- **Authority model:** Colorado is a hybrid, local-primary building-code state for ordinary county and municipal construction. County boards of commissioners are authorized to adopt county building codes in unincorporated areas, and municipalities are treated as local governments with building-code adoption and energy-code duties. Colorado also has statewide trade and special-scope programs, including the State Electrical Board, State Plumbing Board, Division of Fire Prevention and Control, Office of the State Architect, and Division of Housing.

- **Statewide code status:** A single statewide IBC/IRC adoption for all buildings was not verified. The strongest statewide construction-code floor identified here is the energy-code framework: local governments that have adopted and enforced building codes, or that adopt and enforce building codes after 2022-07-01, must adopt and enforce an energy code meeting statutory minimum performance requirements. Separate statewide or state-administered code adoptions exist for electrical work, plumbing and fuel gas, factory-built and certain housing scopes, and Division of Fire Prevention and Control programs.

- **Local enforcement model:** Ordinary building-code enforcement appears primarily local, subject to state trade-code and special-scope overlays. For counties, the board of county commissioners may appoint a county building inspector to administer and enforce the adopted county building code. For public schools, charter schools, institute charter schools, and junior colleges, the Division of Fire Prevention and Control conducts or oversees plan review, permitting, and inspections, with defined delegation to qualified local authorities.

- **Local amendment posture:** County building-code amendments are supported by statute through resolution after public hearing. For plumbing and fuel gas, a city, town, county, or city and county adopting more stringent standards than the Colorado Plumbing Code or Colorado Fuel Gas Code must furnish a copy to the State Plumbing Board. Municipal building-code amendment authority and the full local energy-code amendment/equivalency process remain partially unresolved.

- **Known transition periods or pending changes:** For local governments updating building codes from 2023-07-01 through 2026-06-30, Colorado requires an energy code at least as stringent as the 2021 IECC plus the Colorado Model Electric Ready and Solar Ready Code language. For updates on or after 2026-07-01, the trigger moves to the Model Low Energy and Carbon Code performance threshold, with statutory timing exceptions when the only update is to the NEC, elevator and escalator code, or plumbing code.

- **Production readiness:** narrow_partial_validation_ready

### Key Findings

```yaml
---
key_findings:
- topic: Overall authority model
  finding: Colorado does not appear to have a single statewide IBC/IRC adoption for
    all ordinary buildings; local county and municipal adoption remains central, with
    state energy, trade, and special-program overlays.
  confidence: 0.72
  source_ids:
  - src:usa-co:crs-title-30
  - src:usa-co:crs-title-31
  - src:usa-co:session-law-2022-ch301
- topic: County building-code authority
  finding: County boards of commissioners are authorized to adopt building codes in
    all or part of the county outside incorporated cities and towns.
  confidence: 0.9
  source_ids:
  - src:usa-co:crs-title-30
- topic: Local energy-code floor
  finding: Counties and municipalities with building codes must adopt and enforce
    energy codes meeting Colorado's statutory thresholds when adopting or updating
    building codes.
  confidence: 0.88
  source_ids:
  - src:usa-co:crs-title-30
  - src:usa-co:crs-title-31
- topic: Electrical code
  finding: The State Electrical Board incorporates NFPA 70, National Electrical Code,
    2023 edition, as minimum standards, effective 2023-08-01.
  confidence: 0.92
  source_ids:
  - src:usa-co:3-ccr-710-1
  - src:usa-co:3-ccr-710-1-rule-info
- topic: Plumbing and fuel gas
  finding: The State Plumbing Board rules establish the Colorado Plumbing Code and
    Colorado Fuel Gas Code using incorporated 2018 IPC, IRC, and IFGC provisions,
    with additions, revisions, and deletions.
  confidence: 0.86
  source_ids:
  - src:usa-co:3-ccr-720-1
  - src:usa-co:3-ccr-720-1-rule-info
- topic: DFPC state-administered building and fire code
  finding: DFPC adopts building, fire, life-safety, and related standards for property,
    buildings, and structures subject to DFPC oversight; this is not evidence of universal
    statewide building-code coverage.
  confidence: 0.82
  source_ids:
  - src:usa-co:8-ccr-1507-101
  - src:usa-co:8-ccr-1507-30
- topic: Factory-built and selected housing scopes
  finding: Division of Housing / State Housing Board rules cover factory-built structures,
    tiny homes, installations, and certain hotels, motels, and multi-family structures
    in areas where no standards exist.
  confidence: 0.82
  source_ids:
  - src:usa-co:8-ccr-1302-14
  - src:usa-co:8-ccr-1302-14-rule-info
- topic: Local amendments
  finding: County building-code amendments are verified; municipal and full local
    energy-amendment processes remain partially unresolved.
  confidence: 0.55
  source_ids:
  - src:usa-co:crs-title-30
  - src:usa-co:3-ccr-720-1
```

---

### 2.1 Primary Building Code Authorities

| Authority ID | Authority Name | Authority Type | Legal Basis | Role | Enforcement Model | Source IDs | Verification Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| ahj:usa-co:local-building-code-authorities | County boards of commissioners and municipal governments | local_adopting_authorities | CRS Title 30 county building-code provisions; CRS Title 31 municipal energy-code provisions | Adopt and administer local building codes for ordinary local construction where local code authority exists. | local_primary_with_state_overlays | src:usa-co:crs-title-30, src:usa-co:crs-title-31 | partially_verified |
| ahj:usa-co:energy-code-board | Colorado Energy Code Board / Colorado Energy Office / Department of Local Affairs | state_model_code_and_energy_policy_authority | HB22-1362 / Session Law Ch. 301; CRS energy-code provisions | Develops model energy-code language used in statutory local-government and state-agency energy-code requirements. | state_floor_local_enforcement | src:usa-co:session-law-2022-ch301, src:usa-co:crs-title-30, src:usa-co:crs-title-31 | partially_verified |
| ahj:usa-co:state-electrical-board | Colorado State Electrical Board | state_trade_code_authority | 3 CCR 710-1; Electrical Practice Act references | Adopts and enforces electrical rules and minimum electrical standards. | state_trade_program_with_possible_local_interfaces | src:usa-co:dora-electrical, src:usa-co:3-ccr-710-1, src:usa-co:3-ccr-710-1-rule-info | verified |
| ahj:usa-co:state-plumbing-board | Colorado State Plumbing Board | state_trade_code_authority | 3 CCR 720-1; Title 12 Article 155 references | Adopts Colorado Plumbing Code and Colorado Fuel Gas Code; governs plumbing/fuel gas standards statewide for covered work. | state_trade_program_with_local_more_stringent_reporting | src:usa-co:3-ccr-720-1, src:usa-co:3-ccr-720-1-rule-info | verified |
| ahj:usa-co:dfpc | Colorado Division of Fire Prevention and Control | state_fire_and_life_safety_program_authority | 8 CCR 1507-101; 8 CCR 1507-30 | Adopts building and fire codes for DFPC-administered programs and serves as or delegates fire-code official functions for specified public-school and junior-college scopes. | state_administered_and_delegated_for_specified_scopes | src:usa-co:8-ccr-1507-101, src:usa-co:8-ccr-1507-30 | partially_verified |
| ahj:usa-co:division-of-housing | Colorado Division of Housing / State Housing Board | state_special_building_program_authority | 8 CCR 1302-14; HB22-1362 / Session Law Ch. 301 | Regulates factory-built structures, tiny homes, installations, and certain hotels, motels, and multi-family structures in areas where no construction standards exist. | state_special_scope_program | src:usa-co:8-ccr-1302-14, src:usa-co:8-ccr-1302-14-rule-info, src:usa-co:session-law-2022-ch301 | partially_verified |
| ahj:usa-co:office-of-state-architect | Colorado Office of the State Architect | state_facilities_authority | HB22-1362 / Session Law Ch. 301 | Must adopt and enforce energy-code requirements for construction by state agencies on covered state-owned or state-leased properties and facilities. | state_facility_program | src:usa-co:session-law-2022-ch301 | partially_verified |
| ahj:usa-co:elevator-safety | Elevator and escalator code authority | state_trade_or_special_program_authority | referenced in CRS Title 30 and Title 31 energy-code timing exceptions | Authority and current adopted edition not independently extracted; listed because the statutes distinguish elevator and escalator code updates from ordinary building-code updates. | unresolved | src:usa-co:crs-title-30, src:usa-co:crs-title-31 | unresolved |
| ahj:usa-co:accessibility | Accessibility authority | state_or_local_special_program_authority | unresolved | Accessibility-specific construction-code authority was not extracted. | unresolved | src:usa-co:ccr-portal | unresolved |

### 2.2 Specialized Code Authorities by Code Family

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-co:local-building-code-authorities | County boards and municipalities | Local adoption and enforcement for ordinary building-code scope; DFPC/Division of Housing overlays for specific scopes. | CRS Title 30; CRS Title 31; 8 CCR 1507-101; 8 CCR 1302-14 | src:usa-co:crs-title-30, src:usa-co:crs-title-31, src:usa-co:8-ccr-1507-101, src:usa-co:8-ccr-1302-14 | partially_verified |
| Residential | ahj:usa-co:local-building-code-authorities | County boards and municipalities | Local adoption; state special-scope overlays for factory-built/tiny-home structures and DFPC scopes. | CRS Title 30; CRS Title 31; 8 CCR 1302-14 | src:usa-co:crs-title-30, src:usa-co:crs-title-31, src:usa-co:8-ccr-1302-14 | partially_verified |
| Existing Building / Rehabilitation | ahj:usa-co:dfpc | DFPC for state-administered scopes; local governments otherwise | DFPC adopts 2021 IEBC for its oversight scopes; local-government existing-building treatment otherwise varies. | 8 CCR 1507-101 | src:usa-co:8-ccr-1507-101 | partially_verified |
| Mechanical | ahj:usa-co:dfpc | DFPC for state-administered scopes; local governments otherwise | DFPC adopts 2021 IMC for its oversight scopes; statewide ordinary mechanical code not verified. | 8 CCR 1507-101 | src:usa-co:8-ccr-1507-101 | partially_verified |
| Plumbing | ahj:usa-co:state-plumbing-board | Colorado State Plumbing Board | State plumbing code authority. | 3 CCR 720-1 | src:usa-co:3-ccr-720-1, src:usa-co:3-ccr-720-1-rule-info | verified |
| Fuel Gas | ahj:usa-co:state-plumbing-board | Colorado State Plumbing Board | State fuel-gas code authority. | 3 CCR 720-1 | src:usa-co:3-ccr-720-1, src:usa-co:3-ccr-720-1-rule-info | verified |
| Electrical | ahj:usa-co:state-electrical-board | Colorado State Electrical Board | State electrical code authority. | 3 CCR 710-1 | src:usa-co:3-ccr-710-1, src:usa-co:3-ccr-710-1-rule-info | verified |
| Energy | ahj:usa-co:energy-code-board | Energy Code Board / CEO / DOLA plus local governments | State energy-code floor applied through local government and state-agency adoption/enforcement triggers. | HB22-1362; CRS Title 30; CRS Title 31 | src:usa-co:session-law-2022-ch301, src:usa-co:crs-title-30, src:usa-co:crs-title-31 | verified_core |
| Fire - construction references | ahj:usa-co:dfpc | DFPC | DFPC fire and life-safety code authority for DFPC-administered scopes; local fire AHJ model outside those scopes unresolved. | 8 CCR 1507-101; 8 CCR 1507-30 | src:usa-co:8-ccr-1507-101, src:usa-co:8-ccr-1507-30 | partially_verified |
| Fire - operational / prevention code | ahj:usa-co:dfpc | DFPC and local fire authorities | DFPC rule coverage verified for DFPC-administered scopes; statewide/local operational fire-code adoption outside those scopes unresolved. | 8 CCR 1507-101; 8 CCR 1507-30 | src:usa-co:8-ccr-1507-101, src:usa-co:8-ccr-1507-30 | partially_verified |
| Accessibility | ahj:usa-co:accessibility | Unknown | Not extracted. | Not extracted | src:usa-co:ccr-portal | unresolved |
| Elevator / Conveyance | ahj:usa-co:elevator-safety | Unknown | Referenced as a distinct update category in energy-code timing exceptions; current code edition not extracted. | CRS Title 30; CRS Title 31 | src:usa-co:crs-title-30, src:usa-co:crs-title-31 | unresolved |

### 2.3 Authority Hierarchy Notes

Colorado should not be modeled as a simple single-adopting-authority state. For ordinary buildings, the query path should first determine whether the parcel is inside a municipality or in unincorporated county territory, then identify the applicable local building-code adoption. The state energy-code floor is not a substitute for the local building code; it is a statutory requirement layered onto local governments that have adopted and enforced building codes. Electrical, plumbing/fuel gas, factory-built/tiny-home, state-facility, and DFPC-administered buildings introduce separate state-level overlays.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-co:001 | ahj:usa-co:local-building-code-authorities | adopts_and_enforces | ordinary local building codes, varying by jurisdiction | src:usa-co:crs-title-30, src:usa-co:crs-title-31 | partially_verified |
| edge:usa-co:002 | ahj:usa-co:energy-code-board | supplies_model_code_floor_for | local government energy-code adoption and enforcement | src:usa-co:session-law-2022-ch301, src:usa-co:crs-title-30, src:usa-co:crs-title-31 | verified_core |
| edge:usa-co:003 | ahj:usa-co:state-electrical-board | adopts_minimum_standards | NFPA 70 / NEC 2023 | src:usa-co:3-ccr-710-1 | verified |
| edge:usa-co:004 | ahj:usa-co:state-plumbing-board | adopts_minimum_standards | Colorado Plumbing Code and Colorado Fuel Gas Code | src:usa-co:3-ccr-720-1 | verified |
| edge:usa-co:005 | ahj:usa-co:dfpc | adopts_and_enforces_or_delegates | DFPC-administered building and fire code scopes | src:usa-co:8-ccr-1507-101, src:usa-co:8-ccr-1507-30 | partially_verified |
| edge:usa-co:006 | ahj:usa-co:division-of-housing | adopts_and_administers | factory-built structures, tiny homes, installations, and specified no-local-standard housing scopes | src:usa-co:8-ccr-1302-14, src:usa-co:session-law-2022-ch301 | partially_verified |
| edge:usa-co:007 | ahj:usa-co:office-of-state-architect | adopts_and_enforces | energy code for covered state-agency construction | src:usa-co:session-law-2022-ch301 | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | No single statewide ordinary building code verified | Local model code selection varies; DFPC and Division of Housing have separate state-administered scopes | varies locally; 2021 IBC for DFPC and Division of Housing verified scopes | local_or_special_scope | null | null | null | null | Local adoption varies; state energy-code floor applies when local governments adopt or update building codes. | src:usa-co:crs-title-30, src:usa-co:crs-title-31, src:usa-co:8-ccr-1507-101, src:usa-co:8-ccr-1302-14 |
| Residential | No single statewide ordinary residential code verified | Local model code selection varies; Division of Housing/factory-built scopes use state standards | varies locally; 2021 IRC appears in DFPC and Division of Housing verified scopes | local_or_special_scope | null | null | null | null | Local adoption varies; state energy-code floor applies when local governments adopt or update building codes. | src:usa-co:crs-title-30, src:usa-co:crs-title-31, src:usa-co:8-ccr-1507-101, src:usa-co:8-ccr-1302-14 |
| Existing Building / Rehabilitation | DFPC-administered existing-building scope | International Existing Building Code | 2021 | special_scope_verified | null | null | 2021-07-01 | null | DFPC rule says new work after 2021-07-01 must meet new-construction requirements as amended by IEBC/NFPA 101 provisions; ordinary local scope unresolved. | src:usa-co:8-ccr-1507-101 |
| Mechanical | DFPC-administered mechanical scope | International Mechanical Code | 2021 | special_scope_verified | null | null | null | null | Ordinary local mechanical-code adoption not extracted. | src:usa-co:8-ccr-1507-101 |
| Plumbing | Colorado Plumbing Code | Selected IPC and IRC provisions with Colorado additions, revisions, and deletions | 2018 IPC and 2018 IRC provisions | statewide_trade_code_verified | null | 2026-02-14 | 2026-02-14 | 2026-02-14 | Current rule page effective 2026-02-14; incorporated 2018 provisions do not include later amendments or editions. | src:usa-co:3-ccr-720-1, src:usa-co:3-ccr-720-1-rule-info |
| Fuel Gas | Colorado Fuel Gas Code | Selected IFGC and IRC provisions with Colorado additions, revisions, and deletions | 2018 IFGC and 2018 IRC provisions | statewide_trade_code_verified | null | 2026-02-14 | 2026-02-14 | 2026-02-14 | Current rule page effective 2026-02-14; incorporated 2018 provisions do not include later amendments or editions. | src:usa-co:3-ccr-720-1, src:usa-co:3-ccr-720-1-rule-info |
| Electrical | State Electrical Board minimum electrical standards | NFPA 70 / National Electrical Code | 2023 | statewide_trade_code_verified | null | 2023-08-01 | 2023-08-01 | 2023-08-01 | Current rule page effective 2026-03-30; NEC incorporation effective 2023-08-01. | src:usa-co:3-ccr-710-1, src:usa-co:3-ccr-710-1-rule-info |
| Energy - local governments | Colorado statutory local energy-code floor | 2021 IECC plus Colorado Model Electric Ready and Solar Ready Code until 2026-07-01 update trigger; Model Low Energy and Carbon Code after 2026-07-01 update trigger | performance threshold, not a single pasted model code | statewide_floor_local_enforcement | 2022-06-02 | 2023-07-01 | 2023-07-01 | varies_by_local_update_trigger | See section 4; timing exceptions for NEC, elevator/escalator, and plumbing-only updates. | src:usa-co:session-law-2022-ch301, src:usa-co:crs-title-30, src:usa-co:crs-title-31 |
| Energy - state agencies | State-agency energy-code duty | 2021 IECC plus Model Electric Ready/Solar Ready by 2025-01-01; Model Low Energy and Carbon Code by 2030-01-01 | performance threshold | state_facility_energy_code | 2022-06-02 | 2025-01-01 | 2025-01-01 | 2025-01-01 and 2030-01-01 | Applies to covered state-agency construction on state-owned or specified state-leased properties/facilities. | src:usa-co:session-law-2022-ch301 |
| Fire - construction references | DFPC building/fire/life-safety rules | IBC, IMC, IECC, IEBC, IRC, IFC, NFPA standards, and Colorado model/wildfire codes for DFPC scopes | primarily 2021 ICC editions plus listed NFPA editions and Colorado codes | special_scope_verified | null | null | null | null | Applies only to buildings and life-safety systems subject to DFPC oversight. | src:usa-co:8-ccr-1507-101 |
| Fire - operational / prevention code | DFPC fire code for covered scopes | International Fire Code and listed NFPA standards | 2021 IFC plus listed NFPA editions | special_scope_verified | null | null | null | null | DFPC is fire-code official for specified public-school and junior-college scopes unless delegated to qualified local AHJ. | src:usa-co:8-ccr-1507-101, src:usa-co:8-ccr-1507-30 |
| Factory-built / Tiny homes / selected no-local-standard housing | Division of Housing / State Housing Board rules | 2021 IBC and additional listed standards; full list not exhaustively parsed | 2021 IBC verified; other standards require full extraction | special_scope_partially_verified | null | 2026-03-17 | 2026-03-17 | 2026-03-17 | Current rule page effective 2026-03-17; manufacturers may use prior construction codes for up to 180 days after a new rule amendment takes effect, with possible 180-day extension. | src:usa-co:8-ccr-1302-14, src:usa-co:8-ccr-1302-14-rule-info |
| Accessibility | unresolved | unresolved | unresolved | unresolved | null | null | null | null | Accessibility-specific construction-code authority and editions not extracted. | src:usa-co:ccr-portal |
| Elevator / Conveyance | unresolved | unresolved | unresolved | unresolved | null | null | null | null | Elevator and escalator code updates are distinguished in energy-code timing rules, but the current elevator code was not extracted. | src:usa-co:crs-title-30, src:usa-co:crs-title-31 |

### 3.2 Normalized Adoption Records

```yaml
adoption_records:
  - adoption_id: "adoption:usa-co:county-building-local"
    code_family: "Building"
    authority_id: "ahj:usa-co:local-building-code-authorities"
    code_name: "Locally adopted county building code"
    base_model_code: "varies_by_county"
    edition: "varies_by_county"
    adoption_date: null
    effective_date: null
    operative_date: null
    mandatory_date: null
    applicability: "all or part of a county outside incorporated cities and towns, when adopted by county board of commissioners"
    source_ids:
      - "src:usa-co:crs-title-30"
    verification_status: "partially_verified"

  - adoption_id: "adoption:usa-co:local-energy-2023-to-2026"
    code_family: "Energy"
    authority_id: "ahj:usa-co:energy-code-board"
    code_name: "Local energy-code floor for local governments updating building codes from 2023-07-01 through 2026-06-30"
    base_model_code: "2021 IECC plus Colorado Model Electric Ready and Solar Ready Code language"
    edition: "2021_IECC_plus_colorado_model_electric_ready_solar_ready"
    adoption_date: "2022-06-02"
    effective_date: "2023-07-01"
    operative_date: "2023-07-01"
    mandatory_date: "varies_by_local_building_code_update"
    applicability: "local governments with adopted/enforced building codes when adopting or updating building codes during the statutory window"
    source_ids:
      - "src:usa-co:crs-title-30"
      - "src:usa-co:crs-title-31"
      - "src:usa-co:session-law-2022-ch301"
    verification_status: "verified_core"

  - adoption_id: "adoption:usa-co:local-energy-2026-forward"
    code_family: "Energy"
    authority_id: "ahj:usa-co:energy-code-board"
    code_name: "Local energy-code floor for local governments updating building codes on or after 2026-07-01"
    base_model_code: "Model Low Energy and Carbon Code performance threshold"
    edition: "model_low_energy_and_carbon_code_threshold"
    adoption_date: "2022-06-02"
    effective_date: "2026-07-01"
    operative_date: "2026-07-01"
    mandatory_date: "varies_by_local_building_code_update_or_statutory_backstop"
    applicability: "local governments with adopted/enforced building codes when adopting or updating building codes on or after 2026-07-01, subject to timing exceptions"
    source_ids:
      - "src:usa-co:crs-title-30"
      - "src:usa-co:crs-title-31"
      - "src:usa-co:session-law-2022-ch301"
    verification_status: "verified_core"

  - adoption_id: "adoption:usa-co:electrical-nec-2023"
    code_family: "Electrical"
    authority_id: "ahj:usa-co:state-electrical-board"
    code_name: "Colorado minimum electrical standards"
    base_model_code: "NFPA 70 / National Electrical Code"
    edition: "2023"
    adoption_date: null
    effective_date: "2023-08-01"
    operative_date: "2023-08-01"
    mandatory_date: "2023-08-01"
    applicability: "electrical practice, electrical services, and electrical contracting in Colorado"
    source_ids:
      - "src:usa-co:3-ccr-710-1"
      - "src:usa-co:3-ccr-710-1-rule-info"
    verification_status: "verified"

  - adoption_id: "adoption:usa-co:plumbing-2018"
    code_family: "Plumbing"
    authority_id: "ahj:usa-co:state-plumbing-board"
    code_name: "Colorado Plumbing Code"
    base_model_code: "selected IPC and IRC provisions with Colorado revisions"
    edition: "2018 IPC and 2018 IRC provisions"
    adoption_date: null
    effective_date: "2026-02-14"
    operative_date: "2026-02-14"
    mandatory_date: "2026-02-14"
    applicability: "inspection, installation, alteration, and repair of plumbing fixtures, appliances, and systems throughout Colorado"
    source_ids:
      - "src:usa-co:3-ccr-720-1"
      - "src:usa-co:3-ccr-720-1-rule-info"
    verification_status: "verified"

  - adoption_id: "adoption:usa-co:fuel-gas-2018"
    code_family: "Fuel Gas"
    authority_id: "ahj:usa-co:state-plumbing-board"
    code_name: "Colorado Fuel Gas Code"
    base_model_code: "selected IFGC and IRC provisions with Colorado revisions"
    edition: "2018 IFGC and 2018 IRC provisions"
    adoption_date: null
    effective_date: "2026-02-14"
    operative_date: "2026-02-14"
    mandatory_date: "2026-02-14"
    applicability: "inspection, installation, alteration, and repair of fuel gas piping and systems throughout Colorado"
    source_ids:
      - "src:usa-co:3-ccr-720-1"
      - "src:usa-co:3-ccr-720-1-rule-info"
    verification_status: "verified"

  - adoption_id: "adoption:usa-co:dfpc-1507-101"
    code_family: "Building / Fire / Life Safety"
    authority_id: "ahj:usa-co:dfpc"
    code_name: "DFPC building and fire code adoption for state-administered programs"
    base_model_code: "2021 IBC, 2021 IMC, 2021 IECC, 2021 IEBC, 2021 IRC, 2021 IFC, listed NFPA standards, and Colorado model/wildfire codes"
    edition: "mixed_listed_editions"
    adoption_date: null
    effective_date: null
    operative_date: null
    mandatory_date: null
    applicability: "property, buildings, structures, and life-safety systems subject to DFPC oversight"
    source_ids:
      - "src:usa-co:8-ccr-1507-101"
    verification_status: "partially_verified"

  - adoption_id: "adoption:usa-co:division-of-housing-1302-14"
    code_family: "Factory-built / Tiny homes / selected no-local-standard housing"
    authority_id: "ahj:usa-co:division-of-housing"
    code_name: "Division of Housing / State Housing Board construction standards"
    base_model_code: "2021 IBC and additional listed standards"
    edition: "2021 IBC verified; full list requires extraction"
    adoption_date: null
    effective_date: "2026-03-17"
    operative_date: "2026-03-17"
    mandatory_date: "2026-03-17"
    applicability: "factory-built structures, tiny homes, installations, and certain hotels, motels, and multi-family structures in areas without standards"
    source_ids:
      - "src:usa-co:8-ccr-1302-14"
      - "src:usa-co:8-ccr-1302-14-rule-info"
    verification_status: "partially_verified"
```

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

The clearest statewide transition rules are energy-code transition rules. Colorado separates local-government building-code updates from updates limited to the National Electrical Code, elevator and escalator code, or plumbing code. The rules should be modeled as update-triggered, not as a universal date on which every local jurisdiction automatically has the same complete building-code stack.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-co:energy-local-2023-window | Local government energy codes | update_trigger | 2023-07-01 through 2026-06-30 | Local government with adopted/enforced building code adopts or updates any building code during this period. | Existing local code remains until local update trigger, subject to statutory requirements. | src:usa-co:crs-title-30, src:usa-co:crs-title-31 | verified_core |
| date-rule:usa-co:energy-local-2026-forward | Local government energy codes | update_trigger | on or after 2026-07-01 | Local government with adopted/enforced building code adopts or updates any building code on or after this date. | Limited timing exceptions apply where only NEC, elevator/escalator, or plumbing code is updated. | src:usa-co:crs-title-30, src:usa-co:crs-title-31 | verified_core |
| date-rule:usa-co:energy-nec-elevator-plumbing-exception | Local energy-code timing exceptions | exception | through 2026-06-30 and by 2030-01-01 depending on trigger | If the only building-code update is to the NEC, elevator/escalator rules, or plumbing code, the local government need not update energy code immediately; statutory backstops still apply. | yes, until applicable backstop | src:usa-co:crs-title-30, src:usa-co:crs-title-31 | partially_verified |
| date-rule:usa-co:electrical-nec-2023 | Electrical code | effective_date | 2023-08-01 | State Electrical Board incorporation by reference of NEC 2023. | no later NEC amendments or editions included unless adopted by the Board. | src:usa-co:3-ccr-710-1 | verified |
| date-rule:usa-co:plumbing-current-rule | Plumbing and fuel gas | effective_date | 2026-02-14 | Current 3 CCR 720-1 rule page effective date. | incorporated 2018 provisions remain as listed; no later editions included by reference. | src:usa-co:3-ccr-720-1, src:usa-co:3-ccr-720-1-rule-info | verified |
| date-rule:usa-co:factory-built-current-rule | Division of Housing 8 CCR 1302-14 | effective_date | 2026-03-17 | Current 8 CCR 1302-14 rule page effective date. | rule allows prior construction codes for up to 180 days after amendment effective date, with possible 180-day extension. | src:usa-co:8-ccr-1302-14, src:usa-co:8-ccr-1302-14-rule-info | verified_core |
| date-rule:usa-co:permit-date-ordinary-building | Ordinary local building code permit-date logic | unresolved | null | Permit-date / application-date treatment for ordinary local building codes was not fully extracted. | unresolved | src:usa-co:crs-title-30, src:usa-co:crs-title-31 | unresolved |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Energy - local governments | Model Low Energy and Carbon Code performance threshold | 2022-06-02 | 2022-06-02 | 2026-07-01 | 2026-07-01 | varies_by_local_update_or_backstop | active | src:usa-co:session-law-2022-ch301, src:usa-co:crs-title-30, src:usa-co:crs-title-31 | For local governments with building codes, triggered by adoption/update of building codes on or after 2026-07-01, subject to timing exceptions. |
| Energy - state agencies | Model Low Energy and Carbon Code performance threshold | 2022-06-02 | 2022-06-02 | 2030-01-01 | 2030-01-01 | 2030-01-01 | active | src:usa-co:session-law-2022-ch301 | Applies to Office of State Architect, Division of Housing, and DFPC scopes described in HB22-1362 / Session Law Ch. 301. |
| Electrical | Future NEC edition | null | null | null | null | null | watch | src:usa-co:3-ccr-710-1-rule-info | Watch State Electrical Board rulemaking and current 3 CCR 710-1 rule page. |
| Plumbing / Fuel Gas | Future plumbing/fuel-gas edition | null | null | null | null | null | watch | src:usa-co:3-ccr-720-1-rule-info | Watch State Plumbing Board rulemaking and current 3 CCR 720-1 rule page. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-co:county-building-code | Building | Unincorporated county areas | County board adoption | County building code may apply in all or part of the county outside incorporated cities and towns. | src:usa-co:crs-title-30 | verified |
| applicability-rule:usa-co:energy-rural-exception | Energy | Counties not required to enforce energy code before 2022-07-01 and lacking all four infrastructure services | statutory criteria | Counties satisfying the statutory rural infrastructure exception are treated differently in the local-government energy-code rule. | src:usa-co:crs-title-30 | partially_verified |
| applicability-rule:usa-co:dfpc-school-fire-official | Fire / building | Public schools, institute charter schools, charter schools, and junior colleges | DFPC rule scope | DFPC is the fire-code official unless delegated to a qualified local AHJ; if the local AHJ declines or lacks a certified inspector, DFPC performs plan reviews and inspections. | src:usa-co:8-ccr-1507-30 | verified |
| applicability-rule:usa-co:state-electrical-ev-ready | Electrical / energy-related electrical infrastructure | EV power transfer infrastructure permit applicant | on or after 2024-03-01 | Permit applicant must comply with CRS and Colorado Energy Office requirements in the Model Electric Ready and Solar Ready Code for EV power transfer infrastructure. | src:usa-co:3-ccr-710-1 | verified |
| applicability-rule:usa-co:division-of-housing-no-local-standards | Building / energy / housing | Hotels, motels, and multi-family structures in areas where no standards exist | Division of Housing scope | Division of Housing / State Housing Board rules and energy-code duties apply to specified scopes where no local construction standards exist. | src:usa-co:8-ccr-1302-14, src:usa-co:session-law-2022-ch301 | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

State amendment structure:
- Electrical: State Electrical Board incorporates NEC 2023 and may amend the incorporated standard; 3 CCR 710-1 does not include later NEC amendments or editions unless adopted.
- Plumbing/Fuel Gas: State Plumbing Board adopts selected 2018 IPC, IRC, and IFGC provisions with Colorado additions, revisions, and deletions; later amendments or editions are not included unless adopted.
- DFPC: DFPC adopts listed model codes and standards with provisions applicable to DFPC-administered scopes.
- Local building: county building-code amendments are allowed by county resolution after public hearing.
- Local energy: equivalency and amendment mechanics remain partially unresolved for ordinary local governments; state-agency amendment authority is verified where HB22-1362 says amendments may not decrease effectiveness or energy efficiency.

Where amendments are published:
- CCR rules for state trade and special-scope codes.
- Local ordinances/resolutions for local building-code amendments.
- State model energy-code materials and local adoption instruments for local-government energy-code implementation.

**Amendment parsing status:** partial

### 5.2 State Amendment Sources

| Amendment Source ID | Code Family | Publisher | Publication Path | Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| amendment-source:usa-co:3-ccr-710-1 | Electrical | State Electrical Board / Secretary of State | 3 CCR 710-1 current PDF | State electrical rules and NEC 2023 incorporation/amendments | src:usa-co:3-ccr-710-1, src:usa-co:3-ccr-710-1-rule-info | verified |
| amendment-source:usa-co:3-ccr-720-1 | Plumbing / Fuel Gas | State Plumbing Board / Secretary of State | 3 CCR 720-1 current PDF | Colorado Plumbing Code and Colorado Fuel Gas Code additions, revisions, deletions, and alternate method petition process | src:usa-co:3-ccr-720-1, src:usa-co:3-ccr-720-1-rule-info | verified |
| amendment-source:usa-co:8-ccr-1507-101 | Building / Fire / Life Safety | DFPC / Secretary of State | 8 CCR 1507-101 current PDF | DFPC-administered building, fire, life-safety, and related standards | src:usa-co:8-ccr-1507-101 | partially_verified |
| amendment-source:usa-co:8-ccr-1302-14 | Factory-built / Tiny homes / no-local-standard housing | Division of Housing / Secretary of State | 8 CCR 1302-14 current PDF | State Housing Board and Division of Housing codes and standards | src:usa-co:8-ccr-1302-14, src:usa-co:8-ccr-1302-14-rule-info | partially_verified |
| amendment-source:usa-co:county-building | Local building | County boards of commissioners | County resolutions after public hearing | County building-code changes, alterations, and amendments | src:usa-co:crs-title-30 | verified_core |
| amendment-source:usa-co:municipal-building | Local building | Municipal governments | Local ordinances | Municipal amendment procedure not extracted. | src:usa-co:crs-title-31 | unresolved |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Topic | Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| amendment:usa-co:energy-2023-floor | Energy | 2021 IECC plus electric-ready/solar-ready | Local building-code updates from 2023-07-01 through 2026-06-30 require an energy code meeting or exceeding 2021 IECC plus Colorado Model Electric Ready and Solar Ready Code language. | src:usa-co:crs-title-30, src:usa-co:crs-title-31 | verified_core |
| amendment:usa-co:energy-2026-floor | Energy | Model Low Energy and Carbon Code | Local building-code updates on or after 2026-07-01 require an energy code meeting or exceeding the Model Low Energy and Carbon Code performance threshold, subject to statutory timing exceptions. | src:usa-co:crs-title-30, src:usa-co:crs-title-31 | verified_core |
| amendment:usa-co:electrical-ev-ready | Electrical / Energy | EV power transfer infrastructure | On or after 2024-03-01, EV power transfer infrastructure permit applicants must comply with CRS and Colorado Energy Office requirements in the Model Electric Ready and Solar Ready Code. | src:usa-co:3-ccr-710-1 | verified |
| amendment:usa-co:plumbing-local-more-stringent-reporting | Plumbing / Fuel Gas | Local more-stringent standards | A city, town, county, or city and county adopting more stringent standards than the Colorado Plumbing Code or Colorado Fuel Gas Code must furnish a copy to the State Plumbing Board. | src:usa-co:3-ccr-720-1 | verified |
| amendment:usa-co:division-of-housing-prior-code-grace | Factory-built / Tiny homes / no-local-standard housing | 180-day prior-code use | Manufacturers may use construction codes in effect before a new code amendment for up to 180 days after the amendment effective date, with possible 180-day written extension. | src:usa-co:8-ccr-1302-14 | verified_core |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-co"
  model: "local_primary_with_state_trade_and_special_scope_overlays"
  enforcing_entities:
    - "county building departments / county building inspectors where a county code is adopted"
    - "municipal building departments where a municipal code is adopted"
    - "Colorado State Electrical Board / state electrical inspection program for covered electrical work"
    - "Colorado State Plumbing Board / state or delegated plumbing inspection program for covered plumbing and fuel-gas work"
    - "Colorado Division of Fire Prevention and Control for DFPC-administered scopes"
    - "Colorado Division of Housing for factory-built, tiny-home, installation, and specified no-local-standard scopes"
  required_officials:
    - "county building inspector, when appointed by board of county commissioners"
    - "DFPC-certified or delegated inspectors for DFPC public-school and junior-college scopes"
  state_reserved_activities:
    - "electrical code adoption and minimum standards"
    - "plumbing and fuel-gas code adoption and minimum standards"
    - "DFPC-administered building, fire, and life-safety scopes"
    - "Division of Housing factory-built and specified housing scopes"
    - "state-agency energy-code duties for covered state facilities"
  source_ids:
    - "src:usa-co:crs-title-30"
    - "src:usa-co:3-ccr-710-1"
    - "src:usa-co:3-ccr-720-1"
    - "src:usa-co:8-ccr-1507-30"
    - "src:usa-co:8-ccr-1302-14"
    - "src:usa-co:session-law-2022-ch301"
  verification_status: "partially_verified"
  confidence: 0.70
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-co"
  model: "local_amendments_partially_verified"
  applies_to_code_families:
    - "Building"
    - "Residential"
    - "Energy"
    - "Plumbing"
    - "Fuel Gas"
  approval_required: "varies_by_code_family"
  approving_authority_id: "ahj:usa-co:local-building-code-authorities"
  filing_required: "verified_for_more_stringent_plumbing_fuel_gas_standards; unresolved_for_general_municipal_building_and_energy"
  registry_exists: null
  registry_source_ids: []
  legal_basis_source_ids:
    - "src:usa-co:crs-title-30"
    - "src:usa-co:3-ccr-720-1"
  verification_status: "partially_verified"
  confidence: 0.55
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement and local amendment authority should be stored as separate dimensions. A county may enforce an adopted county building code through a county building inspector, but that does not automatically identify every permissible local amendment or state approval path. Conversely, state trade codes may apply statewide while still allowing local governments to adopt more stringent plumbing/fuel-gas standards if they furnish copies to the State Plumbing Board.

### 6.4 Known Local Amendment Registries

| Registry ID | Code Family | Registry Name | Publisher | URL | Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- |
| registry:usa-co:local-building-amendments | Building / Residential | No statewide local building-amendment registry verified | unresolved | null | unresolved | src:usa-co:crs-title-30, src:usa-co:crs-title-31 | County amendment authority verified; centralized registry not verified. |
| registry:usa-co:plumbing-more-stringent | Plumbing / Fuel Gas | Local more-stringent standard copies filed with State Plumbing Board | State Plumbing Board | null | partially_verified | src:usa-co:3-ccr-720-1 | Rule requires copies to be furnished to the Board, but no public registry URL was extracted. |
| registry:usa-co:energy-local-adoptions | Energy | Local energy-code adoption registry | unresolved | null | unresolved | src:usa-co:session-law-2022-ch301, src:usa-co:crs-title-30, src:usa-co:crs-title-31 | Existence and public availability of any registry not verified. |

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendments were parsed in this pass.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: partially_designed

Jurisdiction stack:

```text
Address
  -> State: Colorado
  -> County
  -> Municipality / unincorporated county
  -> Determine ordinary building-code adopting authority:
       - municipality if inside incorporated municipality
       - county if unincorporated and county has adopted code
       - no verified local ordinary building code if neither condition is supported
  -> Apply Colorado local energy-code floor if the local government has adopted/enforced building codes
  -> Apply state trade-code overlays:
       - electrical: State Electrical Board / NEC 2023
       - plumbing/fuel gas: State Plumbing Board / Colorado Plumbing Code and Colorado Fuel Gas Code
  -> Check state special-scope overlays:
       - public schools / charter schools / junior colleges: DFPC and possible delegated local AHJ
       - factory-built / tiny home / specified no-local-standard housing: Division of Housing
       - covered state-agency construction: Office of the State Architect energy-code duty
  -> Determine local fire AHJ and operational fire-code adoption
  -> Fetch applicable local amendment records
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | Colorado statutory and CCR portals | src:usa-co:crs-title-30, src:usa-co:crs-title-31, src:usa-co:ccr-portal | statewide legal framework | when laws or rules change | partially_verified |
| County | Census TIGER/Line or Colorado state GIS source | none_selected | statewide | unknown | pending |
| Municipality | Census TIGER/Line places/incorporated areas or Colorado state GIS source | none_selected | statewide | unknown | pending |
| Fire District / Fire AHJ | DFPC/local fire department data source not selected | none_selected | statewide | unknown | pending |
| School / Junior College special scopes | DFPC program records | src:usa-co:8-ccr-1507-30 | state program scope | when program/rules change | partially_verified |
| Factory-built / no-local-standard housing scope | Division of Housing program records | src:usa-co:8-ccr-1302-14 | state program scope | when program/rules change | partially_verified |

### 7.3 AHJ Contact Data

No AHJ contact data was populated. This report identifies authority classes and legal overlays, but it does not yet provide address-level AHJ contacts.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Source Type | Publisher | URL | Accessed Date | Snapshot ID | Checksum | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-co:crs-title-30 | Colorado Revised Statutes 2024, Title 30, Government - County | statute_pdf | Colorado General Assembly / Office of Legislative Legal Services | https://content.leg.colorado.gov/sites/default/files/images/olls/crs2024-title-30.pdf | 2026-06-25 | snapshot-pending | snapshot-pending | verified |
| src:usa-co:crs-title-31 | Colorado Revised Statutes 2024, Title 31, Government - Municipal | statute_pdf | Colorado General Assembly / Office of Legislative Legal Services | https://content.leg.colorado.gov/sites/default/files/images/olls/crs2024-title-31.pdf | 2026-06-25 | snapshot-pending | snapshot-pending | verified |
| src:usa-co:session-law-2022-ch301 | Session Laws of Colorado 2022, Chapter 301, HB22-1362 | session_law_pdf | Colorado General Assembly | https://leg.colorado.gov/laws/session-laws/1691/download | 2026-06-25 | snapshot-pending | snapshot-pending | verified |
| src:usa-co:ccr-portal | Code of Colorado Regulations portal | rules_portal | Colorado Secretary of State | https://www.sos.state.co.us/CCR/Welcome.do | 2026-06-25 | snapshot-pending | snapshot-pending | verified |
| src:usa-co:dora-electrical | Electrical Board: Applications and Forms / program page | agency_page | Colorado Department of Regulatory Agencies, Division of Professions and Occupations | https://dpo.colorado.gov/Electrical/Applications | 2026-06-25 | snapshot-pending | snapshot-pending | supplemental_official |
| src:usa-co:3-ccr-710-1-rule-info | 3 CCR 710-1 State Electrical Board Rules and Regulations current rule information | ccr_rule_info_page | Colorado Secretary of State | https://www.sos.state.co.us/CCR/DisplayRule.do?action=ruleinfo&ruleId=211 | 2026-06-25 | snapshot-pending | snapshot-pending | verified |
| src:usa-co:3-ccr-710-1 | 3 CCR 710-1 State Electrical Board Rules and Regulations | ccr_rule_pdf | Colorado Secretary of State | https://www.sos.state.co.us/CCR/GenerateRulePdf.do?fileName=3+CCR+710-1&ruleVersionId=11034 | 2026-06-25 | snapshot-pending | snapshot-pending | verified |
| src:usa-co:3-ccr-720-1-rule-info | 3 CCR 720-1 Plumbing Rules and Regulations current rule information | ccr_rule_info_page | Colorado Secretary of State | https://www.sos.state.co.us/CCR/DisplayRule.do?action=ruleinfo&agencyID=33&agencyName=Division+of+Professions+and+Occupations+-+&deptID=18&deptName=Department+of+Regulatory+Agencies&ruleId=2255 | 2026-06-25 | snapshot-pending | snapshot-pending | verified |
| src:usa-co:3-ccr-720-1 | 3 CCR 720-1 Plumbing Rules and Regulations | ccr_rule_pdf | Colorado Secretary of State | https://www.sos.state.co.us/CCR/GenerateRulePdf.do?fileName=3+CCR+720-1&ruleVersionId=8778 | 2026-06-25 | snapshot-pending | snapshot-pending | verified |
| src:usa-co:8-ccr-1507-101 | 8 CCR 1507-101 Building and Fire Code Adoption and Certification of Inspectors for Fire & Life Safety Programs Administered by the State of Colorado | ccr_rule_pdf | Colorado Secretary of State / Colorado Division of Fire Prevention and Control | https://www.sos.state.co.us/CCR/GenerateRulePdf.do?fileName=8+CCR+1507-101&ruleVersionId=12302 | 2026-06-25 | snapshot-pending | snapshot-pending | verified |
| src:usa-co:8-ccr-1507-30 | 8 CCR 1507-30 Code Enforcement and Certification of Inspectors for Public Schools, Charter Schools and Junior Colleges | ccr_rule_pdf | Colorado Secretary of State / Colorado Division of Fire Prevention and Control | https://www.coloradosos.gov/CCR/GenerateRulePdf.do?fileName=8+CCR+1507-30&ruleVersionId=11809 | 2026-06-25 | snapshot-pending | snapshot-pending | verified |
| src:usa-co:8-ccr-1302-14-rule-info | 8 CCR 1302-14 current rule information | ccr_rule_info_page | Colorado Secretary of State / Colorado Division of Housing | https://www.sos.state.co.us/CCR/DisplayRule.do?action=ruleinfo&agencyID=51&agencyName=Division+of+Housing&deptID=12&deptName=Department+of+Local+Affairs&ruleId=3278&seriesNum=8+CCR+1302-14 | 2026-06-25 | snapshot-pending | snapshot-pending | verified |
| src:usa-co:8-ccr-1302-14 | 8 CCR 1302-14 Non-Residential and Residential Factory-Built Structures and Tiny Homes; Installations; and Selected No-Local-Standard Scopes | ccr_rule_pdf | Colorado Secretary of State / Colorado Division of Housing | https://www.sos.state.co.us/CCR/GenerateRulePdf.do?fileName=8+CCR+1302-14&ruleVersionId=11765 | 2026-06-25 | snapshot-pending | snapshot-pending | verified |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-co:crs-title-30 | annual_statute_pdf | Annual CRS PDF should be checked against current codified law before production release. | authoritative_for_this_report_snapshot |
| src:usa-co:crs-title-31 | annual_statute_pdf | Annual CRS PDF should be checked against current codified law before production release. | authoritative_for_this_report_snapshot |
| src:usa-co:session-law-2022-ch301 | session_law_not_current_codification | Session law is authoritative enactment history but should be reconciled against current CRS codification for final production fields. | authoritative_for_enactment_and_transition_context |
| src:usa-co:ccr-portal | portal_only | The CCR portal confirms the official publication system; individual rule PDFs govern their specific rules. | authoritative_for_rules_publication_system |
| src:usa-co:3-ccr-710-1-rule-info | rule_info_page | Rule-info page gives current effective-date metadata; the PDF governs the rule text. | authoritative_for_rule_metadata |
| src:usa-co:3-ccr-720-1-rule-info | rule_info_page | Rule-info page gives current effective-date metadata; the PDF governs the rule text. | authoritative_for_rule_metadata |
| src:usa-co:8-ccr-1302-14-rule-info | rule_info_page | Rule-info page gives current effective-date metadata; the PDF governs the rule text. | authoritative_for_rule_metadata |
| src:usa-co:3-ccr-720-1 | ccr_pdf_version_check_needed | Current rule-info page should be retained with the PDF to avoid ruleVersionId confusion in automated fetches. | authoritative_when_paired_with_rule_info_page |
| src:usa-co:8-ccr-1507-101 | special_scope_only | DFPC rule is authoritative for DFPC-administered scopes; do not treat it as universal statewide ordinary building-code adoption. | authoritative_for_dfpc_scopes_only |
| src:usa-co:8-ccr-1507-30 | special_scope_only | Public-school/charter-school/junior-college rule is authoritative for listed scopes; do not extrapolate to all fire AHJs. | authoritative_for_listed_scopes_only |
| src:usa-co:dora-electrical | agency_summary | Agency page is official but less precise than CCR rule text for code edition and effective date. | supplemental_only |

### 8.3 Supplemental Sources

No non-official supplemental sources were used in this pass.

### 8.4 Source Extraction Metadata

| Source ID | Parser | Parser Version | Extracted At | Extraction Confidence | Requires OCR | Requires Browser Automation | Manual Review Required |
| --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-co:crs-title-30 | browser_pdf_text_and_manual_review | 0.2 | 2026-06-25T00:00:00Z | 0.86 | no | yes | yes |
| src:usa-co:crs-title-31 | browser_pdf_text_and_manual_review | 0.2 | 2026-06-25T00:00:00Z | 0.84 | no | yes | yes |
| src:usa-co:session-law-2022-ch301 | browser_pdf_text_and_manual_review | 0.2 | 2026-06-25T00:00:00Z | 0.82 | no | yes | yes |
| src:usa-co:ccr-portal | browser_manual | 0.2 | 2026-06-25T00:00:00Z | 0.95 | no | yes | yes |
| src:usa-co:dora-electrical | browser_manual | 0.2 | 2026-06-25T00:00:00Z | 0.80 | no | yes | yes |
| src:usa-co:3-ccr-710-1-rule-info | browser_manual | 0.2 | 2026-06-25T00:00:00Z | 0.95 | no | yes | yes |
| src:usa-co:3-ccr-710-1 | browser_pdf_text_and_manual_review | 0.2 | 2026-06-25T00:00:00Z | 0.92 | no | yes | yes |
| src:usa-co:3-ccr-720-1-rule-info | browser_manual | 0.2 | 2026-06-25T00:00:00Z | 0.95 | no | yes | yes |
| src:usa-co:3-ccr-720-1 | browser_pdf_text_and_manual_review | 0.2 | 2026-06-25T00:00:00Z | 0.88 | no | yes | yes |
| src:usa-co:8-ccr-1507-101 | browser_pdf_text_and_manual_review | 0.2 | 2026-06-25T00:00:00Z | 0.86 | no | yes | yes |
| src:usa-co:8-ccr-1507-30 | browser_pdf_text_and_manual_review | 0.2 | 2026-06-25T00:00:00Z | 0.86 | no | yes | yes |
| src:usa-co:8-ccr-1302-14-rule-info | browser_manual | 0.2 | 2026-06-25T00:00:00Z | 0.95 | no | yes | yes |
| src:usa-co:8-ccr-1302-14 | browser_pdf_text_and_manual_review | 0.2 | 2026-06-25T00:00:00Z | 0.82 | no | yes | yes |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report:usa-co | report.status | partially_verified | verified | 1.00 | src:usa-co:crs-title-30, src:usa-co:crs-title-31, src:usa-co:3-ccr-710-1, src:usa-co:3-ccr-720-1 | Body now contains source-backed authority and adoption fields, but unresolved local and AHJ fields prevent verified status. |
| ahj:usa-co:local-building-code-authorities | authority model | local_primary_with_state_overlays | partially_verified | 0.72 | src:usa-co:crs-title-30, src:usa-co:crs-title-31 | Ordinary local building-code adoption is verified for counties and energy-code duties are verified for local governments; municipal general building-code procedure remains less fully extracted. |
| adoption:usa-co:local-energy-2023-to-2026 | base model code | 2021 IECC plus Colorado Model Electric Ready and Solar Ready Code language | verified_core | 0.88 | src:usa-co:crs-title-30, src:usa-co:crs-title-31 | Applies when local governments with building codes adopt/update building codes in statutory window. |
| adoption:usa-co:local-energy-2026-forward | base model code | Model Low Energy and Carbon Code performance threshold | verified_core | 0.88 | src:usa-co:crs-title-30, src:usa-co:crs-title-31 | Applies to local building-code updates on or after 2026-07-01, subject to timing exceptions. |
| adoption:usa-co:electrical-nec-2023 | edition | 2023 NEC | verified | 0.92 | src:usa-co:3-ccr-710-1 | Effective date 2023-08-01. |
| adoption:usa-co:plumbing-2018 | edition | 2018 IPC and IRC provisions | verified | 0.86 | src:usa-co:3-ccr-720-1, src:usa-co:3-ccr-720-1-rule-info | Paired with current rule-info page effective 2026-02-14. |
| adoption:usa-co:fuel-gas-2018 | edition | 2018 IFGC and IRC provisions | verified | 0.86 | src:usa-co:3-ccr-720-1, src:usa-co:3-ccr-720-1-rule-info | Paired with current rule-info page effective 2026-02-14. |
| local-enforcement:usa-co | model | local_primary_with_state_trade_and_special_scope_overlays | partially_verified | 0.70 | src:usa-co:crs-title-30, src:usa-co:8-ccr-1507-30 | County and DFPC special-scope enforcement verified; municipality and fire-AHJ mapping unresolved. |
| local-amendment-rule:usa-co | model | local_amendments_partially_verified | partially_verified | 0.55 | src:usa-co:crs-title-30, src:usa-co:3-ccr-720-1 | County building-code amendments and plumbing/fuel-gas more-stringent reporting verified; other local amendment dimensions unresolved. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | All `src:usa-co:*` cited in the body are listed in section 8. |
| All authority IDs resolve | pass | All authority IDs used in adoption and edge records are defined in section 2. |
| All current code families have adoption rows | pass | Rows are present for all template code families; unresolved rows remain explicit. |
| Building and operational fire code are separated | pass | DFPC construction/fire scopes are separated from local ordinary building-code adoption. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Separate date columns are preserved. |
| Effective dates are valid ISO dates | pass | Date values use ISO format where known. |
| No impossible date sequences | pass | No adoption/effective/operative contradiction identified in populated records. |
| Transition rules have explicit trigger conditions | pass | Energy-code and special-scope transition rows identify triggers. |
| Permit-date logic is captured where applicable | fail | Ordinary local building-code permit-date/application-date logic is unresolved. |
| Local enforcement model classified | pass | Classified as local-primary with state trade and special-scope overlays. |
| Local amendment rule classified | partial | County and plumbing/fuel-gas dimensions are captured; municipal and energy-amendment details remain unresolved. |
| AHJ confirmation metadata present | fail | Address-level AHJ contacts and boundary source IDs are not populated. |
| Official-source caveats captured | pass | Caveats for CCR rule PDFs, annual CRS PDFs, session law, and special-scope rules are included. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-co:001 | high | municipal building-code authority | General municipal building-code adoption and amendment procedure is not fully extracted beyond energy-code local-government provisions. | Extract municipal building-code enabling statutes and local adoption/amendment procedure from CRS Title 31 and any relevant rules. | null | null | open |
| issue:usa-co:002 | high | local fire AHJ outside DFPC scopes | DFPC school/state-program fire authority is verified, but ordinary local operational fire-code adoption and AHJ mapping are unresolved. | Identify local fire-code enabling law, fire district authority, and any statewide fire-code limits or registries. | null | null | open |
| issue:usa-co:003 | medium | local energy-code equivalency and amendments | The statutory performance floor is verified, but local amendment/equivalency review, reporting, and registry mechanics are not fully modeled. | Extract Energy Code Board/CEO/DOLA model code guidance, local adoption instructions, and any reporting requirements. | null | null | open |
| issue:usa-co:004 | medium | elevator / conveyance | Elevator and escalator code updates are referenced in energy-code timing exceptions, but the actual authority and current code edition were not extracted. | Extract Colorado elevator/conveyance code authority and current adopted editions. | null | null | open |
| issue:usa-co:005 | medium | accessibility | Accessibility-specific construction-code authority and adopted standards are unresolved. | Extract accessibility construction requirements and interaction with local building-code enforcement. | null | null | open |
| issue:usa-co:006 | medium | AHJ boundary and contact data | No boundary layers or AHJ contacts are populated. | Select county, municipal, fire, and state-program boundary/contact sources for resolver integration. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-co:crs-title-30 | src:usa-co:crs-title-30 | statute_diff | monthly | Title 30 county building-code or energy-code provisions change | 2026-06-25 | active |
| watch:usa-co:crs-title-31 | src:usa-co:crs-title-31 | statute_diff | monthly | Title 31 municipal building-code or energy-code provisions change | 2026-06-25 | active |
| watch:usa-co:session-law-energy | src:usa-co:session-law-2022-ch301 | codification_reconciliation | quarterly | energy-code session law provisions need reconciliation against current CRS | 2026-06-25 | active |
| watch:usa-co:3-ccr-710-1 | src:usa-co:3-ccr-710-1-rule-info | ccr_rule_info_diff | weekly | State Electrical Board current rule effective date or PDF changes | 2026-06-25 | active |
| watch:usa-co:3-ccr-720-1 | src:usa-co:3-ccr-720-1-rule-info | ccr_rule_info_diff | weekly | State Plumbing Board current rule effective date or PDF changes | 2026-06-25 | active |
| watch:usa-co:8-ccr-1507-101 | src:usa-co:8-ccr-1507-101 | ccr_rule_pdf_diff | weekly | DFPC building/fire code adoption rule changes | 2026-06-25 | active |
| watch:usa-co:8-ccr-1507-30 | src:usa-co:8-ccr-1507-30 | ccr_rule_pdf_diff | weekly | Public-school/junior-college code enforcement rule changes | 2026-06-25 | active |
| watch:usa-co:8-ccr-1302-14 | src:usa-co:8-ccr-1302-14-rule-info | ccr_rule_info_diff | weekly | Division of Housing factory-built/tiny-home rule changes | 2026-06-25 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-25 | Replaced framework-only Colorado draft with a source-backed partial authority and code-adoption report | report:usa-co, ahj:usa-co:local-building-code-authorities, adoption:usa-co:local-energy-2023-to-2026, adoption:usa-co:electrical-nec-2023, adoption:usa-co:plumbing-2018 | src:usa-co:crs-title-30, src:usa-co:crs-title-31, src:usa-co:3-ccr-710-1, src:usa-co:3-ccr-720-1, src:usa-co:8-ccr-1507-101, src:usa-co:8-ccr-1302-14 | ChatGPT | Upgraded report to partially_verified; unresolved local fire, municipal amendment, accessibility, elevator, and AHJ boundary issues remain explicit. |
