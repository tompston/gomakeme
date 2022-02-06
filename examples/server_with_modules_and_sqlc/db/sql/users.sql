CREATE TABLE IF NOT EXISTS users (
  -- init
  id            BIGSERIAL         PRIMARY KEY,
  created_at    TIMESTAMP         NOT NULL DEFAULT NOW(),
  updated_at    TIMESTAMP         NOT NULL DEFAULT NOW()

  -- new columns below

);

-- when the row is updated, update the "updated_at" timestamp
CREATE TRIGGER set_timestamp BEFORE UPDATE ON users
FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();


-- name: GetUser :one
SELECT      id
FROM        users
WHERE       id = $1 LIMIT 1;

-- name: GetUsers :many
SELECT      id
FROM        users;

-- name: CreateUser :one
INSERT INTO users ( column_name ) 
VALUES      ( $1 )
RETURNING   column_name;

-- name: DeleteUser :exec
DELETE FROM users
WHERE       id = $1;

-- name: UpdateUser :one
UPDATE      users
SET         column_name = $1
WHERE       column_name = $2
RETURNING   column_name;