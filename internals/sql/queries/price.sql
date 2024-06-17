-- name: AddPrice :many
INSERT INTO Prices (PriceDate, Item_id, Price)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetPricebyId :one
SELECT *
FROM Prices
WHERE Item_id = $1
ORDER BY PriceDate DESC;