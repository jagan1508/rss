-- +goose Up
CREATE TABLE users (
    id VARCHAR(40) PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL
);

-- +goose Down
DROP TABLE users;