-- +goose Up
DROP TABLE orders;
CREATE TABLE orders(
  id UUID PRIMARY KEY, 
  user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
  market_name VARCHAR NOT NULL,
  side VARCHAR(4) NOT NULL,
  amount BIGINT NOT NULL,
  filled BIGINT NOT NULL DEFAULT 0,
  price BIGINT NOT NULL,
  type VARCHAR NOT NULL,
  status VARCHAR NOT NULL,
  created_on TIMESTAMP DEFAULT now(),
  updated_on TIMESTAMP DEFAULT current_timestamp
);

-- +goose Down
DROP TABLE orders;
CREATE TABLE orders(
  id UUID PRIMARY KEY, 
  user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
  market_name VARCHAR NOT NULL,
  side VARCHAR(4) NOT NULL,
  size decimal NOT NULL,
  fill decimal NOT NULL,
  price decimal,
  type VARCHAR NOT NULL,
  status VARCHAR NOT NULL,
  created_on TIMESTAMP DEFAULT now(),
  updated_on TIMESTAMP DEFAULT current_timestamp
);