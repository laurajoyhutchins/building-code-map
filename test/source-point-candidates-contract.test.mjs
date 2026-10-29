import assert from 'node:assert/strict';
import { readFileSync } from 'node:fs';
import test from 'node:test';
import Ajv2020 from 'ajv/dist/2020.js';
import Ajv2019 from 'ajv/dist/2019.js';
import addFormats from 'ajv-formats';
import { validatePublicArtifact } from '../scripts/public-boundary.mjs';

const schema = JSON.parse(readFileSync('schemas/source-point-candidates-v1.schema.json','utf8'));
const contract = JSON.parse(readFileSync('schemas/source-point-candidates-v1.odcs.json','utf8'));
const validate = new Ajv2020({allErrors:true,strict:false});
addFormats(validate);
const accepts = validate.compile(schema);

test('polygon-source candidate is closed and cannot authorize adoption or AHJ', () => {
  const row = {
    schema_version:'1.0.0', classification:'source_polygon_candidate_only',
    source_geometry_sha256:'a'.repeat(64), source_layer_role:'municipal_boundary',
    source_features:23, point:{longitude:-105.2,latitude:40.2},
    candidate_feature_ids:['123','456'], boundary_feature_ids:[],
    status:'multiple_source_candidates',
    identity_review:'not_established', boundary_effective_date:'unknown',
    legal_authority:'not_established', regulatory_verdict_authorized:false,
    limitations:'Source polygon containment is not legal authority.',
  };
  assert.equal(accepts(row),true,JSON.stringify(accepts.errors));
  assert.equal(accepts({...row,legal_authority:'verified'}),false);
  assert.equal(accepts({...row,regulatory_verdict_authorized:true}),false);
  assert.equal(accepts({...row,adopted_code:'IBC'}),false);
  assert.equal(accepts({...row,source_layer_role:'verified_AHJ'}),false);
  assert.equal(accepts({...row,source_geometry_sha256:'unhashed'}),false);
  assert.equal(accepts({...row,point:{longitude:181,latitude:40}}),false);
});

test('ODCS 3.1 and source-polygon lookup schema are field-aligned', () => {
  const ajv = new Ajv2019({allErrors:true,strict:false});
  addFormats(ajv);
  assert.equal(ajv.compile(JSON.parse(readFileSync('schemas/odcs-v3.1.0.schema.json','utf8')))(contract),
    true,'paired ODCS must satisfy published contract');
  assert.deepEqual(contract.schema[0].properties.map(x=>x.name),Object.keys(schema.properties));
  assert.deepEqual(schema.required,Object.keys(schema.properties));
  assert.equal(contract.schema[0].properties.every(x=>x.required===true),true);
  assert.deepEqual(validatePublicArtifact('schemas/source-point-candidates-v1.odcs.json',
    JSON.stringify(contract), {}), []);
  const fake = structuredClone(contract);
  fake.authoritativeDefinitions[0].url = 'https://example.com/untrusted.json';
  assert.match(validatePublicArtifact('schemas/source-point-candidates-v1.odcs.json',
    JSON.stringify(fake), {}).join(';'),/must match reviewed schema-only identity/);
  assert.match(validatePublicArtifact('schemas/private-county-source.json',
    JSON.stringify(contract), {}).join(';'),/unreviewed JSON data path/);
});
