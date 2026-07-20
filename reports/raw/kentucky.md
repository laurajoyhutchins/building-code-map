---
state:
  state_id: "US-KY"
  name: "Kentucky"
  abbreviation: "KY"
report:
  report_id: "state-report:usa-ky"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-25"
  last_verified: null
  reviewed_by: null
risk:
  overall_confidence: 0.68 # 0.00 - 1.00
  risk_flags:
    - "core_building_and_residential_authority_verified_from_official_sources"
    - "trade_code_editions_partially_normalized"
    - "fire_prevention_local_program_scope_requires_followup"
    - "ahj_boundary_and_contact_data_not_populated"
  open_questions_count: 4

---

# State Building Code Authority Report: Kentucky

## 1. Executive Summary

- **Authority model:** Kentucky uses a statewide, mandatory, uniform building-code model administered by the Kentucky Department of Housing, Buildings and Construction (DHBC). The primary construction-code regulations are 815 KAR 7:120 for the Kentucky Building Code (KBC) and 815 KAR 7:125 for the Kentucky Residential Code (KRC). Both regulations are based on KRS 198B.040(7) and KRS 198B.050. Source IDs: `src:usa-ky:815-kar-7-120`, `src:usa-ky:815-kar-7-125`, `src:usa-ky:dhbc-building-code-enforcement`.

- **Statewide code status:** The current KBC record verified in this report incorporates the 2015 International Building Code and the 2018 Kentucky Building Code, Fourth Edition, February 2024. The current KRC record verified in this report incorporates the 2015 International Residential Code for One- and Two-Family Dwellings and the 2018 Kentucky Residential Code, Third Edition, August 2024. Source IDs: `src:usa-ky:815-kar-7-120`, `src:usa-ky:kbc-2018-fourth-edition-pdf`, `src:usa-ky:815-kar-7-125`, `src:usa-ky:krc-2018-third-edition-pdf`.

- **Local enforcement model:** Enforcement is shared between DHBC and local city or county building departments. Local governments may operate building-code programs and may request expanded local jurisdiction, but DHBC retains state-reserved plan-review, inspection, and enforcement responsibilities and may preempt deficient expanded-jurisdiction programs. Single-family permits, inspections, and certificates of occupancy are not required unless a local ordinance requires them, although the KRC remains the applicable construction standard. Source IDs: `src:usa-ky:krs-198b-060`, `src:usa-ky:815-kar-7-110`, `src:usa-ky:dhbc-building-codes`, `src:usa-ky:815-kar-7-125`.

- **Local amendment posture:** The KBC and KRC are described by DHBC as “mini/maxi” codes. Local governments are not to adopt or enforce another building code governing commercial construction under the KBC, and local governments are not to adopt or enforce another building code for detached single-family dwellings, two-family dwellings, and townhouses under the KRC. The verified amendment path is through DHBC, the regulatory process, and statewide-adopted amendments. Fire-prevention local ordinance authority under 815 KAR 10:060 and KRS 227.320 is verified only at a summary level here and remains a follow-up item for local fire-code scope. Source IDs: `src:usa-ky:kbc-2018-fourth-edition-pdf`, `src:usa-ky:krc-2018-third-edition-pdf`, `src:usa-ky:krs-198b-080`, `src:usa-ky:815-kar-10-060`.

- **Known transition periods or pending changes:** The KBC states that plans submitted on or after 2019-08-01 must be designed and submitted to conform to the code. The KRC states that, effective 2019-08-01, the code is mandatory and no permit may be issued for construction under another building code. KRC electrical notes defer selected 2023 NEC articles until 2026-07-15 and defer certain GFCI requirements until specified UL standards are revised. DHBC's proposed-regulation page lists 2026 proposed regulations for elevator, boiler, manufactured-housing, and inspector/electrical-inspector rules, but no pending KBC/KRC model-code edition change was normalized in this report. Source IDs: `src:usa-ky:kbc-2018-fourth-edition-pdf`, `src:usa-ky:krc-2018-third-edition-pdf`, `src:usa-ky:dhbc-proposed-kar`.

- **Production readiness:** narrow_validation_ready

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: DHBC is the primary statewide authority for the mandatory uniform building
    code under 815 KAR 7:120 and 815 KAR 7:125.
  confidence: 0.86
  source_ids:
  - src:usa-ky:815-kar-7-120
  - src:usa-ky:815-kar-7-125
  - src:usa-ky:dhbc-building-code-enforcement
- topic: Primary building code edition
  finding: KBC = 2015 IBC plus Kentucky amendments in the 2018 Kentucky Building Code,
    Fourth Edition, February 2024.
  confidence: 0.88
  source_ids:
  - src:usa-ky:815-kar-7-120
  - src:usa-ky:kbc-2018-fourth-edition-pdf
- topic: Primary residential code edition
  finding: KRC = 2015 IRC plus Kentucky amendments in the 2018 Kentucky Residential
    Code, Third Edition, August 2024.
  confidence: 0.88
  source_ids:
  - src:usa-ky:815-kar-7-125
  - src:usa-ky:krc-2018-third-edition-pdf
- topic: Electrical code authority
  finding: Electrical standards are part of the Uniform State Building Code and use
    NFPA 70 / NEC as a minimum standard; KBC/KRC current references identify NFPA
    70-23 with Kentucky transition notes.
  confidence: 0.74
  source_ids:
  - src:usa-ky:krs-227-480
  - src:usa-ky:kbc-2018-fourth-edition-pdf
  - src:usa-ky:krc-2018-third-edition-pdf
  - src:usa-ky:815-kar-35-020
- topic: Fire code authority
  finding: Construction fire-safety requirements are in the KBC where referenced;
    operational and existing-building fire safety is separated through the Kentucky
    Standards of Safety, 815 KAR 10:060, with State Fire Marshal and local fire program
    roles.
  confidence: 0.72
  source_ids:
  - src:usa-ky:kbc-2018-fourth-edition-pdf
  - src:usa-ky:815-kar-10-060
- topic: Local enforcement
  finding: Enforcement is shared between DHBC and city/county building departments,
    with expanded local jurisdiction criteria and state-retained responsibilities.
  confidence: 0.79
  source_ids:
  - src:usa-ky:krs-198b-060
  - src:usa-ky:815-kar-7-110
  - src:usa-ky:dhbc-building-codes
- topic: Local amendments
  finding: KBC/KRC local building-code substitutions are not supported; amendments
    verified here are statewide through DHBC regulatory process.
  confidence: 0.76
  source_ids:
  - src:usa-ky:kbc-2018-fourth-edition-pdf
  - src:usa-ky:krc-2018-third-edition-pdf
  - src:usa-ky:krs-198b-080
- topic: Effective / operative date rule
  finding: KBC plan-submittal and KRC permit rules both use 2019-08-01 as the verified
    mandatory/operative transition date for the 2018 code family; current KAR amendments
    have a separate 2024-12-03 effective history entry.
  confidence: 0.74
  source_ids:
  - src:usa-ky:kbc-2018-fourth-edition-pdf
  - src:usa-ky:krc-2018-third-edition-pdf
  - src:usa-ky:815-kar-7-120
  - src:usa-ky:815-kar-7-125
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | `ahj:usa-ky:dhbc` |
| Authority name | Kentucky Department of Housing, Buildings and Construction |
| Authority type | state_department |
| Legal basis | KRS 198B.040(7); KRS 198B.050; 815 KAR 7:120; 815 KAR 7:125 |
| Role | Promulgates, administers, and enforces the mandatory uniform state building code and the Kentucky amendments to the model codes. |
| Enforcement model | Statewide uniform code with state and local enforcement responsibilities. DHBC retains plan-review, inspection, and enforcement responsibilities for specified buildings and may preempt deficient expanded local programs. |
| Source IDs | `src:usa-ky:815-kar-7-120`; `src:usa-ky:815-kar-7-125`; `src:usa-ky:dhbc-building-code-enforcement`; `src:usa-ky:815-kar-7-110` |
| Verification status | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | `ahj:usa-ky:dhbc` | Kentucky Department of Housing, Buildings and Construction | Adopts and administers the KBC for buildings other than KRC-covered dwellings/townhouses and other listed exceptions. | KRS 198B.040(7); KRS 198B.050; 815 KAR 7:120 | `src:usa-ky:815-kar-7-120`; `src:usa-ky:kbc-2018-fourth-edition-pdf` | partially_verified |
| Residential | `ahj:usa-ky:dhbc` | Kentucky Department of Housing, Buildings and Construction | Adopts and administers the KRC for single-family dwellings, two-family dwellings, and townhouses. | KRS 198B.040(7); KRS 198B.050; 815 KAR 7:125 | `src:usa-ky:815-kar-7-125`; `src:usa-ky:krc-2018-third-edition-pdf` | partially_verified |
| Existing Building / Rehabilitation | `ahj:usa-ky:dhbc` | Kentucky Department of Housing, Buildings and Construction | Applies KBC Chapter 34 and the International Existing Building Code to repair, alteration, change of occupancy, additions, and relocation to the extent referenced. | 815 KAR 7:120; KBC Sections 101.4.7, 102.6, and Chapter 34 | `src:usa-ky:kbc-2018-fourth-edition-pdf` | partially_verified |
| Mechanical | `ahj:usa-ky:dhbc-hvac` | Kentucky Department of Housing, Buildings and Construction, Division of HVAC | Administers HVAC licensing and state/local HVAC permitting and inspection programs; KBC references the International Mechanical Code for mechanical systems. | KRS 198B.650-198B.689; 815 KAR 8:100; KBC Section 101.4.2 | `src:usa-ky:dhbc-hvac`; `src:usa-ky:815-kar-8-100`; `src:usa-ky:kbc-2018-fourth-edition-pdf` | partially_verified |
| Plumbing | `ahj:usa-ky:dhbc-plumbing` | Kentucky Department of Housing, Buildings and Construction, Division of Plumbing | Administers the Kentucky State Plumbing Code and plumbing permitting/licensing framework. | KRS Chapter 318; KRS 318.130; 815 KAR Chapter 20 | `src:usa-ky:dhbc-plumbing-code`; `src:usa-ky:dhbc-plumbing-division`; `src:usa-ky:815-kar-20-130`; `src:usa-ky:kbc-2018-fourth-edition-pdf` | partially_verified |
| Fuel Gas | `ahj:usa-ky:dhbc` | Kentucky Department of Housing, Buildings and Construction | Applies NFPA 54 by reference for gas piping, gas appliances, and related accessories as covered by KBC/KRC. | 815 KAR 7:120; 815 KAR 7:125; KBC Section 101.4.1; KRC referenced standards | `src:usa-ky:kbc-2018-fourth-edition-pdf`; `src:usa-ky:krc-2018-third-edition-pdf` | partially_verified |
| Electrical | `ahj:usa-ky:dhbc-electrical` | Kentucky Department of Housing, Buildings and Construction / certified electrical inspectors | Administers electrical inspection requirements; KRS 227.480 makes Uniform State Building Code electrical standards based on NEC minimum standards. | KRS 227.480; 815 KAR 35:020; 815 KAR 7:120; 815 KAR 7:125 | `src:usa-ky:krs-227-480`; `src:usa-ky:815-kar-35-020`; `src:usa-ky:kbc-2018-fourth-edition-pdf`; `src:usa-ky:krc-2018-third-edition-pdf` | partially_verified |
| Energy | `ahj:usa-ky:dhbc` | Kentucky Department of Housing, Buildings and Construction | Applies IECC-based energy provisions through KBC Chapter 13 and KRC Chapter 11. | 815 KAR 7:120; 815 KAR 7:125; KBC Section 101.4.6; KBC Chapter 13; KRC Chapter 11 | `src:usa-ky:kbc-2018-fourth-edition-pdf`; `src:usa-ky:krc-2018-third-edition-pdf` | partially_verified |
| Fire - construction references | `ahj:usa-ky:dhbc` | Kentucky Department of Housing, Buildings and Construction | Applies International Fire Code provisions only where specifically referenced in the KBC for new construction. | 815 KAR 7:120; KBC Section 101.4.5 | `src:usa-ky:kbc-2018-fourth-edition-pdf` | partially_verified |
| Fire - operational / prevention code | `ahj:usa-ky:state-fire-marshal` | State Fire Marshal / local fire code official | Enforces Kentucky Standards of Safety for existing buildings and continued fire-safety maintenance; State Fire Marshal has primary jurisdiction unless a local fire-inspection program is established by ordinance, with exclusive jurisdiction over specified state and licensed facilities. | KRS 227.300; KRS 227.320; 815 KAR 10:060; KBC Section 101.5 | `src:usa-ky:815-kar-10-060`; `src:usa-ky:kbc-2018-fourth-edition-pdf` | partially_verified |
| Accessibility | `ahj:usa-ky:dhbc` | Kentucky Department of Housing, Buildings and Construction | Applies KBC accessibility provisions as part of building-code plan review and inspection. | 815 KAR 7:120; KBC Chapter 11 | `src:usa-ky:kbc-2018-fourth-edition-pdf`; `src:usa-ky:dhbc-building-code-enforcement` | partially_verified |
| Elevator / Conveyance | `ahj:usa-ky:dhbc-elevator` | Kentucky Department of Housing, Buildings and Construction, Elevator Inspection Section | Approves plans/specifications, issues permits, and administers inspection requirements and fees for elevators and related conveyances. | KRS 198B.400-198B.540; 815 KAR 4:010; 815 KAR 4:025 | `src:usa-ky:dhbc-elevators`; `src:usa-ky:815-kar-4-010`; `src:usa-ky:815-kar-4-025` | partially_verified |

### 2.3 Authority Hierarchy Notes

Kentucky's verified hierarchy is state-led and uniform. DHBC promulgates and administers the KBC/KRC and sets statewide amendments. City and county building departments may enforce within their jurisdictions and may obtain expanded jurisdiction, but the department retains state plan-review, inspection, and enforcement responsibility for specified categories and may preempt local expanded jurisdiction in whole or in part. Fire prevention is a parallel but connected model: the Kentucky Standards of Safety supplement the KBC in fire-safety matters, the State Fire Marshal has primary fire-safety jurisdiction by default, and local fire chiefs may have primary local fire-inspection jurisdiction where a local program is established by ordinance.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| `edge:usa-ky:001` | `ahj:usa-ky:dhbc` | promulgates | Kentucky Building Code and Kentucky Residential Code | `src:usa-ky:815-kar-7-120`; `src:usa-ky:815-kar-7-125` | partially_verified |
| `edge:usa-ky:002` | `ahj:usa-ky:dhbc` | incorporates_by_reference | 2015 IBC plus 2018 KBC, Fourth Edition, February 2024 | `src:usa-ky:815-kar-7-120`; `src:usa-ky:kbc-2018-fourth-edition-pdf` | partially_verified |
| `edge:usa-ky:003` | `ahj:usa-ky:dhbc` | incorporates_by_reference | 2015 IRC plus 2018 KRC, Third Edition, August 2024 | `src:usa-ky:815-kar-7-125`; `src:usa-ky:krc-2018-third-edition-pdf` | partially_verified |
| `edge:usa-ky:004` | `ahj:usa-ky:dhbc` | shares_enforcement_with | city and county building departments | `src:usa-ky:krs-198b-060`; `src:usa-ky:dhbc-building-codes` | partially_verified |
| `edge:usa-ky:005` | `ahj:usa-ky:dhbc` | grants_expanded_jurisdiction_to | local governments meeting 815 KAR 7:110 criteria | `src:usa-ky:815-kar-7-110` | partially_verified |
| `edge:usa-ky:006` | `ahj:usa-ky:dhbc` | may_preempt | deficient local expanded-jurisdiction program | `src:usa-ky:815-kar-7-110` | partially_verified |
| `edge:usa-ky:007` | `ahj:usa-ky:state-fire-marshal` | has_primary_jurisdiction_unless | local fire inspection program established by ordinance | `src:usa-ky:815-kar-10-060` | partially_verified |
| `edge:usa-ky:008` | `ahj:usa-ky:dhbc` | preempts_local_building_code_substitution | local governments adopting/enforcing another building code for KBC/KRC covered scopes | `src:usa-ky:kbc-2018-fourth-edition-pdf`; `src:usa-ky:krc-2018-third-edition-pdf`; `src:usa-ky:krs-198b-080` | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | Kentucky Building Code | 2015 International Building Code plus Kentucky amendments | 2018 Kentucky Building Code, Fourth Edition, February 2024 | mandatory statewide for KBC-covered buildings | null | 2024-12-03 | 2019-08-01 | 2019-08-01 | KBC Section 120.1: all plans submitted on or after 2019-08-01 must conform to the code. Effective date shown is the current 2024 amendment history date, not the original 2018 code-family conversion date. | `src:usa-ky:815-kar-7-120`; `src:usa-ky:kbc-2018-fourth-edition-pdf`; `src:usa-ky:dhbc-building-code-enforcement` |
| Residential | Kentucky Residential Code | 2015 International Residential Code for One- and Two-Family Dwellings plus Kentucky amendments | 2018 Kentucky Residential Code, Third Edition, August 2024 | mandatory statewide for single-family dwellings, two-family dwellings, and townhouses | null | 2024-12-03 | 2019-08-01 | 2019-08-01 | KRC Section R116.1: effective 2019-08-01, code is mandatory and no permit may be issued under another building code. Effective date shown is the current 2024 amendment history date. | `src:usa-ky:815-kar-7-125`; `src:usa-ky:krc-2018-third-edition-pdf`; `src:usa-ky:dhbc-building-code-enforcement` |
| Existing Building / Rehabilitation | Kentucky Building Code existing-building provisions | International Existing Building Code by KBC reference | edition not normalized; KBC Chapter 34 governs scope | mandatory where KBC applies to repair, alteration, change of occupancy, additions, or relocation | null | 2024-12-03 | 2019-08-01 | 2019-08-01 | KBC Sections 101.4.7 and 102.6 apply existing-building provisions to the prescribed extent and generally allow lawful existing occupancy unless otherwise provided. | `src:usa-ky:kbc-2018-fourth-edition-pdf` |
| Mechanical | Kentucky mechanical requirements through KBC/KRC references and HVAC program | International Mechanical Code by KBC reference | edition not normalized in this report | applicable by reference and through HVAC permitting/inspection program | null | null | null | null | KBC Section 101.4.2 applies IMC provisions to mechanical systems; local/state HVAC program details require additional parsing for project-specific applicability. | `src:usa-ky:kbc-2018-fourth-edition-pdf`; `src:usa-ky:dhbc-hvac`; `src:usa-ky:815-kar-8-100` |
| Plumbing | Kentucky State Plumbing Code | State plumbing code / 815 KAR Chapter 20 | current KAR chapter; model edition not applicable | mandatory for plumbing systems | null | null | null | null | KBC Section 101.4.3 applies the Kentucky State Plumbing Code and requires installation under a Kentucky licensed master plumber and state plumbing inspector approval. | `src:usa-ky:kbc-2018-fourth-edition-pdf`; `src:usa-ky:dhbc-plumbing-code`; `src:usa-ky:815-kar-20-130` |
| Fuel Gas | Kentucky fuel-gas requirements | NFPA 54 National Fuel Gas Code | KBC references NFPA 54-12; KRC referenced-standards page references NFPA 54-09 | applicable by reference | null | 2024-12-03 | 2019-08-01 | 2019-08-01 | Applies through KBC/KRC referenced standards and only to the scope covered by the applicable code. | `src:usa-ky:kbc-2018-fourth-edition-pdf`; `src:usa-ky:krc-2018-third-edition-pdf` |
| Electrical | Kentucky electrical requirements | NFPA 70 National Electrical Code | NFPA 70-23 / 2023 NEC, with Kentucky transition notes | mandatory by Uniform State Building Code reference, subject to delayed provisions | null | 2024-12-03 | null | null | KRC delays NEC Articles 210.52(C), 230.67, and 314.27(C) until 2026-07-15; certain 2023 NEC GFCI requirements are delayed until specified UL standards are revised. | `src:usa-ky:krs-227-480`; `src:usa-ky:kbc-2018-fourth-edition-pdf`; `src:usa-ky:krc-2018-third-edition-pdf`; `src:usa-ky:815-kar-35-020` |
| Energy | Kentucky energy provisions | IECC | KBC references 2012 IECC and 2009 IECC for specified residential/Group R exceptions; KRC Chapter 11 uses 2009 IECC provisions | applicable by KBC/KRC | null | null | null | null | KBC Chapter 13 and KRC Chapter 11 control energy compliance to the extent referenced. | `src:usa-ky:kbc-2018-fourth-edition-pdf`; `src:usa-ky:krc-2018-third-edition-pdf` |
| Fire - construction references | KBC fire-prevention references | International Fire Code | edition not normalized in this report | limited by reference | null | 2024-12-03 | 2019-08-01 | 2019-08-01 | KBC Section 101.4.5 applies IFC matters affecting new construction only where specifically referenced in the KBC. | `src:usa-ky:kbc-2018-fourth-edition-pdf` |
| Fire - operational / prevention code | Kentucky Standards of Safety | NFPA-based state fire-safety regulation | 815 KAR 10:060; referenced standards include NFPA 70 2017 for that regulation | active | null | null | null | null | Applies to all buildings except one- and two-family dwellings; State Fire Marshal has primary jurisdiction unless a local ordinance establishes a local fire-inspection program. | `src:usa-ky:815-kar-10-060`; `src:usa-ky:kbc-2018-fourth-edition-pdf` |
| Accessibility | KBC accessibility provisions | 2015 IBC accessibility framework and referenced standards | current KBC edition | mandatory where KBC applies | null | 2024-12-03 | 2019-08-01 | 2019-08-01 | Administered as part of KBC plan review and inspection. | `src:usa-ky:kbc-2018-fourth-edition-pdf`; `src:usa-ky:dhbc-building-code-enforcement` |
| Elevator / Conveyance | Kentucky elevator and conveyance rules | State elevator statutes/regulations and KBC referenced standards | 815 KAR Chapter 4; standards tied to KBC incorporation | active | null | null | null | null | Elevator annual inspections and permits are regulated under 815 KAR 4:010 and 815 KAR 4:025; exact technical standards require device-specific review. | `src:usa-ky:dhbc-elevators`; `src:usa-ky:815-kar-4-010`; `src:usa-ky:815-kar-4-025` |

### 3.2 Adoption Records

#### `adoption:usa-ky:kbc-2018-fourth-edition`

- **Code family:** Building
- **State code name:** Kentucky Building Code
- **Base model code:** 2015 International Building Code
- **Kentucky amendment document:** 2018 Kentucky Building Code, Fourth Edition, February 2024
- **Scope:** all buildings constructed in Kentucky except one-family dwellings, two-family dwellings, and townhouses governed by 815 KAR 7:125, and manufactured homes governed by KRS 227.550 through 227.665.
- **Current amendment effective date:** 2024-12-03
- **Operative / mandatory plan-submittal rule:** 2019-08-01
- **Notes:** DHBC's public Building Code Enforcement page states that it adopted the 2018 KBC Second Edition and 2018 KRC Second Edition on 2019-08-03 and that current Kentucky-specific amendments include the KBC Fourth Edition. The KBC PDF and 815 KAR 7:120 support the current edition; date normalization still distinguishes current-amendment effective dates from 2019 code-family transition dates.
- **Source IDs:** `src:usa-ky:815-kar-7-120`; `src:usa-ky:kbc-2018-fourth-edition-pdf`; `src:usa-ky:dhbc-building-code-enforcement`

#### `adoption:usa-ky:krc-2018-third-edition`

- **Code family:** Residential
- **State code name:** Kentucky Residential Code
- **Base model code:** 2015 International Residential Code for One- and Two-Family Dwellings
- **Kentucky amendment document:** 2018 Kentucky Residential Code, Third Edition, August 2024
- **Scope:** single-family dwellings, two-family dwellings, and townhouses constructed in Kentucky.
- **Current amendment effective date:** 2024-12-03
- **Operative / mandatory permit rule:** 2019-08-01
- **Notes:** Permits, inspections, and certificates of occupancy are not required for a single-family dwelling unless required by local ordinance; this does not displace the KRC as the applicable statewide code.
- **Source IDs:** `src:usa-ky:815-kar-7-125`; `src:usa-ky:krc-2018-third-edition-pdf`; `src:usa-ky:dhbc-building-code-enforcement`

#### `adoption:usa-ky:electrical-nec-2023`

- **Code family:** Electrical
- **State code name:** Kentucky electrical requirements under the Uniform State Building Code
- **Base model code:** NFPA 70 National Electrical Code
- **Edition:** 2023 NEC / NFPA 70-23, with delayed Kentucky applicability notes
- **Mandatory date:** not fully normalized
- **Delayed provisions:** 2023 NEC Articles 210.52(C), 230.67, and 314.27(C) are not mandatory until 2026-07-15, and the corresponding 2017 NEC provisions apply until then. Selected 2023 NEC GFCI requirements are delayed until specified revisions are made to UL 943 and UL 101.
- **Source IDs:** `src:usa-ky:krs-227-480`; `src:usa-ky:kbc-2018-fourth-edition-pdf`; `src:usa-ky:krc-2018-third-edition-pdf`; `src:usa-ky:815-kar-35-020`

#### `adoption:usa-ky:energy-iecc`

- **Code family:** Energy
- **State code name:** Kentucky energy provisions in KBC/KRC
- **Base model code:** International Energy Conservation Code
- **Edition:** KBC references 2012 IECC and 2009 IECC for specified exceptions; KRC Chapter 11 uses 2009 IECC provisions.
- **Date status:** not normalized
- **Source IDs:** `src:usa-ky:kbc-2018-fourth-edition-pdf`; `src:usa-ky:krc-2018-third-edition-pdf`

#### `adoption:usa-ky:fire-standards-of-safety`

- **Code family:** Fire - operational / prevention code
- **State code name:** Kentucky Standards of Safety
- **Base / referenced codes:** 815 KAR 10:060 with NFPA and KBC references
- **Scope:** applies to all buildings except one- and two-family dwellings; supplements KBC in matters of fire safety.
- **Authority model:** State Fire Marshal primary jurisdiction unless a local fire-inspection program is established by ordinance; State Fire Marshal exclusive jurisdiction over state-owned property and Cabinet for Health and Family Services licensed facilities, subject to local fire chief request pathway.
- **Source IDs:** `src:usa-ky:815-kar-10-060`; `src:usa-ky:kbc-2018-fourth-edition-pdf`

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

Kentucky date normalization requires separating at least three date types: DHBC's public adoption note for the 2018 code family, the plan/permit transition dates in the KBC/KRC text, and the current 2024 regulatory amendment effective dates in the KAR histories. This report treats the KBC/KRC 2019-08-01 provisions as the verified operative/mandatory transition dates for plans or permits and treats 2024-12-03 as the current amendment-history effective date for the 2024 KBC/KRC incorporated-material updates.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `date-rule:usa-ky:kbc-plan-submittal-2019` | KBC-covered buildings | plan_submittal | 2019-08-01 | Plans submitted on or after this date must be designed and submitted to conform to the KBC. | unresolved for plans submitted before this date | `src:usa-ky:kbc-2018-fourth-edition-pdf` | partially_verified |
| `date-rule:usa-ky:krc-mandatory-permit-2019` | KRC-covered dwellings and townhouses | permit / mandatory_code | 2019-08-01 | Effective this date, the KRC is mandatory and no permit may be issued for construction under another building code. | no, for permits issued on/after trigger date | `src:usa-ky:krc-2018-third-edition-pdf` | partially_verified |
| `date-rule:usa-ky:kbc-current-kar-amendment-2024` | KBC regulation and incorporated Fourth Edition | regulatory_effective_date | 2024-12-03 | 815 KAR 7:120 history entry for the current amendment incorporating the Fourth Edition. | not applicable | `src:usa-ky:815-kar-7-120` | partially_verified |
| `date-rule:usa-ky:krc-current-kar-amendment-2024` | KRC regulation and incorporated Third Edition | regulatory_effective_date | 2024-12-03 | 815 KAR 7:125 history entry for the current amendment incorporating the Third Edition. | not applicable | `src:usa-ky:815-kar-7-125` | partially_verified |
| `date-rule:usa-ky:nec-delayed-articles-2026` | Electrical / NEC | delayed_mandatory_date | 2026-07-15 | 2023 NEC Articles 210.52(C), 230.67, and 314.27(C) are not mandatory until this date; corresponding 2017 NEC provisions apply until then. | yes, specified 2017 NEC provisions apply until trigger date | `src:usa-ky:krc-2018-third-edition-pdf` | partially_verified |
| `date-rule:usa-ky:nec-gfci-ul-trigger` | Electrical / NEC GFCI provisions | conditional_trigger | null | 2023 NEC GFCI provisions listed in KRC notes become mandatory only after specified UL 943 and UL 101 revisions. | yes, requirement delayed until UL trigger | `src:usa-ky:krc-2018-third-edition-pdf` | partially_verified |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building / Residential | no formal future model-code adoption normalized | null | null | null | null | null | watch | `src:usa-ky:dhbc-proposed-kar`; `src:usa-ky:815-kar-7-120`; `src:usa-ky:815-kar-7-125` | No formal KBC/KRC future edition adoption was normalized from the sources parsed for this report. |
| Elevator / Conveyance | proposed 815 KAR 4:010 and 815 KAR 4:025 amendments | 2026-04-10 | null | null | null | null | active_proposed_regulation | `src:usa-ky:dhbc-proposed-kar`; `src:usa-ky:815-kar-4-010`; `src:usa-ky:815-kar-4-025` | DHBC's proposed-regulation page lists proposed elevator-related amendments filed 2026-04-10. |
| Inspector / Electrical Inspector | proposed 815 KAR 7:070 and 815 KAR 35:015 amendments | 2026-04-10 | null | null | null | null | active_proposed_regulation | `src:usa-ky:dhbc-proposed-kar`; `src:usa-ky:815-kar-7-110` | Proposed changes may affect inspector certification and electrical inspector certification but were not parsed into adoption records. |
| Boiler / Pressure Vessel | proposed 815 KAR Chapter 15 amendments | 2026-04-10 | null | null | null | null | out_of_scope_watch | `src:usa-ky:dhbc-proposed-kar` | Included for monitoring because it is DHBC-regulated but not part of the core KBC/KRC adoption matrix. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| `applicability-rule:usa-ky:kbc-krc-scope-split` | Building / Residential | KRC-covered dwellings and townhouses | building is detached single-family dwelling, two-family dwelling, or townhouse | KBC excludes one-family dwellings, two-family dwellings, and townhouses governed by 815 KAR 7:125; other residential occupancies not within KRC scope comply with KBC. | `src:usa-ky:815-kar-7-120`; `src:usa-ky:815-kar-7-125`; `src:usa-ky:kbc-2018-fourth-edition-pdf`; `src:usa-ky:krc-2018-third-edition-pdf` | partially_verified |
| `applicability-rule:usa-ky:single-family-local-ordinance` | Residential | single-family dwelling | local ordinance requiring permits, inspections, or certificates of occupancy | Permits, inspections, and certificates of occupancy are not required for a single-family dwelling unless required by local ordinance; KRC compliance remains required. | `src:usa-ky:815-kar-7-125`; `src:usa-ky:dhbc-building-codes`; `src:usa-ky:krs-198b-060` | partially_verified |
| `applicability-rule:usa-ky:no-local-inspection-program` | Building / Residential | jurisdictions without local building inspection program | project located in jurisdiction without local building inspection program | All construction projects except single-family dwellings must be submitted to DHBC before construction. Single-family dwellings remain subject to KRC, with mechanical systems inspected by Kentucky HVAC and electrical by the local county electrical inspector when no single-family program exists. | `src:usa-ky:dhbc-building-codes` | partially_verified |
| `applicability-rule:usa-ky:fire-ifc-references` | Fire - construction references | new construction | IFC matter specifically referenced in KBC | IFC provisions apply to matters affecting or relating to new construction only where specifically referenced in KBC. | `src:usa-ky:kbc-2018-fourth-edition-pdf` | partially_verified |
| `applicability-rule:usa-ky:ksos-existing-buildings` | Fire - operational / prevention code | existing buildings | building subject to 815 KAR 10:060 | Kentucky Standards of Safety apply to all buildings except one- and two-family dwellings; KBC construction standards in effect at time of construction supersede conflicting fire-safety construction standards when a lawful certificate of occupancy was issued. | `src:usa-ky:815-kar-10-060`; `src:usa-ky:kbc-2018-fourth-edition-pdf` | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** Kentucky adopts base model codes and Kentucky-specific amendment documents by administrative regulation. For the KBC, 815 KAR 7:120 incorporates the 2015 IBC and the 2018 Kentucky Building Code, Fourth Edition, February 2024. For the KRC, 815 KAR 7:125 incorporates the 2015 IRC and the 2018 Kentucky Residential Code, Third Edition, August 2024.

**Where amendments are published:** DHBC code PDFs, 815 KAR 7:120 / 815 KAR 7:125 incorporation-by-reference records, and the Kentucky Administrative Register / LRC process for regulatory amendments.

**Amendment parsing status:** partially_parsed

### 5.2 State Amendment Sources

| Source ID | Amendment Set | Publication Path | Coverage | Status |
| --- | --- | --- | --- | --- |
| `src:usa-ky:kbc-2018-fourth-edition-pdf` | 2018 Kentucky Building Code, Fourth Edition, February 2024 | DHBC PDF incorporated by 815 KAR 7:120 | KBC amendments to 2015 IBC and referenced-code rules | partially_parsed |
| `src:usa-ky:krc-2018-third-edition-pdf` | 2018 Kentucky Residential Code, Third Edition, August 2024 | DHBC PDF incorporated by 815 KAR 7:125 | KRC amendments to 2015 IRC and referenced-code rules | partially_parsed |
| `src:usa-ky:815-kar-7-120` | KBC administrative regulation | LRC official KAR page | regulatory incorporation and current history | partially_parsed |
| `src:usa-ky:815-kar-7-125` | KRC administrative regulation | LRC official KAR page | regulatory incorporation and current history | partially_parsed |
| `src:usa-ky:krs-198b-080` | Uniform State Building Code amendments statute | LRC official statute page/PDF | statewide amendment process | partially_parsed |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Summary | Impact | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| `amendment:usa-ky:kbc-mini-maxi` | Building | KBC is described as a statewide uniform mandatory “mini/maxi” code; local governments shall not adopt or enforce another building code governing commercial construction. | Statewide uniformity and local building-code preemption for commercial construction. | `src:usa-ky:kbc-2018-fourth-edition-pdf`; `src:usa-ky:krs-198b-080` | partially_verified |
| `amendment:usa-ky:krc-mini-maxi` | Residential | KRC establishes minimum and maximum building-code requirements for detached single-family dwellings, two-family dwellings, and townhouses; local governments shall not adopt or enforce another building code for these units. | Statewide uniformity and local building-code preemption for KRC-covered residential units. | `src:usa-ky:krc-2018-third-edition-pdf`; `src:usa-ky:krs-198b-080` | partially_verified |
| `amendment:usa-ky:single-family-local-permits` | Residential | Single-family permits, inspections, and certificates of occupancy are not required unless required by local ordinance. | Local enforcement trigger differs from statewide code applicability. | `src:usa-ky:815-kar-7-125`; `src:usa-ky:krs-198b-060`; `src:usa-ky:dhbc-building-codes` | partially_verified |
| `amendment:usa-ky:nec-2023-delays` | Electrical | Selected 2023 NEC requirements are delayed until 2026-07-15 or until specified UL revisions occur. | Prevents treating all 2023 NEC provisions as immediately mandatory. | `src:usa-ky:krc-2018-third-edition-pdf` | partially_verified |
| `amendment:usa-ky:ifc-limited-reference` | Fire - construction references | International Fire Code provisions apply to new construction only where specifically referenced in KBC. | Keeps construction fire-code scope separate from operational fire-safety regulation. | `src:usa-ky:kbc-2018-fourth-edition-pdf`; `src:usa-ky:815-kar-10-060` | partially_verified |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-ky"
  model: "statewide_uniform_code_with_state_and_local_enforcement"
  enforcing_entities:
    - "Kentucky Department of Housing, Buildings and Construction / Division of Building Code Enforcement"
    - "city building departments"
    - "county building departments"
    - "local governments granted expanded jurisdiction under 815 KAR 7:110"
    - "State Fire Marshal and local fire chiefs/fire code officials for Kentucky Standards of Safety scope"
    - "certified electrical inspectors"
    - "state plumbing inspectors"
    - "state or approved local HVAC inspection programs"
  required_officials:
    - "certified building inspector, level III, for expanded local jurisdiction"
    - "certified electrical inspector for expanded local jurisdiction and electrical inspections"
    - "Kentucky licensed master plumber supervision and state plumbing inspector approval for plumbing systems"
  state_reserved_activities:
    - "DHBC-retained plan review, inspection, and enforcement responsibilities under expanded-jurisdiction agreements"
    - "institutional buildings"
    - "educational buildings unless otherwise agreed in writing"
    - "Cabinet for Health and Family Services licensed facilities"
    - "jurisdictions without local building inspection programs for non-single-family construction"
    - "State Fire Marshal exclusive jurisdiction over state-owned property and CHFS-licensed facilities under 815 KAR 10:060, subject to request pathway described there"
    - "elevator and conveyance plan approval, permitting, and inspections"
  source_ids:
    - "src:usa-ky:krs-198b-060"
    - "src:usa-ky:815-kar-7-110"
    - "src:usa-ky:dhbc-building-codes"
    - "src:usa-ky:815-kar-10-060"
    - "src:usa-ky:dhbc-elevators"
  verification_status: "partially_verified"
  confidence: 0.79
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-ky"
  model: "statewide_mini_maxi_building_code_with_statewide_amendment_process"
  applies_to_code_families:
    - "Building"
    - "Residential"
    - "Electrical as incorporated into the Uniform State Building Code"
    - "Referenced codes to the extent incorporated by KBC/KRC"
  approval_required: true
  approving_authority_id: "ahj:usa-ky:dhbc"
  filing_required: true
  registry_exists: true
  registry_source_ids:
    - "src:usa-ky:815-kar-7-120"
    - "src:usa-ky:815-kar-7-125"
    - "src:usa-ky:dhbc-building-code-enforcement"
  legal_basis_source_ids:
    - "src:usa-ky:krs-198b-080"
    - "src:usa-ky:kbc-2018-fourth-edition-pdf"
    - "src:usa-ky:krc-2018-third-edition-pdf"
  verification_status: "partially_verified"
  confidence: 0.76
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Kentucky local enforcement and local amendment authority should not be collapsed. Local governments may enforce the statewide KBC/KRC and may request expanded jurisdiction for plan review and inspection. That is separate from amending the code. The verified amendment path is state-level: interested parties may suggest amendments to DHBC, DHBC may amend through the KRS Chapter 13A regulatory process, and adopted amendments are effective statewide. For KBC and KRC building-code coverage, the DHBC PDFs describe local substitution with another building code as barred by the mini/maxi model.

Fire-prevention enforcement is related but distinct. The Kentucky Standards of Safety apply to all buildings except one- and two-family dwellings, supplement the KBC in fire-safety matters, and create State Fire Marshal and local fire-inspection program roles. This report does not fully classify whether local fire-prevention ordinances may impose local operational requirements beyond adopting/enforcing 815 KAR 10:060.

### 6.4 Known Local Amendment Registries

No local building-code amendment registry was verified. The verified registry path is statewide through LRC KAR records and DHBC incorporated-material PDFs.

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendments were parsed. Local ordinances may establish local enforcement programs and single-family permit/inspection requirements, but this report does not treat those ordinances as local code amendments.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: partially_started

Jurisdiction stack:

```text
Address
  -> State of Kentucky
  -> County
  -> Municipality / unincorporated county
  -> Local building inspection program status
  -> Expanded local jurisdiction agreement status, if any
  -> State-retained building categories, if applicable
  -> Fire inspection program / local fire chief jurisdiction status
  -> Trade-specific inspection authorities
  -> Applicable statewide KBC/KRC adoption record
  -> Applicable statewide trade-code or referenced-standard record
  -> Local ordinance triggers for single-family permits/inspections, if applicable
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | not selected | none | statewide | unknown | pending |
| County | not selected | none | statewide | unknown | pending |
| Municipality | not selected | none | statewide | unknown | pending |
| Fire District / local fire program | not selected | none | statewide | unknown | pending |
| Expanded local jurisdiction agreements | DHBC / local program records not yet selected | `src:usa-ky:815-kar-7-110` | statewide program framework; local agreement inventory not populated | unknown | pending |
| Single-family local ordinance triggers | local ordinances not yet selected | `src:usa-ky:815-kar-7-125`; `src:usa-ky:krs-198b-060` | statewide legal framework; local ordinance inventory not populated | unknown | pending |

### 7.3 AHJ Contact Data

No AHJ contact list was populated. DHBC division pages and contact pages are identified as starting points, but production AHJ resolution requires a directory of local building departments, local fire-inspection programs, electrical inspectors, plumbing inspectors, HVAC authorities, and elevator-inspection contacts.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title / Description | Source Type | Publisher | URL | Accessed | Supported Fields |
| --- | --- | --- | --- | --- | --- | --- |
| `src:usa-ky:krs-198b-050` | KRS 198B.050, Uniform State Building Code | statute | Kentucky Legislative Research Commission | https://apps.legislature.ky.gov/law/statutes/statute.aspx?id=46900 | 2026-06-25 | Uniform State Building Code authority |
| `src:usa-ky:krs-198b-060` | KRS 198B.060, Local enforcement of Uniform State Building Code | statute | Kentucky Legislative Research Commission | https://apps.legislature.ky.gov/law/statutes/statute.aspx?id=52659 | 2026-06-25 | Local enforcement, single-family local ordinance trigger |
| `src:usa-ky:krs-198b-080` | KRS 198B.080, Amendments to the Uniform State Building Code | statute | Kentucky Legislative Research Commission | https://apps.legislature.ky.gov/law/statutes/statute.aspx?id=46903 | 2026-06-25 | Statewide amendment process and statewide effect |
| `src:usa-ky:krs-227-480` | KRS 227.480, Authority to require electrical permits and electrical standards | statute | Kentucky Legislative Research Commission | https://apps.legislature.ky.gov/law/statutes/statute.aspx?id=50963 | 2026-06-25 | Electrical standards and NEC minimum standard |
| `src:usa-ky:815-kar-7-120` | 815 KAR 7:120, Kentucky Building Code | regulation | Kentucky Legislative Research Commission | https://apps.legislature.ky.gov/law/kar/titles/815/007/120/ | 2026-06-25 | KBC authority, scope, incorporation by reference, current amendment history |
| `src:usa-ky:815-kar-7-125` | 815 KAR 7:125, Kentucky Residential Code | regulation | Kentucky Legislative Research Commission | https://apps.legislature.ky.gov/law/kar/titles/815/007/125/ | 2026-06-25 | KRC authority, scope, single-family exception, incorporation by reference, current amendment history |
| `src:usa-ky:815-kar-7-110` | 815 KAR 7:110, Criteria for expanded local jurisdiction | regulation | Kentucky Legislative Research Commission | https://apps.legislature.ky.gov/law/kar/titles/815/007/110/ | 2026-06-25 | Expanded local jurisdiction, inspector requirements, state retained jurisdiction, preemption |
| `src:usa-ky:815-kar-10-060` | 815 KAR 10:060, Standards of Safety | regulation | Kentucky Legislative Research Commission | https://apps.legislature.ky.gov/law/kar/titles/815/010/060/ | 2026-06-25 | Fire-safety standards, State Fire Marshal/local fire jurisdiction, existing-building fire rules |
| `src:usa-ky:815-kar-20-130` | 815 KAR 20:130, House sewers and storm water piping; methods of installation | regulation | Kentucky Legislative Research Commission | https://apps.legislature.ky.gov/law/kar/titles/815/020/130/ | 2026-06-25 | Plumbing code authority example under KRS 318.130 and 815 KAR Chapter 20 |
| `src:usa-ky:815-kar-35-020` | 815 KAR 35:020, Electrical inspections | regulation | Kentucky Legislative Research Commission | https://apps.legislature.ky.gov/law/kar/titles/815/035/020/ | 2026-06-25 | Electrical inspection requirements |
| `src:usa-ky:815-kar-8-100` | 815 KAR 8:100, Local HVAC inspection program | regulation | Kentucky Legislative Research Commission | https://apps.legislature.ky.gov/law/kar/titles/815/008/100/ | 2026-06-25 | Local HVAC program framework |
| `src:usa-ky:815-kar-4-010` | 815 KAR 4:010, Annual inspection of elevators, chairlifts, fixed guideway systems, and platform lifts | regulation | Kentucky Legislative Research Commission | https://apps.legislature.ky.gov/law/kar/titles/815/004/010/ | 2026-06-25 | Elevator annual inspection requirements |
| `src:usa-ky:815-kar-4-025` | 815 KAR 4:025, Permit and inspection fees for new and altered elevators, chairlifts, fixed guideway systems, and platform lifts | regulation | Kentucky Legislative Research Commission | https://apps.legislature.ky.gov/law/kar/titles/815/004/025/ | 2026-06-25 | Elevator permit and inspection framework |
| `src:usa-ky:dhbc-building-code-enforcement` | DHBC Division of Building Code Enforcement page | agency_page | Kentucky Department of Housing, Buildings and Construction | https://dhbc.ky.gov/newstatic_info.aspx?static_id=297 | 2026-06-25 | DHBC role, current KBC/KRC amendment links, adoption note, division responsibilities |
| `src:usa-ky:dhbc-building-codes` | DHBC Building Codes page | agency_page | Kentucky Department of Housing, Buildings and Construction | https://dhbc.ky.gov/newstatic_info.aspx?static_id=333 | 2026-06-25 | State/local enforcement split, jurisdictions without local inspection program, KBC/KRC effective agency notes |
| `src:usa-ky:dhbc-plumbing-code` | DHBC Kentucky Plumbing Law, Regulations and Code page | agency_page | Kentucky Department of Housing, Buildings and Construction | https://dhbc.ky.gov/newstatic_info.aspx?static_id=342 | 2026-06-25 | Plumbing law/regulation portal |
| `src:usa-ky:dhbc-plumbing-division` | DHBC Division of Plumbing page | agency_page | Kentucky Department of Housing, Buildings and Construction | https://dhbc.ky.gov/newstatic_info.aspx?static_id=337 | 2026-06-25 | Plumbing permit and licensed plumber requirements |
| `src:usa-ky:dhbc-hvac` | DHBC Division of HVAC page | agency_page | Kentucky Department of Housing, Buildings and Construction | https://dhbc.ky.gov/newstatic_info.aspx?static_id=335 | 2026-06-25 | HVAC administration and licensing framework |
| `src:usa-ky:dhbc-elevators` | DHBC Elevators page | agency_page | Kentucky Department of Housing, Buildings and Construction | https://dhbc.ky.gov/newstatic_info.aspx?static_id=328 | 2026-06-25 | Elevator plan/specification approval and permitting |
| `src:usa-ky:dhbc-proposed-kar` | DHBC Proposed Administrative Regulations page | agency_page | Kentucky Department of Housing, Buildings and Construction | https://dhbc.ky.gov/kar.aspx | 2026-06-25 | Pending/proposed DHBC regulatory changes |
| `src:usa-ky:kbc-2018-fourth-edition-pdf` | 2018 Kentucky Building Code, Fourth Edition, February 2024 | incorporated_code_pdf | Kentucky Department of Housing, Buildings and Construction | https://dhbc.ky.gov/Documents/2018%20Kentucky%20Building%20Code%204th%20Ed.pdf | 2026-06-25 | KBC amendments, mini/maxi model, date rules, referenced-code provisions |
| `src:usa-ky:krc-2018-third-edition-pdf` | 2018 Kentucky Residential Code, Third Edition, August 2024 | incorporated_code_pdf | Kentucky Department of Housing, Buildings and Construction | https://dhbc.ky.gov/Documents/2018%20Kentucky%20Residential%20Code%203d%20Ed.pdf | 2026-06-25 | KRC amendments, mini/maxi model, date rules, NEC transition notes |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| `src:usa-ky:kbc-2018-fourth-edition-pdf` | incorporated_material_pdf | The PDF is official DHBC incorporated material but contains copyrighted/model-code material and Kentucky amendments. Parsed text and screenshots were used for key pages. | Treat as authoritative for the Kentucky amendment document; verify exact code text before quoting or building machine-readable rule clauses. |
| `src:usa-ky:krc-2018-third-edition-pdf` | incorporated_material_pdf | The PDF is official DHBC incorporated material but contains copyrighted/model-code material and Kentucky amendments. Parsed text and screenshots were used for key pages. | Treat as authoritative for the Kentucky amendment document; verify exact code text before quoting or building machine-readable rule clauses. |
| `src:usa-ky:815-kar-7-120` | lrc_alternate_view_duplicate_text | The LRC page may include current, previous, and alternate-view text blocks on one page, including markup from amendments. | Use current-regulation lines and the history entry; avoid relying on strikethrough alternate text as current law. |
| `src:usa-ky:815-kar-7-125` | lrc_alternate_view_duplicate_text | The LRC page may include current, previous, and alternate-view text blocks on one page, including markup from amendments. | Use current-regulation lines and the history entry; avoid relying on strikethrough alternate text as current law. |
| `src:usa-ky:krs-198b-060` | lrc_pdf_text_extract | Search result returned statute PDF text rather than a clean line-addressable HTML page in this pass. | Reopen and archive the official statute PDF in a later validation pass before production use. |
| `src:usa-ky:krs-198b-080` | lrc_pdf_text_extract | Search result returned statute PDF text rather than a clean line-addressable HTML page in this pass. | Reopen and archive the official statute PDF in a later validation pass before production use. |
| `src:usa-ky:krs-227-480` | lrc_pdf_text_extract | Search result returned statute PDF text rather than a clean line-addressable HTML page in this pass. | Reopen and archive the official statute PDF in a later validation pass before production use. |
| `src:usa-ky:dhbc-proposed-kar` | dynamic_agency_page | Proposed-regulation list can change frequently and should be rechecked close to production release. | Use only as monitoring evidence; do not treat proposed rules as effective law. |
| `src:usa-ky:dhbc-building-code-enforcement` | agency_summary_page | Agency page provides summary and links; the binding authority remains the statute/regulation and incorporated material. | Use for agency context and current links; defer to KAR and incorporated materials on code text. |
| `src:usa-ky:dhbc-building-codes` | agency_summary_page | Agency page provides practical jurisdiction summaries, including effective-date notes that differ from code text transition clauses. | Keep as explanatory evidence and preserve date distinctions in normalized records. |

### 8.3 Supplemental Sources

None used in this report.

### 8.4 Source Extraction Metadata

| Extraction ID | Source ID | Extraction Method | Extracted Facts | Extracted On | Notes |
| --- | --- | --- | --- | --- | --- |
| `extract:usa-ky:001` | `src:usa-ky:815-kar-7-120` | web_open | KRS authority, KBC mandatory status, KBC/KRC/manufactured-home scope split, incorporation of 2015 IBC and 2018 KBC Fourth Edition, history effective 2024-12-03 | 2026-06-25 | Official LRC KAR page |
| `extract:usa-ky:002` | `src:usa-ky:815-kar-7-125` | web_open | KRS authority, KRC mandatory status, single-family permit/inspection exception, incorporation of 2015 IRC and 2018 KRC Third Edition, history effective 2024-12-03 | 2026-06-25 | Official LRC KAR page |
| `extract:usa-ky:003` | `src:usa-ky:kbc-2018-fourth-edition-pdf` | pdf_text_and_screenshot | KBC cover, mini/maxi statement, state amendment process, effective date rule, referenced-code provisions | 2026-06-25 | PDF screenshots used for cover/preface/date and reference pages |
| `extract:usa-ky:004` | `src:usa-ky:krc-2018-third-edition-pdf` | pdf_text_and_screenshot | KRC cover, mini/maxi statement, mandatory date rule, NEC 2023 delayed provisions | 2026-06-25 | PDF screenshots used for cover/preface/date and NEC notes |
| `extract:usa-ky:005` | `src:usa-ky:815-kar-7-110` | web_open | expanded local jurisdiction, local inspector requirements, state retained jurisdiction, department preemption | 2026-06-25 | Official LRC KAR page |
| `extract:usa-ky:006` | `src:usa-ky:815-kar-10-060` | web_open | Kentucky Standards of Safety, State Fire Marshal/local fire jurisdiction, existing-building fire construction standards | 2026-06-25 | Official LRC KAR page |
| `extract:usa-ky:007` | `src:usa-ky:dhbc-building-codes` | web_open | state/local enforcement split, lack of local inspection program rules, DHBC practical effective-date summary | 2026-06-25 | Agency guidance page |
| `extract:usa-ky:008` | `src:usa-ky:dhbc-proposed-kar` | web_open | proposed 2026 administrative regulations list | 2026-06-25 | Dynamic agency page |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 1.00 | none | Core KBC/KRC authority and adoption fields have official-source support; remaining trade/date/AHJ gaps are explicit. |
| report | risk.overall_confidence | 0.68 | verified | 1.00 | none | Confidence reflects strong core sources and unresolved trade-code/AHJ/local fire-scope details. |
| `ahj:usa-ky:dhbc` | primary authority | Kentucky Department of Housing, Buildings and Construction | partially_verified | 0.86 | `src:usa-ky:815-kar-7-120`; `src:usa-ky:815-kar-7-125`; `src:usa-ky:dhbc-building-code-enforcement` | Department is defined in the regulations and agency page supports division role. |
| `adoption:usa-ky:kbc-2018-fourth-edition` | base model code | 2015 International Building Code | verified | 0.90 | `src:usa-ky:815-kar-7-120`; `src:usa-ky:kbc-2018-fourth-edition-pdf` | Text contains a likely omission of “Code” after “2015 International Building”; context and PDF support IBC. |
| `adoption:usa-ky:kbc-2018-fourth-edition` | Kentucky amendment edition | 2018 Kentucky Building Code, Fourth Edition, February 2024 | verified | 0.91 | `src:usa-ky:815-kar-7-120`; `src:usa-ky:kbc-2018-fourth-edition-pdf` | Current KAR and DHBC PDF agree. |
| `adoption:usa-ky:krc-2018-third-edition` | base model code | 2015 International Residential Code for One- and Two-Family Dwellings | verified | 0.91 | `src:usa-ky:815-kar-7-125`; `src:usa-ky:krc-2018-third-edition-pdf` | Current KAR and DHBC PDF agree. |
| `adoption:usa-ky:krc-2018-third-edition` | Kentucky amendment edition | 2018 Kentucky Residential Code, Third Edition, August 2024 | verified | 0.91 | `src:usa-ky:815-kar-7-125`; `src:usa-ky:krc-2018-third-edition-pdf` | Current KAR and DHBC PDF agree. |
| `date-rule:usa-ky:kbc-plan-submittal-2019` | operative date | 2019-08-01 | verified | 0.88 | `src:usa-ky:kbc-2018-fourth-edition-pdf` | Plan-submittal trigger is in KBC Section 120.1. |
| `date-rule:usa-ky:krc-mandatory-permit-2019` | mandatory date | 2019-08-01 | verified | 0.88 | `src:usa-ky:krc-2018-third-edition-pdf` | Mandatory/no-other-code permit trigger is in KRC Section R116.1. |
| `local-amendment-rule:usa-ky` | model | statewide_mini_maxi_building_code_with_statewide_amendment_process | partially_verified | 0.76 | `src:usa-ky:kbc-2018-fourth-edition-pdf`; `src:usa-ky:krc-2018-third-edition-pdf`; `src:usa-ky:krs-198b-080` | Local operational fire-safety ordinance scope remains a separate issue. |
| `local-enforcement:usa-ky` | model | statewide_uniform_code_with_state_and_local_enforcement | partially_verified | 0.79 | `src:usa-ky:krs-198b-060`; `src:usa-ky:815-kar-7-110`; `src:usa-ky:dhbc-building-codes` | Local AHJ inventories are not populated. |
| `adoption:usa-ky:electrical-nec-2023` | delayed NEC provisions | 2026-07-15 and UL-triggered delays | partially_verified | 0.80 | `src:usa-ky:krc-2018-third-edition-pdf` | KRC footnotes are clear; KBC/KAR legal interaction should be validated before production. |
| `adoption:usa-ky:fire-standards-of-safety` | operational fire authority | Kentucky Standards of Safety / State Fire Marshal primary jurisdiction | partially_verified | 0.72 | `src:usa-ky:815-kar-10-060`; `src:usa-ky:kbc-2018-fourth-edition-pdf` | KRS 227.320 local ordinance process needs follow-up. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Every `src:usa-ky:*` ID cited in the body is present in section 8.1. |
| All authority IDs resolve | pass | Authority IDs are defined in sections 2.1 and 2.2, with graph edges tied to those IDs. |
| All current code families have adoption records | fail | The adoption matrix is complete, but normalized adoption records were created only for the highest-value supported code families. Several trade rows remain matrix-only. |
| Building and operational fire code are separated | pass | KBC construction fire references and 815 KAR 10:060 operational/existing-building fire safety are separate rows and authorities. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Current amendment effective dates, 2019 operative/mandatory dates, and unresolved adoption dates are separated. |
| Effective dates are valid ISO dates | pass | Entered dates use ISO format; unresolved fields are `null`. |
| No impossible date sequences | pass | The apparent 2024 effective date after 2019 mandatory date is explicitly labeled as the current amendment-history effective date, not the original mandatory date. |
| Transition rules have explicit trigger conditions | pass | KBC, KRC, and NEC transition triggers are captured. |
| Permit-date logic is captured where applicable | pass | KRC no-other-code permit rule and single-family local ordinance trigger are captured. |
| Local enforcement model classified | pass | State/local shared model with expanded jurisdiction and DHBC reserved responsibilities is classified. |
| Local amendment rule classified | pass | Statewide mini/maxi building-code model and state amendment pathway are classified, with fire-scope caveat. |
| AHJ confirmation metadata present | fail | No AHJ directory, contact, or local program inventory was populated. |
| Official-source caveats captured | pass | Section 8.2 flags PDF, LRC alternate-view, statute PDF extract, and agency summary limitations. |
| Leftover template markers removed | pass | No template placeholder markers from the baseline are intentionally present. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `issue:usa-ky:001` | medium | date normalization | Normalize and document the relationship among DHBC's 2019 adoption/effective agency notes, KBC/KRC 2019-08-01 text transition rules, and 2024-12-03 KAR amendment effective histories. | Pull archived 2019 KAR history, DHBC adoption records, and Register notices; add date-specific adoption records if supported. | null | null | open |
| `issue:usa-ky:002` | medium | fire authority | Complete statute-level review of KRS 227.320 and local fire-inspection program authority to classify local fire-prevention ordinance scope. | Extract KRS 227.320 and any State Fire Marshal guidance; update local amendment vs local fire enforcement sections. | null | null | open |
| `issue:usa-ky:003` | medium | trade code editions | Parse KBC/KRC referenced standards and trade chapters for IMC, IEBC, IFC, accessibility, elevator, and energy edition details. | Extract KBC/KRC Chapter 35/reference tables and trade-specific regulations into normalized records. | null | null | open |
| `issue:usa-ky:004` | high | AHJ resolution | No local AHJ directory, boundary data, local building inspection program inventory, local fire program inventory, or single-family ordinance trigger inventory is populated. | Identify authoritative state/local datasets and create AHJ resolver inputs. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| `watch:usa-ky:815-kar-7-120` | `src:usa-ky:815-kar-7-120` | html_diff | monthly | KBC regulation text, incorporated material, or history changes | 2026-06-25 | active |
| `watch:usa-ky:815-kar-7-125` | `src:usa-ky:815-kar-7-125` | html_diff | monthly | KRC regulation text, incorporated material, or history changes | 2026-06-25 | active |
| `watch:usa-ky:dhbc-building-code-enforcement` | `src:usa-ky:dhbc-building-code-enforcement` | html_diff | monthly | New KBC/KRC amendment PDFs or adoption notes | 2026-06-25 | active |
| `watch:usa-ky:dhbc-proposed-kar` | `src:usa-ky:dhbc-proposed-kar` | html_diff | biweekly | New proposed regulations, especially KBC/KRC, electrical inspector, elevator, fire, or trade rules | 2026-06-25 | active |
| `watch:usa-ky:815-kar-10-060` | `src:usa-ky:815-kar-10-060` | html_diff | monthly | Kentucky Standards of Safety or fire-jurisdiction changes | 2026-06-25 | active |
| `watch:usa-ky:815-kar-35-020` | `src:usa-ky:815-kar-35-020` | html_diff | monthly | Electrical inspection requirements or NEC implementation changes | 2026-06-25 | active |
| `watch:usa-ky:815-kar-7-110` | `src:usa-ky:815-kar-7-110` | html_diff | quarterly | Expanded local jurisdiction criteria changes | 2026-06-25 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-24 | Generated baseline draft report | `report:usa-ky` | none | Codex | Baseline had statewide fields unresolved and no primary sources. |
| 2026-06-25 | Populated source-backed Kentucky report | `ahj:usa-ky:dhbc`; `adoption:usa-ky:kbc-2018-fourth-edition`; `adoption:usa-ky:krc-2018-third-edition`; `local-enforcement:usa-ky`; `local-amendment-rule:usa-ky` | `src:usa-ky:815-kar-7-120`; `src:usa-ky:815-kar-7-125`; `src:usa-ky:kbc-2018-fourth-edition-pdf`; `src:usa-ky:krc-2018-third-edition-pdf`; `src:usa-ky:815-kar-7-110`; `src:usa-ky:815-kar-10-060` | ChatGPT | Upgraded to `partially_verified` because core authority and code-adoption fields are supported by official Kentucky sources; unresolved items remain explicit. |
