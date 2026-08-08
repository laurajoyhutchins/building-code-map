CREATE TEMP TABLE IF NOT EXISTS bcm_audit(metric VARCHAR, value BIGINT);
CREATE OR REPLACE TEMP TABLE street_ranges_source AS
SELECT * FROM read_csv_auto({{SOURCE}}, header = true, all_varchar = true);
CREATE OR REPLACE TEMP TABLE street_ranges_ranked AS
SELECT
  trim(source_record_id) AS source_record_id,
  trim(from_number) AS from_number,
  trim(to_number) AS to_number,
  upper(trim(parity)) AS parity,
  trim(street) AS street,
  trim(city) AS city,
  upper(trim(state)) AS state,
  trim(postal_code) AS postal_code,
  trim(from_longitude) AS from_longitude,
  trim(from_latitude) AS from_latitude,
  trim(to_longitude) AS to_longitude,
  trim(to_latitude) AS to_latitude,
  row_number() OVER (
    PARTITION BY trim(source_record_id)
    ORDER BY trim(from_number), trim(to_number), trim(street), trim(city), upper(trim(state))
  ) AS duplicate_rank,
  (
    trim(source_record_id) <> '' AND trim(street) <> '' AND trim(city) <> '' AND
    length(upper(trim(state))) = 2 AND upper(trim(parity)) IN ('B', 'E', 'O') AND
    try_cast(trim(from_number) AS BIGINT) IS NOT NULL AND
    try_cast(trim(to_number) AS BIGINT) IS NOT NULL AND
    try_cast(trim(from_number) AS BIGINT) <> try_cast(trim(to_number) AS BIGINT) AND
    try_cast(trim(from_longitude) AS DOUBLE) BETWEEN -180 AND 180 AND
    try_cast(trim(to_longitude) AS DOUBLE) BETWEEN -180 AND 180 AND
    try_cast(trim(from_latitude) AS DOUBLE) BETWEEN -90 AND 90 AND
    try_cast(trim(to_latitude) AS DOUBLE) BETWEEN -90 AND 90
  ) AS valid_record
FROM street_ranges_source;
COPY (
  SELECT source_record_id, from_number, to_number, parity, street, city, state,
         postal_code, from_longitude, from_latitude, to_longitude, to_latitude
  FROM street_ranges_ranked
  WHERE valid_record AND duplicate_rank = 1
  ORDER BY source_record_id, street, city, state, from_number, to_number
) TO {{OUTPUT}} (header, delimiter ',');
INSERT INTO bcm_audit
SELECT 'accepted', count(*) FROM street_ranges_ranked WHERE valid_record AND duplicate_rank = 1
UNION ALL SELECT 'rejected', count(*) FROM street_ranges_ranked WHERE NOT valid_record
UNION ALL SELECT 'duplicate', count(*) FROM street_ranges_ranked WHERE duplicate_rank > 1
UNION ALL SELECT 'quarantined', 0;
