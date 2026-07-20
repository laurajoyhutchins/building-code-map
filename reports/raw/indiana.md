---
state:
  state_id: "US-IN"
  name: "Indiana"
  abbreviation: "IN"
report:
  report_id: "state-report:usa-in"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-25"
  last_verified: "2026-06-25"
  reviewed_by: null
risk:
  overall_confidence: 0.72 # 0.00 - 1.00
  risk_flags:
    - "official_rule_text_partly_js_gated"
    - "state_amendments_not_section_parsed"
    - "local_AHJ_contacts_not_collected"
    - "pending_2026_code_rulemakings"
  open_questions_count: 6

---

# State Building Code Authority Report: Indiana

## 1. Executive Summary

- **Authority model:** Indiana uses a statewide code model centered on the Fire Prevention and Building Safety Commission (`ahj:usa-in:fpbsc`). The Commission is the primary statewide body for building and fire-safety laws, variance review, rule-modification petitions, and order review. The Indiana Department of Homeland Security (`ahj:usa-in:idhs-dfbs`) supports administration through the Division of Fire and Building Safety, including building plan review, code enforcement support, regulated devices, and rule implementation.

- **Statewide code status:** Indiana has current statewide adoptions in Title 675 of the Indiana Administrative Code, including the 2014 Indiana Building Code, 2020 Indiana Residential Code, 2012 Indiana Plumbing Code, Indiana Electrical Code 2009 Edition, 2014 Indiana Mechanical Code, 2010 Indiana Energy Conservation Code, 2014 Indiana Fire Code, 2014 Indiana Fuel Gas Code, Indiana regulated lifting device rules, and boiler/pressure-vessel rules. The IDHS code-navigation page is an informal index, so production use should verify section text and history lines against the Indiana Register / Indiana Administrative Code.

- **Local enforcement model:** The state reserves key Class 1 structure design-release review to the Division, with the Building Plan Review Section reviewing construction plans for compliance with Commission rules. Cities, towns, and counties may participate in Class 1 local plan review only through certification under 675 IAC 12-7 and still route stamped plans into the Division design-release process. Local building, fire, and enforcement departments also work with the IDHS Code Enforcement Section on training, interpretation, and inspection support.

- **Local amendment posture:** IDHS states that, after HEA 1575, substantive local building and fire-safety ordinance provisions outside mere incorporation of Commission rules are generally unenforceable and may not be adopted. Local administrative provisions, such as departments, permits, plan review, inspection schedules, fines, and order-review processes, remain distinct from substantive standards. Modifications to statewide standards are now routed through Commission rulemaking or local petitions to the Commission, not through the former ordinance-review program. Because IDHS expressly labels this as its position rather than binding legal advice, the local-amendment conclusion is marked `partially_verified` rather than `verified`.

- **Known transition periods or pending changes:** Indiana's general administrative rules include a design-release filing-date rule for projects requiring design release and a work-date rule for construction exempt from design release. Pending or active update work exists for the building code, fire prevention code, electrical code, and plumbing code. A 2025 readoption docket readopted several current rules without changes and proposed repeal of 675 IAC 24-1.

- **Production readiness:** narrow_validation_ready

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: The Fire Prevention and Building Safety Commission is the primary statewide
    body for building and fire-safety laws.
  confidence: 0.9
  source_ids:
  - src:usa-in:idhs-fpbsc-overview
  - src:usa-in:iac-675-12-gar-pdf
- topic: Primary building code edition
  finding: 675 IAC 13-2.6 lists the 2014 Indiana Building Code, based on the 2012
    IBC First Printing, as currently in effect, effective 2014-12-01.
  confidence: 0.86
  source_ids:
  - src:usa-in:idhs-fpbsc-rules-overview
  - src:usa-in:iar-title-675-current
- topic: Residential code edition
  finding: 675 IAC 14-4.4 lists the 2020 Indiana Residential Code, based on the 2018
    IRC First Printing, as currently in effect, effective 2019-12-26.
  confidence: 0.86
  source_ids:
  - src:usa-in:idhs-fpbsc-rules-overview
  - src:usa-in:iar-title-675-current
- topic: Electrical code edition
  finding: 675 IAC 17-1.8 lists the Indiana Electrical Code, 2009 Edition, based on
    NFPA 70 NEC 2008 First Printing, as currently in effect, effective 2009-08-26;
    a 2026 rulemaking is pending to replace it.
  confidence: 0.8
  source_ids:
  - src:usa-in:idhs-fpbsc-rules-overview
  - src:usa-in:idhs-electrical-code-update
  - src:usa-in:iar-electrical-lsa-26-62
- topic: Fire code edition
  finding: 675 IAC 22-2.5 lists the 2014 Indiana Fire Code, based on the 2012 IFC
    First Printing, as currently in effect, effective 2014-12-01.
  confidence: 0.84
  source_ids:
  - src:usa-in:idhs-fpbsc-rules-overview
  - src:usa-in:idhs-fire-code-update
- topic: Local enforcement
  finding: IDHS performs Class 1 plan review and provides code-enforcement inspections,
    guidance, training, interpretation, and on-site inspection assistance; certified
    cities, towns, and counties may perform local plan review under 675 IAC 12-7.
  confidence: 0.78
  source_ids:
  - src:usa-in:idhs-building-plan-review
  - src:usa-in:idhs-code-enforcement
  - src:usa-in:iac-675-12-gar-pdf
- topic: Local amendments
  finding: "Substantive local deviations from Commission building/fire-safety rules\
    \ are generally unavailable under IDHS's HEA 1575 posture; administrative\
    \ local provisions remain."
  confidence: 0.74
  source_ids:
  - src:usa-in:idhs-fpbsc-overview
- topic: Effective / operative date rule
  finding: Design-release projects generally use Commission rules in effect when plans
    and specifications are filed with the Division; exempt construction uses rules
    in effect when work is actually done, with a rebuttable presumption tied to construction
    start.
  confidence: 0.82
  source_ids:
  - src:usa-in:iac-675-12-gar-pdf
  - src:usa-in:cornell-675-12-4-7
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | `ahj:usa-in:fpbsc` |
| Authority name | Indiana Fire Prevention and Building Safety Commission |
| Authority type | statewide commission / rulemaking body |
| Legal basis | Indiana Code 22-12-2; Title 675 IAC; 675 IAC 12 general administrative rules |
| Role | Creates statewide building and fire-safety laws; adopts rules under Title 675; reviews variances, rule-modification petitions, and enforcement orders. |
| Enforcement model | Statewide rules administered by IDHS/Division functions with local administration and certified local plan-review participation where allowed. |
| Source IDs | `src:usa-in:idhs-fpbsc-overview`; `src:usa-in:iac-675-12-gar-pdf`; `src:usa-in:idhs-fpbsc-rules-overview` |
| Verification status | partially_verified |

### 2.1.1 Supporting State Authorities

| Authority ID | Authority Name | Authority Type | Role | Source IDs | Verification Status |
| --- | --- | --- | --- | --- | --- |
| `ahj:usa-in:idhs-dfbs` | Indiana Department of Homeland Security / Division of Fire and Building Safety | state agency division | Supports administration of Commission rules, design-release functions, code enforcement assistance, and regulated-device programs. | `src:usa-in:iac-675-12-gar-pdf`; `src:usa-in:idhs-building-plan-review`; `src:usa-in:idhs-code-enforcement` | partially_verified |
| `ahj:usa-in:state-building-commissioner` | Indiana State Building Commissioner | state official | Oversees the Building Plan Review Section and Variance Section. | `src:usa-in:idhs-building-plan-review` | partially_verified |
| `ahj:usa-in:state-fire-marshal` | Indiana State Fire Marshal | state official | Oversees fire prevention/protection services, fire-code enforcement functions, regulated devices, fire/explosion investigations, and related IDHS sections. | `src:usa-in:idhs-code-enforcement`; `src:usa-in:idhs-fpbsc-rules-overview` | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | `ahj:usa-in:fpbsc` | Fire Prevention and Building Safety Commission | Adopts statewide building-code rules; reviews variances and modification petitions. | IC 22-12-2; IC 22-13; 675 IAC 13 | `src:usa-in:idhs-fpbsc-overview`; `src:usa-in:idhs-fpbsc-rules-overview` | partially_verified |
| Residential | `ahj:usa-in:fpbsc` | Fire Prevention and Building Safety Commission | Adopts statewide one- and two-family dwelling code rules. | 675 IAC 14 | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:idhs-rule-readoption-2025` | partially_verified |
| Existing Building / Rehabilitation | `ahj:usa-in:fpbsc` | Fire Prevention and Building Safety Commission | Administers existing-building, occupancy, alteration, and conversion rules through 675 IAC 12 and 675 IAC 13 references. | 675 IAC 12-4; 675 IAC 12-13; 675 IAC 13-2.6-34 | `src:usa-in:iac-675-12-gar-pdf`; `src:usa-in:cornell-675-12-4-12` | partially_verified |
| Mechanical | `ahj:usa-in:fpbsc` | Fire Prevention and Building Safety Commission | Adopts statewide mechanical-code rules. | 675 IAC 18 | `src:usa-in:idhs-fpbsc-rules-overview` | partially_verified |
| Plumbing | `ahj:usa-in:fpbsc` | Fire Prevention and Building Safety Commission | Adopts statewide plumbing-code rules; plumbing licensure and health-agency interfaces are outside this pass. | 675 IAC 16 | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:iar-plumbing-lsa-26-128` | partially_verified |
| Fuel Gas | `ahj:usa-in:fpbsc` | Fire Prevention and Building Safety Commission | Adopts statewide fuel-gas-code rules. | 675 IAC 25 | `src:usa-in:idhs-fpbsc-rules-overview` | partially_verified |
| Electrical | `ahj:usa-in:fpbsc` | Fire Prevention and Building Safety Commission | Adopts statewide electrical-code rules; 2026 update work is active. | 675 IAC 17 | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:idhs-electrical-code-update`; `src:usa-in:iar-electrical-lsa-26-62` | partially_verified |
| Energy | `ahj:usa-in:fpbsc` | Fire Prevention and Building Safety Commission | Adopts statewide energy-conservation-code rules. | 675 IAC 19 | `src:usa-in:idhs-fpbsc-rules-overview` | partially_verified |
| Fire - construction references | `ahj:usa-in:fpbsc` | Fire Prevention and Building Safety Commission | Adopts construction-related fire-code provisions and referenced NFPA standards. | 675 IAC 22; 675 IAC 28 | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:idhs-fire-code-update` | partially_verified |
| Fire - operational / prevention code | `ahj:usa-in:state-fire-marshal` | Indiana State Fire Marshal / IDHS Code Enforcement Section | Oversees fire prevention/protection services and fire-code enforcement functions; provides inspections, training, interpretation, and assistance statewide. | IC 22 fire-safety laws; 675 IAC 22; 675 IAC 28 | `src:usa-in:idhs-code-enforcement`; `src:usa-in:idhs-fpbsc-rules-overview` | partially_verified |
| Accessibility | `ahj:usa-in:fpbsc` | Fire Prevention and Building Safety Commission | Adopts accessibility requirements through the Indiana Building Code and ANSI A117.1 referenced by 675 IAC 13-2.6. | 675 IAC 13-2.6 | `src:usa-in:idhs-fpbsc-rules-overview` | partially_verified |
| Elevator / Conveyance | `ahj:usa-in:state-fire-marshal` | Indiana State Fire Marshal / Elevators and Amusement Rides Section | Regulates elevators, escalators, manlifts, hoists, amusement rides, and related devices under Commission rules and IDHS programs. | 675 IAC 21; 675 IAC 23 | `src:usa-in:idhs-code-enforcement`; `src:usa-in:idhs-fpbsc-rules-overview` | partially_verified |
| Boiler / Pressure Vessel | `ahj:usa-in:state-fire-marshal` | Indiana State Fire Marshal / Boilers and Pressure Vessels Section | Maintains periodic inspection program and applies 675 IAC 30. | 675 IAC 30 | `src:usa-in:idhs-code-enforcement`; `src:usa-in:idhs-fpbsc-rules-overview` | partially_verified |

### 2.3 Authority Hierarchy Notes

Indiana's code-administration model is statewide in substance and distributed in administration. The Commission adopts the statewide rules contained in Title 675 IAC. IDHS and its Division of Fire and Building Safety administer plan review, design releases, code-enforcement assistance, interpretations, regulated devices, and administrative support. Local governments may adopt and enforce administrative ordinances and may administer permitting/inspection programs, but IDHS states that substantive local building and fire-safety standards outside incorporation of Commission rules generally may not be adopted or enforced. Certified local plan reviewers can participate in Class 1 plan review, but the rule text keeps the Division design-release process central.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| `edge:usa-in:001` | `ahj:usa-in:fpbsc` | adopts_statewide_rules_for | Title 675 building, fire-safety, equipment, and related code articles | `src:usa-in:idhs-fpbsc-overview`; `src:usa-in:iac-675-12-gar-pdf`; `src:usa-in:idhs-fpbsc-rules-overview` | partially_verified |
| `edge:usa-in:002` | `ahj:usa-in:fpbsc` | delegates_administration_support_to | IDHS / Division of Fire and Building Safety | `src:usa-in:iac-675-12-gar-pdf`; `src:usa-in:idhs-building-plan-review` | partially_verified |
| `edge:usa-in:003` | `ahj:usa-in:state-building-commissioner` | oversees | Building Plan Review Section and Variance Section | `src:usa-in:idhs-building-plan-review` | partially_verified |
| `edge:usa-in:004` | `ahj:usa-in:idhs-dfbs` | reviews_for_compliance | Class 1 structure construction plans under Commission rules and state-adopted building codes | `src:usa-in:idhs-building-plan-review`; `src:usa-in:iac-675-12-gar-pdf` | partially_verified |
| `edge:usa-in:005` | `ahj:usa-in:idhs-dfbs` | may_certify | cities, towns, and counties for limited local plan review of Class 1 structures | `src:usa-in:iac-675-12-gar-pdf` | partially_verified |
| `edge:usa-in:006` | local_jurisdictions | may_adopt_and_enforce | administrative building/fire-safety ordinance provisions such as permits, plan review, inspection schedules, fines, and order review | `src:usa-in:idhs-fpbsc-overview` | partially_verified |
| `edge:usa-in:007` | local_jurisdictions | may_not_substantively_modify_without | Commission rulemaking or petition process | `src:usa-in:idhs-fpbsc-overview` | partially_verified |
| `edge:usa-in:008` | `ahj:usa-in:fpbsc` | publishes_nonrule_interpretations_followed_by | Commission, IDHS, local building/fire departments, and AHJs enforcing state building/fire-safety laws | `src:usa-in:idhs-nonrule-interpretations` | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | 2014 Indiana Building Code | International Building Code, 2012 Edition, First Printing; ANSI A117.1, 2009 Edition, First Printing | 2014 Indiana / 2012 IBC | current | null | 2014-12-01 | null | null | Design-release filing-date rule in 675 IAC 12-4-7; local substantive deviations generally unavailable under IDHS HEA 1575 posture. | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:iac-675-12-gar-pdf`; `src:usa-in:cornell-675-12-4-7` |
| Residential | 2020 Indiana Residential Code | 2018 International Residential Code for One- and Two-Family Dwellings, First Printing | 2020 Indiana / 2018 IRC | current | null | 2019-12-26 | null | null | Residential code readopted without changes in 2025 docket; project-specific transition beyond effective date not extracted. | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:idhs-rule-readoption-2025` |
| Existing Building / Rehabilitation | Existing-building rules within 675 IAC 12 and Indiana Building Code Chapter 34 references | 675 IAC 12-4; 675 IAC 12-13; Indiana Building Code Chapter 34 | administrative / 2014 IBC reference | current | null | null | null | null | Existing additions/alterations generally must comply with new-construction rules for the work scope, with explicit limits on triggering whole-building compliance. | `src:usa-in:iac-675-12-gar-pdf`; `src:usa-in:cornell-675-12-4-12` |
| Mechanical | 2014 Indiana Mechanical Code | International Mechanical Code, 2012 Edition, First Printing | 2014 Indiana / 2012 IMC | current | null | 2014-12-01 | null | null | Design-release filing-date rule where design release applies. | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:iac-675-12-gar-pdf` |
| Plumbing | 2012 Indiana Plumbing Code | International Plumbing Code, 2006 Edition, Second Printing | 2012 Indiana / 2006 IPC | current | null | 2012-12-24 | null | null | Design-release filing-date rule where design release applies; 2026 plumbing update proceeding pending. | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:iac-675-12-gar-pdf`; `src:usa-in:iar-plumbing-lsa-26-128` |
| Fuel Gas | 2014 Indiana Fuel Gas Code | International Fuel Gas Code, 2012 Edition, Second Printing | 2014 Indiana / 2012 IFGC | current | null | 2014-12-01 | null | null | Design-release filing-date rule where design release applies. | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:iac-675-12-gar-pdf` |
| Electrical | Indiana Electrical Code, 2009 Edition | NFPA 70 National Electrical Code, 2008 Edition, First Printing | 2009 Indiana / 2008 NEC | current; update pending | null | 2009-08-26 | null | null | Current code remains listed as in effect; 2026 rulemaking proposes 2026 Indiana Electrical Code based on 2023 NFPA 70. | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:idhs-electrical-code-update`; `src:usa-in:iar-electrical-lsa-26-62` |
| Energy | 2010 Indiana Energy Conservation Code | ANSI/ASHRAE 90.1 Energy Standard for Buildings Except Low-Rise Residential Buildings, 2007 Edition, I-P Edition | 2010 Indiana / ASHRAE 90.1-2007 | current | null | 2010-05-06 | null | null | Existing-building energy exception exists for certain alterations of buildings built before 1978; otherwise design-release rules may apply by project. | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:iac-675-12-gar-pdf`; `src:usa-in:cornell-675-12-4-12` |
| Fire - construction references | 2014 Indiana Fire Code | International Fire Code, 2012 Edition, First Printing | 2014 Indiana / 2012 IFC | current | null | 2014-12-01 | null | null | Design-release and existing-building rules apply where fire-code requirements are tied to construction; fire-code update committee is reviewing 2024 model code. | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:idhs-fire-code-update`; `src:usa-in:iac-675-12-gar-pdf` |
| Fire - operational / prevention code | 2014 Indiana Fire Code plus NFPA standards and supplemental fire-safety rules | 2012 IFC; NFPA standards listed in 675 IAC 28 | mixed current standards | current; supplemental subrule change pending | null | 2014-12-01 | null | null | NFPA standards and supplemental rules have separate section histories; 2025 readoption docket indicates repeal of 675 IAC 24-1. | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:idhs-rule-readoption-2025`; `src:usa-in:idhs-code-enforcement` |
| Accessibility | Indiana Building Code accessibility provisions | ANSI A117.1 Accessible and Usable Buildings and Facilities, 2009 Edition, First Printing, via 675 IAC 13-2.6 | 2009 ANSI A117.1 through 2014 IBC adoption | current | null | 2014-12-01 | null | null | Follows the applicable building-code transition rule. | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:iac-675-12-gar-pdf` |
| Elevator / Conveyance | Indiana Regulated Lifting Device Rules | ASME/ANSI A17.1 and related standards listed in 675 IAC 21 | mixed, including A17.1 2007 | current | null | null | null | null | Article-level rule dates vary by section; full device-specific applicability needs article parsing. | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:idhs-code-enforcement` |
| Boiler / Pressure Vessel | Indiana Boiler and Pressure Vessel Rules | ASME Boiler and Pressure Vessel Code 2019 edition; NBIC 2019; API standards listed in 675 IAC 30 | 2019 model set | current | null | 2019-07-01 | null | null | Article-level rules amended 2021-07-01; previous emergency-rule histories need separate extraction for historical projects. | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:idhs-code-enforcement`; `src:usa-in:idhs-rule-readoption-2025` |

### 3.2 Additional Administered Programs Captured in Source Registry

| Program | Rule Citation | Status | Effective / Amended Dates Captured | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- |
| Industrialized building systems and mobile structures | 675 IAC 15 | current | 1986 and 2005 effective dates captured at article-row level | `src:usa-in:idhs-fpbsc-rules-overview` | Not elevated to core building-code matrix because detailed applicability was outside this pass. |
| Swimming pool code | 675 IAC 20 | current | Article-row dates vary | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:idhs-rule-readoption-2025` | Readopted without changes in 2025 docket. |
| Amusement device code | 675 IAC 23 | current | Article-row dates vary | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:idhs-rule-readoption-2025` | Readopted without changes in 2025 docket. |
| Visitability rule | 675 IAC 27 | current | 2005-10-25 effective date captured at article-row level | `src:usa-in:idhs-fpbsc-rules-overview` | One- and two-family dwelling / townhouse visitability rule; not parsed section-by-section. |
| Outdoor event equipment | 675 IAC 29 | current | 2017-04-05 effective date captured | `src:usa-in:idhs-fpbsc-rules-overview` | Based on selected ASCE/SEI 7-10 provisions. |

### 3.3 Normalized Adoption Records

#### `adoption:usa-in:building:2014`

| Field | Value |
| --- | --- |
| authority_id | `ahj:usa-in:fpbsc` |
| code_family | Building |
| rule_citation | 675 IAC 13-2.6 |
| state_code_name | 2014 Indiana Building Code |
| base_model_code | International Building Code, 2012 Edition, First Printing; ANSI A117.1, 2009 Edition, First Printing |
| adoption_date | null |
| effective_date | 2014-12-01 |
| operative_date | null |
| mandatory_date | null |
| status | current |
| source_ids | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:iar-title-675-current` |

#### `adoption:usa-in:residential:2020`

| Field | Value |
| --- | --- |
| authority_id | `ahj:usa-in:fpbsc` |
| code_family | Residential |
| rule_citation | 675 IAC 14-4.4 |
| state_code_name | 2020 Indiana Residential Code |
| base_model_code | 2018 International Residential Code for One- and Two-Family Dwellings, First Printing |
| adoption_date | null |
| effective_date | 2019-12-26 |
| operative_date | null |
| mandatory_date | null |
| status | current |
| source_ids | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:idhs-rule-readoption-2025`; `src:usa-in:iar-title-675-current` |

#### `adoption:usa-in:plumbing:2012`

| Field | Value |
| --- | --- |
| authority_id | `ahj:usa-in:fpbsc` |
| code_family | Plumbing |
| rule_citation | 675 IAC 16-1.4 |
| state_code_name | 2012 Indiana Plumbing Code |
| base_model_code | International Plumbing Code, 2006 Edition, Second Printing |
| adoption_date | null |
| effective_date | 2012-12-24 |
| operative_date | null |
| mandatory_date | null |
| status | current; update pending |
| source_ids | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:iar-plumbing-lsa-26-128`; `src:usa-in:iar-title-675-current` |

#### `adoption:usa-in:electrical:2009`

| Field | Value |
| --- | --- |
| authority_id | `ahj:usa-in:fpbsc` |
| code_family | Electrical |
| rule_citation | 675 IAC 17-1.8 |
| state_code_name | Indiana Electrical Code, 2009 Edition |
| base_model_code | NFPA 70 National Electrical Code, 2008 Edition, First Printing |
| adoption_date | null |
| effective_date | 2009-08-26 |
| operative_date | null |
| mandatory_date | null |
| status | current; update pending |
| source_ids | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:idhs-electrical-code-update`; `src:usa-in:iar-electrical-lsa-26-62`; `src:usa-in:iar-title-675-current` |

#### `adoption:usa-in:mechanical:2014`

| Field | Value |
| --- | --- |
| authority_id | `ahj:usa-in:fpbsc` |
| code_family | Mechanical |
| rule_citation | 675 IAC 18-1.6 |
| state_code_name | 2014 Indiana Mechanical Code |
| base_model_code | International Mechanical Code, 2012 Edition, First Printing |
| adoption_date | null |
| effective_date | 2014-12-01 |
| operative_date | null |
| mandatory_date | null |
| status | current |
| source_ids | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:iar-title-675-current` |

#### `adoption:usa-in:energy:2010`

| Field | Value |
| --- | --- |
| authority_id | `ahj:usa-in:fpbsc` |
| code_family | Energy |
| rule_citation | 675 IAC 19-4 |
| state_code_name | 2010 Indiana Energy Conservation Code |
| base_model_code | ANSI/ASHRAE 90.1-2007, I-P Edition |
| adoption_date | null |
| effective_date | 2010-05-06 |
| operative_date | null |
| mandatory_date | null |
| status | current |
| source_ids | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:iar-title-675-current` |

#### `adoption:usa-in:fire:2014`

| Field | Value |
| --- | --- |
| authority_id | `ahj:usa-in:fpbsc` |
| code_family | Fire - operational / prevention code |
| rule_citation | 675 IAC 22-2.5 |
| state_code_name | 2014 Indiana Fire Code |
| base_model_code | International Fire Code, 2012 Edition, First Printing |
| adoption_date | null |
| effective_date | 2014-12-01 |
| operative_date | null |
| mandatory_date | null |
| status | current; update pending |
| source_ids | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:idhs-fire-code-update`; `src:usa-in:iar-title-675-current` |

#### `adoption:usa-in:fuel-gas:2014`

| Field | Value |
| --- | --- |
| authority_id | `ahj:usa-in:fpbsc` |
| code_family | Fuel Gas |
| rule_citation | 675 IAC 25-3 |
| state_code_name | 2014 Indiana Fuel Gas Code |
| base_model_code | International Fuel Gas Code, 2012 Edition, Second Printing |
| adoption_date | null |
| effective_date | 2014-12-01 |
| operative_date | null |
| mandatory_date | null |
| status | current |
| source_ids | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:iar-title-675-current` |

#### `adoption:usa-in:boiler-pressure-vessel:2019`

| Field | Value |
| --- | --- |
| authority_id | `ahj:usa-in:state-fire-marshal` |
| code_family | Boiler / Pressure Vessel |
| rule_citation | 675 IAC 30 |
| state_code_name | Indiana Boiler and Pressure Vessel Rules |
| base_model_code | ASME BPVC 2019; NBIC 2019; API standards as listed in 675 IAC 30 |
| adoption_date | null |
| effective_date | 2019-07-01 |
| operative_date | null |
| mandatory_date | null |
| status | current; amended 2021-07-01 |
| source_ids | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:idhs-rule-readoption-2025`; `src:usa-in:iar-title-675-current` |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

Indiana separates at least two core timing rules in 675 IAC 12-4-7. For construction projects requiring a design release under 675 IAC 12-6, construction must comply with the Commission rules in effect on the date the plans and specifications were filed with the Division. For construction that is exempt from design-release requirements but still subject to Commission rules, the applicable rules are those in effect when the construction work is actually done, with a rebuttable presumption that all construction for the exempt project is done on the date construction begins. Existing-building additions, alterations, repairs, occupancy changes, and conversion requests have separate rules in 675 IAC 12-4, 675 IAC 12-13, and related Indiana Building Code references.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `date-rule:usa-in:001` | Class 1 and other projects requiring design release under 675 IAC 12-6 | filing_date_controls | date plans and specifications are filed with Division | Construction requires design release. | Yes, if the prior rule was in effect on filing date; a later rule may be used if it becomes effective before the work and an appropriate addenda/revision design release is issued. | `src:usa-in:iac-675-12-gar-pdf`; `src:usa-in:cornell-675-12-4-7` | partially_verified |
| `date-rule:usa-in:002` | Construction exempt from design release but subject to Commission rules | work_date_controls | date construction work is actually done | Construction is exempt from design-release requirement and remains subject to Commission rules. | The applicable rule is tied to work date; there is a rebuttable presumption that exempt-project construction occurs on construction-start date. | `src:usa-in:iac-675-12-gar-pdf`; `src:usa-in:cornell-675-12-4-7` | partially_verified |
| `date-rule:usa-in:003` | Existing building additions and alterations | existing_building_scope_rule | current new-construction rules for scope of work, with existing-condition protections | Addition or alteration to an existing building, structure, or permanent system. | Existing building or system need not fully comply if the work scope meets new-construction rules and does not violate listed limits. | `src:usa-in:iac-675-12-gar-pdf`; `src:usa-in:cornell-675-12-4-12` | partially_verified |
| `date-rule:usa-in:004` | Repairs to permanent building systems | repair_rule | current rules or rules in effect at original installation | Repair to permanent HVAC, electrical, plumbing, sanitary, emergency detection/communication, fire, or explosion systems. | Yes, repairs may use current rules or rules in effect at original installation. | `src:usa-in:iac-675-12-gar-pdf`; `src:usa-in:cornell-675-12-4-12` | partially_verified |
| `date-rule:usa-in:005` | Design release expiration | expiration_rule | one year from design-release date, with one possible extension not exceeding 180 days if conditions are met | Construction work within scope of design release has not commenced. | Not a code-edition transition rule; it controls release validity. | `src:usa-in:iac-675-12-gar-pdf` | partially_verified |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | Potential update based on 2024 International Building Code | null | null | null | null | null | active_committee | `src:usa-in:idhs-building-code-update` | Committee is drafting a proposed rule; Commission retains full rulemaking authority. |
| Fire prevention | Potential update based on 2024 International Fire Prevention Code | null | null | null | null | null | active_committee | `src:usa-in:idhs-fire-code-update` | Committee is drafting a proposed rule; Commission retains full rulemaking authority. |
| Electrical | Proposed 2026 Indiana Electrical Code based on 2023 NFPA 70 National Electrical Code | 2026-03-11 | null | null | null | null | first_public_comment_period | `src:usa-in:idhs-electrical-code-update`; `src:usa-in:iar-electrical-lsa-26-62` | Docket materials describe repeal of the 2009 Indiana Electrical Code and replacement with 2026 Indiana Electrical Code; no final effective date captured. |
| Plumbing | Indiana Plumbing Code update proceeding | 2026-05-20 | null | null | null | null | first_public_comment_period | `src:usa-in:iar-plumbing-lsa-26-128` | First public comment period/public hearing notices located; proposed text and final adoption details not parsed here. |
| Residential / Electrical / Pool / Amusement / Boiler | Rule readoptions without changes; repeal of 675 IAC 24-1 | 2025-09-03 | null | null | null | null | monitor_final_effective_text | `src:usa-in:idhs-rule-readoption-2025`; `src:usa-in:iar-readoption-lsa-24-566` | IDHS docket lists readoptions and proposed repeal. Indiana Register final-rule page is JavaScript-gated in this pass. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| `applicability-rule:usa-in:001` | Building / multi-code | Class 1 structure construction | Construction on a Class 1 structure unless exempt under 675 IAC 12-6-4 | No construction may be done on a Class 1 structure until the Division issues a design release, unless a specific exemption applies. | `src:usa-in:iac-675-12-gar-pdf`; `src:usa-in:idhs-building-plan-review` | partially_verified |
| `applicability-rule:usa-in:002` | Building / fire / mechanical / plumbing / electrical / fuel gas / energy | Local certified plan review | City, town, or county has qualified plan reviewer and applies for certification | Certified local plan review is limited to Class 1 plan review and excludes industrialized building systems, automatic fire-extinguishing or standpipe systems, regulated lifting devices, and boilers/pressure vessels. | `src:usa-in:iac-675-12-gar-pdf` | partially_verified |
| `applicability-rule:usa-in:003` | Existing buildings | Additions, alterations, repairs, occupancy changes | Existing building or permanent building-system work | Existing-building work has separate scope and non-degradation rules; not every alteration forces full building compliance. | `src:usa-in:iac-675-12-gar-pdf`; `src:usa-in:cornell-675-12-4-12` | partially_verified |
| `applicability-rule:usa-in:004` | Local ordinances | Local building/fire-safety ordinances | Local jurisdiction adopts ordinance after HEA 1575 framework | IDHS distinguishes allowed administrative provisions from generally unavailable substantive building/fire-safety standards outside incorporation of Commission rules. | `src:usa-in:idhs-fpbsc-overview` | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** Indiana adopts model codes through Title 675 IAC provisions and modifies those model codes in state rule text. The state's code-navigation page lists the model code editions adopted by reference and notes that model codes are copyrighted; the official Indiana rule text and history lines remain the primary legal source.

**Where amendments are published:** Indiana Administrative Code / Indiana Register under Title 675, plus IDHS courtesy PDFs and code-navigation pages. For user-facing reading, IDHS links to ICC free-access material for the residential code where available.

**Amendment parsing status:** article_level_captured; section_level_pending

### 5.2 State Amendment Sources

| Source Set | Citation / Location | Status | Source IDs | Notes |
| --- | --- | --- | --- | --- |
| Building Code amendments | 675 IAC 13-2.6 | article_level_captured | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:iar-title-675-current` | State code identified; individual amendments not extracted into a section-level change list. |
| Residential Code amendments | 675 IAC 14-4.4 | article_level_captured | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:idhs-rule-readoption-2025`; `src:usa-in:iar-title-675-current` | State code identified; free-access model-with-amendments link exists on IDHS page. |
| Plumbing Code amendments | 675 IAC 16-1.4 | article_level_captured | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:iar-title-675-current` | 2026 update proceeding pending. |
| Electrical Code amendments | 675 IAC 17-1.8 | article_level_captured | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:iar-electrical-lsa-26-62`; `src:usa-in:iar-title-675-current` | 2026 update proceeding pending. |
| Mechanical Code amendments | 675 IAC 18-1.6 | article_level_captured | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:iar-title-675-current` | Section-level amendments pending. |
| Fire Code and NFPA standards | 675 IAC 22; 675 IAC 28 | article_level_captured | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:idhs-fire-code-update`; `src:usa-in:iar-title-675-current` | Operational/construction split needs section-level parsing before automated AHJ use. |
| Fuel Gas Code amendments | 675 IAC 25-3 | article_level_captured | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:iar-title-675-current` | Section-level amendments pending. |
| General administrative rules | 675 IAC 12 | partially_section_parsed | `src:usa-in:iac-675-12-gar-pdf`; `src:usa-in:cornell-675-12-4-7`; `src:usa-in:cornell-675-12-4-12` | Date rules, design release, existing-building, and local plan-review provisions were extracted. |

### 5.3 High-Impact State Amendments

No individual model-code amendment has been promoted to a high-impact amendment record in this pass. The following high-impact candidates are queued for section-level extraction:

| Candidate ID | Code Family | Candidate Topic | Why It Matters | Current Status | Source IDs |
| --- | --- | --- | --- | --- | --- |
| `amendment-candidate:usa-in:001` | Building / Accessibility | 675 IAC 13-2.6 accessibility modifications and ANSI A117.1 incorporation | Affects accessibility scope and construction compliance. | queued | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:iar-title-675-current` |
| `amendment-candidate:usa-in:002` | Existing Building | 675 IAC 12-4 existing-building and repair logic | Affects renovations, changes of occupancy, repairs, and enforcement. | partially_extracted | `src:usa-in:iac-675-12-gar-pdf`; `src:usa-in:cornell-675-12-4-12` |
| `amendment-candidate:usa-in:003` | Fire / Sprinkler / NFPA | 675 IAC 22 and 675 IAC 28 adopted standards and retained historical provisions | Affects fire-protection design, operational inspections, and AHJ interpretation. | queued | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:idhs-fire-code-update` |
| `amendment-candidate:usa-in:004` | Energy | 675 IAC 19 and pre-1978 alteration exception under 675 IAC 12-4-12 | Affects energy scope for alterations. | partially_extracted | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:iac-675-12-gar-pdf`; `src:usa-in:cornell-675-12-4-12` |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-in"
  model: "statewide_substantive_code_with_state_design_release_and_local_administration"
  enforcing_entities:
    - "Indiana Department of Homeland Security / Division of Fire and Building Safety"
    - "Building Plan Review Section"
    - "State Fire Marshal / Code Enforcement Section"
    - "local building departments"
    - "local fire departments"
    - "certified city, town, or county plan-review programs where approved"
  required_officials:
    - "State Building Commissioner for Building Plan Review and Variance Section oversight"
    - "State Fire Marshal for fire prevention/protection services and code-enforcement oversight"
    - "qualified local plan reviewer for certified local plan-review participation"
  state_reserved_activities:
    - "Commission rulemaking and variance review"
    - "Class 1 design release issuance under 675 IAC 12-6"
    - "certification and sanctioning of local plan-review programs"
    - "review of modifications to statewide standards through Commission rulemaking or petition process"
  source_ids:
    - "src:usa-in:idhs-fpbsc-overview"
    - "src:usa-in:idhs-building-plan-review"
    - "src:usa-in:idhs-code-enforcement"
    - "src:usa-in:iac-675-12-gar-pdf"
  verification_status: "partially_verified"
  confidence: 0.78
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-in"
  model: "statewide_unamended_substantive_rules_with_local_administrative_ordinances"
  applies_to_code_families:
    - "building"
    - "fire_safety"
    - "other Commission building/fire-safety rules where substantive standards are at issue"
  approval_required: false
  approving_authority_id: null
  filing_required: null
  registry_exists: false
  registry_source_ids: []
  legal_basis_source_ids:
    - "src:usa-in:idhs-fpbsc-overview"
  verification_status: "partially_verified"
  confidence: 0.74
  notes:
    - "IDHS states that substantive local standards outside incorporation of Commission rules are generally unenforceable and may not be adopted."
    - "IDHS also states that local administrative provisions remain available."
    - "IDHS expressly notes that its statements are its position and not binding legal advice; local enforceability determinations remain with the administering local unit and counsel."
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement and local amendment authority are separate in Indiana. A local jurisdiction may administer permits, plan review, inspection schedules, fines, and order-review processes as administrative provisions, and a certified local plan reviewer may participate in Class 1 plan review under 675 IAC 12-7. Those functions do not create independent authority to enact more stringent, more detailed, or different substantive building/fire-safety standards. For substantive changes, the supported path is Commission rulemaking or a petition process to the Commission.

### 6.4 Known Local Amendment Registries

| Registry ID | Registry Name | Coverage | Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- |
| `registry:usa-in:local-amendments:none-found` | none located | statewide | none_supported | `src:usa-in:idhs-fpbsc-overview` | IDHS states the Commission no longer reviews or approves ordinances. No statewide local-amendment registry was located in this pass. |

### 6.5 Municipality-Specific Known Amendments

No municipality-specific substantive amendments were added. Under the IDHS HEA 1575 posture, municipality-specific records should be treated as administrative ordinance records unless legal review confirms a valid substantive provision or Commission rulemaking/petition outcome.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: initial_statewide_model_ready

Jurisdiction stack:

```text
Address
  -> State of Indiana
  -> County
  -> Municipality / unincorporated county
  -> Local administrative building department, if established
  -> Local fire department / fire prevention bureau, if applicable
  -> Certified local plan-review program, if present
  -> IDHS Division design-release workflow for Class 1 structures
  -> State Fire Marshal / Code Enforcement Section for state inspection, interpretation, and assistance pathways
  -> Applicable Title 675 statewide code adoption records
  -> Applicable local administrative ordinance records
  -> Commission variance / petition / interpretation records
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | TIGER/Line or Indiana state GIS source to be selected | none | statewide | annual or state cadence | pending_selection |
| County | TIGER/Line or Indiana county boundary dataset to be selected | none | statewide | annual or state cadence | pending_selection |
| Municipality | Indiana state GIS municipal boundary source or Census incorporated-place layer to be selected | none | statewide | variable | pending_selection |
| Fire District | State or local fire-district boundary source to be selected | none | incomplete until sourced | variable | pending_selection |
| Special District | State or local special-district sources to be selected | none | incomplete until sourced | variable | pending_selection |

### 7.3 AHJ Contact Data

No AHJ contact roster was populated. For production AHJ resolution, collect at least the following:

| Contact Dataset | Intended Use | Status | Notes |
| --- | --- | --- | --- |
| IDHS Building Plan Review contacts | State design-release routing and Class 1 plan questions | pending | IDHS plan-review page provides general contact pathways but was not normalized into a contact table. |
| IDHS Code Enforcement / State Fire Marshal contacts | Fire-code and inspection assistance routing | pending | Statewide functions confirmed; individual inspector territories were not extracted. |
| Certified local plan-review jurisdictions | Identify when local pre-review exists for Class 1 structures | pending | 675 IAC 12-7 process confirmed; current certified-jurisdiction list not located. |
| Local building departments | Permits, inspections, fees, and local administrative procedures | pending | Must be collected jurisdiction-by-jurisdiction. |
| Local fire departments / fire prevention bureaus | Fire inspections, operational permits, local administrative procedures | pending | Must be collected jurisdiction-by-jurisdiction. |

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Source Type | Publisher | Title / Locator | URL | Key Facts Used | Extraction Status |
| --- | --- | --- | --- | --- | --- | --- |
| `src:usa-in:idhs-fpbsc-overview` | agency_page | Indiana Department of Homeland Security | Boards and Commissions - Fire Prevention and Building Safety Commission | https://www.in.gov/dhs/boards-and-commissions/ | Commission role; statewide code authority; HEA 1575 local-ordinance posture; petition/rulemaking pathway. | extracted_2026-06-25 |
| `src:usa-in:idhs-fpbsc-rules-overview` | agency_page | Indiana Department of Homeland Security | Rules of the Indiana Fire Prevention and Building Safety Commission | https://www.in.gov/dhs/boards-and-commissions/fpbsc-rules/ | Current Title 675 code-adoption matrix, model editions, effective dates, and caveat that page is informal. | extracted_2026-06-25 |
| `src:usa-in:iac-675-12-gar-pdf` | official_pdf_copy | Indiana Department of Homeland Security / Indiana Administrative Code | 675 IAC 12 General Administrative Rules | https://www.in.gov/dhs/files/675-IAC-12-General-Administrative-Rules.pdf | Commission/division definitions, organization of Title 675, design-release date rule, design-release requirements, local plan-review certification, existing-building rules. | extracted_2026-06-25 |
| `src:usa-in:idhs-building-plan-review` | agency_page | Indiana Department of Homeland Security | Building Plan Review | https://www.in.gov/dhs/building-plan-review/ | Class 1 plan review function; State Building Commissioner oversight; state-adopted building codes. | extracted_2026-06-25 |
| `src:usa-in:idhs-code-enforcement` | agency_page | Indiana Department of Homeland Security | State Fire Marshal / Code Enforcement | https://www.in.gov/dhs/fire-and-building-safety/ | State Fire Marshal oversight; Code Enforcement Section inspections, training, guidance, interpretations, and local department assistance; regulated devices and boiler program context. | extracted_2026-06-25 |
| `src:usa-in:idhs-nonrule-interpretations` | agency_page | Indiana Department of Homeland Security | Non-Rule Policy Documents / Interpretations of the Commission's Rules | https://www.in.gov/dhs/boards-and-commissions/non-rule-policy-documentsinterpretations-of-the-commissions-rules/ | Non-rule interpretations generally followed by Commission, IDHS, local departments, and AHJs enforcing state building/fire-safety laws. | extracted_2026-06-25 |
| `src:usa-in:idhs-building-code-update` | agency_page | Indiana Department of Homeland Security | Building Code Update Committee | https://www.in.gov/dhs/boards-and-commissions/building-code-update-committee/ | Pending work to draft proposed rule updating current Indiana Building Code at 675 IAC 13-2.6 using 2024 IBC; Commission retains authority. | extracted_2026-06-25 |
| `src:usa-in:idhs-fire-code-update` | agency_page | Indiana Department of Homeland Security | Fire Prevention Code Update Committee | https://www.in.gov/dhs/boards-and-commissions/fire-prevention-update-committee/ | Pending work to draft proposed rule updating current fire prevention code at 675 IAC 22-2.5 using 2024 model code; Commission retains authority. | extracted_2026-06-25 |
| `src:usa-in:idhs-electrical-code-update` | agency_page | Indiana Department of Homeland Security | Indiana Electrical Code Update Committee | https://www.in.gov/dhs/boards-and-commissions/indiana-electrical-code-update-committee/ | Review of proposed amendments to 2023 NFPA 70 NEC; path to Commission review and rule promulgation. | extracted_2026-06-25 |
| `src:usa-in:idhs-rule-readoption-2025` | agency_pdf | Indiana Department of Homeland Security / Fire Prevention and Building Safety Commission | Rulemaking Docket: Rule Readoption, LSA Document No. 24-566 | https://www.in.gov/dhs/files/Rulemaking-Docket-Rule-Readoptions-2025.pdf | Readoptions without changes for residential, electrical, pool, amusement, and boiler rules; repeal of 675 IAC 24-1; public-comment timetable. | extracted_2026-06-25 |
| `src:usa-in:iar-readoption-lsa-24-566` | register_notice | Indiana Register / Legislative Services Agency | LSA Document 24-566, Rule Readoption | https://iar.iga.in.gov/register/20251210-IR-675240566RFA | Final-rule locator for readoption docket; page content was JavaScript-gated in browser tool but search result confirms official Register page. | locator_verified_2026-06-25 |
| `src:usa-in:iar-electrical-lsa-26-62` | register_notice / agency_pdf | Indiana Register / IDHS | LSA Document 26-62, Indiana Electrical Code, 2026 Edition | https://iar.iga.in.gov/register/20260311-IR-675260062FNA | First public comment period and proposed electrical-code update path; IDHS docket summary indicates 2026 Indiana Electrical Code replacing 2009 code. | extracted_2026-06-25 |
| `src:usa-in:iar-plumbing-lsa-26-128` | register_notice | Indiana Register / Legislative Services Agency | LSA Document 26-128, Indiana Plumbing Code | https://iar.iga.in.gov/register/20260520-IR-675260128FNA | First public comment period for plumbing-code update proceeding. | locator_verified_2026-06-25 |
| `src:usa-in:iar-title-675-current` | administrative_code_portal | Indiana Register / Indiana Administrative Code | Title 675 Fire Prevention and Building Safety Commission | https://iar.iga.in.gov/latestTitle/675 | Official rule-text portal for Title 675 current rules; used as official locator backing DHS article list. | locator_verified_2026-06-25 |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| `src:usa-in:idhs-fpbsc-rules-overview` | informal_agency_index | IDHS says the rules overview is an informal navigation resource and not a substitute for legal research or the provisions/history lines of 675 IAC. | use_for_triage; verify against Indiana Register/IAC for production |
| `src:usa-in:iac-675-12-gar-pdf` | courtesy_pdf | PDF appears to be an Indiana Administrative Code copy hosted by IDHS; it should be spot-checked against official Register/IAC text before final legal publication. | acceptable_for_partial_verification; verify critical fields |
| `src:usa-in:iar-title-675-current` | javascript_gated | Official Indiana Register / Administrative Code portal content is JavaScript-heavy in the browser tool. | preserve as official locator; use available agency/PDF text for extraction until portal text is accessible |
| `src:usa-in:idhs-fpbsc-overview` | agency_position_caveat | IDHS states its HEA 1575 local-ordinance discussion is the agency's position and not binding legal advice. | treat local-amendment conclusion as partially_verified |
| `src:usa-in:iar-readoption-lsa-24-566` | javascript_gated | Official final-rule page was located but not text-extracted because the portal returned a JavaScript app shell. | monitor and verify final effective text through Register export or official filing copy |
| `src:usa-in:iar-plumbing-lsa-26-128` | javascript_gated | Official notice page was located but not text-extracted beyond search-result metadata. | use as pending-change locator; do not infer final code text |

### 8.3 Supplemental Sources

| Source ID | Source Type | Publisher | Title / Locator | URL | Key Facts Used | Production Treatment |
| --- | --- | --- | --- | --- | --- | --- |
| `src:usa-in:cornell-675-12-4-7` | supplemental_rule_text | Legal Information Institute / Cornell Law School | 675 IAC 12-4-7 Application of rules to new construction | https://www.law.cornell.edu/regulations/indiana/675-IAC-12-4-7 | Readable date-rule text for design-release filing date and exempt-construction work date. | corroborates official PDF; verify against Indiana Register for final production |
| `src:usa-in:cornell-675-12-4-12` | supplemental_rule_text | Legal Information Institute / Cornell Law School | 675 IAC 12-4-12 Existing buildings; additions or alterations | https://www.law.cornell.edu/regulations/indiana/675-IAC-12-4-12 | Readable existing-building alteration, repair, and energy exception text. | corroborates official PDF; verify against Indiana Register for final production |
| `src:usa-in:cornell-675-13-2.6` | supplemental_rule_text | Legal Information Institute / Cornell Law School | 675 IAC 13-2.6 2014 Indiana Building Code | https://www.law.cornell.edu/regulations/indiana/675-IAC-13-2.6 | Supplemental locator for building-code rule structure and Chapter 34 reference. | queued; not used for core fields except as a locator candidate |

### 8.4 Source Extraction Metadata

| Extraction ID | Date | Method | Sources | Extracted By | Notes |
| --- | --- | --- | --- | --- | --- |
| `extract:usa-in:2026-06-25:official-pages` | 2026-06-25 | web_open | `src:usa-in:idhs-fpbsc-overview`; `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:idhs-building-plan-review`; `src:usa-in:idhs-code-enforcement`; `src:usa-in:idhs-nonrule-interpretations` | GPT-5.5 Thinking | Official agency HTML parsed for authority, current code list, and administration fields. |
| `extract:usa-in:2026-06-25:675-12-pdf` | 2026-06-25 | pdf_text_and_screenshot | `src:usa-in:iac-675-12-gar-pdf` | GPT-5.5 Thinking | PDF text searched for date rules, design-release requirements, local plan review, and existing-building provisions; key PDF pages were visually checked. |
| `extract:usa-in:2026-06-25:pending-rulemaking` | 2026-06-25 | web_search_and_open | `src:usa-in:idhs-building-code-update`; `src:usa-in:idhs-fire-code-update`; `src:usa-in:idhs-electrical-code-update`; `src:usa-in:idhs-rule-readoption-2025`; `src:usa-in:iar-electrical-lsa-26-62`; `src:usa-in:iar-plumbing-lsa-26-128` | GPT-5.5 Thinking | Pending changes recorded only at docket/committee level where final dates were not available. |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 1.00 | none | Status reflects source-backed authority and adoption fields with unresolved section-level amendment and AHJ details. |
| report | risk.overall_confidence | 0.72 | verified | 1.00 | none | Weighted by official-source coverage for core fields and remaining gaps for local/AHJ details. |
| `ahj:usa-in:fpbsc` | authority_name | Indiana Fire Prevention and Building Safety Commission | partially_verified | 0.90 | `src:usa-in:idhs-fpbsc-overview`; `src:usa-in:iac-675-12-gar-pdf` | Statute-level citations identified through agency and rule text; official IC text should be captured directly in next pass. |
| `adoption:usa-in:building:2014` | effective_date | 2014-12-01 | partially_verified | 0.86 | `src:usa-in:idhs-fpbsc-rules-overview` | IDHS table gives current/effective status; official IAR history line should be captured for full verification. |
| `adoption:usa-in:residential:2020` | effective_date | 2019-12-26 | partially_verified | 0.86 | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:idhs-rule-readoption-2025` | Readoption docket confirms ongoing necessity without changes. |
| `adoption:usa-in:electrical:2009` | current_status | current; update pending | partially_verified | 0.80 | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:idhs-electrical-code-update`; `src:usa-in:iar-electrical-lsa-26-62` | Current edition and pending update both confirmed; final rule not captured. |
| `adoption:usa-in:fire:2014` | effective_date | 2014-12-01 | partially_verified | 0.84 | `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:idhs-fire-code-update` | Fire code update committee confirms 675 IAC 22-2.5 is current code being updated. |
| `date-rule:usa-in:001` | rule_type | filing_date_controls | partially_verified | 0.82 | `src:usa-in:iac-675-12-gar-pdf`; `src:usa-in:cornell-675-12-4-7` | Official PDF and supplemental text agree. |
| `local-amendment-rule:usa-in` | model | statewide_unamended_substantive_rules_with_local_administrative_ordinances | partially_verified | 0.74 | `src:usa-in:idhs-fpbsc-overview` | IDHS caveat prevents final verification without direct statutory/legal review. |
| `local-enforcement:usa-in` | model | statewide_substantive_code_with_state_design_release_and_local_administration | partially_verified | 0.78 | `src:usa-in:idhs-building-plan-review`; `src:usa-in:idhs-code-enforcement`; `src:usa-in:iac-675-12-gar-pdf` | Contact roster and certified local plan-review list remain pending. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Every `src:usa-in:*` cited in body is listed in section 8. |
| All authority IDs resolve | pass | Core authority IDs are used consistently; contact-level local AHJ IDs are intentionally deferred. |
| All current code families have adoption records | partial | Core families have rows; section-level amendments and some ancillary programs need deeper records. |
| Building and operational fire code are separated | pass | Matrix separates construction fire references from operational/prevention code. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Unknown adoption/operative/mandatory dates remain null instead of inferred. |
| Effective dates are valid ISO dates | pass | Captured effective dates use YYYY-MM-DD format. |
| No impossible date sequences | pass | No conflicting date sequences introduced. |
| Transition rules have explicit trigger conditions | pass | Design-release, exempt-construction, and existing-building triggers are stated. |
| Permit-date logic is captured where applicable | partial | Design-release filing-date logic captured; local permit-date logic remains jurisdiction-specific. |
| Local enforcement model classified | pass | Classified as statewide substantive code with state design release and local administration. |
| Local amendment rule classified | pass | Classified with IDHS caveat. |
| AHJ confirmation metadata present | partial | State-level AHJ roles captured; local contact data and boundaries remain pending. |
| Official-source caveats captured | pass | Informal index, courtesy PDF, JavaScript-gated Register pages, and IDHS legal-position caveats are explicit. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `issue:usa-in:001` | high | official IAC text | DHS rules page is an informal index and Indiana Register pages are JavaScript-gated in this pass. | Capture/export official current 675 IAC sections and history lines for all core adoptions. | null | null | open |
| `issue:usa-in:002` | high | section-level amendments | Current code editions are captured, but individual Indiana amendments are not parsed into change records. | Parse 675 IAC 13, 14, 16, 17, 18, 19, 22, 25, 28, and 30 section-by-section. | null | null | open |
| `issue:usa-in:003` | high | 2026 rulemakings | Electrical, plumbing, building, and fire update proceedings may change current editions. | Monitor LSA 26-62, LSA 26-128, building-code committee, and fire-code committee until final rule/effective dates are known. | null | null | open |
| `issue:usa-in:004` | medium | fire authority details | State Fire Marshal role is confirmed at agency level, but detailed fire-inspection powers, local fire official duties, and operational permit rules were not fully extracted. | Extract IC 22 fire-safety provisions, 675 IAC 22, 675 IAC 28, and IDHS Code Enforcement details. | null | null | open |
| `issue:usa-in:005` | medium | certified local plan review list | 675 IAC 12-7 process is captured, but current certified cities/towns/counties were not located. | Find or request current certification roster from IDHS. | null | null | open |
| `issue:usa-in:006` | medium | local administrative ordinances | IDHS local-amendment posture is captured, but local administrative ordinance schemas and sample fields are not normalized. | Build local ordinance metadata model separating permits, fees, inspections, order review, and substantive-code petitions. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| `watch:usa-in:fpbsc-rules` | `src:usa-in:idhs-fpbsc-rules-overview` | html_diff | monthly | current/effective code rows change, caveat changes, new article rows added | 2026-06-25 | active |
| `watch:usa-in:title-675` | `src:usa-in:iar-title-675-current` | official_code_diff | monthly | Title 675 rule text, history lines, or articles updated | 2026-06-25 | active |
| `watch:usa-in:electrical-26-62` | `src:usa-in:iar-electrical-lsa-26-62` | register_docket | weekly | final rule filed, Attorney General/Governor approval, effective date posted | 2026-06-25 | active |
| `watch:usa-in:plumbing-26-128` | `src:usa-in:iar-plumbing-lsa-26-128` | register_docket | weekly | public comments, agency response, final rule, effective date posted | 2026-06-25 | active |
| `watch:usa-in:building-code-update` | `src:usa-in:idhs-building-code-update` | committee_page_diff | monthly | draft proposed rule, committee recommendation, public-comment notice, or rulemaking docket added | 2026-06-25 | active |
| `watch:usa-in:fire-code-update` | `src:usa-in:idhs-fire-code-update` | committee_page_diff | monthly | draft proposed rule, committee recommendation, public-comment notice, or rulemaking docket added | 2026-06-25 | active |
| `watch:usa-in:readoption-24-566` | `src:usa-in:iar-readoption-lsa-24-566` | register_docket | monthly | final effective text or repeal of 675 IAC 24-1 confirmed in accessible official text | 2026-06-25 | active |
| `watch:usa-in:hea-1575-ordinances` | `src:usa-in:idhs-fpbsc-overview` | legal_position_diff | quarterly | IDHS changes HEA 1575 local-ordinance guidance or Commission ordinance process | 2026-06-25 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-25 | Replaced baseline draft with source-backed partially verified Indiana report | `report:usa-in`; `ahj:usa-in:fpbsc`; core adoption records; local enforcement/amendment records | `src:usa-in:idhs-fpbsc-overview`; `src:usa-in:idhs-fpbsc-rules-overview`; `src:usa-in:iac-675-12-gar-pdf`; `src:usa-in:idhs-building-plan-review`; `src:usa-in:idhs-code-enforcement` | GPT-5.5 Thinking | Core authority, adoption matrix, date rules, local amendment posture, source registry, QA, open issues, and monitoring added. |
