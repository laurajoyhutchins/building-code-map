---
state:
  state_id: "US-HI"
  name: "Hawaii"
  abbreviation: "HI"
report:
  report_id: "state-report:usa-hi"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-25"
  last_verified: "2026-06-25"
  reviewed_by: null
risk:
  overall_confidence: 0.86 # 0.00 - 1.00
  risk_flags:
    - "county_adoption_and_state_adoption_are_both_relevant"
    - "mechanical_and_fuel_gas_not_separately_listed"
    - "accessibility_code_not_fully_parsed"
    - "official_source_discrepancy_electrical_edition"
  open_questions_count: 4

---

# State Building Code Authority Report: Hawaii

## 1. Executive Summary

- **Authority model:** Hawaii uses a statewide building-code council attached to the Department of Accounting and General Services, plus a State Fire Council for the fire code. Counties may still update county building codes, so local adoption remains part of the operating model.

- **Statewide code status:** Strong statewide code program with county overlay. The State Building Code Council publishes current state building codes, and the State Fire Code is adopted by the State Fire Council and becomes part of the Hawaii state building codes.

- **Local enforcement model:** Hybrid. State codes exist statewide, but counties may update county building codes and local AHJs administer fire and permit functions under the fire code.

- **Local amendment posture:** County-level adoption and update authority exists, and the 2025 emergency proclamation temporarily suspends some chapter 107 provisions to preserve housing-code stability while counties may still update county building codes.

- **Known transition periods or pending changes:** The Governor's September 23, 2025 emergency proclamation temporarily suspends several Chapter 107 provisions, including the default county adoption deadline for state building codes. A source conflict remains for the electrical-code edition: the SBCC code-list page labels the state electrical code as 2018 NFPA 70, while the linked official State Electrical Code PDF adopts the 2020 NFPA 70 effective March 15, 2022.

- **Production readiness:** usable_with_caution

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: State Building Code Council, authorized by HRS 107-22 and housed administratively
    in DAGS, establishes the Hawaii state building-code baseline through adoption
    of national model codes and state standards.
  confidence: 0.98
  source_ids:
  - src:usa-hi:bcc-home
  - src:usa-hi:bcc-building-code-rules
- topic: Primary building code edition
  finding: State Building Code adopts the 2018 IBC with Hawaii amendments; the SBCC
    adoption letter states adoption and effective date as 2021-04-20, with state-building
    design compliance by 2022-04-20 and county adoption/interim-code timing by 2023-04-20.
  confidence: 0.99
  source_ids:
  - src:usa-hi:bcc-building-code-rules
  - src:usa-hi:state-building-code-2018
- topic: Residential / existing building / energy / plumbing
  finding: SBCC current-code page lists 2018 IRC, 2018 IEBC, 2018 IECC, and 2018 UPC;
    linked official code PDFs support adoption/effective dates for residential, existing-building,
    energy, and plumbing code families.
  confidence: 0.94
  source_ids:
  - src:usa-hi:bcc-building-code-rules
  - src:usa-hi:state-residential-code-2018
  - src:usa-hi:state-existing-building-code-2018
  - src:usa-hi:state-energy-code-2018
  - src:usa-hi:state-plumbing-code-2018
- topic: Electrical code edition conflict
  finding: The SBCC current-code page labels the State Electric Code as 2018 NFPA
    70, but the linked official State Electrical Code PDF adopts the 2020 NFPA 70
    effective 2022-03-15. The report preserves the conflict rather than choosing one
    silently.
  confidence: 0.7
  source_ids:
  - src:usa-hi:bcc-building-code-rules
  - src:usa-hi:state-electrical-code-2020
- topic: Fire code authority
  finding: State Fire Council adopts the 2024 NFPA 1 Fire Code with Hawaii State Amendments;
    the fire-code PDF states an effective date of 2025-09-05 and administration/enforcement
    by the AHJ designated by HRS 132.
  confidence: 0.99
  source_ids:
  - src:usa-hi:bcc-building-code-rules
  - src:usa-hi:state-fire-code-2024
- topic: Local enforcement and amendment posture
  finding: 'Hawaii remains a state-plus-county model: county updates are expressly
    preserved in the 2025 affordable-housing proclamation and the State Fire Code
    uses county/AHJ permit and enforcement paths.'
  confidence: 0.94
  source_ids:
  - src:usa-hi:bcc-building-code-rules
  - src:usa-hi:governor-proclamation-2025-09-23
  - src:usa-hi:state-fire-code-2024
- topic: Production readiness limiter
  finding: Mechanical/fuel-gas standalone scope, accessibility codification, county
    amendment registry, and the electrical-edition conflict remain open issues.
  confidence: 0.88
  source_ids:
  - src:usa-hi:bcc-building-code-rules
  - src:usa-hi:state-electrical-code-2020
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-hi:state-building-code-council |
| Authority name | State Building Code Council |
| Authority type | state_agency |
| Legal basis | HRS 107-22 as referenced by the DAGS Building Code Council page |
| Role | Establishes the state building code through adoption of national model codes and state design standards |
| Enforcement model | hybrid |
| Source IDs | src:usa-hi:bcc-home; src:usa-hi:bcc-building-code-rules |
| Verification status | verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-hi:state-building-code-council | State Building Code Council | Adopts the State Building Code - 2018 IBC | HRS 107-22 / DAGS code page | src:usa-hi:bcc-building-code-rules | verified |
| Residential | ahj:usa-hi:state-building-code-council | State Building Code Council | Adopts the State Residential Code - 2018 IRC | DAGS code page | src:usa-hi:bcc-building-code-rules | verified |
| Existing Building / Rehabilitation | ahj:usa-hi:state-building-code-council | State Building Code Council | Adopts the State Existing Building Code - 2018 IEBC | DAGS code page | src:usa-hi:bcc-building-code-rules | verified |
| Mechanical | ahj:usa-hi:state-building-code-council | State Building Code Council | Not separately listed on the current DAGS page; likely handled through the state code/county code mix | DAGS code page | src:usa-hi:bcc-building-code-rules | unresolved |
| Plumbing | ahj:usa-hi:state-building-code-council | State Building Code Council | Adopts the State Plumbing Code - 2018 UPC | DAGS code page | src:usa-hi:bcc-building-code-rules | verified |
| Fuel Gas | ahj:usa-hi:state-building-code-council | State Building Code Council | Not separately listed on the current DAGS page; not parsed as a statewide standalone code here | DAGS code page | src:usa-hi:bcc-building-code-rules | unresolved |
| Electrical | ahj:usa-hi:state-building-code-council | State Building Code Council | DAGS list says State Electric Code - 2018 NFPA 70; linked official PDF adopts 2020 NFPA 70 effective 2022-03-15 | DAGS code page and linked State Electrical Code PDF | src:usa-hi:bcc-building-code-rules; src:usa-hi:state-electrical-code-2020 | source_conflict |
| Energy | ahj:usa-hi:state-building-code-council | State Building Code Council | Adopts the State Energy Code - 2018 IECC | DAGS code page | src:usa-hi:bcc-building-code-rules | verified |
| Fire - construction references | ahj:usa-hi:state-fire-council | State Fire Council | Adopts the State Fire Code and links it into the state building code framework | HRS 132-3 / DAGS fire-code page | src:usa-hi:bcc-building-code-rules; src:usa-hi:state-fire-code-2024 | verified |
| Fire - operational / prevention code | ahj:usa-hi:state-fire-council | State Fire Council | Adopts and amends the State Fire Code | HRS 132-3 / DAGS fire-code page | src:usa-hi:state-fire-code-2024 | verified |
| Accessibility | ahj:usa-hi:state-building-code-council | State Building Code Council | Not separately parsed here; state code page points at design standards and county update authority, but no accessibility edition was extracted | DAGS code page | src:usa-hi:bcc-building-code-rules | unresolved |
| Elevator / Conveyance | ahj:usa-hi:dlir-hiosh | Hawaii Department of Labor and Industrial Relations, Occupational Safety and Health | State Elevator Code and related systems are handled in DLIR OSH rules | DLIR HIOSH rules page | src:usa-hi:hiosh-elevators | verified |

### 2.3 Authority Hierarchy Notes

Hawaii is a true state-plus-county system. The State Building Code Council sets the state code baseline, the State Fire Council adopts the fire code, and counties retain authority to update county building codes. The September 23, 2025 affordable-housing proclamation temporarily narrowed some Chapter 107 mechanics, but it expressly preserved county ability to update county building codes.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-hi:001 | ahj:usa-hi:state-building-code-council | delegates_enforcement_to | county_building_code_admin | src:usa-hi:bcc-building-code-rules | verified |
| edge:usa-hi:002 | ahj:usa-hi:state-building-code-council | preempts | less_stringent_county_building_standards | src:usa-hi:bcc-building-code-rules | verified |
| edge:usa-hi:003 | ahj:usa-hi:state-fire-council | reserves_review_for | ahj_designated_by_hrs_132 | src:usa-hi:state-fire-code-2024 | verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | State Building Code | IBC | 2018 | current | 2021-04-20 | 2021-04-20 | 2022-04-20 | 2023-04-20 | State-building design compliance one year after adoption; county adoption/interim-county-code path two years after adoption, subject to later Chapter 107 emergency-proclamation caveats. | src:usa-hi:bcc-building-code-rules; src:usa-hi:state-building-code-2018; src:usa-hi:governor-proclamation-2025-09-23 |
| Residential | State Residential Code | IRC | 2018 | current | 2020-11-17 | 2020-11-17 | 2021-11-17 | 2022-11-17 | State-building design compliance one year after adoption; county adoption/interim-county-code path two years after adoption. | src:usa-hi:bcc-building-code-rules; src:usa-hi:state-residential-code-2018 |
| Existing Building / Rehabilitation | State Existing Building Code | IEBC | 2018 | current | 2020-11-17 | 2020-11-17 | 2021-11-17 | 2022-11-17 | State-building design compliance one year after adoption; county adoption/interim-county-code path two years after adoption. | src:usa-hi:bcc-building-code-rules; src:usa-hi:state-existing-building-code-2018 |
| Mechanical | county / local code framework | not verified | null | unresolved | null | null | null | null | Not separately listed on the current SBCC page as a standalone statewide code family. Mechanical content appears inside the IRC, but standalone statewide treatment was not verified. | src:usa-hi:bcc-building-code-rules; src:usa-hi:state-residential-code-2018 |
| Plumbing | State Plumbing Code | UPC | 2018 | current | null | 2020-05-19 | null | null | Current SBCC page links to an image-based plumbing-code PDF showing effective date 2020-05-19; adoption-date and county-deadline letter text were not machine-parsed. | src:usa-hi:bcc-building-code-rules; src:usa-hi:state-plumbing-code-2018 |
| Fuel Gas | county / local code framework | not verified | null | unresolved | null | null | null | null | Not separately listed on the current SBCC page as a standalone statewide code family. Fuel-gas content appears inside the IRC and plumbing-reference framework, but standalone statewide treatment was not verified. | src:usa-hi:bcc-building-code-rules; src:usa-hi:state-residential-code-2018 |
| Electrical | State Electrical Code | NFPA 70 | 2018 / 2020 conflict | source_conflict | 2022-03-15 | 2022-03-15 | 2023-03-14 | 2024-03-14 | SBCC current-code list labels 2018 NFPA 70, but linked official State Electrical Code PDF adopts 2020 NFPA 70; preserve conflict for manual review. | src:usa-hi:bcc-building-code-rules; src:usa-hi:state-electrical-code-2020 |
| Energy | State Energy Code | IECC | 2018 | current | 2020-12-15 | 2020-12-15 | 2021-12-15 | 2022-12-15 | State-building design compliance one year after adoption; county adoption/interim-county-energy-code path two years after adoption. | src:usa-hi:bcc-building-code-rules; src:usa-hi:state-energy-code-2018 |
| Fire - construction references | State Fire Code | NFPA 1 | 2024 | current | null | 2025-09-05 | null | null | Adopted by State Fire Council and made part of the Hawaii state building codes; building-code conflicts defer to building code for new construction except where fire code is specifically referenced. | src:usa-hi:state-fire-code-2024; src:usa-hi:bcc-building-code-rules |
| Fire - operational / prevention code | State Fire Code | NFPA 1 | 2024 | current | null | 2025-09-05 | null | null | State Fire Code effective date is 2025-09-05; administered and enforced by the AHJ designated by HRS 132. | src:usa-hi:state-fire-code-2024 |
| Accessibility | Hawaii design standards / accessibility framework | not verified | null | unresolved | null | null | null | null | IBC amendments include accessibility content, but a separately adopted accessibility-code edition was not verified. | src:usa-hi:bcc-home; src:usa-hi:state-building-code-2018 |
| Elevator / Conveyance | State Elevator Code | not verified | null | current | null | null | null | null | SBCC current-code page lists State Elevator Code and links to HIOSH Part 11 rules for elevators and related systems; base-model edition not verified. | src:usa-hi:bcc-building-code-rules; src:usa-hi:hiosh-elevators |

### 3.2 Adoption Records

```yaml
adoption_id: "adoption:usa-hi:building:state-building-code-2018"
state_id: "US-HI"
code_family: "building"
status: "current"
state_code:
  name: "State Building Code"
  edition_label: "2018 IBC with Hawaii amendments"
  codification: "HRS 107 / SBCC State Building Code PDF"
base_model_code:
  publisher: "ICC"
  code_name: "IBC"
  edition_year: 2018
  incorporated_by_reference: true
authority:
  adopting_authority_id: "ahj:usa-hi:state-building-code-council"
  enforcing_authority_model: "hybrid"
  interpretation_authority_id: "ahj:usa-hi:state-building-code-council"
dates:
  adoption_date: "2021-04-20"
  effective_date: "2021-04-20"
  operative_date: "2022-04-20"
  mandatory_date: "2023-04-20"
  replacement_date: null
applicability:
  date_trigger: "state_design_date_and_county_adoption_deadline"
  applies_to:
    - "new_construction"
    - "alteration"
    - "addition"
    - "state_owned"
  exclusions: []
  special_conditions:
    - "Counties may still update county building codes."
    - "Later emergency proclamation suspends some Chapter 107 mechanics."
transition:
  exists: true
  rule_id: "date-rule:usa-hi:building-2018"
  start_date: "2021-04-20"
  end_date: "2023-04-20"
  prior_code_allowed: true
  prior_code_condition: "State design compliance and county adoption deadlines are stated in the SBCC adoption letter; later Chapter 107 suspension should be checked for newer code-cycle effects."
amendments:
  state_amended: true
  amendment_set_ids:
    - "amendment-set:usa-hi:building:state-building-code"
  amendment_source_ids:
    - "src:usa-hi:state-building-code-2018"
provenance:
  source_ids:
    - "src:usa-hi:bcc-building-code-rules"
    - "src:usa-hi:state-building-code-2018"
    - "src:usa-hi:governor-proclamation-2025-09-23"
  field_sources:
    state_code.name: ["src:usa-hi:bcc-building-code-rules", "src:usa-hi:state-building-code-2018"]
    state_code.edition_label: ["src:usa-hi:state-building-code-2018"]
    dates.adoption_date: ["src:usa-hi:state-building-code-2018"]
    dates.effective_date: ["src:usa-hi:state-building-code-2018"]
    dates.operative_date: ["src:usa-hi:state-building-code-2018"]
    dates.mandatory_date: ["src:usa-hi:state-building-code-2018"]
verification:
  status: "verified"
  confidence: 0.99
  notes: "SBCC adoption letter and code PDF explicitly state edition, adoption/effective date, state design compliance date, and county adoption/interim-code timing."

adoption_id: "adoption:usa-hi:residential:state-residential-code-2018"
state_id: "US-HI"
code_family: "residential"
status: "current"
state_code:
  name: "State Residential Code"
  edition_label: "2018 IRC with Hawaii amendments"
  codification: "HRS 107 / SBCC State Residential Code PDF"
base_model_code:
  publisher: "ICC"
  code_name: "IRC"
  edition_year: 2018
  incorporated_by_reference: true
authority:
  adopting_authority_id: "ahj:usa-hi:state-building-code-council"
  enforcing_authority_model: "hybrid"
  interpretation_authority_id: "ahj:usa-hi:state-building-code-council"
dates:
  adoption_date: "2020-11-17"
  effective_date: "2020-11-17"
  operative_date: "2021-11-17"
  mandatory_date: "2022-11-17"
  replacement_date: null
applicability:
  date_trigger: "state_design_date_and_county_adoption_deadline"
  applies_to:
    - "one_and_two_family_dwellings"
    - "townhouses"
  exclusions: []
  special_conditions:
    - "Each county may require a permit by ordinance for areas regulated by the code."
transition:
  exists: true
  rule_id: "date-rule:usa-hi:residential-2018"
  start_date: "2020-11-17"
  end_date: "2022-11-17"
  prior_code_allowed: true
  prior_code_condition: "State design compliance and county adoption deadlines are stated in the SBCC adoption letter."
amendments:
  state_amended: true
  amendment_set_ids:
    - "amendment-set:usa-hi:residential:state-residential-code"
  amendment_source_ids:
    - "src:usa-hi:state-residential-code-2018"
provenance:
  source_ids:
    - "src:usa-hi:bcc-building-code-rules"
    - "src:usa-hi:state-residential-code-2018"
verification:
  status: "verified"
  confidence: 0.98
  notes: "SBCC residential-code PDF explicitly states edition, adoption/effective date, and transition dates."

adoption_id: "adoption:usa-hi:existing-building:state-existing-building-code-2018"
state_id: "US-HI"
code_family: "existing_building"
status: "current"
state_code:
  name: "State Existing Building Code"
  edition_label: "2018 IEBC with Hawaii amendments"
  codification: "HRS 107 / SBCC State Existing Building Code PDF"
base_model_code:
  publisher: "ICC"
  code_name: "IEBC"
  edition_year: 2018
  incorporated_by_reference: true
authority:
  adopting_authority_id: "ahj:usa-hi:state-building-code-council"
  enforcing_authority_model: "hybrid"
  interpretation_authority_id: "ahj:usa-hi:state-building-code-council"
dates:
  adoption_date: "2020-11-17"
  effective_date: "2020-11-17"
  operative_date: "2021-11-17"
  mandatory_date: "2022-11-17"
  replacement_date: null
applicability:
  date_trigger: "state_design_date_and_county_adoption_deadline"
  applies_to:
    - "alteration"
    - "movement"
    - "repair"
    - "change_of_use_or_occupancy"
    - "demolition"
  exclusions: []
  special_conditions: []
transition:
  exists: true
  rule_id: "date-rule:usa-hi:existing-building-2018"
  start_date: "2020-11-17"
  end_date: "2022-11-17"
  prior_code_allowed: true
  prior_code_condition: "State design compliance and county adoption deadlines are stated in the SBCC adoption letter."
amendments:
  state_amended: true
  amendment_set_ids:
    - "amendment-set:usa-hi:existing-building:state-existing-building-code"
  amendment_source_ids:
    - "src:usa-hi:state-existing-building-code-2018"
provenance:
  source_ids:
    - "src:usa-hi:bcc-building-code-rules"
    - "src:usa-hi:state-existing-building-code-2018"
verification:
  status: "verified"
  confidence: 0.98
  notes: "SBCC existing-building-code PDF explicitly states edition, adoption/effective date, and transition dates."

adoption_id: "adoption:usa-hi:plumbing:state-plumbing-code-2018"
state_id: "US-HI"
code_family: "plumbing"
status: "current"
state_code:
  name: "State Plumbing Code"
  edition_label: "2018 UPC with Hawaii amendments"
  codification: "HRS 107 / SBCC State Plumbing Code PDF"
base_model_code:
  publisher: "IAPMO"
  code_name: "UPC"
  edition_year: 2018
  incorporated_by_reference: true
authority:
  adopting_authority_id: "ahj:usa-hi:state-building-code-council"
  enforcing_authority_model: "hybrid"
  interpretation_authority_id: "ahj:usa-hi:state-building-code-council"
dates:
  adoption_date: null
  effective_date: "2020-05-19"
  operative_date: null
  mandatory_date: null
  replacement_date: null
applicability:
  date_trigger: "effective_date"
  applies_to:
    - "new_construction"
    - "relocation"
    - "alteration"
    - "repair"
    - "reconstruction"
  exclusions: []
  special_conditions:
    - "Each county may require a permit by ordinance for areas regulated by the code."
transition:
  exists: true
  rule_id: "date-rule:usa-hi:plumbing-2018"
  start_date: "2020-05-19"
  end_date: null
  prior_code_allowed: null
  prior_code_condition: "Plumbing-code source is image-based; adoption-date and county-deadline letter text were not extracted."
amendments:
  state_amended: true
  amendment_set_ids:
    - "amendment-set:usa-hi:plumbing:state-plumbing-code"
  amendment_source_ids:
    - "src:usa-hi:state-plumbing-code-2018"
provenance:
  source_ids:
    - "src:usa-hi:bcc-building-code-rules"
    - "src:usa-hi:state-plumbing-code-2018"
verification:
  status: "partially_verified"
  confidence: 0.90
  notes: "Current SBCC page lists 2018 UPC; linked plumbing PDF image shows effective date 2020-05-19 and 2018 UPC adoption language, but text extraction is image-dependent."

adoption_id: "adoption:usa-hi:electrical:state-electrical-code-2020"
state_id: "US-HI"
code_family: "electrical"
status: "source_conflict"
state_code:
  name: "State Electrical Code"
  edition_label: "2020 NFPA 70 per linked PDF; SBCC list labels 2018 NFPA 70"
  codification: "HRS 107 / SBCC State Electrical Code PDF"
base_model_code:
  publisher: "NFPA"
  code_name: "NFPA 70 / NEC"
  edition_year: 2020
  incorporated_by_reference: true
authority:
  adopting_authority_id: "ahj:usa-hi:state-building-code-council"
  enforcing_authority_model: "hybrid"
  interpretation_authority_id: "ahj:usa-hi:state-building-code-council"
dates:
  adoption_date: "2022-03-15"
  effective_date: "2022-03-15"
  operative_date: "2023-03-14"
  mandatory_date: "2024-03-14"
  replacement_date: null
applicability:
  date_trigger: "state_design_date_and_county_adoption_deadline"
  applies_to:
    - "electrical_installations_not_regulated_by_public_utilities_commission"
  exclusions: []
  special_conditions:
    - "Edition conflict between SBCC current-code list and linked official PDF requires manual review."
transition:
  exists: true
  rule_id: "date-rule:usa-hi:electrical-2020"
  start_date: "2022-03-15"
  end_date: "2024-03-14"
  prior_code_allowed: true
  prior_code_condition: "State design compliance and county adoption deadlines are stated in the SBCC adoption letter, but edition conflict remains unresolved."
amendments:
  state_amended: true
  amendment_set_ids:
    - "amendment-set:usa-hi:electrical:state-electrical-code"
  amendment_source_ids:
    - "src:usa-hi:state-electrical-code-2020"
provenance:
  source_ids:
    - "src:usa-hi:bcc-building-code-rules"
    - "src:usa-hi:state-electrical-code-2020"
verification:
  status: "source_conflict"
  confidence: 0.70
  notes: "Do not silently normalize to either 2018 or 2020. The linked official PDF supports 2020 NFPA 70, while the SBCC page text says 2018 NFPA 70."

adoption_id: "adoption:usa-hi:energy:state-energy-code-2018"
state_id: "US-HI"
code_family: "energy"
status: "current"
state_code:
  name: "State Energy Code"
  edition_label: "2018 IECC with Hawaii amendments"
  codification: "HRS 107 / SBCC State Energy Code PDF"
base_model_code:
  publisher: "ICC"
  code_name: "IECC"
  edition_year: 2018
  incorporated_by_reference: true
authority:
  adopting_authority_id: "ahj:usa-hi:state-building-code-council"
  enforcing_authority_model: "hybrid"
  interpretation_authority_id: "ahj:usa-hi:state-building-code-council"
dates:
  adoption_date: "2020-12-15"
  effective_date: "2020-12-15"
  operative_date: "2021-12-15"
  mandatory_date: "2022-12-15"
  replacement_date: null
applicability:
  date_trigger: "state_design_date_and_county_adoption_deadline"
  applies_to:
    - "new_construction"
    - "residential"
    - "commercial"
  exclusions: []
  special_conditions:
    - "Counties may still update county building codes."
transition:
  exists: true
  rule_id: "date-rule:usa-hi:energy-2018"
  start_date: "2020-12-15"
  end_date: "2022-12-15"
  prior_code_allowed: true
  prior_code_condition: "State design compliance and county adoption deadlines are stated in the SBCC adoption letter."
amendments:
  state_amended: true
  amendment_set_ids:
    - "amendment-set:usa-hi:energy:state-energy-code"
  amendment_source_ids:
    - "src:usa-hi:state-energy-code-2018"
provenance:
  source_ids:
    - "src:usa-hi:bcc-building-code-rules"
    - "src:usa-hi:state-energy-code-2018"
verification:
  status: "verified"
  confidence: 0.98
  notes: "SBCC energy-code PDF explicitly states edition, adoption/effective date, and transition dates."

adoption_id: "adoption:usa-hi:fire:state-fire-code-2024"
state_id: "US-HI"
code_family: "fire_operational"
status: "current"
state_code:
  name: "State Fire Code"
  edition_label: "2024 NFPA 1 with Hawaii State Amendments"
  codification: "HRS 132 / State Fire Council 2024 State Fire Code"
base_model_code:
  publisher: "NFPA"
  code_name: "NFPA 1 Fire Code"
  edition_year: 2024
  incorporated_by_reference: true
authority:
  adopting_authority_id: "ahj:usa-hi:state-fire-council"
  enforcing_authority_model: "state_and_county_ahj"
  interpretation_authority_id: "ahj:usa-hi:state-fire-council"
dates:
  adoption_date: null
  effective_date: "2025-09-05"
  operative_date: null
  mandatory_date: null
  replacement_date: null
applicability:
  date_trigger: "effective_date"
  applies_to:
    - "new_construction_when_referenced_from_building_code"
    - "alteration"
    - "existing_buildings"
    - "commercial"
    - "state_owned"
  exclusions: []
  special_conditions:
    - "Administered and enforced by the AHJ designated by HRS 132."
    - "For new building design/construction conflicts, the building code applies unless the fire code is specifically referenced from the building code."
transition:
  exists: true
  rule_id: "date-rule:usa-hi:fire-code-effective-2025-09-05"
  start_date: "2025-09-05"
  end_date: null
  prior_code_allowed: false
  prior_code_condition: "The 2024 NFPA 1 with Hawaii State Amendments is the current state fire code on the SBCC current-code page."
amendments:
  state_amended: true
  amendment_set_ids:
    - "amendment-set:usa-hi:fire:state-fire-code-2024"
  amendment_source_ids:
    - "src:usa-hi:state-fire-code-2024"
provenance:
  source_ids:
    - "src:usa-hi:state-fire-code-2024"
    - "src:usa-hi:bcc-building-code-rules"
verification:
  status: "verified"
  confidence: 0.99
  notes: "Fire-code text explicitly states the edition, effective date, and AHJ enforcement model."

adoption_id: "adoption:usa-hi:elevator:state-elevator-code"
state_id: "US-HI"
code_family: "elevator"
status: "current"
state_code:
  name: "State Elevator Code"
  edition_label: null
  codification: "SBCC current-code list / DLIR HIOSH Part 11"
base_model_code:
  publisher: null
  code_name: null
  edition_year: null
  incorporated_by_reference: null
authority:
  adopting_authority_id: "ahj:usa-hi:dlir-hiosh"
  enforcing_authority_model: "state"
  interpretation_authority_id: "ahj:usa-hi:dlir-hiosh"
dates:
  adoption_date: null
  effective_date: null
  operative_date: null
  mandatory_date: null
  replacement_date: null
applicability:
  date_trigger: null
  applies_to:
    - "elevators"
    - "escalators"
    - "dumbwaiters"
    - "moving_walks"
    - "material_lifts"
    - "personnel_hoists"
    - "platform_chairlifts"
  exclusions: []
  special_conditions:
    - "Base model and edition were not verified during this pass."
transition:
  exists: false
  rule_id: null
  start_date: null
  end_date: null
  prior_code_allowed: null
  prior_code_condition: null
amendments:
  state_amended: null
  amendment_set_ids:
    - "amendment-set:usa-hi:elevator:state-elevator-code"
  amendment_source_ids:
    - "src:usa-hi:hiosh-elevators"
provenance:
  source_ids:
    - "src:usa-hi:bcc-building-code-rules"
    - "src:usa-hi:hiosh-elevators"
verification:
  status: "partially_verified"
  confidence: 0.82
  notes: "SBCC page lists State Elevator Code and DLIR HIOSH Part 11 lists elevator-related administrative rules, but edition and base standard were not verified."
```

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

Hawaii has a state-code-council baseline for building codes and a separate state fire-code adoption path. The SBCC adoption letters separate adoption/effective dates from later state-design and county-adoption deadlines. The September 23, 2025 emergency proclamation temporarily suspended several Chapter 107 provisions, including the ability of the SBCC to amend or update state building codes in some circumstances and the default two-year county adoption deadline; the proclamation expressly says counties may still update county building codes. The fire-code PDF separately states that the 2024 NFPA 1 Fire Code with Hawaii State Amendments is effective 2025-09-05.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-hi:building-2018 | State Building Code | adoption_effective_operational_mandatory_sequence | 2021-04-20 / 2022-04-20 / 2023-04-20 | SBCC adoption letter sets adoption/effective date, state-design compliance deadline, and county adoption/interim-code deadline. | true | src:usa-hi:state-building-code-2018 | verified |
| date-rule:usa-hi:residential-2018 | State Residential Code | adoption_effective_operational_mandatory_sequence | 2020-11-17 / 2021-11-17 / 2022-11-17 | SBCC adoption letter sets adoption/effective date, state-design compliance deadline, and county adoption/interim-code deadline. | true | src:usa-hi:state-residential-code-2018 | verified |
| date-rule:usa-hi:existing-building-2018 | State Existing Building Code | adoption_effective_operational_mandatory_sequence | 2020-11-17 / 2021-11-17 / 2022-11-17 | SBCC adoption letter sets adoption/effective date, state-design compliance deadline, and county adoption/interim-code deadline. | true | src:usa-hi:state-existing-building-code-2018 | verified |
| date-rule:usa-hi:energy-2018 | State Energy Code | adoption_effective_operational_mandatory_sequence | 2020-12-15 / 2021-12-15 / 2022-12-15 | SBCC adoption letter sets adoption/effective date, state-design compliance deadline, and county adoption/interim-code deadline. | true | src:usa-hi:state-energy-code-2018 | verified |
| date-rule:usa-hi:plumbing-2018 | State Plumbing Code | effective_date | 2020-05-19 | Linked plumbing-code PDF image shows effective date. | unresolved | src:usa-hi:state-plumbing-code-2018 | partially_verified |
| date-rule:usa-hi:electrical-2020 | State Electrical Code | source_conflict_adoption_sequence | 2022-03-15 / 2023-03-14 / 2024-03-14 | Linked official PDF adopts 2020 NFPA 70, but SBCC current-code list labels 2018 NFPA 70. | true | src:usa-hi:bcc-building-code-rules; src:usa-hi:state-electrical-code-2020 | source_conflict |
| date-rule:usa-hi:county-update | county building codes | transition_period | ongoing | Counties may still update county building codes. | true | src:usa-hi:bcc-building-code-rules; src:usa-hi:governor-proclamation-2025-09-23 | verified |
| date-rule:usa-hi:emergency-proclamation | chapter 107 procedures | suspension_period | 2025-09-23 | Governor's emergency proclamation suspends several chapter 107 sections and adopts emergency rules. | true | src:usa-hi:bcc-building-code-rules; src:usa-hi:governor-proclamation-2025-09-23 | verified |
| date-rule:usa-hi:fire-code-effective-2025-09-05 | state fire code | effective_date | 2025-09-05 | NFPA 1 2024 with Hawaii State Amendments becomes the State Fire Code. | false | src:usa-hi:state-fire-code-2024 | verified |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| building | SBCC code-cycle update | null | null | null | null | null | monitoring | src:usa-hi:bcc-home; src:usa-hi:bcc-building-code-rules; src:usa-hi:governor-proclamation-2025-09-23 | Watch for proclamation changes and new county code actions. |
| fire_operational | fire code amendments | null | null | null | null | null | monitoring | src:usa-hi:state-fire-code-2024 | Watch State Fire Council updates. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-hi:001 | building | county building codes | county ordinance | Counties may still update county building codes. | src:usa-hi:bcc-building-code-rules; src:usa-hi:governor-proclamation-2025-09-23 | verified |
| applicability-rule:usa-hi:002 | fire_operational | all regulated fire areas | AHJ designation | The fire code is administered and enforced by the AHJ designated by HRS 132. | src:usa-hi:state-fire-code-2024 | verified |
| applicability-rule:usa-hi:003 | elevator | elevators and related systems | OSH standards | DLIR OSH rules list elevators, escalators, dumbwaiters, moving walks, and related systems. | src:usa-hi:hiosh-elevators | verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** state council code adoption with county update authority and a separate state fire-code adoption path

**Where amendments are published:** Building Code Council pages and supporting PDFs on ags.hawaii.gov; fire code amendments on law.hawaii.gov

**Amendment parsing status:** partial

### 5.2 State Amendment Sources

| Amendment Set ID | Code Family | Publication Format | Source IDs | Parsed? | Notes |
| --- | --- | --- | --- | --- | --- |
| amendment-set:usa-hi:building:state-building-code | building | SBCC code PDF | src:usa-hi:state-building-code-2018 | partial | State building code PDF is parsed for core adoption/date fields, not all amendments. |
| amendment-set:usa-hi:residential:state-residential-code | residential | SBCC code PDF | src:usa-hi:state-residential-code-2018 | partial | State residential code PDF is parsed for core adoption/date fields, not all amendments. |
| amendment-set:usa-hi:existing-building:state-existing-building-code | existing_building | SBCC code PDF | src:usa-hi:state-existing-building-code-2018 | partial | State existing building code PDF is parsed for core adoption/date fields, not all amendments. |
| amendment-set:usa-hi:plumbing:state-plumbing-code | plumbing | SBCC image-based PDF | src:usa-hi:state-plumbing-code-2018 | partial | State plumbing code source is image-based and was parsed only from visual review. |
| amendment-set:usa-hi:electrical:state-electrical-code | electrical | SBCC code PDF | src:usa-hi:state-electrical-code-2020 | partial | Core adoption/date fields parsed; edition conflicts with SBCC current-code list. |
| amendment-set:usa-hi:energy:state-energy-code | energy | SBCC code PDF | src:usa-hi:state-energy-code-2018 | partial | State energy code PDF is parsed for core adoption/date fields, not all amendments. |
| amendment-set:usa-hi:fire:state-fire-code-2024 | fire_operational | state fire code PDF | src:usa-hi:state-fire-code-2024 | partial | Fire code text states adoption, effective date, AHJ enforcement, and amendments. |
| amendment-set:usa-hi:elevator:state-elevator-code | elevator | DLIR standards page | src:usa-hi:hiosh-elevators | partial | Elevator-related rules are listed in DLIR HIOSH Part 11; base standard and edition were not verified. |

### 5.3 High-Impact State Amendments

| Code Family | Section | Amendment Type | Summary | Source IDs | Confidence |
| --- | --- | --- | --- | --- | --- |
| building | 107-24 / 107-28 / 107-31 | suspend | The September 23, 2025 emergency proclamation temporarily suspends several Chapter 107 provisions affecting SBCC updates and county adoption deadlines. | src:usa-hi:bcc-building-code-rules; src:usa-hi:governor-proclamation-2025-09-23 | 0.97 |
| fire_operational | chapter 132 / rule 4 | adopt | The 2024 NFPA 1 Fire Code with Hawaii State Amendments is adopted as the State Fire Code and is effective September 5, 2025. | src:usa-hi:state-fire-code-2024 | 0.99 |
| fire_operational | 107-25 / 132-3 | integrate | Once adopted, the State Fire Code becomes part of the Hawaii State Building Codes. | src:usa-hi:bcc-building-code-rules; src:usa-hi:state-fire-code-2024 | 0.98 |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-hi"
  model: "hybrid"
  enforcing_entities:
    - "county_government"
    - "authority_having_jurisdiction"
    - "state_fire_code_ahj"
  required_officials:
    - "building_official"
    - "county_code_official"
    - "ahj"
  state_reserved_activities:
    - "state_building_code_council_adoption"
    - "state_fire_code_adoption"
    - "county_update_deadline"
  source_ids:
    - "src:usa-hi:bcc-building-code-rules"
    - "src:usa-hi:state-fire-code-2024"
    - "src:usa-hi:governor-proclamation-2025-09-23"
  verification_status: "verified"
  confidence: 0.95
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-hi"
  model: "home_rule_local_choice"
  applies_to_code_families:
    - "building"
    - "residential"
    - "existing_building"
    - "plumbing"
    - "electrical"
    - "energy"
    - "fire_operational"
  approval_required: false
  approving_authority_id: null
  filing_required: "unresolved"
  registry_exists: false
  registry_source_ids:
    - "src:usa-hi:bcc-building-code-rules"
  legal_basis_source_ids:
    - "src:usa-hi:bcc-building-code-rules"
    - "src:usa-hi:governor-proclamation-2025-09-23"
  verification_status: "partially_verified"
  confidence: 0.82
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Hawaii is not purely state-enforced or purely local. The state code council sets the state baseline, counties can still update county building codes, and the fire code is enforced by the AHJ designated under Chapter 132. The September 23, 2025 emergency proclamation narrowed some Chapter 107 mechanics but expressly preserved county code updates. Filing or registry requirements for county amendments were not fully verified.

### 6.4 Known Local Amendment Registries

| Registry ID | Scope | Maintained By | Source ID | Machine Readable? | Parsed? | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| registry:usa-hi:local-amendments | county | not verified | src:usa-hi:bcc-building-code-rules | no | no | The page provides a public proposal form but no single machine-readable statewide registry was parsed. |

### 6.5 Municipality-Specific Known Amendments

| Jurisdiction | Code Family | Amendment Set ID | Approval Status | Effective Date | Source IDs | Parsed? |
| --- | --- | --- | --- | --- | --- | --- |
| unknown | building | local-amendment:usa-hi:unknown | unknown | null | none | no |

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: county_level

Jurisdiction stack:

```text
Address
  -> State
  -> County
  -> Municipality / unincorporated county
  -> Special districts, if applicable
  -> Building AHJ
  -> Fire AHJ
  -> Trade-specific AHJs
  -> Applicable state code adoption records
  -> Applicable local amendment records
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- | --- |
| State | Hawaii Department of Accounting and General Services | src:usa-hi:bcc-home | statewide | periodic | active |
| County | county building code systems | src:usa-hi:bcc-building-code-rules | statewide | periodic | active |
| Municipality | county / city charter structures | src:usa-hi:bcc-building-code-rules | statewide | periodic | active |
| Fire District | AHJ designated by HRS 132 | src:usa-hi:state-fire-code-2024 | statewide | periodic | active |
| Special District | not verified | none | unknown | unknown | pending |

### 7.3 AHJ Contact Data

| AHJ ID | Jurisdiction | Department | Role | Website | Phone | Email | Last Verified | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| ahj:hi:state-building-code-council | Statewide | State Building Code Council | building | https://ags.hawaii.gov/bcc/ | 808-586-0400 | [email protected] | 2026-06-25 | src:usa-hi:bcc-home |
| ahj:hi:state-fire-code | Statewide | State Fire Code AHJ | fire | https://law.hawaii.gov/wp-content/uploads/2025/09/NFPA1-2024-Hawaii-State-Fire-Code-amendments.pdf | null | null | 2026-06-25 | src:usa-hi:state-fire-code-2024 |
| ahj:hi:dlir-osh-elevators | Statewide | DLIR Occupational Safety and Health | elevators | https://labor.hawaii.gov/hiosh/standards/standards-adminrules-part-11/ | 808-586-9116 | [email protected] | 2026-06-25 | src:usa-hi:hiosh-elevators |

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Source Type | Publisher | URL | Accessed Date | Snapshot ID | Checksum | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-hi:bcc-home | Building Code Council | agency_page | Department of Accounting and General Services | https://ags.hawaii.gov/bcc/ | 2026-06-25 | snapshot:hi-bcc-home-2026-06-25 | null | active |
| src:usa-hi:bcc-building-code-rules | Building Code Rules | agency_page | Department of Accounting and General Services | https://ags.hawaii.gov/bcc/building-code-rules/ | 2026-06-25 | snapshot:hi-bcc-rules-2026-06-25 | null | active |
| src:usa-hi:state-building-code-2018 | State Building Code - 2018 IBC with Hawaii Amendments | code_pdf | State Building Code Council / Department of Accounting and General Services | https://ags.hawaii.gov/wp-content/uploads/2021/09/2018StateBuildingCode_20210817.pdf | 2026-06-25 | snapshot:hi-building-code-2018-2026-06-25 | null | active |
| src:usa-hi:state-residential-code-2018 | State Residential Code - 2018 IRC with Hawaii Amendments | code_pdf | State Building Code Council / Department of Accounting and General Services | https://ags.hawaii.gov/wp-content/uploads/2021/09/2018StateResidentialCode_20201117rev1.pdf | 2026-06-25 | snapshot:hi-residential-code-2018-2026-06-25 | null | active |
| src:usa-hi:state-existing-building-code-2018 | State Existing Building Code - 2018 IEBC with Hawaii Amendments | code_pdf | State Building Code Council / Department of Accounting and General Services | https://ags.hawaii.gov/wp-content/uploads/2021/09/2018StateExistingBuildingCode_20201117.pdf | 2026-06-25 | snapshot:hi-existing-building-code-2018-2026-06-25 | null | active |
| src:usa-hi:state-energy-code-2018 | State Energy Code - 2018 IECC with Hawaii Amendments | code_pdf | State Building Code Council / Department of Accounting and General Services | https://ags.hawaii.gov/wp-content/uploads/2021/01/soh_bcc_energycode_20201215.pdf | 2026-06-25 | snapshot:hi-energy-code-2018-2026-06-25 | null | active |
| src:usa-hi:state-plumbing-code-2018 | State Plumbing Code - 2018 UPC with Hawaii Amendments | code_pdf | State Building Code Council / Department of Accounting and General Services | https://ags.hawaii.gov/wp-content/uploads/2020/06/StatePlumbingCode_20200519.pdf | 2026-06-25 | snapshot:hi-plumbing-code-2018-2026-06-25 | null | active |
| src:usa-hi:state-electrical-code-2020 | State Electrical Code - 2020 NFPA 70 / NEC | code_pdf | State Building Code Council / Department of Accounting and General Services | https://ags.hawaii.gov/wp-content/uploads/2021/12/StateElectricalCode_03152022.pdf | 2026-06-25 | snapshot:hi-electrical-code-2020-2026-06-25 | null | active |
| src:usa-hi:state-fire-code-2024 | 2024 State Fire Code Amendments - Draft | regulation_pdf | State Fire Council / Department of Law Enforcement | https://law.hawaii.gov/wp-content/uploads/2025/09/NFPA1-2024-Hawaii-State-Fire-Code-amendments.pdf | 2026-06-25 | snapshot:hi-fire-code-2024-2026-06-25 | null | active |
| src:usa-hi:governor-proclamation-2025-09-23 | Fifteenth Proclamation Relating to Affordable Housing | proclamation_pdf | Office of the Governor, State of Hawaii | https://governor.hawaii.gov/wp-content/uploads/2025/09/2509073.pdf | 2026-06-25 | snapshot:hi-affordable-housing-proclamation-2025-09-23 | null | active |
| src:usa-hi:hiosh-elevators | Standards (AdminRules) Part 11 - Elevators and Related Systems | agency_page | Hawaii Department of Labor and Industrial Relations | https://labor.hawaii.gov/hiosh/standards/standards-adminrules-part-11/ | 2026-06-25 | snapshot:hi-hiosh-part11-2026-06-25 | null | active |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-hi:bcc-building-code-rules | official_html | The current page clearly lists current codes and Chapter 107 proclamation effects, but mechanical, fuel gas, and accessibility were not separately parsed as standalone state codes. The electrical-code list item conflicts with the linked official State Electrical Code PDF. | use for current code-family registry and county-update rules; do not infer missing families; preserve electrical conflict |
| src:usa-hi:state-plumbing-code-2018 | image_pdf | The plumbing PDF has poor text extraction; core fields were confirmed visually from rendered pages. | use only for visible effective date and 2018 UPC adoption language; leave unparsed dates unresolved |
| src:usa-hi:state-electrical-code-2020 | official_pdf_conflict | Linked official PDF adopts 2020 NFPA 70 even though the SBCC current-code page labels the State Electric Code as 2018 NFPA 70. | treat electrical edition as source_conflict pending manual agency confirmation |
| src:usa-hi:state-fire-code-2024 | draft_pdf | The PDF is labeled draft in link metadata but contains the current fire-code text, effective date, adoption language, county permit authorization, and AHJ enforcement structure. | use for fire-code core fields and note the draft label |
| src:usa-hi:governor-proclamation-2025-09-23 | pdf_orientation | Some pages include rotated signature/rules text; core Chapter 107 suspension text was readable in parsed text and screenshots. | use for Chapter 107 suspension and county-update preservation only |

### 8.3 Supplemental Sources

| Source ID | Title | Source Type | Publisher | URL | Use | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| none | none | none | none | none | none | none |

### 8.4 Source Extraction Metadata

| Source ID | Parser | Parser Version | Extracted At | Extraction Confidence | Requires OCR | Requires Browser Automation | Manual Review Required |
| --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-hi:bcc-home | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.97 | no | no | no |
| src:usa-hi:bcc-building-code-rules | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.96 | no | no | yes |
| src:usa-hi:state-building-code-2018 | pdf_text | 2026-06-25 | 2026-06-25T00:00:00Z | 0.99 | no | no | no |
| src:usa-hi:state-residential-code-2018 | pdf_text | 2026-06-25 | 2026-06-25T00:00:00Z | 0.99 | no | no | no |
| src:usa-hi:state-existing-building-code-2018 | pdf_text | 2026-06-25 | 2026-06-25T00:00:00Z | 0.99 | no | no | no |
| src:usa-hi:state-energy-code-2018 | pdf_text | 2026-06-25 | 2026-06-25T00:00:00Z | 0.99 | no | no | no |
| src:usa-hi:state-plumbing-code-2018 | pdf_screenshot | 2026-06-25 | 2026-06-25T00:00:00Z | 0.82 | yes | no | yes |
| src:usa-hi:state-electrical-code-2020 | pdf_text | 2026-06-25 | 2026-06-25T00:00:00Z | 0.97 | no | no | yes |
| src:usa-hi:state-fire-code-2024 | pdf_text_and_screenshot | 2026-06-25 | 2026-06-25T00:00:00Z | 0.99 | no | no | no |
| src:usa-hi:governor-proclamation-2025-09-23 | pdf_text_and_screenshot | 2026-06-25 | 2026-06-25T00:00:00Z | 0.95 | no | no | no |
| src:usa-hi:hiosh-elevators | web | 2026-06-25 | 2026-06-25T00:00:00Z | 0.95 | no | no | no |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| adoption:usa-hi:building:state-building-code-2018 | dates.adoption_date | 2021-04-20 | verified | 0.99 | src:usa-hi:state-building-code-2018 | SBCC adoption letter states adoption date. |
| adoption:usa-hi:building:state-building-code-2018 | dates.operative_date | 2022-04-20 | verified | 0.99 | src:usa-hi:state-building-code-2018 | State-building design compliance date. |
| adoption:usa-hi:building:state-building-code-2018 | dates.mandatory_date | 2023-04-20 | verified | 0.99 | src:usa-hi:state-building-code-2018 | County adoption/interim-code date stated in adoption letter. |
| adoption:usa-hi:residential:state-residential-code-2018 | state_code.edition_label | 2018 IRC with Hawaii amendments | verified | 0.98 | src:usa-hi:state-residential-code-2018 | Current-code list and PDF agree on 2018 IRC. |
| adoption:usa-hi:existing-building:state-existing-building-code-2018 | state_code.edition_label | 2018 IEBC with Hawaii amendments | verified | 0.98 | src:usa-hi:state-existing-building-code-2018 | Current-code list and PDF agree on 2018 IEBC. |
| adoption:usa-hi:plumbing:state-plumbing-code-2018 | dates.effective_date | 2020-05-19 | partially_verified | 0.90 | src:usa-hi:state-plumbing-code-2018 | Confirmed visually from image-based PDF. |
| adoption:usa-hi:electrical:state-electrical-code-2020 | state_code.edition_label | 2020 NFPA 70 per linked PDF; SBCC list labels 2018 NFPA 70 | source_conflict | 0.70 | src:usa-hi:bcc-building-code-rules; src:usa-hi:state-electrical-code-2020 | Official sources conflict. |
| adoption:usa-hi:energy:state-energy-code-2018 | dates.adoption_date | 2020-12-15 | verified | 0.98 | src:usa-hi:state-energy-code-2018 | SBCC adoption letter states adoption date. |
| adoption:usa-hi:fire:state-fire-code-2024 | dates.effective_date | 2025-09-05 | verified | 0.99 | src:usa-hi:state-fire-code-2024 | Fire-code PDF states effective date. |
| adoption:usa-hi:fire:state-fire-code-2024 | authority.adopting_authority_id | ahj:usa-hi:state-fire-council | verified | 0.99 | src:usa-hi:state-fire-code-2024 | State Fire Council adoption explicit in PDF. |
| local-enforcement:usa-hi | model | hybrid | verified | 0.95 | src:usa-hi:bcc-building-code-rules; src:usa-hi:state-fire-code-2024; src:usa-hi:governor-proclamation-2025-09-23 | County updates and AHJ enforcement both matter. |
| local-amendment-rule:usa-hi | filing_required | unresolved | unresolved | 0.50 | src:usa-hi:bcc-building-code-rules; src:usa-hi:governor-proclamation-2025-09-23 | County update authority is supported; filing/registry mechanics are not. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | All `src:usa-hi:*` IDs cited in the body have source-registry rows. |
| All authority IDs resolve | pass | Authority IDs are internally consistent. |
| Current verified code families have adoption records | pass | Records now exist for building, residential, existing building, plumbing, electrical, energy, fire, and elevator; unresolved mechanical/fuel-gas/accessibility rows remain explicit. |
| Building and operational fire code are separated | pass | SBCC building code and State Fire Code are separate authorities and records. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Adoption/effective dates, state-design operative dates, county adoption/interim dates, and fire effective date are separated. |
| Effective dates are valid ISO dates | pass | Dates are ISO formatted where known. |
| No impossible date sequences | pass | Known adoption/effective/operative/mandatory sequences are chronological. |
| Transition rules have explicit trigger conditions | pass | SBCC adoption letters, county update logic, proclamation suspension, and fire-code effective date each have rule rows. |
| Permit-date logic is captured where applicable | pass | County permit authorization and AHJ review paths are captured for residential, energy, plumbing, electrical, and fire where sources support it. |
| Local enforcement model classified | pass | Classified as hybrid. |
| Local amendment rule classified | pass | Classified as partially verified; county update authority is supported, registry/filing details remain unresolved. |
| AHJ confirmation metadata present | pass | AHJ contact data is present and dated 2026-06-25. |
| Official-source caveats captured | pass | Caveats note the electrical edition conflict, image-based plumbing PDF, fire-code draft label, and unparsed mechanical/fuel-gas/accessibility gaps. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-hi:001 | medium | mechanical and fuel-gas scope | Confirm whether Hawaii treats mechanical and fuel-gas as separate statewide code families or only as content inside adopted residential/plumbing/local code frameworks. | Extract statutory/rule references or agency guidance addressing standalone mechanical and fuel-gas scope. | null | null | open |
| issue:usa-hi:002 | medium | accessibility code | Confirm the current accessibility code edition and codification. | Extract adopted accessibility text, design standard, or rule beyond the IBC Chapter 11 amendment references. | null | null | open |
| issue:usa-hi:003 | low | county code registry | Determine whether any county or state office maintains a machine-readable code-update/amendment registry. | Review county code pages, ordinances, and DAGS/SBCC proposal or filing workflows. | null | null | open |
| issue:usa-hi:004 | high | electrical edition conflict | SBCC current-code page labels the State Electric Code as 2018 NFPA 70, but the linked official State Electrical Code PDF adopts 2020 NFPA 70 effective 2022-03-15. | Obtain agency confirmation or corrected SBCC page/PDF before marking electrical edition verified. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-hi:bcc-rules | src:usa-hi:bcc-building-code-rules | html_diff | monthly | state code list, electrical-code label, or chapter 107 changes | 2026-06-25 | active |
| watch:usa-hi:bcc-home | src:usa-hi:bcc-home | html_diff | monthly | council mission or code-adoption timeline changes | 2026-06-25 | active |
| watch:usa-hi:fire-code | src:usa-hi:state-fire-code-2024 | pdf_hash_diff | monthly | fire-code PDF hash changes | 2026-06-25 | active |
| watch:usa-hi:electrical-code | src:usa-hi:state-electrical-code-2020 | pdf_hash_diff | monthly | State Electrical Code PDF hash or edition changes | 2026-06-25 | active |
| watch:usa-hi:proclamation | src:usa-hi:governor-proclamation-2025-09-23 | pdf_hash_diff | monthly | proclamation extension, expiration, or chapter 107 suspension changes | 2026-06-25 | active |
| watch:usa-hi:elevators | src:usa-hi:hiosh-elevators | html_diff | quarterly | elevator rules changes | 2026-06-25 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-25 | Upgraded Hawaii report to partially_verified after official-source pass; added code-specific adoption records and dates; preserved electrical edition conflict. | report:usa-hi; adoption:usa-hi:* | src:usa-hi:bcc-home; src:usa-hi:bcc-building-code-rules; src:usa-hi:state-building-code-2018; src:usa-hi:state-residential-code-2018; src:usa-hi:state-existing-building-code-2018; src:usa-hi:state-energy-code-2018; src:usa-hi:state-plumbing-code-2018; src:usa-hi:state-electrical-code-2020; src:usa-hi:state-fire-code-2024; src:usa-hi:governor-proclamation-2025-09-23; src:usa-hi:hiosh-elevators | system | Hawaii is a hybrid state-plus-county code system. Electrical edition remains open because official page text conflicts with the linked official PDF. |