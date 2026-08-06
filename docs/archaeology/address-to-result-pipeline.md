# Address-to-result pipeline

## 1. Preserve the input

The client retains the entered address or coordinate. A coordinate remains a coordinate; it is not converted into a synthetic address.

## 2. Normalize a civic address

Normalization is deterministic and conservative. It standardizes whitespace, punctuation, bounded street suffixes, directions, state names, and ZIP forms while rejecting blank, incomplete, PO-box, and malformed inputs. The original query and normalized comparison fields remain distinct.

## 3. Geocode locally

The SQLite geocoder prefers qualifying address points. Street-range interpolation is used only when no address point qualifies, and only with range containment and parity checks. Candidate ordering is deterministic. Materially tied candidates return ambiguity instead of arbitrary selection.

## 4. Preserve point provenance

A selected candidate includes coordinates, matched address, precision, confidence, source name, source record identifier, and source vintage. `address_point` and `interpolated` are different precision classes. Neither independently proves parcel, rooftop, entrance, or unit accuracy.

## 5. Derive geographic observations

The server tests the point against an admitted boundary snapshot. It may observe state, county, municipality, incorporated status, special areas, American Indian areas, and NERIS jurisdictions.

The observations are not caller-authored trusted context and are not yet legal conclusions. Multiple peer state, county, or municipality matches produce `409 boundary_ambiguous` with all tied evidence.

## 6. Interpret regulatory policy

The resolver selects the compatible state profile, applies date and project-scope rules, expands authority candidates, resolves adoption records, and preserves source and verification references. Special, tribal, and fire-service observations remain contextual unless a policy rule supplies the legal relationship.

## 7. Apply the as-of date

Publication, adoption, effective, mandatory, and repeal dates remain distinct. Pending instruments do not become current merely because their edition is newer. An omitted date is visibly defaulted to the current UTC date for compatibility.

## 8. Emit inspectable outputs

`POST /lookup` returns geocoding and regulatory results. `POST /resolve` is point-only. Results expose location evidence, candidate authorities, adoptions, special conditions, local confirmation items, warnings, jurisdiction structure, and sources.

## Failure boundaries

- invalid address: request rejected;
- geocoder unavailable: address lookup unavailable while coordinate paths may remain usable;
- no candidate: `not_found`;
- tied geocoder candidates: `ambiguous`;
- tied peer boundary observations: `boundary_ambiguous`;
- unsupported geography or missing profile: insufficient evidence;
- missing local adoption, amendment, or enforcement record: local record required;
- conflicting source claims: conflict preserved;
- invalid snapshot: capability not admitted.

Each boundary prevents a weaker stage from manufacturing confidence for the next stage.
