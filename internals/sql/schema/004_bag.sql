-- +goose Up
CREATE TABLE Bag (
    Id UUID PRIMARY KEY,
    User_id UUID NOT NULL REFERENCES Users(Id),
    TotalValue DECIMAL(10, 2) NOT NULL
);
-- +goose Down
DROP TABLE Bag;