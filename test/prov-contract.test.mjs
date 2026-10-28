import test from 'node:test';
import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import { execFileSync, spawnSync } from 'node:child_process';
import Ajv2020 from 'ajv/dist/2020.js';
import Ajv2019 from 'ajv/dist/2019.js';
import addFormats from 'ajv-formats';

const schema = JSON.parse(readFileSync('schemas/prov-graph.schema.json', 'utf8'));
const contract = JSON.parse(readFileSync('schemas/prov-graph.odcs.json', 'utf8'));
const ajv = new Ajv2020({ allErrors: true, strict: false, validateFormats: true });
addFormats(ajv);
const validate = ajv.compile(schema);
const context = {
  prov: 'http://www.w3.org/ns/prov#',
  bcm: 'https://github.com/laurajoyhutchins/building-code-map/ns/provenance#',
  xsd: 'http://www.w3.org/2001/XMLSchema#',
};
function graph() {
  return {
    '@context': context,
    '@graph': [
      {
        '@id': 'urn:bcm:prov:document:1',
        '@type': 'prov:Entity',
        'bcm:sourceID': 'DEMO-DOC',
      },
      {
        '@id': 'urn:bcm:prov:artifact:2',
        '@type': 'prov:Entity',
        'bcm:sha256': 'a'.repeat(64),
        'prov:wasDerivedFrom': [{ '@id': 'urn:bcm:prov:document:1' }],
      },
      {
        '@id': 'urn:bcm:prov:activity:3',
        '@type': 'prov:Activity',
        'prov:used': [{ '@id': 'urn:bcm:prov:artifact:2' }],
        'prov:startedAtTime': { '@value': '2026-10-28T01:00:00Z', '@type': 'xsd:dateTime' },
      },
    ],
  };
}

test('export schema accepts explicit typed PROV-O edges', () => {
  assert.equal(validate(graph()), true, JSON.stringify(validate.errors));
  assert.equal(schema.properties['@context'].properties.prov.const,
    'http://www.w3.org/ns/prov#');
});

test('export schema rejects invalid PROV syntax and fabricated authority', () => {
  const variants = [
    (x) => { x['@context'].prov = 'https://fake.example/prov#'; },
    (x) => { x['@graph'][0]['@type'] = 'prov:VerifiedClaim'; },
    (x) => { x['@graph'][1]['bcm:sha256'] = 'hash-not-pinned'; },
    (x) => { x['@graph'][2]['prov:used'] = ['urn:bcm:prov:artifact:2']; },
    (x) => { x['@graph'][2]['prov:startedAtTime']['@type'] = 'xsd:string'; },
    (x) => { x['@graph'][0]['prov:wasApprovedBy'] = [{ '@id': 'urn:bcm:prov:activity:3' }]; },
    (x) => { x['@graph'][2]['prov:wasDerivedFrom'] = [{ '@id': 'urn:bcm:prov:artifact:2' }]; },
  ];
  for (const [index, mutate] of variants.entries()) {
    const value = structuredClone(graph());
    mutate(value);
    assert.equal(validate(value), false, 'invalid PROV variant ' + index);
  }
});

test('paired ODCS contract validates against pinned public ODCS grammar', () => {
  const grammar = JSON.parse(readFileSync('schemas/odcs-v3.1.0.schema.json', 'utf8'));
  const validator = new Ajv2019({ allErrors: true, strict: false });
  addFormats(validator);
  const valid = validator.compile(grammar);
  assert.equal(valid(contract), true, JSON.stringify(valid.errors));
  assert.equal(contract.apiVersion, 'v3.1.0');
  assert.equal(contract.authoritativeDefinitions[0].url.endsWith('/schemas/prov-graph.schema.json'), true);
});

test('published Go source and tests preserve independent evidence gating', () => {
  const source = readFileSync('engine/prov.go', 'utf8');
  assert.match(source, /validateEvidenceLink/);
  assert.match(source, /bcm:hasEvidenceLink/);
  assert.doesNotMatch(source, /VerifyClaimEvidence\(/);
  assert.doesNotMatch(source, /EvidenceVerified/);
  const core = readFileSync('provenance/graph.go', 'utf8');
  assert.match(core, /prov:wasGeneratedBy/);
  assert.match(core, /prov:wasAssociatedWith/);
  assert.match(core, /prov:used/);
  assert.doesNotMatch(core, /http\.Get\(/);
});

test('the real Go exporter emits schema-valid PROV-O from a canonical synthetic claim', () => {
  const fixture = readFileSync('test/fixtures/prov/claim.json', 'utf8');
  const run = () => execFileSync('go', ['run', './cmd/bcm-prov'], {
    input: fixture, encoding: 'utf8', timeout: 60000,
  });
  const first = run();
  assert.equal(first, run(), 'source graph serialization must be deterministic');
  const actual = JSON.parse(first);
  assert.equal(validate(actual), true, JSON.stringify(validate.errors));
  const entities = actual['@graph'];
  assert.equal(entities.length, 5);
  assert.ok(entities.some((node) => node['bcm:sourceID'] === 'DEMO-ARTIFACT-1' &&
    node['bcm:sha256'] === 'a'.repeat(64)));
  assert.ok(entities.some((node) => node['bcm:sourceID'] === 'DEMO-LINK-1' &&
    node['bcm:evidenceRole'] === 'establishes'));
  assert.ok(entities.every((node) => node['@type'] === 'prov:Entity'),
    'a source claim cannot fabricate review agents or activities');
});

test('the real Go CLI rejects malformed, tampered and trailing source observations', () => {
  const valid = JSON.parse(readFileSync('test/fixtures/prov/claim.json', 'utf8'));
  for (const mutate of [
    (x) => { x.links[0].artifact.sha256 = 'malformed'; },
    (x) => { x.links[0].document.jurisdiction_id = 'US-CO'; },
    (x) => { x.links[0].role = 'verified_by_ai'; },
    (x) => { x.extraUnreviewedClaim = true; },
  ]) {
    const trial = structuredClone(valid);
    mutate(trial);
    const result = spawnSync('go', ['run', './cmd/bcm-prov'], {
      input: JSON.stringify(trial), encoding: 'utf8', timeout: 60000,
    });
    assert.notEqual(result.status, 0, 'unverified input must fail closed');
    assert.equal(result.stdout, '', 'invalid source must not produce a PROV graph');
  }
  const extra = spawnSync('go', ['run', './cmd/bcm-prov'], {
    input: JSON.stringify(valid) + ' ' + JSON.stringify({}), encoding: 'utf8', timeout: 60000,
  });
  assert.notEqual(extra.status, 0);
});

test('external JSON-LD is admitted only through whole-graph Go validation', () => {
  const valid = JSON.parse(readFileSync('test/fixtures/prov/actor.jsonld', 'utf8'));
  assert.equal(validate(valid), true, JSON.stringify(validate.errors));
  const run = (value) => spawnSync('go', ['run', './cmd/bcm-prov', 'validate'], {
    input: JSON.stringify(value), encoding: 'utf8', timeout: 60000,
  });
  const passed = run(valid);
  assert.equal(passed.status, 0, passed.stderr);
  assert.equal(validate(JSON.parse(passed.stdout)), true);
  for (const mutate of [
    (x) => { x['@graph'][2]['prov:used'][0]['@id'] = 'urn:bcm:prov:missing'; },
    (x) => { x['@graph'][1]['prov:wasDerivedFrom'][0]['@id'] = 'urn:bcm:prov:demo:activity'; },
    (x) => { x['@graph'][3]['@id'] = x['@graph'][0]['@id']; },
  ]) {
    const mutant = structuredClone(valid);
    mutate(mutant);
    // This proves the original defect: local structure alone lacks global closure.
    assert.equal(validate(mutant), true, JSON.stringify(validate.errors));
    const result = run(mutant);
    assert.notEqual(result.status, 0, 'whole-graph import must reject structural false positive');
    assert.equal(result.stdout, '', 'invalid graph must not produce a successful output');
  }
});
