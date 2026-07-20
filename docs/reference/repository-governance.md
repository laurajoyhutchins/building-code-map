# Repository Governance Reference

This document defines the intended GitHub settings and contribution controls for the public Building Code Map repository.

## Repository Settings

- Visibility: public
- Default branch: `main`
- Preferred merge method: squash merge
- Merge commits: disabled
- Rebase merge: disabled
- Automatic merge: disabled
- Automatically delete head branches after merge: enabled
- Allow pull-request branches to be updated: enabled

Squash merge keeps one intentional commit on `main` for each reviewed change. Automatic merge remains disabled so landing a change is an explicit action after exact-head verification.

## `main` Ruleset

Use an active ruleset named **Exact-head review clearance** targeting `main`.

Required rules:

- changes must enter through a pull request;
- required approving review count is `0` for the solo-maintainer repository;
- required status checks are `Frontend` and `Backend`;
- the pull-request head must be up to date with `main` before merge;
- all review conversations must be resolved;
- non-fast-forward pushes are blocked;
- branch deletion is blocked;
- linear history is required.

Do not require code-owner approval while the repository has one maintainer, because a pull-request author cannot approve their own change. `CODEOWNERS` remains useful for ownership visibility and automatic review requests from external contributors.

## Exact-Head Verification

Before landing a pull request:

1. confirm the reviewed head SHA is still the current pull-request head;
2. confirm `Frontend` and `Backend` passed for that head;
3. confirm review conversations are resolved;
4. squash merge using an outcome-focused commit title;
5. confirm the resulting commit is present on `main`;
6. confirm the source branch is deleted.

A prior successful run on an older head does not clear a newer head.

## Security Settings

Enable the GitHub features available for the public repository:

- private vulnerability reporting;
- dependency graph;
- Dependabot alerts;
- Dependabot security updates;
- secret scanning;
- push protection for detected secrets.

Version-update pull requests are configured separately in `.github/dependabot.yml` for npm, Go modules, and GitHub Actions.

## Public Contribution Controls

- `CONTRIBUTING.md` defines software, data, provenance, copyright, and restricted-data boundaries.
- `.github/CODEOWNERS` identifies the maintainer for all repository content and explicitly protects governance files.
- `.github/pull_request_template.md` requires impact, verification, and evidence details.
- `.github/ISSUE_TEMPLATE/` separates software bugs, data corrections, and feature requests.
- `SECURITY.md` directs vulnerability reports to private disclosure.
- `DATA_SOURCES.md` defines attribution and redistribution boundaries.

## Review Cadence

Review this policy after any of the following:

- another maintainer receives write access;
- CI job names change;
- the repository begins publishing releases or deployments;
- a new data source introduces different redistribution terms;
- a security incident or provenance failure exposes a missing control.
