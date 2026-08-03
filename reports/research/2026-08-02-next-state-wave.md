# Next-State Regulatory Source Inventory

- **Research date:** 2026-08-02
- **Scope:** Virginia, Oregon, North Carolina, Massachusetts, New York, and California
- **Status:** Initial official-source inventory. This file is not an executable jurisdiction profile and does not claim that any state is pilot-ready.
- **Method:** Prefer statutes, administrative regulations, state code-agency pages, adoption notices, and official local-amendment registries. Preserve unresolved fields rather than inferring them from model-code labels.

## Recommended onboarding sequence

| Wave | State | Initial classification | Main reason |
| --- | --- | --- | --- |
| A | Virginia | Data-first | Uniform statewide construction code with local enforcement and strong preemption evidence. |
| A | Oregon | Data-first | Official code-family pages expose current editions, effective dates, mandatory dates, and phase-in periods. |
| A | North Carolina | Data-first with date blocker | Current 2018 code is clear, but the 2024 replacement date remains conditional rather than calendar-fixed. |
| B | Massachusetts | Resolver extension required | Building code is statewide, but energy policy varies by municipality among Base, Stretch, and Specialized codes. |
| B | New York | Resolver extension required | Statewide Uniform Code excludes New York City, and approved more-restrictive local standards can modify the state baseline. |
| B | California | Dedicated subsystem | Title 24 is statewide, but state-agency jurisdiction, occupancy scope, local amendments, and temporary statutory limits create a high-dimensional policy model. |

## Virginia

### Preliminary findings

- The Board of Housing and Community Development adopts and amends the Virginia Uniform Statewide Building Code (USBC).
- The 2021 Virginia codes became effective on 2024-01-18.
- The 2021 construction code incorporates the 2021 International Building Code, with Virginia amendments.
- Construction and rehabilitation enforcement is generally assigned to local building departments. Small towns may be enforced by their county when they do not elect to administer the code.
- The statewide code supersedes county, municipal, other political-subdivision, and state-agency building codes, subject to statutory exceptions for matters outside its coverage and certain land-use or floodplain regulations.
- The administrative code provides a one-year period beginning on the 2021 code effective date during which qualifying permit applications may use the prior edition. The exact request model and end-date representation should be verified from the full regulation before compilation.

### Official source registry

| Source ID | Source | Supports |
| --- | --- | --- |
| `src:usa-va:dhcd-codes` | https://www.dhcd.virginia.gov/codes | Current code cycle and 2024-01-18 effective date. |
| `src:usa-va:dhcd-usbc` | https://www.dhcd.virginia.gov/virginia-uniform-statewide-building-code-usbc | Adopting authority, USBC composition, local enforcement overview. |
| `src:usa-va:vacode-36-98` | https://law.lis.virginia.gov/vacode/title36/chapter6/section36-98/ | Statewide code authority and preemption. |
| `src:usa-va:vacode-36-105` | https://law.lis.virginia.gov/vacode/title36/chapter6/section36-105/ | Local building-department enforcement and appeals structure. |
| `src:usa-va:vac-13-5-63-10` | https://law.lis.virginia.gov/admincode/title13/agency5/chapter63/section10/ | 2021 IBC incorporation into the Virginia Construction Code. |
| `src:usa-va:vac-13-5-63-30` | https://law.lis.virginia.gov/admincode/title13/agency5/chapter63/section30/ | Applicability and prior-code transition rule. |

### Next extraction tasks

1. Extract all current code-family incorporations and editions from 13VAC5-63 and the fire-code chapter.
2. Represent the one-year prior-edition option as a date rule tied to permit-application date.
3. Verify state-owned building, industrialized building, amusement-device, elevator, and operational fire-code authority paths.
4. Add incorporated and unincorporated fixtures, including a town below the statutory population threshold.

## Oregon

### Preliminary findings

- The Building Codes Division adopts, amends, and interprets the specialty codes that make up the Oregon State Building Code.
- The 2025 Oregon Structural Specialty Code (OSSC) is based on the 2024 IBC, IFC, and IEBC. Its administrative provisions became effective and mandatory on 2025-10-01. Construction provisions became effective on 2025-10-01, permitted either the 2022 or 2025 OSSC during a six-month phase-in, and became mandatory on 2026-04-01.
- The 2025 Oregon Energy Efficiency Specialty Code became effective on 2025-01-01 and mandatory on 2025-07-01; it is based on ASHRAE 90.1-2022.
- The current residential code is the 2023 Oregon Residential Specialty Code. Construction provisions became mandatory on 2024-04-01. A 2026 residential-code adoption process is underway with an anticipated adoption date of 2026-10-01, which must not be treated as an effective date until formally adopted.
- The adopted-code index identifies the 2023 Oregon Electrical Specialty Code based on the 2023 NEC, the 2023 Oregon Plumbing Specialty Code based on the 2021 UPC, and current mechanical-code materials.

### Official source registry

| Source ID | Source | Supports |
| --- | --- | --- |
| `src:usa-or:bcd-state-code` | https://www.oregon.gov/bcd/codes-stand/pages/index.aspx | State code authority and specialty-code program structure. |
| `src:usa-or:bcd-adopted-codes` | https://www.oregon.gov/bcd/codes-stand/pages/adopted-codes.aspx | Current code-family inventory and effective/mandatory dates. |
| `src:usa-or:bcd-ossc-2025` | https://www.oregon.gov/bcd/codes-stand/Pages/ossc-adoption.aspx | 2025 OSSC model-code basis, phase-in, and mandatory date. |
| `src:usa-or:bcd-oeesc-2025` | https://www.oregon.gov/bcd/codes-stand/Pages/oeesc-adoption.aspx | 2025 commercial energy-code basis and dates. |
| `src:usa-or:bcd-orsc-2026-process` | https://www.oregon.gov/bcd/codes-stand/Pages/orsc-adoption.aspx | Pending 2026 residential-code adoption process. |
| `src:usa-or:bcd-code-history` | https://www.oregon.gov/bcd/codes-stand/pages/codebook-history.aspx | Historical edition ranges and replacement dates. |

### Next extraction tasks

1. Build separate adoption records for structural, residential, energy, electrical, mechanical, and plumbing families.
2. Encode phase-in periods independently by code family; Oregon should become the primary transition-rule fixture.
3. Verify statewide preemption, local enforcement delegation, local amendment authority, and state-reserved project classes from Oregon statutes and administrative rules.
4. Capture interim amendments and errata as source-health or amendment-set records rather than silently folding them into the base edition.

## North Carolina

### Preliminary findings

- The Office of the State Fire Marshal still lists the 2018 North Carolina codes under current codes. Those codes became effective on 2019-01-01 and are based primarily on the 2015 International Codes with North Carolina amendments.
- Implementation of the 2024 North Carolina State Building Code was postponed. The published rule is conditional: implementation occurs 12 months after the State Fire Marshal certifies publication/distribution and constitution of the Residential Code Council. A calendar effective date must not be invented.
- The North Carolina State Building Code applies throughout the state, subject to the statutory local-code provisions and specialized allocations.
- The State Fire Marshal has general supervisory and enforcement authority for major construction code families, while local governments must provide inspection services through their own, joint, contracted, county, or state-arranged inspection programs.
- State buildings follow a separate enforcement path through the Department of Administration, Office of State Construction.

### Official source registry

| Source ID | Source | Supports |
| --- | --- | --- |
| `src:usa-nc:osfm-current-past-codes` | https://www.ncosfm.gov/index.php/codes/codes-current-and-past | Current 2018 code status and 2024 code materials. |
| `src:usa-nc:osfm-2024-delay` | https://www.ncosfm.gov/news/press-releases/2025/04/07/north-carolina-delays-implementation-2024-state-building-code | Conditional trigger for future 2024 code implementation. |
| `src:usa-nc:gs-143-138` | https://www.ncleg.gov/EnactedLegislation/Statutes/HTML/BySection/Chapter_143/GS_143-138.html | Statewide code adoption and effect on local codes. |
| `src:usa-nc:gs-143-139` | https://www.ncleg.gov/enactedlegislation/statutes/html/bysection/chapter_143/gs_143-139.html | State Fire Marshal oversight, enforcement allocations, and state-building exception. |
| `src:usa-nc:gs-160d-1102` | https://www.ncleg.gov/EnactedLegislation/Statutes/HTML/BySection/Chapter_160D/GS_160D-1102.html | Required local inspection-service arrangements. |

### Next extraction tasks

1. Confirm whether the State Fire Marshal has issued the statutory certification or any later implementation notice after the 2025 delay publication.
2. Model the 2024 code as `pending` with a condition-based date rule, not a future adoption with a guessed date.
3. Extract current editions for each 2018 code family and their state amendments.
4. Add ordinary local, state-owned, tribal, and state-intervention enforcement fixtures.

## Massachusetts

### Preliminary findings

- The Board of Building Regulations and Standards adopts the Massachusetts State Building Code, 780 CMR.
- The tenth edition is the current building code. It first became effective on 2024-10-11; the concurrency period ended on 2025-06-30, after which the tenth edition became the only current edition.
- The tenth edition is based on modified 2021 I-Codes, including the IBC, IRC, IEBC, IMC, IECC, ISPSC, and IFC.
- Energy policy has three current levels: Base, municipal opt-in Stretch, and municipal opt-in Specialized. The official municipality inventory includes an effective date for each community and was updated with a 2026-06-04 snapshot.
- Municipality-conditioned energy policy cannot be represented accurately by the current state profile's incorporated/unincorporated split alone.

### Official source registry

| Source ID | Source | Supports |
| --- | --- | --- |
| `src:usa-ma:780-cmr` | https://www.mass.gov/massachusetts-state-building-code-780-cmr | BBRS authority and current state building-code structure. |
| `src:usa-ma:tenth-edition` | https://www.mass.gov/handbook/tenth-edition-of-the-ma-state-building-code-780 | Tenth-edition model-code basis and concurrency rule. |
| `src:usa-ma:edition-dates` | https://www.mass.gov/info-details/dates-and-editions-of-massachusetts-building-code-780-cmr | Effective and expiration dates for current and historical editions. |
| `src:usa-ma:energy-codes` | https://www.mass.gov/info-details/massachusetts-building-energy-codes | Base, Stretch, and Specialized energy-code definitions. |
| `src:usa-ma:municipal-energy-map` | https://www.mass.gov/info-details/massachusetts-building-energy-code-adoption-by-municipality | Municipality-level energy-code status and effective dates. |

### Next extraction tasks

1. Extract the tenth-edition adoption records and transition rule for base and residential volumes.
2. Acquire the municipality energy-code table in a machine-readable form and retain its dated snapshot metadata.
3. Add a resolver policy selector keyed by municipality or a maintained geographic policy layer before compiling the energy overrides.
4. Verify local building-official enforcement, state building inspector responsibilities, fire-code authority, and state-reserved project classes.

## New York

### Preliminary findings

- The 2025 New York State Uniform Code and Energy Code became effective on 2025-12-31 and replaced the 2020 editions.
- The state Uniform Code applies throughout New York except New York City, which retains its own construction code.
- New York City currently uses the 2022 NYC Construction Codes for ordinary new construction. Its enacted Existing Building Code is scheduled to become effective on 2027-07-17.
- A 2026 court order suspends implementation and enforcement of the 2025 state provisions prohibiting fossil-fuel equipment and systems in new buildings. The remainder of the 2025 adoption remains effective according to the Department of State notice.
- Local governments generally administer and enforce the Uniform Code. A locality may opt out, after which the Department of State provides enforcement.
- Local governments may enact more restrictive construction standards, but the standards require notice, petition, and Code Council approval before becoming effective.

### Official source registry

| Source ID | Source | Supports |
| --- | --- | --- |
| `src:usa-ny:uniform-code` | https://dos.ny.gov/uniform-fire-prevention-and-building-code | Statewide applicability, NYC exception, and local-standard authority. |
| `src:usa-ny:2025-notice` | https://dos.ny.gov/notice-adoption | 2025 editions, 2025-12-31 effective date, transition, and 2026 court-order notice. |
| `src:usa-ny:local-enforcement` | https://dos.ny.gov/code/local-government-state-agency-enforcement-programs | Local and state enforcement-program framework. |
| `src:usa-ny:local-authority-memo` | https://dos.ny.gov/node/23866 | Local enforcement and approval requirements for more restrictive local standards. |
| `src:usa-ny:nyc-codes` | https://www.nyc.gov/site/buildings/codes/codes.page | Current NYC code family inventory. |
| `src:usa-ny:nyc-existing-code` | https://www.nyc.gov/site/buildings/codes/existing-building-code.page | Future NYC Existing Building Code effective date. |

### Next extraction tasks

1. Create a statewide 2025 adoption timeline and retain the 2020-to-2025 optional transition period.
2. Represent New York City as an explicit geographic exception, not as a local amendment to the state code.
3. Add field- or provision-level suspension support for the court-ordered fossil-fuel provisions; a whole-code status is too coarse.
4. Inventory approved more-restrictive local standards and opted-out enforcement jurisdictions.

## California

### Preliminary findings

- The 2025 California Building Standards Code, Title 24, became effective on 2026-01-01 and is the current triennial edition through 2028-12-31, subject to errata, emergency standards, and intervening supplements.
- Title 24 is divided into multiple parts covering administration, building, residential, electrical, mechanical, plumbing, energy, wildland-urban interface, historical buildings, fire, existing buildings, CALGreen, and referenced standards.
- The California Building Standards Commission administers and coordinates the adoption process, approves standards proposed by multiple state agencies, codifies the approved standards, and directly handles certain state-owned and otherwise unassigned occupancies.
- Cities and counties may adopt qualifying local modifications based on climatic, geological, or topographical findings and must file the findings and amendments with the Commission before they become operative.
- The Commission maintains an edition-specific local-ordinance registry. The 2025 registry is actively receiving and processing filings.
- Statutory amendments effective in 2025 temporarily restrict many new local modifications affecting residential units from 2025-10-01 through 2031-06-01, subject to listed exceptions. This must be represented as a date-sensitive amendment-authority rule.

### Official source registry

| Source ID | Source | Supports |
| --- | --- | --- |
| `src:usa-ca:title-24-codes` | https://www.dgs.ca.gov/en/BSC/Codes | Current 2025 Title 24 parts and 2026-01-01 effective date. |
| `src:usa-ca:cbsc-about` | https://www.dgs.ca.gov/BSC/About | Commission authority and multi-agency adoption structure. |
| `src:usa-ca:local-amendments` | https://www.dgs.ca.gov/bsc/codes/local-amendments-to-building-standards---ordinances | Local-amendment registry and filing program. |
| `src:usa-ca:2025-ordinances` | https://www.dgs.ca.gov/BSC/Codes/2025-Ordinances | Edition-specific 2025 local ordinance index. |
| `src:usa-ca:hsc-17958-5` | https://leginfo.legislature.ca.gov/faces/codes_displaySection.xhtml?lawCode=HSC&sectionNum=17958.5. | Local modification findings and 2025-2031 residential restrictions. |
| `src:usa-ca:hsc-17958-7` | https://leginfo.legislature.ca.gov/faces/codes_displaySection.xhtml?lawCode=HSC&sectionNum=17958.7 | Filing prerequisites and Commission rejection authority. |
| `src:usa-ca:2025-code-changes` | https://www.dgs.ca.gov/BSC/Resources/2025-Title-24-California-Code-Changes | State change summary for the 2025 triennial edition. |

### Next extraction tasks

1. Map each Title 24 part to the proposing/adopting state agency and occupancy or project scope.
2. Create base statewide adoption records before attempting local amendment ingestion.
3. Snapshot and normalize the 2025 ordinance registry by city, county, ordinance number, code family, filing status, and source availability.
4. Add date-sensitive local-amendment-authority rules reflecting the temporary residential restriction and its exceptions.
5. Define resolver inputs for occupancy and state-agency jurisdiction before claiming California pilot readiness.

## Cross-state schema and resolver work exposed by this research

1. **Geography-conditioned policy:** Select a policy by named municipality or special geographic policy layer, required for Massachusetts energy codes and New York City.
2. **Conditional future dates:** Represent an adoption whose operative date depends on an official certification or other future event, required for North Carolina.
3. **Provision-level legal status:** Preserve a court-ordered suspension or conflict affecting only part of an otherwise current code, required for New York.
4. **Edition-specific local amendment registries:** Store and refresh local ordinances separately from the state baseline, required for California and useful for New York.
5. **Multi-agency occupancy jurisdiction:** Route state authority by project or occupancy class, required for California and useful for state-owned or specialized projects elsewhere.
6. **Phase-in periods:** Preserve effective, optional, mandatory, and replacement dates separately, with Oregon as the primary validation fixture.

## Immediate deliverables from this inventory

- Convert Virginia and Oregon into full source-backed state reports first.
- Build a North Carolina draft whose 2024 code remains explicitly pending until an official certification and implementation date are found.
- Design geography-conditioned policy selection before compiling Massachusetts, New York, or California into executable profiles.
- Keep every source in this inventory under source-health monitoring because several states publish interim amendments, court updates, code-cycle notices, or actively changing local registries.
