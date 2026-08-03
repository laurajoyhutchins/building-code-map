# Building Code Map versus Building Code AST

## Building Code Map owns

- address and coordinate input;
- address normalization and geocoding;
- geographic boundary matching;
- jurisdiction and authority relationships;
- adopting, administering, and enforcement candidates;
- code-family and project-scope applicability;
- adoption instruments and model-code edition references;
- adoption, effective, transition, repeal, and supersession dates;
- statewide and local amendment applicability metadata;
- source provenance, review status, warnings, and unresolved local records.

## Building Code AST owns

- source artifact and edition identity;
- publication hierarchy;
- chapters, sections, subsections, paragraphs, and lists;
- definitions;
- exceptions and notes;
- cross-references;
- tables and figures;
- exact source spans and provenance within a publication;
- parser diagnostics and unsupported structures;
- selected provision semantics under separately versioned contracts.

## Intended composition

```text
address or point
  -> Building Code Map
  -> authority and adopted source identity
  -> Building Code AST
  -> publication structure and selected provisions
  -> separately governed review or rule evaluation
```

Neither step proves legal interpretation. Map resolution does not authorize redistribution of code
text. AST parsing does not prove that a publication applies to a location.

## Where the boundary could blur

- amendment records may need stable clause identifiers from an AST;
- code-edition comparison may join adoption timelines to publication structures;
- permit research may need both authority context and provision evidence;
- future clause-level applicability may require a reviewed rule layer between the repositories.

Those integrations should use explicit identifiers and contracts. Building Code Map should not absorb
PDF parsing, and Building Code AST should not infer jurisdictional adoption.

## Electrical Equipment Lineage and other downstream systems

Electrical Equipment Lineage preserves manufacturer, product-family, UL, and replacement claims.
Building Code Map may supply code edition and jurisdiction context for future research, but it does
not determine product compatibility, listing status, installation acceptance, or design compliance.
The same principle applies to permit tools, design-review workflows, and engineering compliance
systems: applicability context is one input, not a universal decision engine.
