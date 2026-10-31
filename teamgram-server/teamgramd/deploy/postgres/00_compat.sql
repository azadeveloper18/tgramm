-- Modified by Kgramm, 2026-10-31.
-- Derived from teamgram-server SQL, which remains under the Apache License 2.0.
-- MySQL dialect (ENGINE, AUTO_INCREMENT counters, charset, ON DUPLICATE KEY) is not preserved.


CREATE EXTENSION IF NOT EXISTS citext;

-- Flag columns are boolean because the Go data objects scan bool.
-- PostgreSQL already has an explicit integer→boolean cast and refuses to drop it,
-- so "deleted = 0" does not compile. DAO statements use false and true instead.
DROP FUNCTION IF EXISTS kgramm_int_to_bool(integer);

-- pgx sends Go strings as text. JSON columns are jsonb.
DO $$ BEGIN
  CREATE CAST (text AS jsonb) WITH INOUT AS IMPLICIT;
EXCEPTION
  WHEN duplicate_object THEN NULL;
END $$;
