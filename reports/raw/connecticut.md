---
state:
  state_id: "US-CT"
  name: "Connecticut"
  abbreviation: "CT"
report:
  report_id: "state-report:usa-ct"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-25"
  last_verified: "2026-06-25"
  reviewed_by: null
risk:
  overall_confidence: 0.88 # 0.00 - 1.00
  risk_flags:
    - "pending_2026_code_cycle"
    - "local_amendment_registry_not_verified"
    - "ahj_contacts_not_populated"
    - "elevator_code_not_fully_parsed"
  open_questions_count: 4

---

# State Building Code Authority Report: Connecticut

## 1. Executive Summary

- **Authority model:** Hybrid statewide/local model. Connecticut adopts statewide building, fire safety, and fire prevention codes through state authorities, while municipal building officials and local fire marshals perform most front-line administration and enforcement.

- **Statewide code status:** Current core statewide codes are verified. The 2022 Connecticut State Building Code applies to projects with permit applications filed from October 1, 2022. The 2022 Connecticut State Fire Safety Code and 2022 Connecticut State Fire Prevention Code are also effective October 1, 2022.

- **Local enforcement model:** Building-code administration is local-first through municipal building officials, with state-level interpretation, modification, variation, exemption, and appeal paths. Fire-code administration uses local fire marshals for initial compliance decisions in most cases, with the State Fire Marshal retaining review, state-property, interpretation, and modification roles.

- **Local amendment posture:** The State Building Code is the building code for all towns, cities, and boroughs. Local governments and interested persons may propose amendments to the State Building Code, including municipality-specific amendments when special local conditions are established, but amendments must be adopted through the state statutory procedure. Fire-prevention local ordinance authority was verified only for permit, certificate, notice, approval, order, and fee-schedule provisions where the State Fire Prevention Code allows it; a statewide local-amendment registry was not verified.

- **Known transition periods or pending changes:** Connecticut has an active 2026 code cycle. The next State Building Code, State Fire Safety Code, and State Fire Prevention Code were expected in mid-2026, but DAS reports that approval by the Legislative Regulation Review Committee was still pending and the effective date may be delayed.

- **Production readiness:** usable_with_caution

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: The State Building Inspector and Codes and Standards Committee, with approval
    of the Commissioner of Administrative Services, adopt and administer the State
    Building Code.
  confidence: 0.97
  source_ids:
  - src:usa-ct:statute:chap-541
- topic: Building-code statewide scope
  finding: The State Building Code, including state-adopted amendments, is the building
    code for all towns, cities, boroughs, state agencies, the Connecticut Airport
    Authority, and the Connecticut Port Authority.
  confidence: 0.97
  source_ids:
  - src:usa-ct:statute:chap-541
  - src:usa-ct:code:2022-csbc-pdf
- topic: Current building-code edition
  finding: The current State Building Code is the 2022 Connecticut State Building
    Code, based on 2021 International Codes, 2020 NFPA 70, and 2017 ICC A117.1; it
    applies to permit applications filed from 2022-10-01.
  confidence: 0.96
  source_ids:
  - src:usa-ct:agency:building-code-page
  - src:usa-ct:code:2022-csbc-pdf
- topic: Fuel gas treatment
  finding: Connecticut does not adopt the International Fuel Gas Code as a separate
    state code; fuel-gas references are redirected to NFPA 2, NFPA 54, and NFPA 58
    as adopted in the 2022 fire codes.
  confidence: 0.91
  source_ids:
  - src:usa-ct:code:2022-csbc-pdf
- topic: Current fire safety code
  finding: The 2022 Connecticut State Fire Safety Code is effective 2022-10-01 and
    adopts the 2021 IFC as Part III and 2021 NFPA 101 as Part IV, with a January 1,
    2006 applicability split.
  confidence: 0.95
  source_ids:
  - src:usa-ct:agency:fire-codes-docs
  - src:usa-ct:code:2022-csfsc-pdf
- topic: Current fire prevention code
  finding: The 2022 Connecticut State Fire Prevention Code is effective 2022-10-01
    and is based on NFPA 1, Fire Code, 2021 edition, with Connecticut amendments.
  confidence: 0.96
  source_ids:
  - src:usa-ct:agency:fire-codes-overview
  - src:usa-ct:code:2022-csfpc-pdf
- topic: Effective-date publication rule
  finding: Approved or deemed-approved State Building Code, State Fire Prevention
    Code, and Fire Safety Code provisions become effective and enforceable upon posting
    on the DAS website unless a later date is required by statute or specified in
    the code.
  confidence: 0.97
  source_ids:
  - src:usa-ct:statute:chap-541
- topic: Local enforcement
  finding: Municipal building officials administer the building code locally; local
    fire marshals inspect occupancies regulated by the fire codes and make most initial
    fire-code compliance decisions.
  confidence: 0.94
  source_ids:
  - src:usa-ct:statute:chap-541
  - src:usa-ct:code:2022-csfpc-pdf
- topic: Elevator / conveyance code
  finding: OSBI administers elevators and escalators in addition to the State Building
    Code; Connecticut adopted ASME A17.1-2013 effective 2018-01-03.
  confidence: 0.82
  source_ids:
  - src:usa-ct:agency:elevator-code-page
  - src:usa-ct:agency:boiler-elevator-inspections
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-ct:state-building-inspector |
| Authority name | State Building Inspector and Codes and Standards Committee, with approval of the Commissioner of Administrative Services |
| Authority type | state_office_and_committee |
| Parent agency | Department of Administrative Services |
| Legal basis | Conn. Gen. Stat. §§ 29-250, 29-251, 29-252, 29-252a, 29-252b, 29-253, 29-254, 29-260 |
| Role | Adopts, administers, amends, interprets, and handles state-level variations, exemptions, equivalent compliance, and appeals for the State Building Code |
| Enforcement model | hybrid |
| Source IDs | src:usa-ct:statute:chap-541; src:usa-ct:code:2022-csbc-pdf |
| Verification status | verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-ct:state-building-inspector | State Building Inspector and Codes and Standards Committee | Adopts and administers the State Building Code | Conn. Gen. Stat. § 29-252 | src:usa-ct:statute:chap-541; src:usa-ct:code:2022-csbc-pdf | verified |
| Residential | ahj:usa-ct:state-building-inspector | State Building Inspector and Codes and Standards Committee | Adopts residential provisions through the 2021 IRC portion of the 2022 Connecticut State Building Code | Conn. Gen. Stat. §§ 29-252, 29-253 | src:usa-ct:agency:building-code-page; src:usa-ct:code:2022-csbc-pdf | verified |
| Existing Building / Rehabilitation | ahj:usa-ct:state-building-inspector | State Building Inspector and Codes and Standards Committee | Adopts existing-building and rehabilitation provisions, including a required rehabilitation subcode | Conn. Gen. Stat. §§ 29-252, 29-256 | src:usa-ct:statute:chap-541; src:usa-ct:code:2022-csbc-pdf | verified |
| Mechanical | ahj:usa-ct:state-building-inspector | State Building Inspector and Codes and Standards Committee | Adopts mechanical provisions through the 2021 IMC portion of the 2022 Connecticut State Building Code | Conn. Gen. Stat. § 29-252 | src:usa-ct:agency:building-code-page; src:usa-ct:code:2022-csbc-pdf | verified |
| Plumbing | ahj:usa-ct:state-building-inspector | State Building Inspector and Codes and Standards Committee | Adopts plumbing provisions through the 2021 IPC portion of the 2022 Connecticut State Building Code | Conn. Gen. Stat. § 29-252 | src:usa-ct:agency:building-code-page; src:usa-ct:code:2022-csbc-pdf | verified |
| Fuel Gas | ahj:usa-ct:state-building-inspector | State Building Inspector and Codes and Standards Committee, with fire-code cross-reference | Connecticut does not separately adopt IFGC; gas requirements are addressed through NFPA 2, NFPA 54, and NFPA 58 references tied to the State Fire Safety and Fire Prevention Codes | Conn. Gen. Stat. §§ 29-252, 29-291a, 29-292 | src:usa-ct:code:2022-csbc-pdf; src:usa-ct:code:2022-csfpc-pdf | partially_verified |
| Electrical | ahj:usa-ct:state-building-inspector | State Building Inspector and Codes and Standards Committee | Adopts electrical provisions through the 2020 NFPA 70 / National Electrical Code portion of the 2022 Connecticut State Building Code | Conn. Gen. Stat. § 29-252 | src:usa-ct:agency:building-code-page; src:usa-ct:code:2022-csbc-pdf | verified |
| Energy | ahj:usa-ct:state-building-inspector | State Building Inspector and Codes and Standards Committee | Adopts energy provisions through the 2021 IECC portion of the 2022 Connecticut State Building Code and statutory energy-efficiency revision requirements | Conn. Gen. Stat. §§ 29-252, 29-256a | src:usa-ct:statute:chap-541; src:usa-ct:agency:building-code-page; src:usa-ct:code:2022-csbc-pdf | verified |
| Swimming Pool and Spa | ahj:usa-ct:state-building-inspector | State Building Inspector and Codes and Standards Committee | Adopts the 2021 International Swimming Pool and Spa Code portion of the 2022 Connecticut State Building Code | Conn. Gen. Stat. § 29-252 | src:usa-ct:agency:building-code-page; src:usa-ct:code:2022-csbc-pdf | verified |
| Fire - construction / life safety | ahj:usa-ct:state-fire-marshal | State Fire Marshal and Codes and Standards Committee | Adopts and administers the Fire Safety Code, including Part III IFC and Part IV NFPA 101 provisions | Conn. Gen. Stat. §§ 29-292, 29-292a | src:usa-ct:statute:chap-541; src:usa-ct:agency:fire-codes-docs; src:usa-ct:code:2022-csfsc-pdf | verified |
| Fire - operational / prevention | ahj:usa-ct:state-fire-marshal | State Fire Marshal, in coordination with the advisory committee | Adopts and administers the State Fire Prevention Code | Conn. Gen. Stat. §§ 29-291a, 29-291e | src:usa-ct:statute:chap-541; src:usa-ct:agency:fire-codes-overview; src:usa-ct:code:2022-csfpc-pdf | verified |
| Accessibility | ahj:usa-ct:state-building-inspector | State Building Inspector and Codes and Standards Committee | Adopts accessibility provisions through 2017 ICC/ANSI A117.1 and state accessibility statutes | Conn. Gen. Stat. §§ 29-269 to 29-275a | src:usa-ct:statute:chap-541; src:usa-ct:agency:building-code-page; src:usa-ct:code:2022-csbc-pdf | verified |
| Elevator / Conveyance | ahj:usa-ct:osbi-elevators | Office of the State Building Inspector | Regulates elevators and escalators and administers ASME A17.1 Safety Code for Elevators and Escalators | Conn. Gen. Stat. § 29-192 | src:usa-ct:agency:elevator-code-page; src:usa-ct:agency:boiler-elevator-inspections; src:usa-ct:agency:modifications-waivers | partially_verified |

### 2.3 Authority Hierarchy Notes

Connecticut uses a statewide code baseline with local execution. The State Building Code is the code for all municipalities, state agencies, the Connecticut Airport Authority, and the Connecticut Port Authority. Local building officials administer the code for their municipalities, but local policies and procedures cannot waive code requirements or create requirements beyond the code. Fire-code administration is similarly split: local fire marshals make most initial determinations and conduct inspections, while the State Fire Marshal reviews local fire marshal decisions, handles state-owned property, and administers modifications and variations where statutes or codes reserve state action.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-ct:001 | ahj:usa-ct:state-building-inspector | adopts_and_administers | State Building Code | src:usa-ct:statute:chap-541; src:usa-ct:code:2022-csbc-pdf | verified |
| edge:usa-ct:002 | ahj:usa-ct:state-building-inspector | state_code_for | towns_cities_boroughs_and_state_agencies | src:usa-ct:statute:chap-541; src:usa-ct:code:2022-csbc-pdf | verified |
| edge:usa-ct:003 | ahj:usa-ct:state-building-inspector | local_enforcement_by | municipal_building_officials | src:usa-ct:statute:chap-541 | verified |
| edge:usa-ct:004 | ahj:usa-ct:state-fire-marshal | adopts_and_administers | State Fire Prevention Code | src:usa-ct:statute:chap-541; src:usa-ct:code:2022-csfpc-pdf | verified |
| edge:usa-ct:005 | ahj:usa-ct:state-fire-marshal | adopts_with | Codes and Standards Committee for Fire Safety Code | src:usa-ct:statute:chap-541; src:usa-ct:agency:fire-codes-overview | verified |
| edge:usa-ct:006 | local_fire_marshal | initial_determination_for | State Fire Prevention Code compliance in local jurisdiction | src:usa-ct:code:2022-csfpc-pdf; src:usa-ct:statute:chap-541 | verified |
| edge:usa-ct:007 | ahj:usa-ct:state-fire-marshal | reviews | local_fire_marshal_decisions_except_limited_local_ordinances | src:usa-ct:statute:chap-541; src:usa-ct:code:2022-csfpc-pdf | verified |
| edge:usa-ct:008 | ahj:usa-ct:osbi-elevators | administers | ASME A17.1 Safety Code for Elevators and Escalators | src:usa-ct:agency:elevator-code-page; src:usa-ct:agency:boiler-elevator-inspections | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | 2022 Connecticut State Building Code | International Building Code | 2021 | current | 2022-10-01 | 2022-10-01 | 2022-10-01 | 2022-10-01 | Applies to projects with permit applications filed from 2022-10-01. | src:usa-ct:agency:building-code-page; src:usa-ct:code:2022-csbc-pdf |
| Residential | 2022 Connecticut State Building Code | International Residential Code | 2021 | current | 2022-10-01 | 2022-10-01 | 2022-10-01 | 2022-10-01 | Applies to detached one- and two-family dwellings and townhouses within the IRC scope for permit applications filed from 2022-10-01. | src:usa-ct:agency:building-code-page; src:usa-ct:code:2022-csbc-pdf |
| Existing Building / Rehabilitation | 2022 Connecticut State Building Code | International Existing Building Code | 2021 | current | 2022-10-01 | 2022-10-01 | 2022-10-01 | 2022-10-01 | Applies through the CSBC for permit applications filed from 2022-10-01; rehabilitation subcode requirement is statutory. | src:usa-ct:statute:chap-541; src:usa-ct:agency:building-code-page; src:usa-ct:code:2022-csbc-pdf |
| Mechanical | 2022 Connecticut State Building Code | International Mechanical Code | 2021 | current | 2022-10-01 | 2022-10-01 | 2022-10-01 | 2022-10-01 | Applies through the CSBC for permit applications filed from 2022-10-01. | src:usa-ct:agency:building-code-page; src:usa-ct:code:2022-csbc-pdf |
| Plumbing | 2022 Connecticut State Building Code | International Plumbing Code | 2021 | current | 2022-10-01 | 2022-10-01 | 2022-10-01 | 2022-10-01 | Applies through the CSBC for permit applications filed from 2022-10-01. | src:usa-ct:agency:building-code-page; src:usa-ct:code:2022-csbc-pdf |
| Fuel Gas | 2022 Connecticut State Building Code / fire-code cross-reference | NFPA 2, NFPA 54, NFPA 58 references; International Fuel Gas Code not adopted | fire-code editions not fully parsed | current | 2022-10-01 | 2022-10-01 | 2022-10-01 | 2022-10-01 | IFGC references are redirected to NFPA fuel-gas requirements as adopted in the 2022 State Fire Safety Code and State Fire Prevention Code. | src:usa-ct:code:2022-csbc-pdf; src:usa-ct:code:2022-csfpc-pdf |
| Electrical | 2022 Connecticut State Building Code | NFPA 70, National Electrical Code | 2020 | current | 2022-10-01 | 2022-10-01 | 2022-10-01 | 2022-10-01 | Applies through the CSBC for electrical systems for permit applications filed from 2022-10-01. | src:usa-ct:agency:building-code-page; src:usa-ct:code:2022-csbc-pdf |
| Energy | 2022 Connecticut State Building Code | International Energy Conservation Code | 2021 | current | 2022-10-01 | 2022-10-01 | 2022-10-01 | 2022-10-01 | Applies through the CSBC for permit applications filed from 2022-10-01. | src:usa-ct:agency:building-code-page; src:usa-ct:code:2022-csbc-pdf |
| Swimming Pool and Spa | 2022 Connecticut State Building Code | International Swimming Pool and Spa Code | 2021 | current | 2022-10-01 | 2022-10-01 | 2022-10-01 | 2022-10-01 | Applies through the CSBC for permit applications filed from 2022-10-01. | src:usa-ct:agency:building-code-page; src:usa-ct:code:2022-csbc-pdf |
| Accessibility | 2022 Connecticut State Building Code | ICC/ANSI A117.1 Accessible and Usable Buildings and Facilities | 2017 | current | 2022-10-01 | 2022-10-01 | 2022-10-01 | 2022-10-01 | Applies as the accessibility standard incorporated into the CSBC. | src:usa-ct:agency:building-code-page; src:usa-ct:code:2022-csbc-pdf |
| Fire - construction / life safety | 2022 Connecticut State Fire Safety Code, Part III | International Fire Code | 2021 | current | 2022-10-01 | 2022-10-01 | 2022-10-01 | 2022-10-01 | Part III applies to new construction, renovations, additions, or changes of use built or initially permitted on or after 2006-01-01, subject to current-code transition rules. | src:usa-ct:agency:fire-codes-docs; src:usa-ct:code:2022-csfsc-pdf |
| Fire - existing buildings / life safety | 2022 Connecticut State Fire Safety Code, Part IV | NFPA 101, Life Safety Code | 2021 | current | 2022-10-01 | 2022-10-01 | 2022-10-01 | 2022-10-01 | Part IV applies to existing buildings or portions that existed or were initially permitted before 2006-01-01. | src:usa-ct:agency:fire-codes-docs; src:usa-ct:code:2022-csfsc-pdf |
| Fire - operational / prevention | 2022 Connecticut State Fire Prevention Code | NFPA 1, Fire Code | 2021 | current | 2022-10-01 | 2022-10-01 | 2022-10-01 | 2022-10-01 | State Fire Prevention Code applies to operational/prevention provisions; local fire marshal makes initial determinations except where reserved. | src:usa-ct:agency:fire-codes-overview; src:usa-ct:code:2022-csfpc-pdf |
| Elevator / Conveyance | Safety Code for Elevators and Escalators | ASME A17.1 Safety Code for Elevators and Escalators | 2013 | current | 2018-01-03 | 2018-01-03 | 2018-01-03 | 2018-01-03 | DAS page states Connecticut adopted ASME A17.1-2013 on 2018-01-03; full RCSA text was not parsed here. | src:usa-ct:agency:elevator-code-page; src:usa-ct:agency:boiler-elevator-inspections |

### 3.2 Adoption Records

```yaml
adoption_id: "adoption:usa-ct:building:2022-csbc"
state_id: "US-CT"
code_family: "building"
status: "current"
state_code:
  name: "2022 Connecticut State Building Code"
  edition_label: "2022"
  codification: "Conn. Gen. Stat. ch. 541, especially §§ 29-252 to 29-254; DAS-posted 2022 CSBC PDF"
base_model_code:
  publisher: "International Code Council / NFPA / ICC A117.1"
  code_name:
    - "2021 International Building Code"
    - "2021 International Residential Code"
    - "2021 International Existing Building Code"
    - "2021 International Plumbing Code"
    - "2021 International Mechanical Code"
    - "2021 International Energy Conservation Code"
    - "2021 International Swimming Pool and Spa Code"
    - "2020 NFPA 70, National Electrical Code"
    - "2017 ICC/ANSI A117.1 Accessible and Usable Buildings and Facilities"
  edition_year: 2022
  incorporated_by_reference: true
authority:
  adopting_authority_id: "ahj:usa-ct:state-building-inspector"
  enforcing_authority_model: "hybrid"
  interpretation_authority_id: "ahj:usa-ct:state-building-inspector"
dates:
  adoption_date: "2022-10-01"
  effective_date: "2022-10-01"
  operative_date: "2022-10-01"
  mandatory_date: "2022-10-01"
  replacement_date: null
applicability:
  date_trigger: "permit_application_date"
  applies_to:
    - "new_construction"
    - "alteration"
    - "repair"
    - "addition"
    - "change_of_occupancy"
    - "state_owned"
    - "municipalities"
  exclusions:
    - "International Fuel Gas Code is not separately adopted"
  special_conditions:
    - "Applies to projects with permit applications filed from 2022-10-01."
    - "Official electronic version is the DAS-posted code unless a later date is required by statute or specified in the code."
transition:
  exists: true
  rule_id: "date-rule:usa-ct:building:2022-permit-application"
  start_date: "2022-10-01"
  end_date: null
  prior_code_allowed: false
  prior_code_condition: "2018 CSBC applies to permit applications filed from 2018-10-01 through 2022-09-30; 2022 CSBC applies from 2022-10-01."
amendments:
  state_amended: true
  amendment_set_ids:
    - "amendment-set:usa-ct:building:2022-csbc"
  amendment_source_ids:
    - "src:usa-ct:code:2022-csbc-pdf"
provenance:
  source_ids:
    - "src:usa-ct:statute:chap-541"
    - "src:usa-ct:agency:building-code-page"
    - "src:usa-ct:code:2022-csbc-pdf"
  field_sources:
    state_code.name: ["src:usa-ct:agency:building-code-page", "src:usa-ct:code:2022-csbc-pdf"]
    base_model_code.code_name: ["src:usa-ct:agency:building-code-page", "src:usa-ct:code:2022-csbc-pdf"]
    authority.adopting_authority_id: ["src:usa-ct:statute:chap-541", "src:usa-ct:code:2022-csbc-pdf"]
    dates.effective_date: ["src:usa-ct:agency:building-code-page", "src:usa-ct:code:2022-csbc-pdf"]
    applicability.date_trigger: ["src:usa-ct:agency:building-code-page", "src:usa-ct:code:2022-csbc-pdf"]
verification:
  status: "verified"
  confidence: 0.96
  notes: "Current edition, base model codes, and permit-application transition date verified from official DAS page and DAS-posted code PDF."

adoption_id: "adoption:usa-ct:fire-safety:2022-csfsc"
state_id: "US-CT"
code_family: "fire_life_safety"
status: "current"
state_code:
  name: "2022 Connecticut State Fire Safety Code"
  edition_label: "2022"
  codification: "Conn. Gen. Stat. §§ 29-292, 29-292a; DAS-posted 2022 CSFSC PDF"
base_model_code:
  publisher: "International Code Council / NFPA"
  code_name:
    - "2021 International Fire Code, Part III"
    - "2021 NFPA 101, Life Safety Code, Part IV"
  edition_year: 2022
  incorporated_by_reference: true
authority:
  adopting_authority_id: "ahj:usa-ct:state-fire-marshal"
  enforcing_authority_model: "hybrid"
  interpretation_authority_id: "ahj:usa-ct:state-fire-marshal"
dates:
  adoption_date: "2022-10-01"
  effective_date: "2022-10-01"
  operative_date: "2022-10-01"
  mandatory_date: "2022-10-01"
  replacement_date: null
applicability:
  date_trigger: "building_permit_application_date_for_part_split"
  applies_to:
    - "new_construction"
    - "renovation"
    - "addition"
    - "change_of_use"
    - "existing_buildings"
  exclusions: []
  special_conditions:
    - "Part III applies to buildings or portions for which an initial building permit application was made on or after 2006-01-01."
    - "Part IV applies to occupancies and uses located in buildings, structures, or portions that existed before 2006-01-01."
transition:
  exists: true
  rule_id: "date-rule:usa-ct:fire-safety:2006-part-split"
  start_date: "2022-10-01"
  end_date: null
  prior_code_allowed: true
  prior_code_condition: "Some existing portions retain Part IV treatment based on pre-2006 existence or permit status; additions, alterations, renovations, and changes of use are handled under Part III rules as stated in the code."
amendments:
  state_amended: true
  amendment_set_ids:
    - "amendment-set:usa-ct:fire-safety:2022-csfsc"
  amendment_source_ids:
    - "src:usa-ct:code:2022-csfsc-pdf"
provenance:
  source_ids:
    - "src:usa-ct:statute:chap-541"
    - "src:usa-ct:agency:fire-codes-docs"
    - "src:usa-ct:code:2022-csfsc-pdf"
  field_sources:
    state_code.name: ["src:usa-ct:agency:fire-codes-docs", "src:usa-ct:code:2022-csfsc-pdf"]
    base_model_code.code_name: ["src:usa-ct:agency:fire-codes-docs", "src:usa-ct:code:2022-csfsc-pdf"]
    authority.adopting_authority_id: ["src:usa-ct:statute:chap-541", "src:usa-ct:agency:fire-codes-overview"]
    dates.effective_date: ["src:usa-ct:agency:fire-codes-docs", "src:usa-ct:code:2022-csfsc-pdf"]
    applicability.date_trigger: ["src:usa-ct:agency:fire-codes-docs", "src:usa-ct:code:2022-csfsc-pdf"]
verification:
  status: "verified"
  confidence: 0.95
  notes: "Current edition, base model codes, effective date, and 2006 Part III/Part IV applicability split verified."

adoption_id: "adoption:usa-ct:fire-prevention:2022-csfpc"
state_id: "US-CT"
code_family: "fire_prevention"
status: "current"
state_code:
  name: "2022 Connecticut State Fire Prevention Code"
  edition_label: "2022"
  codification: "Conn. Gen. Stat. §§ 29-291a, 29-291e; DAS-posted 2022 CSFPC PDF"
base_model_code:
  publisher: "National Fire Protection Association"
  code_name: "NFPA 1, Fire Code"
  edition_year: 2021
  incorporated_by_reference: true
authority:
  adopting_authority_id: "ahj:usa-ct:state-fire-marshal"
  enforcing_authority_model: "hybrid"
  interpretation_authority_id: "ahj:usa-ct:state-fire-marshal"
dates:
  adoption_date: "2022-10-01"
  effective_date: "2022-10-01"
  operative_date: "2022-10-01"
  mandatory_date: "2022-10-01"
  replacement_date: null
applicability:
  date_trigger: "code_effective_date_with_local_initial_determination"
  applies_to:
    - "fire_prevention"
    - "fire_hazard_abatement"
    - "maintenance_and_operations"
    - "state_owned"
  exclusions:
    - "one- and two-family dwellings and qualifying townhouses except where the code specifies"
  special_conditions:
    - "State Fire Marshal is AHJ for administration, application, interpretation, and modification."
    - "Local fire marshal makes initial compliance determination except where reserved."
transition:
  exists: true
  rule_id: "date-rule:usa-ct:fire-prevention:posting"
  start_date: "2022-10-01"
  end_date: null
  prior_code_allowed: false
  prior_code_condition: "The current posted code controls unless a later date is required by statute or specified in the code."
amendments:
  state_amended: true
  amendment_set_ids:
    - "amendment-set:usa-ct:fire-prevention:2022-csfpc"
  amendment_source_ids:
    - "src:usa-ct:code:2022-csfpc-pdf"
provenance:
  source_ids:
    - "src:usa-ct:statute:chap-541"
    - "src:usa-ct:agency:fire-codes-overview"
    - "src:usa-ct:code:2022-csfpc-pdf"
  field_sources:
    state_code.name: ["src:usa-ct:code:2022-csfpc-pdf"]
    base_model_code.code_name: ["src:usa-ct:agency:fire-codes-overview", "src:usa-ct:code:2022-csfpc-pdf"]
    authority.adopting_authority_id: ["src:usa-ct:statute:chap-541", "src:usa-ct:code:2022-csfpc-pdf"]
    dates.effective_date: ["src:usa-ct:code:2022-csfpc-pdf"]
verification:
  status: "verified"
  confidence: 0.96
  notes: "Current edition, base NFPA 1 edition, effective date, and AHJ/local-fire-marshal split verified."

adoption_id: "adoption:usa-ct:elevator:asme-a17-1-2013"
state_id: "US-CT"
code_family: "elevator_conveyance"
status: "current"
state_code:
  name: "Safety Code for Elevators and Escalators"
  edition_label: "ASME A17.1-2013"
  codification: "Conn. Gen. Stat. § 29-192; RCSA §§ 29-192-1e to 29-192-16e cited by DAS history page"
base_model_code:
  publisher: "ASME"
  code_name: "ASME A17.1 Safety Code for Elevators and Escalators"
  edition_year: 2013
  incorporated_by_reference: true
authority:
  adopting_authority_id: "ahj:usa-ct:osbi-elevators"
  enforcing_authority_model: "state_administered"
  interpretation_authority_id: "ahj:usa-ct:osbi-elevators"
dates:
  adoption_date: "2018-01-03"
  effective_date: "2018-01-03"
  operative_date: "2018-01-03"
  mandatory_date: "2018-01-03"
  replacement_date: null
applicability:
  date_trigger: "equipment_registration_or_inspection_date"
  applies_to:
    - "elevators"
    - "escalators"
    - "moving_walks"
  exclusions: []
  special_conditions:
    - "Full RCSA text and later amendments were not parsed."
transition:
  exists: true
  rule_id: "date-rule:usa-ct:elevator:2018-a17-1"
  start_date: "2018-01-03"
  end_date: null
  prior_code_allowed: unresolved
  prior_code_condition: "Historic DAS page lists prior editions; current transition exceptions were not parsed."
amendments:
  state_amended: unresolved
  amendment_set_ids: []
  amendment_source_ids:
    - "src:usa-ct:agency:elevator-code-page"
provenance:
  source_ids:
    - "src:usa-ct:agency:elevator-code-page"
    - "src:usa-ct:agency:boiler-elevator-inspections"
  field_sources:
    state_code.name: ["src:usa-ct:agency:elevator-code-page", "src:usa-ct:agency:boiler-elevator-inspections"]
    base_model_code.code_name: ["src:usa-ct:agency:elevator-code-page"]
    dates.effective_date: ["src:usa-ct:agency:elevator-code-page"]
verification:
  status: "partially_verified"
  confidence: 0.82
  notes: "Official DAS pages verify current ASME A17.1-2013 adoption and OSBI administration; complete regulation text was not parsed."
```

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

Connecticut has two date-rule layers. First, state code adoptions and amendments become effective and enforceable when posted on the Department of Administrative Services website unless a later date is required by statute or by the code. Second, each code edition can define a project-trigger rule. For the 2022 Connecticut State Building Code, the trigger is permit applications filed from 2022-10-01. For the 2022 Connecticut State Fire Safety Code, the current edition is effective 2022-10-01, but Part III and Part IV use a January 1, 2006 building-permit/existing-building split.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-ct:building:posting | State Building Code and amendments | effective_date | upon DAS posting, unless later date applies | Code approved or deemed approved under Conn. Gen. Stat. § 29-252b and posted on DAS website | false | src:usa-ct:statute:chap-541 | verified |
| date-rule:usa-ct:building:2022-permit-application | 2022 Connecticut State Building Code | operative_date / mandatory_date | from 2022-10-01 | Permit application filed on or after 2022-10-01 | false | src:usa-ct:agency:building-code-page; src:usa-ct:code:2022-csbc-pdf | verified |
| date-rule:usa-ct:building:2018-to-2022-window | 2018 Connecticut State Building Code | transition_window | 2018-10-01 to 2022-09-30 | Permit application filed during the 2018 CSBC window | true | src:usa-ct:agency:building-code-page | verified |
| date-rule:usa-ct:fire-prevention:posting | State Fire Prevention Code and amendments | effective_date | upon DAS posting, unless later date applies | Code approved or deemed approved under Conn. Gen. Stat. § 29-291e and posted on DAS website | false | src:usa-ct:statute:chap-541 | verified |
| date-rule:usa-ct:fire-safety:posting | Fire Safety Code and amendments | effective_date | upon DAS posting, unless later date applies | Code approved or deemed approved under Conn. Gen. Stat. § 29-292a and posted on DAS website | false | src:usa-ct:statute:chap-541 | verified |
| date-rule:usa-ct:fire-safety:2022-effective | 2022 Connecticut State Fire Safety Code | effective_date | 2022-10-01 | Current CSFSC edition effective date | false | src:usa-ct:agency:fire-codes-docs; src:usa-ct:code:2022-csfsc-pdf | verified |
| date-rule:usa-ct:fire-safety:2006-part-split | 2022 Connecticut State Fire Safety Code Parts III and IV | applicability_split | 2006-01-01 | Initial building permit application date or existing condition before 2006-01-01 | true | src:usa-ct:agency:fire-codes-docs; src:usa-ct:code:2022-csfsc-pdf | verified |
| date-rule:usa-ct:fire-prevention:2022-effective | 2022 Connecticut State Fire Prevention Code | effective_date | 2022-10-01 | Current CSFPC edition effective date | false | src:usa-ct:code:2022-csfpc-pdf | verified |
| date-rule:usa-ct:elevator:2018-a17-1 | Safety Code for Elevators and Escalators | effective_date | 2018-01-03 | Adoption of ASME A17.1-2013 by Connecticut | unresolved | src:usa-ct:agency:elevator-code-page | partially_verified |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | 2026 Connecticut State Building Code | 2024-01-10 code-amendment-subcommittee process start | null | null | null | null | active_monitoring | src:usa-ct:agency:code-adoption-process | DAS states the next codes were expected in mid-2026, but approval remained pending and the effective date may be delayed. |
| Fire - construction / life safety | 2026 Connecticut State Fire Safety Code | 2024-01-10 code-amendment-subcommittee process start | null | null | null | null | active_monitoring | src:usa-ct:agency:code-adoption-process | Future effective date unresolved. |
| Fire - operational / prevention | 2026 Connecticut State Fire Prevention Code | 2024-01-10 code-amendment-subcommittee process start | null | null | null | null | active_monitoring | src:usa-ct:agency:code-adoption-process | Future effective date unresolved. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-ct:building:state-buildings | building | state, Connecticut Airport Authority, and Connecticut Port Authority buildings | state building permit / state review path | The State Building Code applies to state agencies, the Connecticut Airport Authority, and the Connecticut Port Authority; covered state projects use the State Building Inspector review and permit path. | src:usa-ct:statute:chap-541 | verified |
| applicability-rule:usa-ct:building:gas | fuel_gas | fuel-gas systems | referenced code within CSBC | IFGC is not adopted; references route to NFPA 2, NFPA 54, and NFPA 58 requirements as adopted in fire codes. | src:usa-ct:code:2022-csbc-pdf | verified |
| applicability-rule:usa-ct:fire-safety:part-iii | fire_life_safety | new construction, renovations, additions, changes of use | initial building permit on or after 2006-01-01 | Part III / 2021 IFC provisions apply to the post-2006 scope described in the CSFSC. | src:usa-ct:agency:fire-codes-docs; src:usa-ct:code:2022-csfsc-pdf | verified |
| applicability-rule:usa-ct:fire-safety:part-iv | fire_life_safety | existing buildings / occupancies | building or portion existed or initial permit application was before 2006-01-01 | Part IV / 2021 NFPA 101 provisions apply to the pre-2006 scope described in the CSFSC. | src:usa-ct:agency:fire-codes-docs; src:usa-ct:code:2022-csfsc-pdf | verified |
| applicability-rule:usa-ct:fire-prevention:state-property | fire_prevention | state-owned property | state ownership | The State Fire Marshal makes CSFPC compliance determinations on state-owned property. | src:usa-ct:code:2022-csfpc-pdf | verified |
| applicability-rule:usa-ct:local-fire-ordinance | fire_prevention | local permit, certificate, notice, approval, order, and fee-schedule provisions | local ordinance where enabled by CSFPC § 1.12 | A municipality or fire district may establish specified permit and fee procedures by ordinance; decisions under such local ordinances are treated differently from ordinary CSFPC local determinations. | src:usa-ct:code:2022-csfpc-pdf | verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** state-level amendments to incorporated national model codes, published as DAS-posted code PDFs with Connecticut-specific amendment conventions.

**Where amendments are published:** Department of Administrative Services code pages and code PDFs, with statutory adoption and posting rules in Conn. Gen. Stat. chapter 541.

**Amendment parsing status:** partial. High-level amendment model and several high-impact amendments were parsed; complete section-by-section amendments were not normalized.

### 5.2 State Amendment Sources

| Amendment Set ID | Code Family | Publication Format | Source IDs | Parsed? | Notes |
| --- | --- | --- | --- | --- | --- |
| amendment-set:usa-ct:building:2022-csbc | building / residential / existing building / plumbing / mechanical / energy / electrical / swimming pool and spa / accessibility | DAS-posted PDF with errata | src:usa-ct:agency:building-code-page; src:usa-ct:code:2022-csbc-pdf | partial | Includes amendments to 2021 IBC, 2017 ICC A117.1, 2021 IEBC, 2021 IPC, 2021 IMC, 2021 IECC, 2020 NFPA 70, 2021 ISPSC, and 2021 IRC. |
| amendment-set:usa-ct:fire-safety:2022-csfsc | fire_life_safety | DAS-posted PDF with errata | src:usa-ct:agency:fire-codes-docs; src:usa-ct:code:2022-csfsc-pdf | partial | Includes Part III amendments to 2021 IFC and Part IV amendments to 2021 NFPA 101. |
| amendment-set:usa-ct:fire-prevention:2022-csfpc | fire_prevention | DAS-posted PDF with errata | src:usa-ct:agency:fire-codes-overview; src:usa-ct:code:2022-csfpc-pdf | partial | Includes Connecticut amendments to NFPA 1, Fire Code, 2021 edition. |
| amendment-set:usa-ct:elevator:a17-1-2013 | elevator_conveyance | DAS code page / RCSA adoption history | src:usa-ct:agency:elevator-code-page | no | RCSA text and any Connecticut-specific amendments were not parsed. |

### 5.3 High-Impact State Amendments

| Code Family | Section | Amendment Type | Summary | Source IDs | Confidence |
| --- | --- | --- | --- | --- | --- |
| building | CSBC 101.1 / 101.1.1 | modify / add | 2021 IBC as amended is the 2021 IBC portion of the 2022 CSBC and applies as the building code for towns, cities, boroughs, and state agencies. | src:usa-ct:code:2022-csbc-pdf | 0.96 |
| fuel_gas | CSBC 101.4.1 and parallel provisions | delete / redirect | IFGC is not adopted in Connecticut; fuel-gas references route to NFPA 2, NFPA 54, and NFPA 58 as adopted in the 2022 fire codes. | src:usa-ct:code:2022-csbc-pdf | 0.91 |
| electrical | CSBC 101.4.8 and NFPA 70 amendment section | incorporate / amend | 2020 NFPA 70 applies to electrical system installations and is amended in the 2022 CSBC. | src:usa-ct:code:2022-csbc-pdf | 0.95 |
| energy | CSBC IECC provisions | incorporate / amend | 2021 IECC commercial and residential provisions are incorporated into the 2022 CSBC with Connecticut amendments. | src:usa-ct:agency:building-code-page; src:usa-ct:code:2022-csbc-pdf | 0.95 |
| fire_life_safety | CSFSC Part III / Part IV | incorporate / split applicability | 2022 CSFSC adopts 2021 IFC as Part III for post-2006 new/changed construction scope and 2021 NFPA 101 as Part IV for pre-2006 existing-building scope. | src:usa-ct:agency:fire-codes-docs; src:usa-ct:code:2022-csfsc-pdf | 0.95 |
| fire_prevention | CSFPC 1.0.1 | incorporate / amend | 2022 CSFPC adopts NFPA 1, Fire Code, 2021 edition, with Connecticut amendments. | src:usa-ct:code:2022-csfpc-pdf | 0.96 |
| fire_prevention | CSFPC 1.4.6 | modify | State Fire Marshal is the AHJ for administration, application, interpretation, and modification; local fire marshal makes initial compliance determinations except where reserved. | src:usa-ct:code:2022-csfpc-pdf | 0.94 |
| building / fire_life_safety | Conn. Gen. Stat. § 29-256f | statutory amendment directive | Connecticut directed amendments regarding single-exit-stairway residential occupancies and three-unit/four-unit residential buildings. | src:usa-ct:statute:chap-541 | 0.88 |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-ct"
  model: "hybrid"
  enforcing_entities:
    - "municipal_building_official"
    - "local_fire_marshal"
    - "State Building Inspector for reserved state-building and modification matters"
    - "State Fire Marshal for state-owned property, reviews, and modifications"
  required_officials:
    - "building_official"
    - "assistant_building_official_where_applicable"
    - "local_fire_marshal"
    - "fire_inspector_where_applicable"
  state_reserved_activities:
    - "adoption_and_amendment_of_state_codes"
    - "official_state_code_posting"
    - "state_building_projects_with_state_review_trigger"
    - "State Building Code variations, exemptions, equivalent or alternate compliance"
    - "State Fire Prevention Code official interpretation and review of local decisions"
    - "State Fire Safety Code / Fire Prevention Code variations and exemptions"
    - "elevator_and_escalator_code administration"
  local_activities:
    - "building permit administration by municipal building official"
    - "building code inspection and local code administration"
    - "local fire inspections and initial fire-code compliance determinations"
    - "local fire-prevention permits and fee schedules where authorized by ordinance"
  source_ids:
    - "src:usa-ct:statute:chap-541"
    - "src:usa-ct:code:2022-csbc-pdf"
    - "src:usa-ct:code:2022-csfpc-pdf"
  verification_status: "verified"
  confidence: 0.94
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-ct"
  model: "statewide_code_with_state_amendment_proposal_path"
  applies_to_code_families:
    - "building"
    - "fire_prevention_limited_local_ordinance_provisions"
  local_unilateral_code_amendments_allowed: false
  proposal_allowed: true
  proposal_eligible_parties:
    - "town"
    - "city"
    - "borough"
    - "interested_person"
  municipality_specific_amendments_allowed: true
  municipality_specific_condition: "Conditions must be alleged and established within the municipality that are not generally found within other municipalities."
  approval_required: true
  approving_authority_id: "ahj:usa-ct:state-building-inspector"
  adoption_procedure_source_ids:
    - "src:usa-ct:statute:chap-541"
  filing_required: true
  registry_exists: unresolved
  registry_source_ids: []
  legal_basis_source_ids:
    - "src:usa-ct:statute:chap-541"
    - "src:usa-ct:code:2022-csfpc-pdf"
  verification_status: "partially_verified"
  confidence: 0.87
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement and local amendment authority are separate in Connecticut. Municipal building officials administer the statewide State Building Code, but the code itself is adopted at the state level and is the code for all towns, cities, and boroughs. Towns, cities, boroughs, and interested persons may propose State Building Code amendments, including municipality-specific amendments when special local conditions are established, but the amendment must proceed under the state adoption process. For fire prevention, local fire marshals enforce the code locally and municipalities or fire districts may create certain permit and fee ordinances where the State Fire Prevention Code allows; that is narrower than a general local power to rewrite the state fire codes.

### 6.4 Known Local Amendment Registries

| Registry ID | Scope | Maintained By | Source ID | Machine Readable? | Parsed? | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| registry:usa-ct:state-building-code-variations | State Building Code variations / exemptions for existing buildings | Department of Administrative Services / State Building Inspector and Codes and Standards Committee | src:usa-ct:statute:chap-541; src:usa-ct:agency:modifications-waivers | partial | Statute requires a biennially updated online list of variations or exemptions for existing buildings; DAS also links to an index of State Building Code and State Fire Safety Code modifications. This is not the same as a comprehensive local-amendment registry. |
| registry:usa-ct:local-amendments | municipality-specific amendments | unresolved | none | no | no | No official statewide machine-readable local-amendment registry was verified. |

### 6.5 Municipality-Specific Known Amendments

| Jurisdiction | Code Family | Amendment Set ID | Approval Status | Effective Date | Source IDs | Parsed? |
| --- | --- | --- | --- | --- | --- | --- |
| unresolved | building | unresolved | unresolved | null | none | no |

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: state_plus_municipality_pending_local_contacts

Jurisdiction stack:

```text
Address
  -> State: Connecticut
  -> Municipality / town / city / borough
  -> Building AHJ: municipal building official unless state-building trigger applies
  -> Fire AHJ: local fire marshal for local initial determinations unless state-owned property or state-reserved issue applies
  -> State Building Inspector / Codes and Standards Committee for state-level interpretations, modifications, variations, exemptions, appeals, and code adoption
  -> State Fire Marshal for fire-code interpretations, state-property decisions, local-decision review, variations, exemptions, and code adoption
  -> Applicable statewide code adoption record
  -> Any verified local fire-prevention ordinance or state-approved municipality-specific building amendment
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | Connecticut General Assembly / CT.gov state portals | src:usa-ct:statute:chap-541; src:usa-ct:portal:ctgov-home | statewide | periodic | active |
| Municipality | State statutory municipal building official framework | src:usa-ct:statute:chap-541 | statewide | periodic | active_for_authority_model |
| Fire District | State Fire Prevention Code local ordinance provisions and local fire marshal statutes | src:usa-ct:statute:chap-541; src:usa-ct:code:2022-csfpc-pdf | partial | unknown | partially_verified |
| Special District | unresolved | none | unknown | unknown | unresolved |

### 7.3 AHJ Contact Data

No AHJ contact data was populated for this pass. The report identifies the statewide and local role model, but it does not contain municipal building department, local fire marshal, fire district, elevator inspection, or named personnel contact records.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Source Type | Publisher | URL | Accessed Date | Snapshot ID | Checksum | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-ct:portal:ctgov-home | CT.GOV - Connecticut's Official State Website | agency_page | State of Connecticut | https://portal.ct.gov/ | 2026-06-25 | snapshot:2026-06-25:ctgov-home | null | active |
| src:usa-ct:statute:chap-541 | Connecticut General Statutes, Chapter 541 | statute | Connecticut General Assembly | https://www.cga.ct.gov/current/pub/chap_541.htm | 2026-06-25 | snapshot:2026-06-25:chap-541 | null | active_with_caveat |
| src:usa-ct:agency:building-code-page | Connecticut State Building Code - Regulations | agency_page | Department of Administrative Services | https://portal.ct.gov/das/home/office-of-state-building-inspector/connecticut-state-building-code/regulations?language=en_US | 2026-06-25 | snapshot:2026-06-25:das-building-code-regulations | null | active |
| src:usa-ct:code:2022-csbc-pdf | 2022 Connecticut State Building Code with Errata #1 | code_pdf | Department of Administrative Services / Office of State Building Inspector | https://portal.ct.gov/das/-/media/das/office-of-state-building-inspector/2022-state-codes/2022-csbc-final.pdf | 2026-06-25 | snapshot:2026-06-25:2022-csbc-final | null | active |
| src:usa-ct:agency:fire-codes-overview | CT Fire Safety and Prevention Codes | agency_page | Department of Administrative Services / Office of State Fire Marshal | https://portal.ct.gov/en/das/office-of-state-fire-marshal/ct-fire-safety-and-prevention-codes?language=en_US | 2026-06-25 | snapshot:2026-06-25:das-fire-codes-overview | null | active |
| src:usa-ct:agency:fire-codes-docs | CT Fire Safety and Prevention Codes - Documents/Forms | agency_page | Department of Administrative Services / Office of State Fire Marshal | https://portal.ct.gov/das/office-of-state-fire-marshal/ct-fire-safety-and-prevention-codes/documents | 2026-06-25 | snapshot:2026-06-25:das-fire-codes-docs | null | active |
| src:usa-ct:code:2022-csfsc-pdf | 2022 Connecticut State Fire Safety Code | code_pdf | Department of Administrative Services / Office of State Fire Marshal | https://portal.ct.gov/-/media/das/office-of-state-building-inspector/2022-state-codes/2022-csfsc-final.pdf | 2026-06-25 | snapshot:2026-06-25:2022-csfsc-final | null | active |
| src:usa-ct:code:2022-csfpc-pdf | 2022 Connecticut State Fire Prevention Code | code_pdf | Department of Administrative Services / Office of State Fire Marshal | https://portal.ct.gov/-/media/DAS/Office-of-State-Building-Inspector/2022-State-Codes/2022-CSFPC-Final.pdf | 2026-06-25 | snapshot:2026-06-25:2022-csfpc-final | null | active |
| src:usa-ct:agency:code-adoption-process | Building and Fire Code Adoption Process | agency_page | Department of Administrative Services / Office of State Building Inspector | https://portal.ct.gov/das/office-of-state-building-inspector/building-and-fire-code-adoption-process | 2026-06-25 | snapshot:2026-06-25:das-code-adoption-process | null | active |
| src:usa-ct:agency:modifications-waivers | Building and Fire Code Modifications, Exemptions and Waivers | agency_page | Department of Administrative Services / Office of State Building Inspector and Office of State Fire Marshal | https://portal.ct.gov/das/office-of-state-building-inspector/building-and-fire-code-modifications-exemptions-and-waivers | 2026-06-25 | snapshot:2026-06-25:das-modifications-waivers | null | active |
| src:usa-ct:agency:elevator-code-page | Safety Code for Elevators and Escalators | agency_page | Department of Administrative Services / Office of State Building Inspector | https://portal.ct.gov/das/office-of-state-building-inspector/safety-code-for-elevators-and-escalators | 2026-06-25 | snapshot:2026-06-25:das-elevator-code | null | active |
| src:usa-ct:agency:boiler-elevator-inspections | Apply for boiler, elevator, and escalator inspections and certifications | agency_page | Department of Administrative Services / Office of State Building Inspector | https://portal.ct.gov/das/office-of-state-building-inspector/apply-for-boiler-elevator-and-escalator-inspections-and-certifications | 2026-06-25 | snapshot:2026-06-25:das-boiler-elevator-inspections | null | active |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-ct:statute:chap-541 | statutory_supplement_notice | CGA current chapter page includes a reader notice to consult the 2026 Supplement for statutes amended, repealed, or added during the 2025 legislative sessions. | Use for core authority structure; monitor supplement/current codification before final verification. |
| src:usa-ct:code:2022-csbc-pdf | official_pdf | DAS-posted PDF is an official code source but is not fully section-normalized in this report. | Use for current edition, model codes, and high-impact amendment facts; complete amendment parsing remains open. |
| src:usa-ct:code:2022-csfsc-pdf | official_pdf | DAS-posted fire safety code PDF is parsed only at high-value fields. | Use for current edition, Part III/Part IV split, and base model codes; complete amendment parsing remains open. |
| src:usa-ct:code:2022-csfpc-pdf | official_pdf | DAS-posted fire prevention code PDF is parsed only at high-value fields. | Use for current edition, NFPA 1 adoption, and AHJ split; complete amendment parsing remains open. |
| src:usa-ct:agency:elevator-code-page | official_html | Official DAS page verifies current ASME A17.1-2013 adoption, but full regulation text was not extracted. | Use for current elevator-code edition and effective date; parse RCSA text before final verification. |
| src:usa-ct:agency:code-adoption-process | official_html_dynamic | DAS reports pending 2026 code-cycle status; timing may change quickly. | Monitor frequently until 2026 code adoption is resolved. |

### 8.3 Supplemental Sources

| Source ID | Title | Source Type | Publisher | URL | Use | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| none | none | none | none | none | none | No non-official sources were used. |

### 8.4 Source Extraction Metadata

| Source ID | Parser | Parser Version | Extracted At | Extraction Confidence | Requires OCR | Requires Browser Automation | Manual Review Required |
| --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-ct:statute:chap-541 | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.95 | no | no | yes |
| src:usa-ct:agency:building-code-page | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.96 | no | no | no |
| src:usa-ct:code:2022-csbc-pdf | web_pdf | 2026-06-25 | 2026-06-25T00:00:00Z | 0.94 | no | no | yes |
| src:usa-ct:agency:fire-codes-overview | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.93 | no | no | no |
| src:usa-ct:agency:fire-codes-docs | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.94 | no | no | no |
| src:usa-ct:code:2022-csfsc-pdf | web_pdf | 2026-06-25 | 2026-06-25T00:00:00Z | 0.93 | no | no | yes |
| src:usa-ct:code:2022-csfpc-pdf | web_pdf | 2026-06-25 | 2026-06-25T00:00:00Z | 0.94 | no | no | yes |
| src:usa-ct:agency:code-adoption-process | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.92 | no | no | yes |
| src:usa-ct:agency:modifications-waivers | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.90 | no | no | no |
| src:usa-ct:agency:elevator-code-page | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.88 | no | no | yes |
| src:usa-ct:agency:boiler-elevator-inspections | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.88 | no | no | no |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| adoption:usa-ct:building:2022-csbc | state_code.name | 2022 Connecticut State Building Code | verified | 0.96 | src:usa-ct:agency:building-code-page; src:usa-ct:code:2022-csbc-pdf | Current code page and official PDF agree. |
| adoption:usa-ct:building:2022-csbc | base_model_code.code_name | 2021 IBC, 2021 IRC, 2021 IEBC, 2021 IPC, 2021 IMC, 2021 IECC, 2021 ISPSC, 2020 NFPA 70, 2017 ICC A117.1 | verified | 0.96 | src:usa-ct:agency:building-code-page; src:usa-ct:code:2022-csbc-pdf | Model-code list verified from DAS page and PDF. |
| adoption:usa-ct:building:2022-csbc | dates.effective_date | 2022-10-01 | verified | 0.96 | src:usa-ct:agency:building-code-page; src:usa-ct:code:2022-csbc-pdf | Applies to permit applications filed from 2022-10-01. |
| adoption:usa-ct:building:2022-csbc | applicability.date_trigger | permit_application_date | verified | 0.96 | src:usa-ct:agency:building-code-page; src:usa-ct:code:2022-csbc-pdf | Official DAS page states permit-application trigger. |
| adoption:usa-ct:building:2022-csbc | fuel_gas.ifgc_adopted | false | verified | 0.91 | src:usa-ct:code:2022-csbc-pdf | IFGC not adopted; NFPA fuel-gas references substituted. |
| adoption:usa-ct:fire-safety:2022-csfsc | state_code.name | 2022 Connecticut State Fire Safety Code | verified | 0.95 | src:usa-ct:agency:fire-codes-docs; src:usa-ct:code:2022-csfsc-pdf | Current code and official PDF verified. |
| adoption:usa-ct:fire-safety:2022-csfsc | base_model_code.code_name | 2021 IFC and 2021 NFPA 101 | verified | 0.95 | src:usa-ct:agency:fire-codes-docs; src:usa-ct:code:2022-csfsc-pdf | Part III/Part IV model-code split verified. |
| adoption:usa-ct:fire-safety:2022-csfsc | applicability.part_split_date | 2006-01-01 | verified | 0.95 | src:usa-ct:agency:fire-codes-docs; src:usa-ct:code:2022-csfsc-pdf | Part III / Part IV applicability verified. |
| adoption:usa-ct:fire-prevention:2022-csfpc | base_model_code.code_name | 2021 NFPA 1, Fire Code | verified | 0.96 | src:usa-ct:agency:fire-codes-overview; src:usa-ct:code:2022-csfpc-pdf | Official sources agree. |
| local-enforcement:usa-ct | building_official_model | municipal building official administers code | verified | 0.94 | src:usa-ct:statute:chap-541 | Statute verifies appointment and local administration. |
| local-enforcement:usa-ct | fire_marshal_model | local fire marshal initial determination with State Fire Marshal review | verified | 0.94 | src:usa-ct:statute:chap-541; src:usa-ct:code:2022-csfpc-pdf | Local/state split verified from statute and code text. |
| local-amendment-rule:usa-ct | building_amendment_proposal_path | towns, cities, boroughs, and interested persons may propose amendments, including municipality-specific amendments under special-local-condition showing | verified | 0.90 | src:usa-ct:statute:chap-541 | State adoption process still required. |
| registry:usa-ct:local-amendments | registry_exists | unresolved | unresolved | 0.00 | none | Comprehensive local-amendment registry not found in this pass. |
| adoption:usa-ct:elevator:asme-a17-1-2013 | base_model_code.code_name | ASME A17.1-2013 | partially_verified | 0.82 | src:usa-ct:agency:elevator-code-page | DAS page verified; full regulation text needs parsing. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Every `src:usa-ct:...` cited in the body appears in Section 8. |
| All authority IDs resolve | pass | Authority IDs used in adoption and enforcement sections are defined in Section 2. |
| Frontmatter placeholders replaced | pass | State, abbreviation, report ID, status, and dates are populated. |
| Core authority fields source-backed | pass | State Building Inspector, Codes and Standards Committee, State Fire Marshal, municipal building officials, and local fire marshals are tied to official sources. |
| Current code editions verified for core statewide codes | pass | Building, residential, existing, mechanical, plumbing, electrical, energy, accessibility, fire safety, and fire prevention rows are source-backed. |
| Fuel gas treatment explicit | pass | IFGC is explicitly marked not separately adopted; NFPA fuel-gas routing is cited. |
| Adoption/effective/operative/mandatory dates are separated | pass | Dates are separately recorded even when they fall on the same calendar day. |
| ISO date format used | pass | Dates use YYYY-MM-DD format where known. |
| No impossible date sequences introduced | pass | Unknown future dates and unresolved transitions remain null or unresolved. |
| Date rules have trigger conditions | pass | Permit-application, posting, and 2006 Part III/Part IV triggers are captured. |
| Local enforcement separated from local amendment authority | pass | Local administration is not treated as local code ownership. |
| 2026 pending changes captured | pass | Pending code cycle is recorded without asserting an effective date. |
| AHJ contacts populated | fail | Role model is populated; contact data remains open. |
| Complete state amendment parsing | fail | High-impact amendments are captured, but complete section-normalized amendments are open. |
| Complete local amendment registry | fail | No comprehensive official registry was verified. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-ct:001 | high | 2026 code cycle | DAS reports the 2026 State Building, Fire Safety, and Fire Prevention Codes are pending Legislative Regulation Review Committee approval and may be delayed. | Monitor DAS code-adoption page, LRC materials, and final posted codes; update effective and operative dates when adopted. | null | null | open |
| issue:usa-ct:002 | medium | local amendment registry | State Building Code amendment proposal path is verified, but no comprehensive registry of approved municipality-specific amendments was verified. | Search DAS code-making records, municipal filings, and any LRC/local amendment records. | null | null | open |
| issue:usa-ct:003 | medium | elevator code parsing | DAS verifies ASME A17.1-2013, but RCSA §§ 29-192-1e to 29-192-16e and any amendments were not parsed. | Extract current RCSA elevator regulation text and reconcile with DAS page. | null | null | open |
| issue:usa-ct:004 | low | AHJ contacts | State and local AHJ contact records were not populated. | Collect OSBI, OSFM, municipal building official, local fire marshal, and elevator inspection contact datasets. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-ct:building-code-page | src:usa-ct:agency:building-code-page | html_diff | monthly | Current State Building Code text or current-code PDF link changes | 2026-06-25 | active |
| watch:usa-ct:fire-codes-overview | src:usa-ct:agency:fire-codes-overview | html_diff | monthly | Fire Safety or Fire Prevention current-code links, errata, or adoption status changes | 2026-06-25 | active |
| watch:usa-ct:fire-codes-docs | src:usa-ct:agency:fire-codes-docs | html_diff | monthly | Fire Safety Code history, current code, errata, or effective-date changes | 2026-06-25 | active |
| watch:usa-ct:code-adoption-process | src:usa-ct:agency:code-adoption-process | html_diff | weekly_until_2026_cycle_resolves | Any posting of approval, final 2026 code documents, effective date, or delay update | 2026-06-25 | active |
| watch:usa-ct:statute-chap-541 | src:usa-ct:statute:chap-541 | statute_text_diff | quarterly | Chapter 541 authority, adoption, enforcement, variation, local amendment, or publication-rule changes | 2026-06-25 | active |
| watch:usa-ct:elevator-code | src:usa-ct:agency:elevator-code-page | html_diff | quarterly | Elevator code edition or effective-date change | 2026-06-25 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-25 | Upgraded Connecticut draft to partially verified report; populated current 2022 code editions, base model codes, effective dates, permit-application transition, fuel-gas treatment, fire safety/fire prevention records, 2006 fire-safety split, elevator code row, and pending 2026 code-cycle notes. | report:usa-ct; adoption:usa-ct:building:2022-csbc; adoption:usa-ct:fire-safety:2022-csfsc; adoption:usa-ct:fire-prevention:2022-csfpc; adoption:usa-ct:elevator:asme-a17-1-2013 | src:usa-ct:statute:chap-541; src:usa-ct:agency:building-code-page; src:usa-ct:code:2022-csbc-pdf; src:usa-ct:agency:fire-codes-overview; src:usa-ct:agency:fire-codes-docs; src:usa-ct:code:2022-csfsc-pdf; src:usa-ct:code:2022-csfpc-pdf; src:usa-ct:agency:code-adoption-process; src:usa-ct:agency:elevator-code-page; src:usa-ct:agency:boiler-elevator-inspections; src:usa-ct:agency:modifications-waivers | system | AHJ contacts, local amendment registry, full amendment parsing, and full elevator regulation parsing remain open. |
