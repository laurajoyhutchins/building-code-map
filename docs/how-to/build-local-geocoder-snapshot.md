# Build a Local Geocoder Snapshot

Use this procedure to turn reviewed address-point and street-range CSV files into the SQLite snapshot consumed by the Building Code Map backend.

The builder does not download data. Acquire and review source files separately, confirm their redistribution terms, and retain the source name and vintage used for the build.

## Prepare address points

Address-point CSV files require these columns:

```text
source_record_id,address_number,street,city,state,postal_code,longitude,latitude,matched_address
```

Example:

```csv
source_record_id,address_number,street,city,state,postal_code,longitude,latitude,matched_address
co-denver-1600,1600,N BROADWAY ST,DENVER,CO,80202,-104.9876,39.7411,1600 N Broadway St Denver CO 80202
```

Requirements:

- `source_record_id` must be stable and unique within the named source.
- Coordinates must be WGS 84 longitude and latitude.
- `matched_address` is display text. The builder independently normalizes comparison fields.
- `state` may be a state name or postal abbreviation.
- ZIP codes are optional, but the column must be present.

## Prepare street ranges

Street-range CSV files require these columns:

```text
source_record_id,from_number,to_number,parity,street,city,state,postal_code,from_longitude,from_latitude,to_longitude,to_latitude
```

`parity` must be:

- `B` for both odd and even numbers;
- `E` for even numbers;
- `O` for odd numbers.

Street ranges are a fallback. The runtime selects an authoritative address point before considering interpolation.

## Build the snapshot

Run from `backend/`:

```bash
go run ./cmd/geocoder-build \
  --output data/geocoder.sqlite \
  --address-points ./path/to/address-points.csv \
  --street-ranges ./path/to/street-ranges.csv \
  --source-name "Reviewed local address data" \
  --source-vintage "2026-08-01"
```

At least one of `--address-points` or `--street-ranges` is required.

The builder creates a temporary database, validates rows, loads all records in one transaction, and replaces the target only after the build succeeds. Duplicate source identities, invalid coordinates, malformed ranges, and missing provenance stop the build.

## Start the backend

The default geocoder path is `backend/data/geocoder.sqlite` when the server is launched from `backend/`.

Use another repository-contained path with either:

```bash
go run ./cmd/server --geocoder-data data/my-geocoder.sqlite
```

or:

```text
GEOCODER_DATA_PATH=data/my-geocoder.sqlite
```

A missing or invalid geocoder snapshot disables address endpoints but does not disable coordinate-based `/resolve` requests.

## Check the result

Request geocoding without regulatory resolution:

```bash
curl -X POST http://127.0.0.1:8000/geocode \
  -H "Content-Type: application/json" \
  -d '{"address":"1600 N Broadway, Denver, CO 80202"}'
```

A selected candidate includes:

- matched address;
- longitude and latitude;
- `address_point` or `interpolated` precision;
- confidence;
- source name;
- source record identifier;
- source vintage.

Use `/lookup` to compose the selected point with regulatory resolution. Ambiguous and unmatched addresses do not proceed to the regulatory resolver.

## Publication boundary

Do not commit the resulting SQLite file. The repository ignores database binaries. Commit reproducible ingestion logic, small synthetic or redistribution-cleared test fixtures, and source documentation instead.
