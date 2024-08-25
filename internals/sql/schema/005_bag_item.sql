-- +goose Up
CREATE TABLE Bag_item (
    Bag_id UUID NOT NULL REFERENCES Bag(Id),
    item_classid BIGINT NOT NULL REFERENCES Items(classid),
    Amount INTEGER NOT NULL,
    PRIMARY KEY(item_classid, Bag_id),
    UNIQUE(item_classid, Bag_id)
);
-- +goose Down
DROP TABLE Bag_item;