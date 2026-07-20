---
state:
  state_id: "US-WV"
  name: "West Virginia"
  abbreviation: "WV"
report:
  report_id: "state-report:usa-wv"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-26"
  last_verified: "2026-06-26"
  reviewed_by: null
risk:
  overall_confidence: 0.68 # 0.00 - 1.00
  risk_flags:
    - "fire_code_emergency_to_final_transition_requires_current_date_review"
    - "local_amendment_scope_partially_classified"
    - "state_amendments_not_fully_parsed"
    - "elevator_conveyance_authority_unresolved"
    - "ahj_contact_data_not_populated"
  open_questions_count: 5

---

# State Building Code Authority Report: West Virginia

## 1. Executive Summary

- **Authority model:** West Virginia uses a hybrid model. The State Fire Commission is the statewide rulemaking body for both the State Building Code and State Fire Code, while the Office of the State Fire Marshal holds statewide fire/life-safety enforcement responsibilities and state-level fire/building-code violation enforcement powers. The State Building Code has force and effect in counties and municipalities that adopt it; the State Fire Code is promulgated as a statewide fire-safety code.

- **Statewide code status:** 87 CSR 4, the State Building Code, is active with an effective date of 2022-08-01 and a sunset date of 2027-08-01. It incorporates a 2018 ICC-based building-code package, the 2015 IECC for residential energy provisions, ASHRAE/IESNA 90.1-2013 for commercial energy provisions, the 2020 National Electrical Code, ICC/ANSI A117.1-2017, and the 2018 International Swimming Pool and Spa Code, subject to West Virginia-specific amendments and exclusions.

- **Local enforcement model:** Local jurisdictions that adopt the State Building Code are responsible for enforcement, must notify the State Fire Commission or State Fire Marshal within 30 days of adoption, and must comply with State Fire Marshal certification rules for public-sector building-code officials, inspectors, and plans examiners. 87 CSR 4 also delegates building-code interpretation and enforcement questions to the local jurisdiction unless state law, rule, or incorporated code provides otherwise.

- **Local amendment posture:** Existing local building codes are voided after State Building Code promulgation, and counties or municipalities that choose to adopt a building code must adopt the State Building Code. West Virginia statutes preserve more stringent local laws, ordinances, or regulations only when they are not inconsistent with the State Building Code and are not contrary to recognized engineering standards; the State Fire Commission decides priority when questions arise. The available sources do not establish a central public registry of local amendments.

- **Known transition periods or pending changes:** 87 CSR 1, the State Fire Code, is shown by the Secretary of State as active emergency while the legislative rule proceeds, with an emergency effective date of 2025-03-06 and a final legislative rule effective date of 2026-07-01. The 2026 final Fire Code rule terminates on 2031-08-01 unless renewed or replaced. 87 CSR 4 terminates on 2027-08-01 unless renewed or replaced.

- **Production readiness:** limited_use. The core authority model, building-code adoption package, fire-code status, and local-adoption model are source-backed. The report should not be marked verified until full amendment parsing, AHJ contacts, local amendment registry research, and conveyance/elevator authority review are completed.

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: The State Fire Commission proposes legislative rules for the State Building
    Code and State Fire Code.
  confidence: 0.9
  source_ids:
  - src:usa-wv:statute:15a-11-3
  - src:usa-wv:statute:15a-11-5
  - src:usa-wv:rule:87-04-final-2022
  - src:usa-wv:rule:87-01-final-2026
- topic: Primary building code edition
  finding: 87 CSR 4 adopts a 2018 ICC-family package with West Virginia amendments,
    effective 2022-08-01.
  confidence: 0.86
  source_ids:
  - src:usa-wv:sos:87-04-rule-page
  - src:usa-wv:rule:87-04-final-2022
- topic: Fire code edition/status
  finding: 87 CSR 1 is active emergency as of this report date; the final legislative
    rule is effective 2026-07-01 and adopts NFPA 1-2021 and NFPA 101-2021.
  confidence: 0.8
  source_ids:
  - src:usa-wv:sos:87-01-rule-page
  - src:usa-wv:rule:87-01-emergency-2025
  - src:usa-wv:rule:87-01-final-2026
- topic: Electrical code authority
  finding: The NEC is incorporated through 87 CSR 4; electrician licensing and electrical-inspector
    certification are administered through State Fire Marshal statutes.
  confidence: 0.78
  source_ids:
  - src:usa-wv:rule:87-04-final-2022
  - src:usa-wv:statute:29-3b-2
  - src:usa-wv:statute:29-3b-5
  - src:usa-wv:statute:29-3c-3
  - src:usa-wv:statute:29-3c-4
- topic: Local enforcement
  finding: Counties and municipalities may adopt the State Building Code; adopting
    local jurisdictions enforce it and must notify the State Fire Commission or State
    Fire Marshal.
  confidence: 0.84
  source_ids:
  - src:usa-wv:statute:15a-11-5
  - src:usa-wv:statute:7-1-3n
  - src:usa-wv:statute:8-12-13
  - src:usa-wv:rule:87-04-final-2022
- topic: Local amendments
  finding: Local rules more stringent than the State Building Code may govern only
    if not inconsistent and not contrary to recognized engineering standards; Commission
    priority review is noted.
  confidence: 0.65
  source_ids:
  - src:usa-wv:statute:15a-11-5
- topic: Effective / operative date rule
  finding: 87 CSR 4 has a fixed effective date of 2022-08-01; no permit-date grace
    period was extracted from the reviewed State Building Code rule.
  confidence: 0.75
  source_ids:
  - src:usa-wv:sos:87-04-rule-page
  - src:usa-wv:rule:87-04-final-2022
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-wv:state-fire-commission |
| Authority name | West Virginia State Fire Commission |
| Authority type | state_commission_rulemaking_body |
| Legal basis | W. Va. Code §15A-11-5; 87 CSR 4 |
| Role | Proposes legislative rules for the State Building Code, including building energy codes, to safeguard life and property and ensure construction quality. |
| Enforcement model | State Building Code has force and effect in adopting counties and municipalities; local jurisdictions enforce adopted building-code rules, while the Office of the State Fire Marshal has state-level enforcement authority for fire/building-code violations and related inspections. |
| Source IDs | src:usa-wv:statute:15a-11-5; src:usa-wv:statute:15a-10-2; src:usa-wv:statute:15a-10-3; src:usa-wv:rule:87-04-final-2022 |
| Verification status | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-wv:state-fire-commission | West Virginia State Fire Commission | Adopts State Building Code by legislative rule; local adoption gives local force/effect. | W. Va. Code §15A-11-5; 87 CSR 4 | src:usa-wv:statute:15a-11-5; src:usa-wv:rule:87-04-final-2022 | partially_verified |
| Residential | ahj:usa-wv:state-fire-commission | West Virginia State Fire Commission | Incorporates the 2018 IRC with WV amendments and energy-code exclusions. | 87 CSR 4 | src:usa-wv:rule:87-04-final-2022 | partially_verified |
| Existing Building / Rehabilitation | ahj:usa-wv:state-fire-commission | West Virginia State Fire Commission | Incorporates the 2018 International Existing Building Code with WV amendments. | 87 CSR 4 | src:usa-wv:rule:87-04-final-2022 | partially_verified |
| Mechanical | ahj:usa-wv:state-fire-commission | West Virginia State Fire Commission | Incorporates the 2018 International Mechanical Code. | 87 CSR 4 | src:usa-wv:rule:87-04-final-2022 | partially_verified |
| Plumbing | ahj:usa-wv:state-fire-commission | West Virginia State Fire Commission | Incorporates the 2018 International Plumbing Code; Department of Health and Human Resources plumbing rules take precedence in conflicts noted by 87 CSR 4. | 87 CSR 4 | src:usa-wv:rule:87-04-final-2022 | partially_verified |
| Fuel Gas | ahj:usa-wv:state-fire-commission | West Virginia State Fire Commission | Incorporates the 2018 International Fuel Gas Code. | 87 CSR 4 | src:usa-wv:rule:87-04-final-2022 | partially_verified |
| Electrical | ahj:usa-wv:state-fire-commission; ahj:usa-wv:state-fire-marshal-electrical | State Fire Commission; Office of State Fire Marshal | State Building Code incorporates the 2020 NEC; State Fire Marshal administers electrician licensing and electrical-inspector certification. | 87 CSR 4; W. Va. Code ch. 29, arts. 3B and 3C | src:usa-wv:rule:87-04-final-2022; src:usa-wv:statute:29-3b-2; src:usa-wv:statute:29-3b-5; src:usa-wv:statute:29-3c-3; src:usa-wv:statute:29-3c-4 | partially_verified |
| Energy | ahj:usa-wv:state-fire-commission | West Virginia State Fire Commission | State Building Code includes energy codes; 87 CSR 4 uses 2015 IECC residential and ASHRAE/IESNA 90.1-2013 commercial provisions. | W. Va. Code §15A-11-5; 87 CSR 4 | src:usa-wv:statute:15a-11-5; src:usa-wv:rule:87-04-final-2022 | partially_verified |
| Fire - construction references | ahj:usa-wv:state-fire-commission | West Virginia State Fire Commission | 87 CSR 4 states State Fire Code takes precedence and substitutes NFPA Life Safety Code for International Fire Code references in incorporated ICC codes. | 87 CSR 4 | src:usa-wv:rule:87-04-final-2022 | partially_verified |
| Fire - operational / prevention code | ahj:usa-wv:state-fire-commission; ahj:usa-wv:state-fire-marshal | State Fire Commission; Office of State Fire Marshal | State Fire Commission promulgates the State Fire Code; State Fire Marshal enforces fire programs and state fire/life-safety laws. | W. Va. Code §§15A-11-3, 15A-10-2, 15A-10-3; 87 CSR 1 | src:usa-wv:statute:15a-11-3; src:usa-wv:statute:15a-10-2; src:usa-wv:statute:15a-10-3; src:usa-wv:rule:87-01-emergency-2025; src:usa-wv:rule:87-01-final-2026 | partially_verified |
| Accessibility | ahj:usa-wv:state-fire-commission | West Virginia State Fire Commission | Incorporates ICC/ANSI A117.1-2017 as an accessibility standard within the State Building Code package. | 87 CSR 4 | src:usa-wv:rule:87-04-final-2022 | partially_verified |
| Elevator / Conveyance | ahj:usa-wv:unresolved-conveyance | Unresolved | Conveyance/elevator authority was not researched to a primary-source conclusion in this pass. | unresolved | none | unresolved |

### 2.3 Authority Hierarchy Notes

The State Fire Commission is the rulemaking body for the State Building Code and State Fire Code. The Office of the State Fire Marshal assists the State Fire Commission and has statutory enforcement responsibilities for fire programs and authority to enforce the fire code, building code, and related laws. Counties and municipalities may adopt the State Building Code; once they do, local enforcement is the default for building-code administration. 87 CSR 4 states that building-code interpretation and enforcement questions are delegated to the local jurisdiction unless the State Code, the rule, or the incorporated model codes provide otherwise. The State Fire Code remains a separate statewide fire/life-safety regime with express precedence over the State Building Code where conflicts arise.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-wv:001 | ahj:usa-wv:state-fire-commission | promulgates | State Building Code and State Fire Code legislative rules | src:usa-wv:statute:15a-11-3; src:usa-wv:statute:15a-11-5 | partially_verified |
| edge:usa-wv:002 | ahj:usa-wv:state-fire-marshal | enforces | fire programs, fire code, building code, and related laws within statutory scope | src:usa-wv:statute:15a-10-2; src:usa-wv:statute:15a-10-3 | partially_verified |
| edge:usa-wv:003 | ahj:usa-wv:state-fire-commission | gives_force_effect_to | State Building Code in adopting counties and municipalities | src:usa-wv:statute:15a-11-5; src:usa-wv:statute:7-1-3n; src:usa-wv:statute:8-12-13 | partially_verified |
| edge:usa-wv:004 | ahj:usa-wv:local-adopting-jurisdiction | enforces | adopted State Building Code within local jurisdiction | src:usa-wv:statute:15a-11-5; src:usa-wv:rule:87-04-final-2022 | partially_verified |
| edge:usa-wv:005 | ahj:usa-wv:state-fire-commission | resolves_priority_for | more stringent local laws or ordinances that may conflict with State Building Code | src:usa-wv:statute:15a-11-5 | partially_verified |
| edge:usa-wv:006 | ahj:usa-wv:state-fire-code | takes_precedence_over | State Building Code and incorporated model-code fire references | src:usa-wv:rule:87-04-final-2022; src:usa-wv:rule:87-01-final-2026 | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | West Virginia State Building Code | International Building Code | 2018 | active_for_adopting_jurisdictions | 2022-05-05 | 2022-08-01 | 2022-08-01 | 2022-08-01 for adopting jurisdictions | No permit-date grace period extracted; rule is prospective through adopting counties/municipalities. | src:usa-wv:sos:87-04-rule-page; src:usa-wv:rule:87-04-final-2022 |
| Residential | West Virginia State Building Code | International Residential Code | 2018 | active_for_adopting_jurisdictions | 2022-05-05 | 2022-08-01 | 2022-08-01 | 2022-08-01 for adopting jurisdictions | IRC Chapter 11 energy-efficiency provisions are exempted by 87 CSR 4; use WV energy provisions listed separately. | src:usa-wv:rule:87-04-final-2022 |
| Existing Building / Rehabilitation | West Virginia State Building Code | International Existing Building Code | 2018 | active_for_adopting_jurisdictions | 2022-05-05 | 2022-08-01 | 2022-08-01 | 2022-08-01 for adopting jurisdictions | 87 CSR 4 modifies fire-code cross-references; no permit-date grace period extracted. | src:usa-wv:rule:87-04-final-2022 |
| Mechanical | West Virginia State Building Code | International Mechanical Code | 2018 | active_for_adopting_jurisdictions | 2022-05-05 | 2022-08-01 | 2022-08-01 | 2022-08-01 for adopting jurisdictions | No separate mechanical transition rule extracted. | src:usa-wv:rule:87-04-final-2022 |
| Plumbing | West Virginia State Building Code | International Plumbing Code | 2018 | active_for_adopting_jurisdictions | 2022-05-05 | 2022-08-01 | 2022-08-01 | 2022-08-01 for adopting jurisdictions | 87 CSR 4 states Department of Health and Human Resources plumbing rules take precedence in conflicts. | src:usa-wv:rule:87-04-final-2022 |
| Fuel Gas | West Virginia State Building Code | International Fuel Gas Code | 2018 | active_for_adopting_jurisdictions | 2022-05-05 | 2022-08-01 | 2022-08-01 | 2022-08-01 for adopting jurisdictions | No separate fuel-gas transition rule extracted. | src:usa-wv:rule:87-04-final-2022 |
| Electrical | West Virginia State Building Code | NFPA 70, National Electrical Code | 2020 | active_for_adopting_jurisdictions | 2022-05-05 | 2022-08-01 | 2022-08-01 | 2022-08-01 for adopting jurisdictions | 87 CSR 4 includes NEC-specific WV amendments; licensing/certification statutes are separate. | src:usa-wv:rule:87-04-final-2022; src:usa-wv:statute:29-3b-2; src:usa-wv:statute:29-3c-3 |
| Energy | West Virginia State Building Code | International Energy Conservation Code; ASHRAE/IESNA 90.1 | 2015 residential; 2013 commercial | active_for_adopting_jurisdictions | 2022-05-05 | 2022-08-01 | 2022-08-01 | 2022-08-01 for adopting jurisdictions | Residential: 2015 IECC, with 87 CSR 4 IRC Chapter 11 exemption; Commercial: ASHRAE/IESNA 90.1-2013. | src:usa-wv:statute:15a-11-5; src:usa-wv:rule:87-04-final-2022 |
| Fire - construction references | West Virginia State Building Code | NFPA Life Safety Code substituted for IFC references | 2018/2021 reference ambiguity in 87 CSR 4 | active_with_caveat | 2022-05-05 | 2022-08-01 | 2022-08-01 | 2022-08-01 for adopting jurisdictions | 87 CSR 4 contains an internal Life Safety Code edition ambiguity; see caveat and open issue. | src:usa-wv:rule:87-04-final-2022 |
| Fire - operational / prevention code | West Virginia State Fire Code | NFPA 1; NFPA 101 | 2021 | active_emergency; final_legislative_rule_effective_2026-07-01 | 2025-03-06 | 2025-03-06 | 2025-03-06 | 2025-03-06 during emergency status; 2026-07-01 under final legislative rule | SOS lists 87 CSR 1 as active emergency and states the emergency rule is effective while the legislative rule proceeds. | src:usa-wv:sos:87-01-rule-page; src:usa-wv:rule:87-01-emergency-2025; src:usa-wv:rule:87-01-final-2026 |
| Accessibility | West Virginia State Building Code | ICC/ANSI A117.1 | 2017 | active_for_adopting_jurisdictions | 2022-05-05 | 2022-08-01 | 2022-08-01 | 2022-08-01 for adopting jurisdictions | No separate accessibility transition rule extracted. | src:usa-wv:rule:87-04-final-2022 |
| Swimming Pool / Spa | West Virginia State Building Code | International Swimming Pool and Spa Code | 2018 | active_for_adopting_jurisdictions | 2022-05-05 | 2022-08-01 | 2022-08-01 | 2022-08-01 for adopting jurisdictions | No separate pool/spa transition rule extracted. | src:usa-wv:rule:87-04-final-2022 |
| Property Maintenance | West Virginia State Building Code local option | International Property Maintenance Code | 2018 | local_option_rejection_allowed | 2022-05-05 | 2022-08-01 | local adoption dependent | local adoption dependent | 87 CSR 4 states the 2018 IPMC may be rejected by a local jurisdiction; agricultural property exemptions also apply. | src:usa-wv:rule:87-04-final-2022; src:usa-wv:statute:7-1-3n |
| Elevator / Conveyance | unresolved | unresolved | unresolved | unresolved | null | null | null | null | Conveyance/elevator adoption and authority were not reviewed to conclusion. | none |

### 3.2 Adoption Records

| Adoption Record ID | Code Family | Adopting Authority | State Rule / Source | Base Code | Edition | Adoption Date | Effective Date | Operative Date | Mandatory Date | Applies To | Notes | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| adoption:usa-wv:87-04:ibc-2018 | Building | ahj:usa-wv:state-fire-commission | 87 CSR 4 | International Building Code | 2018 | 2022-05-05 | 2022-08-01 | 2022-08-01 | 2022-08-01 for adopting jurisdictions | Counties and municipalities adopting the State Building Code | Rule terminates 2027-08-01 unless renewed or replaced. | src:usa-wv:sos:87-04-rule-page; src:usa-wv:rule:87-04-final-2022 |
| adoption:usa-wv:87-04:ipc-2018 | Plumbing | ahj:usa-wv:state-fire-commission | 87 CSR 4 | International Plumbing Code | 2018 | 2022-05-05 | 2022-08-01 | 2022-08-01 | 2022-08-01 for adopting jurisdictions | Counties and municipalities adopting the State Building Code | DHHR plumbing rules take precedence in conflicts noted in 87 CSR 4. | src:usa-wv:rule:87-04-final-2022 |
| adoption:usa-wv:87-04:imc-2018 | Mechanical | ahj:usa-wv:state-fire-commission | 87 CSR 4 | International Mechanical Code | 2018 | 2022-05-05 | 2022-08-01 | 2022-08-01 | 2022-08-01 for adopting jurisdictions | Counties and municipalities adopting the State Building Code | No separate state mechanical board was resolved from reviewed sources. | src:usa-wv:rule:87-04-final-2022 |
| adoption:usa-wv:87-04:ifgc-2018 | Fuel Gas | ahj:usa-wv:state-fire-commission | 87 CSR 4 | International Fuel Gas Code | 2018 | 2022-05-05 | 2022-08-01 | 2022-08-01 | 2022-08-01 for adopting jurisdictions | Counties and municipalities adopting the State Building Code | State Fire Code conflict precedence remains applicable. | src:usa-wv:rule:87-04-final-2022 |
| adoption:usa-wv:87-04:irc-2018 | Residential | ahj:usa-wv:state-fire-commission | 87 CSR 4 | International Residential Code | 2018 | 2022-05-05 | 2022-08-01 | 2022-08-01 | 2022-08-01 for adopting jurisdictions | Counties and municipalities adopting the State Building Code | IRC Chapter 11 energy provisions exempted by WV rule. | src:usa-wv:rule:87-04-final-2022 |
| adoption:usa-wv:87-04:iecc-2015 | Energy - residential | ahj:usa-wv:state-fire-commission | 87 CSR 4 | International Energy Conservation Code | 2015 | 2022-05-05 | 2022-08-01 | 2022-08-01 | 2022-08-01 for adopting jurisdictions | Residential energy provisions in adopting jurisdictions | Kept separate from commercial ASHRAE/IESNA record. | src:usa-wv:statute:15a-11-5; src:usa-wv:rule:87-04-final-2022 |
| adoption:usa-wv:87-04:ashrae-90-1-2013 | Energy - commercial | ahj:usa-wv:state-fire-commission | 87 CSR 4 | ASHRAE/IESNA 90.1 | 2013 | 2022-05-05 | 2022-08-01 | 2022-08-01 | 2022-08-01 for adopting jurisdictions | Commercial energy provisions in adopting jurisdictions | Kept separate from residential IECC record. | src:usa-wv:statute:15a-11-5; src:usa-wv:rule:87-04-final-2022 |
| adoption:usa-wv:87-04:nec-2020 | Electrical | ahj:usa-wv:state-fire-commission | 87 CSR 4 | NFPA 70, National Electrical Code | 2020 | 2022-05-05 | 2022-08-01 | 2022-08-01 | 2022-08-01 for adopting jurisdictions | Electrical provisions in adopting jurisdictions | Separate state licensing/certification statutes still apply. | src:usa-wv:rule:87-04-final-2022; src:usa-wv:statute:29-3b-2; src:usa-wv:statute:29-3c-3 |
| adoption:usa-wv:87-04:a117-1-2017 | Accessibility | ahj:usa-wv:state-fire-commission | 87 CSR 4 | ICC/ANSI A117.1 | 2017 | 2022-05-05 | 2022-08-01 | 2022-08-01 | 2022-08-01 for adopting jurisdictions | Accessibility design provisions in adopting jurisdictions | Federal ADA obligations are outside this state-code record. | src:usa-wv:rule:87-04-final-2022 |
| adoption:usa-wv:87-04:iebc-2018 | Existing Building / Rehabilitation | ahj:usa-wv:state-fire-commission | 87 CSR 4 | International Existing Building Code | 2018 | 2022-05-05 | 2022-08-01 | 2022-08-01 | 2022-08-01 for adopting jurisdictions | Existing building provisions in adopting jurisdictions | Fire-code cross-reference amendment caveat applies. | src:usa-wv:rule:87-04-final-2022 |
| adoption:usa-wv:87-04:ispsc-2018 | Swimming Pool / Spa | ahj:usa-wv:state-fire-commission | 87 CSR 4 | International Swimming Pool and Spa Code | 2018 | 2022-05-05 | 2022-08-01 | 2022-08-01 | 2022-08-01 for adopting jurisdictions | Pool/spa provisions in adopting jurisdictions | No separate transition rule extracted. | src:usa-wv:rule:87-04-final-2022 |
| adoption:usa-wv:87-01:emergency-fire-code-2021 | Fire - operational / prevention code | ahj:usa-wv:state-fire-commission | 87 CSR 1 emergency | NFPA 1; NFPA 101 | 2021 | 2025-03-06 | 2025-03-06 | 2025-03-06 | 2025-03-06 | Statewide fire/life-safety code with stated exemptions | SOS lists emergency status while legislative rule proceeds. | src:usa-wv:sos:87-01-rule-page; src:usa-wv:rule:87-01-emergency-2025 |
| adoption:usa-wv:87-01:final-fire-code-2021 | Fire - operational / prevention code | ahj:usa-wv:state-fire-commission | 87 CSR 1 final legislative rule | NFPA 1; NFPA 101 | 2021 | 2026-05-13 | 2026-07-01 | 2026-07-01 | 2026-07-01 | Statewide fire/life-safety code with stated exemptions | Final legislative rule terminates 2031-08-01. | src:usa-wv:sos:87-01-rule-page; src:usa-wv:rule:87-01-final-2026 |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

87 CSR 4 has a final filing date of 2022-05-05, an effective date of 2022-08-01, and a sunset date of 2027-08-01. The State Building Code applies prospectively through local adoption by counties or municipalities rather than retroactively to existing work. No permit-application, permit-issuance, or parallel-use grace period was extracted from the reviewed State Building Code sources.

For the State Fire Code, the Secretary of State rule detail shows 87 CSR 1 as active emergency on 2026-06-26, with the emergency effective date listed as 2025-03-06 and the final legislative rule effective date listed as 2026-07-01. The final rule terminates on 2031-08-01 unless renewed or replaced.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-wv:87-04-effective | State Building Code | fixed_rule_effective_date | 2022-08-01 | 87 CSR 4 final rule effective date after 2022-05-05 filing | no statewide parallel-use period extracted | src:usa-wv:sos:87-04-rule-page; src:usa-wv:rule:87-04-final-2022 | partially_verified |
| date-rule:usa-wv:building-prospective-local-adoption | State Building Code | prospective_local_adoption | local adoption dependent | County or municipality adopts State Building Code | existing local building codes voided after state-code promulgation; local adoption must use State Building Code | src:usa-wv:statute:15a-11-5; src:usa-wv:statute:7-1-3n; src:usa-wv:statute:8-12-13 | partially_verified |
| date-rule:usa-wv:87-04-sunset | State Building Code | sunset | 2027-08-01 | 87 CSR 4 sunset provision | renewal/replacement required to avoid sunset | src:usa-wv:rule:87-04-final-2022 | partially_verified |
| date-rule:usa-wv:87-01-emergency | State Fire Code | emergency_effective_date | 2025-03-06 | SOS-listed 87 CSR 1 emergency effective date | emergency rule remains effective while legislative rule proceeds per SOS note | src:usa-wv:sos:87-01-rule-page; src:usa-wv:rule:87-01-emergency-2025 | partially_verified |
| date-rule:usa-wv:87-01-final | State Fire Code | final_legislative_effective_date | 2026-07-01 | 87 CSR 1 final legislative rule effective date | emergency-to-final handoff requires review after 2026-07-01 | src:usa-wv:sos:87-01-rule-page; src:usa-wv:rule:87-01-final-2026 | partially_verified |
| date-rule:usa-wv:87-01-sunset | State Fire Code | sunset | 2031-08-01 | 87 CSR 1 final rule sunset provision | renewal/replacement required to avoid sunset | src:usa-wv:rule:87-01-final-2026 | partially_verified |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Fire - operational / prevention code | 87 CSR 1 final legislative Fire Code adopting NFPA 1-2021 and NFPA 101-2021 | 2026-05-13 | 2026-05-13 | 2026-07-01 | 2026-07-01 | 2026-07-01 | active_watch | src:usa-wv:sos:87-01-rule-page; src:usa-wv:rule:87-01-final-2026 | As of 2026-06-26, SOS lists the rule as active emergency while final legislative rulemaking proceeds. Recheck on or after 2026-07-01. |
| Building and related construction codes | 87 CSR 4 renewal or replacement | null | null | null | null | null | active_watch | src:usa-wv:sos:87-04-rule-page; src:usa-wv:rule:87-04-final-2022 | Existing 87 CSR 4 sunset is 2027-08-01. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-wv:agricultural-exemption | Building; fire; property maintenance | Buildings or structures used primarily for agricultural purposes, subject to stated exceptions in the Fire Code | Project/use is primarily agricultural | Agricultural uses are exempt from the State Building Code and State Fire Code in the reviewed rules/statutes, with Fire Code exceptions for certain sleeping, health-care, detention, or correctional occupancies. | src:usa-wv:statute:15a-11-5; src:usa-wv:rule:87-04-final-2022; src:usa-wv:rule:87-01-final-2026 | partially_verified |
| applicability-rule:usa-wv:fire-code-exemptions | Fire - operational / prevention code | Personal care homes caring for three or fewer patients; dwelling houses for no more than two families; boilers | Fire Code scope and exemptions | 87 CSR 1 exempts specified small personal-care homes, one- and two-family dwelling houses, agricultural buildings subject to stated limits, and boiler permitting/inspection processes under the Division of Labor. | src:usa-wv:rule:87-01-final-2026 | partially_verified |
| applicability-rule:usa-wv:ipmc-local-option | Property Maintenance | Local adoption/rejection of 2018 IPMC | Local jurisdiction considers State Building Code package | 87 CSR 4 lists the 2018 IPMC and states that it may be rejected by a local jurisdiction. | src:usa-wv:rule:87-04-final-2022 | partially_verified |
| applicability-rule:usa-wv:appendices-local-adoption | Building package appendices and discretionary local provisions | Appendices, penalties, appeals boards, and other local-adoption-dependent provisions | Local ordinance/order adoption | 87 CSR 4 treats model-code appendices and discretionary provisions as unenforceable unless adopted by the local jurisdiction. | src:usa-wv:rule:87-04-final-2022 | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** incorporated_model_codes_with_state_specific_exceptions_and_cross_reference_substitutions

**Where amendments are published:** West Virginia Secretary of State Code of State Rules final PDFs for 87 CSR 4 and 87 CSR 1, with agency indexing on the Office of the State Fire Marshal laws page.

**Amendment parsing status:** partial_key_provisions_only

87 CSR 4 incorporates model codes by reference and then modifies specific provisions, including code-conflict rules, fire-code substitutions, residential energy treatment, NEC-related amendments, local enforcement provisions, local adoption requirements, and agricultural exemptions. This report captures high-impact amendments but does not attempt a section-by-section amendment extraction.

### 5.2 State Amendment Sources

| Amendment Source ID | Code Family | Publication Path | Effective Date | Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| amendment-source:usa-wv:87-04-final-2022 | Building package | SOS final PDF for 87 CSR 4 | 2022-08-01 | current_active_rule | src:usa-wv:sos:87-04-rule-page; src:usa-wv:rule:87-04-final-2022 | Full rule text is an official PDF; detailed model-code amendment extraction remains incomplete. |
| amendment-source:usa-wv:87-01-emergency-2025 | Fire Code | SOS emergency PDF for 87 CSR 1 | 2025-03-06 | active_emergency_per_sos | src:usa-wv:sos:87-01-rule-page; src:usa-wv:rule:87-01-emergency-2025 | Emergency rule active while final rule proceeds according to SOS detail page. |
| amendment-source:usa-wv:87-01-final-2026 | Fire Code | SOS final PDF for 87 CSR 1 | 2026-07-01 | final_filed_future_effective_on_report_date | src:usa-wv:sos:87-01-rule-page; src:usa-wv:rule:87-01-final-2026 | Effective shortly after this report date; recheck current status after 2026-07-01. |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Base Code Section / Topic | Amendment Summary | Source IDs | Confidence | Status |
| --- | --- | --- | --- | --- | --- | --- |
| amendment:usa-wv:87-04:fire-code-precedence | Building; fire | Code conflicts | 87 CSR 4 states the State Fire Code takes precedence over the State Building Code when conflicts occur. | src:usa-wv:rule:87-04-final-2022 | 0.82 | partially_verified |
| amendment:usa-wv:87-04:plumbing-health-precedence | Plumbing | Code conflicts | 87 CSR 4 states Department of Health and Human Resources plumbing rules take precedence when they conflict with the State Building Code. | src:usa-wv:rule:87-04-final-2022 | 0.72 | partially_verified |
| amendment:usa-wv:87-04:no-ifc-reference | Building; existing building; fire references | IFC cross-references | 87 CSR 4 does not adopt the International Fire Code as the operational fire code and substitutes NFPA Life Safety Code references for International Fire Code references in incorporated ICC codes. | src:usa-wv:rule:87-04-final-2022 | 0.70 | partially_verified_with_edition_caveat |
| amendment:usa-wv:87-04:energy-split | Energy | Residential/commercial energy | 87 CSR 4 uses 2015 IECC for residential energy provisions and ASHRAE/IESNA 90.1-2013 for commercial energy provisions, while exempting IRC Chapter 11 energy-efficiency provisions. | src:usa-wv:rule:87-04-final-2022 | 0.82 | partially_verified |
| amendment:usa-wv:87-04:ipmc-local-rejection | Property Maintenance | 2018 IPMC | 87 CSR 4 states the 2018 International Property Maintenance Code may be rejected by a local jurisdiction. | src:usa-wv:rule:87-04-final-2022 | 0.76 | partially_verified |
| amendment:usa-wv:87-04:local-adoption-notice | Local enforcement | Adoption notice | 87 CSR 4 requires each local jurisdiction adopting the State Building Code to notify the State Fire Commission or State Fire Marshal within 30 days. | src:usa-wv:rule:87-04-final-2022 | 0.84 | partially_verified |
| amendment:usa-wv:87-01:nfpa-precedence | Fire Code | NFPA conflicts | 87 CSR 1 states NFPA 101 takes precedence over NFPA 1; the State Fire Code takes precedence over the State Building Code and over incorporated NFPA codes. | src:usa-wv:rule:87-01-final-2026 | 0.80 | partially_verified |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-wv"
  model: "local_adoption_with_local_enforcement_plus_state_fire_marshal_statewide_enforcement_authority"
  enforcing_entities:
    - "county or municipality that adopts the State Building Code"
    - "local jurisdiction identified by 87 CSR 4 after adoption"
    - "Office of the State Fire Marshal for fire programs and state-level fire/building-code enforcement powers"
  required_officials:
    - "building code officials, building code inspectors, and plans examiners subject to State Fire Marshal certification rules for municipal, county, and public-sector officials"
    - "electrical inspectors subject to state certification when performing compensated electrical inspections"
  state_reserved_activities:
    - "State Fire Marshal enforcement of fire programs and state fire/life-safety laws"
    - "State Fire Marshal authority to enforce the fire code, building code, and related laws within statutory scope"
    - "State Fire Commission priority decisions for potentially conflicting local laws, ordinances, or regulations"
  source_ids:
    - "src:usa-wv:statute:15a-11-5"
    - "src:usa-wv:statute:15a-10-2"
    - "src:usa-wv:statute:15a-10-3"
    - "src:usa-wv:statute:7-1-3n"
    - "src:usa-wv:statute:8-12-13"
    - "src:usa-wv:rule:87-04-final-2022"
    - "src:usa-wv:sos:103-06-rule-page"
  verification_status: "partially_verified"
  confidence: 0.78
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-wv"
  model: "state_code_required_for_local_building_code_with_limited_more_stringent_local_rules"
  applies_to_code_families:
    - "State Building Code"
    - "local appendices and local discretionary provisions"
    - "property maintenance where locally adopted or rejected"
  approval_required: "routine preapproval process not resolved; State Fire Commission decides priority when local provisions create questions of conflict or relative priority"
  approving_authority_id: "ahj:usa-wv:state-fire-commission"
  filing_required: "adopting jurisdiction must notify the State Fire Commission or State Fire Marshal within 30 days of adoption; exact filing form or central public registry unresolved"
  registry_exists: "not identified"
  registry_source_ids: []
  legal_basis_source_ids:
    - "src:usa-wv:statute:15a-11-5"
    - "src:usa-wv:statute:7-1-3n"
    - "src:usa-wv:statute:8-12-13"
    - "src:usa-wv:rule:87-04-final-2022"
  verification_status: "partially_verified"
  confidence: 0.63
```

### 6.3 Local Enforcement vs. Local Amendment Summary

West Virginia separates the local power to enforce an adopted State Building Code from the ability to maintain an independent local building code. Counties and municipalities may adopt building/housing codes, but after State Building Code promulgation, existing local building codes are voided and any local building-code adoption must be the State Building Code. Local enforcement is then performed by the adopting jurisdiction. More stringent local laws, ordinances, or regulations may govern only when they are not inconsistent with the State Building Code and are not contrary to recognized engineering standards, and the State Fire Commission decides priority when a conflict or priority question arises.

### 6.4 Known Local Amendment Registries

| Registry ID | Registry Name | Owner | URL / Source ID | Coverage | Status | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| registry:usa-wv:local-amendments:unresolved | unresolved | unresolved | none | statewide local amendments | unresolved | No official statewide local-amendment registry was identified in the reviewed source set. |

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendment set was parsed. Future review should sample large jurisdictions and compare local ordinances against the State Building Code adoption-notice requirement and the State Fire Commission priority rule.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: partial_model_only

Jurisdiction stack:

```text
Address
  -> State of West Virginia
  -> County
  -> Municipality or unincorporated county area
  -> Local adopting jurisdiction, if the State Building Code has been adopted
  -> Building-code official / inspector / plans examiner certified under applicable State Fire Marshal rules
  -> State Fire Marshal fire/life-safety jurisdiction
  -> Electrical licensing and electrical-inspector certification constraints
  -> Applicable 87 CSR 4 adoption records
  -> Applicable 87 CSR 1 fire-code record
  -> Local appendices, local discretionary provisions, and more stringent local rules if validly adopted and not inconsistent
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | not selected | none | statewide | unresolved | pending |
| County | not selected | none | statewide | unresolved | pending |
| Municipality | not selected | none | statewide | unresolved | pending |
| Fire District | not selected | none | statewide | unresolved | pending |
| Special District | not selected | none | statewide | unresolved | pending |

### 7.3 AHJ Contact Data

No AHJ contact data was populated. The core AHJ classification is source-backed, but address-level AHJ resolution requires a local adoption registry or jurisdiction-by-jurisdiction ordinance review.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Publisher | Type | Date / Status | URL | Primary extracted fields | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-wv:statute:15a-11-1 | West Virginia Code §15A-11-1, State Fire Commission | West Virginia Legislature | statute | current online code | https://code.wvlegislature.gov/15A-11-1/ | State Fire Commission existence/composition. | official_html |
| src:usa-wv:statute:15a-11-3 | West Virginia Code §15A-11-3, State Fire Code rule authority | West Virginia Legislature | statute | current online code | https://code.wvlegislature.gov/15A-11-3/ | State Fire Commission authority to propose State Fire Code; statewide force/effect. | official_html |
| src:usa-wv:statute:15a-11-5 | West Virginia Code §15A-11-5, State Building Code rule authority | West Virginia Legislature | statute | current online code | https://code.wvlegislature.gov/15A-11-5/ | State Building Code authority, local adoption force/effect, local priority rule, enforcement responsibility. | official_html |
| src:usa-wv:statute:15a-10-1 | West Virginia Code §15A-10-1, State Fire Marshal | West Virginia Legislature | statute | current online code | https://code.wvlegislature.gov/15A-10-1/ | Office of State Fire Marshal continued and enforcement role context. | official_html |
| src:usa-wv:statute:15a-10-2 | West Virginia Code §15A-10-2, State Fire Marshal powers and duties | West Virginia Legislature | statute | current online code | https://code.wvlegislature.gov/15A-10-2/ | State Fire Marshal enforcement powers; transfer of fire/building-code violation enforcement from Fire Commission to Office of State Fire Marshal. | official_html |
| src:usa-wv:statute:15a-10-3 | West Virginia Code §15A-10-3, State Fire Marshal enforcement authority | West Virginia Legislature | statute | current online code | https://code.wvlegislature.gov/15A-10-3/ | State Fire Marshal authority to enforce fire code, building code, and related laws; inspection authority. | official_html |
| src:usa-wv:statute:7-1-3n | West Virginia Code §7-1-3n, county building and housing code authority | West Virginia Legislature | statute | current online code | https://code.wvlegislature.gov/7-1-3N/ | County authority to adopt building/housing codes; requirement to use State Building Code after voidance of existing county code. | official_html |
| src:usa-wv:statute:8-12-13 | West Virginia Code §8-12-13, municipal building regulation and State Building Code | West Virginia Legislature | statute | current online code | https://code.wvlegislature.gov/8-12-13/ | Municipal building, electrical, plumbing regulatory authority; requirement to use State Building Code after voidance of existing municipal building code. | official_html |
| src:usa-wv:statute:29-3b-2 | West Virginia Code §29-3B-2, electrician licensing definitions and license requirement | West Virginia Legislature | statute | current online code | https://code.wvlegislature.gov/29-3B-2/ | State electrical work licensing requirement; National Electrical Code definition context. | official_html |
| src:usa-wv:statute:29-3b-5 | West Virginia Code §29-3B-5, electrician licensing rules | West Virginia Legislature | statute | current online code | https://code.wvlegislature.gov/29-3B-5/ | State Fire Marshal rulemaking for electrician licensing, examinations, and renewals. | official_html |
| src:usa-wv:statute:29-3b-9 | West Virginia Code §29-3B-9, local electrician-license limitation | West Virginia Legislature | statute | current online code | https://code.wvlegislature.gov/29-3B-9/ | Localities may not require another electrician license from a valid state license holder. | official_html |
| src:usa-wv:statute:29-3c-3 | West Virginia Code §29-3C-3, electrical inspector certification requirement | West Virginia Legislature | statute | current online code | https://code.wvlegislature.gov/29-3C-3/ | State certification requirement for compensated electrical inspectors and inspector-definition bridge. | official_html |
| src:usa-wv:statute:29-3c-4 | West Virginia Code §29-3C-4, electrical inspector certification rules | West Virginia Legislature | statute | current online code | https://code.wvlegislature.gov/29-3C-4/ | State Fire Marshal certification/rulemaking for electrical inspectors. | official_html |
| src:usa-wv:osfm:laws-page | Office of the State Fire Marshal Laws page | West Virginia Office of the State Fire Marshal | agency_page | accessed 2026-06-26 | https://firemarshal.wv.gov/laws | Agency index listing fire, building, licensing, certification, and related legislative rules. | official_html_index |
| src:usa-wv:sos:87-04-rule-page | Secretary of State rule detail, 87 CSR 4, State Building Code | West Virginia Secretary of State | rule_detail | effective 2022-08-01 | https://apps.sos.wv.gov/adlaw/csr/rule.aspx?rule=87-04 | Rule status, filing, effective date, and rule documents for State Building Code. | official_html |
| src:usa-wv:rule:87-04-final-2022 | 87 CSR 4, State Building Code, final filed 2022-05-05 | West Virginia Secretary of State / State Fire Commission | administrative_rule_pdf | effective 2022-08-01; sunset 2027-08-01 | https://apps.sos.wv.gov/adlaw/csr/readfile.aspx?DocId=55263&Format=PDF | Current State Building Code rule text and incorporated model-code package. | official_pdf |
| src:usa-wv:sos:87-01-rule-page | Secretary of State rule detail, 87 CSR 1, Fire Code | West Virginia Secretary of State | rule_detail | active emergency; final effective 2026-07-01 | https://apps.sos.wv.gov/adlaw/csr/rule.aspx?rule=87-01 | Rule status, emergency date, final filing, and effective date for Fire Code. | official_html |
| src:usa-wv:rule:87-01-emergency-2025 | 87 CSR 1, Fire Code, emergency filed 2025-03-06 | West Virginia Secretary of State / State Fire Commission | emergency_rule_pdf | effective 2025-03-06 per SOS rule detail | https://apps.sos.wv.gov/adlaw/csr/readfile.aspx?DocId=57865&Format=PDF | Emergency State Fire Code rule and transition explanation. | official_pdf |
| src:usa-wv:rule:87-01-final-2026 | 87 CSR 1, Fire Code, final filed 2026-05-13 | West Virginia Secretary of State / State Fire Commission | administrative_rule_pdf | effective 2026-07-01; sunset 2031-08-01 | https://apps.sos.wv.gov/adlaw/csr/readfile.aspx?DocId=58962&Format=PDF | Final legislative State Fire Code rule and incorporated NFPA 1 / NFPA 101 package. | official_pdf |
| src:usa-wv:sos:103-06-rule-page | Secretary of State rule detail, 103 CSR 6, public-sector building-code officials, inspectors, and plans examiners | West Virginia Secretary of State / Office of State Fire Marshal | rule_detail | effective 2024-05-01 | https://apps.sos.wv.gov/adlaw/csr/rule.aspx?rule=103-06 | Certification and continuing education rule for municipal, county, and other public-sector building-code officials, building-code inspectors, and plans examiners. | official_html |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-wv:sos:87-04-rule-page | authority_display_caveat | The SOS HTML detail page displays an authority line that differs from the authority section in the 87 CSR 4 final PDF and the current WV Code building-code statute. | Treat 87 CSR 4 final PDF and W. Va. Code §15A-11-5 as controlling for building-code authority pending legal review. |
| src:usa-wv:rule:87-04-final-2022 | pdf_extraction_caveat | The rule is an official PDF; high-impact provisions were extracted, but a line-by-line amendment inventory was not completed. | partial_extraction_accepted_for_core_fields |
| src:usa-wv:rule:87-04-final-2022 | internal_edition_caveat | 87 CSR 4 contains an apparent Life Safety Code edition ambiguity in fire-code substitution language. | Flagged as open issue; do not normalize a single construction-reference fire-code edition beyond the State Fire Code record without review. |
| src:usa-wv:sos:87-01-rule-page | transition_caveat | The SOS rule detail shows 87 CSR 1 as active emergency on 2026-06-26, while also listing a final file effective 2026-07-01. | Recheck after 2026-07-01 before marking fire-code status verified. |
| src:usa-wv:rule:87-01-emergency-2025 | emergency_rule_caveat | Emergency rule filing is temporary and interacts with the final legislative rule. | Use only for current-status explanation through the transition period. |
| src:usa-wv:sos:103-06-rule-page | scope_caveat | Rule page was used to identify the public-sector certification path but the full text was not fully parsed. | Use for local-official certification signal; parse full rule before verified status. |

### 8.3 Supplemental Sources

None used. This report relies on official West Virginia legislative, Secretary of State, and Office of State Fire Marshal sources.

### 8.4 Source Extraction Metadata

| Extraction ID | Source IDs | Extracted On | Extractor | Coverage | Notes |
| --- | --- | --- | --- | --- | --- |
| extraction:usa-wv:2026-06-26:authority | src:usa-wv:statute:15a-11-3; src:usa-wv:statute:15a-11-5; src:usa-wv:statute:15a-10-2; src:usa-wv:statute:15a-10-3 | 2026-06-26 | GPT-5.5 Thinking | authority_structure | Core state authority and enforcement model extracted. |
| extraction:usa-wv:2026-06-26:building-code | src:usa-wv:sos:87-04-rule-page; src:usa-wv:rule:87-04-final-2022 | 2026-06-26 | GPT-5.5 Thinking | current_code_adoptions | Building-code package and high-impact amendments extracted; full amendment inventory remains open. |
| extraction:usa-wv:2026-06-26:fire-code | src:usa-wv:sos:87-01-rule-page; src:usa-wv:rule:87-01-emergency-2025; src:usa-wv:rule:87-01-final-2026 | 2026-06-26 | GPT-5.5 Thinking | fire_code_transition | Emergency/final fire-code transition captured with post-2026-07-01 watch. |
| extraction:usa-wv:2026-06-26:local-adoption | src:usa-wv:statute:7-1-3n; src:usa-wv:statute:8-12-13; src:usa-wv:rule:87-04-final-2022 | 2026-06-26 | GPT-5.5 Thinking | local_enforcement_and_adoption | County/municipal adoption model and local-notice requirement captured. |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 1.00 | none | Status matches current evidence depth: core fields are sourced; full amendment and AHJ work remains open. |
| report | risk.overall_confidence | 0.68 | verified | 1.00 | none | Confidence is intentionally conservative because local amendment registry, AHJ contacts, and full amendment parsing remain unresolved. |
| ahj:usa-wv:state-fire-commission | authority.role | building and fire code rulemaking body | partially_verified | 0.90 | src:usa-wv:statute:15a-11-3; src:usa-wv:statute:15a-11-5 | Core rulemaking role verified from statutes. |
| ahj:usa-wv:state-fire-marshal | authority.role | enforcement and state fire/life-safety authority | partially_verified | 0.84 | src:usa-wv:statute:15a-10-2; src:usa-wv:statute:15a-10-3 | Enforcement role verified; program-specific inspection boundaries may need further parsing. |
| adoption:usa-wv:87-04:ibc-2018 | base_code.edition | 2018 IBC | partially_verified | 0.86 | src:usa-wv:rule:87-04-final-2022 | Extracted from official 87 CSR 4 PDF. |
| adoption:usa-wv:87-04:nec-2020 | base_code.edition | 2020 NEC | partially_verified | 0.84 | src:usa-wv:rule:87-04-final-2022 | Extracted from official 87 CSR 4 PDF. |
| adoption:usa-wv:87-01:emergency-fire-code-2021 | status | active_emergency | partially_verified | 0.80 | src:usa-wv:sos:87-01-rule-page; src:usa-wv:rule:87-01-emergency-2025 | Requires recheck after 2026-07-01. |
| adoption:usa-wv:87-01:final-fire-code-2021 | effective_date | 2026-07-01 | partially_verified | 0.86 | src:usa-wv:sos:87-01-rule-page; src:usa-wv:rule:87-01-final-2026 | Final rule effective date verified from SOS detail and final PDF. |
| local-enforcement:usa-wv | model | local_adoption_with_local_enforcement_plus_state_fire_marshal_statewide_enforcement_authority | partially_verified | 0.78 | src:usa-wv:statute:15a-11-5; src:usa-wv:rule:87-04-final-2022 | Local adoption/enforcement backed; local-adoption registry unresolved. |
| local-amendment-rule:usa-wv | model | state_code_required_for_local_building_code_with_limited_more_stringent_local_rules | partially_verified | 0.63 | src:usa-wv:statute:15a-11-5; src:usa-wv:statute:7-1-3n; src:usa-wv:statute:8-12-13 | Requires full local-amendment procedure review before verified status. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Every `src:usa-wv:*` citation in the body appears in section 8. |
| All authority IDs resolve | pass | All authority IDs used in graph and tables are defined or explicitly labeled unresolved. |
| All current code families have adoption records or explicit unresolved rows | pass | Core families are populated; conveyance/elevator remains an explicit unresolved row. |
| Building and operational fire code are separated | pass | 87 CSR 4 and 87 CSR 1 are separately modeled. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Date columns are separated; null values remain explicit for unresolved conveyance items. |
| Effective dates are valid ISO dates where populated | pass | Populated dates use YYYY-MM-DD format. |
| No impossible date sequences | pass | 87 CSR 1 future final effective date is intentionally handled as a pending transition from emergency status. |
| Transition rules have explicit trigger conditions | partial | Fixed effective/sunset and emergency/final transition rules are captured; no permit-date rule was found in reviewed sources. |
| Permit-date logic is captured where applicable | partial | No permit-date grace rule was extracted from 87 CSR 4; mark as open for full rule/legal review. |
| Local enforcement model classified | pass | Local adoption plus local enforcement and State Fire Marshal authority are captured. |
| Local amendment rule classified | partial | Core local-code constraint captured; registry and detailed procedure unresolved. |
| AHJ confirmation metadata present | fail | Address-level AHJ contacts and boundary source stack are not populated. |
| Official-source caveats captured | pass | PDF, authority-display, edition ambiguity, emergency-transition, and certification-rule caveats are recorded. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-wv:001 | high | fire-code transition | SOS shows 87 CSR 1 active emergency on 2026-06-26 and final effective 2026-07-01. | Recheck SOS 87 CSR 1 after 2026-07-01 and update status from active_emergency to active_final if appropriate. | null | 2026-07-08 | open |
| issue:usa-wv:002 | high | building-code sunset | 87 CSR 4 sunsets on 2027-08-01. | Add watch for any 87 CSR 4 renewal, amendment, replacement, or new model-code adoption package. | null | 2027-02-01 | open |
| issue:usa-wv:003 | medium | local amendment registry | No central public local-amendment/adoption registry was identified from the reviewed source set. | Search OSFM, State Fire Commission meeting materials, and municipal/county filings for adoption notices or registries. | null | null | open |
| issue:usa-wv:004 | medium | full amendment inventory | High-impact WV amendments were captured, but 87 CSR 4 and 87 CSR 1 were not parsed section-by-section. | Perform a rule-by-rule amendment extraction and normalize all code-section amendments. | null | null | open |
| issue:usa-wv:005 | medium | Life Safety Code edition ambiguity | 87 CSR 4 appears to contain conflicting Life Safety Code edition references for ICC fire-code substitutions. | Verify with State Fire Commission or OSFM guidance before normalizing a single construction-reference edition. | null | null | open |
| issue:usa-wv:006 | medium | conveyance/elevator authority | Elevator/conveyance authority and adopted standards were not resolved. | Research West Virginia Division of Labor or other official conveyance authority sources and populate code family row. | null | null | open |
| issue:usa-wv:007 | low | AHJ contact data | Address-level AHJ contacts are absent. | Build local adoption/AHJ dataset from county and municipal sources. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-wv:87-04 | src:usa-wv:sos:87-04-rule-page | html_diff | monthly | New final, emergency, modified, or notice document for State Building Code | 2026-06-26 | active |
| watch:usa-wv:87-01 | src:usa-wv:sos:87-01-rule-page | html_diff | weekly_until_2026-07-15_then_monthly | Status changes from active emergency to active final, or new Fire Code filing | 2026-06-26 | active |
| watch:usa-wv:osfm-laws | src:usa-wv:osfm:laws-page | html_diff | monthly | Agency updates laws/rules index or adds guidance documents | 2026-06-26 | active |
| watch:usa-wv:15a-11-5 | src:usa-wv:statute:15a-11-5 | statute_diff | quarterly | Building-code authority statute amendment | 2026-06-26 | active |
| watch:usa-wv:103-06 | src:usa-wv:sos:103-06-rule-page | html_diff | quarterly | Certification-rule amendment or new public-sector inspector requirements | 2026-06-26 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-24 | Generated baseline draft report | report:usa-wv | none | Codex | Starting stub contained unresolved placeholder content. |
| 2026-06-26 | Populated source-backed West Virginia authority, adoption, local enforcement, fire-code transition, source registry, QA, and open-issue sections | report:usa-wv; ahj:usa-wv:state-fire-commission; adoption:usa-wv:87-04:ibc-2018; adoption:usa-wv:87-01:emergency-fire-code-2021 | src:usa-wv:statute:15a-11-3; src:usa-wv:statute:15a-11-5; src:usa-wv:rule:87-04-final-2022; src:usa-wv:rule:87-01-emergency-2025; src:usa-wv:rule:87-01-final-2026 | GPT-5.5 Thinking | Status upgraded to partially_verified with unresolved items kept explicit. |
