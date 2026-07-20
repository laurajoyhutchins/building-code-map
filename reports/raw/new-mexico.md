---
state:
  state_id: "US-NM"
  name: "New Mexico"
  abbreviation: "NM"
report:
  report_id: "state-report:usa-nm"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-26"
  last_verified: "2026-06-26"
  reviewed_by: null
risk:
  overall_confidence: 0.64 # 0.00 - 1.00
  risk_flags:
    - "accessibility_scope_partially_unresolved"
    - "local_amendment_registry_not_found"
    - "official_statute_text_needs_current-compilation_crosscheck"
    - "some_rule_text_crosschecked_with_supplemental_mirrors"
  open_questions_count: 5

---

# State Building Code Authority Report: New Mexico

## 1. Executive Summary

- **Authority model:** New Mexico uses a statewide construction-code framework administered primarily by the Construction Industries Division (CID) of the Regulation and Licensing Department, with Construction Industries Commission (CIC) approval and trade-bureau recommendations under the Construction Industries Licensing Act. Political subdivisions are subject to codes adopted and approved under the Act as minimum requirements.

- **Statewide code status:** Core construction codes are administered through NMAC Title 14, including 2021 New Mexico commercial, residential, existing building, plumbing, mechanical, and energy codes, the 2020 New Mexico Electrical Code, and the 2019 New Mexico Safety Code for Elevators and Escalators under Chapter 15. Fire prevention and public occupancy rules are administered separately under Title 10, Chapter 25, Part 5.

- **Local enforcement model:** The evidence supports a hybrid model: CID-administered statewide rules with local AHJs, building officials, and fire marshals involved where they have permitting or enforcement authority. The fire rule expressly separates construction-code responsibility for building-permit projects from fire-marshal responsibility for fire protection systems.

- **Local amendment posture:** State construction codes appear to operate as minimum requirements for political subdivisions. More stringent local fire-protection requirements are allowed when they do not conflict with state fire rules. A statewide local-amendment registry or approval workflow was not confirmed.

- **Known transition periods or pending changes:** Residential and existing-building 2021 code transitions ran to 2023-12-14; electrical transition to 2023-09-28; energy transition to 2024-07-30; plumbing and mechanical transitions to 2022-09-10. A 14.7.2 NMAC commercial-building amendment became effective 2025-01-13.

- **Production readiness:** partially_ready_for_narrow_validation

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: CID adopts building codes and minimum standards recommended by trade bureaus
    and approved by CIC under CILA; CIC reviews and approves/disapproves codes, standards,
    and related rules.
  confidence: 0.82
  source_ids:
  - src:usa-nm:cila-rld-2021
  - src:usa-nm:supp-nmsa-60-13-44-justia
- topic: Primary building code edition
  finding: New Mexico adopted the 2021 IBC as the 2021 New Mexico Commercial Building
    Code, effective 2023-07-14.
  confidence: 0.88
  source_ids:
  - src:usa-nm:nm-register-2023-building
  - src:usa-nm:nmac-14-7-2
- topic: Residential code edition
  finding: New Mexico adopted the 2021 IRC as the 2021 New Mexico Residential Building
    Code, effective 2023-07-14 with mandatory use after 2023-12-14.
  confidence: 0.86
  source_ids:
  - src:usa-nm:nm-register-2023-building
  - src:usa-nm:nmac-14-7-3
- topic: Electrical code authority and edition
  finding: The 2020 NEC is adopted as the New Mexico Electrical Code, effective 2023-03-28
    with mandatory use after 2023-09-28.
  confidence: 0.88
  source_ids:
  - src:usa-nm:nm-register-2023-electrical
  - src:usa-nm:nmac-14-10-4
- topic: Fire code authority
  finding: The State Fire Marshal rule adopts the IFC as revised in 10.25.5 NMAC;
    fire protection systems remain under fire-marshal authority while other construction-code
    work remains with the building AHJ.
  confidence: 0.78
  source_ids:
  - src:usa-nm:nmac-10-25-5
  - src:usa-nm:nm-register-2022-fire-amend
  - src:usa-nm:supp-cornell-10-25-5-8
- topic: Local enforcement
  finding: Local AHJs enforce where they have permitting/enforcement authority, but
    the state rule set remains statewide and CID jurisdiction applies where applicable.
  confidence: 0.66
  source_ids:
  - src:usa-nm:nmac-14-5-1
  - src:usa-nm:nmac-14-5-2
  - src:usa-nm:nm-register-2022-fire-amend
- topic: Local amendments
  finding: Political-subdivision codes must meet state minimums; fire AHJs may impose
    more stringent, non-conflicting fire-protection requirements. Approval/filing
    mechanics remain unresolved.
  confidence: 0.58
  source_ids:
  - src:usa-nm:supp-nmsa-60-13-44-justia
  - src:usa-nm:nm-register-2022-fire-amend
- topic: Effective / operative date rule
  finding: NMAC code rules commonly use an effective date, permit-application cutoff,
    and code-specific grace period or mandatory date.
  confidence: 0.82
  source_ids:
  - src:usa-nm:nm-register-2023-building
  - src:usa-nm:nm-register-2023-electrical
  - src:usa-nm:nm-register-2024-energy
- topic: Elevator / conveyance code
  finding: CID now lists Elevator Safety Code parts 14.15.1 through 14.15.8, including
    the 2019 New Mexico Safety Code for Elevators and Escalators; existing commercial
    conveyances had a 2026-01-01 registration deadline.
  confidence: 0.78
  source_ids:
  - src:usa-nm:rld-elevator-bureau
  - src:usa-nm:nm-register-2025-elevator-rules
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-nm:cid |
| Authority name | Construction Industries Division, New Mexico Regulation and Licensing Department |
| Authority type | state_agency |
| Legal basis | Construction Industries Licensing Act; NMSA provisions including Sections 60-13-9 and 60-13-44, as referenced in NMAC code rules |
| Role | Administers statewide construction-code rules; adopts building codes and minimum standards recommended by trade bureaus and approved by CIC; administers permits, inspections, plan review, stop-work, and code enforcement where CID has jurisdiction |
| Enforcement model | hybrid_state_and_local_AHJ |
| Source IDs | src:usa-nm:cila-rld-2021; src:usa-nm:nmac-14-5-1; src:usa-nm:nmac-14-5-2; src:usa-nm:nm-register-2023-building |
| Verification status | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-nm:cid | Construction Industries Division | Administers 14.7.2 NMAC, 2021 New Mexico Commercial Building Code | NMSA 60-13-9, 60-13-10.3, 60-13-44 as cited in rule | src:usa-nm:nm-register-2023-building; src:usa-nm:nmac-14-7-2 | partially_verified |
| Residential | ahj:usa-nm:cid | Construction Industries Division | Administers 14.7.3 NMAC, 2021 New Mexico Residential Building Code | NMSA 60-13-9 and 60-13-44 as cited in rule | src:usa-nm:nm-register-2023-building; src:usa-nm:nmac-14-7-3 | partially_verified |
| Existing Building / Rehabilitation | ahj:usa-nm:cid | Construction Industries Division | Administers 14.7.7 NMAC, 2021 New Mexico Existing Building Code | NMSA 60-13-9 and 60-13-44 as cited in rule | src:usa-nm:nm-register-2023-building; src:usa-nm:nmac-14-7-7 | partially_verified |
| Mechanical | ahj:usa-nm:cid | Construction Industries Division / Mechanical Bureau | Administers 14.9.2 NMAC, 2021 New Mexico Mechanical Code | CILA trade-bureau and code-adoption authority | src:usa-nm:cila-rld-2021; src:usa-nm:nmac-14-9-2 | partially_verified |
| Plumbing | ahj:usa-nm:cid | Construction Industries Division / Mechanical Bureau | Administers 14.8.2 NMAC, 2021 New Mexico Plumbing Code | CILA trade-bureau and code-adoption authority | src:usa-nm:cila-rld-2021; src:usa-nm:nmac-14-8-2 | partially_verified |
| Fuel Gas | ahj:usa-nm:cid | Construction Industries Division / Mechanical Bureau | Fuel gas is handled through NMMC fuel-gas provisions and LP-gas references rather than a standalone IFGC adoption verified in this pass | CILA trade-bureau authority; code cross-reference rules | src:usa-nm:nmac-14-9-2; src:usa-nm:nmac-14-7-6 | partially_verified |
| Electrical | ahj:usa-nm:cid | Construction Industries Division / Electrical Bureau | Administers 14.10.4 NMAC, 2020 New Mexico Electrical Code | CILA trade-bureau and code-adoption authority | src:usa-nm:nm-register-2023-electrical; src:usa-nm:nmac-14-10-4 | partially_verified |
| Energy | ahj:usa-nm:cid | Construction Industries Division | Administers 14.7.6 and 14.7.9 NMAC, 2021 residential and commercial energy conservation codes | NMSA 60-13-9 and 60-13-44 as cited in rules | src:usa-nm:nm-register-2024-energy; src:usa-nm:nmac-14-7-6; src:usa-nm:nmac-14-7-9 | partially_verified |
| Fire - construction references | ahj:usa-nm:building-ahj | Building AHJ / chief building official | Construction-related fire-code sections for projects requiring building permits are handled by the building AHJ, except fire-protection systems | Fire rule construction-responsibility clause | src:usa-nm:nm-register-2022-fire-amend; src:usa-nm:nmac-10-25-5 | partially_verified |
| Fire - operational / prevention code | ahj:usa-nm:sfm | New Mexico State Fire Marshal / fire marshal with authority over permitting | Administers New Mexico Fire Code / fire prevention and public-occupancy rule; retains authority over fire-protection systems | 10.25.5 NMAC | src:usa-nm:nmac-10-25-5; src:usa-nm:nm-register-2022-fire-amend; src:usa-nm:supp-cornell-10-25-5-8 | partially_verified |
| Accessibility | ahj:usa-nm:cid | Construction Industries Division / General Construction Bureau | Recommends accessibility specifications for qualifying public buildings; current standalone accessibility code set was not parsed | CILA Section 60-13-44 mirror; adopted building/electrical code references | src:usa-nm:supp-nmsa-60-13-44-justia; src:usa-nm:nmac-14-7-2 | unresolved_scope |
| Elevator / Conveyance | ahj:usa-nm:cid-elevator-bureau | Construction Industries Division / Elevator Bureau | Administers Elevator Safety Code parts 14.15.1 through 14.15.8, including registration, permitting, inspections, licensing/certification, and the 2019 New Mexico Safety Code for Elevators and Escalators | Elevator Safety Act; Title 14, Chapter 15 Elevator Safety Code | src:usa-nm:rld-elevator-bureau; src:usa-nm:nm-register-2025-elevator-rules; src:usa-nm:nmac-14-7-2; src:usa-nm:supp-cornell-14-7-2-38 | partially_verified |

### 2.3 Authority Hierarchy Notes

The construction-code hierarchy is best represented as a state minimum-code system with CID/CIC adoption and local AHJ enforcement where local entities have permitting authority. CILA gives trade bureaus a role in recommending minimum standards; the commission reviews and adopts standards and codes by rule. Political subdivisions are subject to those state-adopted codes as minimum requirements. Fire prevention and public occupancy sit in a separate State Fire Marshal rule set, with the fire rule preserving fire-marshal authority over fire-protection systems while assigning other construction-code responsibilities on building-permit projects to the chief building official or AHJ.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-nm:001 | ahj:usa-nm:trade-bureaus | recommends_minimum_standards_to | ahj:usa-nm:cid / ahj:usa-nm:cic | src:usa-nm:cila-rld-2021; src:usa-nm:supp-nmsa-60-13-44-justia | partially_verified |
| edge:usa-nm:002 | ahj:usa-nm:cic | approves_or_disapproves | rules, standards, codes, and licensing requirements | src:usa-nm:cila-rld-2021 | partially_verified |
| edge:usa-nm:003 | ahj:usa-nm:cid | administers | Title 14 construction-code rules where CID has jurisdiction | src:usa-nm:nmac-14-5-1; src:usa-nm:nmac-14-5-2 | partially_verified |
| edge:usa-nm:004 | ahj:usa-nm:cid | sets_minimum_requirements_for | political_subdivision_codes | src:usa-nm:supp-nmsa-60-13-44-justia | partially_verified |
| edge:usa-nm:005 | ahj:usa-nm:sfm | administers | New Mexico Fire Code / fire prevention and public occupancy | src:usa-nm:nmac-10-25-5; src:usa-nm:nm-register-2022-fire-amend | partially_verified |
| edge:usa-nm:006 | ahj:usa-nm:sfm | reserves_authority_for | fire_protection_systems | src:usa-nm:nm-register-2022-fire-amend | partially_verified |
| edge:usa-nm:007 | ahj:usa-nm:building-ahj | responsible_for | construction-code sections of fire rule on building-permit projects, except fire-protection systems | src:usa-nm:nm-register-2022-fire-amend | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | 2021 New Mexico Commercial Building Code | International Building Code | 2021 | current | 2023-05-17 | 2023-07-14 | 2023-07-14 | 2023-07-14 | Applies to work on or after 2023-07-14 subject to CID jurisdiction unless a permit application was received before that date; one- and two-family dwellings and townhouses use 14.7.3; existing buildings may use 14.7.2 or 14.7.7 as applicable. | src:usa-nm:nm-register-2023-building; src:usa-nm:nmac-14-7-2 |
| Residential | 2021 New Mexico Residential Building Code | International Residential Code | 2021 | current | 2023-05-17 | 2023-07-14 | 2023-07-14 | 2023-12-14 | From publication through 2023-12-14, permits could issue under the previous rule or the new rule; after 2023-12-14, only the new rule could be used. | src:usa-nm:nm-register-2023-building; src:usa-nm:nmac-14-7-3 |
| Existing Building / Rehabilitation | 2021 New Mexico Existing Building Code | International Existing Building Code | 2021 | current | 2023-05-17 | 2023-07-14 | 2023-07-14 | 2023-12-14 | Same 2023-12-14 transition structure as the residential and existing-building 2021 replacement package. | src:usa-nm:nm-register-2023-building; src:usa-nm:nmac-14-7-7 |
| Mechanical | 2021 New Mexico Mechanical Code | Uniform Mechanical Code | 2021 | current | null | 2022-03-10 | 2022-03-10 | 2022-09-10 | Permits could issue under either the previous or new rule until 2022-09-10; after 2022-09-10 only the new rule could be used. | src:usa-nm:nmac-14-9-2 |
| Plumbing | 2021 New Mexico Plumbing Code | Uniform Plumbing Code | 2021 | current | null | 2022-03-10 | 2022-03-10 | 2022-09-10 | Permits could issue under either the previous or new rule until 2022-09-10; after 2022-09-10 only the new rule could be used. | src:usa-nm:nmac-14-8-2 |
| Fuel Gas | Fuel-gas provisions within the 2021 New Mexico Mechanical Code and LP-gas standards | Uniform Mechanical Code Chapter 13 / LP-gas standards cross-references | 2021 / unresolved for LP-gas details | partially_verified | null | 2022-03-10 | 2022-03-10 | 2022-09-10 | Standalone IFGC adoption was not confirmed; relevant rules deem IFGC references to NMMC or LP-gas standards. | src:usa-nm:nmac-14-9-2; src:usa-nm:nmac-14-7-6 |
| Electrical | 2020 New Mexico Electrical Code | National Electrical Code | 2020 | current | 2023-02-14 | 2023-03-28 | 2023-03-28 | 2023-09-28 | From publication through 2023-09-28, permits could issue under the previous or new rule; after 2023-09-28 only the new rule could be used. | src:usa-nm:nm-register-2023-electrical; src:usa-nm:nmac-14-10-4 |
| Energy | 2021 New Mexico Residential Energy Conservation Code and 2021 New Mexico Commercial Energy Conservation Code | International Energy Conservation Code | 2021 | current | null | 2024-01-30 | 2024-01-30 | 2024-07-30 | From publication through 2024-07-30, permits could issue under the previous or new rule; after 2024-07-30 only the new rule could be used. | src:usa-nm:nm-register-2024-energy; src:usa-nm:nmac-14-7-6; src:usa-nm:nmac-14-7-9 |
| Fire - construction references | New Mexico Fire Code construction interfaces | International Fire Code with New Mexico amendments and NMBC cross-references | 2021 | current | null | 2022-11-01 | 2022-11-01 | 2022-11-01 | Construction-related sections for construction projects requiring building permits are assigned to the chief building official/AHJ, except fire-protection systems. | src:usa-nm:nmac-10-25-5; src:usa-nm:nm-register-2022-fire-amend; src:usa-nm:supp-cornell-10-25-5-8 |
| Fire - operational / prevention code | New Mexico Fire Code / Fire Prevention and Public Occupancy | International Fire Code with New Mexico amendments | 2021 | current | null | 2022-11-01 | 2022-11-01 | 2022-11-01 | Local fire-protection requirements may be more stringent when they do not conflict with the state fire rule. | src:usa-nm:nmac-10-25-5; src:usa-nm:nm-register-2022-fire-amend; src:usa-nm:supp-cornell-10-25-5-8 |
| Accessibility | Public-building accessibility standards and adopted code accessibility provisions | Nationally recognized accessibility standard / IBC accessibility provisions | unresolved | partially_verified | null | null | null | null | CILA supports accessibility specifications for certain public buildings, but the current adopted standard and amendment details need a focused pass. | src:usa-nm:supp-nmsa-60-13-44-justia; src:usa-nm:nmac-14-7-2 |
| Elevator / Conveyance | 2019 New Mexico Safety Code for Elevators and Escalators; NMCBC Chapter 30 construction interface | elevator/escalator safety standards and IBC Chapter 30 interface | 2019 / 2021 interface | current | 2025-08-12 | 2025-08-12 | 2025-08-12 | 2026-01-01 | Title 14 Chapter 15 elevator rules were adopted in New Mexico Register Volume XXXVI, Issue 15; existing commercial conveyances had to be registered with CID by 2026-01-01. | src:usa-nm:rld-elevator-bureau; src:usa-nm:nm-register-2025-elevator-rules; src:usa-nm:nmac-14-7-2; src:usa-nm:supp-cornell-14-7-2-38 |

### 3.2 Adoption Records

```yaml
- adoption_id: adopt:usa-nm:commercial-building:2021-ibc
  code_family: Building
  authority_id: ahj:usa-nm:cid
  state_code_name: "2021 New Mexico Commercial Building Code"
  base_model_code: "2021 International Building Code"
  adoption_date: "2023-05-17"
  effective_date: "2023-07-14"
  operative_date: "2023-07-14"
  mandatory_date: "2023-07-14"
  transition_rule_ids:
    - date-rule:usa-nm:commercial-building-2023-permit-cutoff
  source_ids:
    - src:usa-nm:nm-register-2023-building
    - src:usa-nm:nmac-14-7-2
  confidence: 0.88

- adoption_id: adopt:usa-nm:residential-building:2021-irc
  code_family: Residential
  authority_id: ahj:usa-nm:cid
  state_code_name: "2021 New Mexico Residential Building Code"
  base_model_code: "2021 International Residential Code"
  adoption_date: "2023-05-17"
  effective_date: "2023-07-14"
  operative_date: "2023-07-14"
  mandatory_date: "2023-12-14"
  transition_rule_ids:
    - date-rule:usa-nm:residential-existing-2023-grace
  source_ids:
    - src:usa-nm:nm-register-2023-building
    - src:usa-nm:nmac-14-7-3
  confidence: 0.86

- adoption_id: adopt:usa-nm:existing-building:2021-iebc
  code_family: Existing Building / Rehabilitation
  authority_id: ahj:usa-nm:cid
  state_code_name: "2021 New Mexico Existing Building Code"
  base_model_code: "2021 International Existing Building Code"
  adoption_date: "2023-05-17"
  effective_date: "2023-07-14"
  operative_date: "2023-07-14"
  mandatory_date: "2023-12-14"
  transition_rule_ids:
    - date-rule:usa-nm:residential-existing-2023-grace
  source_ids:
    - src:usa-nm:nm-register-2023-building
    - src:usa-nm:nmac-14-7-7
  confidence: 0.84

- adoption_id: adopt:usa-nm:mechanical:2021-umc
  code_family: Mechanical
  authority_id: ahj:usa-nm:cid
  state_code_name: "2021 New Mexico Mechanical Code"
  base_model_code: "2021 Uniform Mechanical Code"
  adoption_date: null
  effective_date: "2022-03-10"
  operative_date: "2022-03-10"
  mandatory_date: "2022-09-10"
  transition_rule_ids:
    - date-rule:usa-nm:plumbing-mechanical-2022-grace
  source_ids:
    - src:usa-nm:nmac-14-9-2
  confidence: 0.76

- adoption_id: adopt:usa-nm:plumbing:2021-upc
  code_family: Plumbing
  authority_id: ahj:usa-nm:cid
  state_code_name: "2021 New Mexico Plumbing Code"
  base_model_code: "2021 Uniform Plumbing Code"
  adoption_date: null
  effective_date: "2022-03-10"
  operative_date: "2022-03-10"
  mandatory_date: "2022-09-10"
  transition_rule_ids:
    - date-rule:usa-nm:plumbing-mechanical-2022-grace
  source_ids:
    - src:usa-nm:nmac-14-8-2
  confidence: 0.76

- adoption_id: adopt:usa-nm:electrical:2020-nec
  code_family: Electrical
  authority_id: ahj:usa-nm:cid
  state_code_name: "2020 New Mexico Electrical Code"
  base_model_code: "2020 National Electrical Code"
  adoption_date: "2023-02-14"
  effective_date: "2023-03-28"
  operative_date: "2023-03-28"
  mandatory_date: "2023-09-28"
  transition_rule_ids:
    - date-rule:usa-nm:electrical-2023-grace
  source_ids:
    - src:usa-nm:nm-register-2023-electrical
    - src:usa-nm:nmac-14-10-4
  confidence: 0.88

- adoption_id: adopt:usa-nm:energy:2021-iecc
  code_family: Energy
  authority_id: ahj:usa-nm:cid
  state_code_name: "2021 New Mexico Residential Energy Conservation Code / 2021 New Mexico Commercial Energy Conservation Code"
  base_model_code: "2021 International Energy Conservation Code"
  adoption_date: null
  effective_date: "2024-01-30"
  operative_date: "2024-01-30"
  mandatory_date: "2024-07-30"
  transition_rule_ids:
    - date-rule:usa-nm:energy-2024-grace
  source_ids:
    - src:usa-nm:nm-register-2024-energy
    - src:usa-nm:nmac-14-7-6
    - src:usa-nm:nmac-14-7-9
  confidence: 0.86

- adoption_id: adopt:usa-nm:elevator:2019-nmscee
  code_family: Elevator / Conveyance
  authority_id: ahj:usa-nm:cid-elevator-bureau
  state_code_name: "2019 New Mexico Safety Code for Elevators and Escalators"
  base_model_code: "elevator/escalator safety standards; exact incorporated standards not extracted in this pass"
  adoption_date: "2025-08-12"
  effective_date: "2025-08-12"
  operative_date: "2025-08-12"
  mandatory_date: "2026-01-01"
  transition_rule_ids:
    - date-rule:usa-nm:elevator-2025-registration
  source_ids:
    - src:usa-nm:rld-elevator-bureau
    - src:usa-nm:nm-register-2025-elevator-rules
  confidence: 0.78

- adoption_id: adopt:usa-nm:fire:2021-ifc
  code_family: Fire - operational / prevention code
  authority_id: ahj:usa-nm:sfm
  state_code_name: "New Mexico Fire Code"
  base_model_code: "2021 International Fire Code"
  adoption_date: null
  effective_date: "2022-11-01"
  operative_date: "2022-11-01"
  mandatory_date: "2022-11-01"
  transition_rule_ids:
    - date-rule:usa-nm:fire-2022-responsibility-split
  source_ids:
    - src:usa-nm:nmac-10-25-5
    - src:usa-nm:nm-register-2022-fire-amend
    - src:usa-nm:supp-cornell-10-25-5-8
  confidence: 0.78
```

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

New Mexico code rules commonly distinguish the rule effective date from later mandatory dates. Several CID rules allow permits to be issued under either the previous rule or the new rule for a defined grace period. Scope provisions also use permit-application timing: work on or after the effective date is covered by the new rule unless the permit application was received before that date. Elevator Safety Code rules appeared in New Mexico Register Volume XXXVI, Issue 15, and the RLD Elevator Bureau page identifies a 2026-01-01 registration deadline for existing commercial conveyances. Fire-code transition logic is less completely parsed; the 2022 fire amendment provides a 2022-11-01 effective date and clarifies responsibility between fire marshals and building officials.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-nm:commercial-building-2023-permit-cutoff | 14.7.2 commercial building | permit_application_cutoff | 2023-07-14 | Work on or after the effective date is subject to 14.7.2 unless the permit application was received by CID before 2023-07-14. | yes, for applications received before effective date | src:usa-nm:nm-register-2023-building | partially_verified |
| date-rule:usa-nm:residential-existing-2023-grace | 14.7.3 residential; 14.7.7 existing building | grace_period_then_mandatory | publication through 2023-12-14 | Permit may be issued under previous or new rule during grace period; after 2023-12-14 only new rule. | yes, through 2023-12-14 | src:usa-nm:nm-register-2023-building; src:usa-nm:nmac-14-7-3; src:usa-nm:nmac-14-7-7 | partially_verified |
| date-rule:usa-nm:plumbing-mechanical-2022-grace | 14.8.2 plumbing; 14.9.2 mechanical | grace_period_then_mandatory | 2022-03-10 through 2022-09-10 | Permit may be issued under previous or new rule during grace period; after 2022-09-10 only new rule. | yes, through 2022-09-10 | src:usa-nm:nmac-14-8-2; src:usa-nm:nmac-14-9-2 | partially_verified |
| date-rule:usa-nm:electrical-2023-grace | 14.10.4 electrical | grace_period_then_mandatory | 2023-03-28 through 2023-09-28 | Permit may be issued under previous or new rule during grace period; after 2023-09-28 only new rule. | yes, through 2023-09-28 | src:usa-nm:nm-register-2023-electrical; src:usa-nm:nmac-14-10-4 | verified_from_sources_reviewed |
| date-rule:usa-nm:energy-2024-grace | 14.7.6 residential energy; 14.7.9 commercial energy | grace_period_then_mandatory | 2024-01-30 through 2024-07-30 | Permit may be issued under previous or new rule during grace period; after 2024-07-30 only new rule. | yes, through 2024-07-30 | src:usa-nm:nm-register-2024-energy; src:usa-nm:nmac-14-7-6; src:usa-nm:nmac-14-7-9 | verified_from_sources_reviewed |
| date-rule:usa-nm:elevator-2025-registration | 14.15 elevator safety | new_rule_plus_registration_deadline | 2025-08-12 rule issue; 2026-01-01 existing commercial conveyance registration deadline | Elevator Safety Code parts 14.15.1 through 14.15.8 were published as new rules; existing commercial conveyances had to be registered with CID by 2026-01-01. | not a prior-code grace rule | src:usa-nm:rld-elevator-bureau; src:usa-nm:nm-register-2025-elevator-rules | partially_verified |
| date-rule:usa-nm:fire-2022-responsibility-split | 10.25.5 fire prevention / New Mexico Fire Code | effective_date_and_authority_split | 2022-11-01 | Fire rule effective date plus responsibility split for building-permit projects and fire-protection systems. | unresolved | src:usa-nm:nm-register-2022-fire-amend; src:usa-nm:nmac-10-25-5 | partially_verified |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | 14.7.2 NMAC Section 11 amendment | 2024-12-10 | 2024-12-10 | 2025-01-13 | 2025-01-13 | 2025-01-13 | in_force_monitor | src:usa-nm:nm-register-2024-14-7-2-amend | Recent commercial-building amendment found; no newer statewide model-code replacement was captured in this pass. |
| Elevator / Conveyance | Title 14 Chapter 15 Elevator Safety Code | 2025-08-12 | 2025-08-12 | 2025-08-12 | 2025-08-12 | 2026-01-01 | in_force_monitor | src:usa-nm:rld-elevator-bureau; src:usa-nm:nm-register-2025-elevator-rules | Rules are in force; commercial-conveyance registration deadline is 2026-01-01. |
| Multiple code families | next statewide model-code cycle | null | null | null | null | null | watch | src:usa-nm:srca-title14-index; src:usa-nm:rld-building-codes-page | No pending replacement cycle was confirmed from the sources reviewed. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-nm:commercial-residential-split | Building / Residential | Commercial construction versus detached one- and two-family dwellings and townhouses | Project type | Commercial rule excludes detached one- and two-family dwellings and townhouses, which are governed by 14.7.3 NMAC. | src:usa-nm:nm-register-2023-building; src:usa-nm:nmac-14-7-2; src:usa-nm:nmac-14-7-3 | partially_verified |
| applicability-rule:usa-nm:existing-building-option | Building / Existing Building | Existing buildings | Existing-building work | Existing buildings may comply with 14.7.2 or 14.7.7, as applicable. | src:usa-nm:nm-register-2023-building; src:usa-nm:nmac-14-7-7 | partially_verified |
| applicability-rule:usa-nm:energy-residential-commercial | Energy | Residential and commercial energy conservation | Project type | Residential energy is in 14.7.6 NMAC and commercial energy is in 14.7.9 NMAC, both based on the 2021 IECC. | src:usa-nm:nm-register-2024-energy; src:usa-nm:nmac-14-7-6; src:usa-nm:nmac-14-7-9 | verified_from_sources_reviewed |
| applicability-rule:usa-nm:fire-construction-interface | Fire / Building | Construction projects requiring a building permit | Building permit required | Construction-related sections of the fire code are the responsibility of the chief building official/AHJ, except fire-protection systems, which remain under fire-marshal authority. | src:usa-nm:nm-register-2022-fire-amend | partially_verified |
| applicability-rule:usa-nm:fuel-gas-cross-reference | Fuel Gas | Fuel-gas references in adopted codes | IFGC reference appears in adopted code | IFGC references are treated as references to the New Mexico Mechanical Code or LP-gas standards, depending on subject matter. | src:usa-nm:nmac-14-9-2; src:usa-nm:nmac-14-7-6 | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** model-code adoption by reference with New Mexico-specific amendments embedded in NMAC parts. Rule section numbering generally maps to model-code chapters or administrative provisions.

**Where amendments are published:** New Mexico Administrative Code, chiefly Title 14 for construction codes and Title 10, Chapter 25 for fire prevention/public occupancy. Rule adoptions and amendments are published through the New Mexico Register and the State Records Center and Archives.

**Amendment parsing status:** high_level_only. This pass identified adoption editions, effective dates, grace periods, authority clauses, and selected high-impact cross-references. It did not produce a line-by-line amendment delta for each model code.

### 5.2 State Amendment Sources

| Amendment Source ID | Code Family | Publication Path | Parsed Level | Notes |
| --- | --- | --- | --- | --- |
| amend-source:usa-nm:14-7-2 | Building | 14.7.2 NMAC / New Mexico Register 2023 building package | adoption_and_selected_scope | Includes 2021 IBC adoption, effective date, scope, and Chapter 30 elevator cross-reference. |
| amend-source:usa-nm:14-7-3 | Residential | 14.7.3 NMAC / New Mexico Register 2023 building package | adoption_and_transition | Includes 2021 IRC adoption and 2023-12-14 mandatory date. |
| amend-source:usa-nm:14-7-7 | Existing Building | 14.7.7 NMAC / New Mexico Register 2023 building package | adoption_and_transition | Includes 2021 IEBC adoption and transition structure. |
| amend-source:usa-nm:14-7-6 | Residential Energy | 14.7.6 NMAC / New Mexico Register 2024 energy package | adoption_and_transition | Includes 2021 IECC adoption and 2024-07-30 mandatory date. |
| amend-source:usa-nm:14-7-9 | Commercial Energy | 14.7.9 NMAC / New Mexico Register 2024 energy package | adoption_and_transition | Includes 2021 IECC adoption and 2024-07-30 mandatory date. |
| amend-source:usa-nm:14-8-2 | Plumbing | 14.8.2 NMAC | adoption_and_transition | Includes 2021 UPC adoption and 2022-09-10 mandatory date. |
| amend-source:usa-nm:14-9-2 | Mechanical / Fuel Gas | 14.9.2 NMAC | adoption_and_transition | Includes 2021 UMC adoption and fuel-gas treatment within NMMC. |
| amend-source:usa-nm:14-10-4 | Electrical | 14.10.4 NMAC / New Mexico Register 2023 electrical package | adoption_and_transition | Includes 2020 NEC adoption and 2023-09-28 mandatory date. |
| amend-source:usa-nm:14-15-8 | Elevator / Conveyance | Title 14 Chapter 15 Elevator Safety Code / New Mexico Register 2025 Issue 15 | adoption_and_registration | Includes Elevator Safety Code parts 14.15.1 through 14.15.8 and the 2019 New Mexico Safety Code for Elevators and Escalators. |
| amend-source:usa-nm:10-25-5 | Fire | 10.25.5 NMAC / New Mexico Register 2022 fire amendment | adoption_and_authority_split | Includes fire-code adoption references, local fire-protection stringency language, and fire/building responsibility split. |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Amendment Topic | Summary | Source IDs | Confidence |
| --- | --- | --- | --- | --- | --- |
| amend:usa-nm:commercial-scope-2023 | Building | Scope and code family routing | 14.7.2 routes one- and two-family dwellings and townhouses to 14.7.3 and existing-building work to 14.7.2 or 14.7.7 as applicable. | src:usa-nm:nm-register-2023-building; src:usa-nm:nmac-14-7-2 | 0.84 |
| amend:usa-nm:energy-transition-2024 | Energy | Six-month code concurrency | 2021 residential and commercial energy codes used a transition period ending 2024-07-30. | src:usa-nm:nm-register-2024-energy | 0.86 |
| amend:usa-nm:electrical-transition-2023 | Electrical | Six-month code concurrency | 2020 NM Electrical Code used a transition period ending 2023-09-28. | src:usa-nm:nm-register-2023-electrical | 0.88 |
| amend:usa-nm:fire-building-split-2022 | Fire / Building | AHJ responsibility split | Construction-related fire-code sections for building-permit projects are assigned to the chief building official/AHJ, while fire-protection systems remain under fire-marshal authority. | src:usa-nm:nm-register-2022-fire-amend | 0.78 |
| amend:usa-nm:local-fire-stringency-2022 | Fire | Local requirements | Local fire-protection requirements may be more stringent when they do not conflict with the state fire rule. | src:usa-nm:nm-register-2022-fire-amend | 0.74 |
| amend:usa-nm:fuel-gas-ifgc-reference | Fuel Gas | IFGC cross-reference | References to the International Fuel Gas Code are treated as references to the New Mexico Mechanical Code or LP-gas standards rather than a standalone IFGC adoption in this pass. | src:usa-nm:nmac-14-9-2; src:usa-nm:nmac-14-7-6 | 0.68 |
| amend:usa-nm:elevator-code-2025 | Elevator / Conveyance | New elevator safety code chapter | Title 14 Chapter 15 rules add Elevator Safety Code general provisions, permitting, inspections, fees, licensing/certification, inspectors, and the 2019 New Mexico Safety Code for Elevators and Escalators. | src:usa-nm:rld-elevator-bureau; src:usa-nm:nm-register-2025-elevator-rules | 0.78 |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-nm"
  model: "hybrid_state_and_local_AHJ"
  enforcing_entities:
    - "Construction Industries Division where CID has jurisdiction"
    - "municipal, county, or political-subdivision AHJs with permitting/enforcement authority"
    - "chief building official / building AHJ for construction-code responsibilities on building-permit projects"
    - "fire marshal with authority over permitting for fire prevention and fire-protection systems"
  required_officials:
    - "AHJ / building official where local permitting authority exists"
    - "fire marshal where fire-code permitting/enforcement authority exists"
  state_reserved_activities:
    - "statewide construction-code adoption and minimum standards through CID/CIC"
    - "CID enforcement where work is subject to CID jurisdiction"
    - "fire-protection-system authority reserved to the fire marshal under the fire rule"
    - "elevator and commercial-conveyance registration, permitting, and inspection under CID Elevator Bureau rules"
  source_ids:
    - "src:usa-nm:cila-rld-2021"
    - "src:usa-nm:nmac-14-5-1"
    - "src:usa-nm:nmac-14-5-2"
    - "src:usa-nm:nm-register-2022-fire-amend"
  verification_status: "partially_verified"
  confidence: 0.66
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-nm"
  model: "state_minimum_with_local_more_stringent_fire_requirements_allowed_when_nonconflicting"
  applies_to_code_families:
    - "construction codes adopted and approved under the Construction Industries Licensing Act"
    - "fire-protection requirements under the New Mexico Fire Code"
  approval_required: null
  approving_authority_id: null
  filing_required: null
  registry_exists: null
  registry_source_ids: []
  legal_basis_source_ids:
    - "src:usa-nm:supp-nmsa-60-13-44-justia"
    - "src:usa-nm:nm-register-2022-fire-amend"
  verification_status: "partially_verified"
  confidence: 0.58
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement and local amendment authority should not be collapsed. The evidence supports local AHJ enforcement where local entities have permitting authority, but state-adopted construction codes remain minimum requirements for political subdivisions. For fire, local requirements may be more stringent when they do not conflict with state fire rules; however, this pass did not identify a statewide registry, filing workflow, or approval mechanism for local amendments.

### 6.4 Known Local Amendment Registries

| Registry ID | Registry Name | Code Families | Source IDs | Status | Notes |
| --- | --- | --- | --- | --- | --- |
| registry:usa-nm:local-amendments | statewide local amendment registry | unknown | none | not_found | No statewide local-amendment registry was confirmed in this pass. |

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendments were parsed. Albuquerque, Farmington, Santa Fe, Las Cruces, county, and fire-district amendments should be handled in separate local-jurisdiction records rather than inferred into the statewide report.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: partially_started

Jurisdiction stack:

```text
Address
  -> State of New Mexico
  -> County
  -> Municipality / unincorporated county
  -> Special districts, if applicable
  -> CID jurisdiction check
  -> Local building AHJ / chief building official, if applicable
  -> Local or state fire marshal / fire AHJ
  -> Trade-specific code authority under CID rules
  -> Applicable statewide code adoption records
  -> Applicable local enforcement record
  -> Applicable local amendment record, if independently sourced
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

No AHJ contact table was populated. A production resolver should separately collect CID regional office contacts, municipal building departments, county building departments, and fire-marshal/fire-department contacts.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Publisher | URL | Supports | Status |
| --- | --- | --- | --- | --- | --- |
| src:usa-nm:cila-rld-2021 | Construction Industries Licensing Act courtesy PDF | New Mexico Regulation and Licensing Department | https://www.rld.nm.gov/wp-content/uploads/2021/07/Article-13-CILA-7.1.21.pdf | CID/CIC/trade-bureau authority; statutory purpose and definitions | official_courtesy_copy |
| src:usa-nm:srca-title14-index | Title 14 Housing and Construction index | New Mexico State Records Center and Archives | https://www.srca.nm.gov/nmac-home/nmac-titles/title-14-housing-and-construction/ | Title 14 construction-code location and monitoring | official_index |
| src:usa-nm:rld-building-codes-page | CID Rules, Laws, and Building Codes page | New Mexico Regulation and Licensing Department | https://www.rld.nm.gov/construction-industries/rules-laws-and-building-codes/ | RLD/CID code resources and monitoring | official_agency_page |
| src:usa-nm:rld-elevator-bureau | Elevator Bureau page | New Mexico Regulation and Licensing Department | https://www.rld.nm.gov/construction-industries/find-a-bureau/bureaus/elevator-bureau/ | Elevator Safety Act, commercial conveyance registration, Elevator Safety Code 14.15.1 through 14.15.8, contact authority | official_agency_page |
| src:usa-nm:nm-register-2025-elevator-rules | New Mexico Register, Vol. XXXVI, Issue 15, adopted elevator rules | New Mexico State Records Center and Archives | https://www.srca.nm.gov/nmac-home/new-mexico-register/volume-xxxvi-issue-15/adopted-rules-issue-15/ | New Title 14 Chapter 15 elevator rules and 14.15.8 2019 New Mexico Safety Code for Elevators and Escalators | official_register_html |
| src:usa-nm:nmac-14-5-1 | 14.5.1 NMAC General Provisions | New Mexico State Records Center and Archives | https://www.srca.nm.gov/parts/title14/14.005.0001.html | AHJ definitions and Title 14 general provisions | official_rule_html |
| src:usa-nm:nmac-14-5-2 | 14.5.2 NMAC Permits | New Mexico State Records Center and Archives | https://www.srca.nm.gov/parts/title14/14.005.0002.html | Permit applicability and CID jurisdiction context | official_rule_html |
| src:usa-nm:nm-register-2023-building | New Mexico Register, Vol. XXXIV, Issue 11, June 13, 2023, building code package | New Mexico State Records Center and Archives | https://www.srca.nm.gov/nmac/nmregister/pdf/xxxiv11.pdf | 2021 commercial, residential, and existing-building code adoptions; effective dates; scope; grace periods | official_register_pdf |
| src:usa-nm:nmac-14-7-2 | 14.7.2 NMAC 2021 New Mexico Commercial Building Code | New Mexico State Records Center and Archives | https://www.srca.nm.gov/parts/title14/14.007.0002.html | Current commercial building code text; 2021 IBC adoption; amendment monitoring | official_rule_html |
| src:usa-nm:nmac-14-7-3 | 14.7.3 NMAC 2021 New Mexico Residential Building Code | New Mexico State Records Center and Archives | https://www.srca.nm.gov/parts/title14/14.007.0003.html | Current residential building code text; 2021 IRC adoption | official_rule_html |
| src:usa-nm:nmac-14-7-7 | 14.7.7 NMAC 2021 New Mexico Existing Building Code | New Mexico State Records Center and Archives | https://www.srca.nm.gov/parts/title14/14.007.0007.html | Current existing building code text; 2021 IEBC adoption | official_rule_html |
| src:usa-nm:nm-register-2024-14-7-2-amend | 14.7.2 NMAC amendment effective 2025-01-13 | New Mexico State Records Center and Archives | https://www.srca.nm.gov/nmac/nmregister/xxxv/14.7.2amend.html | Recent commercial-building amendment watch item | official_register_html |
| src:usa-nm:nm-register-2024-energy | New Mexico Register, Vol. XXXV, Issue 2, January 30, 2024, energy code package | New Mexico State Records Center and Archives | https://www.srca.nm.gov/nmac/nmregister/pdf/xxxv02.pdf | 2021 residential and commercial energy code adoptions; effective and mandatory dates | official_register_pdf |
| src:usa-nm:nmac-14-7-6 | 14.7.6 NMAC 2021 New Mexico Residential Energy Conservation Code | New Mexico State Records Center and Archives | https://www.srca.nm.gov/parts/title14/14.007.0006.html | Residential energy code; 2021 IECC adoption; cross-references | official_rule_html |
| src:usa-nm:nmac-14-7-9 | 14.7.9 NMAC New Mexico Commercial Energy Conservation Code | New Mexico State Records Center and Archives | https://www.srca.nm.gov/parts/title14/14.007.0009.html | Commercial energy code; 2021 IECC adoption; cross-references | official_rule_html |
| src:usa-nm:nmac-14-8-2 | 14.8.2 NMAC 2021 New Mexico Plumbing Code | New Mexico State Records Center and Archives | https://www.srca.nm.gov/parts/title14/14.008.0002.html | Plumbing code; 2021 UPC adoption; transition dates | official_rule_html |
| src:usa-nm:nmac-14-9-2 | 14.9.2 NMAC 2021 New Mexico Mechanical Code | New Mexico State Records Center and Archives | https://www.srca.nm.gov/parts/title14/14.009.0002.html | Mechanical code; 2021 UMC adoption; fuel-gas provisions | official_rule_html |
| src:usa-nm:nm-register-2023-electrical | New Mexico Register, Vol. XXXIV, Issue 4, February 28, 2023, electrical code package | New Mexico State Records Center and Archives | https://www.srca.nm.gov/nmac/nmregister/pdf/xxxiv04.pdf | 2020 electrical code adoption; effective and mandatory dates | official_register_pdf |
| src:usa-nm:nmac-14-10-4 | 14.10.4 NMAC 2020 New Mexico Electrical Code | New Mexico State Records Center and Archives | https://www.srca.nm.gov/parts/title14/14.010.0004.html | Electrical code; 2020 NEC adoption; administrative enforcement references | official_rule_html |
| src:usa-nm:nmac-10-25-5 | 10.25.5 NMAC Fire Prevention and Public Occupancy | New Mexico State Records Center and Archives | https://www.srca.nm.gov/parts/title10/10.025.0005.html | New Mexico Fire Code; IFC adoption and fire-code framework | official_rule_html |
| src:usa-nm:nm-register-2022-fire-amend | New Mexico Register, Vol. XXXIII, Issue 18, September 27, 2022, fire rule amendment | New Mexico State Records Center and Archives | https://www.srca.nm.gov/nmac/nmregister/pdf/xxxiii18.pdf | Fire effective date; local fire stringency; building/fire responsibility split | official_register_pdf |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-nm:cila-rld-2021 | courtesy_statute_copy | RLD PDF is an official agency-hosted convenience copy dated 2021-07-01, not a live codification endpoint. | Recheck against NMOneSource or the current statutory compilation before verified status. |
| src:usa-nm:nm-register-2023-building | pdf_extraction | Register PDF is official, but extracted text should be page-checked before line-level legal quoting. | Suitable for partially_verified adoption/date fields; page images needed for verified legal quote use. |
| src:usa-nm:nm-register-2024-energy | pdf_extraction | Register PDF is official, but extracted text should be page-checked before line-level legal quoting. | Suitable for partially_verified adoption/date fields; page images needed for verified legal quote use. |
| src:usa-nm:nm-register-2023-electrical | pdf_extraction | Register PDF is official, but extracted text should be page-checked before line-level legal quoting. | Suitable for partially_verified adoption/date fields; page images needed for verified legal quote use. |
| src:usa-nm:nm-register-2022-fire-amend | pdf_extraction | Register PDF is official, but fire-code edition text was crosschecked with supplemental rule mirrors. | Use official PDF plus official current NMAC page for production; verify exact IFC edition in page/PDF view. |
| src:usa-nm:nmac-14-7-2 | html_rule_currentness | Official current-rule page confirms the active part but search extraction may omit some subsection text. | Use official page as registry source; retain supplemental mirror only for subsection extraction caveat. |
| src:usa-nm:rld-elevator-bureau | current_program_page | Page states registration/permitting program status and links to rule/code pages, but does not itself provide line-by-line code text. | Use with New Mexico Register Issue 15 and NMAC Chapter 15 rule pages before verified status. |
| src:usa-nm:nm-register-2025-elevator-rules | effective_date_default | Register page states rules are effective on publication unless otherwise specified; individual Chapter 15 rule text should be checked for later section dates. | Use 2025-08-12 as a partially_verified effective date; page-check individual rules for verified status. |
| src:usa-nm:nmac-10-25-5 | html_rule_currentness | Official current-rule page confirms the active part but edition detail was crosschecked with Cornell mirror. | Verify exact current 10.25.5.8 text before raising fire row to verified. |

### 8.3 Supplemental Sources

| Source ID | Title | Publisher | URL | Supports | Caveat |
| --- | --- | --- | --- | --- | --- |
| src:usa-nm:supp-nmsa-60-13-44-justia | NMSA 1978, Section 60-13-44 mirror | Justia | https://law.justia.com/codes/new-mexico/chapter-60/article-13/section-60-13-44/ | Trade-bureau recommendations, political subdivisions subject to adopted codes as minimum requirements | Supplemental legal mirror; replace with official statute source before verified. |
| src:usa-nm:supp-cornell-14-7-2-38 | 14.7.2.38 NMAC mirror | Cornell Legal Information Institute | https://www.law.cornell.edu/regulations/new-mexico/N-M-Admin-Code-SS-14.7.2.38 | Elevator/conveying systems Chapter 30 cross-reference | Supplemental mirror; used because official page extraction did not isolate subsection 14.7.2.38. |
| src:usa-nm:supp-cornell-10-25-5-8 | 10.25.5.8 NMAC mirror | Cornell Legal Information Institute | https://www.law.cornell.edu/regulations/new-mexico/N-M-Admin-Code-SS-10.25.5.8 | IFC adoption and fire-code cross-references | Supplemental mirror; exact current official text should be checked before verified fire-code status. |

### 8.4 Source Extraction Metadata

| Extraction ID | Date | Extractor | Sources Covered | Method | Result |
| --- | --- | --- | --- | --- | --- |
| extract:usa-nm:2026-06-26:authority | 2026-06-26 | GPT-5.5 Thinking | CILA, NMSA 60-13-44 mirror, 14.5.1, 14.5.2 | official/source search and manual synthesis | Primary authority model populated with caveats. |
| extract:usa-nm:2026-06-26:adoptions | 2026-06-26 | GPT-5.5 Thinking | 14.7.2, 14.7.3, 14.7.7, 14.7.6, 14.7.9, 14.8.2, 14.9.2, 14.10.4, 14.15.1-14.15.8, 10.25.5 | official NMAC/Register source review; supplemental crosscheck for hard-to-extract subsections | Core code adoption matrix populated. |
| extract:usa-nm:2026-06-26:qa | 2026-06-26 | GPT-5.5 Thinking | Complete report | source-ID scan, marker scan, date-sequence sanity check | File is ready for narrow validation, not full verification. |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified_from_current_file | 1.00 | none | Status reflects that core authority and adoption rows are source-backed but not all code families are fully parsed. |
| ahj:usa-nm:cid | authority.name | Construction Industries Division, New Mexico Regulation and Licensing Department | partially_verified | 0.82 | src:usa-nm:cila-rld-2021; src:usa-nm:nm-register-2023-building | Official rules identify CID as issuing agency for core building codes. |
| ahj:usa-nm:cic | authority.role | Reviews and approves/disapproves rules, standards, and codes | partially_verified | 0.78 | src:usa-nm:cila-rld-2021 | Statutory source is an official courtesy copy; current statute recency should be rechecked. |
| adopt:usa-nm:commercial-building:2021-ibc | base_model_code | 2021 International Building Code | verified_from_sources_reviewed | 0.88 | src:usa-nm:nm-register-2023-building; src:usa-nm:nmac-14-7-2 | Adoption and effective date support found. |
| adopt:usa-nm:residential-building:2021-irc | mandatory_date | 2023-12-14 | verified_from_sources_reviewed | 0.86 | src:usa-nm:nm-register-2023-building; src:usa-nm:nmac-14-7-3 | Transition/grace-period support found. |
| adopt:usa-nm:existing-building:2021-iebc | mandatory_date | 2023-12-14 | partially_verified | 0.84 | src:usa-nm:nm-register-2023-building; src:usa-nm:nmac-14-7-7 | Existing-building transition support found, but full amendment delta not parsed. |
| adopt:usa-nm:plumbing:2021-upc | base_model_code | 2021 Uniform Plumbing Code | partially_verified | 0.76 | src:usa-nm:nmac-14-8-2 | Adoption support found; adoption-date field left null. |
| adopt:usa-nm:mechanical:2021-umc | base_model_code | 2021 Uniform Mechanical Code | partially_verified | 0.76 | src:usa-nm:nmac-14-9-2 | Adoption support found; adoption-date field left null. |
| adopt:usa-nm:electrical:2020-nec | mandatory_date | 2023-09-28 | verified_from_sources_reviewed | 0.88 | src:usa-nm:nm-register-2023-electrical; src:usa-nm:nmac-14-10-4 | Adoption, effective date, and mandatory date support found. |
| adopt:usa-nm:energy:2021-iecc | mandatory_date | 2024-07-30 | verified_from_sources_reviewed | 0.86 | src:usa-nm:nm-register-2024-energy; src:usa-nm:nmac-14-7-6; src:usa-nm:nmac-14-7-9 | Residential and commercial energy rules share transition structure. |
| adopt:usa-nm:elevator:2019-nmscee | state_code_name | 2019 New Mexico Safety Code for Elevators and Escalators | partially_verified | 0.78 | src:usa-nm:rld-elevator-bureau; src:usa-nm:nm-register-2025-elevator-rules | Current RLD page and New Mexico Register adopted-rules page support Chapter 15 elevator rules; individual rule text still needs page-checking. |
| adopt:usa-nm:fire:2021-ifc | base_model_code | 2021 International Fire Code | partially_verified | 0.78 | src:usa-nm:nmac-10-25-5; src:usa-nm:nm-register-2022-fire-amend; src:usa-nm:supp-cornell-10-25-5-8 | Current official rule and official amendment were reviewed; exact edition text relies on supplemental crosscheck. |
| local-amendment-rule:usa-nm | model | state minimum with more stringent non-conflicting local fire requirements | partially_verified | 0.58 | src:usa-nm:supp-nmsa-60-13-44-justia; src:usa-nm:nm-register-2022-fire-amend | Filing/approval workflow remains open. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Source-ID scan found body citations are present in section 8. |
| All authority IDs resolve | pass | Authority IDs used in sections 2, 3, and 6 are defined or named in section 2. |
| All current code families have adoption records | partial | Core construction, electrical, energy, elevator/conveyance, and fire rows are populated; accessibility remains scope-limited. |
| Building and operational fire code are separated | pass | Building/fire construction interfaces and operational fire prevention are separate rows. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Separate columns are retained; null values remain where adoption dates were not established. |
| Effective dates are valid ISO dates | pass | All populated date fields use YYYY-MM-DD format. |
| No impossible date sequences | pass | Mandatory dates are on or after effective/operative dates. |
| Transition rules have explicit trigger conditions | pass | Core code-family transition rules include trigger conditions. |
| Permit-date logic is captured where applicable | partial | Permit-application cutoff captured for commercial/building package and referenced generally; other code-family permit text may need direct rule-level review. |
| Local enforcement model classified | partial | Hybrid model identified; AHJ-by-address data not collected. |
| Local amendment rule classified | partial | Minimum-code and fire non-conflict posture captured; approval/filing/registry mechanics unresolved. |
| AHJ confirmation metadata present | fail | No local AHJ contact or boundary dataset was populated. |
| Official-source caveats captured | pass | Courtesy-copy, PDF extraction, and supplemental mirror caveats are explicit. |
| Template marker scan | pass | No template placeholders or prohibited marker phrases remain. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-nm:001 | medium | official statute currentness | CILA authority was supported by an official RLD courtesy PDF and a supplemental statute mirror, but not by a live official codification endpoint. | Replace or crosscheck with current NMOneSource/NM Legislature statute text before verified status. | null | null | open |
| issue:usa-nm:002 | medium | fire-code edition and local amendment workflow | The fire row is partly supported by supplemental extraction for 10.25.5.8 edition text; local fire amendment filing/approval workflow remains unresolved. | Page-check 10.25.5.8 and search SFM/local amendment guidance. | null | null | open |
| issue:usa-nm:003 | medium | accessibility scope | Accessibility authority is supported at the CILA level, but the current adopted accessibility standard and amendment text were not parsed. | Parse current NMCBC accessibility sections and any CID accessibility guidance. | null | null | open |
| issue:usa-nm:004 | low | elevator/conveyance rule details | CID Elevator Bureau and Chapter 15 rules were identified, but individual 14.15.1 through 14.15.8 rule text was not page-checked line by line. | Page-check each Chapter 15 rule and extract the incorporated elevator/escalator standards and any delayed operative dates. | null | null | open |
| issue:usa-nm:005 | medium | local amendment registry | No statewide local-amendment registry was found. | Confirm whether CID, SFM, or another state entity accepts or publishes local amendments. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-nm:title14-index | src:usa-nm:srca-title14-index | html_diff | monthly | New Title 14 construction-code part or updated existing part | 2026-06-26 | active |
| watch:usa-nm:rld-building-codes | src:usa-nm:rld-building-codes-page | html_diff | monthly | CID code page adds/replaces rule PDFs or code guidance | 2026-06-26 | active |
| watch:usa-nm:commercial-building | src:usa-nm:nmac-14-7-2 | html_diff | monthly | 14.7.2 amendment or replacement | 2026-06-26 | active |
| watch:usa-nm:energy | src:usa-nm:nmac-14-7-6; src:usa-nm:nmac-14-7-9 | html_diff | monthly | Energy-code amendment or replacement | 2026-06-26 | active |
| watch:usa-nm:electrical | src:usa-nm:nmac-14-10-4 | html_diff | monthly | NEC edition update or electrical-code amendment | 2026-06-26 | active |
| watch:usa-nm:elevator | src:usa-nm:rld-elevator-bureau; src:usa-nm:nm-register-2025-elevator-rules | html_diff | monthly | Elevator Safety Code amendment, registration/permitting program update, or incorporated standards change | 2026-06-26 | active |
| watch:usa-nm:fire | src:usa-nm:nmac-10-25-5 | html_diff | monthly | Fire-code amendment, IFC edition change, or SFM guidance update | 2026-06-26 | active |
| watch:usa-nm:register | src:usa-nm:srca-title14-index | register_review | monthly | New Mexico Register contains CID, CIC, RLD, or SFM code rulemaking notice | 2026-06-26 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-24 | Generated baseline draft report | report:usa-nm | none | Codex | Initial New Mexico stub contained explicit unresolved fields and no primary-source registry. |
| 2026-06-26 | Populated New Mexico state report with official-source-backed authority and code-adoption fields | report:usa-nm; ahj:usa-nm:cid; ahj:usa-nm:sfm; adopt:usa-nm:commercial-building:2021-ibc; adopt:usa-nm:residential-building:2021-irc; adopt:usa-nm:existing-building:2021-iebc; adopt:usa-nm:plumbing:2021-upc; adopt:usa-nm:mechanical:2021-umc; adopt:usa-nm:electrical:2020-nec; adopt:usa-nm:energy:2021-iecc; adopt:usa-nm:elevator:2019-nmscee; adopt:usa-nm:fire:2021-ifc | src:usa-nm:cila-rld-2021; src:usa-nm:nm-register-2023-building; src:usa-nm:nm-register-2023-electrical; src:usa-nm:nm-register-2024-energy; src:usa-nm:rld-elevator-bureau; src:usa-nm:nm-register-2025-elevator-rules; src:usa-nm:nm-register-2022-fire-amend | GPT-5.5 Thinking | Upgraded status to partially_verified after source registry, adoption matrix, transition rules, caveats, and QA rows were populated. |
