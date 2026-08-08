CREATE TEMP TABLE IF NOT EXISTS bcm_audit(metric VARCHAR, value BIGINT);
CREATE OR REPLACE TEMP TABLE layer_families_source AS
SELECT * FROM read_csv_auto({{LAYER_SOURCE}}, header = true, all_varchar = true);
COPY (
  SELECT trim(key) AS key, trim(label) AS label, trim(martin_layer_id) AS martin_layer_id,
         trim(description) AS description, trim(default_enabled) AS default_enabled
  FROM layer_families_source
  WHERE trim(key) <> '' AND trim(label) <> ''
  ORDER BY key
) TO {{LAYER_OUTPUT}} (header, delimiter ',');

CREATE OR REPLACE TEMP TABLE refresh_status_source AS
SELECT * FROM read_csv_auto({{REFRESH_SOURCE}}, header = true, all_varchar = true);
COPY (
  SELECT trim(status) AS status,
         trim(latest_successful_refresh) AS latest_successful_refresh,
         trim(latest_attempt) AS latest_attempt,
         trim(next_scheduled_refresh) AS next_scheduled_refresh,
         trim(message) AS message
  FROM refresh_status_source
  ORDER BY status, latest_successful_refresh
) TO {{REFRESH_OUTPUT}} (header, delimiter ',');

CREATE OR REPLACE TEMP TABLE boundary_features_source AS
SELECT * FROM read_csv_auto({{FEATURE_SOURCE}}, header = true, all_varchar = true);
CREATE OR REPLACE TEMP TABLE boundary_features_ranked AS
SELECT
  trim(layer_family) AS layer_family,
  trim(feature_id) AS feature_id,
  trim(title) AS title,
  trim(subtitle) AS subtitle,
  trim(source_id) AS source_id,
  trim(geometry_label) AS geometry_label,
  trim(geometry_source) AS geometry_source,
  trim(last_synced_at) AS last_synced_at,
  trim(geometry_json) AS geometry_json,
  trim(attributes_json) AS attributes_json,
  row_number() OVER (
    PARTITION BY trim(layer_family), trim(feature_id)
    ORDER BY trim(source_id), trim(last_synced_at), trim(geometry_json)
  ) AS duplicate_rank,
  (
    trim(layer_family) <> '' AND trim(feature_id) <> '' AND trim(title) <> '' AND
    trim(source_id) <> '' AND trim(geometry_label) <> '' AND
    try_cast(trim(last_synced_at) AS TIMESTAMP) IS NOT NULL AND
    try_cast(trim(attributes_json) AS JSON) IS NOT NULL
  ) AS valid_record,
  {{GEOMETRY_VALID}} AS geometry_valid
FROM boundary_features_source;
COPY (
  SELECT layer_family, feature_id, title, subtitle, source_id, geometry_label,
         geometry_source, last_synced_at, geometry_json, attributes_json
  FROM boundary_features_ranked
  WHERE valid_record AND geometry_valid AND duplicate_rank = 1
  ORDER BY layer_family, feature_id
) TO {{FEATURE_OUTPUT}} (header, delimiter ',');
INSERT INTO bcm_audit
SELECT 'accepted', count(*) FROM boundary_features_ranked
 WHERE valid_record AND geometry_valid AND duplicate_rank = 1
UNION ALL SELECT 'rejected', count(*) FROM boundary_features_ranked WHERE NOT valid_record
UNION ALL SELECT 'duplicate', count(*) FROM boundary_features_ranked WHERE duplicate_rank > 1
UNION ALL SELECT 'quarantined', count(*) FROM boundary_features_ranked
 WHERE valid_record AND NOT geometry_valid;
