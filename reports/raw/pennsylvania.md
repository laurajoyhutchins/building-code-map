---
state:
  state_id: "US-PA"
  name: "Pennsylvania"
  abbreviation: "PA"
report:
  report_id: "state-report:usa-pa"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-26"
  last_verified: null
  reviewed_by: null
risk:
  overall_confidence: 0.72 # 0.00 - 1.00
  risk_flags:
    - "operational_fire_code_authority_unresolved"
    - "standalone_electrical_code_authority_unresolved"
    - "accessibility_act_16_2026_reconciliation_needed"
    - "ahj_boundary_sources_not_selected"
    - "full_amendment_diff_not_parsed"
  open_questions_count: 6
---

# State Building Code Authority Report: Pennsylvania

## 1. Executive Summary

- **Authority model:** Pennsylvania administers a statewide Uniform Construction Code (UCC). The Department of Labor & Industry (DLI), through the UCC program in the Bureau of Occupational and Industrial Safety, is the primary statewide construction-code authority. The Uniform Construction Code Review and Advisory Council participates in the triennial review process, and DLI promulgates final-omitted regulations adopting Council decisions as required by the Pennsylvania Construction Code Act. Source IDs: `src:usa-pa:pcca-act-45`, `src:usa-pa:dli-ucc-home`, `src:usa-pa:dli-ucc-rac-2021-cycle`.

- **Statewide code status:** Pennsylvania's current UCC is based primarily on selected 2021 ICC model codes. The 2021 triennial update was approved by IRRC on `2025-10-16`, published as final-omitted rulemaking in the Pennsylvania Bulletin on `2025-11-08`, and became effective on `2026-01-01`. Source IDs: `src:usa-pa:dli-ucc-rac-2021-cycle`, `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted`, `src:usa-pa:pa-code-403-21`.

- **Local enforcement model:** Municipalities may administer and enforce the UCC if they adopt the current UCC as their municipal building code. A municipality may use its own code officials, certified third-party agencies, intermunicipal arrangements, another municipality, or an agreement with DLI for nonresidential and other structures. Where a municipality does not administer and enforce, the Construction Code Act and regulations provide third-party and DLI pathways depending on project type. Source IDs: `src:usa-pa:pcca-act-45`, `src:usa-pa:pa-code-403-102`.

- **Local amendment posture:** The UCC preempts local construction codes except as provided by statute. Local technical changes may be adopted only if they equal or exceed the UCC and follow the required hearing, notice, filing, and DLI-review process. DLI maintains a municipal code-change ordinance page that functions as a state registry for reviewed local changes. Source IDs: `src:usa-pa:pcca-act-45`, `src:usa-pa:pa-code-403-102`, `src:usa-pa:dli-municipal-code-change-ordinances`.

- **Known transition periods or pending changes:** The 2021 UCC update became effective `2026-01-01`. A contract-signed-before-effective-date transition rule can allow the former code if the permit application is filed within six months after the effective date or within a shorter municipal ordinance period. Act 16 of 2026 amended the Construction Code Act's accessibility provisions and should be reconciled against future UCC accessibility regulations. Source IDs: `src:usa-pa:pcca-act-45`, `src:usa-pa:pa-code-403-1`, `src:usa-pa:act-16-2026`.

- **Production readiness:** partial

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: DLI is the statewide UCC agency; the Construction Code Act requires DLI
    to promulgate regulations adopting Council decisions in the triennial review process.
  confidence: 0.86
  source_ids:
  - src:usa-pa:pcca-act-45
  - src:usa-pa:dli-ucc-rac-2021-cycle
- topic: Primary building code edition
  finding:
  - '2026-01-01'
  confidence: 0.91
  source_ids:
  - src:usa-pa:pa-code-403-21
  - src:usa-pa:pa-bulletin-55-45-ucc-final-omitted
- topic: Residential code edition
  finding: The current UCC adopts the 2021 International Residential Code with Pennsylvania
    exclusions and amendments.
  confidence: 0.86
  source_ids:
  - src:usa-pa:pa-code-403-21
  - src:usa-pa:pa-bulletin-55-45-ucc-final-omitted
- topic: Electrical code authority
  finding: A standalone statewide NEC adoption was not established from the sources
    reviewed; electrical requirements appear through adopted ICC provisions and referenced
    standards where applicable.
  confidence: 0.42
  source_ids:
  - src:usa-pa:pa-code-403-21
- topic: Fire code authority
  finding: Pennsylvania adopts the 2021 International Fire Code only to the extent
    it is referenced by other adopted ICC codes; separate operational fire-prevention
    authority remains unresolved.
  confidence: 0.57
  source_ids:
  - src:usa-pa:pa-code-403-21
  - src:usa-pa:pa-bulletin-55-45-ucc-final-omitted
- topic: Local enforcement
  finding: Municipalities may opt into UCC administration and enforcement; non-administering
    municipalities use statutory third-party or DLI pathways by project type.
  confidence: 0.84
  source_ids:
  - src:usa-pa:pcca-act-45
  - src:usa-pa:pa-code-403-102
- topic: Local amendments
  finding: Local technical changes must equal or exceed the UCC and follow the hearing,
    notice, filing, and DLI-review process.
  confidence: 0.82
  source_ids:
  - src:usa-pa:pcca-act-45
  - src:usa-pa:pa-code-403-102
  - src:usa-pa:dli-municipal-code-change-ordinances
- topic: Effective / transition rule
  finding:
  - '2026-01-01'
  confidence: 0.81
  source_ids:
  - src:usa-pa:dli-ucc-rac-2021-cycle
  - src:usa-pa:pcca-act-45
  - src:usa-pa:pa-code-403-1
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | `ahj:usa-pa:dli-bois-ucc` |
| Authority name | Pennsylvania Department of Labor & Industry, Bureau of Occupational and Industrial Safety, Uniform Construction Code program |
| Authority type | state executive agency |
| Legal basis | Pennsylvania Construction Code Act, Act 45 of 1999, as amended; 34 Pa. Code Part XIV, Uniform Construction Code |
| Role | Statewide UCC administration; promulgation of UCC regulations; triennial-code update implementation; oversight of certification and local/third-party enforcement pathways; selected plan-review and inspection roles where DLI jurisdiction applies |
| Enforcement model | mixed state-local-third-party |
| Source IDs | `src:usa-pa:pcca-act-45`; `src:usa-pa:dli-ucc-home`; `src:usa-pa:dli-ucc-regs-statutes`; `src:usa-pa:pa-code-403-102` |
| Verification status | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | `ahj:usa-pa:dli-bois-ucc` | Pennsylvania Department of Labor & Industry, UCC program | State UCC adoption and oversight; local/third-party enforcement pathways | Construction Code Act; 34 Pa. Code §§ 403.1, 403.21, 403.102 | `src:usa-pa:pcca-act-45`; `src:usa-pa:pa-code-403-1`; `src:usa-pa:pa-code-403-21`; `src:usa-pa:pa-code-403-102` | partially_verified |
| Residential | `ahj:usa-pa:dli-bois-ucc` | Pennsylvania Department of Labor & Industry, UCC program | State UCC adoption and oversight for one- and two-family dwellings and townhouses | Construction Code Act; 34 Pa. Code §§ 403.1, 403.21, 403.63, 403.102 | `src:usa-pa:pcca-act-45`; `src:usa-pa:pa-code-403-21`; `src:usa-pa:pa-code-403-63`; `src:usa-pa:pa-code-403-102` | partially_verified |
| Existing Building / Rehabilitation | `ahj:usa-pa:dli-bois-ucc` | Pennsylvania Department of Labor & Industry, UCC program | State adoption of existing-building compliance paths | 34 Pa. Code § 403.21 | `src:usa-pa:pa-code-403-21`; `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted` | partially_verified |
| Mechanical | `ahj:usa-pa:dli-bois-ucc` | Pennsylvania Department of Labor & Industry, UCC program | State UCC adoption and oversight | 34 Pa. Code § 403.21 | `src:usa-pa:pa-code-403-21`; `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted` | partially_verified |
| Plumbing | `ahj:usa-pa:dli-bois-ucc` | Pennsylvania Department of Labor & Industry, UCC program | State UCC adoption and oversight, subject to second-class-county local-health plumbing caveat | 34 Pa. Code § 403.21; Construction Code Act local-health caveat | `src:usa-pa:pa-code-403-21`; `src:usa-pa:pcca-act-45` | partially_verified |
| Fuel Gas | `ahj:usa-pa:dli-bois-ucc` | Pennsylvania Department of Labor & Industry, UCC program | State UCC adoption and oversight for fuel gas code provisions | 34 Pa. Code § 403.21 | `src:usa-pa:pa-code-403-21`; `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted` | partially_verified |
| Electrical | `ahj:usa-pa:dli-bois-ucc` | Pennsylvania Department of Labor & Industry, UCC program | Electrical requirements are present through adopted ICC code provisions and referenced standards where applicable; standalone statewide NEC adoption not established | 34 Pa. Code § 403.21 | `src:usa-pa:pa-code-403-21` | unresolved |
| Energy | `ahj:usa-pa:dli-bois-ucc` | Pennsylvania Department of Labor & Industry, UCC program | State UCC adoption and oversight for energy provisions and approved compliance software paths | 34 Pa. Code § 403.21 | `src:usa-pa:pa-code-403-21`; `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted` | partially_verified |
| Fire - construction references | `ahj:usa-pa:dli-bois-ucc` | Pennsylvania Department of Labor & Industry, UCC program | IFC provisions apply only to the extent referenced by other adopted ICC construction codes | 34 Pa. Code § 403.21 | `src:usa-pa:pa-code-403-21`; `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted` | partially_verified |
| Fire - operational / prevention code | `ahj:usa-pa:fire-operational-unresolved` | Unresolved | Separate operational fire-prevention code authority was not established from sources reviewed | unresolved | `src:usa-pa:pa-code-403-21` | unresolved |
| Accessibility | `ahj:usa-pa:dli-accessibility-ucc` | Pennsylvania Department of Labor & Industry, accessibility functions under the UCC | State accessibility oversight; Secretary and Accessibility Advisory Board roles require post-Act-16 regulatory reconciliation | Construction Code Act; Act 16 of 2026; 34 Pa. Code § 403.21 | `src:usa-pa:pcca-act-45`; `src:usa-pa:act-16-2026`; `src:usa-pa:pa-code-403-21`; `src:usa-pa:pa-code-403-102` | partially_verified |
| Elevator / Conveyance | `ahj:usa-pa:dli-bois-ucc` | Pennsylvania Department of Labor & Industry, UCC program | IBC elevator/conveyance construction provisions are included in adopted IBC scope; separate elevator program rules were not parsed | 34 Pa. Code § 403.21 | `src:usa-pa:pa-code-403-21` | partially_verified |

### 2.3 Authority Hierarchy Notes

Pennsylvania's UCC model is statewide in adoption but decentralized in administration. DLI is the statewide construction-code authority and the UCC is adopted by state regulation. Municipalities choose whether to administer and enforce the current UCC locally. If they do, they must adopt the current UCC by ordinance and may use municipal code officials, third-party agencies, intermunicipal cooperation, or DLI agreements. If they do not, owners and applicants use the statutory third-party or DLI pathways. Source IDs: `src:usa-pa:pcca-act-45`; `src:usa-pa:pa-code-403-102`.

Local enforcement is separate from local amendment power. A municipality may administer the UCC without creating a local technical amendment. Local technical amendments or code changes must equal or exceed the UCC and follow the state review and filing process. Source IDs: `src:usa-pa:pcca-act-45`; `src:usa-pa:pa-code-403-102`; `src:usa-pa:dli-municipal-code-change-ordinances`.

Operational fire-prevention authority remains unresolved. The current construction-code adoption incorporates the 2021 IFC only to the extent it is referenced by other adopted ICC codes, so this report does not treat the IFC row as a verified statewide operational fire code. Source IDs: `src:usa-pa:pa-code-403-21`.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| `edge:usa-pa:001` | `ahj:usa-pa:dli-bois-ucc` | adopts_by_regulation | Uniform Construction Code / incorporated model codes | `src:usa-pa:pcca-act-45`; `src:usa-pa:pa-code-403-21` | partially_verified |
| `edge:usa-pa:002` | `ahj:usa-pa:dli-bois-ucc` | implements_review_cycle_from | Uniform Construction Code Review and Advisory Council decisions / triennial update process | `src:usa-pa:pcca-act-45`; `src:usa-pa:dli-ucc-rac-2021-cycle` | partially_verified |
| `edge:usa-pa:003` | `ahj:usa-pa:dli-bois-ucc` | permits_local_administration_by | municipalities adopting the current UCC as their municipal building code | `src:usa-pa:pcca-act-45`; `src:usa-pa:pa-code-403-102` | partially_verified |
| `edge:usa-pa:004` | `ahj:usa-pa:dli-bois-ucc` | fallback_enforcement_for | structures or municipalities covered by statutory DLI/third-party pathways | `src:usa-pa:pcca-act-45`; `src:usa-pa:pa-code-403-102` | partially_verified |
| `edge:usa-pa:005` | local municipalities | may_adopt_equal_or_more_stringent_changes_after | public hearing, notice, filing, and DLI review | `src:usa-pa:pcca-act-45`; `src:usa-pa:pa-code-403-102`; `src:usa-pa:dli-municipal-code-change-ordinances` | partially_verified |
| `edge:usa-pa:006` | `ahj:usa-pa:dli-accessibility-ucc` | retains_or_oversees | accessibility modifications and accessibility enforcement issues requiring certified officials or DLI involvement | `src:usa-pa:pcca-act-45`; `src:usa-pa:act-16-2026`; `src:usa-pa:pa-code-403-102` | partially_verified |
| `edge:usa-pa:007` | `ahj:usa-pa:fire-operational-unresolved` | unresolved_relationship_to | statewide operational fire-prevention code, if any | `src:usa-pa:pa-code-403-21` | unresolved |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | Pennsylvania Uniform Construction Code | International Building Code | 2021, selected chapters and Pennsylvania amendments | current | 2025-10-16 | 2026-01-01 | 2026-01-01 | 2026-01-01 | IRRC approval used as adoption date; final-omitted rule published 2025-11-08; qualifying pre-effective contracts may use prior code only within statutory or shorter local transition window | `src:usa-pa:dli-ucc-rac-2021-cycle`; `src:usa-pa:pa-code-403-21`; `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted`; `src:usa-pa:pcca-act-45` |
| Residential | Pennsylvania Uniform Construction Code | International Residential Code | 2021, with Pennsylvania amendments | current | 2025-10-16 | 2026-01-01 | 2026-01-01 | 2026-01-01 | Same 2021 UCC triennial transition rule | `src:usa-pa:dli-ucc-rac-2021-cycle`; `src:usa-pa:pa-code-403-21`; `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted`; `src:usa-pa:pcca-act-45` |
| Existing Building / Rehabilitation | Pennsylvania Uniform Construction Code | International Existing Building Code and IBC Chapter 34 compliance path | 2021 IEBC; 2021 IBC Chapter 34 | current | 2025-10-16 | 2026-01-01 | 2026-01-01 | 2026-01-01 | Existing-building work may comply with IBC Chapter 34 or the 2021 IEBC, subject to UCC amendments and transition rules | `src:usa-pa:pa-code-403-21`; `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted` |
| Mechanical | Pennsylvania Uniform Construction Code | International Mechanical Code | 2021 | current | 2025-10-16 | 2026-01-01 | 2026-01-01 | 2026-01-01 | Same 2021 UCC triennial transition rule | `src:usa-pa:dli-ucc-rac-2021-cycle`; `src:usa-pa:pa-code-403-21`; `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted` |
| Plumbing | Pennsylvania Uniform Construction Code | International Plumbing Code | 2021 | current_with_caveat | 2025-10-16 | 2026-01-01 | 2026-01-01 | 2026-01-01 | Same 2021 UCC triennial transition rule; county-of-second-class local-health plumbing-code caveat applies where statutory conditions are met | `src:usa-pa:dli-ucc-rac-2021-cycle`; `src:usa-pa:pa-code-403-21`; `src:usa-pa:pcca-act-45`; `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted` |
| Fuel Gas | Pennsylvania Uniform Construction Code | International Fuel Gas Code | 2021 | current | 2025-10-16 | 2026-01-01 | 2026-01-01 | 2026-01-01 | Same 2021 UCC triennial transition rule | `src:usa-pa:dli-ucc-rac-2021-cycle`; `src:usa-pa:pa-code-403-21`; `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted` |
| Electrical | Pennsylvania Uniform Construction Code electrical provisions where embedded or referenced | Standalone NEC adoption unresolved | unresolved | unresolved | null | null | null | null | Electrical-specific statewide adoption and authority require additional verification; do not infer a separate NEC edition from this report | `src:usa-pa:pa-code-403-21` |
| Energy | Pennsylvania Uniform Construction Code | International Energy Conservation Code | 2021 | current | 2025-10-16 | 2026-01-01 | 2026-01-01 | 2026-01-01 | IECC compliance may use stated Pennsylvania/DOE compliance paths where allowed by § 403.21; same 2021 UCC transition rule applies | `src:usa-pa:dli-ucc-rac-2021-cycle`; `src:usa-pa:pa-code-403-21`; `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted` |
| Fire - construction references | Pennsylvania Uniform Construction Code | International Fire Code | 2021, only to extent referenced by adopted ICC codes | current_limited_reference | 2025-10-16 | 2026-01-01 | 2026-01-01 | 2026-01-01 | IFC is not treated here as a verified operational fire-prevention code; use only as construction-code reference unless separately verified | `src:usa-pa:pa-code-403-21`; `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted` |
| Fire - operational / prevention code | unresolved | unresolved | unresolved | unresolved | null | null | null | null | Statewide operational fire-code authority and edition remain open | `src:usa-pa:pa-code-403-21` |
| Accessibility | Pennsylvania UCC accessibility provisions | UCC accessibility provisions and referenced model-code provisions, subject to Pennsylvania exclusions and Act 16 of 2026 | partially parsed | partially_verified | 2025-10-16 | 2026-01-01 | 2026-01-01 | 2026-01-01 | Accessibility provisions require post-Act-16 review before treating as fully normalized | `src:usa-pa:pcca-act-45`; `src:usa-pa:act-16-2026`; `src:usa-pa:pa-code-403-21`; `src:usa-pa:pa-code-403-102` |
| Elevator / Conveyance | Pennsylvania Uniform Construction Code | International Building Code elevator/conveyance provisions | 2021 IBC provisions, selected scope | partially_verified | 2025-10-16 | 2026-01-01 | 2026-01-01 | 2026-01-01 | IBC construction provisions were verified; separate conveyance-program rules were not parsed | `src:usa-pa:pa-code-403-21` |

### 3.2 Adoption Records

#### `adopt:usa-pa:ucc-2021-triennial`

| Field | Value |
| --- | --- |
| Applies to | Building, Residential, Mechanical, Plumbing, Fuel Gas, Energy, Existing Building, construction-fire references, selected accessibility and conveyance provisions |
| State code name | Pennsylvania Uniform Construction Code |
| Base adoption action | 2021 UCC triennial update; final-omitted rulemaking amending 34 Pa. Code Chapters 401 and 403 |
| Adoption Date | 2025-10-16 |
| Publication Date | 2025-11-08 |
| Effective Date | 2026-01-01 |
| Operative Date | 2026-01-01 |
| Mandatory Date | 2026-01-01 |
| Transition Rule | Qualifying pre-effective contracts may use the former code only if the permit application is filed within six months after the effective date or within a shorter municipal ordinance period; permits issued before the new regulation remain valid subject to permit-validity rules |
| Source IDs | `src:usa-pa:dli-ucc-rac-2021-cycle`; `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted`; `src:usa-pa:pcca-act-45`; `src:usa-pa:pa-code-403-1`; `src:usa-pa:pa-code-403-21` |
| Verification status | partially_verified |

#### `adopt:usa-pa:ifc-2021-reference-only`

| Field | Value |
| --- | --- |
| Applies to | Fire - construction references |
| State code name | Pennsylvania Uniform Construction Code |
| Base adoption action | 2021 International Fire Code included only to the extent referenced by other adopted ICC codes |
| Adoption Date | 2025-10-16 |
| Publication Date | 2025-11-08 |
| Effective Date | 2026-01-01 |
| Operative Date | 2026-01-01 |
| Mandatory Date | 2026-01-01 |
| Transition Rule | Do not use as a standalone statewide operational fire-prevention-code adoption without separate authority review |
| Source IDs | `src:usa-pa:pa-code-403-21`; `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted` |
| Verification status | partially_verified |

#### `adopt:usa-pa:iecc-2021-energy`

| Field | Value |
| --- | --- |
| Applies to | Energy |
| State code name | Pennsylvania Uniform Construction Code |
| Base adoption action | 2021 International Energy Conservation Code with Pennsylvania amendments and compliance-tool paths stated in § 403.21 |
| Adoption Date | 2025-10-16 |
| Publication Date | 2025-11-08 |
| Effective Date | 2026-01-01 |
| Operative Date | 2026-01-01 |
| Mandatory Date | 2026-01-01 |
| Transition Rule | Same 2021 UCC triennial transition rule |
| Source IDs | `src:usa-pa:pa-code-403-21`; `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted` |
| Verification status | partially_verified |

#### `adopt:usa-pa:electrical-standalone-unresolved`

| Field | Value |
| --- | --- |
| Applies to | Electrical |
| State code name | unresolved |
| Base adoption action | Standalone statewide electrical/NEC adoption not established from the official sources reviewed |
| Adoption Date | null |
| Publication Date | null |
| Effective Date | null |
| Operative Date | null |
| Mandatory Date | null |
| Transition Rule | unresolved |
| Source IDs | `src:usa-pa:pa-code-403-21` |
| Verification status | unresolved |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

The current 2021 UCC triennial update became effective on `2026-01-01`. The primary statewide transition rule distinguishes the regulation effective date from a limited grandfathering pathway for projects with a design or construction contract signed before that effective date. For those qualifying projects, the permit application must be filed within six months after the effective date, or within the shorter period established by municipal ordinance. Source IDs: `src:usa-pa:dli-ucc-rac-2021-cycle`; `src:usa-pa:pcca-act-45`; `src:usa-pa:pa-code-403-1`.

Permit-processing and permit-validity rules remain separate from adoption-effective dates. Commercial permit review is generally tied to a 30-business-day review period. Residential permit review is generally tied to a 15-business-day review period, with a shorter 5-business-day period when the application includes the required licensed-design-professional certification. Issued permits become invalid if work does not begin within 180 days of issuance or if work is abandoned or suspended for 180 days after beginning; a permit may not remain valid for more than five years from issuance. Source IDs: `src:usa-pa:pa-code-403-43`; `src:usa-pa:pa-code-403-63`.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `date-rule:usa-pa:001` | 2021 UCC triennial update | effective_date | 2026-01-01 | Final-omitted UCC rulemaking effective date | no, except qualifying transition cases | `src:usa-pa:dli-ucc-rac-2021-cycle`; `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted` | verified |
| `date-rule:usa-pa:002` | Buildings or renovations under pre-effective design or construction contracts | contract_grandfathering | six months after 2026-01-01, subject to shorter municipal ordinance period | Design or construction contract signed before effective date and permit application filed within allowed period | yes | `src:usa-pa:pcca-act-45`; `src:usa-pa:pa-code-403-1` | partially_verified |
| `date-rule:usa-pa:003` | Permits issued before latest UCC update | prior_permit_validity | permit-specific | Construction permit issued before effective date of latest regulation | yes, under the issued permit | `src:usa-pa:pcca-act-45`; `src:usa-pa:pa-code-403-1` | partially_verified |
| `date-rule:usa-pa:004` | Commercial construction permits | review_clock | 30 business days | Code administrator receives permit application | not applicable | `src:usa-pa:pa-code-403-43` | partially_verified |
| `date-rule:usa-pa:005` | Residential construction permits | review_clock | 15 business days or 5 business days with qualifying design-professional certification | Code administrator receives permit application | not applicable | `src:usa-pa:pa-code-403-63` | partially_verified |
| `date-rule:usa-pa:006` | Issued permits | permit_validity | 180 days to begin work; 180-day abandonment/suspension trigger; five-year maximum permit validity | Permit issued and work has not begun, or work begins and is abandoned or suspended | not applicable | `src:usa-pa:pa-code-403-43`; `src:usa-pa:pa-code-403-63` | partially_verified |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Accessibility | Revised or successor accessibility-code process under Act 16 of 2026 | 2026-05-06 | null | null | null | null | active_watch | `src:usa-pa:act-16-2026`; `src:usa-pa:pcca-act-45` | Act 16 amended the Construction Code Act; the report should be updated when DLI publishes implementing accessibility-rule changes or guidance. |
| All UCC code families | Next triennial ICC-code review cycle | null | null | null | null | null | active_watch | `src:usa-pa:pcca-act-45`; `src:usa-pa:dli-ucc-rac-2021-cycle` | The Construction Code Act and DLI materials describe triennial review; no next adoption action was normalized in this report. |
| Fire - operational / prevention code | unresolved | null | null | null | null | null | unresolved | `src:usa-pa:pa-code-403-21` | Operational fire-prevention code authority remains an open issue. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| `applicability-rule:usa-pa:001` | Plumbing | County of the second class local-health plumbing code | A county of the second class has adopted a plumbing code under the Local Health Administration Law and the code meets UCC minimum requirements | IPC/UCC plumbing provisions have a local-health caveat in second-class counties; do not assume the statewide IPC rule is the only plumbing authority for those jurisdictions | `src:usa-pa:pa-code-403-21`; `src:usa-pa:pcca-act-45` | partially_verified |
| `applicability-rule:usa-pa:002` | Existing Building / Rehabilitation | Existing-building repair, alteration, change of occupancy, addition, and relocation | Existing-building work within UCC scope | Compliance may use 2021 IBC Chapter 34 or the 2021 IEBC, subject to Pennsylvania amendments and exclusions | `src:usa-pa:pa-code-403-21`; `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted` | partially_verified |
| `applicability-rule:usa-pa:003` | Energy | IECC compliance demonstration | Applicant uses permitted prescriptive methods or compliance tools | § 403.21 identifies REScheck / Pennsylvania Alternative Residential Energy Provisions and COMcheck paths for applicable energy compliance demonstrations | `src:usa-pa:pa-code-403-21`; `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted` | partially_verified |
| `applicability-rule:usa-pa:004` | Fire - construction references | IFC referenced by adopted ICC codes | Other adopted ICC codes reference IFC provisions | IFC applies only to the extent referenced; this does not resolve a standalone operational fire-prevention code | `src:usa-pa:pa-code-403-21`; `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted` | partially_verified |
| `applicability-rule:usa-pa:005` | Accessibility | Accessibility modifications and enforcement responsibilities | Accessibility provisions apply or a modification/technical infeasibility issue arises | DLI retains significant accessibility authority; local accessibility enforcement depends on certified accessibility inspectors/plans examiners and Act 16 reconciliation | `src:usa-pa:pcca-act-45`; `src:usa-pa:act-16-2026`; `src:usa-pa:pa-code-403-102` | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** model codes are incorporated by reference into the UCC through 34 Pa. Code § 403.21 with Pennsylvania-specific exclusions, retained older provisions, appendix choices, scope limits, and compliance alternatives.

**Where amendments are published:** Pennsylvania Code, 34 Pa. Code Chapters 401 and 403, and Pennsylvania Bulletin rulemaking publications for the triennial UCC update.

**Amendment parsing status:** partial_high_level

**Primary amendment sources:** `src:usa-pa:pa-code-403-21`; `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted`.

### 5.2 State Amendment Sources

| Amendment Source ID | Publication | Coverage | Status | Source IDs |
| --- | --- | --- | --- | --- |
| `amend-source:usa-pa:403-21-current` | 34 Pa. Code § 403.21, Uniform Construction Code | Current incorporated-code list and Pennsylvania modifications/exclusions | partially_parsed | `src:usa-pa:pa-code-403-21` |
| `amend-source:usa-pa:55-pa-b-7701` | 55 Pa.B. 7701 final-omitted rulemaking | 2021 triennial update amendments to Chapters 401 and 403 | partially_parsed | `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted` |
| `amend-source:usa-pa:local-change-registry` | DLI Municipal Code Change Ordinances page | Local equal-or-more-stringent municipal change decisions and reviews | registry_identified | `src:usa-pa:dli-municipal-code-change-ordinances` |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Summary | Practical Effect | Source IDs | Verification Status |
| --- | --- | --- | --- | --- | --- |
| `amend:usa-pa:ibc-chapter-selection-2021` | Building | UCC incorporates selected chapters of the 2021 IBC rather than the entire IBC without change | Use § 403.21 chapter list and exclusions before treating any IBC section as adopted statewide | `src:usa-pa:pa-code-403-21`; `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted` | partially_verified |
| `amend:usa-pa:ifc-reference-limited-2021` | Fire - construction references | 2021 IFC applies only to the extent referenced by other adopted ICC codes | Do not normalize IFC as a standalone operational fire-prevention code from this source alone | `src:usa-pa:pa-code-403-21`; `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted` | partially_verified |
| `amend:usa-pa:ipc-second-class-county-caveat` | Plumbing | IPC adoption is subject to a caveat for plumbing codes adopted by a second-class county under the Local Health Administration Law when the local code meets or exceeds UCC standards | Jurisdiction resolver should flag Allegheny County / second-class-county plumbing review rather than assuming one statewide plumbing AHJ | `src:usa-pa:pa-code-403-21`; `src:usa-pa:pcca-act-45` | partially_verified |
| `amend:usa-pa:energy-compliance-tools` | Energy | § 403.21 allows specified REScheck / Pennsylvania Alternative Residential Energy Provisions and COMcheck methods for IECC compliance | Energy compliance validation should preserve the allowed compliance-path options rather than only listing IECC edition | `src:usa-pa:pa-code-403-21`; `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted` | partially_verified |
| `amend:usa-pa:appendix-limits` | Multiple | Appendices to adopted codes are not adopted except for the IEBC appendices/resource information and specific IBC Appendices E and H identified in § 403.21 | Appendix applicability must not be assumed unless specifically adopted or otherwise required | `src:usa-pa:pa-code-403-21`; `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted` | partially_verified |
| `amend:usa-pa:accessibility-exclusions` | Accessibility | Some accessibility provisions in model-code sources are excluded or redirected through Pennsylvania accessibility authority | Accessibility adoption should not be treated as a simple model-code carryover until Act 16 and DLI regulatory implementation are reconciled | `src:usa-pa:pcca-act-45`; `src:usa-pa:act-16-2026`; `src:usa-pa:pa-code-403-21` | partially_verified |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-pa"
  model: "municipal_election_with_state_and_third_party_fallback"
  enforcing_entities:
    - "municipal code officials in municipalities that administer and enforce the UCC"
    - "certified third-party agencies retained by municipalities or applicants where permitted"
    - "intermunicipal or contracted municipal enforcement arrangements"
    - "Pennsylvania Department of Labor & Industry for applicable nonresidential or other structures under statutory and regulatory pathways"
  required_officials:
    - "certified code administrators"
    - "certified construction code officials"
    - "certified accessibility inspectors or plans examiners where accessibility enforcement is local"
  state_reserved_activities:
    - "statewide UCC regulation and triennial adoption process"
    - "accessibility modification and technical-infeasibility oversight requiring state reconciliation after Act 16 of 2026"
    - "DLI enforcement or plan-review roles in non-administering jurisdictions and other structures where statute or regulation assigns DLI involvement"
  source_ids:
    - "src:usa-pa:pcca-act-45"
    - "src:usa-pa:pa-code-403-102"
    - "src:usa-pa:dli-ucc-home"
  verification_status: "partially_verified"
  confidence: 0.84
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-pa"
  model: "statewide_minimum_with_equal_or_more_stringent_local_changes_after_state_review"
  applies_to_code_families:
    - "building"
    - "residential"
    - "mechanical"
    - "plumbing"
    - "fuel_gas"
    - "energy"
    - "other UCC code families where a local technical change is authorized"
  approval_required: true
  approving_authority_id: "ahj:usa-pa:dli-bois-ucc"
  filing_required: true
  registry_exists: true
  registry_source_ids:
    - "src:usa-pa:dli-municipal-code-change-ordinances"
  legal_basis_source_ids:
    - "src:usa-pa:pcca-act-45"
    - "src:usa-pa:pa-code-403-102"
  verification_status: "partially_verified"
  confidence: 0.82
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement authority and local amendment authority should be resolved separately. A municipality may administer and enforce the UCC by adopting the current UCC and arranging code-administration services. That does not by itself create a local technical amendment. Local changes to the UCC must equal or exceed the statewide UCC, follow the required public process, and be reviewed by DLI. Source IDs: `src:usa-pa:pcca-act-45`; `src:usa-pa:pa-code-403-102`; `src:usa-pa:dli-municipal-code-change-ordinances`.

### 6.4 Known Local Amendment Registries

| Registry ID | Registry Name | Maintainer | Coverage | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| `registry:usa-pa:dli-municipal-code-change-ordinances` | Municipal Code Change Ordinances | Pennsylvania Department of Labor & Industry | DLI-reviewed municipal code-change ordinances and decisions listed on the DLI UCC site | `src:usa-pa:dli-municipal-code-change-ordinances` | identified_not_exhaustively_parsed |

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendment was parsed into normalized local records in this pass. The DLI municipal code-change ordinance registry was identified and should be parsed separately for jurisdiction-specific amendment coverage. Source ID: `src:usa-pa:dli-municipal-code-change-ordinances`.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: partial_design_only

Jurisdiction stack:

```text
Address
  -> State: Pennsylvania
  -> County
  -> Municipality or unincorporated area
  -> Municipality UCC administration status
  -> Second-class-county plumbing caveat check, if applicable
  -> Building-code administrator or certified third-party agency
  -> DLI pathway for applicable non-administering-jurisdiction structures
  -> Accessibility official / DLI accessibility pathway
  -> Fire AHJ, unresolved for operational/prevention code
  -> Applicable state UCC adoption records
  -> Applicable DLI-reviewed local code-change records
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | not selected | none | statewide | unknown | pending |
| County | not selected | none | statewide | unknown | pending |
| Municipality | not selected | none | statewide | unknown | pending |
| Municipal UCC administration status | not selected | none | statewide | unknown | pending |
| Fire District / Fire AHJ | not selected | none | statewide | unknown | unresolved |
| Local amendment registry | DLI Municipal Code Change Ordinances | `src:usa-pa:dli-municipal-code-change-ordinances` | statewide where DLI has posted reviewed local changes | irregular / page-update based | identified |

### 7.3 AHJ Contact Data

No AHJ contact dataset was populated. The next pass should select authoritative municipal administration data, DLI contact lists, and local code-office contact sources before this report is used for address-level AHJ routing.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Source Name | Issuer | Type | URL / Citation | Records Supported | Caveats |
| --- | --- | --- | --- | --- | --- | --- |
| `src:usa-pa:dli-ucc-home` | Uniform Construction Code Home | Pennsylvania Department of Labor & Industry | agency page | https://www.pa.gov/agencies/dli/programs-services/labor-management-relations/bureau-of-occupational-and-industrial-safety/uniform-construction-code-home | statewide UCC program existence and general description | Agency web page; use statutes and regulations for legal hierarchy |
| `src:usa-pa:dli-ucc-rac-2021-cycle` | Uniform Construction Code Review and Advisory Council / 2021 Code Review Cycle | Pennsylvania Department of Labor & Industry | agency page | https://www.pa.gov/agencies/dli/programs-services/labor-management-relations/bureau-of-occupational-and-industrial-safety/uniform-construction-code-home/review-advisory-council.html | 2021 triennial cycle; IRRC approval date; publication date; effective date | Agency status page; legal text is in PA Code and PA Bulletin |
| `src:usa-pa:dli-ucc-regs-statutes` | UCC Regulations and Statutes | Pennsylvania Department of Labor & Industry | agency page | https://www.pa.gov/agencies/dli/programs-services/labor-management-relations/bureau-of-occupational-and-industrial-safety/uniform-construction-code-home/regulations-statutes.html | Construction Code Act and UCC regulatory source pointers | Portal page; use linked official statute/regulation text for legal reliance |
| `src:usa-pa:pcca-act-45` | Pennsylvania Construction Code Act, Act 45 of 1999, as amended | Pennsylvania General Assembly | statute | https://www.legis.state.pa.us/WU01/LI/LI/US/HTM/1999/0/0045..HTM | statewide authority, preemption, triennial regulation process, municipal administration, transition rules, local changes, accessibility roles | Official unconsolidated act page with amendments and compiler notes; reconcile with current codification and Act 16 implementation |
| `src:usa-pa:act-16-2026` | Act 16 of 2026 / SB 867 amendments to the Construction Code Act | Pennsylvania General Assembly | statute amendment | https://www.legis.state.pa.us/cfdocs/billInfo/billInfo.cfm?sYear=2025&sInd=0&body=S&type=B&bn=867 | accessibility-process risk flag and pending reconciliation | Official bill/law information; detailed regulatory implementation not parsed |
| `src:usa-pa:pa-code-403-1` | 34 Pa. Code § 403.1, Scope | Pennsylvania Code / Legislative Reference Bureau | regulation | https://www.pacodeandbulletin.gov/Display/pacode?file=/secure/pacode/data/034/chapter403/s403.1.html | UCC scope, transition/g grandfathering, prior-permit logic, effective-history note | Official HTML drawn from PA Code; verify printed/code text for production legal citations |
| `src:usa-pa:pa-code-403-21` | 34 Pa. Code § 403.21, Uniform Construction Code | Pennsylvania Code / Legislative Reference Bureau | regulation | https://www.pacodeandbulletin.gov/Display/pacode?file=/secure/pacode/data/034/chapter403/s403.21.html | current model-code adoption matrix, amendments, IFC reference limitation, IEBC/IECC paths, appendices, plumbing caveat | Official HTML drawn from PA Code; extensive code text should be checked against official publication in edge cases |
| `src:usa-pa:pa-code-403-43` | 34 Pa. Code § 403.43, Grant, denial and effect of permits | Pennsylvania Code / Legislative Reference Bureau | regulation | https://www.pacodeandbulletin.gov/Display/pacode?file=/secure/pacode/data/034/chapter403/s403.43.html | commercial permit review and permit-validity rules | Official HTML drawn from PA Code |
| `src:usa-pa:pa-code-403-63` | 34 Pa. Code § 403.63, Grant, denial and effect of permits | Pennsylvania Code / Legislative Reference Bureau | regulation | https://www.pacodeandbulletin.gov/Display/pacode?file=/secure/pacode/data/034/chapter403/s403.63.html | residential permit review and permit-validity rules | Official HTML drawn from PA Code |
| `src:usa-pa:pa-code-403-102` | 34 Pa. Code § 403.102, Municipal and third-party agency administration and enforcement | Pennsylvania Code / Legislative Reference Bureau | regulation | https://www.pacodeandbulletin.gov/secure/pacode/data/034/chapter403/034_0403.pdf | local enforcement options, municipal adoption ordinance, DLI notification, local change review, accessibility official roles | Official PA Code chapter PDF; use section number § 403.102 for stable citation and spot-check current codification |
| `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted` | 55 Pa.B. 7701, UCC final-omitted rulemaking, Volume 55, Number 45 | Pennsylvania Bulletin / Legislative Reference Bureau | official rulemaking PDF | https://pacodeandbulletin.gov/secure/pabulletin/data/vol55/55-45/55_45_rr.pdf | 2021 triennial rulemaking, effective date, code adoption text, amendments to Chapters 401 and 403 | Official PDF; extraction is OCR/text-layout sensitive, so use section numbers and PA Code codification for production |
| `src:usa-pa:dli-municipal-code-change-ordinances` | Municipal Code Change Ordinances | Pennsylvania Department of Labor & Industry | agency registry page | https://www.pa.gov/agencies/dli/programs-services/labor-management-relations/bureau-of-occupational-and-industrial-safety/uniform-construction-code-home/municipal-ordinances.html | local amendment registry, DLI decisions on municipal changes | Registry was identified but not exhaustively parsed into local records |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| `src:usa-pa:pa-code-403-1` | html_codification | Official PA Code HTML may display formatting differently than printed/codified law; use section text and history notes for stable references. | acceptable_with_spot_check |
| `src:usa-pa:pa-code-403-21` | html_codification | The adoption section is long and amendment-heavy; individual model-code provisions must be validated against § 403.21 before downstream normalization. | acceptable_with_full_amendment_parse |
| `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted` | official_pdf_ocr_layout | Rulemaking PDF is official but text extraction may split columns and sections; screenshots and codified PA Code were used for spot checks. | acceptable_as_publication_source |
| `src:usa-pa:pcca-act-45` | amended_act_page | The official Act page includes amendments and compiler notes; accessibility provisions changed in 2026 and require regulatory reconciliation. | acceptable_with_act_16_review |
| `src:usa-pa:dli-municipal-code-change-ordinances` | registry_not_exhaustively_parsed | The registry page confirms a DLI publication path for local changes, but individual municipal amendments were not normalized. | needs_followup_parse |
| `src:usa-pa:pa-code-403-102` | local_enforcement_detail_not_fully_normalized | The local enforcement section contains detailed pathways; this report captures the main structure only. | acceptable_for_statewide_summary |

### 8.3 Supplemental Sources

None used in the report body.

### 8.4 Source Extraction Metadata

| Field | Value |
| --- | --- |
| Extraction date | 2026-06-26 |
| Extraction scope | Pennsylvania UCC authority, 2021 code adoption, transition rules, local enforcement, local amendment authority, source registry |
| Extraction method | Official state agency pages, Pennsylvania Code pages, Pennsylvania General Assembly act page, Pennsylvania Bulletin PDF |
| PDF visual spot check | Pennsylvania Bulletin 55 Pa.B. 7701 pages containing effective date and § 403.21 adoption text were visually checked |
| Not completed | Full model-code amendment diff; local municipal amendment record parsing; operational fire-code authority; standalone electrical-code authority; AHJ boundary/contact data |
| Reviewer | none |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| `report` | `report.status` | partially_verified | verified | 1.00 | none | Status upgraded after source registry, authority, adoption, transition, local enforcement, and local amendment fields were populated and checked. |
| `report` | `risk.overall_confidence` | 0.72 | verified | 1.00 | none | Confidence remains below verified due to unresolved operational fire, standalone electrical, AHJ, and full-amendment parsing issues. |
| `ahj:usa-pa:dli-bois-ucc` | primary authority | Pennsylvania Department of Labor & Industry, UCC program | partially_verified | 0.86 | `src:usa-pa:pcca-act-45`; `src:usa-pa:dli-ucc-home`; `src:usa-pa:dli-ucc-regs-statutes` | Statewide UCC program and statutory authority confirmed. |
| `adopt:usa-pa:ucc-2021-triennial` | effective date | 2026-01-01 | verified | 0.91 | `src:usa-pa:dli-ucc-rac-2021-cycle`; `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted` | Agency page and official rulemaking agree. |
| `adopt:usa-pa:ucc-2021-triennial` | base codes | 2021 IBC, IRC, IMC, IFGC, IPC, IECC, IEBC and limited IFC references | partially_verified | 0.86 | `src:usa-pa:pa-code-403-21`; `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted` | High-level code families captured; detailed amendment diff remains open. |
| `adopt:usa-pa:ifc-2021-reference-only` | fire code scope | IFC only to extent referenced by other adopted ICC codes | partially_verified | 0.57 | `src:usa-pa:pa-code-403-21` | Operational fire-code authority unresolved. |
| `adopt:usa-pa:electrical-standalone-unresolved` | standalone electrical adoption | unresolved | unresolved | 0.42 | `src:usa-pa:pa-code-403-21` | No standalone NEC adoption established from § 403.21 or other official sources reviewed. |
| `local-enforcement:usa-pa` | model | municipal_election_with_state_and_third_party_fallback | partially_verified | 0.84 | `src:usa-pa:pcca-act-45`; `src:usa-pa:pa-code-403-102` | Main model confirmed; address-level AHJ routing not populated. |
| `local-amendment-rule:usa-pa` | model | statewide_minimum_with_equal_or_more_stringent_local_changes_after_state_review | partially_verified | 0.82 | `src:usa-pa:pcca-act-45`; `src:usa-pa:pa-code-403-102`; `src:usa-pa:dli-municipal-code-change-ordinances` | Registry identified; individual municipal records not parsed. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Every `src:usa-pa:*` ID cited in the report appears in Section 8. |
| All authority IDs resolve | pass | Authority IDs used in edges and local rules are defined or explicitly marked unresolved. |
| All current code families have adoption matrix rows | pass | Rows are present for all template code families. |
| All current code families have adoption records | fail | Core UCC adoption records are present, but standalone electrical and operational fire rows remain unresolved. |
| Building and operational fire code are separated | pass | Construction fire references and operational fire-prevention code are separated. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Date columns are separate and transition notes are separate. |
| Effective dates are valid ISO dates | pass | Known dates use ISO format; unresolved fields use `null`. |
| No impossible date sequences | pass | No contradictory date sequences were introduced. |
| Transition rules have explicit trigger conditions | pass | Contract, prior-permit, review-clock, and permit-validity triggers are stated. |
| Permit-date logic is captured where applicable | pass | Permit review and permit validity rules are included in Section 4. |
| Local enforcement model classified | pass | Municipal-election/state-third-party fallback model is captured. |
| Local amendment rule classified | pass | Equal-or-more-stringent local change model and DLI registry are captured. |
| AHJ confirmation metadata present | fail | AHJ contact/boundary data remains absent. |
| Official-source caveats captured | pass | Caveats are captured for each high-risk official source type. |
| Leftover template markers removed | pass | No template placeholders or prohibited stub markers remain after validation. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `issue:usa-pa:001` | high | operational fire code | Confirm whether Pennsylvania has a separate statewide operational fire-prevention code authority, edition, or Fire and Panic Act enforcement pathway distinct from UCC construction references. | Extract Fire and Panic Act savings/implementation provisions, state fire marshal or DLI fire-safety sources, and any current operational fire regulations. | none | null | open |
| `issue:usa-pa:002` | high | standalone electrical code | Confirm whether a standalone statewide electrical code or NEC edition is adopted outside the UCC model-code provisions. | Search DLI/UCC, PA Code, and licensing/enforcement sources for NEC-specific adoption and AHJ rules. | none | null | open |
| `issue:usa-pa:003` | high | accessibility Act 16 reconciliation | Act 16 of 2026 amended Construction Code Act accessibility provisions; current UCC accessibility records need regulatory reconciliation. | Parse Act 16 amendments and any DLI implementation guidance or PA Code updates after `2026-05-06`. | none | null | open |
| `issue:usa-pa:004` | medium | full amendment diff | § 403.21 contains many Pennsylvania-specific exclusions and modified sections that were not fully normalized. | Convert § 403.21 and 55 Pa.B. 7701 amendments into section-level amendment records. | none | null | open |
| `issue:usa-pa:005` | medium | local amendment registry parsing | DLI's municipal ordinance registry was identified but municipality-level amendments were not normalized. | Parse DLI registry entries into local amendment records with municipality, ordinance, decision date, and affected code sections. | none | null | open |
| `issue:usa-pa:006` | medium | AHJ routing data | Address-level AHJ resolution requires municipal UCC administration status, DLI fallback pathways, and contact data. | Select authoritative boundary, municipal administration, DLI, and local office contact sources. | none | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| `watch:usa-pa:dli-ucc-home` | `src:usa-pa:dli-ucc-home` | html_diff | monthly | UCC program page announces new code status, guidance, or portal restructuring | 2026-06-26 | active |
| `watch:usa-pa:rac-cycle` | `src:usa-pa:dli-ucc-rac-2021-cycle` | html_diff | monthly | RAC page posts new review-cycle material, recommendations, or adoption status | 2026-06-26 | active |
| `watch:usa-pa:pa-code-403-21` | `src:usa-pa:pa-code-403-21` | html_diff | monthly | § 403.21 code-adoption text changes | 2026-06-26 | active |
| `watch:usa-pa:pa-bulletin-ucc` | `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted` | publication_watch | monthly | New UCC rulemaking or correction is published | 2026-06-26 | active |
| `watch:usa-pa:act-45` | `src:usa-pa:pcca-act-45` | statute_watch | quarterly | Construction Code Act amended or compiler notes change | 2026-06-26 | active |
| `watch:usa-pa:act-16` | `src:usa-pa:act-16-2026` | implementation_watch | monthly | DLI publishes accessibility implementation guidance or regulations | 2026-06-26 | active |
| `watch:usa-pa:local-ordinances` | `src:usa-pa:dli-municipal-code-change-ordinances` | registry_diff | monthly | DLI adds, removes, or revises local code-change ordinance entries | 2026-06-26 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-26 | Replaced baseline draft with source-backed Pennsylvania UCC authority and adoption report | `report:usa-pa`; `ahj:usa-pa:dli-bois-ucc`; `adopt:usa-pa:ucc-2021-triennial`; `local-enforcement:usa-pa`; `local-amendment-rule:usa-pa` | `src:usa-pa:pcca-act-45`; `src:usa-pa:dli-ucc-home`; `src:usa-pa:dli-ucc-rac-2021-cycle`; `src:usa-pa:pa-code-403-21`; `src:usa-pa:pa-bulletin-55-45-ucc-final-omitted`; `src:usa-pa:pa-code-403-102` | ChatGPT | Status set to `partially_verified`; unresolved operational fire, standalone electrical, accessibility reconciliation, amendment-diff, and AHJ issues left explicit. |
