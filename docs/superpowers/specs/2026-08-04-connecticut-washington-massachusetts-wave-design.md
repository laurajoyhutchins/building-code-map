# Connecticut, Washington, and Massachusetts Mixed State Wave Design

- Design date: 2026-08-04
- Repository: `laurajoyhutchins/building-code-map`
- Status: approved design, awaiting owner review of the committed specification
- Depends on: the Virginia, Oregon, and North Carolina executable wave in draft PR #39, or equivalent changes merged to `main`

## Decision

Implement the next state-expansion wave as a mixed portfolio:

1. Connecticut, using the existing resolver to strengthen permit-application-date history.
2. Washington, using the existing resolver to add a broad statewide code-family pilot while preserving separate authority boundaries.
3. Massachusetts, after adding one generic jurisdiction-conditioned policy-layer capability for municipal energy-code selection.

This wave deliberately combines two low-friction state profiles with one bounded resolver extension. It must not create state-name branches in the resolver or turn Massachusetts into a one-off special case.

## Why this composition

The repository already supports source-backed state profiles, code-family overrides, date rules, pending adoptions, local enforcement candidates, and explicit unresolved records. Connecticut and Washington fit those concepts.

Massachusetts exposes a reusable missing capability: a statewide baseline can be modified by a municipality-selected policy tier whose legal effect begins on a recorded date. The state publishes a current municipality-by-municipality inventory for all 351 municipalities, including Base, Stretch, or Specialized energy-code status and effective date. That makes Massachusetts a suitable first jurisdiction-conditioned policy layer, provided the system preserves source snapshots, canonical municipality identity, temporal validity, and uncertainty.

## Alternatives considered

### Throughput-only wave

Connecticut, Washington, South Carolina, and New Mexico could be implemented without a resolver extension. This would maximize state count but postpone a known architecture gap and leave Massachusetts research-only.

### Capability-only wave

Massachusetts, New York, and California could be used to expand the resolver aggressively. This would combine municipality-conditioned policies, geographic exceptions, provision-level legal status, multi-agency project routing, and local amendment registries in one wave. The result would be too broad to verify cleanly.

### Selected mixed wave

Connecticut and Washington exercise the current model. Massachusetts adds exactly one generic capability. New York and California remain deferred until their distinct routing problems can be designed independently.

## Governing principles

The implementation must preserve the distinction among:

- address and normalized location input;
- geographic boundary observations;
- canonical legal jurisdiction identity;
- statewide code adoptions;
- municipality-conditioned policy selections;
- effective-date observations;
- enforcement authority;
- source snapshots;
- compiler output;
- resolver conclusions;
- unresolved or stale evidence.

A municipality name match, program designation, map color, or current web-table row must not silently become a timeless legal conclusion.

## Scope

### Connecticut pilot

Connecticut is the historical-date pilot.

Minimum executable scope:

- record the current 2022 Connecticut State Building Code;
- record the 2018 and 2016 predecessor editions as historical adoptions;
- use permit-application date as the controlling temporal input for the recorded code families;
- implement at least building and residential code-family resolution;
- preserve the state adopting authority and local enforcement candidates;
- record errata as source evidence without treating an errata publication date as a new code-edition effective date;
- return unresolved results for code families or project paths not independently verified.

Required boundary behavior:

- a permit application dated 2022-09-30 resolves to the 2018 edition for covered families;
- a permit application dated 2022-10-01 resolves to the 2022 edition;
- a permit application dated 2018-09-30 resolves to the 2016 edition;
- a permit application dated 2018-10-01 resolves to the 2018 edition;
- a query without the controlling permit-application date must not pretend to provide a historical edition choice.

Additional code families may be added only when their exact Connecticut adoption, amendment, authority, and temporal scope are independently source-backed. They are not required for pilot completion.

### Washington pilot

Washington is the broad statewide-family pilot.

Minimum executable scope:

- record the 2021 Washington State Building Code family effective 2024-03-15;
- cover building, existing building, residential, mechanical, fire, plumbing, commercial energy, and residential energy where official state sources support exact mappings;
- preserve the State Building Code Council as the adopting authority for those recorded families;
- preserve local administering and enforcement candidates separately from statewide adoption;
- record the 2024 adoption cycle as pending rulemaking without assigning operative dates before final adopted rules provide them;
- keep electrical adoption outside the State Building Code Council family and identify the Department of Labor and Industries as a separate authority path requiring its own source-backed profile increment;
- leave the wildland-urban-interface path unresolved in this pilot because the 2026 reserved-status and mapping dependencies require separate treatment.

Required boundary behavior:

- an applicable date of 2024-03-14 resolves to the prior recorded edition for covered families when the predecessor record is present;
- an applicable date of 2024-03-15 resolves to the 2021 edition;
- pending 2024-cycle records never appear as current adoptions;
- an electrical query does not inherit an International Residential Code electrical chapter or a State Building Code Council adoption by implication;
- unresolved local amendments and enforcement details remain visible.

### Massachusetts pilot

Massachusetts is the capability pilot.

Minimum executable scope:

- record the statewide 10th-edition building-code baseline and its verified effective/concurrency history;
- record statewide Base energy-code adoptions for residential and commercial energy;
- record Stretch and Specialized energy-code adoption bundles from 225 CMR 22 and 225 CMR 23;
- compile a versioned snapshot of the official municipal energy-code inventory covering all 351 municipalities;
- select Base, Stretch, or Specialized policy by canonical municipality identity and applicable date;
- preserve the municipality record, source snapshot, effective date, and verification state in the resolver result;
- keep non-energy local rules and amendments outside the policy-layer mechanism.

The policy-layer capability is limited to selecting among already modeled adoption bundles. It does not interpret ordinance text, infer amendments, or rewrite statewide adoption records.

## Jurisdiction-conditioned policy-layer contract

### Data ownership

The state profile owns statewide authorities and adoption records.

A separate policy-layer artifact owns municipality-conditioned selections. The initial artifact should be stored under a generic regulatory-data namespace rather than embedded in Massachusetts-specific resolver code, for example:

`backend/data/regulatory/policy-layers/massachusetts-energy.json`

The exact filename may follow repository naming conventions, but the compiled concept must remain a generic policy layer.

### Layer fields

Each policy layer must contain:

- `schema_version`;
- stable `layer_id`;
- `state_id`;
- one or more applicable `code_family` values;
- `jurisdiction_kind`, initially `municipality`;
- `selector_key`, referencing the canonical legal-jurisdiction identifier emitted by the jurisdiction pipeline;
- source snapshot metadata;
- expected and observed coverage counts;
- a deterministic content hash;
- one or more policy-tier definitions;
- jurisdiction policy records;
- verification metadata.

### Policy tiers

The Massachusetts layer has exactly three mutually exclusive tiers:

- `base`;
- `stretch`;
- `specialized`.

A tier selects an adoption bundle rather than manufacturing a new adoption at query time.

- Base selects the statewide Base residential or commercial energy adoption.
- Stretch selects the applicable Stretch adoption bundle.
- Specialized selects the applicable Stretch provisions plus the Specialized appendices represented by the modeled Specialized adoption bundle.

The resolver must select exactly one tier for each covered municipality, code family, and applicable date. It must never combine Base, Stretch, and Specialized as three independent current alternatives.

### Jurisdiction records

Each jurisdiction policy record must contain:

- canonical `jurisdiction_id`;
- human-readable municipality name for auditability only;
- `policy_tier_id`;
- `effective_from`;
- optional `effective_through`;
- optional `adopted_on` when an official source establishes it;
- source identifier and source-row locator;
- verification status and confidence;
- optional notes describing incomplete history.

Missing adoption dates remain null. They must not be copied from effective dates or inferred from program membership.

### Canonical identity

Runtime matching must use the canonical legal-jurisdiction identifier already produced by the location and authority pipeline. Free-text municipality-name matching is prohibited at resolution time.

Ingestion may use a reviewed alias table to map official source names to canonical IDs. The compiler must reject:

- duplicate canonical municipality IDs;
- source rows that map to multiple jurisdictions;
- canonical jurisdictions with conflicting active tiers;
- unrecognized source municipalities;
- coverage counts other than 351 for a snapshot declared complete.

### Temporal behavior

Policy records are intervals, not timeless labels.

The evaluator selects the record whose interval contains the applicable date. When the current official inventory supplies only the current tier and its effective date, the compiler must not infer the prior tier. Queries before that effective date remain unresolved unless a predecessor interval is independently sourced.

A future effective date remains pending and does not alter the current result.

### Freshness and source snapshots

The resolver performs no live web scraping.

The official Massachusetts inventory is ingested into a normalized, reviewable snapshot that records:

- source URL and title;
- page observation date;
- source-declared status date when present;
- extraction method;
- source record count;
- normalized record count;
- canonical match count;
- content hash;
- compiler version;
- refresh due date.

The initial snapshot must have a concrete refresh due date of 2026-09-04. After that date, the system may still report what the snapshot observed, but it must mark current status as stale and require renewed official confirmation rather than presenting the tier as freshly verified.

Freshness policy remains data-driven through the existing regulatory rule-pack mechanism. No Massachusetts state-name condition belongs in resolver code.

## Resolution flow

For a query involving a policy-layer code family:

1. Resolve geographic observations and canonical legal jurisdiction through the existing pipeline.
2. Resolve the statewide state profile and baseline adoption candidates.
3. Determine whether the requested code family references a policy layer.
4. Require one unambiguous canonical municipality identifier.
5. Find the time-valid jurisdiction policy record.
6. Select the record's policy tier and its pre-modeled adoption bundle.
7. Return the statewide baseline and municipality policy selection as distinct evidence-bearing objects.
8. Attach snapshot identity, effective date, source locator, freshness state, and unresolved warnings.

The policy selection must not mutate the underlying statewide adoption record. Downstream projections may present the combined practical result, but the normalized response must preserve both claims.

## Failure and uncertainty behavior

The evaluator fails closed when:

- the municipality is missing or ambiguous;
- the official row cannot be mapped to a canonical jurisdiction;
- two active records overlap for the same municipality and family;
- no time-valid record exists;
- a policy tier references a missing adoption bundle;
- source coverage is incomplete while the layer claims complete coverage;
- the snapshot or compiled artifact fails hash or drift validation.

A stale snapshot is not equivalent to a missing snapshot. The resolver may report the last observed tier, but must downgrade verification, expose the observation date, and require current confirmation.

The resolver must not infer a tier from:

- Green Community or Climate Leader Community membership;
- neighboring municipalities;
- population or county;
- a municipality's current website without a normalized source record;
- a prior or later policy interval;
- the presence of Stretch or Specialized code text in the repository.

## Compiler and validation design

The regulatory compiler must validate policy layers alongside state profiles and rule packs.

Validation must include:

- schema validation;
- referential integrity to state profiles, code families, adoption IDs, sources, and canonical jurisdictions;
- interval ordering and overlap detection;
- mutually exclusive tier selection;
- source and normalized coverage counts;
- stable ordering and deterministic serialization;
- content-hash verification;
- generated-artifact drift checks;
- source-health and refresh-due checks.

The nationwide quality report must distinguish:

- state profile coverage;
- code-family coverage;
- municipality policy-layer coverage;
- current versus historical interval completeness;
- fresh, stale, partially verified, and unresolved evidence.

A state must not be reported as historically complete merely because its current municipal snapshot covers all jurisdictions.

## Testing strategy

### Test-first state fixtures

Add failing resolver tests before each profile or policy artifact exists.

Connecticut tests:

- 2016/2018 boundary on 2018-10-01;
- 2018/2022 boundary on 2022-10-01;
- missing permit-application date for a historical query;
- current query for a covered family;
- unresolved family behavior.

Washington tests:

- 2018/2021 boundary on 2024-03-15;
- current resolution for every included family;
- pending 2024-cycle exclusion;
- separate electrical authority behavior;
- unresolved wildland-urban-interface behavior;
- local enforcement candidate preservation.

Massachusetts tests:

- one Base municipality;
- one Stretch municipality;
- one Specialized municipality;
- residential and commercial energy bundles;
- exact effective-date boundary;
- future policy record exclusion;
- missing predecessor interval;
- unknown municipality;
- ambiguous municipality identity;
- duplicate and overlapping records rejected by the compiler;
- 350-row incomplete snapshot rejected when declared complete;
- stale snapshot warning after 2026-09-04;
- deterministic compilation and drift detection.

### Regression requirements

All existing Colorado, Florida, New Jersey, Virginia, Oregon, and North Carolina tests must continue to pass unchanged unless a generic schema migration requires mechanical fixture updates. A generic migration must preserve their observable resolver behavior.

No test may encode a state-name branch as the expected implementation mechanism.

## Implementation sequence

1. Land or rebase onto the Virginia, Oregon, and North Carolina wave.
2. Add Connecticut test-first fixtures, profile, rule pack, pilot report, and historical date records.
3. Add Washington test-first fixtures, profile, rule pack, pilot report, authority separations, and pending-cycle records.
4. Add policy-layer schema and compiler tests without Massachusetts data.
5. Add the generic policy-layer evaluator and result-provenance fields.
6. Ingest and normalize the Massachusetts municipal energy inventory snapshot.
7. Add Massachusetts statewide profile, energy adoption bundles, rule pack, and pilot report.
8. Add Massachusetts resolver, compiler, freshness, and coverage tests.
9. Update readiness documentation, nationwide quality reporting, and public status language.
10. Run exact-head backend, regulatory-data, frontend, and documentation verification.

Implementation commits should remain reviewable by phase. The policy-layer engine and the Massachusetts data snapshot should not arrive as one opaque commit.

## Readiness claims

After this wave, documentation may claim:

- Connecticut is an executable, partially verified pilot for the recorded families and historical permit-date windows.
- Washington is an executable, partially verified pilot for the recorded statewide code families.
- Massachusetts is an executable, partially verified pilot for the statewide baseline and municipality-conditioned energy-code tier selection.
- the Massachusetts current municipal snapshot has complete 351-municipality identity coverage if the compiler proves it.

Documentation must not claim:

- complete code-family coverage for any of the three states;
- complete local amendment coverage;
- complete historical Massachusetts municipal tier history;
- resolved electrical authority in Washington unless a separate source-backed increment is implemented;
- resolved Massachusetts non-energy local rules;
- nationwide jurisdiction-conditioned policy support beyond the generic mechanism and its verified datasets.

## Deferred work

South Carolina and New Mexico remain candidates for the next throughput-oriented state wave.

New York remains deferred until the repository has a separately designed geographic-exception and provision-level legal-status mechanism.

California remains deferred until project/occupancy routing, state-agency jurisdiction, and local-amendment registry ingestion are designed as a bounded subsystem.

## Source baseline

Sources verified during design on 2026-08-04 include:

### Connecticut

- Connecticut Department of Administrative Services, Connecticut State Building Code regulations: `https://portal.ct.gov/en/das/office-of-state-building-inspector/connecticut-state-building-code/regulations`

### Washington

- Washington State Building Code Council, State Codes, Regulations & Guidelines: `https://sbcc.wa.gov/state-codes-regulations-guidelines`
- Washington State Building Code Council, Building Code Amendments: `https://sbcc.wa.gov/state-codes-regulations-guidelines/state-building-code/building-code-amendments`
- Washington State Building Code Council, 2024 Code Adoption Cycle: `https://sbcc.wa.gov/2024-code-adoption-cycle`

### Massachusetts

- Massachusetts Department of Energy Resources, Massachusetts Building Energy Codes: `https://www.mass.gov/info-details/massachusetts-building-energy-codes`
- Massachusetts Department of Energy Resources, Building Energy Code Adoption by Municipality: `https://www.mass.gov/info-details/massachusetts-building-energy-code-adoption-by-municipality`
- Massachusetts 225 CMR 22.00: `https://www.mass.gov/regulations/225-CMR-2200-massachusetts-stretch-code-and-specialized-code-for-low-rise-residential-2025-residential-low-rise-amendments-to-iecc2021-and-irc-2021-chapter-11-energy-efficiency`
- Massachusetts 225 CMR 23.00: `https://www.mass.gov/regulations/225-CMR-2300-massachusetts-stretch-code-and-specialized-code-for-commercial-multi-family-and-all-other-construction-2025-amendments-to-iecc2021-and-ashrae-standards-901-2019`

These sources establish the design baseline, not the complete implementation evidence set. Each executable record still requires exact source attribution and verification during implementation.
