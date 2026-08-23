# Methodology

Building Code Map treats regulatory resolution as an evidence problem, not a string-matching problem.

## Evidence before claims

A result should identify the authority, effective date, code family or rule, and evidence supporting the conclusion. The system should preserve unresolved conditions rather than silently convert uncertainty into a confident answer.

## Fail closed

When the available evidence cannot establish applicability, authority, timing, or a project-specific condition, the result remains conditional or not verified. Missing evidence does not prove non-applicability.

## Project facts matter

A project answer may depend on facts such as location, permit or application date, occupancy, construction type, scope of work, and local administrative conditions. Those facts are inputs to the resolution process rather than after-the-fact annotations.

## Time matters

Regulatory conclusions are evaluated for a stated date. Future effective dates, transition periods, superseded rules, and historical conditions must not be flattened into a single timeless value.

## Public and production data

The public distribution provides schemas, methodology, and synthetic examples. The maintained production regulatory knowledge base is a separate service/data layer and is not required for understanding the public contracts.

## Synthetic fixtures

Public examples use reserved fictional identifiers such as `DEMO-XX`. They are designed to exercise interfaces and uncertainty semantics without reproducing a real jurisdiction's maintained regulatory record.
