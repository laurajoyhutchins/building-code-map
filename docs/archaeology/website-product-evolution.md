# Website and product-design evolution

## Technical workbench

The early product surface centered on a MapLibre explorer, boundary layers, feature inspection, and
resolution panels. That surface was valuable for debugging geometry and source records, but it made
the user reconstruct the central product question.

Map-specific correctness work remained important:

- PR #32 reconciled layers whether map style or registry data arrived first;
- PR #34 replaced globally assumed feature IDs with `(layerFamily, featureId)`;
- unresolved issue #7 calls for truly on-demand boundary details.

## Public search-first product

Merged PR #36 established:

- a restrained public root route;
- the technical GIS console at `/explorer`;
- coordinate input before local geocoding existed;
- code-family and date controls;
- plain-language resolution notices;
- authorities, adopted codes, local confirmation, jurisdiction structure, and sources;
- responsive layout and accessible form semantics.

Merged PR #37 extended the same interface to civic addresses and displayed matched-address precision,
source, vintage, and interpolation warnings.

## Credibility correction

The product explicitly rejected conspicuous trust signaling such as:

- “Source-backed”;
- “Uncertainty preserved”;
- “Verify with the AHJ”;
- compact trust-badge rows;
- decorative metrics and runtime telemetry;
- generic promotional claims.

The concern was not that provenance or caveats were unimportant. The concern was that generic
assurance language made a rigorous domain product look formulaic or synthetically assembled.

## Active design principle

Demonstrate trustworthiness through behavior and detail:

- label the matched address and geocoding precision;
- show effective dates;
- name authority candidates and roles;
- surface special conditions;
- list required local records and warnings;
- expose jurisdiction relationships;
- link the actual sources with access and check dates;
- use restrained, domain-specific language;
- retain necessary engineering and legal caveats without turning them into badges.

## Current implemented visual behavior

Current source uses an off-white page, compact borders, blue/green accent variables, oversized
headline typography, generous spacing, a narrow result measure, and two-column result sections that
collapse to one column below 720 px. It avoids card grids and decorative status pills.

The broader “phthalo” and “web revival” direction is part of historical design context. Only the
specific CSS and component behavior in the current tree should be claimed as implemented.

## Remaining public-readiness concerns

- the site can be public while supported data remains only three partially verified pilots;
- missing snapshots and unsupported locations need precise empty and error states;
- map chunk size remains a build warning;
- accessibility and responsive behavior have tests and source evidence, but no complete independent
  audit is recorded;
- source completeness disclosures must remain tied to actual records, not marketing copy.
