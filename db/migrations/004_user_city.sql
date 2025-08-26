-- +goose Up
ALTER TABLE users
ADD COLUMN city TEXT NOT NULL DEFAULT "";

-- +goose Down
ALTER TABLE users
DROP COLUMN city;
