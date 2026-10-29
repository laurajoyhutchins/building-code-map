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
const check = ajv.compile(load('accuracy-review-queue-v1.schema.json'));

test('blind review selection cannot be falsely marked independently adjudicated', () => {
  const sample = {
    schema_version: '1', status: 'pending_independent_review',
    input_sha256: 'a'.repeat(64), source_samples: 1, max_cases: 100,
    case_count: 1, selected_cohorts: { boundary: 1 },
    cases: [{ case_id: 'review:' + 'b'.repeat(24), workload_id: 'edge:1',
      candidate_tags: ['boundary'], status: 'pending_independent_review' }],
    legal_accuracy_reviewed: false, expected_results_present: false,
  };
  assert.equal(check(sample), true, JSON.stringify(check.errors));
  assert.equal(check({ ...sample, legal_accuracy_reviewed: true }), false);
  assert.equal(check({ ...sample, expected_results_present: true }), false);
  assert.equal(check({ ...sample, expected_code: '2024 IBC' }), false);
  assert.equal(check({ ...sample, cases: [{ ...sample.cases[0], expected_verdict: 'verified' }] }), false);
  assert.equal(check({ ...sample, cases: [{ ...sample.cases[0], reviewer: 'self' }] }), false);
  assert.equal(check({ ...sample, status: 'verified' }), false);
});

test('blind queue contract is schema-only and ODCS pinned', () => {
  const odcs = load('bulk-query.odcs.json');
  const validator = new Ajv2019({ allErrors: true, strict: false });
  addFormats(validator);
  assert.equal(validator.compile(load('odcs-v3.1.0.schema.json'))(odcs), true);
  assert.equal(odcs.schema.length, 9);
  assert.deepEqual(validatePublicArtifact('schemas/accuracy-review-queue-v1.schema.json',
    JSON.stringify(load('accuracy-review-queue-v1.schema.json')), {}), []);
  assert.deepEqual(validatePublicArtifact('schemas/bulk-query.odcs.json', JSON.stringify(odcs), {}), []);
});
