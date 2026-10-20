import fs from 'node:fs';
import path from 'node:path';
import { pathToFileURL } from 'node:url';

const REQUIRED_PUBLIC_DISPOSITIONS = {
  external_material: 'review_required',
  production_regulatory_corpus: 'prohibited',
  restricted_standard_model_code: 'prohibited',
};

export function validatePolicy(policy) {
  const errors = [];

  if (policy?.version !== 1) {
    errors.push('policy.version must be 1');
  }
  if (policy?.default_disposition !== 'review_required') {
    errors.push('default_disposition must be review_required');
  }
  if (typeof policy?.synthetic_jurisdiction_prefix !== 'string' || !policy.synthetic_jurisdiction_prefix) {
    errors.push('synthetic_jurisdiction_prefix must be a non-empty string');
  }
  if (!Array.isArray(policy?.forbidden_top_level_paths)) {
    errors.push('forbidden_top_level_paths must be an array');
  }
  if (!Array.isArray(policy?.approved_top_level_paths) || policy.approved_top_level_paths.length === 0) {
    errors.push('approved_top_level_paths must be a non-empty array');
  }

  for (const [materialClass, expected] of Object.entries(REQUIRED_PUBLIC_DISPOSITIONS)) {
    const actual = policy?.material_classes?.[materialClass]?.public_repository;
    if (actual !== expected) {
      errors.push(`${materialClass}.public_repository must be ${expected}`);
    }
  }

  return errors;
}

function collectJurisdictionIds(value, valuePath = '$', results = []) {
  if (Array.isArray(value)) {
    value.forEach((item, index) => collectJurisdictionIds(item, `${valuePath}[${index}]`, results));
    return results;
  }
  if (value === null || typeof value !== 'object') return results;

  for (const [key, child] of Object.entries(value)) {
    const childPath = valuePath === '$' ? key : `${valuePath}.${key}`;
    if (key === 'jurisdiction_id' || key === 'primary_jurisdiction_id' || key.endsWith('_jurisdiction_id')) {
      results.push({ path: childPath, value: child });
    }
    if (key === 'jurisdictions' && Array.isArray(child)) {
      child.forEach((jurisdiction, index) => {
        if (jurisdiction && typeof jurisdiction === 'object' && !Array.isArray(jurisdiction) && 'id' in jurisdiction) {
          results.push({ path: `${childPath}[${index}].id`, value: jurisdiction.id });
        }
      });
    }
    collectJurisdictionIds(child, childPath, results);
  }

  return results;
}

export function validateSyntheticFixture(fixture, policy) {
  const errors = [];
  const prefix = policy?.synthetic_jurisdiction_prefix;

  for (const identity of collectJurisdictionIds(fixture)) {
    if (typeof identity.value !== 'string' || !identity.value.startsWith(prefix)) {
      errors.push(
        `fixture ${identity.path} ${String(identity.value)} must use synthetic jurisdiction prefix ${prefix}`,
      );
    }
  }

  return errors;
}

// The public tree is a reviewed software distribution, never an implicit
// regulatory corpus distribution. New data paths require a policy review.
const NON_REPOSITORY_DIRECTORIES = new Set(['.git', 'node_modules', 'dist', 'coverage']);
const PROHIBITED_FILE_EXTENSIONS = new Set([
  '.pdf', '.sqlite', '.duckdb', '.parquet', '.geojson', '.csv', '.zip', '.tar',
  '.gz', '.db', '.ndjson',
]);
const SYNTHETIC_JSON_ROOTS = ['demo/', 'test/fixtures/', 'bcastconsumer/testdata/'];

export function validateRepositoryPaths(topLevelPaths, policy) {
  const errors = [];
  const forbidden = new Set(policy?.forbidden_top_level_paths ?? []);
  const approved = new Set(policy?.approved_top_level_paths ?? []);

  for (const entry of topLevelPaths) {
    if (NON_REPOSITORY_DIRECTORIES.has(entry)) continue;
    if (forbidden.has(entry)) {
      errors.push(`forbidden top-level path: ${entry}`);
    } else if (!approved.has(entry)) {
      errors.push(`unreviewed public top-level path: ${entry}`);
    }
  }

  return errors;
}

function collectRepositoryFiles(root, dir = root, result = []) {
  for (const entry of fs.readdirSync(dir, { withFileTypes: true })) {
    if (NON_REPOSITORY_DIRECTORIES.has(entry.name)) continue;
    const fullPath = path.join(dir, entry.name);
    const relativePath = path.relative(root, fullPath).split(path.sep).join('/');
    if (entry.isSymbolicLink()) {
      result.push({ path: relativePath, symlink: true });
    } else if (entry.isDirectory()) {
      collectRepositoryFiles(root, fullPath, result);
    } else if (entry.isFile()) {
      result.push({ path: relativePath, symlink: false });
    }
  }
  return result;
}

// Only specifically owned first-party JSON metadata and schemas, or synthetic
// fixtures, have an implicit public disposition. Every other JSON data path
// needs an explicit policy change and rights review.
function jsonDisposition(relativePath) {
  if (relativePath === 'package.json' || relativePath === 'policy/public-boundary.json') {
    return 'first-party';
  }
  if (relativePath.startsWith('schemas/') && relativePath.endsWith('.schema.json')) {
    return 'schema';
  }
  if (SYNTHETIC_JSON_ROOTS.some((root) => relativePath.startsWith(root))) {
    return 'synthetic';
  }
  return 'review-required';
}

export function validatePublicArtifact(relativePath, contents, policy) {
  const errors = [];
  const normalized = relativePath.replaceAll('\\', '/');
  const parts = normalized.split('/');
  if (parts.some((part) => (policy?.forbidden_top_level_paths ?? []).includes(part))) {
    errors.push('prohibited corpus directory appears within the public path');
  }
  if (PROHIBITED_FILE_EXTENSIONS.has(path.posix.extname(normalized).toLowerCase())) {
    errors.push('data or external-source binary requires private custody and rights review');
  }
  if (!normalized.endsWith('.json')) return errors;

  const disposition = jsonDisposition(normalized);
  if (disposition === 'review-required') {
    errors.push('unreviewed JSON data path requires explicit public disposition');
    return errors;
  }
  let value;
  try {
    value = JSON.parse(contents);
  } catch {
    errors.push('JSON content is invalid');
    return errors;
  }
  if (disposition === 'schema') {
    if (typeof value?.$schema !== 'string' || !value.$schema.includes('json-schema.org')) {
      errors.push('a public schema path must contain an actual JSON Schema');
    }
  }
  if (disposition === 'synthetic') {
    errors.push(...validateSyntheticFixture(value, policy));
    if (!/DEMO-|synthetic|fictional/i.test(contents)) {
      errors.push('fixture lacks an explicit synthetic marker');
    }
  }
  return errors;
}

export function validateRepository(rootDir = process.cwd()) {
  const policyPath = path.join(rootDir, 'policy', 'public-boundary.json');
  const policy = JSON.parse(fs.readFileSync(policyPath, 'utf8'));
  const errors = [...validatePolicy(policy)];

  const topLevelPaths = fs.readdirSync(rootDir, { withFileTypes: true }).map((entry) => entry.name);
  errors.push(...validateRepositoryPaths(topLevelPaths, policy));

  for (const entry of collectRepositoryFiles(rootDir)) {
    if (entry.symlink) {
      errors.push(`${entry.path}: symlink not permitted in the reviewed public tree`);
      continue;
    }
    // Only inspect source text for data-shaped files. Other files have their
    // public ownership controlled by the approved top-level source roots.
    const extension = path.posix.extname(entry.path).toLowerCase();
    const needsContent = extension === '.json';
    const contents = needsContent ? fs.readFileSync(path.join(rootDir, entry.path), 'utf8') : '';
    for (const message of validatePublicArtifact(entry.path, contents, policy)) {
      errors.push(`${entry.path}: ${message}`);
    }
  }
  return errors;
}

function main() {
  const errors = validateRepository();
  if (errors.length > 0) {
    for (const error of errors) console.error(error);
    process.exitCode = 1;
    return;
  }
  console.log('Public boundary validation passed.');
}

if (process.argv[1] && import.meta.url === pathToFileURL(path.resolve(process.argv[1])).href) {
  main();
}
