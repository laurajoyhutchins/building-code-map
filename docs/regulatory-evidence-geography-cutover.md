# Regulatory evidence / geography cutover

Status: **approved ownership direction; not a production cutover** (2026-10-28). Implementation tracking: [BCM #41](https://github.com/laurajoyhutchins/building-code-map/issues/41), [BCM #91](https://github.com/laurajoyhutchins/building-code-map/issues/91), [site attribution #90](https://github.com/laurajoyhutchins/building-code-map/issues/90), [Denver project #74](https://github.com/laurajoyhutchins/building-code-map/issues/74).

## One project answer, two independently established facts

BCM's product question is: for this project location, date, code family and known facts, which governing instruments and authorities can actually be verified? BCM exclusively owns geospatial data processing and jurisdiction determination, plus the canonical Runtime, ProjectVerifier and project-code-basis HTTP path.

Regulatory Source Operations (RSO) owns the evidence **why** a regulation claim is supported: acquisition, original source identity and content, exact text, publication and enactment/exception/amendment interpretation, reviewed effective intervals, provenance and missing regulatory proof obligations. The public [BCAST](https://github.com/laurajoyhutchins/regulatory-package-toolkit) package is an immutable publication representation, not evidence that it became law.

**Neither an official legal source nor a polygon by itself establishes a project verdict.** RSO may say that ordinance X establishes adoption Y for a legally identified municipality and interval, subject to reviewed exceptions. BCM independently verifies that a parcel/point is governed by the relevant territorial and enforcement authority on the project date.

## Contract and authority diagram

1. Publisher/rightfully obtained regulatory source -> RSO extraction, audited source review, claim-obligation closure, bounded adoption/amendment/effective-date evidence.
2. RSO produces rights-scoped, immutable reviewed **claim-evidence** records and W3C PROV-O source lineage, conforming to a jointly versioned **JSON Schema 2020-12 + ODCS** export.
3. Private BCM data/provider assembly references RSO evidence and its generation. It independently binds government entity IDs, date, family, accepted geography and approved evidence-policy inputs. Provider may not certify itself.
4. BCM Runtime verifies claims against a **separately trusted, runtime-owned SourcePolicy**, including exact document/bytes/anchor, date, jurisdiction, contradiction and supersession. RSO's reviewer field and source-policy proposal cannot authorize this step.
5. BCM separately verifies site-to-jurisdiction authority from geocoder precision, dated geometry, legal boundary/delegation instrument and government entity identity. Geographic evidence cannot be inferred from Census name matches, statewide code inheritance, or arbitrary provider-assigned jurisdiction IDs.
6. The existing ProjectVerifier reduces the canonical Runtime result to verified, conditional or not_verified with structured unresolved requirements. No alternative endpoint, graph truth store or verdict reducer.

**Time dimensions:** source acquisition time, source publication/edition, enactment/adoption date, effective/repeal interval, boundary/annexation effective interval, and project as-of date must not be collapsed. **Spatial dimensions:** candidate polygon containment, legally established boundary, enforcing AHJ and project parcel/point must not be collapsed.

## Private/public ownership

| Responsibility | Target owner |
| --- | --- |
| Publication/section identity and neutral package semantics | public BCAST |
| Regulatory originals, rights and review receipts; interpretation and regulatory claim obligations | private RSO |
| Geocoding, GIS, Census crosswalks, governmental/territorial authority; spatial and effective-date algorithms | public BCM software |
| Protected geographic snapshot custody, thin RSO evidence reference/bundle assembly, provider distribution | private bcm-regulatory-data |
| Temporary state research and existing live app, until verified transfer | private building-code-map-alpha |

Alpha remains source custodian until source rights, retained hashes, reviewer evidence, contract binding, real project consumption and rollback are verified. Never import alpha's private Git history into this public repository.

## Acceptance before authority transfer

- A realistic municipal enactment is independently reviewed at original-source text/edition and effective-date scope; the expected reviewed claim is not derived only from the output under test.
- Cross-repo export has paired schema + ODCS; the public W3C PROV graph importer validates reference closure. Any in-flight private state graph using a legacy dialect must migrate before admission.
- Explicit negatives for a fabricated review activity, wrong identity/family/date, altered artifact or anchor digest, unreviewed exception, repeal, contrary ordinance, two enforcing authorities, ambiguous locality and overlapping polygons.
- A real address or parcel has independently evidenced dated legal territorial authority; an interpolated point or source snapshot alone never proves it.
- Existing SourcePolicy stays independently supplied. Without approved exact source witnesses, runtime returns insufficient/not_verified regardless of provider status.
- One project passes the **existing** Runtime/ProjectVerifier/POST /v1/project-code-basis. Missing project facts yield conditional follow-up; source defects remain evidence_defect rather than user questions.
- Source-safe receipts pin the exact revision/export/consumer; unknown statuses survive and the original provider remains recoverable until equivalence and deployed readiness are observed.

## State scale and migration

Keep **identity inventoried**, **official source observed**, **regulatory source independently reviewed**, **legal claim admissible**, **territorial authority verified**, and **project basis verified** as separate metrics by family and applicability date.

One statewide source may satisfy many shared legal obligations but cannot prove thousands of local permitting offices or parcel boundaries. Kansas provides the first candidate locality pilot, Michigan tests statewide inheritance, Tennessee preserves multi-county conflicts, Maine tests source custody, Illinois/ND test provenance and spatial identity joins. Treat those as test cases, not already completed legal verification.

Workstream tasks: [RSO claim export #77](https://github.com/laurajoyhutchins/regulatory-source-operations/issues/77), [RSO proof closure #78](https://github.com/laurajoyhutchins/regulatory-source-operations/issues/78), [private custody #8](https://github.com/laurajoyhutchins/bcm-regulatory-data/issues/8), [alpha inventory #318](https://github.com/laurajoyhutchins/building-code-map-alpha/issues/318). No dependency on Overcenter or duplicated runtime in RSO. All new cross-repo contracts must satisfy JSON Schema + ODCS; internal contracts need JSON Schema.
