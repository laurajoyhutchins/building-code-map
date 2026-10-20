# BCM software and regulatory-data ownership

This document records the approved target architecture. It does **not** claim migration or production cutover is complete.

## Intended source of truth

- `laurajoyhutchins/building-code-map` (this **public clean-history repository**) owns the reusable software, canonical Authority Engine, project-code verification, HTTP contract `POST /v1/project-code-basis`, schemas, synthetic reference application, and eventually the full product frontend, geocoder algorithms, and GIS tooling.
- A **private regulatory-data repository** will own maintained jurisdiction research, source retrieval records, license- or redistribution-sensitive evidence, source-to-claim mappings, private verification fixtures, and the compilation of signed/content-identified regulatory bundles. Until a separate repository is created and independently validated, `building-code-map-alpha` remains the protected custody location for those bytes.
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
6. Switch website/API build and deployment to the public application, confirm health, rollback, and real project-scoped results. Preserve GCP-backed CI without falling back to usage-limited substantive GitHub runners.
7. Only after read-back of public source, private bundle origin, CI, deployed application, and rollback: archive alpha. Do not delete it or make its history public.

## Current transitional state

The public repository now has an executable **synthetic-only** project verification reference app. The private alpha application still owns its deployed/operational frontend and geocoder until the reviewed migration above is complete. Work on the public app and its schema first; use alpha only to preserve and transfer existing behavior and protect private regulatory data.

The public engine's project-basis v0.3 and `POST /v1/project-code-basis` are the intended canonical contracts. The two repo-specific Go engines and endpoint shapes are **not** compatible merely because both use the BCM name.
