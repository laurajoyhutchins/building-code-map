import test from 'node:test';
import assert from 'node:assert/strict';
import fs from 'node:fs';
import os from 'node:os';
import path from 'node:path';
import {
  validatePolicy,
  validateSyntheticFixture,
  validateRepositoryPaths,
  validatePublicArtifact,
  validateRepository,
} from '../scripts/public-boundary.mjs';

const policy = {
  version: 1,
  default_disposition: 'review_required',
  synthetic_jurisdiction_prefix: 'DEMO-',
  forbidden_top_level_paths: ['corpus', 'production-data', 'research'],
  approved_top_level_paths: ['README.md', 'demo', 'docs', 'test', 'policy', 'schemas'],
  material_classes: {
    external_material: { public_repository: 'review_required' },
    production_regulatory_corpus: { public_repository: 'prohibited' },
    restricted_standard_model_code: { public_repository: 'prohibited' },
  },
};

test('external material cannot be admitted to the public repository by default', () => {
  assert.deepEqual(validatePolicy(policy), []);
  const unsafe = structuredClone(policy);
  unsafe.material_classes.external_material.public_repository = 'allowed';
  assert.match(validatePolicy(unsafe).join('\n'), /external_material.*review_required/);
});

test('production corpus and restricted standards cannot be distributed from the public repository', () => {
  const unsafe = structuredClone(policy);
  unsafe.material_classes.production_regulatory_corpus.public_repository = 'allowed';
  unsafe.material_classes.restricted_standard_model_code.public_repository = 'allowed';
  const errors = validatePolicy(unsafe).join('\n');
  assert.match(errors, /production_regulatory_corpus.*prohibited/);
  assert.match(errors, /restricted_standard_model_code.*prohibited/);
});

test('public fixtures must remain synthetic', () => {
  assert.deepEqual(validateSyntheticFixture({ jurisdiction_id: 'DEMO-XX' }, policy), []);
  assert.match(
    validateSyntheticFixture({ jurisdiction_id: 'CO-DENVER' }, policy).join('\n'),
    /synthetic jurisdiction prefix/,
  );
});

test('synthetic demo contracts and nested provider bundles do not require a root jurisdiction_id', () => {
  assert.deepEqual(
    validateSyntheticFixture({ id: 'bcm-provider-demo-xx' }, policy),
    [],
  );
  assert.deepEqual(
    validateSyntheticFixture({
      jurisdictions: [{ id: 'DEMO-XX' }],
      authorities: [{ id: 'DEMO-AUTHORITY-001', jurisdiction_id: 'DEMO-XX' }],
    }, policy),
    [],
  );
});

test('nested jurisdiction identities in public demo data must remain synthetic', () => {
  const errors = validateSyntheticFixture({
    jurisdictions: [{ id: 'CO-DENVER' }],
    authorities: [{ id: 'DEMO-AUTHORITY-001', jurisdiction_id: 'CO-DENVER' }],
  }, policy).join('\n');
  assert.match(errors, /CO-DENVER/);
  assert.match(errors, /synthetic jurisdiction prefix/);
});

test('production-data paths cannot appear at repository top level', () => {
  assert.deepEqual(validateRepositoryPaths(['README.md', 'demo', 'docs'], policy), []);
  assert.match(
    validateRepositoryPaths(['README.md', 'corpus', 'docs'], policy).join('\n'),
    /forbidden top-level path: corpus/,
  );
});

test('unreviewed directories are not silently public-safe', () => {
  assert.match(
    validateRepositoryPaths(['README.md', 'new-regulatory-dataset'], policy).join('\n'),
    /unreviewed public top-level path/,
  );
});

test('non-fictional jurisdiction data fails outside demo and in alternate fixture directories', () => {
  const fixture = JSON.stringify({ jurisdiction_id: 'US-CO', status: 'resolved' });
  for (const place of [
    'test/fixtures/overlooked.json',
    'docs/hidden-corpus.json',
    'unknown-directory/hidden.json',
  ]) {
    assert.notEqual(validatePublicArtifact(place, fixture, policy).length, 0, place);
  }
  assert.deepEqual(
    validatePublicArtifact(
      'test/fixtures/conforming.json',
      JSON.stringify({ jurisdiction_id: 'DEMO-XX', description: 'synthetic fixture' }),
      policy,
    ),
    [],
  );
});

test('regulated attachments and disguised JSON schemas do not bypass review', () => {
  assert.match(
    validatePublicArtifact('docs/corpus/source.txt', '', policy).join('\n'),
    /prohibited corpus directory/,
  );
  assert.match(
    validatePublicArtifact('test/fixtures/agency.pdf', '', policy).join('\n'),
    /private custody/,
  );
  assert.match(
    validatePublicArtifact('schemas/regulation.schema.json', '{"jurisdiction_id":"US-CO"}', policy).join('\n'),
    /actual JSON Schema/,
  );
});

test('repository-wide validation rejects alternate-path regulatory fixture', (t) => {
  const root = fs.mkdtempSync(path.join(os.tmpdir(), 'bcm-public-boundary-'));
  t.after(() => fs.rmSync(root, { recursive: true, force: true }));
  fs.mkdirSync(path.join(root, 'policy'));
  fs.mkdirSync(path.join(root, 'test', 'fixtures'), { recursive: true });
  fs.writeFileSync(path.join(root, 'policy', 'public-boundary.json'), JSON.stringify(policy));
  fs.writeFileSync(
    path.join(root, 'test', 'fixtures', 'leak.json'),
    JSON.stringify({ jurisdiction_id: 'US-CO', title: 'fictional red-team sentinel' }),
  );
  assert.match(
    validateRepository(root).join('\n'),
    /test\/fixtures\/leak.json.*synthetic jurisdiction prefix/,
  );
});

test('only exact project-owned geocoder ranking policy can be public, never a GIS source', () => {
  const config = {
    version: 'geocoder-ranking-1.0',
    address_point_base: 0.7,
    street_range_base: 0.55,
    exact_street: 0.05,
    exact_city: 0.15,
    exact_postal_code: 0.1,
    postal_not_supplied: 0.05,
    parity_matched: 0.05,
    minimum_address_point_quality: 0.85,
    minimum_street_range_quality: 0.75,
    ambiguity_gap: 0.05,
    source_priority: {},
  };
  const official = { ...config, postal_code_not_supplied: config.postal_not_supplied };
  delete official.postal_not_supplied;
  assert.deepEqual(
    validatePublicArtifact('geocoder/config/geocoder-ranking-policy.json', JSON.stringify(official), policy),
    [],
  );
  assert.match(
    validatePublicArtifact('geocoder/config/municipal-boundaries.json', JSON.stringify(official), policy).join('\n'),
    /unreviewed JSON data path/,
  );
  assert.match(
    validatePublicArtifact(
      'geocoder/config/geocoder-ranking-policy.json',
      JSON.stringify({ ...official, source_priority: { 'private-source': 0.9 } }),
      policy,
    ).join('\n'),
    /source priorities/,
  );
  assert.match(
    validatePublicArtifact(
      'geocoder/config/geocoder-ranking-policy.json',
      JSON.stringify({ ...official, claim_kind: 'code_adoption' }),
      policy,
    ).join('\n'),
    /unreviewed fields/,
  );
});
