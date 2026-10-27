# Colorado address-point acquisition (candidate)

The State of Colorado publishes the statewide Colorado Public Address Composite through the official ArcGIS FeatureServer:
https://gis.colorado.gov/public/rest/services/Address_and_Parcel/Colorado_Public_Addresses/FeatureServer/0

The public importer at `scripts/ingest_colorado_addresses.py` extracts a stable, paginated CSV with object identities for use by BCM's existing private Go geocoder snapshot builder. It does **not** add a second geocoder, use a runtime web lookup, or place the downloaded address corpus in this repository.

## Source and terms

The State of Colorado OIT GIS Team aggregates records supplied by counties and municipalities and updates the state layer periodically. The server permits commercial use with attribution, but explicitly prohibits resale. See the service's published copyright text and terms. The output receipt includes the published attribution statement, actual retrieval date, source URL, row counts, and SHA-256.

The provider itself warns that the dataset can be incomplete or stale. An address point supports a geographic observation only, not proof of the local building-code authority or edition.

## Acquire a candidate

Run outside Git or in a private ignored working directory:

```bash
python3 scripts/ingest_colorado_addresses.py \
  --output /private/bcm/candidates/colorado-addresses.csv \
  --retrieved-at YYYY-MM-DD
```

## Receipt-bound publication and recovery

A CSV filename alone is **not** an admitted acquisition. Independently check the
CSV's exact SHA-256 against its adjacent `<csv>.receipt.json` immediately before
promoting bytes to the geocoder snapshot builder:

```bash
python3 scripts/ingest_colorado_addresses.py \\
  --output /private/bcm/candidates/colorado-addresses.csv \\
  --verify-existing
```

The importer stages the receipt before any CSV replacement and rolls back an
existing verified pair on ordinary filesystem write/rename failures. A partial
or mismatched existing pair is refused rather than silently overwritten. Neither
the CSV nor the receipt is itself regulatory authority or production approval.

**Two filenames cannot be replaced as one filesystem transaction.** A process
crash or storage failure between renames can leave a pair incomplete. In this
state, verification must fail closed. The importer retains prior files, when
available, under a sibling hidden `.<csv-name>.recovery.*` directory. Stop
writers, inspect any retained pair and its recorded SHA-256 under protected
custody, restore only a matching pair with the original `output_path`, and
rerun `--verify-existing`. Never infer success merely from a renamed file,
a successful download, or an old receipt. Do not run a release consumer across
a concurrent replacement without retaining and rechecking the exact bytes it
uses.

Default fail-closed behavior refuses any invalid record, duplicate/nonmonotonic ID, missing page, implausible coordinate, or excessive source count. A narrowly reviewed exception for invalid source rows may specify `--max-invalid`, with the full rejection count retained in the receipt. The importer does not publish an approved runtime bundle and cannot set the public site's production flag.

Pass the CSV to the existing verified Go `geocoder-build` in protected custody, using the recorded vintage and a reviewed snapshot manifest template. Check the output manifest, data license, accepted/rejected counts, regional coverage, temporal validity, regulatory catalog, boundary geometry, capability ledger, and unknown/ambiguous test cases before constructing an immutable runtime bundle.

Do not publish these address records to GitHub Pages or a public Git repository. Keep the state attribution accessible in the application documentation. No current state, county, city, or address should be declared a verified code determination simply because geocoding succeeded.
