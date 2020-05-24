-- +goose Up
-- SQL in this section is executed when the migration is applied.
DROP TABLE stories;
CREATE TABLE stories(
  id UUID PRIMARY KEY, 
  author_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
  title VARCHAR NOT NULL,
  content JSONB NOT NULL,
  status VARCHAR NOT NULL,
  created_on TIMESTAMP DEFAULT now(),
  updated_on TIMESTAMP DEFAULT current_timestamp
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE stories;
CREATE TABLE stories(
  id UUID PRIMARY KEY, 
  author_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
  title VARCHAR NOT NULL,
  content VARCHAR NOT NULL,
  status VARCHAR NOT NULL,
  created_on TIMESTAMP DEFAULT now(),
  updated_on TIMESTAMP DEFAULT current_timestamp
);

