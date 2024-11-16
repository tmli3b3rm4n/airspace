-- Create schema if not exists
CREATE SCHEMA IF NOT EXISTS public;

-- Create the flight_restrictions table
CREATE TABLE IF NOT EXISTS flight_restrictions (
                                                   id SERIAL PRIMARY KEY,
                                                   properties JSONB,                   -- For storing GeoJSON properties
                                                   geom geometry(Geometry, 4326)      -- For storing geometries (Polygon, MultiPolygon, etc.)
);

-- Create a spatial index on the geometry column
CREATE INDEX IF NOT EXISTS idx_flight_restrictions_geom
    ON flight_restrictions USING GIST (geom);