# Data Visibility Contract

Building Code Map separates an inspectable public software product from the maintained operational data used to support regulatory answers.

The boundary is intentional. A caller can receive a useful, evidence-backed project answer without receiving a downloadable copy of every underlying source, research record, or maintained jurisdiction dataset.

## Public repository

The repository may contain:

- software, schemas, and interface contracts approved for public release;
- methodology and product documentation;
- synthetic fixtures and invented evidence used for tests;
- citations, identifiers, or examples that are independently safe to publish.

The repository does not contain the maintained jurisdiction-by-jurisdiction production corpus, private project data, restricted source captures, or bulk evidence mappings.

## Hosted service responses

A project-scoped response may expose the information needed to understand and evaluate that result, including:

- normalized regulatory conclusions and their `verified`, `conditional`, or `not_verified` status;
- identified authority and jurisdiction;
- applicable code family or rule, edition, and effective-date information;
- unresolved project conditions;
- citations, source identifiers, URLs, and provenance metadata that are safe to expose.

A hosted response is not a bulk-corpus export. The public contract does not promise raw source captures, research notes, internal verification annotations, complete jurisdiction histories, or unrestricted enumeration of the maintained dataset.

Restricted source text is not part of the default response contract. Quotation or redistribution of third-party text requires a source-specific basis independent of this policy.

## Non-public operational data

Operational systems may contain material necessary to maintain and verify regulatory answers, including jurisdiction records, effective-date histories, research notes, evidence captures, verification annotations, transformation artifacts, license metadata, and source material that cannot be redistributed publicly.

Operational availability does not imply public redistribution rights.

## Derived facts and source material

A fact derived from a source and the source bytes themselves are different artifacts and are classified separately.

Building Code Map may return a project-scoped derived regulatory fact when that disclosure is permitted by the product contract. That does not imply that the underlying source text, the complete maintained dataset, or intermediate research is distributable.

Likewise, a public law or government record being publicly accessible does not automatically settle whether every retrieved representation, compilation, annotation, or third-party hosting copy may be redistributed.

## Default for new interfaces

New schemas, endpoints, exports, examples, and debugging surfaces must fail closed against this boundary:

1. expose the information needed for the product contract and no accidental bulk surface;
2. prefer citations and derived facts over redistributing source bytes;
3. do not add bulk-corpus access by incidental convenience;
4. do not make operational fields public merely because they already exist internally;
5. treat every new export surface as a visibility decision, not ordinary serialization.

If a feature needs broader visibility, change this contract deliberately and review the licensing and product consequences at the same time.
