# Authority Engine Release Verification

Date: 2026-08-06

## Exact identities

- DuckDB build-lane PR: #57
- Verified DuckDB build-lane head: `e4157ea7ec25bc064540f0b71a48835360c0e237`
- DuckDB build-lane CI run: `31104689886`
- Integrated release source commit: `29a5e083750bdeabacdb6cabf0f259273fcdc9a8`
- Source-tree CI run: `31104779790`
- Generated package commit: `9cbb263943f2a8b4b9dedbff5dceb4c8761504b1`
- Release workflow run: `31104776515`
- Go toolchain: `go1.25.0 linux/amd64`
- Cold-room network mode: `fallback`

The source commit is intentionally distinct from the generated package commit. The release package records the exact integrated source tree from which its binary, geocoder snapshot, snapshot manifest, bundle manifests, checksums, and verification receipt were generated. The package commit contains generated artifacts and is not reinterpreted as its own source.

The bot-authored package commit caused GitHub to classify its pull-request workflow as `action_required` without creating jobs. This is a workflow-approval state, not test evidence. The human-authored commit containing this review exists to trigger the complete ordinary repository gate against the final package tree.

## Stack integration evidence

The integrated source includes the complete predecessor stack through the DuckDB snapshot build lane:

1. PR #45 snapshot manifests and activation identity: `3814252a4674cc7de8bd9e8a99aa482bdcc3af26`, CI `31100571518` passed;
2. PR #47 frontend runtime decoders: `0eee12cada8a425dde12bfe6eb1fedd1ab652177`, CI `31100915409` passed;
3. PR #48 geocoder ranking and provenance: `10e04f0ec6733fb5bebaada34b7e5330d42565d3`, CI `31100940991` passed;
4. PR #56 authority-engine contracts: `adf6fc62db7162451b0bc3b3a5ae9dc8619f1e9b`, CI `31100969079` passed;
5. PR #50 reusable engine core: `6f0aa1f1dce7184065bb10839657f22c09f4ca57`, CI `31100997341` passed;
6. PR #51 public packages and HTTP transport: `d6a95f9e59a93b603d5795c236d3832f257fd604`, CI `31101027999` passed;
7. PR #52 CLI and content-addressed bundles: `5a1caf97a668ef2ed357e02305bf69a20d2e28e5`, CI `31101048480` passed;
8. PR #53 MCP adapter: `bd3b541860c323493355820fa157b7a8ccba5475`, CI `31101071571` passed;
9. PR #54 website engine v1 migration: `995aca99d34831fac007ad0e11f871650713f8fa`, CI `31101098125` passed;
10. PR #57 DuckDB snapshot build and audit lane: `e4157ea7ec25bc064540f0b71a48835360c0e237`, CI `31104689886` passed.

PR #55 targets PR #57. The integrated source commit preserves the prior release history and the verified DuckDB branch through two-parent merge commits rather than rebases or force-pushes.

## DuckDB build-lane disposition

PR #57 adds:

- `bcm snapshot audit` and `bcm snapshot build`;
- embedded, versioned DuckDB normalization contracts;
- exact DuckDB executable version and SHA-256 verification;
- exact spatial-extension SHA-256 verification for boundary builds;
- rejection of implicit extension installation and remote executable inputs;
- bounded disposable execution state;
- deterministic accepted, rejected, duplicate, and quarantined counts;
- canonical Go finalization into the existing SQLite runtime formats;
- PR #45 snapshot-manifest finalization;
- build receipts binding tools, extensions, SQL contracts, sources, counts, and outputs.

DuckDB remains a build-time analytical workspace. SQLite remains the runtime snapshot format. Go remains authoritative for final schemas, validation, activation, geocoder ranking and ambiguity, regulatory meaning, and transport behavior.

The repository gate verifies compilation, unit behavior, deterministic boundary SQLite output, contract safety, identity binding, and stack compatibility. It does not claim that the release package contains DuckDB or that a real DuckDB-plus-spatial capsule was executed during release regeneration. That executable pilot remains separate evidence to add when an exact provisioned tool capsule is available.

## Reproducible release evidence

Release workflow `31104776515` completed all of the following against exact source commit `29a5e083750bdeabacdb6cabf0f259273fcdc9a8`:

1. ran `go test ./...`;
2. ran `go vet ./...`;
3. built `bin/bcm` with `CGO_ENABLED=0`, `GOOS=linux`, `GOARCH=amd64`, `-trimpath`, and an empty Go build ID;
4. generated a canonical geocoder manifest template bound to the exact source revision and committed fixture digests;
5. rebuilt the SQLite geocoder twice from the committed address-point and street-range fixtures;
6. compared both geocoder databases byte-for-byte;
7. compared both finalized snapshot manifests byte-for-byte;
8. regenerated the boundary snapshot, six-state regulatory catalog, engine schemas, coverage, bundle manifest, verification receipt, and complete checksum manifest;
9. verified the prepared package with `building-code-engine/verify-offline.sh`;
10. copied the package into a temporary cold-room directory and verified its checksums again;
11. exercised bundle inspection, point resolution, address geocoding, address resolution, HTTP readiness, HTTP resolution, and MCP initialization in the cold room;
12. finalized the verification receipt only after those checks passed;
13. reran both offline and cold-room verification against the finalized receipt;
14. committed the generated package only after the second verification pass succeeded.

The runner could not create a private network namespace, so cold-room verification used the explicit `network-namespace-fallback` path. Commands still ran with a cleared environment containing only bounded `PATH` and `HOME` values. The receipt records `cold_room_mode: fallback`; it does not claim kernel network isolation.

## Snapshot-manifest identity

The generated geocoder manifest records:

- snapshot ID: `geocoder-release-29a5e083750b`;
- builder revision: `29a5e083750bdeabacdb6cabf0f259273fcdc9a8`;
- source fixture digest: `635ae6a70342b720ce319c3d4fef5a9e9f7c3f0f9e3cec3c654bbcac2d175177`;
- accepted records: `5`;
- rejected, duplicate, and quarantined records: `0`;
- output size: `36864` bytes;
- output SHA-256: `2dec9b4dca9c07b8ba4b925dbd98ef707b61fec9a520d88bff77e3daed086ab9`.

## Package identity

- `bin/bcm`: `f7b7559463e89aed3f96510a2e563ae43332c7a3daada43e86a8efbb8df9de5f`
- `data/geocoder.sqlite`: `2dec9b4dca9c07b8ba4b925dbd98ef707b61fec9a520d88bff77e3daed086ab9`
- `data/geocoder.sqlite.manifest.json`: `e83ddb1c1cb94cbbcdf7e8aa7f105c1a601287a9ddc4117d1c06ead4c525bfd2`
- boundary snapshot: `2feb7751a351a84b9b84887bf90991988d0bed0da293563a40b0b55f6e0e2a6f`
- recursive regulatory catalog: `cc0dc65fcecf69d344ccfbf2212d88c542a86317adf62eb4ebaf7b29995a88ac`
- bundle manifest: `0efd2f530af3a403a605ced4f3846f9cf87cca773da97344069797fbaf44b510`
- verification receipt: `a01b6ceb7c431d865b0244b0766d339ba70dc65bcf6a0b717fd5dc47baa8c6a7`

The checksum manifest covers the release README, executable, verification scripts, both package data views, the geocoder snapshot manifest in both views, all six regulatory profiles and rule packs, manifests, verification receipt, and engine schemas. The checksum manifest itself is excluded to avoid recursive self-hashing.

## Coverage

The package declares executable regulatory coverage for Colorado, Florida, New Jersey, North Carolina, Oregon, and Virginia across building, residential, electrical, energy, and fire-operational code families as applicable to the included profiles.

## Disposition

DuckDB build-lane implementation, stack propagation, package regeneration, snapshot-manifest identity, and executable offline/cold-room verification are complete for source commit `29a5e083750bdeabacdb6cabf0f259273fcdc9a8`.

The human-authored commit containing this evidence must pass the complete repository CI gate before PR #55 can claim final exact-head clearance. PRs #57 and #55 remain draft and unmerged pending dependency-order integration and independent review.
