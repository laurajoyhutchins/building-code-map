import test from 'node:test';
import assert from 'node:assert/strict';
import { mkdtempSync, readFileSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';
import { spawnSync } from 'node:child_process';

const fixturePath = 'test/fixtures/provider-contract/valid.json';
const valid = JSON.parse(readFileSync(fixturePath, 'utf8'));

function runContract(contract) {
  const directory = mkdtempSync(join(tmpdir(), 'bcm-provider-contract-'));
  const path = join(directory, 'contract.json');
  writeFileSync(path, `${JSON.stringify(contract, null, 2)}\n`);
  return spawnSync(process.execPath, ['scripts/provider-contract.mjs', path], {
    encoding: 'utf8',
  });
}

test('accepts the synthetic BCM provider contract', () => {
  const result = runContract(valid);
  assert.equal(result.status, 0, result.stderr || result.stdout);
  assert.equal(result.stderr, '');
});

test('requires purpose, usage, and limitations', () => {
  const contract = structuredClone(valid);
  delete contract.description.limitations;
  const result = runContract(contract);
  assert.notEqual(result.status, 0);
  assert.match(result.stderr, /BCM_REQUIRED description\.limitations/);
});

test('rejects unsupported BCM profile versions with a stable diagnostic', () => {
  const contract = structuredClone(valid);
  contract.customProperties.find((item) => item.property === 'bcmProfileVersion').value = '2.0.0';
  const result = runContract(contract);
  assert.notEqual(result.status, 0);
  assert.match(result.stderr, /BCM_PROFILE_VERSION_UNSUPPORTED customProperties\.bcmProfileVersion/);
});

test('rejects invalid stable IDs', () => {
  const contract = structuredClone(valid);
  contract.authoritativeDefinitions[0].id = 'bad id';
  const result = runContract(contract);
  assert.notEqual(result.status, 0);
  assert.match(result.stderr, /BCM_STABLE_ID_INVALID authoritativeDefinitions\[0\]\.id/);
});

test('rejects duplicate stable IDs where references are statically checkable', () => {
  const contract = structuredClone(valid);
  contract.authoritativeDefinitions[1].id = contract.authoritativeDefinitions[0].id;
  const result = runContract(contract);
  assert.notEqual(result.status, 0);
  assert.match(result.stderr, /BCM_STABLE_ID_DUPLICATE authoritativeDefinitions\[1\]\.id/);
});

test('rejects undeclared BCM extensions', () => {
  const contract = structuredClone(valid);
  contract.customProperties.push({
    id: 'bcm_secret_switch',
    property: 'bcmSecretSwitch',
    value: 'nope',
  });
  const result = runContract(contract);
  assert.notEqual(result.status, 0);
  assert.match(result.stderr, /BCM_EXTENSION_UNSUPPORTED customProperties\[2\]\.property/);
});

test('rejects fields outside the pinned ODCS root contract', () => {
  const contract = structuredClone(valid);
  contract.privateStoragePath = 's3://not-public/example';
  const result = runContract(contract);
  assert.notEqual(result.status, 0);
  assert.match(result.stderr, /ODCS_FIELD_UNSUPPORTED privateStoragePath/);
});
