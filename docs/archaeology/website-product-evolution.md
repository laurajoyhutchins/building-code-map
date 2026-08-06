# Website and product evolution

## From explorer to public lookup

The technical GIS explorer proved that local boundary data could be displayed and inspected, but it was not the right public front door. The product evolved into a search-first address or coordinate lookup while retaining `/explorer` as the technical workbench.

## Evidence-bearing result design

The public result is designed to show why a conclusion was produced rather than decorating it with generic trust claims. It can expose:

- original and matched location;
- geocoder precision and source provenance;
- boundary observations and ambiguity;
- candidate authorities and relationships;
- adopted instruments and temporal status;
- verification state;
- required local confirmation;
- warnings and dated sources.

Generic trust badges and slogan-like assurance copy were removed. Credibility is carried by inspectable evidence and bounded uncertainty.

## Current request flow

The public client sends:

- an address to `POST /lookup`, which geocodes locally and resolves the resulting point; or
- coordinates to the point-only `POST /resolve` path.

The client does not supply normalized jurisdiction context as trusted regulatory input. Geographic context is derived by the server from its validated snapshot.

## Degraded and ambiguous operation

The backend now reports capability-specific readiness. A missing geocoder does not necessarily disable coordinate inspection. Missing regulatory profiles can leave geographic exploration available without producing code conclusions.

Boundary ambiguity is an explicit response, not an invisible first-match choice. Frontend runtime decoding and richer degraded or ambiguous evidence rendering remain follow-up work.

## Current product truth

The website can demonstrate six executable pilot profiles and three production-ready declared scopes. It must not present these as national completeness or complete statewide legal coverage.

The product should continue to favor precise language:

- “candidate authority” over “the authority” when evidence is incomplete;
- “declared production-ready scope” over “state complete”;
- “address point” or “interpolated” over “exact location”;
- “current UTC date supplied” over silently implying a project date;
- “local record required” over filling a gap with a likely answer.

## Deferred product capabilities

- nationwide jurisdiction and coverage ledger;
- runtime decoding of every API payload;
- historical location and regulatory resolution;
- project-type intake sufficient for specialized agency routing;
- comprehensive local amendment browsing;
- source-change notifications and stale-result warnings;
- publication-grade deployment backed by hydrated, versioned production snapshots.
