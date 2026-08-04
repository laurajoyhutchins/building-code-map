# Next-State Regulatory Coverage Wave

- Initial research date: 2026-08-02
- Executable-wave verification date: 2026-08-04
- Scope: Virginia, Oregon, North Carolina, Massachusetts, New York, and California
- Status: three executable pilots; three resolver-dependent research tracks

This document separates the implemented state-data wave from states that still need new
resolver capabilities.

Executable does not mean complete. Every pilot remains source-backed, partially verified,
and explicit about required local records.

## Executable pilots

### Virginia

Virginia now has an executable pilot for statewide construction and existing-building
baselines.

The profile records local enforcement candidates and the transition period from 2024-01-18
through 2025-01-17.

Trade codes, operational fire authority, specialized project paths, and complete historical
editions remain unresolved.

### Oregon

Oregon now has an executable pilot for the recorded building, existing-building, energy,
residential, electrical, plumbing, and mechanical specialty-code families.

Optional and mandatory dates remain separate. The anticipated 2026 residential code remains
pending and is not returned as current.

Local program boundaries, interim amendments, alternate methods, and prior-edition records
remain incomplete.

### North Carolina

North Carolina now has an executable pilot for the current 2018 building and residential
codes.

The profile records State Fire Marshal supervision and local inspection candidates. The 2024
code remains pending because implementation depends on a certification event and statutory
delay.

No calendar date is inferred. The resolver continues to return the 2018 baseline and requires
the current certification record.

## Canonical artifacts

- `reports/pilots/virginia.md`
- `reports/pilots/oregon.md`
- `reports/pilots/north-carolina.md`
- `backend/data/regulatory/virginia.json`
- `backend/data/regulatory/oregon.json`
- `backend/data/regulatory/north-carolina.json`
- `backend/data/regulatory/rules/virginia.json`
- `backend/data/regulatory/rules/oregon.json`
- `backend/data/regulatory/rules/north-carolina.json`

## Resolver-dependent research tracks

### Massachusetts

Massachusetts remains research-only.

Municipality-conditioned policy selection is required before Base, Stretch, and Specialized
energy-code paths can be compiled safely.

### New York

New York remains research-only.

The resolver needs an explicit New York City exception and provision-level legal-status
overrides before the state profile can be executable.

### California

California remains research-only.

The state needs project and occupancy routing among state agencies and local authorities. It
also needs a separately refreshed local-ordinance registry.

This is a bounded subsystem, not another flat state profile.

## Cross-state requirements exposed

1. Geography-conditioned policies for named municipalities or maintained policy layers.
2. Conditional future dates based on official certifications or events.
3. Provision-level suspension, conflict, and supersession states.
4. Edition-specific local-amendment registries.
5. Multi-agency project and occupancy jurisdiction.
6. Independent effective, optional, mandatory, and replacement dates.

## Next work

- Expand the executable pilots through source-backed code-family and authority increments.
- Add prior-edition records before claiming complete historical transition results.
- Design municipality-conditioned policy selection before promoting Massachusetts.
- Design geographic exceptions and provision-level status before promoting New York.
- Design California routing and amendment ingestion before promoting California.
