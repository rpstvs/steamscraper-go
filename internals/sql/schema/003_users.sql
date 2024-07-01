-- +goose Up
CREATE TABLE Users (
    Id UUID PRIMARY KEY,
    Name TEXT NOT NULL,
    SteamID TEXT NOT NULL UNIQUE
);
-- +goose Down
DROP TABLE Users;