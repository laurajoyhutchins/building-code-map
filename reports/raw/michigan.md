---
state:
  state_id: "US-MI"
  name: "Michigan"
  abbreviation: "MI"
report:
  report_id: "state-report:usa-mi"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-25"
  last_verified: null
  reviewed_by: null
risk:
  overall_confidence: 0.68 # 0.00 - 1.00
  risk_flags:
    - "fire_operational_code_scope_not_fully_parsed"
    - "local_amendment_scope_unresolved"
    - "mcl_125_1508b_full_text_not_extracted"
    - "residential_2021_and_residential_energy_update_stayed"
    - "permit_date_transition_rule_unresolved"
  open_questions_count: 6

---

# State Building Code Authority Report: Michigan

## 1. Executive Summary

- **Authority model:** Michigan administers a statewide construction-code framework through the Department of Licensing and Regulatory Affairs (LARA), Director's Office and Bureau of Construction Codes (BCC). The state construction code rules cite section 4 of the Stille-DeRossett-Hale Single State Construction Code Act, 1972 PA 230, MCL 125.1504, as the rulemaking authority for the code parts reviewed here. Appeals are handled through the State Construction Code Commission, and elevator/conveyance functions are administered through BCC's Elevator Section with Elevator Safety Board rule/licensing functions.

- **Statewide code status:** Michigan has verified current statewide adoptions for the 2021 Michigan Building Code, 2021 Michigan Rehabilitation Code, 2021 Michigan Mechanical Code, 2021 Michigan Plumbing Code, Michigan Electrical Code based on the 2023 NEC, 2021 Michigan Commercial Energy Code, and Michigan Elevator Code standards. Residential remains on the 2015 Michigan Residential Code while the 2021 residential-code and residential-energy update is subject to a court-order stay. A single statewide operational fire-code edition was not fully reduced to a model-code adoption record in this pass.

- **Local enforcement model:** The best-supported current classification is a state-default / local-assumption model. LARA/BCC is the state construction-code authority, while local governmental subdivisions may act as enforcing agencies where they have assumed responsibility and designated enforcement. The MCL 125.1508b local-enforcement text should be extracted directly from the official MCL source before this field is treated as verified.

- **Local amendment posture:** General local technical-amendment authority remains unresolved. Local enforcement is not treated as local amendment authority. This report verifies only rule-specific local options requiring a local resolution filed with LARA/BCC, such as certain high-rise or sprinkler exceptions in the Michigan Building Code rules.

- **Known transition periods or pending changes:** The 2021 residential and residential-energy code updates were listed by LARA as set to become effective on 2025-08-29, but a 2025-07-07 court order temporarily prevents LARA from implementing those rules while the order remains in effect. LARA's open-rules page also shows 2026 rulemaking activity for Part 4 Building, Part 5 Residential, Part 7 Plumbing, Part 9A Mechanical, Part 10 residential energy, and Part 10A commercial energy.

- **Production readiness:** partially_verified_not_production

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: LARA Director's Office/BCC is the primary statewide construction-code authority
    for the code parts reviewed; rules cite MCL 125.1504 as authority.
  confidence: 0.85
  source_ids:
  - src:usa-mi:mcl-125-1504
  - src:usa-mi:admin-part4-building
  - src:usa-mi:lara-bcc-home
- topic: Primary building code edition
  finding: 2021 Michigan Building Code, based on the 2021 IBC, is current with an
    effective date of 2025-04-09.
  confidence: 0.95
  source_ids:
  - src:usa-mi:lara-code-books
  - src:usa-mi:admin-part4-building
- topic: Residential code edition
  finding: 2015 Michigan Residential Code remains current; 2021 IRC / residential-energy
    update implementation is stayed.
  confidence: 0.9
  source_ids:
  - src:usa-mi:lara-code-books
  - src:usa-mi:lara-open-rules-res-energy-stay
- topic: Mechanical / plumbing / electrical
  finding: Mechanical and plumbing use 2021 model-code bases effective 2024-03-12;
    electrical uses the 2023 NEC effective 2024-03-12.
  confidence: 0.92
  source_ids:
  - src:usa-mi:admin-part9a-mechanical
  - src:usa-mi:admin-part7-plumbing
  - src:usa-mi:admin-part8-electrical
  - src:usa-mi:lara-code-books
- topic: Commercial energy
  finding: Michigan Commercial Energy Code uses the 2021 IECC for nonresidential buildings
    and ASHRAE/IESNA 90.1-2019; effective date 2025-04-22.
  confidence: 0.92
  source_ids:
  - src:usa-mi:admin-part10a-commercial-energy
  - src:usa-mi:lara-code-books
- topic: Fire code authority
  finding: Bureau of Fire Services / State Fire Marshal and the State Fire Safety
    Board are identified, but a complete operational fire-code adoption matrix is
    unresolved.
  confidence: 0.55
  source_ids:
  - src:usa-mi:lara-bfs-home
  - src:usa-mi:fire-prevention-code-act207
  - src:usa-mi:state-fire-safety-board-rules
- topic: Local enforcement
  finding: State-default / local-assumption model is provisionally classified from
    MCL 125.1508b search-surface text and BCC/AHJ materials; direct statutory extraction
    remains open.
  confidence: 0.7
  source_ids:
  - src:usa-mi:mcl-125-1508b
  - src:usa-mi:bcc-building-permits
  - src:usa-mi:admin-part8-electrical
- topic: Local amendments
  finding: General local amendment authority is unresolved; only specific rule-level
    local-resolution options were captured.
  confidence: 0.35
  source_ids:
  - src:usa-mi:admin-part4-building
- topic: Effective / operative date rule
  finding: Code effective dates were captured separately; statewide permit-date or
    issuance-date transition logic remains unresolved except for code-specific permit
    provisions noted below.
  confidence: 0.6
  source_ids:
  - src:usa-mi:lara-code-books
  - src:usa-mi:admin-part7-plumbing
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-mi:lara-bcc-director |
| Authority name | Michigan Department of Licensing and Regulatory Affairs (LARA), Director's Office / Bureau of Construction Codes |
| Authority type | state agency / director |
| Legal basis | Stille-DeRossett-Hale Single State Construction Code Act, 1972 PA 230, MCL 125.1504, as cited in Michigan Administrative Rules for construction-code parts |
| Role | Promulgates and administers state construction-code rules and statewide code adoptions for the construction-code families reviewed here |
| Enforcement model | statewide construction code with state/default and local-assumption enforcement structure; field remains partially verified pending full MCL 125.1508b extraction |
| Source IDs | src:usa-mi:mcl-125-1504; src:usa-mi:admin-part4-building; src:usa-mi:lara-bcc-home |
| Verification status | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-mi:lara-bcc-director | LARA Director's Office / Bureau of Construction Codes | Adopts and administers the Michigan Building Code | MCL 125.1504; Part 4 rules | src:usa-mi:mcl-125-1504; src:usa-mi:admin-part4-building; src:usa-mi:lara-code-books | partially_verified |
| Residential | ahj:usa-mi:lara-bcc-director | LARA Director's Office / Bureau of Construction Codes | Administers Michigan Residential Code; 2021 update stayed | MCL 125.1504; Part 5 / residential rulemaking source not fully parsed | src:usa-mi:lara-code-books; src:usa-mi:lara-open-rules-res-energy-stay; src:usa-mi:lara-open-rules | partially_verified_stayed_update |
| Existing Building / Rehabilitation | ahj:usa-mi:lara-bcc-director | LARA Director's Office / Bureau of Construction Codes | Adopts and administers Michigan Rehabilitation Code | MCL 125.1504; Part 4 adoption references IEBC; LARA code-books page | src:usa-mi:admin-part4-building; src:usa-mi:lara-code-books | partially_verified |
| Mechanical | ahj:usa-mi:lara-bcc-director | LARA Director's Office / Bureau of Construction Codes | Adopts and administers Michigan Mechanical Code | MCL 125.1504; Part 9A rules | src:usa-mi:admin-part9a-mechanical; src:usa-mi:lara-code-books | partially_verified |
| Plumbing | ahj:usa-mi:lara-bcc-director | LARA Director's Office / Bureau of Construction Codes | Adopts and administers Michigan Plumbing Code | MCL 125.1504; Part 7 rules | src:usa-mi:admin-part7-plumbing; src:usa-mi:lara-code-books | partially_verified |
| Fuel Gas | ahj:usa-mi:lara-bcc-director | LARA Director's Office / Bureau of Construction Codes | Administers fuel-gas provisions through mechanical and residential code structure; a separate IFGC adoption was not verified | Part 9A Mechanical rules and residential-code structure | src:usa-mi:admin-part9a-mechanical; src:usa-mi:lara-code-books | partially_verified_scope_limited |
| Electrical | ahj:usa-mi:lara-bcc-director | LARA Director's Office / Bureau of Construction Codes | Adopts Michigan Electrical Code; local AHJ may administer/enforce as designated by governing authority | MCL 125.1504; Part 8 rules | src:usa-mi:admin-part8-electrical; src:usa-mi:lara-code-books | partially_verified |
| Energy - commercial | ahj:usa-mi:lara-bcc-director | LARA Director's Office / Bureau of Construction Codes | Adopts Michigan Commercial Energy Code | MCL 125.1504; Part 10A rules | src:usa-mi:admin-part10a-commercial-energy; src:usa-mi:lara-code-books | partially_verified |
| Energy - residential | ahj:usa-mi:lara-bcc-director | LARA Director's Office / Bureau of Construction Codes | Residential energy rules under Part 10 remain subject to 2025 stay; current residential standards remain in effect | MCL 125.1504; LARA residential-energy stay notice | src:usa-mi:lara-open-rules-res-energy-stay; src:usa-mi:lara-code-books | partially_verified_stayed_update |
| Fire - construction references | ahj:usa-mi:lara-bcc-director | LARA Director's Office / Bureau of Construction Codes | Administers fire/life-safety construction provisions embedded in building, residential, mechanical, plumbing, electrical, and energy codes | Construction-code rules | src:usa-mi:admin-part4-building; src:usa-mi:admin-part8-electrical | partially_verified |
| Fire - operational / prevention code | ahj:usa-mi:lara-bfs-state-fire-marshal | LARA Bureau of Fire Services / State Fire Marshal | Administers fire-prevention and regulated-facility programs; complete statewide operational fire-code adoption not parsed | Fire Prevention Code Act 207 and State Fire Safety Board rules need full extraction | src:usa-mi:lara-bfs-home; src:usa-mi:fire-prevention-code-act207; src:usa-mi:state-fire-safety-board-rules | partially_verified_scope_incomplete |
| Accessibility | ahj:usa-mi:lara-bcc-director | LARA Director's Office / Bureau of Construction Codes | Accessibility and barrier-free provisions are administered within state construction-code framework, but dedicated barrier-free rule parsing was not completed | Part 4 Building rules and BCC scope | src:usa-mi:admin-part4-building; src:usa-mi:lara-bcc-home | unresolved_detail |
| Elevator / Conveyance | ahj:usa-mi:lara-bcc-elevator-section | LARA BCC Elevator Section / Elevator Safety Board | Administers and enforces Michigan Elevator Code; Elevator Safety Board has rule/licensing/examination functions | 1967 PA 227; Elevator Safety Board rules and LARA elevator-code materials | src:usa-mi:lara-elevator-section; src:usa-mi:lara-elevator-permit-info; src:usa-mi:elevator-safety-board; src:usa-mi:lara-code-books | partially_verified |
| Construction-code appeals | ahj:usa-mi:construction-code-commission | State Construction Code Commission | Hears appeals under the state construction-code framework | 1972 PA 230 | src:usa-mi:construction-code-commission | partially_verified |

### 2.3 Authority Hierarchy Notes

Michigan's construction-code hierarchy is organized around LARA/BCC and administrative-rule code parts. The rules reviewed cite MCL 125.1504 as the Director's authority to promulgate state construction-code rules. LARA's BCC is the operational bureau for construction-code administration, code books, permits, and related code-rule resources.

The enforcement hierarchy is not identical to the rulemaking hierarchy. Local governmental subdivisions may serve as enforcing agencies where they have assumed responsibility, and code rules also use discipline-specific AHJ language. Because the official full text of MCL 125.1508b was not extracted directly, local enforcement should remain partially verified rather than verified.

Fire-prevention authority should not be merged into the construction-code authority without more parsing. Construction fire/life-safety provisions are administered through BCC construction codes, while operational/prevention fire programs involve LARA's Bureau of Fire Services, the State Fire Marshal, and State Fire Safety Board rules.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-mi:001 | ahj:usa-mi:lara-bcc-director | promulgates | Michigan construction-code rules under MCL 125.1504 | src:usa-mi:mcl-125-1504; src:usa-mi:admin-part4-building | partially_verified |
| edge:usa-mi:002 | ahj:usa-mi:lara-bcc-director | administers | statewide construction-code families listed in section 3 | src:usa-mi:lara-bcc-home; src:usa-mi:lara-code-books | partially_verified |
| edge:usa-mi:003 | ahj:usa-mi:lara-bcc-director | allows_local_assumption_of_enforcement | governmental subdivisions / enforcing agencies | src:usa-mi:mcl-125-1508b; src:usa-mi:bcc-building-permits | partially_verified_caveated |
| edge:usa-mi:004 | ahj:usa-mi:construction-code-commission | hears_appeals_for | construction-code disputes | src:usa-mi:construction-code-commission | partially_verified |
| edge:usa-mi:005 | ahj:usa-mi:lara-bcc-elevator-section | administers_and_enforces | Michigan Elevator Code inspections, permits, accident/violation investigation | src:usa-mi:lara-elevator-section; src:usa-mi:lara-elevator-permit-info | partially_verified |
| edge:usa-mi:006 | ahj:usa-mi:elevator-safety-board | supports_rules_and_licensing | elevator contractor licensing and board functions | src:usa-mi:elevator-safety-board | partially_verified |
| edge:usa-mi:007 | ahj:usa-mi:lara-bfs-state-fire-marshal | administers | fire-prevention and regulated-facility programs separate from construction-code-only provisions | src:usa-mi:lara-bfs-home; src:usa-mi:state-fire-safety-board-rules | partially_verified_scope_incomplete |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | 2021 Michigan Building Code | 2021 International Building Code; appendices F, G, and H except listed sections | 2021 | current | null | 2025-04-09 | 2025-04-09 | unknown | Effective rule date verified; statewide permit-date trigger unresolved | src:usa-mi:lara-code-books; src:usa-mi:admin-part4-building |
| Residential | 2015 Michigan Residential Code | 2015 International Residential Code | 2015 | current; 2021 update stayed | null | 2016-02-08 | 2016-02-08 | unknown | 2021 IRC / residential energy update stayed by 2025 court order; current 2015 code remains valid while order is in effect | src:usa-mi:lara-code-books; src:usa-mi:lara-open-rules-res-energy-stay |
| Existing Building / Rehabilitation | 2021 Michigan Rehabilitation Code | 2021 International Existing Building Code | 2021 | current | null | 2025-04-09 | 2025-04-09 | unknown | Effective rule date verified; statewide permit-date trigger unresolved | src:usa-mi:lara-code-books; src:usa-mi:admin-part4-building |
| Mechanical | 2021 Michigan Mechanical Code | 2021 International Mechanical Code | 2021 | current | null | 2024-03-12 | 2024-03-12 | unknown | Effective rule date verified; statewide permit-date trigger unresolved | src:usa-mi:lara-code-books; src:usa-mi:admin-part9a-mechanical |
| Plumbing | 2021 Michigan Plumbing Code | 2021 International Plumbing Code, second printing, including appendices B, C, D, and E except listed sections | 2021 | current | null | 2024-03-12 | 2024-03-12 | unknown | Effective rule date verified; permit expiration/reinstatement rule captured in section 4 | src:usa-mi:lara-code-books; src:usa-mi:admin-part7-plumbing |
| Fuel Gas | Michigan Mechanical Code / Michigan Residential Code fuel-gas provisions | IMC fuel-gas provisions and residential fuel-gas provisions; separate IFGC adoption not verified | mixed | partially_verified_scope_limited | null | 2024-03-12 | 2024-03-12 | unknown | Applies through mechanical and residential code structure; scope needs follow-up | src:usa-mi:admin-part9a-mechanical; src:usa-mi:lara-code-books |
| Electrical | Michigan Electrical Code | 2023 National Electrical Code, second printing, with TIA, errata, Annex H except listed sections; NFPA 110 and NFPA 111 also adopted by reference | 2023 NEC | current | null | 2024-03-12 | 2024-03-12 | unknown | Effective rule date verified; AHJ designated by governing authority under Part 8 | src:usa-mi:lara-code-books; src:usa-mi:admin-part8-electrical |
| Energy - commercial | 2021 Michigan Commercial Energy Code | 2021 IECC for nonresidential buildings; ASHRAE/IESNA 90.1-2019 adopted by reference | 2021 IECC / 90.1-2019 | current | null | 2025-04-22 | 2025-04-22 | unknown | Appendix CB mandatory; Appendix CC permissive and not mandated by enforcing agency | src:usa-mi:lara-code-books; src:usa-mi:admin-part10a-commercial-energy |
| Energy - residential | Michigan residential energy standards | 2015 residential energy standards currently remain in effect; 2021 update stayed | 2015 current; 2021 stayed | current_with_stayed_update | null | 2017-09-20 | 2017-09-20 | unknown | 2021 residential energy update set for 2025-08-29 but stayed; 2015 residential energy standards remain valid while order remains in place | src:usa-mi:lara-code-books; src:usa-mi:lara-open-rules-res-energy-stay |
| Fire - construction references | Fire/life-safety construction provisions in Michigan construction codes | IBC/IRC/IMC/IPC/NEC referenced standards and state amendments | mixed | current_within_construction_codes | null | mixed | mixed | unknown | Construction provisions track their host code family; operational/prevention fire code tracked separately | src:usa-mi:admin-part4-building; src:usa-mi:admin-part8-electrical |
| Fire - operational / prevention code | Michigan fire-prevention framework | Fire Prevention Code Act and State Fire Safety Board rules; model-code edition unresolved | unknown | scoped_unresolved | null | null | null | unknown | Operational/prevention code edition not fully reduced to a statewide adoption record | src:usa-mi:lara-bfs-home; src:usa-mi:fire-prevention-code-act207; src:usa-mi:state-fire-safety-board-rules |
| Accessibility | Michigan construction-code accessibility / barrier-free provisions | Building-code accessibility provisions and barrier-free rules not fully parsed | unknown | unresolved_detail | null | null | null | unknown | Dedicated barrier-free rule extraction required | src:usa-mi:admin-part4-building; src:usa-mi:lara-bcc-home |
| Elevator / Conveyance | Michigan Elevator Code | ASME A17.1-2016, A18.1-2017, A90.1-2009, and Elevator Safety Board General Rules | mixed standards | current | null | 2023-06-27 | 2023-06-27 | unknown | Effective date from LARA code-books page; enforcement handled by Elevator Section / appropriate enforcing agency | src:usa-mi:lara-code-books; src:usa-mi:lara-elevator-permit-info; src:usa-mi:lara-elevator-section |

### 3.2 Adoption Records

| Adoption Record ID | Code Family | State Code Name | Base Code / Standard | Edition | Effective Date | Operative Date | Citation Source IDs | Normalized Status | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| adoption:usa-mi:building-2021-mbc | Building | 2021 Michigan Building Code | 2021 IBC | 2021 | 2025-04-09 | 2025-04-09 | src:usa-mi:lara-code-books; src:usa-mi:admin-part4-building | active | Adoption date not separately extracted from final-rule filing in this pass. |
| adoption:usa-mi:rehab-2021-mrc | Existing Building / Rehabilitation | 2021 Michigan Rehabilitation Code | 2021 IEBC | 2021 | 2025-04-09 | 2025-04-09 | src:usa-mi:lara-code-books; src:usa-mi:admin-part4-building | active | Michigan uses "Rehabilitation Code" naming for existing-building/rehab scope. |
| adoption:usa-mi:mechanical-2021 | Mechanical | 2021 Michigan Mechanical Code | 2021 IMC | 2021 | 2024-03-12 | 2024-03-12 | src:usa-mi:lara-code-books; src:usa-mi:admin-part9a-mechanical | active | Part 9A rules adopt the 2021 IMC with exceptions. |
| adoption:usa-mi:plumbing-2021 | Plumbing | 2021 Michigan Plumbing Code | 2021 IPC | 2021 | 2024-03-12 | 2024-03-12 | src:usa-mi:lara-code-books; src:usa-mi:admin-part7-plumbing | active | Part 7 rules adopt the 2021 IPC second printing with listed appendices/exceptions. |
| adoption:usa-mi:electrical-2023-nec | Electrical | Michigan Electrical Code | NFPA 70 / NEC | 2023 | 2024-03-12 | 2024-03-12 | src:usa-mi:lara-code-books; src:usa-mi:admin-part8-electrical | active | Part 8 rules adopt NEC with TIA, errata, Annex H, and listed exceptions. |
| adoption:usa-mi:commercial-energy-2021 | Energy - commercial | 2021 Michigan Commercial Energy Code | 2021 IECC; ASHRAE/IESNA 90.1-2019 | 2021 / 2019 | 2025-04-22 | 2025-04-22 | src:usa-mi:lara-code-books; src:usa-mi:admin-part10a-commercial-energy | active | Appendix CB mandatory; Appendix CC permissive. |
| adoption:usa-mi:residential-2015-current | Residential | 2015 Michigan Residential Code | 2015 IRC | 2015 | 2016-02-08 | 2016-02-08 | src:usa-mi:lara-code-books; src:usa-mi:lara-open-rules-res-energy-stay | active_with_stayed_successor | LARA notice states 2015 IRC and residential IECC standards remain valid/effective while the stay remains in place. |
| adoption:usa-mi:residential-energy-2015-current | Energy - residential | Michigan residential energy standards | 2015 residential IECC standards / Michigan energy code materials | 2015 | 2017-09-20 | 2017-09-20 | src:usa-mi:lara-code-books; src:usa-mi:lara-open-rules-res-energy-stay | active_with_stayed_successor | The precise placement of residential energy provisions should be rechecked against Part 10 and MRC chapter structure. |
| adoption:usa-mi:elevator-2023-code-set | Elevator / Conveyance | Michigan Elevator Code | ASME A17.1-2016; ASME A18.1-2017; ASME A90.1-2009; Elevator Safety Board rules | mixed | 2023-06-27 | 2023-06-27 | src:usa-mi:lara-code-books; src:usa-mi:lara-elevator-permit-info | active | Effective date and standards captured from LARA code-books/elevator materials. |
| pending:usa-mi:residential-2021-stayed | Residential | Michigan Residential Code update | 2021 IRC | 2021 | 2025-08-29 | stayed | src:usa-mi:lara-open-rules-res-energy-stay; src:usa-mi:lara-open-rules | stayed_pending_litigation | LARA advisory states the rules were set to become effective 2025-08-29 but implementation is temporarily prevented. |
| pending:usa-mi:residential-energy-2021-stayed | Energy - residential | Residential energy update | 2021 residential IECC standards | 2021 | 2025-08-29 | stayed | src:usa-mi:lara-open-rules-res-energy-stay; src:usa-mi:lara-open-rules | stayed_pending_litigation | Keep active-current and stayed-successor records separate. |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

Effective dates were captured from LARA code-book listings and the administrative-rule histories. This report does not classify a statewide permit-application-date, permit-issuance-date, or mandatory-use rule for all code families. Where a rule text supplied a specific permit provision, it is captured below at the code-family level rather than generalized.

The residential-code and residential-energy transition is unusual: LARA reports that the 2021 IRC and 2021 IECC residential-standard rules were set to become effective on 2025-08-29, but a 2025-07-07 court order temporarily prevents implementation. While the order remains in place, the currently adopted 2015 IRC and residential IECC standards remain valid and effective.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-mi:building-effective-2025 | Building / Rehabilitation | effective_date | 2025-04-09 | Part 4 rule history and LARA current-code listing | unknown | src:usa-mi:admin-part4-building; src:usa-mi:lara-code-books | partially_verified |
| date-rule:usa-mi:mechanical-effective-2024 | Mechanical | effective_date | 2024-03-12 | Part 9A rule history and LARA current-code listing | unknown | src:usa-mi:admin-part9a-mechanical; src:usa-mi:lara-code-books | partially_verified |
| date-rule:usa-mi:plumbing-effective-2024 | Plumbing | effective_date | 2024-03-12 | Part 7 rule history and LARA current-code listing | unknown | src:usa-mi:admin-part7-plumbing; src:usa-mi:lara-code-books | partially_verified |
| date-rule:usa-mi:electrical-effective-2024 | Electrical | effective_date | 2024-03-12 | Part 8 rule history and LARA current-code listing | unknown | src:usa-mi:admin-part8-electrical; src:usa-mi:lara-code-books | partially_verified |
| date-rule:usa-mi:commercial-energy-effective-2025 | Energy - commercial | effective_date | 2025-04-22 | Part 10A rule history and LARA current-code listing | unknown | src:usa-mi:admin-part10a-commercial-energy; src:usa-mi:lara-code-books | partially_verified |
| date-rule:usa-mi:residential-stay-2025 | Residential and residential energy | court_stay / implementation_hold | 2025-07-07 order; 2025-08-29 effective date stayed | Court order prevents LARA from applying or implementing the 2021 residential and residential-energy rules while the order remains in effect | yes; 2015 standards remain valid/effective while stay remains | src:usa-mi:lara-open-rules-res-energy-stay | partially_verified |
| date-rule:usa-mi:plumbing-permit-expiration | Plumbing | permit_expiration | 180 days | Plumbing permit expires if work is not commenced within 180 days after permit issuance, or if work is suspended or abandoned for 180 days after commencement | conditional; if the code changed and work was not started, a new permit must be obtained before resuming work | src:usa-mi:admin-part7-plumbing | partially_verified |
| date-rule:usa-mi:statewide-permit-trigger | All construction-code families | permit_date_or_issuance_trigger | unknown | Statewide trigger for which code edition applies to a submitted, issued, expired, or reinstated permit was not fully extracted | unknown | src:usa-mi:admin-part4-building; src:usa-mi:bcc-building-permits | unresolved |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Residential | 2021 Michigan Residential Code / 2021 IRC | 2025 LARA advisory | null | 2025-08-29 | stayed | stayed | active_watch | src:usa-mi:lara-open-rules-res-energy-stay; src:usa-mi:lara-open-rules | Rules were set for 2025-08-29 but implementation is stayed by court order. |
| Energy - residential | 2021 residential IECC standards | 2025 LARA advisory | null | 2025-08-29 | stayed | stayed | active_watch | src:usa-mi:lara-open-rules-res-energy-stay; src:usa-mi:lara-open-rules | Same stay affects residential energy update. |
| Building / Rehabilitation / Commercial Energy | Proposed Part 4, Rehabilitation Code, and Part 10A rule drafts | 2026-05-21 public advisory meeting | null | null | null | null | monitor | src:usa-mi:lara-open-rules | Proposed-rule stage only; do not treat as adopted. |
| Plumbing | Proposed 2024 Michigan Plumbing Code update / 2024 IPC | 2026-08-05 public hearing listed | null | 120 days after filing with Secretary of State, if adopted | null | null | monitor | src:usa-mi:lara-open-rules | Proposed-rule stage only. |
| Mechanical | Proposed Part 9A Mechanical Code rule update | 2026 open-rules listing | null | 120 days after filing with Secretary of State, if adopted | null | null | monitor | src:usa-mi:lara-open-rules | Proposed-rule stage only. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-mi:building-permit-before-work | Building | Construction, alteration, demolition, change of occupancy, and similar regulated work | Owner or authorized agent intends regulated work | A permit is required from the building official before regulated building work, subject to ordinary-repair exceptions. | src:usa-mi:admin-part4-building; src:usa-mi:bcc-building-permits | partially_verified |
| applicability-rule:usa-mi:electrical-ahj | Electrical | Electrical code administration/enforcement | Governing authority designates electrical inspector or designee | Part 8 defines AHJ to include the electrical inspector or other governing-body designee and says the code is administered/enforced by the designated AHJ. | src:usa-mi:admin-part8-electrical | partially_verified |
| applicability-rule:usa-mi:commercial-energy-appendices | Energy - commercial | IECC appendices | Project uses commercial energy code compliance path | Part 10A makes Appendix CB mandatory; Appendix CC is permissive and not mandated by the enforcement agency. | src:usa-mi:admin-part10a-commercial-energy | partially_verified |
| applicability-rule:usa-mi:local-resolution-building-exceptions | Building / fire-safety construction provisions | Certain existing-building high-rise and sprinkler exceptions | Local unit has qualifying municipal fire department and governing body resolution filed with LARA/BCC | Certain exceptions require local-government resolution filed with LARA/BCC; this is a specific rule option, not proof of general local amendment authority. | src:usa-mi:admin-part4-building | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** Michigan adopts model codes and standards by reference through Michigan Administrative Rules code parts, then amends, deletes, or supplements specific model-code provisions in those rules.

**Where amendments are published:** Michigan Administrative Rules / Administrative Rules System, Michigan Register histories, LARA/BCC code-book pages, and official LARA integrated-code or code-book resources where available.

**Amendment parsing status:** high_level_parsed; substantive line-by-line amendment extraction remains incomplete.

### 5.2 State Amendment Sources

| Amendment Source ID | Code Family | Publication Path | Coverage | Extraction Status | Source IDs |
| --- | --- | --- | --- | --- | --- |
| amend-source:usa-mi:part4-building | Building / Rehabilitation / construction references | Michigan Administrative Rules Part 4 | 2021 IBC adoption and state amendments; selected local-resolution exceptions | partially_parsed | src:usa-mi:admin-part4-building |
| amend-source:usa-mi:part7-plumbing | Plumbing | Michigan Administrative Rules Part 7 | 2021 IPC adoption, state amendments, permit provisions | partially_parsed | src:usa-mi:admin-part7-plumbing |
| amend-source:usa-mi:part8-electrical | Electrical | Michigan Administrative Rules Part 8 | 2023 NEC adoption, exceptions, NFPA 110/111 references, AHJ language | partially_parsed | src:usa-mi:admin-part8-electrical |
| amend-source:usa-mi:part9a-mechanical | Mechanical / Fuel Gas | Michigan Administrative Rules Part 9A | 2021 IMC adoption and exceptions | partially_parsed | src:usa-mi:admin-part9a-mechanical |
| amend-source:usa-mi:part10a-commercial-energy | Energy - commercial | Michigan Administrative Rules Part 10A | 2021 IECC / ASHRAE 90.1 adoption; Appendix CB/CC treatment | partially_parsed | src:usa-mi:admin-part10a-commercial-energy |
| amend-source:usa-mi:residential-stayed-update | Residential / residential energy | LARA open-rules and advisory materials | Stayed 2021 IRC / residential energy update | high_level_only | src:usa-mi:lara-open-rules-res-energy-stay; src:usa-mi:lara-open-rules |
| amend-source:usa-mi:elevator | Elevator / Conveyance | LARA code-books and elevator-code pages | Elevator code standards and enforcement context | high_level_only | src:usa-mi:lara-code-books; src:usa-mi:lara-elevator-permit-info; src:usa-mi:lara-elevator-section |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Topic | Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| amend:usa-mi:building-act-precedence | Building / construction codes | Act precedence | Part 4 states that the act takes precedence over all code provisions. | src:usa-mi:admin-part4-building | partially_verified |
| amend:usa-mi:building-local-resolution-exceptions | Building / fire-safety construction provisions | Local-resolution exceptions | Certain existing-building exceptions require a qualifying local fire department and a local governing-body resolution filed with LARA/BCC. | src:usa-mi:admin-part4-building | partially_verified |
| amend:usa-mi:electrical-nec-scope | Electrical | NEC adoption with exceptions | Part 8 adopts the 2023 NEC with TIA, errata, Annex H, and listed exceptions, and also adopts NFPA 110/111 by reference. | src:usa-mi:admin-part8-electrical | partially_verified |
| amend:usa-mi:commercial-energy-appendix-cb-cc | Energy - commercial | IECC appendices | Appendix CB is mandatory; Appendix CC is permissive and not mandated by the enforcement agency. | src:usa-mi:admin-part10a-commercial-energy | partially_verified |
| amend:usa-mi:plumbing-permit-expiration | Plumbing | Permit expiration and changed-code effect | A plumbing permit expires after 180 days without commencement or after 180 days of suspension/abandonment; if code changes and work was not started, a new permit is required before resuming. | src:usa-mi:admin-part7-plumbing | partially_verified |
| amend:usa-mi:residential-stayed-update | Residential / residential energy | Stayed 2021 update | The planned 2021 IRC and residential-energy update is not implemented while the court order remains in effect. | src:usa-mi:lara-open-rules-res-energy-stay | partially_verified |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-mi"
  model: "state-default_with_local_assumption"
  enforcing_entities:
    - "LARA/BCC or the director where the state retains/defaults enforcement responsibility"
    - "counties, cities, villages, or townships that have assumed responsibility and designated an enforcing agency"
    - "discipline-specific AHJs such as electrical inspectors or other governing-body designees under Part 8"
    - "Elevator Section or appropriate enforcing agency for elevator installation/alteration plan and permit workflow"
  required_officials:
    - "building official / appropriate enforcing agency"
    - "electrical inspector or other designee of the governing body for electrical code administration"
    - "licensed elevator contractor and licensed elevator journeyperson for elevator installation/alteration work"
  state_reserved_activities:
    - "state construction-code rulemaking and code adoption"
    - "state-level BCC administration where local responsibility has not been assumed or where state jurisdiction is retained"
    - "elevator-code administration, inspection, certificate, accident, and violation functions through LARA/BCC Elevator Section"
  source_ids:
    - "src:usa-mi:mcl-125-1508b"
    - "src:usa-mi:bcc-building-permits"
    - "src:usa-mi:admin-part8-electrical"
    - "src:usa-mi:lara-elevator-section"
    - "src:usa-mi:lara-elevator-permit-info"
  verification_status: "partially_verified_caveated"
  confidence: 0.70
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-mi"
  model: "unresolved_single_state_code_with_rule_specific_local_options"
  applies_to_code_families:
    - "construction codes"
    - "fire-prevention and operational fire-code scope unresolved"
  approval_required: "unknown"
  approving_authority_id: null
  filing_required: "verified only for specific Part 4 building-code exceptions that require local resolution filed with LARA/BCC"
  registry_exists: "unknown"
  registry_source_ids: []
  legal_basis_source_ids:
    - "src:usa-mi:admin-part4-building"
    - "src:usa-mi:mcl-125-1508b"
  verification_status: "unresolved_detail"
  confidence: 0.35
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Michigan local enforcement should be treated separately from local amendment power. The sources reviewed support local administration/enforcement pathways, discipline-specific AHJ roles, and certain rule-specific local-resolution options. They do not yet establish a general local technical-amendment authority, a statewide local-amendment registry, or a complete preemption rule. The report therefore classifies local enforcement as partially verified and local amendment authority as unresolved.

### 6.4 Known Local Amendment Registries

| Registry ID | Name | Scope | Source IDs | Status | Notes |
| --- | --- | --- | --- | --- | --- |
| registry:usa-mi:local-amendments | Local construction-code amendments registry | unknown | none | unresolved | No statewide local-amendment registry was verified. |
| registry:usa-mi:part4-local-resolutions | Part 4 local resolutions filed with LARA/BCC | specific high-rise/sprinkler exceptions | src:usa-mi:admin-part4-building | partially_verified_existence_only | The rules require filing of specific resolutions, but no public registry was located in this pass. |

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendments were parsed. This absence is not evidence that no local amendments or local code ordinances exist.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: partially_started

Jurisdiction stack:

```text
Address
  -> State of Michigan
  -> County
  -> Municipality / township / village / city / unincorporated county
  -> Local governmental subdivision enforcement-assumption status
  -> BCC/state default enforcement where local responsibility has not been assumed or where state jurisdiction is retained
  -> Building official / appropriate enforcing agency
  -> Trade-specific AHJs, including electrical inspector or other governing-body designee
  -> Fire AHJ / Bureau of Fire Services program authority where applicable
  -> Elevator Section / appropriate enforcing agency for conveyances
  -> Applicable state code adoption records
  -> Applicable local resolution or amendment records, if any
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | not selected | none | statewide | unknown | pending |
| County | not selected | none | statewide | unknown | pending |
| City / village / township | not selected | none | statewide | unknown | pending |
| Local enforcing agency status | not selected | none | statewide | unknown | pending |
| Fire district / fire department service area | not selected | none | local / program-specific | unknown | pending |
| Elevator enforcing agency / municipal elevator department | not selected | none | statewide plus approved local programs if any | unknown | pending |

### 7.3 AHJ Contact Data

No AHJ contact data was populated. A production resolver should add BCC enforcing-agency listings, local building department contacts, fire AHJ data, and elevator enforcing-agency coverage if available from official sources.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Source Type | Title | Publisher / Authority | URL | Supports | Date / Version | Extraction Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-mi:lara-bcc-home | agency_page | Bureau of Construction Codes | Michigan LARA | https://www.michigan.gov/lara/bureau-list/bcc | BCC role and mission | current page accessed 2026-06-25 | parsed_summary |
| src:usa-mi:lara-code-books | agency_page | Code Books | Michigan LARA / BCC | https://www.michigan.gov/lara/bureau-list/bcc/rules-acts/codes/code-books | Current code editions and effective dates | current page accessed 2026-06-25 | parsed_key_fields |
| src:usa-mi:mcl-125-1504 | statute | MCL 125.1504, State construction code; rules | Michigan Legislature | https://www.legislature.mi.gov/Laws/MCL?objectName=mcl-125-1504 | State construction-code rulemaking authority | official MCL target | indirect_extract_from_rules |
| src:usa-mi:mcl-125-1508b | statute | MCL 125.1508b, administration and enforcement / local assumption provisions | Michigan Legislature | https://www.legislature.mi.gov/Laws/MCL?objectName=mcl-125-1508b | Local enforcement model | official MCL target | search_surface_only |
| src:usa-mi:admin-part4-building | administrative_rule | Michigan Administrative Rules Part 4, Building Code | Michigan Administrative Rules / LARA | https://ars.apps.lara.state.mi.us/AdminCode/DownloadAdminCodeFile?FileName=R+408.30401++to+408.30499a.pdf&ReturnHTML=True | 2021 IBC adoption, building-code purpose, authority, effective date, local-resolution exceptions | 2024 MR 23; effective 2025-04-09 | parsed_key_fields |
| src:usa-mi:admin-part7-plumbing | administrative_rule | Michigan Administrative Rules Part 7, Plumbing Code | Michigan Administrative Rules / LARA | https://ars.apps.lara.state.mi.us/AdminCode/DownloadAdminCodeFile?FileName=R+408.30701++to+408.30796.pdf&ReturnHTML=True | 2021 IPC adoption, purpose/scope, effective date, permit expiration | 2023 MR 22; effective 2024-03-12 | parsed_key_fields |
| src:usa-mi:admin-part8-electrical | administrative_rule | Michigan Administrative Rules Part 8, Electrical Code | Michigan Administrative Rules / LARA | https://ars.apps.lara.state.mi.us/AdminCode/DownloadAdminCodeFile?FileName=R+408.30801++to+408.30880.pdf&ReturnHTML=True | 2023 NEC adoption, AHJ definition, effective date, scope | 2023 MR 22; effective 2024-03-12 | parsed_key_fields |
| src:usa-mi:admin-part9a-mechanical | administrative_rule | Michigan Administrative Rules Part 9A, Mechanical Code | Michigan Administrative Rules / LARA | https://ars.apps.lara.state.mi.us/AdminCode/DownloadAdminCodeFile?FileName=R+408.30901++to+408.30998.pdf&ReturnHTML=True | 2021 IMC adoption, scope, effective date | 2023 MR 22; effective 2024-03-12 | parsed_key_fields |
| src:usa-mi:admin-part10a-commercial-energy | administrative_rule | Michigan Administrative Rules Part 10A, Michigan Energy Code | Michigan Administrative Rules / LARA | https://ars.apps.lara.state.mi.us/AdminCode/DownloadAdminCodeFile?FileName=R+408.31087++to+408.31099.pdf&ReturnHTML=True | 2021 IECC commercial adoption, ASHRAE 90.1-2019, Appendix CB/CC, effective date | 2024 MR 23; effective 2025-04-22 | parsed_key_fields |
| src:usa-mi:lara-open-rules-res-energy-stay | agency_page_and_advisory | 2021 Residential Code / Residential Energy Code injunction advisory | Michigan LARA / BCC | https://www.michigan.gov/lara/bureau-list/bcc/rules-acts/rules/currently-open | Stayed 2021 residential and residential-energy updates; 2015 standards remain valid/effective | LARA advisory referencing 2025-07-07 order | parsed_key_fields |
| src:usa-mi:lara-open-rules | agency_page | Currently Open Rules | Michigan LARA / BCC | https://www.michigan.gov/lara/bureau-list/bcc/rules-acts/rules/currently-open | Pending rulemaking and hearing/advisory dates | current page accessed 2026-06-25 | parsed_key_fields |
| src:usa-mi:bcc-building-permits | agency_page | Building Permit Information | Michigan LARA / BCC | https://www.michigan.gov/lara/bureau-list/bcc/sections/permits/building-permit-information | Permit-before-work overview and building-code scope | current page accessed 2026-06-25 | parsed_summary |
| src:usa-mi:construction-code-commission | agency_page | State Construction Code Commission | Michigan LARA / BCC | https://www.michigan.gov/lara/bureau-list/bcc/boards/state-construction-code-commission | Commission creation and appeal function | current page accessed 2026-06-25 | parsed_summary |
| src:usa-mi:lara-bfs-home | agency_page | Bureau of Fire Services | Michigan LARA | https://www.michigan.gov/lara/bureau-list/bfs | Fire-prevention mission and BFS context | current page accessed 2026-06-25 | parsed_summary |
| src:usa-mi:fire-prevention-code-act207 | statute | Fire Prevention Code Act 207 of 1941 | Michigan Legislature / official PDF target | https://www.legislature.mi.gov/documents/mcl/pdf/mcl-act-207-of-1941.pdf | Fire-prevention statutory framework | official act PDF target | not_fully_parsed |
| src:usa-mi:state-fire-safety-board-rules | administrative_rule_index | State Fire Safety Board rules index | Michigan Administrative Rules / LARA | https://ars.apps.lara.state.mi.us/AdminCode/DeptBureauAdminCode?Department=Licensing+and+Regulatory+Affairs&Bureau=State+Fire+Safety+Board | Fire-safety administrative rules and review/hearing process | current ARS index accessed 2026-06-25 | parsed_summary |
| src:usa-mi:lara-elevator-permit-info | agency_page | Elevator Permit Information | Michigan LARA / BCC | https://www.michigan.gov/lara/bureau-list/bcc/sections/permits/elevator-permit-information | Elevator-code scope, standards, permit pathway | current page accessed 2026-06-25 | parsed_summary |
| src:usa-mi:lara-elevator-section | agency_page | Elevator Section | Michigan LARA / BCC | https://www.michigan.gov/lara/bureau-list/bcc/sections/elevator-section | Elevator-code administration and enforcement | current page accessed 2026-06-25 | parsed_summary |
| src:usa-mi:elevator-safety-board | agency_page | Elevator Safety Board | Michigan LARA / BCC | https://www.michigan.gov/lara/bureau-list/bcc/boards/elevator-safety-board | Elevator board creation, rule/licensing/exam functions | current page accessed 2026-06-25 | parsed_summary |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-mi:mcl-125-1504 | access_limitation | Official MCL page was used as registry target, but the rule authority text was extracted from official administrative rules citing the statute because the MCL page was not directly retrievable in this pass. | validate_directly_before_verified |
| src:usa-mi:mcl-125-1508b | access_limitation | Local-enforcement classification relies on search-surface official snippets and supporting BCC/AHJ sources; full official statutory text was not extracted. | high_priority_reextract |
| src:usa-mi:admin-part4-building | courtesy_html | ARS HTML is a courtesy rendering of rule/PDF text and includes OCR/formatting artifacts; key fields should be checked against official PDF for litigation-grade use. | acceptable_for_partial_verification |
| src:usa-mi:admin-part7-plumbing | courtesy_html | ARS HTML is a courtesy rendering of rule/PDF text and includes formatting artifacts. | acceptable_for_partial_verification |
| src:usa-mi:admin-part8-electrical | courtesy_html | ARS HTML is a courtesy rendering of rule/PDF text and includes formatting artifacts. | acceptable_for_partial_verification |
| src:usa-mi:admin-part9a-mechanical | courtesy_html | ARS HTML is a courtesy rendering of rule/PDF text and includes formatting artifacts. | acceptable_for_partial_verification |
| src:usa-mi:admin-part10a-commercial-energy | courtesy_html | ARS HTML is a courtesy rendering of rule/PDF text and includes formatting artifacts. | acceptable_for_partial_verification |
| src:usa-mi:lara-code-books | agency_summary | LARA's code-books page is an official agency convenience page for current codes/effective dates, not the full legal text of every rule. | pair_with_rules_for_verified_status |
| src:usa-mi:lara-open-rules-res-energy-stay | litigation_status | The residential/residential-energy stay is time-sensitive and may change with further court or agency action. | monitor_frequently |
| src:usa-mi:fire-prevention-code-act207 | incomplete_extraction | Act 207 was identified as an official fire-prevention framework source, but full text and implementing rule/code edition were not parsed. | unresolved_fire_scope |

### 8.3 Supplemental Sources

None used for report facts. Search snippets and third-party references were used only as leads where noted in caveats; they were not assigned source IDs for primary conclusions.

### 8.4 Source Extraction Metadata

| Extraction ID | Date | Extractor | Scope | Method | Result |
| --- | --- | --- | --- | --- | --- |
| extract:usa-mi:2026-06-25:code-books | 2026-06-25 | ChatGPT | LARA code-books current code editions and effective dates | Web fetch and line review | Key code families populated |
| extract:usa-mi:2026-06-25:admin-rules | 2026-06-25 | ChatGPT | Part 4, Part 7, Part 8, Part 9A, Part 10A administrative rules | Web fetch of ARS HTML renderings | Authority, adoptions, effective dates, and selected amendments captured |
| extract:usa-mi:2026-06-25:res-stay | 2026-06-25 | ChatGPT | Residential/residential-energy injunction advisory | LARA open-rules page and advisory PDF screenshot review | Stayed update and current-code continuity captured |
| extract:usa-mi:2026-06-25:authority-local | 2026-06-25 | ChatGPT | Authority hierarchy, local enforcement, local amendment scoping | Agency pages, rule texts, MCL search-surface leads | Authority partially verified; local amendment unresolved |
| extract:usa-mi:2026-06-25:fire-elevator | 2026-06-25 | ChatGPT | Fire-prevention and elevator authority overview | Agency pages and ARS index review | Elevator partially verified; operational fire code unresolved |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 1.00 | none | Body contains official source registry and source-backed core fields, but unresolved items remain. |
| state | state.state_id | US-MI | verified | 1.00 | none | Preserved from state-specific uploaded draft. |
| authority:primary | authority_id | ahj:usa-mi:lara-bcc-director | partially_verified | 0.85 | src:usa-mi:mcl-125-1504; src:usa-mi:admin-part4-building; src:usa-mi:lara-bcc-home | MCL page requires direct extraction before verified status. |
| adoption:building | effective_date | 2025-04-09 | verified | 0.95 | src:usa-mi:lara-code-books; src:usa-mi:admin-part4-building | Current-code page and rule history align. |
| adoption:residential | current_code | 2015 Michigan Residential Code | partially_verified | 0.90 | src:usa-mi:lara-code-books; src:usa-mi:lara-open-rules-res-energy-stay | Current code verified; successor update stayed. |
| adoption:mechanical | effective_date | 2024-03-12 | verified | 0.92 | src:usa-mi:lara-code-books; src:usa-mi:admin-part9a-mechanical | Current-code page and rule history align. |
| adoption:plumbing | effective_date | 2024-03-12 | verified | 0.92 | src:usa-mi:lara-code-books; src:usa-mi:admin-part7-plumbing | Current-code page and rule history align. |
| adoption:electrical | effective_date | 2024-03-12 | verified | 0.92 | src:usa-mi:lara-code-books; src:usa-mi:admin-part8-electrical | Current-code page and rule history align. |
| adoption:commercial-energy | effective_date | 2025-04-22 | verified | 0.92 | src:usa-mi:lara-code-books; src:usa-mi:admin-part10a-commercial-energy | Current-code page and rule history align. |
| adoption:elevator | effective_date | 2023-06-27 | partially_verified | 0.85 | src:usa-mi:lara-code-books; src:usa-mi:lara-elevator-permit-info; src:usa-mi:lara-elevator-section | Standards captured from LARA pages; detailed rules not fully parsed. |
| local-enforcement | model | state-default_with_local_assumption | partially_verified_caveated | 0.70 | src:usa-mi:mcl-125-1508b; src:usa-mi:bcc-building-permits; src:usa-mi:admin-part8-electrical | Full MCL 125.1508b text extraction remains open. |
| local-amendment | model | unresolved_single_state_code_with_rule_specific_local_options | unresolved_detail | 0.35 | src:usa-mi:admin-part4-building; src:usa-mi:mcl-125-1508b | Rule-specific local filing captured; general local amendment authority unresolved. |
| fire-operational | code_edition | unknown | unresolved_detail | 0.55 | src:usa-mi:lara-bfs-home; src:usa-mi:fire-prevention-code-act207; src:usa-mi:state-fire-safety-board-rules | Need full Act 207 and implementing rule parse. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Every body source ID appears in section 8. |
| All authority IDs resolve | pass | Authorities used in section 2 are defined in section 2. |
| All current code families have adoption rows | pass | Rows are present for construction, residential, trades, energy, fire construction, fire operational, accessibility, and elevators. |
| Building and operational fire code are separated | pass | Construction fire/life-safety references are separate from operational/prevention fire-code authority. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Dates are stored in separate fields; unknown dates remain explicit. |
| Effective dates are valid ISO dates | pass | Known effective dates use YYYY-MM-DD format. |
| No impossible date sequences | pass | No known effective/operative-date conflict introduced. |
| Transition rules have explicit trigger conditions | pass | Captured date rules include trigger conditions; statewide permit-trigger gap remains explicitly unresolved. |
| Permit-date logic is captured where applicable | fail | A statewide permit-date or issuance-date code-version trigger was not verified. |
| Local enforcement model classified | pass | Classified as state-default_with_local_assumption, with caveat. |
| Local amendment rule classified | fail | General local amendment authority remains unresolved. |
| AHJ confirmation metadata present | fail | No AHJ contact or coverage dataset was populated. |
| Official-source caveats captured | pass | Caveats cover MCL access, ARS courtesy renderings, agency summaries, litigation status, and fire-source incompleteness. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-mi:001 | high | MCL 125.1508b local enforcement | Local enforcement classification relies on partial official snippets and supporting sources rather than full direct statutory extraction. | Extract complete official MCL 125.1508b text and update section 6. | null | null | open |
| issue:usa-mi:002 | high | local amendment authority | General local amendment power, preemption scope, approval/filing requirements, and registry status are not resolved. | Extract Act 230 amendment/preemption provisions and BCC local-code guidance. | null | null | open |
| issue:usa-mi:003 | high | operational fire code | Fire-prevention authority is identified, but no complete operational fire-code adoption record was parsed. | Extract Act 207, State Fire Safety Board rules, BFS program rules, and any statewide IFC/NFPA adoption references. | null | null | open |
| issue:usa-mi:004 | medium | statewide permit-date transition | Effective dates are known, but statewide permit application / permit issuance / grace-period rule was not fully verified. | Parse code-specific administrative rules and BCC transition guidance for current code adoptions. | null | null | open |
| issue:usa-mi:005 | high | residential and residential energy stay | 2021 residential and residential-energy updates are stayed; status may change quickly. | Monitor LARA open-rules advisory and court/order updates; update effective/operative status promptly. | null | null | open |
| issue:usa-mi:006 | medium | accessibility and elevator detail | Accessibility/barrier-free rules and elevator municipal/enforcing-agency exceptions need deeper parsing. | Extract barrier-free rules, elevator rules, and municipal elevator-department authority if applicable. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-mi:lara-code-books | src:usa-mi:lara-code-books | html_diff | monthly | Current-code edition, effective date, or code-book link changes | 2026-06-25 | active |
| watch:usa-mi:admin-part4 | src:usa-mi:admin-part4-building | html_diff | monthly | Part 4 rule history or adopted-code text changes | 2026-06-25 | active |
| watch:usa-mi:admin-trades | src:usa-mi:admin-part7-plumbing; src:usa-mi:admin-part8-electrical; src:usa-mi:admin-part9a-mechanical | html_diff | monthly | Trade-code rule histories or adopted editions change | 2026-06-25 | active |
| watch:usa-mi:energy-rules | src:usa-mi:admin-part10a-commercial-energy; src:usa-mi:lara-open-rules-res-energy-stay | html_diff | weekly | Commercial energy rule changes or residential stay status changes | 2026-06-25 | active |
| watch:usa-mi:open-rules | src:usa-mi:lara-open-rules | html_diff | weekly | New hearing notices, filings, or final-rule effective dates | 2026-06-25 | active |
| watch:usa-mi:mcl-act230 | src:usa-mi:mcl-125-1504; src:usa-mi:mcl-125-1508b | statute_diff | quarterly | Act 230 statutory amendments or official text changes | 2026-06-25 | active |
| watch:usa-mi:fire | src:usa-mi:lara-bfs-home; src:usa-mi:fire-prevention-code-act207; src:usa-mi:state-fire-safety-board-rules | html_diff | quarterly | Fire-prevention rules, Act 207 changes, or program-code adoption changes | 2026-06-25 | pending_parse |
| watch:usa-mi:elevator | src:usa-mi:lara-code-books; src:usa-mi:lara-elevator-section; src:usa-mi:elevator-safety-board | html_diff | quarterly | Elevator standards, rule histories, or enforcement guidance changes | 2026-06-25 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-24 | Generated baseline draft report | report:usa-mi | none | Codex | Original uploaded draft had placeholders and no primary-source registry. |
| 2026-06-25 | Populated Michigan state report with official LARA/BCC, administrative-rule, elevator, fire, and open-rule sources | report:usa-mi; authority:primary; adoption:*; local-enforcement:usa-mi; local-amendment-rule:usa-mi | src:usa-mi:lara-code-books; src:usa-mi:admin-part4-building; src:usa-mi:admin-part7-plumbing; src:usa-mi:admin-part8-electrical; src:usa-mi:admin-part9a-mechanical; src:usa-mi:admin-part10a-commercial-energy; src:usa-mi:lara-open-rules-res-energy-stay | ChatGPT | Upgraded to partially_verified; kept fire, local amendment, permit-date trigger, and MCL 125.1508b full-text gaps explicit. |
