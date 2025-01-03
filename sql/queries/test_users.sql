-- name: CreateUser :one
INSERT INTO test_user (id, created_at, updated_at, name)
VALUES (?, ?, ?, ?)
RETURNING *;
