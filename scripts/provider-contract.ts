import { readFileSync } from 'node:fs';
import { resolve } from 'node:path';
import { fileURLToPath } from 'node:url';
import Ajv2019 from 'ajv/dist/2019.js';
import addFormats from 'ajv-formats';
import type { ErrorObject } from 'ajv';

type JsonRecord = Record<string, unknown>;

const upstream = JSON.parse(readFileSync(
  new URL('../schemas/odcs-v3.1.0.schema.json', import.meta.url), 'utf8',
)) as JsonRecord;
const profile = JSON.parse(readFileSync(
  new URL('../schemas/bcm-provider-contract-v1.schema.json', import.meta.url), 'utf8',
)) as JsonRecord;

// Compile the published ODCS schema unchanged and apply BCM's tighter profile.
// All references are local. Validation never downloads a remote schema or edits inputs.
const ajv = new Ajv2019({
  allErrors: true,
  strict: false,
  coerceTypes: false,
  useDefaults: false,
  removeAdditional: false,
});
addFormats(ajv);
ajv.addSchema(upstream, 'https://building-code-map.invalid/schemas/odcs-v3.1.0.schema.json');
const validate = ajv.compile(profile);

function record(value: unknown): value is JsonRecord {
  return typeof value === 'object' && value !== null && !Array.isArray(value);
}

function diagnosticPath(error: ErrorObject): string {
  const parts = error.instancePath.split('/').slice(1).map(
    (part) => part.replaceAll('~1', '/').replaceAll('~0', '~'),
  );
  if (error.keyword === 'required') parts.push(String(error.params.missingProperty));
  if (error.keyword === 'additionalProperties') {
    parts.push(String(error.params.additionalProperty));
  }
  return parts.reduce<string>((result, part) => {
    if (/^\d+$/.test(part)) return result + '[' + part + ']';
    return result ? result + '.' + part : part;
  }, '');
}

function schemaDiagnostic(error: ErrorObject, input: unknown): string {
  const path = diagnosticPath(error);
  if (error.keyword === 'additionalProperties' && error.instancePath === '') {
    return 'ODCS_FIELD_UNSUPPORTED ' + path;
  }
  if (path === 'apiVersion' && error.keyword === 'const') {
    return 'ODCS_API_VERSION_UNSUPPORTED apiVersion';
  }
  if (path === 'kind' && error.keyword === 'const') {
    return 'ODCS_KIND_INVALID kind';
  }
  if (path === 'customProperties' && error.keyword === 'contains') {
    return 'BCM_REQUIRED customProperties.bcmProfileVersion';
  }
  if (/^customProperties\[\d+\]\.value$/.test(path) && error.keyword === 'const') {
    const index = Number(path.match(/\[(\d+)\]/)?.[1]);
    const entry = record(input) && Array.isArray(input.customProperties)
      ? input.customProperties[index] : null;
    if (record(entry) && entry.property === 'bcmProfileVersion') {
      return 'BCM_PROFILE_VERSION_UNSUPPORTED customProperties.bcmProfileVersion';
    }
  }
  if (/^customProperties\[\d+\]\.property$/.test(path) && error.keyword === 'enum') {
    return 'BCM_EXTENSION_UNSUPPORTED ' + path;
  }
  if (/^(authoritativeDefinitions|customProperties)\[\d+\]\.id$/.test(path)
      && (error.keyword === 'pattern' || error.keyword === 'required'
        || error.keyword === 'type')) {
    return 'BCM_STABLE_ID_INVALID ' + path;
  }
  if (error.keyword === 'required') return 'BCM_REQUIRED ' + path;
  return 'BCM_SCHEMA_INVALID ' + (path || '$') + ' ' + (error.message || error.keyword);
}

function duplicateIds(input: unknown): string[] {
  if (!record(input)) return [];
  const errors: string[] = [];
  for (const field of ['authoritativeDefinitions', 'customProperties']) {
    const entries = input[field];
    if (!Array.isArray(entries)) continue;
    const seen = new Set<string>();
    for (const [index, entry] of entries.entries()) {
      if (!record(entry) || typeof entry.id !== 'string') continue;
      if (seen.has(entry.id)) {
        errors.push('BCM_STABLE_ID_DUPLICATE ' + field + '[' + index + '].id');
      }
      seen.add(entry.id);
    }
  }
  return errors;
}

export function validateProviderContract(contract: unknown): string[] {
  const isValid = validate(contract);
  const issues = isValid ? [] : (validate.errors ?? []).map(
    (error) => schemaDiagnostic(error, contract),
  );
  issues.push(...duplicateIds(contract));
  return [...new Set(issues)].sort();
}

export function validateProviderContractFile(path: string): string[] {
  let contract: unknown;
  try {
    contract = JSON.parse(readFileSync(path, 'utf8')) as unknown;
  } catch (error) {
    return ['BCM_CONTRACT_PARSE_ERROR ' + path + ' ' + String(error)];
  }
  return validateProviderContract(contract);
}

function main(paths: string[]): number {
  if (paths.length === 0) {
    console.error('BCM_USAGE provider-contract <contract.json> [...]');
    return 2;
  }
  const errors = paths.flatMap((path) =>
    validateProviderContractFile(path).map((message) => path + ': ' + message),
  );
  if (errors.length) {
    for (const error of errors) console.error(error);
    return 1;
  }
  return 0;
}

if (process.argv[1] && fileURLToPath(import.meta.url) === resolve(process.argv[1])) {
  process.exitCode = main(process.argv.slice(2));
}
