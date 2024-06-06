-- name: CreateItem :one
INSERT INTO Items (Id, ItemName)
VALUES ($1, $2)
RETURNING *;