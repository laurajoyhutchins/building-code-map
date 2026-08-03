---
state:
  state_id: "US-VA"
  name: "Virginia"
  abbreviation: "VA"
report:
  report_id: "state-report:usa-va"
  schema_version: "1.0"
  status: "draft"
  last_updated: "2026-08-02"
  last_verified: "2026-08-02"
risk:
  overall_confidence: 0.78
  risk_flags: ["draft_scope", "code_families_incomplete", "special_projects_unresolved"]
  open_questions_count: 5
---

# Building Code Authority Report: Virginia

Virginia provides a strong statewide-code and local-enforcement model. This draft establishes the construction and existing-building paths without yet claiming complete coverage of electrical, plumbing, mechanical, energy, fire prevention, elevators, industrialized buildings, or state-owned projects.

## Supported findings

- The Board of Housing and Community Development adopts and amends the Virginia Uniform Statewide Building Code.
- The 2021 Virginia codes became effective on 2024-01-18.
- The Virginia Construction Code incorporates the 2021 International Building Code with Virginia amendments.
- The Virginia Existing Building Code incorporates the 2021 International Existing Building Code with Virginia amendments.
- Local building departments generally enforce construction and rehabilitation requirements.
- State law broadly supersedes local building codes, subject to statutory exceptions.
- Permit applications submitted during the one-year transition beginning on the 2021 code effective date may qualify to use the prior edition. The rule pack and fixtures still need to encode and test that transition.

## Canonical Jurisdiction Profile

```json jurisdiction-profile
{
  "schema_version": "1.0",
  "profile_id": "state-profile:us-va",
  "state_id": "US-VA",
  "state_name": "Virginia",
  "state_abbreviation": "VA",
  "state_fips": "51",
  "status": "draft",
  "last_verified": "2026-08-02",
  "sources": [
    {
      "id": "src:us-va:dhcd-codes",
      "title": "Virginia DHCD — Codes",
      "url": "https://www.dhcd.virginia.gov/codes",
      "kind": "official_web",
      "accessed_at": "2026-08-02"
    },
    {
      "id": "src:us-va:dhcd-usbc",
      "title": "Virginia DHCD — Uniform Statewide Building Code",
      "url": "https://www.dhcd.virginia.gov/virginia-uniform-statewide-building-code-usbc",
      "kind": "official_web",
      "accessed_at": "2026-08-02"
    },
    {
      "id": "src:us-va:code-36-98",
      "title": "Code of Virginia § 36-98",
      "url": "https://law.lis.virginia.gov/vacode/title36/chapter6/section36-98/",
      "kind": "statute",
      "accessed_at": "2026-08-02"
    },
    {
      "id": "src:us-va:code-36-105",
      "title": "Code of Virginia § 36-105",
      "url": "https://law.lis.virginia.gov/vacode/title36/chapter6/section36-105/",
      "kind": "statute",
      "accessed_at": "2026-08-02"
    },
    {
      "id": "src:us-va:vac-13-5-63-10",
      "title": "13VAC5-63-10 — Virginia Construction Code",
      "url": "https://law.lis.virginia.gov/admincode/title13/agency5/chapter63/section10/",
      "kind": "administrative_code",
      "accessed_at": "2026-08-02"
    },
    {
      "id": "src:us-va:vac-13-5-63-30",
      "title": "13VAC5-63-30 — Application of code",
      "url": "https://law.lis.virginia.gov/admincode/title13/agency5/chapter63/section30/",
      "kind": "administrative_code",
      "accessed_at": "2026-08-02"
    },
    {
      "id": "src:us-va:vac-13-5-63-400",
      "title": "13VAC5-63-400 — Virginia Existing Building Code",
      "url": "https://law.lis.virginia.gov/admincode/title13/agency5/chapter63/section400/",
      "kind": "administrative_code",
      "accessed_at": "2026-08-02"
    }
  ],
  "authorities": [
    {
      "id": "auth:us-va:board-housing-community-development",
      "name": "Virginia Board of Housing and Community Development",
      "type": "state_board",
      "roles": ["adopts", "amends", "maintains_statewide_code"],
      "source_ids": ["src:us-va:dhcd-usbc", "src:us-va:code-36-98"],
      "verification": {
        "status": "verified",
        "confidence": 0.95
      }
    }
  ],
  "relationships": [
    {
      "id": "relationship:us-va:state-code-preempts-local-building-codes",
      "from_id": "auth:us-va:board-housing-community-development",
      "relationship": "preempts",
      "to": "county, municipal, political-subdivision, and state-agency building codes within USBC scope",
      "scope": ["construction_code"],
      "summary": "The Uniform Statewide Building Code supersedes local and other state-agency building codes within its statutory scope, subject to stated exceptions.",
      "source_ids": ["src:us-va:code-36-98"],
      "verification": {
        "status": "verified",
        "confidence": 0.95
      }
    }
  ],
  "adoptions": [
    {
      "id": "adoption:us-va:building:2021-vcc",
      "code_family": "building",
      "status": "current",
      "state_code_name": "2021 Virginia Construction Code",
      "base_model_code": {
        "publisher": "ICC",
        "name": "International Building Code",
        "edition_year": 2021
      },
      "adopting_authority_id": "auth:us-va:board-housing-community-development",
      "enforcement_model": "local",
      "applies_to": ["new_construction", "additions", "construction_activities"],
      "dates": {
        "effective_date": "2024-01-18",
        "operative_date": "2024-01-18"
      },
      "source_ids": ["src:us-va:dhcd-codes", "src:us-va:vac-13-5-63-10"],
      "verification": {
        "status": "partially_verified",
        "confidence": 0.9
      }
    },
    {
      "id": "adoption:us-va:existing-building:2021-vebc",
      "code_family": "existing_building",
      "status": "current",
      "state_code_name": "2021 Virginia Existing Building Code",
      "base_model_code": {
        "publisher": "ICC",
        "name": "International Existing Building Code",
        "edition_year": 2021
      },
      "adopting_authority_id": "auth:us-va:board-housing-community-development",
      "enforcement_model": "local",
      "applies_to": ["alteration", "repair", "rehabilitation", "change_of_occupancy"],
      "dates": {
        "effective_date": "2024-01-18",
        "operative_date": "2024-01-18"
      },
      "source_ids": ["src:us-va:dhcd-codes", "src:us-va:vac-13-5-63-400"],
      "verification": {
        "status": "partially_verified",
        "confidence": 0.9
      }
    }
  ],
  "defaults": {
    "incorporated": {
      "status": "partially_resolved",
      "authority_candidates": [
        {
          "kind": "state_authority",
          "authority_id": "auth:us-va:board-housing-community-development",
          "roles": ["adopts", "amends"],
          "source_ids": ["src:us-va:dhcd-usbc", "src:us-va:code-36-98"]
        },
        {
          "kind": "municipality",
          "label": "Local building department",
          "roles": ["administers", "enforces", "permits", "inspects"],
          "source_ids": ["src:us-va:code-36-105"]
        }
      ],
      "required_local_records": [
        "Identity and service area of the enforcing local building department",
        "Permit application date and any prior-edition election",
        "Applicable local administrative procedures that do not conflict with the USBC"
      ],
      "warnings": [
        "This draft does not yet resolve all trade, fire-prevention, elevator, industrialized-building, or special-project authorities."
      ],
      "source_ids": ["src:us-va:code-36-98", "src:us-va:code-36-105"]
    },
    "unincorporated": {
      "status": "partially_resolved",
      "authority_candidates": [
        {
          "kind": "state_authority",
          "authority_id": "auth:us-va:board-housing-community-development",
          "roles": ["adopts", "amends"],
          "source_ids": ["src:us-va:dhcd-usbc", "src:us-va:code-36-98"]
        },
        {
          "kind": "county",
          "label": "County or contracted building department",
          "roles": ["administers", "enforces", "permits", "inspects"],
          "source_ids": ["src:us-va:code-36-105"]
        }
      ],
      "required_local_records": [
        "Identity and service area of the enforcing county or contracted building department",
        "Permit application date and any prior-edition election",
        "Applicable local administrative procedures that do not conflict with the USBC"
      ],
      "warnings": [
        "This draft does not yet resolve all trade, fire-prevention, elevator, industrialized-building, or special-project authorities."
      ],
      "source_ids": ["src:us-va:code-36-98", "src:us-va:code-36-105"]
    }
  },
  "code_family_overrides": {
    "building": {
      "status": "partially_resolved",
      "adoption_ids": ["adoption:us-va:building:2021-vcc"],
      "required_local_records": ["Permit application date and prior-edition election, if any"],
      "source_ids": ["src:us-va:dhcd-codes", "src:us-va:vac-13-5-63-10", "src:us-va:vac-13-5-63-30"]
    },
    "existing_building": {
      "status": "partially_resolved",
      "adoption_ids": ["adoption:us-va:existing-building:2021-vebc"],
      "required_local_records": ["Project scope and permit application date"],
      "source_ids": ["src:us-va:dhcd-codes", "src:us-va:vac-13-5-63-400"]
    }
  },
  "verification": {
    "status": "needs_review",
    "confidence": 0.78,
    "notes": "Construction and existing-building authority paths are source-backed; remaining code families and special-project paths are not yet complete."
  }
}
```

## Open questions

1. What current Virginia regulations incorporate the electrical, plumbing, mechanical, fuel-gas, energy, and operational fire codes?
2. Which authorities administer state-owned buildings, public schools, industrialized buildings, elevators, and amusement devices?
3. Does any locality retain a legally meaningful construction-code variation within the USBC framework beyond administrative procedures and statutory exceptions?
4. How should small-town county enforcement be represented in geographic fixtures?
5. What current amendments, errata, and interpretations must be tracked as source-health records?
