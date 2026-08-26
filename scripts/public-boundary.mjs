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

export function validateRepositoryPaths(topLevelPaths, policy) {
  const errors = [];
  const forbidden = new Set(policy?.forbidden_top_level_paths ?? []);

  for (const entry of topLevelPaths) {
    if (forbidden.has(entry)) {
      errors.push(`forbidden top-level path: ${entry}`);
    }
  }

  return errors;
}

function collectJsonFiles(directory) {
  if (!fs.existsSync(directory)) return [];
  const files = [];

  for (const entry of fs.readdirSync(directory, { withFileTypes: true })) {
    const fullPath = path.join(directory, entry.name);
    if (entry.isDirectory()) {
      files.push(...collectJsonFiles(fullPath));
    } else if (entry.isFile() && entry.name.endsWith('.json')) {
      files.push(fullPath);
    }
  }

  return files;
}

export function validateRepository(rootDir = process.cwd()) {
  const policyPath = path.join(rootDir, 'policy', 'public-boundary.json');
  const policy = JSON.parse(fs.readFileSync(policyPath, 'utf8'));
  const errors = [...validatePolicy(policy)];

  const topLevelPaths = fs.readdirSync(rootDir, { withFileTypes: true }).map((entry) => entry.name);
  errors.push(...validateRepositoryPaths(topLevelPaths, policy));

  for (const fixturePath of collectJsonFiles(path.join(rootDir, 'demo'))) {
    const fixture = JSON.parse(fs.readFileSync(fixturePath, 'utf8'));
    errors.push(...validateSyntheticFixture(fixture, policy).map(
      (error) => `${path.relative(rootDir, fixturePath)}: ${error}`,
    ));
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
