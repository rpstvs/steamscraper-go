-- +goose Up
CREATE TABLE Users(
    Id UUID PRIMARY KEY UNIQUE,
    Name TEXT NOT NULL,
    SteamID TEXT NOT NULL UNIQUE,
    Created_at TIMESTAMP NOT NULL,
    Updated_at TIMESTAMP NOT NULL
);
-- +goose Down
DROP TABLE Users;