package geocoder

const schemaVersion = "1"

const schemaSQL = `
PRAGMA foreign_keys = ON;

CREATE TABLE geocoder_metadata (
  key TEXT PRIMARY KEY,
  value TEXT NOT NULL
);

CREATE TABLE address_points (
  id INTEGER PRIMARY KEY,
  address_number TEXT NOT NULL,
  street_name TEXT NOT NULL,
  city TEXT NOT NULL,
  state TEXT NOT NULL,
  postal_code TEXT NOT NULL DEFAULT '',
  matched_address TEXT NOT NULL,
  longitude REAL NOT NULL CHECK (longitude >= -180 AND longitude <= 180),
  latitude REAL NOT NULL CHECK (latitude >= -90 AND latitude <= 90),
  source_name TEXT NOT NULL,
  source_record_id TEXT NOT NULL,
  source_vintage TEXT NOT NULL,
  UNIQUE (source_name, source_record_id)
);

CREATE INDEX address_points_lookup
ON address_points (address_number, street_name, state, city, postal_code);

CREATE TABLE street_ranges (
  id INTEGER PRIMARY KEY,
  from_number INTEGER NOT NULL,
  to_number INTEGER NOT NULL,
  parity TEXT NOT NULL CHECK (parity IN ('B', 'E', 'O')),
  street_name TEXT NOT NULL,
  city TEXT NOT NULL,
  state TEXT NOT NULL,
  postal_code TEXT NOT NULL DEFAULT '',
  from_longitude REAL NOT NULL CHECK (from_longitude >= -180 AND from_longitude <= 180),
  from_latitude REAL NOT NULL CHECK (from_latitude >= -90 AND from_latitude <= 90),
  to_longitude REAL NOT NULL CHECK (to_longitude >= -180 AND to_longitude <= 180),
  to_latitude REAL NOT NULL CHECK (to_latitude >= -90 AND to_latitude <= 90),
  source_name TEXT NOT NULL,
  source_record_id TEXT NOT NULL,
  source_vintage TEXT NOT NULL,
  UNIQUE (source_name, source_record_id)
);

CREATE INDEX street_ranges_lookup
ON street_ranges (street_name, state, city, postal_code, from_number, to_number);
`
