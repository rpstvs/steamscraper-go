-- name: CreateUser :one
INSERT INTO Users (Id, Name, SteamID)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetUserbyId :one
SELECT *
FROM Users
WHERE SteamID = $1;