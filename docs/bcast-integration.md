# BCAST content references in Building Code Map

BCAST and Building Code Map answer different questions. [BCAST `bcast.package/0.2.0`](https://github.com/laurajoyhutchins/bcast/blob/main/spec/package-0.2.0.md) identifies an exact materialization of regulatory publication content. BCM determines whether a requirement applies to a **particular project, jurisdiction and date** from adoption and amendment evidence.

## Consumer integration using existing BCM contracts

1. Obtain a BCAST package through its public API or client, configured for `bcast.package/0.2.0`. Verify the **full package**, including its content-addressed `package_id`, before trusting an identity. API metadata alone cannot verify the digest.
2. Select provisions by public `object_id` / `(kind, locator)`, not by source-provider coordinates. The BCAST `status` of `resolved`, `unresolved` or `unsupported` means representation availability, **not** applicability or legal correctness.
3. When BCM's provider has independently resolved a project-scoped code decision, it may carry source-trace references in the existing `CodeResolution.Evidence` array, for example `{"kind":"bcast.package","id":"bcastpkg:sha256:<64-hex-digest>"}` and `{"kind":"bcast.object","id":"bcastobj:sha256:<64-hex-digest>"}`. They are **trace pointers only**, not evidence-link roles or a substitute for the original source.
4. Adoption, amendment, effective-date and AHJ claims must still use BCM's existing `EvidenceClaim`, `SourceDocument`, `SourceArtifact`, `TextAnchor` and `SourcePolicy` evidence path. An establishing link to an authoritative adoption record must be independently supported. Neither a BCAST hash nor source-citation metadata implies that a jurisdiction adopted a code.

No new field, schema version, adapter framework, or parallel jurisdiction resolver is needed just to retain BCAST source coordinates. BCM's existing `project-code-basis/0.3` `evidence` representation already carries typed references.

## Deliberate fail-closed behavior

A BCAST content reference cannot change `StatusInsufficientEvidence` to `StatusResolved`, produce a `ProjectVerified` verdict, or be copied into `ExactEvidence` as though it were an operative ordinance. A BCAST package with `status: resolved` is not a BCM `verified` project-code basis. These are independently named states with different meanings.

For example, 24 CFR Part 3280 from the eCFR is a federal publication; its presence in BCAST does **not** demonstrate local building-code adoption. The missing jurisdictional evidence remains explicit.

## Downstream Arcata boundary

Once BCM establishes the governing adopted code and edition, Arcata may use that result to select an explicitly supported engineering calculation profile. The BCAST package/object identities may be retained as input-source references, but **must not** reconfigure Arcata's physical `ElectricalSystem`, certify a method, or silently replace the engineer's edition choice. See [Arcata standards and references](https://github.com/laurajoyhutchins/arcata/blob/main/docs/standards-and-references.md).

This documents a tested source-reference handoff, not a production integration with a live BCAST service. Production BCAST API deployment and BCM authority-data sourcing remain separate work.
