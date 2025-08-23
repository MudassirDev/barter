-- +goose Up
CREATE TABLE skills (
  id UUID PRIMARY KEY,
  title TEXT NOT NULL UNIQUE,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE skills;
