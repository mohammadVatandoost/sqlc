CREATE TABLE foo (a integer, b integer);

-- name: InsertValues :batchone
INSERT INTO foo (a, b)
VALUES ($1, $2)
ON CONFLICT DO NOTHING
RETURNING *;

-- name: GetOne :one
SELECT * FROM foo WHERE a = ? AND b = ? LIMIT 1;

-- name: GetAll :many
SELECT * FROM foo;

-- name: GetAllAByB :many
SELECT a FROM foo WHERE b = ?;
