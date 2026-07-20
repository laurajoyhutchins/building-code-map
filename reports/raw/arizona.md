---
state:
  state_id: "US-AZ"
  name: "Arizona"
  abbreviation: "AZ"
report:
  report_id: "state-report:usa-az"
  schema_version: "0.3.0"
  status: "partially_verified"
  last_updated: "2026-06-25"
  last_verified: "2026-06-25"
  reviewed_by: null
risk:
  overall_confidence: 0.78
  risk_flags:
    - "no_single_statewide_private_building_code_verified"
    - "local_code_editions_vary_by_jurisdiction"
    - "local_adoption_records_not_parsed"
    - "state_fire_code_verified_but_modifications_not_fully_parsed"
    - "public_building_fallback_requires_county_largest_city_lookup"
open_questions_count: 4
---

# State Building Code Authority Report: Arizona

## 1. Executive Summary

- **Authority model:** Arizona is primarily a local-adoption state for ordinary building, residential, mechanical, plumbing, electrical, and related construction codes. Cities and towns may adopt codes by reference. Counties may adopt and enforce building and related codes for zoned unincorporated areas, subject to statutory limits.

- **Statewide building code status:** No single statewide IBC/IRC-style building code for ordinary private construction was verified. Statewide rules do exist for specific scopes, especially the State Fire Code, public buildings, and elevator/conveyance safety.

- **State Fire Code:** Verified. Arizona Administrative Code R4-36-201 incorporates the **2024 International Fire Code** as the Arizona State Fire Code, with specified appendices and state modifications in Article 3. The rule states that the incorporated material does not include later amendments or editions.

- **Public buildings:** Verified special rule. Public buildings generally follow the applicable building, plumbing, electrical, fire, and mechanical codes adopted by the city, town, county, or fire district where the building is located. If no local codes have been adopted, the public building must be designed or constructed according to the State Fire Code and the building, plumbing, electrical, and mechanical codes that apply in the largest city in the county.

- **Fire authority:** The Office of the State Fire Marshal is established within the Arizona Department of Forestry and Fire Management and has authority to adopt and enforce a state fire code, with statutory exceptions and local-assumption rules. Fire districts may adopt nationally recognized fire codes approved by the State Fire Marshal after voter approval and may amend or replace their adopted fire code with State Fire Marshal approval.

- **Elevator / conveyance code:** Verified statewide specialty code. Arizona Administrative Code R20-5-507 requires covered conveyances installed, repaired, or altered on or after January 1, 2023 to comply with ASME A17.1-2019 or ASME A17.7-2007 as referenced in ASME A17.1-2019.

- **Production readiness:** state_level_ready_for_authority_model; not_ready_for_address_specific_code_resolution.

### Key Findings

```yaml
---
key_findings:
- topic: State adopting authority
  finding: No single statewide private building code adoption was verified; ordinary
    construction codes are primarily adopted locally by municipalities and counties.
  confidence: 0.82
  source_ids:
  - src:usa-az:ars:9-801
  - src:usa-az:ars:9-802
  - src:usa-az:ars:11-861
- topic: Municipal code adoption
  finding: Municipalities may adopt a code or public record by reference; "code" includes
    building, electrical wiring, fire prevention, wildland-urban interface, and other
    proper subjects of municipal legislation.
  confidence: 0.92
  source_ids:
  - src:usa-az:ars:9-801
  - src:usa-az:ars:9-802
- topic: County code adoption
  finding: Counties with zoning may adopt and enforce building and related codes for
    zoned unincorporated areas; county codes are limited to nationally adopted building/electrical/plumbing/mechanical
    codes or the code adopted by the largest city in the county.
  confidence: 0.93
  source_ids:
  - src:usa-az:ars:11-861
- topic: State Fire Code
  finding: Arizona has adopted the 2024 International Fire Code as the State Fire
    Code, with specified appendices and state modifications.
  confidence: 0.95
  source_ids:
  - src:usa-az:aac:r4-36-201
- topic: Public buildings
  finding: "Public buildings follow applicable local building, plumbing, electrical,\
    \ fire, and mechanical codes; if no local codes exist, they use the State Fire\
    \ Code plus the largest city in the county's building, plumbing, electrical,\
    \ and mechanical codes."
  confidence: 0.95
  source_ids:
  - src:usa-az:ars:34-461
- topic: Fire districts
  finding: Fire districts may adopt a nationally recognized fire code approved by
    the State Fire Marshal after voter approval, and amendments/revisions require
    State Fire Marshal approval.
  confidence: 0.9
  source_ids:
  - src:usa-az:ars:48-805
- topic: Elevator / conveyance
  finding: Covered conveyances installed, repaired, or altered on or after January
    1, 2023 must comply with ASME A17.1-2019 or ASME A17.7-2007 as referenced in ASME
    A17.1-2019.
  confidence: 0.92
  source_ids:
  - src:usa-az:aac:r20-5-507
- topic: Local amendments
  finding: Municipal code amendments follow the same adoption-by-reference process.
    County and fire-district amendments are subject to additional statutory limits.
  confidence: 0.84
  source_ids:
  - src:usa-az:ars:9-802
  - src:usa-az:ars:11-861
  - src:usa-az:ars:48-805
- topic: Effective / transition rules
  finding: Verified for State Fire Code, elevator/conveyance, public buildings, municipal
    adoption-by-reference procedure, and county adoption/update mechanics; not fully
    resolved for each local ordinance.
  confidence: 0.76
  source_ids:
  - src:usa-az:aac:r4-36-201
  - src:usa-az:aac:r20-5-507
  - src:usa-az:ars:34-461
  - src:usa-az:ars:9-802
  - src:usa-az:ars:11-861
```

---

### 2.1 Primary Building Code Authorities

| Authority ID                            | Authority Name                       | Authority Type      | Legal Basis                                  | Role                                                                                                                                                    | Enforcement Model                    | Source IDs                                                            | Verification Status |
| --------------------------------------- | ------------------------------------ | ------------------- | -------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------ | --------------------------------------------------------------------- | ------------------- |
| ahj:usa-az:municipal-code-adoption      | Municipal code adoption by reference | statute             | A.R.S. §§ 9-801, 9-802                       | Enables municipalities to adopt codes by reference, including building, electrical, fire prevention, WUI, and related codes                             | local_municipal                      | src:usa-az:ars:9-801, src:usa-az:ars:9-802                              | verified            |
| ahj:usa-az:county-code-adoption         | County code adoption and enforcement | statute             | A.R.S. §§ 11-861, 11-862                     | Enables counties to adopt and enforce building and related codes in zoned unincorporated areas, with advisory-board requirements and code-source limits | county_unincorporated_local_option   | src:usa-az:ars:11-861, src:usa-az:ars:11-862                            | verified            |
| ahj:usa-az:state-fire-marshal           | Office of the State Fire Marshal     | statute/rule        | A.R.S. §§ 37-1381, 37-1383; A.A.C. R4-36-201 | Adopts and enforces the State Fire Code, with statutory local exceptions                                                                                | statewide_fire_with_local_exceptions | src:usa-az:ars:37-1381, src:usa-az:ars:37-1383, src:usa-az:aac:r4-36-201 | verified            |
| ahj:usa-az:public-building-code-rule    | Public building code applicability   | statute             | A.R.S. § 34-461                              | Determines code compliance for public buildings and fallback codes where no local codes exist                                                           | special_state_public_building_rule   | src:usa-az:ars:34-461                                                  | verified            |
| ahj:usa-az:fire-district-code-authority | Fire district fire-code adoption     | statute             | A.R.S. § 48-805                              | Allows fire districts to adopt nationally recognized fire codes approved by the State Fire Marshal after voter approval                                 | special_district_fire                | src:usa-az:ars:48-805                                                  | verified            |
| ahj:usa-az:elevator-conveyance-safety   | Elevator and conveyance safety       | administrative_rule | A.A.C. R20-5-507                             | Establishes statewide conveyance safety standards                                                                                                       | statewide_specialty_code             | src:usa-az:aac:r20-5-507                                               | verified            |

### 2.2 Specialized Code Authorities

| Code Family                          | Authority ID                                                                                                                                | Authority Name                                               | Role                                                                                                           | Legal Basis                                                         | Source IDs                                                                                                      | Status               |
| ------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------ | -------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | -------------------- |
| Building                             | ahj:usa-az:municipal-code-adoption; ahj:usa-az:county-code-adoption; ahj:usa-az:public-building-code-rule                                   | Municipalities, counties, and public-building fallback       | Local adoption for ordinary construction; public-building fallback rule                                        | A.R.S. §§ 9-801, 9-802, 11-861, 34-461                              | src:usa-az:ars:9-801, src:usa-az:ars:9-802, src:usa-az:ars:11-861, src:usa-az:ars:34-461                            | verified_state_level |
| Residential                          | ahj:usa-az:municipal-code-adoption; ahj:usa-az:county-code-adoption                                                                         | Municipalities and counties                                  | Local adoption; residential sprinkler preemptions apply                                                        | A.R.S. §§ 9-801, 9-802, 9-807, 9-808, 11-861                        | src:usa-az:ars:9-801, src:usa-az:ars:9-802, src:usa-az:ars:9-807, src:usa-az:ars:9-808, src:usa-az:ars:11-861        | verified_state_level |
| Existing Building / Rehabilitation   | ahj:usa-az:municipal-code-adoption; ahj:usa-az:county-code-adoption                                                                         | Municipalities and counties                                  | Local adoption if included in adopted local code set                                                           | A.R.S. §§ 9-801, 9-802, 11-861                                      | src:usa-az:ars:9-801, src:usa-az:ars:9-802, src:usa-az:ars:11-861                                                  | local_records_needed |
| Mechanical                           | ahj:usa-az:municipal-code-adoption; ahj:usa-az:county-code-adoption; ahj:usa-az:public-building-code-rule                                   | Municipalities, counties, public-building fallback           | Local adoption; public-building fallback                                                                       | A.R.S. §§ 9-801, 9-802, 11-861, 34-461                              | src:usa-az:ars:9-801, src:usa-az:ars:9-802, src:usa-az:ars:11-861, src:usa-az:ars:34-461                            | verified_state_level |
| Plumbing                             | ahj:usa-az:municipal-code-adoption; ahj:usa-az:county-code-adoption; ahj:usa-az:public-building-code-rule                                   | Municipalities, counties, public-building fallback           | Local adoption; public-building fallback                                                                       | A.R.S. §§ 9-801, 9-802, 11-861, 34-461                              | src:usa-az:ars:9-801, src:usa-az:ars:9-802, src:usa-az:ars:11-861, src:usa-az:ars:34-461                            | verified_state_level |
| Fuel Gas                             | ahj:usa-az:municipal-code-adoption; ahj:usa-az:county-code-adoption                                                                         | Municipalities and counties                                  | Local adoption if included in adopted mechanical/plumbing/fuel-gas code set                                    | A.R.S. §§ 9-801, 9-802, 11-861                                      | src:usa-az:ars:9-801, src:usa-az:ars:9-802, src:usa-az:ars:11-861                                                  | local_records_needed |
| Electrical                           | ahj:usa-az:municipal-code-adoption; ahj:usa-az:county-code-adoption; ahj:usa-az:public-building-code-rule                                   | Municipalities, counties, public-building fallback           | Local adoption; public-building fallback                                                                       | A.R.S. §§ 9-801, 9-802, 11-861, 34-461                              | src:usa-az:ars:9-801, src:usa-az:ars:9-802, src:usa-az:ars:11-861, src:usa-az:ars:34-461                            | verified_state_level |
| Energy                               | ahj:usa-az:municipal-code-adoption                                                                                                          | Municipalities primarily; local records required             | No statewide energy baseline verified in this pass                                                             | A.R.S. §§ 9-801, 9-802                                              | src:usa-az:ars:9-801, src:usa-az:ars:9-802                                                                        | unresolved_local     |
| Fire - construction references       | ahj:usa-az:state-fire-marshal; ahj:usa-az:municipal-code-adoption; ahj:usa-az:county-code-adoption; ahj:usa-az:fire-district-code-authority | State Fire Marshal, municipalities, counties, fire districts | State Fire Code plus local fire-code authority where allowed                                                   | A.R.S. §§ 37-1381, 37-1383, 9-808, 11-861, 48-805; A.A.C. R4-36-201 | src:usa-az:ars:37-1381, src:usa-az:ars:37-1383, src:usa-az:aac:r4-36-201, src:usa-az:ars:48-805                     | verified             |
| Fire - operational / prevention code | ahj:usa-az:state-fire-marshal; ahj:usa-az:fire-district-code-authority; ahj:usa-az:county-code-adoption; ahj:usa-az:municipal-code-adoption | State Fire Marshal, fire districts, counties, municipalities | State Fire Code and local fire-prevention codes subject to statutory exceptions/preemptions                    | A.R.S. §§ 37-1383, 48-805, 11-861, 9-808; A.A.C. R4-36-201          | src:usa-az:ars:37-1383, src:usa-az:aac:r4-36-201, src:usa-az:ars:48-805, src:usa-az:ars:11-861, src:usa-az:ars:9-808 | verified             |
| Accessibility                        | ahj:usa-az:municipal-code-adoption; ahj:usa-az:county-code-adoption                                                                         | Local adopting authorities                                   | No separate state accessibility code was verified in this pass; federal and local-code analysis still required | A.R.S. §§ 9-801, 9-802, 11-861                                      | src:usa-az:ars:9-801, src:usa-az:ars:9-802, src:usa-az:ars:11-861                                                  | unresolved           |
| Elevator / Conveyance                | ahj:usa-az:elevator-conveyance-safety                                                                                                       | Industrial Commission / ADOSH Elevator Safety rules          | Statewide conveyance safety standards                                                                          | A.A.C. R20-5-507                                                    | src:usa-az:aac:r20-5-507                                                                                         | verified             |

### 2.3 Authority Hierarchy Notes

Arizona should be modeled as a **local construction-code state with statewide specialty overlays**. The verified statewide overlay is strongest for fire protection and conveyances. For ordinary private construction, the controlling building, residential, mechanical, plumbing, electrical, fuel-gas, energy, and existing-building editions generally must be resolved at the municipality or unincorporated-county level.

For public buildings, Arizona has a separate statutory fallback: use applicable local building, plumbing, electrical, fire, and mechanical codes where adopted; otherwise use the State Fire Code and the building, plumbing, electrical, and mechanical codes that apply in the largest city in the county.

### 2.4 Authority Graph Edges

| Edge ID        | From Authority                          | Relationship                   | To Authority / Scope                                        | Source IDs                                     | Status   |
| -------------- | --------------------------------------- | ------------------------------ | ----------------------------------------------------------- | ---------------------------------------------- | -------- |
| edge:usa-az:001 | ahj:usa-az:municipal-code-adoption      | authorizes                     | municipal_code_adoption_by_reference                        | src:usa-az:ars:9-801, src:usa-az:ars:9-802       | verified |
| edge:usa-az:002 | ahj:usa-az:municipal-code-adoption      | allows_amendment_same_manner   | municipal_codes_adopted_by_reference                        | src:usa-az:ars:9-802                            | verified |
| edge:usa-az:003 | ahj:usa-az:county-code-adoption         | authorizes                     | county_building_related_codes_in_zoned_unincorporated_areas | src:usa-az:ars:11-861                           | verified |
| edge:usa-az:004 | ahj:usa-az:county-code-adoption         | limits_code_sources_to         | national_codes_or_largest_city_code                         | src:usa-az:ars:11-861                           | verified |
| edge:usa-az:005 | ahj:usa-az:state-fire-marshal           | adopts                         | arizona_state_fire_code                                     | src:usa-az:ars:37-1383, src:usa-az:aac:r4-36-201 | verified |
| edge:usa-az:006 | ahj:usa-az:public-building-code-rule    | applies                        | public_buildings                                            | src:usa-az:ars:34-461                           | verified |
| edge:usa-az:007 | ahj:usa-az:fire-district-code-authority | authorizes_with_voter_approval | fire_district_nationally_recognized_fire_code               | src:usa-az:ars:48-805                           | verified |
| edge:usa-az:008 | ahj:usa-az:elevator-conveyance-safety   | adopts                         | statewide_conveyance_safety_standards                       | src:usa-az:aac:r20-5-507                        | verified |

---

## 3. Current Statewide Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family                          | State Code Name                                       | Base Model Code                            | Edition                          | Status                       | Adoption Date             | Effective Date                                                                                    | Operative Date           | Mandatory Date                              | Transition Rule                                                        | Source IDs                                                                           |
| ------------------------------------ | ----------------------------------------------------- | ------------------------------------------ | -------------------------------- | ---------------------------- | ------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------ | ------------------------------------------- | ---------------------------------------------------------------------- | ------------------------------------------------------------------------------------ |
| Building                             | No single statewide private building code verified    | Local IBC or other local code, if adopted  | jurisdiction_specific            | local_option_verified        | null                      | local_ordinance_specific                                                                          | local_ordinance_specific | local_ordinance_specific                    | Public-building fallback uses largest city in county if no local codes | src:usa-az:ars:9-801, src:usa-az:ars:9-802, src:usa-az:ars:11-861, src:usa-az:ars:34-461 |
| Residential                          | No single statewide private residential code verified | Local IRC or other local code, if adopted  | jurisdiction_specific            | local_option_verified        | null                      | local_ordinance_specific                                                                          | local_ordinance_specific | local_ordinance_specific                    | Residential sprinkler state preemptions apply                          | src:usa-az:ars:9-807, src:usa-az:ars:9-808, src:usa-az:ars:11-861                       |
| Existing Building / Rehabilitation   | No statewide existing-building code verified          | Local IEBC or similar, if adopted          | jurisdiction_specific            | local_records_needed         | null                      | local_ordinance_specific                                                                          | local_ordinance_specific | local_ordinance_specific                    | Not verified statewide                                                 | src:usa-az:ars:9-801, src:usa-az:ars:9-802, src:usa-az:ars:11-861                       |
| Mechanical                           | No single statewide private mechanical code verified  | Local IMC/UMC or similar, if adopted       | jurisdiction_specific            | local_option_verified        | null                      | local_ordinance_specific                                                                          | local_ordinance_specific | local_ordinance_specific                    | Public-building fallback applies if no local codes                     | src:usa-az:ars:11-861, src:usa-az:ars:34-461                                           |
| Plumbing                             | No single statewide private plumbing code verified    | Local IPC/UPC or similar, if adopted       | jurisdiction_specific            | local_option_verified        | null                      | local_ordinance_specific                                                                          | local_ordinance_specific | local_ordinance_specific                    | Public-building fallback applies if no local codes                     | src:usa-az:ars:11-861, src:usa-az:ars:34-461                                           |
| Fuel Gas                             | No single statewide private fuel-gas code verified    | Local fuel-gas provisions, if adopted      | jurisdiction_specific            | local_records_needed         | null                      | local_ordinance_specific                                                                          | local_ordinance_specific | local_ordinance_specific                    | Not verified statewide                                                 | src:usa-az:ars:9-801, src:usa-az:ars:9-802, src:usa-az:ars:11-861                       |
| Electrical                           | No single statewide private electrical code verified  | Local electrical code, if adopted          | jurisdiction_specific            | local_option_verified        | null                      | local_ordinance_specific                                                                          | local_ordinance_specific | local_ordinance_specific                    | Public-building fallback applies if no local codes                     | src:usa-az:ars:9-801, src:usa-az:ars:11-861, src:usa-az:ars:34-461                      |
| Energy                               | No statewide energy code baseline verified            | Local energy conservation code, if adopted | jurisdiction_specific            | unresolved_local             | null                      | local_ordinance_specific                                                                          | local_ordinance_specific | local_ordinance_specific                    | Local records required                                                 | src:usa-az:ars:9-801, src:usa-az:ars:9-802                                             |
| Fire - construction references       | Arizona State Fire Code                               | International Fire Code                    | 2024                             | statewide_specialty_verified | rulemaking_history_in_aac | 2026-04-13 for 2024 IFC amendment                                                                 | 2026-04-13               | 2026-04-13 unless otherwise provided by law | No later amendments or editions included                               | src:usa-az:aac:r4-36-201                                                              |
| Fire - operational / prevention code | Arizona State Fire Code                               | International Fire Code                    | 2024                             | statewide_specialty_verified | rulemaking_history_in_aac | 2026-04-13 for 2024 IFC amendment                                                                 | 2026-04-13               | 2026-04-13 unless otherwise provided by law | Local exceptions and fire-district rules apply                         | src:usa-az:aac:r4-36-201, src:usa-az:ars:37-1383, src:usa-az:ars:48-805                 |
| Accessibility                        | Not verified as separate Arizona statewide code       | unknown                                    | unknown                          | unresolved                   | null                      | null                                                                                              | null                     | null                                        | Federal and local code analysis required                               | none                                                                                 |
| Elevator / Conveyance                | Arizona elevator and conveyance safety standards      | ASME A17.1 / ASME A17.7                    | ASME A17.1-2019; ASME A17.7-2007 | statewide_specialty_verified | rulemaking_history_in_aac | 2023-01-09 for current rule amendment; Jan. 1, 2023 trigger for installations/repairs/alterations | 2023-01-01 trigger       | 2023-01-01 trigger                          | Different compliance options for pre-2023 and pre-2009 conveyances     | src:usa-az:aac:r20-5-507                                                              |

### 3.2 Adoption Records

| Adoption Record ID                       | Code Family                                                     | Adopting Authority                                                            | Code Adopted                                                              | Edition                          | Scope                                                                        | Effective / Trigger Date                                                                                | Notes                                                             | Source IDs              | Status   |
| ---------------------------------------- | --------------------------------------------------------------- | ----------------------------------------------------------------------------- | ------------------------------------------------------------------------- | -------------------------------- | ---------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------- | -------- |
| adoption:usa-az:fire:2024-ifc             | Fire - operational / prevention; Fire - construction references | Office of the State Fire Marshal / Department of Forestry and Fire Management | Arizona State Fire Code                                                   | 2024 International Fire Code     | Statewide unless otherwise provided by law; modified by Article 3            | 2026-04-13 for current 2024 IFC amendment                                                               | Includes specified appendices; excludes later amendments/editions | src:usa-az:aac:r4-36-201 | verified |
| adoption:usa-az:elevator:asme-2019        | Elevator / Conveyance                                           | Industrial Commission of Arizona / ADOSH Elevator Safety rules                | Safety Code for Elevators and Escalators; Performance-Based Safety Code   | ASME A17.1-2019; ASME A17.7-2007 | Covered conveyances installed, repaired, or altered on or after Jan. 1, 2023 | 2023-01-01 trigger; rule amended effective 2023-01-09                                                   | Older conveyances have separate compliance options                | src:usa-az:aac:r20-5-507 | verified |
| adoption:usa-az:public-buildings:fallback | Public buildings                                                | Arizona Legislature                                                           | Local codes or fallback to State Fire Code plus largest-city county codes | jurisdiction_specific            | Public buildings                                                             | Codes in effect when designed or constructed; current codes for unsafe/fire hazard/dangerous conditions | Applies where local codes are absent                              | src:usa-az:ars:34-461    | verified |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

Arizona does not appear to use a single statewide construction-code adoption date for ordinary private buildings. Date logic is instead scope-specific:

1. **Municipal code adoption:** A municipality may adopt an existing code or public record by reference through an adopting ordinance published in full, with required copies filed and available for public use and inspection. Amendments may be adopted in the same manner.

2. **County code adoption:** A county board of supervisors may adopt building and related codes for zoned unincorporated areas after notice and hearings. If the county adopts the largest city's code, the county must adopt the same change within 90 days after receiving written notice of a city-code change or terminate the adopted city code.

3. **State Fire Code:** The Arizona State Fire Code is currently based on the 2024 International Fire Code as incorporated by R4-36-201, with no later amendments or editions included. The current amendment adopting the 2024 IFC is effective April 13, 2026.

4. **Public buildings:** Public buildings are subject to the applicable codes in effect when the building is designed or constructed. Current codes apply when a building is structurally unsafe, lacks adequate egress, is a fire hazard, or is otherwise dangerous to human life.

5. **Elevator / conveyance:** Covered conveyances installed, repaired, or altered on or after January 1, 2023 use ASME A17.1-2019 or ASME A17.7-2007 as referenced in ASME A17.1-2019. Separate rules apply for conveyances installed, repaired, or altered between May 5, 2009 and December 31, 2022, and before May 5, 2009.

### 4.2 Date Rule Table

| Rule ID                                             | Applies To                                  | Rule Type                              | Date / Period                                    | Trigger Condition                                                                  | Prior Code Allowed?                                             | Source IDs              | Status   |
| --------------------------------------------------- | ------------------------------------------- | -------------------------------------- | ------------------------------------------------ | ---------------------------------------------------------------------------------- | --------------------------------------------------------------- | ----------------------- | -------- |
| date-rule:usa-az:municipal:adoption-by-reference     | Municipal codes                             | adoption_procedure                     | local ordinance date                             | Municipal adopting ordinance published in full; copies filed for inspection        | local ordinance specific                                        | src:usa-az:ars:9-802     | verified |
| date-rule:usa-az:county:notice-hearing               | County codes                                | adoption_procedure                     | local county action                              | Notice and hearings before planning and zoning commission and board of supervisors | local county action specific                                    | src:usa-az:ars:11-861    | verified |
| date-rule:usa-az:county:largest-city-update          | County code based on largest city in county | update_rule                            | 90 days after written notice of city-code change | County has adopted largest city code                                               | County must adopt same change or terminate adopted city code    | src:usa-az:ars:11-861    | verified |
| date-rule:usa-az:fire:2024-ifc                       | State Fire Code                             | effective_date                         | 2026-04-13                                       | R4-36-201 amendment adopting 2024 IFC                                              | No later amendments or editions included                        | src:usa-az:aac:r4-36-201 | verified |
| date-rule:usa-az:public-building:design-construction | Public buildings                            | applicability_date                     | code in effect when designed or constructed      | Public building design/construction                                                | Current codes apply for unsafe/fire hazard/dangerous conditions | src:usa-az:ars:34-461    | verified |
| date-rule:usa-az:elevator:2023-trigger               | Elevator / conveyance                       | installation_repair_alteration_trigger | 2023-01-01                                       | Installed, repaired, or altered on/after Jan. 1, 2023                              | Older conveyances have separate compliance options              | src:usa-az:aac:r20-5-507 | verified |

### 4.3 Pending or Future Code Changes

| Code Family                                                          | Future Code                                 | Announced Date        | Adoption Date         | Effective Date        | Operative Date        | Mandatory Date        | Watch Status             | Source IDs              | Notes                                                         |
| -------------------------------------------------------------------- | ------------------------------------------- | --------------------- | --------------------- | --------------------- | --------------------- | --------------------- | ------------------------ | ----------------------- | ------------------------------------------------------------- |
| Fire                                                                 | none verified beyond current 2024 IFC rule  | null                  | null                  | null                  | null                  | null                  | monitor_aac_and_register | src:usa-az:aac:r4-36-201 | Current rule already reflects 2024 IFC in A.A.C. Supp. 26-1.  |
| Building / Residential / Mechanical / Plumbing / Electrical / Energy | local jurisdiction updates                  | jurisdiction_specific | jurisdiction_specific | jurisdiction_specific | jurisdiction_specific | jurisdiction_specific | monitor_local_ordinances | local_records_needed    | No statewide private construction-code cycle verified.        |
| Elevator / Conveyance                                                | none verified beyond current ASME standards | null                  | null                  | null                  | null                  | null                  | monitor_aac_and_register | src:usa-az:aac:r20-5-507 | Current rule reflects ASME A17.1-2019 / A17.7-2007 structure. |

### 4.4 Special Applicability Rules

| Rule ID                                                   | Code Family                                      | Project Type / Scope                                            | Trigger                                                        | Rule Summary                                                                                                                                                                                                      | Source IDs                                | Status   |
| --------------------------------------------------------- | ------------------------------------------------ | --------------------------------------------------------------- | -------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------- | -------- |
| applicability-rule:usa-az:public-buildings                 | Building, plumbing, electrical, fire, mechanical | Public buildings                                                | Public building design/construction or unsafe/hazard condition | Public buildings follow local adopted codes; if no local codes, use State Fire Code and the largest city in the county's building, plumbing, electrical, and mechanical codes.                                    | src:usa-az:ars:34-461                      | verified |
| applicability-rule:usa-az:state-fire-exceptions            | Fire                                             | State Fire Code enforcement                                     | City population and local code assumption; occupancy type      | State Fire Marshal enforces statewide except in qualifying cities with nationally recognized fire codes and local assumption of jurisdiction, but state/county owned buildings remain subject to State Fire Code. | src:usa-az:ars:37-1383                     | verified |
| applicability-rule:usa-az:county-unincorporated-zoned      | Building and related codes                       | Zoned unincorporated county areas                               | County has adopted zoning                                      | County may adopt and enforce building and related codes in zoned unincorporated areas; rural/unclassified areas may be exempted.                                                                                  | src:usa-az:ars:11-861                      | verified |
| applicability-rule:usa-az:county-fire-district             | Fire prevention                                  | Unincorporated county area without fire-district code           | Fire district has not adopted nationally recognized fire code  | County may adopt fire prevention code; county code remains until a fire district is established and adopts an applicable code.                                                                                    | src:usa-az:ars:11-861                      | verified |
| applicability-rule:usa-az:residential-sprinkler-preemption | Residential/fire                                 | One- and two-family residences and certain accessory structures | Municipal/county fire-code access-road or route requirements   | State preemption limits local rules that indirectly require sprinklers through fire apparatus access-road or route requirements.                                                                                  | src:usa-az:ars:9-808, src:usa-az:ars:11-861 | verified |
| applicability-rule:usa-az:elevator-date-bands              | Elevator / conveyance                            | Conveyances                                                     | Installation, repair, or alteration date                       | Jan. 1, 2023 and older date bands determine available ASME compliance pathways.                                                                                                                                   | src:usa-az:aac:r20-5-507                   | verified |

---

## 5. State Amendments

### 5.1 Amendment Model

State amendment structure:

* **Fire:** Arizona adopts the 2024 International Fire Code by reference as the State Fire Code and modifies it through Article 3 of A.A.C. Title 4, Chapter 36, including definitions, appendices, permits, inspections/enforcement, and other provisions.
* **Elevator / conveyance:** Arizona adopts ASME standards by rule with Arizona-specific applicability and date-band rules.
* **Ordinary building/residential/trade codes:** No statewide amendment set was verified because no single statewide private construction-code adoption was verified.

Where amendments are published:

* State Fire Code modifications: Arizona Administrative Code, Title 4, Chapter 36, Article 3.
* Elevator/conveyance rules: Arizona Administrative Code, Title 20, Chapter 5, Article 5.
* Local building-code amendments: local ordinances and adopted-code records.

**Amendment parsing status:** partial_state_specialty_only.

### 5.2 State Amendment Sources

| Amendment Source ID               | Code Family           | Publisher                                                               | Location                              | Status                      |
| --------------------------------- | --------------------- | ----------------------------------------------------------------------- | ------------------------------------- | --------------------------- |
| amend-src:usa-az:fire:article3     | Fire                  | Arizona Secretary of State / Department of Forestry and Fire Management | A.A.C. Title 4, Chapter 36, Article 3 | identified_not_fully_parsed |
| amend-src:usa-az:elevator:article5 | Elevator / Conveyance | Arizona Secretary of State / Industrial Commission of Arizona           | A.A.C. Title 20, Chapter 5, Article 5 | partially_parsed            |

### 5.3 High-Impact State Amendments

| Amendment ID                                  | Code Family           | Topic                                  | Summary                                                                                                                                                                                                                        | Source IDs                                | Status   |
| --------------------------------------------- | --------------------- | -------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ----------------------------------------- | -------- |
| amendment:usa-az:fire:ifc-2024-appendices      | Fire                  | IFC appendices                         | R4-36-201 includes specified Appendix D sections and Appendices B, C, E, F, G, H, I, J, and N.                                                                                                                                 | src:usa-az:aac:r4-36-201                   | verified |
| amendment:usa-az:fire:no-later-editions        | Fire                  | Incorporation by reference             | State Fire Code incorporation does not include later IFC amendments or editions.                                                                                                                                               | src:usa-az:aac:r4-36-201                   | verified |
| amendment:usa-az:elevator:no-later-editions    | Elevator / Conveyance | Incorporation by reference             | R20-5-507 incorporations do not include later amendments or editions.                                                                                                                                                          | src:usa-az:aac:r20-5-507                   | verified |
| amendment:usa-az:elevator:date-band-compliance | Elevator / Conveyance | Transition / older installations       | Older conveyances have compliance options based on installation/repair/alteration date bands.                                                                                                                                  | src:usa-az:aac:r20-5-507                   | verified |
| amendment:usa-az:local:sprinkler-preemption    | Residential / Fire    | Residential sprinklers and fire access | State law limits local fire-code, ordinance, stipulation, or legal requirements for fire apparatus access roads or routes that directly or indirectly require sprinklers in specified one- or two-family residential contexts. | src:usa-az:ars:9-808, src:usa-az:ars:11-861 | verified |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
enforcement_id: "local-enforcement:usa-az"
model: "local_adoption_with_state_fire_and_specialty_overlays"
enforcing_entities:
- "city"
- "town"
- "county_for_zoned_unincorporated_areas"
- "fire_district"
- "office_of_state_fire_marshal"
- "industrial_commission_elevator_safety"
required_officials:
- "local_building_official_or_equivalent"
- "county_code_enforcement_official"
- "state_fire_marshal_or_designee"
- "fire_district_authority_where_applicable"
- "elevator_safety_inspector"
state_reserved_activities:
- "state_fire_code_adoption"
- "state_fire_code_enforcement_except_statutory_local_assumption_or_exceptions"
- "state_county_and_school_public_building_fire_review_scope"
- "elevator_and_conveyance_safety"
source_ids:
- "src:usa-az:ars:9-801"
- "src:usa-az:ars:9-802"
- "src:usa-az:ars:11-861"
- "src:usa-az:ars:34-461"
- "src:usa-az:ars:37-1381"
- "src:usa-az:ars:37-1383"
- "src:usa-az:aac:r4-36-201"
- "src:usa-az:ars:48-805"
- "src:usa-az:aac:r20-5-507"
verification_status: "verified_state_level"
confidence: 0.86
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
rule_id: "local-amendment-rule:usa-az"
model: "local_ordinance_or_code_by_reference_with_state_limits"
applies_to_code_families:
- "building"
- "residential"
- "existing_building"
- "mechanical"
- "plumbing"
- "fuel_gas"
- "electrical"
- "energy"
- "fire_operational"
- "wildland_urban_interface"
approval_required:
municipality: "local_process_not_state_approval_verified"
county: "county_notice_and_hearing_required; code-source limits apply"
fire_district: "state_fire_marshal_approval_required_for_adoption_code_source_and_for_amendments_or_revisions"
approving_authority_id:
municipality: "local_governing_body"
county: "county_board_of_supervisors"
fire_district: "fire_district_board_voters_and_state_fire_marshal"
filing_required:
municipality: "copies filed with municipal clerk"
county: "county records/local process"
fire_district: "three copies on file for public inspection"
registry_exists: "not_verified"
registry_source_ids: []
legal_basis_source_ids:
- "src:usa-az:ars:9-802"
- "src:usa-az:ars:11-861"
- "src:usa-az:ars:48-805"
verification_status: "verified_state_level"
confidence: 0.82
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Arizona's code model should not be interpreted as "no state code." Rather, it is a layered model:

1. Ordinary private construction code editions are usually local.
2. Counties have specific authority for zoned unincorporated areas.
3. The State Fire Code is statewide unless otherwise provided by law and is subject to statutory local exceptions.
4. Public buildings have a statutory fallback rule.
5. Elevator/conveyance safety is a statewide specialty code system.

### 6.4 Known Local Amendment Registries

No statewide local building-code or amendment registry was verified in this pass.

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendments were parsed in this pass. Required next records include Phoenix, Tucson, Mesa, Chandler, Scottsdale, Gilbert, Glendale, Tempe, Peoria, Surprise, and county unincorporated codes.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

Resolver status: `state_level_model_ready`; `address_level_not_started`.

Jurisdiction stack:

```text
Address
-> State: Arizona
-> County
-> Municipality, if incorporated
-> If unincorporated: county zoning status and county adopted-code status
-> Fire district or municipal/county fire authority
-> Public-building status, if applicable
-> State Fire Marshal jurisdiction or local fire-code assumption / exception
-> Elevator/conveyance jurisdiction, if conveyance involved
-> Applicable local building/residential/trade code adoptions
-> Applicable local amendments and preemptions
-> AHJ contacts and permit offices
```

### 7.2 Boundary Data Sources

| Boundary Type | Source | Source ID | Coverage | Update Frequency | Status |
| --- | --- | --- | --- | --- | --- | --- |
| State | Arizona official statutes and administrative code | src:usa-az:ars:portal, src:usa-az:aac:portal | statewide | when law/rules change | verified |
| County | not selected | none | statewide | unknown | pending |
| Municipality | not selected | none | statewide | unknown | pending |
| Fire District | not selected | none | statewide | unknown | pending |
| Special District | not selected | none | statewide | unknown | pending |

### 7.3 AHJ Contact Data

No AHJ contact data was populated for this pass.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID                        | Title                                                                                                                                                                                    | Source Type         | Publisher                                                               | Accessed Date | Snapshot ID      | Checksum         | Status                 |
| -------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------- | ----------------------------------------------------------------------- | ------------- | ---------------- | ---------------- | ---------------------- |
| src:usa-az:ars:9-801              | A.R.S. § 9-801, Definitions                                                                                                                                                              | statute             | Arizona Legislature                                                     | 2026-06-25    | snapshot-pending | snapshot-pending | verified               |
| src:usa-az:ars:9-802              | A.R.S. § 9-802, Procedure for adoption by reference                                                                                                                                      | statute             | Arizona Legislature                                                     | 2026-06-25    | snapshot-pending | snapshot-pending | verified               |
| src:usa-az:ars:9-807              | A.R.S. § 9-807, Mandated fire sprinklers in certain residences                                                                                                                           | statute             | Arizona Legislature                                                     | 2026-06-25    | snapshot-pending | snapshot-pending | verified               |
| src:usa-az:ars:9-808              | A.R.S. § 9-808, Fire apparatus access road or approved route; fire watch requirements                                                                                                    | statute             | Arizona Legislature                                                     | 2026-06-25    | snapshot-pending | snapshot-pending | verified               |
| src:usa-az:ars:11-861             | A.R.S. § 11-861, Adoption of codes by reference; limitations; method of adoption                                                                                                         | statute             | Arizona Legislature                                                     | 2026-06-25    | snapshot-pending | snapshot-pending | verified               |
| src:usa-az:ars:11-862             | A.R.S. § 11-862, Advisory board; appointment; terms; duties                                                                                                                              | statute             | Arizona Legislature                                                     | 2026-06-25    | snapshot-pending | snapshot-pending | verified               |
| src:usa-az:ars:34-461             | A.R.S. § 34-461, Applicability of local codes; exceptions; definition                                                                                                                    | statute             | Arizona Legislature                                                     | 2026-06-25    | snapshot-pending | snapshot-pending | verified               |
| src:usa-az:ars:37-1381            | A.R.S. § 37-1381, Office of the State Fire Marshal                                                                                                                                       | statute             | Arizona Legislature                                                     | 2026-06-25    | snapshot-pending | snapshot-pending | verified               |
| src:usa-az:ars:37-1383            | A.R.S. § 37-1383, State Fire Marshal powers and duties                                                                                                                                   | statute             | Arizona Legislature                                                     | 2026-06-25    | snapshot-pending | snapshot-pending | verified               |
| src:usa-az:aac:r4-36-201          | A.A.C. R4-36-201, Incorporation by Reference of the International Fire Code                                                                                                              | administrative_rule | Arizona Secretary of State / Department of Forestry and Fire Management | 2026-06-25    | snapshot-pending | snapshot-pending | verified               |
| src:usa-az:ars:48-805             | A.R.S. § 48-805, Fire district powers and duties                                                                                                                                         | statute             | Arizona Legislature                                                     | 2026-06-25    | snapshot-pending | snapshot-pending | verified               |
| src:usa-az:aac:r20-5-507          | A.A.C. R20-5-507, Safety Code for Elevators, Escalators, Dumbwaiters, Moving Walks, Material Lifts, Special Purpose Personnel Elevators, and Dumbwaiters with Automatic Transfer Devices | administrative_rule | Arizona Secretary of State / Industrial Commission of Arizona           | 2026-06-25    | snapshot-pending | snapshot-pending | verified               |
| src:usa-az:constitution:article13 | Arizona Constitution, Article 13, Municipal Corporations                                                                                                                                 | constitution        | Arizona Legislature                                                     | 2026-06-24    | snapshot-pending | snapshot-pending | retained_from_baseline |

### 8.2 Official-Source Caveats

| Source ID                        | Caveat Type               | Caveat Summary                                                                                                           | Production Treatment         |
| -------------------------------- | ------------------------- | ------------------------------------------------------------------------------------------------------------------------ | ---------------------------- |
| src:usa-az:ars:9-801              | municipal_scope           | Defines "code" broadly for municipal adoption, but local ordinance records are still needed to identify actual editions. | use_for_authority_only       |
| src:usa-az:ars:9-802              | adoption_procedure_only   | Establishes municipal adoption-by-reference method, not specific model-code editions.                                    | use_for_process_only         |
| src:usa-az:ars:11-861             | county_scope_limited      | Applies to counties with zoning and zoned unincorporated areas; local county adoption records still required.            | use_for_authority_and_limits |
| src:usa-az:ars:34-461             | public_buildings_only     | Applies to public buildings and contains fallback logic; not a statewide private-building code adoption.                 | scope_limited                |
| src:usa-az:aac:r4-36-201          | fire_only                 | Verifies State Fire Code, not general building/residential/trade codes.                                                  | use_for_fire_only            |
| src:usa-az:aac:r20-5-507          | conveyance_only           | Verifies elevator/conveyance standards, not building-code editions.                                                      | use_for_conveyances_only     |
| src:usa-az:constitution:article13 | constitutional_level_only | Useful for municipal-corporation context, but superseded for code-specific analysis by extracted statutes and rules.     | background_only              |

### 8.3 Supplemental Sources

None used as controlling authority. Only official Arizona Legislature and Arizona Administrative Code sources were used for the updated conclusions.

### 8.4 Source Extraction Metadata

| Source ID               | Parser                  | Parser Version | Extracted At         | Extraction Confidence | Requires OCR | Requires Browser Automation | Manual Review Required |
| ----------------------- | ----------------------- | -------------- | -------------------- | --------------------- | ------------ | --------------------------- | ---------------------- |
| src:usa-az:ars:9-801     | browser_manual          | 0.2            | 2026-06-25T00:00:00Z | 0.98                  | no           | yes                         | yes                    |
| src:usa-az:ars:9-802     | browser_manual          | 0.2            | 2026-06-25T00:00:00Z | 0.98                  | no           | yes                         | yes                    |
| src:usa-az:ars:11-861    | browser_manual          | 0.2            | 2026-06-25T00:00:00Z | 0.98                  | no           | yes                         | yes                    |
| src:usa-az:ars:34-461    | browser_manual          | 0.2            | 2026-06-25T00:00:00Z | 0.98                  | no           | yes                         | yes                    |
| src:usa-az:aac:r4-36-201 | pdf_text_browser_manual | 0.2            | 2026-06-25T00:00:00Z | 0.96                  | no           | yes                         | yes                    |
| src:usa-az:aac:r20-5-507 | pdf_text_browser_manual | 0.2            | 2026-06-25T00:00:00Z | 0.94                  | no           | yes                         | yes                    |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID                                           | Field                   | Value                                                                                                                     | Verification Status  | Confidence | Source IDs                                                                                | Notes                                                                            |
| --------------------------------------------------- | ----------------------- | ------------------------------------------------------------------------------------------------------------------------- | -------------------- | ---------- | ----------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| report:usa-az                                        | report.status           | partially_verified                                                                                                        | verified             | 1.00       | all_verified_sources                                                                      | State-level authority model is verified; local adoption records remain unparsed. |
| report:usa-az                                        | risk.overall_confidence | 0.78                                                                                                                      | partially_verified   | 0.78       | all_verified_sources                                                                      | Confidence improved from baseline due to extraction of statutes/rules.           |
| adoption:usa-az:fire:2024-ifc                        | edition                 | 2024 International Fire Code                                                                                              | verified             | 0.95       | src:usa-az:aac:r4-36-201                                                                   | State Fire Code verified.                                                        |
| adoption:usa-az:elevator:asme-2019                   | edition                 | ASME A17.1-2019 / ASME A17.7-2007                                                                                         | verified             | 0.92       | src:usa-az:aac:r20-5-507                                                                   | Conveyance code verified.                                                        |
| local-enforcement:usa-az                             | model                   | local_adoption_with_state_fire_and_specialty_overlays                                                                     | verified_state_level | 0.86       | src:usa-az:ars:9-802, src:usa-az:ars:11-861, src:usa-az:ars:37-1383, src:usa-az:aac:r4-36-201 | Local editions still require jurisdiction extraction.                            |
| local-amendment-rule:usa-az                          | model                   | local_ordinance_or_code_by_reference_with_state_limits                                                                    | verified_state_level | 0.82       | src:usa-az:ars:9-802, src:usa-az:ars:11-861, src:usa-az:ars:48-805                           | State-level process and limits verified.                                         |
| date-rule:usa-az:public-building:design-construction | rule                    | public buildings use codes in effect when designed or constructed; current codes for unsafe/fire hazard/danger conditions | verified             | 0.95       | src:usa-az:ars:34-461                                                                      | Important for permit/date logic.                                                 |

### 9.2 Validation Checks

| Check                                                  | Status           | Notes                                                                                                                                                 |
| ------------------------------------------------------ | ---------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
| All source IDs resolve                                 | pass             | Source IDs listed in section 8.                                                                                                                       |
| All authority IDs resolve                              | pass             | Authority IDs defined in section 2.                                                                                                                   |
| Statewide ordinary building code verified or disproven | pass_with_caveat | No single statewide private IBC/IRC-style adoption verified; state/local model clarified.                                                             |
| Building and operational fire code are separated       | pass             | Fire is verified separately as State Fire Code; ordinary building remains local.                                                                      |
| Public-building fallback captured                      | pass             | A.R.S. § 34-461 added.                                                                                                                                |
| County adoption model captured                         | pass             | A.R.S. §§ 11-861 and 11-862 added.                                                                                                                    |
| Municipal adoption-by-reference captured               | pass             | A.R.S. §§ 9-801 and 9-802 added.                                                                                                                      |
| Fire district authority captured                       | pass             | A.R.S. § 48-805 added.                                                                                                                                |
| Elevator/conveyance specialty code captured            | pass             | A.A.C. R20-5-507 added.                                                                                                                               |
| Effective dates are valid ISO dates                    | pass             | Known dates normalized.                                                                                                                               |
| No impossible date sequences                           | pass             | No conflicting date sequences found.                                                                                                                  |
| Transition rules have explicit trigger conditions      | pass_partial     | State Fire Code, public buildings, county city-code update, and elevator date bands captured; local ordinance date rules still jurisdiction-specific. |
| Permit-date logic captured where applicable            | pass_partial     | Public building and elevator triggers captured; local permit vesting not exhaustively parsed.                                                         |
| Local enforcement model classified                     | pass             | Classified as local adoption with state fire/specialty overlays.                                                                                      |
| Local amendment rule classified                        | pass             | Classified with municipal, county, and fire-district distinctions.                                                                                    |
| AHJ confirmation metadata present                      | fail             | No AHJ contact database populated.                                                                                                                    |
| Official-source caveats captured                       | pass             | Scope limitations are recorded.                                                                                                                       |

---

## 10. Open Issues and Review Queue

| Issue ID        | Severity | Topic                         | Description                                                                                                                            | Needed Action                                                                          | Owner | Due Date | Status |
| --------------- | -------- | ----------------------------- | -------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | ----- | -------- | ------ |
| issue:usa-az:001 | high     | local adoption records        | Identify current building/residential/mechanical/plumbing/electrical/energy/fire code editions for major municipalities and counties.  | Parse local ordinances and adopted-code pages for top municipalities and all counties. | null  | null     | open   |
| issue:usa-az:002 | high     | AHJ resolution                | Build address-level jurisdiction resolver for incorporated vs. unincorporated areas, fire districts, and public-building overlays.     | Select boundary datasets and AHJ source list.                                          | null  | null     | open   |
| issue:usa-az:003 | medium   | State Fire Code modifications | Parse A.A.C. Title 4, Chapter 36, Article 3 into high-impact amendment records.                                                        | Extract R4-36-301 through applicable Article 3 provisions and Exhibit A.               | null  | null     | open   |
| issue:usa-az:004 | medium   | energy and accessibility      | Confirm whether any Arizona-specific statewide energy or accessibility requirements apply outside local code adoption and federal law. | Search state statutes, rules, and agency guidance; then local adoptions.               | null  | null     | open   |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID                          | Source ID                                    | Watch Type   | Frequency | Trigger Condition                              | Last Checked | Status |
| --------------------------------- | -------------------------------------------- | ------------ | --------- | ---------------------------------------------- | ------------ | ------ |
| watch:usa-az:ars:9-801-802         | src:usa-az:ars:9-801; src:usa-az:ars:9-802     | statute_diff | quarterly | municipal adoption-by-reference changes        | 2026-06-25   | active |
| watch:usa-az:ars:11-861-862        | src:usa-az:ars:11-861; src:usa-az:ars:11-862   | statute_diff | quarterly | county code adoption or advisory-board changes | 2026-06-25   | active |
| watch:usa-az:ars:34-461            | src:usa-az:ars:34-461                         | statute_diff | quarterly | public-building applicability changes          | 2026-06-25   | active |
| watch:usa-az:state-fire-code       | src:usa-az:aac:r4-36-201                      | rule_diff    | monthly   | State Fire Code edition or amendment change    | 2026-06-25   | active |
| watch:usa-az:fire-marshal-statutes | src:usa-az:ars:37-1381; src:usa-az:ars:37-1383 | statute_diff | quarterly | State Fire Marshal authority changes           | 2026-06-25   | active |
| watch:usa-az:fire-districts        | src:usa-az:ars:48-805                         | statute_diff | quarterly | fire district code-adoption authority changes  | 2026-06-25   | active |
| watch:usa-az:elevator-conveyance   | src:usa-az:aac:r20-5-507                      | rule_diff    | quarterly | ASME conveyance-standard updates               | 2026-06-25   | active |

### 11.2 Changelog

| Date       | Change                                                    | Record IDs                                                                                                                         | Source IDs                                                                                                                                                                                                                       | Changed By       | Notes                                                                                                                                  |
| ---------- | --------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------- | -------------------------------------------------------------------------------------------------------------------------------------- |
| 2026-06-24 | Baseline Arizona municipal-corporation report created     | report:usa-az                                                                                                                       | src:usa-az:constitution:article13, src:usa-az:statutes:portal                                                                                                                                                                      | Codex            | Statewide code status unresolved.                                                                                                      |
| 2026-06-25 | Filled state-level statutory and administrative-code gaps | report:usa-az; adoption:usa-az:fire:2024-ifc; adoption:usa-az:elevator:asme-2019; local-enforcement:usa-az; local-amendment-rule:usa-az | src:usa-az:ars:9-801, src:usa-az:ars:9-802, src:usa-az:ars:11-861, src:usa-az:ars:11-862, src:usa-az:ars:34-461, src:usa-az:ars:37-1381, src:usa-az:ars:37-1383, src:usa-az:aac:r4-36-201, src:usa-az:ars:48-805, src:usa-az:aac:r20-5-507 | GPT-5.5 Thinking | Confirmed local-adoption model, State Fire Code, public-building fallback, fire-district authority, and elevator/conveyance standards. |
