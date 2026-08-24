import test from 'node:test';
import assert from 'node:assert/strict';
import {
  validatePolicy,
  validateSyntheticFixture,
  validateRepositoryPaths,
} from '../scripts/public-boundary.mjs';

const policy = {
  version: 1,
  default_disposition: 'review_required',
  synthetic_jurisdiction_prefix: 'DEMO-',
  forbidden_top_level_paths: ['corpus', 'production-data', 'research'],
  material_classes: {
    external_material: { public_repository: 'review_required' },
    production_regulatory_corpus: { public_repository: 'prohibited' },
    restricted_standard_model_code: { public_repository: 'prohibited' },
  },
};

test('external material cannot be admitted to the public repository by default', () => {
  assert.deepEqual(validatePolicy(policy), []);
  const unsafe = structuredClone(policy);
  unsafe.material_classes.external_material.public_repository = 'allowed';
  assert.match(validatePolicy(unsafe).join('\n'), /external_material.*review_required/);
});

test('production corpus and restricted standards cannot be distributed from the public repository', () => {
  const unsafe = structuredClone(policy);
  unsafe.material_classes.production_regulatory_corpus.public_repository = 'allowed';
  unsafe.material_classes.restricted_standard_model_code.public_repository = 'allowed';
  const errors = validatePolicy(unsafe).join('\n');
  assert.match(errors, /production_regulatory_corpus.*prohibited/);
  assert.match(errors, /restricted_standard_model_code.*prohibited/);
});

test('public fixtures must remain synthetic', () => {
  assert.deepEqual(validateSyntheticFixture({ jurisdiction_id: 'DEMO-XX' }, policy), []);
  assert.match(
    validateSyntheticFixture({ jurisdiction_id: 'CO-DENVER' }, policy).join('\n'),
    /synthetic jurisdiction prefix/,
  );
});

test('production-data paths cannot appear at repository top level', () => {
  assert.deepEqual(validateRepositoryPaths(['README.md', 'demo', 'docs'], policy), []);
  assert.match(
    validateRepositoryPaths(['README.md', 'corpus', 'docs'], policy).join('\n'),
    /forbidden top-level path: corpus/,
  );
});
