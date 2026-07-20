---
state:
  state_id: "US-AK"
  name: "Alaska"
  abbreviation: "AK"
report:
  report_id: "state-report:usa-ak"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-26"
  last_verified: "2026-06-26"
  reviewed_by: null
risk:
  overall_confidence: 0.78 # 0.00 - 1.00
  risk_flags:
    - "ahfc_residential_code_scope_is_financing_tied_not_universal"
    - "official_aac_direct_html_for_title_8_not_extracted"
    - "local_deferred_jurisdiction_roster_not_extracted"
    - "municipal_approved_code_registry_not_populated"
    - "trade_code_effective_dates_incomplete"
  open_questions_count: 6
---

# State Building Code Authority Report: Alaska

## 1. Executive Summary

- **Authority model:** Alaska has a hybrid state-and-local model. The Alaska Department of Public Safety, Division of Fire and Life Safety / State Fire Marshal is the primary statewide fire and life-safety code authority for the 2021 IBC, IEBC, IMC, IFGC, and IFC as adopted in 13 AAC 50 and applied through 13 AAC 50 — 13 AAC 55. The Department of Labor and Workforce Development, Mechanical Inspection section separately administers statewide mechanical-inspection programs for plumbing, electrical, boilers/pressure vessels, elevators, and similar installations. Alaska Housing Finance Corporation administers residential and energy standards for AHFC or state-financed housing and other AHFC-linked programs.

- **Statewide code status:** State fire/life-safety adoptions are verified for the 2021 IBC, 2021 IEBC, 2021 IMC, limited 2021 IFGC, and 2021 IFC. State trade-code adoptions are verified at the current-program-page level for 2018 UPC, 2020 NEC, 2017 NESC, boiler/pressure-vessel codes, and elevator/conveyance codes. Residential IRC and energy IECC/BEES adoptions are verified as AHFC/state-financing-linked standards, not as a universal private residential building code for every project in the state.

- **Local enforcement model:** State Fire Marshal enforcement is state-led with deferral/exemption mechanisms for municipalities and state agencies that meet 13 AAC 50.075 criteria. DOLWD Mechanical Inspection states that plumbing, electrical, and elevator inspection authority is deferred to local government inspectors in the Municipality of Anchorage through a memorandum of agreement.

- **Local amendment posture:** For 13 AAC 50 — 13 AAC 55, local political subdivisions may not set less-stringent minimum standards without written State Fire Marshal approval, and proposed local revisions relating to fire and life safety must be submitted to the State Fire Marshal. Electrical and plumbing statutes preserve local authority for not-less-stringent standards, but the current local filing/approval workflow outside the fire/life-safety chapters remains unresolved.

- **Known transition periods or pending changes:** The 2021 State Fire Marshal code amendments show a 2022-10-28 amendment/effective marker in Register 244. AHFC residential and energy code materials use 2019-01-01 as the construction-start applicability date for 2018 IRC/IECC updates. DOLWD Mechanical Inspection announced a separate 2025 licensing/reciprocity regulation change, but that item is not a model-code edition update.

- **Production readiness:** partially_ready

### Key Findings

```yaml
---
key_findings:
- topic: Primary statewide building/fire-life-safety authority
  finding: Alaska DPS Division of Fire and Life Safety / State Fire Marshal adopts
    and enforces the Alaska Fire Code based on the IFC with Alaska-specific amendments
    and administers plan review and inspections.
  confidence: 0.94
  source_ids:
  - src:usa-ak:dps:fire-life-safety-page
  - src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf
- topic: Commercial/state fire-life-safety building code
  finding: 13 AAC 50.020 adopts the 2021 IBC, with chapters and amendments, to regulate
    all occupancies and buildings within its stated scope.
  confidence: 0.92
  source_ids:
  - src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf
- topic: Existing building code
  finding: 13 AAC 50.021 adopts the 2021 IEBC, with amendments, to regulate all occupancies
    and buildings within its stated scope.
  confidence: 0.9
  source_ids:
  - src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf
- topic: Mechanical code
  finding: 13 AAC 50.023 adopts the 2021 IMC, with amendments.
  confidence: 0.9
  source_ids:
  - src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf
- topic: Fuel gas code
  finding: '13 AAC 50.024 adopts only specified 2021 IFGC provisions: Chapter 3 Section
    304 and Chapters 6 and 7.'
  confidence: 0.88
  source_ids:
  - src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf
- topic: Fire code
  finding: 13 AAC 50.025 adopts the 2021 IFC, with specified chapters reserved and
    Alaska amendments.
  confidence: 0.94
  source_ids:
  - src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf
- topic: Electrical code
  finding: DOLWD Mechanical Inspection lists 8 AAC 70.025(a) as the 2020 National
    Electrical Code and 8 AAC 70.025(b) as the 2017 National Electrical Safety Code.
  confidence: 0.86
  source_ids:
  - src:usa-ak:dolwd:mechanical-inspection-page
- topic: Plumbing code
  finding: DOLWD Mechanical Inspection lists 2018 UPC, 2018 Uniform Swimming Pool/Spa/Hot
    Tub Code, and 2018 Uniform Solar Energy/Hydronics/Geothermal Code as adopted codes;
    the April 2020 plumbing publication shows 8 AAC 63.010 adopting those 2018 IAPMO
    codes.
  confidence: 0.85
  source_ids:
  - src:usa-ak:dolwd:mechanical-inspection-page
  - src:usa-ak:dolwd:plumbing-code-pdf
- topic: AHFC residential code
  finding: AHFC materials identify the 2018 IRC with Alaska-specific amendments as
    the applicable AHFC residential building code/minimum construction standard.
  confidence: 0.82
  source_ids:
  - src:usa-ak:ahfc:codes-standards
  - src:usa-ak:ahfc:minimum-construction-standards
  - src:usa-ak:ahfc:2018-code-adoption-packet
- topic: AHFC energy standard
  finding: AHFC BEES is comprised of the 2018 IECC, ASHRAE 62.2-2016, and Alaska-specific
    amendments, with a minimum 5 Star energy rating.
  confidence: 0.88
  source_ids:
  - src:usa-ak:ahfc:bees
- topic: Local fire/life-safety amendments
  finding: Local fire/life-safety revisions must be submitted to the State Fire Marshal,
    and less-stringent local minimum standards require written State Fire Marshal
    approval.
  confidence: 0.92
  source_ids:
  - src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf
- topic: Local general authority
  finding: Home-rule boroughs and cities may exercise all legislative powers not prohibited
    by law or charter, but code-specific state minimums and approval rules remain
    controlling where applicable.
  confidence: 0.86
  source_ids:
  - src:usa-ak:constitution:article-x
  - src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf
```

---

### 2.1 Primary Building Code Authorities

```yaml
---
primary_building_code_authorities:
  authority_id: "auth:us-ak:dps:state-fire-marshal"
  authority_name: "Alaska Department of Public Safety, Division of Fire and Life Safety / State Fire Marshal"
  authority_type: "state_agency_fire_life_safety_code_authority"
  legal_basis:
    - "AS 18.70.010"
    - "AS 18.70.080"
    - "13 AAC 50 — 13 AAC 55"
  role: "Adopts and administers statewide fire and life-safety construction and operational code requirements, including 2021 IBC, IEBC, IMC, limited IFGC, and IFC adoptions with Alaska amendments."
  enforcement_model: "state_led_with_deferred_local_authority"
  source_ids:
    - "src:us-ak:dps:fire-life-safety-page"
    - "src:us-ak:dps:aac-13-50-55-fire-prevention-pdf"
  verification_status: "verified"
```

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-ak:dps:state-fire-marshal | State Fire Marshal / Division of Fire and Life Safety | Adopts and applies 2021 IBC within 13 AAC 50 scope; Chapter 1 scope exceptions include detached one-, two-, and three-family dwellings and specified townhouses. | AS 18.70.080; 13 AAC 50.020 | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | verified |
| Residential | ahj:usa-ak:ahfc:residential-standards | Alaska Housing Finance Corporation | Administers AHFC Minimum Construction Standards / residential code for affected financed housing and residential units outside an approved municipal building code. | 15 AAC 150.030; 15 AAC 150.035; AS 18.56.300 | src:usa-ak:ahfc:codes-standards, src:usa-ak:ahfc:minimum-construction-standards, src:usa-ak:ahfc:2018-code-adoption-packet | partially_verified_scope_limited |
| Existing Building / Rehabilitation | ahj:usa-ak:dps:state-fire-marshal | State Fire Marshal / Division of Fire and Life Safety | Adopts and applies 2021 IEBC within 13 AAC 50 scope. | AS 18.70.080; 13 AAC 50.021 | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | verified |
| Mechanical | ahj:usa-ak:dps:state-fire-marshal | State Fire Marshal / Division of Fire and Life Safety | Adopts and applies 2021 IMC within 13 AAC 50 scope. | AS 18.70.080; 13 AAC 50.023 | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | verified |
| Plumbing | ahj:usa-ak:dolwd:mechanical-inspection | DOLWD Labor Standards and Safety, Mechanical Inspection | Administers plumbing installations; current official program page lists 2018 UPC and related IAPMO codes. | AS 18.60.705; AS 18.60.710; 8 AAC 63.010 / DOLWD program page | src:usa-ak:dolwd:mechanical-inspection-page, src:usa-ak:dolwd:plumbing-code-pdf | partially_verified_official_program_current |
| Fuel Gas | ahj:usa-ak:dps:state-fire-marshal | State Fire Marshal / Division of Fire and Life Safety | Adopts limited 2021 IFGC provisions and cross-references DOLWD plumbing/electrical codes. | AS 18.70.080; 13 AAC 50.024 | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | verified |
| Electrical | ahj:usa-ak:dolwd:mechanical-inspection | DOLWD Labor Standards and Safety, Mechanical Inspection | Administers electrical installations; current official program page lists 2020 NEC and 2017 NESC. | AS 18.60.580; AS 18.60.600; 8 AAC 70.025 / DOLWD program page | src:usa-ak:dolwd:mechanical-inspection-page, src:usa-ak:dolwd:electrical-safety-pdf | partially_verified_official_program_current |
| Energy | ahj:usa-ak:ahfc:bees | Alaska Housing Finance Corporation | Administers BEES for AHFC/state-financing-linked buildings. | 15 AAC 155.010; AS 46.11.040; AHFC program materials | src:usa-ak:ahfc:bees, src:usa-ak:ahfc:codes-standards, src:usa-ak:ahfc:2018-code-adoption-packet | partially_verified_scope_limited |
| Fire - construction references | ahj:usa-ak:dps:state-fire-marshal | State Fire Marshal / Division of Fire and Life Safety | Applies IBC/IEBC/IMC/IFGC/IFC requirements for fire and life safety plan review and construction-related safeguards. | AS 18.70.080; 13 AAC 50.020, .021, .023, .024, .025, .027 | src:usa-ak:dps:fire-life-safety-page, src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | verified |
| Fire - operational / prevention code | ahj:usa-ak:dps:state-fire-marshal | State Fire Marshal / Division of Fire and Life Safety | Adopts and enforces the Alaska Fire Code based on 2021 IFC with Alaska amendments. | AS 18.70.080; 13 AAC 50.025; 13 AAC 55.010 | src:usa-ak:dps:fire-life-safety-page, src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | verified |
| Accessibility | ahj:usa-ak:dps:state-fire-marshal | State Fire Marshal / Division of Fire and Life Safety | Accessibility provisions are likely incorporated through the IBC where the IBC applies; Alaska-specific accessibility enforcement beyond the IBC was not separately extracted. | 13 AAC 50.020 | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | partially_verified_indirect |
| Elevator / Conveyance | ahj:usa-ak:dolwd:mechanical-inspection | DOLWD Labor Standards and Safety, Mechanical Inspection | Administers elevator, escalator, platform lift, chairlift, tramway, amusement ride, ski lift, and related device inspection programs. | 8 AAC 77.005 / DOLWD program page | src:usa-ak:dolwd:mechanical-inspection-page, src:usa-ak:dolwd:elevator-safety-pdf | partially_verified_official_program_current |
| Boiler / Pressure Vessel | ahj:usa-ak:dolwd:mechanical-inspection | DOLWD Labor Standards and Safety, Mechanical Inspection | Administers boilers and unfired pressure vessels and lists current adopted boiler/pressure-vessel standards. | 8 AAC 80.010 / DOLWD program page | src:usa-ak:dolwd:mechanical-inspection-page, src:usa-ak:dolwd:boiler-pressure-vessel-pdf | partially_verified_official_program_current |

### 2.3 Authority Hierarchy Notes

Alaska is not a simple single-board statewide building-code state. For commercial and public occupancies within the State Fire Marshal's scope, the 13 AAC 50 code set is the central statewide fire/life-safety code layer. The IBC adoption itself deletes or modifies parts of Chapter 1 and removes references to the IECC, IRC, IPMC, IPSDC, and IWUIC, so the IBC row should not be treated as a universal adoption of every ICC-family code for every project. Plumbing and electrical references in the DPS code set are replaced with DOLWD-adopted plumbing and electrical codes.

Residential one-, two-, three-, and certain townhouse projects require separate treatment. AHFC administers the 2018 IRC with Alaska-specific amendments and BEES for AHFC or state-financing-linked housing, and AHFC's page frames those standards as a financing eligibility requirement. Local municipal codes may also govern residential projects, particularly in approved municipalities. Home-rule authority remains relevant in the background, but code-specific state minimums, approved municipal code rules, and State Fire Marshal deferral/approval rules are the controlling source trail for this report.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-ak:001 | ahj:usa-ak:dps:state-fire-marshal | adopts_by_reference | 2021 IBC / IEBC / IMC / limited IFGC / IFC with Alaska amendments | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | verified |
| edge:usa-ak:002 | ahj:usa-ak:dps:state-fire-marshal | substitutes | DOLWD electrical code under 8 AAC 70.025 for deleted NFPA 70 references in IBC/IFC/IFGC contexts | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf, src:usa-ak:dolwd:mechanical-inspection-page | verified |
| edge:usa-ak:003 | ahj:usa-ak:dps:state-fire-marshal | substitutes | DOLWD plumbing code under 8 AAC 63.010 for deleted IPC/IFGC references where specified | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf, src:usa-ak:dolwd:mechanical-inspection-page | verified |
| edge:usa-ak:004 | ahj:usa-ak:dps:state-fire-marshal | may_defer_to | municipalities or state agencies meeting 13 AAC 50.075 criteria | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | verified |
| edge:usa-ak:005 | ahj:usa-ak:dolwd:mechanical-inspection | administers | plumbing, electrical, boiler/pressure-vessel, elevator/conveyance programs | src:usa-ak:dolwd:mechanical-inspection-page | verified |
| edge:usa-ak:006 | ahj:usa-ak:ahfc:residential-standards | conditions_financing_on | AHFC Minimum Construction Standards / 2018 IRC with Alaska amendments | src:usa-ak:ahfc:codes-standards, src:usa-ak:ahfc:minimum-construction-standards | partially_verified |
| edge:usa-ak:007 | ahj:usa-ak:ahfc:bees | conditions_financing_or_state_assistance_on | BEES / 2018 IECC / ASHRAE 62.2-2016 / Alaska amendments | src:usa-ak:ahfc:bees | verified |
| edge:usa-ak:008 | ahj:usa-ak:constitution:local-government | confers_general_power_to | home-rule boroughs and cities, subject to prohibitions by law or charter | src:usa-ak:constitution:article-x | verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | 13 AAC 50.020 International Building Code | International Building Code | 2021 | adopted_state_fire_life_safety_scope | null | 2022-10-28 | 2022-10-28 | null | Amendment history shows 2022-10-28 / Register 244; no separate grace period extracted. IBC scope excludes detached one-, two-, and three-family dwellings and specified townhouses. | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf |
| Residential | AHFC Minimum Construction Standards / residential building code | International Residential Code with Alaska-specific amendments | 2018 | adopted_scope_limited_ahfc_financing_and_approved_municipal_code_context | 2018-11-28 | 2018-11-28 | 2019-01-01 | 2019-01-01 | Applies to buildings whose construction began on or after 2019-01-01, with construction-began date determined by foundation installation date in AHFC regulatory materials; current AHFC page presents this as an AHFC/statutory financing eligibility standard. | src:usa-ak:ahfc:codes-standards, src:usa-ak:ahfc:minimum-construction-standards, src:usa-ak:ahfc:2018-code-adoption-packet |
| Existing Building / Rehabilitation | 13 AAC 50.021 International Existing Building Code | International Existing Building Code | 2021 | adopted_state_fire_life_safety_scope | null | 2022-10-28 | 2022-10-28 | null | Amendment history shows 2022-10-28 / Register 244; no separate grace period extracted. | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf |
| Mechanical | 13 AAC 50.023 International Mechanical Code | International Mechanical Code | 2021 | adopted_state_fire_life_safety_scope | null | 2022-10-28 | 2022-10-28 | null | Amendment history shows 2022-10-28 / Register 244; no separate grace period extracted. | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf |
| Plumbing | Alaska minimum plumbing standards / DOLWD Mechanical Inspection adopted code | Uniform Plumbing Code; Uniform Swimming Pool, Spa and Hot Tub Code; Uniform Solar Energy, Hydronics and Geothermal Code | 2018 | adopted_state_minimum_plumbing | null | 2020-04-24 | 2020-04-24 | null | DOLWD current page confirms 2018 IAPMO code editions; April 2020 plumbing publication provides regulatory text and history, but it is marked informational and should be refreshed from official AAC in production. | src:usa-ak:dolwd:mechanical-inspection-page, src:usa-ak:dolwd:plumbing-code-pdf |
| Fuel Gas | 13 AAC 50.024 International Fuel Gas Code | International Fuel Gas Code, limited provisions | 2021 | adopted_limited_state_fire_life_safety_scope | null | 2022-10-28 | 2022-10-28 | null | Only Chapter 3 Section 304 and Chapters 6 and 7 are adopted; IECC and IRC references are deleted from Chapter 1. | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf |
| Electrical | Alaska minimum electrical standards / DOLWD Mechanical Inspection adopted code | National Electrical Code and National Electrical Safety Code | 2020 NEC; 2017 NESC | adopted_state_minimum_electrical | null | null | null | null | Current DOLWD program page lists 2020 NEC and 2017 NESC; official current AAC text/effective date for the 2020 NEC update was not directly extracted. | src:usa-ak:dolwd:mechanical-inspection-page |
| Energy | AHFC Building Energy Efficiency Standard (BEES) | International Energy Conservation Code; ASHRAE 62.2; Alaska-specific amendments | 2018 IECC; ASHRAE 62.2-2016 | adopted_scope_limited_ahfc_or_state_financial_assistance | 2018-11-28 | 2018-11-28 | 2019-01-01 | 2019-01-01 | AHFC page states BEES applies to buildings that began construction on or after 1992 if AHFC or other state financial assistance is used; 2018 update materials identify 2019-01-01 applicability for the 2018 IECC update. | src:usa-ak:ahfc:bees, src:usa-ak:ahfc:codes-standards, src:usa-ak:ahfc:2018-code-adoption-packet |
| Fire - construction references | Alaska Fire Code / 13 AAC 50 construction-related provisions | IBC, IEBC, IMC, limited IFGC, IFC | 2021 | adopted_state_fire_life_safety_scope | null | 2022-10-28 | 2022-10-28 | null | Construction plan review under 13 AAC 50.027 focuses on fire and life-safety review and does not cover all structural/mechanical/electrical review beyond life-safety confirmation. | src:usa-ak:dps:fire-life-safety-page, src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf |
| Fire - operational / prevention code | 13 AAC 50.025 International Fire Code | International Fire Code | 2021 | adopted_statewide_fire_prevention_scope | null | 2022-10-28 | 2022-10-28 | null | Chapters 13-19, 41-49, and 68-79 are reserved; Chapters 1-12, 20-40, 50-67, 80 and Appendices B-I, K, L, and N are adopted with amendments. | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf |
| Accessibility | IBC accessibility provisions within 13 AAC 50.020 scope | International Building Code | 2021 | indirectly_adopted_with_ibc_scope | null | 2022-10-28 | 2022-10-28 | null | IBC adopted, but accessibility-specific Alaska amendments and relationship to federal ADA/FHAct were not separately parsed. | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf |
| Elevator / Conveyance | DOLWD elevator and conveyance standards | ASME A17.1, A17.3, A18.1, A17.2, A17.6, A17.7 and related standards | 2016 / 2011 / 2017 / 2007 as listed | adopted_state_program_current | null | null | null | null | DOLWD current page lists adopted elevator/conveyance standards; official current AAC text/effective dates not fully extracted. | src:usa-ak:dolwd:mechanical-inspection-page, src:usa-ak:dolwd:elevator-safety-pdf |
| Boiler / Pressure Vessel | DOLWD boiler and pressure-vessel standards | ASME Boiler and Pressure Vessel Code; API Pressure Vessel Inspection Code; National Board Inspection Code | 2013 ASME BPVC; 2018 API; current NBIC | adopted_state_program_current | null | null | null | null | DOLWD current page lists adopted boiler/pressure-vessel standards; official current AAC text/effective dates not fully extracted. | src:usa-ak:dolwd:mechanical-inspection-page, src:usa-ak:dolwd:boiler-pressure-vessel-pdf |

### 3.2 Adoption Records

| Record ID | Code Family | Authority ID | Base Code | Edition | State Amendments? | Effective Date | Scope Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| adoption:usa-ak:dps:ibc-2021 | building | ahj:usa-ak:dps:state-fire-marshal | IBC | 2021 | yes | 2022-10-28 | Regulates all occupancies and buildings in 13 AAC 50 scope, subject to Alaska Chapter 1 deletions and residential exceptions. | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | verified |
| adoption:usa-ak:dps:iebc-2021 | existing_building | ahj:usa-ak:dps:state-fire-marshal | IEBC | 2021 | yes | 2022-10-28 | Regulates existing occupancies and buildings in 13 AAC 50 scope. | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | verified |
| adoption:usa-ak:dps:imc-2021 | mechanical | ahj:usa-ak:dps:state-fire-marshal | IMC | 2021 | yes | 2022-10-28 | Regulates mechanical systems in 13 AAC 50 scope; code official is State Fire Marshal or designated/deferred representative. | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | verified |
| adoption:usa-ak:dps:ifgc-2021-limited | fuel_gas | ahj:usa-ak:dps:state-fire-marshal | IFGC | 2021 | yes | 2022-10-28 | Adopts Chapter 3 Section 304 and Chapters 6 and 7 only. | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | verified |
| adoption:usa-ak:dps:ifc-2021 | fire_operational | ahj:usa-ak:dps:state-fire-marshal | IFC | 2021 | yes | 2022-10-28 | Adopts IFC Chapters 1-12, 20-40, 50-67, 80 and Appendices B-I, K, L, N; reserves Chapters 13-19, 41-49, 68-79. | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | verified |
| adoption:usa-ak:dolwd:upc-2018 | plumbing | ahj:usa-ak:dolwd:mechanical-inspection | UPC and related IAPMO codes | 2018 | yes | 2020-04-24 | Minimum plumbing standards; DOLWD jurisdiction noted as communities with populations of 2,500 and above. | src:usa-ak:dolwd:mechanical-inspection-page, src:usa-ak:dolwd:plumbing-code-pdf | partially_verified |
| adoption:usa-ak:dolwd:nec-2020 | electrical | ahj:usa-ak:dolwd:mechanical-inspection | NEC | 2020 | unknown | null | DOLWD current page lists 2020 NEC; official current AAC/effective date not directly extracted. | src:usa-ak:dolwd:mechanical-inspection-page | partially_verified |
| adoption:usa-ak:dolwd:nesc-2017 | electrical_utility | ahj:usa-ak:dolwd:mechanical-inspection | NESC | 2017 | unknown | null | DOLWD current page lists 2017 NESC. | src:usa-ak:dolwd:mechanical-inspection-page | partially_verified |
| adoption:usa-ak:ahfc:irc-2018 | residential | ahj:usa-ak:ahfc:residential-standards | IRC | 2018 | yes | 2018-11-28 | AHFC residential code for 4-or-fewer dwelling-unit buildings in applicable AHFC/approved-municipality context; 2019-01-01 construction-start applicability for 2018 update. | src:usa-ak:ahfc:codes-standards, src:usa-ak:ahfc:minimum-construction-standards, src:usa-ak:ahfc:2018-code-adoption-packet | partially_verified_scope_limited |
| adoption:usa-ak:ahfc:iecc-2018-bees | energy | ahj:usa-ak:ahfc:bees | IECC and ASHRAE 62.2 | 2018 IECC; ASHRAE 62.2-2016 | yes | 2018-11-28 | BEES for buildings using AHFC or other state financial assistance; minimum 5 Star energy rating. | src:usa-ak:ahfc:bees, src:usa-ak:ahfc:2018-code-adoption-packet | partially_verified_scope_limited |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

The State Fire Marshal's 13 AAC 50 code adoptions show an amendment/effective marker of 2022-10-28 in Register 244 for the 2021 IBC/IEBC/IMC/IFGC/IFC package. This report did not find a separate statewide grace period or concurrency rule for those codes. AHFC residential and energy code materials distinguish construction-start applicability, and AHFC regulatory materials identify the foundation installation date as the construction-began date for residential IRC applicability. DOLWD trade-code current editions are verified from the current Mechanical Inspection page, but official current effective dates for the 2020 NEC and some boiler/elevator standards require a direct official AAC extract.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-ak:001 | 13 AAC 50.020 / .021 / .023 / .024 / .025 | regulation_effective_date | 2022-10-28 | 2021 State Fire Marshal code package amendments shown as am 10/28/2022, Register 244. | unresolved | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | verified_for_effective_marker |
| date-rule:usa-ak:002 | AHFC 2018 IRC residential code | construction_start_applicability | 2019-01-01 | Building construction began on or after 2019-01-01; AHFC materials identify foundation installation date as construction-began date. | no for current affected projects | src:usa-ak:ahfc:2018-code-adoption-packet, src:usa-ak:ahfc:minimum-construction-standards | partially_verified |
| date-rule:usa-ak:003 | AHFC BEES / 2018 IECC update | construction_start_applicability | 2019-01-01 | Building construction began on or after 2019-01-01 for the 2018 IECC/ASHRAE 62.2-2016 update; AHFC/state financial assistance scope applies. | no for current affected projects | src:usa-ak:ahfc:bees, src:usa-ak:ahfc:2018-code-adoption-packet | partially_verified |
| date-rule:usa-ak:004 | Plumbing / 2018 UPC | regulation_effective_date | 2020-04-24 | 8 AAC 63.010 history line in DOLWD April 2020 plumbing publication. | unresolved | src:usa-ak:dolwd:plumbing-code-pdf, src:usa-ak:dolwd:mechanical-inspection-page | partially_verified |
| date-rule:usa-ak:005 | Electrical / 2020 NEC | effective_date_unresolved | null | Current DOLWD page lists 2020 NEC, but official AAC effective date not extracted. | unresolved | src:usa-ak:dolwd:mechanical-inspection-page | unresolved |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Residential / Energy | future ICC updates | 2025-02-05 | null | null | null | null | watch | src:usa-ak:ahfc:building-code-advisory-council | AHFC established a Building Code Advisory Council in 2024 and stated the council will review/update existing building codes on a regular rotation, not less than every three years. No adopted future edition identified. |
| Plumbing / Electrical licensing | licensing and reciprocity changes | 2025-07-25 | null | 2025-08-08 | 2025-08-08 | null | monitor_for_code_effect_only | src:usa-ak:dolwd:mechanical-inspection-page | DOLWD page flags a new regulation effective 2025-08-08 affecting reciprocity/provisional licensing; not treated as model-code edition adoption. |
| 13 AAC 50 fire/life safety | unknown | null | null | null | null | null | watch | src:usa-ak:dps:fire-life-safety-page, src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | No future State Fire Marshal model-code edition adoption verified. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-ak:001 | Building / IBC | Detached one-, two-, and three-family dwellings; specified townhouses | 13 AAC 50.020 IBC Chapter 1 amendment | The 2021 IBC adoption excludes detached one-, two-, and three-family dwellings and treats townhouses with separated egress and other criteria under a specific exception; use AHFC/municipal sources for residential analysis. | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | verified |
| applicability-rule:usa-ak:002 | Fire/life-safety plan review | Construction, alteration, repair, occupancy change, fire protection systems | Before work begins, where 13 AAC 50.027 applies | Plans for covered work and fire protection systems must be submitted to the Division of Fire and Life Safety for fire/life-safety examination and approval; review does not cover every structural/mechanical/electrical consideration beyond life-safety confirmation. | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | verified |
| applicability-rule:usa-ak:003 | Plumbing | New or altered plumbing, gas, and fuel piping installations | Community population 2,500 and above | DOLWD Mechanical Inspection states it has jurisdiction over plumbing work in communities with populations of 2,500 and above; Anchorage inspection is deferred through MOA. | src:usa-ak:dolwd:mechanical-inspection-page | verified_program_scope |
| applicability-rule:usa-ak:004 | Electrical | New and altered electrical installations subject to NEC/NESC | Commercial structures and dwellings of three-plex and above | DOLWD Mechanical Inspection states all commercial structures and dwellings of three-plex and above are subject to inspection; Anchorage inspection is deferred through MOA. | src:usa-ak:dolwd:mechanical-inspection-page | verified_program_scope |
| applicability-rule:usa-ak:005 | AHFC residential and BEES | Residential housing / buildings using AHFC or state financial assistance | AHFC financing or other state financial assistance | AHFC pages state homes may only be financed by AHFC when they meet Minimum Construction Standards and BEES, and BEES applies where AHFC or other state financial assistance is used. | src:usa-ak:ahfc:codes-standards, src:usa-ak:ahfc:bees | verified_scope_limited |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** embedded_regulatory_amendments_by_code_family

**Where amendments are published:**
- 13 AAC 50 — 13 AAC 55 fire/life-safety amendments are published by DPS/State Fire Marshal and available through the Alaska Administrative Code / DPS PDF publication.
- DOLWD plumbing/electrical/boiler/elevator code amendments and adopted references are published in Title 8 AAC and DOLWD program publications.
- AHFC residential and BEES amendments are published as Alaska-specific amendments linked from AHFC's code pages.

**Amendment parsing status:** high_level_only

### 5.2 State Amendment Sources

| Amendment Source ID | Applies To | Publication Path | Status | Caveat |
| --- | --- | --- | --- | --- |
| amend-source:usa-ak:dps:13-aac-50 | IBC, IEBC, IMC, IFGC, IFC, NFPA references | 13 AAC 50 — 13 AAC 55 DPS PDF / official AAC | partially_parsed | Many amendments are detailed section-by-section changes; this report captures only high-impact structure and adoption facts. |
| amend-source:usa-ak:dolwd:plumbing-8-aac-63 | UPC and related IAPMO codes | DOLWD Plumbing Code PDF / Title 8 AAC | partially_parsed | DOLWD PDF says it is informational and may not capture later changes; current DOLWD page confirms 2018 editions. |
| amend-source:usa-ak:dolwd:electrical-8-aac-70 | NEC / NESC | DOLWD current page / Title 8 AAC | high_level_only | Current page confirms edition; current official regulation text and effective date still need direct official AAC extraction. |
| amend-source:usa-ak:ahfc:irc-2018 | Residential / IRC | AHFC 2018 IRC Alaska-specific amendments | high_level_only | Amendments document identified but not parsed section-by-section. |
| amend-source:usa-ak:ahfc:bees-2018 | Energy / IECC / ASHRAE 62.2 | AHFC 2018 IECC Alaska-specific amendments / BEES page | high_level_only | Amendments document identified but not parsed section-by-section. |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Summary | Source IDs | Status |
| --- | --- | --- | --- | --- |
| amendment:usa-ak:dps:ibc-delete-related-codes | Building | The 2021 IBC adoption deletes references to IECC, IPMC, IPSDC, IRC, and IWUIC in Chapter 1 and substitutes DOLWD-adopted electrical and plumbing codes for deleted electrical/plumbing/fuel-gas references. | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | verified |
| amendment:usa-ak:dps:ifc-chapters-reserved | Fire operational | The 2021 IFC adoption reserves Chapters 13-19, 41-49, and 68-79 and adopts Chapters 1-12, 20-40, 50-67, 80 and specified appendices. | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | verified |
| amendment:usa-ak:dps:local-less-stringent-approval | Fire/life-safety local amendments | Local political subdivisions may not set minimum standards less stringent than 13 AAC 50 — 13 AAC 55 without written State Fire Marshal approval; all local fire/life-safety revisions must be submitted. | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | verified |
| amendment:usa-ak:ahfc:bees-5-star | Energy | AHFC BEES requires the 2018 IECC, ASHRAE 62.2-2016, Alaska-specific amendments, and a minimum 5 Star energy rating. | src:usa-ak:ahfc:bees | verified |
| amendment:usa-ak:dolwd:plumbing-lp-gas | Plumbing / Fuel Gas | DOLWD plumbing publication includes Alaska-specific UPC revisions, including restrictions on LPG piping serving gas-fired appliances in pits or basements where heavier-than-air gas might collect. | src:usa-ak:dolwd:plumbing-code-pdf | partially_verified |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-ak"
  model: "state_led_with_deferred_local_authority_and_local_home_rule_overlay"
  enforcing_entities:
    - "state_fire_marshal"
    - "division_of_fire_and_life_safety"
    - "registered_fire_department"
    - "deferred_municipality_building_or_fire_official"
    - "dolwd_mechanical_inspection"
    - "anchorage_local_inspectors_under_moa_for_selected_trade_programs"
    - "approved_municipality_for_ahfc_residential_code_context"
  required_officials:
    - "state_fire_marshal_or_designated_representative"
    - "fire_code_official"
    - "building_official_or_fire_chief_of_deferred_jurisdiction_where_applicable"
    - "dolwd_mechanical_inspection_inspector_or_deferred_local_inspector_where_applicable"
  state_reserved_activities:
    - "13_aac_50_027_fire_life_safety_plan_review_unless_exempted_or_deferred"
    - "state_fire_marshal_approval_for_less_stringent_local_fire_life_safety_minimums"
    - "state_fire_marshal_receipt_of_local_fire_life_safety_revisions"
  source_ids:
    - "src:usa-ak:dps:fire-life-safety-page"
    - "src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf"
    - "src:usa-ak:dolwd:mechanical-inspection-page"
  verification_status: "partially_verified"
  confidence: 0.86
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-ak"
  model: "not_less_stringent_with_state_fire_marshal_submission_or_approval_for_fire_life_safety"
  applies_to_code_families:
    - "building_fire_life_safety"
    - "existing_building_fire_life_safety"
    - "mechanical_fire_life_safety"
    - "fuel_gas_fire_life_safety"
    - "fire_operational"
    - "plumbing_not_less_stringent_local_authority_preserved"
    - "electrical_not_less_stringent_local_authority_preserved"
  approval_required: "yes_for_less_stringent_fire_life_safety_minimums"
  approving_authority_id: "ahj:usa-ak:dps:state-fire-marshal"
  filing_required: "yes_for_proposed_local_fire_life_safety_revisions"
  registry_exists: "unresolved"
  registry_source_ids: []
  legal_basis_source_ids:
    - "src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf"
    - "src:usa-ak:dolwd:plumbing-code-pdf"
    - "src:usa-ak:dolwd:electrical-safety-pdf"
    - "src:usa-ak:constitution:article-x"
  verification_status: "partially_verified"
  confidence: 0.84
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement and local amendment authority must be separated. The State Fire Marshal may defer certain plan-review or building fire-safety inspection/enforcement activities to municipalities or state agencies that have the required expertise, ordinances/agreements, and inspection programs. That deferral does not itself mean a municipality may adopt weaker fire/life-safety standards; 13 AAC 55.030(c) separately requires written State Fire Marshal approval for less-stringent local minimums and submission of proposed fire/life-safety revisions.

DOLWD Mechanical Inspection identifies local deferral for Anchorage inspections in electrical, plumbing, and elevator programs through a memorandum of agreement. Plumbing and electrical legal materials preserve not-less-stringent local standards, but this report did not identify a statewide registry for local plumbing/electrical amendments or approved municipal residential codes.

### 6.4 Known Local Amendment Registries

| Registry | Scope | Source IDs | Status | Notes |
| --- | --- | --- | --- | --- |
| State Fire Marshal local fire/life-safety revision submissions | Local revisions to codes relating to fire and life safety | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | registry_unresolved | Submission requirement verified; public registry path not identified. |
| AHFC approved municipal building codes | Municipal building codes accepted for AHFC residential financing compliance | src:usa-ak:ahfc:minimum-construction-standards | registry_unresolved | AHFC references approved municipalities and approved municipal building codes, but a current roster was not extracted. |
| DOLWD Anchorage MOA deferral | Plumbing/electrical/elevator inspections in Municipality of Anchorage | src:usa-ak:dolwd:mechanical-inspection-page | registry_unresolved | MOA existence verified at program-page level; document not extracted. |

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendment text was parsed. Municipality-specific data should be added only after ordinance extraction and, for fire/life-safety deviations, after confirming State Fire Marshal submission/approval status.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: partially_started

Jurisdiction stack:

```text
Address / project
  -> State of Alaska
  -> Occupancy and project type
  -> Fire/life-safety scope under 13 AAC 50 - 13 AAC 55
  -> State Fire Marshal / Division of Fire and Life Safety plan review unless deferred or exempted
  -> Municipality / borough / city / unorganized area
  -> Registered fire department or deferred fire/building official, if applicable
  -> DOLWD Mechanical Inspection trade jurisdiction for plumbing, electrical, boilers, elevators, conveyances
  -> Anchorage MOA local inspectors for selected trades, if applicable
  -> AHFC/state-financing trigger for residential IRC and BEES requirements
  -> Approved municipal building code status for AHFC-financed residential units
  -> Local ordinances and amendments, subject to not-less-stringent and state-submission/approval rules where applicable
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | Alaska Legislature statutes portal / State of Alaska official pages | src:usa-ak:statutes:portal | statewide | when state law changes | verified_as_index |
| Municipal home rule / local government | Alaska Constitution Article X | src:usa-ak:constitution:article-x | statewide governance framework | constitutional amendments | verified |
| Deferred fire/life-safety jurisdictions | State Fire Marshal deferral agreements under 13 AAC 50.075 | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | statewide where agreements exist | audited every two years under rule | rule_verified_roster_unresolved |
| Approved municipal building codes for AHFC | AHFC approved municipality references | src:usa-ak:ahfc:minimum-construction-standards | AHFC-financed residential scope | unknown | roster_unresolved |
| DOLWD trade deferral | DOLWD Mechanical Inspection page, Anchorage MOA references | src:usa-ak:dolwd:mechanical-inspection-page | Anchorage for selected trade inspections | unknown | partially_verified |
| Fire department registration | DPS Division of Fire and Life Safety Fire Department Registration & Directory | src:usa-ak:dps:fire-life-safety-page | statewide fire/emergency response agencies | unknown | directory_identified_not_extracted |

### 7.3 AHJ Contact Data

No AHJ contact dataset was populated. Source pages identify agency-level contacts for DOLWD Mechanical Inspection and AHFC headquarters, but this report does not include AHJ contact records because jurisdiction-specific contact extraction was outside the current scope.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Source Type | Publisher | URL | Accessed Date | Snapshot ID | Checksum | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-ak:dps:fire-life-safety-page | Fire and Life Safety | agency_page | Alaska Department of Public Safety | https://dps.alaska.gov/fire/ | 2026-06-26 | snapshot_pending | checksum_pending | verified |
| src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | 13 AAC 50 to 13 AAC 55 Fire Prevention, effective 2022-10-28 | regulation_pdf | Alaska Department of Public Safety / Division of Fire and Life Safety | https://dps.alaska.gov/wp-content/uploads/2026/05/13-AAC-50-to-13-AAC-55-Fire-Prevention_effective-10-28-2022.pdf | 2026-06-26 | snapshot_pending | checksum_pending | verified |
| src:usa-ak:dolwd:mechanical-inspection-page | Mechanical Inspection | agency_page | Alaska Department of Labor and Workforce Development, Labor Standards and Safety | https://labor.alaska.gov/lss/mihome.htm | 2026-06-26 | snapshot_pending | checksum_pending | verified |
| src:usa-ak:dolwd:plumbing-code-pdf | Plumbing Code, Statutes and Regulations, April 2020 | statute_regulation_publication_pdf | Alaska Department of Labor and Workforce Development, Mechanical Inspection | https://labor.alaska.gov/lss/forms/Plumbing_Code.pdf | 2026-06-26 | snapshot_pending | checksum_pending | caveated_official_publication |
| src:usa-ak:dolwd:electrical-safety-pdf | Electrical Safety, Statutes and Regulations, May 2018 | statute_regulation_publication_pdf | Alaska Department of Labor and Workforce Development, Mechanical Inspection | https://labor.alaska.gov/lss/forms/electrical-stats-regs.pdf | 2026-06-26 | snapshot_pending | checksum_pending | caveated_official_publication_stale_for_current_nec |
| src:usa-ak:dolwd:elevator-safety-pdf | Elevator Safety, Statutes and Regulations | statute_regulation_publication_pdf | Alaska Department of Labor and Workforce Development, Mechanical Inspection | https://labor.alaska.gov/lss/forms/elevator-stats-regs.pdf | 2026-06-26 | snapshot_pending | checksum_pending | caveated_official_publication |
| src:usa-ak:dolwd:boiler-pressure-vessel-pdf | Boiler and Unfired Pressure Vessel Statutes and Regulations | statute_regulation_publication_pdf | Alaska Department of Labor and Workforce Development, Mechanical Inspection | https://labor.alaska.gov/lss/forms/boiler-stats-regs.pdf | 2026-06-26 | snapshot_pending | checksum_pending | caveated_official_publication |
| src:usa-ak:ahfc:codes-standards | Codes & Standards | agency_page | Alaska Housing Finance Corporation | https://www.ahfc.us/efficiency/codes-standards | 2026-06-26 | snapshot_pending | checksum_pending | verified |
| src:usa-ak:ahfc:minimum-construction-standards | Alaska Minimum Construction Standards | agency_page | Alaska Housing Finance Corporation | https://www.ahfc.us/pros/builders/alaska-minimum-construction-standards | 2026-06-26 | snapshot_pending | checksum_pending | verified |
| src:usa-ak:ahfc:bees | Building Energy Efficiency Standard | agency_page | Alaska Housing Finance Corporation | https://www.ahfc.us/pros/builders/building-energy-efficiency-standard | 2026-06-26 | snapshot_pending | checksum_pending | verified |
| src:usa-ak:ahfc:building-code-advisory-council | Ensuring Safety and Energy Efficiency through Building Code Advisory Council | agency_blog_page | Alaska Housing Finance Corporation | https://www.ahfc.us/blog/posts/ensuring-safety-and-energy-efficiency-through-building-code-advisory-council | 2026-06-26 | snapshot_pending | checksum_pending | verified_for_monitoring |
| src:usa-ak:ahfc:2018-code-adoption-packet | AHFC Board Packet, 2018 IRC/IECC Adoption Materials | board_packet_pdf | Alaska Housing Finance Corporation | https://www.ahfc.us/application/files/8215/3436/9454/082218_board_packet_ahfc.pdf | 2026-06-26 | snapshot_pending | checksum_pending | caveated_adoption_materials |
| src:usa-ak:constitution:article-x | Alaska Constitution, Article X, Local Government | constitution | Office of the Lieutenant Governor | https://ltgov.alaska.gov/information/alaskas-constitution/ | 2026-06-26 | snapshot_pending | checksum_pending | verified |
| src:usa-ak:statutes:portal | Alaska Statutes 2024 portal | statute_portal | Alaska State Legislature | https://www.akleg.gov/basis/statutes.asp | 2026-06-26 | snapshot_pending | checksum_pending | verified_as_index |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | official_pdf_publication | DPS PDF states the publication is also available online at the official AAC page; PDF text extraction was usable and screenshots were spot-checked. | authoritative_for_13_aac_50_55_subject_to_official_aac_snapshot |
| src:usa-ak:dolwd:mechanical-inspection-page | program_page_current | Current DOLWD page is official and lists adopted codes, but it is a program page rather than the official AAC text. | use_for_current_edition_confirmation |
| src:usa-ak:dolwd:plumbing-code-pdf | informational_publication | PDF warns it is an informational guide and may contain errors or omissions after publication. Current DOLWD page separately confirms the 2018 UPC edition. | use_with_current_program_page_crosscheck |
| src:usa-ak:dolwd:electrical-safety-pdf | stale_for_current_edition | May 2018 PDF supports statutory authority and old regulatory structure but lists 2017 NEC, while the current DOLWD page lists 2020 NEC. | do_not_use_for_current_nec_edition |
| src:usa-ak:ahfc:2018-code-adoption-packet | board_packet_adoption_material | Official AHFC board/adoption materials contain proposed regulatory language and transition details; current AHFC pages confirm current code resources. | use_for_transition_context_with_current_page_crosscheck |
| src:usa-ak:statutes:portal | portal_index_only | Statutes portal confirms official statutory repository but chapter-specific statutory text was not directly extracted from the portal in this pass. | browser_required_for_chapter_extract |
| src:usa-ak:constitution:article-x | constitutional_level_only | Useful for local government background but not sufficient alone to define building-code adoption scope. | authoritative_for_governance_only |

### 8.3 Supplemental Sources

| Source ID | Title | Source Type | Publisher | URL | Accessed Date | Status | Use Limitation |
| --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-ak:supp:cornell:15-aac-155-010 | 15 AAC 155.010 Adoption of energy standard and amendments | secondary_regulation_copy | Legal Information Institute, Cornell Law School | https://www.law.cornell.edu/regulations/alaska/15-AAC-155.010 | 2026-06-26 | supplemental | Use only to cross-check AHFC energy regulatory text; not treated as official source. |
| src:usa-ak:supp:justia:15-aac-150-035 | 15 AAC 150.035 Adoption of residential building code and amendments | secondary_regulation_copy | Justia | https://regulations.justia.com/states/alaska/title-15/chapter-150/article-1/section-15-aac-150-035/ | 2026-06-26 | supplemental | Use only to cross-check AHFC residential regulatory text; not treated as official source. |

### 8.4 Source Extraction Metadata

| Source ID | Parser | Parser Version | Extracted At | Extraction Confidence | Requires OCR | Requires Browser Automation | Manual Review Required |
| --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-ak:dps:fire-life-safety-page | browser_text | 0.1 | 2026-06-26T00:00:00Z | 0.94 | no | yes | yes |
| src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | browser_pdf_text_plus_screenshots | 0.1 | 2026-06-26T00:00:00Z | 0.96 | no | yes | yes |
| src:usa-ak:dolwd:mechanical-inspection-page | browser_text | 0.1 | 2026-06-26T00:00:00Z | 0.92 | no | yes | yes |
| src:usa-ak:dolwd:plumbing-code-pdf | browser_pdf_text_plus_screenshots | 0.1 | 2026-06-26T00:00:00Z | 0.82 | no | yes | yes |
| src:usa-ak:dolwd:electrical-safety-pdf | browser_pdf_text_plus_screenshots | 0.1 | 2026-06-26T00:00:00Z | 0.65 | no | yes | yes |
| src:usa-ak:dolwd:elevator-safety-pdf | browser_pdf_text | 0.1 | 2026-06-26T00:00:00Z | 0.70 | no | yes | yes |
| src:usa-ak:dolwd:boiler-pressure-vessel-pdf | browser_pdf_text | 0.1 | 2026-06-26T00:00:00Z | 0.70 | no | yes | yes |
| src:usa-ak:ahfc:codes-standards | browser_text | 0.1 | 2026-06-26T00:00:00Z | 0.90 | no | yes | yes |
| src:usa-ak:ahfc:minimum-construction-standards | browser_text | 0.1 | 2026-06-26T00:00:00Z | 0.88 | no | yes | yes |
| src:usa-ak:ahfc:bees | browser_text | 0.1 | 2026-06-26T00:00:00Z | 0.92 | no | yes | yes |
| src:usa-ak:ahfc:building-code-advisory-council | browser_text | 0.1 | 2026-06-26T00:00:00Z | 0.88 | no | yes | yes |
| src:usa-ak:ahfc:2018-code-adoption-packet | search_snippet_and_pdf_metadata | 0.1 | 2026-06-26T00:00:00Z | 0.72 | no | yes | yes |
| src:usa-ak:constitution:article-x | browser_text | 0.1 | 2026-06-26T00:00:00Z | 0.96 | no | yes | yes |
| src:usa-ak:statutes:portal | browser_text | 0.1 | 2026-06-26T00:00:00Z | 0.82 | no | yes | yes |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report:usa-ak | report.status | partially_verified | verified | 1.00 | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf, src:usa-ak:dolwd:mechanical-inspection-page, src:usa-ak:ahfc:bees | Core authority and code-adoption fields are source-backed, but local registries and some effective dates remain unresolved. |
| ahj:usa-ak:dps:state-fire-marshal | authority.name | Alaska Department of Public Safety, Division of Fire and Life Safety / State Fire Marshal | verified | 0.94 | src:usa-ak:dps:fire-life-safety-page, src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | Primary fire/life-safety authority identified. |
| adoption:usa-ak:dps:ibc-2021 | edition | 2021 IBC | verified | 0.92 | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | Section-level adoption captured. |
| adoption:usa-ak:dps:ifc-2021 | edition | 2021 IFC | verified | 0.94 | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | Reserved/adopted chapter split captured. |
| adoption:usa-ak:dolwd:upc-2018 | edition | 2018 UPC | partially_verified | 0.85 | src:usa-ak:dolwd:mechanical-inspection-page, src:usa-ak:dolwd:plumbing-code-pdf | Current page and 2020 PDF align. |
| adoption:usa-ak:dolwd:nec-2020 | edition | 2020 NEC | partially_verified | 0.86 | src:usa-ak:dolwd:mechanical-inspection-page | Current official program page supports edition; official AAC text/effective date still needed. |
| adoption:usa-ak:ahfc:irc-2018 | edition | 2018 IRC with Alaska amendments | partially_verified | 0.82 | src:usa-ak:ahfc:codes-standards, src:usa-ak:ahfc:minimum-construction-standards, src:usa-ak:ahfc:2018-code-adoption-packet | Scope is AHFC/approved-municipal-code context, not universal statewide private residential coverage. |
| adoption:usa-ak:ahfc:iecc-2018-bees | edition | 2018 IECC / ASHRAE 62.2-2016 / Alaska amendments | verified_scope_limited | 0.88 | src:usa-ak:ahfc:bees | BEES scope and minimum 5 Star requirement captured. |
| local-amendment-rule:usa-ak | fire_life_safety_rule | less-stringent local minimums require written State Fire Marshal approval and proposed fire/life-safety revisions must be submitted | verified | 0.92 | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | Applies to 13 AAC 50 — 13 AAC 55 fire/life-safety scope. |
| local-enforcement:usa-ak | model | state_led_with_deferred_local_authority_and_local_home_rule_overlay | partially_verified | 0.86 | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf, src:usa-ak:dolwd:mechanical-inspection-page, src:usa-ak:constitution:article-x | Deferral rules verified; actual roster unresolved. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Source IDs used in the body are listed in section 8. |
| All authority IDs resolve | pass | Authority IDs used in this file are defined in section 2 or tied to an adoption record. |
| All current code families have adoption rows | pass | Code-family rows are present; unresolved details remain explicit. |
| Building and operational fire code are separated | pass | IBC/building and IFC/fire-operational rows are distinct. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Unknown fields are left null and transition notes explain scope. |
| Effective dates are valid ISO dates | pass | Entered dates use YYYY-MM-DD. |
| No impossible date sequences | pass | No contradictory date sequences identified. |
| Transition rules have explicit trigger conditions | partial | AHFC trigger conditions captured; State Fire Marshal grace/concurrency rule unresolved. |
| Permit-date logic is captured where applicable | partial | AHFC construction-start/foundation logic captured; permit-date logic for 13 AAC 50 not identified. |
| Local enforcement model classified | pass | State-led with deferral and local overlay. |
| Local amendment rule classified | partial | Fire/life-safety rule verified; plumbing/electrical and approved-municipal-code registries unresolved. |
| AHJ confirmation metadata present | fail | No jurisdiction-specific AHJ contacts or deferred-jurisdiction roster entered. |
| Official-source caveats captured | pass | Official-source caveats recorded in section 8.2. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-ak:001 | high | official Title 8 AAC current text | DOLWD current program page verifies current adopted trade-code editions, but direct official AAC extraction for 8 AAC 63, 70, 77, and 80 was incomplete. | Extract current official AAC sections from Alaska Legislature AAC site or official published supplement; update effective dates and citations. | null | null | open |
| issue:usa-ak:002 | high | deferred fire/life-safety jurisdictions | 13 AAC 50.075 deferral criteria are verified, but the roster of current deferred municipalities/agencies was not extracted. | Obtain State Fire Marshal deferral/exemption agreements or current list. | null | null | open |
| issue:usa-ak:003 | high | AHFC approved municipal building codes | AHFC references approved municipal building codes, but a current approved municipality/code list was not populated. | Extract AHFC approved municipalities and map AHFC code applicability by jurisdiction. | null | null | open |
| issue:usa-ak:004 | medium | residential universality | AHFC residential code is financing-linked; statewide requirements for non-AHFC private detached residential projects need local ordinance and AHFC/municipal analysis. | Build residential decision tree by financing trigger and municipality. | null | null | open |
| issue:usa-ak:005 | medium | trade-code local amendment workflow | Plumbing and electrical sources preserve not-less-stringent local authority, but filing/approval/registry details outside fire/life-safety scope are unresolved. | Extract local standards rules and any municipal filing process for plumbing/electrical. | null | null | open |
| issue:usa-ak:006 | medium | amendment parsing | Alaska amendments to IBC/IEBC/IMC/IFGC/IFC, AHFC IRC/IECC, and DOLWD trade codes are only high-level parsed. | Convert amendment documents to structured amendment records. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-ak:dps-fire-page | src:usa-ak:dps:fire-life-safety-page | html_diff | monthly | Fire/life-safety page changes, plan-review links, adopted-code language changes | 2026-06-26 | active |
| watch:usa-ak:dps-13-aac-50-pdf | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf | pdf_diff | monthly | New 13 AAC 50 — 55 PDF or amendment effective date appears | 2026-06-26 | active |
| watch:usa-ak:dolwd-mi | src:usa-ak:dolwd:mechanical-inspection-page | html_diff | monthly | Adopted-code list changes for plumbing, electrical, boiler, elevator, or jurisdiction text | 2026-06-26 | active |
| watch:usa-ak:ahfc-codes | src:usa-ak:ahfc:codes-standards | html_diff | monthly | AHFC IRC/IECC links or current code editions change | 2026-06-26 | active |
| watch:usa-ak:ahfc-bees | src:usa-ak:ahfc:bees | html_diff | monthly | BEES edition, 5 Star requirement, or applicability language changes | 2026-06-26 | active |
| watch:usa-ak:constitution | src:usa-ak:constitution:article-x | html_diff | quarterly | Local-government constitutional text changes or site updates | 2026-06-26 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-26 | Replaced Alaska home-rule-only draft with a partially verified hybrid state-and-local report. Added State Fire Marshal 2021 IBC/IEBC/IMC/IFGC/IFC adoptions, DOLWD trade-code program adoptions, AHFC IRC/BEES scope, local amendment constraints, and QA/open issues. | report:usa-ak; ahj:usa-ak:dps:state-fire-marshal; ahj:usa-ak:dolwd:mechanical-inspection; ahj:usa-ak:ahfc:residential-standards; ahj:usa-ak:ahfc:bees | src:usa-ak:dps:aac-13-50-55-fire-prevention-pdf, src:usa-ak:dolwd:mechanical-inspection-page, src:usa-ak:ahfc:bees | ChatGPT | Status remains partially_verified due unresolved local rosters, Title 8 AAC direct extract, and amendment parsing. |
