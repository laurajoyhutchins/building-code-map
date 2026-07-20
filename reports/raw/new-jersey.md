---
state:
  state_id: "US-NJ"
  name: "New Jersey"
  abbreviation: "NJ"
report:
  report_id: "state-report:usa-nj"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-26"
  last_verified: null
  reviewed_by: null
risk:
  overall_confidence: 0.64 # 0.00 - 1.00
  risk_flags:
    - "operational_fire_code_primary_text_not_fully_verified"
    - "state_amendments_not_fully_parsed"
    - "local_fire_maintenance_code_scope_requires_review"
    - "ahj_contact_data_not_populated"
  open_questions_count: 5

---

# State Building Code Authority Report: New Jersey

## 1. Executive Summary

- **Authority model:** New Jersey uses a statewide Uniform Construction Code (UCC) model administered by the New Jersey Department of Community Affairs (DCA), with the Commissioner authorized to adopt the State Uniform Construction Code. The UCC is implemented through N.J.A.C. 5:23 and is intended to provide uniform construction standards throughout the State.

- **Statewide code status:** DCA's current construction-code page lists the Building, Electrical, Energy, Mechanical, One- and Two-Family Dwelling, Fuel Gas, Barrier Free, and Elevator subcodes on the 2021/2020 code cycle with September 6, 2022 adoption-date labels, and lists the Plumbing Subcode on the 2021 National Standard Plumbing Code cycle with a September 19, 2022 adoption-date label. The Rehabilitation Subcode is maintained at N.J.A.C. 5:23-6 and listed as current as of March 6, 2023.

- **Local enforcement model:** The UCC is statewide, but ordinary administration and enforcement are performed through municipal enforcing agencies, construction officials, and subcode officials, with DCA oversight and certain State-reserved activities.

- **Local amendment posture:** For UCC-covered construction standards, N.J.A.C. 5:23 states that the subcode provisions apply uniformly throughout New Jersey and that standards other than those provided in the UCC are void and of no effect. This does not eliminate local zoning, trade/profession licensing, or limited local fire-safety maintenance authority, but local fire maintenance codes cannot conflict with the State Uniform Construction Code.

- **Known transition periods or pending changes:** N.J.A.C. 5:23-1.6 provides a six-month grace period after the operative date of a subcode revision for qualifying complete permit applications and projects already under review. DCA also publishes rule proposals and adoptions that should be monitored for the 2024 model-code cycle; this report did not treat those items as current statewide code adoptions without parsing the final adopted rule text and current-code table.

- **Production readiness:** limited_internal_use

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: The DCA Commissioner is the primary statewide authority for adopting the
    State Uniform Construction Code.
  confidence: 0.82
  source_ids:
  - src:usa-nj:ucc-act-dca-pdf
  - src:usa-nj:njac-5-23-1
- topic: Primary building code edition
  finding: DCA lists the Building Subcode as International Building Code/2021, New
    Jersey edition.
  confidence: 0.86
  source_ids:
  - src:usa-nj:dca-current-codes
  - src:usa-nj:njac-5-23-3
  - src:usa-nj:2021-icode-2020-nec-adoption
- topic: Electrical code authority and edition
  finding: DCA lists the Electrical Subcode as National Electrical Code (NFPA 70)/2020.
  confidence: 0.84
  source_ids:
  - src:usa-nj:dca-current-codes
  - src:usa-nj:2021-icode-2020-nec-adoption
- topic: Fire code authority
  finding: Fire-construction references are part of the UCC; operational fire/prevention
    code authority is associated with DCA Division of Fire Safety and N.J.A.C. 5:70,
    but the current primary text was not fully parsed.
  confidence: 0.55
  source_ids:
  - src:usa-nj:dca-current-codes
  - src:usa-nj:dca-dfs-codes-regs
  - src:usa-nj:fire-code-2023-proposal
- topic: Local enforcement
  finding: Municipal construction officials and subcode officials administer and enforce
    the UCC locally, subject to DCA oversight and State-reserved review/enforcement.
  confidence: 0.78
  source_ids:
  - src:usa-nj:njac-5-23-4
  - src:usa-nj:ucc-act-dca-pdf
- topic: Local amendments
  finding: UCC-covered construction standards are statewide and preempt conflicting
    local construction standards; local zoning, licensing, and non-conflicting fire-safety
    maintenance authority remain distinct issues.
  confidence: 0.72
  source_ids:
  - src:usa-nj:njac-5-23-3
  - src:usa-nj:ucc-act-dca-pdf
- topic: Effective / operative date rule
  finding: Unless a New Jersey Register adoption notice provides otherwise, the effective
    date is the operative date; N.J.A.C. 5:23-1.6 supplies a six-month grace period
    for qualifying projects after a subcode revision.
  confidence: 0.8
  source_ids:
  - src:usa-nj:njac-5-23-1
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-nj:dca-commissioner-and-division-codes-standards |
| Authority name | New Jersey Department of Community Affairs, Commissioner / Division of Codes and Standards |
| Authority type | state_department |
| Legal basis | State Uniform Construction Code Act, N.J.S.A. 52:27D-119 et seq.; N.J.A.C. 5:23 |
| Role | Adopt and administer the State Uniform Construction Code; adopt subcodes; publish and maintain UCC regulations; monitor local enforcing agencies; perform State-reserved review/enforcement where required. |
| Enforcement model | statewide_code_with_local_enforcing_agencies_and_state_reserved_activities |
| Source IDs | src:usa-nj:ucc-act-dca-pdf; src:usa-nj:njac-5-23-1; src:usa-nj:njac-5-23-3; src:usa-nj:njac-5-23-4 |
| Verification status | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-nj:dca-division-codes-standards | DCA Division of Codes and Standards | UCC subcode adoption, maintenance, and oversight | N.J.A.C. 5:23-3.14; UCC Act | src:usa-nj:dca-current-codes; src:usa-nj:njac-5-23-3; src:usa-nj:2021-icode-2020-nec-adoption | partially_verified |
| Residential | ahj:usa-nj:dca-division-codes-standards | DCA Division of Codes and Standards | One- and two-family dwelling subcode adoption, maintenance, and oversight | N.J.A.C. 5:23-3.21; UCC Act | src:usa-nj:dca-current-codes; src:usa-nj:njac-5-23-3; src:usa-nj:2021-icode-2020-nec-adoption | partially_verified |
| Existing Building / Rehabilitation | ahj:usa-nj:dca-division-codes-standards | DCA Division of Codes and Standards | Rehabilitation subcode adoption, maintenance, and oversight | N.J.A.C. 5:23-6 | src:usa-nj:dca-current-codes; src:usa-nj:njac-5-23-6 | partially_verified |
| Mechanical | ahj:usa-nj:dca-division-codes-standards | DCA Division of Codes and Standards | Mechanical subcode adoption, maintenance, and oversight | N.J.A.C. 5:23-3.20; UCC Act | src:usa-nj:dca-current-codes; src:usa-nj:njac-5-23-3; src:usa-nj:2021-icode-2020-nec-adoption | partially_verified |
| Plumbing | ahj:usa-nj:dca-division-codes-standards | DCA Division of Codes and Standards | Plumbing subcode adoption, maintenance, and oversight, with statutory consultation requirement involving health authorities before adoption | N.J.A.C. 5:23-3.15; UCC Act | src:usa-nj:dca-current-codes; src:usa-nj:njac-5-23-3; src:usa-nj:ucc-act-dca-pdf | partially_verified |
| Fuel Gas | ahj:usa-nj:dca-division-codes-standards | DCA Division of Codes and Standards | Fuel gas subcode adoption, maintenance, and oversight | N.J.A.C. 5:23-3.22; UCC Act | src:usa-nj:dca-current-codes; src:usa-nj:njac-5-23-3; src:usa-nj:2021-icode-2020-nec-adoption | partially_verified |
| Electrical | ahj:usa-nj:dca-division-codes-standards | DCA Division of Codes and Standards | Electrical subcode adoption, maintenance, and oversight | N.J.A.C. 5:23-3.16; UCC Act | src:usa-nj:dca-current-codes; src:usa-nj:njac-5-23-3; src:usa-nj:2021-icode-2020-nec-adoption | partially_verified |
| Energy | ahj:usa-nj:dca-division-codes-standards | DCA Division of Codes and Standards | Energy subcode adoption, maintenance, and oversight | N.J.A.C. 5:23-3.18; UCC Act | src:usa-nj:dca-current-codes; src:usa-nj:njac-5-23-3; src:usa-nj:2021-icode-2020-nec-adoption | partially_verified |
| Fire - construction references | ahj:usa-nj:dca-division-codes-standards | DCA Division of Codes and Standards | UCC fire-protection/construction requirements and referenced fire provisions | N.J.A.C. 5:23-3.17; UCC Act | src:usa-nj:dca-current-codes; src:usa-nj:njac-5-23-3; src:usa-nj:2021-icode-2020-nec-adoption | partially_verified |
| Fire - operational / prevention code | ahj:usa-nj:dca-division-fire-safety | DCA Division of Fire Safety | State Fire Prevention Code / Uniform Fire Code administration and rulemaking support | N.J.A.C. 5:70; Uniform Fire Safety Act and DCA fire-code rulemaking page | src:usa-nj:dca-dfs-codes-regs; src:usa-nj:fire-code-2023-proposal; src:usa-nj:ucc-act-dca-pdf | unresolved_primary_text |
| Accessibility | ahj:usa-nj:dca-division-codes-standards | DCA Division of Codes and Standards | Barrier Free Subcode adoption, maintenance, and oversight | Chapter 11 of IBC/2021 and N.J.A.C. 5:23-7 | src:usa-nj:dca-current-codes; src:usa-nj:njac-5-23-3 | partially_verified |
| Elevator / Conveyance | ahj:usa-nj:dca-division-codes-standards | DCA Division of Codes and Standards | Elevator subcode adoption, State-reserved plan review where applicable, and oversight | N.J.A.C. 5:23-12; ASME references in IBC/2021 Chapter 35 | src:usa-nj:dca-current-codes; src:usa-nj:njac-5-23-3; src:usa-nj:njac-5-23-4 | partially_verified |

### 2.3 Authority Hierarchy Notes

New Jersey has a hybrid administration model: the State adopts and maintains a uniform construction code, while municipal enforcing agencies administer permits, plan review, inspection, and enforcement for ordinary local projects. DCA retains oversight authority and may perform or reserve plan review/enforcement for categories specified in N.J.A.C. 5:23, including certain State-reserved activities. Operational fire-code authority is related but should not be merged with UCC fire-protection construction requirements until the current N.J.A.C. 5:70 text is parsed.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-nj:001 | ahj:usa-nj:dca-commissioner-and-division-codes-standards | adopts | State Uniform Construction Code and individual UCC subcodes | src:usa-nj:ucc-act-dca-pdf; src:usa-nj:njac-5-23-3 | partially_verified |
| edge:usa-nj:002 | ahj:usa-nj:dca-commissioner-and-division-codes-standards | delegates_ordinary_enforcement_to | municipal enforcing agencies, construction officials, and subcode officials | src:usa-nj:njac-5-23-4; src:usa-nj:ucc-act-dca-pdf | partially_verified |
| edge:usa-nj:003 | ahj:usa-nj:dca-commissioner-and-division-codes-standards | reserves_review_for | categories listed in N.J.A.C. 5:23, including State-reserved plan review and inspections where specified | src:usa-nj:njac-5-23-3; src:usa-nj:njac-5-23-4 | partially_verified |
| edge:usa-nj:004 | ahj:usa-nj:dca-commissioner-and-division-codes-standards | preempts | non-UCC local construction standards for UCC-covered subject matter | src:usa-nj:njac-5-23-3 | partially_verified |
| edge:usa-nj:005 | local_fire_authorities | may_adopt_nonconflicting_fire_maintenance_code | local fire-safety maintenance provisions, separate from UCC fire-prevention-related construction activities | src:usa-nj:ucc-act-dca-pdf | partially_verified |
| edge:usa-nj:006 | ahj:usa-nj:dca-division-fire-safety | administers_or_supports | State Fire Prevention Code / Uniform Fire Code rulemaking and code access | src:usa-nj:dca-dfs-codes-regs; src:usa-nj:fire-code-2023-proposal | unresolved_primary_text |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | Building Subcode, N.J.A.C. 5:23-3.14 | International Building Code, New Jersey edition | 2021 | current | 2022-08-01 | 2022-09-06 | 2022-09-06 | 2023-03-06 | Six-month UCC grace period after operative date for qualifying complete permit applications; DCA current-code table labels adoption date as 2022-09-06. | src:usa-nj:dca-current-codes; src:usa-nj:2021-icode-2020-nec-adoption; src:usa-nj:njac-5-23-1 |
| Residential | One- and Two-Family Dwelling Subcode, N.J.A.C. 5:23-3.21 | International Residential Code, New Jersey edition | 2021 | current | 2022-08-01 | 2022-09-06 | 2022-09-06 | 2023-03-06 | Six-month UCC grace period after operative date for qualifying complete permit applications; DCA current-code table labels adoption date as 2022-09-06. | src:usa-nj:dca-current-codes; src:usa-nj:2021-icode-2020-nec-adoption; src:usa-nj:njac-5-23-1 |
| Existing Building / Rehabilitation | Rehabilitation Subcode, N.J.A.C. 5:23-6 | New Jersey rehabilitation subcode | state-specific | current | null | null | null | null | Maintained separately and updated as necessary; no single current adoption/effective date normalized during this update. | src:usa-nj:dca-current-codes; src:usa-nj:njac-5-23-6 |
| Mechanical | Mechanical Subcode, N.J.A.C. 5:23-3.20 | International Mechanical Code | 2021 | current | 2022-08-01 | 2022-09-06 | 2022-09-06 | 2023-03-06 | Six-month UCC grace period after operative date for qualifying complete permit applications; DCA current-code table labels adoption date as 2022-09-06. | src:usa-nj:dca-current-codes; src:usa-nj:2021-icode-2020-nec-adoption; src:usa-nj:njac-5-23-1 |
| Plumbing | Plumbing Subcode, N.J.A.C. 5:23-3.15 | National Standard Plumbing Code Illustrated, New Jersey edition | 2021 | current | 2022-09-19 | null | null | null | DCA current-code table provides a 2022-09-19 adoption-date label; final rule notice dates were not separately parsed during this update. | src:usa-nj:dca-current-codes; src:usa-nj:njac-5-23-3 |
| Fuel Gas | Fuel Gas Subcode, N.J.A.C. 5:23-3.22 | International Fuel Gas Code | 2021 | current | 2022-08-01 | 2022-09-06 | 2022-09-06 | 2023-03-06 | Six-month UCC grace period after operative date for qualifying complete permit applications; DCA current-code table labels adoption date as 2022-09-06. | src:usa-nj:dca-current-codes; src:usa-nj:2021-icode-2020-nec-adoption; src:usa-nj:njac-5-23-1 |
| Electrical | Electrical Subcode, N.J.A.C. 5:23-3.16 | National Electrical Code (NFPA 70) | 2020 | current | 2022-08-01 | 2022-09-06 | 2022-09-06 | 2023-03-06 | Six-month UCC grace period after operative date for qualifying complete permit applications; DCA current-code table labels adoption date as 2022-09-06. | src:usa-nj:dca-current-codes; src:usa-nj:2021-icode-2020-nec-adoption; src:usa-nj:njac-5-23-1 |
| Energy | Energy Subcode, N.J.A.C. 5:23-3.18 | IECC for low-rise residential; ASHRAE 90.1 for commercial and all other residential | IECC 2021; ASHRAE 90.1-2019 | current | 2022-08-01 | 2022-09-06 | 2022-09-06 | 2023-03-06 | Six-month UCC grace period after operative date for qualifying complete permit applications; DCA current-code table labels adoption date as 2022-09-06. | src:usa-nj:dca-current-codes; src:usa-nj:2021-icode-2020-nec-adoption; src:usa-nj:njac-5-23-1 |
| Fire - construction references | Fire Protection Subcode and fire-related UCC construction provisions | 2021 I-Code family references, including IFC references where incorporated by UCC | 2021 | current | 2022-08-01 | 2022-09-06 | 2022-09-06 | 2023-03-06 | Treated as UCC construction code; do not use as a substitute for operational fire/prevention code. | src:usa-nj:dca-current-codes; src:usa-nj:njac-5-23-3; src:usa-nj:2021-icode-2020-nec-adoption; src:usa-nj:njac-5-23-1 |
| Fire - operational / prevention code | State Fire Prevention Code / Uniform Fire Code, N.J.A.C. 5:70 | 2018 IFC New Jersey edition is listed by DCA Division of Fire Safety; current primary text not fully parsed | 2018 | partially_verified | null | null | null | null | Operational fire/prevention code remains open for primary-text verification; keep distinct from UCC fire-construction references. | src:usa-nj:dca-dfs-codes-regs; src:usa-nj:fire-code-2023-proposal |
| Accessibility | Barrier Free Subcode | Chapter 11 of IBC/2021 and N.J.A.C. 5:23-7; ICC A117.1-2017 | 2021 / 2017 | current | 2022-09-06 | null | null | null | DCA current-code table provides a 2022-09-06 adoption-date label; final rule notice dates were not separately parsed for the Barrier Free row. | src:usa-nj:dca-current-codes; src:usa-nj:njac-5-23-3 |
| Elevator / Conveyance | Elevator Subcode, N.J.A.C. 5:23-12 | ASME standards as referenced in IBC/2021 Chapter 35 | as referenced | current | 2022-09-06 | null | null | null | DCA current-code table provides a 2022-09-06 adoption-date label; State-reserved review details require subchapter-level parsing. | src:usa-nj:dca-current-codes; src:usa-nj:njac-5-23-3; src:usa-nj:njac-5-23-4 |

### 3.2 Adoption Records

#### record:usa-nj:ucc-2021-icodes-2020-nec

| Field | Value |
| --- | --- |
| Applies to | Building, Residential, Mechanical, Fuel Gas, Electrical, Energy, Fire - construction references |
| Rule / action | R.2022 d.111 adopted amendments to N.J.A.C. 5:23-3.14, 5:23-3.16, 5:23-3.17, 5:23-3.18, 5:23-3.20, 5:23-3.21, and 5:23-3.22 |
| Adopted by Commissioner | 2022-08-01 |
| Filed | 2022-08-12 |
| Effective | 2022-09-06 |
| Operative | 2022-09-06 unless a notice provision states otherwise |
| Mandatory | 2023-03-06 as a conservative normalized six-month post-operative marker for UCC subcode revisions; exact last day of prior-code use should be confirmed in production |
| Source IDs | src:usa-nj:2021-icode-2020-nec-adoption; src:usa-nj:njac-5-23-1; src:usa-nj:dca-current-codes |
| Status | partially_verified |

#### record:usa-nj:ucc-plumbing-2021-nspc

| Field | Value |
| --- | --- |
| Applies to | Plumbing |
| Rule / action | DCA current-code table lists Plumbing Subcode, N.J.A.C. 5:23-3.15, as National Standard Plumbing Code/2021, New Jersey edition |
| Adopted by Commissioner | null |
| Filed | null |
| Effective | null |
| Operative | null |
| Mandatory | null |
| Source IDs | src:usa-nj:dca-current-codes; src:usa-nj:njac-5-23-3 |
| Status | partially_verified |

#### record:usa-nj:rehabilitation-subcode

| Field | Value |
| --- | --- |
| Applies to | Existing Building / Rehabilitation |
| Rule / action | DCA current-code table lists N.J.A.C. 5:23-6 as the Rehabilitation Subcode, updated as necessary and current as of 2023-03-06 |
| Adopted by Commissioner | null |
| Filed | null |
| Effective | null |
| Operative | null |
| Mandatory | null |
| Source IDs | src:usa-nj:dca-current-codes; src:usa-nj:njac-5-23-6 |
| Status | partially_verified |

#### record:usa-nj:operational-fire-code-open

| Field | Value |
| --- | --- |
| Applies to | Fire - operational / prevention code |
| Rule / action | DCA Division of Fire Safety lists a 2018 IFC New Jersey edition item and N.J.A.C. 5:70 rulemaking materials, but the current primary code text and final adoption sequence were not fully parsed |
| Adopted by Commissioner | null |
| Filed | null |
| Effective | null |
| Operative | null |
| Mandatory | null |
| Source IDs | src:usa-nj:dca-dfs-codes-regs; src:usa-nj:fire-code-2023-proposal |
| Status | unresolved_primary_text |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

New Jersey distinguishes effective and operative dates in N.J.A.C. 5:23-1.4. Unless a New Jersey Register adoption notice provides otherwise, the effective date is the operative date. N.J.A.C. 5:23-1.6 provides a six-month grace period following the operative date of a subcode revision. During that period, a complete permit application that includes prior approvals may be reviewed under the immediately preceding code; projects already under review before the operative date may continue under the immediately preceding code.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-nj:001 | UCC rules and subcode revisions | effective_to_operative_default | effective date equals operative date unless adoption notice states otherwise | Effective rule published/adopted without a different operative date in the New Jersey Register notice | no special prior-code allowance from this rule alone | src:usa-nj:njac-5-23-1 | verified |
| date-rule:usa-nj:002 | UCC subcode revisions | six_month_grace_period | six months after operative date | Complete permit application, including prior approvals, submitted during grace period; or project already under review before operative date | yes, under immediately preceding code for qualifying applications/projects | src:usa-nj:njac-5-23-1 | verified |
| date-rule:usa-nj:003 | 2021 I-Code / 2020 NEC UCC adoption package | adoption_notice_effective_date | 2022-09-06 effective / operative normalized | R.2022 d.111 adoption package for listed subcodes | yes, subject to six-month UCC grace rule | src:usa-nj:2021-icode-2020-nec-adoption; src:usa-nj:njac-5-23-1 | partially_verified |
| date-rule:usa-nj:004 | Operational fire/prevention code | unresolved | null | Current N.J.A.C. 5:70 primary text and final adoption notice not fully parsed | unresolved | src:usa-nj:dca-dfs-codes-regs; src:usa-nj:fire-code-2023-proposal | unresolved_primary_text |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| UCC subcodes | 2024 model-code cycle rulemaking materials | null | null | null | null | null | monitor | src:usa-nj:dca-rule-proposals-adoptions; src:usa-nj:dca-current-codes | DCA rulemaking page includes 2024 model-code cycle items. This report does not treat those as current until the final adopted text and current-code page are reconciled. |
| Operational fire/prevention code | N.J.A.C. 5:70 and related fire-code proposals | 2026-01-01 | null | null | null | null | monitor | src:usa-nj:dca-dfs-codes-regs | DCA Division of Fire Safety page lists 2026 rule proposals; no adopted current-code change was normalized during this update. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-nj:001 | UCC subcodes | Projects subject to a revised subcode | Subcode revision becomes operative | Six-month transition pathway for qualifying complete permit applications and projects already under review. | src:usa-nj:njac-5-23-1 | verified |
| applicability-rule:usa-nj:002 | Existing Building / Rehabilitation | Work in pre-existing buildings | Work scope falls within Rehabilitation Subcode | Rehabilitation Subcode is intended to create a uniform statewide rehabilitation framework while preserving certain relationships with fire, hotel, housing, and other regulations. | src:usa-nj:njac-5-23-6 | partially_verified |
| applicability-rule:usa-nj:003 | Fire - operational / prevention code | Operational fire safety / maintenance | Project or occupancy implicates N.J.A.C. 5:70 rather than UCC construction provisions | Keep operational fire code separate from construction-code fire protection subcode until N.J.A.C. 5:70 primary text is fully parsed. | src:usa-nj:dca-dfs-codes-regs; src:usa-nj:ucc-act-dca-pdf | unresolved_primary_text |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** code-family-specific amendments embedded in N.J.A.C. 5:23 subchapter text, New Jersey editions of model codes, errata/corrected pages, formal technical opinions, construction code communications, and New Jersey Register adoption notices.

**Where amendments are published:** DCA UCC rules and subcode PDFs; DCA Current Construction Codes page; rule-adoption PDFs; incorporated New Jersey editions and corrected pages.

**Amendment parsing status:** partial_high_level_only

### 5.2 State Amendment Sources

| Amendment Source ID | Source IDs | Scope | Parsing Status | Notes |
| --- | --- | --- | --- | --- |
| amendment-source:usa-nj:ucc-subchapter-3 | src:usa-nj:njac-5-23-3 | Subcode adoptions and modifications | partial | Establishes subcode adoption structure and uniformity. Full section-by-section amendments were not extracted. |
| amendment-source:usa-nj:2021-icode-2020-nec | src:usa-nj:2021-icode-2020-nec-adoption | 2021 I-Code and 2020 NEC adoption package | partial | Used for date normalization and code-family coverage; not a full amendment digest. |
| amendment-source:usa-nj:dca-current-codes | src:usa-nj:dca-current-codes | Current code list, corrected pages, errata links | partial | Used as current-code source of record; adoption-date labels may not equal legal adopted/filed/effective dates. |
| amendment-source:usa-nj:ucc-bulletins-ftos-cccs | src:usa-nj:dca-current-codes | Bulletins, formal technical opinions, construction code communications | not_started | Listed on DCA current-code page but not parsed. |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Topic | Summary | Source IDs | Confidence | Status |
| --- | --- | --- | --- | --- | --- | --- |
| amendment:usa-nj:001 | All UCC subcodes | Uniform statewide applicability | N.J.A.C. 5:23-3 states that subcode provisions apply uniformly throughout New Jersey and that standards other than those provided are void and of no effect for covered subject matter. | src:usa-nj:njac-5-23-3 | 0.82 | partially_verified |
| amendment:usa-nj:002 | Building / Residential / Mechanical / Fuel Gas / Energy / Electrical / Fire construction | New Jersey 2021 I-Code / 2020 NEC adoption package | R.2022 d.111 updated multiple UCC subcodes to the 2021 I-Code and 2020 NEC cycle. | src:usa-nj:2021-icode-2020-nec-adoption; src:usa-nj:dca-current-codes | 0.76 | partially_verified |
| amendment:usa-nj:003 | Rehabilitation | State-specific rehabilitation framework | N.J.A.C. 5:23-6 establishes a state-specific rehabilitation subcode for pre-existing buildings. | src:usa-nj:njac-5-23-6; src:usa-nj:dca-current-codes | 0.74 | partially_verified |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-nj"
  model: "statewide_code_with_municipal_enforcement_and_state_reserved_activities"
  enforcing_entities:
    - "municipal enforcing agencies"
    - "construction officials"
    - "subcode officials"
    - "DCA Division of Codes and Standards for State-reserved review/enforcement and oversight"
  required_officials:
    - "construction official"
    - "building subcode official"
    - "electrical subcode official"
    - "plumbing subcode official"
    - "fire protection subcode official"
    - "other subcode officials as applicable"
  state_reserved_activities:
    - "State-reserved plan review and inspection categories identified in N.J.A.C. 5:23"
    - "Department oversight of local enforcing agencies"
    - "Department authority to take corrective action or supplant failing local enforcing agencies under the UCC Act"
  source_ids:
    - "src:usa-nj:njac-5-23-4"
    - "src:usa-nj:ucc-act-dca-pdf"
  verification_status: "partially_verified"
  confidence: 0.78
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-nj"
  model: "state_preemption_for_ucc_construction_standards_with_preserved_zoning_trade_licensing_and_nonconflicting_fire_maintenance_authority"
  applies_to_code_families:
    - "Building"
    - "Residential"
    - "Existing Building / Rehabilitation"
    - "Mechanical"
    - "Plumbing"
    - "Fuel Gas"
    - "Electrical"
    - "Energy"
    - "Fire - construction references"
    - "Accessibility"
    - "Elevator / Conveyance"
  approval_required: null
  approving_authority_id: null
  filing_required: null
  registry_exists: null
  registry_source_ids: []
  legal_basis_source_ids:
    - "src:usa-nj:njac-5-23-3"
    - "src:usa-nj:ucc-act-dca-pdf"
  verification_status: "partially_verified"
  confidence: 0.72
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement and local amendment authority must be kept separate in New Jersey. Municipal enforcing agencies administer the statewide UCC, but that enforcement role does not imply authority to amend statewide construction standards. For UCC-covered construction standards, the subcodes are statewide and conflicting local standards are void. Separate local zoning, licensing, and fire-safety maintenance powers may still exist, but the fire-maintenance pathway cannot be used to create construction-code requirements that conflict with the State Uniform Construction Code.

### 6.4 Known Local Amendment Registries

No statewide registry of local construction-code amendments was verified. This is expected under a statewide preemption model, but the absence of a registry should still be confirmed for local fire-maintenance codes and zoning-related overlays before production use.

### 6.5 Municipality-Specific Known Amendments

No municipality-specific construction-code amendments were parsed. Local zoning, land-use, floodplain, historic-preservation, and fire-maintenance ordinances remain outside the scope of this pass unless they purport to modify UCC construction standards.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: partial_model_only

Jurisdiction stack:

```text
Address
  -> State of New Jersey
  -> County
  -> Municipality
  -> Municipal enforcing agency / construction official
  -> Applicable UCC subcode officials
  -> DCA State-reserved review, if project category triggers State review
  -> Fire AHJ / DCA Division of Fire Safety for operational fire-code issues
  -> Applicable statewide UCC adoption records
  -> Applicable non-conflicting local zoning, land-use, fire-maintenance, or special-district overlays
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

No AHJ contact data was populated for this pass. Production AHJ resolution requires a municipal enforcing-agency directory, fire-code official directory, DCA State-reserved review contact mapping, and a policy for municipalities that use shared or interlocal code-enforcement arrangements.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Publisher / Custodian | Source Type | URL | Accessed | Supports | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-nj:ucc-act-dca-pdf | State Uniform Construction Code Act, N.J.S.A. 52:27D-119 et seq. | New Jersey Department of Community Affairs / LexisNexis courtesy copy | statute_courtesy_pdf | https://www.nj.gov/dca/codes/codreg/pdf_regs/52_27D_119.pdf | 2026-06-26 | State adopting authority; Commissioner powers; local enforcement oversight; local zoning/licensing carveouts; local fire maintenance-code limitations | partially_verified |
| src:usa-nj:njac-5-23-1 | N.J.A.C. 5:23-1, General Provisions | New Jersey Department of Community Affairs / LexisNexis courtesy copy | administrative_rule_courtesy_pdf | https://www.nj.gov/dca/codes/codreg/pdf_regs/njac_5_23_1.pdf | 2026-06-26 | Authority basis; definitions; effective/operative date rule; six-month grace-period rule | partially_verified |
| src:usa-nj:njac-5-23-3 | N.J.A.C. 5:23-3, Subcodes | New Jersey Department of Community Affairs / LexisNexis courtesy copy | administrative_rule_courtesy_pdf | https://www.nj.gov/dca/codes/codreg/pdf_regs/njac_5_23_3.pdf | 2026-06-26 | Subcode structure; code adoptions; statewide uniformity; State-reserved activities | partially_verified |
| src:usa-nj:njac-5-23-4 | N.J.A.C. 5:23-4, Enforcing Agencies | New Jersey Department of Community Affairs / LexisNexis courtesy copy | administrative_rule_courtesy_pdf | https://www.nj.gov/dca/codes/codreg/pdf_regs/njac_5_23_4.pdf | 2026-06-26 | Local enforcing agency administration; construction official duties; State enforcement role | partially_verified |
| src:usa-nj:njac-5-23-6 | N.J.A.C. 5:23-6, Rehabilitation Subcode | New Jersey Department of Community Affairs / LexisNexis courtesy copy | administrative_rule_courtesy_pdf | https://www.nj.gov/dca/codes/codreg/pdf_regs/njac_5_23_6.pdf | 2026-06-26 | Rehabilitation subcode scope and statewide rehabilitation framework | partially_verified |
| src:usa-nj:dca-current-codes | Current Construction Codes | New Jersey Department of Community Affairs, Division of Codes and Standards | official_agency_webpage | https://www.nj.gov/dca/codes/codreg/current.shtml | 2026-06-26 | Current code matrix; code family editions; DCA adoption-date labels; errata and official code links | partially_verified |
| src:usa-nj:dca-ucc-portal | UCC Code and Regulations portal | New Jersey Department of Community Affairs, Division of Codes and Standards | official_agency_webpage | https://www.nj.gov/dca/codes/codreg/ucc.shtml | 2026-06-26 | Portal for UCC chapter PDFs and caveat that online PDFs are courtesy copies | verified_for_caveat |
| src:usa-nj:2021-icode-2020-nec-adoption | Adopted Amendments: N.J.A.C. 5:23-3.14, 3.16, 3.17, 3.18, 3.20, 3.21, and 3.22; 2021 I-Code / 2020 NEC package | New Jersey Department of Community Affairs | rule_adoption_pdf | https://www.nj.gov/dca/codes/codreg/pdf_rule_adoptions/2021_ICode_2020_NEC_adopt.pdf | 2026-06-26 | R.2022 d.111 adoption package; adopted/filed/effective dates for major UCC subcode updates | partially_verified |
| src:usa-nj:dca-rule-proposals-adoptions | Rule Proposals and Adoptions | New Jersey Department of Community Affairs, Division of Codes and Standards | official_agency_webpage | https://www.nj.gov/dca/codes/codreg/rule_proposals_adoptions.shtml | 2026-06-26 | Monitoring source for pending and adopted UCC rulemaking | partially_verified |
| src:usa-nj:dca-dfs-codes-regs | Division of Fire Safety Codes and Regulations | New Jersey Department of Community Affairs, Division of Fire Safety | official_agency_webpage | https://www.nj.gov/dca/dfs/codes/ | 2026-06-26 | Fire-code rule proposals/adoptions and access point for N.J.A.C. 5:70-76 | partially_verified |
| src:usa-nj:fire-code-2023-proposal | 2018 IFC / N.J.A.C. 5:70 proposal materials | New Jersey Department of Community Affairs, Division of Fire Safety | rule_proposal_pdf | https://www.nj.gov/dca/dfs/pdf/rules/proposal_2018_ifc.pdf | 2026-06-26 | Context for proposed State Fire Prevention Code / 2018 IFC rulemaking; not a final-adoption substitute | supplemental_for_fire_context |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-nj:ucc-act-dca-pdf | courtesy_copy | DCA-hosted LexisNexis statute PDF is useful for extraction but should be checked against official statutory publication for production. | verify_against_official_statutes |
| src:usa-nj:njac-5-23-1 | courtesy_copy | DCA indicates online UCC chapter PDFs are courtesy copies; official Administrative Rules are available through the New Jersey Office of Administrative Law / LexisNexis arrangement. | verify_against_official_njac |
| src:usa-nj:njac-5-23-3 | courtesy_copy | DCA-hosted UCC chapter PDF is not treated as the final official legal publication without Lexis/OAL confirmation. | verify_against_official_njac |
| src:usa-nj:njac-5-23-4 | courtesy_copy | DCA-hosted UCC chapter PDF is not treated as the final official legal publication without Lexis/OAL confirmation. | verify_against_official_njac |
| src:usa-nj:njac-5-23-6 | courtesy_copy | DCA-hosted UCC chapter PDF is not treated as the final official legal publication without Lexis/OAL confirmation. | verify_against_official_njac |
| src:usa-nj:dca-current-codes | date_semantics | DCA current-code table includes an "Adoption Date" column, but does not always distinguish legal adoption, filing, effective, operative, and mandatory dates. | reconcile_with_rule_notices |
| src:usa-nj:2021-icode-2020-nec-adoption | partial_extraction | Used for major date and coverage facts; full adopted amendments were not parsed section by section. | parse_before_amendment_digest |
| src:usa-nj:dca-dfs-codes-regs | incomplete_primary_text | The Fire Safety page links to N.J.A.C. 5:70-76 and rulemaking materials but this pass did not extract the current official N.J.A.C. 5:70 text. | verify_fire_primary_text |
| src:usa-nj:fire-code-2023-proposal | proposal_not_final | Proposal material is useful context only and should not be cited as final current law without a final adoption notice/current text. | do_not_use_as_final_authority |

### 8.3 Supplemental Sources

No non-government supplemental source was relied on for a final current-code field. Supplemental trade, ICC store, or secondary-law sources may be useful to locate the current fire-code edition, but should not replace primary DCA/OAL/N.J.A.C. evidence.

### 8.4 Source Extraction Metadata

| Extraction ID | Source ID | Extracted By | Extracted On | Method | Coverage | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| extract:usa-nj:001 | src:usa-nj:dca-current-codes | ChatGPT | 2026-06-26 | official web page review | code matrix | Used for current code editions and DCA adoption-date labels. |
| extract:usa-nj:002 | src:usa-nj:ucc-act-dca-pdf | ChatGPT | 2026-06-26 | PDF text extraction and screenshot spot-check | statute authority and local relationship | Courtesy-copy caveat retained. |
| extract:usa-nj:003 | src:usa-nj:njac-5-23-1 | ChatGPT | 2026-06-26 | PDF text extraction and screenshot spot-check | date definitions and grace period | Exact inclusive/exclusive last grace-day should be checked before production. |
| extract:usa-nj:004 | src:usa-nj:njac-5-23-3 | ChatGPT | 2026-06-26 | PDF text extraction and screenshot spot-check | subcode structure and statewide uniformity | Full amendment parse remains open. |
| extract:usa-nj:005 | src:usa-nj:njac-5-23-4 | ChatGPT | 2026-06-26 | PDF text extraction and screenshot spot-check | enforcing agency model | Municipal directory not populated. |
| extract:usa-nj:006 | src:usa-nj:njac-5-23-6 | ChatGPT | 2026-06-26 | PDF text extraction | rehabilitation subcode scope | Adoption history not normalized. |
| extract:usa-nj:007 | src:usa-nj:dca-dfs-codes-regs | ChatGPT | 2026-06-26 | official web page review | fire-code source location and rulemaking context | Current N.J.A.C. 5:70 primary text remains open. |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 0.90 | none | Core authority and construction-code matrix are source-backed; operational fire primary text and amendment digest remain open. |
| report | risk.overall_confidence | 0.64 | verified | 0.80 | none | Moderate confidence for UCC authority and current-code matrix; lower confidence for fire and amendment details. |
| ahj:usa-nj:dca-commissioner-and-division-codes-standards | legal_basis | N.J.S.A. 52:27D-119 et seq.; N.J.A.C. 5:23 | partially_verified | 0.82 | src:usa-nj:ucc-act-dca-pdf; src:usa-nj:njac-5-23-1 | Courtesy-copy caveat applies. |
| matrix:building | edition | International Building Code/2021, New Jersey edition | partially_verified | 0.86 | src:usa-nj:dca-current-codes; src:usa-nj:2021-icode-2020-nec-adoption | Current page and adoption package aligned at high level. |
| matrix:plumbing | edition | National Standard Plumbing Code/2021, New Jersey edition | partially_verified | 0.76 | src:usa-nj:dca-current-codes | DCA current-code table verified; final adoption notice not separately parsed. |
| matrix:electrical | edition | National Electrical Code (NFPA 70)/2020 | partially_verified | 0.84 | src:usa-nj:dca-current-codes; src:usa-nj:2021-icode-2020-nec-adoption | Current page and adoption package aligned at high level. |
| matrix:fire-operational | edition | 2018 IFC New Jersey edition listed by DCA Division of Fire Safety | unresolved_primary_text | 0.50 | src:usa-nj:dca-dfs-codes-regs; src:usa-nj:fire-code-2023-proposal | Current final N.J.A.C. 5:70 text and adoption dates remain open. |
| local-enforcement:usa-nj | model | statewide_code_with_municipal_enforcement_and_state_reserved_activities | partially_verified | 0.78 | src:usa-nj:njac-5-23-4; src:usa-nj:ucc-act-dca-pdf | Municipal directory and interlocal arrangements not populated. |
| local-amendment-rule:usa-nj | model | state_preemption_for_ucc_construction_standards_with_preserved_zoning_trade_licensing_and_nonconflicting_fire_maintenance_authority | partially_verified | 0.72 | src:usa-nj:njac-5-23-3; src:usa-nj:ucc-act-dca-pdf | Local fire-maintenance filing/registry questions remain open. |
| date-rule:usa-nj:001 | effective/operative default | effective date equals operative date unless notice states otherwise | verified | 0.80 | src:usa-nj:njac-5-23-1 | Courtesy-copy caveat applies. |
| date-rule:usa-nj:002 | six-month grace period | qualifying applications/projects may use prior code during grace window | verified | 0.80 | src:usa-nj:njac-5-23-1 | Exact inclusive deadline should be checked before production automation. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Every `src:usa-nj:*` cited in the body appears in section 8. |
| All authority IDs resolve | pass | Authority IDs used in body are declared or described in sections 2.1 and 2.2. |
| All current code families have adoption rows | pass | All template code families remain present with explicit current, partial, or unresolved status. |
| Building and operational fire code are separated | pass | UCC fire-construction references and N.J.A.C. 5:70 operational fire/prevention code are separate rows. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Known dates are separated; unknown dates are `null`. |
| Effective dates are valid ISO dates | pass | Populated date fields use YYYY-MM-DD format. |
| No impossible date sequences | pass | No effective date precedes the adopted-by-Commissioner date in normalized records. |
| Transition rules have explicit trigger conditions | pass | UCC effective/operative and six-month grace rules are stated with triggers. |
| Permit-date logic is captured where applicable | pass | UCC grace-period rule references complete permit applications and projects already under review. |
| Local enforcement model classified | pass | Classified as statewide code with municipal enforcement and State-reserved activities. |
| Local amendment rule classified | pass | Classified with caveat for nonconflicting local fire maintenance and other preserved local powers. |
| AHJ confirmation metadata present | fail | No municipal enforcing-agency or fire-AHJ contact dataset was populated. |
| Official-source caveats captured | pass | Courtesy-copy and proposal/current-text limitations are documented. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-nj:001 | high | operational fire code primary text | Confirm the current text, adopted edition, and adoption/effective/operative dates for N.J.A.C. 5:70 State Fire Prevention Code / Uniform Fire Code. | Extract current N.J.A.C. 5:70 official text and final adoption notices from Lexis/OAL/DCA fire rule-adoption materials. | null | null | open |
| issue:usa-nj:002 | medium | state amendments | Full New Jersey amendments, corrected pages, formal technical opinions, bulletins, and construction code communications were not parsed. | Build code-family amendment digest from N.J.A.C. 5:23 PDFs, adopted NJ editions, errata, FTOs, bulletins, and CCCs. | null | null | open |
| issue:usa-nj:003 | medium | local fire maintenance scope | Local fire maintenance authority exists separately from UCC construction standards, but filing/approval/registry details were not verified. | Review Uniform Fire Safety Act, N.J.A.C. 5:70-76, and municipal/fire-district implementation rules. | null | null | open |
| issue:usa-nj:004 | medium | State-reserved plan review | State-reserved review categories were identified only at a high level. | Parse N.J.A.C. 5:23-3.11 and related provisions into project-type triggers and AHJ routing rules. | null | null | open |
| issue:usa-nj:005 | medium | AHJ contacts and boundary data | Municipal enforcing-agency, fire official, and DCA contact datasets were not populated. | Select boundary datasets and authoritative contact directories; map municipalities and shared-service arrangements. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-nj:dca-current-codes | src:usa-nj:dca-current-codes | html_diff | monthly | Current code edition or adoption-date label changes | 2026-06-26 | active |
| watch:usa-nj:ucc-portal | src:usa-nj:dca-ucc-portal | html_diff | monthly | UCC PDF revision, portal restructuring, or caveat update | 2026-06-26 | active |
| watch:usa-nj:njac-5-23-3 | src:usa-nj:njac-5-23-3 | pdf_diff | monthly | Subcode adoption or amendment text changes | 2026-06-26 | active |
| watch:usa-nj:njac-5-23-1 | src:usa-nj:njac-5-23-1 | pdf_diff | quarterly | Date-rule, grace-period, or definitions changes | 2026-06-26 | active |
| watch:usa-nj:njac-5-23-4 | src:usa-nj:njac-5-23-4 | pdf_diff | quarterly | Enforcement agency or State-reserved enforcement model changes | 2026-06-26 | active |
| watch:usa-nj:dca-rule-proposals-adoptions | src:usa-nj:dca-rule-proposals-adoptions | html_diff | monthly | New UCC rule proposal or adoption, especially 2024 model-code cycle | 2026-06-26 | active |
| watch:usa-nj:dfs-codes | src:usa-nj:dca-dfs-codes-regs | html_diff | monthly | N.J.A.C. 5:70 final adoption or current fire-code page update | 2026-06-26 | active |
| watch:usa-nj:lexis-oal-admin-code | src:usa-nj:dca-ucc-portal | external_official_text_check | quarterly | Courtesy copy differs from official Administrative Code text | 2026-06-26 | pending |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-24 | Generated baseline draft report | report:usa-nj | none | Codex | Initial placeholder-heavy draft. |
| 2026-06-26 | Populated New Jersey report with source-backed UCC authority, current construction-code matrix, date rules, local enforcement model, caveats, and open issues | report:usa-nj; ahj:usa-nj:dca-commissioner-and-division-codes-standards; local-enforcement:usa-nj; local-amendment-rule:usa-nj; date-rule:usa-nj:001; date-rule:usa-nj:002 | src:usa-nj:ucc-act-dca-pdf; src:usa-nj:njac-5-23-1; src:usa-nj:njac-5-23-3; src:usa-nj:njac-5-23-4; src:usa-nj:dca-current-codes | ChatGPT | Status raised to partially_verified because core authority and current construction-code matrix are source-backed; fire primary text and full amendment digest remain open. |
