-- +goose Up
CREATE TABLE currencies (
  symbol VARCHAR PRIMARY KEY,
  name VARCHAR NOT NULL,
  precision float NOT NULL DEFAULT 0.000000000001,
  created_on TIMESTAMP DEFAULT now(),
  updated_on TIMESTAMP DEFAULT current_timestamp
);

CREATE TABLE balances (
  id UUID PRIMARY KEY, 
  user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
  symbol VARCHAR NOT NULL REFERENCES currencies (symbol),
  amount BIGINT NOT NULL DEFAULT 0,
  locked BIGINT NOT NULL DEFAULT 0,       -- this is the amount that is not available either because it is locked in an order or used in a game
  address VARCHAR NOT NULL,
  created_on TIMESTAMP DEFAULT now(),
  updated_on TIMESTAMP DEFAULT current_timestamp
);

INSERT into currencies values ('BTC', 'Bitcoin');
INSERT into currencies values ('LTC', 'Litecoin');

-- +goose Down
DROP TABLE balances;
DROP TABLE currencies; 