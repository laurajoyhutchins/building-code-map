import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import test from 'node:test';
import Ajv2020 from 'ajv/dist/2020.js';
import Ajv2019 from 'ajv/dist/2019.js';
import addFormats from 'ajv-formats';
import { validatePublicArtifact } from '../scripts/public-boundary.mjs';

const load = (name) => JSON.parse(readFileSync('schemas/' + name, 'utf8'));
const sampleSchema = load('bulk-workload-v1.schema.json');
const resultSchema = load('bulk-result-v1.schema.json');
const summarySchema = load('bulk-summary-v1.schema.json');
const canarySchema = load('live-bulk-canary.schema.json');
const odcs = load('bulk-query.odcs.json');
const ajv = new Ajv2020({ allErrors: true, strict: false });
addFormats(ajv);
const validateSample = ajv.compile(sampleSchema);
const validateResult = ajv.compile(resultSchema);
const validateSummary = ajv.compile(summarySchema);
const validateCanary = ajv.compile(canarySchema);

test('bulk sample schema admits points and addresses, rejects ambiguity and malformed grid provenance', () => {
  const point = { id: 'grid:100mi:r1:c2', point: { latitude: 39, longitude: -105 },
    source_sha256: 'a'.repeat(64), cell: { crs: 'EPSG:6933', row: 1,
      column: 2, step_metres: 160934.4, land_fraction: 0.42,
      sample_kind: 'land_representative_point' } };
  assert.equal(validateSample(point), true, JSON.stringify(validateSample.errors));
  assert.equal(validateSample({ id: 'address_1', address: 'Fictional Demo Street',
    applicability_date: '2026-10-28', facts: { occupancy: 'business' } }), true);
  const invalid = [
    { ...point, address: 'second location' },
    { ...point, point: { latitude: 100, longitude: 0 } },
    { ...point, source_sha256: 'unsigned' },
    { ...point, cell: { ...point.cell, land_fraction: 0 } },
    { id: 'missing-location' },
  ];
  invalid.forEach((row) => assert.equal(validateSample(row), false));
});

test('bulk outcome and summary schemas preserve distinct unresolved and defects', () => {
  const outcome = { id: 'sample', category: 'unresolved', http_status: 422,
    error_code: 'boundary_ambiguous', applicability_date: '2026-10-28', latency_ms: 3.5 };
  assert.equal(validateResult(outcome), true, JSON.stringify(validateResult.errors));
  assert.equal(validateResult({ ...outcome, category: 'green' }), false);
  assert.equal(validateResult({ ...outcome, source_text: 'PRIVATE' }), false);
  const summary = { schema_version: '1', input_sha256: 'f'.repeat(64),
    endpoint_path: '/v1/project-code-basis', applicability_date: '2026-10-28',
    samples: 1, skipped_samples: 2, shard_index: 0, shard_count: 2,
    categories: { unresolved: 1 }, http_statuses: { '422': 1 },
    code_families: {}, duration_ms: 100, requests_per_second: 10,
    latency_p50_ms: 2, latency_p95_ms: 3, latency_p99_ms: 3 };
  assert.equal(validateSummary(summary), true, JSON.stringify(validateSummary.errors));
  assert.equal(validateSummary({ ...summary, input_sha256: 'unverified' }), false);
});

test('bulk query ODCS uses pinned schema and the public boundary rejects altered authority', () => {
  const validator = new Ajv2019({ allErrors: true, strict: false });
  addFormats(validator);
  const conforms = validator.compile(load('odcs-v3.1.0.schema.json'));
  assert.equal(conforms(odcs), true, JSON.stringify(conforms.errors));
  assert.deepEqual(validatePublicArtifact('schemas/bulk-query.odcs.json', JSON.stringify(odcs), {}), []);
  const invalid = structuredClone(odcs);
  invalid.authoritativeDefinitions[0].url = 'https://example.invalid/fake-source';
  assert.match(validatePublicArtifact('schemas/bulk-query.odcs.json', JSON.stringify(invalid), {}).join('; '),
    /must match reviewed schema-only identity/);
  assert.match(validatePublicArtifact('schemas/unreviewed-source.odcs.json', JSON.stringify(odcs), {}).join('; '),
    /unreviewed JSON data path/);
});

test('identity mismatch is a first-class error; canary never authorizes real-code claims', () => {
  const outcome = { id: 'g1', category: 'runtime_identity_mismatch',
    http_status: 200, applicability_date: '2026-10-28', latency_ms: 12 };
  assert.equal(validateResult(outcome), true, JSON.stringify(validateResult.errors));
  const canary = { schema_version: '1', category: 'runtime_canary_only',
    canary_input_sha256: 'a'.repeat(64), sample_id: 'canary-1', http_status: 200,
    expected_runtime_identity: { bundle_id: 'sha256:reviewed-bundle',
      source_commit: 'c'.repeat(40) }, project_verdict: 'not_verified',
    jurisdiction_id: 'CO-DENVER', code_families: ['building'],
    latency_ms: 22, legal_accuracy_reviewed: false, bulk_run_authorized: false };
  assert.equal(validateCanary(canary), true, JSON.stringify(validateCanary.errors));
  assert.equal(validateCanary({ ...canary, bulk_run_authorized: true }), false);
  assert.equal(validateCanary({ ...canary, legal_accuracy_reviewed: true }), false);
  assert.deepEqual(validatePublicArtifact('schemas/live-bulk-canary.schema.json',
    JSON.stringify(canarySchema), {}), []);
});

test('circuit breaker receipt must disclose partial input validation', () => {
  const summary = { schema_version: '1', input_sha256: 'f'.repeat(64),
    endpoint_path: '/v1/project-code-basis', applicability_date: '2026-10-29',
    samples: 4, skipped_samples: 0, shard_index: 0, shard_count: 1,
    categories: { service_error: 1, unresolved: 3 },
    http_statuses: { '422': 3, '503': 1 },
    code_families: {}, duration_ms: 1000, requests_per_second: 4,
    latency_p50_ms: 10, latency_p95_ms: 80, latency_p99_ms: 80,
    circuit_breaker_tripped: true, input_validation_complete: false,
    fatal_error_count: 1, abort_after_errors: 1 };
  assert.equal(validateSummary(summary), true, JSON.stringify(validateSummary.errors));
  assert.equal(validateSummary({ ...summary, abort_after_errors: -1 }), false);
  assert.equal(validateSummary({ ...summary, fatal_error_count: -1 }), false);
  assert.equal(validateSummary({ ...summary, secret: 'should not exist' }), false);
});
