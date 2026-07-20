---
state:
  state_id: "US-IA"
  name: "Iowa"
  abbreviation: "IA"
report:
  report_id: "state-report:usa-ia"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-25"
  last_verified: null
  reviewed_by: null
risk:
  overall_confidence: 0.68 # 0.00 - 1.00
  risk_flags:
    - "local_amendment_scope_partially_unresolved"
    - "local_ahj_contacts_not_populated"
    - "factory_built_chapters_not_parsed"
    - "special_occupancy_fire_rules_not_exhaustively_parsed"
  open_questions_count: 6

---

# State Building Code Authority Report: Iowa

## 1. Executive Summary

- **Authority model:** Iowa uses a hybrid state/local authority model. The Department of Inspections, Appeals, and Licensing (DIAL) director serves as, or may designate, the State Building Code Commissioner. The commissioner adopts and amends the State Building Code with approval of the Building Code Advisory Council. Distinct state authorities administer fire-control rules, electrical inspections, plumbing/mechanical systems, and conveyance safety.

- **Statewide code status:** Iowa has a State Building Code, but it is not a universal statewide building code for every construction project. Iowa Code chapter 103A applies the state code to state-owned buildings, local jurisdictions that adopt the state code by ordinance, certain state-appropriated projects, certain larger cities without a qualifying local code, factory-built structures, and statewide energy/accessibility requirements. Current verified adoptions include the 2024 IBC, 2024 IRC, 2024 IEBC, 2024 IFC, 2023 NEC, 2024 UPC, 2024 IMC, and 2012 IECC energy provisions.

- **Local enforcement model:** Local enforcement is code-family specific. For building-code projects, local jurisdictions enforce locally adopted codes and may submit or review projects depending on whether a local building department exists. Plumbing and mechanical permit/inspection/enforcement are primarily local where local jurisdictions administer those programs. Electrical inspections may be handled by political subdivisions that perform electrical inspections, while DIAL administers the state electrical inspection program for other covered work and state-owned property. Fire-code compliance can be recognized through local ordinances meeting state-rule criteria.

- **Local amendment posture:** Local amendment authority is mixed. Plumbing and mechanical rules expressly bar local adoption of different model codes but allow stricter local amendments if copies are provided to the board. Fire rules recognize certain local fire ordinances if they incorporate an approved model fire code and include plan review and regular inspections. General local amendment scope for building-code provisions beyond those express rules remains partially unresolved.

- **Known transition periods or pending changes:** Iowa's 2025 update to 481—Chapter 301 includes a 12-month transition for projects subject to Part 1 or Part 3 that are commenced within 12 months after the revised standards' effective date. The effective date reflected in Chapter 301 is 2025-09-10, so the transition window runs through 2026-09-10 for covered projects. No later statewide code change was verified in this pass.

- **Production readiness:** partial_validation_ready

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: The DIAL director serves as, or designates, the State Building Code Commissioner;
    the commissioner adopts and amends the State Building Code with Building Code
    Advisory Council approval.
  confidence: 0.88
  source_ids:
  - src:usa-ia:code-103a
  - src:usa-ia:iac-481-300
- topic: State code applicability
  finding: The State Building Code applies to specified state, local-adoption, state-funded,
    city-size, factory-built, energy, and accessibility contexts rather than to all
    private construction by default.
  confidence: 0.82
  source_ids:
  - src:usa-ia:code-103a
  - src:usa-ia:iac-481-301
- topic: Primary building/residential/existing code editions
  finding: Chapter 301 adopts the 2024 IBC, 2024 IRC, and 2024 IEBC with Iowa amendments.
  confidence: 0.88
  source_ids:
  - src:usa-ia:iac-481-301
- topic: Energy code edition
  finding: Iowa's verified statewide energy rules in Chapter 301 adopt the 2012 IECC
    Residential Provisions and 2012 IECC Commercial Provisions.
  confidence: 0.86
  source_ids:
  - src:usa-ia:code-103a
  - src:usa-ia:iac-481-301
- topic: Electrical code authority and edition
  finding: The state electrical inspection program is in DIAL; the 2023 NEC is adopted
    for covered electrical installations.
  confidence: 0.82
  source_ids:
  - src:usa-ia:iac-481-404
- topic: Fire code authority and edition
  finding: DIAL fire-control rules are adopted in consultation with the State Fire
    Marshal; Chapter 280 adopts the 2024 IFC with amendments.
  confidence: 0.82
  source_ids:
  - src:usa-ia:code-10a-511
  - src:usa-ia:iac-481-280
- topic: Plumbing/mechanical authority and editions
  finding: The Plumbing and Mechanical Systems Board rules adopt the 2024 UPC and
    2024 IMC; local jurisdictions administer permits and inspections unless another
    law provides otherwise.
  confidence: 0.84
  source_ids:
  - src:usa-ia:iac-481-425
  - src:usa-ia:iac-481-426
- topic: Transition rule
  finding: Chapter 301 permits covered Part 1 and Part 3 projects commenced within
    12 months after the revised standards' effective date to use either the revised
    or prior standards.
  confidence: 0.84
  source_ids:
  - src:usa-ia:iac-481-301
- topic: Local amendments
  finding: Plumbing/mechanical local amendments are partly resolved; general building
    local-amendment scope and a public registry remain unresolved.
  confidence: 0.5
  source_ids:
  - src:usa-ia:iac-481-425
  - src:usa-ia:iac-481-426
  - src:usa-ia:iac-481-280
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-ia:dial-building-code-commissioner |
| Authority name | State Building Code Commissioner; DIAL Director or designee |
| Authority type | statewide building-code adopting and administrative authority |
| Legal basis | Iowa Code chapter 103A, including sections 103A.4 and 103A.7; 481—Chapter 300 |
| Role | Adopts, amends, and promulgates the State Building Code with Building Code Advisory Council approval; administers plan review and inspections for state-code projects identified by statute and rule. |
| Enforcement model | hybrid_state_and_local |
| Source IDs | src:usa-ia:code-103a; src:usa-ia:iac-481-300 |
| Verification status | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-ia:dial-building-code-commissioner | State Building Code Commissioner / DIAL | Adopts and administers State Building Code building provisions, including the 2024 IBC as amended. | Iowa Code ch. 103A; 481—Ch. 300; 481—Ch. 301 | src:usa-ia:code-103a; src:usa-ia:iac-481-300; src:usa-ia:iac-481-301 | verified_core |
| Residential | ahj:usa-ia:dial-building-code-commissioner | State Building Code Commissioner / DIAL | Adopts 2024 IRC provisions for one- and two-family dwellings and townhouses three stories or less, as amended. | 481—301.8 | src:usa-ia:iac-481-301 | verified_core |
| Existing Building / Rehabilitation | ahj:usa-ia:dial-building-code-commissioner | State Building Code Commissioner / DIAL | Adopts 2024 IEBC for repair, alteration, change of occupancy, addition, and relocation of existing buildings. | 481—301.7 | src:usa-ia:iac-481-301 | verified_core |
| Mechanical | ahj:usa-ia:plumbing-mechanical-systems-board | Plumbing and Mechanical Systems Board | Adopts the State Mechanical Code; local jurisdictions may administer permits, inspections, testing, and enforcement. | Iowa Code ch. 105; 481—Ch. 426 | src:usa-ia:iac-481-301; src:usa-ia:iac-481-426 | verified_core |
| Plumbing | ahj:usa-ia:plumbing-mechanical-systems-board | Plumbing and Mechanical Systems Board | Adopts the State Plumbing Code; local jurisdictions may administer permits, inspections, testing, and enforcement. | Iowa Code ch. 105; 481—Ch. 425 | src:usa-ia:iac-481-301; src:usa-ia:iac-481-425 | verified_core |
| Fuel Gas | ahj:usa-ia:plumbing-mechanical-systems-board | Plumbing and Mechanical Systems Board, with fire/LPG cross-references | Fuel-gas piping is tied to plumbing/mechanical rules and referenced NFPA/LPG rules; LPG program details were not fully parsed. | 481—301.9; 481—425.3; 481—426.3 | src:usa-ia:iac-481-301; src:usa-ia:iac-481-425; src:usa-ia:iac-481-426 | partially_verified |
| Electrical | ahj:usa-ia:dial-electrical-inspection-program | DIAL Electrical Inspection Program / Electrical Examining Board | Enforces requirements adopted by the Electrical Examining Board; Chapter 404 adopts the 2023 NEC for covered installations. | 481—Ch. 404 | src:usa-ia:iac-481-404 | verified_core |
| Energy | ahj:usa-ia:dial-building-code-commissioner | State Building Code Commissioner / DIAL | Administers statewide energy-conservation requirements under Part 3 of Chapter 301. | Iowa Code 103A.10; 481—301.23 through 301.25 | src:usa-ia:code-103a; src:usa-ia:iac-481-301 | verified_core |
| Fire - construction references | ahj:usa-ia:dial-fire-control | DIAL Fire Control / Building Code Bureau | Fire rules reference IBC/IEBC compliance for building work and include plan review paths. | Iowa Code 10A.511; 481—Ch. 280 | src:usa-ia:code-10a-511; src:usa-ia:iac-481-280 | partially_verified |
| Fire - operational / prevention code | ahj:usa-ia:dial-fire-control | DIAL Fire Control in consultation with State Fire Marshal | Promulgates fire-safety rules; Chapter 280 adopts the 2024 IFC with Iowa amendments. | Iowa Code 10A.511; 481—Ch. 280 | src:usa-ia:code-10a-511; src:usa-ia:iac-481-280 | verified_core |
| Accessibility | ahj:usa-ia:dial-building-code-commissioner | State Building Code Commissioner / DIAL | Chapter 301 applies public-building and multiunit residential accessibility provisions statewide for covered work. | Iowa Code ch. 103A; 481—301.16 through 301.19 | src:usa-ia:code-103a; src:usa-ia:iac-481-301 | verified_core |
| Elevator / Conveyance | ahj:usa-ia:elevator-safety-board | Elevator Safety Board / DIAL Director | Safety board formulates conveyance rules; DIAL director enforces Iowa State Elevator Code requirements. | Iowa Code ch. 89A; 481—Chs. 371 and 372 | src:usa-ia:code-89a; src:usa-ia:iac-481-371; src:usa-ia:iac-481-372 | partially_verified |

### 2.3 Authority Hierarchy Notes

Iowa's primary building-code authority is centralized for state-code adoption, but enforcement and applicability vary by project type and code family. The State Building Code Commissioner adopts the state code with Building Code Advisory Council approval. The code applies directly to selected statewide categories and to local jurisdictions that adopt the state code by ordinance. Local building departments remain important for locally administered plan review and inspection.

Fire authority is separate from the building-code commissioner model. Iowa Code section 10A.511 gives the DIAL director fire-control duties and requires fire-safety rules to be adopted in consultation with the State Fire Marshal. Electrical, plumbing, mechanical, and elevator/conveyance rules are also specialized and should not be assumed to follow the same local-amendment or local-enforcement rules as general building-code provisions.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-ia:001 | ahj:usa-ia:dial-building-code-commissioner | adopts_with_approval_of | ahj:usa-ia:building-code-advisory-council / State Building Code amendments | src:usa-ia:code-103a; src:usa-ia:iac-481-300 | verified_core |
| edge:usa-ia:002 | ahj:usa-ia:dial-building-code-commissioner | administers_plan_review_for | state-owned, certain state-funded, and other state-code projects identified in 481—300.6 | src:usa-ia:code-103a; src:usa-ia:iac-481-300 | verified_core |
| edge:usa-ia:003 | ahj:usa-ia:local-jurisdictions | may_adopt_by_ordinance | State Building Code, effective no more than six months after local adoption | src:usa-ia:code-103a | verified_core |
| edge:usa-ia:004 | ahj:usa-ia:dial-fire-control | consults_with | ahj:usa-ia:state-fire-marshal / fire-safety rules | src:usa-ia:code-10a-511; src:usa-ia:iac-481-280 | verified_core |
| edge:usa-ia:005 | ahj:usa-ia:dial-fire-control | recognizes_compliance_through | qualifying local fire ordinances with plan review and regular inspections | src:usa-ia:iac-481-280 | partially_verified |
| edge:usa-ia:006 | ahj:usa-ia:plumbing-mechanical-systems-board | allows_local_enforcement_by | local jurisdictions administering plumbing/mechanical permits, inspections, testing, and enforcement | src:usa-ia:iac-481-425; src:usa-ia:iac-481-426 | verified_core |
| edge:usa-ia:007 | ahj:usa-ia:dial-electrical-inspection-program | coordinates_with | political subdivisions performing electrical inspections | src:usa-ia:iac-481-404 | verified_core |
| edge:usa-ia:008 | ahj:usa-ia:elevator-safety-board | formulates_rules_for | conveyance safety; DIAL director enforcement | src:usa-ia:code-89a; src:usa-ia:iac-481-372 | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | Iowa State Building Code, Part 1 | International Building Code | 2024 | current_with_state_amendments | null | 2025-09-10 | 2025-09-10 | 2026-09-10 for covered projects commenced after transition window | 12-month dual-use transition for Part 1 projects commenced within 12 months after effective date | src:usa-ia:iac-481-301 |
| Residential | Iowa State Building Code residential provisions | International Residential Code | 2024 | current_with_state_amendments | null | 2025-09-10 | 2025-09-10 | 2026-09-10 for covered projects commenced after transition window | 12-month dual-use transition for Part 1 projects commenced within 12 months after effective date | src:usa-ia:iac-481-301 |
| Existing Building / Rehabilitation | Iowa State Existing Building Code provisions | International Existing Building Code | 2024 | current_with_state_amendments | null | 2025-09-10 | 2025-09-10 | 2026-09-10 for covered projects commenced after transition window | 12-month dual-use transition for Part 1 projects commenced within 12 months after effective date | src:usa-ia:iac-481-301 |
| Mechanical | State Mechanical Code | International Mechanical Code | 2024 | current_with_state_amendments | null | 2025-03-26 | 2025-03-26 | null | No general transition rule verified for Chapter 426 | src:usa-ia:iac-481-426 |
| Plumbing | State Plumbing Code | Uniform Plumbing Code | 2024 | current_with_state_amendments | null | 2025-03-26 | 2025-03-26 | null | No general transition rule verified for Chapter 425 | src:usa-ia:iac-481-425 |
| Fuel Gas | Iowa fuel-gas piping requirements | UPC Ch. 12; NFPA 54; NFPA 58; state LPG/fire cross-references | 2024 UPC / current NFPA references | partially_verified | null | 2025-03-26 | 2025-03-26 | null | No standalone fuel-gas transition rule verified | src:usa-ia:iac-481-301; src:usa-ia:iac-481-425; src:usa-ia:iac-481-426 |
| Electrical | State Electrical Code / Electrical Inspection Program rules | National Electrical Code | 2023 | current_with_state_amendments | null | 2025-07-01 | 2025-07-01 | null | No general transition rule verified for Chapter 404 | src:usa-ia:iac-481-404 |
| Energy | Iowa statewide energy code provisions | International Energy Conservation Code, Residential and Commercial Provisions | 2012 | current_with_state_amendments | null | 2025-09-10 | 2025-09-10 | 2026-09-10 for covered projects commenced after transition window | 12-month dual-use transition for Part 3 projects commenced within 12 months after effective date | src:usa-ia:code-103a; src:usa-ia:iac-481-301 |
| Fire - construction references | Iowa Fire Control Administration building references | IBC / IEBC cross-references | 2024 | current_with_state_amendments | null | 2025-09-10 | 2025-09-10 | null | No general construction-reference transition rule verified in Chapter 280 | src:usa-ia:iac-481-280 |
| Fire - operational / prevention code | Iowa Fire Code / Fire Control Administration rules | International Fire Code | 2024 | current_with_state_amendments | null | 2025-09-10 | 2025-09-10 | null | No general transition rule verified for Chapter 280 | src:usa-ia:code-10a-511; src:usa-ia:iac-481-280 |
| Accessibility | Iowa accessibility provisions for public and multiunit residential buildings | 2010 ADA Standards or IBC accessibility provisions; Chapter 301 state rules | 2010 ADA / 2024 IBC references | current_with_state_amendments | null | 2025-09-10 | 2025-09-10 | 2026-09-10 where Part 1 transition applies | Chapter 301 transition may apply where accessibility compliance is through Part 1 building provisions | src:usa-ia:code-103a; src:usa-ia:iac-481-301 |
| Elevator / Conveyance | Iowa State Elevator Code / conveyance safety rules | ASME A17.1/CSA B44, ASME A17.3, ASME A17.7, ASME A18.1, ANSI A117.1, NFPA 70 references | mixed by installation date; current new-installation reference includes ASME A17.1-2019/CSA B44-19 | current_with_state_amendments | null | 2025-02-12 | 2025-02-12 | null | Installation-date rules govern the applicable conveyance standard | src:usa-ia:code-89a; src:usa-ia:iac-481-371; src:usa-ia:iac-481-372 |

### 3.2 Adoption Records

| Adoption Record ID | Code Family | State Code / Rule | Base Model Code | Edition | Effective Date | Operative Date | Mandatory Date | Scope Notes | Source IDs | Verification Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| adoption:usa-ia:building:2024-ibc | Building | 481—301.3 | International Building Code | 2024 | 2025-09-10 | 2025-09-10 | 2026-09-10 for post-transition covered projects | Applies to buildings and facilities subject to the State Building Code under Iowa Code ch. 103A. | src:usa-ia:iac-481-301 | verified_core |
| adoption:usa-ia:residential:2024-irc | Residential | 481—301.8 | International Residential Code | 2024 | 2025-09-10 | 2025-09-10 | 2026-09-10 for post-transition covered projects | Covers one- and two-family dwellings and townhouses not more than three stories above grade plane. | src:usa-ia:iac-481-301 | verified_core |
| adoption:usa-ia:existing:2024-iebc | Existing Building / Rehabilitation | 481—301.7 | International Existing Building Code | 2024 | 2025-09-10 | 2025-09-10 | 2026-09-10 for post-transition covered projects | Covers repair, alteration, change of occupancy, addition, and relocation of existing buildings. | src:usa-ia:iac-481-301 | verified_core |
| adoption:usa-ia:mechanical:2024-imc | Mechanical | 481—426.3 | International Mechanical Code | 2024 | 2025-03-26 | 2025-03-26 | null | Applies to mechanical systems in Iowa, with local administration provisions. | src:usa-ia:iac-481-426 | verified_core |
| adoption:usa-ia:plumbing:2024-upc | Plumbing | 481—425.3 | Uniform Plumbing Code | 2024 | 2025-03-26 | 2025-03-26 | null | Applies to plumbing in Iowa, with local administration provisions. | src:usa-ia:iac-481-425 | verified_core |
| adoption:usa-ia:fuel-gas:2024-mixed | Fuel Gas | 481—301.9; 481—425.3; 481—426.3 | UPC Chapter 12 plus NFPA and state fire/LPG cross-references | mixed | 2025-03-26 | 2025-03-26 | null | Fuel gas is not represented as a single standalone state model-code adoption in the parsed rules. | src:usa-ia:iac-481-301; src:usa-ia:iac-481-425; src:usa-ia:iac-481-426 | partially_verified |
| adoption:usa-ia:electrical:2023-nec | Electrical | 481—404.2 | National Electrical Code | 2023 | 2025-07-01 | 2025-07-01 | null | Applies to electrical installations subject to Iowa Code ch. 103. | src:usa-ia:iac-481-404 | verified_core |
| adoption:usa-ia:energy:2012-iecc-res | Energy | 481—301.24 | International Energy Conservation Code, Residential Provisions | 2012 | 2025-09-10 | 2025-09-10 | 2026-09-10 for post-transition covered projects | Residential energy code for applicable residential construction limited to three or fewer stories statewide. | src:usa-ia:code-103a; src:usa-ia:iac-481-301 | verified_core |
| adoption:usa-ia:energy:2012-iecc-com | Energy | 481—301.25 | International Energy Conservation Code, Commercial Provisions | 2012 | 2025-09-10 | 2025-09-10 | 2026-09-10 for post-transition covered projects | Nonresidential energy code for commercial construction and residential construction of four or more stories statewide. | src:usa-ia:code-103a; src:usa-ia:iac-481-301 | verified_core |
| adoption:usa-ia:fire:2024-ifc | Fire - operational / prevention code | 481—280.4 | International Fire Code | 2024 | 2025-09-10 | 2025-09-10 | null | General fire-safety requirements with Iowa amendments and local-compliance recognition paths. | src:usa-ia:code-10a-511; src:usa-ia:iac-481-280 | verified_core |
| adoption:usa-ia:accessibility:2010-ada-ibc | Accessibility | 481—301.16 through 301.19 | 2010 ADA Standards or IBC accessibility provisions | 2010 ADA / 2024 IBC references | 2025-09-10 | 2025-09-10 | null | Public-building and multiunit residential accessibility requirements are applied statewide for covered new work and renovations. | src:usa-ia:code-103a; src:usa-ia:iac-481-301 | verified_core |
| adoption:usa-ia:elevator:asme-a17 | Elevator / Conveyance | 481—372.1 | ASME A17.1/CSA B44 and related standards | mixed by installation date | 2025-02-12 | 2025-02-12 | null | Chapter 372 uses installation-date applicability; current new-installation period begins on or after 2021-06-01. | src:usa-ia:code-89a; src:usa-ia:iac-481-371; src:usa-ia:iac-481-372 | partially_verified |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

The highest-value date rule found in this pass is the 12-month Chapter 301 transition for projects subject to Part 1 or Part 3. Chapter 301 states that a project commenced within 12 months after the effective date of the revised standards may comply with either the revised standards or the prior standards. Chapter 301 also defines "commenced" by reference to preliminary approval from the commissioner or local building department under 481—300.6.

Local adoption of the State Building Code by a governmental subdivision has a separate statutory timing rule: the local ordinance becomes effective no more than six months after the date of adoption. Electrical, plumbing, mechanical, fire, and elevator/conveyance chapters have verified effective dates, but no broad transition rule equivalent to Chapter 301 was extracted for those chapters.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-ia:chapter-301-transition | Building, residential, existing building, accessibility-by-building-code path, and energy provisions in Chapter 301 Parts 1 and 3 | transition_period | 12 months from 2025-09-10 through 2026-09-10 | Project subject to Part 1 or Part 3 is commenced within 12 months after the effective date of the revised standards. | yes | src:usa-ia:iac-481-301 | verified_core |
| date-rule:usa-ia:chapter-301-commenced | Chapter 301 Part 1 and Part 3 projects | commencement_definition | tied to preliminary approval under 481—300.6 | Project has obtained preliminary approval from the commissioner or local building department, as applicable. | not applicable | src:usa-ia:iac-481-301; src:usa-ia:iac-481-300 | verified_core |
| date-rule:usa-ia:local-state-code-ordinance | Local adoption of the State Building Code | local_adoption_effective_deadline | no more than six months after local adoption | Governmental subdivision adopts the State Building Code by ordinance. | unresolved | src:usa-ia:code-103a | verified_core |
| date-rule:usa-ia:electrical-effective-2025 | Electrical | effective_date | 2025-07-01 | Chapter 404 NEC 2023 adoption becomes effective. | unresolved | src:usa-ia:iac-481-404 | partially_verified |
| date-rule:usa-ia:plumbing-mechanical-effective-2025 | Plumbing and mechanical | effective_date | 2025-03-26 | Chapters 425 and 426 2024 model-code adoptions become effective. | unresolved | src:usa-ia:iac-481-425; src:usa-ia:iac-481-426 | partially_verified |
| date-rule:usa-ia:fire-effective-2025 | Fire | effective_date | 2025-09-10 | Chapter 280 IFC 2024 adoption becomes effective. | unresolved | src:usa-ia:iac-481-280 | partially_verified |
| date-rule:usa-ia:elevator-installation-date | Elevator / conveyance | installation_date_applicability | installation-date dependent | Applicable ASME/ANSI/NFPA standard varies by installation date, including the on-or-after 2021-06-01 period. | no general prior-code transition extracted | src:usa-ia:iac-481-372 | partially_verified |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building / residential / energy | none verified | null | null | null | null | null | active_monitoring | src:usa-ia:iac-481-301 | Recent Chapter 301 amendments were effective 2025-09-10; no later official future statewide change was verified. |
| Fire | none verified | null | null | null | null | null | active_monitoring | src:usa-ia:iac-481-280 | Recent Chapter 280 amendments were effective 2025-09-10; special occupancy fire rules were not exhaustively parsed. |
| Electrical | none verified | null | null | null | null | null | active_monitoring | src:usa-ia:iac-481-404 | Current verified NEC adoption is the 2023 NEC, effective 2025-07-01. |
| Plumbing / mechanical | none verified | null | null | null | null | null | active_monitoring | src:usa-ia:iac-481-425; src:usa-ia:iac-481-426 | Current verified plumbing/mechanical adoptions were effective 2025-03-26. |
| Elevator / conveyance | none verified | null | null | null | null | null | active_monitoring | src:usa-ia:iac-481-372 | Conveyance rules were only parsed for core model-standard applicability. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-ia:state-code-covered-projects | Building | State-owned buildings, local jurisdictions adopting the state code, state-appropriated new construction not wholly state-owned, certain larger cities without a local code, and factory-built structures | Statutory applicability under Iowa Code 103A.10 and related provisions | The State Building Code is not universal by default; it attaches to specified categories and local-adoption contexts. | src:usa-ia:code-103a | verified_core |
| applicability-rule:usa-ia:statewide-energy | Energy | New construction and covered lighting/energy-efficiency work statewide | Human-occupancy buildings heated or cooled; lighting standards for lit human-occupancy buildings | State energy conservation requirements apply statewide and supersede certain local minimum energy requirements for new construction. | src:usa-ia:code-103a; src:usa-ia:iac-481-301 | verified_core |
| applicability-rule:usa-ia:statewide-accessibility | Accessibility | Public buildings and multiunit residential buildings | Covered new construction and applicable renovation/rehabilitation work | Chapter 301 accessibility rules apply statewide for specified public-building and residential scopes. | src:usa-ia:iac-481-301 | verified_core |
| applicability-rule:usa-ia:local-fire-ordinance-recognition | Fire | Buildings in local jurisdictions with qualifying local fire ordinances | Local ordinance incorporates any edition of IFC, NFPA 1, or the 1997 Uniform Fire Code and includes plan review and regular inspections | The state rule provides a local-compliance recognition path, subject to exceptions and Chapter 282 requirements for flammable and combustible liquids. | src:usa-ia:iac-481-280 | partially_verified |
| applicability-rule:usa-ia:elevator-installation-period | Elevator / conveyance | Conveyances installed in specified date ranges | Installation date | Chapter 372 assigns standards by installation date, including current new-installation standards for installations on or after 2021-06-01. | src:usa-ia:iac-481-372 | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** integrated administrative-code amendments by chapter.

**Where amendments are published:** Iowa Administrative Code, chiefly 481—Chapters 280, 300, 301, 404, 425, 426, 371, and 372 for the code families parsed in this report.

**Amendment parsing status:** targeted_extract_only

The report identifies model-code editions and selected high-impact Iowa modifications, but it does not claim a complete clause-by-clause amendment inventory.

### 5.2 State Amendment Sources

| Amendment Source ID | Code Family | Publication Path | Coverage | Source IDs | Parsing Status |
| --- | --- | --- | --- | --- | --- |
| amendset:usa-ia:building-301 | Building / residential / existing / energy / accessibility | 481—Chapter 301 | Iowa amendments to IBC, IRC, IEBC, IECC, accessibility, and transition rules | src:usa-ia:iac-481-301 | targeted_extract_only |
| amendset:usa-ia:fire-280 | Fire | 481—Chapter 280 | Iowa amendments to IFC and local fire-ordinance recognition | src:usa-ia:iac-481-280 | targeted_extract_only |
| amendset:usa-ia:electrical-404 | Electrical | 481—Chapter 404 | Iowa amendments to NEC and inspection program rules | src:usa-ia:iac-481-404 | targeted_extract_only |
| amendset:usa-ia:plumbing-425 | Plumbing | 481—Chapter 425 | Iowa amendments to UPC and local-administration rules | src:usa-ia:iac-481-425 | targeted_extract_only |
| amendset:usa-ia:mechanical-426 | Mechanical | 481—Chapter 426 | Iowa amendments to IMC and local-administration rules | src:usa-ia:iac-481-426 | targeted_extract_only |
| amendset:usa-ia:elevator-372 | Elevator / conveyance | 481—Chapter 372 | Iowa conveyance standards and installation-date applicability | src:usa-ia:iac-481-372 | targeted_extract_only |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Rule / Section | Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| amend:usa-ia:irc-sprinklers-deleted | Residential | 481—301.8 | Iowa deletes selected IRC residential sprinkler provisions in the adopted 2024 IRC amendments. | src:usa-ia:iac-481-301 | targeted_verified |
| amend:usa-ia:chapter-301-transition | Building / energy | 481—301.15 | Iowa provides a 12-month transition for covered Part 1 and Part 3 projects after revised standards become effective. | src:usa-ia:iac-481-301 | verified_core |
| amend:usa-ia:energy-2012-iecc | Energy | 481—301.24 and 481—301.25 | Iowa's statewide energy rules adopt 2012 IECC Residential and Commercial Provisions, despite 2024 editions in other code families. | src:usa-ia:iac-481-301 | verified_core |
| amend:usa-ia:fire-local-recognition | Fire | 481—280.11 | Iowa fire rules recognize compliance through qualifying local fire ordinances with plan review and regular inspections, subject to exceptions. | src:usa-ia:iac-481-280 | partially_verified |
| amend:usa-ia:plumbing-local-code-limits | Plumbing | 481—425.3 | Local jurisdictions may not adopt other plumbing codes, but may adopt stricter plumbing amendments and provide copies to the board. | src:usa-ia:iac-481-425 | verified_core |
| amend:usa-ia:mechanical-local-code-limits | Mechanical | 481—426.5 | Local jurisdictions may not adopt other mechanical codes, but may adopt stricter mechanical amendments and provide copies to the board. | src:usa-ia:iac-481-426 | verified_core |
| amend:usa-ia:elevator-local-conflict | Elevator / conveyance | Iowa Code ch. 89A | Iowa State Elevator Code provisions supersede conflicting building-code provisions and local conflicting ordinances. | src:usa-ia:code-89a | partially_verified |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-ia"
  model: "hybrid_state_reserved_and_local"
  enforcing_entities:
    - "DIAL Building Code Bureau / State Building Code Commissioner for state-code plan-review and inspection scopes identified in Iowa Code ch. 103A and 481—Ch. 300"
    - "Local building departments for locally adopted code administration and for projects where local plan review is available"
    - "DIAL fire-control authority and authorized local fire departments for fire inspections under 481—Ch. 280"
    - "Political subdivisions performing electrical inspections, with DIAL electrical inspection program for other covered work and state-owned property"
    - "Local jurisdictions administering plumbing and mechanical permits, inspections, testing, and enforcement"
    - "DIAL director and Elevator Safety Board framework for conveyance safety"
  required_officials:
    - "State Building Code Commissioner"
    - "Building Code Advisory Council"
    - "DIAL director / fire-control administration"
    - "State Fire Marshal consultation role"
    - "Electrical inspectors or political subdivision inspection programs where applicable"
    - "Local building officials where local code administration exists"
  state_reserved_activities:
    - "State-owned building plan review and comment under 481—300.6"
    - "Certain state-appropriated projects where no qualifying local code administration exists"
    - "State electrical inspections for state-owned property under Chapter 404"
    - "Conveyance safety enforcement under Iowa Code ch. 89A"
  source_ids:
    - "src:usa-ia:code-103a"
    - "src:usa-ia:iac-481-300"
    - "src:usa-ia:iac-481-280"
    - "src:usa-ia:iac-481-404"
    - "src:usa-ia:iac-481-425"
    - "src:usa-ia:iac-481-426"
    - "src:usa-ia:code-89a"
  verification_status: "partially_verified"
  confidence: 0.72
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-ia"
  model: "mixed_by_code_family"
  applies_to_code_families:
    - "Building: general local-amendment scope remains partially unresolved; local adoption of the State Building Code by ordinance is verified."
    - "Plumbing: local jurisdictions may not adopt other plumbing codes, but may adopt stricter amendments and provide copies to the board."
    - "Mechanical: local jurisdictions may not adopt other mechanical codes, but may adopt stricter amendments and provide copies to the board."
    - "Fire: local fire ordinances can be recognized if they incorporate an approved model fire code and include plan review and regular inspections."
    - "Energy: statewide energy requirements are verified; local stringency and preemption details require additional review beyond the supersession language extracted from Iowa Code ch. 103A."
    - "Elevator / conveyance: conflicting local provisions are limited by Iowa Code ch. 89A; local inspection acceptance exists but full local-role detail remains partially parsed."
  approval_required: "varies_by_code_family"
  approving_authority_id: "varies_by_code_family"
  filing_required: "plumbing_and_mechanical_copy_to_board_verified; other filing requirements unresolved"
  registry_exists: null
  registry_source_ids: []
  legal_basis_source_ids:
    - "src:usa-ia:code-103a"
    - "src:usa-ia:iac-481-280"
    - "src:usa-ia:iac-481-425"
    - "src:usa-ia:iac-481-426"
    - "src:usa-ia:code-89a"
  verification_status: "partially_verified"
  confidence: 0.58
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement and local amendment authority are not the same in Iowa. A local jurisdiction may enforce or administer a code without having free authority to adopt a different model code. The clearest example is plumbing and mechanical: the state rules allow local administration of permits and inspections, but restrict local jurisdictions from adopting other plumbing or mechanical codes while allowing stricter local amendments with copies provided to the board.

For fire, the state rule recognizes a local compliance path when a local ordinance incorporates specified model fire codes and includes plan review and regular inspections. For general building-code local amendments, the current pass verifies local adoption and state-code applicability but does not fully classify local amendment limits.

### 6.4 Known Local Amendment Registries

| Registry ID | Code Family | Registry Name / Holder | Public URL | Filing Rule | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| registry:usa-ia:plumbing-mechanical-copies | Plumbing / mechanical | Plumbing and Mechanical Systems Board copy requirement | not located | Local jurisdictions adopting stricter plumbing or mechanical amendments provide copies to the board. | src:usa-ia:iac-481-425; src:usa-ia:iac-481-426 | partially_verified_no_public_registry |
| registry:usa-ia:building-local-amendments | Building | none verified | null | General filing or registry rule not extracted. | src:usa-ia:code-103a | unresolved |
| registry:usa-ia:fire-local-ordinances | Fire | none verified | null | Local recognition criteria extracted, but no statewide registry was located. | src:usa-ia:iac-481-280 | unresolved |

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendments were parsed in this pass. A production AHJ resolver should not infer local amendments from this statewide report alone.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: partial_model_only

Jurisdiction stack:

```text
Address
  -> State of Iowa
  -> County
  -> Municipality / unincorporated county
  -> Local building-code adoption status
  -> Local building department and plan-review availability
  -> Fire jurisdiction and local fire-ordinance recognition status
  -> Electrical inspection jurisdiction: state program or political subdivision
  -> Plumbing/mechanical local administration status
  -> Elevator/conveyance state jurisdiction, if applicable
  -> Applicable statewide adoption records
  -> Applicable local amendment records, if verified
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

No AHJ contact data was populated for this pass. The report should be paired with a separate local-adoption and local-department dataset before address-level production use.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Publisher / Owner | Type | URL / Location | Date / Version | Used For | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-ia:code-103a | Iowa Code 2026, Chapter 103A, State Building Code | Iowa Legislature | statute PDF | https://www.legis.iowa.gov/docs/ico/chapter/103A.pdf | Iowa Code 2026 PDF generated 2025-12-09 | State building-code authority, applicability, energy statewide provisions, local adoption timing, local enforcement caveats | official_primary |
| src:usa-ia:iac-481-300 | 481—Chapter 300, State Building Code—Administration | Iowa Legislature / DIAL | administrative rule PDF | https://www.legis.iowa.gov/docs/iac/chapter/02-18-2026.481.300.pdf | IAC Supplement 2025-11-26; ARC 9473C effective 2025-09-10 | Administration, commissioner role, advisory council role, plan review/inspection submission paths | official_primary |
| src:usa-ia:iac-481-301 | 481—Chapter 301, State Building Code | Iowa Legislature / DIAL | administrative rule PDF | https://www.legis.iowa.gov/docs/iac/chapter/01-07-2026.481.301.pdf | IAC Supplement 2025-11-26; ARC 9474C effective 2025-09-10 | IBC/IRC/IEBC adoptions, energy code, accessibility, transition rule, mechanical/electrical/plumbing cross-references | official_primary |
| src:usa-ia:code-10a-511 | Iowa Code 2026, Section 10A.511, Fire Control Duties | Iowa Legislature | statute PDF | https://www.legis.iowa.gov/docs/code/10A.511.pdf | Iowa Code 2026 PDF generated 2025-12-09 | Fire-control statutory authority and State Fire Marshal consultation | official_primary |
| src:usa-ia:iac-481-280 | 481—Chapter 280, Fire Control Administration | Iowa Legislature / DIAL | administrative rule PDF | https://www.legis.iowa.gov/docs/iac/chapter/01-07-2026.481.280.pdf | IAC Supplement 2025-11-26; ARC 9477C effective 2025-09-10 | IFC 2024 adoption, fire inspections, plan submissions, local fire-ordinance recognition | official_primary |
| src:usa-ia:iac-481-404 | 481—Chapter 404, Electrical Inspection Program | Iowa Legislature / DIAL | administrative rule PDF | https://www.legis.iowa.gov/docs/iac/chapter/02-18-2026.481.404.pdf | ARC 9033C effective 2025-07-01 | State electrical inspection program and 2023 NEC adoption | official_primary |
| src:usa-ia:iac-481-425 | 481—Chapter 425, State Plumbing Code | Iowa Legislature / Plumbing and Mechanical Systems Board | administrative rule PDF | https://www.legis.iowa.gov/docs/iac/chapter/02-18-2026.481.425.pdf | ARC 8958C effective 2025-03-26 | 2024 UPC adoption, local administration, local stricter amendments | official_primary |
| src:usa-ia:iac-481-426 | 481—Chapter 426, State Mechanical Code | Iowa Legislature / Plumbing and Mechanical Systems Board | administrative rule PDF | https://www.legis.iowa.gov/docs/iac/chapter/02-18-2026.481.426.pdf | ARC 8958C effective 2025-03-26 | 2024 IMC adoption, fuel-gas references, local administration, local stricter amendments | official_primary |
| src:usa-ia:code-89a | Iowa Code 2026, Chapter 89A, Elevators | Iowa Legislature | statute PDF | https://www.legis.iowa.gov/docs/ico/chapter/89A.pdf | Iowa Code 2026 PDF generated 2025-12-09 | Iowa State Elevator Code authority, safety board, director enforcement, local conflict rule | official_primary |
| src:usa-ia:iac-481-371 | 481—Chapter 371, Administration of Conveyance Safety Program | Iowa Legislature / DIAL | administrative rule PDF | https://www.legis.iowa.gov/docs/iac/chapter/481.371.pdf | IAC source accessed 2026-06-25 | Conveyance program definitions and permit context | official_primary |
| src:usa-ia:iac-481-372 | 481—Chapter 372, Conveyances Installed On or After January 1, 1975 | Iowa Legislature / DIAL | administrative rule PDF | https://www.legis.iowa.gov/docs/iac/chapter/02-18-2026.481.372.pdf | ARC 8774C effective 2025-02-12; IAC Supplement 2025-07-09 | Conveyance safety standards by installation date | official_primary |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-ia:code-103a | official_pdf_extraction | Official Iowa Legislature PDF was parsed from text and spot-checked against rendered pages for high-value authority and applicability fields. | usable_with_spot_checks |
| src:usa-ia:iac-481-300 | official_pdf_extraction | Official Iowa Legislature administrative-rule PDF was parsed from text and spot-checked against rendered pages for commissioner, council, and plan-review fields. | usable_with_spot_checks |
| src:usa-ia:iac-481-301 | official_pdf_extraction | Official Iowa Legislature administrative-rule PDF is long and amendment-heavy; this pass extracted adoption, energy, accessibility, and transition fields but not every amendment. | usable_for_core_fields_only |
| src:usa-ia:iac-481-280 | official_pdf_extraction | Official Iowa Legislature administrative-rule PDF includes special occupancy provisions beyond the general IFC/local-ordinance fields extracted here. | needs_special_occupancy_followup |
| src:usa-ia:iac-481-404 | official_pdf_extraction | Electrical rule was parsed for program authority, NEC edition, and inspection split; full permitting workflow was not exhaustively modeled. | usable_for_core_fields_only |
| src:usa-ia:iac-481-425 | official_pdf_extraction | Plumbing rule was parsed for UPC edition and local-administration/local-amendment posture; full amendments were not exhaustively modeled. | usable_for_core_fields_only |
| src:usa-ia:iac-481-426 | official_pdf_extraction | Mechanical rule was parsed for IMC edition, fuel-gas references, and local-administration/local-amendment posture; full amendments were not exhaustively modeled. | usable_for_core_fields_only |
| src:usa-ia:code-89a | official_pdf_extraction | Elevator statute was parsed for authority and preemption/local-conflict fields; full permit, inspection, and variance process remains partially parsed. | usable_for_core_fields_only |
| src:usa-ia:iac-481-371 | limited_extract | Chapter 371 was used only for conveyance program context and definitions. | needs_followup_if_conveyance_detail_required |
| src:usa-ia:iac-481-372 | official_pdf_extraction | Chapter 372 was parsed for model-standard applicability by installation date; full conveyance amendment inventory was not prepared. | usable_for_core_fields_only |

### 8.3 Supplemental Sources

None used in this pass. All substantive report findings are based on official Iowa Code or Iowa Administrative Code sources.

### 8.4 Source Extraction Metadata

| Extraction ID | Source IDs | Method | Extracted Fields | Extracted On | Notes |
| --- | --- | --- | --- | --- | --- |
| extract:usa-ia:authority-core | src:usa-ia:code-103a; src:usa-ia:iac-481-300 | official PDF text plus rendered-page spot check | primary authority, advisory council, plan review, applicability | 2026-06-25 | Core authority fields populated. |
| extract:usa-ia:adoptions-core | src:usa-ia:iac-481-301; src:usa-ia:iac-481-280; src:usa-ia:iac-481-404; src:usa-ia:iac-481-425; src:usa-ia:iac-481-426 | official PDF text plus targeted rendered-page spot checks | model-code editions, effective dates, transition fields | 2026-06-25 | Adoption matrix and normalized records populated. |
| extract:usa-ia:elevator-core | src:usa-ia:code-89a; src:usa-ia:iac-481-371; src:usa-ia:iac-481-372 | official PDF text plus targeted rendered-page spot check | elevator authority and installation-date standards | 2026-06-25 | Conveyance detail remains partial. |
| extract:usa-ia:local-rules | src:usa-ia:code-103a; src:usa-ia:iac-481-280; src:usa-ia:iac-481-425; src:usa-ia:iac-481-426 | official PDF text | local enforcement and local amendment posture | 2026-06-25 | Building local-amendment scope requires follow-up. |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 1.00 | none | Status reflects source-backed core authority/adoption fields with unresolved local details. |
| report | risk.overall_confidence | 0.68 | verified | 1.00 | none | Confidence remains below verified due local AHJ, local-amendment, factory-built, and special-occupancy gaps. |
| ahj:usa-ia:dial-building-code-commissioner | role | adopts State Building Code with council approval | verified | 0.88 | src:usa-ia:code-103a; src:usa-ia:iac-481-300 | Highest-value authority field. |
| adoption:usa-ia:building:2024-ibc | base_model_code | 2024 IBC | verified | 0.88 | src:usa-ia:iac-481-301 | Adoption rule extracted from Chapter 301. |
| adoption:usa-ia:residential:2024-irc | base_model_code | 2024 IRC | verified | 0.86 | src:usa-ia:iac-481-301 | Adoption rule extracted from Chapter 301. |
| adoption:usa-ia:existing:2024-iebc | base_model_code | 2024 IEBC | verified | 0.86 | src:usa-ia:iac-481-301 | Adoption rule extracted from Chapter 301. |
| adoption:usa-ia:energy:2012-iecc-res | base_model_code | 2012 IECC Residential Provisions | verified | 0.86 | src:usa-ia:iac-481-301 | Statewide energy code field. |
| adoption:usa-ia:fire:2024-ifc | base_model_code | 2024 IFC | verified | 0.84 | src:usa-ia:iac-481-280 | Fire code adoption extracted from Chapter 280. |
| adoption:usa-ia:electrical:2023-nec | base_model_code | 2023 NEC | verified | 0.82 | src:usa-ia:iac-481-404 | Electrical adoption extracted from Chapter 404. |
| adoption:usa-ia:plumbing:2024-upc | base_model_code | 2024 UPC | verified | 0.84 | src:usa-ia:iac-481-425 | Plumbing adoption extracted from Chapter 425. |
| adoption:usa-ia:mechanical:2024-imc | base_model_code | 2024 IMC | verified | 0.84 | src:usa-ia:iac-481-426 | Mechanical adoption extracted from Chapter 426. |
| date-rule:usa-ia:chapter-301-transition | transition_period | 12 months after 2025-09-10 effective date | verified | 0.84 | src:usa-ia:iac-481-301 | Trigger condition and prior-code allowance captured. |
| local-amendment-rule:usa-ia | model | mixed_by_code_family | partially_verified | 0.58 | src:usa-ia:iac-481-280; src:usa-ia:iac-481-425; src:usa-ia:iac-481-426 | Plumbing/mechanical and fire posture captured; general building amendment scope unresolved. |
| local-enforcement:usa-ia | model | hybrid_state_reserved_and_local | partially_verified | 0.72 | src:usa-ia:code-103a; src:usa-ia:iac-481-300; src:usa-ia:iac-481-404; src:usa-ia:iac-481-425; src:usa-ia:iac-481-426 | High-level model supported, but local AHJ dataset missing. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | All `src:usa-ia:*` identifiers cited in the body appear in Section 8. |
| All authority IDs resolve | pass | Authority IDs used in Section 2 are described in the authority tables or graph notes. |
| All current code families have adoption records | pass | Every code family in the adoption matrix has a corresponding normalized adoption row or partial record. |
| Building and operational fire code are separated | pass | Building/fire construction references and operational/prevention fire code are separate rows. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Adoption dates remain null where not directly verified; effective, operative, and mandatory dates are separated. |
| Effective dates are valid ISO dates | pass | Entered effective dates use ISO format. |
| No impossible date sequences | pass | No adoption/effective/mandatory sequence conflict was introduced. |
| Transition rules have explicit trigger conditions | pass | Chapter 301 transition trigger and commencement definition are captured. |
| Permit-date logic is captured where applicable | partial | Chapter 301 preliminary-approval logic is captured; electrical/plumbing/mechanical permit-date transition logic remains unresolved. |
| Local enforcement model classified | partial | State/local split is classified, but local AHJ contact and coverage data remain absent. |
| Local amendment rule classified | partial | Plumbing/mechanical and fire rules are classified; general building local-amendment limits need follow-up. |
| AHJ confirmation metadata present | fail | No address-level AHJ contact or boundary dataset was populated. |
| Official-source caveats captured | pass | Caveats are listed by source in Section 8.2. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-ia:001 | high | local building amendments | General local amendment authority for building-code provisions is only partially resolved. | Extract Iowa local-government enabling statutes, 103A local-code provisions, and any DIAL guidance on local building amendments. | null | null | open |
| issue:usa-ia:002 | high | local AHJ dataset | Address-level production use requires a local adoption, building department, electrical inspection, plumbing/mechanical administration, and fire jurisdiction dataset. | Build or ingest local AHJ coverage and contact records. | null | null | open |
| issue:usa-ia:003 | medium | factory-built structures | Chapter 300 identifies additional state building code chapters for factory-built structures, but Chapters 321, 322, and 323 were not parsed. | Extract 481—Chapters 321, 322, and 323. | null | null | open |
| issue:usa-ia:004 | medium | fire special occupancies | Chapter 280 includes health care, ambulatory surgery, hospital, child care, correctional, and other special occupancy provisions that were not exhaustively modeled. | Extract special occupancy fire rules and confirm relationship to IFC/Life Safety Code. | null | null | open |
| issue:usa-ia:005 | medium | fuel gas / LPG details | Fuel-gas and LPG cross-references were identified but not fully parsed, including 661—Chapter 226 and 481—286.1. | Extract state LP gas/fire rules and reconcile fuel-gas AHJ responsibility. | null | null | open |
| issue:usa-ia:006 | low | amendment inventory | This report does not contain a complete clause-by-clause Iowa amendment inventory. | Parse each IAC chapter amendment-by-amendment for downstream rules engines. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-ia:code-103a | src:usa-ia:code-103a | html_pdf_diff | monthly | Statutory amendments to Iowa Code ch. 103A | 2026-06-25 | active |
| watch:usa-ia:iac-481-300 | src:usa-ia:iac-481-300 | html_pdf_diff | monthly | Administrative changes to State Building Code administration | 2026-06-25 | active |
| watch:usa-ia:iac-481-301 | src:usa-ia:iac-481-301 | html_pdf_diff | monthly | New building, residential, existing building, energy, accessibility, or transition amendments | 2026-06-25 | active |
| watch:usa-ia:iac-481-280 | src:usa-ia:iac-481-280 | html_pdf_diff | monthly | Fire code adoption or special occupancy fire-rule changes | 2026-06-25 | active |
| watch:usa-ia:iac-481-404 | src:usa-ia:iac-481-404 | html_pdf_diff | monthly | Electrical code edition or inspection-program changes | 2026-06-25 | active |
| watch:usa-ia:iac-481-425 | src:usa-ia:iac-481-425 | html_pdf_diff | monthly | Plumbing code edition, local amendment, or enforcement changes | 2026-06-25 | active |
| watch:usa-ia:iac-481-426 | src:usa-ia:iac-481-426 | html_pdf_diff | monthly | Mechanical code edition, fuel-gas, local amendment, or enforcement changes | 2026-06-25 | active |
| watch:usa-ia:elevator | src:usa-ia:code-89a; src:usa-ia:iac-481-372 | html_pdf_diff | quarterly | Elevator statute or conveyance-standard changes | 2026-06-25 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-24 | Generated baseline draft report | report:usa-ia | none | Codex | Initial stub with unresolved values. |
| 2026-06-25 | Populated source-backed core authority, adoption, transition, local enforcement, local amendment, source registry, QA, and monitoring sections | report:usa-ia; ahj:usa-ia:dial-building-code-commissioner; adoption:usa-ia:*; date-rule:usa-ia:*; local-enforcement:usa-ia; local-amendment-rule:usa-ia | src:usa-ia:code-103a; src:usa-ia:iac-481-300; src:usa-ia:iac-481-301; src:usa-ia:code-10a-511; src:usa-ia:iac-481-280; src:usa-ia:iac-481-404; src:usa-ia:iac-481-425; src:usa-ia:iac-481-426; src:usa-ia:code-89a; src:usa-ia:iac-481-371; src:usa-ia:iac-481-372 | ChatGPT | Status upgraded to partially_verified for source-backed core fields; unresolved local/AHJ and special-scope issues remain explicit. |
