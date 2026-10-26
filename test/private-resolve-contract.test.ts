import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import test from 'node:test';
import Ajv2020 from 'ajv/dist/2020.js';
import addFormats from 'ajv-formats';

const read = (path: string): object =>
  JSON.parse(readFileSync(new URL(path, import.meta.url), 'utf8')) as object;

const ajv = new Ajv2020({ strict: false, allErrors: true });
addFormats(ajv);
ajv.addSchema(read('../schemas/evidence.schema.json'));
const validate = ajv.compile(read('../schemas/private-resolve-response-v1.schema.json'));

const response = () => ({
  schema_version: '1',
  resolution: {
    status: 'verified',
    applicability_date: '2026-10-25',
    geography: { jurisdiction_id: 'DEMO-XX', jurisdiction_name: 'Fictional jurisdiction' },
    adoptions: [{
      code_family: 'building',
      verification: { status: 'verified' },
      exact_evidence: [{
        id: 'DEMO-LINK', claim_id: 'DEMO-CLAIM', role: 'establishes',
        document: { id: 'DEMO-DOC', kind: 'synthetic_ordinance', jurisdiction_id: 'DEMO-XX' },
        artifact: {
          id: 'DEMO-ARTIFACT', document_id: 'DEMO-DOC',
          sha256: 'a'.repeat(64), size_bytes: 123,
        },
        anchor: {
          id: 'DEMO-ANCHOR', artifact_id: 'DEMO-ARTIFACT',
          legal_locator: 'Fictional section 1', text_sha256: 'b'.repeat(64),
        },
      }],
    }],
  },
  provenance: { source_commit: 'DEMO-COMMIT', bundle_manifest_digest: 'DEMO-XX-BUNDLE' },
});

test('synthetic cross-repository JSON Schema accepts complete exact evidence', () => {
  assert.equal(validate(response()), true, JSON.stringify(validate.errors));
});

test('cross-repository JSON Schema rejects malformed exact source hashes', () => {
  const candidate = response();
  candidate.resolution.adoptions[0].exact_evidence[0].artifact.sha256 = 'not-a-digest';
  assert.equal(validate(candidate), false);
});

test('cross-repository JSON Schema rejects an invalid applicability date', () => {
  const candidate = response();
  candidate.resolution.applicability_date = 'tomorrow';
  assert.equal(validate(candidate), false);
});

test('schema version is pinned', () => {
  const candidate = response();
  candidate.schema_version = '2';
  assert.equal(validate(candidate), false);
});
