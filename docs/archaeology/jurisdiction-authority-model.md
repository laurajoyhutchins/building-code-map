# Jurisdiction and authority model

## Geographic containment

The boundary snapshot can supply containment matches for:

- state;
- county;
- municipality;
- special area;
- tribal area;
- NERIS fire jurisdiction.

The broader product model anticipates township, borough, utility, flood/NFIP, campus, airport,
federal, and other special jurisdictions where reliable geometry and legal evidence exist. Their
presence in the conceptual model does not mean current nationwide datasets are included.

## Legal authority

Authority records answer different questions:

- who adopted the rule;
- who administers the program;
- who performs plan review or permitting;
- who enforces or inspects;
- whether authority was delegated, reserved, preempted, or shared;
- whether a code family or project class changes the answer.

A polygon may identify a relevant entity without proving any of those roles.

## Why a graph is needed

A flat `city -> edition` table fails when:

- the state adopts a mandatory minimum but a municipality enforces it;
- electrical authority differs from building authority;
- fire prevention is operationally separate from construction subcodes;
- a county governs unincorporated land but not a city;
- a state agency governs public schools or state-owned facilities;
- local amendments are allowed only within bounded scope;
- two official sources conflict;
- a boundary is known but the adopting instrument is missing.

The state profile therefore models authorities, roles, relationships, adoptions, policy rules,
claims, sources, and verification state. The result emits candidate authorities and an
`authority_path`.

## Current contract

`schemas/regulatory/resolution-result.schema.json` fixes `schema_version` to `1.0`. It includes:

- normalized geography;
- resolution status;
- authority candidates;
- authority path relationships;
- adoption records;
- applicable rules;
- supporting claims;
- required local records;
- warnings;
- evidence;
- policy basis.

No separate current schema named `authority_graph` was found on `main`. “Authority graph” is the
domain model expressed through the state profile and result relationship structures.

## Current containment behavior and risks

- polygon holes are respected;
- points on a segment count as contained;
- the first county and municipality match are retained;
- special, tribal, and fire matches accumulate;
- geometry is not historically versioned;
- no explicit boundary-dispute object exists;
- legal authority still depends on policy data after containment;
- missing or conflicting geometry can make context incomplete before policy resolution.

Temporal boundary changes, point-on-boundary policy, overlapping municipalities, and source conflict
handling require stronger data and explicit fixtures before being described as complete.
