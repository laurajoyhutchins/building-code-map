// Public-safe, source-free evidence-ledger validator. This is not an authority
// engine and cannot establish the truth of referenced external receipts.
import { readFileSync } from 'node:fs';
import { fileURLToPath } from 'node:url';
import { resolve } from 'node:path';
import Ajv2020 from 'ajv/dist/2020.js';
import addFormats from 'ajv-formats';

const schema = JSON.parse(readFileSync(new URL('../schemas/code-family-coverage-ledger.schema.json', import.meta.url), 'utf8'));
const ajv = new Ajv2020({ allErrors: true, strict: false });
addFormats(ajv);
const isValidSchema = ajv.compile(schema);

export function validateCoverageLedger(ledger) {
  if (!isValidSchema(ledger)) {
    const details = isValidSchema.errors.map((error) => `${error.instancePath || '/'} ${error.message}`).join('; ');
    throw new Error(`coverage ledger schema invalid: ${details}`);
  }

  const errors = [];
  const seen = new Set();
  const { as_of: asOf, code_family: family, project_class: projectClass } = ledger.scope;
  const denominator = ledger.projection.denominator;
  const allowed = denominator.status === 'independently_audited'
    ? new Set(denominator.entity_ids) : null;

  ledger.records.forEach((record, index) => {
    const prefix = `records[${index}] (${record.entity_id})`;
    if (seen.has(record.entity_id)) errors.push(`${prefix}: duplicate entity_id within family/date/project class`);
    seen.add(record.entity_id);
    if (allowed && !allowed.has(record.entity_id)) errors.push(`${prefix}: not in audited denominator`);

    const candidate = record.adoption_candidate;
    if (candidate) {
      if (candidate.code_family !== family) {
        errors.push(`${prefix}: adoption candidate code family does not match ledger scope`);
      }
      if (candidate.transition_end_on && candidate.transition_end_on < candidate.effective_on) {
        errors.push(`${prefix}: transition ends before effective date`);
      }
      if (candidate.enacted_on && candidate.enacted_on > candidate.effective_on) {
        errors.push(`${prefix}: enactment occurs after effective date`);
      }
      if (candidate.effective_on > asOf && record.authority.status === 'resolved') {
        errors.push(`${prefix}: future effective date cannot be resolved for scope.as_of`);
      }
    }

    if (record.project_verification.status === 'independently_verified') {
      const receipt = record.project_verification.receipt;
      if (receipt.entity_id !== record.entity_id) {
        errors.push(`${prefix}: project receipt entity does not match ledger entity`);
      }
      if (receipt.code_family !== family || receipt.as_of !== asOf || receipt.project_class !== projectClass) {
        errors.push(`${prefix}: project receipt family/date/class does not match ledger scope`);
      }
      if (candidate && candidate.effective_on > asOf) {
        errors.push(`${prefix}: project receipt uses future-effective adoption`);
      }
    }
  });

  if (errors.length) throw new Error(`coverage ledger integrity invalid: ${errors.join('; ')}`);
  return ledger;
}

export function summarizeCoverageLedger(ledger) {
  validateCoverageLedger(ledger);
  const records = ledger.records;
  const denominator = ledger.projection.denominator;
  const audited = denominator.status === 'independently_audited';
  // Even an audited governmental-entity denominator is not an authoritative
  // address-level jurisdiction/authority-closure denominator.
  return {
    code_family: ledger.scope.code_family,
    as_of: ledger.scope.as_of,
    project_class: ledger.scope.project_class,
    projection_snapshot_sha256: ledger.projection.snapshot_sha256,
    denominator_status: denominator.status,
    denominator_entities: audited ? denominator.entity_ids.length : null,
    inspected_entity_records: records.length,
    unrepresented_denominator_entities: audited ? denominator.entity_ids.length - records.length : null,
    adoption_candidate_records: records.filter((r) => !!r.adoption_candidate).length,
    independently_checked_source_records: records.filter((r) => r.source_review.status === 'independently_checked').length,
    legally_resolved_records: records.filter((r) => r.authority.status === 'resolved').length,
    provider_conformance_records: records.filter((r) => r.provider.status === 'conformance_verified').length,
    // Syntactically complete, internally consistent claims only. External
    // execution, reviewer independence and source identity still require audit.
    claimed_independently_verified_project_receipts: records.filter((r) => r.project_verification.status === 'independently_verified').length,
    unverified_records: records.filter((r) => r.project_verification.status !== 'independently_verified').length,
    conflicting_records: records.filter((r) => r.authority.status === 'conflicting' || r.source_review.status === 'conflicting').length,
    national_verified_project_coverage: 'not_established',
    address_level_jurisdiction_coverage: 'not_established',
  };
}

const REQUIRED_CODE_FAMILIES = Object.freeze([
  'building', 'electrical', 'fire', 'energy', 'plumbing', 'mechanical',
]);

/**
 * Admission for multi-family projection. The caller must supply the expected
 * entity-snapshot digest and applicability scope, not infer them from the
 * first submitted family. A missing ledger is UNKNOWN, not a zero numerator.
 *
 * This remains syntactic admission only. External source and reviewer receipts
 * must be independently verified before real-world coverage is asserted.
 */
export function summarizeCodeFamilyMatrix(ledgers, expected) {
  if (!Array.isArray(ledgers) ||
      typeof expected?.as_of !== 'string' ||
      !/^\d{4}-\d{2}-\d{2}$/.test(expected.as_of) ||
      typeof expected?.project_class !== 'string' || !expected.project_class ||
      typeof expected?.projection_snapshot_sha256 !== 'string' ||
      !/^[0-9a-f]{64}$/.test(expected.projection_snapshot_sha256)) {
    throw new Error('family matrix requires ledgers and explicit date/class/projection digest');
  }

  const families = Object.fromEntries(
    REQUIRED_CODE_FAMILIES.map((family) => [
      family, {
        status: 'ledger_missing',
        denominator_entities: null,
        claimed_independently_verified_project_receipts: null,
        national_verified_project_coverage: 'not_established',
      },
    ]),
  );
  const seen = new Set();
  let firstSnapshotId = null;
  let firstSourceRef = null;

  for (const ledger of ledgers) {
    // Reject an invalid supplied family before using it as a property key.
    const summary = summarizeCoverageLedger(ledger);
    const family = summary.code_family;
    if (seen.has(family)) throw new Error(`duplicate code family: ${family}`);
    seen.add(family);
    if (ledger.scope.as_of !== expected.as_of ||
        ledger.scope.project_class !== expected.project_class) {
      throw new Error(`applicability scope mismatch for code family ${family}`);
    }
    if (ledger.projection.snapshot_sha256 !== expected.projection_snapshot_sha256) {
      throw new Error(`projection snapshot mismatch for code family ${family}`);
    }
    if (firstSnapshotId !== null &&
        (ledger.projection.snapshot_id !== firstSnapshotId ||
         ledger.projection.source_ref !== firstSourceRef)) {
      throw new Error(`projection identity mismatch for code family ${family}`);
    }
    firstSnapshotId = ledger.projection.snapshot_id;
    firstSourceRef = ledger.projection.source_ref;
    // The schema admits only canonical built-ins or x_ extension family codes.
    families[family] = { status: 'ledger_present', ...summary };
  }

  return {
    scope: { country_code: 'US', as_of: expected.as_of, project_class: expected.project_class },
    projection_snapshot_sha256: expected.projection_snapshot_sha256,
    families,
    missing_required_families: REQUIRED_CODE_FAMILIES.filter((family) => !seen.has(family)),
    national_verified_project_coverage: 'not_established',
    address_level_jurisdiction_coverage: 'not_established',
  };
}

// Producers should call this at input admission, before any report/publication.
// No partial report is emitted when a file cannot be parsed or validated.
if (process.argv[1] && fileURLToPath(import.meta.url) === resolve(process.argv[1])) {
  try {
    if (process.argv.length !== 4 || process.argv[2] !== 'validate') {
      throw new Error('usage: node scripts/code-family-coverage-ledger.mjs validate <ledger.json>');
    }
    let ledger;
    try {
      ledger = JSON.parse(readFileSync(process.argv[3], 'utf8'));
    } catch (error) {
      throw new Error(`ledger input invalid: ${error.message}`);
    }
    const result = summarizeCoverageLedger(ledger);
    process.stdout.write(`${JSON.stringify(result, null, 2)}\n`);
  } catch (error) {
    console.error(error.message);
    process.exitCode = 1;
  }
}
