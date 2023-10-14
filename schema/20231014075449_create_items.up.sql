-- Extensions ----------------------------------------------------
CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;
COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';

-- Table Definition ----------------------------------------------
CREATE TABLE IF NOT EXISTS items (
    id integer PRIMARY KEY,
    name text
);