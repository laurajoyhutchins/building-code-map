# Pilot production readiness

Issue #31 promotes the Colorado, Florida, and New Jersey pilots only after a defined, testable scope has complete evidence. The ordinary regulatory coverage report remains useful for nationwide inventory, but placeholder fixture identifiers are not evidence and cannot establish production readiness.

## Readiness manifest

Each pilot has a manifest in `reports/production/` that declares:

- the code families and applicability-date window being promoted;
- the source records treated as official primary evidence;
- continuous current and historical adoption timeline segments;
- the resolver fixture classes required for the pilot;
- named fixture definitions, expected resolver status, and supporting state and local records.

The manifest does not alter runtime resolution. It is an acceptance contract over the compiled state profile and rule pack.

## Promotion gate

Run:

```bash
python reports/tools/production_readiness.py audit \
  --compiled-root backend/data/regulatory \
  --manifests-root reports/production \
  --json-output /tmp/pilot-readiness.json \
  --markdown-output /tmp/pilot-readiness.md \
  --fail-on-premature-verification
```

A state is `production_ready` only when all of the following hold:

1. The profile and profile verification are `verified`.
2. Every scoped code family has at least one adoption record and a continuous verified timeline across the entire declared date window.
3. Timeline adoption references and material supported claims are verified and cite only the manifest's primary-source set.
4. Every required fixture class is represented by a fully defined, registered, verified fixture.
5. Incorporated, unincorporated, and production-supported historical fixtures cite official local records, not only statewide summaries.
6. Negative fixtures deliberately verify missing-evidence and conflict behavior without being included in the supported production result set.
7. Every primary source has a current `available` health record.
8. No supported claim remains conflicting or partially verified.

CI fails when a profile is changed to `verified` before this gate is satisfied. The initial manifests intentionally expose the remaining research gaps and do not promote any pilot.

## Source health

The scheduled `Regulatory source health` workflow checks every official URL in the compiled pilot profiles. It records HTTP status, redirects, the final URL, and availability in JSON and Markdown artifacts. A redirect is classified as `moved` so the canonical source URL can be updated instead of silently relying on redirect behavior.

The source-health check is network-dependent and therefore runs on a schedule and by manual dispatch rather than as part of every pull request.
