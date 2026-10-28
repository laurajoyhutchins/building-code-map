import assert from 'node:assert/strict';
import test from 'node:test';
import { summarizeCoverageLedger, summarizeCodeFamilyMatrix } from '../scripts/code-family-coverage-ledger.mjs';
import { spawnSync } from 'node:child_process';
import { mkdtempSync, writeFileSync, rmSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';

// All identities and source references below are fictional, not production
// regulatory observations or actual execution/independent-review receipts.
const sha = (character) => character.repeat(64);
const clone = (value) => structuredClone(value);

function syntheticLedger(family = 'building') {
  return {
    schema_version: '0.1',
    scope: { country_code: 'US', code_family: family, as_of: '2026-10-28', project_class: 'DEMO-new-commercial' },
    projection: {
      snapshot_id: 'DEMO-projection-v1', snapshot_sha256: sha('a'),
      source_ref: 'DEMO-source-snapshot', universe: 'governmental_entities',
      denominator: { status: 'not_established' },
    },
    records: [{
      entity_id: 'DEMO-XX-TOWN', authority_level: 'municipality',
      inventory: { source_ref: 'DEMO-inventory', source_sha256: sha('b') },
      adoption_candidate: {
        source_ref: 'DEMO-ordinance', source_sha256: sha('c'), edition: 'DEMO-2024',
        amendments: ['DEMO-amendment-A'], enacted_on: '2024-05-01', effective_on: '2025-01-01',
      },
      source_review: { status: 'not_checked' }, authority: { status: 'unresolved' },
      geography: { status: 'unresolved' }, provider: { status: 'unknown' },
      project_verification: { status: 'not_verified' }, unknown_reasons: ['source_unreviewed'],
    }],
  };
}

function verifiedLedger() {
  const ledger = syntheticLedger();
  const record = ledger.records[0];
  record.source_review = { status: 'independently_checked', receipt_ref: 'DEMO-source-review' };
  record.authority = { status: 'resolved', basis_ref: 'DEMO-legal-basis' };
  record.geography = { status: 'project_location_verified', receipt_ref: 'DEMO-geometry-review' };
  record.provider = { status: 'conformance_verified', receipt_ref: 'DEMO-provider-conformance' };
  record.project_verification = {
    status: 'independently_verified',
    receipt: {
      project_id: 'DEMO-PROJECT-1', code_family: ledger.scope.code_family,
      as_of: ledger.scope.as_of, project_class: ledger.scope.project_class,
      engine_commit: 'd'.repeat(40), source_revision_sha256: sha('e'),
      provider_release_sha256: sha('f'), result_sha256: sha('1'),
      execution_ref: 'DEMO-run', independent_review_ref: 'DEMO-independent-review',
    },
  };
  record.unknown_reasons = [];
  return ledger;
}

const rejects = (ledger, pattern) => assert.throws(() => summarizeCoverageLedger(ledger), pattern);

test('family projections stay independent: building receipt never counts as electrical coverage', () => {
  const building = summarizeCoverageLedger(verifiedLedger());
  const electrical = summarizeCoverageLedger(syntheticLedger('electrical'));
  assert.equal(building.claimed_independently_verified_project_receipts, 1);
  assert.equal(electrical.claimed_independently_verified_project_receipts, 0);
  assert.equal(electrical.adoption_candidate_records, 1);
  assert.equal(electrical.legally_resolved_records, 0);
  for (const summary of [building, electrical]) {
    assert.equal(summary.denominator_entities, null);
    assert.equal(summary.national_verified_project_coverage, 'not_established');
    assert.equal(summary.address_level_jurisdiction_coverage, 'not_established');
  }
});

test('an inventoried family with no adoption source remains unknown', () => {
  const ledger = syntheticLedger('fire');
  delete ledger.records[0].adoption_candidate;
  ledger.records[0].unknown_reasons = ['no_adoption_source'];
  const summary = summarizeCoverageLedger(ledger);
  assert.equal(summary.inspected_entity_records, 1);
  assert.equal(summary.adoption_candidate_records, 0);
  assert.equal(summary.unverified_records, 1);
});

test('the projection denominator cannot be invented without an independent audit', () => {
  const ledger = syntheticLedger();
  ledger.projection.denominator.entity_ids = ['DEMO-XX-TOWN'];
  rejects(ledger, /coverage ledger schema invalid/);
});

test('an audited denominator tracks missing records but does not claim national coverage', () => {
  const ledger = syntheticLedger('mechanical');
  ledger.projection.denominator = {
    status: 'independently_audited', entity_ids: ['DEMO-XX-TOWN', 'DEMO-XX-COUNTY'],
    audit_receipt_ref: 'DEMO-denominator-review',
  };
  const summary = summarizeCoverageLedger(ledger);
  assert.equal(summary.denominator_entities, 2);
  assert.equal(summary.unrepresented_denominator_entities, 1);
  assert.equal(summary.national_verified_project_coverage, 'not_established');
  ledger.records[0].entity_id = 'DEMO-XX-OUTSIDE';
  rejects(ledger, /not in audited denominator/);
});

test('duplicate entity IDs fail closed rather than inflating a family numerator', () => {
  const ledger = syntheticLedger();
  ledger.records.push(clone(ledger.records[0]));
  rejects(ledger, /duplicate entity_id/);
});

test('candidate adoption alone cannot become a verified project', () => {
  const ledger = syntheticLedger();
  ledger.records[0].project_verification = verifiedLedger().records[0].project_verification;
  ledger.records[0].unknown_reasons = [];
  rejects(ledger, /coverage ledger schema invalid/);
});

test('a verified status without an execution receipt is invalid', () => {
  const ledger = verifiedLedger();
  delete ledger.records[0].project_verification.receipt;
  rejects(ledger, /coverage ledger schema invalid/);
});

test('a receipt for the wrong family, date or project class cannot be promoted', () => {
  for (const [field, value] of [['code_family', 'electrical'], ['as_of', '2024-01-01'], ['project_class', 'DEMO-residential']]) {
    const ledger = verifiedLedger();
    ledger.records[0].project_verification.receipt[field] = value;
    rejects(ledger, /receipt family\/date\/class does not match/);
  }
});

test('future-effective adoption cannot count as resolved on the scoped date', () => {
  const ledger = verifiedLedger();
  ledger.records[0].adoption_candidate.effective_on = '2027-01-01';
  rejects(ledger, /future effective date/);
});

test('inverted transition and enactment dates fail closed', () => {
  const ledger = syntheticLedger();
  ledger.records[0].adoption_candidate.transition_end_on = '2024-01-01';
  rejects(ledger, /transition ends before effective date/);
  const later = syntheticLedger();
  later.records[0].adoption_candidate.enacted_on = '2026-01-01';
  rejects(later, /enactment occurs after effective date/);
});

test('conflicting amendments or overlapping authorities cannot masquerade as independently verified', () => {
  const ledger = verifiedLedger();
  ledger.records[0].authority = { status: 'conflicting', basis_ref: 'DEMO-overlap-dispute' };
  rejects(ledger, /coverage ledger schema invalid/);
  const sourceConflict = verifiedLedger();
  sourceConflict.records[0].source_review = { status: 'conflicting', receipt_ref: 'DEMO-amendment-conflict' };
  rejects(sourceConflict, /coverage ledger schema invalid/);
});

test('an unverified project location is not a verified jurisdictional determination', () => {
  const ledger = verifiedLedger();
  ledger.records[0].geography = { status: 'unresolved' };
  rejects(ledger, /coverage ledger schema invalid/);
});

test('unknown or misspelled families do not silently become building coverage', () => {
  const typo = syntheticLedger('electrcial');
  rejects(typo, /coverage ledger schema invalid/);
  assert.equal(summarizeCoverageLedger(syntheticLedger('x_solar')).code_family, 'x_solar');
});

test('an independently reviewed source is not yet legally resolved or provider-capable', () => {
  const ledger = syntheticLedger('plumbing');
  ledger.records[0].source_review = { status: 'independently_checked', receipt_ref: 'DEMO-source-review' };
  const summary = summarizeCoverageLedger(ledger);
  assert.equal(summary.independently_checked_source_records, 1);
  assert.equal(summary.legally_resolved_records, 0);
  assert.equal(summary.provider_conformance_records, 0);
  assert.equal(summary.claimed_independently_verified_project_receipts, 0);
});

test('missing family ledgers remain explicitly unknown, never zero coverage', () => {
  const matrix = summarizeCodeFamilyMatrix([verifiedLedger()], {
    as_of: '2026-10-28', project_class: 'DEMO-new-commercial',
    projection_snapshot_sha256: sha('a'),
  });
  assert.equal(matrix.families.building.status, 'ledger_present');
  assert.equal(matrix.families.building.claimed_independently_verified_project_receipts, 1);
  for (const family of ['electrical', 'fire', 'energy', 'plumbing', 'mechanical']) {
    assert.equal(matrix.families[family].status, 'ledger_missing');
    assert.equal(matrix.families[family].denominator_entities, null);
    assert.equal(matrix.families[family].claimed_independently_verified_project_receipts, null);
  }
  assert.equal(matrix.national_verified_project_coverage, 'not_established');
  assert.equal(matrix.address_level_jurisdiction_coverage, 'not_established');
});

test('independent source-admission matrix keeps same-snapshot family evidence separate', () => {
  const matrix = summarizeCodeFamilyMatrix([verifiedLedger(), syntheticLedger('electrical')], {
    as_of: '2026-10-28', project_class: 'DEMO-new-commercial',
    projection_snapshot_sha256: sha('a'),
  });
  assert.equal(matrix.families.building.claimed_independently_verified_project_receipts, 1);
  assert.equal(matrix.families.electrical.claimed_independently_verified_project_receipts, 0);
  assert.equal(matrix.families.electrical.adoption_candidate_records, 1);
  assert.equal(matrix.families.fire.status, 'ledger_missing');
});

test('no duplicate family or mixed projection vintage enters aggregate coverage', () => {
  const expected = {
    as_of: '2026-10-28', project_class: 'DEMO-new-commercial',
    projection_snapshot_sha256: sha('a'),
  };
  assert.throws(() => summarizeCodeFamilyMatrix([syntheticLedger(), syntheticLedger()], expected), /duplicate code family/);
  const wrongSnapshot = syntheticLedger('electrical');
  wrongSnapshot.projection.snapshot_sha256 = sha('9');
  assert.throws(() => summarizeCodeFamilyMatrix([syntheticLedger(), wrongSnapshot], expected), /projection snapshot mismatch/);
});

test('mixed applicability dates and project classes are rejected before national coverage projection', () => {
  const expected = {
    as_of: '2026-10-28', project_class: 'DEMO-new-commercial',
    projection_snapshot_sha256: sha('a'),
  };
  const wrongDate = syntheticLedger('energy');
  wrongDate.scope.as_of = '2026-10-27';
  assert.throws(() => summarizeCodeFamilyMatrix([wrongDate], expected), /applicability scope mismatch/);
  const wrongClass = syntheticLedger('fire');
  wrongClass.scope.project_class = 'DEMO-residential';
  assert.throws(() => summarizeCodeFamilyMatrix([wrongClass], expected), /applicability scope mismatch/);
});

test('at-admission CLI accepts only structurally valid ledger before output', () => {
  const dir = mkdtempSync(join(tmpdir(), 'DEMO-family-ledger-'));
  const input = join(dir, 'DEMO-ledger.json');
  try {
    writeFileSync(input, JSON.stringify(syntheticLedger('plumbing')));
    const valid = spawnSync(process.execPath, ['scripts/code-family-coverage-ledger.mjs', 'validate', input], {
      encoding: 'utf8',
    });
    assert.equal(valid.status, 0, valid.stderr);
    assert.match(valid.stdout, /"code_family":\s*"plumbing"/);
    assert.match(valid.stdout, /"national_verified_project_coverage":\s*"not_established"/);

    const invalid = syntheticLedger('plumbing');
    invalid.records[0].authority = { status: 'resolved', basis_ref: 'DEMO-legal-basis' };
    writeFileSync(input, JSON.stringify(invalid));
    const rejected = spawnSync(process.execPath, ['scripts/code-family-coverage-ledger.mjs', 'validate', input], {
      encoding: 'utf8',
    });
    assert.notEqual(rejected.status, 0);
    assert.equal(rejected.stdout, '');
    assert.match(rejected.stderr, /coverage ledger schema invalid/);
  } finally {
    rmSync(dir, { recursive: true, force: true });
  }
});

test('at-admission CLI rejects malformed JSON instead of emitting partial coverage', () => {
  const dir = mkdtempSync(join(tmpdir(), 'DEMO-family-ledger-'));
  const input = join(dir, 'DEMO-ledger.json');
  try {
    writeFileSync(input, '{ "schema_version": ');
    const result = spawnSync(process.execPath, ['scripts/code-family-coverage-ledger.mjs', 'validate', input], {
      encoding: 'utf8',
    });
    assert.notEqual(result.status, 0);
    assert.equal(result.stdout, '');
    assert.match(result.stderr, /ledger input invalid/);
  } finally {
    rmSync(dir, { recursive: true, force: true });
  }
});
