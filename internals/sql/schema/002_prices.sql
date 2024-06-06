-- +goose Up
CREATE TABLE Prices (
    PriceDate TIMESTAMP,
    Item_id INTEGER NOT NULL REFERENCES Items(id),
    Price FLOAT,
    PRIMARY KEY(Item_id, PriceDate)
);
-- +goose Down
DROP TABLE Prices;