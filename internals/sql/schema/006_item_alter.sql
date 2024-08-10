-- +goose Up
ALTER TABLE Items
ADD COLUMN ImageUrl TEXT NOT NULL;
-- +goose Down
ALTER TABLE Items DROP COLUMN ImageUrl;