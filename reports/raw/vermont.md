---
state:
  state_id: "US-VT"
  name: "Vermont"
  abbreviation: "VT"
report:
  report_id: "state-report:usa-vt"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-26"
  last_verified: null
  reviewed_by: null
risk:
  overall_confidence: 0.62 # 0.00 - 1.00
  risk_flags:
    - "fire_building_primary_code_text_not_fully_parsed"
    - "fire_building_effective_date_unresolved"
    - "fire_building_model_editions_partly_supported_by_supplemental_sources"
    - "local_amendment_scope_partially_verified"
    - "elevator_conveyance_authority_unresolved"
    - "energy_code_effective_date_supported_by_guidance_and_statutory_transition_rules"
  open_questions_count: 7

---

# State Building Code Authority Report: Vermont

## 1. Executive Summary

- **Authority model:** Vermont uses a state-administered public-building fire and building safety model. The Commissioner of Public Safety, acting through the Department of Public Safety / Division of Fire Safety, is the primary statewide authority for rules governing the construction, maintenance, operation, fire prevention, and fire-hazard removal of public buildings. Owner-occupied single-family residences are generally outside the statutory "public building" definition unless used for specified public purposes. Source IDs: `src:usa-vt:stat-20-2731`, `src:usa-vt:stat-20-2730`.

- **Statewide code status:** The State has an adopted 2025 Vermont Fire & Building Safety Code rule record, but this pass did not parse the full state code text. Supplemental municipal and training sources indicate the 2025 code cycle updated key referenced construction and fire codes from 2015 editions to 2021 editions, including NFPA 1, NFPA 101, IBC, and IEBC. Electrical and plumbing are more directly supported by state rule records: the 2025 Vermont Electrical Safety Rules adopt the 2023 National Electrical Code, and the 2025 Vermont Plumbing Rules update the Vermont plumbing rules from the 2021 IPC to the 2024 IPC. Source IDs: `src:usa-vt:sos-rule-25p022-fire-building`, `src:usa-vt:burlington-adopted-codes`, `src:usa-vt:bennington-2025-code-release`, `src:usa-vt:sos-rule-25p023-electrical`, `src:usa-vt:sos-rule-25p024-plumbing`.

- **Local enforcement model:** Municipal enforcement is possible, but it is not the default statewide model. A municipality may appoint officials and establish procedures for enforcing rules made by the Commissioner only when the Commissioner assigns enforcement responsibility to that municipality; the assignment does not diminish the Commissioner's authority. Source ID: `src:usa-vt:stat-20-2736`.

- **Local amendment posture:** Municipalities may establish building codes, but municipal codes must be consistent with the rules adopted by the Commissioner of Public Safety, and municipal building codes may not govern electrical installations covered by Vermont's electrical-installation chapter. Copies of any municipal building code must be filed with the municipal building inspector and town clerk. Local enforcement and local code adoption are separate questions in Vermont. Source ID: `src:usa-vt:stat-24-3101`.

- **Known transition periods or pending changes:** Residential and commercial energy standards have explicit transition rules. The Commissioner of Public Service must set an effective date at least three months after adopting amended RBES or CBES; construction or applications before the effective date may use the prior or amended standard, while later work must comply with the most recent standard. Efficiency Vermont guidance identifies the latest RBES and CBES effective date as 2024-07-01 for construction commenced on or after that date. Fire/building, electrical, and plumbing adoption/effective/mandatory dates need full rule-text verification before production use. Source IDs: `src:usa-vt:stat-30-51-rbes`, `src:usa-vt:stat-30-53-cbes`, `src:usa-vt:efficiency-vt-energy-code-support`.

- **Production readiness:** partially_ready_for_narrow_review

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: Commissioner of Public Safety / Division of Fire Safety is the primary
    verified authority for statewide public-building fire and building safety rules.
  confidence: 0.85
  source_ids:
  - src:usa-vt:stat-20-2731
  - src:usa-vt:sos-rule-25p022-fire-building
- topic: Scope of public-building rules
  finding: The statutory public-building definition includes many nonresidential and
    public-use premises, while owner-occupied single-family residences are generally
    excluded unless used for public purposes listed in statute.
  confidence: 0.8
  source_ids:
  - src:usa-vt:stat-20-2730
- topic: Primary building/fire code cycle
  finding: A 2025 Vermont Fire & Building Safety Code rule record is adopted; precise
    model-code edition mapping still needs confirmation in the full state code text.
  confidence: 0.62
  source_ids:
  - src:usa-vt:sos-rule-25p022-fire-building
  - src:usa-vt:burlington-adopted-codes
  - src:usa-vt:bennington-2025-code-release
- topic: Electrical code authority and edition
  finding: Electricians' Licensing Board rules may adopt a nationally recognized electrical
    code; the 2025 Vermont Electrical Safety Rules adopt the 2023 NEC.
  confidence: 0.88
  source_ids:
  - src:usa-vt:stat-26-891-electrical
  - src:usa-vt:sos-rule-25p023-electrical
- topic: Plumbing code authority and edition
  finding: Plumber's Examining Board rules may adopt a nationally recognized plumbing
    code; the 2025 Vermont Plumbing Rules update the rules to the 2024 IPC.
  confidence: 0.86
  source_ids:
  - src:usa-vt:stat-26-2173-plumbing
  - src:usa-vt:sos-rule-25p024-plumbing
- topic: Energy code authority
  finding: Commissioner of Public Service adopts and updates RBES and CBES; latest
    support guidance identifies both 2024 standards as effective 2024-07-01.
  confidence: 0.82
  source_ids:
  - src:usa-vt:stat-30-51-rbes
  - src:usa-vt:stat-30-53-cbes
  - src:usa-vt:efficiency-vt-energy-code-support
- topic: Local enforcement
  finding: Municipal enforcement can occur by Commissioner's assignment, with Commissioner
    authority retained.
  confidence: 0.8
  source_ids:
  - src:usa-vt:stat-20-2736
- topic: Local amendments / local codes
  finding: Municipal building codes are allowed but must be consistent with Commissioner
    rules; electrical installations covered by state electrical law are excluded.
  confidence: 0.78
  source_ids:
  - src:usa-vt:stat-24-3101
- topic: Accessibility
  finding: Vermont has an Access Board rulemaking framework for access standards and
    Department of Public Safety enforcement responsibilities.
  confidence: 0.72
  source_ids:
  - src:usa-vt:stat-20-174-accessibility
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | `ahj:usa-vt:dps-commissioner-fire-safety` |
| Authority name | Commissioner of Public Safety / Department of Public Safety, Division of Fire Safety |
| Authority type | state executive agency / commissioner |
| Legal basis | 20 V.S.A. § 2731 authorizes rules governing construction, maintenance, operation, fire prevention, and fire-hazard removal for public-safety protection. |
| Role | Adopts and administers public-building fire and building safety rules; conducts or oversees inspections; maintains statewide construction-category/code applicability chart. |
| Enforcement model | State-administered inspection and enforcement, with municipal enforcement assignment possible under 20 V.S.A. § 2736. |
| Source IDs | `src:usa-vt:stat-20-2731`, `src:usa-vt:stat-20-2730`, `src:usa-vt:stat-20-2736`, `src:usa-vt:sos-rule-25p022-fire-building` |
| Verification status | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | `ahj:usa-vt:dps-commissioner-fire-safety` | Commissioner of Public Safety / Division of Fire Safety | Adopts and administers public-building construction and safety rules. | 20 V.S.A. § 2731; Rule 25P022 | `src:usa-vt:stat-20-2731`, `src:usa-vt:sos-rule-25p022-fire-building` | partially_verified |
| Residential | `ahj:usa-vt:public-service-commissioner-energy`; `ahj:usa-vt:dps-commissioner-fire-safety` where a dwelling is a public building | Commissioner of Public Service for RBES; Commissioner of Public Safety for public-building uses | RBES applies to residential energy construction; general owner-occupied single-family building-code authority was not established in this pass. | 30 V.S.A. § 51; 20 V.S.A. § 2730 | `src:usa-vt:stat-30-51-rbes`, `src:usa-vt:stat-20-2730`, `src:usa-vt:efficiency-vt-energy-code-support` | partially_verified |
| Existing Building / Rehabilitation | `ahj:usa-vt:dps-commissioner-fire-safety` | Commissioner of Public Safety / Division of Fire Safety | Administers existing-building/public-building safety standards through fire/building code framework. | 20 V.S.A. § 2731; Rule 25P022 | `src:usa-vt:stat-20-2731`, `src:usa-vt:sos-rule-25p022-fire-building`, `src:usa-vt:burlington-adopted-codes` | partially_verified |
| Mechanical | `ahj:usa-vt:dps-commissioner-fire-safety` | Commissioner of Public Safety / Division of Fire Safety | Mechanical requirements appear within the fire/building safety framework; exact referenced mechanical standards were not fully parsed. | 20 V.S.A. § 2731; Rule 25P022 | `src:usa-vt:stat-20-2731`, `src:usa-vt:sos-rule-25p022-fire-building`, `src:usa-vt:burlington-adopted-codes` | partially_verified |
| Plumbing | `ahj:usa-vt:plumbers-examining-board` | Plumber's Examining Board / Department of Public Safety | Makes and revises plumbing rules; may adopt a nationally recognized plumbing code. | 26 V.S.A. § 2173; Rule 25P024 | `src:usa-vt:stat-26-2173-plumbing`, `src:usa-vt:sos-rule-25p024-plumbing` | partially_verified |
| Fuel Gas | `ahj:usa-vt:dps-commissioner-fire-safety` | Commissioner of Public Safety / Division of Fire Safety | Fuel-gas references appear to be handled through fire/building safety referenced standards; exact edition mapping remains unresolved. | 20 V.S.A. § 2731; Rule 25P022 | `src:usa-vt:stat-20-2731`, `src:usa-vt:sos-rule-25p022-fire-building`, `src:usa-vt:burlington-adopted-codes` | partially_verified |
| Electrical | `ahj:usa-vt:electricians-licensing-board` | Electricians' Licensing Board / Department of Public Safety | Adopts rules for electrical installations and may adopt a nationally recognized electrical code. | 26 V.S.A. § 891; Rule 25P023 | `src:usa-vt:stat-26-891-electrical`, `src:usa-vt:sos-rule-25p023-electrical` | verified |
| Energy | `ahj:usa-vt:public-service-commissioner-energy` | Commissioner of Public Service / Department of Public Service | Adopts and updates RBES and CBES. | 30 V.S.A. §§ 51, 53 | `src:usa-vt:stat-30-51-rbes`, `src:usa-vt:stat-30-53-cbes`, `src:usa-vt:efficiency-vt-energy-code-support` | partially_verified |
| Fire - construction references | `ahj:usa-vt:dps-commissioner-fire-safety` | Commissioner of Public Safety / Division of Fire Safety | Adopts public-building fire and building safety construction references. | 20 V.S.A. § 2731; Rule 25P022 | `src:usa-vt:stat-20-2731`, `src:usa-vt:sos-rule-25p022-fire-building` | partially_verified |
| Fire - operational / prevention code | `ahj:usa-vt:dps-commissioner-fire-safety` | Commissioner of Public Safety / Division of Fire Safety | Adopts rules for fire prevention and removal of fire hazards. | 20 V.S.A. § 2731; Rule 25P022 | `src:usa-vt:stat-20-2731`, `src:usa-vt:sos-rule-25p022-fire-building` | partially_verified |
| Accessibility | `ahj:usa-vt:access-board` | Vermont Access Board; Department of Public Safety enforcement role | Adopts access rules; Department of Public Safety has enforcement responsibilities for covered public buildings. | 20 V.S.A. ch. 174 | `src:usa-vt:stat-20-174-accessibility`, `src:usa-vt:burlington-adopted-codes` | partially_verified |
| Elevator / Conveyance | `ahj:usa-vt:unresolved-elevator-conveyance` | unresolved | Elevator/conveyance authority was not resolved in this pass. | unresolved | none | unresolved |

### 2.3 Authority Hierarchy Notes

Vermont's primary verified statewide construction-safety authority is the Commissioner of Public Safety / Division of Fire Safety for public buildings. The public-building definition is central to applicability because owner-occupied single-family residences are generally excluded unless used for public purposes described by statute. Municipalities may enforce Commissioner rules only by assignment, and that assignment does not reduce the Commissioner's authority.

Energy standards are separately administered by the Commissioner of Public Service. RBES applies to residential energy construction, and CBES applies to commercial construction and residential buildings four stories or more, subject to statutory and rule-based scoping. Electrical and plumbing are administered through separate professional boards under Title 26. Municipalities retain local building-code authority, but local building codes must be consistent with Commissioner rules and cannot cover electrical installations governed by the state electrical chapter.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| `edge:usa-vt:001` | `ahj:usa-vt:dps-commissioner-fire-safety` | adopts_rules_for | public-building construction, maintenance, operation, fire prevention, and fire-hazard removal | `src:usa-vt:stat-20-2731` | verified |
| `edge:usa-vt:002` | `ahj:usa-vt:dps-commissioner-fire-safety` | may_assign_enforcement_to | municipalities | `src:usa-vt:stat-20-2736` | verified |
| `edge:usa-vt:003` | `ahj:usa-vt:dps-commissioner-fire-safety` | retains_authority_after_assignment | assigned municipal enforcement programs | `src:usa-vt:stat-20-2736` | verified |
| `edge:usa-vt:004` | municipalities | may_establish_consistent_building_codes | local building-code ordinances | `src:usa-vt:stat-24-3101` | partially_verified |
| `edge:usa-vt:005` | municipalities | may_not_regulate_under_local_building_code | electrical installations covered by 26 V.S.A. ch. 15 | `src:usa-vt:stat-24-3101` | verified |
| `edge:usa-vt:006` | `ahj:usa-vt:public-service-commissioner-energy` | adopts_and_updates | RBES and CBES | `src:usa-vt:stat-30-51-rbes`, `src:usa-vt:stat-30-53-cbes` | verified |
| `edge:usa-vt:007` | `ahj:usa-vt:electricians-licensing-board` | adopts_rules_for | electrical installations / NEC | `src:usa-vt:stat-26-891-electrical`, `src:usa-vt:sos-rule-25p023-electrical` | verified |
| `edge:usa-vt:008` | `ahj:usa-vt:plumbers-examining-board` | adopts_rules_for | plumbing systems / IPC | `src:usa-vt:stat-26-2173-plumbing`, `src:usa-vt:sos-rule-25p024-plumbing` | verified |
| `edge:usa-vt:009` | `ahj:usa-vt:access-board` | adopts_rules_for | access to public buildings | `src:usa-vt:stat-20-174-accessibility` | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | 2025 Vermont Fire & Building Safety Code | International Building Code | 2021, pending full state-code text confirmation | adopted; edition partially_verified | 2025-11-04 | null | null | null | Rule record adopted; implementation/effective-date text not parsed. | `src:usa-vt:sos-rule-25p022-fire-building`, `src:usa-vt:bennington-2025-code-release`, `src:usa-vt:burlington-adopted-codes` |
| Residential | No general statewide owner-occupied single-family building code confirmed; RBES applies for residential energy | RBES; public-building rules where dwelling use falls within public-building definition | 2024 RBES for energy; general building edition unresolved | partially_verified | null | 2024-07-01 for latest RBES | 2024-07-01 for latest RBES | 2024-07-01 for latest RBES | RBES applies to construction commenced on or after the effective date; general residential building-code trigger unresolved. | `src:usa-vt:stat-20-2730`, `src:usa-vt:stat-30-51-rbes`, `src:usa-vt:efficiency-vt-energy-code-support` |
| Existing Building / Rehabilitation | 2025 Vermont Fire & Building Safety Code | International Existing Building Code | 2021, pending full state-code text confirmation | adopted; edition partially_verified | 2025-11-04 | null | null | null | Rule record adopted; implementation/effective-date text not parsed. | `src:usa-vt:sos-rule-25p022-fire-building`, `src:usa-vt:bennington-2025-code-release`, `src:usa-vt:burlington-adopted-codes` |
| Mechanical | 2025 Vermont Fire & Building Safety Code | Mechanical and fire-safety referenced standards | varies; exact editions unresolved | partially_verified | 2025-11-04 | null | null | null | Referenced-standard mapping and transition dates unresolved. | `src:usa-vt:sos-rule-25p022-fire-building`, `src:usa-vt:burlington-adopted-codes` |
| Plumbing | 2025 Vermont Plumbing Rules | International Plumbing Code | 2024 | adopted | null | null | null | null | Rule record adopted; effective/mandatory dates not parsed. | `src:usa-vt:sos-rule-25p024-plumbing`, `src:usa-vt:stat-26-2173-plumbing` |
| Fuel Gas | 2025 Vermont Fire & Building Safety Code | NFPA fuel-gas / LP-gas referenced standards | edition unresolved | partially_verified | 2025-11-04 | null | null | null | Referenced-standard mapping and transition dates unresolved. | `src:usa-vt:sos-rule-25p022-fire-building`, `src:usa-vt:burlington-adopted-codes` |
| Electrical | 2025 Vermont Electrical Safety Rules | NFPA 70, National Electrical Code | 2023 | adopted | null | null | null | null | Rule record adopted; effective/mandatory dates not parsed. | `src:usa-vt:sos-rule-25p023-electrical`, `src:usa-vt:stat-26-891-electrical` |
| Energy | 2024 Vermont Residential Building Energy Standards and 2024 Vermont Commercial Building Energy Standards | RBES: Vermont standards based on prior RBES and IECC elements; CBES: Vermont standards based on IECC/ASHRAE 90.1 elements | 2024 Vermont standards | current | null | 2024-07-01 | 2024-07-01 | 2024-07-01 | Amended RBES/CBES effective date must be at least three months after adoption; pre-effective construction/applications may use prior or amended standards. | `src:usa-vt:stat-30-51-rbes`, `src:usa-vt:stat-30-53-cbes`, `src:usa-vt:efficiency-vt-energy-code-support`, `src:usa-vt:cvr-31-000-004-rbes`, `src:usa-vt:cvr-31-000-003-cbes` |
| Fire - construction references | 2025 Vermont Fire & Building Safety Code | NFPA 1, NFPA 101, IBC, and related construction standards | 2021 for key codes, pending full state-code text confirmation | adopted; edition partially_verified | 2025-11-04 | null | null | null | Rule record adopted; implementation/effective-date text not parsed. | `src:usa-vt:sos-rule-25p022-fire-building`, `src:usa-vt:bennington-2025-code-release`, `src:usa-vt:burlington-adopted-codes` |
| Fire - operational / prevention code | 2025 Vermont Fire & Building Safety Code | NFPA 1 Fire Code and NFPA 101 Life Safety Code | 2021, pending full state-code text confirmation | adopted; edition partially_verified | 2025-11-04 | null | null | null | Rule record adopted; implementation/effective-date text not parsed. | `src:usa-vt:sos-rule-25p022-fire-building`, `src:usa-vt:burlington-adopted-codes` |
| Accessibility | Vermont Access Rules / access standards under 20 V.S.A. ch. 174 | ADAAG / state access rules | 2012 Vermont Access Rules cited by municipal code list; full state access-rule text not parsed | partially_verified | null | null | null | null | Work commencing on public buildings must meet or exceed statutory access standards; detailed edition/date unresolved. | `src:usa-vt:stat-20-174-accessibility`, `src:usa-vt:burlington-adopted-codes` |
| Elevator / Conveyance | unresolved | unresolved | unresolved | unresolved | null | null | null | null | Authority and adoption status unresolved. | none |

### 3.2 Adoption Records

#### `adoption:usa-vt:fire-building-2025`

| Field | Value |
| --- | --- |
| Authority ID | `ahj:usa-vt:dps-commissioner-fire-safety` |
| State code name | 2025 Vermont Fire & Building Safety Code |
| Code families | Building; Existing Building / Rehabilitation; Mechanical where incorporated; Fuel Gas where incorporated; Fire - construction references; Fire - operational / prevention code |
| Base model codes | NFPA 1, NFPA 101, IBC, IEBC, and other referenced standards; exact full list must be confirmed in the full state code text. |
| Edition | 2021 for key NFPA/IBC/IEBC references, based on supplemental municipal/training sources; full state-code text confirmation required. |
| Adoption date | 2025-11-04, based on municipal notice; official SOS page confirms adopted status but not the adoption-date field in the extracted text. |
| Effective date | null |
| Operative date | null |
| Mandatory date | null |
| Status | adopted; edition/date details partially_verified |
| Source IDs | `src:usa-vt:sos-rule-25p022-fire-building`, `src:usa-vt:bennington-2025-code-release`, `src:usa-vt:burlington-adopted-codes`, `src:usa-vt:efficiency-vt-2026-fire-code-training` |

#### `adoption:usa-vt:electrical-2025`

| Field | Value |
| --- | --- |
| Authority ID | `ahj:usa-vt:electricians-licensing-board` |
| State code name | 2025 Vermont Electrical Safety Rules |
| Code family | Electrical |
| Base model code | NFPA 70, National Electrical Code |
| Edition | 2023 |
| Adoption date | null |
| Effective date | null |
| Operative date | null |
| Mandatory date | null |
| Status | adopted |
| Source IDs | `src:usa-vt:sos-rule-25p023-electrical`, `src:usa-vt:stat-26-891-electrical` |

#### `adoption:usa-vt:plumbing-2025`

| Field | Value |
| --- | --- |
| Authority ID | `ahj:usa-vt:plumbers-examining-board` |
| State code name | 2025 Vermont Plumbing Rules |
| Code family | Plumbing |
| Base model code | International Plumbing Code |
| Edition | 2024 |
| Adoption date | null |
| Effective date | null |
| Operative date | null |
| Mandatory date | null |
| Status | adopted |
| Source IDs | `src:usa-vt:sos-rule-25p024-plumbing`, `src:usa-vt:stat-26-2173-plumbing` |

#### `adoption:usa-vt:rbes-2024`

| Field | Value |
| --- | --- |
| Authority ID | `ahj:usa-vt:public-service-commissioner-energy` |
| State code name | 2024 Vermont Residential Building Energy Standards |
| Code family | Energy / residential |
| Base model code | Vermont RBES, based on prior Vermont RBES and IECC elements |
| Edition | 2024 Vermont standard |
| Adoption date | null |
| Effective date | 2024-07-01 |
| Operative date | 2024-07-01 |
| Mandatory date | 2024-07-01 |
| Status | current |
| Source IDs | `src:usa-vt:stat-30-51-rbes`, `src:usa-vt:efficiency-vt-energy-code-support`, `src:usa-vt:cvr-31-000-004-rbes` |

#### `adoption:usa-vt:cbes-2024`

| Field | Value |
| --- | --- |
| Authority ID | `ahj:usa-vt:public-service-commissioner-energy` |
| State code name | 2024 Vermont Commercial Building Energy Standards |
| Code family | Energy / commercial |
| Base model code | Vermont CBES, based on prior Vermont CBES, IECC elements, and ASHRAE 90.1-2019 elements |
| Edition | 2024 Vermont standard |
| Adoption date | null |
| Effective date | 2024-07-01 |
| Operative date | 2024-07-01 |
| Mandatory date | 2024-07-01 |
| Status | current |
| Source IDs | `src:usa-vt:stat-30-53-cbes`, `src:usa-vt:efficiency-vt-energy-code-support`, `src:usa-vt:cvr-31-000-003-cbes` |

#### `adoption:usa-vt:accessibility`

| Field | Value |
| --- | --- |
| Authority ID | `ahj:usa-vt:access-board` |
| State code name | Vermont access standards / Vermont Access Rules |
| Code family | Accessibility |
| Base model code | ADAAG / state access rules |
| Edition | State rule edition unresolved; municipal code list identifies 2012 Vermont Access Rules |
| Adoption date | null |
| Effective date | null |
| Operative date | null |
| Mandatory date | null |
| Status | partially_verified |
| Source IDs | `src:usa-vt:stat-20-174-accessibility`, `src:usa-vt:burlington-adopted-codes` |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

The only well-supported transition logic in this pass is for energy standards. RBES and CBES statutes distinguish adoption from effective date and require amended standards to become effective no sooner than three months after adoption. They also preserve a prior-standard option for projects started or submitted before the new standard's effective date. The 2024 RBES and 2024 CBES effective date is recorded as 2024-07-01 from Efficiency Vermont support guidance and should be confirmed against the official final rule text before marking verified.

Fire/building, electrical, plumbing, and accessibility rule dates remain unresolved because this pass reviewed official rule-record summaries but did not extract the full final rule text or codified rule effective-date sections.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `date-rule:usa-vt:rbes-transition` | RBES amendments | statutory minimum delayed effective date and transition option | Effective date must be specified by rule and be at least three months after adoption. | Residential construction commenced before the effective date of amended RBES. | yes, the earlier or amended standards may be used before the effective date; after the effective date, latest standards apply. | `src:usa-vt:stat-30-51-rbes` | verified |
| `date-rule:usa-vt:cbes-transition` | CBES amendments | statutory minimum delayed effective date and transition option | Effective date must be specified by rule and be at least three months after adoption. | Commercial construction permit application, plan approval application, or plan approval submitted before the effective date of amended CBES. | yes, the earlier or amended standards may be used before the effective date; after the effective date, latest standards apply. | `src:usa-vt:stat-30-53-cbes` | verified |
| `date-rule:usa-vt:rbes-2024-effective` | 2024 RBES | effective / operative / mandatory date | 2024-07-01 | Construction commenced on or after the effective date, according to support guidance. | unresolved in guidance summary; statutory transition rule governs pre-effective work. | `src:usa-vt:efficiency-vt-energy-code-support`, `src:usa-vt:stat-30-51-rbes` | partially_verified |
| `date-rule:usa-vt:cbes-2024-effective` | 2024 CBES | effective / operative / mandatory date | 2024-07-01 | Construction commenced on or after the effective date, according to support guidance; statute separately references permit/plan-approval submission triggers. | unresolved in guidance summary; statutory transition rule governs pre-effective submissions. | `src:usa-vt:efficiency-vt-energy-code-support`, `src:usa-vt:stat-30-53-cbes` | partially_verified |
| `date-rule:usa-vt:fire-building-2025` | 2025 Vermont Fire & Building Safety Code | unresolved | null | unresolved | unresolved | `src:usa-vt:sos-rule-25p022-fire-building` | unresolved |
| `date-rule:usa-vt:electrical-2025` | 2025 Vermont Electrical Safety Rules | unresolved | null | unresolved | unresolved | `src:usa-vt:sos-rule-25p023-electrical` | unresolved |
| `date-rule:usa-vt:plumbing-2025` | 2025 Vermont Plumbing Rules | unresolved | null | unresolved | unresolved | `src:usa-vt:sos-rule-25p024-plumbing` | unresolved |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Fire / Building | Full 2025 Vermont Fire & Building Safety Code final text | null | 2025-11-04 | null | null | null | monitor_for_full_text_and_effective_dates | `src:usa-vt:sos-rule-25p022-fire-building`, `src:usa-vt:bennington-2025-code-release` | Official SOS rule record is adopted, but the full code text and date mechanics were not parsed. |
| Electrical | 2025 Vermont Electrical Safety Rules final text | null | null | null | null | null | monitor_for_full_text_and_effective_dates | `src:usa-vt:sos-rule-25p023-electrical` | NEC edition is verified from the SOS rule summary; rule effective-date mechanics unresolved. |
| Plumbing | 2025 Vermont Plumbing Rules final text | null | null | null | null | null | monitor_for_full_text_and_effective_dates | `src:usa-vt:sos-rule-25p024-plumbing` | IPC edition is verified from the SOS rule summary; rule effective-date mechanics unresolved. |
| Energy | Potential 2026 CBES amendments | null | null | null | null | null | monitor | none | A future CBES-amendment filing should be checked in the Vermont Secretary of State / Legislative Committee on Administrative Rules materials before adding a source-backed adoption record. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| `applicability-rule:usa-vt:public-building-scope` | Fire / Building / Public-building safety | Public buildings, public-use premises, and covered dwelling uses | Whether the premises falls within 20 V.S.A. § 2730 definitions | Public-building status is the first-level applicability filter for Vermont's public-building safety framework. | `src:usa-vt:stat-20-2730` | verified |
| `applicability-rule:usa-vt:rbes-scope` | Energy / residential | Residential buildings and residential buildings three stories or less, subject to RBES scoping | Construction commenced on or after standard effective date | RBES is the residential energy code framework; exact exemptions and detailed scope should be parsed from final RBES text. | `src:usa-vt:stat-30-51-rbes`, `src:usa-vt:efficiency-vt-energy-code-support`, `src:usa-vt:cvr-31-000-004-rbes` | partially_verified |
| `applicability-rule:usa-vt:cbes-scope` | Energy / commercial | Commercial buildings and residential buildings four stories or greater, subject to CBES scoping | Permit or plan approval submission date under statute; construction-start trigger also appears in support guidance | CBES governs commercial energy construction and larger residential buildings; statutory and guidance triggers should be reconciled before verified status. | `src:usa-vt:stat-30-53-cbes`, `src:usa-vt:efficiency-vt-energy-code-support`, `src:usa-vt:cvr-31-000-003-cbes` | partially_verified |
| `applicability-rule:usa-vt:plumbing-scope` | Plumbing | Cities, villages, and towns with public water/sewer systems and public buildings with plumbing/water treatment/heating specialties | Plumbing work within statutory scope | Plumbing rules apply within the scope stated in 26 V.S.A. § 2173. | `src:usa-vt:stat-26-2173-plumbing` | verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** state-amended model-code framework with separate trade and energy rulemaking. Vermont's public-building fire/building safety code is a state rule that references and amends national model codes. Energy standards are Vermont-specific RBES and CBES standards based on IECC/ASHRAE elements. Electrical and plumbing rules adopt national model-code editions with Vermont rule overlays.

**Where amendments are published:** Vermont Secretary of State rule records and agency rule/code publications are the controlling publication path for adopted rules. The official final fire/building, electrical, plumbing, RBES, and CBES texts should be obtained from the agency or Secretary of State before a verified production import. Cornell/Justia regulatory mirrors and municipal adopted-code lists are useful extraction aids but not authoritative final text.

**Amendment parsing status:** partial. Core authority and several model-code editions are captured; full amendments are not parsed.

### 5.2 State Amendment Sources

| Amendment Source ID | Code Family | Publication Path | Parsed? | Notes |
| --- | --- | --- | --- | --- |
| `src:usa-vt:sos-rule-25p022-fire-building` | Fire / Building | Vermont Secretary of State Rules Service rule record | no_full_text | Official adopted rule record; summary only in this pass. |
| `src:usa-vt:sos-rule-25p023-electrical` | Electrical | Vermont Secretary of State Rules Service rule record | no_full_text | Official adopted rule record identifying 2023 NEC adoption. |
| `src:usa-vt:sos-rule-25p024-plumbing` | Plumbing | Vermont Secretary of State Rules Service rule record | no_full_text | Official adopted rule record identifying update from 2021 IPC to 2024 IPC. |
| `src:usa-vt:cvr-31-000-004-rbes` | Energy / RBES | Code of Vermont Rules mirror | partially | Unofficial mirror used to identify 2024 RBES basis and scope. |
| `src:usa-vt:cvr-31-000-003-cbes` | Energy / CBES | Code of Vermont Rules mirror | partially | Unofficial mirror used to identify 2024 CBES basis and scope. |
| `src:usa-vt:burlington-adopted-codes` | Supplemental model-edition list | Municipal adopted-code page | partially | Municipal page used only as a supplemental check for edition lists. |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Topic | Rule Summary | Source IDs | Confidence | Status |
| --- | --- | --- | --- | --- | --- | --- |
| `amendment:usa-vt:fire-building-2025-code-cycle` | Fire / Building | 2025 code-cycle update | Adopted rule establishes minimum standards for fire, explosion, hazardous materials, dangerous structural conditions, and carbon monoxide protection in public buildings; supplemental sources indicate key referenced codes moved to 2021 editions. | `src:usa-vt:sos-rule-25p022-fire-building`, `src:usa-vt:efficiency-vt-2026-fire-code-training`, `src:usa-vt:burlington-adopted-codes` | 0.60 | partially_verified |
| `amendment:usa-vt:electrical-2023-nec` | Electrical | NEC adoption | 2025 Vermont Electrical Safety Rules adopt the 2023 NEC. | `src:usa-vt:sos-rule-25p023-electrical` | 0.88 | verified |
| `amendment:usa-vt:plumbing-2024-ipc` | Plumbing | IPC adoption | 2025 Vermont Plumbing Rules update the Vermont rules from the 2021 IPC to the 2024 IPC. | `src:usa-vt:sos-rule-25p024-plumbing` | 0.86 | verified |
| `amendment:usa-vt:energy-2024-rbes-cbes` | Energy | 2024 RBES and CBES | Latest RBES and CBES support guidance identifies a 2024 standard effective date of 2024-07-01; statutory transition rules govern amended standards. | `src:usa-vt:stat-30-51-rbes`, `src:usa-vt:stat-30-53-cbes`, `src:usa-vt:efficiency-vt-energy-code-support` | 0.78 | partially_verified |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-vt"
  model: "state_administered_with_commissioner_assigned_municipal_enforcement"
  enforcing_entities:
    - "Department of Public Safety / Division of Fire Safety"
    - "municipal officials where enforcement responsibility is assigned by the Commissioner of Public Safety"
    - "local boards of health and Commissioner of Public Safety for plumbing-rule enforcement within the scope of 26 V.S.A. § 2173"
    - "municipal certification and recordation participants for energy certificates, where applicable"
  required_officials:
    - "municipal officials appointed by the municipality when enforcement responsibility is assigned"
    - "building inspector and town clerk for filed municipal building-code copies"
  state_reserved_activities:
    - "Commissioner authority remains intact after municipal assignment"
    - "state rule adoption and code-update responsibility"
    - "state inspections and public-building safety oversight"
  source_ids:
    - "src:usa-vt:stat-20-2736"
    - "src:usa-vt:stat-20-2731"
    - "src:usa-vt:stat-26-2173-plumbing"
    - "src:usa-vt:stat-24-3101"
  verification_status: "partially_verified"
  confidence: 0.80
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-vt"
  model: "municipal_building_codes_allowed_if_consistent_with_state_rules_with_electrical_exclusion"
  applies_to_code_families:
    - "Building"
    - "local building regulations"
    - "RBES stretch code where adopted by municipalities, subject to energy-code rules"
  approval_required: null
  approving_authority_id: null
  filing_required: true
  registry_exists: null
  registry_source_ids: []
  legal_basis_source_ids:
    - "src:usa-vt:stat-24-3101"
    - "src:usa-vt:cvr-31-000-004-rbes"
  verification_status: "partially_verified"
  confidence: 0.70
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement and local amendment authority should not be collapsed in Vermont. A municipality can receive an enforcement assignment from the Commissioner for state rules under 20 V.S.A. § 2736. Separately, a municipality may establish building codes under 24 V.S.A. § 3101, but those codes must be consistent with Commissioner rules and cannot govern electrical installations covered by state electrical law. The approval path for local amendments, any statewide registry, and the practical scope of consistency review remain unresolved.

### 6.4 Known Local Amendment Registries

| Registry ID | Code Family | Registry Name | Maintainer | Source IDs | Status | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| `registry:usa-vt:municipal-building-codes` | Local building codes | no statewide registry verified | unresolved | none | unresolved | 24 V.S.A. § 3101 requires local filing with the building inspector and town clerk, but a statewide registry was not identified in this pass. |
| `registry:usa-vt:rbes-stretch-code` | Energy / RBES stretch | no statewide local-adoption registry verified | unresolved | `src:usa-vt:cvr-31-000-004-rbes` | unresolved | RBES stretch-code adoption by municipalities is identified in mirrored rule text, but local-adoption registry status was not verified. |

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendments were parsed. Burlington and Bennington municipal pages were used only as supplemental source checks for state code-cycle details, not as extracted local amendment records.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: partial_concept_only

Jurisdiction stack:

```text
Address
  -> State of Vermont
  -> Municipality
  -> Public-building applicability under 20 V.S.A. § 2730
  -> Commissioner of Public Safety / Division of Fire Safety jurisdiction
  -> Municipal enforcement assignment status under 20 V.S.A. § 2736
  -> Municipal building-code ordinance status under 24 V.S.A. § 3101
  -> Energy standard applicability: RBES or CBES
  -> Trade-specific AHJs: electrical, plumbing, accessibility, and unresolved elevator/conveyance
  -> Applicable state adoption records
  -> Applicable local code or amendment records
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

No AHJ contact dataset was populated. Future extraction should prioritize Division of Fire Safety regional offices, municipal assignment records, municipal building officials, electrical inspection contacts, plumbing inspection contacts, and energy-code support contacts.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Source Type | Publisher / Maintainer | URL | Date / Version | Used For |
| --- | --- | --- | --- | --- | --- | --- |
| `src:usa-vt:stat-20-2731` | 20 V.S.A. § 2731, Rules; inspections; variances | statute | Vermont General Assembly / Vermont Laws Online | https://legislature.vermont.gov/statutes/section/20/173/02731 | accessed 2026-06-26 | Commissioner of Public Safety authority; inspections; building-code update duties. |
| `src:usa-vt:stat-20-2730` | 20 V.S.A. § 2730, Definitions | statute | Vermont General Assembly / Vermont Laws Online | https://legislature.vermont.gov/statutes/section/20/173/02730 | accessed 2026-06-26 | Public-building scope and exclusions. |
| `src:usa-vt:stat-20-2736` | 20 V.S.A. § 2736, Enforcement; municipal enforcement assignment | statute | Vermont General Assembly / Vermont Laws Online | https://legislature.vermont.gov/statutes/section/20/173/02736 | accessed 2026-06-26 | Municipal enforcement assignment and retained Commissioner authority. |
| `src:usa-vt:stat-24-3101` | 24 V.S.A. § 3101, Establishment of building codes and regulations | statute | Vermont General Assembly / Vermont Laws Online | https://legislature.vermont.gov/statutes/section/24/083/03101 | accessed 2026-06-26 | Municipal building-code authority, state-rule consistency, electrical exclusion, local filing. |
| `src:usa-vt:sos-rule-25p022-fire-building` | Rule 25P022, 2025 Vermont Fire & Building Safety Code | rule record | Vermont Secretary of State, Rules Service | https://secure.vermont.gov/SOS/rules/results.php?keyword=&status=ADOPTED&number=25P022&title=&order=&page=1 | adopted status shown; accessed 2026-06-26 | Adopted fire/building safety rule record, authority, purpose and summary. |
| `src:usa-vt:sos-rule-25p023-electrical` | Rule 25P023, The 2025 Vermont Electrical Safety Rules | rule record | Vermont Secretary of State, Rules Service | https://secure.vermont.gov/SOS/rules/results.php?keyword=&status=ADOPTED&number=25P023&title=&order=&page=1 | adopted status shown; accessed 2026-06-26 | Adopted electrical rule and 2023 NEC adoption. |
| `src:usa-vt:stat-26-891-electrical` | 26 V.S.A. § 891, Electricians' Licensing Board powers; rules | statute | Vermont General Assembly / Vermont Laws Online | https://legislature.vermont.gov/statutes/section/26/015/00891 | accessed 2026-06-26 | Electrical rulemaking authority and authority to adopt a nationally recognized electrical code. |
| `src:usa-vt:sos-rule-25p024-plumbing` | Rule 25P024, The 2025 Vermont Plumbing Rules | rule record | Vermont Secretary of State, Rules Service | https://secure.vermont.gov/SOS/rules/results.php?keyword=&status=ADOPTED&number=25P024&title=&order=&page=1 | adopted status shown; accessed 2026-06-26 | Adopted plumbing rule and 2024 IPC update. |
| `src:usa-vt:stat-26-2173-plumbing` | 26 V.S.A. § 2173, Plumbing rules | statute | Vermont General Assembly / Vermont Laws Online | https://legislature.vermont.gov/statutes/section/26/039/02173 | accessed 2026-06-26 | Plumbing rulemaking authority, scope, enforcement, inspection authority. |
| `src:usa-vt:stat-30-51-rbes` | 30 V.S.A. § 51, Residential Building Energy Standards | statute | Vermont General Assembly / Vermont Laws Online | https://legislature.vermont.gov/statutes/section/30/002/00051 | accessed 2026-06-26 | RBES authority, update requirements, transition rule, certification. |
| `src:usa-vt:stat-30-53-cbes` | 30 V.S.A. § 53, Commercial Building Energy Standards | statute | Vermont General Assembly / Vermont Laws Online | https://legislature.vermont.gov/statutes/section/30/002/00053 | accessed 2026-06-26 | CBES authority, update requirements, transition rule, exemptions/variances. |
| `src:usa-vt:stat-20-174-accessibility` | 20 V.S.A. ch. 174, Access to Public Buildings | statute | Vermont General Assembly / Vermont Laws Online | https://legislature.vermont.gov/statutes/fullchapter/20/174 | accessed 2026-06-26 | Vermont Access Board, access standard requirements, enforcement. |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| `src:usa-vt:stat-20-2731` | online_statute_copy | Vermont Laws Online is a convenient state web source; production legal work should verify against the official Vermont Statutes Annotated or agency counsel. | usable_with_legal_verification |
| `src:usa-vt:stat-20-2730` | online_statute_copy | Same caveat as other Vermont Laws Online statutory pages. | usable_with_legal_verification |
| `src:usa-vt:stat-20-2736` | online_statute_copy | Same caveat as other Vermont Laws Online statutory pages. | usable_with_legal_verification |
| `src:usa-vt:stat-24-3101` | online_statute_copy | Same caveat as other Vermont Laws Online statutory pages. | usable_with_legal_verification |
| `src:usa-vt:sos-rule-25p022-fire-building` | rule_record_summary_only | Official rule record shows adopted status and summary, but this pass did not parse the full final code text or effective-date section. | needs_full_text_before_verified |
| `src:usa-vt:sos-rule-25p023-electrical` | rule_record_summary_only | Official rule record supports 2023 NEC adoption but not all effective-date and transition details. | needs_full_text_before_verified |
| `src:usa-vt:sos-rule-25p024-plumbing` | rule_record_summary_only | Official rule record supports 2024 IPC update but not all effective-date and transition details. | needs_full_text_before_verified |
| `src:usa-vt:stat-30-51-rbes` | online_statute_copy | Same caveat as other Vermont Laws Online statutory pages; final RBES rule text still needed for verified edition parsing. | usable_with_rule_text_verification |
| `src:usa-vt:stat-30-53-cbes` | online_statute_copy | Same caveat as other Vermont Laws Online statutory pages; final CBES rule text still needed for verified edition parsing. | usable_with_rule_text_verification |
| `src:usa-vt:stat-20-174-accessibility` | online_statute_copy | Same caveat as other Vermont Laws Online statutory pages; Access Rules edition/date still need rule-text extraction. | usable_with_rule_text_verification |

### 8.3 Supplemental Sources

| Source ID | Title | Source Type | Publisher / Maintainer | URL | Date / Version | Used For |
| --- | --- | --- | --- | --- | --- | --- |
| `src:usa-vt:efficiency-vt-energy-code-support` | Vermont Energy Code Support | guidance / support page | Efficiency Vermont | https://www.efficiencyvermont.com/trade-partners/energy-code-support | accessed 2026-06-26 | Latest RBES/CBES effective date, practical residential/commercial scope, official code-document links. |
| `src:usa-vt:cvr-31-000-004-rbes` | Code of Vermont Rules 31 000 004, 2024 RBES | regulatory mirror | Justia | https://regulations.justia.com/states/vermont/agency-31/sub-agency-000/chapter-004/section-31-000-004/ | current through June 2025; accessed 2026-06-26 | RBES model-code basis, scope, stretch-code note. |
| `src:usa-vt:cvr-31-000-003-cbes` | Code of Vermont Rules 31 000 003, 2024 CBES Amendments | regulatory mirror | Cornell Legal Information Institute | https://www.law.cornell.edu/regulations/vermont/31-003-Code-Vt-R-31-000-003-X | accessed 2026-06-26 | CBES model-code basis, scope, AHJ caveat, preemption note. |
| `src:usa-vt:burlington-adopted-codes` | Adopted Codes | municipal code-information page | City of Burlington, Vermont | https://www.burlingtonvt.gov/286/Adopted-Codes | accessed 2026-06-26 | Supplemental list of 2025 Vermont adopted model-code editions; not used as a primary statewide source. |
| `src:usa-vt:bennington-2025-code-release` | New 2025 Fire and Building Codes Released | municipal notice | Town of Bennington, Vermont | https://benningtonvt.org/new-2025-fire-and-building-codes-released/ | accessed 2026-06-26 | Supplemental notice identifying 2025 Vermont Fire & Building Safety Code adoption date. |
| `src:usa-vt:efficiency-vt-2026-fire-code-training` | Vermont Fire and Building Safety Code updates referenced construction codes | training / event page | Efficiency Vermont | https://www.efficiencyvermont.com/events/vermont-fire-and-building-safety-code-updates-for-architects-engineers-and-builders | accessed 2026-06-26 | Supplemental confirmation that referenced construction codes updated from 2015 to 2021 editions. |

### 8.4 Source Extraction Metadata

| Extraction ID | Source IDs | Extracted By | Extraction Date | Method | Completeness | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| `extract:usa-vt:authority-statutes` | `src:usa-vt:stat-20-2731`, `src:usa-vt:stat-20-2730`, `src:usa-vt:stat-20-2736`, `src:usa-vt:stat-24-3101` | ChatGPT | 2026-06-26 | HTML review | high_for_core_authority | Core authority and local-enforcement framework captured. |
| `extract:usa-vt:trade-rules` | `src:usa-vt:sos-rule-25p023-electrical`, `src:usa-vt:stat-26-891-electrical`, `src:usa-vt:sos-rule-25p024-plumbing`, `src:usa-vt:stat-26-2173-plumbing` | ChatGPT | 2026-06-26 | HTML review | medium | Model-code edition captured for electrical and plumbing; full final rule text not parsed. |
| `extract:usa-vt:energy` | `src:usa-vt:stat-30-51-rbes`, `src:usa-vt:stat-30-53-cbes`, `src:usa-vt:efficiency-vt-energy-code-support`, `src:usa-vt:cvr-31-000-004-rbes`, `src:usa-vt:cvr-31-000-003-cbes` | ChatGPT | 2026-06-26 | HTML review | medium | Statutory transition logic captured; final official rule text needed for verified edition parsing. |
| `extract:usa-vt:fire-building-2025` | `src:usa-vt:sos-rule-25p022-fire-building`, `src:usa-vt:burlington-adopted-codes`, `src:usa-vt:bennington-2025-code-release`, `src:usa-vt:efficiency-vt-2026-fire-code-training` | ChatGPT | 2026-06-26 | HTML review | medium_low | Rule adopted status captured; model-code edition mapping partially reliant on supplemental sources. |
| `extract:usa-vt:accessibility` | `src:usa-vt:stat-20-174-accessibility`, `src:usa-vt:burlington-adopted-codes` | ChatGPT | 2026-06-26 | HTML review | low_medium | Statutory board/enforcement captured; current access-rule edition and text need direct extraction. |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| `report` | `report.status` | partially_verified | verified | 1.00 | none | Status reflects source-backed core authority and adoption fields with unresolved date/model-code details. |
| `authority:primary` | `authority.name` | Commissioner of Public Safety / Division of Fire Safety | verified | 0.85 | `src:usa-vt:stat-20-2731`, `src:usa-vt:sos-rule-25p022-fire-building` | Primary public-building fire/building authority. |
| `authority:primary` | `scope` | public buildings | verified | 0.80 | `src:usa-vt:stat-20-2730` | Owner-occupied single-family residence exclusion captured. |
| `adoption:fire-building-2025` | `rule_status` | adopted | verified | 0.85 | `src:usa-vt:sos-rule-25p022-fire-building` | SOS rule record confirms adopted status. |
| `adoption:fire-building-2025` | `model_code_editions` | 2021 key code references | partially_verified | 0.55 | `src:usa-vt:burlington-adopted-codes`, `src:usa-vt:efficiency-vt-2026-fire-code-training` | Needs full state-code text confirmation. |
| `adoption:electrical-2025` | `base_model_code` | 2023 NEC | verified | 0.88 | `src:usa-vt:sos-rule-25p023-electrical`, `src:usa-vt:stat-26-891-electrical` | Supported by official rule record and statutory authority. |
| `adoption:plumbing-2025` | `base_model_code` | 2024 IPC | verified | 0.86 | `src:usa-vt:sos-rule-25p024-plumbing`, `src:usa-vt:stat-26-2173-plumbing` | Supported by official rule record and statutory authority. |
| `adoption:rbes-2024` | `effective_date` | 2024-07-01 | partially_verified | 0.78 | `src:usa-vt:efficiency-vt-energy-code-support`, `src:usa-vt:stat-30-51-rbes` | Guidance gives date; statute gives transition framework. |
| `adoption:cbes-2024` | `effective_date` | 2024-07-01 | partially_verified | 0.78 | `src:usa-vt:efficiency-vt-energy-code-support`, `src:usa-vt:stat-30-53-cbes` | Guidance gives date; statute gives transition framework. |
| `local-enforcement` | `model` | state_administered_with_commissioner_assigned_municipal_enforcement | verified | 0.80 | `src:usa-vt:stat-20-2736` | Commissioner assignment and retained authority captured. |
| `local-amendment-rule` | `model` | municipal building codes allowed if consistent with state rules and electrical excluded | partially_verified | 0.70 | `src:usa-vt:stat-24-3101` | Approval and registry mechanics unresolved. |
| `accessibility` | `authority` | Vermont Access Board / Department of Public Safety | partially_verified | 0.72 | `src:usa-vt:stat-20-174-accessibility` | Current access-rule edition remains unresolved. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Every `src:usa-vt:*` used in the body is present in section 8, except `none` placeholders for explicitly unresolved items. |
| All authority IDs resolve | pass | Authority IDs used in section 2 are internally declared or explicitly unresolved. |
| All current code families have adoption matrix rows | pass | Rows are retained for all template code families; unresolved rows are explicit. |
| Normalized adoption records exist only for supported facts | pass | Adoption records are limited to fire/building, electrical, plumbing, energy, and accessibility facts supported by captured sources. |
| Building and operational fire code are separated | pass | Separate rows are present for building, fire-construction references, and operational/prevention fire code. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Unknown dates remain null; energy dates are repeated only where the evidence supports effective/operative/mandatory use. |
| Effective dates are valid ISO dates | pass | Populated dates use YYYY-MM-DD format. |
| No impossible date sequences | pass | No adoption/effective sequence conflict was introduced; unresolved dates remain null. |
| Transition rules have explicit trigger conditions | partial | RBES and CBES transition triggers are captured; fire/building, electrical, and plumbing triggers remain unresolved. |
| Permit-date logic is captured where applicable | partial | CBES permit/plan-approval logic is captured; RBES construction-commencement logic is captured; fire/building permit-date logic remains unresolved. |
| Local enforcement model classified | pass | Municipal enforcement assignment model captured. |
| Local amendment rule classified | partial | Municipal code consistency and electrical exclusion captured; approval and registry details unresolved. |
| AHJ confirmation metadata present | fail | No AHJ contact dataset or municipal assignment dataset was extracted. |
| Official-source caveats captured | pass | Caveats distinguish official rule summaries from full rule text and supplemental sources. |
| Leftover template markers removed | pass | Template placeholders were replaced or converted into explicit unresolved fields. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `issue:usa-vt:001` | high | fire/building final text | Full 2025 Vermont Fire & Building Safety Code text was not parsed, so model-code edition mapping and effective/mandatory dates are only partially verified. | Obtain final adopted code text and parse adoption, effective, operative, mandatory, and transition language. | null | null | open |
| `issue:usa-vt:002` | high | fire/building date mechanics | Adoption date is based on supplemental municipal notice while official SOS summary confirms adopted status; final date fields need primary-source confirmation. | Review adopted rule filing, LCAR materials, agency code publication, and any implementation memo. | null | null | open |
| `issue:usa-vt:003` | medium | local amendment scope | Municipal building-code authority is verified, but state approval requirements, consistency-review process, and statewide registry status remain unresolved. | Review municipal enabling statutes, Division of Fire Safety guidance, and municipal code filing practices. | null | null | open |
| `issue:usa-vt:004` | medium | AHJ resolution | No dataset identifies which municipalities have Commissioner-assigned enforcement responsibility. | Locate or request Division of Fire Safety municipal assignment records and regional inspection coverage. | null | null | open |
| `issue:usa-vt:005` | medium | access-rule edition | Vermont Access Board authority is verified, but current Access Rules edition/date and amendments were not fully parsed. | Obtain Vermont Access Rules final text and rule history. | null | null | open |
| `issue:usa-vt:006` | medium | elevator/conveyance authority | Elevator/conveyance authority and current code adoption were not researched. | Review Vermont conveyance/elevator statutes and Division of Fire Safety program materials. | null | null | open |
| `issue:usa-vt:007` | low | energy code final text | RBES/CBES effective date is supported by guidance and statutes, but official final rule text should be used before verified production status. | Download official RBES/CBES PDFs or Secretary of State rule filings and reconcile scope/trigger language. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| `watch:usa-vt:sos-fire-building-rule` | `src:usa-vt:sos-rule-25p022-fire-building` | html_diff | monthly | rule record changes, links to final text, effective-date clarification, or replacement rule | 2026-06-26 | active |
| `watch:usa-vt:sos-electrical-rule` | `src:usa-vt:sos-rule-25p023-electrical` | html_diff | monthly | electrical rule record changes or new NEC-cycle filing | 2026-06-26 | active |
| `watch:usa-vt:sos-plumbing-rule` | `src:usa-vt:sos-rule-25p024-plumbing` | html_diff | monthly | plumbing rule record changes or new IPC-cycle filing | 2026-06-26 | active |
| `watch:usa-vt:energy-support` | `src:usa-vt:efficiency-vt-energy-code-support` | html_diff | monthly | new RBES/CBES code documents, effective dates, compliance forms, or 2026 amendments | 2026-06-26 | active |
| `watch:usa-vt:stat-20-2731` | `src:usa-vt:stat-20-2731` | statute_diff | quarterly | amendments to public-building authority or Commissioner code-update duties | 2026-06-26 | active |
| `watch:usa-vt:stat-24-3101` | `src:usa-vt:stat-24-3101` | statute_diff | quarterly | amendments to municipal building-code authority or filing requirements | 2026-06-26 | active |
| `watch:usa-vt:access-rules` | `src:usa-vt:stat-20-174-accessibility` | statute_and_rule_diff | quarterly | Access Board or access-rule amendments | 2026-06-26 | pending_rule_text |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-24 | Generated baseline draft report | `report:usa-vt` | none | Codex | Baseline draft contained unresolved placeholders and no primary sources. |
| 2026-06-26 | Populated Vermont report with source-backed core authority, adoption matrix, source registry, QA, and open issues | `state-report:usa-vt`, `ahj:usa-vt:dps-commissioner-fire-safety`, `adoption:usa-vt:fire-building-2025`, `adoption:usa-vt:electrical-2025`, `adoption:usa-vt:plumbing-2025`, `adoption:usa-vt:rbes-2024`, `adoption:usa-vt:cbes-2024` | `src:usa-vt:stat-20-2731`, `src:usa-vt:stat-20-2730`, `src:usa-vt:stat-20-2736`, `src:usa-vt:stat-24-3101`, `src:usa-vt:sos-rule-25p022-fire-building`, `src:usa-vt:sos-rule-25p023-electrical`, `src:usa-vt:sos-rule-25p024-plumbing`, `src:usa-vt:stat-30-51-rbes`, `src:usa-vt:stat-30-53-cbes` | ChatGPT | Upgraded status to partially_verified; kept unresolved date, local-amendment, AHJ, access-rule, and conveyance issues explicit. |
