-- name: CreateUser :one
INSERT INTO Users (Id, Name, SteamID, Created_at, Updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: GetUserbyId :one
SELECT *
FROM Users
WHERE SteamID = $1;