# Jurisdiction and authority model

## Core distinction

A containing polygon is a geographic observation. A governing authority is a legal and administrative conclusion supported by policy and evidence. Building Code Map keeps those concepts separate.

## Geographic observations

The boundary layer may report:

- state;
- county or county equivalent;
- municipality;
- incorporated or unincorporated status;
- special land-use area;
- American Indian area;
- NERIS fire-service jurisdiction.

Peer state, county, or municipality overlaps are explicit ambiguity. Special-area, tribal-area, and NERIS observations remain contextual and are not automatically promoted to adopting or enforcing authority.

## Regulatory interpretation

A state profile and its rule pack may produce:

- candidate adopting authorities;
- candidate enforcing authorities;
- statewide base-code records;
- local adoption or amendment requirements;
- project-type and code-family overrides;
- delegated or special-program paths;
- unresolved local records;
- warnings and source references.

The current result contract is `resolution-result` schema 1.0. Authority is represented through candidates, roles, relationships, and `authority_path`; there is no separate current `authority_graph` schema contract.

## Public trust boundary

The public API accepts a point and derives observations server-side. It does not accept normalized jurisdiction context as trusted caller input. This prevents a plausible city or county label from bypassing the boundary evidence layer.

## Current strengths

- deterministic point-to-observation resolution;
- explicit overlap ambiguity;
- generic profile-driven regulatory resolution;
- separation of adoption, enforcement, and special-program candidates;
- evidence and verification references;
- required-local-record and insufficient-evidence outputs.

## Missing national layer

The runtime still lacks a canonical nationwide governmental-entity registry connecting:

```text
boundary feature
    -> legal entity
    -> adoption and amendment powers
    -> enforcement assignment
    -> inheritance rule
    -> validity interval
```

That missing layer is the main obstacle to measuring and completing every U.S. jurisdiction. State-level pilot profiles can identify local follow-up, but they do not yet enumerate and classify every expected municipality, county, township, consolidated government, state agency, or special authority.

## Safety rule

When legal authority cannot be established from the admitted data and policy, the resolver must return uncertainty or a required confirmation. Geographic plausibility is never enough.
