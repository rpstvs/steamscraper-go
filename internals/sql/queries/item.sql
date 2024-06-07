-- name: CreateItem :one
INSERT INTO Items (id, ItemName)
VALUES ($1, $2)
RETURNING *;
-- name: GetItemsIds :many
SELECT Id
FROM Items;
-- name: GetItemByName :one
SELECT Id
FROM Items
WHERE itemname = $1;