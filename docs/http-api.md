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

## Request correlation

Every response from the public handler includes one effective `X-Request-ID`.

A caller-supplied request ID is preserved only when it is 1 through 128 ASCII characters, begins with an alphanumeric character, and otherwise contains only alphanumeric characters plus `.`, `_`, `:`, or `-`. Missing or invalid values are replaced with an opaque generated ID.

The correlation layer emits a bounded completion record after the request finishes:

```json
{
  "request_id": "req_0123456789abcdef0123456789abcdef",
  "method": "POST",
  "route_class": "project_code_verification",
  "status": 200,
  "duration_ms": 4,
  "runtime_identity": {
    "engine_version": "example",
    "source_commit": "example-commit",
    "bundle_id": "example-bundle"
  }
}
```

The completion record is intentionally not an access log. It never contains the civic address, raw request body, project facts, source excerpts, response body, query parameters, or raw unmatched path. Unknown methods collapse to `OTHER`; unmatched paths collapse to `unmatched`.

`runtime_identity` appears only when the canonical successful `ProjectCodeBasis.provenance` already contains that identity. The middleware receives it through request context from the canonical project handler rather than parsing or copying the response body.

Completion recording is non-authoritative. A sink error or panic cannot change the HTTP result. `JSONLineSink` is a small standard-library sink for one-record-per-line structured output; hosted runtimes may inject another implementation of the narrow `CompletionSink` interface.

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

`httpapi.NewHandler` accepts the narrow verifier interface with the same `VerifyProject` signature as `engine.ProjectVerifier` and an optional `CompletionSink`. Production hosting can inject the canonical verifier backed by its provider and its desired bounded completion sink. Public tests use synthetic providers only.

The public repository therefore defines and tests the HTTP and correlation contracts without bundling maintained regulatory data, production geocoder or boundary stores, restricted source bytes, deployment secrets, or a generalized observability stack.
