# Live project lookup

BCM's public website and API are separate deployables. The public website is hosted from this repository on GitHub Pages. It can call a separately hosted Go service at `POST /v1/project-code-basis` without bundling the private regulatory dataset.

## Source of the answer

The public Go service uses `engine.NewProjectVerifier` for the project verdict. Its optional `privateProvider` adapter calls the existing private BCM Authority Engine at `POST /v1/resolve`. The adapter transfers only jurisdiction identity, code adoption names, source IDs, unresolved local requirements, and content-identified runtime provenance. It does not return the private corpus or stored legal text.

This adapter is deliberately conservative: it does not promote a legacy resolved adoption to a fully verified project claim. Unavailable capabilities, mismatched bundle identities, incomplete geography, and upstream errors fail closed. Project facts are retained in the request but marked as requiring review if the legacy resolver has not evaluated them.

## Deployment inputs

The public Go service supports a configurable remote provider. Without configuration, `go run ./cmd/bcm` remains a fictional DEMO-XX reference app. To run against private BCM, set:

- `BCM_PROVIDER_URL`: an HTTPS origin for the private Authority Engine service, accessible to the Go service but **not** exposed by the browser.
- `BCM_PROVIDER_BUNDLE_DIGEST`: the precise digest from a separately reviewed, non-demo bundle loaded by that private Authority Engine. A digest mismatch prevents live results.
- `BCM_PUBLIC_ORIGIN`: the public website origin, normally `https://laurajoyhutchins.github.io`. Browser CORS is restricted to this origin.
- `BCM_PROVIDER_AUTH_MODE`: required for the live Go service. Use `gcp_id_token` when the private Authority Engine is protected by Cloud Run IAM, or `bearer` only when a separate provider explicitly validates bearer tokens.
- `BCM_PROVIDER_TOKEN`: server-side secret required only in `bearer` mode; prohibited in `gcp_id_token` mode. Never place credentials in GitHub Pages variables.

For Cloud Run, grant the **public API service account** `roles/run.invoker` on the private backend service only. The Go API obtains a short-lived Google-signed identity token from the Cloud Run metadata server with the private service origin as the audience. No service-account key file is required or permitted. Keep the private backend unauthenticated-invoker permission disabled.
- `PORT`: listen port, defaults to `8080` when the live provider is configured.

The server offers `GET /v1/readiness` for Pages to check the private service's address-lookup capability and expected bundle identity. `GET /health` only establishes process liveness. Neither endpoint implies nationwide coverage.

A multi-stage Dockerfile builds the same public Go binary for a private-backed server deployment. Deploy behind authentication and appropriate rate limits until the production data, access policy, error handling, and public exposure have been independently reviewed. Do **not** deploy the checked-in `building-code-map-alpha/backend/bundle.json` as production coverage. It is explicitly a demo fixture.

## Enable the public Pages client

Once the Go API service has been deployed and tested with the reviewed private dataset, set the **public repository variable** `BCM_PUBLIC_API_BASE_URL` to that API's HTTPS origin (no path or credentials). GitHub Pages builds a live-mode browser interface and removes the synthetic sample result file. If the variable is unset, Pages continues serving its fictional sample and makes no live claims.

Live-mode requests go to `POST /v1/project-code-basis` with `project_id`, `address`, `applicability_date`, and optional project facts. The browser validates the versioned public result, matches its requested date and project identity, requires a content-identified bundle, and never accepts DEMO-XX as a real jurisdiction. Changing inputs cancels stale requests.

## Release gate

Before setting `BCM_PUBLIC_API_BASE_URL` or permitting unauthenticated access:

1. Independently verify the production private bundle's component digests, source licenses and distribution constraints, as-of dates, and actual geocoder/boundary coverage.
2. Deploy the private Authority Engine behind a server-to-server authentication boundary, without serving its bulk data endpoints publicly.
3. Verify `GET /v1/readiness`, `GET /v1/bundle`, and a representative real project-scoped `POST /v1/resolve`, including unsupported and ambiguous cases.
4. Deploy the public Go API with a pinned bundle digest, restricted browser origin, server-to-server authentication, rate limits, monitoring and rollback. Verify that `POST /v1/project-code-basis` preserves local-record and evidence defects.
5. Configure Pages only after the public API returns accurate project results under the contract. Independently read back the deployed revision and sample failure cases.

**Current boundary:** This repository contains the live wiring and tests, not an independently confirmed production regulatory deployment. Never interpret a green CI check or a valid synthetic example as live service availability.

## Optional non-regulatory utility context

The live Go backend may independently load a SHA-256-pinned county-candidate
projection of the Electric Utility Catalog from a private local path. Its
`GET /v1/utility-context` operation requires separately established state
and county, never geocodes an address, and never changes the code-basis result.
Configuration, offline projection rules and the serving-utility disclaimer are
in [Utility context](utility-context.md). The static demo remains synthetic;
this route is not automatically deployed by a successful source PR.
