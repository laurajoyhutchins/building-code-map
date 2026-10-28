import test from 'node:test';
import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import { execFileSync } from 'node:child_process';
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
