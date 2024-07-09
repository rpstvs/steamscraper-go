-- name: CreateBag :one
INSERT INTO Bag (Id, TotalValue, User_id)
VALUES ($1, $2, $3)
RETURNING *;
-- name: UpdateBag :one
UPDATE Bag
SET TotalValue = $2
WHERE Id = $1
RETURNING *;
-- name: GetBagbyID :one
SELECT *
FROM Bag
WHERE Id = $1;