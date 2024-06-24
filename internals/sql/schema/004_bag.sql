-- +goose Up
CREATE TABLE Bag (
    Id UUID PRIMARY KEY,
    Item_id UUID [],
    TotalValue DECIMAL(10, 2) NOT NULL
);
-- +goose Down
DROP TABLE Bag;