-- name: AddPrice :many
INSERT INTO Prices (PriceDate, Item_id, Price)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetPricebyId :many
SELECT Price,
    PriceDate
FROM Prices
WHERE Item_id = $1
ORDER BY PriceDate DESC;
-- name: GetLatestPrice :one
SELECT Price,
    Item_id
FROM Prices
WHERE Item_id = $1
ORDER BY PriceDate DESC;
-- name: GetItemRecord :many
Select Price
FROM Prices
WHERE Item_id = $1
ORDER By PriceDate DESC
LIMIT $2;
-- name: GetItemsRecord :many
SELECT Itemname,
    Id,
    CAST (Prices.Price AS NUMERIC(10, 2))
FROM Items
    LEFT JOIN Prices ON Items.Id = Prices.Item_id
WHERE Itemname = $2
ORDER BY PriceDate DESC
Limit $1;