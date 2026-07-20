---
state:
  state_id: "US-MD"
  name: "Maryland"
  abbreviation: "MD"
report:
  report_id: "state-report:usa-md"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-25"
  last_verified: "2026-06-25"
  reviewed_by: null
risk:
  overall_confidence: 0.72 # 0.00 - 1.00
  risk_flags:
    - "hybrid_trade_code_authority"
    - "plumbing_mechanical_current_edition_needs_board_confirmation"
    - "electrical_code_scope_and_local_amendments_not_fully_parsed"
    - "local_amendment_registry_not_cataloged"
    - "elevator_livability_and_safety_glazing_not_fully_parsed"
  open_questions_count: 6

---

# State Building Code Authority Report: Maryland

## 1. Executive Summary

- **Authority model:** Maryland uses a statewide adoption / local enforcement model for the Maryland Building Performance Standards (MBPS). The Maryland Department of Labor, Division of Labor and Industry, Building Codes Administration is the primary statewide building-code authority for the MBPS, while local jurisdictions implement and enforce the standards. Operational fire prevention is governed separately through the State Fire Prevention Code, administered by the State Fire Prevention Commission / State Fire Marshal and legally designated local fire officials. Plumbing, fuel gas, mechanical, electrical, accessibility, rehabilitation, green construction, and conveyance requirements are layered through separate Maryland statutes, COMAR chapters, boards, or agency programs.

- **Statewide code status:** The MBPS currently incorporate the 2021 IBC, 2021 IRC, and 2021 IECC with Maryland modifications. The Maryland Building Rehabilitation Code incorporates the 2021 IEBC. The State Fire Prevention Code currently incorporates NFPA 1 (2024) and NFPA 101 (2024). The International Green Construction Code chapter now incorporates the 2024 IgCC, effective 2026-05-25. Trade-code rows are supported by the 2026 Maryland code matrix and current COMAR text, but plumbing, fuel-gas, and mechanical editions need targeted follow-up because the current COMAR chapters use a "most recent version" formulation while the agency matrix still lists specific editions.

- **Local enforcement model:** Local jurisdictions implement and enforce the MBPS. At minimum, local implementation includes plan review, building permit issuance, inspections, and use-and-occupancy certificates. County or municipal responsibility depends on the statutory local-jurisdiction rules in Public Safety Article §12-505. Fire-code enforcement is by the State Fire Marshal, legally designated county or municipal fire officials, or other persons legally appointed by the State Fire Marshal.

- **Local amendment posture:** Local jurisdictions may adopt local MBPS amendments for local conditions, but may not weaken the energy code, automatic fire sprinkler provisions, or wind / wind-borne-debris requirements. Local MBPS amendments prevail locally when properly adopted and must be submitted to the Maryland Department of Labor within the statutory / regulatory notice windows. Local fire-code, electrical-code, and trade-code local amendment coverage was not exhaustively cataloged in this pass.

- **Known transition periods or pending changes:** The 2021 I-code MBPS cycle became effective 2023-05-29, and Maryland Department of Labor guidance states local jurisdictions had to begin implementing and enforcing the 2021 codes by 2024-05-29. The State Fire Prevention Code was amended effective 2025-06-23 to the 2024 NFPA 1 / NFPA 101 cycle. COMAR 09.12.57 was amended effective 2026-05-25 to incorporate the 2024 IgCC. Future MBPS adoption monitoring is required because Public Safety Article §12-503 requires subsequent IBC / IECC versions to be adopted within statutory timeframes, while the current official matrix still lists the 2021 IBC / IRC / IECC for the MBPS.

- **Production readiness:** partially_ready_for_statewide_authority_and_current_adoption_fields; not_ready_for_full_local_ordinance_or_AHJ_resolution.

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: The Maryland Department of Labor adopts the MBPS by regulation; the Building
    Codes Administration is the relevant Labor unit for statewide building-code administration.
  confidence: 0.9
  source_ids:
  - src:usa-md:ps-12-503
  - src:usa-md:labor-building-admin
  - src:usa-md:comar-09-12-51
- topic: Primary building / residential / energy editions
  finding: MBPS incorporates the 2021 IBC, 2021 IRC, and 2021 IECC with Maryland modifications.
  confidence: 0.92
  source_ids:
  - src:usa-md:comar-09-12-51-04
  - src:usa-md:comar-09-12-51-05
  - src:usa-md:code-matrix-2026
- topic: MBPS transition rule
  finding: Department guidance for the 2021 cycle states an effective date of 2023-05-29
    and local implementation/enforcement by 2024-05-29.
  confidence: 0.86
  source_ids:
  - src:usa-md:labor-2021-cycle
  - src:usa-md:ps-12-505
  - src:usa-md:comar-09-12-51-06
- topic: Existing building / rehabilitation code
  finding: Maryland Building Rehabilitation Code incorporates the 2021 IEBC.
  confidence: 0.88
  source_ids:
  - src:usa-md:comar-09-12-58
  - src:usa-md:code-matrix-2026
- topic: Green construction
  finding: COMAR 09.12.57 incorporates the 2024 IgCC; this supersedes older 2021-cycle
    and March 2026 matrix references for the IgCC row.
  confidence: 0.85
  source_ids:
  - src:usa-md:comar-09-12-57
  - src:usa-md:labor-2021-cycle
  - src:usa-md:code-matrix-2026
- topic: Fire code authority and editions
  finding: The State Fire Prevention Code incorporates NFPA 1 (2024) and NFPA 101
    (2024), with enforcement by the State Fire Marshal or legally designated local
    fire officials.
  confidence: 0.9
  source_ids:
  - src:usa-md:comar-29-06-01
  - src:usa-md:code-matrix-2026
- topic: Local enforcement
  finding: Local jurisdictions enforce MBPS; minimum activities include plan review,
    permit issuance, inspections, and occupancy certificates.
  confidence: 0.9
  source_ids:
  - src:usa-md:ps-12-505
  - src:usa-md:comar-09-12-51-08
- topic: Local amendments
  finding: Local MBPS amendments are allowed but bounded by statutory floors for energy,
    sprinklers, and wind provisions and must be submitted to the Department.
  confidence: 0.88
  source_ids:
  - src:usa-md:ps-12-504
  - src:usa-md:comar-09-12-51-05
- topic: Trade-code editions
  finding: Plumbing, fuel-gas, mechanical, and electrical rows are partially verified
    from the current code matrix and COMAR chapters; edition reconciliation remains
    open for plumbing / fuel gas / mechanical due current COMAR wording.
  confidence: 0.6
  source_ids:
  - src:usa-md:code-matrix-2026
  - src:usa-md:comar-09-20-01
  - src:usa-md:comar-09-15-05
  - src:usa-md:comar-09-12-50-02-1
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | `ahj:usa-md:labor-building-codes` |
| Authority name | Maryland Department of Labor, Division of Labor and Industry, Building Codes Administration |
| Authority type | state agency / code administration unit |
| Legal basis | Public Safety Article §12-503; COMAR 09.12.51 |
| Role | Adopts the MBPS by regulation; publishes and administers statewide building-code programs including MBPS, Maryland Accessibility Code, Maryland Building Rehabilitation Code, Model Performance Code, Minimum Livability Code, and related state-administered code programs. |
| Enforcement model | Statewide adoption with local jurisdiction implementation and enforcement for ordinary MBPS building permits; state-administered programs remain separately administered where assigned by law. |
| Source IDs | `src:usa-md:ps-12-503`; `src:usa-md:labor-building-admin`; `src:usa-md:labor-building-codes`; `src:usa-md:comar-09-12-51` |
| Verification status | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | `ahj:usa-md:labor-building-codes` | Maryland Department of Labor / Building Codes Administration | Adopts and modifies the IBC as part of MBPS; local jurisdictions enforce. | Public Safety Article §§12-501, 12-503, 12-505; COMAR 09.12.51 | `src:usa-md:ps-12-501`; `src:usa-md:ps-12-503`; `src:usa-md:ps-12-505`; `src:usa-md:comar-09-12-51-04` | partially_verified |
| Residential | `ahj:usa-md:labor-building-codes` | Maryland Department of Labor / Building Codes Administration | Adopts and modifies the IRC as part of MBPS; local jurisdictions enforce. | Public Safety Article §§12-501, 12-503, 12-505; COMAR 09.12.51 | `src:usa-md:ps-12-501`; `src:usa-md:ps-12-503`; `src:usa-md:ps-12-505`; `src:usa-md:comar-09-12-51-04` | partially_verified |
| Existing Building / Rehabilitation | `ahj:usa-md:labor-building-codes` | Maryland Department of Labor / Building Codes Administration | Adopts Maryland Building Rehabilitation Code; local jurisdictions are responsible for implementation and enforcement. | Public Safety Article §§12-1004, 12-1007(a); COMAR 09.12.58 | `src:usa-md:comar-09-12-58`; `src:usa-md:code-matrix-2026` | partially_verified |
| Mechanical | `ahj:usa-md:hvacr-board` | Maryland State Board of Heating, Ventilation, Air-Conditioning, and Refrigeration Contractors | Adopts IMC / IRC HVACR provisions for HVACR work; DLI / Model Performance Code applies to modular and state-owned / leased / operated / controlled buildings; local adoption may also apply. | Business Regulation Article §§9A-205 and 9A-310(a)(1)(iii); COMAR 09.15.05; COMAR 09.12.50 | `src:usa-md:comar-09-15-05`; `src:usa-md:comar-09-12-50-02-1`; `src:usa-md:code-matrix-2026` | partially_verified |
| Plumbing | `ahj:usa-md:plumbing-board` | Maryland Board of Plumbing | Adopts plumbing-code standards; DLI / Model Performance Code applies to modular and state-owned / leased / operated / controlled buildings; local adoption may also apply. | Business Occupations and Professions Article §§12-205, 12-207; COMAR 09.20.01; COMAR 09.12.50 | `src:usa-md:comar-09-20-01`; `src:usa-md:comar-09-12-50-02-1`; `src:usa-md:code-matrix-2026` | partially_verified |
| Fuel Gas | `ahj:usa-md:plumbing-board` | Maryland Board of Plumbing | Adopts fuel-gas and LP-gas standards through plumbing-board regulations; edition reconciliation remains open. | Business Occupations and Professions Article §§12-205, 12-207; COMAR 09.20.01 | `src:usa-md:comar-09-20-01`; `src:usa-md:code-matrix-2026` | partially_verified |
| Electrical | `ahj:usa-md:state-fire-marshal` / `ahj:usa-md:local-electrical-ahj` | State Fire Marshal / local electrical AHJs | Electrical code is not treated as an ordinary MBPS family; matrix identifies 2017 NEC under the State Fire Prevention Code for ordinary buildings, 2020 NEC under the Model Performance Code for modular/state-owned scope, and local jurisdiction adoption for ordinary local scope. | COMAR 09.12.50; State Fire Prevention Code / electrical-code statutes need deeper parsing | `src:usa-md:code-matrix-2026`; `src:usa-md:comar-09-12-50-02-1`; `src:usa-md:supp-ps-12-603-justia` | partially_verified |
| Energy | `ahj:usa-md:labor-building-codes` | Maryland Department of Labor / Building Codes Administration | Adopts and modifies IECC as part of MBPS; local jurisdictions may not weaken the IECC / energy provisions. | Public Safety Article §§12-503, 12-504; COMAR 09.12.51 | `src:usa-md:ps-12-503`; `src:usa-md:ps-12-504`; `src:usa-md:comar-09-12-51-04`; `src:usa-md:comar-09-12-51-05` | partially_verified |
| Green Construction | `ahj:usa-md:labor-building-codes` | Maryland Department of Labor / Building Codes Administration | Adopts IgCC by regulation; local jurisdictions may adopt and amend the IgCC. | Public Safety Article §§12-503(d), 12-504(a)(2); COMAR 09.12.57 | `src:usa-md:ps-12-503`; `src:usa-md:ps-12-504`; `src:usa-md:comar-09-12-57` | partially_verified |
| Fire - construction references | `ahj:usa-md:state-fire-prevention-commission` | Maryland State Fire Prevention Commission / State Fire Marshal | State Fire Prevention Code applies to new and existing fire/life-safety conditions, with exclusions; building-code construction references interact with local MBPS AHJs. | Public Safety Article §§6-206, 6-501; COMAR 29.06.01 | `src:usa-md:comar-29-06-01`; `src:usa-md:code-matrix-2026` | partially_verified |
| Fire - operational / prevention code | `ahj:usa-md:state-fire-prevention-commission` | Maryland State Fire Prevention Commission / State Fire Marshal | Adopts and enforces the State Fire Prevention Code through State Fire Marshal and legally designated local fire officials. | Public Safety Article §§6-206, 6-501; COMAR 29.06.01 | `src:usa-md:comar-29-06-01`; `src:usa-md:code-matrix-2026` | partially_verified |
| Accessibility | `ahj:usa-md:labor-building-codes` | Maryland Department of Labor / Building Codes Administration | Administers Maryland Accessibility Code; current matrix lists 2019 MAC and COMAR identifies applicable ADA / Maryland standards. | Public Safety Article §12-202(b); COMAR 09.12.53 | `src:usa-md:comar-09-12-53`; `src:usa-md:code-matrix-2026` | partially_verified |
| Elevator / Conveyance | `ahj:usa-md:labor-conveyance` | Maryland Department of Labor / Division of Labor and Industry | Current matrix identifies 2021 IBC and Labor requirements for elevator / conveying systems; detailed conveyance statute / regulation parsing remains open. | LABOR requirements not fully parsed in this pass | `src:usa-md:code-matrix-2026` | unresolved |
| Minimum Livability / Property Maintenance | `ahj:usa-md:labor-building-codes` | Maryland Department of Labor / Building Codes Administration | Building Codes Administration lists the Minimum Livability Code among state codes; 2012 IPMC is listed in the current code matrix. | COMAR 09.12.54 not parsed in this pass | `src:usa-md:labor-building-admin`; `src:usa-md:code-matrix-2026` | partially_verified |

### 2.3 Authority Hierarchy Notes

Maryland is not a simple single-code-board model. The MBPS are adopted statewide by Labor and enforced locally. The State Fire Prevention Code is a separate statewide fire/life-safety regime administered by the State Fire Prevention Commission / State Fire Marshal and legally designated local fire officials. Plumbing, fuel gas, mechanical, and electrical requirements are layered through board regulations, fire-code rules, Model Performance Code rules for modular / state buildings, and local adoption. For production AHJ resolution, MBPS enforcement should be resolved separately from fire-code enforcement and from trade-code officials.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| `edge:usa-md:001` | `ahj:usa-md:labor-building-codes` | adopts | MBPS: 2021 IBC, 2021 IRC, 2021 IECC with Maryland modifications | `src:usa-md:ps-12-503`; `src:usa-md:comar-09-12-51-04` | partially_verified |
| `edge:usa-md:002` | `ahj:usa-md:labor-building-codes` | assigns_implementation_and_enforcement_to | local jurisdictions for MBPS buildings / structures | `src:usa-md:ps-12-505`; `src:usa-md:comar-09-12-51-08` | partially_verified |
| `edge:usa-md:003` | `ahj:usa-md:local-jurisdiction` | may_amend_with_limits | MBPS and IgCC local amendments | `src:usa-md:ps-12-504`; `src:usa-md:comar-09-12-51-05` | partially_verified |
| `edge:usa-md:004` | `ahj:usa-md:state-fire-prevention-commission` | adopts | State Fire Prevention Code: NFPA 1 (2024) and NFPA 101 (2024) with Maryland amendments | `src:usa-md:comar-29-06-01` | partially_verified |
| `edge:usa-md:005` | `ahj:usa-md:state-fire-marshal` | enforces_or_delegates_enforcement | State Fire Prevention Code through State Fire Marshal, legally designated local fire officials, or appointed persons | `src:usa-md:comar-29-06-01` | partially_verified |
| `edge:usa-md:006` | `ahj:usa-md:plumbing-board` | adopts | Plumbing, fuel gas, NFPA 54, and NFPA 58 standards | `src:usa-md:comar-09-20-01`; `src:usa-md:code-matrix-2026` | partially_verified |
| `edge:usa-md:007` | `ahj:usa-md:hvacr-board` | adopts | Mechanical / HVACR code standards | `src:usa-md:comar-09-15-05`; `src:usa-md:code-matrix-2026` | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | Maryland Building Performance Standards | International Building Code | 2021 | current | unknown | 2023-05-29 | 2023-05-29 for state effective rule; local permit applicability controlled by local implementation | 2024-05-29 for local implementation/enforcement of the 2021 cycle | Local jurisdictions implement/enforce within 12 months after Department amendments; applies to permit applications received by local jurisdiction. | `src:usa-md:comar-09-12-51-04`; `src:usa-md:comar-09-12-51-06`; `src:usa-md:labor-2021-cycle`; `src:usa-md:code-matrix-2026` |
| Residential | Maryland Building Performance Standards | International Residential Code | 2021 | current | unknown | 2023-05-29 | 2023-05-29 for state effective rule; local permit applicability controlled by local implementation | 2024-05-29 for local implementation/enforcement of the 2021 cycle | Same MBPS transition rule as building. | `src:usa-md:comar-09-12-51-04`; `src:usa-md:comar-09-12-51-06`; `src:usa-md:labor-2021-cycle`; `src:usa-md:code-matrix-2026` |
| Existing Building / Rehabilitation | Maryland Building Rehabilitation Code | International Existing Building Code | 2021 | current | unknown | 2023-05-29 | 2023-05-29 | unknown | COMAR chapter updated in the 2021 cycle; local jurisdictions responsible for implementation and enforcement. | `src:usa-md:comar-09-12-58`; `src:usa-md:labor-2021-cycle`; `src:usa-md:code-matrix-2026` |
| Mechanical | Maryland HVACR / Mechanical Code; Model Performance Code for modular and state buildings | International Mechanical Code; IRC HVACR chapters | 2021 per current agency matrix and COMAR cross-reference; current COMAR now uses "most recent version" language | current_partially_verified | unknown | 2026-04-13 for current COMAR amendment to incorporation language | unknown | unknown | Edition reconciliation required between current COMAR wording and matrix row. | `src:usa-md:comar-09-15-05`; `src:usa-md:comar-09-12-50-02-1`; `src:usa-md:code-matrix-2026` |
| Plumbing | Maryland State Plumbing Code; Model Performance Code for modular and state buildings | International Plumbing Code | 2021 per current agency matrix and COMAR 09.12.50 cross-reference; current COMAR 09.20.01 now uses "most recent version" language | current_partially_verified | unknown | 2025-11-24 for current COMAR amendment to incorporation language | unknown | unknown | Edition reconciliation required between current COMAR wording and matrix row. | `src:usa-md:comar-09-20-01`; `src:usa-md:comar-09-12-50-02-1`; `src:usa-md:code-matrix-2026` |
| Fuel Gas | Maryland State Plumbing Code / Fuel Gas provisions | International Fuel Gas Code; NFPA 54; NFPA 58 | 2018 IFGC and 2017 NFPA 58 per current matrix; current COMAR now uses "most recent version" language | current_partially_verified | unknown | 2019-12-30 per matrix for IFGC / NFPA 58 rows; 2025-11-24 for current COMAR amendment to incorporation language | unknown | unknown | Edition reconciliation required. | `src:usa-md:comar-09-20-01`; `src:usa-md:code-matrix-2026` |
| Electrical | State Fire Prevention Code / local electrical code / Model Performance Code | National Electrical Code | 2017 NEC for State Fire Prevention Code matrix row; 2020 NEC for Model Performance Code modular and state-owned scope | current_partially_verified | unknown | unknown | unknown | unknown | Electrical scope is hybrid: State Fire Prevention Code, local electrical codes, and MPC scope differ. | `src:usa-md:code-matrix-2026`; `src:usa-md:comar-09-12-50-02-1`; `src:usa-md:supp-ps-12-603-justia` |
| Energy | Maryland Building Performance Standards | International Energy Conservation Code | 2021 | current | unknown | 2023-05-29 | 2023-05-29 for state effective rule; local permit applicability controlled by local implementation | 2024-05-29 for local implementation/enforcement of the 2021 cycle | Local jurisdictions may not weaken IECC / energy provisions. | `src:usa-md:ps-12-503`; `src:usa-md:ps-12-504`; `src:usa-md:comar-09-12-51-04`; `src:usa-md:labor-2021-cycle`; `src:usa-md:code-matrix-2026` |
| Green Construction | Maryland International Green Construction Code | International Green Construction Code | 2024 | current | unknown | 2026-05-25 | 2026-05-25 | unknown | Current COMAR incorporates 2024 IgCC; local jurisdictions may adopt and amend IgCC. | `src:usa-md:comar-09-12-57`; `src:usa-md:ps-12-503`; `src:usa-md:ps-12-504` |
| Fire - construction references | State Fire Prevention Code | NFPA 1 Fire Code; NFPA 101 Life Safety Code | 2024 | current | unknown | 2025-06-23 | 2025-06-23 for new-building / new-condition definition | unknown | Applies to new and existing buildings and conditions, subject to existing-condition rules and exclusions. | `src:usa-md:comar-29-06-01`; `src:usa-md:code-matrix-2026` |
| Fire - operational / prevention code | State Fire Prevention Code | NFPA 1 Fire Code; NFPA 101 Life Safety Code | 2024 | current | unknown | 2025-06-23 | 2025-06-23 for new-building / new-condition definition | unknown | Does not apply in Baltimore City except specified statutory scope; does not apply to buildings used solely as dwellings for not more than two families. | `src:usa-md:comar-29-06-01`; `src:usa-md:code-matrix-2026` |
| Accessibility | Maryland Accessibility Code | Maryland Accessibility Code / ADA Standards | 2019 MAC per current code matrix; COMAR uses 2010 ADA Standards and Maryland additional requirements | current_partially_verified | unknown | unknown | unknown | unknown | Detailed MAC edition history not fully parsed. | `src:usa-md:comar-09-12-53`; `src:usa-md:code-matrix-2026` |
| Elevator / Conveyance | Maryland Elevator / Conveying System requirements | 2021 IBC plus LABOR requirements | 2021 IBC matrix row; detailed conveyance code not parsed | unresolved | unknown | unknown | unknown | unknown | Matrix row captured; legal basis and amendment path unresolved. | `src:usa-md:code-matrix-2026` |
| Minimum Livability / Property Maintenance | Minimum Livability Code | International Property Maintenance Code | 2012 IPMC per current code matrix | current_partially_verified | unknown | unknown | unknown | unknown | COMAR 09.12.54 not parsed; row retained because matrix identifies statewide code family. | `src:usa-md:labor-building-admin`; `src:usa-md:code-matrix-2026` |

### 3.2 Adoption Records

| Adoption Record ID | Code Family | Authority ID | State Code Name | Base Code | Edition | Effective Date | Mandatory / Local Implementation Date | Scope | Source IDs | Verification Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `adoption:usa-md:mbps-2021-ibc` | Building | `ahj:usa-md:labor-building-codes` | Maryland Building Performance Standards | IBC | 2021 | 2023-05-29 | 2024-05-29 | Statewide MBPS; local jurisdiction enforcement | `src:usa-md:comar-09-12-51-04`; `src:usa-md:labor-2021-cycle`; `src:usa-md:code-matrix-2026` | verified_core |
| `adoption:usa-md:mbps-2021-irc` | Residential | `ahj:usa-md:labor-building-codes` | Maryland Building Performance Standards | IRC | 2021 | 2023-05-29 | 2024-05-29 | Statewide MBPS; local jurisdiction enforcement | `src:usa-md:comar-09-12-51-04`; `src:usa-md:labor-2021-cycle`; `src:usa-md:code-matrix-2026` | verified_core |
| `adoption:usa-md:mbps-2021-iecc` | Energy | `ahj:usa-md:labor-building-codes` | Maryland Building Performance Standards | IECC | 2021 | 2023-05-29 | 2024-05-29 | Statewide MBPS energy code; local amendments may not weaken | `src:usa-md:comar-09-12-51-04`; `src:usa-md:labor-2021-cycle`; `src:usa-md:ps-12-504` | verified_core |
| `adoption:usa-md:mbrc-2021-iebc` | Existing Building / Rehabilitation | `ahj:usa-md:labor-building-codes` | Maryland Building Rehabilitation Code | IEBC | 2021 | 2023-05-29 | unknown | Existing building repair, alteration, addition, and change of occupancy scope | `src:usa-md:comar-09-12-58`; `src:usa-md:labor-2021-cycle`; `src:usa-md:code-matrix-2026` | partially_verified |
| `adoption:usa-md:igcc-2024` | Green Construction | `ahj:usa-md:labor-building-codes` | International Green Construction Code | IgCC | 2024 | 2026-05-25 | unknown | State IgCC chapter; local adoption/amendment permitted | `src:usa-md:comar-09-12-57`; `src:usa-md:ps-12-503`; `src:usa-md:ps-12-504` | partially_verified |
| `adoption:usa-md:sfpc-2024-nfpa-1-101` | Fire - operational / prevention code | `ahj:usa-md:state-fire-prevention-commission` | State Fire Prevention Code | NFPA 1 / NFPA 101 | 2024 | 2025-06-23 | unknown | State fire/life-safety code with Baltimore City and one-/two-family dwelling exclusions | `src:usa-md:comar-29-06-01`; `src:usa-md:code-matrix-2026` | verified_core |
| `adoption:usa-md:plumbing-2021-ipc-matrix` | Plumbing | `ahj:usa-md:plumbing-board` | Maryland State Plumbing Code | IPC | 2021 per 2026 matrix | 2025-11-24 for current COMAR incorporation-language amendment | unknown | Plumbing-board scope plus local / MPC scope; edition reconciliation needed | `src:usa-md:comar-09-20-01`; `src:usa-md:code-matrix-2026` | partially_verified |
| `adoption:usa-md:fuel-gas-2018-ifgc-matrix` | Fuel Gas | `ahj:usa-md:plumbing-board` | State Plumbing Code fuel-gas provisions | IFGC | 2018 per 2026 matrix | 2019-12-30 per matrix; 2025-11-24 for current COMAR incorporation-language amendment | unknown | Fuel-gas scope; edition reconciliation needed | `src:usa-md:comar-09-20-01`; `src:usa-md:code-matrix-2026` | partially_verified |
| `adoption:usa-md:mechanical-2021-imc-matrix` | Mechanical | `ahj:usa-md:hvacr-board` | Maryland HVACR / Mechanical Code | IMC / IRC HVACR chapters | 2021 per 2026 matrix and COMAR 09.12.50 cross-reference | 2026-04-13 for current COMAR incorporation-language amendment | unknown | HVACR board scope plus local / MPC scope; edition reconciliation needed | `src:usa-md:comar-09-15-05`; `src:usa-md:code-matrix-2026`; `src:usa-md:comar-09-12-50-02-1` | partially_verified |
| `adoption:usa-md:electrical-nec-hybrid` | Electrical | `ahj:usa-md:state-fire-marshal` / `ahj:usa-md:local-electrical-ahj` | State Fire Prevention Code / local electrical code / MPC electrical row | NEC | 2017 NEC per State Fire Prevention Code matrix row; 2020 NEC under MPC scope | unknown | unknown | Hybrid electrical scope; official state statute source not fully parsed beyond supplemental source | `src:usa-md:code-matrix-2026`; `src:usa-md:comar-09-12-50-02-1`; `src:usa-md:supp-ps-12-603-justia` | partially_verified |
| `adoption:usa-md:accessibility-mac-2019-matrix` | Accessibility | `ahj:usa-md:labor-building-codes` | Maryland Accessibility Code | MAC / ADA Standards | 2019 MAC per 2026 matrix | unknown | unknown | Accessibility scope under MAC and ADA/Maryland requirements | `src:usa-md:comar-09-12-53`; `src:usa-md:code-matrix-2026` | partially_verified |
| `adoption:usa-md:elevator-2021-ibc-matrix` | Elevator / Conveyance | `ahj:usa-md:labor-conveyance` | Elevator / Conveying System requirements | IBC plus LABOR requirements | 2021 IBC matrix row | unknown | unknown | Matrix row only; full legal basis unresolved | `src:usa-md:code-matrix-2026` | unresolved |
| `adoption:usa-md:livability-2012-ipmc-matrix` | Minimum Livability / Property Maintenance | `ahj:usa-md:labor-building-codes` | Minimum Livability Code | IPMC | 2012 IPMC per 2026 matrix | unknown | unknown | Matrix row and agency list only; COMAR 09.12.54 not parsed | `src:usa-md:labor-building-admin`; `src:usa-md:code-matrix-2026` | partially_verified |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

Maryland separates adoption / effective dates from local implementation duties. For the 2021 MBPS cycle, Maryland Department of Labor guidance provides a statewide effective date of 2023-05-29 and a local implementation/enforcement date of 2024-05-29. Public Safety Article §12-505 and COMAR 09.12.51.06 support the 12-month local implementation concept. The State Fire Prevention Code uses a separate new-building / new-condition rule keyed to building permit issuance, actual construction start where permits are not required, change of occupancy, or commencement of a regulated condition on or after the chapter effective date.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `date-rule:usa-md:mbps-2021-effective` | MBPS 2021 IBC / IRC / IECC | effective_date | 2023-05-29 | Department adoption of 2021 cycle amendments | unknown | `src:usa-md:labor-2021-cycle`; `src:usa-md:comar-09-12-51-04` | partially_verified |
| `date-rule:usa-md:mbps-local-implementation` | MBPS local implementation | 12_month_local_implementation_period | 2024-05-29 for 2021 cycle | Local jurisdictions begin implementing and enforcing no later than 12 months after State adoption / amendment effective date | unknown | `src:usa-md:labor-2021-cycle`; `src:usa-md:ps-12-505`; `src:usa-md:comar-09-12-51-06` | partially_verified |
| `date-rule:usa-md:mbps-permit-application` | MBPS applicability | permit_application_rule | no single date | Building permit application received by local jurisdiction; Standards apply to buildings / structures with permit applications in local jurisdiction scope | unknown | `src:usa-md:ps-12-503`; `src:usa-md:comar-09-12-51-06` | partially_verified |
| `date-rule:usa-md:mbrc-2021-effective` | Maryland Building Rehabilitation Code | effective_date | 2023-05-29 | COMAR 09.12.58.03 amendment / 2021 IEBC incorporation | unknown | `src:usa-md:comar-09-12-58`; `src:usa-md:labor-2021-cycle` | partially_verified |
| `date-rule:usa-md:igcc-2024-effective` | International Green Construction Code | effective_date | 2026-05-25 | COMAR 09.12.57.02 amendment incorporating 2024 IgCC | unknown | `src:usa-md:comar-09-12-57` | partially_verified |
| `date-rule:usa-md:fire-2024-effective` | State Fire Prevention Code | effective_date | 2025-06-23 | COMAR 29.06.01 amendments to NFPA 1 / NFPA 101 incorporation and amendments | unknown | `src:usa-md:comar-29-06-01` | partially_verified |
| `date-rule:usa-md:fire-new-condition` | State Fire Prevention Code | permit_issuance_or_construction_start_rule | tied to effective date of chapter | New building or condition includes building permit issued on or after effective date, actual construction start on or after effective date where no permit is required, change in occupancy, or regulated condition commenced on or after effective date | unknown | `src:usa-md:comar-29-06-01` | partially_verified |
| `date-rule:usa-md:local-amendment-submission` | Local MBPS amendments | filing_notice_rule | 15 days before effective date; 5 days after emergency adoption | Local jurisdiction adopts MBPS local amendment | prior local law controls local adoption process; not fully parsed | `src:usa-md:ps-12-504`; `src:usa-md:comar-09-12-51-05` | partially_verified |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building / Residential / Energy | 2024 IBC / IRC / IECC or later successor cycle | unknown | unknown | unknown | unknown | unknown | active_watch | `src:usa-md:ps-12-503`; `src:usa-md:code-matrix-2026` | Statute requires subsequent versions on a statutory cycle, but official matrix still lists 2021 IBC / IRC / IECC as of 2026-03-18. |
| Green Construction | 2024 IgCC | known via current COMAR | unknown | 2026-05-25 | 2026-05-25 | unknown | captured | `src:usa-md:comar-09-12-57` | Already incorporated in current COMAR; matrix dated before this amendment still listed 2021 IgCC. |
| Plumbing / Fuel Gas | most recent IPC / IFGC / NFPA 54 / NFPA 58 formulation | unknown | unknown | 2025-11-24 | unknown | unknown | active_watch | `src:usa-md:comar-09-20-01`; `src:usa-md:code-matrix-2026` | Reconcile COMAR text with agency matrix before marking verified. |
| Mechanical | most recent IMC / IRC HVACR formulation | unknown | unknown | 2026-04-13 | unknown | unknown | active_watch | `src:usa-md:comar-09-15-05`; `src:usa-md:code-matrix-2026` | Reconcile COMAR text with agency matrix before marking verified. |
| Fire Prevention | next NFPA cycle | unknown | unknown | unknown | unknown | unknown | active_watch | `src:usa-md:comar-29-06-01` | Current code is 2024 NFPA 1 / NFPA 101 effective through 2025 amendments. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| `applicability-rule:usa-md:mbps-agritourism` | MBPS | qualifying agritourism buildings / structures | Project falls within COMAR 09.12.51.06C exception | COMAR identifies a limited exception for agritourism buildings / structures when specified statutory conditions are satisfied; detailed eligibility not parsed. | `src:usa-md:comar-09-12-51-06` | partially_verified |
| `applicability-rule:usa-md:mbrc-existing-buildings` | Existing Building / Rehabilitation | existing buildings undergoing repair, alterations, additions, or change of occupancy | Existing-building work within MBRC scope | MBPS / IRC modifications route existing-building work to the Maryland Building Rehabilitation Code. | `src:usa-md:comar-09-12-51-04`; `src:usa-md:comar-09-12-58` | partially_verified |
| `applicability-rule:usa-md:fire-existing-hazard` | Fire - operational / prevention code | existing buildings and conditions | Existing or previously approved condition predates effective date but AHJ determines hazard | Existing conditions generally are not newly subjected unless the AHJ determines the existing situation constitutes a public-welfare / safety hazard requiring correction; AHJ may modify requirements where impractical if reasonable safety is evident. | `src:usa-md:comar-29-06-01` | partially_verified |
| `applicability-rule:usa-md:fire-baltimore-city-exclusion` | Fire - operational / prevention code | Baltimore City buildings and conditions | Location in Baltimore City | State Fire Prevention Code does not apply in Baltimore City except to buildings and conditions specifically prescribed in Public Safety Article Title 6, Subtitle 4. | `src:usa-md:comar-29-06-01` | partially_verified |
| `applicability-rule:usa-md:fire-one-two-family-exclusion` | Fire - operational / prevention code | one- and two-family dwellings | Building used solely as dwelling house for not more than two families | State Fire Prevention Code does not apply to these dwellings as prescribed in Public Safety Article Title 6, Subtitle 3. | `src:usa-md:comar-29-06-01` | partially_verified |
| `applicability-rule:usa-md:model-performance-code` | Industrialized / modular and state-owned / leased / operated / controlled buildings | State-administered MPC scope | Project falls within Model Performance Code scope | Department page states the MPC applies to industrialized / modular buildings and State-owned / leased / operated / controlled buildings, and is not binding on subdivisions unless adopted locally. | `src:usa-md:labor-building-codes`; `src:usa-md:comar-09-12-50-02-1`; `src:usa-md:code-matrix-2026` | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** Maryland modifies adopted model codes through COMAR. MBPS modifications are in COMAR 09.12.51.04; the MBPS are constituted by the incorporated IBC / IRC / IECC as modified. State Fire Prevention Code amendments are embedded in COMAR 29.06.01. Plumbing and HVACR regulations now primarily incorporate codes by reference, with edition / amendment reconciliation open because recent COMAR amendments changed incorporation wording.

**Where amendments are published:** COMAR chapters and agency matrix / code pages. Local amendments are adopted by local law and submitted to the Maryland Department of Labor; a complete local-amendment registry was not parsed in this pass.

**Amendment parsing status:** core_state_amendment_paths_identified; detailed_section_level_amendments_not_fully_extracted.

### 5.2 State Amendment Sources

| Source / Amendment Set | Code Family | Publication Path | Current Parsing Status | Source IDs |
| --- | --- | --- | --- | --- |
| COMAR 09.12.51.04 | Building / Residential / Energy | Library of Maryland Regulations | incorporation and several high-impact modifications captured; detailed model-section diff not fully extracted | `src:usa-md:comar-09-12-51-04` |
| COMAR 09.12.51.05 | MBPS / local amendments | Library of Maryland Regulations | local amendment authority and state floors captured | `src:usa-md:comar-09-12-51-05` |
| COMAR 09.12.58.03 | Existing Building / Rehabilitation | Library of Maryland Regulations | 2021 IEBC incorporation and local enforcement captured | `src:usa-md:comar-09-12-58` |
| COMAR 09.12.57.02 | Green Construction | Library of Maryland Regulations | 2024 IgCC incorporation captured | `src:usa-md:comar-09-12-57` |
| COMAR 29.06.01 | Fire - operational / prevention code | Library of Maryland Regulations | 2024 NFPA 1 / NFPA 101 incorporation, enforcement, applicability, and major exclusions captured; detailed amendments not fully extracted | `src:usa-md:comar-29-06-01` |
| COMAR 09.20.01 | Plumbing / Fuel Gas | Library of Maryland Regulations | incorporation framework captured; edition reconciliation open | `src:usa-md:comar-09-20-01` |
| COMAR 09.15.05 | Mechanical / HVACR | Library of Maryland Regulations | incorporation framework captured; edition reconciliation open | `src:usa-md:comar-09-15-05` |
| COMAR 09.12.53 | Accessibility | Library of Maryland Regulations | applicable standards captured at high level; complete MAC amendments not extracted | `src:usa-md:comar-09-12-53` |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Summary | Practical Effect | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| `amendment:usa-md:mbps-local-enforcement` | Building / Residential | MBPS modifies model-code administration to identify local jurisdictions as responsible for implementation and enforcement, including plan review, permits, inspections, and occupancy certificates. | AHJ resolution must use local jurisdiction, not only the State adopting agency. | `src:usa-md:comar-09-12-51-04`; `src:usa-md:ps-12-505` | partially_verified |
| `amendment:usa-md:mbps-existing-building-mbrc` | Building / Residential | Existing buildings undergoing repair, alterations, additions, or change of occupancy are routed to the Maryland Building Rehabilitation Code. | Existing-building compliance must query MBRC / IEBC rather than only new-construction IBC / IRC provisions. | `src:usa-md:comar-09-12-51-04`; `src:usa-md:comar-09-12-58` | partially_verified |
| `amendment:usa-md:local-amendment-floors` | MBPS / Energy / Sprinklers / Wind | Local amendments may not weaken IECC / energy provisions, automatic fire sprinkler provisions, or wind design / wind-borne-debris requirements. | Local amendments need floor checks before being applied. | `src:usa-md:ps-12-504`; `src:usa-md:comar-09-12-51-05` | partially_verified |
| `amendment:usa-md:irc-appendices-radon-tiny-houses` | Residential | COMAR adopts IRC Appendices AF Radon Control Methods and AQ Tiny Houses as part of the IRC. | Residential code parser should include these appendices statewide unless locally superseded within allowed limits. | `src:usa-md:comar-09-12-51-04` | partially_verified |
| `amendment:usa-md:fire-baltimore-city-and-two-family-exclusions` | Fire - operational / prevention code | COMAR excludes Baltimore City except specified statutory scope and excludes buildings used solely as dwelling houses for not more than two families. | Fire-code AHJ and applicability routing must account for location and dwelling type. | `src:usa-md:comar-29-06-01` | partially_verified |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-md"
  model: "statewide_adoption_with_local_jurisdiction_implementation_and_enforcement"
  enforcing_entities:
    - "county by default for buildings or structures in the county, unless statutory municipal responsibility or negotiated arrangement applies"
    - "municipal corporation where it has elected or is required to implement and enforce under Public Safety Article §12-505"
    - "State Fire Marshal or legally designated local fire official for State Fire Prevention Code enforcement"
    - "trade-specific authorities and local officials for plumbing, mechanical, fuel gas, and electrical where applicable"
  minimum_required_activities:
    - "review and acceptance of appropriate plans"
    - "issuance of building permits"
    - "inspection of work authorized by building permits"
    - "issuance of appropriate use and occupancy certificates"
  state_reserved_activities:
    - "state adoption of MBPS, MBRC, MAC, MPC, IgCC, and related state code programs"
    - "State Fire Marshal / Fire Prevention Commission fire-code authority and appeal framework"
    - "state-administered modular / industrialized and state-owned building programs under the Model Performance Code"
  source_ids:
    - "src:usa-md:ps-12-505"
    - "src:usa-md:comar-09-12-51-08"
    - "src:usa-md:comar-29-06-01"
  verification_status: "partially_verified"
  confidence: 0.86
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-md"
  model: "local_amendments_allowed_with_state_floor_limits_and_department_submission"
  applies_to_code_families:
    - "MBPS / IBC"
    - "MBPS / IRC"
    - "MBPS / IECC"
    - "IgCC where locally adopted / amended"
  approval_required: "not established as state preapproval for ordinary MBPS local amendments in sources reviewed; local adoption must comply with law and submission rules"
  approving_authority_id: "local_jurisdiction_under_local_law; Department receives copies"
  filing_required: true
  filing_timing:
    ordinary: "at least 15 days before effective date"
    emergency: "within 5 days after emergency adoption"
  prohibited_weakened_areas:
    - "IECC / IBC Chapter 13 energy provisions"
    - "automatic fire sprinkler systems required by Standards, subject to statutory exception"
    - "wind design and wind-borne-debris provisions"
  registry_exists: "partial_repository_or_submission_process_known; complete registry not parsed"
  registry_source_ids:
    - "src:usa-md:labor-building-codes"
  legal_basis_source_ids:
    - "src:usa-md:ps-12-504"
    - "src:usa-md:comar-09-12-51-05"
  verification_status: "partially_verified"
  confidence: 0.84
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement and local amendment authority are related but distinct in Maryland. Local jurisdictions are responsible for MBPS implementation and enforcement, while local amendments are optional legal changes adopted under local authority and constrained by statewide floors. A jurisdiction can be the enforcing AHJ without necessarily having a unique local amendment for a code family. Production address resolution should therefore compute both an enforcing AHJ and a local-amendment overlay.

### 6.4 Known Local Amendment Registries

| Registry ID | Registry / Publication Path | Coverage | Parsed? | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- |
| `registry:usa-md:labor-local-ordinances` | Maryland Department of Labor Building Codes Administration local ordinances / code page references | local amendments / ordinances submitted to Department | no | `src:usa-md:labor-building-codes`; `src:usa-md:labor-building-admin` | Department page confirms local amendments / ordinances are part of the MBPS workflow, but local ordinance files were not individually extracted. |
| `registry:usa-md:local-code-websites` | county and municipal code websites | local building, fire, trade, zoning, and administrative amendments | no | none | Needs jurisdiction-by-jurisdiction crawl and normalization. |

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendments were parsed in this pass. This is an explicit gap; do not assume a local amendment does or does not exist for a specific Maryland county or municipality until the local ordinance source is checked.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: partial_for_statewide_rules; not_ready_for_address_level_AHJ_resolution.

Jurisdiction stack:

```text
Address
  -> State of Maryland
  -> County
  -> Municipality / incorporated town, if applicable
  -> MBPS local jurisdiction under Public Safety Article §12-505
  -> Local building department / building official or contracted enforcement entity
  -> State Fire Marshal or legally designated local fire official for State Fire Prevention Code scope
  -> Plumbing / HVACR / electrical trade AHJs and board rules, where applicable
  -> State-administered MPC program if project is modular / industrialized or State-owned / leased / operated / controlled
  -> Applicable statewide adoption records
  -> Applicable local amendment records
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | U.S. Census / state boundary source not selected in this pass | none | statewide | unknown | pending |
| County | U.S. Census TIGER/Line or Maryland GIS source not selected in this pass | none | statewide | unknown | pending |
| Municipality | Maryland municipal boundary source not selected in this pass | none | statewide | unknown | pending |
| Fire District | State Fire Marshal / local fire official boundary source not selected in this pass | none | partial / local | unknown | pending |
| Special District | not selected | none | unknown | unknown | pending |
| Local code jurisdiction | county / municipal code office and Labor local-ordinance sources | `src:usa-md:ps-12-505`; `src:usa-md:labor-building-codes` | statewide legal framework; local implementation data not loaded | variable | pending |

### 7.3 AHJ Contact Data

No AHJ contact records were populated in this pass. The report identifies the statewide and legal framework, but production routing still requires contact / office data for counties, municipalities, fire officials, and trade-specific AHJs.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Publisher / Custodian | Source Type | URL | Key Fields Supported | Access Date | Caveats |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `src:usa-md:labor-building-codes` | Maryland Building Codes - Building Codes Administration | Maryland Department of Labor | agency page | https://labor.maryland.gov/labor/build/buildcodes.shtml | MBPS summary; local amendment limits summary; MPC scope; code portal | 2026-06-25 | Agency summary page; use COMAR / statutes for controlling text. |
| `src:usa-md:labor-building-admin` | Building Codes Administration | Maryland Department of Labor | agency page | https://www.labor.maryland.gov/labor/build/ | Building Codes Administration mandates and list of state code programs | 2026-06-25 | High-level agency description; not a substitute for code text. |
| `src:usa-md:labor-2021-cycle` | Code Administration News / 2021 Code Cycle | Maryland Department of Labor | agency notice | https://labor.maryland.gov/labor/build/buildadmin.shtml | 2021 IBC / IRC / IECC / IgCC / IEBC adoption notice; 2023-05-29 effective date; 2024-05-29 local implementation date | 2026-06-25 | Page is a dated agency notice; later IgCC amendment supersedes 2021 IgCC row. |
| `src:usa-md:code-matrix-2026` | Current Adopted Building-Related Codes in the State of Maryland | Maryland Department of Labor | agency PDF matrix | https://labor.maryland.gov/labor/build/buildcodematrix.pdf | Current matrix rows for IBC, IRC, IECC, IEBC, IPMC, MAC, IPC, IMC, NEC, NFPA 1, NFPA 101, conveyance, fuel gas, LP gas | 2026-06-25 | PDF dated 2026-03-18; IgCC row appears superseded by COMAR amendment effective 2026-05-25; trade-code edition rows need reconciliation with current COMAR wording. |
| `src:usa-md:ps-12-501` | Public Safety Article §12-501 | Maryland General Assembly | statute | https://mgaleg.maryland.gov/mgawebsite/Laws/StatuteText?article=gps&enactments=false&section=12-501 | Definitions of Department, local jurisdiction, Standards, IBC / IECC / IgCC | 2026-06-25 | Official HTML statute text. |
| `src:usa-md:ps-12-503` | Public Safety Article §12-503 | Maryland General Assembly | statute | https://mgaleg.maryland.gov/mgawebsite/Laws/StatuteText?article=gps&enactments=false&section=12-503 | Department adoption of MBPS; IBC / IECC adoption; subsequent version cycle; IgCC adoption | 2026-06-25 | Official HTML statute text. |
| `src:usa-md:ps-12-504` | Public Safety Article §12-504 | Maryland General Assembly | statute | https://mgaleg.maryland.gov/mgawebsite/laws/StatuteText?article=gps&section=12-504 | Local amendment authority, amendment limits, local amendment submission timing | 2026-06-25 | Official HTML statute text. |
| `src:usa-md:ps-12-505` | Public Safety Article §12-505 | Maryland General Assembly | statute | https://mgaleg.maryland.gov/mgawebsite/Laws/StatuteText?archived=False&article=gps&enactments=False&section=12-505 | Local implementation and enforcement; 12-month local implementation; county / municipal responsibility; minimum enforcement activities | 2026-06-25 | Official HTML statute text. |
| `src:usa-md:comar-09-12-51` | COMAR 09.12.51 Maryland Building Performance Standards | Library of Maryland Regulations | regulation chapter | https://regs.maryland.gov/us/md/exec/comar/09.12.51 | MBPS chapter structure and authority | 2026-06-25 | Chapter page; detailed requirements in child regulations. |
| `src:usa-md:comar-09-12-51-04` | COMAR 09.12.51.04 Incorporation by Reference | Library of Maryland Regulations | regulation | https://regs.maryland.gov/us/md/exec/comar/09.12.51.04 | 2021 IBC / IRC / IECC incorporation; MBPS modifications; existing-building routing; IRC appendices | 2026-06-25 | Long HTML regulation; detailed amendment diff not fully parsed. |
| `src:usa-md:comar-09-12-51-05` | COMAR 09.12.51.05 Maryland Building Performance Standards | Library of Maryland Regulations | regulation | https://regs.maryland.gov/us/md/exec/comar/09.12.51.05 | MBPS composition; local amendments; local amendment limits and submission timing | 2026-06-25 | Official regulatory text. |
| `src:usa-md:comar-09-12-51-06` | COMAR 09.12.51.06 Application | Library of Maryland Regulations | regulation | https://regs.maryland.gov/us/md/exec/comar/09.12.51.06 | MBPS applicability to permit applications; local implementation within 12 months; agritourism exception | 2026-06-25 | Official regulatory text. |
| `src:usa-md:comar-09-12-51-08` | COMAR 09.12.51.08 Enforcement | Library of Maryland Regulations | regulation | https://regs.maryland.gov/us/md/exec/comar/09.12.51.08 | Enforcement responsibility of local jurisdiction | 2026-06-25 | Official regulatory text. |
| `src:usa-md:comar-09-12-51-09` | COMAR 09.12.51.09 State Fire Code | Library of Maryland Regulations | regulation | https://regs.maryland.gov/us/md/exec/comar/09.12.51.09 | Cross-reference to State Fire Code enforcement by State Fire Marshal or authorized fire official | 2026-06-25 | Official regulatory cross-reference; use COMAR 29.06.01 for fire code details. |
| `src:usa-md:comar-09-12-50-02-1` | COMAR 09.12.50.02-1 Incorporation by Reference | Library of Maryland Regulations | regulation | https://regs.maryland.gov/us/md/exec/comar/09.12.50.02-1 | Model Performance Code incorporations including 2021 IBC, 2021 IPC, 2020 NEC, 2021 IRC, 2021 IMC, 2021 IECC | 2026-06-25 | Applies to MPC scope; not ordinary local building-code scope unless adopted locally. |
| `src:usa-md:comar-09-12-57` | COMAR 09.12.57 International Green Construction Code | Library of Maryland Regulations | regulation chapter / regulation | https://regs.maryland.gov/us/md/exec/comar/09.12.57 | IgCC chapter history; 2024 IgCC incorporation effective 2026-05-25 | 2026-06-25 | Current regulation supersedes older matrix / agency notice for IgCC edition. |
| `src:usa-md:comar-09-12-58` | COMAR 09.12.58 Maryland Building Rehabilitation Code Regulations | Library of Maryland Regulations | regulation chapter / regulation | https://regs.maryland.gov/us/md/exec/comar/09.12.58 | 2021 IEBC incorporation and MBRC enforcement | 2026-06-25 | Official regulatory text; full MBRC amendment details not fully parsed. |
| `src:usa-md:comar-09-12-53` | COMAR 09.12.53 Maryland Accessibility Code | Library of Maryland Regulations | regulation chapter / regulation | https://regs.maryland.gov/us/md/exec/comar/09.12.53 | MAC authority and applicable standards | 2026-06-25 | Current matrix lists 2019 MAC; COMAR text identifies ADA / Maryland standards in detail. |
| `src:usa-md:comar-29-06-01` | COMAR 29.06.01 State Fire Prevention Code | Library of Maryland Regulations | regulation chapter | https://regs.maryland.gov/us/md/exec/comar/29.06/index.full.html | State Fire Prevention Code authority, 2024 NFPA 1 / NFPA 101 incorporation, enforcement, applicability, exclusions, new-building rule | 2026-06-25 | Long HTML full-chapter view; detailed NFPA amendments not fully extracted. |
| `src:usa-md:comar-09-20-01` | COMAR 09.20.01 State Plumbing Code | Library of Maryland Regulations | regulation chapter / regulation | https://regs.maryland.gov/us/md/exec/comar/09.20.01 | State Plumbing Code authority and incorporation framework for IPC, IRC, IFGC, NFPA 54, NFPA 58 | 2026-06-25 | Current COMAR uses "most recent version" wording; reconcile with 2026 agency matrix before verified production use. |
| `src:usa-md:comar-09-15-05` | COMAR 09.15.05 Maryland Building Performance Standards for HVACR Services | Library of Maryland Regulations | regulation chapter / regulation | https://regs.maryland.gov/us/md/exec/comar/09.15.05 | HVACR authority and incorporation framework for IMC / IRC HVACR chapters | 2026-06-25 | Current COMAR uses "most recent version" wording; reconcile with 2026 agency matrix before verified production use. |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| `src:usa-md:code-matrix-2026` | agency_matrix_pdf | The matrix is an official agency PDF dated 2026-03-18 and is valuable for cross-code family status, but it is not the controlling legal text and appears superseded for IgCC by a COMAR amendment effective 2026-05-25. | Use with COMAR / statutes; flag conflicts. |
| `src:usa-md:comar-09-20-01` | edition_reconciliation_needed | Current State Plumbing Code text adopts the most recent versions of several codes as amended / modified / updated, while the 2026 matrix lists specific editions. | Confirm with board rulemaking history before verified production use. |
| `src:usa-md:comar-09-15-05` | edition_reconciliation_needed | Current HVACR text adopts the most recent version of IMC / IRC HVACR chapters, while the 2026 matrix lists 2021 editions. | Confirm with board rulemaking history before verified production use. |
| `src:usa-md:comar-29-06-01` | long_regulation_partial_extraction | Fire code chapter contains detailed amendments to NFPA 1 and NFPA 101; only authority, incorporation, enforcement, applicability, and major exclusions were extracted. | Parse section-level amendments for code-content products. |
| `src:usa-md:labor-building-codes` | agency_summary | Labor's code page summarizes legal rules and links to COMAR; summaries can lag or omit legal nuance. | Treat as source registry / orientation; rely on statute and COMAR for controlling fields. |

### 8.3 Supplemental Sources

| Source ID | Title | Publisher | Source Type | URL | Used For | Caveat |
| --- | --- | --- | --- | --- | --- | --- |
| `src:usa-md:supp-ps-12-603-justia` | Maryland Public Safety Code §12-603 Electrical Installation to Conform to Code | Justia | unofficial statute copy | https://law.justia.com/codes/maryland/public-safety/title-12/subtitle-6/section-12-603/ | Electrical-code conformity language: NEC or county electrical code / amendments | Unofficial copy; included only because the official state page was not opened successfully in this pass. Verify against Maryland General Assembly before marking electrical code field verified. |

### 8.4 Source Extraction Metadata

| Extraction ID | Source ID | Extracted Fields | Extracted By | Extraction Date | Quality | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| `extract:usa-md:001` | `src:usa-md:labor-building-codes` | MBPS summary; local amendments; MPC scope | ChatGPT | 2026-06-25 | medium | Agency summary; cross-checked against COMAR and statutes. |
| `extract:usa-md:002` | `src:usa-md:labor-2021-cycle` | 2021 adoption notice; effective and local implementation dates | ChatGPT | 2026-06-25 | high | Dated agency notice. |
| `extract:usa-md:003` | `src:usa-md:code-matrix-2026` | cross-code family adoption matrix | ChatGPT | 2026-06-25 | medium | PDF table visually and textually reviewed; conflicts captured. |
| `extract:usa-md:004` | `src:usa-md:ps-12-503` | statewide adoption authority; subsequent version rule | ChatGPT | 2026-06-25 | high | Official MGA HTML. |
| `extract:usa-md:005` | `src:usa-md:ps-12-504` | local amendment authority and limits | ChatGPT | 2026-06-25 | high | Official MGA HTML. |
| `extract:usa-md:006` | `src:usa-md:ps-12-505` | local enforcement model and minimum enforcement activities | ChatGPT | 2026-06-25 | high | Official MGA HTML. |
| `extract:usa-md:007` | `src:usa-md:comar-09-12-51-04` | 2021 IBC / IRC / IECC incorporation and MBPS modifications | ChatGPT | 2026-06-25 | high | Long regulation; not a complete model-code diff. |
| `extract:usa-md:008` | `src:usa-md:comar-09-12-57` | 2024 IgCC incorporation | ChatGPT | 2026-06-25 | high | Current COMAR amendment. |
| `extract:usa-md:009` | `src:usa-md:comar-29-06-01` | 2024 NFPA 1 / 101, fire enforcement, applicability, exclusions, new-building rule | ChatGPT | 2026-06-25 | high | Detailed amendments not fully parsed. |
| `extract:usa-md:010` | `src:usa-md:comar-09-20-01` | plumbing / fuel gas incorporation framework | ChatGPT | 2026-06-25 | medium | Edition reconciliation open. |
| `extract:usa-md:011` | `src:usa-md:comar-09-15-05` | HVACR / mechanical incorporation framework | ChatGPT | 2026-06-25 | medium | Edition reconciliation open. |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| `report` | `report.status` | `partially_verified` | verified | 1.00 | none | Core authority, MBPS, fire, and local enforcement fields are sourced; several trade and local registry fields remain open. |
| `report` | `risk.overall_confidence` | 0.72 | verified | 1.00 | none | Confidence reflects strong MBPS / fire sources and weaker trade / local registry coverage. |
| `ahj:usa-md:labor-building-codes` | primary authority | Maryland Department of Labor / Building Codes Administration | verified | 0.90 | `src:usa-md:ps-12-503`; `src:usa-md:labor-building-admin`; `src:usa-md:comar-09-12-51` | Primary statewide MBPS authority. |
| `adoption:usa-md:mbps-2021-ibc` | edition | 2021 IBC | verified | 0.92 | `src:usa-md:comar-09-12-51-04`; `src:usa-md:code-matrix-2026` | Supported by COMAR and agency matrix. |
| `adoption:usa-md:mbps-2021-irc` | edition | 2021 IRC | verified | 0.92 | `src:usa-md:comar-09-12-51-04`; `src:usa-md:code-matrix-2026` | Supported by COMAR and agency matrix. |
| `adoption:usa-md:mbps-2021-iecc` | edition | 2021 IECC | verified | 0.92 | `src:usa-md:comar-09-12-51-04`; `src:usa-md:code-matrix-2026` | Supported by COMAR and agency matrix. |
| `date-rule:usa-md:mbps-local-implementation` | mandatory date | 2024-05-29 | verified | 0.86 | `src:usa-md:labor-2021-cycle`; `src:usa-md:ps-12-505` | Agency notice plus statute support local 12-month implementation rule. |
| `local-enforcement:usa-md` | model | statewide adoption with local implementation and enforcement | verified | 0.86 | `src:usa-md:ps-12-505`; `src:usa-md:comar-09-12-51-08` | Address-level AHJ contact data still missing. |
| `local-amendment-rule:usa-md` | model | local amendments allowed with floor limits and Department submission | verified | 0.84 | `src:usa-md:ps-12-504`; `src:usa-md:comar-09-12-51-05` | Local ordinance registry not parsed. |
| `adoption:usa-md:sfpc-2024-nfpa-1-101` | edition | NFPA 1 (2024) and NFPA 101 (2024) | verified | 0.90 | `src:usa-md:comar-29-06-01`; `src:usa-md:code-matrix-2026` | Detailed NFPA amendments not fully extracted. |
| `adoption:usa-md:igcc-2024` | edition | 2024 IgCC | verified | 0.85 | `src:usa-md:comar-09-12-57` | Current COMAR supersedes older matrix row. |
| `adoption:usa-md:plumbing-2021-ipc-matrix` | edition | 2021 IPC per matrix; current COMAR uses "most recent version" language | unresolved | 0.55 | `src:usa-md:comar-09-20-01`; `src:usa-md:code-matrix-2026` | Needs board confirmation. |
| `adoption:usa-md:mechanical-2021-imc-matrix` | edition | 2021 IMC per matrix; current COMAR uses "most recent version" language | unresolved | 0.55 | `src:usa-md:comar-09-15-05`; `src:usa-md:code-matrix-2026` | Needs board confirmation. |
| `adoption:usa-md:electrical-nec-hybrid` | edition / scope | 2017 NEC fire-code matrix row and 2020 NEC MPC row | unresolved | 0.50 | `src:usa-md:code-matrix-2026`; `src:usa-md:comar-09-12-50-02-1`; `src:usa-md:supp-ps-12-603-justia` | Need official state statute / fire-code electrical edition path and local amendment parse. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| Leftover template markers removed | pass | Placeholder-marker scan returned no matches after generation. |
| All source IDs resolve | pass | Each `src:usa-md:...` reference used in the body appears in the source registry. |
| All authority IDs resolve | pass | Authorities used in adoption records and graph edges are defined in Section 2 or explicitly introduced as local / trade AHJ placeholders. |
| Current code families have matrix rows | pass | Matrix rows retained for MBPS, MBRC, mechanical, plumbing, fuel gas, electrical, energy, green, fire, accessibility, elevator/conveyance, and livability. |
| Building and operational fire code are separated | pass | MBPS rows and State Fire Prevention Code rows are separate. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Unknown dates remain unknown; effective and local implementation dates are separate. |
| Effective dates are valid ISO dates | pass | Dates inserted in date fields use `YYYY-MM-DD`. |
| No impossible date sequences | pass | No row introduces an adoption / effective / mandatory sequence that contradicts verified sources. |
| Transition rules have explicit trigger conditions | pass | MBPS, local amendment, and fire new-condition trigger rules are separately captured. |
| Permit-date logic is captured where applicable | pass | MBPS permit-application applicability and fire permit-issuance / construction-start trigger are captured. |
| Local enforcement model classified | pass | Classified as statewide adoption with local implementation and enforcement. |
| Local amendment rule classified | pass | Classified as local amendments allowed with floor limits and Department submission. |
| AHJ confirmation metadata present | fail | Address-level AHJ contacts and local office records are not populated. |
| Trade-code edition conflicts captured | pass | Plumbing, fuel gas, mechanical, and electrical caveats are explicit. |
| Official-source caveats captured | pass | Section 8.2 captures caveats for agency matrix, trade-code edition reconciliation, and long regulatory chapters. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `issue:usa-md:001` | high | plumbing / fuel gas edition reconciliation | Current COMAR 09.20.01 uses "most recent version" incorporation language, while the 2026 agency matrix lists 2021 IPC, 2018 IFGC, and 2017 NFPA 58. | Review Board of Plumbing rulemaking, Maryland Register notices, and any Board guidance to determine current enforceable editions. | null | null | open |
| `issue:usa-md:002` | high | mechanical / HVACR edition reconciliation | Current COMAR 09.15.05 uses "most recent version" incorporation language, while the 2026 matrix lists 2021 IMC and 2021 IRC HVACR amendments. | Review HVACR Board rulemaking, Maryland Register notices, and guidance to determine current enforceable editions and transition rules. | null | null | open |
| `issue:usa-md:003` | high | electrical code scope | Electrical code authority is hybrid; this pass used the agency matrix and a supplemental statute copy, but did not fully parse official electrical statutes, State Fire Prevention Code electrical references, or local amendments. | Open official Public Safety Subtitle 6 pages, parse COMAR fire-code electrical amendments, and inventory county electrical codes. | null | null | open |
| `issue:usa-md:004` | medium | local amendment registry | Local amendments must be submitted to the Department, but local ordinance files were not extracted. | Crawl Maryland Labor local ordinances / amendment repository and county / municipal code sources; normalize by jurisdiction and code family. | null | null | open |
| `issue:usa-md:005` | medium | elevator / conveyance, safety glazing, and livability | Matrix rows were captured, but detailed statutes and regulations were not parsed. | Parse conveyance, safety glazing, and Minimum Livability Code chapters and add authority / adoption records. | null | null | open |
| `issue:usa-md:006` | medium | address-level AHJ contacts | Report does not include county, municipal, fire, plumbing, HVACR, or electrical contacts. | Build county / municipality / fire official contact registry and boundary data sources. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| `watch:usa-md:labor-code-matrix` | `src:usa-md:code-matrix-2026` | pdf_diff | monthly | matrix date changes, new editions, revised trade-code rows | 2026-06-25 | active |
| `watch:usa-md:mbps-comar` | `src:usa-md:comar-09-12-51` | html_diff | monthly | amendment to MBPS chapter or 2024 / later I-code incorporation | 2026-06-25 | active |
| `watch:usa-md:igcc-comar` | `src:usa-md:comar-09-12-57` | html_diff | monthly | IgCC edition, effective date, or local amendment rule changes | 2026-06-25 | active |
| `watch:usa-md:mbrc-comar` | `src:usa-md:comar-09-12-58` | html_diff | monthly | IEBC / MBRC edition or enforcement changes | 2026-06-25 | active |
| `watch:usa-md:fire-comar` | `src:usa-md:comar-29-06-01` | html_diff | monthly | NFPA edition, fire-code amendments, applicability, or enforcement changes | 2026-06-25 | active |
| `watch:usa-md:plumbing-comar` | `src:usa-md:comar-09-20-01` | html_diff | monthly | plumbing / fuel-gas incorporation wording, edition, or amendment changes | 2026-06-25 | active |
| `watch:usa-md:hvacr-comar` | `src:usa-md:comar-09-15-05` | html_diff | monthly | IMC / IRC HVACR incorporation wording, edition, or amendment changes | 2026-06-25 | active |
| `watch:usa-md:statutes-12-503-505` | `src:usa-md:ps-12-503` | statute_diff | quarterly | changes to statewide adoption, local amendment, or local enforcement statutes | 2026-06-25 | active |
| `watch:usa-md:local-ordinances` | `src:usa-md:labor-building-codes` | link_and_file_inventory | monthly | new local ordinance / amendment files or repository restructuring | 2026-06-25 | pending |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-24 | Generated baseline draft report | `report:usa-md` | none | Codex | Starting file contained state-specific frontmatter but unresolved placeholders. |
| 2026-06-25 | Populated Maryland report from official state statutes, COMAR chapters, Labor code pages, and current code matrix; status upgraded to `partially_verified`. | `report:usa-md`; `ahj:usa-md:labor-building-codes`; `adoption:usa-md:mbps-2021-ibc`; `adoption:usa-md:mbps-2021-irc`; `adoption:usa-md:mbps-2021-iecc`; `adoption:usa-md:sfpc-2024-nfpa-1-101` | `src:usa-md:ps-12-503`; `src:usa-md:ps-12-504`; `src:usa-md:ps-12-505`; `src:usa-md:comar-09-12-51-04`; `src:usa-md:comar-29-06-01`; `src:usa-md:code-matrix-2026` | ChatGPT | Core authority and adoption fields are source-backed; unresolved trade / local registry issues retained explicitly. |
