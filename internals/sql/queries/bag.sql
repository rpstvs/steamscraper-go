-- name: CreateBag :one
INSERT INTO Bag (Id, TotalValue, User_id, Created_at, Updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: UpdateBag :one
UPDATE Bag
SET TotalValue = $2
    AND Updated_at = $3
WHERE Id = $1
RETURNING *;
-- name: GetBagbyID :one
SELECT *
FROM Bag
WHERE Id = $1;
-- name: GetBagsByUser :many
SELECT *
FROM Bag
Where User_id = $1;