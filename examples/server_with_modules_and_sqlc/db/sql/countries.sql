CREATE TABLE IF NOT EXISTS countries (
  -- init
  id            BIGSERIAL         PRIMARY KEY,
  created_at    TIMESTAMP         NOT NULL DEFAULT NOW(),
  updated_at    TIMESTAMP         NOT NULL DEFAULT NOW()

  -- new columns below

);

-- when the row is updated, update the "updated_at" timestamp
CREATE TRIGGER set_timestamp BEFORE UPDATE ON countries
FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();


-- name: GetCountry :one
SELECT      id
FROM        countries
WHERE       id = $1 LIMIT 1;

-- name: GetCountries :many
SELECT      id
FROM        countries;

-- name: CreateCountry :one
INSERT INTO countries ( column_name ) 
VALUES      ( $1 )
RETURNING   column_name;

-- name: DeleteCountry :exec
DELETE FROM countries
WHERE       id = $1;

-- name: UpdateCountry :one
UPDATE      countries
SET         column_name = $1
WHERE       column_name = $2
RETURNING   column_name;