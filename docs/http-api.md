# Project Verification HTTP API

Building Code Map exposes the canonical project-verification workflow through a deliberately small HTTP adapter. The adapter is a transport boundary around `engine.ProjectVerifier`; it does not contain a second verdict reducer, evidence model, or regulatory authority path.

## Endpoint

`POST /v1/project-code-basis`

Request bodies use the public `engine.ProjectRequest` shape:

```json
{
  "project_id": "DEMO-HTTP",
  "address": "100 Demo Plaza",
  "applicability_date": "2026-08-25",
  "project_type": "tenant-improvement",
  "facts": {
    "occupancy": "business"
  }
}
```

A caller may use supported coordinates instead of an address. A caller does not select a code family before project verification.

Successful responses are the canonical `engine.ProjectCodeBasis` JSON representation. The HTTP layer does not rewrite project verdicts, code states, exact evidence, unresolved requirements, facts, or provenance.

## Request boundary

The adapter fails closed at the transport boundary:

- only `POST` is accepted; other methods return `405` with `Allow: POST`;
- request bodies are limited to 64 KiB;
- JSON decoding rejects unknown fields;
- malformed JSON and trailing JSON values are rejected;
- when `Content-Type` is supplied it must resolve to `application/json`.

These checks happen before the canonical project verifier executes.

## Error envelope

Transport and engine failures use one bounded envelope:

```json
{
  "error": {
    "code": "invalid_query",
    "message": "applicability_date is required for engine queries",
    "retryable": false
  }
}
```

Typed `engine.EngineError` codes survive the HTTP boundary. Current status mapping is:

| Engine condition | HTTP status |
| --- | ---: |
| invalid query or coordinates | 400 |
| address or regulatory profile not found | 404 |
| ambiguous boundary/address or unsupported coverage | 422 |
| geocoder or regulatory catalog unavailable | 503 |
| invalid provider data or internal failure | 500 |

Transport-only codes include `invalid_json`, `request_too_large`, `unsupported_media_type`, `method_not_allowed`, and `not_found`.

## Dependency injection

`httpapi.NewHandler` accepts a narrow verifier interface with the same `VerifyProject` signature as `engine.ProjectVerifier`. Production hosting can inject the canonical verifier backed by its provider. Public tests use synthetic providers only.

The public repository therefore defines and tests the HTTP contract without bundling maintained regulatory data, production geocoder or boundary stores, restricted source bytes, or deployment secrets.
