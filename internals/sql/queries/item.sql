-- name: CreateItem :one
INSERT INTO Items (id, ItemName)
VALUES ($1, $2)
RETURNING *;