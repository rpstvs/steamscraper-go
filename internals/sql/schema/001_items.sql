-- +goose Up
CREATE TABLE Items (
    Id UUID PRIMARY KEY,
    ItemName TEXT UNIQUE NOT NULL,
    Classid TEXT NOT NULL,
    DayChange DECIMAL(10, 2) NOT NULL,
    WeekChange DECIMAL(10, 2) NOT NULL
);
-- +goose Down
DROP TABLE Items;