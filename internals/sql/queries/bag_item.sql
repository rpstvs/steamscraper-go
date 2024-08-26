-- name: AddItem :one
INSERT INTO Bag_item (Bag_id, item_classid, Amount)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetBagItem :one
SELECT Amount
From Bag_item
WHERE BAG_id = $1
    AND item_classid = $2;
-- name: UpdateBagItem :one
UPDATE Bag_item
SET Amount = $2
WHERE Bag_id = $1
    AND item_classid = $3
RETURNING *;
-- name: DeleteItem :one
DELETE FROM Bag_item
WHERE Bag_id = $1
    AND item_classid = $2
RETURNING *;