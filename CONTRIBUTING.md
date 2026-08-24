# Contributing

Contributions to Building Code Map should improve the public software, contracts, schemas, documentation, or synthetic examples.

## Public contribution boundary

Do not submit private project data, credentials, maintained production regulatory research, bulk jurisdiction-by-jurisdiction corpus material, or restricted source text through this repository.

Public fixtures should use fictional identifiers such as `DEMO-XX` unless a real-world example has been deliberately reviewed and approved for public distribution. Tests should preserve the distinction between `verified`, `conditional`, and `not_verified`, and must not treat missing evidence as proof of non-applicability.

When production data is unnecessary to demonstrate behavior, prefer a synthetic fixture or interface-level test.

## Third-party material

Before adding externally sourced material, confirm authorship, license terms, attribution requirements, and redistribution rights. Public availability of a law, dataset, model code, standard, or government record does not by itself establish that every representation or copied byte may be redistributed here.

See [`docs/publication-policy.md`](docs/publication-policy.md) for the repository's publication rules.

## Verification

Before opening a pull request, run:

```sh
npm test
npm run validate
```

Keep changes focused enough that reviewers can evaluate both behavior and the public distribution boundary.
