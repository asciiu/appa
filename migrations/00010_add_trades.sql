-- +goose Up
CREATE TABLE trades(
  taker_order_id UUID REFERENCES orders (id) ON DELETE CASCADE,
  maker_order_id UUID REFERENCES orders (id) ON DELETE CASCADE,
  amount BIGINT NOT NULL,
  price BIGINT NOT NULL,
  side VARCHAR(4) NOT NULL,
  created_on TIMESTAMP DEFAULT now(),
  updated_on TIMESTAMP DEFAULT current_timestamp
);

-- +goose Down
DROP TABLE trades;
