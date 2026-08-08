# Authority Engine Release Verification

Date: 2026-08-07 (America/Denver)

## Exact identities

- Integrated source commit: `f19664c296ebbc061bf1a836af9609ba20b44c3e`
- Source-tree CI run: `31242936869`
- Generated package commit: `10352f8bedd582bfcdc8edc71d1c47f6627ad5e6`
- Release workflow run: `31242934381`
- Release workflow conclusion: `success`
- Go toolchain recorded by the package: `go1.25.0 linux/amd64`
- Cold-room mode recorded by the package: `fallback`

The source commit is intentionally distinct from the generated package commit. The release package records the exact integrated source tree from which its executable, snapshots, manifests, schemas, checksums, and verification record were generated. The package commit contains generated artifacts and is not reinterpreted as its own source.

The bot-authored generated package commit caused GitHub to classify its pull-request CI run `31242976117` as `action_required` without creating jobs. That state is not test evidence. This human-authored review commit exists to trigger the complete ordinary repository gate against the final package tree.

## Integrated source disposition

The source commit is based on `main` after the corrected authority-engine stack and DuckDB snapshot-build lane were merged predecessor-first. It therefore contains the integrated behavior from PRs #56, #50, #51, #52, #53, #54, and #57 together with the previously landed snapshot-manifest, frontend runtime-decoder, geocoder provenance, and Issue #7 boundary-detail work.

The descendant branches were reconstructed from semantic implementation deltas rather than stale ancestry trees. That repair preserved the current manifest preflight, frontend cancellation and typed error boundary, deterministic geocoder ranking/provenance, compact Explorer boundary payload, and on-demand feature-detail behavior.

Source-tree CI run `31242936869` passed at exact source commit `f19664c296ebbc061bf1a836af9609ba20b44c3e` before package regeneration.

## Release regeneration evidence

Release workflow `31242934381` completed successfully against exact source commit `f19664c296ebbc061bf1a836af9609ba20b44c3e`. The generated verification manifest records successful checks for:

1. Go tests and vet;
2. deterministic Linux AMD64 binary construction;
3. geocoder rebuild and byte reproducibility;
4. geocoder manifest reproducibility;
5. bundle inspection;
6. CLI point resolution and address geocoding;
7. HTTP readiness and resolution;
8. MCP initialization;
9. offline verification; and
10. cold-room verification.

The generated package binds its geocoder source vintage, bundle identity, checksum manifest, and verification manifest to the same exact source commit.

## Package identity

- `bin/bcm`: `ac968ba2581f68833e30f56f2e1fc81dfd15f370d383de4ac32f00f1126f0ea9`
- `data/authority.snapshot.json`: `2feb7751a351a84b9b84887bf90991988d0bed0da293563a40b0b55f6e0e2a6f`
- `data/geocoder.sqlite`: `8e8224190e12e37c62c3917efbeb8fab3cb50b205d2a7ab34a0aa2834e713efe`
- `data/geocoder.sqlite.manifest.json`: `cfeba51ce9ad26789ec33054c3e6a69e68f973832b6d92a31bce544f85e038da`
- recursive regulatory catalog: `cc0dc65fcecf69d344ccfbf2212d88c542a86317adf62eb4ebaf7b29995a88ac`
- `manifests/bundle.json`: `d0d7ef58a8e2e9f5e73c40b1938ff050f9aa7af78fbf160799e402da534fa669`
- `manifests/verification.json`: `344814c0cf4cd0cf598434377f79ef734e5d7a766d13939a1a96c42ac2906047`

The checksum manifest covers the release README, executable, verification scripts, boundary and geocoder data, geocoder snapshot manifest, all included regulatory profiles and rule packs, bundle/coverage/verification manifests, and engine schemas. The checksum manifest itself is excluded to avoid recursive self-hashing.

## Coverage

The bundle declares regulatory coverage for Colorado, Florida, New Jersey, North Carolina, Oregon, and Virginia across the building, electrical, energy, fire-operational, and residential code families as represented by the included profiles. DuckDB remains a build-time snapshot normalization and analytical QA tool. The packaged runtime continues to use canonical SQLite snapshots and the Go authority engine.

## Cold-room limitation

The hosted runner could not create a private network namespace. The cold-room script therefore used its explicit `network-namespace-fallback` path and ran commands with a cleared environment containing only bounded `PATH` and `HOME` values. The package records `cold_room_mode: fallback` and does not claim kernel-enforced network isolation.

Within that stated limitation, the copied package passed checksum verification, bundle inspection, point resolution, address geocoding, address resolution, HTTP readiness/resolution, and MCP initialization without acquiring package dependencies or source data during those verification commands.

## Final disposition

The regenerated package is correctly bound to the integrated source commit, and the release workflow completed all declared reproducibility and executable verification checks. PR #55 is mergeable only after the human-authored commit containing this review passes the complete ordinary repository CI gate at its exact head.
