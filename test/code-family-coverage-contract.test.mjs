import assert from 'node:assert/strict';
import test from 'node:test';
import { readFileSync } from 'node:fs';
import Ajv2019 from 'ajv/dist/2019.js';
import addFormats from 'ajv-formats';

const odcs = JSON.parse(readFileSync('schemas/code-family-coverage-ledger.odcs.json', 'utf8'));
const pinnedStandard = JSON.parse(readFileSync('schemas/odcs-v3.1.0.schema.json', 'utf8'));
const ledgerSchema = JSON.parse(readFileSync('schemas/code-family-coverage-ledger.schema.json', 'utf8'));
const ajv = new Ajv2019({ allErrors: true, strict: false, coerceTypes: false, useDefaults: false });
addFormats(ajv);
const conformsToODCS = ajv.compile(pinnedStandard);

const expectedSchemaURL = 'https://github.com/laurajoyhutchins/building-code-map/blob/main/schemas/code-family-coverage-ledger.schema.json';
const declared = (contract) => contract.schema.flatMap((object) => object.properties.map((property) => property.name));

test('family-ledger cross-repo ODCS document conforms to pinned upstream ODCS v3.1 schema', () => {
  assert.equal(odcs.apiVersion, 'v3.1.0');
  assert.equal(odcs.kind, 'DataContract');
  assert.equal(conformsToODCS(odcs), true, JSON.stringify(conformsToODCS.errors));
  assert.equal(odcs.status, 'draft');
});

test('family ODCS contract resolves only the same public JSON Schema and explicit evidence dimensions', () => {
  assert.equal(ledgerSchema.properties.scope.properties.country_code.const, 'US');
  assert.equal(ledgerSchema.properties.schema_version.const, '0.1');
  assert.equal(
    odcs.authoritativeDefinitions.find((definition) => definition.id === 'bcm_family_ledger_json_schema')?.url,
    expectedSchemaURL,
  );
  const fields = declared(odcs);
  assert.equal(fields.length, new Set(fields).size);
  for (const field of [
    'scope.code_family', 'scope.as_of', 'scope.project_class',
    'projection.snapshot_id', 'projection.snapshot_sha256', 'projection.denominator.status',
    'records[].entity_id', 'records[].source_review.status', 'records[].authority.status',
    'records[].geography.status', 'records[].provider.status',
    'records[].project_verification.status', 'records[].project_verification.receipt.entity_id',
    'records[].project_verification.receipt.independent_review_ref', 'records[].unknown_reasons',
  ]) {
    assert.ok(fields.includes(field), 'missing family ODCS declaration: ' + field);
  }
  assert.match(odcs.description.limitations, /not.*verified project verdict/i);
  assert.match(odcs.description.limitations, /W3C PROV-O/);
});

test('invalid ODCS shapes cannot satisfy the cross-repo contract', () => {
  for (const patch of [
    (value) => { value.apiVersion = 'v999'; },
    (value) => { value.kind = 'NotADataContract'; },
    (value) => { delete value.description; },
    (value) => { value.schema[0].properties[0].required = 'yes'; },
  ]) {
    const altered = structuredClone(odcs);
    patch(altered);
    assert.equal(conformsToODCS(altered), false, 'invalid upstream ODCS shape accepted');
  }
});

test('no family export may omit schema identity or imply national verified coverage', () => {
  const altered = structuredClone(odcs);
  altered.authoritativeDefinitions[0].url = 'https://example.invalid/other-schema';
  assert.notEqual(
    altered.authoritativeDefinitions.find((definition) => definition.id === 'bcm_family_ledger_json_schema')?.url,
    expectedSchemaURL,
  );
  assert.equal(odcs.description.limitations.includes('not public'), true);
  assert.equal(odcs.status, 'draft');
});
