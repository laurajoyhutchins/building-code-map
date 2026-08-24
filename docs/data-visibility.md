# Data Visibility Contract

Building Code Map separates an inspectable public software product from a maintained production regulatory knowledge base.

The boundary is intentional. A caller may receive a useful, evidence-backed regulatory answer without receiving a downloadable copy of the corpus used to produce it.

## Visibility levels

### Public repository

The repository may contain:

- software, schemas, and interface contracts approved for public release;
- methodology and product documentation;
- synthetic fixtures and invented evidence used for tests;
- source citations, identifiers, or examples that are independently safe to publish.

The repository does not contain the maintained jurisdiction-by-jurisdiction production corpus, production research notes, licensed source captures, private project data, or bulk evidence mappings.

### Hosted service response

A project-scoped response may expose the information needed to understand and evaluate that result, including:

- normalized regulatory conclusions and their `verified`, `conditional`, or `not_verified` status;
- identified authority and jurisdiction;
- applicable code family or rule, edition, and effective-date information;
- unresolved project conditions;
- citations, source identifiers, URLs, and provenance metadata that are safe to expose.

A service response is not a bulk-corpus export. The public contract does not promise raw source captures, research notes, internal verification annotations, complete jurisdiction histories, or unrestricted enumeration of the maintained dataset.

Restricted source text is not part of the default response contract. Any quotation or redistribution of third-party text requires a source-specific basis independent of this policy.

### Production/internal

The production layer may contain material necessary to maintain and verify regulatory answers, including jurisdiction records, effective-date histories, research notes, evidence captures, verification annotations, transformation artifacts, license metadata, and source material that cannot be redistributed publicly.

Internal availability does not imply public redistribution rights.

## Derived facts

A fact derived from a source and the source bytes themselves are different artifacts and are classified separately.

Building Code Map may return a project-scoped derived regulatory fact when that disclosure is permitted by the product boundary. That does not imply that the underlying source text, the complete maintained dataset, or all intermediate research is distributable.

Likewise, a public law or government record being publicly accessible does not automatically settle whether every retrieved representation, compilation, annotation, or third-party hosting copy may be redistributed.

## Historical alpha data

The archived BCM alpha product does not define the visibility of the new product.

Alpha code, data, research, generated artifacts, and source captures are treated as historical material and pass through `docs/admission-policy.md` before reuse. Nothing becomes public or service-visible solely because alpha previously exposed or stored it.

## Default for new interfaces

New schemas, endpoints, exports, examples, and debugging surfaces must fail closed against this boundary:

1. expose the minimum information needed for the product contract;
2. prefer citations and derived facts over redistributing source bytes;
3. do not add bulk-corpus access by incidental convenience;
4. do not make internal fields public merely because they already exist;
5. treat a new export surface as a visibility decision, not ordinary serialization.

If a feature needs broader visibility, change this contract deliberately and review the licensing and product consequences at the same time.
