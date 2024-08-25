-- +goose Up
CREATE TABLE Prices (
    PriceDate DATE NOT NULL,
    item_classid BIGINT NOT NULL REFERENCES Items(classid),
    Price DECIMAL(10, 2) NOT NULL,
    PRIMARY KEY(item_classid, PriceDate),
    UNIQUE(PriceDate)
);
-- +goose Down
DROP TABLE Prices;