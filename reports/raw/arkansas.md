---
state:
  state_id: "US-AR"
  name: "Arkansas"
  abbreviation: "AR"
report:
  report_id: "state-report:usa-ar"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-25"
  last_verified: "2026-06-25"
  reviewed_by: null
risk:
  overall_confidence: 0.58 # 0.00 - 1.00
  risk_flags:
    - "afpc_2021_clean_copy_found_but_final_filing_date_unresolved"
    - "trade_code_editions_unresolved"
    - "local_amendment_authority_unresolved"
  open_questions_count: 4

---

# State Building Code Authority Report: Arkansas

## 1. Executive Summary

- **Authority model:** Arkansas has a State Fire Marshal's Office within the Arkansas Department of Public Safety / Division of Emergency Management. The State Fire Marshal's Office states that it is responsible for enforcing, periodically revising, and updating the Arkansas Fire Prevention Code and for reviewing state building projects.

- **Statewide code status:** A State Fire Marshal clean-copy rules PDF identifies the **Arkansas Fire Prevention Code, 2021 edition** as composed of the 2021 International Fire Code, 2021 International Building Code, and 2021 International Residential Code, as amended and adopted by the Arkansas State Fire Marshal. However, the public page labels the linked PDF as a **proposed** fire-code rule, and the PDF states effectiveness as ten days after final filing without giving the final filing date. Treat the AFPC 2021 facts as partially verified and keep the effective date unresolved until final filing is confirmed.

- **Local enforcement model:** Partially verified for fire/building-safety support only. The State Fire Marshal page says the office reviews plans for state building projects, assists local jurisdictions with plan reviews when needed, conducts state-building inspections, and assists local jurisdictions with inspections when requested. A general local building-code enforcement chain was not fully resolved.

- **Local amendment posture:** Unresolved. This pass did not verify a statewide local-amendment procedure, approval requirement, filing requirement, or registry for local amendments.

- **Known transition periods or pending changes:** The AFPC 2021 clean-copy rules PDF states that the rules become effective ten days after final filing with the Secretary of State's Office, but the final filing date was not verified. The Code of Arkansas Rules portal states that the official online rules database became effective on 2025-01-01 and is updated weekly; that portal date is not a building-code adoption date.

- **Production readiness:** partial_narrow_use_only

### Key Findings

```yaml
---
key_findings:
- topic: Fire/building-safety authority
  finding: The State Fire Marshal's Office states that it enforces, revises, and updates
    the Arkansas Fire Prevention Code and performs state building plan reviews and
    inspections.
  confidence: 0.86
  source_ids:
  - src:usa-ar:dps-fire-marshal-page
  - src:usa-ar:fire-marshal-duties-statute-pdf
- topic: AFPC 2021 contents
  finding: A State Fire Marshal clean-copy rules PDF identifies the AFPC 2021 edition
    as using the 2021 IFC, 2021 IBC, and 2021 IRC, as amended and adopted by the Arkansas
    State Fire Marshal.
  confidence: 0.76
  source_ids:
  - src:usa-ar:afpc-2021-rules-clean-copy
- topic: AFPC effective date
  finding: The AFPC 2021 clean-copy rules PDF says the rules become effective ten
    days after final filing with the Secretary of State, but the final filing date
    was not verified.
  confidence: 0.5
  source_ids:
  - src:usa-ar:afpc-2021-rules-clean-copy
- topic: Plumbing / mechanical / fuel gas / energy / electrical cross-references
  finding: The AFPC 2021 clean-copy rules PDF defines model-code references to the
    Arkansas State Plumbing Code, mechanical code for Arkansas, Arkansas State Gas
    Code, Arkansas Energy Code, and electrical code for the State of Arkansas; editions
    and adopting authorities were not verified here.
  confidence: 0.55
  source_ids:
  - src:usa-ar:afpc-2021-rules-clean-copy
- topic: Local enforcement
  finding: State Fire Marshal plan-review/inspection roles and local assistance are
    verified; the general local building AHJ model remains unresolved.
  confidence: 0.55
  source_ids:
  - src:usa-ar:dps-fire-marshal-page
- topic: Local amendments
  finding: Local amendment authority and registry status remain unresolved.
  confidence: 0.1
  source_ids:
  - src:usa-ar:law-section
  - src:usa-ar:rules-portal
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-ar:state-fire-marshal-office |
| Authority name | Arkansas State Fire Marshal's Office, Arkansas Department of Public Safety / Division of Emergency Management |
| Authority type | state_fire_marshal_and_building_safety_code_authority |
| Legal basis | Arkansas Fire Prevention Code authority as described by the State Fire Marshal page and the State Fire Marshal duties statute excerpt; AFPC 2021 clean-copy rules cite Act 841 of 2023, A.C.A. § 20-22-1010, and A.C.A. § 12-13-105. |
| Role | Enforces, periodically revises, and updates the Arkansas Fire Prevention Code; reviews state building project plans; conducts state-building inspections; assists local jurisdictions with plan reviews and inspections when needed or requested. |
| Enforcement model | hybrid_state_fire_marshal_with_local_assistance_and_unresolved_local_AHJ_scope |
| Source IDs | src:usa-ar:dps-fire-marshal-page, src:usa-ar:fire-marshal-duties-statute-pdf, src:usa-ar:afpc-2021-rules-clean-copy |
| Verification status | partially_verified |

### 2.2 State Law and Rules Framework

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-ar:state-law-framework |
| Authority name | Arkansas State Legislature / Arkansas Law Section / Code of Arkansas Rules |
| Authority type | state_legislative_and_rulemaking_publication_framework |
| Legal basis | Arkansas State Legislature official law portal and Code of Arkansas Rules codification portal |
| Role | Publishes or points to Arkansas law resources and codifies general and permanent state agency rules. |
| Enforcement model | publication_framework_only |
| Source IDs | src:usa-ar:legislature-home, src:usa-ar:law-section, src:usa-ar:rules-portal |
| Verification status | verified_for_publication_framework_only |

### 2.3 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-ar:state-fire-marshal-office | Arkansas State Fire Marshal's Office | AFPC Volume II / IBC reference identified in AFPC 2021 clean-copy rules; final filing date unresolved. | AFPC 2021 clean-copy rules; State Fire Marshal page | src:usa-ar:afpc-2021-rules-clean-copy, src:usa-ar:dps-fire-marshal-page | partially_verified |
| Residential | ahj:usa-ar:state-fire-marshal-office | Arkansas State Fire Marshal's Office | AFPC Volume II / IRC reference identified in AFPC 2021 clean-copy rules; final filing date unresolved. | AFPC 2021 clean-copy rules; State Fire Marshal page | src:usa-ar:afpc-2021-rules-clean-copy, src:usa-ar:dps-fire-marshal-page | partially_verified |
| Existing Building / Rehabilitation | ahj:usa-ar:unknown | Unknown | No existing-building or rehabilitation code adoption was verified. | unresolved | src:usa-ar:rules-portal | unresolved |
| Mechanical | ahj:usa-ar:unknown-mechanical | Unresolved Arkansas mechanical-code authority | AFPC cross-reference says International Mechanical Code means the mechanical code for Arkansas; edition and authority unresolved. | AFPC 2021 clean-copy definitions | src:usa-ar:afpc-2021-rules-clean-copy | partially_verified_cross_reference_only |
| Plumbing | ahj:usa-ar:unknown-plumbing | Unresolved Arkansas plumbing-code authority | AFPC cross-reference says International Plumbing Code means the Arkansas State Plumbing Code; edition and authority unresolved. | AFPC 2021 clean-copy definitions | src:usa-ar:afpc-2021-rules-clean-copy | partially_verified_cross_reference_only |
| Fuel Gas | ahj:usa-ar:unknown-fuel-gas | Unresolved Arkansas gas-code authority | AFPC cross-reference says International Fuel Gas Code means the Arkansas State Gas Code; edition and authority unresolved. | AFPC 2021 clean-copy definitions | src:usa-ar:afpc-2021-rules-clean-copy | partially_verified_cross_reference_only |
| Electrical | ahj:usa-ar:unknown-electrical | Unresolved Arkansas electrical-code authority | AFPC cross-reference says National Electrical Code means the electrical code for the State of Arkansas; edition and authority unresolved. | AFPC 2021 clean-copy definitions | src:usa-ar:afpc-2021-rules-clean-copy | partially_verified_cross_reference_only |
| Energy | ahj:usa-ar:unknown-energy | Unresolved Arkansas energy-code authority | AFPC cross-reference says International Energy Conservation Code means the Arkansas Energy Code; edition and authority unresolved. | AFPC 2021 clean-copy definitions | src:usa-ar:afpc-2021-rules-clean-copy | partially_verified_cross_reference_only |
| Fire - construction references | ahj:usa-ar:state-fire-marshal-office | Arkansas State Fire Marshal's Office | AFPC Volume I / IFC reference identified in AFPC 2021 clean-copy rules. | AFPC 2021 clean-copy rules; State Fire Marshal page | src:usa-ar:afpc-2021-rules-clean-copy, src:usa-ar:dps-fire-marshal-page | partially_verified |
| Fire - operational / prevention code | ahj:usa-ar:state-fire-marshal-office | Arkansas State Fire Marshal's Office | Enforces, revises, and updates the Arkansas Fire Prevention Code. | State Fire Marshal page; duties statute excerpt | src:usa-ar:dps-fire-marshal-page, src:usa-ar:fire-marshal-duties-statute-pdf | verified_for_authority |
| Accessibility | ahj:usa-ar:unknown-accessibility | Unknown | No accessibility-code authority was verified. | unresolved | src:usa-ar:rules-portal | unresolved |
| Elevator / Conveyance | ahj:usa-ar:unknown-elevator | Unknown | No elevator or conveyance authority was verified. | unresolved | src:usa-ar:rules-portal | unresolved |

### 2.4 Authority Hierarchy Notes

Arkansas has a verified state-level fire/building-safety authority in the State Fire Marshal's Office. The State Fire Marshal page and duties-statute excerpt support authority over the Arkansas Fire Prevention Code. The AFPC 2021 clean-copy rules PDF supports a code-family structure in which the AFPC incorporates 2021 IFC, 2021 IBC, and 2021 IRC material, while cross-referencing separate Arkansas plumbing, mechanical, gas, energy, and electrical codes. The publication status of the AFPC 2021 clean copy requires follow-up because the hosting page labels the document as proposed and no final filing date was verified.

### 2.5 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-ar:001 | ahj:usa-ar:state-fire-marshal-office | enforces_revises_updates | Arkansas Fire Prevention Code | src:usa-ar:dps-fire-marshal-page, src:usa-ar:fire-marshal-duties-statute-pdf | verified_for_authority |
| edge:usa-ar:002 | ahj:usa-ar:state-fire-marshal-office | reviews_plans_for | state_building_projects | src:usa-ar:dps-fire-marshal-page | verified |
| edge:usa-ar:003 | ahj:usa-ar:state-fire-marshal-office | assists | local_jurisdictions_with_plan_reviews_and_inspections | src:usa-ar:dps-fire-marshal-page | verified |
| edge:usa-ar:004 | ahj:usa-ar:state-fire-marshal-office | adopts_or_amends | Arkansas Fire Prevention Code 2021 edition | src:usa-ar:afpc-2021-rules-clean-copy | partially_verified_final_filing_unresolved |
| edge:usa-ar:005 | ahj:usa-ar:state-law-framework | codifies | general_and_permanent_agency_rules | src:usa-ar:rules-portal | verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | Arkansas Fire Prevention Code, Volume II | International Building Code | 2021 | partially_verified_final_filing_unresolved | null | null | null | null | AFPC 2021 clean-copy rules say effective ten days after final filing with Secretary of State; final filing date unresolved. | src:usa-ar:afpc-2021-rules-clean-copy |
| Residential | Arkansas Fire Prevention Code, Volume II | International Residential Code | 2021 | partially_verified_final_filing_unresolved | null | null | null | null | AFPC 2021 clean-copy rules say effective ten days after final filing with Secretary of State; final filing date unresolved. | src:usa-ar:afpc-2021-rules-clean-copy |
| Existing Building / Rehabilitation | unknown | unknown | unknown | unresolved | null | null | null | null | unresolved | src:usa-ar:rules-portal |
| Mechanical | mechanical code for Arkansas | International Mechanical Code cross-reference | unknown | cross_reference_only | null | null | null | null | unresolved | src:usa-ar:afpc-2021-rules-clean-copy |
| Plumbing | Arkansas State Plumbing Code | International Plumbing Code cross-reference | unknown | cross_reference_only | null | null | null | null | unresolved | src:usa-ar:afpc-2021-rules-clean-copy |
| Fuel Gas | Arkansas State Gas Code | International Fuel Gas Code cross-reference | unknown | cross_reference_only | null | null | null | null | unresolved | src:usa-ar:afpc-2021-rules-clean-copy |
| Electrical | electrical code for the State of Arkansas | National Electrical Code cross-reference | unknown | cross_reference_only | null | null | null | null | unresolved | src:usa-ar:afpc-2021-rules-clean-copy |
| Energy | Arkansas Energy Code | International Energy Conservation Code cross-reference | unknown | cross_reference_only | null | null | null | null | unresolved | src:usa-ar:afpc-2021-rules-clean-copy |
| Fire - construction references | Arkansas Fire Prevention Code, Volume I | International Fire Code | 2021 | partially_verified_final_filing_unresolved | null | null | null | null | AFPC 2021 clean-copy rules say effective ten days after final filing with Secretary of State; final filing date unresolved. | src:usa-ar:afpc-2021-rules-clean-copy |
| Fire - operational / prevention code | Arkansas Fire Prevention Code, Volume I | International Fire Code | 2021 | partially_verified_final_filing_unresolved | null | null | null | null | AFPC 2021 clean-copy rules say effective ten days after final filing with Secretary of State; final filing date unresolved. | src:usa-ar:afpc-2021-rules-clean-copy, src:usa-ar:dps-fire-marshal-page |
| Accessibility | unknown | unknown | unknown | unresolved | null | null | null | null | unresolved | src:usa-ar:rules-portal |
| Elevator / Conveyance | unknown | unknown | unknown | unresolved | null | null | null | null | unresolved | src:usa-ar:rules-portal |

### 3.2 Normalized Adoption Records

| Adoption ID | Code Family | State Code Name | Base Code | Edition | Adopting Authority ID | Adoption Date | Effective Date | Operative Date | Mandatory Date | Source IDs | Verification Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| adoption:usa-ar:afpc-2021-ifc | Fire - operational / prevention code | Arkansas Fire Prevention Code, Volume I | International Fire Code | 2021 | ahj:usa-ar:state-fire-marshal-office | null | null | null | null | src:usa-ar:afpc-2021-rules-clean-copy | partially_verified_final_filing_unresolved |
| adoption:usa-ar:afpc-2021-ibc | Building | Arkansas Fire Prevention Code, Volume II | International Building Code | 2021 | ahj:usa-ar:state-fire-marshal-office | null | null | null | null | src:usa-ar:afpc-2021-rules-clean-copy | partially_verified_final_filing_unresolved |
| adoption:usa-ar:afpc-2021-irc | Residential | Arkansas Fire Prevention Code, Volume II | International Residential Code | 2021 | ahj:usa-ar:state-fire-marshal-office | null | null | null | null | src:usa-ar:afpc-2021-rules-clean-copy | partially_verified_final_filing_unresolved |
| adoption:usa-ar:plumbing-cross-reference | Plumbing | Arkansas State Plumbing Code | International Plumbing Code cross-reference | unknown | ahj:usa-ar:unknown-plumbing | null | null | null | null | src:usa-ar:afpc-2021-rules-clean-copy | cross_reference_only |
| adoption:usa-ar:fuel-gas-cross-reference | Fuel Gas | Arkansas State Gas Code | International Fuel Gas Code cross-reference | unknown | ahj:usa-ar:unknown-fuel-gas | null | null | null | null | src:usa-ar:afpc-2021-rules-clean-copy | cross_reference_only |
| adoption:usa-ar:energy-cross-reference | Energy | Arkansas Energy Code | International Energy Conservation Code cross-reference | unknown | ahj:usa-ar:unknown-energy | null | null | null | null | src:usa-ar:afpc-2021-rules-clean-copy | cross_reference_only |
| adoption:usa-ar:electrical-cross-reference | Electrical | electrical code for the State of Arkansas | National Electrical Code cross-reference | unknown | ahj:usa-ar:unknown-electrical | null | null | null | null | src:usa-ar:afpc-2021-rules-clean-copy | cross_reference_only |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

The only code-specific effective-date rule verified here is conditional: the AFPC 2021 clean-copy rules state that the rules become effective ten days after final filing with the Secretary of State's Office. The final filing date was not verified, so no ISO effective date should be entered. The Code of Arkansas Rules portal's 2025-01-01 effective date applies to the official rules database itself and is not a code-adoption effective date.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-ar:afpc-2021-effective | Arkansas Fire Prevention Code 2021 clean-copy rules | conditional_effective_date | ten_days_after_final_filing | Final filing with the Secretary of State's Office | unknown | src:usa-ar:afpc-2021-rules-clean-copy | partially_verified_final_filing_date_unresolved |
| date-rule:usa-ar:rules-portal-effective | Code of Arkansas Rules portal | portal_effective_date | 2025-01-01 | Official online rules database became effective | not_applicable | src:usa-ar:rules-portal | verified_for_portal_only |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Fire / Building / Residential | Arkansas Fire Prevention Code 2021 edition | null | null | null | null | null | monitor_final_filing | src:usa-ar:dps-fire-marshal-page, src:usa-ar:afpc-2021-rules-clean-copy | DPS page labels the clean-copy PDF as proposed rules; final filing date and current codified status need confirmation. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-ar:state-building-plan-review | Fire / Building safety | State building projects | State building project plan review | The State Fire Marshal's Office reviews plans for state building projects. | src:usa-ar:dps-fire-marshal-page | verified |
| applicability-rule:usa-ar:local-assistance | Fire / Building safety | Local jurisdiction plan reviews and inspections | Assistance requested or needed | The State Fire Marshal's Office assists local jurisdictions with plan reviews when needed and with inspections when requested. | src:usa-ar:dps-fire-marshal-page | verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** partially_verified_for_AFPC_only

**Where amendments are published:** AFPC 2021 clean-copy rules PDF links to ICC-hosted AFPC volumes and states that Arkansas revisions are indicated by dotted margin lines and Arkansas deletions by solid stars. Final codified publication path remains unresolved.

**Amendment parsing status:** not_started

### 5.2 State Amendment Sources

| Source ID | Amendment Set | Format | Status | Notes |
| --- | --- | --- | --- | --- |
| src:usa-ar:afpc-2021-rules-clean-copy | Arkansas Fire Prevention Code 2021 clean-copy rules and linked AFPC volumes | PDF with ICC volume links | partially_verified_source_status_unresolved | Public page labels the clean-copy as proposed rules; final filing and codified status need confirmation. |

### 5.3 High-Impact State Amendments

No high-impact Arkansas-specific amendment was extracted or validated in this pass.

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-ar"
  model: "hybrid_state_fire_marshal_with_unresolved_local_AHJ_scope"
  enforcing_entities:
    - "Arkansas State Fire Marshal's Office for Arkansas Fire Prevention Code authority, state building project plan review, and state-building inspections"
    - "local jurisdictions where they conduct plan reviews or inspections; scope unresolved"
  required_officials:
    - "building official is defined in AFPC 2021 clean-copy rules as any governmental official having authority to enforce that aspect of the code"
  state_reserved_activities:
    - "state building project plan review"
    - "state-building inspections"
  source_ids:
    - "src:usa-ar:dps-fire-marshal-page"
    - "src:usa-ar:afpc-2021-rules-clean-copy"
  verification_status: "partially_verified"
  confidence: 0.55
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-ar"
  model: "unresolved"
  applies_to_code_families:
    - "unresolved"
  approval_required: null
  approving_authority_id: null
  filing_required: null
  registry_exists: null
  registry_source_ids: []
  legal_basis_source_ids: []
  verification_status: "unresolved"
  confidence: 0.10
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement and local amendment authority are separate unresolved questions. This pass verifies that the State Fire Marshal's Office has statewide AFPC responsibilities and assists local jurisdictions with plan reviews and inspections. It does not verify whether local governments may amend the AFPC, whether amendments require state approval, or whether amendments must be filed in a state registry.

### 6.4 Known Local Amendment Registries

No statewide local amendment registry was verified.

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendments were parsed in this pass.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: not_started

Jurisdiction stack:

```text
Address
  -> State of Arkansas
  -> County
  -> Municipality / unincorporated county
  -> State Fire Marshal authority where state building, AFPC, or state fire-safety review applies
  -> Local building official / fire official, where local jurisdiction has enforcement authority
  -> Trade-specific AHJs for plumbing, mechanical, fuel gas, energy, electrical, accessibility, and elevator/conveyance matters
  -> Applicable state code adoption records
  -> Applicable local amendment records, if any
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- | --- |
| State | Arkansas Legislature and official rules portals | src:usa-ar:legislature-home, src:usa-ar:law-section, src:usa-ar:rules-portal | statewide | when state law or rules change | verified_for_publication_framework |
| County | not selected | none | statewide | unknown | pending |
| Municipality | not selected | none | statewide | unknown | pending |
| Fire District / Fire Department | Arkansas State Fire Marshal / Fire Services infrastructure | src:usa-ar:dps-fire-marshal-page | statewide state-fire-marshal support | unknown | partial |
| Special District | not selected | none | statewide | unknown | pending |

### 7.3 AHJ Contact Data

No AHJ contact dataset was populated for this pass.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Source Type | Publisher | URL | Accessed Date | Snapshot ID | Checksum | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-ar:legislature-home | Arkansas State Legislature home page | agency_page | Arkansas State Legislature | https://www.arkleg.state.ar.us/ | 2026-06-25 | snapshot-pending | snapshot-pending | verified |
| src:usa-ar:law-section | Arkansas Law Section | law_portal | Arkansas State Legislature | https://www.arkleg.state.ar.us/ArkansasLaw | 2026-06-25 | snapshot-pending | snapshot-pending | verified |
| src:usa-ar:rules-portal | Code of Arkansas Rules home page | rules_portal | Arkansas State Government | https://codeofarrules.arkansas.gov/ | 2026-06-25 | snapshot-pending | snapshot-pending | verified |
| src:usa-ar:dps-fire-marshal-page | State Fire Marshal's Office | agency_page | Arkansas Department of Public Safety | https://dps.arkansas.gov/emergency-management/adem/state-fire-marshals-office/ | 2026-06-25 | snapshot-pending | snapshot-pending | verified |
| src:usa-ar:afpc-2021-rules-clean-copy | Proposed Fire Code Rules – Clean Copy / Arkansas Fire Prevention Code Rules 2021 Edition | proposed_rule_pdf | Arkansas Department of Public Safety / Division of Emergency Management | https://dps.arkansas.gov/wp-content/uploads/ASP-Foreword-Fire-Code-Rules-Clean_10-2024.pdf | 2026-06-25 | snapshot-pending | snapshot-pending | partial |
| src:usa-ar:fire-marshal-duties-statute-pdf | Duties of State Fire Marshal Enforcement Section, A.C.A. § 12-13-105 excerpt | courtesy_statute_pdf | Arkansas Department of Public Safety | https://dps.arkansas.gov/wp-content/uploads/2020/05/state_fire_marshal_Duties_of_the_Fire_Marshal_Statute_12132018.pdf | 2026-06-25 | snapshot-pending | snapshot-pending | partial |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-ar:rules-portal | codified_rules_only | The Code of Arkansas Rules is the official codification of general and permanent agency rules, became effective as an online official database on 2025-01-01, and is updated weekly; it may not reflect recent rulemaking. | authoritative_for_codified_rules_status_after_rule_lookup |
| src:usa-ar:dps-fire-marshal-page | agency_page_summary | The State Fire Marshal page provides official agency descriptions and links, but does not itself list the final filing date for AFPC 2021. | authoritative_for_agency_role_and_link_discovery |
| src:usa-ar:afpc-2021-rules-clean-copy | proposed_rule_or_clean_copy | The public page labels this document as proposed rules. It supports the content of the AFPC 2021 clean-copy text, but final filing, codification, and effective date must be separately verified before marking current adoption verified. | use_for_partial_AFPC_2021_extraction_only |
| src:usa-ar:fire-marshal-duties-statute-pdf | courtesy_statute_excerpt | This is a DPS-hosted statute excerpt rather than the primary legislative code text. | verify_against_law_portal_before_final |

### 8.3 Supplemental Sources

None used in this pass.

### 8.4 Source Extraction Metadata

| Source ID | Parser | Parser Version | Extracted At | Extraction Confidence | Requires OCR | Requires Browser Automation | Manual Review Required |
| --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-ar:legislature-home | browser_manual | 0.1 | 2026-06-25T00:00:00Z | 0.99 | no | yes | yes |
| src:usa-ar:law-section | browser_manual | 0.1 | 2026-06-25T00:00:00Z | 0.99 | no | yes | yes |
| src:usa-ar:rules-portal | browser_manual | 0.1 | 2026-06-25T00:00:00Z | 0.99 | no | yes | yes |
| src:usa-ar:dps-fire-marshal-page | browser_manual | 0.1 | 2026-06-25T00:00:00Z | 0.98 | no | yes | yes |
| src:usa-ar:afpc-2021-rules-clean-copy | pdf_text_and_visual_spot_check | 0.1 | 2026-06-25T00:00:00Z | 0.86 | no | yes | yes |
| src:usa-ar:fire-marshal-duties-statute-pdf | pdf_text | 0.1 | 2026-06-25T00:00:00Z | 0.72 | no | yes | yes |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| ahj:usa-ar:state-fire-marshal-office | role | enforces, periodically revises, and updates AFPC | verified | 0.86 | src:usa-ar:dps-fire-marshal-page, src:usa-ar:fire-marshal-duties-statute-pdf | Statute excerpt should be checked against primary code text before final verification. |
| adoption:usa-ar:afpc-2021-ifc | base_model_code | International Fire Code 2021 | partially_verified | 0.76 | src:usa-ar:afpc-2021-rules-clean-copy | Source is a clean-copy PDF linked as proposed rules. |
| adoption:usa-ar:afpc-2021-ibc | base_model_code | International Building Code 2021 | partially_verified | 0.76 | src:usa-ar:afpc-2021-rules-clean-copy | Source is a clean-copy PDF linked as proposed rules. |
| adoption:usa-ar:afpc-2021-irc | base_model_code | International Residential Code 2021 | partially_verified | 0.76 | src:usa-ar:afpc-2021-rules-clean-copy | Source is a clean-copy PDF linked as proposed rules. |
| date-rule:usa-ar:afpc-2021-effective | trigger | ten days after final filing | partially_verified | 0.50 | src:usa-ar:afpc-2021-rules-clean-copy | Final filing date unresolved; no ISO effective date entered. |
| local-enforcement:usa-ar | model | hybrid_state_fire_marshal_with_unresolved_local_AHJ_scope | partially_verified | 0.55 | src:usa-ar:dps-fire-marshal-page, src:usa-ar:afpc-2021-rules-clean-copy | State and assistance roles verified; local AHJ scope unresolved. |
| local-amendment-rule:usa-ar | model | unresolved | unresolved | 0.10 | src:usa-ar:law-section, src:usa-ar:rules-portal | No statewide amendment rule extracted. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Every source ID cited in the body appears in section 8. |
| All authority IDs resolve | pass | Authority IDs used for verified or partially verified authorities appear in sections 2.1 through 2.3. |
| All current code families have adoption rows | pass | Adoption matrix rows are present; unresolved rows remain explicit. |
| Building and operational fire code are separated | pass | AFPC Volume I / IFC and AFPC Volume II / IBC/IRC are separated. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Dates are separate; unresolved code dates are null. |
| Effective dates are valid ISO dates | pass | The only ISO date entered in date rules is 2025-01-01 for the rules portal. |
| No impossible date sequences | pass | No adoption/effective/operative/mandatory sequence was invented. |
| Transition rules have explicit trigger conditions | pass | AFPC effective rule uses the final-filing trigger; final filing date unresolved. |
| Permit-date logic is captured where applicable | fail | No permit-date or grace-period rule was verified. |
| Local enforcement model classified | fail | State Fire Marshal role is partially classified; local AHJ scope unresolved. |
| Local amendment rule classified | fail | Local amendment rule remains unresolved. |
| AHJ confirmation metadata present | fail | No AHJ contacts entered. |
| Official-source caveats captured | pass | Proposed-rule, courtesy-statute, agency-page, and rules-portal caveats are captured. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-ar:001 | high | AFPC 2021 final status | The AFPC 2021 clean-copy rules identify 2021 IFC/IBC/IRC components, but the page labels the PDF as proposed and no final filing date was verified. | Confirm Secretary of State final filing date and current Code of Arkansas Rules codified status. | null | null | open |
| issue:usa-ar:002 | high | trade-code editions | AFPC definitions cross-reference Arkansas plumbing, mechanical, gas, energy, and electrical codes, but editions and adopting authorities were not verified. | Extract current rules from the Code of Arkansas Rules and relevant agency pages. | null | null | open |
| issue:usa-ar:003 | medium | local enforcement | State Fire Marshal state-building and assistance roles are verified, but local building/fire AHJ authority remains unresolved. | Review municipal/county enabling statutes and AFPC administration provisions. | null | null | open |
| issue:usa-ar:004 | medium | local amendments | Local amendment permission, approval, filing, and registry requirements remain unresolved. | Review AFPC administration provisions and local-code amendment statutes or rules. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-ar:legislature-home | src:usa-ar:legislature-home | html_diff | monthly | official legislature portal changes | 2026-06-25 | active |
| watch:usa-ar:law-section | src:usa-ar:law-section | html_diff | monthly | law portal navigation or content changes | 2026-06-25 | active |
| watch:usa-ar:rules-portal | src:usa-ar:rules-portal | html_diff | weekly | rule codification or update changes | 2026-06-25 | active |
| watch:usa-ar:dps-fire-marshal-page | src:usa-ar:dps-fire-marshal-page | html_diff | weekly | AFPC rule links, final-rule notices, or State Fire Marshal role text changes | 2026-06-25 | active |
| watch:usa-ar:afpc-2021-rules-clean-copy | src:usa-ar:afpc-2021-rules-clean-copy | pdf_diff | weekly | clean-copy PDF replaced, final filing date added, or proposed label removed | 2026-06-25 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-25 | Updated Arkansas report from broad unresolved stub to partially verified AFPC / State Fire Marshal authority report | report:usa-ar, ahj:usa-ar:state-fire-marshal-office, adoption:usa-ar:afpc-2021-ifc, adoption:usa-ar:afpc-2021-ibc, adoption:usa-ar:afpc-2021-irc | src:usa-ar:dps-fire-marshal-page, src:usa-ar:afpc-2021-rules-clean-copy, src:usa-ar:fire-marshal-duties-statute-pdf, src:usa-ar:rules-portal | ChatGPT | AFPC 2021 source-status and final filing date remain unresolved. |
