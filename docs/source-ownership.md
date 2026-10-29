# BCM software and regulatory-data ownership

This document records the approved target architecture. It does **not** claim migration or production cutover is complete.

## Functional ownership: one complete task per product

The product boundary is determined by the **observable question and independently
verified answer**, not by implementation language, pipeline stages, data format,
or where credentials and source bytes are stored. A public/private pair may be
necessary for data custody without becoming two competing software products.

### BCM's one job

**Input:** a project location, legally relevant applicability date and known
project facts. **Output:** a reproducible project-scoped code basis: candidate
governing authorities, adopted editions and amendments, evidence, verification
state, and unresolved project or source requirements.

BCM owns the algorithms and user experience required to produce that answer:
address and coordinate resolution, jurisdiction-boundary observation, dated
adoption policy, legal-evidence acceptance, one Authority Engine and
`POST /v1/project-code-basis`. The answer must remain explicitly conditional
or unverified when the controlling records, effective dates, jurisdiction or
project facts are insufficient.

BCM **does not** own generic PDF extraction, publication structure packaging,
source-independent rule-to-Rust compilation, engineering calculations, or
electric-utility candidate lookup. It may reference their versioned outputs,
but neither a source package nor a utility candidate can promote a BCM verdict.
Geocoding and GIS remain part of BCM while needed for code applicability; they
are not automatically separate products.

### Neighboring completed tasks and interfaces

| Tool / current implementation | Complete task and result | Boundary presented to BCM |
| --- | --- | --- |
| [BCAST package toolkit](https://github.com/laurajoyhutchins/regulatory-package-toolkit) | Reviewed, normalized publication records -> deterministic content-addressed publication package that an independent client can validate and navigate | Verified package/object identities are *references* only, not evidence of legal adoption |
| [Arcata Regulatory Source Operations](https://github.com/laurajoyhutchins/arcata-rso) | Lawfully acquired and reviewed legal source -> source-audited BCAST package, exact text/enactment/amendment evidence, provenance and regulatory proof frontier | Reviewed claim evidence is referenced by BCM; source review does not itself authorize the BCM verdict or publicize protected bytes |
| `bcast-rules` (currently a Rust crate inside the BCAST toolkit) | Reviewed, edition-pinned rule IR -> standalone Rust predicates and conformance results | Optional separately reviewed rules, never authority to choose a jurisdiction or code edition |
| [Electric Utility Catalog](https://github.com/laurajoyhutchins/electric-utility-catalog) | EIA release -> verified, source-identifiable statistical utility catalog and county-candidate lookup | **No regulatory-decision dependency**; this is a separate product |
| [Obsidian PDF Extractor](https://github.com/laurajoyhutchins/obsidian-pdf-extractor) | User-selected PDF -> previewed, validated structured notes and traceable extraction artifacts | Generic extraction is not authoritative regulatory interpretation |

The Rust rule compiler's output and lifecycle are distinguishable from package
navigation. Do **not** create another repository solely because it is a crate:
a separate release/repository is warranted only after independent users,
a stable input/output contract, and an end-to-end real-source demonstration
establish a separately useful tool. Until then BCAST owns the existing compiler
and must not duplicate it in BCM.

### Transitional or misplaced responsibilities

- `building-code-map-alpha` is a temporary custodian of richer existing
  application behavior and protected research, **not** another product to
  develop indefinitely. Migrate only reviewed, independently verifiable
  functionality and archive alpha after parity, provenance and rollback.
- `arcata-jurisdictions` (the renamed private geographic-data repository)
  owns publisher-native governmental identities, canonical identity/PROV-O
  contracts, source-bound boundary observations and protected source custody.
  Its source-only canonical exports are *candidates* admitted by BCM's
  `ReadJurisdictionIdentityCandidates` consumer; neither a Census match nor
  a polygon observation proves a legal boundary, enforcing AHJ, adoption or
  project verdict.
- A thin private evidence/provider assembler may eventually reference
  `arcata-rso` source-reviewed claim evidence and `arcata-jurisdictions`
  candidate geography. No independently accepted **production provider
  assembly or live source-to-project cutover** is claimed. The previous
  `bcm-regulatory-data` repository name does not create another legal
  source-review engine or project verifier.
- BCM currently has separate EIA county utility-context APIs and an
  `bcm-utility-api` executable. These are **out of BCM's functional scope**,
  despite being deliberately isolated from the code verdict. Do not expand
  them as a second BCM product. Preserve any existing consumer and provenance
  while planning ownership of the utility-facing API with Electric Utility
  Catalog (or a proven independent consumer), then remove duplicate BCM
  serving/ETL after verified equivalence. No speculative new repo.
- Archived `building-code-dashboard` and `building-code-ast-alpha` are
  historical references, not independent active product authorities.

### Required proof before treating a boundary as operational

1. A real or explicitly labeled synthetic input runs to the named output
   through **one canonical implementation**, with source/version identity.
2. Another implementation, reviewer or consumer independently checks the
   output against its inputs; internal self-consistency is not enough.
3. Incompatible, unavailable, ambiguous, expired or unreviewed evidence fails
   closed rather than promoting a result.
4. A consumer uses a documented, versioned artifact/API; source-tree sharing,
   runtime imports of private modules and coordinated multi-repo source edits
   are not ordinary prerequisites.
5. No migration archives an existing working path until exact-head verification,
   authentic source/data readback, deployed equivalence where applicable, and
   rollback are independently observed.

Today BCM has a synthetic Go/API/UI workflow, but its authentic Denver
project-level proof is still an open acceptance task
([BCM #74](https://github.com/laurajoyhutchins/building-code-map/issues/74)).
A BCAST eCFR package has independent consumer/readback evidence
([BCAST #22](https://github.com/laurajoyhutchins/regulatory-package-toolkit/issues/22)),
with explicit fidelity and release-rights limitations. A structural package
does not constitute a verified code interpretation. Neither proof substitutes
for BCM's real legal applicability decision.

## Regulatory evidence versus geographic authority cutover

The [cross-repo evidence and geography cutover plan](regulatory-evidence-geography-cutover.md)
makes `arcata-rso` (Regulatory Source Operations) responsible for legal
source custody, source-text/adoption/amendment/effective-date review and
regulatory proof obligations. Private `arcata-jurisdictions` produces
source-native governmental/territorial observations with unresolved conflicts,
not legal authorization. Public BCM owns the geocoder/GIS algorithms,
independent government identity and dated authority admission, and the sole
`Runtime`/`ProjectVerifier` verdict. The public candidate-only consumer
and separate trusted runtime `SourcePolicy` and `SiteAuthorityPolicy` are
already implemented with synthetic negatives; **no real source-approved RSO
claim plus site-authority pair is accepted**. See [BCM #91](https://github.com/laurajoyhutchins/building-code-map/issues/91),
[site authority #90](https://github.com/laurajoyhutchins/building-code-map/issues/90)
and [RSO #77](https://github.com/laurajoyhutchins/arcata-rso/issues/77).
The older source-ownership text below is an uncompleted migration baseline and must be read together with this approved target split. Do not treat this documentation as permission to archive alpha or as evidence of an accepted real regulatory claim.

## Intended source of truth

- `laurajoyhutchins/building-code-map` (this **public clean-history repository**) owns the reusable software, canonical Authority Engine, project-code verification, HTTP contract `POST /v1/project-code-basis`, schemas, synthetic reference application, and eventually the full product frontend, geocoder algorithms, and GIS tooling.
- Private `arcata-rso` is responsible for regulatory originals, reviewer and
  rights custody, bounded legal evidence and source-safe *candidate* claim
  exports. Private `arcata-jurisdictions` owns source-only governmental identity
  and geographic observation contracts. A thin protected provider adapter
  may eventually combine **approved** references and dated site evidence, but
  is neither a legal reviewer nor a second verdict path. Until individual
  original-source migrations and independent consumer readback pass,
  `building-code-map-alpha` remains custodian of unmigrated originals.
- `building-code-map-alpha` is **not** a second long-term software authority. It remains active temporarily for existing production capabilities and in-flight PRs, then becomes a read-only historical archive after complete verified cutover.

## Non-negotiable boundaries

The public repository must never receive a commit, branch, tag, tree, or Git object copied wholesale from private alpha history. New public source changes must be independently reviewed and committed with **only deliberately releasable bytes**. A Git history rewrite cannot be treated as an adequate substitute for review.

Public software does not contain the maintained corpus; it consumes only a typed, versioned provider interface and a verified, separately distributed data bundle or hosted API. The public synthetic app accepts only a fictional address and reports its lack of real-world authority. Its presence is **not** a declaration of production availability.

Only the public `engine.ProjectVerifier` is the target owner of the project verdict and of project-code-basis serialization. Compatibility adapters may translate private regulator observations into the public Provider contract, but must not independently generate verdicts, build a parallel project code set, or silently substitute missing evidence.

## Exact cutover requirements

1. Preserve and independently verify private repo/corpus backup, access control, object identities, and existing deployed availability.
2. Resolve PR backlog on alpha without blindly transporting generated data or private CI artifacts into public Git history.
3. Reconcile and version the public `0.3` project-basis contract; the old private exported `0.1` example must no longer be treated as authoritative, and the proposed private `/v1/project/verify` endpoint must become a compatibility adapter or be retired in favor of `/v1/project-code-basis`.
4. Move reviewed algorithmic code and frontend behavior to the clean public repository, using synthetic fixtures, open-source dependency licenses, public boundary checks, and full frontend/backend tests. Connect private data through the Provider boundary; keep regulatory materialization within private custody.
5. Verify the same project examples and failure cases against both the public synthetic provider and the private data adapter. Assert fail-closed behavior, temporal identity, evidence provenance, and no bulk-source leakage.
6. Switch website/API build and deployment to the public application, confirm health, rollback, and real project-scoped results. Preserve exact-head verification using repository-appropriate CI capacity; do not infer runtime readiness from a successful source build.
7. Only after read-back of public source, private bundle origin, CI, deployed application, and rollback: archive alpha. Do not delete it or make its history public.

## Current transitional state

The public repository now has an executable **synthetic-only** project verification reference app. The private alpha application still owns its deployed/operational frontend and geocoder until the reviewed migration above is complete. Work on the public app and its schema first; use alpha only to preserve and transfer existing behavior and protect private regulatory data.

The public engine's project-basis v0.3 and `POST /v1/project-code-basis` are the intended canonical contracts. The two repo-specific Go engines and endpoint shapes are **not** compatible merely because both use the BCM name.
