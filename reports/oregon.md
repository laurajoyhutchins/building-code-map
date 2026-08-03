---
state:
  state_id: "US-OR"
  name: "Oregon"
  abbreviation: "OR"
report:
  report_id: "state-report:usa-or"
  schema_version: "1.0"
  status: "draft"
  last_updated: "2026-08-02"
  last_verified: "2026-08-02"
risk:
  overall_confidence: 0.82
  risk_flags: ["draft_scope", "local_authority_unresolved", "interim_amendments_not_normalized"]
  open_questions_count: 5
---

# Building Code Authority Report: Oregon

Oregon publishes unusually strong current-edition and transition data by specialty-code family. This draft captures the primary statewide adoptions and date boundaries but does not yet claim complete local enforcement, amendment, special-project, or interim-amendment coverage.

## Supported findings

- The Oregon Building Codes Division adopts, amends, and interprets specialty codes that compose the Oregon State Building Code.
- The 2025 Oregon Structural Specialty Code is based on the 2024 IBC, IFC, and IEBC. Its construction provisions became effective on 2025-10-01 and mandatory on 2026-04-01 after a six-month phase-in.
- The 2025 Oregon Energy Efficiency Specialty Code became effective on 2025-01-01 and mandatory on 2025-07-01; it is based on ASHRAE 90.1-2022.
- The current residential code is the 2023 Oregon Residential Specialty Code, with construction provisions mandatory from 2024-04-01.
- The 2023 Oregon Electrical Specialty Code is based on the 2023 NEC.
- The 2023 Oregon Plumbing Specialty Code is based on the 2021 UPC.
- A 2026 Oregon Residential Specialty Code adoption process is underway. Its anticipated adoption date is not treated as an operative or mandatory date in this draft.

## Canonical Jurisdiction Profile

```json jurisdiction-profile
{
  "schema_version": "1.0",
  "profile_id": "state-profile:us-or",
  "state_id": "US-OR",
  "state_name": "Oregon",
  "state_abbreviation": "OR",
  "state_fips": "41",
  "status": "draft",
  "last_verified": "2026-08-02",
  "sources": [
    {
      "id": "src:us-or:bcd-state-code",
      "title": "Oregon Building Codes Division — Oregon State Building Code",
      "url": "https://www.oregon.gov/bcd/codes-stand/pages/index.aspx",
      "kind": "official_web",
      "accessed_at": "2026-08-02"
    },
    {
      "id": "src:us-or:bcd-adopted-codes",
      "title": "Oregon Building Codes Division — Adopted codes online",
      "url": "https://www.oregon.gov/bcd/codes-stand/pages/adopted-codes.aspx",
      "kind": "official_web",
      "accessed_at": "2026-08-02"
    },
    {
      "id": "src:us-or:bcd-ossc-2025",
      "title": "Oregon Building Codes Division — 2025 OSSC adoption",
      "url": "https://www.oregon.gov/bcd/codes-stand/Pages/ossc-adoption.aspx",
      "kind": "official_web",
      "accessed_at": "2026-08-02"
    },
    {
      "id": "src:us-or:bcd-oeesc-2025",
      "title": "Oregon Building Codes Division — 2025 OEESC",
      "url": "https://www.oregon.gov/bcd/codes-stand/Pages/oeesc-adoption.aspx",
      "kind": "official_web",
      "accessed_at": "2026-08-02"
    },
    {
      "id": "src:us-or:bcd-orsc-2026-process",
      "title": "Oregon Building Codes Division — 2026 ORSC adoption process",
      "url": "https://www.oregon.gov/bcd/codes-stand/Pages/orsc-adoption.aspx",
      "kind": "official_web",
      "accessed_at": "2026-08-02"
    },
    {
      "id": "src:us-or:bcd-code-history",
      "title": "Oregon Building Codes Division — Codebook history",
      "url": "https://www.oregon.gov/bcd/codes-stand/pages/codebook-history.aspx",
      "kind": "official_web",
      "accessed_at": "2026-08-02"
    }
  ],
  "authorities": [
    {
      "id": "auth:us-or:building-codes-division",
      "name": "Oregon Building Codes Division",
      "type": "state_agency",
      "roles": ["adopts", "amends", "interprets", "administers_specialty_code_programs"],
      "source_ids": ["src:us-or:bcd-state-code", "src:us-or:bcd-adopted-codes"],
      "verification": {
        "status": "verified",
        "confidence": 0.95
      }
    }
  ],
  "relationships": [],
  "adoptions": [
    {
      "id": "adoption:us-or:building:2025-ossc",
      "code_family": "building",
      "status": "current",
      "state_code_name": "2025 Oregon Structural Specialty Code",
      "base_model_code": {
        "publisher": "ICC",
        "name": "International Building Code",
        "edition_year": 2024
      },
      "adopting_authority_id": "auth:us-or:building-codes-division",
      "enforcement_model": "local_with_state_oversight",
      "applies_to": ["commercial", "new_construction", "alteration", "repair"],
      "dates": {
        "effective_date": "2025-10-01",
        "operative_date": "2025-10-01",
        "mandatory_date": "2026-04-01"
      },
      "source_ids": ["src:us-or:bcd-ossc-2025", "src:us-or:bcd-adopted-codes"],
      "verification": {
        "status": "partially_verified",
        "confidence": 0.92
      }
    },
    {
      "id": "adoption:us-or:existing-building:2025-ossc",
      "code_family": "existing_building",
      "status": "current",
      "state_code_name": "2025 Oregon Structural Specialty Code — existing-building provisions",
      "base_model_code": {
        "publisher": "ICC",
        "name": "International Existing Building Code",
        "edition_year": 2024
      },
      "adopting_authority_id": "auth:us-or:building-codes-division",
      "enforcement_model": "local_with_state_oversight",
      "applies_to": ["existing_buildings", "alteration", "repair"],
      "dates": {
        "effective_date": "2025-10-01",
        "operative_date": "2025-10-01",
        "mandatory_date": "2026-04-01"
      },
      "source_ids": ["src:us-or:bcd-ossc-2025", "src:us-or:bcd-adopted-codes"],
      "verification": {
        "status": "partially_verified",
        "confidence": 0.9
      }
    },
    {
      "id": "adoption:us-or:energy:2025-oeesc",
      "code_family": "energy",
      "status": "current",
      "state_code_name": "2025 Oregon Energy Efficiency Specialty Code",
      "base_model_code": {
        "publisher": "ASHRAE",
        "name": "ANSI/ASHRAE/IES Standard 90.1",
        "edition_year": 2022
      },
      "adopting_authority_id": "auth:us-or:building-codes-division",
      "enforcement_model": "local_with_state_oversight",
      "applies_to": ["commercial_energy"],
      "dates": {
        "effective_date": "2025-01-01",
        "operative_date": "2025-01-01",
        "mandatory_date": "2025-07-01"
      },
      "source_ids": ["src:us-or:bcd-oeesc-2025", "src:us-or:bcd-adopted-codes"],
      "verification": {
        "status": "verified",
        "confidence": 0.95
      }
    },
    {
      "id": "adoption:us-or:residential:2023-orsc",
      "code_family": "residential",
      "status": "current",
      "state_code_name": "2023 Oregon Residential Specialty Code",
      "adopting_authority_id": "auth:us-or:building-codes-division",
      "enforcement_model": "local_with_state_oversight",
      "applies_to": ["detached_one_two_family", "townhouses"],
      "dates": {
        "effective_date": "2023-10-01",
        "operative_date": "2023-10-01",
        "mandatory_date": "2024-04-01"
      },
      "source_ids": ["src:us-or:bcd-adopted-codes"],
      "verification": {
        "status": "partially_verified",
        "confidence": 0.88
      }
    },
    {
      "id": "adoption:us-or:electrical:2023-oesc",
      "code_family": "electrical",
      "status": "current",
      "state_code_name": "2023 Oregon Electrical Specialty Code",
      "base_model_code": {
        "publisher": "NFPA",
        "name": "NFPA 70 National Electrical Code",
        "edition_year": 2023
      },
      "adopting_authority_id": "auth:us-or:building-codes-division",
      "enforcement_model": "local_with_state_oversight",
      "applies_to": ["electrical_installations"],
      "dates": {
        "effective_date": "2023-10-01",
        "operative_date": "2023-10-01"
      },
      "source_ids": ["src:us-or:bcd-adopted-codes"],
      "verification": {
        "status": "partially_verified",
        "confidence": 0.9
      }
    },
    {
      "id": "adoption:us-or:plumbing:2023-opsc",
      "code_family": "plumbing",
      "status": "current",
      "state_code_name": "2023 Oregon Plumbing Specialty Code",
      "base_model_code": {
        "publisher": "IAPMO",
        "name": "Uniform Plumbing Code",
        "edition_year": 2021
      },
      "adopting_authority_id": "auth:us-or:building-codes-division",
      "enforcement_model": "local_with_state_oversight",
      "applies_to": ["plumbing_installations"],
      "dates": {
        "effective_date": "2023-10-01",
        "operative_date": "2023-10-01"
      },
      "source_ids": ["src:us-or:bcd-adopted-codes"],
      "verification": {
        "status": "partially_verified",
        "confidence": 0.88
      }
    },
    {
      "id": "adoption:us-or:residential:2026-orsc-pending",
      "code_family": "residential",
      "status": "pending",
      "state_code_name": "2026 Oregon Residential Specialty Code",
      "base_model_code": {
        "publisher": "ICC",
        "name": "International Residential Code",
        "edition_year": 2024
      },
      "adopting_authority_id": "auth:us-or:building-codes-division",
      "enforcement_model": "local_with_state_oversight",
      "applies_to": ["detached_one_two_family", "townhouses"],
      "dates": {},
      "source_ids": ["src:us-or:bcd-orsc-2026-process"],
      "verification": {
        "status": "needs_review",
        "confidence": 0.75,
        "notes": "The state page describes an anticipated adoption process; no operative or mandatory date is asserted here."
      }
    }
  ],
  "defaults": {
    "incorporated": {
      "status": "partially_resolved",
      "authority_candidates": [
        {
          "kind": "state_authority",
          "authority_id": "auth:us-or:building-codes-division",
          "roles": ["adopts", "amends", "interprets"],
          "source_ids": ["src:us-or:bcd-state-code"]
        },
        {
          "kind": "municipality",
          "label": "Local building department or delegated program",
          "roles": ["permits", "inspects", "enforces"],
          "source_ids": ["src:us-or:bcd-state-code"]
        }
      ],
      "required_local_records": [
        "Identity and service area of the enforcing building program",
        "Permit application date and code-family-specific phase-in election",
        "Applicable interim amendments, errata, and statewide alternate methods"
      ],
      "warnings": [
        "Local enforcement delegation and local amendment authority remain to be verified from statutes and administrative rules."
      ],
      "source_ids": ["src:us-or:bcd-state-code", "src:us-or:bcd-adopted-codes"]
    },
    "unincorporated": {
      "status": "partially_resolved",
      "authority_candidates": [
        {
          "kind": "state_authority",
          "authority_id": "auth:us-or:building-codes-division",
          "roles": ["adopts", "amends", "interprets"],
          "source_ids": ["src:us-or:bcd-state-code"]
        },
        {
          "kind": "county",
          "label": "County building department or delegated program",
          "roles": ["permits", "inspects", "enforces"],
          "source_ids": ["src:us-or:bcd-state-code"]
        }
      ],
      "required_local_records": [
        "Identity and service area of the enforcing building program",
        "Permit application date and code-family-specific phase-in election",
        "Applicable interim amendments, errata, and statewide alternate methods"
      ],
      "warnings": [
        "Local enforcement delegation and local amendment authority remain to be verified from statutes and administrative rules."
      ],
      "source_ids": ["src:us-or:bcd-state-code", "src:us-or:bcd-adopted-codes"]
    }
  },
  "code_family_overrides": {
    "building": {
      "status": "partially_resolved",
      "adoption_ids": ["adoption:us-or:building:2025-ossc"],
      "required_local_records": ["Permit application date and any code-cycle election during a phase-in period"],
      "source_ids": ["src:us-or:bcd-ossc-2025", "src:us-or:bcd-adopted-codes"]
    },
    "existing_building": {
      "status": "partially_resolved",
      "adoption_ids": ["adoption:us-or:existing-building:2025-ossc"],
      "required_local_records": ["Project scope, permit application date, and applicable interim amendments"],
      "source_ids": ["src:us-or:bcd-ossc-2025", "src:us-or:bcd-adopted-codes"]
    },
    "energy": {
      "status": "partially_resolved",
      "adoption_ids": ["adoption:us-or:energy:2025-oeesc"],
      "required_local_records": ["Project classification and applicable statewide alternate methods"],
      "source_ids": ["src:us-or:bcd-oeesc-2025"]
    },
    "residential": {
      "status": "partially_resolved",
      "adoption_ids": ["adoption:us-or:residential:2023-orsc"],
      "required_local_records": ["Project classification and applicable interim amendments"],
      "warnings": ["The 2026 residential-code process is pending and is not returned as current."],
      "source_ids": ["src:us-or:bcd-adopted-codes", "src:us-or:bcd-orsc-2026-process"]
    },
    "electrical": {
      "status": "partially_resolved",
      "adoption_ids": ["adoption:us-or:electrical:2023-oesc"],
      "required_local_records": ["Current Oregon amendments and enforcing electrical program"],
      "source_ids": ["src:us-or:bcd-adopted-codes"]
    },
    "plumbing": {
      "status": "partially_resolved",
      "adoption_ids": ["adoption:us-or:plumbing:2023-opsc"],
      "required_local_records": ["Current Oregon amendments and enforcing plumbing program"],
      "source_ids": ["src:us-or:bcd-adopted-codes"]
    }
  },
  "verification": {
    "status": "needs_review",
    "confidence": 0.82,
    "notes": "Current code-family editions and major transition dates are source-backed; local authority, special-project, and amendment details remain incomplete."
  }
}
```

## Open questions

1. What statutes and rules define local program delegation, enforcement jurisdiction, and appeals?
2. To what extent may local governments amend specialty-code requirements?
3. Which state agencies retain authority over state-owned buildings, schools, healthcare, manufactured housing, elevators, and other special project classes?
4. How should interim amendments, errata, and statewide alternate methods be versioned and applied by date?
5. What are the complete current mechanical, fire-operational, accessibility, and elevator adoption records?
