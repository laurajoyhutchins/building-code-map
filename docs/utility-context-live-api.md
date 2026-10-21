# Live EIA county utility-candidate API

`cmd/bcm-utility-api` is deliberately **not** the BCM regulatory backend. It
serves exactly the source-pinned `GET /v1/utility-context?state=CO&county=Denver`
and `GET /healthz`. It offers no address-to-county inference, confirmed service
territory, or building-code verdict. A county must be established independently.

The immutable derivative is retained under the private GCS bucket and supplied
to the Docker build by a main-only, WIF-authenticated engineering-infra workflow.
Do not commit `private-view/` to Git or publish the full derivative to Pages.
The container embeds a byte-for-byte copy after independent SHA-256 verification;
it requires no runtime Google Cloud credentials or GCS access. Its digest is
pinned in the Dockerfile and checked again at startup.

A public Cloud Run transport for this separately scoped EIA data is **not**
equivalent to making the entire BCM application or its regulatory corpus live.
Bound costs with min instances 0 and max instances 1; monitor usage. Its
candidate lookup is subject to the limitations in [utility context](utility-context.md).
