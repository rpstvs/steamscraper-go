-- +goose Up
CREATE TABLE Bag (
    Id UUID PRIMARY KEY,
    TotalValue DECIMAL(10, 2) NOT NULL
);
-- +goose Down
DROP TABLE Bag;