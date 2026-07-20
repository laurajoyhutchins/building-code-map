---
state:
  state_id: "US-MN"
  name: "Minnesota"
  abbreviation: "MN"
report:
  report_id: "state-report:usa-mn"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-25"
  last_verified: null
  reviewed_by: null
risk:
  overall_confidence: 0.72 # 0.00 - 1.00
  risk_flags:
    - "fire_code_effective_date_unresolved"
    - "local_fire_amendment_scope_requires_legal_review"
    - "detailed_state_amendments_not_fully_parsed"
    - "ahj_boundary_and_contact_data_not_populated"
    - "future_code_cycle_final_dates_unresolved"
  open_questions_count: 5

---

# State Building Code Authority Report: Minnesota

## 1. Executive Summary

- **Authority model:** Minnesota has a statewide State Building Code adopted and amended by the Commissioner of Labor and Industry. The Department of Labor and Industry, Construction Codes and Licensing Division, administers the state construction-code program, while the State Fire Code is adopted by the Commissioner of Labor and Industry consistent with State Fire Marshal recommendations and is enforced under the Commissioner of Public Safety through the State Fire Marshal framework. Source IDs: `src:usa-mn:stat-326b-02`, `src:usa-mn:stat-326b-101`, `src:usa-mn:stat-326b-106`, `src:usa-mn:stat-299f-011`, `src:usa-mn:dli-ccld-overview`.

- **Statewide code status:** The State Building Code is the minimum construction standard throughout Minnesota. DLI states that the 2020 Minnesota State Building Code became effective 2020-03-31, except the Minnesota Mechanical Fuel Gas Code became effective 2020-04-06. DLI also identifies the 2024 Minnesota Commercial Energy Code as effective 2024-01-05. Source IDs: `src:usa-mn:dli-state-building-codes-2020`, `src:usa-mn:dli-building-code-overview`.

- **Local enforcement model:** The State Building Code is statewide, but municipal enforceability depends on municipal administration/enforcement. Municipalities with a State Building Code ordinance in effect on 2008-01-01 generally must continue administration and enforcement; other municipalities may elect to administer and enforce by ordinance. State administration remains important for public buildings, state-licensed facilities, and several trade-specific programs. Source IDs: `src:usa-mn:stat-326b-121`, `src:usa-mn:stat-326b-107`, `src:usa-mn:stat-326b-106`.

- **Local amendment posture:** Municipal building-code divergence is generally preempted: a municipality may not require building-code provisions different from the State Building Code through ordinance or development agreement. Limited exceptions include maintenance ordinances that do not exceed the original code basis unless the State Building Code has retroactive provisions, geological-condition ordinances approved by the state building official, and optional provisions authorized in Minnesota Rules chapter 1300. Fire-code local ordinances may be equal, additional, or more stringent under Minnesota Statutes section 299F.011, but must relate to fire/life/property protection and may not exceed applicable State Building Code requirements. Source IDs: `src:usa-mn:stat-326b-121`, `src:usa-mn:rule-1300`, `src:usa-mn:stat-299f-011`.

- **Known transition periods or pending changes:** DLI publishes distinct effective dates for the 2020 State Building Code, the Mechanical Fuel Gas Code, the 2024 Commercial Energy Code, and the 2023 Minnesota Electrical Code. The electrical transition is permit-date based: permits filed before 2023-07-01 follow the 2020 NEC; permits filed on or after 2023-07-01 follow the 2023 NEC. Minnesota statutes also establish recurring review/adoption cycles for model building and energy codes. Source IDs: `src:usa-mn:dli-state-building-codes-2020`, `src:usa-mn:dli-electrical-codes`, `src:usa-mn:dli-energy-conservation`, `src:usa-mn:stat-326b-106`.

- **Production readiness:** partially_verified_for_core_authority_and_current_adoptions; not_ready_for_full_ahj_resolution

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: Commissioner of Labor and Industry adopts, administers, and amends the
    State Building Code under chapter 326B authority.
  confidence: 0.88
  source_ids:
  - src:usa-mn:stat-326b-101
  - src:usa-mn:stat-326b-106
  - src:usa-mn:dli-ccld-overview
- topic: Primary building code edition
  finding: Minnesota Building Code incorporates the 2018 IBC into the Minnesota State
    Building Code as amended in chapter 1305.
  confidence: 0.86
  source_ids:
  - src:usa-mn:rule-1305
  - src:usa-mn:dli-state-building-codes-2020
- topic: Residential code edition
  finding: Minnesota Residential Code incorporates the 2018 IRC, with listed mandatory
    chapters and Minnesota amendments.
  confidence: 0.84
  source_ids:
  - src:usa-mn:rule-1309
  - src:usa-mn:dli-state-building-codes-2020
- topic: Electrical code authority and edition
  finding: Minnesota Electrical Code currently incorporates the 2023 NEC; DLI states
    the Board of Electricity adopted the 2023 NEC effective 2023-07-01.
  confidence: 0.88
  source_ids:
  - src:usa-mn:rule-1315
  - src:usa-mn:dli-electrical-codes
  - src:usa-mn:stat-326b-02
- topic: Fire code authority
  finding: State Fire Code is adopted by the Commissioner of Labor and Industry consistent
    with State Fire Marshal recommendations and enforced through the Public Safety/State
    Fire Marshal structure.
  confidence: 0.78
  source_ids:
  - src:usa-mn:stat-326b-02
  - src:usa-mn:stat-326b-106
  - src:usa-mn:stat-299f-011
  - src:usa-mn:stat-299f-01
- topic: Fire code edition
  finding: Minnesota Rules chapter 7511 incorporates the 2018 IFC into the Minnesota
    State Fire Code as amended. The precise effective date was not independently resolved
    from the rule text in this pass.
  confidence: 0.72
  source_ids:
  - src:usa-mn:rule-7511
- topic: Local enforcement
  finding: The code is statewide, but municipal administration/enforcement is tied
    to required continuation or ordinance adoption.
  confidence: 0.86
  source_ids:
  - src:usa-mn:stat-326b-121
  - src:usa-mn:dli-building-code-overview
- topic: Local amendments
  finding: Different municipal building-code provisions are generally barred; optional
    provisions and approved geological-condition ordinances are limited exceptions.
  confidence: 0.82
  source_ids:
  - src:usa-mn:stat-326b-121
  - src:usa-mn:rule-1300
- topic: Electrical transition rule
  finding: Electrical permits filed before 2023-07-01 follow the 2020 NEC; permits
    filed on or after 2023-07-01 follow the 2023 NEC.
  confidence: 0.9
  source_ids:
  - src:usa-mn:dli-electrical-codes
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | `ahj:usa-mn:dli-commissioner` |
| Authority name | Minnesota Commissioner of Labor and Industry / Department of Labor and Industry, Construction Codes and Licensing Division |
| Authority type | state_agency_commissioner |
| Legal basis | Minnesota Statutes §§ 326B.02, 326B.101, 326B.106 |
| Role | Adopt, administer, and amend the State Building Code; administer construction-code programs through DLI/CCLD; enforce specified trade-code provisions under statutory direction. |
| Enforcement model | statewide_standard_with_municipal_administration_where_required_or_adopted; state_reserved_trade_and_public_facility_roles |
| Source IDs | `src:usa-mn:stat-326b-02`; `src:usa-mn:stat-326b-101`; `src:usa-mn:stat-326b-106`; `src:usa-mn:dli-ccld-overview`; `src:usa-mn:dli-building-code-overview` |
| Verification status | verified_core_authority |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | `ahj:usa-mn:dli-commissioner` | Minnesota Commissioner of Labor and Industry / DLI CCLD | Adopt and administer the State Building Code and Minnesota Building Code. | Minn. Stat. §§ 326B.101, 326B.106; Minn. R. chs. 1300, 1305 | `src:usa-mn:stat-326b-101`; `src:usa-mn:stat-326b-106`; `src:usa-mn:rule-1300`; `src:usa-mn:rule-1305` | verified_core |
| Residential | `ahj:usa-mn:dli-commissioner` | Minnesota Commissioner of Labor and Industry / DLI CCLD | Adopt and administer Minnesota Residential Code. | Minn. Stat. § 326B.106; Minn. R. ch. 1309 | `src:usa-mn:stat-326b-106`; `src:usa-mn:rule-1309` | verified_core |
| Existing Building / Rehabilitation | `ahj:usa-mn:dli-commissioner` | Minnesota Commissioner of Labor and Industry / DLI CCLD | Adopt and administer Minnesota Conservation Code for Existing Buildings. | Minn. Stat. § 326B.106; Minn. R. ch. 1311 | `src:usa-mn:stat-326b-106`; `src:usa-mn:rule-1311` | verified_core |
| Mechanical | `ahj:usa-mn:dli-commissioner` | Minnesota Commissioner of Labor and Industry / DLI CCLD | Adopt and administer Minnesota Mechanical Code. | Minn. Stat. § 326B.106; Minn. R. ch. 1346 | `src:usa-mn:stat-326b-106`; `src:usa-mn:rule-1346` | verified_core |
| Plumbing | `ahj:usa-mn:plumbing-board-dli` | Minnesota Plumbing Board / DLI | Plumbing-code rule chapter and transferred State Plumbing Code responsibilities, with express Plumbing Board exceptions. | Minn. Stat. § 326B.02; Minn. R. ch. 4714 | `src:usa-mn:stat-326b-02`; `src:usa-mn:rule-4714` | verified_core |
| Fuel Gas | `ahj:usa-mn:dli-commissioner` | Minnesota Commissioner of Labor and Industry / DLI CCLD | Adopt and administer Minnesota Fuel Gas Code within chapter 1346. | Minn. Stat. § 326B.106; Minn. R. ch. 1346 | `src:usa-mn:stat-326b-106`; `src:usa-mn:rule-1346` | verified_core |
| Electrical | `ahj:usa-mn:board-electricity-dli` | Minnesota Board of Electricity / DLI | Adopt and administer Minnesota Electrical Code; DLI states the Board of Electricity adopted the 2023 NEC. | Minn. Stat. § 326B.02; Minn. R. ch. 1315 | `src:usa-mn:stat-326b-02`; `src:usa-mn:rule-1315`; `src:usa-mn:dli-electrical-codes` | verified_core |
| Energy | `ahj:usa-mn:dli-commissioner` | Minnesota Commissioner of Labor and Industry / DLI CCLD | Adopt and administer residential and commercial energy codes. | Minn. Stat. § 326B.106; Minn. R. chs. 1322, 1323 | `src:usa-mn:stat-326b-106`; `src:usa-mn:rule-1322`; `src:usa-mn:rule-1323`; `src:usa-mn:dli-energy-conservation` | verified_core |
| Fire - construction references | `ahj:usa-mn:dli-commissioner` | Minnesota Commissioner of Labor and Industry / DLI CCLD | Coordinate construction-code references to fire, fuel gas, mechanical, electrical, and related codes through State Building Code rules. | Minn. Stat. § 326B.106; Minn. R. chs. 1305, 1300 | `src:usa-mn:stat-326b-106`; `src:usa-mn:rule-1305`; `src:usa-mn:rule-1300` | partially_verified |
| Fire - operational / prevention code | `ahj:usa-mn:state-fire-marshal` | Commissioner of Public Safety / State Fire Marshal | Administer and enforce State Fire Code under chapter 299F; State Fire Marshal Division is within Public Safety. | Minn. Stat. §§ 299F.01, 299F.011; Minn. Stat. § 326B.106; Minn. R. ch. 7511 | `src:usa-mn:stat-299f-01`; `src:usa-mn:stat-299f-011`; `src:usa-mn:stat-326b-106`; `src:usa-mn:rule-7511`; `src:usa-mn:rule-7511-0202` | partially_verified |
| Accessibility | `ahj:usa-mn:dli-commissioner` | Minnesota Commissioner of Labor and Industry / DLI CCLD | Adopt and administer Minnesota Accessibility Code. | Minn. Stat. ch. 326B; Minn. R. ch. 1341 | `src:usa-mn:rule-1341`; `src:usa-mn:dli-state-building-codes-2020`; `src:usa-mn:dli-ccld-overview` | verified_core |
| Elevator / Conveyance | `ahj:usa-mn:dli-elevator-program` | Minnesota Commissioner of Labor and Industry / DLI | Administer and enforce elevator-code provisions statewide except statutory municipal option. | Minn. Stat. § 326B.106; Minn. R. ch. 1307 | `src:usa-mn:stat-326b-106`; `src:usa-mn:rule-1307` | verified_core |

### 2.2.1 Additional Named Officials and Programs

| Authority ID | Authority Name | Role | Source IDs | Status |
| --- | --- | --- | --- | --- |
| `ahj:usa-mn:state-building-official` | Minnesota State Building Official | Approves or disapproves more restrictive municipal building ordinances where geological conditions warrant; also appears in public-building/state-licensed-facility plan-review context. | `src:usa-mn:stat-326b-121`; `src:usa-mn:stat-326b-107` | partially_verified |
| `ahj:usa-mn:dli-ccld` | DLI Construction Codes and Licensing Division | DLI program division that administers construction-code and licensing programs under the Commissioner of Labor and Industry authority model. | `src:usa-mn:dli-ccld-overview`; `src:usa-mn:stat-326b-106` | verified_core |

### 2.3 Authority Hierarchy Notes

Minnesota uses a statewide code plus local-administration model. The State Building Code is the statewide standard and supersedes municipal building codes, while municipal administration/enforcement depends on statute and local ordinance status. DLI/CCLD administers the state construction-code program and reserves or directly handles certain state functions, including public-building/state-licensed-facility plan review unless a municipal agreement applies, and statewide elevator enforcement except as statutorily provided. The State Fire Code has a hybrid model: the Commissioner of Labor and Industry adopts the fire code consistent with State Fire Marshal recommendations, while enforcement is under the Commissioner of Public Safety through the State Fire Marshal framework. Source IDs: `src:usa-mn:stat-326b-02`, `src:usa-mn:stat-326b-106`, `src:usa-mn:stat-326b-107`, `src:usa-mn:stat-326b-121`, `src:usa-mn:stat-299f-011`.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| `edge:usa-mn:001` | `ahj:usa-mn:dli-commissioner` | adopts_administers_amends | State Building Code | `src:usa-mn:stat-326b-101`; `src:usa-mn:stat-326b-106` | verified_core |
| `edge:usa-mn:002` | `ahj:usa-mn:dli-ccld` | administers_programs | accessibility, boiler, building, electrical, elevator, energy, high-pressure piping, manufactured structures, plumbing, residential codes | `src:usa-mn:dli-ccld-overview` | verified_core |
| `edge:usa-mn:003` | `ahj:usa-mn:dli-commissioner` | adopts_fire_code_consistent_with | `ahj:usa-mn:state-fire-marshal` recommendations | `src:usa-mn:stat-326b-02` | verified_core |
| `edge:usa-mn:004` | `ahj:usa-mn:state-fire-marshal` | enforces | State Fire Code under Commissioner of Public Safety | `src:usa-mn:stat-326b-106`; `src:usa-mn:stat-299f-011` | verified_core |
| `edge:usa-mn:005` | `ahj:usa-mn:dli-commissioner` | administers_and_enforces | public buildings and state-licensed facilities unless municipal agreement applies | `src:usa-mn:stat-326b-107` | verified_core |
| `edge:usa-mn:006` | `ahj:usa-mn:dli-commissioner` | enforces_statewide | elevator code provisions except statutory municipal option | `src:usa-mn:stat-326b-106`; `src:usa-mn:rule-1307` | verified_core |
| `edge:usa-mn:007` | `ahj:usa-mn:dli-commissioner` | directs_dli_enforcement_of | electrical installations, plumbing, boilers, high-pressure steam piping, ammonia refrigeration piping | `src:usa-mn:stat-326b-106` | verified_core |
| `edge:usa-mn:008` | State Building Code | supersedes | municipal building codes | `src:usa-mn:stat-326b-121` | verified_core |
| `edge:usa-mn:009` | municipalities | administer_and_enforce_when_required_or_adopted | State Building Code within local jurisdiction | `src:usa-mn:stat-326b-121`; `src:usa-mn:dli-building-code-overview` | verified_core |
| `edge:usa-mn:010` | municipalities_with_state_code | may_adopt_without_change | optional chapter 1306 special fire-protection systems and IBC Appendix J grading | `src:usa-mn:rule-1300` | verified_core |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | 2020 Minnesota Building Code | International Building Code | 2018 | current | null | 2020-03-31 | 2020-03-31 | 2020-03-31 where the State Building Code is locally administered; statewide minimum standard | DLI states 2020 State Building Code effective 2020-03-31; municipal administration/enforcement follows § 326B.121. | `src:usa-mn:dli-state-building-codes-2020`; `src:usa-mn:rule-1305`; `src:usa-mn:stat-326b-121` |
| Residential | 2020 Minnesota Residential Code | International Residential Code | 2018 | current | null | 2020-03-31 | 2020-03-31 | 2020-03-31 where the State Building Code is locally administered; statewide minimum standard | Chapter 1309 incorporates 2018 IRC chapters and replaces energy/mechanical/plumbing/electrical content with Minnesota chapters. | `src:usa-mn:dli-state-building-codes-2020`; `src:usa-mn:rule-1309`; `src:usa-mn:stat-326b-121` |
| Existing Building / Rehabilitation | Minnesota Conservation Code for Existing Buildings | International Existing Building Code | 2018 | current | null | 2020-03-31 | 2020-03-31 | 2020-03-31 where the State Building Code is locally administered; statewide minimum standard | IEBC chapters 2 through 16 incorporated as Minnesota amendments. | `src:usa-mn:dli-state-building-codes-2020`; `src:usa-mn:rule-1311` |
| Mechanical | 2020 Minnesota Mechanical Code | International Mechanical Code; NFPA 96; ASHRAE 62.2; ASHRAE 154 | 2018 IMC / 2017 NFPA 96 / 2016 ASHRAE standards | current | null | 2020-04-06 | 2020-04-06 | 2020-04-06 where the State Building Code is locally administered; statewide minimum standard | DLI states the Mechanical Fuel Gas Code effective date separately from the rest of the 2020 State Building Code. | `src:usa-mn:dli-state-building-codes-2020`; `src:usa-mn:rule-1346` |
| Plumbing | Minnesota Plumbing Code | Uniform Plumbing Code | 2018 | current | null | null | null | null | Effective-date field was not resolved from the rule text in this pass; chapter 4714 incorporates the 2018 UPC with Minnesota amendments. | `src:usa-mn:rule-4714`; `src:usa-mn:stat-326b-02` |
| Fuel Gas | Minnesota Fuel Gas Code | International Fuel Gas Code | 2018 | current | null | 2020-04-06 | 2020-04-06 | 2020-04-06 where the State Building Code is locally administered; statewide minimum standard | DLI states the Mechanical Fuel Gas Code effective date separately. | `src:usa-mn:dli-state-building-codes-2020`; `src:usa-mn:rule-1346` |
| Electrical | Minnesota Electrical Code | National Electrical Code / ANSI-NFPA 70 | 2023 | current | null | 2023-07-01 | 2023-07-01 | 2023-07-01 for permits filed on or after that date | Permit-date transition: permits before 2023-07-01 use the 2020 NEC; permits on or after 2023-07-01 use the 2023 NEC. | `src:usa-mn:rule-1315`; `src:usa-mn:dli-electrical-codes` |
| Energy | Minnesota Residential Energy Code; Minnesota Commercial Energy Code | Residential: 2012 IECC Residential Provisions; Commercial: ANSI/ASHRAE/IES 90.1 | Residential 2012 IECC; Commercial 2019 ASHRAE 90.1 | current_mixed | null | residential: 2015-02-14; commercial: 2024-01-05 | residential: 2015-02-14; commercial: 2024-01-05 | same as effective dates where applicable | Energy is split between residential chapter 1322 and commercial chapter 1323; matrix row intentionally points to separate normalized records below. | `src:usa-mn:dli-energy-conservation`; `src:usa-mn:rule-1322`; `src:usa-mn:rule-1323`; `src:usa-mn:stat-326b-106` |
| Fire - construction references | Minnesota construction-code fire references | Coordinated with Minnesota State Fire Code and construction-code chapters | 2018 IFC references plus Minnesota code references | current | null | null | null | null | Building-code cross-references are coordinated by chapter 1305; full fire-code effective-date resolution remains open. | `src:usa-mn:rule-1305`; `src:usa-mn:rule-7511` |
| Fire - operational / prevention code | Minnesota State Fire Code | International Fire Code | 2018 | current | null | null | null | null | Chapter 7511 incorporates the 2018 IFC as amended; precise effective date not resolved from source text captured here. | `src:usa-mn:rule-7511`; `src:usa-mn:stat-299f-011`; `src:usa-mn:rule-7511-0202` |
| Accessibility | 2020 Minnesota Accessibility Code | IBC Chapter 11; IEBC section 305; ICC/ANSI A117.1 | 2018 IBC / 2018 IEBC / 2009 A117.1 | current | null | 2020-03-31 | 2020-03-31 | 2020-03-31 where the State Building Code is locally administered; statewide minimum standard | DLI lists 2020 Accessibility Code under the 2020 State Building Code; chapter 1341 contains the incorporated accessibility provisions. | `src:usa-mn:dli-state-building-codes-2020`; `src:usa-mn:rule-1341` |
| Elevator / Conveyance | Elevators and Related Devices | ASME A17.1/CSA B44, A17.3, A17.5, A17.7, A18.1, A90.1, B20.1 | multiple referenced standards | current | null | null | null | statewide except statutory municipal option | Commissioner administers and enforces elevator provisions statewide except as provided by statute. | `src:usa-mn:stat-326b-106`; `src:usa-mn:rule-1307` |

### 3.2 Adoption Records

| Adoption Record ID | Code Family | State Code | Base Code / Standard | Edition | Adoption Date | Effective Date | Operative Date | Mandatory Date | Applicability | Source IDs | Verification Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `adopt:usa-mn:building:2020` | Building | 2020 Minnesota Building Code | 2018 IBC | 2018 | null | 2020-03-31 | 2020-03-31 | 2020-03-31 where administered locally | IBC chapters 2-35 administered by municipalities that adopted the State Building Code. | `src:usa-mn:dli-state-building-codes-2020`; `src:usa-mn:rule-1305` | verified_core |
| `adopt:usa-mn:residential:2020` | Residential | 2020 Minnesota Residential Code | 2018 IRC | 2018 | null | 2020-03-31 | 2020-03-31 | 2020-03-31 where administered locally | IRC chapters 2-10 and selected appendices/provisions incorporated with Minnesota amendments. | `src:usa-mn:dli-state-building-codes-2020`; `src:usa-mn:rule-1309` | verified_core |
| `adopt:usa-mn:existing:2020` | Existing Building / Rehabilitation | Minnesota Conservation Code for Existing Buildings | 2018 IEBC | 2018 | null | 2020-03-31 | 2020-03-31 | 2020-03-31 where administered locally | IEBC chapters 2-16 incorporated with Minnesota amendments. | `src:usa-mn:dli-state-building-codes-2020`; `src:usa-mn:rule-1311` | verified_core |
| `adopt:usa-mn:mechanical:2020` | Mechanical | 2020 Minnesota Mechanical Code | 2018 IMC plus listed standards | 2018 | null | 2020-04-06 | 2020-04-06 | 2020-04-06 where administered locally | Mechanical code chapter 1346. | `src:usa-mn:dli-state-building-codes-2020`; `src:usa-mn:rule-1346` | verified_core |
| `adopt:usa-mn:fuel-gas:2020` | Fuel Gas | Minnesota Fuel Gas Code | 2018 IFGC | 2018 | null | 2020-04-06 | 2020-04-06 | 2020-04-06 where administered locally | IFGC chapters 2-8 incorporated into chapter 1346. | `src:usa-mn:dli-state-building-codes-2020`; `src:usa-mn:rule-1346` | verified_core |
| `adopt:usa-mn:electrical:2023` | Electrical | Minnesota Electrical Code | 2023 NEC / ANSI-NFPA 70 | 2023 | null | 2023-07-01 | 2023-07-01 | 2023-07-01 for permits filed on or after that date | Permit-date transition confirmed by DLI. | `src:usa-mn:rule-1315`; `src:usa-mn:dli-electrical-codes` | verified_core |
| `adopt:usa-mn:energy-residential:2015` | Energy | Minnesota Residential Energy Code | 2012 IECC Residential Provisions | 2012 | null | 2015-02-14 | 2015-02-14 | 2015-02-14 where applicable | Chapter 1322, residential energy. | `src:usa-mn:dli-energy-conservation`; `src:usa-mn:rule-1322` | verified_core |
| `adopt:usa-mn:energy-commercial:2024` | Energy | Minnesota Commercial Energy Code | ANSI/ASHRAE/IES Standard 90.1 | 2019 | null | 2024-01-05 | 2024-01-05 | 2024-01-05 where applicable | Chapter 1323, commercial energy. | `src:usa-mn:dli-energy-conservation`; `src:usa-mn:rule-1323` | verified_core |
| `adopt:usa-mn:plumbing:2018-upc` | Plumbing | Minnesota Plumbing Code | 2018 UPC | 2018 | null | null | null | null | Chapter 4714 incorporates 2018 UPC with amendments; date fields unresolved. | `src:usa-mn:rule-4714` | partially_verified |
| `adopt:usa-mn:fire:2018-ifc` | Fire - operational / prevention code | Minnesota State Fire Code | 2018 IFC | 2018 | null | null | null | null | Chapter 7511 incorporates 2018 IFC as amended; effective date unresolved. | `src:usa-mn:rule-7511`; `src:usa-mn:stat-299f-011` | partially_verified |
| `adopt:usa-mn:accessibility:2020` | Accessibility | Minnesota Accessibility Code | 2018 IBC Ch. 11; 2018 IEBC § 305; 2009 ICC/ANSI A117.1 | mixed | null | 2020-03-31 | 2020-03-31 | 2020-03-31 where administered locally | Chapter 1341 accessibility provisions. | `src:usa-mn:dli-state-building-codes-2020`; `src:usa-mn:rule-1341` | verified_core |
| `adopt:usa-mn:elevator:1307` | Elevator / Conveyance | Elevators and Related Devices | ASME A17 and related standards | mixed | null | null | null | statewide except statutory municipal option | Chapter 1307 includes referenced elevator/conveyance standards. | `src:usa-mn:rule-1307`; `src:usa-mn:stat-326b-106` | partially_verified |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

Minnesota distinguishes several date concepts in practice, but not all adoption dates were resolved from the available source text. The highest-confidence current dates are DLI-published effective dates for the 2020 State Building Code, Mechanical Fuel Gas Code, Commercial Energy Code, Residential Energy Code, and 2023 Electrical Code. Electrical code transition is explicitly permit-date based. Local administration/enforcement is driven by municipal ordinance status and statutory continuation rules rather than by a single statewide municipal-enforcement date.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `date-rule:usa-mn:2020-state-building-code-effective` | 2020 State Building Code except Mechanical Fuel Gas Code | effective_date | 2020-03-31 | DLI-published effective date for 2020 State Building Code | unresolved | `src:usa-mn:dli-state-building-codes-2020` | verified_core |
| `date-rule:usa-mn:mechanical-fuel-gas-effective` | Minnesota Mechanical Fuel Gas Code | effective_date | 2020-04-06 | DLI-published exception date | unresolved | `src:usa-mn:dli-state-building-codes-2020`; `src:usa-mn:rule-1346` | verified_core |
| `date-rule:usa-mn:electrical-permit-date-2023-nec` | Minnesota Electrical Code | permit_date_transition | 2023-07-01 | Electrical permit filed on or after 2023-07-01 | Yes, for permits filed before 2023-07-01 under the 2020 NEC | `src:usa-mn:dli-electrical-codes`; `src:usa-mn:rule-1315` | verified_core |
| `date-rule:usa-mn:residential-energy-effective` | Residential Energy Code | effective_date | 2015-02-14 | DLI-published chapter 1322 effective date | unresolved | `src:usa-mn:dli-energy-conservation`; `src:usa-mn:rule-1322` | verified_core |
| `date-rule:usa-mn:commercial-energy-effective` | Commercial Energy Code | effective_date | 2024-01-05 | DLI-published chapter 1323 effective date | unresolved | `src:usa-mn:dli-energy-conservation`; `src:usa-mn:rule-1323` | verified_core |
| `date-rule:usa-mn:municipal-continuation-2008` | Local building-code administration/enforcement | local_enforcement_status_date | 2008-01-01 | Municipality had a State Building Code ordinance in effect on 2008-01-01 | Not applicable | `src:usa-mn:stat-326b-121` | verified_core |
| `date-rule:usa-mn:exterior-work-minimum-time` | Municipal exterior-work completion ordinances | minimum_duration_after_permit | 180 days after permit issuance | Local ordinance requiring exterior work completion after permit | Not applicable | `src:usa-mn:stat-326b-121` | verified_core |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building / Residential / Existing / Mechanical | next model-code cycle | null | null | null | null | null | monitor | `src:usa-mn:stat-326b-106`; `src:usa-mn:dli-state-building-codes-2020` | Statute requires review and adoption beginning with the 2018 model building codes and every six years thereafter, within two years of published editions. Final Minnesota effective dates for the next cycle were not resolved here. |
| Commercial Energy | each qualifying ASHRAE 90.1 or more efficient standard beginning 2024 | null | null | null | null | null | monitor | `src:usa-mn:stat-326b-106`; `src:usa-mn:dli-energy-conservation` | Statutory cycle for commercial energy code continues; current DLI effective date for 2024 Commercial Energy Code is 2024-01-05. |
| Residential Energy | future IECC or more efficient standards beginning 2026 | null | null | null | null | null | monitor | `src:usa-mn:stat-326b-106`; `src:usa-mn:dli-energy-conservation` | Statutory residential energy cycle begins in 2026; no final future effective date resolved here. |
| Electrical | next NEC cycle | null | null | null | null | null | monitor | `src:usa-mn:dli-electrical-codes`; `src:usa-mn:rule-1315` | Current effective NEC edition is 2023; future NEC cycle not resolved. |
| Fire | future IFC cycle | null | null | null | null | null | monitor | `src:usa-mn:rule-7511`; `src:usa-mn:stat-326b-02` | Fire-code update path should be monitored through Revisor rule chapter 7511 and DLI/SFM materials. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| `applicability-rule:usa-mn:public-buildings-state-licensed` | State Building Code | Public buildings and state-licensed facilities | Plans for construction, reconstruction, or alteration | Commissioner administers and enforces unless municipal agreement applies; approved plans required before work begins. | `src:usa-mn:stat-326b-107` | verified_core |
| `applicability-rule:usa-mn:municipal-option-1306-appendix-j` | Building / fire-protection systems / grading | Municipalities that adopted State Building Code | Municipality adopts optional provisions by ordinance without change | Chapter 1306 special fire-protection systems and IBC Appendix J grading are optional administrative provisions. | `src:usa-mn:rule-1300` | verified_core |
| `applicability-rule:usa-mn:geological-more-restrictive` | Building | Local ordinance more restrictive than State Building Code | Geological conditions warrant a more restrictive ordinance and state building official approval is obtained | Municipality may adopt more restrictive ordinance with approval; disapproval appeal available. | `src:usa-mn:stat-326b-121` | verified_core |
| `applicability-rule:usa-mn:fire-local-ordinance` | Fire - operational / prevention code | Local fire ordinances | Local government otherwise authorized by law to adopt/enforce fire ordinances | Local fire requirements may be equal, additional, or more stringent, subject to statutory limits and cannot exceed applicable State Building Code requirements. | `src:usa-mn:stat-299f-011` | verified_core |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** Minnesota adopts model codes and standards by reference into Minnesota Rules chapters, then amends them in state rule text. The primary construction-code administrative chapter is Minnesota Rules chapter 1300. Model-code content appears in family-specific chapters, including 1305, 1309, 1311, 1315, 1322, 1323, 1341, 1346, 4714, and 7511.

**Where amendments are published:** Minnesota Rules on the Office of the Revisor of Statutes website, supplemented by official DLI code pages and code books/fact sheets.

**Amendment parsing status:** partial_high_level_only

### 5.2 State Amendment Sources

| Amendment Source ID | Code Family | Publication Path | Coverage | Source IDs | Parsing Status |
| --- | --- | --- | --- | --- | --- |
| `amend-src:usa-mn:rule-1300` | Administration / statewide code composition | Minnesota Rules chapter 1300 | State Building Code composition, optional provisions, definitions, administration | `src:usa-mn:rule-1300` | high_level_parsed |
| `amend-src:usa-mn:rule-1305` | Building | Minnesota Rules chapter 1305 | 2018 IBC incorporation and Minnesota amendments | `src:usa-mn:rule-1305` | high_level_parsed |
| `amend-src:usa-mn:rule-1309` | Residential | Minnesota Rules chapter 1309 | 2018 IRC incorporation and Minnesota amendments | `src:usa-mn:rule-1309` | high_level_parsed |
| `amend-src:usa-mn:rule-1311` | Existing Building / Rehabilitation | Minnesota Rules chapter 1311 | 2018 IEBC incorporation and Minnesota amendments | `src:usa-mn:rule-1311` | high_level_parsed |
| `amend-src:usa-mn:rule-1346` | Mechanical / Fuel Gas | Minnesota Rules chapter 1346 | 2018 IMC and 2018 IFGC incorporation and Minnesota amendments | `src:usa-mn:rule-1346` | high_level_parsed |
| `amend-src:usa-mn:rule-1315` | Electrical | Minnesota Rules chapter 1315 | 2023 NEC incorporation and Minnesota amendments | `src:usa-mn:rule-1315` | high_level_parsed |
| `amend-src:usa-mn:rule-1322` | Residential Energy | Minnesota Rules chapter 1322 | 2012 IECC Residential Provisions and Minnesota amendments | `src:usa-mn:rule-1322` | high_level_parsed |
| `amend-src:usa-mn:rule-1323` | Commercial Energy | Minnesota Rules chapter 1323 | ASHRAE 90.1-2019 and Minnesota amendments | `src:usa-mn:rule-1323` | high_level_parsed |
| `amend-src:usa-mn:rule-1341` | Accessibility | Minnesota Rules chapter 1341 | IBC chapter 11, IEBC § 305, ICC/ANSI A117.1 amendments | `src:usa-mn:rule-1341` | high_level_parsed |
| `amend-src:usa-mn:rule-4714` | Plumbing | Minnesota Rules chapter 4714 | 2018 UPC and Minnesota amendments | `src:usa-mn:rule-4714` | high_level_parsed |
| `amend-src:usa-mn:rule-7511` | Fire | Minnesota Rules chapter 7511 | 2018 IFC and Minnesota amendments | `src:usa-mn:rule-7511` | high_level_parsed |
| `amend-src:usa-mn:rule-1307` | Elevator / Conveyance | Minnesota Rules chapter 1307 | ASME elevator/conveyance standards and Minnesota amendments | `src:usa-mn:rule-1307` | high_level_parsed |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Topic | Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| `amend:usa-mn:ibc-appendices-deleted` | Building | IBC appendices | Minnesota Rules chapter 1305 deletes all IBC appendices, with specific handling of grading and other referenced matters elsewhere. | `src:usa-mn:rule-1305`; `src:usa-mn:rule-1300` | verified_core |
| `amend:usa-mn:ibc-chapter-11-in-1341` | Building / Accessibility | Accessibility chapter location | Chapter 1305 incorporates amendments to IBC chapter 11 by reference, but actual amendments are located in chapter 1341. | `src:usa-mn:rule-1305`; `src:usa-mn:rule-1341` | verified_core |
| `amend:usa-mn:ibc-floodproofing-to-1335` | Building | Flood hazard / floodproofing | IBC flood hazard or floodproofing provisions and references are deleted; floodproofing requirements are located in Minnesota Rules chapter 1335. | `src:usa-mn:rule-1305`; `src:usa-mn:rule-1300` | verified_core |
| `amend:usa-mn:irc-energy-replaced` | Residential / Energy | Residential energy provisions | IRC chapter 11 energy provisions are replaced by Minnesota Rules chapters 1322 and 1323. | `src:usa-mn:rule-1309`; `src:usa-mn:rule-1322`; `src:usa-mn:rule-1323` | verified_core |
| `amend:usa-mn:optional-1306-appendix-j` | Building / Fire protection / Grading | Optional municipal adoption | Chapter 1306 special fire-protection systems and IBC Appendix J grading are optional provisions that municipalities may adopt without change. | `src:usa-mn:rule-1300` | verified_core |
| `amend:usa-mn:fire-local-stricter-limits` | Fire | Local fire ordinance limits | Local fire ordinances may be equal, additional, or more stringent, but must meet statutory conditions and may not exceed applicable State Building Code requirements. | `src:usa-mn:stat-299f-011` | verified_core |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-mn"
  model: "statewide_code_with_municipal_administration_and_state_reserved_roles"
  enforcing_entities:
    - "municipalities required to continue administration and enforcement because they had a State Building Code ordinance on 2008-01-01, except statutory small-population exception"
    - "municipalities that choose to administer and enforce the State Building Code by ordinance"
    - "Commissioner of Labor and Industry for public buildings and state-licensed facilities unless municipal agreement applies"
    - "Department of Labor and Industry for electrical installations, plumbing, boilers, high-pressure steam piping and appurtenances, and ammonia refrigeration piping under statutory direction"
    - "Commissioner / DLI for elevator provisions statewide except statutory municipal option"
    - "State Fire Marshal / Commissioner of Public Safety and municipal fire code officials for State Fire Code roles"
  required_officials:
    - "certified building official where a municipality administers the State Building Code"
    - "fire code official or state fire marshal / authorized representative for fire-code AHJ context"
  state_reserved_activities:
    - "public-building and state-licensed-facility plan review/enforcement unless municipal agreement applies"
    - "statewide elevator enforcement except statutory municipal option"
    - "DLI enforcement of listed trade-code provisions"
  source_ids:
    - "src:usa-mn:stat-326b-121"
    - "src:usa-mn:stat-326b-107"
    - "src:usa-mn:stat-326b-106"
    - "src:usa-mn:rule-1300"
    - "src:usa-mn:rule-7511-0202"
  verification_status: "verified_core"
  confidence: 0.84
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-mn"
  model: "state_preemption_with_limited_exceptions"
  applies_to_code_families:
    - "building"
    - "residential"
    - "existing_building"
    - "mechanical"
    - "fuel_gas"
    - "energy"
    - "accessibility"
    - "fire_operational_with_separate_statutory_local_fire_authority"
  approval_required: "yes_for_more_restrictive_geological_condition_building_ordinance"
  approving_authority_id: "ahj:usa-mn:state-building-official"
  filing_required: "unresolved"
  registry_exists: "unresolved"
  registry_source_ids: []
  legal_basis_source_ids:
    - "src:usa-mn:stat-326b-121"
    - "src:usa-mn:rule-1300"
    - "src:usa-mn:stat-299f-011"
  verification_status: "partially_verified"
  confidence: 0.76
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement and local amendment authority are separate in Minnesota. A municipality may be required or may choose to administer and enforce the State Building Code, but that does not generally allow the municipality to impose building-code provisions that differ from the State Building Code. Local fire ordinances have a separate statutory pathway allowing equal, additional, or more stringent requirements, subject to limits and a prohibition against exceeding applicable State Building Code requirements.

### 6.4 Known Local Amendment Registries

| Registry ID | Registry Name | Coverage | Source IDs | Status | Notes |
| --- | --- | --- | --- | --- | --- |
| `registry:usa-mn:local-building-amendments` | statewide local building amendment registry | unresolved | none | unresolved | No official statewide registry of approved geological-condition ordinances or optional provisions was verified here. |
| `registry:usa-mn:local-fire-ordinances` | local fire ordinances | unresolved | none | unresolved | Local fire ordinances may exist, but no statewide registry was identified in the captured sources. |

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendments were parsed. Any production system should treat local optional adoption of chapter 1306, IBC Appendix J grading, geological-condition ordinances, and local fire ordinances as jurisdiction-specific data requiring AHJ confirmation.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: partially_started_authority_only

Jurisdiction stack:

```text
Address
  -> State: Minnesota
  -> County
  -> Municipality / unincorporated county / town
  -> Determine whether the municipality administers and enforces the State Building Code by continuation rule or ordinance
  -> Building AHJ: certified municipal building official if local administration applies; state building official / DLI for state-administered public buildings and state-licensed facilities
  -> Fire AHJ: municipal fire code official where appointed, or state fire marshal / authorized representative where state role applies
  -> Trade-specific AHJs: DLI / applicable boards or programs for electrical, plumbing, elevator, boiler, high-pressure piping, and related state-reserved scopes
  -> Applicable Minnesota State Building Code adoption records
  -> Applicable local optional provisions, approved geological-condition ordinances, local fire ordinances, and AHJ interpretations
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

No AHJ contact data was populated. Production AHJ resolution still requires municipal code-administration lists, state building official/DLI contacts, State Fire Marshal district or inspection contacts, and local fire official data.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Type | Publisher | URL | Key Fields Supported | Caveats |
| --- | --- | --- | --- | --- | --- | --- |
| `src:usa-mn:dli-state-building-codes-2020` | 2020 Minnesota State Building Codes | agency_page | Minnesota Department of Labor and Industry | https://www.dli.mn.gov/business/codes-and-laws/2020-minnesota-state-building-codes | 2020 code effective dates; current State Building Code family list; Commercial Energy Code effective date | Official agency page; page combines code-family summaries and does not provide every adoption date. |
| `src:usa-mn:dli-building-code-overview` | Overview of Minnesota State Building Code | agency_page | Minnesota Department of Labor and Industry | https://www.dli.mn.gov/business/codes-and-laws/overview-minnesota-state-building-code | Statewide minimum-standard statement; local enforceability by municipal ordinance; adoption-cycle summary | Official agency overview; use statutes/rules for legal text. |
| `src:usa-mn:dli-ccld-overview` | Construction Codes and Licensing | agency_page | Minnesota Department of Labor and Industry | https://www.dli.mn.gov/about-department/about-dli/our-areas-service/construction-codes-and-licensing | DLI/CCLD program scope and administered code families | Official agency page; high-level program description. |
| `src:usa-mn:stat-326b-02` | Minnesota Statutes § 326B.02, Powers | statute | Office of the Revisor of Statutes | https://www.revisor.mn.gov/statutes/cite/326B.02 | Transfers of authority; general rulemaking authority; State Fire Code rulemaking authority; Board of Electricity and Plumbing Board transfer context | Official HTML statute; current version should be checked for session changes. |
| `src:usa-mn:stat-326b-101` | Minnesota Statutes § 326B.101, State Building Code; Application | statute | Office of the Revisor of Statutes | https://www.revisor.mn.gov/statutes/cite/326B.101 | State Building Code application; commissioner administration and amendment of state construction code | Official HTML statute; current version should be checked for session changes. |
| `src:usa-mn:stat-326b-106` | Minnesota Statutes § 326B.106, General Powers of Commissioner of Labor and Industry | statute | Office of the Revisor of Statutes | https://www.revisor.mn.gov/statutes/cite/326B.106 | Commissioner authority; model-code adoption cycle; energy-code cycles; elevator, trade, and fire-code enforcement relationships | Official HTML statute; current version should be checked for session changes. |
| `src:usa-mn:stat-326b-107` | Minnesota Statutes § 326B.107, Public Buildings and State-Licensed Facilities | statute | Office of the Revisor of Statutes | https://www.revisor.mn.gov/statutes/cite/326B.107 | State administration/enforcement for public buildings and state-licensed facilities; municipal agreements | Official HTML statute; current version should be checked for session changes. |
| `src:usa-mn:stat-326b-121` | Minnesota Statutes § 326B.121, State Building Code; Municipal Enforcement | statute | Office of the Revisor of Statutes | https://www.revisor.mn.gov/statutes/cite/326B.121 | Statewide standard; municipal enforcement; local amendment/preemption; geological-condition exception; exterior-work timing | Official HTML statute; current version should be checked for session changes. |
| `src:usa-mn:rule-1300` | Minnesota Rules chapter 1300, State Building Code Administration | regulation | Office of the Revisor of Statutes | https://www.revisor.mn.gov/rules/1300/full | Code composition; optional provisions; definitions; administration | Official HTML rule; detailed administrative provisions not exhaustively parsed. |
| `src:usa-mn:rule-1305` | Minnesota Rules chapter 1305, Minnesota Building Code | regulation | Office of the Revisor of Statutes | https://www.revisor.mn.gov/rules/1305/full | 2018 IBC incorporation; mandatory chapters; major amendments and cross-references | Official HTML rule; detailed amendments not exhaustively parsed. |
| `src:usa-mn:rule-1309` | Minnesota Rules chapter 1309, Minnesota Residential Code | regulation | Office of the Revisor of Statutes | https://www.revisor.mn.gov/rules/1309/full | 2018 IRC incorporation; mandatory chapters; replacement of energy chapter | Official HTML rule; detailed amendments not exhaustively parsed. |
| `src:usa-mn:rule-1311` | Minnesota Rules chapter 1311, Minnesota Conservation Code for Existing Buildings | regulation | Office of the Revisor of Statutes | https://www.revisor.mn.gov/rules/1311/full | 2018 IEBC incorporation; existing-building scope | Official HTML rule; detailed amendments not exhaustively parsed. |
| `src:usa-mn:rule-1346` | Minnesota Rules chapter 1346, Minnesota Mechanical and Fuel Gas Code | regulation | Office of the Revisor of Statutes | https://www.revisor.mn.gov/rules/1346/full | 2018 IMC and 2018 IFGC incorporation; mechanical/fuel-gas references | Official HTML rule; detailed amendments not exhaustively parsed. |
| `src:usa-mn:rule-1315` | Minnesota Rules chapter 1315, Minnesota Electrical Code | regulation | Office of the Revisor of Statutes | https://www.revisor.mn.gov/rules/1315/full | 2023 NEC incorporation and electrical-code rules | Official HTML rule; date transition confirmed through DLI electrical page. |
| `src:usa-mn:dli-electrical-codes` | Electrical codes and standards | agency_page | Minnesota Department of Labor and Industry | https://www.dli.mn.gov/business/electrical-contractors/electrical-codes-and-standards | 2023 NEC effective date and permit-date transition | Official agency page; should be monitored for future NEC cycles. |
| `src:usa-mn:rule-4714` | Minnesota Rules chapter 4714, Minnesota Plumbing Code | regulation | Office of the Revisor of Statutes | https://www.revisor.mn.gov/rules/4714/full | 2018 UPC incorporation; plumbing administrative authority definitions | Official HTML rule; effective date not resolved from captured text. |
| `src:usa-mn:dli-energy-conservation` | Energy conservation | agency_page | Minnesota Department of Labor and Industry | https://www.dli.mn.gov/business/codes-and-laws/energy-conservation | Residential and Commercial Energy Code effective dates | Official agency page; use rules for incorporated standards. |
| `src:usa-mn:rule-1322` | Minnesota Rules chapter 1322, Residential Energy Code | regulation | Office of the Revisor of Statutes | https://www.revisor.mn.gov/rules/1322/full | 2012 IECC Residential Provisions incorporation; residential energy chapter | Official HTML rule; detailed amendments not exhaustively parsed. |
| `src:usa-mn:rule-1323` | Minnesota Rules chapter 1323, Commercial Energy Code | regulation | Office of the Revisor of Statutes | https://www.revisor.mn.gov/rules/1323/full | ASHRAE 90.1-2019 incorporation; commercial energy chapter | Official HTML rule; detailed amendments not exhaustively parsed. |
| `src:usa-mn:rule-7511` | Minnesota Rules chapter 7511, Minnesota State Fire Code | regulation | Office of the Revisor of Statutes | https://www.revisor.mn.gov/rules/7511/full | 2018 IFC incorporation as Minnesota State Fire Code | Official HTML rule; effective date not resolved from captured text. |
| `src:usa-mn:stat-299f-011` | Minnesota Statutes § 299F.011, State Fire Code | statute | Office of the Revisor of Statutes | https://www.revisor.mn.gov/statutes/cite/299F.011 | State Fire Code applicability; local fire ordinance authority and limits | Official HTML statute; current version should be checked for session changes. |
| `src:usa-mn:stat-299f-01` | Minnesota Statutes § 299F.01, State Fire Marshal | statute | Office of the Revisor of Statutes | https://www.revisor.mn.gov/statutes/cite/299F.01 | State Fire Marshal Division within Department of Public Safety | Official HTML statute; current version should be checked for session changes. |
| `src:usa-mn:rule-7511-0202` | Minnesota Rules part 7511.0202, Definitions / AHJ | regulation | Office of the Revisor of Statutes | https://www.revisor.mn.gov/rules/7511.0202/ | Fire-code official and authority-having-jurisdiction definitions | Official HTML rule; use in AHJ logic with local confirmation. |
| `src:usa-mn:rule-1341` | Minnesota Rules chapter 1341, Accessibility for Buildings and Facilities | regulation | Office of the Revisor of Statutes | https://www.revisor.mn.gov/rules/1341/full | IBC chapter 11, IEBC section 305, ICC/ANSI A117.1 incorporation and amendments | Official HTML rule; detailed amendments not exhaustively parsed. |
| `src:usa-mn:rule-1307` | Minnesota Rules chapter 1307, Elevators and Related Devices | regulation | Office of the Revisor of Statutes | https://www.revisor.mn.gov/rules/1307/full | Elevator/conveyance code scope and referenced ASME standards | Official HTML rule; effective date not resolved from captured text. |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| `src:usa-mn:dli-state-building-codes-2020` | agency_summary | Official DLI page states effective dates and lists codes, but does not expose every rule-adoption date field. | use_for_effective_dates_and_code_list; cross-check rules for model-code incorporation |
| `src:usa-mn:rule-7511` | effective_date_gap | Rule text confirms IFC incorporation but captured text did not resolve the exact fire-code effective date. | keep fire effective/adoption dates null until verified |
| `src:usa-mn:rule-4714` | effective_date_gap | Rule text confirms 2018 UPC incorporation but captured text did not resolve the exact effective date. | keep plumbing effective/adoption dates null until verified |
| `src:usa-mn:rule-1307` | mixed_standards | Elevator chapter incorporates several ASME standards rather than a single model-code edition. | model as mixed referenced standards |
| `src:usa-mn:stat-326b-121` | legal_interpretation_needed | Local preemption and exceptions have legal nuance, especially geological-condition approvals and maintenance ordinances. | production legal review recommended for local-amendment resolver |
| `src:usa-mn:stat-299f-011` | legal_interpretation_needed | Local fire ordinance authority is separate from building-code amendment authority and contains guardrails. | model separately from building-code amendments |

### 8.3 Supplemental Sources

None used for populated facts. The report relies on official Minnesota agency, statute, and rule sources.

### 8.4 Source Extraction Metadata

| Extraction ID | Source ID | Extracted On | Extractor | Coverage | Notes |
| --- | --- | --- | --- | --- | --- |
| `extract:usa-mn:2026-06-25:001` | `src:usa-mn:dli-state-building-codes-2020` | 2026-06-25 | ChatGPT | effective dates and code-family list | Official DLI page reviewed. |
| `extract:usa-mn:2026-06-25:002` | `src:usa-mn:dli-building-code-overview` | 2026-06-25 | ChatGPT | statewide status and local enforceability summary | Official DLI page reviewed. |
| `extract:usa-mn:2026-06-25:003` | `src:usa-mn:stat-326b-02` | 2026-06-25 | ChatGPT | transfer/rulemaking/fire-code authority | Revisor statute reviewed. |
| `extract:usa-mn:2026-06-25:004` | `src:usa-mn:stat-326b-106` | 2026-06-25 | ChatGPT | commissioner powers; model-code cycle; state enforcement roles | Revisor statute reviewed. |
| `extract:usa-mn:2026-06-25:005` | `src:usa-mn:stat-326b-121` | 2026-06-25 | ChatGPT | municipal enforcement and local building-code amendment limits | Revisor statute reviewed. |
| `extract:usa-mn:2026-06-25:006` | `src:usa-mn:rule-1300` | 2026-06-25 | ChatGPT | State Building Code composition and optional provisions | Revisor rule reviewed. |
| `extract:usa-mn:2026-06-25:007` | `src:usa-mn:rule-1305`; `src:usa-mn:rule-1309`; `src:usa-mn:rule-1311`; `src:usa-mn:rule-1315`; `src:usa-mn:rule-1322`; `src:usa-mn:rule-1323`; `src:usa-mn:rule-1341`; `src:usa-mn:rule-1346`; `src:usa-mn:rule-4714`; `src:usa-mn:rule-7511`; `src:usa-mn:rule-1307` | 2026-06-25 | ChatGPT | current model-code incorporation chapters | Represents extraction across major current adoption rule chapters. |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| `report` | `report.status` | partially_verified | verified | 1.00 | none | Core authority and adoption matrix are source-backed; unresolved fields remain explicit. |
| `state` | `state.state_id` | US-MN | verified | 1.00 | none | Frontmatter normalized. |
| `ahj:usa-mn:dli-commissioner` | primary_authority | Minnesota Commissioner of Labor and Industry / DLI CCLD | verified_core | 0.88 | `src:usa-mn:stat-326b-101`; `src:usa-mn:stat-326b-106`; `src:usa-mn:dli-ccld-overview` | Primary statewide construction-code authority. |
| `adopt:usa-mn:building:2020` | base_model_code | 2018 IBC | verified_core | 0.86 | `src:usa-mn:rule-1305` | Rule chapter 1305 incorporation. |
| `adopt:usa-mn:building:2020` | effective_date | 2020-03-31 | verified_core | 0.84 | `src:usa-mn:dli-state-building-codes-2020` | Agency effective-date page. |
| `adopt:usa-mn:residential:2020` | base_model_code | 2018 IRC | verified_core | 0.84 | `src:usa-mn:rule-1309` | Rule chapter 1309 incorporation. |
| `adopt:usa-mn:mechanical:2020` | base_model_code | 2018 IMC | verified_core | 0.82 | `src:usa-mn:rule-1346` | Rule chapter 1346 incorporation. |
| `adopt:usa-mn:fuel-gas:2020` | base_model_code | 2018 IFGC | verified_core | 0.82 | `src:usa-mn:rule-1346` | Rule chapter 1346 incorporation. |
| `adopt:usa-mn:electrical:2023` | effective_date | 2023-07-01 | verified_core | 0.90 | `src:usa-mn:dli-electrical-codes`; `src:usa-mn:rule-1315` | DLI page confirms Board adoption and permit-date transition. |
| `adopt:usa-mn:energy-commercial:2024` | effective_date | 2024-01-05 | verified_core | 0.88 | `src:usa-mn:dli-energy-conservation`; `src:usa-mn:rule-1323` | DLI page and Revisor rule support current code. |
| `adopt:usa-mn:energy-residential:2015` | effective_date | 2015-02-14 | verified_core | 0.84 | `src:usa-mn:dli-energy-conservation`; `src:usa-mn:rule-1322` | DLI page and Revisor rule support current code. |
| `adopt:usa-mn:fire:2018-ifc` | base_model_code | 2018 IFC | verified_core | 0.78 | `src:usa-mn:rule-7511` | Effective date remains unresolved. |
| `local-enforcement:usa-mn` | model | statewide_code_with_municipal_administration_and_state_reserved_roles | verified_core | 0.84 | `src:usa-mn:stat-326b-121`; `src:usa-mn:stat-326b-107`; `src:usa-mn:stat-326b-106` | Local enforcement separated from amendment authority. |
| `local-amendment-rule:usa-mn` | model | state_preemption_with_limited_exceptions | partially_verified | 0.76 | `src:usa-mn:stat-326b-121`; `src:usa-mn:rule-1300`; `src:usa-mn:stat-299f-011` | Legal review recommended before production automation. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Every `src:usa-mn:*` identifier cited outside section 8 is present in the source registry. |
| All authority IDs resolve | pass | Authority IDs used in authority tables and edge records are defined or intentionally scoped in section 2. |
| All current code families have adoption records | pass | Matrix rows are explicit and normalized adoption records are provided; some date fields remain null with notes. |
| Building and operational fire code are separated | pass | Fire construction references and operational/prevention fire code are separate rows and authority records. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Separate columns are maintained; unresolved dates remain null. |
| Effective dates are valid ISO dates | pass | Entered effective dates use ISO format. |
| No impossible date sequences | pass | No conflicting date sequences identified. |
| Transition rules have explicit trigger conditions | pass | Electrical permit-date transition, state code effective dates, and local ordinance triggers are recorded. |
| Permit-date logic is captured where applicable | pass | Electrical code permit-date transition is captured. |
| Local enforcement model classified | pass | Local enforcement model classified with source-backed caveats. |
| Local amendment rule classified | pass | Building preemption and fire local ordinance posture captured; registry remains unresolved. |
| AHJ confirmation metadata present | fail | AHJ contact and boundary data have not been populated. |
| Official-source caveats captured | pass | Caveats are listed in section 8.2. |
| Remaining unresolved fields explicit | pass | Unresolved dates and registry gaps are labeled directly. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `issue:usa-mn:001` | high | fire code effective date | Minnesota Rules chapter 7511 confirms 2018 IFC incorporation, but this pass did not resolve the exact effective date from a primary effective-date source. | Verify State Register notice, rule history, or official SFM/DLI publication for effective date and adoption date. | null | null | open |
| `issue:usa-mn:002` | high | local amendment registry | No official statewide registry was captured for geological-condition approvals, optional chapter 1306/Appendix J adoption, or local fire ordinances. | Identify DLI/state building official records and any SFM/local ordinance filing practices. | null | null | open |
| `issue:usa-mn:003` | medium | plumbing effective date | Chapter 4714 incorporation of 2018 UPC is verified, but adoption/effective/operative dates were not resolved. | Verify rule history and official DLI plumbing-code effective-date materials. | null | null | open |
| `issue:usa-mn:004` | medium | AHJ boundary/contact data | Report identifies AHJ model but lacks municipal code-administration lists, local fire official data, and boundary sources. | Add official municipal administration list, GIS/boundary data, and AHJ contacts. | null | null | open |
| `issue:usa-mn:005` | medium | future code cycle | Statutory review/adoption cycles are captured, but final next-cycle effective dates are unresolved. | Monitor DLI code adoption pages, State Register, and Revisor rule updates. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| `watch:usa-mn:dli-2020-codes` | `src:usa-mn:dli-state-building-codes-2020` | html_diff | monthly | Code list, effective dates, or adoption update changes | 2026-06-25 | active |
| `watch:usa-mn:dli-overview` | `src:usa-mn:dli-building-code-overview` | html_diff | quarterly | Changes to statewide/minimum-standard or local enforceability language | 2026-06-25 | active |
| `watch:usa-mn:admin-code` | `src:usa-mn:rule-1300` | html_diff | monthly | Chapter 1300 amendment or optional provision change | 2026-06-25 | active |
| `watch:usa-mn:building-code` | `src:usa-mn:rule-1305` | html_diff | monthly | Building code chapter amendment | 2026-06-25 | active |
| `watch:usa-mn:residential-code` | `src:usa-mn:rule-1309` | html_diff | monthly | Residential code chapter amendment | 2026-06-25 | active |
| `watch:usa-mn:electrical` | `src:usa-mn:dli-electrical-codes` | html_diff | monthly | NEC cycle or permit-date transition update | 2026-06-25 | active |
| `watch:usa-mn:energy` | `src:usa-mn:dli-energy-conservation` | html_diff | monthly | Energy code effective date, ASHRAE, or IECC update | 2026-06-25 | active |
| `watch:usa-mn:fire-code` | `src:usa-mn:rule-7511` | html_diff | monthly | State Fire Code chapter amendment or new IFC incorporation | 2026-06-25 | active |
| `watch:usa-mn:fire-statute` | `src:usa-mn:stat-299f-011` | statute_diff | quarterly | Local fire ordinance authority or State Fire Code applicability changes | 2026-06-25 | active |
| `watch:usa-mn:municipal-enforcement` | `src:usa-mn:stat-326b-121` | statute_diff | quarterly | Local enforcement or local amendment rule changes | 2026-06-25 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-25 | Populated Minnesota report from baseline draft with official source registry, authority model, adoption matrix, date rules, local enforcement/amendment model, QA, and open issues. | `report:usa-mn`; `ahj:usa-mn:dli-commissioner`; `local-enforcement:usa-mn`; `local-amendment-rule:usa-mn`; `adopt:usa-mn:*` | `src:usa-mn:*` | ChatGPT | Status upgraded to `partially_verified` because core authority and current adoption facts are source-backed; unresolved items remain explicit. |
