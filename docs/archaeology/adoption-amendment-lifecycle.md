# Adoption and amendment lifecycle

## Distinct records

Building Code Map treats the following as separate facts:

- model-code publication;
- enabling authority;
- adopting instrument;
- adoption date;
- effective or operative date;
- optional transition period;
- mandatory date;
- amendment or override;
- enforcement assignment;
- repeal or supersession;
- source observation and verification date.

A newer publication is not automatically the governing instrument. A department webpage is not automatically an adoption instrument. Local enforcement does not necessarily imply local adoption power.

## Executable pilot profiles

The repository currently compiles six state profiles:

- Colorado
- Florida
- New Jersey
- Virginia
- Oregon
- North Carolina

The profiles exercise different regimes, including primarily local adoption, statewide baselines with local administration, specialty-code families, state trade-code overrides, prior-edition election windows, and future transitions whose operative date depends on an external legal trigger.

Pending Oregon and North Carolina records remain pending until their required conditions are supported. They do not leak into current results.

## Production-readiness scopes

Colorado, Florida, and New Jersey have scoped production manifests that currently pass the repository's readiness audit. The audit requires declared coverage, continuous verified timelines within scope, primary evidence, source health, and evidence-backed fixtures.

A passing scope means only that the declared scope satisfies the gate. It does not establish complete statewide coverage of every:

- municipality or county;
- code family;
- local amendment;
- enforcement contract;
- special project or occupancy;
- optional appendix;
- design criterion;
- historical interval.

Virginia, Oregon, and North Carolina remain executable, partially verified pilots without production-ready manifests.

## Resolution behavior

For a point, as-of date, and available project context, the resolver:

1. loads the compatible state profile;
2. derives incorporated or unincorporated policy from boundary observations;
3. applies code-family and project-type rules;
4. filters records by temporal status;
5. expands authority and enforcement candidates;
6. reports missing local records, conflicts, warnings, and verification status.

## Required next contract

Nationwide scaling requires explicit inheritance and local-jurisdiction records so a state rule can be represented once and legally applied to classified entities without duplicating the same adoption thousands of times.
