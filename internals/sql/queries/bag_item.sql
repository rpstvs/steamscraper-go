-- name: AddItem :one
INSERT INTO Bag_item (Bag_id, Item_id, Amount)
VALUES ($1, $2, $3)
RETURNING *;