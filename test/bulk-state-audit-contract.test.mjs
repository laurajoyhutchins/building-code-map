import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import test from 'node:test';
import Ajv2020 from 'ajv/dist/2020.js';
import { validatePublicArtifact } from '../scripts/public-boundary.mjs';

const schema = JSON.parse(readFileSync('schemas/bulk-state-audit.schema.json', 'utf8'));
const validate = new Ajv2020({ allErrors: true, strict: false }).compile(schema);

test('candidate-only geography coverage audit cannot be relabeled as legal verification', () => {
  const audit = {
    schema_version: '1', classification: 'source_geography_candidate_only',
    source_geometry_sha256: 'a'.repeat(64), workload_sha256: 'b'.repeat(64),
    samples: 3, states: 2, state_candidate_counts: { '01': 1, '02': 0 },
    unsampled_states: ['02'], unmatched_points: 1, overlapping_candidate_points: 1,
    unmatched_ids: ['out'], overlapping_ids: ['on-border'],
  };
  assert.equal(validate(audit), true, JSON.stringify(validate.errors));
  assert.equal(validate({ ...audit, classification: 'verified_legal_coverage' }), false);
  assert.equal(validate({ ...audit, legal_code_set: ['IBC'] }), false);
  assert.equal(validate({ ...audit, state_candidate_counts: { '01': -1 } }), false);
  assert.equal(validate({ ...audit, source_geometry_sha256: 'unverified' }), false);
  assert.deepEqual(validatePublicArtifact('schemas/bulk-state-audit.schema.json', JSON.stringify(schema), {}), []);
});
