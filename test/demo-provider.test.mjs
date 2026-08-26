import test from 'node:test';
import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import { validateProviderContract } from '../scripts/provider-contract.mjs';

const demoDir = 'demo/DEMO-XX';

function readJson(name) {
  return JSON.parse(readFileSync(`${demoDir}/${name}`, 'utf8'));
}

function loadDemo() {
  return {
    contract: readJson('provider-contract.json'),
    bundle: readJson('provider-bundle.json'),
    project: readJson('project-code-basis.json'),
    readme: readFileSync(`${demoDir}/README.md`, 'utf8'),
  };
}

function customProperty(contract, property) {
  return contract.customProperties.find((item) => item.property === property)?.value;
}

function byId(items, id) {
  return items.find((item) => item.id === id);
}

test('DEMO-XX golden provider contract conforms to the BCM profile validator', () => {
  const { contract } = loadDemo();
  assert.deepEqual(validateProviderContract(contract), []);
});

test('provider contract, provider bundle, and project result share one synthetic bundle identity', () => {
  const { contract, bundle, project } = loadDemo();
  assert.equal(customProperty(contract, 'bcmProviderBundleId'), bundle.id);
  assert.equal(project.provenance.bundle_id, bundle.id);
  assert.equal(bundle.id, 'DEMO-BUNDLE');
});

test('provider bundle reuses the existing DEMO-XX exact-evidence identities', () => {
  const { bundle, project } = loadDemo();
  const exact = project.exact_evidence[0];

  assert.ok(byId(bundle.source_documents, exact.document.id));
  assert.ok(byId(bundle.source_artifacts, exact.artifact.id));
  assert.ok(byId(bundle.text_anchors, exact.anchor.id));
  assert.ok(byId(bundle.claims, exact.claim_id));
  assert.ok(byId(bundle.evidence_links, exact.id));
});

test('provider bundle uses stable synthetic IDs and internally resolves its representative references', () => {
  const { bundle } = loadDemo();
  const collections = [
    bundle.jurisdictions,
    bundle.authorities,
    bundle.regulatory_instruments,
    bundle.adoptions,
    bundle.amendments,
    bundle.code_families,
    bundle.code_editions,
    bundle.source_documents,
    bundle.source_artifacts,
    bundle.text_anchors,
    bundle.claims,
    bundle.evidence_links,
    bundle.source_policy_rules,
  ];

  for (const collection of collections) {
    for (const item of collection) assert.match(item.id, /^DEMO-[A-Za-z0-9_-]+$/);
  }

  const adoption = bundle.adoptions[0];
  assert.ok(byId(bundle.authorities, adoption.authority_id));
  assert.ok(byId(bundle.regulatory_instruments, adoption.instrument_id));
  assert.ok(byId(bundle.code_editions, adoption.code_edition_id));

  const amendment = bundle.amendments[0];
  assert.ok(byId(bundle.authorities, amendment.authority_id));
  assert.ok(byId(bundle.regulatory_instruments, amendment.instrument_id));
  assert.ok(byId(bundle.adoptions, amendment.adoption_id));

  const link = bundle.evidence_links[0];
  assert.ok(byId(bundle.claims, link.claim_id));
  assert.ok(byId(bundle.source_documents, link.document_id));
  assert.ok(byId(bundle.source_artifacts, link.artifact_id));
  assert.ok(byId(bundle.text_anchors, link.anchor_id));
});

test('DEMO-XX documentation explains contract, bundle, and project-result boundaries', () => {
  const { readme } = loadDemo();
  assert.match(readme, /provider-contract\.json/);
  assert.match(readme, /provider-bundle\.json/);
  assert.match(readme, /project-code-basis\.json/);
  assert.match(readme, /illustrative/i);
  assert.match(readme, /not normative/i);
  assert.match(readme, /no real regulatory coverage/i);
});
