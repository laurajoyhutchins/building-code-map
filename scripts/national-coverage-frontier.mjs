// Public-safe national research work queue derived from BCM's existing ledger.
// A gap is never an adoption verdict, evidence receipt, or legal authority.
import { readFileSync } from 'node:fs';
import { fileURLToPath } from 'node:url';
import { resolve } from 'node:path';
import { summarizeCodeFamilyMatrix } from './code-family-coverage-ledger.mjs';

const REQUIRED_FAMILIES = ['building', 'electrical', 'fire', 'energy', 'plumbing', 'mechanical'];
const FAMILY_ORDER = new Map(REQUIRED_FAMILIES.map((family, index) => [family, index]));
const sortedUnique = (values) => [...new Set(values)].sort();

/**
 * Project actionable national research gaps from schema/semantics-validated
 * family ledgers. Project-site and project-fact deficits are a separate lane:
 * they are not part of a national legal adoption coverage denominator.
 * No original sources are fetched, reviewed or self-approved here.
 */
export function buildNationalCoverageFrontier(ledgers, expected) {
  const matrix = summarizeCodeFamilyMatrix(ledgers, expected); // validate ALL inputs first
  const work = [];
  const add = (family, entityId, reason, details = {}) => {
    work.push({ family, entity_id: entityId, reason, ...details });
  };

  for (const family of matrix.missing_required_families) {
    add(family, null, 'family_ledger_missing', { lane: 'inventory' });
  }

  for (const ledger of [...ledgers].sort((a, b) => a.scope.code_family.localeCompare(b.scope.code_family))) {
    const family = ledger.scope.code_family;
    const denominator = ledger.projection.denominator;
    const observed = new Map(ledger.records.map((record) => [record.entity_id, record]));
    if (denominator.status !== 'independently_audited') {
      add(family, null, 'entity_denominator_not_audited', { lane: 'inventory' });
    } else {
      for (const entityId of [...denominator.entity_ids].sort()) {
        if (!observed.has(entityId)) {
          add(family, entityId, 'entity_record_missing', { lane: 'inventory' });
        }
      }
    }

    for (const record of [...ledger.records].sort((a, b) => a.entity_id.localeCompare(b.entity_id))) {
      const entityId = record.entity_id;
      const common = {
        lane: 'adoption',
        source_ref: record.adoption_candidate?.source_ref ?? null,
        source_sha256: record.adoption_candidate?.source_sha256 ?? null,
        recorded_unknown_reasons: sortedUnique(record.unknown_reasons),
      };
      if (!record.adoption_candidate) {
        add(family, entityId, 'adoption_source_missing', common);
      } else if (record.source_review.status === 'not_checked') {
        add(family, entityId, 'adoption_source_review_missing', common);
      } else if (record.source_review.status === 'conflicting') {
        add(family, entityId, 'adoption_source_conflicting', common);
      }
      if (record.authority.status === 'unresolved') {
        add(family, entityId, 'legal_authority_unresolved', common);
      } else if (record.authority.status === 'conflicting') {
        add(family, entityId, 'legal_authority_conflicting', common);
      }

      // A project receipt is site-specific, not a jurisdiction-wide adoption
      // proof. Preserve limitations but never expand them into adoption counts.
      if (record.project_verification.status !== 'independently_verified') {
        const reasons = sortedUnique([
          ...record.unknown_reasons,
          ...(record.geography.status !== 'project_location_verified' ? ['address_jurisdiction_unverified'] : []),
          ...(record.provider.status !== 'conformance_verified' ? ['provider_not_conformant'] : []),
          'project_receipt_missing',
        ]);
        add(family, entityId, 'project_verification_incomplete', {
          lane: 'project', recorded_unknown_reasons: reasons,
        });
      }
    }
  }

  work.sort((a, b) => {
    const familyIndex = (FAMILY_ORDER.get(a.family) ?? 100) - (FAMILY_ORDER.get(b.family) ?? 100);
    return familyIndex || a.family.localeCompare(b.family) ||
      (a.entity_id ?? '').localeCompare(b.entity_id ?? '') ||
      a.lane.localeCompare(b.lane) || a.reason.localeCompare(b.reason);
  });
  return {
    schema_version: '1',
    classification: 'evidence_gap_queue_only',
    scope: matrix.scope,
    projection_snapshot_sha256: matrix.projection_snapshot_sha256,
    national_verified_project_coverage: 'not_established',
    missing_required_families: matrix.missing_required_families,
    counts: {
      inventory_gaps: work.filter((item) => item.lane === 'inventory').length,
      adoption_gaps: work.filter((item) => item.lane === 'adoption').length,
      project_limitations: work.filter((item) => item.lane === 'project').length,
    },
    work,
  };
}

// Input admission is atomic: never print a partial frontier on invalid files.
if (process.argv[1] && fileURLToPath(import.meta.url) === resolve(process.argv[1])) {
  try {
    const [, , expectedPath, ...ledgerPaths] = process.argv;
    if (!expectedPath) throw new Error('usage: node scripts/national-coverage-frontier.mjs <expected.json> [family-ledger.json ...]');
    const read = (path) => {
      try { return JSON.parse(readFileSync(path, 'utf8')); }
      catch (error) { throw new Error('frontier input invalid (' + path + '): ' + error.message); }
    };
    const result = buildNationalCoverageFrontier(ledgerPaths.map(read), read(expectedPath));
    process.stdout.write(JSON.stringify(result, null, 2) + '\n');
  } catch (error) {
    console.error(error.message);
    process.exitCode = 1;
  }
}
