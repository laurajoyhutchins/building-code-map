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
  if (relativePath === 'package.json' || relativePath === 'package-lock.json' || relativePath === 'policy/public-boundary.json') {
    return 'first-party';
  }
  // Reviewed first-party algorithm configuration, not a GIS or regulatory corpus.
  // Permit only this one exact file, never a whole geocoder data directory.
  if (relativePath === 'geocoder/config/geocoder-ranking-policy.json') {
    return 'geocoder-ranking-policy';
  }
  if (relativePath === 'schemas/code-family-coverage-ledger.odcs.json') {
    // Exactly one first-party contract definition. Do not permit arbitrary
    // ODCS-shaped source records or private regulatory corpus in public Git.
    return 'family-ledger-contract';
  }
  if (relativePath === 'schemas/site-authority-policy.odcs.json') {
    // Reviewed public contract definition only, never a real policy or private source.
    return 'site-authority-contract';
  }
  if (relativePath === 'schemas/bulk-query.odcs.json') {
    return 'bulk-query-contract';
  }
  if (relativePath === 'schemas/prov-graph.odcs.json') {
    // A single, reviewed first-party interchange definition, never a blanket
    // permission to commit ODCS-shaped regulatory data.
    return 'prov-contract';
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
  if (disposition === 'geocoder-ranking-policy') {
    const requiredNumbers = [
      'address_point_base', 'street_range_base', 'exact_street', 'exact_city',
      'exact_postal_code', 'postal_code_not_supplied', 'parity_matched',
      'minimum_address_point_quality', 'minimum_street_range_quality', 'ambiguity_gap',
    ];
    const expectedKeys = new Set(['version', 'source_priority', ...requiredNumbers]);
    const keys = value && typeof value === 'object' && !Array.isArray(value)
      ? Object.keys(value) : [];
    if (keys.length !== expectedKeys.size || keys.some((key) => !expectedKeys.has(key))) {
      errors.push('geocoder ranking policy contains unreviewed fields');
    }
    if (typeof value?.version !== 'string' || !value.version.startsWith('geocoder-ranking-')) {
      errors.push('geocoder ranking version must be declared');
    }
    for (const field of requiredNumbers) {
      if (typeof value?.[field] !== 'number' || !Number.isFinite(value[field]) ||
          value[field] < 0 || value[field] > 1) {
        errors.push(`geocoder ranking ${field} must be a bounded number`);
      }
    }
    if (value?.source_priority === null || typeof value?.source_priority !== 'object' ||
        Array.isArray(value?.source_priority) ||
        Object.keys(value.source_priority ?? {}).length !== 0) {
      errors.push('public geocoder ranking must not embed real source priorities');
    }
  }
  if (disposition === 'schema') {
    if (typeof value?.$schema !== 'string' || !value.$schema.includes('json-schema.org')) {
      errors.push('a public schema path must contain an actual JSON Schema');
    }
  }
  if (disposition === 'family-ledger-contract') {
    const schemaUrl = 'https://github.com/laurajoyhutchins/building-code-map/blob/main/schemas/code-family-coverage-ledger.schema.json';
    if (value?.apiVersion !== 'v3.1.0' || value?.kind !== 'DataContract' ||
        value?.id !== 'bcm-code-family-coverage-ledger-v01' ||
        value?.status !== 'draft' ||
        value?.schema?.[0]?.name !== 'code_family_coverage_ledger_v01' ||
        !value?.authoritativeDefinitions?.some((item) => item?.id === 'bcm_family_ledger_json_schema' &&
          item?.url === schemaUrl)) {
      errors.push('public family-ledger contract must match reviewed schema and identity');
    }
  }
  if (disposition === 'site-authority-contract') {
    if (value?.apiVersion !== 'v3.1.0' || value?.kind !== 'DataContract' ||
        value?.id !== 'bcm-site-jurisdiction-authority-policy-v1' ||
        value?.status !== 'draft' ||
        value?.schema?.[0]?.name !== 'site_authority_policy_v1' ||
        !value?.schema?.[0]?.properties?.some((item) => item?.name === 'rules[].legal.instrument_id') ||
        !value?.schema?.[0]?.properties?.some((item) => item?.name === 'rules[].provenance.review_receipt_sha256')) {
      errors.push('public site-authority ODCS must match reviewed schema and contract-only identity');
    }
  }
  if (disposition === 'prov-contract') {
    if (value?.apiVersion !== 'v3.1.0' || value?.kind !== 'DataContract' ||
        value?.id !== 'bcm-prov-o-evidence-graph-v1' ||
        value?.schema?.[0]?.name !== 'prov_o_evidence_graph' ||
        !value?.authoritativeDefinitions?.some(
          (item) => item?.url === 'https://github.com/laurajoyhutchins/building-code-map/blob/main/schemas/prov-graph.schema.json',
        )) {
      errors.push('public PROV contract must match the reviewed schema and identity');
    }
  }
  if (disposition === 'bulk-query-contract') {
    const schemaRoot = 'https://github.com/laurajoyhutchins/building-code-map/blob/main/schemas/';
    const expected = ['bulk-workload-v1.schema.json', 'bulk-result-v1.schema.json',
      'bulk-summary-v1.schema.json', 'live-bulk-canary.schema.json',
      'bulk-boundary-summary-v1.schema.json',
      'bulk-temporal-events-v1.schema.json', 'bulk-temporal-summary-v1.schema.json',
      'bulk-source-gap-summary-v1.schema.json', 'accuracy-review-queue-v1.schema.json'];
    if (value?.apiVersion !== 'v3.1.0' || value?.kind !== 'DataContract' ||
        value?.id !== 'bcm-bulk-query-v1' || value?.status !== 'draft' ||
        value?.schema?.length !== 9 ||
        !expected.every((name) => value?.authoritativeDefinitions?.some((entry) => entry?.url === schemaRoot + name))) {
      errors.push('public bulk query contract must match reviewed schema-only identity');
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
