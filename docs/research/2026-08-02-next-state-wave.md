# Next-State Regulatory Coverage Wave

- Initial research date: 2026-08-02
- Executable-wave verification date: 2026-08-04
- Scope: Virginia, Oregon, North Carolina, Massachusetts, New York, and California
- Status: three executable pilots; three resolver-dependent research tracks

This document records the boundary between the implemented state-data wave and the states that still require new resolver capabilities. Executable does not mean complete: every pilot remains source-backed, partially verified, and explicit about required local records.

## Wave status

| State | Status | Implemented or blocking boundary |
| --- | --- | --- |
| Virginia | Executable pilot | Statewide construction and existing-building baselines, local enforcement candidates, and the 2024-01-18 through 2025-01-17 transition rule. |
| Oregon | Executable pilot | Current specialty-code adoptions, code-family-specific optional and mandatory dates, local-program candidates, and pending 2026 residential adoption. |
| North Carolina | Executable pilot | Current 2018 building and residential baselines, local inspection with state supervision, and a conditionally pending 2024 transition without an invented date. |
| Massachusetts | Research only | Requires municipality-conditioned energy-policy selection before Base, Stretch, and Specialized code paths can be compiled safely. |
| New York | Research only | Requires explicit geographic exceptions for New York City and provision-level legal-status overrides. |
| California | Research only | Requires occupancy and state-agency routing plus a separately refreshed local-amendment registry. |

## Implemented profiles

The human-reviewable source artifacts are:

- `reports/pilots/virginia.md`
- `reports/pilots/oregon.md`
- `reports/pilots/north-carolina.md`

The compiler projects those reports into `backend/data/regulatory/*.json`. Executable rules, claims, source-health observations, and fixture identifiers live in `backend/data/regulatory/rules/*.json`.

### Virginia

The pilot records the Board of Housing and Community Development as the statewide adopting authority, the 2021 Virginia Construction Code and Existing Building Code, local enforcement candidates, statewide preemption within the USBC scope, and the one-year prior-edition election period. Trade codes, operational fire authority, specialized project paths, and complete historical editions remain unresolved.

### Oregon

The pilot records current building, existing-building, commercial-energy, residential, electrical, plumbing, and mechanical specialty-code paths. Optional and mandatory dates are represented separately. The anticipated 2026 Oregon Residential Specialty Code remains pending and is not returned as current. Local program boundaries, interim amendments, alternate methods, and prior-edition records remain incomplete.

### North Carolina

The pilot records the current 2018 building and residential codes, the state code councils, State Fire Marshal supervision, and local inspection candidates. The 2024 code remains pending because its implementation depends on a certification event and statutory delay. No calendar date is inferred. The resolver continues to return the 2018 baseline and requires the current certification record.

## Deferred resolver capabilities

### Municipality-conditioned policy

Massachusetts energy-code resolution depends on maintained municipality-level policy and effective dates. A state-wide default plus an unconstrained local override would be unsafe.

### Geographic exceptions and provision status

New York requires an explicit New York City exclusion from the state Uniform Code and a way to preserve provision-level suspensions or other legal-status changes without invalidating an entire adoption.

### Multi-agency and occupancy routing

California requires project and occupancy routing among state agencies and local enforcing authorities, plus a versioned local-ordinance registry. This is a dedicated subsystem, not another flat state profile.

## Cross-state requirements exposed

1. Geography-conditioned policies for named municipalities or maintained policy layers.
2. Conditional future dates based on official certifications or events.
3. Provision-level suspension, conflict, and supersession states.
4. Edition-specific local-amendment registries.
5. Multi-agency project and occupancy jurisdiction.
6. Independent effective, optional, mandatory, and replacement dates.

## Next work

- Expand the three executable pilots only through source-backed code-family, authority, amendment, and fixture increments.
- Add prior-edition records before claiming complete historical transition results.
- Design and test municipality-conditioned policy selection before promoting Massachusetts.
- Design geographic exceptions and provision-level status before promoting New York.
- Design California occupancy routing and local-amendment ingestion as a bounded subsystem before promoting California.
