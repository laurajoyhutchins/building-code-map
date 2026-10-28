// Public-safe, source-free evidence-ledger validator. This is not an authority
// engine and cannot establish the truth of referenced external receipts.
import { readFileSync } from 'node:fs';
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
