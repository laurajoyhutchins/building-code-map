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
const validateReceipt = ajv.compile(load('bulk-source-gap-summary-v1.schema.json'));

test('source candidate gap point can never claim verified legal jurisdiction', () => {
  const row = { id: 'gap:' + 'a'.repeat(24), source_sha256: 'b'.repeat(64),
    point: { latitude: 40.125, longitude: -105.075 },
    gap: { classification: 'source_geometry_candidate_only',
      kind: 'nested_feature_candidate', source_feature_id: 'town',
      containing_feature_ids: ['county', 'town'],
      baseline_sampled: false, crs: 'EPSG:6933' } };
  assert.equal(validateSample(row), true, JSON.stringify(validateSample.errors));
  assert.equal(validateSample({ ...row, gap: { ...row.gap, authority: 'local AHJ' } }), false);
  assert.equal(validateSample({ ...row, gap: { ...row.gap, classification: 'verified_code' } }), false);
  assert.equal(validateSample({ ...row, gap: { ...row.gap, baseline_sampled: true } }), false);
  assert.equal(validateSample({ ...row, gap: { ...row.gap, kind: 'verified_jurisdiction_gap' } }), false);
  assert.equal(validateSample({ ...row, source_sha256: undefined }), false);
});
test('source-gap summary reports omissions without implying national completeness', () => {
  const receipt = { schema_version: '1', classification: 'source_geometry_candidate_only',
    source_geometry_sha256: 'c'.repeat(64), baseline_sha256: 'd'.repeat(64),
    workload_sha256: 'e'.repeat(64), source_features: 3, baseline_samples: 1,
    baseline_represented_features: 1, missing_source_features: 2,
    generated_samples: 2, categories: { nested_feature_candidate: 1, interior_void_candidate: 1 } };
  assert.equal(validateReceipt(receipt), true, JSON.stringify(validateReceipt.errors));
  assert.equal(validateReceipt({ ...receipt, classification: 'verified_legal_gap' }), false);
  assert.equal(validateReceipt({ ...receipt, legal_authority: 'none' }), false);
  assert.equal(validateReceipt({ ...receipt, workload_sha256: 'bad' }), false);
  assert.deepEqual(validatePublicArtifact('schemas/bulk-source-gap-summary-v1.schema.json',
    JSON.stringify(load('bulk-source-gap-summary-v1.schema.json')), {}), []);
});
test('reviewed ODCS has exactly eight bulk exchanges and source-gap receipt', () => {
  const odcs = load('bulk-query.odcs.json');
  const validator = new Ajv2019({ allErrors: true, strict: false });
  addFormats(validator);
  assert.equal(validator.compile(load('odcs-v3.1.0.schema.json'))(odcs), true);
  assert.equal(odcs.schema.length, 8);
  assert.deepEqual(validatePublicArtifact('schemas/bulk-query.odcs.json', JSON.stringify(odcs), {}), []);
  const modified = structuredClone(odcs);
  modified.authoritativeDefinitions = modified.authoritativeDefinitions.filter(
    x => x.id !== 'bulk_source_gap_summary_json_schema');
  assert.match(validatePublicArtifact('schemas/bulk-query.odcs.json', JSON.stringify(modified), {}).join(';'),
    /must match reviewed schema-only identity/);
});
