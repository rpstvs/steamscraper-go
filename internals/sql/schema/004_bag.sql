-- +goose Up
CREATE TABLE Bag(
    Id UUID UNIQUE PRIMARY KEY,
    User_id UUID NOT NULL REFERENCES Users(Id),
    TotalValue DECIMAL(10, 2) NOT NULL,
    Created_at TIMESTAMP NOT NULL,
    Updated_at TIMESTAMP NOT NULL,
    Name TEXT NOT NULL
);
-- +goose Down
DROP TABLE Bag;