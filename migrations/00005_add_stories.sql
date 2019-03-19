-- +goose Up
CREATE TABLE stories(
  id UUID PRIMARY KEY, 
  author_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
  title VARCHAR NOT NULL,
  content VARCHAR NOT NULL,
  status VARCHAR NOT NULL,
  created_on TIMESTAMP DEFAULT now(),
  updated_on TIMESTAMP DEFAULT current_timestamp
);

-- +goose Down
DROP TABLE stories;