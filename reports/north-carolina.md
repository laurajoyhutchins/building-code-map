---
state:
  state_id: "US-NC"
  name: "North Carolina"
  abbreviation: "NC"
report:
  report_id: "state-report:usa-nc"
  schema_version: "1.0"
  status: "draft"
  last_updated: "2026-08-02"
  last_verified: "2026-08-02"
risk:
  overall_confidence: 0.76
  risk_flags: ["draft_scope", "future_code_condition_unmodeled", "code_families_incomplete"]
  open_questions_count: 6
---

# Building Code Authority Report: North Carolina

North Carolina currently presents a stable statewide 2018-code baseline and an unusual future transition: implementation of the 2024 code is tied to an official certification and a twelve-month delay rather than a fixed calendar date. This draft preserves that uncertainty explicitly.

## Supported findings

- The Office of the State Fire Marshal lists the 2018 North Carolina codes as current and states that they became effective on 2019-01-01.
- The 2018 code family uses the 2015 International Codes with North Carolina amendments.
- Implementation of the 2024 North Carolina State Building Code was postponed. The published trigger is twelve months after the State Fire Marshal certifies publication and distribution of the completed code and constitution of the Residential Code Council.
- The North Carolina State Building Code applies throughout the state, subject to statutory local-code provisions and specialized allocations.
- The State Fire Marshal has general supervisory, administrative, and enforcement authority over major construction-code families.
- Local governments must provide inspection services through their own department, a joint department, a contract, county service, or state-arranged service.
- State buildings follow a separate Department of Administration enforcement path that is not yet compiled into this profile.

## Canonical Jurisdiction Profile

```json jurisdiction-profile
{
  "schema_version": "1.0",
  "profile_id": "state-profile:us-nc",
  "state_id": "US-NC",
  "state_name": "North Carolina",
  "state_abbreviation": "NC",
  "state_fips": "37",
  "status": "draft",
  "last_verified": "2026-08-02",
  "sources": [
    {
      "id": "src:us-nc:osfm-current-past-codes",
      "title": "North Carolina OSFM — Codes: Current and Past",
      "url": "https://www.ncosfm.gov/index.php/codes/codes-current-and-past",
      "kind": "official_web",
      "accessed_at": "2026-08-02"
    },
    {
      "id": "src:us-nc:osfm-2024-delay",
      "title": "North Carolina OSFM — 2024 State Building Code implementation delay",
      "url": "https://www.ncosfm.gov/news/press-releases/2025/04/07/north-carolina-delays-implementation-2024-state-building-code",
      "kind": "official_notice",
      "accessed_at": "2026-08-02"
    },
    {
      "id": "src:us-nc:gs-143-138",
      "title": "North Carolina General Statutes § 143-138",
      "url": "https://www.ncleg.gov/EnactedLegislation/Statutes/HTML/BySection/Chapter_143/GS_143-138.html",
      "kind": "statute",
      "accessed_at": "2026-08-02"
    },
    {
      "id": "src:us-nc:gs-143-139",
      "title": "North Carolina General Statutes § 143-139",
      "url": "https://www.ncleg.gov/enactedlegislation/statutes/html/bysection/chapter_143/gs_143-139.html",
      "kind": "statute",
      "accessed_at": "2026-08-02"
    },
    {
      "id": "src:us-nc:gs-160d-1102",
      "title": "North Carolina General Statutes § 160D-1102",
      "url": "https://www.ncleg.gov/EnactedLegislation/Statutes/HTML/BySection/Chapter_160D/GS_160D-1102.html",
      "kind": "statute",
      "accessed_at": "2026-08-02"
    }
  ],
  "authorities": [
    {
      "id": "auth:us-nc:state-code-councils",
      "name": "North Carolina Building Code Council and Residential Code Council",
      "type": "state_boards",
      "roles": ["adopt", "amend", "maintain_state_building_code"],
      "source_ids": ["src:us-nc:gs-143-138", "src:us-nc:gs-143-139"],
      "verification": {
        "status": "partially_verified",
        "confidence": 0.85
      }
    },
    {
      "id": "auth:us-nc:state-fire-marshal",
      "name": "North Carolina State Fire Marshal",
      "type": "state_official",
      "roles": ["supervises", "administers", "enforces", "arranges_inspection_service_when_local_service_fails"],
      "source_ids": ["src:us-nc:gs-143-139", "src:us-nc:gs-160d-1102"],
      "verification": {
        "status": "verified",
        "confidence": 0.95
      }
    }
  ],
  "relationships": [
    {
      "id": "relationship:us-nc:state-fire-marshal-oversees-local-enforcement",
      "from_id": "auth:us-nc:state-fire-marshal",
      "relationship": "supervises",
      "to": "local building-code inspection and enforcement programs",
      "scope": ["construction_code"],
      "summary": "The State Fire Marshal supervises and administers statewide code enforcement while local governments ordinarily provide inspection services.",
      "source_ids": ["src:us-nc:gs-143-139", "src:us-nc:gs-160d-1102"],
      "verification": {
        "status": "verified",
        "confidence": 0.92
      }
    }
  ],
  "adoptions": [
    {
      "id": "adoption:us-nc:building:2018-code",
      "code_family": "building",
      "status": "current",
      "state_code_name": "2018 North Carolina State Building Code: Building Code",
      "base_model_code": {
        "publisher": "ICC",
        "name": "International Building Code",
        "edition_year": 2015
      },
      "adopting_authority_id": "auth:us-nc:state-code-councils",
      "enforcement_model": "local_with_state_supervision",
      "applies_to": ["new_construction", "commercial", "alteration", "repair"],
      "dates": {
        "effective_date": "2019-01-01",
        "operative_date": "2019-01-01"
      },
      "source_ids": ["src:us-nc:osfm-current-past-codes", "src:us-nc:gs-143-138"],
      "verification": {
        "status": "partially_verified",
        "confidence": 0.88
      }
    },
    {
      "id": "adoption:us-nc:residential:2018-code",
      "code_family": "residential",
      "status": "current",
      "state_code_name": "2018 North Carolina State Building Code: Residential Code",
      "base_model_code": {
        "publisher": "ICC",
        "name": "International Residential Code",
        "edition_year": 2015
      },
      "adopting_authority_id": "auth:us-nc:state-code-councils",
      "enforcement_model": "local_with_state_supervision",
      "applies_to": ["detached_one_two_family", "townhouses"],
      "dates": {
        "effective_date": "2019-01-01",
        "operative_date": "2019-01-01"
      },
      "source_ids": ["src:us-nc:osfm-current-past-codes", "src:us-nc:gs-143-138"],
      "verification": {
        "status": "partially_verified",
        "confidence": 0.86
      }
    },
    {
      "id": "adoption:us-nc:building:2024-code-pending",
      "code_family": "building",
      "status": "pending",
      "state_code_name": "2024 North Carolina State Building Code",
      "adopting_authority_id": "auth:us-nc:state-code-councils",
      "enforcement_model": "local_with_state_supervision",
      "applies_to": ["construction_code_scope"],
      "dates": {},
      "source_ids": ["src:us-nc:osfm-current-past-codes", "src:us-nc:osfm-2024-delay"],
      "verification": {
        "status": "needs_review",
        "confidence": 0.8,
        "notes": "Implementation is conditional and no calendar operative date is asserted."
      }
    }
  ],
  "defaults": {
    "incorporated": {
      "status": "partially_resolved",
      "authority_candidates": [
        {
          "kind": "state_authority",
          "authority_id": "auth:us-nc:state-code-councils",
          "roles": ["adopt", "amend"],
          "source_ids": ["src:us-nc:gs-143-138"]
        },
        {
          "kind": "state_authority",
          "authority_id": "auth:us-nc:state-fire-marshal",
          "roles": ["supervises", "administers"],
          "source_ids": ["src:us-nc:gs-143-139"]
        },
        {
          "kind": "municipality",
          "label": "Local inspection department or contracted inspection service",
          "roles": ["permits", "inspects", "enforces"],
          "source_ids": ["src:us-nc:gs-160d-1102"]
        }
      ],
      "required_local_records": [
        "Identity and service area of the local inspection department or contracted service",
        "Project type and any state-reserved enforcement allocation",
        "Applicable local rules approved under state law"
      ],
      "warnings": [
        "The 2024 code is pending under a condition-based implementation rule and is not returned as current."
      ],
      "source_ids": ["src:us-nc:gs-143-138", "src:us-nc:gs-143-139", "src:us-nc:gs-160d-1102"]
    },
    "unincorporated": {
      "status": "partially_resolved",
      "authority_candidates": [
        {
          "kind": "state_authority",
          "authority_id": "auth:us-nc:state-code-councils",
          "roles": ["adopt", "amend"],
          "source_ids": ["src:us-nc:gs-143-138"]
        },
        {
          "kind": "state_authority",
          "authority_id": "auth:us-nc:state-fire-marshal",
          "roles": ["supervises", "administers"],
          "source_ids": ["src:us-nc:gs-143-139"]
        },
        {
          "kind": "county",
          "label": "County inspection department or contracted inspection service",
          "roles": ["permits", "inspects", "enforces"],
          "source_ids": ["src:us-nc:gs-160d-1102"]
        }
      ],
      "required_local_records": [
        "Identity and service area of the county inspection department or contracted service",
        "Project type and any state-reserved enforcement allocation",
        "Applicable local rules approved under state law"
      ],
      "warnings": [
        "The 2024 code is pending under a condition-based implementation rule and is not returned as current."
      ],
      "source_ids": ["src:us-nc:gs-143-138", "src:us-nc:gs-143-139", "src:us-nc:gs-160d-1102"]
    }
  },
  "code_family_overrides": {
    "building": {
      "status": "partially_resolved",
      "adoption_ids": ["adoption:us-nc:building:2018-code"],
      "required_local_records": ["Project type, permit record, and any approved local rule"],
      "warnings": ["The pending 2024 code is excluded until its condition-based implementation date is established."],
      "source_ids": ["src:us-nc:osfm-current-past-codes", "src:us-nc:osfm-2024-delay"]
    },
    "residential": {
      "status": "partially_resolved",
      "adoption_ids": ["adoption:us-nc:residential:2018-code"],
      "required_local_records": ["Project classification, permit record, and any approved local rule"],
      "warnings": ["The pending 2024 code is excluded until its condition-based implementation date is established."],
      "source_ids": ["src:us-nc:osfm-current-past-codes", "src:us-nc:osfm-2024-delay"]
    }
  },
  "verification": {
    "status": "needs_review",
    "confidence": 0.76,
    "notes": "The current 2018 baseline, local inspection framework, and conditional 2024 delay are source-backed; remaining code families and special allocations are incomplete."
  }
}
```

## Open questions

1. Has the State Fire Marshal issued the certification that starts the twelve-month implementation clock for the 2024 code?
2. How should the schema represent an effective date calculated from a future certification event?
3. What are the complete current editions and amendments for existing-building, mechanical, plumbing, fuel-gas, electrical, energy, fire, accessibility, and elevator families?
4. What local rules have been approved to supersede or supplement the state code in specific jurisdictions?
5. Which project classes are reserved to the Department of Administration or other state agencies?
6. What resolver fixtures best cover local service, joint service, contracted service, county service, tribal service, and state intervention?
