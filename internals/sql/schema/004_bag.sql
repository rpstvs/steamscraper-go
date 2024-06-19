-- +goose Up
CREATE TABLE Bag (
    Id UUID PRIMARY KEY,
    User_id UUID NOT NULL,
    Item_id UUID NOT NULL,
    TotalValue DECIMAL(10, 2) NOT NULL,
);
-- +goose Down
DROP TABLE Bag;