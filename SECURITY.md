# Security Policy

## Supported Version

Security fixes are applied to the current `main` branch. This project has not yet published a stable release series.

## Reporting a Vulnerability

Please report suspected vulnerabilities privately through GitHub's **Report a vulnerability** / private security advisory feature for this repository.

Do not open a public issue containing credentials, exploit details, private data, or a working proof of concept. Include:

- the affected file, endpoint, workflow, or data path;
- the impact and the conditions required to reproduce it;
- a minimal proof of concept when safe to share privately;
- any suggested mitigation;
- whether the issue may affect deployed instances or only local development.

You should receive an acknowledgment after the report is reviewed. Valid reports will be tracked privately until a fix or coordinated disclosure plan is ready.

## Data Accuracy and Provenance

Incorrect building-code adoption facts, authority routing, dates, or source provenance are data-quality issues rather than software vulnerabilities. They may be reported through a normal issue as long as the report contains no private information. Please cite the controlling official source and identify the affected record or report field.

## Secrets and Local Data

The repository must not contain credentials, private environment files, hydrated SQLite or DuckDB snapshots, department-owned NERIS operational data, generated decision-graph exports containing private prompts, or machine-specific home-directory paths.