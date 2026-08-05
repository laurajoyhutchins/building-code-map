# Connecticut, Washington, and Massachusetts Mixed State Wave Design

- Design date: 2026-08-04
- Repository: `laurajoyhutchins/building-code-map`
- Status: approved design, awaiting owner review of this committed specification
- Dependency: the Virginia, Oregon, and North Carolina executable wave in draft PR #39, or equivalent changes merged to `main`

## Decision

Implement the next expansion wave as:

1. Connecticut, exercising existing date-aware resolution.
2. Washington, exercising existing statewide code-family and authority modeling.
3. Massachusetts, adding one generic jurisdiction-conditioned policy-layer capability for municipal energy-code selection.

The wave combines two low-friction pilots with one bounded resolver extension. It must not add state-name branches to the resolver or create a Massachusetts-only data path.

## Why this composition

The repository already supports state profiles, code-family overrides, date rules, pending adoptions, authority candidates, source evidence, and explicit uncertainty. Connecticut and Washington fit those concepts.

Massachusetts exposes one reusable gap: a statewide baseline can be modified by a municipality-selected policy tier whose legal effect begins on a recorded date. Massachusetts publishes a current inventory covering all 351 municipalities and identifies each municipality's Base, Stretch, or Specialized energy-code tier and effective date.

This is a better capability increment than attempting New York geographic exceptions or California multi-agency project routing in the same wave.

## Alternatives considered

### Throughput only

Connecticut, Washington, South Carolina, and New Mexico would increase state count fastest, but would postpone a known resolver gap.

### Capability heavy

Massachusetts, New York, and California would combine municipality-conditioned policies, geographic exceptions, provision-level legal status, agency routing, and local amendment registries. That is too broad for one verifiable wave.

### Selected mixed wave

Connecticut and Washington validate the current architecture. Massachusetts adds exactly one reusable capability. New York and California remain separately bounded future designs.

## Governing boundaries

The implementation must preserve distinctions among:

- location input and normalized address;
- geographic boundary observations;
- legal jurisdiction identity;
- statewide adoption records;
- municipality-conditioned policy selections;
- applicable dates;
- enforcement authority;
- source snapshots;
- compiled records;
- resolver conclusions;
- stale, incomplete, conflicting, or unresolved evidence.

A municipality name, map color, community designation, or current web row must not become a timeless legal conclusion.

## Connecticut pilot

Connecticut is the historical-date pilot.

### Minimum scope

- Record the current 2022 Connecticut State Building Code.
- Record the 2018 and 2016 predecessor editions as historical adoptions.
- Implement at least building and residential resolution.
- Preserve the state adopting authority and local enforcement candidates.
- Record errata as evidence without treating an errata publication date as a new edition effective date.
- Leave unverified families and project paths unresolved.

### Date semantics

The existing request field `applicability_date` is reused. For covered Connecticut families, it represents the permit-application date because the official edition history is keyed to permit applications.

The resolver does not gain a Connecticut-specific date field. The Connecticut profile and rule pack must explain the controlling date meaning and require the permit record as evidence.

A caller that omits `applicability_date` receives the existing current-date default. That result may answer a current query, but must not be presented as a historical permit-date determination.

### Required boundaries

- 2018-09-30 selects the 2016 edition.
- 2018-10-01 selects the 2018 edition.
- 2022-09-30 selects the 2018 edition.
- 2022-10-01 selects the 2022 edition.
- A historical conclusion lacking the permit-application date remains insufficiently supported.

Additional code families are optional and require exact adoption, amendment, authority, and temporal evidence.

## Washington pilot

Washington is the broad statewide-family pilot.

### Minimum scope

- Record the 2021 Washington State Building Code family effective 2024-03-15.
- Cover building, existing building, residential, mechanical, fire, plumbing, commercial energy, and residential energy where official sources support exact mappings.
- Preserve the State Building Code Council as adopting authority for those families.
- Preserve local administration and enforcement separately.
- Record the 2024 adoption cycle as pending rulemaking without inferred operative dates.
- Preserve electrical adoption as a separate Department of Labor and Industries authority path.
- Leave wildland-urban-interface resolution outside this pilot because the 2026 reserved-status and mapping dependencies require separate treatment.

### Required boundaries

- 2024-03-14 selects the prior recorded edition when that predecessor is modeled.
- 2024-03-15 selects the 2021 edition.
- Pending 2024-cycle records never appear as current.
- Electrical queries do not inherit an IRC electrical chapter or State Building Code Council adoption by implication.
- Local amendments and enforcement details remain explicit unresolved records where not verified.

## Massachusetts pilot

Massachusetts is the capability pilot.

### Minimum scope

- Record the statewide 10th-edition building-code baseline and verified concurrency history.
- Record statewide Base residential and commercial energy adoptions.
- Record Stretch and Specialized residential and commercial adoption bundles from 225 CMR 22 and 225 CMR 23.
- Compile a versioned snapshot of the official municipal inventory covering all 351 municipalities.
- Select Base, Stretch, or Specialized by canonical municipality identity and applicable date.
- Preserve the municipality record, source snapshot, effective date, and verification state in the result.
- Keep non-energy local rules outside the policy-layer mechanism.

The new mechanism selects among modeled adoption bundles. It does not interpret ordinances, infer amendments, or rewrite statewide adoption records.

## Jurisdiction-conditioned policy-layer contract

### Data ownership and loading

State profiles continue to own statewide authorities and adoptions.

Policy layers are independent catalog artifacts stored under a generic namespace:

`backend/data/regulatory/policy-layers/`

The Massachusetts artifact is:

`backend/data/regulatory/policy-layers/massachusetts-energy.json`

The catalog loads and validates profiles, rule packs, and policy layers. It indexes policy layers by normalized `state_id` and `code_family`.

State profiles remain schema version 1.0. No mechanical migration of existing profiles is required. Cross-artifact validation ensures that policy-layer adoption IDs exist in the referenced state profile.

The policy-layer schema begins at version 1.0.

### Layer fields

Each layer contains:

- `schema_version`;
- stable `layer_id`;
- `state_id`;
- covered `code_families`;
- `jurisdiction_kind`, initially `municipality`;
- `selector_key`, initially `geography.municipality.feature_id`;
- source snapshot metadata;
- expected and observed coverage counts;
- deterministic content hash;
- policy-tier definitions;
- jurisdiction policy records;
- verification metadata.

Only one active layer may exist for a state and code family.

### Policy tiers

The Massachusetts layer defines exactly three mutually exclusive tiers:

- `base`;
- `stretch`;
- `specialized`.

A tier selects an adoption bundle rather than creating an adoption during resolution.

- Base selects the statewide Base adoption.
- Stretch selects the applicable Stretch bundle.
- Specialized selects the modeled Stretch provisions plus Specialized appendices as one defined bundle.

The resolver selects exactly one tier for each covered municipality, family, and date. It must not present all three as concurrent alternatives.

### Jurisdiction records

Each record contains:

- canonical `jurisdiction_id`, equal to the municipality boundary `feature_id` used by the legal-jurisdiction pipeline;
- municipality name for audit only;
- `policy_tier_id`;
- `effective_from`;
- optional `effective_through`;
- optional `adopted_on` only when an official source establishes it;
- source identifier and row locator;
- verification status and confidence;
- notes identifying incomplete history when needed.

Missing adoption dates remain absent. They are not copied from effective dates or inferred from program membership.

### Canonical identity

Resolution uses `GeographicContext.Municipality.FeatureID`. Runtime free-text municipality matching is prohibited.

Ingestion may use a reviewed alias table to map official names to canonical feature IDs. Compilation rejects:

- duplicate canonical IDs;
- one source row mapped to multiple jurisdictions;
- conflicting active tiers;
- unrecognized source municipalities;
- missing canonical municipalities;
- any snapshot claiming complete coverage with a count other than 351.

### Temporal behavior

Policy records are intervals.

The evaluator selects the record containing `applicability_date`. When the official inventory supplies only a current tier and its effective date, the compiler does not infer the prior tier. Earlier queries remain unresolved unless a predecessor interval is independently sourced.

Future records remain pending until their effective date.

### Source snapshots and freshness

The resolver performs no live scraping.

The normalized snapshot records:

- source URL and title;
- observation date;
- source-declared status date;
- extraction method;
- source and normalized row counts;
- canonical match count;
- content hash;
- compiler version;
- refresh due date.

The initial snapshot refresh due date is 2026-09-04. After that date, the resolver may report the last observed tier, but must mark it stale, expose the observation date, and require current official confirmation.

Freshness remains data-driven through regulatory rules. No Massachusetts condition belongs in resolver code.

## Result contract

The resolution result gains an optional `jurisdiction_policy_selection` object containing:

- layer ID and schema version;
- jurisdiction ID and name;
- selected tier ID;
- selected adoption IDs;
- effective interval;
- source snapshot ID and observation date;
- freshness state;
- verification metadata.

The existing `adoptions` array continues to contain the adoption records selected for practical applicability. The policy-selection object explains why the municipal bundle was selected.

This is an additive public-result change. `ResultSchemaVersion` advances from `1.0` to `1.1`. Existing result fields retain their meanings.

## Resolution flow

For a covered family:

1. Resolve geography and legal jurisdiction through the existing pipeline.
2. Resolve the state profile and statewide baseline.
3. Look up a policy layer by state and code family.
4. Require one municipality `feature_id` when a layer applies.
5. Select the time-valid jurisdiction record.
6. Select the record's pre-modeled adoption bundle.
7. Return adoptions and a distinct policy-selection provenance object.
8. Attach snapshot identity, effective dates, freshness, evidence, and unresolved warnings.

The policy selection does not mutate statewide adoption records.

## Failure and uncertainty behavior

Resolution fails closed when:

- municipality identity is missing or ambiguous;
- a source row cannot map to one canonical jurisdiction;
- active records overlap;
- no time-valid record exists;
- a tier references missing adoptions;
- a complete snapshot has incomplete coverage;
- snapshot or generated-artifact hashes fail;
- multiple layers claim the same state and family.

A stale snapshot is reported as a stale observation, not erased and not presented as freshly verified.

The resolver must not infer a tier from community-program membership, neighbors, county, population, a municipal website outside the normalized snapshot, a prior or future interval, or the presence of code text.

## Compiler and quality reporting

The compiler validates:

- schema and referential integrity;
- jurisdiction IDs;
- interval order and overlap;
- mutually exclusive tiers;
- source and normalized coverage counts;
- deterministic ordering and serialization;
- content hashes;
- generated-artifact drift;
- freshness rules.

The nationwide quality report separately reports:

- state-profile coverage;
- code-family coverage;
- policy-layer jurisdiction coverage;
- current versus historical interval completeness;
- fresh, stale, partially verified, conflicting, and unresolved evidence.

Complete current municipality coverage must not be described as complete historical coverage.

## Testing strategy

Add failing tests before each implementation phase.

### Connecticut

- 2016/2018 boundary;
- 2018/2022 boundary;
- current query using the default applicability date;
- historical query using an explicit permit-application date;
- warning or required-record behavior when historical date evidence is absent;
- unresolved family behavior.

### Washington

- 2018/2021 boundary;
- current resolution for every included family;
- pending 2024-cycle exclusion;
- separate electrical authority behavior;
- unresolved wildland-urban-interface behavior;
- local enforcement candidate preservation.

### Massachusetts

- one Base municipality;
- one Stretch municipality;
- one Specialized municipality;
- residential and commercial bundles;
- exact effective-date boundary;
- future record exclusion;
- missing predecessor interval;
- missing municipality;
- unknown feature ID;
- duplicate and overlapping record rejection;
- 350-row incomplete snapshot rejection;
- stale warning after 2026-09-04;
- deterministic compilation and drift detection;
- result schema version 1.1 and policy-selection provenance.

All existing state tests must continue to pass. Existing profile fixtures should not require migration because the profile schema remains 1.0.

## Implementation sequence

1. Land or rebase onto the Virginia, Oregon, and North Carolina wave.
2. Add Connecticut test-first fixtures, profile, rules, pilot report, and historical records.
3. Add Washington test-first fixtures, profile, rules, pilot report, authority separation, and pending records.
4. Add policy-layer schema, catalog loading, and compiler tests without Massachusetts data.
5. Add the generic evaluator and result-schema 1.1 provenance object.
6. Ingest and normalize the Massachusetts municipal snapshot.
7. Add Massachusetts statewide profile, adoption bundles, rules, and pilot report.
8. Add Massachusetts resolver, freshness, coverage, and drift tests.
9. Update readiness documentation and nationwide quality reporting.
10. Run exact-head backend, regulatory-data, frontend, and documentation verification.

The policy engine and Massachusetts dataset remain separate reviewable commits.

## Permitted readiness claims

After successful implementation, documentation may claim:

- Connecticut is an executable, partially verified pilot for recorded families and permit-date windows.
- Washington is an executable, partially verified pilot for recorded statewide families.
- Massachusetts is an executable, partially verified pilot for statewide baselines and municipality-conditioned energy tiers.
- The Massachusetts current snapshot covers all 351 municipality identities only when compilation proves it.

It must not claim complete code-family coverage, complete local amendments, complete Massachusetts policy history, resolved Washington electrical adoption, resolved Massachusetts non-energy local rules, or nationwide policy-layer completeness.

## Deferred work

South Carolina and New Mexico remain candidates for the next throughput wave.

New York remains deferred pending a geographic-exception and provision-status design.

California remains deferred pending a bounded design for project and occupancy routing, state-agency jurisdiction, and local amendment ingestion.

## Source baseline

Verified on 2026-08-04:

### Connecticut

- Connecticut Department of Administrative Services, Connecticut State Building Code regulations: `https://portal.ct.gov/en/das/office-of-state-building-inspector/connecticut-state-building-code/regulations`

### Washington

- State Codes, Regulations & Guidelines: `https://sbcc.wa.gov/state-codes-regulations-guidelines`
- Building Code Amendments: `https://sbcc.wa.gov/state-codes-regulations-guidelines/state-building-code/building-code-amendments`
- 2024 Code Adoption Cycle: `https://sbcc.wa.gov/2024-code-adoption-cycle`

### Massachusetts

- Massachusetts Building Energy Codes: `https://www.mass.gov/info-details/massachusetts-building-energy-codes`
- Building Energy Code Adoption by Municipality: `https://www.mass.gov/info-details/massachusetts-building-energy-code-adoption-by-municipality`
- 225 CMR 22.00: `https://www.mass.gov/regulations/225-CMR-2200-massachusetts-stretch-code-and-specialized-code-for-low-rise-residential-2025-residential-low-rise-amendments-to-iecc2021-and-irc-2021-chapter-11-energy-efficiency`
- 225 CMR 23.00: `https://www.mass.gov/regulations/225-CMR-2300-massachusetts-stretch-code-and-specialized-code-for-commercial-multi-family-and-all-other-construction-2025-amendments-to-iecc2021-and-ashrae-standards-901-2019`

These sources establish the design baseline. Each executable record still requires exact source attribution and implementation-time verification.
