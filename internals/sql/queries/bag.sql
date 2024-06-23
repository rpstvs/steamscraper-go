-- name: CreateBag :one
INSERT INTO Bag (Id, Item_id, TotalValue)
VALUES ($1, $2, $3)
RETURNING *;
-- name: UpdateBag :one
UPDATE Bag
SET Item_id = $2,
    TotalValue = $3
WHERE Id = $1
RETURNING *;
-- name: GetBagbyID :one
SELECT *
FROM Bag
WHERE Id = $1;