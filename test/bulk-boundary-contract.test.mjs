import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import test from 'node:test';
import Ajv2020 from 'ajv/dist/2020.js';
import Ajv2019 from 'ajv/dist/2019.js';
import addFormats from 'ajv-formats';
import { validatePublicArtifact } from '../scripts/public-boundary.mjs';

const load = (name) => JSON.parse(readFileSync('schemas/' + name, 'utf8'));
const ajv = new Ajv2020({ allErrors: true, strict: false });
addFormats(ajv);
const sample = ajv.compile(load('bulk-workload-v1.schema.json'));
const receipt = ajv.compile(load('bulk-boundary-summary-v1.schema.json'));
const odcs = load('bulk-query.odcs.json');

test('boundary workload is source-only, not jurisdiction or legal evidence', () => {
  const row = {
    id: 'edge:' + 'a'.repeat(24) + ':a',
    point: { longitude: -105, latitude: 40 },
    source_sha256: 'b'.repeat(64),
    boundary: {
      pair_key: 'a'.repeat(24), feature_id: 'synthetic-A',
      neighbor_feature_id: 'synthetic-B', side: 'a',
      offset_metres: 100, crs: 'EPSG:6933',
      classification: 'source_polygon_candidate_only',
    },
  };
  assert.equal(sample(row), true, JSON.stringify(sample.errors));
  assert.equal(sample({ ...row, boundary: { ...row.boundary, classification: 'verified_legal_AHJ' } }), false);
  assert.equal(sample({ ...row, boundary: { ...row.boundary, authority: 'City Hall' } }), false);
  assert.equal(sample({ ...row, source_sha256: undefined }), false);
  assert.equal(sample({ ...row, boundary: { ...row.boundary, offset_metres: 0 } }), false);
  assert.equal(sample({ ...row, point: undefined, address: 'synthetic address' }), false);
});

test('boundary receipts preserve overlap and unresolved diagnostics without legal claims', () => {
  const row = {
    schema_version: '1', classification: 'source_polygon_candidate_only',
    source_geometry_sha256: 'b'.repeat(64), workload_sha256: 'c'.repeat(64),
    source_features: 3, candidate_touch_pairs: 2, sampled_pairs: 1,
    unresolved_touch_pairs: 1, overlapping_source_pairs: 0,
    max_pairs: 10000, requested_offset_metres: 100, samples: 2,
  };
  assert.equal(receipt(row), true, JSON.stringify(receipt.errors));
  assert.equal(receipt({ ...row, classification: 'verified' }), false);
  assert.equal(receipt({ ...row, private_source_text: 'not for publication' }), false);
  assert.equal(receipt({ ...row, requested_offset_metres: 0 }), false);
  assert.equal(receipt({ ...row, workload_sha256: 'unhashed' }), false);
  assert.deepEqual(validatePublicArtifact('schemas/bulk-boundary-summary-v1.schema.json',
    JSON.stringify(load('bulk-boundary-summary-v1.schema.json')), {}), []);
});

test('paired ODCS permits only the reviewed fifth bulk exchange schema', () => {
  const validator = new Ajv2019({ allErrors: true, strict: false });
  addFormats(validator);
  assert.equal(validator.compile(load('odcs-v3.1.0.schema.json'))(odcs), true);
  assert.equal(odcs.schema.length, 9);
  assert.deepEqual(validatePublicArtifact('schemas/bulk-query.odcs.json',
    JSON.stringify(odcs), {}), []);
  const tampered = structuredClone(odcs);
  tampered.authoritativeDefinitions = tampered.authoritativeDefinitions.filter(
    x => x.id !== 'bulk_boundary_summary_json_schema');
  assert.match(validatePublicArtifact('schemas/bulk-query.odcs.json',
    JSON.stringify(tampered), {}).join(';'), /must match reviewed schema-only identity/);
});
