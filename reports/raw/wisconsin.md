---
state:
  state_id: "US-WI"
  name: "Wisconsin"
  abbreviation: "WI"
report:
  report_id: "state-report:usa-wi"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-26"
  last_verified: null
  reviewed_by: null
risk:
  overall_confidence: 0.58 # 0.00 - 1.00
  risk_flags:
    - "fire_prevention_code_update_pending"
    - "electrical_and_plumbing_updates_effective_2026_09_01_need_follow_up"
    - "udc_update_rulemaking_pending"
    - "local_amendment_exceptions_need_statute_level_review"
    - "ahj_contact_data_not_populated"
  open_questions_count: 5

---

# State Building Code Authority Report: Wisconsin

## 1. Executive Summary

- **Authority model:** Wisconsin uses a state-administered building and trades code model centered on the Wisconsin Department of Safety and Professional Services (DSPS), Division of Industry Services (DIS). DIS reviews plans for public buildings and places of employment, sets standards for energy efficiency and HVAC, and coordinates with local certified municipalities for certain plan review and inspection services. The state also administers specialized code chapters for the Uniform Dwelling Code, fire prevention, electrical, plumbing, and conveyances. Sources: `src:usa-wi:dsps-commercial-buildings`, `src:usa-wi:dsps-plan-review`, `src:usa-wi:dsps-rules-index`.

- **Statewide code status:** The Wisconsin Commercial Building Code (CBC), Wis. Admin. Code chs. SPS 361 to 366, is verified as a statewide commercial code framework. DSPS states that the Wisconsin CBC adopts the 2021 ICC code set as of 2025-09-01 and that plans submitted on or after 2025-11-01 must comply with the updated CBC. The commercial code families verified in this pass include building, existing building, commercial energy, mechanical, and fuel gas. Sources: `src:usa-wi:dsps-commercial-buildings`, `src:usa-wi:dsps-plan-review`, `src:usa-wi:sps-361-366`, `src:usa-wi:dsps-cbc-code-template`.

- **Residential statewide status:** The Uniform Dwelling Code (UDC), Wis. Admin. Code chs. SPS 320 to 325, is the statewide building code for one- and two-family dwellings built since 1980-06-01. DIS states that the UDC is enforced in all Wisconsin municipalities. A 2025 DSPS statement of scope shows a pending permanent rulemaking to update and clarify the UDC. Sources: `src:usa-wi:dsps-udc-page`, `src:usa-wi:sps-320-325`, `src:usa-wi:udc-scope-2025`.

- **Local enforcement model:** Enforcement is hybrid. DIS performs state plan review and inspection activities where reserved to the state or not delegated; local certified municipalities or delegated agents perform certain commercial plan review and inspection services; municipalities enforce the UDC statewide; and fire departments or municipalities participating in fire dues have operational fire-prevention duties under the SPS 314 framework. Sources: `src:usa-wi:dsps-commercial-buildings`, `src:usa-wi:dsps-plan-review`, `src:usa-wi:dsps-udc-page`, `src:usa-wi:fire-scope-2025`.

- **Local amendment posture:** Wisconsin generally has a statewide-uniform posture for commercial/public-building construction standards. Wis. Stat. § 101.02(7r) restricts counties, cities, villages, and towns from enacting or enforcing local minimum standards for constructing, altering, or adding to public buildings or places of employment unless the ordinance strictly conforms to the applicable DSPS rules, subject to statutory exceptions. The plumbing code is also described in rulemaking text as uniform statewide. A municipality-specific amendment registry was not located in this pass. Sources: `src:usa-wi:stat-101-02-local-standards`, `src:usa-wi:plumbing-cr26-019`.

- **Known transition periods or pending changes:** CBC 2021 ICC adoption is current, with mandatory compliance tied to plans submitted on or after 2025-11-01. DSPS lists SPS 316 Electrical Code and SPS 381 to 387 Plumbing Code Review as rules effective 2026-09-01. DSPS lists SPS 314 Fire Prevention Code Review and SPS 320 to 325 UDC Update as drafting rules. Sources: `src:usa-wi:dsps-commercial-buildings`, `src:usa-wi:dsps-plan-review`, `src:usa-wi:dsps-pending-rules`, `src:usa-wi:fire-scope-2025`, `src:usa-wi:udc-scope-2025`.

- **Production readiness:** partially_verified_for_core_statewide_authority; not_ready_for_full_ahj_resolution.

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: DSPS/DIS is the primary verified statewide authority for commercial building
    plan review, UDC administration, and state construction/trades code programs.
  confidence: 0.8
  source_ids:
  - src:usa-wi:dsps-commercial-buildings
  - src:usa-wi:dsps-plan-review
  - src:usa-wi:dsps-rules-index
- topic: Primary commercial building code edition
  finding: Wisconsin CBC adopts the 2021 ICC code set as of 2025-09-01, with plans
    submitted on or after 2025-11-01 required to comply.
  confidence: 0.82
  source_ids:
  - src:usa-wi:dsps-commercial-buildings
  - src:usa-wi:dsps-plan-review
  - src:usa-wi:sps-361-366
  - src:usa-wi:dsps-cbc-code-template
- topic: Residential building code
  finding: UDC is statewide for one- and two-family dwellings built since 1980-06-01
    and is enforced in all Wisconsin municipalities.
  confidence: 0.82
  source_ids:
  - src:usa-wi:dsps-udc-page
  - src:usa-wi:sps-320-325
- topic: Fire code authority
  finding: SPS 314 is the fire-prevention chapter; current rules reference 2012 NFPA
    1 and DSPS is considering a 2024 NFPA 1 update.
  confidence: 0.7
  source_ids:
  - src:usa-wi:sps-314
  - src:usa-wi:fire-scope-2025
  - src:usa-wi:dsps-pending-rules
- topic: Electrical code
  finding: SPS 316 is the electrical code chapter. DSPS rulemaking materials state
    the current chapter adopts 2017 NEC and the pending/approved update adopts 2023
    NEC effective 2026-09-01.
  confidence: 0.7
  source_ids:
  - src:usa-wi:sps-316
  - src:usa-wi:electrical-cr26-016
  - src:usa-wi:dsps-pending-rules
- topic: Plumbing code
  finding: SPS 381 to 387 are the Wisconsin Plumbing Code chapters; DSPS pending-rules
    page lists Plumbing Code Review effective 2026-09-01.
  confidence: 0.72
  source_ids:
  - src:usa-wi:sps-381-387
  - src:usa-wi:plumbing-cr26-019
  - src:usa-wi:dsps-pending-rules
- topic: Local amendments
  finding: Commercial/public-building local minimum construction standards generally
    must strictly conform to state rules, with statutory exceptions; registry and
    exception inventory remain unresolved.
  confidence: 0.6
  source_ids:
  - src:usa-wi:stat-101-02-local-standards
  - src:usa-wi:court-abc-madison-2023
- topic: Effective / operative date rule
  finding: CBC uses both an effective date and a later plan-submittal compliance date;
    electrical and plumbing updates have a future effective date of 2026-09-01.
  confidence: 0.74
  source_ids:
  - src:usa-wi:dsps-commercial-buildings
  - src:usa-wi:dsps-plan-review
  - src:usa-wi:dsps-pending-rules
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | `ahj:usa-wi:dsps-dis` |
| Authority name | Wisconsin Department of Safety and Professional Services, Division of Industry Services |
| Authority type | state agency / state code-administration division |
| Legal basis | Wis. Stat. ch. 101 and Wis. Admin. Code SPS construction/trades chapters; exact subsection mapping captured for high-value topics below and remains incomplete for all programs. |
| Role | Administers statewide construction/trades code programs; reviews plans for public buildings and places of employment; sets commercial energy and HVAC standards; administers UDC consultation, credentials, and materials evaluation; coordinates with local certified municipalities and delegated agents. |
| Enforcement model | state-administered with local certified municipality/delegated-agent participation and municipal UDC enforcement |
| Source IDs | `src:usa-wi:dsps-commercial-buildings`; `src:usa-wi:dsps-plan-review`; `src:usa-wi:dsps-udc-page`; `src:usa-wi:dsps-rules-index` |
| Verification status | verified_core |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | `ahj:usa-wi:dsps-dis` | DSPS / DIS | Administers Wisconsin CBC for public buildings and places of employment; plan review for covered commercial buildings. | Wis. Stat. ch. 101; Wis. Admin. Code SPS 361 and SPS 362 | `src:usa-wi:dsps-commercial-buildings`; `src:usa-wi:dsps-plan-review`; `src:usa-wi:sps-361-366` | verified_core |
| Residential | `ahj:usa-wi:dsps-dis` | DSPS / DIS | Administers UDC program for one- and two-family dwellings; municipalities enforce UDC statewide. | Wis. Admin. Code SPS 320 to 325; Wis. Stat. §§ 101.63 and 101.65 | `src:usa-wi:dsps-udc-page`; `src:usa-wi:sps-320-325`; `src:usa-wi:stat-101-63`; `src:usa-wi:stat-101-65` | verified_core |
| Existing Building / Rehabilitation | `ahj:usa-wi:dsps-dis` | DSPS / DIS | Administers existing-building provisions within the commercial code. | Wis. Admin. Code SPS 366 | `src:usa-wi:dsps-rules-index`; `src:usa-wi:sps-361-366`; `src:usa-wi:dsps-cbc-code-template` | verified_core |
| Mechanical | `ahj:usa-wi:dsps-dis` | DSPS / DIS | Administers commercial HVAC/mechanical provisions. | Wis. Admin. Code SPS 364 | `src:usa-wi:dsps-commercial-buildings`; `src:usa-wi:dsps-rules-index`; `src:usa-wi:sps-361-366` | verified_core |
| Plumbing | `ahj:usa-wi:dsps-dis-plumbing` | DSPS / DIS Plumbing Program | Administers Wisconsin Plumbing Code and related plan review. | Wis. Stat. ch. 145; Wis. Admin. Code SPS 381 to 387 | `src:usa-wi:dsps-rules-index`; `src:usa-wi:sps-381-387`; `src:usa-wi:plumbing-cr26-019`; `src:usa-wi:dsps-plan-review` | partially_verified |
| Fuel Gas | `ahj:usa-wi:dsps-dis` | DSPS / DIS | Administers commercial fuel gas appliance provisions. | Wis. Admin. Code SPS 365 | `src:usa-wi:dsps-rules-index`; `src:usa-wi:sps-361-366`; `src:usa-wi:dsps-cbc-code-template` | verified_core |
| Electrical | `ahj:usa-wi:dsps-dis-electrical` | DSPS / DIS Electrical Program | Administers Wisconsin Electrical Code; future SPS 316 rule effective 2026-09-01. | Wis. Stat. § 101.82; Wis. Admin. Code SPS 316 | `src:usa-wi:dsps-rules-index`; `src:usa-wi:sps-316`; `src:usa-wi:stat-101-82`; `src:usa-wi:electrical-cr26-016`; `src:usa-wi:dsps-pending-rules` | partially_verified |
| Energy | `ahj:usa-wi:dsps-dis` | DSPS / DIS | Sets standards for energy efficiency in commercial buildings; administers SPS 363 and UDC SPS 322 energy provisions. | Wis. Admin. Code SPS 363 and SPS 322 | `src:usa-wi:dsps-commercial-buildings`; `src:usa-wi:dsps-rules-index`; `src:usa-wi:sps-361-366`; `src:usa-wi:sps-320-325` | verified_core |
| Fire - construction references | `ahj:usa-wi:dsps-dis` | DSPS / DIS | Reviews fire alarm, fire detection, and fire suppression systems as part of covered public-building/place-of-employment plan review; construction provisions are embedded in CBC. | Wis. Stat. § 101.12; Wis. Admin. Code SPS 361 to 366 | `src:usa-wi:dsps-plan-review`; `src:usa-wi:sps-361-366` | partially_verified |
| Fire - operational / prevention code | `ahj:usa-wi:dsps-dis-fire-prevention` | DSPS / DIS Fire Prevention Program; local fire departments/municipalities for fire-dues duties | Administers SPS 314 fire-prevention framework; local fire departments and municipalities are affected/enforcement participants under the fire-dues framework. | Wis. Stat. §§ 101.02 and 101.14; Wis. Admin. Code SPS 314 | `src:usa-wi:dsps-rules-index`; `src:usa-wi:sps-314`; `src:usa-wi:fire-scope-2025`; `src:usa-wi:stat-101-14` | partially_verified |
| Accessibility | `ahj:usa-wi:dsps-dis` | DSPS / DIS | Accessibility provisions are administered through the commercial building code; specific standard-level parsing was limited to DSPS code-template identification. | Wis. Admin. Code SPS 361 to 366; model-code accessibility provisions | `src:usa-wi:sps-361-366`; `src:usa-wi:dsps-cbc-code-template` | partially_verified |
| Elevator / Conveyance | `ahj:usa-wi:dsps-dis-conveyance` | DSPS / DIS Conveyance Program | Administers elevators, escalators, and lift devices; plan review required for elevators, escalators, lifts, and power dumbwaiters in covered projects. | Wis. Admin. Code SPS 318; Wis. Stat. § 101.12 for relevant plan review | `src:usa-wi:dsps-rules-index`; `src:usa-wi:sps-318`; `src:usa-wi:dsps-plan-review` | partially_verified |

### 2.3 Authority Hierarchy Notes

Wisconsin's verified hierarchy is state-led for code adoption and administration, with local enforcement/delegation elements. DSPS/DIS is the core state authority. For commercial/public-building plan review, DIS performs reviews and coordinates with local certified municipalities that can provide plan review and inspection for certain buildings. For one- and two-family dwellings, the UDC is statewide and enforced in all municipalities; Wis. Stat. § 101.65 and SPS 320.065 support municipal jurisdiction over new one- and two-family dwelling construction and inspection. Fire-prevention enforcement is more distributed because SPS 314 also establishes duties and requirements for fire departments and municipalities receiving fire dues; a full fire-AHJ hierarchy and contact list is not populated here.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| `edge:usa-wi:001` | `ahj:usa-wi:dsps-dis` | administers | statewide CBC chapters SPS 361 to 366 for public buildings and places of employment | `src:usa-wi:dsps-commercial-buildings`; `src:usa-wi:sps-361-366` | verified_core |
| `edge:usa-wi:002` | `ahj:usa-wi:dsps-dis` | coordinates_with | local certified municipalities for certain commercial plan review and inspection services | `src:usa-wi:dsps-commercial-buildings`; `src:usa-wi:dsps-plan-review` | verified_core |
| `edge:usa-wi:003` | `ahj:usa-wi:dsps-dis` | requires_plan_review_for | covered public buildings, places of employment, schools, high-hazard occupancies, 3+ unit residential buildings, and listed systems | `src:usa-wi:dsps-plan-review`; `src:usa-wi:stat-101-12` | verified_core |
| `edge:usa-wi:004` | `ahj:usa-wi:dsps-dis` | administers | statewide UDC chapters SPS 320 to 325 | `src:usa-wi:dsps-udc-page`; `src:usa-wi:sps-320-325` | verified_core |
| `edge:usa-wi:005` | `local_municipalities` | enforce | UDC in all Wisconsin municipalities | `src:usa-wi:dsps-udc-page`; `src:usa-wi:stat-101-65` | verified_core |
| `edge:usa-wi:006` | `ahj:usa-wi:dsps-dis-fire-prevention` | administers | SPS 314 operational fire-prevention code | `src:usa-wi:sps-314`; `src:usa-wi:fire-scope-2025`; `src:usa-wi:stat-101-14` | partially_verified |
| `edge:usa-wi:007` | `ahj:usa-wi:dsps-dis` | preempts_or_limits | local commercial/public-building minimum construction standards that do not strictly conform, subject to statutory exceptions | `src:usa-wi:stat-101-02-local-standards`; `src:usa-wi:court-abc-madison-2023` | partially_verified |
| `edge:usa-wi:008` | `ahj:usa-wi:dsps-dis-plumbing` | administers_uniform_code | statewide Wisconsin Plumbing Code SPS 381 to 387 | `src:usa-wi:sps-381-387`; `src:usa-wi:plumbing-cr26-019` | partially_verified |
| `edge:usa-wi:009` | `ahj:usa-wi:dsps-dis-electrical` | administers | Wisconsin Electrical Code SPS 316 | `src:usa-wi:sps-316`; `src:usa-wi:electrical-cr26-016`; `src:usa-wi:stat-101-82` | partially_verified |
| `edge:usa-wi:010` | `ahj:usa-wi:dsps-dis-conveyance` | administers | elevators, escalators, and lift devices under SPS 318 | `src:usa-wi:sps-318`; `src:usa-wi:dsps-plan-review` | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | Wisconsin Commercial Building Code, SPS 361 to 366 / SPS 362 Buildings and Structures | International Building Code | 2021 | current statewide commercial/public-building code | null | 2025-09-01 | 2025-09-01 | 2025-11-01 | DSPS current pages state the 2021 ICC update is effective 2025-09-01 and plans submitted on or after 2025-11-01 must comply; DSPS earlier releases used shorter grace-period dates and are treated as superseded for the mandatory plan-submittal date. | `src:usa-wi:dsps-commercial-buildings`; `src:usa-wi:dsps-plan-review`; `src:usa-wi:sps-361-366`; `src:usa-wi:dsps-cbc-code-template`; `src:usa-wi:cr23-007`; `src:usa-wi:dsps-cbc-2025-release`; `src:usa-wi:dsps-cbc-extension-2025` |
| Residential | Wisconsin Uniform Dwelling Code, SPS 320 to 325 | Wisconsin state-specific UDC with national standards and codes referenced; wholesale IRC adoption not verified | current SPS 320 to 325; last comprehensive update effective 2016-01-01 | current statewide; update rulemaking pending | null | 2016-01-01 | 2016-01-01 | 1980-06-01 | UDC applies statewide to one- and two-family dwellings built since 1980-06-01; pending rulemaking to update and clarify UDC. | `src:usa-wi:dsps-udc-page`; `src:usa-wi:sps-320-325`; `src:usa-wi:udc-scope-2025` |
| Existing Building / Rehabilitation | Wisconsin Commercial Building Code, SPS 366 Existing Buildings | International Existing Building Code | 2021 | current statewide commercial/public-building existing-building provisions | null | 2025-09-01 | 2025-09-01 | 2025-11-01 | Same CBC transition rule as commercial building plans. | `src:usa-wi:dsps-rules-index`; `src:usa-wi:sps-361-366`; `src:usa-wi:dsps-cbc-code-template`; `src:usa-wi:dsps-plan-review` |
| Mechanical | Wisconsin Commercial Building Code, SPS 364 HVAC | International Mechanical Code | 2021 | current statewide commercial mechanical provisions | null | 2025-09-01 | 2025-09-01 | 2025-11-01 | Same CBC transition rule; supplemental HVAC submissions should align with the code edition for the parent commercial building plan. | `src:usa-wi:dsps-rules-index`; `src:usa-wi:sps-361-366`; `src:usa-wi:dsps-cbc-code-template`; `src:usa-wi:dsps-cbc-2025-release` |
| Plumbing | Wisconsin Plumbing Code, SPS 381 to 387 | Wisconsin state plumbing code; model standard basis not parsed in this pass | current SPS 381 to 387; 2026 update effective 2026-09-01 | current with future update effective 2026-09-01 | null | null | null | 2026-09-01 for 2026 update | DSPS pending-rules page lists SPS 381 to 387 Plumbing Code Review as rule effective 2026-09-01; current chapter set remains SPS 381 to 387. | `src:usa-wi:dsps-rules-index`; `src:usa-wi:sps-381-387`; `src:usa-wi:plumbing-cr26-019`; `src:usa-wi:dsps-pending-rules` |
| Fuel Gas | Wisconsin Commercial Building Code, SPS 365 Fuel Gas Appliances | International Fuel Gas Code | 2021 | current statewide commercial fuel-gas provisions | null | 2025-09-01 | 2025-09-01 | 2025-11-01 | Same CBC transition rule; DSPS has issued program guidance on an SPS 365 fuel-gas venting amendment. | `src:usa-wi:dsps-rules-index`; `src:usa-wi:sps-361-366`; `src:usa-wi:dsps-cbc-code-template`; `src:usa-wi:dsps-commercial-buildings` |
| Electrical | Wisconsin Electrical Code, SPS 316 | NFPA 70, National Electrical Code | 2017 current; 2023 update effective 2026-09-01 | current with future update effective 2026-09-01 | null | null | null | 2026-09-01 for 2023 NEC update | DSPS rulemaking materials state current SPS 316 adopts 2017 NEC; pending-rules page lists SPS 316 Electrical Code as rule effective 2026-09-01. Commercial electrical plan review is voluntary under SPS 316. | `src:usa-wi:dsps-rules-index`; `src:usa-wi:sps-316`; `src:usa-wi:electrical-cr26-016`; `src:usa-wi:dsps-pending-rules`; `src:usa-wi:dsps-plan-review` |
| Energy | Commercial Energy Conservation, SPS 363; UDC Energy Conservation, SPS 322 | International Energy Conservation Code for commercial CBC; UDC energy provisions for one- and two-family dwellings | 2021 IECC for commercial; UDC edition not model-code-normalized | current statewide | null for commercial CBC | 2025-09-01 for commercial CBC | 2025-09-01 for commercial CBC | 2025-11-01 for commercial CBC plans | Commercial energy follows CBC transition; residential energy remains in UDC SPS 322 and pending UDC rulemaking. | `src:usa-wi:dsps-commercial-buildings`; `src:usa-wi:sps-361-366`; `src:usa-wi:dsps-cbc-code-template`; `src:usa-wi:sps-320-325` |
| Fire - construction references | CBC fire/life-safety construction provisions and plan review for fire alarm/detection/suppression systems | IBC and referenced standards within CBC; detailed fire-standard mapping unresolved | 2021 IBC framework | current construction-code references partially verified | null | 2025-09-01 | 2025-09-01 | 2025-11-01 | Construction fire/life-safety items follow CBC; plan review is required for listed fire alarm, fire detection, and fire suppression systems in covered projects. | `src:usa-wi:dsps-plan-review`; `src:usa-wi:sps-361-366`; `src:usa-wi:dsps-cbc-code-template` |
| Fire - operational / prevention code | Fire Prevention Code, SPS 314 | NFPA 1 Fire Code | 2012 current; 2024 NFPA 1 under consideration | current with update rulemaking pending | null | 2014-09-01 | 2014-09-01 | 2014-09-01 | DSPS 2025 scope statement says current SPS 314 rules adopt/reference 2012 NFPA 1 and the update project will consider 2024 NFPA 1. | `src:usa-wi:dsps-rules-index`; `src:usa-wi:sps-314`; `src:usa-wi:fire-scope-2025`; `src:usa-wi:dsps-pending-rules` |
| Accessibility | Accessibility provisions within Wisconsin CBC | IBC accessibility provisions and ICC/ANSI standard identified by DSPS code template | 2021 IBC / 2017 ICC A117.1 in DSPS template | partially verified | null | 2025-09-01 | 2025-09-01 | 2025-11-01 | Accessibility standard was identified in DSPS commercial code template, but section-level amendments were not parsed. | `src:usa-wi:sps-361-366`; `src:usa-wi:dsps-cbc-code-template` |
| Elevator / Conveyance | Elevators, Escalators and Lift Devices, SPS 318 | Conveyance standards in SPS 318; base standard details unresolved | current SPS 318 | current chapter verified; base standard not normalized | null | null | null | null | Plan review is required for elevators, escalators, lifts, and power dumbwaiters in covered public-building/place-of-employment projects; detailed conveyance standard edition not parsed. | `src:usa-wi:dsps-rules-index`; `src:usa-wi:sps-318`; `src:usa-wi:dsps-plan-review` |

### 3.2 Adoption Records

| Adoption Record ID | Code Family | State Code / Instrument | Base Code / Standard | Adoption Date | Effective Date | Operative Date | Mandatory Date | Applicability / Trigger | Source IDs | Verification Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| `adoption-rec:usa-wi:cbc-2021` | Building / Existing / Mechanical / Fuel Gas / Commercial Energy | Wis. Admin. Code SPS 361 to 366; CR 23-007 | 2021 IBC, 2021 IEBC, 2021 IECC, 2021 IMC, 2021 IFGC; 2017 ICC A117.1 noted in DSPS template | null | 2025-09-01 | 2025-09-01 | 2025-11-01 | Commercial/public-building plans submitted on or after 2025-11-01 must comply with updated CBC. | `src:usa-wi:dsps-commercial-buildings`; `src:usa-wi:dsps-plan-review`; `src:usa-wi:sps-361-366`; `src:usa-wi:dsps-cbc-code-template`; `src:usa-wi:cr23-007` | verified_core |
| `adoption-rec:usa-wi:udc-current` | Residential | Wis. Admin. Code SPS 320 to 325 | Wisconsin UDC; national standards/codes referenced; exact model-code crosswalk unresolved | null | 2016-01-01 | 2016-01-01 | 1980-06-01 | Applies statewide to one- and two-family dwellings built since 1980-06-01; enforced in all municipalities. | `src:usa-wi:dsps-udc-page`; `src:usa-wi:sps-320-325`; `src:usa-wi:udc-scope-2025` | verified_core |
| `adoption-rec:usa-wi:fire-nfpa1-2012` | Fire - operational / prevention code | Wis. Admin. Code SPS 314 | NFPA 1 Fire Code | null | 2014-09-01 | 2014-09-01 | 2014-09-01 | Applies primarily to use, operation, and maintenance of public buildings and places of employment; fire department/municipality fire-dues duties also included. | `src:usa-wi:sps-314`; `src:usa-wi:fire-scope-2025`; `src:usa-wi:stat-101-14` | partially_verified |
| `adoption-rec:usa-wi:electrical-2017-nec` | Electrical | Wis. Admin. Code SPS 316 | 2017 NEC | null | null | null | null | Current SPS 316 baseline before 2026 update; commercial electrical plan review generally voluntary under SPS 316. | `src:usa-wi:sps-316`; `src:usa-wi:electrical-cr26-016`; `src:usa-wi:dsps-plan-review` | partially_verified |
| `adoption-rec:usa-wi:electrical-2023-nec-future` | Electrical | CR 26-016 / SPS 316 update | 2023 NEC | null | 2026-09-01 | 2026-09-01 | 2026-09-01 | DSPS pending-rules page lists SPS 316 Electrical Code as rule effective 2026-09-01. | `src:usa-wi:electrical-cr26-016`; `src:usa-wi:dsps-pending-rules` | partially_verified |
| `adoption-rec:usa-wi:plumbing-2026-update` | Plumbing | CR 26-019 / SPS 381 to 387 update | Wisconsin Plumbing Code; model-standard crosswalk unresolved | null | 2026-09-01 | 2026-09-01 | 2026-09-01 | DSPS pending-rules page lists SPS 381 to 387 Plumbing Code Review as rule effective 2026-09-01. | `src:usa-wi:plumbing-cr26-019`; `src:usa-wi:dsps-pending-rules`; `src:usa-wi:sps-381-387` | partially_verified |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

Wisconsin distinguishes at least three date concepts in the current evidence set: (1) the effective date of the commercial code update, (2) the plan-submittal compliance date for CBC projects, and (3) future effective dates for separate electrical and plumbing rule packages. The current DSPS commercial and plan-review pages supersede earlier press-release grace-period language by stating that the CBC adopts the 2021 ICC code set as of 2025-09-01 while plans submitted on or after 2025-11-01 must comply. DSPS pending-rules status lists SPS 316 Electrical Code and SPS 381 to 387 Plumbing Code Review as rules effective 2026-09-01. UDC applicability is not a simple edition date: DSPS states the UDC is statewide for one- and two-family dwellings built since 1980-06-01, and a 2025 scope statement identifies the most recent comprehensive UDC update as effective 2016-01-01.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `date-rule:usa-wi:cbc-2021-effective` | Wisconsin CBC, SPS 361 to 366 | effective_date | 2025-09-01 | State code update published/effective for 2021 ICC adoption. | yes, under compliance grace period for plan submissions | `src:usa-wi:dsps-commercial-buildings`; `src:usa-wi:dsps-plan-review`; `src:usa-wi:dsps-cbc-2025-release`; `src:usa-wi:cr23-007` | verified_core |
| `date-rule:usa-wi:cbc-2021-plan-compliance` | Wisconsin CBC commercial building plans | mandatory_plan_submittal_date | 2025-11-01 | Plans submitted on or after this date must comply with updated CBC. | yes, for plans submitted before this date, subject to DSPS review posture and project facts | `src:usa-wi:dsps-commercial-buildings`; `src:usa-wi:dsps-plan-review`; `src:usa-wi:dsps-cbc-extension-2025` | verified_core |
| `date-rule:usa-wi:cbc-supplemental-alignment` | CBC supplemental submissions, including fire suppression/alarm, HVAC, boilers, elevators, and refrigeration | parent_plan_alignment_rule | project-specific | Supplemental submissions should align with the code edition used for the approved commercial building plan. | yes, if parent plan was validly reviewed under prior code | `src:usa-wi:dsps-cbc-2025-release`; `src:usa-wi:dsps-plan-review` | partially_verified |
| `date-rule:usa-wi:udc-initial-scope` | UDC one- and two-family dwellings | applicability_start_date | 1980-06-01 | One- and two-family dwellings built on or after this date are within statewide UDC scope. | not applicable to pre-scope dwellings without project-specific analysis | `src:usa-wi:dsps-udc-page` | verified_core |
| `date-rule:usa-wi:udc-current-comprehensive-update` | UDC SPS 320 to 325 | prior_comprehensive_update_effective_date | 2016-01-01 | Last comprehensive update identified by DSPS 2025 statement of scope. | unresolved | `src:usa-wi:udc-scope-2025` | partially_verified |
| `date-rule:usa-wi:electrical-2026-update` | SPS 316 Electrical Code | future_rule_effective_date | 2026-09-01 | DSPS pending-rules page lists SPS 316 as rule effective on this date. | unresolved for permit/application transition details | `src:usa-wi:dsps-pending-rules`; `src:usa-wi:electrical-cr26-016` | partially_verified |
| `date-rule:usa-wi:plumbing-2026-update` | SPS 381 to 387 Plumbing Code Review | future_rule_effective_date | 2026-09-01 | DSPS pending-rules page lists SPS 381 to 387 as rule effective on this date. | unresolved for permit/application transition details | `src:usa-wi:dsps-pending-rules`; `src:usa-wi:plumbing-cr26-019` | partially_verified |
| `date-rule:usa-wi:fire-prevention-update` | SPS 314 Fire Prevention Code | pending_rulemaking_status | null | DSPS pending-rules page lists SPS 314 Fire Prevention Code Review as drafting rule; no future effective date verified. | unresolved | `src:usa-wi:dsps-pending-rules`; `src:usa-wi:fire-scope-2025` | unresolved |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Electrical | SPS 316 update adopting 2023 NEC | 2026-03-16 | null | 2026-09-01 | 2026-09-01 | 2026-09-01 | active_watch | `src:usa-wi:electrical-cr26-016`; `src:usa-wi:dsps-pending-rules` | DSPS materials state current SPS 316 adopts 2017 NEC and update adopts 2023 NEC. Confirm final text after effective date. |
| Plumbing | SPS 381 to 387 Plumbing Code Review | 2026-03-16 | null | 2026-09-01 | 2026-09-01 | 2026-09-01 | active_watch | `src:usa-wi:plumbing-cr26-019`; `src:usa-wi:dsps-pending-rules` | DSPS pending-rules page identifies effective date. Confirm final code text after effective date. |
| Fire - operational / prevention code | SPS 314 Fire Prevention Code Review; potential 2024 NFPA 1 consideration | 2025-11-04 | null | null | null | null | active_watch | `src:usa-wi:fire-scope-2025`; `src:usa-wi:dsps-pending-rules` | Status is drafting rule. Existing rules reference 2012 NFPA 1. |
| Residential | SPS 302, 305, and 320 to 325 update to UDC | 2025-03-06 | null | null | null | null | active_watch | `src:usa-wi:udc-scope-2025`; `src:usa-wi:dsps-pending-rules` | Status is drafting rule. |
| Conveyance | SPS 302, 305, 316, 318, 321, 362, 366, 381, and 382 conveyance safety code comprehensive review | null | null | null | null | null | active_watch | `src:usa-wi:dsps-pending-rules`; `src:usa-wi:sps-318` | Listed as drafting rule; not parsed beyond pending-rules status. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| `applicability-rule:usa-wi:commercial-plan-review-volume` | Commercial Building | Commercial, industrial, and other public buildings 25,000 cubic feet or larger; any size school, high-hazard occupancy, and 3+ unit residential building | Plan review threshold and occupancy/use category | DSPS plan review page states plan review is required for these categories and for listed systems under Wis. Stat. § 101.12 and plumbing law references. | `src:usa-wi:dsps-plan-review`; `src:usa-wi:stat-101-12` | verified_core |
| `applicability-rule:usa-wi:commercial-system-review` | Commercial Building / Fire / Plumbing / Conveyance | HVAC, fire alarm, fire detection, fire suppression, industrial exhaust, elevators/escalators/lifts, bleachers, amusement rides, plumbing, POWTS | System plan review | Certain system plan reviews are required even when the building project context differs; commercial electrical plan reviews under SPS 316 are identified by DSPS as voluntary. | `src:usa-wi:dsps-plan-review`; `src:usa-wi:sps-316` | partially_verified |
| `applicability-rule:usa-wi:udc-statewide` | Residential / UDC | One- and two-family dwellings built since 1980-06-01 | Building date and dwelling type | UDC is statewide and enforced in all municipalities. | `src:usa-wi:dsps-udc-page`; `src:usa-wi:sps-320-325` | verified_core |
| `applicability-rule:usa-wi:fire-prevention-scope` | Fire - operational / prevention | Use, operation, and maintenance of public buildings and places of employment; fire department/municipality fire-dues duties | Operational fire prevention | Current SPS 314 framework primarily concerns national fire-prevention standards for use, operation, and maintenance of public buildings and places of employment and duties of fire departments/municipalities receiving fire dues. | `src:usa-wi:fire-scope-2025`; `src:usa-wi:sps-314` | partially_verified |
| `applicability-rule:usa-wi:plumbing-uniform-statewide` | Plumbing | New plumbing installations and plumbing code scope statewide | Plumbing installation or design scope | Rulemaking text describes SPS 381 to 387 as applying uniformly to design, construction, installation, supervision, maintenance, and inspection of plumbing. | `src:usa-wi:plumbing-cr26-019`; `src:usa-wi:sps-381-387` | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** chaptered_state_amendment_model

Wisconsin publishes state-specific chapters that administer and amend model codes rather than relying on a bare model-code adoption. The verified commercial structure is SPS 361 Administration and Enforcement, SPS 362 Buildings and Structures, SPS 363 Energy Conservation, SPS 364 HVAC, SPS 365 Fuel Gas Appliances, and SPS 366 Existing Buildings. Fire prevention is SPS 314 and modifies/references NFPA 1. Electrical is SPS 316 and modifies/references the NEC. Plumbing is the state plumbing code in SPS 381 to 387. The UDC is a state-specific code in SPS 320 to 325.

**Where amendments are published:** Wisconsin Administrative Code on the Wisconsin Legislature / Legislative Reference Bureau site; DSPS program pages, code inserts, code templates, guidance documents, rulemaking pages, and newsletters are supporting publication paths and caveat sources.

**Amendment parsing status:** high_level_structure_parsed; section_level_amendments_not_parsed

### 5.2 State Amendment Sources

| Source ID | Amendment Scope | Publication Path | Status | Notes |
| --- | --- | --- | --- | --- |
| `src:usa-wi:sps-361-366` | Commercial Building Code chapters, including administration, IBC, IECC, IMC, IFGC, and IEBC amendments | Wisconsin Administrative Code | partially_parsed | Model-code editions verified; individual amendments not extracted section-by-section. |
| `src:usa-wi:dsps-cbc-code-template` | DSPS commercial plan code-template summary | DSPS PDF/template | parsed_summary_only | Identifies applicable 2025 Wisconsin SPS 361 to 366, 2021 IBC/IEBC/IECC/IMC/IFGC, and 2017 ICC A117.1; template is not rule text. |
| `src:usa-wi:sps-320-325` | UDC chapters SPS 320 to 325 | Wisconsin Administrative Code | partially_parsed | Current UDC chapter list verified; section-level standards not extracted. |
| `src:usa-wi:sps-314` | Fire Prevention Code | Wisconsin Administrative Code | partially_parsed | Current NFPA 1 edition supported through DSPS 2025 scope statement. |
| `src:usa-wi:sps-316` | Electrical Code | Wisconsin Administrative Code | partially_parsed | Current NEC edition supported through DSPS rulemaking text; pending update effective 2026-09-01. |
| `src:usa-wi:sps-381-387` | Wisconsin Plumbing Code | Wisconsin Administrative Code | partially_parsed | Uniform statewide scope supported through CR 26-019 rulemaking text. |
| `src:usa-wi:sps-318` | Elevators, escalators, and lift devices | Wisconsin Administrative Code | chapter_identified | Base standards and amendments not parsed. |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Topic | Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| `amendment:usa-wi:cbc-code-suite-2021` | Commercial Building | Commercial model-code update | Wisconsin CBC moved to the 2021 ICC code set with Wisconsin-specific adjustments in SPS 361 to 366. | `src:usa-wi:dsps-commercial-buildings`; `src:usa-wi:sps-361-366`; `src:usa-wi:dsps-cbc-code-template` | verified_core |
| `amendment:usa-wi:cbc-transition-nov-2025` | Commercial Building | Transition/grace period | Current DSPS pages require plans submitted on or after 2025-11-01 to comply with the updated CBC. Earlier DSPS press releases used 2025-10-01 language, which is treated as superseded for this file. | `src:usa-wi:dsps-commercial-buildings`; `src:usa-wi:dsps-plan-review`; `src:usa-wi:dsps-cbc-extension-2025`; `src:usa-wi:dsps-cbc-2025-release` | verified_core |
| `amendment:usa-wi:sps365-plastic-venting` | Fuel Gas | Plastic venting materials | DSPS commercial-buildings page flags an SPS 365.0501(4) fuel-gas amendment requiring fuel-fired equipment to use plastic venting materials meeting UL 1738. | `src:usa-wi:dsps-commercial-buildings`; `src:usa-wi:sps-361-366` | partially_verified |
| `amendment:usa-wi:fire-nfpa1-2012` | Fire - operational / prevention | Current NFPA 1 edition | DSPS 2025 scope statement says current SPS 314 rules adopt/reference 2012 NFPA 1 and rulemaking will consider 2024 NFPA 1. | `src:usa-wi:fire-scope-2025`; `src:usa-wi:sps-314` | partially_verified |
| `amendment:usa-wi:electrical-2023-nec` | Electrical | Future NEC update | CR 26-016 rulemaking materials state the update adopts 2023 NEC; DSPS pending-rules page lists effective date 2026-09-01. | `src:usa-wi:electrical-cr26-016`; `src:usa-wi:dsps-pending-rules` | partially_verified |
| `amendment:usa-wi:plumbing-uniformity` | Plumbing | Uniform state plumbing code | CR 26-019 rule text describes SPS 381 to 387 as applying uniformly to plumbing design, construction, installation, supervision, maintenance, and inspection. | `src:usa-wi:plumbing-cr26-019`; `src:usa-wi:sps-381-387` | partially_verified |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-wi"
  model: "hybrid_state_delegated_local"
  enforcing_entities:
    - "DSPS Division of Industry Services for state-administered commercial/public-building and trades plan review and related services"
    - "local certified municipalities and delegated agents for certain commercial building plan review and inspection services"
    - "municipalities for statewide UDC enforcement"
    - "fire departments and municipalities participating in fire-dues programs for operational fire-prevention duties under SPS 314 framework"
  required_officials:
    - "certified commercial inspectors or local certified municipality personnel for delegated commercial review/inspection, where applicable"
    - "UDC-certified inspectors and municipal enforcement officials"
    - "fire inspectors/fire chiefs where SPS 314 or fire-dues duties apply"
  state_reserved_activities:
    - "statewide code adoption and Wisconsin Administrative Code publication"
    - "commercial/public-building plan review where not delegated or reserved to DSPS"
    - "petitions for variance from Wisconsin Administrative Code requirements"
    - "credentialing for UDC inspection and applicable trades programs"
    - "program-level code interpretation and materials/product evaluation"
  source_ids:
    - "src:usa-wi:dsps-commercial-buildings"
    - "src:usa-wi:dsps-plan-review"
    - "src:usa-wi:dsps-udc-page"
    - "src:usa-wi:fire-scope-2025"
  verification_status: "partially_verified"
  confidence: 0.68
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-wi"
  model: "statewide_uniform_with_limited_statutory_exceptions"
  applies_to_code_families:
    - "commercial/public-building construction standards"
    - "fire detection, prevention, or suppression local exceptions under Wis. Stat. § 101.02(7r), if statutory conditions are met"
    - "plumbing code uniformity under SPS 381 to 387 rulemaking text"
    - "UDC technical amendments unresolved; statewide uniform baseline verified"
  approval_required: true
  approving_authority_id: "ahj:usa-wi:dsps-dis"
  filing_required: true
  registry_exists: null
  registry_source_ids: []
  legal_basis_source_ids:
    - "src:usa-wi:stat-101-02-local-standards"
    - "src:usa-wi:court-abc-madison-2023"
    - "src:usa-wi:plumbing-cr26-019"
    - "src:usa-wi:dsps-udc-page"
  verification_status: "partially_verified"
  confidence: 0.58
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Wisconsin local enforcement authority and local amendment authority must be kept separate. Local certified municipalities or delegated agents may review/inspect certain projects, and municipalities enforce the UDC, but that does not by itself authorize more stringent local technical standards. For commercial/public-building construction standards, the verified statutory posture is strict conformity to applicable DSPS rules, subject to statutory exceptions such as qualifying pre-2013 local fire detection, prevention, or suppression ordinances that meet filing/approval conditions. The 2023 Wisconsin Court of Appeals discussion confirms that the local-preemption question can turn on whether an ordinance is effectively a building-code standard rather than a zoning/aesthetic regulation; production use should not classify non-building-code local ordinances without legal review.

### 6.4 Known Local Amendment Registries

| Registry ID | Registry Name | Scope | Source IDs | Status | Notes |
| --- | --- | --- | --- | --- | --- |
| `registry:usa-wi:local-commercial-exceptions` | Local commercial/public-building statutory exception inventory | Local ordinances allowed under Wis. Stat. § 101.02(7r) exceptions | `src:usa-wi:stat-101-02-local-standards` | unresolved | A DSPS registry or dataset was not located in this pass. |
| `registry:usa-wi:delegated-municipalities` | Delegated municipality / local certified municipality list | Commercial plan review and inspection service availability | `src:usa-wi:dsps-commercial-buildings`; `src:usa-wi:dsps-plan-review` | referenced_not_extracted | DSPS pages link to a delegated municipality list; data was not extracted into the report. |

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendment records were parsed into this report. Do not infer a local technical amendment from local enforcement status alone.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: scoped_not_built

Jurisdiction stack:

```text
Address
  -> State of Wisconsin
  -> County
  -> Municipality / unincorporated county area
  -> State-administered or delegated commercial plan review path, if commercial/public-building scope applies
  -> Municipality UDC enforcement path, if one- or two-family dwelling scope applies
  -> Fire department / fire district / municipality for operational fire-prevention duties, if applicable
  -> DSPS trade-specific review path for plumbing, POWTS, conveyances, amusement rides, and other reserved systems
  -> Applicable statewide code adoption records
  -> Applicable local exception/amendment records, only if separately verified
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | not selected | none | statewide | unknown | pending |
| County | not selected | none | statewide | unknown | pending |
| Municipality | not selected | none | statewide | unknown | pending |
| Fire District / Fire Department Service Area | not selected | none | partial / variable | unknown | pending |
| DSPS Delegated Municipality | DSPS delegated municipality list referenced but not extracted | `src:usa-wi:dsps-commercial-buildings`; `src:usa-wi:dsps-plan-review` | statewide where delegated municipalities exist | unknown | pending_extract |
| UDC Municipality Enforcement | DSPS UDC program page | `src:usa-wi:dsps-udc-page` | statewide | unknown | verified_program_level |

### 7.3 AHJ Contact Data

No AHJ contact dataset was populated. The next pass should extract DSPS contacts, delegated municipality lists, UDC municipal enforcement contacts, and fire department/fire inspection contacts into normalized AHJ records.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Source Type | Title / Description | Publisher | URL / Locator | Key Fields Supported | Status |
| --- | --- | --- | --- | --- | --- | --- |
| `src:usa-wi:dsps-commercial-buildings` | agency_page | DSPS Commercial Buildings program page | Wisconsin DSPS | https://dsps.wi.gov/Pages/Programs/CommercialBuildings/Default.aspx | DIS commercial authority, local certified municipalities, CBC 2021 ICC update, 2025-11-01 compliance date, DSPS code inserts, SPS 365 venting note | official_current_page |
| `src:usa-wi:dsps-plan-review` | agency_page | DSPS Division of Industry Services Plan Review page | Wisconsin DSPS | https://dsps.wi.gov/Pages/Programs/PlanReview/Default.aspx | Plan review thresholds, required system reviews, CBC 2021 update dates, delegated agents, voluntary commercial electrical plan review note | official_current_page |
| `src:usa-wi:dsps-rules-index` | agency_page | DSPS Administrative Rules and Statutes page | Wisconsin DSPS | https://dsps.wi.gov/Pages/Programs/AdministrativeRules.aspx | Chapter list for SPS 314, 316, 318, 320 to 325, 361 to 366, 381 to 387; statutes chapters 101 and 145 | official_index |
| `src:usa-wi:dsps-pending-rules` | agency_page | DSPS Pending Rules page | Wisconsin DSPS | https://dsps.wi.gov/Pages/RulesStatutes/PendingRules.aspx | Future rule statuses and effective dates for SPS 316 and SPS 381 to 387; drafting status for SPS 314 and SPS 320 to 325 | official_current_page |
| `src:usa-wi:sps-361-366` | admin_code | Wisconsin Commercial Building Code chapters SPS 361 to 366 | Wisconsin Legislature / LRB | https://docs.legis.wisconsin.gov/code/admin_code/sps/safety_and_buildings_and_environment/361_366 | Commercial code chapter structure; 2021 IBC incorporation through SPS 361.05; related 2021 ICC code set by DSPS template | official_code |
| `src:usa-wi:dsps-cbc-code-template` | agency_template_pdf | Commercial Buildings code template / code legend | Wisconsin DSPS | https://dsps.wi.gov/Documents/Programs/CommercialBuildings/codetemplate.pdf | 2025 Wisconsin SPS 361 to 366, 2021 IBC, 2021 IEBC, 2021 IECC, 2021 IMC, 2021 IFGC, 2017 ICC A117.1 | official_template_caveated |
| `src:usa-wi:cr23-007` | register_rule | CR 23-007 Wisconsin Commercial Building Code update | Wisconsin Legislature / LRB | https://docs.legis.wisconsin.gov/code/register/2025 | Register/order metadata for CBC update; effective date context | official_register_caveated |
| `src:usa-wi:dsps-cbc-2025-release` | agency_news_pdf | Upgraded Commercial Building Code to Take Effect September 1 | Wisconsin DSPS | https://dsps.wi.gov/Documents/NewsMedia/20250730CommercialBuildingCodeNewsRelease.pdf | 2025 CBC update context, Sept. 1 effective date, supplemental submission alignment | official_news_superseded_for_deadline |
| `src:usa-wi:dsps-cbc-extension-2025` | agency_news_pdf | DSPS Provides Extension to Building Professionals | Wisconsin DSPS | https://dsps.wi.gov/Documents/NewsMedia/20250916CBC_ExtensionNewsRelease.pdf | CBC grace-period extension context; current page used for final 2025-11-01 compliance date | official_news_caveated |
| `src:usa-wi:dsps-udc-page` | agency_page | One- & Two-Family Uniform Dwelling Code page | Wisconsin DSPS | https://dsps.wi.gov/Pages/Programs/UDC/Default.aspx | UDC statewide scope since 1980-06-01; enforcement in all municipalities; UDC credentials/materials role | official_current_page |
| `src:usa-wi:sps-320-325` | admin_code | Wisconsin Uniform Dwelling Code chapters SPS 320 to 325 | Wisconsin Legislature / LRB | https://docs.legis.wisconsin.gov/code/admin_code/sps/safety_and_buildings_and_environment/320_325 | UDC chapter structure and administrative rule location | official_code |
| `src:usa-wi:udc-scope-2025` | rulemaking_pdf | Statement of Scope: SPS 302, 305, and 320 to 325, Update to UDC | Wisconsin DSPS | https://dsps.wi.gov/Documents/RulesStatutes/SPS320to325NPPH.pdf | Pending UDC update; statewide uniform UDC purpose; last comprehensive update effective 2016-01-01 | official_rulemaking_pdf |
| `src:usa-wi:sps-314` | admin_code | SPS 314 Fire Prevention | Wisconsin Legislature / LRB | https://docs.legis.wisconsin.gov/code/admin_code/sps/safety_and_buildings_and_environment/314 | Fire prevention chapter location | official_code |
| `src:usa-wi:fire-scope-2025` | rulemaking_pdf | Statement of Scope: SPS 302, 305, and 314 Fire Prevention Code Review | Wisconsin DSPS | https://dsps.wi.gov/Documents/RulesStatutes/SPS314AS2.pdf | Current SPS 314 references 2012 NFPA 1; proposed review may consider 2024 NFPA 1; statutory authority excerpts; affected fire departments/municipalities | official_rulemaking_pdf |
| `src:usa-wi:dsps-fire-prevention` | agency_page | DSPS Fire Prevention program / reporting page | Wisconsin DSPS | https://dsps.wi.gov/Pages/Programs/FirePrevention/Default.aspx | Fire-prevention program reference and reporting context | official_page_limited_use |
| `src:usa-wi:sps-316` | admin_code | SPS 316 Electrical | Wisconsin Legislature / LRB | https://docs.legis.wisconsin.gov/code/admin_code/sps/safety_and_buildings_and_environment/316 | Electrical code chapter location | official_code |
| `src:usa-wi:electrical-cr26-016` | rulemaking_pdf_register | CR 26-016 / SPS 316 Electrical Code update materials | Wisconsin DSPS and Wisconsin Legislature / LRB | https://docs.legis.wisconsin.gov/code/register/2026/842a1/register/rule_notices/cr_26_016_hearing_information/cr_26_016_rule_text | Current SPS 316 adopts 2017 NEC; update adopts 2023 NEC; effective date confirmed by pending-rules page | official_rulemaking_caveated |
| `src:usa-wi:sps-381-387` | admin_code | SPS 381 to 387 Wisconsin Plumbing Code | Wisconsin Legislature / LRB | https://docs.legis.wisconsin.gov/code/admin_code/sps/safety_and_buildings_and_environment/381_387 | Plumbing code chapter set | official_code |
| `src:usa-wi:plumbing-cr26-019` | register_rule | CR 26-019 / SPS 381 to 387 Plumbing Code Review | Wisconsin Legislature / LRB | https://docs.legis.wisconsin.gov/code/register/2026/845a3/register/cr/cr_26_019_rule_text/cr_26_019_rule_text | Wisconsin Plumbing Code uniform applicability; future update context | official_register_rule |
| `src:usa-wi:sps-318` | admin_code | SPS 318 Elevators, Escalators and Lift Devices | Wisconsin Legislature / LRB | https://docs.legis.wisconsin.gov/code/admin_code/sps/safety_and_buildings_and_environment/318 | Conveyance chapter identification | official_code |
| `src:usa-wi:stat-101-02` | statute | Wis. Stat. § 101.02, DSPS powers/duties | Wisconsin Legislature / LRB | https://docs.legis.wisconsin.gov/statutes/statutes/101/02 | General DSPS powers and standards authority | official_statute |
| `src:usa-wi:stat-101-02-local-standards` | statute | Wis. Stat. § 101.02(7r), local minimum standards for public buildings/places of employment | Wisconsin Legislature / LRB | https://docs.legis.wisconsin.gov/document/statutes/101.02%287r%29%28f%29 | Local commercial/public-building strict-conformity rule and exceptions | official_statute_snippet |
| `src:usa-wi:stat-101-12` | statute | Wis. Stat. § 101.12, plan review for public buildings/places of employment | Wisconsin Legislature / LRB | https://docs.legis.wisconsin.gov/statutes/statutes/101/12 | Plan-review legal basis | official_statute |
| `src:usa-wi:stat-101-14` | statute | Wis. Stat. § 101.14, fire detection/prevention/suppression rule authority | Wisconsin Legislature / LRB | https://docs.legis.wisconsin.gov/statutes/statutes/101/14 | Fire-prevention statutory authority | official_statute |
| `src:usa-wi:stat-101-63` | statute | Wis. Stat. § 101.63, UDC standards authority | Wisconsin Legislature / LRB | https://docs.legis.wisconsin.gov/statutes/statutes/101/63 | UDC standards and inspection procedures authority | official_statute |
| `src:usa-wi:stat-101-65` | statute | Wis. Stat. § 101.65, municipal authority over one- and two-family dwellings | Wisconsin Legislature / LRB | https://docs.legis.wisconsin.gov/document/statutes/101.65%281%29%28c%29 | Municipal UDC jurisdiction and enforcement | official_statute_snippet |
| `src:usa-wi:stat-101-82` | statute | Wis. Stat. § 101.82, state electrical wiring code | Wisconsin Legislature / LRB | https://docs.legis.wisconsin.gov/statutes/statutes/101/82 | Electrical code statutory authority | official_statute |
| `src:usa-wi:court-abc-madison-2023` | court_opinion_pdf | Associated Builders & Contractors of Wisconsin, Inc. v. City of Madison discussion of Wis. Stat. § 101.02(7r) | Wisconsin Court of Appeals | https://www.wicourts.gov/ca/opinion/DisplayDocument.pdf?content=pdf&seqNo=711316 | Legal context for preemption/local ordinance classification under § 101.02(7r) | official_court_supplement |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| `src:usa-wi:dsps-cbc-code-template` | template_not_rule_text | DSPS code template is an official program aid, not enforceable rule text. | Use for code-family normalization; verify enforceable text in SPS 361 to 366 for production. |
| `src:usa-wi:dsps-cbc-2025-release` | superseded_deadline_language | July 2025 release indicated an earlier mandatory deadline; current DSPS pages and later extension context support 2025-11-01 for plan compliance. | Do not use for final compliance deadline without current page cross-check. |
| `src:usa-wi:dsps-cbc-extension-2025` | transitional_news | News release explains grace-period evolution, but current agency page is the controlling operational source used here. | Use as transition history only. |
| `src:usa-wi:cr23-007` | register_metadata_not_fully_extracted | CBC clearinghouse/register metadata was not fully parsed section-by-section. | Verify LRB register/order details before promoting to verified. |
| `src:usa-wi:udc-scope-2025` | pending_rulemaking | UDC scope statement identifies objectives and history but does not itself update the code. | Use for pending-status and historical context only. |
| `src:usa-wi:fire-scope-2025` | pending_rulemaking | Fire scope statement describes current 2012 NFPA 1 reference and possible 2024 NFPA 1 consideration; it is not final adoption of 2024 NFPA 1. | Track pending rule; do not mark 2024 NFPA 1 adopted. |
| `src:usa-wi:electrical-cr26-016` | future_rule | Electrical update is listed effective 2026-09-01; final codified text should be checked after effective date. | Treat as future/pending until effective date and codified rule text are confirmed. |
| `src:usa-wi:plumbing-cr26-019` | future_rule | Plumbing update is listed effective 2026-09-01; final codified text should be checked after effective date. | Treat as future/pending until effective date and codified rule text are confirmed. |
| `src:usa-wi:stat-101-02-local-standards` | legal_interpretation_needed | Statutory exceptions and court interpretations require legal analysis before classifying local ordinances. | Use for high-level statewide posture; escalate close calls. |
| `src:usa-wi:court-abc-madison-2023` | legal_context | Court opinion is official but fact-specific. | Use as interpretive context, not as a substitute for ordinance-specific legal review. |

### 8.3 Supplemental Sources

No non-official sources are used for primary field support. Official court opinion material is treated as official supplemental legal context for local-preemption interpretation.

### 8.4 Source Extraction Metadata

| Extraction ID | Source IDs | Extracted On | Method | Coverage | Notes |
| --- | --- | --- | --- | --- | --- |
| `extract:usa-wi:2026-06-26:agency-pages` | `src:usa-wi:dsps-commercial-buildings`; `src:usa-wi:dsps-plan-review`; `src:usa-wi:dsps-udc-page`; `src:usa-wi:dsps-rules-index`; `src:usa-wi:dsps-pending-rules` | 2026-06-26 | HTML page review | Core authority, chapter lists, current/future date status | Current pages preferred over older press releases. |
| `extract:usa-wi:2026-06-26:rulemaking-pdfs` | `src:usa-wi:udc-scope-2025`; `src:usa-wi:fire-scope-2025`; `src:usa-wi:electrical-cr26-016`; `src:usa-wi:plumbing-cr26-019` | 2026-06-26 | PDF/register text review with caveats | Pending updates and current edition context | PDFs/register documents require final codification checks. |
| `extract:usa-wi:2026-06-26:statutes` | `src:usa-wi:stat-101-02`; `src:usa-wi:stat-101-02-local-standards`; `src:usa-wi:stat-101-12`; `src:usa-wi:stat-101-14`; `src:usa-wi:stat-101-63`; `src:usa-wi:stat-101-65`; `src:usa-wi:stat-101-82` | 2026-06-26 | official statute snippets and DSPS cross-references | Authority and local amendment posture | Full subsection extraction remains an open production task. |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| `report` | `report.status` | partially_verified | verified | 1.00 | none | Core authority and adoption fields now have official sources; unresolved items remain explicit. |
| `report` | `report.last_updated` | 2026-06-26 | verified | 1.00 | none | Updated during this population pass. |
| `ahj:usa-wi:dsps-dis` | authority.name | Wisconsin Department of Safety and Professional Services, Division of Industry Services | verified_core | 0.82 | `src:usa-wi:dsps-commercial-buildings`; `src:usa-wi:dsps-plan-review` | Agency/division role verified from DSPS pages. |
| `adoption-rec:usa-wi:cbc-2021` | base code | 2021 ICC code set for CBC | verified_core | 0.82 | `src:usa-wi:dsps-commercial-buildings`; `src:usa-wi:sps-361-366`; `src:usa-wi:dsps-cbc-code-template` | Section-level amendments not parsed. |
| `adoption-rec:usa-wi:cbc-2021` | mandatory date | 2025-11-01 | verified_core | 0.85 | `src:usa-wi:dsps-commercial-buildings`; `src:usa-wi:dsps-plan-review` | Current DSPS pages control over older release language. |
| `adoption-rec:usa-wi:udc-current` | statewide scope | UDC statewide for 1- and 2-family dwellings built since 1980-06-01 | verified_core | 0.84 | `src:usa-wi:dsps-udc-page` | Applies by dwelling type and date. |
| `local-enforcement:usa-wi` | UDC enforcement | enforced in all Wisconsin municipalities | verified_core | 0.84 | `src:usa-wi:dsps-udc-page`; `src:usa-wi:stat-101-65` | Municipality contact data not extracted. |
| `adoption-rec:usa-wi:fire-nfpa1-2012` | current fire standard | 2012 NFPA 1 referenced by current SPS 314 rules | partially_verified | 0.70 | `src:usa-wi:fire-scope-2025`; `src:usa-wi:sps-314` | Source is a DSPS rulemaking scope statement; final SPS 314 text should be section-checked. |
| `adoption-rec:usa-wi:electrical-2017-nec` | current electrical standard | 2017 NEC current under SPS 316 | partially_verified | 0.70 | `src:usa-wi:electrical-cr26-016`; `src:usa-wi:sps-316` | Rulemaking text states current code; current codified section should be rechecked. |
| `adoption-rec:usa-wi:electrical-2023-nec-future` | future effective date | 2026-09-01 | partially_verified | 0.78 | `src:usa-wi:dsps-pending-rules`; `src:usa-wi:electrical-cr26-016` | Follow up after effective date. |
| `adoption-rec:usa-wi:plumbing-2026-update` | future effective date | 2026-09-01 | partially_verified | 0.78 | `src:usa-wi:dsps-pending-rules`; `src:usa-wi:plumbing-cr26-019` | Follow up after effective date. |
| `local-amendment-rule:usa-wi` | model | statewide_uniform_with_limited_statutory_exceptions | partially_verified | 0.58 | `src:usa-wi:stat-101-02-local-standards`; `src:usa-wi:court-abc-madison-2023`; `src:usa-wi:plumbing-cr26-019` | Requires legal review before production classification of any specific local ordinance. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | All `src:usa-wi:*` identifiers cited in the body are listed in Section 8. |
| All authority IDs resolve | pass | Authority IDs used in Section 2 and local enforcement records are declared or intentionally scoped to local entity classes. |
| All current code families have adoption rows | pass | Rows are present for each template family; unresolved subfields remain explicit. |
| Building and operational fire code are separated | pass | CBC construction/fire references and SPS 314 operational fire prevention are separate rows. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Date fields are separate; null is used where no supported date was verified. |
| Effective dates are valid ISO dates | pass | Entered dates use ISO `YYYY-MM-DD` format. |
| No impossible date sequences | pass | Known effective and mandatory dates are sequenced consistently; future rule dates are not backdated. |
| Transition rules have explicit trigger conditions | pass | CBC, UDC, electrical, and plumbing transition rows include trigger descriptions. |
| Permit-date or plan-submittal logic is captured where applicable | pass | CBC plan-submittal rule captured; trade-specific permit transition details remain open. |
| Local enforcement model classified | pass | Hybrid state/delegated/local model is described with caveats. |
| Local amendment rule classified | pass | High-level statewide-uniform/limited-exception posture captured; registry remains unresolved. |
| AHJ confirmation metadata present | fail | AHJ contact records and delegated municipality dataset are not populated. |
| Official-source caveats captured | pass | Caveats are listed in Section 8.2. |
| Leftover template markers removed | pass | Configured marker scan passed; no scanned placeholder markers remained after validation. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `issue:usa-wi:001` | high | fire authority and update | SPS 314 current 2012 NFPA 1 reference is supported by DSPS scope statement, but final rule text and future 2024 NFPA 1 adoption status need follow-up. | Parse SPS 314 sections and monitor DSPS pending-rules page and register documents. | null | null | open |
| `issue:usa-wi:002` | high | electrical and plumbing 2026 updates | DSPS lists SPS 316 and SPS 381 to 387 updates effective 2026-09-01; final codified text and any permit/application transition rules must be confirmed after effective date. | Re-check official Administrative Code and DSPS guidance on or after 2026-09-01. | null | 2026-09-15 | open |
| `issue:usa-wi:003` | medium | UDC update | SPS 302, 305, and 320 to 325 UDC update is in drafting status; final rule and transition provisions are not available in this report. | Monitor DSPS pending rules, UDC Council materials, and LRB register. | null | null | open |
| `issue:usa-wi:004` | medium | local amendment exceptions | Local commercial/public-building standards are generally strict-conformity, but exception inventory and any DSPS filing/approval registry were not located. | Extract Wis. Stat. § 101.02(7r) in full, SPS 361.03 local ordinance provisions, and any DSPS local exception records. | null | null | open |
| `issue:usa-wi:005` | medium | AHJ/delegated municipality data | Delegated municipality list and municipal UDC enforcement contacts were referenced but not extracted. | Build AHJ/delegated municipality dataset and link it to jurisdiction resolver. | null | null | open |
| `issue:usa-wi:006` | low | elevator/conveyance base standards | SPS 318 chapter was identified, but base standard editions and amendments were not parsed. | Parse SPS 318 and DSPS conveyance guidance. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| `watch:usa-wi:commercial-buildings` | `src:usa-wi:dsps-commercial-buildings` | html_diff | monthly | CBC deadline, code inserts, or program guidance changes | 2026-06-26 | active |
| `watch:usa-wi:plan-review` | `src:usa-wi:dsps-plan-review` | html_diff | monthly | Plan-review thresholds, delegated-agent links, or system review requirements change | 2026-06-26 | active |
| `watch:usa-wi:pending-rules` | `src:usa-wi:dsps-pending-rules` | html_diff | biweekly_until_2026_09_15 | SPS 314, 316, 320 to 325, 381 to 387, or conveyance rule status changes | 2026-06-26 | active |
| `watch:usa-wi:admin-code-commercial` | `src:usa-wi:sps-361-366` | html_diff | monthly | SPS 361 to 366 amendment or publication change | 2026-06-26 | active |
| `watch:usa-wi:admin-code-udc` | `src:usa-wi:sps-320-325` | html_diff | monthly | SPS 320 to 325 amendment or publication change | 2026-06-26 | active |
| `watch:usa-wi:admin-code-fire` | `src:usa-wi:sps-314` | html_diff | monthly | SPS 314 amendment or NFPA 1 edition change | 2026-06-26 | active |
| `watch:usa-wi:admin-code-electrical` | `src:usa-wi:sps-316` | html_diff | biweekly_until_2026_09_15 | 2023 NEC update codified or transition guidance published | 2026-06-26 | active |
| `watch:usa-wi:admin-code-plumbing` | `src:usa-wi:sps-381-387` | html_diff | biweekly_until_2026_09_15 | 2026 plumbing update codified or transition guidance published | 2026-06-26 | active |
| `watch:usa-wi:local-standards-statute` | `src:usa-wi:stat-101-02-local-standards` | statute_diff | quarterly | Change to § 101.02(7r) or related local ordinance exceptions | 2026-06-26 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-26 | Populated Wisconsin report from baseline draft; replaced generic placeholders; added source registry, authority model, adoption matrix, transition rules, local enforcement/amendment posture, QA, and open issues. | `report:usa-wi`; `ahj:usa-wi:dsps-dis`; `adoption-rec:usa-wi:cbc-2021`; `local-enforcement:usa-wi`; `local-amendment-rule:usa-wi` | `src:usa-wi:dsps-commercial-buildings`; `src:usa-wi:dsps-plan-review`; `src:usa-wi:dsps-udc-page`; `src:usa-wi:dsps-rules-index`; `src:usa-wi:dsps-pending-rules` | ChatGPT | Status set to `partially_verified` because core authority/adoption fields are official-source-backed while AHJ contacts and several detailed rule transitions remain unresolved. |
