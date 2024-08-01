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
-- name: UpdateDailyChange :exec
UPDATE Items
SET DayChange = $1
WHERE Id = $2;
-- name: UpdateWeeklyChange :exec
UPDATE Items
SET WeekChange = $1
WHERE Id = $2;