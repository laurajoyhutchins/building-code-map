import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import test from 'node:test';
import Ajv2020 from 'ajv/dist/2020.js';
import Ajv2019 from 'ajv/dist/2019.js';
import addFormats from 'ajv-formats';
import { validatePublicArtifact } from '../scripts/public-boundary.mjs';

const load = name => JSON.parse(readFileSync('schemas/' + name, 'utf8'));
const ajv = new Ajv2020({ allErrors: true, strict: false });
addFormats(ajv);
const validateSample = ajv.compile(load('bulk-workload-v1.schema.json'));
const validateEvents = ajv.compile(load('bulk-temporal-events-v1.schema.json'));
const validateSummary = ajv.compile(load('bulk-temporal-summary-v1.schema.json'));

test('temporal candidate events reject legal claims, missing source, bad dates', () => {
  const base = { schema_version: '1', classification: 'source_date_candidate_only',
    events: [{ id: 'synthetic-transition', transition_date: '2025-12-31',
      source_sha256: 'b'.repeat(64) }] };
  assert.equal(validateEvents(base), true, JSON.stringify(validateEvents.errors));
  assert.equal(validateEvents({ ...base, classification: 'enacted_verified' }), false);
  assert.equal(validateEvents({ ...base, legal_authority: true }), false);
  assert.equal(validateEvents({ ...base, events: [{ ...base.events[0], transition_date: '2025-02-30' }] }), false);
  assert.equal(validateEvents({ ...base, events: [{ ...base.events[0], effective_as_law: true }] }), false);
});

test('temporal metadata is still a canonical workload, not an adopted code verdict', () => {
  const row = { id: 'temporal:' + 'a'.repeat(24),
    point: { latitude: 39, longitude: -105 }, applicability_date: '2025-12-30',
    temporal: { event_id: 'synthetic-transition', transition_date: '2025-12-31',
      offset_days: -1, source_sha256: 'b'.repeat(64),
      classification: 'source_date_candidate_only' } };
  assert.equal(validateSample(row), true, JSON.stringify(validateSample.errors));
  assert.equal(validateSample({ ...row, applicability_date: undefined }), false);
  assert.equal(validateSample({ ...row, temporal: { ...row.temporal, offset_days: 5 } }), false);
  assert.equal(validateSample({ ...row, temporal: { ...row.temporal, legally_effective: true } }), false);
  assert.equal(validateSample({ ...row, temporal: { ...row.temporal, classification: 'verified' } }), false);
});

test('temporal receipts and ODCS remain source-only versioned contracts', () => {
  const receipt = { schema_version: '1', classification: 'source_date_candidate_only',
    input_sha256: 'a'.repeat(64), transition_manifest_sha256: 'b'.repeat(64),
    workload_sha256: 'c'.repeat(64), source_samples: 2, candidate_events: 1, samples: 6 };
  assert.equal(validateSummary(receipt), true, JSON.stringify(validateSummary.errors));
  assert.equal(validateSummary({ ...receipt, classification: 'legal_verified' }), false);
  assert.equal(validateSummary({ ...receipt, enacted_code_set: ['2024 IBC'] }), false);
  const odcs = load('bulk-query.odcs.json');
  const ajv2019 = new Ajv2019({ allErrors: true, strict: false });
  addFormats(ajv2019);
  assert.equal(ajv2019.compile(load('odcs-v3.1.0.schema.json'))(odcs), true);
  assert.equal(odcs.schema.length, 8);
  assert.deepEqual(validatePublicArtifact('schemas/bulk-query.odcs.json', JSON.stringify(odcs), {}), []);
  for (const path of ['bulk-temporal-events-v1.schema.json', 'bulk-temporal-summary-v1.schema.json']) {
    assert.deepEqual(validatePublicArtifact('schemas/' + path, JSON.stringify(load(path)), {}), []);
  }
});
