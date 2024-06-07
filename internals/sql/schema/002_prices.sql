-- +goose Up
CREATE TABLE Prices (
    PriceDate TIMESTAMP NOT NULL,
    Item_id INTEGER NOT NULL REFERENCES Items(Id),
    Price DECIMAL(10, 2) NOT NULL,
    PRIMARY KEY(Item_id, PriceDate)
);
-- +goose Down
DROP TABLE Prices;