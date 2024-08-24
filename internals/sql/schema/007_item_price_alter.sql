-- +goose Up
ALTER TABLE Items
ADD COLUMN classid BIGINT NOT NULL UNIQUE;
ALTER TABLE Prices
ADD COLUMN item_classid BIGINT NOT NULL UNIQUE;
-- +goose Down
ALTER TABLE Items DROP COLUMN classid;
ALTER TABLE Prices DROP COLUMN item_classid;