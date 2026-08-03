# Next-State Regulatory Source Inventory

- Research date: 2026-08-02
- Scope: Virginia, Oregon, North Carolina, Massachusetts, New York, and California
- Status: official-source seed, not an executable profile

This inventory records the first authoritative sources, supported facts, unresolved questions, and resolver implications for the next state-coverage wave. Facts remain conservative until they are extracted into full state reports and validated profiles.

## Onboarding order

| State          | Track                   | Primary reason                                                                                     |
| -------------- | ----------------------- | -------------------------------------------------------------------------------------------------- |
| Virginia       | Data-first              | Uniform statewide code, strong preemption, local enforcement.                                      |
| Oregon         | Data-first              | Excellent official edition and phase-in data by code family.                                       |
| North Carolina | Data-first with blocker | Current code is clear; replacement date is conditional.                                            |
| Massachusetts  | Resolver extension      | Energy code varies by municipality.                                                                |
| New York       | Resolver extension      | New York City is excluded from the state code; local stricter standards are approved individually. |
| California     | Dedicated subsystem     | Multi-agency scope, occupancy routing, and large local-amendment registry.                         |

## Virginia

Supported seed facts:

- The Board of Housing and Community Development adopts and amends the Virginia Uniform Statewide Building Code.
- The 2021 Virginia codes became effective on 2024-01-18.
- The Virginia Construction Code incorporates the 2021 IBC with Virginia amendments.
- Local building departments generally enforce construction and rehabilitation requirements.
- State law broadly supersedes local building codes, subject to stated exceptions.
- The 2021 edition included a one-year prior-edition option tied to permit applications; the exact date-rule representation still needs full-text extraction.

Official sources:

- `src:usa-va:dhcd-codes`: https://www.dhcd.virginia.gov/codes
- `src:usa-va:dhcd-usbc`: https://www.dhcd.virginia.gov/virginia-uniform-statewide-building-code-usbc
- `src:usa-va:vacode-36-98`: https://law.lis.virginia.gov/vacode/title36/chapter6/section36-98/
- `src:usa-va:vacode-36-105`: https://law.lis.virginia.gov/vacode/title36/chapter6/section36-105/
- `src:usa-va:vac-13-5-63-10`: https://law.lis.virginia.gov/admincode/title13/agency5/chapter63/section10/
- `src:usa-va:vac-13-5-63-30`: https://law.lis.virginia.gov/admincode/title13/agency5/chapter63/section30/

Next extraction: all code-family incorporations, state-owned and specialized authority paths, fire prevention authority, and incorporated/unincorporated fixtures.

## Oregon

Supported seed facts:

- The Building Codes Division adopts, amends, and interprets the specialty codes composing the Oregon State Building Code.
- The 2025 OSSC is based on the 2024 IBC, IFC, and IEBC. Construction provisions were optional from 2025-10-01 and mandatory from 2026-04-01.
- The 2025 OEESC became effective on 2025-01-01 and mandatory on 2025-07-01, based on ASHRAE 90.1-2022.
- The current residential code is the 2023 ORSC; a 2026 ORSC process is underway with an anticipated adoption date that must not be treated as operative until formally adopted.
- The official adopted-code index identifies current electrical, plumbing, mechanical, and other specialty codes.

Official sources:

- `src:usa-or:bcd-state-code`: https://www.oregon.gov/bcd/codes-stand/pages/index.aspx
- `src:usa-or:bcd-adopted-codes`: https://www.oregon.gov/bcd/codes-stand/pages/adopted-codes.aspx
- `src:usa-or:bcd-ossc-2025`: https://www.oregon.gov/bcd/codes-stand/Pages/ossc-adoption.aspx
- `src:usa-or:bcd-oeesc-2025`: https://www.oregon.gov/bcd/codes-stand/Pages/oeesc-adoption.aspx
- `src:usa-or:bcd-orsc-2026-process`: https://www.oregon.gov/bcd/codes-stand/Pages/orsc-adoption.aspx
- `src:usa-or:bcd-code-history`: https://www.oregon.gov/bcd/codes-stand/pages/codebook-history.aspx

Next extraction: separate adoption records and phase-in rules for every code family, local enforcement and amendment authority, state-reserved projects, interim amendments, and errata.

## North Carolina

Supported seed facts:

- OSFM still lists the 2018 North Carolina codes as current; they became effective on 2019-01-01.
- The 2024 code implementation was postponed and is triggered 12 months after an official certification of specified conditions. No calendar effective date should be inferred.
- The state code applies statewide, with State Fire Marshal supervision and local inspection-service delivery.
- State buildings have a separate Department of Administration enforcement path.

Official sources:

- `src:usa-nc:osfm-current-past-codes`: https://www.ncosfm.gov/index.php/codes/codes-current-and-past
- `src:usa-nc:osfm-2024-delay`: https://www.ncosfm.gov/news/press-releases/2025/04/07/north-carolina-delays-implementation-2024-state-building-code
- `src:usa-nc:gs-143-138`: https://www.ncleg.gov/EnactedLegislation/Statutes/HTML/BySection/Chapter_143/GS_143-138.html
- `src:usa-nc:gs-143-139`: https://www.ncleg.gov/enactedlegislation/statutes/html/bysection/chapter_143/gs_143-139.html
- `src:usa-nc:gs-160d-1102`: https://www.ncleg.gov/EnactedLegislation/Statutes/HTML/BySection/Chapter_160D/GS_160D-1102.html

Next extraction: check for a later certification, encode the 2024 edition as conditionally pending, extract each current 2018 code family, and add local, state-owned, tribal, and state-intervention fixtures.

## Massachusetts

Supported seed facts:

- The Board of Building Regulations and Standards adopts 780 CMR.
- The tenth edition is current. It became effective on 2024-10-11, with the ninth-edition concurrency period ending on 2025-06-30.
- It is based on modified 2021 I-Codes.
- Energy code policy has Base, municipal opt-in Stretch, and municipal opt-in Specialized levels. The state publishes municipality-level status and effective dates.

Official sources:

- `src:usa-ma:780-cmr`: https://www.mass.gov/massachusetts-state-building-code-780-cmr
- `src:usa-ma:tenth-edition`: https://www.mass.gov/handbook/tenth-edition-of-the-ma-state-building-code-780
- `src:usa-ma:edition-dates`: https://www.mass.gov/info-details/dates-and-editions-of-massachusetts-building-code-780-cmr
- `src:usa-ma:energy-codes`: https://www.mass.gov/info-details/massachusetts-building-energy-codes
- `src:usa-ma:municipal-energy-map`: https://www.mass.gov/info-details/massachusetts-building-energy-code-adoption-by-municipality

Resolver impact: municipality-conditioned policy is required before energy overrides can be compiled safely.

## New York

Supported seed facts:

- The 2025 State Uniform Code and Energy Code became effective on 2025-12-31.
- The Uniform Code applies statewide except New York City.
- New York City currently uses its 2022 Construction Codes; its Existing Building Code is scheduled for 2027-07-17.
- A 2026 court order suspends the state provisions prohibiting fossil-fuel equipment and systems while the rest of the 2025 adoption remains in force according to the state notice.
- Local governments generally enforce the Uniform Code. More restrictive local standards require state approval before effectiveness.

Official sources:

- `src:usa-ny:uniform-code`: https://dos.ny.gov/uniform-fire-prevention-and-building-code
- `src:usa-ny:2025-notice`: https://dos.ny.gov/notice-adoption
- `src:usa-ny:local-enforcement`: https://dos.ny.gov/code/local-government-state-agency-enforcement-programs
- `src:usa-ny:local-authority-memo`: https://dos.ny.gov/node/23866
- `src:usa-ny:nyc-codes`: https://www.nyc.gov/site/buildings/codes/codes.page
- `src:usa-ny:nyc-existing-code`: https://www.nyc.gov/site/buildings/codes/existing-building-code.page

Resolver impact: explicit geographic exceptions and provision-level legal-status overrides are required.

## California

Supported seed facts:

- The 2025 California Building Standards Code, Title 24, became effective on 2026-01-01 and is the current triennial edition.
- Title 24 contains multiple code parts and is produced through a multi-agency adoption and approval system coordinated by the California Building Standards Commission.
- Cities and counties may file qualifying local modifications based on required local findings.
- The Commission maintains a 2025 ordinance registry.
- Current law temporarily restricts many new local residential modifications from 2025-10-01 through 2031-06-01, subject to exceptions.

Official sources:

- `src:usa-ca:title-24-codes`: https://www.dgs.ca.gov/en/BSC/Codes
- `src:usa-ca:cbsc-about`: https://www.dgs.ca.gov/BSC/About
- `src:usa-ca:local-amendments`: https://www.dgs.ca.gov/bsc/codes/local-amendments-to-building-standards---ordinances
- `src:usa-ca:2025-ordinances`: https://www.dgs.ca.gov/BSC/Codes/2025-Ordinances
- `src:usa-ca:hsc-17958-5`: https://leginfo.legislature.ca.gov/faces/codes_displaySection.xhtml?lawCode=HSC&sectionNum=17958.5.
- `src:usa-ca:hsc-17958-7`: https://leginfo.legislature.ca.gov/faces/codes_displaySection.xhtml?lawCode=HSC&sectionNum=17958.7
- `src:usa-ca:2025-code-changes`: https://www.dgs.ca.gov/BSC/Resources/2025-Title-24-California-Code-Changes

Resolver impact: California needs occupancy and state-agency routing plus a separately refreshed local-amendment dataset.

## Cross-state requirements exposed

1. Geography-conditioned policies for named municipalities or maintained geographic policy layers.
2. Conditional future dates based on an official certification or event.
3. Provision-level suspension or conflict states.
4. Edition-specific local-amendment registries.
5. Multi-agency project and occupancy jurisdiction.
6. Independent effective, optional, mandatory, and replacement dates.

## Immediate next work

- Convert Virginia and Oregon into complete state reports first.
- Create a North Carolina draft that leaves the 2024 code conditionally pending.
- Design geography-conditioned policy selection before compiling Massachusetts, New York, or California into executable profiles.
