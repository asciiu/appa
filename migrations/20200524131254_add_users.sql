-- +goose Up
-- SQL in this section is executed when the migration is applied.
-- USERS
CREATE TABLE users(
    id uuid PRIMARY KEY,
    email VARCHAR UNIQUE NOT NULL,
    email_verified BOOLEAN NOT NULL DEFAULT false,
    username VARCHAR UNIQUE NOT NULL, 
    avatar_url VARCHAR,
    password_hash VARCHAR NOT NULL,
    created_on TIMESTAMP DEFAULT now(),
    updated_on TIMESTAMP DEFAULT current_timestamp,
    deleted_on TIMESTAMP
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE users;