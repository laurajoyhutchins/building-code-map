---
state:
  state_id: "US-ME"
  name: "Maine"
  abbreviation: "ME"
report:
  report_id: "state-report:usa-me"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-25"
  last_verified: null
  reviewed_by: null
risk:
  overall_confidence: 0.62 # 0.00 - 1.00
  risk_flags:
    - "fire_code_amendments_not_fully_parsed"
    - "plumbing_rule_dates_not_fully_parsed"
    - "elevator_code_editions_unresolved"
    - "local_amendment_scope_requires_legal_review"
    - "ahj_contacts_not_populated"
  open_questions_count: 6

---

# State Building Code Authority Report: Maine

## 1. Executive Summary

- **Authority model:** Maine has a statewide building-code authority model for the Maine Uniform Building and Energy Code (MUBEC). The Technical Building Codes and Standards Board, housed in the Maine Office of Community Affairs, is the primary statewide authority for adopting, amending, and maintaining MUBEC. Fire/life-safety rules are handled separately under the Department of Public Safety and the Office of State Fire Marshal. Electrical, plumbing, fuel-gas, accessibility, and elevator/conveyance requirements are administered through distinct boards, statutes, or state programs.

- **Statewide code status:** MUBEC is statewide in scope and currently includes the 2021 IBC, 2021 IRC, 2021 IEBC, 2021 IECC, 2021 IMC, 2019 ASHRAE 62.1, 2019 ASHRAE 62.2, CSA-F326-M91 as an alternative ventilation standard for single-family dwellings, ASHRAE 90.1, and ASTM E1465-2008. State agency pages identify the current amended MUBEC adoption as enforceable or in effect beginning **2025-04-07**.

- **Local enforcement model:** MUBEC applies statewide, but the enforcement obligation is population-based. Municipalities with more than 4,000 residents must enforce MUBEC. Municipalities with 4,000 or fewer residents are not required to enforce MUBEC, but if they adopt or enforce a building code it must be MUBEC, the Maine Uniform Building Code, the Maine Uniform Energy Code, or another option expressly allowed by statute. Mandatory enforcement municipalities may use local building officials, code enforcement officers, interlocal or county/regional arrangements, or certified third-party inspectors.

- **Local amendment posture:** Local enforcement is distinct from local amendment authority. Political-subdivision ordinances inconsistent with MUBEC are void except where statute preserves local enforcement ordinances and specific exceptions. The Technical Building Codes and Standards Board has a formal amendment process through which municipalities, county/regional governments, state agencies, professional associations, and members of the public may propose amendments. The optional energy/stretch appendix is available for voluntary municipal adoption.

- **Known transition periods or pending changes:** The major verified transition in this pass is the amended MUBEC adoption effective **2025-04-07**. Electrical installations commencing on or after **2024-07-01** must comply with the 2023 NEC as adopted by the Electricians' Examining Board. No future statewide MUBEC effective date was verified in this pass.

- **Production readiness:** partially_ready_for_narrow_validation

### Key Findings

```yaml
---
key_findings:
- topic: Primary statewide MUBEC authority
  finding: Technical Building Codes and Standards Board adopts, amends, and maintains
    MUBEC and resolves conflicts with fire/life-safety codes.
  confidence: 0.85
  source_ids:
  - src:usa-me:statute-10-9722
- topic: Current MUBEC model-code editions
  finding: Current agency-published MUBEC adoption includes the 2021 IBC, IRC, IEBC,
    IECC, and IMC, plus listed ASHRAE/CSA/ASTM standards.
  confidence: 0.85
  source_ids:
  - src:usa-me:fmo-building-codes
  - src:usa-me:energy-building-codes
  - src:usa-me:mubec-standards-amendments
- topic: MUBEC effective/enforcement date
  finding: Maine agency sources identify the current amended MUBEC adoption as in
    effect/enforceable beginning 2025-04-07.
  confidence: 0.82
  source_ids:
  - src:usa-me:fmo-building-codes
  - src:usa-me:energy-building-codes
  - src:usa-me:mubec-standards-amendments
- topic: Local MUBEC enforcement threshold
  finding: Municipalities over 4,000 must enforce MUBEC; municipalities of 4,000 or
    fewer are not required to enforce it but are constrained if they choose to adopt
    or enforce a building code.
  confidence: 0.85
  source_ids:
  - src:usa-me:statute-10-9724
  - src:usa-me:statute-25-2373
  - src:usa-me:fmo-building-codes
- topic: Local amendment / preemption posture
  finding: Inconsistent local ordinances are void except preserved local enforcement
    ordinances and specific statutory exceptions; amendments flow through the statewide
    Board process.
  confidence: 0.75
  source_ids:
  - src:usa-me:statute-10-9722
  - src:usa-me:statute-10-9724
- topic: Electrical installation code
  finding: The Electricians' Examining Board adopted the 2023 NEC with amendments/exclusions;
    installations commencing on or after 2024-07-01 must comply.
  confidence: 0.9
  source_ids:
  - src:usa-me:eeb-home
  - src:usa-me:eeb-ch120
- topic: Fire/life-safety authority
  finding: The Department of Public Safety / State Fire Marshal has statewide fire/life-safety
    authority; published agency lists include NFPA 1-2018 and NFPA 101-2018.
  confidence: 0.75
  source_ids:
  - src:usa-me:statute-25-2396
  - src:usa-me:statute-25-2452
  - src:usa-me:fmo-rules
  - src:usa-me:fmo-nfpa-standards
- topic: Accessibility / barrier-free construction
  finding: Maine statutory construction standards reference the 2010 ADA Standards;
    the State Fire Marshal has plan-review and permit roles for covered public-building
    work.
  confidence: 0.78
  source_ids:
  - src:usa-me:statute-5-4594-g
  - src:usa-me:fmo-barrier-free
  - src:usa-me:fmo-plans-review
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-me:tbcands-board |
| Authority name | Technical Building Codes and Standards Board |
| Authority type | statewide board within Maine Office of Community Affairs |
| Legal basis | 10 M.R.S. § 9722 |
| Role | Adopt, amend, and maintain MUBEC; provide training for municipal building officials, code enforcement officers, and third-party inspectors; resolve conflicts between MUBEC and fire/life-safety codes under 25 M.R.S. §§ 2452 and 2465. |
| Enforcement model | Statewide code with population-based local enforcement and certified third-party inspection options. |
| Source IDs | src:usa-me:statute-10-9722; src:usa-me:statute-10-9723; src:usa-me:statute-10-9724; src:usa-me:statute-25-2373 |
| Verification status | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-me:tbcands-board | Technical Building Codes and Standards Board | Adopts and amends the Commercial Building Code of Maine / MUBEC building provisions. | 10 M.R.S. § 9722; MUBEC rules | src:usa-me:statute-10-9722; src:usa-me:moca-code-rules; src:usa-me:fmo-mubec-rules | partially_verified |
| Residential | ahj:usa-me:tbcands-board | Technical Building Codes and Standards Board | Adopts and amends the Residential Building Code for One- and Two-Family Dwellings in Maine. | 10 M.R.S. § 9722; MUBEC rules | src:usa-me:statute-10-9722; src:usa-me:moca-code-rules; src:usa-me:fmo-mubec-rules | partially_verified |
| Existing Building / Rehabilitation | ahj:usa-me:tbcands-board | Technical Building Codes and Standards Board | Adopts and amends the Existing Building Code of Maine. | 10 M.R.S. § 9722; MUBEC rules | src:usa-me:statute-10-9722; src:usa-me:moca-code-rules; src:usa-me:fmo-mubec-rules | partially_verified |
| Mechanical | ahj:usa-me:tbcands-board | Technical Building Codes and Standards Board | Adopts and amends the Mechanical Code of Maine. | 10 M.R.S. § 9722; MUBEC rules | src:usa-me:statute-10-9722; src:usa-me:moca-code-rules; src:usa-me:fmo-mubec-rules | partially_verified |
| Plumbing | ahj:usa-me:plumbers-examining-board | Plumbers' Examining Board | Publishes/adopts the Maine State Internal Plumbing Code based on the 2021 Uniform Plumbing Code with Maine amendments. | Board rule authority to be extracted; agency page verified in this pass | src:usa-me:moca-plumbing | partially_verified |
| Fuel Gas | ahj:usa-me:fuel-board | Maine Fuel Board | Adopts fuel-burning equipment and gas standards, including NFPA 54 and NFPA 58, with amendments. | 32 M.R.S. ch. 139 as implemented by Board rules | src:usa-me:fuel-board-laws-rules; src:usa-me:fuel-board-rules | partially_verified |
| Electrical | ahj:usa-me:electricians-examining-board | Electricians' Examining Board | Adopts electrical installation standards and licenses/permitting program; current rule incorporates the 2023 NEC with amendments/exclusions. | 32 M.R.S. §§ 1153, 1153-A; 02-318 C.M.R. ch. 120 | src:usa-me:eeb-home; src:usa-me:eeb-ch120 | partially_verified |
| Energy | ahj:usa-me:tbcands-board | Technical Building Codes and Standards Board | Adopts and amends the Maine Uniform Energy Code component of MUBEC and optional stretch-code appendix. | 10 M.R.S. § 9722; MUBEC rules | src:usa-me:statute-10-9722; src:usa-me:energy-building-codes; src:usa-me:moca-code-rules | partially_verified |
| Fire - construction references | ahj:usa-me:dps-state-fire-marshal | Department of Public Safety / Office of State Fire Marshal | Administers fire/life-safety review and fire-code requirements that interact with construction. | 25 M.R.S. §§ 2396, 2452 | src:usa-me:statute-25-2396; src:usa-me:statute-25-2452; src:usa-me:fmo-plans-review | partially_verified |
| Fire - operational / prevention code | ahj:usa-me:dps-state-fire-marshal | Department of Public Safety / Office of State Fire Marshal | Adopts/administers fire prevention and life-safety rules, including listed NFPA standards. | 25 M.R.S. §§ 2396, 2452; State Fire Marshal rules | src:usa-me:statute-25-2396; src:usa-me:statute-25-2452; src:usa-me:fmo-rules; src:usa-me:fmo-nfpa-standards | partially_verified |
| Accessibility | ahj:usa-me:state-fire-marshal-barrier-free | Office of State Fire Marshal / Maine Human Rights Act construction standards | Reviews and permits covered barrier-free/public-building construction; statutory standards reference the 2010 ADA Standards. | 5 M.R.S. § 4594-G | src:usa-me:statute-5-4594-g; src:usa-me:fmo-barrier-free; src:usa-me:fmo-plans-review | partially_verified |
| Elevator / Conveyance | ahj:usa-me:elevator-tramway-program | Elevator and Tramway Safety Program | Administers elevator/tramway inspection and plan-submission requirements under state rules and adopted national codes. | 32 M.R.S. ch. 133 and Program rules; exact current adopted editions unresolved | src:usa-me:elevator-program-laws-rules | unresolved |

### 2.3 Authority Hierarchy Notes

MUBEC is not a purely local code. The Technical Building Codes and Standards Board maintains the statewide building/energy/mechanical code framework, while local enforcement depends on municipal population and municipal administrative choices. Fire/life-safety authority is separate but linked: the Board is expressly tasked with resolving conflicts between MUBEC and fire/life-safety codes under 25 M.R.S. §§ 2452 and 2465.

Specialty trade systems are not fully collapsed into MUBEC. Electrical installations are governed through the Electricians' Examining Board and Chapter 120; fuel-gas and fuel-burning equipment requirements are governed through the Maine Fuel Board; plumbing is tied to the Plumbers' Examining Board's Maine State Internal Plumbing Code; elevator and tramway safety is administered through a separate program.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-me:001 | ahj:usa-me:tbcands-board | adopts_amends_maintains | MUBEC building, residential, existing-building, energy, and mechanical components | src:usa-me:statute-10-9722 | partially_verified |
| edge:usa-me:002 | ahj:usa-me:tbcands-board | resolves_conflicts_with | fire/life-safety codes under 25 M.R.S. §§ 2452 and 2465 | src:usa-me:statute-10-9722 | partially_verified |
| edge:usa-me:003 | ahj:usa-me:tbcands-board | relies_on_enforcement_by | municipalities, code enforcement officers, certified third-party inspectors, and interlocal/county/regional arrangements | src:usa-me:statute-25-2373 | partially_verified |
| edge:usa-me:004 | state_mubec | preempts_inconsistent | political-subdivision ordinances inconsistent with MUBEC, subject to statutory exceptions | src:usa-me:statute-10-9724 | partially_verified |
| edge:usa-me:005 | ahj:usa-me:dps-state-fire-marshal | administers | fire/life-safety, public-building plan review, and statewide NFPA 101 enforcement powers | src:usa-me:statute-25-2396; src:usa-me:statute-25-2452; src:usa-me:fmo-plans-review | partially_verified |
| edge:usa-me:006 | ahj:usa-me:electricians-examining-board | adopts | electrical installation standards / NEC adoption | src:usa-me:eeb-home; src:usa-me:eeb-ch120 | partially_verified |
| edge:usa-me:007 | ahj:usa-me:fuel-board | adopts | fuel-burning equipment and gas standards | src:usa-me:fuel-board-rules | partially_verified |
| edge:usa-me:008 | ahj:usa-me:state-fire-marshal-barrier-free | reviews_permits | covered barrier-free/public-building construction | src:usa-me:statute-5-4594-g; src:usa-me:fmo-barrier-free | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | Commercial Building Code of Maine / MUBEC building component | International Building Code | 2021 | current | null | 2025-04-07 | 2025-04-07 | 2025-04-07 for jurisdictions enforcing MUBEC | Current amended MUBEC is identified by state agency sources as enforceable/in effect beginning 2025-04-07. | src:usa-me:fmo-building-codes; src:usa-me:energy-building-codes; src:usa-me:mubec-standards-amendments; src:usa-me:moca-code-rules |
| Residential | Residential Building Code for One- and Two-Family Dwellings in Maine | International Residential Code | 2021 | current | null | 2025-04-07 | 2025-04-07 | 2025-04-07 for jurisdictions enforcing MUBEC | Same current MUBEC effective/enforcement date; statewide application but enforcement is population-based. | src:usa-me:fmo-building-codes; src:usa-me:energy-building-codes; src:usa-me:mubec-standards-amendments; src:usa-me:moca-code-rules |
| Existing Building / Rehabilitation | Existing Building Code of Maine | International Existing Building Code | 2021 | current | null | 2025-04-07 | 2025-04-07 | 2025-04-07 for jurisdictions enforcing MUBEC | Same current MUBEC effective/enforcement date. | src:usa-me:fmo-building-codes; src:usa-me:energy-building-codes; src:usa-me:mubec-standards-amendments; src:usa-me:moca-code-rules |
| Mechanical | Mechanical Code of Maine | International Mechanical Code | 2021 | current | null | 2025-04-07 | 2025-04-07 | 2025-04-07 for jurisdictions enforcing MUBEC | Same current MUBEC effective/enforcement date. | src:usa-me:fmo-building-codes; src:usa-me:energy-building-codes; src:usa-me:mubec-standards-amendments; src:usa-me:moca-code-rules |
| Plumbing | Maine State Internal Plumbing Code | Uniform Plumbing Code | 2021 | current per agency publication page | null | null | null | null | Plumbers' Examining Board adoption and Maine amendments verified at agency-publication level; rule filing/effective date still unresolved. | src:usa-me:moca-plumbing |
| Fuel Gas | Maine Fuel Board Chapter 6 fuel-gas standards | NFPA 54 / ANSI Z223.1 National Fuel Gas Code; NFPA 58 for LPG | 2021 NFPA 54; 2020 NFPA 58 | current per Board rules reviewed | null | 2023-09-16 | 2023-09-16 | null | Board rules state adopted fuel-gas and LPG standards with amendments; installation-specific transition detail beyond rule effective date remains unresolved. | src:usa-me:fuel-board-laws-rules; src:usa-me:fuel-board-rules |
| Electrical | Electrical Installation Standards | NFPA 70, National Electrical Code | 2023 | current | null | 2024-07-01 | 2024-07-01 | 2024-07-01 for installations commencing on or after that date | Installations commencing on or after 2024-07-01 must comply with the 2023 NEC as adopted and amended/excluded by Chapter 120. | src:usa-me:eeb-home; src:usa-me:eeb-ch120 |
| Energy | Maine Uniform Energy Code component of MUBEC | International Energy Conservation Code; ASHRAE 90.1 | 2021 IECC; 2019 ASHRAE 90.1 | current | null | 2025-04-07 | 2025-04-07 | 2025-04-07 for jurisdictions enforcing MUBEC | MUBEC energy component follows the current MUBEC effective/enforcement date; optional stretch appendix is voluntary for municipal adoption. | src:usa-me:statute-10-9722; src:usa-me:energy-building-codes; src:usa-me:mubec-standards-amendments; src:usa-me:moca-code-rules |
| Fire - construction references | State Fire Marshal fire/life-safety standards | NFPA 101 Life Safety Code and other listed NFPA standards | 2018 NFPA 101 listed | current per Fire Marshal list | null | null | null | null | Exact construction-permit transition date for all listed NFPA standards was not normalized in this pass. | src:usa-me:statute-25-2396; src:usa-me:statute-25-2452; src:usa-me:fmo-rules; src:usa-me:fmo-nfpa-standards |
| Fire - operational / prevention code | State Fire Marshal Fire Prevention Code | NFPA 1 Fire Prevention Code | 2018 | current per Fire Marshal list | null | 2019-11-27 | null | null | Fire Marshal rules page lists Chapter 3 Fire Prevention Code amended 2019-11-27 and identifies NFPA 1-2018; official MAPA-certified rule text still needed for court-grade use. | src:usa-me:fmo-rules; src:usa-me:fmo-nfpa-standards |
| Accessibility | Maine Human Rights Act construction standards / Barrier-free construction review | 2010 ADA Standards for Accessible Design | 2010 | current statutory standard verified | null | null | 2012-03-15 trigger date in statute | null | Applies to covered new construction and alterations where application or construction timing and cost thresholds are met; State Fire Marshal review/permit requirements apply to covered public-building work. | src:usa-me:statute-5-4594-g; src:usa-me:fmo-barrier-free; src:usa-me:fmo-plans-review |
| Elevator / Conveyance | Elevator and Tramway Safety Program rules | National elevator/tramway codes adopted by program rules | exact editions unresolved | unresolved | null | null | null | null | Program rules and plan-submission/inspection framework were identified, but current adopted edition normalization remains an open issue. | src:usa-me:elevator-program-laws-rules |

### 3.2 Adoption Records

#### record:usa-me:mubec-ibc-2021

| Field | Value |
| --- | --- |
| Code family | Building |
| State code name | Commercial Building Code of Maine / MUBEC building component |
| Base model code | International Building Code |
| Edition | 2021 |
| Effective date | 2025-04-07 |
| Operative date | 2025-04-07 |
| Mandatory date | 2025-04-07 for municipalities and other jurisdictions enforcing MUBEC |
| Source IDs | src:usa-me:fmo-building-codes; src:usa-me:energy-building-codes; src:usa-me:mubec-standards-amendments; src:usa-me:moca-code-rules |
| Notes | Adoption-date field remains null because the exact rule-adoption date was not extracted from certified rulemaking records. |

#### record:usa-me:mubec-irc-2021

| Field | Value |
| --- | --- |
| Code family | Residential |
| State code name | Residential Building Code for One- and Two-Family Dwellings in Maine |
| Base model code | International Residential Code |
| Edition | 2021 |
| Effective date | 2025-04-07 |
| Operative date | 2025-04-07 |
| Mandatory date | 2025-04-07 for municipalities and other jurisdictions enforcing MUBEC |
| Source IDs | src:usa-me:fmo-building-codes; src:usa-me:energy-building-codes; src:usa-me:mubec-standards-amendments; src:usa-me:moca-code-rules |
| Notes | Adoption-date field remains null because the exact rule-adoption date was not extracted from certified rulemaking records. |

#### record:usa-me:mubec-iecc-2021

| Field | Value |
| --- | --- |
| Code family | Energy |
| State code name | Maine Uniform Energy Code |
| Base model code | International Energy Conservation Code |
| Edition | 2021 IECC plus 2019 ASHRAE 90.1 |
| Effective date | 2025-04-07 |
| Operative date | 2025-04-07 |
| Mandatory date | 2025-04-07 for municipalities and other jurisdictions enforcing MUBEC |
| Source IDs | src:usa-me:statute-10-9722; src:usa-me:energy-building-codes; src:usa-me:mubec-standards-amendments; src:usa-me:moca-code-rules |
| Notes | The Board must make an optional energy-conservation appendix/stretch code available for voluntary municipal adoption. |

#### record:usa-me:electrical-nec-2023

| Field | Value |
| --- | --- |
| Code family | Electrical |
| State code name | Electrical Installation Standards |
| Base model code | NFPA 70, National Electrical Code |
| Edition | 2023 |
| Effective date | 2024-07-01 |
| Operative date | 2024-07-01 |
| Mandatory date | 2024-07-01 for installations commencing on or after that date |
| Source IDs | src:usa-me:eeb-home; src:usa-me:eeb-ch120 |
| Notes | Chapter 120 adopts/incorporates the 2023 NEC subject to Maine amendments and exclusions. |

#### record:usa-me:fuel-gas-nfpa54-2021

| Field | Value |
| --- | --- |
| Code family | Fuel Gas |
| State code name | Maine Fuel Board Chapter 6 |
| Base model code | NFPA 54 / ANSI Z223.1 National Fuel Gas Code |
| Edition | 2021 |
| Effective date | 2023-09-16 |
| Operative date | 2023-09-16 |
| Mandatory date | null |
| Source IDs | src:usa-me:fuel-board-rules |
| Notes | Installation-specific transition detail was not extracted beyond the rule effective date. |

#### record:usa-me:fire-nfpa1-2018

| Field | Value |
| --- | --- |
| Code family | Fire - operational / prevention code |
| State code name | State Fire Marshal Fire Prevention Code |
| Base model code | NFPA 1 Fire Prevention Code |
| Edition | 2018 |
| Effective date | 2019-11-27 |
| Operative date | null |
| Mandatory date | null |
| Source IDs | src:usa-me:fmo-rules; src:usa-me:fmo-nfpa-standards |
| Notes | The Fire Marshal rules page is not a certified MAPA rule copy; use certified rule text for final legal reliance. |

#### record:usa-me:plumbing-upc-2021

| Field | Value |
| --- | --- |
| Code family | Plumbing |
| State code name | Maine State Internal Plumbing Code |
| Base model code | IAPMO Uniform Plumbing Code |
| Edition | 2021 |
| Effective date | null |
| Operative date | null |
| Mandatory date | null |
| Source IDs | src:usa-me:moca-plumbing |
| Notes | Agency page verifies the 2021 UPC adoption at publication level; Board rule filing/effective dates and amendment text still need extraction. |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

The strongest verified date rules are the MUBEC amended-adoption date, the electrical installation commencement trigger, and the MUBEC population-based enforcement rule. MUBEC agency sources identify 2025-04-07 as the date the current amended adoption became enforceable/in effect. Electrical Chapter 120 uses an installation-commencement trigger: work commencing on or after 2024-07-01 must comply with the 2023 NEC as adopted. Maine statutes preserve separate local enforcement and population-based enforcement rules for MUBEC.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-me:mubec-2025-current | MUBEC building, residential, existing-building, mechanical, and energy components | effective / enforcement | 2025-04-07 | Current amended MUBEC adoption identified by state agency pages as enforceable or in effect beginning on this date. | unresolved | src:usa-me:fmo-building-codes; src:usa-me:energy-building-codes; src:usa-me:mubec-standards-amendments | partially_verified |
| date-rule:usa-me:mubec-local-threshold | MUBEC enforcement | population threshold | ongoing | Municipality has more than 4,000 residents; municipalities with 4,000 or fewer are not required to enforce but are constrained if they adopt/enforce a building code. | not applicable | src:usa-me:statute-10-9724; src:usa-me:statute-25-2373; src:usa-me:fmo-building-codes | partially_verified |
| date-rule:usa-me:electrical-2023-nec | Electrical installations | commencement-date trigger | 2024-07-01 | Installation commences on or after 2024-07-01. | no for covered commencing installations | src:usa-me:eeb-home; src:usa-me:eeb-ch120 | verified_at_rule_level |
| date-rule:usa-me:fuel-board-2023-rules | Fuel-gas / fuel-burning equipment standards | rule effective date | 2023-09-16 | Maine Fuel Board Chapter 6 rules in reviewed combined PDF. | unresolved | src:usa-me:fuel-board-rules | partially_verified |
| date-rule:usa-me:local-ordinance-preemption | Local ordinances inconsistent with MUBEC | preemption date / continuing rule | 2010-12-01 and ongoing | MUBEC replaced prior code frameworks and voided inconsistent political-subdivision ordinances, subject to preserved local enforcement ordinances and exceptions. | no for inconsistent local ordinances | src:usa-me:statute-10-9724 | partially_verified |
| date-rule:usa-me:accessibility-trigger | Covered public-building accessibility work | statutory applicability trigger | 2012-03-15 trigger referenced in statute | Application is completed or received, or construction begins, on or after the statutory trigger date and statutory thresholds/scope are met. | no for covered work | src:usa-me:statute-5-4594-g | partially_verified |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| MUBEC | none verified | null | null | null | null | null | monitor | src:usa-me:fmo-mubec-rules; src:usa-me:moca-code-rules; src:usa-me:energy-building-codes | No future statewide MUBEC date was verified beyond the 2025-04-07 current adoption. |
| Fire / life safety | unresolved | null | null | null | null | null | monitor | src:usa-me:fmo-rules; src:usa-me:fmo-nfpa-standards | Fire Marshal rule and NFPA-standard pages should be monitored for rulemaking and certified-rule updates. |
| Elevator / conveyance | unresolved | null | null | null | null | null | monitor | src:usa-me:elevator-program-laws-rules | Current/proposed elevator-code edition updates were not normalized in this pass. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-me:mubec-statewide-local-enforcement | Building / Residential / Existing / Mechanical / Energy | Constructed or renovated buildings subject to MUBEC | Statewide code applicability; enforcement obligation varies by municipal population/local adoption | MUBEC applies to constructed/renovated buildings statewide, but mandatory local enforcement applies to municipalities with more than 4,000 residents. | src:usa-me:statute-10-9724; src:usa-me:fmo-building-codes | partially_verified |
| applicability-rule:usa-me:mubec-statutory-exceptions | MUBEC | Listed building types | Statutory exceptions | Statute excludes specified categories from MUBEC requirements, including listed log/manufactured housing, post-and-beam/timber-frame construction, and crop warehouses/silos. | src:usa-me:statute-10-9724 | partially_verified |
| applicability-rule:usa-me:electrical-commencement | Electrical | Electrical installations | Commencement date | Installations commencing on or after 2024-07-01 must comply with the 2023 NEC as adopted by Chapter 120. | src:usa-me:eeb-home; src:usa-me:eeb-ch120 | verified_at_rule_level |
| applicability-rule:usa-me:barrier-free-public-buildings | Accessibility | Covered public buildings and alterations | Statutory coverage and State Fire Marshal review/permit path | Covered public-building projects require barrier-free/accessibility review and may require State Fire Marshal plan approval before local permit issuance. | src:usa-me:statute-5-4594-g; src:usa-me:fmo-barrier-free; src:usa-me:fmo-plans-review | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** statewide rule chapters and board-adopted amendments for MUBEC, plus separate specialty-board or agency amendment sets for electrical, fuel-gas, fire/life-safety, plumbing, accessibility, and elevator/conveyance programs.

**Where amendments are published:** MUBEC amendments are published through Technical Building Codes and Standards Board / MOCA / State Fire Marshal rule and guidance pages. Electrical amendments are in Electricians' Examining Board Chapter 120. Fuel-gas amendments are in Maine Fuel Board rules. Fire/life-safety rules are published by the State Fire Marshal, with a caveat that the agency page advises obtaining certified copies from MAPA for official court use.

**Amendment parsing status:** high_level_verified; section-by-section amendment extraction remains incomplete.

### 5.2 State Amendment Sources

| Amendment Source ID | Code Family / Scope | Publication Path | Parsed Level | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| amend:usa-me:mubec-rules | MUBEC building, residential, existing-building, energy, mechanical | MOCA / State Fire Marshal MUBEC rule chapters | chapter-level source identification | src:usa-me:fmo-mubec-rules; src:usa-me:moca-code-rules; src:usa-me:mubec-standards-amendments | partially_verified |
| amend:usa-me:electrical-ch120 | Electrical | Electricians' Examining Board Chapter 120 | rule-level adoption and trigger date | src:usa-me:eeb-ch120 | partially_verified |
| amend:usa-me:fuel-board-ch6 | Fuel-gas / fuel-burning equipment | Maine Fuel Board Chapter 6 | standard-adoption level | src:usa-me:fuel-board-rules | partially_verified |
| amend:usa-me:fire-marshal-rules | Fire/life safety | State Fire Marshal rule chapters and NFPA adopted-standards page | standard-list level | src:usa-me:fmo-rules; src:usa-me:fmo-nfpa-standards | partially_verified |
| amend:usa-me:plumbing-upc | Plumbing | Maine State Internal Plumbing Code / Plumbers' Examining Board adoption page | agency-publication level | src:usa-me:moca-plumbing | unresolved |
| amend:usa-me:accessibility | Accessibility | 5 M.R.S. § 4594-G and State Fire Marshal barrier-free program pages | statutory/program level | src:usa-me:statute-5-4594-g; src:usa-me:fmo-barrier-free | partially_verified |

### 5.3 High-Impact State Amendments

No section-by-section high-impact amendment set was extracted in this pass. Do not rely on this report for detailed design exceptions, climate-zone modifications, sprinkler triggers, local stretch-code details, or code-section text until the relevant C.M.R. chapters and copyrighted model-code amendment documents are parsed separately.

Verified high-level amendment/adoption features:

| Feature | Summary | Source IDs | Status |
| --- | --- | --- | --- |
| MUBEC optional stretch code | The Board must make an optional energy-conservation appendix or stretch code available for voluntary municipal adoption; the energy agency page describes a 15% UA improvement over base IECC performance for commercial and residential buildings. | src:usa-me:statute-10-9722; src:usa-me:energy-building-codes | partially_verified |
| Electrical NEC amendments/exclusions | Chapter 120 adopts the 2023 NEC subject to Maine-specific amendments and exclusions. | src:usa-me:eeb-home; src:usa-me:eeb-ch120 | partially_verified |
| Fuel-gas standards with amendments | Maine Fuel Board rules adopt NFPA 54-2021 and NFPA 58-2020 subject to Maine amendments. | src:usa-me:fuel-board-rules | partially_verified |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-me"
  model: "statewide_code_population_threshold_local_enforcement"
  enforcing_entities:
    - "Municipal building officials and local code enforcement officers in municipalities required to enforce MUBEC"
    - "Interlocal enforcement arrangements with certified building officials"
    - "Contractual county or regional enforcement arrangements"
    - "Certified third-party inspectors"
    - "Municipalities of 4,000 or fewer residents when they elect to adopt/enforce an allowed code"
  required_officials:
    - "Certified municipal building officials"
    - "Local code enforcement officers"
    - "Certified third-party inspectors, where used"
  state_reserved_activities:
    - "Technical Building Codes and Standards Board statewide MUBEC adoption/amendment/maintenance"
    - "State Fire Marshal / Department of Public Safety fire and life-safety authority"
    - "State Fire Marshal covered public-building, barrier-free, and other plan-review/permitting functions"
    - "Electricians' Examining Board electrical standards/licensing/permitting"
    - "Maine Fuel Board fuel-gas and fuel-burning equipment standards"
    - "Plumbers' Examining Board internal plumbing code authority"
    - "Elevator and Tramway Safety Program inspection and plan-submission framework"
  source_ids:
    - "src:usa-me:statute-10-9722"
    - "src:usa-me:statute-10-9723"
    - "src:usa-me:statute-10-9724"
    - "src:usa-me:statute-25-2373"
    - "src:usa-me:fmo-plans-review"
  verification_status: "partially_verified"
  confidence: 0.82

### 6.2 Local Amendment Rule

local_amendment_rule:
  rule_id: "local-amendment-rule:usa-me"
  model: "state_preemption_with_board_amendment_process_and_limited_local_exceptions"
  applies_to_code_families:
    - "MUBEC building"
    - "MUBEC residential"
    - "MUBEC existing-building"
    - "MUBEC energy"
    - "MUBEC mechanical"
  approval_required: true
  approving_authority_id: "ahj:usa-me:tbcands-board"
  filing_required: "unresolved"
  registry_exists: "partial"
  registry_source_ids:
    - "src:usa-me:statute-10-9722"
  legal_basis_source_ids:
    - "src:usa-me:statute-10-9722"
    - "src:usa-me:statute-10-9724"
  verification_status: "partially_verified"
  confidence: 0.70
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Maine local enforcement authority should not be read as general local amendment authority. A municipality may administer and enforce MUBEC through local officials or accepted alternatives, and smaller municipalities may choose whether to enforce an allowed code. That enforcement authority is constrained by statewide preemption: ordinances inconsistent with MUBEC are void unless a statutory exception applies. Local or regional governments may propose amendments through the statewide Board process.

### 6.4 Known Local Amendment Registries

| Registry ID | Scope | Registry Host | Source IDs | Status | Notes |
| --- | --- | --- | --- | --- | --- |
| registry:usa-me:stretch-code-municipalities | Municipal voluntary adoption of optional energy appendix/stretch code | Technical Building Codes and Standards Board / State | src:usa-me:statute-10-9722; src:usa-me:energy-building-codes | unresolved | Statute requires the Board to maintain a list of municipalities adopting the optional energy appendix; the actual registry/list was not captured in this pass. |

### 6.5 Municipality-Specific Known Amendments

No municipality-specific MUBEC amendments were parsed in this pass. Municipal ordinance research should be scoped separately and reconciled against the statewide preemption language and Board amendment process.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: partially_modeled

Jurisdiction stack:

```text
Address
  -> State of Maine
  -> County
  -> Municipality or unorganized territory
  -> Municipal population threshold and local MUBEC adoption/enforcement status
  -> Local building official / code enforcement officer / certified third-party inspector / county or regional enforcement agreement
  -> State Fire Marshal public-building, fire/life-safety, sprinkler, fuel-tank, amusement-ride, gas-station, or accessibility review if triggered
  -> Trade-specific AHJ for electrical, plumbing, fuel-gas/fuel-burning equipment, elevator/tramway, and other specialty systems
  -> Applicable statewide MUBEC adoption records
  -> Applicable optional stretch-code or local enforcement records, if verified
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | not selected | none | statewide | unknown | pending |
| County | not selected | none | statewide | unknown | pending |
| Municipality | not selected | none | statewide | unknown | pending |
| Municipal population threshold | not selected | none | statewide | unknown | pending |
| Fire District | not selected | none | statewide | unknown | pending |
| Special District | not selected | none | statewide | unknown | pending |

### 7.3 AHJ Contact Data

No AHJ contact data was populated for this pass. The next pass should collect municipal code-enforcement contacts for municipalities over 4,000 residents and map State Fire Marshal plan-review triggers by occupancy/project type.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Source Type | Title | Issuing Authority | URL / Citation | Date / Version | Key Facts Used | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-me:statute-10-9722 | statute | 10 M.R.S. § 9722, Technical Building Codes and Standards Board | Maine Legislature | https://www.mainelegislature.org/legis/statutes/10/title10sec9722.html | current HTML accessed 2026-06-25 | Establishes Board within MOCA; requires adoption/amendment/maintenance of MUBEC; lists MUBEC code families; amendment process; conflict-resolution role; optional energy appendix list. | official |
| src:usa-me:statute-10-9723 | statute | 10 M.R.S. § 9723, Training and certification program | Maine Legislature | https://www.mainelegislature.org/legis/statutes/10/title10sec9723.html | current HTML accessed 2026-06-25 | Board committee and MOCA training/certification program for code enforcement officials and third-party inspectors. | official |
| src:usa-me:statute-10-9724 | statute | 10 M.R.S. § 9724, Application | Maine Legislature | https://www.mainelegislature.org/legis/statutes/10/title10sec9724.html | current HTML accessed 2026-06-25 | Mandatory MUBEC enforcement in municipalities over 4,000; smaller municipality rule; preemption/voiding inconsistent ordinances; statutory exceptions. | official |
| src:usa-me:statute-25-2373 | statute | 25 M.R.S. § 2373, Municipal inspection options | Maine Legislature | https://www.mainelegislature.org/legis/statutes/25/title25sec2373.html | current HTML accessed 2026-06-25 | Local enforcement options: local officials, interlocal agreements, county/regional agreements, certified third-party inspectors. | official |
| src:usa-me:fmo-building-codes | agency_page | Building Codes | Maine Department of Public Safety, Office of State Fire Marshal | https://www.maine.gov/dps/fmo/building-codes | accessed 2026-06-25 | Current amended MUBEC enforceable April 7, 2025; current code list; statewide applicability and municipal enforcement threshold. | official |
| src:usa-me:fmo-mubec-rules | agency_page | MUBEC Rules and Laws | Maine Department of Public Safety, Office of State Fire Marshal | https://www.maine.gov/dps/fmo/building-codes/mubec-rules | accessed 2026-06-25 | MUBEC rule chapter list and current/prior rule structure. | official |
| src:usa-me:mubec-standards-amendments | agency_pdf | MUBEC Standards and Amendments | Maine Department of Public Safety / Office of State Fire Marshal | https://www.maine.gov/dps/fmo/sites/maine.gov.dps.fmo/files/inline-files/MUBEC%20Standards%20and%20Amendments.pdf | effective beginning 2025-04-07; accessed 2026-06-25 | Current model-code editions and standards; date and local enforcement summary. | official_courtesy_pdf |
| src:usa-me:moca-code-rules | agency_page | Code Enforcement Rules | Maine Office of Community Affairs | https://www.maine.gov/moca/programs/code-enforcement/laws-rules/rules | accessed 2026-06-25 | Current MUBEC rule chapter list, including amendments to 2021 IBC, IEBC, IRC, IECC, and IMC. | official |
| src:usa-me:energy-building-codes | agency_page | Building & Energy Codes | Maine Department of Energy Resources | https://www.maine.gov/energy/initiatives/energy-efficiency/building-codes | accessed 2026-06-25 | Current codes in effect as of 2025-04-07; 2021 I-code list; optional stretch-code description. | official |
| src:usa-me:statute-25-2396 | statute | 25 M.R.S. § 2396, Office of State Fire Marshal | Maine Legislature | https://www.mainelegislature.org/legis/statutes/25/title25sec2396.html | current HTML accessed 2026-06-25 | Establishes State Fire Marshal; statewide enforcement powers including NFPA 101; fire-prevention areas. | official |
| src:usa-me:statute-25-2452 | statute | 25 M.R.S. § 2452, Life safety and property protection | Maine Legislature | https://www.mainelegislature.org/legis/statutes/25/title25sec2452.html | current HTML accessed 2026-06-25 | Commissioner of Public Safety rulemaking for fire/life-safety around buildings and structures. | official |
| src:usa-me:fmo-rules | agency_page | Rules of the State Fire Marshal | Maine Department of Public Safety, Office of State Fire Marshal | https://www.maine.gov/dps/fmo/laws-rules-policy/rules | accessed 2026-06-25 | Fire Marshal rule list; Chapter 3 Fire Prevention Code / NFPA 1-2018; caveat about certified MAPA copies. | official_courtesy_html |
| src:usa-me:fmo-nfpa-standards | agency_page | State Adopted NFPA Standards | Maine Department of Public Safety, Office of State Fire Marshal | https://www11.maine.gov/dps/fmo/laws-rules-policy/nfpa-standards | accessed 2026-06-25 | Lists adopted NFPA standards including NFPA 1-2018, NFPA 54-2021, NFPA 58-2020, NFPA 101-2018. | official |
| src:usa-me:eeb-home | agency_page | Electricians' Examining Board | Maine Department of Professional and Financial Regulation | https://www.maine.gov/pfr/professionallicensing/professions/electricians | accessed 2026-06-25 | Board purpose; 2023 NEC adoption notice and 2024-07-01 commencement trigger. | official |
| src:usa-me:eeb-ch120 | agency_pdf | 02-318 C.M.R. ch. 120, Electrical Installation Standards | Maine Electricians' Examining Board | https://www.maine.gov/pfr/professionallicensing/sites/maine.gov.pfr.professionallicensing/files/inline-files/Electrician-Examining-Board-Chapter-120.pdf | repealed/replaced effective 2024-07-01 | Rule adoption of 2023 NEC subject to amendments/exclusions; statutory authority; commencement trigger. | official_pdf |
| src:usa-me:moca-plumbing | agency_page | Publications and Manuals: Maine State Internal Plumbing Code | Maine Office of Community Affairs | https://www.maine.gov/moca/programs/code-enforcement/resources/publications-manuals | accessed 2026-06-25 | Plumbers' Examining Board adoption of IAPMO 2021 UPC as Maine plumbing installation standard; amendments available with code. | official |
| src:usa-me:fuel-board-laws-rules | agency_page | Maine Fuel Board Laws and Rules | Maine Department of Professional and Financial Regulation | https://www.maine.gov/pfr/professionallicensing/professions/maine-fuel-board/home/laws-rules | accessed 2026-06-25 | Official entry point for Fuel Board laws/rules and inter-board MOUs. | official |
| src:usa-me:fuel-board-rules | agency_pdf | Maine Fuel Board Rules, including Chapter 6 Adoption of Standards | Maine Fuel Board | https://www.maine.gov/pfr/professionallicensing/sites/maine.gov.pfr.professionallicensing/files/inline-files/MFB%20FUE%20rules%20adopted%20May%2026%2C%202023%20combined%20PDF._0.pdf | effective 2023-09-16 for reviewed rule segment | Adopted fuel standards including NFPA 54-2021 and NFPA 58-2020 subject to amendments; rule effective date. | official_pdf |
| src:usa-me:statute-5-4594-g | statute | 5 M.R.S. § 4594-G, Standards of construction | Maine Legislature | https://www.mainelegislature.org/legis/statutes/5/title5sec4594-G.html | current HTML accessed 2026-06-25 | 2010 ADA Standards definition; applicability; barrier-free certification; State Fire Marshal plan-review and variance role. | official |
| src:usa-me:fmo-barrier-free | agency_page | Barrier Free Construction Permits | Maine Department of Public Safety, Office of State Fire Marshal | https://www.maine.gov/dps/fmo/inspections-plans-review/construction/barrier-free | accessed 2026-06-25 | State Fire Marshal barrier-free permit and plan-review program description; state permits in addition to local permits. | official |
| src:usa-me:fmo-plans-review | agency_page | Plans Review | Maine Department of Public Safety, Office of State Fire Marshal | https://www.maine.gov/dps/fmo/plans-review | accessed 2026-06-25 | Plans Review scope for public buildings, accessibility, sprinklers, fuel tanks, amusement rides, gas stations. | official |
| src:usa-me:elevator-program-laws-rules | agency_pdf | Elevator and Tramway Safety Program Laws and Rules | Maine Department of Professional and Financial Regulation / OPOR | https://www.maine.gov/pfr/professionallicensing/sites/maine.gov.pfr.professionallicensing/files/inline-files/Elevator-and-Tramway-Safety-Program-Laws-and-Rules.pdf | compiled PDF accessed 2026-06-25 | Identifies program rule framework, national-code chapter, plan submission, inspection, and certificate concepts; current editions unresolved. | official_pdf_needs_currency_check |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-me:fmo-rules | certified_rule_copy | The State Fire Marshal rules page warns that files on the page are not official state rules for use before a court and that certified copies should be obtained from MAPA. | Use for research and source discovery; obtain certified rule copies for verified legal production. |
| src:usa-me:mubec-standards-amendments | courtesy_pdf_numbering | The one-page MUBEC standards PDF is useful for effective-date and adopted-code summary, but its chapter numbering differs from the current MOCA/OSFM rule-page structure. | Use as summary support only; cite rule chapters for detailed amendment text. |
| src:usa-me:fmo-barrier-free | agency_page_threshold_conflict | The agency barrier-free page and statute appear to use different dollar thresholds in the excerpts reviewed. | Treat 5 M.R.S. § 4594-G as controlling until agency confirms current operational threshold. |
| src:usa-me:moca-plumbing | rule_date_gap | The agency page verifies 2021 UPC adoption but does not by itself provide normalized adoption/effective/mandatory dates. | Extract Plumbers' Examining Board rule filing and amendments before marking plumbing fully verified. |
| src:usa-me:elevator-program-laws-rules | currency_check_needed | The elevator/tramway PDF identifies the framework but exact current adopted national-code editions and any later rulemaking were not normalized. | Keep elevator code edition unresolved until current C.M.R. chapter and rulemaking records are extracted. |
| src:usa-me:fmo-nfpa-standards | standard_list_not_amendment_text | The adopted-standards page lists NFPA editions but does not provide full amendment text or all transition rules. | Use for edition identification; use certified rules for detailed requirements. |

### 8.3 Supplemental Sources

None used in this pass.

### 8.4 Source Extraction Metadata

| Extraction ID | Source IDs | Extracted By | Extracted On | Method | Notes |
| --- | --- | --- | --- | --- | --- |
| extract:usa-me:2026-06-25-core | src:usa-me:statute-10-9722; src:usa-me:statute-10-9723; src:usa-me:statute-10-9724; src:usa-me:statute-25-2373 | GPT-5.5 Thinking | 2026-06-25 | official statute HTML review | Core MUBEC authority and local enforcement fields populated. |
| extract:usa-me:2026-06-25-codes | src:usa-me:fmo-building-codes; src:usa-me:fmo-mubec-rules; src:usa-me:mubec-standards-amendments; src:usa-me:moca-code-rules; src:usa-me:energy-building-codes | GPT-5.5 Thinking | 2026-06-25 | agency page/PDF review | Current MUBEC code-family editions and 2025-04-07 date populated. |
| extract:usa-me:2026-06-25-specialty | src:usa-me:eeb-home; src:usa-me:eeb-ch120; src:usa-me:moca-plumbing; src:usa-me:fuel-board-laws-rules; src:usa-me:fuel-board-rules; src:usa-me:elevator-program-laws-rules | GPT-5.5 Thinking | 2026-06-25 | agency page/PDF review | Specialty trade rows populated to verified level; plumbing/elevator remain open. |
| extract:usa-me:2026-06-25-fire-access | src:usa-me:statute-25-2396; src:usa-me:statute-25-2452; src:usa-me:fmo-rules; src:usa-me:fmo-nfpa-standards; src:usa-me:statute-5-4594-g; src:usa-me:fmo-barrier-free; src:usa-me:fmo-plans-review | GPT-5.5 Thinking | 2026-06-25 | statute and agency page review | Fire/life-safety and accessibility authority captured at high level with caveats. |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 0.90 | none | Status upgraded only after source registry, source-ID resolution, and validation pass; not marked fully verified. |
| report | risk.overall_confidence | 0.62 | verified | 0.80 | none | Confidence reflects strong core MUBEC/electrical support and unresolved specialty details. |
| ahj:usa-me:tbcands-board | authority.name | Technical Building Codes and Standards Board | partially_verified | 0.85 | src:usa-me:statute-10-9722 | Primary MUBEC authority verified from statute. |
| ahj:usa-me:tbcands-board | role | Adopt, amend, and maintain MUBEC | partially_verified | 0.85 | src:usa-me:statute-10-9722 | Board conflict-resolution role with fire/life safety also captured. |
| record:usa-me:mubec-ibc-2021 | edition | 2021 IBC | partially_verified | 0.85 | src:usa-me:fmo-building-codes; src:usa-me:energy-building-codes; src:usa-me:mubec-standards-amendments | Current MUBEC edition supported by multiple official agency sources. |
| record:usa-me:mubec-ibc-2021 | effective_date | 2025-04-07 | partially_verified | 0.82 | src:usa-me:fmo-building-codes; src:usa-me:energy-building-codes; src:usa-me:mubec-standards-amendments | Exact certified rule-adoption filing date still unresolved. |
| local-enforcement:usa-me | model | statewide_code_population_threshold_local_enforcement | partially_verified | 0.82 | src:usa-me:statute-10-9724; src:usa-me:statute-25-2373; src:usa-me:fmo-building-codes | Population threshold and enforcement options supported by statute and agency page. |
| local-amendment-rule:usa-me | model | state_preemption_with_board_amendment_process_and_limited_local_exceptions | partially_verified | 0.70 | src:usa-me:statute-10-9722; src:usa-me:statute-10-9724 | Needs legal review before production automation. |
| record:usa-me:electrical-nec-2023 | mandatory_date | 2024-07-01 | verified_at_rule_level | 0.90 | src:usa-me:eeb-home; src:usa-me:eeb-ch120 | Commencement trigger directly stated in Board rule/notice. |
| record:usa-me:fire-nfpa1-2018 | edition | NFPA 1-2018 | partially_verified | 0.75 | src:usa-me:fmo-rules; src:usa-me:fmo-nfpa-standards | Certified rule text still needed for final legal production. |
| record:usa-me:plumbing-upc-2021 | edition | 2021 UPC | partially_verified | 0.65 | src:usa-me:moca-plumbing | Rule effective date and amendments unresolved. |
| record:usa-me:fuel-gas-nfpa54-2021 | edition | NFPA 54-2021 | partially_verified | 0.78 | src:usa-me:fuel-board-rules | Standard identified in Board rules; transition details unresolved. |
| record:usa-me:accessibility | standard | 2010 ADA Standards | partially_verified | 0.78 | src:usa-me:statute-5-4594-g; src:usa-me:fmo-barrier-free | Agency/statute threshold discrepancy recorded as caveat. |
| ahj:usa-me:elevator-tramway-program | edition | exact editions unresolved | unresolved | 0.35 | src:usa-me:elevator-program-laws-rules | Framework identified; current adopted national-code edition not normalized. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Every `src:usa-me:*` identifier cited in the body appears in Section 8. |
| All authority IDs resolve | pass | Named authority IDs used in Sections 2 and 6 are defined in authority tables or specialized-authority rows. |
| All current code families have adoption matrix rows | pass | Core and specialty rows are present; unresolved date fields are explicit. |
| Building and operational fire code are separated | pass | MUBEC building rows and Fire Marshal fire-prevention/life-safety rows are separate. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Separate columns are maintained; null is used where dates were not verified. |
| Effective dates are valid ISO dates | pass | Entered dates use YYYY-MM-DD. |
| No impossible date sequences | pass | No conflicting sequence was introduced; unresolved dates remain null. |
| Transition rules have explicit trigger conditions | fail | MUBEC and electrical triggers are captured, but plumbing, fire, elevator, and some fuel-gas transition triggers remain unresolved. |
| Permit-date logic is captured where applicable | fail | Accessibility and electrical triggers are captured; complete permit-date logic across all code families remains incomplete. |
| Local enforcement model classified | pass | Population-threshold local enforcement model is captured. |
| Local amendment rule classified | pass | Classified as state preemption with Board amendment process and limited local exceptions; legal review still required. |
| AHJ confirmation metadata present | fail | AHJ contact data and boundary datasets were not populated. |
| Official-source caveats captured | pass | Certified-copy, courtesy-PDF, threshold-conflict, and rule-date caveats are documented. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-me:001 | medium | fire/life-safety amendments | Fire Marshal NFPA edition list and statutory authority are captured, but certified C.M.R. rule text and detailed amendments/transition rules remain incomplete. | Extract certified MAPA rules for State Fire Marshal chapters and normalize amendment/effective dates. | null | null | open |
| issue:usa-me:002 | medium | plumbing rule dates | Agency page verifies 2021 UPC adoption, but Board rule filing, effective date, and Maine amendment text were not normalized. | Extract Plumbers' Examining Board rules and amendment publication path. | null | null | open |
| issue:usa-me:003 | medium | elevator/conveyance current editions | Elevator/tramway framework identified; exact current national-code editions and any later/proposed rulemaking not normalized. | Extract current C.M.R. Chapter 511 and rulemaking docket status. | null | null | open |
| issue:usa-me:004 | medium | local amendment scope | Preemption and Board amendment process are captured, but production automation needs legal review of local ordinance exceptions and optional stretch-code registry. | Review 10 M.R.S. §§ 9722 and 9724 with municipal home-rule/local ordinance materials; locate stretch-code municipality list. | null | null | open |
| issue:usa-me:005 | low | MUBEC amendment detail | Current MUBEC editions and effective date are captured; section-level Maine amendments are not parsed. | Parse MOCA/OSFM Chapters 4-8 and produce code-section-level amendment records. | null | null | open |
| issue:usa-me:006 | low | AHJ contacts and boundaries | No municipal AHJ contacts, population data source, or boundary dataset was selected. | Add municipal boundary source, population source, and code-enforcement contact dataset. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-me:mubec-rules | src:usa-me:fmo-mubec-rules | html_diff | monthly | MUBEC rule chapter restructuring, new chapters, or revised effective dates | 2026-06-25 | active |
| watch:usa-me:moca-code-rules | src:usa-me:moca-code-rules | html_diff | monthly | New MUBEC chapter documents or amendment updates | 2026-06-25 | active |
| watch:usa-me:energy-building-codes | src:usa-me:energy-building-codes | html_diff | monthly | New energy-code or stretch-code announcement | 2026-06-25 | active |
| watch:usa-me:fire-rules | src:usa-me:fmo-rules | html_diff | monthly | New State Fire Marshal rulemaking, chapter update, or certified-rule notice | 2026-06-25 | active |
| watch:usa-me:nfpa-standards | src:usa-me:fmo-nfpa-standards | html_diff | quarterly | Adopted NFPA edition update | 2026-06-25 | active |
| watch:usa-me:electrical | src:usa-me:eeb-home | html_diff | quarterly | NEC edition update or Chapter 120 notice | 2026-06-25 | active |
| watch:usa-me:plumbing | src:usa-me:moca-plumbing | html_diff | quarterly | UPC edition/amendment update | 2026-06-25 | active |
| watch:usa-me:fuel-board | src:usa-me:fuel-board-laws-rules | html_diff | quarterly | Fuel Board rule update or standard-edition change | 2026-06-25 | active |
| watch:usa-me:elevator | src:usa-me:elevator-program-laws-rules | pdf_diff | quarterly | Elevator/tramway adopted-code edition or rule update | 2026-06-25 | pending |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-24 | Baseline draft stub existed with unresolved authority/adoption/local enforcement fields. | state-report:usa-me | none | prior draft generator | Superseded by 2026-06-25 source-backed partial verification pass. |
| 2026-06-25 | Populated Maine report with official source registry, authority model, current MUBEC adoption, specialty-code rows, local enforcement model, date rules, QA checks, and open issues. | state-report:usa-me; ahj:usa-me:tbcands-board; local-enforcement:usa-me | src:usa-me:statute-10-9722; src:usa-me:statute-10-9724; src:usa-me:fmo-building-codes; src:usa-me:energy-building-codes; src:usa-me:eeb-ch120 | GPT-5.5 Thinking | Status set to partially_verified, not verified; unresolved specialty fields remain explicit. |
