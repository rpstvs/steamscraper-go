-- +goose Up
CREATE TABLE Item (
    Id INTEGER UNIQUE NOT NULL,
    ItemName TEXT NOT NULL,
    Condition TEXT,
    Price FLOAT,
);
-- +goose Down
DROP TABLE users;