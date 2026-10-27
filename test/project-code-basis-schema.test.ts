import test from 'node:test';
import assert from 'node:assert/strict';
import { execFileSync } from 'node:child_process';
import { readFileSync } from 'node:fs';
import Ajv2020 from 'ajv/dist/2020.js';
import addFormats from 'ajv-formats';

const schema = JSON.parse(readFileSync('schemas/project-code-basis.schema.json', 'utf8'));
const evidenceSchema = JSON.parse(readFileSync('schemas/evidence.schema.json', 'utf8'));
const ajv = new Ajv2020({ allErrors: true, strict: false, validateFormats: true });
addFormats(ajv);
ajv.addSchema(evidenceSchema);
const conforms = ajv.compile(schema);

// This executable uses the actual Provider -> Runtime -> ProjectVerifier ->
// HTTP handler. A static JSON golden does not establish Go/schema agreement.
function canonicalHTTPBasis() {
  const output = execFileSync('go', ['run', './test/runtime-schema-fixture'], {
    encoding: 'utf8',
    timeout: 60000,
    maxBuffer: 1024 * 1024,
  });
  return JSON.parse(output);
}

test('canonical runtime-to-HTTP output conforms to public project-code-basis schema', () => {
  const actual = canonicalHTTPBasis();
  assert.equal(actual.schema_version, '0.3');
  assert.equal(actual.project_id, 'DEMO-SCHEMA');
  assert.equal(actual.verdict, 'conditional');
  assert.equal(actual.derived_facts[0].key, 'jurisdiction');
  assert.equal(actual.evidence[0].id, 'DEMO-SOURCE');
  assert.equal(conforms(actual), true, JSON.stringify(conforms.errors));
});

test('the same published schema rejects malformed derived facts and ordinary evidence', () => {
  const actual = canonicalHTTPBasis();
  assert.equal(conforms(actual), true, JSON.stringify(conforms.errors));
  for (const [name, mutate, expectedPath] of [
    ['empty derived key', (basis) => { basis.derived_facts[0].key = ''; }, '/derived_facts/0/key'],
    ['empty derived value', (basis) => { basis.derived_facts[0].value = ''; }, '/derived_facts/0/value'],
    ['empty evidence id', (basis) => { basis.evidence[0].id = ''; }, '/evidence/0/id'],
    ['empty evidence kind', (basis) => { basis.evidence[0].kind = ''; }, '/evidence/0/kind'],
  ]) {
    const mutant = structuredClone(actual);
    mutate(mutant);
    assert.equal(conforms(mutant), false, name);
    assert.ok(conforms.errors?.some((item) => item.instancePath === expectedPath), name);
  }
});
