import assert from 'node:assert/strict';
import test from 'node:test';
import { spawnSync } from 'node:child_process';
import { mkdtempSync, writeFileSync, rmSync } from 'node:fs';
import { join } from 'node:path';
import { tmpdir } from 'node:os';
import { buildNationalCoverageFrontier } from '../scripts/national-coverage-frontier.mjs';

// Only fictional identities and receipts. No source, permit or adoption claimed.
const hash = 'a'.repeat(64);
const expected = { as_of: '2026-11-07', project_class: 'DEMO-commercial', projection_snapshot_sha256: hash };
const record = (id) => ({
  entity_id: id, authority_level: 'municipality',
  inventory: { source_ref: 'DEMO-cgu', source_sha256: 'b'.repeat(64) },
  source_review: { status: 'not_checked' },
  authority: { status: 'unresolved' },
  geography: { status: 'unresolved' },
  provider: { status: 'unknown' },
  project_verification: { status: 'not_verified' },
  unknown_reasons: ['no_adoption_source'],
});
const ledger = (family = 'building', audited = true) => ({
  schema_version: '0.1',
  scope: { country_code: 'US', code_family: family, as_of: expected.as_of, project_class: expected.project_class },
  projection: {
    snapshot_id: 'DEMO-2026', snapshot_sha256: hash, source_ref: 'DEMO-publisher',
    universe: 'governmental_entities',
    denominator: audited ? {
      status: 'independently_audited',
      entity_ids: ['DEMO-TOWN', 'DEMO-COUNTY'],
      audit_receipt_ref: 'DEMO-independent-audit',
    } : { status: 'not_established' },
  },
  records: [record('DEMO-TOWN')],
});

test('missing families and unrepresented entities become explicit inventory gaps', () => {
  const result = buildNationalCoverageFrontier([ledger()], expected);
  assert.equal(result.classification, 'evidence_gap_queue_only');
  assert.equal(result.national_verified_project_coverage, 'not_established');
  assert.equal(result.counts.inventory_gaps, 6);
  assert.ok(result.work.some((x) => x.family === 'building' && x.entity_id === 'DEMO-COUNTY' && x.reason === 'entity_record_missing'));
  assert.ok(result.work.some((x) => x.family === 'electrical' && x.entity_id === null && x.reason === 'family_ledger_missing'));
  assert.ok(!result.work.some((x) => x.family === 'electrical' && x.entity_id === 'DEMO-TOWN'));
});

test('an unaudited denominator does not invent missing government counts', () => {
  const result = buildNationalCoverageFrontier([ledger('fire', false)], expected);
  assert.ok(result.work.some((x) => x.family === 'fire' && x.reason === 'entity_denominator_not_audited'));
  assert.ok(!result.work.some((x) => x.family === 'fire' && x.reason === 'entity_record_missing'));
});

test('separate adoption and project limitations preserve candidate identity', () => {
  const source = ledger();
  source.records[0].adoption_candidate = {
    source_ref: 'DEMO-ordinance', source_sha256: 'c'.repeat(64),
    code_family: 'building', edition: 'DEMO-2024', amendments: [],
    enacted_on: '2025-05-01', effective_on: '2026-01-01',
  };
  source.records[0].unknown_reasons = ['source_unreviewed'];
  const result = buildNationalCoverageFrontier([source], expected);
  assert.ok(result.work.some((x) => x.reason === 'adoption_source_review_missing' && x.source_ref === 'DEMO-ordinance' && x.source_sha256 === 'c'.repeat(64)));
  assert.ok(result.work.some((x) => x.reason === 'legal_authority_unresolved'));
  assert.ok(result.work.some((x) => x.reason === 'project_verification_incomplete' && x.recorded_unknown_reasons.includes('address_jurisdiction_unverified')));
  assert.ok(!result.work.some((x) => x.reason === 'adoption_source_missing'));
  assert.equal(result.counts.project_limitations, 1);
});

test('no family ledgers explicitly identifies six missing families', () => {
  const result = buildNationalCoverageFrontier([], expected);
  assert.equal(result.counts.inventory_gaps, 6);
  assert.equal(result.counts.adoption_gaps, 0);
});

test('mixed source vintages and forged verification receipt fail closed', () => {
  const other = ledger('energy');
  other.projection.snapshot_sha256 = 'f'.repeat(64);
  assert.throws(() => buildNationalCoverageFrontier([ledger(), other], expected), /projection snapshot mismatch/);
  const tampered = ledger();
  tampered.records[0].project_verification.status = 'independently_verified';
  assert.throws(() => buildNationalCoverageFrontier([tampered], expected), /schema invalid/);
});

test('work ordering does not depend on record or file order', () => {
  const building = ledger();
  building.records.push(record('DEMO-COUNTY'));
  const fire = ledger('fire');
  const first = buildNationalCoverageFrontier([fire, building], expected);
  building.records.reverse();
  assert.deepEqual(first, buildNationalCoverageFrontier([building, fire], expected));
});

test('CLI produces no partial output for malformed source input', () => {
  const dir = mkdtempSync(join(tmpdir(), 'DEMO-coverage-frontier-'));
  try {
    const scope = join(dir, 'scope.json');
    const source = join(dir, 'source.json');
    writeFileSync(scope, JSON.stringify(expected));
    writeFileSync(source, JSON.stringify(ledger()));
    const ok = spawnSync(process.execPath, ['scripts/national-coverage-frontier.mjs', scope, source], { encoding: 'utf8' });
    assert.equal(ok.status, 0, ok.stderr);
    assert.equal(JSON.parse(ok.stdout).classification, 'evidence_gap_queue_only');
    writeFileSync(source, '{bad');
    const bad = spawnSync(process.execPath, ['scripts/national-coverage-frontier.mjs', scope, source], { encoding: 'utf8' });
    assert.notEqual(bad.status, 0);
    assert.equal(bad.stdout, '');
  } finally { rmSync(dir, { recursive: true, force: true }); }
});
