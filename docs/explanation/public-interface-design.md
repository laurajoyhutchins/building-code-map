# Public interface design

Building Code Map has two distinct interfaces.

The public interface answers one question: which authorities and adopted building-code records are associated with this location? It begins with a location lookup, presents the answer before supporting detail, and uses ordinary regulatory language. It does not display runtime telemetry, cache counts, API routes, dataset branding, trust slogans, or implementation terminology.

The explorer remains the technical inspection surface. It exposes boundary layers, cached records, source attributes, refresh state, and backend contracts for maintainers and advanced users.

## Public interface

The public route is `/`.

The initial view contains the product name, one sentence describing the task, a coordinate lookup, code-family and applicability-date controls, and a link to the explorer. Address lookup is not presented until a supported geocoder exists.

After lookup, the page presents:

1. matched place;
2. authority names;
3. supported adoption records and effective dates;
4. a plain-language notice when the result is incomplete, ambiguous, conflicting, or requires local records;
5. special conditions and records to confirm locally;
6. links to the supporting sources.

Resolved results receive no decorative status badge. Other states are explained in sentences.

## Explorer

The technical interface remains available at `/explorer`. Its existing map, boundary layers, record inspector, refresh state, and API contract are retained.

## Language

Public labels use `Authorities`, `Adopted codes`, `Special conditions`, `Confirm locally`, `Jurisdiction structure`, and `Sources`.

The public route does not use promotional trust claims or internal labels such as `policy basis`, `supporting claims`, `mirror healthy`, `cached click targets`, or `backend contract`.

## Visual system

The existing restrained palette and typography remain. The public route uses an open layout with a single lookup region and a single result column. It avoids feature grids, metric cards, pills, glows, decorative icons, and nested dashboard panels.

## Failure behavior

Invalid coordinates are rejected before a request. API failures appear beside the lookup without clearing the entered values. Incomplete regulatory results preserve uncertainty and direct the user to the specific records that still require confirmation.
