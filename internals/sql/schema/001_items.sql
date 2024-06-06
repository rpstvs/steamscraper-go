-- +goose Up
CREATE TABLE Items (
    Id INTEGER UNIQUE NOT NULL PRIMARY KEY,
    ItemName TEXT NOT NULL
);
-- +goose Down
DROP TABLE Item;