-- name: GetCommand :many
SELECT * FROM command;

-- name: GetCommandById :one
SELECT * FROM command WHERE id = $1;

-- name: CreateCommand :execresult
INSERT INTO command (command, message, target, args) VALUES ($1, $2, $3, $4);