---
state:
  state_id: "US-NE"
  name: "Nebraska"
  abbreviation: "NE"
report:
  report_id: "state-report:usa-ne"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-26"
  last_verified: "2026-06-26"
  reviewed_by: null
risk:
  overall_confidence: 0.72 # 0.00 - 1.00
  risk_flags:
    - "state_amendments_partially_parsed"
    - "fire_code_chapter_1_is_scanned_stamped_pdf"
    - "jurisdiction_boundary_sources_not_selected"
    - "municipality_specific_amendments_not_parsed"
    - "elevator_boiler_model_code_editions_not_fully_parsed"
  open_questions_count: 5

---

# State Building Code Authority Report: Nebraska

## 1. Executive Summary

- **Authority model:** Nebraska uses a hybrid statutory/local model. The Nebraska Legislature creates and updates the state building code by reference. Counties, cities, and villages may enact, administer, and enforce a local building or construction code if they adopt the state building code or a code that conforms generally with it. If a local jurisdiction does not adopt a qualifying code within two years after a state update, the state building code applies in that jurisdiction, except for farm construction or construction for farm purposes. Source IDs: src:usa-ne:71-6403-state-building-code, src:usa-ne:71-6406-local-building-code.

- **Statewide code status:** The core statutory state building code currently includes the 2018 IBC, 2018 IRC, 2018 IEBC, and 2018 UPC, each with Nebraska-specific exclusions or inclusions. The Nebraska Energy Code is based on the 2018 IECC, and the state energy agency page also identifies ASHRAE 90.1-2016 as part of the 2019 update. Electrical wiring standards are governed separately by the State Electrical Board and are based on the 2023 NEC with statutory exceptions that retain 2017 NEC language for specified sections. Source IDs: src:usa-ne:71-6403-state-building-code, src:usa-ne:81-1611-energy-code, src:usa-ne:dwee-energy-codes, src:usa-ne:81-2104-electrical-board, src:usa-ne:state-electrical-2023-adoption.

- **Local enforcement model:** Local governments are the main building/construction-code enforcement layer where they adopt and administer a conforming local code. The State Fire Marshal separately enforces the State Fire Code and conducts plan review unless delegated or local review authority applies. The Department of Water, Energy, and Environment enforces the energy code where a local jurisdiction has not adopted an energy code. Electrical enforcement is handled under the State Electrical Act through the State Electrical Board/State Electrical Division, with local inspection-program details requiring a separate jurisdictional pass. Source IDs: src:usa-ne:71-6406-local-building-code, src:usa-ne:81-503-01-state-fire-code, src:usa-ne:title153-ch21-plan-review, src:usa-ne:sfm-delegated-local-authority, src:usa-ne:dwee-energy-codes, src:usa-ne:81-2104-electrical-board.

- **Local amendment posture:** Local building/construction-code amendments are allowed within statutory conformance boundaries. A local code is not deemed conforming if it uses a prior edition of any component of the state building code or omits required radon-resistant new-construction standards. Local administrative/enforcement amendments must be published separately. Local governments must notify the Department of Water, Energy, and Environment within 30 days if they delete any portion of IBC chapter 13 or IRC chapter 11 from their local code. Source IDs: src:usa-ne:71-6406-local-building-code, src:usa-ne:76-3504-radon.

- **Known transition periods or pending changes:** LB611 (2025) and LB801 (2026), both proposing building/energy code updates, were indefinitely postponed on 2026-04-17. The report therefore treats the 2018 building-code package and 2018 IECC energy code as the current supported statewide baseline. The operative date for the 2021 building-code update is recorded as a derived date that still needs final legal-calendar confirmation before production reliance. Source IDs: src:usa-ne:lb611-2025-status, src:usa-ne:lb801-2026-status, src:usa-ne:lb131-2021-final, src:usa-ne:lb131-2021-status.

- **Production readiness:** limited_use_research_draft. The core authority and adoption fields are source-backed, but state amendments, local AHJ resolution, local amendment registry status, and mechanical/elevator/boiler code-edition details still require a narrower validation pass.

### Key Findings

```yaml
---
key_findings:
- topic: State building code adoption
  finding: The Legislature creates the state building code and adopts 2018 IBC, 2018
    IRC, 2018 IEBC, and 2018 UPC by reference, with listed exclusions.
  confidence: 0.92
  source_ids:
  - src:usa-ne:71-6403-state-building-code
- topic: Local code enforcement
  finding: Counties, cities, and villages may enact, administer, and enforce local
    building or construction codes if they adopt the state code or a code conforming
    generally with it.
  confidence: 0.88
  source_ids:
  - src:usa-ne:71-6406-local-building-code
- topic: Local fallback rule
  finding: If a local jurisdiction fails to adopt a qualifying code within two years
    after a state update, the state building code applies there, except for farm construction
    or construction for farm purposes.
  confidence: 0.84
  source_ids:
  - src:usa-ne:71-6406-local-building-code
- topic: Local amendment limits
  finding: Local amendments are allowed, but prior editions and omission of radon-resistant
    new-construction minimum standards break general conformity.
  confidence: 0.86
  source_ids:
  - src:usa-ne:71-6406-local-building-code
  - src:usa-ne:76-3504-radon
- topic: Energy code
  finding: Nebraska Energy Code is the 2018 IECC; DWEE states the 2018 IECC went into
    effect on 2020-07-01 and identifies ASHRAE 90.1-2016 in the update.
  confidence: 0.88
  source_ids:
  - src:usa-ne:81-1611-energy-code
  - src:usa-ne:dwee-energy-codes
- topic: Electrical code authority
  finding: The State Electrical Board is the specialized electrical authority; current
    statutory standards begin with the 2023 NEC, subject to listed 2017 NEC exceptions.
  confidence: 0.9
  source_ids:
  - src:usa-ne:81-2104-electrical-board
  - src:usa-ne:state-electrical-2023-adoption
- topic: Fire code authority
  finding: The State Fire Marshal adopts and enforces the State Fire Code and reviews
    fire-code plans; Title 153 Chapter 1 is a scanned/stamped rule source.
  confidence: 0.82
  source_ids:
  - src:usa-ne:81-503-01-state-fire-code
  - src:usa-ne:sfm-regulations
  - src:usa-ne:title153-ch1-fire-code
  - src:usa-ne:title153-ch21-plan-review
- topic: Accessibility authority
  finding: The State Fire Marshal adopts accessibility standards consistent with ADA
    and FHA standards and maintains state/local review-authority pathways.
  confidence: 0.76
  source_ids:
  - src:usa-ne:81-5-147-accessibility
  - src:usa-ne:sfm-regulations
  - src:usa-ne:sfm-delegated-local-authority
- topic: Pending statewide updates
  finding: LB611 and LB801 building/energy-code update bills were indefinitely postponed
    on 2026-04-17.
  confidence: 0.86
  source_ids:
  - src:usa-ne:lb611-2025-status
  - src:usa-ne:lb801-2026-status
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-ne:legislature-building-construction-act |
| Authority name | Nebraska Legislature, through the Building Construction Act |
| Authority type | state legislature / statutory code-adopting authority |
| Legal basis | Neb. Rev. Stat. § 71-6403 creates the state building code and adopts model codes by reference; Neb. Rev. Stat. § 71-6406 controls local adoption, enforcement, conformance, and amendment boundaries. |
| Role | Creates and updates the state building code; defines local conformance and fallback rules. |
| Enforcement model | Local jurisdictions may administer/enforce conforming local codes; the state code applies as a statutory fallback when local adoption is not timely, subject to the farm-purpose exception. |
| Source IDs | src:usa-ne:71-6403-state-building-code, src:usa-ne:71-6406-local-building-code |
| Verification status | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-ne:legislature-building-construction-act | Nebraska Legislature | Adopts 2018 IBC into the state building code; local governments administer/enforce conforming codes. | Neb. Rev. Stat. §§ 71-6403, 71-6406 | src:usa-ne:71-6403-state-building-code, src:usa-ne:71-6406-local-building-code | partially_verified |
| Residential | ahj:usa-ne:legislature-building-construction-act | Nebraska Legislature | Adopts 2018 IRC into the state building code with exclusions. | Neb. Rev. Stat. § 71-6403 | src:usa-ne:71-6403-state-building-code | partially_verified |
| Existing Building / Rehabilitation | ahj:usa-ne:legislature-building-construction-act | Nebraska Legislature | Adopts 2018 IEBC into the state building code with exclusion of section 809. | Neb. Rev. Stat. § 71-6403 | src:usa-ne:71-6403-state-building-code | partially_verified |
| Mechanical | ahj:usa-ne:local-building-authorities | Counties, cities, and villages | No statewide IMC-style adoption was verified; local standard-code authority exists within § 71-6406 and municipal/county enabling statutes referenced there. Boiler/elevator mechanical-safety programs are separate. | Neb. Rev. Stat. § 71-6406 | src:usa-ne:71-6406-local-building-code, src:usa-ne:sfm-mechanical-safety | unresolved |
| Plumbing | ahj:usa-ne:legislature-building-construction-act | Nebraska Legislature | Adopts 2018 UPC into the state building code. Local governments may also adopt plumbing codes under authorities referenced in § 71-6406. | Neb. Rev. Stat. §§ 71-6403, 71-6406 | src:usa-ne:71-6403-state-building-code, src:usa-ne:71-6406-local-building-code | partially_verified |
| Fuel Gas | ahj:usa-ne:state-fire-marshal | Nebraska State Fire Marshal | State Fire Code includes NFPA fuel-gas and LP-gas standards; broader construction-code interface needs further parsing. | Neb. Rev. Stat. § 81-503.01; 153 NAC ch. 1 | src:usa-ne:81-503-01-state-fire-code, src:usa-ne:title153-ch1-fire-code | partially_verified |
| Electrical | ahj:usa-ne:state-electrical-board | Nebraska State Electrical Board / State Electrical Division | Adopts/enforces electrical wiring standards under the State Electrical Act. | Neb. Rev. Stat. § 81-2104 | src:usa-ne:81-2104-electrical-board, src:usa-ne:state-electrical-statutes-rules, src:usa-ne:state-electrical-2023-adoption | partially_verified |
| Energy | ahj:usa-ne:dwee-energy | Nebraska Department of Water, Energy, and Environment; Director of Water, Energy, and Environment | State energy-code administration, alternative standards, and fallback enforcement where no local energy code exists. | Neb. Rev. Stat. §§ 81-1611, 81-1618; DWEE Energy Codes page | src:usa-ne:81-1611-energy-code, src:usa-ne:81-1618-local-energy-code, src:usa-ne:dwee-energy-codes | partially_verified |
| Fire - construction references | ahj:usa-ne:state-fire-marshal | Nebraska State Fire Marshal | Reviews plans for compliance with State Fire Code; plan review required before covered construction, alteration, enlargement, repair, improvement, or conversion unless waived/exempt. | Neb. Rev. Stat. § 81-503.01; 153 NAC ch. 21 | src:usa-ne:81-503-01-state-fire-code, src:usa-ne:title153-ch21-plan-review | partially_verified |
| Fire - operational / prevention code | ahj:usa-ne:state-fire-marshal | Nebraska State Fire Marshal | Adopts, promulgates, and enforces State Fire Code. | Neb. Rev. Stat. § 81-503.01; 153 NAC ch. 1 | src:usa-ne:81-503-01-state-fire-code, src:usa-ne:title153-ch1-fire-code | partially_verified |
| Accessibility | ahj:usa-ne:state-fire-marshal-accessibility | Nebraska State Fire Marshal | Adopts public/commercial building and FHA multifamily accessibility standards; reviews accessibility unless local authority applies. | Neb. Rev. Stat. § 81-5,147; Title 156 listing; SFM delegated/local authority page | src:usa-ne:81-5-147-accessibility, src:usa-ne:sfm-regulations, src:usa-ne:sfm-delegated-local-authority | partially_verified |
| Elevator / Conveyance | ahj:usa-ne:state-fire-marshal-mechanical-safety | Nebraska State Fire Marshal, Mechanical Safety | Elevator/conveyance licensing and inspection program; model-code editions require deeper rule parsing. | Title 230 listing; Neb. Rev. Stat. § 81-5,230 | src:usa-ne:sfm-regulations, src:usa-ne:sfm-mechanical-safety, src:usa-ne:sfm-title230-elevator, src:usa-ne:81-5-230-elevator-licensing | partially_verified |
| Boiler / Pressure Vessel | ahj:usa-ne:state-fire-marshal-mechanical-safety | Nebraska State Fire Marshal, Mechanical Safety | Boiler inspection program; current edition details require deeper rule parsing. | Title 229 listing; Boiler Inspection Act references | src:usa-ne:sfm-regulations, src:usa-ne:sfm-mechanical-safety, src:usa-ne:sfm-title229-boiler | partially_verified |

### 2.3 Authority Hierarchy Notes

The statewide building-code baseline is statutory rather than housed in a building-code board. Local governments are key enforcement bodies for building and construction codes, but they must stay within the conformance framework in § 71-6406. Specialized statewide programs sit outside that general local-building-code layer: the State Fire Marshal controls the State Fire Code, plan review, accessibility review where not locally delegated, boiler inspection, and elevator/conveyance safety; the State Electrical Board controls statewide electrical standards; and DWEE administers the Nebraska Energy Code and fallback energy enforcement.

The State Fire Marshal delegated-authority page states that the State Fire Marshal's jurisdiction encompasses the entire state except areas granted delegated authority, but the same page also says the State Fire Marshal retains the ability to overrule local delegated personnel. The page's narrative says delegated authority is held in “nine cities,” while the displayed list contains more entries. Treat that count/list discrepancy as an agency-page caveat until reconciled with the underlying delegation documents. Source IDs: src:usa-ne:sfm-delegated-local-authority.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-ne:001 | ahj:usa-ne:legislature-building-construction-act | creates_and_adopts | state_building_code:usa-ne | src:usa-ne:71-6403-state-building-code | partially_verified |
| edge:usa-ne:002 | ahj:usa-ne:legislature-building-construction-act | authorizes_enforcement_by | counties_cities_villages_adopting_conforming_code | src:usa-ne:71-6406-local-building-code | partially_verified |
| edge:usa-ne:003 | ahj:usa-ne:legislature-building-construction-act | applies_as_fallback_if | local_jurisdiction_fails_to_adopt_within_two_years_after_update | src:usa-ne:71-6406-local-building-code | partially_verified |
| edge:usa-ne:004 | ahj:usa-ne:state-fire-marshal | adopts_and_enforces | state_fire_code | src:usa-ne:81-503-01-state-fire-code, src:usa-ne:title153-ch1-fire-code | partially_verified |
| edge:usa-ne:005 | ahj:usa-ne:state-fire-marshal | reviews_plans_for | state_fire_code_compliance | src:usa-ne:81-503-01-state-fire-code, src:usa-ne:title153-ch21-plan-review | partially_verified |
| edge:usa-ne:006 | ahj:usa-ne:state-fire-marshal | delegates_some_authority_to | local_fire_prevention_personnel_in_listed_jurisdictions | src:usa-ne:sfm-delegated-local-authority | partially_verified |
| edge:usa-ne:007 | ahj:usa-ne:state-fire-marshal-accessibility | local_review_substitution_for | local_authority_cities_and_counties_conducting_NAG_review | src:usa-ne:sfm-delegated-local-authority | partially_verified |
| edge:usa-ne:008 | ahj:usa-ne:dwee-energy | fallback_enforces | Nebraska_Energy_Code_where_no_local_energy_code | src:usa-ne:dwee-energy-codes | partially_verified |
| edge:usa-ne:009 | ahj:usa-ne:state-electrical-board | adopts_and_enforces | statewide_electrical_wiring_standards | src:usa-ne:81-2104-electrical-board | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | Nebraska state building code | International Building Code | 2018 | adopted | 2021-05-25 | 2021-05-25 | 2021-08-28 | 2023-08-28 | LB131 was approved with emergency clause and specified operative timing for listed sections; local jurisdictions have two years after a state update to adopt the state code or a code conforming generally with it. Operative/mandatory dates are derived and should be validated against the official legislative calendar before verified status. | src:usa-ne:71-6403-state-building-code, src:usa-ne:71-6406-local-building-code, src:usa-ne:lb131-2021-final, src:usa-ne:lb131-2021-status |
| Residential | Nebraska state building code | International Residential Code | 2018 | adopted | 2021-05-25 | 2021-05-25 | 2021-08-28 | 2023-08-28 | Same as building; IRC section R313 and chapters 25-33 are excluded from the statutory state code. | src:usa-ne:71-6403-state-building-code, src:usa-ne:71-6406-local-building-code, src:usa-ne:lb131-2021-final, src:usa-ne:lb131-2021-status |
| Existing Building / Rehabilitation | Nebraska state building code | International Existing Building Code | 2018 | adopted | 2021-05-25 | 2021-05-25 | 2021-08-28 | 2023-08-28 | Same as building; IEBC section 809 is excluded from the statutory state code. | src:usa-ne:71-6403-state-building-code, src:usa-ne:71-6406-local-building-code, src:usa-ne:lb131-2021-final, src:usa-ne:lb131-2021-status |
| Mechanical | Local building/construction codes and specialized SFM mechanical-safety programs | unresolved | unresolved | unresolved | null | null | null | null | No statewide mechanical model-code adoption was verified. Boiler and elevator/conveyance safety are separately administered by the State Fire Marshal. | src:usa-ne:71-6406-local-building-code, src:usa-ne:sfm-mechanical-safety |
| Plumbing | Nebraska state building code | Uniform Plumbing Code | 2018 | adopted | 2021-05-25 | 2021-05-25 | 2021-08-28 | 2023-08-28 | UPC is part of state building code; local plumbing-code authority is also recognized under § 71-6406 through referenced local enabling statutes. | src:usa-ne:71-6403-state-building-code, src:usa-ne:71-6406-local-building-code, src:usa-ne:lb131-2021-final, src:usa-ne:lb131-2021-status |
| Fuel Gas | Nebraska State Fire Code | NFPA 54, National Fuel Gas Code; NFPA 58, Liquefied Petroleum Gas Code | NFPA 54-2012; NFPA 58-2011 | partially_verified | 2021-02-23 | 2021-02-23 | 2021-02-23 | null | Title 153 Chapter 1 is a stamped/scanned adopted-NFPA-code listing. Broader building-code fuel-gas interface is unresolved. | src:usa-ne:title153-ch1-fire-code, src:usa-ne:81-503-01-state-fire-code |
| Electrical | State Electrical Act / State Electrical Board standards | NFPA 70, National Electrical Code | 2023, with specified 2017 NEC exceptions | adopted | 2024-04-19 | 2024-08-01 | 2024-08-01 | 2024-08-01 | Electrical permits received by the State Electrical Division on or after 2024-08-01 use the 2023 NEC; existing permits remain under the 2017 NEC. | src:usa-ne:81-2104-electrical-board, src:usa-ne:state-electrical-2023-adoption |
| Energy | Nebraska Energy Code | International Energy Conservation Code; ASHRAE 90.1 identified by DWEE | 2018 IECC; ASHRAE 90.1-2016 | adopted | 2019-05-08 | 2020-07-01 | 2020-07-01 | 2020-07-01 | DWEE states the 2018 IECC took effect 2020-07-01; if no local energy code is adopted, DWEE enforces. | src:usa-ne:81-1611-energy-code, src:usa-ne:dwee-energy-codes, src:usa-ne:81-1618-local-energy-code |
| Fire - construction references | Nebraska State Fire Code / SFM plan review | NFPA 1; NFPA 101; associated NFPA pamphlets | NFPA 1-2012; NFPA 101-2012 by statute; Title 153 ch. 1 stamped NFPA listing | adopted | 2021-02-23 | 2021-02-23 | 2021-02-23 | null | SFM plan review is required before covered construction/remodeling unless exempt or waived; plan approval expires if work is not commenced within 180 days. | src:usa-ne:81-503-01-state-fire-code, src:usa-ne:title153-ch1-fire-code, src:usa-ne:title153-ch21-plan-review |
| Fire - operational / prevention code | Nebraska State Fire Code | NFPA 1; NFPA 101; associated NFPA pamphlets | NFPA 1-2012; NFPA 101-2012 by statute; Title 153 ch. 1 stamped NFPA listing | adopted | 2021-02-23 | 2021-02-23 | 2021-02-23 | null | State Fire Marshal enforces State Fire Code through inspections, compliance, orders, and plan review. | src:usa-ne:81-503-01-state-fire-code, src:usa-ne:title153-ch1-fire-code |
| Accessibility | Nebraska Accessibility Guidelines / accessibility standards | ADA standards and FHA standards, as incorporated by state rule | current federal standards, edition not parsed | partially_verified | null | null | null | null | SFM adopts accessibility standards; local-authority jurisdictions may conduct NAG review in lieu of SFM, subject to exceptions. | src:usa-ne:81-5-147-accessibility, src:usa-ne:sfm-regulations, src:usa-ne:sfm-delegated-local-authority |
| Elevator / Conveyance | Nebraska conveyance safety program | Title 230 conveyance safety rules | edition unresolved | partially_verified | null | null | null | null | SFM Mechanical Safety administers elevator/conveyance inspections and licensing; full code edition parsing remains open. | src:usa-ne:sfm-regulations, src:usa-ne:sfm-mechanical-safety, src:usa-ne:sfm-title230-elevator, src:usa-ne:81-5-230-elevator-licensing |
| Boiler / Pressure Vessel | Nebraska Boiler Safety Code | Title 229 boiler safety rules | edition unresolved | partially_verified | null | null | null | null | SFM Mechanical Safety administers boiler inspection; full code edition parsing remains open. | src:usa-ne:sfm-regulations, src:usa-ne:sfm-mechanical-safety, src:usa-ne:sfm-title229-boiler |

### 3.2 Adoption Records

```yaml
adoption_records:
  - adoption_id: adoption:usa-ne:building:2018-ibc
    code_family: Building
    state_code_name: Nebraska state building code
    base_model_code: International Building Code
    edition: "2018"
    adoption_date: "2021-05-25"
    effective_date: "2021-05-25"
    operative_date: "2021-08-28"
    mandatory_date: "2023-08-28"
    date_confidence: 0.62
    notes: "Date sequence uses LB131 approval/emergency clause plus operative timing and § 71-6406 two-year local-update rule; verify exact operative date before verified status."
    source_ids:
      - src:usa-ne:71-6403-state-building-code
      - src:usa-ne:71-6406-local-building-code
      - src:usa-ne:lb131-2021-final
      - src:usa-ne:lb131-2021-status

  - adoption_id: adoption:usa-ne:residential:2018-irc
    code_family: Residential
    state_code_name: Nebraska state building code
    base_model_code: International Residential Code
    edition: "2018"
    exclusions: ["section R313", "chapters 25 through 33"]
    adoption_date: "2021-05-25"
    effective_date: "2021-05-25"
    operative_date: "2021-08-28"
    mandatory_date: "2023-08-28"
    date_confidence: 0.62
    source_ids:
      - src:usa-ne:71-6403-state-building-code
      - src:usa-ne:71-6406-local-building-code
      - src:usa-ne:lb131-2021-final
      - src:usa-ne:lb131-2021-status

  - adoption_id: adoption:usa-ne:existing-building:2018-iebc
    code_family: Existing Building / Rehabilitation
    state_code_name: Nebraska state building code
    base_model_code: International Existing Building Code
    edition: "2018"
    exclusions: ["section 809"]
    adoption_date: "2021-05-25"
    effective_date: "2021-05-25"
    operative_date: "2021-08-28"
    mandatory_date: "2023-08-28"
    date_confidence: 0.62
    source_ids:
      - src:usa-ne:71-6403-state-building-code
      - src:usa-ne:71-6406-local-building-code
      - src:usa-ne:lb131-2021-final
      - src:usa-ne:lb131-2021-status

  - adoption_id: adoption:usa-ne:plumbing:2018-upc
    code_family: Plumbing
    state_code_name: Nebraska state building code
    base_model_code: Uniform Plumbing Code
    edition: "2018"
    adoption_date: "2021-05-25"
    effective_date: "2021-05-25"
    operative_date: "2021-08-28"
    mandatory_date: "2023-08-28"
    date_confidence: 0.62
    source_ids:
      - src:usa-ne:71-6403-state-building-code
      - src:usa-ne:71-6406-local-building-code
      - src:usa-ne:lb131-2021-final
      - src:usa-ne:lb131-2021-status

  - adoption_id: adoption:usa-ne:energy:2018-iecc
    code_family: Energy
    state_code_name: Nebraska Energy Code
    base_model_code: International Energy Conservation Code
    edition: "2018"
    companion_standard: "ASHRAE 90.1-2016 identified by DWEE"
    adoption_date: "2019-05-08"
    effective_date: "2020-07-01"
    operative_date: "2020-07-01"
    mandatory_date: "2020-07-01"
    date_confidence: 0.86
    source_ids:
      - src:usa-ne:81-1611-energy-code
      - src:usa-ne:dwee-energy-codes

  - adoption_id: adoption:usa-ne:electrical:2023-nec
    code_family: Electrical
    state_code_name: State Electrical Act / State Electrical Board standards
    base_model_code: NFPA 70 National Electrical Code
    edition: "2023, with statutory 2017 NEC exceptions"
    adoption_date: "2024-04-19"
    effective_date: "2024-08-01"
    operative_date: "2024-08-01"
    mandatory_date: "2024-08-01"
    date_confidence: 0.88
    source_ids:
      - src:usa-ne:81-2104-electrical-board
      - src:usa-ne:state-electrical-2023-adoption

  - adoption_id: adoption:usa-ne:fire:title153-ch1
    code_family: Fire - operational / prevention code
    state_code_name: Nebraska State Fire Code
    base_model_code: NFPA 1, NFPA 101, and associated NFPA standards
    edition: "NFPA 1-2012 and NFPA 101-2012 by statute; Title 153 adopted listing includes NFPA 1-2012"
    adoption_date: "2021-02-23"
    effective_date: "2021-02-23"
    operative_date: "2021-02-23"
    mandatory_date: null
    date_confidence: 0.70
    source_ids:
      - src:usa-ne:81-503-01-state-fire-code
      - src:usa-ne:title153-ch1-fire-code
```

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

Nebraska separates statutory adoption from local implementation. The state building-code update enacted through LB131 was approved by the Governor on 2021-05-25 and included an emergency clause, but the building-code sections were made operative three calendar months after adjournment of the legislative session. The report records 2021-08-28 as the derived operative date and 2023-08-28 as the derived two-year local backstop date; both dates should be confirmed against official calendar interpretation before the report is advanced beyond partially verified. Source IDs: src:usa-ne:lb131-2021-status, src:usa-ne:lb131-2021-final, src:usa-ne:71-6406-local-building-code.

The energy-code date rule is clearer from the agency page: the 2018 IECC went into effect on 2020-07-01, and DWEE enforces the energy code if a local jurisdiction has not adopted an energy code. Electrical transition is permit-based: permits received by the State Electrical Division on or after 2024-08-01 are on the 2023 NEC, while existing permits remain on the 2017 NEC. Source IDs: src:usa-ne:dwee-energy-codes, src:usa-ne:state-electrical-2023-adoption.

Fire-code plan review has its own timing: covered work must submit plans before construction, alteration, enlargement, repair, improvement, or conversion unless exempt or waived; approval expires if work has not commenced within 180 days. A State Fire Marshal certificate of occupancy is required before covered use or occupancy. Source IDs: src:usa-ne:title153-ch21-plan-review, src:usa-ne:title153-ch22-co.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-ne:building-lb131-operative | 2018 IBC/IRC/IEBC/UPC update | operative_date | 2021-08-28 derived | LB131 sections made operative three calendar months after legislative adjournment; exact calendar conversion requires final legal-calendar validation. | prior editions not conforming under § 71-6406 | src:usa-ne:lb131-2021-final, src:usa-ne:lb131-2021-status, src:usa-ne:71-6406-local-building-code | partially_verified |
| date-rule:usa-ne:local-two-year-backstop | Local building/construction code adoption | mandatory_backstop | two years after state building-code update; 2023-08-28 derived for LB131 update | County/city/village does not adopt state code or a generally conforming code within two years after state update. | no, prior editions of state-code components break conformity | src:usa-ne:71-6406-local-building-code | partially_verified |
| date-rule:usa-ne:energy-2018-iecc | Nebraska Energy Code | effective / mandatory | 2020-07-01 | 2018 IECC became Nebraska Energy Code; DWEE fallback enforcement applies where no local energy code exists. | Local energy code may be equivalent if it does not increase energy consumption and is consistent with statutory intent. | src:usa-ne:81-1611-energy-code, src:usa-ne:dwee-energy-codes, src:usa-ne:81-1618-local-energy-code | partially_verified |
| date-rule:usa-ne:electrical-2023-nec | Electrical permits | permit_received_date | 2024-08-01 | Electrical permits received by the State Electrical Division on or after this date use 2023 NEC; existing permits remain on 2017 NEC. | existing permits only | src:usa-ne:state-electrical-2023-adoption, src:usa-ne:81-2104-electrical-board | partially_verified |
| date-rule:usa-ne:sfm-plan-review-before-work | Fire-code plan review | preconstruction_submittal | before covered work begins | Covered erection, construction, enlargement, alteration, repair, improvement, or conversion. | no | src:usa-ne:title153-ch21-plan-review | partially_verified |
| date-rule:usa-ne:sfm-plan-approval-expiration | Fire-code plan review | expiration | 180 days | Work authorized by SFM plan approval is not commenced within 180 days unless extended by SFM. | no | src:usa-ne:title153-ch21-plan-review | partially_verified |
| date-rule:usa-ne:sfm-co-before-occupancy | Fire-code certificate of occupancy | occupancy_trigger | before use or occupancy | No covered building or structure may be used or occupied, and no covered change in structure/portion made, until SFM issues certificate of occupancy. | no | src:usa-ne:title153-ch22-co | partially_verified |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building / Energy | 2021 IECC / 2021 building-code update proposal in LB611 | 2025-01-22 | null | null | null | null | closed_failed | src:usa-ne:lb611-2025-status | LB611 was indefinitely postponed on 2026-04-17. |
| Building / Energy | 2024 IBC/energy chapter update proposal in LB801 | 2026-01-07 | null | null | null | null | closed_failed | src:usa-ne:lb801-2026-status | LB801 was indefinitely postponed on 2026-04-17. |
| Fire / plans administration | SFM plan review fee update | 2025-09-01 | null | 2025-09-01 | 2025-09-01 | 2025-09-01 | monitor_admin_change | src:usa-ne:sfm-regulations, src:usa-ne:title153-ch21-plan-review | Administrative fee update, not a model-code edition update. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-ne:farm-exception | Building | Farm construction / construction for farm purposes | State building code fallback after local non-adoption | State code fallback after the two-year local adoption window does not apply to construction on a farm or for farm purposes. | src:usa-ne:71-6406-local-building-code | partially_verified |
| applicability-rule:usa-ne:radon-new-construction | Building / Residential | New construction intended to be regularly occupied by people | Built after 2019-09-01 | Must use radon-resistant new construction unless statutory exception applies; minimum standards are part of state building code through § 71-6403. | src:usa-ne:76-3504-radon, src:usa-ne:71-6403-state-building-code | partially_verified |
| applicability-rule:usa-ne:state-building-energy | Energy | New state buildings and state-owned building components | State-owned or state-funded project, as described in §§ 72-804 and 72-805 | New state building and listed state-owned building components must meet or exceed 2018 IECC; state-funded building review is handled by DWEE unless reviewed by a local jurisdiction enforcing a local code that includes 2018 IECC requirements. | src:usa-ne:72-804-state-building-energy, src:usa-ne:72-805-state-funded-energy | partially_verified |
| applicability-rule:usa-ne:sfm-plan-review-small-apartment-exemption | Fire / plan review | Apartment buildings with fewer than three dwelling units | Building-plan review under 153 NAC ch. 21 | Building plans review is not required for apartment buildings with fewer than three dwelling units. | src:usa-ne:title153-ch21-plan-review | partially_verified |
| applicability-rule:usa-ne:sfm-delegated-state-owned-exception | Fire / accessibility | State-owned properties and federally certified health care facilities | Project located in delegated/local authority jurisdiction | SFM page states state-owned properties and CMS-certified facilities are exempt from delegated/local authority jurisdiction and plans/NAG review should be submitted to SFM. | src:usa-ne:sfm-delegated-local-authority | partially_verified |
| applicability-rule:usa-ne:nag-religious-exemption | Accessibility | Religious entities | NAG review | SFM local-authority page states all religious entities are exempted from NAG reviews. | src:usa-ne:sfm-delegated-local-authority | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** Nebraska's state amendments appear in several publication paths rather than one unified amendment volume. Core building-code amendments are embedded directly in § 71-6403 as exclusions and special rules. Local amendment boundaries are in § 71-6406. Electrical amendments/exceptions are in § 81-2104. Fire-code amendments are in State Fire Marshal Title 153 rules, including a scanned/stamped Chapter 1 NFPA adoption listing. Energy-code alternative standards and local energy-code equivalence/waiver rules are in §§ 81-1611 and 81-1618.

**Where amendments are published:** Nebraska Revised Statutes, State Fire Marshal regulations, State Electrical Board rules/pages, and local ordinances or resolutions. Local administrative/enforcement amendments must be published separately from the local building or construction code, and local governments adopting standard codes must keep at least one copy available in the clerk's office.

**Amendment parsing status:** partial. The high-impact statutory amendments have been captured, but a complete line-by-line amendment extraction from State Fire Marshal, State Electrical Board, Title 156 accessibility, Title 229 boiler, Title 230 elevator, and local ordinances remains open.

### 5.2 State Amendment Sources

| Source ID | Amendment Area | Publication Path | Parsing Status | Notes |
| --- | --- | --- | --- | --- |
| src:usa-ne:71-6403-state-building-code | Building, residential, existing building, plumbing | Nebraska Revised Statutes | high_impact_fields_parsed | Statutory adoption/exclusions parsed. |
| src:usa-ne:71-6406-local-building-code | Local amendments and conformance | Nebraska Revised Statutes | high_impact_fields_parsed | Local amendment scope and nonconformance triggers parsed. |
| src:usa-ne:76-3504-radon | Radon-resistant new construction | Nebraska Revised Statutes | high_impact_fields_parsed | Statutory minimum standards parsed at summary level only. |
| src:usa-ne:81-2104-electrical-board | Electrical | Nebraska Revised Statutes | high_impact_fields_parsed | 2023 NEC adoption and 2017 NEC exception sections parsed. |
| src:usa-ne:title153-ch1-fire-code | Fire | State Fire Marshal Title 153 Chapter 1 PDF | partial | Scanned/stamped PDF; NFPA adoption list needs a structured extraction pass. |
| src:usa-ne:title153-ch21-plan-review | Fire plan review | State Fire Marshal Title 153 Chapter 21 PDF | partial | Plan-review triggers and 180-day expiration parsed. |
| src:usa-ne:title153-ch22-co | Fire certificate of occupancy | State Fire Marshal Title 153 Chapter 22 PDF | partial | CO trigger parsed. |
| src:usa-ne:81-1611-energy-code | Energy | Nebraska Revised Statutes | high_impact_fields_parsed | 2018 IECC adoption and alternative-standards authority parsed. |
| src:usa-ne:81-1618-local-energy-code | Energy local code / waiver | Nebraska Revised Statutes | high_impact_fields_parsed | Local equivalence/waiver procedure parsed. |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Amendment Summary | Impact | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| amend:usa-ne:ibc:101-4-3-ch29 | Building | State building code adopts the 2018 IBC except section 101.4.3 and chapter 29. | Excludes plumbing-systems reference section and plumbing-fixture chapter from state IBC adoption; UPC is separately adopted. | src:usa-ne:71-6403-state-building-code | partially_verified |
| amend:usa-ne:ibc:care-childcare | Building | IBC section 305.2.3 applies to a facility with twelve or fewer children; section 310.4.1 applies to a care facility for twelve or fewer persons. | Special Nebraska statutory occupancy/care-facility rules. | src:usa-ne:71-6403-state-building-code | partially_verified |
| amend:usa-ne:irc:r313-ch25-33 | Residential | State building code adopts 2018 IRC except section R313 and chapters 25 through 33. | Excludes IRC sprinkler and plumbing-related chapters from state IRC adoption. | src:usa-ne:71-6403-state-building-code | partially_verified |
| amend:usa-ne:iebc:809 | Existing Building / Rehabilitation | State building code adopts 2018 IEBC except section 809. | Excludes IEBC section 809 from state adoption. | src:usa-ne:71-6403-state-building-code | partially_verified |
| amend:usa-ne:upc:2018 | Plumbing | State building code adopts the 2018 Uniform Plumbing Code as an ANSI American National Standard. | Establishes UPC rather than IPC as statewide statutory plumbing code component. | src:usa-ne:71-6403-state-building-code | partially_verified |
| amend:usa-ne:radon:minimums | Building / Residential | Minimum standards for radon-resistant new construction adopted under § 76-3504 are part of the state building code. | New occupiable construction must include statutory radon-resistant construction features unless excepted. | src:usa-ne:71-6403-state-building-code, src:usa-ne:76-3504-radon | partially_verified |
| amend:usa-ne:local:no-prior-edition | Local amendments | Local code is not deemed generally conforming if it includes a prior edition of any component or combination of components of the state building code. | Prevents local backsliding to older editions for state-code components. | src:usa-ne:71-6406-local-building-code | partially_verified |
| amend:usa-ne:local:dwee-notice | Local amendments / energy chapters | Counties, cities, and villages must notify DWEE within 30 days if they delete any part of IBC chapter 13 or IRC chapter 11. | Creates state notice hook for local energy-related deletions. | src:usa-ne:71-6406-local-building-code | partially_verified |
| amend:usa-ne:electrical:2017-exceptions | Electrical | 2023 NEC is the beginning reference, except 2017 NEC minimum standards continue for sections 210.8(A), 210.8(A)(3), 210.8(A)(5), 230.67(A), and 230.85. | Maintains older NEC provisions for selected GFCI/service equipment requirements. | src:usa-ne:81-2104-electrical-board | partially_verified |
| amend:usa-ne:fire:nfpa1-ne-amendments | Fire | Title 153 Chapter 1 lists “Fire Prevention Code (with NE Amendments 11-25-19)” as NFPA 1-2012. | Confirms State Fire Code has Nebraska amendments to NFPA 1. Full amendment text requires separate extraction. | src:usa-ne:title153-ch1-fire-code | partially_verified |
| amend:usa-ne:energy:alternative-standards | Energy | Director of Water, Energy, and Environment may adopt equivalent alternative standards if they do not result in energy consumption greater than strict application of Nebraska Energy Code. | Provides statewide alternative-compliance pathway. | src:usa-ne:81-1611-energy-code | partially_verified |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-ne"
  model: "hybrid_state_fallback_local_primary"
  enforcing_entities:
    - "county, city, or village adopting the state building code or a code conforming generally with it"
    - "state building code as fallback where local jurisdiction fails to adopt a qualifying code within two years after a state update, subject to farm-purpose exception"
    - "Nebraska State Fire Marshal for State Fire Code enforcement, plan review, inspections, compliance, orders, and certificate-of-occupancy functions unless delegated or local authority applies"
    - "local fire prevention personnel in State Fire Marshal delegated-authority jurisdictions, subject to SFM retained override"
    - "Nebraska Department of Water, Energy, and Environment where no local energy code has been adopted"
    - "Nebraska State Electrical Board / State Electrical Division for State Electrical Act standards and enforcement"
  required_officials:
    - "local code authority or building official where local code is adopted"
    - "county, city, or village clerk maintaining adopted standard-code copies"
    - "State Fire Marshal or delegated/local fire/accessibility authority"
    - "State Electrical Board / State Electrical Division inspectors or approved local inspection program, as applicable"
    - "DWEE energy-code reviewer/enforcement role where applicable"
  state_reserved_activities:
    - "State Fire Code adoption, enforcement, plan review, and SFM override of delegated local fire-prevention personnel"
    - "state-owned properties and CMS-certified health care facilities exempt from delegated/local SFM review authority per SFM page"
    - "Nebraska Energy Code alternative standards and fallback enforcement where no local energy code exists"
    - "statewide electrical wiring standards under the State Electrical Act"
  source_ids:
    - "src:usa-ne:71-6406-local-building-code"
    - "src:usa-ne:81-503-01-state-fire-code"
    - "src:usa-ne:title153-ch21-plan-review"
    - "src:usa-ne:title153-ch22-co"
    - "src:usa-ne:sfm-delegated-local-authority"
    - "src:usa-ne:dwee-energy-codes"
    - "src:usa-ne:81-2104-electrical-board"
  verification_status: "partially_verified"
  confidence: 0.78
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-ne"
  model: "local_amendments_allowed_with_statutory_conformance_boundaries"
  applies_to_code_families:
    - "building"
    - "residential"
    - "existing_building"
    - "plumbing"
    - "local standard codes referenced by § 71-6406"
    - "energy, subject to § 81-1618 and § 71-6406 notice rules"
  approval_required: "not generally verified for all local building amendments; local energy-code waivers require analysis submitted to DWEE and DWEE review/response before local enforcement proceeds"
  approving_authority_id: "ahj:usa-ne:dwee-energy for local energy-code waiver review; otherwise unresolved for local building/construction amendments"
  filing_required: "local standard-code copies must be kept in the county/city/village clerk office; local deletion of IBC chapter 13 or IRC chapter 11 requires DWEE notice within 30 days"
  registry_exists: "no statewide local building-amendment registry verified"
  registry_source_ids: []
  legal_basis_source_ids:
    - "src:usa-ne:71-6406-local-building-code"
    - "src:usa-ne:81-1618-local-energy-code"
  verification_status: "partially_verified"
  confidence: 0.79
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Nebraska separates local enforcement authority from local amendment authority. A county, city, or village may administer and enforce a local building or construction code if the code is the state building code or generally conforms with it. Local amendments are permitted for cost, safety, durability, efficiency, best-practice, or special-local-condition reasons, but local governments may not adopt or enforce a local building or construction code outside § 71-6406. Local governments may have enforcement apparatus even when the state building code is the legally applicable code; § 71-6406 says the applicable code remains legally applicable regardless of whether local administration/enforcement provisions have been adopted. Source IDs: src:usa-ne:71-6406-local-building-code.

Energy is separate: local energy codes are allowed if they do not result in greater energy consumption than strict Nebraska Energy Code application and are reasonably consistent with statutory intent. Local waivers of a specific Nebraska Energy Code requirement require a local analysis submitted to DWEE and a DWEE review/response process. Source IDs: src:usa-ne:81-1618-local-energy-code.

### 6.4 Known Local Amendment Registries

No statewide local building-amendment registry was verified. The following state-maintained sources are related to local responsibility but should not be treated as comprehensive amendment registries:

| Registry / List | Coverage | Source IDs | Status | Notes |
| --- | --- | --- | --- | --- |
| SFM Delegated Authority Cities | Local fire-prevention delegated authority areas | src:usa-ne:sfm-delegated-local-authority | partially_verified | Agency page warns jurisdiction lists can change and contains a count/list discrepancy. |
| SFM Local Authority Cities and Counties | Local accessibility/NAG review areas | src:usa-ne:sfm-delegated-local-authority | partially_verified | Agency page states listed local authorities conduct accessibility review in lieu of SFM, subject to stated exceptions. |
| DWEE local energy-code notifications | Local deletion of IBC chapter 13 or IRC chapter 11 | src:usa-ne:71-6406-local-building-code | unresolved | Statute creates a notification requirement; no public registry was verified in this pass. |

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendments were parsed for this report. Because § 71-6406 allows local standards and requires local copies to be held at the clerk's office, municipality-specific coverage should be built from local ordinances and clerk-held code copies rather than inferred from the statewide report.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: designed_not_implemented

Jurisdiction stack:

```text
Address
  -> State of Nebraska
  -> County
  -> Municipality / unincorporated county
  -> Local building or construction code status under § 71-6406
  -> Applicable local amendments and clerk-held code copies
  -> State building code fallback if no timely conforming local code
  -> Fire AHJ: State Fire Marshal or SFM delegated authority city / campus
  -> Accessibility/NAG review: State Fire Marshal or SFM-listed local authority city/county
  -> Energy AHJ: local energy-code jurisdiction or DWEE fallback enforcement
  -> Electrical AHJ: State Electrical Division / State Electrical Board framework and any approved municipal/county inspection program
  -> Boiler/elevator/conveyance AHJ: SFM Mechanical Safety where applicable
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | unresolved | none | statewide | unresolved | pending |
| County | unresolved | none | statewide | unresolved | pending |
| Municipality | unresolved | none | statewide | unresolved | pending |
| Fire delegated authority | SFM delegated-authority page | src:usa-ne:sfm-delegated-local-authority | listed cities/campus entries and radii where shown | agency page updates; exact cadence unresolved | partially_verified |
| Accessibility local authority | SFM local-authority page | src:usa-ne:sfm-delegated-local-authority | listed cities and counties | agency page warns local authority can change without notice | partially_verified |
| Special districts | unresolved | none | statewide | unresolved | pending |

### 7.3 AHJ Contact Data

| AHJ Data Type | Status | Source IDs | Notes |
| --- | --- | --- | --- |
| State Fire Marshal delegated authority contact list | partially_identified | src:usa-ne:sfm-delegated-local-authority | Page links a delegated-authority contact-information document; the linked contact document was not parsed into structured contacts. |
| State Fire Marshal inspector map | unresolved | src:usa-ne:sfm-regulations | Regulations page links maps for scheduling inspections by division and county; the map was not parsed. |
| State Electrical inspectors / municipal and county electrical inspection programs | unresolved | src:usa-ne:state-electrical-statutes-rules | State Electrical site navigation identifies municipal/county inspection programs, but details were not parsed. |
| Local building officials | unresolved | none | Requires local jurisdiction source collection. |
| DWEE energy-code contact | partially_identified | src:usa-ne:dwee-energy-codes | DWEE page lists an energy-code contact, but a structured AHJ contact record was not created. |

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title / Citation | Source Type | URL | Version / Date Captured | Key Fields Supported | Caveat |
| --- | --- | --- | --- | --- | --- | --- |
| src:usa-ne:71-6403-state-building-code | Neb. Rev. Stat. § 71-6403, State building code; adopted; amendments | statute | https://nebraskalegislature.gov/laws/statutes.php?statute=71-6403 | accessed 2026-06-26 | State building code creation; 2018 IBC, IRC, IEBC, UPC; radon standards included | Official legislature HTML; current text can change after access date. |
| src:usa-ne:71-6406-local-building-code | Neb. Rev. Stat. § 71-6406, County, city, or village; building code; adopt; amend; enforce; copy; fees | statute | https://nebraskalegislature.gov/laws/statutes.php?statute=71-6406 | accessed 2026-06-26 | Local adoption/enforcement; two-year fallback; local amendment conformance; clerk-copy rule | Official legislature HTML; current text can change after access date. |
| src:usa-ne:lb131-2021-status | LB131 bill page, 107th Legislature | bill history | https://nebraskalegislature.gov/bills/view_bill.php?DocumentID=43499 | accessed 2026-06-26 | Governor approval date; final reading/emergency-clause status; bill history | Bill page is official; operative date is derived from final bill text plus legislative calendar. |
| src:usa-ne:lb131-2021-final | Legislative Bill 131 Final Reading PDF, 2021 | final bill PDF | https://nebraskalegislature.gov/FloorDocs/107/PDF/Final/LB131.pdf | accessed 2026-06-26 | Operative timing; emergency clause; repeal/original-section update path | PDF parsed by web text; exact section mapping should be verified before verified status. |
| src:usa-ne:81-1611-energy-code | Neb. Rev. Stat. § 81-1611, Nebraska Energy Code; adoption; alternative standards | statute | https://nebraskalegislature.gov/laws/statutes.php?statute=81-1611 | accessed 2026-06-26 | 2018 IECC adoption; alternative-standards authority | Official legislature HTML. |
| src:usa-ne:dwee-energy-codes | Nebraska Department of Water, Energy, and Environment, Energy Codes page | agency page | https://dwee.nebraska.gov/state-energy-information/energy-codes | accessed 2026-06-26 | 2018 IECC effective 2020-07-01; ASHRAE 90.1-2016 reference; DWEE fallback enforcement | Agency page is official; page text should be monitored for updates. |
| src:usa-ne:81-1618-local-energy-code | Neb. Rev. Stat. § 81-1618, Local energy code; fees; waiver; procedure | statute | https://nebraskalegislature.gov/laws/statutes.php?statute=81-1618 | accessed 2026-06-26 | Local energy code equivalence; local waiver submission/review procedure | Official legislature HTML. |
| src:usa-ne:72-804-state-building-energy | Neb. Rev. Stat. § 72-804, New state building; code requirements | statute | https://nebraskalegislature.gov/laws/statutes.php?statute=72-804 | accessed 2026-06-26 | New state building and state-owned component energy requirements | Official legislature HTML; source displays 2026 effective date note for unrelated amendment history. |
| src:usa-ne:72-805-state-funded-energy | Neb. Rev. Stat. § 72-805, Buildings constructed with state funds; code requirements; exceptions | statute | https://nebraskalegislature.gov/laws/statutes.php?statute=72-805 | accessed 2026-06-26 | State-funded building IECC review and exception framework | Official legislature HTML; source displays 2026 effective date note for unrelated amendment history. |
| src:usa-ne:76-3504-radon | Neb. Rev. Stat. § 76-3504, Radon resistant new construction; minimum standards | statute | https://nebraskalegislature.gov/laws/statutes.php?statute=76-3504 | accessed 2026-06-26 | Radon-resistant new-construction minimum standards and date | Official legislature HTML. |
| src:usa-ne:81-503-01-state-fire-code | Neb. Rev. Stat. § 81-503.01, State Fire Code; State Fire Marshal; duties; contents; enforcement; plans | statute | https://nebraskalegislature.gov/laws/statutes.php?statute=81-503.01 | accessed 2026-06-26 | State Fire Marshal adopts/enforces State Fire Code; NFPA 1/NFPA 101 baseline; plan review | Official legislature HTML; statutory NFPA editions may differ from later rule details unless rules are parsed. |
| src:usa-ne:sfm-regulations | Nebraska State Fire Marshal, Regulations page | agency page | https://sfm.nebraska.gov/regulations | accessed 2026-06-26 | Title 153, Title 156, Title 229, Title 230 listings; plan-fee update notice | Official agency page; linked PDFs may be scanned or updated independently. |
| src:usa-ne:title153-ch1-fire-code | Title 153 - State Fire Marshal, Chapter 1 - Nebraska State Fire Code | regulation PDF | https://sfm.nebraska.gov/sites/default/files/doc/Title%20153%20Chapter%201%20Stamped.pdf | accessed 2026-06-26 | State Fire Code NFPA adopted-code listing; NFPA 1-2012 with Nebraska amendments; NFPA 54/58 editions | Stamped/scanned PDF; extraction is partial and should be verified against Secretary of State rules. |
| src:usa-ne:title153-ch21-plan-review | Title 153 - State Fire Marshal, Chapter 21 - Review and Approval of Plans and Fees Assessed | regulation PDF | https://sfm.nebraska.gov/sites/default/files/doc/title153-ch21%20Final%202025.pdf | accessed 2026-06-26 | Plan-review trigger; exemptions; 180-day approval expiration; late fee | Official agency PDF; parsed text available. |
| src:usa-ne:title153-ch22-co | Title 153 - State Fire Marshal, Chapter 22 - Certificate of Occupancy | regulation PDF | https://sfm.nebraska.gov/sites/default/files/doc/title153-ch22.pdf | accessed 2026-06-26 | SFM certificate-of-occupancy rule | Official agency PDF; parsed text available. |
| src:usa-ne:sfm-delegated-local-authority | Nebraska State Fire Marshal, Delegated Authority and Local Authority page | agency page | https://sfm.nebraska.gov/delegated-authority-and-local-authority | accessed 2026-06-26 | SFM delegated authority; local accessibility review authorities; state-owned/CMS exceptions; religious NAG exemption | Page warns local authority can change without notice; narrative count and displayed list should be reconciled. |
| src:usa-ne:81-5-147-accessibility | Neb. Rev. Stat. § 81-5,147, Buildings and facilities; standards, specifications, and exclusions; adoption | statute | https://nebraskalegislature.gov/laws/statutes.php?statute=81-5%2C147 | accessed 2026-06-26 | State Fire Marshal ADA/FHA accessibility-standard adoption authority | Official legislature HTML. |
| src:usa-ne:81-2104-electrical-board | Neb. Rev. Stat. § 81-2104, State Electrical Board; powers enumerated | statute | https://nebraskalegislature.gov/laws/statutes.php?statute=81-2104 | accessed 2026-06-26 | Electrical Board powers; 2023 NEC basis; 2017 NEC exception sections; enforcement authority | Official legislature HTML. |
| src:usa-ne:state-electrical-2023-adoption | Nebraska State Electrical, State Act Update & 2023 Code Adoption | agency page | https://electrical.nebraska.gov/state-act-update-2023-code-adoption | accessed 2026-06-26 | 2023 NEC adoption date/effective date; exam/permit transition | Official agency page; page uses “National Electric Code” wording, statute controls legal title. |
| src:usa-ne:state-electrical-statutes-rules | Nebraska State Electrical, Statutes & Rules | agency page | https://electrical.nebraska.gov/statutes-rules | accessed 2026-06-26 | State Electrical Act and State Electrical Board rules publication path | Official agency page; linked PDFs not fully parsed. |
| src:usa-ne:sfm-mechanical-safety | Nebraska State Fire Marshal, Mechanical Safety | agency page | https://sfm.nebraska.gov/mechanical-safety | accessed 2026-06-26 | Boiler and elevator/conveyance program pages and forms | Official agency page; code editions require linked-rule extraction. |
| src:usa-ne:sfm-title230-elevator | Title 230 Chapter 1, Safety Code for Elevators, Escalators and Other Conveyances | regulation PDF | https://sfm.nebraska.gov/sites/default/files/doc/Chapter-1.pdf | accessed 2026-06-26 | Elevator/conveyance inspection and definitions | Official agency PDF; full rule not parsed. |
| src:usa-ne:sfm-title229-boiler | Nebraska Administrative Code Title 229 - Boiler Safety Code | regulation PDF | https://sfm.nebraska.gov/sites/default/files/doc/Title%20229%20With%20TOC.pdf | accessed 2026-06-26 | Boiler inspection program path and transfer note | Official agency PDF; full rule not parsed. |
| src:usa-ne:81-5-230-elevator-licensing | Neb. Rev. Stat. § 81-5,230, Elevator mechanic license; elevator contractor license; application | statute | https://nebraskalegislature.gov/laws/statutes.php?statute=81-5%2C230 | accessed 2026-06-26 | Elevator mechanic/contractor licensing through State Fire Marshal | Official legislature HTML. |
| src:usa-ne:lb611-2025-status | LB611 bill page, 109th Legislature | bill history | https://nebraskalegislature.gov/bills/view_bill.php?DocumentID=59719 | accessed 2026-06-26 | 2025 building/energy update proposal; indefinitely postponed 2026-04-17 | Official bill page. |
| src:usa-ne:lb801-2026-status | LB801 bill page, 109th Legislature | bill history | https://nebraskalegislature.gov/bills/view_bill.php?DocumentID=62790 | accessed 2026-06-26 | 2026 building/energy update proposal; indefinitely postponed 2026-04-17 | Official bill page. |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-ne:title153-ch1-fire-code | scanned_stamped_pdf | Title 153 Chapter 1 is a stamped/scanned PDF; NFPA adoption list was read from screenshot/PDF text rather than a structured regulation dataset. | verify_against_secretary_of_state_rules |
| src:usa-ne:sfm-delegated-local-authority | agency_list_volatility | Page warns local authority can change without notice and contains a narrative/list-count inconsistency. | monitor_before_ahj_resolution |
| src:usa-ne:lb131-2021-final | derived_date | Operative date is derived from bill text and legislative calendar logic rather than an explicit date printed in current statute. | legal_calendar_confirmation_required |
| src:usa-ne:state-electrical-2023-adoption | agency_summary | Agency page supports transition date; statute controls the legal NEC adoption and exceptions. | use_with_statute |
| src:usa-ne:sfm-title229-boiler | incomplete_extraction | Boiler rule source located but not fully parsed into model-code editions and inspection triggers. | complete_rule_extraction |
| src:usa-ne:sfm-title230-elevator | incomplete_extraction | Elevator/conveyance rule source located but not fully parsed into model-code editions and inspection triggers. | complete_rule_extraction |

### 8.3 Supplemental Sources

None used as controlling authority. Non-official code-adoption summaries should be used only as navigation aids and not as the source of record for production fields.

### 8.4 Source Extraction Metadata

| Extraction ID | Source IDs | Extracted Fields | Method | Extracted On | Extractor Notes |
| --- | --- | --- | --- | --- | --- |
| extract:usa-ne:building-statutes | src:usa-ne:71-6403-state-building-code, src:usa-ne:71-6406-local-building-code | Core building-code adoptions; local conformance; fallback; amendment limits | official HTML review | 2026-06-26 | High-confidence authority and edition fields; operative date still derived. |
| extract:usa-ne:energy | src:usa-ne:81-1611-energy-code, src:usa-ne:81-1618-local-energy-code, src:usa-ne:dwee-energy-codes | 2018 IECC adoption; effective date; local energy code and waiver rules | official HTML review | 2026-06-26 | DWEE agency page supplies 2020-07-01 effective date and fallback enforcement. |
| extract:usa-ne:fire | src:usa-ne:81-503-01-state-fire-code, src:usa-ne:sfm-regulations, src:usa-ne:title153-ch1-fire-code, src:usa-ne:title153-ch21-plan-review, src:usa-ne:title153-ch22-co | SFM authority; NFPA baseline; plan review; CO | official HTML/PDF review plus screenshot for scanned PDF | 2026-06-26 | Fire-code edition list is partial due scanned Title 153 Chapter 1 source. |
| extract:usa-ne:electrical | src:usa-ne:81-2104-electrical-board, src:usa-ne:state-electrical-2023-adoption, src:usa-ne:state-electrical-statutes-rules | State Electrical Board authority; 2023 NEC and transition | official HTML review | 2026-06-26 | Good support for statewide electrical field. |
| extract:usa-ne:accessibility-mechanical | src:usa-ne:81-5-147-accessibility, src:usa-ne:sfm-delegated-local-authority, src:usa-ne:sfm-mechanical-safety, src:usa-ne:sfm-title229-boiler, src:usa-ne:sfm-title230-elevator | Accessibility authority; local review authority; boiler/elevator source locations | official HTML/PDF review | 2026-06-26 | Program authority located; detailed boiler/elevator editions unresolved. |
| extract:usa-ne:pending-bills | src:usa-ne:lb611-2025-status, src:usa-ne:lb801-2026-status | Pending/failed code-update bills | official bill-page review | 2026-06-26 | Both bills show indefinite postponement on 2026-04-17. |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 0.80 | none | Core authority/adoption fields have source backing; unresolved items remain explicit. |
| report | risk.overall_confidence | 0.72 | verified | 0.75 | none | Confidence reflects strong statutory support but incomplete amendments/AHJ parsing. |
| ahj:usa-ne:legislature-building-construction-act | legal_basis | Neb. Rev. Stat. §§ 71-6403 and 71-6406 | verified | 0.92 | src:usa-ne:71-6403-state-building-code, src:usa-ne:71-6406-local-building-code | Primary authority model verified. |
| adoption:usa-ne:building:2018-ibc | edition | 2018 IBC, with statutory exclusions | verified | 0.92 | src:usa-ne:71-6403-state-building-code | Statutory model-code edition verified. |
| adoption:usa-ne:residential:2018-irc | edition | 2018 IRC, except R313 and chapters 25-33 | verified | 0.92 | src:usa-ne:71-6403-state-building-code | Statutory model-code edition verified. |
| adoption:usa-ne:existing-building:2018-iebc | edition | 2018 IEBC, except section 809 | verified | 0.92 | src:usa-ne:71-6403-state-building-code | Statutory model-code edition verified. |
| adoption:usa-ne:plumbing:2018-upc | edition | 2018 UPC | verified | 0.92 | src:usa-ne:71-6403-state-building-code | Statutory model-code edition verified. |
| adoption:usa-ne:building:2018-ibc | operative_date | 2021-08-28 | partially_verified | 0.62 | src:usa-ne:lb131-2021-final, src:usa-ne:lb131-2021-status | Date is derived; requires legal-calendar confirmation. |
| adoption:usa-ne:energy:2018-iecc | effective_date | 2020-07-01 | verified | 0.88 | src:usa-ne:dwee-energy-codes | Agency page directly states effective date. |
| adoption:usa-ne:electrical:2023-nec | effective_date | 2024-08-01 | verified | 0.88 | src:usa-ne:state-electrical-2023-adoption, src:usa-ne:81-2104-electrical-board | Permit-date transition verified from agency page; legal code exceptions verified from statute. |
| local-enforcement:usa-ne | model | hybrid_state_fallback_local_primary | partially_verified | 0.78 | src:usa-ne:71-6406-local-building-code, src:usa-ne:81-503-01-state-fire-code, src:usa-ne:dwee-energy-codes, src:usa-ne:81-2104-electrical-board | Building, fire, energy, and electrical enforcement layers verified at summary level. |
| local-amendment-rule:usa-ne | model | local_amendments_allowed_with_statutory_conformance_boundaries | partially_verified | 0.79 | src:usa-ne:71-6406-local-building-code, src:usa-ne:81-1618-local-energy-code | Statewide local amendment rule verified at statutory level; local registries unresolved. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Each `src:usa-ne:*` used in the body is listed in Section 8. |
| All authority IDs resolve | pass | Authority IDs used in Section 2 are internally defined or described. |
| All current code families have adoption rows | pass | Matrix rows are present for building, residential, existing building, mechanical, plumbing, fuel gas, electrical, energy, fire, accessibility, elevator, and boiler. |
| Building and operational fire code are separated | pass | Building-code adoption under § 71-6403 is distinct from SFM State Fire Code authority under § 81-503.01 and Title 153. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Dates are separated; derived dates are labeled. |
| Effective dates are valid ISO dates | pass | Entered dates use ISO format; unresolved dates are `null`. |
| No impossible date sequences | pass | No effective/operative/mandatory sequence conflict was introduced. |
| Transition rules have explicit trigger conditions | pass | Local two-year fallback, energy fallback, electrical permit-date transition, and SFM plan-review triggers are captured. |
| Permit-date logic is captured where applicable | pass | Electrical permit-date logic is captured for 2024-08-01. |
| Local enforcement model classified | pass | Classified as hybrid_state_fallback_local_primary. |
| Local amendment rule classified | pass | Classified as local_amendments_allowed_with_statutory_conformance_boundaries. |
| AHJ confirmation metadata present | partial | SFM delegated/local authority sources are identified, but structured AHJ contacts and boundary data remain open. |
| Official-source caveats captured | pass | Caveats are recorded for scanned PDFs, agency list volatility, and derived dates. |
| Open issues remain explicit | pass | Unresolved boiler/elevator details, local registries, boundary data, and derived date validation remain listed in Section 10. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-ne:001 | high | LB131 operative date | 2021-08-28 operative date and 2023-08-28 local fallback date are derived from LB131 timing and § 71-6406; verify exact legal calendar treatment. | Confirm with official session-adjournment record or Nebraska Revisor/Secretary of State guidance. | null | null | open |
| issue:usa-ne:002 | high | State Fire Code full NFPA list | Title 153 Chapter 1 is scanned/stamped; the NFPA adopted-code list needs structured extraction, including NFPA 101 location and all Nebraska amendments. | Extract Title 153 Chapter 1 from Secretary of State or clean rule text; reconcile with statute. | null | null | open |
| issue:usa-ne:003 | medium | Local amendment registry | No statewide local building-amendment registry was verified. | Contact DWEE/SFM or search agency portals for local code filings and energy chapter deletion notices. | null | null | open |
| issue:usa-ne:004 | medium | Mechanical / boiler / elevator code editions | SFM Mechanical Safety sources were found, but boiler and elevator/conveyance code editions and inspection triggers were not fully parsed. | Parse Title 229 and Title 230 in detail; extract code editions, applicability, inspection cadence, and exceptions. | null | null | open |
| issue:usa-ne:005 | medium | AHJ boundary and contact resolution | Boundary data and local AHJ contacts are not structured. | Select boundary data sources, parse SFM delegated/local authority contacts, and collect local building-official sources for key jurisdictions. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-ne:building-statute-71-6403 | src:usa-ne:71-6403-state-building-code | html_diff | monthly | statute text changes or source history changes | 2026-06-26 | active |
| watch:usa-ne:local-building-statute-71-6406 | src:usa-ne:71-6406-local-building-code | html_diff | monthly | local adoption/amendment rule changes | 2026-06-26 | active |
| watch:usa-ne:energy-statutes | src:usa-ne:81-1611-energy-code | html_diff | monthly | Nebraska Energy Code edition or alternative-standard authority changes | 2026-06-26 | active |
| watch:usa-ne:dwee-energy-page | src:usa-ne:dwee-energy-codes | html_diff | monthly | new effective date, enforcement guidance, or 2021/2024 IECC update | 2026-06-26 | active |
| watch:usa-ne:fire-regulations | src:usa-ne:sfm-regulations | html_diff | monthly | Title 153, Title 156, Title 229, or Title 230 update | 2026-06-26 | active |
| watch:usa-ne:sfm-delegated-authority | src:usa-ne:sfm-delegated-local-authority | html_diff | monthly | delegated authority/local authority list change | 2026-06-26 | active |
| watch:usa-ne:electrical-81-2104 | src:usa-ne:81-2104-electrical-board | html_diff | monthly | NEC edition or statutory exception changes | 2026-06-26 | active |
| watch:usa-ne:electrical-adoption-page | src:usa-ne:state-electrical-2023-adoption | html_diff | monthly | permit transition or NEC adoption guidance changes | 2026-06-26 | active |
| watch:usa-ne:legislature-current-building-bills | src:usa-ne:lb801-2026-status | bill_search | monthly_in_session | new building, energy, electrical, fire, accessibility, boiler, or elevator code bills | 2026-06-26 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-24 | Generated baseline draft report | report:usa-ne | none | prior draft | Baseline contained unresolved placeholder content. |
| 2026-06-26 | Populated Nebraska report with official-source-backed authority, adoption, local enforcement, local amendment, date-rule, and QA sections | report:usa-ne; ahj:usa-ne:legislature-building-construction-act; adoption:usa-ne:building:2018-ibc; adoption:usa-ne:energy:2018-iecc; adoption:usa-ne:electrical:2023-nec; local-enforcement:usa-ne; local-amendment-rule:usa-ne | src:usa-ne:71-6403-state-building-code, src:usa-ne:71-6406-local-building-code, src:usa-ne:81-1611-energy-code, src:usa-ne:dwee-energy-codes, src:usa-ne:81-2104-electrical-board, src:usa-ne:81-503-01-state-fire-code | ChatGPT | Status advanced to partially_verified with unresolved items explicit. |
