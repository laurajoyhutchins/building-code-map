---
state:
  state_id: "US-IL"
  name: "Illinois"
  abbreviation: "IL"
report:
  report_id: "state-report:usa-il"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-25"
  last_verified: null
  reviewed_by: null
risk:
  overall_confidence: 0.68 # 0.00 - 1.00
  risk_flags:
    - "hybrid_state_local_authority_model"
    - "local_code_directory_not_extracted"
    - "fire_life_safety_local_equivalency_exceptions_need_project_level_review"
    - "private_mechanical_and_fuel_gas_scope_not_fully_resolved"
    - "chicago_and_article_34_school_exceptions_need_project_level_review"
  open_questions_count: 6

---

# State Building Code Authority Report: Illinois

## 1. Executive Summary

- **Authority model:** Illinois uses a hybrid state/local model. The Illinois Capital Development Board (CDB) is the primary statewide building-code coordinating authority for the post-2025 statewide framework, local building-code reporting, state-funded buildings, and commercial buildings in non-building-code jurisdictions. Statewide specialty code authority is split among CDB for energy and accessibility, the Office of the State Fire Marshal (OSFM) for fire/life-safety and elevators, the Illinois Department of Public Health (IDPH) for plumbing, and the Illinois State Board of Education (ISBE) for public schools.

- **Statewide code status:** Beginning January 1, 2025, Illinois requires statewide building-code baselines, but the framework is not a single uniform statewide building code for every private project. Commercial buildings in non-building-code jurisdictions and state-funded buildings must use current or most-recent-preceding IBC/IEBC and NEC plus Illinois energy, stretch-energy where applicable, accessibility, plumbing, and fire-prevention rules. Municipal and county building codes remain operative, but their structural provisions must meet statutory baseline-stringency requirements beginning January 1, 2025.

- **Local enforcement model:** Local building officials remain the primary first-line AHJs where a local code exists. Commercial projects in non-building-code jurisdictions rely on qualified-inspector certification filed with the municipality or county. Fire, plumbing, energy, accessibility, public-school, and elevator rules have separate state or shared enforcement channels.

- **Local amendment posture:** Local building codes and amendments are allowed, but municipalities and counties must report adopted model code titles/editions and amendments to CDB. Beginning January 1, 2025, local structural provisions must satisfy statutory baseline stringency. Energy, fire/life-safety, plumbing, accessibility, public-school, and elevator rules impose separate state constraints, and local fire or elevator rules may be recognized only under their respective state-rule conditions.

- **Known transition periods or pending changes:** Key transition rules include the January 1, 2025 statewide building-code baseline, the permit-application-year rule for commercial non-building-code jurisdictions, the January 1, 2025 public-school design-contract trigger for 2024 ICC codes, and the November 30, 2025 effective date for the 2024 Illinois Energy Conservation Code. Future stretch-energy-code update milestones are statutory watch items for December 31, 2026, December 31, 2029, and December 31, 2032.

- **Production readiness:** partially_ready_for_narrow_validation

### Key Findings

```yaml
---
key_findings:
- topic: State adopting / coordinating authority
  finding: CDB is the primary statewide building-code coordinating authority for the
    post-2025 framework, commercial non-building-code certification, and local building-code
    reporting.
  confidence: 0.82
  source_ids:
  - src:usa-il:cdb-act-20-ilcs-3105
  - src:usa-il:cdb-building-codes-page
- topic: Primary private-building model
  finding: 'Illinois is hybrid: local codes continue where adopted, but post-2025
    statutory baselines and state specialty codes apply.'
  confidence: 0.78
  source_ids:
  - src:usa-il:cdb-act-20-ilcs-3105
  - src:usa-il:cdb-building-codes-page
- topic: Commercial non-building-code baseline
  finding: New or substantially improved commercial buildings in non-building-code
    jurisdictions must certify compliance with current or most-recent-preceding IBC/IEBC
    and NEC, plus Illinois energy/stretch, accessibility, plumbing, and fire-prevention
    rules.
  confidence: 0.85
  source_ids:
  - src:usa-il:cdb-act-20-ilcs-3105
- topic: Residential non-building-code baseline
  finding: In non-building-code jurisdictions, new residential construction contracts
    must adopt an agreed IRC or nearby municipal/county residential code; if no agreement/code
    is stated, the current IRC plus Illinois energy and plumbing codes are adopted
    by law into the contract.
  confidence: 0.82
  source_ids:
  - src:usa-il:residential-building-code-act-815-ilcs-670
- topic: Energy code
  finding: CDB administers the Illinois Energy Conservation Code; the 2024 IECC, as
    amended, is incorporated for privately funded commercial facilities and residential
    buildings, effective 2025-11-30.
  confidence: 0.86
  source_ids:
  - src:usa-il:energy-code-71-iac-600
  - src:usa-il:energy-code-cdb-page
- topic: Fire / life-safety code
  finding: OSFM administers Part 100 Fire Prevention and Safety under the Fire Investigation
    Act and adopts NFPA 101 (2015) subject to OSFM modifications; local AHJs may use
    equal-or-higher local fire/life-safety codes when OSFM criteria are met.
  confidence: 0.8
  source_ids:
  - src:usa-il:fire-prevention-rules-41-iac-100
  - src:usa-il:osfm-life-safety-faq
- topic: Public schools
  finding: Public school facilities outside Article 34 Chicago-school exceptions are
    governed by the Health/Life Safety Code for Public Schools; 2024 IBC and listed
    subcodes apply to projects with design contracts executed on or after 2025-01-01.
  confidence: 0.82
  source_ids:
  - src:usa-il:school-code-105-ilcs-5-2-3-12
  - src:usa-il:hls-code-public-schools-23-iac-180
- topic: Local amendment registry
  finding: Municipalities and counties must report new building-code editions and
    amendments to CDB before effective date; CDB must publish reported code title,
    edition, amendment status, and reporting date.
  confidence: 0.8
  source_ids:
  - src:usa-il:cdb-act-20-ilcs-3105
  - src:usa-il:cdb-building-codes-page
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-il:cdb |
| Authority name | Illinois Capital Development Board |
| Authority type | state board / statewide building-code coordinating authority |
| Legal basis | Capital Development Board Act, especially 20 ILCS 3105/10.09-1 and 20 ILCS 3105/10.18 |
| Role | Coordinates statewide construction-code framework; identifies and publishes reported local building codes; establishes commercial non-building-code certification requirements; administers Illinois energy code rulemaking and accessibility-code rulemaking; applies stretch-energy requirements to CDB-authorized or CDB-funded projects. |
| Enforcement model | hybrid_state_local_specialty: local AHJs enforce local building codes; non-building-code commercial projects use qualified-inspector certification; specialty codes are enforced by state or shared authorities. |
| Source IDs | src:usa-il:cdb-act-20-ilcs-3105; src:usa-il:cdb-building-codes-page; src:usa-il:energy-code-71-iac-600; src:usa-il:accessibility-code-71-iac-400 |
| Verification status | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-il:cdb; ahj:usa-il:local-building-ahj | Illinois Capital Development Board; municipalities/counties | CDB sets statewide baseline/reporting and non-building-code commercial certification requirements; local AHJs administer local adopted building codes. | 20 ILCS 3105/10.09-1; 20 ILCS 3105/10.18 | src:usa-il:cdb-act-20-ilcs-3105; src:usa-il:cdb-building-codes-page | partially_verified |
| Residential | ahj:usa-il:local-building-ahj; ahj:usa-il:cdb | Municipalities/counties; CDB reporting role | Local residential codes may apply; non-building-code jurisdictions use the Residential Building Code Act contract/default mechanism; local structural provisions must satisfy post-2025 baseline rule. | 815 ILCS 670/15; 20 ILCS 3105/10.18 | src:usa-il:residential-building-code-act-815-ilcs-670; src:usa-il:cdb-act-20-ilcs-3105 | partially_verified |
| Existing Building / Rehabilitation | ahj:usa-il:cdb; ahj:usa-il:local-building-ahj | CDB; municipalities/counties | IEBC applies to non-building-code commercial certification; local rehabilitation codes must meet post-2025 baseline structural requirements for existing nonresidential buildings. | 20 ILCS 3105/10.09-1; 20 ILCS 3105/10.18 | src:usa-il:cdb-act-20-ilcs-3105 | partially_verified |
| Mechanical | ahj:usa-il:local-building-ahj; ahj:usa-il:isbe | Local AHJs; ISBE for public schools | Private statewide mechanical-code adoption outside local codes was not fully resolved. Public school projects use 2024 IMC as a subcode of the Health/Life Safety Code for applicable post-2025 design contracts. | 23 Ill. Adm. Code 180.60 | src:usa-il:hls-code-public-schools-23-iac-180 | unresolved_private_scope |
| Plumbing | ahj:usa-il:idph; ahj:usa-il:local-plumbing-ahj | Illinois Department of Public Health; local plumbing inspectors/governmental units | IDPH promulgates Illinois Plumbing Code and retains variance authority; local governmental units may administer permits and inspections through licensed plumbing inspectors. | 225 ILCS 320/35; 225 ILCS 320/37; 77 Ill. Adm. Code 890 | src:usa-il:plumbing-code-77-iac-890; src:usa-il:plumbing-license-law-225-ilcs-320 | partially_verified |
| Fuel Gas | ahj:usa-il:local-building-ahj; ahj:usa-il:isbe | Local AHJs; ISBE for public schools | Public school projects use 2024 IFGC for applicable post-2025 design contracts. Private statewide fuel-gas code authority outside local codes was not fully resolved. | 23 Ill. Adm. Code 180.60 | src:usa-il:hls-code-public-schools-23-iac-180 | unresolved_private_scope |
| Electrical | ahj:usa-il:cdb; ahj:usa-il:local-building-ahj | CDB; local AHJs | Commercial non-building-code certification must include current or most-recent-preceding NEC; local permitting/enforcement details remain jurisdiction-specific. | 20 ILCS 3105/10.09-1 | src:usa-il:cdb-act-20-ilcs-3105 | partially_verified |
| Energy | ahj:usa-il:cdb | Illinois Capital Development Board | CDB adopts/administers Illinois Energy Conservation Code and Illinois Stretch Energy Code. | 20 ILCS 3125/15; 20 ILCS 3125/55; 71 Ill. Adm. Code 600 | src:usa-il:energy-efficient-building-act-20-ilcs-3125; src:usa-il:energy-code-71-iac-600; src:usa-il:energy-code-cdb-page | partially_verified |
| Fire - construction references | ahj:usa-il:osfm; ahj:usa-il:isbe; ahj:usa-il:local-fire-ahj | OSFM; ISBE for public schools; local fire AHJs | Commercial non-building-code certification includes Fire Investigation Act rules; public schools use 2024 IFC excluding Chapter 4; local fire/life-safety codes may be recognized if equal or higher. | 20 ILCS 3105/10.09-1; 41 Ill. Adm. Code 100; 23 Ill. Adm. Code 180.60 | src:usa-il:cdb-act-20-ilcs-3105; src:usa-il:fire-prevention-rules-41-iac-100; src:usa-il:hls-code-public-schools-23-iac-180 | partially_verified |
| Fire - operational / prevention code | ahj:usa-il:osfm; ahj:usa-il:local-fire-ahj | Office of the State Fire Marshal; local fire AHJs | OSFM administers Fire Prevention and Safety Rules, including NFPA 101 (2015) with modifications; local AHJs enforce under OSFM direction unless equal-or-higher local standards apply. | Fire Investigation Act, 425 ILCS 25/9; 41 Ill. Adm. Code 100 | src:usa-il:fire-prevention-rules-41-iac-100; src:usa-il:osfm-life-safety-faq | partially_verified |
| Accessibility | ahj:usa-il:cdb; ahj:usa-il:il-ag | CDB; Illinois Attorney General | CDB adopts Illinois Accessibility Code and issues project-specific interpretations; Illinois Attorney General enforces the Code under the Environmental Barriers Act. | 410 ILCS 25/4; 71 Ill. Adm. Code 400 | src:usa-il:environmental-barriers-act-410-ilcs-25; src:usa-il:accessibility-code-71-iac-400 | partially_verified |
| Elevator / Conveyance | ahj:usa-il:osfm-elevator; ahj:usa-il:elevator-review-board; ahj:usa-il:local-elevator-admin | OSFM Elevator Safety Division; Elevator Safety Review Board; local elevator administrators by agreement | OSFM implements elevator/conveyance registration, inspection, certification, and licensing outside Chicago; municipalities/counties may enter local elevator agreements, with standards at least as stringent as state rules. | Elevator Safety and Regulation Act, 225 ILCS 312; 41 Ill. Adm. Code 1000 | src:usa-il:elevator-safety-rules-41-iac-1000; src:usa-il:elevator-act-225-ilcs-312; src:usa-il:osfm-elevator-page | partially_verified |

### 2.3 Authority Hierarchy Notes

Illinois should not be modeled as a single-code state. The post-2025 statewide framework overlays local codes with minimum structural requirements, non-building-code commercial certification, and statewide specialty codes. A project-level AHJ resolver must first determine whether a municipality or county has reported/adopted a building code, then layer in statewide energy, plumbing, accessibility, fire/life-safety, school, elevator, and other specialty rules. Public schools, state-funded work, buildings in non-building-code jurisdictions, Chicago elevator work, and Article 34 school-district work require special-case routing.

Local enforcement and local amendment authority are separate. Local jurisdictions may administer and amend local building codes, but those amendments must be reported to CDB and cannot undercut statutory structural baseline rules beginning January 1, 2025. Separate specialty codes have their own state-law constraints and may not be overridden merely because a local building department enforces permits.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-il:001 | ahj:usa-il:cdb | publishes_reported_codes_for | municipalities_and_counties | src:usa-il:cdb-act-20-ilcs-3105 | partially_verified |
| edge:usa-il:002 | municipalities_and_counties | must_report_code_editions_and_amendments_to | ahj:usa-il:cdb | src:usa-il:cdb-act-20-ilcs-3105; src:usa-il:cdb-building-codes-page | partially_verified |
| edge:usa-il:003 | ahj:usa-il:cdb | sets_minimum_structural_baseline_for | municipal_and_county_building_codes_beginning_2025-01-01 | src:usa-il:cdb-act-20-ilcs-3105 | partially_verified |
| edge:usa-il:004 | property_owner_in_noncode_commercial_scope | obtains_qualified_inspection_certification_filed_with | municipality_or_county | src:usa-il:cdb-act-20-ilcs-3105 | partially_verified |
| edge:usa-il:005 | ahj:usa-il:osfm | directs_or_shares_fire_prevention_enforcement_with | local_fire_AHJs | src:usa-il:fire-prevention-rules-41-iac-100 | partially_verified |
| edge:usa-il:006 | local_fire_AHJs | may_use_equal_or_higher_fire_life_safety_code_if_criteria_met_by | ahj:usa-il:osfm | src:usa-il:fire-prevention-rules-41-iac-100 | partially_verified |
| edge:usa-il:007 | ahj:usa-il:idph | retains_variance_authority_for | Illinois_Plumbing_Code | src:usa-il:plumbing-code-77-iac-890 | partially_verified |
| edge:usa-il:008 | authorized_local_governmental_units | may_administer_plumbing_permits_and_inspections_through | licensed_plumbing_inspectors | src:usa-il:plumbing-license-law-225-ilcs-320 | partially_verified |
| edge:usa-il:009 | ahj:usa-il:isbe | governs_health_life_safety_code_for | public_schools_outside_Article_34_limited_exception | src:usa-il:school-code-105-ilcs-5-2-3-12; src:usa-il:hls-code-public-schools-23-iac-180 | partially_verified |
| edge:usa-il:010 | ahj:usa-il:osfm-elevator | may_enter_local_elevator_agreements_with | municipalities_or_counties | src:usa-il:elevator-act-225-ilcs-312 | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | Illinois commercial non-building-code certification framework / local structural baseline | IBC | Current or most recent preceding for non-building-code commercial; baseline IBC is edition first published in current or preceding 9 years with least-restrictive structural provisions | mandatory_limited_scope | 2024-01-01 | 2025-01-01 | 2025-01-01 | 2025-01-01 | In non-building-code commercial jurisdictions, permit-application-year January 1 requirements remain for permit duration; local structural baseline begins 2025-01-01. | src:usa-il:cdb-act-20-ilcs-3105; src:usa-il:cdb-building-codes-page |
| Residential | Illinois residential non-building-code contract/default framework / local structural baseline | IRC | Current or most recent preceding IRC where agreed; current IRC if no agreed/stated code; excludes Parts IV and VII | mandatory_limited_scope | 2024-01-01 | 2024-01-01 | 2024-01-01 | 2025-01-01 for local structural baseline | In non-building-code jurisdictions, construction contract controls where valid; absent agreement, IRC/current energy/plumbing defaults apply by law. | src:usa-il:residential-building-code-act-815-ilcs-670; src:usa-il:cdb-act-20-ilcs-3105 |
| Existing Building / Rehabilitation | Illinois existing-building non-building-code certification framework / local structural baseline | IEBC | Current or most recent preceding for non-building-code commercial; baseline IEBC is edition first published in current or preceding 9 years with least-restrictive structural provisions | mandatory_limited_scope | 2024-01-01 | 2025-01-01 | 2025-01-01 | 2025-01-01 | Permit-application-year January 1 rule for non-building-code commercial; local structural baseline for existing nonresidential work begins 2025-01-01. | src:usa-il:cdb-act-20-ilcs-3105 |
| Mechanical | Health/Life Safety Code for Public Schools mechanical subcode | IMC | 2024 | mandatory_limited_public_school_scope | 2025-01-30 | 2025-01-30 | 2025-01-01 | 2025-01-01 | Applies to public-school projects with design contracts executed on or after 2025-01-01; private statewide mechanical-code scope outside local codes remains unresolved. | src:usa-il:hls-code-public-schools-23-iac-180 |
| Plumbing | Illinois Plumbing Code | State plumbing code with incorporated standards | 77 Ill. Adm. Code 890; latest located amendment effective 2014-04-24 | mandatory_statewide_scope_for_plumbing_work | null | 2014-04-24 | 2014-04-24 | 2014-04-24 | Applies to new plumbing and altered/replaced plumbing; existing buildings are covered when plumbing is altered, use changes, or a health/safety hazard exists. | src:usa-il:plumbing-code-77-iac-890; src:usa-il:plumbing-license-law-225-ilcs-320 |
| Fuel Gas | Health/Life Safety Code for Public Schools fuel-gas subcode | IFGC | 2024 | mandatory_limited_public_school_scope | 2025-01-30 | 2025-01-30 | 2025-01-01 | 2025-01-01 | Applies to public-school projects with design contracts executed on or after 2025-01-01; private statewide fuel-gas-code scope outside local codes remains unresolved. | src:usa-il:hls-code-public-schools-23-iac-180 |
| Electrical | Commercial non-building-code NEC certification requirement | NEC | Current or most recent preceding | mandatory_limited_scope | 2024-01-01 | 2025-01-01 | 2025-01-01 | 2025-01-01 | For commercial buildings in non-building-code jurisdictions, current/current-preceding NEC is part of certification package; local electrical permitting remains jurisdiction-specific. | src:usa-il:cdb-act-20-ilcs-3105 |
| Energy | Illinois Energy Conservation Code | IECC | 2024, as amended | mandatory_statewide_scope | 2025-11-30 | 2025-11-30 | 2025-11-30 | 2025-11-30 | Applies to privately funded commercial facilities and residential buildings; local residential energy standards are generally minimum-and-maximum except statutory exceptions and stretch-code jurisdictions. | src:usa-il:energy-code-71-iac-600; src:usa-il:energy-efficient-building-act-20-ilcs-3125; src:usa-il:energy-code-cdb-page |
| Fire - construction references | Fire Investigation Act rules / public-school IFC subcode | NFPA 101; IFC | NFPA 101 2015 in Part 100; 2024 IFC excluding Ch. 4 for public schools | mandatory_with_scope_exceptions | 2020-01-01 for Part 100; 2025-01-30 for school rule update | 2020-01-01; 2025-01-30 | 2020-01-01; 2025-01-01 for school design-contract trigger | 2020-01-01; 2025-01-01 for school design-contract trigger | OSFM Part 100 applies unless local equal-or-higher criteria are met; public-school IFC trigger follows design-contract date. | src:usa-il:fire-prevention-rules-41-iac-100; src:usa-il:hls-code-public-schools-23-iac-180 |
| Fire - operational / prevention code | Fire Prevention and Safety Rules | NFPA 101 Life Safety Code | 2015, subject to OSFM modifications | mandatory_with_local_equal_or_higher_exceptions | 2020-01-01 | 2020-01-01 | 2020-01-01 | 2020-01-01 | Local AHJs may be recognized as equal or higher under OSFM criteria; owner/occupant/lessee remains responsible for compliance. | src:usa-il:fire-prevention-rules-41-iac-100; src:usa-il:osfm-life-safety-faq |
| Accessibility | Illinois Accessibility Code | Illinois Accessibility Code / 2010 ADA Standards comparison | 2018 Part 400 code | mandatory_scope_defined_by_EBA | 2018-10-23 | 2018-10-23 | 2018-10-23 | 2018-10-23 | Applies under Environmental Barriers Act scope; IBC Chapter 11 is displaced by Illinois Accessibility Code in CDB non-building-code commercial and public-school contexts. | src:usa-il:accessibility-code-71-iac-400; src:usa-il:environmental-barriers-act-410-ilcs-25 |
| Elevator / Conveyance | Illinois Elevator Safety Rules | State elevator/conveyance rules with incorporated standards | 41 Ill. Adm. Code 1000; latest located amendment effective 2023-12-28 | mandatory_outside_chicago_with_state_owned_chicago_exception | null | 2023-12-28 | 2023-12-28 | 2023-12-28 | Applies to covered conveyances outside municipalities over 500,000 population, except state-owned buildings in that municipality; local elevator agreements must be at least as stringent as state rules. | src:usa-il:elevator-safety-rules-41-iac-1000; src:usa-il:elevator-act-225-ilcs-312; src:usa-il:osfm-elevator-page |

### 3.2 Adoption Records

```yaml
adoption_records:
  - adoption_id: "adoption:usa-il:commercial-noncode-ibc-iebc-nec-2025"
    code_families: ["Building", "Existing Building / Rehabilitation", "Electrical"]
    state_code_name: "Commercial non-building-code certification framework"
    base_codes:
      - "International Building Code, current edition or most recent preceding edition, including Appendix G and excluding Chapters 11, 13, and 29"
      - "International Existing Building Code, current edition or most recent preceding edition"
      - "National Electrical Code, current edition or most recent preceding edition"
    scope: "Newly constructed or substantially improved commercial buildings in non-building-code jurisdictions"
    adoption_date: "2024-01-01"
    effective_date: "2025-01-01"
    operative_date: "2025-01-01"
    mandatory_date: "2025-01-01"
    transition_rule_ids: ["date-rule:usa-il:commercial-noncode-jan1-permit"]
    source_ids: ["src:usa-il:cdb-act-20-ilcs-3105"]
    confidence: 0.86

  - adoption_id: "adoption:usa-il:local-structural-baseline-2025"
    code_families: ["Building", "Residential", "Existing Building / Rehabilitation"]
    state_code_name: "Municipal and county structural baseline requirement"
    base_codes:
      - "International Building Code baseline for new nonresidential structural design"
      - "International Existing Building Code baseline for existing nonresidential structural rehabilitation"
      - "International Residential Code baseline for residential structural design"
    scope: "Municipal and county building codes beginning 2025-01-01"
    adoption_date: "2024-01-01"
    effective_date: "2025-01-01"
    operative_date: "2025-01-01"
    mandatory_date: "2025-01-01"
    transition_rule_ids: ["date-rule:usa-il:local-structural-baseline-2025"]
    source_ids: ["src:usa-il:cdb-act-20-ilcs-3105"]
    confidence: 0.84

  - adoption_id: "adoption:usa-il:residential-noncode-contract-default"
    code_families: ["Residential"]
    state_code_name: "Residential Building Code Act non-building-code jurisdiction contract/default rule"
    base_codes:
      - "International Residential Code, current edition or most recent preceding edition, excluding Parts IV and VII, where agreed"
      - "Current International Residential Code if no code is agreed or stated"
      - "Illinois Energy Conservation Code"
      - "Illinois Plumbing Code"
    scope: "New residential construction in non-building-code jurisdictions"
    adoption_date: "2024-01-01"
    effective_date: "2024-01-01"
    operative_date: "2024-01-01"
    mandatory_date: "2024-01-01"
    transition_rule_ids: ["date-rule:usa-il:residential-noncode-contract"]
    source_ids: ["src:usa-il:residential-building-code-act-815-ilcs-670"]
    confidence: 0.82

  - adoption_id: "adoption:usa-il:energy-2024-iecc"
    code_families: ["Energy"]
    state_code_name: "Illinois Energy Conservation Code"
    base_codes: ["2024 International Energy Conservation Code, including published errata and excluding published supplements, with Illinois amendments"]
    scope: "Privately funded commercial facilities and residential buildings, subject to statutory exceptions"
    adoption_date: "2025-11-30"
    effective_date: "2025-11-30"
    operative_date: "2025-11-30"
    mandatory_date: "2025-11-30"
    transition_rule_ids: ["date-rule:usa-il:energy-permit-application"]
    source_ids: ["src:usa-il:energy-code-71-iac-600", "src:usa-il:energy-code-cdb-page"]
    confidence: 0.86

  - adoption_id: "adoption:usa-il:stretch-energy-code"
    code_families: ["Energy"]
    state_code_name: "Illinois Stretch Energy Code"
    base_codes: ["Illinois commercial and residential stretch energy code components based on statutory site-energy-index targets"]
    scope: "Municipalities that adopt the Stretch Energy Code and projects authorized or funded by CDB after 2024-07-01"
    adoption_date: "2024-06-30"
    effective_date: "2025-01-01"
    operative_date: "2025-01-01"
    mandatory_date: "2024-07-01"
    notes: "Mandatory date applies to CDB-authorized or CDB-funded projects to which an energy conservation code is applicable; municipal stretch-code applicability depends on local adoption."
    transition_rule_ids: ["date-rule:usa-il:stretch-energy-cdb-funded"]
    source_ids: ["src:usa-il:energy-efficient-building-act-20-ilcs-3125", "src:usa-il:energy-code-71-iac-600"]
    confidence: 0.78

  - adoption_id: "adoption:usa-il:fire-nfpa101-2015"
    code_families: ["Fire - operational / prevention code"]
    state_code_name: "Fire Prevention and Safety Rules"
    base_codes: ["NFPA 101 Life Safety Code, 2015 edition, subject to OSFM modifications"]
    scope: "Fire prevention and life-safety scope under 41 Ill. Adm. Code 100, subject to local equal-or-higher treatment"
    adoption_date: "2020-01-01"
    effective_date: "2020-01-01"
    operative_date: "2020-01-01"
    mandatory_date: "2020-01-01"
    transition_rule_ids: ["date-rule:usa-il:fire-nfpa101-new-existing-2020"]
    source_ids: ["src:usa-il:fire-prevention-rules-41-iac-100", "src:usa-il:osfm-life-safety-faq"]
    confidence: 0.82

  - adoption_id: "adoption:usa-il:plumbing-code-77-iac-890"
    code_families: ["Plumbing"]
    state_code_name: "Illinois Plumbing Code"
    base_codes: ["77 Ill. Adm. Code 890"]
    scope: "Design and installation of new plumbing and alteration/replacement of plumbing systems"
    adoption_date: null
    effective_date: "2014-04-24"
    operative_date: "2014-04-24"
    mandatory_date: "2014-04-24"
    transition_rule_ids: ["date-rule:usa-il:plumbing-scope-existing-buildings"]
    source_ids: ["src:usa-il:plumbing-code-77-iac-890"]
    confidence: 0.80

  - adoption_id: "adoption:usa-il:public-schools-2024-icc"
    code_families: ["Building", "Existing Building / Rehabilitation", "Energy", "Fire - construction references", "Fuel Gas", "Mechanical"]
    state_code_name: "Health/Life Safety Code for Public Schools"
    base_codes:
      - "2024 International Building Code"
      - "2024 International Energy Conservation Code"
      - "2024 International Existing Building Code"
      - "2024 International Fire Code, excluding Chapter 4"
      - "2024 International Fuel Gas Code"
      - "2024 International Mechanical Code"
      - "2024 International Property Maintenance Code"
    scope: "Public school facilities subject to 23 Ill. Adm. Code 180, with Article 34 exception noted"
    adoption_date: "2025-01-30"
    effective_date: "2025-01-30"
    operative_date: "2025-01-01"
    mandatory_date: "2025-01-01"
    transition_rule_ids: ["date-rule:usa-il:public-schools-design-contract-2025"]
    source_ids: ["src:usa-il:school-code-105-ilcs-5-2-3-12", "src:usa-il:hls-code-public-schools-23-iac-180"]
    confidence: 0.84

  - adoption_id: "adoption:usa-il:accessibility-code-2018"
    code_families: ["Accessibility"]
    state_code_name: "Illinois Accessibility Code"
    base_codes: ["71 Ill. Adm. Code 400, primarily Appendix A"]
    scope: "Environmental Barriers Act public-facility and multi-story housing scope"
    adoption_date: "2018-10-23"
    effective_date: "2018-10-23"
    operative_date: "2018-10-23"
    mandatory_date: "2018-10-23"
    transition_rule_ids: ["date-rule:usa-il:accessibility-current-code"]
    source_ids: ["src:usa-il:accessibility-code-71-iac-400", "src:usa-il:environmental-barriers-act-410-ilcs-25"]
    confidence: 0.80

  - adoption_id: "adoption:usa-il:elevator-41-iac-1000"
    code_families: ["Elevator / Conveyance"]
    state_code_name: "Illinois Elevator Safety Rules"
    base_codes: ["41 Ill. Adm. Code 1000, incorporated standards not fully parsed in this update"]
    scope: "Covered conveyances outside municipalities over 500,000 population, except state-owned buildings in such municipality"
    adoption_date: null
    effective_date: "2023-12-28"
    operative_date: "2023-12-28"
    mandatory_date: "2023-12-28"
    transition_rule_ids: ["date-rule:usa-il:elevator-2023-rules"]
    source_ids: ["src:usa-il:elevator-safety-rules-41-iac-1000", "src:usa-il:elevator-act-225-ilcs-312"]
    confidence: 0.78
```

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

Illinois uses several different date triggers. The commercial non-building-code certification rule locks the applicable requirements to the requirements in effect on January 1 of the year when the permit application was made, or January 1 of the year when construction begins if no permit is required. Local building-code structural baseline requirements begin January 1, 2025. Public-school Health/Life Safety Code applicability turns on design-contract execution date for the 2024 ICC code set. The 2024 Illinois Energy Conservation Code became effective November 30, 2025. OSFM's 2015 NFPA 101 adoption uses the January 1, 2020 adoption date as an important new/existing occupancy marker.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-il:commercial-noncode-jan1-permit | Commercial buildings in non-building-code jurisdictions | permit_application_year_lock | January 1 of permit-application year, or January 1 of construction-start year if no permit required | Building permit application or construction start where no permit is required | yes, if it was the applicable requirement on that January 1 | src:usa-il:cdb-act-20-ilcs-3105 | verified |
| date-rule:usa-il:local-structural-baseline-2025 | Municipal and county building codes | mandatory_start | 2025-01-01 | Local building code regulating new nonresidential, existing nonresidential rehabilitation, or residential structural design | local codes remain, but baseline stringency applies | src:usa-il:cdb-act-20-ilcs-3105 | verified |
| date-rule:usa-il:local-code-reporting-30-day | Municipal/county adopted building-code editions and amendments | reporting_before_effective_date | at least 30 days before effective date | Adoption of new building-code edition or amendment | yes; report does not itself replace local adoption date | src:usa-il:cdb-act-20-ilcs-3105; src:usa-il:cdb-building-codes-page | verified |
| date-rule:usa-il:residential-noncode-contract | New residential construction in non-building-code jurisdictions | contract_or_default | code in effect on first day of construction for agreed nearby municipal/county code; current IRC if no agreement | Home builder and purchaser agreement, or lack of stated/agreed residential building code | yes, if validly agreed under statutory options | src:usa-il:residential-building-code-act-815-ilcs-670 | verified |
| date-rule:usa-il:energy-permit-application | Illinois Energy Conservation Code | permit_application | 2025-11-30 current 2024 IECC effective date; statutory trigger uses building-permit application | New building or structure for which a building permit application is received by municipality or county | prior code transition not fully parsed for applications straddling 2025-11-30 | src:usa-il:energy-code-cdb-page; src:usa-il:energy-efficient-building-act-20-ilcs-3125; src:usa-il:energy-code-71-iac-600 | partially_verified |
| date-rule:usa-il:stretch-energy-cdb-funded | Illinois Stretch Energy Code | funding_authorization_scope | after 2024-07-01 | Project to which an energy conservation code applies and that is authorized or funded in any part by CDB | no, if stretch code applies; municipal adoption details must be checked locally | src:usa-il:energy-efficient-building-act-20-ilcs-3125 | partially_verified |
| date-rule:usa-il:fire-nfpa101-new-existing-2020 | OSFM Life Safety Code | adoption_date_new_existing_marker | 2020-01-01 | Determining application of NFPA 101 new vs. existing occupancy provisions | yes, existing occupancy provisions may apply where applicable | src:usa-il:fire-prevention-rules-41-iac-100; src:usa-il:osfm-life-safety-faq | partially_verified |
| date-rule:usa-il:public-schools-design-contract-2025 | Public school Health/Life Safety Code | design_contract_trigger | 2025-01-01 | Design contract executed on or after 2025-01-01 | existing legal occupancy before/on 2024-12-31 may continue until repairs, alterations, occupancy change, relocation, or additions trigger IEBC | src:usa-il:hls-code-public-schools-23-iac-180 | verified |
| date-rule:usa-il:plumbing-scope-existing-buildings | Illinois Plumbing Code | existing_building_scope | 2014-04-24 located latest amendment effective date | Alteration, renovation, replacement, use/classification change, or health/safety hazard | yes for existing buildings with no plumbing alteration/use change/hazard under Part 890 scope | src:usa-il:plumbing-code-77-iac-890 | partially_verified |
| date-rule:usa-il:accessibility-current-code | Illinois Accessibility Code | code_commencement_scope | 2018-10-23 current Part 400 effective date; EBA scope includes construction commencement rules | Public facility or multi-story housing under Environmental Barriers Act scope | project-specific code version depends on construction commencement; requires project-date analysis | src:usa-il:accessibility-code-71-iac-400; src:usa-il:environmental-barriers-act-410-ilcs-25 | partially_verified |
| date-rule:usa-il:elevator-2023-rules | Illinois Elevator Safety Rules | rule_effective_date | 2023-12-28 | Covered conveyance work subject to Part 1000 | no transition parsed beyond latest rule effective date | src:usa-il:elevator-safety-rules-41-iac-1000 | partially_verified |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Energy | Illinois Stretch Energy Code update | null | 2026-12-31 latest statutory deadline for next residential/commercial stretch-code targets | null | null | null | active_watch | src:usa-il:energy-efficient-building-act-20-ilcs-3125 | Statute sets December 31, 2026 site-energy-index milestones; verify rulemaking and final text before updating adoption records. |
| Energy | Illinois Stretch Energy Code later updates | null | 2029-12-31 and 2032-12-31 statutory target dates | null | null | null | active_watch | src:usa-il:energy-efficient-building-act-20-ilcs-3125 | Future statutory milestones; not yet current adopted code records. |
| Elevator / Conveyance | 2025 proposed Part 1000 amendments | 2025-01-01 | null | null | null | null | closed_watch | src:usa-il:osfm-elevator-page | OSFM page indicated the proposed 2025 amendments were withdrawn as of 2026-02-06; monitor for refiling. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-il:non-building-code-commercial | Building / Existing / Electrical / Energy / Accessibility / Plumbing / Fire | Newly constructed or substantially improved commercial building in non-building-code jurisdiction | Occupancy of covered commercial building | Owner or agent must obtain qualified inspection and file certification of compliance with listed codes before occupancy. | src:usa-il:cdb-act-20-ilcs-3105 | verified |
| applicability-rule:usa-il:state-funded-buildings | Building / Energy / Accessibility / Plumbing / Fire | State-funded buildings and CDB-authorized/funded projects | State or CDB funding/authorization | CDB page treats state-funded buildings with commercial non-building-code requirements; Stretch Energy Code applies to CDB-authorized/funded energy-code projects after 2024-07-01. | src:usa-il:cdb-building-codes-page; src:usa-il:energy-efficient-building-act-20-ilcs-3125 | partially_verified |
| applicability-rule:usa-il:article-34-schools | Public schools | School districts governed by Article 34 and charter schools in those districts | Article 34 governance/geographic boundaries | Part 180 applies only to Section 180.250 for these districts; otherwise local building codes apply. | src:usa-il:hls-code-public-schools-23-iac-180 | partially_verified |
| applicability-rule:usa-il:elevator-chicago | Elevator / Conveyance | Municipality with population over 500,000 | Location and ownership | Part 1000 does not apply to a municipality over 500,000 population, except state-owned buildings in that municipality. | src:usa-il:elevator-safety-rules-41-iac-1000; src:usa-il:osfm-elevator-page | verified |
| applicability-rule:usa-il:local-fire-equivalency | Fire - operational / prevention code | Local AHJ with fire/life-safety code | OSFM equal-or-higher criteria | Local AHJ code may be considered equal or higher if it meets criteria such as adopting listed model codes or obtaining OSFM agreement/demonstration. | src:usa-il:fire-prevention-rules-41-iac-100 | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** multi-agency, code-family-specific.

**Where amendments are published:** Illinois Administrative Code Parts and agency pages, including 71 Ill. Adm. Code 600 for energy, 41 Ill. Adm. Code 100 for OSFM fire/life-safety modifications, 71 Ill. Adm. Code 400 for accessibility, 77 Ill. Adm. Code 890 for plumbing, 23 Ill. Adm. Code 180 for public schools, and 41 Ill. Adm. Code 1000 for elevators.

**Amendment parsing status:** partial. High-level amendment locations and several high-impact substitutions were captured; line-by-line amendment extraction remains open.

### 5.2 State Amendment Sources

| Amendment Source ID | Code Family | Publication Path | What It Contains | Source IDs | Parsing Status |
| --- | --- | --- | --- | --- | --- |
| amend-source:usa-il:energy-2024-iecc | Energy | 71 Ill. Adm. Code 600 Appendix A | Illinois amendments to the 2024 IECC for Illinois Energy Conservation Code | src:usa-il:energy-code-71-iac-600 | location_verified_not_fully_parsed |
| amend-source:usa-il:commercial-stretch-energy | Energy | 71 Ill. Adm. Code 600 Appendix B | Illinois Commercial Stretch Energy Code amendments to 2024 IECC final draft | src:usa-il:energy-code-71-iac-600 | location_verified_not_fully_parsed |
| amend-source:usa-il:residential-stretch-energy | Energy | 71 Ill. Adm. Code 600 Appendix C | Illinois Residential Stretch Energy Code amendments to 2024 IECC | src:usa-il:energy-code-71-iac-600 | location_verified_not_fully_parsed |
| amend-source:usa-il:fire-nfpa101-modifications | Fire | 41 Ill. Adm. Code 100.7(c) | OSFM modifications to NFPA 101 (2015) adoption | src:usa-il:fire-prevention-rules-41-iac-100; src:usa-il:osfm-life-safety-faq | location_verified_not_fully_parsed |
| amend-source:usa-il:schools-code-substitutions | Public schools | 23 Ill. Adm. Code 180.60(b) | Public-school substitutions for IFC Chapter 4, IBC admin provisions, IBC Chapter 11 accessibility, plumbing, boiler, elevators, and sprinklers | src:usa-il:hls-code-public-schools-23-iac-180 | partially_parsed |
| amend-source:usa-il:cdb-commercial-exclusions | Commercial non-building-code | 20 ILCS 3105/10.09-1 | IBC included with Appendix G and excluding Chapters 11, 13, and 29 for non-building-code commercial certification | src:usa-il:cdb-act-20-ilcs-3105 | partially_parsed |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Summary | Practical Impact | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| amendment:usa-il:cdb-commercial-ibc-exclusions | Building / Accessibility / Energy / Plumbing | For commercial non-building-code certification, IBC is used with Appendix G but excludes Chapters 11, 13, and 29, while Illinois accessibility, energy/stretch, and plumbing codes are separately required. | Avoid treating model IBC accessibility, energy, or plumbing chapters as controlling where Illinois specialty codes replace them. | src:usa-il:cdb-act-20-ilcs-3105 | partially_verified |
| amendment:usa-il:schools-ibc-chapter-11-substitution | Accessibility / Public schools | Public-school Health/Life Safety Code uses Illinois Accessibility Code instead of IBC Chapter 11. | School projects need Illinois Accessibility Code review, not model IBC Chapter 11 alone. | src:usa-il:hls-code-public-schools-23-iac-180 | verified |
| amendment:usa-il:schools-plumbing-substitution | Plumbing / Public schools | Public-school code uses Illinois Plumbing Code and 2024 IPC 405.3.1 instead of IBC plumbing references. | School plumbing review must include Illinois Plumbing Code. | src:usa-il:hls-code-public-schools-23-iac-180 | verified |
| amendment:usa-il:osfm-nfpa101-modifications | Fire | OSFM adopts NFPA 101 (2015) subject to Part 100.7(c) modifications. | NFPA 101 text alone is insufficient; OSFM modifications must be applied. | src:usa-il:fire-prevention-rules-41-iac-100; src:usa-il:osfm-life-safety-faq | location_verified_not_fully_parsed |
| amendment:usa-il:energy-2024-iecc-appendix-a | Energy | Part 600 incorporates 2024 IECC with Illinois modifications in Appendix A. | Illinois amendments must be applied to base 2024 IECC. | src:usa-il:energy-code-71-iac-600 | location_verified_not_fully_parsed |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-il"
  model: "hybrid_local_state_specialty"
  enforcing_entities:
    - "municipal_building_departments_where_local_code_exists"
    - "county_building_departments_where_county_code_exists_or_unincorporated_scope_applies"
    - "qualified_inspectors_for_commercial_buildings_in_non-building-code_jurisdictions"
    - "local_fire_AHJs_under_OSFM_direction_or_equal_higher_local_fire_codes"
    - "IDPH_and_local_plumbing_inspectors_for_plumbing_scope"
    - "ISBE_regional_superintendents_and_local_school_boards_for_public_school_HLS_scope"
    - "OSFM_Elevator_Safety_Division_and_local_elevator_administrators_by_agreement"
  required_officials:
    - "qualified_inspector_for_commercial_non-building-code_certification"
    - "licensed_plumbing_inspector_where_local_plumbing_inspection_is_administered"
    - "qualified_school_plan_reviewers_and_inspectors_under_23_Ill_Adm_Code_180"
    - "certified_elevator_inspector_for elevator inspections under state/local elevator framework"
  state_reserved_activities:
    - "CDB local building-code reporting/public website function"
    - "CDB energy-code and stretch-energy-code rulemaking"
    - "CDB accessibility-code rulemaking and project-specific interpretations"
    - "Illinois Attorney General accessibility-code enforcement under EBA"
    - "IDPH variances from Illinois Plumbing Code"
    - "OSFM Fire Prevention and Safety Rules and equal-or-higher fire-code determinations"
    - "OSFM elevator/conveyance regulation outside Chicago except authorized local agreements"
    - "ISBE public-school Health/Life Safety Code administration"
  source_ids:
    - "src:usa-il:cdb-act-20-ilcs-3105"
    - "src:usa-il:fire-prevention-rules-41-iac-100"
    - "src:usa-il:plumbing-code-77-iac-890"
    - "src:usa-il:hls-code-public-schools-23-iac-180"
    - "src:usa-il:elevator-safety-rules-41-iac-1000"
  verification_status: "partially_verified"
  confidence: 0.72
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-il"
  model: "local_building_codes_allowed_subject_to_reporting_and_state_specialty_constraints"
  applies_to_code_families:
    - "Building"
    - "Residential"
    - "Existing Building / Rehabilitation"
    - "Energy"
    - "Fire - operational / prevention code"
    - "Elevator / Conveyance"
    - "Plumbing"
    - "Accessibility"
  approval_required: "varies_by_code_family"
  approving_authority_id: "varies: CDB reporting for local building codes; OSFM criteria for equal-or-higher fire/life-safety; OSFM/Elevator Board framework for local elevator variances; IDPH for plumbing variances"
  filing_required: true
  registry_exists: true
  registry_source_ids:
    - "src:usa-il:cdb-act-20-ilcs-3105"
    - "src:usa-il:cdb-building-codes-page"
  legal_basis_source_ids:
    - "src:usa-il:cdb-act-20-ilcs-3105"
    - "src:usa-il:energy-efficient-building-act-20-ilcs-3125"
    - "src:usa-il:fire-prevention-rules-41-iac-100"
    - "src:usa-il:elevator-act-225-ilcs-312"
    - "src:usa-il:plumbing-code-77-iac-890"
  verification_status: "partially_verified"
  confidence: 0.70
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement and local amendment authority diverge in Illinois. A city or county may be the front-line building AHJ and may have local amendments, but the code applied to a project must still satisfy statewide structural baseline rules and separate statewide specialty requirements. A local fire department may enforce fire-safety ordinances, but OSFM Part 100 still matters unless the local code qualifies as equal or higher under OSFM criteria. Local plumbing enforcement does not imply local authority to grant Plumbing Code variances; IDPH retains sole variance discretion for Part 890. Local elevator administration requires a state/local elevator agreement and at-least-as-stringent standards.

### 6.4 Known Local Amendment Registries

| Registry ID | Registry Name | Responsible Entity | Coverage | Source IDs | Extraction Status |
| --- | --- | --- | --- | --- | --- |
| registry:usa-il:cdb-local-code-directory | Illinois Municipal Code Directory / CDB local code reporting | CDB | Reported municipal and county building-code title, edition, amendment status, and reporting date | src:usa-il:cdb-act-20-ilcs-3105; src:usa-il:cdb-building-codes-page | registry_identified_not_extracted |

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendments were extracted into this report. Project-level use should query the CDB local code directory and the relevant municipality/county ordinance record before relying on local code edition or amendment status.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: partially_started

Jurisdiction stack:

```text
Address
  -> State: Illinois
  -> County
  -> Municipality / unincorporated county
  -> Determine whether municipality/county has reported an adopted building code to CDB
  -> If no reported/adopted building code: apply non-building-code jurisdiction rules where project scope fits
  -> Determine project type: commercial, residential, public school, state-funded, CDB-funded/authorized, elevator/conveyance, licensed facility, other specialty scope
  -> Building AHJ: local municipality/county or qualified-inspector certification path
  -> Fire AHJ: OSFM and/or local fire AHJ; check equal-or-higher local fire code status
  -> Plumbing AHJ: IDPH / local plumbing inspection program
  -> Energy AHJ: local code official enforcing Illinois Energy Conservation Code / CDB rule framework; check Stretch Energy Code adoption/funding trigger
  -> Accessibility: Illinois Accessibility Code scope; CDB interpretation / Illinois Attorney General enforcement path
  -> Public school AHJ: ISBE / regional superintendent / local school board, unless Article 34 exception routes to local building codes except sprinkler section
  -> Elevator AHJ: OSFM Elevator Safety Division or local elevator administrator by agreement; Chicago exception check
  -> Applicable state code adoption records
  -> Applicable local amendment records
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | not selected | none | statewide | unknown | pending |
| County | not selected | none | statewide | unknown | pending |
| Municipality | not selected | none | statewide | unknown | pending |
| Fire District | not selected | none | statewide | unknown | pending |
| School District / Article 34 | not selected | none | statewide | unknown | pending |
| Elevator local-administrator jurisdiction | OSFM municipalities list page identified but not extracted | src:usa-il:osfm-elevator-page | statewide outside Chicago with local-agreement overlays | unknown | pending_extraction |

### 7.3 AHJ Contact Data

No AHJ contact data was populated. The report identifies authority classes and source registries, but it does not provide project-level contact routing for a specific address.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Publisher | Source Type | URL / Locator | Supports | Status |
| --- | --- | --- | --- | --- | --- | --- |
| src:usa-il:cdb-act-20-ilcs-3105 | Capital Development Board Act, 20 ILCS 3105 | Illinois General Assembly | statute | https://www.ilga.gov/legislation/ILCS/details?ActID=362&ActName=Capital+Development+Board+Act.&ChapAct=20+ILCS+3105%2F&Chapter=&ChapterID=5&MajorTopic=&Print=True&SeqEnd=6000000&SeqStart=100000 | CDB authority, commercial non-building-code certification, local code reporting, local structural baseline | parsed_core_sections |
| src:usa-il:cdb-building-codes-page | Building Codes and Regulations | Illinois Capital Development Board | agency page | https://cdb.illinois.gov/business/codes/buildingcodesregulations.html | CDB summary of post-2025 code framework, local reporting, public schools routing, CDB directories | parsed_summary_page |
| src:usa-il:energy-efficient-building-act-20-ilcs-3125 | Energy Efficient Building Act, 20 ILCS 3125 | Illinois General Assembly | statute | https://www.ilga.gov/Legislation/ILCS/Articles?ActID=2614&ChapterID=5 | Energy Code adoption authority, applicability, local limits, Stretch Energy Code | parsed_core_sections |
| src:usa-il:energy-code-71-iac-600 | 71 Ill. Adm. Code 600, Illinois Energy Code | Illinois Joint Committee on Administrative Rules / CDB | administrative rule | https://www.ilga.gov/agencies/JCAR/EntirePart?titlepart=07100600 | 2024 IECC incorporation, energy-code amendments, stretch-energy rules | parsed_core_sections |
| src:usa-il:energy-code-cdb-page | Illinois Energy Conservation Code page | Illinois Capital Development Board | agency page | https://cdb.illinois.gov/business/codes/illinois-energy-codes/illinois-energy-conservation-code.html | 2024 IECC amended effective date for privately funded commercial and residential buildings | parsed_summary_page |
| src:usa-il:residential-building-code-act-815-ilcs-670 | Residential Building Code Act, 815 ILCS 670 | Illinois General Assembly | statute | https://www.ilga.gov/Legislation/ILCS/Articles?ActID=2585&ChapterID=67 | Residential non-building-code jurisdiction contract/default rule and definitions | parsed_core_sections |
| src:usa-il:fire-prevention-rules-41-iac-100 | 41 Ill. Adm. Code 100, Fire Prevention and Safety | OSFM / Illinois JCAR | administrative rule | https://www.ilga.gov/agencies/JCAR/EntirePart?titlepart=04100100 | OSFM fire-prevention authority, NFPA 101 adoption, local equal-or-higher criteria | parsed_core_sections |
| src:usa-il:osfm-life-safety-faq | Fire Prevention & Building Safety FAQ | Office of the State Fire Marshal | agency FAQ | https://sfm.illinois.gov/about/divisions/fire-prevention-and-building-safety/frequently-asked-questions.html | OSFM NFPA 101 adoption explanation, modifications, applicability notes | parsed_summary_page |
| src:usa-il:plumbing-code-77-iac-890 | 77 Ill. Adm. Code 890, Illinois Plumbing Code | IDPH / Illinois JCAR | administrative rule | https://www.ilga.gov/agencies/JCAR/EntirePart?titlepart=07700890 | Plumbing-code authority, applicability, variance authority | parsed_core_sections |
| src:usa-il:plumbing-license-law-225-ilcs-320 | Illinois Plumbing License Law, 225 ILCS 320 | Illinois General Assembly | statute | https://www.ilga.gov/Documents/legislation/ilcs/documents/022503200K37.htm | Local plumbing permits/inspection authority and licensed-inspector requirement | parsed_core_section_37 |
| src:usa-il:environmental-barriers-act-410-ilcs-25 | Environmental Barriers Act, 410 ILCS 25 | Illinois General Assembly | statute | https://www.ilga.gov/Legislation/ILCS/Articles?ActID=1519&Chapter=PUBLIC+HEALTH&ChapterID=35&MajorTopic=HEALTH+AND+SAFETY | CDB adoption of Illinois Accessibility Code, project-specific interpretations, scope | parsed_core_sections |
| src:usa-il:accessibility-code-71-iac-400 | 71 Ill. Adm. Code 400, Illinois Accessibility Code | CDB / Illinois JCAR | administrative rule | https://www.ilga.gov/agencies/JCAR/EntirePart?titlepart=07100400 | Accessibility Code purpose, effective date, force of law, AG enforcement | parsed_core_sections |
| src:usa-il:school-code-105-ilcs-5-2-3-12 | School Code, 105 ILCS 5/2-3.12 | Illinois General Assembly | statute | https://www.ilga.gov/documents/legislation/ilcs/documents/010500050K2-3.12.htm | Health/Life Safety Code as governing code for public schools, inspection coordination, Article 34 limit | parsed_core_section |
| src:usa-il:hls-code-public-schools-23-iac-180 | 23 Ill. Adm. Code 180, Health/Life Safety Code for Public Schools | ISBE / Illinois JCAR | administrative rule | https://ilga.gov/agencies/JCAR/EntirePart?titlepart=02300180 | 2024 ICC public-school code adoptions, scope, Article 34 exception, substitutions | parsed_core_sections |
| src:usa-il:elevator-safety-rules-41-iac-1000 | 41 Ill. Adm. Code 1000, Illinois Elevator Safety Rules | OSFM / Illinois JCAR | administrative rule | https://ilga.gov/agencies/JCAR/EntirePart?titlepart=04101000 | Elevator/conveyance rule authority, scope, latest effective date, Chicago exception | parsed_core_sections |
| src:usa-il:elevator-act-225-ilcs-312 | Elevator Safety and Regulation Act, 225 ILCS 312 | Illinois General Assembly | statute | https://ilga.gov/Documents/legislation/ilcs/documents/022503120K140.htm | Local elevator administrator agreements, at-least-as-stringent local standards, variance notice | parsed_section_140 |
| src:usa-il:osfm-elevator-page | Elevator Safety | Office of the State Fire Marshal | agency page | https://sfm.illinois.gov/about/divisions/elevators.html | OSFM Elevator Safety Division role, outside-Chicago summary, rulemaking notices | parsed_summary_page |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-il:cdb-act-20-ilcs-3105 | statutory_database_caveat | ILGA notes its statutory database is updated on an ongoing basis and Public Acts should be checked when source notes include future-effective changes. | verify_against_public_acts_for_final_certification |
| src:usa-il:cdb-building-codes-page | agency_summary | CDB page is a useful official summary, but legal fields should trace back to statutes/rules. | use_as_secondary_official_summary |
| src:usa-il:energy-code-cdb-page | agency_summary | CDB page states effective date and links amendments; use Part 600 for enforceable rule details. | use_with_rule_crosscheck |
| src:usa-il:fire-prevention-rules-41-iac-100 | administrative_rule_html | Part 100 text was accessed via ILGA/JCAR HTML; incorporated copyrighted NFPA text is not reproduced. | use_rule_adoption_and_modifications_only |
| src:usa-il:osfm-life-safety-faq | agency_faq | FAQ explains OSFM practice but should not replace Part 100 where rule text controls. | use_as_context_with_rule |
| src:usa-il:plumbing-code-77-iac-890 | administrative_rule_large_html | Full Part 890 is large; only authority, scope, and variance provisions were extracted. | parse_full_part_before_production_use |
| src:usa-il:hls-code-public-schools-23-iac-180 | administrative_rule_scope_exception | Article 34 school-district exception requires project-specific school-district check. | require_project_level_resolution |
| src:usa-il:elevator-safety-rules-41-iac-1000 | incorporated_standards_not_fully_parsed | Incorporated elevator standard editions were not extracted into this report. | parse_section_1000_60_before_elevator_code_matrix_finalization |

### 8.3 Supplemental Sources

None used. All populated legal and code-adoption fields in this update are based on Illinois official state statutes, administrative rules, or agency pages.

### 8.4 Source Extraction Metadata

| Extraction ID | Source IDs | Extracted Fields | Extracted By | Extraction Date | Confidence |
| --- | --- | --- | --- | --- | --- |
| extract:usa-il:2026-06-25:cdb-framework | src:usa-il:cdb-act-20-ilcs-3105; src:usa-il:cdb-building-codes-page | CDB authority, commercial non-code certification, local reporting, structural baseline | ChatGPT | 2026-06-25 | 0.82 |
| extract:usa-il:2026-06-25:energy | src:usa-il:energy-efficient-building-act-20-ilcs-3125; src:usa-il:energy-code-71-iac-600; src:usa-il:energy-code-cdb-page | 2024 IECC, stretch code, local energy limits | ChatGPT | 2026-06-25 | 0.84 |
| extract:usa-il:2026-06-25:fire | src:usa-il:fire-prevention-rules-41-iac-100; src:usa-il:osfm-life-safety-faq | OSFM authority, NFPA 101 adoption, equal-or-higher local fire codes | ChatGPT | 2026-06-25 | 0.80 |
| extract:usa-il:2026-06-25:plumbing | src:usa-il:plumbing-code-77-iac-890; src:usa-il:plumbing-license-law-225-ilcs-320 | Plumbing Code authority/scope, local inspection authority, variance authority | ChatGPT | 2026-06-25 | 0.78 |
| extract:usa-il:2026-06-25:schools | src:usa-il:school-code-105-ilcs-5-2-3-12; src:usa-il:hls-code-public-schools-23-iac-180 | HLS public-school authority, 2024 ICC code set, design-contract trigger, Article 34 exception | ChatGPT | 2026-06-25 | 0.82 |
| extract:usa-il:2026-06-25:elevator | src:usa-il:elevator-safety-rules-41-iac-1000; src:usa-il:elevator-act-225-ilcs-312; src:usa-il:osfm-elevator-page | Elevator authority, scope, Chicago exception, local agreements | ChatGPT | 2026-06-25 | 0.78 |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 1.00 | none | Status reflects body now includes official state sources, while unresolved fields remain explicit. |
| report | risk.overall_confidence | 0.68 | verified | 1.00 | none | Conservative confidence due to unresolved local directory, mechanical/fuel-gas private scope, and project-level exceptions. |
| ahj:usa-il:cdb | authority_name | Illinois Capital Development Board | partially_verified | 0.88 | src:usa-il:cdb-act-20-ilcs-3105; src:usa-il:cdb-building-codes-page | Primary statewide coordinating authority for this report's building-code framework. |
| adoption:usa-il:commercial-noncode-ibc-iebc-nec-2025 | base_codes | current/current-preceding IBC/IEBC/NEC plus specialty codes | verified | 0.86 | src:usa-il:cdb-act-20-ilcs-3105 | Statutory code list extracted. |
| adoption:usa-il:local-structural-baseline-2025 | mandatory_date | 2025-01-01 | verified | 0.86 | src:usa-il:cdb-act-20-ilcs-3105 | Statutory beginning date extracted. |
| adoption:usa-il:energy-2024-iecc | edition/effective_date | 2024 IECC as amended; 2025-11-30 | verified | 0.86 | src:usa-il:energy-code-71-iac-600; src:usa-il:energy-code-cdb-page | Rule and CDB page support current edition/effective date. |
| adoption:usa-il:fire-nfpa101-2015 | edition/effective_date | NFPA 101 2015; 2020-01-01 | partially_verified | 0.82 | src:usa-il:fire-prevention-rules-41-iac-100; src:usa-il:osfm-life-safety-faq | NFPA incorporated text not reproduced; OSFM modifications not parsed. |
| adoption:usa-il:public-schools-2024-icc | design_contract_trigger | 2025-01-01 | verified | 0.84 | src:usa-il:hls-code-public-schools-23-iac-180 | Part 180 text extracted. |
| local-enforcement:usa-il | model | hybrid_local_state_specialty | partially_verified | 0.72 | src:usa-il:cdb-act-20-ilcs-3105; src:usa-il:fire-prevention-rules-41-iac-100; src:usa-il:plumbing-code-77-iac-890; src:usa-il:hls-code-public-schools-23-iac-180; src:usa-il:elevator-safety-rules-41-iac-1000 | Address-level AHJ contacts not populated. |
| local-amendment-rule:usa-il | filing_required | true for local building-code edition/amendment reporting to CDB | verified | 0.82 | src:usa-il:cdb-act-20-ilcs-3105; src:usa-il:cdb-building-codes-page | Approval varies by code family; local building-code reporting is confirmed. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Every `src:usa-il:*` referenced in body has a registry row in Section 8. |
| All authority IDs resolve | pass | Authority IDs are consistent within the report; no separate authority registry exists in this schema. |
| All current code families have adoption records | pass | Every matrix code family has an explicit row; private mechanical/fuel-gas rows are limited-scope and unresolved outside public schools/local codes. |
| Building and operational fire code are separated | pass | Section 3 separates fire construction references from OSFM operational/prevention Life Safety Code. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Distinct date fields are preserved; null remains where adoption date was not extracted. |
| Effective dates are valid ISO dates | pass | Dates entered in date fields use YYYY-MM-DD. |
| No impossible date sequences | pass | Known delayed operative triggers are explained where adoption/effective and operative dates differ. |
| Transition rules have explicit trigger conditions | pass | Permit, design-contract, reporting, occupancy, and funding triggers are captured. |
| Permit-date logic is captured where applicable | pass | CDB commercial non-building-code January 1 permit-application rule is included. |
| Local enforcement model classified | pass | Classified as hybrid_local_state_specialty. |
| Local amendment rule classified | pass | Classified as local_building_codes_allowed_subject_to_reporting_and_state_specialty_constraints. |
| AHJ confirmation metadata present | fail | No address-level AHJ contact records were populated. |
| Official-source caveats captured | pass | Section 8.2 lists caveats by source. |
| Local code directory extracted | fail | CDB directory was identified but not extracted into municipality/county records. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-il:001 | high | local code directory | CDB local code directory was identified but municipality/county code editions and amendments were not extracted. | Pull directory data and normalize local code/adoption/amendment records. | null | null | open |
| issue:usa-il:002 | high | fire local equivalency | OSFM equal-or-higher local fire-code criteria are captured, but no locality-specific equivalency determinations were extracted. | Identify whether OSFM publishes locality-specific determinations or agreements; otherwise require AHJ confirmation. | null | null | open |
| issue:usa-il:003 | medium | private mechanical/fuel gas scope | Statewide private mechanical and fuel-gas adoption outside local codes and public schools was not fully resolved. | Review CDB code directory, local adoption statutes, and any state specialty rules for mechanical/fuel-gas scope. | null | null | open |
| issue:usa-il:004 | medium | elevator incorporated standards | Part 1000 authority/scope/effective date are captured, but incorporated standard editions in the elevator rules were not parsed. | Extract 41 Ill. Adm. Code 1000.60 and related incorporated standards. | null | null | open |
| issue:usa-il:005 | medium | Article 34 and Chicago exceptions | Public-school Article 34 exception and Chicago elevator exception require address/project-level routing. | Build resolver flags for Chicago and Article 34 school districts. | null | null | open |
| issue:usa-il:006 | low | accessibility enforcement details | Accessibility adoption and AG enforcement are confirmed, but CDB enforcement/interpretation workflows were not fully extracted. | Extract CDB accessibility interpretation and enforcement pages. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-il:cdb-building-codes | src:usa-il:cdb-building-codes-page | html_diff | monthly | CDB changes statewide framework, local directory links, or code summaries | 2026-06-25 | active |
| watch:usa-il:cdb-act | src:usa-il:cdb-act-20-ilcs-3105 | statute_diff | quarterly | Amendments to 20 ILCS 3105/10.09-1 or 10.18 | 2026-06-25 | active |
| watch:usa-il:energy-code | src:usa-il:energy-code-71-iac-600 | admin_rule_diff | monthly | 2024 IECC amendments, 2026 stretch-code proposals, or Part 600 updates | 2026-06-25 | active |
| watch:usa-il:stretch-act | src:usa-il:energy-efficient-building-act-20-ilcs-3125 | statute_diff | quarterly | Stretch Energy Code milestones or local energy preemption changes | 2026-06-25 | active |
| watch:usa-il:osfm-fire | src:usa-il:fire-prevention-rules-41-iac-100 | admin_rule_diff | quarterly | NFPA 101 edition update or OSFM modifications/local equivalency criteria changes | 2026-06-25 | active |
| watch:usa-il:schools-hls | src:usa-il:hls-code-public-schools-23-iac-180 | admin_rule_diff | quarterly | ISBE code edition updates, applicability changes, or public-school substitutions | 2026-06-25 | active |
| watch:usa-il:elevator | src:usa-il:elevator-safety-rules-41-iac-1000 | admin_rule_diff | quarterly | Refiled elevator amendments or incorporated-standard updates | 2026-06-25 | active |
| watch:usa-il:plumbing | src:usa-il:plumbing-code-77-iac-890 | admin_rule_diff | semiannual | Plumbing Code amendments or IDPH variance/enforcement changes | 2026-06-25 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-24 | Generated baseline draft report | report:usa-il | none | Codex | Baseline file contained state metadata but no official source-backed Illinois code analysis. |
| 2026-06-25 | Populated Illinois report with official source-backed authority, adoption, date-rule, local enforcement, amendment, QA, and monitoring sections | report:usa-il; ahj:usa-il:cdb; adoption:usa-il:commercial-noncode-ibc-iebc-nec-2025; adoption:usa-il:energy-2024-iecc; adoption:usa-il:fire-nfpa101-2015; local-enforcement:usa-il; local-amendment-rule:usa-il | src:usa-il:cdb-act-20-ilcs-3105; src:usa-il:cdb-building-codes-page; src:usa-il:energy-code-71-iac-600; src:usa-il:fire-prevention-rules-41-iac-100; src:usa-il:plumbing-code-77-iac-890; src:usa-il:hls-code-public-schools-23-iac-180; src:usa-il:elevator-safety-rules-41-iac-1000 | ChatGPT | Upgraded status to partially_verified after validation pass; unresolved items remain explicit. |
