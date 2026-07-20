---
state:
  state_id: "US-NC"
  name: "North Carolina"
  abbreviation: "NC"
report:
  report_id: "state-report:usa-nc"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-26"
  last_verified: null
  reviewed_by: null
risk:
  overall_confidence: 0.62 # 0.00 - 1.00
  risk_flags:
    - "2024_code_effective_date_is_contingent"
    - "adoption_dates_not_fully_normalized_from_rule_history"
    - "local_fire_prevention_code_approval_registry_not_extracted"
    - "electrical_code_scope_split_requires_second_pass"
  open_questions_count: 6

---

# State Building Code Authority Report: North Carolina

## 1. Executive Summary

- **Authority model:** North Carolina has a statewide code model. The North Carolina State Building Code is prepared and adopted by the Building Code Council and Residential Code Council under Chapter 143, Article 9. The Building Code Council oversees commercial and multi-family code volumes, while the Residential Code Council oversees residential construction code volumes. OSFM also states that the State Building Code is adopted and amended for statewide implementation by the Building Code Council. Source IDs: `src:usa-nc:gs-143-136-138`; `src:usa-nc:osfm-codes`.

- **Statewide code status:** The currently mandatory statewide code baseline remains the 2018 North Carolina State Building Code suite for the major ICC-derived volumes, with a current electrical-code split noted below. The OSFM code portal states that the 2018 NC Codes were effective January 1, 2019, are currently effective, and were available during a July 1, 2018 through December 31, 2018 transition period. A December 5, 2025 OSFM formal interpretation confirms that the 2018 State Building Code remains effective and that the 2024 State Building Code is not yet mandatorily effective. Source IDs: `src:usa-nc:osfm-codes-current-past`; `src:usa-nc:osfm-2025-formal-interpretation-107`.

- **Local enforcement model:** Local governments perform building-code administration through inspection departments, joint departments, contracts, or county-provided inspection services. The State Fire Marshal has general supervisory, administrative, and enforcement authority for most State Building Code construction-code sections, cooperates with local officials, and may arrange inspection services if a local government fails to provide them. Source IDs: `src:usa-nc:gs-143-139`; `src:usa-nc:gs-160d-1102`; `src:usa-nc:gs-160d-art11`.

- **Local amendment posture:** The State Building Code applies throughout North Carolina from adoption, except as otherwise provided. Political subdivisions may adopt fire-prevention codes and floodplain management regulations, but local fire-prevention codes and most local variants require responsible Code Council approval before becoming effective. The scope of any statewide registry of approved local rules was not extracted in this pass. Source ID: `src:usa-nc:gs-143-136-138`.

- **Known transition periods or pending changes:** The 2024 North Carolina State Building Code has been adopted as a future code suite but is delayed. S.L. 2025-2 delays the 2024 Code until 12 months after the first day of the month following State Fire Marshal certification that publication/distribution requirements are complete and the Residential Code Council is fully constituted. OSFM states that, until the official effective date is determined, the 2018 Code remains in effect and the 2024 Code may be used as an alternative method at the owner or agent's request. Source IDs: `src:usa-nc:sl-2025-2`; `src:usa-nc:osfm-2025-delay-press`; `src:usa-nc:osfm-2025-formal-interpretation-107`.

- **Production readiness:** partially_ready_for_core_authority_and_current_code; not_ready_for_full_local_amendment_registry_or_complete_rule-history normalization.

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: The Building Code Council and Residential Code Council may prepare and
    adopt the North Carolina State Building Code; the Building Code Council oversees
    commercial and multi-family code volumes, and the Residential Code Council oversees
    residential code volumes.
  confidence: 0.85
  source_ids:
  - src:usa-nc:gs-143-136-138
  - src:usa-nc:osfm-codes
- topic: Primary building code edition
  finding: The current mandatory statewide building-code baseline is the 2018 NC Code
    suite, effective 2019-01-01, with the 2018 NC Building Code based on the 2015
    IBC with NC amendments.
  confidence: 0.8
  source_ids:
  - src:usa-nc:osfm-codes-current-past
  - src:usa-nc:osfm-2025-formal-interpretation-107
- topic: Electrical code authority / edition
  finding: 'North Carolina Electrical Code is an adopted NEC version with NC amendments.
    OSFM''s code portal lists the 2020 Electrical Code as effective 2021-11-01 and
    currently effective; OSFM''s 2025 formal interpretation identifies a residential/non-residential
    scope split: 2017 NFPA 70 for 2018 Residential Code scope and 2020 NFPA 70 for
    other structures.'
  confidence: 0.65
  source_ids:
  - src:usa-nc:osfm-codes-current-past
  - src:usa-nc:osfm-electrical-division
  - src:usa-nc:osfm-2025-formal-interpretation-107
- topic: Fire code authority
  finding: The North Carolina Fire Code is a State Building Code volume. Local political
    subdivisions may adopt local fire-prevention codes, subject to statutory limitations
    and responsible Code Council approval unless an exception applies.
  confidence: 0.75
  source_ids:
  - src:usa-nc:gs-143-136-138
  - src:usa-nc:osfm-codes-current-past
- topic: Local enforcement
  finding: Local governments must perform inspection duties through their own inspection
    department, a joint department, a contract with another unit, or county-provided
    services; local inspectors inspect work in progress under permits.
  confidence: 0.8
  source_ids:
  - src:usa-nc:gs-160d-1102
  - src:usa-nc:gs-160d-art11
- topic: Local amendments
  finding: Statewide code applies throughout the State. Local fire-prevention codes
    and most local regulations require responsible Code Council approval before becoming
    effective; local floodplain regulations have a distinct allowance.
  confidence: 0.7
  source_ids:
  - src:usa-nc:gs-143-136-138
- topic: Effective / operative date rule
  finding: 'Current 2018 code suite: effective 2019-01-01, with transition from 2018-07-01
    through 2018-12-31. Future 2024 suite: mandatory date contingent under S.L. 2025-2.'
  confidence: 0.75
  source_ids:
  - src:usa-nc:osfm-codes-current-past
  - src:usa-nc:sl-2025-2
  - src:usa-nc:osfm-2025-formal-interpretation-107
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | `ahj:usa-nc:building-code-council` |
| Authority name | North Carolina Building Code Council |
| Authority type | statewide code council |
| Legal basis | N.C. Gen. Stat. § 143-136; N.C. Gen. Stat. § 143-138 |
| Role | Reviews, considers, and oversees revisions or amendments to State Building Code volumes and other codes applicable to commercial or multi-family construction; shares State Building Code preparation/adoption authority with the Residential Code Council. |
| Enforcement model | statewide adoption with State Fire Marshal supervision/administration/enforcement and local inspection departments for local administration. |
| Source IDs | `src:usa-nc:gs-143-136-138`; `src:usa-nc:gs-143-139`; `src:usa-nc:osfm-codes` |
| Verification status | verified_core |

| Field | Value |
| --- | --- |
| Authority ID | `ahj:usa-nc:residential-code-council` |
| Authority name | North Carolina Residential Code Council |
| Authority type | statewide residential code council |
| Legal basis | N.C. Gen. Stat. § 143-136.1; N.C. Gen. Stat. § 143-138 |
| Role | Oversees code volumes applicable to residential construction and contained in the State Building Code volumes; performs residential-code revision duties. |
| Enforcement model | statewide adoption with local inspection departments and State Fire Marshal oversight; formation/confirmation status remains relevant to 2024 Code mandatory effectiveness. |
| Source IDs | `src:usa-nc:gs-143-136-138`; `src:usa-nc:osfm-codes`; `src:usa-nc:sl-2025-2`; `src:usa-nc:osfm-2025-formal-interpretation-107` |
| Verification status | verified_core; current membership trigger unresolved for production monitoring |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | `ahj:usa-nc:building-code-council` | North Carolina Building Code Council | Adopts and amends commercial / multi-family building-code volumes; hears appeals/interpretations within its scope. | N.C. Gen. Stat. §§ 143-136, 143-138 | `src:usa-nc:gs-143-136-138` | verified_core |
| Residential | `ahj:usa-nc:residential-code-council` | North Carolina Residential Code Council | Oversees residential construction code volumes and residential-code revisions. | N.C. Gen. Stat. §§ 143-136.1, 143-138 | `src:usa-nc:gs-143-136-138`; `src:usa-nc:osfm-codes` | verified_core |
| Existing Building / Rehabilitation | `ahj:usa-nc:building-code-council`; `ahj:usa-nc:residential-code-council` | Building Code Council / Residential Code Council | Existing Building Code is a State Building Code volume; scope may depend on commercial/multi-family versus residential application. | N.C. Gen. Stat. § 143-138(a) | `src:usa-nc:gs-143-136-138`; `src:usa-nc:osfm-2025-formal-interpretation-107` | partially_verified |
| Mechanical | `ahj:usa-nc:building-code-council`; `ahj:usa-nc:residential-code-council` | Building Code Council / Residential Code Council | Mechanical Code is a State Building Code volume; RCC has special review authority for residential-applicable mechanical provisions. | N.C. Gen. Stat. § 143-138(a), (d) | `src:usa-nc:gs-143-136-138`; `src:usa-nc:osfm-2025-formal-interpretation-107` | partially_verified |
| Plumbing | `ahj:usa-nc:building-code-council`; `ahj:usa-nc:residential-code-council` | Building Code Council / Residential Code Council | Plumbing Code is a State Building Code volume; State Fire Marshal supervises and administers plumbing-code enforcement except where allocated elsewhere. | N.C. Gen. Stat. §§ 143-138(a), 143-139(b) | `src:usa-nc:gs-143-136-138`; `src:usa-nc:gs-143-139` | verified_core |
| Fuel Gas | `ahj:usa-nc:building-code-council`; `ahj:usa-nc:residential-code-council` | Building Code Council / Residential Code Council | Fuel Gas Code is a State Building Code volume; RCC has special review authority for residential-applicable fuel-gas provisions. | N.C. Gen. Stat. § 143-138(a), (d) | `src:usa-nc:gs-143-136-138`; `src:usa-nc:osfm-2025-formal-interpretation-107` | partially_verified |
| Electrical | `ahj:usa-nc:building-code-council`; `ahj:usa-nc:residential-code-council`; `ahj:usa-nc:state-electrical-division` | Building Code Council / Residential Code Council / OSFM State Electrical Division | Electrical Code is a State Building Code volume; State Electrical Division interprets the State Electrical Code and provides access to adopted amendments. | N.C. Gen. Stat. §§ 143-138(a), 143-139(b); OSFM State Electrical Division page | `src:usa-nc:gs-143-136-138`; `src:usa-nc:gs-143-139`; `src:usa-nc:osfm-electrical-division` | partially_verified |
| Energy | `ahj:usa-nc:building-code-council`; `ahj:usa-nc:residential-code-council` | Building Code Council / Residential Code Council | Energy Conservation Code is a State Building Code volume; RCC has special review authority for residential-applicable energy provisions. | N.C. Gen. Stat. § 143-138(a), (d) | `src:usa-nc:gs-143-136-138`; `src:usa-nc:osfm-2025-formal-interpretation-107` | partially_verified |
| Fire - construction references | `ahj:usa-nc:building-code-council`; `ahj:usa-nc:state-fire-marshal` | Building Code Council / State Fire Marshal | Fire Code is a State Building Code volume; State Fire Marshal has general authority for fire-protection portions of the Code except allocated sections. | N.C. Gen. Stat. §§ 143-138(a), 143-139(b) | `src:usa-nc:gs-143-136-138`; `src:usa-nc:gs-143-139` | verified_core |
| Fire - operational / prevention code | `ahj:usa-nc:building-code-council`; local political subdivisions subject to approval | Building Code Council / local fire-prevention code authority | State Fire Code is a code volume; political subdivisions may adopt fire-prevention codes within jurisdiction, subject to limits and responsible Code Council approval. | N.C. Gen. Stat. § 143-138(e) | `src:usa-nc:gs-143-136-138` | partially_verified |
| Accessibility | `ahj:usa-nc:building-code-council` | North Carolina Building Code Council | Accessibility requirements are tied to the 2018 NC Building Code and ANSI A117.1 standard listed by OSFM. | OSFM current-codes page | `src:usa-nc:osfm-codes-current-past` | partially_verified |
| Elevator / Conveyance | `ahj:usa-nc:department-of-labor-elevator-bureau` | NC Department of Labor, Elevator and Amusement Device Bureau | Department of Labor has general supervision over State Building Code sections pertaining to elevators, moving stairways, and amusement devices; the Bureau enforces the Elevator Safety Act. | N.C. Gen. Stat. § 143-139(d); NC DOL Elevator page | `src:usa-nc:gs-143-139`; `src:usa-nc:ncdol-elevator` | verified_core |

### 2.3 Authority Hierarchy Notes

North Carolina's model is hybrid but statewide. The State Building Code is a statewide code, with the Building Code Council and Residential Code Council responsible for adoption and revisions within their statutory domains. The State Fire Marshal has general supervision, administration, and enforcement authority for most construction-code sections. Local governments handle day-to-day code administration through inspection departments or equivalent service arrangements. Boilers and elevators are allocated away from OSFM's general construction-code enforcement authority to the Department of Labor in N.C. Gen. Stat. § 143-139(c)-(d); elevator details were verified in this pass, but boiler details were not separately normalized. Source IDs: `src:usa-nc:gs-143-136-138`; `src:usa-nc:gs-143-139`; `src:usa-nc:gs-160d-1102`; `src:usa-nc:ncdol-elevator`.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| `edge:usa-nc:001` | `ahj:usa-nc:building-code-council` | prepares_and_adopts | North Carolina State Building Code volumes for commercial and multi-family construction | `src:usa-nc:gs-143-136-138` | verified_core |
| `edge:usa-nc:002` | `ahj:usa-nc:residential-code-council` | prepares_and_adopts | North Carolina State Building Code volumes applicable to residential construction | `src:usa-nc:gs-143-136-138` | verified_core |
| `edge:usa-nc:003` | `ahj:usa-nc:state-fire-marshal` | supervises_administers_enforces | State Building Code sections for plumbing, electrical, general building restrictions, HVAC, fire protection, and construction generally, except allocated sections | `src:usa-nc:gs-143-139` | verified_core |
| `edge:usa-nc:004` | `ahj:usa-nc:local-government-inspection-department` | enforces_locally | Permitted work under applicable State and local laws | `src:usa-nc:gs-160d-1102`; `src:usa-nc:gs-160d-art11` | verified_core |
| `edge:usa-nc:005` | `ahj:usa-nc:state-fire-marshal` | backstops | Local governments that fail to provide inspection services | `src:usa-nc:gs-160d-1102` | verified_core |
| `edge:usa-nc:006` | `ahj:usa-nc:department-of-labor-elevator-bureau` | supervises_enforces | Elevators, moving stairways, amusement devices, and related special equipment | `src:usa-nc:gs-143-139`; `src:usa-nc:ncdol-elevator` | verified_core |
| `edge:usa-nc:007` | local political subdivision | may_adopt_with_limits | Fire-prevention code and floodplain management regulations within jurisdiction | `src:usa-nc:gs-143-136-138` | partially_verified |
| `edge:usa-nc:008` | responsible Code Council | approves_before_effective | Local fire-prevention code or other local code/regulation, except floodplain management regulations and G.S. 160D-1128 items | `src:usa-nc:gs-143-136-138` | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Administrative | 2018 North Carolina Administrative Code and Policies | state administrative provisions | 2018 | current_mandatory | unresolved | 2019-01-01 | 2019-01-01 | 2019-01-01 | 2018 suite available during transition 2018-07-01 through 2018-12-31. | `src:usa-nc:osfm-codes-current-past`; `src:usa-nc:osfm-2025-formal-interpretation-107` |
| Building | 2018 North Carolina Building Code | 2015 International Building Code with NC amendments | 2018 | current_mandatory | unresolved | 2019-01-01 | 2019-01-01 | 2019-01-01 | 2018 suite available during transition 2018-07-01 through 2018-12-31. | `src:usa-nc:osfm-codes-current-past`; `src:usa-nc:osfm-2025-formal-interpretation-107` |
| Residential | 2018 North Carolina Residential Code | 2015 International Residential Code with NC amendments | 2018 | current_mandatory | unresolved | 2019-01-01 | 2019-01-01 | 2019-01-01 | 2018 Residential Code available during transition 2018-07-01 through 2018-12-31. | `src:usa-nc:osfm-codes-current-past`; `src:usa-nc:osfm-2025-formal-interpretation-107` |
| Existing Building / Rehabilitation | 2018 North Carolina Existing Building Code | 2015 International Existing Building Code with NC amendments | 2018 | current_mandatory | unresolved | 2019-01-01 | 2019-01-01 | 2019-01-01 | 2018 suite available during transition 2018-07-01 through 2018-12-31. | `src:usa-nc:osfm-codes-current-past`; `src:usa-nc:osfm-2025-formal-interpretation-107` |
| Mechanical | 2018 North Carolina Mechanical Code | 2015 International Mechanical Code with NC amendments | 2018 | current_mandatory | unresolved | 2019-01-01 | 2019-01-01 | 2019-01-01 | 2018 suite available during transition 2018-07-01 through 2018-12-31. | `src:usa-nc:osfm-codes-current-past`; `src:usa-nc:osfm-2025-formal-interpretation-107` |
| Plumbing | 2018 North Carolina Plumbing Code | 2015 International Plumbing Code with NC amendments | 2018 | current_mandatory | unresolved | 2019-01-01 | 2019-01-01 | 2019-01-01 | 2018 suite available during transition 2018-07-01 through 2018-12-31. | `src:usa-nc:osfm-codes-current-past`; `src:usa-nc:osfm-2025-formal-interpretation-107` |
| Fuel Gas | 2018 North Carolina Fuel Gas Code | 2015 International Fuel Gas Code with NC amendments | 2018 | current_mandatory | unresolved | 2019-01-01 | 2019-01-01 | 2019-01-01 | 2018 suite available during transition 2018-07-01 through 2018-12-31. | `src:usa-nc:osfm-codes-current-past`; `src:usa-nc:osfm-2025-formal-interpretation-107` |
| Electrical | North Carolina Electrical Code; scope split noted | 2017 NFPA 70 for structures in 2018 Residential Code scope; 2020 NFPA 70 for other structures, both with NC amendments | 2017 / 2020 | current_mandatory_with_scope_note | unresolved | 2019-01-01 for residential-scope 2017 provisions; 2021-11-01 for 2020 Electrical Code portal listing | 2019-01-01 / 2021-11-01 | 2019-01-01 / 2021-11-01 | OSFM current page lists 2020 Electrical Code as currently effective; OSFM formal interpretation identifies a scope split. Reconcile in second pass before verified status. | `src:usa-nc:osfm-codes-current-past`; `src:usa-nc:osfm-electrical-division`; `src:usa-nc:osfm-2025-formal-interpretation-107` |
| Energy | 2018 North Carolina Energy Conservation Code | 2015 International Energy Conservation Code with NC amendments | 2018 | current_mandatory | unresolved | 2019-01-01 | 2019-01-01 | 2019-01-01 | 2018 suite available during transition 2018-07-01 through 2018-12-31. | `src:usa-nc:osfm-codes-current-past`; `src:usa-nc:osfm-2025-formal-interpretation-107` |
| Fire - construction references | 2018 North Carolina Fire Code | 2015 International Fire Code with NC amendments | 2018 | current_mandatory | unresolved | 2019-01-01 | 2019-01-01 | 2019-01-01 | Same statewide Fire Code volume; construction-trigger details not parsed beyond code-volume status. | `src:usa-nc:osfm-codes-current-past`; `src:usa-nc:osfm-2025-formal-interpretation-107` |
| Fire - operational / prevention code | 2018 North Carolina Fire Code; approved local fire-prevention codes where applicable | 2015 International Fire Code with NC amendments | 2018 | current_mandatory_with_local_fire_prevention_allowance | unresolved | 2019-01-01 | 2019-01-01 | 2019-01-01 | Local fire-prevention codes require statutory review/approval rules; registry not extracted. | `src:usa-nc:osfm-codes-current-past`; `src:usa-nc:gs-143-136-138` |
| Accessibility | Accessibility provisions in 2018 North Carolina Building Code | 2009 ANSI A117.1 referenced by OSFM for 2018 NCBC | 2018 NCBC / 2009 ANSI A117.1 | current_mandatory | unresolved | 2019-01-01 | 2019-01-01 | 2019-01-01 | Accessibility is tracked via NC Building Code entry; deeper ADA/FHA relationship not parsed. | `src:usa-nc:osfm-codes-current-past` |
| Elevator / Conveyance | Elevator Safety Act / State Building Code elevator provisions | ASME and other referenced standards not normalized | unresolved | current_specialized_program | unresolved | unresolved | unresolved | unresolved | Department of Labor supervision verified; specific adopted edition and state amendments not parsed. | `src:usa-nc:gs-143-139`; `src:usa-nc:ncdol-elevator` |

### 3.2 Adoption Records

#### adoption:usa-nc:2018-state-building-code-suite

| Field | Value |
| --- | --- |
| Applies to | Administrative, Building, Residential, Energy Conservation, Existing Building, Fire, Fuel Gas, Mechanical, Plumbing, and related statewide code administration |
| State code names | 2018 North Carolina Administrative Code and Policies; 2018 North Carolina Building Code; 2018 North Carolina Residential Code; 2018 North Carolina Energy Conservation Code; 2018 North Carolina Existing Building Code; 2018 North Carolina Fire Code; 2018 North Carolina Fuel Gas Code; 2018 North Carolina Mechanical Code; 2018 North Carolina Plumbing Code |
| Base model codes | 2015 International Codes with North Carolina amendments; specific model families listed in the matrix |
| Adoption date | unresolved |
| Effective date | 2019-01-01 |
| Operative date | 2019-01-01 |
| Mandatory date | 2019-01-01 |
| Transition start | 2018-07-01 |
| Transition end | 2018-12-31 |
| Status | current_mandatory |
| Source IDs | `src:usa-nc:osfm-codes-current-past`; `src:usa-nc:osfm-2025-formal-interpretation-107` |
| Notes | OSFM's code portal states the 2018 codes were effective January 1, 2019 and currently effective, with a July 1, 2018 through December 31, 2018 transition period. OSFM's December 2025 formal interpretation confirms the 2018 code remains effective statewide. |

#### adoption:usa-nc:electrical-2020-and-residential-scope-note

| Field | Value |
| --- | --- |
| Applies to | Electrical code provisions |
| State code name | North Carolina Electrical Code |
| Base model code | NEC / NFPA 70 with North Carolina amendments |
| Adoption date | unresolved |
| Effective date | 2021-11-01 for 2020 Electrical Code portal listing; 2019-01-01 for 2017 NFPA 70 provisions within the 2018 Residential Code scope as stated in OSFM's 2025 formal interpretation |
| Operative date | 2021-11-01 / 2019-01-01, scope dependent |
| Mandatory date | 2021-11-01 / 2019-01-01, scope dependent |
| Status | current_mandatory_with_scope_note |
| Source IDs | `src:usa-nc:osfm-codes-current-past`; `src:usa-nc:osfm-electrical-division`; `src:usa-nc:osfm-2025-formal-interpretation-107` |
| Notes | The OSFM current-codes page lists the 2020 Electrical Code as effective 2021-11-01 and currently effective. The OSFM formal interpretation identifies the 2017 North Carolina Electrical Code for buildings/structures included in the 2018 Residential Code scope and the 2020 North Carolina Electrical Code for buildings/structures outside that scope. This split should be reconciled against the adopted electrical amendments before verified status. |

#### adoption:usa-nc:2024-state-building-code-suite-future

| Field | Value |
| --- | --- |
| Applies to | Future 2024 State Building Code suite |
| State code names | 2024 North Carolina Administrative Code and Policies; 2024 NC Building, Residential, Existing Building, Energy Conservation, Fire, Fuel Gas, Mechanical, Plumbing; 2023 North Carolina Electrical Code |
| Base model codes | 2021 International Codes with North Carolina amendments; 2023 NFPA 70 for Electrical; Residential energy/mechanical/fuel-gas exceptions noted in OSFM interpretation |
| Adoption date | various 2023-2024 approvals, not normalized |
| Effective date | contingent; not yet determined |
| Operative date | not mandatorily operative statewide; available as alternative method by owner/agent request |
| Mandatory date | delayed under S.L. 2025-2 until 12 months after the first day of the month following State Fire Marshal certification that statutory publication/distribution and RCC constitution events have occurred |
| Status | future_adopted_not_mandatory |
| Source IDs | `src:usa-nc:sl-2025-2`; `src:usa-nc:osfm-2025-delay-press`; `src:usa-nc:osfm-2025-formal-interpretation-107`; `src:usa-nc:osfm-2024-rrc-links` |
| Notes | The 2024 Code should not be treated as the mandatory statewide code until the S.L. 2025-2 trigger is satisfied and an effective date can be calculated. |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

The current mandatory code baseline is the 2018 suite, effective 2019-01-01. OSFM identifies a pre-effective transition period from 2018-07-01 through 2018-12-31. The future 2024 suite is adopted but not mandatorily effective; S.L. 2025-2 makes its effective date contingent on State Fire Marshal certification after two events: completion of publication/distribution requirements and full constitution of the Residential Code Council. The December 2025 OSFM formal interpretation states that those required events were incomplete as of that letter and that the 2018 Code remains effective while the 2024 Code may be used as an alternative method. Source IDs: `src:usa-nc:osfm-codes-current-past`; `src:usa-nc:sl-2025-2`; `src:usa-nc:osfm-2025-formal-interpretation-107`.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `date-rule:usa-nc:2018-transition` | 2018 NC Codes | transition/grace period | 2018-07-01 through 2018-12-31 | 2018 codes available before 2019-01-01 effective date | Yes, prior code allowed during transition, based on OSFM wording for 2018 transition period. | `src:usa-nc:osfm-codes-current-past` | partially_verified |
| `date-rule:usa-nc:2018-effective` | 2018 NC Codes | effective/mandatory | 2019-01-01 | End of 2018 transition period and OSFM effective-date listing | No current evidence of general prior-code option after 2018-12-31. | `src:usa-nc:osfm-codes-current-past`; `src:usa-nc:osfm-2025-formal-interpretation-107` | verified_core |
| `date-rule:usa-nc:electrical-2020-effective` | 2020 Electrical Code portal listing | effective | 2021-11-01 | OSFM lists the 2020 Electrical Code effective from 2021-11-01 and currently effective. | unresolved for residential-scope split | `src:usa-nc:osfm-codes-current-past`; `src:usa-nc:osfm-2025-formal-interpretation-107` | partially_verified |
| `date-rule:usa-nc:2024-contingent-effective` | 2024 NC State Building Code | contingent future effective date | 12 months after the first day of the month following State Fire Marshal certification | Certification that publication/distribution requirements are complete and the Residential Code Council is fully constituted | Yes, 2018 remains effective until mandatory 2024 effective date; 2024 may be used as an alternative method at owner/agent request. | `src:usa-nc:sl-2025-2`; `src:usa-nc:osfm-2025-delay-press`; `src:usa-nc:osfm-2025-formal-interpretation-107` | verified_core |
| `date-rule:usa-nc:permit-expiration` | Building permits | permit lifecycle | 6 months after issuance if work not commenced; 12 months after discontinuance following commencement | Building permit issued under Article 11 | Not a code-edition transition rule; included for AHJ lifecycle context. | `src:usa-nc:gs-160d-art11` | partially_verified |
| `date-rule:usa-nc:interpretation-choice` | Interpretations after permit issuance | permit interpretation stability | no fixed date | If an interpretation changes after a building permit is issued, permit applicant may choose which interpretation applies unless it would cause harm to life or property. | Yes, for interpretation changes, not code-edition changes. | `src:usa-nc:gs-143-136-138` | partially_verified |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| State Building Code suite | 2024 North Carolina State Building Code suite; 2023 North Carolina Electrical Code | 2025-04-07 OSFM delay press release; 2025-03-19 S.L. 2025-2 enactment | various 2023-2024 approvals, not normalized | contingent; not determined | alternative-method use allowed before mandatory date | contingent under S.L. 2025-2 | active_watch | `src:usa-nc:sl-2025-2`; `src:usa-nc:osfm-2025-delay-press`; `src:usa-nc:osfm-2025-formal-interpretation-107`; `src:usa-nc:osfm-2024-rrc-links` | Calculate only after State Fire Marshal certification letter is identified. |
| Residential code cycle | Future residential six-year revision | statutory | adoption date for first six-year revision unresolved | 2031-01-01 for first six-year revision under statute | 2031-01-01 if statutory sequence remains unchanged | 2031-01-01 | monitor | `src:usa-nc:gs-143-136-138` | Statute says first six-year revision by the Residential Council is adopted to become effective 2031-01-01 and every six years thereafter. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| `applicability-rule:usa-nc:2024-alternative-method` | 2024 NC State Building Code | Projects whose owner or agent requests use of 2024 Code before mandatory effective date | Owner/agent request under 2018 NC Administrative Code and Policies Section 102.5, as summarized by OSFM | 2024 Code may be used as an alternative method before mandatory effectiveness; if used, applicable 2024 requirements apply to that project. | `src:usa-nc:osfm-2025-delay-press`; `src:usa-nc:osfm-2025-formal-interpretation-107` | verified_core |
| `applicability-rule:usa-nc:local-fire-prevention` | Fire prevention | Local fire-prevention codes | Local political subdivision adoption plus responsible Code Council approval, subject to statutory limits | Political subdivisions may adopt fire-prevention codes; local provisions applying to dwellings subject to the Residential Code are limited; approval is required before effectiveness unless exception applies. | `src:usa-nc:gs-143-136-138` | partially_verified |
| `applicability-rule:usa-nc:local-floodplain` | Floodplain management | Buildings/structures in flood hazard areas | Local floodplain regulation | Political subdivisions may adopt floodplain management regulations within jurisdiction; local floodplain regulations may regulate all types and uses of buildings or structures in identified flood hazard areas. | `src:usa-nc:gs-143-136-138` | partially_verified |
| `applicability-rule:usa-nc:low-cost-work-permit-exclusion` | Permit requirements | Certain construction, installation, repair, replacement, or alteration costing $40,000 or less | Work cost threshold and excluded activity list | No permit required for qualifying work performed under the current State Building Code unless it involves listed activities, including load-bearing work, plumbing design changes, HVAC/electrical design changes, materials not permitted by Code, roofing additions, or changes to which the Fire Code applies. | `src:usa-nc:gs-160d-art11` | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** statewide North Carolina amendments to model codes, adopted by the responsible Code Council and approved through rulemaking channels. The OSFM current-codes page links approved amendments for each current code volume. The detailed text of each amendment set was not parsed in this pass.

**Where amendments are published:** OSFM current-and-past codes page and linked ICC / amendment documents; 2024 RRC link index page for future 2024 adopted rules. Source IDs: `src:usa-nc:osfm-codes-current-past`; `src:usa-nc:osfm-2024-rrc-links`.

**Amendment parsing status:** source_registry_created; high-impact amendment extraction not_started.

### 5.2 State Amendment Sources

| Amendment Source ID | Code Volume(s) | Publication Path | Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- |
| `amendsource:usa-nc:2018-current-codes-page` | 2018 Administrative, Building, Residential, Energy, Existing, Fire, Fuel Gas, Mechanical, Plumbing; 2020 Electrical | OSFM Codes - Current and Past page links code references and approved amendments. | located_not_parsed | `src:usa-nc:osfm-codes-current-past` | Use linked amendment PDFs/ICC pages for section-level extraction. |
| `amendsource:usa-nc:2024-rrc-link-index` | 2024 State Building Code suite and 2023 Electrical | OSFM link index for Rules Review Commission links to the 2024 State Building Code. | located_not_parsed | `src:usa-nc:osfm-2024-rrc-links` | Link index page located; underlying approved rules need extraction. |
| `amendsource:usa-nc:2025-formal-interpretation-107` | 2018 and 2024 Administrative Code and Policies; sheathing inspection-related applicability | OSFM formal interpretation dated 2025-12-05. | partially_extracted | `src:usa-nc:osfm-2025-formal-interpretation-107` | Used for effective status and current/future volume lists; not a substitute for full amendment set. |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Section / Topic | Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| `amend:usa-nc:2024-alternative-use-status` | 2024 Code suite | Interim use | 2024 Code may be used before mandatory effectiveness as an alternative method when requested by the owner or agent; the 2018 Code remains mandatory until the 2024 trigger is satisfied. | `src:usa-nc:osfm-2025-delay-press`; `src:usa-nc:osfm-2025-formal-interpretation-107` | verified_core |
| `amend:usa-nc:2024-sheathing-inspection-applicability` | Administrative / inspections | Routine sheathing inspection | OSFM's 2025 formal interpretation states that 2024 sheathing inspection requirements are applicable only when the 2024 Code is used before its mandatory date, and are not generally mandatory until the 2024 Code becomes mandatorily effective. | `src:usa-nc:osfm-2025-formal-interpretation-107` | partially_verified |
| `amend:usa-nc:2018-electrical-scope-split` | Electrical | 2017 / 2020 NEC scope | OSFM's formal interpretation identifies 2017 NFPA 70 with NC amendments for buildings and structures included within the 2018 Residential Code scope and 2020 NFPA 70 with NC amendments for other structures. | `src:usa-nc:osfm-2025-formal-interpretation-107` | needs_second_pass |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-nc"
  model: "statewide_code_with_local_inspection_departments_and_state_backstop"
  enforcing_entities:
    - "local government inspection departments"
    - "joint inspection departments"
    - "contracted inspection services from another unit of local government"
    - "county inspection services for cities when arranged"
    - "State Fire Marshal / OSFM when a local government fails to provide inspection services"
  required_officials:
    - "inspectors appointed by local government with titles such as building, electrical, plumbing, housing, zoning, HVAC, fire prevention, deputy, or assistant inspector"
    - "person designated by each local government for daily oversight of G.S. 160D-1104 duties"
  state_reserved_activities:
    - "State Fire Marshal general supervision, administration, and enforcement for most State Building Code construction sections"
    - "Department of Labor supervision for elevator and moving-stairway sections"
    - "boiler allocation to Department of Labor identified in statute but not normalized in this pass"
  source_ids:
    - "src:usa-nc:gs-143-139"
    - "src:usa-nc:gs-160d-1102"
    - "src:usa-nc:gs-160d-art11"
    - "src:usa-nc:ncdol-elevator"
  verification_status: "partially_verified"
  confidence: 0.78
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-nc"
  model: "statewide_code_with_limited_local_fire_prevention_and_floodplain_authority"
  applies_to_code_families:
    - "fire prevention"
    - "floodplain management"
    - "other local code/regulation variants only where approved or otherwise permitted by statute"
  approval_required: true
  approving_authority_id: "ahj:usa-nc:building-code-council_or_ahj:usa-nc:residential-code-council_as_responsible_code_council"
  filing_required: "unresolved"
  registry_exists: "unresolved"
  registry_source_ids: []
  legal_basis_source_ids:
    - "src:usa-nc:gs-143-136-138"
  verification_status: "partially_verified"
  confidence: 0.70
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement and local amendment authority are separate in North Carolina. Local governments are required to provide code administration/inspection services by one of several mechanisms, but that does not create general authority to adopt local construction codes. The State Building Code applies statewide. Local political subdivisions may adopt local fire-prevention codes and floodplain management regulations, but local fire-prevention codes and most local variants require responsible Code Council approval to become effective. Source IDs: `src:usa-nc:gs-143-136-138`; `src:usa-nc:gs-160d-1102`; `src:usa-nc:gs-160d-art11`.

### 6.4 Known Local Amendment Registries

No comprehensive state registry of approved local fire-prevention codes, floodplain regulations, or other approved local variants was extracted in this pass. The statutory approval pathway is verified; registry availability remains unresolved. Source ID: `src:usa-nc:gs-143-136-138`.

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendments were parsed in this pass. Future work should identify whether OSFM/BCC/RCC meeting records or North Carolina Register notices function as the practical registry of local approvals.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: partially_started

Jurisdiction stack:

```text
Address
  -> State of North Carolina
  -> County
  -> Municipality or unincorporated county jurisdiction
  -> Local government inspection department / contracted inspection provider / county service arrangement
  -> Fire-prevention AHJ where local fire-prevention code applies
  -> State Fire Marshal supervisory / appeal / backstop role
  -> Department of Labor for elevator / conveyance and boiler scope
  -> Applicable statewide State Building Code adoption record
  -> Approved local fire-prevention or floodplain regulation record, if any
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | not selected | none | statewide | unresolved | pending |
| County | not selected | none | statewide | unresolved | pending |
| Municipality | not selected | none | statewide | unresolved | pending |
| Fire District | not selected | none | statewide/local | unresolved | pending |
| Special District | not selected | none | statewide/local | unresolved | pending |

### 7.3 AHJ Contact Data

No AHJ contact data was populated for this pass. The OSFM source registry provides statewide agency contact context but not local AHJ contacts.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Publisher / Custodian | Source Type | URL | Date / Version | Primary Use |
| --- | --- | --- | --- | --- | --- | --- |
| `src:usa-nc:gs-143-136-138` | N.C. Gen. Stat. Chapter 143, Article 9: Building Code Council, Residential Code Council, and North Carolina State Building Code | North Carolina General Assembly | statute HTML | https://www.ncleg.gov/enactedlegislation/statutes/html/byarticle/chapter_143/article_9.html | General Statutes include changes through S.L. 2026-7 per NCGA chapter page | Building Code Council/RCC creation, adoption authority, code volumes, local-code effect, local fire-prevention/floodplain rule, publication/distribution, appeals, interpretation choice. |
| `src:usa-nc:gs-143-139` | N.C. Gen. Stat. § 143-139, Enforcement of the North Carolina State Building Code | North Carolina General Assembly | statute HTML | https://www.ncleg.gov/enactedlegislation/statutes/html/bysection/chapter_143/gs_143-139.html | General Statutes include changes through S.L. 2026-7 per NCGA chapter page | State Fire Marshal general authority; local official cooperation; Department of Labor allocations for boilers/elevators; enforcement remedies. |
| `src:usa-nc:gs-160d-1102` | N.C. Gen. Stat. § 160D-1102, Building code administration | North Carolina General Assembly | statute HTML | https://www.ncleg.gov/EnactedLegislation/Statutes/HTML/BySection/Chapter_160D/GS_160D-1102.html | General Statutes include changes through S.L. 2026-7 per NCGA chapter page | Local inspection department options and State Fire Marshal backstop if local services fail. |
| `src:usa-nc:gs-160d-art11` | N.C. Gen. Stat. Chapter 160D, Article 11, Building Code Enforcement | North Carolina General Assembly | statute HTML | https://www.ncleg.gov/EnactedLegislation/Statutes/HTML/ByArticle/Chapter_160D/Article_11.html | General Statutes include changes through S.L. 2026-7 per NCGA chapter page | Building permits, local review timelines, inspections of work in progress, permit expiration, certificates, records. |
| `src:usa-nc:osfm-codes` | Codes | North Carolina Office of State Fire Marshal | agency webpage | https://www.ncosfm.gov/codes | accessed 2026-06-26 | OSFM summary that NC State Building Code is adopted/amended by Building Code Council for statewide implementation; Residential Code Council overview. |
| `src:usa-nc:osfm-codes-current-past` | Codes - Current and Past | North Carolina Office of State Fire Marshal | agency webpage / code portal | https://www.ncosfm.gov/codes/codes-current-and-past | accessed 2026-06-26 | Current code editions, 2018 code effective and transition dates, 2020 Electrical Code effective date, accessibility standard, current and past code links. |
| `src:usa-nc:osfm-electrical-division` | State Electrical Division | North Carolina Office of State Fire Marshal | agency webpage | https://www.ncosfm.gov/codes/state-electrical-division | accessed 2026-06-26 | State Electrical Division role interpreting State Electrical Code and providing access to adopted amendments. |
| `src:usa-nc:osfm-2025-delay-press` | North Carolina Delays Implementation of 2024 State Building Code | North Carolina Office of State Fire Marshal | agency press release | https://www.ncosfm.gov/news/press-releases/2025/04/07/north-carolina-delays-implementation-2024-state-building-code | 2025-04-07 | 2024 Code delay; continued effectiveness of 2018 Code; alternative-method use of 2024 Code; 2024 code volume list. |
| `src:usa-nc:osfm-2024-rrc-links` | 2025-0403 Rules Review Commission links to the 2024 North Carolina State Building Code | North Carolina Office of State Fire Marshal | agency webpage linking PDF | https://www.ncosfm.gov/2025-0403-rules-review-commission-links-2024-north-carolina-state-building-code | 2025-04-03 | Location of RRC link index for approved 2024 State Building Code rules; underlying links not parsed. |
| `src:usa-nc:osfm-2025-formal-interpretation-107` | December 5, 2025 OSFM Formal Interpretation re 2018 NCACP Section 107 Inspections | North Carolina Office of State Fire Marshal | formal interpretation PDF | https://www.ncosfm.gov/formal-interpretations/2025-1205-nchba-2018-ncacp-section-107-inspections/open | 2025-12-05 | Confirms 2018 Code remains effective, lists current 2018 code volumes and 2024 future volumes, describes S.L. 2025-2 conditions and non-mandatory 2024 status as of the letter. |
| `src:usa-nc:sl-2025-2` | Session Law 2025-2, House Bill 47 | North Carolina General Assembly | session law PDF | https://www.ncleg.gov/EnactedLegislation/SessionLaws/PDF/2025-2026/SL2025-2.pdf | enacted 2025-03-19 | Legal delay of 2024 Code effective date and certification trigger. |
| `src:usa-nc:ncdol-elevator` | Elevator | North Carolina Department of Labor | agency webpage | https://www.labor.nc.gov/safety-and-health/elevator | accessed 2026-06-26 | Elevator and Amusement Device Bureau enforcement of Elevator Safety Act; elevator design/construction/installation/inspection/certification scope. |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| `src:usa-nc:gs-143-136-138` | unofficial_html | NCGA chapter page states General Statutes on website are not official, although current through S.L. 2026-7. | Validate against official codification / session law before verified status. |
| `src:usa-nc:gs-143-139` | unofficial_html | NCGA statute HTML is convenient but carries the same NCGA caveat. | Treat as authoritative for partial verification; confirm for production. |
| `src:usa-nc:gs-160d-1102` | unofficial_html | NCGA statute HTML is convenient but carries the same NCGA caveat. | Treat as authoritative for partial verification; confirm for production. |
| `src:usa-nc:gs-160d-art11` | unofficial_html | NCGA statute HTML is convenient but carries the same NCGA caveat. | Treat as authoritative for partial verification; confirm for production. |
| `src:usa-nc:osfm-codes-current-past` | portal_summary | OSFM current-codes page is official but summarizes and links code texts/amendments; it is not a complete rule-history extraction. | Use for current-code edition status; parse linked amendments for verified status. |
| `src:usa-nc:osfm-2025-formal-interpretation-107` | formal_interpretation_context | The PDF is an OSFM formal interpretation focused on inspections. It is strong evidence for effective status and listed code volumes, but should be paired with source code text and adoption orders for production. | Use for partial verification; do not treat as complete adoption registry. |
| `src:usa-nc:osfm-2024-rrc-links` | link_index_only | Page identifies a PDF link index for 2024 RRC-approved rules; underlying RRC documents were not parsed. | Keep 2024 adoption details at future/pending status until link targets are extracted. |
| `src:usa-nc:sl-2025-2` | pdf_extracted | Session law PDF is official but date calculation requires a later State Fire Marshal certification letter, not found in this pass. | Monitor for certification event before calculating mandatory date. |

### 8.3 Supplemental Sources

None used. Non-official industry, municipal, and commentary sources were not used for report facts.

### 8.4 Source Extraction Metadata

| Extraction ID | Source IDs | Extracted Fields | Extracted By | Extracted Date | Notes |
| --- | --- | --- | --- | --- | --- |
| `extract:usa-nc:2026-06-26-authority` | `src:usa-nc:gs-143-136-138`; `src:usa-nc:gs-143-139`; `src:usa-nc:osfm-codes` | adopting authorities, authority split, State Fire Marshal enforcement role, code volumes | ChatGPT | 2026-06-26 | Core authority extracted from statute HTML and OSFM page. |
| `extract:usa-nc:2026-06-26-current-codes` | `src:usa-nc:osfm-codes-current-past`; `src:usa-nc:osfm-2025-formal-interpretation-107`; `src:usa-nc:osfm-electrical-division` | current editions, effective dates, transition dates, electrical split note | ChatGPT | 2026-06-26 | Electrical code requires second-pass reconciliation. |
| `extract:usa-nc:2026-06-26-local` | `src:usa-nc:gs-143-136-138`; `src:usa-nc:gs-160d-1102`; `src:usa-nc:gs-160d-art11` | local enforcement model, local amendment posture, permit lifecycle items | ChatGPT | 2026-06-26 | Local amendment registry not extracted. |
| `extract:usa-nc:2026-06-26-2024-delay` | `src:usa-nc:sl-2025-2`; `src:usa-nc:osfm-2025-delay-press`; `src:usa-nc:osfm-2025-formal-interpretation-107`; `src:usa-nc:osfm-2024-rrc-links` | 2024 Code future status, contingent trigger, alternative-method use | ChatGPT | 2026-06-26 | Monitor for State Fire Marshal certification. |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 1.00 | none | Core authority, current-code status, and local enforcement model are source-backed; not all local registries and amendment text are parsed. |
| report | risk.overall_confidence | 0.62 | verified | 1.00 | none | Conservative confidence due to unresolved amendment registry and adoption-date normalization gaps. |
| `ahj:usa-nc:building-code-council` | authority exists and duties | Building Code Council created; duties include review/consideration of revisions and amendments for commercial/multi-family code volumes. | verified_core | 0.85 | `src:usa-nc:gs-143-136-138` | Statute-level verified. |
| `ahj:usa-nc:residential-code-council` | authority exists and duties | Residential Code Council created; oversees residential construction code volumes. | verified_core | 0.80 | `src:usa-nc:gs-143-136-138` | Current membership/confirmation status remains a 2024 Code trigger. |
| `adoption:usa-nc:2018-state-building-code-suite` | effective date | 2019-01-01 | verified_core | 0.85 | `src:usa-nc:osfm-codes-current-past`; `src:usa-nc:osfm-2025-formal-interpretation-107` | OSFM portal and formal interpretation align. |
| `adoption:usa-nc:2018-state-building-code-suite` | transition period | 2018-07-01 through 2018-12-31 | partially_verified | 0.75 | `src:usa-nc:osfm-codes-current-past` | OSFM states availability during transition; detailed permit-date mechanics not parsed. |
| `adoption:usa-nc:electrical-2020-and-residential-scope-note` | current status | 2020 Electrical Code current, with formal-interpretation scope note for 2017 residential-scope provisions | partially_verified | 0.65 | `src:usa-nc:osfm-codes-current-past`; `src:usa-nc:osfm-2025-formal-interpretation-107` | Needs second-pass electrical amendment review. |
| `adoption:usa-nc:2024-state-building-code-suite-future` | mandatory date | contingent; not determined | verified_core | 0.85 | `src:usa-nc:sl-2025-2`; `src:usa-nc:osfm-2025-formal-interpretation-107` | Do not calculate until certification letter is found. |
| `local-enforcement:usa-nc` | model | statewide code with local inspection departments and State Fire Marshal backstop | verified_core | 0.78 | `src:usa-nc:gs-143-139`; `src:usa-nc:gs-160d-1102` | Local AHJ contacts not populated. |
| `local-amendment-rule:usa-nc` | model | statewide code with limited local fire-prevention/floodplain authority and Code Council approval requirement | partially_verified | 0.70 | `src:usa-nc:gs-143-136-138` | Registry and filing mechanics unresolved. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Body source IDs are present in section 8. |
| All authority IDs resolve | pass | Authority IDs used in tables are defined in sections 2 and 6. |
| All current code families have adoption records | pass | Matrix rows are present for all template families; current suite and electrical adoption records are explicit. |
| Building and operational fire code are separated | pass | Separate rows and notes distinguish State Fire Code volume from local fire-prevention-code authority. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Separate fields maintained; unresolved adoption dates remain unresolved. |
| Effective dates are valid ISO dates | pass | Entered dates use ISO format. |
| No impossible date sequences | pass | 2018 transition precedes 2019 effective date; 2024 dates remain contingent rather than calculated. |
| Transition rules have explicit trigger conditions | pass | 2018 transition and 2024 contingent trigger are documented. |
| Permit-date logic is captured where applicable | fail | General permit expiration and interpretation-choice rules are captured, but full code-edition permit-choice / vesting mechanics were not extracted. |
| Local enforcement model classified | pass | Model classified from G.S. 160D-1102 and G.S. 143-139. |
| Local amendment rule classified | pass | Statewide-code/local-fire-prevention/floodplain approval model classified. |
| AHJ confirmation metadata present | fail | No local AHJ contacts or jurisdiction boundary sources were populated. |
| Official-source caveats captured | pass | Caveats recorded for NCGA HTML, OSFM summary pages, formal interpretation, and 2024 link index. |
| 2024 Code mandatory status safe | pass | 2024 Code is marked future/not mandatory; trigger remains active_watch. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `issue:usa-nc:001` | high | 2024 Code effective date | The 2024 Code mandatory date is contingent on State Fire Marshal certification; no certification letter was extracted. | Monitor OSFM, NCGA, and Revisor channels for certification and calculate effective date only after certification. | null | null | open |
| `issue:usa-nc:002` | high | adoption-date normalization | 2018 and 2024 adoption dates and RRC approval dates were not fully normalized into per-volume records. | Parse BCC/RCC minutes, RRC approvals, North Carolina Register notices, and linked amendment documents. | null | null | open |
| `issue:usa-nc:003` | medium | electrical scope split | OSFM portal lists 2020 Electrical Code as current; OSFM formal interpretation identifies a 2017 residential-scope / 2020 other-structures split. | Reconcile against adopted electrical amendments and ICC/NFPA publication text. | null | null | open |
| `issue:usa-nc:004` | medium | local amendment registry | Statutory approval requirement for local fire-prevention and local code variants is verified, but registry/source path for approved local codes was not located. | Search BCC/RCC decisions, OSFM records, North Carolina Register notices, and local approval orders. | null | null | open |
| `issue:usa-nc:005` | medium | fire operational AHJ details | Fire Code volume and local fire-prevention code authority are verified, but operational enforcement divisions and local fire marshal roles were not deeply parsed. | Extract fire-code enforcement statutes/rules and OSFM fire-prevention guidance. | null | null | open |
| `issue:usa-nc:006` | low | specialized state programs | Elevator authority is verified; boiler and state-building program details were identified in statute but not separately normalized. | Add Department of Labor boiler source and Department of Administration / State Construction source if report scope requires them. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| `watch:usa-nc:osfm-current-codes` | `src:usa-nc:osfm-codes-current-past` | html_diff | monthly | changes to current-code list, electrical-code note, or 2024 effective-date notice | 2026-06-26 | active |
| `watch:usa-nc:2024-delay-press` | `src:usa-nc:osfm-2025-delay-press` | html_diff | monthly | update or superseding press release on 2024 Code effective date | 2026-06-26 | active |
| `watch:usa-nc:sl-2025-2-trigger` | `src:usa-nc:sl-2025-2` | certification_search | monthly | State Fire Marshal certification letter to Revisor and legislative leaders | 2026-06-26 | active |
| `watch:usa-nc:2024-rrc-links` | `src:usa-nc:osfm-2024-rrc-links` | pdf_link_diff | monthly | new or revised RRC link index for 2024 Code rules | 2026-06-26 | active |
| `watch:usa-nc:chapter-143` | `src:usa-nc:gs-143-136-138` | statute_diff | quarterly | amendments to Article 9 affecting code authority, local code effect, publication, or effective-date rules | 2026-06-26 | active |
| `watch:usa-nc:chapter-160d-article-11` | `src:usa-nc:gs-160d-art11` | statute_diff | quarterly | changes to local inspection, permit, plan review, private inspection, or CO rules | 2026-06-26 | active |
| `watch:usa-nc:dol-elevator` | `src:usa-nc:ncdol-elevator` | html_diff | quarterly | changes to elevator safety program authority or code adoption details | 2026-06-26 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-24 | Generated baseline draft report | `report:usa-nc` | none | Codex | Original uploaded draft contained state-specific frontmatter but unresolved authority and adoption placeholders. |
| 2026-06-26 | Populated core authority, current code matrix, local enforcement model, local amendment rule, source registry, and QA sections | `ahj:usa-nc:*`; `adoption:usa-nc:*`; `local-enforcement:usa-nc`; `local-amendment-rule:usa-nc`; `date-rule:usa-nc:*`; `watch:usa-nc:*` | `src:usa-nc:gs-143-136-138`; `src:usa-nc:gs-143-139`; `src:usa-nc:gs-160d-1102`; `src:usa-nc:gs-160d-art11`; `src:usa-nc:osfm-codes`; `src:usa-nc:osfm-codes-current-past`; `src:usa-nc:osfm-electrical-division`; `src:usa-nc:osfm-2025-delay-press`; `src:usa-nc:osfm-2025-formal-interpretation-107`; `src:usa-nc:sl-2025-2`; `src:usa-nc:ncdol-elevator` | ChatGPT | Upgraded status to partially_verified for core fields; kept unresolved fields explicit. |
