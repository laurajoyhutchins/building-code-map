import test from 'node:test';
import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import Ajv2020 from 'ajv/dist/2020.js';
import addFormats from 'ajv-formats';

const schema = JSON.parse(readFileSync('schemas/site-authority-policy.schema.json', 'utf8'));
const odcs = JSON.parse(readFileSync('schemas/site-authority-policy.odcs.json', 'utf8'));
const ajv = new Ajv2020({ allErrors: true, strict: false, validateFormats: true });
addFormats(ajv);
const validate = ajv.compile(schema);
const h = (x) => x.repeat(64);

const fictional = {
  schema_version: '1.0.0',
  rules: [{
    id: 'synthetic-only',
    location_sha256: h('a'),
    location_precision: 'address_point',
    country_code: 'US',
    jurisdiction_id: 'DEMO-XX',
    authority_id: 'synthetic-authority',
    effective_from: '2026-01-01',
    effective_to: '2026-12-31',
    review_state: 'reviewed',
    boundary: {
      object_uri: 'gs://synthetic-only/placeholder',
      object_generation: '1',
      sha256: h('b'),
      crs: 'EPSG:4326',
      site_relation: 'interior_reviewed',
      topology_review_sha256: h('c'),
      crosswalk_review_sha256: h('d'),
    },
    legal: {
      instrument_id: 'synthetic-ordinance',
      source_artifact_sha256: h('e'),
      exact_text_sha256: h('f'),
      legal_locator: 'Fictional §1',
      effective_from: '2026-01-01',
      effective_to: '2026-12-31',
      enforcement_authority_id: 'synthetic-authority',
      enforcement_mode: 'direct',
    },
    provenance: {
      entity_uri: 'urn:synthetic:source',
      activity_uri: 'urn:synthetic:review',
      agent_uri: 'urn:synthetic:agent',
      review_receipt_sha256: h('a'),
    },
  }],
};

test('site-authority cross-repo schema and ODCS pair describe the same gate', () => {
  assert.equal(schema.$schema, 'https://json-schema.org/draft/2020-12/schema');
  assert.equal(odcs.apiVersion, 'v3.1.0');
  assert.equal(odcs.kind, 'DataContract');
  assert.equal(odcs.version, '1.0.0');
  assert.ok(odcs.schema[0].properties.some((x) => x.name === 'rules[].legal.instrument_id'));
  assert.ok(odcs.schema[0].properties.some((x) => x.name === 'rules[].provenance.activity_uri'));
  assert.equal(validate(fictional), true, JSON.stringify(validate.errors));
});

test('untrusted site policy cannot bypass required boundary, legal or PROV fields', () => {
  const cases = [
    ['candidate review', (v) => { v.rules[0].review_state = 'candidate_only'; }],
    ['invented legal status', (v) => { v.rules[0].legal.enforcement_mode = 'unknown'; }],
    ['interpolated geocode', (v) => { v.rules[0].location_precision = 'interpolated'; }],
    ['missing GCS generation', (v) => { delete v.rules[0].boundary.object_generation; }],
    ['topology conflicts', (v) => { v.rules[0].boundary.site_relation = 'overlaps'; }],
    ['missing original law', (v) => { delete v.rules[0].legal.source_artifact_sha256; }],
    ['delegation without instrument', (v) => { v.rules[0].legal.enforcement_mode = 'delegated'; }],
    ['direct authority with fake delegation', (v) => { v.rules[0].legal.delegation_artifact_sha256 = h('b'); }],
    ['missing W3C PROV agent', (v) => { delete v.rules[0].provenance.agent_uri; }],
    ['digest truncated', (v) => { v.rules[0].provenance.review_receipt_sha256 = 'a'; }],
    ['unknown legal verdict injection', (v) => { v.rules[0].legal_verdict_authorized = true; }],
  ];
  for (const [name, edit] of cases) {
    const variant = structuredClone(fictional);
    edit(variant);
    assert.equal(validate(variant), false, name);
  }
});
