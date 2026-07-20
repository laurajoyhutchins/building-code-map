# Contributing to Building Code Map

Thank you for helping improve Building Code Map. This project combines software, geographic data, and evidence-bearing building-code research, so contributions must preserve both technical correctness and source provenance.

## Before You Start

Use the issue templates to report:

- software defects;
- incorrect or incomplete jurisdiction or code-adoption data;
- narrowly scoped feature requests.

Report suspected security vulnerabilities privately through GitHub's **Report a vulnerability** feature. Do not publish credentials, private data, exploit details, or working proofs of concept in a public issue.

## Local Development

Prerequisites and launch commands are maintained in the [README](README.md). The canonical verification commands are:

```bash
pnpm install --frozen-lockfile
pnpm check
pnpm lint
pnpm format:check
pnpm test
pnpm build
pnpm backend:test
pnpm backend:lint
```

Runtime SQLite and DuckDB snapshots are intentionally not committed. A healthy API can report not-ready until a supported local snapshot is available.

## Contribution Boundaries

### Software changes

- Keep changes focused and consistent with the existing React/TypeScript and Go architecture.
- Add or update tests for behavior changes.
- Update API or configuration documentation when a public contract changes.
- Avoid introducing a new dependency when the standard library or an existing dependency is sufficient.

### Geographic and regulatory data

Every evidence-bearing change must identify:

- the affected jurisdiction or authority;
- the controlling official publisher;
- a stable official source URL;
- the access or verification date;
- the adoption, effective, operative, or mandatory date when applicable;
- unresolved or conflicting facts without guessing.

Do not copy commercial model-code text, standards, or copyrighted commentary merely because a jurisdiction adopts them. Record editions, adoption instruments, amendments, applicability, and authority using legally reusable facts and short necessary excerpts only.

### Restricted and generated material

Do not commit:

- credentials or private environment files;
- hydrated SQLite or DuckDB databases;
- department-owned NERIS operational, incident, personnel, or exposure data;
- private prompts or generated decision-graph exports containing private material;
- machine-specific home-directory paths;
- logs, caches, build outputs, or local tool state.

See [DATA_SOURCES.md](DATA_SOURCES.md) and [SECURITY.md](SECURITY.md) for the maintained publication boundaries.

## Pull Requests

A pull request should:

1. link the relevant GitHub issue;
2. explain the user-visible or architectural outcome;
3. describe frontend, backend, data, and documentation impact;
4. list the verification performed;
5. include screenshots for visible UI changes;
6. identify official sources for data or regulatory changes.

Keep commits focused and use concise imperative subjects. The preferred merge method is squash merge so `main` retains one intentional commit per reviewed change.

## Review Standard

Reviewers should reject changes that:

- weaken provenance or replace unresolved facts with unsupported certainty;
- add redistribution-sensitive data to Git history;
- describe deferred architecture as already implemented;
- bypass tests, formatting, or least-privilege workflow boundaries;
- mix unrelated cleanup with the proposed outcome.
