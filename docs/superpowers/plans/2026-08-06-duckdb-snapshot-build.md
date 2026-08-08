# DuckDB snapshot build lane implementation plan

1. Add a `snapshotbuild` package with explicit tool, source, and output contracts.
2. Embed versioned DuckDB SQL for geocoder and boundary normalization.
3. Reject implicit extension installation and remote executable inputs.
4. Verify exact DuckDB and spatial-extension identities before execution.
5. Add deterministic audit reports and build receipts.
6. Reuse the existing geocoder SQLite builder.
7. Add a canonical boundary SQLite finalizer and validate its output through the runtime loader.
8. Extend `bcm` with `snapshot audit` and `snapshot build`.
9. Generate existing PR #45 component manifests without adding a competing schema.
10. Verify core pipeline behavior with fake executable tests and deterministic boundary-output tests.
11. Insert the PR after website migration PR #54 and before offline release PR #55.
12. Advance PR #55 onto the new branch, then regenerate release-bound identities separately.
