# Independent BCAST consumer evidence (Go)

BCM now has a bounded read-only Go consumer for the public `bcast.package/0.2.0` and `bcast.api/0.2.0` contracts in `bcastconsumer/`. It uses Go's standard library plus the upstream `golang.org/x/text/unicode/norm` NFC validator (pinned at Go 1.25-compatible `v0.41.0`). It knows no eCFR XML, provider IDs, private Python compiler, or source access credentials.

The consumer downloads the **entire package**, checks its SHA-256 materialization ID, publication and object identities, rejects unknown fields, duplicates, unexpected JSON types, missing parents and cycles, then checks the API metadata/object/child endpoints against the verified package. The source-safe fixture is copied from public `bcast/examples/synthetic/valid/package-0.2.0.json` for cross-repository interoperability regression. Ordinary BCM CI runs the Go tests entirely offline through a local HTTP server.

This is a read-only source-structure dependency, **not an adoption integration**. A resolved BCAST object is not an operative adopted code, jurisdiction verdict, professional approval, or BCM evidence-link role. In particular, the existing BCM `engine.ProjectVerifier` and its jurisdiction/adoption proof remain authoritative. Do not let this client set project code status or `ExactEvidence`.

Boundary: publication edition/revision coordinates now admit canonical NFC Unicode and reject non-NFC inputs without changing their identity bytes, using the maintained upstream Go normalization library. An exact source-safe Unicode fixture from public BCAST `8273a20` is retained under `bcastconsumer/testdata/` and exercised end-to-end over the local HTTP consumer path. This does not prove live private BCAST and independently deployed BCM communicate in production. That remains separate acceptance evidence. The public BCAST contract, not this client, defines normative semantics.

See [BCAST issue #22](https://github.com/laurajoyhutchins/bcast/issues/22) and [BCAST 1.0 readiness issue #13](https://github.com/laurajoyhutchins/bcast/issues/13).

## Exact-package HTTP interoperability probe

The standalone Go command verifies **all** public BCAST `0.2.0` GET operations for a caller-selected package, without importing the Python BCAST compiler or using any eCFR-specific coordinates:

```sh
go run ./cmd/bcast-contract-probe \\
  -base-url http://127.0.0.1:8000 \\
  -package-id 'bcastpkg:sha256:EXACT_64_DIGIT_CONTENT_DIGEST'
```

It downloads and independently verifies the complete content-addressed BCAST materialization, checks publication and object IDs and parentage, and compares the public HTTP metadata, object and direct-child results with the verified package. Only the package identifier and object count are printed on success; no actual regulatory prose is logged. This command is an interoperability diagnostic, never a way to approve a jurisdictional code adoption. It requires an already available, authorized BCAST read-only API endpoint and does not create or deploy one.
