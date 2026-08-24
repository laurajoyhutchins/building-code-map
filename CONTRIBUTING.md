# Contributing

Contributions to the public Building Code Map project should improve public software, contracts, schemas, documentation, or synthetic examples.

Do not submit private project data, credentials, maintained production regulatory research, or bulk jurisdiction-by-jurisdiction corpus material through this repository.

Public fixtures must use fictional identifiers such as `DEMO-XX` unless a specific real-world example has been deliberately approved for public distribution. Tests should preserve the distinction between `verified`, `conditional`, and `not_verified`, and must not treat missing evidence as proof of non-applicability.

Changes should be small enough to review against the public/private boundary. When a contribution would require production data to demonstrate correctness, prefer a synthetic fixture or an interface-level test instead.

## Material from BCM alpha or third parties

Do not directly port files, data, generated artifacts, or history from the archived BCM alpha repository. Follow `docs/admission-policy.md`: classify the material, resolve authorship and licensing, choose its destination, and only then clean-copy or reimplement content that belongs in the public product.

Third-party source text, model codes, standards, production research, and real jurisdiction fixtures are not made public merely by inclusion in a contribution.

Before opening a pull request, run:

```sh
npm test
npm run validate
```
