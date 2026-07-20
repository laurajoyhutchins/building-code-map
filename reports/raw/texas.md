---
state:
  state_id: "US-TX"
  name: "Texas"
  abbreviation: "TX"
report:
  report_id: "state-report:usa-tx"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-26"
  last_verified: null
  reviewed_by: null
risk:
  overall_confidence: 0.56 # 0.00 - 1.00
  risk_flags:
    - "hybrid_local_authority_model"
    - "county_code_scope_requires_final_codified_review"
    - "state_amendments_not_fully_parsed"
    - "local_fire_code_scope_partially_verified"
    - "ahj_contact_and_boundary_data_missing"
  open_questions_count: 6

---

# State Building Code Authority Report: Texas

## 1. Executive Summary

- **Authority model:** Texas uses a hybrid model. No single general-purpose statewide agency was verified as adopting and enforcing a complete building code for all site-built construction. Municipal baseline authority for residential and commercial buildings is in Local Government Code Chapter 214; county authority for unincorporated areas is in Local Government Code Chapter 233; energy codes, industrialized housing/buildings, electrical licensing/code, state fire-marshal inspection standards, accessibility, elevators, and coastal windstorm certification each have distinct statewide authorities.

- **Statewide code status:** Several statewide or statewide-minimum programs are verified: SECO/Comptroller energy standards; TDLR Industrialized Housing and Buildings mandatory codes; TDLR state electrical code; TDI State Fire Marshal inspection standards; TDLR accessibility standards; TDLR elevator/conveyance program; and TDI coastal windstorm adopted codes. General municipal site-built building/residential/commercial code provisions are verified at a baseline level, but full state and local amendment parsing is incomplete.

- **Local enforcement model:** Local governments remain central AHJs for most site-built construction. Municipalities administer and enforce municipal building code provisions. Local jurisdictions are responsible for implementing/enforcing the statewide energy code. Local licensing or inspecting authorities perform local electrical permitting/inspection where required. SFMO may inspect buildings statewide, but it cancels requests in areas served by local agencies with certified inspectors. Industrialized buildings and windstorm certification have state-administered components.

- **Local amendment posture:** Municipal local amendments are supported for residential, commercial, electrical, and related code administration paths, subject to source-specific limits. Energy amendments are locally permitted, but nonattainment and affected counties have stringency constraints and ESL review mechanics. Industrialized buildings have state preemption against local amendments as a prerequisite for permits or certificates of occupancy. Local amendment registries were not comprehensively identified.

- **Known transition periods or pending changes:** IHB adopted 2021 I-Codes and 2020 NEC with an effective date of 2024-07-01 and transition completion deadlines by 2024-12-31 for prior plant construction. TDLR adopted the 2023 NEC as the state electrical code effective 2023-09-01. SFMO adopted 2021 NFPA 101 and NFPA 1 as inspection standards effective 2023-09-01. TDI windstorm WPI-1 certification moves to the 2024 IRC or 2024 IBC for applications starting 2026-04-01. SECO's energy-code update process includes notice/comment and a minimum nine-month lead time after adoption.

- **Production readiness:** limited_internal_review

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: Texas is best modeled as a hybrid/local authority state rather than a single
    statewide building-code jurisdiction for all site-built projects.
  confidence: 0.6
  source_ids:
  - src:usa-tx:lg-ch214
  - src:usa-tx:lg-ch233
  - src:usa-tx:seco-local-ordinances
  - src:usa-tx:tdlr-ihb-codes
  - src:usa-tx:tdlr-electrical-compliance
  - src:usa-tx:tdi-sfmo-standards
- topic: Municipal residential/commercial baseline
  finding: Municipal residential and commercial baseline codes are tied to IRC/IBC
    statutory provisions; edition references changed to codes as they existed on 2012-05-01.
  confidence: 0.55
  source_ids:
  - src:usa-tx:lg-ch214
  - src:usa-tx:capitol-hb0738-2021
  - src:usa-tx:supp-findlaw-lg214-212
  - src:usa-tx:supp-justia-lg214-216
- topic: Energy code
  finding: Statewide energy minimums are Chapter 11 of the 2015 IRC for single-family
    homes and the 2015 IECC for other residential, commercial, and industrial buildings;
    local implementation/enforcement remains local.
  confidence: 0.85
  source_ids:
  - src:usa-tx:seco-local-ordinances
  - src:usa-tx:seco-commercial-energy
- topic: Industrialized housing/buildings
  finding: IHB mandatory codes effective 2024-07-01 include 2021 IBC, 2021 IRC, 2021
    IFGC, 2021 IMC, 2021 IPC, 2015 IECC, 2021 IEBC, and 2020 NEC.
  confidence: 0.9
  source_ids:
  - src:usa-tx:tdlr-ihb-codes
  - src:usa-tx:tdlr-ihb-adoption-2024
- topic: Electrical code authority
  finding: TDLR adopted the 2023 NEC as the state electrical code effective 2023-09-01,
    with permit-date and start-of-work date logic.
  confidence: 0.9
  source_ids:
  - src:usa-tx:tdlr-electrical-compliance
  - src:usa-tx:tdlr-electrical-laws-rules
- topic: Fire code / inspection standard
  finding: SFMO uses 2021 NFPA 101 and NFPA 1 as inspection standards effective 2023-09-01;
    local fire-code adoption and local operational fire-prevention codes remain only
    partially mapped.
  confidence: 0.78
  source_ids:
  - src:usa-tx:tdi-sfmo-standards
  - src:usa-tx:tdi-sfmo-rules
  - src:usa-tx:tdi-sfmo-inspections
- topic: Coastal windstorm
  finding: TDI windstorm certification requires WPI-1 applications starting 2026-04-01
    to be certified under the 2024 IRC or 2024 IBC.
  confidence: 0.88
  source_ids:
  - src:usa-tx:tdi-windstorm-adopted-codes
- topic: Local amendments
  finding: Local amendment authority is confirmed for several paths, but no unified
    statewide local-amendment registry was verified.
  confidence: 0.45
  source_ids:
  - src:usa-tx:seco-local-ordinances
  - src:usa-tx:seco-adoption-process
  - src:usa-tx:lg-ch214
  - src:usa-tx:tdlr-ihb-faq
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-tx:hybrid-state-local |
| Authority name | Hybrid Texas state/local construction-code authority model |
| Authority type | hybrid_state_local_model |
| Legal basis | Municipal code authority: Texas Local Government Code Chapter 214. County/unincorporated authority: Texas Local Government Code Chapter 233. Statewide specialized programs are administered through SECO/Comptroller, TDLR, TDI/SFMO, and TDI windstorm program statutes/rules. |
| Role | Coordinates multiple independent authority streams rather than one statewide general building-code adoption body. |
| Enforcement model | Local enforcement for most site-built construction; state-administered or state-standardized enforcement for specialized programs. |
| Source IDs | src:usa-tx:lg-ch214; src:usa-tx:lg-ch233; src:usa-tx:seco-local-ordinances; src:usa-tx:tdlr-ihb-codes; src:usa-tx:tdlr-electrical-laws-rules; src:usa-tx:tdi-sfmo-rules; src:usa-tx:tdi-windstorm-adopted-codes |
| Verification status | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-tx:municipal-building; ahj:usa-tx:county-construction; ahj:usa-tx:tdlr-ihb; ahj:usa-tx:tdi-windstorm | Municipalities; counties for unincorporated areas; TDLR IHB; TDI windstorm | Municipal commercial building baseline; county unincorporated construction authority; state modular/IHB code; windstorm certification in designated areas | Local Government Code Ch. 214; Local Government Code Ch. 233; Occupations Code Ch. 1202 / 16 TAC Ch. 70; Insurance Code/windstorm program sources | src:usa-tx:lg-ch214; src:usa-tx:lg-ch233; src:usa-tx:tdlr-ihb-codes; src:usa-tx:tdi-windstorm-adopted-codes | partially_verified |
| Residential | ahj:usa-tx:municipal-residential; ahj:usa-tx:county-construction; ahj:usa-tx:tdlr-ihb; ahj:usa-tx:tdi-windstorm | Municipalities; counties for unincorporated areas; TDLR IHB; TDI windstorm | Municipal residential baseline; county unincorporated residential construction authority; IHB; coastal windstorm certification | Local Government Code Ch. 214; Local Government Code Ch. 233; 16 TAC Ch. 70; TDI windstorm adopted codes | src:usa-tx:lg-ch214; src:usa-tx:capitol-hb0738-2021; src:usa-tx:tdlr-ihb-codes; src:usa-tx:tdi-windstorm-adopted-codes | partially_verified |
| Existing Building / Rehabilitation | ahj:usa-tx:tdlr-ihb | TDLR Industrialized Housing and Buildings | Verified only for IHB where 2021 IEBC is a mandatory code; general statewide site-built existing-building adoption not parsed | 16 TAC Ch. 70 / IHB mandatory building codes | src:usa-tx:tdlr-ihb-codes; src:usa-tx:tdlr-ihb-adoption-2024 | partial |
| Mechanical | ahj:usa-tx:tdlr-ihb | TDLR Industrialized Housing and Buildings | Verified for IHB only; general site-built statewide mechanical code not resolved | 16 TAC Ch. 70 / IHB mandatory building codes | src:usa-tx:tdlr-ihb-codes | partial |
| Plumbing | ahj:usa-tx:tdlr-ihb | TDLR Industrialized Housing and Buildings | Verified for IHB only; general site-built statewide plumbing code not resolved | 16 TAC Ch. 70 / IHB mandatory building codes | src:usa-tx:tdlr-ihb-codes | partial |
| Fuel Gas | ahj:usa-tx:tdlr-ihb | TDLR Industrialized Housing and Buildings | Verified for IHB only; general site-built statewide fuel-gas code not resolved | 16 TAC Ch. 70 / IHB mandatory building codes | src:usa-tx:tdlr-ihb-codes | partial |
| Electrical | ahj:usa-tx:tdlr-electrical; ahj:usa-tx:municipal-electrical | TDLR; local licensing/inspection authorities | State electrical code adoption and local inspection/permitting where applicable | Occupations Code Ch. 1305; 16 TAC Ch. 73; municipal electrical statutory provisions | src:usa-tx:tdlr-electrical-compliance; src:usa-tx:tdlr-electrical-laws-rules; src:usa-tx:lg-ch214 | partially_verified |
| Energy | ahj:usa-tx:seco | State Energy Conservation Office, Texas Comptroller of Public Accounts | Statewide energy-code standards, adoption process, local ordinance constraints, and state-funded building standards | Health and Safety Code Ch. 388; 34 TAC Ch. 19 references; SECO guidance | src:usa-tx:seco-local-ordinances; src:usa-tx:seco-commercial-energy; src:usa-tx:seco-state-funded; src:usa-tx:seco-adoption-process | verified_core |
| Fire - construction references | ahj:usa-tx:tdi-sfmo; ahj:usa-tx:local-fire-ahj | TDI State Fire Marshal; local fire authorities | SFMO inspection standards affect design/construction/operation where SFMO is AHJ; local fire code scope remains partially mapped | Government Code Ch. 417; 28 TAC Ch. 34; Local Government Code Ch. 233 fire-code provisions | src:usa-tx:tdi-sfmo-standards; src:usa-tx:tdi-sfmo-rules; src:usa-tx:lg-ch233 | partially_verified |
| Fire - operational / prevention code | ahj:usa-tx:tdi-sfmo; ahj:usa-tx:local-fire-ahj | TDI State Fire Marshal; local fire authorities | NFPA 101 and NFPA 1 inspection standards for SFMO; local AHJs in many areas | Government Code Ch. 417; 28 TAC Ch. 34; local fire-code authorities | src:usa-tx:tdi-sfmo-standards; src:usa-tx:tdi-sfmo-inspections; src:usa-tx:lg-ch233 | partially_verified |
| Accessibility | ahj:usa-tx:tdlr-accessibility | TDLR Architectural Barriers Program | Texas Accessibility Standards registration/review and compliance program | Government Code Ch. 469; 16 TAC Ch. 68; 2012 TAS | src:usa-tx:tdlr-tas; src:usa-tx:tdlr-ab-laws-rules | partially_verified |
| Elevator / Conveyance | ahj:usa-tx:tdlr-elevator | TDLR Elevator/Escalator Safety Program | Elevator and escalator safety law/rules; exact adopted ASME edition not parsed | Health and Safety Code Ch. 754; 16 TAC Ch. 74 | src:usa-tx:tdlr-elevator-laws-rules | partially_verified |

### 2.3 Authority Hierarchy Notes

Texas should be represented as a set of parallel and overlapping authority streams:

1. **Municipal site-built code authority:** Local Government Code Chapter 214 establishes municipal residential, commercial, electrical, and related authority streams. Municipalities may administer and enforce their codes and adopt local amendments, subject to statutory and program-specific limits.
2. **County/unincorporated authority:** Local Government Code Chapter 233 establishes county authority for selected unincorporated-area codes, including fire-code authority and later residential/commercial construction-code provisions. The county construction-code text requires final codified review because one source in this pass was an official enrolled/amendment-style publication rather than a fully parsed current-code page.
3. **Statewide energy minimums:** SECO/Comptroller guidance states the statewide energy-code minimums and local ordinance constraints; local jurisdictions implement and enforce.
4. **Statewide specialty programs:** TDLR IHB, TDLR electrical, TDLR architectural barriers, TDLR elevator, TDI/SFMO inspections, and TDI windstorm certification each operate under separate legal authority.
5. **Local fire AHJ interaction:** SFMO conducts inspections statewide within its authority, but its own inspection page distinguishes areas served by local agencies with certified inspectors.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-tx:001 | ahj:usa-tx:hybrid-state-local | includes | ahj:usa-tx:municipal-building / municipal site-built code authority | src:usa-tx:lg-ch214 | partially_verified |
| edge:usa-tx:002 | ahj:usa-tx:hybrid-state-local | includes | ahj:usa-tx:county-construction / unincorporated county authority | src:usa-tx:lg-ch233 | partial |
| edge:usa-tx:003 | ahj:usa-tx:seco | establishes_minimum_for | statewide energy codes; local implementation/enforcement | src:usa-tx:seco-local-ordinances; src:usa-tx:seco-adoption-process | verified_core |
| edge:usa-tx:004 | ahj:usa-tx:tdlr-ihb | preempts_or_limits | local amendments as prerequisite to IHB permits or certificates of occupancy | src:usa-tx:tdlr-ihb-faq | partial |
| edge:usa-tx:005 | ahj:usa-tx:tdlr-electrical | sets_state_code_for | electrical work subject to state electrician law; local authorities inspect/permit where required | src:usa-tx:tdlr-electrical-compliance; src:usa-tx:tdlr-electrical-laws-rules | verified_core |
| edge:usa-tx:006 | ahj:usa-tx:tdi-sfmo | inspects_using | 2021 NFPA 101 and NFPA 1 inspection standards | src:usa-tx:tdi-sfmo-standards; src:usa-tx:tdi-sfmo-inspections | verified_core |
| edge:usa-tx:007 | ahj:usa-tx:tdi-windstorm | certifies_under | 2024 IRC / 2024 IBC for WPI-1 applications beginning 2026-04-01 | src:usa-tx:tdi-windstorm-adopted-codes | verified_core |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | Municipal commercial building code baseline; IHB mandatory building code; TDI windstorm adopted codes | IBC | Municipal commercial: IBC as existed on 2012-05-01; IHB: 2021 IBC; windstorm: 2024 IBC beginning 2026-04-01 | partially_verified | null | IHB: 2024-07-01; windstorm: 2026-04-01 | null | IHB: 2024-07-01 for new design/construction; windstorm WPI-1: 2026-04-01 | IHB prior plant construction transition through 2024-12-31; WPI-1 applications on/after 2026-04-01 use 2024 IRC/IBC | src:usa-tx:lg-ch214; src:usa-tx:capitol-hb0738-2021; src:usa-tx:tdlr-ihb-codes; src:usa-tx:tdlr-ihb-adoption-2024; src:usa-tx:tdi-windstorm-adopted-codes |
| Residential | Municipal residential building code baseline; IHB mandatory residential code; TDI windstorm adopted codes | IRC | Municipal residential: IRC as existed on 2012-05-01; IHB: 2021 IRC; windstorm: 2024 IRC beginning 2026-04-01 | partially_verified | null | IHB: 2024-07-01; windstorm: 2026-04-01 | null | IHB: 2024-07-01 for new design/construction; windstorm WPI-1: 2026-04-01 | IHB transition through 2024-12-31 for prior plant construction; WPI-1 applications on/after 2026-04-01 use 2024 IRC/IBC | src:usa-tx:lg-ch214; src:usa-tx:capitol-hb0738-2021; src:usa-tx:tdlr-ihb-codes; src:usa-tx:tdi-windstorm-adopted-codes |
| Existing Building / Rehabilitation | IHB mandatory existing building code | IEBC | 2021 | partial; IHB only | 2024-05-21 | 2024-07-01 | 2024-07-01 | 2024-07-01 for covered IHB design/construction | Prior construction completed/inspected/labeled by 2024-12-31 may remain under prior code path | src:usa-tx:tdlr-ihb-codes; src:usa-tx:tdlr-ihb-adoption-2024 |
| Mechanical | IHB mandatory mechanical code | IMC | 2021 | partial; IHB only | 2024-05-21 | 2024-07-01 | 2024-07-01 | 2024-07-01 for covered IHB design/construction | Same IHB transition rule | src:usa-tx:tdlr-ihb-codes; src:usa-tx:tdlr-ihb-adoption-2024 |
| Plumbing | IHB mandatory plumbing code | IPC | 2021 | partial; IHB only | 2024-05-21 | 2024-07-01 | 2024-07-01 | 2024-07-01 for covered IHB design/construction | Same IHB transition rule | src:usa-tx:tdlr-ihb-codes; src:usa-tx:tdlr-ihb-adoption-2024 |
| Fuel Gas | IHB mandatory fuel gas code | IFGC | 2021 | partial; IHB only | 2024-05-21 | 2024-07-01 | 2024-07-01 | 2024-07-01 for covered IHB design/construction | Same IHB transition rule | src:usa-tx:tdlr-ihb-codes; src:usa-tx:tdlr-ihb-adoption-2024 |
| Electrical | Texas state electrical code; IHB mandatory electrical code | NEC | State electrical code: 2023 NEC as existed 2022-08-25; IHB: 2020 NEC | partially_verified | null | 2023-09-01 for state electrical code; 2024-07-01 for IHB NEC 2020 | 2023-09-01 for state electrical code | Permit-date or start-of-work date controls under TDLR guidance | If permit required, code in effect on permit date; if no permit required, code in effect when electrician begins installing materials/equipment | src:usa-tx:tdlr-electrical-compliance; src:usa-tx:tdlr-electrical-laws-rules; src:usa-tx:tdlr-ihb-codes |
| Energy | Statewide energy code; state-funded building energy standards; IHB energy code | 2015 IRC Ch. 11; 2015 IECC; ASHRAE 90.1-2016 / 2018 IECC for state-funded buildings; IHB 2015 IECC | 2015 IRC Chapter 11 for single-family homes; 2015 IECC for other residential/commercial/industrial; state-funded: ASHRAE 90.1-2016 or 2018 IECC; IHB: 2015 IECC | verified_core | null | Commercial/non-low-rise: 2016-11-01; state-funded revisions: 2021-01-01; IHB: 2024-07-01 | null | Statewide minimums active; state-funded design assignments on/after 2021-01-01 | SECO future adoptions cannot be effective earlier than nine months after adoption | src:usa-tx:seco-local-ordinances; src:usa-tx:seco-commercial-energy; src:usa-tx:seco-state-funded; src:usa-tx:seco-adoption-process; src:usa-tx:tdlr-ihb-codes |
| Fire - construction references | SFMO inspection standards; county fire-code authority; IHB fire/building references | NFPA 101; NFPA 1; IFC/UFC for county fire-code option; IBC/IRC/IFC references by program | SFMO: 2021 NFPA 101 and 2021 NFPA 1; county fire code: IFC or UFC as existed 2005-05-01; IHB: 2021 IBC/IRC with related fire provisions | partially_verified | null | SFMO: 2023-09-01; IHB: 2024-07-01 | null | SFMO standards active for inspections within SFMO authority | SFMO standard applies in inspections; local AHJ path may supersede request processing in areas served by local certified inspectors | src:usa-tx:tdi-sfmo-standards; src:usa-tx:tdi-sfmo-inspections; src:usa-tx:lg-ch233; src:usa-tx:supp-findlaw-lg233-062 |
| Fire - operational / prevention code | SFMO inspection standards and local/county fire code authorities | NFPA 101; NFPA 1; IFC/UFC local/county paths | SFMO: 2021 NFPA 101 and 2021 NFPA 1 | partially_verified | null | 2023-09-01 | null | Applies when SFMO conducts inspections within its authority | Local operational fire-code adoptions not comprehensively cataloged | src:usa-tx:tdi-sfmo-standards; src:usa-tx:tdi-sfmo-rules; src:usa-tx:tdi-sfmo-inspections; src:usa-tx:lg-ch233 |
| Accessibility | Texas Accessibility Standards | TAS / ADA-aligned accessibility standards | 2012 TAS | partially_verified | null | 2012-03-15 | 2012-03-15 | 2012-03-15 | Registration/review thresholds and compliance obligations vary by project cost/scope | src:usa-tx:tdlr-tas; src:usa-tx:tdlr-ab-laws-rules; src:usa-tx:tdlr-ab-faq |
| Elevator / Conveyance | Elevator and Escalator Safety Program | ASME A17.1 and related conveyance standards | exact edition unresolved | partial | null | null | null | null | Exact edition and transition text require rule-level extraction | src:usa-tx:tdlr-elevator-laws-rules; src:usa-tx:tdlr-elevator-faq |

### 3.2 Adoption Records

#### adoption:usa-tx:municipal-residential-irc

| Field | Value |
| --- | --- |
| Code family | Residential |
| Authority | ahj:usa-tx:municipal-residential |
| Base model code | International Residential Code |
| Edition / version | IRC as existed on 2012-05-01 |
| Adoption date | null |
| Effective date | null |
| Operative date | null |
| Mandatory date | null |
| Applicability | Residential structures in municipalities; source details indicate construction, alteration, remodeling, enlargement, and repair scope. |
| Source IDs | src:usa-tx:lg-ch214; src:usa-tx:capitol-hb0738-2021; src:usa-tx:supp-findlaw-lg214-212 |
| Verification status | partially_verified |
| Notes | Edition reference is supported; final current statutory text should be re-extracted before marking verified. |

#### adoption:usa-tx:municipal-commercial-ibc

| Field | Value |
| --- | --- |
| Code family | Building |
| Authority | ahj:usa-tx:municipal-building |
| Base model code | International Building Code |
| Edition / version | IBC as existed on 2012-05-01 |
| Adoption date | null |
| Effective date | null |
| Operative date | null |
| Mandatory date | null |
| Applicability | Commercial buildings in municipalities; source details indicate construction, alteration, remodeling, enlargement, and repair scope. |
| Source IDs | src:usa-tx:lg-ch214; src:usa-tx:capitol-hb0738-2021; src:usa-tx:supp-justia-lg214-216 |
| Verification status | partially_verified |
| Notes | Edition reference is supported; final current statutory text should be re-extracted before marking verified. |

#### adoption:usa-tx:tdlr-electrical-2023-nec

| Field | Value |
| --- | --- |
| Code family | Electrical |
| Authority | ahj:usa-tx:tdlr-electrical |
| Base model code | National Electrical Code |
| Edition / version | 2023 NEC as it existed on 2022-08-25 |
| Adoption date | null |
| Effective date | 2023-09-01 |
| Operative date | 2023-09-01 |
| Mandatory date | 2023-09-01 |
| Applicability | Electrical installations subject to the state electrical code and local permit/inspection rules. |
| Source IDs | src:usa-tx:tdlr-electrical-compliance; src:usa-tx:tdlr-electrical-laws-rules |
| Verification status | verified_core |
| Notes | TDLR guidance separates permit-required projects from projects where no permit is required. |

#### adoption:usa-tx:seco-statewide-energy-2015

| Field | Value |
| --- | --- |
| Code family | Energy |
| Authority | ahj:usa-tx:seco |
| Base model code | 2015 IRC Chapter 11 and 2015 IECC |
| Edition / version | 2015 |
| Adoption date | null |
| Effective date | 2016-11-01 for commercial, industrial, and residential buildings taller than three stories; single-family date not extracted from the official page in this pass |
| Operative date | null |
| Mandatory date | active statewide minimum |
| Applicability | Single-family homes use Chapter 11 of the 2015 IRC. Other residential buildings and commercial/industrial buildings use the 2015 IECC. |
| Source IDs | src:usa-tx:seco-local-ordinances; src:usa-tx:seco-commercial-energy |
| Verification status | verified_core |
| Notes | Local jurisdictions implement and enforce the energy code. |

#### adoption:usa-tx:seco-state-funded-2021

| Field | Value |
| --- | --- |
| Code family | Energy |
| Authority | ahj:usa-tx:seco |
| Base model code | ASHRAE 90.1-2016 or 2018 IECC; 2018 IECC residential for low-rise state-funded residential |
| Edition / version | 2016 / 2018 |
| Adoption date | null |
| Effective date | 2021-01-01 |
| Operative date | 2021-01-01 |
| Mandatory date | 2021-01-01 for covered state-funded design assignments |
| Applicability | State agencies and institutions of higher education for new construction and major renovation projects. |
| Source IDs | src:usa-tx:seco-state-funded |
| Verification status | verified_core |
| Notes | Requirements apply to design assignments on or after the effective date. |

#### adoption:usa-tx:tdlr-ihb-2024

| Field | Value |
| --- | --- |
| Code family | Building; Residential; Existing Building; Mechanical; Plumbing; Fuel Gas; Energy; Electrical |
| Authority | ahj:usa-tx:tdlr-ihb |
| Base model code | 2021 IBC; 2021 IRC; 2021 IFGC; 2021 IMC; 2021 IPC; 2015 IECC; 2021 IEBC; 2020 NEC |
| Edition / version | listed code editions |
| Adoption date | 2024-05-21 for Commission rule adoption; 2023-11-16 for IHB Council approval |
| Effective date | 2024-07-01 |
| Operative date | 2024-07-01 |
| Mandatory date | 2024-07-01 for covered design/construction after that date |
| Applicability | Industrialized housing and buildings, modules, and modular components. |
| Source IDs | src:usa-tx:tdlr-ihb-codes; src:usa-tx:tdlr-ihb-adoption-2024 |
| Verification status | verified_core |
| Notes | Prior plant construction before 2024-07-01 must be completed, inspected, and labeled by 2024-12-31 for prior code path treatment. |

#### adoption:usa-tx:sfmo-nfpa-2021

| Field | Value |
| --- | --- |
| Code family | Fire - construction references; Fire - operational / prevention code |
| Authority | ahj:usa-tx:tdi-sfmo |
| Base model code | NFPA 101 Life Safety Code; NFPA 1 Fire Code |
| Edition / version | 2021 |
| Adoption date | null |
| Effective date | 2023-09-01 |
| Operative date | 2023-09-01 |
| Mandatory date | 2023-09-01 for SFMO inspection standard use |
| Applicability | SFMO inspections and design/construction/operation determinations where SFMO is AHJ or has inspection authority. |
| Source IDs | src:usa-tx:tdi-sfmo-standards; src:usa-tx:tdi-sfmo-rules |
| Verification status | verified_core |
| Notes | Local AHJ scope and local fire-code editions require jurisdiction-by-jurisdiction handling. |

#### adoption:usa-tx:tdi-windstorm-2024

| Field | Value |
| --- | --- |
| Code family | Building; Residential; windstorm certification |
| Authority | ahj:usa-tx:tdi-windstorm |
| Base model code | 2024 International Residential Code or 2024 International Building Code |
| Edition / version | 2024 |
| Adoption date | null |
| Effective date | 2026-04-01 |
| Operative date | 2026-04-01 |
| Mandatory date | 2026-04-01 for WPI-1 applications |
| Applicability | Windstorm Certificate of Compliance applications in designated catastrophe areas for TWIA eligibility. |
| Source IDs | src:usa-tx:tdi-windstorm-adopted-codes |
| Verification status | verified_core |
| Notes | This is a windstorm certification program, not a general statewide building-code adoption for all projects. |

#### adoption:usa-tx:tdlr-tas-2012

| Field | Value |
| --- | --- |
| Code family | Accessibility |
| Authority | ahj:usa-tx:tdlr-accessibility |
| Base model code | Texas Accessibility Standards |
| Edition / version | 2012 TAS |
| Adoption date | null |
| Effective date | 2012-03-15 |
| Operative date | 2012-03-15 |
| Mandatory date | 2012-03-15 |
| Applicability | Covered buildings/facilities under the Elimination of Architectural Barriers program. |
| Source IDs | src:usa-tx:tdlr-tas; src:usa-tx:tdlr-ab-laws-rules; src:usa-tx:tdlr-ab-faq |
| Verification status | partially_verified |
| Notes | Registration/review thresholds were not fully converted into normalized applicability records. |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

Texas date logic is program-specific. IHB has an explicit 2024-07-01 effective date and 2024-12-31 completion/inspection/label deadline for prior plant construction. Electrical work uses permit-date logic when a permit is required and start-of-work logic when no permit is required. SECO energy-code adoption uses notice/comment and a rule that new code effective dates cannot be earlier than nine months after adoption. TDI windstorm WPI-1 applications move to 2024 IRC/IBC on 2026-04-01. Several municipal and county statutory dates remain unresolved at the field level.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-tx:ihb-2024-effective | Industrialized housing/buildings | effective_date | 2024-07-01 | Design packages approved on/after 2024-07-01 or construction started on/after 2024-07-01 | No, except transition cases | src:usa-tx:tdlr-ihb-codes; src:usa-tx:tdlr-ihb-adoption-2024 | verified_core |
| date-rule:usa-tx:ihb-2024-transition | Industrialized housing/buildings | transition_deadline | 2024-12-31 | Plant construction began before 2024-07-01 and follows prior mandatory code path | Yes, if completed, inspected, and labeled by 2024-12-31 | src:usa-tx:tdlr-ihb-codes; src:usa-tx:tdlr-ihb-adoption-2024 | verified_core |
| date-rule:usa-tx:electrical-permit | Electrical installations | permit_date | permit date | Permit is required by local licensing/inspection authority | Code in effect at permit issuance/obtainment controls | src:usa-tx:tdlr-electrical-compliance | verified_core |
| date-rule:usa-tx:electrical-no-permit | Electrical installations | start_of_work_date | date electrician begins installing materials/equipment | No permit is required | Code in effect on start-of-work date controls | src:usa-tx:tdlr-electrical-compliance | verified_core |
| date-rule:usa-tx:seco-new-energy-code | Future SECO energy-code adoptions | minimum_lead_time | not earlier than nine months after adoption | SECO adopts a new code edition after ESL analysis and notice/comment | Prior code allowed until effective date | src:usa-tx:seco-adoption-process | verified_core |
| date-rule:usa-tx:windstorm-2024 | TDI windstorm WPI-1 certification | application_date | 2026-04-01 | WPI-1 application starts on/after 2026-04-01 | Prior code path not parsed for in-flight applications | src:usa-tx:tdi-windstorm-adopted-codes | verified_core |
| date-rule:usa-tx:sfmo-2021-nfpa | SFMO inspection standards | effective_date | 2023-09-01 | SFMO inspection within its authority | No general transition rule parsed | src:usa-tx:tdi-sfmo-standards | verified_core |
| date-rule:usa-tx:municipal-irc-ibc-2012 | Municipal IRC/IBC statutory edition references | edition_reference_date | 2012-05-01 | Municipal residential/commercial statutory baseline code reference | Transition and local grace-period logic unresolved | src:usa-tx:capitol-hb0738-2021; src:usa-tx:lg-ch214 | partial |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building / Residential / Windstorm | 2024 IRC / 2024 IBC for WPI-1 certification | null | null | 2026-04-01 | 2026-04-01 | 2026-04-01 | active_watch | src:usa-tx:tdi-windstorm-adopted-codes | Applies to windstorm certification, not statewide general building code. |
| Energy | Future IECC / ASHRAE updates | null | null | null | null | null | active_watch | src:usa-tx:seco-adoption-process | SECO adoption process and ESL stringency analysis should be monitored. |
| County residential/commercial construction | IRC/IBC as existed 2018-01-01 for certain unincorporated county construction | null | null | 2025-09-01 | 2025-09-01 | 2025-09-01 | needs_final_codified_review | src:usa-tx:lg-ch233; src:usa-tx:capitol-sb2-2025-amendment | Official legislative publication indicates a 2025 change; codified text must be checked before production use. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-tx:energy-single-family | Energy | Single-family homes | Residential single-family construction | State minimum is Chapter 11 of the 2015 IRC. | src:usa-tx:seco-local-ordinances | verified_core |
| applicability-rule:usa-tx:energy-other-buildings | Energy | Other residential, commercial, industrial buildings | Building is not a single-family home in the SECO local-ordinance description | State minimum is 2015 IECC. | src:usa-tx:seco-local-ordinances; src:usa-tx:seco-commercial-energy | verified_core |
| applicability-rule:usa-tx:state-funded-energy | Energy | State agency and institution of higher education projects | New construction or major renovation | ASHRAE 90.1-2016 is the minimum energy standard, with 2018 IECC alternative; low-rise residential state-funded buildings use the 2018 IECC residential chapter. | src:usa-tx:seco-state-funded | verified_core |
| applicability-rule:usa-tx:ihb | Multiple | Industrialized housing/buildings, modules, modular components | Covered IHB program construction | 2021 I-Codes/2020 NEC/2015 IECC mandatory code set applies. | src:usa-tx:tdlr-ihb-codes | verified_core |
| applicability-rule:usa-tx:sfmo-local-agency | Fire | Fire-safety inspection request | Building is in an area served by a local agency employing certified inspectors | SFMO inspection page indicates requests are canceled and applicant is directed to local fire inspection authority. | src:usa-tx:tdi-sfmo-inspections | verified_core |
| applicability-rule:usa-tx:windstorm | Building / Residential / Windstorm | Windstorm Certificate of Compliance | Location in designated catastrophe area and TWIA eligibility path | WPI-1 applications on/after 2026-04-01 must be certified under 2024 IRC or 2024 IBC. | src:usa-tx:tdi-windstorm-adopted-codes | verified_core |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** mixed_program_model

**Where amendments are published:** Statutes, agency rules, agency mandatory-code pages, state register/rulemaking notices, and program-specific guidance pages. IHB amendments point to 16 TAC 70.101; energy amendments point to SECO/34 TAC Chapter 19 references and local ordinance rules; SFMO standards point to Government Code Chapter 417 and 28 TAC Chapter 34; electrical rules point to 16 TAC Chapter 73.

**Amendment parsing status:** partial

### 5.2 State Amendment Sources

| Amendment Source ID | Code Family | Publication Path | Coverage | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| amendset:usa-tx:ihb-70-101 | IHB multiple code families | 16 TAC 70.101 and TDLR mandatory-code guidance | IHB state amendments to adopted model codes | src:usa-tx:tdlr-ihb-codes; src:usa-tx:tdlr-ihb-adoption-2024 | identified_not_fully_parsed |
| amendset:usa-tx:energy-34-tac-19 | Energy | SECO/Comptroller energy code pages and 34 TAC Chapter 19 references | Statewide minimum energy codes and state-funded building standards | src:usa-tx:seco-commercial-energy; src:usa-tx:seco-state-funded; src:usa-tx:seco-adoption-process | identified_partial |
| amendset:usa-tx:sfmo-28-tac-34 | Fire | 28 TAC Chapter 34 and SFMO inspection standard guidance | NFPA inspection standards, SFMO deviations/approval path | src:usa-tx:tdi-sfmo-standards; src:usa-tx:tdi-sfmo-rules | identified_partial |
| amendset:usa-tx:electrical-16-tac-73 | Electrical | 16 TAC Chapter 73 and TDLR compliance guidance | State electrical code adoption and permit/start-of-work date rules | src:usa-tx:tdlr-electrical-compliance; src:usa-tx:tdlr-electrical-laws-rules | identified_partial |
| amendset:usa-tx:windstorm-tdi | Windstorm | TDI adopted building-code and windstorm certification guidance | Certification under 2024 IRC/IBC for designated catastrophe areas | src:usa-tx:tdi-windstorm-adopted-codes | identified_partial |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Impact Area | Summary | Source IDs | Confidence | Status |
| --- | --- | --- | --- | --- | --- | --- |
| amend:usa-tx:ihb-local-preemption | IHB | local permit / certificate prerequisites | TDLR guidance indicates municipalities may not require or enforce local amendments to the IHB mandatory state codes as a prerequisite for permits or certificates of occupancy for compliant industrialized buildings. | src:usa-tx:tdlr-ihb-faq | 0.70 | partial |
| amend:usa-tx:energy-nonattainment-stringency | Energy | local amendments | SECO guidance states local jurisdictions may adopt amendments, but nonattainment and affected counties must satisfy energy-code stringency rules and may use ESL determinations. | src:usa-tx:seco-local-ordinances; src:usa-tx:seco-adoption-process | 0.85 | verified_core |
| amend:usa-tx:sfmo-lsc-prevails | Fire | inspection/design hierarchy | SFMO guidance states that when other codes are used for design elements, the Life Safety Code prevails and deviations must be approved by SFMO as AHJ or cited as a deficiency. | src:usa-tx:tdi-sfmo-standards | 0.82 | verified_core |
| amend:usa-tx:windstorm-2024-wbd | Windstorm | windborne debris | TDI guidance states windborne debris protection requirements are as specified in the 2024 IRC/IBC for the windstorm certification path. | src:usa-tx:tdi-windstorm-adopted-codes | 0.80 | verified_core |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-tx"
  model: "local_primary_with_state_specialty_programs"
  enforcing_entities:
    - "municipal building departments and local AHJs for municipal site-built construction"
    - "county authorities where Local Government Code Chapter 233 applies"
    - "local jurisdictions for energy-code implementation and enforcement"
    - "local licensing or inspecting authorities for electrical permits and inspections where required"
    - "TDI State Fire Marshal for inspections within SFMO authority, except areas served by local agencies with certified inspectors"
    - "TDLR for Industrialized Housing and Buildings program functions"
    - "TDLR for accessibility and elevator/conveyance program functions"
    - "TDI windstorm program / appointed engineers / inspectors for windstorm certificate path"
  required_officials:
    - "local building official or AHJ, jurisdiction-specific"
    - "local fire inspection authority, jurisdiction-specific"
    - "TDLR or approved IHB program participants, program-specific"
    - "TDI/TWIA windstorm certification participants, program-specific"
  state_reserved_activities:
    - "IHB plan review/approval and mandatory code program elements"
    - "state electrical code adoption"
    - "state energy-code determination/adoption process"
    - "SFMO inspection standards and inspections within SFMO authority"
    - "TAS architectural-barriers registration/review program"
    - "elevator/escalator safety program"
    - "windstorm certificate-of-compliance program in designated catastrophe areas"
  source_ids:
    - "src:usa-tx:lg-ch214"
    - "src:usa-tx:lg-ch233"
    - "src:usa-tx:seco-local-ordinances"
    - "src:usa-tx:tdlr-electrical-compliance"
    - "src:usa-tx:tdi-sfmo-inspections"
    - "src:usa-tx:tdlr-ihb-codes"
    - "src:usa-tx:tdi-windstorm-adopted-codes"
  verification_status: "partially_verified"
  confidence: 0.62
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-tx"
  model: "permitted_by_program_with_limits"
  applies_to_code_families:
    - "municipal residential and commercial codes"
    - "municipal electrical code provisions"
    - "energy code local ordinances subject to stringency constraints in nonattainment/affected counties"
    - "IHB limited/preempted against local amendments as permit or certificate prerequisite"
    - "local fire codes, jurisdiction-specific and not fully cataloged"
  approval_required: "program_specific"
  approving_authority_id: "varies; ESL/SECO review path for certain energy amendments; TDLR state path for IHB; local adoption for municipal amendments"
  filing_required: "unresolved"
  registry_exists: "no unified statewide registry verified"
  registry_source_ids: []
  legal_basis_source_ids:
    - "src:usa-tx:lg-ch214"
    - "src:usa-tx:seco-local-ordinances"
    - "src:usa-tx:seco-adoption-process"
    - "src:usa-tx:tdlr-ihb-faq"
  verification_status: "partially_verified"
  confidence: 0.50
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement is broader than local amendment authority. Texas local governments commonly enforce construction codes, but some state programs set mandatory statewide standards or state-managed certification requirements. For example, local jurisdictions enforce energy codes, while SECO sets statewide minimums and program-specific stringency rules. Local electrical authorities perform permits/inspections where required, while TDLR sets the state electrical code. IHB is a state-administered code stream with local amendment limitations. Local fire AHJs and SFMO authority can overlap, but SFMO guidance explicitly recognizes local agencies with certified inspectors.

### 6.4 Known Local Amendment Registries

| Registry ID | Registry Name | Scope | Source IDs | Status | Notes |
| --- | --- | --- | --- | --- | --- |
| registry:usa-tx:energy-esl-determinations | ESL local energy-code amendment determinations | Energy-code amendments in nonattainment and affected counties | src:usa-tx:seco-local-ordinances; src:usa-tx:seco-adoption-process | identified_partial | Registry mechanics and coverage need extraction into structured records. |
| registry:usa-tx:municipal-building-amendments | Unified municipal building-code amendment registry | Building/residential/commercial/electrical | none | unresolved | No unified statewide registry was identified during this pass. |
| registry:usa-tx:local-fire-code-amendments | Unified local fire-code amendment registry | Fire operational/prevention codes | none | unresolved | Local AHJ-by-AHJ research required. |

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendments were parsed. This report captures only statewide and program-level authority signals.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: scoped_not_built

Jurisdiction stack:

```text
Address
  -> State of Texas
  -> County
  -> Municipality or unincorporated county
  -> Designated catastrophe area / windstorm eligibility area, if applicable
  -> Nonattainment or affected county energy-code classification, if applicable
  -> Local building AHJ
  -> Local fire AHJ or SFMO path
  -> Trade-specific AHJs such as electrical, plumbing, elevator, accessibility, or IHB program authority
  -> Applicable state program code adoption records
  -> Applicable local code and amendment records
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | U.S. Census TIGER/Line or Texas GIS portal candidate | none | statewide | unknown | candidate_not_selected |
| County | U.S. Census TIGER/Line or Texas GIS portal candidate | none | statewide | unknown | candidate_not_selected |
| Municipality | Texas municipal boundary source candidate | none | statewide | unknown | unresolved |
| Designated catastrophe area | TDI windstorm/designated catastrophe area source | src:usa-tx:tdi-windstorm-adopted-codes | coastal / designated counties and areas | agency-updated | partial |
| Energy nonattainment / affected county | SECO / ESL energy-code amendment process source | src:usa-tx:seco-local-ordinances; src:usa-tx:seco-adoption-process | affected energy jurisdictions | agency-updated | partial |
| Fire District / local fire AHJ | local jurisdiction datasets | none | statewide | unknown | unresolved |
| Special District | local/state datasets | none | statewide | unknown | unresolved |

### 7.3 AHJ Contact Data

No AHJ contact data was populated. Address-level AHJ resolution remains out of scope for this source pass.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Publisher | Source Type | URL / Locator | Accessed | Coverage | Reliability | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-tx:lg-ch214 | Texas Local Government Code Chapter 214, Municipal Regulation of Housing and Other Structures | Texas Legislature / Texas Constitution and Statutes | statute | https://statutes.capitol.texas.gov/Docs/LG/htm/LG.214.htm | 2026-06-26 | Municipal residential, commercial, electrical, and related code authority | official | Official statute HTML extraction was incomplete in this pass; section-level details were cross-checked with official legislative material and supplemental courtesy copies. |
| src:usa-tx:capitol-hb0738-2021 | HB 738 official legislative analysis / enrolled bill materials, 87th Legislature | Texas Legislature Online | legislative material | https://capitol.texas.gov/ | 2026-06-26 | Changed municipal IRC/IBC edition references to codes as existed on 2012-05-01 | official | Used to support municipal edition-reference update where current-code HTML was not fully extractable. |
| src:usa-tx:lg-ch233 | Texas Local Government Code Chapter 233, County Regulation of Housing and Other Structures | Texas Legislature / Texas Constitution and Statutes | statute | https://statutes.capitol.texas.gov/Docs/LG/htm/LG.233.htm | 2026-06-26 | County authority including fire-code and unincorporated construction-code streams | official | Official extraction incomplete; 2025 county construction-code provisions require final codified verification. |
| src:usa-tx:capitol-sb2-2025-amendment | Official legislative publication for 2025 county construction-code amendment text | Texas Legislature Online | legislative material | https://capitol.texas.gov/ | 2026-06-26 | Indicates county residential/commercial code editions as existed on 2018-01-01 and 2025-09-01 effective path | official | Use as a watch source only until final codified text is extracted. |
| src:usa-tx:seco-commercial-energy | Commercial and Multi-family Construction | Texas Comptroller / State Energy Conservation Office | agency guidance | https://comptroller.texas.gov/programs/seco/code/commercial.php | 2026-06-26 | 2015 IECC commercial/multifamily status and 2016-11-01 code history | official | Official agency page. |
| src:usa-tx:seco-local-ordinances | Local Ordinances | Texas Comptroller / State Energy Conservation Office | agency guidance | https://comptroller.texas.gov/programs/seco/code/ordinances.php | 2026-06-26 | Statewide energy minimums, local amendments, home-rule/local enforcement language | official | Official agency page. |
| src:usa-tx:seco-state-funded | State-Funded Buildings | Texas Comptroller / State Energy Conservation Office | agency guidance | https://comptroller.texas.gov/programs/seco/code/state-funded.php | 2026-06-26 | State-funded building energy standards effective 2021-01-01 | official | Official agency page. |
| src:usa-tx:seco-adoption-process | Energy Code Adoption | Texas Comptroller / State Energy Conservation Office | agency guidance | https://comptroller.texas.gov/programs/seco/code/ | 2026-06-26 | SECO adoption process, ESL review, comment period, minimum nine-month effective-date lead time, local amendment process | official | Official agency page. |
| src:usa-tx:tdlr-ihb-codes | Industrialized Housing and Buildings Mandatory Building Codes | Texas Department of Licensing and Regulation | agency guidance | https://www.tdlr.texas.gov/ihb/codes.htm | 2026-06-26 | IHB mandatory code list and transition rules effective 2024-07-01 | official | Agency page summarizes 16 TAC 70.100 and 70.101. |
| src:usa-tx:tdlr-ihb-adoption-2024 | New Mandatory Building Codes Effective July 1, 2024 | Texas Department of Licensing and Regulation | agency notice | https://www.tdlr.texas.gov/ihb/ihb.htm | 2026-06-26 | Council approval, Commission rule adoption, effective date, transition details | official | Agency news/notice. |
| src:usa-tx:tdlr-ihb-faq | Industrialized Housing and Buildings FAQs | Texas Department of Licensing and Regulation | agency FAQ | https://www.tdlr.texas.gov/ihb/ihbfaq.htm | 2026-06-26 | Local amendment limitations for IHB and municipal acceptance guidance | official | FAQ; should be paired with Occupations Code Chapter 1202 for production proof. |
| src:usa-tx:tdlr-electrical-compliance | Electricians Compliance Guide | Texas Department of Licensing and Regulation | agency guidance | https://www.tdlr.texas.gov/electricians/compliance-guide.htm | 2026-06-26 | 2023 NEC effective date and permit/start-of-work rule | official | Official agency page. |
| src:usa-tx:tdlr-electrical-laws-rules | Electricians Laws and Rules | Texas Department of Licensing and Regulation | agency guidance / legal index | https://www.tdlr.texas.gov/electricians/laws-rules.htm | 2026-06-26 | Occupations Code Chapter 1305 and 16 TAC Chapter 73 legal basis | official | Legal index page. |
| src:usa-tx:tdi-windstorm-adopted-codes | Adopted Building Codes: Windstorm Inspections | Texas Department of Insurance | agency guidance | https://www.tdi.texas.gov/wind/adoptedbuildingcodes.html | 2026-06-26 | 2024 IRC/IBC WPI-1 windstorm certification effective 2026-04-01 | official | Official agency page, last updated 2026-04-23. |
| src:usa-tx:tdi-sfmo-standards | Codes and Standards Used for Inspections by the State Fire Marshal | Texas Department of Insurance / State Fire Marshal's Office | agency guidance | https://www.tdi.texas.gov/fire/fmcodes.html | 2026-06-26 | 2021 NFPA 101 and NFPA 1 inspection standards effective 2023-09-01 | official | Official agency page. |
| src:usa-tx:tdi-sfmo-rules | Fire Safety Inspection Rules and Statutes | Texas Department of Insurance / State Fire Marshal's Office | agency guidance / legal index | https://www.tdi.texas.gov/fire/fmrules.html | 2026-06-26 | Government Code Chapter 417 and 28 TAC Chapter 34 inspection standards legal basis | official | Legal index page. |
| src:usa-tx:tdi-sfmo-inspections | Fire Safety Inspections Services | Texas Department of Insurance / State Fire Marshal's Office | agency guidance | https://www.tdi.texas.gov/fire/fmfsis.html | 2026-06-26 | SFMO inspection request scope and local certified-inspector handoff | official | Official agency page. |
| src:usa-tx:tdlr-tas | Texas Accessibility Standards | Texas Department of Licensing and Regulation | agency guidance / standards page | https://www.tdlr.texas.gov/ab/abtas.htm | 2026-06-26 | 2012 TAS and effective date | official | Official agency page. |
| src:usa-tx:tdlr-ab-laws-rules | Architectural Barriers Laws and Rules | Texas Department of Licensing and Regulation | agency guidance / legal index | https://www.tdlr.texas.gov/ab/laws-rules.htm | 2026-06-26 | Government Code Chapter 469 and 16 TAC Chapter 68 legal basis | official | Legal index page. |
| src:usa-tx:tdlr-ab-faq | Architectural Barriers FAQs | Texas Department of Licensing and Regulation | agency FAQ | https://www.tdlr.texas.gov/ab/abfaq.htm | 2026-06-26 | TAS compliance and submission thresholds | official | FAQ; use for applicability notes, not as sole legal proof. |
| src:usa-tx:tdlr-elevator-laws-rules | Elevator / Escalator Safety Laws and Rules | Texas Department of Licensing and Regulation | agency guidance / legal index | https://www.tdlr.texas.gov/elevator/laws-rules.htm | 2026-06-26 | Health and Safety Code Chapter 754 and 16 TAC Chapter 74 legal basis | official | Exact ASME edition not parsed. |
| src:usa-tx:tdlr-elevator-faq | Elevator / Escalator Safety FAQs | Texas Department of Licensing and Regulation | agency FAQ | https://www.tdlr.texas.gov/elevator/elfaq.htm | 2026-06-26 | Registration/licensing and program-scope support | official | FAQ; use only as supplemental program-scope support. |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-tx:lg-ch214 | official_html_extraction_gap | The official chapter page was identified, but section-level HTML extraction was incomplete in the tool output. | Re-open official current-code sections before final verification; current report may remain partially_verified. |
| src:usa-tx:lg-ch233 | official_html_extraction_gap | The official chapter page was identified, but some county construction-code details came from legislative-publication snippets and courtesy copies. | Treat county residential/commercial provisions as watch/partial until codified current text is parsed. |
| src:usa-tx:capitol-sb2-2025-amendment | legislative_material_not_codified_text | Official legislative publication appears to reflect amendments effective 2025-09-01, but not a parsed current statute page. | Use for pending/watch row and open issue, not final production field. |
| src:usa-tx:tdlr-ihb-faq | faq_not_primary_rule | FAQ is official agency guidance but not the rule text itself. | Pair with Occupations Code Chapter 1202 and 16 TAC Chapter 70 for verified legal proof. |
| src:usa-tx:tdlr-ab-faq | faq_not_primary_rule | FAQ is official agency guidance for thresholds and process. | Pair with Government Code Chapter 469 and 16 TAC Chapter 68 for verified legal proof. |
| src:usa-tx:tdlr-elevator-faq | faq_not_primary_rule | FAQ is official agency guidance for program scope. | Pair with Health and Safety Code Chapter 754 and 16 TAC Chapter 74 for verified legal proof. |

### 8.3 Supplemental Sources

| Source ID | Title | Publisher | Source Type | URL / Locator | Accessed | Coverage | Caveat |
| --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-tx:supp-findlaw-lg214-212 | Texas Local Government Code §214.212 courtesy text | FindLaw | supplemental legal publisher | https://codes.findlaw.com/tx/local-government-code/loc-gov-t-sect-214-212/ | 2026-06-26 | Municipal IRC applicability and local amendments | Not official; used only to fill extraction gaps after official chapter/source identification. |
| src:usa-tx:supp-findlaw-lg214-214 | Texas Local Government Code §214.214 courtesy text | FindLaw | supplemental legal publisher | https://codes.findlaw.com/tx/local-government-code/loc-gov-t-sect-214-214/ | 2026-06-26 | Municipal NEC applicability and local amendments | Not official; used only to fill extraction gaps. |
| src:usa-tx:supp-justia-lg214-216 | Texas Local Government Code §214.216 courtesy text | Justia | supplemental legal publisher | https://law.justia.com/codes/texas/local-government-code/title-7/subtitle-a/chapter-214/subchapter-g/section-214-216/ | 2026-06-26 | Municipal IBC applicability and local amendments | Not official; used only to fill extraction gaps. |
| src:usa-tx:supp-findlaw-lg233-062 | Texas Local Government Code §233.062 courtesy text | FindLaw | supplemental legal publisher | https://codes.findlaw.com/tx/local-government-code/loc-gov-t-sect-233-062/ | 2026-06-26 | County fire-code scope and IFC/UFC 2005 reference | Not official; used only to fill extraction gaps after official chapter/source identification. |

### 8.4 Source Extraction Metadata

| Extraction ID | Source IDs | Extracted Fields | Extracted On | Method | Confidence | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| extract:usa-tx:001 | src:usa-tx:seco-local-ordinances; src:usa-tx:seco-commercial-energy; src:usa-tx:seco-state-funded; src:usa-tx:seco-adoption-process | energy code editions, local enforcement, local amendments, state-funded standards, adoption timing rules | 2026-06-26 | official agency HTML | 0.85 | Strong official support for core energy fields. |
| extract:usa-tx:002 | src:usa-tx:tdlr-ihb-codes; src:usa-tx:tdlr-ihb-adoption-2024 | IHB mandatory code editions, adoption/effective dates, transition rules | 2026-06-26 | official agency HTML | 0.90 | Strong official support for IHB code matrix. |
| extract:usa-tx:003 | src:usa-tx:tdlr-electrical-compliance; src:usa-tx:tdlr-electrical-laws-rules | state electrical code edition and date rules | 2026-06-26 | official agency HTML | 0.90 | Strong official support for electrical fields. |
| extract:usa-tx:004 | src:usa-tx:tdi-sfmo-standards; src:usa-tx:tdi-sfmo-rules; src:usa-tx:tdi-sfmo-inspections | SFMO inspection standards, local-inspector handoff, legal index | 2026-06-26 | official agency HTML | 0.78 | Local fire-code authority still requires additional jurisdictional mapping. |
| extract:usa-tx:005 | src:usa-tx:tdi-windstorm-adopted-codes | 2024 IRC/IBC windstorm certification path and 2026-04-01 effective date | 2026-06-26 | official agency HTML | 0.88 | Windstorm is a specialized certification program. |
| extract:usa-tx:006 | src:usa-tx:lg-ch214; src:usa-tx:capitol-hb0738-2021; supplemental LG 214 sources | municipal IRC/IBC/NEC authority, local amendment support | 2026-06-26 | official statute/source identification plus supplemental courtesy text | 0.55 | Needs final statute-section extraction for verified status. |
| extract:usa-tx:007 | src:usa-tx:lg-ch233; src:usa-tx:capitol-sb2-2025-amendment; src:usa-tx:supp-findlaw-lg233-062 | county fire-code and emerging construction-code authority | 2026-06-26 | official statute/source identification plus legislative material and supplemental courtesy text | 0.45 | County construction-code fields remain open. |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 1.00 | none | Status matches current source-backed but incomplete report. |
| report | risk.overall_confidence | 0.56 | verified | 1.00 | none | Confidence reflects strong specialty-program support and unresolved local/county gaps. |
| ahj:usa-tx:hybrid-state-local | authority model | hybrid_state_local_model | partially_verified | 0.60 | src:usa-tx:lg-ch214; src:usa-tx:lg-ch233; src:usa-tx:seco-local-ordinances; src:usa-tx:tdlr-ihb-codes; src:usa-tx:tdlr-electrical-laws-rules; src:usa-tx:tdi-sfmo-rules | Model supported by multiple distinct authority streams. |
| adoption:usa-tx:seco-statewide-energy-2015 | code editions | 2015 IRC Chapter 11 and 2015 IECC | verified_core | 0.85 | src:usa-tx:seco-local-ordinances; src:usa-tx:seco-commercial-energy | Single-family and other-building split captured. |
| adoption:usa-tx:tdlr-ihb-2024 | code editions / effective date | 2021 I-Codes, 2015 IECC, 2020 NEC; effective 2024-07-01 | verified_core | 0.90 | src:usa-tx:tdlr-ihb-codes; src:usa-tx:tdlr-ihb-adoption-2024 | Transition captured. |
| adoption:usa-tx:tdlr-electrical-2023-nec | edition / date rule | 2023 NEC effective 2023-09-01; permit/start-of-work logic | verified_core | 0.90 | src:usa-tx:tdlr-electrical-compliance | Strong official agency guidance. |
| adoption:usa-tx:sfmo-nfpa-2021 | fire inspection standards | 2021 NFPA 101 and NFPA 1 effective 2023-09-01 | verified_core | 0.78 | src:usa-tx:tdi-sfmo-standards; src:usa-tx:tdi-sfmo-rules | SFMO inspection standard is verified; local fire-code mapping remains partial. |
| adoption:usa-tx:tdi-windstorm-2024 | windstorm code edition | 2024 IRC/IBC effective 2026-04-01 for WPI-1 | verified_core | 0.88 | src:usa-tx:tdi-windstorm-adopted-codes | Program-specific, not general statewide. |
| local-enforcement:usa-tx | model | local_primary_with_state_specialty_programs | partially_verified | 0.62 | src:usa-tx:seco-local-ordinances; src:usa-tx:tdlr-electrical-compliance; src:usa-tx:tdi-sfmo-inspections | Needs AHJ dataset follow-up. |
| local-amendment-rule:usa-tx | model | permitted_by_program_with_limits | partially_verified | 0.50 | src:usa-tx:lg-ch214; src:usa-tx:seco-local-ordinances; src:usa-tx:seco-adoption-process; src:usa-tx:tdlr-ihb-faq | No unified registry verified. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| Frontmatter state/report placeholders replaced | pass | Texas state identifiers, report_id, and last_updated are populated. |
| Conservative status used | pass | Status is partially_verified because official sources support core fields while open issues remain. |
| All source IDs resolve | pass | Body source IDs are present in Section 8 or intentionally marked none. |
| All authority IDs resolve | pass | Authority IDs are defined in authority tables or YAML-like sections. |
| All current code families have matrix rows | pass | Rows are present for template code families. |
| Building and operational fire code are separated | pass | Fire construction references and operational/prevention standards are distinct rows. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Date fields are separated; unresolved dates are null or explicitly scoped. |
| Effective dates are valid ISO dates | pass | Entered dates use YYYY-MM-DD. |
| No impossible date sequences | pass | No contradictory date sequences were identified. |
| Transition rules have explicit trigger conditions | pass | IHB, electrical, energy adoption, SFMO, and windstorm triggers are captured. |
| Permit-date logic is captured where applicable | pass | Electrical permit/no-permit logic is captured. |
| Local enforcement model classified | pass | Model is classified as local_primary_with_state_specialty_programs. |
| Local amendment rule classified | pass | Model is classified as permitted_by_program_with_limits. |
| AHJ confirmation metadata present | fail | AHJ contact and boundary data were not populated. |
| Full state amendment text parsed | fail | Program amendment sources are identified but not fully parsed. |
| County construction-code provisions final-codified | fail | 2025 county construction-code fields require final codified statute review. |
| Official-source caveats captured | pass | Caveats are recorded in Section 8.2. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-tx:001 | high | county residential/commercial construction codes | Official legislative material indicates 2025 county construction-code changes for certain unincorporated areas, but final codified text was not parsed. | Extract current Local Government Code Chapter 233 sections and confirm effective/applicability rules. | null | null | open |
| issue:usa-tx:002 | high | municipal code sections | Municipal IRC/IBC/NEC authority is supported, but official section-level text extraction was incomplete. | Re-extract LG §§214.212, 214.214, 214.216, and related sections from official current code. | null | null | open |
| issue:usa-tx:003 | medium | non-IHB mechanical/plumbing/fuel gas | Statewide adoption for site-built mechanical, plumbing, and fuel-gas codes outside IHB was not resolved. | Review state plumbing/mechanical statutes and local enabling authority. | null | null | open |
| issue:usa-tx:004 | medium | local fire-code authority and editions | SFMO standards are verified, but local/county fire-code adoption scope and editions remain incomplete. | Extract county and municipal fire-code authority and build jurisdictional fire-AHJ model. | null | null | open |
| issue:usa-tx:005 | medium | state amendments | IHB, energy, SFMO, electrical, and windstorm amendment publication paths are identified but amendment text is not fully parsed. | Parse 16 TAC 70.101, 34 TAC Chapter 19, 16 TAC Chapter 73, 28 TAC Chapter 34, and windstorm administrative rules. | null | null | open |
| issue:usa-tx:006 | medium | AHJ/boundary/contact resolution | No address-level boundary data or AHJ contacts were populated. | Select boundary datasets and AHJ registry/contact sources. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-tx:lg-214 | src:usa-tx:lg-ch214 | html_diff | monthly | municipal code-authority sections amended or recodified | 2026-06-26 | active |
| watch:usa-tx:lg-233 | src:usa-tx:lg-ch233 | html_diff | monthly | county construction/fire-code authority amended or codified text updated | 2026-06-26 | active |
| watch:usa-tx:seco-energy | src:usa-tx:seco-adoption-process | html_diff | monthly | SECO notice/comment, ESL review, or adoption of newer IECC/ASHRAE edition | 2026-06-26 | active |
| watch:usa-tx:tdlr-ihb | src:usa-tx:tdlr-ihb-codes | html_diff | monthly | mandatory IHB code list or transition rule changes | 2026-06-26 | active |
| watch:usa-tx:tdlr-electrical | src:usa-tx:tdlr-electrical-compliance | html_diff | quarterly | NEC edition or permit/start-of-work logic changes | 2026-06-26 | active |
| watch:usa-tx:sfmo | src:usa-tx:tdi-sfmo-standards | html_diff | quarterly | NFPA inspection standard edition or SFMO AHJ guidance changes | 2026-06-26 | active |
| watch:usa-tx:tdi-windstorm | src:usa-tx:tdi-windstorm-adopted-codes | html_diff | monthly through 2026-04-01 transition, then quarterly | windstorm code edition, WPI-1 rule, designated catastrophe area, or debris protection update | 2026-06-26 | active |
| watch:usa-tx:tdlr-accessibility | src:usa-tx:tdlr-tas | html_diff | quarterly | TAS or architectural-barriers threshold changes | 2026-06-26 | active |
| watch:usa-tx:tdlr-elevator | src:usa-tx:tdlr-elevator-laws-rules | html_diff | quarterly | ASME edition or conveyance rule changes | 2026-06-26 | pending |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-24 | Generated baseline Texas draft report | report:usa-tx | none | Codex | Original stub contained unresolved placeholder-style entries and no primary source registry. |
| 2026-06-26 | Populated partially verified Texas report | report:usa-tx; ahj:usa-tx:hybrid-state-local; adoption:usa-tx:seco-statewide-energy-2015; adoption:usa-tx:tdlr-ihb-2024; adoption:usa-tx:tdlr-electrical-2023-nec; adoption:usa-tx:sfmo-nfpa-2021; adoption:usa-tx:tdi-windstorm-2024 | src:usa-tx:lg-ch214; src:usa-tx:lg-ch233; src:usa-tx:seco-local-ordinances; src:usa-tx:tdlr-ihb-codes; src:usa-tx:tdlr-electrical-compliance; src:usa-tx:tdi-sfmo-standards; src:usa-tx:tdi-windstorm-adopted-codes | ChatGPT | Source registry added, authority model populated, code adoption records normalized, unresolved items retained explicitly. |
