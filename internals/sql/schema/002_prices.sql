-- +goose Up
CREATE TABLE Prices (
    PriceDate DATE NOT NULL,
    Item_id UUID NOT NULL REFERENCES Items(Id),
    Price DECIMAL(10, 2) NOT NULL,
    PRIMARY KEY(Item_id, PriceDate),
    UNIQUE(Item_id, PriceDate)
);
-- +goose Down
DROP TABLE Prices;