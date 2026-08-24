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

export function validateSyntheticFixture(fixture, policy) {
  const errors = [];
  const prefix = policy?.synthetic_jurisdiction_prefix;
  const jurisdictionId = fixture?.jurisdiction_id;

  if (typeof jurisdictionId !== 'string' || !jurisdictionId.startsWith(prefix)) {
    errors.push(`fixture jurisdiction_id must use synthetic jurisdiction prefix ${prefix}`);
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
