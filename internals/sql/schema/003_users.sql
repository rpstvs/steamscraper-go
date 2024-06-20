-- +goose Up
CREATE TABLE Users (Id UUID PRIMARY KEY, Bag UUID);
-- +goose Down
DROP TABLE Users;