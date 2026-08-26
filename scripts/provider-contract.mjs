import { readFileSync } from 'node:fs';
import { fileURLToPath } from 'node:url';

const SUPPORTED_PROFILE_VERSION = '1.0.0';
const STABLE_ID = /^[A-Za-z0-9_-]+$/;

// ODCS v3.1.0 root properties, pinned from:
// bitol-io/open-data-contract-standard@v3.1.0
// schema blob 3596e68b07d19fc44fe947d8796021b64b8f86eb.
const ODCS_ROOT_FIELDS = new Set([
  'version',
  'kind',
  'apiVersion',
  'id',
  'name',
  'tenant',
  'tags',
  'status',
  'servers',
  'dataProduct',
  'description',
  'domain',
  'schema',
  'support',
  'price',
  'team',
  'roles',
  'slaDefaultElement',
  'slaProperties',
  'authoritativeDefinitions',
  'customProperties',
  'contractCreatedTs',
]);

const REQUIRED_ROOT_FIELDS = ['version', 'apiVersion', 'kind', 'id', 'status', 'name'];
const REQUIRED_DESCRIPTION_FIELDS = ['purpose', 'usage', 'limitations'];
const BCM_EXTENSION_PROPERTIES = new Set([
  'bcmProfileVersion',
  'bcmProviderBundleId',
]);

function diagnostic(code, path, detail = '') {
  return `${code} ${path}${detail ? ` ${detail}` : ''}`;
}

function isObject(value) {
  return value !== null && typeof value === 'object' && !Array.isArray(value);
}

function validateStableId(value, path, errors) {
  if (typeof value !== 'string' || !STABLE_ID.test(value)) {
    errors.push(diagnostic('BCM_STABLE_ID_INVALID', path));
    return false;
  }
  return true;
}

function validateUniqueStableIds(items, path, errors) {
  if (!Array.isArray(items)) return;
  const seen = new Set();
  for (let index = 0; index < items.length; index += 1) {
    const item = items[index];
    if (!isObject(item)) continue;
    const itemPath = `${path}[${index}].id`;
    if (!validateStableId(item.id, itemPath, errors)) continue;
    if (seen.has(item.id)) {
      errors.push(diagnostic('BCM_STABLE_ID_DUPLICATE', itemPath));
    } else {
      seen.add(item.id);
    }
  }
}

export function validateProviderContract(contract) {
  const errors = [];

  if (!isObject(contract)) {
    return [diagnostic('BCM_CONTRACT_INVALID', '$', 'expected object')];
  }

  for (const field of Object.keys(contract)) {
    if (!ODCS_ROOT_FIELDS.has(field)) {
      errors.push(diagnostic('ODCS_FIELD_UNSUPPORTED', field));
    }
  }

  for (const field of REQUIRED_ROOT_FIELDS) {
    if (contract[field] === undefined || contract[field] === null || contract[field] === '') {
      errors.push(diagnostic('BCM_REQUIRED', field));
    }
  }

  if (contract.apiVersion !== undefined && contract.apiVersion !== 'v3.1.0') {
    errors.push(diagnostic('ODCS_API_VERSION_UNSUPPORTED', 'apiVersion'));
  }
  if (contract.kind !== undefined && contract.kind !== 'DataContract') {
    errors.push(diagnostic('ODCS_KIND_INVALID', 'kind'));
  }
  if (contract.id !== undefined) validateStableId(contract.id, 'id', errors);

  if (!isObject(contract.description)) {
    errors.push(diagnostic('BCM_REQUIRED', 'description'));
  } else {
    for (const field of REQUIRED_DESCRIPTION_FIELDS) {
      const value = contract.description[field];
      if (typeof value !== 'string' || value.trim() === '') {
        errors.push(diagnostic('BCM_REQUIRED', `description.${field}`));
      }
    }
  }

  if (!Array.isArray(contract.authoritativeDefinitions) || contract.authoritativeDefinitions.length === 0) {
    errors.push(diagnostic('BCM_REQUIRED', 'authoritativeDefinitions'));
  } else {
    validateUniqueStableIds(contract.authoritativeDefinitions, 'authoritativeDefinitions', errors);
  }

  if (!Array.isArray(contract.customProperties) || contract.customProperties.length === 0) {
    errors.push(diagnostic('BCM_REQUIRED', 'customProperties'));
  } else {
    validateUniqueStableIds(contract.customProperties, 'customProperties', errors);
    let profileVersion = null;
    for (let index = 0; index < contract.customProperties.length; index += 1) {
      const item = contract.customProperties[index];
      if (!isObject(item)) {
        errors.push(diagnostic('BCM_EXTENSION_INVALID', `customProperties[${index}]`));
        continue;
      }
      if (typeof item.property !== 'string' || item.property.trim() === '') {
        errors.push(diagnostic('BCM_REQUIRED', `customProperties[${index}].property`));
        continue;
      }
      if (item.property.startsWith('bcm') && !BCM_EXTENSION_PROPERTIES.has(item.property)) {
        errors.push(diagnostic('BCM_EXTENSION_UNSUPPORTED', `customProperties[${index}].property`));
      }
      if (item.property === 'bcmProfileVersion') profileVersion = item.value;
    }

    if (profileVersion === null) {
      errors.push(diagnostic('BCM_REQUIRED', 'customProperties.bcmProfileVersion'));
    } else if (profileVersion !== SUPPORTED_PROFILE_VERSION) {
      errors.push(diagnostic('BCM_PROFILE_VERSION_UNSUPPORTED', 'customProperties.bcmProfileVersion'));
    }
  }

  return errors;
}

export function validateProviderContractFile(path) {
  let contract;
  try {
    contract = JSON.parse(readFileSync(path, 'utf8'));
  } catch (error) {
    return [diagnostic('BCM_CONTRACT_PARSE_ERROR', path, error.message)];
  }
  return validateProviderContract(contract);
}

function main(argv) {
  const paths = argv.slice(2);
  if (paths.length === 0) {
    console.error('BCM_USAGE provider-contract <contract.json> [...]');
    return 2;
  }

  const errors = [];
  for (const path of paths) {
    for (const error of validateProviderContractFile(path)) errors.push(`${path}: ${error}`);
  }

  if (errors.length > 0) {
    for (const error of errors) console.error(error);
    return 1;
  }
  return 0;
}

if (process.argv[1] && fileURLToPath(import.meta.url) === process.argv[1]) {
  process.exitCode = main(process.argv);
}
