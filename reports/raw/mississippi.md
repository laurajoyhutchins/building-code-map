---
state:
  state_id: "US-MS"
  name: "Mississippi"
  abbreviation: "MS"
report:
  report_id: "state-report:usa-ms"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-25"
  last_verified: null
  reviewed_by: null
risk:
  overall_confidence: 0.58 # 0.00 - 1.00
  risk_flags:
    - "official_code_portal_not_line_addressable"
    - "mbcc_current_edition_list_not_obtained"
    - "energy_code_authority_unresolved"
    - "accessibility_authority_unresolved"
    - "elevator_conveyance_authority_unresolved"
    - "local_opt_out_registry_not_obtained"
    - "local_fire_code_registry_not_obtained"
  open_questions_count: 7

---

# State Building Code Authority Report: Mississippi

## 1. Executive Summary

- **Authority model:** Mississippi uses a hybrid state/local model. The Mississippi Building Codes Council (MBCC) is the statewide body created by statute to adopt by reference and amend discretionary statewide minimum construction codes. Counties and municipalities are the ordinary adopting and enforcing entities for the State Uniform Construction Code, subject to the statutory opt-out provisions and specific exemptions. Fire prevention is administered through a separate State Fire Marshal / Mississippi Fire Prevention Code framework.

- **Statewide code status:** State law establishes the State Uniform Construction Code around the IBC, IRC, and other codes addressing electrical, plumbing, mechanical, fire, and fuel gas. The verified legal rule is edition-relative: one of the last three adopted editions as adopted and amended by MBCC, rather than a single fixed statewide edition in the statutes reviewed here. Coastal counties named in Miss. Code § 17-2-1 have a separate wind/flood mitigation rule tied to the 2003 IBC/IRC until they adopt the latest editions described in that section, but that section also contains an opt-out mechanism.

- **Local enforcement model:** Local enforcement is explicit. Codes adopted by a county board of supervisors or municipal governing authority under Miss. Code § 17-2-5 are enforced by that county or municipality. Local governments may also use agreements with other governmental entities or certified third-party providers for permitting and enforcement.

- **Local amendment posture:** Local governments may adopt and amend the minimum codes and may adopt construction codes that are not less stringent than the State Uniform Construction Code. A separate rule for fire prevention allows a county or municipality to adopt a fire prevention code with standards not less stringent than the Mississippi Fire Prevention Code; in that case the local code and enforcement mechanism apply instead of the Mississippi Fire Prevention Code for the covered local scope.

- **Known transition periods or pending changes:** The Fire Marshal rule in Mississippi Administrative Code Title 19, Part 7, Chapter 7 states that the current regulation supersedes prior versions and became effective 2025-01-01. No MBCC current-edition adoption order was obtained in this pass.

- **Production readiness:** limited_use. Core authority, local enforcement, and fire-prevention authority are source-backed, but the report should not be treated as complete until the MBCC current adopted-edition list, opt-out jurisdictions, and energy/accessibility/elevator authority paths are verified from primary sources.

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: Mississippi Building Codes Council is created by statute and adopts by
    reference/amends discretionary statewide minimum codes for IBC, IRC, and other
    listed code subjects.
  confidence: 0.78
  source_ids:
  - src:usa-ms:stat-17-2-3
- topic: State Uniform Construction Code
  finding: Counties and municipalities must adopt/amend as minimum codes one of the
    last three adopted editions of IBC or IRC, or other codes addressing electrical,
    plumbing, mechanical, fire, and fuel gas, subject to opt-out and exemptions.
  confidence: 0.74
  source_ids:
  - src:usa-ms:stat-17-2-4
- topic: Coastal wind/flood rule
  finding: Jackson, Harrison, Hancock, Stone, and Pearl River Counties, including
    municipalities, have a separate wind/flood mitigation provision based on the 2003
    IBC/IRC until local adoption of later minimum mandatory codes, with an opt-out
    process.
  confidence: 0.72
  source_ids:
  - src:usa-ms:stat-17-2-1
- topic: Local enforcement
  finding: Local governments enforce codes they adopt; they may use agreements with
    other governmental entities or certified third-party providers for permitting/enforcement.
  confidence: 0.78
  source_ids:
  - src:usa-ms:stat-17-2-5
- topic: Fire prevention authority
  finding: The State Fire Marshal promulgates and enforces the Mississippi Fire Prevention
    Code; the current rule is based on the most current IFC and became effective 2025-01-01.
  confidence: 0.86
  source_ids:
  - src:usa-ms:admin-19-7-ch7
- topic: Local fire code substitution
  finding: A county or municipality may adopt a fire prevention code not less stringent
    than the MFPC, and that local code/enforcement mechanism applies instead for the
    local covered scope.
  confidence: 0.82
  source_ids:
  - src:usa-ms:admin-19-7-ch7
- topic: Current code editions
  finding: Exact current MBCC adopted editions and amendment package were not obtained
    from an official MBCC adoption order in this pass.
  confidence: 0.35
  source_ids:
  - none
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-ms:mbcc |
| Authority name | Mississippi Building Codes Council |
| Authority type | statewide council / code-adoption body |
| Legal basis | Miss. Code § 17-2-3 |
| Role | Adopts by reference and amends one of the last three editions of specified model codes as discretionary statewide minimum codes; produces recommendations for mandatory statewide minimum codes. |
| Enforcement model | Local enforcement for locally adopted construction codes; MBCC is not verified as the direct statewide enforcement agency in this pass. |
| Source IDs | src:usa-ms:stat-17-2-3; src:usa-ms:stat-17-2-5 |
| Verification status | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-ms:mbcc | Mississippi Building Codes Council | Adopts/amends IBC-based discretionary statewide minimum code; local governments adopt/enforce as State Uniform Construction Code subject to opt-out/exemptions. | Miss. Code §§ 17-2-3, 17-2-4, 17-2-5 | src:usa-ms:stat-17-2-3; src:usa-ms:stat-17-2-4; src:usa-ms:stat-17-2-5 | partially_verified |
| Residential | ahj:usa-ms:mbcc | Mississippi Building Codes Council | Adopts/amends IRC-based discretionary statewide minimum code; statutory sprinkler limitation applies to MBCC, while local governments may adopt/modify/enforce codes including sprinkler requirements. | Miss. Code § 17-2-3 | src:usa-ms:stat-17-2-3 | partially_verified |
| Existing Building / Rehabilitation | ahj:usa-ms:mbcc | Mississippi Building Codes Council | IRC Appendix J is adopted by reference in the statutory description; separate IEBC authority not verified. | Miss. Code § 17-2-3 | src:usa-ms:stat-17-2-3 | limited_verified |
| Mechanical | ahj:usa-ms:mbcc | Mississippi Building Codes Council | Included among “other codes” addressing mechanical matters; exact model code edition not parsed. | Miss. Code §§ 17-2-3, 17-2-4 | src:usa-ms:stat-17-2-3; src:usa-ms:stat-17-2-4 | partially_verified |
| Plumbing | ahj:usa-ms:mbcc | Mississippi Building Codes Council | Included among “other codes” addressing plumbing matters; exact model code edition not parsed. | Miss. Code §§ 17-2-3, 17-2-4 | src:usa-ms:stat-17-2-3; src:usa-ms:stat-17-2-4 | partially_verified |
| Fuel Gas | ahj:usa-ms:mbcc | Mississippi Building Codes Council | Included among “other codes” addressing fuel gas matters; exact model code edition not parsed. | Miss. Code §§ 17-2-3, 17-2-4 | src:usa-ms:stat-17-2-3; src:usa-ms:stat-17-2-4 | partially_verified |
| Electrical | ahj:usa-ms:mbcc | Mississippi Building Codes Council | Included among “other codes” addressing electrical matters; exact model code edition not parsed. | Miss. Code §§ 17-2-3, 17-2-4 | src:usa-ms:stat-17-2-3; src:usa-ms:stat-17-2-4 | partially_verified |
| Energy | ahj:usa-ms:unresolved-energy | Unresolved | State energy-code authority and current IECC edition were not verified from primary Mississippi sources in this pass. | unresolved | none | unresolved |
| Fire - construction references | ahj:usa-ms:mbcc | Mississippi Building Codes Council | Fire appears among “other codes” for construction-code adoption; operational fire-prevention code is handled separately by the State Fire Marshal. | Miss. Code §§ 17-2-3, 17-2-4 | src:usa-ms:stat-17-2-3; src:usa-ms:stat-17-2-4 | partially_verified |
| Fire - operational / prevention code | ahj:usa-ms:sfm | State Fire Marshal / Mississippi Insurance Department | Promulgates, interprets, and enforces Mississippi Fire Prevention Code within statutory scope; local not-less-stringent fire codes may displace MFPC for the local covered scope. | Miss. Code §§ 45-11-101 through 45-11-105; Miss. Admin. Code Title 19, Part 7, Chapter 7 | src:usa-ms:admin-19-7-ch7 | verified_scope_limited |
| Accessibility | ahj:usa-ms:unresolved-accessibility | Unresolved | Separate statewide accessibility authority not verified; IBC accessibility provisions may apply through IBC adoption but were not separately parsed. | unresolved | none | unresolved |
| Elevator / Conveyance | ahj:usa-ms:unresolved-elevator | Unresolved | Elevator/conveyance authority not researched in this pass. | unresolved | none | unresolved |

### 2.3 Authority Hierarchy Notes

Mississippi's construction-code model should be treated as layered rather than a single direct statewide enforcement system:

1. MBCC is the statewide code-adoption council for the discretionary statewide minimum construction-code set.
2. Counties and municipalities are the adopting and enforcing entities for construction codes under Miss. Code §§ 17-2-4 and 17-2-5, with statutory opt-out and exemption rules.
3. Coastal counties named in Miss. Code § 17-2-1 have a separate wind/flood mitigation baseline unless they adopted later minimum mandatory codes or opted out under that section.
4. The State Fire Marshal administers the Mississippi Fire Prevention Code for a distinct fire-prevention scope. Local fire codes not less stringent than MFPC may substitute for MFPC and local enforcement in the local jurisdiction, except where state-owned/state-agency scope remains reserved.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-ms:001 | ahj:usa-ms:mbcc | adopts_and_amends | discretionary statewide minimum construction codes | src:usa-ms:stat-17-2-3 | partially_verified |
| edge:usa-ms:002 | ahj:usa-ms:mbcc | informs_minimum_for | county and municipal code adoption | src:usa-ms:stat-17-2-4; src:usa-ms:stat-17-2-5 | partially_verified |
| edge:usa-ms:003 | county_or_municipal_governing_authority | enforces | codes adopted by that county or municipality | src:usa-ms:stat-17-2-5 | partially_verified |
| edge:usa-ms:004 | county_or_municipal_governing_authority | may_contract_with | other governmental entities or certified third-party providers for permits/enforcement | src:usa-ms:stat-17-2-5 | partially_verified |
| edge:usa-ms:005 | ahj:usa-ms:sfm | promulgates_and_enforces | Mississippi Fire Prevention Code | src:usa-ms:admin-19-7-ch7 | verified_scope_limited |
| edge:usa-ms:006 | county_or_municipal_fire_official | may_enforce_or_substitute | local fire prevention code not less stringent than MFPC | src:usa-ms:admin-19-7-ch7 | verified_scope_limited |
| edge:usa-ms:007 | county_or_municipal_governing_authority | may_opt_out | State Uniform Construction Code requirements within statutory opt-out window | src:usa-ms:stat-17-2-4 | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | State Uniform Construction Code / MBCC discretionary statewide minimum code | International Building Code | One of the last three adopted editions as adopted and amended by MBCC; exact current MBCC edition list not obtained | local adoption/enforcement with opt-out; coastal rule has separate baseline | 2014-08-01 | 2014-08-01 | null | null | Jurisdictions could opt out within 120 days after §17-2-4 became effective; coastal counties have separate §17-2-1 transition. | src:usa-ms:stat-17-2-3; src:usa-ms:stat-17-2-4; src:usa-ms:stat-17-2-1 |
| Residential | State Uniform Construction Code / MBCC discretionary statewide minimum code | International Residential Code | One of the last three adopted editions as adopted and amended by MBCC; exact current MBCC edition list not obtained | local adoption/enforcement with opt-out; sprinkler limitation applies to MBCC but local governments may adopt/modify/enforce sprinkler codes | 2014-08-01 | 2014-08-01 | null | null | Jurisdictions could opt out within 120 days after §17-2-4 became effective; coastal counties have separate §17-2-1 transition. | src:usa-ms:stat-17-2-3; src:usa-ms:stat-17-2-4; src:usa-ms:stat-17-2-1 |
| Existing Building / Rehabilitation | IRC Appendix J reference | IRC Appendix J, Existing Buildings and Structures | Adopted by reference in statute; exact MBCC amendment text not obtained | limited_verified | null | null | null | null | Appendix J appears in statutory IRC description; separate IEBC adoption not verified. | src:usa-ms:stat-17-2-3 |
| Mechanical | State Uniform Construction Code other-code component | Other codes addressing mechanical matters | Not parsed | local adoption/enforcement if adopted by county or municipality | 2014-08-01 | 2014-08-01 | null | null | Must be within the code subjects adopted/amended by MBCC and subject to local opt-out/exemptions. | src:usa-ms:stat-17-2-3; src:usa-ms:stat-17-2-4 |
| Plumbing | State Uniform Construction Code other-code component | Other codes addressing plumbing matters | Not parsed | local adoption/enforcement if adopted by county or municipality | 2014-08-01 | 2014-08-01 | null | null | Must be within the code subjects adopted/amended by MBCC and subject to local opt-out/exemptions. | src:usa-ms:stat-17-2-3; src:usa-ms:stat-17-2-4 |
| Fuel Gas | State Uniform Construction Code other-code component | Other codes addressing fuel gas matters | Not parsed | local adoption/enforcement if adopted by county or municipality | 2014-08-01 | 2014-08-01 | null | null | Must be within the code subjects adopted/amended by MBCC and subject to local opt-out/exemptions. | src:usa-ms:stat-17-2-3; src:usa-ms:stat-17-2-4 |
| Electrical | State Uniform Construction Code other-code component | Other codes addressing electrical matters | Not parsed | local adoption/enforcement if adopted by county or municipality | 2014-08-01 | 2014-08-01 | null | null | Must be within the code subjects adopted/amended by MBCC and subject to local opt-out/exemptions. | src:usa-ms:stat-17-2-3; src:usa-ms:stat-17-2-4 |
| Energy | unresolved | IECC or other energy code not verified from primary Mississippi source | unresolved | unresolved | null | null | null | null | State energy-code source path must be verified before completing this row. | none |
| Fire - construction references | State Uniform Construction Code other-code component | Other codes addressing fire matters | Not parsed | local adoption/enforcement if adopted by county or municipality | 2014-08-01 | 2014-08-01 | null | null | Construction-code fire references should be kept separate from MFPC operational/prevention scope. | src:usa-ms:stat-17-2-3; src:usa-ms:stat-17-2-4 |
| Fire - operational / prevention code | Mississippi Fire Prevention Code | International Fire Code | Most current edition of IFC as revised or amended | statewide within MFPC statutory scope unless local not-less-stringent fire code applies | null | 2025-01-01 | 2025-01-01 | null | Current regulation supersedes prior Title 19, Part 7, Chapter 7 versions effective 2025-01-01. | src:usa-ms:admin-19-7-ch7 |
| Accessibility | unresolved | unresolved | unresolved | unresolved | null | null | null | null | Separate accessibility authority not verified; IBC-derived accessibility provisions not parsed. | none |
| Elevator / Conveyance | unresolved | unresolved | unresolved | unresolved | null | null | null | null | Elevator/conveyance authority not researched in this pass. | none |

### 3.2 Adoption Records

| Adoption Record ID | Code Family | Adopting / Responsible Authority | Instrument | Adoption Date | Effective Date | Operative Date | Mandatory Date | Scope | Source IDs | Confidence |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| adoption:usa-ms:mbcc-discretionary-minimum-codes | Building; Residential; other listed trade-code subjects | Mississippi Building Codes Council | Miss. Code § 17-2-3 | null | 2012-07-01 | null | null | MBCC adopts/amends one of the last three editions of IBC, IRC, and other codes addressing electrical, plumbing, mechanical, fire, and fuel gas as discretionary statewide minimum codes. | src:usa-ms:stat-17-2-3 | 0.72 |
| adoption:usa-ms:state-uniform-construction-code-2014 | Building; Residential; other listed trade-code subjects | County boards of supervisors and municipal governing authorities | Miss. Code § 17-2-4, added by Laws 2014, ch. 382, SB 2378 | 2014-08-01 | 2014-08-01 | null | null | Local adoption/amendment of minimum codes as State Uniform Construction Code; local opt-out within 120 days; exemptions apply. | src:usa-ms:stat-17-2-4 | 0.74 |
| adoption:usa-ms:coastal-wind-flood-2006 | Building; Residential wind/flood mitigation | Jackson, Harrison, Hancock, Stone, and Pearl River Counties and municipalities therein | Miss. Code § 17-2-1 | 2006-04-14 | 2006-05-14 | 2006-05-14 | null | 2003 IBC/IRC wind and flood mitigation baseline until adoption of later minimum mandatory codes; opt-out permitted. | src:usa-ms:stat-17-2-1 | 0.70 |
| adoption:usa-ms:local-enforcement-2008 | Local enforcement | County boards of supervisors and municipal governing authorities | Miss. Code § 17-2-5 | 2008-07-01 | 2008-07-01 | 2008-07-01 | null | Local codes adopted after 2008-07-01 must use MBCC minimum codes; local governments enforce and may use intergovernmental or certified third-party arrangements. | src:usa-ms:stat-17-2-5 | 0.78 |
| adoption:usa-ms:mfpc-2025 | Fire - operational / prevention | State Fire Marshal | Miss. Admin. Code Title 19, Part 7, Chapter 7 | null | 2025-01-01 | 2025-01-01 | null | MFPC based on the most current IFC, applied to specified building classes and enforced by State Fire Marshal/state and local governments; local not-less-stringent fire codes may apply instead. | src:usa-ms:admin-19-7-ch7 | 0.86 |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

Mississippi's reviewed sources distinguish multiple date concepts. Section 17-2-4 became effective 2014-08-01 and allowed counties and municipalities to opt out within 120 days after that effective date. Section 17-2-1 for coastal wind/flood mitigation was enacted effective 2006-04-14, but its substantive provisions went into effect 30 days later and local governments could opt out within 60 days after those provisions went into effect. The Fire Marshal regulation states an effective date of 2025-01-01 for the current Title 19, Part 7, Chapter 7 regulation.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-ms:001 | State Uniform Construction Code under §17-2-4 | statutory effective date | 2014-08-01 | Laws 2014, ch. 382, SB 2378 became effective | unresolved | src:usa-ms:stat-17-2-4 | partially_verified |
| date-rule:usa-ms:002 | State Uniform Construction Code opt-out | opt-out/grace period | 120 days after 2014-08-01 | county or municipality adopts resolution and enters it in minutes | yes, if opted out; exact consequence by jurisdiction unresolved | src:usa-ms:stat-17-2-4 | partially_verified |
| date-rule:usa-ms:003 | Coastal wind/flood mitigation rule | operative delay | 30 days from 2006-04-14 | §17-2-1 provisions go into effect | prior practice unresolved | src:usa-ms:stat-17-2-1 | partially_verified |
| date-rule:usa-ms:004 | Coastal wind/flood opt-out | opt-out/grace period | 60 days after §17-2-1 provisions go into effect | county or municipality adopts resolution and enters it in minutes | yes, if opted out; exact opt-out inventory unresolved | src:usa-ms:stat-17-2-1 | partially_verified |
| date-rule:usa-ms:005 | Mississippi Fire Prevention Code | regulatory effective date | 2025-01-01 | current Title 19, Part 7, Chapter 7 supersedes prior versions | prior versions superseded | src:usa-ms:admin-19-7-ch7 | verified_scope_limited |
| date-rule:usa-ms:006 | MFPC plan review | plan approval date rule | version in effect at plan approval | State Fire Marshal plan approval for MFPC-covered classes | no explicit concurrency rule beyond plan-approval version found | src:usa-ms:admin-19-7-ch7 | verified_scope_limited |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Fire - operational / prevention code | Current rule already effective | null | null | 2025-01-01 | 2025-01-01 | null | active_monitoring | src:usa-ms:admin-19-7-ch7 | Watch SOS Administrative Code Title 19, Part 7, Chapter 7 for amendments. |
| Building / Residential / trade codes | unresolved | null | null | null | null | null | needs_source | none | Obtain MBCC current adopted-edition list or meeting/adoption order. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-ms:001 | Building / Residential | Jackson, Harrison, Hancock, Stone, Pearl River Counties and municipalities therein | Coastal county/municipality under §17-2-1 | Wind/flood mitigation requirements based on 2003 IBC/IRC continue until adoption of later minimum mandatory codes, subject to opt-out. | src:usa-ms:stat-17-2-1 | partially_verified |
| applicability-rule:usa-ms:002 | Building / Residential / trade code | Farm structures | Owner files required affidavit; structure qualifies as farm structure | Local governments may not enforce building code portions regulating construction/improvement of farm structures, except NFIP floodplain management ordinances/regulations remain applicable. | src:usa-ms:stat-17-2-7 | partially_verified |
| applicability-rule:usa-ms:003 | Building / Residential / trade code | Industrial and certain other exempt structures | Project falls within §17-2-9 categories | Local governments may not enforce specified code portions for listed industrial facilities, fairground buildings, qualifying private unattached recreational structures, HUD-code manufactured housing, certain Pearl River County salvage/green-timber personal-use construction, with NFIP floodplain exception. | src:usa-ms:stat-17-2-9 | partially_verified |
| applicability-rule:usa-ms:004 | Fire - operational / prevention | State-owned/state-agency buildings, public assembly, high-rise, private correctional, request-based review, private fraternity/sorority houses on state property | Building falls within MFPC scope | MFPC applies to specified classes; local not-less-stringent fire code may apply instead for several non-state categories. | src:usa-ms:admin-19-7-ch7 | verified_scope_limited |
| applicability-rule:usa-ms:005 | State/county/municipal building code | Refrigerants in commercial or residential buildings/construction | Refrigerant is federally acceptable under 42 USC 7671k and equipment is listed/installed under safety standards/use conditions | State, county, or municipal building codes may not prohibit or limit the covered refrigerant. | src:usa-ms:stat-17-2-11 | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** MBCC may adopt by reference and amend the specified model codes. Local governments may adopt and amend as minimum codes under the State Uniform Construction Code and may adopt construction codes not less stringent than the codes in §17-2-4(1). MFPC is rule-based through the State Fire Marshal and includes a local not-less-stringent substitution model.

**Where amendments are published:** unresolved for MBCC current construction-code amendments; MFPC rule text is published in Mississippi Administrative Code Title 19, Part 7, Chapter 7.

**Amendment parsing status:** partial_statutory_framework_only

### 5.2 State Amendment Sources

| Amendment Source ID | Code Family | Publication Path | Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| amendment-source:usa-ms:mbcc | Building; Residential; Mechanical; Plumbing; Fuel Gas; Electrical; Fire construction references | MBCC adoption/amendment materials not obtained | Statutory authority for MBCC amendments verified; actual amendment text unresolved. | src:usa-ms:stat-17-2-3; src:usa-ms:stat-17-2-4 | unresolved_text |
| amendment-source:usa-ms:mfpc | Fire - operational / prevention | Mississippi Administrative Code Title 19, Part 7, Chapter 7 | MFPC rules, definitions, applicability, plan review, enforcement, appeals, effective date. | src:usa-ms:admin-19-7-ch7 | partially_parsed |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Summary | Source IDs | Status |
| --- | --- | --- | --- |
| amendment:usa-ms:residential-sprinkler-mbcc-limitation | Residential | MBCC may not enact an ordinance, bylaw, order, building code, or rule requiring multipurpose residential fire sprinklers or other fire sprinkler systems in new or existing one- or two-family dwellings; county and municipal governing authorities may adopt, modify, and enforce codes adopted by the council, including sprinkler requirements in any structure. | src:usa-ms:stat-17-2-3 | partially_verified |
| amendment:usa-ms:irc-appendix-j | Existing Building / Rehabilitation | IRC Appendix J, Existing Buildings and Structures, is adopted by reference in the statutory IRC description. | src:usa-ms:stat-17-2-3 | partially_verified |
| amendment:usa-ms:mfpc-most-current-ifc | Fire - operational / prevention | MFPC is based on the most current edition of the International Fire Code as revised or amended. | src:usa-ms:admin-19-7-ch7 | verified_scope_limited |
| amendment:usa-ms:refrigerant-protection | Building / Residential | State, county, and municipal building codes may not prohibit or otherwise limit a federally acceptable refrigerant when listed and installed under applicable safety standards and use conditions. | src:usa-ms:stat-17-2-11 | partially_verified |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-ms"
  model: "local_adoption_and_local_enforcement_with_state_minimum_code_framework"
  enforcing_entities:
    - "county boards of supervisors"
    - "municipal governing authorities"
    - "other governmental entities by agreement"
    - "certified third-party providers by agreement"
    - "State Fire Marshal / State Chief Deputy Fire Marshal for MFPC-covered state scope"
    - "local fire officials where local not-less-stringent fire code applies"
  required_officials:
    - "unresolved for construction code"
    - "State Chief Deputy Fire Marshal as fire official for MFPC"
  state_reserved_activities:
    - "MFPC enforcement for state-owned/state-agency scope and specified MFPC-covered classes"
    - "MFPC plan review where required by rule"
  source_ids:
    - "src:usa-ms:stat-17-2-5"
    - "src:usa-ms:admin-19-7-ch7"
  verification_status: "partially_verified"
  confidence: 0.78
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-ms"
  model: "local_not_less_stringent_with_historical_opt_out"
  applies_to_code_families:
    - "building"
    - "residential"
    - "electrical"
    - "plumbing"
    - "mechanical"
    - "fire construction references"
    - "fuel gas"
    - "fire prevention"
  approval_required: "unresolved for construction codes; local fire codes are subject to State Fire Marshal annual examination/listing for not-less-stringent status"
  approving_authority_id: "ahj:usa-ms:sfm for local fire-code not-less-stringency list; construction-code approving authority unresolved"
  filing_required: "local opt-out resolutions must be adopted and entered upon local minutes; filing/registry for amendments unresolved"
  registry_exists: "rule requires annual State Fire Marshal list of county/municipal fire prevention codes, but public registry was not obtained"
  registry_source_ids:
    - "src:usa-ms:admin-19-7-ch7"
  legal_basis_source_ids:
    - "src:usa-ms:stat-17-2-3"
    - "src:usa-ms:stat-17-2-4"
    - "src:usa-ms:stat-17-2-5"
    - "src:usa-ms:admin-19-7-ch7"
  verification_status: "partially_verified"
  confidence: 0.66
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Construction-code enforcement and local amendment authority are related but distinct. The reviewed statutes say counties and municipalities enforce the codes they adopt, and may use agreements or certified third-party providers for permitting/enforcement. Separately, the statutes allow adoption/amendment of minimum codes and additional construction codes not less stringent than the State Uniform Construction Code. For fire prevention, the reviewed rule expressly separates the MFPC from local fire prevention codes: when a county or municipality adopts a not-less-stringent fire prevention code, that local code and enforcement mechanism apply instead of MFPC for the applicable local scope.

### 6.4 Known Local Amendment Registries

| Registry ID | Registry Name | Scope | Publication Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- |
| registry:usa-ms:sfm-local-fire-codes | State Fire Marshal annual list of county/municipal fire prevention codes | Local fire prevention codes and whether they are not less stringent than MFPC | required_by_rule_but_not_obtained | src:usa-ms:admin-19-7-ch7 | Rule 7.03/7.06 references annual examination/listing; public list location unresolved. |
| registry:usa-ms:construction-opt-outs | State Uniform Construction Code local opt-out resolutions | Counties/municipalities opting out of §17-2-4 | unresolved | src:usa-ms:stat-17-2-4 | Local minutes may be the authoritative record unless a central registry exists. |

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendments were parsed in this pass. Add local records only after obtaining the local ordinance, resolution, minutes, or official code portal for the jurisdiction.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: partial_model_only

Jurisdiction stack:

```text
Address
  -> State: Mississippi
  -> County
  -> Municipality if incorporated, or unincorporated county if outside municipality
  -> Coastal-county status: Jackson / Harrison / Hancock / Stone / Pearl River
  -> Local opt-out status under Miss. Code § 17-2-4 and, if coastal, § 17-2-1
  -> Building-code adopting/enforcing authority
  -> Fire-code AHJ: State Fire Marshal unless local not-less-stringent fire code applies for the class of building; state-owned/state-agency scope remains state-relevant
  -> Trade-specific AHJs as locally adopted or state-administered
  -> Applicable code adoption records
  -> Applicable local amendment records
  -> Exemption check: farm structure, listed industrial/fairground/recreational/manufactured-housing exemptions, NFIP floodplain overlay
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | not selected | none | statewide | unresolved | pending |
| County | not selected | none | statewide | unresolved | pending |
| Municipality | not selected | none | statewide | unresolved | pending |
| Fire District | not selected | none | statewide | unresolved | pending |
| Special District | not selected | none | statewide | unresolved | pending |
| Coastal County Flag | statutory list in Miss. Code § 17-2-1 | src:usa-ms:stat-17-2-1 | Jackson, Harrison, Hancock, Stone, Pearl River | statutory | partially_verified |

### 7.3 AHJ Contact Data

No AHJ contact data was populated for this pass. For production use, create AHJ records for state fire marshal contacts, county building departments, municipal building departments, and local fire officials, then link each to local code-adoption and opt-out records.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Source Type | Title | Issuer / Publisher | URL / Location | Date / Version | Key Fields Supported | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-ms:code-public-access | official_code_portal | Mississippi Code Public Access | LexisNexis for Mississippi Code public access | https://www.lexisnexis.com/hottopics/mscode/ | accessed 2026-06-25 | Statutory code portal for Mississippi Code | Official/public code portal is not line-addressable in this workflow; section text was cross-checked against accessible 2024 HTML convenience copies listed as caveated extraction aids. |
| src:usa-ms:stat-17-2-1 | statute | Miss. Code § 17-2-1, Certain counties required to enforce wind/flood mitigation requirements | Mississippi Code / public access copy | https://law.justia.com/codes/mississippi/title-17/chapter-2/section-17-2-1/ | 2024 accessible copy; official portal cross-check required | Coastal wind/flood rule; opt-out window; coastal county list | Caveat: accessible text is Justia convenience HTML with disclaimer; confirm against official code portal before verified status. |
| src:usa-ms:stat-17-2-3 | statute | Miss. Code § 17-2-3, Creation of Mississippi Building Codes Council | Mississippi Code / public access copy | https://law.justia.com/codes/mississippi/title-17/chapter-2/section-17-2-3/ | 2024 accessible copy; official portal cross-check required | MBCC creation; membership; model-code adoption/amendment authority; sprinkler limitation; local modification clause | Caveat: accessible text is Justia convenience HTML with disclaimer; confirm against official code portal before verified status. |
| src:usa-ms:stat-17-2-4 | statute | Miss. Code § 17-2-4, State Uniform Construction Code; exemptions | Mississippi Code / public access copy | https://law.justia.com/codes/mississippi/title-17/chapter-2/section-17-2-4/ | 2024 accessible copy; official portal cross-check required | State Uniform Construction Code; IBC/IRC/other code subjects; local opt-out; effective date | Caveat: accessible text is Justia convenience HTML with disclaimer; confirm against official code portal before verified status. |
| src:usa-ms:stat-17-2-5 | statute | Miss. Code § 17-2-5, Adoption of minimum codes; agreements for enforcement | Mississippi Code / public access copy | https://law.justia.com/codes/mississippi/title-17/chapter-2/section-17-2-5/ | 2024 accessible copy; official portal cross-check required | Local adoption after 2008; local enforcement; third-party/intergovernmental enforcement agreements | Caveat: accessible text is Justia convenience HTML with disclaimer; confirm against official code portal before verified status. |
| src:usa-ms:stat-17-2-7 | statute | Miss. Code § 17-2-7, Farm structures exempt | Mississippi Code / public access copy | https://law.justia.com/codes/mississippi/title-17/chapter-2/section-17-2-7/ | 2024 accessible copy; official portal cross-check required | Farm-structure exemption; affidavit condition; NFIP carveout | Caveat: accessible text is Justia convenience HTML with disclaimer; confirm against official code portal before verified status. |
| src:usa-ms:stat-17-2-9 | statute | Miss. Code § 17-2-9, Certain other buildings/facilities/manufactured housing exempt | Mississippi Code / public access copy | https://law.justia.com/codes/mississippi/title-17/chapter-2/section-17-2-9/ | 2024 accessible copy; official portal cross-check required | Industrial/fairground/recreational/manufactured-housing/salvage timber exemptions; NFIP carveout | Caveat: accessible text is Justia convenience HTML with disclaimer; confirm against official code portal before verified status. |
| src:usa-ms:stat-17-2-11 | statute | Miss. Code § 17-2-11, Refrigerant limitation | Mississippi Code / public access copy | https://law.justia.com/codes/mississippi/title-17/chapter-2/section-17-2-11/ | 2024 accessible copy; official portal cross-check required | Building-code prohibition on limiting federally acceptable refrigerants | Caveat: accessible text is Justia convenience HTML with disclaimer; confirm against official code portal before verified status. |
| src:usa-ms:admin-19-7-ch7 | administrative_rule | Mississippi Administrative Code Title 19, Part 7, Chapter 7, Rules and Regulations for the Mississippi Fire Prevention Code | Mississippi Secretary of State Administrative Code / Mississippi Insurance Department, State Fire Marshal | https://www.sos.ms.gov/adminsearch/ACCode/00000310c.pdf | current compilation accessed 2026-06-25; rule effective 2025-01-01 | MFPC authority, model code, applicability, state/local enforcement, local fire-code substitution, plan review date rule, effective date | Official PDF compilation; text extraction and page screenshots reviewed. |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-ms:code-public-access | official_portal_access | Official code portal is accessible but was not parsed into line-addressable section text in this workflow. | Recheck final statutory language in official portal before verified status. |
| src:usa-ms:stat-17-2-1 | unofficial_html_extraction | Justia convenience HTML used for line-readable text; site includes disclaimer to check official sources. | Treat as partially verified until official code portal text is saved or independently captured. |
| src:usa-ms:stat-17-2-3 | unofficial_html_extraction | Justia convenience HTML used for line-readable text; site includes disclaimer to check official sources. | Treat as partially verified until official code portal text is saved or independently captured. |
| src:usa-ms:stat-17-2-4 | unofficial_html_extraction | Justia convenience HTML used for line-readable text; site includes disclaimer to check official sources. | Treat as partially verified until official code portal text is saved or independently captured. |
| src:usa-ms:stat-17-2-5 | unofficial_html_extraction | Justia convenience HTML used for line-readable text; site includes disclaimer to check official sources. | Treat as partially verified until official code portal text is saved or independently captured. |
| src:usa-ms:stat-17-2-7 | unofficial_html_extraction | Justia convenience HTML used for line-readable text; site includes disclaimer to check official sources. | Treat as partially verified until official code portal text is saved or independently captured. |
| src:usa-ms:stat-17-2-9 | unofficial_html_extraction | Justia convenience HTML used for line-readable text; site includes disclaimer to check official sources. | Treat as partially verified until official code portal text is saved or independently captured. |
| src:usa-ms:stat-17-2-11 | unofficial_html_extraction | Justia convenience HTML used for line-readable text; site includes disclaimer to check official sources. | Treat as partially verified until official code portal text is saved or independently captured. |
| src:usa-ms:admin-19-7-ch7 | pdf_text_extraction | Official PDF compilation; page numbering and extracted text should be preserved if future diffs are required. | Accept for partially_verified; archive a copy for reproducible production use. |

### 8.3 Supplemental Sources

| Source ID | Source Type | Title | Issuer / Publisher | URL / Location | Date / Version | Usage | Caveat |
| --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-ms:supp-doe-energycodes | federal_supplemental | Mississippi state page, Building Energy Codes Program | U.S. Department of Energy | https://www.energycodes.gov/status/states/mississippi | accessed 2026-06-25 | Identifies Mississippi energy-code contacts only; not used to complete adoption matrix. | Supplemental federal page; not a Mississippi primary authority source. |
| src:usa-ms:supp-icc-adoption-map | trade_association_supplemental | ICC Mississippi adoption map | International Code Council | https://www.iccsafe.org/advocacy/adoptions-map/mississippi/ | accessed 2026-06-25 | Background only; not used as primary support for current edition rows. | Non-governmental; use only to guide official-source follow-up. |
| src:usa-ms:supp-fema-bcat | federal_supplemental | FEMA Building Code Adoption Tracking | Federal Emergency Management Agency | https://www.fema.gov/emergency-managers/risk-management/building-science/bcat | accessed 2026-06-25 | Background on code-adoption tracking; not used to set MBCC editions. | Federal tracking may lag or summarize local practice. |

### 8.4 Source Extraction Metadata

| Extraction ID | Source ID | Extracted By | Date | Method | Pages / Lines Reviewed | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| extract:usa-ms:001 | src:usa-ms:stat-17-2-1 | ChatGPT | 2026-06-25 | web text extraction | 2024 HTML lines covering §17-2-1 | Captured coastal county list, 2003 IBC/IRC wind/flood baseline, opt-out language. |
| extract:usa-ms:002 | src:usa-ms:stat-17-2-3 | ChatGPT | 2026-06-25 | web text extraction | 2024 HTML lines covering §17-2-3 | Captured MBCC creation, membership, code-adoption authority, sprinkler limitation. |
| extract:usa-ms:003 | src:usa-ms:stat-17-2-4 | ChatGPT | 2026-06-25 | web text extraction | 2024 HTML lines covering §17-2-4 | Captured State Uniform Construction Code, opt-out, effective date. |
| extract:usa-ms:004 | src:usa-ms:stat-17-2-5 | ChatGPT | 2026-06-25 | web text extraction | 2024 HTML lines covering §17-2-5 | Captured local enforcement and agreements. |
| extract:usa-ms:005 | src:usa-ms:admin-19-7-ch7 | ChatGPT | 2026-06-25 | PDF text extraction plus screenshot review | PDF pages around 134, 140, 150 | Captured MFPC authority, IFC basis, applicability, enforcement, local substitution, effective date. |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 1.00 | none | Core authority and fire-code fields are source-backed; several code-family rows remain unresolved. |
| state | state_id/name/abbreviation | US-MS / Mississippi / MS | verified | 1.00 | none | From uploaded state draft metadata. |
| ahj:usa-ms:mbcc | authority.name | Mississippi Building Codes Council | partially_verified | 0.78 | src:usa-ms:stat-17-2-3 | Statute text captured via caveated accessible copy; official portal recheck required before full verification. |
| ahj:usa-ms:mbcc | role | adopts/amends discretionary statewide minimum codes | partially_verified | 0.78 | src:usa-ms:stat-17-2-3 | Exact current edition list unresolved. |
| adoption:usa-ms:state-uniform-construction-code-2014 | effective_date | 2014-08-01 | partially_verified | 0.74 | src:usa-ms:stat-17-2-4 | Effective date stated in accessible statute copy. |
| adoption:usa-ms:coastal-wind-flood-2006 | coastal_counties | Jackson, Harrison, Hancock, Stone, Pearl River | partially_verified | 0.72 | src:usa-ms:stat-17-2-1 | Includes municipalities within those counties. |
| local-enforcement:usa-ms | model | local adoption and local enforcement with state minimum code framework | partially_verified | 0.78 | src:usa-ms:stat-17-2-5 | Third-party and intergovernmental agreements verified in statute. |
| local-amendment-rule:usa-ms | model | local not less stringent with historical opt-out | partially_verified | 0.66 | src:usa-ms:stat-17-2-4; src:usa-ms:admin-19-7-ch7 | Construction amendment registry unresolved; fire list required by rule but not obtained. |
| adoption:usa-ms:mfpc-2025 | effective_date | 2025-01-01 | verified_scope_limited | 0.86 | src:usa-ms:admin-19-7-ch7 | Source is official SOS Administrative Code PDF. |
| adoption:usa-ms:mfpc-2025 | base_model_code | most current edition of International Fire Code | verified_scope_limited | 0.86 | src:usa-ms:admin-19-7-ch7 | MFPC rule text states the basis. |
| energy | current_state_code | unresolved | unresolved | 0.20 | none | DOE/ICC supplemental leads found, but Mississippi primary source not obtained. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Every `src:usa-ms:*` used in body appears in section 8. |
| All authority IDs resolve | pass | Authority IDs used in sections 2 and 6 are defined or deliberately marked unresolved. |
| All current code families have adoption records | fail | Energy, accessibility, and elevator/conveyance remain unresolved. |
| Building and operational fire code are separated | pass | Construction-code fire references and MFPC operational/prevention scope are separate rows. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Dates are separate; unresolved mandatory dates are null. |
| Effective dates are valid ISO dates | pass | Populated effective dates use YYYY-MM-DD. |
| No impossible date sequences | pass | No date sequence contradiction introduced; computed opt-out end dates are not entered as fixed mandatory dates. |
| Transition rules have explicit trigger conditions | pass | Opt-out and MFPC effective-date triggers are described. |
| Permit-date logic is captured where applicable | partial | MFPC plan-review source states plan version is the MFPC version in effect at plan approval; construction-code permit logic remains unresolved. |
| Local enforcement model classified | pass | Local enforcement classified with statutory support. |
| Local amendment rule classified | partial | Not-less-stringent and opt-out rules classified; construction-code filing/registry still unresolved. |
| AHJ confirmation metadata present | fail | AHJ contacts and local jurisdiction inventory not populated. |
| Official-source caveats captured | pass | Caveats recorded for code portal and convenience HTML extraction. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-ms:001 | high | MBCC current editions | Exact MBCC currently adopted editions and amendment packages were not obtained. | Locate MBCC meeting minutes, order, adoption notice, or official code page listing current editions and amendments. | null | null | open |
| issue:usa-ms:002 | high | official statutory text capture | Statutory section text was extracted from accessible Justia HTML and anchored to official code portal; official portal is not line-addressable here. | Capture or archive official LexisNexis/public-access code text for §§17-2-1 through 17-2-11. | null | null | open |
| issue:usa-ms:003 | high | opt-out inventory | State Uniform Construction Code opt-out jurisdictions and coastal §17-2-1 opt-out status are unresolved. | Locate local resolutions/minutes or any central MBCC/MID/State Fire Marshal opt-out records. | null | null | open |
| issue:usa-ms:004 | medium | energy code | State energy-code authority and current IECC status were not verified from Mississippi primary sources. | Check Mississippi Development Authority / Energy & Natural Resources, MBCC, and administrative code sources. | null | null | open |
| issue:usa-ms:005 | medium | local fire code list | MFPC rule requires annual State Fire Marshal examination/listing of local fire prevention codes, but no public list was obtained. | Request or locate State Fire Marshal local fire-code list. | null | null | open |
| issue:usa-ms:006 | medium | accessibility authority | Separate accessibility authority and amendments were not parsed. | Review IBC adoption amendments and state accessibility statutes/rules. | null | null | open |
| issue:usa-ms:007 | medium | elevator/conveyance authority | Elevator/conveyance code authority was not researched. | Review Mississippi Department of Insurance / Labor / elevator-specific statutes and rules. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-ms:code-public-access | src:usa-ms:code-public-access | manual_statute_check | quarterly | amendments to Title 17, Chapter 2 or Title 45, Chapter 11 | 2026-06-25 | active |
| watch:usa-ms:admin-19-7-ch7 | src:usa-ms:admin-19-7-ch7 | pdf_diff | monthly | new Title 19, Part 7, Chapter 7 PDF or amendments | 2026-06-25 | active |
| watch:usa-ms:mbcc-editions | none | agency_page_or_minutes_search | monthly | MBCC current edition/adoption order found or changed | 2026-06-25 | needs_source |
| watch:usa-ms:local-fire-code-list | src:usa-ms:admin-19-7-ch7 | registry_search | quarterly | State Fire Marshal publishes or updates local fire prevention code list | 2026-06-25 | needs_source |
| watch:usa-ms:energy-code | src:usa-ms:supp-doe-energycodes | primary_source_followup | quarterly | Mississippi primary source confirms IECC/current energy code | 2026-06-25 | needs_source |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-24 | Generated baseline draft report | report:usa-ms | none | Codex | Initial stub with unresolved placeholders. |
| 2026-06-25 | Populated Mississippi authority, State Uniform Construction Code framework, local enforcement, MFPC, source registry, QA, and open issues | ahj:usa-ms:mbcc; ahj:usa-ms:sfm; adoption:usa-ms:state-uniform-construction-code-2014; adoption:usa-ms:mfpc-2025; local-enforcement:usa-ms; local-amendment-rule:usa-ms | src:usa-ms:stat-17-2-1; src:usa-ms:stat-17-2-3; src:usa-ms:stat-17-2-4; src:usa-ms:stat-17-2-5; src:usa-ms:stat-17-2-7; src:usa-ms:stat-17-2-9; src:usa-ms:stat-17-2-11; src:usa-ms:admin-19-7-ch7 | ChatGPT | Upgraded status to partially_verified with explicit caveats and unresolved issues. |
