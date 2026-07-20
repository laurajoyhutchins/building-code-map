---
state:
  state_id: "US-SD"
  name: "South Dakota"
  abbreviation: "SD"
report:
  report_id: "state-report:usa-sd"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-26"
  last_verified: null
  reviewed_by: null
risk:
  overall_confidence: 0.62 # 0.00 - 1.00
  risk_flags:
    - "official_statute_portal_requires_js_extraction_caveat"
    - "2026_code_update_effective_date_requires_post_july_1_codification_check"
    - "mechanical_fuel_gas_existing_building_scope_partially_unresolved"
    - "local_amendment_inventory_not_collected"
    - "elevator_conveyance_authority_unresolved"
  open_questions_count: 5

---

# State Building Code Authority Report: South Dakota

## 1. Executive Summary

- **Authority model:** Hybrid statutory baseline with local administration. South Dakota does not appear to operate a single statewide building department for all construction. Title 11, chapter 10 establishes model-code baselines for local building and property-maintenance ordinances and a default nonresidential design standard when a local unit has not adopted an ordinance. Separate state programs administer electrical, plumbing, and fire-safety rules.

- **Statewide code status:** As of 2026-06-26, the currently verified Title 11 references are the 2021 IBC for nonresidential building ordinances/default nonresidential design standard, the 2021 IPMC for local property-maintenance ordinances, the 2021 IRC for municipal residential-code adoption, and the 2009 IECC as a voluntary standard for new residential buildings. Senate Bill 129 was signed in 2026 and updates these statutory references to the 2024 IBC, 2024 IPMC, and 2024 IRC; the default nonresidential trigger in § 11-10-6 changes to buildings commenced after 2026-07-01.

- **Local enforcement model:** Local units may adopt an ordinance allowing local administration and enforcement of the § 11-10-6 default nonresidential design standard. Local building-code enforcement is therefore not the same as statewide trade inspection: the Electrical Commission and Plumbing Commission administer statewide programs, with plumbing inspection exceptions for cities that have their own inspection requirements.

- **Local amendment posture:** Local units may amend, modify, or delete portions of the IBC before adoption and later amendments become effective when filed with the municipal finance officer or county auditor. Local units have similar IPMC amendment authority. Municipal IRC authority is narrower: municipalities may adopt/amend the IRC, but may not require residential sprinklers or impose requirements more stringent than the referenced IRC edition. Municipal fire regulations must be at least equal to State Fire Marshal standards and may be more stringent.

- **Known transition periods or pending changes:** 2026 SB 129 updates Title 11 references from 2021 to 2024 model codes. The bill text expressly changes the default nonresidential commencement trigger in § 11-10-6 from 2021-07-01 to 2026-07-01. The report should be rechecked after 2026-07-01 against official codified statutes.

- **Production readiness:** partially_ready_for_narrow_validation

### Key Findings

```yaml
---
key_findings:
- topic: State building-code model
  finding: "South Dakota uses a statutory/local-option model: local ordinances must\
    \ track specified I-Code editions, and \xA7 11-10-6 supplies a default nonresidential\
    \ design standard where no local ordinance exists."
  confidence: 0.78
  source_ids:
  - src:usa-sd:sdcl-11-10-5
  - src:usa-sd:sdcl-11-10-6
- topic: Primary building code edition
  finding: Current verified baseline remains 2021 IBC until the 2026 update is operative;
    SB 129 updates the statutory IBC reference to 2024 and changes the default commencement
    trigger to 2026-07-01.
  confidence: 0.75
  source_ids:
  - src:usa-sd:sdcl-11-10-5
  - src:usa-sd:sdcl-11-10-6
  - src:usa-sd:sb-129-2026
- topic: Residential code posture
  finding: Municipalities may adopt/amend the 2021 IRC, but may not require residential
    sprinklers or standards more stringent than the referenced IRC edition; SB 129
    updates the reference to 2024 IRC.
  confidence: 0.74
  source_ids:
  - src:usa-sd:sdcl-11-10-12
  - src:usa-sd:sb-129-2026
- topic: Electrical code authority
  finding: The South Dakota Electrical Commission administers electrical wiring laws/rules;
    permits received on or after 2024-11-12 are inspected under the 2023 NEC.
  confidence: 0.85
  source_ids:
  - src:usa-sd:dlr-electrical-home
  - src:usa-sd:arsd-20-44-22-01
- topic: Plumbing code authority
  finding: The South Dakota Plumbing Commission identifies SDCL 36-25 and ARSD 20:53/20:54
    as governing plumbing law/rules and states the 2024 UPC rule package was provisionally
    effective on 2025-10-28.
  confidence: 0.82
  source_ids:
  - src:usa-sd:dlr-plumbing-laws
  - src:usa-sd:dlr-plumbing-inspections
  - src:usa-sd:arsd-20-54
- topic: Fire code authority
  finding: The Department of Public Safety / State Fire Marshal fire-safety rules
    incorporate the 2015 IBC, 2015 IFC, and 2015 IMC, as modified in ARSD 61:15.
  confidence: 0.76
  source_ids:
  - src:usa-sd:arsd-61-15-01-01
  - src:usa-sd:dps-fire-laws
  - src:usa-sd:sdcl-34-29b-2
- topic: Local amendments
  finding: Local IBC/IPMC amendments are effective after adoption and filing with
    the local official named in statute; fire rules allow political subdivisions to
    enforce more stringent requirements.
  confidence: 0.76
  source_ids:
  - src:usa-sd:sdcl-11-10-5
  - src:usa-sd:sdcl-11-10-11
  - src:usa-sd:sdcl-34-29b-25
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-sd:statutory-local-building-code |
| Authority name | South Dakota Legislature / local units of government under SDCL Title 11, chapter 10 |
| Authority type | hybrid_statutory_baseline_and_local_government_adoption |
| Legal basis | SDCL §§ 11-10-5, 11-10-6, 11-10-11, 11-10-12; 2026 SB 129 for pending 2024-code updates |
| Role | Establishes required model-code editions for local building/property-maintenance ordinances, default nonresidential design standard where a local ordinance is absent, and municipal residential-code authority limits. |
| Enforcement model | Local administration/enforcement where a local unit adopts an ordinance; statewide enforcement not established for general building construction. |
| Source IDs | src:usa-sd:sdcl-11-10-5; src:usa-sd:sdcl-11-10-6; src:usa-sd:sdcl-11-10-11; src:usa-sd:sdcl-11-10-12; src:usa-sd:sb-129-2026 |
| Verification status | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-sd:statutory-local-building-code | Legislature / local units of government | Statutory model-code baseline and local adoption/enforcement; default nonresidential design standard where no local ordinance exists. | SDCL §§ 11-10-5, 11-10-6; 2026 SB 129 | src:usa-sd:sdcl-11-10-5; src:usa-sd:sdcl-11-10-6; src:usa-sd:sb-129-2026 | partially_verified |
| Residential | ahj:usa-sd:municipal-residential-code | Municipal governing bodies | Municipalities may enact residential requirements by adopting the IRC, subject to anti-sprinkler and no-more-stringent limits. | SDCL § 11-10-12; 2026 SB 129 | src:usa-sd:sdcl-11-10-12; src:usa-sd:sb-129-2026 | partially_verified |
| Existing Building / Rehabilitation | ahj:usa-sd:unknown-existing-building | Unresolved | Distinct existing-building or rehabilitation code adoption not established. | Not established | none | unresolved |
| Mechanical | ahj:usa-sd:dps-fire-marshal | Department of Public Safety / State Fire Marshal | Fire-safety rules incorporate the 2015 IMC; broader statewide mechanical-code administration outside fire-safety rules is unresolved. | ARSD 61:15:01:01; SDCL ch. 34-29B | src:usa-sd:arsd-61-15-01-01; src:usa-sd:sdcl-34-29b-2 | partially_verified |
| Plumbing | ahj:usa-sd:plumbing-commission | South Dakota Plumbing Commission | Administers plumbing statutes/rules and state plumbing inspections, with city inspection exceptions. | SDCL ch. 36-25; ARSD 20:53; ARSD 20:54 | src:usa-sd:dlr-plumbing-laws; src:usa-sd:dlr-plumbing-inspections; src:usa-sd:arsd-20-54 | partially_verified |
| Fuel Gas | ahj:usa-sd:unknown-fuel-gas | Unresolved | Distinct statewide fuel-gas code authority not established; possible local or fire/mechanical overlap requires review. | Not established | none | unresolved |
| Electrical | ahj:usa-sd:electrical-commission | South Dakota Electrical Commission | Administers electrical wiring laws/rules, inspects wiring installations, investigates complaints, and licenses electricians. | SDCL ch. 36-16; ARSD 20:44 | src:usa-sd:dlr-electrical-home; src:usa-sd:arsd-20-44; src:usa-sd:arsd-20-44-22-01 | verified_for_core_fields |
| Energy | ahj:usa-sd:statutory-local-building-code | State of South Dakota / Title 11 | 2009 IECC adopted as voluntary standard for new residential buildings. | SDCL § 11-10-7 | src:usa-sd:sdcl-11-10-7 | partially_verified |
| Fire - construction references | ahj:usa-sd:dps-fire-marshal | Department of Public Safety / State Fire Marshal | Fire-safety standards incorporate 2015 IBC, 2015 IFC, and 2015 IMC, as modified. | ARSD 61:15:01:01; SDCL ch. 34-29B | src:usa-sd:arsd-61-15-01-01; src:usa-sd:dps-fire-laws; src:usa-sd:sdcl-34-29b-2 | partially_verified |
| Fire - operational / prevention code | ahj:usa-sd:dps-fire-marshal | Department of Public Safety / State Fire Marshal | Fire prevention, investigation, training, education, fire-safety plan review for listed project types, and enforcement under chapter 34-29B. | SDCL ch. 34-29B; ARSD 61:15 | src:usa-sd:dps-fire-laws; src:usa-sd:sdcl-34-29b-2; src:usa-sd:sdcl-34-29b-3; src:usa-sd:arsd-61-15-01-01 | partially_verified |
| Accessibility | ahj:usa-sd:statutory-local-building-code | Local building AHJ / applicable IBC accessibility provisions | Accessibility applies through adopted IBC provisions where applicable; no separate statewide accessibility board/code was verified. | SDCL §§ 11-10-5, 11-10-6; ARSD 61:15:01:01 | src:usa-sd:sdcl-11-10-5; src:usa-sd:sdcl-11-10-6; src:usa-sd:arsd-61-15-01-01 | partially_verified |
| Elevator / Conveyance | ahj:usa-sd:unknown-elevator | Unresolved | Distinct statewide elevator/conveyance safety program not established from official sources reviewed. | Not established | none | unresolved |

### 2.3 Authority Hierarchy Notes

South Dakota's general construction-code framework is best classified as a statutory/local-option model rather than a centralized statewide building department. For nonresidential buildings, local units may adopt building-code ordinances that must comply with the referenced IBC edition, and a default statutory design standard applies where a local unit has not adopted an ordinance under § 11-10-5. Local enforcement is explicit only where a local unit adopts an ordinance allowing administration and enforcement of the default standard.

Trade and fire-safety authority is separate. The Electrical Commission and Plumbing Commission administer statewide trade rules, while the Department of Public Safety / State Fire Marshal administers fire-safety rules and specified plan-review/code-compliance programs. Local fire regulations may be more stringent than State Fire Marshal standards, but must be at least equal to those standards.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-sd:001 | ahj:usa-sd:statutory-local-building-code | authorizes_adoption_by | local units of government for nonresidential IBC ordinances | src:usa-sd:sdcl-11-10-5 | partially_verified |
| edge:usa-sd:002 | ahj:usa-sd:statutory-local-building-code | supplies_default_design_standard_for | nonresidential buildings in local units without § 11-10-5 ordinances | src:usa-sd:sdcl-11-10-6 | partially_verified |
| edge:usa-sd:003 | ahj:usa-sd:statutory-local-building-code | authorizes_local_administration_enforcement | local units adopting ordinance for § 11-10-6 design standard | src:usa-sd:sdcl-11-10-6 | partially_verified |
| edge:usa-sd:004 | ahj:usa-sd:municipal-residential-code | authorizes_adoption_by | municipalities for residential structures | src:usa-sd:sdcl-11-10-12 | partially_verified |
| edge:usa-sd:005 | ahj:usa-sd:electrical-commission | administers_and_inspects | electrical wiring installations statewide | src:usa-sd:dlr-electrical-home | verified_for_core_fields |
| edge:usa-sd:006 | ahj:usa-sd:plumbing-commission | administers_and_inspects | plumbing installations statewide, except cities with their own inspection requirements | src:usa-sd:dlr-plumbing-inspections | verified_for_core_fields |
| edge:usa-sd:007 | ahj:usa-sd:dps-fire-marshal | establishes_minimum_fire_safety_floor_for | municipal fire regulations | src:usa-sd:sdcl-34-29b-25 | partially_verified |
| edge:usa-sd:008 | ahj:usa-sd:board-technical-professions | convenes_workgroup_to_review | latest model national codes referenced in Title 11, chapter 10 | src:usa-sd:sdcl-11-10-13; src:usa-sd:btp-workgroup-2025 | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | SDCL Title 11 nonresidential construction standards / local building-code ordinance baseline | International Building Code | 2021 current; 2024 pending | current_until_2026-07-01_for_default_rule | null | null | 2021-07-01 for default nonresidential commencement trigger | 2021-07-01 for default nonresidential design standard where no local ordinance exists | § 11-10-6 applies to buildings commenced after 2021-07-01 where the local unit lacks a § 11-10-5 ordinance; SB 129 changes the referenced edition and commencement trigger to 2024 IBC / 2026-07-01. | src:usa-sd:sdcl-11-10-5; src:usa-sd:sdcl-11-10-6; src:usa-sd:sb-129-2026 |
| Residential | Municipal residential structure code authority | International Residential Code | 2021 current; 2024 pending | local_option_current_until_2026-07-01 | null | null | null | null | Municipality may adopt/amend IRC requirements, but may not require sprinklers or impose requirements more stringent than the referenced IRC edition; SB 129 updates the reference to 2024 IRC. | src:usa-sd:sdcl-11-10-12; src:usa-sd:sb-129-2026 |
| Existing Building / Rehabilitation | unresolved | unresolved | unresolved | unresolved | null | null | null | null | Distinct IEBC or rehabilitation-code adoption not established. | none |
| Mechanical | DPS fire-safety standards | International Mechanical Code | 2015 | fire_safety_current | null | null | null | null | ARSD 61:15:01:01 incorporates 2015 IMC for fire-safety standards; broader statewide mechanical scope unresolved. | src:usa-sd:arsd-61-15-01-01 |
| Plumbing | South Dakota State Plumbing Code | Uniform Plumbing Code | 2024 | current_with_provisional_effective_date_source | 2025-10-08 | 2025-10-28 | 2025-10-28 | 2025-10-28 | Plumbing Commission page states adopted amendments were served to the Secretary of State on 2025-10-08 and rules are provisionally effective on the 20th day after filing; inspections page states provisional effectiveness as of 2025-10-28. | src:usa-sd:dlr-plumbing-laws; src:usa-sd:dlr-plumbing-inspections; src:usa-sd:arsd-20-54 |
| Fuel Gas | unresolved | unresolved | unresolved | unresolved | null | null | null | null | Distinct statewide fuel-gas adoption not established; may require local ordinance and fire/mechanical-rule review. | none |
| Electrical | South Dakota electrical installations | NFPA 70 National Electrical Code | 2023 | current | null | 2024-11-12 | 2024-11-12 | 2024-11-12 for permits received on or after this date | Electrical Commission page states all permits received on or after 2024-11-12 are inspected using the 2023 NEC. | src:usa-sd:dlr-electrical-home; src:usa-sd:arsd-20-44-22-01 |
| Energy | Voluntary residential energy standard | International Energy Conservation Code | 2009 | voluntary_current | null | null | null | none | SDCL § 11-10-7 adopts the 2009 IECC as the voluntary standard applying to new residential buildings. | src:usa-sd:sdcl-11-10-7 |
| Fire - construction references | DPS fire-safety standards | International Building Code / International Mechanical Code / International Fire Code | 2015 | current_for_fire_safety_rules | null | null | null | null | ARSD 61:15:01:01 incorporates 2015 IBC, 2015 IFC, and 2015 IMC as fire-safety rules, except as modified in ARSD 61:15. | src:usa-sd:arsd-61-15-01-01 |
| Fire - operational / prevention code | DPS fire-safety standards | International Fire Code | 2015 | current | null | null | null | null | State Fire Marshal/DPS fire-safety rules incorporate the 2015 IFC, with local regulations allowed to be more stringent if at least equal to state standards. | src:usa-sd:arsd-61-15-01-01; src:usa-sd:sdcl-34-29b-25 |
| Accessibility | IBC accessibility provisions where IBC applies | International Building Code accessibility provisions | 2021 current; 2024 pending for Title 11 building baseline | partially_verified | null | null | 2021-07-01 for default nonresidential IBC baseline | context-dependent | Separate statewide accessibility code/board not verified; accessibility is captured only as embedded in IBC/Fire Marshal rules and should be checked with ADA/federal and local requirements. | src:usa-sd:sdcl-11-10-5; src:usa-sd:sdcl-11-10-6; src:usa-sd:arsd-61-15-01-01 |
| Elevator / Conveyance | unresolved | unresolved | unresolved | unresolved | null | null | null | null | Distinct statewide elevator/conveyance code adoption not established from official sources reviewed. | none |
| Property Maintenance | Local property-maintenance ordinance baseline | International Property Maintenance Code | 2021 current; 2024 pending | local_option_current_until_2026-07-01 | null | null | null | null | Local property-maintenance ordinances must comply with the referenced IPMC edition; SB 129 updates the reference to 2024 IPMC. | src:usa-sd:sdcl-11-10-11; src:usa-sd:sb-129-2026 |

### 3.2 Adoption Records

| Adoption ID | Code Family | Jurisdiction Scope | State Code Name | Base Model Code | Edition | Adoption Instrument | Adoption Date | Effective Date | Operative Date | Mandatory Date | Source IDs | Confidence | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| adoption:usa-sd:building-ibc-2021 | Building | statewide statutory baseline / local-option | Title 11 nonresidential construction standards | IBC | 2021 | SDCL §§ 11-10-5 and 11-10-6 | null | null | 2021-07-01 for § 11-10-6 default commencement trigger | 2021-07-01 where § 11-10-6 default applies | src:usa-sd:sdcl-11-10-5; src:usa-sd:sdcl-11-10-6 | 0.78 | Applies to buildings other than residential structures; exclusions include HUD-compliant manufactured homes, certain mobile/manufactured homes, farmstead structures, and specialty resort/vacation home establishments meeting stated criteria. |
| adoption:usa-sd:building-ibc-2024-pending | Building | statewide statutory baseline / local-option | 2026 SB 129 Title 11 updates | IBC | 2024 | 2026 SB 129 | 2026-03-04 signed by Governor | 2026-07-01 expected general effective date, codification recheck required | 2026-07-01 for § 11-10-6 default commencement trigger | 2026-07-01 where § 11-10-6 default applies | src:usa-sd:sb-129-2026 | 0.73 | Bill text changes § 11-10-6 default commencement trigger to buildings commenced after 2026-07-01. |
| adoption:usa-sd:residential-irc-2021 | Residential | municipal local-option | Municipal residential structure requirements | IRC | 2021 | SDCL § 11-10-12 | 2022 legislative source year | null | null | null | src:usa-sd:sdcl-11-10-12 | 0.74 | Municipalities may adopt/amend, but cannot require sprinklers or standards more stringent than the referenced IRC edition. |
| adoption:usa-sd:residential-irc-2024-pending | Residential | municipal local-option | 2026 SB 129 municipal residential update | IRC | 2024 | 2026 SB 129 | 2026-03-04 signed by Governor | 2026-07-01 expected general effective date, codification recheck required | 2026-07-01 expected | 2026-07-01 expected | src:usa-sd:sb-129-2026 | 0.70 | Bill text updates § 11-10-12 reference to the 2024 IRC. |
| adoption:usa-sd:property-maintenance-ipmc-2021 | Property Maintenance | local-option | Local property-maintenance ordinance baseline | IPMC | 2021 | SDCL § 11-10-11 | 2021 legislative source year | null | null | null | src:usa-sd:sdcl-11-10-11 | 0.75 | Applies only where a local unit adopts an ordinance prescribing standards for maintenance of existing structures and premises. |
| adoption:usa-sd:property-maintenance-ipmc-2024-pending | Property Maintenance | local-option | 2026 SB 129 property-maintenance update | IPMC | 2024 | 2026 SB 129 | 2026-03-04 signed by Governor | 2026-07-01 expected general effective date, codification recheck required | 2026-07-01 expected | 2026-07-01 expected | src:usa-sd:sb-129-2026 | 0.70 | Bill text updates § 11-10-11 reference to the 2024 IPMC. |
| adoption:usa-sd:energy-iecc-2009-voluntary | Energy | statewide voluntary residential standard | Voluntary residential energy standard | IECC | 2009 | SDCL § 11-10-7 | 2009 legislative source year | null | null | none | src:usa-sd:sdcl-11-10-7 | 0.82 | Voluntary standard for new residential buildings, not a mandatory statewide energy code. |
| adoption:usa-sd:electrical-nec-2023 | Electrical | statewide electrical installations | Electrical installations | NEC | 2023 | ARSD 20:44 and Electrical Commission implementation notice | null | 2024-11-12 | 2024-11-12 | 2024-11-12 for permits received on/after date | src:usa-sd:dlr-electrical-home; src:usa-sd:arsd-20-44-22-01 | 0.85 | Commission page states permits received on or after 2024-11-12 are inspected under the 2023 NEC. |
| adoption:usa-sd:plumbing-upc-2024 | Plumbing | statewide plumbing installations with local-inspection exceptions | South Dakota State Plumbing Code | UPC | 2024 | ARSD 20:54 / Plumbing Commission rule package | 2025-10-08 | 2025-10-28 | 2025-10-28 | 2025-10-28 | src:usa-sd:dlr-plumbing-laws; src:usa-sd:dlr-plumbing-inspections; src:usa-sd:arsd-20-54 | 0.82 | DLR page describes provisional effective date; production should verify final codified rule text after legislative review. |
| adoption:usa-sd:fire-ifc-2015 | Fire - operational / prevention code | statewide fire-safety rules | DPS fire-safety standards | IFC | 2015 | ARSD 61:15:01:01 | null | null | null | null | src:usa-sd:arsd-61-15-01-01 | 0.76 | 2015 IFC is part of fire-safety standards, except as modified, amended, or deleted in ARSD 61:15. |
| adoption:usa-sd:fire-ibc-imc-2015 | Fire - construction references | statewide fire-safety rules | DPS fire-safety standards | IBC / IMC | 2015 | ARSD 61:15:01:01 | null | null | null | null | src:usa-sd:arsd-61-15-01-01 | 0.74 | Captures Fire Marshal/DPS construction references only, not the separate Title 11 local building-code framework. |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

The most important verified transition rule is the nonresidential default design-standard trigger in SDCL § 11-10-6. The current statute applies the 2021 IBC default to covered buildings commenced after 2021-07-01 when the local unit has not adopted a § 11-10-5 ordinance. SB 129 changes that trigger to 2026-07-01 and the IBC edition to 2024. Electrical transition is permit-receipt based: permits received on or after 2024-11-12 are inspected under the 2023 NEC. Plumbing transition is rule-filing based: the 2024 UPC rule package was served to the Secretary of State on 2025-10-08 and described by the Plumbing Commission as provisionally effective on 2025-10-28.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-sd:001 | IBC default nonresidential design standard under § 11-10-6 | commencement_date | 2021-07-01 | Covered building commenced after this date in a local unit without a § 11-10-5 ordinance | unresolved | src:usa-sd:sdcl-11-10-6 | verified_current |
| date-rule:usa-sd:002 | 2024 IBC update under SB 129 | commencement_date | 2026-07-01 | Covered building commenced after this date in a local unit without a § 11-10-5 ordinance | unresolved | src:usa-sd:sb-129-2026 | pending_coding_recheck |
| date-rule:usa-sd:003 | 2023 NEC | permit_received_date | 2024-11-12 | Electrical permits received on or after this date are inspected using 2023 NEC | unresolved | src:usa-sd:dlr-electrical-home | verified_for_core_fields |
| date-rule:usa-sd:004 | 2024 UPC | filing_plus_20_days | 2025-10-28 | Adopted plumbing amendments served to Secretary of State on 2025-10-08; rules provisionally effective 20th day after filing | unresolved | src:usa-sd:dlr-plumbing-laws; src:usa-sd:dlr-plumbing-inspections | partially_verified |
| date-rule:usa-sd:005 | Local IBC/IPMC amendments | filing_date | date of adoption and local filing | Additional local amendments become effective upon adoption and filing with the municipal finance officer or county auditor | not applicable | src:usa-sd:sdcl-11-10-5; src:usa-sd:sdcl-11-10-11 | partially_verified |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | 2024 IBC | 2026-01-25 introduced | 2026-03-04 signed by Governor | 2026-07-01 expected general effective date; codification recheck required | 2026-07-01 for § 11-10-6 commencement trigger | 2026-07-01 where § 11-10-6 default applies | active | src:usa-sd:sb-129-2026 | Bill text changes IBC reference and default trigger. |
| Property Maintenance | 2024 IPMC | 2026-01-25 introduced | 2026-03-04 signed by Governor | 2026-07-01 expected general effective date; codification recheck required | 2026-07-01 expected | 2026-07-01 expected where local ordinance uses updated statute | active | src:usa-sd:sb-129-2026 | Bill text changes § 11-10-11 IPMC reference. |
| Residential | 2024 IRC | 2026-01-25 introduced | 2026-03-04 signed by Governor | 2026-07-01 expected general effective date; codification recheck required | 2026-07-01 expected | 2026-07-01 expected where municipality adopts under updated statute | active | src:usa-sd:sb-129-2026 | Bill text changes § 11-10-12 IRC reference. |
| Energy | 2009 IECC remains voluntary | null | null | null | null | none | monitor | src:usa-sd:sdcl-11-10-7; src:usa-sd:btp-workgroup-2025 | 2025 workgroup recommended continuing voluntary 2009 IECC approach. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-sd:001 | Building | Residential structures, HUD-compliant manufactured/mobile homes, farmsteads, specialty resort/vacation home establishments | Statutory exclusions in § 11-10-5 and § 11-10-6 | Certain building-code ordinance/default-standard provisions do not apply to the listed categories or are limited by HUD and Group R-3 compliance language. | src:usa-sd:sdcl-11-10-5; src:usa-sd:sdcl-11-10-6 | partially_verified |
| applicability-rule:usa-sd:002 | Residential | Detached one- and two-family dwellings, townhouses not more than three stories with separate means of egress, and accessory structures | Municipal IRC adoption under § 11-10-12 | Municipality may adopt IRC requirements for residential structures; cannot require sprinklers or requirements more stringent than the referenced IRC edition. | src:usa-sd:sdcl-11-10-12; src:usa-sd:sb-129-2026 | partially_verified |
| applicability-rule:usa-sd:003 | Plumbing | Plumbing installations | State plumbing inspection certificate requirement | Plumbing inspection certificate is required for plumbing installations, except for cities with their own inspection requirements. | src:usa-sd:dlr-plumbing-inspections | verified_for_core_fields |
| applicability-rule:usa-sd:004 | Fire | Schools, daycare facilities, flammable/combustible-liquid bulk/processing/retail motor fueling operations, LP bulk/fueling/filling operations | DPS/State Fire Marshal code-compliance plan review | DPS State Fire Laws page identifies ARSD 61:15 plan-document review for these construction-project categories. | src:usa-sd:dps-fire-laws | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** mixed statutory/local and agency-rule amendments.

**State amendment structure:** Title 11 statutory amendments are published in SDCL and session laws; local amendments are filed with the municipal finance officer or county auditor; electrical/plumbing/fire amendments are in ARSD articles 20:44, 20:54, and 61:15.

**Amendment parsing status:** partial_core_only

### 5.2 State Amendment Sources

| Amendment Source ID | Code Family | Publication Path | Source IDs | Parsing Status | Notes |
| --- | --- | --- | --- | --- | --- |
| amendment-source:usa-sd:title-11 | Building / Residential / Property Maintenance / Energy | SDCL Title 11, chapter 10; 2026 SB 129 | src:usa-sd:sdcl-11-10-5; src:usa-sd:sdcl-11-10-6; src:usa-sd:sdcl-11-10-7; src:usa-sd:sdcl-11-10-11; src:usa-sd:sdcl-11-10-12; src:usa-sd:sb-129-2026 | partial_core_only | Core adoption and local-amendment rules captured; detailed model-code amendment text not parsed. |
| amendment-source:usa-sd:electrical | Electrical | ARSD 20:44 and Electrical Commission notices | src:usa-sd:arsd-20-44; src:usa-sd:arsd-20-44-22-01; src:usa-sd:dlr-electrical-home | partial_core_only | NEC edition and effective permit rule captured; detailed exceptions not fully parsed. |
| amendment-source:usa-sd:plumbing | Plumbing | ARSD 20:54 and Plumbing Commission notices | src:usa-sd:arsd-20-54; src:usa-sd:dlr-plumbing-laws; src:usa-sd:dlr-plumbing-inspections | partial_core_only | 2024 UPC adoption/effective status captured; individual amendments not fully parsed. |
| amendment-source:usa-sd:fire | Fire / Mechanical construction references | ARSD 61:15 | src:usa-sd:arsd-61-15-01-01; src:usa-sd:dps-fire-laws | partial_core_only | Base fire-safety model-code editions captured; detailed deletions/amendments not fully parsed. |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Amendment Summary | Source IDs | Status | Notes |
| --- | --- | --- | --- | --- | --- |
| amendment:usa-sd:residential-sprinkler-limit | Residential | Municipalities may not require sprinklers in residential structures or impose requirements more stringent than the referenced IRC edition. | src:usa-sd:sdcl-11-10-12; src:usa-sd:sb-129-2026 | partially_verified | This is a statutory limit, not merely a local amendment. |
| amendment:usa-sd:local-ibc-amendment-filing | Building | Additional municipal IBC amendments are effective upon adoption and filing with the municipal finance officer; county amendments are effective upon adoption and filing with the county auditor. | src:usa-sd:sdcl-11-10-5 | partially_verified | Creates local filing-based effectiveness rule; no statewide registry found. |
| amendment:usa-sd:local-ipmc-amendment-filing | Property Maintenance | Additional municipal/county IPMC amendments are effective upon adoption and filing with the municipal finance officer or county auditor. | src:usa-sd:sdcl-11-10-11 | partially_verified | Creates local filing-based effectiveness rule; no statewide registry found. |
| amendment:usa-sd:fire-local-more-stringent | Fire | Municipal fire regulations must be at least equal to State Fire Marshal standards; political subdivisions may make and enforce more stringent requirements. | src:usa-sd:sdcl-34-29b-25 | partially_verified | Local fire-code posture differs from residential IRC no-more-stringent rule. |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-sd"
  model: "hybrid_local_building_enforcement_with_state_trade_and_fire_programs"
  enforcing_entities:
    - "Local units of government for building-code ordinances and any locally adopted administration/enforcement of the § 11-10-6 default design standard"
    - "South Dakota Electrical Commission / state and approved local electrical inspectors for electrical installations"
    - "South Dakota Plumbing Commission / state plumbing inspectors, except cities with their own plumbing inspection requirements"
    - "Department of Public Safety / State Fire Marshal for fire-safety rules and listed plan-review/code-compliance categories"
  required_officials:
    - "municipal finance officer for filing municipal IBC/IPMC amendments"
    - "county auditor for filing county IBC/IPMC amendments"
    - "state or approved local electrical inspector where applicable"
    - "state plumbing inspector or local city inspection office where applicable"
  state_reserved_activities:
    - "Electrical licensing/rules/inspection program administered by Electrical Commission"
    - "Plumbing licensing/rules/inspection program administered by Plumbing Commission"
    - "State Fire Marshal fire prevention, investigation, training, education, and ARSD 61:15 plan-review/code-compliance activities"
  source_ids:
    - "src:usa-sd:sdcl-11-10-5"
    - "src:usa-sd:sdcl-11-10-6"
    - "src:usa-sd:dlr-electrical-home"
    - "src:usa-sd:dlr-plumbing-inspections"
    - "src:usa-sd:dps-fire-laws"
  verification_status: "partially_verified"
  confidence: 0.72
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-sd"
  model: "local_amendments_allowed_with_code_family_specific_limits"
  applies_to_code_families:
    - "Building / IBC: local units may amend, modify, or delete before adoption; later amendments effective upon local filing"
    - "Property Maintenance / IPMC: local units may amend, modify, or delete before adoption; later amendments effective upon local filing"
    - "Residential / IRC: municipalities may amend, but may not require sprinklers or impose more stringent requirements than the referenced IRC edition"
    - "Fire: municipal standards must be at least equal to State Fire Marshal standards; more stringent political-subdivision requirements are allowed"
  approval_required: false
  approving_authority_id: null
  filing_required: true
  registry_exists: false
  registry_source_ids: []
  legal_basis_source_ids:
    - "src:usa-sd:sdcl-11-10-5"
    - "src:usa-sd:sdcl-11-10-11"
    - "src:usa-sd:sdcl-11-10-12"
    - "src:usa-sd:sdcl-34-29b-25"
  verification_status: "partially_verified"
  confidence: 0.74
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement and local amendment authority should not be collapsed. A local unit may have a building-code ordinance and local filing obligations for amendments, while statewide trade programs still govern electrical and plumbing requirements. Fire authority has a different rule: local fire regulations must be at least equal to State Fire Marshal standards and may be more stringent. Plumbing inspection is also split: state inspection certificates are required for plumbing installations, except cities listed by the Plumbing Commission as having their own inspection requirements.

### 6.4 Known Local Amendment Registries

No statewide local-amendment registry was verified. The statutes identify local filing locations rather than a centralized registry: municipal amendments are filed with the municipal finance officer and county amendments with the county auditor.

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendments were parsed in this pass. The report intentionally avoids using a single city code as a statewide proxy.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: partial_policy_model_only

Jurisdiction stack:

```text
Address
  -> State
  -> County
  -> Municipality / unincorporated county
  -> Local building-code ordinance status under SDCL § 11-10-5
  -> Default nonresidential design standard status under SDCL § 11-10-6, if no local ordinance
  -> Municipal residential-code status under SDCL § 11-10-12
  -> Local property-maintenance ordinance status under SDCL § 11-10-11
  -> State Electrical Commission jurisdiction / local electrical inspector status
  -> State Plumbing Commission jurisdiction / local city inspection-program exception status
  -> State Fire Marshal / local fire AHJ status
  -> Applicable local amendments and filing records
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

No AHJ contact dataset was populated. Plumbing inspection city exceptions identified from the Plumbing Commission page are Aberdeen, Brandon, Rapid City, Sioux Falls, and Spearfish; contact details for those local offices were not collected.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Publisher / Agency | URL | Source Type | Accessed | Supports |
| --- | --- | --- | --- | --- | --- | --- |
| src:usa-sd:sdcl-11-10-5 | SDCL § 11-10-5, New construction standards--Building code ordinance | South Dakota Legislature / Legislative Research Council | https://sdlegislature.gov/Statutes/11-10-5 | statute | 2026-06-26 | 2021 IBC local ordinance baseline; local amendment filing; exclusions |
| src:usa-sd:sdcl-11-10-6 | SDCL § 11-10-6, New construction standards--No building code ordinance | South Dakota Legislature / Legislative Research Council | https://sdlegislature.gov/Statutes/11-10-6 | statute | 2026-06-26 | 2021 IBC default design standard; local administration/enforcement; commencement trigger |
| src:usa-sd:sdcl-11-10-7 | SDCL § 11-10-7, Energy conservation code adopted as voluntary standard | South Dakota Legislature / Legislative Research Council | https://sdlegislature.gov/Statutes/11-10-7 | statute | 2026-06-26 | 2009 IECC voluntary residential standard |
| src:usa-sd:sdcl-11-10-11 | SDCL § 11-10-11, Property maintenance--Local ordinance--Required standards--Modifications | South Dakota Legislature / Legislative Research Council | https://sdlegislature.gov/Statutes/11-10-11 | statute | 2026-06-26 | 2021 IPMC local ordinance baseline; local amendment filing |
| src:usa-sd:sdcl-11-10-12 | SDCL § 11-10-12, Enactment of standards of International Residential Code | South Dakota Legislature / Legislative Research Council | https://sdlegislature.gov/Statutes/11-10-12 | statute | 2026-06-26 | Municipal IRC adoption; anti-sprinkler and no-more-stringent limits |
| src:usa-sd:sdcl-11-10-13 | SDCL § 11-10-13, Commission/workgroup of building codes | South Dakota Legislature / Legislative Research Council | https://sdlegislature.gov/Statutes/11-10-13 | statute | 2026-06-26 | Board of Technical Professions workgroup review cycle |
| src:usa-sd:sb-129-2026 | 2026 Senate Bill 129, update references to 2024 standard building codes | South Dakota Legislature / MyLRC | https://sdlegislature.gov/Session/Bill/27050 and https://mylrc.sdlegislature.gov/api/Documents/297175.pdf | enacted bill / official bill text | 2026-06-26 | Pending 2024 IBC/IPMC/IRC updates; 2026-07-01 default commencement trigger |
| src:usa-sd:btp-workgroup-2025 | State of South Dakota Summary Report, 2025 Building Codes Review | South Dakota Board of Technical Professions / Department of Labor and Regulation | https://dlr.sd.gov/btp/documents/BOTP-Workgroup-Summary-Report.pdf | official workgroup report | 2026-06-26 | Review cycle, referenced Title 11 model codes, recommendation context |
| src:usa-sd:dlr-electrical-home | South Dakota Electrical Commission homepage | South Dakota Department of Labor and Regulation | https://dlr.sd.gov/electrical/default.aspx | agency page | 2026-06-26 | Electrical Commission role; 2023 NEC permit transition |
| src:usa-sd:arsd-20-44 | ARSD Article 20:44, Electricians | South Dakota Legislature / Legislative Research Council | https://sdlegislature.gov/Rules/Administrative/20%3A44 | administrative rule | 2026-06-26 | Electrical rule structure |
| src:usa-sd:arsd-20-44-22-01 | ARSD 20:44:22:01, Adoption of National Electrical Code | South Dakota Legislature / Legislative Research Council | https://sdlegislature.gov/Rules/Administrative/20%3A44%3A22%3A01 | administrative rule | 2026-06-26 | NEC conformity rule |
| src:usa-sd:dlr-plumbing-laws | South Dakota Plumbing Commission, South Dakota Laws and Rules | South Dakota Department of Labor and Regulation | https://dlr.sd.gov/plumbing/laws.aspx | agency page | 2026-06-26 | Plumbing statutes/rules path; 2024 UPC adoption process and filing/effective rule |
| src:usa-sd:dlr-plumbing-inspections | South Dakota Plumbing Commission, Inspections | South Dakota Department of Labor and Regulation | https://dlr.sd.gov/plumbing/inspections.aspx | agency page | 2026-06-26 | 2024 UPC provisional effective date; inspection certificate requirement; city exceptions |
| src:usa-sd:arsd-20-54 | ARSD Article 20:54, State Plumbing Code | South Dakota Legislature / Legislative Research Council | https://sdlegislature.gov/Rules/Administrative/20%3A54 | administrative rule | 2026-06-26 | State plumbing code rules and 2024 UPC references |
| src:usa-sd:dps-fire-laws | DPS State Fire Laws page | South Dakota Department of Public Safety | https://www.sd.gov/dps?id=kb_article_view&sys_kb_id=9673c7194725aa90237fbd51026d43f0 | agency page | 2026-06-26 | State Fire Marshal program summaries and ARSD 61:15 plan-review scope |
| src:usa-sd:sdcl-34-29b-2 | SDCL § 34-29B-2, State Fire Marshal program and rules | South Dakota Legislature / Legislative Research Council | https://sdlegislature.gov/Statutes/34-29B-2 | statute | 2026-06-26 | Fire prevention, investigation, training, education program and rule authority |
| src:usa-sd:sdcl-34-29b-3 | SDCL § 34-29B-3, State Fire Marshal appointment and deputies | South Dakota Legislature / Legislative Research Council | https://sdlegislature.gov/Statutes/34-29B-3 | statute | 2026-06-26 | State Fire Marshal authority structure |
| src:usa-sd:sdcl-34-29b-25 | SDCL § 34-29B-25, Standards of municipal regulations | South Dakota Legislature / Legislative Research Council | https://sdlegislature.gov/Statutes/34-29B-25 | statute | 2026-06-26 | Municipal fire regulations at least equal to state standards; more stringent local requirements allowed |
| src:usa-sd:arsd-61-15-01-01 | ARSD 61:15:01:01, Conformity with national standards | South Dakota Legislature / Legislative Research Council | https://sdlegislature.gov/Rules/Administrative/61%3A15%3A01%3A01 | administrative rule | 2026-06-26 | 2015 IBC/IFC/IMC fire-safety standards |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-sd:sdcl-11-10-5 | js_rendering | Official SD Legislature page is JavaScript-heavy in this tool; exact extraction was cross-checked against indexed official snippets and a Justia mirror. | verify_with_official_api_or_browser_before_verified_status |
| src:usa-sd:sdcl-11-10-6 | js_rendering | Official SD Legislature page is JavaScript-heavy in this tool; exact extraction was cross-checked against indexed official snippets and a Justia mirror. | verify_with_official_api_or_browser_before_verified_status |
| src:usa-sd:sdcl-11-10-7 | js_rendering | Official SD Legislature page is JavaScript-heavy in this tool; exact extraction was cross-checked against indexed official snippets and a Justia mirror. | verify_with_official_api_or_browser_before_verified_status |
| src:usa-sd:sdcl-11-10-11 | js_rendering | Official SD Legislature page is JavaScript-heavy in this tool; exact extraction was cross-checked against indexed official snippets and a Justia mirror. | verify_with_official_api_or_browser_before_verified_status |
| src:usa-sd:sdcl-11-10-12 | js_rendering | Official SD Legislature page is JavaScript-heavy in this tool; exact extraction was cross-checked against indexed official snippets and a Justia mirror. | verify_with_official_api_or_browser_before_verified_status |
| src:usa-sd:sb-129-2026 | effective_date | Official bill text verifies code-edition updates and the 2026-07-01 commencement trigger in § 11-10-6; general act effective date should be confirmed after codification. | recheck_after_2026-07-01 |
| src:usa-sd:btp-workgroup-2025 | advisory_report | Workgroup report is advisory/recommendation context, not itself an adoption instrument. | use_for_monitoring_not_current_law |
| src:usa-sd:arsd-20-54 | recent_rule_update | DLR page describes 2024 UPC rules as provisionally effective; final codified rule text should be rechecked. | recheck_current_arsd_text |
| src:usa-sd:arsd-61-15-01-01 | courtesy_text_extraction | Official ARSD page is JavaScript-heavy; exact text extracted from Justia regulatory mirror and official search snippets. | verify_with_official_api_or_browser_before_verified_status |

### 8.3 Supplemental Sources

| Supplemental Source ID | Title | Publisher | URL | Use | Caveat |
| --- | --- | --- | --- | --- | --- |
| sup:usa-sd:justia-title-11 | 2025 South Dakota Codified Laws, Title 11, Chapter 10 sections | Justia | https://law.justia.com/codes/south-dakota/title-11/chapter-10/ | Text extraction mirror for statutory sections | Unofficial; official SD Legislature sources control. |
| sup:usa-sd:justia-arsd-61-15-01-01 | SD Administrative Rules 61:15:01:01 mirror | Justia Regulations | https://regulations.justia.com/states/south-dakota/title-61/article-61-15/chapter-61-15-01/section-61-15-01-01/ | Text extraction mirror for fire-safety code edition | Unofficial; official ARSD controls. |
| sup:usa-sd:legiscan-sb129 | SD SB129 status tracking | LegiScan | https://legiscan.com/SD/bill/SB129/2026 | Status confirmation that SB 129 was signed on 2026-03-04, with links to official South Dakota source documents | Supplemental only; official South Dakota bill page and MyLRC PDF control. |

### 8.4 Source Extraction Metadata

| Extraction ID | Source IDs | Method | Extracted Fields | Date | Notes |
| --- | --- | --- | --- | --- | --- |
| extract:usa-sd:001 | src:usa-sd:sdcl-11-10-5; src:usa-sd:sdcl-11-10-6; src:usa-sd:sdcl-11-10-7; src:usa-sd:sdcl-11-10-11; src:usa-sd:sdcl-11-10-12; src:usa-sd:sdcl-11-10-13 | web search/open, official snippets, Justia cross-check | Building/residential/property-maintenance/energy authority and current editions | 2026-06-26 | Official site rendered JavaScript fallback in tool. |
| extract:usa-sd:002 | src:usa-sd:sb-129-2026 | official bill page search result and MyLRC PDF text/screenshot | 2024 IBC/IPMC/IRC updates; 2026-07-01 commencement trigger | 2026-06-26 | Official bill PDF reviewed. |
| extract:usa-sd:003 | src:usa-sd:dlr-electrical-home; src:usa-sd:arsd-20-44-22-01 | agency page plus ARSD/search text | 2023 NEC effective permit rule; commission role | 2026-06-26 | ARSD exact rule text should be verified via official API/browser before verified status. |
| extract:usa-sd:004 | src:usa-sd:dlr-plumbing-laws; src:usa-sd:dlr-plumbing-inspections; src:usa-sd:arsd-20-54 | agency pages plus ARSD/search text | 2024 UPC adoption and effective dates; inspection model | 2026-06-26 | Recent rule update; final codified text recheck needed. |
| extract:usa-sd:005 | src:usa-sd:dps-fire-laws; src:usa-sd:arsd-61-15-01-01; src:usa-sd:sdcl-34-29b-2; src:usa-sd:sdcl-34-29b-3; src:usa-sd:sdcl-34-29b-25 | agency page, official snippets, Justia regulatory mirror | Fire Marshal authority, fire-safety model codes, local more-stringent rule | 2026-06-26 | Official ARSD page is JavaScript-heavy in this tool. |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 0.90 | none | Core authority/adoption fields populated from official-state and cross-checked sources, with explicit caveats. |
| report | risk.overall_confidence | 0.62 | verified | 0.80 | none | Confidence limited by JS-heavy official statute portal, recent 2026 update, and unresolved code families. |
| ahj:usa-sd:statutory-local-building-code | authority model | hybrid_statutory_baseline_and_local_government_adoption | partially_verified | 0.78 | src:usa-sd:sdcl-11-10-5; src:usa-sd:sdcl-11-10-6 | Captures Title 11 structure; local ordinance inventory not collected. |
| adoption:usa-sd:building-ibc-2021 | base model code / edition | 2021 IBC | partially_verified | 0.78 | src:usa-sd:sdcl-11-10-5; src:usa-sd:sdcl-11-10-6 | Current until pending 2026 update becomes operative/codified. |
| adoption:usa-sd:building-ibc-2024-pending | future edition / trigger | 2024 IBC; 2026-07-01 commencement trigger | partially_verified | 0.73 | src:usa-sd:sb-129-2026 | Requires post-2026-07-01 codified statute check. |
| adoption:usa-sd:electrical-nec-2023 | effective permit date | 2024-11-12 | verified_for_core_fields | 0.85 | src:usa-sd:dlr-electrical-home | Agency page states permit-receipt transition. |
| adoption:usa-sd:plumbing-upc-2024 | provisional effective date | 2025-10-28 | partially_verified | 0.82 | src:usa-sd:dlr-plumbing-laws; src:usa-sd:dlr-plumbing-inspections | Recent adoption; final codified ARSD text should be rechecked. |
| adoption:usa-sd:fire-ifc-2015 | base model code / edition | 2015 IFC | partially_verified | 0.76 | src:usa-sd:arsd-61-15-01-01 | Exact text extracted from unofficial mirror due JS official page; official source controls. |
| local-amendment-rule:usa-sd | filing_required | true | partially_verified | 0.76 | src:usa-sd:sdcl-11-10-5; src:usa-sd:sdcl-11-10-11 | Filing locations verified for IBC/IPMC local amendments; no statewide registry verified. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Every body `src:usa-sd:*` cited appears in Section 8. |
| All authority IDs resolve | pass | Authority IDs used in tables are introduced in Section 2 or explicitly unresolved. |
| All current code families have adoption records | pass | Rows are present; unresolved families remain explicit. |
| Building and operational fire code are separated | pass | Title 11 building baseline is separated from ARSD 61:15 fire-safety codes. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Date columns remain separate; null retained where not established. |
| Effective dates are valid ISO dates | pass | Dates entered use ISO format. |
| No impossible date sequences | pass | Known sequences align with sources; SB 129 future effective status is flagged for recheck. |
| Transition rules have explicit trigger conditions | pass | Commencement-date, permit-received-date, filing-plus-20-days, and local-filing triggers are explicit. |
| Permit-date logic is captured where applicable | pass | Electrical 2023 NEC transition uses permit-received date. |
| Local enforcement model classified | pass | Classified as hybrid local building enforcement with state trade and fire programs. |
| Local amendment rule classified | pass | Classified by code family with filing and stringency limits. |
| AHJ confirmation metadata present | fail | No address-level AHJ or local contact dataset populated. |
| Official-source caveats captured | pass | JS, advisory-report, recent-rule, and supplemental-source caveats included. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-sd:001 | high | 2026 code update codification | SB 129 is signed and updates Title 11 references to 2024 model codes, but report date is before the 2026-07-01 trigger. | Recheck official SDCL §§ 11-10-5, 11-10-6, 11-10-11, and 11-10-12 after 2026-07-01. | null | 2026-07-15 | open |
| issue:usa-sd:002 | high | mechanical / fuel gas / existing-building scope | ARSD 61:15 verifies 2015 IMC for fire-safety rules, but broader statewide mechanical, fuel-gas, and IEBC/rehabilitation authority remains unresolved. | Search official ARSD/SDCL and major local ordinances; do not infer statewide adoption from local codes. | null | null | open |
| issue:usa-sd:003 | medium | local amendment inventory | Title 11 identifies local filing locations, but no statewide registry was found. | Build local ordinance ingestion for municipalities/counties and capture filed amendments. | null | null | open |
| issue:usa-sd:004 | medium | elevator / conveyance authority | No distinct statewide elevator/conveyance safety program was verified from official sources reviewed. | Confirm with DPS/Fire Marshal, DLR, and ARSD whether any statewide conveyance program exists. | null | null | open |
| issue:usa-sd:005 | medium | AHJ and boundary resolver | Address-level building/fire/trade AHJ resolution has not been implemented. | Select boundary datasets and AHJ contacts; map plumbing city exceptions and local building-code adoption status. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-sd:title-11 | src:usa-sd:sdcl-11-10-5 | statute_diff | monthly_until_2026-08_then_quarterly | codified 2024 IBC/IPMC/IRC update appears or statute text changes | 2026-06-26 | active |
| watch:usa-sd:sb129 | src:usa-sd:sb-129-2026 | bill_status_and_session_law_check | weekly_until_2026-07-15 | session-law/codified version confirms effective date and final text | 2026-06-26 | active |
| watch:usa-sd:electrical | src:usa-sd:dlr-electrical-home | html_diff | monthly | NEC edition, permit-transition rule, or exceptions change | 2026-06-26 | active |
| watch:usa-sd:plumbing | src:usa-sd:dlr-plumbing-laws | html_diff | monthly | UPC edition, effective date, or rule status changes | 2026-06-26 | active |
| watch:usa-sd:fire | src:usa-sd:arsd-61-15-01-01 | rule_diff | quarterly | IFC/IBC/IMC fire-safety edition or amendments change | 2026-06-26 | active |
| watch:usa-sd:btp | src:usa-sd:btp-workgroup-2025 | report_check | annual | new model building code workgroup report or recommendation appears | 2026-06-26 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-24 | Generated baseline draft report | report:usa-sd | none | Codex | Initial state-specific skeleton contained unresolved placeholders. |
| 2026-06-26 | Populated core authority, adoption, transition, source-registry, and QA sections | ahj:usa-sd:statutory-local-building-code; adoption:usa-sd:building-ibc-2021; adoption:usa-sd:electrical-nec-2023; adoption:usa-sd:plumbing-upc-2024; adoption:usa-sd:fire-ifc-2015 | src:usa-sd:sdcl-11-10-5; src:usa-sd:sdcl-11-10-6; src:usa-sd:sdcl-11-10-7; src:usa-sd:sdcl-11-10-11; src:usa-sd:sdcl-11-10-12; src:usa-sd:sb-129-2026; src:usa-sd:dlr-electrical-home; src:usa-sd:dlr-plumbing-laws; src:usa-sd:dlr-plumbing-inspections; src:usa-sd:arsd-61-15-01-01 | ChatGPT | Upgraded status to partially_verified with caveats and open issues. |
