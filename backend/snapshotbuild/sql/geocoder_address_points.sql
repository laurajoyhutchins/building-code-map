CREATE TEMP TABLE IF NOT EXISTS bcm_audit(metric VARCHAR, value BIGINT);
CREATE OR REPLACE TEMP TABLE address_points_source AS
SELECT * FROM read_csv_auto({{SOURCE}}, header = true, all_varchar = true);
CREATE OR REPLACE TEMP TABLE address_points_ranked AS
SELECT
  trim(source_record_id) AS source_record_id,
  trim(address_number) AS address_number,
  trim(street) AS street,
  trim(city) AS city,
  upper(trim(state)) AS state,
  trim(postal_code) AS postal_code,
  trim(longitude) AS longitude,
  trim(latitude) AS latitude,
  trim(matched_address) AS matched_address,
  row_number() OVER (
    PARTITION BY trim(source_record_id)
    ORDER BY trim(matched_address), trim(longitude), trim(latitude)
  ) AS duplicate_rank,
  (
    trim(source_record_id) <> '' AND trim(address_number) <> '' AND
    trim(street) <> '' AND trim(city) <> '' AND length(upper(trim(state))) = 2 AND
    trim(matched_address) <> '' AND
    try_cast(trim(longitude) AS DOUBLE) BETWEEN -180 AND 180 AND
    try_cast(trim(latitude) AS DOUBLE) BETWEEN -90 AND 90
  ) AS valid_record
FROM address_points_source;
COPY (
  SELECT source_record_id, address_number, street, city, state, postal_code,
         longitude, latitude, matched_address
  FROM address_points_ranked
  WHERE valid_record AND duplicate_rank = 1
  ORDER BY source_record_id, address_number, street, city, state, postal_code
) TO {{OUTPUT}} (header, delimiter ',');
INSERT INTO bcm_audit
SELECT 'accepted', count(*) FROM address_points_ranked WHERE valid_record AND duplicate_rank = 1
UNION ALL SELECT 'rejected', count(*) FROM address_points_ranked WHERE NOT valid_record
UNION ALL SELECT 'duplicate', count(*) FROM address_points_ranked WHERE duplicate_rank > 1
UNION ALL SELECT 'quarantined', 0;
