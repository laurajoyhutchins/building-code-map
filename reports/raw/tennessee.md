---
state:
  state_id: "US-TN"
  name: "Tennessee"
  abbreviation: "TN"
report:
  report_id: "state-report:usa-tn"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-26"
  last_verified: null
  reviewed_by: null
risk:
  overall_confidence: 0.67 # 0.00 - 1.00
  risk_flags:
    - "official_rule_pdfs_not_fulltext_extracted_in_this_pass"
    - "statutory_code_text_partly_supported_by_unofficial_mirrors"
    - "local_amendment_registry_not_identified"
    - "operational_fire_scope_requires_additional_title_68_review"
    - "elevator_conveyance_authority_not_checked"
  open_questions_count: 6

---

# State Building Code Authority Report: Tennessee

## 1. Executive Summary

- **Authority model:** Tennessee uses a state-minimum model centered on the Tennessee Department of Commerce and Insurance, Division of Fire Prevention / State Fire Marshal's Office (SFMO). The SFMO promulgates minimum statewide building construction safety standards and the Codes Enforcement section enforces statewide adopted fire and building construction safety codes. Local governments may adopt and enforce qualifying local building construction and fire safety codes, but the state framework retains minimum-code, audit, conflict-resolution, and certain reserved-review roles.

- **Statewide code status:** Tennessee's SFMO currently lists the 2021 IBC, 2021 IFC, 2021 IFGC, 2021 IMC, 2021 IPC, 2021 IPMC, 2021 commercial IECC, 2021 IEBC, limited 2021 NFPA 101 Life Safety Code coverage for specified new small residential board-and-care facilities, the 2018 IRC, the 2018 IECC for one- and two-family dwellings and townhomes, and the 2017 NEC with stated exceptions. The 2021 commercial/fire code update became effective 2025-04-17, with a 120-day previous-code submission transition for SFMO plan-approval projects through 2025-08-15.

- **Local enforcement model:** Hybrid. The SFMO enforces the statewide adopted fire and building codes and acts as the fallback for buildings not covered by an adopting/enforcing local jurisdiction. Local jurisdictions can certify that they have adopted and enforce qualifying code families; their enforcement is subject to SFMO audit and state conflict-resolution rules. Counties and municipalities may also opt out of statewide standards for one- and two-family dwellings, while owners in opt-out or non-code jurisdictions may request SFMO inspection.

- **Local amendment posture:** Local standards may not be less stringent than the state minimum unless permitted by statute or approved in writing by the Division/SFMO. Local jurisdictions adopting and enforcing codes must forward relevant ordinances to the Division within 60 days, keep adopted building/fire code editions within seven years of the most current published edition, and obtain Division review/approval before reducing locally adopted requirements.

- **Known transition periods or pending changes:** The 2025 rule filing became effective 2025-04-17. For projects requiring SFMO plan approval, previous codes could be used for plan submissions for 120 days after that date, through 2025-08-15. No later statewide code-change action was verified in this pass.

- **Production readiness:** partially_ready_for_narrow_validation

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: The SFMO is the primary statewide code authority for minimum building construction
    safety standards.
  confidence: 0.85
  source_ids:
  - src:usa-tn:tdci-codes-enforcement
  - src:usa-tn:lexis-tca-68-120-101
  - src:usa-tn:justia-tca-68-120-101
- topic: Primary commercial building code edition
  finding: 2021 International Building Code, with Tennessee amendments.
  confidence: 0.9
  source_ids:
  - src:usa-tn:tdci-currently-adopted-codes
  - src:usa-tn:rule-0780-02-02
- topic: Residential code edition
  finding: 2018 IRC and 2018 IECC for one- and two-family dwellings/townhomes, with
    Tennessee amendments.
  confidence: 0.85
  source_ids:
  - src:usa-tn:tdci-currently-adopted-codes
  - src:usa-tn:tdci-residential-permits
  - src:usa-tn:rule-0780-02-23
- topic: Electrical code edition
  finding: 2017 NEC, with SFMO-listed exceptions; rule effective/operative dates require
    careful treatment.
  confidence: 0.75
  source_ids:
  - src:usa-tn:tdci-currently-adopted-codes
  - src:usa-tn:rule-0780-02-01
  - src:usa-tn:cornell-0780-02-01-02
- topic: Fire code authority and edition
  finding: 2021 IFC is listed as currently adopted; NFPA 101 is limited to specified
    new small residential board-and-care facilities in the current SFMO list.
  confidence: 0.8
  source_ids:
  - src:usa-tn:tdci-currently-adopted-codes
  - src:usa-tn:rule-0780-02-02
- topic: Local enforcement
  finding: Local jurisdictions may adopt/enforce qualifying codes; otherwise SFMO
    enforces statewide codes for uncovered buildings. SFMO audits local enforcement
    at least every three years.
  confidence: 0.75
  source_ids:
  - src:usa-tn:justia-tca-68-120-101
  - src:usa-tn:cornell-0780-02-02-05
- topic: Local amendments
  finding: Local ordinances must not be less stringent unless approved; ordinances
    must be forwarded within 60 days; less-stringent reductions require Division review/approval.
  confidence: 0.78
  source_ids:
  - src:usa-tn:rule-0780-02-02
  - src:usa-tn:cornell-0780-02-02-05
- topic: Transition rule
  finding: 2025 update effective 2025-04-17; prior codes allowed for SFMO plan submissions
    through 2025-08-15.
  confidence: 0.9
  source_ids:
  - src:usa-tn:tdci-codes-enforcement
  - src:usa-tn:rule-filing-2025-01-17
  - src:usa-tn:rule-0780-02-03
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-tn:state-fire-marshal |
| Authority name | Tennessee Department of Commerce and Insurance, Division of Fire Prevention / State Fire Marshal's Office |
| Authority type | state agency / state fire marshal |
| Legal basis | T.C.A. § 68-120-101; T.C.A. § 68-120-106; Tenn. Comp. R. & Regs. 0780-02-02; Tenn. Comp. R. & Regs. 0780-02-03 |
| Role | Promulgates minimum statewide building construction safety standards; adopts model codes by rule; enforces statewide adopted fire and building construction safety codes; reviews plans and inspections for covered projects. |
| Enforcement model | hybrid_state_minimum_with_local_certification_and_state_reserved_scope |
| Source IDs | src:usa-tn:tdci-codes-enforcement; src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:rule-0780-02-02; src:usa-tn:rule-0780-02-03; src:usa-tn:lexis-tca-68-120-101; src:usa-tn:justia-tca-68-120-101 |
| Verification status | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-tn:state-fire-marshal | Tennessee State Fire Marshal's Office | Adopts/enforces minimum statewide building construction safety standards; current IBC adoption. | T.C.A. § 68-120-101; Tenn. Comp. R. & Regs. 0780-02-02 | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:rule-0780-02-02; src:usa-tn:justia-tca-68-120-101 | verified_core |
| Residential | ahj:usa-tn:state-fire-marshal | Tennessee State Fire Marshal's Office | Administers statewide one- and two-family dwelling and townhouse standards where applicable; recognizes opt-out/non-code jurisdictions and local enforcement. | T.C.A. § 68-120-101; Tenn. Comp. R. & Regs. 0780-02-23 | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:tdci-residential-permits; src:usa-tn:tdci-opt-out; src:usa-tn:rule-0780-02-23 | partially_verified |
| Existing Building / Rehabilitation | ahj:usa-tn:state-fire-marshal | Tennessee State Fire Marshal's Office | Adopts IEBC and state amendments for existing buildings. | Tenn. Comp. R. & Regs. 0780-02-02 | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:rule-0780-02-02 | verified_core |
| Mechanical | ahj:usa-tn:state-fire-marshal | Tennessee State Fire Marshal's Office | Adopts IMC; mechanical-inspector certification referenced in statute. | Tenn. Comp. R. & Regs. 0780-02-02; T.C.A. § 68-120-101 | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:rule-0780-02-02; src:usa-tn:justia-tca-68-120-101 | verified_core |
| Plumbing | ahj:usa-tn:state-fire-marshal | Tennessee State Fire Marshal's Office | Adopts IPC; plumbing-inspector certification referenced in statute. | Tenn. Comp. R. & Regs. 0780-02-02; T.C.A. § 68-120-101 | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:rule-0780-02-02; src:usa-tn:justia-tca-68-120-101 | verified_core |
| Fuel Gas | ahj:usa-tn:state-fire-marshal | Tennessee State Fire Marshal's Office | Adopts IFGC. | Tenn. Comp. R. & Regs. 0780-02-02 | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:rule-0780-02-02 | verified_core |
| Electrical | ahj:usa-tn:state-fire-marshal | Tennessee State Fire Marshal's Office / Department of Commerce and Insurance | Adopts NEC through Electrical Installations rules; NEC prevails for residential electrical chapters deleted from IRC. | Tenn. Comp. R. & Regs. 0780-02-01; Tenn. Comp. R. & Regs. 0780-02-23 | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:rule-0780-02-01; src:usa-tn:cornell-0780-02-01-02; src:usa-tn:rule-0780-02-23 | partially_verified |
| Energy - commercial | ahj:usa-tn:state-fire-marshal | Tennessee State Fire Marshal's Office | Adopts commercial IECC with Tennessee amendments. | Tenn. Comp. R. & Regs. 0780-02-02 | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:rule-0780-02-02 | verified_core |
| Energy - residential | ahj:usa-tn:state-fire-marshal | Tennessee State Fire Marshal's Office | Applies 2018 IECC / IRC energy compliance for one- and two-family dwellings and townhomes. | Tenn. Comp. R. & Regs. 0780-02-23 | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:tdci-residential-permits; src:usa-tn:rule-0780-02-23 | partially_verified |
| Fire - construction references | ahj:usa-tn:state-fire-marshal | Tennessee State Fire Marshal's Office | Adopts IFC construction-related provisions with IBC coordination and amendments. | Tenn. Comp. R. & Regs. 0780-02-02 | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:rule-0780-02-02 | verified_core |
| Fire - operational / prevention code | ahj:usa-tn:state-fire-marshal | Tennessee State Fire Marshal's Office | Adopts/enforces statewide fire prevention and fire protection standards; current list includes IFC and limited NFPA 101 coverage. Operational occupancy-inspection scope remains partially parsed. | T.C.A. §§ 68-102-113, 68-120-101; Tenn. Comp. R. & Regs. 0780-02-02 | src:usa-tn:tdci-codes-enforcement; src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:rule-0780-02-02 | partially_verified |
| Accessibility | ahj:usa-tn:state-fire-marshal | Tennessee State Fire Marshal's Office / Tennessee Public Buildings Accessibility Act framework | IBC Chapter 11 is removed from the Tennessee IBC adoption; accessibility conflicts are governed by the Tennessee Public Buildings Accessibility Act. Full accessibility statute parsing remains open. | T.C.A. §§ 68-120-201 to 68-120-205; Tenn. Comp. R. & Regs. 0780-02-02 | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:tdci-accessibility-act; src:usa-tn:cornell-0780-02-02-01 | partially_verified |
| Elevator / Conveyance | ahj:usa-tn:unknown | unresolved | Elevator/conveyance authority was not checked in this pass. | unresolved | none | unresolved |

### 2.3 Authority Hierarchy Notes

The state-level framework is not pure statewide preemption and not pure local adoption. The SFMO adopts minimum statewide standards. Local governments can adopt and enforce qualifying local codes for all buildings, for buildings other than one- and two-family dwellings, or for one- and two-family dwellings only. Where a local jurisdiction chooses only a partial scope, the SFMO enforces statewide codes for buildings outside that local scope. State buildings, educational occupancies, and other occupancies requiring SFMO inspection for initial licensure remain state-sensitive categories even in local jurisdictions. The SFMO also audits local enforcement and supersedes conflicting local interpretations for the same or substantially identical standards.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-tn:001 | ahj:usa-tn:state-fire-marshal | promulgates | minimum statewide building construction safety standards | src:usa-tn:lexis-tca-68-120-101; src:usa-tn:justia-tca-68-120-101 | partially_verified |
| edge:usa-tn:002 | ahj:usa-tn:state-fire-marshal | adopts_by_reference | IBC, IFC, IFGC, IMC, IPC, IPMC, IECC, IEBC, NFPA 101, NEC, IRC as listed in current rules/pages | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:rule-0780-02-02; src:usa-tn:rule-0780-02-23; src:usa-tn:rule-0780-02-01 | verified_core |
| edge:usa-tn:003 | local_jurisdictions | may_adopt_and_enforce | qualifying building construction and fire safety codes | src:usa-tn:justia-tca-68-120-101; src:usa-tn:cornell-0780-02-02-05 | partially_verified |
| edge:usa-tn:004 | ahj:usa-tn:state-fire-marshal | fallback_enforces | buildings not covered by a local jurisdiction's adopted/enforced scope | src:usa-tn:justia-tca-68-120-101 | partially_verified |
| edge:usa-tn:005 | ahj:usa-tn:state-fire-marshal | audits | local governments choosing to enforce their own code at least every three years | src:usa-tn:justia-tca-68-120-101; src:usa-tn:cornell-0780-02-02-05 | partially_verified |
| edge:usa-tn:006 | ahj:usa-tn:state-fire-marshal | supersedes_conflicting_interpretation | local application/interpretation of same or substantially identical building/fire standards | src:usa-tn:justia-tca-68-120-101; src:usa-tn:findlaw-tca-68-120-106 | partially_verified |
| edge:usa-tn:007 | local_jurisdictions | may_opt_out | statewide one- and two-family dwelling standards within opt-out jurisdiction boundaries | src:usa-tn:tdci-opt-out; src:usa-tn:justia-tca-68-120-101 | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | Tennessee minimum statewide building construction safety standards | International Building Code | 2021 | adopted_with_amendments | 2025-01-17 | 2025-04-17 | 2025-04-17 | 2025-08-16 for SFMO-plan-approval submissions after transition | Previous codes may be submitted for SFMO plan approval through 2025-08-15. | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:tdci-codes-enforcement; src:usa-tn:rule-0780-02-02; src:usa-tn:rule-filing-2025-01-17 |
| Residential | One and Two Family Dwellings and Townhouses standards | International Residential Code | 2018 | adopted_with_amendments | null | 2020-07-16 | 2020-07-16 | null | Applies where statewide standards apply or where local adoption uses compatible standards; opt-out jurisdictions require separate handling. | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:tdci-residential-permits; src:usa-tn:rule-0780-02-23 |
| Existing Building / Rehabilitation | Tennessee minimum standards for existing buildings | International Existing Building Code | 2021 | adopted_with_amendments | 2025-01-17 | 2025-04-17 | 2025-04-17 | 2025-08-16 for SFMO-plan-approval submissions after transition | Previous codes may be submitted for SFMO plan approval through 2025-08-15. | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:tdci-codes-enforcement; src:usa-tn:rule-0780-02-02 |
| Mechanical | Tennessee minimum standards for mechanical systems | International Mechanical Code | 2021 | adopted | 2025-01-17 | 2025-04-17 | 2025-04-17 | 2025-08-16 for SFMO-plan-approval submissions after transition | Previous codes may be submitted for SFMO plan approval through 2025-08-15. | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:tdci-codes-enforcement; src:usa-tn:rule-0780-02-02 |
| Plumbing | Tennessee minimum standards for plumbing systems | International Plumbing Code | 2021 | adopted | 2025-01-17 | 2025-04-17 | 2025-04-17 | 2025-08-16 for SFMO-plan-approval submissions after transition | Previous codes may be submitted for SFMO plan approval through 2025-08-15. | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:tdci-codes-enforcement; src:usa-tn:rule-0780-02-02 |
| Fuel Gas | Tennessee minimum standards for fuel gas systems | International Fuel Gas Code | 2021 | adopted | 2025-01-17 | 2025-04-17 | 2025-04-17 | 2025-08-16 for SFMO-plan-approval submissions after transition | Previous codes may be submitted for SFMO plan approval through 2025-08-15. | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:tdci-codes-enforcement; src:usa-tn:rule-0780-02-02 |
| Electrical | Electrical Installations | National Electrical Code | 2017 | adopted_with_exceptions | 2017-12-19 | 2018-03-19 | 2018-10-01 | null | Rule/effective versus code-operative date split should be rechecked against official PDF. | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:rule-0780-02-01; src:usa-tn:cornell-0780-02-01-02 |
| Energy - commercial | Tennessee minimum commercial energy standards | International Energy Conservation Code | 2021 | adopted_with_amendments | 2025-01-17 | 2025-04-17 | 2025-04-17 | 2025-08-16 for SFMO-plan-approval submissions after transition | Previous codes may be submitted for SFMO plan approval through 2025-08-15; special occupancy exceptions use older editions. | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:tdci-codes-enforcement; src:usa-tn:rule-0780-02-02 |
| Energy - residential | One- and two-family dwelling/townhome energy standards | International Energy Conservation Code / IRC Chapter 11 path | 2018 | adopted_with_amendments | null | 2020-07-16 | 2020-07-16 | null | SFMO FAQ permits compliance through 2018 IRC Chapter 11 or 2018 IECC. | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:tdci-residential-permits; src:usa-tn:rule-0780-02-23 |
| Property Maintenance | Tennessee minimum property maintenance standards | International Property Maintenance Code | 2021 | adopted_with_amendments | 2025-01-17 | 2025-04-17 | 2025-04-17 | 2025-08-16 for SFMO-plan-approval submissions after transition | Previous codes may be submitted for SFMO plan approval through 2025-08-15 where relevant. | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:rule-0780-02-02 |
| Fire - construction references | Tennessee minimum fire construction standards | International Fire Code | 2021 | adopted_with_amendments | 2025-01-17 | 2025-04-17 | 2025-04-17 | 2025-08-16 for SFMO-plan-approval submissions after transition | Previous codes may be submitted for SFMO plan approval through 2025-08-15. | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:tdci-codes-enforcement; src:usa-tn:rule-0780-02-02 |
| Fire - operational / prevention code | Tennessee fire prevention / fire protection standards | International Fire Code; NFPA 101 only for specified new small residential board-and-care facilities | 2021 | adopted_limited_scope | 2025-01-17 | 2025-04-17 | 2025-04-17 | 2025-08-16 for SFMO-plan-approval submissions after transition | Operational enforcement scope and occupancy-specific inspection rules require more parsing. | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:tdci-codes-enforcement; src:usa-tn:rule-0780-02-02 |
| Accessibility | Tennessee Public Buildings Accessibility Act / state accessibility rules | statute/state accessibility framework; IBC Chapter 11 removed | current statute not edition-based | partially_verified | null | null | null | null | IBC Chapter 11 is removed; TPBAA prevails for accessibility conflicts. | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:tdci-accessibility-act; src:usa-tn:cornell-0780-02-02-01 |
| Elevator / Conveyance | unknown | unknown | unknown | unresolved | null | null | null | null | Elevator/conveyance authority not checked. | none |

### 3.2 Adoption Records

```yaml
adoption_records:
  - record_id: "adoption:usa-tn:ibc-2021"
    code_family: "Building"
    model_code: "International Building Code"
    edition: "2021"
    state_amendments: true
    adoption_date: "2025-01-17"
    adoption_date_note: "Rule filing date, not separately confirmed as agency adoption vote date."
    effective_date: "2025-04-17"
    transition_rule_ids:
      - "date-rule:usa-tn:2025-sfmo-plan-transition"
    source_ids:
      - "src:usa-tn:tdci-currently-adopted-codes"
      - "src:usa-tn:tdci-codes-enforcement"
      - "src:usa-tn:rule-0780-02-02"

  - record_id: "adoption:usa-tn:ifc-2021"
    code_family: "Fire - construction references; Fire - operational / prevention code"
    model_code: "International Fire Code"
    edition: "2021"
    state_amendments: true
    adoption_date: "2025-01-17"
    effective_date: "2025-04-17"
    source_ids:
      - "src:usa-tn:tdci-currently-adopted-codes"
      - "src:usa-tn:rule-0780-02-02"

  - record_id: "adoption:usa-tn:ifgc-2021"
    code_family: "Fuel Gas"
    model_code: "International Fuel Gas Code"
    edition: "2021"
    adoption_date: "2025-01-17"
    effective_date: "2025-04-17"
    source_ids:
      - "src:usa-tn:tdci-currently-adopted-codes"
      - "src:usa-tn:rule-0780-02-02"

  - record_id: "adoption:usa-tn:imc-2021"
    code_family: "Mechanical"
    model_code: "International Mechanical Code"
    edition: "2021"
    adoption_date: "2025-01-17"
    effective_date: "2025-04-17"
    source_ids:
      - "src:usa-tn:tdci-currently-adopted-codes"
      - "src:usa-tn:rule-0780-02-02"

  - record_id: "adoption:usa-tn:ipc-2021"
    code_family: "Plumbing"
    model_code: "International Plumbing Code"
    edition: "2021"
    adoption_date: "2025-01-17"
    effective_date: "2025-04-17"
    source_ids:
      - "src:usa-tn:tdci-currently-adopted-codes"
      - "src:usa-tn:rule-0780-02-02"

  - record_id: "adoption:usa-tn:iecc-commercial-2021"
    code_family: "Energy - commercial"
    model_code: "International Energy Conservation Code"
    edition: "2021"
    state_amendments: true
    adoption_date: "2025-01-17"
    effective_date: "2025-04-17"
    source_ids:
      - "src:usa-tn:tdci-currently-adopted-codes"
      - "src:usa-tn:rule-0780-02-02"

  - record_id: "adoption:usa-tn:iebc-2021"
    code_family: "Existing Building / Rehabilitation"
    model_code: "International Existing Building Code"
    edition: "2021"
    state_amendments: true
    adoption_date: "2025-01-17"
    effective_date: "2025-04-17"
    source_ids:
      - "src:usa-tn:tdci-currently-adopted-codes"
      - "src:usa-tn:rule-0780-02-02"

  - record_id: "adoption:usa-tn:ipmc-2021"
    code_family: "Property Maintenance"
    model_code: "International Property Maintenance Code"
    edition: "2021"
    state_amendments: true
    adoption_date: "2025-01-17"
    effective_date: "2025-04-17"
    source_ids:
      - "src:usa-tn:tdci-currently-adopted-codes"
      - "src:usa-tn:rule-0780-02-02"

  - record_id: "adoption:usa-tn:irc-2018"
    code_family: "Residential"
    model_code: "International Residential Code"
    edition: "2018"
    state_amendments: true
    adoption_date: null
    effective_date: "2020-07-16"
    source_ids:
      - "src:usa-tn:tdci-currently-adopted-codes"
      - "src:usa-tn:tdci-residential-permits"
      - "src:usa-tn:rule-0780-02-23"

  - record_id: "adoption:usa-tn:iecc-residential-2018"
    code_family: "Energy - residential"
    model_code: "International Energy Conservation Code"
    edition: "2018"
    state_amendments: true
    adoption_date: null
    effective_date: "2020-07-16"
    source_ids:
      - "src:usa-tn:tdci-currently-adopted-codes"
      - "src:usa-tn:tdci-residential-permits"
      - "src:usa-tn:rule-0780-02-23"

  - record_id: "adoption:usa-tn:nec-2017"
    code_family: "Electrical"
    model_code: "National Electrical Code"
    edition: "2017"
    state_amendments: true
    adoption_date: "2017-12-19"
    effective_date: "2018-03-19"
    operative_date: "2018-10-01"
    date_note: "Current-code page history and rule text appear to distinguish rule effective date from NEC application date; recheck official PDF before final verification."
    source_ids:
      - "src:usa-tn:tdci-currently-adopted-codes"
      - "src:usa-tn:rule-0780-02-01"
      - "src:usa-tn:cornell-0780-02-01-02"
```

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

The 2025 commercial/fire/code update was filed with the Secretary of State on 2025-01-17 and became effective on 2025-04-17. SFMO states that projects requiring SFMO plan approval may be submitted under previous codes for 120 days after 2025-04-17, through 2025-08-15. The current extraction supports a plan-submission transition for SFMO plan-approval projects; it does not prove a universal permit-date or local-jurisdiction transition rule.

Tennessee's code-and-standards rule also includes a nonretroactivity principle for existing buildings: an existing building that complied with legally effective standards at the time of construction is not subject to later adopted standards unless the nonconformity creates a serious life-safety hazard; construction undertaken after the rule's effective date must comply with the standards adopted by reference.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-tn:2025-sfmo-plan-transition | Construction projects requiring SFMO plan approval | grace_period / plan_submission_rule | 2025-04-17 through 2025-08-15 | SFMO approval required; plans submitted within 120 days after 2025-04-17 | yes, for submission under previous codes | src:usa-tn:tdci-codes-enforcement; src:usa-tn:rule-0780-02-03; src:usa-tn:rule-filing-2025-01-17 | verified_core |
| date-rule:usa-tn:retroactive-existing-buildings | Existing buildings | nonretroactivity_life_safety_exception | Standards legally effective at time of construction | Existing building conformed when built and does not pose serious life-safety hazard | yes | src:usa-tn:rule-0780-02-02; src:usa-tn:cornell-0780-02-02-01 | partially_verified |
| date-rule:usa-tn:local-code-seven-year-currency | Local jurisdictions adopting/enforcing codes | local_code_currency_rule | Within seven years of most current published edition | Local jurisdiction enforces its own building construction/fire safety code | conditional; stale local codes can trigger statewide standards | src:usa-tn:justia-tca-68-120-101; src:usa-tn:cornell-0780-02-02-05 | partially_verified |
| date-rule:usa-tn:residential-additions-permit | One- and two-family residential additions | applicability_date | 2011-10-01 | Additions over 30 square feet to existing homes | no | src:usa-tn:tdci-residential-permits | partially_verified |
| date-rule:usa-tn:nec-2017-operative | Electrical installations | operative_date | 2018-10-01 | NEC 2017 application under Electrical Installations rule | unresolved | src:usa-tn:rule-0780-02-01; src:usa-tn:cornell-0780-02-01-02 | partially_verified |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| statewide commercial/fire/building codes | none verified beyond 2025 update | null | null | null | null | null | monitor | src:usa-tn:tdci-codes-enforcement; src:usa-tn:tdci-currently-adopted-codes | 2025 update is already effective; no later pending statewide code adoption was verified. |
| electrical | unknown | null | null | null | null | null | monitor | src:usa-tn:rule-0780-02-01 | A 2025 revised electrical-installations PDF exists, but no newer NEC edition adoption was verified. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-tn:nfpa-101-small-board-care | Fire / Life Safety | New Small Residential Board and Care facilities | Change of occupancy from residential or health care; eight or fewer residents; residents can move reliably to safety as a group within three minutes | 2021 NFPA 101 governs this limited facility class under current SFMO code list. | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:rule-0780-02-02 | verified_core |
| applicability-rule:usa-tn:one-two-family-opt-out | Residential | One- and two-family dwellings | County or municipality opts out by statutory procedure | Counties/municipalities may exempt their jurisdiction from statewide one- and two-family standards; owners may request SFMO inspection in opt-out/non-code jurisdictions. | src:usa-tn:tdci-opt-out; src:usa-tn:justia-tca-68-120-101 | partially_verified |
| applicability-rule:usa-tn:local-sprinkler-vote | Residential fire sprinklers | One- and two-family dwellings | Local government adopts mandatory sprinklers | State standards cannot include mandatory one-/two-family sprinkler requirements, but local governments may adopt stricter requirements using special vote/procedure. | src:usa-tn:justia-tca-68-120-101 | partially_verified |
| applicability-rule:usa-tn:accessibility | Accessibility | Public buildings/accessibility conflicts | Accessibility provision conflict | IBC Chapter 11 is removed from the Tennessee IBC adoption; Tennessee Public Buildings Accessibility Act prevails for accessibility conflicts. | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:tdci-accessibility-act; src:usa-tn:cornell-0780-02-02-01 | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** state_adopts_model_codes_by_rule_with_appendix_and_rule_text_amendments

**Where amendments are published:** SFMO Current Codes page Appendix; Tennessee Secretary of State rules in Tenn. Comp. R. & Regs. 0780-02-02, 0780-02-23, and 0780-02-01; rulemaking filings for recent amendments.

**Amendment parsing status:** partially_parsed_high_impact_only

### 5.2 State Amendment Sources

| Amendment Source ID | Code Family | Publication Path | Parsed Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| amendment-source:usa-tn:0780-02-02 | Building, Fire, Fuel Gas, Mechanical, Plumbing, Commercial Energy, Existing Building, Property Maintenance, limited NFPA 101 | Tenn. Comp. R. & Regs. 0780-02-02 and SFMO Appendix | High-impact summary only | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:rule-0780-02-02; src:usa-tn:cornell-0780-02-02-01 | partially_parsed |
| amendment-source:usa-tn:0780-02-23 | Residential and residential energy | Tenn. Comp. R. & Regs. 0780-02-23 | High-impact summary only | src:usa-tn:tdci-residential-permits; src:usa-tn:rule-0780-02-23; src:usa-tn:cornell-0780-02-23-02 | partially_parsed |
| amendment-source:usa-tn:0780-02-01 | Electrical | Tenn. Comp. R. & Regs. 0780-02-01 | Listed NEC exceptions only | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:rule-0780-02-01; src:usa-tn:cornell-0780-02-01-02 | partially_parsed |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Base Provision | State Treatment | Source IDs | Confidence | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| amendment:usa-tn:ibc-chapter-11 | Building / Accessibility | IBC Chapter 11 Accessibility | Removed from Tennessee IBC adoption; accessibility governed separately. | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:cornell-0780-02-02-01 | 0.80 | Full TPBAA parsing remains open. |
| amendment:usa-tn:ibc-storm-shelters | Building | IBC Section 423 | Storm-shelter requirements are not required. | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:rule-0780-02-02 | 0.85 | Listed in SFMO Appendix. |
| amendment:usa-tn:ibc-ifc-13r | Building / Fire | IBC/IFC NFPA 13R provisions | Tennessee replaces NFPA 13R language for Group R occupancies with state-specific height/story conditions. | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:rule-0780-02-02 | 0.80 | Summary only; full text not normalized. |
| amendment:usa-tn:classroom-locking | Building / Fire | School and university classroom door locking | Tennessee permits classroom doors to be locked to prevent unwanted entry if specified NFPA 2021 requirements are met. | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:rule-0780-02-02 | 0.85 | Applies statewide for Group E and college/university instructional buildings as stated in Appendix. |
| amendment:usa-tn:flammable-gas-2024 | Building / Fire | IBC/IFC flammable gas provisions | Certain flammable-gas provisions are replaced by 2024 IBC/IFC text. | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:rule-0780-02-02 | 0.75 | Full replaced sections not normalized. |
| amendment:usa-tn:iecc-commercial-removed-controls | Energy - commercial | IECC automatic receptacle control, energy monitoring, commissioning | C405.11, C405.12, and C408 are removed. | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:rule-0780-02-02 | 0.80 | Listed in SFMO Appendix. |
| amendment:usa-tn:iecc-f-s-occupancies | Energy - commercial | IECC edition for F-1, F-2, S-1, S-2 | 2006 IECC applies to specified factory/storage occupancies. | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:rule-0780-02-02 | 0.80 | Specific occupancy classifications listed in Appendix. |
| amendment:usa-tn:irc-sprinklers | Residential | IRC Section R313 | Residential sprinklers are not mandatory under state standards. | src:usa-tn:tdci-residential-permits; src:usa-tn:rule-0780-02-23; src:usa-tn:justia-tca-68-120-101 | 0.85 | Local governments may adopt stricter one-/two-family sprinkler requirements under special procedure. |
| amendment:usa-tn:irc-electrical-chapters | Residential / Electrical | IRC Chapters 34-43 | Deleted; electrical standards adopted in 0780-02-01 apply. | src:usa-tn:rule-0780-02-23; src:usa-tn:cornell-0780-02-23-02 | 0.80 | NEC rule should be used for electrical requirements. |
| amendment:usa-tn:nec-afci | Electrical | NEC AFCI / available fault current | SFMO current list states available fault current is optional and certain AFCI applications are optional. | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:rule-0780-02-01 | 0.75 | Full electrical rule parsing remains open. |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-tn"
  model: "hybrid_state_minimum_with_local_certification_opt_out_and_state_fallback"
  enforcing_entities:
    - "Tennessee State Fire Marshal's Office / Codes Enforcement section"
    - "Local jurisdictions that certify adoption and enforcement of qualifying building construction and fire safety codes"
    - "Municipal fire prevention or building officials, fire chiefs, mayors, and county/state officials with concurrent jurisdiction under T.C.A. § 68-120-106, subject to caveat on unofficial statutory mirror"
    - "Deputy building inspectors appointed/authorized for one- and two-family dwelling inspections"
  required_officials_or_programs:
    - "Local jurisdictions adopting/enforcing codes must perform plan review and inspections within statutory/rule timeframes."
    - "SFMO audits local governments choosing to enforce their own codes at least every three years."
    - "The Commissioner/SFMO must provide a program for timely one- and two-family dwelling inspection services."
  state_reserved_activities:
    - "State buildings, educational occupancies, and other occupancies requiring SFMO inspection for initial licensure remain state-sensitive categories."
    - "SFMO enforces statewide codes for building categories not adopted/enforced by a local jurisdiction."
    - "SFMO determination supersedes conflicting local application or interpretation of the same or substantially identical standards."
    - "Owner-requested SFMO inspection is available in opt-out/non-code one- and two-family jurisdictions."
  source_ids:
    - "src:usa-tn:tdci-codes-enforcement"
    - "src:usa-tn:tdci-opt-out"
    - "src:usa-tn:lexis-tca-68-120-101"
    - "src:usa-tn:justia-tca-68-120-101"
    - "src:usa-tn:findlaw-tca-68-120-106"
    - "src:usa-tn:cornell-0780-02-02-05"
  verification_status: "partially_verified"
  confidence: 0.74
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-tn"
  model: "no_less_stringent_unless_statute_or_division_written_approval"
  applies_to_code_families:
    - "building construction safety"
    - "fire prevention"
    - "fire protection"
    - "one- and two-family residential standards where local jurisdiction has not opted out"
  approval_required: true
  approving_authority_id: "ahj:usa-tn:state-fire-marshal"
  filing_required: true
  filing_requirement: "Local jurisdiction adopting and enforcing building construction and fire safety codes must forward all relevant local ordinances to the Division within 60 days after adoption."
  registry_exists: null
  registry_source_ids: []
  legal_basis_source_ids:
    - "src:usa-tn:rule-0780-02-02"
    - "src:usa-tn:cornell-0780-02-02-05"
    - "src:usa-tn:justia-tca-68-120-101"
  verification_status: "partially_verified"
  confidence: 0.78
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement authority and local amendment authority are separate in Tennessee. A local jurisdiction may adopt and enforce qualifying code families, but that does not create unlimited amendment authority. Local standards must remain at least as stringent as Tennessee's state minimum unless a statutory exception applies or the Division approves the reduction in writing. Local jurisdictions also must keep building/fire code publications within seven years of the most current published edition, and their enforcement adequacy is subject to SFMO audit.

### 6.4 Known Local Amendment Registries

No comprehensive statewide local-amendment registry was identified. The rule requiring local jurisdictions to forward relevant ordinances to the Division suggests that SFMO may hold local ordinance information, but a public registry source was not verified.

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendments were normalized in this pass. Local examples found during search were not used because the report scope is statewide and the relevant statewide registry/publication path remains unresolved.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: scoped_not_implemented

Jurisdiction stack:

```text
Address
  -> State of Tennessee
  -> County
  -> Municipality / unincorporated county
  -> Determine whether location is an opt-out/non-code jurisdiction for one- and two-family dwellings
  -> Determine whether local jurisdiction certifies adoption/enforcement for all buildings, nonresidential buildings, or one-/two-family dwellings only
  -> Building AHJ: local enforcing jurisdiction or SFMO fallback
  -> Fire AHJ: local fire/building officials with concurrent jurisdiction plus SFMO conflict/reserved-review authority
  -> Electrical AHJ / inspection program under 0780-02-01
  -> Applicable state code adoption records
  -> Applicable local ordinance/amendment records if available
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | TIGER/Line or state GIS source not selected | none | statewide | unknown | pending |
| County | TIGER/Line or state GIS source not selected | none | statewide | unknown | pending |
| Municipality | TIGER/Line or state/local GIS source not selected | none | statewide | unknown | pending |
| Opt-out / non-code jurisdiction status | SFMO opt-out/non-code jurisdiction page and associated permit workflow | src:usa-tn:tdci-opt-out | statewide residential applicability concept; list/geometry not extracted | unknown | pending_extraction |
| Local code-adopting jurisdictions | SFMO exempt-jurisdiction materials referenced from Codes Enforcement page | src:usa-tn:tdci-codes-enforcement | statewide concept; registry not parsed | unknown | pending_extraction |
| Fire District | not selected | none | statewide | unknown | pending |
| Special District | not selected | none | statewide | unknown | pending |

### 7.3 AHJ Contact Data

| Contact ID | AHJ / Entity | Scope | Contact Data | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| contact:usa-tn:sfmo-codes-enforcement | Tennessee Department of Commerce and Insurance / State Fire Marshal's Office | Statewide fire prevention / codes enforcement | Department of Commerce and Insurance, 500 James Robertson Pkwy, Nashville, TN 37243-0565; main phone 615-741-2241. Residential inspection questions page lists 615-741-7170. | src:usa-tn:tdci-codes-enforcement; src:usa-tn:tdci-opt-out | partially_verified |
| contact:usa-tn:local-ahj | Local building/fire AHJs | Local jurisdictions | Not populated; requires address-specific lookup. | none | unresolved |

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Publisher | Source Type | Date / Currency | Access Path | Supports |
| --- | --- | --- | --- | --- | --- | --- |
| src:usa-tn:tdci-codes-enforcement | Codes Enforcement | Tennessee Department of Commerce and Insurance / State Fire Marshal's Office | agency_page | updated after 2025 filing; accessed 2026-06-26 | https://www.tn.gov/commerce/fire/codes-enforcement.html | SFMO enforcement role; 2025 effective date; 120-day transition; official links to rules and statute |
| src:usa-tn:tdci-currently-adopted-codes | Tennessee State Fire Marshal's Office Currently Adopted Codes | Tennessee Department of Commerce and Insurance / State Fire Marshal's Office | agency_page | page last updated 2025-04-17 | https://www.tn.gov/commerce/fire/codes-enforcement/history.html | Current code editions; adoption history; high-impact amendments appendix |
| src:usa-tn:rule-filing-2025-01-17 | Rulemaking Hearing Rule(s) Filing Form, filed 2025-01-17 | Tennessee Secretary of State | rulemaking_filing_pdf | filed 2025-01-17; effective 2025-04-17 | https://publications.tnsosfiles.com/rules_filings/01-17-25.pdf | 2025 code update filing; effective date; amendments to 0780-02-02 and related rules |
| src:usa-tn:rule-0780-02-02 | Tenn. Comp. R. & Regs. Chapter 0780-02-02, Codes and Standards | Tennessee Secretary of State | administrative_rule_pdf | revised 2025-04-17 | https://publications.tnsosfiles.com/rules/0780/0780-02/0780-02-02.20250417.pdf | Adoption by reference; state amendments; local ordinances; nonretroactivity; conflicts |
| src:usa-tn:rule-0780-02-03 | Tenn. Comp. R. & Regs. Chapter 0780-02-03, Review of Construction Plans and Specifications | Tennessee Secretary of State | administrative_rule_pdf | revised 2025-04-23 | https://publications.tnsosfiles.com/rules/0780/0780-02/0780-02-03.20250423.pdf | SFMO plan-review requirements; transition rule referenced by TDCI |
| src:usa-tn:rule-0780-02-23 | Tenn. Comp. R. & Regs. Chapter 0780-02-23, One and Two Family Dwellings and Townhouses | Tennessee Secretary of State | administrative_rule_pdf | revised 2020-07-16 / later copies found | https://publications.tnsosfiles.com/rules/0780/0780-02/0780-02-23.20200716.pdf | Residential IRC/IECC adoption and amendments |
| src:usa-tn:rule-0780-02-01 | Tenn. Comp. R. & Regs. Chapter 0780-02-01, Electrical Installations | Tennessee Secretary of State | administrative_rule_pdf | revised 2025-07-14 search result; NEC 2017 adoption carried forward | https://publications.tnsosfiles.com/rules/0780/0780-02/0780-02-01.20250714.pdf | Electrical installations; NEC 2017 adoption; electrical exceptions |
| src:usa-tn:lexis-tca-68-120-101 | Tennessee Code Unannotated - Free Public Access | LexisNexis for Tennessee Code public access | statutory_code_portal | current portal accessed 2026-06-26 | https://advance.lexis.com/container?config=014CJAA5ZGVhZjA3NS02MmMzLTRlZWQtOGJjNC00YzQ1MmZlNzc2YWYKAFBvZENhdGFsb2e9zYpNUjTRaIWVfyrur9ud | Official/public code access path for T.C.A. Title 68, Chapter 120 |
| src:usa-tn:tdci-opt-out | Statewide Building Construction Code Inspections in Opt-Out or Non-Code Jurisdictions | Tennessee Department of Commerce and Insurance / State Fire Marshal's Office | agency_page | accessed 2026-06-26 | https://www.tn.gov/commerce/fire/residential-permits/opt-out-jurisdictions.html | Residential opt-out/non-code jurisdictions; owner-requested SFMO inspection |
| src:usa-tn:tdci-residential-permits | Residential Permits and Residential Permit FAQs | Tennessee Department of Commerce and Insurance / State Fire Marshal's Office | agency_page | accessed 2026-06-26 | https://www.tn.gov/commerce/fire/residential-permits.html and https://www.tn.gov/commerce/fire/residential-permits/fire-residential-faqs.html | 2018 IRC/IECC residential adoption; residential sprinkler status; additions over 30 square feet |
| src:usa-tn:tdci-accessibility-act | Tennessee Public Buildings Accessibility Act | Tennessee Department of Commerce and Insurance / State Fire Marshal's Office | agency_page | accessed 2026-06-26 | https://www.tn.gov/commerce/fire/codes-enforcement/tennessee-public-buildings-accessibility-act.html | Accessibility act access path and relationship to Title 68, Chapter 120 |
| src:usa-tn:tncourts-code-link | Tennessee Code - Lexis Law Link | Tennessee Administrative Office of the Courts | official_referral_page | accessed 2026-06-26 | https://www.tncourts.gov/Tennessee%20Code | State court page referring users to Lexis Tennessee Code access |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-tn:rule-0780-02-02 | pdf_access_limit | The Secretary of State PDF is the official rule source, but browser extraction returned access errors in this pass. Text was cross-checked against TDCI pages and mirrored regulation text. | Re-fetch and archive official PDF before final verification. |
| src:usa-tn:rule-0780-02-03 | pdf_access_limit | The official PDF was referenced by TDCI and searched, but full text was not extracted. Transition language is supported by TDCI's official summary. | Re-fetch official PDF and normalize Rule 0780-02-03-.11. |
| src:usa-tn:rule-filing-2025-01-17 | pdf_access_limit | The official filing was linked by TDCI and search-indexed, but full text was not extracted. | Re-fetch and archive official filing before marking verified. |
| src:usa-tn:rule-0780-02-23 | pdf_access_limit | Official residential PDF was search-indexed; mirrored text used for readability. | Re-fetch and compare against current Secretary of State rule page. |
| src:usa-tn:rule-0780-02-01 | rule_date_needs_recheck | Search result found a 2025-07-14 revised PDF, while SFMO adoption-history and mirrored text also reference earlier NEC effective/operative dates. Confirm the exact current PDF and date sequence directly. | Validate the exact current electrical rule PDF and effective dates. |
| src:usa-tn:lexis-tca-68-120-101 | portal_not_line_extracted | Lexis public code portal requires terms acceptance and was not line-extracted. | Use the official portal for final legal text check; keep mirrors as non-authoritative aids. |

### 8.3 Supplemental Sources

| Source ID | Title | Publisher | Source Type | Date / Currency | Access Path | Supports | Caveat |
| --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-tn:justia-tca-68-120-101 | Tennessee Code § 68-120-101, 2024 | Justia | statutory_mirror | 2024 code mirror; accessed 2026-06-26 | https://law.justia.com/codes/tennessee/title-68/safety/chapter-120/part-1/section-68-120-101/ | State fire marshal authority; local adoption/enforcement; opt-out; audits; state fallback; conflicts | Unofficial; confirm against Lexis before final verification. |
| src:usa-tn:findlaw-tca-68-120-106 | Tennessee Code § 68-120-106 | FindLaw | statutory_mirror | current as of 2024-01-02 per page; accessed 2026-06-26 | https://codes.findlaw.com/tn/title-68-health-safety-and-environmental-protection/tn-code-sect-68-120-106/ | Concurrent jurisdiction; state supersedes less stringent local ordinances | Unofficial; confirm against Lexis. |
| src:usa-tn:cornell-0780-02-02-01 | Tenn. Comp. R. & Regs. 0780-02-02-.01 | Cornell Legal Information Institute | regulation_mirror | accessed 2026-06-26 | https://www.law.cornell.edu/regulations/tennessee/T-C-A-SS-0780-02-02-.01 | Adoption by reference and amendments text | Mirror; confirm against TN SOS PDF. |
| src:usa-tn:cornell-0780-02-02-05 | Tenn. Comp. R. & Regs. 0780-02-02-.05 | Cornell Legal Information Institute | regulation_mirror | accessed 2026-06-26 | https://www.law.cornell.edu/regulations/tennessee/T-C-A-SS-0780-02-02-.05 | Local ordinance forwarding, seven-year currency, no-less-stringent rule | Mirror; confirm against TN SOS PDF. |
| src:usa-tn:cornell-0780-02-23-02 | Tenn. Comp. R. & Regs. 0780-02-23-.02 | Cornell Legal Information Institute | regulation_mirror | accessed 2026-06-26 | https://www.law.cornell.edu/regulations/tennessee/Tenn-Comp-R-Regs-0780-02-23-.02 | IRC 2018, Appendix Q, sprinkler and electrical-chapter amendments | Mirror; confirm against TN SOS PDF. |
| src:usa-tn:cornell-0780-02-01-02 | Tenn. Comp. R. & Regs. 0780-02-01-.02 | Cornell Legal Information Institute | regulation_mirror | accessed 2026-06-26 | https://www.law.cornell.edu/regulations/tennessee/Tenn-Comp-R-Regs-0780-02-01-.02 | NEC 2017 adoption and operative-date detail | Mirror; confirm against TN SOS PDF. |

### 8.4 Source Extraction Metadata

| Extraction ID | Source IDs | Extracted By | Extraction Date | Method | Coverage | Quality Notes |
| --- | --- | --- | --- | --- | --- | --- |
| extract:usa-tn:2026-06-26:tdci-pages | src:usa-tn:tdci-codes-enforcement; src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:tdci-opt-out; src:usa-tn:tdci-residential-permits; src:usa-tn:tdci-accessibility-act | ChatGPT | 2026-06-26 | web page retrieval | Core code editions, SFMO role, transition summary, opt-out model | High value; official agency pages. |
| extract:usa-tn:2026-06-26:sos-rules | src:usa-tn:rule-0780-02-02; src:usa-tn:rule-0780-02-03; src:usa-tn:rule-filing-2025-01-17; src:usa-tn:rule-0780-02-23; src:usa-tn:rule-0780-02-01 | ChatGPT | 2026-06-26 | search result snippets and official links; PDF open failed for several sources | Rule identifiers, effective dates, current editions, local ordinance references | Needs direct PDF validation before verified status. |
| extract:usa-tn:2026-06-26:statutes | src:usa-tn:lexis-tca-68-120-101; src:usa-tn:tncourts-code-link; src:usa-tn:justia-tca-68-120-101; src:usa-tn:findlaw-tca-68-120-106 | ChatGPT | 2026-06-26 | official portal discovery plus unofficial mirrors | Authority model, local enforcement, conflict, audits, opt-out | Statutory text must be rechecked against official Lexis portal. |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 0.90 | none | Core authority and code adoption fields are source-backed, but several fields remain open. |
| report | report.last_updated | 2026-06-26 | verified | 1.00 | none | Set during this update pass. |
| ahj:usa-tn:state-fire-marshal | authority.name | Tennessee Department of Commerce and Insurance, Division of Fire Prevention / State Fire Marshal's Office | verified_core | 0.85 | src:usa-tn:tdci-codes-enforcement; src:usa-tn:tdci-currently-adopted-codes | Official agency pages. |
| adoption:usa-tn:ibc-2021 | edition | 2021 IBC | verified_core | 0.90 | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:rule-0780-02-02 | Official current-code page and SOS rule link. |
| adoption:usa-tn:irc-2018 | edition | 2018 IRC | verified_core | 0.85 | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:tdci-residential-permits; src:usa-tn:rule-0780-02-23 | Official page plus rule source. |
| adoption:usa-tn:nec-2017 | edition | 2017 NEC | partially_verified | 0.75 | src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:rule-0780-02-01; src:usa-tn:cornell-0780-02-01-02 | Dates need official PDF recheck. |
| date-rule:usa-tn:2025-sfmo-plan-transition | transition_end | 2025-08-15 | verified_core | 0.90 | src:usa-tn:tdci-codes-enforcement; src:usa-tn:rule-0780-02-03 | TDCI states this date directly. |
| local-enforcement:usa-tn | model | hybrid_state_minimum_with_local_certification_opt_out_and_state_fallback | partially_verified | 0.74 | src:usa-tn:tdci-codes-enforcement; src:usa-tn:tdci-opt-out; src:usa-tn:justia-tca-68-120-101 | Statutory mirror must be checked against Lexis. |
| local-amendment-rule:usa-tn | model | no_less_stringent_unless_statute_or_division_written_approval | partially_verified | 0.78 | src:usa-tn:rule-0780-02-02; src:usa-tn:cornell-0780-02-02-05 | Official PDF needs archival. |
| src-registry | source IDs | All body source IDs should resolve to registry entries or `none` | verified | 1.00 | none | Checked by script on 2026-06-26. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Source-ID check passed against Section 8 registry. |
| All authority IDs resolve | pass | Primary authority is defined; unresolved elevator authority is explicitly labeled. |
| All current code families have adoption rows | pass | Current SFMO-listed code families have rows; elevator/conveyance remains unresolved because it was not part of verified SFMO adoption set. |
| Building and operational fire code are separated | pass | Fire construction references and operational/prevention scope are split. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Dates are kept separate, with nulls where unresolved. |
| Effective dates are valid ISO dates | pass | ISO date fields were checked where dates are present. |
| No impossible date sequences | pass | No adoption/effective/operative sequence conflict found in the populated rows, except NEC caveat explicitly noted. |
| Transition rules have explicit trigger conditions | pass | 2025 transition rule is limited to SFMO plan-approval submissions. |
| Permit-date logic is captured where applicable | fail | The report captures plan-submission transition logic, but not a universal permit-date/issuance-date rule for local jurisdictions. |
| Local enforcement model classified | pass | Model classified as hybrid state-minimum/local-certification/opt-out/fallback. |
| Local amendment rule classified | pass | Rule classified as no-less-stringent unless approved or statutory exception applies. |
| AHJ confirmation metadata present | fail | SFMO contact present, but local AHJ contacts and address-level resolver are not populated. |
| Official-source caveats captured | pass | Official PDF, statute portal, and mirror-use caveats are explicit. |
| No leftover template markers | pass | The standard placeholder-marker search found no matches. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-tn:001 | high | official statutes | T.C.A. § 68-120-101 and § 68-120-106 were interpreted using the Lexis access path plus unofficial mirrors because official Lexis text was not line-extracted. | Recheck statutory text in official Tennessee Code public access portal and archive excerpts. | null | null | open |
| issue:usa-tn:002 | high | official rule PDFs | Several TN SOS PDFs were linked and search-indexed but not directly full-text extracted due access errors. | Download or otherwise capture official PDFs for 0780-02-02, 0780-02-03, 0780-02-23, 0780-02-01, and 01-17-25 filing. | null | null | open |
| issue:usa-tn:003 | medium | operational fire scope | The current report separates IFC and limited NFPA 101, but did not fully parse Title 68, Chapter 102 or all occupancy-inspection rules. | Extract fire prevention statutes and Chapter 0780-02-03 occupancy-specific inspection provisions. | null | null | open |
| issue:usa-tn:004 | medium | local amendment registry | Rule requires forwarding local ordinances, but no public statewide registry was identified. | Identify whether SFMO publishes exempt jurisdiction/local ordinance lists or obtain via records request. | null | null | open |
| issue:usa-tn:005 | medium | AHJ resolver | Local enforcement varies by local adoption scope and residential opt-out status; no address-level resolver or boundary data is attached. | Extract exempt-jurisdiction list, opt-out list, and boundary data. | null | null | open |
| issue:usa-tn:006 | low | elevator/conveyance | Elevator/conveyance authority was not checked. | Search Tennessee elevator/conveyance statutes and TDCI regulatory boards/programs. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-tn:tdci-codes-enforcement | src:usa-tn:tdci-codes-enforcement | html_diff | monthly | New notice, transition rule, filing link, or enforcement scope change | 2026-06-26 | active |
| watch:usa-tn:current-codes | src:usa-tn:tdci-currently-adopted-codes | html_diff | monthly | Current-code edition, appendix, or adoption-history change | 2026-06-26 | active |
| watch:usa-tn:rule-0780-02-02 | src:usa-tn:rule-0780-02-02 | pdf_diff | monthly | Revised Codes and Standards rule PDF | 2026-06-26 | active |
| watch:usa-tn:rule-0780-02-03 | src:usa-tn:rule-0780-02-03 | pdf_diff | monthly | Plan-review transition or submission-rule change | 2026-06-26 | active |
| watch:usa-tn:rule-0780-02-23 | src:usa-tn:rule-0780-02-23 | pdf_diff | monthly | Residential IRC/IECC amendment or edition change | 2026-06-26 | active |
| watch:usa-tn:rule-0780-02-01 | src:usa-tn:rule-0780-02-01 | pdf_diff | monthly | NEC edition or electrical amendment change | 2026-06-26 | active |
| watch:usa-tn:code-portal | src:usa-tn:lexis-tca-68-120-101 | statute_monitor | quarterly | Amendment to T.C.A. Title 68, Chapter 120 or Chapter 102 authority provisions | 2026-06-26 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-24 | Generated baseline draft report | report:usa-tn | none | Codex | Baseline draft had no primary sources. |
| 2026-06-26 | Populated Tennessee report with sourced authority, code-adoption, transition, amendment, local enforcement, source registry, and QA sections | report:usa-tn; ahj:usa-tn:state-fire-marshal; adoption:usa-tn:ibc-2021; adoption:usa-tn:ifc-2021; adoption:usa-tn:irc-2018; local-enforcement:usa-tn; local-amendment-rule:usa-tn | src:usa-tn:tdci-codes-enforcement; src:usa-tn:tdci-currently-adopted-codes; src:usa-tn:rule-0780-02-02; src:usa-tn:rule-0780-02-03; src:usa-tn:rule-0780-02-23; src:usa-tn:rule-0780-02-01; src:usa-tn:justia-tca-68-120-101 | ChatGPT | Upgraded to `partially_verified`; official PDF/statute caveats and open issues remain explicit. |
