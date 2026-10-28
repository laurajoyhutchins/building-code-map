# Site-to-jurisdiction legal authority boundary

BCM #90 adds a **second, independent source-evidence dimension** to the
existing `engine.Runtime` and the same `engine.ProjectVerifier`. An adopted
code may be supported by exact legal text while the geographical site,
enforcing body, legal incorporation/annexation or historical date remains
unverified. It is not a verified project until **both** claims qualify.

The remote Provider remains responsible for candidate observations. It
cannot approve the site where its own code claim applies. The independently
configured `SiteAuthorityPolicy` is an operator-reviewed protected artifact,
separate from `SourcePolicy`. The public code has no accepted real Colorado
rules, geometry, original ordinances or reviewer authority.

## Policy input

The paired cross-repository definitions are
`schemas/site-authority-policy.schema.json` and
`schemas/site-authority-policy.odcs.json` (ODCS v3.1.0). Private data/RSO must
not bypass them. Each reviewed rule binds:

- the exact normalized project address/point via `SiteLocationSHA256`, with
  an independently reviewed address point or parcel (not interpolation);
- a governmental-entity ID, enforcing authority and closed legal effective
  interval, including source-bound official adoption or delegation identity;
- separately reviewed site-interior membership in generation-pinned GCS
  geometry, an explicit topology receipt and Census/source disagreement review;
- original legal instrument and exact operative-text digests and locator;
- a W3C PROV Entity/Activity/Agent and immutable independent-review receipt.

Every digest and PROV reference must correspond to **actual independently
retained and reviewed bytes**. Metadata and a self-asserted `reviewed` label
do not constitute those bytes. The runtime checks structure, matching identities,
dates, contradictory active site authorities and provenance references; it does
**not** independently fetch GCS or authenticate reviewers during a project
request. These upstream trust gates remain necessary before any real policy
can be configured.

Configure both `BCM_SITE_AUTHORITY_POLICY_PATH` and
`BCM_SITE_AUTHORITY_POLICY_SHA256` on the authenticated backend. The loader
requires exact SHA-256, bounded bytes, strict fields and unique JSON keys,
unique rule IDs and structurally complete review observations. The provider
cannot specify either variable. If no policy is configured, resolved code
observations may retain their exact adoption evidence but the canonical
project result becomes `not_verified` with an explicit
`site-jurisdiction-authority` evidence defect, not a prompt to the user.

## Project outputs

The optional `site_authority` output carries a narrow, safe summary:
`status` (reviewed, unresolved, conflicting), `reason`, and for a reviewed
witness its stable rule/provenance entity, review activity, receipt SHA-256
and geometry generation. It intentionally returns no private source body or
bulk geometry. A reviewed geographic witness never independently grants
`legal_verdict_authorized`. The existing exact-code `SourcePolicy` gate
and `ProjectVerifier` still determine the final verdict.

An input at a disputed polygon edge, a Hudson identity collision, two
potential enforcing authorities, an interpolated geocode or a late annexation
must not receive a verified verdict. Tests exercise these rejection paths
with **fictional** witnesses. No positive real Denver project case has yet
been accepted, because no independently reviewed dated legal site witness
and separately source-approved RSO adoption have been demonstrated together.
A synthetic successful test must not be presented as that real result.

The original Colorado candidate projection in `arcata-jurisdictions` and
the alpha custody ledger remain intact and explicitly unpromoted.
