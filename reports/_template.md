---
state:
  state_id: "US-{{XX}}"
  name: "{{State Name}}"
  abbreviation: "{{XX}}"
report:
  report_id: "state-report:usa-{{xx}}"
  schema_version: "0.2.1"
  status: "draft" # draft | partially_verified | verified | deprecated
  last_updated: "{{YYYY-MM-DD}}"
  last_verified: "{{YYYY-MM-DD | null}}"
  reviewed_by: null
risk:
  overall_confidence: null # 0.00 - 1.00
  risk_flags: []
  open_questions_count: 0
---

# Building Code Authority Report: {{State Name}}

## 1. Executive Summary

- **Authority model:** {{short description}}
- **Statewide code status:** {{short description}}
- **Local enforcement model:** {{short description}}
- **Local amendment posture:** {{short description}}
- **Known transition periods or pending changes:** {{short description}}
- **Production readiness:** {{not_ready | partial | usable_with_caution | production_ready}}

### Key Findings

| Topic | Finding | Confidence | Source IDs |
| --- | --- | --- | --- |
| **State adopting authority** | {{finding}} | {{0.00}} | src:... |
| **Primary building code edition** | {{finding}} | {{0.00}} | src:... |
| **Electrical code authority** | {{finding}} | {{0.00}} | src:... |
| **Fire code authority** | {{finding}} | {{0.00}} | src:... |
| **Local enforcement** | {{finding}} | {{0.00}} | src:... |
| **Local amendments** | {{finding}} | {{0.00}} | src:... |
| **Effective / operative date rule** | {{finding}} | {{0.00}} | src:... |

---

## 2. Authority Structure

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| **Authority ID** | ahj:usa-{{xx}}:{{authority-slug}} |
| **Authority name** | {{name}} |
| **Authority type** | {{state_agency}} |
| **Legal basis** | {{statute/regulation citation}} |
| **Role** | {{adopts}} |
| **Enforcement model** | {{state_enforced}} |
| **Source IDs** | src:... |
| **Verification status** | {{verified}} |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| **Building** | auth:... | {{name}} | {{role}} | {{citation}} | src:... | {{status}} |
| **Residential** | auth:... | {{name}} | {{role}} | {{citation}} | src:... | {{status}} |
| **Existing Building / Rehabilitation** | auth:... | {{name}} | {{role}} | {{citation}} | src:... | {{status}} |
| **Mechanical** | auth:... | {{name}} | {{role}} | {{citation}} | src:... | {{status}} |
| **Plumbing** | auth:... | {{name}} | {{role}} | {{citation}} | src:... | {{status}} |
| **Fuel Gas** | auth:... | {{name}} | {{role}} | {{citation}} | src:... | {{status}} |
| **Electrical** | auth:... | {{name}} | {{role}} | {{citation}} | src:... | {{status}} |
| **Energy** | auth:... | {{name}} | {{role}} | {{citation}} | src:... | {{status}} |
| **Fire - construction references** | auth:... | {{name}} | {{role}} | {{citation}} | src:... | {{status}} |
| **Fire - operational / prevention code** | auth:... | {{name}} | {{role}} | {{citation}} | src:... | {{status}} |
| **Accessibility** | auth:... | {{name}} | {{role}} | {{citation}} | src:... | {{status}} |
| **Elevator / Conveyance** | auth:... | {{name}} | {{role}} | {{citation}} | src:... | {{status}} |

### 2.3 Authority Hierarchy Notes

{{Plain-English explanation of how state, county, municipal, fire, trade-specific, and specialized authorities interact.}}

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-{{xx}}:001 | auth:... | delegates_enforcement_to | municipal_ahj | src:... | {{status}} |
| edge:usa-{{xx}}:002 | auth:... | preempts | local_construction_standards | src:... | {{status}} |
| edge:usa-{{xx}}:003 | auth:... | reserves_review_for | {{occupancy/project type}} | src:... | {{status}} |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| **Building** | {{state code name}} | {{IBC/etc.}} | {{year}} | current | {{date/null}} | {{date/null}} | {{date/null}} | {{date/null}} | {{summary}} | src:... |
| **Residential** | {{state code name}} | {{IRC/etc.}} | {{year}} | current | {{date/null}} | {{date/null}} | {{date/null}} | {{date/null}} | {{summary}} | src:... |
| **Existing Building / Rehabilitation** | {{state code name}} | {{IEBC/state rehab code/etc.}} | {{year/null}} | current | {{date/null}} | {{date/null}} | {{date/null}} | {{date/null}} | {{summary}} | src:... |
| **Mechanical** | {{state code name}} | {{IMC/etc.}} | {{year}} | current | {{date/null}} | {{date/null}} | {{date/null}} | {{date/null}} | {{summary}} | src:... |
| **Plumbing** | {{state code name}} | {{IPC/UPC/NSPC/etc.}} | {{year}} | current | {{date/null}} | {{date/null}} | {{date/null}} | {{date/null}} | {{summary}} | src:... |
| **Fuel Gas** | {{state code name}} | {{IFGC/NFPA/etc.}} | {{year}} | current | {{date/null}} | {{date/null}} | {{date/null}} | {{date/null}} | {{summary}} | src:... |
| **Electrical** | {{state code name}} | {{NEC/NFPA 70}} | {{year}} | current | {{date/null}} | {{date/null}} | {{date/null}} | {{date/null}} | {{summary}} | src:... |
| **Energy** | {{state code name}} | {{IECC/ASHRAE/etc.}} | {{year}} | current | {{date/null}} | {{date/null}} | {{date/null}} | {{date/null}} | {{summary}} | src:... |
| **Fire - construction references** | {{state code name}} | {{IFC/NFPA/IBC Ch. refs/etc.}} | {{year}} | current | {{date/null}} | {{date/null}} | {{date/null}} | {{date/null}} | {{summary}} | src:... |
| **Fire - operational / prevention code** | {{state code name}} | {{IFC/NFPA/state fire code/etc.}} | {{year/null}} | current | {{date/null}} | {{date/null}} | {{date/null}} | {{date/null}} | {{summary}} | src:... |
| **Accessibility** | {{state code name}} | {{ADA/IBC Ch. 11/ICC A117.1/state code/etc.}} | {{year/null}} | current | {{date/null}} | {{date/null}} | {{date/null}} | {{date/null}} | {{summary}} | src:... |
| **Elevator / Conveyance** | {{state code name}} | {{ASME A17.1/etc.}} | {{year/null}} | current | {{date/null}} | {{date/null}} | {{date/null}} | {{date/null}} | {{summary}} | src:... |

### 3.2 Adoption Records

`adoption:usa-{{xx}}:{{code_family}}:{{edition_or_slug}}`

```yaml
adoption_id: "adoption:usa-{{xx}}:{{code_family}}:{{edition_or_slug}}"
state_id: "US-{{XX}}"
code_family: "{{building | residential | existing_building | mechanical | plumbing | fuel_gas | electrical | energy | fire_construction | fire_operational | accessibility | elevator}}"
status: "{{current | prior | future | pending | superseded | unknown}}"
state_code:
  name: "{{State-specific code name}}"
  edition_label: "{{edition label | null}}"
  codification: "{{statute/regulation/admin code citation | null}}"
base_model_code:
  publisher: "{{ICC | NFPA | IAPMO | ASHRAE | ASME | state | other | null}}"
  code_name: "{{model code name | null}}"
  edition_year: {{year | null}}
  incorporated_by_reference: {{true | false | null}}
authority:
  adopting_authority_id: "auth:..."
  enforcing_authority_model: "{{state | local | hybrid | special | unknown}}"
  interpretation_authority_id: "auth:... | null"
dates:
  adoption_date: "{{YYYY-MM-DD | null}}"
  effective_date: "{{YYYY-MM-DD | null}}"
  operative_date: "{{YYYY-MM-DD | null}}"
  mandatory_date: "{{YYYY-MM-DD | null}}"
  replacement_date: "{{YYYY-MM-DD | null}}"
applicability:
  date_trigger: "{{permit_application_date | permit_issuance_date | filing_date | construction_start_date | occupancy_date | project_status | unknown}}"
  applies_to:
    - "{{new_construction | alteration | repair | addition | change_of_occupancy | existing_buildings | detached_one_two_family | townhouses | commercial | state_owned | public_school | etc.}}"
  exclusions: []
  special_conditions: []
transition:
  exists: {{true | false | null}}
  rule_id: "date-rule:... | null"
  start_date: "{{YYYY-MM-DD | null}}"
  end_date: "{{YYYY-MM-DD | null}}"
  prior_code_allowed: {{true | false | null}}
  prior_code_condition: "{{description | null}}"
amendments:
  state_amended: {{true | false | null}}
  amendment_set_ids:
    - "amendment-set:..."
  amendment_source_ids:
    - "src:..."
provenance:
  source_ids:
    - "src:..."
  field_sources:
    state_code.name: ["src:..."]
    state_code.codification: ["src:..."]
    base_model_code.code_name: ["src:..."]
    base_model_code.edition_year: ["src:..."]
    authority.adopting_authority_id: ["src:..."]
    dates.adoption_date: ["src:..."]
    dates.effective_date: ["src:..."]
    dates.operative_date: ["src:..."]
    applicability.date_trigger: ["src:..."]
    transition.rule_id: ["src:..."]
verification:
  status: "{{verified | partially_verified | needs_review | unresolved}}"
  confidence: null
  notes: null

```

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

{{Plain-English explanation of adoption dates, effective dates, operative dates, mandatory dates, permit-date logic, grace periods, concurrency periods, and pending future codes.}}

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-{{xx}}:{{slug}} | {{code family/scope}} | {{effective_date, operative_date, mandatory_date, grace_period, transition_period, prior_code_concurrency}} |  |  |  |  |  |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| {{family}} | {{code}} | {{date/null}} | {{date/null}} | {{date/null}} | {{date/null}} | {{date/null}} | {{monitoring, resolved, unknown}} |  |  |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-{{xx}}:{{slug}} | {{family}} | {{scope}} | {{trigger}} | {{summary}} | src:... | {{status}} |

---

## 5. State Amendments

### 5.1 Amendment Model

- **State amendment structure:** {{global amendments | chapter amendments | section amendments | insert pages | administrative rule text | state edition | unknown}}
- **Where amendments are published:** {{description}}
- **Amendment parsing status:** {{not_started | partial | complete | requires_ocr | requires_manual_review}}

### 5.2 State Amendment Sources

| Amendment Set ID | Code Family | Publication Format | Source IDs | Parsed? | Notes |
| --- | --- | --- | --- | --- | --- |
| amendment-set:usa-{{xx}}:{{family}}:{{edition}} | {{family}} | {{PDF/HTML/rule/table/model-code-insert/state-edition}} | src:... | {{yes/no/partial}} | {{notes}} |

### 5.3 High-Impact State Amendments

| Code Family | Section | Amendment Type | Summary | Source IDs | Confidence |
| --- | --- | --- | --- | --- | --- |
| {{family}} | {{section}} | {{modify, delete, add, replace}} |  |  |  |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-{{xx}}"
  model: "{{state_enforced | locally_enforced | hybrid | optional_local_enforcement | unknown}}"
  enforcing_entities:
    - "{{municipality | county | fire_district | state_regional_office | certified_third_party | other}}"
  required_officials:
    - "{{building_official | construction_official | fire_marshal | electrical_inspector | plumbing_inspector | plans_examiner | other}}"
  state_reserved_activities:
    - "{{state_buildings | schools | healthcare | high_hazard | elevators | manufactured_housing | other}}"
  source_ids:
    - "src:..."
  verification_status: "{{verified | partially_verified | needs_review | unresolved}}"
  confidence: null
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-{{xx}}"
  model: "{{prohibited_or_state_preempted | allowed_if_stricter | allowed_with_state_approval | administrative_only | municipality_specific_state_adopted | home_rule_local_choice | unknown}}"
  applies_to_code_families:
    - "{{family}}"
  approval_required: {{true | false | null}}
  approving_authority_id: "auth:... | null"
  filing_required: {{true | false | null}}
  registry_exists: {{true | false | null}}
  registry_source_ids:
    - "src:..."
  legal_basis_source_ids:
    - "src:..."
  verification_status: "{{verified | needs_legal_text_extraction | unresolved}}"
  confidence: null
```

### 6.3 Local Enforcement vs. Local Amendment Summary

{{Plain-English distinction between who enforces the code locally and whether that local entity may amend, supplement, or override the state code.}}

### 6.4 Known Local Amendment Registries

| Registry ID | Scope | Maintained By | Source ID | Machine Readable? | Parsed? | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| registry:usa-{{xx}}:local-amendments | {{statewide/municipal/county/fire/etc.}} | {{agency}} | src:... | {{yes/no/partial}} | {{yes/no/partial}} | {{notes}} |

### 6.5 Municipality-Specific Known Amendments

| Jurisdiction | Code Family | Amendment Set ID | Approval Status | Effective Date | Source IDs | Parsed? |
| --- | --- | --- | --- | --- | --- | --- |
| {{city/county}} | {{family}} | local-amendment:... | {{approved, filed, rejected, unknown}} |  |  |  |

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

- **Resolver status:** {{not_started | state_only | county_level | municipal_level | ahj_level}}

**Example jurisdiction stack:**

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
| --- | --- | --- | --- | --- | --- |
| **State** | {{source}} | src:... | statewide | {{frequency}} | {{status}} |
| **County** | {{source}} | src:... | statewide | {{frequency}} | {{status}} |
| **Municipality** | {{source}} | src:... | statewide | {{frequency}} | {{status}} |
| **Fire District** | {{source}} | src:... | {{coverage}} | {{frequency}} | {{status}} |
| **Special District** | {{source}} | src:... | {{coverage}} | {{frequency}} | {{status}} |

### 7.3 AHJ Contact Data

| AHJ ID | Jurisdiction | Department | Role | Website | Phone | Email | Last Verified | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| ahj:... | {{jurisdiction}} | {{department}} | {{building/fire/electrical/etc.}} | {{url}} | {{phone}} | {{email}} | {{date}} | src:... |

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Source Type | Publisher | URL | Accessed Date | Snapshot ID | Checksum | Status |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-{{xx}}:statute:{{slug}} | {{title}} | statute | {{publisher}} | {{url}} | {{date}} | {{snapshot}} | {{checksum}} | {{status}} |
| src:usa-{{xx}}:regulation:{{slug}} | {{title}} | regulation | {{publisher}} | {{url}} | {{date}} | {{snapshot}} | {{checksum}} | {{status}} |
| src:usa-{{xx}}:agency:{{slug}} | {{title}} | agency_page | {{publisher}} | {{url}} | {{date}} | {{snapshot}} | {{checksum}} | {{status}} |
| src:usa-{{xx}}:amendments:{{family}}:{{edition}} | {{title}} | amendment_pdf | {{publisher}} | {{url}} | {{date}} | {{snapshot}} | {{checksum}} | {{status}} |
| src:usa-{{xx}}:register:{{slug}} | {{title}} | state_register_notice | {{publisher}} | {{url}} | {{date}} | {{snapshot}} | {{checksum}} | {{status}} |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:... | {{courtesy_copy, unofficial_html, official_pdf}} |  |  |

### 8.3 Supplemental Sources

| Source ID | Title | Source Type | Publisher | URL | Use | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| src:... | {{title}} | secondary_summary | {{publisher}} | {{url}} | supplemental_only | {{notes}} |

### 8.4 Source Extraction Metadata

| Source ID | Parser | Parser Version | Extracted At | Extraction Confidence | Requires OCR | Requires Browser Automation | Manual Review Required |
| --- | --- | --- | --- | --- | --- | --- | --- |
| src:... | {{parser}} | {{version}} | {{timestamp}} | {{0.00}} | {{yes/no}} | {{yes/no}} | {{yes/no}} |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| adoption:... | state_code.name | {{value}} | {{status}} | {{0.00}} | src:... | {{notes}} |
| adoption:... | base_model_code.edition_year | {{value}} | {{status}} | {{0.00}} | src:... | {{notes}} |
| adoption:... | dates.effective_date | {{date}} | {{status}} | {{0.00}} | src:... | {{notes}} |
| adoption:... | dates.operative_date | {{date}} | {{status}} | {{0.00}} | src:... | {{notes}} |
| date-rule:... | trigger_condition | {{value}} | {{status}} | {{0.00}} | src:... | {{notes}} |
| local-enforcement:... | model | {{value}} | {{status}} | {{0.00}} | src:... | {{notes}} |
| local-amendment-rule:... | model | {{value}} | {{status}} | {{0.00}} | src:... | {{notes}} |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| **All source IDs resolve** | {{pass/fail}} | {{notes}} |
| **All authority IDs resolve** | {{pass/fail}} | {{notes}} |
| **All current code families have adoption records** | {{pass/fail}} | {{notes}} |
| **Building and operational fire code are separated** | {{pass/fail}} | {{notes}} |
| **Adoption/effective/operative/mandatory dates are not conflated** | {{pass/fail}} | {{notes}} |
| **Effective dates are valid ISO dates** | {{pass/fail}} | {{notes}} |
| **No impossible date sequences** | {{pass/fail}} | {{notes}} |
| **Transition rules have explicit trigger conditions** | {{pass/fail}} | {{notes}} |
| **Permit-date logic is captured where applicable** | {{pass/fail}} | {{notes}} |
| **Local enforcement model classified** | {{pass/fail}} | {{notes}} |
| **Local amendment rule classified** | {{pass/fail}} | {{notes}} |
| **AHJ confirmation metadata present** | {{pass/fail}} | {{notes}} |
| **Official-source caveats captured** | {{pass/fail}} | {{notes}} |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-{{xx}}:001 | {{critical/high/medium/low}} | {{topic}} | {{description}} | {{action}} | {{owner/null}} | {{date/null}} | {{open/in_progress/resolved}} |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-{{xx}}:agency-code-page | src:... | html_diff | monthly | code edition/date text changes | {{date/null}} | {{active/pending}} |
| watch:usa-{{xx}}:amendment-pdf | src:... | checksum_diff | quarterly | PDF hash changes | {{date/null}} | {{active/pending}} |
| watch:usa-{{xx}}:state-register | src:... | register_notice_scan | monthly | proposed/adopted rule activity | {{date/null}} | {{active/pending}} |
| watch:usa-{{xx}}:statute | src:... | statute_text_diff | quarterly | statutory authority changes | {{date/null}} | {{active/pending}} |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| {{date}} | {{description}} | record:... | src:... | {{person/system}} | {{notes}} |