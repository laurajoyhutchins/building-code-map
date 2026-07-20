---
state:
  state_id: "US-NH"
  name: "New Hampshire"
  abbreviation: "NH"
report:
  report_id: "state-report:usa-nh"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-26"
  last_verified: null
  reviewed_by: null
risk:
  overall_confidence: 0.72 # 0.00 - 1.00
  risk_flags:
    - "official_state_pages_blocked_automated_fetch"
    - "state_amendments_not_fully_parsed"
    - "municipal_registry_not_extracted"
    - "future_2024_code_transition_pending_2026_07_01"
  open_questions_count: 6

---

# State Building Code Authority Report: New Hampshire

## 1. Executive Summary

- **Authority model:** New Hampshire uses a statewide minimum-code model. RSA 155-A defines the New Hampshire Building Code / State Building Code by reference to specified ICC codes and NFPA 70, as amended by the State Building Code Review Board and ratified by the legislature. RSA 153 separately defines the New Hampshire Fire Code / State Fire Code by reference to NFPA 101 and NFPA 1, as amended by the State Board of Fire Control and ratified by the General Court.

- **Statewide code status:** Partially verified. As of this report date, RSA 155-A:1, IV lists 2021 IBC, 2021 IEBC, 2021 IPC, 2021 IMC, 2018 IECC, 2021 ISPSC, 2021 IRC, and 2023 NFPA 70 / NEC. RSA 153:1, VI-a lists 2021 NFPA 101 and 2021 NFPA 1 for the State Fire Code. Official state building-code materials also identify a future 2024-code update effective 2026-07-01 with permit requests required to follow those codes beginning 2027-01-01; that future transition should be rechecked after 2026-07-01.

- **Local enforcement model:** Local enforcement is separate from statewide adoption. RSA 155-A:7 gives the local enforcement agency authority to enforce the State Building Code and the local fire chief authority to enforce the State Fire Code. Where there is no local enforcement agency or qualified-third-party contract, the State Fire Marshal or designee may enforce on written municipal request.

- **Local amendment posture:** Local amendments are constrained and changing. RSA 155-A:3 is transitioning to a 2026-07-01 framework that allows local enforcement mechanisms and administrative amendments but prohibits technical amendments except in the limited circumstance where the state building code is more than two editions behind the published model codes. Municipal ordinances adopted under the local-amendment provisions require Building Code Review Board confirmation. Fire-code technical amendments are likewise constrained under RSA 153:5, with administrative amendments reviewed under the fire-code process.

- **Known transition periods or pending changes:** RSA 155-A:2 creates a permit-application date rule and a six-month concurrency period after a code adopted under RSA 155-A:1, IV becomes effective. Current official state materials indicate 2024 building codes effective 2026-07-01 and mandatory for permit requests beginning 2027-01-01.

- **Production readiness:** narrow_use_only

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: The New Hampshire Building Code is established in RSA 155-A and amended
    through the State Building Code Review Board / legislative ratification structure.
  confidence: 0.85
  source_ids:
  - src:usa-nh:rsa-155-a-1
  - src:usa-nh:rsa-155-a-10
  - src:usa-nh:bcrb-page
- topic: Primary building code editions
  finding: Current RSA 155-A:1, IV lists 2021 IBC, IEBC, IPC, IMC, ISPSC, IRC; 2018
    IECC; and 2023 NEC.
  confidence: 0.9
  source_ids:
  - src:usa-nh:rsa-155-a-1
- topic: Fire code authority
  finding: RSA 153:1 defines the State Fire Code as 2021 NFPA 101 and 2021 NFPA 1,
    amended by the State Board of Fire Control and ratified under RSA 153:5.
  confidence: 0.85
  source_ids:
  - src:usa-nh:rsa-153-1
  - src:usa-nh:rsa-153-5
  - src:usa-nh:dfs-state-fire-code
- topic: Local enforcement
  finding: Local enforcement agency enforces the State Building Code; local fire chief
    enforces the State Fire Code; State Fire Marshal may enforce where local enforcement
    is absent and municipality requests it.
  confidence: 0.85
  source_ids:
  - src:usa-nh:rsa-155-a-7
- topic: Local amendments
  finding: On and after 2026-07-01, local technical building-code amendments are generally
    prohibited except when the state code is more than two editions behind the model
    codes; administrative amendments are reserved locally if confirmed.
  confidence: 0.75
  source_ids:
  - src:usa-nh:rsa-155-a-3
  - src:usa-nh:rsa-155-a-10
- topic: Effective / operative date rule
  finding: Code in effect when the building-permit application is received remains
    in effect for the permit work; a six-month concurrency period follows code effective
    dates.
  confidence: 0.85
  source_ids:
  - src:usa-nh:rsa-155-a-2
- topic: Future building-code update
  finding: Official state building-code page identifies 2024 codes effective 2026-07-01
    and permit requests required to follow them beginning 2027-01-01.
  confidence: 0.7
  source_ids:
  - src:usa-nh:dfs-state-building-code
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-nh:bcrb |
| Authority name | New Hampshire State Building Code Review Board |
| Authority type | state board |
| Legal basis | RSA 155-A:10; RSA 155-A:1, IV |
| Role | Reviews and recommends building-code amendments; state building-code amendments are ratified by the legislature. Municipal amendments submitted under RSA 155-A are reviewed/confirmed by the board. |
| Enforcement model | statewide minimum code with local enforcement and limited State Fire Marshal backstop |
| Source IDs | src:usa-nh:rsa-155-a-1; src:usa-nh:rsa-155-a-10; src:usa-nh:bcrb-page |
| Verification status | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-nh:bcrb | State Building Code Review Board | Building-code amendment review and recommendation; legislature ratifies amendments | RSA 155-A:1; RSA 155-A:10 | src:usa-nh:rsa-155-a-1; src:usa-nh:rsa-155-a-10 | partially_verified |
| Residential | ahj:usa-nh:bcrb | State Building Code Review Board | Same as building code for IRC adoption/amendments | RSA 155-A:1; RSA 155-A:10 | src:usa-nh:rsa-155-a-1; src:usa-nh:rsa-155-a-10 | partially_verified |
| Existing Building / Rehabilitation | ahj:usa-nh:bcrb | State Building Code Review Board | Same as building code for IEBC adoption/amendments | RSA 155-A:1; RSA 155-A:10 | src:usa-nh:rsa-155-a-1; src:usa-nh:rsa-155-a-10 | partially_verified |
| Mechanical | ahj:usa-nh:bcrb; ahj:usa-nh:mechanical-licensing-board | State Building Code Review Board; Mechanical Licensing Board | IMC is part of State Building Code; trade licensing/program authority remains with mechanical board subject to RSA 155-A limitations | RSA 155-A:1; RSA 155-A:2, X; RSA 153:27-a | src:usa-nh:rsa-155-a-1; src:usa-nh:rsa-155-a-2; src:usa-nh:oplc-mechanical-board | partially_verified |
| Plumbing | ahj:usa-nh:bcrb; ahj:usa-nh:mechanical-licensing-board | State Building Code Review Board; Mechanical Licensing Board | IPC is part of State Building Code; plumbing licensing/program authority remains with mechanical board subject to RSA 155-A limitations | RSA 155-A:1; RSA 155-A:2, X; RSA 153:27-a | src:usa-nh:rsa-155-a-1; src:usa-nh:rsa-155-a-2; src:usa-nh:oplc-mechanical-board | partially_verified |
| Fuel Gas | ahj:usa-nh:fire-code; ahj:usa-nh:mechanical-licensing-board | State Board of Fire Control / State Fire Marshal; Mechanical Licensing Board | IFGC is not listed as adopted in RSA 155-A:1; fuel-gas systems are treated through the State Fire Code and mechanical licensing framework in amendment materials | RSA 153:1; RSA 153:27-a | src:usa-nh:rsa-153-1; src:usa-nh:summary-2021-building-amendments-2025; src:usa-nh:oplc-mechanical-board | partially_verified |
| Electrical | ahj:usa-nh:bcrb; ahj:usa-nh:electricians-board | State Building Code Review Board; State Electricians' Board | NFPA 70 / NEC is part of State Building Code; electrician program authority is preserved subject to BCRB/legislative approval for code changes | RSA 155-A:1; RSA 155-A:2, X | src:usa-nh:rsa-155-a-1; src:usa-nh:rsa-155-a-2 | partially_verified |
| Energy | ahj:usa-nh:bcrb | State Building Code Review Board | IECC is part of State Building Code | RSA 155-A:1 | src:usa-nh:rsa-155-a-1 | partially_verified |
| Fire - construction references | ahj:usa-nh:fire-code; ahj:usa-nh:bcrb | State Board of Fire Control / State Fire Marshal; State Building Code Review Board | State Building Code and State Fire Code both apply; more life-safety-protective rule controls where they conflict | RSA 155-A:2; RSA 153:1; RSA 153:5 | src:usa-nh:rsa-155-a-2; src:usa-nh:rsa-153-1; src:usa-nh:rsa-153-5 | partially_verified |
| Fire - operational / prevention code | ahj:usa-nh:fire-code | State Board of Fire Control / State Fire Marshal | NFPA 101 and NFPA 1 adoption/amendment framework | RSA 153:1; RSA 153:5 | src:usa-nh:rsa-153-1; src:usa-nh:rsa-153-5; src:usa-nh:dfs-state-fire-code | partially_verified |
| Accessibility | ahj:usa-nh:bcrb | State Building Code Review Board / accessibility certifier framework | Accessibility standards for public buildings are in the State Building Code with certification/enforcement provisions in RSA 155-A:5 through 155-A:5-b | RSA 155-A:5; RSA 155-A:5-a; RSA 155-A:5-b | src:usa-nh:rsa-155-a-5 | partially_verified |
| Elevator / Conveyance | ahj:usa-nh:dol-elevator | New Hampshire Department of Labor, Elevator and Accessibility Lift program | Elevator and accessibility lift inspection/certificate program | RSA 157-B; Lab 1300 rules | src:usa-nh:dol-elevator | partially_verified |

### 2.3 Authority Hierarchy Notes

The statewide code exists independently of local adoption. RSA 155-A:2 requires buildings, building components, and structures constructed in New Hampshire to comply with the State Building Code and State Fire Code. Local jurisdictions are the ordinary building-permit, fee, certificate-of-occupancy, and inspection layer where an enforcement mechanism exists, while the State Fire Marshal handles state-owned buildings and may provide backstop enforcement in municipalities without local enforcement upon request.

The fire-code system is legally distinct from the building-code system. RSA 153:1 defines the State Fire Code and RSA 153:5 gives the State Fire Marshal and local fire chiefs a fire-code role for new construction, additions, and alterations. RSA 155-A:2 resolves building/fire conflicts by applying the code that creates the greater degree of life safety.

Trade boards retain licensing/program authority, but RSA 155-A:2 states that code changes proposed by labor, plumbing/mechanical, or electrical licensing authorities are not enforceable until approved by the State Building Code Review Board and ratified by the legislature.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-nh:001 | ahj:usa-nh:legislature | ratifies | State Building Code amendments reviewed/recommended by ahj:usa-nh:bcrb | src:usa-nh:rsa-155-a-1; src:usa-nh:rsa-155-a-10 | partially_verified |
| edge:usa-nh:002 | ahj:usa-nh:bcrb | reviews_confirms | municipal building-code amendments/ordinances submitted under RSA 155-A | src:usa-nh:rsa-155-a-3; src:usa-nh:rsa-155-a-10 | partially_verified |
| edge:usa-nh:003 | ahj:usa-nh:fire-code | adopts_amendments_with | State Fire Marshal / Commissioner approval / legislative ratification structure under RSA 153:5 | src:usa-nh:rsa-153-5 | partially_verified |
| edge:usa-nh:004 | local_enforcement_agency | enforces | State Building Code in local jurisdiction | src:usa-nh:rsa-155-a-7 | partially_verified |
| edge:usa-nh:005 | local_fire_chief | enforces | State Fire Code in local jurisdiction | src:usa-nh:rsa-155-a-7; src:usa-nh:rsa-153-5 | partially_verified |
| edge:usa-nh:006 | ahj:usa-nh:state-fire-marshal | backstop_enforces | Building and fire code where no local enforcement agency/qualified third party exists and municipality requests enforcement | src:usa-nh:rsa-155-a-7 | partially_verified |
| edge:usa-nh:007 | ahj:usa-nh:dol-elevator | administers | Elevator/accessibility lift inspections and certificates | src:usa-nh:dol-elevator | partially_verified |

### 2.5 Authority ID Registry

| Authority ID | Authority / Actor | Status | Notes |
| --- | --- | --- | --- |
| ahj:usa-nh:legislature | New Hampshire General Court | partially_verified | Ratifies State Building Code and State Fire Code amendments through statutory process. |
| ahj:usa-nh:bcrb | State Building Code Review Board | partially_verified | Primary building-code amendment review and municipal amendment confirmation actor. |
| ahj:usa-nh:fire-code | State Board of Fire Control / State Fire Marshal fire-code system | partially_verified | State Fire Code amendment and enforcement framework under RSA 153. |
| ahj:usa-nh:state-fire-marshal | New Hampshire State Fire Marshal | partially_verified | State-building permit/inspection role and local-enforcement backstop where statute permits. |
| ahj:usa-nh:mechanical-licensing-board | Mechanical Licensing Board | partially_verified | Mechanical, plumbing, and fuel-gas licensing program authority; code changes subject to RSA 155-A review/ratification limits. |
| ahj:usa-nh:electricians-board | State Electricians' Board | unresolved | Electrical licensing authority is referenced by RSA 155-A framework, but board statute/rules were not independently parsed in this pass. |
| ahj:usa-nh:dol-elevator | New Hampshire Department of Labor Elevator and Accessibility Lift program | partially_verified | Elevator/accessibility lift program source located; detailed rule parsing remains open. |
| local_enforcement_agency | Municipal/local building-code enforcement agency | partially_verified | Enforces State Building Code where local enforcement mechanism exists. |
| local_fire_chief | Municipal fire chief | partially_verified | Enforces State Fire Code locally. |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | New Hampshire Building Code / State Building Code | International Building Code | 2021 | current until scheduled 2026-07-01 transition is verified operative | 2025-06-23 | 2025-07-01 | 2025-07-01 | 2026-01-01 for post-effective six-month concurrency if applicable | Permit-application date controls; six-month concurrency after effective date | src:usa-nh:rsa-155-a-1; src:usa-nh:rsa-155-a-2; src:usa-nh:hb134-2025 |
| Residential | New Hampshire Building Code / State Building Code | International Residential Code | 2021 | current until scheduled 2026-07-01 transition is verified operative | 2025-06-23 | 2025-07-01 | 2025-07-01 | 2026-01-01 for post-effective six-month concurrency if applicable | Permit-application date controls; six-month concurrency after effective date | src:usa-nh:rsa-155-a-1; src:usa-nh:rsa-155-a-2; src:usa-nh:hb134-2025 |
| Existing Building / Rehabilitation | New Hampshire Building Code / State Building Code | International Existing Building Code | 2021 | current until scheduled 2026-07-01 transition is verified operative | 2025-06-23 | 2025-07-01 | 2025-07-01 | 2026-01-01 for post-effective six-month concurrency if applicable | Permit-application date controls; six-month concurrency after effective date | src:usa-nh:rsa-155-a-1; src:usa-nh:rsa-155-a-2; src:usa-nh:hb134-2025 |
| Mechanical | New Hampshire Building Code / State Building Code | International Mechanical Code | 2021 | current until scheduled 2026-07-01 transition is verified operative | 2025-06-23 | 2025-07-01 | 2025-07-01 | 2026-01-01 for post-effective six-month concurrency if applicable | Permit-application date controls; six-month concurrency after effective date | src:usa-nh:rsa-155-a-1; src:usa-nh:rsa-155-a-2; src:usa-nh:hb134-2025 |
| Plumbing | New Hampshire Building Code / State Building Code | International Plumbing Code | 2021 | current until scheduled 2026-07-01 transition is verified operative | 2025-06-23 | 2025-07-01 | 2025-07-01 | 2026-01-01 for post-effective six-month concurrency if applicable | Permit-application date controls; six-month concurrency after effective date | src:usa-nh:rsa-155-a-1; src:usa-nh:rsa-155-a-2; src:usa-nh:hb134-2025 |
| Fuel Gas | New Hampshire Fire Code / State Fire Code fuel-gas treatment; not IFGC as a separate RSA 155-A adoption | NFPA-based State Fire Code; IFGC not adopted as standalone building-code family | 2021 NFPA 1 / NFPA 101 framework; fuel-gas provisions per state amendments | current, but amendment details require full parsing | 2024-08-02 | 2024-08-13 | 2024-08-13 | null | Building/fire coordination applies; fuel-gas provisions require amendment-level review | src:usa-nh:rsa-153-1; src:usa-nh:dfs-state-fire-code; src:usa-nh:summary-2021-building-amendments-2025 |
| Electrical | New Hampshire Building Code / State Building Code | NFPA 70, National Electrical Code | 2023 | current until scheduled 2026-07-01 transition is verified operative | 2025-06-23 | 2025-07-01 | 2025-07-01 | 2026-01-01 for post-effective six-month concurrency if applicable | Permit-application date controls; six-month concurrency after effective date; AFCI limitation in RSA 155-A:2, XI | src:usa-nh:rsa-155-a-1; src:usa-nh:rsa-155-a-2; src:usa-nh:hb134-2025 |
| Energy | New Hampshire Building Code / State Building Code | International Energy Conservation Code | 2018 | current until scheduled 2026-07-01 transition is verified operative | 2025-06-23 | 2025-07-01 | 2025-07-01 | 2026-01-01 for post-effective six-month concurrency if applicable | Permit-application date controls; six-month concurrency after effective date | src:usa-nh:rsa-155-a-1; src:usa-nh:rsa-155-a-2; src:usa-nh:hb134-2025 |
| Fire - construction references | New Hampshire Fire Code / State Fire Code | NFPA 101 Life Safety Code; NFPA 1 Fire Code | 2021 | current; potential 2026 legislation should be monitored | 2024-08-02 | 2024-08-13 | 2024-08-13 | null | Local fire chief / State Fire Marshal use State Fire Code; greater life-safety rule controls conflicts | src:usa-nh:rsa-153-1; src:usa-nh:rsa-153-5; src:usa-nh:dfs-state-fire-code |
| Fire - operational / prevention code | New Hampshire Fire Code / State Fire Code | NFPA 101 Life Safety Code; NFPA 1 Fire Code | 2021 | current; potential 2026 legislation should be monitored | 2024-08-02 | 2024-08-13 | 2024-08-13 | null | State Fire Code and associated rules apply; local fire chief and State Fire Marshal roles preserved | src:usa-nh:rsa-153-1; src:usa-nh:rsa-153-5; src:usa-nh:dfs-state-fire-code |
| Accessibility | State Building Code accessibility standards for public buildings | IBC accessibility provisions / referenced accessibility standards | tied to adopted IBC edition | current but detailed standard not parsed | 2025-06-23 | 2025-07-01 | 2025-07-01 | null | Certification/enforcement system under RSA 155-A:5 through 155-A:5-b | src:usa-nh:rsa-155-a-5; src:usa-nh:rsa-155-a-1 |
| Elevator / Conveyance | Elevator and Accessibility Lift Law / Lab 1300 | state elevator/accessibility lift rules; model-code edition unresolved | unresolved | partially_verified program authority; code-edition details unresolved | 2025-07-16 | 2025-07-16 | 2025-07-16 | null | Annual inspection certificate noted in RSA 157-B materials; full adoption matrix requires Lab 1300 parsing | src:usa-nh:dol-elevator |

### 3.2 Adoption Records

| Adoption Record ID | Code Family | Adoption Instrument | Adopted Code / Edition | Adoption Date | Effective Date | Operative Date | Mandatory Date | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| adoption:usa-nh:building-2021-2025 | Building | HB 134 / RSA 155-A:1, IV | 2021 IBC | 2025-06-23 | 2025-07-01 | 2025-07-01 | 2026-01-01 if six-month concurrency applies | src:usa-nh:rsa-155-a-1; src:usa-nh:hb134-2025; src:usa-nh:rsa-155-a-2 | partially_verified |
| adoption:usa-nh:residential-2021-2025 | Residential | HB 134 / RSA 155-A:1, IV | 2021 IRC | 2025-06-23 | 2025-07-01 | 2025-07-01 | 2026-01-01 if six-month concurrency applies | src:usa-nh:rsa-155-a-1; src:usa-nh:hb134-2025; src:usa-nh:rsa-155-a-2 | partially_verified |
| adoption:usa-nh:existing-2021-2025 | Existing Building / Rehabilitation | HB 134 / RSA 155-A:1, IV | 2021 IEBC | 2025-06-23 | 2025-07-01 | 2025-07-01 | 2026-01-01 if six-month concurrency applies | src:usa-nh:rsa-155-a-1; src:usa-nh:hb134-2025; src:usa-nh:rsa-155-a-2 | partially_verified |
| adoption:usa-nh:mechanical-2021-2025 | Mechanical | HB 134 / RSA 155-A:1, IV | 2021 IMC | 2025-06-23 | 2025-07-01 | 2025-07-01 | 2026-01-01 if six-month concurrency applies | src:usa-nh:rsa-155-a-1; src:usa-nh:hb134-2025; src:usa-nh:rsa-155-a-2 | partially_verified |
| adoption:usa-nh:plumbing-2021-2025 | Plumbing | HB 134 / RSA 155-A:1, IV | 2021 IPC | 2025-06-23 | 2025-07-01 | 2025-07-01 | 2026-01-01 if six-month concurrency applies | src:usa-nh:rsa-155-a-1; src:usa-nh:hb134-2025; src:usa-nh:rsa-155-a-2 | partially_verified |
| adoption:usa-nh:energy-2018-2025 | Energy | HB 134 / RSA 155-A:1, IV | 2018 IECC | 2025-06-23 | 2025-07-01 | 2025-07-01 | 2026-01-01 if six-month concurrency applies | src:usa-nh:rsa-155-a-1; src:usa-nh:hb134-2025; src:usa-nh:rsa-155-a-2 | partially_verified |
| adoption:usa-nh:electrical-2023-2025 | Electrical | HB 134 / RSA 155-A:1, IV | 2023 NFPA 70 / NEC | 2025-06-23 | 2025-07-01 | 2025-07-01 | 2026-01-01 if six-month concurrency applies | src:usa-nh:rsa-155-a-1; src:usa-nh:hb134-2025; src:usa-nh:rsa-155-a-2 | partially_verified |
| adoption:usa-nh:ispsc-2021-2025 | Swimming Pool and Spa | HB 134 / RSA 155-A:1, IV | 2021 ISPSC | 2025-06-23 | 2025-07-01 | 2025-07-01 | 2026-01-01 if six-month concurrency applies | src:usa-nh:rsa-155-a-1; src:usa-nh:hb134-2025; src:usa-nh:rsa-155-a-2 | partially_verified |
| adoption:usa-nh:fire-2021-2024 | Fire | RSA 153:1, VI-a / State Fire Code update | 2021 NFPA 101 and 2021 NFPA 1 | 2024-08-02 | 2024-08-13 | 2024-08-13 | null | src:usa-nh:rsa-153-1; src:usa-nh:rsa-153-5; src:usa-nh:dfs-state-fire-code | partially_verified |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

RSA 155-A:2 contains the main building-code date rule. The State Building Code in effect when the permit application is received by the governing authority remains in effect for the duration of the permit work. The same section establishes a six-month concurrency period after the effective date of a code adopted under RSA 155-A:1, IV, allowing applicants to elect either the immediately prior code or the newly adopted code, but not a mixture of the two.

Fire-code effective-date rules were not fully parsed beyond the 2024 State Fire Code update effective date and RSA 153:5 authority model.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-nh:001 | State Building Code | permit-application vesting | code in effect when permit application is received | receipt of building-permit application by governing authority | yes, for duration of covered permit work | src:usa-nh:rsa-155-a-2 | partially_verified |
| date-rule:usa-nh:002 | State Building Code | concurrency period | 6 months after effective date of code adopted under RSA 155-A:1, IV | effective date of new State Building Code adoption | yes; applicant may use prior code or new code, but not both | src:usa-nh:rsa-155-a-2 | partially_verified |
| date-rule:usa-nh:003 | State Building Code 2024 update | future mandatory permit-request date | 2027-01-01 | official state page says permit requests shall follow 2024 codes beginning this date | likely yes from 2026-07-01 through 2026-12-31 under concurrency model; recheck after effective date | src:usa-nh:dfs-state-building-code; src:usa-nh:rsa-155-a-2 | watch |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building / Residential / Existing / Mechanical / Plumbing / Energy / Electrical / ISPSC | 2024 code cycle per official state building-code page | null | unresolved | 2026-07-01 | 2026-07-01 | 2027-01-01 | active | src:usa-nh:dfs-state-building-code | Official state page snippet states "effective July 1, 2026" and "Beginning January 1, 2027, all permit requests shall follow the codes below." Full code-family list needs direct page/PDF extraction after access is available. |
| Fire | Possible 2024 NFPA 1 / NFPA 101 update through 2026 legislation | 2026-02-05 amendment text located in nonofficial bill source | unresolved | unresolved | unresolved | unresolved | active | src:usa-nh:sb494-2026-watch | Treat as a watch item only; not current law in this report. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-nh:001 | Building and Fire | conflict between State Building Code and State Fire Code | conflict in applicable requirements | Code creating the greater degree of life safety takes precedence, subject to RSA 155-A:10 review. | src:usa-nh:rsa-155-a-2 | partially_verified |
| applicability-rule:usa-nh:002 | Building | state-owned, community-college-system, and university-system buildings | permit, inspection, or certificate of occupancy for covered state/system buildings | State Fire Marshal issues permits, conducts inspections, and issues certificates of occupancy for covered buildings unless contracting/authorization applies. | src:usa-nh:rsa-155-a-2 | partially_verified |
| applicability-rule:usa-nh:003 | Building / Electrical | AFCI requirements | NEC provisions exceeding 2014 NEC AFCI requirements | RSA 155-A:2, XI says NEC AFCI provisions exceeding 2014 NEC requirements shall not be enforced under State Building Code or chapter. | src:usa-nh:rsa-155-a-2 | partially_verified |
| applicability-rule:usa-nh:004 | Fire / Residential sprinklers | detached one- or two-family dwellings and certain residential scopes | state or local sprinkler requirement | RSA 153:5 limits State Fire Code sprinkler requirements for detached one- or two-family dwelling units; RSA 155-A:3 also limits municipal enforcement of sprinkler mandates for specified residential dwellings. | src:usa-nh:rsa-153-5; src:usa-nh:rsa-155-a-3 | partially_verified |
| applicability-rule:usa-nh:005 | Fuel Gas | fuel gas systems | fuel-gas provisions in construction codes | Amendment materials indicate IFGC is not adopted in RSA 155-A:1, IV and fuel gas systems comply with the New Hampshire Fire Code. Full amendment text should be parsed before production use. | src:usa-nh:summary-2021-building-amendments-2025 | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** state-board-reviewed amendments with legislative ratification for the State Building Code; State Board of Fire Control / State Fire Marshal amendment process for State Fire Code.

**Where amendments are published:** NH Division of Fire Safety / Fire Marshal building-code and fire-code pages; amendment summary PDFs; board amendment/exhibit pages. Official state PDF access was blocked in automated fetch during this pass, so amendment details are only partially extracted from snippets and statutory text.

**Amendment parsing status:** partial

### 5.2 State Amendment Sources

| Amendment Source ID | Code Family | Publication | Effective Date | Parsing Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| amend:usa-nh:building-2021-2025 | Building, residential, existing, mechanical, plumbing, energy, electrical | New Hampshire Building Code Amendments Effective 2025-07-01 | 2025-07-01 | partial | src:usa-nh:summary-2021-building-amendments-2025 | Search/PDF snippets confirm 2023 NEC update and identify IFGC/IFC non-adoption treatment; full PDF extraction needed. |
| amend:usa-nh:fire-2024 | Fire | State Fire Code page and Saf-FMO 300 / Board of Fire Control amendments | 2024-08-13 | partial | src:usa-nh:dfs-state-fire-code; src:usa-nh:rsa-153-1; src:usa-nh:rsa-153-5 | Full Saf-FMO 300 amendment text not parsed in this pass. |
| amend:usa-nh:building-2026-future | Building and related code families | Official state building-code page / future 2024-code materials | 2026-07-01 | watch | src:usa-nh:dfs-state-building-code | Recheck when direct PDF/page extraction is available and after 2026-07-01. |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Impact Area | Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| amendment:usa-nh:ifc-not-adopted | Building / Fire coordination | referenced fire code | Amendment snippets state the International Fire Code is not adopted by RSA 155-A:1, IV; references to IFC are deemed references to the New Hampshire Fire Code. | src:usa-nh:summary-2021-building-amendments-2025 | partially_verified |
| amendment:usa-nh:ifgc-not-adopted | Fuel Gas | fuel gas systems | Amendment snippets state the International Fuel Gas Code is not adopted by RSA 155-A:1, IV and fuel gas systems must comply with the New Hampshire Fire Code. | src:usa-nh:summary-2021-building-amendments-2025 | partially_verified |
| amendment:usa-nh:nec-afci-limitation | Electrical | AFCI enforcement | RSA 155-A:2, XI limits enforcement of NEC AFCI requirements exceeding 2014 NEC requirements. | src:usa-nh:rsa-155-a-2 | partially_verified |
| amendment:usa-nh:one-stair | Building / residential | egress | RSA 155-A:2, XII permits residential buildings up to four floors above grade plane to have only one stairway under conditions established by the BCRB. | src:usa-nh:rsa-155-a-2 | partially_verified |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-nh"
  model: "statewide_code_local_enforcement_with_state_fire_marshal_backstop"
  enforcing_entities:
    - "local enforcement agency for State Building Code where municipality has an enforcement mechanism"
    - "local fire chief for State Fire Code"
    - "State Fire Marshal or designee for municipalities without local enforcement agency or qualified third-party contract, upon written municipal request"
    - "State Fire Marshal for state-owned, community college system, and university system buildings unless contracted or authorized otherwise"
  required_officials:
    - "building official / local enforcement agency where local enforcement mechanism exists"
    - "local fire chief"
  state_reserved_activities:
    - "State Fire Marshal permits, inspections, and certificates of occupancy for state-owned, community college system, and university system buildings"
    - "Backstop enforcement where no local enforcement agency or qualified third-party contract exists and municipality requests assistance"
  source_ids:
    - "src:usa-nh:rsa-155-a-2"
    - "src:usa-nh:rsa-155-a-3"
    - "src:usa-nh:rsa-155-a-7"
  verification_status: "partially_verified"
  confidence: 0.85
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-nh"
  model: "limited_local_administrative_amendments_with_board_confirmation; technical_amendments_largely_preempted_after_2026_07_01"
  applies_to_code_families:
    - "State Building Code"
    - "State Fire Code, separately under RSA 153:5"
  approval_required: true
  approving_authority_id: "ahj:usa-nh:bcrb for building-code ordinances/amendments; fire-code review path under RSA 153:5 / RSA 153:4-a for fire-code administrative amendments"
  filing_required: true
  registry_exists: "unresolved"
  registry_source_ids:
    - "src:usa-nh:bcrb-page"
  legal_basis_source_ids:
    - "src:usa-nh:rsa-155-a-3"
    - "src:usa-nh:rsa-155-a-10"
    - "src:usa-nh:rsa-153-5"
  verification_status: "partially_verified"
  confidence: 0.75
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement does not equal local code-writing authority. RSA 155-A makes the State Building Code effective statewide and gives local enforcement agencies a code-administration role. Local amendments are a narrower legal question: municipalities may use local enforcement mechanisms and administrative provisions, but technical building-code amendments are restricted under the 2026-07-01 version of RSA 155-A:3. Municipal building-code amendments/ordinances require submission to the State Building Code Review Board for confirmation before enforcement under the statutory process.

### 6.4 Known Local Amendment Registries

| Registry ID | Registry Name | Hosting Authority | Coverage | Source IDs | Status | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| registry:usa-nh:bcrb-local-amendments | BCRB local amendment / board materials | State Building Code Review Board / Division of Fire Safety | unresolved | src:usa-nh:bcrb-page | unresolved | Search snippets and statute confirm BCRB review/confirmation role; a complete registry export was not extracted. |

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendments were parsed into normalized records in this pass. This is an explicit unresolved item, not a finding that no local amendments exist.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: partial_statutory_model_only

Jurisdiction stack:

```text
Address
  -> State of New Hampshire
  -> County
  -> Municipality / unincorporated county or unorganized place
  -> Village district or special district, if applicable
  -> Local enforcement agency / building official, if municipality has enforcement mechanism
  -> Local fire chief
  -> State Fire Marshal backstop or reserved state-building role, if applicable
  -> Trade-specific licensing boards for licensed work/program issues
  -> Applicable State Building Code adoption records
  -> Applicable State Fire Code adoption records
  -> Confirmed local administrative amendments, if any
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | not selected | none | statewide | unknown | pending |
| County | not selected | none | statewide | unknown | pending |
| Municipality | not selected | none | statewide | unknown | pending |
| Village District / Special District | not selected | none | partial | unknown | pending |
| Fire District | not selected | none | statewide | unknown | pending |

### 7.3 AHJ Contact Data

No AHJ contact data was populated for this pass. The statutory model is captured, but production address-level AHJ routing requires current municipal, fire, and state-agency contact datasets.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Publisher / Custodian | Source Type | URL / Locator | Date / Version | Used For | Access Notes |
| --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-nh:dfs-state-building-code | State Building Code | New Hampshire Division of Fire Safety / State Fire Marshal | official agency page | https://www.firemarshal.dos.nh.gov/laws-rules-regulatory/state-building-code | accessed 2026-06-26 | statewide building-code page, future 2024-code transition watch, state-page caveat | Search result accessible; direct automated page open returned 403. |
| src:usa-nh:bcrb-page | State Building Code Review Board | New Hampshire Division of Fire Safety / State Fire Marshal | official agency page | https://www.firemarshal.dos.nh.gov/laws-rules-regulatory/boards-commissions-committees/state-building-code-review-board | accessed 2026-06-26 | BCRB role and amendment review source locator | Search result accessible; direct automated page open returned 403. |
| src:usa-nh:rsa-155-a-1 | RSA 155-A:1 Definitions | New Hampshire General Court / Revised Statutes Online | official statute, mirrored through nonofficial text extraction | https://gc.nh.gov/rsa/html/XII/155-A/155-A-1.htm | 2025 statute text, including future notes | State Building Code definition and current model-code editions | Official locator identified; extractable text came from Justia mirror due General Court navigation/403 issues. |
| src:usa-nh:rsa-155-a-2 | RSA 155-A:2 State Building Code | New Hampshire General Court / Revised Statutes Online | official statute, mirrored through nonofficial text extraction | https://gc.nh.gov/rsa/html/XII/155-A/155-A-2.htm | 2025 statute text, including future notes | compliance mandate, permit-date rule, concurrency period, conflict rule, State Fire Marshal reserved role | Official locator identified; extractable text came from Justia mirror. |
| src:usa-nh:rsa-155-a-3 | RSA 155-A:3 Enforcement Mechanism / Local Amendments | New Hampshire General Court / Revised Statutes Online | official statute, mirrored through nonofficial text extraction | https://gc.nh.gov/rsa/html/XII/155-A/155-A-3.htm | 2025 statute text, future version effective 2026-07-01 | local enforcement mechanism and local amendment constraints | Official locator identified; extractable text came from Justia mirror. |
| src:usa-nh:rsa-155-a-7 | RSA 155-A:7 Enforcement Authority | New Hampshire General Court / Revised Statutes Online | official statute, mirrored through nonofficial text extraction | https://gc.nh.gov/rsa/html/XII/155-A/155-A-7.htm | 2025 statute text, including future notes | local enforcement agency, local fire chief, State Fire Marshal backstop | Official locator identified; extractable text came from Justia mirror. |
| src:usa-nh:rsa-155-a-10 | RSA 155-A:10 State Building Code Review Board | New Hampshire General Court / Revised Statutes Online | official statute, mirrored through nonofficial text extraction | https://gc.nh.gov/rsa/html/XII/155-A/155-A-10.htm | 2025 statute text | BCRB role; municipal amendment confirmation | Extractable lines not fully opened in this pass; search result and related statutory cross-references used. |
| src:usa-nh:rsa-153-1 | RSA 153:1 Definitions | New Hampshire General Court / Revised Statutes Online | official statute, mirrored through nonofficial text extraction | https://gc.nh.gov/rsa/html/XII/153/153-1.htm | 2025 statute text | State Fire Code definition; NFPA 101 / NFPA 1 editions | Official locator identified; extractable text came from Justia mirror. |
| src:usa-nh:rsa-153-5 | RSA 153:5 State Fire Code; Rules | New Hampshire General Court / Revised Statutes Online | official statute, mirrored through nonofficial text extraction | https://gc.nh.gov/rsa/html/XII/153/153-5.htm | 2025 statute text, including future notes | State Fire Code amendment authority; local fire chief / State Fire Marshal use; local fire-code amendment limits | Official locator identified; extractable text came from Justia mirror. |
| src:usa-nh:dfs-state-fire-code | State Fire Code | New Hampshire Division of Fire Safety / State Fire Marshal | official agency page | https://www.firemarshal.dos.nh.gov/laws-rules-regulatory/state-fire-code | accessed 2026-06-26 | 2021 NFPA 1 / NFPA 101 effective 2024-08-13 agency confirmation | Search result accessible; direct automated page open returned 403. |
| src:usa-nh:summary-2021-building-amendments-2025 | New Hampshire Building Code Amendments Effective 2025-07-01 | New Hampshire Division of Fire Safety / State Fire Marshal | official PDF | https://mm.nh.gov/files/uploads/fmo/remote-docs/summary-of-2021-building-code-amendments-effective-1jul2025.pdf | 2025-07-01 | amendment publication path; 2023 NEC update; IFGC/IFC non-adoption snippets | Search/PDF snippets accessible; full automated PDF open blocked. |
| src:usa-nh:oplc-mechanical-board | Mechanical Safety and Licensing Board Laws and Rules | New Hampshire Office of Professional Licensure and Certification | official agency page | https://www.oplc.nh.gov/mechanical-safety-and-licensing-board-laws-and-rules | accessed 2026-06-26 | mechanical/plumbing/fuel-gas licensing program source locator | Search result accessible. |
| src:usa-nh:dol-elevator | Elevators and Accessibility Lifts | New Hampshire Department of Labor | official agency page | https://www.dol.nh.gov/inspections/elevators-and-accessibility-lifts | accessed 2026-06-26 | elevator/accessibility lift authority and rule locator | Search result accessible. |
| src:usa-nh:rsa-155-a-5 | RSA 155-A:5 through 155-A:5-b Accessibility Standards for Public Buildings | New Hampshire General Court / Revised Statutes Online | official statute, mirrored through nonofficial text extraction | https://gc.nh.gov/rsa/html/XII/155-A/155-A-5.htm | 2025 statute text | accessibility authority and certification/enforcement framework | Extractable text came from Justia search/open snippets. |
| src:usa-nh:hb134-2025 | HB 134 / Chapter 135, relative to the state building code | New Hampshire General Court | official enacted bill / chapter law locator | https://www.gencourt.state.nh.us/bill_status/billinfo.aspx?id=134 | signed 2025-06-23; effective 2025-07-01 | 2023 NEC and amendment-ratification-date update | Official bill locator should be verified directly; search snippets from FastDemocracy/BillTrack50 used as temporary extraction aid. |
| src:usa-nh:sb494-2026-watch | SB 494, relative to the state fire code, fire incident reporting and investigations, and duties of the state fire marshal | New Hampshire General Court | bill watch locator | https://www.gencourt.state.nh.us/bill_status/billinfo.aspx?id=494 | 2026 session | watch item for possible 2024 NFPA 1 / NFPA 101 update | Noncurrent-law watch item; not used for current adoption. |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-nh:dfs-state-building-code | automated_access_blocked | Official NH Fire Marshal page was discoverable in search but returned 403 when opened through automated browsing. | Re-open manually or through an approved state-site fetch before verified status. |
| src:usa-nh:dfs-state-fire-code | automated_access_blocked | Official NH Fire Marshal fire-code page was discoverable in search but returned 403 when opened through automated browsing. | Re-open manually or through an approved state-site fetch before verified status. |
| src:usa-nh:rsa-155-a-1 | unofficial_text_mirror_used | Official General Court locator was identified, but line-level extraction used a nonofficial Justia mirror. | Validate against official RSA text before verified status. |
| src:usa-nh:rsa-155-a-2 | unofficial_text_mirror_used | Official General Court locator was identified, but line-level extraction used a nonofficial Justia mirror. | Validate against official RSA text before verified status. |
| src:usa-nh:rsa-155-a-3 | unofficial_text_mirror_used | Official General Court locator was identified, but line-level extraction used a nonofficial Justia mirror. | Validate against official RSA text before verified status. |
| src:usa-nh:rsa-155-a-7 | unofficial_text_mirror_used | Official General Court locator was identified, but line-level extraction used a nonofficial Justia mirror. | Validate against official RSA text before verified status. |
| src:usa-nh:rsa-153-1 | unofficial_text_mirror_used | Official General Court locator was identified, but line-level extraction used a nonofficial Justia mirror. | Validate against official RSA text before verified status. |
| src:usa-nh:rsa-153-5 | unofficial_text_mirror_used | Official General Court locator was identified, but line-level extraction used a nonofficial Justia mirror. | Validate against official RSA text before verified status. |
| src:usa-nh:summary-2021-building-amendments-2025 | partial_pdf_extraction | Official PDF snippets were available, but full PDF text was not parsed. | Full PDF extraction required before using amendment-level details in production. |
| src:usa-nh:hb134-2025 | indirect_bill_metadata | Bill signing/effective-date details were supported by bill-tracking snippets; official bill/chapter page should be opened directly. | Treat adoption/effective date as high confidence but not final verified until official bill record is archived. |
| src:usa-nh:sb494-2026-watch | noncurrent_law | Pending/current-session bill source is used only as a watch item. | Do not treat as current adopted fire code unless enacted and codified. |

### 8.3 Supplemental Sources

| Source ID | Title | Publisher | Source Type | URL / Locator | Used For | Caveat |
| --- | --- | --- | --- | --- | --- | --- |
| supp:usa-nh:justia-155-a | New Hampshire Revised Statutes, Chapter 155-A mirror | Justia | nonofficial statute mirror | https://law.justia.com/codes/new-hampshire/title-xii/chapter-155-a/ | line-level extraction of RSA 155-A content | Must validate against official General Court RSA text. |
| supp:usa-nh:justia-153 | New Hampshire Revised Statutes, Chapter 153 mirror | Justia | nonofficial statute mirror | https://law.justia.com/codes/new-hampshire/title-xii/chapter-153/ | line-level extraction of RSA 153 content | Must validate against official General Court RSA text. |
| supp:usa-nh:fastdemocracy-hb1059 | HB 1059 bill page | FastDemocracy | bill-tracking mirror | nonofficial bill tracker | 2024 adoption-date/effective-date context | Not a substitute for General Court bill/chapter record. |
| supp:usa-nh:nhma-2025-guide | Changes to Building Code Laws in 2025 | New Hampshire Municipal Association | municipal guidance PDF | NHMA guidance PDF | local amendment transition context | Supplemental interpretation only. |

### 8.4 Source Extraction Metadata

| Extraction ID | Source IDs | Extraction Method | Extracted On | Extractor | Notes |
| --- | --- | --- | --- | --- | --- |
| extract:usa-nh:001 | src:usa-nh:rsa-155-a-1; src:usa-nh:rsa-155-a-2; src:usa-nh:rsa-155-a-3; src:usa-nh:rsa-155-a-7 | web search/open of official locator plus nonofficial mirrored statute text | 2026-06-26 | ChatGPT | Official General Court URLs should be revalidated manually due access/navigation issues. |
| extract:usa-nh:002 | src:usa-nh:rsa-153-1; src:usa-nh:rsa-153-5 | web search/open of official locator plus nonofficial mirrored statute text | 2026-06-26 | ChatGPT | Fire-code statute definitions and amendment framework extracted. |
| extract:usa-nh:003 | src:usa-nh:dfs-state-building-code; src:usa-nh:dfs-state-fire-code; src:usa-nh:bcrb-page | web search snippets from official NH Fire Marshal pages | 2026-06-26 | ChatGPT | Direct open returned 403; snippets used only for page role and transition/watch facts. |
| extract:usa-nh:004 | src:usa-nh:summary-2021-building-amendments-2025 | web search/PDF snippets | 2026-06-26 | ChatGPT | Amendment details need full PDF extraction. |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 0.90 | none | Body contains source-backed core authority and adoption fields, with caveats. |
| report | risk.overall_confidence | 0.72 | verified | 0.80 | none | Confidence limited by blocked official pages and incomplete amendment parsing. |
| ahj:usa-nh:bcrb | authority name / role | State Building Code Review Board; building-code amendment and confirmation role | partially_verified | 0.85 | src:usa-nh:rsa-155-a-1; src:usa-nh:rsa-155-a-10; src:usa-nh:bcrb-page | Direct official statute revalidation required. |
| adoption:usa-nh:building-2021-2025 | code edition | 2021 IBC | partially_verified | 0.90 | src:usa-nh:rsa-155-a-1 | Current as of 2026-06-26; future 2026-07-01 transition pending. |
| adoption:usa-nh:electrical-2023-2025 | code edition | 2023 NFPA 70 / NEC | partially_verified | 0.88 | src:usa-nh:rsa-155-a-1; src:usa-nh:hb134-2025 | HB 134 signing/effective metadata needs official bill-record archive. |
| adoption:usa-nh:fire-2021-2024 | code edition | 2021 NFPA 101 and 2021 NFPA 1 | partially_verified | 0.88 | src:usa-nh:rsa-153-1; src:usa-nh:dfs-state-fire-code | Possible 2026 update is watch-only. |
| date-rule:usa-nh:001 | permit-date rule | code in effect when permit application is received remains in effect for permit work | partially_verified | 0.85 | src:usa-nh:rsa-155-a-2 | Statutory text extracted from mirror. |
| date-rule:usa-nh:002 | concurrency period | six months after effective date; prior or new code, not both | partially_verified | 0.85 | src:usa-nh:rsa-155-a-2 | Statutory text extracted from mirror. |
| local-enforcement:usa-nh | model | statewide_code_local_enforcement_with_state_fire_marshal_backstop | partially_verified | 0.85 | src:usa-nh:rsa-155-a-2; src:usa-nh:rsa-155-a-7 | Address-level implementation not populated. |
| local-amendment-rule:usa-nh | model | limited_local_administrative_amendments_with_board_confirmation | partially_verified | 0.75 | src:usa-nh:rsa-155-a-3; src:usa-nh:rsa-155-a-10; src:usa-nh:rsa-153-5 | Need full registry/board process extraction. |
| ahj:usa-nh:dol-elevator | authority | Department of Labor elevator/accessibility lift program | partially_verified | 0.65 | src:usa-nh:dol-elevator | Lab 1300 details not fully parsed. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Every `src:usa-nh:...` used in body appears in section 8. |
| All authority IDs resolve | pass | Authority IDs are defined in sections 2 and 6 or are self-explanatory state/local actors. |
| All current code families have adoption records | pass | Core statewide code families have matrix entries; fuel gas/elevator remain explicitly limited. |
| Building and operational fire code are separated | pass | RSA 155-A and RSA 153 systems are separated. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Dates are separate; null/unresolved retained where not supported. |
| Effective dates are valid ISO dates | pass | Entered dates use ISO format. |
| No impossible date sequences | pass | Dates entered are plausible; future 2026-07-01 transition is flagged as pending. |
| Transition rules have explicit trigger conditions | pass | Permit-receipt and six-month concurrency triggers are captured. |
| Permit-date logic is captured where applicable | pass | RSA 155-A:2 rule captured. |
| Local enforcement model classified | pass | Classified as statewide code with local enforcement and State Fire Marshal backstop. |
| Local amendment rule classified | pass | Classified with 2026-07-01 transition caveat. |
| AHJ confirmation metadata present | fail | No address-level AHJ contacts or boundary sources populated. |
| Official-source caveats captured | pass | Automated-access and mirror-use caveats are recorded. |
| State amendments fully parsed | fail | Amendment PDFs and Saf-FMO 300 were not fully parsed. |
| Future 2024-code transition fully normalized | fail | Official page snippet supports watch item, but full code-family list and adopted instrument require direct extraction. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-nh:001 | high | official statute validation | RSA text was extracted from a nonofficial mirror after official General Court access/navigation failed. | Revalidate RSA 155-A and RSA 153 sections directly against official General Court pages or archived chapter-law PDFs. | null | null | open |
| issue:usa-nh:002 | high | 2026 building-code transition | Official state page indicates 2024 codes effective 2026-07-01 and mandatory for permit requests 2027-01-01, but full official PDF/list was not extracted. | Download/parse official 2024-code update materials and update adoption matrix after 2026-07-01. | null | null | open |
| issue:usa-nh:003 | medium | state amendments | Building-code amendment PDF and fire-code Saf-FMO 300 / Board of Fire Control amendments are only partially parsed. | Extract full amendment documents and normalize high-impact amendments. | null | null | open |
| issue:usa-nh:004 | medium | local amendment registry | BCRB local-amendment confirmation registry was not extracted. | Locate and parse current municipal amendment/ordinance confirmation list, if published. | null | null | open |
| issue:usa-nh:005 | medium | elevator/conveyance code editions | Department of Labor elevator program authority located, but detailed adopted elevator standards/rules were not parsed. | Parse RSA 157-B and Lab 1300 adopted rules. | null | null | open |
| issue:usa-nh:006 | medium | AHJ routing | No boundary, contact, or municipal enforcement mechanism dataset was populated. | Select boundary sources and build municipal/fire/AHJ contact layer. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-nh:state-building-code-page | src:usa-nh:dfs-state-building-code | html_diff | monthly | new code text, 2024-code transition updates, or portal restructuring | 2026-06-26 | active |
| watch:usa-nh:bcrb-page | src:usa-nh:bcrb-page | html_diff | monthly | new BCRB amendments, exhibits, or local amendment registry changes | 2026-06-26 | active |
| watch:usa-nh:state-fire-code-page | src:usa-nh:dfs-state-fire-code | html_diff | monthly | State Fire Code edition or Saf-FMO 300 amendment changes | 2026-06-26 | active |
| watch:usa-nh:rsa-155-a | src:usa-nh:rsa-155-a-1 | statute_diff | monthly through 2026-08-01, then quarterly | changes to State Building Code definition, enforcement, or amendment provisions | 2026-06-26 | active |
| watch:usa-nh:rsa-153 | src:usa-nh:rsa-153-1 | statute_diff | monthly through 2026 session close, then quarterly | changes to State Fire Code definition or fire-code amendment process | 2026-06-26 | active |
| watch:usa-nh:sb494-2026 | src:usa-nh:sb494-2026-watch | bill_status | weekly during session | enactment or codification of 2024 NFPA 1 / NFPA 101 update | 2026-06-26 | active |
| watch:usa-nh:dol-elevator | src:usa-nh:dol-elevator | rule_diff | quarterly | updates to elevator/accessibility lift law or Lab 1300 rules | 2026-06-26 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-24 | Generated baseline draft report | report:usa-nh | none | Codex | Placeholder values removed; state-specific research still needed. |
| 2026-06-26 | Populated New Hampshire authority, adoption, transition, enforcement, amendment, source, and QA sections | report:usa-nh; ahj:usa-nh:bcrb; adoption:usa-nh:*; local-enforcement:usa-nh; local-amendment-rule:usa-nh | src:usa-nh:rsa-155-a-1; src:usa-nh:rsa-155-a-2; src:usa-nh:rsa-155-a-3; src:usa-nh:rsa-155-a-7; src:usa-nh:rsa-153-1; src:usa-nh:rsa-153-5; src:usa-nh:dfs-state-building-code; src:usa-nh:dfs-state-fire-code | ChatGPT | Upgraded to partially_verified with explicit caveats; amendments and future transition remain open issues. |
