---
state:
  state_id: "US-MO"
  name: "Missouri"
  abbreviation: "MO"
report:
  report_id: "state-report:usa-mo"
  schema_version: "0.2.1"
  status: "partially_verified" # draft | partially_verified | verified | deprecated
  last_updated: "2026-06-26"
  last_verified: "2026-06-26"
  reviewed_by: null
risk:
  overall_confidence: 0.62 # 0.00 - 1.00
  risk_flags:
    - "general_statewide_model_building_code_not_identified"
    - "local_code_adoption_varies_by_jurisdiction"
    - "local_amendment_registry_not_identified"
    - "state_statute_pages_uncertified_unofficial_revisor_copy"
    - "csr_effective_dates_require_rule_level_confirmation"
    - "facility_specific_fire_safety_rules_not_fully_parsed"
  open_questions_count: 6

---

# State Building Code Authority Report: Missouri

## 1. Executive Summary

- **Authority model:** Missouri is best modeled as a local-adoption state for general building, residential, mechanical, plumbing, fuel gas, electrical, energy, and local fire-code administration. Missouri statutes authorize communities to adopt technical codes by reference when they otherwise have legal authority; "community" includes counties, fire protection districts, and municipalities. County and municipal authority is not identical and must be resolved by jurisdiction.

- **Statewide code status:** No general statewide IBC/IRC/IMC/IPC/IFGC/NEC/IECC adoption for all private construction was identified in this pass. State-level rules do exist for narrower scopes: elevator and related equipment, state-building energy efficiency, professional conduct evaluation criteria where no local building code exists, seismic design ordinances for certain New Madrid earthquake exposure areas, and fire-safety rules for certain licensed facilities.

- **Local enforcement model:** Local jurisdictions are the primary source for adopted building-code details. Missouri DNR maintains county and city adopted-code listings, but the page states that the lists are general, may omit amendments and trade codes, and should be confirmed with the local jurisdiction.

- **Local amendment posture:** Local code adoption by reference can include amendments, but the code or amendment must be identified by date and source and filed with the community clerk for public inspection for 90 days before adoption. No statewide local-amendment approval workflow or complete statewide amendment registry was identified.

- **Known transition periods or pending changes:** Confirmed transition rules include the 90-day local code-copy filing rule, the county voter-approval condition for certain counties adopting building codes after August 28, 2001, the state-building energy date rules, elevator installation date bands, and seismic applicability rules for covered New Madrid earthquake-risk jurisdictions.

- **Production readiness:** partially_ready_for_state_level_routing; not_ready_for_address_level_code_determination without local ordinance ingestion.

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: No single primary statewide general building-code adopting authority was
    identified; general code adoption is local and statute-enabled.
  confidence: 0.72
  source_ids:
  - src:usa-mo:dnr-codes-by-jurisdiction
  - src:usa-mo:rsmo-67-280
  - src:usa-mo:rsmo-64-170
  - src:usa-mo:rsmo-89-020
- topic: Primary building code edition
  finding: No statewide IBC edition applies to all private construction. The Missouri
    professional board uses the 2018 IBC only as professional-conduct evaluation criteria
    when no local building code exists.
  confidence: 0.75
  source_ids:
  - src:usa-mo:csr-20-2030-2
- topic: Electrical code authority
  finding: No statewide NEC adoption was identified. Counties of the first and second
    classification may regulate electrical wiring/installations under RSMo 64.170;
    communities may adopt electrical technical codes by reference under RSMo 67.280
    if otherwise authorized.
  confidence: 0.6
  source_ids:
  - src:usa-mo:rsmo-64-170
  - src:usa-mo:rsmo-67-280
- topic: Fire code authority
  finding: Missouri Division of Fire Safety states Missouri has no statewide fire
    code, while state fire-safety regulations affect most licensed day cares, nursing
    homes, and other state-licensed facilities.
  confidence: 0.82
  source_ids:
  - src:usa-mo:dfs-fire-faq
- topic: Elevator / conveyance
  finding: Missouri Division of Fire Safety, Elevator Safety Unit administers elevator
    standards; elevators installed after 2020-04-30 must comply with ASME A17.1 2016.
  confidence: 0.86
  source_ids:
  - src:usa-mo:dps-dfs-elevator-code-adoptions
  - src:usa-mo:dps-dfs-elevator-code-info
- topic: State-building energy
  finding: State buildings over 5,000 square feet, substantial renovations, and certain
    state leases/acquisitions are subject to minimum energy efficiency standards.
  confidence: 0.78
  source_ids:
  - src:usa-mo:rsmo-8-812
  - src:usa-mo:csr-10-140-7-010
- topic: Local amendments
  finding: Local code adoption by reference can include amendments with 90-day clerk
    filing, but no complete state-level amendment registry was identified.
  confidence: 0.58
  source_ids:
  - src:usa-mo:rsmo-67-280
  - src:usa-mo:dnr-codes-by-jurisdiction
- topic: Effective / operative date rules
  finding: Statewide local-code effective dates are jurisdiction-specific. Confirmed
    state-level rules include state-building energy dates, elevator installation date
    bands, and seismic applicability dates.
  confidence: 0.64
  source_ids:
  - src:usa-mo:rsmo-67-280
  - src:usa-mo:rsmo-8-812
  - src:usa-mo:csr-10-140-7-010
  - src:usa-mo:dps-dfs-elevator-code-adoptions
  - src:usa-mo:rsmo-319-200
  - src:usa-mo:rsmo-319-203
```

---

### 2.1 Primary Building Code Authorities

| Field | Value |
| --- | --- |
| Authority ID | ahj:usa-mo:local-communities |
| Authority name | Missouri counties, municipalities, and fire protection districts where otherwise authorized by law |
| Authority type | local_adoption_authorities |
| Legal basis | RSMo 67.280; RSMo 64.170; RSMo 89.020; RSMo 321.228 |
| Role | Adopt and administer local technical codes, subject to each entity's underlying statutory authority and scope. |
| Enforcement model | local_adoption_and_local_enforcement |
| Source IDs | src:usa-mo:rsmo-67-280; src:usa-mo:rsmo-64-170; src:usa-mo:rsmo-89-020; src:usa-mo:rsmo-321-228; src:usa-mo:dnr-codes-by-jurisdiction |
| Verification status | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family | Authority ID | Authority Name | Role | Legal Basis | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| Building | ahj:usa-mo:local-communities | Local counties, municipalities, and other authorized communities | Adopt/administer local building codes where authorized; no general statewide model-code adoption identified. | RSMo 67.280; RSMo 64.170; RSMo 89.020 | src:usa-mo:rsmo-67-280; src:usa-mo:rsmo-64-170; src:usa-mo:rsmo-89-020; src:usa-mo:dnr-codes-by-jurisdiction | partially_verified |
| Residential | ahj:usa-mo:local-communities | Local counties and municipalities; fire protection district authority limited where local residential systems exist | Local residential construction regulatory systems may supersede fire protection district residential construction systems; fire districts retain specified hydrant, flow, and fire-lane authority. | RSMo 321.228; RSMo 67.280 | src:usa-mo:rsmo-321-228; src:usa-mo:rsmo-67-280 | partially_verified |
| Existing Building / Rehabilitation | ahj:usa-mo:local-communities | Local counties and municipalities | Existing building or rehabilitation code status is local unless a specific state/facility rule applies. | RSMo 67.280; local ordinances pending | src:usa-mo:rsmo-67-280; src:usa-mo:dnr-codes-by-jurisdiction | unresolved_local_detail |
| Mechanical | ahj:usa-mo:local-communities | Local counties, municipalities, fire protection districts where otherwise authorized | Mechanical technical codes may be adopted by reference by authorized communities. | RSMo 67.280 | src:usa-mo:rsmo-67-280 | partially_verified |
| Plumbing | ahj:usa-mo:local-communities | Local counties and municipalities | Plumbing technical codes may be adopted by reference; first/second class county authority includes plumbing and drain laying. | RSMo 67.280; RSMo 64.170 | src:usa-mo:rsmo-67-280; src:usa-mo:rsmo-64-170 | partially_verified |
| Fuel Gas | ahj:usa-mo:local-communities | Local counties and municipalities | Fuel gas code status was not separately verified; likely local when adopted as a technical code or through model-code package. | RSMo 67.280; local ordinances pending | src:usa-mo:rsmo-67-280 | unresolved_local_detail |
| Electrical | ahj:usa-mo:local-communities | Local counties and municipalities | Electrical technical codes may be adopted by reference; first/second class county authority includes electrical wiring/installations and licensing/inspection. | RSMo 67.280; RSMo 64.170 | src:usa-mo:rsmo-67-280; src:usa-mo:rsmo-64-170 | partially_verified |
| Energy | ahj:usa-mo:local-communities; ahj:usa-mo:dnr-division-energy | Local jurisdictions for private construction; Missouri DNR Division of Energy for state-building minimum energy efficiency standards | Local energy-code adoption varies; state-building energy standard applies to state buildings, substantial renovations, and certain state leases/acquisitions. | RSMo 8.812; 10 CSR 140-7.010 | src:usa-mo:dnr-codes-by-jurisdiction; src:usa-mo:rsmo-8-812; src:usa-mo:csr-10-140-7-010 | partially_verified |
| Fire - construction references | ahj:usa-mo:local-communities; ahj:usa-mo:fire-protection-districts | Local communities and fire protection districts within statutory limits | No statewide fire code; local fire code and facility-specific state fire-safety rules may apply. | RSMo 67.280; RSMo 321.228 | src:usa-mo:dfs-fire-faq; src:usa-mo:rsmo-67-280; src:usa-mo:rsmo-321-228 | partially_verified |
| Fire - operational / prevention code | ahj:usa-mo:dfs-fire-safety; ahj:usa-mo:local-communities | Missouri Division of Fire Safety for licensed-facility fire-safety programs; local jurisdictions for adopted local fire codes | State does not have a statewide fire code; DFS notes state regulations affect most licensed day cares, nursing homes, and other state-licensed facilities. | Facility-specific rules pending; local codes pending | src:usa-mo:dfs-fire-faq | partially_verified_scope_limited |
| Accessibility | ahj:usa-mo:local-communities; ahj:usa-mo:dfs-elevator-safety | Local jurisdictions; DFS Elevator Safety Unit for elevator-related accessibility requirements | General statewide accessibility building-code adoption not resolved; elevator page requires new passenger elevators and platform lifts to comply with ANSI A117.1 2009. | Elevator Safety Act program; local ordinances pending | src:usa-mo:dps-dfs-elevator-code-adoptions | partially_verified_scope_limited |
| Elevator / Conveyance | ahj:usa-mo:dfs-elevator-safety | Missouri Division of Fire Safety, Elevator Safety Unit | Administers elevator and related-equipment code standards, permits, inspections, and operating certificate requirements. | Elevator Safety Act; 11 CSR 40-5.065 referenced by DFS | src:usa-mo:dps-dfs-elevator-code-adoptions; src:usa-mo:dps-dfs-elevator-code-info | partially_verified |
| Professional conduct fallback | ahj:usa-mo:apeplsla-board | Missouri Board for Architects, Professional Engineers, Professional Land Surveyors, and Professional Landscape Architects | Uses 2018 IBC as evaluation criteria for professional conduct when no local building code exists; this is not a general statewide building-code adoption. | 20 CSR 2030-2.040 | src:usa-mo:csr-20-2030-2 | verified_limited_scope |
| Seismic special applicability | ahj:usa-mo:local-communities-seismic | Cities, towns, villages, and counties identified as New Madrid seismic-risk jurisdictions | Must adopt seismic design/construction ordinances or orders for covered new construction, additions, and alterations. | RSMo 319.200 to 319.207 | src:usa-mo:rsmo-319-200; src:usa-mo:rsmo-319-203 | partially_verified |

### 2.3 Authority Hierarchy Notes

Missouri should not be represented as a simple statewide-code state. The practical hierarchy for a building-code lookup is:

1. Determine whether the project is within a municipality, unincorporated county area, fire protection district, or other special jurisdiction.
2. Check whether that jurisdiction has adopted construction, mechanical, plumbing, electrical, fire-prevention, energy, or other technical codes by ordinance/order.
3. Apply state-reserved or state-administered programs only for their limited scopes: elevator/conveyance, state-building energy efficiency, facility-specific fire-safety regulation, professional conduct evaluation, and seismic ordinances in covered earthquake-risk jurisdictions.
4. Confirm local amendments and trade-code editions directly with the local jurisdiction because the DNR adopted-code lists are general and omit specific amendments and trade codes.

### 2.4 Authority Graph Edges

| Edge ID | From Authority | Relationship | To Authority / Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| edge:usa-mo:001 | ahj:usa-mo:local-communities | may_adopt_by_reference | construction, occupancy, mechanical, plumbing, electrical, and fire-prevention technical codes when otherwise authorized | src:usa-mo:rsmo-67-280 | partially_verified |
| edge:usa-mo:002 | ahj:usa-mo:county-commission-first-second-class | may_adopt_and_enforce | building, electrical, plumbing, drain-laying, permit, inspection, and fee regulations for covered counties | src:usa-mo:rsmo-64-170 | partially_verified |
| edge:usa-mo:003 | ahj:usa-mo:municipal-legislative-body | may_regulate | building height, stories, size, yards, density, historical features, and land use through zoning authority | src:usa-mo:rsmo-89-020 | partially_verified |
| edge:usa-mo:004 | ahj:usa-mo:city-county-residential-system | supersedes_enforcement_by | fire protection district residential construction regulatory system, where the city/town/village/county adopts, implements, and enforces such system | src:usa-mo:rsmo-321-228 | partially_verified |
| edge:usa-mo:005 | ahj:usa-mo:fire-protection-districts | retains_authority_for | residential fire hydrant locations/specifications, hydrant flow rates, and fire lanes | src:usa-mo:rsmo-321-228 | partially_verified |
| edge:usa-mo:006 | ahj:usa-mo:dfs-fire-safety | administers_limited_state_rules_for | licensed day cares, nursing homes, and other state-licensed facilities; exact rule sets pending | src:usa-mo:dfs-fire-faq | partially_verified_scope_limited |
| edge:usa-mo:007 | ahj:usa-mo:dfs-elevator-safety | administers | elevators and related equipment operated in Missouri | src:usa-mo:dps-dfs-elevator-code-adoptions; src:usa-mo:dps-dfs-elevator-code-info | partially_verified |
| edge:usa-mo:008 | ahj:usa-mo:dnr-division-energy | administers_limited_state_standard_for | state buildings over 5,000 square feet, substantial renovations, and certain state acquisitions/leases | src:usa-mo:rsmo-8-812; src:usa-mo:csr-10-140-7-010 | partially_verified |
| edge:usa-mo:009 | ahj:usa-mo:apeplsla-board | uses_evaluation_criteria | 2018 IBC for professional-conduct evaluation when no local building code exists | src:usa-mo:csr-20-2030-2 | verified_limited_scope |
| edge:usa-mo:010 | ahj:usa-mo:local-communities-seismic | must_adopt_order_or_ordinance_for | seismic design/construction standards in identified New Madrid earthquake-risk areas | src:usa-mo:rsmo-319-200; src:usa-mo:rsmo-319-203 | partially_verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family | State Code Name | Base Model Code | Edition | Status | Adoption Date | Effective Date | Operative Date | Mandatory Date | Transition Rule | Source IDs |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Building | No general statewide building code identified | none statewide for all private construction | null | local_adoption | null | null | null | null | Local ordinance/order and permit rules control; 2018 IBC is used only as professional-conduct evaluation criteria when no local code exists. | src:usa-mo:dnr-codes-by-jurisdiction; src:usa-mo:rsmo-67-280; src:usa-mo:csr-20-2030-2 |
| Residential | No general statewide residential code identified | none statewide for all private construction | null | local_adoption | null | null | null | null | Local residential construction regulatory systems control; fire protection district authority is limited where city/county residential systems are adopted and enforced. | src:usa-mo:rsmo-321-228; src:usa-mo:rsmo-67-280 |
| Existing Building / Rehabilitation | No general statewide existing building code identified | none statewide for all private construction | null | local_adoption_unresolved_detail | null | null | null | null | Confirm local ordinance and whether IEBC or local rehab provisions were adopted. | src:usa-mo:dnr-codes-by-jurisdiction; src:usa-mo:rsmo-67-280 |
| Mechanical | No general statewide mechanical code identified | none statewide for all private construction | null | local_adoption | null | null | null | null | Authorized communities may adopt mechanical technical codes by reference. | src:usa-mo:rsmo-67-280 |
| Plumbing | No general statewide plumbing code identified | none statewide for all private construction | null | local_adoption | null | null | null | null | Authorized communities may adopt plumbing technical codes by reference; first/second class counties may regulate plumbing and drain laying. | src:usa-mo:rsmo-67-280; src:usa-mo:rsmo-64-170 |
| Fuel Gas | No general statewide fuel gas code identified | none statewide for all private construction | null | local_adoption_unresolved_detail | null | null | null | null | Confirm whether local fuel-gas provisions are adopted through IFGC, IRC, NFPA, utility rules, or local ordinance. | src:usa-mo:rsmo-67-280 |
| Electrical | No general statewide electrical code identified | none statewide for all private construction | null | local_adoption | null | null | null | null | Authorized communities may adopt electrical technical codes; first/second class county authority includes electrical wiring/installations, licensing, and inspections. | src:usa-mo:rsmo-67-280; src:usa-mo:rsmo-64-170 |
| Energy | State Building Minimum Energy Efficiency Standards for state buildings; local energy codes for private construction | State buildings: ASHRAE 90.1 latest except low-rise residential; ASHRAE 90.2 latest or CABO MEC latest for low-rise residential; statute minimum at least IECC 2006/latest | latest-as-published under rule | limited_state_scope_and_local_adoption | null | 2009-07-01 for statutory design trigger; rule effective history includes 2020-01-15 move to 10 CSR | 2009-07-01 for covered state design under statute | 2009-07-01 for covered state design under statute | New edition of minimum energy standard becomes effective three months after publication for projects not started; local private-construction rules vary. | src:usa-mo:rsmo-8-812; src:usa-mo:csr-10-140-7-010; src:usa-mo:dnr-codes-by-jurisdiction |
| Fire - construction references | No statewide fire code identified | none statewide | null | local_adoption_and_facility_specific_state_rules | null | null | null | null | Local fire code or state licensed-facility fire-safety rules may apply. | src:usa-mo:dfs-fire-faq; src:usa-mo:rsmo-67-280 |
| Fire - operational / prevention code | No statewide fire code identified | none statewide | null | local_adoption_and_facility_specific_state_rules | null | null | null | null | DFS states no statewide fire code; contact local jurisdiction for adopted codes; licensed facilities may have state fire-safety rules. | src:usa-mo:dfs-fire-faq |
| Accessibility | No general statewide accessibility building code resolved | elevator-related: ANSI A117.1 | 2009 for new passenger elevators/platform lifts under DFS page | partially_verified_scope_limited | null | null | null | null | General accessibility code posture requires separate review; elevator-related accessibility requirement verified. | src:usa-mo:dps-dfs-elevator-code-adoptions |
| Elevator / Conveyance | Missouri elevator and related-equipment standards | ASME A17.1 for elevators; ASME A18.1 for platform/stair lifts; ANSI A10.4 for personnel hoists; ANSI A117.1 for elevator-related accessibility | ASME A17.1 2016 for elevators installed after 2020-04-30 | state_administered | null | 2020-04-30 | 2020-04-30 | 2020-04-30 | Installation-date bands determine applicable edition; state permit and operating certificate requirements remain. | src:usa-mo:dps-dfs-elevator-code-adoptions; src:usa-mo:dps-dfs-elevator-code-info |
| Professional conduct fallback | Evaluation Criteria for Building Design | International Building Code | 2018 | professional_conduct_evaluation_only | null | 2020-07-30 | 2020-07-30 | null | Applies only in the absence of any local building code and only as evaluation criteria for professional conduct under the board rule. | src:usa-mo:csr-20-2030-2 |
| Seismic special applicability | Local seismic design/construction ordinances or orders in identified New Madrid earthquake-risk jurisdictions | 1990 or later Uniform Building Code or BOCA Code seismic design/construction standards | 1990 or later | required_local_ordinance_in_identified_areas | null | 1996-08-28 for current RSMo 319.200 text; 1991-08-28 for RSMo 319.203 text | 1994-01-01 construction start trigger for certain leased buildings | varies_by_scope | Applies to identified cities, towns, villages, and counties; private structures under 10,000 square feet and single-family/duplex residences are excluded unless local ordinance reaches further. | src:usa-mo:rsmo-319-200; src:usa-mo:rsmo-319-203 |

### 3.2 Adoption Records

#### adoption:usa-mo:local-technical-code-enabling

```yaml
adoption_id: "adoption:usa-mo:local-technical-code-enabling"
code_families:
  - Building
  - Mechanical
  - Plumbing
  - Electrical
  - Fire prevention
state_code_name: "Local technical-code adoption by reference"
base_model_code: "varies_by_local_jurisdiction"
edition: "varies_by_local_jurisdiction"
status: "local_adoption_enabled"
adoption_date: null
effective_date: "2009-08-28"
operative_date: null
mandatory_date: null
scope: "Communities may adopt codes by reference only when otherwise authorized by law."
source_ids:
  - "src:usa-mo:rsmo-67-280"
confidence: 0.72
```

#### adoption:usa-mo:county-building-code-authority-first-second-class

```yaml
adoption_id: "adoption:usa-mo:county-building-code-authority-first-second-class"
code_families:
  - Building
  - Electrical
  - Plumbing
state_code_name: "County building regulations under RSMo 64.170"
base_model_code: "varies_by_county_order_or_ordinance"
edition: "varies_by_county_order_or_ordinance"
status: "county_authority_enabled_with_conditions"
adoption_date: null
effective_date: "2008-08-28"
operative_date: null
mandatory_date: null
scope: "Counties of first and second classification may adopt construction, electrical, plumbing, permit, inspection, licensing, and fee regulations; counties that had not adopted before 2001-08-28 require voter approval."
source_ids:
  - "src:usa-mo:rsmo-64-170"
confidence: 0.68
```

#### adoption:usa-mo:professional-conduct-ibc-2018

```yaml
adoption_id: "adoption:usa-mo:professional-conduct-ibc-2018"
code_family: "Building"
state_code_name: "Evaluation Criteria for Building Design"
base_model_code: "International Building Code"
edition: "2018"
status: "professional_conduct_evaluation_only"
adoption_date: null
effective_date: "2020-07-30"
operative_date: "2020-07-30"
mandatory_date: null
scope: "Applies to evaluation of appropriate conduct for professionals licensed or regulated by the Missouri APEPLSPLA board when no local building code exists."
source_ids:
  - "src:usa-mo:csr-20-2030-2"
confidence: 0.82
```

#### adoption:usa-mo:state-building-energy-standard

```yaml
adoption_id: "adoption:usa-mo:state-building-energy-standard"
code_family: "Energy"
state_code_name: "State Building Minimum Energy Efficiency Standards"
base_model_code:
  non_low_rise_residential: "ASHRAE 90.1 latest edition"
  low_rise_residential: "ASHRAE 90.2 latest edition or CABO Model Energy Code latest edition"
  statutory_floor: "At least IECC 2006 or latest version under RSMo 8.812"
edition: "latest_as_published_under_rule"
status: "state_buildings_and_state_leases_only"
adoption_date: null
effective_date: "2009-07-01"
operative_date: "2009-07-01"
mandatory_date: "2009-07-01"
scope: "State building over 5,000 square feet, substantial renovation of state building over 5,000 square feet where major energy systems are involved, or building over 5,000 square feet considered for state acquisition or lease."
source_ids:
  - "src:usa-mo:rsmo-8-812"
  - "src:usa-mo:csr-10-140-7-010"
confidence: 0.78
```

#### adoption:usa-mo:elevator-asme-a17-1-2016

```yaml
adoption_id: "adoption:usa-mo:elevator-asme-a17-1-2016"
code_family: "Elevator / Conveyance"
state_code_name: "Elevator and related-equipment standards"
base_model_code: "ASME A17.1"
edition: "2016"
status: "state_administered"
adoption_date: null
effective_date: "2020-04-30"
operative_date: "2020-04-30"
mandatory_date: "2020-04-30"
scope: "Elevators installed after 2020-04-30. Other equipment and older installations use separate date bands and referenced standards."
source_ids:
  - "src:usa-mo:dps-dfs-elevator-code-adoptions"
  - "src:usa-mo:dps-dfs-elevator-code-info"
confidence: 0.86
```

#### adoption:usa-mo:seismic-new-madrid-local-ordinances

```yaml
adoption_id: "adoption:usa-mo:seismic-new-madrid-local-ordinances"
code_family: "Seismic special applicability"
state_code_name: "Local seismic design/construction ordinance requirement"
base_model_code: "Uniform Building Code or BOCA Code seismic design/construction standards"
edition: "1990 or later"
status: "required_local_ordinance_in_identified_areas"
adoption_date: null
effective_date: "1996-08-28"
operative_date: "1994-01-01 for certain covered leases/buildings under RSMo 319.203"
mandatory_date: null
scope: "Cities, towns, villages, or counties expected to experience Modified Mercalli VII or above from a New Madrid Fault magnitude 7.6 earthquake."
source_ids:
  - "src:usa-mo:rsmo-319-200"
  - "src:usa-mo:rsmo-319-203"
confidence: 0.70
```

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

Missouri does not have one statewide adoption date for a general private-building code stack. Date resolution must be done by code family and jurisdiction. Confirmed date logic includes local code-copy filing before adoption, conditional county voter approval, state-building energy applicability, elevator installation date bands, professional board evaluation criteria effective date, and New Madrid seismic applicability rules.

### 4.2 Date Rule Table

| Rule ID | Applies To | Rule Type | Date / Period | Trigger Condition | Prior Code Allowed? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| date-rule:usa-mo:001 | Local technical-code adoption by reference | pre_adoption_public_copy_filing | 90 days before adoption | Community adopts code, portion of code, or amendment by reference | not_applicable | src:usa-mo:rsmo-67-280 | verified |
| date-rule:usa-mo:002 | Counties of first and second classification using RSMo 64.170 authority | voter_approval_condition | 2001-08-28 cutoff | County had not adopted a building code before 2001-08-28 | not_applicable | src:usa-mo:rsmo-64-170 | verified |
| date-rule:usa-mo:003 | State-building minimum energy standard | statutory_design_trigger | 2009-07-01 | Design initiated on or after 2009-07-01 for covered state buildings, substantial renovations, state acquisitions, or leases | no for covered design unless exempted by commissioner | src:usa-mo:rsmo-8-812 | verified |
| date-rule:usa-mo:004 | State-building minimum energy standard | model_standard_update | three months after publication | New edition of minimum energy efficiency standard is published and project work has not started | yes for prior design/approval; projects underway should be reviewed and upgraded but prior approval is not displaced | src:usa-mo:csr-10-140-7-010 | verified |
| date-rule:usa-mo:005 | Elevator installations | installation_date_band | after 2020-04-30 | Elevator installed after 2020-04-30 | older installations use older date-band standards | src:usa-mo:dps-dfs-elevator-code-adoptions | verified |
| date-rule:usa-mo:006 | Professional conduct evaluation criteria | rule_effective_date | 2020-07-30 | Professional conduct evaluated under 20 CSR 2030-2.040 where no local building code exists | no subsequent IBC amendments incorporated by this rule | src:usa-mo:csr-20-2030-2 | verified_limited_scope |
| date-rule:usa-mo:007 | New Madrid seismic-risk jurisdictions | local_ordinance_obligation | current RSMo 319.200 text effective 1996-08-28 | State geologist and USGS identify jurisdictions expected to experience MM VII or above from specified New Madrid earthquake scenario | older seismic standard not resolved; statute requires 1990 or later UBC or BOCA seismic standards | src:usa-mo:rsmo-319-200 | partially_verified |
| date-rule:usa-mo:008 | Seismic leased-building applicability | construction_start_trigger | construction begun after 1994-01-01 | Covered leases by political subdivisions, state, or higher education institutions | statutory exemptions apply | src:usa-mo:rsmo-319-203 | verified |
| date-rule:usa-mo:009 | Local private-building code updates | local_effective_date | varies | Local ordinance/order adoption, permit application, permit issuance, or locally defined transition provision | varies_by_jurisdiction | src:usa-mo:dnr-codes-by-jurisdiction | unresolved_local_detail |

### 4.3 Pending or Future Code Changes

| Code Family | Future Code | Announced Date | Adoption Date | Effective Date | Operative Date | Mandatory Date | Watch Status | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| General building / energy | none confirmed | null | null | null | null | null | monitor_legislation_and_register | src:usa-mo:revisor-certification-status; src:usa-mo:sos-csr-current-caveat | No enacted statewide general model-code update was confirmed in this report. |
| Elevator / Conveyance | none confirmed | null | null | null | null | null | monitor_dfs_page_and_csr | src:usa-mo:dps-dfs-elevator-code-adoptions; src:usa-mo:dps-dfs-elevator-code-info | Watch for ASME A17.1 update or DFS/SOS rulemaking. |
| State-building energy | rolling model-standard updates | null | null | three months after publication of new referenced standard for projects not started | same as effective date | same as effective date | monitor_csr_and_dnr | src:usa-mo:csr-10-140-7-010 | Rule uses latest editions and a three-month publication trigger. |

### 4.4 Special Applicability Rules

| Rule ID | Code Family | Project Type / Scope | Trigger | Rule Summary | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- |
| applicability-rule:usa-mo:001 | Energy | State building over 5,000 square feet; substantial renovation over 5,000 square feet with major energy systems; state acquisition/lease over 5,000 square feet | Design initiated on/after 2009-07-01 | Must meet applicable provisions of minimum energy efficiency standard unless exempted by Commissioner of Administration for safety or cost-effectiveness reasons. | src:usa-mo:rsmo-8-812 | verified |
| applicability-rule:usa-mo:002 | Energy | State-owned buildings, additions, substantial renovations, smaller renovations/replacements, and leased space | Rule-specific contract/lease thresholds | 10 CSR 140-7.010 sets mandatory and "should conform" provisions by state-owned and leased-space categories. | src:usa-mo:csr-10-140-7-010 | partially_verified |
| applicability-rule:usa-mo:003 | Elevator / Conveyance | Elevators and related equipment operated in Missouri | Installation date and equipment type | DFS Elevator Safety Unit applies different ASME/ANSI editions by installation date and equipment type; state permit/certificate requirements apply. | src:usa-mo:dps-dfs-elevator-code-adoptions; src:usa-mo:dps-dfs-elevator-code-info | verified |
| applicability-rule:usa-mo:004 | Fire | Licensed day cares, nursing homes, and other state-licensed facilities | Facility licensure | DFS states state regulations affect most licensed day cares, nursing homes, and other state-licensed facilities, despite no statewide fire code. Exact facility-code rules remain to be parsed. | src:usa-mo:dfs-fire-faq | partially_verified_scope_limited |
| applicability-rule:usa-mo:005 | Residential fire / construction | Detached single-family and two-family residential construction | City/town/village/county adopts, implements, and enforces residential construction regulatory system | Fire protection district residential construction regulatory systems become advisory only in that jurisdiction; fire districts retain hydrant, flow, and fire-lane authority. | src:usa-mo:rsmo-321-228 | verified |
| applicability-rule:usa-mo:006 | Seismic | Covered jurisdictions and covered construction/leases in New Madrid earthquake-risk areas | Jurisdiction identified for expected MM VII or above shaking; construction/lease triggers | Identified local governments must adopt seismic design/construction ordinances or orders using 1990-or-later UBC or BOCA seismic standards; statutory exclusions apply. | src:usa-mo:rsmo-319-200; src:usa-mo:rsmo-319-203 | partially_verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** limited state-scope rules plus local amendments.

**Where amendments are published:**

- General local model-code amendments: local ordinances/orders and clerk-filed code/amendment copies under RSMo 67.280.
- Elevator/conveyance: Missouri Division of Fire Safety pages and applicable CSR rules.
- State-building energy: Missouri CSR and Missouri Register in-addition/update process.
- Professional conduct fallback: 20 CSR 2030-2.040; rule incorporates 2018 IBC and does not incorporate subsequent amendments or additions.

**Amendment parsing status:** partial.

### 5.2 State Amendment Sources

| Amendment Source ID | Code Family | Publication Path | Scope | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| amendment-source:usa-mo:local-clerk-filed-codes | local technical codes | local community clerk office and local ordinance records | Local codes and amendments adopted by reference | src:usa-mo:rsmo-67-280 | partially_verified |
| amendment-source:usa-mo:dnr-local-code-lists | local building/energy code listings | DNR Energy Codes by Jurisdiction page / Missouri Data Portal | General adopted code by county/city; not amendments or trade codes | src:usa-mo:dnr-codes-by-jurisdiction | verified_caveated |
| amendment-source:usa-mo:elevator-dfs | elevator/conveyance | DFS Elevator Safety Unit pages and referenced CSR | State elevator standards and installation date bands | src:usa-mo:dps-dfs-elevator-code-adoptions; src:usa-mo:dps-dfs-elevator-code-info | partially_verified |
| amendment-source:usa-mo:state-building-energy | energy | 10 CSR 140-7.010 and Missouri Register in-addition process | State-building minimum energy efficiency standard | src:usa-mo:csr-10-140-7-010 | partially_verified |
| amendment-source:usa-mo:professional-conduct-ibc | building professional conduct | 20 CSR 2030-2.040 | 2018 IBC as evaluation criteria; no subsequent amendments incorporated | src:usa-mo:csr-20-2030-2 | verified_limited_scope |

### 5.3 High-Impact State Amendments

| Amendment ID | Code Family | Amendment / Rule | Effect | Source IDs | Status |
| --- | --- | --- | --- | --- | --- |
| amendment:usa-mo:ibc-2018-no-subsequent-amendments | Building | 20 CSR 2030-2.040 incorporates 2018 IBC for professional-conduct evaluation and does not incorporate subsequent amendments/additions. | Prevents treating newer IBC editions as automatically incorporated for this professional-conduct fallback. | src:usa-mo:csr-20-2030-2 | verified_limited_scope |
| amendment:usa-mo:state-energy-latest-editions | Energy | 10 CSR 140-7.010 defines the minimum energy efficiency standard by latest ASHRAE 90.1 / ASHRAE 90.2 or CABO MEC editions, with a three-month publication trigger. | Requires dynamic model-standard version resolution for covered state projects. | src:usa-mo:csr-10-140-7-010 | partially_verified |
| amendment:usa-mo:elevator-date-bands | Elevator / Conveyance | DFS applies different ASME/ANSI standards by installation date and equipment type. | Requires equipment installation date for code lookup. | src:usa-mo:dps-dfs-elevator-code-adoptions | verified |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-mo"
  model: "local_adoption_and_enforcement_with_limited_state_programs"
  enforcing_entities:
    - "municipalities"
    - "counties where authorized and where a code/order/ordinance exists"
    - "fire protection districts within statutory limits"
    - "Missouri Division of Fire Safety for elevator/conveyance and certain licensed-facility fire-safety programs"
    - "state agencies for state-building energy compliance"
  required_officials:
    - "local building official or local code official where adopted"
    - "local fire code official where adopted"
    - "DFS elevator permit/inspection/certificate process for elevators and related equipment"
  state_reserved_activities:
    - "elevator and related-equipment standards and operating certificates"
    - "state-building minimum energy efficiency standard"
    - "professional conduct evaluation by licensing board"
    - "licensed-facility fire-safety inspections/rules"
    - "seismic ordinance mandate for identified New Madrid earthquake-risk local governments"
  source_ids:
    - "src:usa-mo:dnr-codes-by-jurisdiction"
    - "src:usa-mo:rsmo-67-280"
    - "src:usa-mo:rsmo-64-170"
    - "src:usa-mo:rsmo-89-020"
    - "src:usa-mo:rsmo-321-228"
    - "src:usa-mo:dfs-fire-faq"
    - "src:usa-mo:dps-dfs-elevator-code-adoptions"
    - "src:usa-mo:dps-dfs-elevator-code-info"
  verification_status: "partially_verified"
  confidence: 0.66
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-mo"
  model: "local_amendments_allowed_where_local_code_adoption_power_exists"
  applies_to_code_families:
    - "construction and occupancy"
    - "mechanical"
    - "plumbing"
    - "electrical"
    - "fire prevention"
    - "other local technical codes if authorized elsewhere"
  approval_required: "no statewide approval requirement identified"
  approving_authority_id: null
  filing_required: true
  filing_rule: "At least one copy of the adopted code, portion, or amendment must be filed with the community clerk and available for public use, inspection, and examination for 90 days before adoption by reference."
  registry_exists: "partial_general_code_lists_only"
  registry_source_ids:
    - "src:usa-mo:dnr-codes-by-jurisdiction"
  legal_basis_source_ids:
    - "src:usa-mo:rsmo-67-280"
  verification_status: "partially_verified"
  confidence: 0.58
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Local enforcement and local amendment authority are related but separate. Missouri DNR's adopted-code listings support a local-code lookup workflow, but they do not provide full ordinance text, amendments, trade-code adoptions, permit-date transition rules, or enforcement contacts. RSMo 67.280 supports local adoption by reference only where the community otherwise has legal authority. Therefore an address-level resolver must not infer local code authority from the DNR list alone; it must reconcile local ordinances, county class/status, municipal boundaries, fire protection district boundaries, and any state-reserved programs.

### 6.4 Known Local Amendment Registries

| Registry ID | Registry Name | Maintainer | Coverage | Includes Amendments? | Includes Trade Codes? | Source IDs | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| registry:usa-mo:dnr-county-city-code-lists | Missouri Energy Codes by Jurisdiction / adopted building codes by county and city | Missouri Department of Natural Resources | County and city adopted-code lists | no, per DNR page caveat | no, per DNR page caveat | src:usa-mo:dnr-codes-by-jurisdiction | verified_caveated |
| registry:usa-mo:local-clerk-code-files | Clerk-filed code and amendment copies | Each adopting community | Local community code/adoption records | yes, if locally filed and adopted | yes, if locally filed and adopted | src:usa-mo:rsmo-67-280 | partially_verified_distributed |

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendments were parsed in this draft. DNR-provided county/city adopted-code lists should be ingested as a discovery layer, then local ordinances should be parsed for legally operative code edition, amendments, permit-date transition rules, and enforcement contact data.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

**Resolver status:** state_level_partially_ready; local_boundary_and_ordinance_ingestion_required

Jurisdiction stack:

```text
Address
  -> State: Missouri
  -> County
  -> Municipality or unincorporated county area
  -> County classification and county code-adoption/voter-approval status
  -> Fire protection district and local fire code status
  -> Local building/residential/mechanical/plumbing/electrical/energy/fire ordinances
  -> State special programs:
       - elevator / conveyance
       - state-building energy standard
       - licensed-facility fire-safety rules
       - New Madrid seismic special applicability
       - professional-conduct fallback if no local building code exists
  -> Adopted local code records
  -> Local amendment records
  -> AHJ contacts and permit transition rules
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- |
| State | U.S. Census TIGER/Line or equivalent authoritative boundary source | none | statewide | annual typical; source selection pending | pending |
| County | U.S. Census TIGER/Line counties plus Missouri county class metadata | none | statewide | annual typical; class metadata pending | pending |
| Municipality | U.S. Census TIGER/Line places or state/local municipal boundary source | none | statewide | annual typical; local verification pending | pending |
| Fire District | Missouri fire protection district boundary source not selected | none | statewide if available | unknown | pending |
| Seismic covered jurisdictions | State geologist / USGS / SEMA notification list | none | New Madrid affected jurisdictions | unknown | pending |
| Local code jurisdiction lists | Missouri DNR / Missouri Data Portal county and city adopted-code lists | src:usa-mo:dnr-codes-by-jurisdiction | county/city adopted-code discovery | unknown | source_identified_not_ingested |

### 7.3 AHJ Contact Data

No AHJ contact data was populated. Missouri DNR warns that current code details should be obtained from the local jurisdiction, so AHJ contacts are required before address-level responses can be treated as production-grade.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID | Title | Publisher / Maintainer | Source Type | URL | Accessed | Key Fields Supported | Caveats |
| --- | --- | --- | --- | --- | --- | --- | --- |
| src:usa-mo:dnr-codes-by-jurisdiction | Energy Codes by Jurisdiction | Missouri Department of Natural Resources, Division of Energy | agency_html | https://dnr.mo.gov/energy/efficiency/codes-jurisdiction | 2026-06-26 | Local code adoption discovery; county/city code-list caveat | DNR states lists are general, may omit amendments and trade codes, and local jurisdiction should be contacted for current details. |
| src:usa-mo:rsmo-67-280 | RSMo 67.280, Communities may incorporate by reference certain technical codes | Missouri Revisor of Statutes | statute_html | https://revisor.mo.gov/main/OneSection.aspx?section=67.280 | 2026-06-26 | Local technical-code adoption by reference; community definition; 90-day filing rule | Revisor site states online statutes are uncertified and unofficial. |
| src:usa-mo:rsmo-64-170 | RSMo 64.170, County commissions control construction | Missouri Revisor of Statutes | statute_html | https://revisor.mo.gov/main/OneSection.aspx?section=64.170 | 2026-06-26 | First/second class county building, electrical, plumbing, permit, inspection, fee, and voter-approval authority | Revisor site states online statutes are uncertified and unofficial. |
| src:usa-mo:rsmo-89-020 | RSMo 89.020, Powers of municipal legislative body | Missouri Revisor of Statutes | statute_html | https://revisor.mo.gov/main/OneSection.aspx?section=89.020 | 2026-06-26 | Municipal authority over building height/stories/size and land-use/zoning controls | Zoning authority is not identical to full building-code authority; Revisor caveat applies. |
| src:usa-mo:rsmo-321-228 | RSMo 321.228, Residential construction regulatory system and fire protection districts | Missouri Revisor of Statutes | statute_html | https://revisor.mo.gov/main/OneSection.aspx?section=321.228 | 2026-06-26 | Residential construction regulatory systems; fire district preemption/retained authority | Revisor site states online statutes are uncertified and unofficial. |
| src:usa-mo:dfs-fire-faq | Frequently Asked Questions: Does Missouri have a statewide fire code? | Missouri Department of Public Safety, Division of Fire Safety | agency_html | https://dfs.dps.mo.gov/about/faqs.php | 2026-06-26 | No statewide fire code; licensed-facility caveat; local fire code confirmation | FAQ page is authoritative agency guidance but not a codified rule. |
| src:usa-mo:csr-20-2030-2 | 20 CSR 2030-2.040, Evaluation Criteria for Building Design | Missouri Secretary of State, Code of State Regulations | csr_pdf | https://www.sos.mo.gov/cmsimages/adrules/csr/current/20csr/20c2030-2.pdf | 2026-06-26 | 2018 IBC professional-conduct evaluation criteria where no local building code exists; effective history | PDF includes incorporated material note; CSR current-version caveat applies. |
| src:usa-mo:dps-dfs-elevator-code-adoptions | Code Adoptions - Elevators and Related Equipment | Missouri Department of Public Safety, Division of Fire Safety | agency_html | https://dfs.dps.mo.gov/programs/elevator/2007-code-adoption.php | 2026-06-26 | Elevator/equipment standards by installation date; ASME A17.1 2016 after 2020-04-30; ANSI A117.1 2009 for new passenger elevators/platform lifts | Agency page summarizes adopted standards; confirm against CSR for production use. |
| src:usa-mo:dps-dfs-elevator-code-info | Elevator Safety Code Information | Missouri Department of Public Safety, Division of Fire Safety | agency_html | https://dfs.dps.mo.gov/programs/elevator/elevator-safety-code.php | 2026-06-26 | Elevator code info, permits, operating certificates, local ordinance caveat | Agency page should be reconciled with current CSR and permit forms. |
| src:usa-mo:rsmo-8-812 | RSMo 8.812, Minimum energy efficiency standards for state buildings | Missouri Revisor of Statutes | statute_html | https://revisor.mo.gov/main/OneSection.aspx?section=8.812 | 2026-06-26 | State-building energy standard statutory basis; 2009-07-01 design trigger; exemptions | Revisor site states online statutes are uncertified and unofficial. |
| src:usa-mo:csr-10-140-7-010 | 10 CSR 140-7.010, State Building Minimum Energy Efficiency Standards | Missouri Secretary of State, Code of State Regulations | csr_pdf | https://www.sos.mo.gov/cmsimages/adrules/csr/current/10csr/10c140-7.pdf | 2026-06-26 | Minimum energy efficiency standard definition; ASHRAE 90.1/90.2/CABO MEC latest; three-month update rule; applicability | CSR current-version caveat applies; incorporated standards are not reproduced in full. |
| src:usa-mo:rsmo-319-200 | RSMo 319.200, seismic construction and renovation ordinances | Missouri Revisor of Statutes | statute_html | https://revisor.mo.gov/main/OneSection.aspx?section=319.200 | 2026-06-26 | New Madrid seismic ordinance obligation and model-code standard references | Requires separate list of covered jurisdictions from state geologist/USGS/SEMA. |
| src:usa-mo:rsmo-319-203 | RSMo 319.203, seismic ordinance applicability and exceptions | Missouri Revisor of Statutes | statute_html | https://revisor.mo.gov/main/OneSection.aspx?section=319.203 | 2026-06-26 | Seismic applicability for leases/buildings and exclusions | Revisor site states online statutes are uncertified and unofficial. |
| src:usa-mo:revisor-certification-status | Missouri Revisor homepage certification status | Missouri Revisor of Statutes | official_site_caveat | https://revisor.mo.gov/ | 2026-06-26 | Caveat that online statutes are uncertified and unofficial; effective-date update note | Use certified publication for legal-grade verification. |
| src:usa-mo:sos-csr-current-caveat | Code of State Regulations current-version caveat | Missouri Secretary of State | official_site_caveat | https://www.sos.mo.gov/adrules/csr/current/19csr/19csr | 2026-06-26 | CSR updated monthly; some current rules may not yet be effective; rule effective date controls | Page is a general CSR caveat; apply to current CSR PDFs. |

### 8.2 Official-Source Caveats

| Source ID | Caveat Type | Caveat Summary | Production Treatment |
| --- | --- | --- | --- |
| src:usa-mo:revisor-certification-status | uncertified_online_statutes | The Revisor site states that statutes posted online are uncertified and unofficial, despite being the official website. | Use for draft/partial verification; cite certified statutory publications for verified legal release. |
| src:usa-mo:sos-csr-current-caveat | csr_current_version_effective_date | SOS states the CSR is updated monthly and some rules in the current version may not yet be effective; rule authority/effective history must be checked. | Store rule effective date separately from page access date; compare archived editions where needed. |
| src:usa-mo:dnr-codes-by-jurisdiction | local_code_list_scope_limit | DNR lists general adopted codes but not specific amendments or trade codes; local jurisdiction has current details. | Treat as discovery index, not final legal authority for address-level output. |
| src:usa-mo:csr-20-2030-2 | incorporated_model_code_not_reproduced | 2018 IBC is incorporated by reference and is not reproduced in full in the CSR PDF. | Store incorporation metadata; do not attempt to reproduce copyrighted model-code text. |
| src:usa-mo:csr-10-140-7-010 | incorporated_standards_not_reproduced | ASHRAE/CABO standards are incorporated by reference and not reproduced in full. | Store reference metadata and external standard edition resolution logic. |

### 8.3 Supplemental Sources

None used as binding authority in this report. Secondary summaries may be useful later for completeness checks, but official Missouri sources control the facts captured here.

### 8.4 Source Extraction Metadata

| Source ID | Accessed | Extracted By | Extraction Method | Confidence | Notes |
| --- | --- | --- | --- | --- | --- |
| src:usa-mo:dnr-codes-by-jurisdiction | 2026-06-26 | GPT-5.5 Thinking | html_text_extract | 0.80 | Page content supports discovery registry and caveat, not individual local ordinance records. |
| src:usa-mo:rsmo-67-280 | 2026-06-26 | GPT-5.5 Thinking | html_text_extract | 0.86 | Core local adoption-by-reference statute. |
| src:usa-mo:rsmo-64-170 | 2026-06-26 | GPT-5.5 Thinking | html_text_extract | 0.82 | County authority; additional county-class statutes remain pending. |
| src:usa-mo:rsmo-89-020 | 2026-06-26 | GPT-5.5 Thinking | html_text_extract | 0.74 | Municipal zoning/building-size authority; not full construction-code authority by itself. |
| src:usa-mo:rsmo-321-228 | 2026-06-26 | GPT-5.5 Thinking | html_text_extract | 0.82 | Residential/fire protection district relationship. |
| src:usa-mo:dfs-fire-faq | 2026-06-26 | GPT-5.5 Thinking | html_text_extract | 0.86 | Clear agency answer on statewide fire code. |
| src:usa-mo:csr-20-2030-2 | 2026-06-26 | GPT-5.5 Thinking | pdf_text_extract_plus_visual_spot_check | 0.84 | Official CSR PDF; page screenshots reviewed for rule text. |
| src:usa-mo:dps-dfs-elevator-code-adoptions | 2026-06-26 | GPT-5.5 Thinking | html_text_extract | 0.86 | State elevator installation date bands. |
| src:usa-mo:dps-dfs-elevator-code-info | 2026-06-26 | GPT-5.5 Thinking | html_text_extract | 0.70 | URL should be rechecked during production ingestion. |
| src:usa-mo:rsmo-8-812 | 2026-06-26 | GPT-5.5 Thinking | html_text_extract | 0.84 | State-building energy statutory trigger. |
| src:usa-mo:csr-10-140-7-010 | 2026-06-26 | GPT-5.5 Thinking | pdf_text_extract_plus_visual_spot_check | 0.84 | Official CSR PDF; page screenshots reviewed for rule text. |
| src:usa-mo:rsmo-319-200 | 2026-06-26 | GPT-5.5 Thinking | html_text_extract | 0.78 | Requires covered-jurisdiction list for production. |
| src:usa-mo:rsmo-319-203 | 2026-06-26 | GPT-5.5 Thinking | html_text_extract | 0.80 | Seismic applicability and exceptions. |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID | Field | Value | Verification Status | Confidence | Source IDs | Notes |
| --- | --- | --- | --- | --- | --- | --- |
| report | report.status | partially_verified | verified | 1.00 | none | Core state-level authority sources captured; address-level/local ordinance coverage remains incomplete. |
| report | risk.overall_confidence | 0.62 | verified | 1.00 | none | Moderate confidence for state-level model, lower for local implementation. |
| ahj:usa-mo:local-communities | authority model | local adoption and enforcement | partially_verified | 0.72 | src:usa-mo:dnr-codes-by-jurisdiction; src:usa-mo:rsmo-67-280 | Still needs local ordinance ingestion. |
| adoption:usa-mo:professional-conduct-ibc-2018 | edition | 2018 IBC | verified_limited_scope | 0.82 | src:usa-mo:csr-20-2030-2 | Not a statewide building code adoption. |
| adoption:usa-mo:state-building-energy-standard | statutory trigger | 2009-07-01 | verified | 0.84 | src:usa-mo:rsmo-8-812 | Applies only to covered state buildings/acquisitions/leases. |
| adoption:usa-mo:elevator-asme-a17-1-2016 | operative date | 2020-04-30 | verified | 0.86 | src:usa-mo:dps-dfs-elevator-code-adoptions | Applies to elevators installed after this date. |
| fire | statewide fire code | no statewide fire code | verified | 0.86 | src:usa-mo:dfs-fire-faq | Facility-specific rules remain pending. |
| local-amendment-rule:usa-mo | filing_required | true, 90-day local clerk filing | verified | 0.80 | src:usa-mo:rsmo-67-280 | Statewide registry not identified. |
| seismic | covered jurisdiction list | not populated | unresolved | 0.30 | src:usa-mo:rsmo-319-200; src:usa-mo:rsmo-319-203 | Need SEMA/geologist/USGS list. |
| accessibility | general statewide accessibility code posture | unresolved | unresolved | 0.25 | src:usa-mo:dps-dfs-elevator-code-adoptions | Elevator-related accessibility captured only. |

### 9.2 Validation Checks

| Check | Status | Notes |
| --- | --- | --- |
| All source IDs resolve | pass | Every `src:usa-mo:*` ID used outside the source registry is listed in section 8. |
| All authority IDs resolve | pass | Authority IDs used in sections 2 and 6 are defined or self-descriptive. |
| All current code families have adoption records or explicit disposition rows | pass | Matrix rows distinguish local-adoption rows from limited state-scope rows. |
| Building and operational fire code are separated | pass | Fire construction references and operational/prevention code are separated. |
| Adoption/effective/operative/mandatory dates are not conflated | pass | Separate date fields are preserved. |
| Effective dates are valid ISO dates | pass | Known date fields use ISO format; unknown dates are null. |
| No impossible date sequences | pass | No contradictory date sequence identified. |
| Transition rules have explicit trigger conditions | pass | Confirmed state-level rules have triggers; local rules remain jurisdiction-specific. |
| Permit-date logic is captured where applicable | partial | Local permit-date rules require local ordinance ingestion. |
| Local enforcement model classified | pass | Classified as local adoption/enforcement with limited state programs. |
| Local amendment rule classified | partial | General 90-day filing rule captured; full registry/approval map unresolved. |
| AHJ confirmation metadata present | fail | AHJ contact data not populated. |
| Official-source caveats captured | pass | Revisor, CSR, DNR, and incorporated-standard caveats are recorded. |
| File free of template-residue markers checked by validation regex | pass | Marker scan was run after generation and returned clean. |

---

## 10. Open Issues and Review Queue

| Issue ID | Severity | Topic | Description | Needed Action | Owner | Due Date | Status |
| --- | --- | --- | --- | --- | --- | --- | --- |
| issue:usa-mo:001 | high | local ordinance ingestion | DNR county/city lists do not include amendments, trade codes, permit-date rules, or full local ordinance text. | Ingest DNR/Missouri Data Portal lists, then scrape or manually capture local ordinance sources for MVP jurisdictions. | null | null | open |
| issue:usa-mo:002 | high | AHJ contacts | No local AHJ contact records are populated. | Add building, planning, fire, and trade contact sources by municipality/county. | null | null | open |
| issue:usa-mo:003 | high | fire licensed-facility rules | DFS FAQ confirms state rules for licensed facilities, but those rule sets were not parsed. | Parse day care, long-term care, health care, school, lodging, and other licensed-facility fire-safety rules. | null | null | open |
| issue:usa-mo:004 | medium | county authority completeness | RSMo 64.170 covers first/second class county authority, but county authority for all classes and special forms of government remains incomplete. | Parse Chapters 64 and 65 plus county classification/current county status data. | null | null | open |
| issue:usa-mo:005 | medium | seismic covered jurisdictions | Seismic statutes require local ordinances in identified New Madrid risk jurisdictions, but the covered-jurisdiction list was not captured. | Locate state geologist/USGS/SEMA notice/list and map it to county/municipal boundaries. | null | null | open |
| issue:usa-mo:006 | medium | accessibility | General Missouri accessibility/architectural-barrier rules were not resolved beyond elevator-related ANSI A117.1 requirement. | Review Missouri accessibility statutes/rules and state/federal overlay for public accommodations and state facilities. | null | null | open |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID | Source ID | Watch Type | Frequency | Trigger Condition | Last Checked | Status |
| --- | --- | --- | --- | --- | --- | --- |
| watch:usa-mo:dnr-code-lists | src:usa-mo:dnr-codes-by-jurisdiction | html_diff_and_dataset_diff | monthly | New county/city code-list links, changed DNR caveats, or updated adopted-code data | 2026-06-26 | active |
| watch:usa-mo:rsmo-local-adoption | src:usa-mo:rsmo-67-280 | statute_diff | quarterly | Amendment to local technical-code adoption by reference rule | 2026-06-26 | active |
| watch:usa-mo:rsmo-county-authority | src:usa-mo:rsmo-64-170 | statute_diff | quarterly | Amendment to county building-code authority or voter-approval condition | 2026-06-26 | active |
| watch:usa-mo:dfs-fire-faq | src:usa-mo:dfs-fire-faq | html_diff | monthly | DFS changes statewide fire-code answer or licensed-facility guidance | 2026-06-26 | active |
| watch:usa-mo:elevator | src:usa-mo:dps-dfs-elevator-code-adoptions | html_diff_and_csr_diff | monthly | Updated ASME/ANSI adopted edition, permit rule, or operating certificate requirement | 2026-06-26 | active |
| watch:usa-mo:state-energy | src:usa-mo:csr-10-140-7-010 | csr_pdf_diff | monthly | CSR update or referenced energy-standard update | 2026-06-26 | active |
| watch:usa-mo:professional-board | src:usa-mo:csr-20-2030-2 | csr_pdf_diff | monthly | Updated IBC evaluation edition or professional-conduct rule | 2026-06-26 | active |
| watch:usa-mo:seismic | src:usa-mo:rsmo-319-200 | statute_diff | quarterly | Amendment to seismic design/construction ordinance requirements | 2026-06-26 | active |
| watch:usa-mo:revisor-certification | src:usa-mo:revisor-certification-status | html_diff | quarterly | Change in online statute certification status or supplement notice | 2026-06-26 | active |
| watch:usa-mo:sos-csr-caveat | src:usa-mo:sos-csr-current-caveat | html_diff | quarterly | Change in CSR publication/effective-date caveat | 2026-06-26 | active |

### 11.2 Changelog

| Date | Change | Record IDs | Source IDs | Changed By | Notes |
| --- | --- | --- | --- | --- | --- |
| 2026-06-24 | Generated baseline Missouri draft report | report:usa-mo | none | Codex | Baseline had no primary state sources. |
| 2026-06-26 | Populated Missouri report with official state sources and partial verification | report:usa-mo; ahj:usa-mo:local-communities; adoption:usa-mo:local-technical-code-enabling; adoption:usa-mo:professional-conduct-ibc-2018; adoption:usa-mo:state-building-energy-standard; adoption:usa-mo:elevator-asme-a17-1-2016; adoption:usa-mo:seismic-new-madrid-local-ordinances | src:usa-mo:dnr-codes-by-jurisdiction; src:usa-mo:rsmo-67-280; src:usa-mo:rsmo-64-170; src:usa-mo:rsmo-89-020; src:usa-mo:rsmo-321-228; src:usa-mo:dfs-fire-faq; src:usa-mo:csr-20-2030-2; src:usa-mo:dps-dfs-elevator-code-adoptions; src:usa-mo:rsmo-8-812; src:usa-mo:csr-10-140-7-010; src:usa-mo:rsmo-319-200; src:usa-mo:rsmo-319-203 | GPT-5.5 Thinking | Upgraded to partially_verified for state-level model only; local ordinance and AHJ gaps remain explicit. |
