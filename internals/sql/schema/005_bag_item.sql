-- +goose Up
CREATE TABLE Bag_item (
    Bag_id UUID NOT NULL REFERENCES Bag(Id),
    Item_id UUID NOT NULL REFERENCES Items(Id),
    Amount INTEGER NOT NULL,
    PRIMARY KEY(Item_id, Bag_id),
    UNIQUE(Item_id, Bag_id)
);
-- +goose Down
DROP TABLE Bag_item;