---
state:
  state_id: "US-CO"
  name: "Colorado"
  abbreviation: "CO"
report:
  report_id: "pilot-report:us-co"
  schema_version: "1.0"
  status: "partially_verified"
  last_updated: "2026-07-20"
  last_verified: "2026-07-20"
risk:
  overall_confidence: 0.80
  risk_flags: ["pilot_scope", "local_records_required"]
  open_questions_count: 1
---

# Building Code Authority Pilot: Colorado

Predominantly local adoption with statewide electrical authority.

This pilot is intentionally scoped to resolver behavior. It preserves unresolved local records and does not claim production-ready statewide or municipal coverage.

## Canonical Jurisdiction Profile

```json jurisdiction-profile
{"schema_version":"1.0","profile_id":"state-profile:us-co","state_id":"US-CO","state_name":"Colorado","state_abbreviation":"CO","state_fips":"08","status":"partially_verified","last_verified":"2026-07-20","sources":[{"id":"src:us-co:energy-local","title":"Colorado General Assembly — local energy-code adoption requirements","url":"https://leg.colorado.gov/bills/HB23-1085","kind":"session_law_summary","accessed_at":"2026-07-20"},{"id":"src:us-co:electrical-board","title":"Colorado State Electrical Board — permit and inspection information","url":"https://dpo.colorado.gov/Electrical/PermitInspectionInfo","kind":"official_web","accessed_at":"2026-07-20"},{"id":"src:us-co:accessibility-local","title":"Colorado General Assembly — accessibility standards in locally adopted building codes","url":"https://leg.colorado.gov/bills/HB25-1030","kind":"session_law_summary","accessed_at":"2026-07-20"}],"authorities":[{"id":"auth:us-co:state-electrical-board","name":"Colorado State Electrical Board","type":"state_board","roles":["adopts","inspects_when_no_local_authority","inspects_state_buildings","inspects_public_schools"],"source_ids":["src:us-co:electrical-board"],"verification":{"status":"verified","confidence":0.95}}],"relationships":[],"adoptions":[{"id":"adoption:us-co:electrical:2020-nec","code_family":"electrical","status":"current","state_code_name":"Colorado Electrical Code","base_model_code":{"publisher":"NFPA","name":"NFPA 70 National Electrical Code","edition_year":2020},"adopting_authority_id":"auth:us-co:state-electrical-board","enforcement_model":"hybrid","applies_to":["electrical_installations"],"dates":{},"source_ids":["src:us-co:electrical-board"],"verification":{"status":"partially_verified","confidence":0.8}}],"defaults":{"incorporated":{"status":"local_record_required","authority_candidates":[{"kind":"municipality","source_ids":["src:us-co:accessibility-local","src:us-co:energy-local"],"label":"building department","roles":["adopts","enforces"]},{"kind":"fire_authority","source_ids":["src:us-co:accessibility-local"],"label":"Local fire authority","roles":["adopts_or_enforces"]}],"required_local_records":["Current municipal building-code adoption ordinance and effective date","Current municipal amendments","Current fire-code adoption and enforcing authority"],"warnings":["Colorado does not provide one general statewide building-code edition for ordinary local construction."],"source_ids":["src:us-co:accessibility-local","src:us-co:energy-local"]},"unincorporated":{"status":"local_record_required","authority_candidates":[{"kind":"county","source_ids":["src:us-co:accessibility-local","src:us-co:energy-local"],"label":"building department","roles":["adopts","enforces"]},{"kind":"fire_authority","source_ids":["src:us-co:accessibility-local"],"label":"Local fire authority","roles":["adopts_or_enforces"]}],"required_local_records":["Current county building-code adoption resolution and effective date","Current county amendments","Current fire-code adoption and enforcing authority"],"warnings":["Colorado does not provide one general statewide building-code edition for ordinary local construction."],"source_ids":["src:us-co:accessibility-local","src:us-co:energy-local"]}},"code_family_overrides":{"electrical":{"status":"partially_resolved","authority_candidates":[{"kind":"state_authority","source_ids":["src:us-co:electrical-board"],"authority_id":"auth:us-co:state-electrical-board"}],"adoption_ids":["adoption:us-co:electrical:2020-nec"],"required_local_records":["Whether the location is inspected by a local electrical authority or the State Electrical Board"],"source_ids":["src:us-co:electrical-board"]},"energy":{"status":"local_record_required","required_local_records":["Current locally adopted residential and commercial energy-code editions and effective dates"],"source_ids":["src:us-co:energy-local"]}},"project_type_overrides":{"state_owned":{"status":"partially_resolved","authority_candidates":[{"kind":"state_authority","source_ids":["src:us-co:electrical-board"],"authority_id":"auth:us-co:state-electrical-board"}],"required_local_records":["Controlling state agency building-code and plan-review requirements for the project"],"source_ids":["src:us-co:electrical-board"]},"public_school":{"status":"partially_resolved","authority_candidates":[{"kind":"state_authority","source_ids":["src:us-co:electrical-board"],"authority_id":"auth:us-co:state-electrical-board"}],"required_local_records":["Controlling school-construction and local building-code requirements"],"source_ids":["src:us-co:electrical-board"]}},"verification":{"status":"partially_verified","confidence":0.8}}
```
