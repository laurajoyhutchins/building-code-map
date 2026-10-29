# Regulatory evidence / geography cutover

Status: **approved ownership direction; not a production cutover** (2026-10-28). Implementation tracking: [BCM #41](https://github.com/laurajoyhutchins/building-code-map/issues/41), [BCM #91](https://github.com/laurajoyhutchins/building-code-map/issues/91), [site attribution #90](https://github.com/laurajoyhutchins/building-code-map/issues/90), [Denver project #74](https://github.com/laurajoyhutchins/building-code-map/issues/74).

## Implemented source boundary versus uncompleted legal cutover

The private `arcata-rso` repository contains a **source-safe candidate** claim
exporter with paired JSON Schema/ODCS, W3C PROV-O source entity records and a
missing-regulatory-proof frontier. It records reviewer references as **unadmitted**:
a synthetic export is not a legally reviewed ordinance or a trusted BCM policy.
The private `arcata-jurisdictions` producer emits source-bound governmental and
geographic identity **candidates** under canonical JSON Schema/ODCS and PROV-O
contracts; BCM's `ReadJurisdictionIdentityCandidates` is an offline
candidate-only consumer, not a project authority determination.

Public BCM's `Runtime` and `ProjectVerifier` have two separately configured
trust gates: `SourcePolicy` for code-adoption evidence and
`SiteAuthorityPolicy` for original-source-backed, dated location and enforcing
authority. Both are synthetic-testable; neither an accepted real municipal legal
claim nor a reviewed matching real parcel authority is established. See
[site authority](site-jurisdiction-authority.md) and
[candidate geography](candidate-geography-consumer.md). No live provider
cutover or alpha retirement is implied.

## One project answer, two independently established facts

BCM's product question is: for this project location, date, code family and known facts, which governing instruments and authorities can actually be verified? BCM exclusively owns geospatial data processing and jurisdiction determination, plus the canonical Runtime, ProjectVerifier and project-code-basis HTTP path.

Regulatory Source Operations (RSO) owns the evidence **why** a regulation claim is supported: acquisition, original source identity and content, exact text, publication and enactment/exception/amendment interpretation, reviewed effective intervals, provenance and missing regulatory proof obligations. The public [BCAST](https://github.com/laurajoyhutchins/regulatory-package-toolkit) package is an immutable publication representation, not evidence that it became law.

**Neither an official legal source nor a polygon by itself establishes a project verdict.** RSO may say that ordinance X establishes adoption Y for a legally identified municipality and interval, subject to reviewed exceptions. BCM independently verifies that a parcel/point is governed by the relevant territorial and enforcement authority on the project date.

## Contract and authority diagram

1. Publisher/rightfully obtained regulatory source -> RSO extraction, audited source review, claim-obligation closure, bounded adoption/amendment/effective-date evidence.
2. RSO currently implements deterministic **candidate** claim-evidence records, a missing-proof frontier and source-safe W3C PROV-O provenance under **JSON Schema 2020-12 + ODCS**; only separately authenticated original-source/reviewer receipts and legal adjudication may make a reviewed regulatory claim.
3. The proposed thin private provider adapter may bind independently reviewed RSO references, source-native `arcata-jurisdictions` governmental/geometry observations, dates and family into the public Provider contract. Its operational ownership and live delivery remain **unverified**. Provider output may not certify itself.
4. BCM Runtime evaluates claims against a **separately trusted, runtime-owned SourcePolicy** for exact document/bytes/anchor, date, jurisdiction and conflicting evidence. RSO's reviewer field and source-policy proposal cannot authorize this step.
5. BCM separately verifies site-to-jurisdiction authority through its independently trusted `SiteAuthorityPolicy`, requiring a source-pinned project point/parcel, original legal boundary/delegation instrument, independently reviewed dated topology and enforcing entity identity. A candidate polygon, Census name match, statewide code inheritance or provider-assigned ID does not satisfy this gate.
6. The existing ProjectVerifier reduces the canonical Runtime result to verified, conditional or not_verified with structured unresolved requirements. No alternative endpoint, graph truth store or verdict reducer.

**Time dimensions:** source acquisition time, source publication/edition, enactment/adoption date, effective/repeal interval, boundary/annexation effective interval, and project as-of date must not be collapsed. **Spatial dimensions:** candidate polygon containment, legally established boundary, enforcing AHJ and project parcel/point must not be collapsed.

## Private/public ownership

| Responsibility | Target owner |
| --- | --- |
| Publication/section identity and neutral package semantics | public BCAST |
| Regulatory originals, rights and review receipts; interpretation and regulatory claim obligations | private RSO |
| Geocoding/GIS algorithms, independently trusted dated site and enforcing-authority admission, Runtime and sole project verdict | public BCM software |
| Publisher-native governmental identity, source-only geographic observations, canonical schema/PROV and protected source custody | private `arcata-jurisdictions` |
| Proposed thin private Provider adapter combining approved immutable regulatory evidence and geography | **Not yet an independently accepted/deployed owner** |
| Temporary state research and existing live app, until verified transfer | private building-code-map-alpha |

Alpha remains custodian of unmigrated original material until source rights, retained hashes, reviewer evidence, contract binding, real project consumption and rollback are verified. Never import alpha's private Git history into this public repository.

## Acceptance before authority transfer

- A realistic municipal enactment is independently reviewed at original-source text/edition and effective-date scope; the expected reviewed claim is not derived only from the output under test.
- Cross-repo export has paired schema + ODCS; the public W3C PROV graph importer validates reference closure. Any in-flight private state graph using a legacy dialect must migrate before admission.
- Explicit negatives for a fabricated review activity, wrong identity/family/date, altered artifact or anchor digest, unreviewed exception, repeal, contrary ordinance, two enforcing authorities, ambiguous locality and overlapping polygons.
- A real address or parcel has independently evidenced dated legal territorial authority; an interpolated point or source snapshot alone never proves it.
- Both `SourcePolicy` and `SiteAuthorityPolicy` remain independently supplied and exact-source approved. Without either approved legal source witnesses **or** real site authority, the runtime remains `not_verified` regardless of provider assertions.
- One project passes the **existing** Runtime/ProjectVerifier/POST /v1/project-code-basis. Missing project facts yield conditional follow-up; source defects remain evidence_defect rather than user questions.
- Source-safe receipts pin the exact revision/export/consumer; unknown statuses survive and the original provider remains recoverable until equivalence and deployed readiness are observed.

## State scale and migration

Keep **identity inventoried**, **official source observed**, **regulatory source independently reviewed**, **legal claim admissible**, **territorial authority verified**, and **project basis verified** as separate metrics by family and applicability date.

One statewide source may satisfy many shared legal obligations but cannot prove thousands of local permitting offices or parcel boundaries. Kansas provides the first candidate locality pilot, Michigan tests statewide inheritance, Tennessee preserves multi-county conflicts, Maine tests source custody, Illinois/ND test provenance and spatial identity joins. Treat those as test cases, not already completed legal verification.

Workstream tasks: [RSO claim export #77](https://github.com/laurajoyhutchins/arcata-rso/issues/77), [RSO proof closure #78](https://github.com/laurajoyhutchins/arcata-rso/issues/78), [jurisdiction producer/migration #20](https://github.com/laurajoyhutchins/arcata-jurisdictions/issues/20), [alpha inventory #318](https://github.com/laurajoyhutchins/building-code-map-alpha/issues/318). No dependency on Overcenter or duplicated runtime in RSO. All new cross-repo contracts must satisfy JSON Schema + ODCS; internal contracts need JSON Schema.
