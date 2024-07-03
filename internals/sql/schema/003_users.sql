-- +goose Up
CREATE TABLE Users (
    Id UUID PRIMARY KEY UNIQUE,
    Name TEXT NOT NULL,
    SteamID TEXT NOT NULL UNIQUE,
    Bag Bag []
);
-- +goose Down
DROP TABLE Users;