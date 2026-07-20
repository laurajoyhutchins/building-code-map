---
state:
  state_id: "US-AL"
  name: "Alabama"
  abbreviation: "AL"
report:
  report_id: "state-report:usa-al"
  schema_version: "0.2.1"
  status: "partially_verified"
  last_updated: "2026-06-25"
  last_verified: "2026-06-25"
  reviewed_by: null
risk:
  overall_confidence: 0.74
  risk_flags:
    - "state_building_code_scope_is_limited_not_universal"
    - "residential_code_edition_requires_final_board_publication_review"
    - "commercial_energy_code_transition_requires_general_contractors_board_followup"
    - "fire_marshal_and_dcm_fire_scope_require_project_level_ahj_reconciliation"
    - "local_ahj_registry_not_yet_built"
  open_questions_count: 6
---

# State Building Code Authority Report: Alabama

## 1. Executive Summary

**Authority model:** Alabama uses a hybrid authority model. The Alabama Department of Finance, Real Property Management, Division of Construction Management (“DCM”) administers the Alabama State Building Code for a defined state/public/special-use scope, not for all private construction statewide. The Alabama State Fire Marshal has separate fire and life-safety authority. The Alabama Home Builders Licensure Board has residential building and residential energy code authority. Commercial energy code authority has transitioned away from ADECA and is now tied to the Alabama Licensing Board for General Contractors.

**Statewide code status:** Partially verified. Alabama has verified state-level codes for DCM-covered projects and State Fire Marshal fire-code scope. Residential and commercial energy authority are verified, but final current edition/adoption records still require additional board-level source extraction.

**Local enforcement model:** Partially verified. Residential construction enforcement depends heavily on local jurisdictions with permitting and inspection programs. County commissions may adopt building laws/codes for unincorporated areas. DCM jurisdiction is limited and does not generally cover ordinary residential or commercial projects, except specified categories.

**Local amendment posture:** Partially verified. Local residential building codes are allowed, but post-2027 local residential codes must meet state minimum standards. Local residential energy provisions may not exceed the Alabama Residential Energy Code except for local conditions or federal mandates. A complete local amendment registry was not verified.

**Known transition periods or pending changes:**

* DCM amendments incorporating portions of 2024 model codes became effective March 17, 2025.
* State Fire Marshal Regulation 482-2-101 adopted the 2021 International Fire Code effective January 1, 2025.
* Residential-builder compliance with the applicable residential building code begins January 1, 2027.
* County building-code resolutions generally become effective 120 days after adoption, with a limited insurance-claim-work exception.
* Commercial energy code transition requires follow-up with the Alabama Licensing Board for General Contractors.

**Production readiness:** `partially_ready_for_dcm_and_state_fire_marshal_scope`; `not_ready_for_universal_ahj_routing`

### Key Findings

```yaml
---
key_findings:
- topic: State building-code authority, limited scope
  finding: DCM is the verified state-level authority for the Alabama State Building
    Code within defined public/special-use project scope.
  confidence: 0.9
  source_ids:
  - src:usa-al:dcm:state-building-code
  - src:usa-al:admin-code:355-12-1
- topic: General private statewide building code
  finding: No universal statewide building-code enforcement model for all ordinary
    private commercial/residential projects was verified.
  confidence: 0.78
  source_ids:
  - src:usa-al:dcm:state-building-code
  - src:usa-al:hblb:law-2026
- topic: Primary DCM building code edition
  finding: 2021 IBC, amended with specified 2024 provisions, applies in DCM scope.
  confidence: 0.9
  source_ids:
  - src:usa-al:dcm:state-building-code
  - src:usa-al:admin-code:355-12-1
- topic: DCM companion codes
  finding: DCM scope includes 2021 IEBC, IPC, IFGC, IMC, IFC; 2020 NEC; ASHRAE 90.1-2013;
    2010 ADA Standards; 2020 ICC/NSSA-500; 2019 NFPA 72.
  confidence: 0.88
  source_ids:
  - src:usa-al:dcm:state-building-code
  - src:usa-al:admin-code:355-12-1
- topic: State Fire Marshal authority
  finding: Alabama State Fire Marshal has state-level fire/life-safety authority and
    has adopted the 2021 IFC with exceptions.
  confidence: 0.86
  source_ids:
  - src:usa-al:sfmo:about
  - src:usa-al:sfmo:482-2-101
- topic: Residential code authority
  finding: Alabama Home Builders Licensure Board has sole authority to adopt the Alabama
    Residential Building Code and Residential Energy Code.
  confidence: 0.86
  source_ids:
  - src:usa-al:hblb:law-2026
- topic: Residential enforcement
  finding: Residential code enforcement is local where local permitting/inspection
    programs exist; in jurisdictions without a local code, the board-adopted residential
    code applies.
  confidence: 0.82
  source_ids:
  - src:usa-al:hblb:law-2026
- topic: Commercial energy authority
  finding: Commercial energy authority transferred to the Alabama Licensing Board
    for General Contractors, but final adopted current code edition was not verified
    in this pass.
  confidence: 0.58
  source_ids:
  - src:usa-al:dcm:state-building-code
- topic: Local amendments / local codes
  finding: Local residential codes are allowed subject to state minimums; local residential
    energy provisions are constrained by state energy-code limits.
  confidence: 0.78
  source_ids:
  - src:usa-al:hblb:law-2026
- topic: Elevator / conveyance code
  finding: Not verified in this pass.
  confidence: 0.1
  source_ids:
  - unresolved
```

---

### 2.1 Primary Building Code Authorities

| Authority ID         | Authority Name                                                                               | Authority Type                                                     | Verified Role                                                                                       | Legal / Regulatory Basis                                                        | Enforcement Model                                                  | Status             |
| -------------------- | -------------------------------------------------------------------------------------------- | ------------------------------------------------------------------ | --------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------ | ------------------ |
| ahj:usa-al:dcm       | Alabama Department of Finance, Real Property Management, Division of Construction Management | state agency division                                              | Adopts/administers Alabama State Building Code for defined project scope                            | Ala. Admin. Code ch. 355-12-1; DCM State Building Code page                     | state review/enforcement within limited statutory/regulatory scope | verified           |
| ahj:usa-al:sfmo      | Alabama State Fire Marshal's Office                                                          | state fire authority / division of Alabama Department of Insurance | Adopts/enforces fire, life-safety, and fire-prevention regulations; adopts 2021 IFC with exceptions | Ala. Code Title 36, ch. 19; Ala. Admin. Code 482-2-101                          | state fire/life-safety authority                                   | partially_verified |
| ahj:usa-al:hblb      | Alabama Home Builders Licensure Board                                                        | state licensing/code authority                                     | Sole authority for Alabama Residential Building Code and Residential Energy Code                    | Ala. Code §§ 34-14A-12, 34-14A-80 through 34-14A-83; Ala. Admin. Code ch. 465-X | board adopts; local jurisdictions generally enforce                | verified           |
| ahj:usa-al:hblb-rcac | Residential Building Code Advisory Council                                                   | advisory council                                                   | Develops/proposes residential building and energy code recommendations to HBLB                      | Ala. Code § 34-14A-82                                                           | advisory only; board adopts or rejects                             | verified           |
| ahj:usa-al:genconbd  | Alabama Licensing Board for General Contractors / Commercial Energy Code function            | state board                                                        | Commercial energy code authority after transfer from ADECA                                          | Requires direct extraction from current General Contractors Board materials     | pending / transition                                               | partially_verified |
| ahj:usa-al:local     | Counties and municipalities                                                                  | local governments                                                  | May adopt/enforce local building codes within statutory limits                                      | HBLB law and local enabling authorities                                         | local permitting / inspection / enforcement                        | partially_verified |

### 2.2 Specialized Code Authorities

| Code Family                                | Authority ID                    | Authority Name                                  | Role                                                                                                      | Status                                 |
| ------------------------------------------ | ------------------------------- | ----------------------------------------------- | --------------------------------------------------------------------------------------------------------- | -------------------------------------- |
| Building, DCM-covered scope                | ahj:usa-al:dcm                  | DCM                                             | Adopts/administers Alabama State Building Code for defined covered projects                               | verified                               |
| Residential                                | ahj:usa-al:hblb                 | Alabama Home Builders Licensure Board           | Sole state residential building-code authority                                                            | verified_authority; edition_pending    |
| Existing Building / Rehabilitation         | ahj:usa-al:dcm; ahj:usa-al:sfmo | DCM; State Fire Marshal                         | DCM uses 2021 IEBC in DCM scope; Fire Marshal references require reconciliation                           | partially_verified                     |
| Mechanical                                 | ahj:usa-al:dcm; ahj:usa-al:sfmo | DCM; State Fire Marshal                         | DCM uses 2021 IMC with 2024 amendments; Fire Marshal legacy mechanical regulation requires reconciliation | partially_verified                     |
| Plumbing                                   | ahj:usa-al:dcm                  | DCM                                             | DCM uses 2021 IPC in DCM scope                                                                            | verified_for_dcm_scope                 |
| Fuel Gas                                   | ahj:usa-al:dcm; ahj:usa-al:sfmo | DCM; State Fire Marshal                         | DCM uses 2021 IFGC in DCM scope; Fire Marshal legacy IFGC rule requires reconciliation                    | partially_verified                     |
| Electrical                                 | ahj:usa-al:dcm                  | DCM                                             | DCM uses 2020 NEC in DCM scope                                                                            | verified_for_dcm_scope                 |
| Energy, DCM-covered state-funded buildings | ahj:usa-al:dcm                  | DCM                                             | ASHRAE 90.1-2013 applies to state-funded / state-instrumentality scope                                    | verified_for_dcm_scope                 |
| Residential Energy                         | ahj:usa-al:hblb                 | Alabama Home Builders Licensure Board           | Sole residential energy-code authority                                                                    | verified_authority; edition_pending    |
| Commercial Energy                          | ahj:usa-al:genconbd             | Alabama Licensing Board for General Contractors | Current commercial energy-code adoption authority after transfer                                          | partially_verified; final_code_pending |
| Fire - construction references             | ahj:usa-al:dcm                  | DCM                                             | DCM uses 2021 IFC with amendments for DCM scope                                                           | verified_for_dcm_scope                 |
| Fire - operational / prevention code       | ahj:usa-al:sfmo                 | State Fire Marshal                              | 2021 IFC adopted with exceptions and appendices                                                           | verified                               |
| Accessibility                              | ahj:usa-al:dcm                  | DCM                                             | 2010 ADA Standards; DCM accessibility review for covered public buildings/facilities                      | verified_for_dcm_scope                 |
| Elevator / Conveyance                      | ahj:usa-al:unknown              | Unknown                                         | Not parsed                                                                                                | unresolved                             |

### 2.3 Authority Hierarchy Notes

Alabama should not be modeled as a single statewide building-code authority covering every building type. The correct model is scope-dependent:

1. Determine whether the project falls within DCM-covered scope.
2. Determine whether State Fire Marshal fire/life-safety jurisdiction applies.
3. Determine whether the project is residential and therefore within HBLB residential-code framework.
4. Determine whether a municipality or county has adopted and enforces a local building code.
5. Determine whether commercial energy requirements are governed by the post-transfer General Contractors Board process.
6. Confirm the AHJ directly for the project address and occupancy.

### 2.4 Authority Graph Edges

| Edge ID                 | From Authority      | Relationship             | To Authority / Scope                                                                                 | Status             |
| ----------------------- | ------------------- | ------------------------ | ---------------------------------------------------------------------------------------------------- | ------------------ |
| edge:usa-al:001          | ahj:usa-al:dcm      | adopts_and_administers   | Alabama State Building Code for covered state/public/special-use scope                               | verified           |
| edge:usa-al:002          | ahj:usa-al:dcm      | applies_to               | state buildings/construction, schoolhouses, hotels, movie theaters, and defined accessibility review | verified           |
| edge:usa-al:003          | ahj:usa-al:dcm      | excludes_general_scope   | ordinary residential and ordinary commercial projects except hotels/motels/movie theaters            | verified           |
| edge:usa-al:004          | ahj:usa-al:sfmo     | adopts                   | State Fire Marshal fire code / 2021 IFC with exceptions                                              | verified           |
| edge:usa-al:005          | ahj:usa-al:hblb     | sole_authority_for       | Alabama Residential Building Code and Alabama Residential Energy Code                                | verified           |
| edge:usa-al:hblb:001     | ahj:usa-al:hblb     | local_enforcement_by     | local jurisdictions with permitting and inspection programs                                          | verified           |
| edge:usa-al:local:001    | ahj:usa-al:local    | may_adopt                | local residential building codes subject to state minimums                                           | partially_verified |
| edge:usa-al:genconbd:001 | ahj:usa-al:genconbd | transition_authority_for | commercial energy code                                                                               | partially_verified |

---

## 3. Current Code Adoptions

### 3.1 Code Adoption Matrix

| Code Family                     | State Code Name / Rule                              | Scope                                                       | Base Model Code               | Edition                                 | Status                              | Adoption / Effective Information                                                           | Transition Rule                                                                                  | Source IDs                                                       |
| ------------------------------- | --------------------------------------------------- | ----------------------------------------------------------- | ----------------------------- | --------------------------------------- | ----------------------------------- | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ---------------------------------------------------------------- |
| Building                        | Alabama State Building Code                         | DCM-covered projects                                        | IBC                           | 2021 IBC with specified 2024 amendments | verified_for_dcm_scope              | 2021 IBC adopted July 1, 2022; amended provisions effective March 17, 2025                 | Use DCM amended rule for covered projects                                                        | src:usa-al:dcm:state-building-code; src:usa-al:admin-code:355-12-1 |
| Existing Building               | Alabama State Building Code companion code          | DCM-covered projects                                        | IEBC                          | 2021                                    | verified_for_dcm_scope              | July 1, 2022                                                                               | none parsed                                                                                      | src:usa-al:dcm:state-building-code; src:usa-al:admin-code:355-12-1 |
| Plumbing                        | Alabama State Building Code companion code          | DCM-covered projects                                        | IPC                           | 2021                                    | verified_for_dcm_scope              | July 1, 2022                                                                               | none parsed                                                                                      | src:usa-al:dcm:state-building-code; src:usa-al:admin-code:355-12-1 |
| Fuel Gas                        | Alabama State Building Code companion code          | DCM-covered projects                                        | IFGC                          | 2021                                    | verified_for_dcm_scope              | July 1, 2022                                                                               | none parsed                                                                                      | src:usa-al:dcm:state-building-code; src:usa-al:admin-code:355-12-1 |
| Mechanical                      | Alabama State Building Code companion code          | DCM-covered projects                                        | IMC                           | 2021 IMC with specified 2024 amendments | verified_for_dcm_scope              | 2021 IMC adopted July 1, 2022; amended provisions effective March 17, 2025                 | Use amended DCM rule for covered projects                                                        | src:usa-al:dcm:state-building-code; src:usa-al:admin-code:355-12-1 |
| Electrical                      | Alabama State Building Code companion code          | DCM-covered projects                                        | NFPA 70 / NEC                 | 2020                                    | verified_for_dcm_scope              | July 1, 2022                                                                               | none parsed                                                                                      | src:usa-al:dcm:state-building-code; src:usa-al:admin-code:355-12-1 |
| Fire - construction references  | Alabama State Building Code companion code          | DCM-covered projects                                        | IFC                           | 2021 IFC with specified 2024 amendments | verified_for_dcm_scope              | amended provisions effective March 17, 2025                                                | DCM scope only; reconcile with State Fire Marshal                                                | src:usa-al:dcm:state-building-code; src:usa-al:admin-code:355-12-1 |
| Fire - operational / prevention | State Fire Marshal regulation                       | State Fire Marshal scope                                    | IFC                           | 2021 IFC with appendices and exceptions | verified                            | effective January 1, 2025                                                                  | Applies to projects whose architectural-services contract date is on or after the effective date | src:usa-al:sfmo:482-2-101                                         |
| Energy - DCM state-funded scope | Alabama State Building Code energy reference        | State-funded / state-instrumentality buildings in DCM scope | ASHRAE 90.1                   | 2013                                    | verified_for_dcm_scope              | July 1, 2022                                                                               | none parsed                                                                                      | src:usa-al:dcm:state-building-code; src:usa-al:admin-code:355-12-1 |
| Commercial Energy               | Alabama Commercial Energy Code                      | Commercial buildings outside DCM scope                      | IECC or ASHRAE basis expected | current edition not verified            | transition_pending                  | Authority transferred effective October 1, 2024; final code publication requires follow-up | Proposed-code process pending / unresolved                                                       | src:usa-al:dcm:state-building-code                                |
| Residential Building            | Alabama Residential Building Code                   | Residential construction                                    | IRC-based                     | board-current edition not verified      | authority_verified; edition_pending | HBLB has sole authority; compliance framework begins January 1, 2027                       | Local code may apply; board code applies where no local code exists                              | src:usa-al:hblb:law-2026                                          |
| Residential Energy              | Alabama Residential Energy Code                     | Residential construction                                    | IECC residential provisions   | board-current edition not verified      | authority_verified; edition_pending | HBLB has sole authority; local energy provisions limited by state code                     | Local provisions may not exceed Alabama Residential Energy Code except limited grounds           | src:usa-al:hblb:law-2026                                          |
| Legacy Residential Building     | Alabama Energy and Residential Codes Board rule     | Legacy / prior framework                                    | IRC                           | 2015 IRC as modified                    | legacy_needs_reconciliation         | effective October 1, 2016                                                                  | supersession by HBLB process must be confirmed                                                   | src:usa-al:adeca:305-2-4-legacy                                   |
| Legacy Residential Energy       | Alabama Energy and Residential Codes Board rule     | Legacy / prior framework                                    | IECC                          | 2015 IECC as modified                   | legacy_needs_reconciliation         | effective October 1, 2016                                                                  | supersession by HBLB process must be confirmed                                                   | src:usa-al:adeca:305-2-4-legacy                                   |
| Accessibility                   | Alabama State Building Code accessibility reference | DCM accessibility scope                                     | 2010 ADA Standards            | 2010                                    | verified_for_dcm_scope              | April 18, 2011                                                                             | ADA requirements supersede IBC accessibility where applicable                                    | src:usa-al:dcm:state-building-code; src:usa-al:admin-code:355-12-1 |
| Storm Shelter                   | Alabama State Building Code supplement              | DCM-covered scope where applicable                          | ICC/NSSA-500                  | 2020                                    | verified_for_dcm_scope              | July 1, 2022                                                                               | none parsed                                                                                      | src:usa-al:dcm:state-building-code; src:usa-al:admin-code:355-12-1 |
| Fire Alarm / Signaling          | Alabama State Building Code supplement              | DCM-covered scope where applicable                          | NFPA 72                       | 2019                                    | verified_for_dcm_scope              | July 1, 2022                                                                               | none parsed                                                                                      | src:usa-al:dcm:state-building-code; src:usa-al:admin-code:355-12-1 |
| Elevator / Conveyance           | unknown                                             | unknown                                                     | unknown                       | unknown                                 | unresolved                          | not verified                                                                               | not verified                                                                                     | unresolved                                                       |

### 3.2 Adoption Records

| Adoption Record ID                                | Code Family                   | Authority           | Edition / Rule                          | Effective Date | Verification Status |
| ------------------------------------------------- | ----------------------------- | ------------------- | --------------------------------------- | -------------- | ------------------- |
| adoption:usa-al:dcm:ibc-2021-amended-2025          | Building                      | ahj:usa-al:dcm      | 2021 IBC with specified 2024 amendments | 2025-03-17     | verified            |
| adoption:usa-al:dcm:iebc-2021                      | Existing Building             | ahj:usa-al:dcm      | 2021 IEBC                               | 2022-07-01     | verified            |
| adoption:usa-al:dcm:ipc-2021                       | Plumbing                      | ahj:usa-al:dcm      | 2021 IPC                                | 2022-07-01     | verified            |
| adoption:usa-al:dcm:ifgc-2021                      | Fuel Gas                      | ahj:usa-al:dcm      | 2021 IFGC                               | 2022-07-01     | verified            |
| adoption:usa-al:dcm:imc-2021-amended-2025          | Mechanical                    | ahj:usa-al:dcm      | 2021 IMC with specified 2024 amendments | 2025-03-17     | verified            |
| adoption:usa-al:dcm:nec-2020                       | Electrical                    | ahj:usa-al:dcm      | 2020 NEC                                | 2022-07-01     | verified            |
| adoption:usa-al:dcm:ifc-2021-amended-2025          | Fire construction references  | ahj:usa-al:dcm      | 2021 IFC with specified 2024 amendments | 2025-03-17     | verified            |
| adoption:usa-al:sfmo:ifc-2021                      | Fire operational / prevention | ahj:usa-al:sfmo     | 2021 IFC with exceptions                | 2025-01-01     | verified            |
| adoption:usa-al:dcm:ashrae-90-1-2013               | Energy, DCM scope             | ahj:usa-al:dcm      | ASHRAE 90.1-2013                        | 2022-07-01     | verified            |
| adoption:usa-al:dcm:ada-2010                       | Accessibility                 | ahj:usa-al:dcm      | 2010 ADA Standards                      | 2011-04-18     | verified            |
| adoption:usa-al:hblb:residential-current           | Residential                   | ahj:usa-al:hblb     | Alabama Residential Building Code       | null           | pending             |
| adoption:usa-al:hblb:residential-energy-current    | Residential Energy            | ahj:usa-al:hblb     | Alabama Residential Energy Code         | null           | pending             |
| adoption:usa-al:genconbd:commercial-energy-current | Commercial Energy             | ahj:usa-al:genconbd | Alabama Commercial Energy Code          | null           | pending             |

---

## 4. Date Rules and Transition Logic

### 4.1 Summary

The report now separates DCM effective dates, State Fire Marshal project-trigger dates, residential statutory transition dates, county local-code effective-date rules, and commercial energy transition deadlines. Remaining date risk is concentrated in residential board publication and commercial energy final adoption.

### 4.2 Date Rule Table

| Rule ID                                          | Applies To                      | Rule Type              | Date / Period             | Trigger Condition                                                          | Prior Code Allowed?                 | Status             |
| ------------------------------------------------ | ------------------------------- | ---------------------- | ------------------------- | -------------------------------------------------------------------------- | ----------------------------------- | ------------------ |
| date-rule:usa-al:dcm-2025-amendments              | DCM IBC / IMC / IFC amendments  | effective_date         | 2025-03-17                | DCM-covered project subject to amended Alabama State Building Code         | not parsed                          | verified           |
| date-rule:usa-al:dcm-2022-base-codes              | DCM companion codes             | effective_date         | 2022-07-01                | DCM-covered project subject to Alabama State Building Code                 | not parsed                          | verified           |
| date-rule:usa-al:sfmo-ifc-contract-date           | State Fire Marshal 2021 IFC     | operative_trigger      | 2025-01-01                | Architectural-services contract date on or after effective date            | not parsed                          | verified           |
| date-rule:usa-al:residential-builder-compliance   | Residential construction        | mandatory_date         | 2027-01-01                | Residential builder compliance with applicable residential building code   | local code may apply                | verified           |
| date-rule:usa-al:county-code-effective            | County building-code adoption   | delayed_effective_date | 120 days after resolution | County commission adopts local code for unincorporated area                | limited insurance-claim exception   | verified           |
| date-rule:usa-al:post-2027-local-residential-code | Local residential code adoption | minimum_standard_rule  | 2027-01-01                | County/municipality adopts local residential building code after this date | local code must meet state minimums | verified           |
| date-rule:usa-al:commercial-energy-proposal       | Commercial energy code          | proposed_code_deadline | 2025-12-31                | Advisory committee submits proposed commercial energy code to board        | final adoption not verified         | partially_verified |

### 4.3 Pending or Future Code Changes

| Code Family                  | Future / Pending Code                                          | Announced Date             | Adoption Date | Effective Date                | Watch Status | Notes                                                           |
| ---------------------------- | -------------------------------------------------------------- | -------------------------- | ------------- | ----------------------------- | ------------ | --------------------------------------------------------------- |
| Residential Building         | Alabama Residential Building Code under HBLB                   | statutory process verified | null          | 2027-01-01 compliance trigger | active_watch | Final board-adopted edition and amendments need source capture. |
| Residential Energy           | Alabama Residential Energy Code under HBLB                     | statutory process verified | null          | 2027-01-01 compliance trigger | active_watch | Edition and published amendments need source capture.           |
| Commercial Energy            | Alabama Commercial Energy Code under General Contractors Board | transition verified        | null          | null                          | active_watch | Final code after advisory committee process not verified.       |
| DCM State Building Code      | Future DCM amendments                                          | none verified              | null          | null                          | active_watch | Monitor DCM and Ala. Admin. Code ch. 355-12-1.                  |
| State Fire Marshal Fire Code | Future Fire Marshal amendments                                 | none verified              | null          | null                          | active_watch | Monitor Regulation 482-2-101.                                   |

### 4.4 Special Applicability Rules

| Rule ID                                            | Code Family                  | Project Type / Scope                                                                                      | Trigger                                             | Rule Summary                                                                                                                 | Status   |
| -------------------------------------------------- | ---------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------- | -------- |
| applicability-rule:usa-al:dcm-scope                 | Building and companion codes | State building/construction, schoolhouses, hotels, movie theaters, and defined accessibility-review scope | Project falls within DCM jurisdiction               | DCM State Building Code applies only within defined scope.                                                                   | verified |
| applicability-rule:usa-al:dcm-exclusion             | Building and companion codes | Ordinary residential and ordinary commercial construction                                                 | Project is outside DCM special/public scope         | DCM does not generally have jurisdiction, except hotels/motels and movie theaters.                                           | verified |
| applicability-rule:usa-al:sfmo                      | Fire / life safety           | Fire Marshal-regulated projects                                                                           | Fire Marshal jurisdiction and contract-date trigger | 2021 IFC applies with exceptions for covered Fire Marshal projects.                                                          | verified |
| applicability-rule:usa-al:residential-local-code    | Residential                  | Residential construction in local jurisdiction with adopted building code                                 | Local code exists                                   | Applicable residential code is local adopted code.                                                                           | verified |
| applicability-rule:usa-al:residential-no-local-code | Residential                  | Residential construction where no local code exists                                                       | No local code exists                                | Board-adopted Alabama Residential Building Code applies.                                                                     | verified |
| applicability-rule:usa-al:residential-energy-limit  | Residential Energy           | Local residential energy amendments                                                                       | Local jurisdiction adopts/amends energy provisions  | Local residential energy code may not exceed Alabama Residential Energy Code except for local conditions or federal mandate. | verified |

---

## 5. State Amendments

### 5.1 Amendment Model

**State amendment structure:** split_by_authority

**Where amendments are published:**

* DCM amendments: Alabama State Building Code / Ala. Admin. Code ch. 355-12-1
* Fire Marshal amendments/exceptions: Fire Marshal regulations, especially Regulation 482-2-101
* Residential amendments: HBLB / Residential Building Code Advisory Council materials
* Commercial energy amendments: Alabama Licensing Board for General Contractors process, pending source capture

**Amendment parsing status:** partially_started

### 5.2 State Amendment Sources

| Amendment Source ID                        | Authority           | Scope                                                      | Status                      |
| ------------------------------------------ | ------------------- | ---------------------------------------------------------- | --------------------------- |
| amend-src:usa-al:dcm:355-12-1               | ahj:usa-al:dcm      | DCM State Building Code exceptions and 2024 insertions     | identified_not_fully_parsed |
| amend-src:usa-al:sfmo:482-2-101             | ahj:usa-al:sfmo     | State Fire Marshal IFC exceptions                          | identified_not_fully_parsed |
| amend-src:usa-al:hblb:residential           | ahj:usa-al:hblb     | Alabama Residential Building Code / Energy Code amendments | pending                     |
| amend-src:usa-al:genconbd:commercial-energy | ahj:usa-al:genconbd | Commercial energy code amendments                          | pending                     |

### 5.3 High-Impact State Amendments

| Amendment ID                            | Authority       | Code Family                  | Impact Summary                                                                                                                                                          | Status                      |
| --------------------------------------- | --------------- | ---------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------- |
| amend:usa-al:dcm:2025-ibc-2024-inserts   | ahj:usa-al:dcm  | Building                     | DCM rule incorporates specified 2024 IBC portions into 2021 IBC framework.                                                                                              | identified_not_fully_parsed |
| amend:usa-al:dcm:2025-imc-2024-inserts   | ahj:usa-al:dcm  | Mechanical                   | DCM rule incorporates specified 2024 IMC portions into 2021 IMC framework.                                                                                              | identified_not_fully_parsed |
| amend:usa-al:dcm:2025-ifc-2024-inserts   | ahj:usa-al:dcm  | Fire construction references | DCM rule incorporates specified 2024 IFC portions into 2021 IFC framework.                                                                                              | identified_not_fully_parsed |
| amend:usa-al:sfmo:ifc-exceptions         | ahj:usa-al:sfmo | Fire / life safety           | State Fire Marshal 2021 IFC adoption includes exceptions and appendix selection.                                                                                        | identified_not_fully_parsed |
| amend:usa-al:hblb:sprinkler-choice-limit | ahj:usa-al:hblb | Residential                  | State/local bodies may not adopt rules restricting consumer choice on residential fire sprinkler installation or requiring installation, subject to statutory language. | partially_verified          |

---

## 6. Local Enforcement and Local Amendment Authority

### 6.1 Local Enforcement Model

```yaml
local_enforcement:
  enforcement_id: "local-enforcement:usa-al"
  model: "hybrid_state_limited_scope_with_local_residential_and_local_code_enforcement"
  enforcing_entities:
    - "DCM for defined state/public/special-use construction and accessibility scope"
    - "State Fire Marshal for fire/life-safety scope"
    - "Local jurisdictions with permitting and inspection programs for residential construction"
    - "County commissions in unincorporated areas when local building laws/codes are adopted"
    - "Municipalities where local building codes and inspection programs exist"
  required_officials:
    - "local building official or local inspection authority, where locally established"
    - "State Fire Marshal or deputy, where Fire Marshal jurisdiction applies"
    - "DCM reviewer/inspector, where DCM jurisdiction applies"
  state_reserved_activities:
    - "DCM plan review/enforcement for DCM-covered projects"
    - "State Fire Marshal fire/life-safety enforcement"
    - "HBLB residential code adoption authority"
    - "Commercial energy code adoption authority under General Contractors Board transition"
  source_ids:
    - "src:usa-al:dcm:state-building-code"
    - "src:usa-al:admin-code:355-12-1"
    - "src:usa-al:sfmo:482-2-101"
    - "src:usa-al:hblb:law-2026"
  verification_status: "partially_verified"
  confidence: 0.76
```

### 6.2 Local Amendment Rule

```yaml
local_amendment_rule:
  rule_id: "local-amendment-rule:usa-al"
  model: "local_codes_allowed_subject_to_state_residential_minimums_and_energy_limits"
  applies_to_code_families:
    - "Residential"
    - "Residential Energy"
    - "Local building codes generally, subject to unparsed municipal/county enabling laws"
  approval_required: "not verified as approval; county resolutions and related materials may require filing/copy to HBLB"
  approving_authority_id: null
  filing_required: true
  registry_exists: "not verified"
  registry_source_ids: []
  legal_basis_source_ids:
    - "src:usa-al:hblb:law-2026"
  verification_status: "partially_verified"
  confidence: 0.72
```

### 6.3 Local Enforcement vs. Local Amendment Summary

Residential enforcement and local amendment authority should be modeled separately. HBLB has state adoption authority for residential building and residential energy codes, but local jurisdictions may enforce local codes where they exist. After January 1, 2027, local residential codes adopted by counties or municipalities must meet the minimum standards of the Alabama Residential Building Code and Alabama Residential Energy Code. Residential energy provisions are more constrained because local amendments may not exceed the Alabama Residential Energy Code except under specified conditions.

### 6.4 Known Local Amendment Registries

No statewide local amendment registry was verified in this pass.

### 6.5 Municipality-Specific Known Amendments

No municipality-specific amendments were parsed in this pass.

---

## 7. Jurisdiction and AHJ Resolution

### 7.1 Jurisdiction Resolution Model

**Resolver status:** `partially_started`

Jurisdiction stack:

```text
Address
-> State: Alabama
-> Project funding/ownership/scope screen
   -> DCM covered? state building/construction, schoolhouse, hotel/motel, movie theater, covered public accessibility review
-> Fire/life-safety screen
   -> State Fire Marshal jurisdiction?
-> Residential screen
   -> One- or two-family / residential builder scope?
   -> Local jurisdiction has adopted residential code?
   -> If yes: local code applies
   -> If no: Alabama Residential Building Code applies
-> Commercial energy screen
   -> Determine current Alabama Commercial Energy Code authority and adopted edition
-> County / municipality
   -> incorporated municipality or unincorporated county?
   -> local building department / inspection program?
-> Special districts, if applicable
-> Building AHJ
-> Fire AHJ
-> Trade-specific AHJs
-> Applicable state code adoption records
-> Applicable local amendment records
```

### 7.2 Boundary Data Sources

| Boundary Type    | Source                         | Source ID                         | Coverage  | Update Frequency       | Status   |
| ---------------- | ------------------------------ | --------------------------------- | --------- | ---------------------- | -------- |
| State            | Alabama official legal sources | src:usa-al:agency:legislature-home | statewide | when state law changes | verified |
| County           | not selected                   | src:...                           | statewide | unknown                | pending  |
| Municipality     | not selected                   | src:...                           | statewide | unknown                | pending  |
| Fire District    | not selected                   | src:...                           | unknown   | unknown                | pending  |
| Special District | not selected                   | src:...                           | unknown   | unknown                | pending  |

### 7.3 AHJ Contact Data

No complete AHJ contact dataset was populated. DCM, State Fire Marshal, HBLB, and local building departments should be added as separate contact-source objects in a future pass.

---

## 8. Source Registry

### 8.1 Official Sources

| Source ID                         | Title                                            | Source Type          | Publisher                                                 | Accessed Date | Snapshot ID      | Checksum         | Status             |
| --------------------------------- | ------------------------------------------------ | -------------------- | --------------------------------------------------------- | ------------- | ---------------- | ---------------- | ------------------ |
| src:usa-al:dcm:state-building-code | Alabama State Building Code page                 | agency_page          | Alabama Department of Finance / DCM                       | 2026-06-25    | snapshot-pending | snapshot-pending | verified           |
| src:usa-al:admin-code:355-12-1     | Ala. Admin. Code ch. 355-12-1                    | admin_code_pdf       | Alabama Legislative Services Agency / Administrative Code | 2026-06-25    | snapshot-pending | snapshot-pending | verified           |
| src:usa-al:sfmo:about              | Alabama State Fire Marshal About page            | agency_page          | Alabama State Fire Marshal's Office                       | 2026-06-25    | snapshot-pending | snapshot-pending | verified           |
| src:usa-al:sfmo:regulations        | Alabama State Fire Marshal Regulations page      | agency_page          | Alabama State Fire Marshal's Office                       | 2026-06-25    | snapshot-pending | snapshot-pending | verified           |
| src:usa-al:sfmo:482-2-101          | Fire Marshal Regulation 482-2-101                | regulation_pdf       | Alabama State Fire Marshal's Office                       | 2026-06-25    | snapshot-pending | snapshot-pending | verified           |
| src:usa-al:sfmo:482-2-109          | Fire Marshal Regulation 482-2-109                | regulation_pdf       | Alabama State Fire Marshal's Office                       | 2026-06-25    | snapshot-pending | snapshot-pending | partially_verified |
| src:usa-al:hblb:law-2026           | HBLB law and administrative rules compilation    | statute_and_rule_pdf | Alabama Home Builders Licensure Board                     | 2026-06-25    | snapshot-pending | snapshot-pending | verified           |
| src:usa-al:adeca:305-2-4-legacy    | Alabama Energy and Residential Codes Board rules | admin_code_pdf       | ADECA / AERC Board                                        | 2026-06-25    | snapshot-pending | snapshot-pending | legacy             |
| src:usa-al:agency:legislature-home | Alabama Legislature home                         | agency_page          | Alabama Legislature                                       | 2026-06-24    | snapshot-pending | snapshot-pending | verified           |
| src:usa-al:agency:code-portal      | Code of Alabama portal                           | code_portal          | Alabama Legislature                                       | 2026-06-24    | snapshot-pending | snapshot-pending | discovery_only     |
| src:usa-al:agency:admin-code       | Administrative Code portal                       | admin_code_portal    | Alabama Legislature                                       | 2026-06-24    | snapshot-pending | snapshot-pending | verified           |

### 8.2 Official-Source Caveats

| Source ID                      | Caveat Type              | Caveat Summary                                                                                                                                              | Production Treatment           |
| ------------------------------ | ------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------ |
| src:usa-al:agency:admin-code    | javascript_app           | Administrative Code portal can be JS-driven; direct PDFs/API endpoints are preferable for production extraction.                                            | browser_or_direct_pdf_required |
| src:usa-al:agency:code-portal   | discovery_only           | The Code of Alabama portal confirms the code repository exists, but this pass used more targeted official sources for substantive building-code content.    | discovery_only                 |
| src:usa-al:sfmo:482-2-109       | potential_scope_conflict | Older Fire Marshal fuel gas/mechanical/existing-building regulation must be reconciled with newer Regulation 482-2-101 language and DCM scope.              | manual_review_required         |
| src:usa-al:adeca:305-2-4-legacy | legacy_authority         | ADECA/AERC residential and energy-code rules are affected by post-2024 authority transfers; do not treat as final current authority without reconciliation. | legacy_reference_only          |
| src:usa-al:hblb:law-2026        | board_publication_needed | HBLB authority is verified, but final board-adopted residential code edition and amendments still need direct publication or minutes.                       | follow_up_required             |

### 8.3 Supplemental Sources

None used as controlling authority. The report prioritizes official agency, administrative-code, and board materials.

### 8.4 Source Extraction Metadata

| Source ID                         | Parser                     | Parser Version | Extracted At         | Extraction Confidence | Requires OCR | Requires Browser Automation | Manual Review Required |
| --------------------------------- | -------------------------- | -------------- | -------------------- | --------------------: | ------------ | --------------------------- | ---------------------- |
| src:usa-al:dcm:state-building-code | browser_manual             | 0.2            | 2026-06-25T00:00:00Z |                  0.96 | no           | no                          | yes                    |
| src:usa-al:admin-code:355-12-1     | pdf_text_plus_visual_check | 0.2            | 2026-06-25T00:00:00Z |                  0.94 | no           | no                          | yes                    |
| src:usa-al:sfmo:about              | browser_manual             | 0.2            | 2026-06-25T00:00:00Z |                  0.90 | no           | no                          | yes                    |
| src:usa-al:sfmo:482-2-101          | pdf_text_plus_visual_check | 0.2            | 2026-06-25T00:00:00Z |                  0.92 | no           | no                          | yes                    |
| src:usa-al:sfmo:482-2-109          | pdf_text                   | 0.2            | 2026-06-25T00:00:00Z |                  0.72 | no           | no                          | yes                    |
| src:usa-al:hblb:law-2026           | pdf_text_plus_visual_check | 0.2            | 2026-06-25T00:00:00Z |                  0.88 | no           | no                          | yes                    |
| src:usa-al:adeca:305-2-4-legacy    | pdf_text                   | 0.2            | 2026-06-25T00:00:00Z |                  0.78 | no           | no                          | yes                    |
| src:usa-al:agency:legislature-home | browser_manual             | 0.1            | 2026-06-24T00:00:00Z |                  0.98 | no           | yes                         | yes                    |
| src:usa-al:agency:code-portal      | browser_manual             | 0.1            | 2026-06-24T00:00:00Z |                  0.98 | no           | yes                         | yes                    |
| src:usa-al:agency:admin-code       | browser_manual             | 0.1            | 2026-06-24T00:00:00Z |                  0.90 | no           | yes                         | yes                    |

---

## 9. Verification and QA

### 9.1 Field-Level Verification Summary

| Record ID                  | Field                   | Value                                                                            | Verification Status | Confidence | Source IDs                                                       | Notes                                                                                           |
| -------------------------- | ----------------------- | -------------------------------------------------------------------------------- | ------------------- | ---------: | ---------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| report:usa-al               | report.status           | partially_verified                                                               | verified            |       1.00 | src:usa-al:dcm:state-building-code; src:usa-al:hblb:law-2026       | Reclassified from draft after official-source extraction.                                       |
| report:usa-al               | risk.overall_confidence | 0.74                                                                             | verified            |       1.00 | all verified source IDs                                          | Confidence reflects verified major authority structure with remaining unresolved code editions. |
| ahj:usa-al:dcm             | authority.name          | Alabama Department of Finance / DCM                                              | verified            |       0.95 | src:usa-al:dcm:state-building-code; src:usa-al:admin-code:355-12-1 | Primary DCM scope authority.                                                                    |
| ahj:usa-al:sfmo            | authority.name          | Alabama State Fire Marshal's Office                                              | verified            |       0.90 | src:usa-al:sfmo:about; src:usa-al:sfmo:482-2-101                   | Fire/life-safety authority.                                                                     |
| ahj:usa-al:hblb            | authority.role          | Sole authority for Alabama Residential Building Code and Residential Energy Code | verified            |       0.90 | src:usa-al:hblb:law-2026                                          | Edition still pending.                                                                          |
| local-enforcement:usa-al    | model                   | hybrid_state_limited_scope_with_local_residential_and_local_code_enforcement     | partially_verified  |       0.76 | src:usa-al:dcm:state-building-code; src:usa-al:hblb:law-2026       | AHJ routing still needs local dataset.                                                          |
| local-amendment-rule:usa-al | model                   | local_codes_allowed_subject_to_state_residential_minimums_and_energy_limits      | partially_verified  |       0.72 | src:usa-al:hblb:law-2026                                          | Registry not found.                                                                             |

### 9.2 Validation Checks

| Check                                                          | Status  | Notes                                                                                                               |
| -------------------------------------------------------------- | ------- | ------------------------------------------------------------------------------------------------------------------- |
| All source IDs resolve                                         | pass    | New source IDs listed in section 8.                                                                                 |
| All authority IDs resolve                                      | pass    | DCM, SFMO, HBLB, General Contractors Board transition, and local authorities are defined.                           |
| All current code families have adoption records                | partial | DCM and Fire Marshal records are populated; residential, commercial energy, and elevator/conveyance remain pending. |
| Building and operational fire code are separated               | pass    | DCM fire construction references and State Fire Marshal operational/fire-prevention authority are separated.        |
| Adoption/effective/operative/mandatory dates are not conflated | pass    | DCM dates, Fire Marshal contract trigger, residential 2027 trigger, and county 120-day rule are separated.          |
| Effective dates are valid ISO dates                            | pass    | Verified dates are ISO-formatted.                                                                                   |
| No impossible date sequences                                   | pass    | No conflicting verified date sequences entered.                                                                     |
| Transition rules have explicit trigger conditions              | partial | Major triggers captured; full DCM permit/plan-review transition details still need parsing.                         |
| Permit-date logic is captured where applicable                 | partial | Fire Marshal architectural-services contract trigger captured; DCM permit/plan-review transition not fully parsed.  |
| Local enforcement model classified                             | pass    | Classified as hybrid / limited state scope with local residential enforcement.                                      |
| Local amendment rule classified                                | partial | Residential local-code constraints captured; full municipal/county amendment universe not complete.                 |
| AHJ confirmation metadata present                              | fail    | No complete AHJ contact or jurisdiction boundary dataset entered.                                                   |
| Official-source caveats captured                               | pass    | JS portal, legacy ADECA, and Fire Marshal reconciliation caveats captured.                                          |

---

## 10. Open Issues and Review Queue

| Issue ID        | Severity | Topic                             | Description                                                                                                                                       | Needed Action                                                                                                        | Owner | Due Date | Status |
| --------------- | -------- | --------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | ----- | -------- | ------ |
| issue:usa-al:001 | high     | residential code edition          | HBLB authority is verified, but the final current Alabama Residential Building Code and Residential Energy Code edition/package was not captured. | Pull HBLB board adoption materials, council recommendations, minutes, or published code package.                     | null  | null     | open   |
| issue:usa-al:002 | high     | commercial energy code            | Commercial energy authority transition is verified, but the final current Alabama Commercial Energy Code was not confirmed.                       | Extract Alabama Licensing Board for General Contractors commercial energy code materials and any final board action. | null  | null     | open   |
| issue:usa-al:003 | high     | Fire Marshal scope reconciliation | Fire Marshal Regulation 482-2-109 legacy trade-code adoptions need reconciliation with Regulation 482-2-101 and DCM code references.              | Review current Alabama Administrative Code 482 rules and Fire Marshal regulation history.                            | null  | null     | open   |
| issue:usa-al:004 | medium   | local AHJ routing                 | Local enforcement is classified, but no municipality/county AHJ dataset is built.                                                                 | Select boundary and building department data sources; build address-to-AHJ resolver.                                 | null  | null     | open   |
| issue:usa-al:005 | medium   | local amendment registry          | No statewide registry of local amendments was verified.                                                                                           | Search HBLB, county, municipal, and code-board materials for filing/registry requirements.                           | null  | null     | open   |
| issue:usa-al:006 | medium   | DCM amendments parsing            | DCM 2025 amendments are identified but not fully parsed into high-impact amendment records.                                                       | Parse each DCM exception/amendment by code section.                                                                  | null  | null     | open   |
| issue:usa-al:007 | low      | elevator / conveyance             | Elevator/conveyance authority and code adoption were not parsed.                                                                                  | Search Alabama Department of Labor, insurance, and administrative-code materials.                                    | null  | null     | open   |

---

## 11. Change Monitoring

### 11.1 Watch Targets

| Watch ID                               | Source ID                         | Watch Type      | Frequency | Trigger Condition                                                    | Last Checked | Status  |
| -------------------------------------- | --------------------------------- | --------------- | --------- | -------------------------------------------------------------------- | ------------ | ------- |
| watch:usa-al:dcm-state-building-code    | src:usa-al:dcm:state-building-code | html_diff       | monthly   | DCM code list, edition, amendment, or scope change                   | 2026-06-25   | active  |
| watch:usa-al:admin-code-355-12-1        | src:usa-al:admin-code:355-12-1     | pdf_diff        | monthly   | Alabama State Building Code rule amendment                           | 2026-06-25   | active  |
| watch:usa-al:sfmo-reg-101               | src:usa-al:sfmo:482-2-101          | pdf_diff        | monthly   | Fire Marshal IFC edition or exception change                         | 2026-06-25   | active  |
| watch:usa-al:sfmo-regulations           | src:usa-al:sfmo:regulations        | html_diff       | monthly   | Fire Marshal regulation list change                                  | 2026-06-25   | active  |
| watch:usa-al:hblb-law-rules             | src:usa-al:hblb:law-2026           | pdf_diff        | monthly   | Residential code authority, code edition, or enforcement rule change | 2026-06-25   | active  |
| watch:usa-al:genconbd-commercial-energy | src:usa-al:dcm:state-building-code | manual_followup | monthly   | Commercial energy code final adoption or publication                 | 2026-06-25   | active  |
| watch:usa-al:adeca-legacy               | src:usa-al:adeca:305-2-4-legacy    | legacy_check    | quarterly | Legacy rule superseded, repealed, or amended                         | 2026-06-25   | pending |
| watch:usa-al:code-portal                | src:usa-al:agency:code-portal      | html_diff       | quarterly | Code of Alabama portal restructuring or code text update             | 2026-06-24   | active  |
| watch:usa-al:admin-code-portal          | src:usa-al:agency:admin-code       | html_diff       | quarterly | Administrative Code portal restructuring                             | 2026-06-24   | active  |

### 11.2 Changelog

| Date       | Change                                                                                                                          | Record IDs                                                                              | Source IDs                                                                                                          | Changed By | Notes                                                                                                                    |
| ---------- | ------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | ---------- | ------------------------------------------------------------------------------------------------------------------------ |
| 2026-06-24 | Replaced Alabama stub with conservative draft and official Alabama portal sources                                               | report:usa-al                                                                            | src:usa-al:agency:legislature-home; src:usa-al:agency:code-portal; src:usa-al:agency:admin-code                        | Codex      | Substantive code adoption details remained unresolved.                                                                   |
| 2026-06-25 | Updated report from draft to partially verified; filled DCM, State Fire Marshal, HBLB, local residential, and date-rule records | report:usa-al; ahj:usa-al:dcm; ahj:usa-al:sfmo; ahj:usa-al:hblb; local-enforcement:usa-al | src:usa-al:dcm:state-building-code; src:usa-al:admin-code:355-12-1; src:usa-al:sfmo:482-2-101; src:usa-al:hblb:law-2026 | ChatGPT    | Remaining high-risk gaps are residential code edition, commercial energy final code, AHJ routing, and amendment parsing. |
