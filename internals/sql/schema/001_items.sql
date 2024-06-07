-- +goose Up
CREATE TABLE Items (
    Id UUID PRIMARY KEY,
    ItemName TEXT UNIQUE NOT NULL
);
-- +goose Down
DROP TABLE Items;