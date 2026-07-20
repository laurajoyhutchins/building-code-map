---
state:
  state_id: "US-SC"
  name: "South Carolina"
  abbreviation: "SC"
report:
  report_id: "state-report:usa-sc"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-26"
  last_verified: "2026-06-26"
  reviewed_by: null
risk:
  overall_confidence: 0.72 # 0.00 - 1.00
  risk_flags:
    - "2024_code_cycle_pending_2027_enforcement"
    - "local_permissive_code_adoptions_not_collected"
    - "operational_fire_inspection_overlap_requires_ahj_confirmation"
    - "online_regulation_copy_is_unannotated_courtesy_copy"
    - "ahj_boundary_dataset_not_selected"
    - "elevator_conveyance_authority_unresolved"
  open_questions_count: 6

---

# State Building Code Authority Report: South Carolina

## 1. Executive Summary

- **Authority model:** South Carolina uses a statewide adoption / local enforcement model. The South Carolina Building Codes Council (`ahj:usa-sc:bcc`) is the primary statewide building-code authority for the mandatory construction codes listed in S.C. Code Ann. §6-9-50. Counties and municipalities enforce the adopted building, energy, electrical, plumbing, mechanical, gas, and fire codes within their jurisdictions. The State Fire Marshal (`ahj:usa-sc:osfm`) has separate statewide fire-prevention and life-safety authority under Title 23, Chapter 9.

- **Statewide code status:** The current statewide mandatory construction-code set verified for this report is the 2021 South Carolina Building, Residential, Fire, Plumbing, Mechanical, and Fuel Gas Codes; the 2020 National Electrical Code; and the 2009 International Energy Conservation Code. The 2021 I-Code / 2020 NEC package was adopted by the Building Codes Council on 2021-10-06 with a 2023-01-01 implementation / enforcement date. The 2009 IECC is adopted by statute as the state Energy Standard and became effective 2013-01-01.

- **Local enforcement model:** Local governments are the ordinary enforcement layer. Each county must appoint or contract for a building official for the unincorporated area, and each municipality must appoint or contract for a building official within municipal limits. Local governments may contract with other governmental entities for permitting and enforcement services. State-agency and school-district projects have special statutory treatment and should not be treated as ordinary local-AHJ projects without project-specific review.

- **Local amendment posture:** Local jurisdictions may not write or publish other building codes in whole or in part. Statewide modifications must be approved by the Building Codes Council before the implementation date and are mandatory statewide. Local modifications are possible only after local governing-body approval and Building Codes Council approval, and only for local physical or climatological conditions. Permissive codes, including property maintenance, existing building, performance, and swimming-pool codes, require local ordinance adoption before local enforcement.

- **Known transition periods or pending changes:** The 2024 South Carolina Building Codes / 2023 NEC cycle is in process. LLR reports Council action on 2024 code modifications and a State Register publication on 2026-05-22. Proposed-regulation text for the 2024 cycle states a Building Codes Council enforcement-purpose effective date of 2027-01-01. This report treats the 2024 cycle as a monitored future / transition item until the effective code portal, final regulation text, and local enforcement guidance are reconciled.

- **Production readiness:** partially_ready_for_statewide_mandatory_code_lookup. The primary authority, current statewide mandatory code package, local enforcement model, local modification rule, and permit-date transition rule are source-backed. This file still needs structured amendment parsing, AHJ boundary/contact datasets, local permissive-code adoption collection, elevator/conveyance authority research, and a 2024-cycle legal-effective-date reconciliation before broader production use.

### Key Findings

```yaml
---
key_findings:
- topic: Primary adopting authority
  finding: "The South Carolina Building Codes Council reviews, adopts, modifies, and\
    \ promulgates the mandatory statewide building-code set referenced in \xA76-9-50."
  confidence: 0.9
  source_ids:
  - src:usa-sc:code-title-6-ch9
  - src:usa-sc:regs-ch8-bcc
  - src:usa-sc:llr-bcc-home
- topic: Current construction-code package
  finding: The current verified mandatory package is the 2021 IBC, IRC, IFC, IPC,
    IMC, IFGC, the 2020 NEC, and the 2009 IECC, with South Carolina modifications
    where applicable.
  confidence: 0.9
  source_ids:
  - src:usa-sc:llr-building-codes-in-effect-2022
  - src:usa-sc:llr-building-code-adoption
  - src:usa-sc:llr-2021-code-adoption-notice
- topic: Current 2021/2020 adoption and implementation dates
  finding: The 2021 I-Code / 2020 NEC package was adopted 2021-10-06 and implemented
    / enforceable locally 2023-01-01.
  confidence: 0.9
  source_ids:
  - src:usa-sc:llr-building-code-adoption
  - src:usa-sc:llr-building-codes-in-effect-2022
- topic: Energy code
  finding: The 2009 IECC is adopted by statute as the Energy Standard; all new and
    renovated buildings and additions constructed in South Carolina must comply.
  confidence: 0.88
  source_ids:
  - src:usa-sc:code-title-6-ch10-energy
  - src:usa-sc:llr-building-codes-in-effect-2022
- topic: Local enforcement
  finding: Counties and municipalities enforce the state building-code set and must
    appoint or contract for building officials, with county jurisdiction over unincorporated
    areas and municipal jurisdiction inside municipal limits.
  confidence: 0.86
  source_ids:
  - src:usa-sc:code-title-6-ch9
  - src:usa-sc:regs-ch8-bcc
- topic: Local amendments
  finding: Local code variations require local governing-body approval and Building
    Codes Council approval, and must be justified by physical or climatological conditions.
  confidence: 0.84
  source_ids:
  - src:usa-sc:code-title-6-ch9
  - src:usa-sc:regs-ch8-bcc
- topic: Fire authority
  finding: The Building Codes Council adopts the South Carolina Fire Code for construction-code
    purposes, while the State Fire Marshal retains statewide fire-prevention and fire-protection
    authority.
  confidence: 0.76
  source_ids:
  - src:usa-sc:code-title-6-ch9
  - src:usa-sc:code-title-23-ch9-fire-marshal
  - src:usa-sc:state-fire-codes-standards
- topic: Effective / operative date rule
  finding: Inspections generally follow the code in effect on the date the original
    building permit was issued, with statutory fallback rules where permit or application
    dates are unavailable.
  confidence: 0.88
  source_ids:
  - src:usa-sc:code-title-6-ch9
  - src:usa-sc:regs-ch8-bcc
- topic: Future 2024 cycle
  finding: The 2024 I-Code / 2023 NEC cycle has published modification materials and
    proposed regulations indicating 2027-01-01 enforcement-purpose effective date.
  confidence: 0.74
  source_ids:
  - src:usa-sc:llr-building-code-adoption
  - src:usa-sc:llr-2024-code-intent
  - src:usa-sc:state-register-2025-2024-code-proposed
  - src:usa-sc:llr-2026-2024-code-final-mods
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-sc:bcc |
| Authority name | South Carolina Building Codes Council |
| Authority type | State building-code council within the South Carolina Department of Labor, Licensing and Regulation |
| Legal basis | S.C. Code Ann. §§6-9-40, 6-9-50, 6-9-55, 6-9-105; S.C. Code Regs. Chapter 8, including Regs. 8-236, 8-240, and 8-245 |
| Role | Reviews, adopts, modifies, and promulgates the mandatory statewide building-code set; approves or denies statewide and local code modifications; sets implementation dates; regulates building code officer registration, special inspectors, and the modular building program |
| Enforcement model | State adoption and modification by Council; county and municipal local enforcement; special state/fire/reserved jurisdiction for some project classes |
| Source IDs | src:usa-sc:code-title-6-ch9; src:usa-sc:regs-ch8-bcc; src:usa-sc:llr-bcc-home; src:usa-sc:llr-building-code-adoption |
| Verification status | verified_for_core_fields |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-sc:bcc | South Carolina Building Codes Council | Statewide adoption and modification of South Carolina Building Code / IBC | S.C. Code Ann. §§6-9-40 and 6-9-50; Regs. 8-236, 8-240, Article 8 | src:usa-sc:code-title-6-ch9; src:usa-sc:regs-ch8-bcc; src:usa-sc:llr-building-codes-in-effect-2022 | verified_for_current_package |
| Residential | ahj:usa-sc:bcc | South Carolina Building Codes Council | Statewide adoption and modification of South Carolina Residential Code / IRC; one- and two-family dwelling changes require regulation process before enforcement | S.C. Code Ann. §§6-9-40, 6-9-50, 6-9-55; Regs. 8-236, 8-240, Article 12 | src:usa-sc:code-title-6-ch9; src:usa-sc:regs-ch8-bcc; src:usa-sc:llr-building-codes-in-effect-2022 | verified_for_current_package |
| Existing Building / Rehabilitation | ahj:usa-sc:bcc-and-local-permissive | South Carolina Building Codes Council plus local adopting jurisdictions | Hybrid. Existing building code is a permissive code for local adoption under §6-9-60, but the 2021 South Carolina Building Code modification summary states that the South Carolina Existing Building Code applies to repair, alteration, change of occupancy, addition, and relocation of existing buildings. Scope needs legal reconciliation before automated use. | S.C. Code Ann. §6-9-60; Reg. 8-801 | src:usa-sc:code-title-6-ch9; src:usa-sc:regs-ch8-bcc; src:usa-sc:llr-building-codes-in-effect-2022 | partially_verified_scope_needs_review |
| Mechanical | ahj:usa-sc:bcc | South Carolina Building Codes Council | Statewide adoption and modification of South Carolina Mechanical Code / IMC | S.C. Code Ann. §§6-9-40 and 6-9-50; Regs. 8-236, 8-240, Article 13 | src:usa-sc:code-title-6-ch9; src:usa-sc:regs-ch8-bcc; src:usa-sc:llr-building-codes-in-effect-2022 | verified_for_current_package |
| Plumbing | ahj:usa-sc:bcc | South Carolina Building Codes Council | Statewide adoption and modification of South Carolina Plumbing Code / IPC | S.C. Code Ann. §§6-9-40 and 6-9-50; Regs. 8-236, 8-240, Article 14 | src:usa-sc:code-title-6-ch9; src:usa-sc:regs-ch8-bcc; src:usa-sc:llr-building-codes-in-effect-2022 | verified_for_current_package |
| Fuel Gas | ahj:usa-sc:bcc | South Carolina Building Codes Council | Statewide adoption and modification of South Carolina Fuel Gas Code / IFGC | S.C. Code Ann. §§6-9-40 and 6-9-50; Regs. 8-236, 8-240, Article 10 | src:usa-sc:code-title-6-ch9; src:usa-sc:regs-ch8-bcc; src:usa-sc:llr-building-codes-in-effect-2022 | verified_for_current_package |
| Electrical | ahj:usa-sc:bcc | South Carolina Building Codes Council | Statewide adoption and modification of the National Electrical Code as part of the building-code package | S.C. Code Ann. §6-9-50; Regs. 8-236, 8-240, Article 11 | src:usa-sc:code-title-6-ch9; src:usa-sc:regs-ch8-bcc; src:usa-sc:llr-building-codes-in-effect-2022 | verified_for_current_package |
| Energy | ahj:usa-sc:energy-standard | South Carolina General Assembly / South Carolina Building Codes Council / local building officials | Statutory Energy Standard; local enforcement; Council variance review for special local conditions | S.C. Code Ann. Title 6, Chapter 10; S.C. Code Ann. §6-9-50 | src:usa-sc:code-title-6-ch10-energy; src:usa-sc:llr-building-codes-in-effect-2022 | verified_for_current_standard |
| Fire - construction references | ahj:usa-sc:bcc | South Carolina Building Codes Council | Statewide adoption and modification of South Carolina Fire Code / IFC for construction-code purposes | S.C. Code Ann. §§6-9-40 and 6-9-50; Regs. 8-236, 8-240, Article 9 | src:usa-sc:code-title-6-ch9; src:usa-sc:regs-ch8-bcc; src:usa-sc:llr-building-codes-in-effect-2022 | verified_for_current_package |
| Fire - operational / prevention code | ahj:usa-sc:osfm | Office of State Fire Marshal, Division of Fire and Life Safety, LLR | Statewide fire-prevention, fire-protection, egress, fire alarm / extinguishing systems, fire investigation, citations, and related licensing / permitting functions. Relationship to local fire-code inspection allocation must be resolved by occupancy and project type. | S.C. Code Ann. Title 23, Chapter 9; S.C. Code Ann. §6-9-5(C) | src:usa-sc:code-title-23-ch9-fire-marshal; src:usa-sc:code-title-6-ch9; src:usa-sc:state-fire-codes-standards | partially_verified_scope_needs_ahj_review |
| Accessibility | ahj:usa-sc:bcc-accessibility | South Carolina Building Codes Council / Accessibility Committee / local building officials | Accessibility regulations and ICC/ANSI A117.1 enforcement by local building officials, with Council enforcement where no local building official exists | S.C. Code Ann. Title 10, Chapter 5; Regs. 8-700 through 8-703 | src:usa-sc:code-title-10-ch5-accessibility; src:usa-sc:regs-ch8-bcc; src:usa-sc:llr-building-codes-in-effect-2022 | partially_verified |
| Elevator / Conveyance | ahj:usa-sc:elevator-unresolved | Unresolved in this report | Elevator / conveyance authority was not researched to primary-source level in this pass. | unresolved | none | unresolved |

### 2.3 Authority Hierarchy Notes

South Carolina's statewide building-code model has three distinct layers that must not be collapsed:

1. **State adoption / modification:** The Building Codes Council adopts and modifies the mandatory code families, assigns implementation dates, and approves statewide and local modifications.
2. **Local enforcement:** Counties and municipalities enforce the state code package within their jurisdictions, appoint or contract for building officials, and may adopt local administrative provisions by ordinance where state-level administrative provisions were not adopted.
3. **Specialized / reserved authority:** The State Fire Marshal retains statewide fire-prevention and fire-protection powers under Title 23, and certain state-agency and school-district projects are outside ordinary local permit / license requirements. The modular building program also has a state-level manufacturing-plan and inspection role while local building officials regulate site-specific connections and foundations.

The local enforcement layer is not the same as local amendment authority. Local enforcement is broadly required, but local code variations need Building Codes Council approval and must be justified by physical or climatological conditions.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-sc:001 | ahj:usa-sc:bcc | adopts_and_modifies | Mandatory statewide building, residential, fire, plumbing, mechanical, fuel gas, electrical, and energy code families listed in §6-9-50 | src:usa-sc:code-title-6-ch9; src:usa-sc:regs-ch8-bcc | verified |
| edge:usa-sc:002 | ahj:usa-sc:bcc | sets_implementation_date_for | Council-adopted building codes and modifications | src:usa-sc:code-title-6-ch9; src:usa-sc:llr-building-codes-in-effect-2022 | verified |
| edge:usa-sc:003 | ahj:usa-sc:bcc | relies_on_local_enforcement_by | Counties and municipalities | src:usa-sc:code-title-6-ch9; src:usa-sc:regs-ch8-bcc | verified |
| edge:usa-sc:004 | local_jurisdictions | appoint_or_contract_for | Building official services | src:usa-sc:code-title-6-ch9 | verified |
| edge:usa-sc:005 | local_jurisdictions | may_contract_with | Other governmental entities for permits and code enforcement | src:usa-sc:code-title-6-ch9 | verified |
| edge:usa-sc:006 | ahj:usa-sc:bcc | approves_or_denies | Local building-code modifications based on physical or climatological conditions | src:usa-sc:code-title-6-ch9; src:usa-sc:regs-ch8-bcc | verified |
| edge:usa-sc:007 | ahj:usa-sc:osfm | retains_statewide_authority_for | Fire prevention, fire protection, egress, fire alarm / extinguishing systems, fire investigations, and related fire code / standards matters | src:usa-sc:code-title-23-ch9-fire-marshal; src:usa-sc:code-title-6-ch9 | partially_verified |
| edge:usa-sc:008 | ahj:usa-sc:state-engineer-and-osfm | exercises_reserved_jurisdiction_for | State-agency and school-district projects covered by §6-9-110 | src:usa-sc:code-title-6-ch9; src:usa-sc:state-fire-codes-standards | partially_verified |
| edge:usa-sc:009 | ahj:usa-sc:energy-standard | enforced_by | Local building officials / local jurisdictions | src:usa-sc:code-title-6-ch10-energy | verified |
| edge:usa-sc:010 | ahj:usa-sc:bcc-accessibility | fallback_enforcement_for | Accessibility regulations where a county or municipality has no building official | src:usa-sc:regs-ch8-bcc | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | 2021 South Carolina Building Code | International Building Code | 2021 | current_statewide_mandatory_with_sc_modifications | 2021-10-06 | 2023-01-01 | 2023-01-01 | 2023-01-01 | Permit applications approved before the implementation date may be completed and inspected under the code in effect when the original permit was issued; statutory inspection rule also uses original permit date with fallback dates. | src:usa-sc:llr-building-code-adoption; src:usa-sc:llr-building-codes-in-effect-2022; src:usa-sc:regs-ch8-bcc; src:usa-sc:code-title-6-ch9 |
| Residential | 2021 South Carolina Residential Code | International Residential Code | 2021 | current_statewide_mandatory_with_sc_modifications | 2021-10-06 | 2023-01-01 | 2023-01-01 | 2023-01-01 | Same permit-date / implementation-date transition rule; one- and two-family construction requirements must not be enforced until required regulations are effective. | src:usa-sc:llr-building-code-adoption; src:usa-sc:llr-building-codes-in-effect-2022; src:usa-sc:regs-ch8-bcc; src:usa-sc:code-title-6-ch9 |
| Existing Building / Rehabilitation | South Carolina Existing Building Code / local permissive IEBC adoption | International Existing Building Code | latest edition locally permissive; 2021 SCBC references SC Existing Building Code | hybrid_partially_verified | null | null | null | null | Permissive code local adoption requires ordinance before enforcement; SCBC modification states the South Carolina Existing Building Code applies to repair, alteration, change of occupancy, addition, and relocation of existing buildings. Scope requires legal reconciliation before automated determinations. | src:usa-sc:code-title-6-ch9; src:usa-sc:regs-ch8-bcc; src:usa-sc:llr-building-codes-in-effect-2022 |
| Mechanical | 2021 South Carolina Mechanical Code | International Mechanical Code | 2021 | current_statewide_mandatory_with_sc_modifications | 2021-10-06 | 2023-01-01 | 2023-01-01 | 2023-01-01 | Same permit-date / implementation-date transition rule. | src:usa-sc:llr-building-code-adoption; src:usa-sc:llr-building-codes-in-effect-2022; src:usa-sc:regs-ch8-bcc; src:usa-sc:code-title-6-ch9 |
| Plumbing | 2021 South Carolina Plumbing Code | International Plumbing Code | 2021 | current_statewide_mandatory_with_sc_modifications | 2021-10-06 | 2023-01-01 | 2023-01-01 | 2023-01-01 | Same permit-date / implementation-date transition rule. | src:usa-sc:llr-building-code-adoption; src:usa-sc:llr-building-codes-in-effect-2022; src:usa-sc:regs-ch8-bcc; src:usa-sc:code-title-6-ch9 |
| Fuel Gas | 2021 South Carolina Fuel Gas Code | International Fuel Gas Code | 2021 | current_statewide_mandatory_with_sc_modifications | 2021-10-06 | 2023-01-01 | 2023-01-01 | 2023-01-01 | Same permit-date / implementation-date transition rule. | src:usa-sc:llr-building-code-adoption; src:usa-sc:llr-building-codes-in-effect-2022; src:usa-sc:regs-ch8-bcc; src:usa-sc:code-title-6-ch9 |
| Electrical | 2020 National Electrical Code with South Carolina modifications | NFPA 70, National Electrical Code | 2020 | current_statewide_mandatory_with_sc_modifications | 2021-10-06 | 2023-01-01 | 2023-01-01 | 2023-01-01 | Same permit-date / implementation-date transition rule. | src:usa-sc:llr-building-code-adoption; src:usa-sc:llr-building-codes-in-effect-2022; src:usa-sc:regs-ch8-bcc; src:usa-sc:code-title-6-ch9 |
| Energy | South Carolina Energy Standard | International Energy Conservation Code | 2009 | current_statewide_mandatory_by_statute | 2012-04-02 | 2013-01-01 | 2013-01-01 | 2013-01-01 | State statute applies to new and renovated buildings and additions. Earlier permits are protected by act-specific transition notes; local variance process available for special local conditions. | src:usa-sc:code-title-6-ch10-energy; src:usa-sc:llr-building-codes-in-effect-2022 |
| Fire - construction references | 2021 South Carolina Fire Code | International Fire Code | 2021 | current_statewide_mandatory_with_sc_modifications | 2021-10-06 | 2023-01-01 | 2023-01-01 | 2023-01-01 | Same permit-date / implementation-date transition rule for construction-code enforcement; State Fire Marshal operational authority remains separately applicable. | src:usa-sc:llr-building-code-adoption; src:usa-sc:llr-building-codes-in-effect-2022; src:usa-sc:regs-ch8-bcc; src:usa-sc:code-title-6-ch9; src:usa-sc:code-title-23-ch9-fire-marshal |
| Fire - operational / prevention code | South Carolina Fire Code plus State Fire Marshal regulations / standards | IFC and nationally recognized fire-prevention / protection standards as prescribed by law or regulation | 2021 IFC for current BCC package; OSFM operational standards vary by program | partially_verified | 2021-10-06 for IFC package | 2023-01-01 for IFC package | 2023-01-01 for IFC package | 2023-01-01 for IFC package | Construction-code date logic applies to IFC package; operational fire-prevention inspections and OSFM program rules require occupancy / AHJ review. | src:usa-sc:llr-building-codes-in-effect-2022; src:usa-sc:code-title-23-ch9-fire-marshal; src:usa-sc:state-fire-codes-standards |
| Accessibility | ICC/ANSI A117.1, Accessible and Usable Buildings and Facilities | ICC/ANSI A117.1 | latest published edition | current_statewide_mandatory_for_accessibility_scope | null | null | null | null | Local building officials enforce accessibility regulations; Building Codes Council enforces where no local building official exists. Other accessibility laws may also apply. | src:usa-sc:code-title-10-ch5-accessibility; src:usa-sc:regs-ch8-bcc; src:usa-sc:llr-building-codes-in-effect-2022 |
| Elevator / Conveyance | unresolved | unresolved | unresolved | unresolved | null | null | null | null | Elevator / conveyance authority and current technical standard were not resolved in this report. | none |

### 3.2 Adoption Records

#### adopt:usa-sc:2021-ibc

| Field | Value |
| --- | --- |
| Code family | Building |
| Authority ID | ahj:usa-sc:bcc |
| State code name | 2021 South Carolina Building Code |
| Base model code | 2021 International Building Code |
| Adoption date | 2021-10-06 |
| Effective date | 2023-01-01 |
| Operative date | 2023-01-01 |
| Mandatory date | 2023-01-01 |
| Scope | Mandatory statewide construction-code package with South Carolina modifications |
| Source IDs | src:usa-sc:llr-building-code-adoption; src:usa-sc:llr-building-codes-in-effect-2022; src:usa-sc:regs-ch8-bcc |
| Confidence | 0.90 |

#### adopt:usa-sc:2021-irc

| Field | Value |
| --- | --- |
| Code family | Residential |
| Authority ID | ahj:usa-sc:bcc |
| State code name | 2021 South Carolina Residential Code |
| Base model code | 2021 International Residential Code |
| Adoption date | 2021-10-06 |
| Effective date | 2023-01-01 |
| Operative date | 2023-01-01 |
| Mandatory date | 2023-01-01 |
| Scope | Mandatory statewide construction-code package with South Carolina modifications |
| Source IDs | src:usa-sc:llr-building-code-adoption; src:usa-sc:llr-building-codes-in-effect-2022; src:usa-sc:regs-ch8-bcc; src:usa-sc:code-title-6-ch9 |
| Confidence | 0.88 |

#### adopt:usa-sc:2021-ifc

| Field | Value |
| --- | --- |
| Code family | Fire - construction references |
| Authority ID | ahj:usa-sc:bcc |
| State code name | 2021 South Carolina Fire Code |
| Base model code | 2021 International Fire Code |
| Adoption date | 2021-10-06 |
| Effective date | 2023-01-01 |
| Operative date | 2023-01-01 |
| Mandatory date | 2023-01-01 |
| Scope | Mandatory statewide construction-code package with South Carolina modifications |
| Source IDs | src:usa-sc:llr-building-code-adoption; src:usa-sc:llr-building-codes-in-effect-2022; src:usa-sc:regs-ch8-bcc |
| Confidence | 0.88 |

#### adopt:usa-sc:2021-ipc

| Field | Value |
| --- | --- |
| Code family | Plumbing |
| Authority ID | ahj:usa-sc:bcc |
| State code name | 2021 South Carolina Plumbing Code |
| Base model code | 2021 International Plumbing Code |
| Adoption date | 2021-10-06 |
| Effective date | 2023-01-01 |
| Operative date | 2023-01-01 |
| Mandatory date | 2023-01-01 |
| Scope | Mandatory statewide construction-code package with South Carolina modifications |
| Source IDs | src:usa-sc:llr-building-code-adoption; src:usa-sc:llr-building-codes-in-effect-2022; src:usa-sc:regs-ch8-bcc |
| Confidence | 0.88 |

#### adopt:usa-sc:2021-imc

| Field | Value |
| --- | --- |
| Code family | Mechanical |
| Authority ID | ahj:usa-sc:bcc |
| State code name | 2021 South Carolina Mechanical Code |
| Base model code | 2021 International Mechanical Code |
| Adoption date | 2021-10-06 |
| Effective date | 2023-01-01 |
| Operative date | 2023-01-01 |
| Mandatory date | 2023-01-01 |
| Scope | Mandatory statewide construction-code package with South Carolina modifications |
| Source IDs | src:usa-sc:llr-building-code-adoption; src:usa-sc:llr-building-codes-in-effect-2022; src:usa-sc:regs-ch8-bcc |
| Confidence | 0.88 |

#### adopt:usa-sc:2021-ifgc

| Field | Value |
| --- | --- |
| Code family | Fuel Gas |
| Authority ID | ahj:usa-sc:bcc |
| State code name | 2021 South Carolina Fuel Gas Code |
| Base model code | 2021 International Fuel Gas Code |
| Adoption date | 2021-10-06 |
| Effective date | 2023-01-01 |
| Operative date | 2023-01-01 |
| Mandatory date | 2023-01-01 |
| Scope | Mandatory statewide construction-code package with South Carolina modifications |
| Source IDs | src:usa-sc:llr-building-code-adoption; src:usa-sc:llr-building-codes-in-effect-2022; src:usa-sc:regs-ch8-bcc |
| Confidence | 0.88 |

#### adopt:usa-sc:2020-nec

| Field | Value |
| --- | --- |
| Code family | Electrical |
| Authority ID | ahj:usa-sc:bcc |
| State code name | 2020 National Electrical Code with South Carolina modifications |
| Base model code | NFPA 70, National Electrical Code |
| Adoption date | 2021-10-06 |
| Effective date | 2023-01-01 |
| Operative date | 2023-01-01 |
| Mandatory date | 2023-01-01 |
| Scope | Mandatory statewide construction-code package with South Carolina modifications |
| Source IDs | src:usa-sc:llr-building-code-adoption; src:usa-sc:llr-building-codes-in-effect-2022; src:usa-sc:regs-ch8-bcc |
| Confidence | 0.88 |

#### adopt:usa-sc:2009-iecc

| Field | Value |
| --- | --- |
| Code family | Energy |
| Authority ID | ahj:usa-sc:energy-standard |
| State code name | South Carolina Energy Standard |
| Base model code | 2009 International Energy Conservation Code |
| Adoption date | 2012-04-02 |
| Effective date | 2013-01-01 |
| Operative date | 2013-01-01 |
| Mandatory date | 2013-01-01 |
| Scope | Mandatory statewide statutory energy standard for new and renovated buildings and additions |
| Source IDs | src:usa-sc:code-title-6-ch10-energy; src:usa-sc:llr-building-codes-in-effect-2022 |
| Confidence | 0.88 |

#### adopt:usa-sc:accessibility-a117-1

| Field | Value |
| --- | --- |
| Code family | Accessibility |
| Authority ID | ahj:usa-sc:bcc-accessibility |
| State code name | ICC/ANSI A117.1, Accessible and Usable Buildings and Facilities |
| Base model code | ICC/ANSI A117.1 |
| Adoption date | null |
| Effective date | null |
| Operative date | null |
| Mandatory date | null |
| Scope | Mandatory accessibility standard enforced by local building officials under state accessibility law and regulations |
| Source IDs | src:usa-sc:code-title-10-ch5-accessibility; src:usa-sc:regs-ch8-bcc; src:usa-sc:llr-building-codes-in-effect-2022 |
| Confidence | 0.72 |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

South Carolina distinguishes adoption dates from implementation / enforcement dates. For the current verified mandatory package, the Building Codes Council adopted the 2021 I-Code / 2020 NEC package on 2021-10-06 and established 2023-01-01 as the local-jurisdiction implementation date. The state inspection-date statute generally keys building inspections to the code in effect on the date the original building permit was issued; if no original permit date can be found, the completed application date is used, with a further document-based fallback if neither permit nor application date is available. For changes in structure, the applicable code is tied to the date of application or permit. For changes in use, the applicable code is tied to the date of inspection.

The 2024 code cycle is separately tracked as a future transition. Proposed regulations for the 2024 cycle state that the revised regulations will become effective by law upon General Assembly approval and State Register publication, while the Building Codes Council approved a 2027-01-01 enforcement-purpose effective date. This report does not treat 2024 codes as the current enforceable package until the final codified regulation and LLR current-code page are reconciled.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-sc:adoption-process | Building Codes Council code cycles | Adoption procedure | 180-day minimum written-comment period; Council action within 180 days after comment period ends | Notice of intention to adopt a code, adopt a new edition, or modify an existing code | Not applicable | src:usa-sc:code-title-6-ch9; src:usa-sc:regs-ch8-bcc | verified |
| date-rule:usa-sc:council-effective-date | Building Codes Council adopted or modified codes | Effective date selection | First day of January or July selected by Council | Council adoption / modification of code package | Depends on implementation / permit-date rule | src:usa-sc:code-title-6-ch9 | verified |
| date-rule:usa-sc:2021-package | 2021 IBC, IRC, IFC, IPC, IMC, IFGC and 2020 NEC | Implementation / local enforcement date | 2023-01-01 | 2021 code package adopted 2021-10-06 | Yes, for qualifying pre-implementation permit applications | src:usa-sc:llr-building-code-adoption; src:usa-sc:llr-building-codes-in-effect-2022; src:usa-sc:regs-ch8-bcc | verified |
| date-rule:usa-sc:original-permit-date | Building inspections generally | Permit-date rule | Code in effect on original building permit issuance date | Building inspection under §6-9-130 | Yes | src:usa-sc:code-title-6-ch9; src:usa-sc:regs-ch8-bcc | verified |
| date-rule:usa-sc:no-permit-date-fallback | Building inspections where original permit date is unavailable | Fallback date rule | Completed application date; if unavailable, nearest possible date using available documents | No original permit date, application date, or permit record available | Yes, if earlier date controls | src:usa-sc:code-title-6-ch9 | verified |
| date-rule:usa-sc:change-in-structure | Inspection connected with a change in structure | Application / permit date rule | Date of application or permit | Change in structure | Not specified beyond statutory date rule | src:usa-sc:code-title-6-ch9 | verified |
| date-rule:usa-sc:change-in-use | Inspection connected with change of use | Inspection-date rule | Date of inspection | Change of use for building or space | No, inspection-date rule applies | src:usa-sc:code-title-6-ch9 | verified |
| date-rule:usa-sc:energy-2009-iecc | South Carolina Energy Standard | Statutory effective / mandatory date | 2013-01-01 | 2012 Act 143 update to 2009 IECC | Prior permits protected by statutory transition note | src:usa-sc:code-title-6-ch10-energy; src:usa-sc:llr-building-codes-in-effect-2022 | verified |
| date-rule:usa-sc:2024-cycle-watch | 2024 IBC, IRC, IFC, IPC, IMC, IFGC; 2023 NEC | Future enforcement-purpose effective date | 2027-01-01 | Final approval and State Register publication of 2024-cycle regulations, then Council enforcement date | Expected transition; exact final local enforcement guidance needs reconciliation | src:usa-sc:llr-building-code-adoption; src:usa-sc:state-register-2025-2024-code-proposed; src:usa-sc:llr-2026-2024-code-final-mods | partially_verified |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | 2024 South Carolina Building Code / 2024 IBC with SC modifications | 2025-01-24 | null | 2026-05-22 | 2027-01-01 | 2027-01-01 | active | src:usa-sc:llr-building-code-adoption; src:usa-sc:llr-2024-code-intent; src:usa-sc:state-register-2025-2024-code-proposed; src:usa-sc:llr-2026-2024-code-final-mods | LLR reports final modifications published in the State Register on 2026-05-22. Proposed regulations state 2027-01-01 enforcement-purpose effective date. Treat adoption date as unresolved until final adoption record is parsed. |
| Residential | 2024 South Carolina Residential Code / 2024 IRC with SC modifications | 2025-01-24 | null | 2026-05-22 | 2027-01-01 | 2027-01-01 | active | src:usa-sc:llr-building-code-adoption; src:usa-sc:llr-2024-code-intent; src:usa-sc:state-register-2025-2024-code-proposed; src:usa-sc:llr-2026-2024-code-final-mods | Same 2024 cycle caveat. |
| Fire - construction references | 2024 South Carolina Fire Code / 2024 IFC with SC modifications | 2025-01-24 | null | 2026-05-22 | 2027-01-01 | 2027-01-01 | active | src:usa-sc:llr-building-code-adoption; src:usa-sc:llr-2024-code-intent; src:usa-sc:state-register-2025-2024-code-proposed; src:usa-sc:llr-2026-2024-code-final-mods | Same 2024 cycle caveat; operational fire implications need OSFM confirmation. |
| Mechanical | 2024 South Carolina Mechanical Code / 2024 IMC with SC modifications | 2025-01-24 | null | 2026-05-22 | 2027-01-01 | 2027-01-01 | active | src:usa-sc:llr-building-code-adoption; src:usa-sc:llr-2024-code-intent; src:usa-sc:state-register-2025-2024-code-proposed; src:usa-sc:llr-2026-2024-code-final-mods | Same 2024 cycle caveat. |
| Plumbing | 2024 South Carolina Plumbing Code / 2024 IPC with SC modifications | 2025-01-24 | null | 2026-05-22 | 2027-01-01 | 2027-01-01 | active | src:usa-sc:llr-building-code-adoption; src:usa-sc:llr-2024-code-intent; src:usa-sc:state-register-2025-2024-code-proposed; src:usa-sc:llr-2026-2024-code-final-mods | Same 2024 cycle caveat. |
| Fuel Gas | 2024 South Carolina Fuel Gas Code / 2024 IFGC with SC modifications | 2025-01-24 | null | 2026-05-22 | 2027-01-01 | 2027-01-01 | active | src:usa-sc:llr-building-code-adoption; src:usa-sc:llr-2024-code-intent; src:usa-sc:state-register-2025-2024-code-proposed; src:usa-sc:llr-2026-2024-code-final-mods | Same 2024 cycle caveat. |
| Electrical | 2023 National Electrical Code with SC modifications | 2025-01-24 | null | 2026-05-22 | 2027-01-01 | 2027-01-01 | active | src:usa-sc:llr-building-code-adoption; src:usa-sc:llr-2024-code-intent; src:usa-sc:state-register-2025-2024-code-proposed; src:usa-sc:llr-2026-2024-code-final-mods | Same 2024 cycle caveat. |
| Energy | Future IECC edition | null | null | null | null | null | monitor | src:usa-sc:code-title-6-ch10-energy; src:usa-sc:llr-building-codes-in-effect-2022 | LLR notes future IECC updates require statutory amendment. No later statutory IECC edition was verified for this report. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-sc:state-school-projects | Building / fire / local permitting | State-agency permanent improvement projects and specified school-district facilities | Project is a state department, institution, agency, or covered school-district project | Local ordinances or regulations requiring purchase or acquisition of permits, licenses, or other devices to enforce building standards do not apply; State Engineer and deputy state fire marshal authority may control. | src:usa-sc:code-title-6-ch9; src:usa-sc:state-fire-codes-standards | partially_verified |
| applicability-rule:usa-sc:fire-marshal-title-23 | Fire - operational / prevention code | Fire prevention, explosives / combustibles, fire alarm and extinguishing systems, means of egress, investigations, fire-protection codes and standards | Matter falls within State Fire Marshal statewide authority | State Fire Marshal authority under Title 23 remains applicable and is not displaced by Title 6 local inspection allocation. | src:usa-sc:code-title-23-ch9-fire-marshal; src:usa-sc:code-title-6-ch9 | partially_verified |
| applicability-rule:usa-sc:energy-standard | Energy | New and renovated buildings and additions constructed within South Carolina | Construction, renovation, or addition within statutory Energy Standard scope | 2009 IECC applies statewide as the Energy Standard; local building officials enforce and local jurisdictions may seek Council variance for special local conditions. | src:usa-sc:code-title-6-ch10-energy | verified |
| applicability-rule:usa-sc:permissive-codes | Existing Building / Property Maintenance / Performance / Swimming Pool and Spa | Permissive code families listed in §6-9-60 | Local ordinance adoption before enforcement | Municipalities and counties may adopt only the latest editions of listed permissive codes by reference; local ordinance adoption is required before enforcement. | src:usa-sc:code-title-6-ch9; src:usa-sc:llr-building-codes-in-effect-2022 | verified_for_rule_unverified_for_local_adoptions |
| applicability-rule:usa-sc:admin-chapters | All mandatory and permissive codes | Administrative chapters / administrative procedures | Local government wants to use model-code administrative provisions or local administrative policies | Administrative provisions are not included in mandatory or permissive code adoption unless adopted locally by ordinance or otherwise adopted by Council for uniform application. | src:usa-sc:code-title-6-ch9; src:usa-sc:regs-ch8-bcc; src:usa-sc:llr-building-codes-in-effect-2022 | verified |
| applicability-rule:usa-sc:accessibility | Accessibility | Government and public buildings; scope as provided by accessibility statutes and regulations | Construction, alteration, leasing, renovation, or occupancy change within accessibility regulation scope | Local building officials enforce accessibility regulations; Building Codes Council enforces where no local building official exists. | src:usa-sc:code-title-10-ch5-accessibility; src:usa-sc:regs-ch8-bcc | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** statewide_modifications_by_building_codes_council_with_limited_local_modifications

**State amendment structure:** South Carolina Code of Regulations Chapter 8, State Register notices / final regulations, LLR Building Codes Council code-adoption pages and modification indexes, and integrated South Carolina code publications where available.

**Amendment parsing status:** partially_started

Notes:

- Statewide modifications must be approved by the Building Codes Council before the established implementation date and are mandatory for all jurisdictions in the state.
- Local modifications require approval by the local governing body before Council review and do not take effect until Council approval.
- Local modifications are limited to physical or climatological conditions and apply strictly within the approved jurisdiction.
- The 2024-cycle modification set is published in LLR materials and the State Register; this report captures the cycle and dates but does not parse every modification into field-level records.

### 5.2 State Amendment Sources

| Amendment Source ID | Applies To | Publication Path | Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- |
| amend-source:usa-sc:2021-ch8-article8 | 2021 South Carolina Building Code / IBC | S.C. Code Regs. Chapter 8, Article 8 | parsed_for_authority_and_selected_scope | src:usa-sc:regs-ch8-bcc | Includes modification summary and South Carolina Existing Building Code reference. |
| amend-source:usa-sc:2021-ch8-article9 | 2021 South Carolina Fire Code / IFC | S.C. Code Regs. Chapter 8, Article 9 | source_identified_not_fully_parsed | src:usa-sc:regs-ch8-bcc | Full amendment extraction remains open. |
| amend-source:usa-sc:2021-ch8-article10 | 2021 South Carolina Fuel Gas Code / IFGC | S.C. Code Regs. Chapter 8, Article 10 | source_identified_not_fully_parsed | src:usa-sc:regs-ch8-bcc | Full amendment extraction remains open. |
| amend-source:usa-sc:2021-ch8-article11 | 2020 National Electrical Code | S.C. Code Regs. Chapter 8, Article 11 | source_identified_not_fully_parsed | src:usa-sc:regs-ch8-bcc | Full amendment extraction remains open. |
| amend-source:usa-sc:2021-ch8-article12 | 2021 South Carolina Residential Code / IRC | S.C. Code Regs. Chapter 8, Article 12 | source_identified_not_fully_parsed | src:usa-sc:regs-ch8-bcc | Full amendment extraction remains open. |
| amend-source:usa-sc:2021-ch8-article13 | 2021 South Carolina Mechanical Code / IMC | S.C. Code Regs. Chapter 8, Article 13 | source_identified_not_fully_parsed | src:usa-sc:regs-ch8-bcc | Full amendment extraction remains open. |
| amend-source:usa-sc:2021-ch8-article14 | 2021 South Carolina Plumbing Code / IPC | S.C. Code Regs. Chapter 8, Article 14 | source_identified_not_fully_parsed | src:usa-sc:regs-ch8-bcc | Full amendment extraction remains open. |
| amend-source:usa-sc:2024-cycle | 2024 IBC, IRC, IFC, IPC, IMC, IFGC; 2023 NEC | LLR 2024 modification index and State Register publication | identified_for_future_cycle | src:usa-sc:llr-building-code-adoption; src:usa-sc:llr-2026-2024-code-final-mods; src:usa-sc:state-register-2025-2024-code-proposed | Needs complete field-level extraction before production use. |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Citation / Section | Summary | Source IDs | Verification Status |
| --- | --- | --- | --- | --- | --- |
| amend:usa-sc:ibc-existing-buildings | Building / Existing Building | Reg. 8-801, IBC Section 101.4.7 Existing Buildings | The South Carolina Existing Building Code applies to repair, alteration, change of occupancy, addition, and relocation of existing buildings. | src:usa-sc:regs-ch8-bcc | verified_text_scope_needs_application_review |
| amend:usa-sc:admin-chapters-local | All building codes | S.C. Code Ann. §6-9-50; Reg. 8-236(E) | Administrative chapters / administrative procedures are not broadly adopted by Council; local jurisdictions may adopt model-code administrative chapters or local administrative procedures by ordinance, subject to any general Council-adopted provisions. | src:usa-sc:code-title-6-ch9; src:usa-sc:regs-ch8-bcc; src:usa-sc:llr-building-codes-in-effect-2022 | verified |
| amend:usa-sc:appendices | All building codes | S.C. Code Ann. §6-9-50; Reg. 8-236(D) | Appendices are not enforceable unless specifically referenced in the code text or included by name / letter designation at adoption. | src:usa-sc:code-title-6-ch9; src:usa-sc:regs-ch8-bcc; src:usa-sc:llr-building-codes-in-effect-2022 | verified |
| amend:usa-sc:local-mod-criteria | Local modifications | S.C. Code Ann. §6-9-105; Reg. 8-245 | Local modifications must be approved by Council, apply only within the requesting jurisdiction, and be justified by physical or climatological conditions. | src:usa-sc:code-title-6-ch9; src:usa-sc:regs-ch8-bcc | verified |
| amend:usa-sc:energy-statutory | Energy | S.C. Code Ann. §6-10-30 | The 2009 IECC is adopted as the Energy Standard and is mandatory for new and renovated buildings and additions. | src:usa-sc:code-title-6-ch10-energy | verified |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-sc"
  model: "statewide_state_adoption_local_enforcement"
  enforcing_entities:
    - "counties for unincorporated areas"
    - "municipalities within municipal limits"
    - "other governmental entities by agreement or contract"
    - "State Fire Marshal for reserved or statewide fire-prevention / life-safety functions"
    - "State Engineer / deputy state fire marshal path for covered state and school-district projects"
  required_officials:
    - "county building official or contracted building-official services for unincorporated areas"
    - "municipal building official or contracted building-official services within municipal limits"
    - "local building inspection department for Energy Standard enforcement"
  state_reserved_activities:
    - "Building Codes Council adoption, modification, and implementation-date authority"
    - "Building Codes Council approval of local modifications"
    - "Building Codes Council modular-building manufacturing-plan / inspection program"
    - "State Fire Marshal fire-prevention, fire-protection, fire investigation, and related licensing / permitting authority"
    - "state-agency and school-district project jurisdiction identified in S.C. Code Ann. §6-9-110"
  source_ids:
    - "src:usa-sc:code-title-6-ch9"
    - "src:usa-sc:regs-ch8-bcc"
    - "src:usa-sc:code-title-6-ch10-energy"
    - "src:usa-sc:code-title-23-ch9-fire-marshal"
    - "src:usa-sc:state-fire-codes-standards"
  verification_status: "partially_verified"
  confidence: 0.84
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-sc"
  model: "state_preapproval_required_for_code_variations"
  applies_to_code_families:
    - "building"
    - "residential"
    - "fire"
    - "plumbing"
    - "mechanical"
    - "fuel_gas"
    - "electrical"
    - "energy_variances"
    - "other Council-adopted building codes"
  approval_required: true
  approving_authority_id: "ahj:usa-sc:bcc"
  filing_required: "submission_to_council_required"
  registry_exists: "partial_public_local_modification_references_identified"
  registry_source_ids:
    - "src:usa-sc:llr-building-code-adoption"
  legal_basis_source_ids:
    - "src:usa-sc:code-title-6-ch9"
    - "src:usa-sc:regs-ch8-bcc"
    - "src:usa-sc:code-title-6-ch10-energy"
  verification_status: "verified_for_general_rule_partial_for_registry"
  confidence: 0.82
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Counties and municipalities are local enforcement bodies, but that does not give them independent authority to create different building-code requirements. For mandatory codes, all building codes used within South Carolina must be adopted by the Building Codes Council and enforced by local jurisdictions. Local jurisdictions are prohibited from writing or publishing other building codes in whole or in part. A local jurisdiction that wants a local building-code modification must first secure local governing-body approval, then obtain Council approval, and the request must be based on local physical or climatological conditions.

Permissive code adoption is a different concept. Municipalities and counties may adopt the latest editions of certain permissive ICC codes by ordinance before enforcement, including property maintenance, existing building, performance, and swimming-pool / spa codes. Those local ordinances and local enforcement practices were not collected in this report.

### 6.4 Known Local Amendment Registries

| Registry / Item ID | Jurisdiction | Code Family | Source / Publication Path | Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| local-mod:usa-sc:greenville-radon | Greenville County | Residential | LLR Building Code Adoption page, Greenville County Local Modification | identified_not_structured | src:usa-sc:llr-building-code-adoption | LLR identifies a local modification authorizing adoption of 2021 IRC Appendix AF for radon control methods after Greenville County adopted a local ordinance. This is treated as an example, not a complete statewide registry. |
| local-registry:usa-sc:bcc-local-modifications | Statewide Council local-modification references | Multiple | LLR Building Codes Council pages and modification materials | partial | src:usa-sc:llr-building-code-adoption | A complete extract of every approved local modification was not completed. |

### 6.5 Municipality-Specific Known Amendments

No municipality-wide amendment set was parsed into structured records beyond the Greenville County radon local-modification reference. Local administrative ordinances, permissive-code adoptions, and local fee / permit procedures remain out of scope for this pass.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: partially_specified_not_implemented

Jurisdiction stack:

```text
Address
  -> State: South Carolina
  -> County
  -> Municipality status: municipal limits or unincorporated county
  -> Special district / fire district, if applicable
  -> Building AHJ:
       - municipality building official inside municipal limits
       - county building official in unincorporated area
       - contracted governmental entity or consultant if a local agreement exists
       - state-agency / school-district path for §6-9-110 projects
  -> Fire AHJ:
       - local fire / building officials for locally administered construction code
       - State Fire Marshal for statewide fire-prevention / life-safety functions and reserved projects
  -> Trade-specific AHJs:
       - local building inspection department for ordinary building-code enforcement
       - specialized state agency only where a separate statute / program applies
  -> Applicable statewide code adoption records
  -> Applicable local administrative ordinances, permissive-code adoptions, and approved local modifications
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | U.S. Census TIGER/Line or state GIS source | none | statewide | annual or source-specific | pending_selection |
| County | U.S. Census TIGER/Line or state GIS source | none | statewide | annual or source-specific | pending_selection |
| Municipality | U.S. Census TIGER/Line incorporated places / state municipal boundary data | none | statewide | annual or source-specific | pending_selection |
| Fire District | state or county GIS fire-district layer | none | partial by county likely | source-specific | unresolved |
| Special District | state or county GIS special-district layers | none | partial by county likely | source-specific | unresolved |
| State / school project flag | project owner / facility dataset, not boundary-only | none | project-specific | source-specific | unresolved |

### 7.3 AHJ Contact Data

| Contact ID | Entity | Scope | Contact / Locator | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| contact:usa-sc:bcc | South Carolina Building Codes Council | State building-code adoption, modification, officer registration, modular program | Contact.bcc@llr.sc.gov; LLR Building Codes Council, PO Box 11329, Columbia, SC 29211-1289 | src:usa-sc:llr-building-code-adoption; src:usa-sc:llr-bcc-home | identified |
| contact:usa-sc:osfm | Office of State Fire Marshal | State fire and life-safety code / standards questions | 803-896-9800; State Fire / Office of State Fire Marshal contact path | src:usa-sc:state-fire-codes-standards | identified |
| contact:usa-sc:local-ahj | County / municipal building departments | Local permits, inspections, local administrative procedures, local permissive-code adoption | Not collected in this report | none | unresolved |

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Source Name | Publisher | URL / Locator | Source Type | Key Use | Last Checked | Caveat |
| --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-sc:code-title-6-ch9 | S.C. Code of Laws, Title 6, Chapter 9, Building Codes | South Carolina Legislative Services Agency / Statehouse | https://www.scstatehouse.gov/code/t06c009.php | statute | Primary building-code authority, local enforcement, adoption procedure, local variations, state/school exceptions, inspection-date rules | 2026-06-26 | Online unannotated statutory text; verify against official published code for legal use. |
| src:usa-sc:code-title-6-ch10-energy | S.C. Code of Laws, Title 6, Chapter 10, Building Energy Efficiency Standards | South Carolina Legislative Services Agency / Statehouse | https://www.scstatehouse.gov/code/t06c010.php | statute | 2009 IECC Energy Standard, statewide compliance, local enforcement, variances | 2026-06-26 | Online unannotated statutory text; verify against official published code for legal use. |
| src:usa-sc:code-title-10-ch5-accessibility | S.C. Code of Laws, Title 10, Chapter 5, Construction of Public Buildings for Access by Persons with Disabilities | South Carolina Legislative Services Agency / Statehouse | https://www.scstatehouse.gov/code/t10c005.php | statute | Accessibility Act and ICC/ANSI A117.1 basis | 2026-06-26 | Source was used for high-level authority only; detailed extraction remains partial. |
| src:usa-sc:code-title-23-ch9-fire-marshal | S.C. Code of Laws, Title 23, Chapter 9, State Fire Marshal | South Carolina Legislative Services Agency / Statehouse | https://www.scstatehouse.gov/code/t23c009.php | statute | State Fire Marshal authority and statewide fire-prevention / life-safety scope | 2026-06-26 | Online unannotated statutory text; verify against official published code for legal use. |
| src:usa-sc:regs-status-page | South Carolina Code of State Regulations status / disclaimer page | South Carolina Legislative Services Agency / Statehouse | https://www.scstatehouse.gov/coderegs/statmast.php | official portal notice | Current-through date and unofficial online regulation caveat | 2026-06-26 | States online unannotated regulations are current through State Register Vol. 50, Issue 3, effective 2026-03-27, but are not the official version. |
| src:usa-sc:regs-ch8-bcc | S.C. Code of Regulations, Chapter 8, Building Codes Council | South Carolina Legislative Services Agency / Statehouse | https://www.scstatehouse.gov/coderegs/Chapter%208.pdf | regulation courtesy copy | Council procedures, building-code adoption / modification procedures, current 2021 amendment articles, accessibility regulations, modular program | 2026-06-26 | Online unannotated regulation copy; not official per Statehouse disclaimer. Treat as authoritative portal copy pending legal verification against official published regulations. |
| src:usa-sc:llr-bcc-home | South Carolina Building Codes Council homepage | South Carolina Department of Labor, Licensing and Regulation | https://llr.sc.gov/bcc/ | agency page | Council function, 2021 code effective date, contact path | 2026-06-26 | Agency web page; may change without formal archival history. |
| src:usa-sc:llr-building-code-adoption | Building Code Adoption | South Carolina Department of Labor, Licensing and Regulation, Building Codes Council | https://llr.sc.gov/bcc/BCAdoption.aspx | agency page | 2021 adoption / effective date, 2024-cycle modification and notice materials, Greenville County local modification reference | 2026-06-26 | Agency web page; use with State Register / statute for legal status. |
| src:usa-sc:llr-building-codes-in-effect-2022 | Building Codes in Effect for South Carolina | South Carolina Department of Labor, Licensing and Regulation, Building Codes Council | https://llr.sc.gov/bcc/PDFfiles/Building%20Codes%20in%20Effect%20for%20South%20Carolina%20Final.pdf | agency PDF | Current 2021/2020 mandatory code list, 2009 IECC, accessibility note, adoption and implementation history | 2026-06-26 | PDF dated 12/2022; current for 2021 package but should be reconciled with 2024-cycle materials and any later LLR updates. |
| src:usa-sc:llr-2021-code-adoption-notice | SC Code Adoption Notice, 2021 South Carolina Building Codes | South Carolina Department of Labor, Licensing and Regulation, Building Codes Council | https://llr.sc.gov/bcc/PDFfiles/SC-Code-Adoption-Notice.pdf | agency PDF notice | 2021 adoption notice, mandatory code list, implementation date, modification validity statement | 2026-06-26 | Agency notice; use with Code and regulations for legal status. |
| src:usa-sc:llr-2024-code-intent | Notice of Intent to Modify and Adopt 2024 Building Codes | South Carolina Department of Labor, Licensing and Regulation / State Register notice | LLR Building Codes Council 2024 Notice of Intent PDF | notice / agency PDF | 2024 I-Code and 2023 NEC modification / adoption process and comment window | 2026-06-26 | Used for process timing and code families; final legal effect requires State Register / codified regulation verification. |
| src:usa-sc:state-register-2025-2024-code-proposed | South Carolina State Register, Vol. 49, Issue 10, October 24, 2025, proposed 2024 Building Codes regulations | South Carolina State Register | https://www.scstatehouse.gov/state_register.php?file=Sr49-10.pdf&first=FILE&pdf=1 | proposed regulation publication | 2024-cycle proposed regulations; 2027-01-01 enforcement-purpose effective date stated in proposal | 2026-06-26 | Proposed regulations are not final by themselves; use only for pending-cycle watch unless final publication is reconciled. |
| src:usa-sc:llr-2026-2024-code-final-mods | 2024 Code Modifications as published in the State Register on 2026-05-22 | South Carolina Department of Labor, Licensing and Regulation / State Register | LLR Building Codes Council 2024 modifications link from Building Code Adoption page | final regulation / agency-linked publication | 2024-cycle final modification publication reference | 2026-06-26 | Full final text extraction remains incomplete; reconcile with codified Chapter 8 before replacing current 2021 package. |
| src:usa-sc:state-fire-codes-standards | Codes and Standards | South Carolina State Fire / Office of State Fire Marshal, LLR | https://statefire.llr.sc.gov/osfm/codes.aspx | agency page | State Fire code / standards page, 2021 effective date reference, state/school project path, OSFM contact | 2026-06-26 | Agency web page; operational enforcement scope needs program-specific confirmation. |
| src:usa-sc:osfm-statutes-regulations | Statutes and Regulations | South Carolina State Fire / Office of State Fire Marshal, LLR | https://statefire.llr.sc.gov/osfm/statutes.aspx | agency page | Fire Marshal statutes and regulations navigation | 2026-06-26 | Supplemental navigation only; statute text is preferred for legal assertions. |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-sc:regs-ch8-bcc | unofficial_online_regulation_copy | Statehouse states that online unannotated regulations are current through State Register Vol. 50, Issue 3, effective 2026-03-27, but are not the official version; only current published volumes and relevant acts / joint resolutions contain the official version. | Use for extraction and QA; legal signoff should compare against official published regulations. |
| src:usa-sc:llr-building-codes-in-effect-2022 | dated_agency_pdf | PDF is dated 12/2022 and accurately supports the 2021/2020 package, but later 2024-cycle activity is underway. | Use for current 2021-package fields until 2024 package is reconciled and becomes enforceable. |
| src:usa-sc:state-register-2025-2024-code-proposed | proposed_regulation | Proposed regulations state planned enforcement-purpose date but are not final by themselves. | Use for pending-change / watch fields, not as sole basis to change current code status. |
| src:usa-sc:llr-2026-2024-code-final-mods | final_publication_not_fully_parsed | LLR identifies State Register publication on 2026-05-22, but this pass did not fully parse every final article. | Use as pending-cycle support; perform full extraction before updating normalized adoption records. |
| src:usa-sc:state-fire-codes-standards | agency_page_operational_scope | Page supports current code references and OSFM contact, but occupancy-specific OSFM jurisdiction requires statute / regulation and AHJ confirmation. | Use for source registry and contact; avoid automated OSFM/local fire AHJ conclusions without further review. |
| src:usa-sc:code-title-10-ch5-accessibility | partial_extraction | Accessibility statute was used at high level; detailed accessibility scope and exceptions were not fully parsed. | Keep accessibility status partially verified. |

### 8.3 Supplemental Sources

None used in this pass. Non-official ICC and third-party code portals were not relied on for current statewide adoption status.

### 8.4 Source Extraction Metadata

| Extraction ID | Source ID | Extracted Fields | Extracted By | Extraction Date | Notes |
| --- | --- | --- | --- | --- | --- |
| extract:usa-sc:001 | src:usa-sc:code-title-6-ch9 | BCC authority, local enforcement, adoption procedure, local variation rule, state/school exception, inspection-date rule | GPT-5.5 Thinking | 2026-06-26 | Primary statute for authority model. |
| extract:usa-sc:002 | src:usa-sc:llr-building-code-adoption | 2021 adoption/effective date, 2024 cycle, local modification example, contact | GPT-5.5 Thinking | 2026-06-26 | Agency current-code page. |
| extract:usa-sc:003 | src:usa-sc:llr-building-codes-in-effect-2022 | Mandatory code list, 2021 adoption and implementation dates, 2009 IECC adoption and implementation, accessibility note, permissive-code note | GPT-5.5 Thinking | 2026-06-26 | Core adoption matrix support. |
| extract:usa-sc:004 | src:usa-sc:regs-ch8-bcc | Reg. 8-236, 8-240, 8-245, accessibility administration, 2021 amendment article references | GPT-5.5 Thinking | 2026-06-26 | Used with caveat because online regulation text is not official version. |
| extract:usa-sc:005 | src:usa-sc:code-title-6-ch10-energy | 2009 IECC Energy Standard, compliance scope, local variance rule, enforcement role | GPT-5.5 Thinking | 2026-06-26 | Energy code confirmed by statute. |
| extract:usa-sc:006 | src:usa-sc:code-title-23-ch9-fire-marshal | State Fire Marshal appointment, statewide jurisdiction, duties, citations | GPT-5.5 Thinking | 2026-06-26 | Fire operational authority remains partially verified by scope. |
| extract:usa-sc:007 | src:usa-sc:state-register-2025-2024-code-proposed | 2024-cycle proposed regulation authority and 2027-01-01 enforcement-purpose date | GPT-5.5 Thinking | 2026-06-26 | Pending-cycle source; not used to replace current 2021 package. |
| extract:usa-sc:008 | src:usa-sc:regs-status-page | Current-through date and unofficial online code caveat | GPT-5.5 Thinking | 2026-06-26 | Source caveat support. |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 1.00 | none | Core authority and current mandatory-code fields are source-backed; unresolved items remain explicit. |
| report | risk.overall_confidence | 0.72 | verified | 1.00 | none | Reflects good support for statewide core fields and incomplete AHJ/amendment/local data. |
| ahj:usa-sc:bcc | authority_name | South Carolina Building Codes Council | verified | 0.95 | src:usa-sc:code-title-6-ch9; src:usa-sc:llr-bcc-home | Primary authority identified. |
| ahj:usa-sc:bcc | role | Statewide adoption / modification / implementation-date authority | verified | 0.90 | src:usa-sc:code-title-6-ch9; src:usa-sc:regs-ch8-bcc | Supported by statute and regulations. |
| adopt:usa-sc:2021-ibc | adoption_date | 2021-10-06 | verified | 0.92 | src:usa-sc:llr-building-code-adoption; src:usa-sc:llr-building-codes-in-effect-2022 | Agency page and PDF align. |
| adopt:usa-sc:2021-ibc | mandatory_date | 2023-01-01 | verified | 0.90 | src:usa-sc:llr-building-code-adoption; src:usa-sc:llr-building-codes-in-effect-2022; src:usa-sc:regs-ch8-bcc | Treated as effective / implementation / local enforcement date. |
| adopt:usa-sc:2020-nec | code_edition | 2020 NEC | verified | 0.90 | src:usa-sc:llr-building-codes-in-effect-2022; src:usa-sc:regs-ch8-bcc | Mandatory package list and regulations support field. |
| adopt:usa-sc:2009-iecc | code_edition | 2009 IECC | verified | 0.90 | src:usa-sc:code-title-6-ch10-energy; src:usa-sc:llr-building-codes-in-effect-2022 | Adopted by statute. |
| local-enforcement:usa-sc | model | statewide_state_adoption_local_enforcement | verified | 0.86 | src:usa-sc:code-title-6-ch9; src:usa-sc:regs-ch8-bcc | County / municipality enforcement confirmed. |
| local-amendment-rule:usa-sc | model | state_preapproval_required_for_code_variations | verified | 0.84 | src:usa-sc:code-title-6-ch9; src:usa-sc:regs-ch8-bcc | Local modifications require Council approval and physical / climatological basis. |
| date-rule:usa-sc:original-permit-date | trigger | original building permit issuance date | verified | 0.88 | src:usa-sc:code-title-6-ch9; src:usa-sc:regs-ch8-bcc | Statute and regulations support date rule. |
| ahj:usa-sc:osfm | operational_fire_scope | statewide fire-prevention / fire-protection authority | partially_verified | 0.76 | src:usa-sc:code-title-23-ch9-fire-marshal; src:usa-sc:state-fire-codes-standards | Needs occupancy/project-specific AHJ rules. |
| ahj:usa-sc:elevator-unresolved | authority_name | unresolved | unresolved | 0.10 | none | Elevator / conveyance authority excluded from this pass. |
| future:usa-sc:2024-cycle | mandatory_date | 2027-01-01 | partially_verified | 0.74 | src:usa-sc:llr-building-code-adoption; src:usa-sc:state-register-2025-2024-code-proposed; src:usa-sc:llr-2026-2024-code-final-mods | Final 2024-cycle adoption records need full parsing before replacing current matrix. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Every `src:usa-sc:...` ID cited outside section 8 is present in the source registry; `none` is used only for unresolved rows. |
| All authority IDs resolve | pass | Authority IDs are defined in authority tables or explicitly marked unresolved. |
| All current code families have adoption records | pass_with_caveat | Core mandatory code families have normalized records. Accessibility has a partial record; elevator/conveyance remains explicit unresolved. |
| Building and operational fire code are separated | pass | BCC-adopted Fire Code and OSFM operational authority are separate rows and authorities. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Dates are separate fields; shared dates are repeated only when supported. |
| Effective dates are valid ISO dates | pass | Populated date fields use ISO format. Nulls are retained for unresolved dates. |
| No impossible date sequences | pass | No populated adoption/effective/operative/mandatory sequence conflicts identified. |
| Transition rules have explicit trigger conditions | pass | Permit-date, implementation-date, change-of-structure, and change-of-use triggers are captured. |
| Permit-date logic is captured where applicable | pass | §6-9-130 and Reg. 8-236 permit-date logic captured. |
| Local enforcement model classified | pass | Statewide state-adoption / local-enforcement model classified. |
| Local amendment rule classified | pass | Council preapproval / physical-climatological model classified. |
| AHJ confirmation metadata present | fail | State-level contacts are identified; local AHJ contact and boundary datasets are not collected. |
| Official-source caveats captured | pass | Online regulation, agency PDF, proposed-regulation, and State Register caveats are recorded. |
| 2024-cycle status conservatively handled | pass | Future cycle is tracked as pending / transition, not silently substituted for current enforceable package. |
| Template-marker scan | pass | The marker scan returned no leftover template tokens or placeholder phrases. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-sc:001 | high | 2024 code cycle | Reconcile final 2024 South Carolina Building Codes / 2023 NEC legal effective date, enforcement date, final codified Chapter 8 text, and any LLR implementation guidance before updating current adoption records. | Parse State Register 2026-05-22 final regulations and current Chapter 8; confirm LLR current-code page after final update. | null | null | open |
| issue:usa-sc:002 | high | amendment extraction | State modifications are source-identified but not fully structured by section, applicability, and effect. | Parse Chapter 8 Articles 8-14 and 2024 final modification index into amendment records. | null | null | open |
| issue:usa-sc:003 | medium | fire AHJ scope | Operational fire-prevention authority, local fire-code enforcement, OSFM reserved projects, and local inspection allocation need occupancy/project-specific mapping. | Extract OSFM regulations and program guidance; validate with OSFM / local AHJ workflows. | null | null | open |
| issue:usa-sc:004 | medium | permissive local codes | Property maintenance, existing building, performance, and swimming pool / spa codes may be adopted locally by ordinance, but local adoptions were not collected. | Build municipality/county ordinance collection and permissive-code registry. | null | null | open |
| issue:usa-sc:005 | medium | local administrative provisions | Local administrative chapters / policies require ordinance adoption and vary by jurisdiction. | Add local ordinance extraction for building department administrative procedures. | null | null | open |
| issue:usa-sc:006 | medium | AHJ boundary and contact data | County/municipal boundaries, fire districts, special districts, and local AHJ contacts are not selected or linked. | Select GIS sources; build contact registry; add state/school project flagging. | null | null | open |
| issue:usa-sc:007 | low | elevator / conveyance authority | Elevator and conveyance code authority and current standards were not resolved. | Research LLR Elevators and Amusement Rides statutes/regulations and determine whether this report should include conveyance code data. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-sc:bcc-adoption-page | src:usa-sc:llr-building-code-adoption | html_diff | monthly | 2024 code cycle status change, current-code list update, new modification links, new local modification reference | 2026-06-26 | active |
| watch:usa-sc:bcc-home | src:usa-sc:llr-bcc-home | html_diff | monthly | change to current effective date or Council function / program statements | 2026-06-26 | active |
| watch:usa-sc:regs-ch8 | src:usa-sc:regs-ch8-bcc | pdf_diff | monthly | Chapter 8 code cycle articles updated or 2024 articles codified | 2026-06-26 | active |
| watch:usa-sc:regs-status | src:usa-sc:regs-status-page | html_diff | monthly | current-through date changes; official caveat changes | 2026-06-26 | active |
| watch:usa-sc:state-register | src:usa-sc:state-register-2025-2024-code-proposed | state_register_search | monthly | final regulations, emergency modifications, notices of drafting, or BCC notices | 2026-06-26 | active |
| watch:usa-sc:title-6-ch9 | src:usa-sc:code-title-6-ch9 | statute_diff | quarterly | amendments to Building Codes Act, adoption procedure, local variation rules, or date rules | 2026-06-26 | active |
| watch:usa-sc:title-6-ch10 | src:usa-sc:code-title-6-ch10-energy | statute_diff | quarterly | IECC edition update or energy-standard enforcement change | 2026-06-26 | active |
| watch:usa-sc:state-fire | src:usa-sc:state-fire-codes-standards | html_diff | monthly | OSFM code/standards updates or contact/path changes | 2026-06-26 | active |
| watch:usa-sc:title-23-ch9 | src:usa-sc:code-title-23-ch9-fire-marshal | statute_diff | quarterly | changes to OSFM statewide authority or fire-code enforcement powers | 2026-06-26 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-24 | Generated baseline draft report | report:usa-sc | none | Codex | Initial South Carolina scaffold had placeholders and unresolved fields. |
| 2026-06-26 | Populated core South Carolina report from official state sources | report:usa-sc; ahj:usa-sc:bcc; ahj:usa-sc:osfm; adopt:usa-sc:2021-ibc; adopt:usa-sc:2021-irc; adopt:usa-sc:2021-ifc; adopt:usa-sc:2021-ipc; adopt:usa-sc:2021-imc; adopt:usa-sc:2021-ifgc; adopt:usa-sc:2020-nec; adopt:usa-sc:2009-iecc; local-enforcement:usa-sc; local-amendment-rule:usa-sc | src:usa-sc:code-title-6-ch9; src:usa-sc:regs-ch8-bcc; src:usa-sc:llr-building-code-adoption; src:usa-sc:llr-building-codes-in-effect-2022; src:usa-sc:code-title-6-ch10-energy; src:usa-sc:code-title-23-ch9-fire-marshal | GPT-5.5 Thinking | Upgraded status to partially_verified; retained explicit unresolved issues for 2024 cycle, AHJ data, local permissive codes, amendments, fire operational scope, and elevator/conveyance authority. |
