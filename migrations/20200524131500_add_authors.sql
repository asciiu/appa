-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE writers (
  user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
  title VARCHAR NOT NULL
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE writers;
