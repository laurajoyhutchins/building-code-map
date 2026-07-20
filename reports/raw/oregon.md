---
state:
  state_id: "US-OR"
  name: "Oregon"
  abbreviation: "OR"
report:
  report_id: "state-report:usa-or"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-26"
  last_verified: null
  reviewed_by: null
risk:
  overall_confidence: 0.74 # 0.00 - 1.00
  risk_flags:
    - "fire_enforcement_AHJ_mapping_not_completed"
    - "local_amendment_registry_not_found"
    - "plumbing_and_electrical_mandatory_dates_not_distinctly_identified"
    - "accessibility_treated_as_OSSC_scope_not_separate_codebook"
    - "jurisdiction_boundary_and_contact_data_not_started"
  open_questions_count: 7

---

# State Building Code Authority Report: Oregon

## 1. Executive Summary

- **Authority model:** Oregon uses a statewide building-code model. The Oregon Department of Consumer and Business Services, Building Codes Division (BCD), adopts, amends, interprets, and administers specialty codes that make up the Oregon State Building Code. The Department of the State Fire Marshal separately adopts the Oregon Fire Code for fire-prevention and operational fire-safety functions. Source IDs: `src:usa-or:bcd-state-code`, `src:usa-or:oar-918-008`, `src:usa-or:oar-837-040`, `src:usa-or:ors-455`, `src:usa-or:ors-476`.

- **Statewide code status:** Core statewide adoption records are now populated for the current structural/building, residential, mechanical/fuel-gas, plumbing, electrical, commercial energy, fire, and elevator code families. As of this report date, the 2025 OSSC, 2025 OMSC, 2025 OEESC, 2023 ORSC, 2023 OESC, 2023 OPSC, 2025 OFC, and 2024 Oregon Elevator Specialty Code are treated as current statewide codes within their documented scopes. Source IDs: `src:usa-or:bcd-adopted-codes`, `src:usa-or:oar-918-460-0010`, `src:usa-or:oar-918-440-0010`, `src:usa-or:oar-918-460-0500`, `src:usa-or:oar-918-480-0005`, `src:usa-or:oar-918-305`, `src:usa-or:oar-918-750-0110`, `src:usa-or:oar-837-040`, `src:usa-or:oar-918-400-0455`.

- **Local enforcement model:** Oregon separates statewide code adoption from local administration. Municipalities may administer and enforce full-service or partial building inspection programs under ORS chapter 455 and BCD delegation/renewal rules. BCD retains oversight, code-interpretation, alternate-method, program-review, and local-amendment approval functions. Source IDs: `src:usa-or:oar-918-020`, `src:usa-or:oar-918-008`, `src:usa-or:ors-455`.

- **Local amendment posture:** The Oregon State Building Code is treated as uniform statewide. Local ordinances, rules, or regulations covering the same matters may not impose different requirements unless authorized by the Director of the Department of Consumer and Business Services. Local amendment requests must be submitted to BCD with supporting public-process, impact, and fiscal information, and BCD may approve, condition, deny, review, or terminate the approval. Source IDs: `src:usa-or:ors-455`, `src:usa-or:oar-918-020-0370`.

- **Known transition periods or pending changes:** The 2025 OSSC and 2025 OMSC each had a six-month phase-in from 2025-10-01 through 2026-03-31. The 2025 OFC had a 180-day new-construction plan-review phase-in over the same period. The 2025 OEESC had a 2025-01-01 through 2025-07-01 phase-in. The 2023 ORSC construction provisions had a 2023-10-01 through 2024-03-31 phase-in. A 2026 Oregon Electrical Specialty Code adoption process is active with an anticipated adoption date of 2026-10-01. Source IDs: `src:usa-or:oar-918-460-0010`, `src:usa-or:oar-918-440-0010`, `src:usa-or:oar-837-040`, `src:usa-or:oar-918-460-0500`, `src:usa-or:oar-918-480-0005`, `src:usa-or:bcd-oesc-adoption`.

- **Production readiness:** partially_ready_for_narrow_validation. Core authority and statewide adoption fields are source-backed, but AHJ boundary/contact data, local fire-enforcement mapping, and local amendment registry coverage remain unresolved.

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: BCD adopts, amends, interprets, and administers specialty codes that make
    up the Oregon State Building Code.
  confidence: 0.9
  source_ids:
  - src:usa-or:bcd-state-code
  - src:usa-or:oar-918-008
  - src:usa-or:ors-455
- topic: Primary building code edition
  finding: The current commercial structural/building code is the 2025 Oregon Structural
    Specialty Code, effective 2025-10-01, based on the 2024 IBC, 2024 IFC new-construction
    provisions, and 2024 IEBC, with prior-code phase-in ending 2026-03-31.
  confidence: 0.93
  source_ids:
  - src:usa-or:oar-918-460-0010
  - src:usa-or:bcd-adopted-codes
- topic: Residential code edition
  finding: The current Oregon Residential Specialty Code is the 2023 ORSC, based on
    the 2021 IRC; Chapter 1 was effective and mandatory 2023-10-01 and construction
    provisions became mandatory 2024-04-01.
  confidence: 0.92
  source_ids:
  - src:usa-or:oar-918-480-0005
  - src:usa-or:bcd-residential
- topic: Electrical code authority and edition
  finding: BCD adopts the Oregon Electrical Specialty Code; the current OESC is effective
    2023-10-01 and consists of the 2023 NEC, 2023 NESC, and electrical provisions
    of the Oregon Elevator Specialty Code, as amended.
  confidence: 0.88
  source_ids:
  - src:usa-or:oar-918-305
  - src:usa-or:bcd-electrical
- topic: Fire code authority and edition
  finding: The Department of the State Fire Marshal adopts the Oregon Fire Code; the
    current OFC is the 2025 Oregon Fire Code based on the 2024 IFC, effective 2025-10-01,
    with 2022 OFC concurrency for specified phase-in purposes ending 2026-03-31.
  confidence: 0.88
  source_ids:
  - src:usa-or:oar-837-040
  - src:usa-or:osfm-fire-code
  - src:usa-or:ors-476
- topic: Local enforcement
  finding: Municipalities may administer full-service or partial building inspection
    programs through BCD program assumption/renewal; the exact AHJ for a given address
    remains a jurisdiction-resolution task.
  confidence: 0.78
  source_ids:
  - src:usa-or:oar-918-020
  - src:usa-or:ors-455
- topic: Local amendments
  finding: Local amendments to matters covered by the state building code require
    BCD/Director authorization and must follow the local-amendment request process.
  confidence: 0.88
  source_ids:
  - src:usa-or:oar-918-020-0370
  - src:usa-or:ors-455
- topic: Effective / operative date rule
  finding: For several BCD specialty codes, the code in effect at plan-review or permit-application
    filing controls unless the applicant agrees to later changes; the electrical rule
    separately uses the earlier of plan-review request or permit application, with
    applicant option to use the current Electrical Specialty Code.
  confidence: 0.82
  source_ids:
  - src:usa-or:oar-918-460-0010
  - src:usa-or:oar-918-460-0500
  - src:usa-or:oar-918-480-0005
  - src:usa-or:oar-918-305
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | `ahj:usa-or:dcbs-bcd` |
| Authority name | Oregon Department of Consumer and Business Services, Building Codes Division |
| Authority type | state agency / state code division |
| Legal basis | ORS chapter 455; OAR chapter 918, division 8; code-specific OAR chapter 918 divisions |
| Role | Adopts, amends, interprets, and administers specialty codes that make up the Oregon State Building Code; administers each code through specialized code programs; works with local building officials, advisory boards, and the public. |
| Enforcement model | Statewide code with municipal/local administration through full-service or partial building inspection programs, subject to BCD oversight and program rules. |
| Source IDs | `src:usa-or:bcd-state-code`, `src:usa-or:oar-918-008`, `src:usa-or:oar-918-020`, `src:usa-or:ors-455` |
| Verification status | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | `ahj:usa-or:dcbs-bcd` | Oregon Building Codes Division | Adopts OSSC and related structural specialty-code amendments. | OAR 918-460-0010; OAR chapter 918, division 8; ORS chapter 455 | `src:usa-or:oar-918-460-0010`, `src:usa-or:oar-918-008`, `src:usa-or:ors-455` | verified_core |
| Residential | `ahj:usa-or:dcbs-bcd` | Oregon Building Codes Division | Adopts ORSC / Low-Rise Residential Dwelling Code. | OAR 918-480-0005; ORS 455.610; ORS chapter 455 | `src:usa-or:oar-918-480-0005`, `src:usa-or:bcd-residential`, `src:usa-or:ors-455` | verified_core |
| Existing Building / Rehabilitation | `ahj:usa-or:dcbs-bcd` | Oregon Building Codes Division | Administers existing-building provisions through the OSSC adoption of the 2024 IEBC. | OAR 918-460-0010 | `src:usa-or:oar-918-460-0010` | verified_core |
| Mechanical | `ahj:usa-or:dcbs-bcd` | Oregon Building Codes Division | Adopts OMSC and mechanical specialty-code amendments. | OAR 918-440-0010; ORS chapter 455 | `src:usa-or:oar-918-440-0010`, `src:usa-or:ors-455` | verified_core |
| Plumbing | `ahj:usa-or:dcbs-bcd` | Oregon Building Codes Division | Adopts OPSC and plumbing specialty-code amendments. | OAR 918-750-0110; ORS 447.020; ORS chapter 455 | `src:usa-or:oar-918-750-0110`, `src:usa-or:bcd-adopted-codes`, `src:usa-or:ors-455` | verified_core |
| Fuel Gas | `ahj:usa-or:dcbs-bcd` | Oregon Building Codes Division | Adopts fuel-gas provisions through the OMSC based on the 2024 IFGC; ORSC includes low-rise residential fuel-gas/mechanical provisions. | OAR 918-440-0010; OAR 918-480-0005 | `src:usa-or:oar-918-440-0010`, `src:usa-or:oar-918-480-0005` | verified_core |
| Electrical | `ahj:usa-or:dcbs-bcd` | Oregon Building Codes Division | Adopts OESC and electrical amendments; program activity includes the Electrical and Elevator Board. | OAR 918-305-0100 and 918-305-0105; ORS chapter 479; ORS chapter 455 | `src:usa-or:oar-918-305`, `src:usa-or:bcd-electrical` | verified_core |
| Energy | `ahj:usa-or:dcbs-bcd` | Oregon Building Codes Division, working with Construction Industry Energy Board | Adopts OEESC as the energy provisions of the OSSC. | OAR 918-460-0500; ORS 455.505 and 455.511 | `src:usa-or:oar-918-460-0500`, `src:usa-or:bcd-energy` | verified_core |
| Fire - construction references | `ahj:usa-or:dcbs-bcd` | Oregon Building Codes Division | OSSC incorporates 2024 IFC new-construction provisions; construction plan-review interaction also references OFC phase-in rule. | OAR 918-460-0010; OAR 837-040-0010 | `src:usa-or:oar-918-460-0010`, `src:usa-or:oar-837-040` | verified_core |
| Fire - operational / prevention code | `ahj:usa-or:state-fire-marshal` | Oregon Department of the State Fire Marshal | Adopts Oregon Fire Code; fire-prevention and operational fire-safety rulemaking. | ORS 476.030; OAR 837-040-0010 | `src:usa-or:oar-837-040`, `src:usa-or:osfm-fire-code`, `src:usa-or:ors-476` | partially_verified |
| Accessibility | `ahj:usa-or:dcbs-bcd` | Oregon Building Codes Division | Accessibility is treated as OSSC/related specialty-code scope for this pass, rather than a separately adopted statewide accessibility codebook. | OAR 918-460-0010; OAR 918-305-0030 cross-references ADA and OSSC Chapter 11 for electrical impacts. | `src:usa-or:oar-918-460-0010`, `src:usa-or:oar-918-305` | partially_verified |
| Elevator / Conveyance | `ahj:usa-or:dcbs-bcd` | Oregon Building Codes Division | Adopts Oregon Elevator Specialty Code and elevator amendments. | OAR 918-400-0455; OAR 918-400-0458; ORS chapters 455, 460, and 479 | `src:usa-or:oar-918-400-0455`, `src:usa-or:bcd-elevator` | verified_core |

### 2.3 Authority Hierarchy Notes

The Oregon model is hybrid but statewide: BCD creates and maintains a uniform state building code through Oregon specialty codes, while municipalities can administer and enforce building inspection programs after program assumption/renewal. OAR 918-008 also identifies statewide code interpretations and alternate method rulings that jurisdictions administering the state building code must honor. Fire-prevention code adoption is separate: the Department of the State Fire Marshal adopts the Oregon Fire Code under ORS 476.030 and OAR chapter 837, division 40.

Local enforcement authority and local amendment authority are distinct. A municipality may administer/enforce a building inspection program without having independent authority to change statewide code requirements. Local amendments to matters covered by the state building code require authorization through the BCD local-amendment request process.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| `edge:usa-or:001` | `ahj:usa-or:dcbs-bcd` | adopts_and_amends | Oregon specialty codes / Oregon State Building Code | `src:usa-or:bcd-state-code`, `src:usa-or:oar-918-008` | partially_verified |
| `edge:usa-or:002` | `ahj:usa-or:dcbs-bcd` | administers_through | specialized code programs and advisory-board processes | `src:usa-or:bcd-state-code`, `src:usa-or:oar-918-008` | partially_verified |
| `edge:usa-or:003` | `ahj:usa-or:dcbs-bcd` | authorizes_or_denies | local amendment requests under ORS 455.040 | `src:usa-or:oar-918-020-0370`, `src:usa-or:ors-455` | verified_core |
| `edge:usa-or:004` | `ahj:usa-or:dcbs-bcd` | oversees | municipal full-service and partial building inspection programs | `src:usa-or:oar-918-020`, `src:usa-or:ors-455` | partially_verified |
| `edge:usa-or:005` | Oregon State Building Code | preempts | local ordinances, rules, and regulations covering the same matters unless authorized by the DCBS Director | `src:usa-or:ors-455`, `src:usa-or:oar-918-020-0370` | partially_verified |
| `edge:usa-or:006` | `ahj:usa-or:state-fire-marshal` | adopts | Oregon Fire Code | `src:usa-or:oar-837-040`, `src:usa-or:ors-476` | partially_verified |
| `edge:usa-or:007` | `ahj:usa-or:state-fire-marshal` | coordinates_with | local building official during 2025 OFC new-construction plan-review phase-in | `src:usa-or:oar-837-040` | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | 2025 Oregon Structural Specialty Code (OSSC) | 2024 IBC; 2024 IFC new-construction provisions; 2024 IEBC; energy provisions by OAR 918-460-0500 | 2025 | current_mandatory | 2025-09-26 | 2025-10-01 | null | 2026-04-01 | Chapter 1 effective with no phase-in on 2025-10-01; 2022 OSSC allowed except Chapter 1 from 2025-10-01 through 2026-03-31; code in effect at plan-review or permit-application filing controls unless applicant agrees to later changes. | `src:usa-or:oar-918-460-0010`, `src:usa-or:bcd-adopted-codes` |
| Residential | 2023 Oregon Residential Specialty Code (ORSC) | 2021 IRC; low-rise plumbing provisions in 2023 OPSC; low-rise electrical provisions in 2023 OESC | 2023 | current_mandatory | 2023-09-12 | 2023-10-01 | null | 2024-04-01 | Chapter 1 effective and mandatory 2023-10-01; 2021 ORSC allowed except Chapter 1 from 2023-10-01 through 2024-03-31; all building departments required to accept 2021 or 2023 ORSC designs during phase-in. | `src:usa-or:oar-918-480-0005`, `src:usa-or:bcd-residential` |
| Existing Building / Rehabilitation | 2025 Oregon Structural Specialty Code existing-building provisions | 2024 IEBC as part of OSSC | 2025 | current_mandatory | 2025-09-26 | 2025-10-01 | null | 2026-04-01 | Same OSSC phase-in as Building row; Chapter 34 interim amendments effective 2026-04-01 for select snow-load repair/alteration provisions. | `src:usa-or:oar-918-460-0010`, `src:usa-or:oar-918-460-0015` |
| Mechanical | 2025 Oregon Mechanical Specialty Code (OMSC) | 2024 IMC and 2024 IFGC | 2025 | current_mandatory | 2025-09-26 | 2025-10-01 | null | 2026-04-01 | 2022 OMSC allowed from 2025-10-01 through 2026-03-31; 2022 or 2025 OMSC permitted during phase-in. | `src:usa-or:oar-918-440-0010`, `src:usa-or:bcd-adopted-codes` |
| Plumbing | 2023 Oregon Plumbing Specialty Code (OPSC) | 2021 UPC, First Printing, chapters 1-11 and 13-17, selected appendices, as amended; BCD portal also describes relationship to 2021 OPSC | 2023 | current | 2023-09-21 | 2023-10-01 | null | null | No distinct phase-in or mandatory date was extracted from the current OAR text; local jurisdictions may adopt UPC Appendices L and N. | `src:usa-or:oar-918-750-0110`, `src:usa-or:bcd-adopted-codes` |
| Fuel Gas | 2025 Oregon Mechanical Specialty Code fuel-gas provisions | 2024 IFGC, with Oregon amendments; low-rise residential fuel-gas/mechanical provisions also interact with ORSC | 2025 | current_mandatory | 2025-09-26 | 2025-10-01 | null | 2026-04-01 | Same OMSC phase-in as Mechanical row. | `src:usa-or:oar-918-440-0010`, `src:usa-or:oar-918-480-0005` |
| Electrical | 2023 Oregon Electrical Specialty Code (OESC) | 2023 NFPA 70 NEC; 2023 IEEE C2 NESC; electrical provisions of Oregon Elevator Specialty Code, as amended | 2023 | current | 2023-09-19 | 2023-10-01 | null | null | Applicable code is the earlier of plan-review request or permit application; applicant may choose the current Electrical Specialty Code. No separate mandatory date was extracted. | `src:usa-or:oar-918-305`, `src:usa-or:bcd-electrical` |
| Energy | 2025 Oregon Energy Efficiency Specialty Code (OEESC), Chapter 13 energy provisions of OSSC | ANSI/ASHRAE/IES Standard 90.1-2022 with Oregon amendments | 2025 | current_mandatory | 2024-12-23 | 2025-01-01 | null | 2025-07-01 | 2021 OEESC or 2025 OEESC permitted from 2025-01-01 through 2025-07-01; code in effect at plan-review or permit-application filing controls unless applicant agrees to later changes. | `src:usa-or:oar-918-460-0500`, `src:usa-or:bcd-energy` |
| Fire - construction references | 2025 OSSC fire/life-safety construction provisions | 2024 IFC new-construction provisions incorporated into OSSC; OFC plan-review phase-in also applies to new-construction plan review | 2025 | current_mandatory | 2025-09-26 | 2025-10-01 | null | 2026-04-01 | OSSC phase-in for construction code; OFC new-construction plan-review phase-in from 2025-10-01 through 2026-03-31, reviewed to 2025 or 2022 OFC as directed by the local building official. | `src:usa-or:oar-918-460-0010`, `src:usa-or:oar-837-040` |
| Fire - operational / prevention code | 2025 Oregon Fire Code (OFC) | 2024 IFC as amended by Department of the State Fire Marshal | 2025 | current | 2025-09-23 | 2025-10-01 | null | 2026-04-01 | 2022 OFC remained adopted for specified 180-day phase-in purposes from 2025-10-01 through 2026-03-31; for new-construction plan review, 2025 or 2022 OFC applied as directed by local building official during phase-in. | `src:usa-or:oar-837-040`, `src:usa-or:osfm-fire-code` |
| Accessibility | Accessibility provisions within 2025 OSSC and related specialty-code cross-references | OSSC accessibility scope, including Chapter 11 references; OESC rule cross-references ADA, Fair Housing Act, and OSSC Chapter 11 for electrical impacts | 2025 / current specialty-code cycle | partially_verified | 2025-09-26 | 2025-10-01 | null | 2026-04-01 | Treated as part of current OSSC/specialty-code framework; no separate standalone Oregon accessibility codebook was verified in this pass. | `src:usa-or:oar-918-460-0010`, `src:usa-or:oar-918-305` |
| Elevator / Conveyance | 2024 Oregon Elevator Specialty Code | ASME A17.1-2019; A17.2-2020; A17.3-2020; A17.6-2022; A18.1-2020; A90.1-2015; A17.7-2007 (R2012), as amended | 2024 | current | 2023-12-29 | 2024-01-01 | null | 2024-07-02 | ASME A17.3 existing-elevator standard had a six-month phase-in from 2024-01-01 through 2024-07-01; beginning 2024-07-02, existing installations are inspected under A17.3-2020 as adopted in the 2024 code. | `src:usa-or:oar-918-400-0455`, `src:usa-or:bcd-elevator` |

### 3.2 Adoption Records

#### `adoption-record:usa-or:ossc-2025`

| Field | Value |
| --- | --- |
| Code family | Building; Existing Building / Rehabilitation; Fire - construction references |
| Authority ID | `ahj:usa-or:dcbs-bcd` |
| State code | 2025 Oregon Structural Specialty Code |
| Base model code | 2024 IBC; 2024 IFC new-construction provisions; 2024 IEBC |
| Adoption action date | 2025-09-26 |
| Effective date | 2025-10-01 |
| Operative date | null |
| Mandatory date | 2026-04-01 |
| Transition | Chapter 1 effective without phase-in on 2025-10-01; 2022 OSSC, except Chapter 1, allowed from 2025-10-01 through 2026-03-31; after that the 2025 OSSC is the only current OSSC edition identified. |
| Date-rule hook | `date-rule:usa-or:bcd-application-controls` |
| Source IDs | `src:usa-or:oar-918-460-0010`, `src:usa-or:bcd-adopted-codes` |
| Confidence | 0.93 |

#### `adoption-record:usa-or:orsc-2023`

| Field | Value |
| --- | --- |
| Code family | Residential |
| Authority ID | `ahj:usa-or:dcbs-bcd` |
| State code | 2023 Oregon Residential Specialty Code |
| Base model code | 2021 IRC, with Oregon amendments and exclusions; does not contain separate electrical or plumbing provisions in the integrated ICC codebook |
| Adoption action date | 2023-09-12 |
| Effective date | 2023-10-01 |
| Operative date | null |
| Mandatory date | 2024-04-01 for construction provisions; 2023-10-01 for Chapter 1 |
| Transition | Chapter 1 had no phase-in; 2021 ORSC, except Chapter 1, allowed from 2023-10-01 through 2024-03-31; all building departments required to accept 2021 or 2023 ORSC designs during phase-in. |
| Date-rule hook | `date-rule:usa-or:bcd-application-controls` |
| Source IDs | `src:usa-or:oar-918-480-0005`, `src:usa-or:bcd-residential` |
| Confidence | 0.92 |

#### `adoption-record:usa-or:omsc-2025`

| Field | Value |
| --- | --- |
| Code family | Mechanical; Fuel Gas |
| Authority ID | `ahj:usa-or:dcbs-bcd` |
| State code | 2025 Oregon Mechanical Specialty Code |
| Base model code | 2024 IMC and 2024 IFGC |
| Adoption action date | 2025-09-26 |
| Effective date | 2025-10-01 |
| Operative date | null |
| Mandatory date | 2026-04-01 |
| Transition | 2022 OMSC adopted for phase-in from 2025-10-01 through 2026-03-31; 2022 or 2025 OMSC permitted during that period. |
| Date-rule hook | `date-rule:usa-or:omsc-phase-in` |
| Source IDs | `src:usa-or:oar-918-440-0010`, `src:usa-or:bcd-adopted-codes` |
| Confidence | 0.91 |

#### `adoption-record:usa-or:opsc-2023`

| Field | Value |
| --- | --- |
| Code family | Plumbing |
| Authority ID | `ahj:usa-or:dcbs-bcd` |
| State code | 2023 Oregon Plumbing Specialty Code |
| Base model code | 2021 Uniform Plumbing Code, First Printing, chapters 1-11 and 13-17, appendices A, B, C, D, E, K, and M, as amended; BCD portal also identifies 2021 OPSC relationship |
| Adoption action date | 2023-09-21 |
| Effective date | 2023-10-01 |
| Operative date | null |
| Mandatory date | null |
| Transition | No distinct phase-in or mandatory date was extracted from current OAR text. |
| Date-rule hook | `date-rule:usa-or:opsc-effective-only` |
| Source IDs | `src:usa-or:oar-918-750-0110`, `src:usa-or:bcd-adopted-codes` |
| Confidence | 0.80 |

#### `adoption-record:usa-or:oesc-2023`

| Field | Value |
| --- | --- |
| Code family | Electrical |
| Authority ID | `ahj:usa-or:dcbs-bcd` |
| State code | 2023 Oregon Electrical Specialty Code |
| Base model code | 2023 NFPA 70 NEC; 2023 IEEE C2 NESC; electrical provisions of Oregon Elevator Specialty Code, as amended |
| Adoption action date | 2023-09-19 |
| Effective date | 2023-10-01 |
| Operative date | null |
| Mandatory date | null |
| Transition | Applicable code is the Electrical Specialty Code in effect at the earlier of a plan-review request or permit application; applicant may choose the current Electrical Specialty Code. |
| Date-rule hook | `date-rule:usa-or:oesc-governing-code` |
| Source IDs | `src:usa-or:oar-918-305`, `src:usa-or:bcd-electrical` |
| Confidence | 0.86 |

#### `adoption-record:usa-or:oeesc-2025`

| Field | Value |
| --- | --- |
| Code family | Energy |
| Authority ID | `ahj:usa-or:dcbs-bcd` |
| State code | 2025 Oregon Energy Efficiency Specialty Code |
| Base model code | ANSI/ASHRAE/IES Standard 90.1-2022 with Oregon amendments |
| Adoption action date | 2024-12-23 |
| Effective date | 2025-01-01 |
| Operative date | null |
| Mandatory date | 2025-07-01 |
| Transition | 2021 OEESC or 2025 OEESC permitted from 2025-01-01 through 2025-07-01; after the phase-in, the 2025 OEESC is mandatory for the covered commercial energy scope. |
| Date-rule hook | `date-rule:usa-or:bcd-application-controls` |
| Source IDs | `src:usa-or:oar-918-460-0500`, `src:usa-or:bcd-energy` |
| Confidence | 0.90 |

#### `adoption-record:usa-or:ofc-2025`

| Field | Value |
| --- | --- |
| Code family | Fire - operational / prevention code; Fire - construction plan-review interaction |
| Authority ID | `ahj:usa-or:state-fire-marshal` |
| State code | 2025 Oregon Fire Code |
| Base model code | 2024 International Fire Code, as amended by the Department of the State Fire Marshal |
| Adoption action date | 2025-09-23 |
| Effective date | 2025-10-01 |
| Operative date | null |
| Mandatory date | 2026-04-01 for post-phase-in use inferred from expiration of 2022 OFC concurrency; treat as rule-derived rather than separately stated |
| Transition | 2022 OFC was adopted for a 180-day phase-in from 2025-10-01 through 2026-03-31 for new-construction plan review; review could be to 2025 or 2022 OFC as directed by the local building official. 2022 OFC also remained adopted for RMS integration over the same period. |
| Date-rule hook | `date-rule:usa-or:ofc-phase-in` |
| Source IDs | `src:usa-or:oar-837-040`, `src:usa-or:osfm-fire-code` |
| Confidence | 0.86 |

#### `adoption-record:usa-or:elevator-2024`

| Field | Value |
| --- | --- |
| Code family | Elevator / Conveyance |
| Authority ID | `ahj:usa-or:dcbs-bcd` |
| State code | 2024 Oregon Elevator Specialty Code |
| Base model code | ASME A17.1-2019; A17.2-2020; A17.3-2020; A17.6-2022; A18.1-2020; A90.1-2015; A17.7-2007 (R2012), as amended |
| Adoption action date | 2023-12-29 |
| Effective date | 2024-01-01 |
| Operative date | null |
| Mandatory date | 2024-07-02 for A17.3-2020 existing-installation inspection use |
| Transition | A17.3-2002 or A17.3-2020 permitted from 2024-01-01 through 2024-07-01; beginning 2024-07-02, existing installations are inspected using A17.3-2020 as adopted in the 2024 Oregon Elevator Specialty Code. |
| Date-rule hook | `date-rule:usa-or:elevator-a17-3-phase-in` |
| Source IDs | `src:usa-or:oar-918-400-0455`, `src:usa-or:bcd-elevator` |
| Confidence | 0.88 |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

Oregon uses code-specific phase-in rules and filing-trigger rules rather than one single date rule for every specialty code. The recurring BCD construction-code rule is that code requirements in effect when a plan-review or permit application is filed control the construction under that application unless the applicant agrees to be controlled by later changes. The electrical rule is separately stated: the applicable Electrical Specialty Code is the code in effect at the earlier of a plan-review request or permit application, or the current Electrical Specialty Code at the applicant's option. Fire-code phase-in logic is stated in OAR 837-040 for the OFC and is not identical to the BCD specialty-code text.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `date-rule:usa-or:bcd-application-controls` | OSSC; ORSC; OEESC and related BCD specialty-code applications where rule text repeats this provision | permit_or_plan_review_application_controls | Applies on filing date | Code requirements in effect at the time a plan-review or permit application is filed control construction unless the applicant agrees to subsequent changes. | Yes, if a phase-in rule separately keeps prior code available or if applicant agrees to later changes. | `src:usa-or:oar-918-460-0010`, `src:usa-or:oar-918-460-0500`, `src:usa-or:oar-918-480-0005` | verified_core |
| `date-rule:usa-or:ossc-phase-in` | 2025 OSSC | phase_in | 2025-10-01 through 2026-03-31 | Chapter 1 effective immediately; 2022 OSSC except Chapter 1 remained available during phase-in. | Yes, 2022 OSSC except Chapter 1 through 2026-03-31. | `src:usa-or:oar-918-460-0010`, `src:usa-or:bcd-adopted-codes` | verified_core |
| `date-rule:usa-or:omsc-phase-in` | 2025 OMSC | phase_in | 2025-10-01 through 2026-03-31 | 2022 OMSC and 2025 OMSC both permitted during phase-in. | Yes, 2022 OMSC through 2026-03-31. | `src:usa-or:oar-918-440-0010`, `src:usa-or:bcd-adopted-codes` | verified_core |
| `date-rule:usa-or:orsc-phase-in` | 2023 ORSC | phase_in | 2023-10-01 through 2024-03-31 | Chapter 1 effective immediately; all building departments had to accept plans designed to either 2023 ORSC or 2021 ORSC during phase-in. | Yes, 2021 ORSC except Chapter 1 through 2024-03-31. | `src:usa-or:oar-918-480-0005`, `src:usa-or:bcd-residential` | verified_core |
| `date-rule:usa-or:oeesc-phase-in` | 2025 OEESC | phase_in | 2025-01-01 through 2025-07-01 | 2021 OEESC or 2025 OEESC permitted as Chapter 13 of the OSSC during phase-in. | Yes, 2021 OEESC through 2025-07-01. | `src:usa-or:oar-918-460-0500`, `src:usa-or:bcd-energy` | verified_core |
| `date-rule:usa-or:oesc-governing-code` | OESC | plan_review_or_permit_earlier_controls | Ongoing | Applicable Electrical Specialty Code is in effect at the earlier of request for plan review or permit application; applicant may choose the current Electrical Specialty Code. | Prior-code allowance not separately extracted; applicant current-code option is explicit. | `src:usa-or:oar-918-305` | verified_core |
| `date-rule:usa-or:ofc-phase-in` | 2025 OFC new-construction plan review and RMS integration | phase_in | 2025-10-01 through 2026-03-31 | 2022 OFC remained adopted for specified 180-day phase-in; new construction plan reviews reviewed to 2025 or 2022 OFC as directed by local building official. | Yes, 2022 OFC for specified phase-in purposes through 2026-03-31. | `src:usa-or:oar-837-040`, `src:usa-or:osfm-fire-code` | verified_core |
| `date-rule:usa-or:elevator-a17-3-phase-in` | Existing elevators and escalators under ASME A17.3 | phase_in | 2024-01-01 through 2024-07-01 | Equipment owner could choose A17.3-2002 or A17.3-2020 during phase-in; state required A17.3-2020 inspection beginning 2024-07-02. | Yes, A17.3-2002 through 2024-07-01. | `src:usa-or:oar-918-400-0455`, `src:usa-or:bcd-elevator` | verified_core |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Electrical | 2026 Oregon Electrical Specialty Code | null | 2026-10-01 | null | null | null | active_watch | `src:usa-or:bcd-oesc-adoption` | BCD page reports an active 2026 OESC adoption process with anticipated adoption date 2026-10-01 and base model 2026 NFPA 70 NEC. |
| Structural / existing building | OSSC Chapter 34 snow-load repair/alteration amendments | null | null | 2026-04-01 | null | 2026-04-01 | implemented_monitor | `src:usa-or:oar-918-460-0015` | Already effective on the report date; included here for monitoring because it is an interim/high-impact amendment to an existing current code. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| `applicability-rule:usa-or:orsc-low-rise` | Residential | one- and two-family dwellings, townhouses, and related low-rise residential scope | ORSC scope and Low-Rise Residential Dwelling Code | ORSC is the Oregon low-rise residential code; low-rise plumbing provisions are adopted in the 2023 OPSC and low-rise electrical provisions are adopted in the 2023 OESC. | `src:usa-or:oar-918-480-0005`, `src:usa-or:bcd-residential` | verified_core |
| `applicability-rule:usa-or:opsc-local-appendices` | Plumbing | UPC Appendices L and N | local adoption | OAR 918-750-0110 does not adopt Appendices L and N as part of OPSC but makes them available for local jurisdiction adoption. | `src:usa-or:oar-918-750-0110` | verified_core |
| `applicability-rule:usa-or:ofc-new-construction-phase-in` | Fire | new-construction plan review | 2025 OFC phase-in | During the 180-day OFC phase-in, new-construction plan reviews could be reviewed to 2025 OFC or 2022 OFC as directed by the local building official. | `src:usa-or:oar-837-040` | verified_core |
| `applicability-rule:usa-or:oesc-existing-installations` | Electrical | existing electrical installations | existing wiring complied with code in effect at installation | Existing wiring that complied with the minimum electrical safety code standards in effect at installation is not considered in violation of current OESC unless use or occupancy changes requiring different methods, alterations, or additions. | `src:usa-or:oar-918-305` | verified_core |
| `applicability-rule:usa-or:elevator-existing-a17-3` | Elevator / Conveyance | existing elevators and escalators | inspection after A17.3 phase-in | Beginning 2024-07-02, the state is required to inspect existing elevator/escalator installations using A17.3-2020 as adopted in the 2024 Oregon Elevator Specialty Code. | `src:usa-or:oar-918-400-0455` | verified_core |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** Oregon specialty codes are base model codes and referenced standards adopted by administrative rule with Oregon-specific amendments. BCD's general code-development rule says model codes, standards, and other publications are adopted by reference through administrative rule to create the state building code; code-specific amendment rules place Oregon amendments in the relevant OAR division, tables, or BCD-published amendment documents. The Department of the State Fire Marshal adopts and may amend the Oregon Fire Code under OAR chapter 837, division 40.

**Where amendments are published:** Oregon Administrative Rules, BCD code-adoption pages, code-specific amendment tables, integrated codebooks where available, and State Fire Marshal fire-code rulemaking documents.

**Amendment parsing status:** high_level_only. Amendment existence and several high-impact amendments are recorded, but this report has not extracted full amendment text or built section-level amendment diffs.

Source IDs: `src:usa-or:oar-918-008`, `src:usa-or:oar-918-460-0015`, `src:usa-or:oar-918-480-0010`, `src:usa-or:oar-918-305`, `src:usa-or:oar-918-400-0455`, `src:usa-or:oar-837-040`.

### 5.2 State Amendment Sources

| Amendment Source ID | Code Family | Publication Path | Scope Captured | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| `amend-src:usa-or:ossc-2025` | Building / Existing Building | OAR 918-460-0015 and BCD OSSC adoption page | High-level amendment rule identified; snow-load repairs/alterations amendment captured. | `src:usa-or:oar-918-460-0015`, `src:usa-or:bcd-adopted-codes` | partially_parsed |
| `amend-src:usa-or:orsc-2023` | Residential | OAR 918-480-0010 and BCD residential page | R310, R302.3, R327, and R332 interim amendment topics and dates captured. | `src:usa-or:oar-918-480-0010`, `src:usa-or:bcd-residential` | partially_parsed |
| `amend-src:usa-or:oesc-2023` | Electrical | OAR 918-305-0105 / Table 1-E and BCD electrical page | Amendment table path identified; 2025 errata and 2026 microgrid amendment captured. | `src:usa-or:oar-918-305`, `src:usa-or:bcd-electrical` | partially_parsed |
| `amend-src:usa-or:opsc-2023` | Plumbing | OAR 918-750-0110 and OAR 918-750-0115 | Base adoption and appendices captured; full amendments not parsed. | `src:usa-or:oar-918-750-0110` | partial_path_only |
| `amend-src:usa-or:elevator-2024` | Elevator / Conveyance | OAR 918-400-0458 / Table 2-L | Amendment path identified; full Table 2-L not parsed. | `src:usa-or:oar-918-400-0455`, `src:usa-or:bcd-elevator` | partial_path_only |
| `amend-src:usa-or:ofc-2025` | Fire | OAR 837-040 and OSFM fire-code page | OFC adoption and phase-in captured; full OFC amendments not parsed. | `src:usa-or:oar-837-040`, `src:usa-or:osfm-fire-code` | partially_parsed |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Effective Date | Amendment Summary | Source IDs | Parsing Status |
| --- | --- | --- | --- | --- | --- |
| `amend:usa-or:ossc-2026-snow-load-repair-alteration` | Existing Building / Rehabilitation | 2026-04-01 | OSSC Chapter 34 sections are amended to provide an exception from new-construction provisions for certain repairs and alterations in select locations where new ground snow-load provisions apply. | `src:usa-or:oar-918-460-0015` | summary_only |
| `amend:usa-or:orsc-r327-wildfire-local-adoption` | Residential | 2025-08-05 | ORSC Section R327 is amended for additional wildfire hazard mitigation provisions available for local adoption. | `src:usa-or:oar-918-480-0010`, `src:usa-or:bcd-residential` | summary_only |
| `amend:usa-or:orsc-r332-sleeping-lofts` | Residential | 2026-01-01 | ORSC Section R332 is amended to add sleeping-loft construction provisions and ladder access/egress provisions. | `src:usa-or:oar-918-480-0010`, `src:usa-or:bcd-residential` | summary_only |
| `amend:usa-or:oesc-microgrid-equipment` | Electrical | 2026-01-01 | OESC is amended to require customer-side microgrid equipment and distribution to be installed under applicable current-code sections. | `src:usa-or:oar-918-305` | summary_only |
| `amend:usa-or:elevator-table-2l` | Elevator / Conveyance | 2024-01-01 | OAR 918-400-0458 adopts Oregon-specific amendments to ASME A17.1-2019 in Table 2-L. | `src:usa-or:oar-918-400-0455` | path_only |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-or"
  model: "statewide_code_with_municipal_program_administration"
  enforcing_entities:
    - "municipalities administering full-service building inspection programs"
    - "municipalities administering partial building inspection programs"
    - "Oregon Building Codes Division for state-administered or non-assumed scopes"
    - "local building officials for local program decisions and phase-in direction where assigned"
    - "fire officials / fire jurisdictions for Oregon Fire Code enforcement; address-level mapping unresolved"
  required_officials:
    - "building official or contract building official for municipal building inspection programs"
    - "certified specialty-code inspectors and plan reviewers where required by ORS/OAR"
    - "fire code official or fire jurisdiction authority for OFC enforcement; precise structure unresolved"
  state_reserved_activities:
    - "statewide specialty-code adoption and amendment"
    - "statewide code interpretations"
    - "alternate method rulings"
    - "local amendment approval, conditioning, review, and termination"
    - "program assumption, renewal, and oversight"
    - "elevator existing-installation inspection standard after A17.3 phase-in"
  source_ids:
    - "src:usa-or:oar-918-020"
    - "src:usa-or:oar-918-008"
    - "src:usa-or:oar-918-020-0370"
    - "src:usa-or:oar-918-400-0455"
    - "src:usa-or:oar-837-040"
    - "src:usa-or:ors-455"
  verification_status: "partially_verified"
  confidence: 0.78
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-or"
  model: "statewide_uniform_code_with_director_BCD_authorization_required_for_local_differences"
  applies_to_code_families:
    - "matters covered under the Oregon State Building Code"
    - "plumbing Appendices L and N may be adopted locally under OAR 918-750-0110"
    - "residential wildfire hazard mitigation provisions may be locally adopted where authorized by ORSC amendment path"
  approval_required: true
  approving_authority_id: "ahj:usa-or:dcbs-bcd"
  filing_required: true
  registry_exists: null
  registry_source_ids: []
  legal_basis_source_ids:
    - "src:usa-or:ors-455"
    - "src:usa-or:oar-918-020-0370"
    - "src:usa-or:oar-918-750-0110"
    - "src:usa-or:oar-918-480-0010"
  verification_status: "partially_verified"
  confidence: 0.86
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement is a delegation/administration question: whether a city, county, or other municipality is authorized and staffed to administer a building inspection program for the relevant specialty-code scope. Local amendment authority is a separate preemption/authorization question: Oregon's state building code is uniform statewide, and local code differences for matters covered by the building code require DCBS/BCD authorization. A local jurisdiction may therefore be the enforcement AHJ for a project without having independent power to adopt conflicting local building-code requirements.

### 6.4 Known Local Amendment Registries

| Registry ID | Registry Name | Coverage | Source IDs | Status | Notes |
| --- | --- | --- | --- | --- | --- |
| `registry:usa-or:local-building-code-amendments` | Oregon local building-code amendment approvals | statewide if maintained | none | unresolved | No statewide public registry of all approved local building-code amendments was located in this pass. BCD local-amendment request rule is verified, but registry existence is not. |

### 6.5 Municipality-Specific Known Amendments

| Locality | Code Family | Amendment Topic | Approval Path | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| statewide local-option | Plumbing | UPC Appendices L and N are not part of OPSC but are available for local jurisdiction adoption. | OAR 918-750-0110 | `src:usa-or:oar-918-750-0110` | verified_core |
| local-option municipalities | Residential | ORSC R327 wildfire hazard mitigation provisions are described as available for local adoption. | OAR 918-480-0010 / BCD residential amendment page | `src:usa-or:oar-918-480-0010`, `src:usa-or:bcd-residential` | partially_verified |
| specific municipalities | multiple | locality-specific approved amendments | unresolved | none | not_started |

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: not_started

Jurisdiction stack:

```yaml
Address
  -> State of Oregon
  -> County
  -> Municipality / unincorporated county
  -> BCD program-assumption and renewal status by specialty-code program
  -> Building inspection program AHJ
  -> Fire jurisdiction / fire code official
  -> Trade-specific inspection programs, including electrical, plumbing, mechanical, elevator, and manufactured dwelling scopes
  -> Applicable state specialty-code adoption records
  -> Approved local amendments or local-option provisions
  -> Statewide code interpretations and alternate method rulings
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | Oregon state boundary dataset | none | statewide | unresolved | pending |
| County | Oregon county boundary dataset | none | statewide | unresolved | pending |
| Municipality | Oregon city/municipal boundary dataset | none | statewide | unresolved | pending |
| Fire District | Oregon fire district / fire response boundary dataset | none | statewide | unresolved | pending |
| Building Inspection Program | BCD municipal program administration/assumption records | `src:usa-or:oar-918-020` | statewide program framework, not GIS | renewal/reporting cycle varies by group | framework_only |
| Local Amendments | BCD local amendment approvals | `src:usa-or:oar-918-020-0370` | process verified; registry unresolved | unresolved | process_only |

### 7.3 AHJ Contact Data

No AHJ contact data was populated for this pass. Future AHJ work should map BCD program-assumption/renewal records, municipal building departments, county programs, fire jurisdictions, and state-administered scopes.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Publisher / Custodian | URL | Accessed | Key Fields Supported |
| --- | --- | --- | --- | --- | --- |
| `src:usa-or:bcd-state-code` | Oregon State Building Code: Codes and standards | Oregon Building Codes Division | https://www.oregon.gov/bcd/codes-stand/pages/index.aspx | 2026-06-26 | BCD role; specialty-code programs; adoption/amendment/interpretation overview |
| `src:usa-or:bcd-adopted-codes` | Adopted codes online | Oregon Building Codes Division | https://www.oregon.gov/bcd/codes-stand/pages/adopted-codes.aspx | 2026-06-26 | Statewide code list; effective/mandatory dates; model-code bases; portal caveats |
| `src:usa-or:bcd-residential` | Residential Structures Code Program | Oregon Building Codes Division | https://www.oregon.gov/bcd/codes-stand/pages/residential-structures.aspx | 2026-06-26 | ORSC scope; ORSC phase-in and mandatory dates; 2021 IRC base; interim amendments list |
| `src:usa-or:bcd-electrical` | Electrical Code Program | Oregon Building Codes Division | https://www.oregon.gov/bcd/codes-stand/pages/electrical.aspx | 2026-06-26 | OESC program role; 2023 OESC effective date and NEC base; amendment table path |
| `src:usa-or:bcd-energy` | Oregon Energy Efficiency Specialty Code | Oregon Building Codes Division | https://www.oregon.gov/bcd/codes-stand/pages/oeesc-adoption.aspx | 2026-06-26 | 2025 OEESC effective/mandatory dates; ASHRAE 90.1-2022 base; CIEB participation |
| `src:usa-or:bcd-elevator` | Oregon Elevator Specialty Code | Oregon Building Codes Division | https://www.oregon.gov/bcd/codes-stand/pages/elevator-adoption.aspx | 2026-06-26 | 2024 elevator code components; effective date; A17.3 phase-in |
| `src:usa-or:bcd-oesc-adoption` | Oregon Electrical Specialty Code adoption | Oregon Building Codes Division | https://www.oregon.gov/bcd/codes-stand/pages/oesc-adoption.aspx | 2026-06-26 | Pending 2026 OESC adoption process; anticipated 2026-10-01 adoption; 2026 NEC base |
| `src:usa-or:oar-918-008` | OAR chapter 918, division 8: Division code development rules, in general | Oregon Secretary of State / Building Codes Division | https://secure.sos.state.or.us/oard/displayDivisionRules.action?selectedDivision=4129 | 2026-06-26 | State building code creation by administrative rule; adoption/amendment process; statewide interpretations and alternate method rulings |
| `src:usa-or:oar-918-020` | OAR chapter 918, division 20: Delegation of programs to local jurisdictions | Oregon Secretary of State / Building Codes Division | https://secure.sos.state.or.us/oard/displayDivisionRules.action?selectedDivision=4131 | 2026-06-26 | Municipal program assumption, renewal, reporting periods, program oversight |
| `src:usa-or:oar-918-020-0370` | OAR 918-020-0370: Local Amendment Requests | Oregon Secretary of State / Building Codes Division | https://secure.sos.state.or.us/oard/view.action?ruleNumber=918-020-0370 | 2026-06-26 | Local amendment request requirements, BCD review/approval/denial/termination, public-process and impact-report requirements |
| `src:usa-or:oar-918-305` | OAR chapter 918, division 305: Electrical Codes and Standards | Oregon Secretary of State / Building Codes Division | https://secure.sos.state.or.us/oard/displayDivisionRules.action?selectedDivision=4158 | 2026-06-26 | 2023 OESC adoption; governing-code rule; existing installations; amendments; microgrid amendment |
| `src:usa-or:oar-918-400-0455` | OAR 918-400-0455: Adopted Oregon Elevator Specialty Code | Oregon Secretary of State / Building Codes Division | https://secure.sos.state.or.us/oard/view.action?ruleNumber=918-400-0455 | 2026-06-26 | 2024 Oregon Elevator Specialty Code; ASME components; A17.3 phase-in |
| `src:usa-or:oar-918-440-0010` | OAR 918-440-0010: Adopted Oregon Mechanical Specialty Code | Oregon Secretary of State / Building Codes Division | https://secure.sos.state.or.us/oard/view.action?ruleNumber=918-440-0010 | 2026-06-26 | 2025 OMSC; 2024 IMC/IFGC base; phase-in and effective date |
| `src:usa-or:oar-918-460-0010` | OAR 918-460-0010: Adopted Oregon Structural Specialty Code | Oregon Secretary of State / Building Codes Division | https://secure.sos.state.or.us/oard/view.action?ruleNumber=918-460-0010 | 2026-06-26 | 2025 OSSC; IBC/IFC/IEBC base; phase-in; plan-review/permit application date rule |
| `src:usa-or:oar-918-460-0015` | OAR 918-460-0015: Amendments to the Oregon Structural Specialty Code | Oregon Secretary of State / Building Codes Division | https://secure.sos.state.or.us/oard/view.action?ruleNumber=918-460-0015 | 2026-06-26 | OSSC amendment path; Chapter 34 snow-load repair/alteration amendment effective 2026-04-01 |
| `src:usa-or:oar-918-460-0500` | OAR 918-460-0500: Energy Provisions of the Oregon Structural Specialty Code | Oregon Secretary of State / Building Codes Division | https://secure.sos.state.or.us/oard/view.action?ruleNumber=918-460-0500 | 2026-06-26 | 2025 OEESC; ASHRAE 90.1-2022 base; phase-in; application-date rule |
| `src:usa-or:oar-918-480-0005` | OAR 918-480-0005: Adopted Oregon Residential Specialty Code | Oregon Secretary of State / Building Codes Division | https://secure.sos.state.or.us/oard/view.action?ruleNumber=918-480-0005 | 2026-06-26 | 2023 ORSC; phase-in; low-rise plumbing/electrical relationships; application-date rule |
| `src:usa-or:oar-918-480-0010` | OAR 918-480-0010: Amendments to the Oregon Residential Specialty Code | Oregon Secretary of State / Building Codes Division | https://secure.sos.state.or.us/oard/view.action?ruleNumber=918-480-0010 | 2026-06-26 | ORSC R310/R302.3/R327/R332 interim amendments |
| `src:usa-or:oar-918-750-0110` | OAR 918-750-0110: Oregon Plumbing Specialty Code | Oregon Secretary of State / Building Codes Division | https://secure.sos.state.or.us/oard/view.action?ruleNumber=918-750-0110 | 2026-06-26 | 2023 OPSC; UPC base; adopted/excluded appendices; local adoption of Appendices L and N |
| `src:usa-or:oar-837-040` | OAR chapter 837, division 40: Fire and Life Safety Regulations | Oregon Secretary of State / Department of the State Fire Marshal | https://secure.sos.state.or.us/oard/displayDivisionRules.action?selectedDivision=3805 | 2026-06-26 | 2025 OFC adoption; 2024 IFC base; 180-day OFC phase-in; fire-code authority |
| `src:usa-or:osfm-fire-code` | Oregon Fire Code | Oregon Department of the State Fire Marshal | https://www.oregon.gov/osfm/fire-service-partners/pages/oregon-fire-code.aspx | 2026-06-26 | OFC code portal; 2022 OFC valid-through note; 2025 OFC publication path |
| `src:usa-or:ors-455` | ORS chapter 455: Building Code | Oregon Legislative Assembly | https://www.oregonlegislature.gov/bills_laws/ors/ors455.html | 2026-06-26 | Statutory basis for state building code, Director/DCBS authority, uniformity/preemption, municipal building inspection programs |
| `src:usa-or:ors-476` | ORS chapter 476: State Fire Marshal / fire protection | Oregon Legislative Assembly | https://www.oregonlegislature.gov/bills_laws/ors/ors476.html | 2026-06-26 | Statutory basis for State Fire Marshal rulemaking and enforcement powers |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| `src:usa-or:ors-455` | official_html_fetch_issue | The official Oregon Legislature ORS page is the registered source, but the browser fetch encountered a decoding error. Statutory propositions were cross-checked against official OAR authority citations and official-source search snippets. | retain_official_url; manual_review_before_verified_status |
| `src:usa-or:ors-476` | official_html_fetch_issue | The official Oregon Legislature ORS page is the registered source, but the browser fetch encountered a decoding error. Fire-code authority was corroborated by OAR 837-040, which cites ORS 476.030. | retain_official_url; manual_review_before_verified_status |
| `src:usa-or:bcd-adopted-codes` | summary_portal | BCD adopted-code page is an official summary/portal, not the legal text. Code-specific OAR rules control where they differ. | use_for_cross_check_and_user_facing_dates |
| `src:usa-or:osfm-fire-code` | summary_portal | OSFM fire-code page is an official program page; OAR 837-040 controls adoption text and phase-in. | use_for_cross_check_and_codebook_access |
| `src:usa-or:oar-918-305` | table_not_embedded | OAR display notes that some tables referenced in rule text, such as OESC Table 1-E, are not included in the displayed rule text. | use_rule_for_adoption_dates_and_paths; table_text_requires_separate_extraction |
| `src:usa-or:oar-918-400-0455` | table_not_embedded | The elevator amendment table path is identified, but Table 2-L was not fully extracted. | use_rule_for_adoption_and_phase_in; table_text_requires_separate_extraction |

### 8.3 Supplemental Sources

No non-official sources were used as report authority. Search-result snippets from non-official legal mirrors were consulted only as navigational backup when official ORS pages could not be fetched; they are not cited as authority in this report.

### 8.4 Source Extraction Metadata

| Extraction ID | Source ID | Extracted Fields | Extraction Method | Extracted On | Notes |
| --- | --- | --- | --- | --- | --- |
| `extract:usa-or:001` | `src:usa-or:bcd-state-code` | BCD role, specialty-code programs, uniform code statement | HTML page text | 2026-06-26 | Official agency page |
| `extract:usa-or:002` | `src:usa-or:bcd-adopted-codes` | adopted-code portal dates and base-code descriptions | HTML page text | 2026-06-26 | Cross-checked with OAR where available |
| `extract:usa-or:003` | `src:usa-or:oar-918-460-0010` | OSSC base, effective date, phase-in, filing-date rule | OAR HTML | 2026-06-26 | Legal source for OSSC row |
| `extract:usa-or:004` | `src:usa-or:oar-918-440-0010` | OMSC base, effective date, phase-in | OAR HTML | 2026-06-26 | Legal source for OMSC row |
| `extract:usa-or:005` | `src:usa-or:oar-918-480-0005` | ORSC base/scope, phase-in, low-rise trade relationships, filing-date rule | OAR HTML | 2026-06-26 | Legal source for ORSC row |
| `extract:usa-or:006` | `src:usa-or:oar-918-750-0110` | OPSC base, appendices, local appendix adoption | OAR HTML | 2026-06-26 | Legal source for OPSC row |
| `extract:usa-or:007` | `src:usa-or:oar-918-305` | OESC adoption, governing-code rule, amendments | OAR HTML | 2026-06-26 | Division page used because it exposes multiple relevant electrical rules |
| `extract:usa-or:008` | `src:usa-or:oar-918-460-0500` | OEESC base, effective date, phase-in, filing-date rule | OAR HTML | 2026-06-26 | Legal source for OEESC row |
| `extract:usa-or:009` | `src:usa-or:oar-837-040` | OFC adoption, base, phase-in, local building official direction | OAR HTML | 2026-06-26 | Legal source for OFC row |
| `extract:usa-or:010` | `src:usa-or:oar-918-400-0455` | Elevator code base standards, phase-in | OAR HTML | 2026-06-26 | Legal source for elevator row |
| `extract:usa-or:011` | `src:usa-or:oar-918-020-0370` | Local amendment process | OAR HTML | 2026-06-26 | Legal source for local amendment section |
| `extract:usa-or:012` | `src:usa-or:bcd-oesc-adoption` | 2026 OESC pending adoption | HTML page text | 2026-06-26 | Future watch target |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 1.00 | none | Set after filling source-backed core authority and adoption fields and running validation checks. |
| report | risk.overall_confidence | 0.74 | reviewed | 0.74 | none | Reflects strong core-code coverage but unresolved AHJ, registry, and full amendment-diff work. |
| `ahj:usa-or:dcbs-bcd` | role | adopts, amends, interprets, and administers specialty codes making up Oregon State Building Code | verified | 0.90 | `src:usa-or:bcd-state-code`, `src:usa-or:oar-918-008` | Core authority source-backed. |
| `ahj:usa-or:state-fire-marshal` | OFC role | adopts Oregon Fire Code | partially_verified | 0.88 | `src:usa-or:oar-837-040`, `src:usa-or:ors-476` | ORS official URL retained for manual review; OAR 837-040 provides legal adoption text. |
| `adoption-record:usa-or:ossc-2025` | base/effective/phase-in | 2024 IBC/IFC/IEBC; 2025-10-01; phase-in through 2026-03-31 | verified | 0.93 | `src:usa-or:oar-918-460-0010` | OAR text explicit. |
| `adoption-record:usa-or:orsc-2023` | base/effective/phase-in | 2021 IRC; 2023-10-01; phase-in through 2024-03-31 | verified | 0.92 | `src:usa-or:oar-918-480-0005`, `src:usa-or:bcd-residential` | OAR and BCD portal align. |
| `adoption-record:usa-or:omsc-2025` | base/effective/phase-in | 2024 IMC/IFGC; 2025-10-01; phase-in through 2026-03-31 | verified | 0.91 | `src:usa-or:oar-918-440-0010` | OAR text explicit. |
| `adoption-record:usa-or:opsc-2023` | base/effective | 2021 UPC scope; 2023-10-01 | verified_core | 0.80 | `src:usa-or:oar-918-750-0110` | Mandatory date not separately identified. |
| `adoption-record:usa-or:oesc-2023` | base/effective | 2023 NEC, 2023 NESC, elevator electrical provisions; 2023-10-01 | verified_core | 0.86 | `src:usa-or:oar-918-305` | Mandatory date not separately identified; governing-code rule captured. |
| `adoption-record:usa-or:oeesc-2025` | base/effective/phase-in | ASHRAE 90.1-2022; 2025-01-01; phase-in to 2025-07-01 | verified | 0.90 | `src:usa-or:oar-918-460-0500`, `src:usa-or:bcd-energy` | OAR and BCD portal align. |
| `adoption-record:usa-or:ofc-2025` | base/effective/phase-in | 2024 IFC; 2025-10-01; phase-in through 2026-03-31 | verified_core | 0.86 | `src:usa-or:oar-837-040`, `src:usa-or:osfm-fire-code` | Mandatory date derived from end of 2022 OFC concurrency. |
| `local-amendment-rule:usa-or` | model | approval required for local differences in state building-code matters | verified_core | 0.86 | `src:usa-or:oar-918-020-0370`, `src:usa-or:ors-455` | Registry not found. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Every `src:usa-or:*` cited in body has a Section 8.1 registry entry or an explicit caveat in Section 8.2. |
| All authority IDs resolve | pass | Primary BCD and State Fire Marshal authorities are defined before use. |
| All current code families have adoption records | pass | Adoption matrix rows and normalized records exist for core current statewide families; accessibility is explicitly treated as specialty-code scope. |
| Building and operational fire code are separated | pass | OSSC fire-construction references and OFC operational/prevention code are separate rows. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Dates are separate; `null` is used where not extracted. |
| Effective dates are valid ISO dates | pass | Entered date values use YYYY-MM-DD. |
| No impossible date sequences | pass | Phase-in and mandatory dates follow effective dates. |
| Transition rules have explicit trigger conditions | pass | Primary phase-in and application/permit triggers are recorded. |
| Permit-date logic is captured where applicable | pass | BCD construction-code application-control rule and electrical earlier-of rule are captured. |
| Local enforcement model classified | pass | Classified as statewide code with municipal program administration; address-specific AHJ mapping remains open. |
| Local amendment rule classified | pass | Classified as uniform/preemptive with approval required. |
| AHJ confirmation metadata present | fail | No address-level AHJ contact or boundary data populated. |
| Official-source caveats captured | pass | Official ORS fetch issues and summary-portal limitations are documented. |
| Template markers removed | pass | No template markers, generic placeholders, or stub phrases remain. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| `issue:usa-or:001` | high | AHJ resolution | Address-level AHJ resolution is not built. Municipal program boundaries, county/city coverage, state-administered scopes, and fire jurisdiction boundaries are unresolved. | Build GIS/contact/source stack for BCD program administration, cities/counties, and fire jurisdictions. | null | null | open |
| `issue:usa-or:002` | high | local fire enforcement | OFC adoption authority is verified, but the local fire-code enforcement model and fire-code official hierarchy were not fully parsed. | Extract ORS 476 enforcement provisions, OSFM guidance, and local fire jurisdiction delegation rules. | null | null | open |
| `issue:usa-or:003` | medium | local amendment registry | Local amendment request process is verified, but a statewide registry of approved local amendments was not found. | Search BCD records, public notices, directives, and local amendment approvals. | null | null | open |
| `issue:usa-or:004` | medium | electrical future update | 2026 OESC adoption process is active; current report should be updated near the anticipated 2026-10-01 adoption. | Monitor BCD OESC adoption page and OAR 918-305 updates. | null | null | open |
| `issue:usa-or:005` | medium | amendment text extraction | Full Oregon amendment tables and integrated-code diffs are not parsed. | Extract Table 1-E, Table 2-L, OPSC amendments, OFC amendments, and OSSC/ORSC amendment documents. | null | null | open |
| `issue:usa-or:006` | medium | official ORS text | Official ORS HTML pages could not be parsed by the browser due to a decoding error. | Manually verify ORS 455.020, 455.030, 455.040, 455.148, 455.150, and 476.030 from official ORS source. | null | null | open |
| `issue:usa-or:007` | low | accessibility | Accessibility was treated as OSSC/specialty-code scope, not as a fully separate statewide accessibility authority record. | Extract OSSC Chapter 11, ADA/FHA cross-references, and Oregon-specific accessibility amendments. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| `watch:usa-or:bcd-adopted-codes` | `src:usa-or:bcd-adopted-codes` | html_diff | monthly | adopted-code portal updates or new effective/mandatory dates | 2026-06-26 | active |
| `watch:usa-or:oar-918-305` | `src:usa-or:oar-918-305` | html_diff | monthly through 2026-11-01 | 2026 OESC adoption or amendments to OAR 918-305 | 2026-06-26 | active |
| `watch:usa-or:bcd-oesc-adoption` | `src:usa-or:bcd-oesc-adoption` | html_diff | biweekly through 2026-10-15 | 2026 OESC adoption notice, rules filing, or effective-date change | 2026-06-26 | active |
| `watch:usa-or:oar-837-040` | `src:usa-or:oar-837-040` | html_diff | monthly | OFC amendments, phase-in changes, or OSFM code-cycle updates | 2026-06-26 | active |
| `watch:usa-or:oar-918-460` | `src:usa-or:oar-918-460-0010` | html_diff | monthly | OSSC/OEESC amendments or errata | 2026-06-26 | active |
| `watch:usa-or:oar-918-480` | `src:usa-or:oar-918-480-0005` | html_diff | monthly | ORSC amendments, wildfire local adoption updates, or errata | 2026-06-26 | active |
| `watch:usa-or:local-amendments` | `src:usa-or:oar-918-020-0370` | search_and_record_review | quarterly | newly approved local amendments or published registry source | 2026-06-26 | pending_registry |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-24 | Generated baseline draft report | `report:usa-or` | none | Codex | Original draft contained unresolved placeholders and no state-specific source registry. |
| 2026-06-26 | Populated Oregon authority, code-adoption matrix, date rules, amendment model, source registry, and QA | `ahj:usa-or:dcbs-bcd`; `ahj:usa-or:state-fire-marshal`; `adoption-record:usa-or:*`; `local-amendment-rule:usa-or` | `src:usa-or:bcd-state-code`; `src:usa-or:bcd-adopted-codes`; `src:usa-or:oar-918-008`; `src:usa-or:oar-918-020`; `src:usa-or:oar-918-020-0370`; `src:usa-or:oar-918-305`; `src:usa-or:oar-918-400-0455`; `src:usa-or:oar-918-440-0010`; `src:usa-or:oar-918-460-0010`; `src:usa-or:oar-918-460-0015`; `src:usa-or:oar-918-460-0500`; `src:usa-or:oar-918-480-0005`; `src:usa-or:oar-918-480-0010`; `src:usa-or:oar-918-750-0110`; `src:usa-or:oar-837-040`; `src:usa-or:osfm-fire-code`; `src:usa-or:ors-455`; `src:usa-or:ors-476` | ChatGPT | Upgraded to `partially_verified` after source-backed core fields and validation pass; unresolved AHJ and registry issues remain explicit. |
