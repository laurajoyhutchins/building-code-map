---
state:
  state_id: "US-OK"
  name: "Oklahoma"
  abbreviation: "OK"
report:
  report_id: "state-report:usa-ok"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-26"
  last_verified: null
  reviewed_by: null
risk:
  overall_confidence: 0.72 # 0.00 - 1.00
  risk_flags:
    - "energy_code_scope_unresolved"
    - "accessibility_authority_not_independently_parsed"
    - "local_amendment_process_after_chapter_15_revocation_needs_review"
    - "2024_code_rulemaking_status_needs_monitoring"
    - "date_rules_limited_to_effective_dates"
  open_questions_count: 6

---

# State Building Code Authority Report: Oklahoma

## 1. Executive Summary

- **Authority model:** Oklahoma has a statewide minimum-code adoption model. The Oklahoma Uniform Building Code Commission (OUBCC), located within the Construction Industries Board, is the primary statewide authority for adopting residential and commercial construction codes. OUBCC-adopted codes and standards are minimum standards for residential and commercial construction in the state. Local jurisdictions and other political subdivisions may enact and enforce higher standards, subject to Oklahoma law and Commission approval where local-condition amendments are involved.

- **Statewide code status:** OUBCC currently lists adopted statewide minimum editions for IBC 2018, IFC 2018, IRC 2018, IEBC 2018, IFGC 2018, IMC 2018, IPC 2018, and NEC 2023. This pass verifies the code-family rows and effective dates shown on OUBCC adoption pages. Standalone statewide energy-code status was not resolved from OUBCC adoption pages; residential energy provisions are embedded in the IRC row, but commercial IECC status remains open.

- **Local enforcement model:** OUBCC adopts minimum standards. OUBCC's FAQ states that local jurisdictions, municipalities, and other political subdivisions interpret and enforce adopted minimum codes; in unincorporated areas, the Construction Industries Board handles electrical, mechanical, plumbing, and fuel-gas minimum-code enforcement, while the Oklahoma State Fire Marshal handles commercial building-code enforcement.

- **Local amendment posture:** Local governments may enforce requirements that are higher than the OUBCC minimum standards. Oklahoma statute also states that local changes necessary to accommodate local conditions must be approved by the Commission. The prior Title 748, Chapter 15 local-code-adoption rule chapter is revoked, so the current procedure for submitting or approving local changes needs follow-up.

- **Known transition periods or pending changes:** OUBCC's proposed-rules page states that the Commission began rulemaking to adopt 2024 editions of IBC, IFC, IRC, IEBC, IFGC, IMC, and IPC in late 2025, with public-comment and hearing milestones in January 2026. This pass found proposed-rule activity but did not verify final adoption or an effective date for the 2024 code package.

- **Production readiness:** partially_ready_for_internal_use

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: OUBCC is the primary statewide minimum-code adopting authority for residential
    and commercial construction.
  confidence: 0.9
  source_ids:
  - src:usa-ok:oubcc-statute-act
  - src:usa-ok:oubcc-home
- topic: Primary commercial building code edition
  finding: OUBCC lists the 2018 IBC permanent rule as the current unsuperseded IBC
    adoption, effective 2021-09-14.
  confidence: 0.88
  source_ids:
  - src:usa-ok:oubcc-ibc-adoptions
- topic: Primary residential code edition
  finding: OUBCC lists the 2018 IRC permanent rule as the current unsuperseded IRC
    adoption, effective 2022-09-14.
  confidence: 0.88
  source_ids:
  - src:usa-ok:oubcc-irc-adoptions
- topic: Electrical code edition
  finding: OUBCC lists the 2023 NEC permanent rule as current, effective 2024-09-14.
  confidence: 0.88
  source_ids:
  - src:usa-ok:oubcc-nec-adoptions
- topic: Fire code authority and edition
  finding: OUBCC lists the 2018 IFC permanent rule as current, effective 2021-09-14;
    local/OSFM enforcement allocation should be verified beyond the FAQ before production.
  confidence: 0.78
  source_ids:
  - src:usa-ok:oubcc-ifc-adoptions
  - src:usa-ok:oubcc-faq
- topic: Local enforcement
  finding: "Local jurisdictions enforce adopted minimum codes; CIB and the State Fire\
    \ Marshal have specified roles in unincorporated areas according to OUBCC'\
    s FAQ."
  confidence: 0.76
  source_ids:
  - src:usa-ok:oubcc-faq
- topic: Local amendments
  finding: Higher local standards are allowed, but local-condition amendments require
    Commission approval; the post-revocation administrative procedure needs follow-up.
  confidence: 0.68
  source_ids:
  - src:usa-ok:oubcc-statute-act
  - src:usa-ok:oubcc-admin-2021-rules
- topic: Pending changes
  finding: OUBCC has an active 2024-code rulemaking record, but final adoption was
    not confirmed in this pass.
  confidence: 0.75
  source_ids:
  - src:usa-ok:oubcc-proposed-rules
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-ok:oubcc |
| Authority name | Oklahoma Uniform Building Code Commission |
| Authority type | statewide code commission within Construction Industries Board |
| Legal basis | Oklahoma Uniform Building Code Commission Act, 59 O.S. §§ 1000.20-1000.30; OUBCC's statutory duty is to review and adopt building codes for residential and commercial construction to be used by all entities in the state. |
| Role | Adopts statewide minimum residential and commercial construction codes; uses rulemaking and technical committees; publishes adopted-code and amendment materials. |
| Enforcement model | statewide minimum code adoption with local and special-state enforcement roles |
| Source IDs | src:usa-ok:oubcc-statute-act; src:usa-ok:oubcc-home; src:usa-ok:oubcc-codes-rules |
| Verification status | verified_for_core_authority |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-ok:oubcc | Oklahoma Uniform Building Code Commission | Adopts statewide minimum commercial building code. | 59 O.S. § 1000.23; OAC Title 748, Chapter 20 adoption records | src:usa-ok:oubcc-statute-act; src:usa-ok:oubcc-ibc-adoptions | verified |
| Residential | ahj:usa-ok:oubcc | Oklahoma Uniform Building Code Commission | Adopts statewide minimum residential building code for one- and two-family dwellings and townhouses. | 59 O.S. § 1000.23; OAC Title 748, Chapter 20 adoption records | src:usa-ok:oubcc-statute-act; src:usa-ok:oubcc-irc-adoptions | verified |
| Existing Building / Rehabilitation | ahj:usa-ok:oubcc | Oklahoma Uniform Building Code Commission | Adopts statewide minimum commercial existing-building code. | 59 O.S. § 1000.23; OAC Title 748, Chapter 20 adoption records | src:usa-ok:oubcc-statute-act; src:usa-ok:oubcc-iebc-adoptions | verified |
| Mechanical | ahj:usa-ok:oubcc | Oklahoma Uniform Building Code Commission | Adopts statewide minimum commercial mechanical code. | 59 O.S. § 1000.23; OAC Title 748, Chapter 20 adoption records | src:usa-ok:oubcc-statute-act; src:usa-ok:oubcc-imc-adoptions | verified |
| Plumbing | ahj:usa-ok:oubcc | Oklahoma Uniform Building Code Commission | Adopts statewide minimum commercial plumbing code. | 59 O.S. § 1000.23; OAC Title 748, Chapter 20 adoption records | src:usa-ok:oubcc-statute-act; src:usa-ok:oubcc-ipc-adoptions | verified |
| Fuel Gas | ahj:usa-ok:oubcc | Oklahoma Uniform Building Code Commission | Adopts statewide minimum commercial fuel-gas code. | 59 O.S. § 1000.23; OAC Title 748, Chapter 20 adoption records | src:usa-ok:oubcc-statute-act; src:usa-ok:oubcc-ifgc-adoptions | verified |
| Electrical | ahj:usa-ok:oubcc | Oklahoma Uniform Building Code Commission | Adopts statewide NEC edition and amendments; CIB has enforcement role in unincorporated areas per FAQ. | 59 O.S. § 1000.23; OAC Title 748, Chapter 20 adoption records | src:usa-ok:oubcc-statute-act; src:usa-ok:oubcc-nec-adoptions; src:usa-ok:oubcc-faq | verified_for_adoption; enforcement_partial |
| Energy | ahj:usa-ok:oubcc | Oklahoma Uniform Building Code Commission | OUBCC mission covers codes affecting the built environment; residential energy provisions are in IRC, but standalone commercial IECC adoption was not resolved. | 59 O.S. § 1000.23; IRC adoption records | src:usa-ok:oubcc-statute-act; src:usa-ok:oubcc-irc-adoptions; src:usa-ok:oubcc-home | unresolved_scope |
| Fire - construction references | ahj:usa-ok:oubcc | Oklahoma Uniform Building Code Commission | Adopts statewide IFC edition as part of construction-code set. | 59 O.S. § 1000.23; OAC Title 748, Chapter 20 adoption records | src:usa-ok:oubcc-statute-act; src:usa-ok:oubcc-ifc-adoptions | verified |
| Fire - operational / prevention code | ahj:usa-ok:oubcc | Oklahoma Uniform Building Code Commission | OUBCC lists IFC 2018 adoption; operational enforcement allocation remains partial pending OSFM rule review. | 59 O.S. § 1000.23; OUBCC FAQ for unincorporated enforcement | src:usa-ok:oubcc-ifc-adoptions; src:usa-ok:oubcc-faq | partially_verified |
| Accessibility | ahj:usa-ok:oubcc | Oklahoma Uniform Building Code Commission | Accessibility provisions are included through adopted building/residential/existing-building codes; exact standalone standard path was not parsed. | OAC Title 748, Chapter 20 adoption records | src:usa-ok:oubcc-ibc-adoptions; src:usa-ok:oubcc-irc-adoptions; src:usa-ok:oubcc-iebc-adoptions | unresolved_standard_edition |
| Elevator / Conveyance | ahj:usa-ok:odol | Oklahoma Department of Labor | Administers Elevator Safety Act rules and adopts national elevator/conveyance standards. | 59 O.S. elevator statutes and OAC Title 380, Chapter 70, as published by ODOL | src:usa-ok:odol-elevator-rules | partially_verified |

### 2.3 Authority Hierarchy Notes

Oklahoma's statewide model distinguishes adoption from enforcement. OUBCC establishes minimum statewide codes. Public projects, municipalities, and political subdivisions must meet the minimum standards, while state agencies and local governments may impose higher standards. The OUBCC FAQ assigns day-to-day interpretation and enforcement to local jurisdictions, municipalities, and political subdivisions, with CIB and the State Fire Marshal filling specified unincorporated-area roles.

The local-amendment process is partially unresolved. The statute preserves higher standards and allows local-condition amendments with Commission approval, but the older Title 748, Chapter 15 procedure chapter has been revoked. This means the legal authority is verified, while the current procedural path needs a targeted rule or agency-practice review.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-ok:001 | ahj:usa-ok:oubcc | adopts_minimum_codes_for | residential_and_commercial_construction_statewide | src:usa-ok:oubcc-statute-act; src:usa-ok:oubcc-home | verified |
| edge:usa-ok:002 | ahj:usa-ok:oubcc | publishes_adopted_codes_and_amendments_at | OUBCC Codes & Rules portal | src:usa-ok:oubcc-codes-rules | verified |
| edge:usa-ok:003 | ahj:usa-ok:oubcc | minimum_standards_bind | state_agencies_municipalities_political_subdivisions | src:usa-ok:oubcc-statute-act; src:usa-ok:oubcc-faq | verified |
| edge:usa-ok:004 | local_jurisdictions | interpret_and_enforce | adopted_minimum_codes | src:usa-ok:oubcc-faq | partially_verified |
| edge:usa-ok:005 | ahj:usa-ok:cib | enforces_in_unincorporated_areas | electrical_mechanical_plumbing_fuel_gas_minimum_codes | src:usa-ok:oubcc-faq | partially_verified |
| edge:usa-ok:006 | ahj:usa-ok:state-fire-marshal | enforces_in_unincorporated_areas | commercial_building_code | src:usa-ok:oubcc-faq | partially_verified |
| edge:usa-ok:007 | municipalities_political_subdivisions | may_enact_and_enforce | higher_standards_than_state_minimum | src:usa-ok:oubcc-statute-act; src:usa-ok:oubcc-faq | verified |
| edge:usa-ok:008 | state_agencies_political_subdivisions | local_condition_amendments_require_approval_by | ahj:usa-ok:oubcc | src:usa-ok:oubcc-statute-act | partially_verified |
| edge:usa-ok:009 | ahj:usa-ok:odol | adopts_and_enforces | elevator_and_conveyance_safety_standards | src:usa-ok:odol-elevator-rules | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | Oklahoma adopted IBC | International Building Code | 2018 | current_statewide_minimum | 2021-09-14 | 2021-09-14 | null | null | No separate operative or mandatory date verified; OUBCC page identifies effective date. | src:usa-ok:oubcc-ibc-adoptions |
| Residential | Oklahoma adopted IRC | International Residential Code | 2018 | current_statewide_minimum | 2022-09-14 | 2022-09-14 | null | null | No separate operative or mandatory date verified; OUBCC page identifies effective date. | src:usa-ok:oubcc-irc-adoptions |
| Existing Building / Rehabilitation | Oklahoma adopted IEBC | International Existing Building Code | 2018 | current_statewide_minimum | 2021-09-14 | 2021-09-14 | null | null | No separate operative or mandatory date verified; OUBCC page identifies effective date. | src:usa-ok:oubcc-iebc-adoptions |
| Mechanical | Oklahoma adopted IMC | International Mechanical Code | 2018 | current_statewide_minimum | 2021-09-14 | 2021-09-14 | null | null | No separate operative or mandatory date verified; OUBCC page identifies effective date. | src:usa-ok:oubcc-imc-adoptions |
| Plumbing | Oklahoma adopted IPC | International Plumbing Code | 2018 | current_statewide_minimum | 2021-09-14 | 2021-09-14 | null | null | No separate operative or mandatory date verified; OUBCC page identifies effective date. | src:usa-ok:oubcc-ipc-adoptions |
| Fuel Gas | Oklahoma adopted IFGC | International Fuel Gas Code | 2018 | current_statewide_minimum | 2021-09-14 | 2021-09-14 | null | null | No separate operative or mandatory date verified; OUBCC page identifies effective date. | src:usa-ok:oubcc-ifgc-adoptions |
| Electrical | Oklahoma adopted NEC | National Electrical Code / NFPA 70 | 2023 | current_statewide_minimum | 2024-09-14 | 2024-09-14 | null | null | No separate operative or mandatory date verified; OUBCC page identifies effective date. | src:usa-ok:oubcc-nec-adoptions |
| Energy | Residential energy provisions in Oklahoma adopted IRC; commercial standalone IECC unresolved | IRC Chapter 11 / IECC scope unresolved | 2018 IRC; IECC edition unresolved | partially_verified_scope | 2022-09-14 | 2022-09-14 | null | null | Residential row follows IRC effective date; commercial energy-code row requires targeted review. | src:usa-ok:oubcc-irc-adoptions; src:usa-ok:oubcc-home |
| Fire - construction references | Oklahoma adopted IFC | International Fire Code | 2018 | current_statewide_minimum | 2021-09-14 | 2021-09-14 | null | null | No separate operative or mandatory date verified; OUBCC page identifies effective date. | src:usa-ok:oubcc-ifc-adoptions |
| Fire - operational / prevention code | Oklahoma adopted IFC | International Fire Code | 2018 | current_statewide_minimum_with_enforcement_scope_open | 2021-09-14 | 2021-09-14 | null | null | OUBCC adopts IFC; operational AHJ allocation needs OSFM/local review. | src:usa-ok:oubcc-ifc-adoptions; src:usa-ok:oubcc-faq |
| Accessibility | Accessibility provisions in adopted IBC/IRC/IEBC | IBC/IRC/IEBC accessibility chapters and referenced standards | edition not independently parsed | unresolved_standard_edition | null | null | null | null | Exact accessibility-standard edition and state/federal interface were not parsed. | src:usa-ok:oubcc-ibc-adoptions; src:usa-ok:oubcc-irc-adoptions; src:usa-ok:oubcc-iebc-adoptions |
| Elevator / Conveyance | Oklahoma Elevator Safety Act rules | ASME A17.1, ASME A17.3, ASME A18.1, ASME A90.1, ASSE A10.4, NFPA 70, IBC, ANSI/SIA A92.10-2009 | latest editions/current addenda by rule; ANSI/SIA A92.10-2009 fixed | current_administrative_rule | null | null | null | null | ODOL rule uses moving references; exact edition effective chronology needs follow-up. | src:usa-ok:odol-elevator-rules |

### 3.2 Adoption Records

```yaml
adoption_records:
  - adoption_id: adoption:usa-ok:ibc-2018
    code_family: Building
    adopting_authority_id: ahj:usa-ok:oubcc
    state_code_name: Oklahoma adopted IBC
    base_model_code: International Building Code
    edition: "2018"
    status: current_statewide_minimum
    adoption_date: "2021-09-14"
    effective_date: "2021-09-14"
    operative_date: null
    mandatory_date: null
    source_ids:
      - src:usa-ok:oubcc-ibc-adoptions
    notes: "OUBCC adoption page lists the 2021-09-14 IBC 2018 permanent rule as the current unsuperseded adoption."

  - adoption_id: adoption:usa-ok:irc-2018
    code_family: Residential
    adopting_authority_id: ahj:usa-ok:oubcc
    state_code_name: Oklahoma adopted IRC
    base_model_code: International Residential Code
    edition: "2018"
    status: current_statewide_minimum
    adoption_date: "2022-09-14"
    effective_date: "2022-09-14"
    operative_date: null
    mandatory_date: null
    source_ids:
      - src:usa-ok:oubcc-irc-adoptions
    notes: "OUBCC adoption page lists the 2022-09-14 IRC 2018 permanent rule as the current unsuperseded adoption."

  - adoption_id: adoption:usa-ok:iebc-2018
    code_family: Existing Building / Rehabilitation
    adopting_authority_id: ahj:usa-ok:oubcc
    state_code_name: Oklahoma adopted IEBC
    base_model_code: International Existing Building Code
    edition: "2018"
    status: current_statewide_minimum
    adoption_date: "2021-09-14"
    effective_date: "2021-09-14"
    operative_date: null
    mandatory_date: null
    source_ids:
      - src:usa-ok:oubcc-iebc-adoptions
    notes: "OUBCC adoption page lists the 2021-09-14 IEBC 2018 permanent rule as the current unsuperseded adoption."

  - adoption_id: adoption:usa-ok:ifc-2018
    code_family: Fire - construction references; Fire - operational / prevention code
    adopting_authority_id: ahj:usa-ok:oubcc
    state_code_name: Oklahoma adopted IFC
    base_model_code: International Fire Code
    edition: "2018"
    status: current_statewide_minimum
    adoption_date: "2021-09-14"
    effective_date: "2021-09-14"
    operative_date: null
    mandatory_date: null
    source_ids:
      - src:usa-ok:oubcc-ifc-adoptions
    notes: "OUBCC adoption page lists the 2021-09-14 IFC 2018 permanent rule as the current unsuperseded adoption."

  - adoption_id: adoption:usa-ok:ifgc-2018
    code_family: Fuel Gas
    adopting_authority_id: ahj:usa-ok:oubcc
    state_code_name: Oklahoma adopted IFGC
    base_model_code: International Fuel Gas Code
    edition: "2018"
    status: current_statewide_minimum
    adoption_date: "2021-09-14"
    effective_date: "2021-09-14"
    operative_date: null
    mandatory_date: null
    source_ids:
      - src:usa-ok:oubcc-ifgc-adoptions
    notes: "OUBCC adoption page lists the 2021-09-14 IFGC 2018 permanent rule as the current unsuperseded adoption."

  - adoption_id: adoption:usa-ok:imc-2018
    code_family: Mechanical
    adopting_authority_id: ahj:usa-ok:oubcc
    state_code_name: Oklahoma adopted IMC
    base_model_code: International Mechanical Code
    edition: "2018"
    status: current_statewide_minimum
    adoption_date: "2021-09-14"
    effective_date: "2021-09-14"
    operative_date: null
    mandatory_date: null
    source_ids:
      - src:usa-ok:oubcc-imc-adoptions
    notes: "OUBCC adoption page lists the 2021-09-14 IMC 2018 permanent rule as the current unsuperseded adoption."

  - adoption_id: adoption:usa-ok:ipc-2018
    code_family: Plumbing
    adopting_authority_id: ahj:usa-ok:oubcc
    state_code_name: Oklahoma adopted IPC
    base_model_code: International Plumbing Code
    edition: "2018"
    status: current_statewide_minimum
    adoption_date: "2021-09-14"
    effective_date: "2021-09-14"
    operative_date: null
    mandatory_date: null
    source_ids:
      - src:usa-ok:oubcc-ipc-adoptions
    notes: "OUBCC adoption page lists the 2021-09-14 IPC 2018 permanent rule as the current unsuperseded adoption."

  - adoption_id: adoption:usa-ok:nec-2023
    code_family: Electrical
    adopting_authority_id: ahj:usa-ok:oubcc
    state_code_name: Oklahoma adopted NEC
    base_model_code: National Electrical Code / NFPA 70
    edition: "2023"
    status: current_statewide_minimum
    adoption_date: "2024-09-14"
    effective_date: "2024-09-14"
    operative_date: null
    mandatory_date: null
    source_ids:
      - src:usa-ok:oubcc-nec-adoptions
    notes: "OUBCC adoption page lists the 2023 NEC permanent rule as effective 2024-09-14 and unsuperseded."
```

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

This pass verified effective dates from OUBCC adoption pages for the main statewide code-family rows. The sources reviewed did not identify separate operative dates, mandatory dates, grace periods, permit-application triggers, permit-issuance triggers, or concurrency rules for the current statewide code adoptions. The statewide legal effect is framed as minimum standards for residential and commercial construction, with local jurisdictions required to abide by those minimum standards.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-ok:001 | OUBCC IBC/IFC/IEBC/IFGC/IMC/IPC 2018 adoptions | effective_date | 2021-09-14 | OUBCC permanent-rule effective date listed on code-specific adoption page | not verified | src:usa-ok:oubcc-ibc-adoptions; src:usa-ok:oubcc-ifc-adoptions; src:usa-ok:oubcc-iebc-adoptions; src:usa-ok:oubcc-ifgc-adoptions; src:usa-ok:oubcc-imc-adoptions; src:usa-ok:oubcc-ipc-adoptions | verified_for_effective_date_only |
| date-rule:usa-ok:002 | OUBCC IRC 2018 adoption | effective_date | 2022-09-14 | OUBCC permanent-rule effective date listed on IRC adoption page | not verified | src:usa-ok:oubcc-irc-adoptions | verified_for_effective_date_only |
| date-rule:usa-ok:003 | OUBCC NEC 2023 adoption | effective_date | 2024-09-14 | OUBCC permanent-rule effective date listed on NEC adoption page | not verified | src:usa-ok:oubcc-nec-adoptions | verified_for_effective_date_only |
| date-rule:usa-ok:004 | local adoptions predating OUBCC uniform codes | savings_rule | no fixed date | Codes adopted by state agencies, municipalities, or other political subdivisions before OUBCC uniform-code adoption remain valid until OUBCC adopts uniform codes. | yes, within savings rule | src:usa-ok:oubcc-statute-act | partially_verified |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | IBC 2024 | 2025-11-04 | null | null | null | null | active_watch | src:usa-ok:oubcc-proposed-rules | Proposed-rule page says OUBCC began rulemaking to adopt 2024 IBC. Final adoption not verified. |
| Residential | IRC 2024 | 2025-11-04 | null | null | null | null | active_watch | src:usa-ok:oubcc-proposed-rules | Proposed-rule page says OUBCC began rulemaking to adopt 2024 IRC. Final adoption not verified. |
| Existing Building / Rehabilitation | IEBC 2024 | 2025-11-04 | null | null | null | null | active_watch | src:usa-ok:oubcc-proposed-rules | Proposed-rule page says OUBCC began rulemaking to adopt 2024 IEBC. Final adoption not verified. |
| Fire | IFC 2024 | 2025-11-04 | null | null | null | null | active_watch | src:usa-ok:oubcc-proposed-rules | Proposed-rule page says OUBCC began rulemaking to adopt 2024 IFC. Final adoption not verified. |
| Fuel Gas | IFGC 2024 | 2025-11-04 | null | null | null | null | active_watch | src:usa-ok:oubcc-proposed-rules | Proposed-rule page says OUBCC began rulemaking to adopt 2024 IFGC. Final adoption not verified. |
| Mechanical | IMC 2024 | 2025-11-04 | null | null | null | null | active_watch | src:usa-ok:oubcc-proposed-rules | Proposed-rule page says OUBCC began rulemaking to adopt 2024 IMC. Final adoption not verified. |
| Plumbing | IPC 2024 | 2025-11-04 | null | null | null | null | active_watch | src:usa-ok:oubcc-proposed-rules | Proposed-rule page says OUBCC began rulemaking to adopt 2024 IPC. Final adoption not verified. |
| Electrical | NEC 2027 | 2026-05-26 | null | null | null | null | early_watch | src:usa-ok:oubcc-home | OUBCC home page requests code-change proposals for the 2027 NEC by 2026-06-30. |
| Commercial Energy | IECC 2021 / IECC 2024 | 2026-05-26 | null | null | null | null | early_watch | src:usa-ok:oubcc-home | OUBCC home page calls for volunteers for a Commercial Energy Conservation Technical Code Review Committee reviewing 2021 and 2024 IECC commercial provisions. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-ok:001 | all OUBCC-adopted residential and commercial construction codes | public projects | public-project status | Public projects must abide by OUBCC minimum standards, while state agencies may enforce higher standards. | src:usa-ok:oubcc-statute-act; src:usa-ok:oubcc-faq | verified |
| applicability-rule:usa-ok:002 | all OUBCC-adopted residential and commercial construction codes | municipalities and political subdivisions | local enforcement or local code adoption | Municipalities and political subdivisions must abide by OUBCC minimum standards and may enforce higher standards. | src:usa-ok:oubcc-statute-act; src:usa-ok:oubcc-faq | verified |
| applicability-rule:usa-ok:003 | elevator / conveyance | new installations and covered conveyances | covered conveyance installation, operation, repair, alteration, or inspection | ODOL rules apply ASME and related national standards to elevators and conveyances, with exemptions and inspection schedules in Chapter 70. | src:usa-ok:odol-elevator-rules | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** OUBCC adopts base model codes by reference and publishes Oklahoma amendments/rules by code family. OUBCC's Codes & Rules page states that base model codes have been adopted by reference and amended at the state level, and that only state amendments are available on the OUBCC website for viewing and download.

**Where amendments are published:** OUBCC Codes & Rules pages and OUBCC rule PDFs for the corresponding code family.

**Amendment parsing status:** source_registry_built; detailed_amendment_extraction_not_started

### 5.2 State Amendment Sources

| Amendment Source ID | Code Family | Publication Path | Status | Notes |
| --- | --- | --- | --- | --- |
| src:usa-ok:oubcc-ibc-adoptions | Building | OUBCC IBC adoption page and linked permanent-rule PDFs | indexed_not_parsed | Current adoption page verified; individual amendment provisions not extracted. |
| src:usa-ok:oubcc-irc-adoptions | Residential | OUBCC IRC adoption page and linked permanent-rule PDFs | indexed_not_parsed | Current adoption page verified; individual amendment provisions not extracted. |
| src:usa-ok:oubcc-iebc-adoptions | Existing Building / Rehabilitation | OUBCC IEBC adoption page and linked permanent-rule PDFs | indexed_not_parsed | Current adoption page verified; individual amendment provisions not extracted. |
| src:usa-ok:oubcc-ifc-adoptions | Fire | OUBCC IFC adoption page and linked permanent-rule PDFs | indexed_not_parsed | Current adoption page verified; individual amendment provisions not extracted. |
| src:usa-ok:oubcc-ifgc-adoptions | Fuel Gas | OUBCC IFGC adoption page and linked permanent-rule PDFs | indexed_not_parsed | Current adoption page verified; individual amendment provisions not extracted. |
| src:usa-ok:oubcc-imc-adoptions | Mechanical | OUBCC IMC adoption page and linked permanent-rule PDFs | indexed_not_parsed | Current adoption page verified; individual amendment provisions not extracted. |
| src:usa-ok:oubcc-ipc-adoptions | Plumbing | OUBCC IPC adoption page and linked permanent-rule PDFs | indexed_not_parsed | Current adoption page verified; individual amendment provisions not extracted. |
| src:usa-ok:oubcc-nec-adoptions | Electrical | OUBCC NEC adoption page and linked permanent-rule PDFs | indexed_not_parsed | Current adoption page verified; individual amendment provisions not extracted. |

### 5.3 High-Impact State Amendments

No high-impact Oklahoma amendment was extracted in this pass. The report is suitable for adoption-edition and authority tracking, not amendment-level compliance analysis.

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-ok"
  model: "local enforcement of statewide minimum codes with state fallback/specialized roles"
  enforcing_entities:
    - "local jurisdictions"
    - "municipalities"
    - "other political subdivisions"
    - "Construction Industries Board for electrical, mechanical, plumbing, and fuel gas minimum codes in unincorporated areas"
    - "Oklahoma State Fire Marshal for commercial building code in unincorporated areas"
  required_officials:
    - "local building/code officials where the local jurisdiction administers code enforcement"
    - "CIB inspectors or enforcement personnel for listed trades in unincorporated areas"
    - "State Fire Marshal or designee for commercial building-code enforcement in unincorporated areas"
  state_reserved_activities:
    - "OUBCC adoption of statewide minimum residential and commercial construction codes"
    - "ODOL elevator and conveyance safety standards and inspection program"
  source_ids:
    - "src:usa-ok:oubcc-faq"
    - "src:usa-ok:oubcc-statute-act"
    - "src:usa-ok:odol-elevator-rules"
  verification_status: "partially_verified"
  confidence: 0.76
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-ok"
  model: "statewide minimum with higher local standards allowed; local-condition amendments require OUBCC approval"
  applies_to_code_families:
    - "residential and commercial construction codes adopted by OUBCC"
  approval_required: true
  approving_authority_id: "ahj:usa-ok:oubcc"
  filing_required: "unresolved"
  registry_exists: "statute directs OUBCC to maintain a website listing higher-standard local/state codes, but current operational registry coverage was not verified"
  registry_source_ids:
    - "src:usa-ok:oubcc-statute-act"
  legal_basis_source_ids:
    - "src:usa-ok:oubcc-statute-act"
    - "src:usa-ok:oubcc-faq"
  verification_status: "partially_verified"
  confidence: 0.68
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement and local amendment authority are separate in Oklahoma. OUBCC sets the statewide minimum standards. Local jurisdictions generally interpret and enforce those standards. Local governments may impose higher standards; however, local amendments for local conditions require Commission approval. The current administrative mechanics for filing, registry publication, and approval were not resolved because the older Chapter 15 procedure chapter is revoked.

### 6.4 Known Local Amendment Registries

| Registry | Authority | Coverage | Source IDs | Status | Notes |
| --- | --- | --- | --- | --- | --- |
| OUBCC website listing of higher-standard codes | OUBCC | state agencies, cities, and political subdivisions with higher standards | src:usa-ok:oubcc-statute-act | unresolved_operational_status | Statutory website-listing duty verified; current registry completeness and access path not verified. |
| Jurisdictional Entity On-line Reporting Program | OUBCC | permit fee/reporting data, not confirmed as amendment registry | src:usa-ok:oubcc-home; src:usa-ok:oubcc-codes-rules | not_an_amendment_registry_in_this_pass | Listed as an OUBCC resource, but not used as a local amendment source here. |

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendments were parsed in this pass.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: partial_model_only

Expected jurisdiction stack:

```text
Address
  -> State of Oklahoma
  -> County
  -> Municipality / unincorporated county
  -> Local code-enforcement jurisdiction or political subdivision
  -> Building AHJ
  -> Fire AHJ
  -> Trade-specific AHJs, including CIB for listed trades in unincorporated areas
  -> State Fire Marshal for commercial building code in unincorporated areas
  -> ODOL Elevator Safety, if the project includes covered elevators or conveyances
  -> Applicable OUBCC statewide minimum code adoption records
  -> Applicable local higher-standard or local-condition amendment records
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | not selected | none | statewide | unknown | pending |
| County | not selected | none | statewide | unknown | pending |
| Municipality | not selected | none | statewide | unknown | pending |
| Unincorporated area determination | not selected | none | statewide | unknown | pending |
| Fire District | not selected | none | statewide | unknown | pending |
| Special District | not selected | none | statewide | unknown | pending |

### 7.3 AHJ Contact Data

No AHJ contact data was populated for this pass. OUBCC public contact information exists on the OUBCC site but was not modeled as jurisdiction-level AHJ contact data.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Source Type | Title / Description | Issuer | Accessed | URL | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| src:usa-ok:oubcc-home | agency_page | Uniform Building Code Commission home page | Oklahoma Uniform Building Code Commission / Oklahoma.gov | 2026-06-26 | https://oklahoma.gov/oubcc.html | Used for OUBCC mission, history, publication path, 2026 code-change proposal / energy committee watch items. |
| src:usa-ok:oubcc-codes-rules | agency_page | Codes & Rules | Oklahoma Uniform Building Code Commission / Oklahoma.gov | 2026-06-26 | https://oklahoma.gov/oubcc/codes-and-rules.html | Used for adopted-code publication path and state-amendment publication note. |
| src:usa-ok:oubcc-statute-act | statute_pdf | Uniform Building Code Commission Act, Oklahoma Statute Title 59 §§ 1000.20-1000.30 | Oklahoma Uniform Building Code Commission / Oklahoma.gov | 2026-06-26 | https://oklahoma.gov/content/dam/ok/en/oubcc/documents/statute-and-updated-versions/2023%2011%2001%20Title%2059%20OUBCC%20Statute.pdf | OUBCC labels this as an unofficial copy; used because OSCN was not directly retrievable in this pass. |
| src:usa-ok:oubcc-faq | agency_page | OUBCC FAQs | Oklahoma Uniform Building Code Commission / Oklahoma.gov | 2026-06-26 | https://oklahoma.gov/oubcc/faqs.html | Used for local enforcement, local higher-standard explanation, and public-project minimum standards. |
| src:usa-ok:oubcc-ibc-adoptions | agency_page | International Building Code Adoptions | Oklahoma Uniform Building Code Commission / Oklahoma.gov | 2026-06-26 | https://oklahoma.gov/oubcc/codes-and-rules/international-building-code-adoptions.html | Used for IBC 2018 current-adoption row and effective date. |
| src:usa-ok:oubcc-irc-adoptions | agency_page | International Residential Code Adoptions | Oklahoma Uniform Building Code Commission / Oklahoma.gov | 2026-06-26 | https://oklahoma.gov/oubcc/codes-and-rules/international-residential-code-adoptions.html | Used for IRC 2018 current-adoption row and effective date. |
| src:usa-ok:oubcc-iebc-adoptions | agency_page | International Existing Building Code Adoptions | Oklahoma Uniform Building Code Commission / Oklahoma.gov | 2026-06-26 | https://oklahoma.gov/oubcc/codes-and-rules/international-existing-building-code-adoptions.html | Used for IEBC 2018 current-adoption row and effective date. |
| src:usa-ok:oubcc-ifc-adoptions | agency_page | International Fire Code Adoptions | Oklahoma Uniform Building Code Commission / Oklahoma.gov | 2026-06-26 | https://oklahoma.gov/oubcc/codes-and-rules/international-fire-code-adoptions.html | Used for IFC 2018 current-adoption row and effective date. |
| src:usa-ok:oubcc-ifgc-adoptions | agency_page | International Fuel Gas Code Adoptions | Oklahoma Uniform Building Code Commission / Oklahoma.gov | 2026-06-26 | https://oklahoma.gov/oubcc/codes-and-rules/international-fuel-gas-code-adoptions.html | Used for IFGC 2018 current-adoption row and effective date. |
| src:usa-ok:oubcc-imc-adoptions | agency_page | International Mechanical Code Adoptions | Oklahoma Uniform Building Code Commission / Oklahoma.gov | 2026-06-26 | https://oklahoma.gov/oubcc/codes-and-rules/international-mechanical-code-adoptions.html | Used for IMC 2018 current-adoption row and effective date. |
| src:usa-ok:oubcc-ipc-adoptions | agency_page | International Plumbing Code Adoptions | Oklahoma Uniform Building Code Commission / Oklahoma.gov | 2026-06-26 | https://oklahoma.gov/oubcc/codes-and-rules/international-plumbing-code-adoptions.html | Used for IPC 2018 current-adoption row and effective date. |
| src:usa-ok:oubcc-nec-adoptions | agency_page | National Electrical Code Adoptions | Oklahoma Uniform Building Code Commission / Oklahoma.gov | 2026-06-26 | https://oklahoma.gov/oubcc/codes-and-rules/national-electrical-code-adoptions.html | Used for NEC 2023 current-adoption row and effective date. |
| src:usa-ok:oubcc-proposed-rules | agency_page | Proposed Rules | Oklahoma Uniform Building Code Commission / Oklahoma.gov | 2026-06-26 | https://oklahoma.gov/oubcc/codes-and-rules/proposed-rules.html | Used for 2024-code rulemaking watch items. |
| src:usa-ok:oubcc-admin-2021-rules | rule_pdf | Title 748 Permanent Rules - Administrative Amended | Oklahoma Uniform Building Code Commission / Oklahoma.gov | 2026-06-26 | https://oklahoma.gov/content/dam/ok/en/oubcc/documents/rules/2021%2009%2014%20Permanent%20Rules%20-%20Adminstrative%20Amended.pdf | Used only to flag that prior Chapter 15 local-code procedure rules are revoked. |
| src:usa-ok:odol-elevator-rules | rule_pdf | Elevator Safety Act and Rules | Oklahoma Department of Labor / Oklahoma.gov | 2026-06-26 | https://oklahoma.gov/content/dam/ok/en/labor/documents/safety-and-health/safety-standards/elevators/Elevator%20Safety%20Act%20and%20Rules%209-15-22.pdf | Used for elevator/conveyance authority, adopted national standards, and inspection framework. |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-ok:oubcc-statute-act | unofficial_copy | The OUBCC-hosted statute PDF states that it is an unofficial copy. | Verify against OSCN or another official unannotated statute source before final production. |
| src:usa-ok:oubcc-admin-2021-rules | historical_rule_pdf | Used to identify revocation of prior Chapter 15 local-code procedure rules, not to establish current procedure. | Needs current Secretary of State/OAC confirmation if local-amendment procedure is production-critical. |
| src:usa-ok:odol-elevator-rules | moving_reference | The elevator rules adopt several standards by “latest edition” and “most current addenda,” which can shift without a fixed edition in the rule text. | Confirm currently enforced editions with ODOL before project-specific use. |
| src:usa-ok:oubcc-proposed-rules | pending_rulemaking | Proposed-rule page is not proof of final adoption. | Monitor final adopted rules, Oklahoma Register, and OUBCC adoption pages before updating current-code rows. |
| src:usa-ok:oubcc-ibc-adoptions; src:usa-ok:oubcc-irc-adoptions; src:usa-ok:oubcc-iebc-adoptions; src:usa-ok:oubcc-ifc-adoptions; src:usa-ok:oubcc-ifgc-adoptions; src:usa-ok:oubcc-imc-adoptions; src:usa-ok:oubcc-ipc-adoptions; src:usa-ok:oubcc-nec-adoptions | adoption_page_summary | Adoption pages summarize current and superseded rule actions; detailed amendment text is in linked PDFs, not fully parsed here. | Use for edition/effective-date tracking; parse PDFs before amendment-level reporting. |

### 8.3 Supplemental Sources

None used for report facts. Non-official search results were consulted only for discovery and were not used as report authority.

### 8.4 Source Extraction Metadata

| Extraction ID | Source ID | Extracted Fields | Method | Extracted On | Notes |
| --- | --- | --- | --- | --- | --- |
| extract:usa-ok:001 | src:usa-ok:oubcc-statute-act | OUBCC creation, authority, minimum standards, higher local standards, Commission approval for local-condition amendments | web text extraction plus PDF screenshot spot-check | 2026-06-26 | Source caveat retained. |
| extract:usa-ok:002 | src:usa-ok:oubcc-faq | enforcement model, unincorporated-area CIB/State Fire Marshal allocation, local higher standards | web text extraction | 2026-06-26 | FAQ is official agency guidance but should be supplemented by CIB/OSFM rules for production. |
| extract:usa-ok:003 | src:usa-ok:oubcc-ibc-adoptions | IBC edition and effective date | web text extraction | 2026-06-26 | Current row inferred as unsuperseded top row. |
| extract:usa-ok:004 | src:usa-ok:oubcc-irc-adoptions | IRC edition and effective date | web text extraction | 2026-06-26 | Current row inferred as unsuperseded top row. |
| extract:usa-ok:005 | src:usa-ok:oubcc-iebc-adoptions | IEBC edition and effective date | web text extraction | 2026-06-26 | Current row inferred as unsuperseded top row. |
| extract:usa-ok:006 | src:usa-ok:oubcc-ifc-adoptions | IFC edition and effective date | web text extraction | 2026-06-26 | Current row inferred as unsuperseded top row. |
| extract:usa-ok:007 | src:usa-ok:oubcc-ifgc-adoptions | IFGC edition and effective date | web text extraction | 2026-06-26 | Current row inferred as unsuperseded top row. |
| extract:usa-ok:008 | src:usa-ok:oubcc-imc-adoptions | IMC edition and effective date | web text extraction | 2026-06-26 | Current row inferred as unsuperseded top row. |
| extract:usa-ok:009 | src:usa-ok:oubcc-ipc-adoptions | IPC edition and effective date | web text extraction | 2026-06-26 | Current row inferred as unsuperseded top row. |
| extract:usa-ok:010 | src:usa-ok:oubcc-nec-adoptions | NEC edition and effective date | web text extraction | 2026-06-26 | Current row inferred as unsuperseded top row. |
| extract:usa-ok:011 | src:usa-ok:oubcc-proposed-rules | 2024-code rulemaking watch items | web text extraction | 2026-06-26 | Does not update current-code matrix. |
| extract:usa-ok:012 | src:usa-ok:odol-elevator-rules | elevator authority and adopted national standards | web text extraction plus PDF screenshot spot-check | 2026-06-26 | Moving-reference standards noted. |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 1.00 | none | Core authority and main adoption rows are source-backed; unresolved rows remain explicit. |
| report | risk.overall_confidence | 0.72 | verified | 1.00 | none | Confidence reflects strong core adoption evidence with unresolved energy/accessibility/local procedure details. |
| ahj:usa-ok:oubcc | authority.name | Oklahoma Uniform Building Code Commission | verified | 0.95 | src:usa-ok:oubcc-home; src:usa-ok:oubcc-statute-act | Primary statewide authority. |
| ahj:usa-ok:oubcc | legal_basis | 59 O.S. §§ 1000.20-1000.30 | partially_verified | 0.85 | src:usa-ok:oubcc-statute-act | Source is OUBCC-hosted unofficial statute copy; needs OSCN confirmation. |
| adoption:usa-ok:ibc-2018 | edition/effective_date | IBC 2018 / 2021-09-14 | verified | 0.88 | src:usa-ok:oubcc-ibc-adoptions | OUBCC adoption page. |
| adoption:usa-ok:irc-2018 | edition/effective_date | IRC 2018 / 2022-09-14 | verified | 0.88 | src:usa-ok:oubcc-irc-adoptions | OUBCC adoption page. |
| adoption:usa-ok:iebc-2018 | edition/effective_date | IEBC 2018 / 2021-09-14 | verified | 0.88 | src:usa-ok:oubcc-iebc-adoptions | OUBCC adoption page. |
| adoption:usa-ok:ifc-2018 | edition/effective_date | IFC 2018 / 2021-09-14 | verified | 0.88 | src:usa-ok:oubcc-ifc-adoptions | OUBCC adoption page. |
| adoption:usa-ok:ifgc-2018 | edition/effective_date | IFGC 2018 / 2021-09-14 | verified | 0.88 | src:usa-ok:oubcc-ifgc-adoptions | OUBCC adoption page. |
| adoption:usa-ok:imc-2018 | edition/effective_date | IMC 2018 / 2021-09-14 | verified | 0.88 | src:usa-ok:oubcc-imc-adoptions | OUBCC adoption page. |
| adoption:usa-ok:ipc-2018 | edition/effective_date | IPC 2018 / 2021-09-14 | verified | 0.88 | src:usa-ok:oubcc-ipc-adoptions | OUBCC adoption page. |
| adoption:usa-ok:nec-2023 | edition/effective_date | NEC 2023 / 2024-09-14 | verified | 0.88 | src:usa-ok:oubcc-nec-adoptions | OUBCC adoption page. |
| local-enforcement:usa-ok | model | local enforcement of statewide minimum codes with state fallback/specialized roles | partially_verified | 0.76 | src:usa-ok:oubcc-faq; src:usa-ok:oubcc-statute-act | Requires CIB/OSFM source confirmation for production. |
| local-amendment-rule:usa-ok | model | statewide minimum with higher local standards allowed; local-condition amendments require OUBCC approval | partially_verified | 0.68 | src:usa-ok:oubcc-statute-act; src:usa-ok:oubcc-admin-2021-rules | Procedure unresolved after Chapter 15 revocation. |
| ahj:usa-ok:odol | elevator standards | ASME/NFPA/IBC standards adopted by moving reference | partially_verified | 0.72 | src:usa-ok:odol-elevator-rules | Verify current editions with ODOL before project-level use. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Every `src:usa-ok:*` cited in the body appears in section 8. |
| All authority IDs resolve | pass | OUBCC, CIB, State Fire Marshal, and ODOL IDs are used consistently; CIB/OSFM details are limited to FAQ-backed edges. |
| All current code families have adoption records | fail | Main OUBCC code-family rows have adoption records; energy, accessibility, and elevator rows require additional normalized records or scope decisions. |
| Building and operational fire code are separated | pass | Fire construction references and operational/prevention code rows are separated, with enforcement-scope caveat. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Unsupported operative and mandatory dates remain null. |
| Effective dates are valid ISO dates | pass | Verified effective dates use YYYY-MM-DD format. |
| No impossible date sequences | pass | No unsupported later/earlier relationships were introduced. |
| Transition rules have explicit trigger conditions | pass | Known rules are framed as effective-date rules; permit-trigger logic remains open. |
| Permit-date logic is captured where applicable | fail | No permit-application or permit-issuance transition rule was verified. |
| Local enforcement model classified | pass | Classified with source-backed local and unincorporated-area split. |
| Local amendment rule classified | fail | Legal posture is classified, but current filing/approval procedure remains unresolved. |
| AHJ confirmation metadata present | fail | No jurisdiction-specific AHJ contacts or boundary datasets were entered. |
| Official-source caveats captured | pass | Unofficial statute copy, moving elevator standards, proposed-rule limitations, and adoption-page summary limitations are captured. |
| Leftover template markers removed | pass | The file avoids the template-placeholder patterns checked by the standard validation command. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-ok:001 | high | energy code | Standalone statewide commercial energy-code adoption was not resolved from the OUBCC adoption pages. | Review OAC, OUBCC rule PDFs, State Energy Office materials, and DOE/IECC determinations; distinguish residential IRC energy provisions from commercial IECC adoption. | null | null | open |
| issue:usa-ok:002 | high | 2024 code package | OUBCC proposed-rule page documents 2024-code rulemaking, but final adoption and effective dates were not verified. | Check final adopted rules, Oklahoma Register permanent adoption notices, and OUBCC code-adoption pages. | null | null | open |
| issue:usa-ok:003 | medium | local amendments | Statute verifies higher standards and Commission approval for local-condition amendments, but current filing/approval procedure after Chapter 15 revocation is unresolved. | Review current OAC Title 748, agency forms, Commission agendas, and registry/publication process. | null | null | open |
| issue:usa-ok:004 | medium | operational fire AHJ | OUBCC adopts IFC and FAQ describes State Fire Marshal enforcement in unincorporated areas, but OSFM statutory/rule authority was not independently parsed. | Extract OSFM statutes/rules and compare with OUBCC/local enforcement model. | null | null | open |
| issue:usa-ok:005 | medium | accessibility | Accessibility standard edition and state/federal interaction were not independently parsed. | Extract IBC/IRC/IEBC referenced standards and Oklahoma amendment PDFs; note ADA/ABA/FHA overlay separately. | null | null | open |
| issue:usa-ok:006 | low | elevator edition chronology | ODOL rules use moving references to latest editions/current addenda; exact edition-effective chronology needs confirmation. | Check ODOL current notices, rules, and any municipal delegated programs. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-ok:oubcc-home | src:usa-ok:oubcc-home | html_diff | monthly | OUBCC announces code-change proposal windows, technical committee actions, or new adopted-code status | 2026-06-26 | active |
| watch:usa-ok:oubcc-codes-rules | src:usa-ok:oubcc-codes-rules | html_diff | monthly | new code adoption page, new amendment PDF, or page restructuring | 2026-06-26 | active |
| watch:usa-ok:oubcc-proposed-rules | src:usa-ok:oubcc-proposed-rules | html_diff | biweekly_while_rulemaking_active | 2024-code proposed rules move to final adoption or receive effective dates | 2026-06-26 | active |
| watch:usa-ok:oubcc-nec | src:usa-ok:oubcc-nec-adoptions | html_diff | monthly | NEC 2027 review or adoption activity appears | 2026-06-26 | active |
| watch:usa-ok:odol-elevator | src:usa-ok:odol-elevator-rules | pdf_or_html_diff | quarterly | elevator standards, inspection rules, or delegated municipal program rules change | 2026-06-26 | active |
| watch:usa-ok:statute-act | src:usa-ok:oubcc-statute-act | statute_check | quarterly | amendments to 59 O.S. §§ 1000.20-1000.30 or related local-amendment procedures | 2026-06-26 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-24 | Baseline Oklahoma draft existed before this pass | report:usa-ok | none | prior draft | Generic unresolved fields were present. |
| 2026-06-26 | Populated authority model, source registry, OUBCC adoption matrix, local enforcement model, local amendment posture, pending-rulemaking watch items, and QA/open-issue sections | report:usa-ok; ahj:usa-ok:oubcc; adoption:usa-ok:ibc-2018; adoption:usa-ok:irc-2018; adoption:usa-ok:iebc-2018; adoption:usa-ok:ifc-2018; adoption:usa-ok:ifgc-2018; adoption:usa-ok:imc-2018; adoption:usa-ok:ipc-2018; adoption:usa-ok:nec-2023; local-enforcement:usa-ok; local-amendment-rule:usa-ok | src:usa-ok:oubcc-home; src:usa-ok:oubcc-codes-rules; src:usa-ok:oubcc-statute-act; src:usa-ok:oubcc-faq; src:usa-ok:oubcc-ibc-adoptions; src:usa-ok:oubcc-irc-adoptions; src:usa-ok:oubcc-iebc-adoptions; src:usa-ok:oubcc-ifc-adoptions; src:usa-ok:oubcc-ifgc-adoptions; src:usa-ok:oubcc-imc-adoptions; src:usa-ok:oubcc-ipc-adoptions; src:usa-ok:oubcc-nec-adoptions; src:usa-ok:oubcc-proposed-rules; src:usa-ok:oubcc-admin-2021-rules; src:usa-ok:odol-elevator-rules | ChatGPT | Upgraded status to partially_verified because core authority and main adoption facts are source-backed; unresolved rows remain explicit. |
